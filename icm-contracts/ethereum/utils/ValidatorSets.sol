// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ByteComparator} from "./ByteComparator.sol";
import {ByteSlicer} from "./ByteSlicer.sol";
import {BLST} from "./BLST.sol";

struct Validator {
    bytes blsPublicKey;
    uint64 weight;
}

struct ValidatorSet {
    bytes32 avalancheBlockchainID;
    Validator[] validators;
    uint64 totalWeight;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
}

// The payload that initiates registering / updating a
// Validator set for the chain with ID `avalancheBlockchainID`.
struct ValidatorSetMetadata {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    // A list of hashes for each shard. These are expected to arrive
    // in order
    bytes32[] shardHashes;
}

// A shard to add to a validator set for which partial data has already
// been received. The `avalancheBlockchainID` field will be used to lookup the
// existing partial data
struct ValidatorSetShard {
    // Which shard this is. These are expected to be processed in order
    uint64 shardNumber;
    // The avalanche blockchain ID whose validator set we are updating
    bytes32 avalancheBlockchainID;
}

// A partial Validator set which can be constructed from
//  shards sent across multiple transactions
struct PartialValidatorSet {
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    // A list of hashes for each shard. These are expected to arrive
    // in order
    bytes32[] shardHashes;
    // The amount of shard already received
    uint64 shardsReceived;
    // The validators received so far
    Validator[] validators;
    // The weight of the above validators
    uint64 partialWeight;
    // A flag to indicate that this is now a complete validator set
    bool inProgress;
}

struct ValidatorSetSignature {
    bytes signers;
    bytes signature;
}

// ValidatorChange represents a single validator addition, removal, or modification
struct ValidatorChange {
    bytes20 nodeID; // 20 bytes
    bytes blsPublicKey; // 96 bytes uncompressed
    uint64 previousWeight; // Weight at previous height (0 for additions)
    uint64 currentWeight; // Weight at current height (0 for removals)
}

// ValidatorSetDiffPayload contains the diff information from a signed message
struct ValidatorSetDiffPayload {
    bytes32 avalancheBlockchainID;
    // Previous state
    uint64 previousHeight;
    uint64 previousTimestamp;
    bytes32 previousValidatorSetHash;
    // Current state
    uint64 currentHeight;
    uint64 currentTimestamp;
    bytes32 currentValidatorSetHash;
    // Changes
    ValidatorChange[] added;
    ValidatorChange[] removed;
    ValidatorChange[] modified;
}

