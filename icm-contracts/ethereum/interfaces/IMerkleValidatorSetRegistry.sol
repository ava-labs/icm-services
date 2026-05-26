// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {ICMMessage} from "../../common/ICM.sol";
import {ValidatorSetMerkleCommitment} from "../utils/ValidatorSets.sol";

/**
 * @title IMerkleValidatorSetRegistry
 * @notice Interface for a Merkle-committed Avalanche Validator Registry contract
 * @dev Interface for managing Avalanche validator set commitments and verifying
 * signed messages against them. The full validator set is stored off-chain.
 * Each set is represented on-chain by a single Merkle root commitment over its
 * validators, against which messages are verified using calldata-supplied
 * signers and multi-proofs.
 */
interface IMerkleValidatorSetRegistry {
    event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID);

    /**
     * @notice Registers or updates the Merkle commitment for a validator set keyed by Avalanche
     * blockchain ID.
     *
     * Emits a `ValidatorSetRegistered` event once the new commitment is stored.
     * @dev Unlike the storage-based registry, the entire commitment (root, total weight, P-chain height and timestamp)
     * fits in a single message so no sharding or partial state is needed.
     * The same function handles both first registration and subsequent updates and which case applies
     * is determined by whether the payload's blockchain ID is already registered. The first registration for a
     * given blockchain ID must be signed by the P-chain validator set, while subsequent updates must be signed by
     * the currently registered validator set for that blockchain ID, or the P-chain validator set as a fallback.
     * @param message The ICM message containing the new validator set commitment.
     * The signed warp preimage uses `message.sourceBlockchainID` and `address(0)` as the origin
     * sender, matching how the P-chain Warp precompile emits these messages.
     * @param signingChainID The Avalanche blockchain ID of the validator set that signed the message. This can either
     * be the same as the payload's blockchain ID for updates, or the P-chain ID for first registrations or further updates.
     */
    function registerValidatorSet(ICMMessage calldata message, bytes32 signingChainID) external;

    /**
     * @notice Retrieves the current validator set commitment registered for a given
     * Avalanche blockchain ID.
     * @param avalancheBlockchainID The Avalanche blockchain ID to query the commitment for.
     */
    function getValidatorSetCommitment(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory);
}
