// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/ava-labs/libevm/common"
)

// EventFilter selects the source chain logs that carry a message protocol's messages, following
// the semantics of ethereum.FilterQuery: a log matches if it was emitted by one of [Addresses]
// (any address, if empty) and its topics match [Topics].
// Filters should constrain the emitting address whenever possible: topics can be forged by any
// contract, so a filter that only constrains topics can match logs that were not emitted by the
// message protocol.
type EventFilter struct {
	// Addresses are the contract addresses that emit the protocol's message logs. For protocols
	// that send their messages through the Warp precompile, this is the precompile's address
	// rather than the protocol's own.
	Addresses []common.Address
	// Topics constrain the topics of matching logs, in the format expected by
	// ethereum.FilterQuery.Topics.
	Topics [][]common.Hash
}

// IsEmpty reports whether the filter constrains neither the emitting address nor the log topics.
// Message protocols whose messages are not read from source chain logs return an empty filter.
func (f EventFilter) IsEmpty() bool {
	return len(f.Addresses) == 0 && len(f.Topics) == 0
}
