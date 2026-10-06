// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params/extras"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/core/types"
	"github.com/stretchr/testify/require"
)

var (
	awsRegion string = "us-west-2"
	kmsKey1   string = "test-kms-id1"
)

// GetRelayerAccountPrivateKey tests. Individual cases must be run in their own functions
// because they modify the environment variables.

// setups config json file and writes content
func setupConfigJSON(t *testing.T, rootPath string, value string) string {
	configFilePath := filepath.Join(rootPath, "config.json")
	require.NoError(t, os.WriteFile(configFilePath, []byte(value), 0o600))
	return configFilePath
}

type stubRPCClient struct {
	chainConfig  *params.ChainConfigWithUpgradesJSON
	latestHeader *types.Header
}

func (c *stubRPCClient) ChainConfig(ctx context.Context) (*params.ChainConfigWithUpgradesJSON, error) {
	return c.chainConfig, nil
}

func (c *stubRPCClient) LatestHeader(ctx context.Context) (*types.Header, error) {
	return c.latestHeader, nil
}

func TestMultipleSignersConfig(t *testing.T) {
	testCases := []struct {
		name           string
		baseConfig     Config
		configModifier func(Config) Config
		envSetter      func()
		resultVerifier func(Config) bool
	}{
		{
			name:           "global pk",
			baseConfig:     TestValidConfig,
			configModifier: func(c Config) Config { return c },
			envSetter: func() {
				t.Setenv(accountPrivateKeyEnvVarName, testPk2)
			},
			resultVerifier: func(c Config) bool {
				for _, subnet := range c.DestinationBlockchains {
					pks := set.Of(subnet.AccountPrivateKeys...)
					if !pks.Contains(utils.SanitizeHexString(testPk2)) {
						return false
					}
				}
				return true
			},
		},
		{
			name:           "multiple global pks",
			baseConfig:     TestValidConfig,
			configModifier: func(c Config) Config { return c },
			envSetter: func() {
				t.Setenv(accountPrivateKeyListEnvVarName, strings.Join([]string{testPk1, testPk2}, " "))
			},
			resultVerifier: func(c Config) bool {
				for _, subnet := range c.DestinationBlockchains {
					pks := set.Of(subnet.AccountPrivateKeys...)
					if !pks.Contains(utils.SanitizeHexString(testPk1)) || !pks.Contains(utils.SanitizeHexString(testPk2)) {
						return false
					}
				}
				return true
			},
		},
		{
			name:           "individual and multiple global pks",
			baseConfig:     TestValidConfig,
			configModifier: func(c Config) Config { return c },
			envSetter: func() {
				t.Setenv(accountPrivateKeyEnvVarName, testPk1)
				t.Setenv(accountPrivateKeyListEnvVarName, strings.Join([]string{testPk2, testPk3}, " "))
			},
			resultVerifier: func(c Config) bool {
				for _, subnet := range c.DestinationBlockchains {
					pks := set.Of(subnet.AccountPrivateKeys...)
					if !pks.Contains(utils.SanitizeHexString(testPk1)) ||
						!pks.Contains(utils.SanitizeHexString(testPk2)) ||
						!pks.Contains(utils.SanitizeHexString(testPk3)) {
						return false
					}
				}
				return true
			},
		},
		{
			name:       "destination blockchain pk env",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = ""
				return c
			},
			envSetter: func() {
				varName := fmt.Sprintf(
					"%s_%s",
					accountPrivateKeyEnvVarName,
					TestValidConfig.DestinationBlockchains[0].BlockchainID,
				)
				t.Setenv(varName, testPk1)
			},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk1))
			},
		},
		{
			name: "multiple destination blockchain pks env", baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = ""
				return c
			},
			envSetter: func() {
				varName := fmt.Sprintf(
					"%s_%s",
					accountPrivateKeyListEnvVarName,
					TestValidConfig.DestinationBlockchains[0].BlockchainID,
				)
				t.Setenv(varName, strings.Join([]string{testPk1, testPk2}, " "))
			},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk1)) &&
					pks.Contains(utils.SanitizeHexString(testPk2))
			},
		},
		{
			name:       "individual and multiple destination blockchain pks env",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = ""
				return c
			},
			envSetter: func() {
				varName := fmt.Sprintf(
					"%s_%s",
					accountPrivateKeyListEnvVarName,
					TestValidConfig.DestinationBlockchains[0].BlockchainID,
				)
				t.Setenv(varName, strings.Join([]string{testPk1, testPk2}, " "))

				varName = fmt.Sprintf(
					"%s_%s",
					accountPrivateKeyEnvVarName,
					TestValidConfig.DestinationBlockchains[0].BlockchainID,
				)
				t.Setenv(varName, testPk3)
			},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk1)) &&
					pks.Contains(utils.SanitizeHexString(testPk2)) &&
					pks.Contains(utils.SanitizeHexString(testPk3))
			},
		},
		{
			name:       "destination blockchain pk cfg",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = testPk1
				return c
			},
			envSetter: func() {},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk1))
			},
		},
		{
			name:       "multiple destination blockchain pks cfg",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKeys = []string{testPk2, testPk3}
				return c
			},
			envSetter: func() {},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk2)) &&
					pks.Contains(utils.SanitizeHexString(testPk3))
			},
		},
		{
			name:       "individual and multiple destination blockchain pks cfg",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = testPk1
				c.DestinationBlockchains[0].AccountPrivateKeys = []string{testPk2, testPk3}
				return c
			},
			envSetter: func() {},
			resultVerifier: func(c Config) bool {
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk1)) &&
					pks.Contains(utils.SanitizeHexString(testPk2)) &&
					pks.Contains(utils.SanitizeHexString(testPk3))
			},
		},
		{
			name:       "global env, destination env, and destination cfg",
			baseConfig: TestValidConfig,
			configModifier: func(c Config) Config {
				c.DestinationBlockchains[0].AccountPrivateKey = testPk3
				return c
			},
			envSetter: func() {
				// Global pk
				t.Setenv(accountPrivateKeyEnvVarName, testPk1)

				// Destination pk
				varName := fmt.Sprintf(
					"%s_%s",
					accountPrivateKeyEnvVarName,
					TestValidConfig.DestinationBlockchains[0].BlockchainID,
				)
				t.Setenv(varName, testPk2)
			},
			resultVerifier: func(c Config) bool {
				// Check global pk
				for _, subnet := range c.DestinationBlockchains {
					pks := set.Of(subnet.AccountPrivateKeys...)
					if !pks.Contains(utils.SanitizeHexString(testPk2)) {
						return false
					}
				}

				// Check destination chain specific pk
				pks := set.Of(c.DestinationBlockchains[0].AccountPrivateKeys...)
				return pks.Contains(utils.SanitizeHexString(testPk2)) &&
					pks.Contains(utils.SanitizeHexString(testPk3))
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			root := t.TempDir()

			cfg := testCase.configModifier(testCase.baseConfig)
			cfgBytes, err := json.Marshal(cfg)
			require.NoError(t, err)

			configFile := setupConfigJSON(t, root, string(cfgBytes))

			flags := []string{"--config-file", configFile}
			testCase.envSetter()

			fs := BuildFlagSet()
			if err := fs.Parse(flags); err != nil {
				panic(fmt.Errorf("couldn't parse flags: %w", err))
			}
			v, err := BuildViper(fs)
			require.NoError(t, err)
			parsedCfg, err := NewConfig(v)
			require.NoError(t, err)
			require.NoError(t, parsedCfg.Validate())

			require.True(t, testCase.resultVerifier(parsedCfg))
		})
	}
}

