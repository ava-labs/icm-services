// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {MerkleValidatorSetRegistry} from "../MerkleValidatorSetRegistry.sol";
import {BLST} from "../utils/BLST.sol";
import {
    Validator,
    ValidatorSetMerkleAttestation,
    ValidatorSets,
    ValidatorSetMerkleCommitment
} from "../utils/ValidatorSets.sol";
import {ICMMessage} from "../../common/ICM.sol";
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

    /**
     * @dev Builds the bootstrap 4-validator P-chain set, asserts they are lexicographically
     * sorted by BLS public key, deploys a MerkleValidatorSetRegistry committed to that set,
     * and pushes the validators into the provided storage array.
     */
    function _setUpRegistryWithPChainValidators(
        Validator[] storage validatorsOut
    ) internal returns (MerkleValidatorSetRegistry) {
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
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
            validatorsOut.push(validators[i]);
        }

        return new MerkleValidatorSetRegistry({
            avalancheNetworkID_: NETWORK_ID,
            pChainID_: PCHAIN_BLOCKCHAIN_ID,
            pChainGenesisRoot: _buildRoot(validators),
            pChainTotalWeight: totalWeight,
            pChainHeight: 1,
            pChainTimestamp: 1,
            allowPChainFallback_: true
        });
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

/**
 * @dev Tests for the MerkleValidatorSetRegistry.verifyMessage pipeline.
 * The registry is bootstrapped at construction with a single validator set committed
 * under pChainID and these tests verify messages against that set.
 */
