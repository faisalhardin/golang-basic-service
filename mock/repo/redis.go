// Code generated by MockGen. DO NOT EDIT.
// Source: redis.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	context "context"
	reflect "reflect"
	entity "task1/entity"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/gomodule/redigo/redis"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHandler) Get() redis.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(redis.Conn)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockHandlerMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHandler)(nil).Get))
}

// GetContext mocks base method.
func (m *MockHandler) GetContext(arg0 context.Context) (redis.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext", arg0)
	ret0, _ := ret[0].(redis.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContext indicates an expected call of GetContext.
func (mr *MockHandlerMockRecorder) GetContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockHandler)(nil).GetContext), arg0)
}

// MockStorageInterface is a mock of StorageInterface interface.
type MockStorageInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageInterfaceMockRecorder
}

// MockStorageInterfaceMockRecorder is the mock recorder for MockStorageInterface.
type MockStorageInterfaceMockRecorder struct {
	mock *MockStorageInterface
}

// NewMockStorageInterface creates a new mock instance.
func NewMockStorageInterface(ctrl *gomock.Controller) *MockStorageInterface {
	mock := &MockStorageInterface{ctrl: ctrl}
	mock.recorder = &MockStorageInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageInterface) EXPECT() *MockStorageInterfaceMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockStorageInterface) Del(key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del.
func (mr *MockStorageInterfaceMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockStorageInterface)(nil).Del), key)
}

// HGetSummary mocks base method.
func (m *MockStorageInterface) HGetSummary(key string) (entity.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetSummary", key)
	ret0, _ := ret[0].(entity.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGetSummary indicates an expected call of HGetSummary.
func (mr *MockStorageInterfaceMockRecorder) HGetSummary(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetSummary", reflect.TypeOf((*MockStorageInterface)(nil).HGetSummary), key)
}

// HSet mocks base method.
func (m *MockStorageInterface) HSet(key, field string, value interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", key, field, value)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HSet indicates an expected call of HSet.
func (mr *MockStorageInterfaceMockRecorder) HSet(key, field, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockStorageInterface)(nil).HSet), key, field, value)
}

// HSetSummary mocks base method.
func (m *MockStorageInterface) HSetSummary(key string, summary entity.Summary) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSetSummary", key, summary)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HSetSummary indicates an expected call of HSetSummary.
func (mr *MockStorageInterfaceMockRecorder) HSetSummary(key, summary interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSetSummary", reflect.TypeOf((*MockStorageInterface)(nil).HSetSummary), key, summary)
}

// Ping mocks base method.
func (m *MockStorageInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockStorageInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockStorageInterface)(nil).Ping))
}
