// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package beacon

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/fulu"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
	"github.com/holiman/uint256"
)

// Constants for building the lite beacon state and its Merkle subtrees,
// derived from the Fulu BeaconState and ExecutionPayloadHeader definitions:
// https://github.com/ethereum/consensus-specs/blob/master/specs/fulu/beacon-chain.md#beaconstate
const (
	// numStateFieldLeaves pads the BeaconState's 38 fields to the next
	// power of 2: 2^6 = 64.
	numStateFieldLeaves = 64
	// numExecHeaderLeaves pads the ExecutionPayloadHeader's 17 fields to
	// 2^5 = 32.
	numExecHeaderLeaves = 32

	// Leaf positions in the BeaconState of the two fields we build
	// subtrees for.
	stateRootsLeafIndex        = 6
	execPayloadHeaderLeafIndex = 24
)

// SSZ list limits from the Beacon consensus spec's presets, including padding to the next power
// of two for merkleization.
const (
	historicalRootsLimit       = 1 << 24 // HISTORICAL_ROOTS_LIMIT
	eth1DataVotesLimit         = 2048    // EPOCHS_PER_ETH1_VOTING_PERIOD * SLOTS_PER_EPOCH
	validatorRegistryLimit     = 1 << 40 // VALIDATOR_REGISTRY_LIMIT
	historicalSummariesLimit   = 1 << 24 // HISTORICAL_ROOTS_LIMIT
	pendingDepositsLimit       = 1 << 27 // PENDING_DEPOSITS_LIMIT
	pendingWithdrawalsLimit    = 1 << 27 // PENDING_PARTIAL_WITHDRAWALS_LIMIT
	pendingConsolidationsLimit = 1 << 18 // PENDING_CONSOLIDATIONS_LIMIT
	extraDataLimit             = 32      // MAX_EXTRA_DATA_BYTES
)

// LiteBeaconState is a minimal Fulu beacon state reduced to what proof
// building needs.
//
// The beacon state is a large SSZ container with 38 fields, represented
// compactly by a single Merkle root. Each field is either a value or a
// subtree. Because a Merkle root only depends on its children's hashes, any
// field that is a subtree can be represented by just its root. For our
// purposes, we only need the full subtrees of 2 fields: the state_roots
// vector and the latest_execution_payload_header container. All other
// subtree fields are represented by their Merkle root, avoiding the full
// state tree which is gigabytes in size, dominated by the millions of
// validators in the beacon chain.
type LiteBeaconState struct {
	fieldRoots           [numStateFieldLeaves][32]byte
	stateRoots           [][32]byte
	execHeaderFieldRoots [][32]byte
}

