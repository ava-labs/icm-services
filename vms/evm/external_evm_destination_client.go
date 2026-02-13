// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	validatorregistry "github.com/ava-labs/icm-services/abi-bindings/go/AvalancheValidatorSetRegistry"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm/client"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

const (
	// externalEVMDefaultBaseFeeFactor is the multiplier for base fee when not explicitly set
	externalEVMDefaultBaseFeeFactor = 3
	// externalEVMPoolTxsPerAccount limits pending txs per account in mempool
	externalEVMPoolTxsPerAccount = 16
	// gasLimitForSimulation is the gas limit used when simulating calls
	gasLimitForSimulation = 2_000_000
)

// ExternalEVMDestinationClient handles communication with external EVM chains
// that have AvalancheValidatorSetRegistry contracts deployed.
//
// Implements vms.DestinationClient interface.
type ExternalEVMDestinationClient struct {
	ethClient       client.EthClient
	logger          logging.Logger
	chainID         string
	evmChainID      *big.Int
	registryAddress common.Address
	rpcEndpointURL  string
	blockGasLimit   uint64

	// Gas fee configuration
	maxBaseFee                 *big.Int
	suggestedPriorityFeeBuffer *big.Int
	maxPriorityFeePerGas       *big.Int
	txInclusionTimeout         time.Duration

	// Concurrent senders for transaction processing
	concurrentSenders []*externalEVMConcurrentSender
}

// externalEVMConcurrentSender handles transaction signing and sending for one private key.
// Each sender has its own goroutine that processes transactions sequentially,
// ensuring proper nonce management while allowing concurrent receipt waiting.
type externalEVMConcurrentSender struct {
	logger            logging.Logger
	privateKey        *ecdsa.PrivateKey
	address           common.Address
	currentNonce      uint64
	messageChan       chan externalEVMTxData
	queuedTxSemaphore chan struct{}
	destClient        *ExternalEVMDestinationClient
}

// externalEVMTxData holds the data for a transaction to be sent
type externalEVMTxData struct {
	to            common.Address
	gasLimit      uint64
	gasFeeCap     *big.Int
	gasTipCap     *big.Int
	callData      []byte
	signedMessage *avalancheWarp.Message
	resultChan    chan externalEVMTxResult
}

// externalEVMTxResult holds the result of a transaction
type externalEVMTxResult struct {
	receipt *types.Receipt
	err     error
	txID    common.Hash
}

// NewExternalEVMDestinationClient creates a new external EVM destination client.
func NewExternalEVMDestinationClient(
	logger logging.Logger,
	chainID string,
	rpcEndpointURL string,
	registryAddress common.Address,
	privateKeyHexes []string,
	blockGasLimit uint64,
	maxBaseFee *big.Int,
	suggestedPriorityFeeBuffer *big.Int,
	maxPriorityFeePerGas *big.Int,
	txInclusionTimeoutSeconds uint64,
) (*ExternalEVMDestinationClient, error) {
	logger = logger.With(
		zap.String("chainID", chainID),
		zap.String("registryAddress", registryAddress.Hex()),
	)

	// Parse chain ID
	evmChainID, ok := new(big.Int).SetString(chainID, 10)
	if !ok {
		return nil, fmt.Errorf("invalid chain ID: %s", chainID)
	}

	// Create ethclient connection using libevm/ethclient for external EVM compatibility
	rawClient, err := ethclient.Dial(rpcEndpointURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc endpoint: %w", err)
	}

	// Wrap the client to add Avalanche-specific method stubs
	wrappedClient := client.NewExternalEthClientWrapper(rawClient)

	// Verify chain ID matches
	networkChainID, err := rawClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from endpoint: %w", err)
	}
	if networkChainID.Cmp(evmChainID) != 0 {
		return nil, fmt.Errorf("chain ID mismatch: expected %s, got %s", chainID, networkChainID.String())
	}

	destClient := &ExternalEVMDestinationClient{
		ethClient:                  wrappedClient,
		logger:                     logger,
		chainID:                    chainID,
		evmChainID:                 evmChainID,
		registryAddress:            registryAddress,
		rpcEndpointURL:             rpcEndpointURL,
		blockGasLimit:              blockGasLimit,
		maxBaseFee:                 maxBaseFee,
		suggestedPriorityFeeBuffer: suggestedPriorityFeeBuffer,
		maxPriorityFeePerGas:       maxPriorityFeePerGas,
		txInclusionTimeout:         time.Duration(txInclusionTimeoutSeconds) * time.Second,
	}

	// Initialize concurrent senders from private keys
	concurrentSenders := make([]*externalEVMConcurrentSender, len(privateKeyHexes))
	for i, pkHex := range privateKeyHexes {
		privateKey, err := crypto.HexToECDSA(pkHex)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key %d: %w", i, err)
		}

		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		senderLogger := logger.With(zap.Stringer("senderAddress", address))

		// Get current nonce for this sender
		nonce, err := wrappedClient.NonceAt(context.Background(), address, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce for sender %s: %w", address.Hex(), err)
		}

		cs := &externalEVMConcurrentSender{
			logger:            senderLogger,
			privateKey:        privateKey,
			address:           address,
			currentNonce:      nonce,
			messageChan:       make(chan externalEVMTxData),
			queuedTxSemaphore: make(chan struct{}, externalEVMPoolTxsPerAccount),
			destClient:        destClient,
		}

		// Start the transaction processing goroutine
		go cs.processIncomingTransactions()

		concurrentSenders[i] = cs
		senderLogger.Info("Initialized concurrent sender", zap.Uint64("nonce", nonce))
	}

	destClient.concurrentSenders = concurrentSenders

	logger.Info("Created external EVM destination client",
		zap.String("chainID", chainID),
		zap.String("registryAddress", registryAddress.Hex()),
		zap.String("rpcEndpointURL", rpcEndpointURL),
		zap.Int("numSenders", len(privateKeyHexes)),
	)

	return destClient, nil
}

