# Sending messages from Avalanche

This section of the design document is concerned with messages originating from Avalanche. In this section, we detail how external EVM chains authenticate messages from Avalanche and the infrastructure for sending messages and maintaining  trust.

## Message protocol

Communication between Avalanche L1 (the so-called _"internal interoperability"_) uses the [ICM protocol](../../../icm-contracts/contracts/teleporter/README.md)  which is handled by a set of `Teleporter` contracts. Further protocols can be built on top of `Teleporter`, but this is  out of scope for this design doc. The goal of external interoperability is to support `Teleporter` contracts on external EVM  chains to leverage existing protocols. This will require changes to the `TeleporterMessenger` contracts which will be  deployed on external EVM chains; these changes are detailed in the [TeleporterMessesenger changes](teleporter_messenger_changes.md) section.

The primary difference between internal and external interoperability is how message authentication is accomplished. For internal interoperability, this was achieved used a pre-compiled smart contract called `Warp`. An alternative solution is necessary for external EVM chains because the `Warp` pre-compile is specific to Avalanche.

This will also require changing the relayers. They will need clients for external EVM chains and will need to call different contract functions and provide different data as well. There also new duties required of relayers other than just sending messages that will be required. This is further describe in the [Relayer changes](relayer_changes.md) section.

## Root of trust

The Avalanche P-chain is the root of trust when authenticating messages from the Avalanche C-chain or an Avalanche L1. Specifically, all authentication proofs should derive their authority from an initial  P-chain validator set known to the external EVM chain. While the publication of this initial set on an external EVM chain may be done anyone, it's validity is publicly auditable and thus a trustless procedure.

As P-chain validator sets change over time, a __chain of custody__ must be maintained whereby a quorum of the current P-chain  validator set must sign off on the next validator set (_N.B. we will often say that validator set signs a message, but  this should always be interpreted as meaning a quorum of signatures_). This new set will be published on the external EVM  chain if the signature check passes. From then on, this will be the current P-chain validator set whose signatures represent  authority of Avalanche.

For more details on how P-chain validator sets are published on external EVM chains, see the [Avalanche validator set registry](validator_set_registry.md)  section. As for how validator sets changes are communicated, see the [Updating validator sets]() section.

## Avalanche L1s

Communication to external EVM chains is designed to be possible from Avalanche L1s, as well the Avalanche P-chain. To achieve this, we extend the power to authorize messages from L1s to the validator set of said L1. Much like the P-chain, a current validator set of an L1 is published to the external EVM chain. Messages from an L1 should be signed by a quorum of this validator set in order to be authenticated.

However, the protocol differs when the very first set of validators for an L1 is published to the external EVM chain. This initial set must be signed by the current P-chain validator set that is published on the external EVM chain. After an L1 is registered to the external EVM chain in this way, it may update itself via a quorum of signatures of the currently  published L1 validator set, forming a chain of custody back to the originally registered set. Again, for more details, see the [Avalanche validator set registry]() section. 