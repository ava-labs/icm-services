// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
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

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
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
     *
     * When `currentPartialValSet.isReset` is true this shard belongs to a
     * reset registration whose first shard was already validated in
     * `parseValidatorSetMetadata`. The anchor checks against the stale
     * registered set do not apply to a reset, so they are skipped. The
     * shard's bytes are still authenticated via `shardHashes` in the signed
     * metadata parsed by shard 0, so the reset path relies on that
     * commitment for integrity.
     *
     * Does not emit `ValidatorSetUpdated` directly; the base
     * `updateValidatorSet` checks the reset flag and emits either
     * `ValidatorSetUpdated` or `ValidatorSetReset`.
     */
    function applyShard(
        ValidatorSetShard calldata shard,
        bytes calldata shardBytes
    ) public override {
        bytes32 chainID = shard.avalancheBlockchainID;
        PartialValidatorSet storage currentPartialValSet = _partialValidatorSets[chainID];
        ValidatorSetDiff memory diff =
            ValidatorSets.parseValidatorSetDiff(shardBytes, currentPartialValSet.validators.length);

        require(diff.avalancheBlockchainID == chainID, "Blockchain ID mismatch");

        if (!currentPartialValSet.isReset) {
            ValidatorSet memory currentValidatorSet = this.getValidatorSet(chainID);
            require(
                diff.previousHeight == currentValidatorSet.pChainHeight,
                "Diff anchor height mismatch"
            );
            require(
                diff.previousTimestamp == currentValidatorSet.pChainTimestamp,
                "Diff anchor timestamp mismatch"
            );
            require(diff.currentHeight > currentValidatorSet.pChainHeight, "P-Chain height too low");
            require(
                diff.currentTimestamp > currentValidatorSet.pChainTimestamp,
                "P-Chain timestamp too low"
            );
        }

        // Apply Diff
        Validator[] memory newValidators = applyDiff(chainID, diff);
        // Update Final State
        if (shard.shardNumber == _partialValidatorSets[chainID].shardHashes.length) {
            currentPartialValSet.inProgress = false;
            ValidatorSets.replaceValidators(_validatorSets[chainID].validators, newValidators);
            _validatorSets[chainID].totalWeight = currentPartialValSet.partialWeight;
            _validatorSets[chainID].pChainHeight = currentPartialValSet.pChainHeight;
            _validatorSets[chainID].pChainTimestamp = currentPartialValSet.pChainTimestamp;
        }
    }

    function applyDiff(
        bytes32 chainID,
        ValidatorSetDiff memory diff
    ) public returns (Validator[] memory) {
        PartialValidatorSet storage currentPartialValSet = _partialValidatorSets[chainID];
        (Validator[] memory newValidators, uint64 newWeight) =
            ValidatorSets.applyValidatorSetDiff(currentPartialValSet.validators, diff);
        require(
            newValidators.length > 0 && diff.newSize == newValidators.length,
            "Resulting validator set size mismatch"
        );
        require(newWeight > 0, "Total weight must exceed 0");

        // Update Intermediate State
        currentPartialValSet.pChainHeight = diff.currentHeight;
        currentPartialValSet.pChainTimestamp = diff.currentTimestamp;
        applyPartialUpdate(chainID, newValidators, newWeight);
        currentPartialValSet.shardsReceived += 1;
        return newValidators;
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
     * @dev Returns the validator set that results from applying the first diff.
     *
     * The diff is interpreted as a reset when it carries `previousHeight == 0`
     * and `previousTimestamp == 0`. A reset diff describes the entire current
     * validator set as additions against an empty starting set, which lets
     * the relayer recover when the registered on-chain P-chain height has
     * fallen below the P-chain's minimum retained height and L1-signed
     * incremental updates can no longer be produced. First-time registration
     * also uses this wire format; the caller distinguishes "first
     * registration" from "reset of an already-registered chain" based on
     * whether the chain has a non-zero total weight on-chain.
     */
    function parseValidatorSetMetadata(
        ICMMessage calldata icmMessage,
        bytes calldata shardBytes
    )
        public
        view
        override
        returns (ValidatorSetMetadata memory, Validator[] memory, uint64, bool)
    {
        // Parse
        ValidatorSetMetadata memory validatorSetMetadata =
            ValidatorSets.parseValidatorSetMetadata(icmMessage.rawMessage);
        require(
            validatorSetMetadata.shardHashes[0] == sha256(shardBytes), "Validator set hash mismatch"
        );
        bytes32 chainID = validatorSetMetadata.avalancheBlockchainID;
        ValidatorSet memory currentValidatorSet = this.getValidatorSet(chainID);
        ValidatorSetDiff memory diff =
            ValidatorSets.parseValidatorSetDiff(shardBytes, currentValidatorSet.validators.length);

        // Safety Checks
        require(diff.avalancheBlockchainID == chainID, "Blockchain ID mismatch");

        bool isReset = (diff.previousHeight == 0 && diff.previousTimestamp == 0);

        if (isReset) {
            // A reset must describe only additions: the diff is applied
            // against an empty starting set, so every change entry must add
            // a new validator with positive weight. A weight-0 entry would
            // be a removal against an empty set (impossible) and a
            // weight-only modification is also nonsensical here.
            require(diff.numAdded == diff.changes.length, "Reset diff must contain only additions");
            // The diff was parsed with
            // `parseValidatorSetDiff(shardBytes, currentValidatorSet.validators.length)`,
            // which computes `newSize = currentValidatorCount + numAdded -
            // numRemoved`. For a reset we must instead interpret newSize
            // against an empty starting set so that the subsequent size
            // check matches the actual resulting length.
            diff.newSize = uint256(diff.numAdded);
        } else {
            require(
                diff.previousHeight == currentValidatorSet.pChainHeight, "P-Chain height too low"
            );
            require(
                diff.previousTimestamp == currentValidatorSet.pChainTimestamp,
                "P-Chain timestamp too low"
            );
        }
        require(
            diff.currentHeight > currentValidatorSet.pChainHeight
                && diff.currentHeight == validatorSetMetadata.pChainHeight,
            "P-Chain height too low"
        );
        require(
            diff.currentTimestamp > currentValidatorSet.pChainTimestamp
                && diff.currentTimestamp == validatorSetMetadata.pChainTimestamp,
            "P-Chain timestamp too low"
        );

        // Apply: for a reset the diff is applied against an empty set so the
        // previously registered validators are discarded wholesale. For a
        // regular incremental update the diff is applied against the
        // registered set.
        Validator[] memory startingSet =
            isReset ? new Validator[](0) : currentValidatorSet.validators;
        (Validator[] memory newValidators, uint64 newWeight) =
            ValidatorSets.applyValidatorSetDiff(startingSet, diff);
        require(
            newValidators.length > 0 && diff.newSize == newValidators.length,
            "Resulting validator set size mismatch"
        );
        require(newWeight > 0, "Total weight must exceed 0");
        return (validatorSetMetadata, newValidators, newWeight, isReset);
    }
}
