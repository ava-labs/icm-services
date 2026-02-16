// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_destination_rpc_client.go -package=mocks

package evm

import (
	"context"
	"math/big"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
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
