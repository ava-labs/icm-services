Teleporter is a protocol of interoperability messaging primitives that provides at most one delivery of authenticated ICM messages.  For interoperability between Avalanche L1s and a description of the actual protocol, see [here](../../icm-contracts/contracts/teleporter/README.md). We wish to extend this protocol to external EVM chains, but we cannot use the current contract (`ITeleporterMessenger`). Here we describe the new contracts we will need and the overall architecture.

## Decoupling message passing from Teleporter

 The new `TeleporterV2` contract will implement a new `ITeleporterMessengerV2` interface. The `TeleporterV2` contract should not concern itself with how messages are sent or authenticated. This will depend on the two chains being connected and will be handled by other contracts. Instead, its job is to ensure at most once delivery of messages and to provide the API on which further protocols (such as `ICTT`) are built.

This differs from the `TeleporterMessenger` contracts which had direct access to a pre-compile contract called `Warp`which provides it with the means of retrieving and authenticating messages.  Instead, the `TeleporterV2` contract will need access to a pair of contract addresses that each implement one of the two new interfaces:  `IMessageVerifier` and `IMessageSender`. This allows a `TeleporterV2` contract living on a specific blockchain to be configured to handle interfacing with a specific counter-party blockchain in a generic way.

```solidity
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
    bytes UnsignedMessagebytes;
    // arbitrary data. Used by `IMessageVerifier` contracts to
    // to authenticate the message
    bytes attestation;
}

struct TeleporterMessageInput {
    bytes32 destinationBlockchainID;
    address destinationAddress;
    // This field may not be necessary depending on the finalized design.
    address verifierAddress;
    TeleporterFeeInfo feeInfo;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    bytes message;
}

struct TeleporterMessage {
    uint256 messageNonce;
    address originSenderAddress;
    address originTeleporterAddress;
    bytes32 destinationBlockchainID;
    address destinationAddress;
    // This field may not be necessary depending on the finalized design.
    address verifierAddress;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    TeleporterMessageReceipt[] receipts;
    bytes message;
}

interface IMessageVerifier {
    function verifyMessage(ICMMessage message) external returns (bool);
}

interface IMessageSender {
    /**
    * @notice This function emits the correct type of event to be
    * consumed by a relayer. 
    */
    function sendMessage(TeleporterMessageInput message) external;
}

interface ITeleporterMessengerV2 {
    
    function receiveCrossChainMessage(
        ICMMessage message,
        address relayerRewardAddress
    ) external;
    
    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (bytes32);
	
	...
}
```

Note the attestation field of the `ICMMessage` is a generic blob of bytes, giving it flexibility to be a BLS signature, a certificate for CCIP, a tuple including a ZK proof, state root, and Merkle path to an event for a Boundless ZK Proof, or something entirely different.

This design maintains interface compatibility with applications already built on top of the Teleporter protocol. The architecture looks as follows.

```mermaid
classDiagram
  TeleporterV2 <|..	ITeleporterMessengerV2
	ITeleporterReceiver..|> Application
	ITeleporterReceiver: receiveTeleporterMessage(...)
	ITeleporterMessengerV2: receiveCrossChainMessage(...)
	ITeleporterMessengerV2: sendCrossChainMessage(...)
	Application-->TeleporterV2: sendCrossChainMessage(...)
	TeleporterV2-->Application: receiveTeleporterMessage(...)
	TeleporterV2-->IMessageVerifier: verifyMessage(...)
	TeleporterV2-->IMessageSender: sendMessage(...)
	IMessageVerifier: verifyMessage(...)
	IMessageSender: sendMessage(...)
```

```mermaid
sequenceDiagram
autonumber
actor relayer
note left of relayer:<br> Receive ICM message <br>flow<br>
relayer->>TeleporterV2: receiveCrossChainMessage(...)
TeleporterV2->>IMessageVerifier: verifyMessage(...)
IMessageVerifier-->>TeleporterV2: true / false
TeleporterV2->>Application: receiveTeleporterMessage(...)
```

```mermaid
sequenceDiagram
autonumber
Applicaton->>TeleporterV2: sendCrossChainMessage(...)
TeleporterV2->>IMessageSender: sendMessage(...)
actor relayer
note right of relayer: Send ICM Message flow
IMessageSender-) relayer: emit SendCrossChainMessage event
```

How a `TeleporterV2` contract decides which `IMessageVerifier` to use upon receiving a message depends on a pending design decision. The details of the two proposals can be found [here](icm_message_authentication.md). In either case, it is up to applications to decide the verification scheme they require for receiving messages.

## Message IDs

The `Teleporter` protocol issues a unique ID to each message to help ensure at-most-once delivery. It is used to make sure a message is not received twice as well as to store receipts of received messages. Relayers can send these receipts back to the sending application to communicate successful delivery of a message. It also ties messages to a specific `Teleporter` version.

Message IDs should be unique to each connection and contain nonces to distinguish messages along a given connection. They are a hash of several pieces of information pertaining to the message:

- The source blockchain ID
- The destination blockchain ID
- A nonce (unique per source blockchain)
- The destination `TeleporterV2` contract’s address
- The destination address