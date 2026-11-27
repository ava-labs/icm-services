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

// #[pack(contract="ValidatorSets", visibility="internal")]
// #[unpack(contract="ValidatorSets", visibility="internal", calldata)]
struct Validator {
    // #[pack(method="BLST.unPadUncompressedBlsPublicKey")]
    // #[unpack(
    //    length=BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH,
    //    method="BLST.padUncompressedBLSPublicKey",
    //  )]
    bytes blsPublicKey;
    uint64 weight;
}

/// Compact, constant size on-chain commitment to an Avalanche validator set.
// #[pack(contract="ValidatorSets", visibility="internal")]
// #[unpack(contract="ValidatorSets", visibility="internal", calldata)]
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
// #[pack(contract="ValidatorSets", visibility="internal")]
// #[unpack(contract="ValidatorSets", visibility="internal", calldata)]
struct ValidatorSetMerkleAttestation {
    // #[pack(length=uint32)]
    // #[unpack(length=uint32)]
    Validator[] signers;
    // #[pack(length=uint32)]
    // #[unpack(length=uint32)]
    bytes32[] proof;
    // #[pack(method="packProofFlags")]
    // #[unpack(method="unpackProofFlags")]
    bool[] proofFlags;
    // #[pack(method="abi.encodePacked")]
    // #[unpack(default)]
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
                att.proof, att.proofFlags, comm.root, leaves, _keccakPair
            )) {
            return false;
        }

        // Stake-weighted quorum threshold check
        uint64 signerWeight = 0;
        for (uint256 i = 0; i < numSigners;) {
            signerWeight += att.signers[i].weight;
            unchecked {
                ++i;
            }
        }
        if (!ValidatorSets.verifyWeight(signerWeight, comm.totalWeight)) {
            return false;
        }

        // Aggregate pubkeys and verify BLS signature over the warp preimage.
        bytes memory aggregateKey = att.signers[0].blsPublicKey;
        for (uint256 i = 1; i < numSigners;) {
            aggregateKey = BLST.addG1(aggregateKey, att.signers[i].blsPublicKey);
            unchecked {
                ++i;
            }
        }
        return BLST.verifySignature(aggregateKey, att.aggregateBlsSig, signedData);
    }

    /*
    * @dev Serializes a ValidatorSetMerkleAttestation in the format expected by parseMerkleAttestation.
    */
    function serializeMerkleAttestation(
        ValidatorSetMerkleAttestation memory att
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(CODEC, packValidatorSetMerkleAttestation(att));
    }

    /**
     * @notice Deserialize a bit set back into an array of bools.
     * @dev Inverse of packProofFlags. Reads a uint32 count then ceil(count/8) packed bytes.
     * @return consumed Number of bytes read from data.
     * @return flags The decoded bool array.
     */
    function unpackProofFlags(
        bytes calldata data
    ) internal pure returns (uint256, bool[] memory) {
        uint32 numFlags = uint32(bytes4(data[0:4]));
        uint256 flagBytesLen = Math.ceilDiv(numFlags, 8);
        bool[] memory flags = new bool[](numFlags);
        for (uint256 i = 0; i < numFlags;) {
            uint8 byteVal = uint8(data[4 + (i >> 3)]);
            flags[i] = (byteVal >> uint8(i & 7)) & 1 == 1;
            unchecked {
                ++i;
            }
        }
        return (4 + flagBytesLen, flags);
    }

    /**
     * @notice serialize an array of bools to a bit set
     */
    function packProofFlags(
        bool[] memory flags
    ) internal pure returns (bytes memory) {
        uint32 numFlags = uint32(flags.length);
        uint256 flagBytesLen = Math.ceilDiv(numFlags, 8);
        bytes memory packedFlags = new bytes(flagBytesLen);
        for (uint256 i = 0; i < numFlags;) {
            if (flags[i]) {
                packedFlags[i >> 3] |= bytes1(uint8(1) << uint8(i & 7));
            }
            unchecked {
                ++i;
            }
        }
        return abi.encodePacked(numFlags, packedFlags);
    }

    /**
     * @notice Serializes a ValidatorSetMerkleCommitment payload.
     */
    function serializeMerkleCommitment(
        ValidatorSetMerkleCommitment memory commitment
    ) internal pure returns (bytes memory) {
        bytes4 payloadType = bytes4(0x00000006);
        return abi.encodePacked(CODEC, payloadType, packValidatorSetMerkleCommitment(commitment));
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
        uint256 consumed;
        (consumed, att) = unpackValidatorSetMerkleAttestation(data[2:]);
        att.aggregateBlsSig = data[2 + consumed:2 + consumed + BLST.BLS_SIGNATURE_LENGTH];
        require(
            2 + consumed + BLST.BLS_SIGNATURE_LENGTH == data.length, "Trailing bytes in attestation"
        );
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
        (, ValidatorSetMerkleCommitment memory result) =
            unpackValidatorSetMerkleCommitment(data[6:]);
        return result;
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

    /*
     * @notice Prepends a one to the hash of a validator to avoid second pre-image attacks on
     * the merkle trees. This is a convenience function that should be used to hash validators
     * rather than calling keccak256 on them directly
     */
    function hashValidator(
        Validator memory val
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(1), val.blsPublicKey, val.weight));
    }

    function keccakInternalPair(
        bytes32 a,
        bytes32 b
    ) internal pure returns (bytes32) {
        return _keccakPair(a, b);
    }

    function _keccakPair(
        bytes32 a,
        bytes32 b
    ) internal pure returns (bytes32) {
        return a < b
            ? keccak256(abi.encodePacked(uint256(0), a, b))
            : keccak256(abi.encodePacked(uint256(0), b, a));
    }
}
