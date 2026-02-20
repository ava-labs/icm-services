package ethereum_icm_verification

import (
	"context"

	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	. "github.com/onsi/gomega"
)

func EcdsaVerifier(
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	localEthereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaVerifierByteCodeFile string,
) {
	ctx := context.Background()
	_, fundedKey := localAvalancheNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()

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
	// Deploy the ECDSAVerifier contract on the Ethereum chain
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	utils.DeployWithNicksMethod(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		ecdsaVerifierDeployerTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		fundedEthereumKey,
	)
}
