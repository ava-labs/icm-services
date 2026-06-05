// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
)

// defaultDeliveryQuorumNumerator matches the QUORUM_NUM used by MerkleValidatorSetRegistry
// (67/100) for verifying delivered messages.
const defaultDeliveryQuorumNumerator = 67

// ExternalEVMDestination configures an external EVM chain (e.g. Ethereum) for the relayer.
//
// A single entry can act as a validator-set updater target (keeps the on-chain registry's
// validator set fresh) and/or, when Deliver is set, as a TeleporterV2 message-delivery
// destination (the relayer submits receiveCrossChainMessage transactions to it).
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
	// Contract type: "subset" (default), "diff", or "merkle"
	ContractType string `mapstructure:"contract-type" json:"contract-type,omitempty"`
	// Minimum percentage of total weight that must change before submitting
	// an on-chain update (0–1 range, e.g. 0.05 = 5%). 0 means update on
	// every diff (legacy behavior).
	WeightChangeThresholdPct float64 `mapstructure:"weight-change-threshold-pct" json:"weight-change-threshold-pct,omitempty"` //nolint:lll
	// Maximum duration (in seconds) between on-chain updates. Even if the
	// weight change is below the threshold, an update is forced after this
	// interval. 0 means no staleness cap (legacy behavior).
	MaxUpdateIntervalSeconds uint64 `mapstructure:"max-update-interval-seconds" json:"max-update-interval-seconds,omitempty"` //nolint:lll
	// Maximum suggested gas price (in gwei) on the destination chain at which
	// the relayer will submit a validator-set update transaction. When the
	// network's suggested gas price exceeds this threshold the update is
	// deferred and retried on the next poll.
	MaxGasPriceGwei uint64 `mapstructure:"max-gas-price-gwei" json:"max-gas-price-gwei,omitempty"`

	// --- Message delivery configuration (only used when Deliver is true) ---

	// Deliver enables this external EVM chain as a TeleporterV2 message-delivery
	// destination. The relayer routes messages whose destination blockchain ID equals
	// DestinationBlockchainID to this chain.
	Deliver bool `mapstructure:"deliver" json:"deliver,omitempty"`
	// DestinationBlockchainID is the blockchain ID (cb58 or hex) by which TeleporterV2
	// messages address this external chain (the message's destinationBlockchainID field).
	DestinationBlockchainID string `mapstructure:"destination-blockchain-id" json:"destination-blockchain-id,omitempty"` //nolint:lll
	// TeleporterAddress is the TeleporterMessengerV2 contract address on the external
	// chain. With the universal deployer it is identical on the source chain.
	TeleporterAddress string `mapstructure:"teleporter-address" json:"teleporter-address,omitempty"`
	// QuorumNumerator is the stake-weight quorum (out of 100) required to verify a
	// delivered message. Defaults to 67 when unset.
	QuorumNumerator uint64 `mapstructure:"quorum-numerator" json:"quorum-numerator,omitempty"`
	// BlockGasLimit caps the gas used for delivery transactions.
	BlockGasLimit uint64 `mapstructure:"block-gas-limit" json:"block-gas-limit,omitempty"`
	// TxInclusionTimeoutSeconds bounds how long the relayer waits for a delivery tx to
	// be mined.
	TxInclusionTimeoutSeconds uint64 `mapstructure:"tx-inclusion-timeout-seconds" json:"tx-inclusion-timeout-seconds,omitempty"` //nolint:lll
}

// GetDestinationBlockchainID parses the configured external destination blockchain ID.
func (e *ExternalEVMDestination) GetDestinationBlockchainID() (ids.ID, error) {
	return ids.FromString(e.DestinationBlockchainID)
}

// GetWarpConfig returns the Warp configuration used when signing messages destined for this
// external EVM chain. The source subnet signs, so only the quorum numerator is meaningful.
func (e *ExternalEVMDestination) GetWarpConfig() WarpConfig {
	q := e.QuorumNumerator
	if q == 0 {
		q = defaultDeliveryQuorumNumerator
	}
	return WarpConfig{QuorumNumerator: q}
}

// ValidateDelivery validates and normalizes the fields required when this external destination
// is used for TeleporterV2 message delivery. It is a no-op when Deliver is false.
func (e *ExternalEVMDestination) ValidateDelivery() error {
	if !e.Deliver {
		return nil
	}
	if _, err := e.GetDestinationBlockchainID(); err != nil {
		return fmt.Errorf("invalid destination-blockchain-id %q: %w", e.DestinationBlockchainID, err)
	}
	if !common.IsHexAddress(e.TeleporterAddress) {
		return fmt.Errorf("invalid teleporter-address %q", e.TeleporterAddress)
	}
	if !common.IsHexAddress(e.ContractAddress) {
		return fmt.Errorf("invalid contract-address %q", e.ContractAddress)
	}
	if e.RPCEndpoint == "" {
		return errors.New("rpc-endpoint required for external EVM delivery destination")
	}
	if e.PrivateKey == "" {
		return errors.New("private-key required for external EVM delivery destination")
	}
	if e.QuorumNumerator == 0 {
		e.QuorumNumerator = defaultDeliveryQuorumNumerator
	}
	return nil
}
