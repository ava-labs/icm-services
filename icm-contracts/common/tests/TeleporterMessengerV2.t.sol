// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test, stdStorage, StdStorage} from "@forge-std/Test.sol";
import {
    Initializable
} from "@openzeppelin/contracts-upgradeable@5.1.0/proxy/utils/Initializable.sol";
import {
    IAdapter,
    TeleporterICMMessage,
    TeleporterMessageV2,
    TeleporterMessageInput,
    TeleporterMessageReceipt,
    TeleporterFeeInfo
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";

/// @dev Adapter that accepts every message, so tests can reach the messenger's own checks.
contract AcceptAllAdapter is IAdapter {
    event MessageSent(TeleporterMessageV2 message);

    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        emit MessageSent(message);
    }

    function verifyMessage(
        TeleporterICMMessage calldata
    ) external pure returns (bool) {
        return true;
    }
}

contract TeleporterMessengerV2InitializeTest is Test {
    bytes32 private constant _BLOCKCHAIN_ID = hex"01";
    bytes32 private constant _OTHER_BLOCKCHAIN_ID = hex"02";
    address private constant _INITIALIZER = address(0x1234);
    address private constant _ATTACKER = address(0xBAD);

    AcceptAllAdapter private _adapter;
    TeleporterMessengerV2 private _teleporter;

    function setUp() public {
        _adapter = new AcceptAllAdapter();
        _teleporter = new TeleporterMessengerV2(address(_adapter), _INITIALIZER);
    }

    function testConstructorRevertsZeroInitializer() public {
        vm.expectRevert("TeleporterMessenger: zero initializer address");
        new TeleporterMessengerV2(address(_adapter), address(0));
    }

    function testConstructorRevertsZeroAdapter() public {
        vm.expectRevert("TeleporterMessenger: zero adapter address");
        new TeleporterMessengerV2(address(0), _INITIALIZER);
    }

    function testInitializeByInitializer() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    function testInitializeRevertsUnauthorized() public {
        vm.prank(_ATTACKER);
        vm.expectRevert("TeleporterMessenger: unauthorized initializer");
        _teleporter.initialize(_OTHER_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), bytes32(0));

        // A failed attempt must not consume the one-shot initializer.
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    function testInitializeRevertsZeroBlockchainID() public {
        vm.prank(_INITIALIZER);
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.initialize(bytes32(0));
    }

    function testInitializeRevertsWhenAlreadyInitialized() public {
        vm.startPrank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        _teleporter.initialize(_OTHER_BLOCKCHAIN_ID);
        vm.stopPrank();
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    /// @dev Before initialization the stored blockchain ID is zero. A message addressed to the zero
    /// blockchain ID must not be accepted as if it were addressed to this chain.
    function testReceiveRevertsWhenUninitialized() public {
        TeleporterICMMessage memory message = _buildMessage(bytes32(0));
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.receiveCrossChainMessage(message, address(0));
    }

    /// @dev Sends must also wait for initialization, otherwise the message ID would be derived from
    /// the zero blockchain ID and could never be matched by later retries or receipts.
    function testSendRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.sendCrossChainMessage(_buildMessageInput());
    }

    function testSendSpecifiedReceiptsRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.sendSpecifiedReceipts(
            _OTHER_BLOCKCHAIN_ID,
            new bytes32[](0),
            TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            new address[](0)
        );
    }

    function testGetNextMessageIDRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.getNextMessageID(_OTHER_BLOCKCHAIN_ID);
    }

    function testSendAfterInitialize() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);

        bytes32 expectedMessageID = _teleporter.getNextMessageID(_OTHER_BLOCKCHAIN_ID);
        bytes32 messageID = _teleporter.sendCrossChainMessage(_buildMessageInput());
        assertEq(messageID, expectedMessageID);
        assertEq(messageID, _teleporter.calculateMessageID(_BLOCKCHAIN_ID, _OTHER_BLOCKCHAIN_ID, 1));
    }

    function testReceiveRevertsWrongDestinationAfterInitialize() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);

        TeleporterICMMessage memory message = _buildMessage(_OTHER_BLOCKCHAIN_ID);
        vm.expectRevert("TeleporterMessenger: invalid destination chain ID");
        _teleporter.receiveCrossChainMessage(message, address(0));
    }

    function testConstructorSetsInitializer() public view {
        assertEq(_teleporter.initializerAddress(), _INITIALIZER);
        assertEq(_teleporter.blockchainID(), bytes32(0));
    }

    function _buildMessage(
        bytes32 destinationBlockchainID
    ) private view returns (TeleporterICMMessage memory) {
        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(this),
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: address(this),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"deadbeef"
        });
        return TeleporterICMMessage({
            message: teleporterMessage,
            sourceNetworkID: 1,
            sourceBlockchainID: _OTHER_BLOCKCHAIN_ID,
            attestation: hex""
        });
    }

    function _buildMessageInput() private pure returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: _OTHER_BLOCKCHAIN_ID,
            destinationAddress: address(0xCAFE),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
    }
}

