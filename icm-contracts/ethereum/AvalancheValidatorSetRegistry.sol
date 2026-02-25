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

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

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

    // Mapping of Avalanche blockchain IDs to a partially updated validator set.
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

        PartialValidatorSet storage partialSet = _partialValidatorSets[pChainID];
        partialSet.pChainHeight = initialValidatorSetData.pChainHeight;
        partialSet.pChainTimestamp = initialValidatorSetData.pChainTimestamp;
        partialSet.shardHashes = initialValidatorSetData.shardHashes;
        partialSet.inProgress = true;

        ValidatorSet storage valSet = _validatorSets[pChainID];
        valSet.avalancheBlockchainID = initialValidatorSetData.avalancheBlockchainID;
        valSet.pChainHeight = initialValidatorSetData.pChainHeight;
        valSet.pChainTimestamp = initialValidatorSetData.pChainTimestamp;
    }

    /**
     * @notice Registers a new validator set for a specific Avalanche blockchain ID.
     *
     * Emits an event that a new set has been registered.
     * @dev It may be the case that the entire validator set cannot be passed into this function.
     * If a partial validator set is provided, the chain is still considered registered. To pass
     * the rest of the validator set data, `updateValidatorSet` should be called instead.
     *
     * If this function is called to register a new validator set for chain for a which a partial
     * set exists, this function will revert.
     * @param message The ICM message containing the validator set to register. The message must
     * be signed by the relevant authorizing party
     * @param shardBytes The first shard of the data needed to populate the newly registered
     * validator set.
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes calldata shardBytes
    ) external {
        // Check the network ID and signature
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        // Check that we are not interrupting an existing registration
        if (isRegistrationInProgress(message.sourceBlockchainID)) {
            // check if we are interrupting an existing registration
            revert("A registration is already in progress");
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
        uint256 numValidators = validators.length;

        // This validator set is sharded
        if (validatorSetMetadata.shardHashes.length > 1) {
            // initialize the partial validator set and store it
            PartialValidatorSet storage partialSet =
                _partialValidatorSets[validatorSetMetadata.avalancheBlockchainID];
            partialSet.pChainHeight = validatorSetMetadata.pChainHeight;
            partialSet.pChainTimestamp = validatorSetMetadata.pChainTimestamp;
            partialSet.shardHashes = validatorSetMetadata.shardHashes;
            partialSet.shardsReceived = 1;
            partialSet.partialWeight = validatorWeight;
            partialSet.inProgress = true;
            for (uint256 i = 0; i < numValidators;) {
                partialSet.validators.push(validators[i]);
                unchecked {
                    ++i;
                }
            }

            if (!isRegistered(message.sourceBlockchainID)) {
                ValidatorSet storage valSet =
                    _validatorSets[validatorSetMetadata.avalancheBlockchainID];
                valSet.avalancheBlockchainID = validatorSetMetadata.avalancheBlockchainID;
                valSet.pChainHeight = validatorSetMetadata.pChainHeight;
                valSet.pChainTimestamp = validatorSetMetadata.pChainTimestamp;
            }
        } else {
            // Store the validator set.
            ValidatorSet storage valSet = _validatorSets[validatorSetMetadata.avalancheBlockchainID];
            valSet.avalancheBlockchainID = validatorSetMetadata.avalancheBlockchainID;
            valSet.totalWeight = validatorWeight;
            valSet.pChainHeight = validatorSetMetadata.pChainHeight;
            valSet.pChainTimestamp = validatorSetMetadata.pChainTimestamp;
            for (uint256 i = 0; i < numValidators;) {
                valSet.validators.push(validators[i]);
                unchecked {
                    ++i;
                }
            }
        }
        emit ValidatorSetRegistered(avalancheBlockchainID);
    }

    /**
     * @dev The shard numbers contained the `shard` parameters must be processed in order
     * or else this function will fail.
     */
    function updateValidatorSet(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) external {
        require(
            isRegistrationInProgress(shard.avalancheBlockchainID), "Registration is not in progress"
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
     * is called when registering validator sets. It may also contain a (potentially partially) updated set of
     * the validators that are being registered. This is always considered to be the first shard of
     * the requisite data.
     *
     * @param icmMessage The ICM message containing the validator set metadata
     * @param shardBytes The serialized data used to construct the registered
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
     * @notice Verify the message. Does the following checks
     *   1. Check that the contracts root of trust has been initialized completely
     *   2. Intended for this network ID
     *   3. Has been signed by a quorum of the provided validator set
     * If any check fails, reverts
     */
    function verifyICMMessage(
        ICMMessage calldata message,
        bytes32 avalancheBlockchainID
    ) public view {
        require(pChainInitialized(), "No P-chain validator set registered.");
        require(isRegistered(avalancheBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");
        ValidatorSetSignature memory sig =
            ValidatorSets.parseValidatorSetSignature(message.attestation);
        bytes memory signedData = abi.encodePacked(
            message.sourceNetworkID, message.sourceBlockchainID, message.rawMessage
        );
        require(
            ValidatorSets.verifyValidatorSetSignature(
                sig, signedData, _validatorSets[avalancheBlockchainID]
            ),
            "Failed to verify signatures"
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
