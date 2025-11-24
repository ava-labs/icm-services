// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"context"

	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/proposervm"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/peers/utils"
)

// ProposerVMAPI is a wrapper around a [proposervm.JSONRPCClient],
// and provides additional options for the API passed in the config.
type ProposerVMAPI struct {
	client  *proposervm.JSONRPCClient
	options []rpc.Option
}

func NewProposerVMAPI(uri string, chain string, cfg *config.APIConfig) *ProposerVMAPI {
	return &ProposerVMAPI{
		client:  proposervm.NewJSONRPCClient(uri, chain),
		options: utils.InitializeOptions(cfg),
	}
}

func (p *ProposerVMAPI) GetCurrentEpoch(ctx context.Context) (block.Epoch, error) {
	return p.client.GetCurrentEpoch(ctx, p.options...)
}

func (p *ProposerVMAPI) GetProposedHeight(ctx context.Context) (uint64, error) {
	return p.client.GetProposedHeight(ctx, p.options...)
}
