// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {BLST} from "../utils/BLST.sol";
import {
    Validator,
    ValidatorSet,
    ValidatorSets,
    ValidatorSetSignature,
    ValidatorSetMetadata,
    ValidatorSetShard
} from "../utils/ValidatorSets.sol";

contract ValidatorSetsTest is Test {
    function testFilterValidators() public view {
        // 0100_1000_0110_0000_1000_0000 in hex. This corresponds to validators
        // 2, 5, 10, 11, and 17
        bytes memory signers = hex"486080";
        (Validator[] memory validators,) = _createValidatorSet(20);
        (bytes memory aggregateKey, uint64 aggregateWeight) =
            ValidatorSets.filterValidators(signers, validators);
        bytes memory expectedKey = validators[1].blsPublicKey;
        expectedKey = BLST.addG1(validators[4].blsPublicKey, expectedKey);
        expectedKey = BLST.addG1(validators[9].blsPublicKey, expectedKey);
        expectedKey = BLST.addG1(validators[10].blsPublicKey, expectedKey);
        expectedKey = BLST.addG1(validators[16].blsPublicKey, expectedKey);
        uint64 expectedWeight = 2 + 5 + 10 + 11 + 17;
        assertEq(aggregateWeight, expectedWeight);
        assertEq(aggregateKey, expectedKey);
    }

    /*
     * @dev Test the happy flow when a quorum of a validator set signs a message
     */
    function testVerifyValidatorSetSigPositive(
        bytes memory message
    ) public view {
        (Validator[] memory validators, uint64 totalWeight) = _createValidatorSet(10);
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(0),
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: uint64(0),
            pChainTimestamp: uint64(0)
        });
        // 0000_0111_1100_0000 in hex. This corresponds to validators
        // 6, 7, 8, 9, and  10. This is a quorum (weight 40 out of 55)
        bytes memory signers = hex"07C0";
        uint256[] memory secretKeys = new uint256[](5);
        for (uint256 i = 6; i <= 10; i++) {
            secretKeys[i - 6] = i;
        }
        bytes memory sig = BLST.createAggregateSignature(secretKeys, message);
        ValidatorSetSignature memory signature =
            ValidatorSetSignature({signers: signers, signature: sig});
        bool res = ValidatorSets.verifyValidatorSetSignature(signature, message, validatorSet);
        assert(res);
    }

    /*
     * @dev Test that we reject if the backing signatures don't represent a quorum
     */
    function testVerifyValidatorSetSigNoQuorum(
        bytes memory message
    ) public view {
        (Validator[] memory validators, uint64 totalWeight) = _createValidatorSet(10);
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(0),
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: uint64(0),
            pChainTimestamp: uint64(0)
        });
        // 1111_1000_0000_0000 in hex. This corresponds to validators
        // 1, 2, 3, 4, and  5. This is not a quorum (weight 15 out of 55)
        bytes memory signers = hex"F800";
        uint256[] memory secretKeys = new uint256[](5);
        for (uint256 i = 1; i <= 5; i++) {
            secretKeys[i - 1] = i;
        }
        bytes memory sig = BLST.createAggregateSignature(secretKeys, message);
        ValidatorSetSignature memory signature =
            ValidatorSetSignature({signers: signers, signature: sig});
        bool res = ValidatorSets.verifyValidatorSetSignature(signature, message, validatorSet);
        assertFalse(res);
    }

    /*
     * @dev Test that we reject if the signers bitset doesn't match the actual signers
     */
    function testVerifyValidatorSetSigIncorrectSigners(
        bytes memory message
    ) public view {
        (Validator[] memory validators, uint64 totalWeight) = _createValidatorSet(10);
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: bytes32(0),
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: uint64(0),
            pChainTimestamp: uint64(0)
        });
        // 0000_0111_1100_0000 in hex. This corresponds to validators
        // 6, 7, 8, 9, and  10. This is a quorum (weight 40 out of 55)
        // But this is not the set that signs the message.
        bytes memory signers = hex"07C0";
        uint256[] memory secretKeys = new uint256[](5);
        for (uint256 i = 1; i <= 5; i++) {
            secretKeys[i - 1] = i;
        }
        bytes memory sig = BLST.createAggregateSignature(secretKeys, message);
        ValidatorSetSignature memory signature =
            ValidatorSetSignature({signers: signers, signature: sig});
        bool res = ValidatorSets.verifyValidatorSetSignature(signature, message, validatorSet);
        assert(!res);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSet(
        uint256 numValidators
    ) public pure {
        vm.assume(numValidators < 10);
        uint64 totalWeight = 0;
        Validator[] memory validators = new Validator[](numValidators);
        for (uint256 i = 1; i <= numValidators; i++) {
            validators[i - 1] = Validator({
                // these are not necessarily curve points. It just makes it easy to know the key ordering
                blsPublicKey: BLST.padUncompressedBLSPublicKey(_createPublicKeyFromWords(i, 0, i)),
                weight: uint64(i)
            });
            totalWeight += uint64(i);
        }
        bytes memory serialized = ValidatorSets.serializeValidators(validators);
        (Validator[] memory deserialized, uint64 deserializedTotalWeight) =
            ValidatorSets.parseValidators(serialized);
        assertEq(totalWeight, deserializedTotalWeight);
        assertEq(deserialized.length, validators.length);
        for (uint256 i = 0; i < numValidators; i++) {
            assertEq(validators[i].weight, deserialized[i].weight);
            assertEq(validators[i].blsPublicKey, deserialized[i].blsPublicKey);
        }
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSetPayload(
        bytes32 avalancheBlockchainID,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        bytes32[] memory shardHashes
    ) public pure {
        ValidatorSetMetadata memory payload = ValidatorSetMetadata({
            avalancheBlockchainID: avalancheBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            shardHashes: shardHashes
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetMetadata(payload);
        ValidatorSetMetadata memory deserialized =
            ValidatorSets.parseValidatorSetMetadata(serialized);
        assertEq(payload.avalancheBlockchainID, deserialized.avalancheBlockchainID);
        assertEq(payload.pChainHeight, deserialized.pChainHeight);
        assertEq(payload.pChainTimestamp, deserialized.pChainTimestamp);
        assertEq(payload.shardHashes, deserialized.shardHashes);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSetSignature(
        bytes memory signers,
        bytes32[6] memory signature
    ) public pure {
        bytes memory sig = abi.encodePacked(
            signature[0], signature[1], signature[2], signature[3], signature[4], signature[5]
        );
        ValidatorSetSignature memory validatorSetSig =
            ValidatorSetSignature({signers: signers, signature: sig});
        bytes memory serialized = ValidatorSets.serializeValidatorSetSignature(validatorSetSig);
        ValidatorSetSignature memory deserialized =
            ValidatorSets.parseValidatorSetSignature(serialized);
        assertEq(deserialized.signers, signers);
        assertEq(deserialized.signature, sig);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSetShard(
        uint64 shardNumber,
        bytes32 avalancheBlockchainID
    ) public pure {
        ValidatorSetShard memory validatorSetShard = ValidatorSetShard({
            shardNumber: shardNumber,
            avalancheBlockchainID: avalancheBlockchainID
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetShard(validatorSetShard);
        ValidatorSetShard memory deserialized = ValidatorSets.parseValidatorSetShard(serialized);

        assertEq(deserialized.shardNumber, shardNumber);
        assertEq(deserialized.avalancheBlockchainID, avalancheBlockchainID);
    }

    /*
     * @dev Test util to generate a set of validators. Returns validators and total staking weight
     * N.B. These validators are not sorted by key, so any test requiring that should not use this
     * function
     */
    function _createValidatorSet(
        uint256 numValidators
    ) private view returns (Validator[] memory, uint64) {
        Validator[] memory validators = new Validator[](numValidators);
        uint64 totalWeight = 0;
        for (uint256 i = 1; i <= numValidators; i++) {
            validators[i - 1] =
                Validator({blsPublicKey: BLST.getPublicKeyFromSecret(i), weight: uint64(i)});
            totalWeight += uint64(i);
        }
        return (validators, totalWeight);
    }

    /*
     * @dev Create 96 bytes from three 32 bytes words
     * N.B. These are for testing serialization and key-ordering. They are not real public keys
     * and should not be used as such.
     */
    function _createPublicKeyFromWords(
        uint256 x,
        uint256 y,
        uint256 z
    ) private pure returns (bytes memory) {
        bytes memory pk = new bytes(96);
        /* solhint-disable-next-line no-inline-assembly */
        assembly {
            mstore(add(pk, 0x20), x)
            mstore(add(pk, 0x40), y)
            mstore(add(pk, 0x60), z)
        }
        return pk;
    }
}
