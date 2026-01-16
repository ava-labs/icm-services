// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/**
 * THIS IS LIBRARY IS UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
struct ICMRawMessage {
    // used to distinguish between mainnet and testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // The address that sent the message
    address sourceAddress;
    // Address of verifying contract on the receiving blockchain
    address verifierAddress;
    // The message payload
    bytes payload;
}

struct ICMMessage {
    ICMRawMessage message;
    // The serialized bytes of `message`
    bytes rawMessageBytes;
    // Arbitrary bytes that is used by receiving contracts to
    // authenticate this message.
    bytes attestation;
}

/**
 * @title ICM
 * @notice Utility library for Interchain Messaging (ICM) messages. Mainly (de)serialization.
 * @dev This library provides helper functions for working with ICM signatures and validation
 */
library ICM {
    /* solhint-disable no-inline-assembly */

    uint256 private constant SOURCE_NETWORK_ID_LENGTH = 4;
    uint256 private constant MESSAGE_METADATA_LENGTH = 82;

    /*
     * @notice Deserialize an ICM message from bytes
     */
    function parseICMMessage(
        bytes calldata data
    ) public pure returns (ICMMessage memory) {
        ICMRawMessage memory message = parseICMRawMessage(data);
        uint256 payloadLength = message.payload.length;
        // Parse the unsigned message bytes
        bytes memory rawMessageBytes = data[0:MESSAGE_METADATA_LENGTH + payloadLength];
        // the rest of the bytes are the attestation
        bytes memory attestation = data[MESSAGE_METADATA_LENGTH + payloadLength:];
        return ICMMessage({
            message: message,
            rawMessageBytes: rawMessageBytes,
            attestation: attestation
        });
    }

    /*
     * @notice Deserialize a raw ICM message from bytes
     */
    function parseICMRawMessage(
        bytes calldata data
    ) public pure returns (ICMRawMessage memory) {
        // Validate the codec ID is 0
        require(data[0] == 0 && data[1] == 0, "Invalid codec ID");

        // Parse the sourceNetworkID
        uint32 sourceNetworkID = uint32(bytes4(data[2:2 + SOURCE_NETWORK_ID_LENGTH]));

        // Parse the sourceBlockchainID
        bytes32 sourceBlockchainID = bytes32(data[6:38]);

        // Parse the sourceAddress and verifierAddress
        address sourceAddress = address(bytes20(data[38:58]));
        address verifierAddress = address(bytes20(data[58:78]));

        // Parse the payload
        uint32 payloadLength = uint32(bytes4(data[78:MESSAGE_METADATA_LENGTH]));
        bytes memory payload = data[MESSAGE_METADATA_LENGTH:MESSAGE_METADATA_LENGTH + payloadLength];
        return ICMRawMessage({
            sourceNetworkID: sourceNetworkID,
            sourceBlockchainID: sourceBlockchainID,
            sourceAddress: sourceAddress,
            verifierAddress: verifierAddress,
            payload: payload
        });
    }

    /**
     * @notice Serialize a raw ICM message to bytes.
     */
    function serializeICMRawMessage(
        ICMRawMessage calldata message
    ) public pure returns (bytes memory) {
        // encode the payload length
        bytes2 codec = bytes2(0);
        uint32 sourceNetworkID = uint32(message.sourceNetworkID);
        uint32 payloadLength = uint32(message.payload.length);
        return abi.encodePacked(
            codec,
            sourceNetworkID,
            message.sourceBlockchainID,
            message.sourceAddress,
            message.verifierAddress,
            payloadLength,
            message.payload
        );
    }

    /**
     * @notice Serialize an ICM message to bytes
     */
    function serializeICMMessage(
        ICMMessage calldata message
    ) public pure returns (bytes memory) {
        return abi.encodePacked(message.rawMessageBytes, message.attestation);
    }
}
