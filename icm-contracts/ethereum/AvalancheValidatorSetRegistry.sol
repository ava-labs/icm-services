// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IAvalancheValidatorSetRegistry} from "./interfaces/IAvalancheValidatorSetRegistry.sol";
import {IAvalancheValidatorSetManager} from "./interfaces/IAvalancheValidatorSetManager.sol";
import {ICMMessage} from "../common/ICM.sol";
import {
    PartialValidatorSet,
    ValidatorSetMetadata,
    Validator,
    ValidatorSet,
    ValidatorSetDiffPayload,
    ValidatorSetShard,
    ValidatorSetSignature,
    ValidatorSets
} from "./utils/ValidatorSets.sol";

// A contract for managing Avalanche validator sets which can be used to verify ICM messages
// via a quorum of signatures.
//
// In addition to verifying ICM messages, it contains logic for updating validator sets
// which may need to occur across multiple transactions. This contract outsources the data
// management to another contract which acts as the data layer.
contract AvalancheValidatorSetRegistry is IAvalancheValidatorSetRegistry {
    uint32 public immutable avalancheNetworkID;
    // The Avalanche blockchain ID of the P-chain
    bytes32 public immutable pChainID;

    // The contract responsible for updating validator sets from shards
    address public immutable validatorSetManagerContract;

    constructor(
        uint32 avalancheNetworkID_,
        // The metadata about the initial validator set. This is used
        // allow the actual validator set to be populated across multiple
        // transactions
        ValidatorSetMetadata memory initialValidatorSetData,
        address validatorSetUpdaterContract_
    ) {
        avalancheNetworkID = avalancheNetworkID_;
        pChainID = initialValidatorSetData.avalancheBlockchainID;
        validatorSetManagerContract = validatorSetUpdaterContract_;

        PartialValidatorSet memory partialValidatorSet = PartialValidatorSet({
            pChainHeight: initialValidatorSetData.pChainHeight,
            pChainTimestamp: initialValidatorSetData.pChainTimestamp,
            shardHashes: initialValidatorSetData.shardHashes,
            shardsReceived: 0,
            validators: new Validator[](initialValidatorSetData.totalValidators),
            numValidators: 0,
            partialWeight: 0,
            inProgress: true
        });
        // Store the validator set.
        IAvalancheValidatorSetManager(validatorSetManagerContract).setPartialValidatorSet(
            pChainID, partialValidatorSet
        );
        // pre-allocate the storage slots
        ValidatorSet memory validatorSet = ValidatorSet({
            avalancheBlockchainID: initialValidatorSetData.avalancheBlockchainID,
            pChainHeight: initialValidatorSetData.pChainHeight,
            pChainTimestamp: initialValidatorSetData.pChainTimestamp,
            validators: new Validator[](initialValidatorSetData.totalValidators),
            totalWeight: 0
        });
        // Store the validator set.
        IAvalancheValidatorSetManager(validatorSetManagerContract).setValidatorSet(
            pChainID, validatorSet
        );
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
        require(message.message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        // Check the registration status of the source blockchain ID and verify the
        // message signature appropriately
        if (isRegistrationInProgress(message.message.sourceBlockchainID)) {
            // check if we are interrupting an existing registration
            revert("Can't register to a blockchain ID while another registration is in progress");
        } else if (!isRegistered(message.message.sourceBlockchainID)) {
            // N.B. this message should be signed by the currently registered P-chain validator set
            verifyICMMessage(message, pChainID);
        } else {
            // This blockchain ID has an existing validator set registered to it which should
            // have signed this message
            verifyICMMessage(message, message.message.sourceBlockchainID);
        }

        // Validate that the source address is empty.
        require(message.message.sourceAddress == address(0), "Source address must be empty");
        // Parse and validate the validator set payload
        (
            ValidatorSetMetadata memory validatorSetMetadata,
            Validator[] memory validators,
            uint64 validatorWeight
        ) = IAvalancheValidatorSetManager(validatorSetManagerContract).parseValidatorSetMetadata(
            message, shardBytes
        );
        bytes32 avalancheBlockchainID = validatorSetMetadata.avalancheBlockchainID;
        // This validator set is sharded
        if (validatorSetMetadata.shardHashes.length > 1) {
            // pre-allocate enough storage for the whole validator set
            Validator[] memory valSet = new Validator[](validatorSetMetadata.totalValidators);
            for (uint256 i = 0; i < validators.length; i++) {
                valSet[i] = validators[i];
            }

            // initialize the partial validator set and store it
            PartialValidatorSet memory partialValidatorSet = PartialValidatorSet({
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                shardHashes: validatorSetMetadata.shardHashes,
                shardsReceived: 1,
                validators: valSet,
                numValidators: validators.length,
                partialWeight: validatorWeight,
                inProgress: true
            });
            IAvalancheValidatorSetManager(validatorSetManagerContract).setPartialValidatorSet(
                avalancheBlockchainID, partialValidatorSet
            );
            if (!isRegistered(message.message.sourceBlockchainID)) {
                // pre-allocate storage for the completed validator set
                ValidatorSet memory validatorSet = ValidatorSet({
                    avalancheBlockchainID: avalancheBlockchainID,
                    pChainHeight: validatorSetMetadata.pChainHeight,
                    pChainTimestamp: validatorSetMetadata.pChainTimestamp,
                    validators: new Validator[](validatorSetMetadata.totalValidators),
                    totalWeight: 0
                });
                // Store the validator set.
                IAvalancheValidatorSetManager(validatorSetManagerContract).setValidatorSet(
                    avalancheBlockchainID, validatorSet
                );
            }
        } else {
            // Construct the validator set
            ValidatorSet memory validatorSet = ValidatorSet({
                avalancheBlockchainID: avalancheBlockchainID,
                validators: validators,
                totalWeight: validatorWeight,
                pChainHeight: validatorSetMetadata.pChainHeight,
                pChainTimestamp: validatorSetMetadata.pChainTimestamp
            });
            // Store the validator set.
            IAvalancheValidatorSetManager(validatorSetManagerContract).setValidatorSet(
                avalancheBlockchainID, validatorSet
            );
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
        uint64 shardsReceived = IAvalancheValidatorSetManager(validatorSetManagerContract)
            .getShardsReceived(avalancheBlockchainID);
        require(shardsReceived + 1 == shard.shardNumber, "Received shard out of order");

        bytes32 shardHash = IAvalancheValidatorSetManager(validatorSetManagerContract).getShardHash(
            avalancheBlockchainID, shardsReceived
        );
        require(shardHash == sha256(shardBytes), "Unexpected shard hash");

        IAvalancheValidatorSetManager(validatorSetManagerContract).applyShard(shard, shardBytes);
        if (!isRegistrationInProgress(shard.avalancheBlockchainID)) {
            emit ValidatorSetUpdated(shard.avalancheBlockchainID);
        }
    }

    /**
     * @notice Updates an existing validator set by applying a diff.
     * @dev This function ASSUMES the number of changes in the `diff` 
     * (validators added + removed + modified) is small relative to the total validator set size.
     * @param message The ICM message containing the ValidatorSetDiffPayload.
     */
    function updateValidatorSetWithDiff(
        ICMMessage calldata message
    ) external {

        // Safety checks 
        require(message.message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        bytes32 chainID = message.message.sourceBlockchainID;
        require(isRegistered(chainID), "Validator set not registered");
        require(!isRegistrationInProgress(chainID), "Cannot apply diff with another registration in progress");
        verifyICMMessage(message, chainID);
        
        // Apply the diff
        ValidatorSetDiffPayload memory diff = ValidatorSets.parseValidatorSetDiffPayload(message.message.payload);
        ValidatorSet memory currentValidatorSet = IAvalancheValidatorSetManager(validatorSetManagerContract).getValidatorSet(chainID);
        require(diff.currentHeight > currentValidatorSet.pChainHeight, "Invalid blockchain height");
        (Validator[] memory newValidators, uint64 newWeight) = ValidatorSets.applyValidatorSetDiff(currentValidatorSet.validators, diff);

        // Update state
        ValidatorSet memory newValidatorSet = ValidatorSet({
            avalancheBlockchainID: chainID,
            pChainHeight: diff.currentHeight, 
            pChainTimestamp: diff.currentTimestamp, 
            validators: newValidators,
            totalWeight: newWeight
        });
        IAvalancheValidatorSetManager(validatorSetManagerContract).setValidatorSet(
            chainID, 
            newValidatorSet
        );

        emit ValidatorSetUpdated(chainID);
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
        return IAvalancheValidatorSetManager(validatorSetManagerContract).isRegistered(pChainID);
    }

    /**
     * @notice Check if a **complete** validator set is registered (not just a partial).
     */
    function isRegistered(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return IAvalancheValidatorSetManager(validatorSetManagerContract).isRegistered(
            avalancheBlockchainID
        );
    }

    /**
     * @notice Check if a validator set is registered but awaiting further updates.
     */
    function isRegistrationInProgress(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return IAvalancheValidatorSetManager(validatorSetManagerContract).isRegistrationInProgress(
            avalancheBlockchainID
        );
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
        require(message.message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        ValidatorSetSignature memory sig =
            ValidatorSets.parseValidatorSetSignature(message.attestation);
        ValidatorSet memory validatorSet = IAvalancheValidatorSetManager(
            validatorSetManagerContract
        ).getValidatorSet(avalancheBlockchainID);
        require(
            ValidatorSets.verifyValidatorSetSignature(sig, message.rawMessageBytes, validatorSet),
            "Could not verify ICM message: Signature checks failed"
        );
    }
}
