// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {OracleAdapter, OracleMessage} from "./OracleAdapter.sol";
import {TeleporterICMMessage, TeleporterMessageV2, TeleporterMessageReceipt} from "@common/ITeleporterMessengerV2.sol";
import {WarpMessage, IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";

contract OracleAdapterTest is Test {
    address constant WARP_PRECOMPILE = 0x0200000000000000000000000000000000000005;
    bytes32 constant THIS_CHAIN_ID = bytes32(uint256(0xABCD));
    string constant SOURCE_TYPE = "solana";
    string constant SOURCE_ADDR = "So11111111111111111111111111111111111111112";

    OracleAdapter adapter;
    InlineOracleReceiver receiver;

    function setUp() public {
        vm.mockCall(
            WARP_PRECOMPILE,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(THIS_CHAIN_ID)
        );
        adapter = new OracleAdapter(address(this));
        receiver = new InlineOracleReceiver();
        adapter.setAllowedSource(SOURCE_TYPE, SOURCE_ADDR, true);
    }

    // -------------------------------------------------------------------------
    // verifyMessage — happy path + event
    // -------------------------------------------------------------------------

    function test_verifyMessage_happyPath() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, true);

        vm.expectEmit(true, true, false, true, address(adapter));
        emit OracleAdapter.OracleMessageVerified(
            oracleMsg.nonce, oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.destContract
        );

        bool result = adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
        assertTrue(result);
        assertTrue(adapter.isProcessed(oracleMsg.nonce));
    }

    // -------------------------------------------------------------------------
    // verifyMessage — revert paths
    // -------------------------------------------------------------------------

    function test_verifyMessage_invalidWarp() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, false);

        vm.expectRevert(OracleAdapter.InvalidWarpMessage.selector);
        adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function test_verifyMessage_wrongSourceChain() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        bytes32 wrongChainID = bytes32(uint256(0xDEAD));

        WarpMessage memory warpMsg = WarpMessage({
            sourceChainID: wrongChainID,
            originSenderAddress: address(0),
            payload: _oracleMsgPayload(oracleMsg)
        });
        vm.mockCall(
            WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (0)),
            abi.encode(warpMsg, true)
        );

        vm.expectRevert(
            abi.encodeWithSelector(
                OracleAdapter.WrongSourceChain.selector, wrongChainID, THIS_CHAIN_ID
            )
        );
        adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function test_verifyMessage_payloadMismatch() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, true);

        // Attestation carries a tampered nonce — warp payload no longer matches.
        OracleMessage memory tampered = oracleMsg;
        tampered.nonce = oracleMsg.nonce + 1;

        vm.expectRevert(OracleAdapter.PayloadMismatch.selector);
        adapter.verifyMessage(_buildICMMsg(0, tampered));
    }

    function test_verifyMessage_sourceNotAllowed() public {
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
        adapter.verifyMessage(_buildICMMsg(0, oracleMsg));
    }

    function test_verifyMessage_alreadyProcessed() public {
        OracleMessage memory oracleMsg = _defaultOracleMsg();
        _mockWarp(0, oracleMsg, true);
        adapter.verifyMessage(_buildICMMsg(0, oracleMsg));

        // Second delivery: different warp index, same global nonce.
        _mockWarp(1, oracleMsg, true);
        vm.expectRevert(
            abi.encodeWithSelector(OracleAdapter.AlreadyProcessed.selector, oracleMsg.nonce)
        );
        adapter.verifyMessage(_buildICMMsg(1, oracleMsg));
    }

    // -------------------------------------------------------------------------
    // sendMessage
    // -------------------------------------------------------------------------

    function test_sendMessage_callsWarpPrecompile() public {
        TeleporterMessageV2 memory teleporterMsg = _emptyTeleporterMsg();
        bytes memory expectedPayload = abi.encode(teleporterMsg);

        vm.mockCall(
            WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (expectedPayload)),
            abi.encode(bytes32(0))
        );
        vm.expectCall(
            WARP_PRECOMPILE, abi.encodeCall(IWarpMessenger.sendWarpMessage, (expectedPayload))
        );

        adapter.sendMessage(teleporterMsg);
    }

    // -------------------------------------------------------------------------
    // Admin — onlyOwner guards
    // -------------------------------------------------------------------------

    function test_setAllowedSource_onlyOwner() public {
        vm.prank(address(0xBEEF));
        vm.expectRevert(OracleAdapter.Unauthorized.selector);
        adapter.setAllowedSource("solana", "addr", true);
    }

    function test_transferOwnership_onlyOwner() public {
        vm.prank(address(0xBEEF));
        vm.expectRevert(OracleAdapter.Unauthorized.selector);
        adapter.transferOwnership(address(0xDEAD));
    }

    function test_transferOwnership_zeroAddress() public {
        vm.expectRevert(OracleAdapter.ZeroAddress.selector);
        adapter.transferOwnership(address(0));
    }

    // -------------------------------------------------------------------------
    // Internal helpers
    // -------------------------------------------------------------------------

    function _defaultOracleMsg() internal view returns (OracleMessage memory) {
        return OracleMessage({
            sourceType: SOURCE_TYPE,
            sourceAddress: SOURCE_ADDR,
            destContract: address(receiver),
            sourceBlockHeight: 1_000_000,
            nonce: 1,
            payload: hex"cafebabe"
        });
    }

    function _oracleMsgPayload(OracleMessage memory oracleMsg) internal pure returns (bytes memory) {
        return abi.encode(
            oracleMsg.sourceType,
            oracleMsg.sourceAddress,
            oracleMsg.destContract,
            oracleMsg.sourceBlockHeight,
            oracleMsg.nonce,
            oracleMsg.payload
        );
    }

    function _mockWarp(uint32 warpIndex, OracleMessage memory oracleMsg, bool valid) internal {
        WarpMessage memory warpMsg = WarpMessage({
            sourceChainID: THIS_CHAIN_ID,
            originSenderAddress: address(0),
            payload: _oracleMsgPayload(oracleMsg)
        });
        vm.mockCall(
            WARP_PRECOMPILE,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (warpIndex)),
            abi.encode(warpMsg, valid)
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
                destinationBlockchainID: THIS_CHAIN_ID,
                destinationAddress: oracleMsg.destContract,
                requiredGasLimit: 0,
                allowedRelayerAddresses: relayers,
                receipts: receipts,
                message: hex""
            }),
            sourceNetworkID: 1,
            sourceBlockchainID: THIS_CHAIN_ID,
            attestation: abi.encode(warpIndex, oracleMsg)
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

    function receiveTeleporterMessage(
        bytes32,
        address,
        bytes calldata message
    ) external override {
        lastMessage = message;
    }
}
