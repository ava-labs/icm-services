package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sort"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	"github.com/ava-labs/icm-services/config"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	warppayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
)

/**
* Test roundtrip ICM message verification using MerkleValidatorSetRegistry and ECDASAVerifier. 
* - Ethereum -> Avalanche L1: signed with ECDSA and verified by ECDSAVerifier on the Avalanche L1
* - Avalanche L1 -> Ethereum: signed by the L1 validator set and verified by the MerkleValidatorSetRegistry on Ethereum
*
* The test proceeds in the following steps.
* 1. Deploy MerkleValidatorSetRegistry on Ethereum with the P-Chain genesis commitment.
*	Then, deploy TeleporterMessengerV2 on both chains (ECDSAVerifier adapter on the L1,
*	Merkle adapter on Ethereum). Then, register the L1 validator set on Ethereum.
* 2. Send a cross-chain TeleporterV2 message from Ethereum to the Avalanche L1.
* 3. Manually relay by signing the message with ECDSA and submit to the L1's TeleporterMessengerV2.
* 4. Send a cross-chain TeleporterV2 message from the L1 to Ethereum.
* 5. Manually relay by aggregating the BLS signatures, building the Merkle attestation, and submitting to
*    Ethereum's TeleporterMessengerV2.
*/
func MerkleValidatorSetRegistry(
	ctx context.Context,
	avalancheNetwork *localnetwork.LocalAvalancheNetwork,
	ethereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaSigner *ecdsa.PrivateKey,
	ecdsaVerifierContractAddress common.Address,
) {
	// set top-level variables
	l1Info := avalancheNetwork.GetL1Infos()[0]
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	ethInfo := ethereumNetwork.EthereumTestInfo()
	ethereumBlockchainID := ethInfo.ChainID()
	networkID := avalancheNetwork.GetNetworkID()
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	_, fundedAvalancheKey := avalancheNetwork.GetFundedAccountInfo()
	ethereumOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, ethereumNetwork.ChainID)
	Expect(err).Should(BeNil())
	avalancheOpts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())
	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	// Step 1: Deploy contracts and register the L1 validator set on Ethereum
	// Compute the P-Chain genesis commitment to provide to the registry's constructor
	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, val := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(val.PublicKey.Serialize()),
			Weight:                     val.Weight,
		}
	}
	sort.Slice(pChainValidators, func(i, j int) bool {
		return string(pChainValidators[i].UncompressedPublicKeyBytes[:]) < 
			string(pChainValidators[j].UncompressedPublicKeyBytes[:])
	})
	pChainRoot := validatorupdater.BuildMerkleRoot(pChainValidators)
	var pChainTotalWeight uint64
	for _, v := range pChainValidators {
		pChainTotalWeight += v.Weight
	}

	// Deploy MerkleValidatorSetRegistry on Ethereum
	merkleRegistryAddr := utils.DeployMerkleValidatorSetRegistry(
		ctx, ethInfo, ethFundedKey,
		networkID, constants.PlatformChainID,
		pChainRoot, pChainTotalWeight, pChainHeight, pChainTimestamp,
	)
	
	// Deploy MerkleValidatorSetRegistry on the Avalanche L1
	merkleRegistryAddrL1 := utils.DeployMerkleValidatorSetRegistry(
    	ctx, &l1Info, fundedAvalancheKey,
    	networkID, constants.PlatformChainID,
    	pChainRoot, pChainTotalWeight, pChainHeight, pChainTimestamp,
	)
	Expect(merkleRegistryAddrL1).Should(Equal(merkleRegistryAddr))

	merkleRegistry, err := merkleregistry.NewMerkleValidatorSetRegistry(merkleRegistryAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())

	l1EcdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(
		ecdsaVerifierContractAddress, l1Info.EthClient,
	)
	Expect(err).Should(BeNil())

	// Deploy the Adapter on both chains so each teleporter binds to the same adapter address
	merkleAdapterAddr := utils.DeployAdapter(
		ctx, ethInfo,
		l1Info.BlockchainID, ethInfo.ChainID(),
		ecdsaVerifierContractAddress, merkleRegistryAddr,
		ethFundedKey,
	)
	merkleAdapterAddrL1 := utils.DeployAdapter(
		ctx, &l1Info,
		l1Info.BlockchainID, ethInfo.ChainID(),
		ecdsaVerifierContractAddress, merkleRegistryAddr,
		fundedAvalancheKey,
	)
	Expect(merkleAdapterAddrL1).Should(Equal(merkleAdapterAddr))

	// Deploy TeleporterMessengerV2 on both chains using the same adapter address
	teleporterInfo := localnetwork.NewTeleporterTestInfo(avalancheNetwork, ethereumNetwork)
	l1TeleporterAddr := utils.DeployTeleporterV2(ctx, &l1Info, merkleAdapterAddr, fundedAvalancheKey)
	teleporterInfo.SetTeleporterV2(l1TeleporterAddr, l1Info.BlockchainID)
	ethTeleporterAddr := utils.DeployTeleporterV2(ctx, ethInfo, merkleAdapterAddr, ethFundedKey)
	teleporterInfo.SetTeleporterV2(ethTeleporterAddr, ethInfo.ChainID())
	Expect(l1TeleporterAddr).Should(Equal(ethTeleporterAddr))

	// Register the L1's validator set on Ethereum under the P-Chain root of trust commitment 
	signatureAggregator := avalancheNetwork.GetSignatureAggregator()
	defer signatureAggregator.Shutdown()

	registeredHeight := registerL1ValidatorSet(
		ctx, pChainClient, pChainValidators, signatureAggregator,
		networkID, l1Info, merkleRegistry, ethereumOpts, ethereumNetwork.EthClient,
	)

	// Step 2: Send cross-chain message from Ethereum -> Avalanche L1 verifying against ECDSAVerifier 
	ethTeleporter := teleporterInfo.TeleporterMessengerV2(ethInfo)
	ethMessage := teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: l1Info.BlockchainID,
		DestinationAddress:      common.Address{},
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(10),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte("Hello Avalanche!"),
	}
	tx, err := ethTeleporter.SendCrossChainMessage(ethereumOpts, ethMessage)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, tx.Hash())

	ethEvent, err := utils.GetEventFromLogs(receipt.Logs, l1EcdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())
	Expect(ethEvent.Message.Message).Should(Equal(ethMessage.Message))

	// Step 3: Manually relay message from Ethereum -> Avalanche
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

	l1Teleporter := teleporterInfo.TeleporterMessengerV2(&l1Info)
	tx, err = l1Teleporter.ReceiveCrossChainMessage(avalancheOpts, ethMsg, common.Address{})
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	receiptEvent, err := utils.GetEventFromLogs(receipt.Logs, l1Teleporter.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(receiptEvent.Message.Message).Should(Equal(ethMsg.Message.Message))

  	// Step 4: Send cross-chain message from the Avalanche L1 -> Ethereum verifying against MerkleValidatorSetRegistry
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
	tx, err = l1Teleporter.SendCrossChainMessage(avalancheOpts, avalancheMessage)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	avalancheEvent, err := utils.GetEventFromLogs(receipt.Logs, l1Teleporter.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(avalancheEvent.Message.Message).Should(Equal(avalancheMessage.Message))

	// Step 5: Manually relay Avalanche L1 -> Ethereum verifying against MerkleValidatorSetRegistry
	attestation = signMessageAvalancheMerkle(
		ctx, receipt, l1Info, signatureAggregator,
		pChainClient, registeredHeight,
	)
	avalancheMsg := teleportermessengerv2.TeleporterICMMessage{
		Message:            avalancheEvent.Message,
		SourceNetworkID:    networkID,
		SourceBlockchainID: l1Info.BlockchainID,
		Attestation:        attestation,
	}

	tx, err = ethTeleporter.ReceiveCrossChainMessage(ethereumOpts, avalancheMsg, common.Address{})
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, tx.Hash())

	ethReceiptEvent, err := utils.GetEventFromLogs(receipt.Logs, ethTeleporter.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(ethReceiptEvent.Message.Message).Should(Equal(avalancheMsg.Message.Message))
}

// registerL1ValidatorSet builds a synthetic P-Chain emitted warp message carrying
// the L1's current validator set commitment, has the signature aggregator gather
// P-Chain BLS signatures over it, and submits it to registerValidatorSet on the
// Ethereum-side registry. It returns the P-Chain height the commitment was built from.
func registerL1ValidatorSet(
	ctx context.Context,
	pChainClient *clients.CanonicalValidatorClient,
	pChainValidators []*validatorupdater.Validator,
	aggregator *utils.SignatureAggregator,
	networkID uint32,
	l1Info testinfo.L1TestInfo,
	merkleRegistry *merkleregistry.MerkleValidatorSetRegistry,
	ethereumOpts *bind.TransactOpts,
	ethClient *ethclient.Client,
) uint64 {
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())
	l1Validators := fetchSortedL1ValidatorsAtHeight(ctx, pChainClient, l1Info.SubnetID, pChainHeight)

	l1Commitment, err := validatorupdater.NewValidatorSetMerkleCommitment(
		l1Info.BlockchainID, l1Validators, pChainHeight, pChainTimestamp,
	)
	Expect(err).Should(BeNil())

	addressedCall, err := warppayload.NewAddressedCall(nil, l1Commitment.Bytes())
	Expect(err).Should(BeNil())

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		networkID, constants.PlatformChainID, addressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	signedMsg, err := aggregator.CreateSignedMessage(
    	unsignedMsg,
    	nil,
    	ids.Empty,
    	67,
    	l1Info,
	)
	Expect(err).Should(BeNil())

	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	Expect(ok).Should(BeTrue())

	att, err := validatorupdater.NewValidatorSetMerkleAttestation(pChainValidators, bitSetSig)
	Expect(err).Should(BeNil())

	tx, err := merkleRegistry.RegisterValidatorSet(ethereumOpts, merkleregistry.ICMMessage{
		RawMessage:         l1Commitment.Bytes(),
		SourceNetworkID:    networkID,
		SourceBlockchainID: constants.PlatformChainID,
		Attestation:        att.Bytes(),
	})
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, ethClient, tx.Hash())

	event, err := utils.GetEventFromLogs(receipt.Logs, merkleRegistry.ParseValidatorSetRegistered)
	Expect(err).Should(BeNil())
	Expect(ids.ID(event.AvalancheBlockchainID)).Should(Equal(l1Info.BlockchainID))

	return pChainHeight
}

