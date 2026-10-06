// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"encoding/binary"
	"fmt"
	"math/big"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/common"
)

// The packed wire layout below matches TeleporterMessageV2Parsing.serializeTeleporterMessageV2
// in icm-contracts/common/TeleporterMessageV2.sol. It is abi.encodePacked, NOT abi.encode, so
// the generated ABI bindings (which use abi.encode) cannot parse it. We must (de)serialize it
// manually and keep this in sync with the Solidity implementation.
const (
	nonceLen        = 32
	addressLen      = 20
	blockchainIDLen = 32
	gasLimitLen     = 32
	countLen        = 4  // uint32 length prefixes
	receiptLen      = 52 // 32-byte nonce + 20-byte address

	// Fixed-size header up to (and excluding) the relayer-address count.
	headerLen = nonceLen + addressLen + addressLen + blockchainIDLen + addressLen + gasLimitLen
)

// ParseTeleporterMessageV2 deserializes the packed bytes produced by the Solidity
// serializeTeleporterMessageV2 into a TeleporterMessageV2 struct.
func ParseTeleporterMessageV2(data []byte) (*teleportermessengerv2.TeleporterMessageV2, error) {
	// Minimum length is the fixed header plus the two count prefixes (relayers, receipts).
	// readCount bounds-checks each prefix again; this just fails fast with a clearer error.
	if len(data) < headerLen+countLen+countLen {
		return nil, fmt.Errorf("teleporter message v2 too short: %d bytes", len(data))
	}

	var msg teleportermessengerv2.TeleporterMessageV2
	offset := 0

	msg.MessageNonce = new(big.Int).SetBytes(data[offset : offset+nonceLen])
	offset += nonceLen
	msg.OriginSenderAddress = common.BytesToAddress(data[offset : offset+addressLen])
	offset += addressLen
	msg.OriginTeleporterAddress = common.BytesToAddress(data[offset : offset+addressLen])
	offset += addressLen
	copy(msg.DestinationBlockchainID[:], data[offset:offset+blockchainIDLen])
	offset += blockchainIDLen
	msg.DestinationAddress = common.BytesToAddress(data[offset : offset+addressLen])
	offset += addressLen
	msg.RequiredGasLimit = new(big.Int).SetBytes(data[offset : offset+gasLimitLen])
	offset += gasLimitLen

	// Allowed relayer addresses. The receipts count prefix must also still fit after them.
	numRelayers, err := readCount(data, offset, addressLen, countLen, "relayer addresses")
	if err != nil {
		return nil, err
	}
	offset += countLen
	msg.AllowedRelayerAddresses = make([]common.Address, numRelayers)
	for i := 0; i < numRelayers; i++ {
		msg.AllowedRelayerAddresses[i] = common.BytesToAddress(data[offset : offset+addressLen])
		offset += addressLen
	}

	// Receipts.
	numReceipts, err := readCount(data, offset, receiptLen, 0, "receipts")
	if err != nil {
		return nil, err
	}
	offset += countLen
	msg.Receipts = make([]teleportermessengerv2.TeleporterMessageReceipt, numReceipts)
	for i := 0; i < numReceipts; i++ {
		var receipt teleportermessengerv2.TeleporterMessageReceipt
		receipt.ReceivedMessageNonce = new(big.Int).SetBytes(data[offset : offset+nonceLen])
		offset += nonceLen
		receipt.RelayerRewardAddress = common.BytesToAddress(data[offset : offset+addressLen])
		offset += addressLen
		msg.Receipts[i] = receipt
	}

	// The remainder is the inner application message. Use a non-nil slice so that
	// serialization round-trips are deterministic for empty messages.
	msg.Message = make([]byte, len(data)-offset)
	copy(msg.Message, data[offset:])

	return &msg, nil
}

// readCount reads the uint32 count prefix at data[offset:] and checks that [count] elements of
// [elementLen] bytes, plus [trailingLen] more bytes, fit in the rest of [data]. It checks that
// the prefix itself is in bounds first, so it does not rely on the caller for any bounds check.
//
// This is defensive parsing of untrusted input: messages can arrive via the unauthenticated
// relay API, so a crafted count could overflow the size arithmetic or trigger a huge
// allocation. The check is done in uint64, which cannot wrap here, and before any slice is
// allocated, so a count that passes is always a legitimate, bounded slice length.
func readCount(data []byte, offset, elementLen, trailingLen int, what string) (int, error) {
	// The prefix must be in bounds before it can be read. This also guarantees that
	// [remaining] below cannot go negative.
	if offset < 0 || offset > len(data)-countLen {
		return 0, fmt.Errorf(
			"teleporter message v2 truncated reading %s: count prefix at offset %d, %d bytes total",
			what,
			offset,
			len(data),
		)
	}
	count := binary.BigEndian.Uint32(data[offset : offset+countLen])
	remaining := uint64(len(data) - offset - countLen)
	if uint64(count)*uint64(elementLen)+uint64(trailingLen) > remaining {
		return 0, fmt.Errorf(
			"teleporter message v2 truncated reading %s: %d declared, %d bytes remaining",
			what,
			count,
			remaining,
		)
	}
	return int(count), nil
}

// SerializeTeleporterMessageV2 serializes a TeleporterMessageV2 into the packed format expected by
// the Solidity contracts. Mirrors serializeTeleporterMessageV2 and is primarily used for testing
// round-trips against ParseTeleporterMessageV2.
func SerializeTeleporterMessageV2(msg *teleportermessengerv2.TeleporterMessageV2) []byte {
	size := headerLen + countLen + len(msg.AllowedRelayerAddresses)*addressLen +
		countLen + len(msg.Receipts)*receiptLen + len(msg.Message)
	buf := make([]byte, 0, size)

	buf = appendUint256(buf, msg.MessageNonce)
	buf = append(buf, msg.OriginSenderAddress.Bytes()...)
	buf = append(buf, msg.OriginTeleporterAddress.Bytes()...)
	buf = append(buf, msg.DestinationBlockchainID[:]...)
	buf = append(buf, msg.DestinationAddress.Bytes()...)
	buf = appendUint256(buf, msg.RequiredGasLimit)

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(msg.AllowedRelayerAddresses)))
	for _, addr := range msg.AllowedRelayerAddresses {
		buf = append(buf, addr.Bytes()...)
	}

	buf = binary.BigEndian.AppendUint32(buf, uint32(len(msg.Receipts)))
	for _, receipt := range msg.Receipts {
		buf = appendUint256(buf, receipt.ReceivedMessageNonce)
		buf = append(buf, receipt.RelayerRewardAddress.Bytes()...)
	}

	buf = append(buf, msg.Message...)
	return buf
}

// appendUint256 appends a big.Int as a left-zero-padded 32-byte big-endian value.
func appendUint256(buf []byte, value *big.Int) []byte {
	var padded [nonceLen]byte
	if value != nil {
		value.FillBytes(padded[:])
	}
	return append(buf, padded[:]...)
}
