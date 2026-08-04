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

// BlockHead carries every field a newHeads notification may include across
// the chain families the relayer supports: the standard geth header fields,
// the coreth and subnet-evm extras, and the SAE settlement fields (see
// HeaderSerializable in avalanchego's plugin/evm/customtypes). All fields are
// optional so notifications from chains that omit some of them still decode;
// the upstream serializable types cannot be reused here because each chain
// family's generated decoder requires fields the other family omits, and
// their "hash" is a marshal-only computed field that is dropped on unmarshal.
//
// The node-reported hash is kept verbatim: recomputing it client-side is not
// reliable for chains whose headers carry fields this client cannot encode
// (e.g. SAE chains).
type BlockHead struct {
	Hash             common.Hash      `json:"hash"`
	ParentHash       common.Hash      `json:"parentHash"`
	UncleHash        common.Hash      `json:"sha3Uncles"`
	Coinbase         common.Address   `json:"miner"`
	Root             common.Hash      `json:"stateRoot"`
	TxHash           common.Hash      `json:"transactionsRoot"`
	ReceiptHash      common.Hash      `json:"receiptsRoot"`
	Bloom            types.Bloom      `json:"logsBloom"`
	Difficulty       *hexutil.Big     `json:"difficulty"`
	Number           *hexutil.Big     `json:"number"`
	GasLimit         hexutil.Uint64   `json:"gasLimit"`
	GasUsed          hexutil.Uint64   `json:"gasUsed"`
	Time             hexutil.Uint64   `json:"timestamp"`
	Extra            hexutil.Bytes    `json:"extraData"`
	MixDigest        common.Hash      `json:"mixHash"`
	Nonce            types.BlockNonce `json:"nonce"`
	BaseFee          *hexutil.Big     `json:"baseFeePerGas"`
	WithdrawalsHash  *common.Hash     `json:"withdrawalsRoot"`
	BlobGasUsed      *hexutil.Uint64  `json:"blobGasUsed"`
	ExcessBlobGas    *hexutil.Uint64  `json:"excessBlobGas"`
	ParentBeaconRoot *common.Hash     `json:"parentBeaconBlockRoot"`

	// Avalanche extras (coreth and/or subnet-evm).
	ExtDataHash      common.Hash     `json:"extDataHash"`
	ExtDataGasUsed   *hexutil.Big    `json:"extDataGasUsed"`
	BlockGasCost     *hexutil.Big    `json:"blockGasCost"`
	TimeMilliseconds *hexutil.Uint64 `json:"timestampMilliseconds"`
	MinDelayExcess   *hexutil.Uint64 `json:"minDelayExcess"`
	TargetExponent   *hexutil.Uint64 `json:"targetExponent"`
	MinPriceExponent *hexutil.Uint64 `json:"minPriceExponent"`

	// SAE settlement fields (ACP-194).
	SettledHeight       *hexutil.Uint64 `json:"settledHeight"`
	SettledGasUnix      *hexutil.Uint64 `json:"settledGasUnix"`
	SettledGasNumerator *hexutil.Uint64 `json:"settledGasNumerator"`
	SettledExcess       *hexutil.Uint64 `json:"settledExcess"`
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
	head *BlockHead,
	ethClient ethereum.LogFilterer,
	topics [][]common.Hash,
	isPrimaryNetwork bool,
) (*ICMBlockInfo, error) {
	var (
		logs []types.Log
		err  error
	)
	// Only fetch logs when the block's bloom filter indicates that one of the event
	// topics we filter for (topics[0]) is present. A bloom match may be a false
	// positive; the FilterLogs call below performs the precise filtering.
	//
	// On the primary network (C-Chain) the header bloom filter cannot be relied on
	// as a shortcut: it summarises a settled predecessor range rather than the
	// block's own receipts, so the shortcut would silently miss events. Bypass the
	// bloom check there and always fetch the logs.
	if isPrimaryNetwork || bloomContainsAnyEventTopic(head.Bloom, topics) {
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
				Topics:    topics,
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
