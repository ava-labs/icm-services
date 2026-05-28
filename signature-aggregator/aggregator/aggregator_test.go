package aggregator

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"
	"slices"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	avagocommon "github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
	avago_mocks "github.com/ava-labs/icm-services/peers/avago_mocks"
	client_mocks "github.com/ava-labs/icm-services/peers/clients/mocks"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
)

const defaultSignatureRequestTimeout = 20 * time.Second

func instantiateDefaultAggregator(t *testing.T) (
	*SignatureAggregator,
	*peers.AppRequestNetwork,
	*peers.RelayerExternalHandler,
	*avago_mocks.MockNetwork,
	*client_mocks.MockCanonicalValidatorState,
) {
	return instantiateAggregator(t, defaultSignatureRequestTimeout)
}

func instantiateAggregator(
	t *testing.T,
	signatureRequestTimeout time.Duration,
) (
	*SignatureAggregator,
	*peers.AppRequestNetwork,
	*peers.RelayerExternalHandler, // handler for test access
	*avago_mocks.MockNetwork,
	*client_mocks.MockCanonicalValidatorState,
) {
	mockController := gomock.NewController(t)
	mockNetwork := avago_mocks.NewMockNetwork(mockController)
	mockValidatorClient := client_mocks.NewMockCanonicalValidatorState(mockController)

	// Create a new registry for each test to avoid duplicate registration errors
	registry := prometheus.NewRegistry()

	// Create fresh metrics for each test
	testSigAggMetrics := metrics.NewSignatureAggregatorMetrics(registry)
	testMessageCreator, err := message.NewCreator(
		registry,
		constants.DefaultNetworkCompressionType,
		constants.DefaultNetworkMaximumInboundTimeout,
	)
	require.NoError(t, err)

	// Create handler for AppRequestNetwork
	peerMetrics := peers.NewAppRequestNetworkMetrics(registry)
	handler, err := peers.NewRelayerExternalHandler(
		logging.NoLog{},
		peerMetrics,
		registry,
	)
	require.NoError(t, err)

	// Create a real AppRequestNetwork with mocked dependencies
	manager := validators.NewManager()
	appRequestNetwork := peers.NewAppRequestNetworkForTesting(
		mockNetwork,
		handler,
		logging.NoLog{},
		peerMetrics,
		mockValidatorClient,
		manager,
	)

	aggregator, err := NewSignatureAggregator(
		appRequestNetwork,
		testMessageCreator,
		1024,
		testSigAggMetrics,
		mockValidatorClient,
		signatureRequestTimeout,
	)
	require.NoError(t, err)

	// Return the AppRequestNetwork, handler (for injecting responses), and mocks so tests can set expectations
	return aggregator, appRequestNetwork, handler, mockNetwork, mockValidatorClient
}

// Generate the validator values.
type validatorInfo struct {
	nodeID            ids.NodeID
	blsSigner         *localsigner.LocalSigner
	blsPublicKey      *bls.PublicKey
	blsPublicKeyBytes []byte
	weight            uint64
}

func (v validatorInfo) Compare(o validatorInfo) int {
	return bytes.Compare(v.blsPublicKeyBytes, o.blsPublicKeyBytes)
}

func makeConnectedValidators(validatorCount int) (*peers.CanonicalValidators, []*localsigner.LocalSigner) {
	weights := make([]uint64, validatorCount)
	for i := range weights {
		weights[i] = 1
	}
	return makeConnectedValidatorsWithWeights(weights)
}

func TestCreateSignedMessageFailsInvalidQuorumPercentage(t *testing.T) {
	testCases := []struct {
		name                     string
		requiredQuorumPercentage uint64
		quorumPercentageBuffer   uint64
	}{
		{
			name:                     "Zero required quorum percentage",
			requiredQuorumPercentage: 0,
			quorumPercentageBuffer:   5,
		},
		{
			name:                     "Quorum percentage above 100",
			requiredQuorumPercentage: 96,
			quorumPercentageBuffer:   5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
			signedMsg, err := aggregator.CreateSignedMessage(
				t.Context(),
				logging.NoLog{},
				&warp.UnsignedMessage{},
				nil,
				ids.Empty,
				tc.requiredQuorumPercentage,
				tc.quorumPercentageBuffer,
				pchainapi.ProposedHeight, // Use ProposedHeight for current validators
			)
			require.Nil(t, signedMsg)
			require.ErrorIs(t, err, errInvalidQuorumPercentage)
		})
	}
}

func TestCreateSignedMessageFailsWithNoValidators(t *testing.T) {
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateDefaultAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), ids.Empty).Return(ids.Empty, nil).AnyTimes()
	// TrackSubnet is on AppRequestNetwork - no mock needed
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), ids.Empty).Return(
		validators.WarpSet{
			Validators:  []*validators.Warp{},
			TotalWeight: 0,
		},
		nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{
			ids.Empty: validators.WarpSet{
				Validators:  []*validators.Warp{},
				TotalWeight: 0,
			},
		},
		nil,
	).AnyTimes()
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return([]peer.Info{}).AnyTimes()
	_, err = aggregator.CreateSignedMessage(
		t.Context(), logging.NoLog{}, msg, nil, ids.Empty, 80, 0, pchainapi.ProposedHeight)
	require.ErrorContains(t, err, "no signatures")
}

