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
	Bin: "0x6080604052348015600e575f5ffd5b506109088061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063b771b3bc14610043578063eb97cd2c1461006e578063f1faff0014610083575b5f5ffd5b6100516005600160991b0181565b6040516001600160a01b0390911681526020015b60405180910390f35b61008161007c36600461034e565b6100a6565b005b61009661009136600461038d565b61013c565b6040519015158152602001610065565b6005600160991b016001600160a01b031663ee5b48eb826040516020016100cd9190610576565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016100f89190610692565b6020604051808303815f875af1158015610114573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013891906106c4565b5050565b5f8061014b60608401846106db565b810190610158919061071e565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156101a4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101cb91908101906107be565b915091508061022b5760405162461bcd60e51b815260206004820152602160248201527f57617270416461707465723a20696e76616c69642077617270206d65737361676044820152606560f81b60648201526084015b60405180910390fd5b60208201516001600160a01b0316301461029a5760405162461bcd60e51b815260206004820152602a60248201527f57617270416461707465723a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608401610222565b81516040860135146103005760405162461bcd60e51b815260206004820152602960248201527f57617270416461707465723a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608401610222565b604082015180516020909101205f61031887806108b3565b6040516020016103289190610576565b60408051601f198184030181529190528051602090910120919091149695505050505050565b5f6020828403121561035e575f5ffd5b813567ffffffffffffffff811115610374575f5ffd5b82016101208185031215610386575f5ffd5b9392505050565b5f6020828403121561039d575f5ffd5b813567ffffffffffffffff8111156103b3575f5ffd5b820160808185031215610386575f5ffd5b6001600160a01b03811681146103d8575f5ffd5b50565b80356103e6816103c4565b919050565b5f5f8335601e19843603018112610400575f5ffd5b830160208101925035905067ffffffffffffffff81111561041f575f5ffd5b8060051b3603821315610430575f5ffd5b9250929050565b8183526020830192505f815f5b84811015610475578135610457816103c4565b6001600160a01b031686526020958601959190910190600101610444565b5093949350505050565b5f5f8335601e19843603018112610494575f5ffd5b830160208101925035905067ffffffffffffffff8111156104b3575f5ffd5b8060061b3603821315610430575f5ffd5b8183526020830192505f815f5b84811015610475578135865260208201356104eb816103c4565b6001600160a01b0316602087015260409586019591909101906001016104d1565b5f5f8335601e19843603018112610521575f5ffd5b830160208101925035905067ffffffffffffffff811115610540575f5ffd5b803603821315610430575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f9061058e9084016103db565b6001600160a01b0381166040840152506105aa604084016103db565b6001600160a01b03811660608401525060608301356080838101919091526105d39084016103db565b6001600160a01b03811660a08401525060a083013560c0838101919091526105fd908401846103eb565b61012060e085015261061461014085018284610437565b91505061062460e085018561047f565b848303601f190161010086015261063c8382846104c4565b9250505061064e61010085018561050c565b848303601f190161012086015261066683828461054e565b9695505050505050565b5f5b8381101561068a578181015183820152602001610672565b50505f910152565b602081525f82518060208401526106b0816040850160208701610670565b601f01601f19169190910160400192915050565b5f602082840312156106d4575f5ffd5b5051919050565b5f5f8335601e198436030181126106f0575f5ffd5b83018035915067ffffffffffffffff82111561070a575f5ffd5b602001915036819003821315610430575f5ffd5b5f6020828403121561072e575f5ffd5b813563ffffffff81168114610386575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561077857610778610741565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156107a7576107a7610741565b604052919050565b805180151581146103e6575f5ffd5b5f5f604083850312156107cf575f5ffd5b825167ffffffffffffffff8111156107e5575f5ffd5b8301606081860312156107f6575f5ffd5b6107fe610755565b815181526020820151610810816103c4565b6020820152604082015167ffffffffffffffff81111561082e575f5ffd5b80830192505085601f830112610842575f5ffd5b815167ffffffffffffffff81111561085c5761085c610741565b61086f601f8201601f191660200161077e565b818152876020838601011115610883575f5ffd5b610894826020830160208701610670565b60408301525092506108aa9050602084016107af565b90509250929050565b5f823561011e198336030181126108c8575f5ffd5b919091019291505056fea2646970667358221220386e6f0965a779b7efe604ebba0559f6f3ebba400c6ee64cf099264ba3bd913764736f6c634300081e0033",
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
