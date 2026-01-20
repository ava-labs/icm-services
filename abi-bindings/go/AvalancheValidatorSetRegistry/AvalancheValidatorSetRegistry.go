// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avalanchevalidatorsetregistry

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

// AvalancheValidatorSetRegistryMetaData contains all meta data concerning the AvalancheValidatorSetRegistry contract.
var AvalancheValidatorSetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorSetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorSetID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052348015600e575f5ffd5b50604051610149380380610149833981016040819052602b916038565b63ffffffff166080526060565b5f602082840312156047575f5ffd5b815163ffffffff811681146059575f5ffd5b9392505050565b60805160ce61007b5f395f8181603801526075015260ce5ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c806368531ed014603457806382366d05146073575b5f5ffd5b605a7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200160405180910390f35b7f0000000000000000000000000000000000000000000000000000000000000000605a56fea26469706673582212208a823d82f8a7fe65ce9a905b90ae16e3457dd5e682e11b0c726b4cfc7ff66d6f64736f6c634300081e0033",
}

// AvalancheValidatorSetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.ABI instead.
var AvalancheValidatorSetRegistryABI = AvalancheValidatorSetRegistryMetaData.ABI

// AvalancheValidatorSetRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.Bin instead.
var AvalancheValidatorSetRegistryBin = AvalancheValidatorSetRegistryMetaData.Bin

// DeployAvalancheValidatorSetRegistry deploys a new Ethereum contract, binding an instance of AvalancheValidatorSetRegistry to it.
func DeployAvalancheValidatorSetRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, avalancheNetworkID_ uint32) (common.Address, *types.Transaction, *AvalancheValidatorSetRegistry, error) {
	parsed, err := AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AvalancheValidatorSetRegistryBin), backend, avalancheNetworkID_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AvalancheValidatorSetRegistry{AvalancheValidatorSetRegistryCaller: AvalancheValidatorSetRegistryCaller{contract: contract}, AvalancheValidatorSetRegistryTransactor: AvalancheValidatorSetRegistryTransactor{contract: contract}, AvalancheValidatorSetRegistryFilterer: AvalancheValidatorSetRegistryFilterer{contract: contract}}, nil
}

// AvalancheValidatorSetRegistry is an auto generated Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistry struct {
	AvalancheValidatorSetRegistryCaller     // Read-only binding to the contract
	AvalancheValidatorSetRegistryTransactor // Write-only binding to the contract
	AvalancheValidatorSetRegistryFilterer   // Log filterer for contract events
}

// AvalancheValidatorSetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvalancheValidatorSetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvalancheValidatorSetRegistrySession struct {
	Contract     *AvalancheValidatorSetRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvalancheValidatorSetRegistryCallerSession struct {
	Contract *AvalancheValidatorSetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// AvalancheValidatorSetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvalancheValidatorSetRegistryTransactorSession struct {
	Contract     *AvalancheValidatorSetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryRaw struct {
	Contract *AvalancheValidatorSetRegistry // Generic contract binding to access the raw methods on
}

// AvalancheValidatorSetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryCallerRaw struct {
	Contract *AvalancheValidatorSetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AvalancheValidatorSetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryTransactorRaw struct {
	Contract *AvalancheValidatorSetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvalancheValidatorSetRegistry creates a new instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistry(address common.Address, backend bind.ContractBackend) (*AvalancheValidatorSetRegistry, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistry{AvalancheValidatorSetRegistryCaller: AvalancheValidatorSetRegistryCaller{contract: contract}, AvalancheValidatorSetRegistryTransactor: AvalancheValidatorSetRegistryTransactor{contract: contract}, AvalancheValidatorSetRegistryFilterer: AvalancheValidatorSetRegistryFilterer{contract: contract}}, nil
}

// NewAvalancheValidatorSetRegistryCaller creates a new read-only instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryCaller(address common.Address, caller bind.ContractCaller) (*AvalancheValidatorSetRegistryCaller, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryCaller{contract: contract}, nil
}

// NewAvalancheValidatorSetRegistryTransactor creates a new write-only instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AvalancheValidatorSetRegistryTransactor, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryTransactor{contract: contract}, nil
}

// NewAvalancheValidatorSetRegistryFilterer creates a new log filterer instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AvalancheValidatorSetRegistryFilterer, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryFilterer{contract: contract}, nil
}

// bindAvalancheValidatorSetRegistry binds a generic wrapper to an already deployed contract.
func bindAvalancheValidatorSetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.contract.Transact(opts, method, params...)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) AvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "avalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) AvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) AvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) GetAvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "getAvalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) GetAvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetAvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) GetAvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetAvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// AvalancheValidatorSetRegistryValidatorSetRegisteredIterator is returned from FilterValidatorSetRegistered and is used to iterate over the raw logs and unpacked data for ValidatorSetRegistered events raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetRegisteredIterator struct {
	Event *AvalancheValidatorSetRegistryValidatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalancheValidatorSetRegistryValidatorSetRegistered)
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
		it.Event = new(AvalancheValidatorSetRegistryValidatorSetRegistered)
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
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalancheValidatorSetRegistryValidatorSetRegistered represents a ValidatorSetRegistered event raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetRegistered struct {
	ValidatorSetID        *big.Int
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetRegistered is a free log retrieval operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) FilterValidatorSetRegistered(opts *bind.FilterOpts, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (*AvalancheValidatorSetRegistryValidatorSetRegisteredIterator, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.FilterLogs(opts, "ValidatorSetRegistered", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryValidatorSetRegisteredIterator{contract: _AvalancheValidatorSetRegistry.contract, event: "ValidatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorSetRegistered is a free log subscription operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) WatchValidatorSetRegistered(opts *bind.WatchOpts, sink chan<- *AvalancheValidatorSetRegistryValidatorSetRegistered, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.WatchLogs(opts, "ValidatorSetRegistered", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalancheValidatorSetRegistryValidatorSetRegistered)
				if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
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

// ParseValidatorSetRegistered is a log parse operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) ParseValidatorSetRegistered(log types.Log) (*AvalancheValidatorSetRegistryValidatorSetRegistered, error) {
	event := new(AvalancheValidatorSetRegistryValidatorSetRegistered)
	if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvalancheValidatorSetRegistryValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetUpdatedIterator struct {
	Event *AvalancheValidatorSetRegistryValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalancheValidatorSetRegistryValidatorSetUpdated)
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
		it.Event = new(AvalancheValidatorSetRegistryValidatorSetUpdated)
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
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalancheValidatorSetRegistryValidatorSetUpdated represents a ValidatorSetUpdated event raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetUpdated struct {
	ValidatorSetID        *big.Int
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (*AvalancheValidatorSetRegistryValidatorSetUpdatedIterator, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.FilterLogs(opts, "ValidatorSetUpdated", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryValidatorSetUpdatedIterator{contract: _AvalancheValidatorSetRegistry.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *AvalancheValidatorSetRegistryValidatorSetUpdated, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.WatchLogs(opts, "ValidatorSetUpdated", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalancheValidatorSetRegistryValidatorSetUpdated)
				if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) ParseValidatorSetUpdated(log types.Log) (*AvalancheValidatorSetRegistryValidatorSetUpdated, error) {
	event := new(AvalancheValidatorSetRegistryValidatorSetUpdated)
	if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
