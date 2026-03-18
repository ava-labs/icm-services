package utils

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	adapter "github.com/ava-labs/icm-services/abi-bindings/go/Adapter"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

// DeployAdapter deploys an Adapter contract that delegates to the given IAdapter contract
// depending on the direction of communication. Uses Nick's method.
func DeployAdapter(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	chain1 ids.ID,
	chain2 ids.ID,
	adapterAddress1 common.Address,
	adapterAddress2 common.Address,
	fundedKey *ecdsa.PrivateKey,
) common.Address {
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile("./out/Adapter.sol/Adapter.json")
	Expect(err).Should(BeNil())

	adapterABI, err := adapter.AdapterMetaData.GetAbi()
	Expect(err).Should(BeNil())

	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		adapterABI,
		byteCode,
		chain1,
		chain2,
		adapterAddress1,
		adapterAddress2,
	)
	Expect(err).Should(BeNil())

	transactionBytes, deployerAddress, contractAddress, err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
	)
	Expect(err).Should(BeNil())

	DeployWithNicksMethod(
		ctx,
		testInfo,
		transactionBytes,
		deployerAddress,
		contractAddress,
		fundedKey,
	)
	return contractAddress
}
