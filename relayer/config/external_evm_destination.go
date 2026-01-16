// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/utils/set"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
)

const (
	// Default values for external EVM destinations
	defaultExternalEVMBlockGasLimit             = 15_000_000
	defaultExternalEVMMaxPriorityFeePerGas      = 2500000000 // 2.5 gwei
	defaultExternalEVMTxInclusionTimeoutSeconds = 60
)

// ExternalEVMDestination configures an external EVM chain (e.g., Ethereum)
// where validator set updates should be posted.
type ExternalEVMDestination struct {
	// ChainID is the EVM chain ID (e.g., "1" for Ethereum mainnet, "11155111" for Sepolia)
	ChainID string `mapstructure:"chain-id" json:"chain-id"`

	// RPCEndpoint is the RPC endpoint for the external EVM chain
	RPCEndpoint basecfg.APIConfig `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`

	// RegistryAddress is the address of the AvalancheValidatorSetRegistry contract
	RegistryAddress string `mapstructure:"registry-address" json:"registry-address"`

	// AccountPrivateKey is the private key for signing transactions (single key)
	AccountPrivateKey string `mapstructure:"account-private-key" json:"account-private-key" sensitive:"true"`

	// AccountPrivateKeys is a list of private keys for signing transactions (multiple keys)
	AccountPrivateKeys []string `mapstructure:"account-private-keys-list" json:"account-private-keys-list" sensitive:"true"`

	// BlockGasLimit is the maximum gas limit for transactions
	BlockGasLimit uint64 `mapstructure:"block-gas-limit" json:"block-gas-limit"`

	// MaxBaseFee is the maximum base fee (in wei) to pay for transactions
	MaxBaseFee uint64 `mapstructure:"max-base-fee" json:"max-base-fee"`

	// SuggestedPriorityFeeBuffer is the buffer to add to suggested priority fee (in wei)
	SuggestedPriorityFeeBuffer uint64 `mapstructure:"suggested-priority-fee-buffer" json:"suggested-priority-fee-buffer"`

	// MaxPriorityFeePerGas is the maximum priority fee per gas (in wei)
	MaxPriorityFeePerGas uint64 `mapstructure:"max-priority-fee-per-gas" json:"max-priority-fee-per-gas"`

	// TxInclusionTimeoutSeconds is the timeout for transaction inclusion
	TxInclusionTimeoutSeconds uint64 `mapstructure:"tx-inclusion-timeout-seconds" json:"tx-inclusion-timeout-seconds"`

	// Parsed/validated fields
	registryAddress common.Address
}

// Validate validates the external EVM destination configuration.
func (e *ExternalEVMDestination) Validate() error {
	if e.ChainID == "" {
		return errors.New("chain-id is required for external EVM destination")
	}

	if err := e.RPCEndpoint.Validate(); err != nil {
		return fmt.Errorf("invalid rpc-endpoint in external EVM destination: %w", err)
	}

	if e.RegistryAddress == "" {
		return errors.New("registry-address is required for external EVM destination")
	}

	// Validate registry address
	if !common.IsHexAddress(e.RegistryAddress) {
		return fmt.Errorf("invalid registry-address: %s", e.RegistryAddress)
	}
	e.registryAddress = common.HexToAddress(e.RegistryAddress)

	// Collect and deduplicate the account private keys
	privateKeys := e.AccountPrivateKeys
	if e.AccountPrivateKey != "" {
		privateKeys = append(privateKeys, e.AccountPrivateKey)
	}

	for i, pkey := range privateKeys {
		if _, err := crypto.HexToECDSA(utils.SanitizeHexString(pkey)); err != nil {
			return utils.ErrInvalidPrivateKeyHex
		}
		privateKeys[i] = utils.SanitizeHexString(pkey)
	}
	uniquePks := set.Of(privateKeys...)

	if len(uniquePks) == 0 {
		return errors.New("no account private keys provided for external EVM destination")
	}
	e.AccountPrivateKeys = uniquePks.List()

	// Set defaults
	if e.BlockGasLimit == 0 {
		e.BlockGasLimit = defaultExternalEVMBlockGasLimit
	}
	if e.MaxPriorityFeePerGas == 0 {
		e.MaxPriorityFeePerGas = defaultExternalEVMMaxPriorityFeePerGas
	}
	if e.TxInclusionTimeoutSeconds == 0 {
		e.TxInclusionTimeoutSeconds = defaultExternalEVMTxInclusionTimeoutSeconds
	}

	return nil
}

// GetRegistryAddress returns the parsed registry address.
func (e *ExternalEVMDestination) GetRegistryAddress() common.Address {
	return e.registryAddress
}

const (
	// DefaultValidatorSetUpdaterPollIntervalSeconds is the default polling interval
	// for checking validator set changes when external EVM destinations are configured.
	DefaultValidatorSetUpdaterPollIntervalSeconds = 5 // 1 minute
)
