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
	subsetupdater "github.com/ava-labs/icm-services/abi-bindings/go/SubsetUpdater"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

const (
	diffUpdaterByteCodeFile   = "./out/DiffUpdater.sol/DiffUpdater.json"
	subsetUpdaterByteCodeFile = "./out/SubsetUpdater.sol/SubsetUpdater.json"
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
	shardNumber uint32,
) (common.Address, [][]byte) {
	// Create the shards for initializing the p-chain validator set
	diffs := createShards(ctx, avalancheChainID, avalancheSubnetID, pChainClient, shardNumber)
	shardBytes := make([][]byte, len(diffs))
	shardHashes := make([][32]byte, len(diffs))
	for i, diff := range diffs {
		shardBytes[i] = SerializeValidatorSetDiff(&diff)
		shardHashes[i] = sha256.Sum256(shardBytes[i])
	}

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
	// Large init code / constructor; keep headroom above ~10M default caps
	gasLimit := uint64(16_000_000)
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
	return contractAddress, shardBytes
}

// DeployDiffUpdaterWithMetadata deploys DiffUpdater using Nick's method, matching
// DeploySubsetUpdater. Use the compiled artifact under out/ (libraries linked at build time).
func DeployDiffUpdaterWithMetadata(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	fundedKey *ecdsa.PrivateKey,
	avalancheNetworkID uint32,
	initialValidatorSetData diffupdater.ValidatorSetMetadata,
) common.Address {
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

	gasLimit := uint64(16_000_000)
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

// DeploySubsetUpdater deploys SubsetUpdater using Nick's method (same pattern as DeployDiffUpdater).
func DeploySubsetUpdater(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	fundedKey *ecdsa.PrivateKey,
	avalancheNetworkID uint32,
	initialValidatorSetData subsetupdater.ValidatorSetMetadata,
) common.Address {
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(subsetUpdaterByteCodeFile)
	Expect(err).Should(BeNil())

	subsetUpdaterABI, err := subsetupdater.SubsetUpdaterMetaData.GetAbi()
	Expect(err).Should(BeNil())
	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		subsetUpdaterABI,
		byteCode,
		avalancheNetworkID,
		initialValidatorSetData,
	)
	Expect(err).Should(BeNil())

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

	return contractAddress
}

func createShards(
	ctx context.Context,
	avalancheChainID ids.ID,
	avalancheSubnetID ids.ID,
	pChainClient *platformvm.Client,
	shardNumber uint32,
) []diffupdater.ValidatorSetDiff {
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

	// compute the size of each shard
	shardSize := len(canonicalValidatorSet.Validators) / int(shardNumber)
	// If the number of shards exceeds the number of validators, then shardSize will be 0.
	// In this case, shardSize will be set to 1 and shardNumber will be equal to the number of validators.
	if shardSize == 0 {
		shardSize = 1
		shardNumber = uint32(len(canonicalValidatorSet.Validators))
	}

	diffs := make([]diffupdater.ValidatorSetDiff, shardNumber)
	addedAccumulator := 0

	for s := 0; s < int(shardNumber); s++ {
		var validatorShard []*validators.Warp
		if s == int(shardNumber)-1 {
			validatorShard = canonicalValidatorSet.Validators[s*shardSize:]
		} else {
			validatorShard = canonicalValidatorSet.Validators[s*shardSize : (s+1)*shardSize]
		}
		changes := make([]diffupdater.ValidatorChange, len(validatorShard))
		for i, validator := range validatorShard {
			changes[i] = diffupdater.ValidatorChange{
				BlsPublicKey: validator.PublicKeyBytes,
				Weight:       validator.Weight,
			}
		}
		addedAccumulator += len(changes)
		diffs[s] = diffupdater.ValidatorSetDiff{
			AvalancheBlockchainID: avalancheChainID,
			PreviousHeight:        0,
			PreviousTimestamp:     0,
			CurrentHeight:         pChainHeight,
			CurrentTimestamp:      uint64(banffBlock.Timestamp().Unix()),
			Changes:               changes,
			NumAdded:              uint32(len(changes)),
			NewSize:               big.NewInt(int64(addedAccumulator)),
		}
	}
	return diffs
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

	data := bytes.Join([][]byte{
		codec,
		payloadType,
		diff.AvalancheBlockchainID[:],
		previousHeight,
		previousTimestamp,
		currentHeight,
		currentTimestamp,
		numChanges,
	}, nil)

	for _, change := range diff.Changes {
		data = append(data, change.BlsPublicKey...)
		weightBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(weightBytes, change.Weight)
		data = append(data, weightBytes...)
	}
	numAdded := make([]byte, 4)
	binary.BigEndian.PutUint32(numAdded, diff.NumAdded)
	data = append(data, numAdded...)

	return data
}
