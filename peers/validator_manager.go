package peers

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers/clients"
	sharedUtils "github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
)

type ValidatorManager struct {
	metrics *AppRequestNetworkMetrics
	logger  logging.Logger

	validatorSetLock         *sync.Mutex
	validatorClient          clients.CanonicalValidatorState
	latestSyncedPChainHeight atomic.Uint64
	maxPChainLookback        int64
	manager                  snowVdrs.Manager

	epochedValidatorSetCache *cache.FIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet]
}

func NewValidatorManager(
	cfg Config,
	logger logging.Logger,
	metrics *AppRequestNetworkMetrics,
	validatorSetsCacheSize int,
	manager snowVdrs.Manager,
) *ValidatorManager {
	validatorClient := clients.NewCanonicalValidatorClient(cfg.GetPChainAPI())
	epochedValidatorSetCache := cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](validatorSetsCacheSize)
	return &ValidatorManager{
		logger:                   logger,
		validatorClient:          validatorClient,
		metrics:                  metrics,
		maxPChainLookback:        cfg.GetMaxPChainLookback(),
		epochedValidatorSetCache: epochedValidatorSetCache,
		manager:                  manager,
		validatorSetLock:         new(sync.Mutex),
	}
}

func (v *ValidatorManager) StartCacheValidatorSets(ctx context.Context) {
	// Fetch validators immediately when called, and refresh every ValidatorRefreshPeriod
	ticker := time.NewTicker(ValidatorPreFetchPeriod)
	v.cacheMostRecentValidatorSets(ctx)

	for {
		select {
		case <-ticker.C:
			v.cacheMostRecentValidatorSets(ctx)
		case <-ctx.Done():
			v.logger.Info("Stopping caching validator process...")
			return
		}
	}
}

func (v *ValidatorManager) GetLatestSyncedPChainHeight() uint64 {
	return v.latestSyncedPChainHeight.Load()
}

func (v *ValidatorManager) GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error) {
	return v.validatorClient.GetSubnetID(ctx, blockchainID)
}

func (v *ValidatorManager) GetLatestValidatorSets(ctx context.Context) (map[ids.ID]snowVdrs.WarpSet, error) {
	cctx, cancel := context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	latestPChainHeight, err := v.validatorClient.GetLatestHeight(cctx)
	if err != nil {
		v.logger.Warn("Failed to get latest P-Chain height", zap.Error(err))
		return nil, err
	}

	return v.GetAllValidatorSets(cctx, latestPChainHeight)
}

func (v *ValidatorManager) GetAllValidatorSets(
	ctx context.Context,
	pchainHeight uint64,
) (map[ids.ID]snowVdrs.WarpSet, error) {
	// If we're getting the proposed height, bypass the cache and get the latest data
	// We can't cache this call because we don't know the actual P-Chain height being returned.
	if pchainHeight == pchainapi.ProposedHeight {
		return v.validatorClient.GetAllValidatorSets(ctx, pchainHeight)
	}

	// Use FIFO cache for epoched validators (specific heights) - immutable historical data
	// FIFO cache key is pchainHeight, fetch function uses the passed height
	fetchVdrsFunc := func(height uint64) (map[ids.ID]snowVdrs.WarpSet, error) {
		latestSyncedHeight := v.latestSyncedPChainHeight.Load()
		if v.maxPChainLookback >= 0 && int64(height) < int64(latestSyncedHeight)-v.maxPChainLookback {
			return nil, fmt.Errorf("requested P-Chain height %d is beyond the max lookback of %d from latest height %d",
				height, v.maxPChainLookback, latestSyncedHeight,
			)
		}

		v.logger.Debug("Fetching all canonical validator sets at P-Chain height", zap.Uint64("pchainHeight", height))
		startPChainAPICall := time.Now()
		validatorSet, err := v.validatorClient.GetAllValidatorSets(ctx, height)
		v.metrics.pChainAPICallLatencyMS.Observe(float64(time.Since(startPChainAPICall).Milliseconds()))
		return validatorSet, err
	}

	validatorSets, err := v.epochedValidatorSetCache.Get(pchainHeight, fetchVdrsFunc)
	if err != nil {
		return nil, err
	}

	// If the fetch succeeded, the set is in the cache now so update the latest synced height if greater
	// than the current latest synced height using atomic compare-and-swap
	for {
		current := v.latestSyncedPChainHeight.Load()
		if pchainHeight <= current {
			break
		}
		if v.latestSyncedPChainHeight.CompareAndSwap(current, pchainHeight) {
			break
		}
		// CAS failed, another goroutine updated it, retry
	}

	return validatorSets, nil
}

