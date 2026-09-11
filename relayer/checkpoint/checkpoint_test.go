package checkpoint

import (
	"container/heap"
	"strconv"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	mock_database "github.com/ava-labs/icm-services/database/mocks"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestStageCommittedHeights(t *testing.T) {
	testCases := []struct {
		name                 string
		currentMaxHeight     uint64
		commitRange          blockRange
		pendingRanges        blockRangeHeap
		expectedMaxHeight    uint64
		expectedPendingCount int
	}{
		{
			name:              "commit height is the next height",
			currentMaxHeight:  10,
			commitRange:       blockRange{11, 11},
			expectedMaxHeight: 11,
		},
		{
			name:              "commit height is the next height with pending heights",
			currentMaxHeight:  10,
			commitRange:       blockRange{11, 11},
			pendingRanges:     blockRangeHeap{{12, 12}, {13, 13}},
			expectedMaxHeight: 13,
		},
		{
			name:                 "commit height is not the next height",
			currentMaxHeight:     10,
			commitRange:          blockRange{12, 12},
			expectedMaxHeight:    10,
			expectedPendingCount: 1,
		},
		{
			name:                 "commit height is not the next height with pending heights",
			currentMaxHeight:     10,
			commitRange:          blockRange{12, 12},
			pendingRanges:        blockRangeHeap{{13, 13}, {14, 14}},
			expectedMaxHeight:    10,
			expectedPendingCount: 3,
		},
		{
			name:              "commit height is not the next height with next height pending",
			currentMaxHeight:  10,
			commitRange:       blockRange{12, 12},
			pendingRanges:     blockRangeHeap{{11, 11}},
			expectedMaxHeight: 12,
		},
		{
			name:              "commit range starting at the next height",
			currentMaxHeight:  10,
			commitRange:       blockRange{11, 20},
			expectedMaxHeight: 20,
		},
		{
			name:              "commit range fills the gap before pending ranges",
			currentMaxHeight:  10,
			commitRange:       blockRange{11, 14},
			pendingRanges:     blockRangeHeap{{15, 20}, {21, 21}, {30, 35}},
			expectedMaxHeight: 21,
			// {30, 35} is still waiting for 22-29.
			expectedPendingCount: 1,
		},
		{
			name:              "commit range already committed",
			currentMaxHeight:  10,
			commitRange:       blockRange{5, 10},
			expectedMaxHeight: 10,
		},
		{
			name:              "commit range straddles the committed height",
			currentMaxHeight:  10,
			commitRange:       blockRange{5, 15},
			expectedMaxHeight: 15,
		},
		{
			name:              "pending ranges overlap each other",
			currentMaxHeight:  10,
			commitRange:       blockRange{11, 13},
			pendingRanges:     blockRangeHeap{{12, 15}, {13, 13}, {14, 16}},
			expectedMaxHeight: 16,
		},
		{
			name:              "inverted range is ignored",
			currentMaxHeight:  10,
			commitRange:       blockRange{12, 11},
			expectedMaxHeight: 10,
		},
	}
	db := mock_database.NewMockRelayerDatabase(gomock.NewController(t))
	db.EXPECT().Get(gomock.Any(), gomock.Any()).Return([]byte(strconv.FormatUint(0, 10)), nil).AnyTimes()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			id := database.RelayerID{
				ID: common.BytesToHash(crypto.Keccak256([]byte(test.name))),
			}
			registry := prometheus.NewRegistry()
			metrics := NewCheckpointManagerMetrics(registry)
			cm, err := NewCheckpointManager(logging.NoLog{}, metrics, db, nil, id, test.currentMaxHeight)
			require.NoError(t, err)
			pending := test.pendingRanges
			if pending == nil {
				pending = blockRangeHeap{}
			}
			heap.Init(&pending)
			cm.pendingCommits = &pending
			cm.committedHeight = test.currentMaxHeight
			cm.StageCommittedHeights(test.commitRange.from, test.commitRange.to)
			require.Equal(t, test.expectedMaxHeight, cm.committedHeight)
			require.Equal(t, test.expectedPendingCount, cm.pendingCommits.Len())
		})
	}
}
