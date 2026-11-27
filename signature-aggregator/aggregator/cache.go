// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"maps"
	"math"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/pingcap/errors"
)

type SignatureCache struct {
	// map of warp message ID to a map of public keys to signatures
	signatures *lru.Cache[ids.ID, map[PublicKeyBytes]SignatureBytes]
	// Guards the per-message maps. The LRU synchronizes its own entries but not the maps it
	// stores, which Add writes to (including the first write for a message ID that is not yet
	// cached) and Get copies.
	mu sync.Mutex
}

type PublicKeyBytes [bls.PublicKeyLen]byte
type SignatureBytes [bls.SignatureLen]byte

func NewSignatureCache(size uint64) (*SignatureCache, error) {
	if size > math.MaxInt {
		return nil, errors.New("cache size too big")
	}

	signatureCache, err := lru.New[ids.ID, map[PublicKeyBytes]SignatureBytes](int(size))
	if err != nil {
		return nil, err
	}

	return &SignatureCache{
		signatures: signatureCache,
	}, nil
}

// Get returns a copy of the signatures cached for [msgID]. The cached map itself is never handed
// out: [SignatureCache.Add] keeps writing to it, and an unsynchronized read of a map that is
// being written to is a fatal error.
func (c *SignatureCache) Get(msgID ids.ID) (map[PublicKeyBytes]SignatureBytes, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	sigs, ok := c.signatures.Get(msgID)
	if !ok {
		return nil, false
	}
	return maps.Clone(sigs), true
}

func (c *SignatureCache) Add(
	msgID ids.ID,
	pubKey PublicKeyBytes,
	signature SignatureBytes,
) {
	// The number of signatures cached per message is implicitly bounded
	// by the number of validators registered on-chain.
	// As a result, uncontrolled memory growth is not a concern.
	c.mu.Lock()
	defer c.mu.Unlock()

	sigs, ok := c.signatures.Get(msgID)
	if !ok {
		sigs = make(map[PublicKeyBytes]SignatureBytes)
	}
	sigs[pubKey] = signature
	c.signatures.Add(msgID, sigs)
}
