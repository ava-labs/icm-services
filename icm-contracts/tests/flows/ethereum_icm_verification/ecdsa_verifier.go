package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// EcdsaVerifier Test that we can deploy an ECDSAVerifier contract on both Avalanche and Ethereum and
// uses these contracts to send a message from Ethereum to Avalanche
// 1. Deploy and initialize ECDSAVerifier contract on Avalanche
// 2. Deploy and initializer ECDSAVerifier contract on Ethereum
// 3. Send a message to the contract on Ethereum
// 4. Recover the emitted event and sign it
// 5. Submit the signed message to the contract on Avalanche
func EcdsaVerifier(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	localEthereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaVerifierByteCodeFile string,
) {
	_, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	var ethereumBlockchainID ids.ID
	copy(ethereumBlockchainID[:], localEthereumNetwork.ChainID.Bytes()[:])

	// Get a private key to sign messages from Ethereum
	ecdsaSigner, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	Expect(err).Should(BeNil())

	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(ecdsaVerifierByteCodeFile)
	Expect(err).Should(BeNil())

	// Generate the ECDSAVerifier deployer transaction via Nick's method
	ecdsaVerifierContractTransaction,
		ecdsaVerifierDeployerAddress,
		ecdsaVerifierContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
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
	opts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
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

	// send a message to the ECDSAVerifier on Ethereum and retrieve the event

	// this is a dummy message
	message := ecdsaverifier.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		OriginSenderAddress:     common.Address{},
		OriginTeleporterAddress: common.Address{},
		DestinationBlockchainID: primaryNetworkInfo.BlockchainID,
		DestinationAddress:      common.Address{},
		RequiredGasLimit:        big.NewInt(10),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []ecdsaverifier.TeleporterMessageReceipt{},
		Message:                 []byte("Hello Avalanche!"),
	}
	tx, err = ethereumEcdsaVerifier.SendMessage(opts, message)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())
	// get the event from the receipt
	event, err := utils.GetEventFromLogs(receipt.Logs, ethereumEcdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())
	Expect(event.Message).Should(Equal(message))

	// sign the message
	attestation := signMessage(event.Message, ethereumBlockchainID, ecdsaVerifierContractAddress, ecdsaSigner)
	msg := ecdsaverifier.TeleporterICMMessage{
		Message:            event.Message,
		SourceNetworkID:    0,
		SourceBlockchainID: ethereumBlockchainID,
		Attestation:        attestation,
	}

	// submit the signed message to the ECDSAVerifier contract on Avalanche for verification
	valid, err := avalancheEcdsaVerifier.VerifyMessage(&bind.CallOpts{}, msg)
	Expect(err).Should(BeNil())
	Expect(valid).Should(BeTrue())
}

// Signs a message expected by the ECDSAVerifier according to the EIP-191 standard
func signMessage(
	message ecdsaverifier.TeleporterMessageV2,
	blockchainID ids.ID,
	ecdsaVerifierContractAddress common.Address,
	signer *ecdsa.PrivateKey,
) []byte {
	parsed, err := ecdsaverifier.ECDSAVerifierMetaData.GetAbi()
	Expect(err).Should(BeNil())

	// Find the TeleporterMessageV2 type from the contract ABI
	// It's used in the sendMessage function
	sendMessageMethod := parsed.Methods["sendMessage"]
	teleporterMessageType := sendMessageMethod.Inputs[0].Type
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	arguments := abi.Arguments{
		{Type: teleporterMessageType},
		{Type: bytes32Ty},
		{Type: addressTy},
	}
	encoded, err := arguments.Pack(message, blockchainID, ecdsaVerifierContractAddress)
	Expect(err).Should(BeNil())
	eip191Prefix := []byte("\x19Ethereum Signed Message:\n32")
	digest := crypto.Keccak256(append(eip191Prefix, crypto.Keccak256(encoded)...))
	sig, err := crypto.Sign(digest, signer)
	Expect(err).Should(BeNil())
	// Adjust V from 0/1 to 27/28 format
	if sig[64] < 27 {
		sig[64] += 27
	}
	return sig
}
