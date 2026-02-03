// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

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

struct ICMMessage {
    // TODO we should serialize this as bytes, and we can also have different underlying message types for different
    // use cases other than Teleporter, but let's leave it like this for now for clarity.
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
        ICMMessage calldata message
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
