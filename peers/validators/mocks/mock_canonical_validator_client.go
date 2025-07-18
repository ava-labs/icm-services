// Code generated by MockGen. DO NOT EDIT.
// Source: canonical_validator_client.go
//
// Generated by this command:
//
//	mockgen -source=canonical_validator_client.go -destination=./mocks/mock_canonical_validator_client.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	validators "github.com/ava-labs/avalanchego/snow/validators"
	warp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	gomock "go.uber.org/mock/gomock"
)

// MockCanonicalValidatorState is a mock of CanonicalValidatorState interface.
type MockCanonicalValidatorState struct {
	ctrl     *gomock.Controller
	recorder *MockCanonicalValidatorStateMockRecorder
	isgomock struct{}
}

// MockCanonicalValidatorStateMockRecorder is the mock recorder for MockCanonicalValidatorState.
type MockCanonicalValidatorStateMockRecorder struct {
	mock *MockCanonicalValidatorState
}

// NewMockCanonicalValidatorState creates a new mock instance.
func NewMockCanonicalValidatorState(ctrl *gomock.Controller) *MockCanonicalValidatorState {
	mock := &MockCanonicalValidatorState{ctrl: ctrl}
	mock.recorder = &MockCanonicalValidatorStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCanonicalValidatorState) EXPECT() *MockCanonicalValidatorStateMockRecorder {
	return m.recorder
}

// GetCurrentCanonicalValidatorSet mocks base method.
func (m *MockCanonicalValidatorState) GetCurrentCanonicalValidatorSet(ctx context.Context, subnetID ids.ID) (warp.CanonicalValidatorSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentCanonicalValidatorSet", ctx, subnetID)
	ret0, _ := ret[0].(warp.CanonicalValidatorSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentCanonicalValidatorSet indicates an expected call of GetCurrentCanonicalValidatorSet.
func (mr *MockCanonicalValidatorStateMockRecorder) GetCurrentCanonicalValidatorSet(ctx, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentCanonicalValidatorSet", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetCurrentCanonicalValidatorSet), ctx, subnetID)
}

// GetCurrentHeight mocks base method.
func (m *MockCanonicalValidatorState) GetCurrentHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentHeight indicates an expected call of GetCurrentHeight.
func (mr *MockCanonicalValidatorStateMockRecorder) GetCurrentHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentHeight", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetCurrentHeight), arg0)
}

// GetCurrentValidatorSet mocks base method.
func (m *MockCanonicalValidatorState) GetCurrentValidatorSet(ctx context.Context, subnetID ids.ID) (map[ids.ID]*validators.GetCurrentValidatorOutput, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentValidatorSet", ctx, subnetID)
	ret0, _ := ret[0].(map[ids.ID]*validators.GetCurrentValidatorOutput)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCurrentValidatorSet indicates an expected call of GetCurrentValidatorSet.
func (mr *MockCanonicalValidatorStateMockRecorder) GetCurrentValidatorSet(ctx, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentValidatorSet", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetCurrentValidatorSet), ctx, subnetID)
}

// GetMinimumHeight mocks base method.
func (m *MockCanonicalValidatorState) GetMinimumHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinimumHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinimumHeight indicates an expected call of GetMinimumHeight.
func (mr *MockCanonicalValidatorStateMockRecorder) GetMinimumHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinimumHeight", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetMinimumHeight), arg0)
}

// GetProposedValidators mocks base method.
func (m *MockCanonicalValidatorState) GetProposedValidators(ctx context.Context, subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposedValidators", ctx, subnetID)
	ret0, _ := ret[0].(map[ids.NodeID]*validators.GetValidatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposedValidators indicates an expected call of GetProposedValidators.
func (mr *MockCanonicalValidatorStateMockRecorder) GetProposedValidators(ctx, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposedValidators", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetProposedValidators), ctx, subnetID)
}

// GetSubnetID mocks base method.
func (m *MockCanonicalValidatorState) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetID", ctx, chainID)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetID indicates an expected call of GetSubnetID.
func (mr *MockCanonicalValidatorStateMockRecorder) GetSubnetID(ctx, chainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetID", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetSubnetID), ctx, chainID)
}

// GetValidatorSet mocks base method.
func (m *MockCanonicalValidatorState) GetValidatorSet(ctx context.Context, height uint64, subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSet", ctx, height, subnetID)
	ret0, _ := ret[0].(map[ids.NodeID]*validators.GetValidatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorSet indicates an expected call of GetValidatorSet.
func (mr *MockCanonicalValidatorStateMockRecorder) GetValidatorSet(ctx, height, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSet", reflect.TypeOf((*MockCanonicalValidatorState)(nil).GetValidatorSet), ctx, height, subnetID)
}
