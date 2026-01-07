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

### Initializing the contract

One special validator set is used to initialize this contract: The primary network validator set which is current at the time of contract instantiation. This will be used as the _root of trust_.

A challenge in passing these validator sets to the registry contract is that they may constitute a data payload in excess of blockchain transaction limits. As such, it may be necessary to break up the serialized validator set data across multiple transactions. This will certainly be the case when initializing the contract with a primary network validator set.

Each of these transactions should consist of two pieces of data to pass into the registry contract: An ICM message and a shard of the serialized data set. The ICM message should contain a validator set payload given by the following Solidity data type: 
```solidity
struct ValidatorSetStatePayload {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    bytes32 validatorSetHash;
    uint64 shardNumber;
}
```
This payload should be identical for each shard, except for the `shardNumber` field. The contained data identifies which validator set we are updating. The `validatorSetHash` acts as a commitment to the full validator set we intend to communicate. The contract will need to buffer the shards internally until the hash of all shards (sorted by `shardNumber`) matches the  `validatorSetHash` field. The data shards are simply an array of bytes. Each of these transactions should be signed by the current validator set of the chain identified by `avalancheBlockchainID` (if updating the set) or by the current primary network validator set (if registering the L1 for the first time).

This scheme allows the sharding of update sets, which will be necessary for initializing the contract, but may not be necessary otherwise. Many L1s will not have large validator sets in the first place; the main concern is updating the P-chain. Once a validator set is registered, the data shards above can be switched to diffs, which we expect to be significantly smaller. That is to say, we can specify that the register function expects data shards to be a raw validator set, but the update function expects them to be diffs.

A challenge when applying diffs is that all Avalanche validator sets have a canonical ordering. The diffs must include sufficient information so that after the new validator set is reconstructed, its ordering is also computable. _A final decision about this issue and how best to resolve it has not yet been made_.

## Registering and updating new L1s

As alluded to in the previous section, only the P-chain will be registered on contract initialization. Any L1 wishing to be registered will need to call `registerValidatorSet` contract function. This message should be signed by the current primary network validator set, after which it will be added to the `registry` field under its blockchain ID.

To update the validator set for a given blockchain ID, the `updateValidatorSet` contract function is used instead. The message passed to this function is not signed by the primary network validators, but by the current validator set for the blockchain ID. 

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

The registry uses the `sourceBlockchainID` to look up a validator set. The `attestation` should deserialize to an `ICMSinature` instance. The `signers` field is a bit set indicating which validators have signed this message, utilizing the canonical validator ordering. The signers should represent a quorum of staking weight for the message to be verified. The `signature` is an aggregate BLS signature of the validators specified by the `signers` field. BLS computations requires an EVM chain to be on EVM version `prague` or later. 
  