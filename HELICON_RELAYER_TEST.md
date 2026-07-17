# Helicon Devnet ICM Relayer - C to C Smoke Test Checkpoint

**Date:** 2026-07-16
**Ticket:** "Confirm that relayer works with helicon devnet c-chain <-> c-chain teleporter messages"
**Status:** Relayer works. Messages sent while it is running are relayed. Messages sent while it is down get picked up via the catchup path once a live block triggers it. See §10.7.

---

## 1. Goal

Verify that the ICM relayer (`icm-services/relayer`) can pick up a Teleporter message emitted on the Helicon devnet C-Chain (network 76, EVM chainId 43117) and deliver it back to the same C-Chain. Single-chain self-loop is enough to exercise the whole pipeline: source event -> BLS aggregation -> destination delivery.

## 2. Environment

| Item | Value |
|---|---|
| Working dir | `/Users/bora.erinc/Desktop/avalanche` |
| Repos | `avalanchego/` (`main`), `icm-services/` (`master`) |
| Network ID | 76 (Helicon devnet) |
| C-Chain EVM chainId | 43117 (`0xa86d`) |
| C-Chain blockchain ID (cb58) | `2CpuZMeuT19nECGuqUo1oZveNFvsjXV7xbVapiaaqSPnTKuWzH` |
| C-Chain blockchain ID (hex) | `0x9e8c642a15436e41c7892fc96ce89d205fb83c5f052d46c8bbbba3957af2a97a` |
| Public RPC | `https://api.avax-dev.network/ext/bc/C/rpc` |
| Public WS | `wss://api.avax-dev.network/ext/bc/C/ws` |
| Primary Network subnet ID | `11111111111111111111111111111111LpoYY` (32 zero bytes cb58) |
| Faucet | https://build.avax.network/console/primary-network/devnet-faucet (`@avalabs.org` login) |

## 3. Environment variables used

Set once per shell session:

```bash
export RPC=https://api.avax-dev.network/ext/bc/C/rpc
export KEY=<DEPLOYER_PRIVATE_KEY>         # cast wallet new, fund via faucet
export DEPLOYER=0x287D0E3B4b3c01E764BBFA82F8ed0a9e678322da   # cast wallet address --private-key $KEY
```

## 4. Contracts deployed

All three deployed via plain `forge create` from `$DEPLOYER`, **not** via Nick's method. See §5.1 for why I abandoned the canonical path.

| Contract | Address | Deploy source |
|---|---|---|
| `TeleporterMessenger` | `0x9AC66D0172343FEd917147c1fA0c46a43cca2BAa` | `icm-contracts/avalanche/teleporter/TeleporterMessenger.sol` |
| `TeleporterRegistry` | `0x6084e475c26f05897d1f8d5e352560b2adc1f573` | `icm-contracts/avalanche/teleporter/registry/TeleporterRegistry.sol` |
| `TestMessenger` (fixture) | `0x86a9d25DaA4C217d01FfE63570D716aD86f0D9aE` | `icm-contracts/avalanche/teleporter/tests/TestMessenger.sol` |

Export:

```bash
export MESSENGER=0x9AC66D0172343FEd917147c1fA0c46a43cca2BAa
export REGISTRY=0x6084e475c26f05897d1f8d5e352560b2adc1f573
export TESTMSG=0x86a9d25DaA4C217d01FfE63570D716aD86f0D9aE
```

## 5. Helicon-specific RPC quirks encountered

These bit me in the rear a few times.

### 5.1 `state not available for pending block`

`cast send` and `forge create` default to reading nonce and estimating gas against the `pending` block tag. The public devnet RPC does not serve pending state. Any tx-building operation errors out with `-32000: state not available for pending block`.

**Workaround:** pre-fetch nonce against `latest` and hardcode gas params:

```bash
NONCE=$(cast nonce --rpc-url $RPC $DEPLOYER)
# then pass:  --legacy --nonce $NONCE --gas-price 30gwei --gas-limit <N>
```

Applies to every send/deploy in this doc.

### 5.2 `only replay-protected (EIP-155) transactions allowed over RPC`

The devnet RPC rejects pre-EIP-155 txs. Nick's method deployment for canonical `TeleporterMessenger` at `0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf` is inherently pre-EIP-155 (that's how the same address is produced on every chain). Two options:

- Run your own tracking node with that flag, **or**
- Skip Nick's method and deploy Teleporter at a non-canonical address (what I did, see §6).

For a self-loop C-Chain to C-Chain test this is fine: source Messenger == destination Messenger by construction, so the "same address on every chain" property holds. I chose this route.

