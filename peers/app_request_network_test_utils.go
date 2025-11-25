package peers

import (
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers/clients"
)

// NewAppRequestNetworkForTesting creates an AppRequestNetwork instance for testing
// with injected mock dependencies
func NewAppRequestNetworkForTesting(
	mockNetwork network.Network,
	handler *RelayerExternalHandler,
	logger logging.Logger,
	metrics *AppRequestNetworkMetrics,
	mockValidatorClient clients.CanonicalValidatorState,
	manager snowVdrs.Manager,
) *AppRequestNetwork {
	// Initialize cache for testing
	epochedValidatorSetCache := cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](10)

	validatorManager := &ValidatorManager{
		logger:                   logger,
		validatorClient:          mockValidatorClient,
		metrics:                  metrics,
		maxPChainLookback:        -1, // No lookback limit for tests
		epochedValidatorSetCache: epochedValidatorSetCache,
		manager:                  manager,
		validatorSetLock:         new(sync.Mutex),
	}

	return &AppRequestNetwork{
		network:            mockNetwork,
		handler:            handler,
		logger:             logger,
		metrics:            metrics,
		trackedSubnets:     set.NewSet[ids.ID](0),
		lruSubnets:         linked.NewHashmap[ids.ID, interface{}](),
		trackedSubnetsLock: new(sync.RWMutex),
		validatorManager:   validatorManager,
	}
}
