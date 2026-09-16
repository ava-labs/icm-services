// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"fmt"

	"github.com/ava-labs/libevm/common"
)

// Verifier types selectable via the "verifier-type" setting. The verifier type must match the
// adapter contract the destination TeleporterMessengerV2 was constructed with, since it
// determines both how the source message payload is decoded and how the relayer builds the
// attestation the adapter verifies.
const (
	// VerifierTypeMerkle verifies messages against a MerkleValidatorSetRegistry adapter using a
	// Merkle attestation (signers + multi-proof + aggregate BLS signature).
	VerifierTypeMerkle = "merkle"
	// VerifierTypeWarp verifies messages with the WarpAdapter, which reads the signed warp
	// message from the transaction predicate via the warp precompile.
	VerifierTypeWarp = "warp"
)

// Config holds the settings for the TeleporterV2 message handler.
//
// TeleporterMessengerV2 delegates message verification to an adapter contract, so the way the
// relayer attests to a message depends on which adapter the destination messenger was deployed
// with. The "verifier-type" setting selects that path:
//
//   - "merkle": regular ICM messages are verified on the destination chain by a
//     MerkleValidatorSetRegistry that acts as the TeleporterMessengerV2 adapter. The relayer
//     attaches a Merkle attestation (signers + multi-proof + aggregate BLS signature) instead of
//     reading the signed warp message in the transaction access list.
//   - "warp": messages are verified by the WarpAdapter through the warp precompile. The relayer
//     includes the signed warp message as a transaction predicate and the attestation is the
//     index of that message in the predicate.
type Config struct {
	// RewardAddress is the address credited as the relayer reward recipient on the source chain.
	RewardAddress string `json:"reward-address"`
	// VerifierType selects the verification path ("merkle" or "warp"). Defaults to "merkle" when
	// unset.
	VerifierType string `json:"verifier-type"`
	// RegistryAddress is the MerkleValidatorSetRegistry deployed on the destination chain. It is
	// queried for the committed P-chain height so the relayer builds the attestation against the
	// exact validator set committed under the stored Merkle root. This is the same contract that
	// acts as the TeleporterMessengerV2 adapter and is the warp message's origin sender.
	// Required for (and only used by) the "merkle" verifier type.
	RegistryAddress string `json:"registry-address"`
	// TeleporterAddress is the TeleporterMessengerV2 contract on the destination chain. It is a
	// different contract from the adapter and is the target of receiveCrossChainMessage,
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

	verifierType := VerifierTypeMerkle
	if v, ok := m["verifier-type"]; ok {
		verifierType, ok = v.(string)
		if !ok {
			return nil, fmt.Errorf("invalid verifier-type: %v", v)
		}
		if verifierType != VerifierTypeMerkle && verifierType != VerifierTypeWarp {
			return nil, fmt.Errorf("invalid verifier-type: %s", verifierType)
		}
	}

	var registryAddress string
	if verifierType == VerifierTypeMerkle {
		registryAddress, ok = m["registry-address"].(string)
		if !ok {
			return nil, fmt.Errorf("registry-address not found")
		}
		if !common.IsHexAddress(registryAddress) {
			return nil, fmt.Errorf("invalid registry address: %s", registryAddress)
		}
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
		VerifierType:      verifierType,
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
