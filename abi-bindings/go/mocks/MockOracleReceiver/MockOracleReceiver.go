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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessenger_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceBlockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"OracleMessageReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"lastNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSourceAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSourceBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSourceChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSourceType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161080238038061080283398101604081905261002e91610066565b6001600160a01b0381166100555760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052610093565b5f60208284031215610076575f5ffd5b81516001600160a01b038116811461008c575f5ffd5b9392505050565b60805161074a6100b85f395f818160eb015281816101f40152610236015261074a5ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c806352631ab41161006357806352631ab4146100d45780637ba85592146100dd5780639b3e5803146100e6578063b15f209414610125578063c868efaa1461012e575f5ffd5b80630a4e00bf146100945780630d768691146100af57806344c19139146100c45780634cc2aa3c146100cc575b5f5ffd5b61009c5f5481565b6040519081526020015b60405180910390f35b6100b7610143565b6040516100a6919061035b565b6100b76101cf565b6100b76101dc565b61009c60045481565b61009c60065481565b61010d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a6565b61009c60035481565b61014161013c366004610374565b6101e9565b005b6002805461015090610404565b80601f016020809104026020016040519081016040528092919081815260200182805461017c90610404565b80156101c75780601f1061019e576101008083540402835291602001916101c7565b820191905f5260205f20905b8154815290600101906020018083116101aa57829003601f168201915b505050505081565b6001805461015090610404565b6005805461015090610404565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610267576040516385faaab560e01b81523360048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602482015260440160405180910390fd5b5f80808080610278868801886104e5565b5f8e9055939850919650945092509050600161029486826105e5565b5060026102a185826105e5565b506003839055600482905560056102b882826105e5565b5060065f81546102c7906106a0565b9091555060405189907f96175bc7a2d20af3c200044ed8440af964a33879077ba89e6e5a68fa3b2eb2219061030590889088908890889088906106c4565b60405180910390a2505050505050505050565b5f81518084525f5b8181101561033c57602081850181015186830182015201610320565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f61036d6020830184610318565b9392505050565b5f5f5f5f60608587031215610387575f5ffd5b8435935060208501356001600160a01b03811681146103a4575f5ffd5b9250604085013567ffffffffffffffff8111156103bf575f5ffd5b8501601f810187136103cf575f5ffd5b803567ffffffffffffffff8111156103e5575f5ffd5b8760208284010111156103f6575f5ffd5b949793965060200194505050565b600181811c9082168061041857607f821691505b60208210810361043657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f67ffffffffffffffff84111561046a5761046a61043c565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff821117156104995761049961043c565b6040528381529050808284018510156104b0575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f8301126104d6575f5ffd5b61036d83833560208501610450565b5f5f5f5f5f60a086880312156104f9575f5ffd5b853567ffffffffffffffff81111561050f575f5ffd5b61051b888289016104c7565b955050602086013567ffffffffffffffff811115610537575f5ffd5b610543888289016104c7565b9450506040860135925060608601359150608086013567ffffffffffffffff81111561056d575f5ffd5b8601601f8101881361057d575f5ffd5b61058c88823560208401610450565b9150509295509295909350565b601f8211156105e057805f5260205f20601f840160051c810160208510156105be5750805b601f840160051c820191505b818110156105dd575f81556001016105ca565b50505b505050565b815167ffffffffffffffff8111156105ff576105ff61043c565b6106138161060d8454610404565b84610599565b6020601f821160018114610645575f831561062e5750848201515b5f19600385901b1c1916600184901b1784556105dd565b5f84815260208120601f198516915b828110156106745787850151825560209485019460019092019101610654565b508482101561069157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f600182016106bd57634e487b7160e01b5f52601160045260245ffd5b5060010190565b60a081525f6106d660a0830188610318565b82810360208401526106e88188610318565b905085604084015284606084015282810360808401526107088185610318565b9897505050505050505056fea26469706673582212209dd72fa44660990fd37d0522235e1b625b37ebb5e6b89b6510f3754db8db4eed64736f6c634300081e0033",
}

// MockOracleReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockOracleReceiverMetaData.ABI instead.
var MockOracleReceiverABI = MockOracleReceiverMetaData.ABI

// MockOracleReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockOracleReceiverMetaData.Bin instead.
var MockOracleReceiverBin = MockOracleReceiverMetaData.Bin

// DeployMockOracleReceiver deploys a new Ethereum contract, binding an instance of MockOracleReceiver to it.
func DeployMockOracleReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterMessenger_ common.Address) (common.Address, *types.Transaction, *MockOracleReceiver, error) {
	parsed, err := MockOracleReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockOracleReceiverBin), backend, teleporterMessenger_)
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
// Solidity: function lastNonce() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastNonce is a free data retrieval call binding the contract method 0x52631ab4.
//
// Solidity: function lastNonce() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverSession) LastNonce() (*big.Int, error) {
	return _MockOracleReceiver.Contract.LastNonce(&_MockOracleReceiver.CallOpts)
}

