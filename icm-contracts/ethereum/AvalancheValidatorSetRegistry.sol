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

// A contract for managing Avalanche validator sets which can be used to verify ICM messages
// via a quorum of signatures.
//
// In addition to verifying ICM messages, it contains logic for updating validator sets
// which may need to occur across multiple transactions. This contract is agnostic on
// how the data is sharded across these transactions. Two virtual functions should
// be overridden in a child contract to specify this.
contract AvalancheValidatorSetRegistry is IAvalancheValidatorSetRegistry {
    uint32 public immutable avalancheNetworkID;
    // The Avalanche blockchain ID of the P-chain
    bytes32 public immutable pChainID;
    // Mapping of Avalanche blockchain IDs to their complete validator set data.
    mapping(bytes32 => ValidatorSet) internal _validatorSets;

    // Mapping of Avalanche blockchain IDs to a subset of said validator set.
    // Used to allow updating validator sets across multiple transactions
    mapping(bytes32 => PartialValidatorSet) internal _partialValidatorSets;

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
            shardHashes: initialValidatorSetData.shardHashes,
            shardsReceived: 0,
            validators: new Validator[](0),
            partialWeight: 0,
            inProgress: true
        });

        _validatorSets[pChainID] = ValidatorSet({
            avalancheBlockchainID: initialValidatorSetData.avalancheBlockchainID,
            pChainHeight: initialValidatorSetData.pChainHeight,
            pChainTimestamp: initialValidatorSetData.pChainTimestamp,
            validators: new Validator[](0),
            totalWeight: 0
        });
    }

    /**
     * @notice Registers a new validator set for a specific Avalanche blockchain ID.
     * A registered set may require several txs to populate due to size/gas constraints.
     * The population process should not be interrupted with a new register call. If this
     * function is called to register a new validator set for chain for a which a partial
     * set exists, this function will revert.
     *
     * Emits an event that a new set has been registered.
     * @dev It may be the case that the entire validator set cannot be passed into this function.
     * If a partial validator set is provided, the chain is still considered registered. To pass
     * the rest of the validator set data, `updateValidatorSet` should be called instead.
     * @param message The ICM message containing the validator set to register. The message must
     * be signed by the relevant authorizing party
     * @param shardBytes The first shard of the data needed to populate the newly registered
     * validator set.
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes calldata shardBytes
    ) external override {
        // Check the network ID and signature
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        // Check that we are not interrupting an existing registration
        if (isRegistrationInProgress(message.sourceBlockchainID)) {
            // check if we are interrupting an existing registration
            revert("Can't register to a blockchain ID while another registration is in progress");
        }

        // Check if this is the first time this blockchain is registering a validator set
        if (!isRegistered(message.sourceBlockchainID)) {
            // N.B. this message should be signed by the currently registered P-chain validator set
            verifyICMMessage(message, pChainID);
        } else {
            // This blockchain ID has an existing validator set registered to it which should
            // have signed this message
            verifyICMMessage(message, message.sourceBlockchainID);
        }

        // Parse and validate the validator set payload
        (
            ValidatorSetMetadata memory validatorSetMetadata,
            Validator[] memory validators,
            uint64 validatorWeight
        ) = parseValidatorSetMetadata(message, shardBytes);
        bytes32 avalancheBlockchainID = validatorSetMetadata.avalancheBlockchainID;
        require(message.sourceBlockchainID == avalancheBlockchainID, "Source chain ID mismatch");

        // This validator set is sharded
        if (validatorSetMetadata.shardHashes.length > 1) {
            // initialize the partial validator set and store it
            _partialValidatorSets[validatorSetMetadata.avalancheBlockchainID] = PartialValidatorSet({
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                shardHashes: validatorSetMetadata.shardHashes,
                shardsReceived: 1,
                validators: validators,
                partialWeight: validatorWeight,
                inProgress: true
            });
            if (!isRegistered(message.sourceBlockchainID)) {
                // pre-allocate storage for the completed validator set
                _validatorSets[validatorSetMetadata.avalancheBlockchainID] = ValidatorSet({
                    avalancheBlockchainID: validatorSetMetadata.avalancheBlockchainID,
                    pChainHeight: validatorSetMetadata.pChainHeight,
                    pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                    validators: new Validator[](0),
                    totalWeight: 0
                });
            }
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
        }
        emit ValidatorSetRegistered(avalancheBlockchainID);
    }

    /**
     * @notice Apply a shard to a registered validator set
     *
     * Emits an event if a partial set becomes fully populated.
     * @param shard indicates the shard number and blockchain ID of the partial
     * set that is being updated.
     * @param shardBytes The next shard of data needed to populate a registered partial
     * validator set
     *
     */
    function updateValidatorSet(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) external {
        require(
            isRegistrationInProgress(shard.avalancheBlockchainID),
            "Cannot apply shard if registration is not in progress"
        );
        bytes32 avalancheBlockchainID = shard.avalancheBlockchainID;
        require(
            _partialValidatorSets[avalancheBlockchainID].shardsReceived + 1 == shard.shardNumber,
            "Received shard out of order"
        );
        require(
            _partialValidatorSets[avalancheBlockchainID].shardHashes[shard.shardNumber - 1]
                == sha256(shardBytes),
            "Unexpected shard hash"
        );
        applyShard(shard, shardBytes);
        if (!isRegistrationInProgress(shard.avalancheBlockchainID)) {
            emit ValidatorSetUpdated(shard.avalancheBlockchainID);
        }
    }

    /**
     * @notice  Validate and apply a shard to a partial validator set. If the set is completed by this shard, copy
     * it over to the `_validatorSets` mapping.
     * @param shard Indicates the sequence number of the shard and blockchain affected by this update
     * @param shardBytes the actual data of the shard which
     */
    /* solhint-disable-next-line no-unused-vars */
    function applyShard(ValidatorSetShard calldata shard, bytes memory shardBytes) public virtual {
        revert("Not implemented");
    }

    /**
     * @notice Parses and validates metadata about a validator set data from an ICM message. This
     * is called when registering validator sets. It may also contain a (potentially partial) subset of
     * the validators that are being registered. This is always considered to be the first shard of
     * the requisite data.
     *
     * @param icmMessage The ICM message containing the validator set metadata
     * @param shardBytes The serialized data used to construct a subset of the registered
     * validator set
     * @return The parsed validator set metadata
     * @return A parsed validators array
     * @return The total weight of the parsed validators
     */
    function parseValidatorSetMetadata(
        /* solhint-disable-next-line no-unused-vars */
        ICMMessage calldata icmMessage,
        /* solhint-disable-next-line no-unused-vars */
        bytes calldata shardBytes
    ) public view virtual returns (ValidatorSetMetadata memory, Validator[] memory, uint64) {
        revert("Not implemented");
    }

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() public view returns (uint32) {
        return avalancheNetworkID;
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
            isRegistered(avalancheBlockchainID),
            "No validator set is registered for the provided Avalanche blockchain ID"
        );
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        ValidatorSetSignature memory sig =
            ValidatorSets.parseValidatorSetSignature(message.attestation);
        require(
            ValidatorSets.verifyValidatorSetSignature(
                sig, message.rawMessage, _validatorSets[avalancheBlockchainID]
            ),
            "Could not verify ICM message: Signature checks failed"
        );
    }

    /**
     * @notice Check if a **complete** validator set is registered (not just a partial).
     */
    function isRegistered(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return _validatorSets[avalancheBlockchainID].totalWeight != 0;
    }

    /**
     * @notice Check if a validator set is registered but awaiting further updates.
     */
    function isRegistrationInProgress(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return _partialValidatorSets[avalancheBlockchainID].inProgress;
    }
}

// This contract specifies that shards of validator sets is a serialized subsequence of
// the entire validator set that has been cryptographically committed to.
contract FullSetUpdater is AvalancheValidatorSetRegistry {
    constructor(
        uint32 avalancheNetworkID_,
        // The metadata about the initial validator set. This is used
        // allow the actual validator set to be populated across multiple
        // transactions
        ValidatorSetMetadata memory initialValidatorSetData
    ) AvalancheValidatorSetRegistry(avalancheNetworkID_, initialValidatorSetData) {}

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
        require(validatorWeight > 0, "Total weight must be greater than 0");

        // update the partial validator set
        for (uint256 i = 0; i < validators.length; i++) {
            _partialValidatorSets[avalancheBlockchainID].validators.push(validators[i]);
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
            "P-Chain height must be greater than the current validator set"
        );
        require(
            _validatorSets[avalancheBlockchainID].pChainTimestamp
                < validatorSetMetadata.pChainTimestamp,
            "P-Chain timestamp must be greater than the current validator set"
        );
        require(validators.length > 0, "Validator set cannot be empty");
        require(totalWeight > 0, "Total weight must be greater than 0");
        return (validatorSetMetadata, validators, totalWeight);
    }
}
