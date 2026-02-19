package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/rpc"
	testmessenger "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/tests/TestMessenger"
	"github.com/ava-labs/icm-services/icm-contracts/tests/testinfo"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func DeployWithNicksMethod2(
	ctx context.Context,
	l1 *testinfo.L1TestInfo,
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) {
	// Fund the deployer address
	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11)) // 11 AVAX
	fundDeployerTx := CreateNativeTransferTransaction(
		ctx, &l1.EVMTestInfo, fundedKey, deployerAddress, fundAmount,
	)
	SendTransactionAndWaitForSuccess(ctx, l1.EthClient, fundDeployerTx)

	log.Info("Finished funding contract deployer", zap.String("blockchainID", l1.BlockchainID.Hex()))

	// Deploy contract
	rpcClient, err := rpc.DialContext(
		ctx,
		HttpToRPCURI(l1.NodeURIs[0], l1.BlockchainID.String()),
	)
	Expect(err).Should(BeNil())
	defer rpcClient.Close()

	txHash := common.Hash{}
	err = rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1.EthClient, txHash)

	contractCode, err := l1.EthClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(contractCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
}

func DeployWithNicksMethod(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) {
	// Fund the deployer address
	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11)) // 11 native token
	evmInfo := testInfo.GetEVMTestInfo()
	fundDeployerTx := CreateNativeTransferTransaction(
		ctx, evmInfo, fundedKey, deployerAddress, fundAmount,
	)
	SendTransactionAndWaitForSuccess(ctx, evmInfo.EthClient, fundDeployerTx)

	log.Info("Finished funding contract deployer", zap.String("blockchainID", testInfo.ChainID()))

	// Deploy contract
	rpcClient := testInfo.RPCClient(ctx)
	defer rpcClient.Close()

	txHash := common.Hash{}
	err := rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, txHash)

	contractCode, err := evmInfo.EthClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(contractCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
}

func DeployTestMessenger(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	teleporterManager common.Address,
	registryAddress common.Address,
	evmInfo testinfo.EVMTestInfo,
) (common.Address, *testmessenger.TestMessenger) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		evmInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, exampleMessenger, err := testmessenger.DeployTestMessenger(
		opts,
		evmInfo.EthClient,
		registryAddress,
		teleporterManager,
		big.NewInt(1),
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, tx.Hash())

	return address, exampleMessenger
}
