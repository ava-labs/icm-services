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
    TeleporterMessage,
    TeleporterMessageV2,
    TeleporterMessageInput,
    TeleporterMessageReceipt,
    TeleporterFeeInfo
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
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

/// @dev Receiver that reverts while `shouldRevert` is set, and otherwise records what was delivered to it.
contract SwitchableReceiver is ITeleporterReceiver {
    address public immutable teleporterContract;
    bool public shouldRevert = true;
    bytes32 public lastSourceBlockchainID;
    address public lastOriginSenderAddress;
    bytes public lastMessage;

    constructor(
        address teleporterContract_
    ) {
        teleporterContract = teleporterContract_;
    }

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
        require(msg.sender == teleporterContract, "SwitchableReceiver: unauthorized");
        require(!shouldRevert, "SwitchableReceiver: execution failed");
        lastSourceBlockchainID = sourceBlockchainID;
        lastOriginSenderAddress = originSenderAddress;
        lastMessage = message;
    }
}

contract TeleporterMessengerV2RetryMessageExecutionTest is Test {
    bytes32 private constant _BLOCKCHAIN_ID = hex"01";
    bytes32 private constant _OTHER_BLOCKCHAIN_ID = hex"02";
    address private constant _INITIALIZER = address(0x1234);
    address private constant _ORIGIN_SENDER = address(0x5E4D);
    address private constant _RELAYER_REWARD_ADDRESS = address(0x5EED);
    uint256 private constant _MESSAGE_NONCE = 42;
    bytes private constant _PAYLOAD = hex"deadbeef";

    AcceptAllAdapter private _adapter;
    TeleporterMessengerV2 private _teleporter;
    SwitchableReceiver private _receiver;

    event MessageExecutionFailed(
        bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, TeleporterMessage message
    );
    event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID);

    function setUp() public {
        _adapter = new AcceptAllAdapter();
        _teleporter = new TeleporterMessengerV2(address(_adapter), _INITIALIZER);
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        _receiver = new SwitchableReceiver(address(_teleporter));
    }

    /// @dev A failed message is stored as the hash of the {TeleporterMessage} emitted in
    /// {MessageExecutionFailed}, so passing that struct back must execute the message.
    function testRetryExecutesFailedMessage() public {
        (bytes32 messageID, TeleporterMessage memory message) =
            _receiveFailedMessage(address(_receiver));

        _receiver.setShouldRevert(false);
        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit MessageExecuted(messageID, _OTHER_BLOCKCHAIN_ID);
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);

        assertEq(_receiver.lastSourceBlockchainID(), _OTHER_BLOCKCHAIN_ID);
        assertEq(_receiver.lastOriginSenderAddress(), _ORIGIN_SENDER);
        assertEq(_receiver.lastMessage(), _PAYLOAD);
        assertEq(_teleporter.receivedFailedMessageHashes(messageID), bytes32(0));

        // A message can be executed at most once.
        vm.expectRevert("TeleporterMessenger: message not found");
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);
    }

    function testRetryRevertsWhenExecutionFailsAgain() public {
        (bytes32 messageID, TeleporterMessage memory message) =
            _receiveFailedMessage(address(_receiver));

        vm.expectRevert("TeleporterMessenger: retry execution failed");
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);

        // The message stays retryable.
        assertEq(_teleporter.receivedFailedMessageHashes(messageID), keccak256(abi.encode(message)));
    }

    function testRetryRevertsInvalidMessageHash() public {
        (bytes32 messageID, TeleporterMessage memory message) =
            _receiveFailedMessage(address(_receiver));
        bytes32 failedMessageHash = _teleporter.receivedFailedMessageHashes(messageID);
        _receiver.setShouldRevert(false);

        message.message = hex"beef";
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);

        assertEq(_teleporter.receivedFailedMessageHashes(messageID), failedMessageHash);
        assertEq(_receiver.lastMessage(), hex"");
    }

    function testRetryRevertsMessageNotFound() public {
        TeleporterMessage memory message =
            _toTeleporterMessage(_buildMessage(address(_receiver)).message);

        vm.expectRevert("TeleporterMessenger: message not found");
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);
    }

    /// @dev A message delivered to an address without code is stored as failed, and can be retried
    /// once a contract has been deployed there.
    function testRetryAfterCodeDeployedAtDestination() public {
        address destination = address(0xCAFE);
        (bytes32 messageID, TeleporterMessage memory message) = _receiveFailedMessage(destination);

        vm.expectRevert("TeleporterMessenger: destination address has no code");
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);

        // Deploy the receiver's code at the destination. Its storage is empty there, so it does
        // not revert.
        vm.etch(destination, address(_receiver).code);
        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit MessageExecuted(messageID, _OTHER_BLOCKCHAIN_ID);
        _teleporter.retryMessageExecution(_OTHER_BLOCKCHAIN_ID, message);

        assertEq(SwitchableReceiver(destination).lastSourceBlockchainID(), _OTHER_BLOCKCHAIN_ID);
        assertEq(SwitchableReceiver(destination).lastOriginSenderAddress(), _ORIGIN_SENDER);
        assertEq(SwitchableReceiver(destination).lastMessage(), _PAYLOAD);
        assertEq(_teleporter.receivedFailedMessageHashes(messageID), bytes32(0));
    }

    /// @dev Delivers a message to `destination` whose execution fails, and returns its message ID
    /// together with the {TeleporterMessage} emitted in {MessageExecutionFailed}.
    function _receiveFailedMessage(
        address destination
    ) private returns (bytes32, TeleporterMessage memory) {
        TeleporterICMMessage memory icmMessage = _buildMessage(destination);
        TeleporterMessage memory message = _toTeleporterMessage(icmMessage.message);
        bytes32 messageID =
            _teleporter.calculateMessageID(_OTHER_BLOCKCHAIN_ID, _BLOCKCHAIN_ID, _MESSAGE_NONCE);

        vm.expectEmit(true, true, true, true, address(_teleporter));
        emit MessageExecutionFailed(messageID, _OTHER_BLOCKCHAIN_ID, message);
        _teleporter.receiveCrossChainMessage(icmMessage, _RELAYER_REWARD_ADDRESS);

        // The message counts as delivered, and its hash is stored for retries.
        assertTrue(_teleporter.messageReceived(messageID));
        assertEq(_teleporter.getRelayerRewardAddress(messageID), _RELAYER_REWARD_ADDRESS);
        assertEq(_teleporter.receivedFailedMessageHashes(messageID), keccak256(abi.encode(message)));
        return (messageID, message);
    }

    /// @dev Builds a message with non-empty relayer and receipt lists, so that the dynamic fields
    /// take part in the hash comparison.
    function _buildMessage(
        address destination
    ) private view returns (TeleporterICMMessage memory) {
        address[] memory allowedRelayerAddresses = new address[](1);
        allowedRelayerAddresses[0] = address(this);
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: 7, relayerRewardAddress: _RELAYER_REWARD_ADDRESS
        });
        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: _MESSAGE_NONCE,
            originSenderAddress: _ORIGIN_SENDER,
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: _BLOCKCHAIN_ID,
            destinationAddress: destination,
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: allowedRelayerAddresses,
            receipts: receipts,
            message: _PAYLOAD
        });
        return TeleporterICMMessage({
            message: teleporterMessage,
            sourceNetworkID: 1,
            sourceBlockchainID: _OTHER_BLOCKCHAIN_ID,
            attestation: hex""
        });
    }

    /// @dev Mirrors the conversion the messenger performs on delivery.
    function _toTeleporterMessage(
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
