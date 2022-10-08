// Code generated by MockGen. DO NOT EDIT.
// Source: burn_fuel.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBurnFuel is a mock of IBurnFuel interface.
type MockIBurnFuel struct {
	ctrl     *gomock.Controller
	recorder *MockIBurnFuelMockRecorder
}

// MockIBurnFuelMockRecorder is the mock recorder for MockIBurnFuel.
type MockIBurnFuelMockRecorder struct {
	mock *MockIBurnFuel
}

// NewMockIBurnFuel creates a new mock instance.
func NewMockIBurnFuel(ctrl *gomock.Controller) *MockIBurnFuel {
	mock := &MockIBurnFuel{ctrl: ctrl}
	mock.recorder = &MockIBurnFuelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBurnFuel) EXPECT() *MockIBurnFuelMockRecorder {
	return m.recorder
}

// GetFuel mocks base method.
func (m *MockIBurnFuel) GetFuel() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuel")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFuel indicates an expected call of GetFuel.
func (mr *MockIBurnFuelMockRecorder) GetFuel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuel", reflect.TypeOf((*MockIBurnFuel)(nil).GetFuel))
}

// GetFuelConsumption mocks base method.
func (m *MockIBurnFuel) GetFuelConsumption() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuelConsumption")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFuelConsumption indicates an expected call of GetFuelConsumption.
func (mr *MockIBurnFuelMockRecorder) GetFuelConsumption() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuelConsumption", reflect.TypeOf((*MockIBurnFuel)(nil).GetFuelConsumption))
}

// SetFuel mocks base method.
func (m *MockIBurnFuel) SetFuel(fuel int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFuel", fuel)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFuel indicates an expected call of SetFuel.
func (mr *MockIBurnFuelMockRecorder) SetFuel(fuel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFuel", reflect.TypeOf((*MockIBurnFuel)(nil).SetFuel), fuel)
}
