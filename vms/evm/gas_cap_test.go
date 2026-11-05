// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	ethereum "github.com/ava-labs/libevm"
	"github.com/stretchr/testify/require"
)

// estimateOnlyClient is a Client that answers gas estimations, and records the estimated call.
type estimateOnlyClient struct {
	Client
	estimated uint64
	err       error
	called    bool
	call      ethereum.CallMsg
}

func (c *estimateOnlyClient) EstimateGas(_ context.Context, call ethereum.CallMsg) (uint64, error) {
	c.called = true
	c.call = call
	return c.estimated, c.err
}

// The check asks the node to estimate a call that cannot complete, bounded by the block gas
// limit, and reads the bound the node ran out of gas at. Only a node that states a lower bound
// fails the check: one that cannot be read is let through with a warning.
func TestCheckRPCGasCap(t *testing.T) {
	const blockGasLimit = uint64(12_000_000)
	tests := []struct {
		name      string
		estimated uint64
		err       error
		wantErr   bool
	}{
		{
			name: "node tries the block gas limit",
			err:  errors.New("gas required exceeds allowance (12000000)"),
		},
		{
			name:    "node cap below the block gas limit",
			err:     errors.New("gas required exceeds allowance (8000000)"),
			wantErr: true,
		},
		{
			name: "bound not stated",
			err:  errors.New("gas required exceeds allowance"),
		},
		{
			name: "node refuses the probe",
			err:  errors.New("sender address not allow listed"),
		},
		{
			name: "transport failure",
			err:  errors.New("connection refused"),
		},
		{
			name:      "probe unexpectedly estimated",
			estimated: 21_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &estimateOnlyClient{estimated: tt.estimated, err: tt.err}

			err := CheckRPCGasCap(context.Background(), logging.NoLog{}, client, blockGasLimit)
			if tt.wantErr {
				require.ErrorContains(t, err, "8000000")
				require.ErrorContains(t, err, "12000000")
			} else {
				require.NoError(t, err)
			}
			require.True(t, client.called)
			require.Equal(t, blockGasLimit, client.call.Gas)
			require.Equal(t, gasCapProbeInitCode, client.call.Data)
			require.Nil(t, client.call.To, "the probe must be a contract creation")
		})
	}

	t.Run("no block gas limit configured", func(t *testing.T) {
		client := &estimateOnlyClient{err: errors.New("must not be called")}
		require.NoError(t, CheckRPCGasCap(context.Background(), logging.NoLog{}, client, 0))
		require.False(t, client.called)
	})
}