### 5.3 `deploy_teleporter.sh` patched but not used

`icm-services/scripts/deploy_teleporter.sh` was patched to (a) skip the `cast estimate --create` pre-flight (hits §5.1) and (b) hardcode nonce/gas on the funding `cast send`. Patches are in place, script is functional up to `cast publish`, which then hits §5.2. Kept committed for future use if `allow-unprotected-txs` gets enabled.

## 6. Deployment commands

Each step depends on the previous. All from `/icm-services`.

### 6.1 One-time build

```bash
cd icm-contracts
forge build
cd ..
```

### 6.2 TeleporterMessenger

```bash
NONCE=$(cast nonce --rpc-url $RPC $DEPLOYER)

forge create \
  icm-contracts/avalanche/teleporter/TeleporterMessenger.sol:TeleporterMessenger \
  --rpc-url $RPC --private-key $KEY --broadcast \
  --legacy --nonce $NONCE --gas-price 30gwei --gas-limit 5000000

export MESSENGER=<Deployed to: address from output>
```

### 6.3 TeleporterRegistry

Struct: `ProtocolRegistryEntry { uint256 version; address protocolAddress; }`

```bash
NONCE=$(cast nonce --rpc-url $RPC $DEPLOYER)

forge create \
  icm-contracts/avalanche/teleporter/registry/TeleporterRegistry.sol:TeleporterRegistry \
  --rpc-url $RPC --private-key $KEY --broadcast \
  --legacy --nonce $NONCE --gas-price 30gwei --gas-limit 3000000 \
  --constructor-args "[(1,$MESSENGER)]"

export REGISTRY=<Deployed to>
```

### 6.4 TestMessenger

`TestMessenger` inherits from `TeleporterRegistryOwnableAppUpgradeable` so its constructor has the `initializer` modifier so a direct deploy also initializes state.

Constructor: `(address teleporterRegistryAddress, address teleporterManager, uint256 minTeleporterVersion)`

```bash
NONCE=$(cast nonce --rpc-url $RPC $DEPLOYER)

forge create \
  icm-contracts/avalanche/teleporter/tests/TestMessenger.sol:TestMessenger \
  --rpc-url $RPC --private-key $KEY --broadcast \
  --legacy --nonce $NONCE --gas-price 30gwei --gas-limit 4000000 \
  --constructor-args $REGISTRY $DEPLOYER 1

export TESTMSG=<Deployed to>
```

### 6.5 Post-deploy sanity check

```bash
cast code $MESSENGER --rpc-url $RPC | head -c 20
cast code $REGISTRY  --rpc-url $RPC | head -c 20
cast code $TESTMSG   --rpc-url $RPC | head -c 20

# Registry maps version 1 -> Messenger
cast call $REGISTRY "getAddressFromVersion(uint256)(address)" 1 --rpc-url $RPC
# -> $MESSENGER

# TestMessenger stores registry in ERC-7201 namespaced storage.
# There is NO public getter, so read the slot directly.
# Slot = keccak256(abi.encode(uint256(keccak256("teleporter.storage.TeleporterRegistryApp")) - 1)) & ~0xff
cast storage $TESTMSG \
  0xde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00 \
  --rpc-url $RPC
# Returns bytes32, last 20 bytes = $REGISTRY

cast call $TESTMSG "getMinTeleporterVersion()(uint256)" --rpc-url $RPC
# -> 1
```

## 7. Relayer setup

### 7.1 Config file

Location: `icm-services/helicon-relayer-config.json`

```json
{
  "log-level": "debug",
  "info-api":    { "base-url": "https://api.avax-dev.network" },
  "p-chain-api": { "base-url": "https://api.avax-dev.network" },
  "storage-location": "./icm-relayer-storage",
  "process-missed-blocks": false,
  "source-blockchains": [
    {
      "subnet-id": "11111111111111111111111111111111LpoYY",
      "blockchain-id": "2CpuZMeuT19nECGuqUo1oZveNFvsjXV7xbVapiaaqSPnTKuWzH",
      "rpc-endpoint": { "base-url": "https://api.avax-dev.network/ext/bc/C/rpc" },
      "ws-endpoint":  { "base-url": "wss://api.avax-dev.network/ext/bc/C/ws" },
      "message-contracts": {
        "0x9ac66d0172343fed917147c1fa0c46a43cca2baa": {
          "message-format": "teleporter",
          "settings": { "reward-address": "0xF4979a31f70cB1b3Ac4d202D41C3f5EDbfBc0716" }
        }
      }
    }
  ],
  "destination-blockchains": [
    {
      "subnet-id": "11111111111111111111111111111111LpoYY",
      "blockchain-id": "2CpuZMeuT19nECGuqUo1oZveNFvsjXV7xbVapiaaqSPnTKuWzH",
      "rpc-endpoint": { "base-url": "https://api.avax-dev.network/ext/bc/C/rpc" },
      "account-private-key": "<RELAYER_KEY, a fresh key from `cast wallet new`, funded via faucet>"
    }
  ]
}
```

