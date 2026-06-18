// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecdsaverifier

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

// TeleporterICMMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterICMMessage struct {
	Message            TeleporterMessageV2
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessageV2 is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageV2 struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	OriginTeleporterAddress common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// ECDSAVerifierMetaData contains all meta data concerning the ECDSAVerifier contract.
var ECDSAVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ECDSAVerifierSendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761070a908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163c4d66de81461020257508063eb97cd2c1461018f5763f1faff001461003d575f80fd5b3461018b57602036600319011261018b5760043567ffffffffffffffff811161018b57803603608060031982011261018b57600482013561012219820181121561018b576040516100c1816100a3602082019460608652600460808401918901016103b0565b6044870135604083015230606083015203601f19810183528261056c565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f52601c52603c5f20916064810135916022190182121561018b570160048101359067ffffffffffffffff821161018b576024810190823603821361018b576040519061013c601f8501601f19166020018361056c565b838252602060048536930101011161018b576020935f8585610172966101699683870137840101526105a2565b909291926105dc565b5f546040516001600160a01b0392831691909216148152f35b5f80fd5b3461018b57602036600319011261018b5760043567ffffffffffffffff811161018b57610120600319823603011261018b576101fd7f7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8916040519182916020835260208301906004016103b0565b0390a1005b3461018b57602036600319011261018b576004356001600160a01b0381169081900361018b575f5160206106de5f395f51905f52549160ff8360401c16159267ffffffffffffffff811680159081610394575b600114908161038a575b159081610381575b506103725767ffffffffffffffff1981166001175f5160206106de5f395f51905f525583610346575b50811561030b57506bffffffffffffffffffffffff60a01b5f5416175f556102b457005b68ff0000000000000000195f5160206106de5f395f51905f5254165f5160206106de5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b62461bcd60e51b8152602060048201526016602482015275496e76616c6964207369676e6572206164647265737360501b6044820152606490fd5b68ffffffffffffffffff191668010000000000000001175f5160206106de5f395f51905f525583610290565b63f92ee8a960e01b5f5260045ffd5b90501585610267565b303b15915061025f565b859150610255565b35906001600160a01b038216820361018b57565b803582526101208201906001600160a01b036103ce6020830161039c565b1660208401526001600160a01b036103e86040830161039c565b166040840152606081810135908401526001600160a01b0361040c6080830161039c565b16608084015260a081013560a084015260c0810135601e1982360301908181121561018b5782016020813591019367ffffffffffffffff821161018b578160051b3603851361018b57819061012060c088015252610140850193905f5b8181106105465750505060e08201358181121561018b5782016020813591019367ffffffffffffffff821161018b578160061b3603851361018b5785810360e087015281815260200193905f5b818110610516575050506101008201359081121561018b57016020813591019267ffffffffffffffff821161018b57813603841361018b576020938161010084938603910152818452848401375f828201840152601f01601f1916010190565b90919460408060019288358152838060a01b0361053560208b0161039c565b1660208201520196019291016104b6565b909194602080600192838060a01b0361055e8a61039c565b168152019601929101610469565b90601f8019910116810190811067ffffffffffffffff82111761058e57604052565b634e487b7160e01b5f52604160045260245ffd5b81519190604183036105d2576105cb9250602082015190606060408401519301515f1a90610650565b9192909190565b50505f9160029190565b600481101561063c57806105ee575050565b600181036106055763f645eedf60e01b5f5260045ffd5b60028103610620575063fce698f760e01b5f5260045260245ffd5b60031461062a5750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084116106d2579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa156106c7575f516001600160a01b038116156106bd57905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f916003919056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a164736f6c634300081e000a",
}

// ECDSAVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAVerifierMetaData.ABI instead.
var ECDSAVerifierABI = ECDSAVerifierMetaData.ABI

// ECDSAVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAVerifierMetaData.Bin instead.
var ECDSAVerifierBin = ECDSAVerifierMetaData.Bin

