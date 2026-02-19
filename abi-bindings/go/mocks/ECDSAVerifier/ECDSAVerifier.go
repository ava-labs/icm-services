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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ECDSAVerifierSendMessage\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516108c03803806108c083398101604081905261002e91610099565b6001600160a01b0381166100885760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964207369676e6572206164647265737300000000000000000000604482015260640160405180910390fd5b6001600160a01b03166080526100c6565b5f602082840312156100a9575f5ffd5b81516001600160a01b03811681146100bf575f5ffd5b9392505050565b6080516107dc6100e45f395f8181608501526101bf01526107dc5ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063eb97cd2c14610043578063f1faff0014610058578063f74d548014610080575b5f5ffd5b6100566100513660046103f6565b6100bf565b005b61006b610066366004610435565b6100f9565b60405190151581526020015b60405180910390f35b6100a77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610077565b7f7f79990a356de554936f38da80d42a1fb6ea1198955703669370ad6bfcf297d8816040516100ee91906106ec565b60405180910390a150565b5f8061010583806106fe565b83604001353060405160200161011d9392919061071d565b6040516020818303038152906040528051906020012090505f61016c827f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f908152601c91909152603c902090565b90505f6101bb61017f606087018761074f565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525086939250506101fc9050565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316149350505050919050565b5f5f5f5f61020a8686610224565b92509250925061021a828261026d565b5090949350505050565b5f5f5f835160410361025b576020840151604085015160608601515f1a61024d8882858561032e565b955095509550505050610266565b505081515f91506002905b9250925092565b5f82600381111561028057610280610792565b03610289575050565b600182600381111561029d5761029d610792565b036102bb5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156102cf576102cf610792565b036102f55760405163fce698f760e01b8152600481018290526024015b60405180910390fd5b600382600381111561030957610309610792565b0361032a576040516335e2f38360e21b8152600481018290526024016102ec565b5050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561036757505f915060039050826103ec565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156103b8573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166103e357505f9250600191508290506103ec565b92505f91508190505b9450945094915050565b5f60208284031215610406575f5ffd5b813567ffffffffffffffff81111561041c575f5ffd5b8201610120818503121561042e575f5ffd5b9392505050565b5f60208284031215610445575f5ffd5b813567ffffffffffffffff81111561045b575f5ffd5b82016080818503121561042e575f5ffd5b80356001600160a01b0381168114610482575f5ffd5b919050565b5f5f8335601e1984360301811261049c575f5ffd5b830160208101925035905067ffffffffffffffff8111156104bb575f5ffd5b8060051b36038213156104cc575f5ffd5b9250929050565b8183526020830192505f815f5b8481101561050f576001600160a01b036104f98361046c565b16865260209586019591909101906001016104e0565b5093949350505050565b5f5f8335601e1984360301811261052e575f5ffd5b830160208101925035905067ffffffffffffffff81111561054d575f5ffd5b8060061b36038213156104cc575f5ffd5b8183526020830192505f815f5b8481101561050f57813586526001600160a01b0361058b6020840161046c565b166020870152604095860195919091019060010161056b565b5f5f8335601e198436030181126105b9575f5ffd5b830160208101925035905067ffffffffffffffff8111156105d8575f5ffd5b8036038213156104cc575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f61061f6020830161046c565b6001600160a01b031660208401526106396040830161046c565b6001600160a01b031660408401526060828101359084015261065d6080830161046c565b6001600160a01b0316608084015260a0828101359084015261068260c0830183610487565b61012060c0860152610699610120860182846104d3565b9150506106a960e0840184610519565b85830360e08701526106bc83828461055e565b925050506106ce6101008401846105a4565b8583036101008701526106e28382846105e6565b9695505050505050565b602081525f61042e602083018461060e565b5f823561011e19833603018112610713575f5ffd5b9190910192915050565b606081525f61072f606083018661060e565b6020830194909452506001600160a01b0391909116604090910152919050565b5f5f8335601e19843603018112610764575f5ffd5b83018035915067ffffffffffffffff82111561077e575f5ffd5b6020019150368190038213156104cc575f5ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220f42f1704d1b001338e1cae58be770a86ba87827873bc595ba88ef8f2834e57ef64736f6c634300081e0033",
}

// ECDSAVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAVerifierMetaData.ABI instead.
var ECDSAVerifierABI = ECDSAVerifierMetaData.ABI

// ECDSAVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAVerifierMetaData.Bin instead.
var ECDSAVerifierBin = ECDSAVerifierMetaData.Bin

// DeployECDSAVerifier deploys a new Ethereum contract, binding an instance of ECDSAVerifier to it.
func DeployECDSAVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, signer common.Address) (common.Address, *types.Transaction, *ECDSAVerifier, error) {
	parsed, err := ECDSAVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSAVerifierBin), backend, signer)
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

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_ECDSAVerifier *ECDSAVerifierCaller) TrustedSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSAVerifier.contract.Call(opts, &out, "trustedSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_ECDSAVerifier *ECDSAVerifierSession) TrustedSigner() (common.Address, error) {
	return _ECDSAVerifier.Contract.TrustedSigner(&_ECDSAVerifier.CallOpts)
}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_ECDSAVerifier *ECDSAVerifierCallerSession) TrustedSigner() (common.Address, error) {
	return _ECDSAVerifier.Contract.TrustedSigner(&_ECDSAVerifier.CallOpts)
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
