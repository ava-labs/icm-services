// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {IOracleMessageReceiver} from "../teleporterV2/IOracleMessageReceiver.sol";

/**
 * @dev Test receiver for oracle E2E tests. Records the last delivered oracle message
 *      so tests can assert on payload correctness.
 *
 *      NOT FOR PRODUCTION USE.
 */
contract MockOracleReceiver is IOracleMessageReceiver {
    address public immutable oracleAdapter;

    bytes32 public lastSourceChainID;
    string public lastSourceType;
    string public lastSourceAddress;
    uint64 public lastNonce;
    bytes public lastPayload;
    uint256 public receiveCount;

    error OnlyOracleAdapter(address caller, address expected);

    constructor(
        address oracleAdapter_
    ) {
        oracleAdapter = oracleAdapter_;
    }

    /**
     * @inheritdoc IOracleMessageReceiver
     */
    function receiveOracleMessage(
        bytes32 sourceChainID,
        string calldata sourceType,
        string calldata sourceAddress,
        uint64 nonce,
        bytes calldata payload
    ) external override {
        if (msg.sender != oracleAdapter) {
            revert OnlyOracleAdapter(msg.sender, oracleAdapter);
        }

        lastSourceChainID = sourceChainID;
        lastSourceType = sourceType;
        lastSourceAddress = sourceAddress;
        lastNonce = nonce;
        lastPayload = payload;
        receiveCount++;
    }
}
