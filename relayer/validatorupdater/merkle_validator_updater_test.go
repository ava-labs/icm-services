// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	merklevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

// fakeRegistryCaller serves getValidatorSetCommitment the way the contract does:
// a zero commitment for any blockchain ID it has no entry for.
type fakeRegistryCaller struct {
	commitments map[ids.ID]merklevalidatorsetregistry.ValidatorSetMerkleCommitment
	err         error
	queried     []ids.ID
}

func (f *fakeRegistryCaller) GetValidatorSetCommitment(
	_ *bind.CallOpts,
	avalancheBlockchainID [32]byte,
) (merklevalidatorsetregistry.ValidatorSetMerkleCommitment, error) {
	f.queried = append(f.queried, ids.ID(avalancheBlockchainID))
	if f.err != nil {
		return merklevalidatorsetregistry.ValidatorSetMerkleCommitment{}, f.err
	}
	return f.commitments[ids.ID(avalancheBlockchainID)], nil
}

// TestPChainAnchorHeight checks that P-chain-signed updates are anchored at the
// height of the registry's stored P-chain entry, and not at the entry of the
// chain being updated.
func TestPChainAnchorHeight(t *testing.T) {
	l1BlockchainID := ids.GenerateTestID()
	registry := &fakeRegistryCaller{
		commitments: map[ids.ID]merklevalidatorsetregistry.ValidatorSetMerkleCommitment{
			constants.PlatformChainID: {
				AvalancheBlockchainID: constants.PlatformChainID,
				TotalWeight:           100,
				PChainHeight:          1234,
			},
			l1BlockchainID: {
				AvalancheBlockchainID: l1BlockchainID,
				TotalWeight:           50,
				PChainHeight:          9999,
			},
		},
	}

	height, err := pChainAnchorHeight(context.Background(), registry)
	require.NoError(t, err)
	require.Equal(t, uint64(1234), height)
	require.Equal(t, []ids.ID{constants.PlatformChainID}, registry.queried)
}

// TestPChainAnchorHeight_PChainNotRegistered checks that a registry with no
// P-chain entry is rejected instead of yielding the zero commitment's height 0,
// which the P-chain would serve as a literal (genesis) height.
func TestPChainAnchorHeight_PChainNotRegistered(t *testing.T) {
	_, err := pChainAnchorHeight(context.Background(), &fakeRegistryCaller{})
	require.ErrorIs(t, err, errPChainNotRegistered)
}

func TestPChainAnchorHeight_Error(t *testing.T) {
	rpcErr := errors.New("rpc down")
	_, err := pChainAnchorHeight(context.Background(), &fakeRegistryCaller{err: rpcErr})
	require.ErrorIs(t, err, rpcErr)
}
