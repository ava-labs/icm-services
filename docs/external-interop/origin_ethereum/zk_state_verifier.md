# Verifying ICM Messages with ZK Proofs

As described in [Authenticating ICM Messages](https://github.com/ava-labs/icm-services/blob/main/docs/external-interop/icm_message_authentication.md), any contract that implements the `IMessageVerifier` interface can be used by a `TeleporterMessengerV2` instance to authenticate inbound messages. The protocol is agnostic to how a message is authenticated; that responsibility is delegated entirely to the verifier.

The `ZKAdapter` is one such `IMessageVerifier` implementation. While [`AvalancheValidatorSetRegistry`](https://github.com/ava-labs/icm-services/blob/main/docs/external-interop/origin_avalanche/validator_set_registry.md) authenticates messages originating from Avalanche L1s on an external EVM chain by checking a quorum of BLS validator signatures, the `ZKAdapter` covers the opposite direction: it authenticates messages originating on Ethereum so that they can be consumed on Avalanche L1s. It does so by leveraging zero-knowledge (ZK) proofs of the source chain's consensus rather than trusting a signing committee. From the trusted consensus state, it proves execution-layer events using various Merkle proofs from the beacon state root down to a receipt log event. 

The current supported implementation is [The Signal](https://github.com/boundless-xyz/Signal-Ethereum) by Boundless, an open-source ZK consensus client that compresses Ethereum's finalized beacon-chain checkpoints into a single ZK proof that any chain or contract can verify directly, with the proof carried in the message's attestation and checked on-chain via a RISC Zero deployed verifier.

In short, `ZKAdapter` lets an Avalanche chain trustlessly confirm that a given log/event was emitted on Ethereum, and treats that proven event as the attestation for an ICM message.

## Architecture
 
`ZKAdapter` is composed of two layers:
 
* **`ZKStateManager`** — a ZK light client that tracks the finalized beacon-chain state of a single source chain and can prove that a specific execution-layer log was emitted on that chain.
* **`ZKAdapter`** — a wrapper that inherits `ZKStateManager` and implements the full [`IAdapter`](https://github.com/ava-labs/icm-services/blob/main/icm-contracts/common/ITeleporterMessengerV2.sol) interface (`IMessageSender` + `IMessageVerifier`), adapting the message-emitting and log-proving entry points to the TeleporterV2 format.
```solidity
contract ZKAdapter is ZKStateManager, IAdapter { ... }
```
 
Because `ZKAdapter` implements the full `IAdapter`, the same contract can serve both halves of a connection: `sendMessage` originates a message on the source chain by emitting it as a log, and `verifyMessage` authenticates an inbound message on the destination chain by proving that log. For the Ethereum to Avalanche direction, an instance is deployed on each side at the same address (via Nick's method); the source-side instance emits, and the destination-side instance verifies. See the [Teleporter architecture](teleporter_contracts.md) for how a `TeleporterMessengerV2` is bound to an adapter.
 
## The `ZKStateManager`
 
The `ZKStateManager` maintains a trusted view of the source chain's beacon chain. It is initialized in its constructor with:
 
* `sourceChainId` — the chain whose state it tracks,
* `startingState` — the initial `Consensus.State` used as the root of trust,
* `beaconConfig` — the `Execution.BeaconConfig` used to verify execution-layer data,
* `verifier` / `imageID` — the RISC Zero verifier contract and the program image ID of the consensus-transition circuit (the "Signal Ethereum" program),
* `permissibleTimespan` — a bound used to reject stale transitions,
* `admin` / `superAdmin` — role holders for the privileged functions below.
It exposes two core flows.
 
### 1. Advancing consensus state — `transition`
 
```solidity
function transition(ConsensusData calldata consensus) external;
```
 
`transition` advances the tracked beacon state from the contract's current state to a new finalized checkpoint. It:
 
1. Decodes the `Journal` (pre-state, post-state, finalized slot) from `consensus.journalData`.
2. Verifies the supplied RISC Zero proof (`consensus.seal`) against `imageID` and the journal hash, after first checking that `journal.preState` matches the contract's stored `_currentState` and that the transition is within `permissibleTimespan`. A successful proof attests that `journal.postState` follows from `journal.preState` under Ethereum's Casper FFG consensus rules.
3. Updates `_currentState` to the post-state and records the finalized beacon block root for the finalized slot in the `_allowedBeaconBlocks` mapping (`slot => beaconBlockRoot`).
Each successful `transition` therefore extends the set of finalized beacon block roots the contract considers trustworthy. These roots are the anchors that later log proofs are checked against.
 
### 2. Proving a log — `verifyLogAndExtract`
 
```solidity
function verifyLogAndExtract(
    Execution.Proof calldata execProof,
    Receipt.Proof calldata logProof
) external returns (bytes memory logData);
```
 
This is the entry point that proves a specific log was emitted on the source chain. It performs two verifications:
 
1. **Execution verification.** It looks up the trusted beacon block root for `execProof.anchorSlot` in `_allowedBeaconBlocks` (reverting if that slot has not been confirmed). It then calls `Execution.verify` to link the claimed `targetReceiptsRoot` back to that anchor root, traversing from the beacon layer to the execution layer via Merkle inclusion proofs.
2. **Log verification.** It calls `Receipt.verifyAndExtractLog` against the now-verified `targetReceiptsRoot` to prove inclusion of the receipt and its log, and returns the raw, non-indexed `logData`.
On success it emits `ZKEventImported` and returns `logData` to the caller; the caller is responsible for decoding and interpreting that data.
 
Because step 1 requires the anchor slot's beacon root to already be present, **`verifyLogAndExtract` can only succeed after the relevant finalized state has been synced via `transition`.** Keeping the state manager current is an operational responsibility of the relayer.
 
### Administrative functions
 
`ZKStateManager` also exposes role-gated maintenance functions: `updateImageID`, `updateVerifier`, and `updatePermissibleTimespan` (all `ADMIN_ROLE`), plus `manualTransition`, which applies a state transition *without* proof verification for emergency recovery. `getBeaconBlockRoot` is a public view helper for inspecting confirmed roots.
 
## The `ZKAdapter`
 
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
ZKStateManager->>RiscZeroVerifier: verify(seal, imageID, journalHash)
RiscZeroVerifier-->>ZKStateManager: ok / revert
note over ZKStateManager: record finalized beacon block root
 
note left of relayer: Deliver ICM message
relayer->>TeleporterMessengerV2: receiveCrossChainMessage(...)
TeleporterMessengerV2->>ZKAdapter: verifyMessage(message)
ZKAdapter->>ZKAdapter: check source chain, network, emitter, topic
ZKAdapter->>ZKAdapter: this.verifyLogAndExtract(execProof, logProof)
ZKAdapter->>ZKAdapter: bind proven log payload to message
ZKAdapter-->>TeleporterMessengerV2: true / revert
```
 



