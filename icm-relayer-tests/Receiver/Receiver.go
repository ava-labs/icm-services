// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiver

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

// ReceiverMetaData contains all meta data concerning the Receiver contract.
var ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedStr2\",\"type\":\"string\"}],\"name\":\"LogData\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"lastMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405273253b2784c75e510dd0ff1da844684a1ac0aa5fcf6080523480156026575f80fd5b5060805161055b6100455f395f81816066015261014b015261055b5ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806332970710146100435780633cb747bf14610061578063c868efaa146100a0575b5f80fd5b61004b6100b5565b604051610058919061024c565b60405180910390f35b6100887f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610058565b6100b36100ae366004610298565b610140565b005b5f80546100c190610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100ed90610325565b80156101385780601f1061010f57610100808354040283529160200191610138565b820191905f5260205f20905b81548152906001019060200180831161011b57829003601f168201915b505050505081565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101d75760405162461bcd60e51b815260206004820152603260248201527f52656365697665724f6e5375626e65743a20756e617574686f72697a6564205460448201527132b632b837b93a32b926b2b9b9b2b733b2b960711b606482015260840160405180910390fd5b6040516210d85d60ea1b81526003016040519081900381206448656c6c6f60d81b825290600501604051908190038120907ff41b1b5f2f5226e38f3ba40820fca31c08530df874fb47e7c3c5ffd60c825edb905f90a361023981830183610371565b5f906102459082610465565b5050505050565b5f602080835283518060208501525f5b818110156102785785810183015185820160400152820161025c565b505f604082860101526040601f19601f8301168501019250505092915050565b5f805f80606085870312156102ab575f80fd5b8435935060208501356001600160a01b03811681146102c8575f80fd5b9250604085013567ffffffffffffffff808211156102e4575f80fd5b818701915087601f8301126102f7575f80fd5b813581811115610305575f80fd5b886020828501011115610316575f80fd5b95989497505060200194505050565b600181811c9082168061033957607f821691505b60208210810361035757634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610381575f80fd5b813567ffffffffffffffff80821115610398575f80fd5b818401915084601f8301126103ab575f80fd5b8135818111156103bd576103bd61035d565b604051601f8201601f19908116603f011681019083821181831017156103e5576103e561035d565b816040528281528760208487010111156103fd575f80fd5b826020860160208301375f928101602001929092525095945050505050565b601f82111561046057805f5260205f20601f840160051c810160208510156104415750805b601f840160051c820191505b81811015610245575f815560010161044d565b505050565b815167ffffffffffffffff81111561047f5761047f61035d565b6104938161048d8454610325565b8461041c565b602080601f8311600181146104c6575f84156104af5750858301515b5f19600386901b1c1916600185901b17855561051d565b5f85815260208120601f198616915b828110156104f4578886015182559484019460019091019084016104d5565b508582101561051157878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212203c0c3288bd8799ab6c2323c694e55efaf330dc5764778b6089c5682915a8e7be64736f6c63430008190033",
}

// ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiverMetaData.ABI instead.
var ReceiverABI = ReceiverMetaData.ABI

// ReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiverMetaData.Bin instead.
var ReceiverBin = ReceiverMetaData.Bin

// DeployReceiver deploys a new Ethereum contract, binding an instance of Receiver to it.
func DeployReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Receiver, error) {
	parsed, err := ReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// Receiver is an auto generated Go binding around an Ethereum contract.
type Receiver struct {
	ReceiverCaller     // Read-only binding to the contract
	ReceiverTransactor // Write-only binding to the contract
	ReceiverFilterer   // Log filterer for contract events
}

// ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiverSession struct {
	Contract     *Receiver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiverCallerSession struct {
	Contract *ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiverTransactorSession struct {
	Contract     *ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiverRaw struct {
	Contract *Receiver // Generic contract binding to access the raw methods on
}

// ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiverCallerRaw struct {
	Contract *ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiverTransactorRaw struct {
	Contract *ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiver creates a new instance of Receiver, bound to a specific deployed contract.
func NewReceiver(address common.Address, backend bind.ContractBackend) (*Receiver, error) {
	contract, err := bindReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Receiver{ReceiverCaller: ReceiverCaller{contract: contract}, ReceiverTransactor: ReceiverTransactor{contract: contract}, ReceiverFilterer: ReceiverFilterer{contract: contract}}, nil
}

// NewReceiverCaller creates a new read-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverCaller(address common.Address, caller bind.ContractCaller) (*ReceiverCaller, error) {
	contract, err := bindReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverCaller{contract: contract}, nil
}

// NewReceiverTransactor creates a new write-only instance of Receiver, bound to a specific deployed contract.
func NewReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverTransactor, error) {
	contract, err := bindReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverTransactor{contract: contract}, nil
}

// NewReceiverFilterer creates a new log filterer instance of Receiver, bound to a specific deployed contract.
func NewReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverFilterer, error) {
	contract, err := bindReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverFilterer{contract: contract}, nil
}

// bindReceiver binds a generic wrapper to an already deployed contract.
func bindReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Receiver *ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Receiver *ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Receiver *ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Receiver.Contract.contract.Transact(opts, method, params...)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_Receiver *ReceiverCaller) LastMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Receiver.contract.Call(opts, &out, "lastMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_Receiver *ReceiverSession) LastMessage() (string, error) {
	return _Receiver.Contract.LastMessage(&_Receiver.CallOpts)
}

