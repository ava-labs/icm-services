package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
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
//  3. Manually relay the message: Recover the emitted event and sign it and submit the signed
//     message to the contract on Avalanche
//  4. Send a cross-chain message from Avalanche to Ethereum
//  5. Manually relay the message: Uses a mock signature aggregator to sign the message and
//     submit the signed message to the contract on Ethereum
func AvalancheValidatorSetRegistry(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	localEthereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaSigner *ecdsa.PrivateKey,
	ecdsaVerifierContractAddress common.Address,
	adapterContractAddress common.Address,
	mockSigner *utils.MockSignatureAggregator,
) {
	// set top-level variables
	_, fundedEthereumKey := localEthereumNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	_, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	ethereumNetworkInfo := localEthereumNetwork.EthereumTestInfo()
	ethereumBlockchainID := localEthereumNetwork.EthereumTestInfo().ChainID()
	teleporterInfo := localnetwork.NewTeleporterTestInfo(
		localAvalancheNetwork,
		localEthereumNetwork,
	)

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
	Expect(err).Should(BeNil())

	// =========================================================================
	// Step 2: Send cross-chain message from Ethereum to Avalanche
	// =========================================================================

	// send a message to the TeleporterMessengerV2 on Ethereum and retrieve the event
	ethTeleporter := teleporterInfo.TeleporterMessengerV2(ethereumNetworkInfo)

	ethMessage := teleportermessengerv2.TeleporterMessageInput{
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
	tx, err := ethTeleporter.SendCrossChainMessage(opts, ethMessage)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, localEthereumNetwork.EthClient, tx.Hash())

	// get the event from the receipt
	ethEvent, err := utils.GetEventFromLogs(receipt.Logs, ethereumEcdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())
	// check that the event contains the correct message
	Expect(ethEvent.Message.Message).Should(Equal(ethMessage.Message))

	// =========================================================================
	// Step 3: Manually relay the message
	// =========================================================================

	// sign the message
	attestation := signMessageEcdsa(ethEvent.Message, ethereumBlockchainID, ecdsaVerifierContractAddress, ecdsaSigner)
	ethMsg := teleportermessengerv2.TeleporterICMMessage{
		Message: teleportermessengerv2.TeleporterMessageV2{
			MessageNonce:            ethEvent.Message.MessageNonce,
			OriginSenderAddress:     ethEvent.Message.OriginSenderAddress,
			OriginTeleporterAddress: ethEvent.Message.OriginTeleporterAddress,
			DestinationBlockchainID: ethEvent.Message.DestinationBlockchainID,
			DestinationAddress:      ethEvent.Message.DestinationAddress,
			RequiredGasLimit:        ethEvent.Message.RequiredGasLimit,
			AllowedRelayerAddresses: ethEvent.Message.AllowedRelayerAddresses,
			Message:                 ethEvent.Message.Message,
		},
		SourceNetworkID:    0,
		SourceBlockchainID: ethereumBlockchainID,
		Attestation:        attestation,
	}

	// submit the signed message to the TeleporterMessengerV2 contract on Avalanche for verification
	avalancheTeleporter := teleporterInfo.TeleporterMessengerV2(&primaryNetworkInfo)
	opts, err = bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = avalancheTeleporter.ReceiveCrossChainMessage(opts, ethMsg, common.Address{})
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())
	// get the event from the receipt
	receiptEvent, err := utils.GetEventFromLogs(receipt.Logs, avalancheTeleporter.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(receiptEvent.Message.Message).Should(Equal(ethMsg.Message.Message))
	Expect(true).Should(Equal(true))

	// =========================================================================
	// Step 4: Send cross-chain message from Avalanche to Ethereum
	// =========================================================================

	// send a message to the TeleporterMessengerV2 on Avalanche and retrieve the event
	avalancheMessage := teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: ethereumBlockchainID,
		DestinationAddress:      common.Address{},
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(10),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte("Hello Ethereum!"),
	}
	opts, err = bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = avalancheTeleporter.SendCrossChainMessage(opts, avalancheMessage)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())

	// get the event from the receipt
	avalancheEvent, err := utils.GetEventFromLogs(receipt.Logs, avalancheTeleporter.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	// check that the event contains the correct message
	Expect(avalancheEvent.Message.Message).Should(Equal(avalancheMessage.Message))

	// =========================================================================
	// Step 5: Manually relay the message to Ethereum
	// =========================================================================
	// sign the message
	if mockSigner != nil {
		attestation = mockSignMessageAvalanche(
			primaryNetworkInfo,
			avalancheEvent.Message,
			mockSigner,
		)
	} else {
		signatureAggregator := localAvalancheNetwork.GetSignatureAggregator()
		defer signatureAggregator.Shutdown()
		attestation = signMessageAvalanche(
			primaryNetworkInfo,
			avalancheEvent.Message,
			signatureAggregator,
		).Bytes()
	}
	avalancheMsg := teleportermessengerv2.TeleporterICMMessage{
		Message:            avalancheEvent.Message,
		SourceNetworkID:    1,
		SourceBlockchainID: primaryNetworkInfo.BlockchainID,
		Attestation:        attestation,
	}

	// submit the signed message to the TeleporterMessengerV2 contract on Ethereum for verification
	ethereumTeleporter := teleporterInfo.TeleporterMessengerV2(ethereumNetworkInfo)
	opts, err = bind.NewKeyedTransactorWithChainID(fundedEthereumKey, ethereumNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = ethereumTeleporter.ReceiveCrossChainMessage(opts, avalancheMsg, common.Address{})
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, ethereumNetworkInfo.EthClient, tx.Hash())
	// get the event from the receipt
	ethReceiptEvent, err := utils.GetEventFromLogs(receipt.Logs, ethereumTeleporter.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(ethReceiptEvent.Message.Message).Should(Equal(avalancheMsg.Message.Message))
	Expect(true).Should(Equal(true))
}

// Signs a message expected by the ECDSAVerifier according to the EIP-191 standard
func signMessageEcdsa(
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

func mockSignMessageAvalanche(
	primaryNetworkInfo testinfo.L1TestInfo,
	message teleportermessengerv2.TeleporterMessageV2,
	mockSigner *utils.MockSignatureAggregator,
) []byte {
	addressedCall, err := payload.NewAddressedCall(
		nil,
		utils.SerializeTeleporterMessageV2(message))
	Expect(err).Should(BeNil())
	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		1,
		primaryNetworkInfo.ChainID(),
		addressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	// Get the signers for this blockchain
	signers := mockSigner.Signers[primaryNetworkInfo.ChainID()]

	// Sign the message with each signer
	signatures := make([]*bls.Signature, len(signers))
	for i, signer := range signers {
		sig, err := signer.Sign(unsignedMsg.Bytes())
		Expect(err).Should(BeNil())
		signatures[i] = sig
	}

	// Aggregate all signatures
	aggregatedSig, err := bls.AggregateSignatures(signatures)
	Expect(err).Should(BeNil())
	numBytes := (len(signers) + 7) / 8
	bitsetBytes := make([]byte, numBytes)
	for i := range signers {
		byteOffset := i / 8
		byteIdx := numBytes - 1 - byteOffset // Big-endian: reverse byte order
		bitPos := i % 8
		bitsetBytes[byteIdx] |= (1 << bitPos)
	}
	// Get uncompressed signature (192 bytes) - Solidity BLST requires uncompressed
	uncompressedSig := aggregatedSig.Serialize()

	// Construct attestation bytes: bitset || uncompressed signature
	// This matches the format expected by parseValidatorSetSignature
	return append(bitsetBytes, uncompressedSig...)
}

func signMessageAvalanche(
	primaryNetworkInfo testinfo.L1TestInfo,
	message teleportermessengerv2.TeleporterMessageV2,
	aggregator *utils.SignatureAggregator,
) *avalancheWarp.Message {
	addressedCall, err := payload.NewAddressedCall(
		nil,
		utils.SerializeTeleporterMessageV2(message))
	Expect(err).Should(BeNil())
	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		1,
		primaryNetworkInfo.ChainID(),
		addressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())
	signedMsg := utils.GetSignedMessage(
		primaryNetworkInfo,
		primaryNetworkInfo,
		unsignedMsg,
		nil,
		aggregator,
	)
	return signedMsg
}
