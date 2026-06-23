// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
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

// WarpMessageInfo describes the transaction information for the Warp message
// sent on the source chain.
// WarpMessageInfo instances are either derived from the logs of a block or
// from the manual Warp message information provided via configuration.
type WarpMessageInfo struct {
	SourceAddress   common.Address
	SourceTxID      common.Hash
	UnsignedMessage *avalancheWarp.UnsignedMessage
}

// Extract Warp logs from the block, if they exist.
func NewICMBlockInfo(
	logger logging.Logger,
	header *types.Header,
	ethClient ethereum.LogFilterer,
	topics [][]common.Hash,
) (*ICMBlockInfo, error) {
	var (
		logs []types.Log
		err  error
	)
	// Check if the block contains warp logs, and fetch them from the client if it does.
	// We test the block bloom against the event-signature topics we actually filter for
	// (topics[0]) rather than a single hardcoded event hash. The warp precompile emits a
	// SendWarpMessage log whenever a message is sent (for both the Teleporter v1 path and the
	// TeleporterV2/Merkle adapter path), so gating on the warp event signature reliably detects
	// relayable blocks across protocol versions. A bloom match may be a false positive; the
	// FilterLogs call below performs the precise filtering.
	if bloomContainsAnyEventTopic(header.Bloom, topics) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		operation := func() (err error) {
			logs, err = ethClient.FilterLogs(cctx, ethereum.FilterQuery{
				Topics:    topics,
				FromBlock: header.Number,
				ToBlock:   header.Number,
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

		// We increase the timeout here to 30 seconds reducing the chance of hitting a race condition
		// where the block header is received via websocket subscription before the block's
		// logs are available via RPC. This is a known behavior in EVM nodes due to
		// asynchronous log/index processing after a block becomes canonical.
		timeout := utils.DefaultRPCTimeout * 6
		err = utils.WithRetriesTimeout(operation, notify, timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to get logs for block: %w", err)
		}
	}

	return &ICMBlockInfo{
		BlockNumber: header.Number.Uint64(),
		Logs:        logs,
		IsCatchup:   false,
	}, nil
}

// bloomContainsAnyEventTopic reports whether the block's bloom filter indicates the presence of at
// least one of the event-signature topics being filtered for (the first topic position, topics[0]).
// If no event-signature topics are provided, it conservatively returns true so that logs are still
// fetched. A positive result may be a false positive due to the probabilistic nature of bloom
// filters; callers must perform precise filtering afterwards.
func bloomContainsAnyEventTopic(bloom types.Bloom, topics [][]common.Hash) bool {
	if len(topics) == 0 || len(topics[0]) == 0 {
		return true
	}
	for _, eventTopic := range topics[0] {
		if bloom.Test(eventTopic[:]) {
			return true
		}
	}
	return false
}

// Extract the Warp message information from the raw log
func NewWarpMessageInfo(log types.Log) (*WarpMessageInfo, error) {
	if len(log.Topics) != 3 {
		return nil, ErrInvalidLog
	}
	if log.Topics[0] != WarpPrecompileLogFilter {
		return nil, ErrInvalidLog
	}
	unsignedMsg, err := UnpackWarpMessage(log.Data)
	if err != nil {
		return nil, err
	}

	return &WarpMessageInfo{
		SourceAddress:   common.BytesToAddress(log.Topics[1][:]),
		SourceTxID:      log.TxHash,
		UnsignedMessage: unsignedMsg,
	}, nil
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
