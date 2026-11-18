# ICM Relayer

Reference relayer implementation for cross-chain Avalanche Warp Message delivery.

ICM Relayer listens for Warp message events on a set of source blockchains, and constructs transactions to relay the Warp message to the intended destination blockchain. The relayer does so by querying the source blockchain validator nodes for their BLS signatures on the Warp message, combining the individual BLS signatures into a single aggregate BLS signature, and packaging the aggregate BLS signature into a transaction according to the destination blockchain VM Warp message verification rules.

## Installation

### Dev Container & Codespace

To get started easily, we provide a Dev Container specification, that can be used using GitHub Codespace or locally using Docker and VS Code. [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers) are a concept that utilizes containerization to create consistent and isolated development environment. You can run them directly on Github by clicking **Code**, switching to the **Codespaces** tab and clicking **Create codespace on main**. Alternatively, you can run them locally with the extensions for [VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) or other code editors.

### Download Prebuilt Binaries

Prebuilt binaries are available for download from the [releases page](https://github.com/ava-labs/icm-services/releases).

The following commands demonstrate how to download and install the v0.2.13 release of the relayer on MacOS. The exact commands will vary by platform.

```bash
# Download the release tarball and checksums
curl -w '%{http_code}' -sL -o ~/Downloads/icm-relayer_0.2.13_darwin_arm64.tar.gz https://github.com/ava-labs/icm-services/releases/download/v0.2.13/icm-relayer_0.2.13_darwin_arm64.tar.gz
curl -w '%{http_code}' -sL -o ~/Downloads/icm-relayer_0.2.13_checksums.txt https://github.com/ava-labs/icm-services/releases/download/v0.2.13/icm-relayer_0.2.13_checksums.txt

# (Optional) Verify the checksums
cd ~/Downloads
# Confirm that the following two commands output the same checksum
grep "icm-relayer_0.2.13_darwin_arm64.tar.gz" "icm-relayer_0.2.13_checksums.txt" 2>/dev/null
shasum -a 256 "icm-relayer_0.2.13_darwin_arm64.tar.gz" 2>/dev/null

# Extract the tarball and install the relayer binary
tar -xzf icm-relayer_0.2.13_darwin_arm64.tar.gz
sudo install icm-relayer /usr/local/bin
```

_Note:_ If downloading the binaries through a browser on MacOS, the browser may mark the binary as quarantined since it has not been verified through the App Store. To remove the quarantine, run the following command:

```bash
xattr -d com.apple.quarantine /usr/local/bin/icm-relayer
```

### Download Docker Image

The published Docker image can be pulled from `avaplatform/icm-relayer:latest` on dockerhub.

### Build from Source

See the [Building](#building) section for instructions on how to build the relayer from source.

### System Requirements

- Ubuntu 22.04 or later
  - Tested on x86_64/amd64 architecture.
- MacOS 14.3 or later
  - Tested on arm64 architecture (Apple silicon).

### API Requirements

- ICM Relayer requires access to Avalanche API nodes for the P-Chain as well as any connected Subnets. The API nodes must have the following methods enabled:
  - Each Subnet API node must have enabled:
    - eth API (RPC and WS)
  - The P-Chain API node must have enabled:
    - platform.getHeight
    - platform.validatedBy
    - platform.getValidatorsAt OR platform.getCurrentValidators
  - The Info API node must have enabled:
    - info.peers
    - info.getNetworkID
  - If the Info API node is also a Subnet validator, it must have enabled:
    - info.getNodeID
    - info.getNodeIP

The Fuji and Mainnet [public API nodes](https://docs.avax.network/tooling/rpc-providers) provided by Avalanche have these methods enabled, and are suitable for use with the relayer.

### Peer-to-Peer Connections

- The ICM relayer implementation gathers BLS signatures from the validators of the source Subnet via peer-to-peer `AppRequest` messages. Validator nodes need to be configured to accept incoming peer connections. Otherwise, the relayer will fail to gather Warp message signatures. For example, networking rules may need to be adjusted to allow traffic on the default AvalancheGo P2P port (9651), or the public IP may need to be manually set in the [node configuration](https://docs.avax.network/nodes/configure/avalanchego-config-flags#public-ip).

### Private Key Management

- Each configured destination blockchain requires a private key to sign transactions. This key can be provided as a hex-encoded string in the configuration (see `account-private-key` in [Configuration](#configuration)) or environment variable, or stored in KMS and used to sign transactions remotely (see `kms-key-id` and `kms-aws-region` in [Configuration](#configuration)).
- **Each private key used by the relayer should not be used to sign transactions outside of the relayer**, as this may cause the relayer to fail to sign transactions due to nonce mismatches.

## Usage

### Options

The relayer binary accepts the following command line options. Other configuration options are not supported via the command line and must be provided via the configuration JSON file or environment variable.

```bash
icm-relayer --config-file path-to-config                Specifies the relayer config file and begin relaying messages.
icm-relayer --version                                   Display icm-relayer version and exit.
icm-relayer --help                                      Display icm-relayer usage and exit.
```

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Building

Before building, be sure to install Go, which is required even if you're just building the Docker image.

Build the relayer by running the script:

```bash
./scripts/build.sh
```

Build a Docker image by running the script:

```bash
./scripts/build_local_image.sh
```

### Configuration

The relayer is configured via a JSON file. Its path is passed in via the `--config-file` command line argument or the `CONFIG_FILE` environment variable, which takes precedence if both are set. See [`sample-relayer-config.json`](../sample-relayer-config.json) for a minimal example.

The relayer rejects unknown keys: it exits at startup if the file contains a key that is not listed in this section.

Top-level options that hold a single value can also be set via environment variable, which takes precedence over the value in the file. To get the environment variable corresponding to a key, upper case the key and change the delimiter from "-" to "_". For example, `LOG_LEVEL` sets the `"log-level"` JSON key. There are two limitations:

- Options nested inside an object or a list (for example, anything under `"p-chain-api"` or `"source-blockchains"`) can only be set in the file. The exception is the private keys of a destination blockchain, which have their own environment variables described under [Destination blockchains](#destination-blockchains).
- `REDIS_URL` and `ALLOW_PRIVATE_IPS` are only read if the `"redis-url"` or `"allow-private-ips"` key is also present in the file.

Blockchain IDs and Subnet IDs are cb58-encoded. Addresses are hex-encoded. An option is optional unless it is marked as required.

The configuration options are grouped as follows:

- [General](#general)
- [Avalanche API nodes](#avalanche-api-nodes)
- [Peer-to-peer network](#peer-to-peer-network)
- [Source blockchains](#source-blockchains)
- [Message protocols](#message-protocols)
- [Destination blockchains](#destination-blockchains)
- [External EVM destinations](#external-evm-destinations)

#### General

`"log-level": "verbo" | "debug" | "trace" | "info" | "warn" | "error" | "fatal" | "off"`

- The log level for the relayer. Defaults to `info`. Logs from the peer-to-peer network are only emitted at the `error` level, unless the log level is `debug` or `verbo`.

`"storage-location": string`

- The path to the directory in which the relayer will store its state as JSON files. Defaults to `./.icm-relayer-storage`. Not used if `redis-url` is set.

`"redis-url": string`

- The URL of the Redis server to use to manage state. This URL should specify the user, password, host, port, DB index, and protocol version. For example, `"redis://user:password@localhost:6379/0?protocol=3"`. Overrides `storage-location` if provided.

`"process-missed-blocks": boolean`

- Whether or not to process missed blocks after restarting. Defaults to `true`. If set to false, the relayer will start processing blocks from the chain head. See [Processing Missed Blocks](#processing-missed-blocks).

`"db-write-interval-seconds": unsigned integer`

- The interval in seconds at which the relayer will write the latest processed block heights to storage. Must be between `1` and `600`. Defaults to `10`.

`"api-port": unsigned integer`

- The port on which the relayer will listen for [API](#api) requests. Defaults to `8080`.

`"metrics-port": unsigned integer`

- The port on which the relayer will expose Prometheus metrics. Defaults to `9090`.

`"max-concurrent-messages": unsigned integer`

- The maximum number of messages the application will attempt to process concurrently. Processing messages involves making potentially multiple RPC requests, and issuing too many requests at once may cause failures. Must be greater than 0. Defaults to `250`.

`"signature-cache-size": unsigned integer`

- The maximum number of Warp messages for which the relayer caches the validator signatures it has collected. Once the cache is full, the least recently used message is evicted. Defaults to `1048576`.

#### Avalanche API nodes

`"p-chain-api": APIConfig`

- Required. The configuration for the Avalanche P-Chain API node. This API node needs to have the following methods enabled:
  - platform.getHeight
  - platform.validatedBy
  - platform.getValidatorsAt OR platform.getCurrentValidators

`"info-api": APIConfig`

- Required. The configuration for the Avalanche Info API node. This API node needs to have the following methods enabled:

  - info.peers
  - info.getNetworkID

- Additionally, if the Info API node is also a validator, it must have enabled:
  - info.getNodeID
  - info.getNodeIP

An `APIConfig` describes how to reach an API. It is used for the two API nodes above, and for the RPC and WebSocket endpoints of the source and destination blockchains. It has the following fields:

`"base-url": string`

- Required. The URL of the API. For `p-chain-api` and `info-api`, this is the URL of the node without a path, for example `https://api.avax.network`. For a blockchain's `rpc-endpoint` and `ws-endpoint`, this is the full URL of the endpoint, for example `https://api.avax.network/ext/bc/C/rpc`.

`"query-parameters": map[string]string`

- A map of query parameters to values that will be added to the URL of each request.

`"http-headers": map[string]string`

- A map of HTTP headers to include in each request.

#### Peer-to-peer network

`"manually-tracked-peers": []PeerConfig`

- A list of peers that the relayer will always connect to, in addition to the peers it discovers on its own. This is useful if the relayer cannot discover the validators of a source blockchain on its own, for example on a private network. Each `PeerConfig` has the following configuration:

  `"id": string`

  - The node ID of the peer, for example `NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg`.

  `"ip": string`

  - The IP address and P2P port of the peer, for example `203.0.113.10:9651`. Host names are not supported.

`"allow-private-ips": boolean`

- Whether or not the relayer is allowed to connect to peers that have private IP addresses. This is needed to connect to a local or private network. Defaults to `false`.

`"tls-cert-path": string`

- The path to a TLS cert file. Should only be set if a static NodeID is required for connecting to private networks. The relayer's NodeID is derived from this cert. If not set, the relayer generates a new cert, and therefore a new NodeID, each time it starts.

`"tls-key-path": string`

- The path to a TLS key file. Should only be set if a static NodeID is required for connecting to private networks.

- `tls-cert-path` and `tls-key-path` must be set together. If neither file exists yet, the relayer generates a cert and key and writes them to these paths. If only one of the files exists, the relayer exits. Environment variables in the paths are expanded, for example `$HOME/relayer.crt`.

`"initial-connection-timeout-seconds": unsigned integer`

- The maximum number of seconds to wait during start up to connect to sufficient stake weight of each supported chain. The relayer exits if it does not reach sufficient stake weight in that time. Must be greater than 0. Defaults to `300`.

`"validator-refresh-period-seconds": unsigned integer`

- The interval in seconds at which the relayer refreshes the validator sets of the Subnets it tracks. Must be greater than 0. Defaults to `60`.

#### Source blockchains

`"source-blockchains": []SourceBlockchain`

- Required. The list of source blockchains to support. At least one is required, each must have a unique `blockchain-id`, and together they can belong to at most 16 Subnets. Each `SourceBlockchain` has the following configuration:

  `"subnet-id": string`

  - Required. The ID of the Subnet that validates the source blockchain. The Subnet ID of the Primary Network is `11111111111111111111111111111111LpoYY`.

  `"blockchain-id": string`

  - Required. The ID of the source blockchain.

  `"rpc-endpoint": APIConfig`

  - Required. The RPC endpoint configuration of the source blockchain's API node. The API node is expected to have all standard ETH endpoints enabled.

  `"ws-endpoint": APIConfig`

  - Required. The WebSocket endpoint configuration of the source blockchain's API node. The API node is expected to accept eth_subscribe connections.

  `"message-contracts": map[string]MessageProtocolConfig`

  - Map of contract addresses to the config options of the protocol at that address. Which contract address to use depends on the protocol, see [Message protocols](#message-protocols). Each `MessageProtocolConfig` has the following configuration:

    `"message-format": "teleporter" | "teleporter-v2" | "off-chain-registry"`

    - Required. The message protocol of the contract.

    `"settings": object`

    - The settings of the message protocol, see [Message protocols](#message-protocols).

  `"supported-destinations": []SupportedDestination`

  - List of destinations that the source blockchain supports. If `supported-destinations` is empty, then all destination blockchains (and therefore all addresses on those destination blockchains) are supported. Each `SupportedDestination` has the following configuration:

    `"blockchain-id": string`

    - Required. The ID of the destination blockchain. It must be the `blockchain-id` of one of the `destination-blockchains`, or the `destination-blockchain-id` of one of the `external-evm-destinations`.

    `"addresses": []string`

    - List of addresses on that destination blockchain that the relayer supports delivering Warp messages to. The destination address is defined by the message protocol. For example, it could be the address called from the message protocol contract. If no supported addresses are provided, all addresses are allowed on that blockchain. The zero address is not allowed.

  `"process-historical-blocks-from-height": unsigned integer`

  - The block height at which to back-process transactions from the source blockchain. If the database already contains a later block height for the source blockchain, then that will be used instead. If zero or left unset, and the database does not contain a block height for the source blockchain, the relayer starts from the chain head. Will only be used if `process-missed-blocks` is set to `true`.

  `"allowed-origin-sender-addresses": []string`

  - List of addresses on this source blockchain to relay Warp messages from. The sending address is defined by the message protocol. For example, it could be defined as the EOA that initiates the transaction, or the address that calls the message protocol contract. If empty, then all addresses are allowed. The zero address is not allowed.

#### Message protocols

Each entry of a source blockchain's `message-contracts` selects a message protocol with `message-format`, and configures it with `settings`. The supported protocols are:

`"message-format": "teleporter"`

- Relays [Teleporter](../icm-contracts/avalanche/teleporter/README.md) messages. The contract address of the entry is the address of the `TeleporterMessenger` contract on the source blockchain. It has the following `settings`:

  `"reward-address": string`

  - Required. The address that will receive the rewards for the messages delivered by the relayer.

`"message-format": "teleporter-v2"`

- Relays TeleporterV2 messages. `TeleporterMessengerV2` verifies messages through an adapter contract, and the contract address of the entry is the address of that adapter on the source blockchain: the `WarpAdapter` for the `warp` verifier type, or the `MerkleValidatorSetRegistry` for the `merkle` verifier type. It has the following `settings`:

  `"reward-address": string`

  - Required. The address that will receive the rewards for the messages delivered by the relayer.

  `"teleporter-address": string`

  - Required. The address of the `TeleporterMessengerV2` contract on the destination blockchain.

  `"verifier-type": "merkle" | "warp"`

  - How messages are verified on the destination blockchain. It must match the adapter that the destination `TeleporterMessengerV2` was deployed with. Defaults to `merkle`.
    - `merkle`: messages are verified by a `MerkleValidatorSetRegistry`. This is the verifier type for [external EVM destinations](#external-evm-destinations).
    - `warp`: messages are verified by the `WarpAdapter`, through the Warp precompile.

  `"registry-address": string`

  - The address of the `MerkleValidatorSetRegistry` contract on the destination blockchain. Required for the `merkle` verifier type, and not used for the `warp` verifier type.

`"message-format": "off-chain-registry"`

- Relays off-chain Warp messages to a `TeleporterRegistry` contract, which are used to [register a new Teleporter version](../icm-contracts/avalanche/teleporter/registry/UPGRADING.md). These messages are not sent on the source blockchain, so the relayer does not listen for them. They are submitted through the [`/relay/message`](#relaymessage) API instead. The contract address of the entry is the zero address, `0x0000000000000000000000000000000000000000`. It has the following `settings`:

  `"teleporter-registry-address": string`

  - Required. The address of the `TeleporterRegistry` contract that the messages are delivered to.

#### Destination blockchains

`"destination-blockchains": []DestinationBlockchain`

- Required. The list of destination blockchains to support. At least one is required, and each must have a unique `blockchain-id`. Each `DestinationBlockchain` has the following configuration:

  `"subnet-id": string`

  - Required. The ID of the Subnet that validates the destination blockchain.

  `"blockchain-id": string`

  - Required. The ID of the destination blockchain.

  `"rpc-endpoint": APIConfig`

  - Required. The RPC endpoint configuration of the destination blockchain's API node.

  Each destination blockchain requires at least one key to sign transactions with. Keys can be provided through any combination of the private key and KMS options below, and of the environment variables described after them. If more than one key is provided, the relayer spreads its transactions across all of the accounts, which increases the number of transactions it can have pending at once.

  `"account-private-key": string`

  - The hex-encoded private key to use for signing transactions on the destination blockchain.
  - Please note that the private key should be exclusive to the relayer, see [Private Key Management](#private-key-management).

  `"account-private-keys-list": []string`

  - A list of hex-encoded private keys for signing transactions on the destination blockchain.
  - Please note that all private keys should be exclusive to the relayer, see [Private Key Management](#private-key-management).

  `"kms-key-id": string`

  - The ID of the KMS key to use for signing transactions on the destination blockchain. If `kms-key-id` is provided, then `kms-aws-region` is required.
  - Please note that the private key in KMS should be exclusive to the relayer, see [Private Key Management](#private-key-management).

  `"kms-aws-region": string`

  - The AWS region in which the KMS key is located. Required if `kms-key-id` is provided.

  `"kms-keys": []KMSKey`

  - A list of KMS Keys that may be used for signing transactions on the destination blockchain. A `KMSKey` has the following fields:

    `"key-id": string`

    - Required. The ID of the KMS key to use for signing transactions on the destination blockchain.

    `"aws-region": string`

    - Required. The AWS region in which the KMS key is located.

  Private keys may also be provided by the following environment variables. They are used in addition to any keys in the file.

  - `ACCOUNT_PRIVATE_KEY`: a hex-encoded private key that is used for every destination blockchain.
  - `ACCOUNT_PRIVATE_KEYS_LIST`: a space-delimited list of hex-encoded private keys that are used for every destination blockchain.
  - `ACCOUNT_PRIVATE_KEY_<blockchain-id>` and `ACCOUNT_PRIVATE_KEYS_LIST_<blockchain-id>`: the same, but only used for the destination blockchain with that `blockchain-id`. For example, `ACCOUNT_PRIVATE_KEY_2D8RG4UpSXbPbvPCAWppNJyqTG2i2CAXSkTgmTBBvs7GKNZjsY`.

  `"block-gas-limit": unsigned integer`

  - The maximum amount of gas that can be used in a single block on this blockchain. The relayer will not attempt to deliver messages that require more gas than this limit to the given chain. Defaults to `12000000`. Cannot be set higher than `12000000` for a blockchain on the Primary Network.

  `"max-base-fee": unsigned integer`

  - The maximum base fee gas price (in WEI) the relayer is willing to pay on this blockchain. If zero or left unset, the relayer will use three times the current base fee estimation, and not have an explicit maximum.

  `"suggested-priority-fee-buffer": unsigned integer`

  - The fee per gas (in WEI) that will be added to the priority fee rate on top of the suggested priority fee fetched via RPC. A higher suggested priority fee buffer will make transaction inclusion more likely in cases where a chain's capacity is being exhausted, but also results in higher effective gas prices for transactions when potentially not necessary. The maximum priority fee per gas is also capped by `max-priority-fee-per-gas`, as described below. Defaults to `0`.

  `"max-priority-fee-per-gas": unsigned integer`

  - The maximum priority fee per gas (in WEI) that the relayer is willing to pay to incentivize transactions being included on this blockchain. The relayer will use the current estimation of the required gas tip cap for this blockchain plus the `suggested-priority-fee-buffer`, up to a maximum of this configured value. Defaults to `2500000000` (2.5 GWEI).

  `"tx-inclusion-timeout-seconds": unsigned integer`

  - The time in seconds to wait for a sent transaction to be included in a block when verifying transaction receipts before erroring out. Defaults to `30`.

#### External EVM destinations

`"external-evm-destinations": []ExternalEVMDestination`

- The list of EVM chains outside of Avalanche, such as Ethereum, to support. For each entry, the relayer keeps the validator set of one Avalanche blockchain up to date in a `MerkleValidatorSetRegistry` contract on the external chain. If `destination-blockchain-id` is set, the relayer also delivers TeleporterV2 messages from that Avalanche blockchain to the external chain. See [Relayer changes](../docs/external-interop/origin_avalanche/relayer_changes.md) for how this works. Each `ExternalEVMDestination` has the following configuration:

  `"rpc-endpoint": string`

  - Required. The URL of the RPC endpoint of the external chain's API node. Unlike the other endpoints, this is a string and not an `APIConfig`.

  `"contract-address": string`

  - Required. The address of the `MerkleValidatorSetRegistry` contract on the external chain.

  `"blockchain-id": string`

  - Required. The ID of the Avalanche blockchain whose validator set the relayer registers on the external chain. If message delivery is enabled, this is also the source blockchain of the delivered messages.

  `"subnet-id": string`

  - Required. The ID of the Subnet that validates that blockchain.

  `"private-key": string`

  - Required. The hex-encoded private key to use for signing the transactions that update the validator set on the external chain. It must not have a "0x" prefix.
  - Please note that the private key should be exclusive to the relayer, see [Private Key Management](#private-key-management).

  `"poll-interval-seconds": unsigned integer`

  - The interval in seconds at which the relayer checks whether the validator set has changed. Defaults to `10`.

  `"max-update-interval-seconds": unsigned integer`

  - The maximum time in seconds between two updates of the validator set on the external chain. Once it has elapsed, the relayer sends an update even if the validator set has not changed. If zero or left unset, the relayer only sends an update when the validator set changes.

  `"max-gas-price-gwei": unsigned integer`

  - The maximum gas price (in GWEI) at which the relayer will send a validator set update. If the gas price suggested by the external chain is higher, the update is deferred to the next poll. If zero or left unset, there is no maximum.

  The remaining options configure message delivery, and are only used if `destination-blockchain-id` is set. To deliver messages, the source blockchain also needs a [`teleporter-v2`](#message-protocols) entry in its `message-contracts` that uses the `merkle` verifier type.

  `"destination-blockchain-id": string`

  - The blockchain ID that TeleporterV2 messages use to address the external chain. Setting it enables message delivery to the external chain. If left unset, the relayer only updates the validator set.

  `"delivery-private-key": string`

  - The hex-encoded private key to use for signing the transactions that deliver messages. Required if `destination-blockchain-id` is set. It must be different from `private-key`, because the validator set updates and the message deliveries would otherwise conflict on the nonce of a shared account.
  - Please note that the private key should be exclusive to the relayer, see [Private Key Management](#private-key-management).

  `"teleporter-address": string`

  - The address of the `TeleporterMessengerV2` contract on the external chain. Required if `destination-blockchain-id` is set. Messages are delivered to the `teleporter-address` in the `settings` of the source blockchain's `teleporter-v2` message contract, so the two should be the same.

  `"quorum-numerator": unsigned integer`

  - The percentage of stake weight that must sign a message delivered to the external chain. Defaults to `67`, which is the quorum that the `MerkleValidatorSetRegistry` contract requires. The contract rejects messages signed by a smaller quorum.

  `"block-gas-limit": unsigned integer`

  - The maximum amount of gas that a transaction delivering a message can use. The relayer will not attempt to deliver messages that require more gas than this limit. Defaults to `12000000`.

  `"tx-inclusion-timeout-seconds": unsigned integer`

  - The time in seconds to wait for a transaction delivering a message to be included in a block. Defaults to `30`.

## Architecture

### Components

The relayer consists of the following components:

- At the global level:
  - P2P app network: issues signature `AppRequests`
  - P-Chain client: gets the validators for a Subnet
  - Relayer database: stores latest processed block for each Application Relayer
    - Currently supports Redis and local JSON file storage
- Per Source Blockchain
  - Subscriber: listens for logs pertaining to cross-chain message transactions
  - Source RPC client: queries for missed blocks on startup
- Per Destination Blockchain
  - Destination RPC client: broadcasts transactions to the destination
- Application Relayers
  - Relay messages from a specific source blockchain and source address to a specific destination blockchain and destination address

### Data Flow

<figure>
  <img src="../resources/relayer-diagram.png?raw=true"></img>
  <figcaption>Figure 1: Relayer Data Flow</figcaption>
</figure>

### Processing Missed Blocks

On startup, the relayer will process any blocks that it missed while offline if `process-missed-blocks` is set to `true` in the configuration. For each configured `source-blockchain`, the starting block height is set as the _minimum_ block height that is stored in the relayer's database across keys that pertain to that blockchain. These keys correspond to distinct sending addresses (as specified in `source-blockchain.allowed-origin-sender-addresses`), meaning that on startup, the relayer will begin processing from the _minimum_ block height across all configured sending addresses for each `source-blockchain`. Note that an empty `source-blockchain.allowed-origin-sender-addresses` list is treated as its own distinct key. If no keys are found, then the relayer begins processing from the current chain tip.

Once the starting block height is calculated, all blocks between it and the current tip of the `source-blockchain` are processed according to the _current_ configuration rules. 

_Note:_ Given these semantics for computing the starting block height, it's possible for blocks that have previously been ignored under different configuration options to be relayed on a subsequent run. For example, consider the following scenario consisting of three subsequent runs (see Figure 2 below):

- **Run 1**: Suppose that on Blockchain A we set `allowed-origin-sender-addresses=[0x1234]`, meaning that the relayer should _ignore_ messages sent by any other address. The relayer processes through block `100`, and that height is marked as processed for `0x1234`'s key.

- **Run 2**: The relayer is then restarted with `allowed-origin-sender-addresses=[0xabcd]`, replacing `0x1234`. A message is sent from address `0x1234` at block height `200`. The relayer will decide to ignore this message, and will mark block `200` as processed for `0xabcd`'s key.

- **Run 3**: The relayer is then restarted again with the original configuration of `allowed-origin-sender-addresses=[0x1234]`. The relayer will calculate the starting block height as `100`, and process blocks `100` through the current chain tip, reprocessing block `200` along the way. Instead of ignoring the message in this block, however, the relayer will relay it to the destination.

<figure>
  <img src="../resources/catch-up-example.png?raw=true"></img>
  <figcaption>Figure 2: Processing Missed Blocks Example</figcaption>
</figure>

### API

Request bodies are limited to 48 KiB, which fits the largest Warp message that can be relayed. A larger request is rejected with `HTTP 413`.

#### `/relay`
- Used to manually relay a Warp message. The body of the request must contain the following JSON:

```json
{
 "blockchain-id": "<cb58-encoded or '0x' prefixed hex-encoded of blockchain ID>",
 "message-id": "<cb58-encoded or '0x' prefixed hex-encoded of Warp message ID>",
 "block-num": "<Block number that the message was sent in>"
}
```

- If successful, the endpoint will return the following JSON:

```json
{
 "transaction-hash": "<Transaction hash that includes the delivered warp message>"
}
```

#### `/relay/message`

- Used to manually relay a warp message. The body of the request must contain the following JSON:

```json
{
 "unsigned-message-bytes": "<Hex encoded byte array containing the unsigned warp message>",
 "source-address": "<Hex encoding of address that sent the warp message>"
}
```

- If successful, the endpoint will return the following JSON:

```json
{
 "transaction-hash": "<Transaction hash that includes the delivered Warp message>",
}
```

#### `/health`

- Takes no arguments. Returns a `200` status code if all Application Relayers are healthy. Returns a `503` status if any of the Application Relayers have experienced an unrecoverable error. Here is an example return body:

```json
{
  "status": "down",
  "details": {
    "relayers-all": {
      "status": "down",
      "timestamp": "2024-06-01T05:06:07.685522Z",
      "error": "<List of cb-58 encoded IDs for unhealthy relayers>"
    }
  }
}
```

## Testing

### Unit Tests

Unit tests can be ran locally by running the command in the root of the project:

```bash
./scripts/test.sh
```

If your temporary directory is not writable, the unit tests may fail with messages like `fork/exec /tmp/go-build2296620589/b247/config.test: permission denied`. To fix this, set the `TMPDIR` environment variable to something writable, for example `export TMPDIR=~/tmp`.

### E2E Tests

See [Build and Test Requirements](#build-and-test-requirements) for the tooling the E2E tests depend on. Run the tests using the dedicated script:

```bash
./scripts/e2e_test.sh
```

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for [test descriptions](./tests/e2e_test.go#L68) that match the provided input. For example, to run the `Basic Relay` test:

```bash
GINKGO_FOCUS="Basic" ./scripts/e2e_test.sh
```

The E2E tests use the `TeleporterMessenger` contract deployment transaction specified in the following files:

- `tests/utils/UniversalTeleporterDeployerAddress.txt`
- `tests/utils/UniversalTeleporterDeployerTransaction.txt`
- `tests/utils/UniversalTeleporterMessagerContractAddress.txt`
  To update the version of Teleporter used by the E2E tests, update these values with the latest contract deployment information. For more information on how to deploy the Teleporter contract, see the [Teleporter documentation](https://github.com/ava-labs/icm-contracts/tree/main/utils/contract-deployment).

### Generate Mocks

[Gomock](https://pkg.go.dev/go.uber.org/mock/gomock) is used to generate mocks for testing. To generate mocks, run the following command at the root of the project:

```bash
go generate ./...
```

### Generate Abi Bindings

[subnet-evm](https://github.com/ava-labs/subnet-evm/tree/6e67777132dd7d0d556e1e34d68ad8e27b22ebef/cmd/abigen) is used to generate abi binding `.go` files for solidity contracts. If you change any of the smart contracts, you will have to update the abi bindings. To generate these files, run the following command at the root of the project:

```bash
./scripts/abi_bindings.sh
```
