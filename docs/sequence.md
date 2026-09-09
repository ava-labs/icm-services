# Processing the Warp messages in a new block
Illustration of the sequence of events triggered by a new block containing Warp messages to relay. The Subscriber is only notified of blocks that contain logs matching the listener's message protocol; the blocks in between, which contain no messages, are checkpointed together with the next block that does.
```mermaid
sequenceDiagram
    participant Subscriber
    participant Listener
    participant ApplicationRelayer
    participant MessageHandler
    participant AppRequestNetwork
    participant DestinationClient
    participant CheckpointManager
    participant RelayerDatabase

    Subscriber->>Listener : (async) New block with Warp logs
    activate Listener
    Listener->>Listener : Extend block range over preceding empty blocks
    Listener->>Listener : Create handlers for each message
    Listener-->>ApplicationRelayer : (async) Pass message handlers for block range
    deactivate Listener
    activate ApplicationRelayer
    par foreach Warp message in block
        ApplicationRelayer->>MessageHandler : ShouldSendMessage
        activate MessageHandler
        MessageHandler->>ApplicationRelayer : true

        par foreach canonical validator
            ApplicationRelayer->>AppRequestNetwork : Request signatures
            activate AppRequestNetwork
            AppRequestNetwork->>ApplicationRelayer : signature
        end
        ApplicationRelayer->>ApplicationRelayer : Aggregate signatures

        ApplicationRelayer->>MessageHandler : SendMessage
        activate MessageHandler
        MessageHandler->>DestinationClient : SendTx
        activate DestinationClient
        DestinationClient->>MessageHandler : txHash
        MessageHandler->>MessageHandler : Wait for receipt
        MessageHandler->>ApplicationRelayer : return
    end
    ApplicationRelayer->>CheckpointManager : Stage block range
    deactivate ApplicationRelayer
    CheckpointManager-->>RelayerDatabase : (async) Write height
```