// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"errors"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

var (
	WarpPrecompileLogFilter = warp.WarpABI.Events["SendWarpMessage"].ID
	ErrInvalidLog           = errors.New("invalid warp message log")
)

// SourceMessage describes a message sent from a source blockchain that the relayer may deliver to
// its destination chain. It is agnostic to the message protocol that sent the message: each
// protocol's MessageHandlerFactory is responsible for decoding [Payload] into that protocol's own
// message representation.
// SourceMessage instances are either derived from the logs of a block, or from the message
// information provided directly to the relayer API.
type SourceMessage struct {
	// SourceBlockchainID is the ID of the blockchain that the message was sent from.
	SourceBlockchainID ids.ID
	// ProtocolAddress is the address of the message protocol contract that sent the message. The
	// message is relayed by the message handler registered for this address.
	ProtocolAddress common.Address
	// Payload is the message as encoded by the message protocol that sent it. Protocols that send
	// their messages through the Warp precompile set this to the encoded unsigned Warp message.
	// Protocols that emit their messages as contract events set it to the event data.
	Payload []byte
	// SourceTxID is the hash of the transaction that the message was emitted in. It is the zero
	// hash for messages provided directly to the relayer API rather than read from a log.
	SourceTxID common.Hash
}

// NewSourceMessage returns the message contained in [log], which was emitted on
// [sourceBlockchainID] by the message protocol deployed at [protocolAddress].
// The log data is the message protocol's encoding of the message, so it is passed through
// unparsed for the message protocol's handler to decode.
func NewSourceMessage(
	sourceBlockchainID ids.ID,
	protocolAddress common.Address,
	log types.Log,
) *SourceMessage {
	return &SourceMessage{
		SourceBlockchainID: sourceBlockchainID,
		ProtocolAddress:    protocolAddress,
		Payload:            log.Data,
		SourceTxID:         log.TxHash,
	}
}

// NewSourceMessageFromWarpLog returns the message contained in [log], which must be a
// SendWarpMessage event emitted by the Warp precompile on [sourceBlockchainID]. The message
// protocol that sent the message is the sender of the Warp message, which is indexed as the
// second topic of the log.
func NewSourceMessageFromWarpLog(sourceBlockchainID ids.ID, log types.Log) (*SourceMessage, error) {
	if log.Address != warp.ContractAddress {
		return nil, ErrInvalidLog
	}
	if len(log.Topics) != 3 {
		return nil, ErrInvalidLog
	}
	if log.Topics[0] != WarpPrecompileLogFilter {
		return nil, ErrInvalidLog
	}
	return NewSourceMessage(sourceBlockchainID, common.BytesToAddress(log.Topics[1][:]), log), nil
}

// WarpEventFilter returns the source chain log filter matching the Warp messages sent by the
// message protocol contract at [protocolAddress]. Warp messages are emitted as SendWarpMessage
// events by the Warp precompile, with the sending protocol's address as the first indexed topic.
func WarpEventFilter(protocolAddress common.Address) evm.EventFilter {
	return evm.EventFilter{
		Addresses: []common.Address{warp.ContractAddress},
		Topics: [][]common.Hash{
			{WarpPrecompileLogFilter},
			{common.BytesToHash(protocolAddress[:])},
		},
	}
}
