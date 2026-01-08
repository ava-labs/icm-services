// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
)

// EthClient is a common interface for EVM clients that both
// subnet-evm/ethclient.Client and libevm/ethclient.Client can satisfy.
// This enables the DestinationClient interface to work with both
// Avalanche L1 chains and external EVM chains.
//
// Embeds bind.ContractBackend to support ABI binding operations.
type EthClient interface {
	bind.ContractBackend

	// ChainID returns the chain ID of the connected network
	ChainID(ctx context.Context) (*big.Int, error)

	// SendTransaction sends a signed transaction to the network
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	// TransactionReceipt returns the receipt of a transaction by transaction hash
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}
