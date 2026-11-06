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
			// Startup parses the key without stripping a 0x prefix, so validation rejects it too.
			name:       "private-key with 0x prefix",
			modify:     func(z *ZKConsensusUpdater) { z.PrivateKey = "0x" + validKey },
			invalidKey: true,
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
			err := cfg.Validate()

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
