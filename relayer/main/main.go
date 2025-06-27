// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	offchainregistry "github.com/ava-labs/icm-services/messages/off-chain-registry"
	"github.com/ava-labs/icm-services/messages/teleporter"
	metricsServer "github.com/ava-labs/icm-services/metrics"
	"github.com/ava-labs/icm-services/peers"
	peerUtils "github.com/ava-labs/icm-services/peers/utils"
	"github.com/ava-labs/icm-services/relayer"
	"github.com/ava-labs/icm-services/relayer/api"
	"github.com/ava-labs/icm-services/relayer/checkpoint"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	sigAggMetrics "github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Sets GOMAXPROCS to the CPU quota for containerized environments
	_ "go.uber.org/automaxprocs"
)

var version = "v0.0.0-dev"

const (
	relayerMetricsPrefix     = "app"
	peerNetworkMetricsPrefix = "peers"
)

func main() {
	cfg := buildConfig()

	// Modify the default http.DefaultClient globally
	// TODO: Remove this temporary fix once the RPC clients used by the relayer
	// start accepting custom underlying http clients.
	{
		// Set the timeout conservatively to catch any potential cases where the context is not used
		// and the request hangs indefinitely.
		http.DefaultClient.Timeout = 2 * utils.DefaultRPCTimeout
		maxConns := 10_000
		http.DefaultClient.Transport = &http.Transport{
			MaxConnsPerHost:     maxConns,
			MaxIdleConns:        maxConns,
			MaxIdleConnsPerHost: maxConns,
			IdleConnTimeout:     0, // Unlimited since handled by context and timeout on the client level.
		}
	}
	logLevel, err := logging.ToLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("error reading log level from config: %s", err)
	}
	logger := logging.NewLogger(
		"icm-relayer",
		logging.NewWrappedCore(
			logLevel,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)
	logger.Info("Initializing icm-relayer")

	// Initialize the Warp Config values and trackedSubnets by fetching via RPC
	// We do this here so that BuildConfig doesn't need to make RPC calls
	if err = cfg.Initialize(); err != nil {
		logger.Fatal("couldn't initialize config", zap.Error(err))
		os.Exit(1)
	}

	// Initialize all destination clients
	logger.Info("Initializing destination clients")
	destinationClients, err := vms.CreateDestinationClients(logger, cfg)
	if err != nil {
		logger.Fatal("Failed to create destination clients", zap.Error(err))
		os.Exit(1)
	}

	// Initialize all source clients
	logger.Info("Initializing source clients")
	sourceClients, err := createSourceClients(context.Background(), logger, &cfg)
	if err != nil {
		logger.Fatal("Failed to create source clients", zap.Error(err))
		os.Exit(1)
	}

	// Initialize metrics gathered through prometheus
	registries, err := metricsServer.StartMetricsServer(
		logger,
		cfg.MetricsPort,
		[]string{
			relayerMetricsPrefix,
			peerNetworkMetricsPrefix,
		},
	)
	if err != nil {
		logger.Fatal("Failed to start metrics server", zap.Error(err))
		os.Exit(1)
	}
	relayerMetricsRegistry := registries[relayerMetricsPrefix]
	peerNetworkMetricsRegistry := registries[peerNetworkMetricsPrefix]

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

	// Initialize message creator passed down to relayers for creating app requests.
	// We do not collect metrics for the message creator.
	messageCreator, err := message.NewCreator(
		prometheus.NewRegistry(), // isolate this from the rest of the metrics
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

	network, err := peers.NewNetwork(
		networkLogger,
		relayerMetricsRegistry,
		peerNetworkMetricsRegistry,
		cfg.GetTrackedSubnets(),
		manuallyTrackedPeers,
		&cfg,
	)
	if err != nil {
		logger.Fatal("Failed to create app request network", zap.Error(err))
		os.Exit(1)
	}
	defer network.Shutdown()

	err = relayer.InitializeConnectionsAndCheckStake(logger, network, &cfg)
	if err != nil {
		logger.Fatal("Failed to initialize connections and check stake", zap.Error(err))
		os.Exit(1)
	}

	relayerMetrics, err := relayer.NewApplicationRelayerMetrics(relayerMetricsRegistry)
	if err != nil {
		logger.Fatal("Failed to create application relayer metrics", zap.Error(err))
		os.Exit(1)
	}

	// Initialize the database
	db, err := database.NewDatabase(logger, &cfg)
	if err != nil {
		logger.Fatal("Failed to create database", zap.Error(err))
		os.Exit(1)
	}

	// Initialize the global write ticker
	ticker := utils.NewTicker(cfg.DBWriteIntervalSeconds)
	go ticker.Run()

	relayerHealth := createHealthTrackers(&cfg)

	deciderConnection, err := createDeciderConnection(cfg.DeciderURL)
	if err != nil {
		logger.Fatal(
			"Failed to instantiate decider connection",
			zap.Error(err),
		)
		os.Exit(1)
	}

	messageHandlerFactories, err := createMessageHandlerFactories(
		logger,
		&cfg,
		deciderConnection,
	)
	if err != nil {
		logger.Fatal("Failed to create message handler factories", zap.Error(err))
		os.Exit(1)
	}

	signatureAggregator, err := aggregator.NewSignatureAggregator(
		network,
		messageCreator,
		cfg.SignatureCacheSize,
		sigAggMetrics.NewSignatureAggregatorMetrics(
			relayerMetricsRegistry,
		),
		platformvm.NewClient(cfg.GetPChainAPI().BaseURL),
		peerUtils.InitializeOptions(cfg.GetPChainAPI()),
	)
	if err != nil {
		logger.Fatal("Failed to create signature aggregator", zap.Error(err))
		os.Exit(1)
	}

	// Limits the global number of messages that can be processed concurrently by the application
	// to avoid trying to issue too many requests at once.
	processMessageSemaphore := make(chan struct{}, cfg.MaxConcurrentMessages)

	applicationRelayers, minHeights, err := createApplicationRelayers(
		context.Background(),
		logger,
		relayerMetrics,
		db,
		ticker,
		network,
		&cfg,
		sourceClients,
		destinationClients,
		signatureAggregator,
		processMessageSemaphore,
	)
	if err != nil {
		logger.Fatal("Failed to create application relayers", zap.Error(err))
		os.Exit(1)
	}
	messageCoordinator := relayer.NewMessageCoordinator(
		logger,
		messageHandlerFactories,
		applicationRelayers,
		sourceClients,
	)

	networkHealthFunc := peers.GetNetworkHealthFunc(network, cfg.GetTrackedSubnets().List())

	// Each Listener goroutine will have an atomic bool that it can set to false to indicate an unrecoverable error
	api.HandleHealthCheck(logger, relayerHealth, networkHealthFunc)
	api.HandleRelay(logger, messageCoordinator)
	api.HandleRelayMessage(logger, messageCoordinator)

	// start the health check server
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.APIPort), nil)
		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Health check server closed")
		} else if err != nil {
			logger.Fatal("Health check server exited with error", zap.Error(err))
			os.Exit(1)
		}
	}()

	// Create listeners for each of the subnets configured as a source
	errGroup, ctx := errgroup.WithContext(context.Background())
	for _, s := range cfg.SourceBlockchains {
		sourceBlockchain := s

		// errgroup will cancel the context when the first goroutine returns an error
		errGroup.Go(func() error {
			// runListener runs until it errors or the context is canceled by another goroutine
			return relayer.RunListener(
				ctx,
				logger,
				*sourceBlockchain,
				sourceClients[sourceBlockchain.GetBlockchainID()],
				relayerHealth[sourceBlockchain.GetBlockchainID()],
				minHeights[sourceBlockchain.GetBlockchainID()],
				messageCoordinator,
			)
		})
	}
	logger.Info("Initialization complete")
	err = errGroup.Wait()
	logger.Error("Relayer exiting.", zap.Error(err))
}

