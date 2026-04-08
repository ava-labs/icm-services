/**
 * Generates a test fixture for ZKAdapter e2e tests.
 * See README.md for additional details regarding functionality, usage, and dependencies. 
 */

import { ssz } from "@lodestar/types";
import { Tree } from "@chainsafe/persistent-merkle-tree";
import { ethers } from "ethers";
import { Trie } from "@ethereumjs/trie";
import { RLP } from "@ethereumjs/rlp";
import * as fs from "fs";
import * as path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Configuration
const BEACON_API_URL = process.env.BEACON_API_URL!;
const ETH_RPC_URL = process.env.ETH_RPC_URL!;
const TX_HASH = process.env.TX_HASH!;
const LOG_INDEX = parseInt(process.env.LOG_INDEX || "0");

if (!BEACON_API_URL || !ETH_RPC_URL || !TX_HASH) {
  console.error("Required env vars: BEACON_API_URL, ETH_RPC_URL, TX_HASH");
  console.error("Optional: LOG_INDEX (default: 0)");
  process.exit(1);
}


// Fulu generalized indices 
const G_INDEX_BLOCK_STATE_ROOT  = BigInt(11);  
const G_INDEX_BASE_STATE_ROOTS  = BigInt(70);  
const STATE_ROOTS_DEPTH         = BigInt(13);  
const G_INDEX_EXEC_ROOT         = BigInt(88);   
const G_INDEX_RECEIPTS_ROOT     = BigInt(35);  
const STATE_ROOTS_VECTOR_SIZE   = BigInt(8192);

// Ethereum beacon chain genesis time
const MAINNET_BEACON_GENESIS_TIME = 1606824023;
const SECONDS_PER_SLOT = 12;

// Helpers
function toHex(bytes: Uint8Array): string {
  return "0x" + Buffer.from(bytes).toString("hex");
}

function blockTimestampToSlot(timestamp: number): number {
  return Math.floor((timestamp - MAINNET_BEACON_GENESIS_TIME) / SECONDS_PER_SLOT);
}

// Fetches raw SSZ-encoded bytes from the beacon API 
async function fetchSSZBytes(url: string): Promise<Uint8Array> {
  const res = await fetch(url, {
    headers: { Accept: "application/octet-stream" },
  });
  if (!res.ok) throw new Error(`Fetch failed: ${res.status} ${res.statusText} for ${url}`);
  return new Uint8Array(await res.arrayBuffer());
}

// Fetches the beacon block root for a given slot
async function fetchBeaconBlockRoot(slot: number): Promise<string> {
  const res = await fetch(`${BEACON_API_URL}/eth/v1/beacon/blocks/${slot}/root`);
  if (!res.ok) throw new Error(`Block root fetch failed: ${res.status}`);
  const json = (await res.json()) as { data: { root: string } };
  return json.data.root;
}

// Scans forward from startSlot until a slot with a block is found, skipping missed slots
async function findNextValidSlot(startSlot: number, maxAttempts: number = 32): Promise<number> {
  for (let i = 0; i < maxAttempts; i++) {
    const slot = startSlot + i;
    const res = await fetch(`${BEACON_API_URL}/eth/v1/beacon/blocks/${slot}/root`);
    if (res.ok) return slot;
  }
  throw new Error(`No valid block found in ${maxAttempts} slots starting from ${startSlot}`);
}

// Extracts an SSZ Merkle proof for a given generalized index from an SSZ tree node
function extractProof(node: any, gIndex: bigint): string[] {
  const tree = new Tree(node);
  const proof = tree.getSingleProof(gIndex);
  return proof.map((p: Uint8Array) => toHex(p));
}

