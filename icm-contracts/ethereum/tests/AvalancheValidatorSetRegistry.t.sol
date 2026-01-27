// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ICM, ICMMessage, ICMRawMessage} from "../../common/ICM.sol";
import {BLST} from "../utils/BLST.sol";
import {AvalancheValidatorSetRegistry} from "../AvalancheValidatorSetRegistry.sol";
import {
    Validator,
    ValidatorSet,
    ValidatorSets,
    ValidatorSetMetadata,
    ValidatorSetShard,
    ValidatorSetSignature
} from "../utils/ValidatorSets.sol";

// Common utility functions and fixtures for the suites in this file
contract AvalancheValidatorSetRegistryCommon is Test {
    uint32 public constant NETWORK_ID = 1;

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
            totalWeight += uint64(i + 1);
        }

        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(0),
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: 0,
            pChainTimestamp: 0
        });
        return (validatorSet, sha256(ValidatorSets.serializeValidators(validators)));
    }

    /**
     * @dev A fixture that returns a validator set along with an ICM message to register a
     * this set to a new chain ID. This set is split into two shards with one validator each.
     */
    function registerValidatorSetFixture()
        public
        view
        returns (Validator[] memory, ICMMessage memory)
    {
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
            totalWeight += uint64(i + 1);
            // compute the shard hash
            Validator[] memory validatorShard = new Validator[](1);
            validatorShard[0] = validators[i];
            shardHashes[i] = sha256(ValidatorSets.serializeValidators(validatorShard));
        }

        ValidatorSetMetadata memory metadata = ValidatorSetMetadata({
            avalancheBlockchainID: 0x3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7,
            pChainHeight: 10,
            pChainTimestamp: 10,
            validatorSetHash: sha256(ValidatorSets.serializeValidators(validators)),
            totalValidators: 2,
            shardHashes: shardHashes
        });

        // construct the ICM message
        ICMRawMessage memory raw = ICMRawMessage({
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: 0x3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7,
            sourceAddress: address(0),
            verifierAddress: address(0),
            payload: ValidatorSets.serializeValidatorSetMetadata(metadata)
        });
        bytes memory rawMessageBytes = ICM.serializeICMRawMessage(raw);
        // sign the message
        bytes memory signature = dummyPChainValidatorSetSign(rawMessageBytes);

        ICMMessage memory message =
            ICMMessage({message: raw, rawMessageBytes: rawMessageBytes, attestation: signature});

        return (validators, message);
    }

    /**
     * @dev Sign the input payload with the dummy validator set created above
     */
    function dummyPChainValidatorSetSign(
        bytes memory payload
    ) public view returns (bytes memory) {
        uint256[] memory secretKeys = dummyPChainValidatorSetSecretKeys();
        bytes memory rawSig = BLST.createAggregateSignature(secretKeys, payload);
        ValidatorSetSignature memory signature = ValidatorSetSignature({
            // all five validators sign
            signers: hex"F8",
            signature: rawSig
        });
        return ValidatorSets.serializeValidatorSetSignature(signature);
    }

    /**
     * @dev A fixture holding the secret keys of the P-chain validator set. These are sorted
     * so that their public counterparts are in increasing order
     */
    function dummyPChainValidatorSetSecretKeys() public pure returns (uint256[] memory) {
        uint256[] memory secretKeys = new uint256[](5);
        /* solhint-disable-next-line no-inline-assembly */
        assembly {
            mstore(secretKeys, 5)
            mstore(add(secretKeys, 0x20), 2)
            mstore(add(secretKeys, 0x40), 3)
            mstore(add(secretKeys, 0x60), 4)
            mstore(add(secretKeys, 0x80), 5)
            mstore(add(secretKeys, 0xA0), 1)
        }
        return secretKeys;
    }
}

// Test suite for testing the initialization of the first P-chain validator set before
// engaging in normal operation
contract AvalancheValidatorSetRegistryInitialization is AvalancheValidatorSetRegistryCommon {
    AvalancheValidatorSetRegistry private _registry;

    function setUp() public {
        (ValidatorSet memory validatorSet, bytes32 validatorSetHash) = dummyPChainValidatorSet();
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = validatorSetHash;
        ValidatorSetMetadata memory initialValidatorSetData = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            validatorSetHash: validatorSetHash,
            totalValidators: 5,
            shardHashes: shardHashes
        });
        _registry = new AvalancheValidatorSetRegistry(NETWORK_ID, initialValidatorSetData);
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
        vm.expectRevert(bytes("Cannot apply shard if registration is not in progress"));
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
        (Validator[] memory validators, ICMMessage memory message) = registerValidatorSetFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        vm.expectRevert(
            bytes("A complete P-chain validator must be registered to verify ICM messages")
        );
        _registry.registerValidatorSet(message, validatorBytes);
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(_registry.getAvalancheNetworkID(), NETWORK_ID);
    }
}

// Test suite for functionality after the initial P-chain set has been registered
contract AvalancheValidatorSetRegistryPostInitialization is AvalancheValidatorSetRegistryCommon {
    AvalancheValidatorSetRegistry private _registry;

    function setUp() public {
        (ValidatorSet memory validatorSet, bytes32 validatorSetHash) = dummyPChainValidatorSet();
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = validatorSetHash;
        ValidatorSetMetadata memory initialValidatorSetData = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            validatorSetHash: validatorSetHash,
            totalValidators: 5,
            shardHashes: shardHashes
        });
        _registry = new AvalancheValidatorSetRegistry(NETWORK_ID, initialValidatorSetData);
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
        (Validator[] memory validators, ICMMessage memory message) = registerValidatorSetFixture();
        Validator[] memory validatorShard = new Validator[](1);
        validatorShard[0] = validators[0];
        bytes memory validatorBytes = ValidatorSets.serializeValidators(validatorShard);
        _registry.registerValidatorSet(message, validatorBytes);
    }
}
