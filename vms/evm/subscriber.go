// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"cmp"
	"context"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
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

// BlockHeader is the subset of a block header the relayer uses, decoded leniently
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
type BlockHeader struct {
	Hash   common.Hash  `json:"hash"`
	Number *hexutil.Big `json:"number"`
}

type SubscriberRPCClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	// BlockHeaderByNumber returns the block header with its node-reported hash, or an
	// error — ethereum.NotFound when the serving node does not have the block,
	// which callers treat as retryable.
	BlockHeaderByNumber(ctx context.Context, number *big.Int) (*BlockHeader, error)
	ethereum.LogFilterer
}

// RPCHeaderClient augments an ethclient with verbatim-hash header fetches by
// number, satisfying SubscriberRPCClient. The hash must come from the node
// rather than be recomputed client-side, which is unreliable for chains whose
// headers carry fields this client cannot encode (e.g. SAE chains).
type RPCHeaderClient struct {
	*ethclient.Client
}

func NewRPCHeaderClient(client *ethclient.Client) RPCHeaderClient {
	return RPCHeaderClient{Client: client}
}

func (c RPCHeaderClient) BlockHeaderByNumber(ctx context.Context, number *big.Int) (*BlockHeader, error) {
	var header *BlockHeader
	err := c.Client.Client().CallContext(ctx, &header, "eth_getBlockByNumber", hexutil.EncodeBig(number), false)
	if err == nil && header == nil {
		err = ethereum.NotFound
	}
	return header, err
}

// SubscriberWSClient is the client for the WS connection that delivers the
// log subscription. All of its methods must be served by the node the
// subscription is open with, i.e. over that same connection: the subscriber
// relies on BlockNumber to bound which blocks the subscription will deliver,
// and on FilterLogs reaching a node that has the receipts of every block up to
// that bound (see filterLogsByBlockHash). An *ethclient.Client over the WS
// connection satisfies this.
type SubscriberWSClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	ethereum.LogFilterer
}

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
	Logs []types.Log
}

type Subscriber struct {
	wsClient     SubscriberWSClient
	rpcClient    SubscriberRPCClient
	blockchainID ids.ID
	filter       EventFilter
	logs         chan types.Log
	icmBlocks    chan *ICMBlockInfo
	sub          ethereum.Subscription

	// highestDispatchedBlock is the highest source chain block that has been
	// dispatched for processing, either by catch-up or from the subscription.
	// Every block up to it is accounted for, so when the subscription reports
	// a block above it, the blocks in between contain no matching logs and
	// are folded into that block's range. It is updated both by Subscribe and
	// by the goroutine that consumes the subscription.
	highestDispatchedBlock uint64
	// dispatchedSinceIdleCheck records whether the subscription has reported a
	// new block since the last call to ProcessIdleBlocks. Guarded by
	// highestDispatchedLock.
	dispatchedSinceIdleCheck bool
	highestDispatchedLock    sync.Mutex

	errChan chan error

	logger logging.Logger
}

// NewSubscriber returns a Subscriber that accounts for every block from
// [startingHeight] onward.
func NewSubscriber(
	logger logging.Logger,
	blockchainID ids.ID,
	wsClient SubscriberWSClient,
	rpcClient SubscriberRPCClient,
	errChan chan error,
	filter EventFilter,
	startingHeight uint64,
) *Subscriber {
	highestDispatchedBlock := uint64(0)
	if startingHeight > 0 {
		highestDispatchedBlock = startingHeight - 1
	}
	
	subscriber := &Subscriber{
		blockchainID:           blockchainID,
		filter:                 filter,
		wsClient:               wsClient,
		rpcClient:              rpcClient,
		logger:                 logger,
		icmBlocks:              make(chan *ICMBlockInfo, maxClientSubscriptionBuffer),
		logs:                   make(chan types.Log, maxClientSubscriptionBuffer),
		highestDispatchedBlock: highestDispatchedBlock,
		errChan:                errChan,
	}
	go subscriber.blocksInfoFromLogs()
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
	var header *BlockHeader
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		header, err = s.rpcClient.BlockHeaderByNumber(cctx, new(big.Int).SetUint64(height))
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"get block header by number failed, retrying...",
			zap.Uint64("blockNumber", height),
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	// Same window as the live path: headers near the chain tip may not yet be
	// known to every node behind a load-balanced endpoint.
	if err := utils.WithRetriesTimeout(operation, notify, utils.DefaultRPCTimeout*6); err != nil {
		return fmt.Errorf("failed to get header for block %d: %w", height, err)
	}

	logs, err := s.filterLogsByBlockHash(header.Hash)
	if err != nil {
		return err
	}
	s.icmBlocks <- &ICMBlockInfo{
		FromBlock: height,
		ToBlock:   height,
		Logs:      logs,
	}
	return nil
}

