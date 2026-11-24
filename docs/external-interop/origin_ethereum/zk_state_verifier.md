# Verifying ICM Messages with ZK Proofs

As described in [Authenticating ICM Messages](https://github.com/ava-labs/icm-services/blob/main/docs/external-interop/icm_message_authentication.md), any contract that implements the `IMessageVerifier` interface can be used by a `TeleporterMessengerV2` instance to authenticate inbound messages. The protocol is agnostic to how a message is authenticated as that responsibility is delegated entirely to the verifier.

The `ZKAdapter` is one such `IMessageVerifier` implementation. While the [`MerkleValidatorSetRegistry`](https://github.com/ava-labs/icm-services/blob/main/docs/external-interop/origin_avalanche/validator_set_registry.md) authenticates messages originating from Avalanche L1s on an external EVM chain by checking a quorum of BLS validator signatures, the `ZKAdapter` covers the opposite direction: it authenticates messages originating on Ethereum so that they can be consumed on Avalanche L1s. It does so by leveraging zero-knowledge (ZK) proofs of the source chain's consensus rather than trusting a signing committee. From the trusted consensus state, it proves execution-layer events using various Merkle proofs from the beacon state root down to a receipt log event.

The current supported implementation is [SP1 Helios](https://github.com/succinctlabs/sp1-helios) by Succinct, an open-source ZK Ethereum light client that proves, in a single SP1 proof, that the beacon chain's sync committee signed a new finalized beacon block header. The proof is checked on-chain via a deployed SP1 verifier.

In short, `ZKAdapter` lets an Avalanche chain trustlessly confirm that a given log/event was emitted on Ethereum, and treats that proven event as the attestation for an ICM message.

## Background: Ethereum beacon-chain terms

A few Ethereum consensus-layer concepts are used throughout this document:

* **Slot** — Ethereum's consensus layer, the "beacon chain", advances in fixed time intervals called slots. Each slot may contain one beacon block.
* **Finalized checkpoint** — periodically, validators *finalize* a checkpoint: a slot that the chain commits to irreversibly under Casper FFG. A finalized checkpoint is identified by its epoch and the beacon block root at that boundary.
* **Sync committee** — a set of 512 validators, rotated every sync committee period (8192 slots, about 27 hours), that signs the head of the chain. Light clients such as SP1 Helios follow finality by verifying these signatures, and learn each period's committee from the previous one.
* **Beacon state root** — each beacon block commits to a Merkle root over the entire consensus state (the "beacon state"). Through that state it also commits to the **execution state** which contains the transactions, receipts, and logs of the corresponding EVM block.

Because all of this is Merkle-committed, a single trusted beacon block root is enough to prove, via inclusion proofs, that a specific execution-layer log was emitted. That is exactly what this adapter does: it trusts a set of finalized beacon roots, established by ZK proof, and then proves individual logs against them.

## Architecture

`ZKAdapter` is composed of two layers:

* **`ZKStateManager`** — a ZK light client that tracks the finalized beacon-chain state of a single source chain and can prove that a specific execution-layer log was emitted on that chain.
* **`ZKAdapter`** — a wrapper that inherits `ZKStateManager` and implements the full [`IAdapter`](https://github.com/ava-labs/icm-services/blob/main/icm-contracts/common/ITeleporterMessengerV2.sol) interface, which implements`IMessageSender` and `IMessageVerifier`, adapting the message-emitting and log-proving entry points to the TeleporterV2 format.

```solidity
contract ZKAdapter is ZKStateManager, IAdapter { ... }
```

Because `ZKAdapter` implements the full `IAdapter`, the same contract can serve both halves of a connection: `sendMessage` originates a message on the source chain by emitting it as a log, and `verifyMessage` authenticates an inbound message on the destination chain by proving that log. For the Ethereum to Avalanche direction, an instance is deployed on each side at the same address (via Nick's method); the source-side instance emits, and the destination-side instance verifies. See the [Teleporter architecture](https://github.com/ava-labs/icm-services/blob/main/docs/external-interop/teleporter_contracts.md) for how a `TeleporterMessengerV2` is bound to an adapter.

## `ZKStateManager`

The `ZKStateManager` maintains a trusted view of the source chain's beacon chain. It is initialized in its constructor with:

* `sourceChainId` — the chain whose state it tracks,
* `startingState` — the initial `Consensus.State` used as the root of trust; its header is confirmed as the first anchor,
* `beaconConfig` — the `Execution.BeaconConfig` used to verify execution-layer data,
* `genesisTime` — the Unix timestamp of the beacon chain's genesis, used to convert slots to wall-clock time,
* `verifier` / `lightClientVkey` — the SP1 verifier contract and the verification key of the SP1 Helios light client program,
* `permissibleTimespan` — the maximum age, in seconds, of a transition: the new head must have begun no more than this long before the current block time, and a single transition may not span more chain time than this,
* `admin` / `superAdmin` — role holders for the privileged functions below.

The root of trust is a `Consensus.State`: the latest finalized head slot, its beacon block root, and the sync committee of its period, which verifies the next update.

```solidity
library Consensus {
    struct State {
        uint64 head;               // slot of the latest finalized beacon block
        bytes32 header;            // beacon block root at head
        bytes32 syncCommitteeHash; // sync committee hash of head's period
    }
}
```

Internally, the manager keeps three pieces of contract state that the flows below read and write:

* `_currentState` — a `Consensus.State`; the single, latest trusted consensus state. Each transition overwrites it.
* `syncCommittees` — a state mapping `period => syncCommitteeHash`. Each update can carry the next period's committee, which is what lets the light client cross into a new period.
* `_allowedBeaconBlocks` — a state mapping `slot => beaconBlockRoot` that accumulates every finalized beacon block root the contract has confirmed. Entries are added and never overwritten, so it grows into the full set of trusted anchors.

It exposes two core flows.

### 1. Advancing consensus state — `transition`

```solidity
function transition(ConsensusData calldata consensus) external;

struct ConsensusData {
    bytes publicValues; // ABI-encoded ProofOutputs (the proof's public values)
    bytes proof;        // the SP1 proof
}

// Public values of the SP1 Helios light client program (abridged).
struct ProofOutputs {
    bytes32 prevHeader;            // must match the contract's stored _currentState
    uint256 prevHead;
    bytes32 prevSyncCommitteeHash;
    uint256 newHead;               // the newly finalized checkpoint slot
    bytes32 newHeader;             // its beacon block root
    bytes32 syncCommitteeHash;     // sync committee of newHead's period
    bytes32 nextSyncCommitteeHash; // sync committee of the next period, if known
    ...
}
```

`transition` advances the tracked beacon state from the contract's current state to a new finalized header. It:

1. Decodes the `ProofOutputs` from `consensus.publicValues`.
2. Verifies the supplied SP1 proof (`consensus.proof`) against `lightClientVkey` and the public values, after first checking that the previous head, header, and sync committee match the contract's stored `_currentState`, that the head advances, and that the transition is within `permissibleTimespan` (its new head is recent and it does not span too much chain time). The recency check means a stale but otherwise valid proof cannot be replayed, for example after an admin has reset the state with `manualTransition`. A successful proof attests that the current sync committee signed `newHeader` as finalized.
3. Stores any newly learned sync committees, updates the contract's `_currentState` to the new head, and records its beacon block root in `_allowedBeaconBlocks`.

Because each update is verified by the sync committee of the current head's period, an update must land at least once per sync committee period (about 27 hours). Otherwise the contract cannot advance until an admin recovers it with `manualTransition`.

Each successful `transition` therefore extends the set of finalized beacon block roots the contract considers trustworthy. These roots are the anchors that later log proofs are checked against.

### 2. Proving a log — `verifyLogAndExtract`

```solidity
function verifyLogAndExtract(
    Execution.Proof calldata execProof,
    Receipt.Proof calldata logProof
) external returns (bytes memory logData);
```

The two proof structs it consumes trace a chain of trust from a beacon block root down to a single log. `Execution.Proof` links a receipts root back to a trusted beacon block root:

```solidity
struct Proof {
    uint64 anchorSlot;                    // slot whose trusted beacon root anchors the proof
    uint64 targetSlot;                    // slot where the target transaction occurred
    bytes32 anchorBeaconStateRoot;        // trusted block -> anchor beacon state
    bytes32[] anchorBeaconStateProof;
    bytes32 targetBeaconStateRoot;        // anchor state -> target beacon state (history)
    bytes32[] targetBeaconStateProof;
    bytes32 targetExecutionHeaderRoot;    // target state -> execution header
    bytes32[] targetExecutionHeaderProof;
    bytes32 targetReceiptsRoot;           // execution header -> receipts root
    bytes32[] targetReceiptsProof;
}
```

`Receipt.Proof` proves a specific log lives inside that receipts root:

```solidity
struct Proof {
    bytes[] proof;            // Merkle-Patricia inclusion proof (RLP-encoded trie nodes)
    bytes key;               // RLP-encoded transaction index
    bytes value;             // RLP-encoded receipt
    uint256 logIndex;        // which log in the receipt
    address expectedEmitter; // contract that must have emitted the log
    bytes32 expectedTopic0;  // event signature the log must match
}
```

This is the entry point that proves a specific log was emitted on the source chain. It performs two verifications:

1. **Execution verification.** It looks up the trusted beacon block root for `execProof.anchorSlot` in the `_allowedBeaconBlocks` state mapping (reverting if that slot has not been confirmed). It then calls `Execution.verify` to link the claimed `targetReceiptsRoot` back to that anchor root, traversing from the beacon layer to the execution layer via Merkle inclusion proofs.
2. **Log verification.** It calls `Receipt.verifyAndExtractLog` against the now-verified `targetReceiptsRoot` to prove inclusion of the receipt and its log, and returns the raw, non-indexed `logData`.

On success it emits `ZKEventImported` and returns `logData` to the caller; the caller is responsible for decoding and interpreting that data.

Because step 1 requires the anchor slot's beacon root to already be present in `_allowedBeaconBlocks`, **`verifyLogAndExtract` can only succeed after the relevant finalized state has been synced via `transition`.** Keeping the state manager current is an operational responsibility of the relayer.

### Administrative functions

`ZKStateManager` also exposes role-gated maintenance functions: `updateLightClientVkey`, `updateVerifier`, and `updatePermissibleTimespan` (all `ADMIN_ROLE`), plus `manualTransition`, which applies a state transition *without* proof verification for emergency recovery, overwriting any stored sync committees. `getBeaconBlockRoot` is a public view helper for inspecting confirmed roots.

## `ZKAdapter`

`ZKAdapter` adapts the `ZKStateManager` flows to the `IAdapter` interface so it can be plugged into a `TeleporterMessengerV2`.

### `sendMessage`

```solidity
function sendMessage(TeleporterMessageV2 calldata message) external {
    require(msg.sender == message.originTeleporterAddress, "unauthorized sender");
    emit TeleporterV2MessageSent(abi.encode(message));
}
```

On the source chain, `sendMessage` originates a message by emitting it as a `TeleporterV2MessageSent` log that the relayer later proves on the destination chain. The `require` enforces that the caller is the originating Teleporter messenger: the messenger sets `originTeleporterAddress` to its own address (`address(this)`) before calling the adapter, so this check guarantees only the genuine messenger can emit a message. Without it, anyone could call `sendMessage` directly with an attacker-authored message and produce a log that the destination side would accept.

### Attestation format

The `attestation` field of the `TeleporterICMMessage` is a generic `bytes` blob that must ABI-encode the following struct:

```solidity
struct Attestation {
    Execution.Proof execProof;
    Receipt.Proof logProof;
}
```

These are exactly the two proofs `verifyLogAndExtract` consumes: the execution proof anchoring a receipts root to a trusted beacon block, and the log proof establishing the specific log within that receipts root.

### `verifyMessage`

```solidity
function verifyMessage(
    TeleporterICMMessage calldata message
) external returns (bool) {
    Attestation memory att = abi.decode(message.attestation, (Attestation));
    require(message.sourceBlockchainID == bytes32(sourceChainId), "bad source chain");
    require(message.sourceNetworkID == sourceNetworkId, "bad network");

    require(att.logProof.expectedEmitter == address(this), "bad emitter");
    require(att.logProof.expectedTopic0 == _MESSAGE_SENT_TOPIC, "bad topic");

    bytes memory logData = this.verifyLogAndExtract(att.execProof, att.logProof);
    bytes memory emitted = abi.decode(logData, (bytes));
    require(keccak256(emitted) == keccak256(abi.encode(message.message)), "payload mismatch");

    return true;
}
```

`verifyMessage` confirms the message really came from the source chain and isn't something the relayer made up. The relayer supplies the proof, so the adapter can't trust the fields in it; instead it checks each one against a value it already knows. It confirms the message names the chain and network this adapter tracks, that the proven log was emitted by the trusted adapter (the same contract address on both chains, via Nick's method), and that the log is a `TeleporterV2MessageSent` event. It then runs `verifyLogAndExtract` to check the proof against the synced beacon state. Finally, it checks that the log's contents match the message being verified — the step that ties the proof to *this* message, since otherwise a real proof of some unrelated event would pass. If any check fails the call reverts, and that revert propagates out of `receiveCrossChainMessage`.

## Verification flow

```mermaid
sequenceDiagram
autonumber
actor relayer
note left of relayer: Keep source-chain state synced
relayer->>ZKStateManager: transition(consensus)
ZKStateManager->>SP1Verifier: verifyProof(lightClientVkey, publicValues, proof)
SP1Verifier-->>ZKStateManager: ok / revert
note over ZKStateManager: record finalized beacon block root

note left of relayer: Deliver ICM message
relayer->>TeleporterMessengerV2: receiveCrossChainMessage(...)
TeleporterMessengerV2->>ZKAdapter: verifyMessage(message)
ZKAdapter->>ZKAdapter: check source chain, network, emitter, topic
ZKAdapter->>ZKAdapter: this.verifyLogAndExtract(execProof, logProof)
ZKAdapter->>ZKAdapter: bind proven log payload to message
ZKAdapter-->>TeleporterMessengerV2: true / revert
```
