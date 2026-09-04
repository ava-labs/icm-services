// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/teleporterv2"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
)

var (
	messageProtocolAddress       = common.HexToAddress("0xa0000000000000000000000000000000000001")
	destinationTeleporterAddress = common.HexToAddress("0xa000000000000000000000000000000000002")
	rewardAddress                = common.HexToAddress("0xa000000000000000000000000000000000003")

	originSenderAddress     = common.HexToAddress("0xa000000000000000000000000000000000004")
	sourceTeleporterAddress = common.HexToAddress("0xa000000000000000000000000000000000005")
	destinationAppAddress   = common.HexToAddress("0xa000000000000000000000000000000000006")

	messageProtocolConfig = config.MessageProtocolConfig{
		MessageFormat: "zk", // TODO: match the registered protocol type
		Settings: map[string]interface{}{
			"reward-address":      rewardAddress.Hex(),
			"adapter-address":     messageProtocolAddress.Hex(),
			"teleporter-address":  destinationTeleporterAddress.Hex(),
			"beacon-genesis-time": float64(1606824023),
			"seconds-per-slot":    float64(12),
			"beacon-api-url":      "http://localhost:5052",
		},
	}
	destinationBlockchainIDString = "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD"
	destinationBlockchainID       ids.ID
	sourceBlockchainID            = ids.Empty
	validTeleporterMessage        *teleportermessengerv2.TeleporterMessageV2
	validEventPayload             []byte
)

func init() {
	var err error
	destinationBlockchainID, err = ids.FromString(destinationBlockchainIDString)
	if err != nil {
		panic(err)
	}
	validTeleporterMessage = &teleportermessengerv2.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		OriginSenderAddress:     originSenderAddress,
		OriginTeleporterAddress: sourceTeleporterAddress,
		DestinationBlockchainID: destinationBlockchainID,
		DestinationAddress:      destinationAppAddress,
		RequiredGasLimit:        big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []teleportermessengerv2.TeleporterMessageReceipt{},
		Message:                 []byte{1, 2, 3, 4},
	}

	// Build the event payload the adapter would emit
	serialized := teleporterv2.SerializeTeleporterMessageV2(validTeleporterMessage)
	validEventPayload, err = zkAdapterABI.Events["TeleporterV2MessageSent"].Inputs.Pack(serialized)
	if err != nil {
		panic(err)
	}
}

// sourceMessage returns the source chain message as the relayer would pass it
// to the message handler factory.
func sourceMessage(payload []byte) *messages.SourceMessage {
	return &messages.SourceMessage{
		SourceBlockchainID: sourceBlockchainID,
		ProtocolAddress:    messageProtocolAddress,
		Payload:            payload,
		SourceTxID:         common.HexToHash("0x01"),
	}
}

func TestParseMessage(t *testing.T) {
	testCases := []struct {
		name    string
		payload []byte
		isError bool
	}{
		{name: "valid", payload: validEventPayload, isError: false},
		{name: "empty payload", payload: []byte{}, isError: true},
		{name: "garbage payload", payload: []byte{1, 2, 3, 4}, isError: true},
		{
			name: "valid framing, truncated message",
			payload: func() []byte {
				p, err := zkAdapterABI.Events["TeleporterV2MessageSent"].Inputs.Pack([]byte{1, 2, 3})
				require.NoError(t, err)
				return p
			}(),
			isError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := parseMessage(sourceMessage(test.payload))
			if test.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, validTeleporterMessage, parsed)
		})
	}
}

func TestGetMessageRoutingInfo(t *testing.T) {
	factory, err := NewMessageHandlerFactory(
		messageProtocolAddress,
		messageProtocolConfig,
		nil,
	)
	require.NoError(t, err)

	info, err := factory.GetMessageRoutingInfo(sourceMessage(validEventPayload))
	require.NoError(t, err)
	require.Equal(t, sourceBlockchainID, info.SourceChainID)
	require.Equal(t, messageProtocolAddress, info.SenderAddress)
	require.Equal(t, destinationBlockchainID, info.DestinationChainID)
	require.Equal(t, messageProtocolAddress, info.DestinationAddress)
}

func TestFindMessageLog(t *testing.T) {
	makeLog := func(addr common.Address, topic common.Hash, data []byte) *types.Log {
		return &types.Log{Address: addr, Topics: []common.Hash{topic}, Data: data}
	}
	otherAddress := common.HexToAddress("0xffff")
	otherTopic := common.HexToHash("0xeeee")

	receipt := &types.Receipt{Logs: []*types.Log{
		makeLog(otherAddress, messageSentTopic, validEventPayload),           // wrong emitter
		makeLog(messageProtocolAddress, otherTopic, validEventPayload),       // wrong event type
		makeLog(messageProtocolAddress, messageSentTopic, []byte{9, 9}),      // wrong payload
		makeLog(messageProtocolAddress, messageSentTopic, validEventPayload), // ours
	}}

	index, err := findMessageLog(receipt, messageProtocolAddress, validEventPayload)
	require.NoError(t, err)
	require.Equal(t, uint(3), index)

	_, err = findMessageLog(receipt, messageProtocolAddress, []byte{1, 2, 3})
	require.Error(t, err)
}
