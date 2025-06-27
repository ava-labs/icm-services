// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"fmt"

	"github.com/ava-labs/libevm/common"
)

type Config struct {
	RewardAddress string `json:"reward-address"`
}

func ConfigFromMap(m map[string]any) (*Config, error) {
	rewardAddress, ok := m["reward-address"].(string)
	if !ok {
		return nil, fmt.Errorf("reward-address not found")
	}

	if !common.IsHexAddress(rewardAddress) {
		return nil, fmt.Errorf("invalid reward address for EVM source subnet: %s", rewardAddress)
	}

	return &Config{
		RewardAddress: rewardAddress,
	}, nil
}
