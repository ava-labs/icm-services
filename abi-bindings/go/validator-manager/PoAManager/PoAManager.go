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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIValidatorManagerExternalOwnable\",\"name\":\"validatorManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"newWeight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferValidatorManagerOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610b60380380610b6083398101604081905261002e916100f2565b816001600160a01b03811661005c57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100658161008c565b50600180546001600160a01b0319166001600160a01b03929092169190911790555061012a565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146100ef575f80fd5b50565b5f8060408385031215610103575f80fd5b825161010e816100db565b602084015190925061011f816100db565b809150509250929050565b610a29806101375f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c80639cb7624e116100635780639cb7624e1461012f578063a3a65e4814610142578063b6e6a2ca14610155578063ce161f1414610168578063f2fde38b14610199575f80fd5b8063661096691461009f578063715018a6146100d757806389f9f85b146100e15780638da5cb5b146100f45780639681d9401461010e575b5f80fd5b6100b26100ad366004610576565b6101ac565b6040805167ffffffffffffffff90931683526020830191909152015b60405180910390f35b6100df61023c565b005b6100df6100ef3660046105bf565b61024f565b5f546040516001600160a01b0390911681526020016100ce565b61012161011c3660046105f2565b6102b5565b6040519081526020016100ce565b61012161013d3660046107aa565b61032d565b6101216101503660046105f2565b6103b8565b6100df610163366004610864565b6103ee565b61017b6101763660046105f2565b61042a565b6040805192835267ffffffffffffffff9091166020830152016100ce565b6100df6101a73660046105bf565b6104a7565b5f806101b66104e6565b600154604051636610966960e01b81526004810186905267ffffffffffffffff851660248201526001600160a01b039091169063661096699060440160408051808303815f875af115801561020d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610231919061087b565b915091509250929050565b6102446104e6565b61024d5f610512565b565b6102576104e6565b60015460405163f2fde38b60e01b81526001600160a01b0383811660048301529091169063f2fde38b906024015b5f604051808303815f87803b15801561029c575f80fd5b505af11580156102ae573d5f803e3d5ffd5b5050505050565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af1158015610303573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032791906108a7565b92915050565b5f6103366104e6565b600154604051634e5bb12760e11b81526001600160a01b0390911690639cb7624e9061036e9089908990899089908990600401610967565b6020604051808303815f875af115801561038a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ae91906108a7565b9695505050505050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016102e7565b6103f66104e6565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca90602401610285565b50565b60015460405163338587c560e21b815263ffffffff831660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af115801561047a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049e91906109d0565b91509150915091565b6104af6104e6565b6001600160a01b0381166104dd57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61042781610512565b5f546001600160a01b0316331461024d5760405163118cdaa760e01b81523360048201526024016104d4565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b67ffffffffffffffff81168114610427575f80fd5b5f8060408385031215610587575f80fd5b82359150602083013561059981610561565b809150509250929050565b80356001600160a01b03811681146105ba575f80fd5b919050565b5f602082840312156105cf575f80fd5b6105d8826105a4565b9392505050565b803563ffffffff811681146105ba575f80fd5b5f60208284031215610602575f80fd5b6105d8826105df565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156106425761064261060b565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156106715761067161060b565b604052919050565b5f82601f830112610688575f80fd5b813567ffffffffffffffff8111156106a2576106a261060b565b6106b5601f8201601f1916602001610648565b8181528460208386010111156106c9575f80fd5b816020850160208301375f918101602001919091529392505050565b5f604082840312156106f5575f80fd5b6106fd61061f565b9050610708826105df565b815260208083013567ffffffffffffffff80821115610725575f80fd5b818501915085601f830112610738575f80fd5b81358181111561074a5761074a61060b565b8060051b915061075b848301610648565b8181529183018401918481019088841115610774575f80fd5b938501935b838510156107995761078a856105a4565b82529385019390850190610779565b808688015250505050505092915050565b5f805f805f60a086880312156107be575f80fd5b853567ffffffffffffffff808211156107d5575f80fd5b6107e189838a01610679565b965060208801359150808211156107f6575f80fd5b61080289838a01610679565b95506040880135915080821115610817575f80fd5b61082389838a016106e5565b94506060880135915080821115610838575f80fd5b50610845888289016106e5565b925050608086013561085681610561565b809150509295509295909350565b5f60208284031215610874575f80fd5b5035919050565b5f806040838503121561088c575f80fd5b825161089781610561565b6020939093015192949293505050565b5f602082840312156108b7575f80fd5b5051919050565b5f81518084525f5b818110156108e2576020818501810151868301820152016108c6565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b8083101561095c5784516001600160a01b03168252938301936001929092019190830190610933565b509695505050505050565b60a081525f61097960a08301886108be565b828103602084015261098b81886108be565b9050828103604084015261099f8187610901565b905082810360608401526109b38186610901565b91505067ffffffffffffffff831660808301529695505050505050565b5f80604083850312156109e1575f80fd5b8251915060208301516105998161056156fea2646970667358221220d508e71c3692c1792236951288b4596893cca2e0dd3d00055398f0f43e3bdaa864736f6c63430008190033",
}

// PoAManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoAManagerMetaData.ABI instead.
var PoAManagerABI = PoAManagerMetaData.ABI

// PoAManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoAManagerMetaData.Bin instead.
var PoAManagerBin = PoAManagerMetaData.Bin

// DeployPoAManager deploys a new Ethereum contract, binding an instance of PoAManager to it.
func DeployPoAManager(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, validatorManager common.Address) (common.Address, *types.Transaction, *PoAManager, error) {
	parsed, err := PoAManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoAManagerBin), backend, owner, validatorManager)
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

// TransferValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x89f9f85b.
//
// Solidity: function transferValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactor) TransferValidatorManagerOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.contract.Transact(opts, "transferValidatorManagerOwnership", newOwner)
}

// TransferValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x89f9f85b.
//
// Solidity: function transferValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerSession) TransferValidatorManagerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferValidatorManagerOwnership(&_PoAManager.TransactOpts, newOwner)
}

// TransferValidatorManagerOwnership is a paid mutator transaction binding the contract method 0x89f9f85b.
//
// Solidity: function transferValidatorManagerOwnership(address newOwner) returns()
func (_PoAManager *PoAManagerTransactorSession) TransferValidatorManagerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAManager.Contract.TransferValidatorManagerOwnership(&_PoAManager.TransactOpts, newOwner)
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
