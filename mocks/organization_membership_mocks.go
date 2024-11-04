// Code generated by MockGen. DO NOT EDIT.
// Source: organization_membership.go
//
// Generated by this command:
//
//	mockgen -source=organization_membership.go -destination=mocks/organization_membership_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockOrganizationMemberships is a mock of OrganizationMemberships interface.
type MockOrganizationMemberships struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationMembershipsMockRecorder
}

// MockOrganizationMembershipsMockRecorder is the mock recorder for MockOrganizationMemberships.
type MockOrganizationMembershipsMockRecorder struct {
	mock *MockOrganizationMemberships
}

// NewMockOrganizationMemberships creates a new mock instance.
func NewMockOrganizationMemberships(ctrl *gomock.Controller) *MockOrganizationMemberships {
	mock := &MockOrganizationMemberships{ctrl: ctrl}
	mock.recorder = &MockOrganizationMembershipsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationMemberships) EXPECT() *MockOrganizationMembershipsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganizationMemberships) Create(ctx context.Context, organization string, options tfe.OrganizationMembershipCreateOptions) (*tfe.OrganizationMembership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.OrganizationMembership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationMembershipsMockRecorder) Create(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationMemberships)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockOrganizationMemberships) Delete(ctx context.Context, organizationMembershipID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organizationMembershipID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationMembershipsMockRecorder) Delete(ctx, organizationMembershipID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationMemberships)(nil).Delete), ctx, organizationMembershipID)
}

// List mocks base method.
func (m *MockOrganizationMemberships) List(ctx context.Context, organization string, options *tfe.OrganizationMembershipListOptions) (*tfe.OrganizationMembershipList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.OrganizationMembershipList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrganizationMembershipsMockRecorder) List(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationMemberships)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockOrganizationMemberships) Read(ctx context.Context, organizationMembershipID string) (*tfe.OrganizationMembership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organizationMembershipID)
	ret0, _ := ret[0].(*tfe.OrganizationMembership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOrganizationMembershipsMockRecorder) Read(ctx, organizationMembershipID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOrganizationMemberships)(nil).Read), ctx, organizationMembershipID)
}

// ReadWithOptions mocks base method.
func (m *MockOrganizationMemberships) ReadWithOptions(ctx context.Context, organizationMembershipID string, options tfe.OrganizationMembershipReadOptions) (*tfe.OrganizationMembership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, organizationMembershipID, options)
	ret0, _ := ret[0].(*tfe.OrganizationMembership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockOrganizationMembershipsMockRecorder) ReadWithOptions(ctx, organizationMembershipID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockOrganizationMemberships)(nil).ReadWithOptions), ctx, organizationMembershipID, options)
}
