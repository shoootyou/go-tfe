// Code generated by MockGen. DO NOT EDIT.
// Source: workspace.go
//
// Generated by this command:
//
//	mockgen -source=workspace.go -destination=mocks/workspace_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkspaces is a mock of Workspaces interface.
type MockWorkspaces struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspacesMockRecorder
}

// MockWorkspacesMockRecorder is the mock recorder for MockWorkspaces.
type MockWorkspacesMockRecorder struct {
	mock *MockWorkspaces
}

// NewMockWorkspaces creates a new mock instance.
func NewMockWorkspaces(ctrl *gomock.Controller) *MockWorkspaces {
	mock := &MockWorkspaces{ctrl: ctrl}
	mock.recorder = &MockWorkspacesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaces) EXPECT() *MockWorkspacesMockRecorder {
	return m.recorder
}

// AddRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) AddRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceAddRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRemoteStateConsumers indicates an expected call of AddRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) AddRemoteStateConsumers(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).AddRemoteStateConsumers), ctx, workspaceID, options)
}

// AddTagBindings mocks base method.
func (m *MockWorkspaces) AddTagBindings(ctx context.Context, workspaceID string, options tfe.WorkspaceAddTagBindingsOptions) ([]*tfe.TagBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagBindings", ctx, workspaceID, options)
	ret0, _ := ret[0].([]*tfe.TagBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTagBindings indicates an expected call of AddTagBindings.
func (mr *MockWorkspacesMockRecorder) AddTagBindings(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagBindings", reflect.TypeOf((*MockWorkspaces)(nil).AddTagBindings), ctx, workspaceID, options)
}

// AddTags mocks base method.
func (m *MockWorkspaces) AddTags(ctx context.Context, workspaceID string, options tfe.WorkspaceAddTagsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTags", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTags indicates an expected call of AddTags.
func (mr *MockWorkspacesMockRecorder) AddTags(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTags", reflect.TypeOf((*MockWorkspaces)(nil).AddTags), ctx, workspaceID, options)
}

// AssignSSHKey mocks base method.
func (m *MockWorkspaces) AssignSSHKey(ctx context.Context, workspaceID string, options tfe.WorkspaceAssignSSHKeyOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignSSHKey", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignSSHKey indicates an expected call of AssignSSHKey.
func (mr *MockWorkspacesMockRecorder) AssignSSHKey(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).AssignSSHKey), ctx, workspaceID, options)
}

// Create mocks base method.
func (m *MockWorkspaces) Create(ctx context.Context, organization string, options tfe.WorkspaceCreateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWorkspacesMockRecorder) Create(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWorkspaces)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockWorkspaces) Delete(ctx context.Context, organization, workspace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization, workspace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWorkspacesMockRecorder) Delete(ctx, organization, workspace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWorkspaces)(nil).Delete), ctx, organization, workspace)
}

// DeleteByID mocks base method.
func (m *MockWorkspaces) DeleteByID(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockWorkspacesMockRecorder) DeleteByID(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockWorkspaces)(nil).DeleteByID), ctx, workspaceID)
}

// DeleteDataRetentionPolicy mocks base method.
func (m *MockWorkspaces) DeleteDataRetentionPolicy(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataRetentionPolicy", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataRetentionPolicy indicates an expected call of DeleteDataRetentionPolicy.
func (mr *MockWorkspacesMockRecorder) DeleteDataRetentionPolicy(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataRetentionPolicy", reflect.TypeOf((*MockWorkspaces)(nil).DeleteDataRetentionPolicy), ctx, workspaceID)
}

// ForceUnlock mocks base method.
func (m *MockWorkspaces) ForceUnlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUnlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForceUnlock indicates an expected call of ForceUnlock.
func (mr *MockWorkspacesMockRecorder) ForceUnlock(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUnlock", reflect.TypeOf((*MockWorkspaces)(nil).ForceUnlock), ctx, workspaceID)
}

// List mocks base method.
func (m *MockWorkspaces) List(ctx context.Context, organization string, options *tfe.WorkspaceListOptions) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkspacesMockRecorder) List(ctx, organization, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkspaces)(nil).List), ctx, organization, options)
}

// ListRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) ListRemoteStateConsumers(ctx context.Context, workspaceID string, options *tfe.RemoteStateConsumersListOptions) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRemoteStateConsumers indicates an expected call of ListRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) ListRemoteStateConsumers(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).ListRemoteStateConsumers), ctx, workspaceID, options)
}

