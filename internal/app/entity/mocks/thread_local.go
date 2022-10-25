// Code generated by MockGen. DO NOT EDIT.
// Source: thread_local.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entity "github.com/dosovma/otus_arch/internal/app/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIThreadLocal is a mock of IThreadLocal interface.
type MockIThreadLocal struct {
	ctrl     *gomock.Controller
	recorder *MockIThreadLocalMockRecorder
}

// MockIThreadLocalMockRecorder is the mock recorder for MockIThreadLocal.
type MockIThreadLocalMockRecorder struct {
	mock *MockIThreadLocal
}

// NewMockIThreadLocal creates a new mock instance.
func NewMockIThreadLocal(ctrl *gomock.Controller) *MockIThreadLocal {
	mock := &MockIThreadLocal{ctrl: ctrl}
	mock.recorder = &MockIThreadLocalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIThreadLocal) EXPECT() *MockIThreadLocalMockRecorder {
	return m.recorder
}

// GetCurrentScope mocks base method.
func (m *MockIThreadLocal) GetCurrentScope() map[string]func(entity.UObject) entity.Executable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentScope")
	ret0, _ := ret[0].(map[string]func(entity.UObject) entity.Executable)
	return ret0
}

// GetCurrentScope indicates an expected call of GetCurrentScope.
func (mr *MockIThreadLocalMockRecorder) GetCurrentScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentScope", reflect.TypeOf((*MockIThreadLocal)(nil).GetCurrentScope))
}

// GetValue mocks base method.
func (m *MockIThreadLocal) GetValue(key string) (map[string]func(entity.UObject) entity.Executable, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", key)
	ret0, _ := ret[0].(map[string]func(entity.UObject) entity.Executable)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValue indicates an expected call of GetValue.
func (mr *MockIThreadLocalMockRecorder) GetValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockIThreadLocal)(nil).GetValue), key)
}

// SetCurrentScope mocks base method.
func (m *MockIThreadLocal) SetCurrentScope(scope map[string]func(entity.UObject) entity.Executable) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentScope", scope)
}

// SetCurrentScope indicates an expected call of SetCurrentScope.
func (mr *MockIThreadLocalMockRecorder) SetCurrentScope(scope interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentScope", reflect.TypeOf((*MockIThreadLocal)(nil).SetCurrentScope), scope)
}

// SetValue mocks base method.
func (m *MockIThreadLocal) SetValue(key string, value map[string]func(entity.UObject) entity.Executable) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValue", key, value)
}

// SetValue indicates an expected call of SetValue.
func (mr *MockIThreadLocalMockRecorder) SetValue(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValue", reflect.TypeOf((*MockIThreadLocal)(nil).SetValue), key, value)
}
