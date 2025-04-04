// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/errgroup"

	"go.uber.org/zap"
)

const (
	retryTimeout = 10 * time.Second
)

// Errors
var errFailedToGetAggSig = errors.New("failed to get aggregate signature from node endpoint")

// CheckpointManager stores committed heights in the database
type CheckpointManager interface {
	// Run starts a go routine that periodically stores the last committed height in the Database
	Run()
	// StageCommittedHeight queues a height to be written to the database.
	// Heights are committed in sequence, so if height is not exactly one
	// greater than the current committedHeight, it is instead cached in memory
	// to potentially be committed later.
	StageCommittedHeight(height uint64)
}

// ApplicationRelayers define a Warp message route from a specific source address on a specific source blockchain
// to a specific destination address on a specific destination blockchain. This routing information is
// encapsulated in [relayerID], which also represents the database key for an ApplicationRelayer.
type ApplicationRelayer struct {
	logger                    logging.Logger
	metrics                   *ApplicationRelayerMetrics
	network                   peers.AppRequestNetwork
	sourceBlockchain          config.SourceBlockchain
	signingSubnetID           ids.ID
	destinationClient         vms.DestinationClient
	relayerID                 database.RelayerID
	warpConfig                config.WarpConfig
	checkpointManager         CheckpointManager
	sourceWarpSignatureClient *rpc.Client // nil if configured to fetch signatures via AppRequest for the source blockchain
	signatureAggregator       *aggregator.SignatureAggregator
}

