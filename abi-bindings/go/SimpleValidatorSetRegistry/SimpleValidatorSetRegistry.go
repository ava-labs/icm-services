// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplevalidatorsetregistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
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

// SimpleValidatorSetRegistryValidator is an auto generated low-level Go binding around an user-defined struct.
type SimpleValidatorSetRegistryValidator struct {
	BlsPublicKey []byte
	Weight       uint64
}

// SimpleValidatorSetRegistryValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type SimpleValidatorSetRegistryValidatorSet struct {
	AvalancheBlockchainID [32]byte
	Validators            []SimpleValidatorSetRegistryValidator
	TotalWeight           uint64
	PChainHeight          uint64
}

// SimplevalidatorsetregistryMetaData contains all meta data concerning the Simplevalidatorsetregistry contract.
var SimplevalidatorsetregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avalancheNetworkID\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_avalancheBlockChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avalancheBlockChainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avalancheNetworkID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentValidatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSimpleValidatorSetRegistry.ValidatorSet\",\"components\":[{\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structSimpleValidatorSetRegistry.Validator[]\",\"components\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"totalWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pChainHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorSet\",\"inputs\":[{\"name\":\"validatorSetID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSimpleValidatorSetRegistry.ValidatorSet\",\"components\":[{\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structSimpleValidatorSetRegistry.Validator[]\",\"components\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"totalWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pChainHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextValidatorSetID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerValidatorSet\",\"inputs\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structSimpleValidatorSetRegistry.Validator[]\",\"components\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pChainHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateValidatorSet\",\"inputs\":[{\"name\":\"validatorSetID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structSimpleValidatorSetRegistry.Validator[]\",\"components\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"pChainHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ValidatorSetRegistered\",\"inputs\":[{\"name\":\"validatorSetID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetUpdated\",\"inputs\":[{\"name\":\"validatorSetID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x60c06040525f805463ffffffff1916905534801561001b575f5ffd5b50604051610f75380380610f7583398101604081905261003a9161004d565b63ffffffff90911660805260a052610081565b5f5f6040838503121561005e575f5ffd5b825163ffffffff81168114610071575f5ffd5b6020939093015192949293505050565b60805160a051610ebf6100b65f395f8181610135015281816103e0015281816104ed015261076801525f60fb0152610ebf5ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80635cf6785a116100585780635cf6785a146100e157806368531ed0146100f6578063a04e82931461011d578063b490c33314610130575f5ffd5b80630209fdd01461007e5780631405d5ef1461009c57806338153622146100bd575b5f5ffd5b610086610157565b6040516100939190610a26565b60405180910390f35b6100af6100aa366004610b76565b610351565b604051908152602001610093565b5f546100cc9063ffffffff1681565b60405163ffffffff9091168152602001610093565b6100f46100ef366004610bc8565b61053f565b005b6100cc7f000000000000000000000000000000000000000000000000000000000000000081565b61008661012b366004610c22565b6107b8565b6100af7f000000000000000000000000000000000000000000000000000000000000000081565b61019160405180608001604052805f8152602001606081526020015f6001600160401b031681526020015f6001600160401b031681525090565b5f5463ffffffff166101ea5760405162461bcd60e51b815260206004820152601760248201527f4e6f2076616c696461746f72207365747320657869737400000000000000000060448201526064015b60405180910390fd5b5f80546001919061020290839063ffffffff16610c4d565b63ffffffff1681526020019081526020015f206040518060800160405290815f820154815260200160018201805480602002602001604051908101604052809291908181526020015f905b8282101561031f578382905f5260205f2090600202016040518060400160405290815f8201805461027d90610c6f565b80601f01602080910402602001604051908101604052809291908181526020018280546102a990610c6f565b80156102f45780601f106102cb576101008083540402835291602001916102f4565b820191905f5260205f20905b8154815290600101906020018083116102d757829003601f168201915b50505091835250506001918201546001600160401b031660209182015291835292909201910161024d565b50505090825250600291909101546001600160401b038082166020840152600160401b90910416604090910152919050565b5f8261039f5760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d70747900000060448201526064016101e1565b5f805463ffffffff1681806103b383610ca7565b82546101009290920a63ffffffff818102199093169183160217909155165f8181526001602052604081207f0000000000000000000000000000000000000000000000000000000000000000815560028101805467ffffffffffffffff60401b1916600160401b6001600160401b03891602179055919250805b868110156104c9578260010188888381811061044b5761044b610ccb565b905060200281019061045d9190610cdf565b81546001810183555f928352602090922090916002020161047e8282610d55565b505087878281811061049257610492610ccb565b90506020028101906104a49190610cdf565b6104b5906040810190602001610e71565b6104bf9083610e93565b915060010161042d565b5060028201805467ffffffffffffffff19166001600160401b0383161790556040517f00000000000000000000000000000000000000000000000000000000000000009084907fe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583905f90a3509095945050505050565b5f5463ffffffff1684106105955760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f722073657420646f6573206e6f742065786973740000000060448201526064016101e1565b816105e25760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d70747900000060448201526064016101e1565b5f84815260016020526040902060028101546001600160401b03600160401b9091048116908316116106565760405162461bcd60e51b815260206004820152601e60248201527f502d436861696e20686569676874206d7573742062652067726561746572000060448201526064016101e1565b5f85815260016020819052604082206106719291019061098c565b5f85815260016020526040812060028101805467ffffffffffffffff60401b1916600160401b6001600160401b0387160217905590805b8581101561074457826001018787838181106106c6576106c6610ccb565b90506020028101906106d89190610cdf565b81546001810183555f92835260209092209091600202016106f98282610d55565b505086868281811061070d5761070d610ccb565b905060200281019061071f9190610cdf565b610730906040810190602001610e71565b61073a9083610e93565b91506001016106a8565b5060028201805467ffffffffffffffff19166001600160401b0383161790556040517f00000000000000000000000000000000000000000000000000000000000000009088907fd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9905f90a350505050505050565b6107f260405180608001604052805f8152602001606081526020015f6001600160401b031681526020015f6001600160401b031681525090565b5f5463ffffffff1682106108485760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f722073657420646f6573206e6f742065786973740000000060448201526064016101e1565b5f828152600160208181526040808420815160808101835281548152938101805483518186028101860190945280845294959194868501949192909184015b82821015610959578382905f5260205f2090600202016040518060400160405290815f820180546108b790610c6f565b80601f01602080910402602001604051908101604052809291908181526020018280546108e390610c6f565b801561092e5780601f106109055761010080835404028352916020019161092e565b820191905f5260205f20905b81548152906001019060200180831161091157829003601f168201915b50505091835250506001918201546001600160401b0316602091820152918352929092019101610887565b50505090825250600291909101546001600160401b038082166020840152600160401b9091041660409091015292915050565b5080545f8255600202905f5260205f20908101906109aa91906109ad565b50565b808211156109dc575f6109c082826109e0565b5060018101805467ffffffffffffffff191690556002016109ad565b5090565b5080546109ec90610c6f565b5f825580601f106109fb575050565b601f0160209004905f5260205f20908101906109aa91905b808211156109dc575f8155600101610a13565b602081525f60a082018351602084015260208401516080604085015281815180845260c08601915060c08160051b87010193506020830192505f5b81811015610ae65760bf19878603018352835180516040875280518060408901525f5b81811015610aa157602081840181015160608b8401015201610a84565b505f6060828a0101526001600160401b0360208401511660208901526060601f19601f8301168901019750505050602084019350602083019250600181019050610a61565b5050505060408401516001600160401b03811660608501525060608401516001600160401b0381166080850152509392505050565b5f5f83601f840112610b2b575f5ffd5b5081356001600160401b03811115610b41575f5ffd5b6020830191508360208260051b8501011115610b5b575f5ffd5b9250929050565b6001600160401b03811681146109aa575f5ffd5b5f5f5f60408486031215610b88575f5ffd5b83356001600160401b03811115610b9d575f5ffd5b610ba986828701610b1b565b9094509250506020840135610bbd81610b62565b809150509250925092565b5f5f5f5f60608587031215610bdb575f5ffd5b8435935060208501356001600160401b03811115610bf7575f5ffd5b610c0387828801610b1b565b9094509250506040850135610c1781610b62565b939692955090935050565b5f60208284031215610c32575f5ffd5b5035919050565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8281168282160390811115610c6957610c69610c39565b92915050565b600181811c90821680610c8357607f821691505b602082108103610ca157634e487b7160e01b5f52602260045260245ffd5b50919050565b5f63ffffffff821663ffffffff8103610cc257610cc2610c39565b60010192915050565b634e487b7160e01b5f52603260045260245ffd5b5f8235603e19833603018112610cf3575f5ffd5b9190910192915050565b601f821115610d4457805f5260205f20601f840160051c81016020851015610d225750805b601f840160051c820191505b81811015610d41575f8155600101610d2e565b50505b505050565b5f8135610c6981610b62565b8135601e19833603018112610d68575f5ffd5b820180356001600160401b0381118015610d80575f5ffd5b813603602084011315610d91575f5ffd5b5f905050610da981610da38554610c6f565b85610cfd565b5f601f821160018114610ddd575f8315610dc65750838201602001355b5f19600385901b1c1916600184901b178555610e39565b5f85815260208120601f198516915b82811015610e0e57602085880181013583559485019460019092019101610dec565b5084821015610e2d575f1960f88660031b161c19602085880101351681555b505060018360011b0185555b50505050610e6d610e4c60208401610d49565b600183016001600160401b0382166001600160401b03198254161781555050565b5050565b5f60208284031215610e81575f5ffd5b8135610e8c81610b62565b9392505050565b6001600160401b038181168382160190811115610c6957610c69610c3956fea164736f6c634300081e000a",
}

// SimplevalidatorsetregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplevalidatorsetregistryMetaData.ABI instead.
var SimplevalidatorsetregistryABI = SimplevalidatorsetregistryMetaData.ABI

// SimplevalidatorsetregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplevalidatorsetregistryMetaData.Bin instead.
var SimplevalidatorsetregistryBin = SimplevalidatorsetregistryMetaData.Bin

// DeploySimplevalidatorsetregistry deploys a new Ethereum contract, binding an instance of Simplevalidatorsetregistry to it.
func DeploySimplevalidatorsetregistry(auth *bind.TransactOpts, backend bind.ContractBackend, _avalancheNetworkID uint32, _avalancheBlockChainID [32]byte) (common.Address, *types.Transaction, *Simplevalidatorsetregistry, error) {
	parsed, err := SimplevalidatorsetregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplevalidatorsetregistryBin), backend, _avalancheNetworkID, _avalancheBlockChainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplevalidatorsetregistry{SimplevalidatorsetregistryCaller: SimplevalidatorsetregistryCaller{contract: contract}, SimplevalidatorsetregistryTransactor: SimplevalidatorsetregistryTransactor{contract: contract}, SimplevalidatorsetregistryFilterer: SimplevalidatorsetregistryFilterer{contract: contract}}, nil
}

