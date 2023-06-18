package entity

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -self_package=src/entity -destination=../_mocks/redis/mock_redis.go -package=redis src/entity Handler
type Handler interface {
	Get() redis.Conn
	GetContext(context.Context) (redis.Conn, error)
}

type StorageInterface interface {
	Ping() error
	HSet(key, field string, value interface{}) (int64, error)
	HGetSummary(key string) (summary Summary, err error)
	HSetSummary(key string, summary Summary) (int64, error)
	Del(key string) (int64, error)
}
