// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {TeleporterICMMessage, TeleporterMessageV2} from "./TeleporterMessageV2.sol";

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