// LastNonce is a free data retrieval call binding the contract method 0x52631ab4.
//
// Solidity: function lastNonce() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastNonce() (*big.Int, error) {
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

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastSourceBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastSourceBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverSession) LastSourceBlockHeight() (*big.Int, error) {
	return _MockOracleReceiver.Contract.LastSourceBlockHeight(&_MockOracleReceiver.CallOpts)
}

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint256)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastSourceBlockHeight() (*big.Int, error) {
	return _MockOracleReceiver.Contract.LastSourceBlockHeight(&_MockOracleReceiver.CallOpts)
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

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverSession) TeleporterMessenger() (common.Address, error) {
	return _MockOracleReceiver.Contract.TeleporterMessenger(&_MockOracleReceiver.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) TeleporterMessenger() (common.Address, error) {
	return _MockOracleReceiver.Contract.TeleporterMessenger(&_MockOracleReceiver.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address , bytes message) returns()
func (_MockOracleReceiver *MockOracleReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address , bytes message) returns()
func (_MockOracleReceiver *MockOracleReceiverSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.ReceiveTeleporterMessage(&_MockOracleReceiver.TransactOpts, sourceBlockchainID, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address , bytes message) returns()
func (_MockOracleReceiver *MockOracleReceiverTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _MockOracleReceiver.Contract.ReceiveTeleporterMessage(&_MockOracleReceiver.TransactOpts, sourceBlockchainID, arg1, message)
}

// MockOracleReceiverOracleMessageReceivedIterator is returned from FilterOracleMessageReceived and is used to iterate over the raw logs and unpacked data for OracleMessageReceived events raised by the MockOracleReceiver contract.
type MockOracleReceiverOracleMessageReceivedIterator struct {
	Event *MockOracleReceiverOracleMessageReceived // Event containing the contract specifics and raw log

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
func (it *MockOracleReceiverOracleMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOracleReceiverOracleMessageReceived)
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
		it.Event = new(MockOracleReceiverOracleMessageReceived)
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
func (it *MockOracleReceiverOracleMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOracleReceiverOracleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOracleReceiverOracleMessageReceived represents a OracleMessageReceived event raised by the MockOracleReceiver contract.
type MockOracleReceiverOracleMessageReceived struct {
	SourceBlockchainID [32]byte
	SourceType         string
	SourceAddress      string
	SourceBlockHeight  *big.Int
	Nonce              *big.Int
	Payload            []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOracleMessageReceived is a free log retrieval operation binding the contract event 0x96175bc7a2d20af3c200044ed8440af964a33879077ba89e6e5a68fa3b2eb221.
//
// Solidity: event OracleMessageReceived(bytes32 indexed sourceBlockchainID, string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload)
func (_MockOracleReceiver *MockOracleReceiverFilterer) FilterOracleMessageReceived(opts *bind.FilterOpts, sourceBlockchainID [][32]byte) (*MockOracleReceiverOracleMessageReceivedIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _MockOracleReceiver.contract.FilterLogs(opts, "OracleMessageReceived", sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &MockOracleReceiverOracleMessageReceivedIterator{contract: _MockOracleReceiver.contract, event: "OracleMessageReceived", logs: logs, sub: sub}, nil
}

// WatchOracleMessageReceived is a free log subscription operation binding the contract event 0x96175bc7a2d20af3c200044ed8440af964a33879077ba89e6e5a68fa3b2eb221.
//
// Solidity: event OracleMessageReceived(bytes32 indexed sourceBlockchainID, string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload)
func (_MockOracleReceiver *MockOracleReceiverFilterer) WatchOracleMessageReceived(opts *bind.WatchOpts, sink chan<- *MockOracleReceiverOracleMessageReceived, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _MockOracleReceiver.contract.WatchLogs(opts, "OracleMessageReceived", sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOracleReceiverOracleMessageReceived)
				if err := _MockOracleReceiver.contract.UnpackLog(event, "OracleMessageReceived", log); err != nil {
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

// ParseOracleMessageReceived is a log parse operation binding the contract event 0x96175bc7a2d20af3c200044ed8440af964a33879077ba89e6e5a68fa3b2eb221.
//
// Solidity: event OracleMessageReceived(bytes32 indexed sourceBlockchainID, string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload)
func (_MockOracleReceiver *MockOracleReceiverFilterer) ParseOracleMessageReceived(log types.Log) (*MockOracleReceiverOracleMessageReceived, error) {
	event := new(MockOracleReceiverOracleMessageReceived)
	if err := _MockOracleReceiver.contract.UnpackLog(event, "OracleMessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
