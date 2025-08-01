package checkpoint

import (
	"container/heap"
	"strconv"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	mock_database "github.com/ava-labs/icm-services/database/mocks"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCommitHeight(t *testing.T) {
	testCases := []struct {
		name              string
		currentMaxHeight  uint64
		commitHeight      uint64
		pendingHeights    *utils.UInt64Heap
		expectedMaxHeight uint64
	}{
		{
			name:              "commit height is the next height",
			currentMaxHeight:  10,
			commitHeight:      11,
			pendingHeights:    &utils.UInt64Heap{},
			expectedMaxHeight: 11,
		},
		{
			name:              "commit height is the next height with pending heights",
			currentMaxHeight:  10,
			commitHeight:      11,
			pendingHeights:    &utils.UInt64Heap{12, 13},
			expectedMaxHeight: 13,
		},
		{
			name:              "commit height is not the next height",
			currentMaxHeight:  10,
			commitHeight:      12,
			pendingHeights:    &utils.UInt64Heap{},
			expectedMaxHeight: 10,
		},
		{
			name:              "commit height is not the next height with pending heights",
			currentMaxHeight:  10,
			commitHeight:      12,
			pendingHeights:    &utils.UInt64Heap{13, 14},
			expectedMaxHeight: 10,
		},
		{
			name:              "commit height is not the next height with next height pending",
			currentMaxHeight:  10,
			commitHeight:      12,
			pendingHeights:    &utils.UInt64Heap{11},
			expectedMaxHeight: 12,
		},
	}
	db := mock_database.NewMockRelayerDatabase(gomock.NewController(t))
	db.EXPECT().Get(gomock.Any(), gomock.Any()).Return([]byte(strconv.FormatUint(0, 10)), nil).AnyTimes()
	for _, test := range testCases {
		id := database.RelayerID{
			ID: common.BytesToHash(crypto.Keccak256([]byte(test.name))),
		}
		cm, err := NewCheckpointManager(logging.NoLog{}, db, nil, id, test.currentMaxHeight)
		require.NoError(t, err)
		heap.Init(test.pendingHeights)
		cm.pendingCommits = test.pendingHeights
		cm.committedHeight = test.currentMaxHeight
		cm.StageCommittedHeight(test.commitHeight)
		require.Equal(t, test.expectedMaxHeight, cm.committedHeight, test.name)
	}
}
