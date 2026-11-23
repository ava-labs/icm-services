package ethereum_icm_verification

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"time"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	zkstatemanager "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKStateManager"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

// Fixtures
type EthereumFixture struct {
	AnchorBeaconBlockRoot string                `json:"anchorBeaconBlockRoot"`
	ExecutionProof        ExecutionProofFixture `json:"executionProof"`
	ReceiptProof          ReceiptProofFixture   `json:"receiptProof"`
}

type ExecutionProofFixture struct {
	AnchorSlot                 uint64   `json:"anchorSlot"`
	TargetSlot                 uint64   `json:"targetSlot"`
	AnchorBeaconStateRoot      string   `json:"anchorBeaconStateRoot"`
	AnchorBeaconStateProof     []string `json:"anchorBeaconStateProof"`
	TargetBeaconStateRoot      string   `json:"targetBeaconStateRoot"`
	TargetBeaconStateProof     []string `json:"targetBeaconStateProof"`
	TargetExecutionHeaderRoot  string   `json:"targetExecutionHeaderRoot"`
	TargetExecutionHeaderProof []string `json:"targetExecutionHeaderProof"`
	TargetReceiptsRoot         string   `json:"targetReceiptsRoot"`
	TargetReceiptsProof        []string `json:"targetReceiptsProof"`
}

