// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package zk

import (
	"fmt"
	"math/big"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
)

// zkAdapterABI is the parsed ZKAdapter ABI sourced from the generated bindings
var zkAdapterABI = getZkAdapterABI()

func getZkAdapterABI() *abi.ABI {
	parsed, err := zkadapter.ZKAdapterMetaData.GetAbi()
	if err != nil {
		panic(err) // bindings are corrupted
	}
	return parsed
}

// messageSentTopic is the topic hash (topic0) of the adapter's TeleporterV2MessageSent event
var messageSentTopic = zkAdapterABI.Events["TeleporterV2MessageSent"].ID

// messageArguments is the ABI of the TeleporterMessageV2 struct, taken from the adapter's
// sendMessage input. The adapter emits abi.encode(message), which is this argument list's encoding.
var messageArguments = zkAdapterABI.Methods["sendMessage"].Inputs

// attestationArguments is the ABI of the two proofs that make up the adapter's Attestation
// struct, taken from verifyLogAndExtract(execProof, logProof). The struct itself is not in the
// ABI since no external function takes it directly.
var attestationArguments = zkAdapterABI.Methods["verifyLogAndExtract"].Inputs

// EventFilter returns the source chain log filter matching messages sent by the zk message
// protocol. Duplicate deliveries are rejected at the TeleporterV2.receiveCrossChainMessage
// level, allowing the relayer to remain stateless.
func (f *factory) EventFilter() evm.EventFilter {
	return evm.EventFilter{
		Addresses: []common.Address{f.protocolAddress},
		Topics:    [][]common.Hash{{messageSentTopic}},
	}
}

// parseMessage decodes the TeleporterV2 message that [message] was sent as. The payload is
// the raw data of a TeleporterV2MessageSent(bytes) event, so decoding is two steps: unpack
// the ABI-encoded event to get the inner bytes, then parse those bytes into the message struct.
func parseMessage(
	message *messages.SourceMessage,
) (*teleportermessengerv2.TeleporterMessageV2, error) {
	event, err := zkAdapterABI.Unpack("TeleporterV2MessageSent", message.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed unpacking TeleporterV2MessageSent event: %w", err)
	}
	if len(event) != 1 {
		return nil, fmt.Errorf("unexpected TeleporterV2MessageSent argument count: %d", len(event))
	}
	encodedMessage, ok := event[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("failed extracting encoded message from event")
	}
	decoded, err := messageArguments.Unpack(encodedMessage)
	if err != nil {
		return nil, fmt.Errorf("failed parsing TeleporterV2 message: %w", err)
	}
	if len(decoded) != 1 {
		return nil, fmt.Errorf("unexpected TeleporterV2 message argument count: %d", len(decoded))
	}
	teleporterMessage := abi.ConvertType(
		decoded[0], new(teleportermessengerv2.TeleporterMessageV2),
	).(*teleportermessengerv2.TeleporterMessageV2)
	return teleporterMessage, nil
}

// encodeAttestation ABI-encodes the proofs as the Attestation struct that ZKAdapter.verifyMessage
// decodes with abi.decode(attestation, (Attestation)). Decoding a single dynamic struct expects
// a leading 0x20 offset word followed by the struct's encoding, and the struct's encoding is the
// same as its fields packed as an argument list, which is what verifyLogAndExtract's inputs
// produce. So the attestation is that packed body with the offset word prepended.
//
//nolint:unused // Called by ProcessMessage once message delivery lands in the follow-up PR.
func encodeAttestation(
	execProof *zkadapter.ExecutionProof,
	receiptProof *proofs.ReceiptProof,
) ([]byte, error) {
	body, err := attestationArguments.Pack(*execProof, toBindingReceiptProof(receiptProof))
	if err != nil {
		return nil, fmt.Errorf("failed to pack attestation proofs: %w", err)
	}

	offset := make([]byte, 32)
	offset[31] = 0x20
	return append(offset, body...), nil
}

// toBindingReceiptProof converts the proofs package's receipt proof to the adapter's
// Receipt.Proof binding struct.
func toBindingReceiptProof(p *proofs.ReceiptProof) zkadapter.ReceiptProof {
	return zkadapter.ReceiptProof{
		Proof:           p.Proof,
		Key:             p.Key,
		Value:           p.Value,
		LogIndex:        new(big.Int).SetUint64(uint64(p.LogIndex)),
		ExpectedEmitter: p.ExpectedEmitter,
		ExpectedTopic0:  p.ExpectedTopic0,
	}
}
