package repo

import (
	"fmt"
	"log"
	"task1/entity"

	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

type RedisOptions struct {
	conn    redis.Conn
	Address string
}

func NewRedisRepo(opt *RedisOptions) *RedisOptions {
	conn, err := redis.Dial("tcp", opt.Address) //"127.0.0.1:6379"
	if err != nil {
		log.Default().Print("[Redis Pool]:", err.Error())
	}

	opt.conn = conn

	return opt
}

func (opt RedisOptions) Ping() error {
	_, err := opt.conn.Do("PING")
	return err
}

func (opt RedisOptions) SetKeyValue(key string, value interface{}) (string, error) {
	defer opt.conn.Close()
	resp, err := redis.String(opt.conn.Do("SET", key, value))
	return resp, err
}

func (opt RedisOptions) Get(key string) (string, error) {

	defer opt.conn.Close()
	resp, err := redis.String(opt.conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", fmt.Errorf("not found")
	}
	return resp, err
}

// HGet key and value
func (opt RedisOptions) HGet(key, field string) (string, error) {

	defer opt.conn.Close()
	return redis.String(opt.conn.Do("HGET", key, field))
}

// HGetAll key and value
func (opt RedisOptions) HGetAll(key string) ([]string, error) {
	defer opt.conn.Close()
	return redis.Strings(opt.conn.Do("HGETALL", key))
}

//HSet set has map
func (opt RedisOptions) HSet(key, field string, value interface{}) (int64, error) {
	defer opt.conn.Close()
	resp, err := redis.Int64(opt.conn.Do("HSET", key, field, value))
	return resp, err
}

func (opt RedisOptions) HGetSummary(key string) (summary entity.Summary, err error) {
	defer opt.conn.Close()
	resp, err := redis.Strings(opt.conn.Do("HGETALL", key))
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

func (opt RedisOptions) HSetSummary(key string, summary entity.Summary) (int64, error) {
	defer opt.conn.Close()
	resp, err := redis.Int64(opt.conn.Do("HSET", key,
		"previous_price", summary.PreviousPrice,
		"open_price", summary.OpenPrice,
		"highest_price", summary.HighestPrice,
		"lowest_price", summary.LowestPrice,
		"close_price", summary.ClosePrice,
		"volume", summary.Volume,
		"value", summary.Value))

	return resp, err
}

func (opt RedisOptions) Del(key string) (int64, error) {
	defer opt.conn.Close()
	resp, err := redis.Int64(opt.conn.Do("DEL", key))
	if err != nil && errors.Is(err, redis.ErrNil) {
		err = nil
	}
	return resp, err
}
