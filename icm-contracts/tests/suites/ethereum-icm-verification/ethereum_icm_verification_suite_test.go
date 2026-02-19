package ethereum_icm_verification

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	ecdsaVerifierByteCodeFile = "./out/ECDSAVerifier.sol/ECDSAVerifier.json"
	warpGenesisTemplateFile   = "./tests/utils/warp-genesis-template.json"
)

var (
	localAvalancheNetworkInstance *localnetwork.LocalAvalancheNetwork
	localEthereumNetworkInstance  *localnetwork.LocalEthereumNetwork
	teleporterInfo                utils.TeleporterTestInfo
	e2eFlags                      *e2e.FlagVars
)

var _ = ginkgo.BeforeSuite(func() {
	// Create the local network instances
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	teleporterContractAddress,
		teleporterDeployerAddress,
		teleporterDeployedByteCode := utils.TeleporterDeploymentValues()

	localAvalancheNetworkInstance = localnetwork.NewLocalAvalancheNetwork(
		ctx,
		"ethereum-icm-verification-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:                       "L1",
				EVMChainID:                 12345,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				TeleporterDeployedBytecode: teleporterDeployedByteCode,
				NodeCount:                  1,
			},
		},
		1,
		1,
		e2eFlags,
	)
	log.Info("Started local Avalanche network", zap.Any("networkID", localAvalancheNetworkInstance.NetworkID))
	_, fundedKey := localAvalancheNetworkInstance.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetworkInstance.GetPrimaryNetworkInfo()

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	teleporterInfo = utils.NewTeleporterTestInfo(localAvalancheNetworkInstance.GetAllL1Infos())
	teleporterDeployerTransaction := utils.TeleporterDeployerTransaction()
	utils.DeployWithNicksMethod(
		ctx,
		&primaryNetworkInfo,
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	for _, l1 := range localAvalancheNetworkInstance.GetAllL1Infos() {
		teleporterInfo.SetTeleporter(teleporterContractAddress, l1)
		teleporterInfo.DeployTeleporterRegistry(l1, fundedKey)
	}

	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(ecdsaVerifierByteCodeFile)
	Expect(err).Should(BeNil())

	// Generate the ECDSAVerifier deployer transaction via Nick's method
	ecdsaVerifierDeployerTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		false,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
	)
	Expect(err).Should(BeNil())
	// Deploy the ECDSAVerifier contract on the C-Chain
	utils.DeployWithNicksMethod(
		ctx,
		&primaryNetworkInfo,
		ecdsaVerifierDeployerTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedKey,
	)

	localEthereumNetworkInstance = localnetwork.NewLocalEthereumNetwork(ctx)
	log.Info("Started local Ethereum network", zap.Any("chainID", localEthereumNetworkInstance.ChainID))
	// Deploy the ECDSAVerifier contract on the Ethereum chain
	_, fundedEthereumKey := localEthereumNetworkInstance.GetFundedAccountInfo()
	utils.DeployWithNicksMethod(
		ctx,
		localEthereumNetworkInstance.EthereumTestInfo(),
		ecdsaVerifierDeployerTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedEthereumKey,
	)

	log.Info("Set up ginkgo before suite")
})
