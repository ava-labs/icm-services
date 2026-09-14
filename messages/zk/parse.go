// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"fmt"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/teleporterv2"
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
	teleporterMessage, err := teleporterv2.ParseTeleporterMessageV2(encodedMessage)
	if err != nil {
		return nil, fmt.Errorf("failed parsing TeleporterV2 message: %w", err)
	}
	return teleporterMessage, nil
}
