package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// AvalancheValidatorSetRegistry Test that we can deploy a DiffUpdater contract on Ethereum and
// populate it with the validator set from the Avalanche network.
//  1. Deploy the TeleporterMessengerV2 contracts with the ECDSAVerifier/DiffUpdater adapter
//  2. Send a cross-chain message from Ethereum to Avalanche
//  3. Manually relay the message: Recover the emitted event and sign it  and submit the signed
//     message to the contract on Avalanche
func AvalancheValidatorSetRegistry(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	localEthereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaSigner *ecdsa.PrivateKey,
	ecdsaVerifierContractAddress common.Address,
	adapterContractAddress common.Address,
	teleporterInfo utils.TeleporterTestInfo,
) {
	// set top-level variables
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	_, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	ethereumNetworkInfo := localEthereumNetwork.EthereumTestInfo()
	ethereumBlockchainID := localEthereumNetwork.EthereumTestInfo().ChainID()

	// =========================================================================
	// Step 1: Deploy the TeleporterMessengerV2 contract on both chains
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
		ethereumNetworkInfo,
		adapterContractAddress,
		fundedEthereumKey,
	)
	teleporterInfo.SetTeleporterV2(teleporterContractAddress, ethereumNetworkInfo.ChainID())

	ethereumEcdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(
		ecdsaVerifierContractAddress,
		localEthereumNetwork.EthClient,
	)

	// =========================================================================
	// Step 2: Send cross-chain message from Ethereum to Avalanche
	// =========================================================================

	// send a message to the TeleporterMessengerV2 on Ethereum and retrieve the event
	ethTeleporter := teleporterInfo.TeleporterMessengerV2(ethereumNetworkInfo)

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
	opts, err := bind.NewKeyedTransactorWithChainID(fundedEthereumKey, localEthereumNetwork.ChainID)
	Expect(err).Should(BeNil())
	tx, err := ethTeleporter.SendCrossChainMessage(opts, message)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())

	// get the event from the receipt
	event, err := utils.GetEventFromLogs(receipt.Logs, ethereumEcdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())
	// check that the event contains the correct message
	Expect(event.Message.Message).Should(Equal(message.Message))

	// =========================================================================
	// Step 3: Manually relay the message
	// =========================================================================

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
	Expect(true).Should(Equal(true))
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
