// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

const (
	zkAdapterByteCodeFile = "./out/ZKAdapter.sol/ZKAdapter.json"

	// ZKAdapterSourceChainID is the source chain ID the test ZKAdapter is deployed with. The
	// adapter's constructor only accepts Ethereum mainnet (1) or Sepolia (11155111), and the local
	// Ethereum network runs as 1337, so tests use Sepolia's ID. The adapter never checks it against
	// the chain it actually runs on: it only uses it to label where messages come from, and the
	// relayer reads it from the adapter, so both sides agree.
	ZKAdapterSourceChainID = 11155111

	zkAdapterPermissibleTimespan = 24 * time.Hour
)

// zkAdapterFuluBeaconConfig holds the generalized indices for proving a receipts root from a Fulu
// beacon block root. Matches fuluBeaconConfig in
// icm-contracts/tests/flows/ethereum_icm_verification/zk_verifier.go.
var zkAdapterFuluBeaconConfig = zkadapter.ExecutionBeaconConfig{
	GIndexBlockStateRoot: big.NewInt(11),
	GIndexExecRoot:       big.NewInt(88),
	GIndexBaseStateRoots: big.NewInt(70),
	StateRootsDepth:      big.NewInt(13),
	GIndexReceiptsRoot:   big.NewInt(35),
	StateRootsVectorSize: big.NewInt(8192),
}

// DeployZKAdapter deploys a ZKAdapter contract using Nick's method. The contract's address
// depends on its constructor arguments, so deploying with the same [beaconGenesisTime] and
// [admin] on several chains gives the adapter the same address on each, which verifyMessage
// requires of the source-side adapter.
//
// The adapter starts from an empty consensus state at slot 0 and is not given a verifier, so
// it accepts no proven transition. Tests set its anchors with the admin's manualTransition.
// The constructor rejects a starting slot that begins in the future, so [beaconGenesisTime]
// must not be later than the latest block time on any chain the adapter is deployed to.
func DeployZKAdapter(
	ctx context.Context,
	testInfo testinfo.NetworkTestInfo,
	fundedKey *ecdsa.PrivateKey,
	beaconGenesisTime uint64,
	admin common.Address,
) common.Address {
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(zkAdapterByteCodeFile)
	Expect(err).Should(BeNil())

	zkAdapterABI, err := zkadapter.ZKAdapterMetaData.GetAbi()
	Expect(err).Should(BeNil())

	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		zkAdapterABI,
		byteCode,
		big.NewInt(ZKAdapterSourceChainID),
		zkadapter.ConsensusState{},
		zkAdapterFuluBeaconConfig,
		new(big.Int).SetUint64(beaconGenesisTime),
		big.NewInt(int64(zkAdapterPermissibleTimespan.Seconds())),
		// The verifier and light client vkey are only used by transition(), which checks consensus
		// proofs. Tests set anchors with manualTransition, which skips that check.
		// TODO: deploy a mock verifier and set its light client vkey once the e2e tests cover the
		// consensus updater applying proven transitions.
		// Issue: https://github.com/ava-labs/icm-services/issues/1627
		common.Address{}, // verifier
		[32]byte{},       // light client vkey
		admin,
		admin, // super admin
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
