// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {
    TeleporterMessengerTest,
    TeleporterMessageReceipt,
    TeleporterMessage,
    TeleporterFeeInfo,
    WarpMessage
} from "./TeleporterMessengerTest.t.sol";
import {stdStorage, StdStorage} from "@forge-std/Test.sol";

contract MarkReceiptTest is TeleporterMessengerTest {
    using stdStorage for StdStorage;

    struct FeeRewardInfo {
        uint256 feeAmount;
        address relayerRewardAddress;
    }

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    // This test first mocks sending 3 messages (each with fees) to a different L1.
    // It then mocks receiving a message back from that L1 with the receipt information
    // for each of the 3 messages sent, and checks that the relayers can redeem their rewards.
    function testCheckRelayersUponReceipt() public {
        // Submit a few mock messages to be sent.
        FeeRewardInfo[3] memory feeRewardInfos = [
            FeeRewardInfo(1111111111111111, 0x52A258ED593C793251a89bfd36caE158EE9fC4F8),
            FeeRewardInfo(2222222222222222, 0xeF6ed43EB8Ff15E336D64d1468947cA1046824E6),
            FeeRewardInfo(3333333333333333, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542)
        ];

        uint256[3] memory messageNonces;
        bytes32[3] memory messageIDs;
        for (uint256 i; i < feeRewardInfos.length; ++i) {
            messageNonces[i] = _getNextMessageNonce();
            messageIDs[i] =
                _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, feeRewardInfos[i].feeAmount);
        }

        // Mock receiving a message with the 3 receipts from the mock messages sent above.
        TeleporterMessageReceipt[] memory receipts =
            new TeleporterMessageReceipt[](feeRewardInfos.length);
        for (uint256 i; i < receipts.length; ++i) {
            receipts[i] = TeleporterMessageReceipt({
                receivedMessageNonce: messageNonces[i],
                relayerRewardAddress: feeRewardInfos[i].relayerRewardAddress
            });
        }
        uint256 receivedMessageNonce = 1;
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(receivedMessageNonce, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x93753a9eA4C9D6eeed9f64eA92E97ce1f5FBAeDe;

        for (uint256 i; i < feeRewardInfos.length; ++i) {
            vm.expectEmit(true, true, true, true, address(teleporterMessenger));
            emit ReceiptReceived(
                messageIDs[i],
                DEFAULT_SOURCE_BLOCKCHAIN_ID,
                feeRewardInfos[i].relayerRewardAddress,
                TeleporterFeeInfo({
                    feeTokenAddress: address(_mockFeeAsset), amount: feeRewardInfos[i].feeAmount
                })
            );
        }
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the relayers have redeemable balances
        for (uint256 i; i < feeRewardInfos.length; ++i) {
            assertEq(
                teleporterMessenger.checkRelayerRewardAmount(
                    feeRewardInfos[i].relayerRewardAddress, address(_mockFeeAsset)
                ),
                feeRewardInfos[i].feeAmount
            );
        }

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, receivedMessageNonce
        );
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(expectedMessageID),
            expectedRelayerRewardAddress
        );
        assertTrue(teleporterMessenger.messageReceived(expectedMessageID));

        // Check that the message hashes for the message receipts we received have been cleared.
        for (uint256 i; i < receipts.length; ++i) {
            assertEq(teleporterMessenger.getMessageHash(messageIDs[i]), bytes32(0));
        }
    }

    function testReceiptForNoFeeMessage() public {
        // Submit a a mock message with no fee.
        uint256 sentMessageNonce = _getNextMessageNonce();
        bytes32 sentMessageID = _sendTestMessageWithNoFee(DEFAULT_SOURCE_BLOCKCHAIN_ID);

        // Mock receiving a message with the a receipts of the mock message sent above.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: sentMessageNonce,
            relayerRewardAddress: DEFAULT_RELAYER_REWARD_ADDRESS
        });
        uint256 receivedMessageNonce = 42;
        bytes32 receivedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, receivedMessageNonce
        );
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(receivedMessageNonce, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x2F20537C2F5c57231866DE9D0CE33d0681a200D4;

        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiptReceived(
            sentMessageID,
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            DEFAULT_RELAYER_REWARD_ADDRESS,
            TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0})
        );
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(receivedMessageID),
            expectedRelayerRewardAddress
        );
        assertTrue(teleporterMessenger.messageReceived(receivedMessageID));
    }

    function testDuplicateReceiptAllowed() public {
        // Submit a mock message to be sent.
        FeeRewardInfo memory feeRewardInfo =
            FeeRewardInfo(1111111111111111, 0x52A258ED593C793251a89bfd36caE158EE9fC4F8);
        uint256 sentMessageNonce = _getNextMessageNonce();
        _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, feeRewardInfo.feeAmount);

        // Mock receiving a message with the 2 receipts for the  same mock message above.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](2);
        TeleporterMessageReceipt memory receipt = TeleporterMessageReceipt({
            receivedMessageNonce: sentMessageNonce,
            relayerRewardAddress: feeRewardInfo.relayerRewardAddress
        });
        receipts[0] = receipt;
        receipts[1] = receipt;
        uint256 receivedMessageNonce = 12;
        bytes32 receivedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, receivedMessageNonce
        );
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(receivedMessageNonce, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x6DAEF0D63ea44C801b655Fd97fe3804B9bFCC097;
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the relayer redeemable balances was only added once.
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(
                feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset)
            ),
            feeRewardInfo.feeAmount
        );

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(receivedMessageID),
            expectedRelayerRewardAddress
        );
        assertTrue(teleporterMessenger.messageReceived(receivedMessageID));

        // Check that the message hashes for the message receipts we received have been cleared.
        assertEq(teleporterMessenger.getMessageHash(receivedMessageID), bytes32(0));
    }

    // Receipts whose fees would overflow the accumulated reward of the relayer reward address must
    // still be processed, so that a fee token with broken accounting can not block delivery of the
    // message carrying the receipts. The accumulated reward saturates at the maximum uint256 value.
    function testReceiptRewardSaturatesInsteadOfOverflowing() public {
        address relayerRewardAddress = 0x52A258ED593C793251a89bfd36caE158EE9fC4F8;

        // Submit a mock message with the maximum possible fee amount.
        uint256 firstMessageNonce = _getNextMessageNonce();
        bytes32 firstMessageID =
            _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, type(uint256).max);

        // The mock fee asset now reports the maximum uint256 balance for the Teleporter contract,
        // so reset it to be able to collect the fee of a second message using the same asset.
        stdstore.target(address(_mockFeeAsset)).sig(_mockFeeAsset.mockBalances.selector)
            .with_key(address(teleporterMessenger)).checked_write(uint256(0));

        // Submit a second mock message with the same fee asset.
        uint256 secondMessageNonce = _getNextMessageNonce();
        bytes32 secondMessageID = _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, 1);

        // Mock receiving a message with the receipts for both messages sent above, both of which
        // attribute the reward to the same relayer reward address.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](2);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: firstMessageNonce, relayerRewardAddress: relayerRewardAddress
        });
        receipts[1] = TeleporterMessageReceipt({
            receivedMessageNonce: secondMessageNonce, relayerRewardAddress: relayerRewardAddress
        });
        uint256 receivedMessageNonce = 7;
        bytes32 receivedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, receivedMessageNonce
        );
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(receivedMessageNonce, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message. Both receipts are processed, and the second one saturates the
        // reward rather than reverting the delivery.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiptReceived(
            firstMessageID,
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            relayerRewardAddress,
            TeleporterFeeInfo({feeTokenAddress: address(_mockFeeAsset), amount: type(uint256).max})
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiptReceived(
            secondMessageID,
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            relayerRewardAddress,
            TeleporterFeeInfo({feeTokenAddress: address(_mockFeeAsset), amount: 1})
        );
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the reward saturated at the maximum uint256 value.
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(
                relayerRewardAddress, address(_mockFeeAsset)
            ),
            type(uint256).max
        );

        // Check that the message received is considered delivered.
        assertTrue(teleporterMessenger.messageReceived(receivedMessageID));

        // Check that the message hashes for the message receipts we received have been cleared.
        assertEq(teleporterMessenger.getMessageHash(firstMessageID), bytes32(0));
        assertEq(teleporterMessenger.getMessageHash(secondMessageID), bytes32(0));
    }
}
