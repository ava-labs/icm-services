// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
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
	// protects against the race condition where multiple goroutines are trying to
	// add a signature for a message ID that is not currently in the cache.
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

func (c *SignatureCache) Get(msgID ids.ID) (map[PublicKeyBytes]SignatureBytes, bool) {
	return c.signatures.Get(msgID)
}

func (c *SignatureCache) Add(
	msgID ids.ID,
	pubKey PublicKeyBytes,
	signature SignatureBytes,
) {
	var (
		sigs map[PublicKeyBytes]SignatureBytes
		ok   bool
	)

	// The number of signatures cached per message is implicitly bounded
	// by the number of validators registered on-chain.
	// As a result, uncontrolled memory growth is not a concern.
	c.mu.Lock()
	if sigs, ok = c.Get(msgID); !ok {
		sigs = make(map[PublicKeyBytes]SignatureBytes)
	}
	sigs[pubKey] = signature
	c.signatures.Add(msgID, sigs)
	c.mu.Unlock()
}
