package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// The block gas limit that can be specified for a Teleporter message
	// Based on the C-Chain 15_000_000 gas limit per block, with other Warp message gas overhead conservatively estimated.
	defaultBlockGasLimit = 12_000_000
)

// Destination blockchain configuration. Specifies how to connect to and issue
// transactions on the destination blockchain.
type DestinationBlockchain struct {
	SubnetID           string            `mapstructure:"subnet-id" json:"subnet-id"`
	BlockchainID       string            `mapstructure:"blockchain-id" json:"blockchain-id"`
	VM                 string            `mapstructure:"vm" json:"vm"`
	RPCEndpoint        basecfg.APIConfig `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	KMSKeyIDs          []string          `mapstructure:"kms-key-ids" json:"kms-key-ids"`
	KMSAWSRegions      []string          `mapstructure:"kms-aws-regions" json:"kms-aws-regions"`
	AccountPrivateKeys []string          `mapstructure:"account-private-keys" json:"account-private-keys"`
	BlockGasLimit      uint64            `mapstructure:"block-gas-limit" json:"block-gas-limit"`

	// Fetched from the chain after startup
	warpConfig WarpConfig

	// convenience fields to access parsed data after initialization
	subnetID     ids.ID
	blockchainID ids.ID
}

// Validates the destination subnet configuration
func (s *DestinationBlockchain) Validate() error {
	if s.BlockGasLimit == 0 {
		s.BlockGasLimit = defaultBlockGasLimit
	}
	if err := s.RPCEndpoint.Validate(); err != nil {
		return fmt.Errorf("invalid rpc-endpoint in destination subnet configuration: %w", err)
	}
	if len(s.KMSKeyIDs) != 0 {
		if len(s.KMSAWSRegions) != len(s.KMSKeyIDs) {
			return errors.New("KMS key IDs provided without matching AWS regions")
		}
	} else {
		if len(s.AccountPrivateKeys) == 0 {
			return errors.New("No keys provided")
		}
		for _, pkey := range s.AccountPrivateKeys {
			if _, err := crypto.HexToECDSA(utils.SanitizeHexString(pkey)); err != nil {
				return utils.ErrInvalidPrivateKeyHex
			}
		}
	}

	// Validate the VM specific settings
	vm := ParseVM(s.VM)
	if vm == UNKNOWN_VM {
		return fmt.Errorf("unsupported VM type for source subnet: %s", s.VM)
	}

	// Validate and store the subnet and blockchain IDs for future use
	blockchainID, err := utils.HexOrCB58ToID(s.BlockchainID)
	if err != nil {
		return fmt.Errorf("invalid blockchainID '%s' in configuration. error: %w", s.BlockchainID, err)
	}
	s.blockchainID = blockchainID
	subnetID, err := utils.HexOrCB58ToID(s.SubnetID)
	if err != nil {
		return fmt.Errorf("invalid subnetID '%s' in configuration. error: %w", s.SubnetID, err)
	}
	s.subnetID = subnetID

	if s.subnetID == constants.PrimaryNetworkID &&
		s.BlockGasLimit > defaultBlockGasLimit {
		return fmt.Errorf("C-Chain block-gas-limit '%d' exceeded", s.BlockGasLimit)
	}

	return nil
}

func (s *DestinationBlockchain) GetSubnetID() ids.ID {
	return s.subnetID
}

func (s *DestinationBlockchain) GetBlockchainID() ids.ID {
	return s.blockchainID
}

func (s *DestinationBlockchain) initializeWarpConfigs() error {
	blockchainID, err := ids.FromString(s.BlockchainID)
	if err != nil {
		return fmt.Errorf("invalid blockchainID in configuration. error: %w", err)
	}
	subnetID, err := ids.FromString(s.SubnetID)
	if err != nil {
		return fmt.Errorf("invalid subnetID in configuration. error: %w", err)
	}
	// If the destination blockchain is the primary network, use the default quorum
	// primary network signers here are irrelevant and can be left at default value
	if subnetID == constants.PrimaryNetworkID {
		s.warpConfig = WarpConfig{
			QuorumNumerator: warp.WarpDefaultQuorumNumerator,
		}
		return nil
	}

	client, err := utils.NewEthClientWithConfig(
		context.Background(),
		s.RPCEndpoint.BaseURL,
		s.RPCEndpoint.HTTPHeaders,
		s.RPCEndpoint.QueryParams,
	)
	defer client.Close()
	if err != nil {
		return fmt.Errorf("failed to dial destination blockchain %s: %w", blockchainID, err)
	}
	subnetWarpConfig, err := getWarpConfig(client)
	if err != nil {
		return fmt.Errorf("failed to fetch warp config for blockchain %s: %w", blockchainID, err)
	}
	s.warpConfig = warpConfigFromSubnetWarpConfig(*subnetWarpConfig)
	return nil
}

// Warp Configuration, fetched from the chain config
type WarpConfig struct {
	QuorumNumerator              uint64
	RequirePrimaryNetworkSigners bool
}
