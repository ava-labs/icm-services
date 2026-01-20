// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/subnet-evm/ethclient"
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
	Subscriber                   *evm.Subscriber
	currentRequestID             uint32
	logger                       logging.Logger
	sourceBlockchain             config.SourceBlockchain
	healthStatus                 *atomic.Bool
	ethClient                    ethclient.Client
	messageCoordinator           *MessageCoordinator
	maxConcurrentMsg             uint64
	errChan                      chan error
	lastSubscriberBlockProcessed uint64
}

// RunListener creates a Listener instance and the ApplicationRelayers for a subnet.
// The Listener listens for warp messages on that subnet, and the ApplicationRelayers handle delivery to the destination
func RunListener(
	ctx context.Context,
	logger logging.Logger,
	sourceBlockchain config.SourceBlockchain,
	ethRPCClient ethclient.Client,
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
	ethRPCClient ethclient.Client,
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
	errChan := make(chan error, maxConcurrentMsg)
	sub := evm.NewSubscriber(logger, blockchainID, ethWSClient, ethRPCClient, errChan)

	logger.Info("Creating relayer")
	lstnr := Listener{
		Subscriber:                   sub,
		currentRequestID:             rand.Uint32(), // Initialize to a random value to mitigate requestID collision
		logger:                       logger,
		sourceBlockchain:             sourceBlockchain,
		errChan:                      errChan,
		healthStatus:                 relayerHealth,
		ethClient:                    ethRPCClient,
		messageCoordinator:           messageCoordinator,
		maxConcurrentMsg:             maxConcurrentMsg,
		lastSubscriberBlockProcessed: startingHeight - 1,
	}

	// Open the subscription. We must do this before processing any missed messages, otherwise we may
	// miss an incoming message in between fetching the latest block and subscribing.
	err = lstnr.Subscriber.Subscribe(retrySubscribeTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to node: %w", err)
	}

	return &lstnr, nil
}

// Listens to the Subscriber logs channel to process them.
// On subscriber error, attempts to reconnect and errors if unable.
// Exits if context is cancelled by another goroutine.
func (lstnr *Listener) processLogs(ctx context.Context) error {
	// Error channel for application relayer errors
	needsCatchup := true
	for {
		select {
		case err := <-lstnr.errChan:
			lstnr.healthStatus.Store(false)
			lstnr.logger.Error("Listener received error", zap.Error(err))
			return fmt.Errorf("listener received error: %w", err)
		case icmBlockInfo := <-lstnr.Subscriber.ICMBlocks():
			// Catchup should run on startup, and after any reconnects. It will wait for the first block
			// received from the subscriber, so that it has an accurate bound on which blocks to process.
			if needsCatchup && !icmBlockInfo.IsCatchup {
				needsCatchup = false
				go lstnr.Subscriber.ProcessFromHeight(
					lstnr.lastSubscriberBlockProcessed+1,
					icmBlockInfo.BlockNumber-1,
				)
			}

			if !icmBlockInfo.IsCatchup && icmBlockInfo.BlockNumber > lstnr.lastSubscriberBlockProcessed {
				lstnr.lastSubscriberBlockProcessed = icmBlockInfo.BlockNumber
			}

			go lstnr.messageCoordinator.ProcessBlock(
				icmBlockInfo,
				lstnr.sourceBlockchain.GetBlockchainID(),
				lstnr.errChan,
			)
		case subError := <-lstnr.Subscriber.SubscribeErr():
			needsCatchup = true
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
