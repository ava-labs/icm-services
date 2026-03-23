// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"crypto/rand"
	"crypto/sha256"
	"sort"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/stretchr/testify/require"

	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	blst "github.com/supranational/blst/bindings/go"
)

// TestValidatorSetHashRoundTrip verifies that our serializeValidatorsForHash and
// computeValidatorSetHash produce output compatible with Solidity's format.
// Creates validators, computes hash, builds a diff, parses it back, and verifies
// the hash matches what we'd expect from applying the diff.
func TestValidatorSetHashRoundTrip(t *testing.T) {
	// Create validators using blst (same library as P-chain)
	validators := make([]*message.Validator, 3)
	for i := 0; i < 3; i++ {
		var ikm [32]byte
		_, _ = rand.Read(ikm[:])
		sk := blst.KeyGen(ikm[:])
		require.NotNil(t, sk)
		pk := new(bls.PublicKey).From(sk)
		require.NotNil(t, pk)
		pkBytes := pk.Serialize()
		require.Len(t, pkBytes, 96, "BLS uncompressed public key must be 96 bytes")
		var pkArr [96]byte
		copy(pkArr[:], pkBytes)
		validators[i] = &message.Validator{
			UncompressedPublicKeyBytes: pkArr,
			Weight:                     uint64(i + 1),
		}
	}

	// Sort by key (required for serialization)
	sort.Slice(validators, func(i, j int) bool {
		return string(validators[i].UncompressedPublicKeyBytes[:]) <
			string(validators[j].UncompressedPublicKeyBytes[:])
	})

	// Compute our hash
	hash := computeValidatorSetHash(validators)
	require.NotEqual(t, ids.Empty, hash)

	// Build a diff with this hash
	blockchainID := ids.ID{'t', 'e', 's', 't'}
	changes := make([]message.ValidatorChange, len(validators))
	for i, v := range validators {
		changes[i] = message.ValidatorChange{
			UncompressedPublicKeyBytes: v.UncompressedPublicKeyBytes,
			Weight:                     v.Weight,
		}
	}

	diff, err := message.NewValidatorSetDiff(
		blockchainID,
		0, 0, // prev
		1, 1, // curr
		hash,
		changes,
		uint32(len(changes)),
	)
	require.NoError(t, err)

	// Parse back (simulates contract parsing)
	parsed, err := message.ParseValidatorSetDiff(diff.Bytes())
	require.NoError(t, err)
	require.Equal(t, blockchainID, parsed.BlockchainID)
	require.Len(t, parsed.Changes, len(changes))

	// Simulate apply: for empty current set, result = all changes with weight > 0
	var applied []*message.Validator
	for _, c := range parsed.Changes {
		if c.Weight > 0 {
			applied = append(applied, &message.Validator{
				UncompressedPublicKeyBytes: c.UncompressedPublicKeyBytes,
				Weight:                     c.Weight,
			})
		}
	}

	// Recompute hash of applied set - must match what we put in the diff
	appliedHash := computeValidatorSetHash(applied)
	require.Equal(t, hash, appliedHash, "Hash after apply must match diff.currentValidatorSetHash")

	// Verify serialized format: first 6 bytes should be 0x0000 (codec) + 0x00000005 (type) or similar
	serialized := serializeValidatorsForHash(validators)
	require.GreaterOrEqual(t, len(serialized), 6)
	require.Equal(t, uint8(0), serialized[0])
	require.Equal(t, uint8(0), serialized[1])
	// bytes 2-5 = validator count (big-endian)
	count := uint32(serialized[2])<<24 | uint32(serialized[3])<<16 | uint32(serialized[4])<<8 | uint32(serialized[5])
	require.Equal(t, uint32(len(validators)), count)
	// Total length: 2 + 4 + n*(96+8) = 6 + n*104
	require.Equal(t, 6+len(validators)*104, len(serialized))

	// Hash should be deterministic
	hash2 := ids.ID(sha256.Sum256(serialized))
	require.Equal(t, hash, hash2)
}

// padBlsPublicKeyForTest replicates Solidity's BLST.padUncompressedBLSPublicKey
// to produce 128-byte padded keys for testing onChainValidatorsToMessage.
func padBlsPublicKeyForTest(pk96 [96]byte) []byte {
	padded := make([]byte, 128)
	copy(padded[16:48], pk96[0:32])
	copy(padded[32:64], pk96[16:48])
	copy(padded[80:112], pk96[48:80])
	copy(padded[96:128], pk96[64:96])
	return padded
}

// TestOnChainValidatorsToMessageUnpad verifies that onChainValidatorsToMessage
// correctly unpads 128-byte BLS keys to match Solidity's unPadUncompressedBlsPublicKey.
func TestOnChainValidatorsToMessageUnpad(t *testing.T) {
	// Create a 96-byte BLS key
	var ikm [32]byte
	_, _ = rand.Read(ikm[:])
	sk := blst.KeyGen(ikm[:])
	require.NotNil(t, sk)
	pk := new(bls.PublicKey).From(sk)
	require.NotNil(t, pk)
	pkBytes := pk.Serialize()
	require.Len(t, pkBytes, 96)
	var pk96 [96]byte
	copy(pk96[:], pkBytes)

	// Pad it the same way Solidity does (as stored on-chain)
	padded := padBlsPublicKeyForTest(pk96)
	require.Len(t, padded, 128)

	// Convert via onChainValidatorsToMessage (simulates reading from chain)
	onChainValidators := []diffupdater.Validator{
		{BlsPublicKey: padded, Weight: 1000},
	}
	result := onChainValidatorsToMessage(onChainValidators)
	require.Len(t, result, 1)

	// Unpad must produce the original 96-byte key
	require.Equal(t, pk96, result[0].UncompressedPublicKeyBytes,
		"Unpad must match Solidity's unPadUncompressedBlsPublicKey output")
}
