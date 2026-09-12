package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"

	"github.com/ava-labs/avalanchego/ids"
	merklevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

const (
	merkleValidatorSetRegistryByteCodeFile = "./out/MerkleValidatorSetRegistry.sol/MerkleValidatorSetRegistry.json"
)

// DeployMerkleValidatorSetRegistry deploys a MerkleValidatorSetRegistry contract using Nick's method,
// bootstrapping the P-chain with the provided genesis merkle root, total weight, height, and timestamp.
func DeployMerkleValidatorSetRegistry(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	fundedKey *ecdsa.PrivateKey,
	avalancheNetworkID uint32,
	pChainID ids.ID,
	pChainGenesisRoot [32]byte,
	pChainTotalWeight uint64,
	pChainHeight uint64,
	pChainTimestamp uint64,
	allowPChainFallack bool,
) common.Address {
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(merkleValidatorSetRegistryByteCodeFile)
	Expect(err).Should(BeNil())

	merkleRegistryABI, err := merklevalidatorsetregistry.MerkleValidatorSetRegistryMetaData.GetAbi()
	Expect(err).Should(BeNil())

	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		merkleRegistryABI,
		byteCode,
		avalancheNetworkID,
		pChainID,
		pChainGenesisRoot,
		pChainTotalWeight,
		pChainHeight,
		pChainTimestamp,
		allowPChainFallack,
	)
	Expect(err).Should(BeNil())

	gasLimit := uint64(10_000_000)
	transactionBytes, deployerAddress, contractAddress, err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
		&gasLimit,
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

// SerializeTeleporterMessageV2 Serializes a `TeleporterMessageV2` to bytes in the same manner as the
// `TeleporterMessengerV2` contract expects it to be serialized.
func SerializeTeleporterMessageV2(message teleportermessengerv2.TeleporterMessageV2) []byte {
	nonceBytes := make([]byte, 32)
	message.MessageNonce.FillBytes(nonceBytes)
	gasLimitBytes := make([]byte, 32)
	message.RequiredGasLimit.FillBytes(gasLimitBytes)
	relayerCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(relayerCountBytes, uint32(len(message.AllowedRelayerAddresses)))

	result := bytes.Join([][]byte{
		nonceBytes,
		message.OriginSenderAddress.Bytes(),
		message.OriginTeleporterAddress.Bytes(),
		message.DestinationBlockchainID[:],
		message.DestinationAddress.Bytes(),
		gasLimitBytes,
		relayerCountBytes,
	}, nil)

	for _, addr := range message.AllowedRelayerAddresses {
		result = append(result, addr.Bytes()...)
	}

	receiptsCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(receiptsCountBytes, uint32(len(message.Receipts)))
	result = append(result, receiptsCountBytes...)

	for _, receipt := range message.Receipts {
		receipt.ReceivedMessageNonce.FillBytes(nonceBytes)
		result = append(result, nonceBytes...)
		result = append(result, receipt.RelayerRewardAddress.Bytes()...)
	}

	result = append(result, message.Message...)

	return result
}

// SortValidators sorts validators in ascending lexicographic order of their
// uncompressed BLS public key bytes. This matches the canonical order required
// by both the contracts and the signature aggregator.
func SortValidators(validators []*validatorupdater.Validator) {
	validatorupdater.SortValidators(validators)
}
