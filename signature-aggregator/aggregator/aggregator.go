// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"bytes"
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
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator/cache"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/icm-services/utils"
	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type blsSignatureBuf [bls.SignatureLen]byte

const (
	// Maximum amount of time to spend waiting (in addition to network round trip time per attempt)
	// during relayer signature query routine
	signatureRequestTimeout = 5 * time.Second
	// Maximum amount of time to spend waiting for a connection to a quorum of validators for
	// a given subnetID
	connectToValidatorsTimeout = 5 * time.Second
)

var (
	// Errors
	errNotEnoughSignatures     = errors.New("failed to collect a threshold of signatures")
	errNotEnoughConnectedStake = errors.New("failed to connect to a threshold of stake")
)

type SignatureAggregator struct {
	network peers.AppRequestNetwork
	// protected by subnetsMapLock
	subnetIDsByBlockchainID map[ids.ID]ids.ID
	messageCreator          message.Creator
	currentRequestID        atomic.Uint32
	subnetsMapLock          sync.RWMutex
	metrics                 *metrics.SignatureAggregatorMetrics
	cache                   *cache.Cache
}

func NewSignatureAggregator(
	network peers.AppRequestNetwork,
	messageCreator message.Creator,
	signatureCacheSize uint64,
	metrics *metrics.SignatureAggregatorMetrics,
) (*SignatureAggregator, error) {
	cache, err := cache.NewCache(signatureCacheSize)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create signature cache: %w",
			err,
		)
	}
	sa := SignatureAggregator{
		network:                 network,
		subnetIDsByBlockchainID: map[ids.ID]ids.ID{},
		metrics:                 metrics,
		currentRequestID:        atomic.Uint32{},
		cache:                   cache,
		messageCreator:          messageCreator,
	}
	sa.currentRequestID.Store(rand.Uint32())
	return &sa, nil
}

func (s *SignatureAggregator) Shutdown() {
	s.network.Shutdown()
}

func (s *SignatureAggregator) connectToQuorumValidators(
	log logging.Logger,
	signingSubnet ids.ID,
	quorumPercentage uint64,
) (*peers.ConnectedCanonicalValidators, error) {
	s.network.TrackSubnet(signingSubnet)

	var connectedValidators *peers.ConnectedCanonicalValidators
	var err error
	connectOp := func() error {
		connectedValidators, err = s.network.GetConnectedCanonicalValidators(signingSubnet)
		if err != nil {
			msg := "Failed to fetch connected canonical validators"
			log.Error(
				msg,
				zap.Error(err),
			)
			s.metrics.FailuresToGetValidatorSet.Inc()
			return fmt.Errorf("%s: %w", msg, err)
		}
		s.metrics.ConnectedStakeWeightPercentage.WithLabelValues(
			signingSubnet.String(),
		).Set(
			float64(connectedValidators.ConnectedWeight) /
				float64(connectedValidators.ValidatorSet.TotalWeight) * 100,
		)
		if !utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(connectedValidators.ConnectedWeight),
			connectedValidators.ValidatorSet.TotalWeight,
			quorumPercentage,
		) {
			log.Warn(
				"Failed to connect to a threshold of stake",
				zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
				zap.Uint64("totalValidatorWeight", connectedValidators.ValidatorSet.TotalWeight),
				zap.Uint64("quorumPercentage", quorumPercentage),
			)
			s.metrics.FailuresToConnectToSufficientStake.Inc()
			return errNotEnoughConnectedStake
		}
		return nil
	}
	err = utils.WithRetriesTimeout(log, connectOp, connectToValidatorsTimeout)
	if err != nil {
		return nil, err
	}
	return connectedValidators, nil
}

