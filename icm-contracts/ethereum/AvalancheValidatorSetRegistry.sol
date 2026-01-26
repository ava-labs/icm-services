// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IAvalancheValidatorSetRegistry} from "./interfaces/IAvalancheValidatorSetRegistry.sol";
import {ICMMessage} from "../common/ICM.sol";
import {
    PartialValidatorSet,
    ValidatorSetMetadata,
    Validator,
    ValidatorSet,
    ValidatorSetShard,
    ValidatorSetSignature,
    ValidatorSets
} from "./utils/ValidatorSets.sol";

contract AvalancheValidatorSetRegistry is IAvalancheValidatorSetRegistry {
    uint32 public immutable avalancheNetworkID;
    // The Avalanche blockchain ID of the P-chain
    bytes32 public immutable pChainID;
    // Mapping of Avalanche blockchain IDs to their complete validator set data.
    mapping(bytes32 => ValidatorSet) private _validatorSets;

    // Mapping of Avalanche blockchain IDs to a subset of said validator set.
    // Used to allow updating validator sets across multiple transactions
    mapping(bytes32 => PartialValidatorSet) private _partialValidatorSets;

    constructor(
        uint32 avalancheNetworkID_,
        // The metadata about the initial validator set. This is used
        // allow the actual validator set to be populated across multiple
        // transactions
        ValidatorSetMetadata memory initialValidatorSetData
    ) {
        avalancheNetworkID = avalancheNetworkID_;
        pChainID = initialValidatorSetData.avalancheBlockchainID;

        _partialValidatorSets[pChainID] = PartialValidatorSet({
            pChainHeight: initialValidatorSetData.pChainHeight,
            pChainTimestamp: initialValidatorSetData.pChainTimestamp,
            validatorSetHash: initialValidatorSetData.validatorSetHash,
            shardHashes: initialValidatorSetData.shardHashes,
            shardsReceived: 0,
            validators: new Validator[](initialValidatorSetData.totalValidators),
            numValidators: 0,
            partialWeight: 0
        });
        // pre-allocate the storage slots
        _validatorSets[pChainID] = ValidatorSet({
            avalancheBlockchainID: initialValidatorSetData.avalancheBlockchainID,
            pChainHeight: initialValidatorSetData.pChainHeight,
            pChainTimestamp: initialValidatorSetData.pChainTimestamp,
            validators: new Validator[](initialValidatorSetData.totalValidators),
            totalWeight: 0
        });
    }

    /**
     * @notice Registers a new validator set
     * @dev It may be the case that the entire validator set cannot be passed into this function. If a partial
     * validator set is provided, the chain is still considered registered. To pass the rest of the validator set
     * data, `updateValidatorSet` should be called instead.
     * @param message The ICM message containing the validator set to register. The message must be signed by
     * `pChainValidatorSet`
     * @param validatorBytes The serialized validator set to register.
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes calldata validatorBytes
    ) external override {
        // Check the network ID and signature
        // N.B. this message should be signed by the currently registered P-chain validator set
        verifyICMMessage(message, pChainID);

        // Parse and validate the validator set payload
        (
            ValidatorSetMetadata memory validatorSetMetadata,
            Validator[] memory validators,
            uint64 validatorWeight
        ) = _parseAndValidateValidatorSetPayload(message, validatorBytes);

        // a blockchain cannot be registered more than once.
        if (_isRegistered(validatorSetMetadata.avalancheBlockchainID)) {
            revert("A blockchain with this ID has already been registered");
        }

        // This validator set is sharded
        if (validatorSetMetadata.shardHashes.length > 1) {
            // pre-allocate enough storage for the whole validator set
            Validator[] memory valSet = new Validator[](validatorSetMetadata.totalValidators);
            for (uint256 i = 0; i < validators.length; i++) {
                valSet[i] = validators[i];
            }

            // initialize the partial validator set and store it
            _partialValidatorSets[validatorSetMetadata.avalancheBlockchainID] = PartialValidatorSet({
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                validatorSetHash: validatorSetMetadata.validatorSetHash,
                shardHashes: validatorSetMetadata.shardHashes,
                shardsReceived: 1,
                validators: valSet,
                numValidators: validators.length,
                partialWeight: validatorWeight
            });
            // pre-allocate storage for the completed validator set
            _validatorSets[validatorSetMetadata.avalancheBlockchainID] = ValidatorSet({
                avalancheBlockchainID: validatorSetMetadata.avalancheBlockchainID,
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                validators: new Validator[](validatorSetMetadata.totalValidators),
                totalWeight: 0
            });

            // TODO: emit event here?
        } else {
            // Construct the validator set
            ValidatorSet memory validatorSet = ValidatorSet({
                avalancheBlockchainID: validatorSetMetadata.avalancheBlockchainID,
                validators: validators,
                totalWeight: validatorWeight,
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp
            });
            // Store the validator set.
            _validatorSets[validatorSet.avalancheBlockchainID] = validatorSet;
            emit ValidatorSetRegistered(validatorSet.avalancheBlockchainID);
        }
    }

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() external view returns (uint32) {
        return avalancheNetworkID;
    }

    /**
     * @notice Apply a shard as part of initializing the very first P-chain validator set
     */
    function updateInitialPChainValidatorSet(
        ValidatorSetShard calldata pChainShard,
        bytes memory validatorBytes
    ) public {
        require(!pChainInitialized(), "The initial P-chain validator set is already initialized");
        _applyShard(pChainShard, validatorBytes);
    }

    /**
     * @dev Check if a P-chain validator set has been completely registered.
     * Until is has, no functions other than adding to this validator set are
     * permitted.
     */
    function pChainInitialized() public view returns (bool) {
        return _validatorSets[pChainID].totalWeight != 0;
    }

    /**
     * @notice Verify that the message is
     *   1. Check that the contracts root of trust has been initialized completely
     *   2. Intended for this network ID
     *   3. Has been signed by a quorum of the provided validator set
     * If any check fails, reverts
     */
    function verifyICMMessage(
        ICMMessage calldata message,
        bytes32 avalancheBlockchainID
    ) public view {
        require(
            pChainInitialized(),
            "A complete P-chain validator must be registered to verify ICM messages"
        );
        require(
            _isRegistered(avalancheBlockchainID),
            "No validator set is registered for the provided Avalanche blockchain ID"
        );
        require(message.message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        ValidatorSetSignature memory sig =
            ValidatorSets.parseValidatorSetSignature(message.attestation);
        require(
            ValidatorSets.verifyValidatorSetSignature(
                sig, message.rawMessageBytes, _validatorSets[avalancheBlockchainID]
            ),
            "Could not register validator set: Signature checks failed"
        );
    }

    /**
     * @dev Apply a shard to a partial validator set. If the set is completed by this shard, copy
     * it over to the `_validatorSets` mapping.
     */
    function _applyShard(ValidatorSetShard calldata shard, bytes memory validatorBytes) private {
        bytes32 avalancheBlockchainID = shard.avalancheBlockchainID;
        require(
            _partialValidatorSets[avalancheBlockchainID].shardsReceived + 1 == shard.shardNumber,
            "Received shard out of order"
        );
        require(
            _partialValidatorSets[avalancheBlockchainID].shardHashes[shard.shardNumber - 1]
                == sha256(validatorBytes),
            "Unexpected shard hash"
        );
        // Parse the validators.
        (Validator[] memory validators, uint64 validatorWeight) =
            ValidatorSets.parseValidators(validatorBytes);
        require(validators.length > 0, "Validator set cannot be empty");
        require(validatorWeight > 0, "Total weight must be greater than 0");

        // update the partial validator set
        uint256 offset = _partialValidatorSets[avalancheBlockchainID].numValidators;
        for (uint256 i = 0; i < validators.length; i++) {
            _partialValidatorSets[avalancheBlockchainID].validators[offset + i] = validators[i];
        }
        _partialValidatorSets[avalancheBlockchainID].numValidators += validators.length;
        _partialValidatorSets[avalancheBlockchainID].partialWeight += validatorWeight;
        _partialValidatorSets[avalancheBlockchainID].shardsReceived += 1;

        // We have received all shards. Place this validator set into the mapping
        if (shard.shardNumber == _partialValidatorSets[avalancheBlockchainID].shardHashes.length) {
            _validatorSets[avalancheBlockchainID].validators =
                _partialValidatorSets[avalancheBlockchainID].validators;
            _validatorSets[avalancheBlockchainID].totalWeight =
                _partialValidatorSets[avalancheBlockchainID].partialWeight;
        }
    }

    /**
     * @dev Check if a validator set is already registered under the given `avalancheBlockchainID`.
     */
    function _isRegistered(
        bytes32 avalancheBlockchainID
    ) private view returns (bool) {
        return _validatorSets[avalancheBlockchainID].totalWeight != 0;
    }

    /**
     * @notice Parses and validates (potentially partial) validator set data from an ICM message. This
     * is called when registering validator sets.
     * @dev This is a helper function that consolidates common validation logic
     * @param icmMessage The ICM message containing the validator set data
     * @param validatorBytes The serialized validator set
     * @return The parsed validator set state payload
     * @return The parsed validators array
     * @return The total weight of all validators
     */
    function _parseAndValidateValidatorSetPayload(
        ICMMessage calldata icmMessage,
        bytes calldata validatorBytes
    ) private pure returns (ValidatorSetMetadata memory, Validator[] memory, uint64) {
        // Validate that the source address is empty.
        require(icmMessage.message.sourceAddress == address(0), "Source address must be empty");

        // Parse the validator set state payload.
        ValidatorSetMetadata memory validatorSetStatePayload =
            ValidatorSets.parseValidatorSetMetadata(icmMessage.message.payload);
        // Check that the first validator set shard hash matches the hash of the serialized validator set.
        require(
            validatorSetStatePayload.shardHashes[0] == sha256(validatorBytes),
            "Validator set hash mismatch"
        );
        // Parse the validators.
        (Validator[] memory validators, uint64 totalWeight) =
            ValidatorSets.parseValidators(validatorBytes);

        require(validators.length > 0, "Validator set cannot be empty");
        require(totalWeight > 0, "Total weight must be greater than 0");
        return (validatorSetStatePayload, validators, totalWeight);
    }
}
