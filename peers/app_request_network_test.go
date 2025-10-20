// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peers

import (
	"context"
	"sync"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/peer"
	snowVdrs "github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/crypto/bls/signer/localsigner"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/icm-services/cache"
	"github.com/ava-labs/icm-services/peers/avago_mocks"
	validator_mocks "github.com/ava-labs/icm-services/peers/validators/mocks"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var metrics = newAppRequestNetworkMetrics(prometheus.DefaultRegisterer)

func TestCalculateConnectedWeight(t *testing.T) {
	vdr1 := makeValidator(t, 10, 1)
	vdr2 := makeValidator(t, 20, 1)
	vdr3 := makeValidator(t, 30, 2)
	vdrs := []*snowVdrs.Warp{&vdr1, &vdr2, &vdr3}
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

	testCases := []struct {
		name                    string
		validators              []*snowVdrs.Warp
		connectedNodes          []ids.NodeID
		expectedConnectedWeight uint64
		expectedTotalWeight     uint64
	}{
		{
			name:                    "no connected nodes, one validator",
			validators:              []*snowVdrs.Warp{&validator1_1},
			connectedNodes:          []ids.NodeID{},
			expectedConnectedWeight: 0,
			expectedTotalWeight:     1,
		},
		{
			name:       "all validators, missing one nodeID",
			validators: []*snowVdrs.Warp{&validator1_1, &validator2_1, &validator3_2},
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
			validators: []*snowVdrs.Warp{&validator1_1, &validator2_1, &validator3_2},
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
			validators: []*snowVdrs.Warp{&validator1_1, &validator2_1, &validator3_2},
			connectedNodes: []ids.NodeID{
				validator1_1.NodeIDs[0],
				validator2_1.NodeIDs[0],
			},
			expectedConnectedWeight: 3,
			expectedTotalWeight:     6,
		},
		{
			name:       "irrelevant nodes",
			validators: []*snowVdrs.Warp{&validator1_1, &validator2_1},
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
			vdrsCache := cache.NewTTLCache[ids.ID, snowVdrs.WarpSet](canonicalValidatorSetCacheTTL)
			arNetwork := appRequestNetwork{
				network:                    mockNetwork,
				validatorClient:            mockValidatorClient,
				metrics:                    metrics,
				logger:                     logging.NoLog{},
				canonicalValidatorSetCache: vdrsCache,
				epochedValidatorSetCache:   cache.NewFIFOCache[uint64, map[ids.ID]snowVdrs.WarpSet](100),
			}
			var totalWeight uint64
			for _, vdr := range testCase.validators {
				totalWeight += vdr.Weight
			}
			mockValidatorClient.EXPECT().GetProposedValidators(
				gomock.Any(), subnetID).Return(
				snowVdrs.WarpSet{
					Validators:  testCase.validators,
					TotalWeight: testCase.expectedTotalWeight,
				},
				nil,
			).Times(1)

			peerInfo := make([]peer.Info, len(testCase.validators))
			for _, node := range testCase.connectedNodes {
				peerInfo = append(peerInfo, peer.Info{
					ID: node,
				})
			}
			mockNetwork.EXPECT().PeerInfo(gomock.Any()).Return(peerInfo).Times(1)

			ret, err := arNetwork.GetCanonicalValidators(context.Background(), subnetID, false, uint64(pchainapi.ProposedHeight))
			require.Equal(t, testCase.expectedConnectedWeight, ret.ConnectedWeight)
			require.Equal(t, testCase.expectedTotalWeight, ret.ValidatorSet.TotalWeight)
			require.NoError(t, err)
		})
	}
}

func TestTrackSubnets(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockNetwork := avago_mocks.NewMockNetwork(ctrl)
	mockValidatorClient := validator_mocks.NewMockCanonicalValidatorState(ctrl)
	arNetwork := appRequestNetwork{
		network:            mockNetwork,
		logger:             logging.NoLog{},
		validatorClient:    mockValidatorClient,
		metrics:            metrics,
		manager:            snowVdrs.NewManager(),
		lruSubnets:         linked.NewHashmapWithSize[ids.ID, interface{}](maxNumSubnets),
		validatorSetLock:   new(sync.Mutex),
		trackedSubnetsLock: new(sync.RWMutex),
	}
	require.Zero(t, arNetwork.trackedSubnets.Len())
	require.Zero(t, arNetwork.lruSubnets.Len())
	mockValidatorClient.EXPECT().GetProposedValidators(
		gomock.Any(), gomock.Any(),
	).Return(snowVdrs.WarpSet{}, nil).AnyTimes()
	for range maxNumSubnets {
		arNetwork.TrackSubnet(t.Context(), ids.GenerateTestID())
	}
	require.Equal(t, arNetwork.trackedSubnets.Len(), arNetwork.lruSubnets.Len())
	require.Equal(t, arNetwork.trackedSubnets.Len(), maxNumSubnets)

	// Add one more subnet, which should evict the oldest subnet
	newSubnetID := ids.GenerateTestID()
	oldestSubnetID, _, ok := arNetwork.lruSubnets.Oldest()
	require.True(t, ok)

	arNetwork.TrackSubnet(t.Context(), newSubnetID)
	require.Equal(t, maxNumSubnets, arNetwork.trackedSubnets.Len())
	require.Equal(t, maxNumSubnets, arNetwork.lruSubnets.Len())
	require.False(t, arNetwork.trackedSubnets.Contains(oldestSubnetID))
	_, has := arNetwork.lruSubnets.Get(oldestSubnetID)
	require.False(t, has)

	it := arNetwork.lruSubnets.NewIterator()
	require.NotNil(t, it)
	for range maxNumSubnets {
		require.True(t, it.Next())
		subnetID := it.Key()
		// confirm that they are still in sync
		require.True(t, arNetwork.trackedSubnets.Contains(subnetID))
	}
	// confirm that the iterator is exhausted after maxNumSubnets iterations
	require.False(t, it.Next())
}

func makeValidator(t *testing.T, weight uint64, numNodeIDs int) snowVdrs.Warp {
	localSigner, err := localsigner.New()
	require.NoError(t, err)
	pk := localSigner.PublicKey()

	nodeIDs := make([]ids.NodeID, numNodeIDs)
	for i := 0; i < numNodeIDs; i++ {
		nodeIDs[i] = ids.GenerateTestNodeID()
	}
	return snowVdrs.Warp{
		PublicKey:      pk,
		PublicKeyBytes: bls.PublicKeyToUncompressedBytes(pk),
		Weight:         weight,
		NodeIDs:        nodeIDs,
	}
}
