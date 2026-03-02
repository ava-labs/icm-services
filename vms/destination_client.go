// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_destination_client.go -package=mocks

package vms

import (
	"context"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// DestinationClient is the interface for the destination chain client. Methods that interact with
// the destination chain should generally be implemented in a thread-safe way, as they will be called
// concurrently by the application relayers.
type DestinationClient interface {
	// SendTx constructs the transaction from warp primitives, and sends to the configured destination chain endpoint.
	// Returns the hash of the sent transaction.
	// TODO: Make generic for any VM.
	SendTx(
		signedMessage *warp.Message,
		deliverers set.Set[common.Address],
		toAddress string,
		gasLimit uint64,
		callData []byte,
	) (*types.Receipt, error)

	// Client returns the underlying client for the destination chain
	Client() evm.Client

	// SenderAddresses returns the addresses of the relayer on the destination chain
	SenderAddresses() []common.Address

	// DestinationBlockchainID returns the ID of the destination chain
	DestinationBlockchainID() ids.ID

	// BlockGasLimit returns destination blockchain block gas limit
	BlockGasLimit() uint64

	// GetRPCEndpointURL returns the RPC endpoint URL for this destination blockchain
	GetRPCEndpointURL() string

	// GetPChainHeightForDestination determines the appropriate P-Chain height for validator set selection.
	// The epoch is cached per destination blockchain to avoid per-message fetches.
	GetPChainHeightForDestination(
		ctx context.Context,
	) (uint64, error)
}

// CreateDestinationClients creates destination clients for all subnets configured as destinations
func CreateDestinationClients(
	logger logging.Logger,
	relayerConfig *config.Config,
) (map[ids.ID]DestinationClient, error) {
	// Fetch epoch duration once since it's global across all blockchains
	var epochDuration time.Duration
	infoAPIConfig := relayerConfig.GetInfoAPI()
	infoAPI, err := clients.NewInfoAPI(infoAPIConfig)
	if err != nil {
		logger.Error("Failed to create info API for epoch duration", zap.Error(err))
		return nil, fmt.Errorf("failed to create info API: %w", err)
	}
	upgradeConfig, err := infoAPI.Upgrades(context.Background())
	if err != nil {
		logger.Error("Failed to get upgrade config for epoch duration", zap.Error(err))
		return nil, fmt.Errorf("failed to get upgrade config: %w", err)
	}
	epochDuration = upgradeConfig.GraniteEpochDuration
	logger.Info("Fetched Granite epoch duration",
		zap.Duration("epochDuration", epochDuration),
	)

	destinationClients := make(map[ids.ID]DestinationClient)
	for _, subnetInfo := range relayerConfig.DestinationBlockchains {
		log := logger.With(
			zap.String("blockchainID", subnetInfo.BlockchainID),
		)
		blockchainID, err := ids.FromString(subnetInfo.BlockchainID)
		if err != nil {
			log.Error("Failed to decode base-58 encoded source chain ID", zap.Error(err))
			return nil, err
		}
		if _, ok := destinationClients[blockchainID]; ok {
			log.Info("Destination client already found for blockchainID. Continuing")
			continue
		}

		destinationClient, err := evm.NewDestinationClient(log, subnetInfo, epochDuration)
		if err != nil {
			log.Error("Could not create destination client", zap.Error(err))
			return nil, err
		}

		destinationClients[blockchainID] = destinationClient
	}
	return destinationClients, nil
}
