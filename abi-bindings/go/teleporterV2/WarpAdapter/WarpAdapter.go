// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package warpadapter

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

// WarpAdapterMetaData contains all meta data concerning the WarpAdapter contract.
var WarpAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761069d908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163b771b3bc1461044e57508063eb97cd2c1461038a5763f1faff001461003d575f80fd5b3461016d57602036600319011261016d5760043567ffffffffffffffff811161016d5780360390608060031983011261016d576064810135602219830181121561016d578101600481013567ffffffffffffffff811161016d576024820191813603831361016d576020918101031261016d573563ffffffff811680910361016d576040516306f8253560e41b815260048101919091525f816024816005600160991b015afa801561037f575f915f9161026f575b50156102205760208101516001600160a01b031630036101c857805160448301350361017157604001516020815191012090806004013592610122190183121561016d57600461016160209461015360405193849288840196010185610481565b03601f19810183528261064d565b51902014604051908152f35b5f80fd5b60405162461bcd60e51b815260206004820152602960248201527f57617270416461707465723a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608490fd5b60405162461bcd60e51b815260206004820152602a60248201527f57617270416461707465723a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602160248201527f57617270416461707465723a20696e76616c69642077617270206d65737361676044820152606560f81b6064820152608490fd5b9150503d805f833e610281818361064d565b810160408282031261016d57815167ffffffffffffffff811161016d5782019160608383031261016d57604051926060840184811067ffffffffffffffff82111761036b576040528051845260208101516001600160a01b038116810361016d57602085015260408101519067ffffffffffffffff821161016d57019180601f8401121561016d5782519267ffffffffffffffff841161036b5760405191610333601f8601601f19166020018461064d565b8483526020858301011161016d576020936103539185808501910161066f565b6040840152015190811515820361016d57905f6100f2565b634e487b7160e01b5f52604160045260245ffd5b6040513d5f823e3d90fd5b3461016d57602036600319011261016d5760043567ffffffffffffffff811161016d57610120600319823603011261016d57602061040460446040516103da816101538682019760040188610481565b604051948593849263ee5b48eb60e01b84528660048501525180928160248601528585019061066f565b601f01601f191681010301815f6005600160991b015af1801561037f5761042757005b602090813d8311610447575b61043d818361064d565b8101031261016d57005b503d610433565b3461016d575f36600319011261016d576005600160991b018152602090f35b35906001600160a01b038216820361016d57565b602080825282358183015290916101408301916001600160a01b03906104a890830161046d565b16604084015260018060a01b036104c16040830161046d565b1660608401526060810135608084015260018060a01b036104e46080830161046d565b1660a084015260a081013560c084015260c0810135601e1982360301908181121561016d5782016020813591019367ffffffffffffffff821161016d578160051b3603851361016d57819061012060e088015252610160850193905f5b8181106106275750505060e08201358181121561016d5782016020813591019367ffffffffffffffff821161016d578160061b3603851361016d57858103601f190161010087015281815260200193905f5b8181106105f7575050506101008201359081121561016d57016020813591019267ffffffffffffffff821161016d57813603841361016d576020938291610120601f1982870301910152818452848401375f828201840152601f01601f1916010190565b90919460408060019288358152838060a01b0361061660208b0161046d565b166020820152019601929101610593565b909194602080600192838060a01b0361063f8a61046d565b168152019601929101610541565b90601f8019910116810190811067ffffffffffffffff82111761036b57604052565b5f5b8381106106805750505f910152565b818101518382015260200161067156fea164736f6c634300081e000a",
}

// WarpAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use WarpAdapterMetaData.ABI instead.
var WarpAdapterABI = WarpAdapterMetaData.ABI

// WarpAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WarpAdapterMetaData.Bin instead.
var WarpAdapterBin = WarpAdapterMetaData.Bin

// DeployWarpAdapter deploys a new Ethereum contract, binding an instance of WarpAdapter to it.
func DeployWarpAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WarpAdapter, error) {
	parsed, err := WarpAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WarpAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WarpAdapter{WarpAdapterCaller: WarpAdapterCaller{contract: contract}, WarpAdapterTransactor: WarpAdapterTransactor{contract: contract}, WarpAdapterFilterer: WarpAdapterFilterer{contract: contract}}, nil
}

// WarpAdapter is an auto generated Go binding around an Ethereum contract.
type WarpAdapter struct {
	WarpAdapterCaller     // Read-only binding to the contract
	WarpAdapterTransactor // Write-only binding to the contract
	WarpAdapterFilterer   // Log filterer for contract events
}

// WarpAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarpAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarpAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarpAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarpAdapterSession struct {
	Contract     *WarpAdapter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WarpAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarpAdapterCallerSession struct {
	Contract *WarpAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WarpAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarpAdapterTransactorSession struct {
	Contract     *WarpAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WarpAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarpAdapterRaw struct {
	Contract *WarpAdapter // Generic contract binding to access the raw methods on
}

// WarpAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarpAdapterCallerRaw struct {
	Contract *WarpAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// WarpAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarpAdapterTransactorRaw struct {
	Contract *WarpAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarpAdapter creates a new instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapter(address common.Address, backend bind.ContractBackend) (*WarpAdapter, error) {
	contract, err := bindWarpAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarpAdapter{WarpAdapterCaller: WarpAdapterCaller{contract: contract}, WarpAdapterTransactor: WarpAdapterTransactor{contract: contract}, WarpAdapterFilterer: WarpAdapterFilterer{contract: contract}}, nil
}

// NewWarpAdapterCaller creates a new read-only instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterCaller(address common.Address, caller bind.ContractCaller) (*WarpAdapterCaller, error) {
	contract, err := bindWarpAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterCaller{contract: contract}, nil
}

// NewWarpAdapterTransactor creates a new write-only instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*WarpAdapterTransactor, error) {
	contract, err := bindWarpAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterTransactor{contract: contract}, nil
}

// NewWarpAdapterFilterer creates a new log filterer instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*WarpAdapterFilterer, error) {
	contract, err := bindWarpAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterFilterer{contract: contract}, nil
}

// bindWarpAdapter binds a generic wrapper to an already deployed contract.
func bindWarpAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarpAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarpAdapter *WarpAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarpAdapter.Contract.WarpAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarpAdapter *WarpAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarpAdapter.Contract.WarpAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarpAdapter *WarpAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarpAdapter.Contract.WarpAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarpAdapter *WarpAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarpAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarpAdapter *WarpAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarpAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarpAdapter *WarpAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarpAdapter.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarpAdapter.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterSession) WARPMESSENGER() (common.Address, error) {
	return _WarpAdapter.Contract.WARPMESSENGER(&_WarpAdapter.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterCallerSession) WARPMESSENGER() (common.Address, error) {
	return _WarpAdapter.Contract.WARPMESSENGER(&_WarpAdapter.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterCaller) VerifyMessage(opts *bind.CallOpts, message TeleporterICMMessage) (bool, error) {
	var out []interface{}
	err := _WarpAdapter.contract.Call(opts, &out, "verifyMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterSession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _WarpAdapter.Contract.VerifyMessage(&_WarpAdapter.CallOpts, message)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterCallerSession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _WarpAdapter.Contract.VerifyMessage(&_WarpAdapter.CallOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _WarpAdapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _WarpAdapter.Contract.SendMessage(&_WarpAdapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _WarpAdapter.Contract.SendMessage(&_WarpAdapter.TransactOpts, message)
}
