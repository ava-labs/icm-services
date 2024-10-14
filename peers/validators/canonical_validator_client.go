// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_canonical_validator_client.go -package=mocks

package validators

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/awm-relayer/config"
	"github.com/ava-labs/awm-relayer/peers/utils"
	"go.uber.org/zap"
)

type CanonicalValidatorClient interface {
	validators.State

	GetBlockByHeight(context.Context, uint64) ([]byte, error)
}

// CanonicalValidatorClient wraps platformvm.Client and implements validators.State
type canonicalValidatorClient struct {
	logger  logging.Logger
	client  platformvm.Client
	options []rpc.Option
}

func NewCanonicalValidatorClient(logger logging.Logger, apiConfig *config.APIConfig) CanonicalValidatorClient {
	client := platformvm.NewClient(apiConfig.BaseURL)
	options := utils.InitializeOptions(apiConfig)
	return &canonicalValidatorClient{
		logger:  logger,
		client:  client,
		options: options,
	}
}

func (v *canonicalValidatorClient) GetMinimumHeight(ctx context.Context) (uint64, error) {
	return v.client.GetHeight(ctx, v.options...)
}

func (v *canonicalValidatorClient) GetCurrentHeight(ctx context.Context) (uint64, error) {
	return v.client.GetHeight(ctx, v.options...)
}

func (v *canonicalValidatorClient) GetBlockByHeight(ctx context.Context, height uint64) ([]byte, error) {
	return v.client.GetBlockByHeight(ctx, height, v.options...)
}

func (v *canonicalValidatorClient) GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error) {
	return v.client.ValidatedBy(ctx, blockchainID, v.options...)
}

// Gets the validator set of the given subnet at the given P-chain block height.
func (v *canonicalValidatorClient) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*validators.GetValidatorOutput, error) {
	// First, attempt to use the "getValidatorsAt" RPC method. This method may not be available on
	// all API nodes, in which case we can fall back to using "getCurrentValidators" if needed.
	res, err := v.client.GetValidatorsAt(ctx, subnetID, height, v.options...)
	if err != nil {
		v.logger.Warn(
			"P-chain RPC to getValidatorAt returned error. Falling back to getCurrentValidators",
			zap.String("subnetID", subnetID.String()),
			zap.Uint64("pChainHeight", height),
			zap.Error(err))
		return nil, err
	}
	return res, nil
}
