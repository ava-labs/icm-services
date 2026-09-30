// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"encoding/binary"
	"math"
	"math/big"
	"testing"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

func TestTeleporterMessageV2SerializationRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		msg  teleportermessengerv2.TeleporterMessageV2
	}{
		{
			name: "minimal",
			msg: teleportermessengerv2.TeleporterMessageV2{
				MessageNonce:            big.NewInt(1),
				OriginSenderAddress:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
				OriginTeleporterAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
				DestinationBlockchainID: [32]byte{0xaa, 0xbb},
				DestinationAddress:      common.HexToAddress("0x3333333333333333333333333333333333333333"),
				RequiredGasLimit:        big.NewInt(100000),
				AllowedRelayerAddresses: []common.Address{},
				Receipts:                []teleportermessengerv2.TeleporterMessageReceipt{},
				Message:                 []byte("hello"),
			},
		},
		{
			name: "with relayers and receipts",
			msg: teleportermessengerv2.TeleporterMessageV2{
				MessageNonce:            new(big.Int).SetBytes(common.FromHex("0xffffffffffffffffffffffffffffffff")),
				OriginSenderAddress:     common.HexToAddress("0x4444444444444444444444444444444444444444"),
				OriginTeleporterAddress: common.HexToAddress("0x5555555555555555555555555555555555555555"),
				DestinationBlockchainID: [32]byte{0x01, 0x02, 0x03, 0x04, 0x05},
				DestinationAddress:      common.HexToAddress("0x6666666666666666666666666666666666666666"),
				RequiredGasLimit:        big.NewInt(250000),
				AllowedRelayerAddresses: []common.Address{
					common.HexToAddress("0x7777777777777777777777777777777777777777"),
					common.HexToAddress("0x8888888888888888888888888888888888888888"),
				},
				Receipts: []teleportermessengerv2.TeleporterMessageReceipt{
					{
						ReceivedMessageNonce: big.NewInt(42),
						RelayerRewardAddress: common.HexToAddress("0x9999999999999999999999999999999999999999"),
					},
				},
				Message: []byte{0xde, 0xad, 0xbe, 0xef},
			},
		},
		{
			name: "empty inner message",
			msg: teleportermessengerv2.TeleporterMessageV2{
				MessageNonce:            big.NewInt(7),
				OriginSenderAddress:     common.HexToAddress("0xabcdef0123456789abcdef0123456789abcdef01"),
				OriginTeleporterAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				DestinationBlockchainID: [32]byte{0xff},
				DestinationAddress:      common.Address{},
				RequiredGasLimit:        big.NewInt(0),
				AllowedRelayerAddresses: []common.Address{
					common.HexToAddress("0x0000000000000000000000000000000000000001"),
				},
				Receipts: []teleportermessengerv2.TeleporterMessageReceipt{},
				Message:  []byte{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialized := SerializeTeleporterMessageV2(&tt.msg)
			parsed, err := ParseTeleporterMessageV2(serialized)
			require.NoError(t, err)

			require.Equal(t, 0, tt.msg.MessageNonce.Cmp(parsed.MessageNonce))
			require.Equal(t, tt.msg.OriginSenderAddress, parsed.OriginSenderAddress)
			require.Equal(t, tt.msg.OriginTeleporterAddress, parsed.OriginTeleporterAddress)
			require.Equal(t, tt.msg.DestinationBlockchainID, parsed.DestinationBlockchainID)
			require.Equal(t, tt.msg.DestinationAddress, parsed.DestinationAddress)
			require.Equal(t, 0, tt.msg.RequiredGasLimit.Cmp(parsed.RequiredGasLimit))
			require.Equal(t, tt.msg.AllowedRelayerAddresses, parsed.AllowedRelayerAddresses)
			require.Len(t, parsed.Receipts, len(tt.msg.Receipts))
			for i := range tt.msg.Receipts {
				require.Equal(t, 0, tt.msg.Receipts[i].ReceivedMessageNonce.Cmp(parsed.Receipts[i].ReceivedMessageNonce))
				require.Equal(t, tt.msg.Receipts[i].RelayerRewardAddress, parsed.Receipts[i].RelayerRewardAddress)
			}
			require.Equal(t, tt.msg.Message, parsed.Message)

			// Re-serializing the parsed message must yield identical bytes.
			require.Equal(t, serialized, SerializeTeleporterMessageV2(parsed))
		})
	}
}

