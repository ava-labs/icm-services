// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

// Supported Message Protocols
type MessageProtocol int

const (
	UNKNOWN_MESSAGE_PROTOCOL MessageProtocol = iota
	TELEPORTER
	OFF_CHAIN_REGISTRY
	TELEPORTER_V2
)

func (msg MessageProtocol) String() string {
	switch msg {
	case TELEPORTER:
		return "teleporter"
	case OFF_CHAIN_REGISTRY:
		return "off-chain-registry"
	case TELEPORTER_V2:
		return "teleporter-v2"
	default:
		return "unknown"
	}
}

// ParseMessageProtocol returns the MessageProtocol corresponding to [msg]
func ParseMessageProtocol(msg string) MessageProtocol {
	switch msg {
	case "teleporter":
		return TELEPORTER
	case "off-chain-registry":
		return OFF_CHAIN_REGISTRY
	case "teleporter-v2":
		return TELEPORTER_V2
	default:
		return UNKNOWN_MESSAGE_PROTOCOL
	}
}