// DeployECDSAVerifier deploys a new Ethereum contract, binding an instance of ECDSAVerifier to it.
func DeployECDSAVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSAVerifier, error) {
	parsed, err := ECDSAVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSAVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSAVerifier{ECDSAVerifierCaller: ECDSAVerifierCaller{contract: contract}, ECDSAVerifierTransactor: ECDSAVerifierTransactor{contract: contract}, ECDSAVerifierFilterer: ECDSAVerifierFilterer{contract: contract}}, nil
}

// ECDSAVerifier is an auto generated Go binding around an Ethereum contract.
type ECDSAVerifier struct {
	ECDSAVerifierCaller     // Read-only binding to the contract
	ECDSAVerifierTransactor // Write-only binding to the contract
	ECDSAVerifierFilterer   // Log filterer for contract events
}

// ECDSAVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSAVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSAVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSAVerifierSession struct {
	Contract     *ECDSAVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSAVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSAVerifierCallerSession struct {
	Contract *ECDSAVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ECDSAVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSAVerifierTransactorSession struct {
	Contract     *ECDSAVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ECDSAVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSAVerifierRaw struct {
	Contract *ECDSAVerifier // Generic contract binding to access the raw methods on
}

// ECDSAVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSAVerifierCallerRaw struct {
	Contract *ECDSAVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSAVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSAVerifierTransactorRaw struct {
	Contract *ECDSAVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSAVerifier creates a new instance of ECDSAVerifier, bound to a specific deployed contract.
func NewECDSAVerifier(address common.Address, backend bind.ContractBackend) (*ECDSAVerifier, error) {
	contract, err := bindECDSAVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifier{ECDSAVerifierCaller: ECDSAVerifierCaller{contract: contract}, ECDSAVerifierTransactor: ECDSAVerifierTransactor{contract: contract}, ECDSAVerifierFilterer: ECDSAVerifierFilterer{contract: contract}}, nil
}

// NewECDSAVerifierCaller creates a new read-only instance of ECDSAVerifier, bound to a specific deployed contract.
func NewECDSAVerifierCaller(address common.Address, caller bind.ContractCaller) (*ECDSAVerifierCaller, error) {
	contract, err := bindECDSAVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifierCaller{contract: contract}, nil
}

// NewECDSAVerifierTransactor creates a new write-only instance of ECDSAVerifier, bound to a specific deployed contract.
func NewECDSAVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSAVerifierTransactor, error) {
	contract, err := bindECDSAVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifierTransactor{contract: contract}, nil
}

// NewECDSAVerifierFilterer creates a new log filterer instance of ECDSAVerifier, bound to a specific deployed contract.
func NewECDSAVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAVerifierFilterer, error) {
	contract, err := bindECDSAVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifierFilterer{contract: contract}, nil
}

// bindECDSAVerifier binds a generic wrapper to an already deployed contract.
func bindECDSAVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAVerifier *ECDSAVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAVerifier.Contract.ECDSAVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAVerifier *ECDSAVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.ECDSAVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAVerifier *ECDSAVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.ECDSAVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAVerifier *ECDSAVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAVerifier *ECDSAVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAVerifier *ECDSAVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ECDSAVerifier *ECDSAVerifierCaller) VerifyMessage(opts *bind.CallOpts, message TeleporterICMMessage) (bool, error) {
	var out []interface{}
	err := _ECDSAVerifier.contract.Call(opts, &out, "verifyMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ECDSAVerifier *ECDSAVerifierSession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _ECDSAVerifier.Contract.VerifyMessage(&_ECDSAVerifier.CallOpts, message)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ECDSAVerifier *ECDSAVerifierCallerSession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _ECDSAVerifier.Contract.VerifyMessage(&_ECDSAVerifier.CallOpts, message)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address signer) returns()
func (_ECDSAVerifier *ECDSAVerifierTransactor) Initialize(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _ECDSAVerifier.contract.Transact(opts, "initialize", signer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address signer) returns()
func (_ECDSAVerifier *ECDSAVerifierSession) Initialize(signer common.Address) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.Initialize(&_ECDSAVerifier.TransactOpts, signer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address signer) returns()
func (_ECDSAVerifier *ECDSAVerifierTransactorSession) Initialize(signer common.Address) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.Initialize(&_ECDSAVerifier.TransactOpts, signer)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ECDSAVerifier *ECDSAVerifierTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _ECDSAVerifier.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ECDSAVerifier *ECDSAVerifierSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.SendMessage(&_ECDSAVerifier.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ECDSAVerifier *ECDSAVerifierTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ECDSAVerifier.Contract.SendMessage(&_ECDSAVerifier.TransactOpts, message)
}

// ECDSAVerifierECDSAVerifierSendMessageIterator is returned from FilterECDSAVerifierSendMessage and is used to iterate over the raw logs and unpacked data for ECDSAVerifierSendMessage events raised by the ECDSAVerifier contract.
type ECDSAVerifierECDSAVerifierSendMessageIterator struct {
	Event *ECDSAVerifierECDSAVerifierSendMessage // Event containing the contract specifics and raw log

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
func (it *ECDSAVerifierECDSAVerifierSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAVerifierECDSAVerifierSendMessage)
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
		it.Event = new(ECDSAVerifierECDSAVerifierSendMessage)
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
func (it *ECDSAVerifierECDSAVerifierSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAVerifierECDSAVerifierSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAVerifierECDSAVerifierSendMessage represents a ECDSAVerifierSendMessage event raised by the ECDSAVerifier contract.
type ECDSAVerifierECDSAVerifierSendMessage struct {
	Message TeleporterMessageV2
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterECDSAVerifierSendMessage is a free log retrieval operation binding the contract event 0x7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8.
//
// Solidity: event ECDSAVerifierSendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ECDSAVerifier *ECDSAVerifierFilterer) FilterECDSAVerifierSendMessage(opts *bind.FilterOpts) (*ECDSAVerifierECDSAVerifierSendMessageIterator, error) {

	logs, sub, err := _ECDSAVerifier.contract.FilterLogs(opts, "ECDSAVerifierSendMessage")
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifierECDSAVerifierSendMessageIterator{contract: _ECDSAVerifier.contract, event: "ECDSAVerifierSendMessage", logs: logs, sub: sub}, nil
}

// WatchECDSAVerifierSendMessage is a free log subscription operation binding the contract event 0x7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8.
//
// Solidity: event ECDSAVerifierSendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ECDSAVerifier *ECDSAVerifierFilterer) WatchECDSAVerifierSendMessage(opts *bind.WatchOpts, sink chan<- *ECDSAVerifierECDSAVerifierSendMessage) (event.Subscription, error) {

	logs, sub, err := _ECDSAVerifier.contract.WatchLogs(opts, "ECDSAVerifierSendMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAVerifierECDSAVerifierSendMessage)
				if err := _ECDSAVerifier.contract.UnpackLog(event, "ECDSAVerifierSendMessage", log); err != nil {
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

// ParseECDSAVerifierSendMessage is a log parse operation binding the contract event 0x7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8.
//
// Solidity: event ECDSAVerifierSendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ECDSAVerifier *ECDSAVerifierFilterer) ParseECDSAVerifierSendMessage(log types.Log) (*ECDSAVerifierECDSAVerifierSendMessage, error) {
	event := new(ECDSAVerifierECDSAVerifierSendMessage)
	if err := _ECDSAVerifier.contract.UnpackLog(event, "ECDSAVerifierSendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ECDSAVerifier contract.
type ECDSAVerifierInitializedIterator struct {
	Event *ECDSAVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *ECDSAVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAVerifierInitialized)
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
		it.Event = new(ECDSAVerifierInitialized)
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
func (it *ECDSAVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAVerifierInitialized represents a Initialized event raised by the ECDSAVerifier contract.
type ECDSAVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ECDSAVerifier *ECDSAVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*ECDSAVerifierInitializedIterator, error) {

	logs, sub, err := _ECDSAVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ECDSAVerifierInitializedIterator{contract: _ECDSAVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ECDSAVerifier *ECDSAVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ECDSAVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _ECDSAVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAVerifierInitialized)
				if err := _ECDSAVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ECDSAVerifier *ECDSAVerifierFilterer) ParseInitialized(log types.Log) (*ECDSAVerifierInitialized, error) {
	event := new(ECDSAVerifierInitialized)
	if err := _ECDSAVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