// ListTagBindings mocks base method.
func (m *MockWorkspaces) ListTagBindings(ctx context.Context, workspaceID string) ([]*tfe.TagBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagBindings", ctx, workspaceID)
	ret0, _ := ret[0].([]*tfe.TagBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagBindings indicates an expected call of ListTagBindings.
func (mr *MockWorkspacesMockRecorder) ListTagBindings(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagBindings", reflect.TypeOf((*MockWorkspaces)(nil).ListTagBindings), ctx, workspaceID)
}

// ListTags mocks base method.
func (m *MockWorkspaces) ListTags(ctx context.Context, workspaceID string, options *tfe.WorkspaceTagListOptions) (*tfe.TagList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.TagList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockWorkspacesMockRecorder) ListTags(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockWorkspaces)(nil).ListTags), ctx, workspaceID, options)
}

// Lock mocks base method.
func (m *MockWorkspaces) Lock(ctx context.Context, workspaceID string, options tfe.WorkspaceLockOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockWorkspacesMockRecorder) Lock(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockWorkspaces)(nil).Lock), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockWorkspaces) Read(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockWorkspacesMockRecorder) Read(ctx, organization, workspace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockWorkspaces)(nil).Read), ctx, organization, workspace)
}

// ReadByID mocks base method.
func (m *MockWorkspaces) ReadByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByID indicates an expected call of ReadByID.
func (mr *MockWorkspacesMockRecorder) ReadByID(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByID", reflect.TypeOf((*MockWorkspaces)(nil).ReadByID), ctx, workspaceID)
}

// ReadByIDWithOptions mocks base method.
func (m *MockWorkspaces) ReadByIDWithOptions(ctx context.Context, workspaceID string, options *tfe.WorkspaceReadOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByIDWithOptions", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByIDWithOptions indicates an expected call of ReadByIDWithOptions.
func (mr *MockWorkspacesMockRecorder) ReadByIDWithOptions(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByIDWithOptions", reflect.TypeOf((*MockWorkspaces)(nil).ReadByIDWithOptions), ctx, workspaceID, options)
}

// ReadDataRetentionPolicy mocks base method.
func (m *MockWorkspaces) ReadDataRetentionPolicy(ctx context.Context, workspaceID string) (*tfe.DataRetentionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDataRetentionPolicy", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.DataRetentionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDataRetentionPolicy indicates an expected call of ReadDataRetentionPolicy.
func (mr *MockWorkspacesMockRecorder) ReadDataRetentionPolicy(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDataRetentionPolicy", reflect.TypeOf((*MockWorkspaces)(nil).ReadDataRetentionPolicy), ctx, workspaceID)
}

// ReadDataRetentionPolicyChoice mocks base method.
func (m *MockWorkspaces) ReadDataRetentionPolicyChoice(ctx context.Context, workspaceID string) (*tfe.DataRetentionPolicyChoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDataRetentionPolicyChoice", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.DataRetentionPolicyChoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDataRetentionPolicyChoice indicates an expected call of ReadDataRetentionPolicyChoice.
func (mr *MockWorkspacesMockRecorder) ReadDataRetentionPolicyChoice(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDataRetentionPolicyChoice", reflect.TypeOf((*MockWorkspaces)(nil).ReadDataRetentionPolicyChoice), ctx, workspaceID)
}

// ReadWithOptions mocks base method.
func (m *MockWorkspaces) ReadWithOptions(ctx context.Context, organization, workspace string, options *tfe.WorkspaceReadOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, organization, workspace, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockWorkspacesMockRecorder) ReadWithOptions(ctx, organization, workspace, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockWorkspaces)(nil).ReadWithOptions), ctx, organization, workspace, options)
}

// Readme mocks base method.
func (m *MockWorkspaces) Readme(ctx context.Context, workspaceID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readme", ctx, workspaceID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readme indicates an expected call of Readme.
func (mr *MockWorkspacesMockRecorder) Readme(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readme", reflect.TypeOf((*MockWorkspaces)(nil).Readme), ctx, workspaceID)
}

// RemoveRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) RemoveRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceRemoveRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRemoteStateConsumers indicates an expected call of RemoveRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) RemoveRemoteStateConsumers(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).RemoveRemoteStateConsumers), ctx, workspaceID, options)
}

// RemoveTags mocks base method.
func (m *MockWorkspaces) RemoveTags(ctx context.Context, workspaceID string, options tfe.WorkspaceRemoveTagsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTags", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTags indicates an expected call of RemoveTags.
func (mr *MockWorkspacesMockRecorder) RemoveTags(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTags", reflect.TypeOf((*MockWorkspaces)(nil).RemoveTags), ctx, workspaceID, options)
}

// RemoveVCSConnection mocks base method.
func (m *MockWorkspaces) RemoveVCSConnection(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnection", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnection indicates an expected call of RemoveVCSConnection.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnection(ctx, organization, workspace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnection", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnection), ctx, organization, workspace)
}

