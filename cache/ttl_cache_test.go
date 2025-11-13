package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTTLCacheSingleKey(t *testing.T) {
	tests := []struct {
		name           string
		key            string
		expectedValue  int
		skipCache      bool
		waitBeforeNext time.Duration
		expectedCount  int
	}{
		{
			name:           "fresh cache, fetch",
			waitBeforeNext: 0,
			skipCache:      false,
			expectedCount:  1,
		},
		{
			name:           "use cache, no fetch",
			waitBeforeNext: 0,
			skipCache:      false,
			expectedCount:  1,
		},
		{
			name:           "skipCache=true, fetch",
			waitBeforeNext: 0,
			skipCache:      true,
			expectedCount:  2,
		},
		{
			name:           "ttl expired, fetch",
			waitBeforeNext: 2 * time.Second,
			skipCache:      false,
			expectedCount:  3,
		},
	}
	cache := NewTTLCache[string, int](1 * time.Second)
	fetchCount := 0
	fetchFunc := func(_ string) (int, error) {
		fetchCount++
		return 42, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			if tt.waitBeforeNext > 0 {
				time.Sleep(tt.waitBeforeNext)
			}

			val, err := cache.Get("test", fetchFunc, tt.skipCache)
			require.NoError(err)
			require.Equal(42, val)
			require.Equal(tt.expectedCount, fetchCount)
		})
	}
}
