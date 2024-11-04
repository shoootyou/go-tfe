// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go
//
// Generated by this command:
//
//	mockgen -source=comment.go -destination=mocks/comment_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockComments is a mock of Comments interface.
type MockComments struct {
	ctrl     *gomock.Controller
	recorder *MockCommentsMockRecorder
}

// MockCommentsMockRecorder is the mock recorder for MockComments.
type MockCommentsMockRecorder struct {
	mock *MockComments
}

// NewMockComments creates a new mock instance.
func NewMockComments(ctrl *gomock.Controller) *MockComments {
	mock := &MockComments{ctrl: ctrl}
	mock.recorder = &MockCommentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComments) EXPECT() *MockCommentsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockComments) Create(ctx context.Context, runID string, options tfe.CommentCreateOptions) (*tfe.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, runID, options)
	ret0, _ := ret[0].(*tfe.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentsMockRecorder) Create(ctx, runID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockComments)(nil).Create), ctx, runID, options)
}

// List mocks base method.
func (m *MockComments) List(ctx context.Context, runID string) (*tfe.CommentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, runID)
	ret0, _ := ret[0].(*tfe.CommentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCommentsMockRecorder) List(ctx, runID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockComments)(nil).List), ctx, runID)
}

// Read mocks base method.
func (m *MockComments) Read(ctx context.Context, commentID string) (*tfe.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, commentID)
	ret0, _ := ret[0].(*tfe.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCommentsMockRecorder) Read(ctx, commentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockComments)(nil).Read), ctx, commentID)
}
