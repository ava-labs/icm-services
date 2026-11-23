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

// ConsensusData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusData struct {
	PublicValues []byte
	Proof        []byte
}

// ConsensusState is an auto generated low-level Go binding around an user-defined struct.
type ConsensusState struct {
	Head              uint64
	Header            [32]byte
	SyncCommitteeHash [32]byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"head\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"header\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.State\",\"name\":\"startingState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gIndexBlockStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexExecRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexBaseStateRoots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexReceiptsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsVectorSize\",\"type\":\"uint256\"}],\"internalType\":\"structExecution.BeaconConfig\",\"name\":\"beaconConfig_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"permissibleTimespan_\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"lightClientVkey_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPreState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"given\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"}],\"name\":\"NextSyncCommitteeMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"NonCheckpointSlot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastConfirmedSlot\",\"type\":\"uint64\"}],\"name\":\"NonIncreasingSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissibleTimespanLapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedSourceChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedBeaconBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newVkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldVkey\",\"type\":\"bytes32\"}],\"name\":\"LightClientVkeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"permissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"PermissibleTimespanUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"SyncCommitteeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedMessage\",\"type\":\"bytes\"}],\"name\":\"TeleporterV2MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"preHead\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"postHead\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"head\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"header\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"preState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"head\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"header\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"postState\",\"type\":\"tuple\"}],\"name\":\"Transitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beaconSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"ZKEventImported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOTS_PER_PERIOD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOT_PER_EPOCH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDEFINED_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getBeaconBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"head\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"header\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastConfirmedSlot\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightClientVkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"}],\"name\":\"manualTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissibleTimespan\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"slotTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceNetworkId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"syncCommittees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structConsensusData\",\"name\":\"consensus\",\"type\":\"tuple\"}],\"name\":\"transition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newLightClientVkey\",\"type\":\"bytes32\"}],\"name\":\"updateLightClientVkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newPermissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"updatePermissibleTimespan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"anchorSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"anchorBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"anchorBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetExecutionHeaderRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetExecutionHeaderProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetReceiptsProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structExecution.Proof\",\"name\":\"execProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"expectedEmitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"expectedTopic0\",\"type\":\"bytes32\"}],\"internalType\":\"structReceipt.Proof\",\"name\":\"logProof\",\"type\":\"tuple\"}],\"name\":\"verifyLogAndExtract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051614fee380380614fee83398101604081905261002e916104c4565b888888888888888888885f0361007e5760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b60448201526064015b60405180910390fd5b60808990525f8690036100d35760405162461bcd60e51b815260206004820152601460248201527f496e76616c69642067656e657369732074696d650000000000000000000000006044820152606401610075565b60a08690526101027fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177583610263565b5061010d5f82610263565b508751429061011b9061030c565b11156101695760405162461bcd60e51b815260206004820152601c60248201527f5374617274696e6720737461746520696e2074686520667574757265000000006044820152606401610075565b8751600180546001600160401b0319166001600160401b03831617905560208901516002556040890151600381905590600b905f906101a79061032e565b81526020019081526020015f20819055506101cf885f0151896020015161034a60201b60201c565b50508451600455602085015160055560408501516006556060850151600755608085015160085560a090940151600955600e80546001600160b81b031916600160a01b62ffffff94909416939093026001600160a01b031916929092176001600160a01b039190911617905550600d555061024b9050896103a4565b63ffffffff1660c0525061060f975050505050505050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610303575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556102bb3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610306565b505f5b92915050565b5f610321600c6001600160401b0384166105ab565b60a05161030691906105c2565b5f61033b612000836105d5565b6001600160401b031692915050565b600c80546001600160401b0319166001600160401b0384169081179091555f818152600a6020526040808220849055518392917f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d91a35050565b5f816001036103b557506001919050565b8162aa36a7036103c757506002919050565b6040516372c559db60e01b815260048101839052602401610075565b604051606081016001600160401b038111828210171561041157634e487b7160e01b5f52604160045260245ffd5b60405290565b5f60c08284031215610427575f5ffd5b60405160c081016001600160401b038111828210171561045557634e487b7160e01b5f52604160045260245ffd5b604090815283518252602080850151908301528381015190820152606080840151908201526080808401519082015260a0928301519281019290925250919050565b805162ffffff811681146104a9575f5ffd5b919050565b80516001600160a01b03811681146104a9575f5ffd5b5f5f5f5f5f5f5f5f5f898b036102008112156104de575f5ffd5b8a5199506060601f19820112156104f3575f5ffd5b506104fc6103e3565b60208b01516001600160401b0381168114610515575f5ffd5b815260408b810151602083015260608c015190820152975061053a8b60808c01610417565b6101408b015190975095506105526101608b01610497565b94506105616101808b016104ae565b6101a08b015190945092506105796101c08b016104ae565b91506105886101e08b016104ae565b90509295985092959850929598565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761030657610306610597565b8082018082111561030657610306610597565b5f6001600160401b038316806105f957634e487b7160e01b5f52601260045260245ffd5b6001600160401b03929092169190910492915050565b60805160a05160c0516149996106555f395f81816103840152610b3001525f8181610328015261067301525f818161023701528181610a340152610aca01526149995ff3fe608060405234801561000f575f5ffd5b50600436106101dc575f3560e01c80636b3ba09011610109578063bfad28321161009e578063d547741f1161006e578063d547741f146104a1578063eb97cd2c146104b4578063ebeced0e146104c7578063f1faff00146104e7575f5ffd5b8063bfad283214610427578063c4a6dba11461043a578063cf7923d41461044d578063d085070814610460575f5ffd5b806391d14854116100d957806391d14854146103ee578063978c0ef51461040157806397fc007c14610414578063a217fddf1461034a575f5ffd5b80636b3ba0901461036c5780636fda4c8b1461037f57806375b238fc146103bb5780637640d76b146103cf575f5ffd5b8063304b90711161017f57806342c6498a1161014f57806342c6498a1461032357806354747fd41461034a578063547923a31461035157806365242a0e14610364575f5ffd5b8063304b9071146102ea57806332e0c9a7146102f257806336568abe146102fb578063378aa7011461030e575f5ffd5b80632073ee70116101ba5780632073ee7014610267578063248a9ca3146102885780632b7ac3f3146102aa5780632f2ff15d146102d5575f5ffd5b806301ffc9a7146101e057806304ae4327146102085780631544298e14610232575b5f5ffd5b6101f36101ee366004613791565b6104fa565b60405190151581526020015b60405180910390f35b600e5461021e90600160a01b900462ffffff1681565b60405162ffffff90911681526020016101ff565b6102597f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101ff565b61027061200081565b6040516001600160401b0390911681526020016101ff565b6102596102963660046137b8565b5f9081526020819052604090206001015490565b600e546102bd906001600160a01b031681565b6040516001600160a01b0390911681526020016101ff565b6102e86102e33660046137ea565b610530565b005b610270600c81565b610259600d5481565b6102e86103093660046137ea565b61055a565b610316610592565b6040516101ff9190613814565b6102597f000000000000000000000000000000000000000000000000000000000000000081565b6102595f81565b6102e861035f36600461383e565b6105ea565b610270602081565b61025961037a366004613876565b610658565b6103a67f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101ff565b6102595f5160206149445f395f51905f5281565b6102596103dd3660046137b8565b600b6020525f908152604090205481565b6101f36103fc3660046137ea565b610697565b6102e861040f3660046137b8565b6106bf565b6102e861042236600461388f565b6106fe565b600c54610270906001600160401b031681565b6102e86104483660046138a8565b6107c2565b6102e861045b366004613914565b610800565b61048c61046e366004613876565b6001600160401b03165f908152600a60205260409020549081151590565b604080519283529015156020830152016101ff565b6102e86104af3660046137ea565b610886565b6102e86104c236600461394a565b6108aa565b6104da6104d5366004613981565b610967565b6040516101ff9190613a3c565b6101f36104f5366004613a4e565b610aa5565b5f6001600160e01b03198216637965db0b60e01b148061052a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f8281526020819052604090206001015461054a81610d5c565b6105548383610d69565b50505050565b6001600160a01b03811633146105835760405163334bd91960e11b815260040160405180910390fd5b61058d8282610df8565b505050565b6105bc60405180606001604052805f6001600160401b031681526020015f81526020015f81525090565b50604080516060810182526001546001600160401b0316815260025460208201526003549181019190915290565b5f5160206149445f395f51905f5261060181610d5c565b600e5462ffffff600160a01b9091048116908316036106335760405163a9cb9e0d60e01b815260040160405180910390fd5b50600e805462ffffff909216600160a01b0262ffffff60a01b19909216919091179055565b5f61066d600c6001600160401b038416613a98565b61052a907f0000000000000000000000000000000000000000000000000000000000000000613aaf565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f5160206149445f395f51905f526106d681610d5c565b600d5482036106f85760405163a9cb9e0d60e01b815260040160405180910390fd5b50600d55565b5f5160206149445f395f51905f5261071581610d5c565b6001600160a01b0382166107705760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076657269666965722061646472657373000000000000000060448201526064015b60405180910390fd5b600e546001600160a01b039081169083160361079f5760405163a9cb9e0d60e01b815260040160405180910390fd5b50600e80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5160206149445f395f51905f526107d981610d5c565b5f6107e683850185613c4d565b9050610554816107f98360600151610e61565b6001610e9a565b5f61080b8280613d15565b8101906108189190613c4d565b90505f6108288260600151610e61565b9050610835828285611072565b600c546001600160401b039081169082161161087b57600c5460405163f907012760e01b81526001600160401b0380841660048301529091166024820152604401610767565b61058d82825f610e9a565b5f828152602081905260409020600101546108a081610d5c565b6105548383610df8565b6108ba606082016040830161388f565b6001600160a01b0316336001600160a01b0316146109105760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b6044820152606401610767565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816040516020016109429190613edb565b60408051601f198184030181529082905261095c91613a3c565b60405180910390a150565b60605f600a8161097a6020870187613876565b6001600160401b0316815260208101919091526040015f20549050806109e25760405162461bcd60e51b815260206004820152601860248201527f416e63686f7220736c6f7420697320756e646566696e656400000000000000006044820152606401610767565b6109ee818560046111a5565b5f6109fe8561010001358561158a565b9050610a1060a085016080860161388f565b6001600160a01b0316610a296040870160208801613876565b6001600160401b03167f00000000000000000000000000000000000000000000000000000000000000007f17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb18860c001358860600135604051610a95929190918252602082015260400190565b60405180910390a4949350505050565b5f80610ab46060840184613d15565b810190610ac191906141ed565b905060408301357f000000000000000000000000000000000000000000000000000000000000000014610b295760405162461bcd60e51b815260206004820152601060248201526f3130b21039b7bab931b29031b430b4b760811b6044820152606401610767565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610b60604085016020860161437e565b63ffffffff1614610ba15760405162461bcd60e51b815260206004820152600b60248201526a626164206e6574776f726b60a81b6044820152606401610767565b6020810151608001516001600160a01b03163014610bef5760405162461bcd60e51b815260206004820152600b60248201526a3130b21032b6b4ba3a32b960a91b6044820152606401610767565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816020015160a0015114610c525760405162461bcd60e51b815260206004820152600960248201526862616420746f70696360b81b6044820152606401610767565b805160208201516040516375f6768760e11b81525f92309263ebeced0e92610c7e92919060040161449d565b5f604051808303815f875af1158015610c99573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610cc0919081019061458a565b90505f81806020019051810190610cd7919061458a565b9050610ce385806145f2565b604051602001610cf39190613edb565b60405160208183030381529060405280519060200120818051906020012014610d515760405162461bcd60e51b815260206004820152601060248201526f0e0c2f2d8dec2c840dad2e6dac2e8c6d60831b6044820152606401610767565b506001949350505050565b610d6681336118d7565b50565b5f610d748383610697565b610df1575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610da93390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161052a565b505f61052a565b5f610e038383610697565b15610df1575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161052a565b5f6001600160401b03821115610e9657604080516306dfcc6560e41b8152600481019190915260248101839052604401610767565b5090565b610ea5602083614625565b6001600160401b031615610ed75760405163c71ab6c360e01b81526001600160401b0383166004820152602401610767565b5f610ee183611914565b90508180610efa57505f818152600b6020526040902054155b15610f0d57610f0d818560e00151611930565b61010084015115610f9c575f600b81610f27846001613aaf565b81526020019081526020015f205490508280610f41575080155b15610f6457610f5f610f54836001613aaf565b866101000151611930565b610f9a565b8461010001518114610f9a57610100850151604051634b7329e360e01b8152600481019190915260248101829052604401610767565b505b6040805160608082018352600180546001600160401b0380821685526002805460208088019190915260038054888a0152885196870189528b8416875260808d01518783019081525f8b8152600b9093529189902054878a018190528751851667ffffffffffffffff19909616861790965590519091559290925583519451939492939092909116907fcb2448a2e087d2fb1bdcd5fb470f39df745400c80391a660a606a2b735aafaa0906110549086908690614652565b60405180910390a361106a85876080015161196e565b505050505050565b5f604051806060016040528061108b8660200151610e61565b6001600160401b0390811682528651602080840191909152604080890151938101939093528251606081018452600154909216825260025490820152600354918101919091529091506110de90826119c9565b6110fb57604051634247f10960e11b815260040160405180910390fd5b80516111079084611a0a565b61112457604051632d776d6b60e11b815260040160405180910390fd5b600e54600d546001600160a01b03909116906341493c60906111468580613d15565b6111536020880188613d15565b6040518663ffffffff1660e01b81526004016111739594939291906146a0565b5f6040518083038186803b158015611189575f5ffd5b505afa15801561119b573d5f5f3e3d5ffd5b5050505050505050565b6111b26020830183613876565b6001600160401b03166111cb6040840160208501613876565b6001600160401b0316106112335760405162461bcd60e51b815260206004820152602960248201527f54617267657420736c6f74206d757374206265206c657373207468616e20616e60448201526818da1bdc881cdb1bdd60ba1b6064820152608401610767565b60058101546112486040840160208501613876565b6112556020850185613876565b61125f91906146d8565b6001600160401b031611156112af5760405162461bcd60e51b815260206004820152601660248201527515185c99d95d081cdb1bdd081a5cc81d1bdbc81bdb1960521b6044820152606401610767565b80545f90611302906040850135906112ca60608701876146f7565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508a9250611aaa915050565b9050806113515760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420616e63686f7220737461746520726f6f74000000000000006044820152606401610767565b60058201545f906113686040860160208701613876565b6001600160401b031661137b919061473c565b90505f8184600301548560020154901b6113959190613aaf565b90505f6113e86080870135836113ae60a08a018a6146f7565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060408a0135611aaa565b9050806114375760405162461bcd60e51b815260206004820181905260248201527f496e76616c69642074617267657420626561636f6e20737461746520726f6f746044820152606401610767565b5f61148f8760c001358760010154898060e0019061145591906146f7565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060808b0135611aaa565b9050806114de5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c696420657865637574696f6e20726f6f742070726f6f66000000006044820152606401610767565b5f61153888610100013588600401548a8061012001906114fe91906146f7565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060c08c0135611aaa565b90508061157f5760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081c9958d95a5c1d1cc81c9bdbdd605a1b6044820152606401610767565b505050505050505050565b6040805160018082528183019092526060915f9190816020015b60608152602001906001900390816115a45790505090506115c86020840184613d15565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92018290525085518694509092501515905061160f5761160f61474f565b60209081029190910101525f6116388561162986806146f7565b61163291614763565b84611ac1565b9050805160011461169e5760405162461bcd60e51b815260206004820152602a60248201527f496e76616c6964206e756d626572206f6620726573756c747320696e2072656360448201526932b4b83a10383937b7b360b11b6064820152608401610767565b5f815f815181106116b1576116b161474f565b60200260200101516020015151116117035760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103932b1b2b4b83a10383937b7b360591b6044820152606401610767565b5f61173261172d835f8151811061171c5761171c61474f565b602002602001015160200151611f24565b611f50565b90508060800151518560600135106117805760405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840d8dece40d2dcc8caf607b1b6044820152606401610767565b5f816080015186606001358151811061179b5761179b61474f565b602002602001015190508560800160208101906117b8919061388f565b6001600160a01b0316815f01516001600160a01b0316146118115760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103637b39032b6b4ba3a32b960691b6044820152606401610767565b5f816020015151116118595760405162461bcd60e51b81526020600482015260116024820152704c6f6720686173206e6f20746f7069637360781b6044820152606401610767565b8560a0013581602001515f815181106118745761187461474f565b6020026020010151146118c95760405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206576656e74207369676e61747572650000000000000000006044820152606401610767565b604001519695505050505050565b6118e18282610697565b6119105760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610767565b5050565b5f6119216120008361476f565b6001600160401b031692915050565b5f828152600b602052604080822083905551829184917f0a1aa8597c2315495a28a7e3a0f54e59bebad4c28bd2884a8f0284d01b5f5c8c9190a35050565b600c805467ffffffffffffffff19166001600160401b0384169081179091555f818152600a6020526040808220849055518392917f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d91a35050565b805182515f916001600160401b0391821691161480156119f0575081602001518360200151145b8015611a03575081604001518360400151145b9392505050565b5f826001600160401b0316826001600160401b031611611a2b57505f61052a565b5f611a3583610658565b905042811115611a48575f91505061052a565b600e54600160a01b900462ffffff16611a61824261479c565b1115611a70575f91505061052a565b5f600c611a7d86866146d8565b6001600160401b0316611a909190613a98565b600e54600160a01b900462ffffff16101595945050505050565b5f81611ab7868686612203565b1495945050505050565b60605f82516001600160401b03811115611add57611add613ac2565b604051908082528060200260200182016040528015611b2257816020015b6040805180820190915260608082526020820152815260200190600190039081611afb5790505b5090505f84516001600160401b03811115611b3f57611b3f613ac2565b604051908082528060200260200182016040528015611b8457816020015b604080518082019091525f815260606020820152815260200190600190039081611b5d5790505b5090505f5b8551811015611c05576040518060400160405280878381518110611baf57611baf61474f565b6020026020010151805190602001208152602001878381518110611bd557611bd561474f565b6020026020010151815250828281518110611bf257611bf261474f565b6020908102919091010152600101611b89565b505f5b8451811015611f1957848181518110611c2357611c2361474f565b6020026020010151838281518110611c3d57611c3d61474f565b60200260200101515f01819052505f6040518060400160405280878481518110611c6957611c6961474f565b602002602001015181526020015f81525090505f611c8f611c8a858b612371565b61240c565b905060015b8015611f0e57604080516080810182525f8082526020820181905291810191909152606080820152611cc583612560565b15611d8e575f611cd484612574565b90505f60028660200151611ce8919061473c565b15611d0e5760028660200151611cfe91906147af565b611d09906001613aaf565b611d1f565b60028660200151611d1f91906147af565b90506040518060400160405280611d39885f015184612619565b81526020015f8152509550611d51825f0151876126e1565b15611d8657611d64888360200151612706565b898881518110611d7657611d7661474f565b6020026020010151602001819052505b505050611f0e565b60c083015115611e27575f611da28461274e565b9050611db185825f015161286a565b15611e1a575f611dc3825f0151612886565b8660200151611dd29190613aaf565b90506040518060400160405280611df8885f0151600285611df391906147af565b612619565b8152602001611e0860028461473c565b81525095508160200151925050611e21565b5050611f0e565b50611eeb565b60e083015115611ede575f611e3b846128a7565b9050611e4685612a9b565b15611e90578051511515600103611e1a57611e6887825f015160200151612706565b888781518110611e7a57611e7a61474f565b6020026020010151602001819052505050611f0e565b5f8160200151611ea0875f612aac565b60108110611eb057611eb061474f565b602002015180519091501515600103611d8657611ece866001612b32565b9550806020015192505050611eeb565b825115611eeb5750611f0e565b611ef8611c8a8783612706565b9250508080611f06906147c2565b915050611c94565b505050600101611c08565b509095945050505050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805160a0810182525f808252606060208301819052928201528181018290526080810191909152611f8282612b73565b61201d57602082015180515f1a60018110801590611fa4575060048160ff1611155b611fe65760405162461bcd60e51b8152602060048201526013602482015272556e737570706f72746564207478207479706560681b6044820152606401610767565b60ff8116835260408051808201909152845181906120069060019061479c565b8152602001612016846001613aaf565b9052935050505b5f61202783612baa565b905080516004146120855760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420726563656970743a2065787065637465642034206669656c604482015261647360f01b6064820152608401610767565b6120a7815f8151811061209a5761209a61474f565b6020026020010151612caf565b82602001819052506120d2816001815181106120c5576120c561474f565b6020026020010151612d20565b6001600160401b0316604083015280516120f9908290600290811061209a5761209a61474f565b82606001819052505f612125826003815181106121185761211861474f565b6020026020010151612baa565b8051909150806001600160401b0381111561214257612142613ac2565b60405190808252806020026020018201604052801561219f57816020015b61218c60405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b8152602001906001900390816121605790505b5060808501525f5b818110156121fa576121d18382815181106121c4576121c461474f565b6020026020010151612d6b565b856080015182815181106121e7576121e761474f565b60209081029190910101526001016121a7565b50505050919050565b5f83815b846001146123685761221a60028661473c565b6001036122b75760028482815181106122355761223561474f565b602002602001015183604051602001612258929190918252602082015260400190565b60408051601f1981840301815290829052612272916147da565b602060405180830381855afa15801561228d573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906122b091906147eb565b9150612349565b6002828583815181106122cc576122cc61474f565b60200260200101516040516020016122ee929190918252602082015260400190565b60408051601f1981840301815290829052612308916147da565b602060405180830381855afa158015612323573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061234691906147eb565b91505b6123546002866147af565b945080612360816147c2565b915050612207565b50949350505050565b60605f5b83518110156123cf57828482815181106123915761239161474f565b60200260200101515f0151036123c7578381815181106123b3576123b361474f565b60200260200101516020015191505061052a565b600101612375565b5060405162461bcd60e51b8152602060048201526011602482015270496e636f6d706c6574652050726f6f662160781b6044820152606401610767565b612414613658565b61241c613658565b5f60405180604001604052808581526020015f8152509050612473846040518060400160405280602081526020017f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421815250612f3f565b1561248357506001815292915050565b5f61249561249086611f24565b612baa565b80519091505f8190036124af575050600182525092915050565b80600203612507575f6124cd835f8151811061209a5761209a61474f565b602081015190915060fc1c60028114806124e75750806003145b156124f85760016020870152612500565b600160c08701525b5050612552565b8060110361251b57600160e0850152612552565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206461746160a01b6044820152606401610767565b505061012082015292915050565b5f81602001518061052a5750506040015190565b61257c6136bd565b6125846136bd565b5f6125996124908561012001515f0151611f24565b90505f6125b28260018151811061209a5761209a61474f565b905060405180604001604052806125de6125d7855f8151811061209a5761209a61474f565b6001612f6d565b81525f602091820181905291855260408051608081018252838152808301939093526001908301526060820192909252908301525092915050565b81516060905f612629848361479c565b90508382101561266f5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606401610767565b60608115801561268d5760405191505f8252602082016040526126d7565b6040519150601f8316801560200281840101848101888315602002848c0101015b818310156126c65780518352602092830192016126ae565b5050848452601f01601f1916604052505b5095945050505050565b5f6126eb82612886565b6126f484612886565b148015611a035750611a03838361286a565b606081604001511561271d5750606081015161052a565b81511561273957612732838360200151612371565b905061052a565b5060408051602081019091525f815292915050565b6127566136bd565b61275e6136bd565b5f6127736124908561012001515f0151611f24565b90505f61278c8260018151811061209a5761209a61474f565b90505f60046127a6845f8151811061209a5761209a61474f565b5f815181106127b7576127b761474f565b602001015160f81c60f81b6001600160f81b031916901c60f81c600116905060405180604001604052806128186127f9865f8151811061209a5761209a61474f565b6002612806866001614802565b612810919061481b565b60ff16612f6d565b815260ff831660209182015290855260408051608081019091526001815290810161284284612fa3565b81525f6020808301829052604080519283528282018152909201528501525091949350505050565b5f61287482612886565b61287e848461300a565b149392505050565b5f81602001516002835f01515161289d9190613a98565b61052a919061479c565b6128af61370c565b6128b761370c565b5f6128cc6124908561012001515f0151611f24565b90506128d6613728565b5f5b60108110156129cf575f6128f784838151811061209a5761209a61474f565b9050805160201461295e576040805180820182525f8082528251608081018452818152602080820183905281850183905284519283528281019094526060810191909152918101919091528383601081106129545761295461474f565b60200201526129c6565b5f61296882612fa3565b6040805180820182526001808252825160808101845290815260208082018590525f82850181905284519081528082019094526060820193909352918101919091529091508484601081106129bf576129bf61474f565b6020020152505b506001016128d8565b506129ee6129e98360108151811061209a5761209a61474f565b61317c565b15612a3a576040805180820182525f8082528251608081018452818152602080820183905281850183905284519283528281019094526060810191909152918101919091528352612a8f565b604051806040016040528060011515815260200160405180608001604052805f151581526020015f5f1b8152602001600115158152602001612a888660108151811061209a5761209a61474f565b9052905283525b60208301525092915050565b5f612aa582612886565b1592915050565b5f5f6002838560200151612ac09190613aaf565b612aca91906147af565b90505f6002848660200151612adf9190613aaf565b612ae9919061473c565b90505f855f01518381518110612b0157612b0161474f565b016020015160f81c905060018214612b205760048160ff16901c612b25565b80600f165b60ff169695505050505050565b60408051808201909152606081525f60208201526040518060400160405280845f01518152602001838560200151612b6a9190613aaf565b90529392505050565b80515f908103612b8457505f919050565b602082015180515f1a9060c0821015612ba057505f9392505050565b5060019392505050565b6060612bb582612b73565b612bbd575f5ffd5b5f612bc7836131e4565b90505f816001600160401b03811115612be257612be2613ac2565b604051908082528060200260200182016040528015612c2657816020015b604080518082019091525f8082526020820152815260200190600190039081612c005790505b5090505f612c378560200151613267565b8560200151612c469190613aaf565b90505f805b84811015612ca457612c5c836132e0565b9150604051806040016040528083815260200184815250848281518110612c8557612c8561474f565b6020908102919091010152612c9a8284613aaf565b9250600101612c4b565b509195945050505050565b8051606090612cbc575f5ffd5b5f5f612cc784613388565b915091505f816001600160401b03811115612ce457612ce4613ac2565b6040519080825280601f01601f191660200182016040528015612d0e576020820181803683370190505b509050602081016123688482856133ca565b80515f9015801590612d3457508151602110155b612d3c575f5ffd5b5f5f612d4784613388565b815191935091506020821015612d635760208290036101000a90045b949350505050565b612d9760405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b5f612da183612baa565b90508051600314612df45760405162461bcd60e51b815260206004820152601e60248201527f496e76616c6964206c6f673a2065787065637465642033206669656c647300006044820152606401610767565b612e16815f81518110612e0957612e0961474f565b602002602001015161344a565b6001600160a01b031682528051612e3a908290600290811061209a5761209a61474f565b82604001819052505f612e59826001815181106121185761211861474f565b80519091506004811115612eaf5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206c6f673a20746f6f206d616e7920746f70696373000000006044820152606401610767565b806001600160401b03811115612ec757612ec7613ac2565b604051908082528060200260200182016040528015612ef0578160200160208202803683370190505b5060208501525f5b818110156121fa57612f158382815181106120c5576120c561474f565b6020860151805183908110612f2c57612f2c61474f565b6020908102919091010152600101612ef8565b5f8151835114612f5057505f61052a565b825160208381018281209186019283209091145b95945050505050565b60608251821115612f7c575f5ffd5b5f828451612f8a919061479c565b905060208401612f64612f9d8583613aaf565b83613462565b5f6020825110156130025760405162461bcd60e51b8152602060048201526024808201527f42797465733a3a20746f427974657333323a206461746120697320746f20736860448201526337b93a1760e11b6064820152608401610767565b506020015190565b5f5f6002846020015161301d919061473c565b90505f60028460200151613031919061473c565b9050808203613121575f6002866020015161304c91906147af565b90505f6002866020015161306091906147af565b90505f84156130de57865180518390811061307d5761307d61474f565b016020015188518051600f60f81b909216918590811061309f5761309f61474f565b0160200151600f60f81b16146130bc575f9550505050505061052a565b6130c5836147c2565b92506130d0826147c2565b91506130db816147c2565b90505b5f6130ec895f015185612619565b90505f6130fc895f015185612619565b90508261310983836134b8565b6131139190613aaf565b97505050505050505061052a565b5f61313c61312e87612886565b61313787612886565b61359e565b90505f5b81811015613171576131528682612aac565b61315c8883612aac565b036131715761316a816147c2565b9050613140565b935061052a92505050565b5f5f825111801561052a5750815f8151811061319a5761319a61474f565b6020910101516001600160f81b031916600360fe1b148061052a5750815f815181106131c8576131c861474f565b6020910101516001600160f81b031916600160ff1b1492915050565b80515f9081036131f557505f919050565b5f5f90505f6132078460200151613267565b84602001516132169190613aaf565b90505f845f0151856020015161322c9190613aaf565b90505b8082101561325e57613240826132e0565b61324a9083613aaf565b915082613256816147c2565b93505061322f565b50909392505050565b80515f90811a608081101561327e57505f92915050565b60b8811080613299575060c08110801590613299575060f881105b156132a75750600192915050565b60c08110156132d4576132bc600160b861483c565b6132c99060ff168261479c565b611a03906001613aaf565b6132bc600160f861483c565b80515f908190811a60808110156132fa5760019150613381565b60b88110156133205761330e60808261479c565b613319906001613aaf565b9150613381565b60c081101561334d5760b78103600185019450806020036101000a85510460018201810193505050613381565b60f88110156133615761330e60c08261479c565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b5f5f5f6133988460200151613267565b90505f8185602001516133ab9190613aaf565b90505f82865f01516133bd919061479c565b9196919550909350505050565b805f036133d657505050565b6020811061340e57825182526133ed602084613aaf565b92506133fa602083613aaf565b915061340760208261479c565b90506133d6565b801561058d575f600161342283602061479c565b61342e90610100614938565b613438919061479c565b84518451821691191617835250505050565b80515f90601514613459575f5ffd5b61052a82612d20565b6060816001600160401b0381111561347c5761347c613ac2565b6040519080825280601f01601f1916602001820160405280156134a6576020820181803683370190505b509050602081016133818482856135b3565b5f5f6134c68451845161359e565b90505f5b81811015613593578381815181106134e4576134e461474f565b602001015160f81c60f81b6001600160f81b03191685828151811061350b5761350b61474f565b01602001516001600160f81b03191614613583576135658582815181106135345761353461474f565b602001015160f81c60f81b8583815181106135515761355161474f565b01602001516001600160f81b03191661361a565b613570600283613a98565b61357a9190613aaf565b9250505061052a565b61358c816147c2565b90506134ca565b612f64600282613a98565b5f8183106135ac5781611a03565b5090919050565b602081106135eb57825182526135ca602083613aaf565b91506135d7602084613aaf565b92506135e460208261479c565b90506135b3565b5f81156105545760016135ff83602061479c565b61360b90610100614938565b613615919061479c565b613438565b5f6001600160f81b0319808316908416036136375750600261052a565b8160f81c60f01660ff168360f81c60f01660ff1603610df15750600161052a565b60408051610140810182525f808252602080830182905282840182905260608084018390526080840183905260a0840183905260c0840183905260e084018390526101008401839052845180860190955284528301529061012082015290565b905290565b60405180604001604052806136e46040518060400160405280606081526020015f81525090565b8152604080516080810182525f80825260208281018290529282015260608082015291015290565b604051806040016040528061371f613756565b81526020016136b85b6040518061020001604052806010905b613740613756565b8152602001906001900390816137385790505090565b60405180604001604052805f151581526020016136b8604080516080810182525f808252602082018190529181019190915260608082015290565b5f602082840312156137a1575f5ffd5b81356001600160e01b031981168114611a03575f5ffd5b5f602082840312156137c8575f5ffd5b5035919050565b80356001600160a01b03811681146137e5575f5ffd5b919050565b5f5f604083850312156137fb575f5ffd5b8235915061380b602084016137cf565b90509250929050565b81516001600160401b0316815260208083015190820152604080830151908201526060810161052a565b5f6020828403121561384e575f5ffd5b813562ffffff81168114611a03575f5ffd5b80356001600160401b03811681146137e5575f5ffd5b5f60208284031215613886575f5ffd5b611a0382613860565b5f6020828403121561389f575f5ffd5b611a03826137cf565b5f5f602083850312156138b9575f5ffd5b82356001600160401b038111156138ce575f5ffd5b8301601f810185136138de575f5ffd5b80356001600160401b038111156138f3575f5ffd5b856020828401011115613904575f5ffd5b6020919091019590945092505050565b5f60208284031215613924575f5ffd5b81356001600160401b03811115613939575f5ffd5b820160408185031215611a03575f5ffd5b5f6020828403121561395a575f5ffd5b81356001600160401b0381111561396f575f5ffd5b82016101208185031215611a03575f5ffd5b5f5f60408385031215613992575f5ffd5b82356001600160401b038111156139a7575f5ffd5b830161014081860312156139b9575f5ffd5b915060208301356001600160401b038111156139d3575f5ffd5b830160c081860312156139e4575f5ffd5b809150509250929050565b5f5b83811015613a095781810151838201526020016139f1565b50505f910152565b5f8151808452613a288160208601602086016139ef565b601f01601f19169290920160200192915050565b602081525f611a036020830184613a11565b5f60208284031215613a5e575f5ffd5b81356001600160401b03811115613a73575f5ffd5b820160808185031215611a03575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761052a5761052a613a84565b8082018082111561052a5761052a613a84565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715613af857613af8613ac2565b60405290565b60405161014081016001600160401b0381118282101715613af857613af8613ac2565b60405160c081016001600160401b0381118282101715613af857613af8613ac2565b604080519081016001600160401b0381118282101715613af857613af8613ac2565b604051601f8201601f191681016001600160401b0381118282101715613b8d57613b8d613ac2565b604052919050565b5f6001600160401b03821115613bad57613bad613ac2565b5060051b60200190565b5f82601f830112613bc6575f5ffd5b8135613bd9613bd482613b95565b613b65565b80828252602082019150602060608402860101925085831115613bfa575f5ffd5b602085015b838110156126d75760608188031215613c16575f5ffd5b613c1e613ad6565b8135815260208083013590820152613c38604083016137cf565b60408201528352602090920191606001613bff565b5f60208284031215613c5d575f5ffd5b81356001600160401b03811115613c72575f5ffd5b82016101408185031215613c84575f5ffd5b613c8c613afe565b813581526020808301359082015260408083013590820152606080830135908201526080808301359082015260a0808301359082015260c0808301359082015260e0808301359082015261010080830135908201526101208201356001600160401b03811115613cfa575f5ffd5b613d0686828501613bb7565b61012083015250949350505050565b5f5f8335601e19843603018112613d2a575f5ffd5b8301803591506001600160401b03821115613d43575f5ffd5b602001915036819003821315613d57575f5ffd5b9250929050565b5f5f8335601e19843603018112613d73575f5ffd5b83016020810192503590506001600160401b03811115613d91575f5ffd5b8060051b3603821315613d57575f5ffd5b8183526020830192505f815f5b84811015613dde576001600160a01b03613dc8836137cf565b1686526020958601959190910190600101613daf565b5093949350505050565b5f5f8335601e19843603018112613dfd575f5ffd5b83016020810192503590506001600160401b03811115613e1b575f5ffd5b8060061b3603821315613d57575f5ffd5b8183526020830192505f815f5b84811015613dde57813586526001600160a01b03613e59602084016137cf565b1660208701526040958601959190910190600101613e39565b5f5f8335601e19843603018112613e87575f5ffd5b83016020810192503590506001600160401b03811115613ea5575f5ffd5b803603821315613d57575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f90613ef39084016137cf565b6001600160a01b038116604084015250613f0f604084016137cf565b6001600160a01b0381166060840152506060830135608083810191909152613f389084016137cf565b6001600160a01b03811660a08401525060a083013560c083810191909152613f6290840184613d5e565b61012060e0850152613f7961014085018284613da2565b915050613f8960e0850185613de8565b848303601f1901610100860152613fa1838284613e2c565b92505050613fb3610100850185613e72565b848303601f1901610120860152613fcb838284613eb3565b9695505050505050565b5f82601f830112613fe4575f5ffd5b8135613ff2613bd482613b95565b8082825260208201915060208360051b860101925085831115614013575f5ffd5b602085015b838110156126d7578035835260209283019201614018565b5f6001600160401b0382111561404857614048613ac2565b50601f01601f191660200190565b5f82601f830112614065575f5ffd5b8135614073613bd482614030565b818152846020838601011115614087575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6140b0613bd484613b95565b838152905060208101600584901b8301858111156140cc575f5ffd5b835b818110156141065780356001600160401b038111156140eb575f5ffd5b6140f788828801614056565b845250602092830192016140ce565b5050509392505050565b5f82601f83011261411f575f5ffd5b611a03838335602085016140a3565b5f60c0828403121561413e575f5ffd5b614146613b21565b905081356001600160401b0381111561415d575f5ffd5b61416984828501614110565b82525060208201356001600160401b03811115614184575f5ffd5b61419084828501614056565b60208301525060408201356001600160401b038111156141ae575f5ffd5b6141ba84828501614056565b604083015250606082810135908201526141d6608083016137cf565b608082015260a09182013591810191909152919050565b5f602082840312156141fd575f5ffd5b81356001600160401b03811115614212575f5ffd5b820160408185031215614223575f5ffd5b61422b613b43565b81356001600160401b03811115614240575f5ffd5b82016101408187031215614252575f5ffd5b61425a613afe565b61426382613860565b815261427160208301613860565b60208201526040828101359082015260608201356001600160401b03811115614298575f5ffd5b6142a488828501613fd5565b6060830152506080828101359082015260a08201356001600160401b038111156142cc575f5ffd5b6142d888828501613fd5565b60a08301525060c0828101359082015260e08201356001600160401b03811115614300575f5ffd5b61430c88828501613fd5565b60e08301525061010082810135908201526101208201356001600160401b03811115614336575f5ffd5b61434288828501613fd5565b6101208301525082525060208201356001600160401b03811115614364575f5ffd5b6143708682850161412e565b602083015250949350505050565b5f6020828403121561438e575f5ffd5b813563ffffffff81168114611a03575f5ffd5b5f8151808452602084019350602083015f5b82811015613dde5781518652602095860195909101906001016143b3565b5f60c08301825160c0855281815180845260e08701915060e08160051b88010193506020830192505f5b8181101561442c5760df19888603018352614417858551613a11565b945060209384019392909201916001016143fb565b50505050602083015184820360208601526144478282613a11565b915050604083015184820360408601526144618282613a11565b91505060608301516060850152608083015161448860808601826001600160a01b03169052565b5060a083015160a08501528091505092915050565b604081526144b76040820184516001600160401b03169052565b5f60208401516144d260608401826001600160401b03169052565b5060408401516080830152606084015161014060a08401526144f86101808401826143a1565b9050608085015160c084015260a0850151603f198483030160e085015261451f82826143a1565b91505060c085015161010084015260e0850151603f198483030161012085015261454982826143a1565b915050610100850151610140840152610120850151603f198483030161016085015261457582826143a1565b9150508281036020840152612f6481856143d1565b5f6020828403121561459a575f5ffd5b81516001600160401b038111156145af575f5ffd5b8201601f810184136145bf575f5ffd5b80516145cd613bd482614030565b8181528560208385010111156145e1575f5ffd5b612f648260208301602086016139ef565b5f823561011e19833603018112614607575f5ffd5b9190910192915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b0383168061463d5761463d614611565b806001600160401b0384160691505092915050565b82516001600160401b03168152602080840151908201526040808401519082015260c0810182516001600160401b0316606083015260208301516080830152604083015160a0830152611a03565b858152606060208201525f6146b9606083018688613eb3565b82810360408401526146cc818587613eb3565b98975050505050505050565b6001600160401b03828116828216039081111561052a5761052a613a84565b5f5f8335601e1984360301811261470c575f5ffd5b8301803591506001600160401b03821115614725575f5ffd5b6020019150600581901b3603821315613d57575f5ffd5b5f8261474a5761474a614611565b500690565b634e487b7160e01b5f52603260045260245ffd5b5f611a033684846140a3565b5f6001600160401b0383168061478757614787614611565b806001600160401b0384160491505092915050565b8181038181111561052a5761052a613a84565b5f826147bd576147bd614611565b500490565b5f600182016147d3576147d3613a84565b5060010190565b5f82516146078184602087016139ef565b5f602082840312156147fb575f5ffd5b5051919050565b60ff818116838216019081111561052a5761052a613a84565b5f60ff83168061482d5761482d614611565b8060ff84160691505092915050565b60ff828116828216039081111561052a5761052a613a84565b6001815b60018411156148905780850481111561487457614874613a84565b600184161561488257908102905b60019390931c928002614859565b935093915050565b5f826148a65750600161052a565b816148b257505f61052a565b81600181146148c857600281146148d2576148ee565b600191505061052a565b60ff8411156148e3576148e3613a84565b50506001821b61052a565b5060208310610133831016604e8410600b8410161715614911575081810a61052a565b61491d5f198484614855565b805f190482111561493057614930613a84565b029392505050565b5f611a03838361489856fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a264697066735822122083e08485b86e70ea382623451c3c7648dd29b8de39c92b1f3e50b70d8465e9ca64736f6c634300081e0033",
}

// ZKAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKAdapterMetaData.ABI instead.
var ZKAdapterABI = ZKAdapterMetaData.ABI

// ZKAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKAdapterMetaData.Bin instead.
var ZKAdapterBin = ZKAdapterMetaData.Bin

// DeployZKAdapter deploys a new Ethereum contract, binding an instance of ZKAdapter to it.
func DeployZKAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId_ *big.Int, startingState ConsensusState, beaconConfig_ ExecutionBeaconConfig, genesisTime_ *big.Int, permissibleTimespan_ *big.Int, verifier common.Address, lightClientVkey_ [32]byte, admin common.Address, superAdmin common.Address) (common.Address, *types.Transaction, *ZKAdapter, error) {
	parsed, err := ZKAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKAdapterBin), backend, sourceChainId_, startingState, beaconConfig_, genesisTime_, permissibleTimespan_, verifier, lightClientVkey_, admin, superAdmin)
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

// SLOTSPERPERIOD is a free data retrieval call binding the contract method 0x2073ee70.
//
// Solidity: function SLOTS_PER_PERIOD() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) SLOTSPERPERIOD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "SLOTS_PER_PERIOD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SLOTSPERPERIOD is a free data retrieval call binding the contract method 0x2073ee70.
//
// Solidity: function SLOTS_PER_PERIOD() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) SLOTSPERPERIOD() (uint64, error) {
	return _ZKAdapter.Contract.SLOTSPERPERIOD(&_ZKAdapter.CallOpts)
}

