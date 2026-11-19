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
// seen mined, which in practice means not within the inclusion timeout. The transaction neither
// succeeded nor failed: the tracker keeps it and [TxTracker.Resolve] follows it up.
var ErrTxUnconfirmed = errors.New("transaction broadcast but not seen mined")

// replacementFeeBumpPercent is how much more a transaction must offer than the pending one it
// replaces at the same nonce, in both its fee cap and its tip, for a node to accept it. This is
// the default of geth and of the clients that follow it.
const replacementFeeBumpPercent = 10

// maxInconclusiveLookups is how many calls of [TxTracker.Resolve] in a row may find a transaction
// neither mined nor pending before it is given up as settled. One such answer is not proof. An
// RPC endpoint that fronts several nodes can be answered by one that has not seen a transaction
// the others hold, and a node can report a transaction as mined before it serves the receipt.
const maxInconclusiveLookups = 3

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
	// inconclusiveLookups counts the calls of Resolve in a row that found unconfirmed neither
	// mined nor pending. See maxInconclusiveLookups.
	inconclusiveLookups int
	// unacknowledged is the last replacement of unconfirmed whose send failed. A send can fail
	// after the node has accepted the transaction, for example when the answer is lost or comes
	// too late. The replacement then holds the nonce, and it is unconfirmed that the node no
	// longer knows.
	unacknowledged *types.Transaction
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
// transaction and the error wraps [ErrTxUnconfirmed]. So it does if the wait fails in any other
// way, in which case the error wraps the cause too. If [ctx] ends first, its error is returned
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
	// The transaction is broadcast and was not seen mined. Whatever ended the wait, its outcome
	// is unknown, so it is kept to be followed up.
	t.follow(tx)
	log := t.logger.With(
		zap.Stringer("txHash", tx.Hash()),
		zap.Uint64("nonce", tx.Nonce()),
	)
	if !errors.Is(err, context.DeadlineExceeded) {
		// bind.WaitMined gives up only when its context ends, which with [ctx] still live leaves
		// the inclusion timeout. Should it ever fail for another reason, that reason is reported
		// instead of a timeout that did not happen.
		log.Warn("Wait for transaction to be mined failed", zap.Error(err))
		return nil, fmt.Errorf("%w: %s: %w", ErrTxUnconfirmed, tx.Hash(), err)
	}
	log.Warn("Transaction not mined within the inclusion timeout",
		zap.Duration("inclusionTimeout", t.inclusionTimeout),
	)
	return nil, fmt.Errorf("%w: %s not mined within %s", ErrTxUnconfirmed, tx.Hash(), t.inclusionTimeout)
}

// Unconfirmed returns the transaction whose outcome is still unknown, or nil if there is none.
func (t *TxTracker) Unconfirmed() *types.Transaction {
	return t.unconfirmed
}

// follow makes [tx] the unconfirmed transaction. With nil there is none.
func (t *TxTracker) follow(tx *types.Transaction) {
	t.unconfirmed = tx
	t.inconclusiveLookups = 0
	t.unacknowledged = nil
}

// Resolve finds out what became of the unconfirmed transaction and reports whether it is
// settled, meaning that it can no longer take effect and no longer holds its nonce. With no
// unconfirmed transaction there is nothing to settle and it returns true.
//
//   - Mined, successfully or not: settled.
//   - Still pending: not settled. It is left alone while it still pays the tip the node suggests
//     now, and replaced once it does not.
//   - Neither, because the node does not know the transaction, or reports it as mined and serves
//     no receipt for it: not settled at first, since one such answer is not proof. It is settled
//     once [maxInconclusiveLookups] calls in a row have found nothing else. By then it was
//     dropped from the pool, or another transaction took its nonce, or it is mined and only its
//     receipt is missing.
//
// A replacement is the same call from the same account at the same nonce, priced the way bind
// prices a new transaction and raised to the minimum a node accepts for a replacement. It becomes
// the unconfirmed transaction. It is only made for a transaction that no longer pays the
// suggested tip, never because a freshly priced one would merely carry a higher fee cap, so the
// fees paid follow the market instead of compounding on every call. [allowReplacement], if not
// nil, is asked first and can defer it, for example while gas is above what the sender is
// willing to pay.
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
		t.follow(nil)
		return true, nil
	case !errors.Is(err, ethereum.NotFound):
		return false, fmt.Errorf("failed to get receipt of unconfirmed transaction %s: %w", tx.Hash(), err)
	}

	_, isPending, err := t.client.TransactionByHash(ctx, tx.Hash())
	switch {
	case errors.Is(err, ethereum.NotFound):
		// A replacement the node accepted without the tracker learning of it would have pushed
		// the transaction out of the pool, so that is ruled out before it counts as missing.
		adopted, err := t.adoptUnacknowledged(ctx, log)
		if err != nil || adopted {
			return false, err
		}
		return t.inconclusiveLookup(log, "Unconfirmed transaction is neither mined nor pending"), nil
	case err != nil:
		return false, fmt.Errorf("failed to look up unconfirmed transaction %s: %w", tx.Hash(), err)
	case !isPending:
		return t.inconclusiveLookup(log, "Unconfirmed transaction is reported mined, but has no receipt"), nil
	}

	t.inconclusiveLookups = 0
	return false, t.replaceIfUnderpriced(ctx, log, tx, allowReplacement)
}

