// Code generated by MockGen. DO NOT EDIT.
// Source: admin_workspace.go
//
// Generated by this command:
//
//	mockgen -source=admin_workspace.go -destination=mocks/admin_workspace_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminWorkspaces is a mock of AdminWorkspaces interface.
type MockAdminWorkspaces struct {
	ctrl     *gomock.Controller
	recorder *MockAdminWorkspacesMockRecorder
}

// MockAdminWorkspacesMockRecorder is the mock recorder for MockAdminWorkspaces.
type MockAdminWorkspacesMockRecorder struct {
	mock *MockAdminWorkspaces
}

// NewMockAdminWorkspaces creates a new mock instance.
func NewMockAdminWorkspaces(ctrl *gomock.Controller) *MockAdminWorkspaces {
	mock := &MockAdminWorkspaces{ctrl: ctrl}
	mock.recorder = &MockAdminWorkspacesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminWorkspaces) EXPECT() *MockAdminWorkspacesMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAdminWorkspaces) Delete(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminWorkspacesMockRecorder) Delete(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminWorkspaces)(nil).Delete), ctx, workspaceID)
}

// List mocks base method.
func (m *MockAdminWorkspaces) List(ctx context.Context, options *tfe.AdminWorkspaceListOptions) (*tfe.AdminWorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminWorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminWorkspacesMockRecorder) List(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdminWorkspaces)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockAdminWorkspaces) Read(ctx context.Context, workspaceID string) (*tfe.AdminWorkspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.AdminWorkspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAdminWorkspacesMockRecorder) Read(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAdminWorkspaces)(nil).Read), ctx, workspaceID)
}