// getFeePerGas calculates the gas fee cap and gas tip cap for transactions.
func (c *ExternalEVMDestinationClient) getFeePerGas() (*big.Int, *big.Int, error) {
	var maxBaseFee *big.Int
	if c.maxBaseFee != nil && c.maxBaseFee.Cmp(big.NewInt(0)) > 0 {
		maxBaseFee = c.maxBaseFee
	} else {
		// Get base fee from the latest block header
		baseFeeCtx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()

		header, err := c.ethClient.HeaderByNumber(baseFeeCtx, nil) // nil = latest block
		if err != nil {
			c.logger.Error("Failed to get latest block header", zap.Error(err))
			return nil, nil, err
		}
		if header.BaseFee == nil {
			c.logger.Error("Chain does not support EIP-1559 (no base fee in header)")
			return nil, nil, errors.New("chain does not support EIP-1559")
		}
		maxBaseFee = new(big.Int).Mul(header.BaseFee, big.NewInt(externalEVMDefaultBaseFeeFactor))
	}

	// Get suggested gas tip cap
	gasTipCapCtx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer cancel()

	gasTipCap, err := c.ethClient.SuggestGasTipCap(gasTipCapCtx)
	if err != nil {
		c.logger.Error("Failed to get gas tip cap", zap.Error(err))
		return nil, nil, err
	}

	// Add buffer to suggested tip
	if c.suggestedPriorityFeeBuffer != nil {
		gasTipCap = new(big.Int).Add(gasTipCap, c.suggestedPriorityFeeBuffer)
	}

	// Cap at max priority fee
	if c.maxPriorityFeePerGas != nil && gasTipCap.Cmp(c.maxPriorityFeePerGas) > 0 {
		gasTipCap = c.maxPriorityFeePerGas
	}

	gasFeeCap := new(big.Int).Add(maxBaseFee, gasTipCap)
	return gasFeeCap, gasTipCap, nil
}

