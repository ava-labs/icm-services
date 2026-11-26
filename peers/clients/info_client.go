// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package clients

import (
	"context"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/icm-services/config"
)

// InfoAPI is a wrapper around the info.Client,
// and provides additional options for the API
// passed in the config.
type InfoAPI struct {
	client  info.Client
	options []rpc.Option
}

func NewInfoAPI(apiConfig *config.APIConfig) (*InfoAPI, error) {
	client := info.NewClient(apiConfig.BaseURL)
	options := apiConfig.Options()
	return &InfoAPI{
		client:  *client,
		options: options,
	}, nil
}

func (i *InfoAPI) GetNetworkID(ctx context.Context) (uint32, error) {
	return i.client.GetNetworkID(ctx, i.options...)
}

func (i *InfoAPI) Peers(ctx context.Context, nodeIDs []ids.NodeID) ([]info.Peer, error) {
	return i.client.Peers(ctx, nodeIDs, i.options...)
}

func (i *InfoAPI) Upgrades(ctx context.Context) (*upgrade.Config, error) {
	return i.client.Upgrades(ctx, i.options...)
}
