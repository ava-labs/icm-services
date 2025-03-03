package cache

import (
	"math"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/pingcap/errors"
)

type Cache struct {
	// map of warp message ID to a map of public keys to signatures
	signatures *lru.Cache[ids.ID, map[PublicKeyBytes]SignatureBytes]
}

type PublicKeyBytes [bls.PublicKeyLen]byte
type SignatureBytes [bls.SignatureLen]byte

func NewCache(size uint64) (*Cache, error) {
	if size > math.MaxInt {
		return nil, errors.New("cache size too big")
	}

	signatureCache, err := lru.New[ids.ID, map[PublicKeyBytes]SignatureBytes](int(size))
	if err != nil {
		return nil, err
	}

	return &Cache{
		signatures: signatureCache,
	}, nil
}

func (c *Cache) Get(msgID ids.ID) (map[PublicKeyBytes]SignatureBytes, bool) {
	return c.signatures.Get(msgID)
}

func (c *Cache) Add(
	msgID ids.ID,
	pubKey PublicKeyBytes,
	signature SignatureBytes,
) {
	var (
		sigs map[PublicKeyBytes]SignatureBytes
		ok   bool
	)
	if sigs, ok = c.Get(msgID); !ok {
		sigs = make(map[PublicKeyBytes]SignatureBytes)
	}
	sigs[pubKey] = signature
	c.signatures.Add(msgID, sigs)
}
