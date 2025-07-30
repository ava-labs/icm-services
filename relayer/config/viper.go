// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"fmt"
	"os"
	"strings"

	commonConfig "github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func NewConfig(v *viper.Viper) (Config, error) {
	cfg, err := BuildConfig(v)
	if err != nil {
		return cfg, err
	}
	if err = cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("failed to validate configuration: %w", err)
	}
	return cfg, nil
}

// Build the viper instance. The config file must be provided via the command line flag or environment variable.
// All config keys may be provided via config file or environment variable.
func BuildViper(fs *pflag.FlagSet) (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()
	// Map flag names to env var names. Flags are capitalized, and hyphens are replaced with underscores.
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	var err error
	var filename = os.Getenv(ConfigFileEnvKey)
	if filename == "" {
		filename, err = fs.GetString(ConfigFileKey)
		if err != nil {
			return nil, fmt.Errorf("config file not set via flag or environment variable: %w", err)
		}
	}

	v.SetConfigFile(filename)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func SetDefaultConfigValues(v *viper.Viper) {
	v.SetDefault(LogLevelKey, defaultLogLevel)
	v.SetDefault(StorageLocationKey, defaultStorageLocation)
	v.SetDefault(ProcessMissedBlocksKey, defaultProcessMissedBlocks)
	v.SetDefault(APIPortKey, defaultAPIPort)
	v.SetDefault(MetricsPortKey, defaultMetricsPort)
	v.SetDefault(DBWriteIntervalSecondsKey, defaultIntervalSeconds)
	v.SetDefault(
		SignatureCacheSizeKey,
		defaultSignatureCacheSize,
	)
	v.SetDefault(InitialConnectionTimeoutSecondsKey, defaultInitialConnectionTimeoutSeconds)
	v.SetDefault(MaxConcurrentMessagesKey, defaultMaxConcurrentMessages)
}

// BuildConfig constructs the relayer config using Viper.
// The following precedence order is used. Each item takes precedence over the item below it:
//  1. Flags
//  2. Environment variables
//  3. Config file
//
// Returns the Config
func BuildConfig(v *viper.Viper) (Config, error) {
	// Set default values
	SetDefaultConfigValues(v)

	// Build the config from Viper
	var cfg Config

	if err := v.UnmarshalExact(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to unmarshal viper config: %w", err)
	}

	// Collect global private keys that will be available for all destination blockchains
	accountPrivateKeys := v.GetStringSlice(AccountPrivateKeysKey)
	accountPrivateKey := v.GetString(AccountPrivateKeyKey)
	if accountPrivateKey != "" {
		accountPrivateKeys = append(accountPrivateKeys, accountPrivateKey)
	}

	// Fetch destination chain private keys from the environment variables
	for i, subnet := range cfg.DestinationBlockchains {
		privateKeys := make([]string, len(accountPrivateKeys))
		copy(privateKeys, accountPrivateKeys)
		if privateKeyFromEnv := os.Getenv(fmt.Sprintf(
			"%s_%s",
			accountPrivateKeyEnvVarName,
			subnet.BlockchainID,
		)); privateKeyFromEnv != "" {
			privateKeys = append(privateKeys, privateKeyFromEnv)
		}
		if privateKeysFromEnv := os.Getenv(fmt.Sprintf(
			"%s_%s",
			accountPrivateKeyListEnvVarName,
			subnet.BlockchainID,
		)); privateKeysFromEnv != "" {
			privateKeys = append(privateKeys, strings.Split(privateKeysFromEnv, " ")...)
		}

		for i, privateKey := range privateKeys {
			privateKeys[i] = utils.SanitizeHexString(privateKey)
		}

		// Append keys provided via flag or environment variable to the config
		cfg.DestinationBlockchains[i].AccountPrivateKeys = append(
			cfg.DestinationBlockchains[i].AccountPrivateKeys,
			privateKeys...,
		)
	}

	if v.IsSet(commonConfig.TLSKeyPathKey) || v.IsSet(commonConfig.TLSCertPathKey) {
		cert, err := commonConfig.GetTLSCertFromFile(v)
		if err != nil {
			return cfg, fmt.Errorf("failed to initialize TLS certificate: %w", err)
		}
		cfg.tlsCert = cert
	}

	return cfg, nil
}
