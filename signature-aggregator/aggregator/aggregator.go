// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	networkP2P "github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/units"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type blsSignatureBuf [bls.SignatureLen]byte

const (
	// When selecting which validators to query for signatures, validators are taken in
	// descending weight order until their cumulative weight covers this percentage of the
	// total stake. Querying more stake than the quorum threshold leaves a buffer so that a
	// few non-responding validators don't prevent us from reaching quorum.
	queryStakePercentage = 95

	// Once [queryStakePercentage] of stake is already covered, validators whose individual
	// weight is below this percentage of the total stake are not queried. This avoids
	// querying a long tail of tiny validators for a negligible amount of additional stake.
	minQueryWeightPercentage = 1

	// The minimum balance that an L1 validator must maintain in order to participate
	// in the aggregate signature.
	minimumL1ValidatorBalance = 2048 * units.NanoAvax

	// The amount of time to cache L1 validator balances
	l1ValidatorBalanceTTL = 2 * time.Second
)

var (
	// Errors
	errInvalidQuorumPercentage = errors.New("invalid total quorum percentage")
	errNotEnoughSignatures     = errors.New("failed to collect a threshold of signatures")
	errNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
)

type SignatureAggregator struct {
	network                 *peers.AppRequestNetwork
	messageCreator          message.Creator
	currentRequestID        atomic.Uint32
	metrics                 *metrics.SignatureAggregatorMetrics
	signatureCache          *SignatureCache
	validatorClient         clients.CanonicalValidatorState
	underfundedL1NodeCache  *cache.TTLCache[ids.ID, set.Set[ids.NodeID]]
	signatureRequestTimeout time.Duration

	subnetMapsLock sync.Mutex

	// following block of fields is protected by the subnetMapsLock
	subnetIDsByBlockchainID map[ids.ID]ids.ID
	subnetIDIsL1            map[ids.ID]bool
}

func NewSignatureAggregator(
	network *peers.AppRequestNetwork,
	messageCreator message.Creator,
	signatureCacheSize uint64,
	metrics *metrics.SignatureAggregatorMetrics,
	validatorClient clients.CanonicalValidatorState,
	signatureRequestTimeout time.Duration,
) (*SignatureAggregator, error) {
	signatureCache, err := NewSignatureCache(signatureCacheSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create signature cache: %w", err)
	}
	sa := SignatureAggregator{
		network:                 network,
		subnetIDsByBlockchainID: map[ids.ID]ids.ID{},
		subnetIDIsL1:            map[ids.ID]bool{},
		metrics:                 metrics,
		currentRequestID:        atomic.Uint32{},
		signatureCache:          signatureCache,
		messageCreator:          messageCreator,
		validatorClient:         validatorClient,
		underfundedL1NodeCache:  cache.NewTTLCache[ids.ID, set.Set[ids.NodeID]](l1ValidatorBalanceTTL),
		signatureRequestTimeout: signatureRequestTimeout,
	}
	// invariant: requestIDs for AppRequests must be odd numbered
	sa.currentRequestID.Store(rand.Uint32() | 1)
	return &sa, nil
}

func (s *SignatureAggregator) connectToQuorumValidators(
	ctx context.Context,
	logger logging.Logger,
	signingSubnet ids.ID,
	quorumPercentage uint64,
	pchainHeight uint64,
) (*peers.CanonicalValidators, error) {
	s.network.TrackSubnet(ctx, signingSubnet)

	var vdrs *peers.CanonicalValidators
	var err error
	connectOp := func() error {
		vdrs, err = s.network.GetCanonicalValidators(ctx, signingSubnet, pchainHeight)
		if err != nil {
			msg := "Failed to fetch connected canonical validators"
			logger.Warn(msg, zap.Error(err))
			s.metrics.FailuresToGetValidatorSet.Inc()
			return fmt.Errorf("%s: %w", msg, err)
		}
		s.metrics.ConnectedStakeWeightPercentage.WithLabelValues(
			signingSubnet.String(),
		).Set(
			float64(vdrs.ConnectedWeight) /
				float64(vdrs.ValidatorSet.TotalWeight) * 100,
		)
		if !utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(vdrs.ConnectedWeight),
			vdrs.ValidatorSet.TotalWeight,
			quorumPercentage,
		) {
			// Log details of each connected validator for troubleshooting
			if logger.Enabled(logging.Debug) {
				for _, nodeID := range vdrs.ConnectedNodes.List() {
					vdr, _ := vdrs.GetValidator(nodeID)
					logger.Debug(
						"Connected validator details",
						zap.Stringer("signingSubnet", signingSubnet),
						zap.Stringer("nodeID", nodeID),
						zap.Uint64("weight", vdr.Weight),
					)
				}
			}
			logger.Info(
				"Failed to connect to a threshold of stake",
				zap.Stringer("signingSubnet", signingSubnet),
				zap.Uint64("connectedWeight", vdrs.ConnectedWeight),
				zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
				zap.Uint64("quorumPercentage", quorumPercentage),
				zap.Int("numConnectedPeers", vdrs.ConnectedNodes.Len()),
			)
			s.metrics.FailuresToConnectToSufficientStake.Inc()
			return errNotEnoughConnectedStake
		}
		return nil
	}
	notify := func(err error, duration time.Duration) {
		logger.Debug(
			"connect to validators failed, retrying...",
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}
	err = utils.WithRetriesTimeout(connectOp, notify, utils.ConnectToValidatorsTimeout)
	if err != nil {
		return nil, err
	}
	return vdrs, nil
}

