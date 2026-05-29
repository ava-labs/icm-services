// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

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

// AdapterMetaData contains all meta data concerning the Adapter contract.
var AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chain1_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chain2_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"adapter1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter2_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"adapter1\",\"outputs\":[{\"internalType\":\"contractIAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapter2\",\"outputs\":[{\"internalType\":\"contractIAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chain1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chain2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610100346100e357601f6108d438819003918201601f19168301916001600160401b038311848410176100e7578084926080946040528339810103126100e357805190602081015161005f6060610058604085016100fb565b93016100fb565b60809390935260a0526001600160a01b0390811660c0521660e0526040516107c490816101108239608051818181610180015281816104160152610674015260a0518181816075015281816104d40152610732015260c0518181816101b50152818161043e01526106be015260e05181818160b1015281816104fa015261077c0152f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100e35756fe6080806040526004361015610012575f80fd5b5f3560e01c9081631fcaecb8146101a357508063ac60b05914610169578063eb97cd2c14610129578063f1faff00146100e0578063f59e20751461009c5763f72916301461005e575f80fd5b34610098575f3660031901126100985760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b5f80fd5b34610098575f366003190112610098576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346100985760203660031901126100985760043567ffffffffffffffff811161009857608060031982360301126100985761011f602091600401610664565b6040519015158152f35b346100985760203660031901126100985760043567ffffffffffffffff8111610098576101206003198236030112610098576101679060040161040d565b005b34610098575f3660031901126100985760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34610098575f366003190112610098577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b90601f8019910116810190811067ffffffffffffffff82111761020657604052565b634e487b7160e01b5f52604160045260245ffd5b35906001600160a01b038216820361009857565b9035601e198236030181121561009857016020813591019167ffffffffffffffff821161009857813603831361009857565b908060209392818452848401375f828201840152601f01601f1916010190565b8035825291906101208101906001600160a01b036102a06020860161021a565b1660208201526001600160a01b036102ba6040860161021a565b166040820152606084810135908201526001600160a01b036102de6080860161021a565b16608082015260a084013560a082015260c0840135601e198536030190818112156100985785016020813591019367ffffffffffffffff8211610098578160051b3603851361009857819061012060c086015252610140830193905f5b8181106103e75750505060e0850135908112156100985784016020813591019267ffffffffffffffff8211610098578160061b360384136100985782810360e084015281815260200192905f5b8181106103b7575050506103a5846101006103b49596019061022e565b91610100818503910152610260565b90565b90919360408060019287358152838060a01b036103d660208a0161021a565b166020820152019501929101610388565b909194602080600192838060a01b036103ff8a61021a565b16815201960192910161033b565b5f9060608101357f000000000000000000000000000000000000000000000000000000000000000081036104cf57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b156104cb5761049983928392604051948580948193633ae5f34b60e21b8352602060048401526024830190610280565b03925af180156104c0576104ab575050565b6104b68280926101e4565b6104bd5750565b80fd5b6040513d84823e3d90fd5b8280fd5b9091507f00000000000000000000000000000000000000000000000000000000000000000361057e577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561009857604051633ae5f34b60e21b815260206004820152915f918391829084908290610556906024830190610280565b03925af18015610573576105675750565b5f610571916101e4565b565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152601860248201527f556e657870656374656420626c6f636b636861696e20494400000000000000006044820152606490fd5b90359061011e1981360301821215610098570190565b90816020910312610098575180151581036100985790565b60208152813561011e19833603018112156100985761061d90608060208401528360a084019101610280565b90602083013563ffffffff8116809103610098576103b49361065291604084015260408101356060840152606081019061022e565b916080601f1982860301910152610260565b606061067082806105c3565b01357f0000000000000000000000000000000000000000000000000000000000000000145f146107225760206106b9916040518093819262f1faff60e81b8352600483016105f1565b03815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610573575f916106f9575090565b6103b4915060203d60201161071b575b61071381836101e4565b8101906105d9565b503d610709565b606061072e82806105c3565b01357f0000000000000000000000000000000000000000000000000000000000000000145f1461057e576020610777916040518093819262f1faff60e81b8352600483016105f1565b03815f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af1908115610573575f916106f957509056fea164736f6c634300081e000a",
}

// AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use AdapterMetaData.ABI instead.
var AdapterABI = AdapterMetaData.ABI

// AdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdapterMetaData.Bin instead.
var AdapterBin = AdapterMetaData.Bin

// DeployAdapter deploys a new Ethereum contract, binding an instance of Adapter to it.
func DeployAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, chain1_ [32]byte, chain2_ [32]byte, adapter1_ common.Address, adapter2_ common.Address) (common.Address, *types.Transaction, *Adapter, error) {
	parsed, err := AdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdapterBin), backend, chain1_, chain2_, adapter1_, adapter2_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adapter{AdapterCaller: AdapterCaller{contract: contract}, AdapterTransactor: AdapterTransactor{contract: contract}, AdapterFilterer: AdapterFilterer{contract: contract}}, nil
}

// Adapter is an auto generated Go binding around an Ethereum contract.
type Adapter struct {
	AdapterCaller     // Read-only binding to the contract
	AdapterTransactor // Write-only binding to the contract
	AdapterFilterer   // Log filterer for contract events
}

// AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdapterSession struct {
	Contract     *Adapter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdapterCallerSession struct {
	Contract *AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdapterTransactorSession struct {
	Contract     *AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdapterRaw struct {
	Contract *Adapter // Generic contract binding to access the raw methods on
}

// AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdapterCallerRaw struct {
	Contract *AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdapterTransactorRaw struct {
	Contract *AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdapter creates a new instance of Adapter, bound to a specific deployed contract.
func NewAdapter(address common.Address, backend bind.ContractBackend) (*Adapter, error) {
	contract, err := bindAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adapter{AdapterCaller: AdapterCaller{contract: contract}, AdapterTransactor: AdapterTransactor{contract: contract}, AdapterFilterer: AdapterFilterer{contract: contract}}, nil
}

// NewAdapterCaller creates a new read-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterCaller(address common.Address, caller bind.ContractCaller) (*AdapterCaller, error) {
	contract, err := bindAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterCaller{contract: contract}, nil
}

// NewAdapterTransactor creates a new write-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*AdapterTransactor, error) {
	contract, err := bindAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterTransactor{contract: contract}, nil
}

// NewAdapterFilterer creates a new log filterer instance of Adapter, bound to a specific deployed contract.
func NewAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*AdapterFilterer, error) {
	contract, err := bindAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdapterFilterer{contract: contract}, nil
}

// bindAdapter binds a generic wrapper to an already deployed contract.
func bindAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transact(opts, method, params...)
}

// Adapter1 is a free data retrieval call binding the contract method 0x1fcaecb8.
//
// Solidity: function adapter1() view returns(address)
func (_Adapter *AdapterCaller) Adapter1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "adapter1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adapter1 is a free data retrieval call binding the contract method 0x1fcaecb8.
//
// Solidity: function adapter1() view returns(address)
func (_Adapter *AdapterSession) Adapter1() (common.Address, error) {
	return _Adapter.Contract.Adapter1(&_Adapter.CallOpts)
}

// Adapter1 is a free data retrieval call binding the contract method 0x1fcaecb8.
//
// Solidity: function adapter1() view returns(address)
func (_Adapter *AdapterCallerSession) Adapter1() (common.Address, error) {
	return _Adapter.Contract.Adapter1(&_Adapter.CallOpts)
}

// Adapter2 is a free data retrieval call binding the contract method 0xf59e2075.
//
// Solidity: function adapter2() view returns(address)
func (_Adapter *AdapterCaller) Adapter2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "adapter2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adapter2 is a free data retrieval call binding the contract method 0xf59e2075.
//
// Solidity: function adapter2() view returns(address)
func (_Adapter *AdapterSession) Adapter2() (common.Address, error) {
	return _Adapter.Contract.Adapter2(&_Adapter.CallOpts)
}

// Adapter2 is a free data retrieval call binding the contract method 0xf59e2075.
//
// Solidity: function adapter2() view returns(address)
func (_Adapter *AdapterCallerSession) Adapter2() (common.Address, error) {
	return _Adapter.Contract.Adapter2(&_Adapter.CallOpts)
}

// Chain1 is a free data retrieval call binding the contract method 0xac60b059.
//
// Solidity: function chain1() view returns(bytes32)
func (_Adapter *AdapterCaller) Chain1(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "chain1")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Chain1 is a free data retrieval call binding the contract method 0xac60b059.
//
// Solidity: function chain1() view returns(bytes32)
func (_Adapter *AdapterSession) Chain1() ([32]byte, error) {
	return _Adapter.Contract.Chain1(&_Adapter.CallOpts)
}

// Chain1 is a free data retrieval call binding the contract method 0xac60b059.
//
// Solidity: function chain1() view returns(bytes32)
func (_Adapter *AdapterCallerSession) Chain1() ([32]byte, error) {
	return _Adapter.Contract.Chain1(&_Adapter.CallOpts)
}

// Chain2 is a free data retrieval call binding the contract method 0xf7291630.
//
// Solidity: function chain2() view returns(bytes32)
func (_Adapter *AdapterCaller) Chain2(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "chain2")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Chain2 is a free data retrieval call binding the contract method 0xf7291630.
//
// Solidity: function chain2() view returns(bytes32)
func (_Adapter *AdapterSession) Chain2() ([32]byte, error) {
	return _Adapter.Contract.Chain2(&_Adapter.CallOpts)
}

// Chain2 is a free data retrieval call binding the contract method 0xf7291630.
//
// Solidity: function chain2() view returns(bytes32)
func (_Adapter *AdapterCallerSession) Chain2() ([32]byte, error) {
	return _Adapter.Contract.Chain2(&_Adapter.CallOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Adapter *AdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Adapter *AdapterSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _Adapter.Contract.SendMessage(&_Adapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Adapter *AdapterTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _Adapter.Contract.SendMessage(&_Adapter.TransactOpts, message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_Adapter *AdapterTransactor) VerifyMessage(opts *bind.TransactOpts, message TeleporterICMMessage) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "verifyMessage", message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_Adapter *AdapterSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _Adapter.Contract.VerifyMessage(&_Adapter.TransactOpts, message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_Adapter *AdapterTransactorSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _Adapter.Contract.VerifyMessage(&_Adapter.TransactOpts, message)
}
