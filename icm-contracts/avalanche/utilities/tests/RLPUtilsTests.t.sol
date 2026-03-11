// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";
import {RLPUtils} from "../RLPUtils.sol";

/// @dev Minimal RLP encoder for constructing test fixtures.
library RLP {
    function encodeUint(
        uint256 val
    ) internal pure returns (bytes memory) {
        if (val == 0) return encodeBytes(new bytes(0));
        return encodeBytes(_toBinary(val));
    }

    function encodeAddress(
        address addr
    ) internal pure returns (bytes memory) {
        return encodeBytes(abi.encodePacked(addr));
    }

    function encodeBytes(
        bytes memory data
    ) internal pure returns (bytes memory) {
        if (data.length == 1 && uint8(data[0]) <= 0x7f) {
            return data;
        } else if (data.length <= 55) {
            return abi.encodePacked(uint8(0x80 + data.length), data);
        } else {
            bytes memory lenBytes = _toBinary(data.length);
            return abi.encodePacked(uint8(0xb7 + lenBytes.length), lenBytes, data);
        }
    }

    function encodeList(
        bytes memory items
    ) internal pure returns (bytes memory) {
        if (items.length <= 55) {
            return abi.encodePacked(uint8(0xc0 + items.length), items);
        } else {
            bytes memory lenBytes = _toBinary(items.length);
            return abi.encodePacked(uint8(0xf7 + lenBytes.length), lenBytes, items);
        }
    }

    function _toBinary(
        uint256 val
    ) private pure returns (bytes memory) {
        if (val == 0) return new bytes(0);
        uint256 len;
        uint256 temp = val;
        while (temp > 0) {
            len++;
            temp >>= 8;
        }
        bytes memory result = new bytes(len);
        for (uint256 i = len; i > 0; i--) {
            result[i - 1] = bytes1(uint8(val));
            val >>= 8;
        }
        return result;
    }
}

/// @dev Exposes internal library functions for testing.
contract RLPUtilsHarness {
    using RLPReader for bytes;

    function decodeReceipt(
        bytes memory encoded
    ) external pure returns (RLPUtils.EVMReceipt memory) {
        return RLPUtils.decodeReceipt(encoded.toRlpItem());
    }

    function decodeLog(
        bytes memory encoded
    ) external pure returns (RLPUtils.EVMLog memory) {
        return RLPUtils.decodeLog(encoded.toRlpItem());
    }
}