// LastMessage is a free data retrieval call binding the contract method 0x32970710.
//
// Solidity: function lastMessage() view returns(string)
func (_Receiver *ReceiverCallerSession) LastMessage() (string, error) {
	return _Receiver.Contract.LastMessage(&_Receiver.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Receiver *ReceiverCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Receiver.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Receiver *ReceiverSession) Messenger() (common.Address, error) {
	return _Receiver.Contract.Messenger(&_Receiver.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Receiver *ReceiverCallerSession) Messenger() (common.Address, error) {
	return _Receiver.Contract.Messenger(&_Receiver.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_Receiver *ReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _Receiver.contract.Transact(opts, "receiveTeleporterMessage", arg0, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_Receiver *ReceiverSession) ReceiveTeleporterMessage(arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveTeleporterMessage(&_Receiver.TransactOpts, arg0, arg1, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 , address , bytes message) returns()
func (_Receiver *ReceiverTransactorSession) ReceiveTeleporterMessage(arg0 [32]byte, arg1 common.Address, message []byte) (*types.Transaction, error) {
	return _Receiver.Contract.ReceiveTeleporterMessage(&_Receiver.TransactOpts, arg0, arg1, message)
}

// ReceiverLogDataIterator is returned from FilterLogData and is used to iterate over the raw logs and unpacked data for LogData events raised by the Receiver contract.
type ReceiverLogDataIterator struct {
	Event *ReceiverLogData // Event containing the contract specifics and raw log

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
func (it *ReceiverLogDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiverLogData)
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
		it.Event = new(ReceiverLogData)
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
func (it *ReceiverLogDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiverLogDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiverLogData represents a LogData event raised by the Receiver contract.
type ReceiverLogData struct {
	IndexedStr  common.Hash
	IndexedStr2 common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogData is a free log retrieval operation binding the contract event 0xf41b1b5f2f5226e38f3ba40820fca31c08530df874fb47e7c3c5ffd60c825edb.
//
// Solidity: event LogData(string indexed indexedStr, string indexed indexedStr2)
func (_Receiver *ReceiverFilterer) FilterLogData(opts *bind.FilterOpts, indexedStr []string, indexedStr2 []string) (*ReceiverLogDataIterator, error) {

	var indexedStrRule []interface{}
	for _, indexedStrItem := range indexedStr {
		indexedStrRule = append(indexedStrRule, indexedStrItem)
	}
	var indexedStr2Rule []interface{}
	for _, indexedStr2Item := range indexedStr2 {
		indexedStr2Rule = append(indexedStr2Rule, indexedStr2Item)
	}

	logs, sub, err := _Receiver.contract.FilterLogs(opts, "LogData", indexedStrRule, indexedStr2Rule)
	if err != nil {
		return nil, err
	}
	return &ReceiverLogDataIterator{contract: _Receiver.contract, event: "LogData", logs: logs, sub: sub}, nil
}

// WatchLogData is a free log subscription operation binding the contract event 0xf41b1b5f2f5226e38f3ba40820fca31c08530df874fb47e7c3c5ffd60c825edb.
//
// Solidity: event LogData(string indexed indexedStr, string indexed indexedStr2)
func (_Receiver *ReceiverFilterer) WatchLogData(opts *bind.WatchOpts, sink chan<- *ReceiverLogData, indexedStr []string, indexedStr2 []string) (event.Subscription, error) {

	var indexedStrRule []interface{}
	for _, indexedStrItem := range indexedStr {
		indexedStrRule = append(indexedStrRule, indexedStrItem)
	}
	var indexedStr2Rule []interface{}
	for _, indexedStr2Item := range indexedStr2 {
		indexedStr2Rule = append(indexedStr2Rule, indexedStr2Item)
	}

	logs, sub, err := _Receiver.contract.WatchLogs(opts, "LogData", indexedStrRule, indexedStr2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiverLogData)
				if err := _Receiver.contract.UnpackLog(event, "LogData", log); err != nil {
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

// ParseLogData is a log parse operation binding the contract event 0xf41b1b5f2f5226e38f3ba40820fca31c08530df874fb47e7c3c5ffd60c825edb.
//
// Solidity: event LogData(string indexed indexedStr, string indexed indexedStr2)
func (_Receiver *ReceiverFilterer) ParseLogData(log types.Log) (*ReceiverLogData, error) {
	event := new(ReceiverLogData)
	if err := _Receiver.contract.UnpackLog(event, "LogData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
