package aggregator

import (
	"bytes"
	"context"
	"crypto/rand"
	"slices"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
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
	validatorValues := make([]validatorInfo, validatorCount)
	for i := 0; i < validatorCount; i++ {
		localSigner, err := localsigner.New()
		if err != nil {
			panic(err)
		}
		pubKey := localSigner.PublicKey()
		nodeID := ids.GenerateTestNodeID()
		validatorValues[i] = validatorInfo{
			nodeID:            nodeID,
			blsSigner:         localSigner,
			blsPublicKey:      pubKey,
			blsPublicKeyBytes: bls.PublicKeyToUncompressedBytes(pubKey),
		}
	}

	// Sort the validators by public key to construct the NodeValidatorIndexMap
	utils.Sort(validatorValues)

	// Placeholder for results
	validatorSet := make([]*validators.Warp, validatorCount)
	validatorSigners := make([]*localsigner.LocalSigner, validatorCount)
	nodeValidatorIndexMap := make(map[ids.NodeID]int)
	connectedNodes := set.NewSet[ids.NodeID](validatorCount)
	for i, validator := range validatorValues {
		validatorSigners[i] = validator.blsSigner
		validatorSet[i] = &validators.Warp{
			PublicKey:      validator.blsPublicKey,
			PublicKeyBytes: validator.blsPublicKeyBytes,
			Weight:         1,
			NodeIDs:        []ids.NodeID{validator.nodeID},
		}
		nodeValidatorIndexMap[validator.nodeID] = i
		connectedNodes.Add(validator.nodeID)
	}

	return &peers.CanonicalValidators{
		ConnectedWeight: uint64(validatorCount),
		ConnectedNodes:  connectedNodes,
		ValidatorSet: validators.WarpSet{
			Validators:  validatorSet,
			TotalWeight: uint64(validatorCount),
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

// makeConnectedValidatorsWithWeights mirrors makeConnectedValidators but lets each
// validator have its own weight. Validators are returned in canonical (pubkey)
// order, paired with their signers in the same order.
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
		}
	}

	// Canonical order is by uncompressed public-key bytes; preserve the input
	// weight alongside each info by tagging the index before sort.
	type tagged struct {
		info   validatorInfo
		weight uint64
	}
	tagged_ := make([]tagged, count)
	for i := 0; i < count; i++ {
		tagged_[i] = tagged{info: infos[i], weight: weights[i]}
	}
	utils.Sort(infos)
	weightByPubKey := make(map[string]uint64, count)
	for _, t := range tagged_ {
		weightByPubKey[string(t.info.blsPublicKeyBytes)] = t.weight
	}

	validatorSet := make([]*validators.Warp, count)
	validatorSigners := make([]*localsigner.LocalSigner, count)
	nodeValidatorIndexMap := make(map[ids.NodeID]int, count)
	connectedNodes := set.NewSet[ids.NodeID](count)
	var totalWeight uint64
	for i, info := range infos {
		w := weightByPubKey[string(info.blsPublicKeyBytes)]
		validatorSigners[i] = info.blsSigner
		validatorSet[i] = &validators.Warp{
			PublicKey:      info.blsPublicKey,
			PublicKeyBytes: info.blsPublicKeyBytes,
			Weight:         w,
			NodeIDs:        []ids.NodeID{info.nodeID},
		}
		nodeValidatorIndexMap[info.nodeID] = i
		connectedNodes.Add(info.nodeID)
		totalWeight += w
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

// buildFullySignedMessage signs unsignedMessage with every signer, BLS-aggregates
// the result, and populates aggregator.signatureCache so PruneBitSetSignatureToQuorum
// can find each per-signer signature.
func buildFullySignedMessage(
	t *testing.T,
	aggregator *SignatureAggregator,
	unsignedMessage *warp.UnsignedMessage,
	vdrs *peers.CanonicalValidators,
	signers []*localsigner.LocalSigner,
) *warp.Message {
	t.Helper()
	require.Len(t, signers, len(vdrs.ValidatorSet.Validators))

	sigs := make([]*bls.Signature, len(signers))
	bitSet := set.NewBits()
	for i, signer := range signers {
		sig, err := signer.Sign(unsignedMessage.Bytes())
		require.NoError(t, err)
		sigs[i] = sig
		bitSet.Add(i)
		aggregator.signatureCache.Add(
			unsignedMessage.ID(),
			PublicKeyBytes(bls.PublicKeyToUncompressedBytes(signer.PublicKey())),
			SignatureBytes(bls.SignatureToBytes(sig)),
		)
	}
	aggSig, err := bls.AggregateSignatures(sigs)
	require.NoError(t, err)

	signedMsg, err := warp.NewMessage(
		unsignedMessage,
		&warp.BitSetSignature{
			Signers:   bitSet.Bytes(),
			Signature: *(*[bls.SignatureLen]byte)(bls.SignatureToBytes(aggSig)),
		},
	)
	require.NoError(t, err)
	return signedMsg
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

func TestPruneBitSetSignatureToQuorum_PrunesExcessSigners(t *testing.T) {
	const networkID = constants.UnitTestID
	// 5 validators, each weight 1, total 5. At 60% quorum, 3 signers suffice (3/5 = 60%).
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1, 1, 1})

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("payload"))
	require.NoError(t, err)
	signed := buildFullySignedMessage(t, aggregator, unsigned, vdrs, signers)
	bitSetSig := signed.Signature.(*warp.BitSetSignature)

	pruned, err := aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, bitSetSig, vdrs.ValidatorSet, 60,
	)
	require.NoError(t, err)
	require.NotSame(t, bitSetSig, pruned)

	prunedMsg, err := warp.NewMessage(unsigned, pruned)
	require.NoError(t, err)

	// Should retain exactly 3 of 5 signers.
	require.Len(t, signerIndices(t, prunedMsg), 3)

	require.NoError(t, prunedMsg.Signature.Verify(
		unsigned,
		networkID,
		vdrs.ValidatorSet,
		60,
		100,
	))
}

