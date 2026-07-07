// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/relayer/checkpoint"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	retryTimeout  = 10 * time.Second
	maxRetryCount = 5
)

// ApplicationRelayers define a Warp message route from a specific source address on a specific source blockchain
// to a specific destination address on a specific destination blockchain. This routing information is
// encapsulated in [relayerID], which also represents the database key for an ApplicationRelayer.
type ApplicationRelayer struct {
	logger                  logging.Logger
	metrics                 messages.Metrics
	network                 *peers.AppRequestNetwork
	signingSubnetID         ids.ID
	destinationClient       vms.DestinationClient
	relayerID               database.RelayerID
	warpConfig              config.WarpConfig
	checkpointManager       *checkpoint.CheckpointManager
	signatureAggregator     *aggregator.SignatureAggregator
	processMessageSemaphore chan struct{}
}

func NewApplicationRelayer(
	logger logging.Logger,
	metrics *ApplicationRelayerMetrics,
	network *peers.AppRequestNetwork,
	relayerID database.RelayerID,
	destinationClient vms.DestinationClient,
	sourceBlockchain *config.SourceBlockchain,
	checkpointManager *checkpoint.CheckpointManager,
	cfg *config.Config,
	signatureAggregator *aggregator.SignatureAggregator,
	processMessageSemaphore chan struct{},
) (*ApplicationRelayer, error) {
	warpConfig, err := cfg.GetWarpConfig(relayerID.DestinationBlockchainID)
	if err != nil {
		logger.Error(
			"Failed to get warp config. Relayer may not be configured to deliver to the destination chain.",
			zap.Error(err),
		)
		return nil, err
	}

	var signingSubnet ids.ID
	if sourceBlockchain.GetSubnetID() == constants.PrimaryNetworkID && !warpConfig.RequirePrimaryNetworkSigners {
		// If the message originates from the primary network, and the primary network is validated by
		// the destination subnet we can "self-sign" the message using the validators of the destination subnet.
		logger.Info("Self-signing message originating from primary network")
		signingSubnet = cfg.GetSubnetID(relayerID.DestinationBlockchainID)
	} else {
		// Otherwise, the source subnet signs the message.
		signingSubnet = sourceBlockchain.GetSubnetID()
	}

	checkpointManager.Run()

	return &ApplicationRelayer{
		logger:                  logger.With(zap.Stringer("signingSubnetID", signingSubnet)),
		metrics:                 metrics.forRelayer(relayerID),
		network:                 network,
		destinationClient:       destinationClient,
		relayerID:               relayerID,
		signingSubnetID:         signingSubnet,
		warpConfig:              warpConfig,
		checkpointManager:       checkpointManager,
		signatureAggregator:     signatureAggregator,
		processMessageSemaphore: processMessageSemaphore,
	}, nil
}

// Process [msgs] at height [height] by relaying each message to the destination chain.
// Checkpoints the height with the checkpoint manager when all messages are relayed.
// ProcessHeight is expected to be called for every block greater than or equal to the
// [startingHeight] provided in the constructor.
func (r *ApplicationRelayer) ProcessHeight(
	height uint64,
	handlers []messages.MessageHandler,
	errChan chan error,
) {
	logger := r.logger.With(
		zap.Uint64("height", height),
		zap.Int("numMessages", len(handlers)),
	)
	logger.Verbo("Processing block")

	var eg errgroup.Group
	for _, handler := range handlers {
		// Acquire the semaphore to limit the number of messages being processed concurrently globally.
		r.processMessageSemaphore <- struct{}{}

		eg.Go(func() error {
			defer func() {
				<-r.processMessageSemaphore
			}()
			_, err := r.ProcessMessage(handler)
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		logger.Error("Failed to process block", zap.Error(err))
		errChan <- err
		return
	}
	r.checkpointManager.StageCommittedHeight(height)
	logger.Verbo("Processed block")
}

func (r *ApplicationRelayer) ProcessMessage(handler messages.MessageHandler) (common.Hash, error) {
	var err error
	// Retry processing the message if it fails to account for cases where the signature is successfully aggregated
	// but the message fails to verify on the destination chain due to validator churn
	// No delays are implemented between retries since the failure scenario here involves timing differences
	// and the signature aggregator will not re-query the individual validators from which it has already
	// acquired the signatures.
	for i := 0; i < maxRetryCount; i++ {
		var txHash common.Hash
		startProcessMessageTime := time.Now()
		// Skip the cache if this is not the first attempt
		txHash, err = handler.ProcessMessage()
		if err == nil {
			return txHash, nil
		}
		r.logger.Warn(
			"failed to process message",
			zap.Int("attempt", i+1),
			zap.Int64("latencyMS", time.Since(startProcessMessageTime).Milliseconds()),
			zap.Error(err),
		)
	}
	r.logger.Error("failed to process message after max retries", zap.Error(err))
	return common.Hash{}, err
}
