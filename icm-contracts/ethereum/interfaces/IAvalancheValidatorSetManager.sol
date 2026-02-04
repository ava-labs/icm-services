// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../../common/ICM.sol";
import {
    PartialValidatorSet,
    Validator,
    ValidatorSetShard,
    ValidatorSetMetadata,
    ValidatorSet
} from "../utils/ValidatorSets.sol";

/**
 * @title IAvalancheValidatorManager
 * @notice Interface for an Avalanche Validator Manager contract
 * @dev This interface defines the data layer for managing the validator sets and updating them
 */
interface IAvalancheValidatorSetManager {
    /**
     * @notice This contract controls data that is owned by another contract. We
     * specify the owner to indicate who is authorized to change the data.
     * @dev We want to support deploying an instance of this contract first
     * and then specifying the owner later.
     */
    function initialize(
        address owner
    ) external;

    /**
     * @notice Set the **complete** validator set most recently registered
     */
    function setValidatorSet(
        bytes32 avalancheBlockchainID,
        ValidatorSet memory validatorSet
    ) external;

    /**
     * @notice Set a partial validator set as most recently registered
     */
    function setPartialValidatorSet(
        bytes32 avalancheBlockchainID,
        PartialValidatorSet memory partialValidatorSet
    ) external;

    /**
     * @notice  Validate and apply a shard to a partial validator set. If the set is completed by this shard, copy
     * it over to the `_validatorSets` mapping.
     * @param shard Indicates the sequence number of the shard and blockchain affected by this update
     * @param shardBytes the actual data of the shard which
     */
    function applyShard(ValidatorSetShard calldata shard, bytes memory shardBytes) external;

    /**
     * @notice Parses and validates metadata about a validator set data from an ICM message. This
     * is called when registering validator sets. It may also contain a (potentially partial) subset of
     * the validators that are being registered. This is always considered to be the first shard of
     * the requisite data.
     *
     * @param icmMessage The ICM message containing the validator set metadata
     * @param shardBytes The serialized data used to construct a subset of the registered
     * validator set
     * @return The parsed validator set metadata
     * @return A parsed validators array
     * @return The total weight of the parsed validators
     */
    function parseValidatorSetMetadata(
        ICMMessage calldata icmMessage,
        bytes calldata shardBytes
    ) external view returns (ValidatorSetMetadata memory, Validator[] memory, uint64);

    /**
     * @notice Check if a **complete** validator set is registered (not just a partial).
     */
    function isRegistered(
        bytes32 avalancheBlockchainID
    ) external view returns (bool);

    /**
     * @notice Check if a validator set is registered but awaiting further updates.
     */
    function isRegistrationInProgress(
        bytes32 avalancheBlockchainID
    ) external view returns (bool);

    /**
     * @notice Get the number of shards received for the most recent validator set registered
     * to this blockchain ID
     */
    function getShardsReceived(
        bytes32 avalancheBlockchainID
    ) external view returns (uint64);

    /**
     * @notice Get the hash of the shard committed to at the requested index
     */
    function getShardHash(
        bytes32 avalancheBlockchainID,
        uint256 index
    ) external view returns (bytes32);

    /**
     * @notice Get the **complete** validator set most recently registered
     */
    function getValidatorSet(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSet memory);
}
