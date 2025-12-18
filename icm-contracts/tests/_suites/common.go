package suites

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"
	teleporterByteCodeFile  = "./out/TeleporterMessenger.sol/TeleporterMessenger.json"
)

func StartDefaultNetwork(ctx context.Context, networkName string, e2eFlags *e2e.FlagVars) *network.LocalNetwork {
	teleporterContractAddress,
		teleporterDeployerAddress,
		teleporterDeployedByteCode := utils.TeleporterDeploymentValues()

	ctx, cancel := context.WithTimeout(ctx, 240*2*time.Second)
	defer cancel()

	localNetworkInstance := network.NewLocalNetwork(
		ctx,
		networkName,
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:                         "A",
				EVMChainID:                   12345,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
			{
				Name:                         "B",
				EVMChainID:                   54321,
				TeleporterContractAddress:    teleporterContractAddress,
				TeleporterDeployedBytecode:   teleporterDeployedByteCode,
				TeleporterDeployerAddress:    teleporterDeployerAddress,
				NodeCount:                    5,
				RequirePrimaryNetworkSigners: true,
			},
		},
		4,
		4,
		e2eFlags,
	)
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	_, fundedKey := localNetworkInstance.GetFundedAccountInfo()
	utils.DeployTeleporterMessenger(
		ctx,
		localNetworkInstance.GetPrimaryNetworkInfo(),
		utils.TeleporterDeployerTransaction(),
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	return localNetworkInstance
}