// signMessageAvalancheMerkle aggregates BLS signatures over the warp message and packages
// them with a Merkle multi-inclusion proof against the L1 commitment registered on Ethereum.
func signMessageAvalancheMerkle(
	ctx context.Context,
	receipt *types.Receipt,
	l1Info testinfo.L1TestInfo,
	aggregator *utils.SignatureAggregator,
	pChainClient *clients.CanonicalValidatorClient,
	registeredHeight uint64,
) []byte {
	signedMsg := utils.ConstructSignedWarpMessage(ctx, receipt, l1Info, l1Info, nil, aggregator)
	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	Expect(ok).Should(BeTrue())

	l1Validators := fetchSortedL1ValidatorsAtHeight(ctx, pChainClient, l1Info.SubnetID, registeredHeight)
	att, err := validatorupdater.NewValidatorSetMerkleAttestation(l1Validators, bitSetSig)
	Expect(err).Should(BeNil())
	return att.Bytes()
}

func fetchSortedL1ValidatorsAtHeight(
	ctx context.Context,
	pChainClient *clients.CanonicalValidatorClient,
	subnetID ids.ID,
	height uint64,
) []*validatorupdater.Validator {
	allSets, err := pChainClient.GetAllValidatorSets(ctx, height)
	Expect(err).Should(BeNil())
	vdrSet, ok := allSets[subnetID]
	Expect(ok).Should(BeTrue(), "subnet validators should exist at P-chain height %d", height)
	validators := make([]*validatorupdater.Validator, len(vdrSet.Validators))
	for i, vdr := range vdrSet.Validators {
		validators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(validators, func(i, j int) bool {
		return string(validators[i].UncompressedPublicKeyBytes[:]) < string(validators[j].UncompressedPublicKeyBytes[:])
	})
	return validators
}