func TestIndividualSignersConfig(t *testing.T) {
	dstCfg := *TestValidConfig.DestinationBlockchains[0]
	// Zero out all fields under test
	dstCfg.AccountPrivateKey = ""
	dstCfg.AccountPrivateKeys = nil
	dstCfg.KMSKeyID = ""
	dstCfg.KMSAWSRegion = ""
	dstCfg.KMSKeys = nil

	testCases := []struct {
		name   string
		dstCfg func() DestinationBlockchain
		valid  bool
	}{
		{
			name: "kms supplied",
			dstCfg: func() DestinationBlockchain {
				cfg := dstCfg
				cfg.KMSKeyID = kmsKey1
				cfg.KMSAWSRegion = awsRegion
				return cfg
			},
			valid: true,
		},
		{
			name: "account private key supplied",
			dstCfg: func() DestinationBlockchain {
				cfg := dstCfg
				cfg.AccountPrivateKey = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
				return cfg
			},
			valid: true,
		},
		{
			name: "neither supplied",
			dstCfg: func() DestinationBlockchain {
				return dstCfg
			},
			valid: false,
		},
		{
			name: "both supplied",
			dstCfg: func() DestinationBlockchain {
				cfg := dstCfg
				cfg.KMSKeyID = kmsKey1
				cfg.KMSAWSRegion = awsRegion
				cfg.AccountPrivateKey = "0x56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
				return cfg
			},
			valid: true,
		},
		{
			name: "missing aws region",
			dstCfg: func() DestinationBlockchain {
				cfg := dstCfg
				cfg.KMSKeyID = kmsKey1
				// Missing AWS region
				return cfg
			},
			valid: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dstCfg := testCase.dstCfg()
			err := dstCfg.Validate()
			if testCase.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestGetWarpConfig(t *testing.T) {
	// This is necessary to support fetching warp config from the genesis block
	evm.RegisterAllLibEVMExtras()
	blockchainID, err := ids.FromString("p433wpuXyJiDhyazPYyZMJeaoPSW76CBZ2x7wrVPLgvokotXz")
	require.NoError(t, err)
	subnetID, err := ids.FromString("2PsShLjrFFwR51DMcAh8pyuwzLn1Ym3zRhuXLTmLCR1STk2mL6")
	require.NoError(t, err)

	beforeCurrentBlockTime1 := uint64(time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC).Unix())
	beforeCurrentBlockTime2 := uint64(time.Date(2025, 9, 10, 0, 0, 0, 0, time.UTC).Unix())
	afterCurrentBlockTime := uint64(time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC).Unix())

	currentBlockTime := uint64(time.Date(2025, 9, 15, 0, 0, 0, 0, time.UTC).Unix())
	currentBlock := types.NewBlock(&types.Header{
		Time: currentBlockTime,
	}, nil, nil, nil, nil)

	testCases := []struct {
		name                string
		blockchainID        ids.ID
		subnetID            ids.ID
		chainConfig         params.ChainConfigWithUpgradesJSON
		getChainConfigCalls int
		expectedError       error
		expectedWarpConfig  WarpConfig
	}{
		{
			name:                "subnet genesis precompile",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator: 0,
							},
						},
					},
				),
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              warp.WarpDefaultQuorumNumerator,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "subnet genesis precompile non-default",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator: 50,
							},
						},
					},
				),
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              50,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "subnet upgrade precompile",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				UpgradeConfig: extras.UpgradeConfig{
					PrecompileUpgrades: []extras.PrecompileUpgrade{
						{
							Config: &warp.Config{
								QuorumNumerator: 0,
							},
						},
					},
				},
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              warp.WarpDefaultQuorumNumerator,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "subnet upgrade precompile non-default",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				UpgradeConfig: extras.UpgradeConfig{
					PrecompileUpgrades: []extras.PrecompileUpgrade{
						{
							Config: &warp.Config{
								QuorumNumerator: 50,
							},
						},
					},
				},
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              50,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "subnet multiple already activated upgrades",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				UpgradeConfig: extras.UpgradeConfig{
					PrecompileUpgrades: []extras.PrecompileUpgrade{
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &beforeCurrentBlockTime1,
								},
								QuorumNumerator: 50,
							},
						},
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &beforeCurrentBlockTime2,
								},
								QuorumNumerator:              60,
								RequirePrimaryNetworkSigners: true,
							},
						},
					},
				},
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              60,
				RequirePrimaryNetworkSigners: true,
			},
		},
		{
			name:                "require primary network signers",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator:              0,
								RequirePrimaryNetworkSigners: true,
							},
						},
					},
				),
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              warp.WarpDefaultQuorumNumerator,
				RequirePrimaryNetworkSigners: true,
			},
		},
		{
			name:                "require primary network signers explicit false",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator:              0,
								RequirePrimaryNetworkSigners: false,
							},
						},
					},
				),
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              warp.WarpDefaultQuorumNumerator,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "require primary network signers true in future",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator:              0,
								RequirePrimaryNetworkSigners: false,
							},
						},
					},
				),
				UpgradeConfig: extras.UpgradeConfig{
					PrecompileUpgrades: []extras.PrecompileUpgrade{
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &beforeCurrentBlockTime1,
								},
								QuorumNumerator:              80,
								RequirePrimaryNetworkSigners: false,
							},
						},
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &afterCurrentBlockTime,
								},
								QuorumNumerator:              67,
								RequirePrimaryNetworkSigners: true,
							},
						},
					},
				},
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              80,
				RequirePrimaryNetworkSigners: false,
			},
		},
		{
			name:                "upgrades listed out of order",
			blockchainID:        blockchainID,
			subnetID:            subnetID,
			getChainConfigCalls: 1,
			chainConfig: params.ChainConfigWithUpgradesJSON{
				ChainConfig: *params.WithExtra(
					&params.ChainConfig{},
					&extras.ChainConfig{
						GenesisPrecompiles: extras.Precompiles{
							warpConfigKey: &warp.Config{
								QuorumNumerator:              0,
								RequirePrimaryNetworkSigners: false,
							},
						},
					},
				),
				UpgradeConfig: extras.UpgradeConfig{
					PrecompileUpgrades: []extras.PrecompileUpgrade{
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &afterCurrentBlockTime,
								},
								QuorumNumerator:              80,
								RequirePrimaryNetworkSigners: true,
							},
						},
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &beforeCurrentBlockTime2,
								},
								QuorumNumerator:              90,
								RequirePrimaryNetworkSigners: true,
							},
						},
						{
							Config: &warp.Config{
								Upgrade: precompileconfig.Upgrade{
									BlockTimestamp: &beforeCurrentBlockTime1,
								},
								QuorumNumerator:              80,
								RequirePrimaryNetworkSigners: false,
							},
						},
					},
				},
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              90,
				RequirePrimaryNetworkSigners: true,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client := stubRPCClient{
				chainConfig:  &testCase.chainConfig,
				latestHeader: currentBlock.Header(),
			}

			subnetWarpConfig, err := getWarpConfig(&client)
			require.Equal(t, testCase.expectedError, err)
			expectedWarpConfig := warpConfigFromSubnetWarpConfig(*subnetWarpConfig)
			require.Equal(t, testCase.expectedWarpConfig, expectedWarpConfig)
		})
	}
}

