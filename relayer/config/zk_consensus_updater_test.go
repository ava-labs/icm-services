// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"testing"

	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
)

func TestZKConsensusUpdaterValidate(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	validKey := common.Bytes2Hex(crypto.FromECDSA(key))

	valid := func() ZKConsensusUpdater {
		return ZKConsensusUpdater{
			RPCEndpoint:    "https://rpc.example.com",
			PrivateKey:     validKey,
			AdapterAddress: "0x0000000000000000000000000000000000000001",
			SubgraphURL:    "https://subgraph.example.com",
		}
	}

	testCases := []struct {
		name   string
		modify func(*ZKConsensusUpdater)
		// errContains is empty for a config that must validate.
		errContains string
		// invalidKey marks cases that must fail with ErrInvalidPrivateKeyHex.
		invalidKey bool
	}{
		{name: "valid", modify: func(*ZKConsensusUpdater) {}},
		{
			name:        "missing rpc-endpoint",
			modify:      func(z *ZKConsensusUpdater) { z.RPCEndpoint = "" },
			errContains: "rpc-endpoint required",
		},
		{
			name:        "missing private-key",
			modify:      func(z *ZKConsensusUpdater) { z.PrivateKey = "" },
			errContains: "private-key required",
		},
		{
			name:       "private-key not hex",
			modify:     func(z *ZKConsensusUpdater) { z.PrivateKey = "not-a-hex-private-key" },
			invalidKey: true,
		},
		{
			name:       "private-key wrong length",
			modify:     func(z *ZKConsensusUpdater) { z.PrivateKey = "abcd" },
			invalidKey: true,
		},
		{
			// A 0x prefix is stripped, as for destination account keys and the sample config.
			name:   "private-key with 0x prefix",
			modify: func(z *ZKConsensusUpdater) { z.PrivateKey = "0x" + validKey },
		},
		{
			name:        "invalid adapter-address",
			modify:      func(z *ZKConsensusUpdater) { z.AdapterAddress = "0x1234" },
			errContains: "invalid adapter-address",
		},
		{
			name:        "missing subgraph-url",
			modify:      func(z *ZKConsensusUpdater) { z.SubgraphURL = "" },
			errContains: "subgraph-url required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := valid()
			tc.modify(&cfg)
			err := cfg.Validate(nil)

			switch {
			case tc.invalidKey:
				require.ErrorIs(t, err, utils.ErrInvalidPrivateKeyHex)
				// The key is sensitive and must not leak into the error.
				require.NotContains(t, err.Error(), cfg.PrivateKey)
			case tc.errContains != "":
				require.ErrorContains(t, err, tc.errContains)
			default:
				require.NoError(t, err)
			}
		})
	}
}

// The updater's account must not be shared with any other signing account the relayer is
// configured with, or their transactions' nonces collide on-chain.
func TestZKConsensusUpdaterRejectsSharedSender(t *testing.T) {
	newKeyHex := func() string {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		return common.Bytes2Hex(crypto.FromECDSA(key))
	}
	zkKey, destinationKey, updaterKey, deliveryKey := newKeyHex(), newKeyHex(), newKeyHex(), newKeyHex()

	newConfig := func(destinationKey, deliveryKey string) *Config {
		return &Config{
			DestinationBlockchains: []*DestinationBlockchain{
				{BlockchainID: "destination-1", AccountPrivateKeys: []string{destinationKey}},
			},
			ExternalEVMDestinations: []*ExternalEVMDestination{
				{ContractAddress: "0x02", PrivateKey: updaterKey, DeliveryPrivateKey: deliveryKey},
			},
		}
	}

	testCases := []struct {
		name string
		cfg  *Config
		// errContains is empty when the accounts are distinct.
		errContains string
	}{
		{
			name: "distinct accounts",
			cfg:  newConfig(destinationKey, deliveryKey),
		},
		{
			name:        "shared with a destination account key",
			cfg:         newConfig(zkKey, deliveryKey),
			errContains: "destination blockchain destination-1 account key",
		},
		{
			// Keys are compared by derived address, so a 0x prefix does not hide a match.
			name:        "shared with an external delivery-private-key",
			cfg:         newConfig(destinationKey, "0x"+zkKey),
			errContains: "external EVM destination 0x02 delivery-private-key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			updater := ZKConsensusUpdater{
				RPCEndpoint:    "https://rpc.example.com",
				PrivateKey:     zkKey,
				AdapterAddress: "0x0000000000000000000000000000000000000001",
				SubgraphURL:    "https://subgraph.example.com",
			}
			err := updater.Validate(tc.cfg.senderAccountAddresses())
			if tc.errContains == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.errContains)
			require.ErrorContains(t, err, "nonces would collide")
			// The key is sensitive and must not leak into the error.
			require.NotContains(t, err.Error(), zkKey)
		})
	}
}