// Update the tracked validators for a single subnet. This is used when tracking a new subnet for the first time.
func (v *ValidatorManager) UpdateTrackedValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) error {
	cctx, cancel := context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	vdrs, err := v.validatorClient.GetProposedValidators(cctx, subnetID)
	if err != nil {
		return err
	}

	return v.updatedTrackedValidators(subnetID, vdrs)
}

func (v *ValidatorManager) updatedTrackedValidators(
	subnetID ids.ID,
	vdrs snowVdrs.WarpSet,
) error {
	v.validatorSetLock.Lock()
	defer v.validatorSetLock.Unlock()

	nodeIDs := clients.NodeIDs(vdrs)

	// Remove any elements from the manager that are not in the new validator set
	currentVdrs := v.manager.GetValidatorIDs(subnetID)
	for _, nodeID := range currentVdrs {
		if !nodeIDs.Contains(nodeID) {
			v.logger.Debug("Removing validator",
				zap.Stringer("nodeID", nodeID),
				zap.Stringer("subnetID", subnetID),
			)
			weight := v.manager.GetWeight(subnetID, nodeID)
			if err := v.manager.RemoveWeight(subnetID, nodeID, weight); err != nil {
				return err
			}
		}
	}

	// Add any elements from the new validator set that are not in the manager
	for _, vdr := range vdrs.Validators {
		for _, nodeID := range vdr.NodeIDs {
			if _, ok := v.manager.GetValidator(subnetID, nodeID); !ok {
				v.logger.Debug("Adding validator",
					zap.Stringer("nodeID", nodeID),
					zap.Stringer("subnetID", subnetID),
				)
				if err := v.manager.AddStaker(
					subnetID,
					nodeID,
					vdr.PublicKey,
					ids.Empty,
					vdr.Weight,
				); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ValidatorManager) cacheMostRecentValidatorSets(ctx context.Context) {
	latestPChainHeight, err := v.validatorClient.GetLatestHeight(ctx)
	if err != nil {
		// This is not a critical error, just log and return
		v.logger.Error("Failed to get P-Chain height", zap.Error(err))
		return
	}

	currentSyncedHeight := v.latestSyncedPChainHeight.Load()
	if currentSyncedHeight == 0 {
		// Setting the current synced height to be one less than the latest P-Chain upon initialization makes it
		// such that we only fetch the validator sets at the latest P-Chain height to start.
		currentSyncedHeight = latestPChainHeight - 1
		v.latestSyncedPChainHeight.Store(currentSyncedHeight)
		v.logger.Info("Initializing P-Chain height", zap.Uint64("height", currentSyncedHeight))
	}

	for currentSyncedHeight < latestPChainHeight {
		currentSyncedHeight++
		// GetAllValidatorSets will update latestSyncedPChainHeight after successful cache
		_, err := v.GetAllValidatorSets(ctx, currentSyncedHeight)
		// If we fail to get the validator sets for this height, log and check the next height.
		if err != nil {
			v.logger.Error("Failed to get canonical validators",
				zap.Uint64("height", currentSyncedHeight),
				zap.Error(err),
			)
			continue
		}
	}
}