func TestCreateSignedMessageFailsWithoutSufficientConnectedStake(t *testing.T) {
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateDefaultAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), ids.Empty).Return(ids.Empty, nil)
	// TrackSubnet is on AppRequestNetwork - no mock needed
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{
			ids.Empty: validators.WarpSet{
				Validators:  []*validators.Warp{},
				TotalWeight: 1,
			},
		},
		nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), ids.Empty).Return(
		validators.WarpSet{
			Validators:  []*validators.Warp{},
			TotalWeight: 1,
		},
		nil,
	).AnyTimes()
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return([]peer.Info{}).AnyTimes()
	_, err = aggregator.CreateSignedMessage(
		t.Context(), logging.NoLog{}, msg, nil, ids.Empty, 80, 0, pchainapi.ProposedHeight)
	require.ErrorContains(
		t,
		err,
		"failed to connect to a threshold of stake",
	)
}

func makeAppRequests(
	chainID ids.ID,
	requestID uint32,
	connectedValidators *peers.CanonicalValidators,
) []ids.RequestID {
	var appRequests []ids.RequestID
	for _, validator := range connectedValidators.ValidatorSet.Validators {
		for _, nodeID := range validator.NodeIDs {
			appRequests = append(
				appRequests,
				ids.RequestID{
					NodeID:    nodeID,
					ChainID:   chainID,
					RequestID: requestID,
					Op: byte(
						message.AppResponseOp,
					),
				},
			)
		}
	}
	return appRequests
}

func TestCreateSignedMessageRetriesAndFailsWithoutP2PResponses(t *testing.T) {
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateAggregator(t, 100*time.Millisecond)

	var (
		connectedValidators, _ = makeConnectedValidators(2)
		requestID              = aggregator.currentRequestID.Load() + 2
	)

	chainID := ids.GenerateTestID()

	msg, err := warp.NewUnsignedMessage(0, chainID, []byte{})
	require.NoError(t, err)

	subnetID := ids.GenerateTestID()
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(
		subnetID,
		nil,
	).AnyTimes()

	// TrackSubnet is on AppRequestNetwork - no mock needed
	mockValidatorClient.EXPECT().GetProposedValidators(
		gomock.Any(), subnetID,
	).Return(
		connectedValidators.ValidatorSet,
		nil,
	).AnyTimes()

	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{
			subnetID: connectedValidators.ValidatorSet,
		},
		nil,
	).AnyTimes()

	// Mock PeerInfo to return connected peers
	var peerInfos []peer.Info
	for nodeID := range connectedValidators.ConnectedNodes {
		peerInfos = append(peerInfos, peer.Info{ID: nodeID})
	}
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()

	appRequests := makeAppRequests(chainID, requestID, connectedValidators)
	var nodeIDs set.Set[ids.NodeID]
	for _, appRequest := range appRequests {
		nodeIDs.Add(appRequest.NodeID)
	}

	mockNetwork.EXPECT().Send(
		gomock.Any(),
		gomock.Any(),
		subnetID,
		subnets.NoOpAllower,
	).AnyTimes()

	mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{},
		nil,
	).Times(1)

	_, err = aggregator.CreateSignedMessage(
		t.Context(), logging.NoLog{}, msg, nil, subnetID, 80, 0, pchainapi.ProposedHeight)
	require.ErrorIs(
		t,
		err,
		errNotEnoughSignatures,
	)
}

