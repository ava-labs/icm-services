// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkadapter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ConsensusCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ConsensusCheckpoint struct {
	Epoch uint64
	Root  [32]byte
}

// ConsensusData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusData struct {
	JournalData []byte
	Seal        []byte
}

// ConsensusState is an auto generated low-level Go binding around an user-defined struct.
type ConsensusState struct {
	CurrentJustifiedCheckpoint ConsensusCheckpoint
	FinalizedCheckpoint        ConsensusCheckpoint
}

// ExecutionBeaconConfig is an auto generated low-level Go binding around an user-defined struct.
type ExecutionBeaconConfig struct {
	GIndexBlockStateRoot *big.Int
	GIndexExecRoot       *big.Int
	GIndexBaseStateRoots *big.Int
	StateRootsDepth      *big.Int
	GIndexReceiptsRoot   *big.Int
	StateRootsVectorSize *big.Int
}

// ExecutionProof is an auto generated low-level Go binding around an user-defined struct.
type ExecutionProof struct {
	AnchorSlot                 uint64
	TargetSlot                 uint64
	AnchorBeaconStateRoot      [32]byte
	AnchorBeaconStateProof     [][32]byte
	TargetBeaconStateRoot      [32]byte
	TargetBeaconStateProof     [][32]byte
	TargetExecutionHeaderRoot  [32]byte
	TargetExecutionHeaderProof [][32]byte
	TargetReceiptsRoot         [32]byte
	TargetReceiptsProof        [][32]byte
}

// ReceiptProof is an auto generated low-level Go binding around an user-defined struct.
type ReceiptProof struct {
	Proof           [][]byte
	Key             []byte
	Value           []byte
	LogIndex        *big.Int
	ExpectedEmitter common.Address
	ExpectedTopic0  [32]byte
}

// TeleporterICMMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterICMMessage struct {
	Message            TeleporterMessageV2
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessageV2 is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageV2 struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	OriginTeleporterAddress common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// ZKAdapterMetaData contains all meta data concerning the ZKAdapter contract.
var ZKAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId_\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"internalType\":\"structConsensus.State\",\"name\":\"startingState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gIndexBlockStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexExecRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexBaseStateRoots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexReceiptsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsVectorSize\",\"type\":\"uint256\"}],\"internalType\":\"structExecution.BeaconConfig\",\"name\":\"beaconConfig_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"permissibleTimespan_\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"imageID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPreState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastConfirmedSlot\",\"type\":\"uint64\"}],\"name\":\"NonIncreasingSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissibleTimespanLapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedSourceChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedBeaconBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldImageID\",\"type\":\"bytes32\"}],\"name\":\"ImageIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"permissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"PermissibleTimespanUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedMessage\",\"type\":\"bytes\"}],\"name\":\"TeleporterV2MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"preEpoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"postEpoch\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"preState\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"postState\",\"type\":\"tuple\"}],\"name\":\"Transitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beaconSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"ZKEventImported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOT_PER_EPOCH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDEFINED_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"epochTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getBeaconBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"internalType\":\"structConsensus.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastConfirmedSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"}],\"name\":\"manualTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissibleTimespan\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceNetworkId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"seal\",\"type\":\"bytes\"}],\"internalType\":\"structConsensusData\",\"name\":\"consensus\",\"type\":\"tuple\"}],\"name\":\"transition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"}],\"name\":\"updateImageID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newPermissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"updatePermissibleTimespan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"anchorSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"anchorBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"anchorBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetExecutionHeaderRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetExecutionHeaderProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetReceiptsProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structExecution.Proof\",\"name\":\"execProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"expectedEmitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"expectedTopic0\",\"type\":\"bytes32\"}],\"internalType\":\"structReceipt.Proof\",\"name\":\"logProof\",\"type\":\"tuple\"}],\"name\":\"verifyLogAndExtract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051614cc4380380614cc483398101604081905261002e91610484565b888888888888888888885f0361007e5760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b60448201526064015b60405180910390fd5b60808990525f8690036100d35760405162461bcd60e51b815260206004820152601460248201527f496e76616c69642067656e657369732074696d650000000000000000000000006044820152606401610075565b60a08690526101027fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177583610248565b5061010d5f82610248565b50602088015151429061011f906102f1565b111561016d5760405162461bcd60e51b815260206004820152601c60248201527f5374617274696e6720737461746520696e2074686520667574757265000000006044820152606401610075565b505085518051600180546001600160401b039283166001600160401b031991821617909155602092830151600255978201518051600380549190931699169890981790559586015160045584516005559484015160065560408401516007556060840151600855608084015160095560a090930151600a55600e80546001600160a01b039094166001600160a01b031962ffffff909316600160a01b02929092166001600160b81b0319909416939093171790915550600d55506102308961031f565b63ffffffff1660c05250610587975050505050505050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166102e8575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556102a03390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102eb565b505f5b92915050565b5f600c61030860206001600160401b03851661055d565b610312919061055d565b60a0516102eb9190610574565b5f8160010361033057506001919050565b8162aa36a70361034257506002919050565b6040516372c559db60e01b815260048101839052602401610075565b604080519081016001600160401b038111828210171561038c57634e487b7160e01b5f52604160045260245ffd5b60405290565b5f604082840312156103a2575f5ffd5b6103aa61035e565b82519091506001600160401b03811681146103c3575f5ffd5b815260209182015191810191909152919050565b5f60c082840312156103e7575f5ffd5b60405160c081016001600160401b038111828210171561041557634e487b7160e01b5f52604160045260245ffd5b604090815283518252602080850151908301528381015190820152606080840151908201526080808401519082015260a0928301519281019290925250919050565b805162ffffff81168114610469575f5ffd5b919050565b80516001600160a01b0381168114610469575f5ffd5b5f5f5f5f5f5f5f5f5f898b0361022081121561049e575f5ffd5b8a5199506080601f19820112156104b3575f5ffd5b506104bc61035e565b6104c98c60208d01610392565b81526104d88c60608d01610392565b602082015297506104ec8b60a08c016103d7565b6101608b015190975095506105046101808b01610457565b94506105136101a08b0161046e565b6101c08b0151909450925061052b6101e08b0161046e565b915061053a6102008b0161046e565b90509295985092959850929598565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176102eb576102eb610549565b808201808211156102eb576102eb610549565b60805160a05160c0516146f76105cd5f395f81816103e00152610b1501525f818161037b01526105f601525f818161022101528181610a190152610aaf01526146f75ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c806365242a0e116100fe578063c4a6dba11161009e578063e97eb3621161006e578063e97eb362146104de578063eb97cd2c146104f1578063ebeced0e14610504578063f1faff0014610524575f5ffd5b8063c4a6dba114610464578063cf7923d414610477578063d08507081461048a578063d547741f146104cb575f5ffd5b806391d14854116100d957806391d148541461042b57806397fc007c1461043e578063a217fddf146103b9578063bfad283214610451575f5ffd5b806365242a0e146103d35780636fda4c8b146103db57806375b238fc14610417575f5ffd5b806336568abe116101695780634b0d0fa8116101445780634b0d0fa81461039d5780634d3499b5146103b057806354747fd4146103b9578063547923a3146103c0575f5ffd5b806336568abe146102d3578063378aa701146102e657806342c6498a14610376575f5ffd5b8063248a9ca3116101a4578063248a9ca3146102515780632b7ac3f3146102735780632f2ff15d1461029e578063304b9071146102b3575f5ffd5b806301ffc9a7146101ca57806304ae4327146101f25780631544298e1461021c575b5f5ffd5b6101dd6101d836600461361c565b610537565b60405190151581526020015b60405180910390f35b600e5461020890600160a01b900462ffffff1681565b60405162ffffff90911681526020016101e9565b6102437f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101e9565b61024361025f366004613643565b5f9081526020819052604090206001015490565b600e54610286906001600160a01b031681565b6040516001600160a01b0390911681526020016101e9565b6102b16102ac366004613675565b61056d565b005b6102bb600c81565b6040516001600160401b0390911681526020016101e9565b6102b16102e1366004613675565b610597565b610369604080516080810182525f818301818152606083018290528252825180840190935280835260208084019190915281019190915250604080516080810182526001546001600160401b0390811682840190815260025460608401528252825180840190935260035416825260045460208381019190915281019190915290565b6040516101e991906136dd565b6102437f000000000000000000000000000000000000000000000000000000000000000081565b6102436103ab366004613701565b6105cf565b610243600d5481565b6102435f81565b6102b16103ce36600461371a565b61061a565b6102bb602081565b6104027f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101e9565b6102435f5160206146cb5f395f51905f5281565b6101dd610439366004613675565b610688565b6102b161044c36600461373c565b6106b0565b600c546102bb906001600160401b031681565b6102b1610472366004613755565b610774565b6102b16104853660046137c1565b6107a8565b6104b6610498366004613701565b6001600160401b03165f908152600b60205260409020549081151590565b604080519283529015156020830152016101e9565b6102b16104d9366004613675565b61082c565b6102b16104ec366004613643565b610850565b6102b16104ff3660046137f7565b61088f565b61051761051236600461382e565b61094c565b6040516101e991906138e9565b6101dd6105323660046138fb565b610a8a565b5f6001600160e01b03198216637965db0b60e01b148061056757506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f8281526020819052604090206001015461058781610d41565b6105918383610d4e565b50505050565b6001600160a01b03811633146105c05760405163334bd91960e11b815260040160405180910390fd5b6105ca8282610ddd565b505050565b5f600c6105e660206001600160401b038516613945565b6105f09190613945565b610567907f000000000000000000000000000000000000000000000000000000000000000061395c565b5f5160206146cb5f395f51905f5261063181610d41565b600e5462ffffff600160a01b9091048116908316036106635760405163a9cb9e0d60e01b815260040160405180910390fd5b50600e805462ffffff909216600160a01b0262ffffff60a01b19909216919091179055565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f5160206146cb5f395f51905f526106c781610d41565b6001600160a01b0382166107225760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076657269666965722061646472657373000000000000000060448201526064015b60405180910390fd5b600e546001600160a01b03908116908316036107515760405163a9cb9e0d60e01b815260040160405180910390fd5b50600e80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5160206146cb5f395f51905f5261078b81610d41565b5f61079883850185613a95565b9050610591818260400151610e46565b5f6107b38280613b04565b8101906107c09190613a95565b90506107cc8183610eed565b600c5460408201516001600160401b0391821691161161081a57604081810151600c54915163f907012760e01b81526001600160401b03918216600482015291166024820152604401610719565b610828818260400151610e46565b5050565b5f8281526020819052604090206001015461084681610d41565b6105918383610ddd565b5f5160206146cb5f395f51905f5261086781610d41565b600d5482036108895760405163a9cb9e0d60e01b815260040160405180910390fd5b50600d55565b61089f606082016040830161373c565b6001600160a01b0316336001600160a01b0316146108f55760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b6044820152606401610719565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816040516020016109279190613cca565b60408051601f1981840301815290829052610941916138e9565b60405180910390a150565b60605f600b8161095f6020870187613701565b6001600160401b0316815260208101919091526040015f20549050806109c75760405162461bcd60e51b815260206004820152601860248201527f416e63686f7220736c6f7420697320756e646566696e656400000000000000006044820152606401610719565b6109d38185600561105a565b5f6109e38561010001358561143f565b90506109f560a085016080860161373c565b6001600160a01b0316610a0e6040870160208801613701565b6001600160401b03167f00000000000000000000000000000000000000000000000000000000000000007f17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb18860c001358860600135604051610a7a929190918252602082015260400190565b60405180910390a4949350505050565b5f80610a996060840184613b04565b810190610aa69190614003565b905060408301357f000000000000000000000000000000000000000000000000000000000000000014610b0e5760405162461bcd60e51b815260206004820152601060248201526f3130b21039b7bab931b29031b430b4b760811b6044820152606401610719565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610b456040850160208601614194565b63ffffffff1614610b865760405162461bcd60e51b815260206004820152600b60248201526a626164206e6574776f726b60a81b6044820152606401610719565b6020810151608001516001600160a01b03163014610bd45760405162461bcd60e51b815260206004820152600b60248201526a3130b21032b6b4ba3a32b960a91b6044820152606401610719565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816020015160a0015114610c375760405162461bcd60e51b815260206004820152600960248201526862616420746f70696360b81b6044820152606401610719565b805160208201516040516375f6768760e11b81525f92309263ebeced0e92610c639291906004016142b3565b5f604051808303815f875af1158015610c7e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ca591908101906143a0565b90505f81806020019051810190610cbc91906143a0565b9050610cc88580614408565b604051602001610cd89190613cca565b60405160208183030381529060405280519060200120818051906020012014610d365760405162461bcd60e51b815260206004820152601060248201526f0e0c2f2d8dec2c840dad2e6dac2e8c6d60831b6044820152606401610719565b506001949350505050565b610d4b813361178c565b50565b5f610d598383610688565b610dd6575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610d8e3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610567565b505f610567565b5f610de88383610688565b15610dd6575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610567565b60208083015180518051600180546001600160401b0392831667ffffffffffffffff199182161790915591840151600255828401518051600380549184169190941681179093558401516004558551938401515160405192949116927f955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d092610ecf929190614427565b60405180910390a361082881836020015160200151602001516117c5565b604080516080810182526001546001600160401b039081168284019081526002546060840152825282518084019093526003541682526004546020838101919091528101919091528251610f419190611820565b610f5e57604051634247f10960e11b815260040160405180910390fd5b610f6f825f0151836020015161184f565b610f8c57604051632d776d6b60e11b815260040160405180910390fd5b5f6002610f998380613b04565b604051610fa7929190614443565b602060405180830381855afa158015610fc2573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610fe59190614452565b600e549091506001600160a01b031663ab750e756110066020850185613b04565b600d54856040518563ffffffff1660e01b81526004016110299493929190614469565b5f6040518083038186803b15801561103f575f5ffd5b505afa158015611051573d5f5f3e3d5ffd5b50505050505050565b6110676020830183613701565b6001600160401b03166110806040840160208501613701565b6001600160401b0316106110e85760405162461bcd60e51b815260206004820152602960248201527f54617267657420736c6f74206d757374206265206c657373207468616e20616e60448201526818da1bdc881cdb1bdd60ba1b6064820152608401610719565b60058101546110fd6040840160208501613701565b61110a6020850185613701565b611114919061448f565b6001600160401b031611156111645760405162461bcd60e51b815260206004820152601660248201527515185c99d95d081cdb1bdd081a5cc81d1bdbc81bdb1960521b6044820152606401610719565b80545f906111b79060408501359061117f60608701876144ae565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508a925061190c915050565b9050806112065760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420616e63686f7220737461746520726f6f74000000000000006044820152606401610719565b60058201545f9061121d6040860160208701613701565b6001600160401b03166112309190614507565b90505f8184600301548560020154901b61124a919061395c565b90505f61129d60808701358361126360a08a018a6144ae565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060408a013561190c565b9050806112ec5760405162461bcd60e51b815260206004820181905260248201527f496e76616c69642074617267657420626561636f6e20737461746520726f6f746044820152606401610719565b5f6113448760c001358760010154898060e0019061130a91906144ae565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060808b013561190c565b9050806113935760405162461bcd60e51b815260206004820152601c60248201527f496e76616c696420657865637574696f6e20726f6f742070726f6f66000000006044820152606401610719565b5f6113ed88610100013588600401548a8061012001906113b391906144ae565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060c08c013561190c565b9050806114345760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081c9958d95a5c1d1cc81c9bdbdd605a1b6044820152606401610719565b505050505050505050565b6040805160018082528183019092526060915f9190816020015b606081526020019060019003908161145957905050905061147d6020840184613b04565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052508551869450909250151590506114c4576114c461451a565b60209081029190910101525f6114ed856114de86806144ae565b6114e79161452e565b84611923565b905080516001146115535760405162461bcd60e51b815260206004820152602a60248201527f496e76616c6964206e756d626572206f6620726573756c747320696e2072656360448201526932b4b83a10383937b7b360b11b6064820152608401610719565b5f815f815181106115665761156661451a565b60200260200101516020015151116115b85760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103932b1b2b4b83a10383937b7b360591b6044820152606401610719565b5f6115e76115e2835f815181106115d1576115d161451a565b602002602001015160200151611d86565b611db2565b90508060800151518560600135106116355760405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840d8dece40d2dcc8caf607b1b6044820152606401610719565b5f81608001518660600135815181106116505761165061451a565b6020026020010151905085608001602081019061166d919061373c565b6001600160a01b0316815f01516001600160a01b0316146116c65760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103637b39032b6b4ba3a32b960691b6044820152606401610719565b5f8160200151511161170e5760405162461bcd60e51b81526020600482015260116024820152704c6f6720686173206e6f20746f7069637360781b6044820152606401610719565b8560a0013581602001515f815181106117295761172961451a565b60200260200101511461177e5760405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206576656e74207369676e61747572650000000000000000006044820152606401610719565b604001519695505050505050565b6117968282610688565b6108285760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610719565b600c805467ffffffffffffffff19166001600160401b0384169081179091555f818152600b6020526040808220849055518392917f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d91a35050565b815181515f9161182f91612065565b8015611848575061184883602001518360200151612065565b9392505050565b6020808301515190820151515f91906001600160401b038083169082161161187b575f92505050610567565b5f611885826105cf565b90504281111561189a575f9350505050610567565b600e54600160a01b900462ffffff166118b3824261453a565b11156118c4575f9350505050610567565b5f600c60206118d3868661448f565b6001600160401b03166118e69190613945565b6118f09190613945565b600e54600160a01b900462ffffff161015979650505050505050565b5f8161191986868661208e565b1495945050505050565b60605f82516001600160401b0381111561193f5761193f61396f565b60405190808252806020026020018201604052801561198457816020015b604080518082019091526060808252602082015281526020019060019003908161195d5790505b5090505f84516001600160401b038111156119a1576119a161396f565b6040519080825280602002602001820160405280156119e657816020015b604080518082019091525f8152606060208201528152602001906001900390816119bf5790505b5090505f5b8551811015611a67576040518060400160405280878381518110611a1157611a1161451a565b6020026020010151805190602001208152602001878381518110611a3757611a3761451a565b6020026020010151815250828281518110611a5457611a5461451a565b60209081029190910101526001016119eb565b505f5b8451811015611d7b57848181518110611a8557611a8561451a565b6020026020010151838281518110611a9f57611a9f61451a565b60200260200101515f01819052505f6040518060400160405280878481518110611acb57611acb61451a565b602002602001015181526020015f81525090505f611af1611aec858b6121fc565b612297565b905060015b8015611d7057604080516080810182525f8082526020820181905291810191909152606080820152611b27836123eb565b15611bf0575f611b36846123ff565b90505f60028660200151611b4a9190614507565b15611b705760028660200151611b60919061454d565b611b6b90600161395c565b611b81565b60028660200151611b81919061454d565b90506040518060400160405280611b9b885f0151846124a4565b81526020015f8152509550611bb3825f01518761256c565b15611be857611bc6888360200151612591565b898881518110611bd857611bd861451a565b6020026020010151602001819052505b505050611d70565b60c083015115611c89575f611c04846125d9565b9050611c1385825f01516126f5565b15611c7c575f611c25825f0151612711565b8660200151611c34919061395c565b90506040518060400160405280611c5a885f0151600285611c55919061454d565b6124a4565b8152602001611c6a600284614507565b81525095508160200151925050611c83565b5050611d70565b50611d4d565b60e083015115611d40575f611c9d84612732565b9050611ca885612926565b15611cf2578051511515600103611c7c57611cca87825f015160200151612591565b888781518110611cdc57611cdc61451a565b6020026020010151602001819052505050611d70565b5f8160200151611d02875f612937565b60108110611d1257611d1261451a565b602002015180519091501515600103611be857611d308660016129bd565b9550806020015192505050611d4d565b825115611d4d5750611d70565b611d5a611aec8783612591565b9250508080611d6890614560565b915050611af6565b505050600101611a6a565b509095945050505050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805160a0810182525f808252606060208301819052928201528181018290526080810191909152611de4826129fe565b611e7f57602082015180515f1a60018110801590611e06575060048160ff1611155b611e485760405162461bcd60e51b8152602060048201526013602482015272556e737570706f72746564207478207479706560681b6044820152606401610719565b60ff811683526040805180820190915284518190611e689060019061453a565b8152602001611e7884600161395c565b9052935050505b5f611e8983612a35565b90508051600414611ee75760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420726563656970743a2065787065637465642034206669656c604482015261647360f01b6064820152608401610719565b611f09815f81518110611efc57611efc61451a565b6020026020010151612b3a565b8260200181905250611f3481600181518110611f2757611f2761451a565b6020026020010151612bab565b6001600160401b031660408301528051611f5b9082906002908110611efc57611efc61451a565b82606001819052505f611f8782600381518110611f7a57611f7a61451a565b6020026020010151612a35565b8051909150806001600160401b03811115611fa457611fa461396f565b60405190808252806020026020018201604052801561200157816020015b611fee60405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b815260200190600190039081611fc25790505b5060808501525f5b8181101561205c576120338382815181106120265761202661451a565b6020026020010151612bf6565b856080015182815181106120495761204961451a565b6020908102919091010152600101612009565b50505050919050565b805182515f916001600160401b0391821691161480156118485750506020908101519101511490565b5f83815b846001146121f3576120a5600286614507565b6001036121425760028482815181106120c0576120c061451a565b6020026020010151836040516020016120e3929190918252602082015260400190565b60408051601f19818403018152908290526120fd91614578565b602060405180830381855afa158015612118573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061213b9190614452565b91506121d4565b6002828583815181106121575761215761451a565b6020026020010151604051602001612179929190918252602082015260400190565b60408051601f198184030181529082905261219391614578565b602060405180830381855afa1580156121ae573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906121d19190614452565b91505b6121df60028661454d565b9450806121eb81614560565b915050612092565b50949350505050565b60605f5b835181101561225a578284828151811061221c5761221c61451a565b60200260200101515f0151036122525783818151811061223e5761223e61451a565b602002602001015160200151915050610567565b600101612200565b5060405162461bcd60e51b8152602060048201526011602482015270496e636f6d706c6574652050726f6f662160781b6044820152606401610719565b61229f6134e8565b6122a76134e8565b5f60405180604001604052808581526020015f81525090506122fe846040518060400160405280602081526020017f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421815250612dca565b1561230e57506001815292915050565b5f61232061231b86611d86565b612a35565b80519091505f81900361233a575050600182525092915050565b80600203612392575f612358835f81518110611efc57611efc61451a565b602081015190915060fc1c60028114806123725750806003145b15612383576001602087015261238b565b600160c08701525b50506123dd565b806011036123a657600160e08501526123dd565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206461746160a01b6044820152606401610719565b505061012082015292915050565b5f8160200151806105675750506040015190565b612407613548565b61240f613548565b5f61242461231b8561012001515f0151611d86565b90505f61243d82600181518110611efc57611efc61451a565b90506040518060400160405280612469612462855f81518110611efc57611efc61451a565b6001612df8565b81525f602091820181905291855260408051608081018252838152808301939093526001908301526060820192909252908301525092915050565b81516060905f6124b4848361453a565b9050838210156124fa5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606401610719565b6060811580156125185760405191505f825260208201604052612562565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015612551578051835260209283019201612539565b5050848452601f01601f1916604052505b5095945050505050565b5f61257682612711565b61257f84612711565b148015611848575061184883836126f5565b60608160400151156125a857506060810151610567565b8151156125c4576125bd8383602001516121fc565b9050610567565b5060408051602081019091525f815292915050565b6125e1613548565b6125e9613548565b5f6125fe61231b8561012001515f0151611d86565b90505f61261782600181518110611efc57611efc61451a565b90505f6004612631845f81518110611efc57611efc61451a565b5f815181106126425761264261451a565b602001015160f81c60f81b6001600160f81b031916901c60f81c600116905060405180604001604052806126a3612684865f81518110611efc57611efc61451a565b6002612691866001614589565b61269b91906145a2565b60ff16612df8565b815260ff83166020918201529085526040805160808101909152600181529081016126cd84612e2e565b81525f6020808301829052604080519283528282018152909201528501525091949350505050565b5f6126ff82612711565b6127098484612e95565b149392505050565b5f81602001516002835f0151516127289190613945565b610567919061453a565b61273a613597565b612742613597565b5f61275761231b8561012001515f0151611d86565b90506127616135b3565b5f5b601081101561285a575f612782848381518110611efc57611efc61451a565b905080516020146127e9576040805180820182525f8082528251608081018452818152602080820183905281850183905284519283528281019094526060810191909152918101919091528383601081106127df576127df61451a565b6020020152612851565b5f6127f382612e2e565b6040805180820182526001808252825160808101845290815260208082018590525f828501819052845190815280820190945260608201939093529181019190915290915084846010811061284a5761284a61451a565b6020020152505b50600101612763565b5061287961287483601081518110611efc57611efc61451a565b613007565b156128c5576040805180820182525f808252825160808101845281815260208082018390528185018390528451928352828101909452606081019190915291810191909152835261291a565b604051806040016040528060011515815260200160405180608001604052805f151581526020015f5f1b815260200160011515815260200161291386601081518110611efc57611efc61451a565b9052905283525b60208301525092915050565b5f61293082612711565b1592915050565b5f5f600283856020015161294b919061395c565b612955919061454d565b90505f600284866020015161296a919061395c565b6129749190614507565b90505f855f0151838151811061298c5761298c61451a565b016020015160f81c9050600182146129ab5760048160ff16901c6129b0565b80600f165b60ff169695505050505050565b60408051808201909152606081525f60208201526040518060400160405280845f015181526020018385602001516129f5919061395c565b90529392505050565b80515f908103612a0f57505f919050565b602082015180515f1a9060c0821015612a2b57505f9392505050565b5060019392505050565b6060612a40826129fe565b612a48575f5ffd5b5f612a528361306f565b90505f816001600160401b03811115612a6d57612a6d61396f565b604051908082528060200260200182016040528015612ab157816020015b604080518082019091525f8082526020820152815260200190600190039081612a8b5790505b5090505f612ac285602001516130f2565b8560200151612ad1919061395c565b90505f805b84811015612b2f57612ae78361316b565b9150604051806040016040528083815260200184815250848281518110612b1057612b1061451a565b6020908102919091010152612b25828461395c565b9250600101612ad6565b509195945050505050565b8051606090612b47575f5ffd5b5f5f612b5284613213565b915091505f816001600160401b03811115612b6f57612b6f61396f565b6040519080825280601f01601f191660200182016040528015612b99576020820181803683370190505b509050602081016121f3848285613255565b80515f9015801590612bbf57508151602110155b612bc7575f5ffd5b5f5f612bd284613213565b815191935091506020821015612bee5760208290036101000a90045b949350505050565b612c2260405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b5f612c2c83612a35565b90508051600314612c7f5760405162461bcd60e51b815260206004820152601e60248201527f496e76616c6964206c6f673a2065787065637465642033206669656c647300006044820152606401610719565b612ca1815f81518110612c9457612c9461451a565b60200260200101516132d5565b6001600160a01b031682528051612cc59082906002908110611efc57611efc61451a565b82604001819052505f612ce482600181518110611f7a57611f7a61451a565b80519091506004811115612d3a5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206c6f673a20746f6f206d616e7920746f70696373000000006044820152606401610719565b806001600160401b03811115612d5257612d5261396f565b604051908082528060200260200182016040528015612d7b578160200160208202803683370190505b5060208501525f5b8181101561205c57612da0838281518110611f2757611f2761451a565b6020860151805183908110612db757612db761451a565b6020908102919091010152600101612d83565b5f8151835114612ddb57505f610567565b825160208381018281209186019283209091145b95945050505050565b60608251821115612e07575f5ffd5b5f828451612e15919061453a565b905060208401612def612e28858361395c565b836132ed565b5f602082511015612e8d5760405162461bcd60e51b8152602060048201526024808201527f42797465733a3a20746f427974657333323a206461746120697320746f20736860448201526337b93a1760e11b6064820152608401610719565b506020015190565b5f5f60028460200151612ea89190614507565b90505f60028460200151612ebc9190614507565b9050808203612fac575f60028660200151612ed7919061454d565b90505f60028660200151612eeb919061454d565b90505f8415612f69578651805183908110612f0857612f0861451a565b016020015188518051600f60f81b9092169185908110612f2a57612f2a61451a565b0160200151600f60f81b1614612f47575f95505050505050610567565b612f5083614560565b9250612f5b82614560565b9150612f6681614560565b90505b5f612f77895f0151856124a4565b90505f612f87895f0151856124a4565b905082612f948383613343565b612f9e919061395c565b975050505050505050610567565b5f612fc7612fb987612711565b612fc287612711565b613429565b90505f5b81811015612ffc57612fdd8682612937565b612fe78883612937565b03612ffc57612ff581614560565b9050612fcb565b935061056792505050565b5f5f82511180156105675750815f815181106130255761302561451a565b6020910101516001600160f81b031916600360fe1b14806105675750815f815181106130535761305361451a565b6020910101516001600160f81b031916600160ff1b1492915050565b80515f90810361308057505f919050565b5f5f90505f61309284602001516130f2565b84602001516130a1919061395c565b90505f845f015185602001516130b7919061395c565b90505b808210156130e9576130cb8261316b565b6130d5908361395c565b9150826130e181614560565b9350506130ba565b50909392505050565b80515f90811a608081101561310957505f92915050565b60b8811080613124575060c08110801590613124575060f881105b156131325750600192915050565b60c081101561315f57613147600160b86145c3565b6131549060ff168261453a565b61184890600161395c565b613147600160f86145c3565b80515f908190811a6080811015613185576001915061320c565b60b88110156131ab5761319960808261453a565b6131a490600161395c565b915061320c565b60c08110156131d85760b78103600185019450806020036101000a8551046001820181019350505061320c565b60f88110156131ec5761319960c08261453a565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b5f5f5f61322384602001516130f2565b90505f818560200151613236919061395c565b90505f82865f0151613248919061453a565b9196919550909350505050565b805f0361326157505050565b60208110613299578251825261327860208461395c565b925061328560208361395c565b915061329260208261453a565b9050613261565b80156105ca575f60016132ad83602061453a565b6132b9906101006146bf565b6132c3919061453a565b84518451821691191617835250505050565b80515f906015146132e4575f5ffd5b61056782612bab565b6060816001600160401b038111156133075761330761396f565b6040519080825280601f01601f191660200182016040528015613331576020820181803683370190505b5090506020810161320c84828561343e565b5f5f61335184518451613429565b90505f5b8181101561341e5783818151811061336f5761336f61451a565b602001015160f81c60f81b6001600160f81b0319168582815181106133965761339661451a565b01602001516001600160f81b0319161461340e576133f08582815181106133bf576133bf61451a565b602001015160f81c60f81b8583815181106133dc576133dc61451a565b01602001516001600160f81b0319166134a5565b6133fb600283613945565b613405919061395c565b92505050610567565b61341781614560565b9050613355565b612def600282613945565b5f8183106134375781611848565b5090919050565b60208110613476578251825261345560208361395c565b915061346260208461395c565b925061346f60208261453a565b905061343e565b5f811561059157600161348a83602061453a565b613496906101006146bf565b6134a0919061453a565b6132c3565b5f6001600160f81b0319808316908416036134c257506002610567565b8160f81c60f01660ff168360f81c60f01660ff1603610dd657506001610567565b905290565b60408051610140810182525f808252602080830182905282840182905260608084018390526080840183905260a0840183905260c0840183905260e084018390526101008401839052845180860190955284528301529061012082015290565b604051806040016040528061356f6040518060400160405280606081526020015f81525090565b8152604080516080810182525f80825260208281018290529282015260608082015291015290565b60405180604001604052806135aa6135e1565b81526020016134e35b6040518061020001604052806010905b6135cb6135e1565b8152602001906001900390816135c35790505090565b60405180604001604052805f151581526020016134e3604080516080810182525f808252602082018190529181019190915260608082015290565b5f6020828403121561362c575f5ffd5b81356001600160e01b031981168114611848575f5ffd5b5f60208284031215613653575f5ffd5b5035919050565b80356001600160a01b0381168114613670575f5ffd5b919050565b5f5f60408385031215613686575f5ffd5b823591506136966020840161365a565b90509250929050565b6136bd82825180516001600160401b03168252602090810151910152565b60209081015180516001600160401b031660408401520151606090910152565b60808101610567828461369f565b80356001600160401b0381168114613670575f5ffd5b5f60208284031215613711575f5ffd5b611848826136eb565b5f6020828403121561372a575f5ffd5b813562ffffff81168114611848575f5ffd5b5f6020828403121561374c575f5ffd5b6118488261365a565b5f5f60208385031215613766575f5ffd5b82356001600160401b0381111561377b575f5ffd5b8301601f8101851361378b575f5ffd5b80356001600160401b038111156137a0575f5ffd5b8560208284010111156137b1575f5ffd5b6020919091019590945092505050565b5f602082840312156137d1575f5ffd5b81356001600160401b038111156137e6575f5ffd5b820160408185031215611848575f5ffd5b5f60208284031215613807575f5ffd5b81356001600160401b0381111561381c575f5ffd5b82016101208185031215611848575f5ffd5b5f5f6040838503121561383f575f5ffd5b82356001600160401b03811115613854575f5ffd5b83016101408186031215613866575f5ffd5b915060208301356001600160401b03811115613880575f5ffd5b830160c08186031215613891575f5ffd5b809150509250929050565b5f5b838110156138b657818101518382015260200161389e565b50505f910152565b5f81518084526138d581602086016020860161389c565b601f01601f19169290920160200192915050565b602081525f61184860208301846138be565b5f6020828403121561390b575f5ffd5b81356001600160401b03811115613920575f5ffd5b820160808185031215611848575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761056757610567613931565b8082018082111561056757610567613931565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156139a5576139a561396f565b60405290565b60405160c081016001600160401b03811182821017156139a5576139a561396f565b60405161014081016001600160401b03811182821017156139a5576139a561396f565b604051601f8201601f191681016001600160401b0381118282101715613a1857613a1861396f565b604052919050565b5f60408284031215613a30575f5ffd5b613a38613983565b9050613a43826136eb565b815260209182013591810191909152919050565b5f60808284031215613a67575f5ffd5b613a6f613983565b9050613a7b8383613a20565b8152613a8a8360408401613a20565b602082015292915050565b5f610120828403128015613aa7575f5ffd5b50604051606081016001600160401b0381118282101715613aca57613aca61396f565b604052613ad78484613a57565b8152613ae68460808501613a57565b6020820152613af861010084016136eb565b60408201529392505050565b5f5f8335601e19843603018112613b19575f5ffd5b8301803591506001600160401b03821115613b32575f5ffd5b602001915036819003821315613b46575f5ffd5b9250929050565b5f5f8335601e19843603018112613b62575f5ffd5b83016020810192503590506001600160401b03811115613b80575f5ffd5b8060051b3603821315613b46575f5ffd5b8183526020830192505f815f5b84811015613bcd576001600160a01b03613bb78361365a565b1686526020958601959190910190600101613b9e565b5093949350505050565b5f5f8335601e19843603018112613bec575f5ffd5b83016020810192503590506001600160401b03811115613c0a575f5ffd5b8060061b3603821315613b46575f5ffd5b8183526020830192505f815f5b84811015613bcd57813586526001600160a01b03613c486020840161365a565b1660208701526040958601959190910190600101613c28565b5f5f8335601e19843603018112613c76575f5ffd5b83016020810192503590506001600160401b03811115613c94575f5ffd5b803603821315613b46575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f90613ce290840161365a565b6001600160a01b038116604084015250613cfe6040840161365a565b6001600160a01b0381166060840152506060830135608083810191909152613d2790840161365a565b6001600160a01b03811660a08401525060a083013560c083810191909152613d5190840184613b4d565b61012060e0850152613d6861014085018284613b91565b915050613d7860e0850185613bd7565b848303601f1901610100860152613d90838284613c1b565b92505050613da2610100850185613c61565b848303601f1901610120860152613dba838284613ca2565b9695505050505050565b5f6001600160401b03821115613ddc57613ddc61396f565b5060051b60200190565b5f82601f830112613df5575f5ffd5b8135613e08613e0382613dc4565b6139f0565b8082825260208201915060208360051b860101925085831115613e29575f5ffd5b602085015b83811015612562578035835260209283019201613e2e565b5f6001600160401b03821115613e5e57613e5e61396f565b50601f01601f191660200190565b5f82601f830112613e7b575f5ffd5b8135613e89613e0382613e46565b818152846020838601011115613e9d575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f613ec6613e0384613dc4565b838152905060208101600584901b830185811115613ee2575f5ffd5b835b81811015613f1c5780356001600160401b03811115613f01575f5ffd5b613f0d88828801613e6c565b84525060209283019201613ee4565b5050509392505050565b5f82601f830112613f35575f5ffd5b61184883833560208501613eb9565b5f60c08284031215613f54575f5ffd5b613f5c6139ab565b905081356001600160401b03811115613f73575f5ffd5b613f7f84828501613f26565b82525060208201356001600160401b03811115613f9a575f5ffd5b613fa684828501613e6c565b60208301525060408201356001600160401b03811115613fc4575f5ffd5b613fd084828501613e6c565b60408301525060608281013590820152613fec6080830161365a565b608082015260a09182013591810191909152919050565b5f60208284031215614013575f5ffd5b81356001600160401b03811115614028575f5ffd5b820160408185031215614039575f5ffd5b614041613983565b81356001600160401b03811115614056575f5ffd5b82016101408187031215614068575f5ffd5b6140706139cd565b614079826136eb565b8152614087602083016136eb565b60208201526040828101359082015260608201356001600160401b038111156140ae575f5ffd5b6140ba88828501613de6565b6060830152506080828101359082015260a08201356001600160401b038111156140e2575f5ffd5b6140ee88828501613de6565b60a08301525060c0828101359082015260e08201356001600160401b03811115614116575f5ffd5b61412288828501613de6565b60e08301525061010082810135908201526101208201356001600160401b0381111561414c575f5ffd5b61415888828501613de6565b6101208301525082525060208201356001600160401b0381111561417a575f5ffd5b61418686828501613f44565b602083015250949350505050565b5f602082840312156141a4575f5ffd5b813563ffffffff81168114611848575f5ffd5b5f8151808452602084019350602083015f5b82811015613bcd5781518652602095860195909101906001016141c9565b5f60c08301825160c0855281815180845260e08701915060e08160051b88010193506020830192505f5b818110156142425760df1988860301835261422d8585516138be565b94506020938401939290920191600101614211565b505050506020830151848203602086015261425d82826138be565b9150506040830151848203604086015261427782826138be565b91505060608301516060850152608083015161429e60808601826001600160a01b03169052565b5060a083015160a08501528091505092915050565b604081526142cd6040820184516001600160401b03169052565b5f60208401516142e860608401826001600160401b03169052565b5060408401516080830152606084015161014060a084015261430e6101808401826141b7565b9050608085015160c084015260a0850151603f198483030160e085015261433582826141b7565b91505060c085015161010084015260e0850151603f198483030161012085015261435f82826141b7565b915050610100850151610140840152610120850151603f198483030161016085015261438b82826141b7565b9150508281036020840152612def81856141e7565b5f602082840312156143b0575f5ffd5b81516001600160401b038111156143c5575f5ffd5b8201601f810184136143d5575f5ffd5b80516143e3613e0382613e46565b8181528560208385010111156143f7575f5ffd5b612def82602083016020860161389c565b5f823561011e1983360301811261441d575f5ffd5b9190910192915050565b6101008101614436828561369f565b611848608083018461369f565b818382375f9101908152919050565b5f60208284031215614462575f5ffd5b5051919050565b606081525f61447c606083018688613ca2565b6020830194909452506040015292915050565b6001600160401b03828116828216039081111561056757610567613931565b5f5f8335601e198436030181126144c3575f5ffd5b8301803591506001600160401b038211156144dc575f5ffd5b6020019150600581901b3603821315613b46575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82614515576145156144f3565b500690565b634e487b7160e01b5f52603260045260245ffd5b5f611848368484613eb9565b8181038181111561056757610567613931565b5f8261455b5761455b6144f3565b500490565b5f6001820161457157614571613931565b5060010190565b5f825161441d81846020870161389c565b60ff818116838216019081111561056757610567613931565b5f60ff8316806145b4576145b46144f3565b8060ff84160691505092915050565b60ff828116828216039081111561056757610567613931565b6001815b6001841115614617578085048111156145fb576145fb613931565b600184161561460957908102905b60019390931c9280026145e0565b935093915050565b5f8261462d57506001610567565b8161463957505f610567565b816001811461464f576002811461465957614675565b6001915050610567565b60ff84111561466a5761466a613931565b50506001821b610567565b5060208310610133831016604e8410600b8410161715614698575081810a610567565b6146a45f1984846145dc565b805f19048211156146b7576146b7613931565b029392505050565b5f611848838361461f56fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a164736f6c634300081e000a",
}

// ZKAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKAdapterMetaData.ABI instead.
var ZKAdapterABI = ZKAdapterMetaData.ABI

// ZKAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKAdapterMetaData.Bin instead.
var ZKAdapterBin = ZKAdapterMetaData.Bin

// DeployZKAdapter deploys a new Ethereum contract, binding an instance of ZKAdapter to it.
func DeployZKAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId_ *big.Int, startingState ConsensusState, beaconConfig_ ExecutionBeaconConfig, genesisTime_ *big.Int, permissibleTimespan_ *big.Int, verifier common.Address, imageID_ [32]byte, admin common.Address, superAdmin common.Address) (common.Address, *types.Transaction, *ZKAdapter, error) {
	parsed, err := ZKAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKAdapterBin), backend, sourceChainId_, startingState, beaconConfig_, genesisTime_, permissibleTimespan_, verifier, imageID_, admin, superAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKAdapter{ZKAdapterCaller: ZKAdapterCaller{contract: contract}, ZKAdapterTransactor: ZKAdapterTransactor{contract: contract}, ZKAdapterFilterer: ZKAdapterFilterer{contract: contract}}, nil
}

// ZKAdapter is an auto generated Go binding around an Ethereum contract.
type ZKAdapter struct {
	ZKAdapterCaller     // Read-only binding to the contract
	ZKAdapterTransactor // Write-only binding to the contract
	ZKAdapterFilterer   // Log filterer for contract events
}

// ZKAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKAdapterSession struct {
	Contract     *ZKAdapter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKAdapterCallerSession struct {
	Contract *ZKAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZKAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKAdapterTransactorSession struct {
	Contract     *ZKAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZKAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKAdapterRaw struct {
	Contract *ZKAdapter // Generic contract binding to access the raw methods on
}

// ZKAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKAdapterCallerRaw struct {
	Contract *ZKAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// ZKAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKAdapterTransactorRaw struct {
	Contract *ZKAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKAdapter creates a new instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapter(address common.Address, backend bind.ContractBackend) (*ZKAdapter, error) {
	contract, err := bindZKAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKAdapter{ZKAdapterCaller: ZKAdapterCaller{contract: contract}, ZKAdapterTransactor: ZKAdapterTransactor{contract: contract}, ZKAdapterFilterer: ZKAdapterFilterer{contract: contract}}, nil
}

// NewZKAdapterCaller creates a new read-only instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterCaller(address common.Address, caller bind.ContractCaller) (*ZKAdapterCaller, error) {
	contract, err := bindZKAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterCaller{contract: contract}, nil
}

// NewZKAdapterTransactor creates a new write-only instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKAdapterTransactor, error) {
	contract, err := bindZKAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTransactor{contract: contract}, nil
}

