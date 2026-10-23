// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"math/big"
	"testing"

	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// getAnchor builds a test anchor at the given slot with a root derived from the slot.
func getAnchor(slot uint64) anchor {
	return anchor{slot: slot, root: common.BigToHash(new(big.Int).SetUint64(slot))}
}

// Checks that selectAnchor picks the newest anchor within the 8192-slot window of the target,
// ignoring anchors at or before the target and anchors too far ahead.
func TestSelectAnchor(t *testing.T) {
	const target = uint64(10_000)
	window := uint64(proofs.StateRootsVectorSize)

	tests := []struct {
		name    string
		anchors []anchor
		want    uint64
	}{
		{
			name:    "newest anchor in window is chosen",
			anchors: []anchor{getAnchor(target + 100), getAnchor(target + 500), getAnchor(target + 300)},
			want:    target + 500,
		},
		{
			name:    "anchor at exactly the window edge qualifies",
			anchors: []anchor{getAnchor(target + 1), getAnchor(target + window)},
			want:    target + window,
		},
		{
			name:    "anchor beyond the window is skipped for an older one in window",
			anchors: []anchor{getAnchor(target + 100), getAnchor(target + window + 1)},
			want:    target + 100,
		},
		{
			name:    "anchors at or before the target are ignored",
			anchors: []anchor{getAnchor(target - 1), getAnchor(target), getAnchor(target + 7)},
			want:    target + 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := selectAnchor(tt.anchors, target)
			require.NoError(t, err)
			require.Equal(t, getAnchor(tt.want), got)
		})
	}
}

// Checks that selectAnchor returns a retryable error when no anchor is newer than the target yet.
func TestSelectAnchorNoneNewerYet(t *testing.T) {
	const target = uint64(10_000)

	for _, anchors := range [][]anchor{nil, {getAnchor(target)}, {getAnchor(target - 1), getAnchor(target - 500)}} {
		_, err := selectAnchor(anchors, target)
		require.ErrorContains(t, err, "no confirmed anchor newer than the message slot")
		require.NotErrorIs(t, err, messages.ErrNonRetryable)
	}
}

// Checks that selectAnchor returns a non-retryable error when every newer anchor is beyond the
// 8192-slot window, since no future anchor can cover the target either.
func TestSelectAnchorUnprovable(t *testing.T) {
	const target = uint64(10_000)
	window := uint64(proofs.StateRootsVectorSize)

	anchors := []anchor{getAnchor(target - 5), getAnchor(target + window + 1), getAnchor(target + window + 50)}
	_, err := selectAnchor(anchors, target)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
}

// Checks that containsAnchorAtOrBefore detects an anchor at or before the given slot.
func TestContainsAnchorAtOrBefore(t *testing.T) {
	require.False(t, containsAnchorAtOrBefore(nil, 100))
	require.False(t, containsAnchorAtOrBefore([]anchor{getAnchor(101), getAnchor(200)}, 100))
	require.True(t, containsAnchorAtOrBefore([]anchor{getAnchor(200), getAnchor(100)}, 100))
	require.True(t, containsAnchorAtOrBefore([]anchor{getAnchor(200), getAnchor(99)}, 100))
}
