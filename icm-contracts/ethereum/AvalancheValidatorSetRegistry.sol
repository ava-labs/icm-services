// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IAvalancheValidatorSetRegistry} from "./interfaces/IAvalancheValidatorSetRegistry.sol";

contract AvalancheValidatorSetRegistry is IAvalancheValidatorSetRegistry {
    uint32 public immutable avalancheNetworkID;

    constructor(
        uint32 avalancheNetworkID_
    ) {
        avalancheNetworkID = avalancheNetworkID_;
    }

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() external view returns (uint32) {
        return avalancheNetworkID;
    }
}
