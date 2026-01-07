# Authenticating ICM Messages

Authenticating ICM messages will be handled by contracts implementing the `IMessageVerifier` interface, as described in the [Teleporter architecture](teleporter_contracts.md). The `Teleporter` protocol is complelely agnostic to the details of how messages are authenticated. It is up to applications to decide how to authenticate messages.

For internal interoperability, all message authentication is handled by the `Warp` pre-compile contract. New contracts will be written to wrap this pre-compile so that it is compatible with the `TeleporterV2` contract / architecture. Furthermore, for external EVM chains, it is expected that many applications will authenticate messages by checking a quorum of validator signatures. The contract implementing this logic is detailed [here](origin_avalanche/validator_set_registry.md). Other verification schemes may be developed, and the architecture is designed to support them.

## Determining the `IMessageVerifier` contract
As detailed in the description of the [new Teleporter architecture](teleporter_contracts.md), message authentication is handled by a contract implementing the `IMessageVerifier` interface. The architecture did not specify how a specific `TeleporterV2` contract knows which contract to use, and currently there are two proposals for handling this. In this section, we describe them both.

### Hardcoding verifiers to `TeleporterV2` contracts

The first option is that a `TeleporterV2` contract has both `IMessageVerifier` and `IMessageSender` implementations hard-coded as part of its state. To support multiple verification schemes, multiple `TeleporterV2` contracts may be deployed.  

We define a particular choice of `IMessageVerifier` and `IMessageSender` implementations as a _scheme_. Two `TeleporterV2` contracts may talk to one another if they are part of the same scheme. To achieve this, we use Nick's method as follows:
 * The same implementations of an `IMessageVerifier` / `IMessageSender` contract are deployed to the same address on each chain.
 * All `TeleporterV2` contracts implementing the same scheme are deployed to the same address on each chain. 

When a TeleporterV2 contract receives a message, it checks if it originated from the same contract address as its own. If so, it knows that it is part of the same scheme and can verify the message using the hard-coded `IMessageVerifier` implementation.

Note that an `IMesssageVerifier` implementation for a particular scheme may need to perform different logic depending on which chain it lives. So a switch statement should be included in its logic to dispatch to the right method depending on which chain it is being called. While many of the code paths will be dead on each chain, it means that the same code is deployed on-chain for a particular scheme. The same is true of `IMessageSender` implementations.

When an application is called by a `TeleporterV2` contract, it checks that the `msg.sender` variable contains the address of a `TeleporterV2` contract using a scheme that it supports.

### Allowing applications to specify verifiers

An alternative proposal is to have a single `TelepoterV2` instance per chain, all deployed to the same address using Nick's method. `TeleporterMessage` instances will then contain an address to an `IMessageVerifier` contract that they can use to verify messages. 

When an application is called by a `TeleporterV2` contract, it checks that the `msg.sender` variable contains the address of the `TeleporterV2` contract and additionally checks that the `verifierAddress` field of the `ICMMessage` instance matches the address of an `IMessageVerifier` contract that it supports.