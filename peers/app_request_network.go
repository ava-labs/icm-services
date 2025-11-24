// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

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
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
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
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/peers/utils"
	sharedUtils "github.com/ava-labs/icm-services/utils"
)

const (
	InboundMessageChannelSize = 1000
	ValidatorRefreshPeriod    = time.Minute * 1
	ValidatorPreFetchPeriod   = time.Second * 5
	NumBootstrapNodes         = 5
	// Maximum number of subnets that can be tracked by the app request network
	// This value is defined in avalanchego peers package
	// TODO: use the avalanchego constant when it is exported
	maxNumSubnets = 16

	// The amount of time to cache canonical validator sets
	canonicalValidatorSetCacheTTL = 2 * time.Second
)

var (
	ErrNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
	errTrackingTooManySubnets  = fmt.Errorf("cannot track more than %d subnets", maxNumSubnets)
)

type AppRequestNetwork struct {
	network          network.Network
	handler          *RelayerExternalHandler
	infoAPI          *clients.InfoAPI
	logger           logging.Logger
	validatorSetLock *sync.Mutex
	validatorClient  clients.CanonicalValidatorState
	metrics          *AppRequestNetworkMetrics

	// The set of subnetIDs to track. Shared with the underlying Network object, so access
	// must be protected by the trackedSubnetsLock
	trackedSubnets set.Set[ids.ID]
	// invariant: members of lruSubnets should always be exactly the same as trackedSubnets
	// and the size of lruSubnets should be less than or equal to maxNumSubnets
	lruSubnets         *linked.Hashmap[ids.ID, interface{}]
	trackedSubnetsLock *sync.RWMutex

	latestSyncedPChainHeight atomic.Uint64
	// Used by the signature aggregator to limit how far back in P-Chain history it will look
	maxPChainLookback int64

	manager                    snowVdrs.Manager
	canonicalValidatorSetCache *cache.TTLCache[ids.ID, snowVdrs.WarpSet]
	epochedValidatorSetCache   *cache.FIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet]

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
	validatorSetsCacheSize uint64,
) (*AppRequestNetwork, error) {
	metrics := newAppRequestNetworkMetrics(relayerRegistry)

	// Create the handler for handling inbound app responses
	handler, err := NewRelayerExternalHandler(logger, metrics, timeoutManagerRegistry)
	if err != nil {
		return nil, fmt.Errorf("failed to create p2p network handler: %w", err)
	}

	infoAPI, err := clients.NewInfoAPI(cfg.GetInfoAPI())
	if err != nil {
		return nil, fmt.Errorf("failed to create info API: %w", err)
	}
	networkID, err := infoAPI.GetNetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %w", err)
	}

	upgradeConfig, err := infoAPI.Upgrades(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get upgrades: %w", err)
	}

	validatorClient := clients.NewCanonicalValidatorClient(cfg.GetPChainAPI())
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
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create test network config: %w", err)
	}
	testNetworkConfig.AllowPrivateIPs = cfg.GetAllowPrivateIPs()
	testNetworkConfig.ConnectToAllValidators = true
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
		return nil, fmt.Errorf("failed to parse cert: %w", err)
	}
	nodeID := ids.NodeIDFromCert(parsedCert)
	logger.Info("Network starting with NodeID", zap.Stringer("NodeID", nodeID))

	testNetwork, err := network.NewTestNetwork(
		logger,
		peerNetworkRegistry,
		testNetworkConfig,
		handler,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create test network: %w", err)
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
	peers, err := infoAPI.Peers(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get peers: %w", err)
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
		return nil, fmt.Errorf("failed to get current validators: %w", err)
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
	epochedVdrsCache := cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](int(validatorSetsCacheSize))

	localTrackedSubnets := set.NewSet[ids.ID](maxNumSubnets)

	for _, subnetID := range trackedSubnets.List() {
		localTrackedSubnets.Add(subnetID)
	}

	arNetwork := &AppRequestNetwork{
		network:                    testNetwork,
		handler:                    handler,
		infoAPI:                    infoAPI,
		logger:                     logger,
		validatorSetLock:           new(sync.Mutex),
		validatorClient:            validatorClient,
		metrics:                    metrics,
		trackedSubnets:             localTrackedSubnets,
		trackedSubnetsLock:         trackedSubnetsLock,
		manager:                    manager,
		lruSubnets:                 lruSubnets,
		canonicalValidatorSetCache: vdrsCache,
		epochedValidatorSetCache:   epochedVdrsCache,
		maxPChainLookback:          cfg.GetMaxPChainLookback(),
		networkUpgradeConfig:       upgradeConfig,
		// latestSyncedPChainHeight is initialized to 0 by default (atomic.Uint64 zero value)
	}

	go arNetwork.startUpdateTrackedValidators(ctx)

	return arNetwork, nil
}

