// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/platformvm"
)

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_p_chain_client.go -package=mocks

var _ PChainClient = &platformvm.Client{}

type PChainClient interface {
	GetCurrentValidators(
		ctx context.Context,
		subnetID ids.ID,
		nodeIDs []ids.NodeID,
		options ...rpc.Option,
	) ([]platformvm.ClientPermissionlessValidator, error)
	GetSubnet(ctx context.Context, subnetID ids.ID, options ...rpc.Option) (platformvm.GetSubnetClientResponse, error)
}
