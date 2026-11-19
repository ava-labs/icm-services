// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poamanager

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

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoAManagerMetaData contains all meta data concerning the PoAManager contract.
var PoAManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIValidatorManagerExternalOwnable\",\"name\":\"validatorManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"newWeight\",\"type\":\"uint64\"}],\"name\":\"initiateValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferValidatorManagerOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a03461010857601f610ada38819003918201601f19168301916001600160401b0383118484101761010c5780849260409485528339810103126101085780516001600160a01b038116919082900361010857602001516001600160a01b03811681036101085781156100f5575f80546001600160a01b031981168417825560405193916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36080526109b99081610121823960805181818161014f01528181610212015281816102cb01528181610475015281816105040152818161059001526106d00152f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f5f3560e01c80636610966914610669578063715018a61461061257806389f9f85b1461056c5780638da5cb5b146105455780639681d940146104c65780639cb7624e14610350578063a3a65e481461028d578063b6e6a2ca146101f5578063ce161f14146101105763f2fde38b14610088575f80fd5b3461010d57602036600319011261010d576100a1610763565b6100a9610986565b6001600160a01b031680156100f95781546001600160a01b03198116821783556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b631e4fbdf760e01b82526004829052602482fd5b80fd5b503461010d57602036600319011261010d5761012a610779565b6040805163338587c560e21b815263ffffffff909216600483015290919082602481847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19182156101e8578180936101a4575b5050604080519182526001600160401b03929092166020820152f35b915091506040823d6040116101e0575b816101c16040938361078c565b8101031261010d57506101d86020825192016108df565b905f80610188565b3d91506101b4565b50604051903d90823e3d90fd5b503461010d57602036600319011261010d5761020f610986565b807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561028a57818091602460405180948193635b73516560e11b835260043560048401525af1801561027f5761026e5750f35b816102789161078c565b61010d5780f35b6040513d84823e3d90fd5b50fd5b503461010d57602036600319011261010d576102a7610779565b604051631474cbc960e31b815263ffffffff909116600482015290602082602481847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610344579061030d575b602090604051908152f35b506020813d60201161033c575b816103276020938361078c565b810103126103385760209051610302565b5f80fd5b3d915061031a565b604051903d90823e3d90fd5b503461010d5760a036600319011261010d576004356001600160401b0381116104c2576103819036906004016107c1565b906024356001600160401b0381116104c2576103a19036906004016107c1565b916044356001600160401b0381116104be576103c1903690600401610816565b6064356001600160401b0381116104ba576103e0903690600401610816565b608435916001600160401b0383168093036104b65791610468602094926104566104329561040c610986565b6104446040519a8b988998634e5bb12760e11b8a5260a060048b015260a48a01906108f3565b8881036003190160248a0152906108f3565b86810360031901604488015290610932565b84810360031901606486015290610932565b60848301919091520381847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610344579061030d57602090604051908152f35b8480fd5b8380fd5b8280fd5b5080fd5b503461010d57602036600319011261010d576104e0610779565b60405163025a076560e61b815263ffffffff909116600482015290602082602481847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610344579061030d57602090604051908152f35b503461010d578060031936011261010d57546040516001600160a01b039091168152602090f35b503461033857602036600319011261033857610586610763565b61058e610986565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156103385760405163f2fde38b60e01b81526001600160a01b039091166004820152905f908290602490829084905af18015610607576105f9575080f35b61060591505f9061078c565b005b6040513d5f823e3d90fd5b34610338575f3660031901126103385761062a610986565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b34610338576040366003190112610338576024356001600160401b03811681036103385760406106cb9161069b610986565b8151636610966960e01b815260048035908201526001600160401b03909116602482015291829081906044820190565b03815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015610607575f905f90610722575b604092506001600160401b038351921682526020820152f35b50506040813d60401161075b575b8161073d6040938361078c565b81010312610338578060206107536040936108df565b910151610709565b3d9150610730565b600435906001600160a01b038216820361033857565b6004359063ffffffff8216820361033857565b90601f801991011681019081106001600160401b038211176107ad57604052565b634e487b7160e01b5f52604160045260245ffd5b81601f82011215610338578035906001600160401b0382116107ad57604051926107f5601f8401601f19166020018561078c565b8284526020838301011161033857815f926020809301838601378301015290565b919060408382031261033857604051604081018181106001600160401b038211176107ad576040528093803563ffffffff811681036103385782526020810135906001600160401b038211610338570182601f82011215610338578035926001600160401b0384116107ad578360051b9160405194610898602085018761078c565b855260208086019382010191821161033857602001915b8183106108bf5750505060200152565b82356001600160a01b0381168103610338578152602092830192016108af565b51906001600160401b038216820361033857565b91908251928382525f5b84811061091d575050825f602080949584010152601f8019910116010190565b806020809284010151828286010152016108fd565b6020606081604085019363ffffffff81511686520151936040838201528451809452019201905f5b8181106109675750505090565b82516001600160a01b031684526020938401939092019160010161095a565b5f546001600160a01b0316330361099957565b63118cdaa760e01b5f523360045260245ffdfea164736f6c634300081e000a",
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

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
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
