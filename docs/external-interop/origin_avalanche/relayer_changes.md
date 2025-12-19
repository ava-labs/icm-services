# Relayer changes

Existing relayers need to be extended to handle external EVM chains. Not only will new clients be necessary and  the way messages are sent change, but new types of messages now need to be relayed. For example, for the `AvalancheValidatorSetRegistry` contracts, new validator sets will need to be relayed periodically.

## Updating Validator sets

Avalanche will require a `ValidatorSetUpdater` component that monitors validator set changes for Avalanche L1s and posts  updates to external EVM chains. All L1 validators are registered on the P-chain, which maintains separate validator sets  for each L1. The updater queries the P-chain for each configured L1's validator set and posts updates to  `AvalancheValidatorSetRegistry` contracts on external chains.  Before signing messages, the relayer queries the destination registry to determine what P-chain height was last registered for the source L1, then signs with the specific validators.  External EVM chains need to verify ICM messages from Avalanche L1s by checking validator signatures.

The `ValidatorSetUpdater` component polls P-chain at configurable intervals. On each iteration it performs the following:

1. Query current P-chain height using `GetLatestHeight` 
2. If height changed, fetch ALL validator sets: `GetAllValidatorSets(height)` returns `map[subnetID]ValidatorSet`
3. Update the `AvalancheValidatorSetRegistry` for each L1 and P-chain   
   1. This is done when an L1's validator set reaches some configurable expiration time (base on the P-chain height/timestamp of registration) or exceeds a churn threshold.
   2. If there is a newer validator set available, post update to the external EVM chains   
   3. Call for each L1:       
       ```go 
       updateValidatorSet(
           uint64 pchainHeight, 
           bytes32 sourceSubnetID, 
           ValidatorInfo[] memory validators
       )
       ```  
      The `ValidatorSetUpdater` struct should look like the following:

```go
type ValidatorSetUpdater struct {
    validatorManager      *peers.ValidatorManager
    externalEVMClients    map[string]*vms.DestinationClient // chainID -> destination client interface
    sourceSubnetIDs       []ids.ID
    pollingInterval       time.Duration
    database              database.Database

    // Per-subnet tracking
    lastPostedHeights     map[ids.ID]uint64  // subnetID -> P-chain height
    lastPostedHashes      map[ids.ID][]byte  // subnetID -> validator set hash
}

// Main loop
func (u *ValidatorSetUpdater) Run(ctx context.Context) error {
    ticker := time.NewTicker(time.Duration(u.config.PollIntervalSeconds) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            if err := u.checkAndUpdate(ctx); err != nil {
                u.logger.Error("Failed to check/update", zap.Error(err))
            }
        case <-ctx.Done():
            return nil
        }
    }
}

func (u *ValidatorSetUpdater) checkAndUpdate(ctx context.Context) error {
    // 1. Get current P-chain height
    currentHeight := u.validatorManager.GetLatestSyncedPChainHeight()

    // 2. Fetch ALL validator sets (one call gets all L1s)
    allValidatorSets, err := u.validatorManager.GetAllValidatorSets(ctx, currentHeight)
    if err != nil {
        return err
    }
    
    // When to post the change: 
    // - percent changes of the validator?
    //   > x% percentage change we post
    // 
    
    // Check if there is validator set updated 
    // If yes, call the externalEVMClient to call the updateValidatorSets.
    
    return nil
```

### Configuration

To configure the relayer, we will add `destination-external-evm-blockchains` blocks to the relayer configuration that look as follows: 
```json
{
  "source-blockchains": [{    
          "subnet-id": "11111111111111111111111111111111LpoYY",
          "blockchain-id": "yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp",
          "rpc-endpoint": {"base-url": "https://api.avax.network/ext/bc/C/rpc"},
          ...
  }], 
  "destination-external-evm-blockchains": [{
    "chain-id": "1",
    "registry-contract-address": "0xABCD...",
    "rpc-endpoint": {"base-url": "https://eth-mainnet.g.alchemy.com/v2/..."},	    
  }]
}
```

### External Evm Destination Client

A new client, `ExternalEVMDestinationCient` will be created implementing the `vms.DestinationClient` interface. The implementation of `GetPChainHeightForDestination` function will look as follows:
```go
type externalEVMDestinationClient struct {
    client              ethclient.Client
    registryAddress     common.Address
    registryContract    *validatorregistry.AvalancheValidatorSetRegistry // go binding of registry contract
}

func (c *externalEVMDestinationClient) GetPChainHeightForDestination(
    ctx context.Context,
) (uint64, error) {
    // Query the registry contract on external EVM chain
    height, err := c.registryContract.GetCurrentPChainHeight(
        &bind.CallOpts{Context: ctx},
    )
    if err != nil {
        return 0, fmt.Errorf("failed to query registry contract: %w", err)
    }
    
    return height, nil
}
```

The destination client factory will create separate client types for Avalanche L1s and external EVM chains: 
```go
// vms/destination_client.go
// https://github.com/ava-labs/icm-services/blob/ba3c9944b0eba2b24fd5c455325de54189a32bd3/vms/destination_client.go#L64

func CreateDestinationClients(
    logger logging.Logger,
    relayerConfig *config.Config,
) (map[ids.ID]DestinationClient, error) {
    destinationClients := make(map[ids.ID]DestinationClient)
    
    // Create clients for Avalanche L1 destinations
    for _, destConfig := range relayerConfig.DestinationBlockchains {
        blockchainID := ids.FromString(destConfig.BlockchainID)
        
        // Standard Avalanche L1 client with ProposerVM support
        client, err := evm.NewDestinationClient(logger, destConfig, epochDuration)
        if err != nil {
            return nil, err
        }
        destinationClients[blockchainID] = client
    }
    
    // Create clients for external EVM destinations
    for _, extConfig := range relayerConfig.DestinationExternalEVMBlockchains {
        // Use chain ID as the identifier for external chains
        chainIDBytes := [32]byte{}
        binary.BigEndian.PutUint64(chainIDBytes[24:], extConfig.chainID)
        blockchainID := ids.ID(chainIDBytes)
        
        // External EVM client with registry contract support
        client, err := evm.NewExternalEVMDestinationClient(
            extConfig,
        )
        if err != nil {
            return nil, err
        }
        destinationClients[blockchainID] = client
    }
    
    return destinationClients, nil
}
```