// Process Warp messages from the block range [fromBlock, toBlock], inclusive.
// Each block that contains matching logs is reported as its own [ICMBlockInfo],
// covering the empty blocks before it as well. Empty blocks after the last such
// block are reported as one trailing range, so the whole of [fromBlock, toBlock]
// is accounted for.
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
		return fmt.Errorf("failed to get logs for block range [%d, %d] after max attempts: %w", fromBlock, toBlock, err)
	}
	// eth_getLogs returns logs in block order; sort defensively so that the
	// ranges below are never inverted.
	slices.SortStableFunc(logs, func(a, b types.Log) int {
		return cmp.Compare(a.BlockNumber, b.BlockNumber)
	})

	nextBlock := fromBlock
	for start := 0; start < len(logs); {
		blockNumber := logs[start].BlockNumber
		end := start
		for end < len(logs) && logs[end].BlockNumber == blockNumber {
			end++
		}
		s.icmBlocks <- &ICMBlockInfo{
			FromBlock: nextBlock,
			ToBlock:   blockNumber,
			Logs:      logs[start:end],
		}
		nextBlock = blockNumber + 1
		start = end
	}
	if nextBlock <= toBlock {
		s.icmBlocks <- &ICMBlockInfo{
			FromBlock: nextBlock,
			ToBlock:   toBlock,
		}
	}
	return nil
}

