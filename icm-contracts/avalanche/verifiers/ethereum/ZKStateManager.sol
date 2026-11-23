// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Based on code from: https://github.com/boundless-xyz/boundless-transceiver/blob/fa9d4c67872ad6ea916e647da3cc971c8eeee209/src/BlockRootOracle.sol
// Original code licensed under Apache-2.0 (see external/LICENSE)
// Modified: Added execution layer and receipt proof verification, and replaced the RISC Zero
// consensus proof with an SP1 Helios light client proof (https://github.com/succinctlabs/sp1-helios)

pragma solidity ^0.8.30;

import {AccessControl} from "@openzeppelin/contracts@5.1.0/access/AccessControl.sol";
import {SafeCast} from "@openzeppelin/contracts@5.1.0/utils/math/SafeCast.sol";
import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";
import {Consensus, Execution, Receipt} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Contains the data required to verify and perform a beacon chain state transition.
 * @dev This struct bundles the SP1 proof with the public values it commits to, which attest that
 * the beacon chain's sync committee signed a new finalized header.
 */
struct ConsensusData {
    /// @dev ABI-encoded ProofOutputs committed by the SP1 Helios light client program
    bytes publicValues;
    /// @dev The SP1 proof validating the light client update
    bytes proof;
}

/**
 * @notice A storage slot attested by the SP1 Helios program.
 * @dev Not used by this contract, but part of the ProofOutputs encoding.
 */
struct StorageSlot {
    bytes32 key;
    bytes32 value;
    address contractAddress;
}

/**
 * @notice The public values committed by the SP1 Helios light client program.
 * @dev Must match the program's output exactly, as the proof commits to its ABI encoding.
 * See ProofOutputs in https://github.com/succinctlabs/sp1-helios/blob/main/contracts/src/SP1Helios.sol
 */
struct ProofOutputs {
    /// @dev The previous beacon block header hash (beacon block root)
    bytes32 prevHeader;
    /// @dev The slot of the previous head
    uint256 prevHead;
    /// @dev The sync committee hash which was used to verify the proof
    bytes32 prevSyncCommitteeHash;
    /// @dev The slot of the new head
    uint256 newHead;
    /// @dev The new beacon block header hash (beacon block root)
    bytes32 newHeader;
    /// @dev The execution state root of the new beacon block. Unused by this contract.
    bytes32 executionStateRoot;
    /// @dev The execution block number of the new beacon block. Unused by this contract.
    uint256 executionBlockNumber;
    /// @dev The sync committee hash of the new head's period
    bytes32 syncCommitteeHash;
    /// @dev The sync committee hash of the next period, or zero if not yet known
    bytes32 nextSyncCommitteeHash;
    /// @dev Attested storage slots. Unused by this contract.
    StorageSlot[] storageSlots;
}

/**
 * @notice Information about an imported event from the beacon chain.
 */
struct ZKEventInfo {
    uint256 sourceChainId;
    uint256 beaconSlot;
    bytes32 executionRoot;
    uint256 logIndex;
    bytes logData;
}

