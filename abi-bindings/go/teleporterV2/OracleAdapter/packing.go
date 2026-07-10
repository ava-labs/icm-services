// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package oracleadapter

import (
	"fmt"
	"math/big"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
)

// OracleMessage mirrors the Solidity OracleMessage struct in OracleAdapter.sol.
type OracleMessage struct {
	SourceType        string
	SourceAddress     string
	DestContract      common.Address
	SourceBlockHeight *big.Int
	Nonce             *big.Int
	Payload           []byte
}

var (
	oracleAttestationArgs  abi.Arguments
	oracleMessageBytesArgs abi.Arguments
)

func init() {
	stringT, err := abi.NewType("string", "", nil)
	if err != nil {
		panic(fmt.Sprintf("oracle packing: failed to create string type: %v", err))
	}
	uint32T, err := abi.NewType("uint32", "", nil)
	if err != nil {
		panic(fmt.Sprintf("oracle packing: failed to create uint32 type: %v", err))
	}
	uint256T, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(fmt.Sprintf("oracle packing: failed to create uint256 type: %v", err))
	}
	bytesT, err := abi.NewType("bytes", "", nil)
	if err != nil {
		panic(fmt.Sprintf("oracle packing: failed to create bytes type: %v", err))
	}
	// abi.encode(uint32 warpIndex) — stored in attestation
	// The oracle message is decoded directly from the BLS-verified warp payload.
	oracleAttestationArgs = abi.Arguments{
		{Name: "warpIndex", Type: uint32T},
	}

	// abi.encode(sourceType, sourceAddress, sourceBlockHeight, nonce, payload)
	// stored in TeleporterMessageV2.message and decoded by destination contracts
	oracleMessageBytesArgs = abi.Arguments{
		{Name: "sourceType", Type: stringT},
		{Name: "sourceAddress", Type: stringT},
		{Name: "sourceBlockHeight", Type: uint256T},
		{Name: "nonce", Type: uint256T},
		{Name: "payload", Type: bytesT},
	}
}

// PackOracleAttestation encodes the warp index into the attestation bytes expected
// by OracleAdapter.verifyMessage. The oracle message is decoded by the contract
// directly from the BLS-verified warp payload.
func PackOracleAttestation(warpIndex uint32) ([]byte, error) {
	return oracleAttestationArgs.Pack(warpIndex)
}

// PackOracleMessageBytes encodes oracle fields into the message bytes stored in
// TeleporterMessageV2.message. Destination contracts decode these fields via
// abi.decode(message, (string, string, uint256, uint256, bytes)).
func PackOracleMessageBytes(oracleMsg OracleMessage) ([]byte, error) {
	return oracleMessageBytesArgs.Pack(
		oracleMsg.SourceType,
		oracleMsg.SourceAddress,
		oracleMsg.SourceBlockHeight,
		oracleMsg.Nonce,
		oracleMsg.Payload,
	)
}

// BuildOracleICMMessage constructs a TeleporterICMMessage for an oracle delivery.
//
// The relayer supplies:
//   - warpIndex: index of the verified warp message in predicate storage
//   - oracleMsg: the oracle message fields
//   - teleporterAddress: address of TeleporterMessengerV2 on this chain (used as originTeleporterAddress)
//   - thisChainID: this L1's blockchain ID (used as both sourceBlockchainID and destinationBlockchainID)
//   - networkID: ICM network ID of this chain
//   - requiredGasLimit: gas limit for the destination contract call
func BuildOracleICMMessage(
	warpIndex uint32,
	oracleMsg OracleMessage,
	teleporterAddress common.Address,
	thisChainID [32]byte,
	networkID uint32,
	requiredGasLimit *big.Int,
) (teleportermessengerv2.TeleporterICMMessage, error) {
	attestation, err := PackOracleAttestation(warpIndex)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, fmt.Errorf("pack oracle attestation: %w", err)
	}

	msgBytes, err := PackOracleMessageBytes(oracleMsg)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, fmt.Errorf("pack oracle message bytes: %w", err)
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
