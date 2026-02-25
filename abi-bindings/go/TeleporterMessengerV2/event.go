// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessengerv2

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// Event is a Teleporter log event
type Event uint8

const (
	Unknown Event = iota
	SendCrossChainMessage
	ReceiveCrossChainMessage
	AddFeeAmount
	MessageExecutionFailed
	MessageExecuted
	RelayerRewardsRedeemed
	ReceiptReceived

	sendCrossChainMessageStr    = "SendCrossChainMessage"
	receiveCrossChainMessageStr = "ReceiveCrossChainMessage"
	addFeeAmountStr             = "AddFeeAmount"
	messageExecutionFailedStr   = "MessageExecutionFailed"
	messageExecutedStr          = "MessageExecuted"
	relayerRewardsRedeemedStr   = "RelayerRewardsRedeemed"
	receiptReceivedStr          = "ReceiptReceived"
	unknownStr                  = "Unknown"
)

// String returns the string representation of an Event
func (e Event) String() string {
	switch e {
	case SendCrossChainMessage:
		return sendCrossChainMessageStr
	case ReceiveCrossChainMessage:
		return receiveCrossChainMessageStr
	case AddFeeAmount:
		return addFeeAmountStr
	case MessageExecutionFailed:
		return messageExecutionFailedStr
	case MessageExecuted:
		return messageExecutedStr
	case RelayerRewardsRedeemed:
		return relayerRewardsRedeemedStr
	case ReceiptReceived:
		return receiptReceivedStr
	default:
		return unknownStr
	}
}

// ToEvent converts a string to an Event
func ToEvent(e string) (Event, error) {
	switch strings.ToLower(e) {
	case strings.ToLower(sendCrossChainMessageStr):
		return SendCrossChainMessage, nil
	case strings.ToLower(receiveCrossChainMessageStr):
		return ReceiveCrossChainMessage, nil
	case strings.ToLower(addFeeAmountStr):
		return AddFeeAmount, nil
	case strings.ToLower(messageExecutionFailedStr):
		return MessageExecutionFailed, nil
	case strings.ToLower(messageExecutedStr):
		return MessageExecuted, nil
	case strings.ToLower(relayerRewardsRedeemedStr):
		return RelayerRewardsRedeemed, nil
	case strings.ToLower(receiptReceivedStr):
		return ReceiptReceived, nil
	default:
		return Unknown, fmt.Errorf("unknown event %s", e)
	}
}

// FilterTeleporterEvents parses the topics and data of a Teleporter log into the corresponding Teleporter event
func FilterTeleporterEvents(topics []common.Hash, data []byte, event string) (fmt.Stringer, error) {
	e, err := ToEvent(event)
	if err != nil {
		return nil, err
	}
	var out fmt.Stringer
	switch e {
	case SendCrossChainMessage:
		out = new(TeleporterMessengerV2SendCrossChainMessage)
	case ReceiveCrossChainMessage:
		out = new(TeleporterMessengerV2ReceiveCrossChainMessage)
	case AddFeeAmount:
		out = new(TeleporterMessengerV2AddFeeAmount)
	case MessageExecutionFailed:
		out = new(TeleporterMessengerV2MessageExecutionFailed)
	case MessageExecuted:
		out = new(TeleporterMessengerV2MessageExecuted)
	case RelayerRewardsRedeemed:
		out = new(TeleporterMessengerV2RelayerRewardsRedeemed)
	case ReceiptReceived:
		out = new(TeleporterMessengerV2ReceiptReceived)
	default:
		return nil, fmt.Errorf("unknown event %s", e.String())
	}
	if err := UnpackEvent(out, e.String(), topics, data); err != nil {
		return nil, err
	}
	return out, nil
}

func (t TeleporterMessengerV2SendCrossChainMessage) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2SendCrossChainMessage{
		MessageID:               common.Hash(t.MessageID),
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		Message:                 toReadableTeleporterMessageV2(t.Message),
		FeeInfo:                 t.FeeInfo,
		Raw:                     t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2SendCrossChainMessage struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	Message                 ReadableTeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

func (t TeleporterMessengerV2ReceiveCrossChainMessage) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2ReceiveCrossChainMessage{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Deliverer:          t.Deliverer,
		RewardRedeemer:     t.RewardRedeemer,
		Message:            toReadableTeleporterMessage(t.Message),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2ReceiveCrossChainMessage struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

func (t TeleporterMessengerV2AddFeeAmount) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2AddFeeAmount{
		MessageID:      common.Hash(t.MessageID),
		UpdatedFeeInfo: t.UpdatedFeeInfo,
		Raw:            t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2AddFeeAmount struct {
	MessageID      common.Hash
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log
}

func (t TeleporterMessengerV2MessageExecutionFailed) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2MessageExecutionFailed{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Message:            toReadableTeleporterMessage(t.Message),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2MessageExecutionFailed struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

func (t TeleporterMessengerV2MessageExecuted) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2MessageExecuted{
		MessageID:          common.Hash(t.MessageID),
		SourceBlockchainID: ids.ID(t.SourceBlockchainID),
		Raw:                t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2MessageExecuted struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Raw                types.Log
}

func (t TeleporterMessengerV2RelayerRewardsRedeemed) String() string {
	outJson, _ := json.MarshalIndent(t, "", "  ")

	return string(outJson)
}

func (t TeleporterMessengerV2ReceiptReceived) String() string {
	outJson, _ := json.MarshalIndent(ReadableTeleporterMessengerV2ReceiptReceived{
		MessageID:               common.Hash(t.MessageID),
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		RelayerRewardAddress:    t.RelayerRewardAddress,
		FeeInfo:                 t.FeeInfo,
		Raw:                     t.Raw,
	}, "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessengerV2ReceiptReceived struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

func toReadableTeleporterMessage(t TeleporterMessage) ReadableTeleporterMessage {
	return ReadableTeleporterMessage{
		MessageNonce:            t.MessageNonce,
		OriginSenderAddress:     t.OriginSenderAddress,
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		DestinationAddress:      t.DestinationAddress,
		RequiredGasLimit:        t.RequiredGasLimit,
		AllowedRelayerAddresses: t.AllowedRelayerAddresses,
		Receipts:                t.Receipts,
		Message:                 t.Message,
	}
}

func toReadableTeleporterMessageV2(t TeleporterMessageV2) ReadableTeleporterMessage {
	return ReadableTeleporterMessage{
		MessageNonce:            t.MessageNonce,
		OriginSenderAddress:     t.OriginSenderAddress,
		DestinationBlockchainID: ids.ID(t.DestinationBlockchainID),
		DestinationAddress:      t.DestinationAddress,
		RequiredGasLimit:        t.RequiredGasLimit,
		AllowedRelayerAddresses: t.AllowedRelayerAddresses,
		Receipts:                t.Receipts,
		Message:                 t.Message,
	}
}

func (t TeleporterMessage) String() string {
	outJson, _ := json.MarshalIndent(toReadableTeleporterMessage(t), "", "  ")

	return string(outJson)
}

func (t TeleporterMessageV2) String() string {
	outJson, _ := json.MarshalIndent(toReadableTeleporterMessageV2(t), "", "  ")

	return string(outJson)
}

type ReadableTeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID ids.ID
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}
