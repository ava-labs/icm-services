# Relayer changes

Existing relayers need to be extended to handle external EVM chains. Not only are new clients necessary and the way messages are sent changed, but new types of messages now need to be relayed. For the [`MerkleValidatorSetRegistry`](validator_set_registry.md) contracts, new validator set commitments need to be relayed periodically.

## Updating validator sets

The relayer runs a `MerkleSetUpdater` component that monitors validator set changes for Avalanche L1s and posts updates to external EVM chains. All L1 validators are registered on the P-chain, which maintains separate validator sets for each L1. The updater queries the P-chain for each configured L1's validator set and posts commitment updates to `MerkleValidatorSetRegistry` contracts on external chains.

Before signing a delivered message, the relayer queries the destination registry to determine which P-chain height the source L1's stored commitment was built from, then gathers signatures from the validator set at exactly that height. This is what lets the external chain verify the signatures against the root it has stored.

One `MerkleSetUpdater` is started per configured external EVM destination. Each polls the P-chain at a configurable interval, and on each iteration:

1. Fetches the current P-chain height and the L1's validator set at that height.
2. Sorts the validators into canonical order (ascending by uncompressed BLS public key) and computes the Merkle root over them.
3. Posts an update when either the root differs from the last one it posted, or the last update is older than `max-update-interval-seconds`. Because the commitment is a fixed-size struct, the update is always a single transaction.
4. Skips the update when the destination chain's suggested gas price exceeds `max-gas-price-gwei`, retrying on the next poll.

To post an update the relayer builds a `ValidatorSetMerkleCommitment` payload, has it signed by the appropriate Avalanche validator set, and submits it to the registry's `registerValidatorSet` function. The same function handles both the first registration for a blockchain ID and all subsequent updates; see [Validator set registry](validator_set_registry.md) for which validator set must sign in each case.

### Configuration

External EVM chains are configured with `external-evm-destinations` blocks in the relayer configuration:

```json
{
  "source-blockchains": [{
    "subnet-id": "11111111111111111111111111111111LpoYY",
    "blockchain-id": "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
    "rpc-endpoint": {"base-url": "https://api.avax.network/ext/bc/C/rpc"}
  }],
  "external-evm-destinations": [{
    "rpc-endpoint": "https://eth-mainnet.g.alchemy.com/v2/...",
    "contract-address": "0xABCD...",
    "blockchain-id": "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
    "subnet-id": "11111111111111111111111111111111LpoYY",
    "private-key": "...",
    "poll-interval-seconds": 10,
    "max-update-interval-seconds": 3600,
    "max-gas-price-gwei": 50
  }]
}
```

`blockchain-id` and `subnet-id` identify the Avalanche L1 whose validator set is being tracked; `contract-address` is the registry deployed on the external chain. A single entry can additionally act as a TeleporterV2 message-delivery destination by setting `destination-blockchain-id`, `delivery-private-key`, and `teleporter-address`. The delivery key must differ from `private-key` so the updater and the delivery client never share a sender account and collide on nonces.

### External EVM destination client

`ExternalEVMDestinationClient` implements the `vms.DestinationClient` interface for external chains. Its `GetPChainHeightForDestination` reads the height directly out of the registry's stored commitment for the source blockchain, so signatures are gathered at exactly the height the stored root was built from:

```go
func (c *ExternalEVMDestinationClient) GetPChainHeightForDestination(
    ctx context.Context,
) (uint64, error) {
    registry, err := merkleregistry.NewMerkleValidatorSetRegistry(c.registryAddress, c.ethClient)
    if err != nil {
        return 0, fmt.Errorf("failed to bind merkle registry: %w", err)
    }
    commitment, err := registry.GetValidatorSetCommitment(
        &bind.CallOpts{Context: ctx}, c.sourceBlockchainID,
    )
    if err != nil {
        return 0, fmt.Errorf("failed to read merkle commitment: %w", err)
    }
    return commitment.PChainHeight, nil
}
```

The destination client factory in [`vms/destination_client.go`](../../../vms/destination_client.go) creates standard Avalanche L1 clients from `destination-blockchains` and external EVM clients from the `external-evm-destinations` entries that have `destination-blockchain-id` set.
