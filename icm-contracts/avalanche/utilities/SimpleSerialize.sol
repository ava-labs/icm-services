// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Reference: This is core logic from the Succinct Telepathy Library, which can be found at https://github.com/succinctlabs/telepathy-contracts/blob/main/src/libraries/SimpleSerialize.sol
pragma solidity ^0.8.30;

library SSZ {
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
