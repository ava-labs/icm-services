// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"sync"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Get must return a snapshot. The cache keeps adding to the map it stores for a message, so
// handing that map out would let a reader race with the next Add.
func TestSignatureCacheGetReturnsCopy(t *testing.T) {
	cache, err := NewSignatureCache(10)
	require.NoError(t, err)

	msgID := ids.GenerateTestID()
	_, ok := cache.Get(msgID)
	require.False(t, ok)

	cache.Add(msgID, PublicKeyBytes{1}, SignatureBytes{1})
	first, ok := cache.Get(msgID)
	require.True(t, ok)
	require.Equal(t, map[PublicKeyBytes]SignatureBytes{{1}: {1}}, first)

	// A later Add is not visible through a map returned earlier...
	cache.Add(msgID, PublicKeyBytes{2}, SignatureBytes{2})
	require.Len(t, first, 1)

	// ...and changing a returned map does not change the cache.
	first[PublicKeyBytes{3}] = SignatureBytes{3}
	second, ok := cache.Get(msgID)
	require.True(t, ok)
	require.Equal(t, map[PublicKeyBytes]SignatureBytes{{1}: {1}, {2}: {2}}, second)
}

// Concurrent requests for the same message read its cached signatures while responses for it
// are still being added. Run with -race; unsynchronized access to the same map is otherwise
// only detected by the runtime some of the time.
func TestSignatureCacheConcurrentAddAndGet(t *testing.T) {
	cache, err := NewSignatureCache(10)
	require.NoError(t, err)

	const (
		writers         = 4
		readers         = 4
		keysPerWriter   = 200
		readsPerReader  = 200
		expectedEntries = writers * keysPerWriter
	)
	msgID := ids.GenerateTestID()

	var wg sync.WaitGroup
	for w := range writers {
		wg.Go(func() {
			for k := range keysPerWriter {
				cache.Add(msgID, PublicKeyBytes{byte(w), byte(k)}, SignatureBytes{byte(w), byte(k)})
			}
		})
	}
	for range readers {
		wg.Go(func() {
			for range readsPerReader {
				sigs, _ := cache.Get(msgID)
				// assert, not require: FailNow must not be called off the test goroutine.
				for pubKey, sig := range sigs {
					assert.Equal(t, pubKey[:2], sig[:2])
				}
			}
		})
	}
	wg.Wait()

	sigs, ok := cache.Get(msgID)
	require.True(t, ok)
	require.Len(t, sigs, expectedEntries)
}
