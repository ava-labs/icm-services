// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Based on code from: https://github.com/boundless-xyz/boundless-transceiver
// Modifications: Builds a State Manager for the Ethereum beacon chain, including storing beacon state roots, 
// execution state roots, and receipt roots, on top of the infrastructure provided by the boundless-transceiver repository.

pragma solidity ^0.8.30;

import { AccessControl } from "@openzeppelin/contracts/access/AccessControl.sol";
import { IRiscZeroVerifier } from "@risc0/contracts/IRiscZeroVerifier.sol";
import { ConsensusState, Checkpoint } from "./tseth.sol";
import { MerkleProof } from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import { SSZ } from "../../utilities/SimpleSerialize.sol";
import { MerklePatricia, StorageValue } from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";
import { RLPUtils } from "./RLPUtils.sol";

/// @notice Information about an imported event from the beacon chain.
struct ZKEventInfo {
    uint256 sourceChainId;     
    uint256 beaconSlot;        
    bytes32 executionRoot;     
    uint256 logIndex;    
    bytes logData;            
}

/// @notice Contains all the data required to prove a specific log was emitted.
struct LogProof {
    bytes[] proof;           // The Merkle Patricia Trie inclusion proof (array of rlp-encoded nodes).
    bytes key;               // The RLP-encoded transaction index
    bytes value;             // The RLP-encoded receipt data
    uint256 logIndex;        // The specific index of the log in the receipt
    address expectedEmitter; // Contract address that should have emitted the log
    bytes32 expectedTopic0;  // The event signature (keccak256("Event(args)"))
}

/// @notice Contains the public state transition data for the ZK proof
/// @dev The ZK proof attests that `postState` is the valid successor to `preState` according to consensus rules.
struct Journal {
    /// @dev The starting trust anchor (must match the contract's stored `currentState`).
    ConsensusState preState;
    /// @dev The new consensus state proven to be valid by the ZK circuit.
    ConsensusState postState;
    /// @dev The finalized slot number extracted from `postState` (used for indexing storage).
    uint64 finalizedSlot;
}

/// @notice Contains the data required to verify and perform a beacon chain state transition.
/// @dev This struct bundles the ZK proof ("seal") with the public inputs ("journal") required to 
/// verify that the Beacon Chain has successfully advanced to a new finalized state.
struct ConsensusData{
    /// @dev Encoded Journal struct containing pre- and post-consensus states and finalized slot
    bytes journalData;
    /// @dev The RISC Zero ZK proof validating the beacon state transition
    bytes seal; 
    /// @dev The beacon chain slot corresponding to the new finalized checkpoint 
    uint64 finalizedSlot; 
}

/// @notice A cryptographic proof bundle establishing that an execution layer receipt root is valid for a specified beacon chain slot. 
/// @dev This struct containing inclusion proofs required to verify an execution layer event (receipt) statelessly.
/// 1. Anchor Check: Verifies the `anchorBeaconState` is valid against a trusted beacon block root stored in this contract (specified by `anchorSlot`).
/// 2. History Check: Verifies the `targetBeaconState` exists within the `anchorBeaconState`'s historical state roots vector.
/// 3. Execution Check: Verifies the `targetExecutionHeader` root is included in the `targetBeaconState`.
/// 4. Receipts Check: Verifies the `targetReceiptsRoot` is included in the `targetExecutionHeader`.
struct ExecutionProof {
    // The specific slot where the transaction happened
    uint64 targetSlot;
    // The specific slot for the beacon block root we are using as the anchor
    uint64 anchorSlot;
    // The Anchor State Proof (Block -> State)
    bytes32 anchorBeaconStateRoot;
    bytes32[] anchorBeaconStateProof;
    // The History Proof (Anchor Beacon State -> Target Beacon State)
    bytes32 targetBeaconStateRoot;
    bytes32[] targetBeaconStateProof;
    // The Execution Proof (Target Beacon State -> Execution Header Root)
    bytes32 targetExecutionHeaderRoot;
    bytes32[] targetExecutionHeaderProof;
    // The Receipts Proof (Execution Header Root -> Receipts Root)
    bytes32 targetReceiptsRoot; 
    bytes32[] targetReceiptsProof;
}

