// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

var (
	WarpPrecompileLogFilter = warp.WarpABI.Events["SendWarpMessage"].ID
	ErrInvalidLog           = errors.New("invalid warp message log")
	ErrFailedToProcessLogs  = errors.New("failed to process logs")
)

// ICMBlockInfo describes the block height and logs needed to process Warp messages.
// ICMBlockInfo instances are populated by the subscriber, and forwarded to the Listener to process.
type ICMBlockInfo struct {
	BlockNumber uint64
	Logs        []types.Log
	IsCatchup   bool
}

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

// EventFilter selects the source chain logs that carry a message protocol's messages, following
// the semantics of ethereum.FilterQuery: a log matches if it was emitted by one of [Addresses]
// (any address, if empty) and its topics match [Topics].
// Filters should constrain the emitting address whenever possible: topics can be forged by any
// contract, so a filter that only constrains topics can match logs that were not emitted by the
// message protocol.
type EventFilter struct {
	// Addresses are the contract addresses that emit the protocol's message logs. For protocols
	// that send their messages through the Warp precompile, this is the precompile's address
	// rather than the protocol's own.
	Addresses []common.Address
	// Topics constrain the topics of matching logs, in the format expected by
	// ethereum.FilterQuery.Topics.
	Topics [][]common.Hash
}

// IsEmpty reports whether the filter constrains neither the emitting address nor the log topics.
// Message protocols whose messages are not read from source chain logs return an empty filter.
func (f EventFilter) IsEmpty() bool {
	return len(f.Addresses) == 0 && len(f.Topics) == 0
}

// BlockHead is the subset of a newHeads notification the relayer uses,
// decoded leniently so it works across chain families. The node-reported hash
// is kept verbatim: recomputing it client-side is not reliable for chains
// whose headers carry fields this client cannot encode (e.g. SAE chains).
// The upstream header types cannot be reused here: each family's generated
// decoder requires fields the other family omits, and their "hash" is a
// marshal-only computed field that is dropped on unmarshal.
//
// If more notification fields are needed later, the wire format is defined by
// HeaderSerializable in avalanchego's
// graft/coreth/plugin/evm/customtypes/header_ext.go (C-Chain, including the
// SAE settlement fields) and
// graft/subnet-evm/plugin/evm/customtypes/header_ext.go (subnet-evm chains).
type BlockHead struct {
	Hash   common.Hash  `json:"hash"`
	Number *hexutil.Big `json:"number"`
	Bloom  types.Bloom  `json:"logsBloom"`
}

// WarpMessageInfo describes the transaction information for the Warp message
// sent on the source chain.
// WarpMessageInfo instances are either derived from the logs of a block or
// from the manual Warp message information provided via configuration.
type WarpMessageInfo struct {
	SourceAddress   common.Address
	SourceTxID      common.Hash
	UnsignedMessage *avalancheWarp.UnsignedMessage
}

func NewICMBlockInfo(
	logger logging.Logger,
	head *BlockHead,
	ethClient ethereum.LogFilterer,
	filter EventFilter,
	isPrimaryNetwork bool,
) (*ICMBlockInfo, error) {
	var (
		logs []types.Log
		err  error
	)
	// Only fetch logs when the block's bloom filter indicates that the filter's emitting
	// addresses and event topics are present. A bloom match may be a false positive; the
	// FilterLogs call below performs the precise filtering.
	//
	// On the primary network (C-Chain) the header bloom filter cannot be relied on
	// as a shortcut: it summarises a settled predecessor range rather than the
	// block's own receipts, so the shortcut would silently miss events. Bypass the
	// bloom check there and always fetch the logs.
	if isPrimaryNetwork || bloomMatchesFilter(head.Bloom, filter) {
		// Query by hash: a node that doesn't know the block errors ("unknown
		// block") and is retried below, whereas a by-number query would return
		// empty logs with no error and the block would be silently skipped.
		blockHash := head.Hash
		operation := func() (err error) {
			// Fresh context per attempt so retries aren't killed by an
			// already-expired deadline.
			cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
			defer cancel()
			logs, err = ethClient.FilterLogs(cctx, ethereum.FilterQuery{
				Addresses: filter.Addresses,
				Topics:    filter.Topics,
				BlockHash: &blockHash,
			})
			return err
		}
		notify := func(err error, duration time.Duration) {
			logger.Info(
				"getting ICM block from logs failed, retrying...",
				zap.Duration("retryIn", duration),
				zap.Error(err),
			)
		}

		// Headers arrive via WS before every node behind a load-balanced RPC
		// endpoint knows the block, so allow several retries for the "unknown
		// block" case above.
		timeout := utils.DefaultRPCTimeout * 6
		err = utils.WithRetriesTimeout(operation, notify, timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to get logs for block: %w", err)
		}
	}

	return &ICMBlockInfo{
		BlockNumber: head.Number.ToInt().Uint64(),
		Logs:        logs,
		IsCatchup:   false,
	}, nil
}

// bloomMatchesFilter reports whether the block's bloom filter indicates the presence of logs
// matching [filter]: at least one of the filter's emitting addresses and at least one of its
// event-signature topics (the first topic position, Topics[0]). Filter dimensions that are not
// constrained are conservatively treated as matching so that logs are still fetched. A positive
// result may be a false positive due to the probabilistic nature of bloom filters; callers must
// perform precise filtering afterwards.
func bloomMatchesFilter(bloom types.Bloom, filter EventFilter) bool {
	if len(filter.Addresses) > 0 {
		anyAddress := false
		for _, address := range filter.Addresses {
			if bloom.Test(address[:]) {
				anyAddress = true
				break
			}
		}
		if !anyAddress {
			return false
		}
	}
	if len(filter.Topics) > 0 && len(filter.Topics[0]) > 0 {
		anyTopic := false
		for _, eventTopic := range filter.Topics[0] {
			if bloom.Test(eventTopic[:]) {
				anyTopic = true
				break
			}
		}
		if !anyTopic {
			return false
		}
	}
	return true
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
func WarpEventFilter(protocolAddress common.Address) EventFilter {
	return EventFilter{
		Addresses: []common.Address{warp.ContractAddress},
		Topics: [][]common.Hash{
			{WarpPrecompileLogFilter},
			{common.BytesToHash(protocolAddress[:])},
		},
	}
}

func UnpackWarpMessage(unsignedMsgBytes []byte) (*avalancheWarp.UnsignedMessage, error) {
	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(unsignedMsgBytes)
	if err != nil {
		// If we failed to parse the message as a log, attempt to parse it as a standalone message
		var standaloneErr error
		unsignedMsg, standaloneErr = avalancheWarp.ParseUnsignedMessage(unsignedMsgBytes)
		if standaloneErr != nil {
			err = errors.Join(err, standaloneErr)
			return nil, err
		}
	}
	return unsignedMsg, nil
}
