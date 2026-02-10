// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ICM, ICMMessage} from "../ICM.sol";

contract ICMTest is Test {
    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripICMMessage(
        bytes4 sourceNetworkID,
        bytes32 sourceBlockchainID,
        bytes memory payload,
        bytes memory attestation
    ) public pure {
        ICMMessage memory message = ICMMessage({
            payload: payload,
            sourceNetworkID: uint32(sourceNetworkID),
            sourceBlockchainID: sourceBlockchainID,
            attestation: attestation
        });

        bytes memory serialized = ICM.serializeICMMessage(message);
        ICMMessage memory deserialized = ICM.parseICMMessage(serialized);

        assertEq(uint32(sourceNetworkID), deserialized.sourceNetworkID);
        assertEq(sourceBlockchainID, deserialized.sourceBlockchainID);
        assertEq(rawMessage, deserialized.rawMessage);
        assertEq(attestation, deserialized.attestation);
    }
}