// filterLogsByBlockHash fetches the logs of the block with hash [blockHash]
// that match the event filter.
//
// Logs are fetched over the WS connection rather than the HTTP client: the
// subscribed node is guaranteed to have the block's receipts on disk, both for
// blocks it notified about and for catch-up blocks, which never go beyond its
// chain head. An HTTP request may instead be routed to a different node behind
// a load balancer that has not yet executed the block. On SAE chains such a
// node returns empty logs without an error, which would cause the block's
// messages to be silently skipped.
func (s *Subscriber) filterLogsByBlockHash(blockHash common.Hash) ([]types.Log, error) {
	var logs []types.Log
	// Query by hash: a node that doesn't know the block errors ("unknown
	// block") and is retried below, whereas a by-number query would return
	// empty logs with no error and the block would be silently skipped.
	operation := func() (err error) {
		// Fresh context per attempt so retries aren't killed by an
		// already-expired deadline.
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		logs, err = s.wsClient.FilterLogs(cctx, ethereum.FilterQuery{
			Addresses: s.filter.Addresses,
			Topics:    s.filter.Topics,
			BlockHash: &blockHash,
		})
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
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

// Subscribe subscribes to the source chain logs matching the event filter,
// replacing the current subscription if there is one, and dispatches catch-up
// of the blocks the subscription will not deliver: those up to and including
// the subscribed node's chain head that have not been dispatched yet, i.e. the
// blocks missed while the relayer was down or the previous subscription was
// broken. The subscription is opened before catch-up is bounded, so no block
// can fall between the two.
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

	// The head must be read from the subscribed node (over the WS connection)
	// after the subscription is open. A node that is behind the subscribed
	// node could report a height that the subscription will never deliver
	// logs for, leaving a gap between catch-up and the subscription.
	return s.catchUpToHead()
}

// ProcessIdleBlocks dispatches catch-up from the highest dispatched block to
// the subscribed node's current chain head, unless the subscription has
// reported a new block since the previous call. It is meant to be called
// periodically so that the checkpointed height keeps advancing on chains that
// go long periods without producing logs matching the filter; the subscription
// alone only reports blocks that do.
//
// The range is caught up rather than assumed empty: a notification the node
// has already sent may not have been consumed yet when the head is read, and
// catch-up will find and relay its logs instead of checkpointing past them.
func (s *Subscriber) ProcessIdleBlocks() error {
	s.highestDispatchedLock.Lock()
	dispatched := s.dispatchedSinceIdleCheck
	s.dispatchedSinceIdleCheck = false
	s.highestDispatchedLock.Unlock()
	if dispatched {
		return nil
	}
	return s.catchUpToHead()
}

// catchUpToHead reads the subscribed node's chain head and dispatches catch-up
// of the blocks between the highest dispatched block and it.
func (s *Subscriber) catchUpToHead() error {
	head, err := s.headBlockNumber()
	if err != nil {
		return fmt.Errorf("failed to get chain head of subscribed node: %w", err)
	}

	s.highestDispatchedLock.Lock()
	catchupStart := s.highestDispatchedBlock + 1
	s.highestDispatchedBlock = max(s.highestDispatchedBlock, head)
	s.highestDispatchedLock.Unlock()

	// Run catch-up in a separate goroutine so that new blocks can be processed
	// as soon as possible. ProcessFromHeight returns immediately if there is
	// nothing to catch up on, e.g. if the newly subscribed node is behind the
	// previously subscribed one.
	go s.ProcessFromHeight(catchupStart, head)
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

// headBlockNumber returns the current chain head of the subscribed node.
func (s *Subscriber) headBlockNumber() (uint64, error) {
	var head uint64
	operation := func() (err error) {
		cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()
		head, err = s.wsClient.BlockNumber(cctx)
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"get block number failed, retrying...",
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	if err := utils.WithRetriesTimeout(operation, notify, utils.DefaultRPCTimeout); err != nil {
		return 0, fmt.Errorf("failed to get block number: %w", err)
	}
	return head, nil
}

// blocksInfoFromLogs listens to the log channel of the subscription, converts
// the logs to one [ICMBlockInfo] per block and writes them to the blocks
// channel consumed by the listener.
//
// A block's matching logs are delivered one notification at a time with no
// indication of which is the last, so the block's complete set of matching
// logs is fetched when its first notification arrives and the block's remaining
// notifications are ignored. Notifications arrive in block order, so a block's
// notifications are contiguous.
func (s *Subscriber) blocksInfoFromLogs() {
	var lastBlockHash common.Hash
	for log := range s.logs {
		if log.Removed {
			// Removed logs are only sent when a block is reorged out, which
			// does not happen to accepted blocks on Avalanche chains.
			s.logger.Warn(
				"Ignoring removed log",
				zap.Uint64("blockNumber", log.BlockNumber),
				zap.Stringer("blockHash", log.BlockHash),
				zap.Stringer("txHash", log.TxHash),
			)
			continue
		}
		if log.BlockHash == lastBlockHash {
			continue
		}

		logs, err := s.filterLogsByBlockHash(log.BlockHash)
		if err != nil {
			s.errChan <- fmt.Errorf("getting ICM block logs: %w", err)
			return
		}
		s.icmBlocks <- &ICMBlockInfo{
			FromBlock: s.dispatchLiveBlock(log.BlockNumber),
			ToBlock:   log.BlockNumber,
			Logs:      logs,
		}
		lastBlockHash = log.BlockHash
	}
}

// dispatchLiveBlock records that the subscription reported [blockNumber] and
// returns the first block of the range it should be reported as. The
// subscription only reports blocks that contain matching logs, and reports
// them in order, so the blocks between the highest dispatched block and this
// one contain no matching logs and are folded into its range for the
// checkpoint manager to account for. A block at or below the highest
// dispatched block was already covered by catch-up and is reported as is; it
// is processed again, which is safe because relaying is idempotent.
func (s *Subscriber) dispatchLiveBlock(blockNumber uint64) uint64 {
	s.highestDispatchedLock.Lock()
	defer s.highestDispatchedLock.Unlock()

	if blockNumber <= s.highestDispatchedBlock {
		// Only expected for blocks accepted between opening the subscription
		// and reading the chain head, which catch-up also covers, or after
		// resubscribing to a node that is behind the previously subscribed one.
		s.logger.Warn(
			"Subscription reported an already dispatched block",
			zap.Uint64("blockNumber", blockNumber),
			zap.Uint64("highestDispatchedBlock", s.highestDispatchedBlock),
		)
		return blockNumber
	}
	fromBlock := s.highestDispatchedBlock + 1
	s.highestDispatchedBlock = blockNumber
	s.dispatchedSinceIdleCheck = true
	return fromBlock
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
