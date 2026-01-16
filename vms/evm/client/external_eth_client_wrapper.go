// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
)

// ExternalEthClientWrapper wraps libevm/ethclient.Client to implement EthClient.
// It adds stub implementations for Avalanche-specific methods that don't exist
// in standard go-ethereum clients.
type ExternalEthClientWrapper struct {
	*ethclient.Client
}

// Compile-time check that ExternalEthClientWrapper implements EthClient
var _ EthClient = (*ExternalEthClientWrapper)(nil)

// NewExternalEthClientWrapper creates a new wrapper around libevm/ethclient.Client
func NewExternalEthClientWrapper(client *ethclient.Client) *ExternalEthClientWrapper {
	return &ExternalEthClientWrapper{Client: client}
}

// AcceptedCodeAt returns the code at the latest block for external EVMs.
// External EVMs don't have the "accepted" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) AcceptedCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	// For external EVMs, use latest block instead of "accepted" state
	return w.CodeAt(ctx, account, nil)
}

// AcceptedCallContract executes a call against the latest block for external EVMs.
// External EVMs don't have the "accepted" state concept from Avalanche,
// so we fall back to using the latest block.
func (w *ExternalEthClientWrapper) AcceptedCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	// For external EVMs, use latest block instead of "accepted" state
	return w.CallContract(ctx, call, nil)
}

// Ensure the wrapper satisfies bind.ContractBackend by implementing AcceptedContractCaller
var _ bind.AcceptedContractCaller = (*ExternalEthClientWrapper)(nil)
