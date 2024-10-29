// Code generated by MockGen. DO NOT EDIT.
// Source: task_result.go
//
// Generated by this command:
//
//	mockgen -source=task_result.go -destination=mocks/task_result_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskResults is a mock of TaskResults interface.
type MockTaskResults struct {
	ctrl     *gomock.Controller
	recorder *MockTaskResultsMockRecorder
}

// MockTaskResultsMockRecorder is the mock recorder for MockTaskResults.
type MockTaskResultsMockRecorder struct {
	mock *MockTaskResults
}

// NewMockTaskResults creates a new mock instance.
func NewMockTaskResults(ctrl *gomock.Controller) *MockTaskResults {
	mock := &MockTaskResults{ctrl: ctrl}
	mock.recorder = &MockTaskResultsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskResults) EXPECT() *MockTaskResultsMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockTaskResults) Read(ctx context.Context, taskResultID string) (*tfe.TaskResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, taskResultID)
	ret0, _ := ret[0].(*tfe.TaskResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTaskResultsMockRecorder) Read(ctx, taskResultID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTaskResults)(nil).Read), ctx, taskResultID)
}
