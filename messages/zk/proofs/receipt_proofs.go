// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb"
)

// ethClient is the subset of ethclient.Client used for proof building.
type ethClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	BlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]*types.Receipt, error)
}

// ReceiptProof is a Merkle Patricia Tree (MPT) inclusion proof for
// a specific log inside a receipt and is against a block's receipts root.
// The field layout below mirrors the ZKAdapter's Receipt.Proof struct
// See: https://github.com/ava-labs/icm-services/blob/39eb811943f8017843e30dfeb059e829fe65099e/icm-contracts/avalanche/verifiers/ethereum/StateManagerLibrary.sol#L160
//
//nolint:lll
type ReceiptProof struct {
	// Proof is the list of trie nodes on the path from the receipts root to
	// the receipt, in root-to-leaf order.
	Proof [][]byte
	// Key is the RLP-encoded transaction index of the receipt in the block.
	Key []byte
	// Value is the RLP-encoded (EIP-2718 typed) receipt data.
	Value []byte
	// LogIndex of the message event within the receipt.
	LogIndex uint
	// ExpectedEmitter is the contract that emitted the log.
	ExpectedEmitter common.Address
	// ExpectedTopic0 is the log's topic0 (event signature hash).
	ExpectedTopic0 common.Hash
}

// BuildReceiptProof reconstructs the receipts trie for the block containing
// txHash, verifies its root against the block header, and extracts the
// inclusion proof for that transaction's receipt and the log at logIndex.
func BuildReceiptProof(
	ctx context.Context,
	client ethClient,
	blockNumber uint64,
	txHash common.Hash,
	logIndex uint,
) (*ReceiptProof, error) {
	// Get the target receipt corresponding to txHash and its log.
	header, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch header %d: %w", blockNumber, err)
	}
	// RequireCanonical=false: we want the receipts of exactly this header's block. Whether or not
	// the block is canonical is already enforced by finality depth and by proof verification against
	// a confirmed consensus anchor.
	receipts, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHashWithHash(header.Hash(), false))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch receipts for block %d: %w", blockNumber, err)
	}
	targetReceipt, err := receiptByTxHash(receipts, txHash)
	if err != nil {
		return nil, err
	}
	if int(logIndex) >= len(targetReceipt.Logs) {
		return nil, fmt.Errorf(
			"log index %d out of range (total %d logs) in receipt %s",
			logIndex, 
			len(targetReceipt.Logs), 
			txHash.Hex(),
		)
	}
	targetLog := targetReceipt.Logs[logIndex]

	// Build the receipts trie and verify its root against the header receipts root.
	receiptsTrie, targetValue, err := buildReceiptsTrie(
		receipts,
		targetReceipt.TransactionIndex,
		header.ReceiptHash,
	)
	if err != nil {
		return nil, err
	}

	// Compute the inclusion proof for the target receipt.
	key, err := rlp.EncodeToBytes(targetReceipt.TransactionIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to encode target trie key: %w", err)
	}
	var proof proofList
	if err := receiptsTrie.Prove(key, &proof); err != nil {
		return nil, fmt.Errorf("failed to generate receipt proof: %w", err)
	}

	return &ReceiptProof{
		Proof:           proof,
		Key:             key,
		Value:           targetValue,
		LogIndex:        logIndex,
		ExpectedEmitter: targetLog.Address,
		ExpectedTopic0:  targetLog.Topics[0],
	}, nil
}

// buildReceiptsTrie reconstructs the block's receipts trie and returns it
// along with the encoded target receipt. The trie root is verified against
// the header's receipts root, so any encoding drift fails here rather than
// producing proofs against the wrong tree.
func buildReceiptsTrie(
	receipts []*types.Receipt,
	targetIndex uint,
	receiptsRoot common.Hash,
) (*trie.Trie, []byte, error) {
	receiptsTrie := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	var valueBuf bytes.Buffer
	var targetValue []byte
	receiptList := types.Receipts(receipts)
	for i := range receipts {
		key, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to encode trie key %d: %w", i, err)
		}
		valueBuf.Reset()
		receiptList.EncodeIndex(i, &valueBuf)
		value := bytes.Clone(valueBuf.Bytes())
		if err := receiptsTrie.Update(key, value); err != nil {
			return nil, nil, fmt.Errorf("failed to insert receipt %d into trie: %w", i, err)
		}
		if uint(i) == targetIndex {
			targetValue = value
		}
	}

	if receiptsTrie.Hash() != receiptsRoot {
		return nil, nil, fmt.Errorf(
			"receipts trie root mismatch: computed %s, header %s",
			receiptsTrie.Hash().Hex(),
			receiptsRoot.Hex(),
		)
	}
	return receiptsTrie, targetValue, nil
}

func receiptByTxHash(receipts []*types.Receipt, txHash common.Hash) (*types.Receipt, error) {
	for _, r := range receipts {
		if r.TxHash == txHash {
			return r, nil
		}
	}
	return nil, fmt.Errorf("receipt for tx %s not found in block", txHash.Hex())
}

// proofList collects trie proof nodes in insertion (root-to-leaf) order.
type proofList [][]byte

func (p *proofList) Put(_ []byte, value []byte) error {
	*p = append(*p, value)
	return nil
}

func (p *proofList) Delete([]byte) error { return nil }
