// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {IAdapter} from "./ITeleporterMessengerV2.sol";
import {TeleporterICMMessage, TeleporterMessageV2} from "./TeleporterMessageV2.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// A contract that allows a single contract to be used for both directions of
// communication between two chains even though the logic needed for each direction
// may be different. This is done by determining which blockchain it is on and then
// delegating the logic to an appropriate contract.
contract Adapter is IAdapter {
    bytes32 public immutable chain1;
    bytes32 public immutable chain2;
    IAdapter public immutable adapter1;
    IAdapter public immutable adapter2;

    constructor(bytes32 chain1_, bytes32 chain2_, address adapter1_, address adapter2_) {
        chain1 = chain1_;
        chain2 = chain2_;
        adapter1 = IAdapter(adapter1_);
        adapter2 = IAdapter(adapter2_);
    }

    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        if (message.message.destinationBlockchainID == chain1) {
            return adapter1.verifyMessage(message);
        } else if (message.message.destinationBlockchainID == chain2) {
            return adapter2.verifyMessage(message);
        } else {
            revert("Unexpected blockchain ID");
        }
    }

    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        if (message.destinationBlockchainID == chain1) {
            return adapter1.sendMessage(message);
        } else if (message.destinationBlockchainID == chain2) {
            return adapter2.sendMessage(message);
        } else {
            revert("Unexpected blockchain ID");
        }
    }
}
