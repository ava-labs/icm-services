// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {IAvalancheValidatorSetManager} from "./interfaces/IAvalancheValidatorSetManager.sol";
import {ICMMessage} from "../common/ICM.sol";
import {
    PartialValidatorSet,
    ValidatorSetMetadata,
    Validator,
    ValidatorSet,
    ValidatorSetShard,
    ValidatorSets
} from "./utils/ValidatorSets.sol";

// This contract stores complete and partial validator sets as mappings keyed by blockchain ID.
//
// It further specifies that shards of validator sets is a serialized subsequence of
// the entire validator set that has been cryptographically committed to.
contract AvalancheValidatorSetManager is IAvalancheValidatorSetManager, Initializable {
    // This contract should only be modified by a particular contract so that malicious
    // parties cannot alter the data stored here
    address private _owner;

    // Mapping of Avalanche blockchain IDs to their complete validator set data.
    mapping(bytes32 => ValidatorSet) internal _validatorSets;

    // Mapping of Avalanche blockchain IDs to a subset of said validator set.
    // Used to allow updating validator sets across multiple transactions
    mapping(bytes32 => PartialValidatorSet) internal _partialValidatorSets;

    modifier onlyOwner() {
        require(msg.sender == _owner, "An unauthorized address attempted to call this function");
        _;
    }

    function initialize(
        address owner
    ) external initializer {
        _owner = owner;
    }

    function setValidatorSet(
        bytes32 avalancheBlockchainID,
        ValidatorSet memory validatorSet
    ) external onlyOwner {
        _validatorSets[avalancheBlockchainID] = validatorSet;
    }

    function setPartialValidatorSet(
        bytes32 avalancheBlockchainID,
        PartialValidatorSet memory partialValidatorSet
    ) external onlyOwner {
        _partialValidatorSets[avalancheBlockchainID] = partialValidatorSet;
    }

    /**
     * @dev Applies a set of validators to a partial set that has been registered.
     */
    function applyShard(
        ValidatorSetShard calldata shard,
        bytes memory shardBytes
    ) external onlyOwner {
        bytes32 avalancheBlockchainID = shard.avalancheBlockchainID;
        // Parse the validators.
        (Validator[] memory validators, uint64 validatorWeight) =
            ValidatorSets.parseValidators(shardBytes);
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
    ) external view returns (ValidatorSetMetadata memory, Validator[] memory, uint64) {
        // Parse the validator set state payload.
        ValidatorSetMetadata memory validatorSetMetadata =
            ValidatorSets.parseValidatorSetMetadata(icmMessage.message.payload);
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

    function isRegistered(
        bytes32 avalancheBlockchainID
    ) external view returns (bool) {
        return _validatorSets[avalancheBlockchainID].totalWeight != 0;
    }

    function isRegistrationInProgress(
        bytes32 avalancheBlockchainID
    ) external view returns (bool) {
        return _partialValidatorSets[avalancheBlockchainID].inProgress;
    }

    function getShardsReceived(
        bytes32 avalancheBlockchainID
    ) external view returns (uint64) {
        return _partialValidatorSets[avalancheBlockchainID].shardsReceived;
    }

    function getShardHash(
        bytes32 avalancheBlockchainID,
        uint256 index
    ) external view returns (bytes32) {
        return _partialValidatorSets[avalancheBlockchainID].shardHashes[index];
    }

    function getValidatorSet(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSet memory) {
        return _validatorSets[avalancheBlockchainID];
    }

    /**
     * @notice Get the owner of this contract
     */
    function getOwner() external view returns (address) {
        return _owner;
    }
}