// getUnderfundedL1Nodes fetches the set of L1 nodes that are underfunded
// It uses the underfundedL1NodeCache to avoid fetching the same data multiple times
func (s *SignatureAggregator) getUnderfundedL1Nodes(
	ctx context.Context,
	log logging.Logger,
	signingSubnet ids.ID,
) (set.Set[ids.NodeID], error) {
	fetchUnderfundedL1Nodes := func(subnetID ids.ID) (set.Set[ids.NodeID], error) {
		validators, err := s.validatorClient.GetCurrentValidators(ctx, subnetID)
		if err != nil {
			log.Error("Failed to fetch current L1 validators", zap.Error(err))
			return nil, err
		}

		underfundedL1Nodes := set.NewSet[ids.NodeID](0)
		for _, v := range validators {
			log = log.With(zap.Stringer("nodeID", v.NodeID))
			if v.ClientL1Validator.ValidationID == nil {
				log.Debug("Skipping non-L1 validator")
				continue
			}

			l1Validator := v.ClientL1Validator

			if l1Validator.Balance == nil {
				underfundedL1Nodes.Add(v.NodeID)
				log.Warn("Node has nil balance")
				continue
			}

			if *l1Validator.Balance < minimumL1ValidatorBalance {
				underfundedL1Nodes.Add(v.NodeID)
				log.Debug(
					"Node has insufficient balance",
					zap.Uint64("balance", *l1Validator.Balance),
				)
			}
		}

		return underfundedL1Nodes, nil
	}

	underfundedL1Nodes, err := s.underfundedL1NodeCache.Get(
		signingSubnet,
		fetchUnderfundedL1Nodes,
		false,
	)
	if err != nil {
		log.Error("Failed to get underfunded L1 nodes", zap.Error(err))
		return nil, err
	}

	return underfundedL1Nodes, nil
}

func (s *SignatureAggregator) getExcludedValidators(
	ctx context.Context,
	log logging.Logger,
	signingSubnet ids.ID,
	vdrs *peers.CanonicalValidators,
) (set.Set[int], error) {
	log = log.With(zap.Stringer("signingSubnetID", signingSubnet))

	underfundedL1Nodes, err := s.getUnderfundedL1Nodes(ctx, log, signingSubnet)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch underfunded L1 nodes: %w", err)
	}

	excludedValidators := set.NewSet[int](0)
	// Only exclude a canonical validator if all of its nodes are underfunded L1 validators.
	for i, validator := range vdrs.ValidatorSet.Validators {
		exclude := true
		for _, nodeID := range validator.NodeIDs {
			// Filter out L1 validators that do not have minimumL1ValidatorBalance
			if !underfundedL1Nodes.Contains(nodeID) {
				exclude = false
				break
			}
		}
		if exclude {
			log.Debug(
				"Excluding validator",
				zap.Int("index", i),
				zap.Any("nodeIDs", validator.NodeIDs),
			)
			excludedValidators.Add(i)
		}
	}

	return excludedValidators, nil
}

func validateQuorumPercentages(required, buffer uint64) error {
	if required == 0 || required+buffer > 100 {
		return errInvalidQuorumPercentage
	}
	return nil
}

func (s *SignatureAggregator) selectSigningSubnet(
	ctx context.Context,
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	inputSigningSubnet ids.ID,
) (ids.ID, ids.ID, error) {
	sourceSubnetID, err := s.getSubnetID(ctx, log, unsignedMessage.SourceChainID)
	if err != nil {
		return ids.ID{}, ids.ID{}, err
	}

	var signingSubnetID ids.ID
	if inputSigningSubnet == ids.Empty {
		signingSubnetID = sourceSubnetID
	} else {
		signingSubnetID = inputSigningSubnet
	}
	return signingSubnetID, sourceSubnetID, nil
}

