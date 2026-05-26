package aggregator

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	"github.com/ava-labs/avalanchego/snow/engine/common"
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

func instantiateAggregator(t *testing.T) (
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

// makeConnectedValidatorsWithWeights builds a canonical validator set where
// each validator has the corresponding weight from the provided slice. All
// validators are marked as connected. The returned slice of signers is
// indexed in canonical (public-key sorted) order, matching the order of the
// validators in ValidatorSet.Validators.
func makeConnectedValidatorsWithWeights(weights []uint64) (*peers.CanonicalValidators, []*localsigner.LocalSigner) {
	type weightedValidator struct {
		info   validatorInfo
		weight uint64
	}

	validatorValues := make([]weightedValidator, len(weights))
	for i, w := range weights {
		localSigner, err := localsigner.New()
		if err != nil {
			panic(err)
		}
		pubKey := localSigner.PublicKey()
		nodeID := ids.GenerateTestNodeID()
		validatorValues[i] = weightedValidator{
			info: validatorInfo{
				nodeID:            nodeID,
				blsSigner:         localSigner,
				blsPublicKey:      pubKey,
				blsPublicKeyBytes: bls.PublicKeyToUncompressedBytes(pubKey),
			},
			weight: w,
		}
	}

	// Sort by public key to match canonical ordering used in the aggregator.
	slices.SortFunc(validatorValues, func(a, b weightedValidator) int {
		return a.info.Compare(b.info)
	})

	totalWeight := uint64(0)
	validatorSet := make([]*validators.Warp, len(weights))
	validatorSigners := make([]*localsigner.LocalSigner, len(weights))
	nodeValidatorIndexMap := make(map[ids.NodeID]int)
	connectedNodes := set.NewSet[ids.NodeID](len(weights))
	for i, validator := range validatorValues {
		validatorSigners[i] = validator.info.blsSigner
		validatorSet[i] = &validators.Warp{
			PublicKey:      validator.info.blsPublicKey,
			PublicKeyBytes: validator.info.blsPublicKeyBytes,
			Weight:         validator.weight,
			NodeIDs:        []ids.NodeID{validator.info.nodeID},
		}
		nodeValidatorIndexMap[validator.info.nodeID] = i
		connectedNodes.Add(validator.info.nodeID)
		totalWeight += validator.weight
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
			aggregator, _, _, _, _ := instantiateAggregator(t)
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
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateAggregator(t)
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
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateAggregator(t)
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
	aggregator, _, _, mockNetwork, mockValidatorClient := instantiateAggregator(t)

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

			aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateAggregator(t)

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

func TestUnmarshalResponse(t *testing.T) {
	aggregator, _, _, _, _ := instantiateAggregator(t)

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
			aggregator, _, _, _, mockValidatorClient := instantiateAggregator(t)
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
	aggregator, _, _, _, _ := instantiateAggregator(t)
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
	aggregator, _, _, _, _ := instantiateAggregator(t)
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

func TestSortedCandidateValidators(t *testing.T) {
	// Build a validator set with weights at canonical indices [10, 50, 30, 20, 40].
	// We will:
	//   - drop the only nodeID of index 2 from ConnectedNodes (unreachable),
	//   - mark index 4 as excluded (insufficient L1 balance),
	//   - pretend index 1 already has a cached signature.
	// Expected candidates ordered by weight desc: [3 (w=20), 0 (w=10)].
	weights := []uint64{10, 50, 30, 20, 40}
	nodeIDs := make([]ids.NodeID, len(weights))
	validatorSet := make([]*validators.Warp, len(weights))
	connectedNodes := set.NewSet[ids.NodeID](len(weights))
	for i, w := range weights {
		nodeIDs[i] = ids.GenerateTestNodeID()
		validatorSet[i] = &validators.Warp{
			Weight:  w,
			NodeIDs: []ids.NodeID{nodeIDs[i]},
		}
		if i != 2 {
			connectedNodes.Add(nodeIDs[i])
		}
	}

	vdrs := &peers.CanonicalValidators{
		ConnectedNodes: connectedNodes,
		ValidatorSet: validators.WarpSet{
			Validators:  validatorSet,
			TotalWeight: 150,
		},
	}

	signatureMap := map[int][bls.SignatureLen]byte{1: {}}
	excluded := set.NewSet[int](0)
	excluded.Add(4)

	got := sortedCandidateValidators(vdrs, signatureMap, excluded)
	require.Equal(t, []int{3, 0}, got)
}

func TestSortedCandidateValidatorsStableForEqualWeight(t *testing.T) {
	weights := []uint64{5, 5, 5}
	nodeIDs := make([]ids.NodeID, len(weights))
	validatorSet := make([]*validators.Warp, len(weights))
	connectedNodes := set.NewSet[ids.NodeID](len(weights))
	for i, w := range weights {
		nodeIDs[i] = ids.GenerateTestNodeID()
		validatorSet[i] = &validators.Warp{Weight: w, NodeIDs: []ids.NodeID{nodeIDs[i]}}
		connectedNodes.Add(nodeIDs[i])
	}
	vdrs := &peers.CanonicalValidators{
		ConnectedNodes: connectedNodes,
		ValidatorSet: validators.WarpSet{
			Validators:  validatorSet,
			TotalWeight: 15,
		},
	}
	got := sortedCandidateValidators(vdrs, nil, set.NewSet[int](0))
	// Stable sort preserves canonical order when weights are equal.
	require.Equal(t, []int{0, 1, 2}, got)
}

func TestPickNextBatch(t *testing.T) {
	// Validators already in priority order (highest weight first).
	vdrs := &peers.CanonicalValidators{
		ValidatorSet: validators.WarpSet{
			Validators: []*validators.Warp{
				{Weight: 40}, {Weight: 30}, {Weight: 20}, {Weight: 10}, {Weight: 5},
			},
			TotalWeight: 105,
		},
	}
	candidates := []int{0, 1, 2, 3, 4}

	// Need projected weight to exceed 50% of 105 (= 52.5). 40 alone is not
	// enough; 40+30=70 crosses the threshold. Expect a 2-element batch.
	batch, newOffset := pickNextBatch(vdrs, candidates, 0, big.NewInt(0), 50)
	require.Equal(t, []int{0, 1}, batch)
	require.Equal(t, 2, newOffset)

	// Starting with 70 already accumulated and aiming for 80% (= 84), we need
	// 14 more. Index 2 alone (weight 20) brings projected to 90 which exceeds
	// the threshold. Expect a 1-element batch.
	batch, newOffset = pickNextBatch(vdrs, candidates, 2, big.NewInt(70), 80)
	require.Equal(t, []int{2}, batch)
	require.Equal(t, 3, newOffset)

	// No remaining candidates returns an empty batch.
	batch, newOffset = pickNextBatch(vdrs, candidates, 5, big.NewInt(0), 80)
	require.Empty(t, batch)
	require.Equal(t, 5, newOffset)

	// If candidates can't cover the target, return everything left.
	batch, newOffset = pickNextBatch(vdrs, candidates, 0, big.NewInt(0), 100)
	require.Equal(t, []int{0, 1, 2, 3, 4}, batch)
	require.Equal(t, 5, newOffset)
}

func TestCreateSignedMessageSendsOnlyToTopWeightedValidators(t *testing.T) {
	// Build 5 validators with weights [1, 2, 3, 4, 5]. Total weight = 15.
	// Required quorum = 67%, no buffer => target weight = 11.
	// Top-weighted picks: 5 + 4 + 3 = 12, which is the smallest prefix that
	// crosses the threshold. We expect the aggregator to send to the top 3
	// validators in a single batch and not query the remaining 2.
	msg, err := warp.NewUnsignedMessage(
		constants.UnitTestID,
		ids.GenerateTestID(),
		utils.RandomBytes(64),
	)
	require.NoError(t, err)

	connectedValidators, validatorSigners := makeConnectedValidatorsWithWeights([]uint64{1, 2, 3, 4, 5})

	// Identify the nodeIDs corresponding to the top 3 validators by weight.
	type vdrEntry struct {
		index  int
		weight uint64
		nodeID ids.NodeID
	}
	entries := make([]vdrEntry, len(connectedValidators.ValidatorSet.Validators))
	for i, v := range connectedValidators.ValidatorSet.Validators {
		entries[i] = vdrEntry{i, v.Weight, v.NodeIDs[0]}
	}
	slices.SortFunc(entries, func(a, b vdrEntry) int {
		switch {
		case a.weight > b.weight:
			return -1
		case a.weight < b.weight:
			return 1
		default:
			return 0
		}
	})
	expectedTopNodes := set.NewSet[ids.NodeID](3)
	for i := 0; i < 3; i++ {
		expectedTopNodes.Add(entries[i].nodeID)
	}

	aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateAggregator(t)
	subnetID := ids.GenerateTestID()
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), msg.SourceChainID).Return(subnetID, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), subnetID).Return(
		connectedValidators.ValidatorSet, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{subnetID: connectedValidators.ValidatorSet}, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{}, nil,
	).Times(1)

	peerInfos := make([]peer.Info, 0, connectedValidators.ConnectedNodes.Len())
	for nodeID := range connectedValidators.ConnectedNodes {
		peerInfos = append(peerInfos, peer.Info{ID: nodeID})
	}
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()

	requestID := aggregator.currentRequestID.Load() + 2

	// Expect exactly one Send call and assert it targets only the top
	// 3 validators (not all 5).
	mockNetwork.EXPECT().Send(
		gomock.Any(),
		gomock.Any(),
		subnetID,
		subnets.NoOpAllower,
	).Times(1).DoAndReturn(
		func(
			_ *message.OutboundMessage,
			config common.SendConfig,
			_ ids.ID,
			_ interface{},
		) set.Set[ids.NodeID] {
			require.Equal(t, expectedTopNodes.Len(), config.NodeIDs.Len(),
				"expected Send to target exactly the top 3 weighted validators")
			for nodeID := range config.NodeIDs {
				require.True(t, expectedTopNodes.Contains(nodeID),
					"unexpected non-top-weighted validator %s in batch", nodeID)
			}

			go func() {
				time.Sleep(10 * time.Millisecond)
				for nodeID := range config.NodeIDs {
					vdrIdx := connectedValidators.NodeValidatorIndexMap[nodeID]
					sig, signErr := validatorSigners[vdrIdx].Sign(msg.Bytes())
					if signErr != nil {
						t.Logf("Failed to sign: %v", signErr)
						continue
					}
					respBytes, mErr := proto.Marshal(&sdk.SignatureResponse{
						Signature: bls.SignatureToBytes(sig),
					})
					if mErr != nil {
						t.Logf("Failed to marshal: %v", mErr)
						continue
					}
					inboundMsg := message.InboundAppResponse(
						msg.SourceChainID, requestID, respBytes, nodeID,
					)
					handler.HandleInbound(t.Context(), inboundMsg)
				}
			}()
			return config.NodeIDs
		},
	)

	signedMessage, err := aggregator.CreateSignedMessage(
		t.Context(),
		logging.NoLog{},
		msg,
		nil,
		subnetID,
		67, // required
		0,  // buffer
		pchainapi.ProposedHeight,
	)
	require.NoError(t, err)
	require.NoError(t, signedMessage.Signature.Verify(
		msg,
		constants.UnitTestID,
		connectedValidators.ValidatorSet,
		67,
		100,
	))
}

func TestCreateSignedMessageDispatchesAdditionalBatchAfterShortfall(t *testing.T) {
	// 5 equal-weight validators. Required quorum = 80%, no buffer => target
	// weight = 4. First batch is the top 4 by canonical order. The test
	// returns an invalid signature from one of those validators so the first
	// batch only contributes 3 weight (60%), forcing the aggregator to
	// dispatch a second batch that picks the remaining validator and
	// reaches quorum.
	msg, err := warp.NewUnsignedMessage(
		constants.UnitTestID,
		ids.GenerateTestID(),
		utils.RandomBytes(64),
	)
	require.NoError(t, err)

	connectedValidators, validatorSigners := makeConnectedValidators(5)

	aggregator, _, handler, mockNetwork, mockValidatorClient := instantiateAggregator(t)
	subnetID := ids.GenerateTestID()
	mockValidatorClient.EXPECT().GetSubnetID(gomock.Any(), msg.SourceChainID).Return(subnetID, nil).AnyTimes()
	mockValidatorClient.EXPECT().GetProposedValidators(gomock.Any(), subnetID).Return(
		connectedValidators.ValidatorSet, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetAllValidatorSets(gomock.Any(), gomock.Any()).Return(
		map[ids.ID]validators.WarpSet{subnetID: connectedValidators.ValidatorSet}, nil,
	).AnyTimes()
	mockValidatorClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{}, nil,
	).Times(1)
	peerInfos := make([]peer.Info, 0, connectedValidators.ConnectedNodes.Len())
	for nodeID := range connectedValidators.ConnectedNodes {
		peerInfos = append(peerInfos, peer.Info{ID: nodeID})
	}
	mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfos).AnyTimes()

	// Map canonical index -> nodeID for response injection convenience.
	indexToNodeID := make(map[int]ids.NodeID)
	for nodeID, idx := range connectedValidators.NodeValidatorIndexMap {
		indexToNodeID[idx] = nodeID
	}

	var (
		sendsMu  sync.Mutex
		sendsSeq []set.Set[ids.NodeID]
	)

	mockNetwork.EXPECT().Send(
		gomock.Any(),
		gomock.Any(),
		subnetID,
		subnets.NoOpAllower,
	).MinTimes(2).DoAndReturn(
		func(
			_ *message.OutboundMessage,
			config common.SendConfig,
			_ ids.ID,
			_ interface{},
		) set.Set[ids.NodeID] {
			sendsMu.Lock()
			batchIdx := len(sendsSeq)
			sendsSeq = append(sendsSeq, config.NodeIDs)
			sendsMu.Unlock()

			// requestID for the Nth batch is the initial request ID + 2*N.
			batchReqID := aggregator.currentRequestID.Load()

			go func(batchReqID uint32, batchIdx int, batchNodes set.Set[ids.NodeID]) {
				time.Sleep(10 * time.Millisecond)
				for nodeID := range batchNodes {
					var sigBytes []byte
					vdrIdx := connectedValidators.NodeValidatorIndexMap[nodeID]
					// On the first batch (which should have 4 nodes), return
					// an invalid (empty) signature from the lowest canonical
					// index so that only 3 valid signatures are collected
					// and the aggregator must dispatch a second batch.
					if batchIdx == 0 && vdrIdx == 0 {
						sigBytes = []byte{}
					} else {
						sig, signErr := validatorSigners[vdrIdx].Sign(msg.Bytes())
						if signErr != nil {
							t.Logf("Failed to sign: %v", signErr)
							continue
						}
						sigBytes = bls.SignatureToBytes(sig)
					}
					respBytes, mErr := proto.Marshal(&sdk.SignatureResponse{
						Signature: sigBytes,
					})
					if mErr != nil {
						t.Logf("Failed to marshal: %v", mErr)
						continue
					}
					inboundMsg := message.InboundAppResponse(
						msg.SourceChainID, batchReqID, respBytes, nodeID,
					)
					handler.HandleInbound(t.Context(), inboundMsg)
				}
			}(batchReqID, batchIdx, config.NodeIDs)
			return config.NodeIDs
		},
	)

	signedMessage, err := aggregator.CreateSignedMessage(
		t.Context(),
		logging.NoLog{},
		msg,
		nil,
		subnetID,
		80,
		0,
		pchainapi.ProposedHeight,
	)
	require.NoError(t, err)
	require.NoError(t, signedMessage.Signature.Verify(
		msg,
		constants.UnitTestID,
		connectedValidators.ValidatorSet,
		80,
		100,
	))

	sendsMu.Lock()
	defer sendsMu.Unlock()
	require.GreaterOrEqual(t, len(sendsSeq), 2,
		"expected at least 2 Send invocations after first batch shortfall")
	// First batch should target only the top 4 by canonical order (not all 5).
	require.Equal(t, 4, sendsSeq[0].Len(),
		"first batch should target the 4-validator prefix needed for 80%")
	// Second batch should only target the remaining 1 validator needed to
	// fill the deficit, not the validators we already asked.
	require.Equal(t, 1, sendsSeq[1].Len(),
		"second batch should query just the additional validator needed")
	for nodeID := range sendsSeq[1] {
		require.False(t, sendsSeq[0].Contains(nodeID),
			"second batch should not re-query a validator from the first batch")
	}
}