func TestCreateSignedMessageSucceeds(t *testing.T) {
	// The test sets up valid signature responses from 4 of 5 equally weighted validators.
	testCases := []struct {
		name                     string
		requiredQuorumPercentage uint64
		quorumPercentageBuffer   uint64
	}{
		{
			name:                     "Succeeds with buffer",
			requiredQuorumPercentage: 67,
			quorumPercentageBuffer:   5,
		},
		{
			name:                     "Succeeds without buffer",
			requiredQuorumPercentage: 80,
			quorumPercentageBuffer:   5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var msg *warp.UnsignedMessage // to be signed
			chainID := ids.GenerateTestID()
			networkID := constants.UnitTestID
			msg, err := warp.NewUnsignedMessage(
				networkID,
				chainID,
				utils.RandomBytes(1234),
			)
			require.NoError(t, err)

			// the signers:
			connectedValidators, validatorSigners := makeConnectedValidators(5)

			// prime the aggregator:

			aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateDefaultAggregator(t)

			subnetID := ids.GenerateTestID()
			mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(
				subnetID,
				nil,
			).AnyTimes()

			// TrackSubnet is on AppRequestNetwork - no mock needed
			mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), subnetID).Return(
				connectedValidators.ValidatorSet,
				nil,
			).AnyTimes()

			mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
				map[ids.ID]validators.WarpSet{
					subnetID: connectedValidators.ValidatorSet,
				},
				nil,
			).AnyTimes()

			// Mock PeerInfo to return connected peers
			var peerInfos []peer.Info
			for nodeID := range connectedValidators.ConnectedNodes {
				peerInfos = append(peerInfos, peer.Info{ID: nodeID})
			}
			mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()

			mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
				platformvm.GetSubnetClientResponse{},
				nil,
			).Times(1)

			// prime the signers' responses:

			requestID := aggregator.currentRequestID.Load() + 2

			appRequests := makeAppRequests(chainID, requestID, connectedValidators)

			var nodeIDs set.Set[ids.NodeID]
			for _, appRequest := range appRequests {
				nodeIDs.Add(appRequest.NodeID)
			}

			// Set up mock to inject responses when Send is called
			mockNetwork.EXPECT().Send(
				gomock.Any(),
				gomock.Any(), // common.SendConfig
				subnetID,
				subnets.NoOpAllower,
			).Times(1).DoAndReturn(
				func(
					outboundMsg *message.OutboundMessage,
					config interface{},
					subnetID ids.ID,
					allower interface{},
				) set.Set[ids.NodeID] {
					// Inject responses in a goroutine after Send is called
					// This simulates the network receiving responses from validators
					go func() {
						// Small delay to ensure the aggregator has registered and is waiting
						time.Sleep(10 * time.Millisecond)

						// Send responses through the handler
						for i, appRequest := range appRequests {
							validatorSigner := validatorSigners[connectedValidators.NodeValidatorIndexMap[appRequest.NodeID]]

							// Simulate 1 of 5 validators responding with an invalid signature
							var signatureBytes []byte
							if i == len(appRequests)-1 {
								signatureBytes = make([]byte, 0)
							} else {
								signature, signErr := validatorSigner.Sign(msg.Bytes())
								if signErr != nil {
									t.Logf("Failed to sign: %v", signErr)
									continue
								}
								signatureBytes = bls.SignatureToBytes(signature)
							}

							responseBytes, marshalErr := proto.Marshal(
								&sdk.SignatureResponse{
									Signature: signatureBytes,
								},
							)
							if marshalErr != nil {
								t.Logf("Failed to marshal: %v", marshalErr)
								continue
							}

							// Create an inbound app response message and send it through the handler
							inboundMsg := message.InboundAppResponse(
								chainID,
								requestID,
								responseBytes,
								appRequest.NodeID,
							)
							// Call the handler directly to inject the response
							handler.HandleInbound(context.Background(), inboundMsg)
						}
					}()
					return nodeIDs
				})

			// aggregate the signatures:
			// This should still succeed because we have 4 out of 5 valid signatures,
			// even though we're not able to get the quorum percentage buffer.
			signedMessage, err := aggregator.CreateSignedMessage(
				t.Context(),
				logging.NoLog{},
				msg,
				nil,
				subnetID,
				tc.requiredQuorumPercentage,
				tc.quorumPercentageBuffer,
				pchainapi.ProposedHeight, // Use ProposedHeight for current validators
			)
			require.NoError(t, err)

			verifyErr := signedMessage.Signature.Verify(
				msg,
				networkID,
				connectedValidators.ValidatorSet,
				tc.requiredQuorumPercentage,
				100,
			)
			require.NoError(t, verifyErr)
		})
	}
}

func TestSortedQueryableValidators(t *testing.T) {
	// Three connected validators with distinct weights. Build the canonical set so we can
	// reason about the expected indices after the (BLS-pubkey based) canonical ordering.
	connected, _ := makeConnectedValidatorsWithWeights([]uint64{5, 10, 1})

	indexByWeight := map[uint64]int{}
	for i, v := range connected.ValidatorSet.Validators {
		indexByWeight[v.Weight] = i
	}

	candidates := sortedQueryableValidators(connected, map[int][bls.SignatureLen]byte{})
	require.Len(t, candidates, 3)
	require.Equal(t, indexByWeight[10], candidates[0].index)
	require.Equal(t, indexByWeight[5], candidates[1].index)
	require.Equal(t, indexByWeight[1], candidates[2].index)
	require.Equal(t, uint64(10), candidates[0].weight)
	require.Equal(t, uint64(5), candidates[1].weight)
	require.Equal(t, uint64(1), candidates[2].weight)

	// Validators with a cached signature should be excluded from the candidate list.
	cached := map[int][bls.SignatureLen]byte{
		indexByWeight[10]: {},
	}
	candidates = sortedQueryableValidators(connected, cached)
	require.Len(t, candidates, 2)
	require.Equal(t, indexByWeight[5], candidates[0].index)
	require.Equal(t, indexByWeight[1], candidates[1].index)
}

func TestPickNextBatch(t *testing.T) {
	connected, _ := makeConnectedValidatorsWithWeights([]uint64{7, 3, 1, 1})
	candidates := sortedQueryableValidators(connected, map[int][bls.SignatureLen]byte{})
	require.Len(t, candidates, 4)

	indexToNodeID := func(idx int) ids.NodeID {
		return connected.ValidatorSet.Validators[idx].NodeIDs[0]
	}

	// Requesting 5 of weight should pull only the highest-weight validator (weight 7 >= 5).
	cursor := 0
	queried := set.NewSet[ids.NodeID](0)
	batch, batchWeight := pickNextBatch(connected, candidates, &cursor, queried, 5)
	require.Equal(t, uint64(7), batchWeight)
	require.Equal(t, 1, cursor)
	require.True(t, batch.Contains(indexToNodeID(candidates[0].index)))
	require.True(t, queried.Contains(indexToNodeID(candidates[0].index)))

	// Asking for another batch of weight 2 should pull validator with weight 3.
	batch, batchWeight = pickNextBatch(connected, candidates, &cursor, queried, 2)
	require.Equal(t, uint64(3), batchWeight)
	require.Equal(t, 2, cursor)
	require.True(t, batch.Contains(indexToNodeID(candidates[1].index)))

	// Asking for a batch larger than the remaining weight returns the rest of the validators.
	batch, batchWeight = pickNextBatch(connected, candidates, &cursor, queried, 100)
	require.Equal(t, uint64(2), batchWeight)
	require.Equal(t, 4, cursor)
	require.Equal(t, 2, batch.Len())

	// No more candidates left to query.
	batch, batchWeight = pickNextBatch(connected, candidates, &cursor, queried, 1)
	require.Equal(t, uint64(0), batchWeight)
	require.Equal(t, 0, batch.Len())
}

