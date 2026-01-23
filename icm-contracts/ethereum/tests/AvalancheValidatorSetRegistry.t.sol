// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {AvalancheValidatorSetRegistry} from "../AvalancheValidatorSetRegistry.sol";
import {Validator, ValidatorSet, ValidatorSets, ValidatorSetMetadata} from "../utils/ValidatorSets.sol";
import {BLST} from "../utils/BLST.sol";

contract AvalancheValidatorSetRegistryTest is Test {
    uint32 private constant _NETWORK_ID = 1;

    AvalancheValidatorSetRegistry private _registry;

    function setUp() public {

        (ValidatorSet memory validatorSet, bytes32 validatorSetHash) = _dummyPChainValidatorSet();
        bytes32[] memory shardHashes = new bytes32[](1);
        shardHashes[0] = validatorSetHash;
        ValidatorSetMetadata memory initialValidatorSetData = ValidatorSetMetadata({
            avalancheBlockchainID: validatorSet.avalancheBlockchainID,
            pChainHeight: validatorSet.pChainHeight,
            pChainTimestamp: validatorSet.pChainTimestamp,
            validatorSetHash: validatorSetHash,
            totalValidators: 5,
            shardHashes: shardHashes
        });
        _registry = new AvalancheValidatorSetRegistry(_NETWORK_ID, initialValidatorSetData);
    }

    function testGetAvalancheNetworkID() public view {
        assertEq(_registry.getAvalancheNetworkID(), _NETWORK_ID);
    }

    /*
     * @dev Create a dummy set of P-chain validators to initialize the `AvalancheValidatorSetRegistry` with.
     * @return Returns the validator set and hash of the validators
     */
    function _dummyPChainValidatorSet() private view returns (ValidatorSet memory validatorSet, bytes32 ) {
        Validator[] memory validators = new Validator[](5);
        uint64 totalWeight = 0;
        bytes memory previousPublicKey = new bytes(BLST.BLS_UNCOMPRESSED_PUBLIC_KEY_INPUT_LENGTH);
        uint8[5] memory secretKeys = [2, 3, 4, 5, 1];
        for (uint256 i = 0; i < 5; i++) {
            validators[i] = Validator({
                blsPublicKey: BLST.getPublicKeyFromSecret(secretKeys[i]),
                weight: uint64(i + 1)
            });
            // check that the validators are ordered by public key
            assertEq(BLST.comparePublicKeys(
                BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey),
                previousPublicKey
            ), 1);
            previousPublicKey = BLST.unPadUncompressedBlsPublicKey(validators[i].blsPublicKey);
            totalWeight += uint64(i + 1);
        }

        validatorSet = ValidatorSet ({
            avalancheBlockchainID: bytes32(0),
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: 0,
            pChainTimestamp: 0
        });
        return (validatorSet, sha256(ValidatorSets.serializeValidators(validators)));
    }
}
