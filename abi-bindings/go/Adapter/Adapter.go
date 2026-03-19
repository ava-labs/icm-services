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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chain1_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chain2_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"adapter1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter2_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"adapter1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapter2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chain1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chain2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051610a80380380610a8083398101604081905261002f9161006d565b60809390935260a0919091526001600160a01b0390811660c0521660e0526100b1565b80516001600160a01b0381168114610068575f5ffd5b919050565b5f5f5f5f60808587031215610080575f5ffd5b84516020860151909450925061009860408601610052565b91506100a660608601610052565b905092959194509250565b60805160a05160c05160e05161096661011a5f395f818161011a01528181610307015261054101525f8181606901528181610204015261042901525f81816101410152818161026c01526104a701525f818160ad01528181610169015261038f01526109665ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80631fcaecb814610064578063ac60b059146100a8578063eb97cd2c146100dd578063f1faff00146100f2578063f59e207514610115578063f72916301461013c575b5f5ffd5b61008b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100cf7f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161009f565b6100f06100eb366004610576565b610163565b005b6101056101003660046105ae565b610388565b604051901515815260200161009f565b61008b7f000000000000000000000000000000000000000000000000000000000000000081565b6100cf7f000000000000000000000000000000000000000000000000000000000000000081565b5f3390507f0000000000000000000000000000000000000000000000000000000000000000816001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e891906105e5565b0361026a57604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c9061023990859060040161087c565b5f604051808303815f87803b158015610250575f5ffd5b505af1158015610262573d5f5f3e3d5ffd5b505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000816001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102eb91906105e5565b0361033c57604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c9061023990859060040161087c565b60405162461bcd60e51b815260206004820152601860248201527f556e657870656374656420626c6f636b636861696e2049440000000000000000604482015260640160405180910390fd5b5f5f3390507f0000000000000000000000000000000000000000000000000000000000000000816001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ea573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061040e91906105e5565b036104a55760405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff009061045e90869060040161088e565b6020604051808303815f875af115801561047a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049e9190610911565b9392505050565b7f0000000000000000000000000000000000000000000000000000000000000000816001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610502573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052691906105e5565b0361033c5760405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff009061045e90869060040161088e565b5f60208284031215610586575f5ffd5b813567ffffffffffffffff81111561059c575f5ffd5b8201610120818503121561049e575f5ffd5b5f602082840312156105be575f5ffd5b813567ffffffffffffffff8111156105d4575f5ffd5b82016080818503121561049e575f5ffd5b5f602082840312156105f5575f5ffd5b5051919050565b80356001600160a01b0381168114610612575f5ffd5b919050565b5f5f8335601e1984360301811261062c575f5ffd5b830160208101925035905067ffffffffffffffff81111561064b575f5ffd5b8060051b360382131561065c575f5ffd5b9250929050565b8183526020830192505f815f5b8481101561069f576001600160a01b03610689836105fc565b1686526020958601959190910190600101610670565b5093949350505050565b5f5f8335601e198436030181126106be575f5ffd5b830160208101925035905067ffffffffffffffff8111156106dd575f5ffd5b8060061b360382131561065c575f5ffd5b8183526020830192505f815f5b8481101561069f57813586526001600160a01b0361071b602084016105fc565b16602087015260409586019591909101906001016106fb565b5f5f8335601e19843603018112610749575f5ffd5b830160208101925035905067ffffffffffffffff811115610768575f5ffd5b80360382131561065c575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f6107af602083016105fc565b6001600160a01b031660208401526107c9604083016105fc565b6001600160a01b03166040840152606082810135908401526107ed608083016105fc565b6001600160a01b0316608084015260a0828101359084015261081260c0830183610617565b61012060c086015261082961012086018284610663565b91505061083960e08401846106a9565b85830360e087015261084c8382846106ee565b9250505061085e610100840184610734565b858303610100870152610872838284610776565b9695505050505050565b602081525f61049e602083018461079e565b602081525f823561011e198436030181126108a7575f5ffd5b608060208401526108bd60a0840185830161079e565b9050602084013563ffffffff81168082146108d6575f5ffd5b80604086015250505f60408501359050806060850152506108fa6060850185610734565b848303601f19016080860152610872838284610776565b5f60208284031215610921575f5ffd5b8151801515811461049e575f5ffdfea2646970667358221220bbc48779aa2c938d7568708d7430858d6f57a4072073ba3e9e15e978e8f014ba64736f6c634300081e0033",
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
