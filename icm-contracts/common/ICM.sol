// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/**
 * THIS IS LIBRARY IS UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
struct ICMMessage {
    // The serialized bytes of the raw message. The data and serializations formats
    // for this data will be app / contract specific
    bytes rawMessage;
    // used to distinguish between mainnet and testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // Arbitrary bytes that is used by receiving contracts to
    // authenticate this message.
    bytes attestation;
}

/**
 * @title ICM
 * @notice Utility library for Interchain Messaging (ICM) messages. Mainly (de)serialization.
 */
library ICM {
    // The number of bytes encoding the length of a raw ICM message
    uint256 private constant MESSAGE_LENGTH_BYTES = 4;

    /*
     * @notice Deserialize an ICM message from bytes
     */
    function parseICMMessage(
        bytes calldata data
    ) public pure returns (ICMMessage memory) {
        bytes memory message = extractICMRawMessage(data);
        // The position in data after the raw message bytes
        uint256 messageOffset = message.length + MESSAGE_LENGTH_BYTES;
        // parse the source Network ID
        uint32 sourceNetworkID = uint32(bytes4(data[messageOffset:messageOffset + 4]));
        // parse the source blockchain ID
        bytes32 sourceBlockchainID = bytes32(data[messageOffset + 4:messageOffset + 36]);

        // the rest of the bytes are the attestation
        bytes memory attestation = data[messageOffset + 36:];
        return ICMMessage({
            rawMessage: message,
            sourceNetworkID: sourceNetworkID,
            sourceBlockchainID: sourceBlockchainID,
            attestation: attestation
        });
    }

    /*
     * @notice Extract the raw ICM message bytes
     */
    function extractICMRawMessage(
        bytes calldata data
    ) public pure returns (bytes memory) {
        uint32 messageLength = uint32(bytes4(data[0:MESSAGE_LENGTH_BYTES]));
        bytes memory rawMessage = data[MESSAGE_LENGTH_BYTES:MESSAGE_LENGTH_BYTES + messageLength];
        return rawMessage;
    }

    /**
     * @notice Serialize an ICM message to bytes
     */
    function serializeICMMessage(
        ICMMessage calldata message
    ) public pure returns (bytes memory) {
        uint32 messageLength = uint32(message.rawMessage.length);
        return abi.encodePacked(
            messageLength,
            message.rawMessage,
            message.sourceNetworkID,
            message.sourceBlockchainID,
            message.attestation
        );
    }
}
