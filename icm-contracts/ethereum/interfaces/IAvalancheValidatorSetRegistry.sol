// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../../common/ICM.sol";
import {ValidatorSetShard} from "../utils/ValidatorSets.sol";

/**
 * @title IAvalancheValidatorSetRegistry
 * @notice Interface for an Avalanche Validator Registry contract
 * @dev This interface defines the events and functions for managing Avalanche validators and validating signed message
 * from a given validator set.
 */
interface IAvalancheValidatorSetRegistry {
    event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID);

    event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID);

    /**
     * @notice Registers a new validator set for a specific Avalanche blockchain ID. A registered
     * set may require several txs to populate due to size/gas constraints. This is done by calls to
     * `updateValidatorSet` as necessary.
     *
     * Emits a `ValidatorSetRegistered` event upon successful registration.
     * @dev A validator set can be registered by anyone, but it must be signed by the current P-chain validator set
     * known to this contract if no previous registration to the same blockchain ID exists. Otherwise, it must be
     * signed by the latest validator set registered to the blockchain ID so that the P-chain validator sets have no
     * elevated permission after registration.
     * @param message The ICM message containing the metadata about the validator set to register. See
     * `ValidatorSetMetadata` for further details
     * @param shardBytes The first shard of the data needed to populate the newly registered validator set.
     */
    function registerValidatorSet(ICMMessage calldata message, bytes memory shardBytes) external;

    /**
     * @notice Apply a shard to a registered validator set which is only partially populated.
     *
     * Emits a `ValidatorSetUpdated` event if a partial set becomes fully populated.
     * @param shard indicates the shard number and blockchain ID of the partial
     * set that is being updated.
     * @param shardBytes The next shard of data needed to populate a registered partial
     * validator set
     */
    function updateValidatorSet(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) external;
}
