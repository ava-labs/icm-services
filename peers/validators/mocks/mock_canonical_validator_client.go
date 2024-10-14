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
	gomock "go.uber.org/mock/gomock"
)

// MockCanonicalValidatorClient is a mock of CanonicalValidatorClient interface.
type MockCanonicalValidatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockCanonicalValidatorClientMockRecorder
}

// MockCanonicalValidatorClientMockRecorder is the mock recorder for MockCanonicalValidatorClient.
type MockCanonicalValidatorClientMockRecorder struct {
	mock *MockCanonicalValidatorClient
}

// NewMockCanonicalValidatorClient creates a new mock instance.
func NewMockCanonicalValidatorClient(ctrl *gomock.Controller) *MockCanonicalValidatorClient {
	mock := &MockCanonicalValidatorClient{ctrl: ctrl}
	mock.recorder = &MockCanonicalValidatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCanonicalValidatorClient) EXPECT() *MockCanonicalValidatorClientMockRecorder {
	return m.recorder
}

// GetBlockByHeight mocks base method.
func (m *MockCanonicalValidatorClient) GetBlockByHeight(arg0 context.Context, arg1 uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHeight", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight.
func (mr *MockCanonicalValidatorClientMockRecorder) GetBlockByHeight(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockCanonicalValidatorClient)(nil).GetBlockByHeight), arg0, arg1)
}

// GetCurrentHeight mocks base method.
func (m *MockCanonicalValidatorClient) GetCurrentHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentHeight indicates an expected call of GetCurrentHeight.
func (mr *MockCanonicalValidatorClientMockRecorder) GetCurrentHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentHeight", reflect.TypeOf((*MockCanonicalValidatorClient)(nil).GetCurrentHeight), arg0)
}

// GetMinimumHeight mocks base method.
func (m *MockCanonicalValidatorClient) GetMinimumHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinimumHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinimumHeight indicates an expected call of GetMinimumHeight.
func (mr *MockCanonicalValidatorClientMockRecorder) GetMinimumHeight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinimumHeight", reflect.TypeOf((*MockCanonicalValidatorClient)(nil).GetMinimumHeight), arg0)
}

// GetSubnetID mocks base method.
func (m *MockCanonicalValidatorClient) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetID", ctx, chainID)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetID indicates an expected call of GetSubnetID.
func (mr *MockCanonicalValidatorClientMockRecorder) GetSubnetID(ctx, chainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetID", reflect.TypeOf((*MockCanonicalValidatorClient)(nil).GetSubnetID), ctx, chainID)
}

// GetValidatorSet mocks base method.
func (m *MockCanonicalValidatorClient) GetValidatorSet(ctx context.Context, height uint64, subnetID ids.ID) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSet", ctx, height, subnetID)
	ret0, _ := ret[0].(map[ids.NodeID]*validators.GetValidatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorSet indicates an expected call of GetValidatorSet.
func (mr *MockCanonicalValidatorClientMockRecorder) GetValidatorSet(ctx, height, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSet", reflect.TypeOf((*MockCanonicalValidatorClient)(nil).GetValidatorSet), ctx, height, subnetID)
}
