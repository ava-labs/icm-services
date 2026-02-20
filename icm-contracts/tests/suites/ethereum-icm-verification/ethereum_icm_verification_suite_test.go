package ethereum_icm_verification

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	ethereumIcmVerification "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ethereum_icm_verification"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	ecdsaVerifierByteCodeFile    = "./out/ECDSAVerifier.sol/ECDSAVerifier.json"
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
)

var (
	localAvalancheNetworkInstance *localnetwork.LocalAvalancheNetwork
	localEthereumNetworkInstance  *localnetwork.LocalEthereumNetwork
	e2eFlags                      *e2e.FlagVars
)

func TestEthereumICMVerification(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	e2eFlags = e2e.RegisterFlags()
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Ethereum ICM Verification e2e test")
}

var _ = ginkgo.BeforeSuite(func() {
	// Create the local network instances
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
		1,
		1,
		e2eFlags,
	)
	log.Info("Started local Avalanche network", zap.Any("networkID", localAvalancheNetworkInstance.NetworkID))

	localEthereumNetworkInstance = localnetwork.StartLocalEthereumNetwork(ctx)
	log.Info("Started local Ethereum network", zap.Any("chainID", localEthereumNetworkInstance.ChainID))
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	localAvalancheNetworkInstance.TearDownNetwork()
	localAvalancheNetworkInstance = nil
	localEthereumNetworkInstance.TearDownNetwork()
	localEthereumNetworkInstance = nil
})

var _ = ginkgo.Describe("[Ethereum ICM Verification integration tests]", func() {
	// Ethereum ICM Verification tests
	ginkgo.It("Test ECDSAVerifier",
		ginkgo.Label(ethereumICMVerificationLabel),
		func() {
			ethereumIcmVerification.EcdsaVerifier(localAvalancheNetworkInstance, localEthereumNetworkInstance, ecdsaVerifierByteCodeFile)
		})
})
