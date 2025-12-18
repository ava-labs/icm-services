package suites

import (
	"context"
	"flag"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	icttFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ictt"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/common"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/segmentio/encoding/json"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"

	icttLabel              = "ICTT"
	erc20TokenHomeLabel    = "ERC20TokenHome"
	erc20TokenRemoteLabel  = "ERC20TokenRemote"
	nativeTokenHomeLabel   = "NativeTokenHome"
	nativeTokenRemoteLabel = "NativeTokenRemote"
	multiHopLabel          = "MultiHop"
	sendAndCallLabel       = "SendAndCall"
	registrationLabel      = "Registration"
	upgradabilityLabel     = "upgradability"

	teleporterRegistryAddressFile = "TeleporterRegistryAddress.json"
)

var (
	LocalNetworkInstance *localnetwork.LocalNetwork
	TeleporterInfo       utils.TeleporterTestInfo
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestICTT(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ICTT e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	teleporterContractAddress,
		teleporterDeployerAddress,
		teleporterDeployedByteCode := utils.TeleporterDeploymentValues()

	teleporterDeployerTransaction := utils.TeleporterDeployerTransaction()

	// Create the local network instance
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	LocalNetworkInstance = localnetwork.NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:                       "A",
				EVMChainID:                 12345,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployedBytecode: teleporterDeployedByteCode,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				NodeCount:                  2,
			},
			{
				Name:                       "B",
				EVMChainID:                 54321,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployedBytecode: teleporterDeployedByteCode,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				NodeCount:                  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	TeleporterInfo = utils.NewTeleporterTestInfo(LocalNetworkInstance.GetAllL1Infos())
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()

	if e2eFlags.NetworkDir() == "" {
		// Only deploy Teleporter if we are not reusing an existing network
		utils.DeployTeleporterMessenger(
			ctx,
			LocalNetworkInstance.GetPrimaryNetworkInfo(),
			teleporterDeployerTransaction,
			teleporterDeployerAddress,
			teleporterContractAddress,
			fundedKey,
		)

		for _, l1 := range LocalNetworkInstance.GetAllL1Infos() {
			TeleporterInfo.SetTeleporter(teleporterContractAddress, l1)
			TeleporterInfo.InitializeBlockchainID(l1, fundedKey)
			TeleporterInfo.DeployTeleporterRegistry(l1, fundedKey)
		}

		registryAddresseses := make(map[string]string)
		for l1, teleporterInfo := range TeleporterInfo {
			registryAddresseses[l1.Hex()] = teleporterInfo.TeleporterRegistryAddress.Hex()
		}

		jsonData, err := json.Marshal(registryAddresseses)
		Expect(err).Should(BeNil())
		err = os.WriteFile(teleporterRegistryAddressFile, jsonData, fs.ModePerm)
		Expect(err).Should(BeNil())

	} else {
		// Read the Teleporter registry address from the file
		registryAddresseses := make(map[string]string)
		data, err := os.ReadFile(teleporterRegistryAddressFile)
		Expect(err).Should(BeNil())
		err = json.Unmarshal(data, &registryAddresseses)
		Expect(err).Should(BeNil())

		for _, l1 := range LocalNetworkInstance.GetAllL1Infos() {
			TeleporterInfo.SetTeleporter(teleporterContractAddress, l1)
			TeleporterInfo.SetTeleporterRegistry(
				common.HexToAddress(registryAddresseses[l1.BlockchainID.Hex()]),
				l1,
			)
		}
	}

})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
	LocalNetworkInstance = nil
})

var _ = ginkgo.Describe("[ICTT integration tests]", func() {
	// ICTT tests
	ginkgo.It("Transfer an ERC20 token between two L1s",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, erc20TokenRemoteLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeERC20TokenRemote(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer a native token to an ERC20 token",
		ginkgo.Label(icttLabel, nativeTokenHomeLabel, erc20TokenRemoteLabel),
		func(ctx context.Context) {
			icttFlows.NativeTokenHomeERC20TokenRemote(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer a native token to a native token",
		ginkgo.Label(icttLabel, nativeTokenHomeLabel, nativeTokenRemoteLabel),
		func(ctx context.Context) {
			icttFlows.NativeTokenHomeNativeDestination(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer an ERC20 token with ERC20TokenHome multi-hop",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, erc20TokenRemoteLabel, multiHopLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeERC20TokenRemoteMultiHop(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer a native token with NativeTokenHome multi-hop",
		ginkgo.Label(icttLabel, nativeTokenHomeLabel, erc20TokenRemoteLabel, multiHopLabel),
		func(ctx context.Context) {
			icttFlows.NativeTokenHomeERC20TokenRemoteMultiHop(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer an ERC20 token to a native token",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, nativeTokenRemoteLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeNativeTokenRemote(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer a native token with ERC20TokenHome multi-hop",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, nativeTokenRemoteLabel, multiHopLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeNativeTokenRemoteMultiHop(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer a native token to a native token multi-hop",
		ginkgo.Label(icttLabel, nativeTokenHomeLabel, nativeTokenRemoteLabel, multiHopLabel),
		func(ctx context.Context) {
			icttFlows.NativeTokenHomeNativeTokenRemoteMultiHop(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transfer an ERC20 token using sendAndCall",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, erc20TokenRemoteLabel, sendAndCallLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeERC20TokenRemoteSendAndCall(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Registration and collateral checks",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, nativeTokenRemoteLabel, registrationLabel),
		func(ctx context.Context) {
			icttFlows.RegistrationAndCollateralCheck(ctx, LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Transparent proxy upgrade",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, erc20TokenRemoteLabel, upgradabilityLabel),
		func(ctx context.Context) {
			icttFlows.TransparentUpgradeableProxy(ctx, LocalNetworkInstance, TeleporterInfo)
		})
})
