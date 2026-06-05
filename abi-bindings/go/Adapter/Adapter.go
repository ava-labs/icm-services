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
	Bin: "0x610100604052348015610010575f5ffd5b506040516109ab3803806109ab83398101604081905261002f9161006d565b60809390935260a0919091526001600160a01b0390811660c0521660e0526100b1565b80516001600160a01b0381168114610068575f5ffd5b919050565b5f5f5f5f60808587031215610080575f5ffd5b84516020860151909450925061009860408601610052565b91506100a660608601610052565b905092959194509250565b60805160a05160c05160e05161089161011a5f395f818161011a015281816102b7015261044401525f8181606901528181610210015261037f01525f81816101410152818161027701526103fc01525f818160ad015281816101d0015261033701526108915ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80631fcaecb814610064578063ac60b059146100a8578063eb97cd2c146100dd578063f1faff00146100f2578063f59e207514610115578063f72916301461013c575b5f5ffd5b61008b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100cf7f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161009f565b6100f06100eb36600461047e565b610163565b005b6101056101003660046104bd565b610334565b604051901515815260200161009f565b61008b7f000000000000000000000000000000000000000000000000000000000000000081565b6100cf7f000000000000000000000000000000000000000000000000000000000000000081565b610173606082016040830161050a565b6001600160a01b0316336001600160a01b0316146101ce5760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b60448201526064015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000081606001350361027557604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90610245908490600401610788565b5f604051808303815f87803b15801561025c575f5ffd5b505af115801561026e573d5f5f3e3d5ffd5b5050505050565b7f00000000000000000000000000000000000000000000000000000000000000008160600135036102ec57604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90610245908490600401610788565b60405162461bcd60e51b815260206004820152601860248201527f556e657870656374656420626c6f636b636861696e204944000000000000000060448201526064016101c5565b5f7f0000000000000000000000000000000000000000000000000000000000000000610360838061079a565b60600135036103fa5760405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff00906103b49085906004016107b9565b6020604051808303815f875af11580156103d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f4919061083c565b92915050565b7f0000000000000000000000000000000000000000000000000000000000000000610425838061079a565b60600135036102ec5760405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff00906103b49085906004016107b9565b919050565b5f6020828403121561048e575f5ffd5b813567ffffffffffffffff8111156104a4575f5ffd5b820161012081850312156104b6575f5ffd5b9392505050565b5f602082840312156104cd575f5ffd5b813567ffffffffffffffff8111156104e3575f5ffd5b8201608081850312156104b6575f5ffd5b80356001600160a01b0381168114610479575f5ffd5b5f6020828403121561051a575f5ffd5b6104b6826104f4565b5f5f8335601e19843603018112610538575f5ffd5b830160208101925035905067ffffffffffffffff811115610557575f5ffd5b8060051b3603821315610568575f5ffd5b9250929050565b8183526020830192505f815f5b848110156105ab576001600160a01b03610595836104f4565b168652602095860195919091019060010161057c565b5093949350505050565b5f5f8335601e198436030181126105ca575f5ffd5b830160208101925035905067ffffffffffffffff8111156105e9575f5ffd5b8060061b3603821315610568575f5ffd5b8183526020830192505f815f5b848110156105ab57813586526001600160a01b03610627602084016104f4565b1660208701526040958601959190910190600101610607565b5f5f8335601e19843603018112610655575f5ffd5b830160208101925035905067ffffffffffffffff811115610674575f5ffd5b803603821315610568575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f6106bb602083016104f4565b6001600160a01b031660208401526106d5604083016104f4565b6001600160a01b03166040840152606082810135908401526106f9608083016104f4565b6001600160a01b0316608084015260a0828101359084015261071e60c0830183610523565b61012060c08601526107356101208601828461056f565b91505061074560e08401846105b5565b85830360e08701526107588382846105fa565b9250505061076a610100840184610640565b85830361010087015261077e838284610682565b9695505050505050565b602081525f6104b660208301846106aa565b5f823561011e198336030181126107af575f5ffd5b9190910192915050565b602081525f823561011e198436030181126107d2575f5ffd5b608060208401526107e860a084018583016106aa565b9050602084013563ffffffff8116808214610801575f5ffd5b80604086015250505f60408501359050806060850152506108256060850185610640565b848303601f1901608086015261077e838284610682565b5f6020828403121561084c575f5ffd5b815180151581146104b6575f5ffdfea2646970667358221220944fb1b1eedcdbf7433076efdc614decb507798e3a76e1f21e00e910712d746f64736f6c634300081e0033",
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