func TestValidateSourceBlockchain(t *testing.T) {
	validSourceCfg := SourceBlockchain{
		BlockchainID: testBlockchainID,
		RPCEndpoint: basecfg.APIConfig{
			BaseURL: fmt.Sprintf("http://test.avax.network/ext/bc/%s/rpc", testBlockchainID),
		},
		WSEndpoint: basecfg.APIConfig{
			BaseURL: fmt.Sprintf("ws://test.avax.network/ext/bc/%s/ws", testBlockchainID),
		},
		SubnetID: testSubnetID,
		SupportedDestinations: []*SupportedDestination{
			{
				BlockchainID: testBlockchainID,
			},
		},
		MessageContracts: map[string]MessageProtocolConfig{
			testAddress: {
				MessageFormat: TELEPORTER.String(),
			},
		},
	}
	testCases := []struct {
		name                          string
		sourceSubnet                  func() SourceBlockchain
		destinationBlockchainIDs      []string
		expectError                   bool
		expectedSupportedDestinations []string
	}{
		{
			name:                          "valid source subnet; explicitly supported destination",
			sourceSubnet:                  func() SourceBlockchain { return validSourceCfg },
			destinationBlockchainIDs:      []string{testBlockchainID},
			expectError:                   false,
			expectedSupportedDestinations: []string{testBlockchainID},
		},
		{
			name: "valid source subnet; implicitly supported destination",
			sourceSubnet: func() SourceBlockchain {
				cfg := validSourceCfg
				cfg.SupportedDestinations = nil
				return cfg
			},
			destinationBlockchainIDs:      []string{testBlockchainID},
			expectError:                   false,
			expectedSupportedDestinations: []string{testBlockchainID},
		},
		{
			name:                          "valid source subnet; partially supported destinations",
			sourceSubnet:                  func() SourceBlockchain { return validSourceCfg },
			destinationBlockchainIDs:      []string{testBlockchainID, testBlockchainID2},
			expectError:                   false,
			expectedSupportedDestinations: []string{testBlockchainID},
		},
		{
			name:                          "valid source subnet; unsupported destinations",
			sourceSubnet:                  func() SourceBlockchain { return validSourceCfg },
			destinationBlockchainIDs:      []string{testBlockchainID2},
			expectError:                   true,
			expectedSupportedDestinations: []string{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			blockchainIDs := set.NewSet[string](len(testCase.destinationBlockchainIDs))
			for _, id := range testCase.destinationBlockchainIDs {
				blockchainIDs.Add(id)
			}

			sourceSubnet := testCase.sourceSubnet()
			res := sourceSubnet.Validate(&blockchainIDs)
			if testCase.expectError {
				require.Error(t, res)
			} else {
				require.NoError(t, res)
			}
			// check the supported destinations
			for _, idStr := range testCase.expectedSupportedDestinations {
				id, err := ids.FromString(idStr)
				require.NoError(t, err)
				require.True(t, func() bool {
					for _, dest := range sourceSubnet.SupportedDestinations {
						if dest.GetBlockchainID() == id {
							return true
						}
					}
					return false
				}())
			}
		})
	}
}