// buildConfig parses the flags and builds the config
// Errors here should call log.Fatalf to exit the program
// since these errors are prior to building the logger struct
func buildConfig() config.Config {
	fs := config.BuildFlagSet()
	// Parse the flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		config.DisplayUsageText()
		log.Fatalf("couldn't parse flags: %s", err)
	}
	// If the version flag is set, display the version then exit
	displayVersion, err := fs.GetBool(config.VersionKey)
	if err != nil {
		log.Fatalf("error reading %s flag value: %s", config.VersionKey, err)
	}
	if displayVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}
	// If the help flag is set, output the usage text then exit
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

func createMessageHandlerFactories(
	logger logging.Logger,
	globalConfig *config.Config,
	deciderConnection *grpc.ClientConn,
) (map[ids.ID]map[common.Address]messages.MessageHandlerFactory, error) {
	messageHandlerFactories := make(map[ids.ID]map[common.Address]messages.MessageHandlerFactory)
	for _, sourceBlockchain := range globalConfig.SourceBlockchains {
		messageHandlerFactoriesForSource := make(map[common.Address]messages.MessageHandlerFactory)
		// Create message handler factories for each supported message protocol
		for addressStr, cfg := range sourceBlockchain.MessageContracts {
			address := common.HexToAddress(addressStr)
			format := cfg.MessageFormat
			var (
				m   messages.MessageHandlerFactory
				err error
			)
			switch config.ParseMessageProtocol(format) {
			case config.TELEPORTER:
				m, err = teleporter.NewMessageHandlerFactory(
					logger,
					address,
					cfg,
					deciderConnection,
				)
			case config.OFF_CHAIN_REGISTRY:
				m, err = offchainregistry.NewMessageHandlerFactory(
					logger,
					cfg,
				)
			default:
				m, err = nil, fmt.Errorf("invalid message format %s", format)
			}
			if err != nil {
				logger.Error("Failed to create message handler factory", zap.Error(err))
				return nil, err
			}
			messageHandlerFactoriesForSource[address] = m
		}
		messageHandlerFactories[sourceBlockchain.GetBlockchainID()] = messageHandlerFactoriesForSource
	}
	return messageHandlerFactories, nil
}

