package proofs

import (
	"context"
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/stretchr/testify/require"
)

// fakeEthClient returns fixed test data instead of querying a real chain
type fakeEthClient struct {
	header   *types.Header
	block    *types.Block
	receipts map[common.Hash]*types.Receipt
}

func (f *fakeEthClient) HeaderByNumber(_ context.Context, _ *big.Int) (*types.Header, error) {
	return f.header, nil
}

func (f *fakeEthClient) BlockByNumber(_ context.Context, _ *big.Int) (*types.Block, error) {
	return f.block, nil
}

func (f *fakeEthClient) TransactionReceipt(_ context.Context, txHash common.Hash) (*types.Receipt, error) {
	return f.receipts[txHash], nil
}

// makeTestBlock builds a synthetic block of n transactions with receipts,
// whose header carries the correct receipts root for those receipts.
func makeTestBlock(t *testing.T, n int) (*fakeEthClient, types.Receipts) {
	t.Helper()

	txs := make([]*types.Transaction, n)
	receipts := make(types.Receipts, n)
	byHash := make(map[common.Hash]*types.Receipt, n)
	for i := 0; i < n; i++ {
		to := common.BytesToAddress([]byte{byte(i + 1)})
		txs[i] = types.NewTransaction(uint64(i), to, big.NewInt(1), 21000, big.NewInt(1), nil)

		receipt := &types.Receipt{
			Type:              types.LegacyTxType,
			Status:            types.ReceiptStatusSuccessful,
			CumulativeGasUsed: uint64(21000 * (i + 1)),
			Logs: []*types.Log{
				{
					Address: common.BytesToAddress([]byte{0xEE, byte(i)}),
					Topics:  []common.Hash{common.BytesToHash([]byte{0xAA, byte(i)})},
					Data:    []byte{byte(i)},
				},
				{
					Address: common.BytesToAddress([]byte{0xFF, byte(i)}),
					Topics:  []common.Hash{common.BytesToHash([]byte{0xBB, byte(i)})},
					Data:    []byte{byte(i), byte(i)},
				},
			},
			TxHash:           txs[i].Hash(),
			TransactionIndex: uint(i),
		}
		receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
		receipts[i] = receipt
		byHash[txs[i].Hash()] = receipt
	}

	header := &types.Header{
		Number:      big.NewInt(100),
		ReceiptHash: types.DeriveSha(receipts, trie.NewStackTrie(nil)),
	}
	block := types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil))

	return &fakeEthClient{header: header, block: block, receipts: byHash}, receipts
}

// verifyProofAgainstRoot cryptographically verifies the returned proof nodes
// against the receipts root and returns the proven value.
func verifyProofAgainstRoot(t *testing.T, root common.Hash, key []byte, nodes [][]byte) []byte {
	t.Helper()
	proofDB := rawdb.NewMemoryDatabase()
	for _, node := range nodes {
		require.NoError(t, proofDB.Put(crypto.Keccak256(node), node))
	}
	value, err := trie.VerifyProof(root, key, proofDB)
	require.NoError(t, err)
	return value
}

// Happy path: the proof verifies against the header's receipts root and the
// proven value is the canonical encoding of the target receipt.
func TestBuildReceiptProof(t *testing.T) {
	client, receipts := makeTestBlock(t, 3)
	targetIdx := 1
	targetLogIdx := uint(1)
	targetHash := receipts[targetIdx].TxHash

	proof, err := BuildReceiptProof(context.Background(), client, 100, targetHash, targetLogIdx)
	require.NoError(t, err)

	// The key must be the RLP-encoded transaction index.
	expectedKey, err := rlp.EncodeToBytes(uint(targetIdx))
	require.NoError(t, err)
	require.Equal(t, expectedKey, proof.Key)

	// The proof must verify against the header's receipts root, and the
	// proven value must equal the returned Value.
	provenValue := verifyProofAgainstRoot(t, client.header.ReceiptHash, proof.Key, proof.Proof)
	require.Equal(t, proof.Value, provenValue)

	// Log metadata must identify the requested log.
	targetLog := receipts[targetIdx].Logs[targetLogIdx]
	require.Equal(t, targetLogIdx, proof.LogIndex)
	require.Equal(t, targetLog.Address, proof.ExpectedEmitter)
	require.Equal(t, targetLog.Topics[0], proof.ExpectedTopic0)
}

// Every receipt in the block must be provable, not just a particular index.
func TestBuildReceiptProofAllIndices(t *testing.T) {
	client, receipts := makeTestBlock(t, 5)
	for i, receipt := range receipts {
		proof, err := BuildReceiptProof(context.Background(), client, 100, receipt.TxHash, 0)
		require.NoError(t, err, "receipt %d", i)
		provenValue := verifyProofAgainstRoot(t, client.header.ReceiptHash, proof.Key, proof.Proof)
		require.Equal(t, proof.Value, provenValue, "receipt %d", i)
	}
}

// A log index beyond the receipt's logs must error.
func TestBuildReceiptProofLogIndexOutOfRange(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 99)
	require.ErrorContains(t, err, "log index")
}

// A transaction hash not present in the block must error.
func TestBuildReceiptProofTxNotInBlock(t *testing.T) {
	client, _ := makeTestBlock(t, 2)
	_, err := BuildReceiptProof(context.Background(), client, 100, common.HexToHash("0xdead"), 0)
	require.ErrorContains(t, err, "not found in block")
}

// A header whose receipts root doesn't match the block's receipts must error
// rather than produce a proof against the wrong tree.
func TestBuildReceiptProofRootMismatch(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	client.header.ReceiptHash = common.HexToHash("0xbadbadbad")
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 0)
	require.ErrorContains(t, err, "root mismatch")
}
