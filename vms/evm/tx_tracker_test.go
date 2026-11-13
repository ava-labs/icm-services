// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
)

// fakeTxTrackerClient is a chain as a TxTracker sees it through its RPC client. A transaction it
// holds no receipt for is not mined, and one that is not pending either is unknown to it.
type fakeTxTrackerClient struct {
	receipts map[common.Hash]*types.Receipt
	pending  map[common.Hash]bool
	// receiptErr fails receipt lookups, like a node that is down.
	receiptErr error
	// hang makes receipt lookups block until the caller gives up, like a node that accepts
	// requests and never answers them.
	hang bool

	baseFee  *big.Int
	tipCap   *big.Int
	gasPrice *big.Int

	sent []*types.Transaction
}

func (c *fakeTxTrackerClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	if c.hang {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	if c.receiptErr != nil {
		return nil, c.receiptErr
	}
	if receipt, ok := c.receipts[txHash]; ok {
		return receipt, nil
	}
	return nil, ethereum.NotFound
}

func (*fakeTxTrackerClient) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return nil, nil
}

func (c *fakeTxTrackerClient) TransactionByHash(
	_ context.Context,
	txHash common.Hash,
) (*types.Transaction, bool, error) {
	if c.pending[txHash] {
		return nil, true, nil
	}
	return nil, false, ethereum.NotFound
}

func (c *fakeTxTrackerClient) HeaderByNumber(context.Context, *big.Int) (*types.Header, error) {
	return &types.Header{BaseFee: c.baseFee}, nil
}

func (c *fakeTxTrackerClient) SuggestGasPrice(context.Context) (*big.Int, error) {
	return c.gasPrice, nil
}

func (c *fakeTxTrackerClient) SuggestGasTipCap(context.Context) (*big.Int, error) {
	return c.tipCap, nil
}

func (c *fakeTxTrackerClient) SendTransaction(_ context.Context, tx *types.Transaction) error {
	c.sent = append(c.sent, tx)
	return nil
}

const txTrackerTestInclusionTimeout = 50 * time.Millisecond

var (
	txTrackerTestChainID   = big.NewInt(1337)
	txTrackerTestRecipient = common.HexToAddress("0x00000000000000000000000000000000000000aa")
)

// newTestTxTracker returns a tracker for the transactions of a fresh account on [client].
func newTestTxTracker(t *testing.T, client *fakeTxTrackerClient) *TxTracker {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	txOpts, err := bind.NewKeyedTransactorWithChainID(key, txTrackerTestChainID)
	require.NoError(t, err)
	return NewTxTracker(logging.NoLog{}, client, txOpts.From, txOpts.Signer, txTrackerTestInclusionTimeout)
}

// newTrackedTx returns a transaction from the tracked account, signed the way a contract binding
// signs a call.
func newTrackedTx(t *testing.T, tracker *TxTracker, tipCap int64, feeCap int64) *types.Transaction {
	t.Helper()
	tx, err := tracker.signer(tracker.from, types.NewTx(&types.DynamicFeeTx{
		ChainID:   txTrackerTestChainID,
		Nonce:     7,
		GasTipCap: big.NewInt(tipCap),
		GasFeeCap: big.NewInt(feeCap),
		Gas:       300_000,
		To:        &txTrackerTestRecipient,
		Value:     new(big.Int),
		Data:      []byte{0x01, 0x02, 0x03},
	}))
	require.NoError(t, err)
	return tx
}

// requireReplaces checks that [replacement] is what the tracker now follows in place of [tx]:
// the same call from the same account at the same nonce.
func requireReplaces(t *testing.T, tracker *TxTracker, tx *types.Transaction, replacement *types.Transaction) {
	t.Helper()
	require.Same(t, replacement, tracker.Unconfirmed())
	sender, err := types.Sender(types.LatestSignerForChainID(txTrackerTestChainID), replacement)
	require.NoError(t, err)
	require.Equal(t, tracker.from, sender)
	require.Equal(t, tx.Type(), replacement.Type())
	require.Equal(t, tx.Nonce(), replacement.Nonce())
	require.Equal(t, tx.To(), replacement.To())
	require.Equal(t, tx.Data(), replacement.Data())
	require.Equal(t, tx.Gas(), replacement.Gas())
}

// A node can accept a transaction and never produce a receipt for it. The wait for one must end
// at the inclusion timeout, leaving the transaction to be followed up, instead of holding the
// caller for good.
func TestTxTrackerWaitMinedIsBounded(t *testing.T) {
	clients := map[string]*fakeTxTrackerClient{
		"no receipt":         {},
		"node never answers": {hang: true},
	}
	for name, client := range clients {
		t.Run(name, func(t *testing.T) {
			tracker := newTestTxTracker(t, client)
			tx := newTrackedTx(t, tracker, 2, 82)

			start := time.Now()
			_, err := tracker.WaitMined(context.Background(), tx)
			require.ErrorIs(t, err, ErrTxUnconfirmed)
			// Receipts are polled once a second, so returning sooner than that shows it is
			// the timeout that ended the wait.
			require.Less(t, time.Since(start), time.Second)
			require.Same(t, tx, tracker.Unconfirmed())
		})
	}
}

