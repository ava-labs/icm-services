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
	cctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(cfg.InitialConnectionTimeoutSeconds)*time.Second,
	)
	defer cancel()

	allValidators, err := network.GetAllValidatorSets(cctx, pchainapi.ProposedHeight)
	if err != nil {
		return fmt.Errorf("failed to get all canonical validators: %w", err)
	}

	eg, ectx := errgroup.WithContext(cctx)
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		vdrs := network.BuildCanonicalValidators(allValidators[sourceBlockchain.GetBlockchainID()])
		log := logger.With(
			zap.Stringer("subnetID", sourceBlockchain.GetSubnetID()),
			zap.Stringer("blockchainID", sourceBlockchain.GetBlockchainID()),
			zap.Uint64("connectedWeight", vdrs.ConnectedWeight),
			zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
			zap.Int("numConnectedPeers", vdrs.ConnectedNodes.Len()),
		)
		// Loop over destination blockchains here to confirm connections to a threshold of stake
		// which is determined by the Warp Quorum configs of the destination blockchains.
		maxQuorumNumerator, err := getMaxQuorumNumerator(logger, cfg, sourceBlockchain)
		if err != nil {
			return fmt.Errorf("failed to calculate max quorum numerator: %w", err)
		}
		log = log.With(zap.Uint64("maxQuorumNumerator", maxQuorumNumerator))

		eg.Go(func() error {
			log.Info("Checking sufficient stake for source blockchain")
			return checkSufficientConnectedStake(
				ectx,
				log,
				vdrs,
				maxQuorumNumerator,
			)
		})
	}
	return eg.Wait()
}

func getMaxQuorumNumerator(
	logger logging.Logger,
	cfg *config.Config,
	sourceBlockchain *config.SourceBlockchain,
) (uint64, error) {
	var maxQuorumNumerator uint64
	for _, destination := range sourceBlockchain.SupportedDestinations {
		destinationBlockchainID := destination.GetBlockchainID()
		warpConfig, err := cfg.GetWarpConfig(destinationBlockchainID)
		if err != nil {
			return 0, fmt.Errorf("failed to get warp config for destination %s: %w", destinationBlockchainID, err)
		}
		maxQuorumNumerator = max(maxQuorumNumerator, warpConfig.QuorumNumerator)
	}
	return maxQuorumNumerator, nil
}

// Returns when connected to sufficient stake for all supported destinations of the source blockchain or
// in case of errors or timeouts.
func checkSufficientConnectedStake(
	ctx context.Context,
	logger logging.Logger,
	vdrs *peers.CanonicalValidators,
	maxQuorumNumerator uint64,
) error {
	checkConns := func() error {
		// Log details of each connected validator (nodeID and weight).
		for _, nodeID := range vdrs.ConnectedNodes.List() {
			vdr, _ := vdrs.GetValidator(nodeID)
			logger.Debug(
				"Connected validator details",
				zap.Stringer("nodeID", nodeID),
				zap.Uint64("weight", vdr.Weight),
			)
		}

		if !utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(vdrs.ConnectedWeight),
			vdrs.ValidatorSet.TotalWeight,
			maxQuorumNumerator,
		) {
			logger.Info("Failed to connect to a threshold of stake, retrying...")
			return fmt.Errorf("failed to connect to sufficient stake")
		}

		logger.Info("Connected to sufficient stake")
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