func TestPickNextBatchSkipsUnconnectedValidators(t *testing.T) {
	connected, _ := makeConnectedValidatorsWithWeights([]uint64{10, 5, 1})
	// Mark the highest-weight validator as not connected.
	candidates := sortedQueryableValidators(connected, map[int][bls.SignatureLen]byte{})
	require.Equal(t, uint64(10), candidates[0].weight)
	connected.ConnectedNodes.Remove(connected.ValidatorSet.Validators[candidates[0].index].NodeIDs[0])

	cursor := 0
	queried := set.NewSet[ids.NodeID](0)
	// Requesting weight 5: cursor must advance past the disconnected weight-10 validator
	// and instead include the weight-5 validator.
	batch, batchWeight := pickNextBatch(connected, candidates, &cursor, queried, 5)
	require.Equal(t, uint64(5), batchWeight)
	require.Equal(t, 2, cursor)
	require.Equal(t, 1, batch.Len())
}

func TestWeightGap(t *testing.T) {
	tests := []struct {
		name              string
		accumulatedWeight uint64
		totalWeight       uint64
		quorumPercentage  uint64
		expectedGap       uint64
	}{
		{
			name:              "zero accumulated",
			accumulatedWeight: 0,
			totalWeight:       100,
			quorumPercentage:  67,
			expectedGap:       67,
		},
		{
			name:              "threshold met exactly",
			accumulatedWeight: 67,
			totalWeight:       100,
			quorumPercentage:  67,
			expectedGap:       0,
		},
		{
			name:              "threshold exceeded",
			accumulatedWeight: 80,
			totalWeight:       100,
			quorumPercentage:  67,
			expectedGap:       0,
		},
		{
			name:              "ceil rounding",
			accumulatedWeight: 0,
			totalWeight:       7,
			quorumPercentage:  67,
			expectedGap:       5, // ceil(7 * 67 / 100) = ceil(4.69) = 5
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gap := weightGap(
				new(big.Int).SetUint64(tc.accumulatedWeight),
				tc.totalWeight,
				tc.quorumPercentage,
			)
			require.Equal(t, tc.expectedGap, gap)
		})
	}
}

// TestCreateSignedMessagePrioritizesHighWeightValidators verifies that the aggregator
// queries only the smallest number of (highest-weight) validators needed to reach the
// quorum threshold rather than fanning out to the entire validator set.
func TestCreateSignedMessagePrioritizesHighWeightValidators(t *testing.T) {
	// One dominant validator at weight 100 and four small validators at weight 1 each.
	// A 51% quorum is satisfiable by the dominant validator alone, so the aggregator
	// should issue a single batched request to just that validator.
	connectedValidators, validatorSigners := makeConnectedValidatorsWithWeights(
		[]uint64{100, 1, 1, 1, 1},
	)
	// Locate the dominant validator's canonical index.
	dominantIdx := -1
	for i, v := range connectedValidators.ValidatorSet.Validators {
		if v.Weight == 100 {
			dominantIdx = i
			break
		}
	}
	require.NotEqual(t, -1, dominantIdx)
	dominantNodeID := connectedValidators.ValidatorSet.Validators[dominantIdx].NodeIDs[0]

	chainID := ids.GenerateTestID()
	networkID := constants.UnitTestID
	msg, err := warp.NewUnsignedMessage(networkID, chainID, utils.RandomBytes(64))
	require.NoError(t, err)

	aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateDefaultAggregator(t)

	subnetID := ids.GenerateTestID()
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(subnetID, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), subnetID).Return(
		connectedValidators.ValidatorSet, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{subnetID: connectedValidators.ValidatorSet}, nil,
	).AnyTimes()

	var peerInfos []peer.Info
	for nodeID := range connectedValidators.ConnectedNodes {
		peerInfos = append(peerInfos, peer.Info{ID: nodeID})
	}
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()
	mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{}, nil,
	).Times(1)

	// The aggregator should only ever issue a single request, and it must target only
	// the dominant validator since its weight alone meets the 51% threshold.
	mockNetwork.EXPECT().Send(
		gomock.Any(), gomock.Any(), subnetID, subnets.NoOpAllower,
	).Times(1).DoAndReturn(
		func(
			outboundMsg *message.OutboundMessage,
			config interface{},
			subnetID ids.ID,
			allower interface{},
		) set.Set[ids.NodeID] {
			// Verify that only the dominant validator was targeted by this Send call.
			sendConfig, ok := config.(avagocommon.SendConfig)
			require.True(t, ok)
			require.Equal(t, 1, sendConfig.NodeIDs.Len())
			require.True(t, sendConfig.NodeIDs.Contains(dominantNodeID))

			currentRequestID := aggregator.currentRequestID.Load()
			go func() {
				time.Sleep(10 * time.Millisecond)
				signature, err := validatorSigners[dominantIdx].Sign(msg.Bytes())
				if err != nil {
					t.Logf("failed to sign: %v", err)
					return
				}
				responseBytes, err := proto.Marshal(&sdk.SignatureResponse{
					Signature: bls.SignatureToBytes(signature),
				})
				if err != nil {
					t.Logf("failed to marshal: %v", err)
					return
				}
				handler.HandleInbound(
					context.Background(),
					message.InboundAppResponse(chainID, currentRequestID, responseBytes, dominantNodeID),
				)
			}()
			return sendConfig.NodeIDs
		},
	)

	signedMessage, err := aggregator.CreateSignedMessage(
		t.Context(), logging.NoLog{}, msg, nil, subnetID,
		51, 0, // 51% required quorum, no buffer.
		pchainapi.ProposedHeight,
	)
	require.NoError(t, err)
	require.NoError(t, signedMessage.Signature.Verify(
		msg, networkID, connectedValidators.ValidatorSet, 51, 100,
	))
}

