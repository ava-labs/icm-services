# fixture-gen

Generates JSON test fixtures for ZKAdapter e2e tests. One fixture is needed:

1. **Ethereum fixture** — SSZ Merkle proofs and MPT receipt proofs fetched from a beacon node and execution layer client

Note that to mitigate supply-chain attack risk, it is recommended to run these scripts and npm dependencies in an isolated AWS environment.

## Prerequisites

- Node.js v24+ (https://nodejs.org/en/download)
- An Ethereum mainnet execution layer RPC endpoint (e.g., Infura)
- An Ethereum mainnet beacon node API endpoint with access to recent state history and
  the `eth/v2/debug/beacon/states` endpoint (e.g., QuickNode)

**Note:** The beacon node must have the state for the target slot available. Standard nodes
only retain states for the last few epochs (~50 slots). Use a recent transaction to stay
within this window, or use an archive beacon node for any historical transaction. Mainnet
beacon states are large; increase Node's heap via `--max-old-space-size` if needed.

## Setup
```bash
npm install
```
## Background

The ZKAdapter verifies Ethereum events on Avalanche in two stages: an SP1 Helios ZK proof
establishes a trusted finalized beacon block root on-chain (the anchor), then Merkle proofs
link that anchor to a specific transaction's event log.

The e2e test does not yet submit an SP1 Helios proof. Instead it deploys the ZKAdapter with the
fixture's anchor as its trusted starting state, which the constructor confirms as an anchor.

Key terms:

- `anchorSlot` — the slot of the trusted anchor, a finalized checkpoint after the transaction
- `anchorBeaconBlockRoot` — the beacon block root at `anchorSlot`
- `executionProof` — SSZ Merkle proofs linking the anchor to the target slot's receipts root
- `receiptProof` — MPT proof that the transaction receipt (and its event log) is in the receipts trie

## Choosing the anchor slot

The anchor must be a finalized checkpoint after the transaction's slot, within 8192 slots of
it. To wait for one, set the required environment variables and run the finality script:

```bash
export BEACON_API_URL=...  # Ethereum mainnet beacon API
export ETH_RPC_URL=...     # Ethereum mainnet execution RPC
export TX_HASH=0x...       # Transaction the anchor must cover

node wait_for_finality.mts  # writes to testdata/anchor.json
```

## Generating the Ethereum fixture

The `targetSlot` (the transaction's slot) must be below the `anchorSlot`, within 8192 slots.

Set `ANCHOR_SLOT` from the anchor, then run:

```bash
export BEACON_API_URL=...
export ETH_RPC_URL=...
export TX_HASH=0x...      # Same transaction used for the anchor
export ANCHOR_SLOT=$(jq -r '.anchorSlot' testdata/anchor.json)

NODE_OPTIONS="--max-old-space-size=8192" node generate_fixture.mts
```

## Running the e2e tests

The Go e2e test loads the fixture from the path given by the `ETHEREUM_FIXTURE_PATH`
environment variable, defaulting to `./tests/testdata/ethereum_fixture.json` (relative to the
repo root).

Either copy the generated fixture to that default location:

```bash
cp testdata/ethereum_fixture.json ../../../tests/testdata/
```

or point the variable at it directly:

```bash
cd ../../..   # or the repo root
ETHEREUM_FIXTURE_PATH=scripts/tools/fixture-gen/testdata/ethereum_fixture.json \
GINKGO_FOCUS="ZKAdapter" ./scripts/e2e_test.sh 
```
