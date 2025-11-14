// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

import (
	"sync"
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
	inflight   map[K]*call[V]
	inflightLk sync.Mutex
}

// call represents an in-flight fetch operation
type call[V any] struct {
	wg  sync.WaitGroup
	val V
	err error
}

// NewFIFOCache creates a new FIFO cache with the given capacity and fetch function
func NewFIFOCache[K comparable, V any](capacity int) *FIFOCache[K, V] {
	return &FIFOCache[K, V]{
		cache:    make(map[K]V),
		queue:    make([]K, 0, capacity),
		capacity: capacity,
		inflight: make(map[K]*call[V]),
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

	// Single-flight mechanism
	c.inflightLk.Lock()
	if cl, ok := c.inflight[key]; ok {
		// Another goroutine is already fetching this key
		c.inflightLk.Unlock()
		cl.wg.Wait()
		return cl.val, cl.err
	}

	// We're the first to request this key
	cl := &call[V]{}
	cl.wg.Add(1)
	c.inflight[key] = cl
	c.inflightLk.Unlock()

	// Fetch the value
	val, err := fetchFunc(key)
	cl.val = val
	cl.err = err

	if err == nil {
		// Store in cache
		c.lk.Lock()
		c.set(key, val)
		c.lk.Unlock()
	}

	// Clean up inflight tracking
	c.inflightLk.Lock()
	delete(c.inflight, key)
	c.inflightLk.Unlock()

	cl.wg.Done()

	return val, err
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
