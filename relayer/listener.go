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
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

const (
	retrySubscribeTimeout = 10 * time.Second
	// TODO attempt to resubscribe in perpetuity once we are able to process missed blocks and
	// refresh the chain config on reconnect.
	retryResubscribeTimeout = 10 * time.Second
	// idleCheckpointInterval is how often the listener checkpoints up to the chain head when the
	// subscription has not reported any block with matching logs in the meantime.
	idleCheckpointInterval = time.Hour
)

// Listener handles all messages sent from a given source chain
type Listener struct {
	Subscriber         *evm.Subscriber
	currentRequestID   uint32
	logger             logging.Logger
	sourceBlockchainID ids.ID
	healthStatus       *atomic.Bool
	ethClient          *ethclient.Client
	messageCoordinator *MessageCoordinator
	maxConcurrentMsg   uint64
	errChan            chan error
	protocol           config.Protocol
}

// RunListener creates a Listener instance and the ApplicationRelayers for a subnet.
// The Listener listens for warp messages on that subnet, and the ApplicationRelayers handle delivery to the destination
func RunListener(
	ctx context.Context,
	logger logging.Logger,
	protocol config.Protocol,
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
		zap.String("protocolAddress", protocol.Address.String()),
	)
	// Create the Listener
	listener, err := newListener(
		ctx,
		logger,
		protocol,
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
	protocol config.Protocol,
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

	// The message protocol decides which source chain logs carry its messages, so the filter the
	// subscriber uses comes from the protocol's message handler factory.
	messageHandlerFactory, ok := messageCoordinator.messageHandlerFactories[blockchainID][protocol.Address]
	if !ok {
		return nil, fmt.Errorf(
			"no message handler configured for protocol address %s on blockchain %s",
			protocol.Address,
			blockchainID,
		)
	}
	eventFilter := messageHandlerFactory.EventFilter()
	if eventFilter.IsEmpty() {
		return nil, fmt.Errorf(
			"message protocol %s does not send messages on a source chain",
			protocol.Type,
		)
	}

	wsRPCClient, err := utils.DialWithConfig(
		ctx,
		sourceBlockchain.WSEndpoint.BaseURL,
		sourceBlockchain.WSEndpoint.HTTPHeaders,
		sourceBlockchain.WSEndpoint.QueryParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to node via WS: %w", err)
	}

	errChan := make(chan error, maxConcurrentMsg)
	sub := evm.NewSubscriber(
		logger,
		blockchainID,
		ethclient.NewClient(wsRPCClient),
		evm.NewRPCHeaderClient(ethRPCClient),
		errChan,
		eventFilter,
		startingHeight,
	)

	logger.Info("Creating relayer")
	lstnr := Listener{
		Subscriber:         sub,
		currentRequestID:   rand.Uint32(), // Initialize to a random value to mitigate requestID collision
		logger:             logger,
		sourceBlockchainID: blockchainID,
		errChan:            errChan,
		healthStatus:       relayerHealth,
		ethClient:          ethRPCClient,
		messageCoordinator: messageCoordinator,
		maxConcurrentMsg:   maxConcurrentMsg,
		protocol:           protocol,
	}

	// Open the subscription, which also dispatches catch-up of any missed blocks.
	err = lstnr.Subscriber.Subscribe(retrySubscribeTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to node: %w", err)
	}

	return &lstnr, nil
}

// Listens to the Subscriber blocks channel to process them.
// On subscriber error, attempts to reconnect and errors if unable.
// Exits if context is cancelled by another goroutine.
func (lstnr *Listener) processLogs(ctx context.Context) error {
	idleCheckpointTicker := time.NewTicker(idleCheckpointInterval)
	defer idleCheckpointTicker.Stop()
	for {
		select {
		case err := <-lstnr.errChan:
			lstnr.healthStatus.Store(false)
			lstnr.logger.Error("Listener received error", zap.Error(err))
			return fmt.Errorf("listener received error: %w", err)
		case icmBlockInfo := <-lstnr.Subscriber.ICMBlocks():
			go lstnr.messageCoordinator.ProcessBlock(
				icmBlockInfo,
				lstnr.sourceBlockchainID,
				lstnr.protocol.Address,
				lstnr.errChan,
			)
		case <-idleCheckpointTicker.C:
			// Not fatal: the next tick, or the next block with matching logs, will try again.
			if err := lstnr.Subscriber.ProcessIdleBlocks(); err != nil {
				lstnr.logger.Warn("Failed to checkpoint idle blocks", zap.Error(err))
			}
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

// reconnectToSubscriber reopens the subscription, which also dispatches catch-up of the blocks
// missed while it was broken.
func (lstnr *Listener) reconnectToSubscriber() error {
	err := lstnr.Subscriber.Subscribe(retryResubscribeTimeout)
	if err != nil {
		return fmt.Errorf("failed to resubscribe to node: %w", err)
	}

	// Success
	lstnr.healthStatus.Store(true)
	return nil
}
