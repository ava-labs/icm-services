// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {IMerkleValidatorSetRegistry} from "./interfaces/IMerkleValidatorSetRegistry.sol";
import {ICMMessage} from "@common/ICM.sol";
import {ValidatorSetMerkleCommitment, ValidatorSets} from "./utils/ValidatorSets.sol";
import {
    TeleporterMessageV2Parsing,
    TeleporterICMMessage,
    TeleporterMessageV2
} from "@common/TeleporterMessageV2.sol";
import {IAdapter} from "@common/ITeleporterMessengerV2.sol";
import {IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// A contract for managing Avalanche validator sets which can be used to verify ICM messages
// via a quorum of signatures.
//
// Validator sets are anchored on-chain by a Merkle root commitment over the set, rather
// than by storing the full validator set. Updates and message verification supply signing
// validators in calldata along with a Merkle multi-inclusion proof against the stored root,
// trading per-validator storage costs for per-message calldata and proof verification work.
contract MerkleValidatorSetRegistry is IMerkleValidatorSetRegistry, IAdapter {
    address private constant _WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    uint32 public immutable avalancheNetworkID;
    // The Avalanche blockchain ID of the P-chain
    bytes32 public immutable pChainID;
    // Whether the P-Chain may sign off on updates to already registered L1s
    bool public immutable allowPChainFallback;
    // Mapping of Avalanche blockchain IDs to their validator set commitments.
    mapping(bytes32 => ValidatorSetMerkleCommitment) internal _valSetCommitments;

    // Constructs a new registry instance with the initial validator set commitment registered on the P-chain.
    constructor(
        uint32 avalancheNetworkID_,
        bytes32 pChainID_,
        bytes32 pChainGenesisRoot,
        uint64 pChainTotalWeight,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        bool allowPChainFallback_
    ) {
        avalancheNetworkID = avalancheNetworkID_;
        pChainID = pChainID_;
        allowPChainFallback = allowPChainFallback_;
        _valSetCommitments[pChainID_] = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: pChainID_,
            root: pChainGenesisRoot,
            totalWeight: pChainTotalWeight,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
    }

    /// @dev sendMessage has no msg.sender == originTeleporterAddress check. This check must live in
    /// the contract that the messenger (Teleporter) calls directly, currently in Adapter.sol,
    /// in order to prevent unauthorized calls to sendMessage. To clarify, this contract cannot be used as the
    /// direct messaging entry-point.
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        IWarpMessenger(_WARP_PRECOMPILE_ADDRESS)
            .sendWarpMessage(TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message));
    }

    /**
     * @notice Registers or updates the Merkle commitment for a validator set keyed by Avalanche
     * blockchain ID.
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes32 signingChainID
    ) external {
        require(pChainInitialized(), "No P-chain validator set registered.");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        ValidatorSetMerkleCommitment memory newCommitment =
            ValidatorSets.parseMerkleCommitment(message.rawMessage);
        bytes32 payloadBlockchainID = newCommitment.avalancheBlockchainID;

        // Initial registration must always be signed by the P-Chain.
        // For subsequent updates, the signing authority must be either the target chain itself if already registered
        // or the P-Chain if allowPChainFalback is enabled
        bool selfSigned = (signingChainID == payloadBlockchainID);
        bool pChainSigned = (signingChainID == pChainID);

        if (!isRegistered(payloadBlockchainID)) {
            require(pChainSigned, "Initial registration must be signed by the P-Chain");
        } else {
            require(selfSigned || (pChainSigned && allowPChainFallback), "Invalid signing chain");
        }

        verifyICMMessage(message, signingChainID);

        _valSetCommitments[payloadBlockchainID] = newCommitment;
        emit ValidatorSetRegistered(payloadBlockchainID);
    }

    /**
     * @notice Verifies that a TeleporterICMMessage was signed by the validator set committed
     * to under the Merkle root registered for `message.sourceBlockchainID`.
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external view returns (bool) {
        require(pChainInitialized(), "No P-chain validator set registered.");
        require(isRegistered(message.sourceBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            message.sourceNetworkID,
            message.sourceBlockchainID,
            address(this),
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message.message)
        );
        return ValidatorSets.verifyMerkleAttestation(
            message.attestation, signedData, _valSetCommitments[message.sourceBlockchainID]
        );
    }

    /**
     * @notice Returns the current validator set commitment registered for the given Avalanche blockchain ID.
     */
    function getValidatorSetCommitment(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory) {
        return _valSetCommitments[avalancheBlockchainID];
    }

    /**
     * @notice Verify the message. Does the following checks:
     *   1. P-chain root of trust has been initialized.
     *   2. The given chain has a registered validator set commitment.
     *   3. The message is intended for this network.
     *   4. The attestation is signed by a quorum of validators committed to under
     * `avalancheBlockchainID`'s stored Merkle root.
     * Reverts if any check fails.
     *
     * @dev Used to verify validator set update messages and called by `registerValidatorSet`.
     * These messages differ from Teleporter application messages in that they are emitted
     * directly by the P-chain rather than from a contract invoking `sendWarpMessage`,
     * so the warp preimage's `originSenderAddress` is `address(0)` instead of
     * `address(this)`.
     */
    function verifyICMMessage(
        ICMMessage calldata message,
        bytes32 avalancheBlockchainID
    ) public view {
        require(pChainInitialized(), "P-chain not initialized");
        require(isRegistered(avalancheBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            message.sourceNetworkID, message.sourceBlockchainID, address(0), message.rawMessage
        );
        require(
            ValidatorSets.verifyMerkleAttestation(
                message.attestation, signedData, _valSetCommitments[avalancheBlockchainID]
            ),
            "Failed to verify attestation"
        );
    }

    /**
     * @dev Check if a P-chain validator set has been completely registered.
     * Until it has, no functions other than adding to this validator set are
     * permitted.
     */
    function pChainInitialized() public view returns (bool) {
        return _valSetCommitments[pChainID].totalWeight != 0;
    }

    /**
     * @notice Check if a validator set is registered.
     */
    function isRegistered(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return _valSetCommitments[avalancheBlockchainID].totalWeight != 0;
    }
}
