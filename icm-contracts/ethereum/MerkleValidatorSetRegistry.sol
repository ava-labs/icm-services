// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {IMerkleValidatorSetRegistry} from "./interfaces/IMerkleValidatorSetRegistry.sol";
import {ICMMessage} from "../common/ICM.sol";
import {ValidatorSetMerkleCommitment, ValidatorSets} from "./utils/ValidatorSets.sol";
import {
    TeleporterMessageV2Parsing,
    TeleporterICMMessage,
    TeleporterMessageV2
} from "../common/TeleporterMessageV2.sol";
import {IAdapter} from "../common/ITeleporterMessengerV2.sol";
import {IWarpMessenger} from "../avalanche/subnet-evm/IWarpMessenger.sol";

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
    // Mapping of Avalanche blockchain IDs to their validator set commitments.
    mapping(bytes32 => ValidatorSetMerkleCommitment) internal _valSetCommitments;
    // Constructs a new registry instance with the initial validator set commitment registered on the P-chain.

    constructor(
        uint32 avalancheNetworkID_,
        bytes32 pChainID_,
        bytes32 pChainGenesisRoot,
        uint64 pChainTotalWeight,
        uint64 pChainHeight,
        uint64 pChainTimestamp
    ) {
        avalancheNetworkID = avalancheNetworkID_;
        pChainID = pChainID_;
        _valSetCommitments[pChainID_] = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: pChainID_,
            root: pChainGenesisRoot,
            totalWeight: pChainTotalWeight,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
    }

    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        IWarpMessenger(_WARP_PRECOMPILE_ADDRESS).sendWarpMessage(
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message)
        );
    }

    /// TODO: Implement in follow-up work.
    function registerValidatorSet(
        ICMMessage calldata, /* message */
        bytes memory /* validatorSetBytes */
    ) external {
        revert("Not implemented");
    }

    /// TODO: Implement in follow-up work.
    function applyValidatorSetUpdate(
        ICMMessage calldata /* message */
    ) external {
        revert("Not implemented");
    }

    /// @notice Verifies a TeleporterICMMessage against the registered validator set commitment for message.sourceBlockchainID.
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external view returns (bool) {
        require(pChainInitialized(), "P-chain not initialized");
        require(isRegistered(message.sourceBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            message.sourceNetworkID,
            message.sourceBlockchainID,
            address(this),
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message.message)
        );
        ValidatorSetMerkleCommitment storage comm = _valSetCommitments[message.sourceBlockchainID];
        return ValidatorSets.verifyMerkleAttestation(message.attestation, signedData, comm);
    }

    /**
     * @notice Returns the current validator set commitment registered for the given Avalanche blockchain ID.
     */
    function getValidatorSet(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory) {
        return _valSetCommitments[avalancheBlockchainID];
    }

    /**
     * @dev Check if a P-chain validator set has been completely registered.
     * Until is has, no functions other than adding to this validator set are
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
