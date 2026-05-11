// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {MerkleValidatorSetRegistry} from "../MerkleValidatorSetRegistry.sol";
import {BLST} from "../utils/BLST.sol";
import {Validator, ValidatorSetMerkleAttestation, ValidatorSets} from "../utils/ValidatorSets.sol";
import {
    TeleporterMessageV2,
    TeleporterMessageV2Parsing,
    TeleporterICMMessage
} from "../../common/TeleporterMessageV2.sol";

contract MerkleValidatorSetRegistryCommon is Test {
    uint32 public constant NETWORK_ID = 1;
    bytes32 public constant PCHAIN_BLOCKCHAIN_ID =
        0x3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7;
    bytes32 public constant UNREGISTERED_BLOCKCHAIN_ID =
        0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef;

    function dummyPChainValidatorSetSecretKeys() public pure returns (uint256[] memory) {
        uint256[] memory secretKeys = new uint256[](4);
        secretKeys[0] = 2;
        secretKeys[1] = 3;
        secretKeys[2] = 4;
        secretKeys[3] = 5;
        return secretKeys;
    }
}

/**
 * @dev Tests for the MerkleValidatorSetRegistry.verifyMessage pipeline.
 * The registry is bootstrapped at construction with a single validator set committed
 * under pChainID and these tests verify messages against that set.
 *
 * TODO: Tests covering verification against L1-registered sets require the ability to register
 * validator set commitments under arbitrary blockchain IDs, which is currently not implemented.
 */
