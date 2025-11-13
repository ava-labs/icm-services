// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cache

import (
	"fmt"

	"golang.org/x/sync/singleflight"
)

// SingleFlight provides a generic singleflight mechanism for deduplicating concurrent function calls.
// This is useful for caching scenarios where multiple goroutines might request the same value simultaneously.
type SingleFlight struct {
	group singleflight.Group
}

// NewSingleFlight creates a new SingleFlight instance.
func NewSingleFlight() *SingleFlight {
	return &SingleFlight{}
}

// Do executes and returns the results of the given function, making sure that only one execution
// is in-flight for a given key at a time. If a duplicate call comes in, the duplicate caller
// waits for the original to complete and receives the same results.
func (sf *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error, bool) {
	return sf.group.Do(key, fn)
}

// keyToString converts a comparable key to a string for use with singleflight.
// It handles both fmt.Stringer and primitive types.
func keyToString[K comparable](key K) string {
	if s, ok := any(key).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%v", key)
}
