// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
)

// ZKConsensusUpdater configures a consensus updater for a ZKAdapter deployed on a destination
// chain. The updater polls the Boundless subgraph for Ethereum consensus proofs and submits them
// to the adapter as transitions, keeping its confirmed beacon anchors current so that zk messages
// from Ethereum can be proven against them.
type ZKConsensusUpdater struct {
	// RPC endpoint of the destination chain the ZKAdapter is deployed on.
	RPCEndpoint string `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	// Hex-encoded private key used to sign transition transactions. It must not be shared with
	// a message-delivery account, or their nonces collide on-chain.
	PrivateKey string `mapstructure:"private-key" json:"private-key" sensitive:"true"`
	// Address of the deployed ZKAdapter contract.
	AdapterAddress string `mapstructure:"adapter-address" json:"adapter-address"`
	// GraphQL endpoint of the Boundless subgraph indexing consensus proofs.
	SubgraphURL string `mapstructure:"subgraph-url" json:"subgraph-url"`
	// Poll interval in seconds (default 10)
	PollIntervalSeconds uint64 `mapstructure:"poll-interval-seconds" json:"poll-interval-seconds"`
	// Maximum suggested gas price (in gwei) on the destination chain at which the updater will
	// submit a transition. When exceeded the transition is deferred to the next poll. 0 disables
	// the check.
	MaxGasPriceGwei uint64 `mapstructure:"max-gas-price-gwei" json:"max-gas-price-gwei,omitempty"`

	// senderAddress is the address derived from PrivateKey, set by Validate.
	senderAddress common.Address
}

// Validate checks the fields required to run the consensus updater. otherSenders maps the sender
// addresses of the relayer's other signing accounts to a description of each. The updater's
// account must not be one of them, or their transactions' nonces collide on-chain.
func (z *ZKConsensusUpdater) Validate(otherSenders map[common.Address]string) error {
	if z.RPCEndpoint == "" {
		return errors.New("rpc-endpoint required for zk consensus updater")
	}
	if z.PrivateKey == "" {
		return errors.New("private-key required for zk consensus updater")
	}
	// Parse the key the same way startup does, so a malformed key fails config validation
	// instead of the updater's startup. The parse error is not wrapped, as the key is sensitive.
	privateKey, err := crypto.HexToECDSA(utils.SanitizeHexString(z.PrivateKey))
	if err != nil {
		return fmt.Errorf("invalid private-key for zk consensus updater: %w", utils.ErrInvalidPrivateKeyHex)
	}
	sender := crypto.PubkeyToAddress(privateKey.PublicKey)
	if owner, ok := otherSenders[sender]; ok {
		return fmt.Errorf(
			"private-key for zk consensus updater signs as %s, which is also used by %s; their nonces would collide",
			sender, owner,
		)
	}
	z.senderAddress = sender
	if !common.IsHexAddress(z.AdapterAddress) {
		return fmt.Errorf("invalid adapter-address %q", z.AdapterAddress)
	}
	if z.SubgraphURL == "" {
		return errors.New("subgraph-url required for zk consensus updater")
	}
	return nil
}

// senderAccountAddresses returns the sender address of every private key the relayer signs
// destination and external EVM transactions with, mapped to a description of its account, for
// ZKConsensusUpdater.Validate. It must run after the destinations are validated. Keys that fail
// to parse are skipped, since their own validation reports them. KMS-backed accounts are not
// included, as their addresses are only known after querying KMS.
func (c *Config) senderAccountAddresses() map[common.Address]string {
	senders := make(map[common.Address]string)
	add := func(key, owner string) {
		privateKey, err := crypto.HexToECDSA(utils.SanitizeHexString(key))
		if err != nil {
			return
		}
		senders[crypto.PubkeyToAddress(privateKey.PublicKey)] = owner
	}
	for _, d := range c.DestinationBlockchains {
		for _, key := range d.AccountPrivateKeys {
			add(key, fmt.Sprintf("destination blockchain %s account key", d.BlockchainID))
		}
	}
	for _, e := range c.ExternalEVMDestinations {
		add(e.PrivateKey, fmt.Sprintf("external EVM destination %s private-key", e.ContractAddress))
		add(e.DeliveryPrivateKey, fmt.Sprintf("external EVM destination %s delivery-private-key", e.ContractAddress))
	}
	return senders
}

// GetPollInterval returns how often the updater polls the subgraph for new proofs, falling
// back to defaultPollInterval when PollIntervalSeconds is unset.
func (z *ZKConsensusUpdater) GetPollInterval() time.Duration {
	if z.PollIntervalSeconds == 0 {
		return defaultPollInterval
	}
	return time.Duration(z.PollIntervalSeconds) * time.Second
}
