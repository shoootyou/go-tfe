// Code generated by MockGen. DO NOT EDIT.
// Source: admin_terraform_version.go
//
// Generated by this command:
//
//	mockgen -source=admin_terraform_version.go -destination=mocks/admin_terraform_version_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminTerraformVersions is a mock of AdminTerraformVersions interface.
type MockAdminTerraformVersions struct {
	ctrl     *gomock.Controller
	recorder *MockAdminTerraformVersionsMockRecorder
}

// MockAdminTerraformVersionsMockRecorder is the mock recorder for MockAdminTerraformVersions.
type MockAdminTerraformVersionsMockRecorder struct {
	mock *MockAdminTerraformVersions
}

// NewMockAdminTerraformVersions creates a new mock instance.
func NewMockAdminTerraformVersions(ctrl *gomock.Controller) *MockAdminTerraformVersions {
	mock := &MockAdminTerraformVersions{ctrl: ctrl}
	mock.recorder = &MockAdminTerraformVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminTerraformVersions) EXPECT() *MockAdminTerraformVersionsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdminTerraformVersions) Create(ctx context.Context, options tfe.AdminTerraformVersionCreateOptions) (*tfe.AdminTerraformVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminTerraformVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminTerraformVersionsMockRecorder) Create(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdminTerraformVersions)(nil).Create), ctx, options)
}

// Delete mocks base method.
func (m *MockAdminTerraformVersions) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminTerraformVersionsMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminTerraformVersions)(nil).Delete), ctx, id)
}

// List mocks base method.
func (m *MockAdminTerraformVersions) List(ctx context.Context, options *tfe.AdminTerraformVersionsListOptions) (*tfe.AdminTerraformVersionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminTerraformVersionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminTerraformVersionsMockRecorder) List(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdminTerraformVersions)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockAdminTerraformVersions) Read(ctx context.Context, id string) (*tfe.AdminTerraformVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, id)
	ret0, _ := ret[0].(*tfe.AdminTerraformVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAdminTerraformVersionsMockRecorder) Read(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAdminTerraformVersions)(nil).Read), ctx, id)
}

// Update mocks base method.
func (m *MockAdminTerraformVersions) Update(ctx context.Context, id string, options tfe.AdminTerraformVersionUpdateOptions) (*tfe.AdminTerraformVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, options)
	ret0, _ := ret[0].(*tfe.AdminTerraformVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminTerraformVersionsMockRecorder) Update(ctx, id, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminTerraformVersions)(nil).Update), ctx, id, options)
}
