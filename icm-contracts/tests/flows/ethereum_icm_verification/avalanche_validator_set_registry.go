package ethereum_icm_verification

import (
	"context"

	"github.com/ava-labs/avalanchego/vms/platformvm"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	. "github.com/onsi/gomega"
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
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	contractAddress, serializedShards := utils.DeployDiffUpdater(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		fundedEthereumKey,
		1,
		primaryNetworkInfo.BlockchainID,
		primaryNetworkInfo.SubnetID,
		platformvm.NewClient(primaryNetworkInfo.NodeURIs[0]),
	)

	avalancheValidatorSetRegistry, err := diffupdater.NewDiffUpdater(
		contractAddress,
		localEthereumNetwork.EthClient,
	)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(fundedEthereumKey, localEthereumNetwork.ChainID)
	Expect(err).Should(BeNil())

	// apply the shards to initialize the validator set
	for i, shardBytes := range serializedShards {
		shard := diffupdater.ValidatorSetShard{
			AvalancheBlockchainID: primaryNetworkInfo.BlockchainID,
			ShardNumber:           uint64(i) + 1,
		}
		tx, err := avalancheValidatorSetRegistry.UpdateValidatorSet(opts, shard, shardBytes)
		Expect(err).Should(BeNil())
		receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())
		if i == len(serializedShards) {
			event, err := utils.GetEventFromLogs(receipt.Logs, avalancheValidatorSetRegistry.ParseValidatorSetUpdated)
			Expect(err).Should(BeNil())
			Expect(event.AvalancheBlockchainID).Should(Equal(primaryNetworkInfo.BlockchainID))
		}
	}
	registered, err := avalancheValidatorSetRegistry.IsRegistered(&bind.CallOpts{}, primaryNetworkInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(registered).Should(BeTrue())
}
