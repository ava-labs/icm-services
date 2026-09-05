package teleporterV2_test

import (
	"context"
	"flag"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/utils/logging"
	icttFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ictt"
	teleporterV2Flows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/teleporterV2"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
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
	registrationLabel      = "Registration"

	teleporterMessengerV2Label = "TeleporterMessengerV2"

	teleporterRegistryAddressFile = "TeleporterRegistryAddress.json"
	teleporterV2AddressesFile     = "TeleporterV2ContractAddresses.json"
)

// teleporterV2Addresses is persisted per blockchain so that a reused network
// can reconstruct the TeleporterMessengerV2 and WarpAdapter deployments.
type teleporterV2Addresses struct {
	TeleporterMessenger string `json:"teleporterMessenger"`
	WarpAdapter         string `json:"warpAdapter"`
}

var (
	log logging.Logger

	localNetworkInstance *localnetwork.LocalAvalancheNetwork
	teleporterInfo       utils.TeleporterTestInfo
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestTeleporterV2(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "TeleporterV2 e2e test")
}

var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	log = logging.NewLogger(
		"teleporterV2",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.JSON.ConsoleEncoder(),
		),
	)

	log.Info("Building all ICM service executables")
	utils.BuildAllExecutables(ctx, log)

	// Create the local network instance
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	localNetworkInstance = localnetwork.NewLocalAvalancheNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:       "A",
				EVMChainID: 12345,
				NodeCount:  2,
			},
			{
				Name:       "B",
				EVMChainID: 54321,
				NodeCount:  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	teleporterInfo = localnetwork.NewTeleporterTestInfo(localNetworkInstance)
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the L1 chains.
	_, fundedKey := localNetworkInstance.GetFundedAccountInfo()

	if e2eFlags.NetworkDir() == "" {
		for _, l1 := range localNetworkInstance.GetAllL1Infos() {
			warpAdapterAddress := utils.DeployWarpAdapterContract(ctx, &l1, fundedKey)
			teleporterContractAddress := utils.DeployTeleporterV2(ctx, &l1, warpAdapterAddress, fundedKey)
			teleporterInfo.SetTeleporterV2WarpAdapter(teleporterContractAddress, warpAdapterAddress, l1.BlockchainID)
			teleporterInfo.DeployTeleporterRegistry(ctx, &l1, fundedKey)
		}

		registryAddresseses := make(map[string]string)
		v2Addresses := make(map[string]teleporterV2Addresses)
		for blockchainID := range teleporterInfo {
			registryAddresseses[blockchainID.Hex()] = teleporterInfo.TeleporterRegistryAddress(blockchainID).Hex()
			v2Addresses[blockchainID.Hex()] = teleporterV2Addresses{
				TeleporterMessenger: teleporterInfo.TeleporterMessengerAddress(blockchainID).Hex(),
				WarpAdapter:         teleporterInfo.WarpAdapterAddress(blockchainID).Hex(),
			}
		}

		jsonData, err := json.Marshal(registryAddresseses)
		Expect(err).Should(BeNil())
		err = os.WriteFile(teleporterRegistryAddressFile, jsonData, fs.ModePerm)
		Expect(err).Should(BeNil())

		jsonData, err = json.Marshal(v2Addresses)
		Expect(err).Should(BeNil())
		err = os.WriteFile(teleporterV2AddressesFile, jsonData, fs.ModePerm)
		Expect(err).Should(BeNil())

	} else {
		// Read the Teleporter registry address from the file
		registryAddresseses := make(map[string]string)
		data, err := os.ReadFile(teleporterRegistryAddressFile)
		Expect(err).Should(BeNil())
		err = json.Unmarshal(data, &registryAddresseses)
		Expect(err).Should(BeNil())

		// Read the TeleporterMessengerV2 and WarpAdapter addresses from the file
		v2Addresses := make(map[string]teleporterV2Addresses)
		data, err = os.ReadFile(teleporterV2AddressesFile)
		Expect(err).Should(BeNil())
		err = json.Unmarshal(data, &v2Addresses)
		Expect(err).Should(BeNil())

		for _, l1 := range localNetworkInstance.GetAllL1Infos() {
			addresses, ok := v2Addresses[l1.BlockchainID.Hex()]
			Expect(ok).Should(BeTrue(), "no saved TeleporterV2 addresses for blockchain %s", l1.BlockchainID.Hex())
			teleporterInfo.SetTeleporterV2WarpAdapter(
				common.HexToAddress(addresses.TeleporterMessenger),
				common.HexToAddress(addresses.WarpAdapter),
				l1.BlockchainID,
			)
			teleporterInfo.SetTeleporterRegistry(
				common.HexToAddress(registryAddresseses[l1.BlockchainID.Hex()]),
				l1.BlockchainID,
			)
		}
	}
})

var _ = ginkgo.AfterSuite(func() {
	localNetworkInstance.TearDownNetwork()
	localNetworkInstance = nil
})

var _ = ginkgo.Describe("[ICTT Teleporter V2 integration tests]", func() {
	// Teleporter V2 tests
	ginkgo.It("Basic Relay",
		ginkgo.Label(teleporterMessengerV2Label),
		func(ctx context.Context) {
			teleporterV2Flows.BasicRelay(ctx, log, localNetworkInstance, teleporterInfo)
		})

	// ICTT tests
	ginkgo.It("Transfer an ERC20 token between two L1s",
		ginkgo.Label(icttLabel, erc20TokenHomeLabel, erc20TokenRemoteLabel),
		func(ctx context.Context) {
			icttFlows.ERC20TokenHomeERC20TokenRemote(ctx, localNetworkInstance, teleporterInfo)
		})
})