// RemoveVCSConnectionByID mocks base method.
func (m *MockWorkspaces) RemoveVCSConnectionByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnectionByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnectionByID indicates an expected call of RemoveVCSConnectionByID.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnectionByID(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnectionByID", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnectionByID), ctx, workspaceID)
}

// SafeDelete mocks base method.
func (m *MockWorkspaces) SafeDelete(ctx context.Context, organization, workspace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeDelete", ctx, organization, workspace)
	ret0, _ := ret[0].(error)
	return ret0
}

// SafeDelete indicates an expected call of SafeDelete.
func (mr *MockWorkspacesMockRecorder) SafeDelete(ctx, organization, workspace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeDelete", reflect.TypeOf((*MockWorkspaces)(nil).SafeDelete), ctx, organization, workspace)
}

// SafeDeleteByID mocks base method.
func (m *MockWorkspaces) SafeDeleteByID(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeDeleteByID", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SafeDeleteByID indicates an expected call of SafeDeleteByID.
func (mr *MockWorkspacesMockRecorder) SafeDeleteByID(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeDeleteByID", reflect.TypeOf((*MockWorkspaces)(nil).SafeDeleteByID), ctx, workspaceID)
}

// SetDataRetentionPolicy mocks base method.
func (m *MockWorkspaces) SetDataRetentionPolicy(ctx context.Context, workspaceID string, options tfe.DataRetentionPolicySetOptions) (*tfe.DataRetentionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDataRetentionPolicy", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.DataRetentionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDataRetentionPolicy indicates an expected call of SetDataRetentionPolicy.
func (mr *MockWorkspacesMockRecorder) SetDataRetentionPolicy(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDataRetentionPolicy", reflect.TypeOf((*MockWorkspaces)(nil).SetDataRetentionPolicy), ctx, workspaceID, options)
}

// SetDataRetentionPolicyDeleteOlder mocks base method.
func (m *MockWorkspaces) SetDataRetentionPolicyDeleteOlder(ctx context.Context, workspaceID string, options tfe.DataRetentionPolicyDeleteOlderSetOptions) (*tfe.DataRetentionPolicyDeleteOlder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDataRetentionPolicyDeleteOlder", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.DataRetentionPolicyDeleteOlder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDataRetentionPolicyDeleteOlder indicates an expected call of SetDataRetentionPolicyDeleteOlder.
func (mr *MockWorkspacesMockRecorder) SetDataRetentionPolicyDeleteOlder(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDataRetentionPolicyDeleteOlder", reflect.TypeOf((*MockWorkspaces)(nil).SetDataRetentionPolicyDeleteOlder), ctx, workspaceID, options)
}

// SetDataRetentionPolicyDontDelete mocks base method.
func (m *MockWorkspaces) SetDataRetentionPolicyDontDelete(ctx context.Context, workspaceID string, options tfe.DataRetentionPolicyDontDeleteSetOptions) (*tfe.DataRetentionPolicyDontDelete, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDataRetentionPolicyDontDelete", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.DataRetentionPolicyDontDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDataRetentionPolicyDontDelete indicates an expected call of SetDataRetentionPolicyDontDelete.
func (mr *MockWorkspacesMockRecorder) SetDataRetentionPolicyDontDelete(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDataRetentionPolicyDontDelete", reflect.TypeOf((*MockWorkspaces)(nil).SetDataRetentionPolicyDontDelete), ctx, workspaceID, options)
}

// UnassignSSHKey mocks base method.
func (m *MockWorkspaces) UnassignSSHKey(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignSSHKey", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnassignSSHKey indicates an expected call of UnassignSSHKey.
func (mr *MockWorkspacesMockRecorder) UnassignSSHKey(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).UnassignSSHKey), ctx, workspaceID)
}

// Unlock mocks base method.
func (m *MockWorkspaces) Unlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockWorkspacesMockRecorder) Unlock(ctx, workspaceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockWorkspaces)(nil).Unlock), ctx, workspaceID)
}

// Update mocks base method.
func (m *MockWorkspaces) Update(ctx context.Context, organization, workspace string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organization, workspace, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWorkspacesMockRecorder) Update(ctx, organization, workspace, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWorkspaces)(nil).Update), ctx, organization, workspace, options)
}

// UpdateByID mocks base method.
func (m *MockWorkspaces) UpdateByID(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockWorkspacesMockRecorder) UpdateByID(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockWorkspaces)(nil).UpdateByID), ctx, workspaceID, options)
}

// UpdateRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) UpdateRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteStateConsumers indicates an expected call of UpdateRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) UpdateRemoteStateConsumers(ctx, workspaceID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).UpdateRemoteStateConsumers), ctx, workspaceID, options)
}
