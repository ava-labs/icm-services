// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {BLST} from "./BLST.sol";
import {Math} from "@openzeppelin/contracts@5.1.0/utils/math/Math.sol";
import {MerkleProof} from "@openzeppelin/contracts@5.1.0/utils/cryptography/MerkleProof.sol";

/**
 * THIS IS LIBRARY IS UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
struct Validator {
    bytes blsPublicKey;
    uint64 weight;
}

/// Compact, constant size on-chain commitment to an Avalanche validator set.
struct ValidatorSetMerkleCommitment {
    bytes32 avalancheBlockchainID;
    bytes32 root;
    uint64 totalWeight;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
}

/// Attestation envelope for a BLS-signed ICM message. Contains the
/// signing validators for a specific chain, a Merkle multi-inclusion proof binding
/// them to the registry's stored root, and the aggregate BLS signature.
/// @dev The signers must be stored in increasing lexicographic order by their BLS public key.
struct ValidatorSetMerkleAttestation {
    Validator[] signers;
    bytes32[] proof;
    bool[] proofFlags;
    bytes aggregateBlsSig;
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
     * @dev The (de)serialization codec prefix
     */
    bytes2 private constant CODEC = bytes2(0);

    /*
     * @dev The number of bytes in the `Validator` struct
     */
    uint256 private constant VALIDATOR_BYTES =
        BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH + VALIDATOR_WEIGHT_LENGTH;

    /**
     * @dev ICM message verification scheme using Merkle attestations. Includes the following steps:
     * 1. Reconstruct each signer's leaf
     * 2. Verify all leaves against the stored root via a Merkle multi-inclusion proof
     * 3. Enforce uniqueness of signers
     * 4. Perform a stake-weighted quorum threshold check
     * 5. Perform aggregate BLS signature verification
     */
    function verifyMerkleAttestation(
        bytes calldata rawAttestation,
        bytes memory signedData,
        ValidatorSetMerkleCommitment storage comm
    ) internal view returns (bool) {
        ValidatorSetMerkleAttestation memory att =
            ValidatorSets.parseMerkleAttestation(rawAttestation);
        uint256 numSigners = att.signers.length;
        require(numSigners > 0, "No signers");

        // Reconstruct leaves
        bytes32[] memory leaves = new bytes32[](numSigners);
        for (uint256 i = 0; i < numSigners;) {
            leaves[i] = hashValidator(att.signers[i]);
            unchecked {
                ++i;
            }
        }

        // Perform Merkle multi-inclusion proof verification against the stored root
        // TODO: Switch to multiProofVerifyCalldata once parseMerkleAttestation returns calldata slice offsets instead of a memory struct for additional gas savings
        if (!MerkleProof.multiProofVerify(
                att.proof, att.proofFlags, comm.root, leaves, keccakInternalPair
            )) {
            return false;
        }

        // Stake-weighted quorum threshold check
        uint64 signerWeight = 0;
        for (uint256 i = 0; i < numSigners; ++i) {
            signerWeight += att.signers[i].weight;
        }
        if (!ValidatorSets.verifyWeight(signerWeight, comm.totalWeight)) {
            return false;
        }

        // Aggregate pubkeys and verify BLS signature over the warp preimage.
        bytes memory aggregateKey = att.signers[0].blsPublicKey;
        for (uint256 i = 1; i < numSigners; ++i) {
            aggregateKey = BLST.addG1(aggregateKey, att.signers[i].blsPublicKey);
        }
        return BLST.verifySignature(aggregateKey, att.aggregateBlsSig, signedData);
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
    ) internal pure returns (Validator[] memory, uint64) {
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
                    data[currentPosition
                            + BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH:currentPosition
                                + VALIDATOR_BYTES
                    ]
                )
            );
            require(weight > 0, "Validator weight must be greater than 0");
            validators[i] = Validator({
                blsPublicKey: BLST.padUncompressedBLSPublicKey(unformattedPublicKey), weight: weight
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
    ) internal pure returns (bytes memory) {
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
                    add(serialized, add(offset, 0x20)),
                    mload(add(uncompressedBlsPublicKey, 0x20))
                )
                mstore(
                    add(serialized, add(offset, 0x40)),
                    mload(add(uncompressedBlsPublicKey, 0x40))
                )
                mstore(
                    add(serialized, add(offset, 0x60)),
                    mload(add(uncompressedBlsPublicKey, 0x60))
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

    /*
    * @dev Serializes a ValidatorSetMerkleAttestation in the format expected by parseMerkleAttestation.
    */
    function serializeMerkleAttestation(
        ValidatorSetMerkleAttestation memory att
    ) internal pure returns (bytes memory) {
        bytes memory data = abi.encodePacked(CODEC, uint32(att.signers.length));
        // Encode public keys
        for (uint256 i = 0; i < att.signers.length; i++) {
            data = abi.encodePacked(
                data,
                BLST.unPadUncompressedBlsPublicKey(att.signers[i].blsPublicKey),
                att.signers[i].weight
            );
        }
        // Encode proof
        data = abi.encodePacked(data, uint32(att.proof.length));
        for (uint256 i = 0; i < att.proof.length; i++) {
            data = abi.encodePacked(data, att.proof[i]);
        }
        // Encode proof flags
        uint32 numFlags = uint32(att.proofFlags.length);
        uint256 flagBytesLen = Math.ceilDiv(numFlags, 8);
        bytes memory packedFlags = new bytes(flagBytesLen);
        for (uint256 i = 0; i < numFlags; i++) {
            if (att.proofFlags[i]) {
                packedFlags[i >> 3] |= bytes1(uint8(1) << uint8(i & 7));
            }
        }
        data = abi.encodePacked(data, numFlags, packedFlags);
        return abi.encodePacked(data, att.aggregateBlsSig);
    }

    /**
     * @notice Serializes a ValidatorSetMerkleCommitment payload.
     */
    function serializeMerkleCommitment(
        ValidatorSetMerkleCommitment memory commitment
    ) internal pure returns (bytes memory) {
        bytes4 payloadType = bytes4(0x00000006);
        return abi.encodePacked(
            CODEC,
            payloadType,
            commitment.avalancheBlockchainID,
            commitment.root,
            commitment.totalWeight,
            commitment.pChainHeight,
            commitment.pChainTimestamp
        );
    }

    /**
     * @notice Parses a ValidatorSetMerkleAttestation from serialized bytes
     * @param data The serialized attestation
     * @return att The parsed ValidatorSetMerkleAttestation
     */
    function parseMerkleAttestation(
        bytes calldata data
    ) internal pure returns (ValidatorSetMerkleAttestation memory att) {
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");
        uint256 offset = 2;

        // Parse the signers
        {
            uint32 numSigners = uint32(bytes4(data[offset:offset + NUM_VALIDATOR_LENGTH]));
            offset += NUM_VALIDATOR_LENGTH;
            att.signers = new Validator[](numSigners);
            for (uint32 i = 0; i < numSigners;) {
                bytes memory unformattedPublicKey =
                    data[offset:offset + BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH];
                offset += BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH;
                uint64 weight = uint64(bytes8(data[offset:offset + VALIDATOR_WEIGHT_LENGTH]));
                offset += VALIDATOR_WEIGHT_LENGTH;
                att.signers[i] = Validator({
                    blsPublicKey: BLST.padUncompressedBLSPublicKey(unformattedPublicKey),
                    weight: weight
                });
                unchecked {
                    ++i;
                }
            }
        }
        // Parse the proof
        {
            uint32 numProofHashes = uint32(bytes4(data[offset:offset + 4]));
            offset += 4;
            att.proof = new bytes32[](numProofHashes);
            for (uint32 i = 0; i < numProofHashes;) {
                att.proof[i] = bytes32(data[offset:offset + 32]);
                offset += 32;
                unchecked {
                    ++i;
                }
            }
            uint32 numFlags = uint32(bytes4(data[offset:offset + 4]));
            offset += 4;
            uint256 flagBytesLen = Math.ceilDiv(numFlags, 8);

            // Parse the proof flags
            att.proofFlags = new bool[](numFlags);
            for (uint256 i = 0; i < numFlags;) {
                uint8 byteVal = uint8(data[offset + (i >> 3)]);
                att.proofFlags[i] = (byteVal >> uint8(i & 7)) & 1 == 1;
                unchecked {
                    ++i;
                }
            }

            offset += flagBytesLen;
        }
        // Parse the signature
        att.aggregateBlsSig = data[offset:offset + BLST.BLS_SIGNATURE_LENGTH];
        offset += BLST.BLS_SIGNATURE_LENGTH;
        require(offset == data.length, "Trailing bytes in attestation");
        return att;
    }

    /**
     * @notice Parses a ValidatorSetMerkleCommitment payload from serialized bytes.
     * @dev The serialized format is:
     * - 2 bytes:  codec ID
     * - 4 bytes:  payload type ID
     * - 32 bytes: Avalanche blockchain ID
     * - 32 bytes: Merkle root over the validator set
     * - 8 bytes:  total validator weight
     * - 8 bytes:  P-chain height
     * - 8 bytes:  P-chain timestamp
     */
    function parseMerkleCommitment(
        bytes calldata data
    ) internal pure returns (ValidatorSetMerkleCommitment memory) {
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");
        uint32 payloadTypeID = uint32(bytes4(data[2:6]));
        require(payloadTypeID == 6, "Invalid ValidatorSetMerkleCommitment payload type ID");

        return ValidatorSetMerkleCommitment({
            avalancheBlockchainID: bytes32(data[6:38]),
            root: bytes32(data[38:70]),
            totalWeight: uint64(bytes8(data[70:78])),
            pChainHeight: uint64(bytes8(data[78:86])),
            pChainTimestamp: uint64(bytes8(data[86:94]))
        });
    }

    /**
     * @notice Applies a presorted diff to a presorted validator array in O(N+M) time.
     * @param currentValSet The existing validators (must be sorted by BLS public key).
     * @param diff The diff containing the changes and newSize (must be sorted by BLS public key).
     * @return modified The updated validator array.
     * @return newTotalWeight The total weight of the new set.
     */
    function applyValidatorSetDiff(
        Validator[] memory currentValSet,
        ValidatorSetDiff memory diff
    ) internal pure returns (Validator[] memory, uint64) {
        uint64 newTotalWeight = 0;
        uint256 newSize = diff.newSize;
        Validator[] memory modified = new Validator[](newSize);
        uint256 valSetPos = 0;
        uint256 changePos = 0;

        _verifySortedValidatorChanges(diff.changes);

        for (uint256 i = 0; i < newSize;) {
            int256 compare;
            ValidatorChange memory nextChange;
            Validator memory nextVal;

            // Safety checks
            if (changePos >= diff.changes.length) {
                // case 1: we've exhausted the changes, so we just take the remaining validators
                nextVal = currentValSet[valSetPos];
                compare = -1;
            } else if (valSetPos >= currentValSet.length) {
                // case 2: we've exhausted the current validator set, so we just take the remaining changes (which must all be additions)
                nextChange = diff.changes[changePos];
                compare = 1;
            } else {
                // case 3: we have both validators and changes left, so we compare the next validator and next change
                nextVal = currentValSet[valSetPos];
                nextChange = diff.changes[changePos];
                compare = BLST.comparePublicKeys(
                    BLST.unPadUncompressedBlsPublicKey(nextVal.blsPublicKey),
                    BLST.unPadUncompressedBlsPublicKey(nextChange.blsPublicKey)
                );
            }

            // Compare
            if (compare < 0) {
                // case 1: the next change occurs after this validator
                modified[i] = nextVal;
                newTotalWeight += nextVal.weight;
                valSetPos++;
                unchecked {
                    ++i;
                }
            } else if (compare > 0) {
                // case 2. this is a new validator
                modified[i] =
                    Validator({blsPublicKey: nextChange.blsPublicKey, weight: nextChange.weight});
                newTotalWeight += nextChange.weight;
                changePos++;
                unchecked {
                    ++i;
                }
            } else {
                // case 3. this is a change to an existing validator
                // If the next weight is zero, skip
                if (nextChange.weight != 0) {
                    nextVal.weight = nextChange.weight;
                    modified[i] = nextVal;
                    newTotalWeight += nextChange.weight;
                    unchecked {
                        ++i;
                    }
                }
                changePos++;
                valSetPos++;
            }
        }
        require(modified[newSize - 1].weight > 0, "Incorrect size given in diff");
        return (modified, newTotalWeight);
    }

    /*
     * @notice Serialize a ValidatorSetStatePayload
     */
    function serializeValidatorSetMetadata(
        ValidatorSetMetadata memory payload
    ) internal pure returns (bytes memory) {
        bytes4 payloadType = bytes4(0x00000004);
        // Shard list: uint32 count then each bytes32 with no padding — same as
        // abi.encodePacked(uint32(length), ...hashes); static-sized elements in
        // abi.encode would add offset/length words, so we use encodePacked only.
        return abi.encodePacked(
            CODEC,
            payloadType,
            payload.avalancheBlockchainID,
            payload.pChainHeight,
            payload.pChainTimestamp,
            uint32(payload.shardHashes.length),
            payload.shardHashes
        );
    }

    /*
     * @notice Serialize a ValidatorSetDiff
     */
    function serializeValidatorSetDiff(
        ValidatorSetDiff memory diff
    ) internal pure returns (bytes memory) {
        bytes4 payloadType = bytes4(0x00000005);
        bytes memory data = abi.encodePacked(
            CODEC,
            payloadType,
            diff.avalancheBlockchainID,
            diff.previousHeight,
            diff.previousTimestamp,
            diff.currentHeight,
            diff.currentTimestamp,
            uint32(diff.changes.length)
        );
        for (uint256 i = 0; i < diff.changes.length; i++) {
            data = abi.encodePacked(data, serializeValidatorChange(diff.changes[i]));
        }
        data = abi.encodePacked(data, uint32(diff.numAdded));
        return data;
    }

    /**
     * @notice Serializes a single ValidatorChange
     */
    function serializeValidatorChange(
        ValidatorChange memory change
    ) internal pure returns (bytes memory) {
        return
            abi.encodePacked(BLST.unPadUncompressedBlsPublicKey(change.blsPublicKey), change.weight);
    }

    /*
     * @notice Deserialize bytes into `ValidatorSetSignature`
     */
    function parseValidatorSetSignature(
        bytes calldata signatureBytes
    ) internal pure returns (ValidatorSetSignature memory) {
        bytes memory signers = signatureBytes[0:signatureBytes.length - BLST.BLS_SIGNATURE_LENGTH];
        bytes memory signature = signatureBytes[signatureBytes.length - BLST.BLS_SIGNATURE_LENGTH:];
        return ValidatorSetSignature({signers: signers, signature: signature});
    }

    /*
     * @notice Serialize `ValidatorSetSignature` to bytes
     */
    function serializeValidatorSetSignature(
        ValidatorSetSignature memory signature
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(signature.signers, signature.signature);
    }

    /*
     * @notice Deserialize bytes into `ValidatorSetShard`
     */
    function parseValidatorSetShard(
        bytes calldata shardBytes
    ) internal pure returns (ValidatorSetShard memory) {
        uint64 shardNumber = uint64(bytes8(shardBytes[0:8]));
        bytes32 avalancheBlockchainID = bytes32(shardBytes[8:40]);
        return
            ValidatorSetShard({
                shardNumber: shardNumber, avalancheBlockchainID: avalancheBlockchainID
            });
    }

    /*
     * @notice Serialize `ValidatorSetShard` to bytes
     */
    function serializeValidatorSetShard(
        ValidatorSetShard memory shard
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(shard.shardNumber, shard.avalancheBlockchainID);
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
     * @dev Reconstructs the unsigned Warp message bytes that validators sign for ICM addressed-call payloads.
     * If senderAddress is address(0), the sender field is omitted from the AddressedCall encoding — this is
     * the case for validator set update messages, which carry no sender address.
     * If senderAddress is non-zero (e.g. address(this) for ICM messages sent by a contract), the
     * warp precompile sets msg.sender as the source address and it must be included in the reconstruction.
     *
     * Layout without sender: warpCodec(2) | networkID(4) | sourceChainID(32) | payloadFieldLen(4)
     *                        | addressedCallCodec(2) | typeID(4) | srcAddrLen(4)=0 | innerPayloadLen(4) | payload
     * Layout with sender:    warpCodec(2) | networkID(4) | sourceChainID(32) | payloadFieldLen(4)
     *                        | addressedCallCodec(2) | typeID(4) | srcAddrLen(4)=20 | srcAddr(20) | innerPayloadLen(4) | payload
     */
    function buildUnsignedWarpMessage(
        uint32 networkID,
        bytes32 sourceBlockchainID,
        address senderAddress,
        bytes memory payload
    ) internal pure returns (bytes memory) {
        if (senderAddress == address(0)) {
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
        return abi.encodePacked(
            bytes2(0),
            networkID,
            sourceBlockchainID,
            uint32(payload.length + 34),
            bytes2(0),
            uint32(1),
            uint32(20),
            bytes20(senderAddress),
            uint32(payload.length),
            payload
        );
    }

    function keccakInternalPair(
        bytes32 a,
        bytes32 b
    ) internal pure returns (bytes32) {
        return a < b
            ? keccak256(abi.encodePacked(uint256(0), a, b))
            : keccak256(abi.encodePacked(uint256(0), b, a));
    }

    /*
     * @notice Prepends a one to the hash of a validator to avoid second pre-image attacks on
     * the merkle trees. This is a convenience function that should be used to hash validators
     * rather than calling sha256 on them directly
     */
    function hashValidator(
        Validator memory val
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(1), val.blsPublicKey, val.weight));
    }
}
