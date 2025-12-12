// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

const (
	retrySubscribeTimeout = 10 * time.Second
	// TODO attempt to resubscribe in perpetuity once we are able to process missed blocks and
	// refresh the chain config on reconnect.
	retryResubscribeTimeout = 10 * time.Second
)

// Listener handles all messages sent from a given source chain
type Listener struct {
	Subscriber         *evm.Subscriber
	currentRequestID   uint32
	logger             logging.Logger
	sourceBlockchain   config.SourceBlockchain
	catchUpResultChan  chan bool
	healthStatus       *atomic.Bool
	ethClient          *ethclient.Client
	messageCoordinator *MessageCoordinator
	maxConcurrentMsg   uint64
}

// RunListener creates a Listener instance and the ApplicationRelayers for a subnet.
// The Listener listens for warp messages on that subnet, and the ApplicationRelayers handle delivery to the destination
func RunListener(
	ctx context.Context,
	logger logging.Logger,
	sourceBlockchain config.SourceBlockchain,
	ethRPCClient *ethclient.Client,
	relayerHealth *atomic.Bool,
	startingHeight uint64,
	messageCoordinator *MessageCoordinator,
	maxConcurrentMsg uint64,
) error {
	logger = logger.With(
		zap.Stringer("subnetID", sourceBlockchain.GetSubnetID()),
		zap.String("subnetIDHex", sourceBlockchain.GetSubnetID().Hex()),
		zap.Stringer("blockchainID", sourceBlockchain.GetBlockchainID()),
		zap.String("blockchainIDHex", sourceBlockchain.GetBlockchainID().Hex()),
	)
	// Create the Listener
	listener, err := newListener(
		ctx,
		logger,
		sourceBlockchain,
		ethRPCClient,
		relayerHealth,
		startingHeight,
		messageCoordinator,
		maxConcurrentMsg,
	)
	if err != nil {
		return fmt.Errorf("failed to create listener instance: %w", err)
	}

	logger.Info("Listener initialized. Listening for messages to relay.")

	// Wait for logs from the subscribed node
	// Will only return on error or context cancellation
	return listener.processLogs(ctx)
}

func newListener(
	ctx context.Context,
	logger logging.Logger,
	sourceBlockchain config.SourceBlockchain,
	ethRPCClient *ethclient.Client,
	relayerHealth *atomic.Bool,
	startingHeight uint64,
	messageCoordinator *MessageCoordinator,
	maxConcurrentMsg uint64,
) (*Listener, error) {
	blockchainID, err := ids.FromString(sourceBlockchain.BlockchainID)
	if err != nil {
		return nil, fmt.Errorf("invalid blockchainID provided to subscriber: %w", err)
	}

	ethWSClient, err := utils.NewEthClientWithConfig(
		ctx,
		sourceBlockchain.WSEndpoint.BaseURL,
		sourceBlockchain.WSEndpoint.HTTPHeaders,
		sourceBlockchain.WSEndpoint.QueryParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to node via WS: %w", err)
	}
	sub := evm.NewSubscriber(logger, blockchainID, ethWSClient, ethRPCClient)

	// Marks when the listener has finished the catch-up process on startup.
	// Until that time, we do not know the order in which messages are processed,
	// since the catch-up process occurs concurrently with normal message processing
	// via the subscriber's Subscribe method. As a result, we cannot safely write the
	// latest processed block to the database without risking missing a block in a fault
	// scenario.
	catchUpResultChan := make(chan bool, 1)

	logger.Info("Creating relayer")
	lstnr := Listener{
		Subscriber:         sub,
		currentRequestID:   rand.Uint32(), // Initialize to a random value to mitigate requestID collision
		logger:             logger,
		sourceBlockchain:   sourceBlockchain,
		catchUpResultChan:  catchUpResultChan,
		healthStatus:       relayerHealth,
		ethClient:          ethRPCClient,
		messageCoordinator: messageCoordinator,
		maxConcurrentMsg:   maxConcurrentMsg,
	}

	// Open the subscription. We must do this before processing any missed messages, otherwise we may
	// miss an incoming message in between fetching the latest block and subscribing.
	err = lstnr.Subscriber.Subscribe(retrySubscribeTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to node: %w", err)
	}

	// Process historical blocks in a separate goroutine so that the main processing loop can
	// start processing new blocks as soon as possible. Otherwise, it's possible for
	// ProcessFromHeight to overload the message queue and cause a deadlock.
	go sub.ProcessFromHeight(big.NewInt(0).SetUint64(startingHeight), lstnr.catchUpResultChan)

	return &lstnr, nil
}

// Listens to the Subscriber logs channel to process them.
// On subscriber error, attempts to reconnect and errors if unable.
// Exits if context is cancelled by another goroutine.
func (lstnr *Listener) processLogs(ctx context.Context) error {
	// Error channel for application relayer errors
	errChan := make(chan error, lstnr.maxConcurrentMsg)
	for {
		select {
		case err := <-errChan:
			lstnr.healthStatus.Store(false)
			lstnr.logger.Error("Received error from application relayer", zap.Error(err))
		case catchUpResult, ok := <-lstnr.catchUpResultChan:
			// As soon as we've received anything on the channel, there are no more values expected.
			// The expected case is that the channel is closed by the subscriber after writing a value to it,
			// but we also defensively handle an unexpected close.
			lstnr.catchUpResultChan = nil

			// Mark the relayer as unhealthy if the catch-up process fails or if the catch-up channel is unexpectedly closed.
			if !ok {
				lstnr.healthStatus.Store(false)
				lstnr.logger.Error("Catch-up channel unexpectedly closed. Exiting listener goroutine.")
				return fmt.Errorf("catch-up channel unexpectedly closed")
			}
			if !catchUpResult {
				lstnr.healthStatus.Store(false)
				lstnr.logger.Error("Failed to catch up on historical blocks. Exiting listener goroutine.")
				return fmt.Errorf("failed to catch up on historical blocks")
			}
		case icmBlockInfo := <-lstnr.Subscriber.ICMBlocks():
			go lstnr.messageCoordinator.ProcessBlock(
				icmBlockInfo,
				lstnr.sourceBlockchain.GetBlockchainID(),
				errChan,
			)
		case err := <-lstnr.Subscriber.Err():
			lstnr.healthStatus.Store(false)
			lstnr.logger.Error("Error processing logs. Relayer goroutine exiting", zap.Error(err))
			return fmt.Errorf("error processing logs: %w", err)
		case subError := <-lstnr.Subscriber.SubscribeErr():
			lstnr.logger.Info("Received error from subscribed node", zap.Error(subError))
			subError = lstnr.reconnectToSubscriber()
			if subError != nil {
				lstnr.healthStatus.Store(false)
				lstnr.logger.Error("Relayer goroutine exiting.", zap.Error(subError))
				return fmt.Errorf("listener goroutine exiting: %w", subError)
			}
		case <-ctx.Done():
			lstnr.healthStatus.Store(false)
			lstnr.logger.Info("Exiting listener because context cancelled")
			return nil
		}
	}
}

func (lstnr *Listener) reconnectToSubscriber() error {
	// Attempt to reconnect the subscription
	err := lstnr.Subscriber.Subscribe(retryResubscribeTimeout)
	if err != nil {
		return fmt.Errorf("failed to resubscribe to node: %w", err)
	}

	// Success
	lstnr.healthStatus.Store(true)
	return nil
}
