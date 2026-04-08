// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

// ExternalEVMDestination configures a validator set updater for an external EVM chain.
type ExternalEVMDestination struct {
	// RPC endpoint of the external EVM chain (e.g. local geth node)
	RPCEndpoint string `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	// Hex-encoded private key for signing transactions
	PrivateKey string `mapstructure:"private-key" json:"private-key" sensitive:"true"`
	// Address of the deployed updater contract
	ContractAddress string `mapstructure:"contract-address" json:"contract-address"`
	// The blockchain ID (on the Avalanche side) whose validator set to track
	BlockchainID string `mapstructure:"blockchain-id" json:"blockchain-id"`
	// The subnet ID that the blockchain belongs to
	SubnetID string `mapstructure:"subnet-id" json:"subnet-id"`
	// Number of validators per shard (default 10)
	ShardSize uint32 `mapstructure:"shard-size" json:"shard-size"`
	// Poll interval in seconds (default 10)
	PollIntervalSeconds uint64 `mapstructure:"poll-interval-seconds" json:"poll-interval-seconds"`
	// Contract type: "subset" (default) or "diff"
	ContractType string `mapstructure:"contract-type" json:"contract-type,omitempty"`
	// Minimum percentage of total weight that must change before submitting an
	// on-chain update (0–1 range, e.g. 0.05 = 5%). 0 means update on every diff
	// (legacy behavior).
	WeightChangeThresholdPct float64 `mapstructure:"weight-change-threshold-pct" json:"weight-change-threshold-pct,omitempty"`
	// Maximum duration (in seconds) between on-chain updates. Even if the weight
	// change is below the threshold, an update is forced after this interval.
	// 0 means no staleness cap (legacy behavior).
	MaxUpdateIntervalSeconds uint64 `mapstructure:"max-update-interval-seconds" json:"max-update-interval-seconds,omitempty"`
}
