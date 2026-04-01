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
}
