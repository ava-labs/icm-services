# fixture-gen

Generates JSON test fixtures for ZKAdapter e2e tests. Two fixtures are needed:

1. **Ethereum fixture** — SSZ Merkle proofs and MPT receipt proofs fetched from a beacon node and execution layer client
2. **Boundless fixture** — a ZK consensus proof from the Boundless network (journalData, seal, preState, postState)

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

The ZKAdapter verifies Ethereum events on Avalanche in two stages: a Boundless ZK proof
establishes a trusted finalized beacon block root on-chain (the anchor), then Merkle proofs
link that anchor to a specific transaction's event log.

Key terms:

- `finalizedSlot` / `anchorSlot` — the slot of the trusted anchor; must match across both fixtures
- `journalData`, `seal` — the Boundless ZK proof, verified on-chain to store the anchor
- `executionProof` — SSZ Merkle proofs linking the anchor to the target slot's receipts root
- `receiptProof` — MPT proof that the transaction receipt (and its event log) is in the receipts trie

## Generating the Boundless fixture

Boundless ZK consensus proofs can be queried from the Signal Ethereum subgraph. To deploy your own subgraph instance, see https://github.com/austinabell/signal-ethereum-subgraph. Note the subgraph indexes the Boundless Marketplace contract on **Base**, where Signal Ethereum mainnet proofs are delivered.

The Signal ZK proof orders for Ethereum mainnet on Boundless can be seen here: https://explorer.boundless.network/base/requestors/0x734df7809c4ef94da037449c287166d114503198
The `SIGNAL_ETHEREUM_IMAGE_ID` should match the latest image ID on the explorer link above. 
This can be seen by clicking on a request, viewing its details, and copying the image ID. 
Currently, this is `0x0ccb3d146a7f64e78cc1d146acc26912138ea39bb79b4ca74423389d61b2c30e`.

To generate the fixture, set the required environment variables and run the polling script, which waits until a proof covering the target transaction's slot is available:

```bash
export SUBGRAPH_URL=...   # Signal Ethereum subgraph GraphQL endpoint
export ETH_RPC_URL=...    # Ethereum mainnet execution RPC
export TX_HASH=0x...      # Transaction the proof must cover

node poll_boundless_proofs.mts   # writes to testdata/boundless_fixture.json
```

The fixture must contain `preState`, `postState`, `journalData`, `seal`, and `finalizedSlot`. The `finalizedSlot` determines which beacon block root gets stored on-chain, and the Ethereum fixture must be generated against a slot within an appropriate range of this value.

Move the fixture to `tests/testdata/boundless_fixture.json` before running the e2e tests.

## Generating the Ethereum fixture

The Ethereum fixture must be aligned with the Boundless fixture: the `anchorSlot` must equal the `finalizedSlot` from the Boundless fixture, and the `targetSlot` (the transaction's slot) must be below the `anchorSlot`, within 8192 slots.

Set `ANCHOR_SLOT` from the Boundless fixture, then run:

```bash
export BEACON_API_URL=...
export ETH_RPC_URL=...
export TX_HASH=0x...      # Same transaction used for the Boundless fixture
export ANCHOR_SLOT=$(jq -r '.finalizedSlot' testdata/boundless_fixture.json)

NODE_OPTIONS="--max-old-space-size=8192" node generate_fixture.mts
```

## Running the e2e tests

The Go e2e test loads the fixtures from paths given by the `ETHEREUM_FIXTURE_PATH` and
`BOUNDLESS_FIXTURE_PATH` environment variables, defaulting to
`./tests/testdata/ethereum_fixture.json` and `./tests/testdata/boundless_fixture.json`
(relative to the repo root).

Either copy the generated fixtures to those default locations:

```bash
cp testdata/boundless_fixture.json ../../../tests/testdata/
cp testdata/ethereum_fixture.json ../../../tests/testdata/
```

or point the variables at them directly:

```bash
cd ../../..   # or the repo root
ETHEREUM_FIXTURE_PATH=scripts/tools/fixture-gen/testdata/ethereum_fixture.json \
BOUNDLESS_FIXTURE_PATH=scripts/tools/fixture-gen/testdata/boundless_fixture.json \
GINKGO_FOCUS="ZKAdapter" ./scripts/e2e_test.sh 
```