type ReceiptProofFixture struct {
	Proof           []string `json:"proof"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	LogIndex        uint64   `json:"logIndex"`
	ExpectedEmitter string   `json:"expectedEmitter"`
	ExpectedTopic0  string   `json:"expectedTopic0"`
}

// Helpers
func loadEthereumFixture(path string) EthereumFixture {
	data, err := os.ReadFile(path)
	Expect(err).Should(BeNil())
	var fixture EthereumFixture
	err = json.Unmarshal(data, &fixture)
	Expect(err).Should(BeNil())
	return fixture
}

func hexStringsToBytes32Array(hexStrings []string) [][32]byte {
	result := make([][32]byte, len(hexStrings))
	for i, s := range hexStrings {
		result[i] = common.HexToHash(s)
	}
	return result
}

func hexStringsToByteSlices(hexStrings []string) [][]byte {
	result := make([][]byte, len(hexStrings))
	for i, s := range hexStrings {
		result[i] = common.FromHex(s)
	}
	return result
}

// mainnetBeaconGenesisTime is the Unix timestamp of Ethereum mainnet's beacon chain genesis.
const mainnetBeaconGenesisTime = 1_606_824_023

// Fulu Ethereum Beacon Config
var fuluBeaconConfig = zkstatemanager.ExecutionBeaconConfig{
	GIndexBlockStateRoot: big.NewInt(11),
	GIndexExecRoot:       big.NewInt(88),
	GIndexBaseStateRoots: big.NewInt(70),
	StateRootsDepth:      big.NewInt(13),
	GIndexReceiptsRoot:   big.NewInt(35),
	StateRootsVectorSize: big.NewInt(8192),
}

// ZKAdapterVerifier tests the full ZKAdapter verification flow:
// 1. Deploy ZKAdapter on Avalanche C-Chain
// 2. Start it from the fixture's anchor beacon block root, which the constructor confirms
// 3. Verify an Ethereum event using pre-generated SSZ and MPT Merkle proofs from test fixtures
//
// TODO: Confirm the anchor via transition with a real SP1 Helios proof fixture, once a proof
// source for SP1 Helios proofs exists. Issue: https://github.com/ava-labs/icm-services/issues/1627
func ZKAdapterVerifier(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	zkAdapterByteCodeFile string,
	ethereumFixturePath string,
) {
	// Initialize
	fundedAddress, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	fixture := loadEthereumFixture(ethereumFixturePath)

	// Create a transaction signer
	opts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())

	// Extract, encode, and construct transaction
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(zkAdapterByteCodeFile)
	Expect(err).Should(BeNil())

	zkAdapterABI, err := zkadapter.ZKAdapterMetaData.GetAbi()
	Expect(err).Should(BeNil())

	// Start from the fixture's anchor, which the constructor confirms as a trusted beacon block root.
	// The sync committee is only used to verify transitions, which this test does not submit.
	startingState := zkstatemanager.ConsensusState{
		Head:   fixture.ExecutionProof.AnchorSlot,
		Header: common.HexToHash(fixture.AnchorBeaconBlockRoot),
	}
	permissibleTimespan := big.NewInt(int64((24 * time.Hour).Seconds()))
	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		zkAdapterABI,
		byteCode,
		big.NewInt(1), // Ethereum Mainnet
		startingState,
		fuluBeaconConfig,
		big.NewInt(mainnetBeaconGenesisTime),
		permissibleTimespan,
		// The verifier and light client vkey are only used by transition(), which this test does
		// not submit.
		common.Address{}, // verifier
		[32]byte{},       // light client vkey
		fundedAddress,
		fundedAddress,
	)
	Expect(err).Should(BeNil())

	gasLimit := uint64(8_000_000)
	zkadapterContractTransaction,
		zkadapterDeployerAddress,
		zkadapterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
		&gasLimit,
	)
	Expect(err).Should(BeNil())

	// Deploy
	utils.DeployWithNicksMethod(
		ctx,
		&primaryNetworkInfo,
		zkadapterContractTransaction,
		zkadapterDeployerAddress,
		zkadapterContractAddress,
		fundedAvalancheKey,
	)

	avalancheZkadapter, err := zkadapter.NewZKAdapter(zkadapterContractAddress, primaryNetworkInfo.EthClient)
	Expect(err).Should(BeNil())

	// Verify the starting beacon block root was confirmed as an anchor
	result, err := avalancheZkadapter.GetBeaconBlockRoot(
		&bind.CallOpts{},
		fixture.ExecutionProof.AnchorSlot,
	)
	Expect(err).Should(BeNil())
	Expect(result.Valid).Should(BeTrue())
	Expect(result.Root).Should(Equal(common.HexToHash(fixture.AnchorBeaconBlockRoot)))

	// Construct execution and receipt proofs from the fixtures
	execProof := zkadapter.ExecutionProof{
		AnchorSlot:                 fixture.ExecutionProof.AnchorSlot,
		TargetSlot:                 fixture.ExecutionProof.TargetSlot,
		AnchorBeaconStateRoot:      common.HexToHash(fixture.ExecutionProof.AnchorBeaconStateRoot),
		AnchorBeaconStateProof:     hexStringsToBytes32Array(fixture.ExecutionProof.AnchorBeaconStateProof),
		TargetBeaconStateRoot:      common.HexToHash(fixture.ExecutionProof.TargetBeaconStateRoot),
		TargetBeaconStateProof:     hexStringsToBytes32Array(fixture.ExecutionProof.TargetBeaconStateProof),
		TargetExecutionHeaderRoot:  common.HexToHash(fixture.ExecutionProof.TargetExecutionHeaderRoot),
		TargetExecutionHeaderProof: hexStringsToBytes32Array(fixture.ExecutionProof.TargetExecutionHeaderProof),
		TargetReceiptsRoot:         common.HexToHash(fixture.ExecutionProof.TargetReceiptsRoot),
		TargetReceiptsProof:        hexStringsToBytes32Array(fixture.ExecutionProof.TargetReceiptsProof),
	}

	receiptProof := zkadapter.ReceiptProof{
		Proof:           hexStringsToByteSlices(fixture.ReceiptProof.Proof),
		Key:             common.FromHex(fixture.ReceiptProof.Key),
		Value:           common.FromHex(fixture.ReceiptProof.Value),
		LogIndex:        big.NewInt(int64(fixture.ReceiptProof.LogIndex)),
		ExpectedEmitter: common.HexToAddress(fixture.ReceiptProof.ExpectedEmitter),
		ExpectedTopic0:  common.HexToHash(fixture.ReceiptProof.ExpectedTopic0),
	}

	// Submit the proofs to the contract
	tx, err := avalancheZkadapter.VerifyLogAndExtract(opts, execProof, receiptProof)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())

	// Verify the ZKEventImported event was emitted
	eventImported, err := utils.GetEventFromLogs(receipt.Logs, avalancheZkadapter.ParseZKEventImported)
	Expect(err).Should(BeNil())
	Expect(eventImported.BeaconSlot.Uint64()).Should(Equal(fixture.ExecutionProof.TargetSlot))
	Expect(eventImported.Emitter).Should(Equal(common.HexToAddress(fixture.ReceiptProof.ExpectedEmitter)))
}