func (s *SignatureAggregator) CreateSignedMessage(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	inputSigningSubnet ids.ID,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	log = log.With(zap.String("warpMessageID", unsignedMessage.ID().String()))
	log.Debug("Creating signed message")
	var signingSubnet ids.ID
	var err error
	// If signingSubnet is not set we default to the subnet of the source blockchain
	sourceSubnet, err := s.getSubnetID(log, unsignedMessage.SourceChainID)
	if err != nil {
		return nil, fmt.Errorf(
			"source message subnet not found for chainID %s",
			unsignedMessage.SourceChainID,
		)
	}
	if inputSigningSubnet == ids.Empty {
		signingSubnet = sourceSubnet
	} else {
		signingSubnet = inputSigningSubnet
	}
	log.Debug(
		"Creating signed message with signing subnet",
		zap.Stringer("signingSubnet", signingSubnet),
	)

	connectedValidators, err := s.connectToQuorumValidators(log, signingSubnet, quorumPercentage)
	if err != nil {
		log.Error(
			"Failed to fetch quorum of connected canonical validators",
			zap.Stringer("signingSubnet", signingSubnet),
			zap.Error(err),
		)
		return nil, err
	}

	accumulatedSignatureWeight := big.NewInt(0)
	signatureMap := make(map[int][bls.SignatureLen]byte)
	if cachedSignatures, ok := s.cache.Get(unsignedMessage.ID()); ok {
		log.Debug("Found cached signatures", zap.Int("signatureCount", len(cachedSignatures)))
		for i, validator := range connectedValidators.ValidatorSet.Validators {
			cachedSignature, found := cachedSignatures[cache.PublicKeyBytes(validator.PublicKeyBytes)]
			if found {
				signatureMap[i] = cachedSignature
				accumulatedSignatureWeight.Add(
					accumulatedSignatureWeight,
					new(big.Int).SetUint64(validator.Weight),
				)
			}
		}
		s.metrics.SignatureCacheHits.Add(float64(len(signatureMap)))
	}
	if signedMsg, err := s.aggregateIfSufficientWeight(
		log,
		unsignedMessage,
		signatureMap,
		accumulatedSignatureWeight,
		connectedValidators,
		quorumPercentage,
	); err != nil {
		return nil, err
	} else if signedMsg != nil {
		return signedMsg, nil
	}
	if len(signatureMap) > 0 {
		s.metrics.SignatureCacheMisses.Add(float64(
			len(connectedValidators.ValidatorSet.Validators) - len(signatureMap),
		))
	}

	reqBytes, err := s.marshalRequest(unsignedMessage, justification, sourceSubnet)
	if err != nil {
		msg := "Failed to marshal request bytes"
		log.Error(
			msg,
			zap.Error(err),
		)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Construct the AppRequest
	requestID := s.currentRequestID.Add(1)
	outMsg, err := s.messageCreator.AppRequest(
		unsignedMessage.SourceChainID,
		requestID,
		peers.DefaultAppRequestTimeout,
		reqBytes,
	)
	if err != nil {
		msg := "Failed to create app request message"
		log.Error(
			msg,
			zap.Error(err),
		)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	var signedMsg *avalancheWarp.Message
	// Query the validators with retries. On each retry, query one node per unique BLS pubkey
	operation := func() error {
		responsesExpected := len(connectedValidators.ValidatorSet.Validators) - len(signatureMap)
		log.Debug(
			"Aggregator collecting signatures from peers.",
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.String("signingSubnetID", signingSubnet.String()),
			zap.Int("validatorSetSize", len(connectedValidators.ValidatorSet.Validators)),
			zap.Int("signatureMapSize", len(signatureMap)),
			zap.Int("responsesExpected", responsesExpected),
		)

		vdrSet := set.NewSet[ids.NodeID](len(connectedValidators.ValidatorSet.Validators))
		for i, vdr := range connectedValidators.ValidatorSet.Validators {
			// If we already have the signature for this validator, do not query any of the composite nodes again
			if _, ok := signatureMap[i]; ok {
				continue
			}

			// TODO: Track failures and iterate through the validator's node list on subsequent query attempts
			nodeID := vdr.NodeIDs[0]
			vdrSet.Add(nodeID)
			log.Debug(
				"Added node ID to query.",
				zap.String("nodeID", nodeID.String()),
				zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			)

			// Register a timeout response for each queried node
			reqID := ids.RequestID{
				NodeID:    nodeID,
				ChainID:   unsignedMessage.SourceChainID,
				RequestID: requestID,
				Op:        byte(message.AppResponseOp),
			}
			s.network.RegisterAppRequest(reqID)
		}
		responseChan := s.network.RegisterRequestID(requestID, vdrSet.Len())

		sentTo := s.network.Send(outMsg, vdrSet, sourceSubnet, subnets.NoOpAllower)
		s.metrics.AppRequestCount.Inc()
		log.Debug(
			"Sent signature request to network",
			zap.Any("sentTo", sentTo),
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.String("sourceSubnetID", sourceSubnet.String()),
			zap.String("signingSubnetID", signingSubnet.String()),
		)

		failedSendNodes := make([]ids.NodeID, 0, responsesExpected)
		for nodeID := range vdrSet {
			if !sentTo.Contains(nodeID) {
				log.Debug(
					"Failed to make async request to node",
					zap.String("nodeID", nodeID.String()),
					zap.Error(err),
				)
				responsesExpected--
				failedSendNodes = append(failedSendNodes, nodeID)
				s.metrics.FailuresSendingToNode.Inc()
			}
		}
		if len(failedSendNodes) > 0 {
			log.Warn(
				"Failed to make async request to some nodes",
				zap.Int("numSent", responsesExpected),
				zap.Int("numFailures", len(failedSendNodes)),
				zap.Stringers("failedNodes", failedSendNodes),
			)
		}

		responseCount := 0
		if responsesExpected > 0 {
			for response := range responseChan {
				log.Debug(
					"Processing response from node",
					zap.String("nodeID", response.NodeID().String()),
					zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
				)
				var relevant bool
				signedMsg, relevant, err = s.handleResponse(
					log,
					response,
					sentTo,
					requestID,
					connectedValidators,
					unsignedMessage,
					signatureMap,
					accumulatedSignatureWeight,
					quorumPercentage,
				)
				if err != nil {
					// don't increase node failures metric here, because we did
					// it in handleResponse
					return backoff.Permanent(fmt.Errorf(
						"failed to handle response: %w",
						err,
					))
				}
				if relevant {
					responseCount++
				}
				// If we have sufficient signatures, return here.
				if signedMsg != nil {
					log.Info(
						"Created signed message.",
						zap.Uint64("signatureWeight", accumulatedSignatureWeight.Uint64()),
						zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
					)
					return nil
				}
				// Break once we've had successful or unsuccessful responses from each requested node
				if responseCount == responsesExpected {
					break
				}
			}
		}
		return errNotEnoughSignatures
	}

	err = utils.WithRetriesTimeout(log, operation, signatureRequestTimeout)
	if err != nil {
		log.Warn(
			"Failed to collect a threshold of signatures",
			zap.Uint64("accumulatedWeight", accumulatedSignatureWeight.Uint64()),
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
		)
		return nil, errNotEnoughSignatures
	}
	return signedMsg, nil
}

func (s *SignatureAggregator) getSubnetID(log logging.Logger, blockchainID ids.ID) (ids.ID, error) {
	s.subnetsMapLock.RLock()
	subnetID, ok := s.subnetIDsByBlockchainID[blockchainID]
	s.subnetsMapLock.RUnlock()
	if ok {
		return subnetID, nil
	}
	log.Info("Signing subnet not found, requesting from PChain", zap.String("blockchainID", blockchainID.String()))
	subnetID, err := s.network.GetSubnetID(blockchainID)
	if err != nil {
		return ids.ID{}, fmt.Errorf("source blockchain not found for chain ID %s", blockchainID)
	}
	s.setSubnetID(blockchainID, subnetID)
	return subnetID, nil
}

func (s *SignatureAggregator) setSubnetID(blockchainID ids.ID, subnetID ids.ID) {
	s.subnetsMapLock.Lock()
	s.subnetIDsByBlockchainID[blockchainID] = subnetID
	s.subnetsMapLock.Unlock()
}

// Attempts to create a signed Warp message from the accumulated responses.
// Returns a non-nil Warp message if [accumulatedSignatureWeight] exceeds the signature verification threshold.
// Returns false in the second return parameter if the app response is not relevant to the current signature
// aggregation request. Returns an error only if a non-recoverable error occurs, otherwise returns a nil error
// to continue processing responses.
func (s *SignatureAggregator) handleResponse(
	log logging.Logger,
	response message.InboundMessage,
	sentTo set.Set[ids.NodeID],
	requestID uint32,
	connectedValidators *peers.ConnectedCanonicalValidators,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	accumulatedSignatureWeight *big.Int,
	quorumPercentage uint64,
) (*avalancheWarp.Message, bool, error) {
	// Regardless of the response's relevance, call it's finished handler once this function returns
	defer response.OnFinishedHandling()

	// Check if this is an expected response.
	m := response.Message()
	rcvReqID, ok := message.GetRequestID(m)
	if !ok {
		// This should never occur, since inbound message validity is already checked by the inbound handler
		log.Error("Could not get requestID from message")
		return nil, false, nil
	}
	nodeID := response.NodeID()
	if !sentTo.Contains(nodeID) || rcvReqID != requestID {
		log.Debug("Skipping irrelevant app response")
		return nil, false, nil
	}

	// If we receive an AppRequestFailed, then the request timed out.
	// This is still a relevant response, since we are no longer expecting a response from that node.
	if response.Op() == message.AppErrorOp {
		log.Debug("Request timed out")
		s.metrics.ValidatorTimeouts.Inc()
		return nil, true, nil
	}

	validator, vdrIndex := connectedValidators.GetValidator(nodeID)
	signature, valid := s.isValidSignatureResponse(log, unsignedMessage, response, validator.PublicKey)
	if valid {
		log.Debug(
			"Got valid signature response",
			zap.String("nodeID", nodeID.String()),
			zap.Uint64("stakeWeight", validator.Weight),
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
		)
		signatureMap[vdrIndex] = signature
		s.cache.Add(
			unsignedMessage.ID(),
			cache.PublicKeyBytes(validator.PublicKeyBytes),
			cache.SignatureBytes(signature),
		)
		accumulatedSignatureWeight.Add(accumulatedSignatureWeight, new(big.Int).SetUint64(validator.Weight))
	} else {
		log.Debug(
			"Got invalid signature response",
			zap.String("nodeID", nodeID.String()),
			zap.Uint64("stakeWeight", validator.Weight),
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
		)
		s.metrics.InvalidSignatureResponses.Inc()
		return nil, true, nil
	}

	if signedMsg, err := s.aggregateIfSufficientWeight(
		log,
		unsignedMessage,
		signatureMap,
		accumulatedSignatureWeight,
		connectedValidators,
		quorumPercentage,
	); err != nil {
		return nil, true, err
	} else if signedMsg != nil {
		return signedMsg, true, nil
	}

	// Not enough signatures, continue processing messages
	return nil, true, nil
}

func (s *SignatureAggregator) aggregateIfSufficientWeight(
	log logging.Logger,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	accumulatedSignatureWeight *big.Int,
	connectedValidators *peers.ConnectedCanonicalValidators,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	// As soon as the signatures exceed the stake weight threshold we try to aggregate and send the transaction.
	if !utils.CheckStakeWeightExceedsThreshold(
		accumulatedSignatureWeight,
		connectedValidators.ValidatorSet.TotalWeight,
		quorumPercentage,
	) {
		// Not enough signatures, continue processing messages
		return nil, nil
	}
	aggSig, vdrBitSet, err := s.aggregateSignatures(log, signatureMap)
	if err != nil {
		msg := "Failed to aggregate signature."
		log.Error(
			msg,
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.Error(err),
		)
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
		log.Error(
			msg,
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.Error(err),
		)
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
	// If the handler returned an error response, count the response and continue
	if response.Op() == message.AppErrorOp {
		log.Debug(
			"Relayer async response failed",
			zap.String("nodeID", response.NodeID().String()),
		)
		return blsSignatureBuf{}, false
	}

	appResponse, ok := response.Message().(*p2p.AppResponse)
	if !ok {
		log.Debug(
			"Relayer async response was not an AppResponse",
			zap.String("nodeID", response.NodeID().String()),
		)
		return blsSignatureBuf{}, false
	}

	signature, err := s.unmarshalResponse(appResponse.GetAppBytes())
	if err != nil {
		log.Error(
			"Error unmarshaling signature response",
			zap.Error(err),
		)
	}

	// If the node returned an empty signature, then it has not yet seen the warp message. Retry later.
	emptySignature := blsSignatureBuf{}
	if bytes.Equal(signature[:], emptySignature[:]) {
		log.Debug(
			"Response contained an empty signature",
			zap.String("nodeID", response.NodeID().String()),
		)
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
		log.Debug(
			"Failed to create signature from response",
		)
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
