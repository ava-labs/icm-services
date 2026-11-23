/**
 * Given an Ethereum mainnet tx hash, polls the beacon API until a finalized checkpoint is after
 * the tx's slot, then writes that checkpoint's slot as the anchor slot for the Ethereum fixture.
 *
 * Required env vars:
 *   BEACON_API_URL - Ethereum beacon API endpoint
 *   ETH_RPC_URL    - Ethereum execution layer RPC
 *   TX_HASH        - Ethereum tx hash to anchor
 *
 * Output:
 *   testdata/anchor.json
 */

import { ethers } from "ethers";
import * as fs from "fs";
import * as path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const BEACON_API_URL = process.env.BEACON_API_URL;
const ETH_RPC_URL = process.env.ETH_RPC_URL;

if (!BEACON_API_URL || !ETH_RPC_URL) {
  console.error("Required env vars: BEACON_API_URL, ETH_RPC_URL");
  process.exit(1);
}

const BEACON_GENESIS_TIME = 1606824023;
const SECONDS_PER_SLOT = 12;
// An anchor proves slots through its state_roots vector, which reaches back this many slots.
const STATE_ROOTS_VECTOR_SIZE = 8192;

// How long to wait for finality before giving up (1 hour). Finality normally takes ~13 minutes.
const MAX_WAIT_MS = 60 * 60 * 1000;
// How often to poll the beacon API (1 minute)
const POLL_INTERVAL_MS = 60 * 1000;

// Resolve tx hash from env var or from tx_info.json (written by send_ethereum_message.mts)
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
  return Math.floor((timestamp - BEACON_GENESIS_TIME) / SECONDS_PER_SLOT);
}

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function beaconGet(route: string): Promise<any> {
  const res = await fetch(`${BEACON_API_URL}${route}`);
  if (!res.ok) throw new Error(`GET ${route} failed: ${res.status} ${res.statusText}`);
  return res.json();
}

// Returns the slot of the latest finalized checkpoint's block. The block's slot can precede its
// epoch's first slot when that slot was skipped, so it is read from the block header.
async function finalizedCheckpointSlot(): Promise<number> {
  const checkpoints = await beaconGet("/eth/v1/beacon/states/head/finality_checkpoints");
  const root = checkpoints.data.finalized.root;
  const header = await beaconGet(`/eth/v1/beacon/headers/${root}`);
  return parseInt(header.data.header.message.slot);
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
  console.log(`tx is in block ${receipt.blockNumber}, slot ${txSlot}\n`);

  // The anchor must be after the tx's slot so its state_roots vector can reach it.
  console.log("=== Step 2: Waiting for a finalized checkpoint after the tx ===");
  const startTime = Date.now();
  let anchorSlot = await finalizedCheckpointSlot();
  while (anchorSlot <= txSlot) {
    const elapsed = Date.now() - startTime;
    if (elapsed > MAX_WAIT_MS) {
      throw new Error(`Timed out after ${MAX_WAIT_MS / 60000} minutes waiting for finality`);
    }
    const minutesLeft = ((MAX_WAIT_MS - elapsed) / 60000).toFixed(1);
    console.log(
      `Finalized slot ${anchorSlot} is not after tx slot ${txSlot}. ` +
        `Retrying in ${POLL_INTERVAL_MS / 1000}s (${minutesLeft}min remaining)...`
    );
    await sleep(POLL_INTERVAL_MS);
    anchorSlot = await finalizedCheckpointSlot();
  }

  // Verify the tx is within range
  const slotDiff = anchorSlot - txSlot;
  if (slotDiff > STATE_ROOTS_VECTOR_SIZE) {
    throw new Error(
      `Tx slot ${txSlot} is too far from anchor slot ${anchorSlot} (diff: ${slotDiff}, max: ${STATE_ROOTS_VECTOR_SIZE})`
    );
  }
  console.log(`Anchor slot: ${anchorSlot}, Tx slot: ${txSlot}, diff: ${slotDiff} (within ${STATE_ROOTS_VECTOR_SIZE} limit)\n`);

  // Save the anchor
  console.log("=== Step 3: Writing anchor ===");
  const anchorPath = path.join(__dirname, "testdata", "anchor.json");
  fs.mkdirSync(path.dirname(anchorPath), { recursive: true });
  fs.writeFileSync(anchorPath, JSON.stringify({ anchorSlot }, null, 2));
  console.log(`Anchor written to ${anchorPath}`);
  console.log(`\nANCHOR_SLOT=${anchorSlot}`);
  console.log(`TX_HASH=${txHash}`);
}

main().catch((err) => {
  console.error("\nError:", err.message || err);
  process.exit(1);
});
