# fixture-gen

Generates JSON test fixtures for ZKAdapter e2e tests. Two fixtures are needed:

1. **Ethereum fixture** — SSZ Merkle proofs and MPT receipt proofs fetched from a beacon node and execution layer client
2. **Boundless fixture** — a ZK consensus proof from the Boundless network (journalData, seal, preState, postState)

Note that to mitigate supply-chain attack risk, it is recommended to run these scripts and npm dependencies in an isolated AWS environment.

## Prerequisites

- Node.js v24+ (https://nodejs.org/en/download)
- A Sepolia execution layer RPC endpoint (e.g., Infura)
- A Sepolia beacon node API endpoint with access to recent state history (e.g., QuickNode)

**Note:** The beacon node must have the state for the target slot available. Standard nodes
only retain states for the last few epochs (~50 slots). Use a recent transaction to stay
within this window, or use an archive beacon node for any historical transaction.

## Setup
```bash
npm install
```

## Generating the Boundless fixture

Boundless ZK consensus proofs can be queried from the Signal Ethereum subgraph. To deploy your own subgraph instance, see https://github.com/austinabell/signal-ethereum-subgraph.

The fixture must contain `preState`, `postState`, `journalData`, `seal`, and `finalizedSlot`. The `finalizedSlot` determines which beacon block root gets stored on-chain, and the Ethereum fixture must be generated against a slot within an appropriate range of this value. 

Move the fixture to `tests/testdata/boundless_fixture.json` before running the e2e tests.

## Generating the Ethereum fixture

The Ethereum fixture must be aligned with the Boundless fixture. The `anchorSlot` in the Ethereum fixture must equal the `finalizedSlot` from the Boundless fixture, and the `targetSlot` must be below the `anchorSlot` (within 8192 slots).

To achieve this, pick a transaction at a slot that is exactly 64 slots below the `finalizedSlot`. This way the script's default `anchorSlot = targetSlot + 64` will produce the correct anchor.

Set the required environment variables and run:
```bash
export BEACON_API_URL=...
export ETH_RPC_URL=...
export TX_HASH=0x...

NODE_OPTIONS="--max-old-space-size=8192" node generate_fixtures.mts
```

Move the fixture to `tests/testdata/ethereum_fixture.json` before running the e2e tests. The fixture output includes:

- `anchorBeaconBlockRoot` — the beacon block root used as the trusted anchor
- `metadata` — transaction and slot information for reference (not used by the e2e test)
- `executionProof` — SSZ Merkle proofs linking the beacon state to the execution payload
- `receiptProof` — MPT proof for the target transaction receipt
