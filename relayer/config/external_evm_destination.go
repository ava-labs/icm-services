// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
)

// DefaultRegistryQuorumNumerator matches the QUORUM_NUM used by MerkleValidatorSetRegistry
// (67/100) for verifying delivered messages and validator-set attestations. Signatures
// aggregated with a lower quorum are rejected on-chain, so this is a hard requirement of the
// registry contract rather than a configurable policy default.
const DefaultRegistryQuorumNumerator uint64 = 67

// defaultPollInterval is how often the validator-set updater polls when
// PollIntervalSeconds is unset.
const defaultPollInterval = 10 * time.Second

// ExternalEVMDestination configures an external EVM chain (e.g. Ethereum) for the relayer.
//
// A single entry can act as a validator-set updater target (keeps the on-chain registry's
// validator set fresh) and, when DestinationBlockchainID is set, as a TeleporterV2
// message-delivery destination (the relayer submits receiveCrossChainMessage transactions to it).
type ExternalEVMDestination struct {
	// RPC endpoint of the external EVM chain (e.g. local geth node)
	RPCEndpoint string `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	// Hex-encoded private key used by the validator-set updater to sign
	// registerValidatorSet/updateValidatorSet transactions.
	PrivateKey string `mapstructure:"private-key" json:"private-key" sensitive:"true"`
	// Address of the deployed updater contract
	ContractAddress string `mapstructure:"contract-address" json:"contract-address"`
	// The blockchain ID (on the Avalanche side) whose validator set to track
	BlockchainID string `mapstructure:"blockchain-id" json:"blockchain-id"`
	// The subnet ID that the blockchain belongs to
	SubnetID string `mapstructure:"subnet-id" json:"subnet-id"`
	// Poll interval in seconds (default 10)
	PollIntervalSeconds uint64 `mapstructure:"poll-interval-seconds" json:"poll-interval-seconds"`
	// Maximum duration (in seconds) between on-chain updates. Even if the
	// weight change is below the threshold, an update is forced after this
	// interval. 0 means no staleness cap (legacy behavior).
	MaxUpdateIntervalSeconds uint64 `mapstructure:"max-update-interval-seconds" json:"max-update-interval-seconds,omitempty"` //nolint:lll
	// Maximum suggested gas price (in gwei) on the destination chain at which
	// the relayer will submit a validator-set update transaction. When the
	// network's suggested gas price exceeds this threshold the update is
	// deferred and retried on the next poll.
	MaxGasPriceGwei uint64 `mapstructure:"max-gas-price-gwei" json:"max-gas-price-gwei,omitempty"`

	// --- Message delivery configuration (only used when DestinationBlockchainID is set) ---

	// DestinationBlockchainID is the blockchain ID (cb58 or hex) by which TeleporterV2
	// messages address this external chain (the message's destinationBlockchainID field).
	// Setting it enables this external EVM chain as a TeleporterV2 message-delivery
	// destination; leaving it empty makes the entry a validator-set updater only.
	DestinationBlockchainID string `mapstructure:"destination-blockchain-id" json:"destination-blockchain-id,omitempty"` //nolint:lll
	// DeliveryPrivateKey is the hex-encoded private key used by the relayer to sign
	// message-delivery (receiveCrossChainMessage) transactions. It must be different from
	// PrivateKey so the validator-set updater and the message-delivery client never share a
	// sender account: sharing an account causes their nonces to collide on-chain. Required
	// when DestinationBlockchainID is set.
	DeliveryPrivateKey string `mapstructure:"delivery-private-key" json:"delivery-private-key,omitempty" sensitive:"true"` //nolint:lll
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

// DeliversMessages reports whether this external destination is configured as a TeleporterV2
// message-delivery target, which is enabled by setting DestinationBlockchainID. When empty the
// entry acts as a validator-set updater only.
func (e *ExternalEVMDestination) DeliversMessages() bool {
	return e.DestinationBlockchainID != ""
}

// GetPollInterval returns how often the validator-set updater should poll for validator set
// changes, falling back to defaultPollInterval when PollIntervalSeconds is unset.
func (e *ExternalEVMDestination) GetPollInterval() time.Duration {
	if e.PollIntervalSeconds == 0 {
		return defaultPollInterval
	}
	return time.Duration(e.PollIntervalSeconds) * time.Second
}

// GetWarpConfig returns the Warp configuration used when signing messages destined for this
// external EVM chain. The source subnet signs, so only the quorum numerator is meaningful.
func (e *ExternalEVMDestination) GetWarpConfig() WarpConfig {
	q := e.QuorumNumerator
	if q == 0 {
		q = DefaultRegistryQuorumNumerator
	}
	return WarpConfig{QuorumNumerator: q}
}

// ValidateDelivery validates and normalizes the fields required when this external destination
// is used for TeleporterV2 message delivery. It is a no-op when delivery is not configured.
func (e *ExternalEVMDestination) ValidateDelivery() error {
	if !e.DeliversMessages() {
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
	if e.DeliveryPrivateKey == "" {
		return errors.New("delivery-private-key required for external EVM delivery destination")
	}
	// The validator-set updater (private-key) and the message-delivery client
	// (delivery-private-key) must use different sender accounts; sharing one account
	// makes their nonces collide on-chain.
	if normalizeHexKey(e.DeliveryPrivateKey) == normalizeHexKey(e.PrivateKey) {
		return errors.New("delivery-private-key must be different from private-key")
	}
	if e.QuorumNumerator == 0 {
		e.QuorumNumerator = DefaultRegistryQuorumNumerator
	}
	return nil
}

// normalizeHexKey lowercases a hex-encoded private key and strips an optional "0x" prefix so
// keys can be compared regardless of casing or prefix.
func normalizeHexKey(key string) string {
	return strings.TrimPrefix(strings.ToLower(key), "0x")
}
