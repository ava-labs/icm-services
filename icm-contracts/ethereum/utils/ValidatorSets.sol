// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

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
}