// ParseLiteBeaconState deserializes an SSZ encoded Fulu BeaconState and
// reduces it to a LiteBeaconState. Field roots are computed per field, in
// consensus spec order; the full state tree is never built. A wrong or
// missing field changes the assembled root and is rejected by the tree
// constructors' root check.
func ParseLiteBeaconState(stateSSZ []byte) (*LiteBeaconState, error) {
	var state fulu.BeaconState
	if err := state.UnmarshalSSZ(stateSSZ); err != nil {
		return nil, fmt.Errorf("failed to deserialize beacon state: %w", err)
	}

	lite := &LiteBeaconState{}
	r := lite.fieldRoots[:]
	r[0] = uint64Root(state.GenesisTime)
	r[1] = state.GenesisValidatorsRoot
	r[2] = uint64Root(uint64(state.Slot))
	r[3] = mustRoot(state.Fork.HashTreeRoot())
	r[4] = mustRoot(state.LatestBlockHeader.HashTreeRoot())
	r[5] = rootsVectorRoot(state.BlockRoots)
	r[6] = rootsVectorRoot(state.StateRoots)
	r[7] = rootsListRoot(state.HistoricalRoots, historicalRootsLimit)
	r[8] = mustRoot(state.ETH1Data.HashTreeRoot())
	r[9] = containersListRoot(state.ETH1DataVotes, eth1DataVotesLimit)
	r[10] = uint64Root(state.ETH1DepositIndex)
	r[11] = containersListRoot(state.Validators, validatorRegistryLimit)
	r[12] = uint64sListRoot(gweiToUint64s(state.Balances), validatorRegistryLimit)
	r[13] = rootsVectorRoot(state.RANDAOMixes)
	r[14] = uint64sVectorRoot(gweiToUint64s(state.Slashings))
	r[15] = bytesListRoot(participationToBytes(state.PreviousEpochParticipation), validatorRegistryLimit)
	r[16] = bytesListRoot(participationToBytes(state.CurrentEpochParticipation), validatorRegistryLimit)
	r[17] = bytesChunkRoot(state.JustificationBits)
	r[18] = mustRoot(state.PreviousJustifiedCheckpoint.HashTreeRoot())
	r[19] = mustRoot(state.CurrentJustifiedCheckpoint.HashTreeRoot())
	r[20] = mustRoot(state.FinalizedCheckpoint.HashTreeRoot())
	r[21] = uint64sListRoot(state.InactivityScores, validatorRegistryLimit)
	r[22] = mustRoot(state.CurrentSyncCommittee.HashTreeRoot())
	r[23] = mustRoot(state.NextSyncCommittee.HashTreeRoot())
	// r[24], latest_execution_payload_header, is set from its field roots
	// below so it is byte-identical to the expanded subtree's root.
	r[25] = uint64Root(uint64(state.NextWithdrawalIndex))
	r[26] = uint64Root(uint64(state.NextWithdrawalValidatorIndex))
	r[27] = containersListRoot(state.HistoricalSummaries, historicalSummariesLimit)
	r[28] = uint64Root(state.DepositRequestsStartIndex)
	r[29] = uint64Root(uint64(state.DepositBalanceToConsume))
	r[30] = uint64Root(uint64(state.ExitBalanceToConsume))
	r[31] = uint64Root(uint64(state.EarliestExitEpoch))
	r[32] = uint64Root(uint64(state.ConsolidationBalanceToConsume))
	r[33] = uint64Root(uint64(state.EarliestConsolidationEpoch))
	r[34] = containersListRoot(state.PendingDeposits, pendingDepositsLimit)
	r[35] = containersListRoot(state.PendingPartialWithdrawals, pendingWithdrawalsLimit)
	r[36] = containersListRoot(state.PendingConsolidations, pendingConsolidationsLimit)
	r[37] = uint64sVectorRoot(validatorIndicesToUint64s(state.ProposerLookahead))

	lite.stateRoots = rootsToFixed(state.StateRoots)
	lite.execHeaderFieldRoots = execHeaderFieldRoots(state.LatestExecutionPayloadHeader)
	r[execPayloadHeaderLeafIndex] = merkleizeChunks(lite.execHeaderFieldRoots)

	return lite, nil
}

// AnchorStateTree builds the anchor lite beacon state tree, verified against
// expectedRoot. state_roots is expanded as a real subtree and every other
// field is kept as a root leaf. To be clear, this is a 64-leaf tree, all
// collapsed to field roots except state_roots, which is a real 8192-leaf subtree
// inside it.
func (s *LiteBeaconState) AnchorStateTree(expectedRoot common.Hash) (*ssz.Node, error) {
	stateRootsSubtree, err := ssz.TreeFromChunks(toByteSlices(s.stateRoots))
	if err != nil {
		return nil, fmt.Errorf("failed to build state_roots subtree: %w", err)
	}
	return s.assemble(map[int]*ssz.Node{stateRootsLeafIndex: stateRootsSubtree}, expectedRoot)
}

// TargetStateTree returns two trees:
//
//  1. The target lite beacon state tree, whose root is found at the target
//     slot's position in the anchor state's state_roots vector.
//  2. The execution payload header subtree, contained inside the target
//     beacon state tree. Every other field is kept as a root leaf.
//
// Two trees are returned, instead of one like AnchorStateTree, because
// target state execution proofs need both trees: the header inclusion proof
// against the state tree, and the receipts root proof within the header
// subtree itself.
func (s *LiteBeaconState) TargetStateTree(expectedRoot common.Hash) (*ssz.Node, *ssz.Node, error) {
	execHeaderSubtree, err := ssz.TreeFromChunks(toByteSlices(s.execHeaderFieldRoots))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build execution header subtree: %w", err)
	}
	beaconStateTree, err := s.assemble(map[int]*ssz.Node{execPayloadHeaderLeafIndex: execHeaderSubtree}, expectedRoot)
	if err != nil {
		return nil, nil, err
	}
	return beaconStateTree, execHeaderSubtree, nil
}

