// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_message_handler.go -package=mocks

package messages

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ethereum/go-ethereum/common"
)

// MessageManager is specific to each message protocol. The interface handles choosing which messages to send
// for each message protocol, and performs the sending to the destination chain.
type MessageHandlerFactory interface {
	// Create a message handler to relay the Warp message
	NewMessageHandler(
		unsignedMessage *warp.UnsignedMessage,
		destinationClient vms.DestinationClient,
	) (MessageHandler, error)

	// Return info for routing the message to the correct relayer
	GetMessageRoutingInfo(unsignedMessage *warp.UnsignedMessage) (MessageRoutingInfo, error)
}

// Struct containing fields for routing messages to the correct relayer.
type MessageRoutingInfo struct {
	SourceChainID      ids.ID
	SenderAddress      common.Address
	DestinationChainID ids.ID
	DestinationAddress common.Address
}

// MessageHandlers relay a single Warp message. A new instance should be created for each Warp message.
type MessageHandler interface {
	// ShouldSendMessage returns true if the message should be sent to the destination chain
	// If an error is returned, the boolean should be ignored by the caller.
	ShouldSendMessage() (bool, error)

	// SendMessage sends the signed message to the destination chain. The payload parsed according to
	// the VM rules is also passed in, since MessageManager does not assume any particular VM
	// returns the transaction hash if the transaction is successful.
	SendMessage(signedMessage *warp.Message) (common.Hash, error)

	// LoggerWithContext returns a logger with the message context
	LoggerWithContext(logging.Logger) logging.Logger

	// GetUnsignedMessage returns the unsigned message
	GetUnsignedMessage() *warp.UnsignedMessage
}
