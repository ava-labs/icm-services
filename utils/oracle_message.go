// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"fmt"
	"math/big"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
)

// The oracle wire formats are defined by the pure encode/decode helpers on
// OracleAdapter.sol (encodeOracleMessage, encodeOracleAttestation,
// encodeReceiverPayload, decodeOracleMessage). All packing below goes through
// the generated ABI of those helpers, keeping the Solidity contract the single
// source of truth for the encodings.

// PackOracleWarpPayload encodes an oracle message into the warp payload format
// decoded by OracleAdapter.decodeOracleMessage.
func PackOracleWarpPayload(msg oracleadapter.OracleMessage) ([]byte, error) {
	inputs, err := oracleAdapterMethodInputs("encodeOracleMessage")
	if err != nil {
		return nil, err
	}
	return inputs.Pack(
		msg.SourceType,
		msg.SourceAddress,
		msg.DestContract,
		msg.SourceBlockHeight,
		msg.Nonce,
		msg.Payload,
	)
}

// PackOracleAttestation encodes the warp index into the attestation bytes
// expected by OracleAdapter.verifyMessage.
func PackOracleAttestation(warpIndex uint32) ([]byte, error) {
	inputs, err := oracleAdapterMethodInputs("encodeOracleAttestation")
	if err != nil {
		return nil, err
	}
	return inputs.Pack(warpIndex)
}

// PackOracleReceiverPayload encodes oracle fields into the message bytes stored
// in TeleporterMessageV2.message. Destination contracts decode these fields via
// abi.decode(message, (string, string, uint256, uint256, bytes)).
func PackOracleReceiverPayload(msg oracleadapter.OracleMessage) ([]byte, error) {
	inputs, err := oracleAdapterMethodInputs("encodeReceiverPayload")
	if err != nil {
		return nil, err
	}
	return inputs.Pack(
		msg.SourceType,
		msg.SourceAddress,
		msg.SourceBlockHeight,
		msg.Nonce,
		msg.Payload,
	)
}

// BuildOracleICMMessage constructs a TeleporterICMMessage for an oracle delivery.
//
//   - warpIndex: index of the verified warp message in predicate storage
//   - oracleMsg: the oracle message fields
//   - teleporterAddress: address of TeleporterMessengerV2 on this chain (used as originTeleporterAddress)
//   - thisChainID: this L1's blockchain ID (used as both sourceBlockchainID and destinationBlockchainID)
//   - networkID: ICM network ID of this chain
//   - requiredGasLimit: gas limit for the destination contract call
func BuildOracleICMMessage(
	warpIndex uint32,
	oracleMsg oracleadapter.OracleMessage,
	teleporterAddress common.Address,
	thisChainID [32]byte,
	networkID uint32,
	requiredGasLimit *big.Int,
) (teleportermessengerv2.TeleporterICMMessage, error) {
	attestation, err := PackOracleAttestation(warpIndex)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, fmt.Errorf("pack oracle attestation: %w", err)
	}

	msgBytes, err := PackOracleReceiverPayload(oracleMsg)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, fmt.Errorf("pack oracle receiver payload: %w", err)
	}

	return teleportermessengerv2.TeleporterICMMessage{
		Message: teleportermessengerv2.TeleporterMessageV2{
			MessageNonce:            oracleMsg.Nonce,
			OriginSenderAddress:     common.Address{}, // address(0) — no EVM sender for oracle messages
			OriginTeleporterAddress: teleporterAddress,
			DestinationBlockchainID: thisChainID,
			DestinationAddress:      oracleMsg.DestContract,
			RequiredGasLimit:        requiredGasLimit,
			AllowedRelayerAddresses: []common.Address{},
			Receipts:                []teleportermessengerv2.TeleporterMessageReceipt{},
			Message:                 msgBytes,
		},
		SourceNetworkID:    networkID,
		SourceBlockchainID: thisChainID,
		Attestation:        attestation,
	}, nil
}

// oracleAdapterMethodInputs returns the input argument list of the named
// OracleAdapter encoding helper from the generated ABI.
func oracleAdapterMethodInputs(method string) (abi.Arguments, error) {
	oracleABI, err := oracleadapter.OracleAdapterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("get OracleAdapter ABI: %w", err)
	}
	m, ok := oracleABI.Methods[method]
	if !ok {
		return nil, fmt.Errorf("OracleAdapter ABI is missing method %s", method)
	}
	return m.Inputs, nil
}
