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
     * reset registration. The anchor checks against the stale registered
     * set are skipped because by construction a reset diff carries
     * `previousHeight == 0 && previousTimestamp == 0` (so they would
     * trivially fail), and the diff height/timestamp comparisons against
     * the registered set are redundant: the bound between the signed
     * metadata's `pChainHeight`/`pChainTimestamp` and the registered set
     * was already enforced when the first shard was processed in
     * `parseValidatorSetMetadata`. Each subsequent shard's bytes are
     * authenticated via `shardHashes` in that signed metadata, so the
     * reset path relies on that commitment for integrity.
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
     * The diff carries `previousHeight == 0 && previousTimestamp == 0` in
     * two cases: (a) a first-time registration of a chain (no prior
     * state) and (b) a "reset" of an already-registered chain. In both
     * cases the diff describes the entire current validator set as
     * additions against an empty starting set, which lets the relayer
     * recover from the chain's registered P-chain height falling below
     * the P-chain's minimum retained height (so L1-signed incremental
     * updates can no longer be produced).
     *
     * Only the second case is reported as `isReset == true` to the
     * parent contract; the parent then emits `ValidatorSetReset` instead
     * of `ValidatorSetRegistered`.
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

        bool startsFromEmpty = (diff.previousHeight == 0 && diff.previousTimestamp == 0);
        // A reset is a wholesale replacement of an already registered set.
        // When the chain is not yet registered the prev=0/0 markers simply
        // describe a fresh first-time registration, not a reset.
        bool isReset = startsFromEmpty && isRegistered(chainID);

        if (startsFromEmpty) {
            // The diff is applied against an empty starting set, so every
            // change entry must add a new validator with positive weight.
            // A weight-0 entry would be a removal against an empty set
            // (impossible) and a weight-only modification is also
            // nonsensical here.
            require(
                diff.numAdded == diff.changes.length,
                "Diff against empty set must contain only additions"
            );
            // The diff was parsed with
            // `parseValidatorSetDiff(shardBytes, currentValidatorSet.validators.length)`,
            // which computes `newSize = currentValidatorCount + numAdded -
            // numRemoved`. For a reset of an already-registered chain the
            // currentValidatorCount is non-zero, so override to interpret
            // newSize against an empty starting set so that the subsequent
            // size check matches the actual resulting length.
            diff.newSize = uint256(diff.numAdded);
            // The prev anchor checks against the registered set are
            // skipped here: prev=0/0 is the marker for "start from
            // empty", so by construction it cannot match the registered
            // set's anchor on a reset, and on a first registration the
            // registered set's anchor is 0/0 so the check would be a
            // tautology.
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

        // Apply: when starting from empty, the diff is applied against an
        // empty set so any previously registered validators are discarded
        // wholesale. For a regular incremental update the diff is applied
        // against the registered set.
        Validator[] memory startingSet =
            startsFromEmpty ? new Validator[](0) : currentValidatorSet.validators;
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