contract TeleporterMessengerV2MarkReceiptTest is Test {
    using stdStorage for StdStorage;

    bytes32 private constant _BLOCKCHAIN_ID = hex"01";
    bytes32 private constant _OTHER_BLOCKCHAIN_ID = hex"02";
    address private constant _INITIALIZER = address(0x1234);
    address private constant _RELAYER_REWARD_ADDRESS = address(0x5EED);

    AcceptAllAdapter private _adapter;
    TeleporterMessengerV2 private _teleporter;
    UnitTestMockERC20 private _feeAsset;

    event ReceiptReceived(
        bytes32 indexed messageID,
        bytes32 indexed destinationBlockchainID,
        address indexed relayerRewardAddress,
        TeleporterFeeInfo feeInfo
    );

    function setUp() public {
        _adapter = new AcceptAllAdapter();
        _teleporter = new TeleporterMessengerV2(address(_adapter), _INITIALIZER);
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        _feeAsset = new UnitTestMockERC20();
    }

    // Receipts whose fees would overflow the accumulated reward of the relayer reward address must
    // still be processed, so that a fee token with broken accounting can not block delivery of the
    // message carrying the receipts. The accumulated reward saturates at the maximum uint256 value.
    function testReceiptRewardSaturatesInsteadOfOverflowing() public {
        // Send a message with the maximum possible fee amount.
        bytes32 firstMessageID =
            _teleporter.sendCrossChainMessage(_buildMessageInput(type(uint256).max));

        // The mock fee asset now reports the maximum uint256 balance for the Teleporter contract,
        // so reset it to be able to collect the fee of a second message using the same asset.
        stdstore.target(address(_feeAsset)).sig(_feeAsset.mockBalances.selector)
            .with_key(address(_teleporter)).checked_write(uint256(0));

        // Send a second message with the same fee asset.
        bytes32 secondMessageID = _teleporter.sendCrossChainMessage(_buildMessageInput(1));

        // Receive a message with the receipts for both messages sent above, both of which
        // attribute the reward to the same relayer reward address.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](2);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: 1, relayerRewardAddress: _RELAYER_REWARD_ADDRESS
        });
        receipts[1] = TeleporterMessageReceipt({
            receivedMessageNonce: 2, relayerRewardAddress: _RELAYER_REWARD_ADDRESS
        });
        TeleporterICMMessage memory message = _buildMessage(receipts);

        // Both receipts are processed, and the second one saturates the reward rather than
        // reverting the delivery.
        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit ReceiptReceived(
            firstMessageID,
            _OTHER_BLOCKCHAIN_ID,
            _RELAYER_REWARD_ADDRESS,
            TeleporterFeeInfo({feeTokenAddress: address(_feeAsset), amount: type(uint256).max})
        );
        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit ReceiptReceived(
            secondMessageID,
            _OTHER_BLOCKCHAIN_ID,
            _RELAYER_REWARD_ADDRESS,
            TeleporterFeeInfo({feeTokenAddress: address(_feeAsset), amount: 1})
        );
        _teleporter.receiveCrossChainMessage(message, address(0));

        // Check that the reward saturated at the maximum uint256 value.
        assertEq(
            _teleporter.checkRelayerRewardAmount(_RELAYER_REWARD_ADDRESS, address(_feeAsset)),
            type(uint256).max
        );

        // Check that the message received is considered delivered.
        assertTrue(
            _teleporter.messageReceived(
                _teleporter.calculateMessageID(_OTHER_BLOCKCHAIN_ID, _BLOCKCHAIN_ID, 1)
            )
        );

        // Check that the message hashes for the message receipts we received have been cleared.
        assertEq(_teleporter.getMessageHash(firstMessageID), bytes32(0));
        assertEq(_teleporter.getMessageHash(secondMessageID), bytes32(0));
    }

    function _buildMessage(
        TeleporterMessageReceipt[] memory receipts
    ) private view returns (TeleporterICMMessage memory) {
        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(this),
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: _BLOCKCHAIN_ID,
            destinationAddress: address(this),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            receipts: receipts,
            message: hex""
        });
        return TeleporterICMMessage({
            message: teleporterMessage,
            sourceNetworkID: 1,
            sourceBlockchainID: _OTHER_BLOCKCHAIN_ID,
            attestation: hex""
        });
    }

    function _buildMessageInput(
        uint256 feeAmount
    ) private view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: _OTHER_BLOCKCHAIN_ID,
            destinationAddress: address(0xCAFE),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(_feeAsset), amount: feeAmount}),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
    }
}
