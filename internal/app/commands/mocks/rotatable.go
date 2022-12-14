// Code generated by MockGen. DO NOT EDIT.
// Source: rotatable.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRotatable is a mock of Rotatable interface.
type MockRotatable struct {
	ctrl     *gomock.Controller
	recorder *MockRotatableMockRecorder
}

// MockRotatableMockRecorder is the mock recorder for MockRotatable.
type MockRotatableMockRecorder struct {
	mock *MockRotatable
}

// NewMockRotatable creates a new mock instance.
func NewMockRotatable(ctrl *gomock.Controller) *MockRotatable {
	mock := &MockRotatable{ctrl: ctrl}
	mock.recorder = &MockRotatableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRotatable) EXPECT() *MockRotatableMockRecorder {
	return m.recorder
}

// GetAngularVelocity mocks base method.
func (m *MockRotatable) GetAngularVelocity() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAngularVelocity")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAngularVelocity indicates an expected call of GetAngularVelocity.
func (mr *MockRotatableMockRecorder) GetAngularVelocity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAngularVelocity", reflect.TypeOf((*MockRotatable)(nil).GetAngularVelocity))
}

// GetDirection mocks base method.
func (m *MockRotatable) GetDirection() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirection")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirection indicates an expected call of GetDirection.
func (mr *MockRotatableMockRecorder) GetDirection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirection", reflect.TypeOf((*MockRotatable)(nil).GetDirection))
}

// GetMaxDirections mocks base method.
func (m *MockRotatable) GetMaxDirections() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxDirections")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxDirections indicates an expected call of GetMaxDirections.
func (mr *MockRotatableMockRecorder) GetMaxDirections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxDirections", reflect.TypeOf((*MockRotatable)(nil).GetMaxDirections))
}

// SetDirection mocks base method.
func (m *MockRotatable) SetDirection(direction int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDirection", direction)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDirection indicates an expected call of SetDirection.
func (mr *MockRotatableMockRecorder) SetDirection(direction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDirection", reflect.TypeOf((*MockRotatable)(nil).SetDirection), direction)
}
