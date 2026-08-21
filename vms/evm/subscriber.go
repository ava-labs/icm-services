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
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/libevm/rpc"
	"go.uber.org/zap"
)

const (
	// Max buffer size for ethereum subscription channels
	maxClientSubscriptionBuffer = 20000
	MaxBlocksPerRequest         = 200

	// strictTailBlocks is the number of most recent blocks of a catch-up range
	// that are processed block-by-block (existence check by number, logs by
	// node-reported hash) instead of via by-number range queries. Nodes behind
	// a load-balanced endpoint may briefly disagree about the newest blocks,
	// and a range query served by a lagging node silently omits blocks it does
	// not yet have. Observed skew is sub-second; 10 blocks is a generous
	// margin. Blocks older than this exist on every node, so range queries
	// remain safe and fast for deep history.
	strictTailBlocks = 10

	// defaultScanInterval bounds live-loop scan frequency and is the fallback
	// poll that advances the watermark through empty blocks. See (*Subscriber).Run.
	defaultScanInterval = time.Second

	// defaultConfirmations is how far behind the tip the live loop scans.
	// 0 processes up to the tip; increase to trade latency for reorg safety.
	defaultConfirmations = 0
)

type SubscriberRPCClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	// HeadByNumber returns the block head with its node-reported hash, or an
	// error — ethereum.NotFound when the serving node does not have the block,
	// which callers treat as retryable.
	HeadByNumber(ctx context.Context, number *big.Int) (*BlockHead, error)
	ethereum.LogFilterer
}

// RPCHeadClient augments an ethclient with verbatim-hash head fetches by
// number, satisfying SubscriberRPCClient. The hash must come from the node
// rather than be recomputed client-side, which is unreliable for chains whose
// headers carry fields this client cannot encode (e.g. SAE chains).
type RPCHeadClient struct {
	*ethclient.Client
}

func NewRPCHeadClient(client *ethclient.Client) RPCHeadClient {
	return RPCHeadClient{Client: client}
}

func (c RPCHeadClient) HeadByNumber(ctx context.Context, number *big.Int) (*BlockHead, error) {
	var head *BlockHead
	err := c.Client.Client().CallContext(ctx, &head, "eth_getBlockByNumber", hexutil.EncodeBig(number), false)
	if err == nil && head == nil {
		err = ethereum.NotFound
	}
	return head, err
}

type SubscriberWSClient interface {
	// SubscribeFilterLogs is used only as a low-latency trigger to wake the live
	// loop; log completeness is guaranteed by the scan over the HTTP rpcClient.
	ethereum.LogFilterer
}

// WSHeadClient exposes the LogFilterer methods of an ethclient.Client.
type WSHeadClient struct {
	*ethclient.Client
}

func NewWSHeadClient(client *rpc.Client) WSHeadClient {
	return WSHeadClient{Client: ethclient.NewClient(client)}
}

type Subscriber struct {
	wsClient         SubscriberWSClient
	rpcClient        SubscriberRPCClient
	blockchainID     ids.ID
	isPrimaryNetwork bool
	filter           EventFilter
	// logs receives pushed logs used only as a signal to scan sooner.
	logs      chan types.Log
	icmBlocks chan *ICMBlockInfo
	sub       ethereum.Subscription

	confirmations uint64
	scanInterval  time.Duration

	errChan chan error

	logger logging.Logger
}

// NewSubscriber returns a Subscriber
func NewSubscriber(
	logger logging.Logger,
	blockchainID ids.ID,
	isPrimaryNetwork bool,
	wsClient SubscriberWSClient,
	rpcClient SubscriberRPCClient,
	errChan chan error,
	filter EventFilter,
) *Subscriber {
	return &Subscriber{
		blockchainID:     blockchainID,
		isPrimaryNetwork: isPrimaryNetwork,
		filter:           filter,
		wsClient:         wsClient,
		rpcClient:        rpcClient,
		logger:           logger,
		icmBlocks:        make(chan *ICMBlockInfo, maxClientSubscriptionBuffer),
		logs:             make(chan types.Log, maxClientSubscriptionBuffer),
		confirmations:    defaultConfirmations,
		scanInterval:     defaultScanInterval,
		errChan:          errChan,
	}
}

// ProcessFromHeight processes logs from the starting block to the ending block,
// inclusive, marking the emitted blocks as catch-up. Writes to the error channel
// if an error occurs.
func (s *Subscriber) ProcessFromHeight(startingHeight uint64, endingHeight uint64) {
	log := s.logger.With(
		zap.Uint64("fromBlockHeight", startingHeight),
		zap.Uint64("toBlockHeight", endingHeight),
	)
	log.Info("Processing historical logs")
	if err := s.scanRange(startingHeight, endingHeight, true); err != nil {
		s.errChan <- err
		return
	}
	log.Info("Finished processing historical logs")
}

