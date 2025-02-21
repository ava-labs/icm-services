package aggregator

import (
	"bytes"
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
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers"
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

func instantiateAggregator(t *testing.T) (
	*SignatureAggregator,
	*mocks.MockAppRequestNetwork,
) {
	mockNetwork := mocks.NewMockAppRequestNetwork(gomock.NewController(t))
	if sigAggMetrics == nil {
		sigAggMetrics = metrics.NewSignatureAggregatorMetrics(prometheus.DefaultRegisterer)
	}
	if messageCreator == nil {
		var err error
		messageCreator, err = message.NewCreator(
			logging.NoLog{},
			prometheus.DefaultRegisterer,
			constants.DefaultNetworkCompressionType,
			constants.DefaultNetworkMaximumInboundTimeout,
		)
		require.NoError(t, err)
	}
	aggregator, err := NewSignatureAggregator(
		mockNetwork,
		messageCreator,
		1024,
		sigAggMetrics,
	)
	require.NoError(t, err)
	return aggregator, mockNetwork
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
	for i, validator := range validatorValues {
		validatorSigners[i] = validator.blsSigner
		validatorSet[i] = &warp.Validator{
			PublicKey:      validator.blsPublicKey,
			PublicKeyBytes: validator.blsPublicKeyBytes,
			Weight:         1,
			NodeIDs:        []ids.NodeID{validator.nodeID},
		}
		nodeValidatorIndexMap[validator.nodeID] = i
	}

	return &peers.ConnectedCanonicalValidators{
		ConnectedWeight: uint64(validatorCount),
		ValidatorSet: warp.CanonicalValidatorSet{
			Validators:  validatorSet,
			TotalWeight: uint64(validatorCount),
		},
		NodeValidatorIndexMap: nodeValidatorIndexMap,
	}, validatorSigners
}

func TestCreateSignedMessageFailsWithNoValidators(t *testing.T) {
	aggregator, mockNetwork := instantiateAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockNetwork.EXPECT().GetSubnetID(ids.Empty).Return(ids.Empty, nil)
	mockNetwork.EXPECT().TrackSubnet(ids.Empty)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(ids.Empty).Return(
		&peers.ConnectedCanonicalValidators{
			ConnectedWeight: 0,
			ValidatorSet: warp.CanonicalValidatorSet{
				Validators:  []*warp.Validator{},
				TotalWeight: 0,
			},
		},
		nil,
	)
	_, err = aggregator.CreateSignedMessage(logging.NoLog{}, msg, nil, ids.Empty, 80)
	require.ErrorContains(t, err, "no signatures")
}

func TestCreateSignedMessageFailsWithoutSufficientConnectedStake(t *testing.T) {
	aggregator, mockNetwork := instantiateAggregator(t)
	msg, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{})
	require.NoError(t, err)
	mockNetwork.EXPECT().GetSubnetID(ids.Empty).Return(ids.Empty, nil)
	mockNetwork.EXPECT().TrackSubnet(ids.Empty)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(ids.Empty).Return(
		&peers.ConnectedCanonicalValidators{
			ConnectedWeight: 0,
			ValidatorSet: warp.CanonicalValidatorSet{
				Validators:  []*warp.Validator{},
				TotalWeight: 1,
			},
		},
		nil,
	).AnyTimes()
	_, err = aggregator.CreateSignedMessage(logging.NoLog{}, msg, nil, ids.Empty, 80)
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
	aggregator, mockNetwork := instantiateAggregator(t)

	var (
		connectedValidators, _ = makeConnectedValidators(2)
		requestID              = aggregator.currentRequestID.Load() + 1
	)

	chainID := ids.GenerateTestID()

	msg, err := warp.NewUnsignedMessage(0, chainID, []byte{})
	require.NoError(t, err)

	subnetID := ids.GenerateTestID()
	mockNetwork.EXPECT().GetSubnetID(chainID).Return(
		subnetID,
		nil,
	)

	mockNetwork.EXPECT().TrackSubnet(subnetID)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(subnetID).Return(
		connectedValidators,
		nil,
	)

	appRequests := makeAppRequests(chainID, requestID, connectedValidators)
	for _, appRequest := range appRequests {
		mockNetwork.EXPECT().RegisterAppRequest(appRequest).AnyTimes()
	}

	mockNetwork.EXPECT().RegisterRequestID(
		requestID,
		len(appRequests),
	).Return(
		make(chan message.InboundMessage, len(appRequests)),
	).AnyTimes()

	var nodeIDs set.Set[ids.NodeID]
	for _, appRequest := range appRequests {
		nodeIDs.Add(appRequest.NodeID)
	}
	mockNetwork.EXPECT().Send(
		gomock.Any(),
		nodeIDs,
		subnetID,
		subnets.NoOpAllower,
	).AnyTimes()

	_, err = aggregator.CreateSignedMessage(logging.NoLog{}, msg, nil, subnetID, 80)
	require.ErrorIs(
		t,
		err,
		errNotEnoughSignatures,
	)
}

func TestCreateSignedMessageSucceeds(t *testing.T) {
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

	aggregator, mockNetwork := instantiateAggregator(t)

	subnetID := ids.GenerateTestID()
	mockNetwork.EXPECT().GetSubnetID(chainID).Return(
		subnetID,
		nil,
	)

	mockNetwork.EXPECT().TrackSubnet(subnetID)
	mockNetwork.EXPECT().GetConnectedCanonicalValidators(subnetID).Return(
		connectedValidators,
		nil,
	)

	// prime the signers' responses:

	requestID := aggregator.currentRequestID.Load() + 1

	appRequests := makeAppRequests(chainID, requestID, connectedValidators)
	for _, appRequest := range appRequests {
		mockNetwork.EXPECT().RegisterAppRequest(appRequest).Times(1)
	}

	var nodeIDs set.Set[ids.NodeID]
	responseChan := make(chan message.InboundMessage, len(appRequests))
	for _, appRequest := range appRequests {
		nodeIDs.Add(appRequest.NodeID)
		validatorSigner := validatorSigners[connectedValidators.NodeValidatorIndexMap[appRequest.NodeID]]

		signature, err := validatorSigner.Sign(msg.Bytes())
		require.NoError(t, err)
		responseBytes, err := proto.Marshal(
			&sdk.SignatureResponse{
				Signature: bls.SignatureToBytes(
					signature,
				),
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
		len(appRequests),
	).Return(responseChan).Times(1)

	mockNetwork.EXPECT().Send(
		gomock.Any(),
		nodeIDs,
		subnetID,
		subnets.NoOpAllower,
	).Times(1).Return(nodeIDs)

	// aggregate the signatures:
	var quorumPercentage uint64 = 80
	signedMessage, err := aggregator.CreateSignedMessage(
		logging.NoLog{},
		msg,
		nil,
		subnetID,
		quorumPercentage,
	)
	require.NoError(t, err)

	verifyErr := signedMessage.Signature.Verify(
		msg,
		networkID,
		connectedValidators.ValidatorSet,
		quorumPercentage,
		100,
	)
	require.NoError(t, verifyErr)
}

func TestUnmarshalResponse(t *testing.T) {
	aggregator, _ := instantiateAggregator(t)

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
