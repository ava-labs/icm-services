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
	Bin: "0x6080604052348015600e575f5ffd5b50610a388061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063b771b3bc14610043578063eb97cd2c1461006e578063f1faff0014610083575b5f5ffd5b6100516005600160991b0181565b6040516001600160a01b0390911681526020015b60405180910390f35b61008161007c366004610448565b6100a6565b005b610096610091366004610487565b61023b565b6040519015158152602001610065565b6100b660608201604083016104e5565b6001600160a01b0316336001600160a01b0316148061015457506100e060608201604083016104e5565b6001600160a01b031663d67bdd256040518163ffffffff1660e01b8152600401602060405180830381865afa15801561011b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013f9190610500565b6001600160a01b0316336001600160a01b0316145b6101a55760405162461bcd60e51b815260206004820181905260248201527f57617270416461707465723a20756e617574686f72697a65642073656e64657260448201526064015b60405180910390fd5b6005600160991b016001600160a01b031663ee5b48eb826040516020016101cc91906106a6565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016101f791906107c2565b6020604051808303815f875af1158015610213573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061023791906107f4565b5050565b5f8061024a606084018461080b565b810190610257919061084e565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156102a3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526102ca91908101906108ee565b91509150806103255760405162461bcd60e51b815260206004820152602160248201527f57617270416461707465723a20696e76616c69642077617270206d65737361676044820152606560f81b606482015260840161019c565b60208201516001600160a01b031630146103945760405162461bcd60e51b815260206004820152602a60248201527f57617270416461707465723a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b606482015260840161019c565b81516040860135146103fa5760405162461bcd60e51b815260206004820152602960248201527f57617270416461707465723a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b606482015260840161019c565b604082015180516020909101205f61041287806109e3565b60405160200161042291906106a6565b60408051601f198184030181529190528051602090910120919091149695505050505050565b5f60208284031215610458575f5ffd5b813567ffffffffffffffff81111561046e575f5ffd5b82016101208185031215610480575f5ffd5b9392505050565b5f60208284031215610497575f5ffd5b813567ffffffffffffffff8111156104ad575f5ffd5b820160808185031215610480575f5ffd5b6001600160a01b03811681146104d2575f5ffd5b50565b80356104e0816104be565b919050565b5f602082840312156104f5575f5ffd5b8135610480816104be565b5f60208284031215610510575f5ffd5b8151610480816104be565b5f5f8335601e19843603018112610530575f5ffd5b830160208101925035905067ffffffffffffffff81111561054f575f5ffd5b8060051b3603821315610560575f5ffd5b9250929050565b8183526020830192505f815f5b848110156105a5578135610587816104be565b6001600160a01b031686526020958601959190910190600101610574565b5093949350505050565b5f5f8335601e198436030181126105c4575f5ffd5b830160208101925035905067ffffffffffffffff8111156105e3575f5ffd5b8060061b3603821315610560575f5ffd5b8183526020830192505f815f5b848110156105a55781358652602082013561061b816104be565b6001600160a01b031660208701526040958601959190910190600101610601565b5f5f8335601e19843603018112610651575f5ffd5b830160208101925035905067ffffffffffffffff811115610670575f5ffd5b803603821315610560575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f906106be9084016104d5565b6001600160a01b0381166040840152506106da604084016104d5565b6001600160a01b03811660608401525060608301356080838101919091526107039084016104d5565b6001600160a01b03811660a08401525060a083013560c08381019190915261072d9084018461051b565b61012060e085015261074461014085018284610567565b91505061075460e08501856105af565b848303601f190161010086015261076c8382846105f4565b9250505061077e61010085018561063c565b848303601f190161012086015261079683828461067e565b9695505050505050565b5f5b838110156107ba5781810151838201526020016107a2565b50505f910152565b602081525f82518060208401526107e08160408501602087016107a0565b601f01601f19169190910160400192915050565b5f60208284031215610804575f5ffd5b5051919050565b5f5f8335601e19843603018112610820575f5ffd5b83018035915067ffffffffffffffff82111561083a575f5ffd5b602001915036819003821315610560575f5ffd5b5f6020828403121561085e575f5ffd5b813563ffffffff81168114610480575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156108a8576108a8610871565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108d7576108d7610871565b604052919050565b805180151581146104e0575f5ffd5b5f5f604083850312156108ff575f5ffd5b825167ffffffffffffffff811115610915575f5ffd5b830160608186031215610926575f5ffd5b61092e610885565b815181526020820151610940816104be565b6020820152604082015167ffffffffffffffff81111561095e575f5ffd5b80830192505085601f830112610972575f5ffd5b815167ffffffffffffffff81111561098c5761098c610871565b61099f601f8201601f19166020016108ae565b8181528760208386010111156109b3575f5ffd5b6109c48260208301602087016107a0565b60408301525092506109da9050602084016108df565b90509250929050565b5f823561011e198336030181126109f8575f5ffd5b919091019291505056fea2646970667358221220ceb9bc8887713a00fdeee399444a96c0baf633d15c540b0aec12c8a56f91f49c64736f6c634300081e0033",
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