func TestParseTeleporterMessageV2Errors(t *testing.T) {
	_, err := ParseTeleporterMessageV2([]byte{0x00, 0x01, 0x02})
	require.Error(t, err)
}

// TestParseTeleporterMessageV2RejectsCraftedCounts verifies that attacker-controlled count
// prefixes are rejected with an error before anything is allocated, including values that would
// go negative or wrap when multiplied by the element size in 32-bit int arithmetic. None of
// these inputs may panic or allocate.
func TestParseTeleporterMessageV2RejectsCraftedCounts(t *testing.T) {
	// A valid message with no relayers, no receipts, and no inner message: the relayer count
	// prefix sits at headerLen and the receipts count prefix directly after it.
	base := SerializeTeleporterMessageV2(&teleportermessengerv2.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []teleportermessengerv2.TeleporterMessageReceipt{},
		Message:                 []byte{},
	})
	require.Len(t, base, headerLen+countLen+countLen)
	relayerCountOffset := headerLen
	receiptCountOffset := headerLen + countLen

	const (
		// Smallest counts whose byte size (count*elementLen) exceeds math.MaxInt32, so the
		// product wraps negative in 32-bit int arithmetic while the count itself stays positive.
		relayerCountWrapsInt32 = uint32(math.MaxInt32/addressLen + 1)
		receiptCountWrapsInt32 = uint32(math.MaxInt32/receiptLen + 1)
	)

	tests := []struct {
		name         string
		countOffset  int
		count        uint32
		trailerBytes int // bytes appended after the count prefixes
	}{
		{"one relayer declared, none present", relayerCountOffset, 1, 0},
		{"one relayer declared, one byte short", relayerCountOffset, 1, addressLen - 1},
		{"relayer count negative as int32", relayerCountOffset, 1 << 31, 0},
		{"relayer count max uint32", relayerCountOffset, math.MaxUint32, 0},
		{"relayer byte size wraps int32", relayerCountOffset, relayerCountWrapsInt32, 0},
		{"one receipt declared, none present", receiptCountOffset, 1, 0},
		{"one receipt declared, one byte short", receiptCountOffset, 1, receiptLen - 1},
		{"receipt count negative as int32", receiptCountOffset, 1 << 31, 0},
		{"receipt count max uint32", receiptCountOffset, math.MaxUint32, 0},
		{"receipt byte size wraps int32", receiptCountOffset, receiptCountWrapsInt32, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := append([]byte{}, base...)
			binary.BigEndian.PutUint32(data[tt.countOffset:], tt.count)
			data = append(data, make([]byte, tt.trailerBytes)...)

			require.NotPanics(t, func() {
				parsed, err := ParseTeleporterMessageV2(data)
				require.ErrorContains(t, err, "truncated")
				require.Nil(t, parsed)
			})
		})
	}

	// Sanity check the count offsets used above against a message that has real entries.
	withEntries := SerializeTeleporterMessageV2(&teleportermessengerv2.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{{0x01}, {0x02}},
		Receipts: []teleportermessengerv2.TeleporterMessageReceipt{
			{ReceivedMessageNonce: big.NewInt(1)},
		},
		Message: []byte{},
	})
	require.Equal(t, uint32(2), binary.BigEndian.Uint32(withEntries[relayerCountOffset:]))
	require.Equal(t, uint32(1), binary.BigEndian.Uint32(withEntries[receiptCountOffset+2*addressLen:]))
}
