// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockoraclereceiver

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

// MockOracleReceiverMetaData contains all meta data concerning the MockOracleReceiver contract.
var MockOracleReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"oracleAdapter_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastPayload\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleAdapter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveOracleMessage\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"OnlyOracleAdapter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a0604052348015600e575f5ffd5b50604051610655380380610655833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f5ffd5b81516001600160a01b0381168114605f575f5ffd5b9392505050565b6080516105ca61008b5f395f818160d3015281816101ea015261022c01526105ca5ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c806344c191391161005857806344c191391461010d5780634cc2aa3c1461011557806352631ab41461011d5780637ba855921461014a575f5ffd5b80630a4e00bf146100895780630d768691146100a45780631c67934b146100b95780633d9f3163146100ce575b5f5ffd5b6100915f5481565b6040519081526020015b60405180910390f35b6100ac610153565b60405161009b9190610323565b6100cc6100c7366004610381565b6101df565b005b6100f57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161009b565b6100ac6102c6565b6100ac6102d3565b6003546101319067ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161009b565b61009160055481565b6002805461016090610447565b80601f016020809104026020016040519081016040528092919081815260200182805461018c90610447565b80156101d75780601f106101ae576101008083540402835291602001916101d7565b820191905f5260205f20905b8154815290600101906020018083116101ba57829003601f168201915b505050505081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461025d5760405163b091d6af60e01b81523360048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602482015260440160405180910390fd5b5f889055600161026e8789836104df565b50600261027c8587836104df565b506003805467ffffffffffffffff191667ffffffffffffffff851617905560046102a78284836104df565b5060058054905f6102b783610599565b91905055505050505050505050565b6001805461016090610447565b6004805461016090610447565b5f81518084525f5b81811015610304576020818501810151868301820152016102e8565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f61033560208301846102e0565b9392505050565b5f5f83601f84011261034c575f5ffd5b50813567ffffffffffffffff811115610363575f5ffd5b60208301915083602082850101111561037a575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60a0898b031215610398575f5ffd5b88359750602089013567ffffffffffffffff8111156103b5575f5ffd5b6103c18b828c0161033c565b909850965050604089013567ffffffffffffffff8111156103e0575f5ffd5b6103ec8b828c0161033c565b909650945050606089013567ffffffffffffffff8116811461040c575f5ffd5b9250608089013567ffffffffffffffff811115610427575f5ffd5b6104338b828c0161033c565b999c989b5096995094979396929594505050565b600181811c9082168061045b57607f821691505b60208210810361047957634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b601f8211156104da57805f5260205f20601f840160051c810160208510156104b85750805b601f840160051c820191505b818110156104d7575f81556001016104c4565b50505b505050565b67ffffffffffffffff8311156104f7576104f761047f565b61050b836105058354610447565b83610493565b5f601f84116001811461053c575f85156105255750838201355b5f19600387901b1c1916600186901b1783556104d7565b5f83815260208120601f198716915b8281101561056b578685013582556020948501946001909201910161054b565b5086821015610587575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f600182016105b657634e487b7160e01b5f52601160045260245ffd5b506001019056fea164736f6c634300081e000a",
}

// MockOracleReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockOracleReceiverMetaData.ABI instead.
var MockOracleReceiverABI = MockOracleReceiverMetaData.ABI

// MockOracleReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockOracleReceiverMetaData.Bin instead.
var MockOracleReceiverBin = MockOracleReceiverMetaData.Bin

// DeployMockOracleReceiver deploys a new Ethereum contract, binding an instance of MockOracleReceiver to it.
func DeployMockOracleReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, oracleAdapter_ common.Address) (common.Address, *types.Transaction, *MockOracleReceiver, error) {
	parsed, err := MockOracleReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockOracleReceiverBin), backend, oracleAdapter_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockOracleReceiver{MockOracleReceiverCaller: MockOracleReceiverCaller{contract: contract}, MockOracleReceiverTransactor: MockOracleReceiverTransactor{contract: contract}, MockOracleReceiverFilterer: MockOracleReceiverFilterer{contract: contract}}, nil
}

// MockOracleReceiver is an auto generated Go binding around an Ethereum contract.
type MockOracleReceiver struct {
	MockOracleReceiverCaller     // Read-only binding to the contract
	MockOracleReceiverTransactor // Write-only binding to the contract
	MockOracleReceiverFilterer   // Log filterer for contract events
}

// MockOracleReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockOracleReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockOracleReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockOracleReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOracleReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockOracleReceiverSession struct {
	Contract     *MockOracleReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockOracleReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockOracleReceiverCallerSession struct {
	Contract *MockOracleReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockOracleReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockOracleReceiverTransactorSession struct {
	Contract     *MockOracleReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockOracleReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockOracleReceiverRaw struct {
	Contract *MockOracleReceiver // Generic contract binding to access the raw methods on
}

// MockOracleReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockOracleReceiverCallerRaw struct {
	Contract *MockOracleReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockOracleReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockOracleReceiverTransactorRaw struct {
	Contract *MockOracleReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockOracleReceiver creates a new instance of MockOracleReceiver, bound to a specific deployed contract.
func NewMockOracleReceiver(address common.Address, backend bind.ContractBackend) (*MockOracleReceiver, error) {
	contract, err := bindMockOracleReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockOracleReceiver{MockOracleReceiverCaller: MockOracleReceiverCaller{contract: contract}, MockOracleReceiverTransactor: MockOracleReceiverTransactor{contract: contract}, MockOracleReceiverFilterer: MockOracleReceiverFilterer{contract: contract}}, nil
}

// NewMockOracleReceiverCaller creates a new read-only instance of MockOracleReceiver, bound to a specific deployed contract.
func NewMockOracleReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockOracleReceiverCaller, error) {
	contract, err := bindMockOracleReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleReceiverCaller{contract: contract}, nil
}

// NewMockOracleReceiverTransactor creates a new write-only instance of MockOracleReceiver, bound to a specific deployed contract.
func NewMockOracleReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockOracleReceiverTransactor, error) {
	contract, err := bindMockOracleReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockOracleReceiverTransactor{contract: contract}, nil
}

// NewMockOracleReceiverFilterer creates a new log filterer instance of MockOracleReceiver, bound to a specific deployed contract.
func NewMockOracleReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockOracleReceiverFilterer, error) {
	contract, err := bindMockOracleReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockOracleReceiverFilterer{contract: contract}, nil
}

// bindMockOracleReceiver binds a generic wrapper to an already deployed contract.
func bindMockOracleReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockOracleReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracleReceiver *MockOracleReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracleReceiver.Contract.MockOracleReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracleReceiver *MockOracleReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.MockOracleReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracleReceiver *MockOracleReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.MockOracleReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOracleReceiver *MockOracleReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOracleReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOracleReceiver *MockOracleReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOracleReceiver *MockOracleReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.contract.Transact(opts, method, params...)
}

// LastNonce is a free data retrieval call binding the contract method 0x52631ab4.
//
// Solidity: function lastNonce() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastNonce is a free data retrieval call binding the contract method 0x52631ab4.
//
// Solidity: function lastNonce() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverSession) LastNonce() (uint64, error) {
	return _MockOracleReceiver.Contract.LastNonce(&_MockOracleReceiver.CallOpts)
}

// LastNonce is a free data retrieval call binding the contract method 0x52631ab4.
//
// Solidity: function lastNonce() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastNonce() (uint64, error) {
	return _MockOracleReceiver.Contract.LastNonce(&_MockOracleReceiver.CallOpts)
}

// LastPayload is a free data retrieval call binding the contract method 0x4cc2aa3c.
//
// Solidity: function lastPayload() view returns(bytes)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastPayload(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastPayload")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LastPayload is a free data retrieval call binding the contract method 0x4cc2aa3c.
//
// Solidity: function lastPayload() view returns(bytes)
func (_MockOracleReceiver *MockOracleReceiverSession) LastPayload() ([]byte, error) {
	return _MockOracleReceiver.Contract.LastPayload(&_MockOracleReceiver.CallOpts)
}

// LastPayload is a free data retrieval call binding the contract method 0x4cc2aa3c.
//
// Solidity: function lastPayload() view returns(bytes)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastPayload() ([]byte, error) {
	return _MockOracleReceiver.Contract.LastPayload(&_MockOracleReceiver.CallOpts)
}

// LastSourceAddress is a free data retrieval call binding the contract method 0x0d768691.
//
// Solidity: function lastSourceAddress() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastSourceAddress(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastSourceAddress")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastSourceAddress is a free data retrieval call binding the contract method 0x0d768691.
//
// Solidity: function lastSourceAddress() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverSession) LastSourceAddress() (string, error) {
	return _MockOracleReceiver.Contract.LastSourceAddress(&_MockOracleReceiver.CallOpts)
}

// LastSourceAddress is a free data retrieval call binding the contract method 0x0d768691.
//
// Solidity: function lastSourceAddress() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastSourceAddress() (string, error) {
	return _MockOracleReceiver.Contract.LastSourceAddress(&_MockOracleReceiver.CallOpts)
}

