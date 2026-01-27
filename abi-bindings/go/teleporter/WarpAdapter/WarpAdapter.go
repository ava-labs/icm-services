// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package warpadapter

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

// ICMMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMMessage struct {
	UnsignedMessage    TeleporterMessage
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
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

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// WarpAdapterMetaData contains all meta data concerning the WarpAdapter contract.
var WarpAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"unsignedMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506109088061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063b771b3bc14610043578063b85958cf1461006e578063eb97cd2c14610091575b5f5ffd5b6100516005600160991b0181565b6040516001600160a01b0390911681526020015b60405180910390f35b61008161007c36600461034e565b6100a6565b6040519015158152602001610065565b6100a461009f36600461038c565b6102b8565b005b5f806100b560408401846103c4565b8101906100c2919061040e565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa15801561010e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261013591908101906104ec565b91509150806101955760405162461bcd60e51b815260206004820152602160248201527f57617270416461707465723a20696e76616c69642077617270206d65737361676044820152606560f81b60648201526084015b60405180910390fd5b60208201516001600160a01b031630146102045760405162461bcd60e51b815260206004820152602a60248201527f57617270416461707465723a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b606482015260840161018c565b815160208601351461026a5760405162461bcd60e51b815260206004820152602960248201527f57617270416461707465723a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b606482015260840161018c565b604082015180516020909101205f61028287806105e1565b604051602001610292919061078f565b60408051601f198184030181529190528051602090910120919091149695505050505050565b6005600160991b016001600160a01b031663ee5b48eb826040516020016102df919061078f565b6040516020818303038152906040526040518263ffffffff1660e01b815260040161030a9190610889565b6020604051808303815f875af1158015610326573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061034a91906108bb565b5050565b5f6020828403121561035e575f5ffd5b813567ffffffffffffffff811115610374575f5ffd5b820160608185031215610385575f5ffd5b9392505050565b5f6020828403121561039c575f5ffd5b813567ffffffffffffffff8111156103b2575f5ffd5b82016101208185031215610385575f5ffd5b5f5f8335601e198436030181126103d9575f5ffd5b83018035915067ffffffffffffffff8211156103f3575f5ffd5b602001915036819003821315610407575f5ffd5b9250929050565b5f6020828403121561041e575f5ffd5b813563ffffffff81168114610385575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561046857610468610431565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561049757610497610431565b604052919050565b6001600160a01b03811681146104b3575f5ffd5b50565b5f5b838110156104d05781810151838201526020016104b8565b50505f910152565b805180151581146104e7575f5ffd5b919050565b5f5f604083850312156104fd575f5ffd5b825167ffffffffffffffff811115610513575f5ffd5b830160608186031215610524575f5ffd5b61052c610445565b81518152602082015161053e8161049f565b6020820152604082015167ffffffffffffffff81111561055c575f5ffd5b80830192505085601f830112610570575f5ffd5b815167ffffffffffffffff81111561058a5761058a610431565b61059d601f8201601f191660200161046e565b8181528760208386010111156105b1575f5ffd5b6105c28260208301602087016104b6565b60408301525092506105d89050602084016104d8565b90509250929050565b5f823561011e198336030181126105f6575f5ffd5b9190910192915050565b80356104e78161049f565b5f5f8335601e19843603018112610620575f5ffd5b830160208101925035905067ffffffffffffffff81111561063f575f5ffd5b8060051b3603821315610407575f5ffd5b8183526020830192505f815f5b8481101561068e5781356106708161049f565b6001600160a01b03168652602095860195919091019060010161065d565b5093949350505050565b5f5f8335601e198436030181126106ad575f5ffd5b830160208101925035905067ffffffffffffffff8111156106cc575f5ffd5b8060061b3603821315610407575f5ffd5b8183526020830192505f815f5b8481101561068e578135865260208201356107048161049f565b6001600160a01b0316602087015260409586019591909101906001016106ea565b5f5f8335601e1984360301811261073a575f5ffd5b830160208101925035905067ffffffffffffffff811115610759575f5ffd5b803603821315610407575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f906107a7908401610600565b6001600160a01b0381166040840152506107c360408401610600565b6001600160a01b03811660608401525060608301356080838101919091526107ec908401610600565b6001600160a01b03811660a08401525060a083013560c0838101919091526108169084018461060b565b61012060e085015261082d61014085018284610650565b91505061083d60e0850185610698565b848303601f19016101008601526108558382846106dd565b92505050610867610100850185610725565b848303601f190161012086015261087f838284610767565b9695505050505050565b602081525f82518060208401526108a78160408501602087016104b6565b601f01601f19169190910160400192915050565b5f602082840312156108cb575f5ffd5b505191905056fea2646970667358221220f686d3bfd99af7dce517585a868ec8efe8d8f5622d60f9e1c0bc34951aabd6db64736f6c634300081e0033",
}

// WarpAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use WarpAdapterMetaData.ABI instead.
var WarpAdapterABI = WarpAdapterMetaData.ABI

// WarpAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WarpAdapterMetaData.Bin instead.
var WarpAdapterBin = WarpAdapterMetaData.Bin

// DeployWarpAdapter deploys a new Ethereum contract, binding an instance of WarpAdapter to it.
func DeployWarpAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WarpAdapter, error) {
	parsed, err := WarpAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WarpAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WarpAdapter{WarpAdapterCaller: WarpAdapterCaller{contract: contract}, WarpAdapterTransactor: WarpAdapterTransactor{contract: contract}, WarpAdapterFilterer: WarpAdapterFilterer{contract: contract}}, nil
}

// WarpAdapter is an auto generated Go binding around an Ethereum contract.
type WarpAdapter struct {
	WarpAdapterCaller     // Read-only binding to the contract
	WarpAdapterTransactor // Write-only binding to the contract
	WarpAdapterFilterer   // Log filterer for contract events
}

// WarpAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarpAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarpAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarpAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarpAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarpAdapterSession struct {
	Contract     *WarpAdapter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WarpAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarpAdapterCallerSession struct {
	Contract *WarpAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WarpAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarpAdapterTransactorSession struct {
	Contract     *WarpAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WarpAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarpAdapterRaw struct {
	Contract *WarpAdapter // Generic contract binding to access the raw methods on
}

// WarpAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarpAdapterCallerRaw struct {
	Contract *WarpAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// WarpAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarpAdapterTransactorRaw struct {
	Contract *WarpAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarpAdapter creates a new instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapter(address common.Address, backend bind.ContractBackend) (*WarpAdapter, error) {
	contract, err := bindWarpAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarpAdapter{WarpAdapterCaller: WarpAdapterCaller{contract: contract}, WarpAdapterTransactor: WarpAdapterTransactor{contract: contract}, WarpAdapterFilterer: WarpAdapterFilterer{contract: contract}}, nil
}

// NewWarpAdapterCaller creates a new read-only instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterCaller(address common.Address, caller bind.ContractCaller) (*WarpAdapterCaller, error) {
	contract, err := bindWarpAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterCaller{contract: contract}, nil
}

// NewWarpAdapterTransactor creates a new write-only instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*WarpAdapterTransactor, error) {
	contract, err := bindWarpAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterTransactor{contract: contract}, nil
}

// NewWarpAdapterFilterer creates a new log filterer instance of WarpAdapter, bound to a specific deployed contract.
func NewWarpAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*WarpAdapterFilterer, error) {
	contract, err := bindWarpAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarpAdapterFilterer{contract: contract}, nil
}

// bindWarpAdapter binds a generic wrapper to an already deployed contract.
func bindWarpAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarpAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarpAdapter *WarpAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarpAdapter.Contract.WarpAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarpAdapter *WarpAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarpAdapter.Contract.WarpAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarpAdapter *WarpAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarpAdapter.Contract.WarpAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarpAdapter *WarpAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarpAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarpAdapter *WarpAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarpAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarpAdapter *WarpAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarpAdapter.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarpAdapter.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterSession) WARPMESSENGER() (common.Address, error) {
	return _WarpAdapter.Contract.WARPMESSENGER(&_WarpAdapter.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_WarpAdapter *WarpAdapterCallerSession) WARPMESSENGER() (common.Address, error) {
	return _WarpAdapter.Contract.WARPMESSENGER(&_WarpAdapter.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xb85958cf.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterCaller) VerifyMessage(opts *bind.CallOpts, message ICMMessage) (bool, error) {
	var out []interface{}
	err := _WarpAdapter.contract.Call(opts, &out, "verifyMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0xb85958cf.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterSession) VerifyMessage(message ICMMessage) (bool, error) {
	return _WarpAdapter.Contract.VerifyMessage(&_WarpAdapter.CallOpts, message)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xb85958cf.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message) view returns(bool)
func (_WarpAdapter *WarpAdapterCallerSession) VerifyMessage(message ICMMessage) (bool, error) {
	return _WarpAdapter.Contract.VerifyMessage(&_WarpAdapter.CallOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _WarpAdapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterSession) SendMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _WarpAdapter.Contract.SendMessage(&_WarpAdapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_WarpAdapter *WarpAdapterTransactorSession) SendMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _WarpAdapter.Contract.SendMessage(&_WarpAdapter.TransactOpts, message)
}
