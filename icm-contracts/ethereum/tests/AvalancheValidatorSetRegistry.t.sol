// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
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

// TODO: Address issue to consolidate tests across implementations of AvalanceValidatorSetRegistry: https://github.com/ava-labs/icm-services/issues/1201

// Common utility functions and fixtures for the suites in this file
contract AvalancheValidatorSetRegistryCommon is Test {
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

    /**
     * @dev Sign the input payload with the dummy validator set created above
     */
    function dummyPChainValidatorSetSign(
        bytes32 avalancheBlockchainID,
        bytes memory payload
    ) public view returns (bytes memory) {
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
        bytes memory signedData = abi.encodePacked(NETWORK_ID, avalancheBlockchainID, payload);
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, signedData);
        ValidatorSetSignature memory signature = ValidatorSetSignature({
            // all five validators sign
            signers: hex"F8",
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
        bytes memory signedData = abi.encodePacked(NETWORK_ID, L1_BLOCKCHAIN_ID, payload);
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, signedData);
        ValidatorSetSignature memory signature = ValidatorSetSignature({
            //  both validators sign
            signers: hex"C0",
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
        bytes memory signedData = abi.encodePacked(NETWORK_ID, chainID, payload);
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
}

// Test suite for testing the initialization of the first P-chain validator set before
// engaging in normal operation
contract AvalancheValidatorSetRegistryInitialization is AvalancheValidatorSetRegistryCommon {
    SubsetUpdater private _registry;

    function setUp() public {
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

    /**
     * @dev Test we successfully initialize the first P-chain validator set
     */
    function testPChainInitialization() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorSet.validators);
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(_registry.pChainInitialized());
        _registry.updateValidatorSet(shard, validatorBytes);
        assertTrue(_registry.pChainInitialized());
    }

    /**
     * @dev Test that attempting to continue to add the initial P-chain validator set is rejected
     * once it has been completed
     */
    function testPChainCantBeModifiedAfterInitialization() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorSet.validators);
        ValidatorSetShard memory shard1 = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(_registry.pChainInitialized());
        _registry.updateValidatorSet(shard1, validatorBytes);
        assertTrue(_registry.pChainInitialized());
        ValidatorSetShard memory shard2 = ValidatorSetShard({
            shardNumber: 2,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        vm.expectRevert(bytes("Registration is not in progress"));
        _registry.updateValidatorSet(shard2, validatorBytes);
    }

    /**
     * @dev Test that shards with the incorrect number are rejected
     */
    function testApplyShardOutOfOrder() public {
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorSet.validators);
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 2,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(_registry.pChainInitialized());
        vm.expectRevert(bytes("Received shard out of order"));
        _registry.updateValidatorSet(shard, validatorBytes);
    }

    /**
     * @dev Test that the wrong validators bytes input causes a shard hash mismatch
     */
    function testApplyWrongShard() public {
        // This is a valid subset of the validators, but not what was committed to
        // via the shard hashes
        (ValidatorSet memory validatorSet,) = dummyPChainValidatorSet();
        Validator memory removed = validatorSet.validators[4];
        Validator[] memory subset = new Validator[](4);
        // copy over all validators but the last
        for (uint256 i = 0; i < 4; i++) {
            subset[i] = validatorSet.validators[i];
        }
        validatorSet.validators = subset;
        validatorSet.totalWeight -= removed.weight;

        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorSet.validators);
        ValidatorSetShard memory shard = ValidatorSetShard({
            shardNumber: 1,
            avalancheBlockchainID: validatorSet.avalancheBlockchainID
        });
        assertFalse(_registry.pChainInitialized());
        vm.expectRevert(bytes("Unexpected shard hash"));
        _registry.updateValidatorSet(shard, validatorBytes);
    }

    /**
     * @dev Check that if the initial P-chain validator set has not been fully initialized,
     * attempts to register other validator sets fails.
     */
    function testRegisterBeforeInitializationFails() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        vm.expectRevert(bytes("No P-chain validator set registered."));
        _registry.registerValidatorSet(message, validatorBytes);
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(_registry.getAvalancheNetworkID(), NETWORK_ID);
    }
}

