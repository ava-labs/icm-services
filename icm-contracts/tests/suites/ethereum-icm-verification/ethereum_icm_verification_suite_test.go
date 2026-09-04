package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"flag"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	ethereumIcmVerification "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ethereum_icm_verification"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	ecdsaVerifierByteCodeFile    = "./out/ECDSAVerifier.sol/ECDSAVerifier.json"
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
	zkAdapterByteCodeFile        = "./out/ZKAdapter.sol/ZKAdapter.json"
)

var (
	ethereumFixturePath  = envOrDefault("ETHEREUM_FIXTURE_PATH", "./tests/testdata/ethereum_fixture.json")
	boundlessFixturePath = envOrDefault("BOUNDLESS_FIXTURE_PATH", "./tests/testdata/boundless_fixture.json")
)

var (
	localAvalancheNetworkInstance *localnetwork.LocalAvalancheNetwork
	localEthereumNetworkInstance  *localnetwork.LocalEthereumNetwork
	e2eFlags                      *e2e.FlagVars
	ecdsaVerifierContractAddress  common.Address
	ecdsaSigner                   *ecdsa.PrivateKey
)

func envOrDefault(key, default_path string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return default_path
}

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

//  1. Deploy an ECDSAVerifier contract on every chain. Each flow below deploys and wires
//     up whatever validator set registry and adapter it needs on top of this.
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

	// set top-level variables
	_, fundedEthereumKey := localEthereumNetworkInstance.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetworkInstance.GetPrimaryNetworkInfo()
	_, fundedAvalancheKey := localAvalancheNetworkInstance.GetFundedAccountInfo()
	ethereumNetworkInfo := localEthereumNetworkInstance.GetNetworkInfo()[0]
	// Get a private key to sign messages from Ethereum
	var err error
	ecdsaSigner, err = ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	Expect(err).Should(BeNil())

	// =========================================================================
	// Step 1: Deploy the ECDSA verifier contract on all chains (Ethereum, Avalanche L1, Avalanche C-Chain)
	// =========================================================================
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(ecdsaVerifierByteCodeFile)
	Expect(err).Should(BeNil())

	// Generate the ECDSAVerifier deployer transaction via Nick's method
	var (
		ecdsaVerifierContractTransaction []byte
		ecdsaVerifierDeployerAddress     common.Address
	)
	ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		err = deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
		nil,
	)
	Expect(err).Should(BeNil())

	l1Info := localAvalancheNetworkInstance.GetL1Infos()[0]

	type ecdsaDeployInfo struct {
		name      string
		chainInfo testinfo.NetworkTestInfo
		chainID   *big.Int
		fundedKey *ecdsa.PrivateKey
		ethClient *ethclient.Client
	}
	deployInfo := []ecdsaDeployInfo{
		{
			name:      "C-Chain",
			chainInfo: &primaryNetworkInfo,
			chainID:   primaryNetworkInfo.EVMChainID,
			fundedKey: fundedAvalancheKey,
			ethClient: primaryNetworkInfo.EthClient,
		},
		{
			name:      "L1",
			chainInfo: &l1Info,
			chainID:   l1Info.EVMChainID,
			fundedKey: fundedAvalancheKey,
			ethClient: l1Info.EthClient,
		},
		{
			name:      "Ethereum",
			chainInfo: ethereumNetworkInfo,
			chainID:   localEthereumNetworkInstance.ChainID,
			fundedKey: fundedEthereumKey,
			ethClient: localEthereumNetworkInstance.EthClient,
		},
	}

	for _, t := range deployInfo {
		// Deploy the ECDSAVerifier contract
		utils.DeployWithNicksMethod(
			ctx,
			t.chainInfo,
			ecdsaVerifierContractTransaction,
			ecdsaVerifierDeployerAddress,
			ecdsaVerifierContractAddress,
			t.fundedKey,
		)

		// Bind and initialize
		verifier, err := ecdsaverifier.NewECDSAVerifier(ecdsaVerifierContractAddress, t.ethClient)
		Expect(err).Should(BeNil())
		opts, err := bind.NewKeyedTransactorWithChainID(t.fundedKey, t.chainID)
		Expect(err).Should(BeNil())
		tx, err := verifier.Initialize(opts, crypto.PubkeyToAddress(ecdsaSigner.PublicKey))
		Expect(err).Should(BeNil())
		// Wait for the transaction to be accepted
		utils.WaitForTransactionSuccess(ctx, t.ethClient, tx.Hash())
	}

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

	ginkgo.It("Test MerkleValidatorSetRegistry",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.MerkleValidatorSetRegistry(
				ctx,
				localAvalancheNetworkInstance,
				localEthereumNetworkInstance,
				ecdsaSigner,
				ecdsaVerifierContractAddress,
			)
		})

	ginkgo.It("Test ZKValidatorSetRegistry",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.ZKValidatorSetRegistry(
				ctx,
				localEthereumNetworkInstance,
			)
		})
})
