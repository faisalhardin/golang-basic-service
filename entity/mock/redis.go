package mockrepo

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

var GetFunc func() redis.Conn
var GetContextFunc func(context.Context) (redis.Conn, error)

var GetCloseFunc func() error
var GetErrFunc func() error
var GetDoFunc func(commandName string, args ...interface{}) (reply interface{}, err error)
var GetSendFunc func(commandName string, args ...interface{}) error
var GetFlushFunc func() error
var GetReceiveFunc func() (reply interface{}, err error)

type MochHandler struct {
	GetFunc        func() redis.Conn
	GetContextFunc func(context.Context) (redis.Conn, error)
}

func (h *MochHandler) Get() redis.Conn {
	return GetFunc()
}

func (h *MochHandler) GetContext(ctx context.Context) (redis.Conn, error) {
	return GetContextFunc(ctx)
}

type MockRedisConn struct {
	GetCloseFunc   func() error
	GetErrFunc     func() error
	GetDoFunc      func(commandName string, args ...interface{}) (reply interface{}, err error)
	GetSendFunc    func(commandName string, args ...interface{}) error
	GetFlushFunc   func() error
	GetReceiveFunc func() (reply interface{}, err error)
}

func (rc *MockRedisConn) Close() error {
	return GetCloseFunc()
}

func (rc *MockRedisConn) Err() error {
	return GetErrFunc()
}

func (rc *MockRedisConn) Do(c string, args ...interface{}) (reply interface{}, err error) {
	return GetDoFunc(c, args)
}

func (rc *MockRedisConn) Send(c string, args ...interface{}) error {
	return GetSendFunc(c, args...)
}

func (rc *MockRedisConn) Flush() error {
	return GetFlushFunc()
}

func (rc *MockRedisConn) Receive() (reply interface{}, err error) {
	return GetReceiveFunc()
}