// LastSourceChainID is a free data retrieval call binding the contract method 0x0a4e00bf.
//
// Solidity: function lastSourceChainID() view returns(bytes32)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastSourceChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastSourceChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastSourceChainID is a free data retrieval call binding the contract method 0x0a4e00bf.
//
// Solidity: function lastSourceChainID() view returns(bytes32)
func (_MockOracleReceiver *MockOracleReceiverSession) LastSourceChainID() ([32]byte, error) {
	return _MockOracleReceiver.Contract.LastSourceChainID(&_MockOracleReceiver.CallOpts)
}

// LastSourceChainID is a free data retrieval call binding the contract method 0x0a4e00bf.
//
// Solidity: function lastSourceChainID() view returns(bytes32)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastSourceChainID() ([32]byte, error) {
	return _MockOracleReceiver.Contract.LastSourceChainID(&_MockOracleReceiver.CallOpts)
}

// LastSourceType is a free data retrieval call binding the contract method 0x44c19139.
//
// Solidity: function lastSourceType() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastSourceType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastSourceType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastSourceType is a free data retrieval call binding the contract method 0x44c19139.
//
// Solidity: function lastSourceType() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverSession) LastSourceType() (string, error) {
	return _MockOracleReceiver.Contract.LastSourceType(&_MockOracleReceiver.CallOpts)
}

// LastSourceType is a free data retrieval call binding the contract method 0x44c19139.
//
// Solidity: function lastSourceType() view returns(string)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastSourceType() (string, error) {
	return _MockOracleReceiver.Contract.LastSourceType(&_MockOracleReceiver.CallOpts)
}

// OracleAdapter is a free data retrieval call binding the contract method 0x3d9f3163.
//
// Solidity: function oracleAdapter() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverCaller) OracleAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "oracleAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAdapter is a free data retrieval call binding the contract method 0x3d9f3163.
//
// Solidity: function oracleAdapter() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverSession) OracleAdapter() (common.Address, error) {
	return _MockOracleReceiver.Contract.OracleAdapter(&_MockOracleReceiver.CallOpts)
}

// OracleAdapter is a free data retrieval call binding the contract method 0x3d9f3163.
//
// Solidity: function oracleAdapter() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) OracleAdapter() (common.Address, error) {
	return _MockOracleReceiver.Contract.OracleAdapter(&_MockOracleReceiver.CallOpts)
}

// ReceiveCount is a free data retrieval call binding the contract method 0x7ba85592.
//
// Solidity: function receiveCount() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCaller) ReceiveCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "receiveCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceiveCount is a free data retrieval call binding the contract method 0x7ba85592.
//
// Solidity: function receiveCount() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverSession) ReceiveCount() (*big.Int, error) {
	return _MockOracleReceiver.Contract.ReceiveCount(&_MockOracleReceiver.CallOpts)
}

// ReceiveCount is a free data retrieval call binding the contract method 0x7ba85592.
//
// Solidity: function receiveCount() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) ReceiveCount() (*big.Int, error) {
	return _MockOracleReceiver.Contract.ReceiveCount(&_MockOracleReceiver.CallOpts)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1c67934b.
//
// Solidity: function receiveOracleMessage(bytes32 sourceChainID, string sourceType, string sourceAddress, uint64 nonce, bytes payload) returns()
func (_MockOracleReceiver *MockOracleReceiverTransactor) ReceiveOracleMessage(opts *bind.TransactOpts, sourceChainID [32]byte, sourceType string, sourceAddress string, nonce uint64, payload []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.contract.Transact(opts, "receiveOracleMessage", sourceChainID, sourceType, sourceAddress, nonce, payload)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1c67934b.
//
// Solidity: function receiveOracleMessage(bytes32 sourceChainID, string sourceType, string sourceAddress, uint64 nonce, bytes payload) returns()
func (_MockOracleReceiver *MockOracleReceiverSession) ReceiveOracleMessage(sourceChainID [32]byte, sourceType string, sourceAddress string, nonce uint64, payload []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.ReceiveOracleMessage(&_MockOracleReceiver.TransactOpts, sourceChainID, sourceType, sourceAddress, nonce, payload)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1c67934b.
//
// Solidity: function receiveOracleMessage(bytes32 sourceChainID, string sourceType, string sourceAddress, uint64 nonce, bytes payload) returns()
func (_MockOracleReceiver *MockOracleReceiverTransactorSession) ReceiveOracleMessage(sourceChainID [32]byte, sourceType string, sourceAddress string, nonce uint64, payload []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.ReceiveOracleMessage(&_MockOracleReceiver.TransactOpts, sourceChainID, sourceType, sourceAddress, nonce, payload)
}