// SendTx sends a transaction to an external EVM chain.
// Uses channel-based concurrency for nonce management.
func (c *ExternalEVMDestinationClient) SendTx(
	signedMessage *avalancheWarp.Message,
	deliverers set.Set[common.Address],
	toAddress string,
	gasLimit uint64,
	callData []byte,
) (*types.Receipt, error) {
	gasFeeCap, gasTipCap, err := c.getFeePerGas()
	if err != nil {
		return nil, err
	}

	resultChan := make(chan externalEVMTxResult)
	to := common.HexToAddress(toAddress)
	txData := externalEVMTxData{
		to:            to,
		gasLimit:      gasLimit,
		gasFeeCap:     gasFeeCap,
		gasTipCap:     gasTipCap,
		callData:      callData,
		signedMessage: signedMessage,
		resultChan:    resultChan,
	}

	// Build select cases for eligible senders
	var cases []reflect.SelectCase
	for _, cs := range c.concurrentSenders {
		if deliverers.Len() != 0 && !deliverers.Contains(cs.address) {
			c.logger.Debug("Sender not eligible to deliver message",
				zap.Stringer("address", cs.address))
			continue
		}
		c.logger.Debug("Sender eligible to deliver message",
			zap.Stringer("address", cs.address))
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(cs.messageChan),
			Send: reflect.ValueOf(txData),
		})
	}

	if len(cases) == 0 {
		return nil, errors.New("no eligible senders available")
	}

	// Select an available, eligible sender
	reflect.Select(cases)

	// Wait for result with timeout
	timeout := time.NewTimer(c.txInclusionTimeout + utils.DefaultRPCTimeout)
	defer timeout.Stop()

	select {
	case result, ok := <-resultChan:
		if !ok {
			return nil, errors.New("result channel closed unexpectedly")
		}
		if result.err != nil {
			c.logger.Error("Transaction failed",
				zap.Error(result.err),
				zap.Stringer("txID", result.txID))
			return nil, result.err
		}
		return result.receipt, nil
	case <-timeout.C:
		return nil, errors.New("timed out waiting for transaction result")
	}
}

// processIncomingTransactions is the worker goroutine for each sender.
// It processes transactions sequentially to ensure proper nonce management.
func (s *externalEVMConcurrentSender) processIncomingTransactions() {
	for {
		// Acquire semaphore slot (blocks if too many pending txs)
		s.queuedTxSemaphore <- struct{}{}
		s.logger.Debug("Waiting for incoming transaction")

		txData := <-s.messageChan

		err := s.issueTransaction(txData)
		if err != nil {
			s.logger.Error("Failed to issue transaction", zap.Error(err))
			// Release semaphore and send error result
			<-s.queuedTxSemaphore
			txData.resultChan <- externalEVMTxResult{
				receipt: nil,
				err:     err,
			}
			close(txData.resultChan)
		}
	}
}

// issueTransaction sends a transaction to the external EVM.
// Must be called sequentially per sender (managed by processIncomingTransactions).
// The callData should be pre-encoded by the caller using ABI bindings.
func (s *externalEVMConcurrentSender) issueTransaction(data externalEVMTxData) error {
	s.logger.Debug("Processing transaction", zap.Stringer("to", data.to))

	// Create the transaction with pre-encoded callData
	// The callData is packed by the caller (e.g., registryABI.Pack("updateValidatorSet", ...))
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.destClient.evmChainID,
		Nonce:     s.currentNonce,
		To:        &data.to,
		Gas:       data.gasLimit,
		GasFeeCap: data.gasFeeCap,
		GasTipCap: data.gasTipCap,
		Value:     big.NewInt(0),
		Data:      data.callData, // Pre-encoded by caller
	})

	// Create signer and sign the transaction
	signer := types.LatestSignerForChainID(s.destClient.evmChainID)
	signedTx, err := types.SignTx(tx, signer, s.privateKey)
	if err != nil {
		s.logger.Error("Failed to sign transaction", zap.Error(err))
		return err
	}

	// Send the transaction
	sendCtx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
	defer cancel()

	s.logger.Info("Sending transaction to external EVM",
		zap.Stringer("txID", signedTx.Hash()),
		zap.Stringer("from", s.address),
		zap.Stringer("to", data.to),
		zap.Uint64("nonce", s.currentNonce),
		zap.Uint64("gasLimit", data.gasLimit),
		zap.Stringer("gasFeeCap", data.gasFeeCap),
		zap.Stringer("gasTipCap", data.gasTipCap),
	)

	if err := s.destClient.ethClient.SendTransaction(sendCtx, signedTx); err != nil {
		s.logger.Error("Failed to send transaction", zap.Error(err))
		return err
	}

	s.logger.Info("Transaction sent successfully", zap.Stringer("txID", signedTx.Hash()))
	s.currentNonce++

	// Wait for receipt asynchronously
	go s.waitForReceipt(signedTx.Hash(), data.resultChan)

	return nil
}

