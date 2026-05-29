// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorsetsig

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

// ValidatorSetSigMessage is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetSigMessage struct {
	TargetBlockchainID     [32]byte
	ValidatorSetSigAddress common.Address
	TargetContractAddress  common.Address
	Nonce                  *big.Int
	Value                  *big.Int
	Payload                []byte
}

// ValidatorSetSigMetaData contains all meta data concerning the ValidatorSetSig contract.
var ValidatorSetSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validatorBlockchainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Delivered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"targetBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorSetSigAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorSetSigMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"validateMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c080604052346100c4576020816109dc803803809161001f82856100db565b8339810103126100c4575160015f5560805260405163084279ef60e31b81526020816004817302000000000000000000000000000000000000055afa9081156100d0575f9161009a575b5060a0526040516108c99081610113823960805181818160ec0152610299015260a051818181608c01526107620152f35b90506020813d6020116100c8575b816100b5602093836100db565b810103126100c457515f610069565b5f80fd5b3d91506100a8565b6040513d5f823e3d90fd5b601f909101601f19168101906001600160401b038211908210176100fe57604052565b634e487b7160e01b5f52604160045260245ffdfe608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c9081630731775d14610666575080635433da42146102315780637d969c341461014b5780637ecebe001461010f5780638d6e579d146100d5578063b771b3bc146100b35763d127dc9b14610075575f61000f565b346100af575f3660031901126100af5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b5f80fd5b346100af575f3660031901126100af576040516005600160991b018152602090f35b346100af575f3660031901126100af5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346100af5760203660031901126100af576004356001600160a01b038116908190036100af575f526001602052602060405f2054604051908152f35b346100af5760203660031901126100af5760043567ffffffffffffffff81116100af5760c060031982360301126100af57604051906101898261067e565b8060040135825261019c602482016106bc565b60208301526101ad604482016106bc565b6040830152606481013560608301526084810135608083015260a48101359067ffffffffffffffff82116100af570190366023830112156100af5760048201356101f6816106d0565b90610204604051928361069a565b80825236602485830101116100af576020815f92602461001a9701838601378301015260a082015261075e565b60203660031901126100af5760043563ffffffff81168091036100af5760025f54146106575760025f9081556040516306f8253560e41b81526004810192909252816024816005600160991b015afa801561064c575f915f9161058f575b501561053c5780517f0000000000000000000000000000000000000000000000000000000000000000036104e85760208101516001600160a01b031661048d576040015180518101906020818303126100af5760208101519067ffffffffffffffff82116100af570160c0818303126100af576040519061030f8261067e565b60208101518252610322604082016106ec565b6020830152610333606082016106ec565b91604081019283526080820151936060820194855260a0830151926080830193845260c08101519067ffffffffffffffff82116100af57602061037e928161038b9501920101610700565b9160a0810192835261075e565b83519060018201809211610479575f92839260018060a01b0386511684526001602052604084205560018060a01b038551169051915191602083519301915af13d15610474573d6103db816106d0565b906103e9604051928361069a565b81525f60203d92013e5b1561042f57519051906001600160a01b03167f5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed95f80a360015f55005b60405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f725365745369673a2063616c6c206661696c6564000000006044820152606490fd5b6103f3565b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f725365745369673a206e6f6e2d7a65726f206f726967696e60448201526c53656e6465724164647265737360981b6064820152608490fd5b60405162461bcd60e51b815260206004820152602660248201527f56616c696461746f725365745369673a20696e76616c696420736f75726365436044820152651a185a5b925160d21b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f56616c696461746f725365745369673a20696e76616c69642077617270206d65604482015264737361676560d81b6064820152608490fd5b9150503d805f833e6105a1818361069a565b81016040828203126100af57815167ffffffffffffffff81116100af578201916060838303126100af57604051926060840184811067ffffffffffffffff82111761063857604052805184526105f9602082016106ec565b602085015260408101519267ffffffffffffffff84116100af576020936106209201610700565b604084015201519081151582036100af57908261028f565b634e487b7160e01b5f52604160045260245ffd5b6040513d5f823e3d90fd5b633ee5aeb560e01b5f5260045ffd5b346100af575f3660031901126100af57805f60209252f35b60c0810190811067ffffffffffffffff82111761063857604052565b90601f8019910116810190811067ffffffffffffffff82111761063857604052565b35906001600160a01b03821682036100af57565b67ffffffffffffffff811161063857601f01601f191660200190565b51906001600160a01b03821682036100af57565b81601f820112156100af57805190610717826106d0565b92610725604051948561069a565b828452602083830101116100af575f5b82811061074957505060205f918301015290565b80602080928401015182828701015201610735565b80517f0000000000000000000000000000000000000000000000000000000000000000036108635760208101516001600160a01b03163003610806576040818101516001600160a01b03165f90815260016020522054606090910151036107c157565b60405162461bcd60e51b815260206004820152601e60248201527f56616c696461746f725365745369673a20696e76616c6964206e6f6e636500006044820152606490fd5b60405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f725365745369673a20696e76616c69642076616c6964617460448201526e6f725365745369674164647265737360881b6064820152608490fd5b60405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f725365745369673a20696e76616c6964207461726765744260448201526a1b1bd8dad8da185a5b925160aa1b6064820152608490fdfea164736f6c634300081e000a",
}

// ValidatorSetSigABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetSigMetaData.ABI instead.
var ValidatorSetSigABI = ValidatorSetSigMetaData.ABI

// ValidatorSetSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorSetSigMetaData.Bin instead.
var ValidatorSetSigBin = ValidatorSetSigMetaData.Bin

// DeployValidatorSetSig deploys a new Ethereum contract, binding an instance of ValidatorSetSig to it.
func DeployValidatorSetSig(auth *bind.TransactOpts, backend bind.ContractBackend, validatorBlockchainID_ [32]byte) (common.Address, *types.Transaction, *ValidatorSetSig, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorSetSigBin), backend, validatorBlockchainID_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// ValidatorSetSig is an auto generated Go binding around an Ethereum contract.
type ValidatorSetSig struct {
	ValidatorSetSigCaller     // Read-only binding to the contract
	ValidatorSetSigTransactor // Write-only binding to the contract
	ValidatorSetSigFilterer   // Log filterer for contract events
}

// ValidatorSetSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSigSession struct {
	Contract     *ValidatorSetSig  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetSigCallerSession struct {
	Contract *ValidatorSetSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ValidatorSetSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetSigTransactorSession struct {
	Contract     *ValidatorSetSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ValidatorSetSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetSigRaw struct {
	Contract *ValidatorSetSig // Generic contract binding to access the raw methods on
}

// ValidatorSetSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetSigCallerRaw struct {
	Contract *ValidatorSetSigCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactorRaw struct {
	Contract *ValidatorSetSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSetSig creates a new instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSig(address common.Address, backend bind.ContractBackend) (*ValidatorSetSig, error) {
	contract, err := bindValidatorSetSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// NewValidatorSetSigCaller creates a new read-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetSigCaller, error) {
	contract, err := bindValidatorSetSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigCaller{contract: contract}, nil
}

// NewValidatorSetSigTransactor creates a new write-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetSigTransactor, error) {
	contract, err := bindValidatorSetSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigTransactor{contract: contract}, nil
}

// NewValidatorSetSigFilterer creates a new log filterer instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetSigFilterer, error) {
	contract, err := bindValidatorSetSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigFilterer{contract: contract}, nil
}

// bindValidatorSetSig binds a generic wrapper to an already deployed contract.
func bindValidatorSetSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.ValidatorSetSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCaller) Nonces(opts *bind.CallOpts, targetContractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "nonces", targetContractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidateMessage(opts *bind.CallOpts, message ValidatorSetSigMessage) error {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validateMessage", message)

	if err != nil {
		return err
	}

	return err

}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidatorBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validatorBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactor) ExecuteCall(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.contract.Transact(opts, "executeCall", messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactorSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigSession) Receive() (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.Receive(&_ValidatorSetSig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.Receive(&_ValidatorSetSig.TransactOpts)
}

// ValidatorSetSigDeliveredIterator is returned from FilterDelivered and is used to iterate over the raw logs and unpacked data for Delivered events raised by the ValidatorSetSig contract.
type ValidatorSetSigDeliveredIterator struct {
	Event *ValidatorSetSigDelivered // Event containing the contract specifics and raw log

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
func (it *ValidatorSetSigDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetSigDelivered)
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
		it.Event = new(ValidatorSetSigDelivered)
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
func (it *ValidatorSetSigDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetSigDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetSigDelivered represents a Delivered event raised by the ValidatorSetSig contract.
type ValidatorSetSigDelivered struct {
	TargetContractAddress common.Address
	Nonce                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelivered is a free log retrieval operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) FilterDelivered(opts *bind.FilterOpts, targetContractAddress []common.Address, nonce []*big.Int) (*ValidatorSetSigDeliveredIterator, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.FilterLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigDeliveredIterator{contract: _ValidatorSetSig.contract, event: "Delivered", logs: logs, sub: sub}, nil
}

// WatchDelivered is a free log subscription operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) WatchDelivered(opts *bind.WatchOpts, sink chan<- *ValidatorSetSigDelivered, targetContractAddress []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.WatchLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetSigDelivered)
				if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
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

// ParseDelivered is a log parse operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) ParseDelivered(log types.Log) (*ValidatorSetSigDelivered, error) {
	event := new(ValidatorSetSigDelivered)
	if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
