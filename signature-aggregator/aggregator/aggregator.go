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
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	networkP2P "github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
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
	// Maximum amount of time to spend waiting (in addition to network round trip time per attempt)
	// during relayer signature query routine
	signatureRequestTimeout = 2 * utils.DefaultAppRequestTimeout
	// Maximum amount of time to spend waiting for a connection to a quorum of validators for
	// a given subnetID
	connectToValidatorsTimeout = 5 * time.Second

	// The minimum balance that an L1 validator must maintain in order to participate
	// in the aggregate signature.
	minimumL1ValidatorBalance = 2048 * units.NanoAvax

	// The amount of time to cache L1 validator balances
	l1ValidatorBalanceTTL = 2 * time.Second

	// signatureBatchInterval bounds the time the aggregator waits between
	// dispatching successive batches of signature requests. If the current
	// batch's responses arrive (or its requests time out) sooner, the next
	// batch is sent immediately. If the interval elapses with the deficit
	// unmet, a new batch sized to cover the remaining weight deficit is
	// dispatched without re-querying validators we have already asked.
	signatureBatchInterval = 1 * time.Second
)

var (
	// Errors
	errInvalidQuorumPercentage = errors.New("invalid total quorum percentage")
	errNotEnoughSignatures     = errors.New("failed to collect a threshold of signatures")
	errNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
)

type SignatureAggregator struct {
	network                *peers.AppRequestNetwork
	messageCreator         message.Creator
	currentRequestID       atomic.Uint32
	metrics                *metrics.SignatureAggregatorMetrics
	signatureCache         *SignatureCache
	validatorClient        clients.CanonicalValidatorState
	underfundedL1NodeCache *cache.TTLCache[ids.ID, set.Set[ids.NodeID]]

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
	err = utils.WithRetriesTimeout(connectOp, notify, connectToValidatorsTimeout)
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

// sortedCandidateValidators returns the indices of canonical validators that
// can still be queried for additional signatures, ordered by weight
// descending. A validator is considered queryable when:
//   - we do not already have its signature (from the cache or a prior batch),
//   - it is not in excludedValidators (e.g. an L1 validator with insufficient
//     balance, which cannot contribute to the aggregate signature), and
//   - at least one of its composite node IDs is currently connected.
//
// The returned slice is the priority order in which validators will be
// queried; the highest-weighted validators come first so that we can reach
// quorum with the fewest possible requests.
func sortedCandidateValidators(
	vdrs *peers.CanonicalValidators,
	signatureMap map[int][bls.SignatureLen]byte,
	excludedValidators set.Set[int],
) []int {
	candidates := make([]int, 0, len(vdrs.ValidatorSet.Validators))
	for i, vdr := range vdrs.ValidatorSet.Validators {
		if _, ok := signatureMap[i]; ok {
			continue
		}
		if excludedValidators.Contains(i) {
			continue
		}
		connected := false
		for _, nodeID := range vdr.NodeIDs {
			if vdrs.ConnectedNodes.Contains(nodeID) {
				connected = true
				break
			}
		}
		if !connected {
			continue
		}
		candidates = append(candidates, i)
	}
	sort.SliceStable(candidates, func(a, b int) bool {
		return vdrs.ValidatorSet.Validators[candidates[a]].Weight >
			vdrs.ValidatorSet.Validators[candidates[b]].Weight
	})
	return candidates
}

// pickNextBatch selects the smallest prefix of candidates (starting at
// offset) whose combined weight, added to accumulatedWeight, would cross
// the target threshold. Returns the picked canonical validator indices and
// the new offset to use for the next call. If the remaining candidates do
// not cover the deficit, all remaining candidates are returned.
func pickNextBatch(
	vdrs *peers.CanonicalValidators,
	candidates []int,
	offset int,
	accumulatedWeight *big.Int,
	targetQuorumPercentage uint64,
) (batch []int, newOffset int) {
	if offset >= len(candidates) {
		return nil, offset
	}
	projected := new(big.Int).Set(accumulatedWeight)
	for i := offset; i < len(candidates); i++ {
		idx := candidates[i]
		weight := vdrs.ValidatorSet.Validators[idx].Weight
		projected.Add(projected, new(big.Int).SetUint64(weight))
		batch = append(batch, idx)
		if utils.CheckStakeWeightExceedsThreshold(
			projected,
			vdrs.ValidatorSet.TotalWeight,
			targetQuorumPercentage,
		) {
			return batch, i + 1
		}
	}
	return batch, len(candidates)
}

// signatureCollection holds the shared per-call state used while collecting
// signatures for a single CreateSignedMessage invocation. All mutable
// fields must be accessed under mu.
type signatureCollection struct {
	// Immutable configuration set at construction.
	vdrs                   *peers.CanonicalValidators
	excludedValidators     set.Set[int]
	targetQuorumPercentage uint64
	unsignedMessage        *avalancheWarp.UnsignedMessage

	// Shared mutable state protected by mu.
	mu                         sync.Mutex
	signatureMap               map[int][bls.SignatureLen]byte
	accumulatedSignatureWeight *big.Int

	// quorumReached is closed (exactly once, via quorumOnce) when the
	// accumulated weight first reaches targetQuorumPercentage. It lets the
	// coordinator short-circuit additional batches as soon as quorum is
	// achieved by any in-flight response processor.
	quorumOnce    sync.Once
	quorumReached chan struct{}
}

func (c *signatureCollection) signalQuorum() {
	c.quorumOnce.Do(func() { close(c.quorumReached) })
}

// snapshotWeight returns a copy of accumulatedSignatureWeight taken under mu.
func (c *signatureCollection) snapshotWeight() *big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return new(big.Int).Set(c.accumulatedSignatureWeight)
}

