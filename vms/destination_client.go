// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_destination_client.go -package=mocks

package vms

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// DestinationClient is the interface for the destination chain client. Methods that interact with
// the destination chain should generally be implemented in a thread safe way, as they will be called
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
	Client() interface{}

	// SenderAddresses returns the addresses of the relayer on the destination chain
	SenderAddresses() []common.Address

	// DestinationBlockchainID returns the ID of the destination chain
	DestinationBlockchainID() ids.ID

	// BlockGasLimit returns destination blockchain block gas limit
	BlockGasLimit() uint64
}

func NewDestinationClient(
	logger logging.Logger, subnetInfo *config.DestinationBlockchain,
) (DestinationClient, error) {
	switch config.ParseVM(subnetInfo.VM) {
	case config.EVM:
		return evm.NewDestinationClient(logger, subnetInfo)
	default:
		return nil, fmt.Errorf("invalid vm")
	}
}

// CreateDestinationClients creates destination clients for all subnets configured as destinations
func CreateDestinationClients(
	logger logging.Logger,
	relayerConfig config.Config,
) (map[ids.ID]DestinationClient, error) {
	destinationClients := make(map[ids.ID]DestinationClient)
	for _, subnetInfo := range relayerConfig.DestinationBlockchains {
		blockchainID, err := ids.FromString(subnetInfo.BlockchainID)
		if err != nil {
			logger.Error(
				"Failed to decode base-58 encoded source chain ID",
				zap.String("blockchainID", subnetInfo.BlockchainID),
				zap.Error(err),
			)
			return nil, err
		}
		if _, ok := destinationClients[blockchainID]; ok {
			logger.Info(
				"Destination client already found for blockchainID. Continuing",
				zap.String("blockchainID", blockchainID.String()),
			)
			continue
		}

		destinationClient, err := NewDestinationClient(logger, subnetInfo)
		if err != nil {
			logger.Error(
				"Could not create destination client",
				zap.String("blockchainID", blockchainID.String()),
				zap.Error(err),
			)
			return nil, err
		}

		destinationClients[blockchainID] = destinationClient
	}
	return destinationClients, nil
}