// RLP-encodes a transaction receipt, prefixing with the transaction type for EIP-2718 typed receipts
function encodeReceipt(receipt: ethers.providers.TransactionReceipt): Buffer {
  const status = receipt.status === 1 ? Buffer.from([1]) : Buffer.from([]);
  
  // Strip 0x prefix and leading zeros, then pad to even length for valid RLP encoding
  let gasHex = receipt.cumulativeGasUsed.toHexString().slice(2).replace(/^0+/, "") || "0";
  if (gasHex.length % 2 !== 0) gasHex = "0" + gasHex;
  const cumulativeGas = Buffer.from(gasHex, "hex");
  
  // Strip 0x prefix and convert logs bloom to bytes
  const logsBloom = Buffer.from(receipt.logsBloom.slice(2), "hex");

  // Encode logs as [address, topics, data] byte arrays, stripping 0x prefixes
  const encodedLogs = receipt.logs.map((log) => [
    Buffer.from(log.address.slice(2), "hex"),
    log.topics.map((t) => Buffer.from(t.slice(2), "hex")),
    Buffer.from(log.data.slice(2), "hex"),
  ]);

  // RLP encode the receipt
  const body = RLP.encode([status, cumulativeGas, logsBloom, encodedLogs]);

  // From EIP-2718, prefix with transaction type for non-legacy receipts
  if (receipt.type === 0 || receipt.type === undefined) {
    return Buffer.from(body);
  }

  // Prepend the transaction type byte to the RLP-encoded body
  const typed = Buffer.alloc(1 + body.length);
  typed[0] = receipt.type;
  Buffer.from(body).copy(typed, 1);
  return typed;
}

// Reconstructs the receipts Merkle Patricia Tree (MPT) from all transactions in the block, verifies the root matches
// the block header, then generates and returns a MPT proof for the target transaction's receipt
async function buildReceiptProof(
  provider: ethers.providers.JsonRpcProvider,
  txHash: string
): Promise<{
  receiptsRoot: string;
  proofNodes: string[];
  key: string;
  value: string;
  receipt: ethers.providers.TransactionReceipt;
}> {
  console.log("Fetching transaction receipt...");
  const receipt = await provider.getTransactionReceipt(txHash);
  if (!receipt) throw new Error(`Receipt not found for ${txHash}`);
  console.log(`Transaction in block ${receipt.blockNumber}, index ${receipt.transactionIndex}`);

  const rawBlock = await provider.send("eth_getBlockByNumber", [
    ethers.utils.hexValue(receipt.blockNumber),
    false,
  ]);
  const receiptsRoot: string = rawBlock.receiptsRoot;
  console.log(`Block receipts root from header: ${receiptsRoot}`);

  console.log("Fetching all receipts in block...");
  const blockWithTxs = await provider.getBlockWithTransactions(receipt.blockNumber);
  const allReceipts: ethers.providers.TransactionReceipt[] = [];
  for (const tx of blockWithTxs.transactions) {
    const r = await provider.getTransactionReceipt(tx.hash);
    allReceipts.push(r);
  }
  console.log(`Fetched ${allReceipts.length} receipts`);

  allReceipts.sort((a, b) => a.transactionIndex - b.transactionIndex);

  console.log("Building receipts trie...");
  const trie = new Trie();
  for (let i = 0; i < allReceipts.length; i++) {
    const key = Buffer.from(RLP.encode(i));
    const value = encodeReceipt(allReceipts[i]);
    await trie.put(key, value);
  }

  const trieRoot = toHex(trie.root());
  console.log(`Computed trie root:  ${trieRoot}`);
  if (trieRoot.toLowerCase() !== receiptsRoot.toLowerCase()) {
    throw new Error(`Trie root mismatch. Computed: ${trieRoot}, Expected: ${receiptsRoot}`);
  }
  console.log("Receipts trie root verified against expected value. \n");

  const key = Buffer.from(RLP.encode(receipt.transactionIndex));
  const proof = await trie.createProof(key);
  const value = await trie.get(key);
  if (!value) throw new Error("Could not retrieve value from trie");

  return {
    receiptsRoot,
    proofNodes: proof.map((node) => toHex(node)),
    key: toHex(key),
    value: toHex(Buffer.from(value)),
    receipt,
  };
}