func NewApplicationRelayer(
	logger logging.Logger,
	metrics *ApplicationRelayerMetrics,
	network peers.AppRequestNetwork,
	relayerID database.RelayerID,
	destinationClient vms.DestinationClient,
	sourceBlockchain config.SourceBlockchain,
	checkpointManager CheckpointManager,
	cfg *config.Config,
	signatureAggregator *aggregator.SignatureAggregator,
) (*ApplicationRelayer, error) {
	warpConfig, err := cfg.GetWarpConfig(relayerID.DestinationBlockchainID)
	if err != nil {
		logger.Error(
			"Failed to get warp config. Relayer may not be configured to deliver to the destination chain.",
			zap.String("destinationBlockchainID", relayerID.DestinationBlockchainID.String()),
			zap.Error(err),
		)
		return nil, err
	}

	var signingSubnet ids.ID
	if sourceBlockchain.GetSubnetID() == constants.PrimaryNetworkID && !warpConfig.RequirePrimaryNetworkSigners {
		// If the message originates from the primary network, and the primary network is validated by
		// the destination subnet we can "self-sign" the message using the validators of the destination subnet.
		logger.Info(
			"Self-signing message originating from primary network",
			zap.String("destinationBlockchainID", relayerID.DestinationBlockchainID.String()),
		)
		signingSubnet = cfg.GetSubnetID(relayerID.DestinationBlockchainID)
	} else {
		// Otherwise, the source subnet signs the message.
		signingSubnet = sourceBlockchain.GetSubnetID()
	}

	checkpointManager.Run()

	var warpClient *rpc.Client
	if !sourceBlockchain.UseAppRequestNetwork() {
		// The subnet-evm Warp API client does not support query parameters or HTTP headers
		// and expects the URI to be in a specific form.
		// Instead, we invoke the Warp API directly via the RPC client.
		warpClient, err = utils.DialWithConfig(
			context.Background(),
			sourceBlockchain.WarpAPIEndpoint.BaseURL,
			sourceBlockchain.WarpAPIEndpoint.HTTPHeaders,
			sourceBlockchain.WarpAPIEndpoint.QueryParams,
		)
		if err != nil {
			logger.Error(
				"Failed to create Warp API client",
				zap.Error(err),
			)
			return nil, err
		}
	}

	ar := ApplicationRelayer{
		logger:                    logger,
		metrics:                   metrics,
		network:                   network,
		sourceBlockchain:          sourceBlockchain,
		destinationClient:         destinationClient,
		relayerID:                 relayerID,
		signingSubnetID:           signingSubnet,
		warpConfig:                warpConfig,
		checkpointManager:         checkpointManager,
		sourceWarpSignatureClient: warpClient,
		signatureAggregator:       signatureAggregator,
	}

	return &ar, nil
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
	r.logger.Verbo(
		"Processing block",
		zap.Uint64("height", height),
		zap.Stringer("relayerID", r.relayerID.ID),
		zap.Stringer("blockchainID", r.relayerID.SourceBlockchainID),
		zap.Int("numMessages", len(handlers)),
	)
	var eg errgroup.Group
	for _, handler := range handlers {
		eg.Go(func() error {
			_, err := r.ProcessMessage(handler)
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		r.logger.Error(
			"Failed to process block",
			zap.Uint64("height", height),
			zap.String("relayerID", r.relayerID.ID.String()),
			zap.Error(err),
		)
		errChan <- err
		return
	}
	r.checkpointManager.StageCommittedHeight(height)
	r.logger.Verbo(
		"Processed block",
		zap.Uint64("height", height),
		zap.String("sourceBlockchainID", r.relayerID.SourceBlockchainID.String()),
		zap.String("relayerID", r.relayerID.ID.String()),
		zap.Int("numMessages", len(handlers)),
	)
}

// Relays a message to the destination chain. Does not checkpoint the height.
// returns the transaction hash if the message is successfully relayed.
func (r *ApplicationRelayer) ProcessMessage(handler messages.MessageHandler) (common.Hash, error) {
	r.logger.Debug(
		"Relaying message",
		zap.String("sourceBlockchainID", r.sourceBlockchain.BlockchainID),
		zap.String("relayerID", r.relayerID.ID.String()),
	)
	shouldSend, err := handler.ShouldSendMessage(r.destinationClient)
	if err != nil {
		r.logger.Error(
			"Failed to check if message should be sent",
			zap.Error(err),
		)
		r.incFailedRelayMessageCount("failed to check if message should be sent")
		return common.Hash{}, err
	}
	if !shouldSend {
		r.logger.Info("Message should not be sent")
		return common.Hash{}, nil
	}
	unsignedMessage := handler.GetUnsignedMessage()

	startCreateSignedMessageTime := time.Now()
	// Query nodes on the origin chain for signatures, and construct the signed warp message.
	var signedMessage *avalancheWarp.Message

	// sourceWarpSignatureClient is nil iff the source blockchain is configured to fetch signatures via AppRequest
	if r.sourceWarpSignatureClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DefaultCreateSignedMessageTimeout)
		defer cancel()

		signedMessage, err = r.signatureAggregator.CreateSignedMessage(
			ctx,
			unsignedMessage,
			nil,
			r.signingSubnetID,
			r.warpConfig.QuorumNumerator,
		)
		r.incFetchSignatureAppRequestCount()
		if err != nil {
			r.logger.Error(
				"Failed to create signed warp message via AppRequest network",
				zap.Error(err),
			)
			r.incFailedRelayMessageCount("failed to create signed warp message via AppRequest network")
			return common.Hash{}, err
		}
	} else {
		r.incFetchSignatureRPCCount()
		signedMessage, err = r.createSignedMessage(unsignedMessage)
		if err != nil {
			r.logger.Error(
				"Failed to create signed warp message via RPC",
				zap.Error(err),
			)
			r.incFailedRelayMessageCount("failed to create signed warp message via RPC")
			return common.Hash{}, err
		}
	}

	// create signed message latency (ms)
	r.setCreateSignedMessageLatencyMS(float64(time.Since(startCreateSignedMessageTime).Milliseconds()))

	txHash, err := handler.SendMessage(signedMessage, r.destinationClient)
	if err != nil {
		r.logger.Error(
			"Failed to send warp message",
			zap.Error(err),
		)
		r.incFailedRelayMessageCount("failed to send warp message")
		return common.Hash{}, err
	}
	r.logger.Info(
		"Finished relaying message to destination chain",
		zap.String("destinationBlockchainID", r.relayerID.DestinationBlockchainID.String()),
		zap.String("txHash", txHash.Hex()),
	)
	r.incSuccessfulRelayMessageCount()

	return txHash, nil
}

