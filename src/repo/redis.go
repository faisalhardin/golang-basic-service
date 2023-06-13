package repo

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Options struct {
	conn    redis.Conn
	Address string
}

func New(opt *Options) *Options {
	conn, err := redis.Dial("tcp", opt.Address) //"127.0.0.1:6379"
	if err != nil {
		log.Default().Print("[Redis Pool]:", err.Error())
	}

	opt.conn = conn

	return opt
}

func (opt Options) Ping() error {
	_, err := opt.conn.Do("PING")
	return err
}

func (opt Options) SetKeyValue(key string, value interface{}) (string, error) {
	defer opt.conn.Close()
	resp, err := redis.String(opt.conn.Do("SET", key, value))
	return resp, err
}

func (opt Options) Get(key string) (string, error) {

	defer opt.conn.Close()
	resp, err := redis.String(opt.conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", fmt.Errorf("not found")
	}
	return resp, err
}

// HGet key and value
func (opt Options) HGet(key, field string) (string, error) {

	defer opt.conn.Close()
	return redis.String(opt.conn.Do("HGET", key, field))
}

// HGetAll key and value
func (opt Options) HGetAll(key string) ([]string, error) {
	defer opt.conn.Close()
	return redis.Strings(opt.conn.Do("HGETALL", key))
}

//HSet set has map
func (opt Options) HSet(key, field string, value interface{}) (int64, error) {
	defer opt.conn.Close()
	resp, err := redis.Int64(opt.conn.Do("HSET", key, field, value))
	return resp, err
}
