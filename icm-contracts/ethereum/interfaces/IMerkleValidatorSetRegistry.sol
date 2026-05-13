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
    event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID);

    /// TODO: Implement in follow-up work.
    function registerValidatorSet(ICMMessage calldata message, bytes memory shardBytes) external;

    /// TODO: Implement in follow-up work.
    function applyValidatorSetUpdate(
        ICMMessage calldata message
    ) external;

    /**
     * @notice Retrieves the current validator set commitment registered for a given
     * Avalanche blockchain ID.
     * @param avalancheBlockchainID The Avalanche blockchain ID to query the commitment for.
     */
    function getValidatorSet(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory);
}
