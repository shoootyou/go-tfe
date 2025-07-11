// Code generated by MockGen. DO NOT EDIT.
// Source: gpg_key.go
//
// Generated by this command:
//
//	mockgen -source=gpg_key.go -destination=mocks/gpg_key_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockGPGKeys is a mock of GPGKeys interface.
type MockGPGKeys struct {
	ctrl     *gomock.Controller
	recorder *MockGPGKeysMockRecorder
}

// MockGPGKeysMockRecorder is the mock recorder for MockGPGKeys.
type MockGPGKeysMockRecorder struct {
	mock *MockGPGKeys
}

// NewMockGPGKeys creates a new mock instance.
func NewMockGPGKeys(ctrl *gomock.Controller) *MockGPGKeys {
	mock := &MockGPGKeys{ctrl: ctrl}
	mock.recorder = &MockGPGKeysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGPGKeys) EXPECT() *MockGPGKeysMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGPGKeys) Create(ctx context.Context, registryName tfe.RegistryName, options tfe.GPGKeyCreateOptions) (*tfe.GPGKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, registryName, options)
	ret0, _ := ret[0].(*tfe.GPGKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockGPGKeysMockRecorder) Create(ctx, registryName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGPGKeys)(nil).Create), ctx, registryName, options)
}

// Delete mocks base method.
func (m *MockGPGKeys) Delete(ctx context.Context, keyID tfe.GPGKeyID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, keyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGPGKeysMockRecorder) Delete(ctx, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGPGKeys)(nil).Delete), ctx, keyID)
}

// ListPrivate mocks base method.
func (m *MockGPGKeys) ListPrivate(ctx context.Context, options tfe.GPGKeyListOptions) (*tfe.GPGKeyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPrivate", ctx, options)
	ret0, _ := ret[0].(*tfe.GPGKeyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPrivate indicates an expected call of ListPrivate.
func (mr *MockGPGKeysMockRecorder) ListPrivate(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrivate", reflect.TypeOf((*MockGPGKeys)(nil).ListPrivate), ctx, options)
}

// Read mocks base method.
func (m *MockGPGKeys) Read(ctx context.Context, keyID tfe.GPGKeyID) (*tfe.GPGKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, keyID)
	ret0, _ := ret[0].(*tfe.GPGKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockGPGKeysMockRecorder) Read(ctx, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockGPGKeys)(nil).Read), ctx, keyID)
}

// Update mocks base method.
func (m *MockGPGKeys) Update(ctx context.Context, keyID tfe.GPGKeyID, options tfe.GPGKeyUpdateOptions) (*tfe.GPGKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, keyID, options)
	ret0, _ := ret[0].(*tfe.GPGKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockGPGKeysMockRecorder) Update(ctx, keyID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGPGKeys)(nil).Update), ctx, keyID, options)
}
