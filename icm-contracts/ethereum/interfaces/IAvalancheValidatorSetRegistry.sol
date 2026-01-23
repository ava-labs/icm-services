// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../../common/ICM.sol";

/**
 * @title IAvalancheValidatorSetRegistry
 * @notice Interface for an Avalanche Validator Registry contract
 * @dev This interface defines the events and functions for managing Avalanche validators and validating signed message
 * from a given validator set.
 */
interface IAvalancheValidatorSetRegistry {
    event ValidatorSetRegistered(
        bytes32 indexed avalancheBlockchainID
    );

    event ValidatorSetUpdated(
        bytes32 indexed avalancheBlockchainID
    );

    /**
     * @notice Registers a new validator set
     * @dev A validator set must can be registered by anyone, but it must be signed by the current P-chain validator set
     * known to this contract. Updates to any validator set after registration are always authenticated by a
     * signed ICM message from the current validator set, giving P-chain validator sets no elevated permission after
     * registration.
     * @param message The ICM message containing the validator set to register.
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external;
}
