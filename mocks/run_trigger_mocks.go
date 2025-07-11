// Code generated by MockGen. DO NOT EDIT.
// Source: run_trigger.go
//
// Generated by this command:
//
//	mockgen -source=run_trigger.go -destination=mocks/run_trigger_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockRunTriggers is a mock of RunTriggers interface.
type MockRunTriggers struct {
	ctrl     *gomock.Controller
	recorder *MockRunTriggersMockRecorder
}

// MockRunTriggersMockRecorder is the mock recorder for MockRunTriggers.
type MockRunTriggersMockRecorder struct {
	mock *MockRunTriggers
}

// NewMockRunTriggers creates a new mock instance.
func NewMockRunTriggers(ctrl *gomock.Controller) *MockRunTriggers {
	mock := &MockRunTriggers{ctrl: ctrl}
	mock.recorder = &MockRunTriggersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunTriggers) EXPECT() *MockRunTriggersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRunTriggers) Create(ctx context.Context, workspaceID string, options tfe.RunTriggerCreateOptions) (*tfe.RunTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.RunTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRunTriggersMockRecorder) Create(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRunTriggers)(nil).Create), ctx, workspaceID, options)
}

// Delete mocks base method.
func (m *MockRunTriggers) Delete(ctx context.Context, RunTriggerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, RunTriggerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRunTriggersMockRecorder) Delete(ctx, RunTriggerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRunTriggers)(nil).Delete), ctx, RunTriggerID)
}

// List mocks base method.
func (m *MockRunTriggers) List(ctx context.Context, workspaceID string, options *tfe.RunTriggerListOptions) (*tfe.RunTriggerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.RunTriggerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRunTriggersMockRecorder) List(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRunTriggers)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockRunTriggers) Read(ctx context.Context, RunTriggerID string) (*tfe.RunTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, RunTriggerID)
	ret0, _ := ret[0].(*tfe.RunTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRunTriggersMockRecorder) Read(ctx, RunTriggerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRunTriggers)(nil).Read), ctx, RunTriggerID)
}
