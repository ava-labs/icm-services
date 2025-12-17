# Validator Set Registry

External EVM chains authenticate messages originating from Avalanche or Avalanche L1 by checking it is signed by a 
quorum of signatures of the relevant validator set. In order for an external EVM chain to know what the relevant
validator set is, a smart contract will be deployed that maintains a registry of such validator sets. This contract
is called the `AvalancheValidatorSetRegisty`.

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
A map from `avalancheBlockchainID` to the current validator set will be maintained in the contract. This blockchain ID
may belong to the P-chain or any Avalanche L1. We do not need to keep any validator set of a given blockchain ID other
than the most current as __this contract assumes that a quorum of signatures is always possible to be acquired for the
most current validator set__. How this assumption is ensured is not the concern of this contract. 

One special validator set is used to initialize this contract: The P-chain validator set which is current at the time
of contract instantiation. This will be used as the _root of trust_ as discussed in the [overview](summary.md).

A challenge in passing these validator sets to the registry contract is that they may constitute a data payload in excess
of blockchain transaction limits. As such, it may be necessary to break up the serialized validator set data across
 multiple transactions. This will certainly be the case when initializing the contract with a P-chain validator set.

Each of these transaction should consist of two pieces of data to pass into the registry contract:
An ICM message and a shard of the serialized data set. The ICM message should contain a validator set payload given
by the following Solidity data type:
```solidity
struct ValidatorSetStatePayload {
    bytes32 avalancheBlockchainID;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
    bytes32 validatorSetHash;
    uint64 shardNumber;
}
```
This payload should be identical for each shard, except for the `shardNumber` field. The contained data identifies which validator
set we are updating. The `validatorSetHash` acts as a commitment to full validator set we intend to communicate. The 
contract will need to buffer the shards internally until the hash of all shards (sorted by `shardNumber`) matches the 
`validatorSetHash` field. The data shards are simply an array of bytes. Each of these transactions should be signed by 
the current validator set of the chain identified by `avalancheBlockchainID` (if updating the set) or by the current 
P-chain validator set (if registering the L1 for the first time).

This scheme allows the sharding of update sets, which will be necessary for initializing the contract, but may not be
necessary otherwise. Many L1s will not have large validator sets in the first place; the main concern is updating the
P-chain. Once a validator set is registered, the data shards above can be switched to diffs, which we expect to be 
significantly smaller. That is to say, we can specify that the register function expects data shards to be a raw 
validator set, but the update function expects them to be diffs.

A challenge when applying diffs is that all Avalanche validator sets have a canonical ordering. The diffs must include
sufficient information so that after the new validator set is reconstructed, its ordering is also computable. _A final
decision about this issue and how best to resolve it has not yet been made_.

## Registering and updating new L1s

As alluded to in the previous section, only the P-chain will be registered on contract initialization. Any L1 wishing to
be registered will need to call `registerValidatorSet` contract function. This message should be signed by the current
P-chain validator set, after which it will be added to registry under its blockchain ID.

To update the validator set for a given blockchain ID, the `updateValidatorSet` contract function is used instead. The
message passed to this function is not signed by the P-chain validators, but by the current validator set for the
blockchain ID. 

## Verifying messages

A crucial part of the `AvalancheValidatorSetRegistry` contract is to authenticate messages received by `TeleporterMessenger` 
contracts on external EVM chains.  The `TeleporterMessenger` does this by calling into the `verifyICMMessage` function with
an ICM message, which is described by the following Solidity data types:
```solidity
struct ICMSignature {
    bytes signers;
    bytes signature;
}

struct ICMUnsignedMessage {
    // used to distinguish between main- and testnets
    uint32 avalancheNetworkID;
    bytes32 avalancheSourceBlockchainID;
    bytes payload;
}

struct ICMMessage {
    ICMUnsignedMessage unsignedMessage;
    bytes unsignedMessageBytes;
    ICMSignature signature;
}
```

The registry uses the `avalancheSourceBlockchainID` to look up a validator set. The `signers` is a bit set indicating which
validators have signed this message, utilizing the canonical validator ordering. The signers should represent a quorum
of staking weight for the message to be verified. The `signature` is an aggregate BLS signature of the validators specified
by the `signers` field. BLS computations requires an EVM chain to be on EVM version `prague` or later. 

In order for the `TeleporterMessenger` contracts to receive ICM messages on an external EVM chain, the current contracts
will need to be modified to handle the above flow. This is described further in the [Teleporter Messenger changes](teleporter_messenger_changes.md)
section. 

If the verification is successful, we transform this message into a `WarpMessage`, which is what is expected by
the `TeleportMessenger` contract. 
