// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	relayerTypes "github.com/ava-labs/icm-services/types"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"go.uber.org/zap"
)

const (
	// Max buffer size for ethereum subscription channels
	maxClientSubscriptionBuffer = 20000
	MaxBlocksPerRequest         = 200
)

type Subscriber struct {
	wsClient     ethclient.Client
	rpcClient    ethclient.Client
	blockchainID ids.ID
	headers      chan *types.Header
	icmBlocks    chan *relayerTypes.WarpBlockInfo
	sub          ethereum.Subscription

	errChan chan error

	logger logging.Logger
}

// NewSubscriber returns a Subscriber
func NewSubscriber(
	logger logging.Logger,
	blockchainID ids.ID,
	wsClient ethclient.Client,
	rpcClient ethclient.Client,
	errChan chan error,
) *Subscriber {
	subscriber := &Subscriber{
		blockchainID: blockchainID,
		wsClient:     wsClient,
		rpcClient:    rpcClient,
		logger:       logger,
		icmBlocks:    make(chan *relayerTypes.WarpBlockInfo, maxClientSubscriptionBuffer),
		headers:      make(chan *types.Header, maxClientSubscriptionBuffer),
		errChan:      errChan,
	}
	go subscriber.blocksInfoFromHeaders()
	return subscriber
}

// Process logs from the starting block to the ending block, inclusive. Limits the
// number of blocks retrieved in a single eth_getLogs request to
// `MaxBlocksPerRequest`; if processing more than that, multiple eth_getLogs
// requests will be made.
// Writes to the error channel if an error occurs
func (s *Subscriber) ProcessFromHeight(startingHeight uint64, endingHeight uint64) {
	log := s.logger.With(
		zap.Uint64("fromBlockHeight", startingHeight),
		zap.Uint64("toBlockHeight", endingHeight),
	)
	log.Info("Processing historical logs")

	for fromBlock := startingHeight; fromBlock <= endingHeight; fromBlock += MaxBlocksPerRequest {
		toBlock := min(fromBlock+MaxBlocksPerRequest-1, endingHeight)

		err := s.processBlockRange(fromBlock, toBlock)
		if err != nil {
			s.errChan <- fmt.Errorf("failed to process block range: %w", err)
			return
		}
	}
	log.Info("Finished processing historical logs")
}

// Process Warp messages from the block range [fromBlock, toBlock], inclusive
func (s *Subscriber) processBlockRange(
	fromBlock, toBlock uint64,
) error {
	s.logger.Info(
		"Processing block range",
		zap.Uint64("fromBlockHeight", fromBlock),
		zap.Uint64("toBlockHeight", toBlock),
	)
	logs, err := s.getFilterLogsByBlockRangeRetryable(fromBlock, toBlock)
	if err != nil {
		return fmt.Errorf("failed to get header by number after max attempts: %w", err)
	}

	blocksWithICMMessages, err := relayerTypes.LogsToBlocks(logs)
	if err != nil {
		s.logger.Error("Failed to convert logs to blocks", zap.Error(err))
		return err
	}
	for i := fromBlock; i <= toBlock; i++ {
		if block, ok := blocksWithICMMessages[i]; ok {
			s.icmBlocks <- block
		} else {
			// Blocks with no ICM messages also need to be explicitly processed.
			s.icmBlocks <- &relayerTypes.WarpBlockInfo{
				BlockNumber: i,
				Messages:    []*relayerTypes.WarpMessageInfo{},
				IsCatchup:   true,
			}
		}
	}
	return nil
}

func (s *Subscriber) getFilterLogsByBlockRangeRetryable(fromBlock, toBlock uint64) ([]types.Log, error) {
	var logs []types.Log
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		logs, err = s.rpcClient.FilterLogs(cctx, ethereum.FilterQuery{
			Topics:    [][]common.Hash{{relayerTypes.WarpPrecompileLogFilter}},
			Addresses: []common.Address{warp.ContractAddress},
			FromBlock: new(big.Int).SetUint64(fromBlock),
			ToBlock:   new(big.Int).SetUint64(toBlock),
		})
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"get filter logs by block range failed, retrying...",
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	err := utils.WithRetriesTimeout(operation, notify, utils.DefaultRPCTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to get filter logs by block range: %w", err)
	}
	return logs, nil
}

// Loops forever iff maxResubscribeAttempts == 0
func (s *Subscriber) Subscribe(retryTimeout time.Duration) error {
	// Unsubscribe before resubscribing
	// s.sub should only be nil on the first call to Subscribe
	if s.sub != nil {
		s.sub.Unsubscribe()
	}

	err := s.subscribe(retryTimeout)
	if err != nil {
		return fmt.Errorf("failed to subscribe to node: %w", err)
	}
	return nil
}

// subscribe until it succeeds or reached timeout.
func (s *Subscriber) subscribe(retryTimeout time.Duration) error {
	var sub ethereum.Subscription
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		sub, err = s.wsClient.SubscribeNewHead(cctx, s.headers)
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"subscribe failed, retrying...",
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	err := utils.WithRetriesTimeout(operation, notify, retryTimeout)
	if err != nil {
		return fmt.Errorf("failed to subscribe to node: %w", err)
	}
	s.sub = sub

	return nil
}

// blocksInfoFromHeaders listens to the header channel and converts the headers to [relayerTypes.WarpBlockInfo]
// and writes them to the blocks channel consumed by the listener
func (s *Subscriber) blocksInfoFromHeaders() {
	for header := range s.headers {
		block, err := relayerTypes.NewWarpBlockInfo(s.logger, header, s.rpcClient)
		if err != nil {
			s.errChan <- fmt.Errorf("creating warp block info: %w", err)
			return
		}
		s.icmBlocks <- block
	}
}

func (s *Subscriber) ICMBlocks() <-chan *relayerTypes.WarpBlockInfo {
	return s.icmBlocks
}

// SubscribeErr returns the error channel for the underlying subscription
func (s *Subscriber) SubscribeErr() <-chan error {
	return s.sub.Err()
}

// Err returns the error channel for miscellaneous errors not recoverable from
// by resubscribing.
func (s *Subscriber) Err() <-chan error {
	return s.errChan
}
