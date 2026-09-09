// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
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
	sourceBlockchainID ids.ID
	healthStatus       *atomic.Bool
	ethClient          *ethclient.Client
	messageCoordinator *MessageCoordinator
	maxConcurrentMsg   uint64
	errChan            chan error
	// highestDispatchedBlock is the highest source chain block that has been
	// dispatched for processing, either by catch-up or from the subscription.
	// Every block up to it is accounted for, so when the subscription reports
	// a block above it, the blocks in between contain no matching logs.
	highestDispatchedBlock uint64
	protocol               config.Protocol
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
		sourceBlockchain.GetSubnetID() == constants.PrimaryNetworkID,
		ethclient.NewClient(wsRPCClient),
		evm.NewRPCHeaderClient(ethRPCClient),
		errChan,
		eventFilter,
	)

	logger.Info("Creating relayer")
	lstnr := Listener{
		Subscriber:             sub,
		currentRequestID:       rand.Uint32(), // Initialize to a random value to mitigate requestID collision
		logger:                 logger,
		sourceBlockchainID:     blockchainID,
		errChan:                errChan,
		healthStatus:           relayerHealth,
		ethClient:              ethRPCClient,
		messageCoordinator:     messageCoordinator,
		maxConcurrentMsg:       maxConcurrentMsg,
		highestDispatchedBlock: startingHeight - 1,
		protocol:               protocol,
	}

	// Open the subscription, which also dispatches catch-up of any missed blocks.
	err = lstnr.subscribe(retrySubscribeTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to node: %w", err)
	}

	return &lstnr, nil
}

// subscribe opens the subscription to the source chain's logs and dispatches catch-up of the blocks
// the subscription will not deliver: those up to and including the subscribed node's head that have
// not been dispatched yet, i.e. the blocks missed while the relayer was down or the previous
// subscription was broken. The subscription is opened before catch-up is bounded, so no block can
// fall between the two.
func (lstnr *Listener) subscribe(retryTimeout time.Duration) error {
	head, err := lstnr.Subscriber.Subscribe(retryTimeout)
	if err != nil {
		return err
	}
	// Run catch-up in a separate goroutine so that the main processing loop can start processing
	// new blocks as soon as possible. ProcessFromHeight returns immediately if there is nothing to
	// catch up on, e.g. if the newly subscribed node is behind the previously subscribed one.
	go lstnr.Subscriber.ProcessFromHeight(lstnr.highestDispatchedBlock+1, head)
	lstnr.highestDispatchedBlock = max(lstnr.highestDispatchedBlock, head)
	return nil
}

// Listens to the Subscriber blocks channel to process them.
// On subscriber error, attempts to reconnect and errors if unable.
// Exits if context is cancelled by another goroutine.
func (lstnr *Listener) processLogs(ctx context.Context) error {
	for {
		select {
		case err := <-lstnr.errChan:
			lstnr.healthStatus.Store(false)
			lstnr.logger.Error("Listener received error", zap.Error(err))
			return fmt.Errorf("listener received error: %w", err)
		case icmBlockInfo := <-lstnr.Subscriber.ICMBlocks():
			if !icmBlockInfo.IsCatchup {
				// The subscription only reports blocks that contain matching logs, and reports
				// them in order. The blocks between the highest dispatched block and this one
				// were therefore not reported because they contain no matching logs, so they
				// are folded into this block's range for the checkpoint manager to account
				// for. A block at or below the highest dispatched block was already covered
				// by catch-up and is processed again as is; relaying is idempotent.
				if icmBlockInfo.FromBlock > lstnr.highestDispatchedBlock+1 {
					icmBlockInfo.FromBlock = lstnr.highestDispatchedBlock + 1
				}
				lstnr.highestDispatchedBlock = max(lstnr.highestDispatchedBlock, icmBlockInfo.ToBlock)
			}

			go lstnr.messageCoordinator.ProcessBlock(
				icmBlockInfo,
				lstnr.sourceBlockchainID,
				lstnr.protocol.Address,
				lstnr.errChan,
			)
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
	err := lstnr.subscribe(retryResubscribeTimeout)
	if err != nil {
		return fmt.Errorf("failed to resubscribe to node: %w", err)
	}

	// Success
	lstnr.healthStatus.Store(true)
	return nil
}