func TestCountSuppliedSubnets(t *testing.T) {
	config := Config{
		SourceBlockchains: []*SourceBlockchain{
			{
				SubnetID: "1",
			},
			{
				SubnetID: "2",
			},
			{
				SubnetID: "1",
			},
		},
	}
	require.Equal(t, 2, config.countSuppliedSubnets())
}

func TestInitializeTrackedSubnets(t *testing.T) {
	sourceSubnetID1 := ids.GenerateTestID()
	sourceSubnetID2 := ids.GenerateTestID()
	destSubnetID1 := ids.GenerateTestID()
	destSubnetID2 := ids.GenerateTestID()

	destBlockchainID1 := ids.GenerateTestID()
	destBlockchainID2 := ids.GenerateTestID()

	cfg := &Config{
		SourceBlockchains: []*SourceBlockchain{
			{
				subnetID: sourceSubnetID1,
				SupportedDestinations: []*SupportedDestination{
					&SupportedDestination{
						BlockchainID: destBlockchainID1.String(),
					},
				},
			},
			{
				subnetID: sourceSubnetID2,
				SupportedDestinations: []*SupportedDestination{
					&SupportedDestination{
						BlockchainID: destBlockchainID2.String(),
					},
				},
			},
		},
		DestinationBlockchains: []*DestinationBlockchain{
			{
				subnetID:     destSubnetID1,
				blockchainID: destBlockchainID1,
				warpConfig: WarpConfig{
					RequirePrimaryNetworkSigners: false,
				},
			},
			{
				subnetID:     destSubnetID2,
				blockchainID: destBlockchainID2,
				warpConfig: WarpConfig{
					RequirePrimaryNetworkSigners: true,
				},
			},
		},
	}

	err := cfg.initializeTrackedSubnets()
	require.NoError(t, err)

	expectedSubnets := set.NewSet[ids.ID](3)
	expectedSubnets.Add(sourceSubnetID1)
	expectedSubnets.Add(sourceSubnetID2)
	expectedSubnets.Add(destSubnetID1)

	require.True(t, expectedSubnets.Equals(cfg.GetTrackedSubnets()))
}

