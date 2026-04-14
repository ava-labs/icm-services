package ethereum_icm_verification

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	ethereumIcmVerification "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ethereum_icm_verification"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	ecdsaVerifierByteCodeFile    = "./out/ECDSAVerifier.sol/ECDSAVerifier.json"
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
	zkAdapterByteCodeFile        = "./out/ZKAdapter.sol/ZKAdapter.json"
	ethereumFixturePath          = "./tests/testdata/sepolia_fixture.json"
	boundlessFixturePath         = "./tests/testdata/boundless_fixture.json"
)

var (
	localAvalancheNetworkInstance *localnetwork.LocalAvalancheNetwork
	localEthereumNetworkInstance  *localnetwork.LocalEthereumNetwork
	teleporterInfo                utils.TeleporterTestInfo
	e2eFlags                      *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestEthereumICMVerification(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Ethereum ICM Verification e2e test")
}

var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	// Create the local network instances
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	localAvalancheNetworkInstance = localnetwork.NewLocalAvalancheNetwork(
		ctx,
		"ethereum-icm-verification-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:       "L1",
				EVMChainID: 12345,
				NodeCount:  1,
			},
		},
		4,
		1,
		e2eFlags,
	)
	log.Info("Started local Avalanche network", zap.Any("networkID", localAvalancheNetworkInstance.NetworkID))

	localEthereumNetworkInstance = localnetwork.StartLocalEthereumNetwork(ctx)
	log.Info("Started local Ethereum network", zap.Any("chainID", localEthereumNetworkInstance.ChainID))

	teleporterInfo = localnetwork.NewTeleporterTestInfo(
		localAvalancheNetworkInstance,
		localEthereumNetworkInstance,
	)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	localEthereumNetworkInstance.TearDownNetwork()
	localAvalancheNetworkInstance.TearDownNetwork()
	localAvalancheNetworkInstance = nil
	localEthereumNetworkInstance = nil
})

var _ = ginkgo.Describe("[Ethereum ICM Verification integration tests]", func() {
	// Ethereum ICM Verification tests
	ginkgo.It("Test ECDSAVerifier",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.EcdsaVerifier(
				ctx,
				localAvalancheNetworkInstance,
				localEthereumNetworkInstance,
				ecdsaVerifierByteCodeFile,
				teleporterInfo,
			)
		})
	ginkgo.It("Test AvalancheValidatorSetRegistry",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.AvalancheValidatorSetRegistry(
				ctx,
				localAvalancheNetworkInstance,
				localEthereumNetworkInstance,
			)
		})

	ginkgo.It("Test ZKAdapterVerifier",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.ZKAdapterVerifier(
				ctx,
				localAvalancheNetworkInstance,
				zkAdapterByteCodeFile,
				ethereumFixturePath,
				boundlessFixturePath,
			)
		})
})
