// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {BLST} from "../utils/BLST.sol";
import {
    Validator,
    ValidatorChange,
    ValidatorSet,
    ValidatorSets,
    ValidatorSetDiff,
    ValidatorSetSignature,
    ValidatorSetMetadata,
    ValidatorSetShard
} from "../utils/ValidatorSets.sol";

contract ValidatorSetsTestHarness {
    function parseValidatorSetDiff(
        bytes calldata data,
        uint256 currentValidatorCount
    ) public pure returns (ValidatorSetDiff memory) {
        return ValidatorSets.parseValidatorSetDiff(data, currentValidatorCount);
    }

    function parseValidatorChange(
        bytes calldata data,
        uint256 offset
    ) public pure returns (ValidatorChange memory, uint256) {
        return ValidatorSets.parseValidatorChange(data, offset);
    }

    function parseValidators(
        bytes calldata data
    ) public pure returns (Validator[] memory, uint64) {
        return ValidatorSets.parseValidators(data);
    }

    function parseValidatorSetMetadata(
        bytes calldata data
    ) public pure returns (ValidatorSetMetadata memory) {
        return ValidatorSets.parseValidatorSetMetadata(data);
    }

    function parseValidatorSetSignature(
        bytes calldata signatureBytes
    ) public pure returns (ValidatorSetSignature memory) {
        return ValidatorSets.parseValidatorSetSignature(signatureBytes);
    }

    function parseValidatorSetShard(
        bytes calldata shardBytes
    ) public pure returns (ValidatorSetShard memory) {
        return ValidatorSets.parseValidatorSetShard(shardBytes);
    }
}

contract ValidatorSetsTest is Test {
    ValidatorSetsTestHarness private _harness;

    function setUp() public {
        _harness = new ValidatorSetsTestHarness();
    }

    function testFilterValidators() public view {
        // big.Int with bits 1,4,9,10,16 set = 0x010612. This corresponds to
        // validators at indices 1,4,9,10,16 (0-indexed) with weights 2,5,10,11,17
        bytes memory signers = hex"010612";
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
        // big.Int with bits 5,6,7,8,9 set = 0x03E0. This corresponds to
        // validators at indices 5-9 (0-indexed) with weights 6-10 (quorum: 40/55)
        bytes memory signers = hex"03E0";
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
        // big.Int with bits 0,1,2,3,4 set = 0x1F. This corresponds to
        // validators at indices 0-4 (0-indexed) with weights 1-5 (not a quorum: 15/55)
        bytes memory signers = hex"1F";
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
        // big.Int with bits 5,6,7,8,9 set = 0x03E0. This corresponds to
        // validators at indices 5-9 (0-indexed) with weights 6-10 (quorum: 40/55).
        // But this is not the set that signs the message.
        bytes memory signers = hex"03E0";
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
    function testRoundTripValidatorSetDiff(
        bytes32 avalancheBlockchainID,
        uint64 previousHeight,
        uint64 previousTimestamp,
        uint64 currentHeight,
        uint64 currentTimestamp,
        uint256 numChanges
    ) public view {
        vm.assume(numChanges < 10);
        ValidatorChange[] memory changes = new ValidatorChange[](numChanges);
        uint32 numAdded = 0;
        uint32 numRemoved = 0;
        for (uint256 i; i < numChanges; i++) {
            // we start from index 1 to ensure there are at least as many additions as removals
            (bool added, ValidatorChange memory change) = _createValidatorChange(i + 1);
            changes[i] = change;
            if (added) {
                ++numAdded;
            } else {
                ++numRemoved;
            }
        }
        ValidatorSetDiff memory valsetDiff = ValidatorSetDiff({
            avalancheBlockchainID: avalancheBlockchainID,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp,
            currentHeight: currentHeight,
            currentTimestamp: currentTimestamp,
            changes: changes,
            numAdded: numAdded,
            newSize: numAdded - numRemoved
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetDiff(valsetDiff);
        ValidatorSetDiff memory deserialized = _harness.parseValidatorSetDiff(serialized, 0);
        assertEq(valsetDiff.avalancheBlockchainID, deserialized.avalancheBlockchainID);
        assertEq(valsetDiff.previousHeight, deserialized.previousHeight);
        assertEq(valsetDiff.previousTimestamp, deserialized.previousTimestamp);
        assertEq(valsetDiff.currentHeight, deserialized.currentHeight);
        assertEq(valsetDiff.currentTimestamp, deserialized.currentTimestamp);
        for (uint256 i; i < numChanges; i++) {
            assertEq(valsetDiff.changes[i].blsPublicKey, deserialized.changes[i].blsPublicKey);
            assertEq(valsetDiff.changes[i].weight, deserialized.changes[i].weight);
        }
        assertEq(valsetDiff.numAdded, deserialized.numAdded);
        assertEq(valsetDiff.newSize, deserialized.newSize);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorChange(uint64 weight, uint256 secretKey) public view {
        ValidatorChange memory valChange =
            ValidatorChange({blsPublicKey: BLST.getPublicKeyFromSecret(secretKey), weight: weight});
        bytes memory serialized = ValidatorSets.serializeValidatorChange(valChange);
        /* solhint-disable-next-line no-unused-vars */
        (ValidatorChange memory deserialized, uint256 offset) =
            _harness.parseValidatorChange(serialized, 0);
        assertEq(valChange.blsPublicKey, deserialized.blsPublicKey);
        assertEq(valChange.weight, deserialized.weight);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSet(
        uint256 numValidators
    ) public view {
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
            _harness.parseValidators(serialized);
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
    ) public view {
        ValidatorSetMetadata memory payload = ValidatorSetMetadata({
            avalancheBlockchainID: avalancheBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            shardHashes: shardHashes
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetMetadata(payload);
        ValidatorSetMetadata memory deserialized = _harness.parseValidatorSetMetadata(serialized);
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
    ) public view {
        bytes memory sig = abi.encodePacked(
            signature[0], signature[1], signature[2], signature[3], signature[4], signature[5]
        );
        ValidatorSetSignature memory validatorSetSig =
            ValidatorSetSignature({signers: signers, signature: sig});
        bytes memory serialized = ValidatorSets.serializeValidatorSetSignature(validatorSetSig);
        ValidatorSetSignature memory deserialized = _harness.parseValidatorSetSignature(serialized);
        assertEq(deserialized.signers, signers);
        assertEq(deserialized.signature, sig);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripValidatorSetShard(
        uint64 shardNumber,
        bytes32 avalancheBlockchainID
    ) public view {
        ValidatorSetShard memory validatorSetShard = ValidatorSetShard({
            shardNumber: shardNumber,
            avalancheBlockchainID: avalancheBlockchainID
        });
        bytes memory serialized = ValidatorSets.serializeValidatorSetShard(validatorSetShard);
        ValidatorSetShard memory deserialized = _harness.parseValidatorSetShard(serialized);

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
     * @dev Create a validator change from an index
     */
    function _createValidatorChange(
        uint256 i
    ) private view returns (bool, ValidatorChange memory) {
        return (
            i % 2 != 0,
            ValidatorChange({blsPublicKey: BLST.getPublicKeyFromSecret(i), weight: uint64(i % 2)})
        );
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
