// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/utils"
	ethereum "github.com/ava-labs/libevm"
	"go.uber.org/zap"
)

// gasCapProbeInitCode is the init code of a contract creation that never completes: a JUMPDEST
// followed by a jump back to it (PUSH1 0, JUMP). Estimating its gas runs the node out of all the
// gas it is willing to try, which it then reports.
var gasCapProbeInitCode = []byte{0x5b, 0x60, 0x00, 0x56}

// CheckRPCGasCap verifies that the node behind [client] estimates calls with at least
// [blockGasLimit] gas, and returns an error naming the cap it reports if that is lower. The
// relayer vets every delivery by having the node estimate it at the block gas limit and trusts
// the verdict: a delivery that does not fit is abandoned without a transaction. A node whose RPC
// gas cap is lower runs deliveries with less gas than the transaction will carry, and condemns
// ones that would have been mined. Nodes do not expose their cap, but estimation reveals it:
// asked to estimate a call that never completes, bounded by [blockGasLimit], a node answers with
// the bound it ran out of gas at, which is the lower of that and its cap.
//
// A node that answers in any other way, such as one that refuses contract creation from the zero
// address, cannot be checked. That is logged and startup continues, so that the requirement is
// enforced wherever it can be observed. No check is made when [blockGasLimit] is zero.
func CheckRPCGasCap(
	ctx context.Context,
	logger logging.Logger,
	client Client,
	blockGasLimit uint64,
) error {
	if blockGasLimit == 0 {
		return nil
	}
	// The node stops at the first execution, which fails, so this is one RPC and costs the
	// node at most the block gas limit's worth of jumps.
	ctx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	_, err := client.EstimateGas(ctx, ethereum.CallMsg{
		Gas:  blockGasLimit,
		Data: gasCapProbeInitCode,
	})
	allowance, exceeds := ExceedsGasAllowance(err)
	switch {
	case exceeds && allowance >= blockGasLimit:
		logger.Debug(
			"Destination node allows calls at the block gas limit",
			zap.Uint64("blockGasLimit", blockGasLimit),
		)
		return nil
	case exceeds && allowance != 0:
		return fmt.Errorf(
			"destination node RPC gas cap (%d) is below the configured block-gas-limit (%d): "+
				"raise the node's rpc-gas-cap or lower block-gas-limit",
			allowance,
			blockGasLimit,
		)
	default:
		logger.Warn(
			"Could not determine the destination node's RPC gas cap; "+
				"deliveries needing more gas than the cap would be abandoned as not fitting in a block",
			zap.Uint64("blockGasLimit", blockGasLimit),
			zap.Error(err),
		)
		return nil
	}
}