// Test suite for functionality after the initial P-chain set has been registered
contract AvalancheValidatorSetRegistryPostInitialization is AvalancheValidatorSetRegistryCommon {
    SubsetUpdater private _registry;

    function setUp() public {
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

    /**
     * @dev Test registering a new chain
     */
    function testRegisterNewChain() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        // check that no validator set has ever been registered to this blockchain ID before
        assertFalse(_registry.isRegistered(message.sourceBlockchainID));
        _registry.registerValidatorSet(message, validatorBytes);
        // check that a registration has started but is still in progress
        assertTrue(_registry.isRegistrationInProgress(message.sourceBlockchainID));
        // check that no complete registration has occurred for this blockchain ID
        assertFalse(_registry.isRegistered(message.sourceBlockchainID));
    }

    /**
     * @dev Test that we can register a new chain across two shards.
     */
    function testRegisterNewChainMultipleShards() public {
        // same setup as above test, so we skip the assertions done there
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // add the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);
        // We should not have a fully registered validator set
        assertTrue(_registry.isRegistered(message.sourceBlockchainID));
        // There should be no registrations in progress
        assertFalse(_registry.isRegistrationInProgress(message.sourceBlockchainID));
    }

    /**
     * @dev Test that if we try register a validator set to a blockchain ID that is currently
     * awaiting updates from a previous registration, the tx reverts
     */
    function testInterruptingRegistrationFails() public {
        // same setup as `testRegisterNewChain`, so we skip the assertions done there
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // check that the interruption is caught and rejected
        vm.expectRevert(bytes("A registration is already in progress"));
        _registry.registerValidatorSet(message, validatorBytes);
    }

    /**
     * @dev Test that if we try to register a chain that has been registered before, if the message is
     * signed by the p-chain, it is rejected
     */
    function testRegisterChainWronglySignedByPChain() public {
        // register a full validator set
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // add the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);

        // Try to register this set again, still signed by the P-chain
        validatorShard[0] = validators[0];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        vm.expectRevert(bytes("Failed to verify signatures"));
        _registry.registerValidatorSet(message, validatorBytes);
    }

    /**
     * @dev Test that if we try to register a chain that has not been registered before, if the message is
     * not signed by the p-chain, it is rejected
     */
    function testRegisterChainWronglySignedByL1() public {
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetAgainFixture(10, 10);
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        vm.expectRevert(bytes("Failed to verify signatures"));
        _registry.registerValidatorSet(message, validatorBytes);
    }

    /**
     * @dev Test the happy path when registering a validator set to a blockchain ID
     * which had a validator set registered to it prior
     */
    function testRegisterL1Again() public {
        // register an initial set
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // add the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);

        // register a new set
        (validators, message) = registerValidatorSetAgainFixture(10, 10);
        validatorShard[0] = validators[0];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);

        // register the first shard
        _registry.registerValidatorSet(message, validatorBytes);
        assertTrue(_registry.isRegistered(message.sourceBlockchainID));

        // a set has been registered previously to this blockchain ID
        // a registration is in progress
        assertTrue(_registry.isRegistrationInProgress(message.sourceBlockchainID));

        // register the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);
        // check that the registration is no longer in progress
        assertFalse(_registry.isRegistrationInProgress(message.sourceBlockchainID));
    }

    /**
     * @dev Same as`testRegisterL1Again` but the second registration should fail
     * because the P-chain height has not strictly increased
     */
    function testRegisterL1AgainBadPchainHeight() public {
        // register an initial set
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // add the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);

        // register a new set
        (validators, message) = registerValidatorSetAgainFixture(1, 2);
        validatorShard[0] = validators[0];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        // registering the first shard should fail
        vm.expectRevert(bytes("P-Chain height too low"));
        _registry.registerValidatorSet(message, validatorBytes);
    }

    /**
     * @dev Same as`testRegisterL1AgainBadPChainHeight` but we test a bad P-chain timestamp instead
     */
    function testRegisterL1AgainBadPchainTimestamp() public {
        // register an initial set
        (Validator[] memory validators, ICMMessage memory message) =
            registerValidatorSetInitialFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);

        // add the second shard
        validatorShard[0] = validators[1];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        ValidatorSetShard memory shard =
            ValidatorSetShard({shardNumber: 2, avalancheBlockchainID: message.sourceBlockchainID});
        _registry.updateValidatorSet(shard, validatorBytes);

        // register a new set
        (validators, message) = registerValidatorSetAgainFixture(2, 1);
        validatorShard[0] = validators[0];
        validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        // registering the first shard should fail
        vm.expectRevert(bytes("P-Chain timestamp too low"));
        _registry.registerValidatorSet(message, validatorBytes);
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
        uint64 pChainTimestamp,
        Validator[] memory validators
    ) public {
        ValidatorChange[] memory postChanges = new ValidatorChange[](1);
        postChanges[0] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: BLST.getPublicKeyFromSecret(7),
            weight: 5
        });
        // Diff
        ValidatorSetDiff memory postDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: pChainHeight + 1,
            previousTimestamp: pChainTimestamp + 1,
            previousValidatorSetHash: sha256(ValidatorSets.serializeValidators(validators)),
            currentHeight: pChainHeight + 2,
            currentTimestamp: pChainTimestamp + 2,
            currentValidatorSetHash: bytes32(0),
            changes: postChanges,
            numAdded: 1,
            newSize: 6
        });
        (Validator[] memory validatorsAfterDiff,) =
            ValidatorSets.applyValidatorSetDiff(validators, postDiff);
        postDiff.currentValidatorSetHash =
            sha256(ValidatorSets.serializeValidators(validatorsAfterDiff));
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
        bytes memory explicitBitmask = hex"F8";
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
                nodeID: bytes20(0),
                blsPublicKey: validatorSet.validators[i].blsPublicKey,
                weight: validatorSet.validators[i].weight
            });
        }
        _sortValidatorChanges(initialChanges);
        ValidatorSetDiff memory initialDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            previousValidatorSetHash: sha256(ValidatorSets.serializeValidators(new Validator[](0))),
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            currentValidatorSetHash: sha256(ValidatorSets.serializeValidators(validatorSet.validators)),
            changes: initialChanges,
            numAdded: uint32(validatorSet.validators.length),
            newSize: validatorSet.validators.length
        });
        bytes memory initialDiffBytes = ValidatorSets.serializeValidatorSetDiff(initialDiff);
        shardHashes[0] = sha256(initialDiffBytes);

        // Compute second diff
        ValidatorChange[] memory changes = new ValidatorChange[](3);
        changes[0] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: validatorSet.validators[0].blsPublicKey,
            weight: 0
        }); // Remove
        changes[1] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: validatorSet.validators[1].blsPublicKey,
            weight: uint64(validatorSet.validators[1].weight + 10)
        }); // Modify
        changes[2] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: BLST.getPublicKeyFromSecret(6),
            weight: 1
        }); // Add
        _sortValidatorChanges(changes);
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            previousValidatorSetHash: sha256(ValidatorSets.serializeValidators(validatorSet.validators)),
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            currentValidatorSetHash: bytes32(0),
            changes: changes,
            numAdded: 1,
            newSize: validatorSet.validators.length
        });
        (Validator[] memory validatorsAfterDiff,) =
            ValidatorSets.applyValidatorSetDiff(validatorSet.validators, diff);
        diff.currentValidatorSetHash =
            sha256(ValidatorSets.serializeValidators(validatorsAfterDiff));
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
            chainID, validatorSet.pChainHeight, validatorSet.pChainTimestamp, validatorsAfterDiff
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
                nodeID: bytes20(0),
                blsPublicKey: validatorSet.validators[i].blsPublicKey,
                weight: validatorSet.validators[i].weight
            });
        }
        _sortValidatorChanges(initialChanges);
        ValidatorSetDiff memory initialDiff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            previousValidatorSetHash: sha256(ValidatorSets.serializeValidators(new Validator[](0))),
            currentHeight: validatorSet.pChainHeight + 1,
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            currentValidatorSetHash: sha256(ValidatorSets.serializeValidators(validatorSet.validators)),
            changes: initialChanges,
            numAdded: uint32(validatorSet.validators.length),
            newSize: validatorSet.validators.length
        });
        bytes memory initialDiffBytes = ValidatorSets.serializeValidatorSetDiff(initialDiff);
        shardHashes[0] = sha256(initialDiffBytes);

        // Compute invalid diff
        ValidatorChange[] memory changes = new ValidatorChange[](3);
        changes[0] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: validatorSet.validators[0].blsPublicKey,
            weight: 0
        }); // Remove
        changes[1] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: validatorSet.validators[1].blsPublicKey,
            weight: uint64(validatorSet.validators[1].weight + 10)
        }); // Modify
        changes[2] = ValidatorChange({
            nodeID: bytes20(0),
            blsPublicKey: BLST.getPublicKeyFromSecret(6),
            weight: 1
        }); // Add
        _sortValidatorChanges(changes);
        ValidatorSetDiff memory diff = ValidatorSetDiff({
            avalancheBlockchainID: chainID,
            previousHeight: validatorSet.pChainHeight,
            previousTimestamp: validatorSet.pChainTimestamp,
            previousValidatorSetHash: sha256(ValidatorSets.serializeValidators(validatorSet.validators)),
            currentHeight: validatorSet.pChainHeight, // Invalid height
            currentTimestamp: validatorSet.pChainTimestamp + 1,
            currentValidatorSetHash: bytes32(0),
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
        vm.expectRevert(bytes("Invalid blockchain height"));
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