func TestPruneBitSetSignatureToQuorum_GreedyByWeight(t *testing.T) {
	const networkID = constants.UnitTestID
	// Weights chosen so that greedy-by-weight-desc is unambiguous:
	// total = 28; 67% threshold = 18.76 → need >=19.
	// Largest-first: 10 → 10 (insufficient), 10+8 → 18 (insufficient), 10+8+5 → 23 (ok).
	// Any subset of size 2 has weight at most 10+8 = 18 < 19. Greedy selection of
	// size 3 is the unique minimum-cardinality subset that meets quorum.
	weights := []uint64{10, 8, 5, 4, 1}
	vdrs, signers := makeConnectedValidatorsWithWeights(weights)

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("greedy"))
	require.NoError(t, err)
	signed := buildFullySignedMessage(t, aggregator, unsigned, vdrs, signers)
	bitSetSig := signed.Signature.(*warp.BitSetSignature)

	pruned, err := aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, bitSetSig, vdrs.ValidatorSet, 67,
	)
	require.NoError(t, err)

	prunedMsg, err := warp.NewMessage(unsigned, pruned)
	require.NoError(t, err)

	// Pruned set must consist of the 3 heaviest validators by weight (10, 8, 5).
	got := signerIndices(t, prunedMsg)
	require.Len(t, got, 3)

	// Map indices back to weights and confirm they're the top 3.
	gotWeights := make([]uint64, 0, len(got))
	for _, i := range got {
		gotWeights = append(gotWeights, vdrs.ValidatorSet.Validators[i].Weight)
	}
	slices.Sort(gotWeights)
	require.Equal(t, []uint64{5, 8, 10}, gotWeights)

	require.NoError(t, prunedMsg.Signature.Verify(
		unsigned,
		networkID,
		vdrs.ValidatorSet,
		67,
		100,
	))
}

func TestPruneBitSetSignatureToQuorum_SingleHeavySigner(t *testing.T) {
	const networkID = constants.UnitTestID
	// Validator 0 alone exceeds 67% of total weight; pruning must keep only that signer.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{70, 10, 10, 10})

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("heavy"))
	require.NoError(t, err)
	signed := buildFullySignedMessage(t, aggregator, unsigned, vdrs, signers)
	bitSetSig := signed.Signature.(*warp.BitSetSignature)

	pruned, err := aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, bitSetSig, vdrs.ValidatorSet, 67,
	)
	require.NoError(t, err)

	prunedMsg, err := warp.NewMessage(unsigned, pruned)
	require.NoError(t, err)

	got := signerIndices(t, prunedMsg)
	require.Len(t, got, 1)
	require.Equal(t, uint64(70), vdrs.ValidatorSet.Validators[got[0]].Weight)

	require.NoError(t, prunedMsg.Signature.Verify(
		unsigned,
		networkID,
		vdrs.ValidatorSet,
		67,
		100,
	))
}

