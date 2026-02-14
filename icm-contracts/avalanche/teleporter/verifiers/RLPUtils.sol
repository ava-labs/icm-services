// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

/**
 * THIS IS AN EXAMPLE INTERFACE THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

import { RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";

library RLPUtils {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;

    /*
    * @notice Struct representing an EVM event log.
    */
    struct EVMLog {
        address loggerAddress;
        bytes32[] topics;
        bytes data;
    }
    
    /*
    * @notice Struct representing an EVM transaction receipt.
    */
    struct EVMReceipt {
        uint8 txType;
        bytes postStateOrStatus;
        uint64 cumulativeGasUsed;
        bytes bloom;
        EVMLog[] logs;
    }

    /*
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
     * @notice Decodes a Receipt from raw RLP bytes, handling EIP-2718 types.
     * @dev Supports Legacy (Type 0), EIP-2930 (Type 1), EIP-1559 (Type 2), and EIP-4844 (Type 3).
     */
    function decodeReceipt(RLPReader.RLPItem memory encodedReceipt) internal pure returns (EVMReceipt memory) {
        uint8 txType = 0;

        // 1. Handle Transaction Types (EIP-2718)
        // If the RLP item is NOT a list, it's a typed receipt: [TypeByte] + [RLP_Receipt]
        if (!encodedReceipt.isList()) {
            uint256 memptr = encodedReceipt.memPtr;
            
            // Extract the first byte (Type Byte)
            // solhint-disable-next-line no-inline-assembly
            assembly {
                txType := byte(0, mload(memptr))
            }
            
            // PATCH: Added support for Type 3 (Blob Transactions)
            require(txType >= 1 && txType <= 3, "Unsupported tx type");

            // Advance the pointer to skip the Type Byte and read the actual RLP list
            encodedReceipt = RLPReader.RLPItem({
                len: encodedReceipt.len - 1, 
                memPtr: encodedReceipt.memPtr + 1
            });
        }

        // 2. Decode the Standard Receipt Fields
        // Structure: [status, cumulativeGas, bloom, logs]
        RLPReader.RLPItem[] memory receiptItems = encodedReceipt.toList();
        require(receiptItems.length == 4, "Invalid receipt structure");

        EVMReceipt memory result;
        result.txType = txType;
        result.postStateOrStatus = receiptItems[0].toBytes();
        result.cumulativeGasUsed = uint64(receiptItems[1].toUint());
        result.bloom = receiptItems[2].toBytes();

        // 3. Decode Logs
        RLPReader.RLPItem[] memory logs = receiptItems[3].toList();
        result.logs = new EVMLog[](logs.length);
        
        for (uint256 i = 0; i < logs.length; i++) {
            result.logs[i] = decodeLog(logs[i]);
        }

        return result;
    }

    /**
     * @notice Decodes a single RLP-encoded Log.
     */
    function decodeLog(RLPReader.RLPItem memory encodedLog) internal pure returns (EVMLog memory) {
        RLPReader.RLPItem[] memory log = encodedLog.toList();
        
        // Structure: [address, [topics], data]
        require(log.length == 3, "Invalid log structure");

        EVMLog memory evmLog;
        evmLog.loggerAddress = log[0].toAddress();

        // Decode Topics List
        RLPReader.RLPItem[] memory topics = log[1].toList();
        evmLog.topics = new bytes32[](topics.length);
        for (uint256 i = 0; i < topics.length; i++) {
            evmLog.topics[i] = bytes32(topics[i].toUint());
        }

        evmLog.data = log[2].toBytes();
        return evmLog;
    }
}