// Simplevalidatorsetregistry is an auto generated Go binding around an Ethereum contract.
type Simplevalidatorsetregistry struct {
	SimplevalidatorsetregistryCaller     // Read-only binding to the contract
	SimplevalidatorsetregistryTransactor // Write-only binding to the contract
	SimplevalidatorsetregistryFilterer   // Log filterer for contract events
}

// SimplevalidatorsetregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplevalidatorsetregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplevalidatorsetregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplevalidatorsetregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplevalidatorsetregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplevalidatorsetregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplevalidatorsetregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplevalidatorsetregistrySession struct {
	Contract     *Simplevalidatorsetregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SimplevalidatorsetregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplevalidatorsetregistryCallerSession struct {
	Contract *SimplevalidatorsetregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// SimplevalidatorsetregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplevalidatorsetregistryTransactorSession struct {
	Contract     *SimplevalidatorsetregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// SimplevalidatorsetregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplevalidatorsetregistryRaw struct {
	Contract *Simplevalidatorsetregistry // Generic contract binding to access the raw methods on
}

// SimplevalidatorsetregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplevalidatorsetregistryCallerRaw struct {
	Contract *SimplevalidatorsetregistryCaller // Generic read-only contract binding to access the raw methods on
}

// SimplevalidatorsetregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplevalidatorsetregistryTransactorRaw struct {
	Contract *SimplevalidatorsetregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplevalidatorsetregistry creates a new instance of Simplevalidatorsetregistry, bound to a specific deployed contract.
func NewSimplevalidatorsetregistry(address common.Address, backend bind.ContractBackend) (*Simplevalidatorsetregistry, error) {
	contract, err := bindSimplevalidatorsetregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplevalidatorsetregistry{SimplevalidatorsetregistryCaller: SimplevalidatorsetregistryCaller{contract: contract}, SimplevalidatorsetregistryTransactor: SimplevalidatorsetregistryTransactor{contract: contract}, SimplevalidatorsetregistryFilterer: SimplevalidatorsetregistryFilterer{contract: contract}}, nil
}