// NewZKAdapterFilterer creates a new log filterer instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKAdapterFilterer, error) {
	contract, err := bindZKAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterFilterer{contract: contract}, nil
}

// bindZKAdapter binds a generic wrapper to an already deployed contract.
func bindZKAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKAdapter *ZKAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKAdapter.Contract.ZKAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKAdapter *ZKAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ZKAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKAdapter *ZKAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ZKAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKAdapter *ZKAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKAdapter *ZKAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKAdapter *ZKAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKAdapter.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) ADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.ADMINROLE(&_ZKAdapter.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) ADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.ADMINROLE(&_ZKAdapter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.DEFAULTADMINROLE(&_ZKAdapter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.DEFAULTADMINROLE(&_ZKAdapter.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) SECONDSPERSLOT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) SECONDSPERSLOT() (uint64, error) {
	return _ZKAdapter.Contract.SECONDSPERSLOT(&_ZKAdapter.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) SECONDSPERSLOT() (uint64, error) {
	return _ZKAdapter.Contract.SECONDSPERSLOT(&_ZKAdapter.CallOpts)
}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) SLOTPEREPOCH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "SLOT_PER_EPOCH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) SLOTPEREPOCH() (uint64, error) {
	return _ZKAdapter.Contract.SLOTPEREPOCH(&_ZKAdapter.CallOpts)
}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) SLOTPEREPOCH() (uint64, error) {
	return _ZKAdapter.Contract.SLOTPEREPOCH(&_ZKAdapter.CallOpts)
}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) UNDEFINEDROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "UNDEFINED_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) UNDEFINEDROOT() ([32]byte, error) {
	return _ZKAdapter.Contract.UNDEFINEDROOT(&_ZKAdapter.CallOpts)
}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) UNDEFINEDROOT() ([32]byte, error) {
	return _ZKAdapter.Contract.UNDEFINEDROOT(&_ZKAdapter.CallOpts)
}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) EpochTimestamp(opts *bind.CallOpts, epoch uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "epochTimestamp", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) EpochTimestamp(epoch uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.EpochTimestamp(&_ZKAdapter.CallOpts, epoch)
}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) EpochTimestamp(epoch uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.EpochTimestamp(&_ZKAdapter.CallOpts, epoch)
}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) GenesisTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "genesisTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) GenesisTime() (*big.Int, error) {
	return _ZKAdapter.Contract.GenesisTime(&_ZKAdapter.CallOpts)
}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) GenesisTime() (*big.Int, error) {
	return _ZKAdapter.Contract.GenesisTime(&_ZKAdapter.CallOpts)
}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterCaller) GetBeaconBlockRoot(opts *bind.CallOpts, slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "getBeaconBlockRoot", slot)

	outstruct := new(struct {
		Root  [32]byte
		Valid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterSession) GetBeaconBlockRoot(slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	return _ZKAdapter.Contract.GetBeaconBlockRoot(&_ZKAdapter.CallOpts, slot)
}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterCallerSession) GetBeaconBlockRoot(slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	return _ZKAdapter.Contract.GetBeaconBlockRoot(&_ZKAdapter.CallOpts, slot)
}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(((uint64,bytes32),(uint64,bytes32)))
func (_ZKAdapter *ZKAdapterCaller) GetCurrentState(opts *bind.CallOpts) (ConsensusState, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "getCurrentState")

	if err != nil {
		return *new(ConsensusState), err
	}

	out0 := *abi.ConvertType(out[0], new(ConsensusState)).(*ConsensusState)

	return out0, err

}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(((uint64,bytes32),(uint64,bytes32)))
func (_ZKAdapter *ZKAdapterSession) GetCurrentState() (ConsensusState, error) {
	return _ZKAdapter.Contract.GetCurrentState(&_ZKAdapter.CallOpts)
}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(((uint64,bytes32),(uint64,bytes32)))
func (_ZKAdapter *ZKAdapterCallerSession) GetCurrentState() (ConsensusState, error) {
	return _ZKAdapter.Contract.GetCurrentState(&_ZKAdapter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKAdapter.Contract.GetRoleAdmin(&_ZKAdapter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKAdapter.Contract.GetRoleAdmin(&_ZKAdapter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKAdapter.Contract.HasRole(&_ZKAdapter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKAdapter.Contract.HasRole(&_ZKAdapter.CallOpts, role, account)
}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) ImageID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "imageID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) ImageID() ([32]byte, error) {
	return _ZKAdapter.Contract.ImageID(&_ZKAdapter.CallOpts)
}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) ImageID() ([32]byte, error) {
	return _ZKAdapter.Contract.ImageID(&_ZKAdapter.CallOpts)
}

// LastConfirmedSlot is a free data retrieval call binding the contract method 0xbfad2832.
//
// Solidity: function lastConfirmedSlot() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) LastConfirmedSlot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "lastConfirmedSlot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastConfirmedSlot is a free data retrieval call binding the contract method 0xbfad2832.
//
// Solidity: function lastConfirmedSlot() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) LastConfirmedSlot() (uint64, error) {
	return _ZKAdapter.Contract.LastConfirmedSlot(&_ZKAdapter.CallOpts)
}

