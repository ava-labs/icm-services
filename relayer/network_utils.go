// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const retryPeriodSeconds = 5

// Convenience function to initialize connections and check stake for all source blockchains.
// This function blocks until it successfully connects to sufficient stake for all source blockchains
// or returns an error if unable to fetch warpConfigs or to connect to sufficient stake before timeout.
//
// Sufficient stake is determined by the Warp quora of the configured supported destinations,
// or if the subnet supports all destinations, by the quora of all configured destinations.
func InitializeConnectionsAndCheckStake(
	ctx context.Context,
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
) error {
	for _, subnet := range cfg.GetTrackedSubnets().List() {
		network.TrackSubnet(ctx, subnet)
	}
	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(cfg.InitialConnectionTimeoutSeconds)*time.Second,
	)
	defer cancel()

	eg, cctx := errgroup.WithContext(ctx)
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		eg.Go(func() error {
			logger.Info("Checking sufficient stake for source blockchain",
				zap.Stringer("subnetID", sourceBlockchain.GetSubnetID()),
				zap.String("blockchainID", sourceBlockchain.GetBlockchainID().String()),
			)
			return checkSufficientConnectedStake(cctx, logger, network, cfg, sourceBlockchain)
		})
	}
	return eg.Wait()
}

// Returns when connected to sufficient stake for all supported destinations of the source blockchain or
// in case of errors or timeouts.
func checkSufficientConnectedStake(
	ctx context.Context,
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
	sourceBlockchain *config.SourceBlockchain,
) error {
	subnetID := sourceBlockchain.GetSubnetID()
	// Loop over destination blockchains here to confirm connections to a threshold of stake
	// which is determined by the Warp Quorum configs of the destination blockchains.
	var maxQuorumNumerator uint64

	for _, destination := range sourceBlockchain.SupportedDestinations {
		destinationBlockchainID := destination.GetBlockchainID()
		warpConfig, err := cfg.GetWarpConfig(destinationBlockchainID)
		logger.Debug("Fetched warp config for destination",
			zap.Stringer("destinationBlockchainID", destinationBlockchainID),
			zap.Any("warpConfig", warpConfig),
			zap.Error(err),
		)
		if err != nil {
			logger.Error(
				"Failed to get warp config from chain config",
				zap.Stringer("destinationBlockchainID", destinationBlockchainID),
				zap.Error(err),
			)
			return err
		}
		maxQuorumNumerator = max(maxQuorumNumerator, warpConfig.QuorumNumerator)
	}

	checkConns := func() error {
		vdrs, err := network.GetCanonicalValidators(ctx, subnetID, false, uint64(pchainapi.ProposedHeight))
		if err != nil {
			logger.Error(
				"Failed to retrieve currently connected validators",
				zap.Stringer("subnetID", subnetID),
				zap.Error(err),
			)
			return err
		}

		// Log details of each connected validator (nodeID and weight).
		for _, nodeID := range vdrs.ConnectedNodes.List() {
			vdr, _ := vdrs.GetValidator(nodeID)
			logger.Debug(
				"Connected validator details",
				zap.Stringer("subnetID", subnetID),
				zap.String("nodeID", nodeID.String()),
				zap.Uint64("weight", vdr.Weight),
			)
		}

		if !utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(vdrs.ConnectedWeight),
			vdrs.ValidatorSet.TotalWeight,
			maxQuorumNumerator,
		) {
			logger.Warn(
				"Failed to connect to a threshold of stake, retrying...",
				zap.Stringer("subnetID", subnetID),
				zap.Uint64("quorumNumerator", maxQuorumNumerator),
				zap.Uint64("connectedWeight", vdrs.ConnectedWeight),
				zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
				zap.Int("numConnectedPeers", vdrs.ConnectedNodes.Len()),
			)
			return fmt.Errorf("failed to connect to sufficient stake")
		}

		logger.Info(
			"Connected to sufficient stake",
			zap.Stringer("subnetID", subnetID),
			zap.Uint64("quorumNumerator", maxQuorumNumerator),
			zap.Uint64("connectedWeight", vdrs.ConnectedWeight),
			zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
			zap.Int("numConnectedPeers", vdrs.ConnectedNodes.Len()),
		)
		return nil
	}

	ticker := time.Tick(retryPeriodSeconds * time.Second)
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("failed to connect to sufficient stake: %w", ctx.Err())
		case <-ticker:
			if err := checkConns(); err == nil {
				return nil
			}
		}
	}
}
