// Code generated by MockGen. DO NOT EDIT.
// Source: D:\Programming\Go\url-shortner\kafka-server\internal\database\conn.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDBInterface is a mock of DBInterface interface.
type MockDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBInterfaceMockRecorder
}

// MockDBInterfaceMockRecorder is the mock recorder for MockDBInterface.
type MockDBInterfaceMockRecorder struct {
	mock *MockDBInterface
}

// NewMockDBInterface creates a new mock instance.
func NewMockDBInterface(ctrl *gomock.Controller) *MockDBInterface {
	mock := &MockDBInterface{ctrl: ctrl}
	mock.recorder = &MockDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBInterface) EXPECT() *MockDBInterfaceMockRecorder {
	return m.recorder
}

// InsertOne mocks base method.
func (m *MockDBInterface) InsertOne(document interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", document)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockDBInterfaceMockRecorder) InsertOne(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDBInterface)(nil).InsertOne), document)
}