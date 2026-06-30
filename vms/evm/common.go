// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -destination=./mocks/mock_destination_rpc_client.go -package=mocks github.com/ava-labs/icm-services/vms/evm DestinationRPCClient
package evm

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm/signer"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"go.uber.org/zap"
)

// DestionationRPCClient interface represents the minimal interface needed for querying RPC endpoints.
type DestinationRPCClient interface {
	BlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error)
	ChainID(ctx context.Context) (*big.Int, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BlockNumber(ctx context.Context) (uint64, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	EstimateBaseFee(ctx context.Context) (*big.Int, error)
}

type txData struct {
	to         common.Address
	gasLimit   uint64
	gasFeeCap  *big.Int
	gasTipCap  *big.Int
	chainID    *big.Int
	callData   []byte
	accessList types.AccessList
	resultChan chan txResult
}

type txResult struct {
	receipt *types.Receipt
	err     error
	txID    common.Hash
}

type GasFeeConfig struct {
	maxBaseFee                 *big.Int
	suggestedPriorityFeeBuffer *big.Int
	maxPriorityFeePerGas       *big.Int
}

// Type alias for the destinationClient to have access to the fields but not the methods of the concurrentSigner.
type readonlyConcurrentSigner concurrentSigner

type concurrentSigner struct {
	logger       logging.Logger
	signer       signer.Signer
	currentNonce uint64
	// Unbuffered channel to receive messages to be processed
	messageChan chan txData
	// Semaphore to limit the number of transactions in the mempool for
	// each account, otherwise they may be dropped.
	queuedTxSemaphore  chan struct{}
	txInclusionTimeout time.Duration
	destinationClient  DestinationRPCClient
}

// processIncomingTransactions is a worker that issues transactions from a given concurrentSigner.
// Must be called at most once per concurrentSigner.
// It guarantees that for any messageData read from s.messageChan,
// exactly 1 value is written to messageData.resultChan.
func (s *concurrentSigner) processIncomingTransactions() {
	for {
		// We can only get to listen to messageChan if there is an open queued tx slot
		s.queuedTxSemaphore <- struct{}{}
		s.logger.Debug("Waiting for incoming transaction")

		messageData := <-s.messageChan

		err := s.issueTransaction(messageData)
		if err != nil {
			s.logger.Error(
				"Failed to issue transaction",
				zap.Error(err),
			)
			// If issueTransaction fails, we have not passed the resultChan to waitForReceipt
			// so we need to release the semaphore slot here and send the error result
			<-s.queuedTxSemaphore
			messageData.resultChan <- txResult{
				receipt: nil,
				err:     err,
			}
			close(messageData.resultChan)
		}
	}
}

// maxNonceResyncAttempts bounds how many times issueTransaction will resync the
// cached nonce from the chain and retry when a send fails with "nonce too low".
const maxNonceResyncAttempts = 3

// issueTransaction sends the transaction but does not wait for confirmation.
// In order to properly manage the in-memory nonce, this function must not be
// called concurrently for a given concurrentSigner instance.
// Access to this function should be managed by processIncomingTransactions().
func (s *concurrentSigner) issueTransaction(
	data txData,
) error {
	s.logger.Debug(
		"Processing transaction",
		zap.Stringer("to", data.to),
	)

	var signedTx *types.Transaction
	for attempt := 0; ; attempt++ {
		// Create a standard EIP-1559 transaction with the predicate access list
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:    data.chainID,
			Nonce:      s.currentNonce,
			To:         &data.to,
			Gas:        data.gasLimit,
			GasFeeCap:  data.gasFeeCap,
			GasTipCap:  data.gasTipCap,
			Value:      big.NewInt(0),
			Data:       data.callData,
			AccessList: data.accessList,
		})

		// Sign and send the transaction on the destination chain
		var err error
		signedTx, err = s.signer.SignTx(tx, data.chainID)
		if err != nil {
			s.logger.Error(
				"Failed to sign transaction",
				zap.Error(err),
			)
			return err
		}

		log := s.logger.With(
			zap.Stringer("txID", signedTx.Hash()),
			zap.Uint64("gasLimit", data.gasLimit),
			zap.Stringer("from", s.signer.Address()),
			zap.Stringer("to", data.to),
			zap.Stringer("gasFeeCap", data.gasFeeCap),
			zap.Stringer("gasTipCap", data.gasTipCap),
			zap.Uint64("nonce", s.currentNonce),
		)

		log.Info("Sending transaction")

		sendTxCtx, sendTxCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		err = s.destinationClient.SendTransaction(sendTxCtx, signedTx)
		sendTxCtxCancel()
		if err == nil {
			log.Info("Sent transaction")
			break
		}

		// The cached nonce can fall behind the chain when another component sharing this
		// account (e.g. the Merkle validator-set updater submitting registerValidatorSet)
		// consumes a nonce out-of-band. Resync the nonce from the chain and retry rather
		// than failing the delivery outright.
		if isNonceTooLowError(err) && attempt < maxNonceResyncAttempts {
			log.Warn("Nonce too low, resyncing nonce from chain and retrying", zap.Error(err))
			if resyncErr := s.resyncNonce(); resyncErr != nil {
				log.Error("Failed to resync nonce from chain", zap.Error(resyncErr))
				return err
			}
			continue
		}

		log.Error(
			"Failed to send transaction",
			zap.Error(err),
		)
		return err
	}

	s.currentNonce++

	// We wait for the transaction receipt asynchronously because the transaction has already
	// been accepted by the mempool, so we can send another transaction using the same key
	// while we wait for the receipt of the previous transaction.
	go s.waitForReceipt(signedTx.Hash(), data.resultChan)

	return nil
}

