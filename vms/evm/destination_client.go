// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_eth_client.go -package=mocks

package evm

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm/signer"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/rpc"
	"go.uber.org/zap"
)

const (
	// If the max base fee is not explicitly set, use 3x the current base fee estimate
	defaultBaseFeeFactor          = 3
	poolTxsPerAccount             = 16
	pendingTxRefreshInterval      = 2 * time.Second
	defaultBlockAcceptanceTimeout = 30 * time.Second
)

// Client interface wraps the ethclient.Client interface for mocking purposes.
type Client interface {
	ethclient.Client
}

// Implements DestinationClient
type destinationClient struct {
	client ethclient.Client

	readonlyConcurrentSigners []*readonlyConcurrentSigner

	destinationBlockchainID ids.ID
	evmChainID              *big.Int
	blockGasLimit           uint64
	maxBaseFee              *big.Int
	maxPriorityFeePerGas    *big.Int
	logger                  logging.Logger
	txInclusionTimeout      time.Duration
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
	queuedTxSemaphore chan struct{}
	destinationClient *destinationClient
}

type txData struct {
	to            common.Address
	gasLimit      uint64
	gasFeeCap     *big.Int
	gasTipCap     *big.Int
	callData      []byte
	signedMessage *avalancheWarp.Message
	resultChan    chan txResult
}

type txResult struct {
	receipt *types.Receipt
	err     error
}

func NewDestinationClient(
	logger logging.Logger,
	destinationBlockchain *config.DestinationBlockchain,
) (*destinationClient, error) {
	logger = logger.With(zap.String("blockchainID", destinationBlockchain.BlockchainID))

	destinationID, err := ids.FromString(destinationBlockchain.BlockchainID)
	if err != nil {
		logger.Error(
			"Could not decode destination chain ID from string",
			zap.Error(err),
		)
		return nil, err
	}

	signers, err := signer.NewSigners(destinationBlockchain)
	if err != nil {
		logger.Error(
			"Failed to create signer",
			zap.Error(err),
		)
		return nil, err
	}

	// Dial the destination RPC endpoint
	client, err := utils.NewEthClientWithConfig(
		context.Background(),
		destinationBlockchain.RPCEndpoint.BaseURL,
		destinationBlockchain.RPCEndpoint.HTTPHeaders,
		destinationBlockchain.RPCEndpoint.QueryParams,
	)
	if err != nil {
		logger.Error(
			"Failed to dial rpc endpoint",
			zap.Error(err),
		)
		return nil, err
	}

	evmChainID, err := client.ChainID(context.Background())
	if err != nil {
		logger.Error(
			"Failed to get chain ID from destination chain endpoint",
			zap.Error(err),
		)
		return nil, err
	}

	var (
		destClient                 destinationClient
		pendingNonce, currentNonce uint64
		readonlyConcurrentSigners  = make([]*readonlyConcurrentSigner, len(signers))
	)

	// Block until all pending txs are accepted
	ticker := time.NewTicker(pendingTxRefreshInterval)
	defer ticker.Stop()
	for i, signer := range signers {
		for {
			pendingNonce, err = client.NonceAt(context.Background(), signer.Address(), big.NewInt(int64(rpc.PendingBlockNumber)))
			if err != nil {
				logger.Error(
					"Failed to get pending nonce",
					zap.Error(err),
				)
				return nil, err
			}

			currentNonce, err = client.NonceAt(context.Background(), signer.Address(), nil)
			if err != nil {
				logger.Error(
					"Failed to get current nonce",
					zap.Error(err),
				)
				return nil, err
			}

			// If the pending nonce is not equal to the current nonce, wait and check again
			if pendingNonce != currentNonce {
				logger.Info(
					"Waiting for pending txs to be accepted",
					zap.Uint64("pendingNonce", pendingNonce),
					zap.Uint64("currentNonce", currentNonce),
					zap.Stringer("address", signer.Address()),
				)
				<-ticker.C
				continue
			}

			logger.Debug(
				"Pending txs accepted",
				zap.Stringer("address", signer.Address()),
			)

			concurrentSigner := &concurrentSigner{
				logger:            logger.With(zap.Stringer("senderAddress", signer.Address())),
				signer:            signer,
				currentNonce:      currentNonce,
				messageChan:       make(chan txData),
				queuedTxSemaphore: make(chan struct{}, poolTxsPerAccount),
				destinationClient: &destClient,
			}

			go concurrentSigner.processIncomingTransactions()

			readonlyConcurrentSigners[i] = (*readonlyConcurrentSigner)(concurrentSigner)

			break
		}
	}

	logger.Info(
		"Initialized destination client",
		zap.String("evmChainID", evmChainID.String()),
		zap.Uint64("nonce", pendingNonce),
	)

	destClient = destinationClient{
		client:                    client,
		readonlyConcurrentSigners: readonlyConcurrentSigners,
		destinationBlockchainID:   destinationID,
		evmChainID:                evmChainID,
		logger:                    logger,
		blockGasLimit:             destinationBlockchain.BlockGasLimit,
		maxBaseFee:                new(big.Int).SetUint64(destinationBlockchain.MaxBaseFee),
		maxPriorityFeePerGas:      new(big.Int).SetUint64(destinationBlockchain.MaxPriorityFeePerGas),
		txInclusionTimeout:        time.Duration(destinationBlockchain.TxInclusionTimeoutSeconds) * time.Second,
	}

	return &destClient, nil
}

