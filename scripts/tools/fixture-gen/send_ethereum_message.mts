/**
 * Sends a TeleporterMessageV2 via an ECDSAVerifier contract on Ethereum.
 * The contract must already be deployed and initialized.
 * 
 * Required env vars:
 *   ETH_RPC_URL        - Ethereum execution layer RPC
 *   SENDER_PRIVATE_KEY - Private key for the Ethereum account
 *   SENDER_CONTRACT    - Address of the deployed ECDSAVerifier on Ethereum
 *
 * Output:
 *   testdata/tx_info.json — TX hash, log index, emitter, and event topic
 */

import { ethers } from "ethers";
import * as fs from "fs";
import * as path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const ETH_RPC_URL = process.env.ETH_RPC_URL;
const SENDER_PRIVATE_KEY = process.env.SENDER_PRIVATE_KEY;
const SENDER_CONTRACT = process.env.SENDER_CONTRACT;
const MAX_FEE_GWEI = process.env.MAX_FEE_GWEI ?? "30";

if (!ETH_RPC_URL || !SENDER_PRIVATE_KEY || !SENDER_CONTRACT) {
  console.error("Required env vars: ETH_RPC_URL, SENDER_PRIVATE_KEY, SENDER_CONTRACT");
  process.exit(1);
}

// Generated from ABI Go bindings for ECDSAVerifier
// TODO: Swap out with DiffUpdater ABI once sendMessage implementation is complete. Issue: https://github.com/ava-labs/icm-services/issues/1282
// NOTE: If ECDSAVerifier.sendMessage's signature changes, update this ABI string to match the new function signature.
const ECDSA_VERIFIER_ABI = [
  "function sendMessage(tuple(uint256 messageNonce, address originSenderAddress, address originTeleporterAddress, bytes32 destinationBlockchainID, address destinationAddress, uint256 requiredGasLimit, address[] allowedRelayerAddresses, tuple(uint256 receivedMessageNonce, address relayerRewardAddress)[] receipts, bytes message) message) external",
];

// Generated from ABI Go bindings for ECDSAVerifier
// ECDSAVerifierSendMessage event topic. This is the keccak256 of the ECDSAVerifierSendMessage event signature. 
// NOTE: If the contract is regenerated and this hash changes, update here too.
const EVENT_TOPIC = "0x7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8";

async function main() {
  const provider = new ethers.providers.JsonRpcProvider(ETH_RPC_URL);
  const wallet = new ethers.Wallet(SENDER_PRIVATE_KEY, provider);
  const contract = new ethers.Contract(SENDER_CONTRACT, ECDSA_VERIFIER_ABI, wallet);
  const balance = await wallet.getBalance();

   // Check the network
  const network = await provider.getNetwork();
  if (network.chainId !== 1) {
    throw new Error(`Expected mainnet (chainId 1), got chainId ${network.chainId}`);
  }

  console.log(`Sender balance: ${ethers.utils.formatEther(balance)} ETH`);

  // Set a max fee for the transaction
  const feeData = await provider.getFeeData();
  const maxFee = ethers.utils.parseUnits(MAX_FEE_GWEI, "gwei");
  if (feeData.maxFeePerGas && feeData.maxFeePerGas.gt(maxFee)) {
    throw new Error(
      `Current maxFeePerGas ${ethers.utils.formatUnits(feeData.maxFeePerGas, "gwei")} gwei exceeds cap ${MAX_FEE_GWEI} gwei; skipping run`
    );
  }

  console.log("=== Sending message on Ethereum ===");
  console.log(`Sender:   ${wallet.address}`);
  console.log(`Contract: ${SENDER_CONTRACT}`);

  const message = {
    messageNonce: 1,
    originSenderAddress: wallet.address,
    originTeleporterAddress: wallet.address,
    destinationBlockchainID: ethers.utils.hexZeroPad("0x01", 32),
    destinationAddress: wallet.address,
    requiredGasLimit: 100000,
    allowedRelayerAddresses: [],
    receipts: [],
    message: ethers.utils.hexlify(ethers.utils.toUtf8Bytes("hello from ethereum")),
  };

  const tx = await contract.sendMessage(message, {
    maxFeePerGas: maxFee,
    maxPriorityFeePerGas: ethers.utils.parseUnits("1", "gwei"),
  });
  console.log(`Tx submitted: ${tx.hash}`);
  console.log("Waiting for confirmation...");

  const receipt = await tx.wait();
  console.log(`Tx confirmed in block ${receipt.blockNumber}`);
  console.log(`Gas used: ${receipt.gasUsed.toString()}`);
  console.log(`Logs emitted: ${receipt.logs.length}`);

  // Find the ECDSAVerifierSendMessage event log index
  const logIndex = receipt.logs.findIndex(
    (log: ethers.providers.Log) =>
      log.address.toLowerCase() === SENDER_CONTRACT.toLowerCase() &&
      log.topics[0] === EVENT_TOPIC
  );

  if (logIndex === -1) {
    throw new Error("ECDSAVerifierSendMessage event not found in receipt. This will fail if EVENT_TOPIC has changed");
  }
  console.log(`ECDSAVerifierSendMessage event at log index: ${logIndex}`);

  // Write out tx info
  const outputPath = path.join(__dirname, "testdata", "tx_info.json");
  fs.mkdirSync(path.dirname(outputPath), { recursive: true });
  fs.writeFileSync(outputPath, JSON.stringify({
    txHash: tx.hash,
    blockNumber: receipt.blockNumber,
    logIndex,
    emitter: SENDER_CONTRACT,
    eventTopic: EVENT_TOPIC,
  }, null, 2));

  console.log(`\nTx info written to ${outputPath}`);
  console.log(`TX_HASH=${tx.hash}`);
  console.log(`LOG_INDEX=${logIndex}`);
}

main().catch((err) => {
  console.error("\nError:", err.message || err);
  process.exit(1);
});