// LastConfirmedSlot is a free data retrieval call binding the contract method 0xbfad2832.
//
// Solidity: function lastConfirmedSlot() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) LastConfirmedSlot() (uint64, error) {
	return _ZKAdapter.Contract.LastConfirmedSlot(&_ZKAdapter.CallOpts)
}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterCaller) PermissibleTimespan(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "permissibleTimespan")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterSession) PermissibleTimespan() (*big.Int, error) {
	return _ZKAdapter.Contract.PermissibleTimespan(&_ZKAdapter.CallOpts)
}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterCallerSession) PermissibleTimespan() (*big.Int, error) {
	return _ZKAdapter.Contract.PermissibleTimespan(&_ZKAdapter.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) SourceChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "sourceChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) SourceChainId() (*big.Int, error) {
	return _ZKAdapter.Contract.SourceChainId(&_ZKAdapter.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) SourceChainId() (*big.Int, error) {
	return _ZKAdapter.Contract.SourceChainId(&_ZKAdapter.CallOpts)
}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterCaller) SourceNetworkId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "sourceNetworkId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterSession) SourceNetworkId() (uint32, error) {
	return _ZKAdapter.Contract.SourceNetworkId(&_ZKAdapter.CallOpts)
}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterCallerSession) SourceNetworkId() (uint32, error) {
	return _ZKAdapter.Contract.SourceNetworkId(&_ZKAdapter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKAdapter.Contract.SupportsInterface(&_ZKAdapter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKAdapter.Contract.SupportsInterface(&_ZKAdapter.CallOpts, interfaceId)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterSession) Verifier() (common.Address, error) {
	return _ZKAdapter.Contract.Verifier(&_ZKAdapter.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterCallerSession) Verifier() (common.Address, error) {
	return _ZKAdapter.Contract.Verifier(&_ZKAdapter.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.GrantRole(&_ZKAdapter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.GrantRole(&_ZKAdapter.TransactOpts, role, account)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc4a6dba1.
//
// Solidity: function manualTransition(bytes journalData) returns()
func (_ZKAdapter *ZKAdapterTransactor) ManualTransition(opts *bind.TransactOpts, journalData []byte) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "manualTransition", journalData)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc4a6dba1.
//
// Solidity: function manualTransition(bytes journalData) returns()
func (_ZKAdapter *ZKAdapterSession) ManualTransition(journalData []byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, journalData)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc4a6dba1.
//
// Solidity: function manualTransition(bytes journalData) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) ManualTransition(journalData []byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, journalData)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RenounceRole(&_ZKAdapter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RenounceRole(&_ZKAdapter.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RevokeRole(&_ZKAdapter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RevokeRole(&_ZKAdapter.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.Contract.SendMessage(&_ZKAdapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.Contract.SendMessage(&_ZKAdapter.TransactOpts, message)
}

// Transition is a paid mutator transaction binding the contract method 0xcf7923d4.
//
// Solidity: function transition((bytes,bytes) consensus) returns()
func (_ZKAdapter *ZKAdapterTransactor) Transition(opts *bind.TransactOpts, consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "transition", consensus)
}

// Transition is a paid mutator transaction binding the contract method 0xcf7923d4.
//
// Solidity: function transition((bytes,bytes) consensus) returns()
func (_ZKAdapter *ZKAdapterSession) Transition(consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.Contract.Transition(&_ZKAdapter.TransactOpts, consensus)
}

// Transition is a paid mutator transaction binding the contract method 0xcf7923d4.
//
// Solidity: function transition((bytes,bytes) consensus) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) Transition(consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.Contract.Transition(&_ZKAdapter.TransactOpts, consensus)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdateImageID(opts *bind.TransactOpts, newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updateImageID", newImageID)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterSession) UpdateImageID(newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateImageID(&_ZKAdapter.TransactOpts, newImageID)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdateImageID(newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateImageID(&_ZKAdapter.TransactOpts, newImageID)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdatePermissibleTimespan(opts *bind.TransactOpts, newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updatePermissibleTimespan", newPermissibleTimespan)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterSession) UpdatePermissibleTimespan(newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdatePermissibleTimespan(&_ZKAdapter.TransactOpts, newPermissibleTimespan)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdatePermissibleTimespan(newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdatePermissibleTimespan(&_ZKAdapter.TransactOpts, newPermissibleTimespan)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdateVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updateVerifier", newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterSession) UpdateVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateVerifier(&_ZKAdapter.TransactOpts, newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdateVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateVerifier(&_ZKAdapter.TransactOpts, newVerifier)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterTransactor) VerifyLogAndExtract(opts *bind.TransactOpts, execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "verifyLogAndExtract", execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterSession) VerifyLogAndExtract(execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyLogAndExtract(&_ZKAdapter.TransactOpts, execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterTransactorSession) VerifyLogAndExtract(execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyLogAndExtract(&_ZKAdapter.TransactOpts, execProof, logProof)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterTransactor) VerifyMessage(opts *bind.TransactOpts, message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "verifyMessage", message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyMessage(&_ZKAdapter.TransactOpts, message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterTransactorSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyMessage(&_ZKAdapter.TransactOpts, message)
}

// ZKAdapterConfirmedBeaconBlockIterator is returned from FilterConfirmedBeaconBlock and is used to iterate over the raw logs and unpacked data for ConfirmedBeaconBlock events raised by the ZKAdapter contract.
type ZKAdapterConfirmedBeaconBlockIterator struct {
	Event *ZKAdapterConfirmedBeaconBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterConfirmedBeaconBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterConfirmedBeaconBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterConfirmedBeaconBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterConfirmedBeaconBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterConfirmedBeaconBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterConfirmedBeaconBlock represents a ConfirmedBeaconBlock event raised by the ZKAdapter contract.
type ZKAdapterConfirmedBeaconBlock struct {
	Slot uint64
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterConfirmedBeaconBlock is a free log retrieval operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) FilterConfirmedBeaconBlock(opts *bind.FilterOpts, slot []uint64, root [][32]byte) (*ZKAdapterConfirmedBeaconBlockIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ConfirmedBeaconBlock", slotRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterConfirmedBeaconBlockIterator{contract: _ZKAdapter.contract, event: "ConfirmedBeaconBlock", logs: logs, sub: sub}, nil
}

// WatchConfirmedBeaconBlock is a free log subscription operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) WatchConfirmedBeaconBlock(opts *bind.WatchOpts, sink chan<- *ZKAdapterConfirmedBeaconBlock, slot []uint64, root [][32]byte) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ConfirmedBeaconBlock", slotRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterConfirmedBeaconBlock)
				if err := _ZKAdapter.contract.UnpackLog(event, "ConfirmedBeaconBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedBeaconBlock is a log parse operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) ParseConfirmedBeaconBlock(log types.Log) (*ZKAdapterConfirmedBeaconBlock, error) {
	event := new(ZKAdapterConfirmedBeaconBlock)
	if err := _ZKAdapter.contract.UnpackLog(event, "ConfirmedBeaconBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterImageIDUpdatedIterator is returned from FilterImageIDUpdated and is used to iterate over the raw logs and unpacked data for ImageIDUpdated events raised by the ZKAdapter contract.
type ZKAdapterImageIDUpdatedIterator struct {
	Event *ZKAdapterImageIDUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterImageIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterImageIDUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterImageIDUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterImageIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterImageIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterImageIDUpdated represents a ImageIDUpdated event raised by the ZKAdapter contract.
type ZKAdapterImageIDUpdated struct {
	NewImageID [32]byte
	OldImageID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterImageIDUpdated is a free log retrieval operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) FilterImageIDUpdated(opts *bind.FilterOpts, newImageID [][32]byte, oldImageID [][32]byte) (*ZKAdapterImageIDUpdatedIterator, error) {

	var newImageIDRule []interface{}
	for _, newImageIDItem := range newImageID {
		newImageIDRule = append(newImageIDRule, newImageIDItem)
	}
	var oldImageIDRule []interface{}
	for _, oldImageIDItem := range oldImageID {
		oldImageIDRule = append(oldImageIDRule, oldImageIDItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ImageIDUpdated", newImageIDRule, oldImageIDRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterImageIDUpdatedIterator{contract: _ZKAdapter.contract, event: "ImageIDUpdated", logs: logs, sub: sub}, nil
}

// WatchImageIDUpdated is a free log subscription operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) WatchImageIDUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterImageIDUpdated, newImageID [][32]byte, oldImageID [][32]byte) (event.Subscription, error) {

	var newImageIDRule []interface{}
	for _, newImageIDItem := range newImageID {
		newImageIDRule = append(newImageIDRule, newImageIDItem)
	}
	var oldImageIDRule []interface{}
	for _, oldImageIDItem := range oldImageID {
		oldImageIDRule = append(oldImageIDRule, oldImageIDItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ImageIDUpdated", newImageIDRule, oldImageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterImageIDUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "ImageIDUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseImageIDUpdated is a log parse operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) ParseImageIDUpdated(log types.Log) (*ZKAdapterImageIDUpdated, error) {
	event := new(ZKAdapterImageIDUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "ImageIDUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterPermissibleTimespanUpdatedIterator is returned from FilterPermissibleTimespanUpdated and is used to iterate over the raw logs and unpacked data for PermissibleTimespanUpdated events raised by the ZKAdapter contract.
type ZKAdapterPermissibleTimespanUpdatedIterator struct {
	Event *ZKAdapterPermissibleTimespanUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterPermissibleTimespanUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterPermissibleTimespanUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterPermissibleTimespanUpdated represents a PermissibleTimespanUpdated event raised by the ZKAdapter contract.
type ZKAdapterPermissibleTimespanUpdated struct {
	PermissibleTimespan *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPermissibleTimespanUpdated is a free log retrieval operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) FilterPermissibleTimespanUpdated(opts *bind.FilterOpts, permissibleTimespan []*big.Int) (*ZKAdapterPermissibleTimespanUpdatedIterator, error) {

	var permissibleTimespanRule []interface{}
	for _, permissibleTimespanItem := range permissibleTimespan {
		permissibleTimespanRule = append(permissibleTimespanRule, permissibleTimespanItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "PermissibleTimespanUpdated", permissibleTimespanRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterPermissibleTimespanUpdatedIterator{contract: _ZKAdapter.contract, event: "PermissibleTimespanUpdated", logs: logs, sub: sub}, nil
}

// WatchPermissibleTimespanUpdated is a free log subscription operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) WatchPermissibleTimespanUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterPermissibleTimespanUpdated, permissibleTimespan []*big.Int) (event.Subscription, error) {

	var permissibleTimespanRule []interface{}
	for _, permissibleTimespanItem := range permissibleTimespan {
		permissibleTimespanRule = append(permissibleTimespanRule, permissibleTimespanItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "PermissibleTimespanUpdated", permissibleTimespanRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterPermissibleTimespanUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "PermissibleTimespanUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermissibleTimespanUpdated is a log parse operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) ParsePermissibleTimespanUpdated(log types.Log) (*ZKAdapterPermissibleTimespanUpdated, error) {
	event := new(ZKAdapterPermissibleTimespanUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "PermissibleTimespanUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZKAdapter contract.
type ZKAdapterRoleAdminChangedIterator struct {
	Event *ZKAdapterRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleAdminChanged represents a RoleAdminChanged event raised by the ZKAdapter contract.
type ZKAdapterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZKAdapterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleAdminChangedIterator{contract: _ZKAdapter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleAdminChanged)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleAdminChanged(log types.Log) (*ZKAdapterRoleAdminChanged, error) {
	event := new(ZKAdapterRoleAdminChanged)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZKAdapter contract.
type ZKAdapterRoleGrantedIterator struct {
	Event *ZKAdapterRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleGranted represents a RoleGranted event raised by the ZKAdapter contract.
type ZKAdapterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKAdapterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleGrantedIterator{contract: _ZKAdapter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleGranted)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleGranted(log types.Log) (*ZKAdapterRoleGranted, error) {
	event := new(ZKAdapterRoleGranted)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZKAdapter contract.
type ZKAdapterRoleRevokedIterator struct {
	Event *ZKAdapterRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleRevoked represents a RoleRevoked event raised by the ZKAdapter contract.
type ZKAdapterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKAdapterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleRevokedIterator{contract: _ZKAdapter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleRevoked)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleRevoked(log types.Log) (*ZKAdapterRoleRevoked, error) {
	event := new(ZKAdapterRoleRevoked)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterTeleporterV2MessageSentIterator is returned from FilterTeleporterV2MessageSent and is used to iterate over the raw logs and unpacked data for TeleporterV2MessageSent events raised by the ZKAdapter contract.
type ZKAdapterTeleporterV2MessageSentIterator struct {
	Event *ZKAdapterTeleporterV2MessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterTeleporterV2MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterTeleporterV2MessageSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterTeleporterV2MessageSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterTeleporterV2MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterTeleporterV2MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterTeleporterV2MessageSent represents a TeleporterV2MessageSent event raised by the ZKAdapter contract.
type ZKAdapterTeleporterV2MessageSent struct {
	EncodedMessage []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeleporterV2MessageSent is a free log retrieval operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) FilterTeleporterV2MessageSent(opts *bind.FilterOpts) (*ZKAdapterTeleporterV2MessageSentIterator, error) {

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "TeleporterV2MessageSent")
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTeleporterV2MessageSentIterator{contract: _ZKAdapter.contract, event: "TeleporterV2MessageSent", logs: logs, sub: sub}, nil
}

// WatchTeleporterV2MessageSent is a free log subscription operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) WatchTeleporterV2MessageSent(opts *bind.WatchOpts, sink chan<- *ZKAdapterTeleporterV2MessageSent) (event.Subscription, error) {

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "TeleporterV2MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterTeleporterV2MessageSent)
				if err := _ZKAdapter.contract.UnpackLog(event, "TeleporterV2MessageSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterV2MessageSent is a log parse operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) ParseTeleporterV2MessageSent(log types.Log) (*ZKAdapterTeleporterV2MessageSent, error) {
	event := new(ZKAdapterTeleporterV2MessageSent)
	if err := _ZKAdapter.contract.UnpackLog(event, "TeleporterV2MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterTransitionedIterator is returned from FilterTransitioned and is used to iterate over the raw logs and unpacked data for Transitioned events raised by the ZKAdapter contract.
type ZKAdapterTransitionedIterator struct {
	Event *ZKAdapterTransitioned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterTransitionedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterTransitioned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterTransitioned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterTransitionedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterTransitionedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterTransitioned represents a Transitioned event raised by the ZKAdapter contract.
type ZKAdapterTransitioned struct {
	PreEpoch  uint64
	PostEpoch uint64
	PreState  ConsensusState
	PostState ConsensusState
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransitioned is a free log retrieval operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) FilterTransitioned(opts *bind.FilterOpts, preEpoch []uint64, postEpoch []uint64) (*ZKAdapterTransitionedIterator, error) {

	var preEpochRule []interface{}
	for _, preEpochItem := range preEpoch {
		preEpochRule = append(preEpochRule, preEpochItem)
	}
	var postEpochRule []interface{}
	for _, postEpochItem := range postEpoch {
		postEpochRule = append(postEpochRule, postEpochItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "Transitioned", preEpochRule, postEpochRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTransitionedIterator{contract: _ZKAdapter.contract, event: "Transitioned", logs: logs, sub: sub}, nil
}

// WatchTransitioned is a free log subscription operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) WatchTransitioned(opts *bind.WatchOpts, sink chan<- *ZKAdapterTransitioned, preEpoch []uint64, postEpoch []uint64) (event.Subscription, error) {

	var preEpochRule []interface{}
	for _, preEpochItem := range preEpoch {
		preEpochRule = append(preEpochRule, preEpochItem)
	}
	var postEpochRule []interface{}
	for _, postEpochItem := range postEpoch {
		postEpochRule = append(postEpochRule, postEpochItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "Transitioned", preEpochRule, postEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterTransitioned)
				if err := _ZKAdapter.contract.UnpackLog(event, "Transitioned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransitioned is a log parse operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) ParseTransitioned(log types.Log) (*ZKAdapterTransitioned, error) {
	event := new(ZKAdapterTransitioned)
	if err := _ZKAdapter.contract.UnpackLog(event, "Transitioned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterZKEventImportedIterator is returned from FilterZKEventImported and is used to iterate over the raw logs and unpacked data for ZKEventImported events raised by the ZKAdapter contract.
type ZKAdapterZKEventImportedIterator struct {
	Event *ZKAdapterZKEventImported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZKAdapterZKEventImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterZKEventImported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZKAdapterZKEventImported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZKAdapterZKEventImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterZKEventImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterZKEventImported represents a ZKEventImported event raised by the ZKAdapter contract.
type ZKAdapterZKEventImported struct {
	SourceChainId *big.Int
	BeaconSlot    *big.Int
	Emitter       common.Address
	ExecutionRoot [32]byte
	LogIndex      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZKEventImported is a free log retrieval operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) FilterZKEventImported(opts *bind.FilterOpts, sourceChainId []*big.Int, beaconSlot []*big.Int, emitter []common.Address) (*ZKAdapterZKEventImportedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var beaconSlotRule []interface{}
	for _, beaconSlotItem := range beaconSlot {
		beaconSlotRule = append(beaconSlotRule, beaconSlotItem)
	}
	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ZKEventImported", sourceChainIdRule, beaconSlotRule, emitterRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterZKEventImportedIterator{contract: _ZKAdapter.contract, event: "ZKEventImported", logs: logs, sub: sub}, nil
}

// WatchZKEventImported is a free log subscription operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) WatchZKEventImported(opts *bind.WatchOpts, sink chan<- *ZKAdapterZKEventImported, sourceChainId []*big.Int, beaconSlot []*big.Int, emitter []common.Address) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var beaconSlotRule []interface{}
	for _, beaconSlotItem := range beaconSlot {
		beaconSlotRule = append(beaconSlotRule, beaconSlotItem)
	}
	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ZKEventImported", sourceChainIdRule, beaconSlotRule, emitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterZKEventImported)
				if err := _ZKAdapter.contract.UnpackLog(event, "ZKEventImported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZKEventImported is a log parse operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) ParseZKEventImported(log types.Log) (*ZKAdapterZKEventImported, error) {
	event := new(ZKAdapterZKEventImported)
	if err := _ZKAdapter.contract.UnpackLog(event, "ZKEventImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
