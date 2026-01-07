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
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/sampler"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
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
)

var (
	ErrNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
	errTrackingTooManySubnets  = fmt.Errorf("cannot track more than %d subnets", maxNumSubnets)
)

type AppRequestNetwork struct {
	network network.Network
	handler *RelayerExternalHandler
	logger  logging.Logger

	metrics *AppRequestNetworkMetrics

	// The set of subnetIDs to track. Shared with the underlying Network object, so access
	// must be protected by the trackedSubnetsLock
	trackedSubnets set.Set[ids.ID]
	// invariant: members of lruSubnets should always be exactly the same as trackedSubnets
	// and the size of lruSubnets should be less than or equal to maxNumSubnets
	lruSubnets         *linked.Hashmap[ids.ID, interface{}]
	trackedSubnetsLock *sync.RWMutex

	validatorManager *ValidatorManager
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
	metrics := NewAppRequestNetworkMetrics(relayerRegistry)

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

	pClient := clients.NewCanonicalValidatorClient(cfg.GetPChainAPI())

	vdrs, err := pClient.GetCurrentValidators(ctx, constants.PrimaryNetworkID)
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

	localTrackedSubnets := set.NewSet[ids.ID](maxNumSubnets)

	for _, subnetID := range trackedSubnets.List() {
		localTrackedSubnets.Add(subnetID)
	}

	validatorManager := NewValidatorManager(cfg, logger, metrics, int(validatorSetsCacheSize), manager)

	arNetwork := &AppRequestNetwork{
		network:            testNetwork,
		handler:            handler,
		logger:             logger,
		metrics:            metrics,
		trackedSubnets:     localTrackedSubnets,
		trackedSubnetsLock: trackedSubnetsLock,
		lruSubnets:         lruSubnets,
		validatorManager:   validatorManager,
	}

	go arNetwork.startUpdateTrackedValidators(ctx)

	return arNetwork, nil
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
		n.validatorManager.UpdateTrackedValidatorSet(ctx, subnetID)
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
	n.validatorManager.StartCacheValidatorSets(ctx)
}

func (n *AppRequestNetwork) updateTrackedValidatorSets(ctx context.Context) {
	allValidators, err := n.validatorManager.GetLatestValidatorSets(ctx)
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
				zap.Uint64("pchainHeight", n.validatorManager.GetLatestSyncedPChainHeight()),
			)
			continue
		}
		// If we fail to get the validator sets for this subnet, log and continue to the next subnet
		err := n.validatorManager.updatedTrackedValidators(subnetID, vdrs)
		if err != nil {
			n.logger.Error("Failed to update tracked validators",
				zap.Stringer("subnetID", subnetID),
				zap.Error(err),
			)
		}
	}
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

// GetCanonicalValidators returns the validator information in canonical ordering for the given subnet
// at the specified P-Chain height, as well as the total weight of the validators that this network is connected to
// The caller determines the appropriate P-Chain height (ProposedHeight for current, specific height for epoched)
func (n *AppRequestNetwork) GetCanonicalValidators(
	ctx context.Context,
	subnetID ids.ID,
	pchainHeight uint64,
) (*CanonicalValidators, error) {
	allValidators, err := n.validatorManager.GetAllValidatorSets(ctx, pchainHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get all validators at P-Chain height %d: %w", pchainHeight, err)
	}

	validatorSet, ok := allValidators[subnetID]
	if !ok {
		return nil, fmt.Errorf("no validators for subnet %s at P-Chain height %d", subnetID, pchainHeight)
	}

	return n.buildCanonicalValidators(validatorSet), nil
}

// buildCanonicalValidators builds the CanonicalValidators struct from a validator set
func (n *AppRequestNetwork) buildCanonicalValidators(
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
	return n.validatorManager.GetSubnetID(ctx, blockchainID)
}

// GetNetworkHealthFunc returns a health check function for the network
func (n *AppRequestNetwork) GetNetworkHealthFunc(subnetIDs []ids.ID) func(context.Context) error {
	return func(ctx context.Context) error {
		cachedHeight := n.validatorManager.GetLatestSyncedPChainHeight()
		if cachedHeight == 0 {
			// This should only happen at startup when the cache is not yet initialized.
			n.logger.Info("No cached P-Chain height, skipping network health check")
			return nil
		}

		allValidatorSets, err := n.validatorManager.GetAllValidatorSets(
			ctx,
			cachedHeight,
		)
		if err != nil {
			n.logger.Error("Failed to get all validator sets", zap.Error(err))
			return fmt.Errorf("failed to get all validator sets: %w", err)
		}

		for _, subnetID := range subnetIDs {
			vdrs, ok := allValidatorSets[subnetID]
			if !ok {
				n.logger.Error("No validators for subnet", zap.Stringer("subnetID", subnetID))
				return fmt.Errorf("no validators for subnet %s", subnetID)
			}
			canonicalSet := n.buildCanonicalValidators(vdrs)

			if !utils.CheckStakeWeightExceedsThreshold(
				big.NewInt(0).SetUint64(canonicalSet.ConnectedWeight),
				canonicalSet.ValidatorSet.TotalWeight,
				warp.WarpDefaultQuorumNumerator,
			) {
				n.logger.Error("Not enough connected stake for subnet",
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

// Non-receiver util functions

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
