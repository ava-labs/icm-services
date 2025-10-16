// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=./mocks/mock_app_request_network.go -package=mocks
//go:generate go run go.uber.org/mock/mockgen -destination=./avago_mocks/mock_network.go -package=avago_mocks github.com/ava-labs/avalanchego/network Network

package peers

import (
	"context"
	"crypto"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
	vdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/staking"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/sampler"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers/utils"
	"github.com/ava-labs/icm-services/peers/validators"
	subnetWarp "github.com/ava-labs/subnet-evm/precompile/contracts/warp"

	sharedUtils "github.com/ava-labs/icm-services/utils"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

const (
	InboundMessageChannelSize = 1000
	ValidatorRefreshPeriod    = time.Second * 5
	NumBootstrapNodes         = 5
	// Maximum number of subnets that can be tracked by the app request network
	// This value is defined in avalanchego peers package
	// TODO: use the avalanchego constant when it is exported
	maxNumSubnets = 16

	// The amount of time to cache canonical validator sets
	canonicalValidatorSetCacheTTL = 2 * time.Second

	// The size of the LRU cache for epoched validator sets per subnet
	// Each subnet gets its own LRU cache of (P-chain height -> validator set)
	// This handles both current validators (ProposedHeight) and epoched validators (specific heights)
	epochedValidatorSetCacheSizePerSubnet = 100
)

var _ AppRequestNetwork = (*appRequestNetwork)(nil)

var (
	ErrNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
	errTrackingTooManySubnets  = fmt.Errorf("cannot track more than %d subnets", maxNumSubnets)
)

type AppRequestNetwork interface {
	GetCanonicalValidators(ctx context.Context, subnetID ids.ID, skipCache bool, pchainHeight uint64) (
		*CanonicalValidators,
		error,
	)
	GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error)
	RegisterAppRequest(requestID ids.RequestID)
	RegisterRequestID(
		requestID uint32,
		requestedNodes set.Set[ids.NodeID],
	) chan message.InboundMessage
	Send(
		msg message.OutboundMessage,
		nodeIDs set.Set[ids.NodeID],
		subnetID ids.ID,
		allower subnets.Allower,
	) set.Set[ids.NodeID]
	Shutdown()
	TrackSubnet(ctx context.Context, subnetID ids.ID)
	GetNetworkID() uint32
	IsGraniteActivated() bool
}

type appRequestNetwork struct {
	networkID        uint32
	network          network.Network
	handler          *RelayerExternalHandler
	infoAPI          *InfoAPI
	logger           logging.Logger
	validatorSetLock *sync.Mutex
	validatorClient  validators.CanonicalValidatorState
	metrics          *AppRequestNetworkMetrics

	// The set of subnetIDs to track. Shared with the underlying Network object, so access
	// must be protected by the trackedSubnetsLock
	trackedSubnets set.Set[ids.ID]
	// invariant: members of lruSubnets should always be exactly the same as trackedSubnets
	// and the size of lruSubnets should be less than or equal to maxNumSubnets
	lruSubnets         *linked.Hashmap[ids.ID, interface{}]
	trackedSubnetsLock *sync.RWMutex

	manager                    vdrs.Manager
	canonicalValidatorSetCache *cache.TTLCache[ids.ID, snowVdrs.WarpSet]
	epochedValidatorSetCache   map[ids.ID]*cache.LRUCache[uint64, snowVdrs.WarpSet]
	epochedCacheLock           sync.RWMutex // protects epochedValidatorSetCache map

	networkUpgradeConfig *upgrade.Config
}

