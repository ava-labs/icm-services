package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFIFOCache(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		expectedValue int
		expectedCount int
	}{
		{
			name:          "fresh cache, fetch",
			key:           "test1",
			expectedValue: 42,
			expectedCount: 1,
		},
		{
			name:          "use cache, no fetch",
			key:           "test1",
			expectedValue: 42,
			expectedCount: 1, // Same count as previous
		},
		{
			name:          "different key, fetch",
			key:           "test2",
			expectedValue: 42,
			expectedCount: 2,
		},
		{
			name:          "different key, fetch",
			key:           "test3",
			expectedValue: 42,
			expectedCount: 3,
		},
		{
			name:          "first item evicted, fetch",
			key:           "test1",
			expectedValue: 42,
			expectedCount: 4,
		},
	}

	cache := NewFIFOCache[string, int](2)
	fetchCount := 0
	fetchFunc := func(key string) (int, error) {
		fetchCount++
		return 42, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			val, err := cache.Get(tt.key, fetchFunc)
			require.NoError(err)
			require.Equal(tt.expectedValue, val)
			require.Equal(tt.expectedCount, fetchCount)
		})
	}
}
