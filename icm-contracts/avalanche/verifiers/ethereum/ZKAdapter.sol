// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {ZKStateManager} from "./ZKStateManager.sol";
import {
    TeleporterICMMessage, TeleporterMessageV2
} from "../../../common/ITeleporterMessengerV2.sol";
import {Consensus, Execution, Receipt} from "./StateManagerLibrary.sol";
import {IAdapter} from "../../../common/ITeleporterMessengerV2.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
struct Attestation {
    Execution.Proof execProof;
    Receipt.Proof logProof;
}

/**
 * @notice Adapts ZKStateManager for the Teleporter messaging protocol.
 */
contract ZKAdapter is ZKStateManager, IAdapter {
    bytes32 private constant _MESSAGE_SENT_TOPIC = keccak256("TeleporterMessageSent(bytes)");

    uint32 public immutable sourceNetworkId;

    event TeleporterMessageSent(bytes encodedMessage);

    constructor(
        uint256 sourceChainId_,
        uint32 sourceNetworkId_,
        Consensus.State memory startingState,
        Execution.BeaconConfig memory beaconConfig_,
        uint24 permissibleTimespan_,
        address verifier,
        bytes32 imageID_,
        address admin,
        address superAdmin
    )
        ZKStateManager(
            sourceChainId_,
            startingState,
            beaconConfig_,
            permissibleTimespan_,
            verifier,
            imageID_,
            admin,
            superAdmin
        )
    {
        sourceNetworkId = sourceNetworkId_;
    }

    /**
     * @notice Originates a message by emitting it as a log for the relayer to prove.
     */
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        require(msg.sender == message.originTeleporterAddress, "unauthorized sender");
        emit TeleporterMessageSent(abi.encode(message));
    }

    /**
     * @notice Verifies a Teleporter message was emitted on the source chain by validating the attestation.
     * @dev Reverts unless the attestation proves that a `TeleporterMessageSent` log was emitted by a
     * trusted source-side adapter whose payload equals the message being verified. The source-side adapter
     * has the same contract address as the destination-side adapter, since both are deployed via Nick's method.
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        Attestation memory att = abi.decode(message.attestation, (Attestation));
        require(message.sourceBlockchainID == bytes32(sourceChainId), "bad source chain");
        require(message.sourceNetworkID == sourceNetworkId, "bad network");

        require(att.logProof.expectedEmitter == address(this), "bad emitter");
        require(att.logProof.expectedTopic0 == _MESSAGE_SENT_TOPIC, "bad topic");

        bytes memory logData = this.verifyLogAndExtract(att.execProof, att.logProof);
        bytes memory emitted = abi.decode(logData, (bytes));
        require(keccak256(emitted) == keccak256(abi.encode(message.message)), "payload mismatch");

        return true;
    }
}
