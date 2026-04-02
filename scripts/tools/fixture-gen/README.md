# fixture-gen

Generates a JSON test fixture for ZKAdapter e2e tests by fetching real SSZ Merkle proofs
from a beacon node and MPT receipt proofs from an execution layer client.

Note that to mitigate supply-chain attack risk, it is recommended to run these scripts and npm dependencies in an isolated AWS environment.

## Prerequisites
- Node.js v24+ (e.g., https://nodejs.org/en/download)
- A Sepolia execution layer RPC endpoint (e.g., Infura) 
- A Sepolia beacon node API endpoint with access to recent state history (e.g., QuickNode)

**Note:** The beacon node must have the state for the target slot available. Standard nodes
only retain states for the last few epochs (~50 slots). Use a recent transaction to stay
within this window, or use an archive beacon node for any historical transaction.

## Setup
```bash
npm install
```

## Usage

Set the required environment variables, then compile and run:
```bash
export BEACON_API_URL=https://your-beacon-node
export ETH_RPC_URL=https://your-eth-rpc
export TX_HASH=0x...

npx tsc
node dist/generate_fixtures.mjs
```

## Output

The fixture is written to `dist/sepolia_fixture.json`. Move it to `tests/testdata/sepolia_fixture.json` before running the e2e tests.

- `anchorBeaconBlockRoot` — the beacon block root used as the trusted anchor
- `metadata` — transaction and slot information for reference (this data is not used by the e2e test)
- `executionProof` — SSZ Merkle proofs linking the beacon state to the execution payload
- `receiptProof` — MPT proof for the target transaction receipt