// NewSimplevalidatorsetregistryCaller creates a new read-only instance of Simplevalidatorsetregistry, bound to a specific deployed contract.
func NewSimplevalidatorsetregistryCaller(address common.Address, caller bind.ContractCaller) (*SimplevalidatorsetregistryCaller, error) {
	contract, err := bindSimplevalidatorsetregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplevalidatorsetregistryCaller{contract: contract}, nil
}

// NewSimplevalidatorsetregistryTransactor creates a new write-only instance of Simplevalidatorsetregistry, bound to a specific deployed contract.
func NewSimplevalidatorsetregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplevalidatorsetregistryTransactor, error) {
	contract, err := bindSimplevalidatorsetregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplevalidatorsetregistryTransactor{contract: contract}, nil
}

// NewSimplevalidatorsetregistryFilterer creates a new log filterer instance of Simplevalidatorsetregistry, bound to a specific deployed contract.
func NewSimplevalidatorsetregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplevalidatorsetregistryFilterer, error) {
	contract, err := bindSimplevalidatorsetregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplevalidatorsetregistryFilterer{contract: contract}, nil
}

// bindSimplevalidatorsetregistry binds a generic wrapper to an already deployed contract.
func bindSimplevalidatorsetregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplevalidatorsetregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplevalidatorsetregistry.Contract.SimplevalidatorsetregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.SimplevalidatorsetregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.SimplevalidatorsetregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplevalidatorsetregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.contract.Transact(opts, method, params...)
}

// AvalancheBlockChainID is a free data retrieval call binding the contract method 0xb490c333.
//
// Solidity: function avalancheBlockChainID() view returns(bytes32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCaller) AvalancheBlockChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Simplevalidatorsetregistry.contract.Call(opts, &out, "avalancheBlockChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AvalancheBlockChainID is a free data retrieval call binding the contract method 0xb490c333.
//
// Solidity: function avalancheBlockChainID() view returns(bytes32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) AvalancheBlockChainID() ([32]byte, error) {
	return _Simplevalidatorsetregistry.Contract.AvalancheBlockChainID(&_Simplevalidatorsetregistry.CallOpts)
}

// AvalancheBlockChainID is a free data retrieval call binding the contract method 0xb490c333.
//
// Solidity: function avalancheBlockChainID() view returns(bytes32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerSession) AvalancheBlockChainID() ([32]byte, error) {
	return _Simplevalidatorsetregistry.Contract.AvalancheBlockChainID(&_Simplevalidatorsetregistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCaller) AvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Simplevalidatorsetregistry.contract.Call(opts, &out, "avalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) AvalancheNetworkID() (uint32, error) {
	return _Simplevalidatorsetregistry.Contract.AvalancheNetworkID(&_Simplevalidatorsetregistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerSession) AvalancheNetworkID() (uint32, error) {
	return _Simplevalidatorsetregistry.Contract.AvalancheNetworkID(&_Simplevalidatorsetregistry.CallOpts)
}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCaller) GetCurrentValidatorSet(opts *bind.CallOpts) (SimpleValidatorSetRegistryValidatorSet, error) {
	var out []interface{}
	err := _Simplevalidatorsetregistry.contract.Call(opts, &out, "getCurrentValidatorSet")

	if err != nil {
		return *new(SimpleValidatorSetRegistryValidatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(SimpleValidatorSetRegistryValidatorSet)).(*SimpleValidatorSetRegistryValidatorSet)

	return out0, err

}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) GetCurrentValidatorSet() (SimpleValidatorSetRegistryValidatorSet, error) {
	return _Simplevalidatorsetregistry.Contract.GetCurrentValidatorSet(&_Simplevalidatorsetregistry.CallOpts)
}

