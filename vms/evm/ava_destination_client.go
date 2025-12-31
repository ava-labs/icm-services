// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"math/big"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/libevm/rpc"
)

var _ DestinationRPCClient = (*avaDestinationClient)(nil)

// avaDestinationClient wraps libevm's ethclient.Client and implements DestinationRPCClient.
// It delegates to the underlying ethclient.Client for methods that are available,
// and provides stubs for methods that need manual implementation.
type avaDestinationClient struct {
	ethClient *ethclient.Client
	rpcClient *rpc.Client
}

// NewAvaDestinationClient creates a new avaDestinationClient that wraps the provided ethclient.Client.
// The rpcClient parameter should be the same RPC client that was used to create the ethclient.
// This allows direct CallContext access to other endpoints.
func NewAvaDestinationClient(ethClient *ethclient.Client, rpcClient *rpc.Client) *avaDestinationClient {
	return &avaDestinationClient{
		ethClient: ethClient,
		rpcClient: rpcClient,
	}
}

func (c *avaDestinationClient) RPCClient() *rpc.Client {
	return c.rpcClient
}

func (c *avaDestinationClient) BlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	return c.ethClient.BlockByNumber(ctx, blockNumber)
}

func (c *avaDestinationClient) ChainID(ctx context.Context) (*big.Int, error) {
	return c.ethClient.ChainID(ctx)
}

func (c *avaDestinationClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return c.ethClient.NonceAt(ctx, account, blockNumber)
}

func (c *avaDestinationClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.ethClient.SuggestGasTipCap(ctx)
}

func (c *avaDestinationClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}

func (c *avaDestinationClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.ethClient.TransactionReceipt(ctx, txHash)
}

func (c *avaDestinationClient) BlockNumber(ctx context.Context) (uint64, error) {
	return c.ethClient.BlockNumber(ctx)
}

func (c *avaDestinationClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.CallContract(ctx, msg, blockNumber)
}

// EstimateBaseFee is a stub that needs to be implemented manually.
// This method is not available in libevm's ethclient.Client.
// TODO: Handle base fee estimation for both upstream evm and subnet-evm.
func (c *avaDestinationClient) EstimateBaseFee(ctx context.Context) (*big.Int, error) {
	var hex hexutil.Big
	err := c.rpcClient.CallContext(ctx, &hex, "eth_baseFee")
	if err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}
