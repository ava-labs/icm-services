// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

struct ICMRawMessage {
    // used to distinguish between mainnet and testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // The address that sent the message
    address sourceAddress;
    // Address of verifying contract on the receiving blockchain
    address verifierAddress;
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

    /*
     * @notice Deserialize an ICM message from bytes
     */
    function parseICMMessage(
        bytes calldata data
    ) public pure returns (ICMMessage memory) {
        ICMRawMessage memory message = parseICMRawMessage(data);
        uint256 payloadLength = message.payload.length;
        // Parse the unsigned message bytes
        bytes memory rawMessageBytes = data[0: 82 + payloadLength];
        // the rest of the bytes are the attestation
        bytes memory attestation = data[82 + payloadLength:];
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
        uint32 sourceNetworkID = uint32(bytes4(data[2:6]));


        // Parse the sourceBlockchainID
        bytes32 sourceBlockchainID = bytes32(data[6:38]);

        // Parse the sourceAddress and verifierAddress
        address sourceAddress = address(bytes20(data[38:58]));
        address verifierAddress = address(bytes20(data[58:78]));

        // Parse the payload
        uint32 payloadLength = uint32(bytes4(data[78:82]));
        bytes memory payload = data[82:82 + payloadLength];
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
        bytes memory serialized = new bytes(
        // the codec
            2
            // the sourceNetworkID
            + 4
            // the sourceBlockchainID
            + 32
            // the sourceAddress
            + 20
            // the verifierAddress
            + 20
            // the payload length
            + 4
            // the message payload
            + message.payload.length
        );
        // encode the payload length
        uint256 payloadLength = message.payload.length;
        bytes4 payloadLengthBytes = bytes4(uint32(message.payload.length));
        assembly ("memory-safe"){
            // set start after the array size
            let s := add(serialized, 32)
            // Store Codec
            mstore(s, 0x0000)
            // store the sourceNetworkID
            // the shift is to transform a 32 byte value to a 4 byte value
            mstore(add(s, 2), shl(224, calldataload(message)))
            // store the sourceBlockchainID
            mstore(add(s, 6), calldataload(add(message, 0x20)))
            // store the sourceAddress, shifting 12 bytes to get the 20 byte address
            mstore(add(s, 38), shl(96, calldataload(add(message, 0x40))))
            // store the verifier address, shifting 12 bytes to get the 20 byte address
            mstore(add(s, 58),  shl(96, calldataload(add(message, 0x60))))
            // store the payload length
            mstore(add(s, 78), payloadLengthBytes)
            // store the payload
            let payloadOffset := calldataload(add(message, 0x80))
            let payload := add(add(message, payloadOffset), 0x20)
            calldatacopy(add(s, 82), payload, payloadLength)
        }

        return serialized;
    }

    /**
     * @notice Serialize an ICM message to bytes
     */
    function serializeICMMessage(
        ICMMessage calldata message
    ) public pure returns (bytes memory) {
        bytes memory data = new bytes(
            // unsigned message length
            message.rawMessageBytes.length
            // the length of the attestation
            + message.attestation.length
        );
        uint256 rawMessageLength = message.rawMessageBytes.length;
        uint256 attestationLength = message.attestation.length;
        assembly ("memory-safe") {
            // store the raw message bytes
            let rawMessageOffset := calldataload(add(message, 0x20))
            let rawMessage := add(add(message, rawMessageOffset), 0x20)
            calldatacopy(add(data, 0x20), rawMessage, rawMessageLength)
            // store the attestation bytes
            let attestationOffset := calldataload(add(message, 0x40))
            let attestation := add(add(message, attestationOffset), 0x20)
            calldatacopy(add(data, add(0x20, rawMessageLength)), attestation, attestationLength)
        }
        return data;
    }

}
