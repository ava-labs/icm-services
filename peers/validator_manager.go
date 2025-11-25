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

	canonicalValidatorSetCache *cache.TTLCache[ids.ID, snowVdrs.WarpSet]
	epochedValidatorSetCache   *cache.FIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet]
}

func NewValidatorManager(
	cfg Config,
	logger logging.Logger,
	metrics *AppRequestNetworkMetrics,
	validatorSetsCacheSize int,
	manager snowVdrs.Manager,
) *ValidatorManager {
	validatorClient := clients.NewCanonicalValidatorClient(cfg.GetPChainAPI())
	canonicalValidatorSetCache := cache.NewTTLCache[ids.ID, snowVdrs.WarpSet](canonicalValidatorSetCacheTTL)
	epochedValidatorSetCache := cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](validatorSetsCacheSize)
	return &ValidatorManager{
		logger:                     logger,
		validatorClient:            validatorClient,
		metrics:                    metrics,
		maxPChainLookback:          cfg.GetMaxPChainLookback(),
		canonicalValidatorSetCache: canonicalValidatorSetCache,
		epochedValidatorSetCache:   epochedValidatorSetCache,
		manager:                    manager,
		validatorSetLock:           new(sync.Mutex),
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
	// ProposedHeight is not supported because it's not cacheable and returns an unknown height.
	// Callers should use GetLatestValidatorSets() instead, which fetches the latest height
	// and then gets validators for that specific height.
	if pchainHeight == pchainapi.ProposedHeight {
		v.logger.Warn("ProposedHeight passed to GetAllValidatorSets - Calling GetLatestValidatorSets() instead.")
		return v.GetLatestValidatorSets(ctx)
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

// GetCanonicalValidators returns the validator information in canonical ordering for the given subnet
// at the specified P-Chain height, as well as the total weight of the validators that this network is connected to
// The caller determines the appropriate P-Chain height (ProposedHeight for current, specific height for epoched)
func (v *ValidatorManager) GetValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
	skipCache bool,
	pchainHeight uint64,
) (*snowVdrs.WarpSet, error) {
	v.logger.Debug("Getting validator set at P-Chain height",
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("pchainHeight", pchainHeight),
		zap.Bool("isProposedHeight", pchainHeight == pchainapi.ProposedHeight),
	)

	var validatorSet snowVdrs.WarpSet
	var err error

	if pchainHeight == pchainapi.ProposedHeight {
		// Get the subnet's current canonical validator set
		fetchVdrsFunc := func(subnetID ids.ID) (snowVdrs.WarpSet, error) {
			startPChainAPICall := time.Now()
			validatorSet, err := v.validatorClient.GetProposedValidators(ctx, subnetID)
			v.metrics.pChainAPICallLatencyMS.Observe(float64(time.Since(startPChainAPICall).Milliseconds()))
			if err != nil {
				return snowVdrs.WarpSet{}, err
			}
			return validatorSet, nil
		}
		validatorSet, err = v.canonicalValidatorSetCache.Get(subnetID, fetchVdrsFunc, skipCache)
	} else {
		validatorSet, err = v.getValidatorSetGranite(ctx, subnetID, pchainHeight)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get validator set at P-Chain height %d: %w", pchainHeight, err)
	}

	return &validatorSet, nil
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

func (v *ValidatorManager) getValidatorSetGranite(
	ctx context.Context,
	subnetID ids.ID,
	pchainHeight uint64,
) (snowVdrs.WarpSet, error) {
	allValidators, err := v.GetAllValidatorSets(ctx, pchainHeight)
	if err != nil {
		return snowVdrs.WarpSet{}, fmt.Errorf("failed to get all validators at P-Chain height %d: %w", pchainHeight, err)
	}

	validatorSet, ok := allValidators[subnetID]
	if !ok {
		return snowVdrs.WarpSet{}, fmt.Errorf("no validators for subnet %s at P-Chain height %d", subnetID, pchainHeight)
	}
	return validatorSet, nil
}
