// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
)

// EthClient extends the bind.ContractBackend interface to provide
// additional methods for interacting with Ethereum networks.
// This enables the DestinationClient interface to work with both
// Avalanche L1 chains and external EVM chains.
//
// Embeds bind.ContractBackend to support ABI binding operations.
type EthClient interface {
	bind.ContractBackend
	DestinationRPCClient
}

// ExternalEthClientWrapper wraps libevm/bind ContractBackend.
// It adds stub implementations for Avalanche-specific methods that don't exist
// in standard go-ethereum clients.
type ExternalEthClientWrapper struct {
	*ethclient.Client
}

// NewExternalEthClientWrapper creates a new wrapper around libevm/ethclient.Client
func NewExternalEthClientWrapper(client *ethclient.Client) *ExternalEthClientWrapper {
	return &ExternalEthClientWrapper{Client: client}
}

// PendingCodeAt returns the code at the latest block for external EVMs.
// External EVMs don't have the "pending" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	// For external EVMs, use the latest block instead of "accepted" state
	return w.CodeAt(ctx, account, nil)
}

// PendingCallContract executes a call against the latest block for external EVMs.
// External EVMs don't have the "pending" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	// For external EVMs, use the latest block instead of "accepted" state
	return w.CallContract(ctx, call, nil)
}

func (w *ExternalEthClientWrapper) EstimateBaseFee(ctx context.Context) (*big.Int, error) {
	// Get base fee from the latest block header
	baseFeeCtx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()

	header, err := w.HeaderByNumber(baseFeeCtx, nil) // nil = latest block
	if err != nil {
		return nil, err
	}
	if header.BaseFee == nil {
		return nil, errors.New("chain does not support EIP-1559")
	}
	return header.BaseFee, nil
}

var _ EthClient = (*ExternalEthClientWrapper)(nil)
