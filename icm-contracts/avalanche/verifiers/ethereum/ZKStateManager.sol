// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Based on code from: https://github.com/boundless-xyz/boundless-transceiver/blob/fa9d4c67872ad6ea916e647da3cc971c8eeee209/src/BlockRootOracle.sol
// Original code licensed under Apache-2.0 (see external/LICENSE)
// Modified: Added execution layer and receipt proof verification

pragma solidity ^0.8.30;

import {AccessControl} from "@openzeppelin/contracts@5.0.2/access/AccessControl.sol";
import {IRiscZeroVerifier} from "./external/IRiscZeroVerifier.sol";
import {Consensus, Execution, Receipt} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// TODO: Add e2e testing for verifying Ethereum events on the C-chain. See https://github.com/ava-labs/icm-services/issues/1239

/**
 * @notice Contains the data required to verify and perform a beacon chain state transition.
 * @dev This struct bundles the ZK proof (seal) with the public inputs (journal) required to
 * verify that the Beacon Chain has successfully advanced to a new finalized state.
 */
struct ConsensusData {
    /// @dev Encoded Journal struct containing pre- and post-consensus states and finalized slot
    bytes journalData;
    /// @dev The RISC Zero ZK proof validating the beacon state transition
    bytes seal;
    /// @dev The beacon chain slot corresponding to the new finalized checkpoint
    uint64 finalizedSlot;
}

/**
 * @notice Contains the complete state transition data for ZK proof verification
 * @dev Used as journal data in RISC Zero proofs to validate beacon state transitions
 */
struct Journal {
    /// @dev The consensus state before the transition (must match the contract's stored `_currentState`).
    Consensus.State preState;
    /// @dev The consensus state after the transition
    Consensus.State postState;
    /// @dev The beacon chain slot that was finalized in this transition
    uint64 finalizedSlot;
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
    uint64 public constant SLOT_PER_EPOCH = 32; // TODO: Currently unused, but should be used in _permissibleTimespan

    uint256 public immutable sourceChainId;

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
     * @notice The RISC Zero program ID used to verify Ethereum consensus transitions
     * @dev This image ID corresponds to the ZK program (Signal Ethereum) that validates state transitions.
     * Normally it stays constant, but it can be updated if the consensus program is upgraded, for example.
     * All proofs are verified against this program.
     */
    bytes32 public imageID;

    /**
     * @notice The address of the RISC Zero verifier contract
     * @dev Used to validate zero-knowledge proofs of beacon state transitions
     */
    address public verifier;

    /**
     * @notice Maximum allowed time span for state transitions in seconds
     * @dev Used to prevent acceptance of stale beacon state transitions
     */
    uint24 public permissibleTimespan;