library ValidatorSets {
    /* solhint-disable no-inline-assembly */
    uint256 private constant QUORUM_NUM = 67;
    uint256 private constant QUORUM_DEN = 100;

    /*
     * @dev The number of validators is expected to fit in 4 bytes
     */
    uint256 private constant PAYLOAD_TYPE_ID_LENGTH = 4;

    /*
     * @dev The number of validators is expected to fit in 4 bytes
     */
    uint256 private constant NUM_VALIDATOR_LENGTH = 4;
    /*
     * @dev The weight of a validator is expected to fit in 8 bytes
     */
    uint256 private constant VALIDATOR_WEIGHT_LENGTH = 8;

    /*
     * @dev The number of bytes in the `Validator` struct
     */
    uint256 private constant VALIDATOR_BYTES =
        BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH + VALIDATOR_WEIGHT_LENGTH;

    /*
     * @notice Verifies that a quorum of the input validator set produced the input signature over the input message
     */
    function verifyValidatorSetSignature(
        ValidatorSetSignature calldata signature,
        bytes calldata message,
        ValidatorSet calldata validatorSet
    ) public view returns (bool) {
        (bytes memory aggregateKey, uint64 aggregateWeight) =
            filterValidators(signature.signers, validatorSet.validators);
        if (!verifyWeight(aggregateWeight, validatorSet.totalWeight)) {
            return false;
        }
        return BLST.verifySignature(aggregateKey, signature.signature, message);
    }

    /**
     * @notice Parses the validators from a serialized validator set
     * @param data The serialized validator set. The serialized format is:
     * - 2 bytes: codec ID (0x0000)
     * - 4 bytes: number of validators
     * - 96 bytes per validator: unformatted BLS public key
     * - 8 bytes per validator: weight
     * @return validators The parsed validators
     * @return totalWeight The total weight of all validators
     */
    function parseValidators(
        bytes calldata data
    ) public pure returns (Validator[] memory, uint64) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the number of validators
        uint32 numValidators = uint32(bytes4(data[2:2 + NUM_VALIDATOR_LENGTH]));

        // Parse the validators
        Validator[] memory validators = new Validator[](numValidators);
        uint64 totalWeight = 0;
        // The position in the bytes were are currently parsing
        uint256 currentPosition = 6;
        // we will require that public keys appear in increasing order,
        // lexicographically sorted as bytes
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);

        for (uint32 i = 0; i < numValidators; i++) {
            bytes memory unformattedPublicKey = data[
                currentPosition:currentPosition + BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH
            ];
            require(
                BLST.comparePublicKeys(unformattedPublicKey, previousPublicKey) > 0,
                "BLS public key must be greater than the latest public key"
            );
            // A serialized validator consists of
            // VALIDATOR_BYTES = BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH + VALIDATOR_WEIGHT_LENGTH
            // bytes. To get the the weight, we skip the first  BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH
            // bytes up to VALIDATOR_BYTES from the current byte position.
            uint64 weight = uint64(
                bytes8(
                    data[
                        currentPosition + BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH:
                            currentPosition + VALIDATOR_BYTES
                    ]
                )
            );
            require(weight > 0, "Validator weight must be greater than 0");
            validators[i] = Validator({
                blsPublicKey: BLST.padUncompressedBLSPublicKey(unformattedPublicKey),
                weight: weight
            });
            previousPublicKey = unformattedPublicKey;
            totalWeight += weight;
            // move the current position past the bytes of this validator
            currentPosition += VALIDATOR_BYTES;
        }
        return (validators, totalWeight);
    }

    function serializeValidators(
        Validator[] memory validators
    ) public pure returns (bytes memory) {
        bytes memory serialized = new bytes(2 + 4 + validators.length * VALIDATOR_BYTES);

        assembly ("memory-safe") {
            let validator_length := mload(validators)
            //encode the number of validators
            // convert bytes32 to bytes4 and store the length
            mstore(add(serialized, 0x22), shl(224, validator_length))
        }

        // encode the validators
        uint256 offset = 6;
        for (uint256 i = 0; i < validators.length; i++) {
            // encode the 96-bytes uncompressed BLS public key
            bytes memory uncompressedBlsPublicKey =
                BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey);
            assembly ("memory-safe") {
                mstore(
                    add(serialized, add(offset, 0x20)), mload(add(uncompressedBlsPublicKey, 0x20))
                )
                mstore(
                    add(serialized, add(offset, 0x40)), mload(add(uncompressedBlsPublicKey, 0x40))
                )
                mstore(
                    add(serialized, add(offset, 0x60)), mload(add(uncompressedBlsPublicKey, 0x60))
                )
            }
            offset += BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH;
            uint64 weight = validators[i].weight;
            assembly ("memory-safe") {
                // encode the validator weight
                // shift left to convert bytes32 to bytes8
                mstore(add(serialized, add(offset, 0x20)), shl(192, weight))
            }
            offset += VALIDATOR_WEIGHT_LENGTH;
        }
        return serialized;
    }

    /**
     * @notice Deserializes the validator state payload
     * @param data The serialized validator state payload. The serialized format is:
     * - 2 bytes: codec ID (0x0000)
     * - 4 bytes: payload type id
     * - 32 bytes: Avalanche blockchain ID
     * - 8 bytes: Avalanche P-chain height
     * - 8 bytes: Avalanche P-chain timestamp
     * - the remainder is an abi-encoded array of hashes
     * @return ValidatorSetStatePayload instance
     */
    function parseValidatorSetMetadata(
        bytes calldata data
    ) public pure returns (ValidatorSetMetadata memory) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the payload type ID, and confirm it is 4 for ValidatorSetState
        uint32 payloadTypeID = uint32(bytes4(data[2:2 + PAYLOAD_TYPE_ID_LENGTH]));
        require(
            payloadTypeID == PAYLOAD_TYPE_ID_LENGTH, "Invalid ValidatorSetState payload type ID"
        );

        // Parse the avalancheBlockchainID
        bytes32 avalancheBlockchainID = bytes32(data[2 + PAYLOAD_TYPE_ID_LENGTH:38]);

        // Parse the pChainHeight
        uint64 pChainHeight = uint64(bytes8(data[38:46]));

        // Parse the pChainTimestamp`
        uint64 pChainTimestamp = uint64(bytes8(data[46:54]));

        // Parse the shardHashes
        bytes32[] memory shardHashes = abi.decode(data[54:], (bytes32[]));

        return ValidatorSetMetadata({
            avalancheBlockchainID: avalancheBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            shardHashes: shardHashes
        });
    }

    /**
     * @notice Parses a ValidatorSetDiff payload from serialized bytes
     * @dev The payload type ID must be 5 for ValidatorSetDiff
     * @param data The serialized ValidatorSetDiff payload
     * @return The parsed ValidatorSetDiffPayload
     */
    function parseValidatorSetDiffPayload(
        bytes memory data
    ) public pure returns (ValidatorSetDiffPayload memory) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the payload type ID, and confirm it is 5 for ValidatorSetDiff
        uint32 payloadTypeID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));
        require(payloadTypeID == 5, "Invalid ValidatorSetDiff payload type ID");

        uint256 offset = 6;

        // Parse the avalancheBlockchainID (32 bytes)
        bytes32 avalancheBlockchainID = abi.decode(ByteSlicer.slice(data, offset, 32), (bytes32));
        offset += 32;

        // Parse previous state
        uint64 previousHeight = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;
        uint64 previousTimestamp = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;
        bytes32 previousValidatorSetHash = abi.decode(ByteSlicer.slice(data, offset, 32), (bytes32));
        offset += 32;

        // Parse current state
        uint64 currentHeight = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;
        uint64 currentTimestamp = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;
        bytes32 currentValidatorSetHash = abi.decode(ByteSlicer.slice(data, offset, 32), (bytes32));
        offset += 32;

        // Parse added validators
        uint32 numAdded = uint32(bytes4(ByteSlicer.slice(data, offset, 4)));
        offset += 4;
        ValidatorChange[] memory added = new ValidatorChange[](numAdded);
        for (uint32 i = 0; i < numAdded; i++) {
            (added[i], offset) = parseValidatorChange(data, offset);
        }

        // Parse removed validators
        uint32 numRemoved = uint32(bytes4(ByteSlicer.slice(data, offset, 4)));
        offset += 4;
        ValidatorChange[] memory removed = new ValidatorChange[](numRemoved);
        for (uint32 i = 0; i < numRemoved; i++) {
            (removed[i], offset) = parseValidatorChange(data, offset);
        }

        // Parse modified validators
        uint32 numModified = uint32(bytes4(ByteSlicer.slice(data, offset, 4)));
        offset += 4;
        ValidatorChange[] memory modified = new ValidatorChange[](numModified);
        for (uint32 i = 0; i < numModified; i++) {
            (modified[i], offset) = parseValidatorChange(data, offset);
        }

        return ValidatorSetDiffPayload({
            avalancheBlockchainID: avalancheBlockchainID,
            previousHeight: previousHeight,
            previousTimestamp: previousTimestamp,
            previousValidatorSetHash: previousValidatorSetHash,
            currentHeight: currentHeight,
            currentTimestamp: currentTimestamp,
            currentValidatorSetHash: currentValidatorSetHash,
            added: added,
            removed: removed,
            modified: modified
        });
    }

    /**
     * @notice Parses a single ValidatorChange from serialized bytes
     * @param data The serialized data
     * @param offset The offset to start parsing from
     * @return change The parsed ValidatorChange
     * @return newOffset The new offset after parsing
     */
    function parseValidatorChange(
        bytes memory data,
        uint256 offset
    ) public pure returns (ValidatorChange memory change, uint256 newOffset) {
        // Parse nodeID (20 bytes)
        bytes20 nodeID = bytes20(ByteSlicer.slice(data, offset, 20));
        offset += 20;

        // Parse uncompressed BLS public key (96 bytes)
        bytes memory unformattedPublicKey = ByteSlicer.slice(data, offset, 96);
        bytes memory blsPublicKey = BLST.formatUncompressedBLSPublicKey(unformattedPublicKey);
        offset += 96;

        // Parse previousWeight (8 bytes)
        uint64 previousWeight = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;

        // Parse currentWeight (8 bytes)
        uint64 currentWeight = uint64(bytes8(ByteSlicer.slice(data, offset, 8)));
        offset += 8;

        change = ValidatorChange({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            previousWeight: previousWeight,
            currentWeight: currentWeight
        });

        return (change, offset);
    }

    /**
     * @notice Applies a diff to a validator array
     * @dev Creates a new validator array with the diff applied
     * @param currentValidators The current validator array
     * @param diff The diff to apply
     * @return newValidators The new validator array after applying the diff
     * @return newTotalWeight The new total weight
     */
    function applyValidatorSetDiff(
        Validator[] memory currentValidators,
        ValidatorSetDiffPayload memory diff
    ) public pure returns (Validator[] memory newValidators, uint64 newTotalWeight) {
        // Calculate new size: current - removed + added
        uint256 newSize = currentValidators.length - diff.removed.length + diff.added.length;
        newValidators = new Validator[](newSize);
        newTotalWeight = 0;

        // Create a mapping of validators to remove (by public key hash)
        // We'll use a O(n * logm) approach since the validators are sorted by public key
        uint256 newIndex = 0;

        // Copy validators that are not being removed, applying modifications
        for (uint256 i = 0; i < currentValidators.length; i++) {
            bytes memory currentKey = currentValidators[i].blsPublicKey;

            // Check if this validator is being removed
            (bool isRemoved,) = _searchValidators(diff.removed, currentKey);
            if (isRemoved) continue;

            // Check if this validator is being modified
            (bool isModified, ValidatorChange memory mod) =
                _searchValidators(diff.modified, currentKey);
            uint64 weight = isModified ? mod.currentWeight : currentValidators[i].weight;

            // Add to new list
            require(newIndex < newValidators.length, "Index out of bounds.");
            newValidators[newIndex] = Validator({blsPublicKey: currentKey, weight: weight});
            newTotalWeight += weight;
            newIndex++;
        }

        // Add new validators
        for (uint256 i = 0; i < diff.added.length; i++) {
            newValidators[newIndex] = Validator({
                blsPublicKey: diff.added[i].blsPublicKey,
                weight: diff.added[i].currentWeight
            });
            newTotalWeight += diff.added[i].currentWeight;
            newIndex++;
        }

        // Sort validators by public key (required for canonical ordering)
        _sortValidators(newValidators);

        return (newValidators, newTotalWeight);
    }

    /*
     * @notice Serialize a ValidatorSetStatePayload
     */
    function serializeValidatorSetMetadata(
        ValidatorSetMetadata memory payload
    ) public pure returns (bytes memory) {
        bytes2 codec = bytes2(0);
        bytes4 payloadType = bytes4(0x00000004);
        return abi.encodePacked(
            codec,
            payloadType,
            payload.avalancheBlockchainID,
            payload.pChainHeight,
            payload.pChainTimestamp,
            abi.encode(payload.shardHashes)
        );
    }

    /*
     * @notice Serialize a ValidatorSetDiffPayload
     */
    function serializeValidatorSetDiffPayload(
        ValidatorSetDiffPayload memory payload
    ) public pure returns (bytes memory) {
        bytes2 codec = bytes2(0);
        bytes4 payloadType = bytes4(0x00000005);
        bytes memory data = abi.encodePacked(
            codec,
            payloadType,
            payload.avalancheBlockchainID,
            payload.previousHeight,
            payload.previousTimestamp,
            payload.previousValidatorSetHash,
            payload.currentHeight,
            payload.currentTimestamp,
            payload.currentValidatorSetHash,
            uint32(payload.added.length)
        );
        // Encode added validators
        for (uint256 i = 0; i < payload.added.length; i++) {
            data = abi.encodePacked(data, serializeValidatorChange(payload.added[i]));
        }
        // Encode removed validators
        data = abi.encodePacked(data, uint32(payload.removed.length));
        for (uint256 i = 0; i < payload.removed.length; i++) {
            data = abi.encodePacked(data, serializeValidatorChange(payload.removed[i]));
        }
        // Encode modified validators
        data = abi.encodePacked(data, uint32(payload.modified.length));
        for (uint256 i = 0; i < payload.modified.length; i++) {
            data = abi.encodePacked(data, serializeValidatorChange(payload.modified[i]));
        }
        return data;
    }

    /**
     * @notice Serializes a single ValidatorChange
     */
    function serializeValidatorChange(
        ValidatorChange memory change
    ) public pure returns (bytes memory) {
        return abi.encodePacked(
            change.nodeID, change.blsPublicKey, change.previousWeight, change.currentWeight
        );
    }

    /*
     * @notice Deserialize bytes into `ValidatorSetSignature`
     */
    function parseValidatorSetSignature(
        bytes calldata signatureBytes
    ) public pure returns (ValidatorSetSignature memory) {
        bytes memory signers = signatureBytes[0:signatureBytes.length - BLST.BLS_SIGNATURE_LENGTH];
        bytes memory signature = signatureBytes[signatureBytes.length - BLST.BLS_SIGNATURE_LENGTH:];
        return ValidatorSetSignature({signers: signers, signature: signature});
    }

    /*
     * @notice Serialize `ValidatorSetSignature` to bytes
     */
    function serializeValidatorSetSignature(
        ValidatorSetSignature memory signature
    ) public pure returns (bytes memory) {
        return abi.encodePacked(signature.signers, signature.signature);
    }

    /*
     * @notice Deserialize bytes into `ValidatorSetShard`
     */
    function parseValidatorSetShard(
        bytes calldata shardBytes
    ) public pure returns (ValidatorSetShard memory) {
        uint64 shardNumber = uint64(bytes8(shardBytes[0:8]));
        bytes32 avalancheBlockchainID = bytes32(shardBytes[8:40]);
        return ValidatorSetShard({
            shardNumber: shardNumber,
            avalancheBlockchainID: avalancheBlockchainID
        });
    }

    /*
     * @notice Serialize `ValidatorSetShard` to bytes
     */
    function serializeValidatorSetShard(
        ValidatorSetShard memory shard
    ) public pure returns (bytes memory) {
        return abi.encodePacked(shard.shardNumber, shard.avalancheBlockchainID);
    }

    /*
     * @dev Traverse the bits in signers from left to right, using it as bitvector to determine
     * which validators to select from the provided list.
     * @return The aggregate public key and stake weight of the filtered validators
     */
    function filterValidators(
        bytes memory signers,
        Validator[] memory validators
    ) internal view returns (bytes memory, uint64) {
        bytes memory aggregatePublicKey;
        uint64 aggregateWeight = 0;
        if (validators.length == 0) {
            revert("Cannot validate against an empty list of validators");
        }

        uint256 byteIndex = 0;
        uint8 bitMask = 1 << 7;
        uint8 currentByte = uint8(signers[byteIndex]);

        // we traverse the validator set from left to right
        for (uint256 i = 0; i < validators.length; i++) {
            // check if the bit is set
            if (currentByte & bitMask == bitMask) {
                Validator memory validator = validators[i];
                if (aggregateWeight > 0) {
                    aggregatePublicKey = BLST.addG1(aggregatePublicKey, validator.blsPublicKey);
                } else {
                    aggregatePublicKey = validator.blsPublicKey;
                }
                aggregateWeight += validator.weight;
            }

            // shift one bit to the right
            bitMask = bitMask >> 1;
            if (bitMask == 0) {
                byteIndex += 1;
                currentByte = uint8(signers[byteIndex]);
                bitMask = 1 << 7;
            }
        }
        return (aggregatePublicKey, aggregateWeight);
    }

    /*
     * @notice Verifies that quorumNum * totalWeight <= quorumDen * signatureWeight
     */
    function verifyWeight(
        uint64 signatureWeight,
        uint64 totalWeight
    ) internal pure returns (bool) {
        uint256 scaledTotalWeight = QUORUM_NUM * uint256(totalWeight);
        uint256 scaledSignatureWeight = QUORUM_DEN * uint256(signatureWeight);
        return scaledTotalWeight <= scaledSignatureWeight;
    }

    /**
     * @notice Generic Binary Search.
     * @return found True if the key exists.
     * @return item The ValidatorChange struct found (or empty if not found).
     */
    function _searchValidators(
        ValidatorChange[] memory list,
        bytes memory key
    ) internal pure returns (bool found, ValidatorChange memory item) {
        if (list.length == 0) return (false, item);

        uint256 low = 0;
        uint256 high = list.length - 1;

        while (low <= high) {
            uint256 mid = (low + high) / 2;
            int256 cmp = ByteComparator.compare(list[mid].blsPublicKey, key);

            if (cmp == 0) {
                return (true, list[mid]);
            }

            if (cmp < 0) {
                low = mid + 1;
            } else {
                if (mid == 0) break;
                high = mid - 1;
            }
        }
        return (false, item);
    }

    /**
     * @notice Sorts validators by their uncompressed public key bytes
     * @dev Uses insertion sort which is efficient for small arrays
     */
    function _sortValidators(
        Validator[] memory validators
    ) private pure {
        for (uint256 i = 1; i < validators.length; i++) {
            Validator memory key = validators[i];
            bytes memory keyPubKey = BLST.getUncompressedBlsPublicKey(key.blsPublicKey);
            int256 j = int256(i) - 1;

            while (j >= 0) {
                bytes memory jPubKey =
                    BLST.getUncompressedBlsPublicKey(validators[uint256(j)].blsPublicKey);
                if (ByteComparator.compare(jPubKey, keyPubKey) <= 0) {
                    break;
                }
                validators[uint256(j + 1)] = validators[uint256(j)];
                j--;
            }
            validators[uint256(j + 1)] = key;
        }
    }
}
