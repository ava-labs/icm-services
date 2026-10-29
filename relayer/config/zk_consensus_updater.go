// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/libevm/common"
)

// ZKConsensusUpdater configures a consensus updater for a ZKAdapter deployed on a destination
// chain. The updater polls the Boundless subgraph for Ethereum consensus proofs and submits them
// to the adapter as transitions, keeping its confirmed beacon anchors current so that zk messages
// from Ethereum can be proven against them.
type ZKConsensusUpdater struct {
	// RPC endpoint of the destination chain the ZKAdapter is deployed on.
	RPCEndpoint string `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	// Hex-encoded private key used to sign transition transactions. It must not be shared with
	// a message-delivery account, or their nonces collide on-chain.
	PrivateKey string `mapstructure:"private-key" json:"private-key" sensitive:"true"`
	// Address of the deployed ZKAdapter contract.
	AdapterAddress string `mapstructure:"adapter-address" json:"adapter-address"`
	// GraphQL endpoint of the Boundless subgraph indexing consensus proofs.
	SubgraphURL string `mapstructure:"subgraph-url" json:"subgraph-url"`
	// Poll interval in seconds (default 10)
	PollIntervalSeconds uint64 `mapstructure:"poll-interval-seconds" json:"poll-interval-seconds"`
	// Maximum suggested gas price (in gwei) on the destination chain at which the updater will
	// submit a transition. When exceeded the transition is deferred to the next poll. 0 disables
	// the check.
	MaxGasPriceGwei uint64 `mapstructure:"max-gas-price-gwei" json:"max-gas-price-gwei,omitempty"`
}

// Validate checks the fields required to run the consensus updater.
func (z *ZKConsensusUpdater) Validate() error {
	if z.RPCEndpoint == "" {
		return errors.New("rpc-endpoint required for zk consensus updater")
	}
	if z.PrivateKey == "" {
		return errors.New("private-key required for zk consensus updater")
	}
	if !common.IsHexAddress(z.AdapterAddress) {
		return fmt.Errorf("invalid adapter-address %q", z.AdapterAddress)
	}
	if z.SubgraphURL == "" {
		return errors.New("subgraph-url required for zk consensus updater")
	}
	return nil
}

// GetPollInterval returns how often the updater polls the subgraph for new proofs, falling
// back to defaultPollInterval when PollIntervalSeconds is unset.
func (z *ZKConsensusUpdater) GetPollInterval() time.Duration {
	if z.PollIntervalSeconds == 0 {
		return defaultPollInterval
	}
	return time.Duration(z.PollIntervalSeconds) * time.Second
}
