// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IAvalancheValidatorSetRegistry} from "./interfaces/IAvalancheValidatorSetRegistry.sol";
import {
    Validator,
    ValidatorSet,
    ValidatorSetStatePayload,
    ValidatorSetDiffPayload,
    ValidatorChange,
    ValidatorSets
} from "./utils/ValidatorSets.sol";
import {ICMMessage} from "@avalabs/avalanche-contracts/teleporter/ITeleporterMessenger.sol";
import {ICM, AddressedCall} from "./utils/ICM.sol";
import {IVerifyICMMessage} from "./interfaces/IVerifyWarpMessage.sol";

/**
 * @title AvalancheValidatorSetRegistry
 * @notice Registry for managing Avalanche validator sets
 * @dev This contract allows registration and updates of validator sets for Avalanche blockchains.
 * Updates are authenticated through signed ICM messages from the current validator set.
 */
contract AvalancheValidatorSetRegistry is
    IAvalancheValidatorSetRegistry,
    IVerifyICMMessage
{
    uint32 public immutable avalancheNetworkID;
    /**
     * @notice The blockchain this registry maintains validator sets for
     * @dev This should be a blockchain for which the registered validators
     * represents the entire validator set. E.g., if this contract instance is
     * verifying Avalanche L1 instances, this ID should be the L1 ID, not the
     * P-chain ID.
     */
    bytes32 public immutable avalancheBlockChainID;
    uint32 public nextValidatorSetID = 0;

    // Mapping of validator set IDs to their complete validator set data.
    mapping(uint256 => ValidatorSet) private _validatorSets;

    constructor(
        uint32 _avalancheNetworkID,
        bytes32 _avalancheBlockChainID
    ) {
        avalancheNetworkID = _avalancheNetworkID;
        avalancheBlockChainID = _avalancheBlockChainID;
    }

    /**
     * @notice Get the current (latest) validator set
     */
    function getCurrentValidatorSet() public view returns (ValidatorSet memory) {
        require(0 < nextValidatorSetID, "No validator sets exist");
        return _validatorSets[nextValidatorSetID - 1];
    }

    /**
     * @notice Get the ID of the current (latest) validator set
     */
    function getCurrentValidatorSetID() public view returns (uint256) {
        require(0 < nextValidatorSetID, "No validator sets exist");
        return nextValidatorSetID - 1;
    }

    /**
     * @notice Registers a new validator set
     * @dev A validator set can be registered by anyone. The correctness should be verified
     * with the actual validator set on the Avalanche P-Chain.
     * @param message The ICM message containing the validator set to register. The message must be signed by validator set
     * @param validatorBytes The serialized validator set to register.
     * @return The ID of the registered validator set
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external override returns (uint256) {
        // Parse and validate the validator set data
        (
            ValidatorSetStatePayload memory validatorSetStatePayload,
            Validator[] memory validators,
            uint64 totalWeight
        ) = _parseAndValidateValidatorSetData(message, validatorBytes);

        // Construct the validator set and confirm the ICM was self-signed by it.
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: validatorSetStatePayload.avalancheBlockchainID,
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: validatorSetStatePayload.pChainHeight,
            pChainTimestamp: validatorSetStatePayload.pChainTimestamp
        });

        // Check that blockchain ID matches the current validator set.
        require(
            validatorSetStatePayload.avalancheBlockchainID == avalancheBlockChainID,
            "Blockchain ID mismatch"
        );
        ICM.verifyICMMessage(message, avalancheNetworkID,  avalancheBlockChainID,validatorSet);

        // Store the validator set.
        uint256 validatorSetID = nextValidatorSetID++;
        _validatorSets[validatorSetID] = validatorSet;

        emit ValidatorSetRegistered(validatorSetID, validatorSet.avalancheBlockchainID);
        return validatorSetID;
    }

    /**
     * @notice Updates a validator set
     * @dev Updates are authenticated by a signed ICM message from the current validator set
     * @param validatorSetID The ID of the validator set to update
     * @param message The ICM message containing the update
     */
    function updateValidatorSet(
        uint256 validatorSetID,
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external override {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");

        ValidatorSet storage currentValidatorSet = _validatorSets[validatorSetID];

        // Verify the ICM message using the current validator set
        ICM.verifyICMMessage(message, avalancheNetworkID,  avalancheBlockChainID,currentValidatorSet);

        // Parse and validate the validator set data
        (
            ValidatorSetStatePayload memory validatorSetStatePayload,
            Validator[] memory validators,
            uint64 totalWeight
        ) = _parseAndValidateValidatorSetData(message, validatorBytes);

        // Check that blockchain ID matches the current validator set.
        require(
            validatorSetStatePayload.avalancheBlockchainID
                == currentValidatorSet.avalancheBlockchainID,
            "Blockchain ID mismatch"
        );

        // Check that the pChain height is greater than the current validator set.
        require(
            validatorSetStatePayload.pChainHeight > currentValidatorSet.pChainHeight,
            "P-Chain height must be greater than the current validator set"
        );

        // Check that the pChain timestamp is greater than the current validator set.
        require(
            validatorSetStatePayload.pChainTimestamp > currentValidatorSet.pChainTimestamp,
            "P-Chain timestamp must be greater than the current validator set"
        );

        // Update the validator set
        _validatorSets[validatorSetID] = ValidatorSet({
            avalancheBlockchainID: validatorSetStatePayload.avalancheBlockchainID,
            validators: validators,
            totalWeight: totalWeight,
            pChainHeight: validatorSetStatePayload.pChainHeight,
            pChainTimestamp: validatorSetStatePayload.pChainTimestamp
        });

        emit ValidatorSetUpdated(validatorSetID, validatorSetStatePayload.avalancheBlockchainID);
    }

    /**
     * @notice Updates a validator set using a diff (incremental update)
     * @dev This is more gas-efficient than updateValidatorSet when only a few validators change.
     * The diff is verified using the "cryptographic sandwich" technique:
     * 1. Verify the message was signed by the current validator set
     * 2. Verify the previous hash matches our current state
     * 3. Apply the diff to get the new validator set
     * 4. Verify the resulting hash matches the signed current hash
     * @param validatorSetID The ID of the validator set to update
     * @param message The ICM message containing the ValidatorSetDiff payload
     */
    function updateValidatorSetWithDiff(
        uint256 validatorSetID,
        ICMMessage calldata message
    ) external {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");

        ValidatorSet storage currentValidatorSet = _validatorSets[validatorSetID];

        // Verify the ICM message using the current validator set (at previousHeight)
        ICM.verifyICMMessage(message, avalancheNetworkID, avalancheBlockChainID, currentValidatorSet);

        // Parse the addressed call and validate that the source address is empty
        AddressedCall memory addressedCall = ICM.parseAddressedCall(message.unsignedMessage.payload);
        require(addressedCall.sourceAddress.length == 0, "Source address must be empty");

        // Parse the ValidatorSetDiff payload
        ValidatorSetDiffPayload memory diffPayload = ValidatorSets.parseValidatorSetDiffPayload(addressedCall.payload);

        // Verify blockchain ID matches
        require(
            diffPayload.avalancheBlockchainID == currentValidatorSet.avalancheBlockchainID,
            "Blockchain ID mismatch"
        );

        // Step 1: Verify previous hash matches current state (state continuity check)
        bytes memory currentValidatorsBytes = ValidatorSets.serializeValidators(currentValidatorSet.validators);
        bytes32 currentHash = sha256(currentValidatorsBytes);
        require(
            currentHash == diffPayload.previousValidatorSetHash,
            "Previous validator set hash mismatch - state out of sync"
        );

        // Verify height progression
        require(
            diffPayload.previousHeight == currentValidatorSet.pChainHeight,
            "Previous height must match current validator set height"
        );
        require(
            diffPayload.currentHeight > currentValidatorSet.pChainHeight,
            "Current height must be greater than previous height"
        );

        // Verify timestamp progression
        require(
            diffPayload.currentTimestamp > currentValidatorSet.pChainTimestamp,
            "Current timestamp must be greater than previous timestamp"
        );

        // Step 2: Apply the diff to get new validators
        (Validator[] memory newValidators, uint64 newTotalWeight) = ValidatorSets.applyDiff(
            currentValidatorSet.validators,
            diffPayload
        );

        require(newValidators.length > 0, "Resulting validator set cannot be empty");
        require(newTotalWeight > 0, "Total weight must be greater than 0");

        // Step 3: Verify resulting hash matches signed commitment (tamper detection)
        bytes memory newValidatorsBytes = ValidatorSets.serializeValidators(newValidators);
        bytes32 newHash = sha256(newValidatorsBytes);
        require(
            newHash == diffPayload.currentValidatorSetHash,
            "Resulting validator set hash mismatch - diff application error or tampering"
        );

        // Update the validator set
        _validatorSets[validatorSetID] = ValidatorSet({
            avalancheBlockchainID: diffPayload.avalancheBlockchainID,
            validators: newValidators,
            totalWeight: newTotalWeight,
            pChainHeight: diffPayload.currentHeight,
            pChainTimestamp: diffPayload.currentTimestamp
        });

        emit ValidatorSetDiffApplied(
            validatorSetID,
            diffPayload.avalancheBlockchainID,
            diffPayload.added.length,
            diffPayload.removed.length,
            diffPayload.modified.length
        );
    }

    /**
     * @notice Gets a validator set by its ID
     * @param validatorSetID The ID of the validator set to get
     * @return The validator set
     */
    function getValidatorSet(
        uint256 validatorSetID
    ) external view override returns (ValidatorSet memory) {
        return _getValidatorSet(validatorSetID);
    }

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() external view returns (uint32) {
        return avalancheNetworkID;
    }

    /**
     * @notice Verifies an ICM message against a validator set or reverts
     * @dev This function validates that the message is properly signed by a sufficient quorum of validators
     * from the validator set identified by validatorSetID. The verification includes checking the network ID,
     * blockchain ID, and cryptographic signature verification.
     * @param validatorSetID The ID of the validator set to verify the message against
     * @param message The ICM message to verify
     */
    function verifyICMMessageWithID(
        uint256 validatorSetID,
        ICMMessage calldata message
    ) external view {
        ValidatorSet memory validatorSet = _getValidatorSet(validatorSetID);
        ICM.verifyICMMessage(message, avalancheNetworkID, avalancheBlockChainID, validatorSet);
    }

    /**
     * @notice Verify the signatures of an ICM message against the latest validator set or reverts
     * @param message The ICM message to be verified
     */
    function verifyICMMessage(
        ICMMessage calldata message
    ) external view {
        ValidatorSet memory validatorSet = getCurrentValidatorSet();
        ICM.verifyICMMessage(message, avalancheNetworkID, avalancheBlockChainID, validatorSet);
    }

    function _getValidatorSet(
        uint256 validatorSetID
    ) private view returns (ValidatorSet memory) {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");
        return _validatorSets[validatorSetID];
    }

    /**
     * @notice Parses and validates validator set data from an ICM message
     * @dev This is a helper function that consolidates common validation logic
     * @param message The ICM message containing the validator set data
     * @param validatorBytes The serialized validator set
     * @return validatorSetStatePayload The parsed validator set state payload
     * @return validators The parsed validators array
     * @return totalWeight The total weight of all validators
     */
    function _parseAndValidateValidatorSetData(
        ICMMessage calldata message,
        bytes memory validatorBytes
    )
    private
    pure
    returns (
        ValidatorSetStatePayload memory validatorSetStatePayload,
        Validator[] memory validators,
        uint64 totalWeight
    )
    {
        // Parse the addressed call and validate that the source address is empty.
        AddressedCall memory addressedCall = ICM.parseAddressedCall(message.unsignedMessage.payload);
        require(addressedCall.sourceAddress.length == 0, "Source address must be empty");

        // Parse the validator set state payload.
        validatorSetStatePayload =
            ValidatorSets.parseValidatorSetStatePayload(addressedCall.payload);

        // Check that the validator set hash matches the hash of the serialized validator set.
        require(
            validatorSetStatePayload.validatorSetHash == sha256(validatorBytes),
            "Validator set hash mismatch"
        );

        // Parse the validators.
        (validators, totalWeight) = ValidatorSets.parseValidators(validatorBytes);
        require(validators.length > 0, "Validator set cannot be empty");
        require(totalWeight > 0, "Total weight must be greater than 0");
    }
}
