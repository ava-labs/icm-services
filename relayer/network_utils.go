// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
)

const initialConnectionTimeoutSeconds = 300

// Convenience function to initialize connections and check stake for all source blockchains.
// Only returns an error if it fails to get a list of canonical validator or a valid warp config.
//
// Failing a sufficient stake check will only log an error but still return successfully
// since each attempted relay will make an attempt at reconnecting to any missing validators.
//
// Sufficient stake is determined by the Warp quora of the configured supported destinations,
// or if the subnet supports all destinations, by the quora of all configured destinations.
func InitializeConnectionsAndCheckStake(
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
) error {
	for _, sourceBlockchainConfig := range cfg.SourceBlockchains {
		network.TrackSubnet(sourceBlockchainConfig.GetSubnetID())
	}
	ctx, cancel := context.WithTimeout(context.Background(), initialConnectionTimeoutSeconds*time.Second)
	defer cancel()
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		if sourceBlockchain.GetSubnetID() == constants.PrimaryNetworkID {
			if err := connectToPrimaryNetworkPeers(ctx, logger, network, cfg, sourceBlockchain); err != nil {
				return fmt.Errorf(
					"failed to connect to primary network peers: %w",
					err,
				)
			}
		} else {
			if err := connectToNonPrimaryNetworkPeers(ctx, logger, network, cfg, sourceBlockchain); err != nil {
				return fmt.Errorf(
					"failed to connect to non-primary network peers: %w",
					err,
				)
			}
		}
	}
	return nil
}

// Connect to the validators of the source blockchain. For each destination blockchain,
// verify that we have connected to a threshold of stake.
func connectToNonPrimaryNetworkPeers(
	ctx context.Context,
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
	sourceBlockchain *config.SourceBlockchain,
) error {
	subnetID := sourceBlockchain.GetSubnetID()
	network.TrackSubnet(subnetID)
	for _, destination := range sourceBlockchain.SupportedDestinations {
		blockchainID := destination.GetBlockchainID()
		for {
			connectedValidators, err := network.GetConnectedCanonicalValidators(subnetID)
			if err != nil {
				logger.Error(
					"Failed to connect to canonical validators",
					zap.String("subnetID", subnetID.String()),
					zap.Error(err),
				)
				return err
			}
			if ok, warpConfig, err := checkForSufficientConnectedStake(logger, cfg, connectedValidators, blockchainID); ok && err == nil {
				break
			} else {
				logger.Warn(
					"Failed to connect to a threshold of stake, retrying...",
					zap.String("destinationBlockchainID", blockchainID.String()),
					zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
					zap.Uint64("totalValidatorWeight", connectedValidators.ValidatorSet.TotalWeight),
					zap.Any("WarpConfig", warpConfig),
				)
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					time.Sleep(5 * time.Second) // Retry after a short delay
				}
			}
		}
	}
	return nil
}

// Connect to the validators of the destination blockchains. Verify that we have connected
// to a threshold of stake for each blockchain.
func connectToPrimaryNetworkPeers(
	ctx context.Context,
	logger logging.Logger,
	network peers.AppRequestNetwork,
	cfg *config.Config,
	sourceBlockchain *config.SourceBlockchain,
) error {
	for _, destination := range sourceBlockchain.SupportedDestinations {
		blockchainID := destination.GetBlockchainID()
		subnetID := cfg.GetSubnetID(blockchainID)
		network.TrackSubnet(subnetID)

		for {
			connectedValidators, err := network.GetConnectedCanonicalValidators(subnetID)
			if err != nil {
				logger.Error(
					"Failed to connect to canonical validators",
					zap.String("subnetID", subnetID.String()),
					zap.Error(err),
				)
				return err
			}
			if ok, warpConfig, err := checkForSufficientConnectedStake(logger, cfg, connectedValidators, blockchainID); ok && err == nil {
				break
			} else {
				logger.Warn(
					"Failed to connect to a threshold of stake, retrying...",
					zap.String("destinationBlockchainID", blockchainID.String()),
					zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
					zap.Uint64("totalValidatorWeight", connectedValidators.ValidatorSet.TotalWeight),
					zap.Any("WarpConfig", warpConfig),
				)
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					time.Sleep(1 * time.Second) // Retry after a short delay
				}
			}
		}
	}
	return nil
}

// Fetch the warp config from the destination chain config and check if the connected stake exceeds the threshold
func checkForSufficientConnectedStake(
	logger logging.Logger,
	cfg *config.Config,
	connectedValidators *peers.ConnectedCanonicalValidators,
	destinationBlockchainID ids.ID,
) (bool, *config.WarpConfig, error) {
	warpConfig, err := cfg.GetWarpConfig(destinationBlockchainID)
	if err != nil {
		logger.Error(
			"Failed to get warp config from chain config",
			zap.String("destinationBlockchainID", destinationBlockchainID.String()),
			zap.Error(err),
		)
		return false, nil, err
	}
	return utils.CheckStakeWeightExceedsThreshold(
		big.NewInt(0).SetUint64(connectedValidators.ConnectedWeight),
		connectedValidators.ValidatorSet.TotalWeight,
		warpConfig.QuorumNumerator,
	), &warpConfig, nil
}
