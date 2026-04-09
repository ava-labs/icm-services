package database

import (
	"fmt"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var testingProtocolAddress = common.HexToAddress("0xd81545385803bCD83bd59f58Ba2d2c0562387F83")

func TestIsKeyNotFoundError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "key not found error",
			err:      ErrKeyNotFound,
			expected: true,
		},
		{
			name:     "relayer key not found error",
			err:      ErrRelayerIDNotFound,
			expected: true,
		},
		{
			name:     "unknown error",
			err:      errors.New("unknown error"),
			expected: false,
		},
	}
	for _, testCase := range testCases {
		result := IsKeyNotFoundError(testCase.err)
		require.Equal(t, testCase.expected, result, testCase.name)
	}
}

func TestCalculateRelayerID(t *testing.T) {
	id1, _ := ids.FromString("S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD")
	id2, _ := ids.FromString("2TGBXcnwx5PqiXWiqxAKUaNSqDguXNh1mxnp82jui68hxJSZAx")
	testCases := []struct {
		name                    string
		sourceBlockchainID      ids.ID
		destinationBlockchainID ids.ID
		originSenderAddress     common.Address
		destinationAddress      common.Address
		expected                common.Hash
	}{
		{
			name:                    "all zero",
			sourceBlockchainID:      id1,
			destinationBlockchainID: id2,
			originSenderAddress:     AllAllowedAddress,
			destinationAddress:      AllAllowedAddress,
			expected:                common.HexToHash("0xbeea4a95895a40befa2295a8cd56586cce9ea6d28d8616569c542cc1a632edb3"),
		},
		{
			name:                    "zero source address",
			sourceBlockchainID:      id1,
			destinationBlockchainID: id2,
			originSenderAddress:     AllAllowedAddress,
			destinationAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			expected:                common.HexToHash("0x256b953633895cebec219cca02360c095afc40879958ff76cf7b71ece6f64f10"),
		},
		{
			name:                    "zero destination address",
			sourceBlockchainID:      id1,
			destinationBlockchainID: id2,
			originSenderAddress:     common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			destinationAddress:      AllAllowedAddress,
			expected:                common.HexToHash("0x90acac45798533ca7263420a73f924b22860184920cd6650aab99a19695f9abb"),
		},
		{
			name:                    "all non-zero",
			sourceBlockchainID:      id1,
			destinationBlockchainID: id2,
			originSenderAddress:     common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			destinationAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			expected:                common.HexToHash("0xefe8c09dd4232671543a04507abba2a6a3d60ed0b65a2296dc3707202738a7b4"),
		},
	}
	for _, testCase := range testCases {
		result := CalculateRelayerID(
			testingProtocolAddress,
			testCase.sourceBlockchainID,
			testCase.destinationBlockchainID,
			testCase.originSenderAddress,
			testCase.destinationAddress,
		)
		require.Equal(t, testCase.expected, result, testCase.name)
	}
}

func TestGetConfigRelayerKeys(t *testing.T) {
	allowedAddress := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
	dstCfg1 := config.TestValidDestinationBlockchainConfig

	// All destination chains and source and destination addresses are allowed
	srcCfg1 := config.TestValidSourceBlockchainConfig

	// All destination chains, but only a single source address is allowed
	srcCfg2 := config.TestValidSourceBlockchainConfig
	srcCfg2.BlockchainID = ids.GenerateTestID().String()
	srcCfg2.AllowedOriginSenderAddresses = []string{allowedAddress.String()}

	// Restricted to a single destination chain, but all source and destination addresses are allowed
	srcCfg3 := config.TestValidSourceBlockchainConfig
	srcCfg3.BlockchainID = ids.GenerateTestID().String()
	srcCfg3.SupportedDestinations = []*config.SupportedDestination{
		{
			BlockchainID: dstCfg1.BlockchainID,
		},
	}

	// Restricted to a single destination chain, but only a single source address is allowed
	srcCfg4 := config.TestValidSourceBlockchainConfig
	srcCfg4.BlockchainID = ids.GenerateTestID().String()
	srcCfg4.AllowedOriginSenderAddresses = []string{allowedAddress.String()}
	srcCfg4.SupportedDestinations = []*config.SupportedDestination{
		{
			BlockchainID: dstCfg1.BlockchainID,
		},
	}

	// Restricted to a single destination chain, but only a single destination address is allowed
	srcCfg5 := config.TestValidSourceBlockchainConfig
	srcCfg5.BlockchainID = ids.GenerateTestID().String()
	srcCfg5.SupportedDestinations = []*config.SupportedDestination{
		{
			BlockchainID: dstCfg1.BlockchainID,
			Addresses:    []string{allowedAddress.String()},
		},
	}

	// Restricted to a single destination, but only a single source and destination address is allowed
	srcCfg6 := config.TestValidSourceBlockchainConfig
	srcCfg6.BlockchainID = ids.GenerateTestID().String()
	srcCfg6.AllowedOriginSenderAddresses = []string{allowedAddress.String()}
	srcCfg6.SupportedDestinations = []*config.SupportedDestination{
		{
			BlockchainID: dstCfg1.BlockchainID,
			Addresses:    []string{allowedAddress.String()},
		},
	}

	//

	err := dstCfg1.Validate()
	require.ErrorIs(t, err, nil)

	allowedDestinations := set.NewSet[string](1)
	allowedDestinations.Add(dstCfg1.BlockchainID)
	err = srcCfg1.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)
	err = srcCfg2.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)
	err = srcCfg3.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)
	err = srcCfg4.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)
	err = srcCfg5.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)
	err = srcCfg6.Validate(&allowedDestinations)
	require.ErrorIs(t, err, nil)

	cfg := &config.Config{
		SourceBlockchains:      []*config.SourceBlockchain{&srcCfg1, &srcCfg2, &srcCfg3, &srcCfg4, &srcCfg5, &srcCfg6},
		DestinationBlockchains: []*config.DestinationBlockchain{&dstCfg1},
	}

	targetIDs := []RelayerID{
		NewRelayerID(
			testingProtocolAddress,
			srcCfg1.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			AllAllowedAddress,
			AllAllowedAddress,
		),
		NewRelayerID(
			testingProtocolAddress,
			srcCfg2.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			allowedAddress,
			AllAllowedAddress,
		),
		NewRelayerID(
			testingProtocolAddress,
			srcCfg3.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			AllAllowedAddress,
			AllAllowedAddress,
		),
		NewRelayerID(
			testingProtocolAddress,
			srcCfg4.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			allowedAddress,
			AllAllowedAddress,
		),
		NewRelayerID(
			testingProtocolAddress,
			srcCfg5.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			AllAllowedAddress,
			allowedAddress,
		),
		NewRelayerID(
			testingProtocolAddress,
			srcCfg6.GetBlockchainID(),
			dstCfg1.GetBlockchainID(),
			allowedAddress,
			allowedAddress,
		),
	}

	relayerIDs := GetConfigRelayerIDs(cfg)

	// Test that all target IDs are present
	for i, id := range targetIDs {
		require.True(t,
			func(ids []RelayerID, target RelayerID) bool {
				for _, id := range ids {
					if id.ID == target.ID {
						return true
					}
				}
				return false
			}(relayerIDs, id),
			fmt.Sprintf("target ID %d not found", i),
		)
	}
}
