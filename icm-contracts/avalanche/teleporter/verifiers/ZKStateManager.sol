// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.30;

import { AccessControl } from "@openzeppelin/contracts/access/AccessControl.sol";
import { IRiscZeroVerifier } from "@risc0/contracts/IRiscZeroVerifier.sol";
import { ConsensusState, Checkpoint } from "./tseth.sol";
import {MerkleProof} from "@openzeppelin/contracts//utils/cryptography/MerkleProof.sol";

contract ZKStateManager is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 constant UNDEFINED_ROOT = bytes32(0);

    /// @notice Contains the complete state transition data for verification
    /// @dev Used as journal data in RISC Zero proofs to validate beacon state transitions
    struct Journal {
        /// @dev The consensus state before the transition
        ConsensusState preState;
        /// @dev The consensus state after the transition
        ConsensusState postState;
        /// @dev The beacon chain slot that was finalized in this transition
        uint64 finalizedSlot;
    }

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
    address public immutable VERIFIER;

    /// @notice Maximum allowed time span for state transitions in seconds
    /// @dev Used to prevent acceptance of stale beacon state transitions
    uint24 public permissibleTimespan;

    /// @notice Maps a beacon chain slot to its verified beacon state root 
    /// @dev Used to track the beacon chain slots that are trusted and finalized through verified state transitions
    mapping(uint64 slot => bytes32 beaconRoot) public allowedBeaconStates;

    /// @notice Maps a beacon chain slot to its verified execution state root 
    /// @dev Used to track the execution state slots that are trusted and finalized against verified beacon state roots 
    mapping(uint64 slot => bytes32 execRoot) public allowedExecutionStates;

    /// @notice Maps a beacon chain slot to its verified receipts trie root
    /// @dev Used to track receipt roots that are trusted and finalized against verified execution state roots
    mapping(uint64 slot => bytes32 receiptRoots) public allowedReceiptRoots;

    event Transitioned(
        uint64 indexed preEpoch, uint64 indexed postEpoch, ConsensusState preState, ConsensusState postState
    );
    event Confirmed(uint64 indexed slot, bytes32 indexed root);
    event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID);
    event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan);

    error InvalidArgument();
    error InvalidPreState();
    error PermissibleTimespanLapsed();
    error UnauthorizedEmitterChainId();
    error UnauthorizedEmitterAddress();
    error InvalidBeaconState(); 
    error InvalidProof();

    /// @notice Initializes the ZKStateManager contract with all required parameters
    /// @dev Sets up the initial consensus state, configures verification parameters, and establishes cross-chain
    /// communication
    /// @param startingState The initial consensus state of the beacon chain
    /// @param permissibleTimespan_ Maximum allowed time span for state transitions in seconds
    /// @param verifier Address of the RISC Zero verifier contract for proof validation
    /// @param imageID_ The RISC Zero image ID for the beacon state transition program
    /// @param admin Address to be granted the ADMIN_ROLE
    /// @param superAdmin Address to be granted the DEFAULT_ADMIN_ROLE
    constructor(
        ConsensusState memory startingState,
        uint24 permissibleTimespan_,
        address verifier,
        bytes32 imageID_,
        address admin,
        address superAdmin
    ) {
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(DEFAULT_ADMIN_ROLE, superAdmin);

        currentState = startingState;
        permissibleTimespan = permissibleTimespan_;
        imageID = imageID_;
        VERIFIER = verifier;
    }
     
    /// @notice Validates and applies a beacon state transition using RISC Zero proof
    /// @dev Performs a consensus state transition by verifying that the transition from journal.preState to journal.postState is valid using the ZK proof. 
    /// In more detail, preState = (latestfinalizedCheckpoint, currentJustifiedCheckpoint) and postState = (newFinalizedCheckpoint, newJustifiedCheckpoint). 
    /// Once the ZK proof is verified, currentJustifiedCheckpoint becomes the latest finalized checkpoint, and newFinalizedCheckpoint becomes the current justified checkpoint. 
    /// In other words, the checkpoint that was previously justified is now finalized, and the checkpoint that was previously unconfirmed is now justified.
    /// @param journalData Encoded Journal struct containing pre/post states and finalized slot
    /// @param seal RISC Zero cryptographic proof validating the state transition
    function _verifyBeaconStateTransition(bytes calldata journalData, bytes calldata seal) internal view {
        Journal memory journal = abi.decode(journalData, (Journal));

        // Check if the consensus state transition is allowed; i.e., that the preState matches the current state and that the timespan is permissible.
        if (!_compareConsensusState(currentState, journal.preState)) {
            revert InvalidPreState();
        }
        if (!_permissibleTransition(journal.preState, journal.postState)) {
            revert PermissibleTimespanLapsed();
        }

        // Verify the consensus state transition using the ZK proof. 
        bytes32 journalHash = sha256(journalData);
        IRiscZeroVerifier(VERIFIER).verify(seal, imageID, journalHash);
    }

    /// @notice Validates and applies an execution state transition using Merkle proof authentication
    /// @dev Verify the inclusion of the execution payload root in the beacon state root, and validate 
    /// the chain of intermediate execution state root proofs within that epoch.
    /// @param finalizedBeaconStateRoot The beacon state root of a trusted and finalized consensus checkpoint
    /// @param endExecutionStateRoot The claimed (untrusted) execution state root for the finalized beacon state root `beaconStateRoot` at the end of the epoch boundary
    /// @param endExecutionStateProof Merkle proof authenticating the inclusion of `finalizedExecutionStateRoot` in `beaconStateRoot`
    /// @param executionStateRootsInEpoch Array of intermediate execution state roots within the epoch
    /// @param executionStateProofsInEpoch Array of Merkle proofs for each intermediate execution state roots within the epoch 
    /// @param startSlot The starting slot of the epoch for which the execution state roots are being verified
    function _verifyExecutionStateTransition(
        bytes32 finalizedBeaconStateRoot, 
        bytes32 endExecutionStateRoot, 
        bytes32[] calldata endExecutionStateProof,  
        bytes32[] calldata executionStateRootsInEpoch, 
        bytes32[][] calldata executionStateProofsInEpoch, 
        uint64 startSlot) internal view {
        // Verify the Merkle inclusion of the execution payload root against the beacon state root. 
        bytes32[] memory proof = endExecutionStateProof;
        if (!MerkleProof.verify(proof, finalizedBeaconStateRoot, endExecutionStateRoot)) {
            revert InvalidProof();
        }   

        // Verify all intermediate execution states in the new epoch, starting at the execution state root which corresponds to the prestate finalized checkpoint, and ending at the poststate finalized checkpoint.
        bytes32 startExecutionStateRoot = allowedExecutionStates[startSlot];
        _verifyExecutionStateMerkleChain(startExecutionStateRoot, endExecutionStateRoot, executionStateRootsInEpoch, executionStateProofsInEpoch); 
    }

    /// @notice Internal function to verify a Merkle chain of execution state roots
    /// @dev Verifies a chained sequence of Merkle proofs from a starting execution state root to a final root. 
    /// Each step ensures that leaves[i] is a leaf in the tree with root leaves[i+1], using proofs[i] as the Merkle path.
    /// @param startRoot The initial execution state root at the start of the epoch
    /// @param endRoot The final execution state root at the end of the epoch
    /// @param leaves The ordered list of intermediate execution state roots forming the Merkle chain
    /// @param proofs The corresponding Merkle proofs for each intermediate root in `leaves`
    function _verifyExecutionStateMerkleChain(bytes32 startRoot, bytes32 endRoot, bytes32[] calldata leaves, bytes32[][] calldata proofs) internal pure {
        // Verify intermediate roots in the Merkle chain. 
        bytes32 leaf = startRoot;
        for (uint256 i = 0; i < leaves.length; i++) {
            bytes32 root = leaves[i];
            bytes32[] memory proof = proofs[i];
            if (!MerkleProof.verify(proof, root, leaf)) {
                revert InvalidProof();
            }
            leaf = leaves[i];
        }

        // Check if the final leaf in the chain matches the trusted ending root. 
        require(leaf == endRoot, "Invalid final root in Merkle chain.");
    }

    /// @notice Internal function that verifies the receipts roots against the execution state roots using Merkle proofs.
    /// @dev For each epoch or slot, checks that `receiptRoots[i]` is a leaf in the Merkle tree rooted at 
    /// `executionStateRoots[i]` using the provided `receiptProofs[i]`.
    /// @param executionStateRoots The ordered list of execution state roots, one per slot/epoch
    /// @param receiptProofs The ordered list of Merkle proofs for each receipt root
    /// @param receiptRoots The ordered list of claimed receipt roots to verify against the execution state roots
     function _verifyReceiptRoots(bytes32[] calldata executionStateRoots, bytes32[][] calldata receiptProofs, bytes32[] calldata receiptRoots) internal pure {
        // Verify each receipts root against the corresponding execution state root.
        require(executionStateRoots.length == receiptRoots.length, "Mismatch in execution and receipt roots.");
        for (uint256 i = 0; i < receiptRoots.length; i++) {
            bytes32 root = executionStateRoots[i];
            bytes32 leaf = receiptRoots[i]; 
            bytes32[] memory proof = receiptProofs[i];
            if (!MerkleProof.verify(proof, root, leaf)) {
                revert InvalidProof();
            }
        }
    }

    /// @notice Performs a full state transition from a trusted beacon checkpoint to a new finalized state.
    /// @dev The function performs the following steps:
    /// 1. Verifies the ZK proof of the beacon (consensus) state transition to the new finalized checkpoint.
    /// 2. Validates the Merkle chain of execution state roots from the prior epoch boundary to the new finalized execution state root.
    /// 3. Verifies that all receipt roots are included in their respective execution state roots using Merkle proofs.
    /// 4. Updates the global state to reflect the new trusted consensus, execution, and receipt roots.
    /// @param journalData Encoded Journal struct containing pre- and post-consensus states and finalized slot
    /// @param seal The ZK proof validating the beacon state transition
    /// @param finalizedSlot The beacon chain slot corresponding to the new finalized checkpoint
    /// @param finalizedExecutionStateRoot The claimed execution state root at the epoch boundary to be verified
    /// @param finalizedExecutionStateProof Merkle proof for the finalized execution state root inclusion
    /// @param executionStateRootsInEpoch The ordered list of intermediate execution state roots within the epoch
    /// @param executionStateProofsInEpoch Merkle proofs for each intermediate execution state root
    /// @param receiptRootsInEpoch The ordered list of claimed receipt roots for the epoch
    /// @param receiptRootProofsInEpoch Merkle proofs for each receipt root
    function transition(
        bytes calldata journalData,
        bytes calldata seal,
        uint64 finalizedSlot,
        bytes32 finalizedExecutionStateRoot,
        bytes32[] calldata finalizedExecutionStateProof,
        bytes32[] calldata executionStateRootsInEpoch,
        bytes32[][] calldata executionStateProofsInEpoch,
        bytes32[] calldata receiptRootsInEpoch, 
        bytes32[][] calldata receiptRootProofsInEpoch
    ) external {
        // Decode the journal data. 
        Journal memory journal = abi.decode(journalData, (Journal));
        
        // Transition the consensus state to the new finalized checkpoint after verifying the ZK proof.
        _verifyBeaconStateTransition(journalData, seal);
        
        // Verify the execution state of the new finalized checkpoint, along with the intermediate execution states in the latest epoch.
        _verifyExecutionStateTransition({
            finalizedBeaconStateRoot: journal.postState.finalizedCheckpoint.root,
            endExecutionStateRoot: finalizedExecutionStateRoot,
            endExecutionStateProof: finalizedExecutionStateProof,
            executionStateRootsInEpoch: executionStateRootsInEpoch,
            executionStateProofsInEpoch: executionStateProofsInEpoch,
            startSlot: journal.preState.finalizedCheckpoint.epoch * SLOT_PER_EPOCH
        });

        // Verify the receipts roots against the execution state roots.
        _verifyReceiptRoots(executionStateRootsInEpoch, receiptRootProofsInEpoch, receiptRootsInEpoch);
        
        // Finally, update the global state to reflect the new trusted consensus, execution, and receipts states. 
        _transitionBeaconState(journal, finalizedSlot); 
        _transitionExecutionState(journal.preState.finalizedCheckpoint.epoch * SLOT_PER_EPOCH + 1, executionStateRootsInEpoch);
        _transitionReceiptRootsState(journal.preState.finalizedCheckpoint.epoch * SLOT_PER_EPOCH + 1, receiptRootsInEpoch);
    }

    
    /// @notice Checks if a receipt root is trusted.
    /// @dev The caller must provide the slot hint to avoid expensive iteration.
    /// @param slot The beacon chain slot where this root was finalized.
    /// @param root The receipt root to validate.
    function isValidReceiptsRoot(uint64 slot, bytes32 root) external view returns (bool) {
        return (root != bytes32(0) && allowedReceiptRoots[slot] == root);
    }

    // Perform a manual transition of the consensus state. This function is restricted to ADMIN_ROLE.
    function manualTransition(bytes calldata journalData, uint64 finalizedSlot) external onlyRole(ADMIN_ROLE) {
        Journal memory journal = abi.decode(journalData, (Journal));
        _transitionBeaconState(journal, finalizedSlot);
    }

    /**
     * @notice the root associated with the provided `slot`. If the confirmation level isn't met or the root is not
     * set, `valid` will be false
     *
     * TODO: Add in link ref to confirmation levels
     *
     * @param slot the beacon chain slot to look up
     */
    function blockRoot(uint64 slot) external view returns (bytes32 root, bool valid) {
        root = allowedBeaconStates[slot];
        if (root == UNDEFINED_ROOT) {
            valid = false;
        }
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

    function _transitionBeaconState(Journal memory journal, uint64 finalizedSlot) internal {
        currentState = journal.postState;
        emit Transitioned(
            journal.preState.finalizedCheckpoint.epoch,
            journal.postState.finalizedCheckpoint.epoch,
            journal.preState,
            journal.postState
        );

        Checkpoint memory finalizedCheckpoint = journal.postState.finalizedCheckpoint;
        _confirmBeaconState(finalizedSlot, finalizedCheckpoint.root);
    }

    function _transitionExecutionState(uint64 startSlot, bytes32[] memory execRoots) internal {
        for (uint64 i = 0; i < execRoots.length; i++) {
            _confirmExecutionState(startSlot + uint64(i), execRoots[i]);
        }
    }

    function _transitionReceiptRootsState(uint64 startSlot, bytes32[] memory receiptRoots) internal {
        for (uint64 i = 0; i < receiptRoots.length; i++) {
            _confirmReceiptsRoot(startSlot + uint64(i), receiptRoots[i]);
        }
    }

    function _confirmBeaconState(uint64 slot, bytes32 root) internal {
        if (allowedBeaconStates[slot] == UNDEFINED_ROOT) {
            allowedBeaconStates[slot] = root;
        }
        emit Confirmed(slot, root);
    }

    function _confirmExecutionState(uint64 slot, bytes32 root) internal {
        if (allowedExecutionStates[slot] == UNDEFINED_ROOT) {
            allowedExecutionStates[slot] = root;
        }
        /// TODO: Add confirmed event for execution root
    }

    function _confirmReceiptsRoot(uint64 slot, bytes32 root) internal {
        if (allowedReceiptRoots[slot] == UNDEFINED_ROOT) {
            allowedReceiptRoots[slot] = root;
        }
        /// TODO: Add confirmed event for receipts root
    }

    function _compareConsensusState(ConsensusState memory a, ConsensusState memory b) internal pure returns (bool) {
        return _compareCheckpoint(a.currentJustifiedCheckpoint, b.currentJustifiedCheckpoint)
            && _compareCheckpoint(a.finalizedCheckpoint, b.finalizedCheckpoint);
    }

    function _compareCheckpoint(Checkpoint memory a, Checkpoint memory b) internal pure returns (bool) {
        return a.epoch == b.epoch && a.root == b.root;
    }

    /// TODO: Fix this 
    function _permissibleTransition(
        ConsensusState memory pre,
        ConsensusState memory post
    )
        internal
        view
        returns (bool)
    {
        // uint256 transitionTimespan = block.timestamp
        //     - Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, post.finalizedCheckpoint.epoch);
        // TODO: Come back to this logic.
        // uint256 transitionTimespan = Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, post.finalizedCheckpoint.epoch)
        // - Beacon.epochTimestamp(Beacon.ETHEREUM_GENESIS_BEACON_BLOCK_TIMESTAMP, pre.finalizedCheckpoint.epoch);
        // return transitionTimespan <= uint256(permissibleTimespan);
        return true;
    }

    /// @notice Generates a unique hash for block that was included in the chain at the given slot
    function _checkpointHash(uint64 slot, bytes32 root) internal pure returns (bytes32 hash) {
        hash = keccak256(abi.encodePacked(slot, root));
    }
}