// SLOTSPERPERIOD is a free data retrieval call binding the contract method 0x2073ee70.
//
// Solidity: function SLOTS_PER_PERIOD() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) SLOTSPERPERIOD() (uint64, error) {
	return _ZKAdapter.Contract.SLOTSPERPERIOD(&_ZKAdapter.CallOpts)
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
// Solidity: function getCurrentState() view returns((uint64,bytes32,bytes32))
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
// Solidity: function getCurrentState() view returns((uint64,bytes32,bytes32))
func (_ZKAdapter *ZKAdapterSession) GetCurrentState() (ConsensusState, error) {
	return _ZKAdapter.Contract.GetCurrentState(&_ZKAdapter.CallOpts)
}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns((uint64,bytes32,bytes32))
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

// LightClientVkey is a free data retrieval call binding the contract method 0x32e0c9a7.
//
// Solidity: function lightClientVkey() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) LightClientVkey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "lightClientVkey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LightClientVkey is a free data retrieval call binding the contract method 0x32e0c9a7.
//
// Solidity: function lightClientVkey() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) LightClientVkey() ([32]byte, error) {
	return _ZKAdapter.Contract.LightClientVkey(&_ZKAdapter.CallOpts)
}

// LightClientVkey is a free data retrieval call binding the contract method 0x32e0c9a7.
//
// Solidity: function lightClientVkey() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) LightClientVkey() ([32]byte, error) {
	return _ZKAdapter.Contract.LightClientVkey(&_ZKAdapter.CallOpts)
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

