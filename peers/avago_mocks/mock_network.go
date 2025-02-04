// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/network (interfaces: Network)
//
// Generated by this command:
//
//	mockgen -destination=./avago_mocks/mock_network.go -package=avago_mocks github.com/ava-labs/avalanchego/network Network
//

// Package avago_mocks is a generated GoMock package.
package avago_mocks

import (
	context "context"
	netip "net/netip"
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	message "github.com/ava-labs/avalanchego/message"
	network "github.com/ava-labs/avalanchego/network"
	peer "github.com/ava-labs/avalanchego/network/peer"
	common "github.com/ava-labs/avalanchego/snow/engine/common"
	subnets "github.com/ava-labs/avalanchego/subnets"
	bloom "github.com/ava-labs/avalanchego/utils/bloom"
	ips "github.com/ava-labs/avalanchego/utils/ips"
	set "github.com/ava-labs/avalanchego/utils/set"
	gomock "go.uber.org/mock/gomock"
)

// MockNetwork is a mock of Network interface.
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
	isgomock struct{}
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork.
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance.
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// AllowConnection mocks base method.
func (m *MockNetwork) AllowConnection(peerID ids.NodeID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowConnection", peerID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowConnection indicates an expected call of AllowConnection.
func (mr *MockNetworkMockRecorder) AllowConnection(peerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowConnection", reflect.TypeOf((*MockNetwork)(nil).AllowConnection), peerID)
}

// Connected mocks base method.
func (m *MockNetwork) Connected(peerID ids.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connected", peerID)
}

// Connected indicates an expected call of Connected.
func (mr *MockNetworkMockRecorder) Connected(peerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockNetwork)(nil).Connected), peerID)
}

// Disconnected mocks base method.
func (m *MockNetwork) Disconnected(peerID ids.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnected", peerID)
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockNetworkMockRecorder) Disconnected(peerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockNetwork)(nil).Disconnected), peerID)
}

// Dispatch mocks base method.
func (m *MockNetwork) Dispatch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch")
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockNetworkMockRecorder) Dispatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockNetwork)(nil).Dispatch))
}

// HealthCheck mocks base method.
func (m *MockNetwork) HealthCheck(arg0 context.Context) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockNetworkMockRecorder) HealthCheck(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockNetwork)(nil).HealthCheck), arg0)
}

// KnownPeers mocks base method.
func (m *MockNetwork) KnownPeers() ([]byte, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KnownPeers")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// KnownPeers indicates an expected call of KnownPeers.
func (mr *MockNetworkMockRecorder) KnownPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownPeers", reflect.TypeOf((*MockNetwork)(nil).KnownPeers))
}

// ManuallyTrack mocks base method.
func (m *MockNetwork) ManuallyTrack(nodeID ids.NodeID, ip netip.AddrPort) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ManuallyTrack", nodeID, ip)
}

// ManuallyTrack indicates an expected call of ManuallyTrack.
func (mr *MockNetworkMockRecorder) ManuallyTrack(nodeID, ip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManuallyTrack", reflect.TypeOf((*MockNetwork)(nil).ManuallyTrack), nodeID, ip)
}

// NodeUptime mocks base method.
func (m *MockNetwork) NodeUptime() (network.UptimeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeUptime")
	ret0, _ := ret[0].(network.UptimeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeUptime indicates an expected call of NodeUptime.
func (mr *MockNetworkMockRecorder) NodeUptime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeUptime", reflect.TypeOf((*MockNetwork)(nil).NodeUptime))
}

// PeerInfo mocks base method.
func (m *MockNetwork) PeerInfo(nodeIDs []ids.NodeID) []peer.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerInfo", nodeIDs)
	ret0, _ := ret[0].([]peer.Info)
	return ret0
}

// PeerInfo indicates an expected call of PeerInfo.
func (mr *MockNetworkMockRecorder) PeerInfo(nodeIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerInfo", reflect.TypeOf((*MockNetwork)(nil).PeerInfo), nodeIDs)
}

// Peers mocks base method.
func (m *MockNetwork) Peers(peerID ids.NodeID, trackedSubnets set.Set[ids.ID], requestAllPeers bool, knownPeers *bloom.ReadFilter, peerSalt []byte) []*ips.ClaimedIPPort {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers", peerID, trackedSubnets, requestAllPeers, knownPeers, peerSalt)
	ret0, _ := ret[0].([]*ips.ClaimedIPPort)
	return ret0
}

// Peers indicates an expected call of Peers.
func (mr *MockNetworkMockRecorder) Peers(peerID, trackedSubnets, requestAllPeers, knownPeers, peerSalt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockNetwork)(nil).Peers), peerID, trackedSubnets, requestAllPeers, knownPeers, peerSalt)
}

// Send mocks base method.
func (m *MockNetwork) Send(msg message.OutboundMessage, config common.SendConfig, subnetID ids.ID, allower subnets.Allower) set.Set[ids.NodeID] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", msg, config, subnetID, allower)
	ret0, _ := ret[0].(set.Set[ids.NodeID])
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNetworkMockRecorder) Send(msg, config, subnetID, allower any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNetwork)(nil).Send), msg, config, subnetID, allower)
}

// StartClose mocks base method.
func (m *MockNetwork) StartClose() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartClose")
}

// StartClose indicates an expected call of StartClose.
func (mr *MockNetworkMockRecorder) StartClose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartClose", reflect.TypeOf((*MockNetwork)(nil).StartClose))
}

// Track mocks base method.
func (m *MockNetwork) Track(ips []*ips.ClaimedIPPort) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Track", ips)
	ret0, _ := ret[0].(error)
	return ret0
}

// Track indicates an expected call of Track.
func (mr *MockNetworkMockRecorder) Track(ips any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockNetwork)(nil).Track), ips)
}
