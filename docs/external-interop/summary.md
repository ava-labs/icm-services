# External Interoperability

The goal of external interoperability is to provide a solution for sending messages to and from Avalanche and Avalanche 
L1s to other EVM chains like Ethereum. The underlying protocol should be decentralized and permissionless, allowing any 
party to verify the validity of messages sent and state transitions initiated by said messages. This is a design 
document laying out the protocol, architecture, contracts and code changes, trust assumptions, etc.

## Table of Contents
 - [Sending messages from Avalanche](origin_avalanche/summary.md)
   - [Validator set registry](origin_avalanche/validator_set_registry.md)
   - [Teleporter Messenger changes](origin_avalanche/teleporter_messenger_changes.md)
   - [Relayer Changes](origin_avalanche/relayer_changes.md)