#!/usr/bin/env bash
set -euo pipefail

# Nightly ZK integration test
#
# Sends a message on Sepolia and polls the Boundless subgraph until a ZK consensus proof
# covering the transaction is available.
#
#
# Required env vars:
#   ETH_RPC_URL        - Sepolia execution layer RPC
#   SENDER_PRIVATE_KEY - Private key for the funded Sepolia account
#   SENDER_CONTRACT    - Address of the ECDSAVerifier on Sepolia
#   SUBGRAPH_URL       - Boundless subgraph GraphQL endpoint
#   BEACON_API_URL     - Sepolia beacon API endpoint 

: "${ETH_RPC_URL:?Set ETH_RPC_URL to a Sepolia execution RPC}"
: "${SENDER_PRIVATE_KEY:?Set SENDER_PRIVATE_KEY to a funded Sepolia private key}"
: "${SENDER_CONTRACT:?Set SENDER_CONTRACT to the ECDSAVerifier address on Sepolia}"
: "${SUBGRAPH_URL:?Set SUBGRAPH_URL to the Boundless subgraph endpoint}"
: "${BEACON_API_URL:?Set BEACON_API_URL to a Sepolia beacon API}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
FIXTURE_DIR="$SCRIPT_DIR/tools/fixture-gen"
TESTDATA="$ROOT_DIR/tests/testdata/nightly"
mkdir -p "$TESTDATA"

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
  -v "$TESTDATA:/app/testdata" \
  fixture-gen node send_sepolia_message.mts

TX_HASH=$(jq -r '.txHash' "$TESTDATA/tx_info.json")
echo "Transaction hash: $TX_HASH"
echo ""

# Wait to index ZK proof via the Boundless subgraph 
echo "Waiting to index the Boundless ZK proof"
docker run --rm \
  -e SUBGRAPH_URL="$SUBGRAPH_URL" \
  -e ETH_RPC_URL="$ETH_RPC_URL" \
  -e TX_HASH="$TX_HASH" \
  -v "$TESTDATA:/app/testdata" \
  fixture-gen node poll_boundless_proofs.mts # writes nightly/boundless_fixture.json 

ANCHOR_SLOT=$(jq -r '.finalizedSlot' "$TESTDATA/boundless_fixture.json")
echo ""
echo " ******************************"
echo "  Boundless ZK proof indexed successfully"
echo "  TX hash:     $TX_HASH"
echo "  Anchor slot: $ANCHOR_SLOT"
echo "  Fixtures:    $TESTDATA"
echo " ******************************"

# Generate Ethereum SSZ/MPT proofs against the finalized beacon state for receipt log inclusion
echo "=== Step 4: Generating Ethereum receipt log inclusion proofs ==="
LOG_INDEX=$(jq -r '.logIndex' "$TESTDATA/tx_info.json")
docker run --rm \
  -e BEACON_API_URL="$BEACON_API_URL" \
  -e ETH_RPC_URL="$ETH_RPC_URL" \
  -e TX_HASH="$TX_HASH" \
  -e LOG_INDEX="$LOG_INDEX" \
  -e ANCHOR_SLOT="$ANCHOR_SLOT" \
  -v "$TESTDATA:/app/testdata" \
  fixture-gen node --max-old-space-size=8192 generate_fixture.mts # writes nightly/sepolia_fixture.json 
echo ""

# Run Go e2e test
echo "=== Step 5: Running Go e2e test ==="
cd "$ROOT_DIR"
ETHEREUM_FIXTURE_PATH="$TESTDATA/sepolia_fixture.json" \
BOUNDLESS_FIXTURE_PATH="$TESTDATA/boundless_fixture.json" \
GINKGO_FOCUS="ZKAdapter" ./scripts/e2e_test.sh --components ethereum-icm-verification

echo ""
echo " ******************************"
echo "  Nightly ZK Integration Test PASSED!"
echo " ******************************"
