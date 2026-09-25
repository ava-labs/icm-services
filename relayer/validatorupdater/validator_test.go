// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers/clients/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// newSharedKeyValidatorSet returns a per-node validator set in which two nodes
// share a BLS key, one node has its own key, and one node has no BLS key. It
// returns the set alongside the signer whose key is shared.
func newSharedKeyValidatorSet(t *testing.T) (
	map[ids.NodeID]*validators.GetValidatorOutput,
	*localsigner.LocalSigner,
	*localsigner.LocalSigner,
) {
	sharedSigner, err := localsigner.New()
	require.NoError(t, err)
	soloSigner, err := localsigner.New()
	require.NoError(t, err)

	nodeIDs := make([]ids.NodeID, 4)
	for i := range nodeIDs {
		nodeIDs[i] = ids.GenerateTestNodeID()
	}
	nodeValidators := map[ids.NodeID]*validators.GetValidatorOutput{
		nodeIDs[0]: {NodeID: nodeIDs[0], PublicKey: sharedSigner.PublicKey(), Weight: 10},
		nodeIDs[1]: {NodeID: nodeIDs[1], PublicKey: sharedSigner.PublicKey(), Weight: 20},
		nodeIDs[2]: {NodeID: nodeIDs[2], PublicKey: soloSigner.PublicKey(), Weight: 5},
		nodeIDs[3]: {NodeID: nodeIDs[3], PublicKey: nil, Weight: 7},
	}
	return nodeValidators, sharedSigner, soloSigner
}

func TestFetchCanonicalValidators(t *testing.T) {
	nodeValidators, sharedSigner, soloSigner := newSharedKeyValidatorSet(t)
	subnetID := ids.GenerateTestID()
	const pChainHeight = uint64(42)

	ctrl := gomock.NewController(t)
	client := mocks.NewMockCanonicalValidatorState(ctrl)
	client.EXPECT().
		GetValidatorsAt(gomock.Any(), subnetID, pChainHeight).
		Return(nodeValidators, nil)

	got, err := FetchCanonicalValidators(context.Background(), client, subnetID, pChainHeight)
	require.NoError(t, err)

	// The keyless node is omitted and the two nodes sharing a key are merged
	// into one leaf with their weights summed.
	require.Len(t, got, 2)
	sharedKey := [96]byte(sharedSigner.PublicKey().Serialize())
	soloKey := [96]byte(soloSigner.PublicKey().Serialize())
	expected := []*Validator{
		{UncompressedPublicKeyBytes: sharedKey, Weight: 30},
		{UncompressedPublicKeyBytes: soloKey, Weight: 5},
	}
	SortValidators(expected)
	require.Equal(t, expected, got)

	// The result must match the canonical warp set exactly, in the same order.
	warpSet, err := validators.FlattenValidatorSet(nodeValidators)
	require.NoError(t, err)
	require.Equal(t, ValidatorsFromWarpSet(warpSet), got)
	require.True(t, bytes.Compare(
		got[0].UncompressedPublicKeyBytes[:],
		got[1].UncompressedPublicKeyBytes[:],
	) < 0)
}

func TestFetchCanonicalValidators_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mocks.NewMockCanonicalValidatorState(ctrl)
	client.EXPECT().
		GetValidatorsAt(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, context.DeadlineExceeded)

	_, err := FetchCanonicalValidators(context.Background(), client, ids.GenerateTestID(), 1)
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

// TestMerkleAttestationSignerBitsetIsCanonical checks that a signer bitset
// built over canonical (deduplicated) validator indices, as the signature
// aggregator produces, selects the merged validator when the attestation is
// built from the canonical set.
func TestMerkleAttestationSignerBitsetIsCanonical(t *testing.T) {
	nodeValidators, sharedSigner, _ := newSharedKeyValidatorSet(t)
	warpSet, err := validators.FlattenValidatorSet(nodeValidators)
	require.NoError(t, err)
	canonical := ValidatorsFromWarpSet(warpSet)
	require.Len(t, canonical, 2)

	sharedKey := [96]byte(sharedSigner.PublicKey().Serialize())
	sharedIndex := -1
	for i, vdr := range canonical {
		if vdr.UncompressedPublicKeyBytes == sharedKey {
			sharedIndex = i
		}
	}
	require.NotEqual(t, -1, sharedIndex)

	msg := []byte("merkle attestation")
	sig, err := sharedSigner.Sign(msg)
	require.NoError(t, err)

	signerBits := set.NewBits(sharedIndex)
	bitSetSig := &avalancheWarp.BitSetSignature{
		Signers:   signerBits.Bytes(),
		Signature: [bls.SignatureLen]byte(bls.SignatureToBytes(sig)),
	}

	attestation, err := NewValidatorSetMerkleAttestation(canonical, bitSetSig)
	require.NoError(t, err)

	// The single signer leaf is the merged entry: the shared key with the
	// summed weight of both nodes, which is the key the signature verifies against.
	require.Len(t, attestation.Signers, 1)
	require.Equal(t, sharedKey, attestation.Signers[0].UncompressedPublicKeyBytes)
	require.Equal(t, uint64(30), attestation.Signers[0].Weight)
	require.Equal(t, [192]byte(sig.Serialize()), attestation.AggregateSignature)
	require.True(t, bls.Verify(sharedSigner.PublicKey(), sig, msg))
}