// SlotTimestamp is a free data retrieval call binding the contract method 0x6b3ba090.
//
// Solidity: function slotTimestamp(uint64 slot) view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) SlotTimestamp(opts *bind.CallOpts, slot uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "slotTimestamp", slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotTimestamp is a free data retrieval call binding the contract method 0x6b3ba090.
//
// Solidity: function slotTimestamp(uint64 slot) view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) SlotTimestamp(slot uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.SlotTimestamp(&_ZKAdapter.CallOpts, slot)
}

// SlotTimestamp is a free data retrieval call binding the contract method 0x6b3ba090.
//
// Solidity: function slotTimestamp(uint64 slot) view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) SlotTimestamp(slot uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.SlotTimestamp(&_ZKAdapter.CallOpts, slot)
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

// SyncCommittees is a free data retrieval call binding the contract method 0x7640d76b.
//
// Solidity: function syncCommittees(uint256 period) view returns(bytes32 syncCommitteeHash)
func (_ZKAdapter *ZKAdapterCaller) SyncCommittees(opts *bind.CallOpts, period *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "syncCommittees", period)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SyncCommittees is a free data retrieval call binding the contract method 0x7640d76b.
//
// Solidity: function syncCommittees(uint256 period) view returns(bytes32 syncCommitteeHash)
func (_ZKAdapter *ZKAdapterSession) SyncCommittees(period *big.Int) ([32]byte, error) {
	return _ZKAdapter.Contract.SyncCommittees(&_ZKAdapter.CallOpts, period)
}

