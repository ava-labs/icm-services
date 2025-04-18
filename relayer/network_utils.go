// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
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
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
) error {
	for _, subnet := range cfg.GetTrackedSubnets().List() {
		network.TrackSubnet(subnet)
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.InitialConnectionTimeoutSeconds)*time.Second,
	)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		eg.Go(func() error {
			return checkSufficientConnectedStake(ctx, logger, network, cfg, sourceBlockchain)
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
		connectedValidators, err := network.GetConnectedCanonicalValidators(subnetID)
		if err != nil {
			logger.Error(
				"Failed to retrieve currently connected validators",
				zap.Stringer("subnetID", subnetID),
				zap.Error(err),
			)
			return err
		}
		if utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(connectedValidators.ConnectedWeight),
			connectedValidators.ValidatorSet.TotalWeight,
			maxQuorumNumerator,
		) {
			logger.Info(
				"Connected to sufficient stake",
				zap.Stringer("subnetID", subnetID),
				zap.Uint64("quorumNumerator", maxQuorumNumerator),
				zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
				zap.Uint64("totalValidatorWeight", connectedValidators.ValidatorSet.TotalWeight),
				zap.Int("numConnectedPeers", network.NumConnectedPeers()),
			)
			return nil
		}
		logger.Warn(
			"Failed to connect to a threshold of stake, retrying...",
			zap.Stringer("subnetID", subnetID),
			zap.Uint64("quorumNumerator", maxQuorumNumerator),
			zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
			zap.Uint64("totalValidatorWeight", connectedValidators.ValidatorSet.TotalWeight),
			zap.Int("numConnectedPeers", network.NumConnectedPeers()),
		)
		return fmt.Errorf("failed to connect to sufficient stake")
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