contract RLPUtilsTest is Test {
    using RLP for uint256;
    using RLP for address;
    using RLP for bytes;

    address private constant _LOG_ADDR = 0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa;
    bytes32 private constant _TOPIC_0 = keccak256("Transfer(address,address,uint256)");
    bytes32 private constant _TOPIC_1 = bytes32(uint256(uint160(0xBEEF)));
    bytes32 private constant _TOPIC_2 = bytes32(uint256(uint160(0xCAFE)));
    bytes private constant _LOG_DATA =
        hex"0000000000000000000000000000000000000000000000000000000000000064";
    uint64 private constant _GAS_USED = 21000;

    RLPUtilsHarness private _harness;

    function setUp() public {
        _harness = new RLPUtilsHarness();
    }

    function testDecodeReceiptLegacyNoLogs() public {
        bytes[] memory logs = new bytes[](0);
        bytes memory encoded = _buildReceiptRLP(1, _GAS_USED, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);

        assertEq(r.txType, 0);
        assertEq(uint8(r.postStateOrStatus[0]), 1);
        assertEq(r.cumulativeGasUsed, _GAS_USED);
        assertEq(r.bloom.length, 256);
        assertEq(r.logs.length, 0);
    }

    function testDecodeReceiptLegacyWithLogs() public {
        bytes32[] memory topics = _threeTopicLog();
        bytes memory log = _buildLog(_LOG_ADDR, topics, _LOG_DATA);
        bytes[] memory logs = new bytes[](1);
        logs[0] = log;
        bytes memory encoded = _buildReceiptRLP(1, _GAS_USED, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);

        assertEq(r.txType, 0);
        assertEq(r.logs.length, 1);
        assertEq(r.logs[0].loggerAddress, _LOG_ADDR);
        assertEq(r.logs[0].topics.length, 3);
        assertEq(r.logs[0].topics[0], _TOPIC_0);
        assertEq(r.logs[0].topics[1], _TOPIC_1);
        assertEq(r.logs[0].topics[2], _TOPIC_2);
        assertEq(r.logs[0].data, _LOG_DATA);
    }

    function testDecodeReceiptLegacyFailedTx() public {
        bytes[] memory logs = new bytes[](0);
        bytes memory encoded = _buildReceiptRLP(0, _GAS_USED, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);

        assertEq(r.postStateOrStatus.length, 0);
    }

    function testDecodeReceiptLegacyMultipleLogs() public {
        bytes32[] memory topics1 = _singleTopicLog();
        bytes32[] memory topics2 = _threeTopicLog();
        bytes memory log1 = _buildLog(_LOG_ADDR, topics1, _LOG_DATA);
        bytes memory log2 = _buildLog(address(0xBEEF), topics2, hex"");
        bytes[] memory logs = new bytes[](2);
        logs[0] = log1;
        logs[1] = log2;
        bytes memory encoded = _buildReceiptRLP(1, 42000, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);

        assertEq(r.logs.length, 2);
        assertEq(r.logs[0].loggerAddress, _LOG_ADDR);
        assertEq(r.logs[0].topics.length, 1);
        assertEq(r.logs[1].loggerAddress, address(0xBEEF));
        assertEq(r.logs[1].topics.length, 3);
        assertEq(r.cumulativeGasUsed, 42000);
    }

    function testDecodeReceiptTypedWithoutLogs() public {
        bytes[] memory logs = new bytes[](0);
        bytes memory encoded = _buildTypedReceipt(1, 1, _GAS_USED, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);
        assertEq(r.txType, 1);
        assertEq(r.cumulativeGasUsed, _GAS_USED);
    }

    function testDecodeReceiptTypedWithLogs() public {
        bytes32[] memory topics = _singleTopicLog();
        bytes memory log = _buildLog(_LOG_ADDR, topics, _LOG_DATA);
        bytes[] memory logs = new bytes[](1);
        logs[0] = log;
        bytes memory encoded = _buildTypedReceipt(2, 1, 50000, _emptyBloom(), logs);

        RLPUtils.EVMReceipt memory r = _harness.decodeReceipt(encoded);
        assertEq(r.txType, 2);
        assertEq(r.cumulativeGasUsed, 50000);
        assertEq(r.logs.length, 1);
        assertEq(r.logs[0].loggerAddress, _LOG_ADDR);
    }

    function testDecodeReceiptRevertsUnsupportedType() public {
        bytes[] memory logs = new bytes[](0);
        bytes memory encoded = _buildTypedReceipt(5, 1, _GAS_USED, _emptyBloom(), logs);

        vm.expectRevert("Unsupported tx type");
        _harness.decodeReceipt(encoded);
    }

    function testDecodeReceiptRevertsInvalidFieldCount() public {
        bytes memory items = abi.encodePacked(
            uint256(1).encodeUint(), uint256(_GAS_USED).encodeUint(), _emptyBloom().encodeBytes()
        );
        bytes memory encoded = items.encodeList();

        vm.expectRevert("Invalid receipt: expected 4 fields");
        _harness.decodeReceipt(encoded);
    }

    function testDecodeLogZeroTopics() public {
        bytes32[] memory topics = new bytes32[](0);
        bytes memory encoded = _buildLog(_LOG_ADDR, topics, _LOG_DATA);

        RLPUtils.EVMLog memory log = _harness.decodeLog(encoded);

        assertEq(log.loggerAddress, _LOG_ADDR);
        assertEq(log.topics.length, 0);
        assertEq(log.data, _LOG_DATA);
    }

    function testDecodeLogFourTopics() public {
        bytes32[] memory topics = new bytes32[](4);
        topics[0] = _TOPIC_0;
        topics[1] = _TOPIC_1;
        topics[2] = _TOPIC_2;
        topics[3] = bytes32(uint256(0xDEAD));
        bytes memory encoded = _buildLog(_LOG_ADDR, topics, _LOG_DATA);

        RLPUtils.EVMLog memory log = _harness.decodeLog(encoded);

        assertEq(log.topics.length, 4);
        assertEq(log.topics[0], _TOPIC_0);
        assertEq(log.topics[3], bytes32(uint256(0xDEAD)));
    }

    function testDecodeLogEmptyData() public {
        bytes32[] memory topics = _singleTopicLog();
        bytes memory encoded = _buildLog(_LOG_ADDR, topics, hex"");

        RLPUtils.EVMLog memory log = _harness.decodeLog(encoded);

        assertEq(log.data.length, 0);
    }

    function testDecodeLogRevertsTooManyTopics() public {
        bytes32[] memory topics = new bytes32[](5);
        for (uint256 i; i < 5; i++) {
            topics[i] = bytes32(i);
        }
        bytes memory encoded = _buildLog(_LOG_ADDR, topics, _LOG_DATA);

        vm.expectRevert("Invalid log: too many topics");
        _harness.decodeLog(encoded);
    }

    function testDecodeLogRevertsInvalidFieldCount() public {
        bytes memory items = abi.encodePacked(_LOG_ADDR.encodeAddress(), _LOG_DATA.encodeBytes());
        bytes memory encoded = items.encodeList();

        vm.expectRevert("Invalid log: expected 3 fields");
        _harness.decodeLog(encoded);
    }

    // Helpers

    function _emptyBloom() internal pure returns (bytes memory) {
        return new bytes(256);
    }

    /// @dev Builds an RLP-encoded log: [address, [topics...], data]
    function _buildLog(
        address addr,
        bytes32[] memory topics,
        bytes memory data
    ) internal pure returns (bytes memory) {
        bytes memory topicsConcat;
        for (uint256 i; i < topics.length; i++) {
            topicsConcat = abi.encodePacked(topicsConcat, uint256(topics[i]).encodeUint());
        }
        bytes memory encodedTopics = topicsConcat.encodeList();
        bytes memory items =
            abi.encodePacked(addr.encodeAddress(), encodedTopics, data.encodeBytes());
        return items.encodeList();
    }

    /// @dev Builds a full RLP-encoded receipt without type prefix.
    function _buildReceiptRLP(
        uint8 status,
        uint64 gas,
        bytes memory bloom,
        bytes[] memory encodedLogs
    ) internal pure returns (bytes memory) {
        bytes memory logsConcat;
        for (uint256 i; i < encodedLogs.length; i++) {
            logsConcat = abi.encodePacked(logsConcat, encodedLogs[i]);
        }
        bytes memory logsList = logsConcat.encodeList();
        bytes memory items = abi.encodePacked(
            uint256(status).encodeUint(), uint256(gas).encodeUint(), bloom.encodeBytes(), logsList
        );
        return items.encodeList();
    }

    /// @dev Builds a typed receipt: type_byte || rlp([...])
    function _buildTypedReceipt(
        uint8 txType,
        uint8 status,
        uint64 gas,
        bytes memory bloom,
        bytes[] memory encodedLogs
    ) internal pure returns (bytes memory) {
        bytes memory receiptRLP = _buildReceiptRLP(status, gas, bloom, encodedLogs);
        return abi.encodePacked(bytes1(txType), receiptRLP);
    }

    function _singleTopicLog() internal pure returns (bytes32[] memory topics) {
        topics = new bytes32[](1);
        topics[0] = _TOPIC_0;
    }

    function _threeTopicLog() internal pure returns (bytes32[] memory topics) {
        topics = new bytes32[](3);
        topics[0] = _TOPIC_0;
        topics[1] = _TOPIC_1;
        topics[2] = _TOPIC_2;
    }
}
