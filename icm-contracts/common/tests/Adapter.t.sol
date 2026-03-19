// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {
    TeleporterMessageReceipt,
    TeleporterMessageInput,
    TeleporterMessageV2,
    TeleporterFeeInfo,
    IAdapter,
    TeleporterICMMessage
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";
import {Adapter} from "../Adapter.sol";

contract AdapterTest is Test {
    bytes32 private constant _BLOCKCHAIN1 = hex"01";
    bytes32 private constant _BLOCKCHAIN2 = hex"02";
    bytes32 private constant _BLOCKCHAIN3 = hex"03";
    TeleporterMessengerV2 private _teleporter1;
    TeleporterMessengerV2 private _teleporter2;
    TeleporterMessengerV2 private _teleporter3;
    Adapter private _adapter;
    IAdapter private _adapter1;
    IAdapter private _adapter2;

    function setUp() public {
        _adapter1 = new AtoB();
        _adapter2 = new BtoA();
        _adapter = new Adapter(_BLOCKCHAIN1, _BLOCKCHAIN2, address(_adapter1), address(_adapter2));
        _teleporter1 = new TeleporterMessengerV2(address(_adapter));
        _teleporter1.initialize(_BLOCKCHAIN1);
        _teleporter2 = new TeleporterMessengerV2(address(_adapter));
        _teleporter2.initialize(_BLOCKCHAIN2);
        _teleporter3 = new TeleporterMessengerV2(address(_adapter));
        _teleporter3.initialize(_BLOCKCHAIN3);
    }

    /**
     * @dev Test that a different adapter is used to send a message depending on the
     * blockchain ID field of the calling TeleporterMessengerV2 contract
     */
    function testSending() public {
        TeleporterMessageInput memory input1 = TeleporterMessageInput({
            destinationBlockchainID: _BLOCKCHAIN2,
            destinationAddress: address(_adapter2),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
        vm.expectRevert("Adapter 1");
        _teleporter1.sendCrossChainMessage(input1);

        TeleporterMessageInput memory input2 = TeleporterMessageInput({
            destinationBlockchainID: _BLOCKCHAIN1,
            destinationAddress: address(_adapter1),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
        vm.expectRevert("Adapter 2");
        _teleporter2.sendCrossChainMessage(input2);

        TeleporterMessageInput memory input3 = TeleporterMessageInput({
            destinationBlockchainID: _BLOCKCHAIN3,
            destinationAddress: address(_adapter2),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
        vm.expectRevert("Unexpected blockchain ID");
        _teleporter3.sendCrossChainMessage(input3);
    }

    /**
     * @dev Test that a different adapter is used to receive a message depending on the
     * blockchain ID field of the calling TeleporterMessengerV2 contract
     */
    function testReceiving() public {
        TeleporterICMMessage memory message1 = TeleporterICMMessage({
            message: TeleporterMessageV2({
                messageNonce: 0,
                originSenderAddress: address(_adapter2),
                originTeleporterAddress: address(_teleporter2),
                destinationBlockchainID: _BLOCKCHAIN1,
                destinationAddress: address(_adapter1),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                receipts: new TeleporterMessageReceipt[](0),
                message: hex"deadbeef"
            }),
            sourceNetworkID: 1,
            sourceBlockchainID: _BLOCKCHAIN2,
            attestation: new bytes(0)
        });
        vm.expectRevert("Adapter 1");
        _teleporter1.receiveCrossChainMessage(message1, address(0));

        TeleporterICMMessage memory message2 = TeleporterICMMessage({
            message: TeleporterMessageV2({
                messageNonce: 0,
                originSenderAddress: address(_adapter1),
                originTeleporterAddress: address(_teleporter1),
                destinationBlockchainID: _BLOCKCHAIN2,
                destinationAddress: address(_adapter2),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                receipts: new TeleporterMessageReceipt[](0),
                message: hex"deadbeef"
            }),
            sourceNetworkID: 1,
            sourceBlockchainID: _BLOCKCHAIN1,
            attestation: new bytes(0)
        });
        vm.expectRevert("Adapter 2");
        _teleporter2.receiveCrossChainMessage(message2, address(0));

        TeleporterICMMessage memory message3 = TeleporterICMMessage({
            message: TeleporterMessageV2({
                messageNonce: 0,
                originSenderAddress: address(_adapter1),
                originTeleporterAddress: address(_teleporter3),
                destinationBlockchainID: _BLOCKCHAIN3,
                destinationAddress: address(_adapter2),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                receipts: new TeleporterMessageReceipt[](0),
                message: hex"deadbeef"
            }),
            sourceNetworkID: 1,
            sourceBlockchainID: _BLOCKCHAIN3,
            attestation: new bytes(0)
        });
        vm.expectRevert("Unexpected blockchain ID");
        _teleporter3.receiveCrossChainMessage(message3, address(0));
    }
}

contract AtoB is IAdapter {
    function sendMessage(
        /* solhint-disable-next-line no-unused-vars */
        TeleporterMessageV2 calldata message
    ) external {
        revert("Adapter 1");
    }

    function verifyMessage(
        /* solhint-disable-next-line no-unused-vars */
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        revert("Adapter 1");
    }
}

contract BtoA is IAdapter {
    function sendMessage(
        /* solhint-disable-next-line no-unused-vars */
        TeleporterMessageV2 calldata message
    ) external pure {
        revert("Adapter 2");
    }

    function verifyMessage(
        /* solhint-disable-next-line no-unused-vars */
        TeleporterICMMessage calldata message
    ) external pure returns (bool) {
        revert("Adapter 2");
    }
}