Gotchas that cost time:
- **`info-api`/`p-chain-api`** take the **host root only** (`https://api.avax-dev.network`). The SDKs (`info.NewClient`, `platformvm.NewClient`) append `/ext/info` and `/ext/bc/P` internally.
- **`rpc-endpoint`/`ws-endpoint`** take the **full path** (`.../ext/bc/C/rpc`, `.../ext/bc/C/ws`).
- Source and destination `blockchain-id` are the **same** for a C to C loop.
- `reward-address` and the relayer's `account-private-key` should be for a **different EOA** than `$DEPLOYER`. The relayer manages its own nonce, so do not use the destination key for anything else.

### 7.2 Build and run

```bash
cd /Users/bora.erinc/Desktop/avalanche/icm-services
./scripts/build.sh
./build/icm-relayer --config-file ./helicon-relayer-config.json
```

## 8. Test procedure

### 8.1 Trigger a message

Get the C-Chain blockchain ID as `bytes32`:

```bash
export CCHAIN_BYTES32=$(cast call \
  0x0200000000000000000000000000000000000005 "getBlockchainID()(bytes32)" \
  --rpc-url $RPC)
```

Send a Teleporter message from `$TESTMSG` to itself:

```bash
NONCE=$(cast nonce --rpc-url $RPC $DEPLOYER)

cast send $TESTMSG \
  "sendMessage(bytes32,address,address,uint256,uint256,string)" \
  $CCHAIN_BYTES32 $TESTMSG 0x0000000000000000000000000000000000000000 0 300000 "hello helicon" \
  --private-key $KEY --rpc-url $RPC \
  --legacy --nonce $NONCE --gas-price 30gwei --gas-limit 500000
```

### 8.2 Verify source-side event

Source tx from run: `0x541484a24302cf219cc0cb7bf93df6a66558bb94067ce1caa144cb05c2e8369e`, block **1618** (`0x652`), status **1 (success)**.

Receipt logs (in order): `TestMessenger.SendMessage`, `TeleporterMessenger.BlockchainIDInitialized`, **`TeleporterMessenger.SendCrossChainMessage`**, `WarpMessenger.SendWarpMessage`. All four present, source side worked exactly as expected.

### 8.3 Expected relayer log lines (from source code)

| Log | Where | Meaning |
|---|---|---|
| `Listener initialized. Listening for messages to relay.` | `relayer/listener.go:97` | subscription live |
| `Processing block` (debug) | `relayer/message_coordinator.go:243` | a block was routed for processing |
| `Relaying message` (info) | `messages/teleporter/message_handler.go:327` | log matched, will attempt delivery |
| `Sending message to destination chain` | `.../message_handler.go:253` | signatures aggregated, submitting delivery tx |
| `Finished relaying message to destination chain` | `.../message_handler.go:374` | success |

## 9. What was actually observed

Relayer started at `11:10:05`, chain head captured as `startingHeight=1618`. Then, after the source tx landed at `11:12:37`:

```
"Processing historical logs" fromBlockHeight:1618 toBlockHeight:1617    // no-op range
"Finished processing historical logs" fromBlockHeight:1618 toBlockHeight:1617
"Processing block" blockNumber:1618
"Attempting to commit height less than or equal to the committed height. Skipping."
    stagingHeight:1618  committedHeight:1618
```

**Missing wereevery downstream log**  
 - No `Relaying message`
 - No `Sending message` 
 - No `Finished relaying`

The relayer processed block 1618 with an **empty logs slice**, marked it committed, moved on. The `SendCrossChainMessage` event was never seen.

The p2p-network debug lines (`failed to get primary network uptime: node is not a validator`, `reset ip tracker bloom filter`) were observed.

## 10. Diagnostic journey

Kept as a running log. Bottom line at the end.

### 10.1 Initial (wrong) hypothesis: the bloom shortcut is deterministically broken under SAE

The code at `icm-services/types/types.go:59` uses a shortcut:

```go
// Check if the block contains warp logs, and fetch them from the client if it does
if header.Bloom.Test(TeleporterPrecompileLogFilter[:]) {
    // ... FilterLogs ...
}
```

