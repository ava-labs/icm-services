// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package validatorupdater

import (
	"errors"
	"fmt"
	"math/big"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
)

// proofOutputs mirrors the ProofOutputs struct in ZKStateManager.sol: the public values an
// SP1 Helios light client proof commits to.
type proofOutputs struct {
	PrevHeader            [32]byte
	PrevHead              *big.Int
	PrevSyncCommitteeHash [32]byte
	NewHead               *big.Int
	NewHeader             [32]byte
	ExecutionStateRoot    [32]byte
	ExecutionBlockNumber  *big.Int
	SyncCommitteeHash     [32]byte
	NextSyncCommitteeHash [32]byte
	StorageSlots          []storageSlot
}

// storageSlot mirrors the StorageSlot struct in ZKStateManager.sol.
type storageSlot struct {
	Key             [32]byte
	Value           [32]byte
	ContractAddress common.Address
}

// consensusTransition is the part of a proof's public values the updater selects proofs by: a
// transition from PreState to PostState, finalizing PostState.Head. ZKStateManager's transition
// requires PreState to equal the contract's current state.
type consensusTransition struct {
	PreState  zkadapter.ConsensusState
	PostState zkadapter.ConsensusState
}

// proofOutputsArguments describes the ABI encoding of ZKStateManager's ProofOutputs struct, as
// decoded by abi.decode(publicValues, (ProofOutputs)) on-chain.
var proofOutputsArguments = func() abi.Arguments {
	outputsType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "prevHeader", Type: "bytes32"},
		{Name: "prevHead", Type: "uint256"},
		{Name: "prevSyncCommitteeHash", Type: "bytes32"},
		{Name: "newHead", Type: "uint256"},
		{Name: "newHeader", Type: "bytes32"},
		{Name: "executionStateRoot", Type: "bytes32"},
		{Name: "executionBlockNumber", Type: "uint256"},
		{Name: "syncCommitteeHash", Type: "bytes32"},
		{Name: "nextSyncCommitteeHash", Type: "bytes32"},
		{Name: "storageSlots", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "key", Type: "bytes32"},
			{Name: "value", Type: "bytes32"},
			{Name: "contractAddress", Type: "address"},
		}},
	})
	if err != nil {
		panic(fmt.Sprintf("invalid proof outputs ABI type: %v", err))
	}
	return abi.Arguments{{Type: outputsType}}
}()

// decodePublicValues decodes a proof's ABI-encoded public values into the transition it proves.
// The post-state's sync committee is the one the proof reports for the new head's period, which
// the contract stores unless it already holds one for that period.
func decodePublicValues(publicValues []byte) (*consensusTransition, error) {
	values, err := proofOutputsArguments.Unpack(publicValues)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public values: %w", err)
	}
	// Copy assigns a single argument to the first field of a struct destination, so the outputs
	// are copied into a one-field wrapper rather than into proofOutputs directly.
	var out struct{ Outputs proofOutputs }
	if err := proofOutputsArguments.Copy(&out, values); err != nil {
		return nil, fmt.Errorf("failed to decode public values: %w", err)
	}
	outputs := out.Outputs
	// The contract casts both heads to uint64, reverting if either does not fit.
	if !outputs.PrevHead.IsUint64() || !outputs.NewHead.IsUint64() {
		return nil, errors.New("failed to decode public values: head does not fit in uint64")
	}
	return &consensusTransition{
		PreState: zkadapter.ConsensusState{
			Head:              outputs.PrevHead.Uint64(),
			Header:            outputs.PrevHeader,
			SyncCommitteeHash: outputs.PrevSyncCommitteeHash,
		},
		PostState: zkadapter.ConsensusState{
			Head:              outputs.NewHead.Uint64(),
			Header:            outputs.NewHeader,
			SyncCommitteeHash: outputs.SyncCommitteeHash,
		},
	}, nil
}