// SyncCommittees is a free data retrieval call binding the contract method 0x7640d76b.
//
// Solidity: function syncCommittees(uint256 period) view returns(bytes32 syncCommitteeHash)
func (_ZKAdapter *ZKAdapterCallerSession) SyncCommittees(period *big.Int) ([32]byte, error) {
	return _ZKAdapter.Contract.SyncCommittees(&_ZKAdapter.CallOpts, period)
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
// Solidity: function manualTransition(bytes publicValues) returns()
func (_ZKAdapter *ZKAdapterTransactor) ManualTransition(opts *bind.TransactOpts, publicValues []byte) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "manualTransition", publicValues)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc4a6dba1.
//
// Solidity: function manualTransition(bytes publicValues) returns()
func (_ZKAdapter *ZKAdapterSession) ManualTransition(publicValues []byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, publicValues)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc4a6dba1.
//
// Solidity: function manualTransition(bytes publicValues) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) ManualTransition(publicValues []byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, publicValues)
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

// UpdateLightClientVkey is a paid mutator transaction binding the contract method 0x978c0ef5.
//
// Solidity: function updateLightClientVkey(bytes32 newLightClientVkey) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdateLightClientVkey(opts *bind.TransactOpts, newLightClientVkey [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updateLightClientVkey", newLightClientVkey)
}