contract ZKStateManager is AccessControl {

    using RLPReader for bytes;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 constant UNDEFINED_ROOT = bytes32(0);

    // Generalized Indices (Constants for Deneb / Electra)
    uint256 constant G_INDEX_BLOCK_STATE_ROOT = 11; // BlockHeader -> StateRoot
    uint256 constant G_INDEX_EXEC_ROOT = 1794;      // BeaconState -> ExecutionPayload -> StateRoot
    uint256 constant G_INDEX_BASE_STATE_ROOTS = 38; // BeaconState -> StateRoots Vector
    uint256 constant STATE_ROOTS_DEPTH = 13;                  // Depth of StateRoots Vector (8192 entries)
    uint256 constant G_INDEX_RECEIPTS_ROOT = 1795; // ExecutionPayload -> ReceiptsRoot

    uint256 public immutable SOURCE_CHAIN_ID;
    
    /// @notice Number of slots per epoch in the beacon chain
    uint64 SLOT_PER_EPOCH = 32;

    /// @notice The current consensus state of the beacon chain
    /// @dev Updated atomically through state transitions to ensure consistency
    ConsensusState private currentState;

    /// @notice The RISC Zero program ID used to verify Ethereum consensus transitions
    /// @dev This image ID corresponds to the ZK program (Signal Ethereum) that validates state transitions. 
    /// Normally it stays constant, but it can be updated if the consensus program is upgraded, for example. 
    /// All proofs are verified against this program.
    bytes32 public imageID;

    /// @notice The address of the RISC Zero verifier contract
    /// @dev Used to validate zero-knowledge proofs of beacon state transitions
    address public VERIFIER;

    /// @notice Maximum allowed time span for state transitions in seconds
    /// @dev Used to prevent acceptance of stale beacon state transitions
    uint24 public permissibleTimespan;

    /// @notice Maps a beacon chain slot to its verified beacon block root 
    /// @dev Used to track the beacon chain slots that are trusted and finalized through verified state transitions
    mapping(uint64 slot => bytes32 beaconBlockRoot) public allowedBeaconBlocks;

    /// @notice Maps a beacon chain slot to its verified receipts trie root
    /// @dev Used to track receipt roots that are trusted and finalized against verified execution state roots
    mapping(uint64 slot => bytes32 receiptRoots) public allowedReceiptRoots;

    event Transitioned(
        uint64 indexed preEpoch, uint64 indexed postEpoch, ConsensusState preState, ConsensusState postState
    );
    event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root);
    event ConfirmedBeaconState(uint64 indexed slot, bytes32 indexed root);
    event ConfirmedExecutionState(uint64 indexed slot, bytes32 indexed root);
    event ConfirmedReceiptsState(uint64 indexed slot, bytes32 indexed root);
    event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID);
    event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan);
    event ZKEventImported(
        uint256 indexed sourceChainId,
        uint256 indexed beaconSlot,   
        address indexed emitter,      
        bytes32 executionRoot,        
        uint256 logIndex
    );

    error InvalidArgument();
    error InvalidPreState();
    error PermissibleTimespanLapsed();

    /// @notice Initializes the ZKStateManager contract with all required parameters
    /// @dev Sets up the initial consensus state, configures verification parameters, and establishes cross-chain
    /// communication
    /// @param _sourceChainId The ID of the chain this contract will track
    /// @param startingState The initial consensus state of the beacon chain
    /// @param permissibleTimespan_ Maximum allowed time span for state transitions in seconds
    /// @param verifier Address of the RISC Zero verifier contract for proof validation
    /// @param imageID_ The RISC Zero image ID for the beacon state transition program
    /// @param admin Address to be granted the ADMIN_ROLE
    /// @param superAdmin Address to be granted the DEFAULT_ADMIN_ROLE
    constructor(
        uint256 _sourceChainId,
        ConsensusState memory startingState,
        uint24 permissibleTimespan_,
        address verifier,
        bytes32 imageID_,
        address admin,
        address superAdmin
    ) {
        require(_sourceChainId != 0, "Invalid Chain ID");
        SOURCE_CHAIN_ID = _sourceChainId;

        _grantRole(ADMIN_ROLE, admin);
        _grantRole(DEFAULT_ADMIN_ROLE, superAdmin);

        currentState = startingState;
        permissibleTimespan = permissibleTimespan_;
        imageID = imageID_;
        VERIFIER = verifier;
    }

    /// @notice Performs a full state transition from a trusted beacon checkpoint to a new finalized state.
    /// @dev The function performs the following steps:
    /// 1. Verifies the ZK proof of the beacon (consensus) state transition to the new finalized checkpoint.
    /// 2. Validates the Merkle chain of execution state roots from the prior epoch boundary to the new finalized execution state root.
    /// 3. Verifies that all receipt roots are included in their respective execution state roots using Merkle proofs.
    /// 4. Updates the global state to reflect the new trusted consensus, execution, and receipt roots.
    function transition(
        ConsensusData calldata consensus
    ) external {
        // Decode the journal data. 
        Journal memory journal = abi.decode(consensus.journalData, (Journal));
        // Verify the consensus and execution state transitions. 
        _verifyConsensusTransition(journal, consensus);
        // Finally, update the global state to reflect the new trusted consensus, execution, and receipts states. 
        _transitionConsensusState(journal, consensus.finalizedSlot);
    }

    /// @notice Proves that a specific log was emitted on the beacon chain and executes application logic.
    /// @dev This is the main entry point for bridging events. It performs two key verifications:
    /// 1. Execution Verification: Validates that the provided `targetReceiptsRoot` is part of the canonical beacon chain history using the `execProof`.
    /// 2. Log Verification: Uses the validated `targetReceiptsRoot` to verify the inclusion  of a specific Receipt (and Log) via the `logProof`.
    /// Once verified, it emits a `ZKEventImported` event and passes the verified data to the internal `_onEventImport` handler for application-specific processing.
    /// @param execProof The execution proof linking a receipt root to a trusted beacon block root.
    /// @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
    function proveLogAndExecute(
        ExecutionProof calldata execProof, 
        LogProof calldata logProof
    ) external {
        // Verify the execution state proof, proving that the contained receipts root is legitimate against a stored beacon block root. 
        _verifyExecutionState(execProof);

        // Extract the trusted receipts root to be used. 
        bytes32 trustedReceiptsRoot = execProof.targetReceiptsRoot;

        // Verify the receipt and log, and extract the log data.
        bytes memory logData = _verifyReceiptAndLog(trustedReceiptsRoot, logProof);

        // Emit the event. 
        emit ZKEventImported(
            SOURCE_CHAIN_ID,
            execProof.targetSlot,
            logProof.expectedEmitter,
            execProof.targetExecutionHeaderRoot,
            logProof.logIndex
        );

        // Construct the event info struct. 
        ZKEventInfo memory info = ZKEventInfo({
            sourceChainId: SOURCE_CHAIN_ID,
            beaconSlot: execProof.targetSlot,
            executionRoot: execProof.targetExecutionHeaderRoot,
            logIndex: logProof.logIndex,
            logData: logData
        });

        // Hand off the event to application logic.
        _onEventImport(info);
    }

    function updateImageID(bytes32 newImageID) external onlyRole(ADMIN_ROLE) {
        if (newImageID == imageID) revert InvalidArgument();
        imageID = newImageID;
    }

    function updatePermissibleTimespan(uint24 newPermissibleTimespan) external onlyRole(ADMIN_ROLE) {
        if (newPermissibleTimespan == permissibleTimespan) {
            revert InvalidArgument();
        }
        permissibleTimespan = newPermissibleTimespan;
    }

    /// @notice Checks if a receipt root is trusted.
    /// @param slot The beacon chain slot where this root was finalized.
    /// @param root The receipt root to validate.
    function isValidReceiptsRoot(uint64 slot, bytes32 root) external view returns (bool) {
        return (root != bytes32(0) && allowedReceiptRoots[slot] == root);
    }

    // Perform a manual transition of the consensus state. This function is restricted to ADMIN_ROLE.
    // TODO: Come back to this
    // function manualTransition(bytes calldata journalData, uint64 finalizedSlot) external onlyRole(ADMIN_ROLE) {
    //     Journal memory journal = abi.decode(journalData, (Journal));
    //     _transitionBeaconState(journal, finalizedSlot);
    // }

    /// @notice Outputs the beacon block root associated with the provided `slot`.
    /// @param slot The beacon chain slot to look up
    function getBeaconBlockRoot(uint64 slot) external view returns (bytes32 root, bool valid) {
        root = allowedBeaconBlocks[slot];
        if (root == UNDEFINED_ROOT) {
            valid = false;
        }
    }

     /// @notice Application-specific logic to handle the imported event.
    function _onEventImport(ZKEventInfo memory eventInfo) internal virtual {
        // Placeholder for application logic to handle the imported event.
    }

    /// @notice Transitions and updates the consensus state of the contract to the new post-state. 
    function _transitionConsensusState(Journal memory journal, uint64 finalizedSlot) internal {
        currentState = journal.postState;
        emit Transitioned(
            journal.preState.finalizedCheckpoint.epoch,
            journal.postState.finalizedCheckpoint.epoch,
            journal.preState,
            journal.postState
        );
        _confirmBeaconBlock(finalizedSlot, journal.postState.finalizedCheckpoint.root);
    }

    /// @notice Confirms and stores a beacon block root for a given slot.
    function _confirmBeaconBlock(uint64 slot, bytes32 root) internal {
        if (allowedBeaconBlocks[slot] == UNDEFINED_ROOT) {
            allowedBeaconBlocks[slot] = root;
        }
        emit ConfirmedBeaconBlock(slot, root);
    }

    /// @notice Verifies the consensus state transition using the provided ZK proof and journal data.
    /// @dev Verifies that the transition from `journal.preState` to `journal.postState` is valid using the cryptographic ZK proof. 
    /// This function implements two key checks:
    /// 1. The consensus pre-state of the journal matches the current trusted state of the contract. 
    /// 2. The consensus post-state can be transitioned to following Ethereum consensus rules (Casper FFG) starting at the pre-state. This step is verified by the ZK proof.
    /// Note: This function solely performs verification. The state update must be handled by the caller. 
    function _verifyConsensusTransition(
        Journal memory journal,
        ConsensusData calldata consensus
    ) internal view {
        // Ensure the proof is anchored to the current contract state.
        // The `preState` claimed in the ZK journal must match the `currentState` actually stored in this contract.
        if (!_compareConsensusState(currentState, journal.preState)) {
            revert InvalidPreState();
        }
        // if (!_permissibleTransition(journal.preState, journal.postState)) {
        //     revert PermissibleTimespanLapsed();
        // }

        // Verify the seal (proof) against the Image ID (circuit verification key) and the Journal (public inputs/outputs).
        // A successful verification confirms that `journal.postState` is the correct result of applying the beacon chain consensus rules to `journal.preState`.
        bytes32 journalHash = sha256(consensus.journalData);
        IRiscZeroVerifier(VERIFIER).verify(consensus.seal, imageID, journalHash);
    }

     /**
     * @notice Verifies an execution-layer receipts root by tracing it back through a chain of trust to a beacon block root.
     * @dev Performs a stateless verification using a 4-step chain of trust:
     * 1. Anchor Check: Validates an anchor beacon state against the trusted block root stored in `allowedBeaconBlocks`.
     * 2. History Check: Validates the target beacon state against the anchor beacon state's `state_roots` history vector. 
     * 3. Execution Check: Validates the execution payload header against the target beacon state.
     * 4. Receipts Check: Validates the receipts root against the execution payload header.
     * Trusted Block -> Anchor State -> Target State -> Execution Header -> Receipts Root
     * @param proof The `ExecutionProof` struct containing the target/anchor slots, relevant roots, and all Merkle proofs required for the verification path.
     */
    function _verifyExecutionState(ExecutionProof calldata proof) internal view {
        // Anchor Check: Verify the anchor beacon state root is valid agains the trusted beacon block root stored in the contract state
        bytes32 anchorBeaconBlockRoot = allowedBeaconBlocks[proof.anchorSlot];
        bool validAnchor = SSZ.isValidMerkleProof(proof.anchorBeaconStateRoot, G_INDEX_BLOCK_STATE_ROOT, proof.anchorBeaconStateProof, anchorBeaconBlockRoot);
        require(validAnchor, "Invalid anchor state root");

        // Verify the target beacon state root is in the anchor's history. This is possible since beacon states contain a vector of historical state roots 'state_roots' (referencing the last 8192 slots). 
        // Safety Check. Can only prove history within the vector limit (8192 slots).
        require(proof.targetSlot < proof.anchorSlot, "Target must be in the past");
        require(proof.anchorSlot - proof.targetSlot <= 8192, "Target too old");

        // Calculate the specific G-Index for 'state_roots[targetSlot]' within the beacon state SSZ structure.
        uint256 vectorIndex = proof.targetSlot % 8192;
        uint256 targetGIndex = (G_INDEX_BASE_STATE_ROOTS << STATE_ROOTS_DEPTH) + vectorIndex;

        // History Check: Prove that the target beacon state root is in the anchor beacon state's history using the G-index and the SSZ Merkle proof. 
        bool validBeacon = SSZ.isValidMerkleProof(
            proof.targetBeaconStateRoot,
            targetGIndex,
            proof.targetBeaconStateProof,
            proof.anchorBeaconStateRoot
        );
        require(validBeacon, "Invalid target beacon state root");

        // Execution Check: Verify the exection header root is in the target beacon state. 
        bool validExecutionHeader = SSZ.isValidMerkleProof(
            proof.targetExecutionHeaderRoot,
            G_INDEX_EXEC_ROOT,
            proof.targetExecutionHeaderProof,
            proof.targetBeaconStateRoot
        );
        require(validExecutionHeader, "Invalid execution root proof");

        // Receipts Check: Prove the that the target receipts root is in the execution header. 
        bool validReceiptsRoot = SSZ.isValidMerkleProof(
            proof.targetReceiptsRoot,       
            G_INDEX_RECEIPTS_ROOT,                            
            proof.targetReceiptsProof,            
            proof.targetExecutionHeaderRoot  
        );
        require(validReceiptsRoot, "Invalid receipts root");
    }

    /**
     * @notice Verifies that a specific receipt exists in the receipts Merkle trie and that it contains a specific log. 
     * @param trustedReceiptsRoot A trusted receipts root to verify against. 
     * @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
     * @return logData The non-indexed data from the log (for the caller to decode).
     */
    function _verifyReceiptAndLog(bytes32 trustedReceiptsRoot, LogProof calldata logProof) internal pure returns (bytes memory) {
        
        // Construct the key of the trie receipt proof.
        bytes[] memory keys = new bytes[](1);
        keys[0] = logProof.key; // The RLP encoded index passed by the caller

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(
            trustedReceiptsRoot, 
            logProof.proof, 
            keys
        );
        require(results.length == 1, "Invalid number of results in receipt proof");
        require(results[0].value.length > 0, "Invalid receipt proof");
        require(keccak256(results[0].value) == keccak256(logProof.value), "Proven value does not match expected receipt");

        // Decode the receipt. 
        RLPUtils.EVMReceipt memory receipt = RLPUtils.decodeReceipt(results[0].value.toRlpItem());
        require(logProof.logIndex < receipt.logs.length, "Invalid log index");

        // Verify the log within the receipt.
        RLPUtils.EVMLog memory log = receipt.logs[logProof.logIndex];
        require(log.loggerAddress == logProof.expectedEmitter, "Invalid log emitter");
        require(log.topics.length > 0, "Log has no topics");
        require(log.topics[0] == logProof.expectedTopic0, "Invalid event signature");

        return log.data;
    }

    /// @notice Compares two consensus states structs for equality. 
    function _compareConsensusState(ConsensusState memory a, ConsensusState memory b) internal pure returns (bool) {
        return _compareCheckpoint(a.currentJustifiedCheckpoint, b.currentJustifiedCheckpoint)
            && _compareCheckpoint(a.finalizedCheckpoint, b.finalizedCheckpoint);
    }

    /// @notice Compares two consensus checkpoints by checking if they have the same epoch number and root. 
    function _compareCheckpoint(Checkpoint memory a, Checkpoint memory b) internal pure returns (bool) {
        return a.epoch == b.epoch && a.root == b.root;
    }

    /// @notice Generates a unique hash for block that was included in the chain at the given slot
    function _checkpointHash(uint64 slot, bytes32 root) internal pure returns (bytes32 hash) {
        hash = keccak256(abi.encodePacked(slot, root));
    }

    // /// TODO: Update this 
    // function _permissibleTransition(
    //     ConsensusState memory pre,
    //     ConsensusState memory post
    // )
    //     internal
    //     view
    //     returns (bool)
    // {
    //     // uint256 transitionTimespan = block.timestamp
    //     //     - Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, post.finalizedCheckpoint.epoch);
    //     // TODO: Come back to this
    //     // uint256 transitionTimespan = Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, post.finalizedCheckpoint.epoch)
    //     // - Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, pre.finalizedCheckpoint.epoch);
    //     // return transitionTimespan <= uint256(permissibleTimespan);
    //     return true;
    // }
}

