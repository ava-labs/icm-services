// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_app_request_network.go -package=mocks

package peers

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/set"
)

// AppRequestNetworkTestHelper is an interface for testing purposes only.
// It allows tests to mock the AppRequestNetwork functionality.
type AppRequestNetworkTestHelper interface {
	GetCanonicalValidators(
		ctx context.Context,
		subnetID ids.ID,
		skipCache bool,
		pchainHeight uint64,
	) (*CanonicalValidators, error)
	GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error)
	RegisterAppRequest(requestID ids.RequestID)
	RegisterRequestID(
		requestID uint32,
		requestedNodes set.Set[ids.NodeID],
	) chan message.InboundMessage
	Send(
		msg message.OutboundMessage,
		nodeIDs set.Set[ids.NodeID],
		subnetID ids.ID,
		allower subnets.Allower,
	) set.Set[ids.NodeID]
	TrackSubnet(ctx context.Context, subnetID ids.ID)
}

// Verify that AppRequestNetwork implements the test helper interface
var _ AppRequestNetworkTestHelper = (*AppRequestNetwork)(nil)