func TestConfigSanitization(t *testing.T) {
	testCases := []struct {
		name           string
		config         *Config
		expectedFields map[string]interface{}
		checkField     string
		expectedValue  interface{}
	}{
		{
			name: "sensitive fields are redacted",
			config: &Config{
				LogLevel:   "info",
				RedisURL:   "redis://user:pass@localhost:6379",
				TLSKeyPath: "/path/to/secret.key",
				APIPort:    8080,
			},
			expectedFields: map[string]interface{}{
				"log-level":    "info",
				"redis-url":    "[REDACTED]",
				"tls-key-path": "/path/to/secret.key",
				"api-port":     uint16(8080),
			},
		},
		{
			name: "non-sensitive fields are preserved",
			config: &Config{
				LogLevel:            "debug",
				StorageLocation:     "/tmp/storage",
				APIPort:             9090,
				MetricsPort:         3000,
				ProcessMissedBlocks: true,
				SignatureCacheSize:  1024,
				AllowPrivateIPs:     false,
			},
			checkField:    "log-level",
			expectedValue: "debug",
		},
		{
			name: "nested API config is sanitized",
			config: &Config{
				LogLevel: "info",
				PChainAPI: &basecfg.APIConfig{
					BaseURL: "http://node.example.com:9650",
					QueryParams: map[string]string{
						"api-key": "secret123",
						"timeout": "30s",
					},
					HTTPHeaders: map[string]string{
						"Authorization": "Bearer token123",
						"User-Agent":    "icm-relayer/1.0",
					},
				},
			},
		},
		{
			name: "empty and nil values handled correctly",
			config: &Config{
				LogLevel:   "info",
				RedisURL:   "",
				TLSKeyPath: "",
				PChainAPI:  nil,
				InfoAPI:    nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.config.sanitizeForLogging()
			require.IsType(t, map[string]interface{}{}, result)

			if tc.expectedFields != nil {
				for field, expected := range tc.expectedFields {
					actual, exists := result[field]
					require.True(t, exists, "Expected field %s to exist", field)
					require.Equal(t, expected, actual, "Field %s has unexpected value", field)
				}
			}

			if tc.checkField != "" {
				actual, exists := result[tc.checkField]
				require.True(t, exists, "Expected field %s to exist", tc.checkField)
				require.Equal(t, tc.expectedValue, actual)
			}

			// Verify sensitive fields are redacted if they exist and have values
			if tc.config.RedisURL != "" {
				require.Equal(t, "[REDACTED]", result["redis-url"])
			}
		})
	}
}

