// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
)

var (
	WarpPrecompileLogFilter = warp.WarpABI.Events["SendWarpMessage"].ID
	ErrInvalidLog           = errors.New("invalid warp message log")
	ErrFailedToProcessLogs  = errors.New("failed to process logs")
)

// WarpBlockInfo describes the block height and logs needed to process Warp messages.
// WarpBlockInfo instances are populated by the subscriber, and forwarded to the Listener to process.
type WarpBlockInfo struct {
	BlockNumber uint64
	Messages    []*WarpMessageInfo
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

// Extract Warp logs from the block, if they exist
func NewWarpBlockInfo(logger logging.Logger, header *types.Header, ethClient ethclient.Client) (*WarpBlockInfo, error) {
	var (
		logs []types.Log
		err  error
	)
	// Check if the block contains warp logs, and fetch them from the client if it does
	if header.Bloom.Test(WarpPrecompileLogFilter[:]) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		operation := func() (err error) {
			logs, err = ethClient.FilterLogs(cctx, ethereum.FilterQuery{
				Topics:    [][]common.Hash{{WarpPrecompileLogFilter}},
				Addresses: []common.Address{warp.ContractAddress},
				FromBlock: header.Number,
				ToBlock:   header.Number,
			})
			return err
		}

		// We increase the timeout here to 30 seconds reducing the chance of hitting a race condition
		// where the block header is received via websocket subscription before the block's
		// logs are available via RPC. This is a known behavior in EVM nodes due to
		// asynchronous log/index processing after a block becomes canonical.
		timeout := utils.DefaultRPCTimeout * 6
		err = utils.WithRetriesTimeout(logger, operation, timeout, "get warp logs from block")
		if err != nil {
			return nil, err
		}
	}
	messages := make([]*WarpMessageInfo, len(logs))
	for i, log := range logs {
		warpLog, err := NewWarpMessageInfo(log)
		if err != nil {
			return nil, err
		}
		messages[i] = warpLog
	}

	return &WarpBlockInfo{
		BlockNumber: header.Number.Uint64(),
		Messages:    messages,
	}, nil
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

func LogsToBlocks(logs []types.Log) (map[uint64]*WarpBlockInfo, error) {
	blocks := make(map[uint64]*WarpBlockInfo)
	for _, log := range logs {
		warpMessageInfo, err := NewWarpMessageInfo(log)
		if err != nil {
			return nil, err
		}
		if block, ok := blocks[log.BlockNumber]; ok {
			block.Messages = append(block.Messages, warpMessageInfo)
		} else {
			blocks[log.BlockNumber] = &WarpBlockInfo{
				BlockNumber: log.BlockNumber,
				Messages:    []*WarpMessageInfo{warpMessageInfo},
			}
		}
	}
	return blocks, nil
}
