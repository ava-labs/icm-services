package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
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
	teleporterInfo utils.TeleporterTestInfo,
) {
	_, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	ethereumBlockchainID := localEthereumNetwork.EthereumTestInfo().ChainID()

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

	// Deploy the TeleporterV2 contracts to Avalanche and Ethereum
	var teleporterContractAddress common.Address
	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	for _, l1 := range localAvalancheNetwork.GetAllL1Infos() {
		teleporterContractAddress = utils.DeployTeleporterV2(ctx, &l1, ecdsaVerifierContractAddress, fundedAvalancheKey)
		teleporterInfo.SetTeleporterV2(teleporterContractAddress, l1.BlockchainID)
	}
	// Deploy the Teleporter registry contracts to the Ethereum network.
	teleporterContractAddress = utils.DeployTeleporterV2(
		ctx,
		localEthereumNetwork.EthereumTestInfo(),
		ecdsaVerifierContractAddress,
		fundedEthereumKey,
	)
	teleporterInfo.SetTeleporterV2(teleporterContractAddress, localEthereumNetwork.EthereumTestInfo().ChainID())

	// send a message to the TeleporterMessengerV2 on Ethereum and retrieve the event
	ethTeleporter := teleporterInfo.TeleporterMessengerV2(localEthereumNetwork.EthereumTestInfo())

	message := teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: primaryNetworkInfo.BlockchainID,
		DestinationAddress:      common.Address{},
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(10),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte("Hello Avalanche!"),
	}
	tx, err = ethTeleporter.SendCrossChainMessage(opts, message)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())
	// get the event from the receipt
	event, err := utils.GetEventFromLogs(receipt.Logs, ethereumEcdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())
	// check that the event contains the correct message
	Expect(event.Message.Message).Should(Equal(message.Message))

	// sign the message
	attestation := signMessage(event.Message, ethereumBlockchainID, ecdsaVerifierContractAddress, ecdsaSigner)
	msg := teleportermessengerv2.TeleporterICMMessage{
		Message: teleportermessengerv2.TeleporterMessageV2{
			MessageNonce:            event.Message.MessageNonce,
			OriginSenderAddress:     event.Message.OriginSenderAddress,
			OriginTeleporterAddress: event.Message.OriginTeleporterAddress,
			DestinationBlockchainID: event.Message.DestinationBlockchainID,
			DestinationAddress:      event.Message.DestinationAddress,
			RequiredGasLimit:        event.Message.RequiredGasLimit,
			AllowedRelayerAddresses: event.Message.AllowedRelayerAddresses,
			Message:                 event.Message.Message,
		},
		SourceNetworkID:    0,
		SourceBlockchainID: ethereumBlockchainID,
		Attestation:        attestation,
	}

	// submit the signed message to the TeleporterMessengerV2 contract on Avalanche for verification
	avalancheTeleporter := teleporterInfo.TeleporterMessengerV2(&primaryNetworkInfo)
	opts, err = bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = avalancheTeleporter.ReceiveCrossChainMessage(opts, msg, common.Address{})
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())
	// get the event from the receipt
	receiptEvent, err := utils.GetEventFromLogs(receipt.Logs, avalancheTeleporter.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(receiptEvent.Message.Message).Should(Equal(msg.Message.Message))
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
