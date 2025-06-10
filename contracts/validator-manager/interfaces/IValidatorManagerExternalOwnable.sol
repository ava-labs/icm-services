// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.25;

import {IValidatorManager} from "./IValidatorManager.sol";

interface IValidatorManagerExternalOwnable is IValidatorManager {
    /**
     * @notice Transfers ownership of the validator manager to a new owner.
     * @param newOwner The address of the new owner.
     * @dev This function can only be called by the current owner.
     */
    function transferOwnership(
        address newOwner
    ) external;
}