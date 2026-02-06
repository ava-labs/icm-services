# Validator Set Registry

Many applications on external EVM chains will want to authenticate messages originating from Avalanche or Avalanche L1s by checking it is signed by a quorum of the relevant validator set. In order for an external EVM chain to know what the relevant validator set is, a smart contract will be deployed that maintains a registry of such validator sets. This contract is called the `AvalancheValidatorSetRegisty`.

## Validator sets

A validator set consists of several pieces of data, represented in Solidity by the data types: 
```solidity
struct Validator {
    bytes blsPublicKey;
    uint64 weight;
}

struct ValidatorSet {
    bytes32 avalancheBlockchainID;
    Validator[] validators;
    uint64 totalWeight;
    uint64 pChainHeight;
    uint64 pChainTimeStamp;
}
```

A map from `avalancheBlockchainID` to the current validator set will be maintained in the contract. This blockchain ID may belong to the P-chain or any Avalanche L1. We do not need to keep any validator set of a given blockchain ID other than the most current as __this contract assumes that a quorum of signatures is always possible to be acquired for the most current validator set__. How this assumption is ensured is not the concern of this contract.

A challenge in passing these validator sets to the registry contract is that they may constitute a data payload in excess of blockchain transaction limits. As such, the contract should support the ability to break up the data necessary to update the validator set across multiple transactions. This will certainly be the case when initializing the contract with a primary network validator set, for instance.

To support this, there should be two main endpoints for updates to be passed to: registering a validator set and updating it. Registering a validator set is an authorized cryptographic commitment to the validator set being uploaded. If it cannot be populated in a single transaction, the update endpoint may be used as necessary to add the remaining data.


## Verifying messages

A crucial part of the `AvalancheValidatorSetRegistry` contract is to authenticate messages received by `TeleporterV2` contracts on external EVM chains.  The `TeleporterMessenger` does this by calling into the `verifyICMMessage` function with an ICM message, which is described by the following Solidity data types:
```solidity
struct ICMSignature {
    bytes signers;
    bytes signature;
}

struct ICMUnsignedMessage {
    // used to distinguish between mainnet and various testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // The address that sent the message
    address sourceSenderAddress;
    // This field may not be necessary depending on the finalized design.
    address verifierAddress;
    bytes payload;
}

struct ICMMessage {
    ICMUnsignedMessage unsignedMessage;
    bytes unsignedMessageBytes;
    // This should contain a serialized `ICMSignature`
    bytes attestation;
}
```

The registry uses the `sourceBlockchainID` to look up a validator set. The `attestation` should deserialize to an `ICMSinature` instance. The `signers` field is a bit set indicating which validators have signed this message, utilizing the canonical validator ordering. The signers should represent a quorum of staking weight for the message to be verified. The `signature` is an aggregate BLS signature of the validators specified by the `signers` field. BLS computations require an EVM chain to be on EVM version `prague` or later.

## Registering a validator set

When an Avalanche L1 wishes to add one of its validator sets to the registry, it will need to call the `registerValidatorSet` function. This function expects an ICM message along with a byte payload. The ICM message should contain a `ValidatorSetMetadata` instance as its payload, defined as:
```solidity
struct ValidatorSetMetadata {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    uint64 totalValidators;
    bytes32[] shardHashes;
}
```
This data will be used as a cryptographic commitment to the validator set being registered. The sha256 hash of the serialized validator set should match the `validatorSetHash` field. The shard hashes are the hashes of each chunk of data that will be passed to the `updateValidatorSet` function. The included byte payload to the `registerValidatorSet` function is the first shard of the overall update.

If this L1 does not currently have one of its validator sets registered, the ICM message containing the metadata should be signed by the current primary network validator set registered to the contract. Otherwise, the current validator registered to the L1 should sign the message.

If there is only one shard for the update, this function will place the new validator set into the internal mapping as the current validator set for the given blockchain ID. Otherwise, this action will be postponed until all shards have been received.

If a validator set is registered and only contains a partial amount of the requisite data, attempts to register a validator set to the same blockchain ID will fail. 

## Updating a validator set

If an Avalanche L1 has registered a validator set that requires more than one shard, the remaining data should be passed to the `updateValidatorSet` function with as many transactions as necessary. This function expects `ValidatorSetShard` instance and a byte payload. 
```solidity
struct ValidatorSetShard {
    uint64 shardNumber;
    bytes32 avalancheBlockchainID;
}
```

This function should check that the byte payload's hash matches the `shardHashes` at index `shardNumber`. If so, this data can be applied to the partial validator set. How this is done is up to the specific implementation. It may be a serialization of validators or a partial diff. If this is the last shard, the validator set will be updated in the registry mapping as the most current.

Note that this implies a few things. The first is that shards must be delivered in the expected order. Secondly, it should not be possible to successfully call this function if the last registered metadata corresponds to a validator set that has been fully populated.

Lastly, because the first shard was passed as part of the `registerValidatorSet` function, the first shard passed to `updateValidatorSet` should have shard number 2.

### Initializing the contract

One special validator set is used to initialize this contract: The primary network validator set which is current at the time of contract instantiation. This will be used as the _root of trust_.

This follows a nearly identical flow as registering and updating validators described above. In the constructor, a `ValidatorSetMetadata` instance is passed, committing to the initial validator set. This is not passed via an ICM message and is not authenticated. If deploying via Nick's method, users and determine if an `AvalancheValidatorSetRegistry` contract was deployed with the correct input.  

After deployment, `updateValidatorSet` must be called until the initial validator set is fully populated. Until then, no validation or otherwise stateful function should be available.

It should be noted that no shards will be passed to the constructor, so the first call to `updateValidatorSet` should contain the shard with value 1. This is contrast to the above flow where the first call to `updateValidatorSet` should contain the shard with value 2.