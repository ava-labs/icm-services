// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sender

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

// SenderMetaData contains all meta data concerning the Sender contract.
var SenderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LogData\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"allowedRelayer\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405273253b2784c75e510dd0ff1da844684a1ac0aa5fcf6080523480156026575f80fd5b5060805161049f6100445f395f81816052015260e9015261049f5ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80632a638943146100385780633cb747bf1461004d575b5f80fd5b61004b61004636600461024a565b610090565b005b6100747f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6040805160018082528183019092525f916020808301908036833701905050905081815f815181106100c4576100c46102e0565b6001600160a01b039092166020928302919091019091015260015b60058111610226577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052808781526020018a6001600160a01b0316815260200160405180604001604052805f6001600160a01b031681526020015f8152508152602001620186a08152602001858152602001898960405160200161017b9291906102f4565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016101a991906103a8565b6020604051808303815f875af11580156101c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e9919061042e565b5060405142907fabf2763c2d37d1584cb198a2d5412ad4d09b51eb849dba580b05877010596436905f90a28061021e81610445565b9150506100df565b50505050505050565b80356001600160a01b0381168114610245575f80fd5b919050565b5f805f805f6080868803121561025e575f80fd5b6102678661022f565b9450602086013567ffffffffffffffff80821115610283575f80fd5b818801915088601f830112610296575f80fd5b8135818111156102a4575f80fd5b8960208285010111156102b5575f80fd5b602083019650809550505050604086013591506102d46060870161022f565b90509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f815180845260208085019450602084015f5b8381101561035a5781516001600160a01b031687529582019590820190600101610335565b509495945050505050565b5f81518084525f5b818110156103895760208185018101518683018201520161036d565b505f602082860101526020601f19601f83011685010191505092915050565b60208152815160208201525f602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152610408610100840182610322565b905060a0840151601f198483030160e08501526104258282610365565b95945050505050565b5f6020828403121561043e575f80fd5b5051919050565b5f6001820161046257634e487b7160e01b5f52601160045260245ffd5b506001019056fea26469706673582212203549ab191d020e63ebb711a2f10c912c1064d2f8f33f2791c3f5315ecbb2feae64736f6c63430008190033",
}

// SenderABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderMetaData.ABI instead.
var SenderABI = SenderMetaData.ABI

// SenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderMetaData.Bin instead.
var SenderBin = SenderMetaData.Bin

// DeploySender deploys a new Ethereum contract, binding an instance of Sender to it.
func DeploySender(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sender, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// Sender is an auto generated Go binding around an Ethereum contract.
type Sender struct {
	SenderCaller     // Read-only binding to the contract
	SenderTransactor // Write-only binding to the contract
	SenderFilterer   // Log filterer for contract events
}

// SenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderSession struct {
	Contract     *Sender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderCallerSession struct {
	Contract *SenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderTransactorSession struct {
	Contract     *SenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderRaw struct {
	Contract *Sender // Generic contract binding to access the raw methods on
}

// SenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderCallerRaw struct {
	Contract *SenderCaller // Generic read-only contract binding to access the raw methods on
}

// SenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderTransactorRaw struct {
	Contract *SenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSender creates a new instance of Sender, bound to a specific deployed contract.
func NewSender(address common.Address, backend bind.ContractBackend) (*Sender, error) {
	contract, err := bindSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sender{SenderCaller: SenderCaller{contract: contract}, SenderTransactor: SenderTransactor{contract: contract}, SenderFilterer: SenderFilterer{contract: contract}}, nil
}

// NewSenderCaller creates a new read-only instance of Sender, bound to a specific deployed contract.
func NewSenderCaller(address common.Address, caller bind.ContractCaller) (*SenderCaller, error) {
	contract, err := bindSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderCaller{contract: contract}, nil
}

// NewSenderTransactor creates a new write-only instance of Sender, bound to a specific deployed contract.
func NewSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderTransactor, error) {
	contract, err := bindSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderTransactor{contract: contract}, nil
}

// NewSenderFilterer creates a new log filterer instance of Sender, bound to a specific deployed contract.
func NewSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderFilterer, error) {
	contract, err := bindSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderFilterer{contract: contract}, nil
}