// inconclusiveLookup records that a call of Resolve found the unconfirmed transaction neither
// mined nor pending, which [msg] describes, and reports whether that settles it: it does once
// maxInconclusiveLookups calls in a row have found the same.
func (t *TxTracker) inconclusiveLookup(log logging.Logger, msg string) bool {
	t.inconclusiveLookups++
	settled := t.inconclusiveLookups >= maxInconclusiveLookups
	log.Warn(msg,
		zap.Int("lookups", t.inconclusiveLookups),
		zap.Int("settledAfter", maxInconclusiveLookups),
		zap.Bool("settled", settled),
	)
	if settled {
		t.follow(nil)
	}
	return settled
}

// adoptUnacknowledged makes the replacement whose send failed the unconfirmed transaction if the
// node knows it after all, and reports whether it did. A replacement the node does not know is
// kept: the node that answers next may.
func (t *TxTracker) adoptUnacknowledged(ctx context.Context, log logging.Logger) (bool, error) {
	replacement := t.unacknowledged
	if replacement == nil {
		return false, nil
	}
	_, _, err := t.client.TransactionByHash(ctx, replacement.Hash())
	switch {
	case errors.Is(err, ethereum.NotFound):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("failed to look up replacement transaction %s: %w", replacement.Hash(), err)
	}
	t.follow(replacement)
	log.Warn("Node accepted a replacement whose send failed, following it instead",
		zap.Stringer("replacementTxHash", replacement.Hash()),
	)
	return true, nil
}

// replaceIfUnderpriced replaces the pending transaction [tx] if it is underpriced: its tip is
// below the one the node suggests now, or its fee cap no longer pays that tip at the current base
// fee. Otherwise it is left to be mined. A base fee that has risen but is still covered by the
// fee cap is no reason to replace it.
//
// [allowReplacement], if not nil, is asked once [tx] is found underpriced, before anything is
// signed or sent. The replacement goes ahead only if it returns true. If it returns false or an
// error, [tx] is left as it is and the error, if any, is returned.
func (t *TxTracker) replaceIfUnderpriced(
	ctx context.Context,
	log logging.Logger,
	tx *types.Transaction,
	allowReplacement func(context.Context) (bool, error),
) error {
	tipCap, minFeeCap, feeCap, err := t.suggestFees(ctx)
	if err != nil {
		return err
	}
	if tx.GasTipCap().Cmp(tipCap) >= 0 && tx.GasFeeCap().Cmp(minFeeCap) >= 0 {
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
		// The node may hold the replacement all the same, so it is kept to be looked for if
		// [tx] goes missing.
		t.unacknowledged = replacement
		return fmt.Errorf("failed to replace underpriced transaction %s with %s: %w",
			tx.Hash(), replacement.Hash(), err)
	}

	t.follow(replacement)
	log.Warn("Replaced underpriced transaction",
		zap.Stringer("replacementTxHash", replacement.Hash()),
		zap.Stringer("gasTipCap", tipCap),
		zap.Stringer("gasFeeCap", feeCap),
	)
	return nil
}

// suggestFees returns what the node suggests a transaction pay now:
//   - tipCap, the suggested tip;
//   - minFeeCap, the lowest fee cap that pays that tip at the current base fee, which is their
//     sum. A pending transaction offering less is underpriced;
//   - feeCap, the fee cap bind gives a new transaction when the caller sets none: the tip plus
//     twice the current base fee, which leaves room for the base fee to rise before the
//     transaction is mined.
//
// On a chain without a base fee all three are the node's suggested gas price.
func (t *TxTracker) suggestFees(ctx context.Context) (tipCap, minFeeCap, feeCap *big.Int, err error) {
	head, err := t.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get latest header: %w", err)
	}
	if head.BaseFee == nil {
		gasPrice, gasPriceErr := t.client.SuggestGasPrice(ctx)
		if gasPriceErr != nil {
			return nil, nil, nil, fmt.Errorf("failed to get suggested gas price: %w", gasPriceErr)
		}
		return gasPrice, gasPrice, gasPrice, nil
	}
	tipCap, err = t.client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get suggested gas tip cap: %w", err)
	}
	minFeeCap = new(big.Int).Add(head.BaseFee, tipCap)
	feeCap = new(big.Int).Add(minFeeCap, head.BaseFee)
	return tipCap, minFeeCap, feeCap, nil
}

// replacementFee returns the lowest fee a node accepts from a transaction that replaces one
// offering [fee]: replacementFeeBumpPercent more, rounded up.
func replacementFee(fee *big.Int) *big.Int {
	bumped := new(big.Int).Mul(fee, big.NewInt(100+replacementFeeBumpPercent))
	bumped.Add(bumped, big.NewInt(99))
	return bumped.Div(bumped, big.NewInt(100))
}
