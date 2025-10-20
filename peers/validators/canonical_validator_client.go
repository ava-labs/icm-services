// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_canonical_validator_client.go -package=mocks

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"go.uber.org/zap"

	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/peers/utils"
)

var _ CanonicalValidatorState = &CanonicalValidatorClient{}

// CanonicalValidatorState is an interface that wraps [avalancheWarp.ValidatorState] and adds additional
// convenience methods for fetching current and proposed validator sets.
type CanonicalValidatorState interface {
	GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error)
	GetLatestHeight(ctx context.Context) (uint64, error)
	GetAllValidatorSets(ctx context.Context, pchainHeight uint64) (map[ids.ID]validators.WarpSet, error)
	GetProposedValidators(ctx context.Context, subnetID ids.ID) (validators.WarpSet, error)
}

// CanonicalValidatorClient wraps [platformvm.Client] and implements [CanonicalValidatorState]
type CanonicalValidatorClient struct {
	logger  logging.Logger
	client  *platformvm.Client
	options []rpc.Option
}

func NewCanonicalValidatorClient(logger logging.Logger, apiConfig *config.APIConfig) *CanonicalValidatorClient {
	client := platformvm.NewClient(apiConfig.BaseURL)
	options := utils.InitializeOptions(apiConfig)
	return &CanonicalValidatorClient{
		logger:  logger,
		client:  client,
		options: options,
	}
}

func (v *CanonicalValidatorClient) GetLatestHeight(ctx context.Context) (uint64, error) {
	height, err := v.client.GetHeight(ctx, v.options...)
	if err != nil {
		v.logger.Error("Failed to get latest height", zap.Error(err))
		return 0, err
	}
	return height, nil
}

func (v *CanonicalValidatorClient) GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error) {
	return v.client.ValidatedBy(ctx, blockchainID, v.options...)
}

func (v *CanonicalValidatorClient) GetProposedValidators(
	ctx context.Context,
	subnetID ids.ID,
) (validators.WarpSet, error) {
	res, err := v.client.GetValidatorsAt(ctx, subnetID, pchainapi.ProposedHeight, v.options...)
	if err != nil {
		v.logger.Debug(
			"Error fetching proposed validators",
			zap.String("subnetID", subnetID.String()),
			zap.Error(err),
		)
		return validators.WarpSet{}, err
	}
	return validators.FlattenValidatorSet(res)
}

// Gets the validator set of the given subnet at the given P-chain block height.
// Uses [platform.getValidatorsAt] with supplied height
func (v *CanonicalValidatorClient) GetAllValidatorSets(
	ctx context.Context,
	height uint64,
) (map[ids.ID]validators.WarpSet, error) {
	res, err := v.client.GetAllValidatorsAt(ctx, pchainapi.Height(height), v.options...)
	if err != nil {
		v.logger.Debug(
			"Error fetching validators at height",
			zap.Uint64("pChainHeight", height),
			zap.Error(err),
		)
		return nil, err
	}
	return res, nil
}

func NodeIDs(vdrs validators.WarpSet) set.Set[ids.NodeID] {
	nodeIDSet := set.NewSet[ids.NodeID](len(vdrs.Validators))
	for _, validator := range vdrs.Validators {
		for _, nodeID := range validator.NodeIDs {
			nodeIDSet.Add(nodeID)
		}
	}
	return nodeIDSet
}
