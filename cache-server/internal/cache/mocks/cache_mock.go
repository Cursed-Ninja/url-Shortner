// Code generated by MockGen. DO NOT EDIT.
// Source: D:\Programming\Go\url-shortner\cache-server\internal\cache\cache.go

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCacheInterface is a mock of CacheInterface interface.
type MockCacheInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCacheInterfaceMockRecorder
}

// MockCacheInterfaceMockRecorder is the mock recorder for MockCacheInterface.
type MockCacheInterfaceMockRecorder struct {
	mock *MockCacheInterface
}

// NewMockCacheInterface creates a new mock instance.
func NewMockCacheInterface(ctrl *gomock.Controller) *MockCacheInterface {
	mock := &MockCacheInterface{ctrl: ctrl}
	mock.recorder = &MockCacheInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheInterface) EXPECT() *MockCacheInterfaceMockRecorder {
	return m.recorder
}

// GetValue mocks base method.
func (m *MockCacheInterface) GetValue(key, requestId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", key, requestId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValue indicates an expected call of GetValue.
func (mr *MockCacheInterfaceMockRecorder) GetValue(key, requestId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockCacheInterface)(nil).GetValue), key, requestId)
}

// SetValue mocks base method.
func (m *MockCacheInterface) SetValue(key, value, requestId string, expiryTime time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetValue", key, value, requestId, expiryTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetValue indicates an expected call of SetValue.
func (mr *MockCacheInterfaceMockRecorder) SetValue(key, value, requestId, expiryTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValue", reflect.TypeOf((*MockCacheInterface)(nil).SetValue), key, value, requestId, expiryTime)
}