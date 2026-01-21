// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/**
 * @title IAvalancheValidatorSetRegistry
 * @notice Interface for an Avalanche Validator Registry contract
 * @dev This interface defines the events and functions for managing Avalanche validators and validating signed message
 * from a given validator set.
 */
interface IAvalancheValidatorSetRegistry {
    event ValidatorSetRegistered(
        uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID
    );

    event ValidatorSetUpdated(
        uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID
    );
}
