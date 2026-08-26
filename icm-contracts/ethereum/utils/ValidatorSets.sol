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

// The payload that initiates registering / updating a
// Validator set for the chain with ID `avalancheBlockchainID`.
// A shard to add to a validator set for which partial data has already
// been received. The `avalancheBlockchainID` field will be used to lookup the
// existing partial data
// A partial Validator set which can be constructed from
//  shards sent across multiple transactions
// ValidatorChange represents a single validator addition, removal, or modification
// ValidatorSetDiff contains the diff information from a signed message
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
        if (
            !MerkleProof.multiProofVerify(
                att.proof, att.proofFlags, comm.root, leaves, keccakInternalPair
            )
        ) {
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

    function keccakInternalPair(bytes32 a, bytes32 b) internal pure returns (bytes32) {
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