// NewNetwork creates a P2P network client for interacting with validators
func NewNetwork(
	ctx context.Context,
	logger logging.Logger,
	relayerRegistry prometheus.Registerer,
	peerNetworkRegistry prometheus.Registerer,
	timeoutManagerRegistry prometheus.Registerer,
	trackedSubnets set.Set[ids.ID],
	manuallyTrackedPeers []info.Peer,
	cfg Config,
) (AppRequestNetwork, error) {
	metrics := newAppRequestNetworkMetrics(relayerRegistry)

	// Create the handler for handling inbound app responses
	handler, err := NewRelayerExternalHandler(logger, metrics, timeoutManagerRegistry)
	if err != nil {
		logger.Error(
			"Failed to create p2p network handler",
			zap.Error(err),
		)
		return nil, err
	}

	infoAPI, err := NewInfoAPI(cfg.GetInfoAPI())
	if err != nil {
		logger.Error(
			"Failed to create info API",
			zap.Error(err),
		)
		return nil, err
	}
	networkID, err := infoAPI.GetNetworkID(ctx)
	if err != nil {
		logger.Error(
			"Failed to get network ID",
			zap.Error(err),
		)
		return nil, err
	}

	upgradeConfig, err := infoAPI.Upgrades(ctx)
	if err != nil {
		return nil, err
	}

	validatorClient := validators.NewCanonicalValidatorClient(logger, cfg.GetPChainAPI())
	manager := snowVdrs.NewManager()

	// Primary network must not be explicitly tracked so removing it prior to creating TestNetworkConfig
	trackedSubnets.Remove(constants.PrimaryNetworkID)
	if trackedSubnets.Len() > maxNumSubnets {
		return nil, errTrackingTooManySubnets
	}
	trackedSubnetsLock := new(sync.RWMutex)
	testNetworkConfig, err := network.NewTestNetworkConfig(
		peerNetworkRegistry,
		networkID,
		manager,
		trackedSubnets,
		trackedSubnetsLock,
	)
	if err != nil {
		logger.Error(
			"Failed to create test network config",
			zap.Error(err),
		)
		return nil, err
	}
	testNetworkConfig.AllowPrivateIPs = cfg.GetAllowPrivateIPs()
	// Set the TLS config if exists and log the NodeID
	var cert *tls.Certificate
	if cert = cfg.GetTLSCert(); cert != nil {
		testNetworkConfig.TLSConfig = peer.TLSConfig(*cert, nil)
		testNetworkConfig.TLSKey = cert.PrivateKey.(crypto.Signer)
	} else {
		cert = &testNetworkConfig.TLSConfig.Certificates[0]
	}
	parsedCert, err := staking.ParseCertificate(cert.Leaf.Raw)
	if err != nil {
		return nil, err
	}
	nodeID := ids.NodeIDFromCert(parsedCert)
	logger.Info("Network starting with NodeID", zap.Stringer("NodeID", nodeID))

	// Set the activation time for the latest network upgrade

	upgradeTime := upgrade.GetConfig(networkID).GraniteTime
	testNetwork, err := network.NewTestNetwork(
		logger,
		peerNetworkRegistry,
		testNetworkConfig,
		handler,
		upgradeTime,
	)
	if err != nil {
		logger.Error(
			"Failed to create test network",
			zap.Error(err),
		)
		return nil, err
	}

	for _, peer := range manuallyTrackedPeers {
		logger.Info(
			"Manually Tracking peer (startup)",
			zap.Stringer("ID", peer.ID),
			zap.Stringer("IP", peer.PublicIP),
		)
		testNetwork.ManuallyTrack(peer.ID, peer.PublicIP)
	}

	// Connect to a sample of the primary network validators, with connection
	// info pulled from the info API
	peers, err := infoAPI.Peers(context.Background(), nil)
	if err != nil {
		logger.Error(
			"Failed to get peers",
			zap.Error(err),
		)
		return nil, err
	}
	peersMap := make(map[ids.NodeID]info.Peer)
	for _, peer := range peers {
		peersMap[peer.ID] = peer
	}

	pClient := platformvm.NewClient(cfg.GetPChainAPI().BaseURL)
	options := utils.InitializeOptions(cfg.GetPChainAPI())
	vdrs, err := pClient.GetCurrentValidators(
		ctx,
		constants.PrimaryNetworkID,
		nil,
		options...,
	)
	if err != nil {
		logger.Error("Failed to get current validators", zap.Error(err))
		return nil, err
	}

	// Sample until we've connected to the target number of bootstrap nodes
	s := sampler.NewUniform()
	s.Initialize(uint64(len(vdrs)))
	numConnected := 0
	for numConnected < NumBootstrapNodes {
		i, ok := s.Next()
		if !ok {
			// If we've sampled all the nodes and still haven't connected to the target number of bootstrap nodes,
			// then warn and stop sampling by either returning an error or breaking
			logger.Warn(
				"Failed to connect to enough bootstrap nodes",
				zap.Int("targetBootstrapNodes", NumBootstrapNodes),
				zap.Int("numAvailablePeers", len(peers)),
				zap.Int("connectedBootstrapNodes", numConnected),
			)
			if numConnected == 0 {
				return nil, fmt.Errorf("failed to connect to any bootstrap nodes")
			}
			break
		}
		if peer, ok := peersMap[vdrs[i].NodeID]; ok {
			logger.Info(
				"Manually tracking bootstrap node",
				zap.Stringer("ID", peer.ID),
				zap.Stringer("IP", peer.PublicIP),
			)
			testNetwork.ManuallyTrack(peer.ID, peer.PublicIP)
			numConnected++
		}
	}

	go logger.RecoverAndPanic(func() {
		testNetwork.Dispatch()
	})
	lruSubnets := linked.NewHashmapWithSize[ids.ID, interface{}](maxNumSubnets)
	for _, subnetID := range trackedSubnets.List() {
		lruSubnets.Put(subnetID, nil)
	}
	vdrsCache := cache.NewTTLCache[ids.ID, snowVdrs.WarpSet](canonicalValidatorSetCacheTTL)

	arNetwork := &appRequestNetwork{
		networkID:                  networkID,
		network:                    testNetwork,
		handler:                    handler,
		infoAPI:                    infoAPI,
		logger:                     logger,
		validatorSetLock:           new(sync.Mutex),
		validatorClient:            validatorClient,
		metrics:                    metrics,
		trackedSubnets:             trackedSubnets,
		trackedSubnetsLock:         trackedSubnetsLock,
		manager:                    manager,
		lruSubnets:                 lruSubnets,
		canonicalValidatorSetCache: vdrsCache,
		epochedValidatorSetCache:   make(map[ids.ID]*cache.LRUCache[uint64, snowVdrs.WarpSet]),
		epochedCacheLock:           sync.RWMutex{},
		networkUpgradeConfig:       upgradeConfig,
	}

	go arNetwork.startUpdateValidators(ctx)

	return arNetwork, nil
}