func TestSanitizeStruct(t *testing.T) {
	type TestStruct struct {
		PublicField    string `json:"public-field"`
		SensitiveField string `json:"sensitive-field" sensitive:"true"`
		NumericField   int    `json:"numeric-field"`
		unexported     string // Should be ignored
	}

	testStruct := TestStruct{
		PublicField:    "public-value",
		SensitiveField: "secret-value",
		NumericField:   42,
		unexported:     "ignored",
	}

	v := reflect.ValueOf(testStruct)
	structType := reflect.TypeOf(testStruct)
	result := sanitizeStruct(v, structType)

	require.Equal(t, "public-value", result["public-field"])
	require.Equal(t, "[REDACTED]", result["sensitive-field"])
	require.Equal(t, 42, result["numeric-field"])
	require.NotContains(t, result, "unexported")
}

func TestSanitizeSlice(t *testing.T) {
	type TestStruct struct {
		Value  string `json:"value"`
		Secret string `json:"secret" sensitive:"true"`
	}

	slice := []TestStruct{
		{Value: "value1", Secret: "secret1"},
		{Value: "value2", Secret: "secret2"},
	}

	v := reflect.ValueOf(slice)
	sliceType := reflect.TypeOf(slice)
	result := sanitizeSlice(v, sliceType)

	require.Len(t, result, 2)

	// Check first element
	firstElem, ok := result[0].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "value1", firstElem["value"])
	require.Equal(t, "[REDACTED]", firstElem["secret"])

	// Check second element
	secondElem, ok := result[1].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "value2", secondElem["value"])
	require.Equal(t, "[REDACTED]", secondElem["secret"])
}