    /// @notice Events
    event Transitioned(
        uint64 indexed preEpoch,
        uint64 indexed postEpoch,
        Consensus.State preState,
        Consensus.State postState
    );
    event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root);
    event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID);
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

    /**
     * @notice Initializes the ZKStateManager contract with all required parameters
     * @dev Sets up the initial consensus state, configures verification parameters, and establishes cross-chain
     * communication
     * @param newSourceChainId The ID of the chain this contract will track
     * @param startingState The initial consensus state of the beacon chain
     * @param beaconConfig The beacon config used to verify execution layer data
     * @param permissibleTimespan_ Maximum allowed time span for state transitions in seconds
     * @param verifier_ Address of the RISC Zero verifier contract for proof validation
     * @param imageID_ The RISC Zero image ID for the beacon state transition program
     * @param admin Address to be granted the ADMIN_ROLE
     * @param superAdmin Address to be granted the DEFAULT_ADMIN_ROLE
     */
    constructor(
        uint256 newSourceChainId,
        Consensus.State memory startingState,
        Execution.BeaconConfig memory beaconConfig,
        uint24 permissibleTimespan_,
        address verifier_,
        bytes32 imageID_,
        address admin,
        address superAdmin
    ) {
        require(newSourceChainId != 0, "Invalid chain ID");
        sourceChainId = newSourceChainId;

        _grantRole(ADMIN_ROLE, admin);
        _grantRole(DEFAULT_ADMIN_ROLE, superAdmin);

        _currentState = startingState;
        _beaconConfig = beaconConfig;
        permissibleTimespan = permissibleTimespan_;
        verifier = verifier_;
        imageID = imageID_;
    }

    /**
     * @notice Performs a full state transition from a trusted beacon checkpoint to a new finalized state.
     * @dev The function performs the following steps:
     * 1. Verifies the ZK proof of the beacon (consensus) state transition to the new finalized checkpoint.
     * 2. Validates the Merkle chain of execution state roots from the prior epoch boundary to the new finalized execution state root.
     * 3. Verifies that all receipt roots are included in their respective execution state roots using Merkle proofs.
     * 4. Updates the global state to reflect the new trusted consensus, execution, and receipt roots.
     */
    function transition(
        ConsensusData calldata consensus
    ) external {
        // Decode the journal data.
        Journal memory journal = abi.decode(consensus.journalData, (Journal));
        // Verify the beacon state transition.
        _verify(journal, consensus);
        // Update the contract state to reflect the transition.
        _transition(journal, consensus.finalizedSlot);
    }

    /**
     * @notice Proves that a specific log was emitted on the beacon chain and executes application logic.
     * @dev This is the main entry point for bridging events. It performs two key verifications:
     * 1. Execution Verification: Validates that the provided `targetReceiptsRoot` is part of the canonical beacon chain history using the `execProof`.
     * 2. Log Verification: Uses the validated `targetReceiptsRoot` to verify the inclusion  of a specific Receipt (and Log) via the `logProof`.
     * Once verified, it emits a `ZKEventImported` event and passes the verified data to the internal `_onEventImport` handler for application-specific processing.
     * @param execProof The execution proof linking a receipt root to a trusted beacon block root.
     * @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
     */
    function proveLogAndExecute(
        Execution.Proof calldata execProof,
        Receipt.Proof calldata logProof
    ) external {
        // Verify the execution state proof against a stored beacon block root.
        bytes32 anchorRoot = _allowedBeaconBlocks[execProof.anchorSlot];
        require(anchorRoot != UNDEFINED_ROOT, "Anchor slot is undefined");
        Execution.verify(anchorRoot, execProof, _beaconConfig);

        // Verify the receipt and log, and extract the log data.
        bytes memory logData = Receipt.verifyAndExtractLog(execProof.targetReceiptsRoot, logProof);

        // Emit the event.
        emit ZKEventImported(
            sourceChainId,
            execProof.targetSlot,
            logProof.expectedEmitter,
            execProof.targetExecutionHeaderRoot,
            logProof.logIndex
        );

        // Hand off the event to application logic.
        _onEventImport(
            ZKEventInfo({
                sourceChainId: sourceChainId,
                beaconSlot: execProof.targetSlot,
                executionRoot: execProof.targetExecutionHeaderRoot,
                logIndex: logProof.logIndex,
                logData: logData
            })
        );
    }

    function updateImageID(
        bytes32 newImageID
    ) external onlyRole(ADMIN_ROLE) {
        if (newImageID == imageID) revert InvalidArgument();
        imageID = newImageID;
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
     * @dev Admin-only function for emergency state updates.
     */
    function manualTransition(
        bytes calldata journalData,
        uint64 finalizedSlot
    ) external onlyRole(ADMIN_ROLE) {
        Journal memory journal = abi.decode(journalData, (Journal));
        _transition(journal, finalizedSlot);
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
     * @notice Application-specific logic to handle the imported event.
     * @dev Override this in derived contracts to process verified cross-chain events.
     */
    function _onEventImport(
        ZKEventInfo memory eventInfo
    ) internal virtual 
    // solhint-disable-next-line no-empty-blocks
    {}

    /**
     * @notice Transitions and updates the consensus state of the contract to the new post-state.
     */
    function _transition(Journal memory journal, uint64 finalizedSlot) internal {
        _currentState = journal.postState;
        emit Transitioned(
            journal.preState.finalizedCheckpoint.epoch,
            journal.postState.finalizedCheckpoint.epoch,
            journal.preState,
            journal.postState
        );
        _confirmBeaconBlock(finalizedSlot, journal.postState.finalizedCheckpoint.root);
    }

    /**
     * @notice Confirms and stores a beacon block root for a given slot.
     */
    function _confirmBeaconBlock(uint64 slot, bytes32 root) internal {
        if (_allowedBeaconBlocks[slot] == UNDEFINED_ROOT) {
            _allowedBeaconBlocks[slot] = root;
        }
        emit ConfirmedBeaconBlock(slot, root);
    }

    /**
     * @notice Verifies the consensus state transition using the provided ZK proof and journal data.
     * @dev Verifies that the transition from `journal.preState` to `journal.postState` is valid using the cryptographic ZK proof.
     *
     * This function implements two key checks:
     * 1. The consensus pre-state of the journal matches the current trusted state of the contract.
     * 2. The consensus post-state can be transitioned to following Ethereum consensus rules (Casper FFG) starting at the pre-state.
     * This step is verified by the ZK proof. Note: This function solely performs verification. The state update must be handled by the caller.
     */
    function _verify(Journal memory journal, ConsensusData calldata consensus) internal view {
        // Ensure the proof is anchored to the current contract state.
        // The `preState` claimed in the ZK journal must match the `_currentState` actually stored in this contract.
        if (!Consensus.compareState(_currentState, journal.preState)) {
            revert InvalidPreState();
        }
        // Ensure the transition is not stale.
        if (!_permissibleTransition(journal.preState, journal.postState)) {
            revert PermissibleTimespanLapsed();
        }

        // Verify the seal (proof) against the Image ID (circuit verification key) and the Journal (public inputs/outputs).
        // A successful verification confirms that `journal.postState` is the correct result of applying the beacon chain consensus rules to `journal.preState`.
        bytes32 journalHash = sha256(consensus.journalData);
        IRiscZeroVerifier(verifier).verify(consensus.seal, imageID, journalHash);
    }

    function _permissibleTransition(
        Consensus.State memory,
        Consensus.State memory
    ) internal pure returns (bool) {
        // TODO: Add permissible transition check to prevent stale state transitions
        return true;
    }
}