First hypothesis: under ACP-194 (Streaming Asynchronous Execution), block acceptance and execution are decoupled. If WS `newHeads` fires at acceptance-time, the delivered header would have an empty `Bloom` (bloom being an output of execution) and the shortcut would always skip log fetching.

### 10.2 Reading the SAE VM code disproved the timing part

Followed the chain:

- `avalanchego/vms/saevm/saexec/subscription.go:12-15`, `sendPostExecutionEvents(b, results)` sends the `ChainHeadEvent` **after** execution finishes (function name is on the nose).
- `avalanchego/vms/saevm/saexec/execution.go:309-321`, the send is the very last step after `b.MarkExecuted(...)`, gated by a comment explicitly ordering the sequence.
- `avalanchego/vms/saevm/sae/consensus.go:144-147`, the VM's `SubscribeChainHeadEvent` routes to the executor's post-execution feed:
  > *"SubscribeChainHeadEvent returns a new subscription for each ChainHeadEvent emitted after a block has been executed."*
- `avalanchego/vms/saevm/docs/invariants.md`, canonical mapping table:
  ```
  SAE      ->  rawdb
  Accepted ->  Canonical, HeadFast
  Executed ->  Head
  Settled  ->  Finalized
  ```
  API `"latest"` = last-executed block.

So `newHeads` under SAE fires **post-execution**, not at acceptance.

### 10.3 Experiments

**Experiment A (forcing catch-up path):** flipped `process-missed-blocks: true`, set `process-historical-blocks-from-height: 1617`, cleared `./icm-relayer-storage/`, restarted, and sent a wake-up self-transfer to trigger the historical catch-up. The catch-up path (`subscriber.go:107`, `processBlockRange`) does an **unconditional** `FilterLogs`, no bloom shortcut. Result: block 1618 reprocessed, `Relaying message` fired, delivery tx submitted, `TestMessenger.getCurrentMessage(...)` returned `"hello helicon"`. Confirmed the source-side event exists and can be picked up when the bloom shortcut is bypassed.

**Experiment B (steady-state live-path send):** sent `"hello helicon 2"` while the relayer was warm and healthy. Block 1624 arrived via live path (`Processing block blockNumber:1624` fired), but no `Unpacked warp message`, no `Relaying message`. Silent miss at the time.

**Experiment C (more steady-state live-path sends):** sent `"hello helicon 3"`, `"hello helicon 4"`, `"hello helicon 5"` from the same warm relayer. All picked up by the **live path**, `Unpacked warp message` -> `Relaying message` -> `Finished relaying message to destination chain`, and `TestMessenger.getCurrentMessage(...)` advanced accordingly.

### 10.4 What I know at this point

- `newHeads` under SAE fires post-execution and the header carries the correct bloom, so **hypothesis 10.1 is falsified**.
- The live path is **not persistently broken**, also falsified.
- The live path appears intermittent from the outside. Helicon 1 (cold-start) and helicon 2 (warm) were dropped by the live path. Helicon 3-5 (warm) succeeded from the same process without any config change or restart.
- The catch-up path (unconditional `FilterLogs`) has not been observed to fail.

At this point I still suspected the bloom shortcut at `types.go:59` was the culprit since it was the only meaningful behavioral difference between the two paths. §10.6 disproved that.

### 10.6 Instrumentation experiment

Added a debug log at `types.go:59` (unconditional, right before the bloom test):

```go
logger.Info("bloom debug",
    zap.Uint64("blockNumber", header.Number.Uint64()),
    zap.String("blockHash", header.Hash().Hex()),
    zap.String("bloom", common.Bytes2Hex(header.Bloom[:])),
    zap.String("topicWanted", common.Bytes2Hex(TeleporterPrecompileLogFilter)),
    zap.Bool("bloomTest", header.Bloom.Test(TeleporterPrecompileLogFilter[:])),
)
```

This only fires on the **live path** (`NewICMBlockInfo`, called by `blocksInfoFromHeaders`). The catch-up path (`processBlockRange` at `subscriber.go:107-124`) calls `FilterLogs` directly and constructs `ICMBlockInfo` objects manually, bypassing `NewICMBlockInfo` and therefore this log entirely.

Reproduced a clean miss. Restarted the relayer, immediately sent `"miss test 1"`, block landed at 1639, no live-path delivery. Sent `"miss test 2"` a bit later, block 1642, delivered normally. Then grepped the log:

- `bloom debug ... "blockNumber":1639` returned **no matches**. The live path never called `NewICMBlockInfo` for block 1639.
- `bloom debug ... "blockNumber":1642` returned one entry with `bloom` populated and `bloomTest:true`. Live path works normally two blocks later.
- Manual trace of block 1639's message: its `warpMessageID` first appears in the log *after* a `Processing historical blocks` range that begins at 1640. The catch-up path rescued it.