// TestCreateSignedMessageAdvancesToNextBatchOnTimeout verifies that when the highest-weight
// validators fail to respond within the batch timeout window, the aggregator falls through
// to query additional validators rather than waiting indefinitely.
func TestCreateSignedMessageAdvancesToNextBatchOnTimeout(t *testing.T) {
	// Dominant validator (weight 100) that will never respond, plus four small validators
	// (weight 1 each) that will respond. The required 3% quorum can be met by either the
	// dominant validator alone OR three of the small validators (3/104 ≈ 2.88%, ceil to 4).
	// In practice ceil(104 * 3 / 100) = 4, so we need >= 4 weight from the small ones.
	connectedValidators, validatorSigners := makeConnectedValidatorsWithWeights(
		[]uint64{100, 1, 1, 1, 1},
	)
	dominantIdx := -1
	smallIndices := []int{}
	for i, v := range connectedValidators.ValidatorSet.Validators {
		if v.Weight == 100 {
			dominantIdx = i
		} else {
			smallIndices = append(smallIndices, i)
		}
	}
	require.NotEqual(t, -1, dominantIdx)
	require.Len(t, smallIndices, 4)

	chainID := ids.GenerateTestID()
	networkID := constants.UnitTestID
	msg, err := warp.NewUnsignedMessage(networkID, chainID, utils.RandomBytes(64))
	require.NoError(t, err)

	aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateDefaultAggregator(t)

	subnetID := ids.GenerateTestID()
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(subnetID, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), subnetID).Return(
		connectedValidators.ValidatorSet, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{subnetID: connectedValidators.ValidatorSet}, nil,
	).AnyTimes()

	var peerInfos []peer.Info
	for nodeID := range connectedValidators.ConnectedNodes {
		peerInfos = append(peerInfos, peer.Info{ID: nodeID})
	}
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()
	mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{}, nil,
	).Times(1)

	// Track how many distinct Sends we observe so we can ensure the aggregator advanced
	// past the unresponsive dominant validator and queried at least one small validator.
	var sendCallCount int32
	mockNetwork.EXPECT().Send(
		gomock.Any(), gomock.Any(), subnetID, subnets.NoOpAllower,
	).MinTimes(2).DoAndReturn(
		func(
			outboundMsg *message.OutboundMessage,
			config interface{},
			subnetID ids.ID,
			allower interface{},
		) set.Set[ids.NodeID] {
			callIdx := atomic.AddInt32(&sendCallCount, 1)
			sendConfig, ok := config.(avagocommon.SendConfig)
			require.True(t, ok)
			currentRequestID := aggregator.currentRequestID.Load()

			if callIdx == 1 {
				// First batch: must target only the dominant validator. Do not respond
				// so that the per-batch timeout fires and the aggregator moves on.
				require.True(t, sendConfig.NodeIDs.Contains(
					connectedValidators.ValidatorSet.Validators[dominantIdx].NodeIDs[0],
				))
				return sendConfig.NodeIDs
			}

			// Subsequent batch(es): respond on behalf of every small validator targeted.
			go func() {
				time.Sleep(10 * time.Millisecond)
				for nodeID := range sendConfig.NodeIDs {
					idx := connectedValidators.NodeValidatorIndexMap[nodeID]
					signature, err := validatorSigners[idx].Sign(msg.Bytes())
					if err != nil {
						t.Logf("failed to sign: %v", err)
						continue
					}
					responseBytes, err := proto.Marshal(&sdk.SignatureResponse{
						Signature: bls.SignatureToBytes(signature),
					})
					if err != nil {
						t.Logf("failed to marshal: %v", err)
						continue
					}
					handler.HandleInbound(
						context.Background(),
						message.InboundAppResponse(chainID, currentRequestID, responseBytes, nodeID),
					)
				}
			}()
			return sendConfig.NodeIDs
		},
	)

	signedMessage, err := aggregator.CreateSignedMessage(
		t.Context(), logging.NoLog{}, msg, nil, subnetID,
		3, 0, // 3% required quorum, no buffer.
		pchainapi.ProposedHeight,
	)
	require.NoError(t, err)
	require.GreaterOrEqual(t, atomic.LoadInt32(&sendCallCount), int32(2))
	require.NoError(t, signedMessage.Signature.Verify(
		msg, networkID, connectedValidators.ValidatorSet, 3, 100,
	))
}

