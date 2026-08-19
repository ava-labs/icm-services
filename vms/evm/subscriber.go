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
	SubscribeNewHeads(ctx context.Context, ch chan<- *BlockHead) (ethereum.Subscription, error)
	// Used to fetch logs for live blocks; see blocksInfoFromHeaders for why
	// these fetches must use the WS connection.
	ethereum.LogFilterer
}

// WSHeadClient adapts a raw rpc.Client to SubscriberWSClient. It decodes
// newHeads notifications into BlockHead, keeping the node-reported block hash,
// which cannot be recomputed client-side for chains whose headers carry fields
// unknown to this client (e.g. SAE chains).
//
// The embedded ethclient.Client wraps the same rpc.Client, so log fetches are
// issued over the same connection as the newHeads notification that triggered
// them, and therefore reach the same node.
type WSHeadClient struct {
	client *rpc.Client
	*ethclient.Client
}

func NewWSHeadClient(client *rpc.Client) WSHeadClient {
	return WSHeadClient{
		client: client,
		Client: ethclient.NewClient(client),
	}
}

func (w WSHeadClient) SubscribeNewHeads(
	ctx context.Context,
	ch chan<- *BlockHead,
) (ethereum.Subscription, error) {
	return w.client.EthSubscribe(ctx, ch, "newHeads")
}

type Subscriber struct {
	wsClient         SubscriberWSClient
	rpcClient        SubscriberRPCClient
	blockchainID     ids.ID
	isPrimaryNetwork bool
	filter           EventFilter
	headers          chan *BlockHead
	icmBlocks        chan *ICMBlockInfo
	sub              ethereum.Subscription

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
	subscriber := &Subscriber{
		blockchainID:     blockchainID,
		isPrimaryNetwork: isPrimaryNetwork,
		filter:           filter,
		wsClient:         wsClient,
		rpcClient:        rpcClient,
		logger:           logger,
		icmBlocks:        make(chan *ICMBlockInfo, maxClientSubscriptionBuffer),
		headers:          make(chan *BlockHead, maxClientSubscriptionBuffer),
		errChan:          errChan,
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

	if endingHeight < startingHeight {
		log.Info("Finished processing historical logs")
		return
	}

	// Range queries are only trustworthy for blocks old enough that every node
	// behind a load-balanced endpoint has them; the newest blocks are checked
	// strictly, one by one. See strictTailBlocks.
	strictStart := startingHeight
	if endingHeight-startingHeight+1 > strictTailBlocks {
		strictStart = endingHeight - strictTailBlocks + 1
		for fromBlock := startingHeight; fromBlock < strictStart; fromBlock += MaxBlocksPerRequest {
			toBlock := min(fromBlock+MaxBlocksPerRequest-1, strictStart-1)

			err := s.processBlockRange(fromBlock, toBlock)
			if err != nil {
				s.errChan <- fmt.Errorf("failed to process block range: %w", err)
				return
			}
		}
	}

	for height := strictStart; height <= endingHeight; height++ {
		if err := s.processBlockStrict(height); err != nil {
			s.errChan <- fmt.Errorf("failed to process block %d: %w", height, err)
			return
		}
	}
	log.Info("Finished processing historical logs")
}

// processBlockStrict processes a single block with the same guarantees as the
// live path: existence is confirmed by number (a node that does not have the
// block answers ethereum.NotFound, which is retried) and logs are fetched by
// the node-reported hash. A by-number range query over these blocks could be
// served by a lagging node and silently omit them.
func (s *Subscriber) processBlockStrict(height uint64) error {
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
	block.IsCatchup = true
	s.icmBlocks <- block
	return nil
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
			IsCatchup:   true,
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
		sub, err = s.wsClient.SubscribeNewHeads(cctx, s.headers)
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

// blocksInfoFromHeaders listens to the header channel and converts the headers to [ICMBlockInfo]
// and writes them to the blocks channel consumed by the listener
func (s *Subscriber) blocksInfoFromHeaders() {
	for head := range s.headers {
		// Fetch logs over the WS connection rather than the HTTP client: the
		// node that emitted this head is guaranteed to have the block's
		// receipts on disk, whereas an HTTP request may be routed to a
		// different node behind a load balancer that has not yet executed the
		// block. On SAE chains such a node returns empty logs without an
		// error, which would cause this block's messages to be silently
		// skipped.
		block, err := NewICMBlockInfo(s.logger, head, s.wsClient, s.filter, s.isPrimaryNetwork)
		if err != nil {
			s.errChan <- fmt.Errorf("creating warp block info: %w", err)
			return
		}
		s.icmBlocks <- block
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