// A mined transaction has a known outcome whether it succeeded or reverted, so it is returned
// with its receipt and not kept.
func TestTxTrackerWaitMinedReturnsReceipt(t *testing.T) {
	for _, status := range []uint64{types.ReceiptStatusSuccessful, types.ReceiptStatusFailed} {
		client := &fakeTxTrackerClient{}
		tracker := newTestTxTracker(t, client)
		tx := newTrackedTx(t, tracker, 2, 82)
		client.receipts = map[common.Hash]*types.Receipt{
			tx.Hash(): {Status: status, BlockNumber: big.NewInt(5)},
		}

		receipt, err := tracker.WaitMined(context.Background(), tx)
		require.NoError(t, err)
		require.Equal(t, status, receipt.Status)
		require.Nil(t, tracker.Unconfirmed())
	}
}

// The wait also ends when the caller's context does, which must not be mistaken for a
// transaction that was not mined in time.
func TestTxTrackerWaitMinedEndsWithContext(t *testing.T) {
	tracker := newTestTxTracker(t, &fakeTxTrackerClient{})
	tx := newTrackedTx(t, tracker, 2, 82)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := tracker.WaitMined(ctx, tx)
	require.ErrorIs(t, err, context.Canceled)
	require.NotErrorIs(t, err, ErrTxUnconfirmed)
	require.Nil(t, tracker.Unconfirmed())
}

func TestTxTrackerResolveWithNothingUnconfirmed(t *testing.T) {
	tracker := newTestTxTracker(t, &fakeTxTrackerClient{})

	settled, err := tracker.Resolve(context.Background(), nil)
	require.NoError(t, err)
	require.True(t, settled)
}

// A transaction that was mined after the wait gave up, or that the node no longer knows, is
// settled: the tracker forgets it and sends nothing.
func TestTxTrackerResolveSettles(t *testing.T) {
	for _, mined := range []bool{true, false} {
		client := &fakeTxTrackerClient{}
		tracker := newTestTxTracker(t, client)
		tx := newTrackedTx(t, tracker, 2, 82)
		if mined {
			client.receipts = map[common.Hash]*types.Receipt{
				tx.Hash(): {Status: types.ReceiptStatusSuccessful, BlockNumber: big.NewInt(5)},
			}
		}
		tracker.unconfirmed = tx

		settled, err := tracker.Resolve(context.Background(), nil)
		require.NoError(t, err)
		require.True(t, settled)
		require.Nil(t, tracker.Unconfirmed())
		require.Empty(t, client.sent)
	}
}

// When the node cannot be asked, the outcome stays unknown and the transaction is kept.
func TestTxTrackerResolveKeepsTxOnRPCFailure(t *testing.T) {
	rpcErr := errors.New("rpc down")
	tracker := newTestTxTracker(t, &fakeTxTrackerClient{receiptErr: rpcErr})
	tx := newTrackedTx(t, tracker, 2, 82)
	tracker.unconfirmed = tx

	settled, err := tracker.Resolve(context.Background(), nil)
	require.ErrorIs(t, err, rpcErr)
	require.False(t, settled)
	require.Same(t, tx, tracker.Unconfirmed())
}

// A transaction that is still pending at the fees the node suggests now is left to be mined.
func TestTxTrackerResolveWaitsForPendingTx(t *testing.T) {
	// The node suggests a tip of 2 and, with it, a fee cap of 2 + 2*40 = 82.
	client := &fakeTxTrackerClient{baseFee: big.NewInt(40), tipCap: big.NewInt(2)}
	tracker := newTestTxTracker(t, client)
	tx := newTrackedTx(t, tracker, 2, 82)
	client.pending = map[common.Hash]bool{tx.Hash(): true}
	tracker.unconfirmed = tx

	settled, err := tracker.Resolve(context.Background(), nil)
	require.NoError(t, err)
	require.False(t, settled)
	require.Same(t, tx, tracker.Unconfirmed())
	require.Empty(t, client.sent)
}

