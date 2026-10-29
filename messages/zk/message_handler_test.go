// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/accounts/abi"
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
		MessageFormat: "zk",
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
	serialized, err := messageArguments.Pack(*validTeleporterMessage)
	if err != nil {
		panic(err)
	}
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

func testExecProof() *zkadapter.ExecutionProof {
	return &zkadapter.ExecutionProof{
		AnchorSlot:                 200,
		TargetSlot:                 100,
		AnchorBeaconStateRoot:      [32]byte{0x01},
		AnchorBeaconStateProof:     [][32]byte{{0x02}, {0x03}, {0x04}},
		TargetBeaconStateRoot:      [32]byte{0x05},
		TargetBeaconStateProof:     [][32]byte{{0x06}},
		TargetExecutionHeaderRoot:  [32]byte{0x07},
		TargetExecutionHeaderProof: [][32]byte{{0x08}, {0x09}},
		TargetReceiptsRoot:         [32]byte{0x0a},
		TargetReceiptsProof:        [][32]byte{{0x0b}},
	}
}

func testReceiptProof() *proofs.ReceiptProof {
	return &proofs.ReceiptProof{
		Proof:           [][]byte{{0xaa, 0xbb}, {0xcc}},
		Key:             []byte{0x0f},
		Value:           []byte{0x01, 0x02, 0x03},
		LogIndex:        2,
		ExpectedEmitter: common.HexToAddress("0xa0000000000000000000000000000000000001"),
		ExpectedTopic0:  common.HexToHash("0xdd"),
	}
}

// Checks that the attestation is the one-element-tuple encoding verifyMessage's abi.decode
// expects, and that the proofs decode back to the same values.
func TestEncodeAttestation(t *testing.T) {
	execProof := testExecProof()
	receiptProof := testReceiptProof()

	attestation, err := encodeAttestation(execProof, receiptProof)
	require.NoError(t, err)

	// Leading word is the offset to the struct body.
	require.Equal(t, common.LeftPadBytes([]byte{0x20}, 32), attestation[:32])

	// The body decodes back to the same two proofs through the same ABI.
	decoded, err := attestationArguments.Unpack(attestation[32:])
	require.NoError(t, err)
	require.Len(t, decoded, 2)

	gotExec := abi.ConvertType(decoded[0], new(zkadapter.ExecutionProof)).(*zkadapter.ExecutionProof)
	require.Equal(t, execProof, gotExec)

	gotReceipt := abi.ConvertType(decoded[1], new(zkadapter.ReceiptProof)).(*zkadapter.ReceiptProof)
	require.Equal(t, toBindingReceiptProof(receiptProof), *gotReceipt)
}

// Checks the conversion to the adapter's Receipt.Proof binding struct, including the
// uint to *big.Int log index.
func TestToBindingReceiptProof(t *testing.T) {
	p := testReceiptProof()
	got := toBindingReceiptProof(p)

	require.Equal(t, p.Proof, got.Proof)
	require.Equal(t, p.Key, got.Key)
	require.Equal(t, p.Value, got.Value)
	require.Equal(t, big.NewInt(2), got.LogIndex)
	require.Equal(t, p.ExpectedEmitter, got.ExpectedEmitter)
	require.Equal(t, [32]byte(p.ExpectedTopic0), got.ExpectedTopic0)
}
