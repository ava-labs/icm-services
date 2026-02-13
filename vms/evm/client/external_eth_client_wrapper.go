// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
)

// ExternalEthClientWrapper wraps libevm/bind ContractBackend.
// It adds stub implementations for Avalanche-specific methods that don't exist
// in standard go-ethereum clients.
type ExternalEthClientWrapper struct {
	EthClient
}

// NewExternalEthClientWrapper creates a new wrapper around libevm/ethclient.Client
func NewExternalEthClientWrapper(client EthClient) *ExternalEthClientWrapper {
	return &ExternalEthClientWrapper{EthClient: client}
}

// PendingCodeAt returns the code at the latest block for external EVMs.
// External EVMs don't have the "pending" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	// For external EVMs, use latest block instead of "accepted" state
	return w.CodeAt(ctx, account, nil)
}

// PendingCallContract executes a call against the latest block for external EVMs.
// External EVMs don't have the "pending" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	// For external EVMs, use latest block instead of "accepted" state
	return w.CallContract(ctx, call, nil)
}

// Ensure the wrapper satisfies bind.ContractBackend by implementing PendingContractCaller
var _ bind.PendingContractCaller = (*ExternalEthClientWrapper)(nil)
var _ bind.ContractBackend = (*ExternalEthClientWrapper)(nil)
var _ EthClient = (*ExternalEthClientWrapper)(nil)