// bindSender binds a generic wrapper to an already deployed contract.
func bindSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.SenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.SenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sender *SenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sender *SenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sender *SenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sender.Contract.contract.Transact(opts, method, params...)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Sender *SenderCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sender.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Sender *SenderSession) Messenger() (common.Address, error) {
	return _Sender.Contract.Messenger(&_Sender.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Sender *SenderCallerSession) Messenger() (common.Address, error) {
	return _Sender.Contract.Messenger(&_Sender.CallOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x2a638943.
//
// Solidity: function sendMessage(address destinationAddress, string message, bytes32 destinationBlockchainID, address allowedRelayer) returns()
func (_Sender *SenderTransactor) SendMessage(opts *bind.TransactOpts, destinationAddress common.Address, message string, destinationBlockchainID [32]byte, allowedRelayer common.Address) (*types.Transaction, error) {
	return _Sender.contract.Transact(opts, "sendMessage", destinationAddress, message, destinationBlockchainID, allowedRelayer)
}

// SendMessage is a paid mutator transaction binding the contract method 0x2a638943.
//
// Solidity: function sendMessage(address destinationAddress, string message, bytes32 destinationBlockchainID, address allowedRelayer) returns()
func (_Sender *SenderSession) SendMessage(destinationAddress common.Address, message string, destinationBlockchainID [32]byte, allowedRelayer common.Address) (*types.Transaction, error) {
	return _Sender.Contract.SendMessage(&_Sender.TransactOpts, destinationAddress, message, destinationBlockchainID, allowedRelayer)
}

// SendMessage is a paid mutator transaction binding the contract method 0x2a638943.
//
// Solidity: function sendMessage(address destinationAddress, string message, bytes32 destinationBlockchainID, address allowedRelayer) returns()
func (_Sender *SenderTransactorSession) SendMessage(destinationAddress common.Address, message string, destinationBlockchainID [32]byte, allowedRelayer common.Address) (*types.Transaction, error) {
	return _Sender.Contract.SendMessage(&_Sender.TransactOpts, destinationAddress, message, destinationBlockchainID, allowedRelayer)
}

// SenderLogDataIterator is returned from FilterLogData and is used to iterate over the raw logs and unpacked data for LogData events raised by the Sender contract.
type SenderLogDataIterator struct {
	Event *SenderLogData // Event containing the contract specifics and raw log

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
func (it *SenderLogDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SenderLogData)
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
		it.Event = new(SenderLogData)
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
func (it *SenderLogDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SenderLogDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SenderLogData represents a LogData event raised by the Sender contract.
type SenderLogData struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogData is a free log retrieval operation binding the contract event 0xabf2763c2d37d1584cb198a2d5412ad4d09b51eb849dba580b05877010596436.
//
// Solidity: event LogData(uint256 indexed timestamp)
func (_Sender *SenderFilterer) FilterLogData(opts *bind.FilterOpts, timestamp []*big.Int) (*SenderLogDataIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Sender.contract.FilterLogs(opts, "LogData", timestampRule)
	if err != nil {
		return nil, err
	}
	return &SenderLogDataIterator{contract: _Sender.contract, event: "LogData", logs: logs, sub: sub}, nil
}

// WatchLogData is a free log subscription operation binding the contract event 0xabf2763c2d37d1584cb198a2d5412ad4d09b51eb849dba580b05877010596436.
//
// Solidity: event LogData(uint256 indexed timestamp)
func (_Sender *SenderFilterer) WatchLogData(opts *bind.WatchOpts, sink chan<- *SenderLogData, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Sender.contract.WatchLogs(opts, "LogData", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SenderLogData)
				if err := _Sender.contract.UnpackLog(event, "LogData", log); err != nil {
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

// ParseLogData is a log parse operation binding the contract event 0xabf2763c2d37d1584cb198a2d5412ad4d09b51eb849dba580b05877010596436.
//
// Solidity: event LogData(uint256 indexed timestamp)
func (_Sender *SenderFilterer) ParseLogData(log types.Log) (*SenderLogData, error) {
	event := new(SenderLogData)
	if err := _Sender.contract.UnpackLog(event, "LogData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
