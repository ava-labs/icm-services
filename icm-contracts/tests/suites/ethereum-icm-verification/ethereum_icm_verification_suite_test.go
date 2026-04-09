package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	adapter "github.com/ava-labs/icm-services/abi-bindings/go/Adapter"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	ethereumIcmVerification "github.com/ava-labs/icm-services/icm-contracts/tests/flows/ethereum_icm_verification"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	ecdsaVerifierByteCodeFile    = "./out/ECDSAVerifier.sol/ECDSAVerifier.json"
	adapterByteCodeFile          = "./out/Adapter.sol/Adapter.json"
	warpGenesisTemplateFile      = "./tests/utils/warp-genesis-template.json"
	ethereumICMVerificationLabel = "ethereum-icm-verification"
	zkAdapterByteCodeFile        = "./out/ZKAdapter.sol/ZKAdapter.json"
	sepoliaFixturePath           = "./tests/testdata/sepolia_fixture.json"
)

var (
	localAvalancheNetworkInstance *localnetwork.LocalAvalancheNetwork
	localEthereumNetworkInstance  *localnetwork.LocalEthereumNetwork
	teleporterInfo                utils.TeleporterTestInfo
	e2eFlags                      *e2e.FlagVars
	ecdsaVerifierContractAddress  common.Address
	ecdsaSigner                   *ecdsa.PrivateKey
	adapterContractAddress        common.Address
	mockSignatureAggregator       *utils.MockSignatureAggregator
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

//  1. Deploy a DiffUpdater contract on both chains
//  2. Apply the shards to initialize the initial validator set on Ethereu (not necessary on Avalanche)
//  3. Deploy an ECDSAVerifier contract on both chains
//  4. Deploy an Adapter contract on both chains and initialize it with the ECDSAVerifier contract
//     and DiffUpdater contracts
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

	// set top-level variables
	_, fundedEthereumKey := localEthereumNetworkInstance.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetworkInstance.GetPrimaryNetworkInfo()
	_, fundedAvalancheKey := localAvalancheNetworkInstance.GetFundedAccountInfo()
	ethereumNetworkInfo := localEthereumNetworkInstance.GetNetworkInfo()[0]
	// Get a private key to sign messages from Ethereum
	var err error
	ecdsaSigner, err = ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	Expect(err).Should(BeNil())
	// Create a mock signature aggregator
	// TODO: Replace this with a real signature aggregator
	mockSignatureAggregator = utils.NewMockSignatureAggregator(primaryNetworkInfo.BlockchainID, 4)

	// =========================================================================
	// Step 1: Deploy the DiffUpdater contract on both chains
	// =========================================================================

	// first deploy the DiffUpdater contract on the Ethereum network and initialize it
	registryContractAddress, serializedShards := utils.DeployDiffUpdater(
		ctx,
		ethereumNetworkInfo,
		fundedEthereumKey,
		1,
		primaryNetworkInfo.BlockchainID,
		primaryNetworkInfo.SubnetID,
		platformvm.NewClient(primaryNetworkInfo.NodeURIs[0]),
		5,
		mockSignatureAggregator,
	)
	// sanity check
	Expect(len(serializedShards)).Should(Equal(4))

	avalancheValidatorSetRegistry, err := diffupdater.NewDiffUpdater(
		registryContractAddress,
		localEthereumNetworkInstance.EthClient,
	)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(fundedEthereumKey, localEthereumNetworkInstance.ChainID)
	Expect(err).Should(BeNil())

	// Deploy the DiffUpdater contract on Avalanche
	// N.B. We don't need to initialize the contract as it will only be used for sending messages
	contractAddress, _ := utils.DeployDiffUpdater(
		ctx,
		&primaryNetworkInfo,
		fundedAvalancheKey,
		1,
		primaryNetworkInfo.BlockchainID,
		primaryNetworkInfo.SubnetID,
		platformvm.NewClient(primaryNetworkInfo.NodeURIs[0]),
		// N.B. This must be the same as above so that the constructor arguments match
		// for both deployments
		5,
		mockSignatureAggregator,
	)
	// Ensure that the contract address is the same as the one deployed on Ethereum
	Expect(contractAddress).Should(Equal(registryContractAddress))

	// =========================================================================
	// Step 2: Apply the shards to initialize the validator set on Ethereum
	// =========================================================================
	for i, shardBytes := range serializedShards {
		shard := diffupdater.ValidatorSetShard{
			AvalancheBlockchainID: primaryNetworkInfo.BlockchainID,
			ShardNumber:           uint64(i) + 1,
		}
		tx, err := avalancheValidatorSetRegistry.UpdateValidatorSet(opts, shard, shardBytes)
		Expect(err).Should(BeNil())
		receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetworkInstance.EthClient, tx.Hash())
		if i+1 == len(serializedShards) {
			event, err := utils.GetEventFromLogs(receipt.Logs, avalancheValidatorSetRegistry.ParseValidatorSetUpdated)
			Expect(err).Should(BeNil())
			Expect(event.AvalancheBlockchainID).Should(Equal([32]byte(primaryNetworkInfo.BlockchainID)))
		}
	}
	registered, err := avalancheValidatorSetRegistry.IsRegistered(&bind.CallOpts{}, primaryNetworkInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(registered).Should(BeTrue())

	// =========================================================================
	// Step 3: Deploy the ECDSA verifier contract on both chains
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
	// Deploy the ECDSAVerifier contract on the C-Chain
	utils.DeployWithNicksMethod(
		ctx,
		&primaryNetworkInfo,
		ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedAvalancheKey,
	)

	// Initialize the ECDSAVerifier contract on the C-Chain with the `ecdsaSigner` address
	avalancheEcdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(
		ecdsaVerifierContractAddress,
		primaryNetworkInfo.EthClient,
	)
	Expect(err).Should(BeNil())
	opts, err = bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := avalancheEcdsaVerifier.Initialize(opts, crypto.PubkeyToAddress(ecdsaSigner.PublicKey))
	Expect(err).Should(BeNil())
	// Wait for the transaction to be accepted
	utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())

	// Deploy the ECDSAVerifier contract on the Ethereum chain
	utils.DeployWithNicksMethod(
		ctx,
		ethereumNetworkInfo,
		ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedEthereumKey,
	)
	// Initialize the ECDSAVerifier contract on Ethereum with the `ecdsaSigner` address
	ethereumEcdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(
		ecdsaVerifierContractAddress,
		localEthereumNetworkInstance.EthClient,
	)
	Expect(err).Should(BeNil())
	opts, err = bind.NewKeyedTransactorWithChainID(fundedEthereumKey, localEthereumNetworkInstance.ChainID)
	Expect(err).Should(BeNil())
	tx, err = ethereumEcdsaVerifier.Initialize(opts, crypto.PubkeyToAddress(ecdsaSigner.PublicKey))
	Expect(err).Should(BeNil())
	// Wait for the transaction to be accepted
	utils.WaitForTransactionSuccess(ctx, localEthereumNetworkInstance.EthClient, tx.Hash())

	// =========================================================================
	// Step 4: Deploy the Adapter contract on both chains
	// =========================================================================
	byteCode, err = deploymentUtils.ExtractByteCodeFromFile(adapterByteCodeFile)
	Expect(err).Should(BeNil())

	// Generate the Adapter deployer transaction via Nick's method
	adapterABI, err := adapter.AdapterMetaData.GetAbi()
	Expect(err).Should(BeNil())
	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		adapterABI,
		byteCode,
		primaryNetworkInfo.BlockchainID,
		ethereumNetworkInfo.ChainID(),
		ecdsaVerifierContractAddress,
		registryContractAddress,
	)
	Expect(err).Should(BeNil())
	var (
		adapterContractTransaction []byte
		adapterDeployerAddress     common.Address
	)
	adapterContractTransaction,
		adapterDeployerAddress,
		adapterContractAddress,
		err = deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
		nil,
	)
	Expect(err).Should(BeNil())
	// Deploy the Adapter contract on the C-Chain
	utils.DeployWithNicksMethod(
		ctx,
		&primaryNetworkInfo,
		adapterContractTransaction,
		adapterDeployerAddress,
		adapterContractAddress,
		fundedAvalancheKey,
	)
	// Deploy the Adapter contract on Ethereum
	utils.DeployWithNicksMethod(
		ctx,
		ethereumNetworkInfo,
		adapterContractTransaction,
		adapterDeployerAddress,
		adapterContractAddress,
		fundedEthereumKey,
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
	ginkgo.It("Test AvalancheValidatorSetRegistry",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.AvalancheValidatorSetRegistry(
				ctx,
				localAvalancheNetworkInstance,
				localEthereumNetworkInstance,
				ecdsaSigner,
				ecdsaVerifierContractAddress,
				adapterContractAddress,
				teleporterInfo,
				mockSignatureAggregator,
			)
		})

	ginkgo.It("Test ZKAdapterVerifier",
		ginkgo.Label(ethereumICMVerificationLabel),
		func(ctx context.Context) {
			ethereumIcmVerification.ZKAdapterVerifier(
				ctx,
				localAvalancheNetworkInstance,
				zkAdapterByteCodeFile,
				sepoliaFixturePath,
			)
		})
})
