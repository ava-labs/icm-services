// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/p2p/acp118"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator/cache"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"go.uber.org/zap"
)

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

// Combination of a subnet ID and a blockchain ID
// used to key the p2pSignatureAggregators map
type subnetChainID struct {
	subnetID ids.ID // signing subnetID of the message being signed
	chainID  ids.ID // sourceChainID of the message being signed
}

type SignatureAggregator struct {
	network peers.AppRequestNetwork
	// protected by subnetsMapLock
	subnetIDsByBlockchainID map[ids.ID]ids.ID
	logger                  logging.Logger
	messageCreator          message.Creator
	currentRequestID        atomic.Uint32
	subnetsMapLock          sync.RWMutex
	metrics                 *metrics.SignatureAggregatorMetrics
	cache                   *cache.Cache

	p2pSigAggLock           sync.RWMutex
	p2pSignatureAggregators map[subnetChainID]*acp118.SignatureAggregator
}

func NewSignatureAggregator(
	network peers.AppRequestNetwork,
	logger logging.Logger,
	messageCreator message.Creator,
	signatureCacheSize uint64,
	metrics *metrics.SignatureAggregatorMetrics,
) (*SignatureAggregator, error) {
	cache, err := cache.NewCache(signatureCacheSize, logger)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create signature cache: %w",
			err,
		)
	}
	sa := SignatureAggregator{
		network:                 network,
		subnetIDsByBlockchainID: map[ids.ID]ids.ID{},
		logger:                  logger,
		metrics:                 metrics,
		currentRequestID:        atomic.Uint32{},
		cache:                   cache,
		messageCreator:          messageCreator,
		p2pSignatureAggregators: make(map[subnetChainID]*acp118.SignatureAggregator),
	}
	sa.currentRequestID.Store(rand.Uint32())

	return &sa, nil
}

func (s *SignatureAggregator) Shutdown() {
	s.network.Shutdown()
}

func (s *SignatureAggregator) connectToQuorumValidators(
	signingSubnet ids.ID,
	quorumPercentage uint64,
) (*peers.ConnectedCanonicalValidators, error) {
	s.network.TrackSubnet(signingSubnet)

	var connectedValidators *peers.ConnectedCanonicalValidators
	var err error
	connectOp := func() error {
		connectedValidators, err = s.network.GetConnectedCanonicalValidators(signingSubnet)
		if err != nil {
			msg := "Failed to connect to canonical validators"
			s.logger.Error(
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
				float64(connectedValidators.TotalValidatorWeight) * 100,
		)
		if !utils.CheckStakeWeightExceedsThreshold(
			big.NewInt(0).SetUint64(connectedValidators.ConnectedWeight),
			connectedValidators.TotalValidatorWeight,
			quorumPercentage,
		) {
			s.logger.Warn(
				"Failed to connect to a threshold of stake",
				zap.Uint64("connectedWeight", connectedValidators.ConnectedWeight),
				zap.Uint64("totalValidatorWeight", connectedValidators.TotalValidatorWeight),
				zap.Uint64("quorumPercentage", quorumPercentage),
			)
			s.metrics.FailuresToConnectToSufficientStake.Inc()
			return errNotEnoughConnectedStake
		}
		return nil
	}
	err = utils.WithRetriesTimeout(s.logger, connectOp, connectToValidatorsTimeout)
	if err != nil {
		return nil, err
	}
	return connectedValidators, nil
}

func (s *SignatureAggregator) CreateSignedMessage(
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	inputSigningSubnet ids.ID,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	s.logger.Debug("Creating signed message", zap.String("warpMessageID", unsignedMessage.ID().String()))
	var signingSubnet ids.ID
	var err error
	// If signingSubnet is not set we default to the subnet of the source blockchain
	sourceSubnet, err := s.getSubnetID(unsignedMessage.SourceChainID)
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
	s.logger.Debug(
		"Creating signed message with signing subnet",
		zap.String("warpMessageID", unsignedMessage.ID().String()),
		zap.Stringer("signingSubnet", signingSubnet),
	)

	aggregator, err := s.getAggregator(signingSubnet, unsignedMessage.SourceChainID)
	if err != nil {
		s.logger.Error(
			"Failed to get aggregator",
			zap.String("warpMessageID", unsignedMessage.ID().String()),
			zap.Error(err),
		)
		return nil, err
	}
	connectedValidators, err := s.connectToQuorumValidators(signingSubnet, quorumPercentage)
	if err != nil {
		s.logger.Error(
			"Failed to connect to quorum of validators",
		)
		return nil, err
	}

	accumulatedSignatureWeight := big.NewInt(0)
	signatureMap := make(map[int][bls.SignatureLen]byte)
	if cachedSignatures, ok := s.cache.Get(unsignedMessage.ID()); ok {
		for i, validator := range connectedValidators.ValidatorSet {
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
			len(connectedValidators.ValidatorSet) - len(signatureMap),
		))
	}

	msg := avalancheWarp.Message{
		UnsignedMessage: *unsignedMessage,
		Signature:       &avalancheWarp.BitSetSignature{},
	}

	s.logger.Info("Unsigned message", zap.String("warpMessageID", unsignedMessage.ID().String()))
	signedMessage := &avalancheWarp.Message{}
	operation := func() error {
		var aggregatedStake *big.Int
		var totalStake *big.Int
		var finished bool
		signedMessage, aggregatedStake, totalStake, finished, err = aggregator.AggregateSignatures(
			context.Background(),
			&msg,
			justification,
			connectedValidators.ValidatorSet,
			quorumPercentage,
			warp.WarpQuorumDenominator,
		)
		if err != nil {
			msg := "Failed to aggregate signatures via p2p client"
			s.logger.Error(
				msg,
				zap.String("warpMessageID", unsignedMessage.ID().String()),
				zap.Error(err),
			)
			return fmt.Errorf("%s: %w", msg, err)
		}
		if !finished {
			msg := "Failed to aggregate sufficient stake"
			s.logger.Error(
				msg,
				zap.String("warpMessageID", unsignedMessage.ID().String()),
				zap.String("aggregatedStake", aggregatedStake.String()),
				zap.String("totalStake", totalStake.String()),
			)
			return fmt.Errorf("%s", msg)
		}
		s.logger.Info(
			"Successfully aggregated signatures",
			zap.String("warpMessageID", unsignedMessage.ID().String()),
			zap.String("aggregatedStake", aggregatedStake.String()),
			zap.String("totalStake", totalStake.String()))
		return nil
	}
	err = utils.WithRetriesTimeout(s.logger, operation, signatureRequestTimeout)
	if err != nil {
		return nil, err
	}
	return signedMessage, nil
}

