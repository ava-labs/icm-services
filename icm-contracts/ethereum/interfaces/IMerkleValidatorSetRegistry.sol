// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {ICMMessage} from "../../common/ICM.sol";
import {ValidatorSetMerkleCommitment} from "../utils/ValidatorSets.sol";

/**
 * @title IMerkleValidatorSetRegistry
 * @notice Interface for a Merkle-committed Avalanche Validator Registry contract
 * @dev This interface defines the events and functions for managing Avalanche validator set
 * commitments and verifying signed messages against them. Validator data lives off-chain;
 * the contract retains only a Merkle root over the set, and signers supply Merkle inclusion
 * proofs alongside their attestations.
 */
interface IMerkleValidatorSetRegistry {
    event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID);

    /**
     * @notice Installs a new validator set commitment for the Avalanche blockchain ID
     * named in the payload. The new commitment fits in a single transaction at any
     * validator-set size, so no sharding or multi-transaction coordination is required.
     *
     * Emits a `ValidatorSetUpdated` event upon successful update.
     * @dev A validator set update can be submitted by anyone, but it must be signed by the
     * current P-chain validator set known to this contract if no previous registration to the
     * same blockchain ID exists. Otherwise, it must be signed by the latest validator set
     * registered to the blockchain ID so that the P-chain validator sets have no elevated
     * permission after registration. On updates, the new P-chain height and timestamp must
     * both strictly exceed the currently stored values.
     * @param message The ICM message containing the new commitment payload and a
     * Merkle attestation signed by the appropriate authorizing set. See
     * `ValidatorSetCommitment` and `MerkleAttestation` for further details.
     */
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