// Gets all of the signatures for the given message that have been cached from the connected validators.
// Excludes previously fetched signatures from any validators now inactive.
// Returns the valid cached signatures to be used, and the total weight of the validators those signatures represent.
func (s *SignatureAggregator) getCachedSignaturesForMessage(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	vdrs *peers.CanonicalValidators,
	excludedValidators set.Set[int],
) (map[int][bls.SignatureLen]byte, *big.Int) {
	signatureMap := make(map[int][bls.SignatureLen]byte)
	accumulatedSignatureWeight := big.NewInt(0)
	if cachedSignatures, ok := s.signatureCache.Get(unsignedMessage.ID()); ok {
		log.Debug("Found cached signatures", zap.Int("signatureCount", len(cachedSignatures)))
		for i, validator := range vdrs.ValidatorSet.Validators {
			cachedSignature, found := cachedSignatures[PublicKeyBytes(validator.PublicKeyBytes)]
			// Do not include explicitly excluded validators in the aggregation
			if found && !excludedValidators.Contains(i) {
				signatureMap[i] = cachedSignature
				accumulatedSignatureWeight.Add(
					accumulatedSignatureWeight,
					new(big.Int).SetUint64(validator.Weight),
				)
			}
		}
	}
	s.metrics.SignatureCacheHits.Add(float64(len(signatureMap)))
	return signatureMap, accumulatedSignatureWeight
}

// queryableValidator is a single canonical validator that still needs to be queried,
// paired with its weight and the connected node IDs to send the request to.
type queryableValidator struct {
	weight  uint64
	nodeIDs []ids.NodeID
}

// queryableValidatorsByWeight returns the validators that still need to be queried
// (those without a cached signature in [signatureMap] and with at least one connected
// node), sorted by descending weight. The canonical validator set is already ordered by
// public key, so a stable sort preserves that order for validators of equal weight.
func queryableValidatorsByWeight(
	vdrs *peers.CanonicalValidators,
	signatureMap map[int][bls.SignatureLen]byte,
) []queryableValidator {
	queryable := make([]queryableValidator, 0, len(vdrs.ValidatorSet.Validators))
	for i, v := range vdrs.ValidatorSet.Validators {
		if _, has := signatureMap[i]; has {
			continue
		}
		var nodeIDs []ids.NodeID
		for _, nodeID := range v.NodeIDs {
			if vdrs.ConnectedNodes.Contains(nodeID) {
				nodeIDs = append(nodeIDs, nodeID)
			}
		}
		// Skip validators we have no connected node to query.
		if len(nodeIDs) == 0 {
			continue
		}
		queryable = append(queryable, queryableValidator{weight: v.Weight, nodeIDs: nodeIDs})
	}
	utils.SortByWeightDescending(queryable, func(v queryableValidator) uint64 {
		return v.weight
	})
	return queryable
}

// nodesToQuery returns the set of connected node IDs to request signatures from, preferring
// higher-weight validators. Validators in [queryable] (already sorted by descending weight)
// are included until their cumulative weight covers [coverageGoal] percent of [totalWeight];
// beyond that point validators are included only while their individual weight is at least
// [minQueryWeightPercentage] of [totalWeight], so the long tail of tiny validators is skipped.
func nodesToQuery(queryable []queryableValidator, totalWeight, coverageGoal uint64) set.Set[ids.NodeID] {
	nodes := set.NewSet[ids.NodeID](len(queryable))
	var cumulative uint64
	for _, v := range queryable {
		covered := weightAtLeastPercent(cumulative, totalWeight, coverageGoal)
		tiny := !weightAtLeastPercent(v.weight, totalWeight, minQueryWeightPercentage)
		if covered && tiny {
			break
		}
		cumulative += v.weight
		for _, nodeID := range v.nodeIDs {
			nodes.Add(nodeID)
		}
	}
	return nodes
}

// weightAtLeastPercent reports whether [weight] is at least [percent] percent of
// [totalWeight], i.e. weight >= totalWeight * percent / 100, computed with big.Int to
// avoid overflowing uint64.
func weightAtLeastPercent(weight, totalWeight, percent uint64) bool {
	lhs := new(big.Int).Mul(new(big.Int).SetUint64(weight), big.NewInt(100))
	rhs := new(big.Int).Mul(new(big.Int).SetUint64(totalWeight), new(big.Int).SetUint64(percent))
	return lhs.Cmp(rhs) >= 0
}

