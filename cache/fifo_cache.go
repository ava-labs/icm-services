// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

import (
	"sync"

	"golang.org/x/sync/singleflight"
)

// FetchFunc is the function signature for fetching values
type FetchFunc[K comparable, V any] func(key K) (V, error)

// FIFOCache is a thread-safe FIFO cache with single-flight fetching
type FIFOCache[K comparable, V any] struct {
	lk       sync.RWMutex
	cache    map[K]V
	queue    []K
	capacity int

	// Single-flight mechanism
	singleFlight singleflight.Group
}

// NewFIFOCache creates a new FIFO cache with the given capacity and fetch function
func NewFIFOCache[K comparable, V any](capacity int) *FIFOCache[K, V] {
	return &FIFOCache[K, V]{
		cache:    make(map[K]V),
		queue:    make([]K, 0, capacity),
		capacity: capacity,
	}
}

// Get retrieves a value from the cache or fetches it if not present
// If multiple goroutines call Get for the same key concurrently, only one fetch occurs
func (c *FIFOCache[K, V]) Get(key K, fetchFunc FetchFunc[K, V]) (V, error) {
	// Fast path: check if it's already in cache
	c.lk.RLock()
	if val, ok := c.cache[key]; ok {
		c.lk.RUnlock()
		return val, nil
	}
	c.lk.RUnlock()

	// Use singleflight to deduplicate concurrent fetches
	keyStr := keyToString(key)
	result, err, _ := c.singleFlight.Do(keyStr, func() (interface{}, error) {
		val, fetchErr := fetchFunc(key)
		if fetchErr != nil {
			return *new(V), fetchErr
		}

		// Store in cache
		c.lk.Lock()
		c.set(key, val)
		c.lk.Unlock()

		return val, nil
	})

	if err != nil {
		return *new(V), err
	}

	return result.(V), nil
}

// set adds a key-value pair to the cache (caller must hold write lock)
func (c *FIFOCache[K, V]) set(key K, val V) {
	// If key already exists, don't add to queue again
	if _, exists := c.cache[key]; exists {
		c.cache[key] = val
		return
	}

	// Evict oldest if at capacity
	if len(c.queue) >= c.capacity {
		oldest := c.queue[0]
		c.queue = c.queue[1:]
		delete(c.cache, oldest)
	}

	c.cache[key] = val
	c.queue = append(c.queue, key)
}

// Len returns the current number of items in the cache
func (c *FIFOCache[K, V]) Len() int {
	c.lk.RLock()
	defer c.lk.RUnlock()
	return len(c.cache)
}
