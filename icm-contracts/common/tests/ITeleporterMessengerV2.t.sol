// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {
    TeleporterICMMessage, TeleporterMessageV2, ICMTeleporterV2
} from "../TeleporterMessageV2.sol";
import {TeleporterMessageReceipt} from "@teleporter/ITeleporterMessenger.sol";

contract ICMTest is Test {
    function testTeleporterMessageV2RoundTrip(
        bytes memory payload,
        uint8 numRelayerAddresses,
        address relayerAddress,
        uint8 numReceipts,
        uint256 messageNonce
    ) public pure {
        vm.assume(numRelayerAddresses < 10);
        vm.assume(numReceipts < 10);
        address[] memory allowedRelayerAddresses = new address[](numRelayerAddresses);
        for (uint256 i = 0; i < numRelayerAddresses; i++) {
            allowedRelayerAddresses[i] = relayerAddress;
        }

        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](numReceipts);
        for (uint256 i = 0; i < numReceipts; i++) {
            receipts[i] = TeleporterMessageReceipt({
                receivedMessageNonce: messageNonce,
                relayerRewardAddress: relayerAddress
            });
        }

        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 0,
            originSenderAddress: address(123),
            originTeleporterAddress: address(456),
            destinationBlockchainID: bytes32(hex"abcd"),
            destinationAddress: address(987),
            requiredGasLimit: 100000,
            allowedRelayerAddresses: allowedRelayerAddresses,
            receipts: receipts,
            message: payload
        });
        bytes memory serialized = ICMTeleporterV2.serializeTeleporterMessageV2(teleporterMessage);
        TeleporterMessageV2 memory deserialized =
            ICMTeleporterV2.parseTeleporterMessageV2(serialized);

        assertEq(deserialized.messageNonce, teleporterMessage.messageNonce);
        assertEq(deserialized.originSenderAddress, teleporterMessage.originSenderAddress);
        assertEq(deserialized.originTeleporterAddress, teleporterMessage.originTeleporterAddress);
        assertEq(deserialized.destinationBlockchainID, teleporterMessage.destinationBlockchainID);
        assertEq(deserialized.destinationAddress, teleporterMessage.destinationAddress);
        assertEq(deserialized.requiredGasLimit, teleporterMessage.requiredGasLimit);
        assertEq(deserialized.allowedRelayerAddresses, teleporterMessage.allowedRelayerAddresses);
        for (uint256 i = 0; i < allowedRelayerAddresses.length; i++) {
            assertEq(
                teleporterMessage.allowedRelayerAddresses[i],
                deserialized.allowedRelayerAddresses[i]
            );
        }

        assertEq(deserialized.message, teleporterMessage.message);
    }

    /*
     * @dev Test to make sure a round trip of serialization is a no-op
     */
    function testRoundTripTeleporterICMMessage(
        uint32 sourceNetworkID,
        bytes32 sourceBlockchainID,
        bytes memory payload,
        uint8 numRelayerAddresses,
        address relayerAddress,
        uint8 numReceipts,
        uint256 messageNonce
    ) public pure {
        vm.assume(numRelayerAddresses < 10);
        vm.assume(numReceipts < 10);
        address[] memory allowedRelayerAddresses = new address[](numRelayerAddresses);
        for (uint256 i = 0; i < numRelayerAddresses; i++) {
            allowedRelayerAddresses[i] = relayerAddress;
        }

        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](numReceipts);
        for (uint256 i = 0; i < numReceipts; i++) {
            receipts[i] = TeleporterMessageReceipt({
                receivedMessageNonce: messageNonce,
                relayerRewardAddress: relayerAddress
            });
        }

        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 0,
            originSenderAddress: address(123),
            originTeleporterAddress: address(456),
            destinationBlockchainID: bytes32(hex"abcd"),
            destinationAddress: address(987),
            requiredGasLimit: 100000,
            allowedRelayerAddresses: allowedRelayerAddresses,
            receipts: receipts,
            message: payload
        });

        TeleporterICMMessage memory icmMessage = TeleporterICMMessage({
            message: teleporterMessage,
            sourceNetworkID: sourceNetworkID,
            sourceBlockchainID: sourceBlockchainID,
            attestation: abi.encode(1)
        });

        bytes memory serialized = ICMTeleporterV2.serializeTeleporterICMMessage(icmMessage);
        TeleporterICMMessage memory deserializedICM =
            ICMTeleporterV2.parseTeleporterICMMessage(serialized);
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
