// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {AvalancheValidatorSetRegistry} from "../AvalancheValidatorSetRegistry.sol";

contract AvalancheValidatorSetRegistryTest is Test {
    uint32 private constant _NETWORK_ID = 1;

    AvalancheValidatorSetRegistry private _registry;

    function setUp() public {
        _registry = new AvalancheValidatorSetRegistry(_NETWORK_ID);
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(_registry.getAvalancheNetworkID(), _NETWORK_ID);
    }
}
