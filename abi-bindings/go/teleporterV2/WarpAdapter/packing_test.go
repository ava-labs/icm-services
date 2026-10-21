// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warpadapter

import (
	"math"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

func testTeleporterMessage() teleportermessengerv2.TeleporterMessageV2 {
	addr := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
	return teleportermessengerv2.TeleporterMessageV2{
		MessageNonce:            big.NewInt(5),
		OriginSenderAddress:     addr,
		OriginTeleporterAddress: addr,
		DestinationBlockchainID: ids.ID{1, 2, 3, 4},
		DestinationAddress:      addr,
		RequiredGasLimit:        big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{addr},
		Receipts: []teleportermessengerv2.TeleporterMessageReceipt{
			{ReceivedMessageNonce: big.NewInt(1), RelayerRewardAddress: addr},
		},
		Message: []byte{1, 2, 3, 4},
	}
}

// unpackAttestation returns the attestation carried by receiveCrossChainMessage calldata.
func unpackAttestation(t *testing.T, callData []byte) []byte {
	teleporterABI, err := teleportermessengerv2.TeleporterMessengerV2MetaData.GetAbi()
	require.NoError(t, err)
	method, ok := teleporterABI.Methods["receiveCrossChainMessage"]
	require.True(t, ok)
	require.Equal(t, method.ID, callData[:4])

	args, err := method.Inputs.Unpack(callData[4:])
	require.NoError(t, err)
	icmMessage := *abi.ConvertType(
		args[0], new(teleportermessengerv2.TeleporterICMMessage),
	).(*teleportermessengerv2.TeleporterICMMessage)
	return icmMessage.Attestation
}

// The attestation must be the Warp message index as a 32-byte word that WarpAdapter.sol can
// abi.decode as a uint32, for every index that type can hold.
func TestPackReceiveCrossChainMessageAttestation(t *testing.T) {
	uint32Type, err := abi.NewType("uint32", "", nil)
	require.NoError(t, err)
	uint32Args := abi.Arguments{{Type: uint32Type}}

	for _, index := range []uint64{0, 1, 7, 1 << 31, math.MaxUint32} {
		callData, err := PackReceiveCrossChainMessage(
			testTeleporterMessage(), ids.ID{9}, index, common.Address{1},
		)
		require.NoError(t, err, "index %d", index)

		attestation := unpackAttestation(t, callData)
		expected := make([]byte, 32)
		new(big.Int).SetUint64(index).FillBytes(expected)
		require.Equal(t, expected, attestation, "index %d", index)

		decoded, err := uint32Args.Unpack(attestation)
		require.NoError(t, err)
		require.Equal(t, uint32(index), decoded[0].(uint32), "index %d", index)
	}
}

// An index WarpAdapter.sol cannot decode is rejected instead of being encoded incorrectly:
// math.MaxUint64 in particular used to be cast through int64 and encoded as 1.
func TestPackReceiveCrossChainMessageRejectsIndexBeyondUint32(t *testing.T) {
	for _, index := range []uint64{math.MaxUint32 + 1, 1 << 63, math.MaxUint64} {
		callData, err := PackReceiveCrossChainMessage(
			testTeleporterMessage(), ids.ID{9}, index, common.Address{1},
		)
		require.ErrorContains(t, err, "exceeds the uint32", "index %d", index)
		require.Nil(t, callData, "index %d", index)
	}
}
