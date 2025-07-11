// Code generated by MockGen. DO NOT EDIT.
// Source: workspace_run_task.go
//
// Generated by this command:
//
//	mockgen -source=workspace_run_task.go -destination=mocks/workspace_run_tasks_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkspaceRunTasks is a mock of WorkspaceRunTasks interface.
type MockWorkspaceRunTasks struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceRunTasksMockRecorder
}

// MockWorkspaceRunTasksMockRecorder is the mock recorder for MockWorkspaceRunTasks.
type MockWorkspaceRunTasksMockRecorder struct {
	mock *MockWorkspaceRunTasks
}

// NewMockWorkspaceRunTasks creates a new mock instance.
func NewMockWorkspaceRunTasks(ctrl *gomock.Controller) *MockWorkspaceRunTasks {
	mock := &MockWorkspaceRunTasks{ctrl: ctrl}
	mock.recorder = &MockWorkspaceRunTasksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceRunTasks) EXPECT() *MockWorkspaceRunTasksMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWorkspaceRunTasks) Create(ctx context.Context, workspaceID string, options tfe.WorkspaceRunTaskCreateOptions) (*tfe.WorkspaceRunTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceRunTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWorkspaceRunTasksMockRecorder) Create(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWorkspaceRunTasks)(nil).Create), ctx, workspaceID, options)
}

// Delete mocks base method.
func (m *MockWorkspaceRunTasks) Delete(ctx context.Context, workspaceID, workspaceTaskID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, workspaceID, workspaceTaskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWorkspaceRunTasksMockRecorder) Delete(ctx, workspaceID, workspaceTaskID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWorkspaceRunTasks)(nil).Delete), ctx, workspaceID, workspaceTaskID)
}

// List mocks base method.
func (m *MockWorkspaceRunTasks) List(ctx context.Context, workspaceID string, options *tfe.WorkspaceRunTaskListOptions) (*tfe.WorkspaceRunTaskList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceRunTaskList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkspaceRunTasksMockRecorder) List(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkspaceRunTasks)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockWorkspaceRunTasks) Read(ctx context.Context, workspaceID, workspaceTaskID string) (*tfe.WorkspaceRunTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, workspaceID, workspaceTaskID)
	ret0, _ := ret[0].(*tfe.WorkspaceRunTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockWorkspaceRunTasksMockRecorder) Read(ctx, workspaceID, workspaceTaskID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockWorkspaceRunTasks)(nil).Read), ctx, workspaceID, workspaceTaskID)
}

// Update mocks base method.
func (m *MockWorkspaceRunTasks) Update(ctx context.Context, workspaceID, workspaceTaskID string, options tfe.WorkspaceRunTaskUpdateOptions) (*tfe.WorkspaceRunTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, workspaceID, workspaceTaskID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceRunTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWorkspaceRunTasksMockRecorder) Update(ctx, workspaceID, workspaceTaskID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWorkspaceRunTasks)(nil).Update), ctx, workspaceID, workspaceTaskID, options)
}