// Every value of an operator-supplied string map is redacted regardless of its key: credential
// parameter and header names vary by RPC provider ("key", "token", ...) and cannot be enumerated.
// Keys are kept so the configured set remains visible.
func TestSanitizeMap(t *testing.T) {
	input := map[string]string{
		"api-key":       "secret123",
		"key":           "provider-key-not-on-any-denylist",
		"Authorization": "Bearer token",
		"timeout":       "30s",
		"user-agent":    "test-agent",
	}
	result, ok := sanitizeMap(reflect.ValueOf(input), reflect.TypeOf(input)).(map[string]any)
	require.True(t, ok, "Expected result to be a map[string]any")
	require.Len(t, result, len(input))
	for key := range input {
		require.Equal(t, redactedValue, result[key], "value for key %q must be redacted", key)
	}

	// Maps with non-string values are sanitized element-wise instead.
	nested := map[string]basecfg.APIConfig{
		"primary": {BaseURL: "https://rpc.example.com/v2/KEY", QueryParams: map[string]string{"key": "K"}},
	}
	nestedResult, ok := sanitizeMap(reflect.ValueOf(nested), reflect.TypeOf(nested)).(map[string]any)
	require.True(t, ok)
	primary, ok := nestedResult["primary"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "https://rpc.example.com/"+redactedValue, primary["base-url"])
	require.Equal(t, map[string]any{"key": redactedValue}, primary["query-parameters"])
}

func TestRedactURLCredentials(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"not a URL", "/path/to/storage", "/path/to/storage"},
		{"empty", "", ""},
		{"plain word", "info", "info"},
		{"bare host", "http://127.0.0.1:9650", "http://127.0.0.1:9650"},
		{"bare host with trailing slash", "https://node.example.com:9650/", "https://node.example.com:9650/"},
		{"Infura-style path key", "https://mainnet.infura.io/v3/0123456789abcdef", "https://mainnet.infura.io/[REDACTED]"},
		{"Alchemy-style path key", "wss://eth-mainnet.g.alchemy.com/v2/abcDEF", "wss://eth-mainnet.g.alchemy.com/[REDACTED]"},
		{"Avalanche node path", "http://127.0.0.1:9650/ext/bc/C/rpc", "http://127.0.0.1:9650/[REDACTED]"},
		{"query string key", "https://rpc.example.com/?key=abc", "https://rpc.example.com/[REDACTED]"},
		{"query without path", "https://rpc.example.com?key=abc", "https://rpc.example.com/[REDACTED]"},
		{"basic auth userinfo", "https://user:pass@rpc.example.com", "https://rpc.example.com/[REDACTED]"},
		{"userinfo and port", "https://user:pass@rpc.example.com:8545/", "https://rpc.example.com:8545/[REDACTED]"},
		{"fragment only", "https://rpc.example.com#frag", "https://rpc.example.com/[REDACTED]"},
		{"scheme but unparseable", "http://[::1", "[REDACTED]"},
		{"scheme without host", "file:///etc/passwd", "[REDACTED]"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, redactURLCredentials(tc.input))
		})
	}
}

