// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/drone/autoscaler (interfaces: ServerStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	autoscaler "github.com/drone/autoscaler"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServerStore is a mock of ServerStore interface
type MockServerStore struct {
	ctrl     *gomock.Controller
	recorder *MockServerStoreMockRecorder
}

// MockServerStoreMockRecorder is the mock recorder for MockServerStore
type MockServerStoreMockRecorder struct {
	mock *MockServerStore
}

// NewMockServerStore creates a new mock instance
func NewMockServerStore(ctrl *gomock.Controller) *MockServerStore {
	mock := &MockServerStore{ctrl: ctrl}
	mock.recorder = &MockServerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerStore) EXPECT() *MockServerStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServerStore) Create(arg0 context.Context, arg1 *autoscaler.Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServerStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServerStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockServerStore) Delete(arg0 context.Context, arg1 *autoscaler.Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServerStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServerStore)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockServerStore) Find(arg0 context.Context, arg1 string) (*autoscaler.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockServerStoreMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockServerStore)(nil).Find), arg0, arg1)
}

// List mocks base method
func (m *MockServerStore) List(arg0 context.Context) ([]*autoscaler.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServerStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServerStore)(nil).List), arg0)
}

// ListState mocks base method
func (m *MockServerStore) ListState(arg0 context.Context, arg1 autoscaler.ServerState) ([]*autoscaler.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListState", arg0, arg1)
	ret0, _ := ret[0].([]*autoscaler.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListState indicates an expected call of ListState
func (mr *MockServerStoreMockRecorder) ListState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListState", reflect.TypeOf((*MockServerStore)(nil).ListState), arg0, arg1)
}

// Purge mocks base method
func (m *MockServerStore) Purge(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Purge", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Purge indicates an expected call of Purge
func (mr *MockServerStoreMockRecorder) Purge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Purge", reflect.TypeOf((*MockServerStore)(nil).Purge), arg0, arg1)
}

// Update mocks base method
func (m *MockServerStore) Update(arg0 context.Context, arg1 *autoscaler.Server) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockServerStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServerStore)(nil).Update), arg0, arg1)
}
