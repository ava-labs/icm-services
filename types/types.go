// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/customtypes"
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
	// On SAE-enabled chains (Helicon+), header.LogsBloom summarises the block's
	// newly-settled predecessor range rather than the block's own receipts (see
	// ACP-194 and avalanchego/vms/saevm/sae/rpc/indexing.go's bloomOverrider),
	// so the shortcut would silently miss events. On classic subnet-evm the
	// shortcut is safe. Detect SAE per-header via the settlement marker in the
	// customtypes header extra data, and always fetch when it's present.
	if isSAEHeader(header) || bloomContainsAnyEventTopic(header.Bloom, topics) {
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

// isSAEHeader reports whether the given header was produced by an SAE-enabled
// chain (per ACP-194). On SAE, the block builder encodes settlement metadata
// into HeaderExtra via customtypes; the presence of SettledHeight is a
// sufficient marker. On classic subnet-evm the field is nil.
//
// Referencing customtypes.HeaderExtra directly (rather than parsing raw extra
// bytes) means any upstream rename or type change surfaces at compile time.
func isSAEHeader(h *types.Header) bool {
	return customtypes.GetHeaderExtra(h).SettledHeight != nil
}

// bloomContainsAnyEventTopic reports whether the header's bloom indicates the
// presence of at least one of the event-signature topics being filtered for
// (topics[0]). Only meaningful on classic subnet-evm; on SAE, callers must
// bypass this check because header.Bloom refers to the settled predecessor
// range rather than this block's own receipts.
//
// If no event-signature topics are provided, conservatively returns true so
// logs are still fetched. A positive result may be a false positive due to
// the probabilistic nature of bloom filters; callers must perform precise
// filtering afterwards.
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
