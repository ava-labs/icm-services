/**
 * Given a Sepolia tx hash, polls the Boundless subgraph
 * until a ZK proof covers the tx's slot, then writes the Boundless fixture.
 * 
 * Required env vars:
 *   SUBGRAPH_URL   - Boundless subgraph GraphQL endpoint
 *   ETH_RPC_URL    - Sepolia execution layer RPC
 *   TX_HASH        - Sepolia tx hash to look for
 *
 * Output:
 *   testdata/boundless_fixture.json
 */

import { ethers } from "ethers";
import * as fs from "fs";
import * as path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const SUBGRAPH_URL = process.env.SUBGRAPH_URL!;
const ETH_RPC_URL = process.env.ETH_RPC_URL!;

const SEPOLIA_BEACON_GENESIS_TIME = 1655733600;
const SECONDS_PER_SLOT = 12;

// How long to wait for a proof before giving up (3 hours)
const MAX_WAIT_MS = 180 * 60 * 1000; 
// How often to poll the subgraph (1 minute)
const POLL_INTERVAL_MS = 60 * 1000;

if (!SUBGRAPH_URL || !ETH_RPC_URL) {
  console.error("Required env vars: SUBGRAPH_URL, ETH_RPC_URL");
  process.exit(1);
}

// Resolve tx hash from env var or from tx_info.json (written by send_sepolia_message.mts)
function resolveTxHash(): string {
  if (process.env.TX_HASH) return process.env.TX_HASH;
  
  const txInfoPath = path.join(__dirname, "testdata", "tx_info.json");
  if (fs.existsSync(txInfoPath)) {
    const txInfo = JSON.parse(fs.readFileSync(txInfoPath, "utf-8"));
    console.log(`Read tx hash from ${txInfoPath}`);
    return txInfo.txHash;
  }

  console.error("TX_HASH env var not set and testdata/tx_info.json not found");
  process.exit(1);
}

function blockTimestampToSlot(timestamp: number): number {
  return Math.floor((timestamp - SEPOLIA_BEACON_GENESIS_TIME) / SECONDS_PER_SLOT);
}

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function queryProof(minSlot: number): Promise<{
  finalizedSlot: number;
  preState: string;
  postState: string;
  journalData: string;
  seal: string;
} | null> {
  const res = await fetch(SUBGRAPH_URL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      query: `{
        signalEthereumProofs(
          first: 1,
          orderBy: finalizedSlot,
          orderDirection: desc,
          where: { finalizedSlot_gt: "${minSlot}" }
        ) {
          finalizedSlot
          preState
          postState
          journalData
          seal
        }
      }`,
    }),
  });

  const json = (await res.json()) as any;
  const proof = json.data?.signalEthereumProofs?.[0];
  if (!proof) return null;

  return {
    finalizedSlot: parseInt(proof.finalizedSlot),
    preState: proof.preState,
    postState: proof.postState,
    journalData: proof.journalData,
    seal: proof.seal,
  };
}

async function main() {
  const txHash = resolveTxHash();
  const provider = new ethers.providers.JsonRpcProvider(ETH_RPC_URL);

  // Look up the tx and compute its slot
  console.log("=== Step 1: Looking up tx slot ===");
  const receipt = await provider.getTransactionReceipt(txHash);
  if (!receipt) throw new Error(`Transaction ${txHash} not found`);
  const block = await provider.getBlock(receipt.blockNumber);
  const txSlot = blockTimestampToSlot(block.timestamp);
  console.log(`tx is in block ${receipt.blockNumber}, slot ${txSlot}`);

  // The proof's finalizedSlot must be after the tx's slot so the anchor can reach it.
  // We need finalizedSlot > txSlot and the tx must be within 8192 slots of the anchor.
  const minFinalizedSlot = txSlot + 1;
  console.log(`The minimum finalized slot is ${minFinalizedSlot}\n`);

  // Poll the Boundless subgraph until a suitable proof is available
  console.log("=== Step 2: Waiting for Boundless ZK proof ===");
  const startTime = Date.now();
  let proof = await queryProof(minFinalizedSlot);

  while (!proof) {
    const elapsed = Date.now() - startTime;
    if (elapsed > MAX_WAIT_MS) {
      throw new Error(`Timed out after ${MAX_WAIT_MS / 60000} minutes waiting for proof`);
    }
    const minutesLeft = ((MAX_WAIT_MS - elapsed) / 60000).toFixed(1);
    console.log(`No proof yet. Retrying in ${POLL_INTERVAL_MS / 1000}s (${minutesLeft}min remaining)...`);
    await sleep(POLL_INTERVAL_MS);
    proof = await queryProof(minFinalizedSlot);
  }

  console.log(`Found proof at finalized slot: ${proof.finalizedSlot}`);

  // Verify the tx is within range
  const slotDiff = proof.finalizedSlot - txSlot;
  if (slotDiff > 8192) {
    throw new Error(
      `Tx slot ${txSlot} is too far from finalized slot ${proof.finalizedSlot} (diff: ${slotDiff}, max: 8192)`
    );
  }
  console.log(`Anchor slot: ${proof.finalizedSlot}, Tx slot: ${txSlot}, diff: ${slotDiff} (within 8192 limit)\n`);

  // Save Boundless fixture
  console.log("=== Step 3: Writing Boundless fixture ===");
  const boundlessFixture = {
    finalizedSlot: proof.finalizedSlot.toString(),
    preState: proof.preState,
    postState: proof.postState,
    journalData: proof.journalData,
    seal: proof.seal,
  };

  const boundlessPath = path.join(__dirname, "testdata", "boundless_fixture.json");
  fs.mkdirSync(path.dirname(boundlessPath), { recursive: true });
  fs.writeFileSync(boundlessPath, JSON.stringify(boundlessFixture, null, 2));
  console.log(`Boundless fixture written to ${boundlessPath}`);
  console.log(`\nANCHOR_SLOT=${proof.finalizedSlot}`);
  console.log(`TX_HASH=${txHash}`);
}

main().catch((err) => {
  console.error("\nError:", err.message || err);
  process.exit(1);
});
