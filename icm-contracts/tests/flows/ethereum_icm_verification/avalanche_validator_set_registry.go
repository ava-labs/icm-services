package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ava-labs/avalanchego/vms/platformvm"
	adapter "github.com/ava-labs/icm-services/abi-bindings/go/Adapter"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
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
	ecdsaVerifierByteCodeFile string,
	adapterByteCodeFile string,
	teleporterInfo utils.TeleporterTestInfo,
) {
	// set top-level variables
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	_, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	ethereumNetworkInfo := localEthereumNetwork.GetNetworkInfo()[0]
	// Get a private key to sign messages from Ethereum
	ecdsaSigner, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	Expect(err).Should(BeNil())

	// =========================================================================
	// Step 1: Deploy the DiffUpdater contract on both chains
	// =========================================================================

	// first deploy the DiffUpdater contract on the Ethereum network and initialize it
	registryContractAddress, serializedShards := utils.DeployDiffUpdater(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		fundedEthereumKey,
		1,
		primaryNetworkInfo.BlockchainID,
		primaryNetworkInfo.SubnetID,
		platformvm.NewClient(primaryNetworkInfo.NodeURIs[0]),
		5,
	)
	// sanity check
	Expect(len(serializedShards)).Should(Equal(4))

	avalancheValidatorSetRegistry, err := diffupdater.NewDiffUpdater(
		registryContractAddress,
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
		if i+1 == len(serializedShards) {
			event, err := utils.GetEventFromLogs(receipt.Logs, avalancheValidatorSetRegistry.ParseValidatorSetUpdated)
			Expect(err).Should(BeNil())
			Expect(event.AvalancheBlockchainID).Should(Equal([32]byte(primaryNetworkInfo.BlockchainID)))
		}
	}
	registered, err := avalancheValidatorSetRegistry.IsRegistered(&bind.CallOpts{}, primaryNetworkInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(registered).Should(BeTrue())

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
	)
	// Ensure that the contract address is the same as the one deployed on Ethereum
	Expect(contractAddress).Should(Equal(registryContractAddress))

	// =========================================================================
	// Step 2: Deploy the ECDSA verifier contract on both chains
	// =========================================================================

	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(ecdsaVerifierByteCodeFile)
	Expect(err).Should(BeNil())

	// Generate the ECDSAVerifier deployer transaction via Nick's method
	// N.B. Use a different gas price to generate a unique deployer address
	// 		This only shows up as an issue because we re-use Avalanche networks across test
	ecdsaGasPrice := big.NewInt(0).Add(deploymentUtils.GetDefaultContractCreationGasPrice(), big.NewInt(1))
	ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		ecdsaGasPrice,
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
		localEthereumNetwork.EthereumTestInfo(),
		ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedEthereumKey,
	)
	// Initialize the ECDSAVerifier contract on Ethereum with the `ecdsaSigner` address
	ethereumEcdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(
		ecdsaVerifierContractAddress,
		localEthereumNetwork.EthClient,
	)
	Expect(err).Should(BeNil())
	opts, err = bind.NewKeyedTransactorWithChainID(fundedEthereumKey, localEthereumNetwork.ChainID)
	Expect(err).Should(BeNil())
	tx, err = ethereumEcdsaVerifier.Initialize(opts, crypto.PubkeyToAddress(ecdsaSigner.PublicKey))
	Expect(err).Should(BeNil())
	// Wait for the transaction to be accepted
	utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())

	// =========================================================================
	// Step 3: Deploy the Adapter contract on both chains
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
	adapterContractTransaction,
		adapterDeployerAddress,
		adapterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
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

	// =========================================================================
	// Step 4: Deploy the TeleporterMessengerV2 contract on both chains
	// =========================================================================

	var teleporterContractAddress common.Address
	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	for _, l1 := range localAvalancheNetwork.GetAllL1Infos() {
		teleporterContractAddress = utils.DeployTeleporterV2(ctx, &l1, adapterContractAddress, fundedAvalancheKey)
		teleporterInfo.SetTeleporterV2(teleporterContractAddress, l1.BlockchainID)
	}
	// Deploy the Teleporter registry contracts to the Ethereum network.
	teleporterContractAddress = utils.DeployTeleporterV2(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		adapterContractAddress,
		fundedEthereumKey,
	)
	teleporterInfo.SetTeleporterV2(teleporterContractAddress, localEthereumNetwork.EthereumTestInfo().ChainID())

}
