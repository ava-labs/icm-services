// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/upgrade"

	"github.com/ava-labs/icm-services/peers/clients"
)

// NetworkInfo provides network metadata and upgrade information
type NetworkInfo struct {
	infoAPI              *clients.InfoAPI
	networkUpgradeConfig *upgrade.Config
}

// NewNetworkInfo creates a new network info provider
func NewNetworkInfo(
	ctx context.Context,
	cfg Config,
) (*NetworkInfo, error) {
	infoAPI, err := clients.NewInfoAPI(cfg.GetInfoAPI())
	if err != nil {
		return nil, fmt.Errorf("failed to create info API: %w", err)
	}
	upgradeConfig, err := infoAPI.Upgrades(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get upgrades: %w", err)
	}
	return &NetworkInfo{
		infoAPI:              infoAPI,
		networkUpgradeConfig: upgradeConfig,
	}, nil
}

// IsGraniteActivated returns whether the Granite upgrade is activated
func (n *NetworkInfo) IsGraniteActivated() bool {
	return n.networkUpgradeConfig.IsGraniteActivated(time.Now())
}

// GetGraniteEpochDuration returns the Granite epoch duration from the network upgrade config.
// Returns 0 if Granite is not activated or epoch duration is not configured.
func (n *NetworkInfo) GetGraniteEpochDuration() time.Duration {
	if n.networkUpgradeConfig == nil {
		return 0
	}
	return n.networkUpgradeConfig.GraniteEpochDuration
}
