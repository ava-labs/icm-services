package ethereum_icm_verification

import (
	"context"
	"encoding/json"
	"math/big"
	"os"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	zkstatemanager "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKStateManager"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	deploymentUtils "github.com/ava-labs/icm-services/icm-contracts/utils/deployment-utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

// TODO: Add the fixture and testing flow for Boundless ZK proof verification
// See Issue: https://github.com/ava-labs/icm-services/issues/1248

// Fixtures
type SepoliaFixture struct {
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

// TODO: Define Boundless fixture once the Boundless ZK proof flow is implemented.

// Helpers
func loadSepoliaFixture(path string) SepoliaFixture {
	data, err := os.ReadFile(path)
	Expect(err).Should(BeNil())
	var fixture SepoliaFixture
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

func encodeJournal(
	preState zkstatemanager.ConsensusState,
	postState zkstatemanager.ConsensusState,
	finalizedSlot uint64,
) []byte {
	stateType, _ := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "currentJustifiedCheckpoint", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "epoch", Type: "uint64"},
			{Name: "root", Type: "bytes32"},
		}},
		{Name: "finalizedCheckpoint", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "epoch", Type: "uint64"},
			{Name: "root", Type: "bytes32"},
		}},
	})
	uint64Type, _ := abi.NewType("uint64", "uint64", nil)

	journalData, err := abi.Arguments{
		{Type: stateType},
		{Type: stateType},
		{Type: uint64Type},
	}.Pack(preState, postState, finalizedSlot)
	Expect(err).Should(BeNil())

	return journalData
}

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
// 2. Submit a trusted beacon block root into the contract state via manualTransition (calleable by admin)
// 3. Verify an Ethereum event using pre-generated SSZ and MPT Merkle proofs from test fixtures
func ZKAdapterVerifier(
	ctx context.Context,
	localAvalancheNetwork *localnetwork.LocalAvalancheNetwork,
	zkAdapterByteCodeFile string,
	sepoliaFixturePath string,
) {
	// Initialize
	fundedAddress, fundedAvalancheKey := localAvalancheNetwork.GetFundedAccountInfo()
	primaryNetworkInfo := localAvalancheNetwork.GetPrimaryNetworkInfo()
	fixture := loadSepoliaFixture(sepoliaFixturePath)

	// Create a starting consensus state
	// The starting state is an empty placeholder. This test uses manualTransition to update the current consensus state,
	// so an arbitrary starting state is acceptable.
	startingState := zkstatemanager.ConsensusState{ // TODO: Replace with a real starting state
		// once the Boundless ZK proof flow is implemented.
		CurrentJustifiedCheckpoint: zkstatemanager.ConsensusCheckpoint{
			Epoch: 0,
			Root:  [32]byte{},
		},
		FinalizedCheckpoint: zkstatemanager.ConsensusCheckpoint{
			Epoch: 0,
			Root:  [32]byte{},
		},
	}

	// Extract, encode, and construct transaction
	byteCode, err := deploymentUtils.ExtractByteCodeFromFile(zkAdapterByteCodeFile)
	Expect(err).Should(BeNil())

	zkAdapterABI, err := zkadapter.ZKAdapterMetaData.GetAbi()
	Expect(err).Should(BeNil())

	byteCode, err = deploymentUtils.AddConstructorArgsToByteCode(
		zkAdapterABI,
		byteCode,
		big.NewInt(11155111), // Sepolia
		startingState,
		fuluBeaconConfig,
		big.NewInt(86400),
		common.Address{}, // TODO: Replace with a real verifier address
		// once the Boundless ZK proof flow is implemented.
		[32]byte{}, // TODO: Replace with a real Image Id
		// once the Boundless ZK proof flow is implemented.
		fundedAddress,
		fundedAddress,
	)
	Expect(err).Should(BeNil())

	zkadapterContractTransaction,
		zkadapterDeployerAddress,
		zkadapterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		byteCode,
		nil,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
		nil,
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

	opts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, primaryNetworkInfo.EVMChainID)
	Expect(err).Should(BeNil())

	anchorBeaconBlockRoot := common.HexToHash(fixture.AnchorBeaconBlockRoot)

	// Compute the journal
	journalPostState := zkstatemanager.ConsensusState{ // TODO: Replace with a real journal state once
		// the Boundless ZK proof flow is implemented.
		CurrentJustifiedCheckpoint: zkstatemanager.ConsensusCheckpoint{
			Epoch: 0,
			Root:  [32]byte{},
		},
		FinalizedCheckpoint: zkstatemanager.ConsensusCheckpoint{
			Epoch: fixture.ExecutionProof.AnchorSlot / 32,
			Root:  anchorBeaconBlockRoot,
		},
	}

	// Transition the contract to the beacon block root in the fixture
	journalData := encodeJournal(startingState, journalPostState, fixture.ExecutionProof.AnchorSlot)
	tx, err := avalancheZkadapter.ManualTransition(opts, journalData, fixture.ExecutionProof.AnchorSlot)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())

	// Verify the beacon block root was stored
	result, err := avalancheZkadapter.GetBeaconBlockRoot(&bind.CallOpts{}, fixture.ExecutionProof.AnchorSlot)
	Expect(err).Should(BeNil())
	Expect(result.Valid).Should(BeTrue())
	Expect(result.Root).Should(Equal([32]byte(anchorBeaconBlockRoot)))

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
	tx, err = avalancheZkadapter.ProveLogAndExecute(opts, execProof, receiptProof)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, primaryNetworkInfo.EthClient, tx.Hash())

	// Verify the ZKEventImported event was emitted
	eventImported, err := utils.GetEventFromLogs(receipt.Logs, avalancheZkadapter.ParseZKEventImported)
	Expect(err).Should(BeNil())
	Expect(eventImported.BeaconSlot.Uint64()).Should(Equal(fixture.ExecutionProof.TargetSlot))
	Expect(eventImported.Emitter).Should(Equal(common.HexToAddress(fixture.ReceiptProof.ExpectedEmitter)))
}
