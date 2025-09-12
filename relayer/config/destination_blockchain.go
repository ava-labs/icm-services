package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/set"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
)

const (
	// The block gas limit that can be specified for a Teleporter message
	// Based on the C-Chain 15_000_000 gas limit per block, with other Warp message gas overhead conservatively estimated.
	defaultBlockGasLimit             = 12_000_000
	defaultMaxPriorityFeePerGas      = 2500000000 // 2.5 gwei
	defaultTxInclusionTimeoutSeconds = 30
)

type KMSKey struct {
	KeyID     string `mapstructure:"key-id" json:"key-id"`
	AWSRegion string `mapstructure:"aws-region" json:"aws-region"`
}

// Destination blockchain configuration. Specifies how to connect to and issue
// transactions on the destination blockchain.
type DestinationBlockchain struct {
	SubnetID             string            `mapstructure:"subnet-id" json:"subnet-id"`
	BlockchainID         string            `mapstructure:"blockchain-id" json:"blockchain-id"`
	VM                   string            `mapstructure:"vm" json:"vm"`
	RPCEndpoint          basecfg.APIConfig `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
	KMSKeyID             string            `mapstructure:"kms-key-id" json:"kms-key-id"`
	KMSAWSRegion         string            `mapstructure:"kms-aws-region" json:"kms-aws-region"`
	AccountPrivateKey    string            `mapstructure:"account-private-key" json:"account-private-key" sensitive:"true"`
	KMSKeys              []KMSKey          `mapstructure:"kms-keys" json:"kms-keys" sensitive:"true"`
	AccountPrivateKeys   []string          `mapstructure:"account-private-keys-list" json:"account-private-keys-list" sensitive:"true"` //nolint:lll
	BlockGasLimit        uint64            `mapstructure:"block-gas-limit" json:"block-gas-limit"`
	MaxBaseFee           uint64            `mapstructure:"max-base-fee" json:"max-base-fee"`
	MaxPriorityFeePerGas uint64            `mapstructure:"max-priority-fee-per-gas" json:"max-priority-fee-per-gas"`

	TxInclusionTimeoutSeconds uint64 `mapstructure:"tx-inclusion-timeout-seconds" json:"tx-inclusion-timeout-seconds"`

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
	if s.KMSKeyID != "" {
		if s.KMSAWSRegion == "" {
			return errors.New("KMS key ID provided without an AWS region")
		}
	}
	for _, id := range s.KMSKeys {
		if id.KeyID == "" {
			return errors.New("KMS key ID provided without an AWS region")
		}
		if id.AWSRegion == "" {
			return errors.New("KMS AWS region provided without a key ID")
		}
	}

	// Collect and deduplicate the account private keys
	privateKeys := s.AccountPrivateKeys
	if s.AccountPrivateKey != "" {
		privateKeys = append(privateKeys, s.AccountPrivateKey)
	}

	for i, pkey := range privateKeys {
		if _, err := crypto.HexToECDSA(utils.SanitizeHexString(pkey)); err != nil {
			return utils.ErrInvalidPrivateKeyHex
		}
		privateKeys[i] = utils.SanitizeHexString(pkey)
	}
	uniquePks := set.Of(privateKeys...)

	// Deduplicate the KMS key IDs and AWS regions
	kmsKeys := s.KMSKeys
	if s.KMSKeyID != "" {
		kmsKeys = append(kmsKeys, KMSKey{
			KeyID:     s.KMSKeyID,
			AWSRegion: s.KMSAWSRegion,
		})
	}
	uniqueKmsKeys := set.Of(kmsKeys...)

	if len(uniqueKmsKeys) == 0 && len(uniquePks) == 0 {
		return errors.New("no account private keys or KMS keys provided")
	}

	s.AccountPrivateKeys = uniquePks.List()
	s.KMSKeys = uniqueKmsKeys.List()

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
		return fmt.Errorf("c-chain block-gas-limit '%d' exceeded", s.BlockGasLimit)
	}

	// If not set, use the default value for the maximum priority fee per gas.
	// We do not set any default for the max base fee. Instead, if unset, we
	// will use the current base fee from the chain at the time of the transaction.
	if s.MaxPriorityFeePerGas == 0 {
		s.MaxPriorityFeePerGas = defaultMaxPriorityFeePerGas
	}

	if s.TxInclusionTimeoutSeconds == 0 {
		s.TxInclusionTimeoutSeconds = defaultTxInclusionTimeoutSeconds
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