// snapshot returns copies of signatureMap and accumulatedSignatureWeight
// taken atomically under mu. Callers can safely use the returned values
// while in-flight response handlers continue to mutate the originals.
func (c *signatureCollection) snapshot() (map[int][bls.SignatureLen]byte, *big.Int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	snapMap := make(map[int][bls.SignatureLen]byte, len(c.signatureMap))
	for k, v := range c.signatureMap {
		snapMap[k] = v
	}
	return snapMap, new(big.Int).Set(c.accumulatedSignatureWeight)
}

// processSignatureResponse validates a single inbound response and, if it
// carries a valid signature for a non-excluded validator we don't already
// have, integrates it into the shared collection. It also signals
// quorumReached on the collection the first time the target threshold is
// crossed. It is safe to call concurrently from multiple goroutines.
func (s *SignatureAggregator) processSignatureResponse(
	log logging.Logger,
	c *signatureCollection,
	response message.InboundMessage,
	sentTo set.Set[ids.NodeID],
	requestID uint32,
) {
	defer response.OnFinishedHandling()

	rcvReqID, ok := message.GetRequestID(response.Message)
	if !ok {
		log.Error("Could not get requestID from message")
		return
	}
	nodeID := response.NodeID
	if !sentTo.Contains(nodeID) || rcvReqID != requestID {
		log.Debug("Skipping irrelevant app response")
		return
	}
	if response.Op == message.AppErrorOp {
		log.Debug("Request timed out", zap.Stringer("nodeID", nodeID))
		s.metrics.ValidatorTimeouts.Inc()
		return
	}

	validator, vdrIndex := c.vdrs.GetValidator(nodeID)
	signature, valid := s.isValidSignatureResponse(log, c.unsignedMessage, response, validator.PublicKey)
	if !valid {
		log.Debug(
			"Got invalid signature response",
			zap.Stringer("nodeID", nodeID),
			zap.Uint64("stakeWeight", validator.Weight),
		)
		s.metrics.InvalidSignatureResponses.Inc()
		return
	}
	log.Debug(
		"Got valid signature response",
		zap.Stringer("nodeID", nodeID),
		zap.Uint64("stakeWeight", validator.Weight),
	)
	// Cache the signature regardless of whether the validator is included
	// in this aggregation so that future requests for this message can
	// reuse it (e.g. if the validator was previously excluded).
	s.signatureCache.Add(
		c.unsignedMessage.ID(),
		PublicKeyBytes(validator.PublicKeyBytes),
		SignatureBytes(signature),
	)

	c.mu.Lock()
	defer c.mu.Unlock()
	if _, already := c.signatureMap[vdrIndex]; already {
		return
	}
	if c.excludedValidators.Contains(vdrIndex) {
		return
	}
	c.signatureMap[vdrIndex] = signature
	c.accumulatedSignatureWeight.Add(
		c.accumulatedSignatureWeight,
		new(big.Int).SetUint64(validator.Weight),
	)
	if utils.CheckStakeWeightExceedsThreshold(
		c.accumulatedSignatureWeight,
		c.vdrs.ValidatorSet.TotalWeight,
		c.targetQuorumPercentage,
	) {
		c.signalQuorum()
	}
}

