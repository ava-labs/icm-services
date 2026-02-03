// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {ICM, ICMMessage} from "./ICM.sol";
import {TeleporterMessageReceipt} from "@teleporter/ITeleporterMessenger.sol";

struct TeleporterMessageV2 {
    uint256 messageNonce;
    address originSenderAddress;
    // This is needed because we have an extra level of abstraction now. Before TeleporterMessenger was calling
    // WarpMessenger directly, so the `senderAddress` of the warp message was the address of the TeleporterMessenger contract.
    // Now, WarpAdapter is calling WarpMessenger, so the `originSenderAddress` of the warp message is the address of the WarpAdapter contract.
    address originTeleporterAddress;
    bytes32 destinationBlockchainID;
    address destinationAddress;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    TeleporterMessageReceipt[] receipts;
    bytes message;
}

// A domain type for the specific kind of ICM message TeleporterV2 expects
struct TeleporterICMMessage {
    TeleporterMessageV2 message;
    // used to distinguish between mainnet and testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // Arbitrary bytes that is used by receiving contracts to
    // authenticate this message.
    bytes attestation;
}

interface IMessageVerifier {
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool);
}

// This function signature can be changed to accept bytes to make it more generic, but I think
// having the TeleporterMessage struct is more clear for now.
interface IMessageSender {
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external;
}

// solhint-disable-next-line no-empty-blocks
interface IAdapter is IMessageSender, IMessageVerifier {}

/**
 * @title ICMTeleportV2
 * @notice Utility library for making TeleporterV2 work with ICM messages. Mainly (de)serialization
 * @dev This structuring of the code is a filthy hack. In order for
 * `TeleporterICMMessageParsing.parseTeleporterICMMessage` to call `TeleporterMessageV2Parsing.parseTeleporterMessageV2`
 * and have the memory bytes passed in as calldata, the functions cannot belong to the same library. To hide this
 * downstream, this library consolidates all the calls in one pl0ace.
 */
library ICMTeleporterV2 {
    function parseTeleporterMessageV2(
        bytes calldata data
    ) public pure returns (TeleporterMessageV2 memory) {
        return TeleporterMessageV2Parsing.parseTeleporterMessageV2(data);
    }

    function serializeTeleporterMessageV2(
        TeleporterMessageV2 memory message
    ) public pure returns (bytes memory) {
        return TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message);
    }

    function parseTeleporterICMMessage(
        bytes calldata message
    ) public pure returns (TeleporterICMMessage memory) {
        return TeleporterICMMessageParsing.parseTeleporterICMMessage(message);
    }

    function serializeTeleporterICMMessage(
        TeleporterICMMessage memory message
    ) public pure returns (bytes memory) {
        return TeleporterICMMessageParsing.serializeTeleporterICMMessage(message);
    }
}

