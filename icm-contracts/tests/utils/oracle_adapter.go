// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"math/big"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	relayerUtils "github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

// Thin gomega wrappers around the oracle packing helpers in
// github.com/ava-labs/icm-services/utils, which pack through the generated ABI
// of the pure encode/decode helpers on OracleAdapter.sol. Production consumers
// (e.g. solana-observer) use the error-returning versions directly.

// PackOracleWarpPayload encodes an oracle message into the warp payload format
// decoded by OracleAdapter.decodeOracleMessage.
func PackOracleWarpPayload(msg oracleadapter.OracleMessage) []byte {
	packed, err := relayerUtils.PackOracleWarpPayload(msg)
	Expect(err).Should(BeNil())
	return packed
}

// PackOracleAttestation encodes the warp index into the attestation bytes
// expected by OracleAdapter.verifyMessage.
func PackOracleAttestation(warpIndex uint32) []byte {
	packed, err := relayerUtils.PackOracleAttestation(warpIndex)
	Expect(err).Should(BeNil())
	return packed
}

// PackOracleReceiverPayload encodes oracle fields into the message bytes stored
// in TeleporterMessageV2.message.
func PackOracleReceiverPayload(msg oracleadapter.OracleMessage) []byte {
	packed, err := relayerUtils.PackOracleReceiverPayload(msg)
	Expect(err).Should(BeNil())
	return packed
}

// BuildOracleICMMessage constructs a TeleporterICMMessage for an oracle delivery.
func BuildOracleICMMessage(
	warpIndex uint32,
	oracleMsg oracleadapter.OracleMessage,
	teleporterAddress common.Address,
	thisChainID [32]byte,
	networkID uint32,
	requiredGasLimit *big.Int,
) teleportermessengerv2.TeleporterICMMessage {
	icmMsg, err := relayerUtils.BuildOracleICMMessage(
		warpIndex, oracleMsg, teleporterAddress, thisChainID, networkID, requiredGasLimit,
	)
	Expect(err).Should(BeNil())
	return icmMsg
}
