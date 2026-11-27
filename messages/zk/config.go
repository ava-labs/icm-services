// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package zk

import (
	"fmt"
	"math"

	"github.com/ava-labs/libevm/common"
)

// Config holds the settings for the TeleporterV2 zk message handler.
//
// Messages originating on Ethereum are verified on the destination chain by a ZKAdapter
// that is called by TeleporterMessengerV2.ReceiveCrossChainMessage. The relayer updates
// the state of the destination-side ZKAdapter in 2 possible (independent) ways:
// 1. It provides a Boundless zero-knowledge proof of a beacon state (consensus) transition that
// ZKAdapter verifies and stores as a confirmed beacon anchor.
// 2. It provides a Merkle proof tuple (SSZ execution state proof, MPT receipt proof)
// against an already verified beacon state anchor stored in ZKAdapter.
//
// In summary, the relayer will periodically update ZKAdapter with consensus state transitions by
// submitting proofs as they are made available on the Boundless network. Additionally, when it sees
// a TeleporterV2 message sent on the source chain, it will submit a Merkle proof tuple to ZKAdapter.
type Config struct {
	// RewardAddress is the address credited as the relayer reward recipient on the source chain.
	RewardAddress string `json:"reward-address"`
	// Destination AdapterAddress is the ZKAdapter deployed on the destination chain. It is queried for
	// confirmed beacon anchors (getBeaconBlockRoot) so the relayer builds proofs against a slot
	// the contract has verified via Boundless consensus transitions. This is the same contract
	// as the source-side adapter emitting TeleporterV2MessageSent deployed via Nick's method.
	AdapterAddress string `json:"adapter-address"`
	// TeleporterAddress is the TeleporterMessengerV2 contract on the destination chain. It is a
	// different contract from the adapter and is the target of receiveCrossChainMessage,
	// the contract used to compute the Teleporter message ID, and the one queried for delivery
	// status. With the universal deployer it is identical on the source chain.
	TeleporterAddress string `json:"teleporter-address"`
	// BeaconGenesisTime is the source beacon chain's genesis time in unix seconds. It is
	// network specific (e.g., mainnet: 1606824023, testnet differs) and is used with SecondsPerSlot to
	// map source block timestamps to beacon slots for anchor selection and proof construction.
	BeaconGenesisTime uint64 `json:"beacon-genesis-time"`
	// SecondsPerSlot is the source beacon chain's slot duration in seconds.
	SecondsPerSlot uint64 `json:"seconds-per-slot"`
	// BeaconAPIURL is a beacon node endpoint used to fetch beacon blocks and states when
	// building the SSZ execution proof.
	BeaconAPIURL string `json:"beacon-api-url"`
}

// ConfigFromMap parses and validates the zk protocol settings.
func ConfigFromMap(m map[string]any) (*Config, error) {
	rewardAddress, ok := m["reward-address"].(string)
	if !ok {
		return nil, fmt.Errorf("reward-address not found")
	}
	if !common.IsHexAddress(rewardAddress) {
		return nil, fmt.Errorf("invalid reward-address: %s", rewardAddress)
	}
	adapterAddress, ok := m["adapter-address"].(string)
	if !ok {
		return nil, fmt.Errorf("adapter-address not found")
	}
	if !common.IsHexAddress(adapterAddress) {
		return nil, fmt.Errorf("invalid adapter-address: %s", adapterAddress)
	}
	teleporterAddress, ok := m["teleporter-address"].(string)
	if !ok {
		return nil, fmt.Errorf("teleporter-address not found")
	}
	if !common.IsHexAddress(teleporterAddress) {
		return nil, fmt.Errorf("invalid teleporter-address: %s", teleporterAddress)
	}
	beaconGenesisTime, err := uintSetting(m, "beacon-genesis-time")
	if err != nil {
		return nil, err
	}
	secondsPerSlot, err := uintSetting(m, "seconds-per-slot")
	if err != nil {
		return nil, err
	}
	// The slot of a source block is computed by dividing by the slot duration.
	if secondsPerSlot == 0 {
		return nil, fmt.Errorf("seconds-per-slot must be greater than 0")
	}
	beaconAPIURL, ok := m["beacon-api-url"].(string)
	if !ok || beaconAPIURL == "" {
		return nil, fmt.Errorf("beacon-api-url not found")
	}

	return &Config{
		RewardAddress:     rewardAddress,
		AdapterAddress:    adapterAddress,
		TeleporterAddress: teleporterAddress,
		BeaconGenesisTime: beaconGenesisTime,
		SecondsPerSlot:    secondsPerSlot,
		BeaconAPIURL:      beaconAPIURL,
	}, nil
}

// uintSetting reads the whole, non-negative number stored under [key]. Numbers in the settings
// map are decoded from JSON as float64, and converting one to uint64 unchecked would silently
// truncate a fraction (0.5 becomes 0) and has no defined result for a negative value.
func uintSetting(m map[string]any, key string) (uint64, error) {
	value, ok := m[key].(float64)
	if !ok {
		return 0, fmt.Errorf("%s not found", key)
	}
	// The comparison with Trunc also rejects NaN. MaxUint64 rounds up to 2^64 as a float64, the
	// first value that no longer fits.
	if value < 0 || value != math.Trunc(value) || value >= math.MaxUint64 {
		return 0, fmt.Errorf("%s must be a non-negative integer, got %v", key, value)
	}
	return uint64(value), nil
}

func (c *Config) adapterAddress() common.Address {
	return common.HexToAddress(c.AdapterAddress)
}

func (c *Config) teleporterAddress() common.Address {
	return common.HexToAddress(c.TeleporterAddress)
}

func (c *Config) rewardAddress() common.Address {
	return common.HexToAddress(c.RewardAddress)
}