// sendRequest issues a single AppRequest to [queryNodes], registering response and timeout
// tracking with the network only for the nodes actually reached. It returns the channel the
// network handler delivers responses (and timeouts) on, plus the number of responses to
// expect (the number of nodes successfully reached). If no node could be reached, it returns
// a nil channel and a count of 0.
func (s *SignatureAggregator) sendRequest(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	reqBytes []byte,
	sourceSubnet ids.ID,
	queryNodes set.Set[ids.NodeID],
) (chan message.InboundMessage, int, error) {
	requestID := s.currentRequestID.Add(2)
	outMsg, err := s.messageCreator.AppRequest(
		unsignedMessage.SourceChainID,
		requestID,
		utils.DefaultAppRequestTimeout,
		reqBytes,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create app request message: %w", err)
	}

	// Send first, then register the request and timeouts only for the nodes the network
	// actually reached, so the handler's expected-response count agrees with the count
	// returned to the caller.
	sentTo := s.network.Send(outMsg, queryNodes, sourceSubnet, subnets.NoOpAllower)
	s.metrics.AppRequestCount.Inc()

	var failedSendNodes []ids.NodeID
	for nodeID := range queryNodes {
		if !sentTo.Contains(nodeID) {
			failedSendNodes = append(failedSendNodes, nodeID)
			s.metrics.FailuresSendingToNode.Inc()
		}
	}
	if len(failedSendNodes) > 0 {
		log.Info(
			"Failed to make async request to some nodes",
			zap.Uint32("requestID", requestID),
			zap.Int("numSent", sentTo.Len()),
			zap.Int("numFailures", len(failedSendNodes)),
			zap.Stringers("failedNodes", failedSendNodes),
		)
	}

	if sentTo.Len() == 0 {
		return nil, 0, nil
	}

	for nodeID := range sentTo {
		s.network.RegisterAppRequest(ids.RequestID{
			NodeID:    nodeID,
			ChainID:   unsignedMessage.SourceChainID,
			RequestID: requestID,
			Op:        byte(message.AppResponseOp),
		})
	}

	responseChan := s.network.RegisterRequestID(requestID, sentTo)
	if responseChan == nil {
		return nil, 0, fmt.Errorf("failed to register request ID %d", requestID)
	}

	log.Debug(
		"Sent signature request to network",
		zap.Uint32("requestID", requestID),
		zap.Int("expectedResponses", sentTo.Len()),
	)

	return responseChan, sentTo.Len(), nil
}

// finishResponses drains and releases any responses left on [responseChan] that the caller
// did not process (e.g. because quorum was reached or the context expired). The handler
// closes the channel once every expected response or timeout has arrived, so this returns
// once all outstanding requests resolve.
func finishResponses(responseChan <-chan message.InboundMessage) {
	for resp := range responseChan {
		resp.OnFinishedHandling()
	}
}

// collectSignatures queries validators for signatures, prioritizing higher-weight validators.
// It selects the highest-weight validators that together cover most of the stake (see
// [nodesToQuery]), sends them all a single signature request, and aggregates the responses as
// they arrive. It returns as soon as the accumulated weight meets the required quorum plus
// buffer, and otherwise falls back to the required quorum once all responses are in. The
// overall collection is bounded by [signatureRequestTimeout].
func (s *SignatureAggregator) collectSignatures(
	ctx context.Context,
	logger logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	reqBytes []byte,
	sourceSubnet, signingSubnet ids.ID,
	vdrs *peers.CanonicalValidators,
	signatureMap map[int][bls.SignatureLen]byte,
	excludedValidators set.Set[int],
	accumulatedSignatureWeight *big.Int,
	requiredQuorumPercentage, quorumPercentageBuffer uint64,
) (*avalancheWarp.Message, error) {
	log := logger.With(
		zap.Stringer("sourceBlockchainID", unsignedMessage.SourceChainID),
		zap.Stringer("signingSubnetID", signingSubnet),
	)

	ctx, cancel := context.WithTimeout(ctx, s.signatureRequestTimeout)
	defer cancel()

	totalWeight := vdrs.ValidatorSet.TotalWeight
	targetThreshold := requiredQuorumPercentage + quorumPercentageBuffer

	// Select the highest-weight validators to query. Always cover enough stake to reach the
	// target threshold if everyone responds, while skipping the long tail of tiny validators.
	queryable := queryableValidatorsByWeight(vdrs, signatureMap)
	coverageGoal := max(queryStakePercentage, targetThreshold)
	queryNodes := nodesToQuery(queryable, totalWeight, coverageGoal)

	log.Debug(
		"Aggregator collecting signatures from weight-prioritized validators.",
		zap.Int("validatorSetSize", len(vdrs.ValidatorSet.Validators)),
		zap.Int("signatureMapSize", len(signatureMap)),
		zap.Int("queryableValidators", len(queryable)),
		zap.Int("queryNodes", queryNodes.Len()),
	)

	if queryNodes.Len() > 0 {
		signedMsg, err := s.requestSignatures(
			ctx,
			log,
			unsignedMessage,
			reqBytes,
			sourceSubnet,
			queryNodes,
			vdrs,
			signatureMap,
			excludedValidators,
			accumulatedSignatureWeight,
			targetThreshold,
		)
		if err != nil {
			return nil, err
		}
		if signedMsg != nil {
			log.Info(
				"Created signed message.",
				zap.Uint64("signatureWeight", accumulatedSignatureWeight.Uint64()),
				zap.Uint64("totalValidatorWeight", totalWeight),
			)
			return signedMsg, nil
		}
	}

	// We did not reach the (required + buffer) threshold, so try aggregating with just the
	// required quorum.
	signedMsg, err := s.aggregateIfSufficientWeight(
		log,
		unsignedMessage,
		signatureMap,
		accumulatedSignatureWeight,
		vdrs.ValidatorSet.Validators,
		totalWeight,
		requiredQuorumPercentage,
	)
	if err != nil {
		return nil, err
	}
	if signedMsg != nil {
		log.Info(
			"Created signed message.",
			zap.Uint64("signatureWeight", accumulatedSignatureWeight.Uint64()),
			zap.Uint64("totalValidatorWeight", totalWeight),
		)
		return signedMsg, nil
	}

	// The caller logs this failure (with the same context) at error level, so avoid
	// logging it twice here.
	return nil, errNotEnoughSignatures
}

