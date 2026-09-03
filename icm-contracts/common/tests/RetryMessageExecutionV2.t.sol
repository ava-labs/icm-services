// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {
    TeleporterICMMessage,
    TeleporterMessageReceipt,
    TeleporterMessageV2,
    IAdapter
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";
import {TeleporterMessage} from "@teleporter/ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";

// An adapter that accepts every message, so that the tests can focus on the
// receive and retry logic of TeleporterMessengerV2 rather than on verification.
contract AcceptAllAdapter is IAdapter {
    // solhint-disable-next-line no-empty-blocks
    function sendMessage(
        TeleporterMessageV2 calldata
    ) external {}

    function verifyMessage(
        TeleporterICMMessage calldata
    ) external pure returns (bool) {
        return true;
    }
}

// A receiver that fails until `setShouldRevert(false)` is called, used to make the
// initial message execution fail and a subsequent retry succeed.
contract FlakyMessageReceiverV2 is ITeleporterReceiver {
    bool public shouldRevert = true;
    bytes32 public latestSourceBlockchainID;
    address public latestOriginSenderAddress;
    bytes public latestMessage;

    function setShouldRevert(
        bool shouldRevert_
    ) external {
        shouldRevert = shouldRevert_;
    }

    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        require(!shouldRevert, "receiver: intentional failure");
        latestSourceBlockchainID = sourceBlockchainID;
        latestOriginSenderAddress = originSenderAddress;
        latestMessage = message;
    }
}

contract RetryMessageExecutionV2Test is Test {
    bytes32 private constant _LOCAL_BLOCKCHAIN_ID = bytes32(uint256(0xabcd));
    bytes32 private constant _SOURCE_BLOCKCHAIN_ID = bytes32(uint256(0x1234));
    address private constant _RELAYER_REWARD_ADDRESS = address(0x5678);
    uint256 private constant _MESSAGE_NONCE = 42;
    uint256 private constant _REQUIRED_GAS_LIMIT = 100_000;

    TeleporterMessengerV2 private _teleporter;
    FlakyMessageReceiverV2 private _receiver;

    event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID);
    event MessageExecutionFailed(
        bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, TeleporterMessage message
    );

    function setUp() public {
        _teleporter = new TeleporterMessengerV2(address(new AcceptAllAdapter()));
        _teleporter.initialize(_LOCAL_BLOCKCHAIN_ID);
        _receiver = new FlakyMessageReceiverV2();
    }

    /// @dev A message whose initial execution failed can be executed later via retryMessageExecution.
    function testRetrySucceedsAfterInitialExecutionFailure() public {
        TeleporterMessageV2 memory message = _receiveFailedMessage(address(_receiver));
        bytes32 messageID = _messageID(message.messageNonce);

        // Let the receiver succeed, then retry the execution.
        _receiver.setShouldRevert(false);

        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit MessageExecuted(messageID, _SOURCE_BLOCKCHAIN_ID);
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);

        assertEq(_receiver.latestSourceBlockchainID(), _SOURCE_BLOCKCHAIN_ID);
        assertEq(_receiver.latestOriginSenderAddress(), message.originSenderAddress);
        assertEq(_receiver.latestMessage(), message.message);

        // The failed message hash is cleared, so the message cannot be executed twice.
        assertEq(_teleporter.receivedFailedMessageHashes(messageID), bytes32(0));
        vm.expectRevert("TeleporterMessenger: message not found");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);
    }

    /// @dev A message delivered to an address without code can be executed once a contract is deployed there.
    function testRetrySucceedsAfterReceiverIsDeployed() public {
        address destinationAddress = address(0xdead);
        TeleporterMessageV2 memory message = _receiveFailedMessage(destinationAddress);

        // Mock a receiver being deployed to the destination address after the failed delivery.
        vm.etch(destinationAddress, address(_receiver).code);
        FlakyMessageReceiverV2(destinationAddress).setShouldRevert(false);

        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);

        assertEq(
            FlakyMessageReceiverV2(destinationAddress).latestSourceBlockchainID(),
            _SOURCE_BLOCKCHAIN_ID
        );
        assertEq(FlakyMessageReceiverV2(destinationAddress).latestMessage(), message.message);
    }

    /// @dev Retrying with any altered field is rejected, including the V2-only originTeleporterAddress.
    function testRetryRejectsAlteredMessage() public {
        TeleporterMessageV2 memory message = _receiveFailedMessage(address(_receiver));
        _receiver.setShouldRevert(false);

        TeleporterMessageV2 memory altered = message;

        altered.message = "altered message";
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, altered);
        altered.message = message.message;

        altered.originSenderAddress = address(0xbad);
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, altered);
        altered.originSenderAddress = message.originSenderAddress;

        altered.originTeleporterAddress = address(0xbad);
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, altered);
        altered.originTeleporterAddress = message.originTeleporterAddress;

        altered.requiredGasLimit = _REQUIRED_GAS_LIMIT + 1;
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, altered);
    }

    /// @dev Delivers a message that is received successfully but whose execution fails.
    function _receiveFailedMessage(
        address destinationAddress
    ) private returns (TeleporterMessageV2 memory) {
        TeleporterMessageV2 memory message = TeleporterMessageV2({
            messageNonce: _MESSAGE_NONCE,
            originSenderAddress: address(this),
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: _LOCAL_BLOCKCHAIN_ID,
            destinationAddress: destinationAddress,
            requiredGasLimit: _REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"deadbeef"
        });
        bytes32 messageID = _messageID(message.messageNonce);

        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit MessageExecutionFailed(messageID, _SOURCE_BLOCKCHAIN_ID, _toLegacyMessage(message));
        _teleporter.receiveCrossChainMessage(
            TeleporterICMMessage({
                message: message,
                sourceNetworkID: 1,
                sourceBlockchainID: _SOURCE_BLOCKCHAIN_ID,
                attestation: new bytes(0)
            }),
            _RELAYER_REWARD_ADDRESS
        );

        // The message is received, and its failed execution is recorded for a later retry.
        assertTrue(_teleporter.messageReceived(messageID));
        assertTrue(_teleporter.receivedFailedMessageHashes(messageID) != bytes32(0));

        return message;
    }

    function _messageID(
        uint256 nonce
    ) private view returns (bytes32) {
        return _teleporter.calculateMessageID(_SOURCE_BLOCKCHAIN_ID, _LOCAL_BLOCKCHAIN_ID, nonce);
    }

    function _toLegacyMessage(
        TeleporterMessageV2 memory message
    ) private pure returns (TeleporterMessage memory) {
        return TeleporterMessage({
            messageNonce: message.messageNonce,
            originSenderAddress: message.originSenderAddress,
            destinationBlockchainID: message.destinationBlockchainID,
            destinationAddress: message.destinationAddress,
            requiredGasLimit: message.requiredGasLimit,
            allowedRelayerAddresses: message.allowedRelayerAddresses,
            receipts: message.receipts,
            message: message.message
        });
    }
}
