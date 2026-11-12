// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	merklevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	"github.com/ava-labs/icm-services/peers/clients/mocks"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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

// fakeRegistry extends fakeRegistryCaller to everything the updater uses of the
// registry binding. It refuses updates, so a test that submits one fails.
type fakeRegistry struct {
	fakeRegistryCaller
}

func (*fakeRegistry) AllowPChainFallback(*bind.CallOpts) (bool, error) {
	return true, nil
}

func (*fakeRegistry) RegisterValidatorSet(
	*bind.TransactOpts,
	merklevalidatorsetregistry.ICMMessage,
	[32]byte,
) (*types.Transaction, error) {
	return nil, errors.New("unexpected registerValidatorSet")
}

// fakeTxTracker stands in for evm.TxTracker, whose own tests cover how a
// transaction is followed. It holds [unconfirmed] and resolves it as [settled],
// or fails to with [resolveErr].
type fakeTxTracker struct {
	unconfirmed *types.Transaction
	settled     bool
	resolveErr  error
	// notMined makes WaitMined give up on the transaction it is handed and keep
	// it as unconfirmed, as the tracker does at its inclusion timeout. Without it
	// a test that waits for a transaction fails.
	notMined bool

	// gateAsked and gateAllowed record the answer of the replacement gate the
	// updater passed to Resolve.
	gateAsked   bool
	gateAllowed bool
}

func (f *fakeTxTracker) WaitMined(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if !f.notMined {
		return nil, errors.New("unexpected WaitMined")
	}
	f.unconfirmed = tx
	return nil, fmt.Errorf("%w: %s", evm.ErrTxUnconfirmed, tx.Hash())
}

func (f *fakeTxTracker) Unconfirmed() *types.Transaction {
	return f.unconfirmed
}

func (f *fakeTxTracker) Resolve(
	ctx context.Context,
	allowReplacement func(context.Context) (bool, error),
) (bool, error) {
	if allowReplacement != nil {
		f.gateAsked = true
		f.gateAllowed, _ = allowReplacement(ctx)
	}
	if f.resolveErr != nil {
		return false, f.resolveErr
	}
	if f.settled {
		f.unconfirmed = nil
	}
	return f.settled, nil
}

// newTestUpdater returns an initialized updater whose transactions are followed
// by [tracker]. It has no P-chain client, registry or signature aggregator, so a
// test in which the updater reaches for one of them panics: the paths that must
// not build an update are held to that by construction.
func newTestUpdater(tracker *fakeTxTracker) *MerkleSetUpdater {
	return &MerkleSetUpdater{
		logger:         logging.NoLog{},
		txTracker:      tracker,
		blockchainID:   ids.GenerateTestID(),
		subnetID:       ids.GenerateTestID(),
		maxGasPriceWei: new(big.Int),
		initialized:    true,
	}
}

func newUnconfirmedTx() *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{Nonce: 7})
}

// While an update transaction is unconfirmed, the poll does nothing but follow
// it up: it builds no update and leaves the local state alone. Replacing the
// transaction is subject to the updater's gas price gate.
func TestCheckAndUpdateFollowsUpUnconfirmedTx(t *testing.T) {
	tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx()}
	s := newTestUpdater(tracker)

	require.NoError(t, s.checkAndUpdate(context.Background()))
	require.NotNil(t, tracker.Unconfirmed())
	require.True(t, s.initialized)
	// No max gas price is configured, so the gate lets a replacement through.
	require.True(t, tracker.gateAsked)
	require.True(t, tracker.gateAllowed)

	tracker.resolveErr = errors.New("rpc down")
	require.ErrorIs(t, s.checkAndUpdate(context.Background()), tracker.resolveErr)
	require.True(t, s.initialized)
}

// While an update transaction is unconfirmed, no second one is built behind its
// nonce: not by the P-Chain fallback on the same poll, and not on later polls.
func TestPerformUpdateRefusesSecondTransaction(t *testing.T) {
	s := newTestUpdater(&fakeTxTracker{unconfirmed: newUnconfirmedTx()})

	for _, signingChain := range []ids.ID{s.subnetID, constants.PrimaryNetworkID} {
		err := s.performUpdate(context.Background(), nil, signingChain)
		require.ErrorIs(t, err, evm.ErrTxUnconfirmed)
	}
}

// Once the unconfirmed transaction is settled, the updater re-reads its state
// from the contract before doing anything else, since it cannot know what the
// transaction did. It is then free to build the next update.
func TestCheckAndUpdateResyncsAfterSettledTx(t *testing.T) {
	const localPChainHeight = uint64(10)
	staleSince := time.Now().Add(-time.Hour)

	tests := map[string]struct {
		onChainHeight     uint64
		keepsStalenessAge bool
	}{
		// The update never reached the contract, which still holds the commitment the
		// updater mirrors. The staleness clock keeps running, so a refresh whose
		// transaction was dropped is rebuilt on the next poll.
		"update did not land": {
			onChainHeight:     localPChainHeight,
			keepsStalenessAge: true,
		},
		// The transaction was mined after the wait gave up: the contract holds the
		// new commitment.
		"update landed": {
			onChainHeight: localPChainHeight + 5,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tracker := &fakeTxTracker{unconfirmed: newUnconfirmedTx(), settled: true}
			s := newTestUpdater(tracker)
			s.localPChainHeight = localPChainHeight
			s.lastUpdateTime = staleSince

			registry := &fakeRegistry{fakeRegistryCaller{
				commitments: map[ids.ID]merklevalidatorsetregistry.ValidatorSetMerkleCommitment{
					s.blockchainID: {
						AvalancheBlockchainID: s.blockchainID,
						TotalWeight:           35,
						PChainHeight:          test.onChainHeight,
					},
				},
			}}
			s.contract = registry
			nodeValidators, _, _ := newSharedKeyValidatorSet(t)
			pChainClient := mocks.NewMockCanonicalValidatorState(gomock.NewController(t))
			pChainClient.EXPECT().
				GetValidatorsAt(gomock.Any(), s.subnetID, test.onChainHeight).
				Return(nodeValidators, nil)
			s.pChainClient = pChainClient

			require.NoError(t, s.checkAndUpdate(context.Background()))
			require.Nil(t, tracker.Unconfirmed())
			require.Equal(t, []ids.ID{s.blockchainID}, registry.queried)
			require.True(t, s.initialized)
			require.Equal(t, test.onChainHeight, s.localPChainHeight)
			require.Equal(t, uint64(35), s.localTotalWeight)
			if test.keepsStalenessAge {
				require.Equal(t, staleSince, s.lastUpdateTime)
			} else {
				require.WithinDuration(t, time.Now(), s.lastUpdateTime, time.Minute)
			}
		})
	}
}
