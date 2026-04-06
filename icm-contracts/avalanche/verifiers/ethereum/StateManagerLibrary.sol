// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {SSZ} from "../../utilities/SimpleSerialize.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees-local/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees-local/trie/ethereum/RLPReader.sol";
import {RLPUtils} from "../../utilities/RLPUtils.sol";

/**
 * THIS IS LIBRARY IS UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
library Consensus {
    struct Checkpoint {
        uint64 epoch;
        bytes32 root; // beacon block root for epoch boundary block
    }

    struct State {
        Checkpoint currentJustifiedCheckpoint;
        Checkpoint finalizedCheckpoint;
    }

    /**
     * @notice Compares two consensus states structs for equality.
     */
    function compareState(State memory a, State memory b) internal pure returns (bool) {
        return compareCheckpoint(a.currentJustifiedCheckpoint, b.currentJustifiedCheckpoint)
            && compareCheckpoint(a.finalizedCheckpoint, b.finalizedCheckpoint);
    }

    /**
     * @notice Compares two consensus checkpoints by checking if they have the same epoch number and root.
     */
    function compareCheckpoint(
        Checkpoint memory a,
        Checkpoint memory b
    ) internal pure returns (bool) {
        return a.epoch == b.epoch && a.root == b.root;
    }

    /**
     * @notice Generates a unique hash for block that was included in the chain at the given slot
     */
    function checkpointHash(uint64 slot, bytes32 root) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(slot, root));
    }
}

library Execution {
    struct BeaconConfig {
        uint256 gIndexBlockStateRoot;
        uint256 gIndexExecRoot;
        uint256 gIndexBaseStateRoots;
        uint256 stateRootsDepth;
        uint256 gIndexReceiptsRoot;
        uint256 stateRootsVectorSize;
    }

    /**
     * @notice A cryptographic proof bundle establishing that an execution layer receipt root is valid for a specified beacon chain slot.
     * @dev This struct contains inclusion proofs required to verify an execution layer event, i.e., a transaction receipt.
     * 1. Anchor Check: Verifies the `anchorBeaconState` is valid against a trusted beacon block root.
     * 2. History Check: Verifies the `targetBeaconState` exists within the `anchorBeaconState` historical state roots vector.
     * 3. Execution Check: Verifies the `targetExecutionHeader` root is included in the `targetBeaconState`.
     * 4. Receipts Check: Verifies the `targetReceiptsRoot` is included in the `targetExecutionHeader`.
     */
    struct Proof {
        // The specific slot for the beacon block root we are using as the anchor
        uint64 anchorSlot;
        // The specific slot where the transaction happened
        uint64 targetSlot;
        // The Anchor State Proof (Trusted Beacon Block -> Anchor Beacon State)
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

    /**
     * @notice Verifies an execution-layer receipts root by tracing it back through a chain of trust to a beacon block root.
     * @dev Chain of trust: Trusted Beacon Block -> Anchor Beacon State -> Target Inter-Epoch Beacon State -> Execution Header -> Receipts Root
     * 1. Anchor check: Validates an anchor beacon state against a trusted block root.
     * 2. History check: Validates the target beacon state against the anchor beacon state's `state_roots` history vector.
     * 3. Execution check: Validates the execution payload header against the target beacon state.
     * 4. Receipts check: Validates the receipts root against the execution payload header.
     * @param proof The `ExecutionProof` struct containing the target/anchor slots, relevant roots, and all Merkle proofs required for the verification path.
     */
    function verify(
        bytes32 trustedBeaconBlockRoot,
        Proof calldata proof,
        BeaconConfig storage config
    ) internal view {
        // Safety Check. Can only prove history within the state roots vector size.
        require(proof.targetSlot < proof.anchorSlot, "Target slot must be less than anchor slot");
        require(
            proof.anchorSlot - proof.targetSlot <= config.stateRootsVectorSize,
            "Target slot is too old"
        );

        // Anchor check: Verify the anchor beacon state root is valid against a trusted beacon block root
        bool validAnchor = SSZ.isValidMerkleProof(
            proof.anchorBeaconStateRoot,
            config.gIndexBlockStateRoot,
            proof.anchorBeaconStateProof,
            trustedBeaconBlockRoot
        );
        require(validAnchor, "Invalid anchor state root");

        // Calculate the specific G-Index for 'state_roots[targetSlot]' within the beacon state SSZ structure.
        uint256 vectorIndex = proof.targetSlot % config.stateRootsVectorSize;
        uint256 targetGIndex = (config.gIndexBaseStateRoots << config.stateRootsDepth) + vectorIndex;

        // History check: Prove that the target beacon state root is in the anchor beacon state's history using the G-index and the SSZ Merkle proof.
        // This is possible since beacon states contain a vector of historical state roots `state_roots` referencing the last 8192 slots.
        bool validBeacon = SSZ.isValidMerkleProof(
            proof.targetBeaconStateRoot,
            targetGIndex,
            proof.targetBeaconStateProof,
            proof.anchorBeaconStateRoot
        );
        require(validBeacon, "Invalid target beacon state root");

        // Execution check: Verify the execution header root is in the target beacon state.
        bool validExecutionHeader = SSZ.isValidMerkleProof(
            proof.targetExecutionHeaderRoot,
            config.gIndexExecRoot,
            proof.targetExecutionHeaderProof,
            proof.targetBeaconStateRoot
        );
        require(validExecutionHeader, "Invalid execution root proof");

        // Receipts check: Prove the that the target receipts root is in the execution header.
        bool validReceiptsRoot = SSZ.isValidMerkleProof(
            proof.targetReceiptsRoot,
            config.gIndexReceiptsRoot,
            proof.targetReceiptsProof,
            proof.targetExecutionHeaderRoot
        );
        require(validReceiptsRoot, "Invalid receipts root");
    }
}

library Receipt {
    using RLPReader for bytes;

    /**
     * @notice Contains all the data required to prove a specific log/event was emitted.
     */
    struct Proof {
        // The Merkle Patricia Trie inclusion proof (array of rlp-encoded nodes).
        bytes[] proof;
        // The RLP-encoded transaction index
        bytes key;
        // The RLP-encoded receipt data
        bytes value;
        // The specific index of the log in the receipt
        uint256 logIndex;
        // Contract address that should have emitted the log
        address expectedEmitter;
        // The event signature
        bytes32 expectedTopic0;
    }

    /**
     * @notice Verifies that a specific receipt exists in the receipts Merkle trie and that it contains a specific log.
     * @param trustedReceiptsRoot A trusted receipts root to verify against.
     * @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
     * @return logData The non-indexed data from the log. The caller must decode this.
     */
    function verifyAndExtractLog(
        bytes32 trustedReceiptsRoot,
        Proof calldata logProof
    ) internal pure returns (bytes memory) {
        // Construct the key of the trie receipt proof.
        bytes[] memory keys = new bytes[](1);
        keys[0] = logProof.key;

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results =
            MerklePatricia.VerifyEthereumProof(trustedReceiptsRoot, logProof.proof, keys);
        require(results.length == 1, "Invalid number of results in receipt proof");
        require(results[0].value.length > 0, "Invalid receipt proof");

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
}