contract MerkleValidatorSetRegistryVerifyMessageTest is MerkleValidatorSetRegistryCommon {
    MerkleValidatorSetRegistry private _registry;
    Validator[] private _validators;

    function setUp() public {
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
        // Build a 4-validator set
        Validator[] memory validators = new Validator[](4);
        uint64 totalWeight = 0;
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        for (uint256 i = 0; i < 4; i++) {
            validators[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(secretKeys[i]),
                weight: uint64(i + 1)
            });
            assertEq(
                BLST.comparePublicKeys(
                    BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey),
                    previousPublicKey
                ),
                1
            );
            previousPublicKey = BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey);
            totalWeight += validators[i].weight;
            _validators.push(validators[i]);
        }
        // Initialize
        _registry = new MerkleValidatorSetRegistry({
            avalancheNetworkID_: NETWORK_ID,
            pChainID_: PCHAIN_BLOCKCHAIN_ID,
            pChainGenesisRoot: _buildRoot(validators),
            pChainTotalWeight: totalWeight,
            pChainHeight: 1,
            pChainTimestamp: 1
        });
    }

    /// @dev verifyMessage reverts when the source blockchain ID has no commitment.
    function testVerifyMessageUnregisteredSourceReverts() public {
        TeleporterICMMessage memory message = _emptyMessage();
        message.sourceNetworkID = NETWORK_ID;
        message.sourceBlockchainID = UNREGISTERED_BLOCKCHAIN_ID;

        vm.expectRevert(bytes("No validator set registered to given ID"));
        _registry.verifyMessage(message);
    }

    /// @dev verifyMessage reverts when the attestation contains zero signers.
    function testVerifyMessageEmptySignersReverts() public {
        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: new Validator[](0),
            proof: new bytes32[](0),
            proofFlags: new bool[](0),
            aggregateBlsSig: new bytes(BLST.BLS_SIGNATURE_LENGTH)
        });
        TeleporterICMMessage memory message = _emptyMessage();
        message.sourceNetworkID = NETWORK_ID;
        message.sourceBlockchainID = PCHAIN_BLOCKCHAIN_ID;
        message.attestation = ValidatorSets.serializeMerkleAttestation(att);

        vm.expectRevert(bytes("No signers"));
        _registry.verifyMessage(message);
    }

    /// @dev verifyMessage returns false when the Merkle proof does not reconstruct the stored root.
    function testVerifyMessageBadMerkleProofReturnsFalse() public {
        Validator[] memory signers = new Validator[](1);
        signers[0] = _validators[0];

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: signers,
            proof: new bytes32[](0),
            proofFlags: new bool[](0),
            aggregateBlsSig: new bytes(BLST.BLS_SIGNATURE_LENGTH)
        });

        TeleporterICMMessage memory message = _emptyMessage();
        message.sourceNetworkID = NETWORK_ID;
        message.sourceBlockchainID = PCHAIN_BLOCKCHAIN_ID;
        message.attestation = ValidatorSets.serializeMerkleAttestation(att);

        assertFalse(_registry.verifyMessage(message));
    }

    /**
     * @dev verifyMessage returns true when all 4 validators sign and the proof reconstructs the stored root.
     *           root
     *          /    \
     *        AB      CD
     *       /  \    /  \
     *      L0  L1  L2  L3
     */
    function testVerifyMessageSuccessAllSignersBalancedTree() public view {
        // Construct unsigned warp message
        TeleporterMessageV2 memory inner;
        bytes memory innerSerialized =
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(inner);
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, PCHAIN_BLOCKCHAIN_ID, address(_registry), innerSerialized
        );

        // Sign the message
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
        bytes memory aggregateBlsSig = BLST.createAggregateSignature(secretKeys, signedData);

        // Build the multi-proof. The proof is empty since all validators are signers, so the flags simply instruct the verifier to hash up the tree from the leaves.
        // Recall the number of flags is computed as #leaves + #proofHashes - 1.
        bool[] memory proofFlags = new bool[](3);
        proofFlags[0] = true;
        proofFlags[1] = true;
        proofFlags[2] = true;

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: _validators,
            proof: new bytes32[](0),
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateBlsSig
        });

        TeleporterICMMessage memory message = TeleporterICMMessage({
            message: inner,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });

        assertTrue(_registry.verifyMessage(message));
    }

    /**
     * @dev verifyMessage returns true with a partial signer set requiring a non-empty
     * Merkle multi-proof. Signers are validators 0, 2, 3 (skipping validator 1), so
     * L1 must be supplied as an external sibling hash.
     *
     *           root
     *          /    \
     *        AB      CD
     *       /  \    /  \
     *      L0  L1  L2  L3
     *
     * Signers = [L0, L2, L3], proof = [L1]
     *
     * Steps to compute proof; #flags = #leaves + #proofHashes - 1 = 3 + 1 - 1 flags
     *   1. AB = hash(L0, L1)  indicates to take L0 from leaves and L1 from proof, so the flag is false
     *   2. CD   = hash(L2, L3) indicates to take both from leaves, so the flag is true
     *   3. root = hash(AB, CD)  indicates to take both sibling nodes, so the flag is true
     */
    function testVerifyMessageSuccessPartialSignersBalancedTree() public view {
        // Reconstruct the unsigned warp message
        TeleporterMessageV2 memory inner;
        bytes memory innerSerialized =
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(inner);
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, PCHAIN_BLOCKCHAIN_ID, address(_registry), innerSerialized
        );

        // Sign the message
        uint256[] memory secretKeys = new uint256[](3);
        secretKeys[0] = 2;
        secretKeys[1] = 4;
        secretKeys[2] = 5;
        bytes memory aggregateSig = BLST.createAggregateSignature(secretKeys, signedData);

        // Add validators who signed the message
        Validator[] memory signers = new Validator[](3);
        signers[0] = _validators[0];
        signers[1] = _validators[2];
        signers[2] = _validators[3];

        // Build the proof
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = sha256(abi.encodePacked(_validators[1].blsPublicKey, _validators[1].weight));

        // Flags per combination step, see comment above
        bool[] memory proofFlags = new bool[](3);
        proofFlags[0] = false;
        proofFlags[1] = true;
        proofFlags[2] = true;

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: signers,
            proof: proof,
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateSig
        });

        TeleporterICMMessage memory message = TeleporterICMMessage({
            message: inner,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });
        assertTrue(_registry.verifyMessage(message));
    }

    /**
     * @dev Helper function that builds the Merkle root over the validator leaves using the same
     * scheme that verifyMerkleAttestation expects. The leaves are computed as leaf := sha256(pubkey || weight),
     * where the public keys are lexiographically sorted.
     * NOTE: Assumes number of validators is a power of 2 otherwise this won't match OZ's MerkleProof verifier.
     */
    function _buildRoot(
        Validator[] memory validators
    ) internal pure returns (bytes32) {
        uint256 n = validators.length;
        bytes32[] memory layer = new bytes32[](n);
        for (uint256 i = 0; i < n; i++) {
            layer[i] = sha256(abi.encodePacked(validators[i].blsPublicKey, validators[i].weight));
        }
        while (layer.length > 1) {
            uint256 nextLen = (layer.length + 1) / 2;
            bytes32[] memory next = new bytes32[](nextLen);
            for (uint256 i = 0; i < nextLen; i++) {
                if (2 * i + 1 < layer.length) {
                    bytes32 a = layer[2 * i];
                    bytes32 b = layer[2 * i + 1];
                    next[i] =
                        a < b ? sha256(abi.encodePacked(a, b)) : sha256(abi.encodePacked(b, a));
                } else {
                    next[i] = layer[2 * i];
                }
            }
            layer = next;
        }
        return layer[0];
    }

    function _emptyMessage() internal pure returns (TeleporterICMMessage memory) {
        TeleporterMessageV2 memory inner;
        return TeleporterICMMessage({
            message: inner,
            sourceNetworkID: 0,
            sourceBlockchainID: bytes32(0),
            attestation: new bytes(0)
        });
    }
}