// assemble builds the depth-6 beacon state tree.
// Expanded subtrees are provided in the map, keyed by their leaf index.
func (s *LiteBeaconState) assemble(
	expanded map[int]*ssz.Node,
	expectedRoot common.Hash,
) (*ssz.Node, error) {
	leaves := make([]*ssz.Node, numStateFieldLeaves)
	for i := range leaves {
		if subtree, ok := expanded[i]; ok {
			leaves[i] = subtree
			continue
		}
		root := s.fieldRoots[i]
		leaves[i] = ssz.LeafFromBytes(root[:])
	}
	tree, err := ssz.TreeFromNodes(leaves, len(leaves))
	if err != nil {
		return nil, fmt.Errorf("failed to assemble lite state tree: %w", err)
	}

	if got := common.BytesToHash(tree.Hash()); got != expectedRoot {
		return nil, fmt.Errorf("lite state root mismatch: computed %s, expected %s", got, expectedRoot)
	}
	return tree, nil
}

// execHeaderFieldRoots reduces the execution payload header to per-field
// roots, padded to numExecHeaderLeaves. receipts_root is field 3 (gindex 35).
func execHeaderFieldRoots(h *deneb.ExecutionPayloadHeader) [][32]byte {
	roots := make([][32]byte, numExecHeaderLeaves)
	roots[0] = h.ParentHash
	roots[1] = addressRoot(h.FeeRecipient)
	roots[2] = h.StateRoot
	roots[3] = h.ReceiptsRoot
	roots[4] = bloomRoot(h.LogsBloom)
	roots[5] = h.PrevRandao
	roots[6] = uint64Root(h.BlockNumber)
	roots[7] = uint64Root(h.GasLimit)
	roots[8] = uint64Root(h.GasUsed)
	roots[9] = uint64Root(h.Timestamp)
	roots[10] = bytesListRoot(h.ExtraData, extraDataLimit)
	roots[11] = uint256Root(h.BaseFeePerGas)
	roots[12] = h.BlockHash
	roots[13] = h.TransactionsRoot
	roots[14] = h.WithdrawalsRoot
	roots[15] = uint64Root(h.BlobGasUsed)
	roots[16] = uint64Root(h.ExcessBlobGas)
	return roots
}

// --- field root helpers --------------------------------------------------
//
// Every SSZ field reduces to a 32-byte root in one of a few shapes:
// basic values pack little-endian into a chunk; roots pass through;
// containers use their own HashTreeRoot; vectors merkleize their chunks;
// lists merkleize padded to their limit, then mix in the length.

// uint64Root packs a basic value into a 32-byte chunk.
func uint64Root(v uint64) [32]byte {
	var root [32]byte
	binary.LittleEndian.PutUint64(root[:8], v)
	return root
}

// uint256Root packs a uint256 little-endian into a chunk.
func uint256Root(v *uint256.Int) [32]byte {
	if v == nil {
		return [32]byte{}
	}
	be := v.Bytes32()
	var root [32]byte
	for i := range be {
		root[i] = be[31-i]
	}
	return root
}

// addressRoot right-pads a 20-byte address into a chunk.
func addressRoot(addr [20]byte) [32]byte {
	var root [32]byte
	copy(root[:20], addr[:])
	return root
}

// bytesChunkRoot packs a short byte value (e.g. a bitvector) into a chunk.
func bytesChunkRoot(b []byte) [32]byte {
	var root [32]byte
	copy(root[:], b)
	return root
}

// bloomRoot merkleizes a 256-byte logs bloom as its 8 chunks.
func bloomRoot(bloom [256]byte) [32]byte {
	chunks := make([][32]byte, 8)
	for i := range chunks {
		copy(chunks[i][:], bloom[i*32:(i+1)*32])
	}
	return merkleizeChunks(chunks)
}

// rootsVectorRoot merkleizes a fixed-size vector of roots.
func rootsVectorRoot(roots []phase0.Root) [32]byte {
	return merkleizeChunks(rootsToFixed(roots))
}

// uint64sVectorRoot merkleizes a fixed-size vector of uint64s, packed four
// per chunk.
func uint64sVectorRoot(values []uint64) [32]byte {
	return merkleizeChunks(packUint64s(values))
}

// rootsListRoot merkleizes a list of roots padded to limit, mixing in the
// element count.
func rootsListRoot(roots []phase0.Root, limit uint64) [32]byte {
	return mixinLength(merkleizeToLimit(rootsToFixed(roots), limit), uint64(len(roots)))
}

