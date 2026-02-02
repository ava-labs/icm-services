// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {TeleporterMessageV2, ICMMessage} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessageReceipt} from "@teleporter/ITeleporterMessenger.sol";

contract ICMTest is Test {
    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripRawMessage(
        bytes4 sourceNetworkID,
        bytes32 sourceBlockchainID,
        bytes memory payload
    ) public pure {
        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 0,
            originSenderAddress: address(123),
            originTeleporterAddress: address(456),
            destinationBlockchainID: bytes32(hex"abcd"),
            destinationAddress: address(987),
            requiredGasLimit: 100000,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: payload
        });

        ICMMessage memory icmMessage = ICMMessage({
            message: teleporterMessage,
            sourceNetworkID: uint32(sourceNetworkID),
            sourceBlockchainID: sourceBlockchainID,
            attestation: abi.encode(1)
        });

        bytes memory serialized = abi.encode(icmMessage);
        ICMMessage memory deserializedICM = abi.decode(serialized, (ICMMessage));
        assertEq(icmMessage.sourceNetworkID, deserializedICM.sourceNetworkID);
        assertEq(icmMessage.sourceBlockchainID, deserializedICM.sourceBlockchainID);
        assertEq(icmMessage.attestation, deserializedICM.attestation);

        TeleporterMessageV2 memory deserializedTeleporterMessage = deserializedICM.message;
        assertEq(teleporterMessage.messageNonce, deserializedTeleporterMessage.messageNonce);
        assertEq(
            teleporterMessage.originSenderAddress, deserializedTeleporterMessage.originSenderAddress
        );
        assertEq(
            teleporterMessage.originTeleporterAddress,
            deserializedTeleporterMessage.originTeleporterAddress
        );
        assertEq(
            teleporterMessage.destinationBlockchainID,
            deserializedTeleporterMessage.destinationBlockchainID
        );
        assertEq(
            teleporterMessage.destinationAddress, deserializedTeleporterMessage.destinationAddress
        );
        assertEq(teleporterMessage.requiredGasLimit, deserializedTeleporterMessage.requiredGasLimit);
        assertEq(
            teleporterMessage.allowedRelayerAddresses.length,
            deserializedTeleporterMessage.allowedRelayerAddresses.length
        );
        assertEq(teleporterMessage.receipts.length, deserializedTeleporterMessage.receipts.length);
        assertEq(
            keccak256(teleporterMessage.message), keccak256(deserializedTeleporterMessage.message)
        );
    }
}