func createSourceClients(
	ctx context.Context,
	logger logging.Logger,
	cfg *config.Config,
) (map[ids.ID]ethclient.Client, error) {
	var err error
	clients := make(map[ids.ID]ethclient.Client)

	for _, sourceBlockchain := range cfg.SourceBlockchains {
		clients[sourceBlockchain.GetBlockchainID()], err = utils.NewEthClientWithConfig(
			ctx,
			sourceBlockchain.RPCEndpoint.BaseURL,
			sourceBlockchain.RPCEndpoint.HTTPHeaders,
			sourceBlockchain.RPCEndpoint.QueryParams,
		)
		if err != nil {
			logger.Error(
				"Failed to connect to node via RPC",
				zap.String("blockchainID", sourceBlockchain.BlockchainID),
				zap.Error(err),
			)
			return nil, err
		}
	}
	return clients, nil
}

// Returns a map of application relayers, as well as a map of source blockchain IDs to starting heights.
func createApplicationRelayers(
	ctx context.Context,
	logger logging.Logger,
	relayerMetrics *relayer.ApplicationRelayerMetrics,
	db database.RelayerDatabase,
	ticker *utils.Ticker,
	network peers.AppRequestNetwork,
	cfg *config.Config,
	sourceClients map[ids.ID]ethclient.Client,
	destinationClients map[ids.ID]vms.DestinationClient,
	signatureAggregator *aggregator.SignatureAggregator,
	processMessagesSemaphore chan struct{},
) (map[common.Hash]*relayer.ApplicationRelayer, map[ids.ID]uint64, error) {
	applicationRelayers := make(map[common.Hash]*relayer.ApplicationRelayer)
	minHeights := make(map[ids.ID]uint64)
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		currentHeight, err := sourceClients[sourceBlockchain.GetBlockchainID()].BlockNumber(ctx)
		if err != nil {
			logger.Error("Failed to get current block height", zap.Error(err))
			return nil, nil, err
		}

		// Create the ApplicationRelayers
		applicationRelayersForSource, minHeight, err := createApplicationRelayersForSourceChain(
			ctx,
			logger,
			relayerMetrics,
			db,
			ticker,
			*sourceBlockchain,
			network,
			cfg,
			currentHeight,
			destinationClients,
			signatureAggregator,
			processMessagesSemaphore,
		)
		if err != nil {
			logger.Error(
				"Failed to create application relayers",
				zap.String("blockchainID", sourceBlockchain.BlockchainID),
				zap.Error(err),
			)
			return nil, nil, err
		}

		for relayerID, applicationRelayer := range applicationRelayersForSource {
			applicationRelayers[relayerID] = applicationRelayer
		}
		minHeights[sourceBlockchain.GetBlockchainID()] = minHeight

		logger.Info(
			"Created application relayers",
			zap.String("blockchainID", sourceBlockchain.BlockchainID),
		)
	}
	return applicationRelayers, minHeights, nil
}

