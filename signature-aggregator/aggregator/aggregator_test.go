package aggregator

import (
	"bytes"
	"context"
	"slices"
	"testing"

	"crypto/rand"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
	avago_mocks "github.com/ava-labs/icm-services/peers/avago_mocks"
	"github.com/ava-labs/icm-services/peers/mocks"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"
)

var (
	sigAggMetrics  *metrics.SignatureAggregatorMetrics
	messageCreator message.Creator
)

const (
	maxAppRequestRetries uint32 = 8 // Used to calculate expected RequestID values
)

func instantiateAggregator(t *testing.T) (
	*SignatureAggregator,
	*mocks.MockAppRequestNetwork,
	*avago_mocks.MockClient,
) {
	mockController := gomock.NewController(t)
	mockNetwork := mocks.NewMockAppRequestNetwork(mockController)
	if sigAggMetrics == nil {
		sigAggMetrics = metrics.NewSignatureAggregatorMetrics(prometheus.DefaultRegisterer)
	}
	if messageCreator == nil {
		var err error
		messageCreator, err = message.NewCreator(
			prometheus.DefaultRegisterer,
			constants.DefaultNetworkCompressionType,
			constants.DefaultNetworkMaximumInboundTimeout,
		)
		require.NoError(t, err)
	}
	mockPClient := avago_mocks.NewMockClient(mockController)
	aggregator, err := NewSignatureAggregator(
		mockNetwork,
		messageCreator,
		1024,
		sigAggMetrics,
		mockPClient,
		nil,
	)
	require.NoError(t, err)
	return aggregator, mockNetwork, mockPClient
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

func makeConnectedValidators(validatorCount int) (*peers.ConnectedCanonicalValidators, []*localsigner.LocalSigner) {
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
	validatorSet := make([]*warp.Validator, validatorCount)
	validatorSigners := make([]*localsigner.LocalSigner, validatorCount)
	nodeValidatorIndexMap := make(map[ids.NodeID]int)
	connectedNodes := set.NewSet[ids.NodeID](validatorCount)
	for i, validator := range validatorValues {
		validatorSigners[i] = validator.blsSigner
		validatorSet[i] = &warp.Validator{
			PublicKey:      validator.blsPublicKey,
			PublicKeyBytes: validator.blsPublicKeyBytes,
			Weight:         1,
			NodeIDs:        []ids.NodeID{validator.nodeID},
		}
		nodeValidatorIndexMap[validator.nodeID] = i
		connectedNodes.Add(validator.nodeID)
	}

	return &peers.ConnectedCanonicalValidators{
		ConnectedWeight: uint64(validatorCount),
		ConnectedNodes:  connectedNodes,
		ValidatorSet: warp.CanonicalValidatorSet{
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
			aggregator, _, _ := instantiateAggregator(t)
			signedMsg, err := aggregator.CreateSignedMessage(
				context.Background(),
				logging.NoLog{},
				nil,
				nil,
				ids.Empty,
				tc.requiredQuorumPercentage,
				tc.quorumPercentageBuffer,
				false,
			)
			require.Nil(t, signedMsg)
			require.ErrorIs(t, err, errInvalidQuorumPercentage)
		})
	}
}

func TestCreateSignedMessageFailsWithNoValidators(t *testing.T) {
	aggregator, mockNetwork, _ := instantiateAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockNetwork.EXPECT().GetSubnetID(gomock.Any(), ids.Empty).Return(ids.Empty, nil)
	mockNetwork.EXPECT().TrackSubnet(ids.Empty)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(gomock.Any(), ids.Empty, false).Return(
		&peers.ConnectedCanonicalValidators{
			ConnectedWeight: 0,
			ValidatorSet: warp.CanonicalValidatorSet{
				Validators:  []*warp.Validator{},
				TotalWeight: 0,
			},
		},
		nil,
	)
	_, err = aggregator.CreateSignedMessage(context.Background(), logging.NoLog{}, msg, nil, ids.Empty, 80, 0, false)
	require.ErrorContains(t, err, "no signatures")
}

func TestCreateSignedMessageFailsWithoutSufficientConnectedStake(t *testing.T) {
	aggregator, mockNetwork, _ := instantiateAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockNetwork.EXPECT().GetSubnetID(gomock.Any(), ids.Empty).Return(ids.Empty, nil)
	mockNetwork.EXPECT().TrackSubnet(ids.Empty)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(gomock.Any(), ids.Empty, false).Return(
		&peers.ConnectedCanonicalValidators{
			ConnectedWeight: 0,
			ValidatorSet: warp.CanonicalValidatorSet{
				Validators:  []*warp.Validator{},
				TotalWeight: 1,
			},
		},
		nil,
	).AnyTimes()
	_, err = aggregator.CreateSignedMessage(context.Background(), logging.NoLog{}, msg, nil, ids.Empty, 80, 0, false)
	require.ErrorContains(
		t,
		err,
		"failed to connect to a threshold of stake",
	)
}