// requestSignatures sends a single signature request to [queryNodes] and processes the
// responses as they arrive, returning a signed message as soon as [quorumPercentage] of stake
// is accumulated. It returns (nil, nil) if all responses are processed (or the context
// expires) without reaching the threshold.
func (s *SignatureAggregator) requestSignatures(
	ctx context.Context,
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	reqBytes []byte,
	sourceSubnet ids.ID,
	queryNodes set.Set[ids.NodeID],
	vdrs *peers.CanonicalValidators,
	signatureMap map[int][bls.SignatureLen]byte,
	excludedValidators set.Set[int],
	accumulatedSignatureWeight *big.Int,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	responseChan, expectedResponses, err := s.sendRequest(
		log,
		unsignedMessage,
		reqBytes,
		sourceSubnet,
		queryNodes,
	)
	if err != nil {
		return nil, err
	}
	if responseChan == nil {
		return nil, nil
	}
	// Release any responses we don't process (because quorum was reached or the context
	// expired) so their handlers are always finished. The drain runs in the background so
	// returning early isn't blocked waiting for outstanding requests to time out.
	defer func() { go finishResponses(responseChan) }()

	for i := 0; i < expectedResponses; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case response, ok := <-responseChan:
			if !ok {
				return nil, nil
			}
			signedMsg, err := s.handleResponse(
				log,
				response,
				vdrs,
				unsignedMessage,
				signatureMap,
				excludedValidators,
				accumulatedSignatureWeight,
				quorumPercentage,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to handle response: %w", err)
			}
			if signedMsg != nil {
				return signedMsg, nil
			}
		}
	}
	return nil, nil
}

