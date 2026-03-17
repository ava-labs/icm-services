package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

const (
	diffUpdaterByteCodeFile = "./out/DiffUpdater.sol/DiffUpdater.json"
)

// DeployDiffUpdater Deploys an instance of the `DiffUpdater` contract using
// Nick's method. Crucially, this function assumes that the initial validator
// set size is small enough so that sharding is not required.
//
// Returns the shard bytes needed to initialize the first validator set
func DeployDiffUpdater(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	fundedKey *ecdsa.PrivateKey,
	avalancheNetworkID uint32,
	avalancheChainID ids.ID,
	avalancheSubnetID ids.ID,
	pChainClient *platformvm.Client,
) (common.Address, [][]byte) {
	// Get the p-chain block height
	pChainHeight, err := pChainClient.GetHeight(ctx)
	Expect(err).Should(BeNil())

	// Get the p-chain block at the given height
	pChainBlockBytes, err := pChainClient.GetBlockByHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())
	pChainBlock, err := block.Parse(block.Codec, pChainBlockBytes)
	Expect(err).Should(BeNil())
	banffBlock, ok := pChainBlock.(block.BanffBlock)
	Expect(ok).Should(BeTrue())

	// Get the validators from the block height
	rawValidators, err := pChainClient.GetValidatorsAt(ctx, avalancheSubnetID, api.Height(pChainHeight))
	Expect(err).Should(BeNil())
	canonicalValidatorSet, err := validators.FlattenValidatorSet(rawValidators)
	Expect(err).Should(BeNil())

	changes := make([]diffupdater.ValidatorChange, len(canonicalValidatorSet.Validators))
	for i, validator := range canonicalValidatorSet.Validators {
		changes[i] = diffupdater.ValidatorChange{
			BlsPublicKey: validator.PublicKeyBytes,
			Weight:       validator.Weight,
		}
	}

	// Create the initial validator set diff
	initialDiff := diffupdater.ValidatorSetDiff{
		AvalancheBlockchainID: avalancheChainID,
		PreviousHeight:        0,
		PreviousTimestamp:     0,
		CurrentHeight:         pChainHeight,
		CurrentTimestamp:      uint64(banffBlock.Timestamp().Unix()),
		Changes:               changes,
		NumAdded:              uint32(len(canonicalValidatorSet.Validators)),
		NewSize:               big.NewInt(int64(len(canonicalValidatorSet.Validators))),
	}
	serializedDiff := SerializeValidatorSetDiff(&initialDiff)
	shardHashes := make([][32]byte, 1)
	shardHashes[0] = sha256.Sum256(serializedDiff)

	// Create the metadata used to initialize the `DiffUpdater` contract
	initialValidatorSetData := diffupdater.ValidatorSetMetadata{
		AvalancheBlockchainID: avalancheChainID,
		PChainHeight:          0,
		PChainTimestamp:       0,
		ShardHashes:           shardHashes,
	}

	// Deploy the `DiffUpdater` contract
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(diffUpdaterByteCodeFile)
	Expect(err).Should(BeNil())

	diffUpdaterABI, err := diffupdater.DiffUpdaterMetaData.GetAbi()
	Expect(err).Should(BeNil())
	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		diffUpdaterABI,
		byteCode,
		avalancheNetworkID,
		initialValidatorSetData,
	)
	Expect(err).Should(BeNil())
	// Use a higher gas limit for DiffUpdater due to large init code and constructor storage operations
	gasLimit := uint64(10000000)
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

	// Return the shard bytes needed to initialize the first validator set
	shardBytes := make([][]byte, 1)
	shardBytes[0] = serializedDiff
	return contractAddress, shardBytes
}

// InitialValidatorSetHash Calculates the hash of the initial validator set being registered.
func InitialValidatorSetHash(
	changes []diffupdater.ValidatorChange,
) [32]byte {
	codec := []byte{0x00, 0x00}
	numValidators := make([]byte, 4)
	binary.BigEndian.PutUint32(numValidators, uint32(len(changes)))
	data := bytes.Join([][]byte{codec, numValidators}, nil)

	// Serialize each validator change
	for _, change := range changes {
		// Append the 96-byte uncompressed BLS public key
		data = append(data, change.BlsPublicKey...)

		// Append the 8-byte weight
		weightBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(weightBytes, change.Weight)
		data = append(data, weightBytes...)
	}
	return sha256.Sum256(data)
}

// SerializeValidatorSetDiff Serializes a `ValidatorSetDiff` to bytes in the same manner as the
// `DiffUpdater` contract expects it to be serialized. This is based on the
// `serializeValidatorSetDiff` function in the `ValidatorSets` library.
func SerializeValidatorSetDiff(
	diff *diffupdater.ValidatorSetDiff,
) []byte {
	codec := []byte{0x00, 0x00}
	payloadType := []byte{0x00, 0x00, 0x00, 0x05}

	previousHeight := make([]byte, 8)
	binary.BigEndian.PutUint64(previousHeight, diff.PreviousHeight)
	previousTimestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(previousTimestamp, diff.PreviousTimestamp)

	currentHeight := make([]byte, 8)
	binary.BigEndian.PutUint64(currentHeight, diff.CurrentHeight)
	currentTimestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(currentTimestamp, diff.CurrentTimestamp)

	numChanges := make([]byte, 4)
	binary.BigEndian.PutUint32(numChanges, uint32(len(diff.Changes)))
	numAdded := make([]byte, 4)
	binary.BigEndian.PutUint32(numAdded, diff.NumAdded)

	data := bytes.Join([][]byte{
		codec,
		payloadType,
		diff.AvalancheBlockchainID[:],
		previousHeight,
		previousTimestamp,
		currentHeight,
		currentTimestamp,
		numChanges,
		numAdded,
	}, nil)

	// Serialize each validator change
	for _, change := range diff.Changes {
		// Append the 96-byte uncompressed BLS public key
		data = append(data, change.BlsPublicKey...)

		// Append the 8-byte weight
		weightBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(weightBytes, change.Weight)
		data = append(data, weightBytes...)
	}

	return data
}