// uint64sListRoot merkleizes a list of uint64s (packed four per chunk) to
// its limit, mixing in the element count.
func uint64sListRoot(values []uint64, limit uint64) [32]byte {
	return mixinLength(merkleizeToLimit(packUint64s(values), (limit+3)/4), uint64(len(values)))
}

// bytesListRoot merkleizes a byte list (packed 32 per chunk) to its limit,
// mixing in the byte count.
func bytesListRoot(data []byte, limit uint64) [32]byte {
	chunks := make([][32]byte, (len(data)+31)/32)
	for i := range chunks {
		copy(chunks[i][:], data[i*32:min((i+1)*32, len(data))])
	}
	return mixinLength(merkleizeToLimit(chunks, (limit+31)/32), uint64(len(data)))
}

// containersListRoot merkleizes a list of SSZ containers to its limit,
// mixing in the element count. Each element's root comes from its own
// HashTreeRoot.
func containersListRoot[T interface{ HashTreeRoot() ([32]byte, error) }](items []T, limit uint64) [32]byte {
	roots := make([][32]byte, len(items))
	for i, item := range items {
		roots[i] = mustRoot(item.HashTreeRoot())
	}
	return mixinLength(merkleizeToLimit(roots, limit), uint64(len(items)))
}

// --- merkleization primitives ---------------------------------------------

// merkleizeChunks pairwise-hashes chunks up to a single root, zero-padding
// to the next power of two.
func merkleizeChunks(chunks [][32]byte) [32]byte {
	return merkleizeToLimit(chunks, uint64(nextPowerOfTwo(len(chunks))))
}

// merkleizeToLimit merkleizes chunks as the leaf layer of a tree with
// limit leaves (virtually padded with zero subtrees), per SSZ merkleization.
func merkleizeToLimit(chunks [][32]byte, limit uint64) [32]byte {
	depth := 0
	for 1<<depth < int(limit) {
		depth++
	}

	nodes := make([][32]byte, len(chunks))
	copy(nodes, chunks)
	zero := [32]byte{}
	for level := 0; level < depth; level++ {
		if len(nodes)%2 == 1 {
			nodes = append(nodes, zero)
		}
		for i := 0; i < len(nodes)/2; i++ {
			nodes[i] = hashPair(nodes[2*i], nodes[2*i+1])
		}
		nodes = nodes[:len(nodes)/2]
		zero = hashPair(zero, zero)
	}
	if len(nodes) == 0 {
		return zero
	}
	return nodes[0]
}

// mixinLength completes list merkleization: hash(chunksRoot, length).
func mixinLength(root [32]byte, length uint64) [32]byte {
	return hashPair(root, uint64Root(length))
}

func hashPair(a, b [32]byte) [32]byte {
	return sha256.Sum256(append(a[:], b[:]...))
}

func nextPowerOfTwo(n int) int {
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

// mustRoot unwraps typed hashing, which cannot fail on deserialized data.
func mustRoot(root [32]byte, err error) [32]byte {
	if err != nil {
		panic(err)
	}
	return root
}

// --- type conversions ------------------------------------------------------

// rootsToFixed converts attestantio's typed roots to plain 32-byte arrays.
func rootsToFixed(roots []phase0.Root) [][32]byte {
	out := make([][32]byte, len(roots))
	for i, r := range roots {
		out[i] = r
	}
	return out
}

// packUint64s packs values little-endian, four per 32-byte chunk.
func packUint64s(values []uint64) [][32]byte {
	chunks := make([][32]byte, (len(values)+3)/4)
	for i, v := range values {
		binary.LittleEndian.PutUint64(chunks[i/4][(i%4)*8:], v)
	}
	return chunks
}

func gweiToUint64s(values []phase0.Gwei) []uint64 {
	out := make([]uint64, len(values))
	for i, v := range values {
		out[i] = uint64(v)
	}
	return out
}

func validatorIndicesToUint64s(values []phase0.ValidatorIndex) []uint64 {
	out := make([]uint64, len(values))
	for i, v := range values {
		out[i] = uint64(v)
	}
	return out
}

func participationToBytes(flags []altair.ParticipationFlags) []byte {
	out := make([]byte, len(flags))
	for i, f := range flags {
		out[i] = byte(f)
	}
	return out
}

// toByteSlices converts fixed roots to the [][]byte TreeFromChunks takes.
func toByteSlices(roots [][32]byte) [][]byte {
	out := make([][]byte, len(roots))
	for i := range roots {
		out[i] = make([]byte, 32)
		copy(out[i], roots[i][:])
	}
	return out
}
