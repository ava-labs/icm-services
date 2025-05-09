// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	relayerTypes "github.com/ava-labs/icm-services/types"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	// Max buffer size for ethereum subscription channels
	maxClientSubscriptionBuffer = 20000
	MaxBlocksPerRequest         = 200
)

// subscriber implements Subscriber
type subscriber struct {
	wsClient     ethclient.Client
	rpcClient    ethclient.Client
	blockchainID ids.ID
	headers      chan *types.Header
	icmBlocks    chan *relayerTypes.WarpBlockInfo
	sub          interfaces.Subscription

	errChan chan error

	logger logging.Logger
}

// NewSubscriber returns a subscriber
func NewSubscriber(
	logger logging.Logger,
	blockchainID ids.ID,
	wsClient ethclient.Client,
	rpcClient ethclient.Client,
) *subscriber {
	return &subscriber{
		blockchainID: blockchainID,
		wsClient:     wsClient,
		rpcClient:    rpcClient,
		logger:       logger,
		icmBlocks:    make(chan *relayerTypes.WarpBlockInfo, maxClientSubscriptionBuffer),
		headers:      make(chan *types.Header, maxClientSubscriptionBuffer),
		errChan:      make(chan error),
	}
}

// Process logs from the given block height to the latest block. Limits the
// number of blocks retrieved in a single eth_getLogs request to
// `MaxBlocksPerRequest`; if processing more than that, multiple eth_getLogs
// requests will be made.
// Writes true to the done channel when finished, or false if an error occurs
func (s *subscriber) ProcessFromHeight(height *big.Int, done chan bool) {
	defer close(done)
	s.logger.Info(
		"Processing historical logs",
		zap.String("fromBlockHeight", height.String()),
		zap.String("blockchainID", s.blockchainID.String()),
	)
	if height == nil {
		s.logger.Error("Cannot process logs from nil height")
		done <- false
		return
	}

	// Grab the latest block before filtering logs so we don't miss any before updating the db
	latestBlockHeightCtx, latestBlockHeightCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer latestBlockHeightCtxCancel()
	latestBlockHeight, err := s.rpcClient.BlockNumber(latestBlockHeightCtx)
	if err != nil {
		s.logger.Error(
			"Failed to get latest block",
			zap.String("blockchainID", s.blockchainID.String()),
			zap.Error(err),
		)
		done <- false
		return
	}

	bigLatestBlockHeight := big.NewInt(0).SetUint64(latestBlockHeight)

	//nolint:lll
	for fromBlock := big.NewInt(0).Set(height); fromBlock.Cmp(bigLatestBlockHeight) <= 0; fromBlock.Add(fromBlock, big.NewInt(MaxBlocksPerRequest)) {
		toBlock := big.NewInt(0).Add(fromBlock, big.NewInt(MaxBlocksPerRequest-1))

		// clamp to latest known block because we've already subscribed
		// to new blocks and we don't want to double-process any blocks
		// created after that subscription but before the determination
		// of this "latest"
		if toBlock.Cmp(bigLatestBlockHeight) > 0 {
			toBlock.Set(bigLatestBlockHeight)
		}

		err = s.processBlockRange(fromBlock, toBlock)
		if err != nil {
			s.logger.Error("Failed to process block range", zap.Error(err))
			done <- false
			return
		}
	}
	done <- true
}

// Process Warp messages from the block range [fromBlock, toBlock], inclusive
func (s *subscriber) processBlockRange(
	fromBlock, toBlock *big.Int,
) error {
	logs, err := s.getFilterLogsByBlockRangeRetryable(fromBlock, toBlock)
	if err != nil {
		s.logger.Error(
			"Failed to get header by number after max attempts",
			zap.String("blockchainID", s.blockchainID.String()),
			zap.Error(err),
		)
		return err
	}

	blocksWithICMMessages, err := relayerTypes.LogsToBlocks(logs)
	if err != nil {
		s.logger.Error("Failed to convert logs to blocks", zap.Error(err))
		return err
	}
	blocks, err := fillBlockRange(fromBlock.Uint64(), toBlock.Uint64(), blocksWithICMMessages)
	if err != nil {
		s.logger.Error("Failed to fill block range", zap.Error(err))
		return err
	}
	for _, block := range blocks {
		s.icmBlocks <- block
	}
	return nil
}

