// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ByteSlicer} from "./ByteSlicer.sol";
import {BLST} from "./BLST.sol";
import {ByteComparator} from "./ByteComparator.sol";

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

struct ValidatorSetStatePayload {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    bytes32 validatorSetHash;
}

// ValidatorChange represents a single validator addition, removal, or modification
struct ValidatorChange {
    bytes20 nodeID;           // 20 bytes
    bytes blsPublicKey;       // 96 bytes uncompressed
    uint64 previousWeight;    // Weight at previous height (0 for additions)
    uint64 currentWeight;     // Weight at current height (0 for removals)
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
        bytes memory data
    ) internal pure returns (Validator[] memory, uint64) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the number of validators
        uint32 numValidators = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));

        // Parse the validators
        Validator[] memory validators = new Validator[](numValidators);
        uint64 totalWeight = 0;
        uint256 offset = 6;
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        for (uint32 i = 0; i < numValidators; i++) {
            bytes memory unformattedPublicKey = ByteSlicer.slice(data, offset, 96);
            require(
                ByteComparator.compare(unformattedPublicKey, previousPublicKey) > 0,
                "BLS public key must be greater than the latest public key"
            );
            uint64 weight = uint64(bytes8(ByteSlicer.slice(data, offset + 96, 8)));
            require(weight > 0, "Validator weight must be greater than 0");
            validators[i] = Validator({
                blsPublicKey: BLST.formatUncompressedBLSPublicKey(unformattedPublicKey),
                weight: weight
            });
            totalWeight += weight;
            offset += 104;
        }
        return (validators, totalWeight);
    }

    function serializeValidators(
        Validator[] memory validators
    ) public pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            2 + 4 + validators.length * (96 + 8)
        );
        // the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        //encode the number of validators
        bytes4 numValidators = bytes4(uint32(validators.length));
        for (uint256 i = 0; i < 4; i++) {
            serialized[2 + i] = numValidators[i];
        }
        // encode the validators
        uint256 offset = 6;
        for (uint256 i = 0; i < validators.length; i++) {
            // encode the 96-bytes uncompressed BLS public key
            bytes memory uncompressedBlsPublicKey = BLST.getUncompressedBlsPublicKey(validators[i].blsPublicKey);
            for(uint256 j = 0; j < 96; j++ ) {
                serialized[j + offset] = uncompressedBlsPublicKey[j];
            }
            offset += 96;
            // encode the validator weight
            bytes8 weight = bytes8(validators[i].weight);
            for(uint256 j = 0; j < 8; j++ ) {
                serialized[j + offset] = weight[j];
            }
            offset += 8;
        }
        return serialized;
    }

    function parseValidatorSetStatePayload(
        bytes memory data
    ) internal pure returns (ValidatorSetStatePayload memory) {
        // Check the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the payload type ID, and confirm it is 4 for ValidatorSetState
        uint32 payloadTypeID = uint32(bytes4(ByteSlicer.slice(data, 2, 4)));
        require(payloadTypeID == 4, "Invalid ValidatorSetState payload type ID");

        // Parse the avalancheBlockchainID
        bytes32 avalancheBlockchainID = abi.decode(ByteSlicer.slice(data, 6, 32), (bytes32));

        // Parse the pChainHeight
        uint64 pChainHeight = uint64(bytes8(ByteSlicer.slice(data, 38, 8)));

        // Parse the pChainTimestamp
        uint64 pChainTimestamp = uint64(bytes8(ByteSlicer.slice(data, 46, 8)));

        // Parse the validatorSetHash
        bytes32 validatorSetHash = abi.decode(ByteSlicer.slice(data, 54, 32), (bytes32));

        return ValidatorSetStatePayload({
            avalancheBlockchainID: avalancheBlockchainID,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp,
            validatorSetHash: validatorSetHash
        });
    }

    function serializeValidatorSetStatePayload(
        ValidatorSetStatePayload memory payload
    ) public pure returns (bytes memory) {
        bytes memory serialized = new bytes(
            // codec
            2
            // payload type ID
            + 4
            // Avalanche block chain ID
            + 32
            // the P-chain height
            + 8
            // the P-chain timestamp
            + 8
            // hash of validator set
            + 32
        );
        // encode the codec
        serialized[0] = 0x00;
        serialized[1] = 0x00;
        // the payload type ID
        serialized[2] = 0x00;
        serialized[3] = 0x00;
        serialized[4] = 0x00;
        serialized[5] = 0x04;
        // encode avalancheBlockchainID
        for (uint256 i = 0; i < 32; i++) {
            serialized[6 + i] = payload.avalancheBlockchainID[i];
        }
        // encode the P-chain block height
        bytes8 pChainHeight = bytes8(payload.pChainHeight);
        for (uint256 i = 0; i < 8; i++) {
            serialized[38 + i] = pChainHeight[i];
        }
        // encode the P-chain timestamp
        bytes8 pChainTimestamp = bytes8(payload.pChainTimestamp);
        for (uint256 i = 0; i < 8; i++) {
            serialized[46 + i] = pChainTimestamp[i];
        }
        // encode the validator set hash
        for (uint256 i = 0; i < 32; i++) {
            serialized[54 + i] = payload.validatorSetHash[i];
        }
        return serialized;
    }

    /**
     * @notice Parses a ValidatorSetDiff payload from serialized bytes
     * @dev The payload type ID must be 5 for ValidatorSetDiff
     * @param data The serialized ValidatorSetDiff payload
     * @return The parsed ValidatorSetDiffPayload
     */
    function parseValidatorSetDiffPayload(
        bytes memory data
    ) internal pure returns (ValidatorSetDiffPayload memory) {
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
    ) internal pure returns (ValidatorChange memory change, uint256 newOffset) {
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
    function applyDiff(
        Validator[] memory currentValidators,
        ValidatorSetDiffPayload memory diff
    ) internal pure returns (Validator[] memory newValidators, uint64 newTotalWeight) {
        // Calculate new size: current - removed + added
        uint256 newSize = currentValidators.length - diff.removed.length + diff.added.length;
        newValidators = new Validator[](newSize);
        newTotalWeight = 0;

        // Create a mapping of validators to remove (by public key hash)
        // We'll use a simple O(n*m) approach since arrays are typically small
        uint256 newIndex = 0;

        // Copy validators that are not being removed, applying modifications
        for (uint256 i = 0; i < currentValidators.length; i++) {
            bool isRemoved = false;
            
            // Check if this validator is being removed
            for (uint256 j = 0; j < diff.removed.length; j++) {
                if (keccak256(currentValidators[i].blsPublicKey) == keccak256(diff.removed[j].blsPublicKey)) {
                    isRemoved = true;
                    break;
                }
            }
            
            if (!isRemoved) {
                // Check if this validator is being modified
                uint64 weight = currentValidators[i].weight;
                for (uint256 j = 0; j < diff.modified.length; j++) {
                    if (keccak256(currentValidators[i].blsPublicKey) == keccak256(diff.modified[j].blsPublicKey)) {
                        weight = diff.modified[j].currentWeight;
                        break;
                    }
                }
                
                newValidators[newIndex] = Validator({
                    blsPublicKey: currentValidators[i].blsPublicKey,
                    weight: weight
                });
                newTotalWeight += weight;
                newIndex++;
            }
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

    /**
     * @notice Sorts validators by their uncompressed public key bytes
     * @dev Uses insertion sort which is efficient for small arrays
     */
    function _sortValidators(Validator[] memory validators) private pure {
        for (uint256 i = 1; i < validators.length; i++) {
            Validator memory key = validators[i];
            bytes memory keyPubKey = BLST.getUncompressedBlsPublicKey(key.blsPublicKey);
            int256 j = int256(i) - 1;
            
            while (j >= 0) {
                bytes memory jPubKey = BLST.getUncompressedBlsPublicKey(validators[uint256(j)].blsPublicKey);
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