// waitForReceipt polls for transaction receipt and sends result to channel.
func (s *externalEVMConcurrentSender) waitForReceipt(
	txHash common.Hash,
	resultChan chan<- externalEVMTxResult,
) {
	defer close(resultChan)

	var receipt *types.Receipt
	operation := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DefaultRPCTimeout)
		defer cancel()

		var err error
		receipt, err = s.destClient.ethClient.TransactionReceipt(ctx, txHash)
		return err
	}

	notify := func(err error, duration time.Duration) {
		s.logger.Debug("Waiting for receipt, retrying...",
			zap.Stringer("txID", txHash),
			zap.Duration("retryIn", duration),
			zap.Error(err))
	}

	err := utils.WithRetriesTimeout(operation, notify, s.destClient.txInclusionTimeout)
	if err != nil {
		resultChan <- externalEVMTxResult{
			receipt: nil,
			err:     fmt.Errorf("failed to get transaction receipt: %w", err),
			txID:    txHash,
		}
		return
	}

	<-s.queuedTxSemaphore

	resultChan <- externalEVMTxResult{
		receipt: receipt,
		err:     nil,
		txID:    txHash,
	}
}

// SimulateCall simulates a contract call and returns the revert reason if it fails.
// This is useful for debugging why a transaction might fail before sending it.
func (c *ExternalEVMDestinationClient) SimulateCall(
	ctx context.Context,
	toAddress string,
	callData []byte,
) ([]byte, error) {
	to := common.HexToAddress(toAddress)

	// Use the first sender's address as the "from" address for simulation
	var fromAddress common.Address
	if len(c.concurrentSenders) > 0 {
		fromAddress = c.concurrentSenders[0].address
	}

	callMsg := ethereum.CallMsg{
		From: fromAddress,
		To:   &to,
		Gas:  gasLimitForSimulation, // Use same gas limit as actual transactions
		Data: callData,
	}

	result, err := c.ethClient.CallContract(ctx, callMsg, nil) // nil = latest block
	if err != nil {
		// Try to extract revert reason from error
		c.logger.Debug("SimulateCall error details",
			zap.Error(err),
			zap.String("errorType", fmt.Sprintf("%T", err)),
		)
	}
	return result, err
}

// SimulateCallAtBlock simulates a contract call at a specific block number.
func (c *ExternalEVMDestinationClient) SimulateCallAtBlock(
	ctx context.Context,
	toAddress string,
	callData []byte,
	blockNumber *big.Int,
) ([]byte, error) {
	to := common.HexToAddress(toAddress)

	var fromAddress common.Address
	if len(c.concurrentSenders) > 0 {
		fromAddress = c.concurrentSenders[0].address
	}

	callMsg := ethereum.CallMsg{
		From: fromAddress,
		To:   &to,
		Gas:  gasLimitForSimulation,
		Data: callData,
	}

	result, err := c.ethClient.CallContract(ctx, callMsg, blockNumber)
	return result, err
}

// GetPChainHeightForDestination queries the registry contract for its known P-chain height.
func (c *ExternalEVMDestinationClient) GetPChainHeightForDestination(
	ctx context.Context,
) (uint64, error) {
	c.logger.Debug("Querying registry for P-chain height",
		zap.String("registryAddress", c.registryAddress.Hex()))

	// Get current validator set to find its P-chain height
	registryABI, err := validatorregistry.AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return 0, fmt.Errorf("failed to get registry ABI: %w", err)
	}

	// Check if any validator set exists
	nextID, err := c.GetNextValidatorSetID(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get next validator set ID: %w", err)
	}
	if nextID == 0 {
		// No validator sets registered yet
		return 0, nil
	}

	// Get current validator set ID
	currentID, err := c.GetCurrentValidatorSetID(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get current validator set ID: %w", err)
	}

	// Call getValidatorSet to get the validator set details
	callData, err := registryABI.Pack("getValidatorSet", new(big.Int).SetUint64(currentID))
	if err != nil {
		return 0, fmt.Errorf("failed to pack getValidatorSet call: %w", err)
	}

	result, err := c.ethClient.CallContract(ctx, ethereum.CallMsg{
		To:   &c.registryAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to call getValidatorSet: %w", err)
	}

	// Unpack the result - ValidatorSet struct has pChainHeight at index 3
	unpacked, err := registryABI.Unpack("getValidatorSet", result)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getValidatorSet result: %w", err)
	}

	// The ABI unpacker returns anonymous structs, so we need to use reflection
	// to extract the pChainHeight field
	if len(unpacked) > 0 {
		// Use type assertion with anonymous struct that matches ABI unpacker output
		validatorSet := unpacked[0].(struct {
			AvalancheBlockchainID [32]byte `json:"avalancheBlockchainID"`
			Validators            []struct {
				BlsPublicKey []uint8 `json:"blsPublicKey"`
				Weight       uint64  `json:"weight"`
			} `json:"validators"`
			TotalWeight     uint64 `json:"totalWeight"`
			PChainHeight    uint64 `json:"pChainHeight"`
			PChainTimestamp uint64 `json:"pChainTimestamp"`
		})
		return validatorSet.PChainHeight, nil
	}

	return 0, fmt.Errorf("unexpected result format from getValidatorSet")
}