func (s *subscriber) getFilterLogsByBlockRangeRetryable(fromBlock, toBlock *big.Int) ([]types.Log, error) {
	var (
		err  error
		logs []types.Log
	)
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		logs, err = s.rpcClient.FilterLogs(cctx, interfaces.FilterQuery{
			Topics:    [][]common.Hash{{relayerTypes.WarpPrecompileLogFilter}},
			Addresses: []common.Address{warp.ContractAddress},
			FromBlock: fromBlock,
			ToBlock:   toBlock,
		})
		return err
	}
	err = utils.WithRetriesTimeout(s.logger, operation, utils.DefaultRPCTimeout, "get filter logs by block range")
	if err != nil {
		s.logger.Error(
			"Failed to get filter logs by block range",
			zap.String("blockchainID", s.blockchainID.String()),
			zap.Error(err),
		)
		return nil, relayerTypes.ErrFailedToProcessLogs
	}
	return logs, nil
}

// Loops forever iff maxResubscribeAttempts == 0
func (s *subscriber) Subscribe(retryTimeout time.Duration) error {
	// Unsubscribe before resubscribing
	// s.sub should only be nil on the first call to Subscribe
	if s.sub != nil {
		s.sub.Unsubscribe()
	}

	err := s.subscribe(retryTimeout)
	if err != nil {
		return errors.New("failed to subscribe to node")
	}
	return nil
}

// subscribe until it succeeds or reached timeout.
func (s *subscriber) subscribe(retryTimeout time.Duration) error {
	var sub interfaces.Subscription
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		sub, err = s.wsClient.SubscribeNewHead(cctx, s.headers)
		return err
	}
	err := utils.WithRetriesTimeout(s.logger, operation, retryTimeout, "subscribe")
	if err != nil {
		s.logger.Error(
			"Failed to subscribe to node",
			zap.String("blockchainID", s.blockchainID.String()),
		)
		return err
	}
	s.sub = sub

	go s.blocksInfoFromHeaders()
	return nil
}

// blocksInfoFromHeaders listens to the header channel and converts the headers to [relayerTypes.WarpBlockInfo]
// and writes them to the blocks channel consumed by the listener
func (s *subscriber) blocksInfoFromHeaders() {
	for header := range s.headers {
		block, err := relayerTypes.NewWarpBlockInfo(s.logger, header, s.rpcClient)
		if err != nil {
			s.logger.Error("Failed to create Warp block info", zap.Error(err))
			s.errChan <- err
			return
		}
		s.icmBlocks <- block
	}
}

func (s *subscriber) ICMBlocks() <-chan *relayerTypes.WarpBlockInfo {
	return s.icmBlocks
}

// SubscribeErr returns the error channel for the underlying subscription
func (s *subscriber) SubscribeErr() <-chan error {
	return s.sub.Err()
}

// Err returns the error channel for miscellaneous errors not recoverable from
// by resubscribing.
func (s *subscriber) Err() <-chan error {
	return s.errChan
}

func (s *subscriber) Cancel() {
	// Nothing to do here, the ethclient manages both the log and err channels
}

// Takes a block range and a slice of blocks containing ICM messages and fills in the missing blocks with empty blocks
func fillBlockRange(fromBlock, toBlock uint64, icmBlocks []*relayerTypes.WarpBlockInfo) ([]*relayerTypes.WarpBlockInfo, error) {
	var filledBlocks []*relayerTypes.WarpBlockInfo
	currentBlock := fromBlock
	for _, block := range icmBlocks {
		filledBlocks = append(filledBlocks, makeEmptyBlocks(currentBlock, block.BlockNumber-1)...)
		filledBlocks = append(filledBlocks, block)
		currentBlock = block.BlockNumber + 1
	}
	// Fill the remainder of the range with empty blocks after the last block
	// with an ICM message was appended to the slice.
	filledBlocks = append(filledBlocks, makeEmptyBlocks(currentBlock, toBlock)...)
	return filledBlocks, nil
}

func makeEmptyBlocks(fromBlock, toBlock uint64) []*relayerTypes.WarpBlockInfo {
	var emptyBlocks []*relayerTypes.WarpBlockInfo
	for i := fromBlock; i <= toBlock; i++ {
		emptyBlocks = append(emptyBlocks, &relayerTypes.WarpBlockInfo{
			BlockNumber: i,
			Messages:    []*relayerTypes.WarpMessageInfo{},
		})
	}
	return emptyBlocks
}
