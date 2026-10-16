// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package checkpoint

// blockRange is an inclusive range of block heights.
type blockRange struct {
	from uint64
	to   uint64
}

// blockRangeHeap is a min-heap of block ranges ordered by their first height.
// Adapted from https://pkg.go.dev/container/heap#example-package-IntHeap
type blockRangeHeap []blockRange

func (h blockRangeHeap) Len() int           { return len(h) }
func (h blockRangeHeap) Less(i, j int) bool { return h[i].from < h[j].from }
func (h blockRangeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h blockRangeHeap) Peek() blockRange   { return h[0] }

func (h *blockRangeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(blockRange))
}

func (h *blockRangeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