// GetNextValidatorSetID queries the registry contract for the next validator set ID.
// If this returns 0, no validator sets have been registered yet.
func (c *ExternalEVMDestinationClient) GetNextValidatorSetID(ctx context.Context) (uint32, error) {
	registryABI, err := validatorregistry.AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return 0, fmt.Errorf("failed to get registry ABI: %w", err)
	}

	callData, err := registryABI.Pack("nextValidatorSetID")
	if err != nil {
		return 0, fmt.Errorf("failed to pack nextValidatorSetID call: %w", err)
	}

	result, err := c.ethClient.CallContract(ctx, ethereum.CallMsg{
		To:   &c.registryAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to call nextValidatorSetID: %w", err)
	}

	unpacked, err := registryABI.Unpack("nextValidatorSetID", result)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack nextValidatorSetID result: %w", err)
	}

	if len(unpacked) > 0 {
		return unpacked[0].(uint32), nil
	}

	return 0, fmt.Errorf("unexpected result format from nextValidatorSetID")
}

// GetCurrentValidatorSetID queries the registry contract for the current validator set ID.
func (c *ExternalEVMDestinationClient) GetCurrentValidatorSetID(ctx context.Context) (uint64, error) {
	registryABI, err := validatorregistry.AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return 0, fmt.Errorf("failed to get registry ABI: %w", err)
	}

	callData, err := registryABI.Pack("getCurrentValidatorSetID")
	if err != nil {
		return 0, fmt.Errorf("failed to pack getCurrentValidatorSetID call: %w", err)
	}

	result, err := c.ethClient.CallContract(ctx, ethereum.CallMsg{
		To:   &c.registryAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to call getCurrentValidatorSetID: %w", err)
	}

	unpacked, err := registryABI.Unpack("getCurrentValidatorSetID", result)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getCurrentValidatorSetID result: %w", err)
	}

	if len(unpacked) > 0 {
		// Returns *big.Int, convert to uint64
		return unpacked[0].(*big.Int).Uint64(), nil
	}

	return 0, fmt.Errorf("unexpected result format from getCurrentValidatorSetID")
}

// Client returns the underlying ethclient.
func (c *ExternalEVMDestinationClient) Client() Client {
	return c.ethClient
}

// SenderAddresses returns the addresses of all senders.
func (c *ExternalEVMDestinationClient) SenderAddresses() []common.Address {
	addresses := make([]common.Address, len(c.concurrentSenders))
	for i, cs := range c.concurrentSenders {
		addresses[i] = cs.address
	}
	return addresses
}

// DestinationBlockchainID returns empty for external chains.
// External chains don't have Avalanche blockchain IDs.
// This method is required by the interface but not used for external EVMs.
func (c *ExternalEVMDestinationClient) DestinationBlockchainID() ids.ID {
	return ids.Empty
}

// BlockGasLimit returns the configured gas limit for transactions.
func (c *ExternalEVMDestinationClient) BlockGasLimit() uint64 {
	return c.blockGasLimit
}

// GetRPCEndpointURL returns the RPC endpoint URL for this external chain.
func (c *ExternalEVMDestinationClient) GetRPCEndpointURL() string {
	return c.rpcEndpointURL
}

// RegistryAddress returns the registry contract address for this external chain.
func (c *ExternalEVMDestinationClient) RegistryAddress() common.Address {
	return c.registryAddress
}
