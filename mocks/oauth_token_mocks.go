// Code generated by MockGen. DO NOT EDIT.
// Source: oauth_token.go
//
// Generated by this command:
//
//	mockgen -source=oauth_token.go -destination=mocks/oauth_token_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockOAuthTokens is a mock of OAuthTokens interface.
type MockOAuthTokens struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthTokensMockRecorder
}

// MockOAuthTokensMockRecorder is the mock recorder for MockOAuthTokens.
type MockOAuthTokensMockRecorder struct {
	mock *MockOAuthTokens
}

// NewMockOAuthTokens creates a new mock instance.
func NewMockOAuthTokens(ctrl *gomock.Controller) *MockOAuthTokens {
	mock := &MockOAuthTokens{ctrl: ctrl}
	mock.recorder = &MockOAuthTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthTokens) EXPECT() *MockOAuthTokensMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOAuthTokens) Delete(ctx context.Context, oAuthTokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, oAuthTokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOAuthTokensMockRecorder) Delete(ctx, oAuthTokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOAuthTokens)(nil).Delete), ctx, oAuthTokenID)
}

// List mocks base method.
func (m *MockOAuthTokens) List(ctx context.Context, organization string, options *tfe.OAuthTokenListOptions) (*tfe.OAuthTokenList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.OAuthTokenList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOAuthTokensMockRecorder) List(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOAuthTokens)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockOAuthTokens) Read(ctx context.Context, oAuthTokenID string) (*tfe.OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, oAuthTokenID)
	ret0, _ := ret[0].(*tfe.OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOAuthTokensMockRecorder) Read(ctx, oAuthTokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOAuthTokens)(nil).Read), ctx, oAuthTokenID)
}

// Update mocks base method.
func (m *MockOAuthTokens) Update(ctx context.Context, oAuthTokenID string, options tfe.OAuthTokenUpdateOptions) (*tfe.OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, oAuthTokenID, options)
	ret0, _ := ret[0].(*tfe.OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOAuthTokensMockRecorder) Update(ctx, oAuthTokenID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOAuthTokens)(nil).Update), ctx, oAuthTokenID, options)
}
