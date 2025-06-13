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
	ABI: "[{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"allowedRelayer\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405273253b2784c75e510dd0ff1da844684a1ac0aa5fcf6080523480156026575f80fd5b5060805161043b6100445f395f81816052015260e6015261043b5ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80632a638943146100385780633cb747bf1461004d575b5f80fd5b61004b61004636600461020a565b610090565b005b6100747f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6040805160018082528183019092525f916020808301908036833701905050905081815f815181106100c4576100c46102a0565b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c00160405280868152602001896001600160a01b0316815260200160405180604001604052805f6001600160a01b031681526020015f8152508152602001620186a0815260200184815260200188886040516020016101789291906102b4565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016101a69190610368565b6020604051808303815f875af11580156101c2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e691906103ee565b50505050505050565b80356001600160a01b0381168114610205575f80fd5b919050565b5f805f805f6080868803121561021e575f80fd5b610227866101ef565b9450602086013567ffffffffffffffff80821115610243575f80fd5b818801915088601f830112610256575f80fd5b813581811115610264575f80fd5b896020828501011115610275575f80fd5b60208301965080955050505060408601359150610294606087016101ef565b90509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f815180845260208085019450602084015f5b8381101561031a5781516001600160a01b0316875295820195908201906001016102f5565b509495945050505050565b5f81518084525f5b818110156103495760208185018101518683018201520161032d565b505f602082860101526020601f19601f83011685010191505092915050565b60208152815160208201525f602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526103c86101008401826102e2565b905060a0840151601f198483030160e08501526103e58282610325565b95945050505050565b5f602082840312156103fe575f80fd5b505191905056fea2646970667358221220758ce67fd7e390c574a3d97f17c64b694dcc32c111256ac876589213d78aac4d64736f6c63430008190033",
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