// GetCurrentValidatorSet is a free data retrieval call binding the contract method 0x0209fdd0.
//
// Solidity: function getCurrentValidatorSet() view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerSession) GetCurrentValidatorSet() (SimpleValidatorSetRegistryValidatorSet, error) {
	return _Simplevalidatorsetregistry.Contract.GetCurrentValidatorSet(&_Simplevalidatorsetregistry.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xa04e8293.
//
// Solidity: function getValidatorSet(uint256 validatorSetID) view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCaller) GetValidatorSet(opts *bind.CallOpts, validatorSetID *big.Int) (SimpleValidatorSetRegistryValidatorSet, error) {
	var out []interface{}
	err := _Simplevalidatorsetregistry.contract.Call(opts, &out, "getValidatorSet", validatorSetID)

	if err != nil {
		return *new(SimpleValidatorSetRegistryValidatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(SimpleValidatorSetRegistryValidatorSet)).(*SimpleValidatorSetRegistryValidatorSet)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0xa04e8293.
//
// Solidity: function getValidatorSet(uint256 validatorSetID) view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) GetValidatorSet(validatorSetID *big.Int) (SimpleValidatorSetRegistryValidatorSet, error) {
	return _Simplevalidatorsetregistry.Contract.GetValidatorSet(&_Simplevalidatorsetregistry.CallOpts, validatorSetID)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0xa04e8293.
//
// Solidity: function getValidatorSet(uint256 validatorSetID) view returns((bytes32,(bytes,uint64)[],uint64,uint64))
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerSession) GetValidatorSet(validatorSetID *big.Int) (SimpleValidatorSetRegistryValidatorSet, error) {
	return _Simplevalidatorsetregistry.Contract.GetValidatorSet(&_Simplevalidatorsetregistry.CallOpts, validatorSetID)
}

// NextValidatorSetID is a free data retrieval call binding the contract method 0x38153622.
//
// Solidity: function nextValidatorSetID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCaller) NextValidatorSetID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Simplevalidatorsetregistry.contract.Call(opts, &out, "nextValidatorSetID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextValidatorSetID is a free data retrieval call binding the contract method 0x38153622.
//
// Solidity: function nextValidatorSetID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) NextValidatorSetID() (uint32, error) {
	return _Simplevalidatorsetregistry.Contract.NextValidatorSetID(&_Simplevalidatorsetregistry.CallOpts)
}

// NextValidatorSetID is a free data retrieval call binding the contract method 0x38153622.
//
// Solidity: function nextValidatorSetID() view returns(uint32)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryCallerSession) NextValidatorSetID() (uint32, error) {
	return _Simplevalidatorsetregistry.Contract.NextValidatorSetID(&_Simplevalidatorsetregistry.CallOpts)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x1405d5ef.
//
// Solidity: function registerValidatorSet((bytes,uint64)[] validators, uint64 pChainHeight) returns(uint256)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactor) RegisterValidatorSet(opts *bind.TransactOpts, validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.contract.Transact(opts, "registerValidatorSet", validators, pChainHeight)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x1405d5ef.
//
// Solidity: function registerValidatorSet((bytes,uint64)[] validators, uint64 pChainHeight) returns(uint256)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) RegisterValidatorSet(validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.RegisterValidatorSet(&_Simplevalidatorsetregistry.TransactOpts, validators, pChainHeight)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x1405d5ef.
//
// Solidity: function registerValidatorSet((bytes,uint64)[] validators, uint64 pChainHeight) returns(uint256)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactorSession) RegisterValidatorSet(validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.RegisterValidatorSet(&_Simplevalidatorsetregistry.TransactOpts, validators, pChainHeight)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x5cf6785a.
//
// Solidity: function updateValidatorSet(uint256 validatorSetID, (bytes,uint64)[] validators, uint64 pChainHeight) returns()
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactor) UpdateValidatorSet(opts *bind.TransactOpts, validatorSetID *big.Int, validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.contract.Transact(opts, "updateValidatorSet", validatorSetID, validators, pChainHeight)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x5cf6785a.
//
// Solidity: function updateValidatorSet(uint256 validatorSetID, (bytes,uint64)[] validators, uint64 pChainHeight) returns()
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistrySession) UpdateValidatorSet(validatorSetID *big.Int, validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.UpdateValidatorSet(&_Simplevalidatorsetregistry.TransactOpts, validatorSetID, validators, pChainHeight)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x5cf6785a.
//
// Solidity: function updateValidatorSet(uint256 validatorSetID, (bytes,uint64)[] validators, uint64 pChainHeight) returns()
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryTransactorSession) UpdateValidatorSet(validatorSetID *big.Int, validators []SimpleValidatorSetRegistryValidator, pChainHeight uint64) (*types.Transaction, error) {
	return _Simplevalidatorsetregistry.Contract.UpdateValidatorSet(&_Simplevalidatorsetregistry.TransactOpts, validatorSetID, validators, pChainHeight)
}