library TeleporterMessageV2Parsing {
    /**
     * @notice Parse a serialized `TeleporterMessageV2` instance
     */
    function parseTeleporterMessageV2(
        bytes calldata data
    ) public pure returns (TeleporterMessageV2 memory teleporterMessage) {
        teleporterMessage.messageNonce = uint256(bytes32(data[0:32]));
        teleporterMessage.originSenderAddress = address(bytes20(data[32:52]));
        teleporterMessage.originTeleporterAddress = address(bytes20(data[52:72]));
        teleporterMessage.destinationBlockchainID = bytes32(data[72:104]);
        teleporterMessage.destinationAddress = address(bytes20(data[104:124]));
        teleporterMessage.requiredGasLimit = uint256(bytes32(data[124:156]));

        {
            // get the number of addressed in the array
            uint32 numRelayerAddresses = uint32(bytes4(data[156:160]));
            uint256 offsetAddresses = 160;
            address[] memory allowedRelayerAddresses = new address[](numRelayerAddresses);
            // parse the addresses
            for (uint256 i = 0; i < numRelayerAddresses; i++) {
                allowedRelayerAddresses[i] =
                    address(bytes20(data[offsetAddresses:offsetAddresses + 20]));
                offsetAddresses += 20;
            }
            teleporterMessage.allowedRelayerAddresses = allowedRelayerAddresses;
        }

        {
            uint256 offsetReceipts = 160 + 20 * teleporterMessage.allowedRelayerAddresses.length;
            // get the number of receipts
            uint32 numReceipts = uint32(bytes4(data[offsetReceipts:offsetReceipts + 4]));
            TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](numReceipts);
            offsetReceipts += 4;
            // parse the receipts
            for (uint256 i = 0; i < teleporterMessage.receipts.length; i++) {
                receipts[i] = TeleporterMessageReceipt({
                    receivedMessageNonce: uint256(bytes32(data[offsetReceipts:offsetReceipts + 32])),
                    relayerRewardAddress: address(
                        bytes20(data[offsetReceipts + 32:offsetReceipts + 52])
                    )
                });
                offsetReceipts += 52;
            }
            teleporterMessage.receipts = receipts;
        }
        uint256 offset = 164 + (52 * teleporterMessage.receipts.length)
            + (20 * teleporterMessage.allowedRelayerAddresses.length);
        // the remaining data is the inner message
        teleporterMessage.message = bytes(data[offset:]);
        return teleporterMessage;
    }

    /**
     * @notice Serialize a `TeleporterMessageV2` instance
     */
    function serializeTeleporterMessageV2(
        TeleporterMessageV2 memory message
    ) public pure returns (bytes memory) {
        return abi.encodePacked(
            message.messageNonce,
            message.originSenderAddress,
            message.originTeleporterAddress,
            message.destinationBlockchainID,
            message.destinationAddress,
            message.requiredGasLimit,
            uint32(message.allowedRelayerAddresses.length),
            serializeAddresses(message.allowedRelayerAddresses),
            uint32(message.receipts.length),
            serializeTeleporterReceipts(message.receipts),
            message.message
        );
    }

    /**
     * @notice serialize an array of `address` instances
     */
    function serializeAddresses(
        address[] memory allowedRelayerAddresses
    ) internal pure returns (bytes memory) {
        bytes memory serialized = new bytes(allowedRelayerAddresses.length * 20);
        /* solhint-disable no-inline-assembly */
        assembly ("memory-safe") {
            let s := add(serialized, 32)
            let r := add(allowedRelayerAddresses, 44)
            let l := mload(allowedRelayerAddresses)
            for { let i := 0 } lt(i, l) { i := add(i, 1) } {
                mstore(s, mload(r))
                s := add(s, 20)
                r := add(r, 32)
            }
        }
        /* solhint-enable no-inline-assembly */
        return serialized;
    }

    /**
     * @notice serialize an array of `TeleporterMessageReceipt` instances
     */
    function serializeTeleporterReceipts(
        TeleporterMessageReceipt[] memory receipts
    ) internal pure returns (bytes memory) {
        bytes memory serialized = new bytes(receipts.length * 52);
        uint256 length = receipts.length;
        /* solhint-disable no-inline-assembly */
        assembly ("memory-safe") {
            let s := add(serialized, 0x20)
            let r := add(receipts, 0x20)
            for { let i := 0 } lt(i, length) { i := add(i, 1) } {
                mstore(s, mload(r))
                s := add(s, 32)
                r := add(r, 32)
                mstore(s, mload(r))
                s := add(s, 20)
                r := add(r, 32)
            }
        }
        /* solhint-enable no-inline-assembly */
        return serialized;
    }
}

library TeleporterICMMessageParsing {
    /**
     * @notice parse a serialized `TeleporterICMMessage` instance
     */
    function parseTeleporterICMMessage(
        bytes calldata message
    ) public pure returns (TeleporterICMMessage memory) {
        ICMMessage memory icmMessage = ICM.parseICMMessage(message);
        return TeleporterICMMessage({
            message: ICMTeleporterV2.parseTeleporterMessageV2(icmMessage.rawMessage),
            sourceNetworkID: icmMessage.sourceNetworkID,
            sourceBlockchainID: icmMessage.sourceBlockchainID,
            attestation: icmMessage.attestation
        });
    }

    /**
     * @notice Serialize a `TeleporterICMMessage` instance
     */
    function serializeTeleporterICMMessage(
        TeleporterICMMessage memory message
    ) public pure returns (bytes memory) {
        ICMMessage memory icmMessage = ICMMessage({
            rawMessage: ICMTeleporterV2.serializeTeleporterMessageV2(message.message),
            sourceNetworkID: message.sourceNetworkID,
            sourceBlockchainID: message.sourceBlockchainID,
            attestation: message.attestation
        });
        return ICM.serializeICMMessage(icmMessage);
    }
}
