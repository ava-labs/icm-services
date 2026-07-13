// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {OracleAdapter, OracleMessage} from "../OracleAdapter.sol";
import {
    TeleporterICMMessage,
    TeleporterMessageV2,
    TeleporterMessageReceipt
} from "@common/ITeleporterMessengerV2.sol";
import {WarpMessage, IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";

contract OracleAdapterTest is Test {
    address private constant _WARP_PRECOMPILE = 0x0200000000000000000000000000000000000005;
    bytes32 private constant _THIS_CHAIN_ID = bytes32(uint256(0xABCD));
    string private constant _SOURCE_TYPE = "solana";
    string private constant _SOURCE_ADDR = "So11111111111111111111111111111111111111112";

    OracleAdapter private _adapter;
    InlineOracleReceiver private _receiver;

    function setUp() public {
        vm.mockCall(
            _WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.getBlockchainID, ()),
            abi.encode(_THIS_CHAIN_ID)
        );
        _adapter = new OracleAdapter(address(this));
        _receiver = new InlineOracleReceiver();
        _adapter.setAllowedSource(_SOURCE_TYPE, _SOURCE_ADDR, true);
    }

    // -------------------------------------------------------------------------
    // verifyMessage — happy path + event
    // -------------------------------------------------------------------------

    function testVerifyMessageHappyPath() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, true);

        vm.expectEmit(true, true, false, true, address(_adapter));
        emit OracleAdapter.OracleMessageVerified(
            oracleMsg.nonce, oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.destContract
        );

        bool result = _adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
        assertTrue(result);
        assertTrue(_adapter.isProcessed(oracleMsg.nonce));
    }

    // -------------------------------------------------------------------------
    // verifyMessage — revert paths
    // -------------------------------------------------------------------------

    function testVerifyMessageInvalidWarp() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, false);

        vm.expectRevert(OracleAdapter.InvalidWarpMessage.selector);
        _adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function testVerifyMessageWrongSourceChain() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        bytes32 wrongChainID = bytes32(uint256(0xDEAD));

        WarpMessage memory warpMsg = WarpMessage({
            sourceChainID: wrongChainID,
            originSenderAddress: address(0),
            payload: _oracleMsgPayload(oracleMsg)
        });
        vm.mockCall(
            _WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (0)),
            abi.encode(warpMsg, true)
        );

        vm.expectRevert(
            abi.encodeWithSelector(
                OracleAdapter.WrongSourceChain.selector, wrongChainID, _THIS_CHAIN_ID
            )
        );
        _adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function testVerifyMessageSourceNotAllowed() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        oracleMsg.sourceType = "bitcoin";
        oracleMsg.sourceAddress = "bc1qunlisted";
        _mockWarp(0, oracleMsg, true);

        vm.expectRevert(
            abi.encodeWithSelector(
                OracleAdapter.SourceNotAllowed.selector,
                oracleMsg.sourceType,
                oracleMsg.sourceAddress
            )
        );
        _adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function testVerifyMessageAlreadyProcessed() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, true);
        _adapter.verifyMessage(_buildICMMsg(0, oracleMsg));

        // Second delivery: different warp index, same global nonce.
        _mockWarp(1, oracleMsg, true);
        vm.expectRevert(
            abi.encodeWithSelector(OracleAdapter.AlreadyProcessed.selector, oracleMsg.nonce)
        );
        _adapter.verifyMessage(_buildICMMsg(1, oracleMsg));
    }

    // -------------------------------------------------------------------------
    // sendMessage
    // -------------------------------------------------------------------------

    function testSendMessageEmitsEvent() public {
        TeleporterMessageV2 memory teleporterMsg = _emptyTeleporterMsg();

        vm.expectEmit(false, false, false, true, address(_adapter));
        emit OracleAdapter.OracleMessageSent(teleporterMsg);

        _adapter.sendMessage(teleporterMsg);
    }

    // -------------------------------------------------------------------------
    // Admin — onlyOwner guards
    // -------------------------------------------------------------------------

    function testSetAllowedSourceOnlyOwner() public {
        vm.prank(address(0xBEEF));
        vm.expectRevert(OracleAdapter.Unauthorized.selector);
        _adapter.setAllowedSource("solana", "addr", true);
    }

    function testTransferOwnershipOnlyOwner() public {
        vm.prank(address(0xBEEF));
        vm.expectRevert(OracleAdapter.Unauthorized.selector);
        _adapter.transferOwnership(address(0xDEAD));
    }

    function testTransferOwnershipZeroAddress() public {
        vm.expectRevert(OracleAdapter.ZeroAddress.selector);
        _adapter.transferOwnership(address(0));
    }

    // -------------------------------------------------------------------------
    // Internal helpers
    // -------------------------------------------------------------------------

    function _mockWarp(uint32 warpIndex, OracleMessage memory oracleMsg, bool valid) internal {
        WarpMessage memory warpMsg = WarpMessage({
            sourceChainID: _THIS_CHAIN_ID,
            originSenderAddress: address(0),
            payload: _oracleMsgPayload(oracleMsg)
        });
        vm.mockCall(
            _WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (warpIndex)),
            abi.encode(warpMsg, valid)
        );
    }

    function _defaultOracleMsg() internal view returns (OracleMessage memory) {
        return OracleMessage({
            sourceType: _SOURCE_TYPE,
            sourceAddress: _SOURCE_ADDR,
            destContract: address(_receiver),
            sourceBlockHeight: 1_000_000,
            nonce: 1,
            payload: hex"cafebabe"
        });
    }

    function _oracleMsgPayload(
        OracleMessage memory oracleMsg
    ) internal pure returns (bytes memory) {
        return abi.encode(
            oracleMsg.sourceType,
            oracleMsg.sourceAddress,
            oracleMsg.destContract,
            oracleMsg.sourceBlockHeight,
            oracleMsg.nonce,
            oracleMsg.payload
        );
    }

    function _buildICMMsg(
        uint32 warpIndex,
        OracleMessage memory oracleMsg
    ) internal pure returns (TeleporterICMMessage memory) {
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](0);
        address[] memory relayers = new address[](0);
        return TeleporterICMMessage({
            message: TeleporterMessageV2({
                messageNonce: oracleMsg.nonce,
                originSenderAddress: address(0),
                originTeleporterAddress: address(0),
                destinationBlockchainID: _THIS_CHAIN_ID,
                destinationAddress: oracleMsg.destContract,
                requiredGasLimit: 0,
                allowedRelayerAddresses: relayers,
                receipts: receipts,
                message: hex""
            }),
            sourceNetworkID: 1,
            sourceBlockchainID: _THIS_CHAIN_ID,
            attestation: abi.encode(warpIndex)
        });
    }

    function _emptyTeleporterMsg() internal pure returns (TeleporterMessageV2 memory) {
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](0);
        address[] memory relayers = new address[](0);
        return TeleporterMessageV2({
            messageNonce: 0,
            originSenderAddress: address(0),
            originTeleporterAddress: address(0),
            destinationBlockchainID: bytes32(0),
            destinationAddress: address(0),
            requiredGasLimit: 0,
            allowedRelayerAddresses: relayers,
            receipts: receipts,
            message: hex""
        });
    }
}

contract InlineOracleReceiver is ITeleporterReceiver {
    bytes public lastMessage;

    function receiveTeleporterMessage(bytes32, address, bytes calldata message) external override {
        lastMessage = message;
    }
}
