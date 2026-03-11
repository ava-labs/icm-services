// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

/**
 * THIS IS AN EXAMPLE LIBRARY THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";

library RLPUtils {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;

    /**
     * @notice Struct representing an EVM event log.
     */
    struct EVMLog {
        address loggerAddress;
        bytes32[] topics;
        bytes data;
    }

    /**
     * @notice Struct representing an EVM transaction receipt.
     */
    struct EVMReceipt {
        uint8 txType;
        bytes postStateOrStatus;
        uint64 cumulativeGasUsed;
        bytes bloom;
        EVMLog[] logs;
    }

    /**
     * @notice Struct representing an EVM event log and its associated metadata.
     */
    struct EVMEventInfo {
        bytes32 blockchainID;
        uint256 blockNumber;
        uint256 txIndex;
        uint256 logIndex;
        EVMLog log;
    }

    /**
     * @notice Decodes a Receipt from raw RLP bytes, handling EIP-2718 typed envelopes.
     */
    function decodeReceipt(
        RLPReader.RLPItem memory encodedReceipt
    ) internal pure returns (EVMReceipt memory result) {
        // If the encoded receipt is not a list, the first byte is the transaction type,
        // followed by the RLP encoding of the receipt. If the encoded receipt is already a list itself,
        // then the transaction type is 0 (legacy tx).
        if (!encodedReceipt.isList()) {
            uint256 memptr = encodedReceipt.memPtr;
            uint8 txType;
            // solhint-disable-next-line no-inline-assembly
            assembly {
                txType := byte(0, mload(memptr))
            }
            require(txType >= 1 && txType <= 4, "Unsupported tx type");
            result.txType = txType;
            encodedReceipt = RLPReader.RLPItem({len: encodedReceipt.len - 1, memPtr: memptr + 1});
        }

        // Decode receipt
        // Receipt RLP structure: [postStateOrStatus, cumulativeGasUsed, bloom, logs[]]
        RLPReader.RLPItem[] memory items = encodedReceipt.toList();
        require(items.length == 4, "Invalid receipt: expected 4 fields");
        result.postStateOrStatus = items[0].toBytes();
        result.cumulativeGasUsed = uint64(items[1].toUint());
        result.bloom = items[2].toBytes();

        // Decode logs
        RLPReader.RLPItem[] memory logs = items[3].toList();
        uint256 logsLen = logs.length;
        result.logs = new EVMLog[](logsLen);
        for (uint256 i; i < logsLen;) {
            result.logs[i] = decodeLog(logs[i]);
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Decodes a single RLP-encoded log entry.
     */
    function decodeLog(
        RLPReader.RLPItem memory encodedLog
    ) internal pure returns (EVMLog memory result) {
        // Log RLP structure: [address, [topic0, topic1, ...], data]. Topics are capped at 4.
        RLPReader.RLPItem[] memory fields = encodedLog.toList();
        require(fields.length == 3, "Invalid log: expected 3 fields");
        result.loggerAddress = fields[0].toAddress();
        result.data = fields[2].toBytes();

        // Decode topics
        RLPReader.RLPItem[] memory topics = fields[1].toList();
        uint256 topicsLen = topics.length;
        require(topicsLen <= 4, "Invalid log: too many topics");
        result.topics = new bytes32[](topicsLen);
        for (uint256 i; i < topicsLen;) {
            result.topics[i] = bytes32(topics[i].toUint());
            unchecked {
                ++i;
            }
        }
    }
}
