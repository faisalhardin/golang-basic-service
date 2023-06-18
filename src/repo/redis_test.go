package repo

import (
	"context"
	"reflect"
	"task1/entity"
	"testing"

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

func Test_Ping(t *testing.T) {

	mockRedisConn := MockRedisConn{}

	mockRedisHandler := MochHandler{}

	type args struct {
		Address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name: "Successful ping",
			args: args{
				Address: "",
			},
			patch: func() {
				GetDoFunc = func(commandName string, args ...interface{}) (reply interface{}, err error) {
					return nil, nil
				}

				GetFunc = func() redis.Conn {
					return &mockRedisConn
				}

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				Pool: &mockRedisHandler,
			}
			tt.patch()
			err := storage.Ping()
			if (err != nil) != tt.wantErr {
				t.Errorf("red.Ping() err = %v, wantErr = %v", err, tt.wantErr)
			}

		})
	}
}

func Test_HSetSummary(t *testing.T) {
	mockRedisConn := MockRedisConn{}

	mockRedisHandler := MochHandler{}

	mockSummary := entity.Summary{
		PreviousPrice: 8000,
		OpenPrice:     8050,
		HighestPrice:  8100,
		LowestPrice:   7950,
		ClosePrice:    8100,
		Volume:        900,
		Value:         7210000,
		IsNewDay:      0,
	}
	type args struct {
		Address string
		Key     string
		Summary entity.Summary
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    int64
		patch   func()
	}{
		{
			name: "Successful set",
			args: args{
				Address: "",
				Key:     "",
				Summary: mockSummary,
			},
			wantErr: false,
			want: func() int64 {
				numOfFields := reflect.TypeOf(entity.Summary{}).NumField()
				return int64(numOfFields)
			}(),
			patch: func() {
				GetDoFunc = func(commandName string, args ...interface{}) (reply interface{}, err error) {
					return int64(8), nil
				}

				GetCloseFunc = func() error {
					return nil
				}

				GetFunc = func() redis.Conn {
					return &mockRedisConn
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				Pool: &mockRedisHandler,
			}
			tt.patch()
			got, err := storage.HSetSummary(tt.args.Key, tt.args.Summary)
			if (err != nil) != tt.wantErr {
				t.Errorf("storage.HSetAll() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("storage.HSetAll() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_HGetSummary(t *testing.T) {
	mockRedisConn := MockRedisConn{}

	mockRedisHandler := MochHandler{}

	wantSummary := entity.Summary{
		PreviousPrice: 8000,
		OpenPrice:     8050,
		HighestPrice:  8100,
		LowestPrice:   7950,
		ClosePrice:    8100,
		Volume:        900,
		Value:         7210000,
		IsNewDay:      0,
	}
	type args struct {
		Address string
		Key     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    entity.Summary
		patch   func()
	}{
		{
			name: "Successful set",
			args: args{
				Address: "",
				Key:     "",
			},
			want:    wantSummary,
			wantErr: false,
			patch: func() {
				GetDoFunc = func(commandName string, args ...interface{}) (reply interface{}, err error) {
					return nil, nil
				}

				GetCloseFunc = func() error {
					return nil
				}

				GetFunc = func() redis.Conn {
					return &mockRedisConn
				}

				redisStrings = func(reply interface{}, err error) ([]string, error) {
					response := []string{
						"previous_price", "8000",
						"open_price", "8050",
						"highest_price", "8100",
						"lowest_price", "7950",
						"close_price", "8100",
						"volume", "900",
						"value", "7210000",
						"is_new_day", "0",
					}

					return response, nil
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Storage{
				Pool: &mockRedisHandler,
			}

			tt.patch()
			got, err := storage.HGetSummary(tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("storage.HGet() err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("storage.HGet() = %v, want = %v", got, tt.want)
			}
		})
	}
}
