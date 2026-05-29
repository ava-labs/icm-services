// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.30;

import {ICM, ICMMessage} from "./ICM.sol";
import {TeleporterMessageReceipt} from "@teleporter/ITeleporterMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// #[pack(contract="ICMTeleporterV2")]
// #[unpack(contract="ICMTeleporterV2", calldata)]
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
    // #[pack(method="serializeTeleporterReceipts")]
    // #[unpack(method="parseTeleporterReceipts")]
    TeleporterMessageReceipt[] receipts;
    bytes message;
}

// A domain type for the specific kind of ICM message TeleporterV2 expects
// #[pack(contract="ICMTeleporterV2")]
// #[unpack(contract="ICMTeleporterV2", calldata)]
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

/**
 * @title ICMTeleportV2
 * @notice Utility library for making TeleporterV2 work with ICM messages. Mainly (de)serialization
 */
library ICMTeleporterV2 {

}


/**
 * @notice serialize an array of `TeleporterMessageReceipt` instances
 */
function serializeTeleporterReceipts(
        TeleporterMessageReceipt[] memory receipts
    ) pure returns (bytes memory) {
        bytes memory serialized = new bytes(receipts.length * 52);
        uint256 length = receipts.length;

        /* solhint-disable no-inline-assembly */
        assembly ("memory-safe") {
            let s := add(serialized, 0x20)
            let r := add(receipts, 0x20)
            for { let i := 0 } lt(i, length) { i := add(i, 1) } {
                // get the value at index number i
                let value := mload(r)
                mstore(s, mload(value))
                s := add(s, 32)
                // shift the address left 12 bytes
                mstore(s, shl(96, mload(add(value, 0x20))))
                s := add(s, 20)
                r := add(r, 32)
            }
        }
        /* solhint-enable no-inline-assembly */
        return abi.encodePacked(uint32(receipts.length), serialized);
    }

/**
 * @notice Parse a serialized array  of `TeleporterMessageReceipt` instances
 */
function parseTeleporterReceipts(bytes calldata data) pure returns (uint256, TeleporterMessageReceipt[] memory) {
        // get the number of receipts
        uint32 numReceipts = uint32(bytes4(data[:4]));
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](numReceipts);
        uint256 offsetReceipts = 4;
        // parse the receipts
        for (uint256 i = 0; i < numReceipts;) {
            receipts[i] = TeleporterMessageReceipt({
                receivedMessageNonce: uint256(
                    bytes32(data[offsetReceipts:offsetReceipts + 32])
                ),
                relayerRewardAddress: address(
                    bytes20(data[offsetReceipts + 32:offsetReceipts + 52])
                )
            });
            offsetReceipts += 52;
            unchecked { i++; }
        }
        return (offsetReceipts, receipts);
    }