func TestUnmarshalResponse(t *testing.T) {
	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)

	emptySignatureResponse, err := proto.Marshal(&sdk.SignatureResponse{Signature: []byte{}})
	require.NoError(t, err)

	randSignature := make([]byte, 96)
	_, err = rand.Read(randSignature)
	require.NoError(t, err)

	randSignatureResponse, err := proto.Marshal(&sdk.SignatureResponse{Signature: randSignature})
	require.NoError(t, err)

	testCases := []struct {
		name              string
		appResponseBytes  []byte
		expectedSignature blsSignatureBuf
	}{
		{
			name:              "empty slice",
			appResponseBytes:  []byte{},
			expectedSignature: blsSignatureBuf{},
		},
		{
			name:              "nil slice",
			appResponseBytes:  nil,
			expectedSignature: blsSignatureBuf{},
		},
		{
			name:              "empty signature",
			appResponseBytes:  emptySignatureResponse,
			expectedSignature: blsSignatureBuf{},
		},
		{
			name:              "random signature",
			appResponseBytes:  randSignatureResponse,
			expectedSignature: blsSignatureBuf(randSignature),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			signature, err := aggregator.unmarshalResponse(tc.appResponseBytes)
			require.NoError(t, err)
			require.Equal(t, tc.expectedSignature, signature)
		})
	}
}

func TestGetExcludedValidators(t *testing.T) {
	underFunded := minimumL1ValidatorBalance - 1
	funded := minimumL1ValidatorBalance

	nodeID1 := ids.GenerateTestNodeID()
	validationID1 := ids.GenerateTestID()
	nodeID2 := ids.GenerateTestNodeID()
	validationID2 := ids.GenerateTestID()
	nodeID3 := ids.GenerateTestNodeID()
	validationID3 := ids.GenerateTestID()

	testCases := []struct {
		name         string
		l1Validators []platformvm.ClientPermissionlessValidator
		connected    *peers.CanonicalValidators
		excludedIdx  []int // Indices of validators that should be excluded
	}{
		{
			name: "all underfunded",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      &underFunded,
					},
				},
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID2},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID2,
						Balance:      &underFunded,
					},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1}},
						{NodeIDs: []ids.NodeID{nodeID2}},
					},
				},
			},
			excludedIdx: []int{0, 1},
		},
		{
			name: "all funded",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      &funded,
					},
				},
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID2},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID2,
						Balance:      &funded,
					},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1}},
						{NodeIDs: []ids.NodeID{nodeID2}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name: "one underfunded, one funded",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      &funded,
					},
				},
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID2},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID2,
						Balance:      &funded,
					},
				},
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID3},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID3,
						Balance:      &underFunded,
					},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1}},
						{NodeIDs: []ids.NodeID{nodeID2, nodeID3}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name: "mixed L1/non-L1",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      &funded,
					},
				},
				{
					// non-L1
					ClientStaker: platformvm.ClientStaker{
						NodeID: nodeID2,
					},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1}},
						{NodeIDs: []ids.NodeID{nodeID2}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name: "nil balance",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      nil,
					},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1}},
					},
				},
			},
			excludedIdx: []int{0},
		},
		{
			name: "multiple nodeIDs per validator",
			l1Validators: []platformvm.ClientPermissionlessValidator{
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID1},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID1,
						Balance:      &funded},
				},
				{
					ClientStaker: platformvm.ClientStaker{NodeID: nodeID2},
					ClientL1Validator: platformvm.ClientL1Validator{
						ValidationID: &validationID2,
						Balance:      &funded},
				},
			},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID1, nodeID2}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name:         "no L1 validators",
			l1Validators: []platformvm.ClientPermissionlessValidator{},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{
						{NodeIDs: []ids.NodeID{nodeID3}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name:         "empty validator set",
			l1Validators: []platformvm.ClientPermissionlessValidator{},
			connected: &peers.CanonicalValidators{
				ValidatorSet: validators.WarpSet{
					Validators: []*validators.Warp{},
				},
			},
			excludedIdx: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aggregator, _, _, _, mockValidatorClient := instantiateDefaultAggregator(t)
			ctx := t.Context()
			log := logging.NoLog{}
			signingSubnet := ids.GenerateTestID()

			mockValidatorClient.EXPECT().
				GetCurrentValidators(gomock.Any(), signingSubnet).
				Return(tc.l1Validators, nil)

			excluded, err := aggregator.getExcludedValidators(ctx, log, signingSubnet, tc.connected)
			require.NoError(t, err)
			for idx := range tc.connected.ValidatorSet.Validators {
				shouldBeExcluded := slices.Contains(tc.excludedIdx, idx)
				if shouldBeExcluded {
					require.True(t, excluded.Contains(idx), "validator %d should be excluded", idx)
				} else {
					require.False(t, excluded.Contains(idx), "validator %d should NOT be excluded", idx)
				}
			}
		})
	}
}

