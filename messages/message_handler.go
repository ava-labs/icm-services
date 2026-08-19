// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_message_handler.go -package=mocks

package messages

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
)

// MessageManager is specific to each message protocol. The interface handles choosing which messages to send
// for each message protocol, and performs the sending to the destination chain.
// Message protocols are only handed the protocol agnostic [SourceMessage] observed on the source
// chain, so that any contract that originates messages, e.g. an IAdapter implementation, can be relayed
// without the relayer having to know how that protocol encodes or proves its messages.
type MessageHandlerFactory interface {
	// EventFilter returns the source chain log filter matching the messages sent by this message
	// protocol. The data of the logs it matches is the [SourceMessage] payload passed back
	// to this factory. The filter should constrain the emitting contract address, not just the
	// topics, so that other contracts emitting the same event signature are not matched.
	// Returns an empty filter for protocols whose messages are not read from source chain logs.
	EventFilter() evm.EventFilter

	// Create a message handler to relay the message
	NewMessageHandler(
		logger logging.Logger,
		message *SourceMessage,
		destinationClient vms.DestinationClient,
		signatureAggregator *aggregator.SignatureAggregator,
		metrics Metrics,
		signingSubnetID ids.ID,
		quorumNumerator uint64,
	) (MessageHandler, error)

	// Return info for routing the message to the correct relayer
	GetMessageRoutingInfo(message *SourceMessage) (MessageRoutingInfo, error)
}

// Struct containing fields for routing messages to the correct relayer.
type MessageRoutingInfo struct {
	SourceChainID      ids.ID
	SenderAddress      common.Address
	DestinationChainID ids.ID
	DestinationAddress common.Address
}

// Metrics records the outcome of relaying a single message. It is implemented by the
// caller (e.g. the relayer) so that the message handler can emit metrics without depending
// on the relayer package.
type Metrics interface {
	IncSuccessfulRelayMessageCount()
	IncFailedRelayMessageCount(failureReason string)
	IncFetchSignatureAppRequestCount()
	SetCreateSignedMessageLatencyMS(latency float64)
}

// MessageHandlers relay a single message. A new instance should be created for each message.
type MessageHandler interface {
	// ProcessMessage relays the message to the destination chain by aggregating a signature for it
	// and sending it via SendMessage. It does not retry on failure or checkpoint the height.
	// Returns the transaction hash if the message is successfully relayed.
	ProcessMessage() (common.Hash, error)
}
