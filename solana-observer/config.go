// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ava-labs/libevm/common"
)

// Config is the on-disk configuration for the Solana observer daemon.
//
// The observer watches a Solana WebSocket endpoint for transactions that mention
// a configured Memo program, verifies the transaction was finalized, wraps it as
// an OracleMessage, drives BLS aggregation via the signature-aggregator, and
// submits the resulting warp message to the destination L1's TeleporterMessengerV2.
//
// Solana-side settings live under `solana`, L1-side settings under `l1`, and the
// aggregator endpoint alongside them. The program allowlist is read from the
// sidecar's own config file at `sidecar_config_path`, keeping the sidecar as
// the single source of truth (same pattern as the validator; see
// avalanchego/graft/subnet-evm/plugin/evm/vm.go).
type Config struct {
	Solana     SolanaConfig `json:"solana"`
	L1         L1Config     `json:"l1"`
	Aggregator string       `json:"aggregator_url"`

	// SidecarConfigPath is the path to the same sidecar config file the
	// validators consume. The observer reads the `verifiers.solana.allowed_programs`
	// list out of it to decide which Solana program IDs to subscribe to,
	// unless SubscriptionPrograms is set.
	SidecarConfigPath string `json:"sidecar_config_path"`

	// SubscriptionPrograms, if non-empty, overrides the subscription program list
	// derived from the sidecar config. Use this when the programs you want to
	// watch (e.g. an escrow program) differ from the programs the sidecar
	// verifies (e.g. the Memo program whose instruction data carries the payload).
	SubscriptionPrograms []string `json:"subscription_programs,omitempty"`

	// NonceFile is a path where the observer persists the next OracleMessage
	// nonce per (source_type, source_address) tuple, so restarts don't collide
	// with previously delivered nonces on the L1 side.
	NonceFile string `json:"nonce_file"`
}

type SolanaConfig struct {
	// WSURL is the Solana WebSocket endpoint (e.g. "wss://api.devnet.solana.com").
	WSURL string `json:"ws_url"`
	// RPCURL is the HTTP JSON-RPC endpoint used to fetch full transaction bodies
	// after a logsSubscribe notification (e.g. "https://api.devnet.solana.com").
	RPCURL string `json:"rpc_url"`
	// Commitment is the finality level to subscribe at ("finalized" or "confirmed").
	// The sidecar's verifier requires finalized commitment to accept the tx, so
	// this must be "finalized" for the pipeline to converge.
	Commitment string `json:"commitment"`
}

type L1Config struct {
	// RPCURL is the destination L1's EVM JSON-RPC endpoint.
	RPCURL string `json:"rpc_url"`
	// ChainID is the destination L1's EVM chain ID.
	ChainID uint64 `json:"chain_id"`
	// BlockchainID is the destination L1's Avalanche blockchain ID (base58 CB58).
	// The observer needs this to construct the warp UnsignedMessage's SourceChainID.
	BlockchainID string `json:"blockchain_id"`
	// NetworkID is the ICM network ID.
	NetworkID uint32 `json:"network_id"`
	// TeleporterAddress is the deployed TeleporterMessengerV2 contract on the L1.
	TeleporterAddress common.Address `json:"teleporter_address"`
	// DestContract is the target application contract that implements
	// IOracleMessageReceiver. The observer sets OracleMessage.DestContract to
	// this address on every message.
	DestContract common.Address `json:"dest_contract"`
	// SubnetID is the L1's subnet ID; passed to the aggregator as signing-subnet-id.
	SubnetID string `json:"subnet_id"`
	// DeliveryPrivateKeyHex is the hex-encoded ECDSA private key the observer
	// uses to pay gas for L1 delivery transactions. This account must be funded
	// on the destination L1.
	DeliveryPrivateKeyHex string `json:"delivery_private_key_hex"`
}

// Load reads and parses an observer config file.
func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read observer config %s: %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse observer config %s: %w", path, err)
	}
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) validate() error {
	if c.Solana.WSURL == "" {
		return errors.New("solana.ws_url is required")
	}
	if c.Solana.RPCURL == "" {
		return errors.New("solana.rpc_url is required")
	}
	if c.Solana.Commitment == "" {
		c.Solana.Commitment = "finalized"
	}
	if c.Aggregator == "" {
		return errors.New("aggregator_url is required")
	}
	if c.SidecarConfigPath == "" {
		return errors.New("sidecar_config_path is required")
	}
	if c.L1.RPCURL == "" {
		return errors.New("l1.rpc_url is required")
	}
	if c.L1.BlockchainID == "" {
		return errors.New("l1.blockchain_id is required")
	}
	if c.L1.DeliveryPrivateKeyHex == "" {
		return errors.New("l1.delivery_private_key_hex is required")
	}
	if c.NonceFile == "" {
		return errors.New("nonce_file is required")
	}
	return nil
}
