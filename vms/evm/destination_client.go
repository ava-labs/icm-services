// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_eth_client.go -package=mocks

package evm

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/rpc"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	predicateutils "github.com/ava-labs/avalanchego/vms/evm/predicate"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm/signer"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

const (
	// If the max base fee is not explicitly set, use 3x the current base fee estimate
	defaultBaseFeeFactor          = 3
	poolTxsPerAccount             = 16
	pendingTxRefreshInterval      = 2 * time.Second
	defaultBlockAcceptanceTimeout = 30 * time.Second

	// epochCacheKey is the singleflight key for epoch caching
	// Since we only cache one epoch per destination, the key is constant
	epochCacheKey = "epoch"
)

// Client interface wraps the ethclient.Client interface for mocking purposes.
type Client interface {
	bind.ContractBackend
}

// Implements DestinationClient
type destinationClient struct {
	avaRPCClient DestinationRPCClient
	ethClient    bind.ContractBackend

	readonlyConcurrentSigners []*readonlyConcurrentSigner

	destinationBlockchainID ids.ID
	rpcEndpointURL          string
	evmChainID              *big.Int
	blockGasLimit           uint64
	gasFeeConfig            *GasFeeConfig
	logger                  logging.Logger
	txInclusionTimeout      time.Duration

	// Epoch cache for Granite - cached per destination blockchain
	epochValue        block.Epoch
	epochExpiration   time.Time
	epochSingleFlight singleflight.Group
	proposerClient    *clients.ProposerVMAPI
	epochDuration     time.Duration
}

func NewDestinationClient(
	logger logging.Logger,
	destinationBlockchain *config.DestinationBlockchain,
	epochDuration time.Duration,
) (*destinationClient, error) {
	destinationID, err := ids.FromString(destinationBlockchain.BlockchainID)
	if err != nil {
		return nil, fmt.Errorf("could not decode destination chain ID from string: %w", err)
	}

	signers, err := signer.NewSigners(destinationBlockchain)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer: %w", err)
	}

	// Dial the destination RPC endpoint
	rpcClient, err := utils.DialWithConfig(
		context.Background(),
		destinationBlockchain.RPCEndpoint.BaseURL,
		destinationBlockchain.RPCEndpoint.HTTPHeaders,
		destinationBlockchain.RPCEndpoint.QueryParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc endpoint: %w", err)
	}
	ethClient := ethclient.NewClient(rpcClient)

	evmChainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from destination chain endpoint: %w", err)
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
		log := logger.With(
			zap.Stringer("senderAddress", signer.Address()),
		)
		for {
			pendingNonce, err = ethClient.NonceAt(
				context.Background(),
				signer.Address(),
				big.NewInt(int64(rpc.PendingBlockNumber)),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to get pending nonce: %w", err)
			}

			currentNonce, err = ethClient.NonceAt(context.Background(), signer.Address(), nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get current nonce: %w", err)
			}

			// If the pending nonce is not equal to the current nonce, wait and check again
			if pendingNonce != currentNonce {
				log.Info(
					"Waiting for pending txs to be accepted",
					zap.Uint64("pendingNonce", pendingNonce),
					zap.Uint64("currentNonce", currentNonce),
				)
				<-ticker.C
				continue
			}

			log.Debug("Pending txs accepted")

			concurrentSigner := &concurrentSigner{
				logger:            log,
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
		zap.Stringer("evmChainID", evmChainID),
		zap.Uint64("nonce", pendingNonce),
	)

	// Create ProposerVM client for the destination chain
	endpoint, err := url.Parse(destinationBlockchain.RPCEndpoint.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rpc endpoint for ProposerVM client: %w", err)
	}

	baseURL := fmt.Sprintf("%s://%s", endpoint.Scheme, endpoint.Host)
	blockchainID := destinationBlockchain.BlockchainID
	proposerClient := clients.NewProposerVMAPI(baseURL, blockchainID, &destinationBlockchain.RPCEndpoint)
	gasFeeConfig := &GasFeeConfig{
		maxBaseFee:                 new(big.Int).SetUint64(destinationBlockchain.MaxBaseFee),
		suggestedPriorityFeeBuffer: new(big.Int).SetUint64(destinationBlockchain.SuggestedPriorityFeeBuffer),
		maxPriorityFeePerGas:       new(big.Int).SetUint64(destinationBlockchain.MaxPriorityFeePerGas),
	}
	destClient = destinationClient{
		avaRPCClient:              NewAvaDestinationClient(ethClient, rpcClient),
		ethClient:                 ethClient,
		readonlyConcurrentSigners: readonlyConcurrentSigners,
		destinationBlockchainID:   destinationID,
		rpcEndpointURL:            destinationBlockchain.RPCEndpoint.BaseURL,
		evmChainID:                evmChainID,
		logger:                    logger,
		blockGasLimit:             destinationBlockchain.BlockGasLimit,
		gasFeeConfig:              gasFeeConfig,
		txInclusionTimeout:        time.Duration(destinationBlockchain.TxInclusionTimeoutSeconds) * time.Second,
		proposerClient:            proposerClient,
		epochDuration:             epochDuration,
	}

	return &destClient, nil
}

