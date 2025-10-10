// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

import (
	"sync"

	"github.com/ava-labs/libevm/common/lru"
)

// LRUCache wraps LRU cache with TTL-like Get interface for historical/immutable data
type LRUCache[K comparable, V any] struct {
	cache *lru.Cache[K, V]
	lock  sync.RWMutex
}

func NewLRUCache[K comparable, V any](size int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		cache: lru.NewCache[K, V](size),
	}
}

// Get checks if the cached value exists for a given key, otherwise fetches
// the value using fetchFunc. Similar interface to TTLCache.Get but without TTL expiration.
// If [invalidate] is true, the value will be cleared from the cache prior to fetching.
// This is designed for historical/immutable data that doesn't need expiration.
func (c *LRUCache[K, V]) Get(key K, fetchFunc func(K) (V, error), invalidate bool) (V, error) {
	if invalidate {
		c.lock.Lock()
		c.cache.Remove(key)
		c.lock.Unlock()
	} else {
		c.lock.RLock()
		if value, found := c.cache.Get(key); found {
			c.lock.RUnlock()
			return value, nil
		}
		c.lock.RUnlock()
	}

	// Fetch new value
	newValue, err := fetchFunc(key)
	if err != nil {
		var zero V
		return zero, err
	}

	// Cache the result
	c.lock.Lock()
	c.cache.Add(key, newValue)
	c.lock.Unlock()

	return newValue, nil
}
