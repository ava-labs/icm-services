// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"context"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rpc"
	"github.com/ava-labs/subnet-evm/params"
)

var _ configRPCClient = (*rpcClient)(nil)

type configRPCClient interface {
	ChainConfig(ctx context.Context) (*params.ChainConfigWithUpgradesJSON, error)
	LatestHeader(ctx context.Context) (*types.Header, error)
}

type rpcClient struct {
	c *rpc.Client
}

func (rc *rpcClient) ChainConfig(ctx context.Context) (*params.ChainConfigWithUpgradesJSON, error) {
	var result *params.ChainConfigWithUpgradesJSON
	err := rc.c.CallContext(ctx, &result, "eth_getChainConfig")
	if err != nil {
		return nil, err
	}
	return result, err
}

func (rc *rpcClient) LatestHeader(ctx context.Context) (*types.Header, error) {
	var header *types.Header
	err := rc.c.CallContext(ctx, &header, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		err = ethereum.NotFound
	}
	return header, err
}
