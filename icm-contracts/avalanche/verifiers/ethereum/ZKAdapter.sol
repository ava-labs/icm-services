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
    /// @notice topic0 of the TeleporterV2MessageSent event which is the only log topic verifyMessage accepts
    bytes32 private constant _MESSAGE_SENT_TOPIC = keccak256("TeleporterV2MessageSent(bytes)");

    /// @notice The network ID (mainnet/testnet/local) messages must originate from
    uint32 public immutable sourceNetworkId;

    /// @notice Address of the trusted source chain adapter that emits TeleporterV2MessageSent logs
    address public immutable sourceEmitter;

    event TeleporterV2MessageSent(bytes encodedMessage);

    constructor(
        uint256 sourceChainId_,
        uint32 sourceNetworkId_,
        address sourceEmitter_,
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
        sourceEmitter = sourceEmitter_;
    }

    /**
     * @notice Originates a message by emitting it as a log for the relayer to prove.
     */
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        require(msg.sender == message.originTeleporterAddress, "unauthorized sender");
        emit TeleporterV2MessageSent(abi.encode(message));
    }

    /**
     * @notice Verifies a Teleporter message was emitted on the source chain by validating the attestation.
     * @dev Reverts unless the attestation proves that a `TeleporterV2MessageSent` log was emitted by the
     * trusted source-side adapter (`sourceEmitter`) whose payload equals the message being verified.
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        Attestation memory att = abi.decode(message.attestation, (Attestation));
        require(message.sourceBlockchainID == bytes32(sourceChainId), "bad source chain");
        require(message.sourceNetworkID == sourceNetworkId, "bad network");

        require(att.logProof.expectedEmitter == sourceEmitter, "bad emitter");
        require(att.logProof.expectedTopic0 == _MESSAGE_SENT_TOPIC, "bad topic");

        bytes memory logData = this.verifyLogAndExtract(att.execProof, att.logProof);
        bytes memory emitted = abi.decode(logData, (bytes));
        require(keccak256(emitted) == keccak256(abi.encode(message.message)), "payload mismatch");

        return true;
    }
}
