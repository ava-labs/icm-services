#!/usr/bin/env bash
set -euo pipefail

# Nightly ZK integration test
#
# Sends a message on Sepolia and polls the Boundless subgraph until a ZK consensus proof
# covering the transaction is available.
#
# TODO: Add Ethereum fixture generation and run the full Go e2e test against a local Avalanche network
#
# Required env vars:
#   ETH_RPC_URL        - Sepolia execution layer RPC
#   SENDER_PRIVATE_KEY - Private key for the funded Sepolia account
#   SENDER_CONTRACT    - Address of the ECDSAVerifier on Sepolia
#   SUBGRAPH_URL       - Boundless subgraph GraphQL endpoint

: "${ETH_RPC_URL:?Set ETH_RPC_URL to a Sepolia execution RPC}"
: "${SENDER_PRIVATE_KEY:?Set SENDER_PRIVATE_KEY to a funded Sepolia private key}"
: "${SENDER_CONTRACT:?Set SENDER_CONTRACT to the ECDSAVerifier address on Sepolia}"
: "${SUBGRAPH_URL:?Set SUBGRAPH_URL to the Boundless subgraph endpoint}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
FIXTURE_DIR="$SCRIPT_DIR/tools/fixture-gen"
SHARED_TESTDATA="$ROOT_DIR/tests/testdata"

mkdir -p "$SHARED_TESTDATA"

echo " ******************************"
echo "  Nightly ZK Integration Test"
echo " ******************************"
echo ""

# Build Docker image
echo "Building Docker image." 
docker build -t fixture-gen "$FIXTURE_DIR"
echo ""

# Send message on Sepolia
echo " Sending message from ECDSAVerifier contract on Sepolia"
docker run --rm \
  -e ETH_RPC_URL="$ETH_RPC_URL" \
  -e SENDER_PRIVATE_KEY="$SENDER_PRIVATE_KEY" \
  -e SENDER_CONTRACT="$SENDER_CONTRACT" \
  -v "$SHARED_TESTDATA:/app/testdata" \
  fixture-gen node send_message.mts

TX_HASH=$(jq -r '.txHash' "$SHARED_TESTDATA/tx_info.json")
echo "Transaction hash: $TX_HASH"
echo ""

# Wait to index ZK proof via the Boundless subgraph 
echo "Waiting to index the Boundless ZK proof"
docker run --rm \
  -e SUBGRAPH_URL="$SUBGRAPH_URL" \
  -e ETH_RPC_URL="$ETH_RPC_URL" \
  -e TX_HASH="$TX_HASH" \
  -v "$SHARED_TESTDATA:/app/testdata" \
  fixture-gen node prepare_fixtures.mts

ANCHOR_SLOT=$(jq -r '.finalizedSlot' "$SHARED_TESTDATA/boundless_fixture.json")
echo ""
echo " ******************************"
echo "  Phase 1 Complete"
echo "  TX hash:     $TX_HASH"
echo "  Anchor slot: $ANCHOR_SLOT"
echo "  Fixtures:    $SHARED_TESTDATA"
echo " ******************************"
