// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Reference: This is core logic from the Succinct Telepathy Library, which can be found at https://github.com/succinctlabs/telepathy-contracts/blob/main/src/libraries/SimpleSerialize.sol
pragma solidity ^0.8.30;

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

library SSZ {
    /// @notice Merkleizes 32 leaves from calldata using an iterative stack
    /// @param leaves The leaves data to Merkleize
    /// @return bytes32 The computed Merkle Root
    // Requires that the tree is balanced, i.e., the number of leaves is a power of two
    function merkleize(
        bytes32[] memory leaves
    ) internal pure returns (bytes32) {
        // Safety checks
        uint256 count = leaves.length;
        require(count > 0, "Leaves must be non-empty");
        require((count & (count - 1)) == 0, "Leaves must be power of 2");
        // Allocate the stack to store intermediate nodes
        uint256 depth = Math.log2(count, Math.Rounding.Floor);
        bytes32[] memory stack = new bytes32[](depth + 1);
        // Loop through every leaf
        for (uint256 i = 0; i < count; i++) {
            bytes32 node = leaves[i];
            // Merge up logic.
            // If the current index is odd, it means we just finished a right child
            // We must hash it with the left child waiting in the stack
            uint256 size = i + 1;
            uint256 level = 0;
            // While size is even, we are a right sibling, then merge up
            while ((size % 2) == 0) {
                node = sha256(bytes.concat(stack[level], node));
                level++;
                size /= 2;
            }
            // Store the pending node, waiting for its right sibling
            stack[level] = node;
        }
        // Return the root
        return stack[depth];
    }

    /// @notice Computes the Merkle Root using a leaf, its index, and the proof
    /// @param leaf The leaf data to verify
    /// @param index The Generalized Index (gIndex) of the leaf
    /// @param proof The Merkle proof (array of sibling hashes)
    /// @return bytes32 The reconstructed Merkle Root.
    function restoreMerkleRoot(
        bytes32 leaf,
        uint256 index,
        bytes32[] memory proof
    ) internal pure returns (bytes32) {
        bytes32 value = leaf;
        uint256 i = 0;
        while (index != 1) {
            if (index % 2 == 1) {
                // If index is odd, we are a Right child. Sibling is on the Left.
                value = sha256(bytes.concat(proof[i], value));
            } else {
                // If index is even, we are a Left child. Sibling is on the Right.
                value = sha256(bytes.concat(value, proof[i]));
            }
            index /= 2;
            i++;
        }
        return value;
    }

    /// @notice Verifies that a Merkle proof correctly connects a leaf to a root
    function isValidMerkleProof(
        bytes32 leaf,
        uint256 index,
        bytes32[] memory proof,
        bytes32 root
    ) internal pure returns (bool) {
        return restoreMerkleRoot(leaf, index, proof) == root;
    }
}
