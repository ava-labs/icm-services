// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"sync"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/icm-services/cache"
	validator_mocks "github.com/ava-labs/icm-services/peers/clients/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetLatestSyncedPChainHeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockValidatorClient := validator_mocks.NewMockCanonicalValidatorState(ctrl)

	subnetID := ids.GenerateTestID()
	validatorSet := map[ids.ID]snowVdrs.WarpSet{
		subnetID: {
			Validators:  []*snowVdrs.Warp{},
			TotalWeight: 0,
		},
	}

	testCases := []struct {
		name                   string
		fetchedHeight          uint64
		expectedSyncedHeight   uint64
		shouldCallValidatorAPI bool
		setupMock              func()
	}{
		{
			name:                   "initially returns zero",
			fetchedHeight:          0,
			expectedSyncedHeight:   0,
			shouldCallValidatorAPI: false,
			setupMock:              func() {},
		},
		{
			name:                   "updates after caching first height",
			fetchedHeight:          100,
			expectedSyncedHeight:   100,
			shouldCallValidatorAPI: true,
			setupMock: func() {
				mockValidatorClient.EXPECT().GetAllValidatorSets(
					gomock.Any(), uint64(100),
				).Return(validatorSet, nil).Times(1)
			},
		},
		{
			name:                   "does not update when caching lower height",
			fetchedHeight:          50,
			expectedSyncedHeight:   100,
			shouldCallValidatorAPI: true,
			setupMock: func() {
				mockValidatorClient.EXPECT().GetAllValidatorSets(
					gomock.Any(), uint64(50),
				).Return(validatorSet, nil).Times(1)
			},
		},
		{
			name:                   "updates when caching higher height",
			fetchedHeight:          200,
			expectedSyncedHeight:   200,
			shouldCallValidatorAPI: true,
			setupMock: func() {
				mockValidatorClient.EXPECT().GetAllValidatorSets(
					gomock.Any(), uint64(200),
				).Return(validatorSet, nil).Times(1)
			},
		},
		{
			name:                   "does not cache when fetching proposed height",
			fetchedHeight:          pchainapi.ProposedHeight,
			expectedSyncedHeight:   200,
			shouldCallValidatorAPI: true,
			setupMock: func() {
				mockValidatorClient.EXPECT().GetAllValidatorSets(
					gomock.Any(), gomock.Any(),
				).Return(validatorSet, nil).Times(1)
			},
		},
	}

	validatorManager := ValidatorManager{
		validatorClient:            mockValidatorClient,
		metrics:                    metrics,
		logger:                     logging.NoLog{},
		canonicalValidatorSetCache: cache.NewTTLCache[ids.ID, snowVdrs.WarpSet](canonicalValidatorSetCacheTTL),
		epochedValidatorSetCache:   cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](100),
		maxPChainLookback:          -1, // Disable lookback check for testing
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupMock()

			if testCase.shouldCallValidatorAPI {
				_, err := validatorManager.GetAllValidatorSets(t.Context(), testCase.fetchedHeight)
				require.NoError(t, err)
			}

			require.Equal(t, testCase.expectedSyncedHeight, validatorManager.GetLatestSyncedPChainHeight())
		})
	}
}

func TestConcurrentGetAllValidatorSetsUpdatesLatestSyncedHeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockValidatorClient := validator_mocks.NewMockCanonicalValidatorState(ctrl)

	validatorManager := ValidatorManager{
		validatorClient:            mockValidatorClient,
		metrics:                    metrics,
		logger:                     logging.NoLog{},
		canonicalValidatorSetCache: cache.NewTTLCache[ids.ID, snowVdrs.WarpSet](canonicalValidatorSetCacheTTL),
		epochedValidatorSetCache:   cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](100),
		maxPChainLookback:          -1, // Disable lookback check for testing
	}

	subnetID := ids.GenerateTestID()
	validatorSet := map[ids.ID]snowVdrs.WarpSet{
		subnetID: {
			Validators:  []*snowVdrs.Warp{},
			TotalWeight: 0,
		},
	}

	// Set up expectations for concurrent calls
	mockValidatorClient.EXPECT().GetAllValidatorSets(
		gomock.Any(), uint64(10),
	).Return(validatorSet, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(
		gomock.Any(), uint64(20),
	).Return(validatorSet, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(
		gomock.Any(), uint64(30),
	).Return(validatorSet, nil).AnyTimes()

	// Run concurrent calls
	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines * 3) // 3 heights per goroutine

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_, err := validatorManager.GetAllValidatorSets(t.Context(), 10)
			require.NoError(t, err)
		}()
		go func() {
			defer wg.Done()
			_, err := validatorManager.GetAllValidatorSets(t.Context(), 20)
			require.NoError(t, err)
		}()
		go func() {
			defer wg.Done()
			_, err := validatorManager.GetAllValidatorSets(t.Context(), 30)
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	// After all concurrent calls, latestSyncedPChainHeight should be 30 (the highest)
	require.Equal(t, uint64(30), validatorManager.GetLatestSyncedPChainHeight())
}