func makeAppRequests(
	chainID ids.ID,
	requestID uint32,
	connectedValidators *peers.ConnectedCanonicalValidators,
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
	aggregator, mockNetwork, mockPClient := instantiateAggregator(t)

	var (
		connectedValidators, _ = makeConnectedValidators(2)
		requestID              = aggregator.currentRequestID.Load() + 2
	)

	chainID := ids.GenerateTestID()

	msg, err := warp.NewUnsignedMessage(0, chainID, []byte{})
	require.NoError(t, err)

	subnetID := ids.GenerateTestID()
	mockNetwork.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(
		subnetID,
		nil,
	)

	mockNetwork.EXPECT().TrackSubnet(subnetID)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(gomock.Any(), subnetID, false).Return(
		connectedValidators,
		nil,
	)

	appRequests := makeAppRequests(chainID, requestID, connectedValidators)
	var nodeIDs set.Set[ids.NodeID]
	for _, appRequest := range appRequests {
		nodeIDs.Add(appRequest.NodeID)
		// Expect at most one call to RegisterAppRequest per node per retry for up to [maxAppRequestRetries] retries
		for i := uint32(0); i < maxAppRequestRetries; i++ {
			appRequestCopy := appRequest
			appRequestCopy.RequestID = appRequest.RequestID + i*2
			mockNetwork.EXPECT().RegisterAppRequest(appRequestCopy).MaxTimes(1)
		}
	}

	for i := uint32(0); i < maxAppRequestRetries; i++ {
		mockNetwork.EXPECT().RegisterRequestID(
			requestID+i*2,
			nodeIDs,
		).Return(
			make(chan message.InboundMessage, len(appRequests)),
		).MaxTimes(1)
	}

	mockNetwork.EXPECT().Send(
		gomock.Any(),
		nodeIDs,
		subnetID,
		subnets.NoOpAllower,
	).AnyTimes()

	mockPClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
		platformvm.GetSubnetClientResponse{},
		nil,
	).Times(1)

	_, err = aggregator.CreateSignedMessage(context.Background(), logging.NoLog{}, msg, nil, subnetID, 80, 0, false)
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

			aggregator, mockNetwork, mockPClient := instantiateAggregator(t)

			subnetID := ids.GenerateTestID()
			mockNetwork.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(
				subnetID,
				nil,
			)

			mockNetwork.EXPECT().TrackSubnet(subnetID)
			mockNetwork.EXPECT().GetConnectedCanonicalValidators(gomock.Any(), subnetID, false).Return(
				connectedValidators,
				nil,
			)

			mockPClient.EXPECT().GetSubnet(gomock.Any(), subnetID).Return(
				platformvm.GetSubnetClientResponse{},
				nil,
			).Times(1)

			// prime the signers' responses:

			requestID := aggregator.currentRequestID.Load() + 2

			appRequests := makeAppRequests(chainID, requestID, connectedValidators)
			for _, appRequest := range appRequests {
				mockNetwork.EXPECT().RegisterAppRequest(appRequest).Times(1)
			}

			var nodeIDs set.Set[ids.NodeID]
			responseChan := make(chan message.InboundMessage, len(appRequests))
			for i, appRequest := range appRequests {
				nodeIDs.Add(appRequest.NodeID)
				validatorSigner := validatorSigners[connectedValidators.NodeValidatorIndexMap[appRequest.NodeID]]

				// Simulate 1 of 5 validators responding with an invalid signature
				var signatureBytes []byte
				if i == len(appRequests)-1 {
					signatureBytes = make([]byte, 0)
				} else {
					signature, err := validatorSigner.Sign(msg.Bytes())
					require.NoError(t, err)
					signatureBytes = bls.SignatureToBytes(signature)
				}

				responseBytes, err := proto.Marshal(
					&sdk.SignatureResponse{
						Signature: signatureBytes,
					},
				)
				require.NoError(t, err)
				responseChan <- message.InboundAppResponse(
					chainID,
					requestID,
					responseBytes,
					appRequest.NodeID,
				)
			}
			mockNetwork.EXPECT().RegisterRequestID(
				requestID,
				nodeIDs,
			).Return(responseChan).Times(1)

			mockNetwork.EXPECT().Send(
				gomock.Any(),
				nodeIDs,
				subnetID,
				subnets.NoOpAllower,
			).Times(1).Return(nodeIDs)

			// aggregate the signatures:
			// This should still succeed because we have 4 out of 5 valid signatures,
			// even though we're not able to get the quorum percentage buffer.
			signedMessage, err := aggregator.CreateSignedMessage(
				context.Background(),
				logging.NoLog{},
				msg,
				nil,
				subnetID,
				tc.requiredQuorumPercentage,
				tc.quorumPercentageBuffer,
				false,
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
	aggregator, _, _ := instantiateAggregator(t)

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
		connected    *peers.ConnectedCanonicalValidators
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
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
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
						{NodeIDs: []ids.NodeID{nodeID1, nodeID2}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name:         "no L1 validators",
			l1Validators: []platformvm.ClientPermissionlessValidator{},
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{
						{NodeIDs: []ids.NodeID{nodeID3}},
					},
				},
			},
			excludedIdx: []int{},
		},
		{
			name:         "empty validator set",
			l1Validators: []platformvm.ClientPermissionlessValidator{},
			connected: &peers.ConnectedCanonicalValidators{
				ValidatorSet: warp.CanonicalValidatorSet{
					Validators: []*warp.Validator{},
				},
			},
			excludedIdx: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aggregator, _, mockPClient := instantiateAggregator(t)
			ctx := context.Background()
			log := logging.NoLog{}
			signingSubnet := ids.GenerateTestID()

			mockPClient.EXPECT().
				GetCurrentValidators(gomock.Any(), signingSubnet, gomock.Nil(), gomock.Any()).
				Return(tc.l1Validators, nil)

			excluded, err := aggregator.getExcludedValidators(ctx, log, signingSubnet, tc.connected, true)
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
