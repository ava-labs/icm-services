// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ICM, ICMMessage, ICMRawMessage} from "../ICM.sol";

contract ICMTest is Test {

    /*
     * @dev Fuzz test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripRawMessage(
        bytes4 sourceNetworkID,
        bytes32 sourceBlockchainID,
        address sourceAddress,
        address verifierAddress,
        bytes memory payload
    ) public pure {
        ICMRawMessage memory raw = ICMRawMessage({
            sourceNetworkID: uint32(sourceNetworkID),
            sourceBlockchainID: sourceBlockchainID,
            sourceAddress: sourceAddress,
            verifierAddress: verifierAddress,
            payload: payload
        });

        bytes memory serialized = ICM.serializeICMRawMessage(raw);
        ICMRawMessage memory deserialized = ICM.parseICMRawMessage(serialized);
        assertEq(raw.sourceNetworkID, deserialized.sourceNetworkID);
        assertEq(raw.sourceBlockchainID, deserialized.sourceBlockchainID);
        assertEq(raw.sourceAddress, deserialized.sourceAddress);
        assertEq(raw.verifierAddress, deserialized.verifierAddress);
        assertEq(raw.payload, deserialized.payload);
    }

    function testRoundTripICMessage(
        bytes4 sourceNetworkID,
        bytes32 sourceBlockchainID,
        address sourceAddress,
        address verifierAddress,
        bytes memory payload,
        bytes memory attestation
    ) public pure {
        ICMRawMessage memory raw = ICMRawMessage({
            sourceNetworkID: uint32(sourceNetworkID),
            sourceBlockchainID: sourceBlockchainID,
            sourceAddress: sourceAddress,
            verifierAddress: verifierAddress,
            payload: payload
        });

        bytes memory rawMessageBytes = ICM.serializeICMRawMessage(raw);

        ICMMessage memory message= ICMMessage({
            message: raw,
            rawMessageBytes: rawMessageBytes,
            attestation: attestation
        });
        bytes memory serialized = ICM.serializeICMMessage(message);
        ICMMessage memory deserialized = ICM.parseICMMessage(serialized);

        assertEq(raw.sourceNetworkID, deserialized.message.sourceNetworkID);
        assertEq(raw.sourceBlockchainID, deserialized.message.sourceBlockchainID);
        assertEq(raw.sourceAddress, deserialized.message.sourceAddress);
        assertEq(raw.verifierAddress, deserialized.message.verifierAddress);
        assertEq(raw.payload, deserialized.message.payload);
        assertEq(rawMessageBytes, deserialized.rawMessageBytes);
        assertEq(attestation, deserialized.attestation);
    }
}