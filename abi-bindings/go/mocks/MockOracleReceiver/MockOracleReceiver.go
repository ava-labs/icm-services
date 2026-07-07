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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"teleporterMessenger_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastPayload\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceBlockHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSourceType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"teleporterMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"OnlyTeleporterMessenger\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516107c33803806107c383398101604081905261002e91610066565b6001600160a01b0381166100555760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052610093565b5f60208284031215610076575f5ffd5b81516001600160a01b038116811461008c575f5ffd5b9392505050565b60805161070a6100b95f395f8181610114015281816102270152610269015261070a5ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c806352631ab41161006357806352631ab4146100d45780637ba85592146101065780639b3e58031461010f578063b15f20941461014e578063c868efaa14610161575f5ffd5b80630a4e00bf146100945780630d768691146100af57806344c19139146100c45780634cc2aa3c146100cc575b5f5ffd5b61009c5f5481565b6040519081526020015b60405180910390f35b6100b7610176565b6040516100a69190610373565b6100b7610202565b6100b761020f565b6003546100ee90600160401b90046001600160401b031681565b6040516001600160401b0390911681526020016100a6565b61009c60055481565b6101367f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100a6565b6003546100ee906001600160401b031681565b61017461016f36600461038c565b61021c565b005b600280546101839061041a565b80601f01602080910402602001604051908101604052809291908181526020018280546101af9061041a565b80156101fa5780601f106101d1576101008083540402835291602001916101fa565b820191905f5260205f20905b8154815290600101906020018083116101dd57829003601f168201915b505050505081565b600180546101839061041a565b600480546101839061041a565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461029a576040516370931f1360e01b81523360048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602482015260440160405180910390fd5b5f808080806102ab86880188610514565b5f8e905593985091965094509250905060016102c7868261061f565b5060026102d4858261061f565b50600380546001600160401b03848116600160401b026fffffffffffffffffffffffffffffffff19909216908616171790556004610312828261061f565b5060055f8154610321906106d9565b90915550505050505050505050565b5f81518084525f5b8181101561035457602081850181015186830182015201610338565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6103856020830184610330565b9392505050565b5f5f5f5f6060858703121561039f575f5ffd5b8435935060208501356001600160a01b03811681146103bc575f5ffd5b925060408501356001600160401b038111156103d6575f5ffd5b8501601f810187136103e6575f5ffd5b80356001600160401b038111156103fb575f5ffd5b87602082840101111561040c575f5ffd5b949793965060200194505050565b600181811c9082168061042e57607f821691505b60208210810361044c57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f6001600160401b0384111561047f5761047f610452565b50604051601f19601f85018116603f011681018181106001600160401b03821117156104ad576104ad610452565b6040528381529050808284018510156104c4575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f8301126104ea575f5ffd5b61038583833560208501610466565b80356001600160401b038116811461050f575f5ffd5b919050565b5f5f5f5f5f60a08688031215610528575f5ffd5b85356001600160401b0381111561053d575f5ffd5b610549888289016104db565b95505060208601356001600160401b03811115610564575f5ffd5b610570888289016104db565b94505061057f604087016104f9565b925061058d606087016104f9565b915060808601356001600160401b038111156105a7575f5ffd5b8601601f810188136105b7575f5ffd5b6105c688823560208401610466565b9150509295509295909350565b601f82111561061a57805f5260205f20601f840160051c810160208510156105f85750805b601f840160051c820191505b81811015610617575f8155600101610604565b50505b505050565b81516001600160401b0381111561063857610638610452565b61064c81610646845461041a565b846105d3565b6020601f82116001811461067e575f83156106675750848201515b5f19600385901b1c1916600184901b178455610617565b5f84815260208120601f198516915b828110156106ad578785015182556020948501946001909201910161068d565b50848210156106ca57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f600182016106f657634e487b7160e01b5f52601160045260245ffd5b506001019056fea164736f6c634300081e000a",
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

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverCaller) LastSourceBlockHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockOracleReceiver.contract.Call(opts, &out, "lastSourceBlockHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverSession) LastSourceBlockHeight() (uint64, error) {
	return _MockOracleReceiver.Contract.LastSourceBlockHeight(&_MockOracleReceiver.CallOpts)
}

// LastSourceBlockHeight is a free data retrieval call binding the contract method 0xb15f2094.
//
// Solidity: function lastSourceBlockHeight() view returns(uint64)
func (_MockOracleReceiver *MockOracleReceiverCallerSession) LastSourceBlockHeight() (uint64, error) {
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
