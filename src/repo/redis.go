package repo

import (
	"log"
	"sync"
	"task1/entity"
	"time"

	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

var (
	redisStrings func(reply interface{}, err error) ([]string, error)
)

type RedisOptions struct {
	Address       string
	MaxActiveConn int
	MaxIdleConn   int
	Timeout       int
}

type Storage struct {
	Pool  entity.Handler
	mutex sync.Mutex
}

func NewRedisRepo(opt *RedisOptions) *Storage {

	storage := &Storage{
		mutex: sync.Mutex{},
		Pool: &redis.Pool{
			MaxActive:   opt.MaxActiveConn,
			MaxIdle:     opt.MaxIdleConn,
			IdleTimeout: time.Duration(opt.Timeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", opt.Address)
				if err != nil {
					log.Default().Print("[Redis Pool]:", err.Error())
				}

				return conn, err
			},
		},
	}

	redisStrings = redis.Strings

	return storage
}

func (storage *Storage) Ping() error {
	conn := storage.Pool.Get()
	_, err := conn.Do("PING")
	return err
}

//HSet set has map
func (storage *Storage) HSet(key, field string, value interface{}) (int64, error) {
	conn := storage.Pool.Get()
	defer conn.Close()
	resp, err := redis.Int64(conn.Do("HSET", key, field, value))
	return resp, err
}

func (storage *Storage) HGetSummary(key string) (summary entity.Summary, err error) {
	conn := storage.Pool.Get()
	defer conn.Close()
	resp, err := redisStrings(conn.Do("HGETALL", key))
	if err != nil {
		err = errors.Wrap(err, "HGetSummary")
		return summary, err
	}
	if len(resp) == 0 {
		err = errors.Wrap(redis.ErrNil, "HGetSummary")
		return
	}

	summary.ConvertFromHGetAllToStruct(resp)

	return summary, nil
}

func (storage Storage) HSetSummary(key string, summary entity.Summary) (int64, error) {
	conn := storage.Pool.Get()
	defer conn.Close()
	resp, err := redis.Int64(conn.Do("HSET", key,
		"previous_price", summary.PreviousPrice,
		"open_price", summary.OpenPrice,
		"highest_price", summary.HighestPrice,
		"lowest_price", summary.LowestPrice,
		"close_price", summary.ClosePrice,
		"volume", summary.Volume,
		"value", summary.Value,
		"is_new_day", summary.IsNewDay))

	return resp, err
}

func (storage *Storage) Del(key string) (int64, error) {
	conn := storage.Pool.Get()
	defer conn.Close()
	resp, err := redis.Int64(conn.Do("DEL", key))
	if err != nil && errors.Is(err, redis.ErrNil) {
		err = nil
	}
	return resp, err
}