// createApplicationRelayersForSourceChain creates Application Relayers for a given source blockchain.
func createApplicationRelayersForSourceChain(
	ctx context.Context,
	logger logging.Logger,
	metrics *relayer.ApplicationRelayerMetrics,
	db database.RelayerDatabase,
	ticker *utils.Ticker,
	sourceBlockchain config.SourceBlockchain,
	network peers.AppRequestNetwork,
	cfg *config.Config,
	currentHeight uint64,
	destinationClients map[ids.ID]vms.DestinationClient,
	signatureAggregator *aggregator.SignatureAggregator,
	processMessageSemaphore chan struct{},
) (map[common.Hash]*relayer.ApplicationRelayer, uint64, error) {
	// Create the ApplicationRelayers
	logger.Info(
		"Creating application relayers",
		zap.String("originBlockchainID", sourceBlockchain.BlockchainID),
	)
	applicationRelayers := make(map[common.Hash]*relayer.ApplicationRelayer)

	// Each ApplicationRelayer determines its starting height based on the configuration and database state.
	// The Listener begins processing messages starting from the minimum height across all the ApplicationRelayers
	// If catch up is disabled, the first block the ApplicationRelayer processes is the next block after the current height
	var height, minHeight uint64
	if !cfg.ProcessMissedBlocks {
		logger.Info(
			"processed-missed-blocks set to false, starting processing from chain head",
			zap.String("blockchainID", sourceBlockchain.GetBlockchainID().String()),
		)
		height = currentHeight + 1
		minHeight = height
	}

	for _, relayerID := range database.GetSourceBlockchainRelayerIDs(&sourceBlockchain) {
		// Calculate the catch-up starting block height, and update the min height if necessary
		if cfg.ProcessMissedBlocks {
			var err error
			height, err = database.CalculateStartingBlockHeight(
				logger,
				db,
				relayerID,
				sourceBlockchain.ProcessHistoricalBlocksFromHeight,
				currentHeight,
			)
			if err != nil {
				logger.Error(
					"Failed to calculate starting block height",
					zap.String("relayerID", relayerID.ID.String()),
					zap.Error(err),
				)
				return nil, 0, err
			}

			// Update the min height. This is the height that the listener will start processing from
			if minHeight == 0 || height < minHeight {
				minHeight = height
			}
		}

		checkpointManager := checkpoint.NewCheckpointManager(
			logger,
			db,
			ticker.Subscribe(),
			relayerID,
			height,
		)

		applicationRelayer, err := relayer.NewApplicationRelayer(
			logger,
			metrics,
			network,
			relayerID,
			destinationClients[relayerID.DestinationBlockchainID],
			sourceBlockchain,
			checkpointManager,
			cfg,
			signatureAggregator,
			processMessageSemaphore,
		)
		if err != nil {
			logger.Error(
				"Failed to create application relayer",
				zap.String("relayerID", relayerID.ID.String()),
				zap.Error(err),
			)
			return nil, 0, err
		}
		applicationRelayers[relayerID.ID] = applicationRelayer

		logger.Info(
			"Created application relayer",
			zap.String("relayerID", relayerID.ID.String()),
			zap.String("sourceBlockchainID", relayerID.SourceBlockchainID.String()),
			zap.String("destinationBlockchainID", relayerID.DestinationBlockchainID.String()),
			zap.String("originSenderAddress", relayerID.OriginSenderAddress.String()),
			zap.String("destinationAddress", relayerID.DestinationAddress.String()),
		)
	}
	return applicationRelayers, minHeight, nil
}

// create a connection to the "should send message" decider service.
// if url is unspecified, returns a nil client pointer
func createDeciderConnection(url string) (*grpc.ClientConn, error) {
	if len(url) == 0 {
		return nil, nil
	}

	connection, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"Failed to instantiate grpc client: %w",
			err,
		)
	}

	runtime.SetFinalizer(
		connection,
		func(c *grpc.ClientConn) { c.Close() },
	)

	return connection, nil
}

func createHealthTrackers(cfg *config.Config) map[ids.ID]*atomic.Bool {
	healthTrackers := make(map[ids.ID]*atomic.Bool, len(cfg.SourceBlockchains))
	for _, sourceBlockchain := range cfg.SourceBlockchains {
		healthTrackers[sourceBlockchain.GetBlockchainID()] = atomic.NewBool(true)
	}
	return healthTrackers
}
