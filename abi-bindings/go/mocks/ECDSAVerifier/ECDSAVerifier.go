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
	Bin: "0x6080604052348015600e575f5ffd5b5061090e8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063c4d66de814610043578063eb97cd2c14610058578063f1faff001461006b575b5f5ffd5b61005661005136600461052a565b610092565b005b61005661006636600461054a565b610205565b61007e610079366004610582565b61023f565b604051901515815260200160405180910390f35b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156100d75750825b90505f8267ffffffffffffffff1660011480156100f35750303b155b905081158015610101575080155b1561011f5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561014957845460ff60401b1916600160401b1785555b6001600160a01b03861661019d5760405162461bcd60e51b8152602060048201526016602482015275496e76616c6964207369676e6572206164647265737360501b60448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b03881617905583156101fd57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d881604051610234919061081e565b60405180910390a150565b5f8061024b8380610830565b8360400135306040516020016102639392919061084f565b6040516020818303038152906040528051906020012090505f6102b2827f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f908152601c91909152603c902090565b90505f6103016102c56060870187610881565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250869392505061031a9050565b5f546001600160a01b0391821691161495945050505050565b5f5f5f5f6103288686610342565b925092509250610338828261038b565b5090949350505050565b5f5f5f8351604103610379576020840151604085015160608601515f1a61036b88828585610447565b955095509550505050610384565b505081515f91506002905b9250925092565b5f82600381111561039e5761039e6108c4565b036103a7575050565b60018260038111156103bb576103bb6108c4565b036103d95760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156103ed576103ed6108c4565b0361040e5760405163fce698f760e01b815260048101829052602401610194565b6003826003811115610422576104226108c4565b03610443576040516335e2f38360e21b815260048101829052602401610194565b5050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561048057505f91506003905082610505565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156104d1573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166104fc57505f925060019150829050610505565b92505f91508190505b9450945094915050565b80356001600160a01b0381168114610525575f5ffd5b919050565b5f6020828403121561053a575f5ffd5b6105438261050f565b9392505050565b5f6020828403121561055a575f5ffd5b813567ffffffffffffffff811115610570575f5ffd5b82016101208185031215610543575f5ffd5b5f60208284031215610592575f5ffd5b813567ffffffffffffffff8111156105a8575f5ffd5b820160808185031215610543575f5ffd5b5f5f8335601e198436030181126105ce575f5ffd5b830160208101925035905067ffffffffffffffff8111156105ed575f5ffd5b8060051b36038213156105fe575f5ffd5b9250929050565b8183526020830192505f815f5b84811015610641576001600160a01b0361062b8361050f565b1686526020958601959190910190600101610612565b5093949350505050565b5f5f8335601e19843603018112610660575f5ffd5b830160208101925035905067ffffffffffffffff81111561067f575f5ffd5b8060061b36038213156105fe575f5ffd5b8183526020830192505f815f5b8481101561064157813586526001600160a01b036106bd6020840161050f565b166020870152604095860195919091019060010161069d565b5f5f8335601e198436030181126106eb575f5ffd5b830160208101925035905067ffffffffffffffff81111561070a575f5ffd5b8036038213156105fe575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f6107516020830161050f565b6001600160a01b0316602084015261076b6040830161050f565b6001600160a01b031660408401526060828101359084015261078f6080830161050f565b6001600160a01b0316608084015260a082810135908401526107b460c08301836105b9565b61012060c08601526107cb61012086018284610605565b9150506107db60e084018461064b565b85830360e08701526107ee838284610690565b925050506108006101008401846106d6565b858303610100870152610814838284610718565b9695505050505050565b602081525f6105436020830184610740565b5f823561011e19833603018112610845575f5ffd5b9190910192915050565b606081525f6108616060830186610740565b6020830194909452506001600160a01b0391909116604090910152919050565b5f5f8335601e19843603018112610896575f5ffd5b83018035915067ffffffffffffffff8211156108b0575f5ffd5b6020019150368190038213156105fe575f5ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220d5dc5a894c5c0f4b2d271cd3e07bbd3bdfb8ae8e21f2b5fd326b6afea191bf5064736f6c634300081e0033",
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
