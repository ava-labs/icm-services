// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Reference: This is core logic from the Succinct Telepathy Library, which can be found at https://github.com/succinctlabs/telepathy-contracts/blob/main/test/libraries/SimpleSerialize.t.sol
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {SSZ} from "../SimpleSerialize.sol";

contract SSZTest is Test {
    function testFinalityProofWhenEthereum() public pure {
        uint256 index = 105;
        bytes32[] memory branch = new bytes32[](6);
        branch[0] = bytes32(0xe424020000000000000000000000000000000000000000000000000000000000);
        branch[1] = bytes32(0x75410a8f37f9506fb3f972cce6ece955e381e51037e432ce4ca47479c9cd9158);
        branch[2] = bytes32(0xe6af38835c0ac3c2b0d561dfaec168171d7d77c1c2e8e74ff9b1891cf43faf8d);
        branch[3] = bytes32(0x3e4fb2d12bd835bc6ee23b5ec65a43f4493e32f5ef45d46bd2c38830b17672bb);
        branch[4] = bytes32(0x880548f4df2d4003f7be2fbbde112eb46b8f756b5e33202e04863000e4383f3b);
        branch[5] = bytes32(0x88475251bcec25245a44bddd92b2c36db6c9c48bc6d91b5d0da78af3229ff783);
        bytes32 root = bytes32(0xe81a65c5c0f2a36e40b6872fcfdd62dbb67d47f3d49a6b978c0d4440341e723f);
        bytes32 leaf = bytes32(0xd85d3181f1178b07e89691aa2bfcd4d88837f011fcda3326b4ce9a68ec6d9e44);
        assertTrue(SSZ.isValidMerkleProof(leaf, index, branch, root));
    }

    function testMerkleizeLeavesBalanced() public pure {
        // Define leaves
        bytes32[] memory leaves = new bytes32[](8);
        leaves[0] = bytes32(0xe424020000000000000000000000000000000000000000000000000000000000);
        leaves[1] = bytes32(0x75410a8f37f9506fb3f972cce6ece955e381e51037e432ce4ca47479c9cd9158);
        leaves[2] = bytes32(0xe6af38835c0ac3c2b0d561dfaec168171d7d77c1c2e8e74ff9b1891cf43faf8d);
        leaves[3] = bytes32(0x3e4fb2d12bd835bc6ee23b5ec65a43f4493e32f5ef45d46bd2c38830b17672bb);
        leaves[4] = bytes32(0x880548f4df2d4003f7be2fbbde112eb46b8f756b5e33202e04863000e4383f3b);
        leaves[5] = bytes32(0x88475251bcec25245a44bddd92b2c36db6c9c48bc6d91b5d0da78af3229ff783);
        leaves[6] = bytes32(0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563);
        leaves[7] = bytes32(0x5b38da6a701c568545dcfcb03fcb875f56beddc4f7fe24ebd506de91ee3134b2);
        // Define intermediate nodes
        bytes32 h01 = sha256(bytes.concat(leaves[0], leaves[1]));
        bytes32 h23 = sha256(bytes.concat(leaves[2], leaves[3]));
        bytes32 h45 = sha256(bytes.concat(leaves[4], leaves[5]));
        bytes32 h67 = sha256(bytes.concat(leaves[6], leaves[7]));
        bytes32 h0123 = sha256(bytes.concat(h01, h23));
        bytes32 h4567 = sha256(bytes.concat(h45, h67));
        // Assert
        bytes32 expectedRoot = sha256(bytes.concat(h0123, h4567));
        bytes32 computedRoot = SSZ.merkleize(leaves);
        assertEq(expectedRoot, computedRoot, "Merkleize mismatch");
    }
}