// resyncNonce refreshes the cached nonce from the latest mined state on the chain. It is
// used to recover from "nonce too low" errors that arise when the account is also used by
// another component (such as the validator-set updater).
func (s *concurrentSigner) resyncNonce() error {
	ctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer cancel()
	nonce, err := s.destinationClient.NonceAt(ctx, s.signer.Address(), nil)
	if err != nil {
		return err
	}
	s.logger.Info(
		"Resynced nonce from chain",
		zap.Uint64("previousNonce", s.currentNonce),
		zap.Uint64("newNonce", nonce),
	)
	s.currentNonce = nonce
	return nil
}

// isNonceTooLowError reports whether err indicates the submitted transaction's nonce was
// lower than the account's next expected nonce.
func isNonceTooLowError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "nonce too low")
}

// waitForReceipt always writes to the result channel,
// always closes the result channel,
// may be called concurrently on a given concurrentSigner instance
func (s *concurrentSigner) waitForReceipt(
	txHash common.Hash,
	resultChan chan<- txResult,
) {
	defer close(resultChan)

	var receipt *types.Receipt
	operation := func() (err error) {
		callCtx, callCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer callCtxCancel()
		receipt, err = s.destinationClient.TransactionReceipt(callCtx, txHash)
		return err
	}
	notify := func(err error, duration time.Duration) {
		s.logger.Info(
			"waiting for receipt failed, retrying...",
			zap.Stringer("txID", txHash),
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}

	err := utils.WithRetriesTimeout(operation, notify, s.txInclusionTimeout)
	if err != nil {
		resultChan <- txResult{
			receipt: nil,
			err:     fmt.Errorf("failed to get transaction receipt: %w", err),
			txID:    txHash,
		}
		return
	}

	// Release the queued tx slot
	<-s.queuedTxSemaphore

	resultChan <- txResult{
		receipt: receipt,
		err:     nil,
		txID:    txHash,
	}
}

// getFeePerGas returns the gas fee cap and gas tip cap for the destination chain.
// If the maximum base fee value is not configured, the maximum base is calculated as the current base
// fee multiplied by the default base fee factor. The maximum priority fee per gas is set the minimum
// of the suggested gas tip cap plus the configured suggested priority fee buffer and the configured
// maximum priority fee per gas. The max fee per gas is set to the sum of the max base fee and the
// max priority fee per gas.
func getFeePerGas(
	client DestinationRPCClient,
	gasFeeConfig *GasFeeConfig,
) (*big.Int, *big.Int, error) {
	// If the max base fee isn't explicitly set, then default to fetching the
	// current base fee estimate and multiply it by `defaultMaxBaseFee` to allow for
	// an increase prior to the transaction being included in a block.
	var maxBaseFee *big.Int
	if gasFeeConfig.maxBaseFee.Cmp(big.NewInt(0)) > 0 {
		maxBaseFee = gasFeeConfig.maxBaseFee
	} else {
		// Get the current base fee estimation for the chain.
		baseFeeCtx, baseFeeCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer baseFeeCtxCancel()
		baseFee, err := client.EstimateBaseFee(baseFeeCtx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get base fee: %w", err)
		}
		maxBaseFee = new(big.Int).Mul(baseFee, big.NewInt(defaultBaseFeeFactor))
	}

	// Get the suggested gas tip cap of the network
	gasTipCapCtx, gasTipCapCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer gasTipCapCtxCancel()
	gasTipCap, err := client.SuggestGasTipCap(gasTipCapCtx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get gas tip cap: %w", err)
	}
	gasTipCap = new(big.Int).Add(gasTipCap, gasFeeConfig.suggestedPriorityFeeBuffer)
	if gasTipCap.Cmp(gasFeeConfig.maxPriorityFeePerGas) > 0 {
		gasTipCap = gasFeeConfig.maxPriorityFeePerGas
	}

	gasFeeCap := new(big.Int).Add(maxBaseFee, gasTipCap)

	return gasFeeCap, gasTipCap, nil
}

func SendTx(
	logger logging.Logger,
	client DestinationRPCClient,
	gasFeeConfig *GasFeeConfig,
	concurrentSigners []*readonlyConcurrentSigner,
	accessList types.AccessList,
	chainID *big.Int,
	deliverers set.Set[common.Address],
	toAddress common.Address,
	gasLimit uint64,
	callData []byte,
	txInclusionTimeout time.Duration,
) (*types.Receipt, error) {
	gasFeeCap, gasTipCap, err := getFeePerGas(client, gasFeeConfig)
	if err != nil {
		logger.Error("Failed to calculate gas fee", zap.Error(err))
		return nil, err
	}

	resultChan := make(chan txResult)
	messageData := txData{
		to:         toAddress,
		gasLimit:   gasLimit,
		gasFeeCap:  gasFeeCap,
		gasTipCap:  gasTipCap,
		chainID:    chainID,
		callData:   callData,
		accessList: accessList,
		resultChan: resultChan,
	}

	var cases []reflect.SelectCase
	for _, concurrentSigner := range concurrentSigners {
		signerAddress := concurrentSigner.signer.Address()
		if deliverers.Len() != 0 && !deliverers.Contains(signerAddress) {
			logger.Debug(
				"Signer not eligible to deliver message",
				zap.Any("address", signerAddress),
			)
			continue
		}
		logger.Debug(
			"Signer eligible to deliver message",
			zap.Any("address", signerAddress),
		)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(concurrentSigner.messageChan),
			Send: reflect.ValueOf(messageData),
		})
	}

	// Select an available, eligible signer
	reflect.Select(cases)

	// Wait for the receipt or error to be returned
	// We need to wait for the transaction inclusion, and also the receipt to be returned.
	timeout := time.NewTimer(txInclusionTimeout + utils.DefaultRPCTimeout)
	defer timeout.Stop()
	var result txResult
	var ok bool

	select {
	case result, ok = <-resultChan:
		if !ok {
			return nil, errors.New("channel closed unexpectedly")
		}
	case <-timeout.C:
		return nil, errors.New("timed out waiting for transaction result")
	}

	if result.err != nil {
		logger.Error(
			"Transaction failed to be issued or confirmed",
			zap.Error(result.err),
			zap.Stringer("txID", result.txID),
		)
		return nil, result.err
	}

	return result.receipt, nil
}

func SenderAddresses(concurrentSigners []*readonlyConcurrentSigner) []common.Address {
	addresses := make([]common.Address, len(concurrentSigners))
	for i, concurrentSigner := range concurrentSigners {
		addresses[i] = concurrentSigner.signer.Address()
	}
	return addresses
}