// SendTx constructs, signs, and broadcast a transaction to deliver the given {signedMessage}
// to this chain with the provided {callData}. If the maximum base fee value is not configured, the
// maximum base is calculated as the current base fee multiplied by the default base fee factor.
// The maximum priority fee per gas is set the minimum of the suggested gas tip cap and the configured
// maximum priority fee per gas. The max fee per gas is set to the sum of the max base fee and the
// max priority fee per gas.
func (c *destinationClient) SendTx(
	signedMessage *avalancheWarp.Message,
	deliverers set.Set[common.Address],
	toAddress string,
	gasLimit uint64,
	callData []byte,
) (*types.Receipt, error) {
	// If the max base fee isn't explicitly set, then default to fetching the
	// current base fee estimate and multiply it by `BaseFeeFactor` to allow for
	// an increase prior to the transaction being included in a block.
	var maxBaseFee *big.Int
	if c.maxBaseFee.Cmp(big.NewInt(0)) > 0 {
		maxBaseFee = c.maxBaseFee
	} else {
		// Get the current base fee estimation for the chain.
		baseFeeCtx, baseFeeCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer baseFeeCtxCancel()
		baseFee, err := c.client.EstimateBaseFee(baseFeeCtx)
		if err != nil {
			c.logger.Error(
				"Failed to get base fee",
				zap.Error(err),
			)
			return nil, err
		}
		maxBaseFee = new(big.Int).Mul(baseFee, big.NewInt(defaultBaseFeeFactor))
	}

	// Get the suggested gas tip cap of the network
	gasTipCapCtx, gasTipCapCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer gasTipCapCtxCancel()
	gasTipCap, err := c.client.SuggestGasTipCap(gasTipCapCtx)
	if err != nil {
		c.logger.Error(
			"Failed to get gas tip cap",
			zap.Error(err),
		)
		return nil, err
	}
	if gasTipCap.Cmp(c.maxPriorityFeePerGas) > 0 {
		gasTipCap = c.maxPriorityFeePerGas
	}

	to := common.HexToAddress(toAddress)
	gasFeeCap := new(big.Int).Add(maxBaseFee, gasTipCap)

	resultChan := make(chan txResult)

	messageData := txData{
		to:            to,
		gasLimit:      gasLimit,
		gasFeeCap:     gasFeeCap,
		gasTipCap:     gasTipCap,
		callData:      callData,
		signedMessage: signedMessage,
		resultChan:    resultChan,
	}

	var cases []reflect.SelectCase
	for _, concurrentSigner := range c.readonlyConcurrentSigners {
		signerAddress := concurrentSigner.signer.Address()
		if deliverers.Len() != 0 && !deliverers.Contains(signerAddress) {
			c.logger.Debug(
				"Signer not eligible to deliver message",
				zap.Any("address", signerAddress),
			)
			continue
		}
		c.logger.Debug(
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
	timeout := time.NewTimer(c.txInclusionTimeout + utils.DefaultRPCTimeout)
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
		c.logger.Error(
			"Failed to get transaction receipt",
			zap.Error(result.err),
		)
		return nil, result.err
	}
	c.logger.Debug(
		"Sent transaction",
	)

	return result.receipt, nil
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
			messageData.resultChan <- txResult{
				receipt: nil,
				err:     err,
			}
			close(messageData.resultChan)
		}
	}
}