func (s *SignatureAggregator) getSubnetID(blockchainID ids.ID) (ids.ID, error) {
	s.subnetsMapLock.RLock()
	subnetID, ok := s.subnetIDsByBlockchainID[blockchainID]
	s.subnetsMapLock.RUnlock()
	if ok {
		return subnetID, nil
	}
	s.logger.Info("Signing subnet not found, requesting from PChain", zap.String("blockchainID", blockchainID.String()))
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

func (s *SignatureAggregator) getAggregator(subnetID ids.ID, chainID ids.ID) (*acp118.SignatureAggregator, error) {
	s.p2pSigAggLock.RLock()
	aggregator, ok := s.p2pSignatureAggregators[subnetChainID{subnetID: subnetID, chainID: chainID}]
	s.p2pSigAggLock.RUnlock()
	if ok {
		return aggregator, nil
	}
	s.logger.Info("No P2P signature aggregator found, creating a new one", zap.String("subnetID", subnetID.String()))
	p2pClient, err := peers.NewP2PClient(s.logger, s.network, s.messageCreator, subnetID, chainID)
	if err != nil {
		return nil, err
	}
	aggregator = acp118.NewSignatureAggregator(s.logger, p2pClient)
	s.setAggregator(subnetID, chainID, aggregator)
	return aggregator, nil
}

func (s *SignatureAggregator) setAggregator(subnetID ids.ID, chainID ids.ID, aggregator *acp118.SignatureAggregator) {
	s.p2pSigAggLock.Lock()
	s.p2pSignatureAggregators[subnetChainID{subnetID: subnetID, chainID: chainID}] = aggregator
	s.p2pSigAggLock.Unlock()
}

func (s *SignatureAggregator) aggregateIfSufficientWeight(
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureMap map[int][bls.SignatureLen]byte,
	accumulatedSignatureWeight *big.Int,
	connectedValidators *peers.ConnectedCanonicalValidators,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	// As soon as the signatures exceed the stake weight threshold we try to aggregate and send the transaction.
	if !utils.CheckStakeWeightExceedsThreshold(
		accumulatedSignatureWeight,
		connectedValidators.TotalValidatorWeight,
		quorumPercentage,
	) {
		// Not enough signatures, continue processing messages
		return nil, nil
	}
	aggSig, vdrBitSet, err := s.aggregateSignatures(signatureMap)
	if err != nil {
		msg := "Failed to aggregate signature."
		s.logger.Error(
			msg,
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.String("warpMessageID", unsignedMessage.ID().String()),
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
		s.logger.Error(
			msg,
			zap.String("sourceBlockchainID", unsignedMessage.SourceChainID.String()),
			zap.String("warpMessageID", unsignedMessage.ID().String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}
	return signedMsg, nil
}

// aggregateSignatures constructs a BLS aggregate signature from the collected validator signatures. Also
// returns a bit set representing the validators that are represented in the aggregate signature. The bit
// set is in canonical validator order.
func (s *SignatureAggregator) aggregateSignatures(
	signatureMap map[int][bls.SignatureLen]byte,
) (*bls.Signature, set.Bits, error) {
	// Aggregate the signatures
	signatures := make([]*bls.Signature, 0, len(signatureMap))
	vdrBitSet := set.NewBits()

	for i, sigBytes := range signatureMap {
		sig, err := bls.SignatureFromBytes(sigBytes[:])
		if err != nil {
			msg := "Failed to unmarshal signature"
			s.logger.Error(msg, zap.Error(err))
			return nil, set.Bits{}, fmt.Errorf("%s: %w", msg, err)
		}
		signatures = append(signatures, sig)
		vdrBitSet.Add(i)
	}

	aggSig, err := bls.AggregateSignatures(signatures)
	if err != nil {
		msg := "Failed to aggregate signatures"
		s.logger.Error(msg, zap.Error(err))
		return nil, set.Bits{}, fmt.Errorf("%s: %w", msg, err)
	}
	return aggSig, vdrBitSet, nil
}
