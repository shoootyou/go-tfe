// Code generated by MockGen. DO NOT EDIT.
// Source: admin_opa_version.go
//
// Generated by this command:
//
//	mockgen -source=admin_opa_version.go -destination=mocks/admin_opa_version_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminOPAVersions is a mock of AdminOPAVersions interface.
type MockAdminOPAVersions struct {
	ctrl     *gomock.Controller
	recorder *MockAdminOPAVersionsMockRecorder
}

// MockAdminOPAVersionsMockRecorder is the mock recorder for MockAdminOPAVersions.
type MockAdminOPAVersionsMockRecorder struct {
	mock *MockAdminOPAVersions
}

// NewMockAdminOPAVersions creates a new mock instance.
func NewMockAdminOPAVersions(ctrl *gomock.Controller) *MockAdminOPAVersions {
	mock := &MockAdminOPAVersions{ctrl: ctrl}
	mock.recorder = &MockAdminOPAVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminOPAVersions) EXPECT() *MockAdminOPAVersionsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdminOPAVersions) Create(ctx context.Context, options tfe.AdminOPAVersionCreateOptions) (*tfe.AdminOPAVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminOPAVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminOPAVersionsMockRecorder) Create(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdminOPAVersions)(nil).Create), ctx, options)
}

// Delete mocks base method.
func (m *MockAdminOPAVersions) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminOPAVersionsMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminOPAVersions)(nil).Delete), ctx, id)
}

// List mocks base method.
func (m *MockAdminOPAVersions) List(ctx context.Context, options *tfe.AdminOPAVersionsListOptions) (*tfe.AdminOPAVersionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminOPAVersionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminOPAVersionsMockRecorder) List(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdminOPAVersions)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockAdminOPAVersions) Read(ctx context.Context, id string) (*tfe.AdminOPAVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, id)
	ret0, _ := ret[0].(*tfe.AdminOPAVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAdminOPAVersionsMockRecorder) Read(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAdminOPAVersions)(nil).Read), ctx, id)
}

// Update mocks base method.
func (m *MockAdminOPAVersions) Update(ctx context.Context, id string, options tfe.AdminOPAVersionUpdateOptions) (*tfe.AdminOPAVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, options)
	ret0, _ := ret[0].(*tfe.AdminOPAVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminOPAVersionsMockRecorder) Update(ctx, id, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminOPAVersions)(nil).Update), ctx, id, options)
}
