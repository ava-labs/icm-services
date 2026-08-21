package proofs

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb"
)

// ethClient is the subset of ethclient.Client used by proof building.
// *ethclient.Client satisfies it; tests substitute a fake.
type ethClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// ReceiptProof is a Merkle Patricia Tree (MPT) inclusion proof for
// a specific log inside a receipt and is against a block's receipts root.
// The field layout below mirrors the ZKAdapter's Receipt.Proof struct
// nolint:lll
// See: https://github.com/ava-labs/icm-services/blob/39eb811943f8017843e30dfeb059e829fe65099e/icm-contracts/avalanche/verifiers/ethereum/StateManagerLibrary.sol#L160
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
	header, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch header %d: %w", blockNumber, err)
	}

	receipts, err := blockReceipts(ctx, client, header)
	if err != nil {
		return nil, err
	}

	// Locate the target receipt and its log.
	targetReceipt, err := receiptByTxHash(receipts, txHash)
	if err != nil {
		return nil, err
	}
	if int(logIndex) >= len(targetReceipt.Logs) {
		return nil, fmt.Errorf("log index %d out of range (%d logs) in receipt %s",
			logIndex, len(targetReceipt.Logs), txHash.Hex())
	}
	targetLog := targetReceipt.Logs[logIndex]

	// Rebuild the receipts trie. types.Receipts implements DerivableList, so
	// EncodeIndex produces the canonical EIP-2718 typed receipt encoding.
	receiptsTrie := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	var valueBuf bytes.Buffer
	var targetValue []byte
	receiptList := types.Receipts(receipts)
	for i := range receipts {
		key, err := rlp.EncodeToBytes(uint(i))
		if err != nil {
			return nil, fmt.Errorf("failed to encode trie key %d: %w", i, err)
		}
		valueBuf.Reset()
		receiptList.EncodeIndex(i, &valueBuf)
		value := bytes.Clone(valueBuf.Bytes())
		if err := receiptsTrie.Update(key, value); err != nil {
			return nil, fmt.Errorf("failed to insert receipt %d into trie: %w", i, err)
		}
		if uint(i) == targetReceipt.TransactionIndex {
			targetValue = value
		}
	}

	// The reconstructed root must match the block header, or the proof would
	// be built against the wrong tree.
	if receiptsTrie.Hash() != header.ReceiptHash {
		return nil, fmt.Errorf("receipts trie root mismatch: computed %s, header %s",
			receiptsTrie.Hash().Hex(), header.ReceiptHash.Hex())
	}

	// Extract the inclusion proof for the target receipt.
	key, err := rlp.EncodeToBytes(uint(targetReceipt.TransactionIndex))
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

// blockReceipts fetches every receipt in the block, preferring the batch
// eth_getBlockReceipts endpoint and falling back to per-transaction fetches.
func blockReceipts(
	ctx context.Context,
	client ethClient,
	header *types.Header,
) ([]*types.Receipt, error) {
	// TODO: use client.BlockReceipts if the libevm ethclient version exposes
	// it; the per-tx fallback below is correct but slower for large blocks.
	block, err := client.BlockByNumber(ctx, header.Number)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch block %d: %w", header.Number, err)
	}
	receipts := make([]*types.Receipt, len(block.Transactions()))
	for i, tx := range block.Transactions() {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, fmt.Errorf("failed to fetch receipt %s: %w", tx.Hash().Hex(), err)
		}
		receipts[i] = receipt
	}
	return receipts, nil
}

// receiptByTxHash returns the receipt for the given transaction hash.
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
