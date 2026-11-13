// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/math"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

// ErrTxUnconfirmed is returned by [TxTracker.WaitMined] when a transaction was broadcast but not
// mined within the inclusion timeout. The transaction neither succeeded nor failed: the tracker
// keeps it and [TxTracker.Resolve] follows it up.
var ErrTxUnconfirmed = errors.New("transaction not mined within the inclusion timeout")

// replacementFeeBumpPercent is how much more a transaction must offer than the pending one it
// replaces at the same nonce, in both its fee cap and its tip, for a node to accept it. This is
// the default of geth and of the clients that follow it.
const replacementFeeBumpPercent = 10

// TxTrackerClient is the subset of an EVM chain's RPC client a [TxTracker] uses to follow, price
// and replace transactions.
type TxTrackerClient interface {
	// TransactionReceipt and CodeAt, which bind.WaitMined needs.
	bind.DeployBackend
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

var _ TxTrackerClient = (*ethclient.Client)(nil)

// TxTracker follows the transactions of a sender that has one transaction in flight at a time,
// such as a polling loop that submits through a contract binding and waits for the result. It
// takes over once the transaction is broadcast and sees it through to a known outcome.
//
// Nothing guarantees that a transaction a node has accepted is ever mined: it can be dropped from
// the pool, or outbid for as long as fees stay above its cap. So [TxTracker.WaitMined] waits for
// a receipt only up to the inclusion timeout. A transaction that is not mined by then is kept as
// unconfirmed, and the sender calls [TxTracker.Resolve] before it sends anything else, typically
// at the top of its next poll, because a second transaction would only queue behind the nonce of
// the first.
//
// A TxTracker is not safe for concurrent use.
type TxTracker struct {
	logger logging.Logger
	client TxTrackerClient
	// from and signer are the account the tracked transactions are sent from. They sign the
	// replacement of a transaction that was outbid.
	from             common.Address
	signer           bind.SignerFn
	inclusionTimeout time.Duration

	// unconfirmed is the last transaction that was not mined within inclusionTimeout. Its outcome
	// is unknown: it may yet be mined, still be waiting in the transaction pool, or have been
	// dropped.
	unconfirmed *types.Transaction
}

// NewTxTracker returns a tracker for the transactions [from] sends through [client]. [signer]
// must sign for [from]; both are the fields of the same name of the bind.TransactOpts the
// transactions are sent with.
func NewTxTracker(
	logger logging.Logger,
	client TxTrackerClient,
	from common.Address,
	signer bind.SignerFn,
	inclusionTimeout time.Duration,
) *TxTracker {
	return &TxTracker{
		logger:           logger,
		client:           client,
		from:             from,
		signer:           signer,
		inclusionTimeout: inclusionTimeout,
	}
}

// WaitMined waits for the broadcast transaction [tx] to be mined, for at most the inclusion
// timeout, and returns its receipt. A mined transaction may still have reverted, which the
// receipt's status tells. If the transaction is not mined in time, it becomes the unconfirmed
// transaction and the error wraps [ErrTxUnconfirmed]. If [ctx] ends first, its error is returned
// and the transaction is not kept.
func (t *TxTracker) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	waitCtx, cancel := context.WithTimeout(ctx, t.inclusionTimeout)
	defer cancel()

	receipt, err := bind.WaitMined(waitCtx, t.client, tx)
	if err == nil {
		return receipt, nil
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	t.unconfirmed = tx
	t.logger.Warn("Transaction not mined within the inclusion timeout",
		zap.Stringer("txHash", tx.Hash()),
		zap.Uint64("nonce", tx.Nonce()),
		zap.Duration("inclusionTimeout", t.inclusionTimeout),
	)
	return nil, fmt.Errorf("%w: %s", ErrTxUnconfirmed, tx.Hash())
}

// Unconfirmed returns the transaction whose outcome is still unknown, or nil if there is none.
func (t *TxTracker) Unconfirmed() *types.Transaction {
	return t.unconfirmed
}

// Resolve finds out what became of the unconfirmed transaction and reports whether it is
// settled, meaning that it can no longer take effect and no longer holds its nonce. With no
// unconfirmed transaction there is nothing to settle and it returns true.
//
//   - Mined, successfully or not: settled.
//   - Unknown to the node, so neither mined nor pending: settled. It was dropped from the pool,
//     or another transaction took its nonce.
//   - Still pending: not settled. It is left alone while the fees it offers are at least the
//     ones the node suggests now, and replaced once they are not.
//
// A replacement is the same call from the same account at the same nonce, priced the way bind
// prices a new transaction and raised to the minimum a node accepts for a replacement. It becomes
// the unconfirmed transaction. It is only ever made when the node's suggestion exceeds what the
// pending transaction offers, so the fees paid follow the suggestion instead of compounding on
// every call. [allowReplacement], if not nil, is asked first and can defer it, for example while
// gas is above what the sender is willing to pay.
//
// The tracker forgets a settled transaction. Whatever it changed, the caller has not observed, so
// a caller that mirrors on-chain state should re-read it before sending again.
func (t *TxTracker) Resolve(
	ctx context.Context,
	allowReplacement func(context.Context) (bool, error),
) (bool, error) {
	tx := t.unconfirmed
	if tx == nil {
		return true, nil
	}
	// What follows is a handful of RPCs. They share one deadline, so that a node that stops
	// answering costs the caller one call, not its loop.
	ctx, cancel := context.WithTimeout(ctx, t.inclusionTimeout)
	defer cancel()

	log := t.logger.With(
		zap.Stringer("txHash", tx.Hash()),
		zap.Uint64("nonce", tx.Nonce()),
	)

	receipt, err := t.client.TransactionReceipt(ctx, tx.Hash())
	switch {
	case err == nil:
		log.Info("Unconfirmed transaction was mined",
			zap.Uint64("status", receipt.Status),
			zap.Stringer("blockNumber", receipt.BlockNumber),
		)
		t.unconfirmed = nil
		return true, nil
	case !errors.Is(err, ethereum.NotFound):
		return false, fmt.Errorf("failed to get receipt of unconfirmed transaction %s: %w", tx.Hash(), err)
	}

	_, isPending, err := t.client.TransactionByHash(ctx, tx.Hash())
	switch {
	case errors.Is(err, ethereum.NotFound):
		log.Warn("Unconfirmed transaction is neither mined nor pending")
		t.unconfirmed = nil
		return true, nil
	case err != nil:
		return false, fmt.Errorf("failed to look up unconfirmed transaction %s: %w", tx.Hash(), err)
	case !isPending:
		// Mined, but the receipt is not being served yet. The next call will find it.
		return false, nil
	}

	return false, t.replaceIfUnderpriced(ctx, log, tx, allowReplacement)
}

// replaceIfUnderpriced replaces the pending transaction [tx] if the fees the node suggests now
// exceed the ones it offers. Otherwise there is nothing better to offer and it is left to be
// mined.
func (t *TxTracker) replaceIfUnderpriced(
	ctx context.Context,
	log logging.Logger,
	tx *types.Transaction,
	allowReplacement func(context.Context) (bool, error),
) error {
	tipCap, feeCap, err := t.suggestFees(ctx)
	if err != nil {
		return err
	}
	if tipCap.Cmp(tx.GasTipCap()) <= 0 && feeCap.Cmp(tx.GasFeeCap()) <= 0 {
		log.Info("Unconfirmed transaction is still pending at current fees, waiting",
			zap.Stringer("gasTipCap", tx.GasTipCap()),
			zap.Stringer("gasFeeCap", tx.GasFeeCap()),
		)
		return nil
	}

	if allowReplacement != nil {
		ok, err := allowReplacement(ctx)
		if err != nil || !ok {
			return err
		}
	}

	tipCap = math.BigMax(tipCap, replacementFee(tx.GasTipCap()))
	feeCap = math.BigMax(feeCap, replacementFee(tx.GasFeeCap()))
	var txData types.TxData
	if tx.Type() == types.DynamicFeeTxType {
		txData = &types.DynamicFeeTx{
			ChainID:    tx.ChainId(),
			Nonce:      tx.Nonce(),
			GasTipCap:  tipCap,
			GasFeeCap:  feeCap,
			Gas:        tx.Gas(),
			To:         tx.To(),
			Value:      tx.Value(),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}
	} else {
		txData = &types.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: feeCap,
			Gas:      tx.Gas(),
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		}
	}
	replacement, err := t.signer(t.from, types.NewTx(txData))
	if err != nil {
		return fmt.Errorf("failed to sign replacement for transaction %s: %w", tx.Hash(), err)
	}
	if err := t.client.SendTransaction(ctx, replacement); err != nil {
		return fmt.Errorf("failed to replace underpriced transaction %s: %w", tx.Hash(), err)
	}

	t.unconfirmed = replacement
	log.Warn("Replaced underpriced transaction",
		zap.Stringer("replacementTxHash", replacement.Hash()),
		zap.Stringer("gasTipCap", tipCap),
		zap.Stringer("gasFeeCap", feeCap),
	)
	return nil
}

// suggestFees returns the tip and fee caps a transaction sent now would carry, priced the way
// bind prices one when the caller sets neither: the node's suggested tip, and a fee cap of that
// tip plus twice the current base fee. On a chain without a base fee both are the node's
// suggested gas price.
func (t *TxTracker) suggestFees(ctx context.Context) (*big.Int, *big.Int, error) {
	head, err := t.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get latest header: %w", err)
	}
	if head.BaseFee == nil {
		gasPrice, err := t.client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get suggested gas price: %w", err)
		}
		return gasPrice, gasPrice, nil
	}
	tipCap, err := t.client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get suggested gas tip cap: %w", err)
	}
	feeCap := new(big.Int).Mul(head.BaseFee, big.NewInt(2))
	return tipCap, feeCap.Add(feeCap, tipCap), nil
}

// replacementFee returns the lowest fee a node accepts from a transaction that replaces one
// offering [fee]: replacementFeeBumpPercent more, rounded up.
func replacementFee(fee *big.Int) *big.Int {
	bumped := new(big.Int).Mul(fee, big.NewInt(100+replacementFeeBumpPercent))
	bumped.Add(bumped, big.NewInt(99))
	return bumped.Div(bumped, big.NewInt(100))
}
