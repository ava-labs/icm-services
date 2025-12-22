// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

/**
 * @title SimpleValidatorSetRegistry
 * @notice A simplified validator set registry for testing purposes
 * @dev This contract stores validator sets without ICM message verification
 */
contract SimpleValidatorSetRegistry {
    struct Validator {
        bytes blsPublicKey;
        uint64 weight;
    }

    struct ValidatorSet {
        bytes32 avalancheBlockchainID;
        Validator[] validators;
        uint64 totalWeight;
        uint64 pChainHeight;
    }

    uint32 public immutable avalancheNetworkID;
    bytes32 public immutable avalancheBlockChainID;
    uint32 public nextValidatorSetID = 0;

    mapping(uint256 => ValidatorSet) private _validatorSets;

    event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID);
    event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID);

    constructor(uint32 _avalancheNetworkID, bytes32 _avalancheBlockChainID) {
        avalancheNetworkID = _avalancheNetworkID;
        avalancheBlockChainID = _avalancheBlockChainID;
    }

    /**
     * @notice Register a new validator set (no signature verification)
     * @param validators Array of validators with BLS public keys and weights
     * @param pChainHeight The P-chain height for this validator set
     * @return validatorSetID The ID of the registered validator set
     */
    function registerValidatorSet(
        Validator[] calldata validators,
        uint64 pChainHeight
    ) external returns (uint256) {
        require(validators.length > 0, "Validator set cannot be empty");

        uint256 validatorSetID = nextValidatorSetID++;
        ValidatorSet storage vs = _validatorSets[validatorSetID];
        vs.avalancheBlockchainID = avalancheBlockChainID;
        vs.pChainHeight = pChainHeight;

        uint64 totalWeight = 0;
        for (uint256 i = 0; i < validators.length; i++) {
            vs.validators.push(validators[i]);
            totalWeight += validators[i].weight;
        }
        vs.totalWeight = totalWeight;

        emit ValidatorSetRegistered(validatorSetID, avalancheBlockChainID);
        return validatorSetID;
    }

    /**
     * @notice Update an existing validator set (no signature verification)
     * @param validatorSetID The ID of the validator set to update
     * @param validators Array of validators with BLS public keys and weights
     * @param pChainHeight The P-chain height for this validator set
     */
    function updateValidatorSet(
        uint256 validatorSetID,
        Validator[] calldata validators,
        uint64 pChainHeight
    ) external {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");
        require(validators.length > 0, "Validator set cannot be empty");

        ValidatorSet storage currentVs = _validatorSets[validatorSetID];
        require(pChainHeight > currentVs.pChainHeight, "P-Chain height must be greater");

        // Clear existing validators
        delete _validatorSets[validatorSetID].validators;

        ValidatorSet storage vs = _validatorSets[validatorSetID];
        vs.pChainHeight = pChainHeight;

        uint64 totalWeight = 0;
        for (uint256 i = 0; i < validators.length; i++) {
            vs.validators.push(validators[i]);
            totalWeight += validators[i].weight;
        }
        vs.totalWeight = totalWeight;

        emit ValidatorSetUpdated(validatorSetID, avalancheBlockChainID);
    }

    /**
     * @notice Get the current (latest) validator set
     */
    function getCurrentValidatorSet() external view returns (ValidatorSet memory) {
        require(nextValidatorSetID > 0, "No validator sets exist");
        return _validatorSets[nextValidatorSetID - 1];
    }

    /**
     * @notice Get a validator set by ID
     */
    function getValidatorSet(uint256 validatorSetID) external view returns (ValidatorSet memory) {
        require(validatorSetID < nextValidatorSetID, "Validator set does not exist");
        return _validatorSets[validatorSetID];
    }
}