// collectSignatures dispatches signature requests in successive batches,
// each sized to cover the remaining weight deficit, with validators chosen
// in descending weight order. After each batch the coordinator waits up to
// signatureBatchInterval for responses (or until the batch fully resolves)
// before dispatching the next batch that covers what is still missing,
// without ever re-querying a validator. The overall process is bounded by
// signatureRequestTimeout.
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
	targetQuorumPercentage := requiredQuorumPercentage + quorumPercentageBuffer

	deadlineCtx, cancel := context.WithTimeout(ctx, signatureRequestTimeout)
	defer cancel()

	c := &signatureCollection{
		vdrs:                       vdrs,
		excludedValidators:         excludedValidators,
		targetQuorumPercentage:     targetQuorumPercentage,
		unsignedMessage:            unsignedMessage,
		signatureMap:               signatureMap,
		accumulatedSignatureWeight: accumulatedSignatureWeight,
		quorumReached:              make(chan struct{}),
	}

	candidates := sortedCandidateValidators(vdrs, signatureMap, excludedValidators)
	logger.Debug(
		"Built sorted candidate validator list",
		zap.Int("candidateCount", len(candidates)),
		zap.Int("totalValidators", len(vdrs.ValidatorSet.Validators)),
	)

	var (
		wg     sync.WaitGroup
		cursor int
	)

	// dispatchNextBatch runs a single iteration of the coordinator: it picks
	// the next batch sized to the remaining weight deficit, dispatches the
	// AppRequest, spawns a goroutine to drain its response channel, and then
	// waits for the earliest of (target quorum reached, batch fully resolved,
	// global deadline, per-batch interval). It returns true when the
	// coordinator should stop dispatching further batches.
	dispatchNextBatch := func() (stop bool) {
		select {
		case <-c.quorumReached:
			return true
		case <-deadlineCtx.Done():
			return true
		default:
		}

		batchIndices, newCursor := pickNextBatch(
			vdrs, candidates, cursor, c.snapshotWeight(), targetQuorumPercentage,
		)
		cursor = newCursor
		if len(batchIndices) == 0 {
			// No more validators to query; let any in-flight batches finish.
			return true
		}

		vdrSet := set.NewSet[ids.NodeID](len(batchIndices))
		for _, idx := range batchIndices {
			for _, nodeID := range vdrs.ValidatorSet.Validators[idx].NodeIDs {
				if vdrs.ConnectedNodes.Contains(nodeID) {
					vdrSet.Add(nodeID)
				}
			}
		}
		if vdrSet.Len() == 0 {
			return false
		}

		requestID := s.currentRequestID.Add(2)
		log := logger.With(
			zap.Int("requestID", int(requestID)),
			zap.Stringer("sourceBlockchainID", unsignedMessage.SourceChainID),
			zap.Stringer("signingSubnetID", signingSubnet),
		)

		outMsg, err := s.messageCreator.AppRequest(
			unsignedMessage.SourceChainID,
			requestID,
			utils.DefaultAppRequestTimeout,
			reqBytes,
		)
		if err != nil {
			log.Warn("Failed to create app request message", zap.Error(err))
			return false
		}

		for nodeID := range vdrSet {
			s.network.RegisterAppRequest(ids.RequestID{
				NodeID:    nodeID,
				ChainID:   unsignedMessage.SourceChainID,
				RequestID: requestID,
				Op:        byte(message.AppResponseOp),
			})
		}
		respCh := s.network.RegisterRequestID(requestID, vdrSet)
		if respCh == nil {
			log.Error("Failed to register request ID")
			return false
		}

		sentTo := s.network.Send(outMsg, vdrSet, sourceSubnet, subnets.NoOpAllower)
		s.metrics.AppRequestCount.Inc()
		log.Debug(
			"Sent signature request batch",
			zap.Int("batchValidatorCount", len(batchIndices)),
			zap.Int("batchNodeCount", vdrSet.Len()),
			zap.Int("sentToCount", sentTo.Len()),
		)
		for nodeID := range vdrSet {
			if !sentTo.Contains(nodeID) {
				s.metrics.FailuresSendingToNode.Inc()
				log.Debug("Failed to send to node", zap.Stringer("nodeID", nodeID))
			}
		}

		batchDone := make(chan struct{})
		wg.Add(1)
		go func(
			respCh chan message.InboundMessage,
			sentTo set.Set[ids.NodeID],
			requestID uint32,
			log logging.Logger,
		) {
			defer wg.Done()
			defer close(batchDone)
			for resp := range respCh {
				s.processSignatureResponse(log, c, resp, sentTo, requestID)
			}
		}(respCh, sentTo, requestID, log)

		select {
		case <-c.quorumReached:
		case <-deadlineCtx.Done():
		case <-batchDone:
		case <-time.After(signatureBatchInterval):
		}
		return false
	}

	for !dispatchNextBatch() {
	}

	// Give any in-flight batches a chance to deliver responses (bounded by
	// the global deadline) before we snapshot for aggregation.
	allBatchesDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(allBatchesDone)
	}()
	select {
	case <-c.quorumReached:
	case <-allBatchesDone:
	case <-deadlineCtx.Done():
	}

	// Snapshot under lock so that any goroutines that have not yet exited
	// (their AppRequests have not yet timed out) do not race with the
	// aggregation step below.
	snapshotMap, snapshotWeight := c.snapshot()

	// Try to build the signed message at the optimistic target first
	// (required + buffer). This matches the previous "first try target"
	// behavior of collectSignaturesWithRetries.
	if signedMsg, err := s.aggregateIfSufficientWeight(
		logger, unsignedMessage, snapshotMap, snapshotWeight,
		vdrs.ValidatorSet.TotalWeight, targetQuorumPercentage,
	); err != nil {
		return nil, err
	} else if signedMsg != nil {
		logger.Info(
			"Created signed message.",
			zap.Uint64("signatureWeight", snapshotWeight.Uint64()),
			zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
		)
		return signedMsg, nil
	}

	// Fall back to the minimum required quorum if the optimistic target was
	// not met. The verifying node only requires the required percentage; the
	// buffer is best-effort overshoot.
	if signedMsg, err := s.aggregateIfSufficientWeight(
		logger, unsignedMessage, snapshotMap, snapshotWeight,
		vdrs.ValidatorSet.TotalWeight, requiredQuorumPercentage,
	); err != nil {
		return nil, err
	} else if signedMsg != nil {
		logger.Info(
			"Created signed message at minimum required quorum.",
			zap.Uint64("signatureWeight", snapshotWeight.Uint64()),
			zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
		)
		return signedMsg, nil
	}

	logger.Warn(
		"Failed to collect a threshold of signatures",
		zap.Uint64("accumulatedWeight", snapshotWeight.Uint64()),
		zap.Uint64("totalValidatorWeight", vdrs.ValidatorSet.TotalWeight),
	)
	return nil, errNotEnoughSignatures
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

	// Collect signatures by querying validators in successive batches,
	// each sized to the remaining weight deficit and ordered by stake
	// weight (highest first).
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

func (s *SignatureAggregator) aggregateIfSufficientWeight(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	accumulatedSignatureWeight *big.Int,
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
	aggSig, vdrBitSet, err := s.aggregateSignatures(log, signatureMap)
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