// A pending transaction that the node's suggested fees have left behind is replaced by the same
// call from the same account at the same nonce, at the suggested fees and at least the minimum a
// node accepts for a replacement.
func TestTxTrackerResolveReplacesUnderpricedTx(t *testing.T) {
	tests := map[string]struct {
		baseFee, suggestedTip int64
		oldTip, oldFeeCap     int64
		newTip, newFeeCap     int64
	}{
		// The base fee rose from 40 to 100, past the old fee cap. The suggested fee cap is
		// 2 + 2*100; the suggested tip is unchanged, so it gets the minimum bump.
		"base fee rose": {
			baseFee: 100, suggestedTip: 2,
			oldTip: 2, oldFeeCap: 82,
			newTip: 3, newFeeCap: 202,
		},
		// The suggested tip rose from 2 to 5, while the suggested fee cap of 5 + 2*90 is below
		// the old one, which gets the minimum bump.
		"tip rose": {
			baseFee: 90, suggestedTip: 5,
			oldTip: 2, oldFeeCap: 200,
			newTip: 5, newFeeCap: 220,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &fakeTxTrackerClient{
				baseFee: big.NewInt(test.baseFee),
				tipCap:  big.NewInt(test.suggestedTip),
			}
			tracker := newTestTxTracker(t, client)
			tx := newTrackedTx(t, tracker, test.oldTip, test.oldFeeCap)
			client.pending = map[common.Hash]bool{tx.Hash(): true}
			tracker.unconfirmed = tx

			settled, err := tracker.Resolve(context.Background(), nil)
			require.NoError(t, err)
			require.False(t, settled)

			require.Len(t, client.sent, 1)
			replacement := client.sent[0]
			requireReplaces(t, tracker, tx, replacement)
			require.Equal(t, test.newTip, replacement.GasTipCap().Int64())
			require.Equal(t, test.newFeeCap, replacement.GasFeeCap().Int64())
		})
	}
}

// On a chain without a base fee a contract binding sends a legacy transaction. Its replacement is
// one too, at the node's suggested gas price.
func TestTxTrackerResolveReplacesUnderpricedLegacyTx(t *testing.T) {
	client := &fakeTxTrackerClient{gasPrice: big.NewInt(50)}
	tracker := newTestTxTracker(t, client)
	tx, err := tracker.signer(tracker.from, types.NewTx(&types.LegacyTx{
		Nonce:    7,
		GasPrice: big.NewInt(40),
		Gas:      300_000,
		To:       &txTrackerTestRecipient,
		Data:     []byte{0x01, 0x02, 0x03},
	}))
	require.NoError(t, err)
	client.pending = map[common.Hash]bool{tx.Hash(): true}
	tracker.unconfirmed = tx

	settled, err := tracker.Resolve(context.Background(), nil)
	require.NoError(t, err)
	require.False(t, settled)

	require.Len(t, client.sent, 1)
	replacement := client.sent[0]
	requireReplaces(t, tracker, tx, replacement)
	require.Equal(t, int64(50), replacement.GasPrice().Int64())
}

// The caller decides whether an outbid transaction may be replaced right now. It is asked only
// once a replacement is due, and its refusal or failure leaves the transaction as it is.
func TestTxTrackerResolveAsksBeforeReplacing(t *testing.T) {
	gateErr := errors.New("gas price unavailable")
	tests := map[string]struct {
		underpriced bool
		allow       bool
		allowErr    error
		asked       bool
		replaced    bool
	}{
		"not asked while priced right": {underpriced: false, allow: true},
		"allowed":                      {underpriced: true, allow: true, asked: true, replaced: true},
		"refused":                      {underpriced: true, allow: false, asked: true},
		"failed":                       {underpriced: true, allowErr: gateErr, asked: true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			baseFee := int64(40)
			if test.underpriced {
				baseFee = 100
			}
			client := &fakeTxTrackerClient{baseFee: big.NewInt(baseFee), tipCap: big.NewInt(2)}
			tracker := newTestTxTracker(t, client)
			tx := newTrackedTx(t, tracker, 2, 82)
			client.pending = map[common.Hash]bool{tx.Hash(): true}
			tracker.unconfirmed = tx

			asked := false
			settled, err := tracker.Resolve(context.Background(), func(context.Context) (bool, error) {
				asked = true
				return test.allow, test.allowErr
			})
			require.ErrorIs(t, err, test.allowErr)
			require.False(t, settled)
			require.Equal(t, test.asked, asked)
			if test.replaced {
				require.Len(t, client.sent, 1)
				require.Same(t, client.sent[0], tracker.Unconfirmed())
			} else {
				require.Empty(t, client.sent)
				require.Same(t, tx, tracker.Unconfirmed())
			}
		})
	}
}

func TestReplacementFee(t *testing.T) {
	for fee, expected := range map[int64]int64{0: 0, 1: 2, 10: 11, 100: 110, 101: 112} {
		require.Equal(t, expected, replacementFee(big.NewInt(fee)).Int64(), "fee %d", fee)
	}
}