func (c *destinationClient) RPCClient() DestinationRPCClient {
	return c.avaRPCClient
}

func (c *destinationClient) EVMChainID() *big.Int {
	return c.evmChainID
}

func (c *destinationClient) Logger() logging.Logger {
	return c.logger
}

func (c *destinationClient) GasFeeConfig() *GasFeeConfig {
	return c.gasFeeConfig
}

func (c *destinationClient) FeeFactor() int64 {
	return defaultBaseFeeFactor
}

func (c *destinationClient) ConcurrentSigners() []*readonlyConcurrentSigner {
	return c.readonlyConcurrentSigners
}

func (c *destinationClient) AccessList(data txData) types.AccessList {
	// Construct the actual transaction to broadcast on the destination chain
	// Create predicate from the signed warp message
	predicate := predicateutils.New(data.signedMessage.Bytes())

	// Create access list with the predicate for the warp precompile
	return types.AccessList{
		{
			Address:     warp.ContractAddress,
			StorageKeys: predicate,
		},
	}
}

func (c *destinationClient) TxInclusionTimeout() time.Duration {
	return c.txInclusionTimeout
}

func (c *destinationClient) getFeePerGas() (*big.Int, *big.Int, error) {
	return getFeePerGas(c)
}

// SendTx constructs, signs, and broadcast a transaction to deliver the given {signedMessage}
// to this chain with the provided {callData}.
func (c *destinationClient) SendTx(
	signedMessage *avalancheWarp.Message,
	deliverers set.Set[common.Address],
	toAddress string,
	gasLimit uint64,
	callData []byte,
) (*types.Receipt, error) {
	return SendTx(c, signedMessage, deliverers, toAddress, gasLimit, callData, c.txInclusionTimeout)
}

func (c *destinationClient) SenderAddresses() []common.Address {
	return SenderAddresses(c)
}

func (c *destinationClient) Client() Client {
	return c.ethClient
}

func (c *destinationClient) DestinationBlockchainID() ids.ID {
	return c.destinationBlockchainID
}

func (c *destinationClient) BlockGasLimit() uint64 {
	return c.blockGasLimit
}

func (c *destinationClient) GetRPCEndpointURL() string {
	return c.rpcEndpointURL
}

// GetPChainHeightForDestination determines the appropriate P-Chain height for validator set selection.
// The epoch is cached per destination blockchain to avoid per-message fetches.
func (c *destinationClient) GetPChainHeightForDestination(
	ctx context.Context,
) (uint64, error) {
	// Use singleflight to deduplicate concurrent fetches and serialize cache access
	result, err, _ := c.epochSingleFlight.Do(epochCacheKey, func() (interface{}, error) {
		// Check if cached epoch is still valid
		if !c.epochExpiration.IsZero() && time.Now().Before(c.epochExpiration) {
			return c.epochValue, nil
		}

		// Fetch new epoch
		epoch, fetchErr := c.proposerClient.GetCurrentEpoch(ctx)
		if fetchErr != nil {
			return block.Epoch{}, fetchErr
		}

		c.logger.Info("Successfully retrieved epoch from ProposerVM",
			zap.Stringer("destinationBlockchainID", c.destinationBlockchainID),
			zap.Any("epoch", epoch),
			zap.Duration("epochDuration", c.epochDuration),
		)

		// Calculate expiration time based on epoch.StartTime + epochDuration
		// epoch.StartTime is in nanoseconds (Unix timestamp)
		// Update cache
		c.epochValue = epoch
		c.epochExpiration = time.Unix(0, epoch.StartTime).Add(c.epochDuration)

		c.logger.Debug("Calculated epoch expiration",
			zap.Stringer("destinationBlockchainID", c.destinationBlockchainID),
			zap.Uint64("epochNumber", c.epochValue.Number),
			zap.Time("epochExpiration", c.epochExpiration),
		)

		return epoch, nil
	})

	if err != nil {
		c.logger.Error("Failed to get current epoch from destination chain ProposerVM",
			zap.Stringer("destinationBlockchainID", c.destinationBlockchainID),
			zap.Error(err),
		)
		return 0, err
	}

	epoch := result.(block.Epoch)

	// This should only be the case around activation time
	// but should be safe to keep this as a failsafe.
	if epoch.Number == 0 {
		c.logger.Info("Epoch number is 0, using current validators (ProposedHeight)")
		return pchainapi.ProposedHeight, nil
	}

	return epoch.PChainHeight, nil
}