// UpdateLightClientVkey is a paid mutator transaction binding the contract method 0x978c0ef5.
//
// Solidity: function updateLightClientVkey(bytes32 newLightClientVkey) returns()
func (_ZKAdapter *ZKAdapterSession) UpdateLightClientVkey(newLightClientVkey [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateLightClientVkey(&_ZKAdapter.TransactOpts, newLightClientVkey)
}

// UpdateLightClientVkey is a paid mutator transaction binding the contract method 0x978c0ef5.
//
// Solidity: function updateLightClientVkey(bytes32 newLightClientVkey) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdateLightClientVkey(newLightClientVkey [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateLightClientVkey(&_ZKAdapter.TransactOpts, newLightClientVkey)
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

// ZKAdapterLightClientVkeyUpdatedIterator is returned from FilterLightClientVkeyUpdated and is used to iterate over the raw logs and unpacked data for LightClientVkeyUpdated events raised by the ZKAdapter contract.
type ZKAdapterLightClientVkeyUpdatedIterator struct {
	Event *ZKAdapterLightClientVkeyUpdated // Event containing the contract specifics and raw log

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
func (it *ZKAdapterLightClientVkeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterLightClientVkeyUpdated)
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
		it.Event = new(ZKAdapterLightClientVkeyUpdated)
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
func (it *ZKAdapterLightClientVkeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterLightClientVkeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterLightClientVkeyUpdated represents a LightClientVkeyUpdated event raised by the ZKAdapter contract.
type ZKAdapterLightClientVkeyUpdated struct {
	NewVkey [32]byte
	OldVkey [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLightClientVkeyUpdated is a free log retrieval operation binding the contract event 0xe8baf1f3fe086b66369b84759bbe663914aea7e136766a910b2d22d2a7d1f287.
//
// Solidity: event LightClientVkeyUpdated(bytes32 indexed newVkey, bytes32 indexed oldVkey)
func (_ZKAdapter *ZKAdapterFilterer) FilterLightClientVkeyUpdated(opts *bind.FilterOpts, newVkey [][32]byte, oldVkey [][32]byte) (*ZKAdapterLightClientVkeyUpdatedIterator, error) {

	var newVkeyRule []interface{}
	for _, newVkeyItem := range newVkey {
		newVkeyRule = append(newVkeyRule, newVkeyItem)
	}
	var oldVkeyRule []interface{}
	for _, oldVkeyItem := range oldVkey {
		oldVkeyRule = append(oldVkeyRule, oldVkeyItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "LightClientVkeyUpdated", newVkeyRule, oldVkeyRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterLightClientVkeyUpdatedIterator{contract: _ZKAdapter.contract, event: "LightClientVkeyUpdated", logs: logs, sub: sub}, nil
}

// WatchLightClientVkeyUpdated is a free log subscription operation binding the contract event 0xe8baf1f3fe086b66369b84759bbe663914aea7e136766a910b2d22d2a7d1f287.
//
// Solidity: event LightClientVkeyUpdated(bytes32 indexed newVkey, bytes32 indexed oldVkey)
func (_ZKAdapter *ZKAdapterFilterer) WatchLightClientVkeyUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterLightClientVkeyUpdated, newVkey [][32]byte, oldVkey [][32]byte) (event.Subscription, error) {

	var newVkeyRule []interface{}
	for _, newVkeyItem := range newVkey {
		newVkeyRule = append(newVkeyRule, newVkeyItem)
	}
	var oldVkeyRule []interface{}
	for _, oldVkeyItem := range oldVkey {
		oldVkeyRule = append(oldVkeyRule, oldVkeyItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "LightClientVkeyUpdated", newVkeyRule, oldVkeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterLightClientVkeyUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "LightClientVkeyUpdated", log); err != nil {
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

// ParseLightClientVkeyUpdated is a log parse operation binding the contract event 0xe8baf1f3fe086b66369b84759bbe663914aea7e136766a910b2d22d2a7d1f287.
//
// Solidity: event LightClientVkeyUpdated(bytes32 indexed newVkey, bytes32 indexed oldVkey)
func (_ZKAdapter *ZKAdapterFilterer) ParseLightClientVkeyUpdated(log types.Log) (*ZKAdapterLightClientVkeyUpdated, error) {
	event := new(ZKAdapterLightClientVkeyUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "LightClientVkeyUpdated", log); err != nil {
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

// ZKAdapterSyncCommitteeUpdatedIterator is returned from FilterSyncCommitteeUpdated and is used to iterate over the raw logs and unpacked data for SyncCommitteeUpdated events raised by the ZKAdapter contract.
type ZKAdapterSyncCommitteeUpdatedIterator struct {
	Event *ZKAdapterSyncCommitteeUpdated // Event containing the contract specifics and raw log

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
func (it *ZKAdapterSyncCommitteeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterSyncCommitteeUpdated)
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
		it.Event = new(ZKAdapterSyncCommitteeUpdated)
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
func (it *ZKAdapterSyncCommitteeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterSyncCommitteeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterSyncCommitteeUpdated represents a SyncCommitteeUpdated event raised by the ZKAdapter contract.
type ZKAdapterSyncCommitteeUpdated struct {
	Period *big.Int
	Root   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSyncCommitteeUpdated is a free log retrieval operation binding the contract event 0x0a1aa8597c2315495a28a7e3a0f54e59bebad4c28bd2884a8f0284d01b5f5c8c.
//
// Solidity: event SyncCommitteeUpdated(uint256 indexed period, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) FilterSyncCommitteeUpdated(opts *bind.FilterOpts, period []*big.Int, root [][32]byte) (*ZKAdapterSyncCommitteeUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "SyncCommitteeUpdated", periodRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterSyncCommitteeUpdatedIterator{contract: _ZKAdapter.contract, event: "SyncCommitteeUpdated", logs: logs, sub: sub}, nil
}

// WatchSyncCommitteeUpdated is a free log subscription operation binding the contract event 0x0a1aa8597c2315495a28a7e3a0f54e59bebad4c28bd2884a8f0284d01b5f5c8c.
//
// Solidity: event SyncCommitteeUpdated(uint256 indexed period, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) WatchSyncCommitteeUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterSyncCommitteeUpdated, period []*big.Int, root [][32]byte) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "SyncCommitteeUpdated", periodRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterSyncCommitteeUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "SyncCommitteeUpdated", log); err != nil {
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

// ParseSyncCommitteeUpdated is a log parse operation binding the contract event 0x0a1aa8597c2315495a28a7e3a0f54e59bebad4c28bd2884a8f0284d01b5f5c8c.
//
// Solidity: event SyncCommitteeUpdated(uint256 indexed period, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) ParseSyncCommitteeUpdated(log types.Log) (*ZKAdapterSyncCommitteeUpdated, error) {
	event := new(ZKAdapterSyncCommitteeUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "SyncCommitteeUpdated", log); err != nil {
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
	PreHead   uint64
	PostHead  uint64
	PreState  ConsensusState
	PostState ConsensusState
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransitioned is a free log retrieval operation binding the contract event 0xcb2448a2e087d2fb1bdcd5fb470f39df745400c80391a660a606a2b735aafaa0.
//
// Solidity: event Transitioned(uint64 indexed preHead, uint64 indexed postHead, (uint64,bytes32,bytes32) preState, (uint64,bytes32,bytes32) postState)
func (_ZKAdapter *ZKAdapterFilterer) FilterTransitioned(opts *bind.FilterOpts, preHead []uint64, postHead []uint64) (*ZKAdapterTransitionedIterator, error) {

	var preHeadRule []interface{}
	for _, preHeadItem := range preHead {
		preHeadRule = append(preHeadRule, preHeadItem)
	}
	var postHeadRule []interface{}
	for _, postHeadItem := range postHead {
		postHeadRule = append(postHeadRule, postHeadItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "Transitioned", preHeadRule, postHeadRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTransitionedIterator{contract: _ZKAdapter.contract, event: "Transitioned", logs: logs, sub: sub}, nil
}

// WatchTransitioned is a free log subscription operation binding the contract event 0xcb2448a2e087d2fb1bdcd5fb470f39df745400c80391a660a606a2b735aafaa0.
//
// Solidity: event Transitioned(uint64 indexed preHead, uint64 indexed postHead, (uint64,bytes32,bytes32) preState, (uint64,bytes32,bytes32) postState)
func (_ZKAdapter *ZKAdapterFilterer) WatchTransitioned(opts *bind.WatchOpts, sink chan<- *ZKAdapterTransitioned, preHead []uint64, postHead []uint64) (event.Subscription, error) {

	var preHeadRule []interface{}
	for _, preHeadItem := range preHead {
		preHeadRule = append(preHeadRule, preHeadItem)
	}
	var postHeadRule []interface{}
	for _, postHeadItem := range postHead {
		postHeadRule = append(postHeadRule, postHeadItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "Transitioned", preHeadRule, postHeadRule)
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

// ParseTransitioned is a log parse operation binding the contract event 0xcb2448a2e087d2fb1bdcd5fb470f39df745400c80391a660a606a2b735aafaa0.
//
// Solidity: event Transitioned(uint64 indexed preHead, uint64 indexed postHead, (uint64,bytes32,bytes32) preState, (uint64,bytes32,bytes32) postState)
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
