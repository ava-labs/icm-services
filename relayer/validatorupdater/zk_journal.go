// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"fmt"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/accounts/abi"
)

// consensusJournal is the public output of a Boundless consensus proof: a transition from
// PreState to PostState that finalizes FinalizedSlot. It mirrors the Journal struct in
// ZKStateManager.sol, whose transition requires PreState to equal the contract's current state.
type consensusJournal struct {
	PreState      zkadapter.ConsensusState
	PostState     zkadapter.ConsensusState
	FinalizedSlot uint64
}

// journalArguments describes the ABI encoding of ZKStateManager's Journal struct, as decoded by
// abi.decode(journalData, (Journal)) on-chain.
var journalArguments = func() abi.Arguments {
	checkpoint := []abi.ArgumentMarshaling{
		{Name: "epoch", Type: "uint64"},
		{Name: "root", Type: "bytes32"},
	}
	state := []abi.ArgumentMarshaling{
		{Name: "currentJustifiedCheckpoint", Type: "tuple", Components: checkpoint},
		{Name: "finalizedCheckpoint", Type: "tuple", Components: checkpoint},
	}
	journalType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "preState", Type: "tuple", Components: state},
		{Name: "postState", Type: "tuple", Components: state},
		{Name: "finalizedSlot", Type: "uint64"},
	})
	if err != nil {
		panic(fmt.Sprintf("invalid journal ABI type: %v", err))
	}
	return abi.Arguments{{Type: journalType}}
}()

// decodeJournal decodes a proof's ABI-encoded journal.
func decodeJournal(journalData []byte) (*consensusJournal, error) {
	values, err := journalArguments.Unpack(journalData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode consensus journal: %w", err)
	}
	// Copy assigns a single argument to the first field of a struct destination, so the journal
	// is copied into a one-field wrapper rather than into consensusJournal directly.
	var out struct{ Journal consensusJournal }
	if err := journalArguments.Copy(&out, values); err != nil {
		return nil, fmt.Errorf("failed to decode consensus journal: %w", err)
	}
	return &out.Journal, nil
}
