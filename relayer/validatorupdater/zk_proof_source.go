// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import "context"

// ConsensusProof is one SP1 Helios light client update proof. PublicValues is the ABI-encoded
// ProofOutputs the proof commits to, and NewHead is the slot it finalizes.
type ConsensusProof struct {
	NewHead      uint64
	PublicValues []byte
	Proof        []byte
}

// ProofSource provides SP1 Helios light client update proofs to the consensus updater.
//
// TODO: Implement a ProofSource. The Boundless subgraph that previously served consensus proofs
// does not index SP1 Helios proofs. Issue: https://github.com/ava-labs/icm-services/issues/1626
type ProofSource interface {
	// ConsensusProofsAfterSlot returns up to limit proofs with NewHead > afterSlot, in ascending
	// order. Ascending order matters because transitions are applied sequentially by the caller.
	// Each proof's previous values must match the contract's current state.
	ConsensusProofsAfterSlot(ctx context.Context, afterSlot uint64, limit int) ([]*ConsensusProof, error)
}