func (n *appRequestNetwork) GetNetworkID() uint32 {
	return n.networkID
}

func (n *appRequestNetwork) IsGraniteActivated() bool {
	return n.networkUpgradeConfig.IsGraniteActivated(time.Now())
}

// Helper to scope lock acquisition
func (n *appRequestNetwork) trackSubnet(subnetID ids.ID) {
	n.trackedSubnetsLock.Lock()
	defer n.trackedSubnetsLock.Unlock()
	if n.trackedSubnets.Contains(subnetID) {
		// update the access to keep it in the LRU
		n.lruSubnets.Put(subnetID, nil)
		return
	}
	if n.lruSubnets.Len() >= maxNumSubnets {
		oldestSubnetID, _, _ := n.lruSubnets.Oldest()
		if !n.trackedSubnets.Contains(oldestSubnetID) {
			panic(fmt.Sprintf("SubnetID present in LRU but not in trackedSubnets: %s", oldestSubnetID))
		}
		n.trackedSubnets.Remove(oldestSubnetID)
		n.lruSubnets.Delete(oldestSubnetID)
		n.logger.Info("Removing LRU subnetID from tracked subnets", zap.Stringer("subnetID", oldestSubnetID))
	}
	n.logger.Info("Tracking subnet", zap.Stringer("subnetID", subnetID))
	n.lruSubnets.Put(subnetID, nil)
	n.trackedSubnets.Add(subnetID)
}

// TrackSubnet adds the subnet to the list of tracked subnets
// and initiates the connections to the subnet's validators asynchronously
func (n *appRequestNetwork) TrackSubnet(ctx context.Context, subnetID ids.ID) {
	n.trackSubnet(subnetID)
	n.updateValidatorSet(ctx, subnetID)
}

func (n *appRequestNetwork) startUpdateValidators(ctx context.Context) {
	// Fetch validators immediately when called, and refresh every ValidatorRefreshPeriod
	ticker := time.NewTicker(ValidatorRefreshPeriod)
	for {
		select {
		case <-ticker.C:
			n.trackedSubnetsLock.RLock()
			subnets := n.trackedSubnets.List()
			n.trackedSubnetsLock.RUnlock()

			n.logger.Debug(
				"Fetching validators for subnets",
				zap.Any("subnetIDs", append([]ids.ID{constants.PrimaryNetworkID}, subnets...)),
			)

			n.updateValidatorSet(ctx, constants.PrimaryNetworkID)
			for _, subnet := range subnets {
				n.updateValidatorSet(ctx, subnet)
			}

		case <-ctx.Done():
			n.logger.Info("Stopping updating validator process...")
			return
		}
	}
}

