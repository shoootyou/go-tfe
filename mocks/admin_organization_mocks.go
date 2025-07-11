// Code generated by MockGen. DO NOT EDIT.
// Source: admin_organization.go
//
// Generated by this command:
//
//	mockgen -source=admin_organization.go -destination=mocks/admin_organization_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminOrganizations is a mock of AdminOrganizations interface.
type MockAdminOrganizations struct {
	ctrl     *gomock.Controller
	recorder *MockAdminOrganizationsMockRecorder
}

// MockAdminOrganizationsMockRecorder is the mock recorder for MockAdminOrganizations.
type MockAdminOrganizationsMockRecorder struct {
	mock *MockAdminOrganizations
}

// NewMockAdminOrganizations creates a new mock instance.
func NewMockAdminOrganizations(ctrl *gomock.Controller) *MockAdminOrganizations {
	mock := &MockAdminOrganizations{ctrl: ctrl}
	mock.recorder = &MockAdminOrganizationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminOrganizations) EXPECT() *MockAdminOrganizationsMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAdminOrganizations) Delete(ctx context.Context, organization string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminOrganizationsMockRecorder) Delete(ctx, organization any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminOrganizations)(nil).Delete), ctx, organization)
}

// List mocks base method.
func (m *MockAdminOrganizations) List(ctx context.Context, options *tfe.AdminOrganizationListOptions) (*tfe.AdminOrganizationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminOrganizationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminOrganizationsMockRecorder) List(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdminOrganizations)(nil).List), ctx, options)
}

// ListModuleConsumers mocks base method.
func (m *MockAdminOrganizations) ListModuleConsumers(ctx context.Context, organization string, options *tfe.AdminOrganizationListModuleConsumersOptions) (*tfe.AdminOrganizationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModuleConsumers", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.AdminOrganizationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModuleConsumers indicates an expected call of ListModuleConsumers.
func (mr *MockAdminOrganizationsMockRecorder) ListModuleConsumers(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModuleConsumers", reflect.TypeOf((*MockAdminOrganizations)(nil).ListModuleConsumers), ctx, organization, options)
}

// Read mocks base method.
func (m *MockAdminOrganizations) Read(ctx context.Context, organization string) (*tfe.AdminOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization)
	ret0, _ := ret[0].(*tfe.AdminOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAdminOrganizationsMockRecorder) Read(ctx, organization any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAdminOrganizations)(nil).Read), ctx, organization)
}

// Update mocks base method.
func (m *MockAdminOrganizations) Update(ctx context.Context, organization string, options tfe.AdminOrganizationUpdateOptions) (*tfe.AdminOrganization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.AdminOrganization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminOrganizationsMockRecorder) Update(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminOrganizations)(nil).Update), ctx, organization, options)
}

// UpdateModuleConsumers mocks base method.
func (m *MockAdminOrganizations) UpdateModuleConsumers(ctx context.Context, organization string, consumerOrganizations []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateModuleConsumers", ctx, organization, consumerOrganizations)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateModuleConsumers indicates an expected call of UpdateModuleConsumers.
func (mr *MockAdminOrganizationsMockRecorder) UpdateModuleConsumers(ctx, organization, consumerOrganizations any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModuleConsumers", reflect.TypeOf((*MockAdminOrganizations)(nil).UpdateModuleConsumers), ctx, organization, consumerOrganizations)
}
