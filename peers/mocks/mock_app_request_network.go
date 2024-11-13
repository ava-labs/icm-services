// Code generated by MockGen. DO NOT EDIT.
// Source: app_request_network.go
//
// Generated by this command:
//
//	mockgen -source=app_request_network.go -destination=./mocks/mock_app_request_network.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	message "github.com/ava-labs/avalanchego/message"
	subnets "github.com/ava-labs/avalanchego/subnets"
	set "github.com/ava-labs/avalanchego/utils/set"
	peers "github.com/ava-labs/icm-relayer/peers"
	gomock "go.uber.org/mock/gomock"
)

// MockAppRequestNetwork is a mock of AppRequestNetwork interface.
type MockAppRequestNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockAppRequestNetworkMockRecorder
	isgomock struct{}
}

// MockAppRequestNetworkMockRecorder is the mock recorder for MockAppRequestNetwork.
type MockAppRequestNetworkMockRecorder struct {
	mock *MockAppRequestNetwork
}

// NewMockAppRequestNetwork creates a new mock instance.
func NewMockAppRequestNetwork(ctrl *gomock.Controller) *MockAppRequestNetwork {
	mock := &MockAppRequestNetwork{ctrl: ctrl}
	mock.recorder = &MockAppRequestNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppRequestNetwork) EXPECT() *MockAppRequestNetworkMockRecorder {
	return m.recorder
}

// ConnectPeers mocks base method.
func (m *MockAppRequestNetwork) ConnectPeers(nodeIDs set.Set[ids.NodeID]) set.Set[ids.NodeID] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectPeers", nodeIDs)
	ret0, _ := ret[0].(set.Set[ids.NodeID])
	return ret0
}

// ConnectPeers indicates an expected call of ConnectPeers.
func (mr *MockAppRequestNetworkMockRecorder) ConnectPeers(nodeIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectPeers", reflect.TypeOf((*MockAppRequestNetwork)(nil).ConnectPeers), nodeIDs)
}

// ConnectToCanonicalValidators mocks base method.
func (m *MockAppRequestNetwork) ConnectToCanonicalValidators(subnetID ids.ID) (*peers.ConnectedCanonicalValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToCanonicalValidators", subnetID)
	ret0, _ := ret[0].(*peers.ConnectedCanonicalValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToCanonicalValidators indicates an expected call of ConnectToCanonicalValidators.
func (mr *MockAppRequestNetworkMockRecorder) ConnectToCanonicalValidators(subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToCanonicalValidators", reflect.TypeOf((*MockAppRequestNetwork)(nil).ConnectToCanonicalValidators), subnetID)
}

// GetSubnetID mocks base method.
func (m *MockAppRequestNetwork) GetSubnetID(blockchainID ids.ID) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetID", blockchainID)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetID indicates an expected call of GetSubnetID.
func (mr *MockAppRequestNetworkMockRecorder) GetSubnetID(blockchainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetID", reflect.TypeOf((*MockAppRequestNetwork)(nil).GetSubnetID), blockchainID)
}

// RegisterAppRequest mocks base method.
func (m *MockAppRequestNetwork) RegisterAppRequest(requestID ids.RequestID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAppRequest", requestID)
}

// RegisterAppRequest indicates an expected call of RegisterAppRequest.
func (mr *MockAppRequestNetworkMockRecorder) RegisterAppRequest(requestID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAppRequest", reflect.TypeOf((*MockAppRequestNetwork)(nil).RegisterAppRequest), requestID)
}

// RegisterRequestID mocks base method.
func (m *MockAppRequestNetwork) RegisterRequestID(requestID uint32, numExpectedResponse int) chan message.InboundMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRequestID", requestID, numExpectedResponse)
	ret0, _ := ret[0].(chan message.InboundMessage)
	return ret0
}

// RegisterRequestID indicates an expected call of RegisterRequestID.
func (mr *MockAppRequestNetworkMockRecorder) RegisterRequestID(requestID, numExpectedResponse any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRequestID", reflect.TypeOf((*MockAppRequestNetwork)(nil).RegisterRequestID), requestID, numExpectedResponse)
}

// Send mocks base method.
func (m *MockAppRequestNetwork) Send(msg message.OutboundMessage, nodeIDs set.Set[ids.NodeID], subnetID ids.ID, allower subnets.Allower) set.Set[ids.NodeID] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", msg, nodeIDs, subnetID, allower)
	ret0, _ := ret[0].(set.Set[ids.NodeID])
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAppRequestNetworkMockRecorder) Send(msg, nodeIDs, subnetID, allower any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAppRequestNetwork)(nil).Send), msg, nodeIDs, subnetID, allower)
}