func TestPruneBitSetSignatureToQuorum_NoPruningWhenAllSignersNeeded(t *testing.T) {
	const networkID = constants.UnitTestID
	// 3 equally-weighted validators at 67% quorum: 2/3 = 66.67% < 67%, so all 3 are needed.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1})

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("tight"))
	require.NoError(t, err)
	signed := buildFullySignedMessage(t, aggregator, unsigned, vdrs, signers)
	bitSetSig := signed.Signature.(*warp.BitSetSignature)

	pruned, err := aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, bitSetSig, vdrs.ValidatorSet, 67,
	)
	require.NoError(t, err)

	// Same bitset returned (no signers could be dropped).
	require.Same(t, bitSetSig, pruned)
}

func TestPruneBitSetSignatureToQuorum_CacheMissReturnsOriginal(t *testing.T) {
	const networkID = constants.UnitTestID
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1, 1, 1})

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("cache-miss"))
	require.NoError(t, err)
	signed := buildFullySignedMessage(t, aggregator, unsigned, vdrs, signers)
	bitSetSig := signed.Signature.(*warp.BitSetSignature)

	// Drop the cached signatures so pruning can't safely re-aggregate.
	freshCache, err := NewSignatureCache(1024)
	require.NoError(t, err)
	aggregator.signatureCache = freshCache

	pruned, err := aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, bitSetSig, vdrs.ValidatorSet, 60,
	)
	require.NoError(t, err)
	require.Same(t, bitSetSig, pruned)
}

func TestPruneBitSetSignatureToQuorum_ExistingSignersBelowQuorumErrors(t *testing.T) {
	const networkID = constants.UnitTestID
	// 5 validators (each weight 1, total 5); only 2 of them are in the signer set
	// — that's 40% of the total, below 67% quorum.
	vdrs, signers := makeConnectedValidatorsWithWeights([]uint64{1, 1, 1, 1, 1})

	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(networkID, ids.GenerateTestID(), []byte("low"))
	require.NoError(t, err)

	// Manually build a signed message with only validators 0 and 1 in the signer set.
	sigs := make([]*bls.Signature, 2)
	bitSet := set.NewBits()
	for i := 0; i < 2; i++ {
		sig, signErr := signers[i].Sign(unsigned.Bytes())
		require.NoError(t, signErr)
		sigs[i] = sig
		bitSet.Add(i)
		aggregator.signatureCache.Add(
			unsigned.ID(),
			PublicKeyBytes(bls.PublicKeyToUncompressedBytes(signers[i].PublicKey())),
			SignatureBytes(bls.SignatureToBytes(sig)),
		)
	}
	aggSig, err := bls.AggregateSignatures(sigs)
	require.NoError(t, err)
	signed, err := warp.NewMessage(
		unsigned,
		&warp.BitSetSignature{
			Signers:   bitSet.Bytes(),
			Signature: *(*[bls.SignatureLen]byte)(bls.SignatureToBytes(aggSig)),
		},
	)
	require.NoError(t, err)

	_, err = aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, signed.Signature.(*warp.BitSetSignature), vdrs.ValidatorSet, 67,
	)
	require.ErrorContains(t, err, "does not meet quorum")
}

func TestPruneBitSetSignatureToQuorum_InvalidBitSetErrors(t *testing.T) {
	aggregator, _, _, _, _ := instantiateAggregator(t)
	unsigned, err := warp.NewUnsignedMessage(constants.UnitTestID, ids.GenerateTestID(), []byte("x"))
	require.NoError(t, err)

	vdrs, _ := makeConnectedValidatorsWithWeights([]uint64{1, 1})
	// Invalid bitset: non-minimal encoding (extra zero padding).
	invalidBitSet := &warp.BitSetSignature{Signers: []byte{0, 0, 0}}

	_, err = aggregator.PruneBitSetSignatureToQuorum(
		logging.NoLog{}, unsigned, invalidBitSet, vdrs.ValidatorSet, 51,
	)
	require.ErrorContains(t, err, "invalid signer bitset")
}