// End-to-end check of the startup config log: endpoint URLs carrying credentials in the path,
// query string, or userinfo, and arbitrary query parameters and headers, must not appear in the
// sanitized output in any form.
func TestConfigSanitizationRedactsEndpointCredentials(t *testing.T) {
	const (
		pathKey   = "INFURA_PATH_KEY_0123456789"
		queryKey  = "PROVIDER_QUERY_KEY_abcdef"
		password  = "basic-auth-PASSWORD"
		headerVal = "Bearer HEADER_TOKEN_xyz"
		userAgent = "icm-relayer/1.0"
	)
	cfg := &Config{
		LogLevel: "info",
		PChainAPI: &basecfg.APIConfig{
			BaseURL:     "https://mainnet.infura.io/v3/" + pathKey,
			QueryParams: map[string]string{"key": queryKey, "timeout": "30s"},
			HTTPHeaders: map[string]string{"Authorization": headerVal, "User-Agent": userAgent},
		},
		InfoAPI: &basecfg.APIConfig{
			BaseURL: "https://operator:" + password + "@node.internal.example.com:9650/ext/info",
		},
	}

	sanitized := cfg.sanitizeForLogging()
	pchain, ok := sanitized["p-chain-api"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "https://mainnet.infura.io/"+redactedValue, pchain["base-url"])
	require.Equal(t, map[string]any{"key": redactedValue, "timeout": redactedValue}, pchain["query-parameters"])
	require.Equal(t,
		map[string]any{"Authorization": redactedValue, "User-Agent": redactedValue},
		pchain["http-headers"],
	)
	info, ok := sanitized["info-api"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "https://node.internal.example.com:9650/"+redactedValue, info["base-url"])

	// The same guarantee for what actually reaches the log sink.
	logged, err := json.Marshal(sanitized)
	require.NoError(t, err)
	for _, secret := range []string{pathKey, queryKey, password, headerVal, userAgent, "operator:", "/ext/info", "/v3/"} {
		require.NotContains(t, string(logged), secret)
	}
	require.Contains(t, string(logged), "mainnet.infura.io")
	require.Contains(t, string(logged), "node.internal.example.com:9650")

	// External EVM destinations use a plain string endpoint that flows through the same path.
	ext := ExternalEVMDestination{RPCEndpoint: "https://eth-mainnet.g.alchemy.com/v2/" + pathKey}
	extSanitized := sanitizeStruct(reflect.ValueOf(ext), reflect.TypeOf(ext))
	require.Equal(t, "https://eth-mainnet.g.alchemy.com/"+redactedValue, extSanitized["rpc-endpoint"])
}

func TestGetJSONTag(t *testing.T) {
	testCases := []struct {
		name        string
		tag         string
		expectedTag string
	}{
		{
			name:        "simple json tag",
			tag:         `json:"field-name"`,
			expectedTag: "field-name",
		},
		{
			name:        "json tag with omitempty",
			tag:         `json:"field-name,omitempty"`,
			expectedTag: "field-name",
		},
		{
			name:        "json tag with multiple options",
			tag:         `json:"field-name,omitempty,string"`,
			expectedTag: "field-name",
		},
		{
			name:        "no json tag",
			tag:         `mapstructure:"field-name"`,
			expectedTag: "TestField", // Should return field name
		},
		{
			name:        "json skip tag",
			tag:         `json:"-"`,
			expectedTag: "-",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			field := reflect.StructField{
				Name: "TestField",
				Tag:  reflect.StructTag(tc.tag),
			}
			result := getJSONTag(field)
			require.Equal(t, tc.expectedTag, result)
		})
	}
}
