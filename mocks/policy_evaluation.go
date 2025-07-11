// Code generated by MockGen. DO NOT EDIT.
// Source: policy_evaluation.go
//
// Generated by this command:
//
//	mockgen -source=policy_evaluation.go -destination=mocks/policy_evaluation.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/shoootyou/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockPolicyEvaluations is a mock of PolicyEvaluations interface.
type MockPolicyEvaluations struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyEvaluationsMockRecorder
}

// MockPolicyEvaluationsMockRecorder is the mock recorder for MockPolicyEvaluations.
type MockPolicyEvaluationsMockRecorder struct {
	mock *MockPolicyEvaluations
}

// NewMockPolicyEvaluations creates a new mock instance.
func NewMockPolicyEvaluations(ctrl *gomock.Controller) *MockPolicyEvaluations {
	mock := &MockPolicyEvaluations{ctrl: ctrl}
	mock.recorder = &MockPolicyEvaluationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyEvaluations) EXPECT() *MockPolicyEvaluationsMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPolicyEvaluations) List(ctx context.Context, taskStageID string, options *tfe.PolicyEvaluationListOptions) (*tfe.PolicyEvaluationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, taskStageID, options)
	ret0, _ := ret[0].(*tfe.PolicyEvaluationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPolicyEvaluationsMockRecorder) List(ctx, taskStageID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicyEvaluations)(nil).List), ctx, taskStageID, options)
}

// MockPolicySetOutcomes is a mock of PolicySetOutcomes interface.
type MockPolicySetOutcomes struct {
	ctrl     *gomock.Controller
	recorder *MockPolicySetOutcomesMockRecorder
}

// MockPolicySetOutcomesMockRecorder is the mock recorder for MockPolicySetOutcomes.
type MockPolicySetOutcomesMockRecorder struct {
	mock *MockPolicySetOutcomes
}

// NewMockPolicySetOutcomes creates a new mock instance.
func NewMockPolicySetOutcomes(ctrl *gomock.Controller) *MockPolicySetOutcomes {
	mock := &MockPolicySetOutcomes{ctrl: ctrl}
	mock.recorder = &MockPolicySetOutcomesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicySetOutcomes) EXPECT() *MockPolicySetOutcomesMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPolicySetOutcomes) List(ctx context.Context, policyEvaluationID string, options *tfe.PolicySetOutcomeListOptions) (*tfe.PolicySetOutcomeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, policyEvaluationID, options)
	ret0, _ := ret[0].(*tfe.PolicySetOutcomeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPolicySetOutcomesMockRecorder) List(ctx, policyEvaluationID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicySetOutcomes)(nil).List), ctx, policyEvaluationID, options)
}

// Read mocks base method.
func (m *MockPolicySetOutcomes) Read(ctx context.Context, policySetOutcomeID string) (*tfe.PolicySetOutcome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, policySetOutcomeID)
	ret0, _ := ret[0].(*tfe.PolicySetOutcome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPolicySetOutcomesMockRecorder) Read(ctx, policySetOutcomeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPolicySetOutcomes)(nil).Read), ctx, policySetOutcomeID)
}