I re-ran the same experiment against a **local Helicon tracking node** (built from the `helicon-devnet` branch, WS at `ws://127.0.0.1:9650/ext/bc/C/ws`) and saw the same shape. The startup window belongs to the relayer's own subscription setup, not to a quirk of the shared public endpoint.

### 10.7 Conclusion

The relayer works. Two paths deliver messages, and between them they cover every case that matters.

- **Live path.** Once the WS subscription is fully established, blocks flow through `NewICMBlockInfo`, the bloom shortcut passes (verified populated + `bloomTest:true` for block 1642), and messages are relayed within seconds.
- **Catch-up path.** On startup, and on any subscription error or reconnect, the relayer runs `ProcessFromHeight` over the gap between its last processed block and the first delivered live block. This path calls `FilterLogs` unconditionally, no bloom shortcut. It picks up anything the live path did not see, including messages sent during the brief startup window.

The behavior I first read as "the relayer missed a message" was the relayer doing exactly what it is designed to do. If it is not up (or not yet done initializing) when a message is sent, that message sits on-chain until the next live block triggers catch-up, and then gets delivered. This is intended.

The bloom shortcut at `types.go:59` is not the bug. Confirmed innocent by the grep-2 result above.

## 11. Startup behavior worth knowing about

Nothing here needs fixing, but the pattern is worth documenting so the next person doesn't spend an afternoon chasing it the way I did.

- Send a message within a few seconds of `./build/icm-relayer` starting: the live path does not process it in real time. The message is picked up by the next catch-up cycle once a subsequent live block arrives.
- Send a message once the relayer is warm: delivered promptly through the live path.
- Restart the relayer with `process-missed-blocks: true` and any staged-up messages get processed via catch-up.

The catch-up trigger is any live block delivered after startup (`listener.go:181-186`), so on an idle chain nothing will move until the next tx of any kind lands, even a self-transfer.

## 12. Cleanup

- The `bloom debug` log added at `types/types.go:59` in §10.6 is a debugging aid. Revert or move to `Debug` level before any PR.
- `icm-services/scripts/deploy_teleporter.sh` was patched in §5.3 to work around the public endpoint quirks. If those quirks get fixed upstream (allow-unprotected-txs enabled, pending state served), the patch can be reverted.

---

## Appendix A - files modified in this session

- `icm-services/scripts/deploy_teleporter.sh`. Commented out `cast estimate --create` (§5.3). Added `--legacy --nonce --gas-price --gas-limit` workarounds to the funding `cast send`.
- `icm-services/helicon-relayer-config.json`. Created.
- `icm-services/local-helicon-relayer-config.json`. Created. Same as above but pointing at a local Helicon tracking node.
- `icm-services/types/types.go:59`. Added a temporary `bloom debug` info log (see §10.6). Revert before any PR.

## Appendix B - key log excerpts (verbatim)

Source tx receipt (abbreviated):
```
blockHash            0xa7804117bbfd2fb198a5d9884cf81bc11c3ece50cebab6a9c7ffce950e60b42c
blockNumber          0x652   (= 1618)
status               1 (success)
transactionHash      0x541484a24302cf219cc0cb7bf93df6a66558bb94067ce1caa144cb05c2e8369e
logs                 [SendMessage@TESTMSG, BlockchainIDInitialized@MESSENGER,
                      SendCrossChainMessage@MESSENGER, SendWarpMessage@0x0200...05]
```

Relayer log (abbreviated timeline):
```
11:10:05  Creating application relayers                       sourceBlockchainID=2CpuZMe...
11:10:05  processed-missed-blocks set to false, starting from chain head
11:10:05  Creating checkpoint manager                         startingHeight=1618
11:10:05  Self-signing message originating from primary network
11:10:05  Created application relayer / relayers
11:10:05  Initialization complete
11:10:05  Creating relayer                                    protocolAddress=0x9AC66D...
11:10:05  Listener initialized. Listening for messages to relay.
11:10:15  db put   latestProcessedBlock=1618
11:12:37  Processing historical logs                          fromBlockHeight=1618 toBlockHeight=1617
11:12:37  Finished processing historical logs                 fromBlockHeight=1618 toBlockHeight=1617
11:12:37  Processing block                                    blockNumber=1618
11:12:37  Attempting to commit height <= committed. Skipping.  stagingHeight=1618 committedHeight=1618
[silence, no Relaying message, no Sending message, no Finished]
```
