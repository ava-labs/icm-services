# Validator Set Registry

Many applications on external EVM chains will want to authenticate messages originating from Avalanche or Avalanche L1s by checking that they are signed by a quorum of the relevant validator set. In order for an external EVM chain to know what the relevant validator set is, a smart contract is deployed that maintains a registry of such validator sets. That contract is the `MerkleValidatorSetRegistry`.

Rather than storing validator sets on-chain, the registry keeps a single Merkle root commitment per Avalanche blockchain ID. The validator sets themselves live off-chain, and the signers for any given message are supplied in calldata together with a Merkle multi-proof binding them to the stored root. This keeps both storage and update costs constant in the size of the validator set.

`ZKValidatorSetRegistry` implements the same `IMerkleValidatorSetRegistry` interface, but authorizes commitment updates with a zero-knowledge proof of P-chain state rather than with a BLS-signed ICM message.

## Validator set commitments

A validator set is represented on-chain by a commitment, represented in Solidity by the data types:

```solidity
struct Validator {
    bytes blsPublicKey;
    uint64 weight;
}

struct ValidatorSetMerkleCommitment {
    bytes32 avalancheBlockchainID;
    bytes32 root;
    uint64 totalWeight;
    uint64 pChainHeight;
    uint64 pChainTimestamp;
}
```

A map from `avalancheBlockchainID` to the current commitment is maintained in the contract. This blockchain ID may belong to the P-chain or any Avalanche L1. We do not need to keep any commitment for a given blockchain ID other than the most current one, as __this contract assumes that a quorum of signatures is always possible to acquire for the most current validator set__. How this assumption is ensured is not the concern of this contract.

The `root` is the root of a Merkle tree over the validator set, with the validators sorted in ascending lexicographic order of their uncompressed BLS public keys — the same canonical ordering used by the signature aggregator. `totalWeight` is the sum of the validators' weights and is what quorum is measured against.

Because a commitment is a fixed-size struct, an update always fits in a single transaction. There is no sharding, no partial state, and no multi-transaction registration flow.

## Verifying messages

A crucial part of the registry is to authenticate messages received by `TeleporterMessengerV2` contracts on external EVM chains. The `TeleporterMessengerV2` contract does this by calling into the `verifyICMMessage` function with an ICM message, which is described by the following Solidity data types:

```solidity
struct ValidatorSetMerkleAttestation {
    Validator[] signers;
    bytes32[] proof;
    bool[] proofFlags;
    bytes aggregateBlsSig;
}

struct ICMMessage {
    // The serialized bytes of raw message. The data and serialization formats
    // for this data will be app / contract specific
    bytes rawMessage;
    uint32 sourceNetworkID;
    bytes32 sourceBlockchainID;
    // This should contain a serialized `ValidatorSetMerkleAttestation`
    bytes attestation;
}
```

The registry uses the `sourceBlockchainID` to look up the corresponding commitment. The `attestation` deserializes to a `ValidatorSetMerkleAttestation`. Verification then proceeds in three steps:

1. The `signers` are hashed into Merkle leaves and checked against the stored `root` using the multi-proof given by `proof` and `proofFlags`. This proves that every claimed signer really is a member of the committed validator set, without the contract ever holding the set.
2. The summed weight of the `signers` is checked to be at least the quorum fraction (67/100) of the commitment's `totalWeight`.
3. `aggregateBlsSig` is verified as an aggregate BLS signature over the message by the aggregated public keys of the `signers`.

The `signers` must be supplied in increasing lexicographic order by BLS public key. BLS computations require an EVM chain to be on EVM version `prague` or later.

## Registering and updating a validator set

Both first registration and subsequent updates go through the same `registerValidatorSet` function, which takes an ICM message and the blockchain ID of the validator set that signed it:

```solidity
function registerValidatorSet(ICMMessage calldata message, bytes32 signingChainID) external;
```

The message payload is a serialized `ValidatorSetMerkleCommitment`. Which case applies is determined by whether the payload's blockchain ID is already registered:

- The first registration for a given blockchain ID must be signed by the P-chain validator set.
- Subsequent updates must be signed by the validator set currently registered for that blockchain ID. When the registry was deployed with `allowPChainFallback` enabled, the P-chain validator set may also sign updates.

A `ValidatorSetRegistered` event is emitted once the new commitment is stored. Because the whole commitment arrives in one message, an update either takes effect completely or reverts — there is no state in which a blockchain ID has a partially applied commitment.

### Initializing the contract

One special validator set is used to initialize this contract: the primary network (P-chain) validator set which is current at the time of contract instantiation. This is the _root of trust_.

The P-chain commitment is passed directly to the constructor:

```solidity
constructor(
    uint32 avalancheNetworkID_,
    bytes32 pChainID_,
    bytes32 pChainGenesisRoot,
    uint64 pChainTotalWeight,
    uint64 pChainHeight,
    uint64 pChainTimestamp,
    bool allowPChainFallback_
)
```

This commitment is not delivered via an ICM message and is not authenticated by the contract. If deploying via Nick's method, users can determine whether a registry contract was deployed with the correct input by checking the constructor arguments. Unlike a sharded registry, the contract is fully usable as soon as it is deployed — no follow-up transactions are needed to populate the initial set.
