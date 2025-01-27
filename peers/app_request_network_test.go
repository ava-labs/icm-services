// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/peers/avago_mocks"
	validator_mocks "github.com/ava-labs/icm-services/peers/validators/mocks"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCalculateConnectedWeight(t *testing.T) {
	vdr1 := makeValidator(t, 10, 1)
	vdr2 := makeValidator(t, 20, 1)
	vdr3 := makeValidator(t, 30, 2)
	vdrs := []*warp.Validator{&vdr1, &vdr2, &vdr3}
	nodeValidatorIndexMap := map[ids.NodeID]int{
		vdr1.NodeIDs[0]: 0,
		vdr2.NodeIDs[0]: 1,
		vdr3.NodeIDs[0]: 2,
		vdr3.NodeIDs[1]: 2,
	}
	var connectedNodes set.Set[ids.NodeID]
	connectedNodes.Add(vdr1.NodeIDs[0])
	connectedNodes.Add(vdr2.NodeIDs[0])

	// vdr1 and vdr2 are connected, so their weight should be added
	require.Equal(t, 2, connectedNodes.Len())
	connectedWeight := calculateConnectedWeight(vdrs, nodeValidatorIndexMap, connectedNodes)
	require.Equal(t, uint64(30), connectedWeight)

	// Add one of the vdr3's nodeIDs to the connected nodes
	// and confirm that it adds vdr3's weight
	connectedNodes.Add(vdr3.NodeIDs[0])
	require.Equal(t, 3, connectedNodes.Len())
	connectedWeight2 := calculateConnectedWeight(vdrs, nodeValidatorIndexMap, connectedNodes)
	require.Equal(t, uint64(60), connectedWeight2)

	// Add another of vdr3's nodeIDs to the connected nodes
	// and confirm that it's weight isn't double counted
	connectedNodes.Add(vdr3.NodeIDs[1])
	require.Equal(t, 4, connectedNodes.Len())
	connectedWeight3 := calculateConnectedWeight(vdrs, nodeValidatorIndexMap, connectedNodes)
	require.Equal(t, uint64(60), connectedWeight3)
}

func TestConnectToCanonicalValidators(t *testing.T) {
	ctrl := gomock.NewController(t)

	subnetID := ids.GenerateTestID()
	validator1_1 := makeValidator(t, 1, 1)
	validator2_1 := makeValidator(t, 2, 1)
	validator3_2 := makeValidator(t, 3, 2)
	metrics, _ := newAppRequestNetworkMetrics(prometheus.DefaultRegisterer)

	testCases := []struct {
		name                    string
		validators              []*warp.Validator
		connectedNodes          []ids.NodeID
		expectedConnectedWeight uint64
		expectedTotalWeight     uint64
	}{
		{
			name:                    "no connected nodes, one validator",
			validators:              []*warp.Validator{&validator1_1},
			connectedNodes:          []ids.NodeID{},
			expectedConnectedWeight: 0,
			expectedTotalWeight:     1,
		},
		{
			name:       "all validators, missing one nodeID",
			validators: []*warp.Validator{&validator1_1, &validator2_1, &validator3_2},
			connectedNodes: []ids.NodeID{
				validator1_1.NodeIDs[0],
				validator2_1.NodeIDs[0],
				validator3_2.NodeIDs[0],
				validator3_2.NodeIDs[1],
			},
			expectedConnectedWeight: 6,
			expectedTotalWeight:     6,
		},
		{
			name:       "fully connected",
			validators: []*warp.Validator{&validator1_1, &validator2_1, &validator3_2},
			connectedNodes: []ids.NodeID{
				validator1_1.NodeIDs[0],
				validator2_1.NodeIDs[0],
				validator3_2.NodeIDs[0],
				validator3_2.NodeIDs[1],
			},
			expectedConnectedWeight: 6,
			expectedTotalWeight:     6,
		},
		{
			name:       "missing conn to double node validator",
			validators: []*warp.Validator{&validator1_1, &validator2_1, &validator3_2},
			connectedNodes: []ids.NodeID{
				validator1_1.NodeIDs[0],
				validator2_1.NodeIDs[0],
			},
			expectedConnectedWeight: 3,
			expectedTotalWeight:     6,
		},
		{
			name:       "irrelevant nodes",
			validators: []*warp.Validator{&validator1_1, &validator2_1},
			connectedNodes: []ids.NodeID{
				validator1_1.NodeIDs[0],
				validator2_1.NodeIDs[0],
				validator3_2.NodeIDs[0], // this nodeID does not map to the validator
			},
			expectedConnectedWeight: 3,
			expectedTotalWeight:     3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockNetwork := avago_mocks.NewMockNetwork(ctrl)
			mockValidatorClient := validator_mocks.NewMockCanonicalValidatorState(ctrl)
			arNetwork := appRequestNetwork{
				network:         mockNetwork,
				validatorClient: mockValidatorClient,
				metrics:         metrics,
			}
			var totalWeight uint64
			for _, vdr := range testCase.validators {
				totalWeight += vdr.Weight
			}
			mockValidatorClient.EXPECT().GetCurrentCanonicalValidatorSet(subnetID).Return(
				testCase.validators,
				totalWeight,
				nil,
			).Times(1)

			peerInfo := make([]peer.Info, len(testCase.validators))
			for _, node := range testCase.connectedNodes {
				peerInfo = append(peerInfo, peer.Info{
					ID: node,
				})
			}
			mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfo).Times(1)

			ret, err := arNetwork.GetConnectedCanonicalValidators(subnetID)
			require.Equal(t, testCase.expectedConnectedWeight, ret.ConnectedWeight)
			require.Equal(t, testCase.expectedTotalWeight, ret.TotalValidatorWeight)
			require.NoError(t, err)
		})
	}
}

func makeValidator(t *testing.T, weight uint64, numNodeIDs int) warp.Validator {
	sk, err := bls.NewSecretKey()
	require.NoError(t, err)
	pk := bls.PublicFromSecretKey(sk)

	nodeIDs := make([]ids.NodeID, numNodeIDs)
	for i := 0; i < numNodeIDs; i++ {
		nodeIDs[i] = ids.GenerateTestNodeID()
	}
	return warp.Validator{
		PublicKey:      pk,
		PublicKeyBytes: bls.PublicKeyToUncompressedBytes(pk),
		Weight:         weight,
		NodeIDs:        nodeIDs,
	}
}
