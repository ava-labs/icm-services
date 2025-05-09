// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vms

import (
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/relayer/config"
	relayerTypes "github.com/ava-labs/icm-services/types"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/subnet-evm/ethclient"
)

// Subscriber subscribes to VM events containing Warp message data. The events written to the
// channel returned by Logs() are assumed to be in block order. Logs within individual blocks
// may be in any order.
type Subscriber interface {
	// ProcessFromHeight processes events from {height} to the latest block.
	// Writes true to the channel on success, false on failure
	ProcessFromHeight(height *big.Int, done chan bool)

	// Subscribe registers a subscription. After Subscribe is called,
	// log events that match [filter] are written to the channel returned
	// by Logs
	Subscribe(retryTimeout time.Duration) error

	// ICMBlocks returns the channel that the subscription writes ICM block info to
	ICMBlocks() <-chan *relayerTypes.WarpBlockInfo

	// SubscribeErr returns the channel that the subscription writes errors to
	// If an error is sent to this channel, the subscription should be closed
	SubscribeErr() <-chan error

	// Err returns the channel that the subscriber writes miscellaneous errors to
	// that are not recoverable from by resubscribing.
	Err() <-chan error

	// Cancel cancels the subscription
	Cancel()
}

// NewSubscriber returns a concrete Subscriber according to the VM specified by [subnetInfo]
func NewSubscriber(
	logger logging.Logger,
	vm config.VM,
	blockchainID ids.ID,
	ethWSClient ethclient.Client,
	ethRPCClient ethclient.Client,
) Subscriber {
	switch vm {
	case config.EVM:
		return evm.NewSubscriber(logger, blockchainID, ethWSClient, ethRPCClient)
	default:
		return nil
	}
}