func (r *ApplicationRelayer) RelayerID() database.RelayerID {
	return r.relayerID
}

// createSignedMessage fetches the signed Warp message from the source chain via RPC.
// Each VM may implement their own RPC method to construct the aggregate signature, which
// will need to be accounted for here.
func (r *ApplicationRelayer) createSignedMessage(
	unsignedMessage *avalancheWarp.UnsignedMessage,
) (*avalancheWarp.Message, error) {
	r.logger.Info("Fetching aggregate signature from the source chain validators via API")

	var (
		signedWarpMessageBytes hexutil.Bytes
		err                    error
	)
	cctx, cancel := context.WithTimeout(context.Background(), utils.DefaultCreateSignedMessageTimeout)
	defer cancel()
	operation := func() error {
		return r.sourceWarpSignatureClient.CallContext(
			cctx,
			&signedWarpMessageBytes,
			"warp_getMessageAggregateSignature",
			unsignedMessage.ID(),
			r.warpConfig.QuorumNumerator,
			r.signingSubnetID.String(),
		)
	}
	err = utils.WithRetriesTimeout(r.logger, operation, retryTimeout)
	if err != nil {
		r.logger.Error(
			"Failed to get aggregate signature from node endpoint.",
			zap.String("sourceBlockchainID", r.sourceBlockchain.GetBlockchainID().String()),
			zap.String("destinationBlockchainID", r.relayerID.DestinationBlockchainID.String()),
			zap.String("signingSubnetID", r.signingSubnetID.String()),
		)
		return nil, errFailedToGetAggSig
	}
	warpMsg, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
	if err != nil {
		r.logger.Error(
			"Failed to parse signed warp message",
			zap.Error(err),
		)
		return nil, err
	}
	return warpMsg, nil
}

//
// Metrics
//

func (r *ApplicationRelayer) incSuccessfulRelayMessageCount() {
	r.metrics.successfulRelayMessageCount.
		WithLabelValues(
			r.relayerID.DestinationBlockchainID.String(),
			r.sourceBlockchain.GetBlockchainID().String(),
			r.sourceBlockchain.GetSubnetID().String()).Inc()
}

func (r *ApplicationRelayer) incFailedRelayMessageCount(failureReason string) {
	r.metrics.failedRelayMessageCount.
		WithLabelValues(
			r.relayerID.DestinationBlockchainID.String(),
			r.sourceBlockchain.GetBlockchainID().String(),
			r.sourceBlockchain.GetSubnetID().String(),
			failureReason).Inc()
}

func (r *ApplicationRelayer) setCreateSignedMessageLatencyMS(latency float64) {
	r.metrics.createSignedMessageLatencyMS.
		WithLabelValues(
			r.relayerID.DestinationBlockchainID.String(),
			r.sourceBlockchain.GetBlockchainID().String(),
			r.sourceBlockchain.GetSubnetID().String()).Set(latency)
}

func (r *ApplicationRelayer) incFetchSignatureRPCCount() {
	r.metrics.fetchSignatureRPCCount.
		WithLabelValues(
			r.relayerID.DestinationBlockchainID.String(),
			r.sourceBlockchain.GetBlockchainID().String(),
			r.sourceBlockchain.GetSubnetID().String()).Inc()
}

func (r *ApplicationRelayer) incFetchSignatureAppRequestCount() {
	r.metrics.fetchSignatureAppRequestCount.
		WithLabelValues(
			r.relayerID.DestinationBlockchainID.String(),
			r.sourceBlockchain.GetBlockchainID().String(),
			r.sourceBlockchain.GetSubnetID().String()).Inc()
}
