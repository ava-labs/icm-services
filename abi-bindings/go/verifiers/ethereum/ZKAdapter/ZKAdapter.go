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
	Bin: "0x60e060405234610361576040516143b73881900390601f8201601f19168301906001600160401b0382118483101761036557829160405283398101818103916102208312610361578051916080601f198501126103615760c061007c610063610379565b926100718160208701610398565b845260608501610398565b6020830190815294609f190112610361576040519160c083016001600160401b038111848210176103655760405260a0810151835260c08101516020840190815260e0820151604085019081526101008301516060860190815261012084015191608087019283526101408501519360a08801948552610160860151976101808701519762ffffff89168903610361576101196101a089016103ce565b976101c08101519a61013b6102006101346101e085016103ce565b93016103ce565b908d15610329578d60805280156102e4576101629261015c9160a052610414565b506104ac565b508b515160a05190600581901b6020600160451b0316906001600160401b0316801590820460201417156102d057600c810290808204600c14901517156102d05781018091116102d057421061028b576102429b602080925160018060401b0381511660018060401b0319600154161760015501516002555160018060401b0381511660018060401b031960035416176003550151600455516005555160065551600755516008555160095551600a55600e549060018060a01b03169162ffffff60a01b9060a01b169060018060b81b0319161717600e55600d556103e2565b60c052604051613e39908161051e823960805181818161027f01528181610c230152611326015260a05181818161046b0152610f1c015260c0518181816105a80152610c560152f35b60405162461bcd60e51b815260206004820152601c60248201527f5374617274696e6720737461746520696e2074686520667574757265000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152601460248201527f496e76616c69642067656e657369732074696d650000000000000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b6044820152606490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b60408051919082016001600160401b0381118382101761036557604052565b9190826040910312610361576103ac610379565b825190926001600160401b038216820361036157602091845201516020830152565b51906001600160a01b038216820361036157565b906001821461040e5762aa36a7821461040857506372c559db60e01b5f5260045260245ffd5b60029150565b60019150565b6001600160a01b0381165f9081525f5160206143775f395f51905f52602052604090205460ff166104a7576001600160a01b03165f8181525f5160206143775f395f51905f5260205260408120805460ff191660011790553391907fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775905f5160206143575f395f51905f529080a4600190565b505f90565b6001600160a01b0381165f9081525f5160206143975f395f51905f52602052604090205460ff166104a7576001600160a01b03165f8181525f5160206143975f395f51905f5260205260408120805460ff191660011790553391905f5160206143575f395f51905f528180a460019056fe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146101df57806304ae4327146101da5780631544298e146101d5578063248a9ca3146101d05780632b7ac3f3146101cb5780632f2ff15d146101c6578063304b9071146101c157806336568abe146101bc578063378aa701146101b757806342c6498a146101b25780634b0d0fa8146101ad5780634d3499b5146101a857806354747fd414610185578063547923a3146101a357806365242a0e1461019e5780636fda4c8b1461019957806375b238fc1461019457806391d148541461018f57806397fc007c1461018a578063a217fddf14610185578063bfad283214610180578063c4a6dba11461017b578063cf7923d414610176578063d085070814610171578063d547741f1461016c578063e97eb36214610167578063eb97cd2c14610162578063ebeced0e1461015d5763f1faff0014610158575f80fd5b610bbf565b610b57565b610a18565b6109ee565b6109ae565b61096c565b6107a6565b610729565b610703565b6104ee565b610657565b610606565b6105cc565b61058c565b610572565b610508565b6104d1565b6104aa565b610454565b610408565b61037c565b610361565b61031a565b6102d4565b6102a2565b610268565b610243565b346102355760203660031901126102355760043563ffffffff60e01b811680910361023557602090637965db0b60e01b8114908115610224575b506040519015158152f35b6301ffc9a760e01b1490505f610219565b5f80fd5b5f91031261023557565b34610235575f36600319011261023557602062ffffff600e5460a01c16604051908152f35b34610235575f3660031901126102355760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346102355760203660031901126102355760206102cc6004355f525f602052600160405f20015490565b604051908152f35b34610235575f36600319011261023557600e546040516001600160a01b039091168152602090f35b6001600160a01b0381160361023557565b3590610318826102fc565b565b346102355760403660031901126102355761035f60243560043561033d826102fc565b61035a610355825f525f602052600160405f20015490565b611ad9565b611b11565b005b34610235575f366003190112610235576020604051600c8152f35b346102355760403660031901126102355760043560243561039c816102fc565b336001600160a01b038216036103b55761035f91611b9d565b63334bd91960e11b5f5260045ffd5b9060406020610318936103eb848251602080916001600160401b0381511684520151910152565b0151910190602080916001600160401b0381511684520151910152565b34610235575f3660031901126102355760405161042481610d93565b61042c610e2c565b81526020610438610e2c565b9101526080610445610e6b565b61045260405180926103c4565bf35b34610235575f3660031901126102355760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b6001600160401b0381160361023557565b35906103188261048e565b346102355760203660031901126102355760206102cc6004356104cc8161048e565b610f0c565b34610235575f366003190112610235576020600d54604051908152f35b34610235575f3660031901126102355760206040515f8152f35b346102355760203660031901126102355760043562ffffff81169081810361023557610532611a6a565b600e549162ffffff8360a01c16146105635762ffffff60a01b1990911660a09190911b62ffffff60a01b1617600e55005b63a9cb9e0d60e01b5f5260045ffd5b34610235575f366003190112610235576020604051818152f35b34610235575f36600319011261023557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610235575f3660031901126102355760206040517fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758152f35b3461023557604036600319011261023557602060ff61064b60243560043561062d826102fc565b5f525f845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b3461023557602036600319011261023557600435610674816102fc565b61067c611a6a565b6001600160a01b03811680156106be57600e546001600160a01b03161461056357600e80546001600160a01b0319166001600160a01b03909216919091179055005b60405162461bcd60e51b815260206004820152601860248201527f496e76616c6964207665726966696572206164647265737300000000000000006044820152606490fd5b34610235575f3660031901126102355760206001600160401b03600c5416604051908152f35b34610235576020366003190112610235576004356001600160401b0381116102355736602382011215610235578060040135906001600160401b0382116102355736602483830101116102355761035f91602461079292019061078a611a6a565b810190610fb8565b6001600160401b0360408201511690611c40565b34610235576020366003190112610235576004356001600160401b038111610235578060040160406003198336030112610235576107e761078a8280611004565b916108016107fd84516107f8610e6b565b61248c565b1590565b61095d576108186107fd8451602086015190612500565b61094e5760205f6108298480611004565b9061083960405180938193611d23565b039060025afa15610949576108775f5192602461086f610863610863600e5460018060a01b031690565b6001600160a01b031690565b930190611004565b9091600d5491813b15610235575f936108a66040519687958694859463ab750e7560e01b865260048601611d30565b03915afa80156109495761092f575b5060408101906108cc82516001600160401b031690565b600c546001600160401b0316906001600160401b0382166001600160401b038216111561090e5761035f8361090886516001600160401b031690565b90611c40565b63f907012760e01b5f526001600160401b039081166004521660245260445ffd5b8061093d5f61094393610dce565b80610239565b5f6108b5565b611a0a565b632d776d6b60e11b5f5260045ffd5b634247f10960e11b5f5260045ffd5b34610235576020366003190112610235576001600160401b036004356109918161048e565b165f52600b6020526040805f205481519080825215156020820152f35b346102355760403660031901126102355761035f6024356004356109d1826102fc565b6109e9610355825f525f602052600160405f20015490565b611b9d565b3461023557602036600319011261023557600435610a0a611a6a565b600d54811461056357600d55005b34610235576020366003190112610235576004356001600160401b038111610235576101206003198236030112610235576044810135610a57816102fc565b6001600160a01b03163303610ac257610abd610aa3610ab17f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b69360405192839160040160208301611189565b03601f198101835282610dce565b60405191829182610b43565b0390a1005b60405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b6044820152606490fd5b5f5b838110610b0e5750505f910152565b8181015183820152602001610aff565b90602091610b3781518092818552858086019101610afd565b601f01601f1916010190565b906020610b54928181520190610b1e565b90565b34610235576040366003190112610235576004356001600160401b038111610235576101406003198236030112610235576024356001600160401b0381116102355760c0600319823603011261023557610bbb91610ab1916004019060040161127c565b0390f35b34610235576020366003190112610235576004356001600160401b038111610235578060040160806003198336030112610235575f82610c7e610c4e6024610c18610c106064610cf7990188611004565b8101906115a0565b93610c4860448201357f0000000000000000000000000000000000000000000000000000000000000000146116f9565b01611738565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614611749565b60208101805160800151909190610ca8906001600160a01b03166001600160a01b03163014611783565b610cd87f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b660a084510151146117bd565b5190516040516375f6768760e11b815294859283929060048401611937565b038183305af1801561094957610d25610d3391610d51945f91610d5d575b50602080825183010191016117f5565b602081519101209180611a15565b604051610d4881610aa3602082019485611189565b51902014611a2b565b60405160018152602090f35b610d7991503d805f833e610d718183610dce565b8101906117f5565b5f610d15565b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b03821117610dae57604052565b610d7f565b606081019081106001600160401b03821117610dae57604052565b90601f801991011681019081106001600160401b03821117610dae57604052565b6040519061031860c083610dce565b6040519061031861014083610dce565b60405190610318604083610dce565b60405190610318608083610dce565b60405190610e3982610d93565b5f6020838281520152565b90604051610e5181610d93565b6020600182946001600160401b0381541684520154910152565b60405190610e7882610d93565b81610e836001610e44565b81526020610e916003610e44565b910152565b634e487b7160e01b5f52601160045260245ffd5b908160051b9180830460201490151715610ec057565b610e96565b90600c820291808304600c1490151715610ec057565b908160011b9180830460021490151715610ec057565b9060018201809211610ec057565b91908201809211610ec057565b610f47610f426001600160401b037f00000000000000000000000000000000000000000000000000000000000000009316610eaa565b610ec5565b8101809111610ec05790565b919082604091031261023557604051610f6b81610d93565b60208082948035610f7b8161048e565b84520135910152565b9190608083820312610235576020610e9160405192610fa284610d93565b60408496610fb08382610f53565b865201610f53565b610120818303126102355761010090610fec60405193610fd785610db3565b610fe18184610f84565b855260808301610f84565b60208401520135610ffc8161048e565b604082015290565b903590601e198136030182121561023557018035906001600160401b0382116102355760200191813603831361023557565b35610b54816102fc565b9035601e19823603018112156102355701602081359101916001600160401b038211610235578160051b3603831361023557565b916020908281520191905f5b81811061108d5750505090565b90919260208060019286356110a1816102fc565b848060a01b031681520194019101919091611080565b9035601e19823603018112156102355701602081359101916001600160401b038211610235578160061b3603831361023557565b916020908281520191905f5b8181106111045750505090565b90919260408060019286358152602087013561111f816102fc565b848060a01b0316602082015201940191019190916110f7565b9035601e19823603018112156102355701602081359101916001600160401b03821161023557813603831361023557565b908060209392818452848401375f828201840152601f01601f1916010190565b90610b549160208152813560208201526111b86111a86020840161030d565b6001600160a01b03166040830152565b6111d76111c76040840161030d565b6001600160a01b03166060830152565b606082013560808201526112006111f06080840161030d565b6001600160a01b031660a0830152565b60a082013560c082015261125f61125361123361122060c0860186611040565b61012060e0870152610140860191611074565b61124060e08601866110b7565b858303601f1901610100870152906110eb565b92610100810190611138565b91610120601f1982860301910152611169565b35610b548161048e565b906001600160401b0382356112908161048e565b165f52600b60205260405f2054801561134a57826112ad91611f91565b6112bc81610100840135612327565b916112c960208201611272565b917f17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb16001600160401b03606060c061130360808601611036565b6040805192909701358252919094013560208501526001600160a01b03169416927f000000000000000000000000000000000000000000000000000000000000000092a490565b60405162461bcd60e51b815260206004820152601860248201527f416e63686f7220736c6f7420697320756e646566696e656400000000000000006044820152606490fd5b6001600160401b038111610dae5760051b60200190565b9291906113b28161138f565b936113c06040519586610dce565b602085838152019160051b810192831161023557905b8282106113e257505050565b81358152602091820191016113d6565b9080601f8301121561023557816020610b54933591016113a6565b6001600160401b038111610dae57601f01601f191660200190565b9291926114348261140d565b916114426040519384610dce565b829481845281830111610235578281602093845f960137010152565b9080601f8301121561023557816020610b5493359101611428565b9291906114858161138f565b936114936040519586610dce565b602085838152019160051b8101918383116102355781905b8382106114b9575050505050565b81356001600160401b038111610235576020916114d9878493870161145e565b8152019101906114ab565b9080601f8301121561023557816020610b5493359101611479565b919060c08382031261023557611513610def565b9280356001600160401b03811161023557826115309183016114e4565b845260208101356001600160401b038111610235578261155191830161145e565b602085015260408101356001600160401b0381116102355760a09261157791830161145e565b6040850152606081013560608501526115926080820161030d565b6080850152013560a0830152565b602081830312610235578035906001600160401b038211610235570160408183031261023557604051916115d383610d93565b81356001600160401b03811161023557820161014081830312610235576115f8610dfe565b906116028161049f565b82526116106020820161049f565b60208301526040810135604083015260608101356001600160401b038111610235578361163e9183016113f2565b60608301526080810135608083015260a08101356001600160401b038111610235578361166c9183016113f2565b60a083015260c081013560c083015260e08101356001600160401b038111610235578361169a9183016113f2565b60e0830152610100810135610100830152610120810135906001600160401b038211610235576116cc918491016113f2565b610120820152835260208201356001600160401b038111610235576116f192016114ff565b602082015290565b1561170057565b60405162461bcd60e51b815260206004820152601060248201526f3130b21039b7bab931b29031b430b4b760811b6044820152606490fd5b3563ffffffff811681036102355790565b1561175057565b60405162461bcd60e51b815260206004820152600b60248201526a626164206e6574776f726b60a81b6044820152606490fd5b1561178a57565b60405162461bcd60e51b815260206004820152600b60248201526a3130b21032b6b4ba3a32b960a91b6044820152606490fd5b156117c457565b60405162461bcd60e51b815260206004820152600960248201526862616420746f70696360b81b6044820152606490fd5b602081830312610235578051906001600160401b038211610235570181601f820112156102355780516118278161140d565b926118356040519485610dce565b8184526020828401011161023557610b549160208085019101610afd565b90602080835192838152019201905f5b8181106118705750505090565b8251845260209384019390920191600101611863565b9060c0810182519060c08352815180915260e0830190602060e08260051b8601019301915f905b82821061190c575050505060a06118e66118d4829360208701518682036020880152610b1e565b60408601518582036040870152610b1e565b606080860151908501526080808601516001600160a01b03169085015293015191015290565b9091929360208061192960019360df198a82030186528851610b1e565b9601920192019092916118ad565b90916119fc610b54936040845261195a6040850182516001600160401b03169052565b60208101516001600160401b03166060850152604081015160808501526101206119db6119bb61199b606085015161014060a08a0152610180890190611853565b608085015160c089015260a0850151888203603f190160e08a0152611853565b60c084015161010088015260e0840151878203603f190184890152611853565b610100830151610140870152910151848203603f1901610160860152611853565b916020818403910152611886565b6040513d5f823e3d90fd5b90359061011e1981360301821215610235570190565b15611a3257565b60405162461bcd60e51b815260206004820152601060248201526f0e0c2f2d8dec2c840dad2e6dac2e8c6d60831b6044820152606490fd5b335f9081527f7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ec602052604090205460ff1615611aa257565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b5f8181526020818152604080832033845290915290205460ff1615611afb5750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff16611b97575f818152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615611b97575f818152602081815260408083206001600160a01b03861684529091529020805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b916080610318929493611c39816101008101976103c4565b01906103c4565b9060208061031893818101908151838151611c7b6001600160401b038251166001600160401b03166001600160401b03196001541617600155565b01516002557f955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d06001600160401b038583015193611cd1828651166001600160401b03166001600160401b03196003541617600355565b868501516004555192611d15611d03611cf5898701516001600160401b0390511690565b96516001600160401b031690565b91836040519485941697169583611c21565b0390a351015101519061243d565b908092918237015f815290565b90611d4a9060409396959496606084526060840191611169565b9460208201520152565b15611d5b57565b60405162461bcd60e51b815260206004820152602960248201527f54617267657420736c6f74206d757374206265206c657373207468616e20616e60448201526818da1bdc881cdb1bdd60ba1b6064820152608490fd5b906001600160401b03809116911603906001600160401b038211610ec057565b15611dd957565b60405162461bcd60e51b815260206004820152601660248201527515185c99d95d081cdb1bdd081a5cc81d1bdbc81bdb1960521b6044820152606490fd5b903590601e198136030182121561023557018035906001600160401b03821161023557602001918160051b3603831361023557565b15611e5357565b60405162461bcd60e51b815260206004820152601960248201527f496e76616c696420616e63686f7220737461746520726f6f74000000000000006044820152606490fd5b8115611ea2570690565b634e487b7160e01b5f52601260045260245ffd5b15611ebd57565b606460405162461bcd60e51b815260206004820152602060248201527f496e76616c69642074617267657420626561636f6e20737461746520726f6f746044820152fd5b15611f0857565b60405162461bcd60e51b815260206004820152601c60248201527f496e76616c696420657865637574696f6e20726f6f742070726f6f66000000006044820152606490fd5b15611f5457565b60405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081c9958d95a5c1d1cc81c9bdbdd605a1b6044820152606490fd5b6120bc906103189261207c61207761205961204d6001600160401b038061204760208801611fe2611fc182611272565b84611fda611fce8d611272565b6001600160401b031690565b911610611d54565b61200e612000611ff18b611272565b611ffa84611272565b90611db2565b600a54958691161115611dd2565b61204261203d60408b01359b8c8c612037612030600554926060810190611e17565b36916113a6565b91612570565b611e4c565b611272565b16611e98565b6007546008541b610eff565b9360808401359461207061203060a0870187611e17565b9086612570565b611eb6565b6120a061209b60c08301359360065461207061203060e0870187611e17565b611f01565b6101008101359061203761203060095492610120810190611e17565b611f4d565b604080519091906120d28382610dce565b6001815291601f1901825f5b8281106120ea57505050565b8060606020809385010152016120de565b634e487b7160e01b5f52603260045260245ffd5b80511561211c5760200190565b6120fb565b80516001101561211c5760400190565b80516002101561211c5760600190565b80516003101561211c5760800190565b80516010101561211c576102200190565b805182101561211c5760209160051b010190565b1561217d57565b60405162461bcd60e51b815260206004820152602a60248201527f496e76616c6964206e756d626572206f6620726573756c747320696e2072656360448201526932b4b83a10383937b7b360b11b6064820152608490fd5b156121dc57565b60405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103932b1b2b4b83a10383937b7b360591b6044820152606490fd5b1561222057565b60405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840d8dece40d2dcc8caf607b1b6044820152606490fd5b1561226057565b60405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103637b39032b6b4ba3a32b960691b6044820152606490fd5b156122a257565b60405162461bcd60e51b81526020600482015260116024820152704c6f6720686173206e6f20746f7069637360781b6044820152606490fd5b156122e257565b60405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206576656e74207369676e61747572650000000000000000006044820152606490fd5b906124386123df6123c36123be60206123b661238c6040986123476120c1565b9061235e612357868c018c611004565b3691611428565b6123678361210f565b526123718261210f565b5061238661237f8b80611e17565b3691611479565b9061270a565b6123996001825114612176565b6123b1836123a68361210f565b5101515115156121d5565b61210f565b5101516129d3565b612b3a565b6080606085013591016123d98151518310612219565b51612162565b518051909290612412906001600160a01b031661240161086360808501611036565b6001600160a01b0390911614612259565b60a061242e60208501612428815151151561229b565b5161210f565b51910135146122db565b015190565b6001600160401b0316806001600160401b0319600c541617600c55805f52600b6020528160405f20557f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d5f80a3565b9061249a8251825190612c7b565b91826124a557505090565b610b54925060208091015191015190612c7b565b5f19810191908211610ec057565b6020039060208211610ec057565b60bf19810191908211610ec057565b607f19810191908211610ec057565b91908203918211610ec057565b60209081015151910151516001600160401b03908116911680821115611b975761252982610f0c565b42811161256957420391428311610ec057600e5460a01c62ffffff168093116125695761255f611fce61256493610f4293611db2565b610eaa565b111590565b5050505f90565b915f905b60018303612583575050501490565b9091925f60209160016125968760011690565b1482146125e6576125b46125c0916125ae8787612162565b51612cad565b60405191828092612ccc565b039060025afa15610949576125e06125da5f519460011c90565b926126b7565b90612574565b6125b46125c0916125f78787612162565b5190612cad565b906126088261138f565b6126156040519182610dce565b8281528092612626601f199161138f565b01905f5b82811061263657505050565b60209060405161264581610d93565b606081526060838201528282850101520161262a565b906126658261138f565b6126726040519182610dce565b8281528092612683601f199161138f565b01905f5b82811061269357505050565b6020906040516126a281610d93565b5f815260608382015282828501015201612687565b5f198114610ec05760010190565b60405190608082018281106001600160401b03821117610dae57604052606080835f81525f60208201525f60408201520152565b90601081101561211c5760051b0190565b9261271583516125fe565b61271f835161265b565b945f5b8451811015612777578061273860019287612162565b516020815191012061274a8288612162565b51612753610e0e565b9182526020820152612765828a612162565b526127708189612162565b5001612722565b5091939092505f5b81518110156129cb576127928183612162565b5161279d8286612162565b51526127a98183612162565b516127b2610e0e565b9081525f6020820152906127ce6127c98589612ce3565b612e06565b91600192835b6127e5575b5050600191500161277f565b6127ed6126c5565b936127f782612f0a565b1561288e57506001935061283061281061284892613488565b9260208101612820815160011690565b612876575160011c905b5161340f565b612838610e0e565b9081525f602082015282516135a0565b612853575b806127d9565b60206128619101518861323e565b602061286d8388612162565b5101525f61284d565b516128889060011c610ef1565b610ef1565b9061282a565b91929361289e60c0830151151590565b1561293057506128ad906132ae565b916128b9835182613391565b15612928576129196127c96020836129006128f96128e78461291f9801516128e18b516133a8565b90610eff565b92516128f38460011c90565b9061340f565b9160011690565b612908610e0e565b918252828201529501515b8b61323e565b916126b7565b929190836127d4565b8392506127d9565b60e0820151929392156129b3575061294790612fe6565b9161295182613182565b15612969575050518051600192919015158314612848565b9091602061298291015161297c8461319c565b906126f9565b5191612992836001905115151490565b15612928576129196127c960206129ab61291f9461320c565b950151612913565b9051919291612928576129196127c961291f92612913565b509193505050565b6129db610e2c565b506020815191604051926129ee84610d93565b835201602082015290565b6040519060a082018281106001600160401b03821117610dae5760405260606080835f81528260208201525f604082015282808201520152565b15612a3a57565b60405162461bcd60e51b8152602060048201526013602482015272556e737570706f72746564207478207479706560681b6044820152606490fd5b15612a7c57565b60405162461bcd60e51b815260206004820152602260248201527f496e76616c696420726563656970743a2065787065637465642034206669656c604482015261647360f01b6064820152608490fd5b60405190612ad982610db3565b60606040835f81528260208201520152565b90612af58261138f565b612b026040519182610dce565b8281528092612b13601f199161138f565b01905f5b828110612b2357505050565b602090612b2e612acc565b82828501015201612b17565b90612bca612bc4612b62612b4c6129f9565b94612b596107fd826135cf565b612c21576135f3565b612b6f6004825114612a75565b612b81612b7b8261210f565b516136b8565b6020860152612bae612b9e611fce612b9884612121565b516136e2565b6001600160401b03166040870152565b612bba612b7b82612131565b6060860152612141565b516135f3565b805190612bd682612aeb565b90608085019182525f5b838110612bed5750505050565b80612c1a81612c07612c0160019587612162565b516137e4565b865190612c148383612162565b52612162565b5001612be0565b612c5b612c55602083015192612c4f84515f1a600181101580612c70575b612c4890612a33565b60ff168a52565b516124b9565b91610ef1565b612c63610e0e565b91825260208201526135f3565b506004811115612c3f565b906001600160401b038251166001600160401b03825116149182612c9e57505090565b60209192508101519101511490565b9190604051926020840152604083015260408252610318606083610dce565b90612cdf60209282815194859201610afd565b0190565b5f5b8151811015612d1b5782612cf98284612162565b515114612d0857600101612ce5565b60209250612d1591612162565b51015190565b60405162461bcd60e51b8152602060048201526011602482015270496e636f6d706c6574652050726f6f662160781b6044820152606490fd5b60405190612d6182610d93565b5f602083606081520152565b6040519061014082018281106001600160401b03821117610dae57604052815f81525f60208201525f60408201525f60608201525f60808201525f60a08201525f60c08201525f60e08201525f610100820152610120610e91612d54565b60405190612dda604083610dce565b602082527f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b4216020830152565b612e0e612d6d565b50612e17612d6d565b90612e20610e0e565b908082525f6020830152612e3b612e35612dcb565b82613893565b612f0157612e4b612e50916129d3565b6135f3565b805180612e61575050600182525090565b60028103612eb75750612e78612b7b60209261210f565b015160fc1c60028114908115612eac575b5015612ea057600160208301525b61012082015290565b600160c0830152612e97565b60039150145f612e89565b601114159050612ecd57600160e0830152612e97565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206461746160a01b6044820152606490fd5b50506001815290565b60208101511590811591612f1c575090565b604091500151151590565b60405190612f3482610d93565b815f81526020610e916126c5565b60405190610200612f538184610dce565b825f5b828110612f6257505050565b602090612f6d612f27565b8184015201612f56565b60405190612f8482610d93565b81612f8d612f27565b81526020610e91612f42565b60405190612fa8602083610dce565b5f808352366020840137565b90612fbe8261140d565b612fcb6040519182610dce565b8281528092612fdc601f199161140d565b0190602036910137565b612fee612f77565b50613009612e4b610120613000612f77565b930151516129d3565b613011612f42565b905f5b601081106130ba575061303161302c612b7b83612151565b61390e565b15613077575061303f612f99565b613047610e1d565b905f82525f60208301525f60408301526060820152613064610e0e565b905f825260208201528252602082015290565b612b7b61308391612151565b61308b610e1d565b5f8082526020820152600160408201529060608201526130a9610e0e565b6001815290602082015282526116f1565b806130ca612b7b60019385612162565b805160201461312757506130dc612f99565b6130e4610e1d565b905f82525f60208301525f60408301526060820152613101610e0e565b905f8252602082015261311482866126f9565b5261311f81856126f9565b505b01613014565b613130906138ad565b613138612f99565b613140610e1d565b600181529160208301525f6040830152606082015261315d610e0e565b6001815290602082015261317182866126f9565b5261317c81856126f9565b50613121565b6107fd906133a8565b90815181101561211c570160200190565b600f60ff91600160025f856131bd60208601958651965196861c905161318b565b5160f81c169308036131ce57161690565b60041c161690565b906020820191825192828401809411610ec05751908282018211610ec0576002600f9360ff6131bd6001948298861c905161318b565b613214612d54565b506020815191015160018101809111610ec0576040519161323483610d93565b8252602082015290565b60408201511561325057506060015190565b815161326c575050604051613266602082610dce565b5f815290565b6020610b5492015190612ce3565b6040519061328782610d93565b81613290612d54565b81526020610e916126c5565b60ff60019116019060ff8211610ec057565b6132b661327a565b506133646132d4612e4b6101206132cb61327a565b940151516129d3565b6132e0612b7b82612121565b9061334c61332b612b7b600161332461331e6133126133046123b1612b7b8a61210f565b516001600160f81b03191690565b60041c600f60f81b1690565b60f81c90565b169361210f565b61334661334061333a8561329c565b60011690565b60ff1690565b90613965565b90613355610e0e565b918252602082015283526138ad565b61336c612f99565b613374610e1d565b600181529160208301525f60408301526060820152602082015290565b61339e826133a492613a01565b916133a8565b1490565b805151908160011b9180830460021490151715610ec057602001518103908111610ec05790565b156133d657565b60405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606490fd5b9081516134288261342081846124f3565b9210156133cf565b80613440575050506040515f81526020810160405290565b60405192601f821692831560051b80858701019484860193010101905b8084106134755750508252601f01601f191660405290565b909260208091855181520193019061345d565b61349061327a565b506134a2612e4b61012061300061327a565b80516001101561211c576134c5612b7b6134bf60408401516136b8565b9261210f565b80519081600111610235575f198201918211610ec05760218101906020018110610ec0576134f282612fb4565b916020830191905b6020811015613570578061355157505f19905b5182518216911916179052613520610e0e565b9081525f60208201528252613533610e1d565b5f808252602082015260016040820152906060820152602082015290565b61356561356061356a926124c7565b613cb0565b6124b9565b9061350d565b90918251815260208101809111610ec0579160208101809111610ec05790601f198101908111156134fa57610e96565b906135aa826133a8565b6135b3826133a8565b1491826135bf57505090565b816133a492935061339e91613a01565b8051156135ee57602060c0910151515f1a106135ea57600190565b5f90565b505f90565b6135fc816135cf565b156102355761360a81613b3f565b6136138161138f565b916136216040519384610dce565b818352601f196136308361138f565b015f5b8181106136a1575050602061364d9101516128e181613ba0565b5f905b82821061365d5750505090565b6136998161366c600193613c06565b90613675610e0e565b8281528160208201526136888689612162565b526136938588612162565b50610eff565b910190613650565b6020906136ac610e2c565b82828801015201613633565b805115610235576136cb610b5491613c84565b6136d781939293612fb4565b928360200190613cbf565b80518015159081613718575b5015610235576136fd90613c84565b9051906020811061370c575090565b6020036101000a900490565b6021915011155f6136ee565b1561372b57565b60405162461bcd60e51b815260206004820152601e60248201527f496e76616c6964206c6f673a2065787065637465642033206669656c647300006044820152606490fd5b1561377757565b60405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206c6f673a20746f6f206d616e7920746f70696373000000006044820152606490fd5b906137c68261138f565b6137d36040519182610dce565b8281528092612fdc601f199161138f565b90613841612bc46137fc6137f6612acc565b946135f3565b6138096003825114613724565b61382b61381e6138188361210f565b51613d2d565b6001600160a01b03168652565b613837612b7b82612131565b6040860152612121565b8051906138516004831115613770565b61385a826137bc565b90602085019182525f5b8381106138715750505050565b80613881612b9860019385612162565b61388c828651612162565b5201613864565b908151908051820361256957816020809201209201201490565b60208151106138bd576020015190565b60405162461bcd60e51b8152602060048201526024808201527f42797465733a3a20746f427974657333323a206461746120697320746f20736860448201526337b93a1760e11b6064820152608490fd5b805115908115918261391f57505090565b90915061211c576020810180516001600160f81b031916600360fe1b1491821561394857505090565b909150511561211c57516001600160f81b031916600160ff1b1490565b9081519182821161023557818303928311610ec057602080928201019182910111610ec05761399382612fb4565b916020830191905b60208110156139d157806139bc57505f19905b518251821691191617905290565b6135656135606139cb926124c7565b906139ae565b90918251815260208101809111610ec0579160208101809111610ec05790601f1981019081111561399b57610e96565b6020810190613a11825160011690565b916020840190613a22825160011690565b8403613ada5751613a389060011c915160011c90565b915f93613a68575b613a6392613a56610b549693613a5d935161340f565b925161340f565b90613d5e565b610eff565b925092600f60f81b613a7e61330486865161318b565b16613aa3600f60f81b613a9561330486865161318b565b166001600160f81b03191690565b03613ad257613a5d613a6392613a56613ac16125da610b54986126b7565b916001969350939650509250613a40565b505050505f90565b5050919050613afa613aeb836133a8565b613af4836133a8565b90613d4c565b905f925b828410613b0c575b50505090565b909192613b1981836131d6565b613b2382856131d6565b03613b3957613b31906126b7565b929190613afe565b92613b06565b8051156135ee575f9060208101908151613b5881613ba0565b8101809111610ec057915190518101809111610ec05791905b828110613b7e5750905090565b613b8781613c06565b8101809111610ec057613b9a90916126b7565b90613b71565b515f1a6080811015613bb157505f90565b60b881108015613bf0575b15613bc75750600190565b60c0811015613be157610b54906128839060b75b906124f3565b610b54906128839060f7613bdb565b5060c08110158015613bbc575060f88110613bbc565b80515f1a906080821015613c1b575050600190565b60b8821015613c315750612883610b54916124e4565b60c0821015613c545760010151602082900360b7016101000a90040160b5190190565b60f8821015613c6a5750612883610b54916124d5565b60010151602082900360f7016101000a90040160f5190190565b906020820191613c948351613ba0565b925190838201809211610ec05751928303928311610ec0579190565b601f8111610ec0576101000a90565b90918015613d28575b6020811015613cf85780613cdb57505050565b613565613560613cea926124c7565b905182518216911916179052565b91908051825260208101809111610ec0579060208101809111610ec05791601f19810190811115613cc857610e96565b505050565b6015815103610235576001600160a01b0390613d48906136e2565b1690565b9080821015613d59575090565b905090565b613d6b8151835190613d4c565b915f915b838310613d82575050610b549150610edb565b90916001600160f81b0319613d97828461318b565b51166001600160f81b0319613dac838661318b565b511603613dc357613dbc906126b7565b9190613d6f565b80935060011b9183830460021484151715610ec057613df861330485613df2613304610b54986128e19761318b565b9361318b565b6001600160f81b031980831690821603613e13575050600290565b60f0809160f81c169160f81c16145f146135ea5760019056fea164736f6c634300081e000a2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ecad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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