func TestValidateQuorumPercentages(t *testing.T) {
	tests := []struct {
		name     string
		required uint64
		buffer   uint64
		wantErr  bool
	}{
		{
			name:     "valid",
			required: 80,
			buffer:   5,
			wantErr:  false,
		},
		{
			name:     "zero required",
			required: 0,
			buffer:   5,
			wantErr:  true},
		{
			name:     "sum over 100",
			required: 98, buffer: 5,
			wantErr: true,
		},
		{
			name:     "exactly 100",
			required: 100,
			buffer:   0,
			wantErr:  false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateQuorumPercentages(tc.required, tc.buffer)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSelectSigningSubnet(t *testing.T) {
	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	ctx := t.Context()
	log := logging.NoLog{}
	chainID := ids.GenerateTestID()
	msg, err := warp.NewUnsignedMessage(0, chainID, []byte{})
	require.NoError(t, err)

	// Mock getSubnetID to return a specific subnet
	wantSubnet := ids.GenerateTestID()
	aggregator.subnetIDsByBlockchainID[chainID] = wantSubnet

	// Case: inputSigningSubnet is Empty
	signingSubnet, sourceSubnet, err := aggregator.selectSigningSubnet(ctx, log, msg, ids.Empty)
	require.NoError(t, err)
	require.Equal(t, wantSubnet, signingSubnet)
	require.Equal(t, wantSubnet, sourceSubnet)

	// Case: inputSigningSubnet is set
	otherSubnet := ids.GenerateTestID()
	signingSubnet, sourceSubnet, err = aggregator.selectSigningSubnet(ctx, log, msg, otherSubnet)
	require.NoError(t, err)
	require.Equal(t, otherSubnet, signingSubnet)
	require.Equal(t, wantSubnet, sourceSubnet)
}

func TestPopulateSignatureMapFromCache(t *testing.T) {
	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	connectedValidators, signers := makeConnectedValidators(2)
	msg, err := warp.NewUnsignedMessage(0, ids.GenerateTestID(), []byte("test"))
	require.NoError(t, err)

	// Simulate a cached signature for the first validator
	sig, err := signers[0].Sign(msg.Bytes())
	require.NoError(t, err)
	pubKeyBytes := bls.PublicKeyToUncompressedBytes(signers[0].PublicKey())

	// Add the signature to the aggregator's cache
	aggregator.signatureCache.Add(
		msg.ID(),
		PublicKeyBytes(pubKeyBytes),
		SignatureBytes(bls.SignatureToBytes(sig)),
	)

	excluded := set.NewSet[int](0)
	sigMap, accWeight := aggregator.getCachedSignaturesForMessage(logging.NoLog{}, msg, connectedValidators, excluded)
	require.Len(t, sigMap, 1)
	// The expected weight is the weight of the first validator
	require.Equal(t, connectedValidators.ValidatorSet.Validators[0].Weight, accWeight.Uint64())
}

// makeConnectedValidatorsWithWeights creates len(weights) connected validators,
// assigning each its corresponding weight. Validators are returned in canonical
// (pubkey) order, paired with their signers in the same order.
func makeConnectedValidatorsWithWeights(
	weights []uint64,
) (*peers.CanonicalValidators, []*localsigner.LocalSigner) {
	count := len(weights)
	infos := make([]validatorInfo, count)
	for i := 0; i < count; i++ {
		localSigner, err := localsigner.New()
		if err != nil {
			panic(err)
		}
		pubKey := localSigner.PublicKey()
		infos[i] = validatorInfo{
			nodeID:            ids.GenerateTestNodeID(),
			blsSigner:         localSigner,
			blsPublicKey:      pubKey,
			blsPublicKeyBytes: bls.PublicKeyToUncompressedBytes(pubKey),
			weight:            weights[i],
		}
	}

	// Canonical order is by uncompressed public-key bytes. Sort moves weights
	// alongside their owning validatorInfo, so post-sort indices stay aligned.
	utils.Sort(infos)

	validatorSet := make([]*validators.Warp, count)
	validatorSigners := make([]*localsigner.LocalSigner, count)
	nodeValidatorIndexMap := make(map[ids.NodeID]int, count)
	connectedNodes := set.NewSet[ids.NodeID](count)
	var totalWeight uint64
	for i, info := range infos {
		validatorSigners[i] = info.blsSigner
		validatorSet[i] = &validators.Warp{
			PublicKey:      info.blsPublicKey,
			PublicKeyBytes: info.blsPublicKeyBytes,
			Weight:         info.weight,
			NodeIDs:        []ids.NodeID{info.nodeID},
		}
		nodeValidatorIndexMap[info.nodeID] = i
		connectedNodes.Add(info.nodeID)
		totalWeight += info.weight
	}

	return &peers.CanonicalValidators{
		ConnectedWeight: totalWeight,
		ConnectedNodes:  connectedNodes,
		ValidatorSet: validators.WarpSet{
			Validators:  validatorSet,
			TotalWeight: totalWeight,
		},
		NodeValidatorIndexMap: nodeValidatorIndexMap,
	}, validatorSigners
}

// buildFullSignatureMap signs unsignedMessage with every signer and returns the
// resulting map keyed by canonical validator index along with the cumulative
// signed weight.
func buildFullSignatureMap(
	t *testing.T,
	unsignedMessage *warp.UnsignedMessage,
	vdrs *peers.CanonicalValidators,
	signers []*localsigner.LocalSigner,
) (map[int][bls.SignatureLen]byte, *big.Int) {
	t.Helper()
	require.Len(t, signers, len(vdrs.ValidatorSet.Validators))

	sigMap := make(map[int][bls.SignatureLen]byte, len(signers))
	totalSigned := big.NewInt(0)
	for i, signer := range signers {
		sig, err := signer.Sign(unsignedMessage.Bytes())
		require.NoError(t, err)
		sigMap[i] = [bls.SignatureLen]byte(bls.SignatureToBytes(sig))
		totalSigned.Add(totalSigned, new(big.Int).SetUint64(vdrs.ValidatorSet.Validators[i].Weight))
	}
	return sigMap, totalSigned
}

// signerIndices returns the sorted list of signer indices encoded in a BitSetSignature.
func signerIndices(t *testing.T, msg *warp.Message) []int {
	t.Helper()
	bitSetSig, ok := msg.Signature.(*warp.BitSetSignature)
	require.True(t, ok)
	bits := set.BitsFromBytes(bitSetSig.Signers)
	out := make([]int, 0, bits.Len())
	for i := 0; i < bits.BitLen(); i++ {
		if bits.Contains(i) {
			out = append(out, i)
		}
	}
	return out
}

func TestAggregateIfSufficientWeight_PrunesGreedyByWeight(t *testing.T) {
	const networkID = constants.UnitTestID
	// Weights chosen so that greedy-by-weight-desc is unambiguous:
	// total = 28; 67% threshold = 18.76 → need >=19.
	// Largest-first: 10 → 10 (insufficient), 10+8 → 18 (insufficient), 10+8+5 → 23 (ok).
	// Any subset of size 2 has weight at most 10+8 = 18 < 19. Greedy selection of
	// size 3 is the unique minimum-cardinality subset that meets quorum.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{10, 8, 5, 4, 1})

	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("greedy"))
	require.NoError(t, err)
	sigMap, signedWeight := buildFullSignatureMap(t, unsigned, vdrs, signers)

	signed, err := aggregator.aggregateIfSufficientWeight(
		logging.NoLog{},
		unsigned,
		sigMap,
		signedWeight,
		vdrs.ValidatorSet.Validators,
		vdrs.ValidatorSet.TotalWeight,
		67,
	)
	require.NoError(t, err)
	require.NotNil(t, signed)

	// Pruned set must consist of the 3 heaviest validators by weight (10, 8, 5).
	got := signerIndices(t, signed)
	require.Len(t, got, 3)
	gotWeights := make([]uint64, 0, len(got))
	for _, i := range got {
		gotWeights = append(gotWeights, vdrs.ValidatorSet.Validators[i].Weight)
	}
	slices.Sort(gotWeights)
	require.Equal(t, []uint64{5, 8, 10}, gotWeights)

	require.NoError(t, signed.Signature.Verify(
		unsigned,
		networkID,
		vdrs.ValidatorSet,
		67,
		100,
	))
}

func TestAggregateIfSufficientWeight_SingleHeavySigner(t *testing.T) {
	const networkID = constants.UnitTestID
	// Validator 0 alone exceeds 67% of total weight; pruning must keep only that signer.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{70, 10, 10, 10})

	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("heavy"))
	require.NoError(t, err)
	sigMap, signedWeight := buildFullSignatureMap(t, unsigned, vdrs, signers)

	signed, err := aggregator.aggregateIfSufficientWeight(
		logging.NoLog{},
		unsigned,
		sigMap,
		signedWeight,
		vdrs.ValidatorSet.Validators,
		vdrs.ValidatorSet.TotalWeight,
		67,
	)
	require.NoError(t, err)
	require.NotNil(t, signed)

	got := signerIndices(t, signed)
	require.Len(t, got, 1)
	require.Equal(t, uint64(70), vdrs.ValidatorSet.Validators[got[0]].Weight)

	require.NoError(t, signed.Signature.Verify(
		unsigned,
		networkID,
		vdrs.ValidatorSet,
		67,
		100,
	))
}

func TestAggregateIfSufficientWeight_NoPruningWhenAllSignersNeeded(t *testing.T) {
	const networkID = constants.UnitTestID
	// 3 equally-weighted validators at 67% quorum: 2/3 = 66.67% < 67%, so all 3 are needed.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1})

	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("tight"))
	require.NoError(t, err)
	sigMap, signedWeight := buildFullSignatureMap(t, unsigned, vdrs, signers)

	signed, err := aggregator.aggregateIfSufficientWeight(
		logging.NoLog{},
		unsigned,
		sigMap,
		signedWeight,
		vdrs.ValidatorSet.Validators,
		vdrs.ValidatorSet.TotalWeight,
		67,
	)
	require.NoError(t, err)
	require.NotNil(t, signed)
	require.Len(t, signerIndices(t, signed), 3)
}

func TestAggregateIfSufficientWeight_BelowQuorumReturnsNil(t *testing.T) {
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1, 1, 1})

	aggregator, _, _, _, _ := instantiateDefaultAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(constants.UnitTestID, ids.GenerateTestID(), []byte("low"))
	require.NoError(t, err)

	// Only 2/5 = 40% signed; 60% quorum is not met.
	partialMap := make(map[int][bls.SignatureLen]byte, 2)
	signedWeight := big.NewInt(0)
	for i := 0; i < 2; i++ {
		sig, err := signers[i].Sign(unsigned.Bytes())
		require.NoError(t, err)
		partialMap[i] = [bls.SignatureLen]byte(bls.SignatureToBytes(sig))
		signedWeight.Add(signedWeight, new(big.Int).SetUint64(vdrs.ValidatorSet.Validators[i].Weight))
	}

	signed, err := aggregator.aggregateIfSufficientWeight(
		logging.NoLog{},
		unsigned,
		partialMap,
		signedWeight,
		vdrs.ValidatorSet.Validators,
		vdrs.ValidatorSet.TotalWeight,
		60,
	)
	require.NoError(t, err)
	require.Nil(t, signed)
}
