// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"fmt"

	"github.com/ava-labs/libevm/common"
)

// Config holds the settings for the TeleporterV2 merkle message handler.
//
// Unlike the warp-precompile based handler, regular ICM messages are verified on the
// destination chain by a MerkleValidatorSetRegistry that acts as the TeleporterMessengerV2
// adapter. The relayer attaches a Merkle attestation (signers + multi-proof + aggregate BLS
// signature) instead of riding the signed warp message in the transaction access list.
type Config struct {
	// RewardAddress is the address credited as the relayer reward recipient on the source chain.
	RewardAddress string `json:"reward-address"`
	// RegistryAddress is the MerkleValidatorSetRegistry deployed on the destination chain. It is
	// queried for the committed P-chain height so the relayer builds the attestation against the
	// exact validator set committed under the stored Merkle root. This is the same contract that
	// acts as the TeleporterMessengerV2 adapter and is the warp message's origin sender.
	RegistryAddress string `json:"registry-address"`
	// TeleporterAddress is the TeleporterMessengerV2 contract on the destination chain. It is a
	// different contract from the registry/adapter and is the target of receiveCrossChainMessage,
	// the contract used to compute the Teleporter message ID, and the one queried for delivery
	// status. With the universal deployer it is identical on the source chain.
	TeleporterAddress string `json:"teleporter-address"`
}

func ConfigFromMap(m map[string]any) (*Config, error) {
	rewardAddress, ok := m["reward-address"].(string)
	if !ok {
		return nil, fmt.Errorf("reward-address not found")
	}
	if !common.IsHexAddress(rewardAddress) {
		return nil, fmt.Errorf("invalid reward address: %s", rewardAddress)
	}

	registryAddress, ok := m["registry-address"].(string)
	if !ok {
		return nil, fmt.Errorf("registry-address not found")
	}
	if !common.IsHexAddress(registryAddress) {
		return nil, fmt.Errorf("invalid registry address: %s", registryAddress)
	}

	teleporterAddress, ok := m["teleporter-address"].(string)
	if !ok {
		return nil, fmt.Errorf("teleporter-address not found")
	}
	if !common.IsHexAddress(teleporterAddress) {
		return nil, fmt.Errorf("invalid teleporter address: %s", teleporterAddress)
	}

	return &Config{
		RewardAddress:     rewardAddress,
		RegistryAddress:   registryAddress,
		TeleporterAddress: teleporterAddress,
	}, nil
}

func (c *Config) registryAddress() common.Address {
	return common.HexToAddress(c.RegistryAddress)
}

func (c *Config) teleporterAddress() common.Address {
	return common.HexToAddress(c.TeleporterAddress)
}

func (c *Config) rewardAddress() common.Address {
	return common.HexToAddress(c.RewardAddress)
}
