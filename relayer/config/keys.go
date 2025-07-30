// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

const (
	// Command line option keys
	ConfigFileKey = "config-file"
	VersionKey    = "version"
	HelpKey       = "help"

	// Environment variable keys
	ConfigFileEnvKey = "CONFIG_FILE"

	// Top-level configuration keys
	LogLevelKey                        = "log-level"
	APIPortKey                         = "api-port"
	MetricsPortKey                     = "metrics-port"
	AccountPrivateKeyKey               = "account-private-key"
	AccountPrivateKeysKey              = "account-private-keys-list"
	StorageLocationKey                 = "storage-location"
	ProcessMissedBlocksKey             = "process-missed-blocks"
	DBWriteIntervalSecondsKey          = "db-write-interval-seconds"
	SignatureCacheSizeKey              = "signature-cache-size"
	InitialConnectionTimeoutSecondsKey = "initial-connection-timeout-seconds"
	MaxConcurrentMessagesKey           = "max-concurrent-messages"
)
