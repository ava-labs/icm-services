package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		skipCache     bool
		invalidate    bool
		expectedValue int
		expectedCount int
	}{
		{
			name:          "fresh cache, fetch",
			key:           "test1",
			skipCache:     false,
			invalidate:    false,
			expectedValue: 42,
			expectedCount: 1,
		},
		{
			name:          "use cache, no fetch",
			key:           "test1",
			skipCache:     false,
			invalidate:    false,
			expectedValue: 42,
			expectedCount: 1, // Same count as previous
		},
		{
			name:          "skipCache=true, fetch again",
			key:           "test1",
			skipCache:     true,
			invalidate:    false,
			expectedValue: 42,
			expectedCount: 1,
		},
		{
			name:          "invalidate=true, fetch again",
			key:           "test1",
			skipCache:     false,
			invalidate:    true,
			expectedValue: 42,
			expectedCount: 2,
		},
		{
			name:          "different key, fetch",
			key:           "test2",
			skipCache:     false,
			invalidate:    false,
			expectedValue: 42,
			expectedCount: 3,
		},
	}

	cache := NewLRUCache[string, int](10) // Size 10
	fetchCount := 0
	fetchFunc := func(key string) (int, error) {
		fetchCount++
		return 42, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			val, err := cache.Get(tt.key, fetchFunc, tt.invalidate)
			require.NoError(err)
			require.Equal(tt.expectedValue, val)
			require.Equal(tt.expectedCount, fetchCount)
		})
	}
}