// SimplevalidatorsetregistryValidatorSetRegisteredIterator is returned from FilterValidatorSetRegistered and is used to iterate over the raw logs and unpacked data for ValidatorSetRegistered events raised by the Simplevalidatorsetregistry contract.
type SimplevalidatorsetregistryValidatorSetRegisteredIterator struct {
	Event *SimplevalidatorsetregistryValidatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *SimplevalidatorsetregistryValidatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplevalidatorsetregistryValidatorSetRegistered)
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
		it.Event = new(SimplevalidatorsetregistryValidatorSetRegistered)
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
func (it *SimplevalidatorsetregistryValidatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplevalidatorsetregistryValidatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplevalidatorsetregistryValidatorSetRegistered represents a ValidatorSetRegistered event raised by the Simplevalidatorsetregistry contract.
type SimplevalidatorsetregistryValidatorSetRegistered struct {
	ValidatorSetID        *big.Int
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetRegistered is a free log retrieval operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) FilterValidatorSetRegistered(opts *bind.FilterOpts, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (*SimplevalidatorsetregistryValidatorSetRegisteredIterator, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _Simplevalidatorsetregistry.contract.FilterLogs(opts, "ValidatorSetRegistered", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &SimplevalidatorsetregistryValidatorSetRegisteredIterator{contract: _Simplevalidatorsetregistry.contract, event: "ValidatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorSetRegistered is a free log subscription operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) WatchValidatorSetRegistered(opts *bind.WatchOpts, sink chan<- *SimplevalidatorsetregistryValidatorSetRegistered, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _Simplevalidatorsetregistry.contract.WatchLogs(opts, "ValidatorSetRegistered", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplevalidatorsetregistryValidatorSetRegistered)
				if err := _Simplevalidatorsetregistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
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

// ParseValidatorSetRegistered is a log parse operation binding the contract event 0xe93e9f47e7810153341664fc2050adcb29c88899748615c477d17b712d621583.
//
// Solidity: event ValidatorSetRegistered(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) ParseValidatorSetRegistered(log types.Log) (*SimplevalidatorsetregistryValidatorSetRegistered, error) {
	event := new(SimplevalidatorsetregistryValidatorSetRegistered)
	if err := _Simplevalidatorsetregistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimplevalidatorsetregistryValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Simplevalidatorsetregistry contract.
type SimplevalidatorsetregistryValidatorSetUpdatedIterator struct {
	Event *SimplevalidatorsetregistryValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *SimplevalidatorsetregistryValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplevalidatorsetregistryValidatorSetUpdated)
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
		it.Event = new(SimplevalidatorsetregistryValidatorSetUpdated)
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
func (it *SimplevalidatorsetregistryValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplevalidatorsetregistryValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplevalidatorsetregistryValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Simplevalidatorsetregistry contract.
type SimplevalidatorsetregistryValidatorSetUpdated struct {
	ValidatorSetID        *big.Int
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (*SimplevalidatorsetregistryValidatorSetUpdatedIterator, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _Simplevalidatorsetregistry.contract.FilterLogs(opts, "ValidatorSetUpdated", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &SimplevalidatorsetregistryValidatorSetUpdatedIterator{contract: _Simplevalidatorsetregistry.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *SimplevalidatorsetregistryValidatorSetUpdated, validatorSetID []*big.Int, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var validatorSetIDRule []interface{}
	for _, validatorSetIDItem := range validatorSetID {
		validatorSetIDRule = append(validatorSetIDRule, validatorSetIDItem)
	}
	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _Simplevalidatorsetregistry.contract.WatchLogs(opts, "ValidatorSetUpdated", validatorSetIDRule, avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplevalidatorsetregistryValidatorSetUpdated)
				if err := _Simplevalidatorsetregistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xd48741f16bef6492997e28d107c7a13b06376de704072bdb37a9b02e502ea1f9.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID)
func (_Simplevalidatorsetregistry *SimplevalidatorsetregistryFilterer) ParseValidatorSetUpdated(log types.Log) (*SimplevalidatorsetregistryValidatorSetUpdated, error) {
	event := new(SimplevalidatorsetregistryValidatorSetUpdated)
	if err := _Simplevalidatorsetregistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
