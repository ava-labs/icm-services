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
    event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID);

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
     * the currently registered validator set for that blockchain ID.
     * @param message The ICM message containing the new validator set commitment.
     * The signed warp preimage uses `message.sourceBlockchainID` and `address(0)` as the origin
     * sender, matching how the P-chain Warp precompile emits these messages.
     */
    function registerValidatorSet(
        ICMMessage calldata message
    ) external;

    /**
     * @notice Replace the registered Merkle commitment for a validator set with a new one
     * signed by that set's current validators.
     *
     * Emits a `ValidatorSetUpdated` event once the new commitment is stored.
     * @param message The ICM message containing the new validator set commitment. Must be
     * signed by the currently registered validator set for the blockchain ID declared in
     * the payload; reverts if the chain is not yet registered (use `registerValidatorSet`
     * for first registrations) or if the attestation fails to verify.
     */
    function updateValidatorSet(
        ICMMessage calldata message
    ) external;

    /**
     * @notice Retrieves the current validator set commitment registered for a given
     * Avalanche blockchain ID.
     * @param avalancheBlockchainID The Avalanche blockchain ID to query the commitment for.
     */
    function getValidatorSetCommitment(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory);
}
