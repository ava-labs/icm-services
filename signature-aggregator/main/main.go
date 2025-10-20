// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"

	metricsServer "github.com/ava-labs/icm-services/metrics"
	"github.com/ava-labs/icm-services/peers"
	peerUtils "github.com/ava-labs/icm-services/peers/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	"github.com/ava-labs/icm-services/signature-aggregator/config"
	"github.com/ava-labs/icm-services/signature-aggregator/healthcheck"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var version = "v0.0.0-dev"

const (
	sigAggMetricsPrefix         = "signature-aggregator"
	msgCreatorPrefix            = "msgcreator"
	timeoutManagerMetricsPrefix = "timeoutmanager"

	// The size of the FIFO cache for epoched validator sets
	// The Cache will store validator sets for the most recent N P-Chain heights.
	validatorSetCacheSize = 750
)

func main() {
	// Register all libevm extras in order to be able to get pre-compile information from the genesis block
	evm.RegisterAllLibEVMExtras()
	cfg := buildConfig()

	logLevel, err := logging.ToLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("error reading log level from config: %s", err)
	}

	logger := logging.NewLogger(
		"signature-aggregator",
		logging.NewWrappedCore(
			logLevel,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	logger.Info("Initializing signature-aggregator")

	// Initialize the global app request network
	logger.Info("Initializing app request network")
	// The app request network generates P2P networking logs that are verbose at the info level.
	// Unless the log level is debug or lower, set the network log level to error to avoid spamming the logs.
	// We do not collect metrics for the network.
	networkLogLevel := logging.Error
	if logLevel <= logging.Debug {
		networkLogLevel = logLevel
	}
	networkLogger := logging.NewLogger(
		"p2p-network",
		logging.NewWrappedCore(
			networkLogLevel,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	registries, err := metricsServer.StartMetricsServer(
		logger,
		cfg.MetricsPort,
		[]string{
			sigAggMetricsPrefix,
			msgCreatorPrefix,
			timeoutManagerMetricsPrefix,
		},
	)
	if err != nil {
		logger.Fatal("Failed to start metrics server", zap.Error(err))
		os.Exit(1)
	}

	// Initialize message creator passed down to relayers for creating app requests.
	// We do not collect metrics for the message creator.
	messageCreator, err := message.NewCreator(
		registries[msgCreatorPrefix],
		constants.DefaultNetworkCompressionType,
		constants.DefaultNetworkMaximumInboundTimeout,
	)
	if err != nil {
		logger.Fatal("Failed to create message creator", zap.Error(err))
		os.Exit(1)
	}

	var manuallyTrackedPeers []info.Peer
	for _, p := range cfg.ManuallyTrackedPeers {
		manuallyTrackedPeers = append(manuallyTrackedPeers, info.Peer{
			Info: peer.Info{
				PublicIP: p.GetIP(),
				ID:       p.GetID(),
			},
		})
	}

	// Create parent context with cancel function
	parentCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create errgroup with parent context
	errGroup, ctx := errgroup.WithContext(parentCtx)

	// Clamp the validator set cache size based on the max lookback if set. This will prevent thrashing the cache.
	var vdrCacheSize uint64
	if cfg.MaxPChainLookback > 0 && uint64(cfg.MaxPChainLookback) < validatorSetCacheSize {
		vdrCacheSize = uint64(cfg.MaxPChainLookback)
	} else {
		vdrCacheSize = validatorSetCacheSize
	}

	network, err := peers.NewNetwork(
		ctx,
		networkLogger,
		prometheus.DefaultRegisterer,
		prometheus.DefaultRegisterer,
		registries[timeoutManagerMetricsPrefix],
		cfg.GetTrackedSubnets(),
		manuallyTrackedPeers,
		&cfg,
		vdrCacheSize,
	)
	if err != nil {
		logger.Fatal("Failed to create app request network", zap.Error(err))
		os.Exit(1)
	}
	defer network.Shutdown()

	go network.StartCacheValidatorSets(ctx)

	metricsInstance := metrics.NewSignatureAggregatorMetrics(registries[sigAggMetricsPrefix])

	signatureAggregator, err := aggregator.NewSignatureAggregator(
		network,
		messageCreator,
		cfg.SignatureCacheSize,
		metricsInstance,
		platformvm.NewClient(cfg.GetPChainAPI().BaseURL),
		peerUtils.InitializeOptions(cfg.GetPChainAPI()),
	)
	if err != nil {
		logger.Fatal("Failed to create signature aggregator", zap.Error(err))
		os.Exit(1)
	}

	api.HandleAggregateSignaturesByRawMsgRequest(
		logger,
		metricsInstance,
		signatureAggregator,
	)

	healthCheckSubnets := cfg.GetTrackedSubnets().List()
	healthCheckSubnets = append(healthCheckSubnets, constants.PrimaryNetworkID)
	networkHealthcheckFunc := peers.GetNetworkHealthFunc(network, healthCheckSubnets)
	healthcheck.HandleHealthCheckRequest(networkHealthcheckFunc)

	errGroup.Go(func() error {
		httpServer := &http.Server{
			Addr: fmt.Sprintf(":%d", cfg.APIPort),
		}
		// Handle graceful shutdown
		go func() {
			<-ctx.Done()
			if err := httpServer.Shutdown(context.Background()); err != nil {
				logger.Error("Failed to shutdown server", zap.Error(err))
			}
		}()

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("Failed to start server: %w", err)
		}

		return nil
	})

	// Handle os signal
	errGroup.Go(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigChan
		logger.Info("Receive os signal", zap.String("signal", sig.String()))

		// Cancel the parent context
		// This will cascade to errgroup context
		cancel()

		// No error for graceful shutdown
		return nil
	})

	logger.Info("Initialization complete")

	if err := errGroup.Wait(); err != nil {
		logger.Fatal("Exited with error", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Exited gracefully")
}

// buildConfig parses the flags and builds the config
// Errors here should call log.Fatalf to exit the program
// since these errors are prior to building the logger struct
func buildConfig() config.Config {
	fs := config.BuildFlagSet()
	if err := fs.Parse(os.Args[1:]); err != nil {
		config.DisplayUsageText()
		log.Fatalf("Failed to parse flags: %s", err)
	}

	displayVersion, err := fs.GetBool(config.VersionKey)
	if err != nil {
		log.Fatalf("error reading %s flag: %s", config.VersionKey, err)
	}
	if displayVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	help, err := fs.GetBool(config.HelpKey)
	if err != nil {
		log.Fatalf("error reading %s flag value: %s", config.HelpKey, err)
	}
	if help {
		config.DisplayUsageText()
		os.Exit(0)
	}
	v, err := config.BuildViper(fs)
	if err != nil {
		log.Fatalf("couldn't configure flags: %s", err)
	}

	cfg, err := config.NewConfig(v)
	if err != nil {
		log.Fatalf("couldn't build config: %s", err)
	}
	return cfg
}
