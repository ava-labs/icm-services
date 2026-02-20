// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	validatorregistry "github.com/ava-labs/icm-services/abi-bindings/go/SubsetUpdater"
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

type PrivateKeySigner struct {
	privateKey *ecdsa.PrivateKey
}

func (p *PrivateKeySigner) Address() common.Address {
	return crypto.PubkeyToAddress(p.privateKey.PublicKey)
}

func (p *PrivateKeySigner) SignTx(tx *types.Transaction, evmChainID *big.Int) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(evmChainID)
	return types.SignTx(tx, signer, p.privateKey)
}

// ExternalEVMDestinationClient handles communication with external EVM chains
// that have AvalancheValidatorSetRegistry contracts deployed.
//
// Implements vms.DestinationClient interface.
type ExternalEVMDestinationClient struct {
	ethClient       EthClient
	logger          logging.Logger
	chainID         string
	evmChainID      *big.Int
	registryAddress common.Address
	rpcEndpointURL  string
	blockGasLimit   uint64

	// Gas fee configuration
	gasFeeConfig       *GasFeeConfig
	txInclusionTimeout time.Duration

	// Concurrent senders for transaction processing
	concurrentSenders []*readonlyConcurrentSigner
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
	wrappedClient := NewExternalEthClientWrapper(rawClient)

	// Verify chain ID matches
	networkChainID, err := rawClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from endpoint: %w", err)
	}
	if networkChainID.Cmp(evmChainID) != 0 {
		return nil, fmt.Errorf("chain ID mismatch: expected %s, got %s", chainID, networkChainID.String())
	}
	gasFeeData := &GasFeeConfig{
		maxBaseFee:                 maxBaseFee,
		suggestedPriorityFeeBuffer: suggestedPriorityFeeBuffer,
		maxPriorityFeePerGas:       maxPriorityFeePerGas,
	}
	destClient := &ExternalEVMDestinationClient{
		ethClient:          wrappedClient,
		logger:             logger,
		chainID:            chainID,
		evmChainID:         evmChainID,
		registryAddress:    registryAddress,
		rpcEndpointURL:     rpcEndpointURL,
		blockGasLimit:      blockGasLimit,
		gasFeeConfig:       gasFeeData,
		txInclusionTimeout: time.Duration(txInclusionTimeoutSeconds) * time.Second,
	}

	// Initialize concurrent senders from private keys
	concurrentSenders := make([]*readonlyConcurrentSigner, len(privateKeyHexes))
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

		cs := &concurrentSigner{
			logger:            senderLogger,
			signer:            &PrivateKeySigner{privateKey: privateKey},
			currentNonce:      nonce,
			messageChan:       make(chan txData),
			queuedTxSemaphore: make(chan struct{}, externalEVMPoolTxsPerAccount),
			destinationClient: destClient,
		}

		// Start the transaction processing goroutine
		go cs.processIncomingTransactions()

		concurrentSenders[i] = (*readonlyConcurrentSigner)(cs)
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

func (c *ExternalEVMDestinationClient) EVMChainID() *big.Int {
	return c.evmChainID
}

func (c *ExternalEVMDestinationClient) RPCClient() DestinationRPCClient {
	return c.ethClient
}

func (c *ExternalEVMDestinationClient) Logger() logging.Logger {
	return c.logger
}

func (c *ExternalEVMDestinationClient) GasFeeConfig() *GasFeeConfig {
	return c.gasFeeConfig
}

func (c *ExternalEVMDestinationClient) FeeFactor() int64 {
	return externalEVMDefaultBaseFeeFactor
}

func (c *ExternalEVMDestinationClient) ConcurrentSigners() []*readonlyConcurrentSigner {
	return c.concurrentSenders
}

func (c *ExternalEVMDestinationClient) AccessList(_ txData) types.AccessList {
	return types.AccessList{}
}

func (c *ExternalEVMDestinationClient) TxInclusionTimeout() time.Duration {
	return c.txInclusionTimeout
}

// getFeePerGas calculates the gas fee cap and gas tip cap for transactions.
// nolint:unused
func (c *ExternalEVMDestinationClient) getFeePerGas() (*big.Int, *big.Int, error) {
	return getFeePerGas(c)
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
	return SendTx(c, signedMessage, deliverers, toAddress, gasLimit, callData, c.txInclusionTimeout)
}

// SenderAddresses returns the addresses of all senders.
func (c *ExternalEVMDestinationClient) SenderAddresses() []common.Address {
	return SenderAddresses(c)
}

// Client returns the underlying ethclient.
func (c *ExternalEVMDestinationClient) Client() Client {
	return c.ethClient
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

// GetPChainHeightForDestination queries the registry contract for its known P-chain height.
func (c *ExternalEVMDestinationClient) GetPChainHeightForDestination(
	ctx context.Context,
) (uint64, error) {
	c.logger.Debug("Querying registry for P-chain height",
		zap.String("registryAddress", c.registryAddress.Hex()))

	// Get the current validator set to find its P-chain height
	registryABI, err := validatorregistry.SubsetUpdaterMetaData.GetAbi()
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
	registryABI, err := validatorregistry.SubsetUpdaterMetaData.GetAbi()
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
	registryABI, err := validatorregistry.SubsetUpdaterMetaData.GetAbi()
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
		fromAddress = c.concurrentSenders[0].signer.Address()
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
		fromAddress = c.concurrentSenders[0].signer.Address()
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
