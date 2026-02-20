// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../common/ICM.sol";
import {AvalancheValidatorSetRegistry} from "./AvalancheValidatorSetRegistry.sol";
import {
    ValidatorSetMetadata,
    Validator,
    ValidatorSet,
    ValidatorSetDiff,
    ValidatorSetShard,
    ValidatorSets
} from "./utils/ValidatorSets.sol";

contract DiffUpdater is AvalancheValidatorSetRegistry {
    constructor(
        uint32 avalancheNetworkID_,
        // The metadata about the initial validator set. This is used
        // allow the actual validator set to be populated across multiple
        // transactions
        ValidatorSetMetadata memory initialValidatorSetData
    ) payable AvalancheValidatorSetRegistry(avalancheNetworkID_, initialValidatorSetData) {}

    /**
     * @notice Updates an existing validator set by applying a diff.
     * @dev This function assumes the number of changes in the `diff`
     * (validators added + removed + modified) is small relative to the total validator set size.
     * @param message The ICM message containing the ValidatorSetDiff.
     */
    function updateValidatorSetWithDiff(
        ICMMessage calldata message
    ) external {
        // Safety checks
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        bytes32 chainID = message.sourceBlockchainID;
        require(isRegistered(chainID), "Validator set not registered");
        require(
            !isRegistrationInProgress(chainID),
            "Cannot apply diff with another registration in progress"
        );
        verifyICMMessage(message, chainID);

        // Apply the diff
        ValidatorSet memory currentValidatorSet = this.getValidatorSet(chainID);
        ValidatorSetDiff memory diff = ValidatorSets.parseValidatorSetDiff(
            message.rawMessage, currentValidatorSet.validators.length
        );
        require(diff.currentHeight > currentValidatorSet.pChainHeight, "Invalid blockchain height");

        (Validator[] memory newValidators, uint64 newWeight) =
            ValidatorSets.applyValidatorSetDiff(currentValidatorSet.validators, diff);

        // Update state
        // NOTE: We can simplify this code using --via-ir once a stable compiler version is out.
        ValidatorSet storage storageSet = _validatorSets[chainID];
        storageSet.pChainHeight = diff.currentHeight;
        storageSet.pChainTimestamp = diff.currentTimestamp;
        storageSet.totalWeight = newWeight;
        {
            uint256 oldValLen = storageSet.validators.length;
            uint256 newValLen = newValidators.length;
            uint256 minValLen = newValLen < oldValLen ? newValLen : oldValLen;
            // Overwrite
            uint256 i = 0;
            for (; i < minValLen;) {
                storageSet.validators[i] = newValidators[i];
                unchecked {
                    ++i;
                }
            }
            // Push
            if (newValLen > oldValLen) {
                for (; i < newValLen;) {
                    storageSet.validators.push(newValidators[i]);
                    unchecked {
                        ++i;
                    }
                }
                // Pop
            } else if (oldValLen > newValLen) {
                uint256 diffLen = oldValLen - newValLen;
                for (uint256 j = 0; j < diffLen;) {
                    storageSet.validators.pop();
                    unchecked {
                        ++j;
                    }
                }
            }
        }
        emit ValidatorSetUpdated(chainID);
    }

    /**
     * @dev Applies a set of validators to partial to a set that has been registered.
     */
    function applyShard(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) public override {
        bytes32 avalancheBlockchainID = shard.avalancheBlockchainID;
        // Parse the validators.
        (Validator[] memory validators, uint64 validatorWeight) =
            ValidatorSets.parseValidators(shardBytes);
        require(validators.length > 0, "Validator set cannot be empty");
        require(validatorWeight > 0, "Total weight must exceed 0");

        // update the partial validator set
        for (uint256 i = 0; i < validators.length;) {
            _partialValidatorSets[avalancheBlockchainID].validators.push(validators[i]);
            unchecked {
                ++i;
            }
        }
        _partialValidatorSets[avalancheBlockchainID].partialWeight += validatorWeight;
        _partialValidatorSets[avalancheBlockchainID].shardsReceived += 1;

        // We have received all shards. Place this validator set into the mapping
        if (shard.shardNumber == _partialValidatorSets[avalancheBlockchainID].shardHashes.length) {
            // mark this set as complete
            _partialValidatorSets[avalancheBlockchainID].inProgress = false;
            _validatorSets[avalancheBlockchainID].validators =
                _partialValidatorSets[avalancheBlockchainID].validators;
            _validatorSets[avalancheBlockchainID].totalWeight =
                _partialValidatorSets[avalancheBlockchainID].partialWeight;
        }
    }

    /**
     * @dev The partial validator set is simply a serialized subset of the registered validator set
     */
    function parseValidatorSetMetadata(
        ICMMessage calldata icmMessage,
        bytes calldata shardBytes
    ) public view override returns (ValidatorSetMetadata memory, Validator[] memory, uint64) {
        // Parse the validator set state payload.
        ValidatorSetMetadata memory validatorSetMetadata =
            ValidatorSets.parseValidatorSetMetadata(icmMessage.rawMessage);
        // Check that the first validator set shard hash matches the hash of the serialized validator set.
        require(
            validatorSetMetadata.shardHashes[0] == sha256(shardBytes), "Validator set hash mismatch"
        );
        // Parse the validators.
        (Validator[] memory validators, uint64 totalWeight) =
            ValidatorSets.parseValidators(shardBytes);
        bytes32 avalancheBlockchainID = validatorSetMetadata.avalancheBlockchainID;
        require(
            _validatorSets[avalancheBlockchainID].pChainHeight < validatorSetMetadata.pChainHeight,
            "P-Chain height too low"
        );
        require(
            _validatorSets[avalancheBlockchainID].pChainTimestamp
                < validatorSetMetadata.pChainTimestamp,
            "P-Chain timestamp too low"
        );
        require(validators.length > 0, "Validator set cannot be empty");
        require(totalWeight > 0, "Total weight must exceed 0");
        return (validatorSetMetadata, validators, totalWeight);
    }
}
