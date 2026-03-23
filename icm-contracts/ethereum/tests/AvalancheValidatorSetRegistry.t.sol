// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {AvalancheValidatorSetRegistry} from "../AvalancheValidatorSetRegistry.sol";
import {ICMMessage} from "../../common/ICM.sol";
import {BLST} from "../utils/BLST.sol";
import {DiffUpdater} from "../DiffUpdater.sol";
import {SubsetUpdater} from "../SubsetUpdater.sol";
import {
    Validator,
    ValidatorSet,
    ValidatorSets,
    ValidatorChange,
    ValidatorSetDiff,
    ValidatorSetMetadata,
    ValidatorSetShard,
    ValidatorSetSignature
} from "../utils/ValidatorSets.sol";

// Common utility functions and fixtures for the suites in this file
contract AvalancheValidatorSetRegistryCommon is Test {
    struct RegistryTestCase {
        AvalancheValidatorSetRegistry registry;
        bytes shardBytes;
        ICMMessage message;
    }

    struct RegistryTestCaseMultiStep {
        AvalancheValidatorSetRegistry registry;
        bytes[] shardBytes;
        ICMMessage[] messages;
    }

    uint32 public constant NETWORK_ID = 1;
    bytes32 public constant PCHAIN_BLOCKCHAIN_ID =
        0x3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7;
    bytes32 public constant L1_BLOCKCHAIN_ID =
        0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef;

    /**
     * @dev Create a dummy set of P-chain validators to initialize the `AvalancheValidatorSetRegistry` with.
     * @return Returns the validator set and hash of the validators
     */
    function dummyPChainValidatorSet() public view returns (ValidatorSet memory, bytes32) {
        Validator[] memory validators = new Validator[](5);
        uint64 totalWeight = 0;

        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();

        for (uint256 i = 0; i < 5; i++) {
            validators[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(secretKeys[i]),
                weight: uint64(i + 1)
            });
            // check that the validators are ordered by public key
            assertEq(
                BLST.comparePublicKeys(
                    BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey),
                    previousPublicKey
                ),
                1
            );
            previousPublicKey = BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey);
            totalWeight += validators[i].weight;
        }

        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: 0,
            pChainTimestamp: 0
        });
        return (validatorSet, sha256(ValidatorSets.serializeValidators(validators)));
    }

    /**
     * @dev Create a dummy set of P-chain validators for diffs to initialize the `AvalancheValidatorSetRegistry` with.
     * @return Returns the validator set and hash of the validators
     */
    function dummyPChainValidatorSetForDiff() public view returns (ValidatorSet memory, bytes32) {
        (ValidatorSet memory validatorSet, bytes32 hash) = dummyPChainValidatorSet();
        validatorSet.pChainHeight = 1;
        validatorSet.pChainTimestamp = 1;
        return (validatorSet, hash);
    }

    /**
     * @dev A fixture that returns a validator set along with a raw ICM message to register
     * this set. This set is split into two shards with one validator each.
     */
    function registerValidatorSetFixture(
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) public view returns (Validator[] memory, bytes memory) {
        // create the total validator set and two shards
        Validator[] memory validators = new Validator[](2);
        bytes32[] memory shardHashes = new bytes32[](2);

        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        uint64 totalWeight = 0;
        for (uint256 i = 0; i <= 1; i++) {
            validators[i] =
                Validator({blsPublicKey: BLST.getPublicKeyFromSecret(i + 2), weight: uint64(i + 2)});
            // check that the validators are ordered by public key
            assertEq(
                BLST.comparePublicKeys(
                    BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey),
                    previousPublicKey
                ),
                1
            );
            previousPublicKey = BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey);
            totalWeight += validators[i].weight;
            // compute the shard hash
            Validator[] memory validatorShard = new Validator[](1);
            validatorShard[0] = validators[i];
            shardHashes[i] = sha256(ValidatorSets.serializeValidators(validatorShard));
        }

        ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
            avalancheBlockchainID: L1_BLOCKCHAIN_ID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            shardHashes: shardHashes
        });

        bytes memory raw = ValidatorSets.serializeValidatorSetMetadata(metadata);

        return (validators, raw);
    }

    /**
     * @dev A fixture that returns a validator set along with a raw ICM message to register
     * this set. This set is split into two validator diff shards with one validator each.
     */
    function registerValidatorSetDiffFixture(
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint256 secretKeyOffset,
        Validator[] memory existingValidators
    ) public view returns (Validator[] memory, bytes memory) {
        Validator[] memory newValidators = new Validator[](2);
        bytes32[] memory shardHashes = new bytes32[](2);

        // Create validators
        {
            bytes memory previousPublicKey =
                new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
            for (uint256 i = 0; i <= 1; i++) {
                newValidators[i] = Validator({
                    blsPublicKey: BLST.getPublicKeyFromSecret(i + secretKeyOffset),
                    weight: uint64(i + secretKeyOffset)
                });
                assertEq(
                    BLST.comparePublicKeys(
                        BLST.unPadUncompressedBlsPublicKey(newValidators[i].blsPublicKey),
                        previousPublicKey
                    ),
                    1
                );
                previousPublicKey =
                    BLST.unPadUncompressedBlsPublicKey(newValidators[i].blsPublicKey);
            }
        }
        // Shard 1: remove existing validators (if any), add newValidators[0]
        shardHashes[0] = _computeDiffShard1Hash(
            existingValidators, newValidators[0], pChainHeight, pChainTimestamp
        );
        // Shard 2: add newValidators[1]
        shardHashes[1] = _computeDiffShard2Hash(
            newValidators[0], newValidators[1], pChainHeight, pChainTimestamp
        );
        // Metadata
        {
            ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
                avalancheBlockchainID: L1_BLOCKCHAIN_ID,
                pChainHeight: pChainHeight,
                pChainTimestamp: pChainTimestamp,
                shardHashes: shardHashes
            });
            bytes memory raw = ValidatorSets.serializeValidatorSetMetadata(metadata);
            return (newValidators, raw);
        }
    }

    /**
     * @dev If this validator set is being registered for the first time, it must be signed
     * by the P-chain
     */
    function registerValidatorSetInitialFixture()
        public
        view
        returns (Validator[] memory, ICMMessage memory)
    {
        (Validator[] memory validators, bytes memory raw) = registerValidatorSetFixture(1, 1);

        // sign the message
        bytes memory signature = dummyPChainValidatorSetSign(L1_BLOCKCHAIN_ID, raw);
        ICMMessage memory message = ICMMessage({
            rawMessage: raw,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: L1_BLOCKCHAIN_ID,
            attestation: signature
        });
        return (validators, message);
    }

    /**
     * @dev If this validator diff is being registered for the first time, it must be signed
     * by the P-chain
     */
    function registerValidatorSetInitialDiffFixture()
        public
        view
        returns (Validator[] memory, ICMMessage memory)
    {
        (Validator[] memory validators, bytes memory raw) =
            registerValidatorSetDiffFixture(1, 1, 2, new Validator[](0));

        // sign the message
        bytes memory signature = dummyPChainValidatorSetSign(L1_BLOCKCHAIN_ID, raw);
        ICMMessage memory message = ICMMessage({
            rawMessage: raw,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: L1_BLOCKCHAIN_ID,
            attestation: signature
        });
        return (validators, message);
    }

    /**
     * @dev If this validator set is not being registered for the first time, it must be signed
     * by the validator set previously registered to this blockchain ID
     */
    function registerValidatorSetAgainFixture(
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) public view returns (Validator[] memory, ICMMessage memory) {
        (Validator[] memory validators, bytes memory raw) =
            registerValidatorSetFixture(pChainHeight, pChainTimestamp);
        // sign the message
        bytes memory signature = l1ValidatorSetSign(raw);
        ICMMessage memory message = ICMMessage({
            rawMessage: raw,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: L1_BLOCKCHAIN_ID,
            attestation: signature
        });
        return (validators, message);
    }

    function registerValidatorSetAgainDiffFixture(
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) public view returns (Validator[] memory, ICMMessage memory) {
        (Validator[] memory validators, bytes memory raw) =
            registerValidatorSetDiffFixture(pChainHeight, pChainTimestamp, 2, new Validator[](0));
        bytes memory signature = l1ValidatorSetSign(raw);
        ICMMessage memory message = ICMMessage({
            rawMessage: raw,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: L1_BLOCKCHAIN_ID,
            attestation: signature
        });
        return (validators, message);
    }

    function reRegisterValidatorSetDiffFixture(
        Validator[] memory existingValidators,
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) public view returns (Validator[] memory, ICMMessage memory) {
        (Validator[] memory validators, bytes memory raw) =
            registerValidatorSetDiffFixture(pChainHeight, pChainTimestamp, 6, existingValidators);
        bytes memory signature = l1ValidatorSetSign(raw);
        ICMMessage memory message = ICMMessage({
            rawMessage: raw,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: L1_BLOCKCHAIN_ID,
            attestation: signature
        });
        return (validators, message);
    }

    /**
     * @dev Sign the input payload with the dummy validator set created above
     */
    function dummyPChainValidatorSetSign(
        bytes32 avalancheBlockchainID,
        bytes memory payload
    ) public view returns (bytes memory) {
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
        bytes memory signedData =
            _buildUnsignedWarpMessage(NETWORK_ID, avalancheBlockchainID, payload);
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, signedData);
        ValidatorSetSignature memory signature = ValidatorSetSignature({
            // all five validators sign (bits 0-4 set = 0x1F)
            signers: hex"1F",
            signature: rawSig
        });
        return ValidatorSets.serializeValidatorSetSignature(signature);
    }

    /**
     * @dev Sign the input payload with the L1 validator set
     */
    function l1ValidatorSetSign(
        bytes memory payload
    ) public view returns (bytes memory) {
        uint256[] memory secretKeys = new uint256[](2);
        secretKeys[0] = 2;
        secretKeys[1] = 3;
        bytes memory signedData = _buildUnsignedWarpMessage(NETWORK_ID, L1_BLOCKCHAIN_ID, payload);
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, signedData);
        ValidatorSetSignature memory signature = ValidatorSetSignature({
            // both validators sign (bits 0-1 set = 0x03)
            signers: hex"03",
            signature: rawSig
        });
        return ValidatorSets.serializeValidatorSetSignature(signature);
    }

    /**
     * @dev Signs the payload with specific secret keys and explicitly applies the provided bitmask.
     */
    function customValidatorSetSign(
        bytes32 chainID,
        uint256[] memory secretKeys,
        bytes memory signersBitmask,
        bytes memory payload
    ) public view returns (bytes memory) {
        bytes memory signedData = _buildUnsignedWarpMessage(NETWORK_ID, chainID, payload);
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, signedData);
        ValidatorSetSignature memory signature =
            ValidatorSetSignature({signers: signersBitmask, signature: rawSig});
        return ValidatorSets.serializeValidatorSetSignature(signature);
    }

    /**
     * @dev Creates a fully signed ICM Registration message using a custom bitmask.
     */
    function createCustomRegistrationMessage(
        bytes32 chainID,
        ValidatorSetMetadata memory metadata,
        uint256[] memory secretKeys,
        bytes memory signersBitmask
    ) public view returns (ICMMessage memory) {
        bytes memory rawMessage = ValidatorSets.serializeValidatorSetMetadata(metadata);
        bytes memory attestation =
            customValidatorSetSign(chainID, secretKeys, signersBitmask, rawMessage);
        return ICMMessage({
            rawMessage: rawMessage,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: chainID,
            attestation: attestation
        });
    }

    /**
     * @dev A fixture holding the secret keys of the P-chain validator set. These are sorted
     * so that their public counterparts are in increasing order
     */
    function dummyPChainValidatorSetSecretKeys() public pure returns (uint256[] memory) {
        uint256[] memory secretKeys = new uint256[](5);
        secretKeys[0] = 2;
        secretKeys[1] = 3;
        secretKeys[2] = 4;
        secretKeys[3] = 5;
        secretKeys[4] = 1;
        return secretKeys;
    }

    /**
     * @notice Sorts validator changes by their uncompressed public key bytes
     * @dev Uses insertion sort which is efficient for small arrays
     */
    function sortValidatorChanges(
        ValidatorChange[] memory changes
    ) public pure {
        for (uint256 i = 1; i < changes.length; i++) {
            ValidatorChange memory key = changes[i];
            bytes memory keyPubKey = key.blsPublicKey;
            int256 j = int256(i) - 1;
            while (j >= 0) {
                bytes memory jPubKey = changes[uint256(j)].blsPublicKey;
                if (
                    BLST.comparePublicKeys(
                        BLST.unPadUncompressedBlsPublicKey(jPubKey),
                        BLST.unPadUncompressedBlsPublicKey(keyPubKey)
                    ) <= 0
                ) {
                    break;
                }
                changes[uint256(j + 1)] = changes[uint256(j)];
                j--;
            }
            changes[uint256(j + 1)] = key;
        }
    }

    function allAdditionsValidatorSetDiff(
        ValidatorSet memory validatorSet,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) public pure returns (ValidatorSetDiff memory) {
        ValidatorChange[] memory changes = new ValidatorChange[](validatorSet.validators.length);
        for (uint256 i = 0; i < validatorSet.validators.length; i++) {
            changes[i] = ValidatorChange({
                blsPublicKey: validatorSet.validators[i].blsPublicKey,
                weight: validatorSet.validators[i].weight
            });
        }
        sortValidatorChanges(changes);
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp,
            currentHeight: validatorSet.pChainHeight,
            currentTimestamp: validatorSet.pChainTimestamp,
            changes: changes,
            numAdded: uint32(validatorSet.validators.length),
            newSize: validatorSet.validators.length
        });
        return diff;
    }

    function customValidatorSetDiff(
        ValidatorSet memory currentPartialValidatorSet,
        uint64 previousHeight,
        uint64 previousTimestamp,
        uint32 numAdded,
        ValidatorChange[] memory changes
    ) public pure returns (ValidatorSetDiff memory) {
        sortValidatorChanges(changes);
        uint32 numRemoved = 0;
        for (uint256 i = 0; i < changes.length; i++) {
            if (changes[i].weight == 0) {
                numRemoved++;
            }
        }
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: currentPartialValidatorSet.avalancheBlockchainID,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp,
            currentHeight: currentPartialValidatorSet.pChainHeight,
            currentTimestamp: currentPartialValidatorSet.pChainTimestamp,
            changes: changes,
            numAdded: numAdded,
            newSize: 0
        });
        diff.newSize = currentPartialValidatorSet.validators.length + numAdded - numRemoved;
        return diff;
    }

    function _buildUnsignedWarpMessage(
        uint32 networkID,
        bytes32 sourceBlockchainID,
        bytes memory payload
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            bytes2(0),
            networkID,
            sourceBlockchainID,
            uint32(payload.length + 14),
            bytes2(0),
            uint32(1),
            uint32(0),
            uint32(payload.length),
            payload
        );
    }

    function _emptyICMMessage() internal pure returns (ICMMessage memory) {
        return ICMMessage({
            rawMessage: new bytes(0),
            sourceNetworkID: 0,
            sourceBlockchainID: bytes32(0),
            attestation: new bytes(0)
        });
    }

    function _computeDiffShard1Hash(
        Validator[] memory existingValidators,
        Validator memory newValidator,
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) internal pure returns (bytes32) {
        ValidatorChange[] memory changes = new ValidatorChange[](existingValidators.length + 1);
        for (uint256 i = 0; i < existingValidators.length; i++) {
            changes[i] =
                ValidatorChange({blsPublicKey: existingValidators[i].blsPublicKey, weight: 0});
        }
        changes[existingValidators.length] =
            ValidatorChange({blsPublicKey: newValidator.blsPublicKey, weight: newValidator.weight});
        Validator[] memory startSet =
            existingValidators.length > 0 ? existingValidators : new Validator[](0);
        ValidatorSet memory vs = ValidatorSet({
            avalancheBlockchainID: L1_BLOCKCHAIN_ID,
            validators: startSet,
            totalWeight: 0,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
        return sha256(
            ValidatorSets.serializeValidatorSetDiff(
                customValidatorSetDiff(vs, pChainHeight - 1, pChainTimestamp - 1, 1, changes)
            )
        );
    }

    function _computeDiffShard2Hash(
        Validator memory shard1Validator,
        Validator memory newValidator,
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) internal pure returns (bytes32) {
        Validator[] memory shard1Result = new Validator[](1);
        shard1Result[0] = shard1Validator;
        ValidatorSet memory afterShard1 = ValidatorSet({
            avalancheBlockchainID: L1_BLOCKCHAIN_ID,
            validators: shard1Result,
            totalWeight: uint64(shard1Validator.weight),
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
        ValidatorChange[] memory changes = new ValidatorChange[](1);
        changes[0] =
            ValidatorChange({blsPublicKey: newValidator.blsPublicKey, weight: newValidator.weight});
        return sha256(
            ValidatorSets.serializeValidatorSetDiff(
                customValidatorSetDiff(
                    afterShard1, pChainHeight - 1, pChainTimestamp - 1, 1, changes
                )
            )
        );
    }
}

// Test suite for testing the initialization of the first P-chain validator set before
// engaging in normal operation
contract AvalancheValidatorSetRegistryInitialization is AvalancheValidatorSetRegistryCommon {
    SubsetUpdater private _registry;
    DiffUpdater private _diffRegistry;

    function setUp() public {
        setUpSubsetUpdater();
        setUpDiffUpdater();
    }

    function setUpSubsetUpdater() public {
        (ValidatorSet memory validatorSet, bytes32 validatorSetHash) = dummyPChainValidatorSet();
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = validatorSetHash;
        ValidatorSetMetadata memory initialValidatorSetMetadata = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            shardHashes: shardHashes
        });
        _registry = new SubsetUpdater(NETWORK_ID, initialValidatorSetMetadata);
    }

    function setUpDiffUpdater() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSetForDiff();
        ValidatorSetDiff memory diff = allAdditionsValidatorSetDiff(validatorSet, 0, 0);
        bytes memory diffBytes = ValidatorSets.serializeValidatorSetDiff(diff);
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = sha256(diffBytes);
        require(
            validatorSet.pChainHeight > 0,
            "pChainHeight being transitioned to must be greater than 0 for diff setup"
        );
        ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight - 1,
            pChainTimestamp: validatorSet.pChainTimestamp - 1,
            shardHashes: shardHashes
        });
        _diffRegistry = new DiffUpdater(NETWORK_ID, metadata);
    }

    /**
     * @dev Test we successfully initialize the first P-chain validator set
     */
    function testPChainInitialization() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        (ValidatorSet memory validatorSetForDiff,) = dummyPChainValidatorSetForDiff();

        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorSet.validators),
            message: _emptyICMMessage()
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: _emptyICMMessage()
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testPChainInitialization(testCases[i]);
        }
    }

    /**
     * @dev Helper to test that attempting to continue to add the initial P-chain validator set is rejected
     * once it has been completed
     */
    function testPChainCantBeModifiedAfterInitialization() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        (ValidatorSet memory validatorSetForDiff,) = dummyPChainValidatorSetForDiff();

        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorSet.validators),
            message: _emptyICMMessage()
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: _emptyICMMessage()
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testPChainCantBeModifiedAfterInitialization(testCases[i]);
        }
    }

    /**
     * @dev Test that shards with the incorrect number are rejected
     */
    function testApplyShardOutOfOrder() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        (ValidatorSet memory validatorSetForDiff,) = dummyPChainValidatorSetForDiff();

        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorSet.validators),
            message: _emptyICMMessage()
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: _emptyICMMessage()
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testApplyShardOutOfOrder(testCases[i]);
        }
    }

    /**
     * @dev Test that the wrong validators bytes input causes a shard hash mismatch
     */
    function testApplyWrongShard() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        (ValidatorSet memory validatorSetForDiff,) = dummyPChainValidatorSetForDiff();

        Validator memory removed = validatorSet.validators[4];
        Validator[] memory subset = new Validator[](4);
        // copy over all validators but the last
        for (uint256 i = 0; i < 4; i++) {
            subset[i] = validatorSet.validators[i];
        }
        // This is a valid subset of the validators, but not what was committed to
        // via the shard hashes
        validatorSet.validators = subset;
        validatorSet.totalWeight -= removed.weight;
        validatorSetForDiff.validators = subset;
        validatorSetForDiff.totalWeight -= removed.weight;

        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorSet.validators),
            message: _emptyICMMessage()
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: _emptyICMMessage()
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testApplyWrongShard(testCases[i]);
        }
    }

    /**
     * @dev Check that if the initial P-chain validator set has not been fully initialized,
     * attempts to register other validator sets fails.
     */
    function testRegisterBeforeInitializationFails() public {
        // Setup
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        ValidatorSet memory validatorSetForDiff = ValidatorSet({
            avalancheBlockchainID: message.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(validatorShard[0].weight),
            pChainHeight: 1,
            pChainTimestamp: 1
        });
        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: message
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: messageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterBeforeInitializationFails(testCases[i]);
        }
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(_registry.getAvalancheNetworkID(), NETWORK_ID);
    }

    function _testPChainInitialization(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;

        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(registry.pChainInitialized());
        registry.updateValidatorSet(shard, shardBytes);
        assertTrue(registry.pChainInitialized());
    }

    function _testPChainCantBeModifiedAfterInitialization(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;

        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        ValidatorSetShard memory shard1 = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(registry.pChainInitialized());
        registry.updateValidatorSet(shard1, shardBytes);
        assertTrue(registry.pChainInitialized());
        ValidatorSetShard memory shard2 = ValidatorSetShard({
            shardNumber: 2,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        vm.expectRevert(bytes("Registration is not in progress"));
        registry.updateValidatorSet(shard2, shardBytes);
    }

    function _testApplyShardOutOfOrder(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;

        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 2,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(registry.pChainInitialized());
        vm.expectRevert(bytes("Received shard out of order"));
        registry.updateValidatorSet(shard, shardBytes);
    }

    function _testApplyWrongShard(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;

        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(registry.pChainInitialized());
        vm.expectRevert(bytes("Unexpected shard hash"));
        registry.updateValidatorSet(shard, shardBytes);
    }

    function _testRegisterBeforeInitializationFails(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;
        ICMMessage memory message = testCase.message;

        assertFalse(registry.pChainInitialized());
        vm.expectRevert(bytes("No P-chain validator set registered."));
        registry.registerValidatorSet(message, shardBytes);
    }
}

// Test suite for functionality after the initial P-chain set has been registered
contract AvalancheValidatorSetRegistryPostInitialization is AvalancheValidatorSetRegistryCommon {
    SubsetUpdater private _registry;
    DiffUpdater private _diffRegistry;

    function setUp() public {
        setUpSubsetUpdater();
        setUpDiffUpdater();
    }

    function setUpSubsetUpdater() public {
        (ValidatorSet memory validatorSet, bytes32 validatorSetHash) = dummyPChainValidatorSet();
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = validatorSetHash;
        ValidatorSetMetadata memory initialValidatorSetData = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            shardHashes: shardHashes
        });
        _registry = new SubsetUpdater(NETWORK_ID, initialValidatorSetData);
        // initialize the entire P-chain validator set
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorSet.validators);
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        _registry.updateValidatorSet(shard, validatorBytes);
    }

    function setUpDiffUpdater() public {
        (ValidatorSet memory validatorSetForDiff,) = dummyPChainValidatorSetForDiff();
        ValidatorSetDiff memory diff = allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0);
        bytes memory diffBytes = ValidatorSets.serializeValidatorSetDiff(diff);
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = sha256(diffBytes);
        ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSetForDiff.avalancheBlockchainID,
            pChainHeight: 0,
            pChainTimestamp: 0,
            shardHashes: shardHashes
        });
        _diffRegistry = new DiffUpdater(NETWORK_ID, metadata);
        // initialize the entire P-chain validator set
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSetForDiff.avalancheBlockchainID
        });
        _diffRegistry.updateValidatorSet(shard, diffBytes);
    }

    /**
     * @dev Test registering a new chain
     */
    function testRegisterNewChain() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();

        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        ValidatorSet memory validatorSetForDiff = ValidatorSet({
            avalancheBlockchainID: messageForDiff.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(validatorShard[0].weight),
            pChainHeight: 1,
            pChainTimestamp: 1
        });

        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: message
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(validatorSetForDiff, 0, 0)
            ),
            message: messageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterNewChain(testCases[i]);
        }
    }
    /**
     * @dev Test that we can register a new chain across two shards.
     */

    function testRegisterNewChainMultipleShards() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();

        Validator[] memory validatorShard = new Validator[](1);
        bytes[] memory subsetShards = new bytes[](2);
        bytes[] memory diffShards = new bytes[](2);

        // Compute subset updater shards
        {
            validatorShard[0] = validators[0];
            subsetShards[0] = ValidatorSets.serializeValidators(validatorShard);
            validatorShard[0] = validators[1];
            subsetShards[1] = ValidatorSets.serializeValidators(validatorShard);
        }

        // Compute diff shards
        {
            // Shard 1
            validatorShard[0] = validators[0];
            ValidatorSet memory vs1 = ValidatorSet({
                avalancheBlockchainID: messageForDiff.sourceBlockchainID,
                validators: validatorShard,
                totalWeight: uint64(validators[0].weight),
                pChainHeight: 1,
                pChainTimestamp: 1
            });
            diffShards[0] =
                ValidatorSets.serializeValidatorSetDiff(allAdditionsValidatorSetDiff(vs1, 0, 0));

            // Shard 2
            Validator[] memory cumulative = new Validator[](2);
            cumulative[0] = validators[0];
            cumulative[1] = validators[1];
            ValidatorChange[] memory shard2Changes = new ValidatorChange[](1);
            shard2Changes[0] = ValidatorChange({
                blsPublicKey: validators[1].blsPublicKey,
                weight: validators[1].weight
            });
            diffShards[1] = ValidatorSets.serializeValidatorSetDiff(
                customValidatorSetDiff(vs1, 0, 0, 1, shard2Changes)
            );
        }

        // Test
        ICMMessage[] memory subsetMessages = new ICMMessage[](1);
        subsetMessages[0] = message;
        ICMMessage[] memory diffMessages = new ICMMessage[](1);
        diffMessages[0] = messageForDiff;

        RegistryTestCaseMultiStep[] memory testCases = new RegistryTestCaseMultiStep[](2);
        testCases[0] = RegistryTestCaseMultiStep({
            registry: _registry,
            shardBytes: subsetShards,
            messages: subsetMessages
        });
        testCases[1] = RegistryTestCaseMultiStep({
            registry: _diffRegistry,
            shardBytes: diffShards,
            messages: diffMessages
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterNewChainMultipleShards(testCases[i]);
        }
    }

    /**
     * @dev Test that if we try register a validator set to a blockchain ID that is currently
     * awaiting updates from a previous registration, the tx reverts
     */
    function testInterruptingRegistrationFails() public {
        // same setup as `testRegisterNewChain`, so we skip the assertions done there
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        ValidatorSet memory vsForDiff = ValidatorSet({
            avalancheBlockchainID: messageForDiff.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(validators[0].weight),
            pChainHeight: 1,
            pChainTimestamp: 1
        });

        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: message
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(vsForDiff, 0, 0)
            ),
            message: messageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testInterruptingRegistrationFails(testCases[i]);
        }
    }

    /**
     * @dev Test that if we try to register a chain that has been registered before, if the message is
     * signed by the p-chain, it is rejected
     */
    function testRegisterChainWronglySignedByPChain() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();

        Validator[] memory validatorShard = new Validator[](1);
        bytes[] memory subsetShards = new bytes[](3);
        bytes[] memory diffShards = new bytes[](3);

        // Compute subset shards
        {
            validatorShard[0] = validators[0];
            subsetShards[0] = ValidatorSets.serializeValidators(validatorShard);
            validatorShard[0] = validators[1];
            subsetShards[1] = ValidatorSets.serializeValidators(validatorShard);
            validatorShard[0] = validators[0];
            subsetShards[2] = ValidatorSets.serializeValidators(validatorShard);
        }

        // Compute diff shards
        {
            validatorShard[0] = validators[0];
            ValidatorSet memory vs1 = ValidatorSet({
                avalancheBlockchainID: messageForDiff.sourceBlockchainID,
                validators: validatorShard,
                totalWeight: uint64(validators[0].weight),
                pChainHeight: 1,
                pChainTimestamp: 1
            });
            diffShards[0] =
                ValidatorSets.serializeValidatorSetDiff(allAdditionsValidatorSetDiff(vs1, 0, 0));

            Validator[] memory cumulative = new Validator[](2);
            cumulative[0] = validators[0];
            cumulative[1] = validators[1];
            ValidatorChange[] memory shard2Changes = new ValidatorChange[](1);
            shard2Changes[0] = ValidatorChange({
                blsPublicKey: validators[1].blsPublicKey,
                weight: validators[1].weight
            });
            diffShards[1] = ValidatorSets.serializeValidatorSetDiff(
                customValidatorSetDiff(vs1, 0, 0, 1, shard2Changes)
            );
            // Re-registration attempt uses same diff as shard 1
            diffShards[2] = diffShards[0];
        }

        ICMMessage[] memory subsetMessages = new ICMMessage[](1);
        subsetMessages[0] = message;
        ICMMessage[] memory diffMessages = new ICMMessage[](1);
        diffMessages[0] = messageForDiff;

        RegistryTestCaseMultiStep[] memory testCases = new RegistryTestCaseMultiStep[](2);
        testCases[0] = RegistryTestCaseMultiStep({
            registry: _registry,
            shardBytes: subsetShards,
            messages: subsetMessages
        });
        testCases[1] = RegistryTestCaseMultiStep({
            registry: _diffRegistry,
            shardBytes: diffShards,
            messages: diffMessages
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterChainWronglySignedByPChain(testCases[i]);
        }
    }

    /**
     * @dev Test that if we try to register a chain that has not been registered before, if the message is
     * not signed by the p-chain, it is rejected
     */
    function testRegisterChainWronglySignedByL1() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetAgainFixture(10, 10);
        (, ICMMessage memory messageForDiff) = registerValidatorSetAgainDiffFixture(10, 10);
        // Compute shards
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        ValidatorSet memory vsForDiff = ValidatorSet({
            avalancheBlockchainID: messageForDiff.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(validators[0].weight),
            pChainHeight: 10,
            pChainTimestamp: 10
        });
        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: message
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(vsForDiff, 9, 9)
            ),
            message: messageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterChainWronglySignedByL1(testCases[i]);
        }
    }

    /**
     * @dev Test the happy path when registering a validator set to a blockchain ID
     * which had a validator set registered to it prior
     */
    function testRegisterL1Again() public {
        // register an initial set
        _registerAndApplyInitialL1Set();

        // register a new set
        RegistryTestCaseMultiStep[] memory testCases = _buildReRegistrationTestCases(2, 2, 1, 1);
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterL1Again(testCases[i]);
        }
    }

    /**
     * @dev Same as`testRegisterL1Again` but the second registration should fail
     * because the P-chain height has not strictly increased
     */
    function testRegisterL1AgainBadPchainHeight() public {
        // Register
        _registerAndApplyInitialL1Set();
        // Re-register with bad height
        (Validator[] memory newValidators, ICMMessage memory newMessage) =
            registerValidatorSetAgainFixture(1, 2);
        (, ICMMessage memory newMessageForDiff) = registerValidatorSetAgainDiffFixture(1, 2);

        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = newValidators[0];

        ValidatorSet memory newValidatorSetForDiff = ValidatorSet({
            avalancheBlockchainID: newMessageForDiff.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(newValidators[0].weight),
            pChainHeight: 1,
            pChainTimestamp: 2
        });
        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: newMessage
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(newValidatorSetForDiff, 0, 1)
            ),
            message: newMessageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterL1AgainBadPchainHeight(testCases[i]);
        }
    }

    /**
     * @dev Same as`testRegisterL1AgainBadPChainHeight` but we test a bad P-chain timestamp instead
     */
    function testRegisterL1AgainBadPchainTimestamp() public {
        // Register
        _registerAndApplyInitialL1Set();

        // Re-register with bad timestamp
        (Validator[] memory newValidators, ICMMessage memory newMessage) =
            registerValidatorSetAgainFixture(2, 1);
        (, ICMMessage memory newMessageForDiff) = registerValidatorSetAgainDiffFixture(2, 1);

        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = newValidators[0];

        validatorShard[0] = newValidators[0];
        ValidatorSet memory newVsForDiff = ValidatorSet({
            avalancheBlockchainID: newMessageForDiff.sourceBlockchainID,
            validators: validatorShard,
            totalWeight: uint64(newValidators[0].weight),
            pChainHeight: 2,
            pChainTimestamp: 1
        });
        // Test
        RegistryTestCase[] memory testCases = new RegistryTestCase[](2);
        testCases[0] = RegistryTestCase({
            registry: _registry,
            shardBytes: ValidatorSets.serializeValidators(validatorShard),
            message: newMessage
        });
        testCases[1] = RegistryTestCase({
            registry: _diffRegistry,
            shardBytes: ValidatorSets.serializeValidatorSetDiff(
                allAdditionsValidatorSetDiff(newVsForDiff, 1, 0)
            ),
            message: newMessageForDiff
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterL1AgainBadPchainTimestamp(testCases[i]);
        }
    }

    function _testRegisterNewChain(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;
        ICMMessage memory message = testCase.message;
        // check that no validator set has ever been registered to this blockchain ID before
        assertFalse(registry.isRegistered(message.sourceBlockchainID));
        registry.registerValidatorSet(message, shardBytes);
        // check that a registration has started but is still in progress
        assertTrue(registry.isRegistrationInProgress(message.sourceBlockchainID));
        // check that no complete registration has occurred for this blockchain ID
        assertFalse(registry.isRegistered(message.sourceBlockchainID));
    }

    function _testRegisterNewChainMultipleShards(
        RegistryTestCaseMultiStep memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes[] memory shardBytes = testCase.shardBytes;
        ICMMessage memory message = testCase.messages[0];
        // register first shard
        registry.registerValidatorSet(message, shardBytes[0]);
        assertTrue(registry.isRegistrationInProgress(message.sourceBlockchainID));
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        // add the second shard
        registry.updateValidatorSet(shard, shardBytes[1]);
        // We should not have a fully registered validator set
        assertTrue(registry.isRegistered(message.sourceBlockchainID));
        // There should be no registrations in progress
        assertFalse(registry.isRegistrationInProgress(message.sourceBlockchainID));
    }

    function _testInterruptingRegistrationFails(
        RegistryTestCase memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        bytes memory shardBytes = testCase.shardBytes;
        ICMMessage memory message = testCase.message;
        registry.registerValidatorSet(message, shardBytes);

        // check that the interruption is caught and rejected
        vm.expectRevert(bytes("A registration is already in progress"));
        registry.registerValidatorSet(message, shardBytes);
    }

    function _testRegisterChainWronglySignedByPChain(
        RegistryTestCaseMultiStep memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        ICMMessage memory message = testCase.messages[0];
        // Register first shard
        registry.registerValidatorSet(message, testCase.shardBytes[0]);
        // Add second shard
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        registry.updateValidatorSet(shard, testCase.shardBytes[1]);
        // Try re-registering with same P-chain signature
        vm.expectRevert(bytes("Failed to verify signatures"));
        registry.registerValidatorSet(message, testCase.shardBytes[2]);
    }

    function _testRegisterChainWronglySignedByL1(
        RegistryTestCase memory testCase
    ) internal {
        vm.expectRevert(bytes("Failed to verify signatures"));
        testCase.registry.registerValidatorSet(testCase.message, testCase.shardBytes);
    }

    /**
     * @dev Registers a full 2 shard L1 validator set on both _registry and _diffRegistry.
     * After this call:
     * - Both registries have L1_BLOCKCHAIN_ID registered with 2 validators
     * - The P-Chain height and P-Chain timestamp are both 1
     * - No validator registration is in progress
     */
    function _registerAndApplyInitialL1Set() internal {
        // Setup
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        (, ICMMessage memory messageForDiff) = registerValidatorSetInitialDiffFixture();

        Validator[] memory validatorShard = new Validator[](1);
        bytes[] memory subsetShards = new bytes[](2);
        bytes[] memory diffShards = new bytes[](2);
        // Compute subset shards
        {
            validatorShard[0] = validators[0];
            subsetShards[0] = ValidatorSets.serializeValidators(validatorShard);
            validatorShard[0] = validators[1];
            subsetShards[1] = ValidatorSets.serializeValidators(validatorShard);
        }
        // Compute diff shards
        {
            validatorShard[0] = validators[0];
            ValidatorSet memory vs1 = ValidatorSet({
                avalancheBlockchainID: messageForDiff.sourceBlockchainID,
                validators: validatorShard,
                totalWeight: uint64(validators[0].weight),
                pChainHeight: 1,
                pChainTimestamp: 1
            });
            diffShards[0] =
                ValidatorSets.serializeValidatorSetDiff(allAdditionsValidatorSetDiff(vs1, 0, 0));

            Validator[] memory cumulative = new Validator[](2);
            cumulative[0] = validators[0];
            cumulative[1] = validators[1];
            ValidatorChange[] memory shard2Changes = new ValidatorChange[](1);
            shard2Changes[0] = ValidatorChange({
                blsPublicKey: validators[1].blsPublicKey,
                weight: validators[1].weight
            });
            diffShards[1] = ValidatorSets.serializeValidatorSetDiff(
                customValidatorSetDiff(vs1, 0, 0, 1, shard2Changes)
            );
        }
        ICMMessage[] memory subsetMessages = new ICMMessage[](1);
        subsetMessages[0] = message;
        ICMMessage[] memory diffMessages = new ICMMessage[](1);
        diffMessages[0] = messageForDiff;

        // Test
        RegistryTestCaseMultiStep[] memory testCases = new RegistryTestCaseMultiStep[](2);
        testCases[0] = RegistryTestCaseMultiStep({
            registry: _registry,
            shardBytes: subsetShards,
            messages: subsetMessages
        });
        testCases[1] = RegistryTestCaseMultiStep({
            registry: _diffRegistry,
            shardBytes: diffShards,
            messages: diffMessages
        });
        for (uint256 i = 0; i < testCases.length; i++) {
            _testRegisterNewChainMultipleShards(testCases[i]);
        }
    }

    function _testRegisterL1AgainBadPchainHeight(
        RegistryTestCase memory testCase
    ) internal {
        vm.expectRevert(bytes("P-Chain height too low"));
        testCase.registry.registerValidatorSet(testCase.message, testCase.shardBytes);
    }

    function _testRegisterL1AgainBadPchainTimestamp(
        RegistryTestCase memory testCase
    ) internal {
        vm.expectRevert(bytes("P-Chain timestamp too low"));
        testCase.registry.registerValidatorSet(testCase.message, testCase.shardBytes);
    }

    /**
     * @dev Builds RegistryTestCaseMultiStep for re-registering an L1 set on both registries.
     */
    function _buildReRegistrationTestCases(
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) internal returns (RegistryTestCaseMultiStep[] memory) {
        RegistryTestCaseMultiStep[] memory testCases = new RegistryTestCaseMultiStep[](2);
        testCases[0] = _buildSubsetReRegistrationTestCase(pChainHeight, pChainTimestamp);
        testCases[1] = _buildDiffReRegistrationTestCase(
            pChainHeight, pChainTimestamp, previousHeight, previousTimestamp
        );
        return testCases;
    }

    function _buildSubsetReRegistrationTestCase(
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) internal returns (RegistryTestCaseMultiStep memory) {
        (Validator[] memory newValidators, ICMMessage memory newMessage) =
            registerValidatorSetAgainFixture(pChainHeight, pChainTimestamp);
        Validator[] memory validatorShard = new Validator[](1);
        bytes[] memory subsetShards = new bytes[](2);
        validatorShard[0] = newValidators[0];
        subsetShards[0] = ValidatorSets.serializeValidators(validatorShard);
        validatorShard[0] = newValidators[1];
        subsetShards[1] = ValidatorSets.serializeValidators(validatorShard);
        ICMMessage[] memory messages = new ICMMessage[](1);
        messages[0] = newMessage;
        return RegistryTestCaseMultiStep({
            registry: _registry,
            shardBytes: subsetShards,
            messages: messages
        });
    }

    function _buildDiffReRegistrationTestCase(
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) internal returns (RegistryTestCaseMultiStep memory) {
        Validator[] memory existingValidators = new Validator[](2);
        existingValidators[0] = Validator({blsPublicKey: BLST.getPublicKeyFromSecret(2), weight: 2});
        existingValidators[1] = Validator({blsPublicKey: BLST.getPublicKeyFromSecret(3), weight: 3});
        (Validator[] memory newValidatorsForDiff, ICMMessage memory newMessageForDiff) =
            reRegisterValidatorSetDiffFixture(existingValidators, pChainHeight, pChainTimestamp);
        bytes[] memory diffShards = _buildReRegDiffShards({
            existingValidators: existingValidators,
            newValidatorsForDiff: newValidatorsForDiff,
            blockchainID: newMessageForDiff.sourceBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp
        });
        ICMMessage[] memory messages = new ICMMessage[](1);
        messages[0] = newMessageForDiff;
        return RegistryTestCaseMultiStep({
            registry: _diffRegistry,
            shardBytes: diffShards,
            messages: messages
        });
    }

    function _buildReRegDiffShards(
        Validator[] memory existingValidators,
        Validator[] memory newValidatorsForDiff,
        bytes32 blockchainID,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) internal returns (bytes[] memory) {
        bytes[] memory diffShards = new bytes[](2);
        diffShards[0] = _computeReRegDiffShard1({
            existingValidators: existingValidators,
            newValidator: newValidatorsForDiff[0],
            blockchainID: blockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp
        });
        diffShards[1] = _computeReRegDiffShard2({
            shard1Validator: newValidatorsForDiff[0],
            newValidator: newValidatorsForDiff[1],
            blockchainID: blockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp
        });
        return diffShards;
    }

    function _computeReRegDiffShard1(
        Validator[] memory existingValidators,
        Validator memory newValidator,
        bytes32 blockchainID,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) internal returns (bytes memory) {
        ValidatorChange[] memory changes = new ValidatorChange[](existingValidators.length + 1);
        for (uint256 i = 0; i < existingValidators.length; i++) {
            changes[i] =
                ValidatorChange({blsPublicKey: existingValidators[i].blsPublicKey, weight: 0});
        }
        changes[existingValidators.length] =
            ValidatorChange({blsPublicKey: newValidator.blsPublicKey, weight: newValidator.weight});
        ValidatorSet memory existingSet = ValidatorSet({
            avalancheBlockchainID: blockchainID,
            validators: existingValidators,
            totalWeight: 0,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
        return ValidatorSets.serializeValidatorSetDiff(
            customValidatorSetDiff(existingSet, previousHeight, previousTimestamp, 1, changes)
        );
    }

    function _computeReRegDiffShard2(
        Validator memory shard1Validator,
        Validator memory newValidator,
        bytes32 blockchainID,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        uint64 previousHeight,
        uint64 previousTimestamp
    ) internal returns (bytes memory) {
        Validator[] memory shard1Result = new Validator[](1);
        shard1Result[0] = shard1Validator;
        ValidatorSet memory afterShard1 = ValidatorSet({
            avalancheBlockchainID: blockchainID,
            validators: shard1Result,
            totalWeight: uint64(shard1Validator.weight),
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
        ValidatorChange[] memory changes = new ValidatorChange[](1);
        changes[0] =
            ValidatorChange({blsPublicKey: newValidator.blsPublicKey, weight: newValidator.weight});
        return ValidatorSets.serializeValidatorSetDiff(
            customValidatorSetDiff(afterShard1, previousHeight, previousTimestamp, 1, changes)
        );
    }

    function _testRegisterL1Again(
        RegistryTestCaseMultiStep memory testCase
    ) internal {
        AvalancheValidatorSetRegistry registry = testCase.registry;
        ICMMessage memory message = testCase.messages[0];

        // register the first shard
        registry.registerValidatorSet(message, testCase.shardBytes[0]);
        assertTrue(registry.isRegistered(message.sourceBlockchainID));

        // a set has been registered previously to this blockchain ID
        // a registration is in progress
        assertTrue(registry.isRegistrationInProgress(message.sourceBlockchainID));

        // register the second shard
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        registry.updateValidatorSet(shard, testCase.shardBytes[1]);
        // check that the registration is no longer in progress
        assertFalse(registry.isRegistrationInProgress(message.sourceBlockchainID));
    }
}

contract AvalancheValidatorSetRegistryTests is AvalancheValidatorSetRegistryCommon {
    DiffUpdater private _diffRegistry;

    /**
     *  @dev Creates a new instance of DiffUpdater pre-configured with the provided expected shard hashes.
     */
    function setUpDiffWithHashes(
        bytes32[] memory shardHashes
    ) public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            shardHashes: shardHashes
        });
        _diffRegistry = new DiffUpdater(NETWORK_ID, metadata);
    }

    /**
     * @dev Helper function for testUpdateValidatorSetWithDiffSuccess() that registers a validator set via a diff.
     */
    function helperTestRegisterValidatorSetWithDiffSuccess(
        bytes32 chainID,
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) public {
        ValidatorChange[] memory postChanges = new ValidatorChange[](1);
        postChanges[0] = ValidatorChange({blsPublicKey: BLST.getPublicKeyFromSecret(7), weight: 5});
        // Diff
        ValidatorSetDiff memory postDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: pChainHeight + 1,
            previousTimestamp: pChainTimestamp + 1,
            currentHeight: pChainHeight + 2,
            currentTimestamp: pChainTimestamp + 2,
            changes: postChanges,
            numAdded: 1,
            newSize: 6
        });
        bytes memory postDiffBytes = ValidatorSets.serializeValidatorSetDiff(postDiff);
        bytes32[] memory postShardHashes = new bytes32[](1);
        postShardHashes[0] = sha256(postDiffBytes);
        ValidatorSetMetadata memory nextMetadata = ValidatorSetMetadata({
            avalancheBlockchainID: chainID,
            pChainHeight: pChainHeight + 2,
            pChainTimestamp: pChainTimestamp + 2,
            shardHashes: postShardHashes
        });
        // Sign
        uint256[] memory signers = new uint256[](5);
        signers[0] = 1;
        signers[1] = 3;
        signers[2] = 4;
        signers[3] = 5;
        signers[4] = 6;
        bytes memory explicitBitmask = hex"1F";
        ICMMessage memory icmMsg =
            createCustomRegistrationMessage(chainID, nextMetadata, signers, explicitBitmask);
        // Register
        _diffRegistry.registerValidatorSet(icmMsg, postDiffBytes);
        assertFalse(_diffRegistry.isRegistrationInProgress(chainID));
    }

    /**
     * @dev Verifies workflow of initializing a validator set, then applying sharded diffs, and then registering a new validator set via a diff.
     */
    function testUpdateValidatorSetWithDiffSuccess() public {
        // Initialize
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        bytes32 chainID = validatorSet.avalancheBlockchainID;
        bytes32[] memory shardHashes = new bytes32[](2);

        // Compute first diff
        ValidatorChange[] memory initialChanges =
            new ValidatorChange[](validatorSet.validators.length);
        for (uint256 i = 0; i < validatorSet.validators.length; i++) {
            initialChanges[i] = ValidatorChange({
                blsPublicKey: validatorSet.validators[i].blsPublicKey,
                weight: validatorSet.validators[i].weight
            });
        }
        _sortValidatorChanges(initialChanges);
        ValidatorSetDiff memory initialDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            changes: initialChanges,
            numAdded: uint32(validatorSet.validators.length),
            newSize: validatorSet.validators.length
        });
        bytes memory initialDiffBytes = ValidatorSets.serializeValidatorSetDiff(initialDiff);
        shardHashes[0] = sha256(initialDiffBytes);

        // Compute second diff
        ValidatorChange[] memory changes = new ValidatorChange[](3);
        changes[0] =
            ValidatorChange({blsPublicKey: validatorSet.validators[0].blsPublicKey, weight: 0}); // Remove
        changes[1] = ValidatorChange({
            blsPublicKey: validatorSet.validators[1].blsPublicKey,
            weight: uint64(validatorSet.validators[1].weight + 10)
        }); // Modify
        changes[2] = ValidatorChange({blsPublicKey: BLST.getPublicKeyFromSecret(6), weight: 1}); // Add
        _sortValidatorChanges(changes);
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            changes: changes,
            numAdded: 1,
            newSize: validatorSet.validators.length
        });
        bytes memory diffBytes = ValidatorSets.serializeValidatorSetDiff(diff);
        shardHashes[1] = sha256(diffBytes);

        // Test applying diffs
        setUpDiffWithHashes(shardHashes);
        _diffRegistry.updateValidatorSet(
            ValidatorSetShard({shardNumber: 1, avalancheBlockchainID: chainID}), initialDiffBytes
        );
        _diffRegistry.updateValidatorSet(
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: chainID}), diffBytes
        );

        // Assert
        assertTrue(_diffRegistry.pChainInitialized());
        assertTrue(_diffRegistry.isRegistered(chainID));
        assertFalse(_diffRegistry.isRegistrationInProgress(chainID));

        // Test registering a new validator set
        helperTestRegisterValidatorSetWithDiffSuccess(
            chainID, validatorSet.pChainHeight, validatorSet.pChainTimestamp
        );
    }

    /**
     * @dev Verifies that invalid diff updates revert correctly and prevent concurrent registrations.
     */
    function testUpdateValidatorSetWithDiffFailure() public {
        // Initialize
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        bytes32 chainID = validatorSet.avalancheBlockchainID;
        bytes32[] memory shardHashes = new bytes32[](2);

        // Compute initial diff
        ValidatorChange[] memory initialChanges =
            new ValidatorChange[](validatorSet.validators.length);
        for (uint256 i = 0; i < validatorSet.validators.length; i++) {
            initialChanges[i] = ValidatorChange({
                blsPublicKey: validatorSet.validators[i].blsPublicKey,
                weight: validatorSet.validators[i].weight
            });
        }
        _sortValidatorChanges(initialChanges);
        ValidatorSetDiff memory initialDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            changes: initialChanges,
            numAdded: uint32(validatorSet.validators.length),
            newSize: validatorSet.validators.length
        });
        bytes memory initialDiffBytes = ValidatorSets.serializeValidatorSetDiff(initialDiff);
        shardHashes[0] = sha256(initialDiffBytes);

        // Compute invalid diff
        ValidatorChange[] memory changes = new ValidatorChange[](3);
        changes[0] =
            ValidatorChange({blsPublicKey: validatorSet.validators[0].blsPublicKey, weight: 0}); // Remove
        changes[1] = ValidatorChange({
            blsPublicKey: validatorSet.validators[1].blsPublicKey,
            weight: uint64(validatorSet.validators[1].weight + 10)
        }); // Modify
        changes[2] = ValidatorChange({blsPublicKey: BLST.getPublicKeyFromSecret(6), weight: 1}); // Add
        _sortValidatorChanges(changes);
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            currentHeight: validatorSet.pChainHeight, // Invalid height
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            changes: changes,
            numAdded: 1,
            newSize: validatorSet.validators.length
        });
        bytes memory diffBytes = ValidatorSets.serializeValidatorSetDiff(diff);
        shardHashes[1] = sha256(diffBytes);

        // Test
        setUpDiffWithHashes(shardHashes);
        _diffRegistry.updateValidatorSet(
            ValidatorSetShard({shardNumber: 1, avalancheBlockchainID: chainID}), initialDiffBytes
        );
        bytes32 wrongHash = sha256("wrongData");
        vm.expectRevert(bytes("Unexpected shard hash"));
        _diffRegistry.updateValidatorSet(
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: chainID}),
            abi.encodePacked(wrongHash)
        );
        vm.expectRevert(bytes("P-Chain height too low"));
        _diffRegistry.updateValidatorSet(
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: chainID}), diffBytes
        );
        ICMMessage memory dummyMessage = ICMMessage({
            rawMessage: new bytes(0),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: chainID,
            attestation: new bytes(0)
        });
        vm.expectRevert(bytes("A registration is already in progress"));
        _diffRegistry.registerValidatorSet(dummyMessage, new bytes(0));
    }

    /**
     * @notice Sorts validator changes by their uncompressed public key bytes
     * @dev Uses insertion sort which is efficient for small arrays
     */
    function _sortValidatorChanges(
        ValidatorChange[] memory changes
    ) private pure {
        for (uint256 i = 1; i < changes.length; i++) {
            ValidatorChange memory key = changes[i];
            bytes memory keyPubKey = key.blsPublicKey;
            int256 j = int256(i) - 1;
            while (j >= 0) {
                bytes memory jPubKey = changes[uint256(j)].blsPublicKey;
                if (
                    BLST.comparePublicKeys(
                        BLST.unPadUncompressedBlsPublicKey(jPubKey),
                        BLST.unPadUncompressedBlsPublicKey(keyPubKey)
                    ) <= 0
                ) {
                    break;
                }
                changes[uint256(j + 1)] = changes[uint256(j)];
                j--;
            }
            changes[uint256(j + 1)] = key;
        }
    }
}
