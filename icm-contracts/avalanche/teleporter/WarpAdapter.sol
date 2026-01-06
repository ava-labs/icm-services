// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {WarpMessage, IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {
    IMessageSender,
    IMessageVerifier,
    TeleporterMessage,
    ICMMessage
} from "./ITeleporterMessenger.sol";

/**
 * @dev Implementation of the {IMessageSender, IMessageVerifier} interfaces.
 *
 * This implementation is used to send messages cross chain using the WarpMessenger precompile,
 * and to receive messages sent from other chains. Teleporter contracts should be deployed through Nick's method
 * of universal deployer, such that the same contract is deployed at the same address on all chains.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract WarpAdapter is IMessageSender, IMessageVerifier {
    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    // This function signature can be changed to accept bytes to make it more generic, but I think
    // having the TeleporterMessage struct is more clear for now.
    // We should also consider having the warp message just be the hash of the teleporter message to save gas.
    // As-is, we need to pass the full teleporter message in the transaction predicate, and the transaction data.
    function sendMessage(TeleporterMessage calldata message) external override {
        // Submit the message to the AWM precompile.
        WARP_MESSENGER.sendWarpMessage(abi.encode(message));
    }

    // We can talk about whether we want to revert or return false on invalid messages. For now, I'm
    // just going to do either.
    // We should pass in the hash of the unsigned message instead of the full message to save gas. I'm
    // going to leave the type as is for now for clarity of what exactly is being verified.
    function verifyMessage(ICMMessage calldata message) external view override returns (bool) {
        uint32 messageIndex =  abi.decode(message.attestation, (uint32));

        // Verify and parse the cross chain message included in the transaction access list
        // using the warp message precompile.
        (WarpMessage memory warpMessage, bool success) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(success, "WarpAdapter: invalid warp message");

        // Only allow for messages to be received from the same address as this WarpAdapter contract.
        // The contract should be deployed using the universal deployer pattern, such that it knows messages
        // received from the same address on other chains were constructed using the same bytecode of this contract.
        // This allows for trusting the message format and uniqueness as specified by sendCrossChainMessage.
        require(
            warpMessage.originSenderAddress == address(this),
            "WarpAdapter: invalid origin sender address"
        );

        // Verify that the message was sent by the blockchain stated in the message.
        require(
            message.sourceBlockchainID == warpMessage.sourceChainID,
            "WarpAdapter: invalid source blockchain ID"
        );

        bytes32 warpHash = keccak256(abi.encode(warpMessage.payload));
        bytes32 inputHash = keccak256(abi.encode(message.unsignedMessage));

        return (warpHash == inputHash);
    }
}