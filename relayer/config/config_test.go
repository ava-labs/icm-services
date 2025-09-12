// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	basecfg "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	mock_ethclient "github.com/ava-labs/icm-services/vms/evm/mocks"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/params/extras"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
	blockchainID, err := ids.FromString("p433wpuXyJiDhyazPYyZMJeaoPSW76CBZ2x7wrVPLgvokotXz")
	require.NoError(t, err)
	subnetID, err := ids.FromString("2PsShLjrFFwR51DMcAh8pyuwzLn1Ym3zRhuXLTmLCR1STk2mL6")
	require.NoError(t, err)

	// beforeCurrentBlockTime := uint64(time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC).Unix())
	afterCurrentBlockTime := uint64(time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC).Unix())

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
						UpgradeConfig: extras.UpgradeConfig{
							PrecompileUpgrades: []extras.PrecompileUpgrade{
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
				),
			},
			expectedError: nil,
			expectedWarpConfig: WarpConfig{
				QuorumNumerator:              warp.WarpDefaultQuorumNumerator,
				RequirePrimaryNetworkSigners: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client := mock_ethclient.NewMockClient(gomock.NewController(t))
			gomock.InOrder(
				client.EXPECT().ChainConfig(gomock.Any()).Return(
					&testCase.chainConfig,
					nil,
				).Times(testCase.getChainConfigCalls),
				client.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).Return(
					currentBlock,
					nil,
				).Times(testCase.getChainConfigCalls),
			)

			subnetWarpConfig, err := getWarpConfig(client)
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
		VM:       "evm",
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
