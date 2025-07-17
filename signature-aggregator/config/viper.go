// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"fmt"
	"strings"

	commonConfig "github.com/ava-labs/icm-services/config"
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

	filename, err := fs.GetString(ConfigFileKey)
	if err != nil {
		DisplayUsageText()
		return nil, fmt.Errorf("failed to get config file flag: %w", err)
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
	v.SetDefault(APIPortKey, defaultAPIPort)
	v.SetDefault(MetricsPortKey, defaultMetricsPort)
	v.SetDefault(
		SignatureCacheSizeKey,
		DefaultSignatureCacheSize,
	)
}

// BuildConfig constructs the signature aggregator config using Viper.
// The following precedence order is used. Each item takes precedence over the item below it:
//  1. Flags
//  2. Config file
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

	if v.IsSet(commonConfig.TLSKeyPathKey) || v.IsSet(commonConfig.TLSCertPathKey) {
		cert, err := commonConfig.GetTLSCertFromFile(v)
		if err != nil {
			return cfg, fmt.Errorf("failed to initialize TLS certificate: %w", err)
		}
		cfg.tlsCert = cert
	}

	return cfg, nil
}