contract ZKStateManager is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant UNDEFINED_ROOT = bytes32(0);

    /// @notice Number of slots per epoch in the beacon chain
    uint64 public constant SLOT_PER_EPOCH = 32;

    /// @notice Number of slots per sync committee period in the beacon chain
    uint64 public constant SLOTS_PER_PERIOD = 8192;

    /// @notice Seconds per slot on the Ethereum beacon chain
    uint64 public constant SECONDS_PER_SLOT = 12;

    uint256 public immutable sourceChainId;

    /**
     * @notice Unix timestamp of the tracked beacon chain's genesis, i.e. the start of slot 0
     * @dev Converts the slots carried in consensus states to wall-clock time for the
     * staleness check in {_permissibleTransition}.
     */
    uint256 public immutable genesisTime;

    /**
     * @notice The current consensus state of the beacon chain
     * @dev Updated atomically through state transitions to ensure consistency
     */
    Consensus.State private _currentState;

    /**
     * @notice The beacon chain config used to verify execution layer data.
     */
    Execution.BeaconConfig internal _beaconConfig;

    /**
     * @notice Maps a beacon chain slot to its verified beacon block root
     * @dev Used to track the beacon chain slots that are trusted and finalized through verified state transitions
     */
    mapping(uint64 slot => bytes32 beaconBlockRoot) internal _allowedBeaconBlocks;

    /**
     * @notice Maps a sync committee period to the hash of its sync committee
     * @dev Helios learns the next period's committee ahead of time, so an update can be verified
     * once the head crosses into a new period.
     */
    mapping(uint256 period => bytes32 syncCommitteeHash) public syncCommittees;

    /**
     * @notice The slot of the most recently confirmed beacon block
     * @dev Proof-driven transitions must confirm a higher slot than this (see transition()), so
     * consumers of ConfirmedBeaconBlock events can rely on anchors arriving in slot order. Admin
     * recovery via manualTransition may rewind it.
     */
    uint64 public lastConfirmedSlot;

    /**
     * @notice The SP1 verification key of the Helios light client program
     * @dev Normally it stays constant, but it can be updated if the light client program is upgraded, for example.
     * All proofs are verified against this program.
     */
    bytes32 public lightClientVkey;

    /**
     * @notice The address of the SP1 verifier contract
     * @dev Used to validate zero-knowledge proofs of beacon state transitions
     */
    address public verifier;

    /**
     * @notice Maximum age of a state transition, in seconds
     * @dev Bounds both how far behind the current block time a post-state's head may
     * lie and how much chain time a single transition may span. See {_permissibleTransition}.
     */
    uint24 public permissibleTimespan;

    /// @notice Events
    event Transitioned(
        uint64 indexed preHead,
        uint64 indexed postHead,
        Consensus.State preState,
        Consensus.State postState
    );
    event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root);
    event SyncCommitteeUpdated(uint256 indexed period, bytes32 indexed root);
    event LightClientVkeyUpdated(bytes32 indexed newVkey, bytes32 indexed oldVkey);
    event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan);
    event ZKEventImported(
        uint256 indexed sourceChainId,
        uint256 indexed beaconSlot,
        address indexed emitter,
        bytes32 executionRoot,
        uint256 logIndex
    );

    /// @notice Errors
    error InvalidArgument();
    error InvalidPreState();
    error PermissibleTimespanLapsed();
    error NonIncreasingSlot(uint64 slot, uint64 lastConfirmedSlot);
    error NonCheckpointSlot(uint64 slot);
    error NextSyncCommitteeMismatch(bytes32 given, bytes32 expected);

    /**
     * @notice Initializes the ZKStateManager contract with all required parameters
     * @dev Sets up the initial consensus state, configures verification parameters, and establishes cross-chain
     * communication
     * @param newSourceChainId The ID of the chain this contract will track
     * @param startingState The initial consensus state of the beacon chain
     * @param beaconConfig The beacon config used to verify execution layer data
     * @param genesisTime_ Unix timestamp of the beacon chain's genesis (start of slot 0)
     * @param permissibleTimespan_ Maximum age of a state transition in seconds
     * @param verifier_ Address of the SP1 verifier contract for proof validation
     * @param lightClientVkey_ The SP1 verification key of the Helios light client program
     * @param admin Address to be granted the ADMIN_ROLE
     * @param superAdmin Address to be granted the DEFAULT_ADMIN_ROLE
     */
    constructor(
        uint256 newSourceChainId,
        Consensus.State memory startingState,
        Execution.BeaconConfig memory beaconConfig,
        uint256 genesisTime_,
        uint24 permissibleTimespan_,
        address verifier_,
        bytes32 lightClientVkey_,
        address admin,
        address superAdmin
    ) {
        require(newSourceChainId != 0, "Invalid chain ID");
        sourceChainId = newSourceChainId;

        require(genesisTime_ != 0, "Invalid genesis time");
        genesisTime = genesisTime_;

        _grantRole(ADMIN_ROLE, admin);
        _grantRole(DEFAULT_ADMIN_ROLE, superAdmin);

        require(
            slotTimestamp(startingState.head) <= block.timestamp, "Starting state in the future"
        );
        _currentState = startingState;
        syncCommittees[_syncCommitteePeriod(startingState.head)] = startingState.syncCommitteeHash;
        // The starting header is trusted, so it can anchor messages before the first transition.
        _confirmBeaconBlock(startingState.head, startingState.header);
        _beaconConfig = beaconConfig;
        permissibleTimespan = permissibleTimespan_;
        verifier = verifier_;
        lightClientVkey = lightClientVkey_;
    }

    /**
     * @notice Performs a light client update from the trusted head to a new finalized header.
     * @dev The function performs the following steps:
     * 1. Verifies the SP1 proof that the current sync committee signed the new finalized header.
     * 2. Stores any newly learned sync committees.
     * 3. Updates the consensus state and confirms the new header as an anchor.
     */
    function transition(
        ConsensusData calldata consensus
    ) external {
        // Decode the public values.
        ProofOutputs memory outputs = abi.decode(consensus.publicValues, (ProofOutputs));
        uint64 newHead = SafeCast.toUint64(outputs.newHead);
        // Verify the light client update.
        _verify(outputs, newHead, consensus);
        // Proof-driven confirmations must move forward, but admin recovery via manualTransition may rewind.
        if (newHead <= lastConfirmedSlot) {
            revert NonIncreasingSlot(newHead, lastConfirmedSlot);
        }
        // Update the contract state to reflect the transition.
        _transition(outputs, newHead, false);
    }

    /**
     * @notice Verifies that a specific log was emitted on the beacon chain and returns the log.
     * @dev This is the main entry point for bridging events. It performs two key verifications:
     * 1. Execution Verification: Validates that the provided `targetReceiptsRoot` is part of the canonical beacon chain history using the `execProof`.
     * 2. Log Verification: Uses the validated `targetReceiptsRoot` to verify the inclusion  of a specific Receipt (and Log) via the `logProof`.
     * Once verified, it emits a `ZKEventImported` event.
     * @param execProof The execution proof linking a receipt root to a trusted beacon block root.
     * @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
     */
    function verifyLogAndExtract(
        Execution.Proof calldata execProof,
        Receipt.Proof calldata logProof
    ) external returns (bytes memory) {
        bytes32 anchorRoot = _allowedBeaconBlocks[execProof.anchorSlot];
        require(anchorRoot != UNDEFINED_ROOT, "Anchor slot is undefined");
        Execution.verify(anchorRoot, execProof, _beaconConfig);
        bytes memory logData = Receipt.verifyAndExtractLog(execProof.targetReceiptsRoot, logProof);
        emit ZKEventImported(
            sourceChainId,
            execProof.targetSlot,
            logProof.expectedEmitter,
            execProof.targetExecutionHeaderRoot,
            logProof.logIndex
        );
        return logData;
    }

    function updateLightClientVkey(
        bytes32 newLightClientVkey
    ) external onlyRole(ADMIN_ROLE) {
        if (newLightClientVkey == lightClientVkey) revert InvalidArgument();
        lightClientVkey = newLightClientVkey;
    }

    function updateVerifier(
        address newVerifier
    ) external onlyRole(ADMIN_ROLE) {
        require(newVerifier != address(0), "Invalid verifier address");
        if (newVerifier == verifier) revert InvalidArgument();
        verifier = newVerifier;
    }

    function updatePermissibleTimespan(
        uint24 newPermissibleTimespan
    ) external onlyRole(ADMIN_ROLE) {
        if (newPermissibleTimespan == permissibleTimespan) revert InvalidArgument();
        permissibleTimespan = newPermissibleTimespan;
    }

    /**
     * @notice Manually applies a beacon state transition without proof verification
     * @dev Admin-only function for emergency state updates, for example when no update landed
     * within a sync committee period. The previous values in `publicValues` are ignored, and the
     * sync committees it carries overwrite any stored ones.
     */
    function manualTransition(
        bytes calldata publicValues
    ) external onlyRole(ADMIN_ROLE) {
        ProofOutputs memory outputs = abi.decode(publicValues, (ProofOutputs));
        _transition(outputs, SafeCast.toUint64(outputs.newHead), true);
    }

    /**
     * @notice Outputs the beacon block root associated with the provided `slot`.
     * @param slot The beacon chain slot to look up
     */
    function getBeaconBlockRoot(
        uint64 slot
    ) external view returns (bytes32 root, bool valid) {
        root = _allowedBeaconBlocks[slot];
        valid = root != UNDEFINED_ROOT;
    }

    /**
     * @notice Returns the current trusted consensus state, which the previous values of the next
     * transition's public values must match.
     * @dev Lets off-chain updaters select the proof that extends this state. `lastConfirmedSlot`
     * alone cannot, as it is zero until the first transition.
     */
    function getCurrentState() external view returns (Consensus.State memory) {
        return _currentState;
    }

    /**
     * @notice Returns the wall-clock time at which `slot` begins on the tracked beacon chain.
     */
    function slotTimestamp(
        uint64 slot
    ) public view returns (uint256) {
        return genesisTime + uint256(slot) * SECONDS_PER_SLOT;
    }

    /**
     * @notice Transitions and updates the consensus state of the contract to the new head.
     * @dev Stores the sync committees of the new head's period and the next period if unset. A
     * conflicting next sync committee reverts, unless `overwrite` is set for admin recovery.
     */
    function _transition(
        ProofOutputs memory outputs,
        uint64 newHead,
        bool overwrite
    ) internal {
        // Helios only proves checkpoint slots, as CL nodes typically only serve proofs for those.
        if (newHead % SLOT_PER_EPOCH != 0) {
            revert NonCheckpointSlot(newHead);
        }

        uint256 period = _syncCommitteePeriod(newHead);
        if (overwrite || syncCommittees[period] == bytes32(0)) {
            _setSyncCommittee(period, outputs.syncCommitteeHash);
        }
        if (outputs.nextSyncCommitteeHash != bytes32(0)) {
            bytes32 storedNext = syncCommittees[period + 1];
            if (overwrite || storedNext == bytes32(0)) {
                _setSyncCommittee(period + 1, outputs.nextSyncCommitteeHash);
            } else if (storedNext != outputs.nextSyncCommitteeHash) {
                revert NextSyncCommitteeMismatch(outputs.nextSyncCommitteeHash, storedNext);
            }
        }

        Consensus.State memory preState = _currentState;
        Consensus.State memory postState = Consensus.State({
            head: newHead,
            header: outputs.newHeader,
            syncCommitteeHash: syncCommittees[period]
        });
        _currentState = postState;
        emit Transitioned(preState.head, postState.head, preState, postState);
        _confirmBeaconBlock(newHead, outputs.newHeader);
    }

    function _setSyncCommittee(
        uint256 period,
        bytes32 syncCommitteeHash
    ) internal {
        syncCommittees[period] = syncCommitteeHash;
        emit SyncCommitteeUpdated(period, syncCommitteeHash);
    }

    /**
     * @notice Confirms and stores a beacon block root for a given slot.
     * @dev Sets the slot high-water mark. Proof-driven transitions require it to increase (see
     * {transition}), while admin recovery via {manualTransition} may rewind it.
     */
    function _confirmBeaconBlock(
        uint64 slot,
        bytes32 root
    ) internal {
        lastConfirmedSlot = slot;
        _allowedBeaconBlocks[slot] = root;
        emit ConfirmedBeaconBlock(slot, root);
    }

    /**
     * @notice Verifies the light client update using the provided SP1 proof and public values.
     * @dev This function implements two key checks:
     * 1. The previous values of the public values match the current trusted state of the contract.
     * 2. The current sync committee signed the new finalized header. This step is verified by the
     * SP1 proof. Note: This function solely performs verification. The state update must be
     * handled by the caller.
     */
    function _verify(
        ProofOutputs memory outputs,
        uint64 newHead,
        ConsensusData calldata consensus
    ) internal view {
        // Ensure the proof is anchored to the current contract state.
        // The previous values in the public values must match the `_currentState` actually stored in this contract.
        Consensus.State memory preState = Consensus.State({
            head: SafeCast.toUint64(outputs.prevHead),
            header: outputs.prevHeader,
            syncCommitteeHash: outputs.prevSyncCommitteeHash
        });
        if (!Consensus.compareState(_currentState, preState)) {
            revert InvalidPreState();
        }
        // Ensure the transition is not stale.
        if (!_permissibleTransition(preState.head, newHead)) {
            revert PermissibleTimespanLapsed();
        }

        // Verify the proof against the verification key and the public values. This reverts if the proof is invalid.
        ISP1Verifier(verifier).verifyProof(lightClientVkey, consensus.publicValues, consensus.proof);
    }

    /**
     * @notice Checks that a transition advances the head, is recent, and does not span more chain
     * time than `permissibleTimespan`.
     * @dev Three conditions must hold:
     * 1. The post-state's head is later than the pre-state's, so the head only advances.
     * 2. The post-state's head began no more than `permissibleTimespan` seconds before
     *    the current block time, and not in the future. This is what makes a proof unusable once
     *    it is stale: a valid historical proof whose pre-state happens to match the contract's
     *    state again, for example after an admin recovery via {manualTransition}, is rejected
     *    rather than replayed to move the contract onto an outdated view of the chain.
     * 3. The chain time between the pre- and post-state heads is at most
     *    `permissibleTimespan`, so a single transition cannot jump the contract across a long gap.
     * Slots are converted to seconds with {slotTimestamp}.
     */
    function _permissibleTransition(
        uint64 preHead,
        uint64 postHead
    ) internal view returns (bool) {
        if (postHead <= preHead) {
            return false;
        }

        uint256 postTimestamp = slotTimestamp(postHead);
        if (postTimestamp > block.timestamp) {
            return false;
        }
        if (block.timestamp - postTimestamp > permissibleTimespan) {
            return false;
        }

        uint256 transitionTimespan = uint256(postHead - preHead) * SECONDS_PER_SLOT;
        return transitionTimespan <= permissibleTimespan;
    }

    function _syncCommitteePeriod(
        uint64 slot
    ) internal pure returns (uint256) {
        return slot / SLOTS_PER_PERIOD;
    }
}