async function main() {
  const provider = new ethers.providers.JsonRpcProvider(ETH_RPC_URL);

  // Get block info and compute slot
  console.log("=== Step 1: Fetching transaction data ===");
  const receipt = await provider.getTransactionReceipt(TX_HASH);
  if (!receipt) throw new Error(`Transaction ${TX_HASH} not found`);

  const block = await provider.getBlock(receipt.blockNumber);
  const targetSlot = blockTimestampToSlot(block.timestamp);
  console.log(`Block: ${block.number}, Target slot: ${targetSlot}`);

  // Find a valid anchor slot
  // Scan forward from targetSlot + 64 (2 epochs) to handle missed slots and to ensure slot is finalized
  console.log("Scanning for valid anchor slot...");
  const anchorSlot = await findNextValidSlot(targetSlot + 64);
  console.log(`Anchor slot: ${anchorSlot}\n`);

  if (BigInt(anchorSlot - targetSlot) > STATE_ROOTS_VECTOR_SIZE) {
    throw new Error("Target slot is too far from anchor slot");
  }

  // Fetch anchor beacon block root 
  console.log("=== Step 2: Fetching anchor beacon block root ===");
  const anchorBeaconBlockRoot = await fetchBeaconBlockRoot(anchorSlot);
  console.log(`Anchor beacon block root: ${anchorBeaconBlockRoot}\n`);

  // Fetch beacon block anchor proof (beacon block root -> beacon block state) 
  console.log("=== Step 3: Extracting anchor state proof (block -> state) ===");
  const anchorBlockBytes = await fetchSSZBytes(
    `${BEACON_API_URL}/eth/v2/beacon/blocks/${anchorSlot}`
  );
  console.log(`Anchor block SSZ: ${(anchorBlockBytes.length / 1024).toFixed(1)}KB`);

  const anchorBlockView = ssz.fulu.SignedBeaconBlock.deserializeToView(anchorBlockBytes);
  const anchorBeaconStateProof = extractProof(
    anchorBlockView.message.node,
    G_INDEX_BLOCK_STATE_ROOT
  );
  console.log(`Anchor state proof: ${anchorBeaconStateProof.length} siblings\n`);

  // Fetch anchor beacon state for history proof 
  console.log("=== Step 4: Fetching anchor beacon state ===");
  const anchorStateBytes = await fetchSSZBytes(
    `${BEACON_API_URL}/eth/v2/debug/beacon/states/${anchorSlot}`
  );
  console.log(`Anchor state SSZ: ${(anchorStateBytes.length / 1024 / 1024).toFixed(1)}MB`);

  const anchorStateView = ssz.fulu.BeaconState.deserializeToView(anchorStateBytes);
  const anchorBeaconStateRoot = toHex(anchorStateView.hashTreeRoot());
  console.log(`Anchor beacon state root: ${anchorBeaconStateRoot}`);

  
  // Compute the history proof. First, compute the gIndex for the target slot's state root 
  // within the anchor's state_roots vector, then extract the proof and the target beacon state root
  const vectorIndex = BigInt(targetSlot) % STATE_ROOTS_VECTOR_SIZE;
  const historyGIndex = (G_INDEX_BASE_STATE_ROOTS << STATE_ROOTS_DEPTH) + vectorIndex;
  console.log(`History gIndex: ${historyGIndex} (vectorIndex: ${vectorIndex})`);
  const targetBeaconStateProof = extractProof(anchorStateView.node, historyGIndex);
  const targetBeaconStateRoot = toHex(anchorStateView.stateRoots.get(Number(vectorIndex)));
  console.log(`Target beacon state root: ${targetBeaconStateRoot}`);
  console.log(`History proof: ${targetBeaconStateProof.length} siblings\n`);

  // Fetch the target beacon state for execution and receipts proofs 
  console.log("=== Step 5: Fetching target beacon state ===");
  const targetStateBytes = await fetchSSZBytes(
    `${BEACON_API_URL}/eth/v2/debug/beacon/states/${targetSlot}`
  );
  console.log(`Target state SSZ: ${(targetStateBytes.length / 1024 / 1024).toFixed(1)}MB`);

  const targetStateView = ssz.fulu.BeaconState.deserializeToView(targetStateBytes);

  // Compute the Execution proof. Prove that the execution payload header is part of the target beacon state
  const targetExecutionHeaderProof = extractProof(targetStateView.node, G_INDEX_EXEC_ROOT);
  const targetExecutionHeaderRoot = toHex(
    targetStateView.latestExecutionPayloadHeader.hashTreeRoot()
  );
  console.log(`Execution header root: ${targetExecutionHeaderRoot}`);
  console.log(`Execution proof: ${targetExecutionHeaderProof.length} siblings`);
  
  // Compute the Receipts proof. Prove that the receipts root is part of the execution payload header
  const execPayloadHeaderNode = targetStateView.latestExecutionPayloadHeader.node;
  const targetReceiptsProof = extractProof(execPayloadHeaderNode, G_INDEX_RECEIPTS_ROOT);
  const targetReceiptsRoot = toHex(
    targetStateView.latestExecutionPayloadHeader.receiptsRoot
  );
  console.log(`Receipts root (from SSZ): ${targetReceiptsRoot}`);
  console.log(`Receipts proof: ${targetReceiptsProof.length} siblings\n`);

  // Next, prove that the receipt is part of the receipts trie. 
  console.log("=== Step 6: Building MPT receipt proof ===");
  const receiptProofData = await buildReceiptProof(provider, TX_HASH);

  // Verify receipts root from SSZ matches what MPT expects
  if (targetReceiptsRoot.toLowerCase() !== receiptProofData.receiptsRoot.toLowerCase()) {
    throw new Error(
      `Receipts root mismatch! SSZ: ${targetReceiptsRoot}, Block header: ${receiptProofData.receiptsRoot}`
    );
  }
  console.log("SSZ receipts root matches block header receipts root!\n");

  // Extract event data 
  console.log("=== Step 7: Extracting event data ===");
  const targetLog = receiptProofData.receipt.logs[LOG_INDEX];
  if (!targetLog) throw new Error(`Log index ${LOG_INDEX} not found in receipt`);

  console.log(`Emitter: ${targetLog.address}`);
  console.log(`Topic0:  ${targetLog.topics[0]}`);
  console.log(`Data:    ${targetLog.data}`);
  console.log(`Log index in receipt: ${LOG_INDEX}\n`);

  // Write fixture 
  console.log("=== Step 8: Writing fixture ===");
  const fixture = {
    anchorBeaconBlockRoot: anchorBeaconBlockRoot,
    metadata: {
      description: "ZKAdapter e2e test fixture generated from Ethereum Mainnet (Fulu fork)",
      txHash: TX_HASH,
      blockNumber: receiptProofData.receipt.blockNumber,
      targetSlot,
      anchorSlot,
      fork: "fulu",
      network: "mainnet",
    },
    executionProof: {
      anchorSlot,
      targetSlot,
      anchorBeaconStateRoot,
      anchorBeaconStateProof,
      targetBeaconStateRoot,
      targetBeaconStateProof,
      targetExecutionHeaderRoot,
      targetExecutionHeaderProof,
      targetReceiptsRoot,
      targetReceiptsProof,
    },
    receiptProof: {
      proof: receiptProofData.proofNodes,
      key: receiptProofData.key,
      value: receiptProofData.value,
      logIndex: LOG_INDEX,
      expectedEmitter: targetLog.address,
      expectedTopic0: targetLog.topics[0],
    },
  };

  const outPath = path.join(__dirname, "testdata", "ethereum_fixture.json");
  fs.mkdirSync(path.dirname(outPath), { recursive: true });
  fs.writeFileSync(outPath, JSON.stringify(fixture, null, 2));
  console.log(`Fixture written to ${outPath}`);
  console.log(`\nTotal proof sizes:`);
  console.log(`  Anchor state proof:    ${anchorBeaconStateProof.length} siblings`);
  console.log(`  History proof:         ${targetBeaconStateProof.length} siblings`);
  console.log(`  Execution proof:       ${targetExecutionHeaderProof.length} siblings`);
  console.log(`  Receipts proof:        ${targetReceiptsProof.length} siblings`);
  console.log(`  MPT receipt proof:     ${receiptProofData.proofNodes.length} nodes`);
}

main().catch((err) => {
  console.error("\nError:", err.message || err);
  process.exit(1);
});
