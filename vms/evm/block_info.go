// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// ICMBlockInfo describes a contiguous range of source chain blocks together with the logs they
// contain that match the subscriber's event filter. ICMBlockInfo instances are populated by the
// subscriber, and forwarded to the Listener to process.
//
// Blocks without matching logs are not reported on their own: the subscriber folds them into the
// range of a neighbouring block, so that processing every ICMBlockInfo accounts for every block
// and the checkpoint manager can advance past blocks that produced no logs.
type ICMBlockInfo struct {
	// FromBlock and ToBlock are the first and last heights, inclusive, of the blocks covered.
	FromBlock uint64
	ToBlock   uint64
	// Logs are the logs of the covered blocks that match the subscriber's event filter, in order.
	Logs      []types.Log
	IsCatchup bool
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

// BlockHead is the subset of a block header the relayer uses, decoded leniently
// so it works across chain families. The node-reported hash is kept verbatim:
// recomputing it client-side is not reliable for chains whose headers carry
// fields this client cannot encode (e.g. SAE chains). The upstream header types
// cannot be reused here: each family's generated decoder requires fields the
// other family omits, and their "hash" is a marshal-only computed field that is
// dropped on unmarshal.
//
// If more header fields are needed later, the wire format is defined by
// HeaderSerializable in avalanchego's
// graft/coreth/plugin/evm/customtypes/header_ext.go (C-Chain, including the
// SAE settlement fields) and
// graft/subnet-evm/plugin/evm/customtypes/header_ext.go (subnet-evm chains).
type BlockHead struct {
	Hash   common.Hash  `json:"hash"`
	Number *hexutil.Big `json:"number"`
	Bloom  types.Bloom  `json:"logsBloom"`
}

// NewICMBlockInfo extracts the logs matching [filter] from the block with head [head], if any.
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
		logs, err = FilterLogsByBlockHash(logger, ethClient, filter, head.Hash)
		if err != nil {
			return nil, err
		}
	}

	blockNumber := head.Number.ToInt().Uint64()
	return &ICMBlockInfo{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Logs:      logs,
		IsCatchup: false,
	}, nil
}

// FilterLogsByBlockHash fetches the logs of the block with hash [blockHash] that match [filter].
func FilterLogsByBlockHash(
	logger logging.Logger,
	ethClient ethereum.LogFilterer,
	filter EventFilter,
	blockHash common.Hash,
) ([]types.Log, error) {
	var logs []types.Log
	// Query by hash: a node that doesn't know the block errors ("unknown
	// block") and is retried below, whereas a by-number query would return
	// empty logs with no error and the block would be silently skipped.
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

	// Blocks are learned of via WS before every node behind a load-balanced RPC
	// endpoint knows them, so allow several retries for the "unknown block"
	// case above.
	timeout := utils.DefaultRPCTimeout * 6
	if err := utils.WithRetriesTimeout(operation, notify, timeout); err != nil {
		return nil, fmt.Errorf("failed to get logs for block: %w", err)
	}
	return logs, nil
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