// scanRange emits one ICMBlockInfo per height in [startingHeight, endingHeight],
// including empty blocks so the checkpoint can advance contiguously. Deep
// history is fetched in eth_getLogs chunks of MaxBlocksPerRequest; the newest
// strictTailBlocks are fetched one by one.
func (s *Subscriber) scanRange(startingHeight, endingHeight uint64, isCatchup bool) error {
	if endingHeight < startingHeight {
		return nil
	}

	// Range queries are only trustworthy for blocks old enough that every node
	// behind a load-balanced endpoint has them; the newest blocks are checked
	// strictly, one by one. See strictTailBlocks.
	strictStart := startingHeight
	if endingHeight-startingHeight+1 > strictTailBlocks {
		strictStart = endingHeight - strictTailBlocks + 1
		for fromBlock := startingHeight; fromBlock < strictStart; fromBlock += MaxBlocksPerRequest {
			toBlock := min(fromBlock+MaxBlocksPerRequest-1, strictStart-1)

			if err := s.processBlockRange(fromBlock, toBlock, isCatchup); err != nil {
				return fmt.Errorf("failed to process block range: %w", err)
			}
		}
	}

	for height := strictStart; height <= endingHeight; height++ {
		if err := s.processBlockStrict(height, isCatchup); err != nil {
			return fmt.Errorf("failed to process block %d: %w", height, err)
		}
	}
	return nil
}

// processBlockStrict processes a single block with the same guarantees as the
// live path: existence is confirmed by number (a node that does not have the
// block answers ethereum.NotFound, which is retried) and logs are fetched by
// the node-reported hash. A by-number range query over these blocks could be
// served by a lagging node and silently omit them.
func (s *Subscriber) processBlockStrict(height uint64, isCatchup bool) error {
	var head *BlockHead
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		head, err = s.rpcClient.HeadByNumber(cctx, new(big.Int).SetUint64(height))
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"get head by number failed, retrying...",
			zap.Uint64("blockNumber", height),
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	// Same window as the live path: heads near the chain tip may not yet be
	// known to every node behind a load-balanced endpoint.
	if err := utils.WithRetriesTimeout(operation, notify, utils.DefaultRPCTimeout*6); err != nil {
		return fmt.Errorf("failed to get head for block %d: %w", height, err)
	}

	block, err := NewICMBlockInfo(s.logger, head, s.rpcClient, s.filter, s.isPrimaryNetwork)
	if err != nil {
		return err
	}
	block.IsCatchup = isCatchup
	s.icmBlocks <- block
	return nil
}

// Process Warp messages from the block range [fromBlock, toBlock], inclusive
func (s *Subscriber) processBlockRange(
	fromBlock, toBlock uint64,
	isCatchup bool,
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

	logIndex := 0
	for i := fromBlock; i <= toBlock; i++ {
		blockLogs := []types.Log{}
		for logIndex < len(logs) && logs[logIndex].BlockNumber == i {
			blockLogs = append(blockLogs, logs[logIndex])
			logIndex++
		}
		// Blocks with no ICM messages also need to be explicitly processed.
		s.icmBlocks <- &ICMBlockInfo{
			BlockNumber: i,
			Logs:        blockLogs,
			IsCatchup:   isCatchup,
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
			Addresses: s.filter.Addresses,
			Topics:    s.filter.Topics,
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
		sub, err = s.wsClient.SubscribeFilterLogs(cctx, ethereum.FilterQuery{
			Addresses: s.filter.Addresses,
			Topics:    s.filter.Topics,
		}, s.logs)
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

// Run drives the live path until ctx is cancelled. Triggered by a pushed log or
// by scanInterval, it scans [watermark+1, tip-confirmations], emitting every
// height (including empty blocks) and advancing the watermark. The first scan
// covers any historical gap from startingHeight, unifying catch-up and live.
func (s *Subscriber) Run(ctx context.Context, startingHeight uint64) {
	var watermark uint64
	if startingHeight > 0 {
		watermark = startingHeight - 1
	}

	ticker := time.NewTicker(s.scanInterval)
	defer ticker.Stop()

	var lastScan time.Time
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.logs:
		case <-ticker.C:
		}

		// At most one scan per scanInterval, coalescing bursts of logs.
		if wait := s.scanInterval - time.Since(lastScan); wait > 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(wait):
			}
		}
		drainLogs(s.logs)
		lastScan = time.Now()

		tip, err := s.currentBlockNumber(ctx)
		if err != nil {
			s.logger.Warn("failed to get block number, will retry", zap.Error(err))
			continue
		}
		if tip < s.confirmations {
			continue
		}
		safeTip := tip - s.confirmations
		if safeTip <= watermark {
			continue
		}

		if err := s.scanRange(watermark+1, safeTip, false); err != nil {
			s.errChan <- fmt.Errorf("failed to scan live range: %w", err)
			return
		}
		watermark = safeTip
	}
}

func (s *Subscriber) currentBlockNumber(ctx context.Context) (uint64, error) {
	cctx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	return s.rpcClient.BlockNumber(cctx)
}

// drainLogs empties the trigger channel so pending triggers coalesce into one scan.
func drainLogs(ch <-chan types.Log) {
	for {
		select {
		case <-ch:
		default:
			return
		}
	}
}

func (s *Subscriber) ICMBlocks() <-chan *ICMBlockInfo {
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