// issueTransaction sends the transaction, but does not wait for confirmation.
// In order to properly manage the in-memory nonce, this function must not be
// called concurrently for a given concurrentSigner instance.
// Access to this function should be managed by processIncomingTransactions().
func (s *concurrentSigner) issueTransaction(
	data txData,
) error {
	s.logger.Debug(
		"Processing transaction",
		zap.String("to", data.to.String()),
	)

	// Construct the actual transaction to broadcast on the destination chain
	tx := predicateutils.NewPredicateTx(
		s.destinationClient.evmChainID,
		s.currentNonce,
		&data.to,
		data.gasLimit,
		data.gasFeeCap,
		data.gasTipCap,
		big.NewInt(0),
		data.callData,
		types.AccessList{},
		warp.ContractAddress,
		data.signedMessage.Bytes(),
	)

	// Sign and send the transaction on the destination chain
	signedTx, err := s.signer.SignTx(tx, s.destinationClient.evmChainID)
	if err != nil {
		s.logger.Error(
			"Failed to sign transaction",
			zap.Error(err),
		)
		return err
	}

	sendTxCtx, sendTxCtxCancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer sendTxCtxCancel()

	s.logger.Info(
		"Sending transaction",
		zap.String("txID", signedTx.Hash().String()),
		zap.Uint64("nonce", s.currentNonce),
	)

	if err := s.destinationClient.client.SendTransaction(sendTxCtx, signedTx); err != nil {
		s.logger.Error(
			"Failed to send transaction",
			zap.Error(err),
		)
		return err
	}
	s.logger.Info(
		"Sent transaction",
		zap.String("txID", signedTx.Hash().String()),
		zap.Uint64("nonce", s.currentNonce),
	)

	s.currentNonce++

	// We wait for the transaction receipt asynchronously because the transaction has already
	// been accepted by the mempool, so we can send another transaction using the same key
	// while we wait for the receipt of the previous transaction.
	go s.waitForReceipt(signedTx.Hash(), data.resultChan)

	return nil
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
		receipt, err = s.destinationClient.client.TransactionReceipt(callCtx, txHash)
		return err
	}
	err := utils.WithRetriesTimeout(s.logger, operation, s.destinationClient.txInclusionTimeout, "waitForReceipt")
	if err != nil {
		s.logger.Error(
			"Failed to get transaction receipt",
			zap.Error(err),
		)
		resultChan <- txResult{
			receipt: nil,
			err:     err,
		}
		return
	}

	// Release the queued tx slot
	<-s.queuedTxSemaphore

	resultChan <- txResult{
		receipt: receipt,
		err:     nil,
	}
}

func (c *destinationClient) Client() interface{} {
	return c.client
}

func (c *destinationClient) SenderAddresses() []common.Address {
	addresses := make([]common.Address, len(c.readonlyConcurrentSigners))
	for i, concurrentSigner := range c.readonlyConcurrentSigners {
		addresses[i] = concurrentSigner.signer.Address()
	}
	return addresses
}

func (c *destinationClient) DestinationBlockchainID() ids.ID {
	return c.destinationBlockchainID
}

func (c *destinationClient) BlockGasLimit() uint64 {
	return c.blockGasLimit
}