contract MerkleValidatorSetRegistryVerifyMessageTest is MerkleValidatorSetRegistryCommon {
    MerkleValidatorSetRegistry private _registry;
    Validator[] private _validators;

    function setUp() public {
        _registry = _setUpRegistryWithPChainValidators(_validators);
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
}

/**
 * @dev Reuses the bootstrapped P-chain validator set from the common setUp. Register tests prove new chains can be
 * added under P-chain signatures, update tests prove an existing chain's commitment
 * can be replaced under its own validators' signatures.
 */
contract MerkleValidatorSetRegistryRegisterUpdateTest is MerkleValidatorSetRegistryCommon {
    bytes32 internal constant _NEW_CHAIN_ID =
        0xabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd;

    MerkleValidatorSetRegistry private _registry;
    Validator[] private _pchainValidators;

    function setUp() public {
        _registry = _setUpRegistryWithPChainValidators(_pchainValidators);
    }

    /// @dev First registration of a new chain where P-chain validators sign the new commitment.
    function testRegisterNewChainSuccess() public {
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });

        ICMMessage memory message = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(newCommitment),
            dummyPChainValidatorSetSecretKeys()
        );

        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.avalancheBlockchainID, _NEW_CHAIN_ID);
        assertEq(stored.root, newCommitment.root);
        assertEq(stored.totalWeight, newCommitment.totalWeight);
        assertEq(stored.pChainHeight, newCommitment.pChainHeight);
        assertEq(stored.pChainTimestamp, newCommitment.pChainTimestamp);
    }

    /// @dev registerValidatorSet reverts when the message's network ID doesn't match.
    function testRegisterRevertsOnNetworkMismatch() public {
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory message = ICMMessage({
            rawMessage: ValidatorSets.serializeMerkleCommitment(newCommitment),
            sourceNetworkID: NETWORK_ID + 1, // wrong
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Network ID mismatch"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev Update of P-chain validator set, signed by the currently stored P-chain validators.
    /// Here we update the P-chain's own commitment, since it's already registered at construction.
    function testUpdateExistingChainSuccess() public {
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory message = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(newCommitment),
            dummyPChainValidatorSetSecretKeys()
        );

        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(PCHAIN_BLOCKCHAIN_ID);
        assertEq(stored.root, newCommitment.root);
        assertEq(stored.totalWeight, newCommitment.totalWeight);
        assertEq(stored.pChainHeight, newCommitment.pChainHeight);
    }

    /// @dev regValidatorSet reverts when the payload chain isn't yet registered.
    function testUpdateRevertsOnUnregisteredChain() public {
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: UNREGISTERED_BLOCKCHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory message = ICMMessage({
            rawMessage: ValidatorSets.serializeMerkleCommitment(newCommitment),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: UNREGISTERED_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Initial registration must be signed by the P-Chain"));
        _registry.registerValidatorSet(message, UNREGISTERED_BLOCKCHAIN_ID);
    }

    /// @dev Once a chain is registered, attempts to update it with a signing chain that is
    /// neither the chain itself nor the P-Chain are rejected.
    function testRegisterRevertsOnInvalidSigningChain() public {
        // First register _NEW_CHAIN_ID
        ValidatorSetMerkleCommitment memory initial = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory firstMessage = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(initial),
            dummyPChainValidatorSetSecretKeys()
        );
        _registry.registerValidatorSet(firstMessage, PCHAIN_BLOCKCHAIN_ID);

        // Attempt to update _NEW_CHAIN_ID with a wrong signing chain that is not
        // the chain itself or PCHAIN_BLOCKCHAIN_ID
        ValidatorSetMerkleCommitment memory update = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xcafe)),
            totalWeight: 20,
            pChainHeight: 5,
            pChainTimestamp: 5
        });
        ICMMessage memory secondMessage = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(update),
            dummyPChainValidatorSetSecretKeys()
        );

        vm.expectRevert(bytes("Invalid signing chain"));
        _registry.registerValidatorSet(secondMessage, UNREGISTERED_BLOCKCHAIN_ID);
    }

    /// @dev Reverts when caller specifies the target chain as signer but the target isn't registered yet.
    function testRegisterRevertsOnSelfSignedFirstTimeRegistration() public {
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory message = _signedRegistrationMessage(
            _NEW_CHAIN_ID, // source is the new chain itself
            ValidatorSets.serializeMerkleCommitment(newCommitment),
            dummyPChainValidatorSetSecretKeys()
        );

        // Can't self-sign before being registered
        vm.expectRevert(bytes("Initial registration must be signed by the P-Chain"));
        _registry.registerValidatorSet(message, _NEW_CHAIN_ID);
    }

    /// @dev registerValidatorSet reverts when the attestation's signers are not included
    /// against the registered chain's stored Merkle root. Here we attempt to update P-chain's
    /// commitment but sign with an unregistered set of validators. The multi-proof
    /// can't reconstruct the stored P-chain root from these leaves.
    function testUpdateRevertsOnWrongSigners() public {
        uint256[] memory wrongSecretKeys = new uint256[](4);
        wrongSecretKeys[0] = 6;
        wrongSecretKeys[1] = 7;
        wrongSecretKeys[2] = 8;
        wrongSecretKeys[3] = 9;

        Validator[] memory wrongValidators = new Validator[](4);
        for (uint256 i = 0; i < 4; i++) {
            wrongValidators[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(wrongSecretKeys[i]),
                weight: uint64(i + 1)
            });
        }

        // New wrong commitment
        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            root: bytes32(uint256(0xbeef)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        bytes memory rawMessage = ValidatorSets.serializeMerkleCommitment(newCommitment);

        // Sign with the wrong validator keys
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, PCHAIN_BLOCKCHAIN_ID, address(0), rawMessage
        );
        bytes memory aggregateBlsSig = BLST.createAggregateSignature(wrongSecretKeys, signedData);

        bool[] memory proofFlags = new bool[](3);
        proofFlags[0] = true;
        proofFlags[1] = true;
        proofFlags[2] = true;

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: wrongValidators,
            proof: new bytes32[](0), // all validators sign so multi-proof is empty
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateBlsSig
        });

        ICMMessage memory message = ICMMessage({
            rawMessage: rawMessage,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });

        // Test
        vm.expectRevert(bytes("Failed to verify attestation"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /**
     * @dev registerValidatorSet reverts when the signing validators are correct, but
     * quorum threshold is not reached.
     *
     * Signers are validators 0 and 1 with weights 1 and 2 respectively against total weight 10.
     * This puts quorum threshold at 7. The Merkle multi-proof reconstructs the
     * stored root correctly, so verification proceeds to the weight check and fails there.
     *
     * Below, CD is in the external proof, and L0 and L1 are the signers.
     *           root
     *          /    \
     *        AB      CD
     *       /  \    /  \
     *      L0  L1  L2  L3
     */
    function testUpdateRevertsBelowQuorum() public {
        // Sign with validators 0 and 1 only
        uint256[] memory partialSecretKeys = new uint256[](2);
        partialSecretKeys[0] = 2;
        partialSecretKeys[1] = 3;

        Validator[] memory signers = new Validator[](2);
        signers[0] = _pchainValidators[0];
        signers[1] = _pchainValidators[1];

        ValidatorSetMerkleCommitment memory newCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 3, // below quorum threshold
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        bytes memory rawMessage = ValidatorSets.serializeMerkleCommitment(newCommitment);

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, PCHAIN_BLOCKCHAIN_ID, address(0), rawMessage
        );
        bytes memory aggregateBlsSig = BLST.createAggregateSignature(partialSecretKeys, signedData);

        // Build the multi-proof
        bytes32 l2 =
            sha256(abi.encodePacked(_pchainValidators[2].blsPublicKey, _pchainValidators[2].weight));
        bytes32 l3 =
            sha256(abi.encodePacked(_pchainValidators[3].blsPublicKey, _pchainValidators[3].weight));
        bytes32 cd = l2 < l3 ? sha256(abi.encodePacked(l2, l3)) : sha256(abi.encodePacked(l3, l2));
        bytes32[] memory proof = new bytes32[](1);
        proof[0] = cd;

        // Set flags
        bool[] memory proofFlags = new bool[](2);
        proofFlags[0] = true;
        proofFlags[1] = false;

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: signers,
            proof: proof,
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateBlsSig
        });

        ICMMessage memory message = ICMMessage({
            rawMessage: rawMessage,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });

        // Test
        vm.expectRevert(bytes("Failed to verify attestation"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /**
     * @dev The fallback path: an L1's validator set commitment is updated via P-Chain
     * signatures rather than the L1's own validators. This covers the scenario where
     * an intermediate update was missed and the L1's previous validators can no longer
     * form a quorum.
     */
    function testUpdateRegisteredChainViaPChainSignature() public {
        // First, register a new L1 (signed by P-Chain — initial registration)
        ValidatorSetMerkleCommitment memory initial = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory firstMessage = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(initial),
            dummyPChainValidatorSetSecretKeys()
        );
        _registry.registerValidatorSet(firstMessage, PCHAIN_BLOCKCHAIN_ID);

        // Now update that L1's commitment, signed by P-Chain rather than the L1's validators
        ValidatorSetMerkleCommitment memory updated = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xcafe)),
            totalWeight: 20,
            pChainHeight: 5,
            pChainTimestamp: 5
        });
        ICMMessage memory secondMessage = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(updated),
            dummyPChainValidatorSetSecretKeys()
        );
        _registry.registerValidatorSet(secondMessage, PCHAIN_BLOCKCHAIN_ID);

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.root, updated.root);
        assertEq(stored.totalWeight, updated.totalWeight);
        assertEq(stored.pChainHeight, updated.pChainHeight);
    }

    /**
     * @dev Happy path for the L1 self-signed update flow:
     *   1. The P-Chain registers a new L1 with a Merkle root over a validator set.
     *   2. The L1's own validators sign an update to the L1's commitment.
     *   3. The contract is updated.
     */
    function testL1SelfSignedUpdateSuccess() public {
        // Build L1 validator set
        uint256[] memory l1SecretKeys = new uint256[](4);
        l1SecretKeys[0] = 11;
        l1SecretKeys[1] = 12;
        l1SecretKeys[2] = 13;
        l1SecretKeys[3] = 10;
        Validator[] memory l1Validators = new Validator[](l1SecretKeys.length);
        uint64 l1TotalWeight = 0;
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        for (uint256 i = 0; i < l1SecretKeys.length; i++) {
            l1Validators[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(l1SecretKeys[i]),
                weight: uint64(i + 1)
            });
            assertEq(
                BLST.comparePublicKeys(
                    BLST.unPadUncompressedBlsPublicKey(l1Validators[i].blsPublicKey),
                    previousPublicKey
                ),
                1,
                "L1 keys must be lexicographically sorted"
            );
            previousPublicKey = BLST.unPadUncompressedBlsPublicKey(l1Validators[i].blsPublicKey);
            l1TotalWeight += l1Validators[i].weight;
        }
        bytes32 l1Root = _buildRoot(l1Validators);

        // Register validator set
        ValidatorSetMerkleCommitment memory initialCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: l1Root,
            totalWeight: l1TotalWeight,
            pChainHeight: 2,
            pChainTimestamp: 2
        });
        ICMMessage memory registrationMessage = _signedRegistrationMessage(
            PCHAIN_BLOCKCHAIN_ID,
            ValidatorSets.serializeMerkleCommitment(initialCommitment),
            dummyPChainValidatorSetSecretKeys()
        );
        _registry.registerValidatorSet(registrationMessage, PCHAIN_BLOCKCHAIN_ID);

        // Check the L1 is now registered
        ValidatorSetMerkleCommitment memory afterRegistration =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(afterRegistration.root, l1Root);
        assertEq(afterRegistration.totalWeight, l1TotalWeight);

        // Construct the update signed by the L1's own validators
        ValidatorSetMerkleCommitment memory updatedCommitment = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: _NEW_CHAIN_ID,
            root: bytes32(uint256(0xababa)),
            totalWeight: 10,
            pChainHeight: 5,
            pChainTimestamp: 5
        });
        bytes memory rawUpdate = ValidatorSets.serializeMerkleCommitment(updatedCommitment);
        bytes memory signedData =
            ValidatorSets.buildUnsignedWarpMessage(NETWORK_ID, _NEW_CHAIN_ID, address(0), rawUpdate);
        bytes memory aggregateSig = BLST.createAggregateSignature(l1SecretKeys, signedData);
        bool[] memory proofFlags = new bool[](3);
        proofFlags[0] = true;
        proofFlags[1] = true;
        proofFlags[2] = true;
        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: l1Validators,
            proof: new bytes32[](0),
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateSig
        });
        ICMMessage memory updateMessage = ICMMessage({
            rawMessage: rawUpdate,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: _NEW_CHAIN_ID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });

        // Submit the update
        _registry.registerValidatorSet(updateMessage, _NEW_CHAIN_ID);
        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.root, updatedCommitment.root);
        assertEq(stored.totalWeight, updatedCommitment.totalWeight);
        assertEq(stored.pChainHeight, updatedCommitment.pChainHeight);
        assertEq(stored.pChainTimestamp, updatedCommitment.pChainTimestamp);
    }

    /**
     * @dev Builds an ICMMessage signed by all 4 P-chain validators, suitable for register/update
     * calls. This uses address(0) as origin sender which matches what registerValidatorSet
     * reconstructs in its warp preimage.
     */
    function _signedRegistrationMessage(
        bytes32 sourceBlockchainID,
        bytes memory rawMessage,
        uint256[] memory secretKeys
    ) internal view returns (ICMMessage memory) {
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, sourceBlockchainID, address(0), rawMessage
        );
        bytes memory aggregateBlsSig = BLST.createAggregateSignature(secretKeys, signedData);

        bool[] memory proofFlags = new bool[](_pchainValidators.length - 1);
        for (uint256 i = 0; i < proofFlags.length; i++) {
            proofFlags[i] = true;
        }

        ValidatorSetMerkleAttestation memory att = ValidatorSetMerkleAttestation({
            signers: _pchainValidators,
            proof: new bytes32[](0),
            proofFlags: proofFlags,
            aggregateBlsSig: aggregateBlsSig
        });

        return ICMMessage({
            rawMessage: rawMessage,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: sourceBlockchainID,
            attestation: ValidatorSets.serializeMerkleAttestation(att)
        });
    }
}
