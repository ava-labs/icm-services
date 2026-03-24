# Authenticating ICM Messages

Authenticating ICM messages will be handled by contracts implementing the `IMessageVerifier` interface, as described in the [Teleporter architecture](teleporter_contracts.md). The protocol is completely agnostic to the details of how messages are authenticated; it is up to applications to decide how to authenticate messages. Applications do this by using a `TeleporterMessengerV2`instance implementing an authentication scheme they support.

For internal interoperability, all message authentication is handled by the `Warp` pre-compile contract. New contracts will be written to wrap this pre-compile so that it is compatible with the `TeleporterV2` contract / architecture. Furthermore, for external EVM chains, it is expected that many applications will authenticate messages by checking a quorum of validator signatures. The contract implementing this logic is detailed [here](origin_avalanche/validator_set_registry.md). Other verification schemes may be developed, and the architecture is designed to support them.

## Determining the `IMessageVerifier` contract
As detailed in the description of the [new Teleporter architecture](teleporter_contracts.md), message authentication is handled by a contract implementing the `IMessageVerifier` interface. The architecture did not specify how a specific `TeleporterMessengerV2` contract knows which contract to use. We describe the here in further detail.

 `TeleporterMessengerV2` will be deployed via Nick's method; its constructor taking the address of a contract implementing the  `IAdapter` interface. Note that this requires the `IAdapter` contract to also be deployed via Nick's method. To support multiple verification schemes, multiple `TeleporterMessengerV2` contracts may be deployed. To summarize:
 * The same implementations of an `IAdapter` contract are deployed to the same address on each chain.
 * All `TeleporterV2` contracts implementing the same scheme are deployed to the same address on each chain.

Two `TeleporterV2` contracts may talk to one another if they use the same `IAdapter` contract. When a TeleporterV2 contract receives a message, it checks if it originated from the same contract address as its own. If so, it knows that it is part of the same scheme and can verify the message using the hard-coded `IAdapater` implementation.

Note that an `IAdapater` implementation for a particular scheme may need to perform different logic depending on which chain it lives. So a switch statement should be included in its logic to dispatch to the right method depending on which chain it is being called. While many of the code paths will be dead on each chain, it means that the same code is deployed on-chain for a particular scheme.

When an application is called by a `TeleporterMessengerV2` contract, it checks that the `msg.sender` variable contains the address of a `TeleporterMessengerV2` contract using an `IAdapter` contract that it supports.
