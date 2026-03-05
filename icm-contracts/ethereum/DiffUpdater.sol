// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../common/ICM.sol";
import {AvalancheValidatorSetRegistry} from "./AvalancheValidatorSetRegistry.sol";
import {
    PartialValidatorSet,
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
     */
    function applyShard(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) public override {
        bytes32 chainID = shard.avalancheBlockchainID;
        PartialValidatorSet storage currentPartialValSet = _partialValidatorSets[chainID];
        ValidatorSetDiff memory diff =
            ValidatorSets.parseValidatorSetDiff(shardBytes, currentPartialValSet.validators.length);

        // Safety Checks
        ValidatorSet memory currentValidatorSet = this.getValidatorSet(chainID);
        require(diff.avalancheBlockchainID == chainID, "Blockchain ID mismatch");
        require(
            diff.previousHeight == currentValidatorSet.pChainHeight, "Diff anchor height mismatch"
        );
        require(
            diff.previousTimestamp == currentValidatorSet.pChainTimestamp,
            "Diff anchor timestamp mismatch"
        );
        require(diff.currentHeight > currentValidatorSet.pChainHeight, "Invalid blockchain height");
        require(
            diff.currentTimestamp > currentValidatorSet.pChainTimestamp,
            "Invalid blockchain timestamp"
        );

        // Apply Diff
        (Validator[] memory newValidators, uint64 newWeight) =
            ValidatorSets.applyValidatorSetDiff(currentPartialValSet.validators, diff);
        require(
            newValidators.length > 0 && diff.newSize == newValidators.length,
            "Resulting validator set size mismatch"
        );
        require(newWeight > 0, "Total weight must exceed 0");
        require(
            diff.currentValidatorSetHash == sha256(ValidatorSets.serializeValidators(newValidators)),
            "Current validator hash mismatch"
        );

        // Update Intermediate State
        currentPartialValSet.pChainHeight = diff.currentHeight;
        currentPartialValSet.pChainTimestamp = diff.currentTimestamp;
        applyPartialUpdate(chainID, newValidators, newWeight);
        currentPartialValSet.shardsReceived += 1;

        // Update Final State
        if (shard.shardNumber == _partialValidatorSets[chainID].shardHashes.length) {
            currentPartialValSet.inProgress = false;
            ValidatorSets.replaceValidators(_validatorSets[chainID].validators, newValidators);
            _validatorSets[chainID].totalWeight = currentPartialValSet.partialWeight;
            _validatorSets[chainID].pChainHeight = currentPartialValSet.pChainHeight;
            _validatorSets[chainID].pChainTimestamp = currentPartialValSet.pChainTimestamp;
            emit ValidatorSetUpdated(chainID);
        }
    }

    /**
     * @dev Note that the operations of changes from a diff are commutative. So we can successively
     * apply a subset of them and always keep the latest result as the partial update.
     */
    function applyPartialUpdate(
        bytes32 avalancheBlockchainID,
        Validator[] memory partialUpdate,
        uint64 partialWeightUpdate
    ) public override {
        PartialValidatorSet storage pSet = _partialValidatorSets[avalancheBlockchainID];
        pSet.partialWeight = partialWeightUpdate;
        ValidatorSets.replaceValidators(pSet.validators, partialUpdate);
    }

    /**
     * @dev The returned validator set if the result of applying the first diff to the currently registered
     * set. The weight returned is the weight of the resulting set.
     */
    function parseValidatorSetMetadata(
        ICMMessage calldata icmMessage,
        bytes calldata shardBytes
    ) public view override returns (ValidatorSetMetadata memory, Validator[] memory, uint64) {
        // Parse
        ValidatorSetMetadata memory validatorSetMetadata =
            ValidatorSets.parseValidatorSetMetadata(icmMessage.rawMessage);
        require(
            validatorSetMetadata.shardHashes[0] == sha256(shardBytes), "Validator set hash mismatch"
        );
        bytes32 chainID = icmMessage.sourceBlockchainID;
        ValidatorSet memory currentValidatorSet = this.getValidatorSet(chainID);
        ValidatorSetDiff memory diff =
            ValidatorSets.parseValidatorSetDiff(shardBytes, currentValidatorSet.validators.length);

        // Safety Checks
        require(diff.avalancheBlockchainID == chainID, "Blockchain ID mismatch");
        require(
            diff.previousHeight == currentValidatorSet.pChainHeight, "Diff anchor height mismatch"
        );
        require(
            diff.previousTimestamp == currentValidatorSet.pChainTimestamp,
            "Diff anchor timestamp mismatch"
        );
        require(
            diff.currentHeight > currentValidatorSet.pChainHeight
                && diff.currentHeight == validatorSetMetadata.pChainHeight,
            "Invalid diff height"
        );
        require(
            diff.currentTimestamp > currentValidatorSet.pChainTimestamp
                && diff.currentTimestamp == validatorSetMetadata.pChainTimestamp,
            "Invalid diff timestamp"
        );

        // Apply
        (Validator[] memory newValidators, uint64 newWeight) =
            ValidatorSets.applyValidatorSetDiff(currentValidatorSet.validators, diff);
        require(
            newValidators.length > 0 && diff.newSize == newValidators.length,
            "Resulting validator set size mismatch"
        );
        require(newWeight > 0, "Total weight must exceed 0");
        require(
            diff.currentValidatorSetHash == sha256(ValidatorSets.serializeValidators(newValidators)),
            "Current validator hash mismatch"
        );
        return (validatorSetMetadata, newValidators, newWeight);
    }
}
