// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {BLST} from "../utils/BLST.sol";
import {
    Validator,
    ValidatorSets,
    ValidatorSetMerkleAttestation,
    ValidatorSetMerkleCommitment
} from "../utils/ValidatorSets.sol";

contract ValidatorSetsTestHarness {
    function parseMerkleAttestation(
        bytes calldata attestationBytes
    ) public pure returns (ValidatorSetMerkleAttestation memory) {
        return ValidatorSets.parseMerkleAttestation(attestationBytes);
    }

    function parseMerkleCommitment(
        bytes calldata data
    ) public pure returns (ValidatorSetMerkleCommitment memory) {
        return ValidatorSets.parseMerkleCommitment(data);
    }
}

contract ValidatorSetsTest is Test {
    ValidatorSetsTestHarness private _harness;

    function setUp() public {
        _harness = new ValidatorSetsTestHarness();
    }

    /*
    * @dev Test to make sure a round trip of serialization is a no-op for a Merkle attestation.
    */
    function testRoundTripMerkleAttestation(
        uint256 numSignersRaw,
        uint256 numProofHashesRaw
    ) public view {
        uint256 numSigners = bound(numSignersRaw, 1, 10);
        uint256 numProofHashes = bound(numProofHashesRaw, 1, 20);
        // Build signers
        Validator[] memory signers = new Validator[](numSigners);
        for (uint256 i; i < numSigners; i++) {
            signers[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(i + 1),
                weight: uint64((i + 1) % 2)
            });
        }
        // Build dummy proof hashes
        bytes32[] memory proof = new bytes32[](numProofHashes);
        for (uint256 i; i < numProofHashes; i++) {
            proof[i] = sha256(abi.encodePacked("proof", i));
        }
        // Build dummy proof flags.
        uint256 numFlags = numSigners + numProofHashes - 1;
        bool[] memory proofFlags = new bool[](numFlags);
        for (uint256 i; i < numFlags; i++) {
            proofFlags[i] = (i % 2 == 0);
        }
        // Build aggregate dummy signature
        bytes memory aggregateBlsSig = new bytes(BLST.BLS_SIGNATURE_LENGTH);
        for (uint256 i; i < BLST.BLS_SIGNATURE_LENGTH; i++) {
            aggregateBlsSig[i] = bytes1(uint8(i & 0xff));
        }
        ValidatorSetMerkleAttestation memory attestation = ValidatorSetMerkleAttestation({
            signers: signers,
            proof: proof,
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateBlsSig
        });
        // Serialize
        bytes memory serialized = ValidatorSets.serializeMerkleAttestation(attestation);
        ValidatorSetMerkleAttestation memory deserialized =
            _harness.parseMerkleAttestation(serialized);
        // Test
        assertEq(deserialized.signers.length, numSigners);
        for (uint256 i; i < numSigners; i++) {
            assertEq(attestation.signers[i].blsPublicKey, deserialized.signers[i].blsPublicKey);
            assertEq(attestation.signers[i].weight, deserialized.signers[i].weight);
        }
        assertEq(deserialized.proof.length, numProofHashes);
        for (uint256 i; i < numProofHashes; i++) {
            assertEq(attestation.proof[i], deserialized.proof[i]);
        }
        assertEq(deserialized.proofFlags.length, numFlags);
        for (uint256 i; i < numFlags; i++) {
            assertEq(attestation.proofFlags[i], deserialized.proofFlags[i]);
        }
        assertEq(attestation.aggregateBlsSig, deserialized.aggregateBlsSig);
    }

    /**
     * @dev Serialize/parse roundtrip for ValidatorSetMerkleCommitment payloads.
     */
    function testRoundtripMerkleCommitment() public view {
        ValidatorSetMerkleCommitment memory original = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: 0x1234567890123456789012345678901234567890123456789012345678901234,
            root: 0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef,
            totalWeight: 1_000_000,
            pChainHeight: 42,
            pChainTimestamp: 1_700_000_000
        });
        bytes memory serialized = ValidatorSets.serializeMerkleCommitment(original);
        ValidatorSetMerkleCommitment memory parsed = _harness.parseMerkleCommitment(serialized);

        assertEq(parsed.avalancheBlockchainID, original.avalancheBlockchainID);
        assertEq(parsed.root, original.root);
        assertEq(parsed.totalWeight, original.totalWeight);
        assertEq(parsed.pChainHeight, original.pChainHeight);
        assertEq(parsed.pChainTimestamp, original.pChainTimestamp);
    }
}