func (n *appRequestNetwork) updateValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) error {
	n.validatorSetLock.Lock()
	defer n.validatorSetLock.Unlock()

	// Fetch the subnet validators from the P-Chain
	cctx, cancel := context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	validators, err := n.validatorClient.GetProposedValidators(cctx, subnetID)
	if err != nil {
		return err
	}

	// Remove any elements from the manager that are not in the new validator set
	currentVdrs := n.manager.GetValidatorIDs(subnetID)
	for _, nodeID := range currentVdrs {
		if _, ok := validators[nodeID]; !ok {
			n.logger.Debug("Removing validator",
				zap.Stringer("nodeID", nodeID),
				zap.Stringer("subnetID", subnetID),
			)
			weight := n.manager.GetWeight(subnetID, nodeID)
			if err := n.manager.RemoveWeight(subnetID, nodeID, weight); err != nil {
				return err
			}
		}
	}

	// Add any elements from the new validator set that are not in the manager
	for _, vdr := range validators {
		if _, ok := n.manager.GetValidator(subnetID, vdr.NodeID); !ok {
			n.logger.Debug("Adding validator",
				zap.Stringer("nodeID", vdr.NodeID),
				zap.Stringer("subnetID", subnetID),
			)
			if err := n.manager.AddStaker(
				subnetID,
				vdr.NodeID,
				vdr.PublicKey,
				ids.Empty,
				vdr.Weight,
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *appRequestNetwork) Shutdown() {
	n.network.StartClose()
}

// Helper struct to hold connected validator information
// Warp Validators sharing the same BLS key may consist of multiple nodes,
// so we need to track the node ID to validator index mapping
type CanonicalValidators struct {
	ConnectedWeight uint64
	ConnectedNodes  set.Set[ids.NodeID]
	// ValidatorSet is the full canonical validator set for the subnet
	// and not only the connected nodes.
	ValidatorSet          snowVdrs.WarpSet
	NodeValidatorIndexMap map[ids.NodeID]int
}

// Returns the Warp Validator and its index in the canonical Validator ordering for a given nodeID
func (c *CanonicalValidators) GetValidator(nodeID ids.NodeID) (*warp.Validator, int) {
	return c.ValidatorSet.Validators[c.NodeValidatorIndexMap[nodeID]], c.NodeValidatorIndexMap[nodeID]
}

// GetCanonicalValidators returns the validator information in canonical ordering for the given subnet
// at the specified P-Chain height, as well as the total weight of the validators that this network is connected to
// The caller determines the appropriate P-Chain height (ProposedHeight for current, specific height for epoched)
func (n *appRequestNetwork) GetCanonicalValidators(
	ctx context.Context,
	subnetID ids.ID,
	skipCache bool,
	pchainHeight uint64,
) (*CanonicalValidators, error) {
	n.logger.Debug("Getting validator set at P-Chain height",
		zap.Stringer("subnetID", subnetID),
		zap.Uint64("pchainHeight", pchainHeight),
		zap.Bool("isProposedHeight", pchainHeight == pchainapi.ProposedHeight),
	)

	var validatorSet snowVdrs.WarpSet
	var err error

	if pchainHeight == pchainapi.ProposedHeight {
		// Get the subnet's current canonical validator set
		fetchVdrsFunc := func(subnetID ids.ID) (snowVdrs.WarpSet, error) {
			startPChainAPICall := time.Now()
			validatorSet, err := n.validatorClient.GetCurrentCanonicalValidatorSet(ctx, subnetID, pchainHeight)
			n.setPChainAPICallLatencyMS(time.Since(startPChainAPICall).Milliseconds())
			return validatorSet, err
		}
		validatorSet, err = n.canonicalValidatorSetCache.Get(subnetID, fetchVdrsFunc, skipCache)
		if err != nil {
			return nil, err
		}
	} else {
		// Use LRU cache for epoched validators (specific heights) - immutable historical data
		// LRU cache key is pchainHeight, fetch function uses the passed height
		epochedCache := n.getOrCreateEpochedCache(subnetID)
		fetchVdrsFunc := func(height uint64) (snowVdrs.WarpSet, error) {
			startPChainAPICall := time.Now()
			validatorSet, err := n.validatorClient.GetCurrentCanonicalValidatorSet(ctx, subnetID, height)
			n.setPChainAPICallLatencyMS(time.Since(startPChainAPICall).Milliseconds())
			return validatorSet, err
		}
		validatorSet, err = epochedCache.Get(pchainHeight, fetchVdrsFunc, skipCache)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get validator set at P-Chain height %d: %w", pchainHeight, err)
	}

	return n.buildCanonicalValidators(validatorSet), nil
}

// getOrCreateEpochedCache returns or creates the LRU cache for a specific subnet
func (n *appRequestNetwork) getOrCreateEpochedCache(
	subnetID ids.ID,
) *cache.LRUCache[uint64, snowVdrs.WarpSet] {
	n.epochedCacheLock.Lock()
	defer n.epochedCacheLock.Unlock()
	lruCache, exists := n.epochedValidatorSetCache[subnetID]
	if !exists {
		lruCache = cache.NewLRUCache[uint64, snowVdrs.WarpSet](epochedValidatorSetCacheSizePerSubnet)
		n.epochedValidatorSetCache[subnetID] = lruCache
	}

	return lruCache
}

// buildCanonicalValidators builds the CanonicalValidators struct from a validator set
func (n *appRequestNetwork) buildCanonicalValidators(
	validatorSet snowVdrs.WarpSet,
) *CanonicalValidators {
	// We make queries to node IDs, not unique validators as represented by a BLS pubkey, so we need this map to track
	// responses from nodes and populate the signatureMap with the corresponding validator signature
	// This maps node IDs to the index in the canonical validator set
	nodeValidatorIndexMap := make(map[ids.NodeID]int)
	nodeIDs := set.NewSet[ids.NodeID](len(nodeValidatorIndexMap))
	for i, vdr := range validatorSet.Validators {
		for _, node := range vdr.NodeIDs {
			nodeValidatorIndexMap[node] = i
			nodeIDs.Add(node)
		}
	}

	peerInfo := n.network.PeerInfo(nodeIDs.List())
	connectedPeers := set.NewSet[ids.NodeID](len(nodeIDs))
	for _, peer := range peerInfo {
		if nodeIDs.Contains(peer.ID) {
			connectedPeers.Add(peer.ID)
		}
	}

	// Calculate the total weight of connected validators.
	connectedWeight := calculateConnectedWeight(
		validatorSet.Validators,
		nodeValidatorIndexMap,
		connectedPeers,
	)

	return &CanonicalValidators{
		ConnectedWeight:       connectedWeight,
		ConnectedNodes:        connectedPeers,
		ValidatorSet:          validatorSet,
		NodeValidatorIndexMap: nodeValidatorIndexMap,
	}
}

func (n *appRequestNetwork) Send(
	msg message.OutboundMessage,
	nodeIDs set.Set[ids.NodeID],
	subnetID ids.ID,
	allower subnets.Allower,
) set.Set[ids.NodeID] {
	return n.network.Send(msg, common.SendConfig{NodeIDs: nodeIDs}, subnetID, allower)
}

func (n *appRequestNetwork) RegisterAppRequest(requestID ids.RequestID) {
	n.handler.RegisterAppRequest(requestID)
}

func (n *appRequestNetwork) RegisterRequestID(
	requestID uint32,
	requestedNodes set.Set[ids.NodeID],
) chan message.InboundMessage {
	return n.handler.RegisterRequestID(requestID, requestedNodes)
}

func (n *appRequestNetwork) GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error) {
	return n.validatorClient.GetSubnetID(ctx, blockchainID)
}

//
// Metrics
//

func (n *appRequestNetwork) setPChainAPICallLatencyMS(latency int64) {
	n.metrics.pChainAPICallLatencyMS.Observe(float64(latency))
}

// Non-receiver util functions

func GetNetworkHealthFunc(network AppRequestNetwork, subnetIDs []ids.ID) func(context.Context) error {
	return func(ctx context.Context) error {
		for _, subnetID := range subnetIDs {
			vdrs, err := network.GetCanonicalValidators(
				ctx,
				subnetID,
				false,
				pchainapi.ProposedHeight,
			)
			if err != nil {
				return fmt.Errorf("failed to get connected validators: %s, %w", subnetID, err)
			}
			if !sharedUtils.CheckStakeWeightExceedsThreshold(
				big.NewInt(0).SetUint64(vdrs.ConnectedWeight),
				vdrs.ValidatorSet.TotalWeight,
				subnetWarp.WarpDefaultQuorumNumerator,
			) {
				return ErrNotEnoughConnectedStake
			}
		}
		return nil
	}
}

func calculateConnectedWeight(
	validatorSet []*snowVdrs.Warp,
	nodeValidatorIndexMap map[ids.NodeID]int,
	connectedNodes set.Set[ids.NodeID],
) uint64 {
	connectedBLSPubKeys := set.NewSet[string](len(validatorSet))
	connectedWeight := uint64(0)
	for node := range connectedNodes {
		vdr := validatorSet[nodeValidatorIndexMap[node]]
		blsPubKey := hex.EncodeToString(vdr.PublicKeyBytes)
		if connectedBLSPubKeys.Contains(blsPubKey) {
			continue
		}
		connectedWeight += vdr.Weight
		connectedBLSPubKeys.Add(blsPubKey)
	}
	return connectedWeight
}
