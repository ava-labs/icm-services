// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";

/**
 * @dev Test receiver for oracle E2E tests. Records the last delivered oracle message
 * so tests can assert on payload correctness.
 * Implements ITeleporterReceiver; decodes oracle fields from the message bytes.
 * message is abi.encode(sourceType, sourceAddress, sourceBlockHeight, nonce, payload).
 * NOT FOR PRODUCTION USE.
 */
contract MockOracleReceiver is ITeleporterReceiver {
    address public immutable teleporterMessenger;

    bytes32 public lastSourceChainID;
    string public lastSourceType;
    string public lastSourceAddress;
    uint256 public lastSourceBlockHeight;
    uint256 public lastNonce;
    bytes public lastPayload;
    uint256 public receiveCount;

    event OracleMessageReceived(
        bytes32 indexed sourceBlockchainID,
        string sourceType,
        string sourceAddress,
        uint256 sourceBlockHeight,
        uint256 nonce,
        bytes payload
    );

    error UnauthorizedSender(address caller, address expected);
    error ZeroAddress();

    constructor(
        address teleporterMessenger_
    ) {
        if (teleporterMessenger_ == address(0)) revert ZeroAddress();
        teleporterMessenger = teleporterMessenger_;
    }

    /**
     * @inheritdoc ITeleporterReceiver
     */
    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address, // originSenderAddress is address(0) for oracle messages
        bytes calldata message
    ) external override {
        if (msg.sender != teleporterMessenger) {
            revert UnauthorizedSender(msg.sender, teleporterMessenger);
        }

        (
            string memory sourceType,
            string memory sourceAddress,
            uint256 sourceBlockHeight,
            uint256 nonce,
            bytes memory oraclePayload
        ) = abi.decode(message, (string, string, uint256, uint256, bytes));

        lastSourceChainID = sourceBlockchainID;
        lastSourceType = sourceType;
        lastSourceAddress = sourceAddress;
        lastSourceBlockHeight = sourceBlockHeight;
        lastNonce = nonce;
        lastPayload = oraclePayload;
        ++receiveCount;
        emit OracleMessageReceived({
            sourceBlockchainID: sourceBlockchainID,
            sourceType: sourceType,
            sourceAddress: sourceAddress,
            sourceBlockHeight: sourceBlockHeight,
            nonce: nonce,
            payload: oraclePayload
        });
    }
}
