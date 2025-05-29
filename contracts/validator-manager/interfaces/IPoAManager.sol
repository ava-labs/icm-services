// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./IACP99Manager.sol";

/**
 * @dev Proof-of-authority manager interface meant for to be used as the owner of an
 * {IValidatorManager} contract instance that only allows modifications to the validator set
 * to be initiated by a fixed address. Once initiated, modifications are meant to be able to be
 * completed by any address to allow for imrpoved UX flows.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IPoAManager {
    /**
     * @notice Initiates validator registration. Only callable by the contract owner.
     * @dev Delegates to IValidatorManager.initiateValidatorRegistration.
     * @param nodeID The node ID of the validator.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The P-chain owner for remaining balance.
     * @param disableOwner The P-chain owner for disabling.
     * @param weight The validator's weight.
     * @return validationID The unique identifier for the validator registration.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external returns (bytes32);

    /**
     * @notice Initiates validator removal. Only callable by the contract owner.
     * @dev Delegates to IValidatorManager.initiateValidatorRemoval.
     * @param validationID The unique identifier of the validator to remove.
     */
    function initiateValidatorRemoval(
        bytes32 validationID
    ) external;

    /**
     * @notice Initiates a validator weight update. Only callable by the contract owner.
     * @dev Delegates to IValidatorManager.initiateValidatorWeightUpdate.
     * @param validationID The unique identifier of the validator.
     * @param newWeight The new weight for the validator.
     * @return messageIndex The index of the message for the update.
     * @return updateID The unique identifier for the weight update.
     */
    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external returns (uint64, bytes32);

    /**
     * @notice Completes validator registration. Callable by anyone.
     * @dev Delegates to IValidatorManager.completeValidatorRegistration.
     * @param messageIndex The index of the registration message.
     * @return validationID The unique identifier for the completed registration.
     */
    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32);

    /**
     * @notice Completes validator removal. Callable by anyone.
     * @dev Delegates to IValidatorManager.completeValidatorRemoval.
     * @param messageIndex The index of the removal message.
     * @return validationID The unique identifier for the completed removal.
     */
    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32);

    /**
     * @notice Completes a validator weight update. Callable by anyone.
     * @dev Delegates to IValidatorManager.completeValidatorWeightUpdate.
     * @param messageIndex The index of the weight update message.
     * @return updateID The unique identifier for the completed weight update.
     * @return newWeight The new weight for the validator.
     */
    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) external returns (bytes32, uint64);

    /**
     * @notice Transfers ownership of the underlying validator manager contract.
     * @dev Only callable by the current owner of this contract.
     * @param newOwner The address to transfer ownership to.
     */
    function transferUnderlyingValidatorManagerOwnership(address newOwner) external;
}