func (n *AppRequestNetwork) IsGraniteActivated() bool {
	return n.networkUpgradeConfig.IsGraniteActivated(time.Now())
}

// GetLatestSyncedPChainHeight returns the highest P-Chain height that has been successfully cached.
func (n *AppRequestNetwork) GetLatestSyncedPChainHeight() uint64 {
	return n.latestSyncedPChainHeight.Load()
}

// GetGraniteEpochDuration returns the Granite epoch duration from the network upgrade config.
// Returns 0 if Granite is not activated or epoch duration is not configured.
func (n *AppRequestNetwork) GetGraniteEpochDuration() time.Duration {
	if n.networkUpgradeConfig == nil {
		return 0
	}
	return n.networkUpgradeConfig.GraniteEpochDuration
}

// trackSubnet adds the subnetID to the set of tracked subnets. Returns true iff the subnet was already being tracked.
func (n *AppRequestNetwork) trackSubnet(subnetID ids.ID) bool {
	n.trackedSubnetsLock.Lock()
	defer n.trackedSubnetsLock.Unlock()
	if n.trackedSubnets.Contains(subnetID) {
		// update the access to keep it in the LRU
		n.lruSubnets.Put(subnetID, nil)
		return true
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
	return false
}

// TrackSubnet adds the subnet to the list of tracked subnets
// and initiates the connections to the subnet's validators asynchronously
func (n *AppRequestNetwork) TrackSubnet(ctx context.Context, subnetID ids.ID) {
	// Track the subnet. Update the validator set if we weren't already tracking it.
	if !n.trackSubnet(subnetID) {
		n.updateTrackedValidatorSet(ctx, subnetID)
	}
}

func (n *AppRequestNetwork) startUpdateTrackedValidators(ctx context.Context) {
	// Fetch validators immediately when called, and refresh every ValidatorRefreshPeriod
	ticker := time.NewTicker(ValidatorRefreshPeriod)
	n.updateTrackedValidatorSets(ctx)

	for {
		select {
		case <-ticker.C:
			n.updateTrackedValidatorSets(ctx)
		case <-ctx.Done():
			n.logger.Info("Stopping updating validator process...")
			return
		}
	}
}

func (n *AppRequestNetwork) StartCacheValidatorSets(ctx context.Context) {
	// Fetch validators immediately when called, and refresh every ValidatorRefreshPeriod
	ticker := time.NewTicker(ValidatorPreFetchPeriod)
	n.cacheMostRecentValidatorSets(ctx)

	for {
		select {
		case <-ticker.C:
			n.cacheMostRecentValidatorSets(ctx)
		case <-ctx.Done():
			n.logger.Info("Stopping caching validator process...")
			return
		}
	}
}

func (n *AppRequestNetwork) cacheMostRecentValidatorSets(ctx context.Context) {
	latestPChainHeight, err := n.validatorClient.GetLatestHeight(ctx)
	if err != nil {
		// This is not a critical error, just log and return
		n.logger.Error("Failed to get P-Chain height", zap.Error(err))
		return
	}

	currentSyncedHeight := n.latestSyncedPChainHeight.Load()
	if currentSyncedHeight == 0 {
		// Setting the current synced height to be one less than the latest P-Chain upon initialization makes it
		// such that we only fetch the validator sets at the latest P-Chain height to start.
		currentSyncedHeight = latestPChainHeight - 1
		n.latestSyncedPChainHeight.Store(currentSyncedHeight)
		n.logger.Info("Initializing P-Chain height", zap.Uint64("height", currentSyncedHeight))
	}

	for currentSyncedHeight < latestPChainHeight {
		currentSyncedHeight++
		// GetAllValidatorSets will update latestSyncedPChainHeight after successful cache
		_, err := n.GetAllValidatorSets(ctx, currentSyncedHeight)
		// If we fail to get the validator sets for this height, log and check the next height.
		if err != nil {
			n.logger.Error("Failed to get canonical validators",
				zap.Uint64("height", currentSyncedHeight),
				zap.Error(err),
			)
			continue
		}
	}
}

func (n *AppRequestNetwork) updateTrackedValidatorSets(ctx context.Context) {
	cctx, cancel := context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	latestPChainHeight, err := n.validatorClient.GetLatestHeight(cctx)
	if err != nil {
		n.logger.Warn("Failed to get latest P-Chain height", zap.Error(err))
		return
	}
	cctx, cancel = context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	allValidators, err := n.GetAllValidatorSets(cctx, latestPChainHeight)
	// If we fail to get the validator sets, log and return
	if err != nil {
		n.logger.Warn("Failed to get latest validators", zap.Error(err))
		return
	}

	n.trackedSubnetsLock.RLock()
	subnets := append(n.trackedSubnets.List(), constants.PrimaryNetworkID)
	n.trackedSubnetsLock.RUnlock()

	// Update the validators for each tracked subnet for the most recent height
	for _, subnetID := range subnets {
		vdrs, ok := allValidators[subnetID]
		if !ok {
			n.logger.Warn("No validator set found for tracked subnet",
				zap.Stringer("subnetID", subnetID),
				zap.Uint64("pchainHeight", n.latestSyncedPChainHeight.Load()),
			)
			continue
		}
		// If we fail to get the validator sets for this subnet, log and continue to the next subnet
		err := n.updatedTrackedValidators(subnetID, vdrs)
		if err != nil {
			n.logger.Error("Failed to update tracked validators",
				zap.Stringer("subnetID", subnetID),
				zap.Error(err),
			)
		}
	}
}

// Update the tracked validators for a single subnet. This is used when tracking a new subnet for the first time.
func (n *AppRequestNetwork) updateTrackedValidatorSet(
	ctx context.Context,
	subnetID ids.ID,
) error {
	cctx, cancel := context.WithTimeout(ctx, sharedUtils.DefaultRPCTimeout)
	defer cancel()
	vdrs, err := n.validatorClient.GetProposedValidators(cctx, subnetID)
	if err != nil {
		return err
	}

	return n.updatedTrackedValidators(subnetID, vdrs)
}

func (n *AppRequestNetwork) updatedTrackedValidators(
	subnetID ids.ID,
	vdrs snowVdrs.WarpSet,
) error {
	n.validatorSetLock.Lock()
	defer n.validatorSetLock.Unlock()

	nodeIDs := clients.NodeIDs(vdrs)

	// Remove any elements from the manager that are not in the new validator set
	currentVdrs := n.manager.GetValidatorIDs(subnetID)
	for _, nodeID := range currentVdrs {
		if !nodeIDs.Contains(nodeID) {
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
	for _, vdr := range vdrs.Validators {
		for _, nodeID := range vdr.NodeIDs {
			if _, ok := n.manager.GetValidator(subnetID, nodeID); !ok {
				n.logger.Debug("Adding validator",
					zap.Stringer("nodeID", nodeID),
					zap.Stringer("subnetID", subnetID),
				)
				if err := n.manager.AddStaker(
					subnetID,
					nodeID,
					vdr.PublicKey,
					ids.Empty,
					vdr.Weight,
				); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (n *AppRequestNetwork) Shutdown() {
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
func (c *CanonicalValidators) GetValidator(nodeID ids.NodeID) (*snowVdrs.Warp, int) {
	return c.ValidatorSet.Validators[c.NodeValidatorIndexMap[nodeID]], c.NodeValidatorIndexMap[nodeID]
}

func (n *AppRequestNetwork) getValidatorSetGranite(
	ctx context.Context,
	subnetID ids.ID,
	pchainHeight uint64,
) (snowVdrs.WarpSet, error) {
	allValidators, err := n.GetAllValidatorSets(ctx, pchainHeight)
	if err != nil {
		return snowVdrs.WarpSet{}, fmt.Errorf("failed to get all validators at P-Chain height %d: %w", pchainHeight, err)
	}

	validatorSet, ok := allValidators[subnetID]
	if !ok {
		return snowVdrs.WarpSet{}, fmt.Errorf("no validators for subnet %s at P-Chain height %d", subnetID, pchainHeight)
	}
	return validatorSet, nil
}

func (n *AppRequestNetwork) GetAllValidatorSets(
	ctx context.Context,
	pchainHeight uint64,
) (map[ids.ID]snowVdrs.WarpSet, error) {
	// If we're getting the proposed height, bypass the cache and get the latest data
	// We can't cache this call because we don't know the actual P-Chain height being returned.
	if pchainHeight == pchainapi.ProposedHeight {
		return n.validatorClient.GetAllValidatorSets(ctx, pchainHeight)
	}

	// Use FIFO cache for epoched validators (specific heights) - immutable historical data
	// FIFO cache key is pchainHeight, fetch function uses the passed height
	fetchVdrsFunc := func(height uint64) (map[ids.ID]snowVdrs.WarpSet, error) {
		latestSyncedHeight := n.latestSyncedPChainHeight.Load()
		if n.maxPChainLookback >= 0 && int64(height) < int64(latestSyncedHeight)-n.maxPChainLookback {
			return nil, fmt.Errorf("requested P-Chain height %d is beyond the max lookback of %d from latest height %d",
				height, n.maxPChainLookback, latestSyncedHeight,
			)
		}

		n.logger.Debug("Fetching all canonical validator sets at P-Chain height", zap.Uint64("pchainHeight", height))
		startPChainAPICall := time.Now()
		validatorSet, err := n.validatorClient.GetAllValidatorSets(ctx, height)
		n.setPChainAPICallLatencyMS(time.Since(startPChainAPICall).Milliseconds())
		return validatorSet, err
	}

	validatorSets, err := n.epochedValidatorSetCache.Get(pchainHeight, fetchVdrsFunc)
	if err != nil {
		return nil, err
	}

	// If the fetch succeeded, the set is in the cache now so update the latest synced height if greater
	// than the current latest synced height using atomic compare-and-swap
	for {
		current := n.latestSyncedPChainHeight.Load()
		if pchainHeight <= current {
			break
		}
		if n.latestSyncedPChainHeight.CompareAndSwap(current, pchainHeight) {
			break
		}
		// CAS failed, another goroutine updated it, retry
	}

	return validatorSets, nil
}

// GetCanonicalValidators returns the validator information in canonical ordering for the given subnet
// at the specified P-Chain height, as well as the total weight of the validators that this network is connected to
// The caller determines the appropriate P-Chain height (ProposedHeight for current, specific height for epoched)
func (n *AppRequestNetwork) GetCanonicalValidators(
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
			validatorSet, err := n.validatorClient.GetProposedValidators(ctx, subnetID)
			n.setPChainAPICallLatencyMS(time.Since(startPChainAPICall).Milliseconds())
			if err != nil {
				return snowVdrs.WarpSet{}, err
			}
			return validatorSet, nil
		}
		validatorSet, err = n.canonicalValidatorSetCache.Get(subnetID, fetchVdrsFunc, skipCache)
	} else {
		validatorSet, err = n.getValidatorSetGranite(ctx, subnetID, pchainHeight)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get validator set at P-Chain height %d: %w", pchainHeight, err)
	}

	return n.BuildCanonicalValidators(validatorSet), nil
}

// BuildCanonicalValidators builds the CanonicalValidators struct from a validator set
func (n *AppRequestNetwork) BuildCanonicalValidators(
	validatorSet snowVdrs.WarpSet,
) *CanonicalValidators {
	// We make queries to node IDs, not unique validators as represented by a BLS pubkey, so we need this map to track
	// responses from nodes and populate the signatureMap with the corresponding validator signature
	// This maps node IDs to the index in the canonical validator set
	nodeValidatorIndexMap := make(map[ids.NodeID]int)
	nodeIDs := set.NewSet[ids.NodeID](len(validatorSet.Validators))
	for i, vdr := range validatorSet.Validators {
		for _, nodeID := range vdr.NodeIDs {
			nodeValidatorIndexMap[nodeID] = i
			nodeIDs.Add(nodeID)
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

func (n *AppRequestNetwork) Send(
	msg message.OutboundMessage,
	nodeIDs set.Set[ids.NodeID],
	subnetID ids.ID,
	allower subnets.Allower,
) set.Set[ids.NodeID] {
	return n.network.Send(msg, common.SendConfig{NodeIDs: nodeIDs}, subnetID, allower)
}

func (n *AppRequestNetwork) RegisterAppRequest(requestID ids.RequestID) {
	n.handler.RegisterAppRequest(requestID)
}

func (n *AppRequestNetwork) RegisterRequestID(
	requestID uint32,
	requestedNodes set.Set[ids.NodeID],
) chan message.InboundMessage {
	return n.handler.RegisterRequestID(requestID, requestedNodes)
}

func (n *AppRequestNetwork) GetSubnetID(ctx context.Context, blockchainID ids.ID) (ids.ID, error) {
	return n.validatorClient.GetSubnetID(ctx, blockchainID)
}

//
// Metrics
//

func (n *AppRequestNetwork) setPChainAPICallLatencyMS(latency int64) {
	n.metrics.pChainAPICallLatencyMS.Observe(float64(latency))
}

// Non-receiver util functions

func GetNetworkHealthFunc(
	logger logging.Logger,
	network *AppRequestNetwork,
	subnetIDs []ids.ID,
) func(context.Context) error {
	return func(ctx context.Context) error {
		cachedHeight := network.GetLatestSyncedPChainHeight()
		if cachedHeight == 0 {
			// This should only happen at startup when the cache is not yet initialized.
			logger.Info("No cached P-Chain height, skipping network health check")
			return nil
		}

		allValidatorSets, err := network.GetAllValidatorSets(
			ctx,
			cachedHeight,
		)
		if err != nil {
			logger.Error("Failed to get all validator sets", zap.Error(err))
			return fmt.Errorf("failed to get all validator sets: %w", err)
		}

		for _, subnetID := range subnetIDs {
			vdrs, ok := allValidatorSets[subnetID]
			if !ok {
				logger.Error("No validators for subnet", zap.Stringer("subnetID", subnetID))
				return fmt.Errorf("no validators for subnet %s", subnetID)
			}
			canonicalSet := network.BuildCanonicalValidators(vdrs)

			if !sharedUtils.CheckStakeWeightExceedsThreshold(
				big.NewInt(0).SetUint64(canonicalSet.ConnectedWeight),
				canonicalSet.ValidatorSet.TotalWeight,
				warp.WarpDefaultQuorumNumerator,
			) {
				logger.Error("Not enough connected stake for subnet",
					zap.Stringer("subnetID", subnetID),
					zap.Uint64("connectedWeight", canonicalSet.ConnectedWeight),
					zap.Uint64("totalWeight", canonicalSet.ValidatorSet.TotalWeight),
				)
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
