// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {WarpAdapter} from "../WarpAdapter.sol";
import {IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {TeleporterMessageV2} from "@common/ITeleporterMessengerV2.sol";

/// @dev Minimal stand-in for a TeleporterMessengerV2: exposes the adapter it sends through.
contract WarpAdapterTestTeleporter {
    address public immutable messageSender;

    constructor(
        address messageSender_
    ) {
        messageSender = messageSender_;
    }
}

/**
 * @dev sendMessage turns a caller-supplied message into a Warp message that verifyMessage on the
 * destination chain accepts as authentic, so only the messenger named in the message, or the adapter
 * that messenger sends through, may call it. Anyone else calling it directly would otherwise be able
 * to forge messages from any application.
 */
contract WarpAdapterSendMessageTest is Test {
    address private constant _WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address private constant _WRAPPER_ADAPTER = address(0xA11CE);
    address private constant _ATTACKER = address(0xBAD);

    WarpAdapter private _adapter;
    WarpAdapterTestTeleporter private _teleporter;

    function setUp() public {
        _adapter = new WarpAdapter();
        _teleporter = new WarpAdapterTestTeleporter(_WRAPPER_ADAPTER);
        // The Warp precompile does not exist in the test VM. Give it code so its calls can be mocked.
        vm.etch(_WARP_PRECOMPILE_ADDRESS, hex"00");
        vm.mockCall(
            _WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );
    }

    function testSendMessageFromOriginTeleporter() public {
        TeleporterMessageV2 memory message = _message();
        vm.expectCall(
            _WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(message)))
        );
        vm.prank(address(_teleporter));
        _adapter.sendMessage(message);
    }

    function testSendMessageFromOriginTeleporterAdapter() public {
        TeleporterMessageV2 memory message = _message();
        vm.expectCall(
            _WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(message)))
        );
        vm.prank(_WRAPPER_ADAPTER);
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsForUnauthorizedSender() public {
        TeleporterMessageV2 memory message = _message();
        vm.prank(_ATTACKER);
        vm.expectRevert("WarpAdapter: unauthorized sender");
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsWhenOriginTeleporterIsNotAContract() public {
        // Looking up the messenger's adapter fails on an address without code, which rejects the send.
        TeleporterMessageV2 memory message = _message();
        message.originTeleporterAddress = address(0xBEEF);
        vm.prank(_ATTACKER);
        vm.expectRevert();
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsForAdapterOfAnotherTeleporter() public {
        // The attacker's own messenger may name any adapter, but a message it originates carries the
        // attacker's messenger address, which the destination messenger rejects. Naming the real
        // messenger while calling from another adapter must fail here.
        WarpAdapterTestTeleporter attackerTeleporter = new WarpAdapterTestTeleporter(_ATTACKER);
        TeleporterMessageV2 memory message = _message();
        message.originTeleporterAddress = address(attackerTeleporter);
        vm.prank(_WRAPPER_ADAPTER);
        vm.expectRevert("WarpAdapter: unauthorized sender");
        _adapter.sendMessage(message);
    }

    function _message() internal view returns (TeleporterMessageV2 memory message) {
        message.originSenderAddress = address(0x5E11DE5);
        message.originTeleporterAddress = address(_teleporter);
        message.destinationBlockchainID = bytes32(uint256(2));
        message.message = hex"deadbeef";
    }
}
