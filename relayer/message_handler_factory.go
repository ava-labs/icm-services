// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"fmt"

	"github.com/ava-labs/icm-services/messages"
	offchainregistry "github.com/ava-labs/icm-services/messages/off-chain-registry"
	"github.com/ava-labs/icm-services/messages/teleporter"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"google.golang.org/grpc"
)

// NewMessageHandlerFactory returns the MessageHandlerFactory for the protocol identified
// by cfg.MessageFormat. Returns an error for unknown or unsupported protocols.
func NewMessageHandlerFactory(
	address common.Address,
	cfg config.MessageProtocolConfig,
	deciderConnection *grpc.ClientConn,
) (messages.MessageHandlerFactory, error) {
	switch config.ParseMessageProtocol(cfg.MessageFormat) {
	case config.TELEPORTER:
		return teleporter.NewMessageHandlerFactory(address, cfg, deciderConnection)
	case config.OFF_CHAIN_REGISTRY:
		return offchainregistry.NewMessageHandlerFactory(cfg)
	case config.TELEPORTER_V2:
		return nil, fmt.Errorf("teleporter v2 is not yet supported")
	default:
		return nil, fmt.Errorf("invalid message format %s", cfg.MessageFormat)
	}
}
