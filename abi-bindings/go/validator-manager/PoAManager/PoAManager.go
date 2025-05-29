// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poamanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoAManagerMetaData contains all meta data concerning the PoAManager contract.
var PoAManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"POA_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"newWeight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferUnderlyingValidatorManagerOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610d4a8061001c5f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80639681d9401161006e5780639681d940146101795780639cb7624e1461018c578063a3a65e481461019f578063b6e6a2ca146101b2578063ce161f14146101c5578063f2fde38b146101f6575f80fd5b80632fabc10a146100b5578063485cc955146100dc57806364accaf4146100f15780636610966914610104578063715018a6146101375780638da5cb5b1461013f575b5f80fd5b6100c95f80516020610cf583398151915281565b6040519081526020015b60405180910390f35b6100ef6100ea366004610853565b610209565b005b6100ef6100ff366004610884565b610319565b6101176101123660046108b2565b61038d565b6040805167ffffffffffffffff90931683526020830191909152016100d3565b6100ef61042d565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020016100d3565b6100c96101873660046108f3565b610440565b6100c961019a366004610aab565b6104c7565b6100c96101ad3660046108f3565b610560565b6100ef6101c0366004610b65565b6105a4565b6101d86101d33660046108f3565b6105ed565b6040805192835267ffffffffffffffff9091166020830152016100d3565b6100ef610204366004610884565b610678565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f8115801561024e5750825b90505f8267ffffffffffffffff16600114801561026a5750303b155b905081158015610278575080155b156102965760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156102c057845460ff60401b1916600160401b1785555b6102ca87876106b7565b831561031057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6103216106d5565b5f80516020610cf5833981519152805460405163f2fde38b60e01b81526001600160a01b0384811660048301529091169063f2fde38b906024015b5f604051808303815f87803b158015610373575f80fd5b505af1158015610385573d5f803e3d5ffd5b505050505050565b5f806103976106d5565b5f5f80516020610cf58339815191528054604051636610966960e01b81526004810188905267ffffffffffffffff871660248201529192506001600160a01b03169063661096699060440160408051808303815f875af11580156103fd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104219190610b7c565b92509250509250929050565b6104356106d5565b61043e5f610730565b565b5f80516020610cf5833981519152805460405163025a076560e61b815263ffffffff841660048201525f92916001600160a01b031690639681d940906024015b6020604051808303815f875af115801561049c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c09190610ba8565b9392505050565b5f6104d06106d5565b5f80516020610cf58339815191528054604051634e5bb12760e11b81526001600160a01b0390911690639cb7624e90610515908a908a908a908a908a90600401610c68565b6020604051808303815f875af1158015610531573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105559190610ba8565b979650505050505050565b5f80516020610cf58339815191528054604051631474cbc960e31b815263ffffffff841660048201525f92916001600160a01b03169063a3a65e4890602401610480565b6105ac6106d5565b5f80516020610cf58339815191528054604051635b73516560e11b8152600481018490526001600160a01b039091169063b6e6a2ca9060240161035c565b50565b5f80805f80516020610cf5833981519152805460405163338587c560e21b815263ffffffff871660048201529192506001600160a01b03169063ce161f149060240160408051808303815f875af115801561064a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061066e9190610cd1565b9250925050915091565b6106806106d5565b6001600160a01b0381166106ae57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6105ea81610730565b6106bf6107a0565b6106c8826107e9565b6106d1816107fa565b5050565b336107077f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461043e5760405163118cdaa760e01b81523360048201526024016106a5565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661043e57604051631afcd79f60e31b815260040160405180910390fd5b6107f16107a0565b6105ea81610830565b6108026107a0565b5f80516020610cf583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6106806107a0565b80356001600160a01b038116811461084e575f80fd5b919050565b5f8060408385031215610864575f80fd5b61086d83610838565b915061087b60208401610838565b90509250929050565b5f60208284031215610894575f80fd5b6104c082610838565b67ffffffffffffffff811681146105ea575f80fd5b5f80604083850312156108c3575f80fd5b8235915060208301356108d58161089d565b809150509250929050565b803563ffffffff8116811461084e575f80fd5b5f60208284031215610903575f80fd5b6104c0826108e0565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156109435761094361090c565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156109725761097261090c565b604052919050565b5f82601f830112610989575f80fd5b813567ffffffffffffffff8111156109a3576109a361090c565b6109b6601f8201601f1916602001610949565b8181528460208386010111156109ca575f80fd5b816020850160208301375f918101602001919091529392505050565b5f604082840312156109f6575f80fd5b6109fe610920565b9050610a09826108e0565b815260208083013567ffffffffffffffff80821115610a26575f80fd5b818501915085601f830112610a39575f80fd5b813581811115610a4b57610a4b61090c565b8060051b9150610a5c848301610949565b8181529183018401918481019088841115610a75575f80fd5b938501935b83851015610a9a57610a8b85610838565b82529385019390850190610a7a565b808688015250505050505092915050565b5f805f805f60a08688031215610abf575f80fd5b853567ffffffffffffffff80821115610ad6575f80fd5b610ae289838a0161097a565b96506020880135915080821115610af7575f80fd5b610b0389838a0161097a565b95506040880135915080821115610b18575f80fd5b610b2489838a016109e6565b94506060880135915080821115610b39575f80fd5b50610b46888289016109e6565b9250506080860135610b578161089d565b809150509295509295909350565b5f60208284031215610b75575f80fd5b5035919050565b5f8060408385031215610b8d575f80fd5b8251610b988161089d565b6020939093015192949293505050565b5f60208284031215610bb8575f80fd5b5051919050565b5f81518084525f5b81811015610be357602081850181015186830182015201610bc7565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015610c5d5784516001600160a01b03168252938301936001929092019190830190610c34565b509695505050505050565b60a081525f610c7a60a0830188610bbf565b8281036020840152610c8c8188610bbf565b90508281036040840152610ca08187610c02565b90508281036060840152610cb48186610c02565b91505067ffffffffffffffff831660808301529695505050505050565b5f8060408385031215610ce2575f80fd5b8251915060208301516108d58161089d56fe8e2427ab32c2585abb2a107c76f30b8d77c153bac188f081d4c40ff3fcf13200a264697066735822122007634c4461b16cc05bca5eadc115ecc728ea8c56ca9ece523eba652a6794c82064736f6c63430008190033",
}

// PoAManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoAManagerMetaData.ABI instead.
var PoAManagerABI = PoAManagerMetaData.ABI

// PoAManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoAManagerMetaData.Bin instead.
var PoAManagerBin = PoAManagerMetaData.Bin

// DeployPoAManager deploys a new Ethereum contract, binding an instance of PoAManager to it.
func DeployPoAManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoAManager, error) {
	parsed, err := PoAManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoAManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoAManager{PoAManagerCaller: PoAManagerCaller{contract: contract}, PoAManagerTransactor: PoAManagerTransactor{contract: contract}, PoAManagerFilterer: PoAManagerFilterer{contract: contract}}, nil
}

// PoAManager is an auto generated Go binding around an Ethereum contract.
type PoAManager struct {
	PoAManagerCaller     // Read-only binding to the contract
	PoAManagerTransactor // Write-only binding to the contract
	PoAManagerFilterer   // Log filterer for contract events
}

// PoAManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoAManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoAManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoAManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoAManagerSession struct {
	Contract     *PoAManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoAManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoAManagerCallerSession struct {
	Contract *PoAManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PoAManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoAManagerTransactorSession struct {
	Contract     *PoAManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PoAManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoAManagerRaw struct {
	Contract *PoAManager // Generic contract binding to access the raw methods on
}

// PoAManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoAManagerCallerRaw struct {
	Contract *PoAManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoAManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoAManagerTransactorRaw struct {
	Contract *PoAManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoAManager creates a new instance of PoAManager, bound to a specific deployed contract.
func NewPoAManager(address common.Address, backend bind.ContractBackend) (*PoAManager, error) {
	contract, err := bindPoAManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoAManager{PoAManagerCaller: PoAManagerCaller{contract: contract}, PoAManagerTransactor: PoAManagerTransactor{contract: contract}, PoAManagerFilterer: PoAManagerFilterer{contract: contract}}, nil
}

// NewPoAManagerCaller creates a new read-only instance of PoAManager, bound to a specific deployed contract.
func NewPoAManagerCaller(address common.Address, caller bind.ContractCaller) (*PoAManagerCaller, error) {
	contract, err := bindPoAManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoAManagerCaller{contract: contract}, nil
}

// NewPoAManagerTransactor creates a new write-only instance of PoAManager, bound to a specific deployed contract.
func NewPoAManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoAManagerTransactor, error) {
	contract, err := bindPoAManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoAManagerTransactor{contract: contract}, nil
}

// NewPoAManagerFilterer creates a new log filterer instance of PoAManager, bound to a specific deployed contract.
func NewPoAManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoAManagerFilterer, error) {
	contract, err := bindPoAManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoAManagerFilterer{contract: contract}, nil
}

// bindPoAManager binds a generic wrapper to an already deployed contract.
func bindPoAManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoAManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAManager *PoAManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAManager.Contract.PoAManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAManager *PoAManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAManager.Contract.PoAManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAManager *PoAManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAManager.Contract.PoAManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAManager *PoAManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAManager *PoAManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAManager *PoAManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAManager.Contract.contract.Transact(opts, method, params...)
}

// POAMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x2fabc10a.
//
// Solidity: function POA_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoAManager *PoAManagerCaller) POAMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoAManager.contract.Call(opts, &out, "POA_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POAMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x2fabc10a.
//
// Solidity: function POA_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoAManager *PoAManagerSession) POAMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoAManager.Contract.POAMANAGERSTORAGELOCATION(&_PoAManager.CallOpts)
}

// POAMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x2fabc10a.
//
// Solidity: function POA_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoAManager *PoAManagerCallerSession) POAMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoAManager.Contract.POAMANAGERSTORAGELOCATION(&_PoAManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAManager *PoAManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAManager *PoAManagerSession) Owner() (common.Address, error) {
	return _PoAManager.Contract.Owner(&_PoAManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAManager *PoAManagerCallerSession) Owner() (common.Address, error) {
	return _PoAManager.Contract.Owner(&_PoAManager.CallOpts)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorRegistration(&_PoAManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorRegistration(&_PoAManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorRemoval(&_PoAManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_PoAManager *PoAManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorRemoval(&_PoAManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_PoAManager *PoAManagerTransactor) CompleteValidatorWeightUpdate(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "completeValidatorWeightUpdate", messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_PoAManager *PoAManagerSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorWeightUpdate(&_PoAManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32, uint64)
func (_PoAManager *PoAManagerTransactorSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _PoAManager.Contract.CompleteValidatorWeightUpdate(&_PoAManager.TransactOpts, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address validatorManager) returns()
func (_PoAManager *PoAManagerTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, validatorManager common.Address) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "initialize", owner, validatorManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address validatorManager) returns()
func (_PoAManager *PoAManagerSession) Initialize(owner common.Address, validatorManager common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.Initialize(&_PoAManager.TransactOpts, owner, validatorManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address validatorManager) returns()
func (_PoAManager *PoAManagerTransactorSession) Initialize(owner common.Address, validatorManager common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.Initialize(&_PoAManager.TransactOpts, owner, validatorManager)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x9cb7624e.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint64 weight) returns(bytes32)
func (_PoAManager *PoAManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, weight uint64) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x9cb7624e.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint64 weight) returns(bytes32)
func (_PoAManager *PoAManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, weight uint64) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorRegistration(&_PoAManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x9cb7624e.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint64 weight) returns(bytes32)
func (_PoAManager *PoAManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, weight uint64) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorRegistration(&_PoAManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_PoAManager *PoAManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_PoAManager *PoAManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorRemoval(&_PoAManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_PoAManager *PoAManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorRemoval(&_PoAManager.TransactOpts, validationID)
}

// InitiateValidatorWeightUpdate is a paid mutator transaction binding the contract method 0x66109669.
//
// Solidity: function initiateValidatorWeightUpdate(bytes32 validationID, uint64 newWeight) returns(uint64, bytes32)
func (_PoAManager *PoAManagerTransactor) InitiateValidatorWeightUpdate(opts *bind.TransactOpts, validationID [32]byte, newWeight uint64) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "initiateValidatorWeightUpdate", validationID, newWeight)
}

// InitiateValidatorWeightUpdate is a paid mutator transaction binding the contract method 0x66109669.
//
// Solidity: function initiateValidatorWeightUpdate(bytes32 validationID, uint64 newWeight) returns(uint64, bytes32)
func (_PoAManager *PoAManagerSession) InitiateValidatorWeightUpdate(validationID [32]byte, newWeight uint64) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorWeightUpdate(&_PoAManager.TransactOpts, validationID, newWeight)
}

// InitiateValidatorWeightUpdate is a paid mutator transaction binding the contract method 0x66109669.
//
// Solidity: function initiateValidatorWeightUpdate(bytes32 validationID, uint64 newWeight) returns(uint64, bytes32)
func (_PoAManager *PoAManagerTransactorSession) InitiateValidatorWeightUpdate(validationID [32]byte, newWeight uint64) (*types.Transaction, error) {
	return _PoAManager.Contract.InitiateValidatorWeightUpdate(&_PoAManager.TransactOpts, validationID, newWeight)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAManager *PoAManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAManager *PoAManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAManager.Contract.RenounceOwnership(&_PoAManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAManager *PoAManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAManager.Contract.RenounceOwnership(&_PoAManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferOwnership(&_PoAManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferOwnership(&_PoAManager.TransactOpts, newOwner)
}

// TransferUnderlyingValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x64accaf4.
//
// Solidity: function transferUnderlyingValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactor) TransferUnderlyingValidatorManagerOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "transferUnderlyingValidatorManagerOwnership", newOwner)
}

// TransferUnderlyingValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x64accaf4.
//
// Solidity: function transferUnderlyingValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerSession) TransferUnderlyingValidatorManagerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferUnderlyingValidatorManagerOwnership(&_PoAManager.TransactOpts, newOwner)
}

// TransferUnderlyingValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x64accaf4.
//
// Solidity: function transferUnderlyingValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactorSession) TransferUnderlyingValidatorManagerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferUnderlyingValidatorManagerOwnership(&_PoAManager.TransactOpts, newOwner)
}

// PoAManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoAManager contract.
type PoAManagerInitializedIterator struct {
	Event *PoAManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAManagerInitialized)
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
		it.Event = new(PoAManagerInitialized)
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
func (it *PoAManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAManagerInitialized represents a Initialized event raised by the PoAManager contract.
type PoAManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAManager *PoAManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoAManagerInitializedIterator, error) {

	logs, sub, err := _PoAManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoAManagerInitializedIterator{contract: _PoAManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAManager *PoAManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoAManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoAManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAManagerInitialized)
				if err := _PoAManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAManager *PoAManagerFilterer) ParseInitialized(log types.Log) (*PoAManagerInitialized, error) {
	event := new(PoAManagerInitialized)
	if err := _PoAManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoAManager contract.
type PoAManagerOwnershipTransferredIterator struct {
	Event *PoAManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoAManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAManagerOwnershipTransferred)
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
		it.Event = new(PoAManagerOwnershipTransferred)
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
func (it *PoAManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PoAManager contract.
type PoAManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAManager *PoAManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoAManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoAManagerOwnershipTransferredIterator{contract: _PoAManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAManager *PoAManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoAManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAManagerOwnershipTransferred)
				if err := _PoAManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAManager *PoAManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoAManagerOwnershipTransferred, error) {
	event := new(PoAManagerOwnershipTransferred)
	if err := _PoAManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
