// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package zk

import (
	"context"
	"errors"
	"fmt"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
)

// destinationBlockRange is the number of destination chain blocks covered by each
// ConfirmedBeaconBlock log query. Confirmed anchor blocks are emitted on the destination
// chain via events, so this is a destination-side scan (Avalanche), not source-chain (Ethereum).
const destinationBlockRange = 10_000

// anchor is a beacon slot the ZKAdapter has confirmed, with the beacon block root it stores for
// that slot. The adapter emits a ConfirmedBeaconBlock event for each one.
type anchor struct {
	slot uint64
	root common.Hash
}

// findAnchor returns the confirmed anchor to prove a message at targetSlot against.
//
// Anchors are read from the ZKAdapter's ConfirmedBeaconBlock events, scanning the destination
// chain backwards from the latest block in ranges of destinationBlockRange blocks. The scan
// stops as soon as it reaches an anchor at or before targetSlot.
func findAnchor(
	ctx context.Context,
	client evm.Client,
	adapter *zkadapter.ZKAdapter,
	targetSlot uint64,
) (anchor, error) {
	// Get latest destination chain block header.
	header, err := client.HeaderByNumber(ctx, nil) // nil is latest block
	if err != nil {
		return anchor{}, fmt.Errorf("failed to get latest destination header: %w", err)
	}

	var anchors []anchor
	// Scan backwards on destination in chunks of destinationBlockRange blocks until anchor block is found.
	for end := header.Number.Uint64(); ; end -= destinationBlockRange {
		start := uint64(0)
		if end >= destinationBlockRange {
			start = end - destinationBlockRange + 1
		}

		found, err := fetchConfirmedAnchors(ctx, adapter, start, end)
		if err != nil {
			return anchor{}, err
		}
		anchors = append(anchors, found...)

		selected, err := selectAnchor(anchors, targetSlot)
		if err == nil {
			return selected, nil
		}
		// Once an anchor at or before the target has been seen, or the scan has reached
		// genesis, no older block range can work.
		if containsAnchorAtOrBefore(found, targetSlot) || start == 0 {
			return anchor{}, err
		}
	}
}

// fetchConfirmedAnchors returns the anchors confirmed in destination blocks in range start to end.
func fetchConfirmedAnchors(
	ctx context.Context,
	adapter *zkadapter.ZKAdapter,
	start, end uint64,
) ([]anchor, error) {
	// Filter the ConfirmedBeaconBlock events from the adapter contract in the given block range.
	iter, err := adapter.FilterConfirmedBeaconBlock(
		&bind.FilterOpts{Context: ctx, Start: start, End: &end}, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to filter ConfirmedBeaconBlock events: %w", err)
	}
	defer iter.Close()

	// Get the anchors from the events.
	var anchors []anchor
	for iter.Next() {
		anchors = append(anchors, anchor{slot: iter.Event.Slot, root: iter.Event.Root})
	}
	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("failed to iterate ConfirmedBeaconBlock events: %w", err)
	}
	return anchors, nil
}

func containsAnchorAtOrBefore(anchors []anchor, slot uint64) bool {
	for _, a := range anchors {
		if a.slot <= slot {
			return true
		}
	}
	return false
}

// selectAnchor selects the anchor to prove a message at targetSlot against. An anchor
// is selected if it is newer than the target slot and within StateRootsVectorSize slots of it,
// since that is how far back its state_roots history reaches. There can be multiple valid
// anchors so we just pick the newest one on the destination chain.
func selectAnchor(anchors []anchor, targetSlot uint64) (anchor, error) {
	var best *anchor
	var doesSuitableAnchorExist bool // for error logging
	// Scans list of anchors for the newest one that is within the state_roots window of the target slot.
	for i := range anchors {
		a := anchors[i]

		// An anchor can only prove slots historcially before it (from state_roots vector).
		if a.slot <= targetSlot {
			continue
		}
		doesSuitableAnchorExist = true

		// The target must still be within the anchor's state_roots history.
		if a.slot-targetSlot > proofs.StateRootsVectorSize {
			continue
		}

		// Select the latest anchor.
		if best == nil || a.slot > best.slot {
			best = &anchors[i]
		}
	}
	// Returns suitable anchor.
	if best != nil {
		return *best, nil
	}
	// No suitable anchors exist.
	if !doesSuitableAnchorExist {
		return anchor{}, errors.New("no confirmed anchor newer than the message slot yet")
	}
	// Newer suitable anchors do exist but all are too far ahead.
	return anchor{}, fmt.Errorf(
		"%w: target slot %d is outside every confirmed anchor's %d-slot window",
		messages.ErrNonRetryable, targetSlot, proofs.StateRootsVectorSize,
	)
}
