// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

const (
	// Command line option keys
	ConfigFileKey = "config-file"
	VersionKey    = "version"
	HelpKey       = "help"

	// Top-level configuration keys
	LogLevelKey                        = "log-level"
	PChainAPIKey                       = "p-chain-api"
	InfoAPIKey                         = "info-api"
	APIPortKey                         = "api-port"
	MetricsPortKey                     = "metrics-port"
	SourceBlockchainsKey               = "source-blockchains"
	DestinationBlockchainsKey          = "destination-blockchains"
	AccountPrivateKeyKey               = "account-private-key"
	AccountPrivateKeysKey              = "account-private-keys-list"
	StorageLocationKey                 = "storage-location"
	RedisURLKey                        = "redis-url"
	ProcessMissedBlocksKey             = "process-missed-blocks"
	ManualWarpMessagesKey              = "manual-warp-messages"
	DBWriteIntervalSecondsKey          = "db-write-interval-seconds"
	SignatureCacheSizeKey              = "signature-cache-size"
	InitialConnectionTimeoutSecondsKey = "initial-connection-timeout-seconds"
	MaxConcurrentMessagesKey           = "max-concurrent-messages"
)
