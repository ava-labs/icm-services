package ethereum_icm_verification

import (
	"context"

	"github.com/ava-labs/avalanchego/vms/platformvm"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
)

// AvalancheValidatorSetRegistry Test that we can deploy a DiffUpdater contract on Ethereum and
// populate it with the validator set from the Avalanche network.
// 1. Deploy a DiffUpdater contract on Ethereum
// 2. Apply the shards to initialize the initialize validator set
func AvalancheValidatorSetRegistry(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	localEthereumNetwork *localnetwork.LocalEthereumNetwork,
) {
	_, ethFundedKey := localEthereumNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	serializedShards := utils.DeployDiffUpdater(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		ethFundedKey,
		1,
		primaryNetworkInfo.BlockchainID,
		primaryNetworkInfo.SubnetID,
		platformvm.NewClient(primaryNetworkInfo.NodeURIs[0]),
	)

	// apply the shards to initialize the validator set
	for _, shard := range serializedShards {

	}
}
