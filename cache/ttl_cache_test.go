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

func TestTTLCacheGetWithExpiration(t *testing.T) {
	cache := NewTTLCache[string, int](1 * time.Second) // Fixed TTL is 1s
	fetchCount := 0
	fetchFunc := func(_ string) (int, error) {
		fetchCount++
		return 100, nil
	}

	// Test 1: Fresh cache with custom expiration, fetch
	t.Run("fresh cache with custom expiration, fetch", func(t *testing.T) {
		fetchCount = 0
		expirationFunc := func(_ int) time.Time {
			return time.Now().Add(5 * time.Second)
		}
		val, err := cache.GetWithExpiration("test1", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 1, fetchCount)
	})

	// Test 2: Use cache with custom expiration, no fetch
	t.Run("use cache with custom expiration, no fetch", func(t *testing.T) {
		expirationFunc := func(_ int) time.Time {
			return time.Now().Add(5 * time.Second)
		}
		val, err := cache.GetWithExpiration("test1", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 1, fetchCount) // Still 1, using cached value
	})

	// Test 3: Custom expiration not expired, use cache
	t.Run("custom expiration not expired, use cache", func(t *testing.T) {
		fetchCount = 0
		expirationFunc := func(_ int) time.Time {
			return time.Now().Add(5 * time.Second)
		}
		val, err := cache.GetWithExpiration("test2", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 1, fetchCount)

		// Wait 2 seconds - custom TTL is 5s, so still valid
		time.Sleep(2 * time.Second)
		val, err = cache.GetWithExpiration("test2", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 1, fetchCount) // Still cached, custom TTL is 5s
	})

	// Test 4: Custom expiration expired, fetch
	t.Run("custom expiration expired, fetch", func(t *testing.T) {
		fetchCount = 0
		expirationFunc := func(_ int) time.Time {
			return time.Now().Add(1 * time.Second)
		}
		val, err := cache.GetWithExpiration("test3", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 1, fetchCount)

		// Wait 2 seconds - custom TTL is 1s, so expired
		time.Sleep(2 * time.Second)
		val, err = cache.GetWithExpiration("test3", fetchFunc, expirationFunc, false)
		require.NoError(t, err)
		require.Equal(t, 100, val)
		require.Equal(t, 2, fetchCount) // Expired, fetch again
	})
}

func TestTTLCacheGetWithExpirationInvalidate(t *testing.T) {
	cache := NewTTLCache[string, int](10 * time.Second)
	fetchCount := 0
	fetchFunc := func(_ string) (int, error) {
		fetchCount++
		return 200, nil
	}

	expirationFunc := func(_ int) time.Time {
		return time.Now().Add(5 * time.Second)
	}

	// First fetch
	val, err := cache.GetWithExpiration("test", fetchFunc, expirationFunc, false)
	require.NoError(t, err)
	require.Equal(t, 200, val)
	require.Equal(t, 1, fetchCount)

	// Second fetch - should use cache
	val, err = cache.GetWithExpiration("test", fetchFunc, expirationFunc, false)
	require.NoError(t, err)
	require.Equal(t, 200, val)
	require.Equal(t, 1, fetchCount) // Still 1, using cache

	// Invalidate and fetch again
	val, err = cache.GetWithExpiration("test", fetchFunc, expirationFunc, true)
	require.NoError(t, err)
	require.Equal(t, 200, val)
	require.Equal(t, 2, fetchCount) // Fetched again after invalidation
}

func TestTTLCacheGetWithExpirationDifferentKeys(t *testing.T) {
	cache := NewTTLCache[string, int](1 * time.Second)
	fetchCount := 0
	fetchFunc := func(key string) (int, error) {
		fetchCount++
		if key == "key1" {
			return 10, nil
		}
		return 20, nil
	}

	// Key1 with 5 second expiration
	expirationFunc1 := func(_ int) time.Time {
		return time.Now().Add(5 * time.Second)
	}

	// Key2 with 2 second expiration
	expirationFunc2 := func(_ int) time.Time {
		return time.Now().Add(2 * time.Second)
	}

	// Fetch key1
	val1, err := cache.GetWithExpiration("key1", fetchFunc, expirationFunc1, false)
	require.NoError(t, err)
	require.Equal(t, 10, val1)
	require.Equal(t, 1, fetchCount)

	// Fetch key2
	val2, err := cache.GetWithExpiration("key2", fetchFunc, expirationFunc2, false)
	require.NoError(t, err)
	require.Equal(t, 20, val2)
	require.Equal(t, 2, fetchCount)

	// Both should be cached
	val1, err = cache.GetWithExpiration("key1", fetchFunc, expirationFunc1, false)
	require.NoError(t, err)
	require.Equal(t, 10, val1)

	val2, err = cache.GetWithExpiration("key2", fetchFunc, expirationFunc2, false)
	require.NoError(t, err)
	require.Equal(t, 20, val2)
	require.Equal(t, 2, fetchCount) // Both from cache

	// Wait for key2 to expire (2 seconds) but key1 should still be valid (5 seconds)
	time.Sleep(3 * time.Second)

	// Key1 should still be cached
	val1, err = cache.GetWithExpiration("key1", fetchFunc, expirationFunc1, false)
	require.NoError(t, err)
	require.Equal(t, 10, val1)
	require.Equal(t, 2, fetchCount) // Still from cache

	// Key2 should be expired and refetched
	val2, err = cache.GetWithExpiration("key2", fetchFunc, expirationFunc2, false)
	require.NoError(t, err)
	require.Equal(t, 20, val2)
	require.Equal(t, 3, fetchCount) // Refetched key2
}