func (s *SignatureAggregator) CreateSignedMessage(
	ctx context.Context,
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	inputSigningSubnet ids.ID,
	requiredQuorumPercentage uint64,
	quorumPercentageBuffer uint64,
	pchainHeight uint64,
) (*avalancheWarp.Message, error) {
	log = log.With(
		zap.Uint64("requiredQuorumPercentage", requiredQuorumPercentage),
		zap.Uint64("quorumPercentageBuffer", quorumPercentageBuffer),
		zap.Uint64("pchainHeight", pchainHeight),
		zap.Stringer("sourceBlockchainID", unsignedMessage.SourceChainID),
	)
	log.Info("Creating signed message")
	if err := validateQuorumPercentages(requiredQuorumPercentage, quorumPercentageBuffer); err != nil {
		log.Error("Invalid quorum percentages")
		return nil, err
	}

	log.Debug("Creating signed message")
	// Select signing subnet
	signingSubnet, sourceSubnet, err := s.selectSigningSubnet(ctx, log, unsignedMessage, inputSigningSubnet)
	if err != nil {
		return nil, err
	}

	log = log.With(zap.Stringer("signingSubnet", signingSubnet))
	log.Debug("Creating signed message with signing subnet")

	vdrs, err := s.connectToQuorumValidators(
		ctx,
		log,
		signingSubnet,
		requiredQuorumPercentage,
		pchainHeight,
	)
	if err != nil {
		log.Error("Failed to fetch quorum of connected canonical validators", zap.Error(err))
		return nil, err
	}

	isL1 := false
	if signingSubnet != constants.PrimaryNetworkID {
		isL1, err = s.isSubnetL1(ctx, signingSubnet)
		if err != nil {
			log.Error("Failed to check if signing subnet is L1", zap.Error(err))
			return nil, err
		}
	}

	// Tracks all collected signatures.
	// For L1s, we must take care to *not* include inactive validators in the signature map.
	// Inactive validator's stake weight still contributes to the total weight, but the verifying
	// node will not be able to verify the aggregate signature if it includes an inactive validator.
	var excludedValidators set.Set[int]

	// Fetch L1 validators and find the node IDs with Balance < minimumL1ValidatorBalance
	// Find the corresponding canonical validator set index for each of these, and add to the exclusion list
	// if ALL of the node IDs for a validator have Balance < minimumL1ValidatorBalance
	if isL1 {
		log.Debug("Checking L1 validators for zero balance nodes")
		excludedValidators, err = s.getExcludedValidators(ctx, log, signingSubnet, vdrs)
		if err != nil {
			log.Error("Failed to get excluded validators", zap.Error(err))
			return nil, fmt.Errorf("failed to get excluded validators: %w", err)
		}
	}

	// Populate signature map from cache
	signatureMap, accumulatedSignatureWeight := s.getCachedSignaturesForMessage(
		log,
		unsignedMessage,
		vdrs,
		excludedValidators,
	)

	// Only return early if we have enough signatures to meet the quorum percentage
	// plus the buffer percentage.
	if signedMsg, err := s.aggregateIfSufficientWeight(
		log,
		unsignedMessage,
		signatureMap,
		accumulatedSignatureWeight,
		vdrs.ValidatorSet.Validators,
		vdrs.ValidatorSet.TotalWeight,
		requiredQuorumPercentage+quorumPercentageBuffer,
	); err != nil {
		return nil, err
	} else if signedMsg != nil {
		return signedMsg, nil
	}
	if len(signatureMap) > 0 {
		s.metrics.SignatureCacheMisses.Add(float64(
			len(vdrs.ValidatorSet.Validators) - len(signatureMap),
		))
	}

	reqBytes, err := s.marshalRequest(unsignedMessage, justification, sourceSubnet)
	if err != nil {
		msg := "Failed to marshal request bytes"
		log.Error(msg, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Collect signatures from validators in weight-prioritized batches.
	signedMsg, err := s.collectSignatures(
		ctx,
		log,
		unsignedMessage,
		reqBytes,
		sourceSubnet,
		signingSubnet,
		vdrs,
		signatureMap,
		excludedValidators,
		accumulatedSignatureWeight,
		requiredQuorumPercentage,
		quorumPercentageBuffer,
	)
	if err != nil {
		log.Error(
			"Failed to collect signatures",
			zap.Uint64("accumulatedWeight", accumulatedSignatureWeight.Uint64()),
			zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
			zap.Error(err),
		)
		return nil, err
	}
	return signedMsg, nil
}

func (s *SignatureAggregator) getSubnetID(
	ctx context.Context,
	log logging.Logger,
	blockchainID ids.ID,
) (ids.ID, error) {
	s.subnetMapsLock.Lock()
	defer s.subnetMapsLock.Unlock()

	subnetID, ok := s.subnetIDsByBlockchainID[blockchainID]
	if ok {
		return subnetID, nil
	}
	log.Info("Signing subnet not found, requesting from PChain", zap.Stringer("blockchainID", blockchainID))
	getSubnetIDCtx, cancel := context.WithTimeout(ctx, utils.DefaultRPCTimeout)
	defer cancel()
	subnetID, err := s.network.GetSubnetID(getSubnetIDCtx, blockchainID)
	if err != nil {
		return ids.ID{}, fmt.Errorf("source blockchain not found for chain ID %s", blockchainID)
	}
	s.subnetIDsByBlockchainID[blockchainID] = subnetID
	return subnetID, nil
}

// Looks up whether a subnet is an L1 and caches the result in the map for the lifetime of the application.
// since this value can change only once in the lifetime of the subnet.
func (s *SignatureAggregator) isSubnetL1(ctx context.Context, subnetID ids.ID) (bool, error) {
	s.subnetMapsLock.Lock()
	defer s.subnetMapsLock.Unlock()
	isL1, ok := s.subnetIDIsL1[subnetID]
	if !ok {
		subnet, err := s.validatorClient.GetSubnet(ctx, subnetID)
		if err != nil {
			return false, fmt.Errorf("failed to get subnet: %w", err)
		}
		isL1 = subnet.ConversionID != ids.Empty
		s.subnetIDIsL1[subnetID] = isL1
	}
	return isL1, nil
}

// handleResponse processes a single response from a validator that the aggregator is
// awaiting. It caches any valid signature, updates [signatureMap] / [accumulatedSignatureWeight]
// if the corresponding validator is not in [excludedValidators], and returns a non-nil signed
// Warp message if [accumulatedSignatureWeight] now meets [quorumPercentage] of the validator
// set's total weight. Returns an error only if a non-recoverable error occurs, otherwise
// returns a nil error to continue processing further responses.
func (s *SignatureAggregator) handleResponse(
	log logging.Logger,
	response message.InboundMessage,
	connectedValidators *peers.CanonicalValidators,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	excludedValidators set.Set[int],
	accumulatedSignatureWeight *big.Int,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	// Regardless of the response's relevance, call its finished handler once this function returns.
	defer response.OnFinishedHandling()

	nodeID := response.NodeID

	// If we receive an AppRequestFailed, then the request timed out at the network layer.
	// We treat this as a relevant response, since we are no longer expecting a real response
	// from that node for this request.
	if response.Op == message.AppErrorOp {
		log.Debug("Request timed out", zap.Stringer("nodeID", nodeID))
		s.metrics.ValidatorTimeouts.Inc()
		return nil, nil
	}

	validator, vdrIndex := connectedValidators.GetValidator(nodeID)
	signature, valid := s.isValidSignatureResponse(log, unsignedMessage, response, validator.PublicKey)
	// Cache any valid signature, but only include in the aggregation if the validator is not
	// explicitly excluded, so that we can reuse the cached signature on future requests if
	// the validator is no longer excluded.
	if valid {
		log.Debug(
			"Got valid signature response",
			zap.Stringer("nodeID", nodeID),
			zap.Uint64("stakeWeight", validator.Weight),
			zap.Stringer("sourceBlockchainID", unsignedMessage.SourceChainID),
		)
		s.signatureCache.Add(
			unsignedMessage.ID(),
			PublicKeyBytes(validator.PublicKeyBytes),
			SignatureBytes(signature),
		)
		// A validator may be reached through more than one node ID, so guard against
		// counting its weight twice if multiple of its nodes return a valid signature.
		if _, alreadyCounted := signatureMap[vdrIndex]; !alreadyCounted && !excludedValidators.Contains(vdrIndex) {
			signatureMap[vdrIndex] = signature
			accumulatedSignatureWeight.Add(accumulatedSignatureWeight, new(big.Int).SetUint64(validator.Weight))
		}
	} else {
		log.Debug(
			"Got invalid signature response",
			zap.Stringer("nodeID", nodeID),
			zap.Uint64("stakeWeight", validator.Weight),
			zap.Stringer("sourceBlockchainID", unsignedMessage.SourceChainID),
		)
		s.metrics.InvalidSignatureResponses.Inc()
		return nil, nil
	}

	return s.aggregateIfSufficientWeight(
		log,
		unsignedMessage,
		signatureMap,
		accumulatedSignatureWeight,
		connectedValidators.ValidatorSet.Validators,
		connectedValidators.ValidatorSet.TotalWeight,
		quorumPercentage,
	)
}

func (s *SignatureAggregator) aggregateIfSufficientWeight(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	accumulatedSignatureWeight *big.Int,
	canonicalValidators []*validators.Warp,
	totalWeight uint64,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	// As soon as the signatures exceed the stake weight threshold we try to aggregate and send the transaction.
	if !utils.CheckStakeWeightExceedsThreshold(
		accumulatedSignatureWeight,
		totalWeight,
		quorumPercentage,
	) {
		// Not enough signatures, continue processing messages
		return nil, nil
	}

	// Prune the signature map to the smallest subset whose combined weight still
	// meets quorum. Fewer signers means smaller calldata on destination chains.
	prunedSigMap := pruneSignatureMapToQuorum(
		signatureMap,
		canonicalValidators,
		totalWeight,
		quorumPercentage,
	)

	aggSig, vdrBitSet, err := s.aggregateSignatures(log, prunedSigMap)
	if err != nil {
		msg := "Failed to aggregate signature."
		log.Error(msg, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	signedMsg, err := avalancheWarp.NewMessage(
		unsignedMessage,
		&avalancheWarp.BitSetSignature{
			Signers:   vdrBitSet.Bytes(),
			Signature: *(*[bls.SignatureLen]byte)(bls.SignatureToBytes(aggSig)),
		},
	)
	if err != nil {
		msg := "Failed to create new signed message"
		log.Error(msg, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msg, err)
	}
	return signedMsg, nil
}

// pruneSignatureMapToQuorum returns the subset of signatureMap with the
// heaviest validators (greedy by weight desc) whose combined weight is the
// minimum that still meets quorumPercentage. canonicalValidators must be the
// canonical set the signatureMap indices refer to.
func pruneSignatureMapToQuorum(
	signatureMap map[int][bls.SignatureLen]byte,
	canonicalValidators []*validators.Warp,
	totalWeight uint64,
	quorumPercentage uint64,
) map[int][bls.SignatureLen]byte {
	type signerEntry struct {
		idx    int
		weight uint64
	}
	// Build in canonical index order so SliceStable yields a deterministic
	// ordering for validators of equal weight.
	signers := make([]signerEntry, 0, len(signatureMap))
	for i, v := range canonicalValidators {
		if _, ok := signatureMap[i]; !ok {
			continue
		}
		signers = append(signers, signerEntry{idx: i, weight: v.Weight})
	}
	utils.SortByWeightDescending(signers, func(s signerEntry) uint64 {
		return s.weight
	})

	requiredWeight := utils.RequiredSignatureWeight(totalWeight, quorumPercentage)
	pruned := make(map[int][bls.SignatureLen]byte, len(signers))
	accumulated := new(big.Int)
	weightBuf := new(big.Int)
	for _, sgn := range signers {
		pruned[sgn.idx] = signatureMap[sgn.idx]
		accumulated.Add(accumulated, weightBuf.SetUint64(sgn.weight))
		if accumulated.Cmp(requiredWeight) >= 0 {
			break
		}
	}
	return pruned
}

// isValidSignatureResponse tries to generate a signature from the peer.AsyncResponse, then verifies
// the signature against the node's public key. If we are unable to generate the signature or verify
// correctly, false will be returned to indicate no valid signature was found in response.
func (s *SignatureAggregator) isValidSignatureResponse(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	response message.InboundMessage,
	pubKey *bls.PublicKey,
) (blsSignatureBuf, bool) {
	log = log.With(zap.Stringer("nodeID", response.NodeID))
	// If the handler returned an error response, count the response and continue
	if response.Op == message.AppErrorOp {
		log.Debug("Relayer async response failed")
		return blsSignatureBuf{}, false
	}

	appResponse, ok := response.Message.(*p2p.AppResponse)
	if !ok {
		log.Debug("Relayer async response was not an AppResponse")
		return blsSignatureBuf{}, false
	}

	signature, err := s.unmarshalResponse(appResponse.GetAppBytes())
	if err != nil {
		log.Error("Error unmarshaling signature response", zap.Error(err))
	}

	// If the node returned an empty signature, then it has not yet seen the warp message. Retry later.
	emptySignature := blsSignatureBuf{}
	if bytes.Equal(signature[:], emptySignature[:]) {
		log.Debug("Response contained an empty signature")
		return blsSignatureBuf{}, false
	}

	if len(signature) != bls.SignatureLen {
		log.Debug(
			"Response signature has incorrect length",
			zap.Int("actual", len(signature)),
			zap.Int("expected", bls.SignatureLen),
		)
		return blsSignatureBuf{}, false
	}

	sig, err := bls.SignatureFromBytes(signature[:])
	if err != nil {
		log.Debug("Failed to create signature from response")
		return blsSignatureBuf{}, false
	}

	if !bls.Verify(pubKey, sig, unsignedMessage.Bytes()) {
		log.Debug(
			"Failed verification for signature",
			zap.String("pubKey", hex.EncodeToString(bls.PublicKeyToUncompressedBytes(pubKey))),
		)
		return blsSignatureBuf{}, false
	}

	return signature, true
}

// aggregateSignatures constructs a BLS aggregate signature from the collected validator signatures. Also
// returns a bit set representing the validators that are represented in the aggregate signature. The bit
// set is in canonical validator order.
func (s *SignatureAggregator) aggregateSignatures(
	log logging.Logger,
	signatureMap map[int][bls.SignatureLen]byte,
) (*bls.Signature, set.Bits, error) {
	// Aggregate the signatures
	signatures := make([]*bls.Signature, 0, len(signatureMap))
	vdrBitSet := set.NewBits()

	for i, sigBytes := range signatureMap {
		sig, err := bls.SignatureFromBytes(sigBytes[:])
		if err != nil {
			msg := "Failed to unmarshal signature"
			log.Error(msg, zap.Error(err))
			return nil, set.Bits{}, fmt.Errorf("%s: %w", msg, err)
		}
		signatures = append(signatures, sig)
		vdrBitSet.Add(i)
	}

	aggSig, err := bls.AggregateSignatures(signatures)
	if err != nil {
		msg := "Failed to aggregate signatures"
		log.Error(msg, zap.Error(err))
		return nil, set.Bits{}, fmt.Errorf("%s: %w", msg, err)
	}
	return aggSig, vdrBitSet, nil
}

func (s *SignatureAggregator) marshalRequest(
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	sourceSubnet ids.ID,
) ([]byte, error) {
	messageBytes, err := proto.Marshal(
		&sdk.SignatureRequest{
			Message:       unsignedMessage.Bytes(),
			Justification: justification,
		},
	)
	if err != nil {
		return nil, err
	}
	return networkP2P.PrefixMessage(
		networkP2P.ProtocolPrefix(networkP2P.SignatureRequestHandlerID),
		messageBytes,
	), nil
}

func (s *SignatureAggregator) unmarshalResponse(responseBytes []byte) (blsSignatureBuf, error) {
	// empty responses are valid and indicate the node has not seen the message
	if len(responseBytes) == 0 {
		return blsSignatureBuf{}, nil
	}
	var sigResponse sdk.SignatureResponse
	err := proto.Unmarshal(responseBytes, &sigResponse)
	if err != nil {
		return blsSignatureBuf{}, err
	}
	return blsSignatureBuf(sigResponse.Signature), nil
}
