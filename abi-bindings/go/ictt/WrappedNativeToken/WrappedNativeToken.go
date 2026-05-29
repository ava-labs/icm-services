// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappednativetoken

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

// WrappedNativeTokenMetaData contains all meta data concerning the WrappedNativeToken contract.
var WrappedNativeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608080604052346103ae57610bff803803809161001c82856103b2565b83398101906020818303126103ae578051906001600160401b0382116103ae570181601f820112156103ae578051906001600160401b0382116102c15760405191610071601f8201601f1916602001846103b2565b8083526020830193602082840101116103ae576100f29261009b60219286602061010396016103d5565b604051946702bb930b83832b2160c51b60208701526100d660288784516100c581848401876103d5565b81010301601f1981018852876103b2565b604051948592605760f81b6020850152518092858501906103d5565b81010301601f1981018352826103b2565b81516001600160401b0381116102c157600354600181811c911680156103a4575b60208210146102a357601f8111610341575b50602092601f82116001146102e057928192935f926102d5575b50508160011b915f199060031b1c1916176003555b80516001600160401b0381116102c157600454600181811c911680156102b7575b60208210146102a357601f8111610240575b50602091601f82116001146101e0579181925f926101d5575b50508160011b915f199060031b1c1916176004555b60405161080890816103f78239f35b015190505f806101b1565b601f1982169260045f52805f20915f5b85811061022857508360019510610210575b505050811b016004556101c6565b01515f1960f88460031b161c191690555f8080610202565b919260206001819286850151815501940192016101f0565b60045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f830160051c81019160208410610299575b601f0160051c01905b81811061028e5750610198565b5f8155600101610281565b9091508190610278565b634e487b7160e01b5f52602260045260245ffd5b90607f1690610186565b634e487b7160e01b5f52604160045260245ffd5b015190505f80610150565b601f1982169360035f52805f20915f5b8681106103295750836001959610610311575b505050811b01600355610165565b01515f1960f88460031b161c191690555f8080610303565b919260206001819286850151815501940192016102f0565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c8101916020841061039a575b601f0160051c01905b81811061038f5750610136565b5f8155600101610382565b9091508190610379565b90607f1690610124565b5f80fd5b601f909101601f19168101906001600160401b038211908210176102c157604052565b5f5b8381106103e65750505f910152565b81810151838201526020016103d756fe6080806040526004361015610028575b5036156100205761001e6106bb565b005b61001e6106bb565b5f3560e01c90816306fdde031461056e57508063095ea7b3146104ec57806318160ddd146104cf57806323b872dd146103ef5780632e1a7d4d1461029b578063313ce5671461028057806370a082311461024957806395d89b4114610145578063a9059cbb14610114578063d0e30db0146101015763dd62ed3e146100ad575f61000f565b346100fd5760403660031901126100fd576100c661066d565b6100ce610683565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b5f80fd5b5f3660031901126100fd5761001e6106bb565b346100fd5760403660031901126100fd5761013a61013061066d565b6024359033610764565b602060405160018152f35b346100fd575f3660031901126100fd576040515f6004548060011c9060018116801561023f575b60208310811461022b5782855290811561020757506001146101a9575b6101a58361019981850382610699565b60405191829182610626565b0390f35b91905060045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b915f905b8082106101ed57509091508101602001610199610189565b9192600181602092548385880101520191019092916101d5565b60ff191660208086019190915291151560051b840190910191506101999050610189565b634e487b7160e01b5f52602260045260245ffd5b91607f169161016c565b346100fd5760203660031901126100fd576001600160a01b0361026a61066d565b165f525f602052602060405f2054604051908152f35b346100fd575f3660031901126100fd57602060405160128152f35b346100fd5760203660031901126100fd5760043533156103dc57335f525f6020528060405f20548181106103c357335f525f6020520360405f205580600254036002555f6040518281527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a36040518181527f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b6560203392a28047106103ad575f80808093335af13d156103a8573d67ffffffffffffffff81116103945760405190610374601f8201601f191660200183610699565b81525f60203d92013e5b1561038557005b63d6bda27560e01b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b61037e565b4763cf47918160e01b5f5260045260245260445ffd5b63391434e360e21b5f523360045260245260445260645ffd5b634b637e8f60e11b5f525f60045260245ffd5b346100fd5760603660031901126100fd5761040861066d565b610410610683565b6001600160a01b0382165f81815260016020818152604080842033855290915290912054919360443593929091810161044f575b5061013a9350610764565b8381106104b45784156104a157331561048e5761013a945f52600160205260405f2060018060a01b0333165f526020528360405f209103905584610444565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b346100fd575f3660031901126100fd576020600254604051908152f35b346100fd5760403660031901126100fd5761050561066d565b6024359033156104a1576001600160a01b031690811561048e57335f52600160205260405f20825f526020528060405f20556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b346100fd575f3660031901126100fd575f6003548060011c9060018116801561061c575b60208310811461022b5782855290811561020757506001146105be576101a58361019981850382610699565b91905060035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b915f905b80821061060257509091508101602001610199610189565b9192600181602092548385880101520191019092916105ea565b91607f1691610592565b9190916020815282518060208301525f5b818110610657575060409293505f838284010152601f8019910116010190565b8060208092870101516040828601015201610637565b600435906001600160a01b03821682036100fd57565b602435906001600160a01b03821682036100fd57565b90601f8019910116810190811067ffffffffffffffff82111761039457604052565b33156107515760025434810180911161073d57600255335f525f60205260405f203481540190556040513481525f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203393a36040513481527fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c60203392a2565b634e487b7160e01b5f52601160045260245ffd5b63ec442f0560e01b5f525f60045260245ffd5b6001600160a01b03169081156103dc576001600160a01b031691821561075157815f525f60205260405f20548181106107e257817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92602092855f525f84520360405f2055845f525f825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffdfea164736f6c634300081e000a",
}

// WrappedNativeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedNativeTokenMetaData.ABI instead.
var WrappedNativeTokenABI = WrappedNativeTokenMetaData.ABI

// WrappedNativeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedNativeTokenMetaData.Bin instead.
var WrappedNativeTokenBin = WrappedNativeTokenMetaData.Bin

// DeployWrappedNativeToken deploys a new Ethereum contract, binding an instance of WrappedNativeToken to it.
func DeployWrappedNativeToken(auth *bind.TransactOpts, backend bind.ContractBackend, symbol string) (common.Address, *types.Transaction, *WrappedNativeToken, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedNativeTokenBin), backend, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// WrappedNativeToken is an auto generated Go binding around an Ethereum contract.
type WrappedNativeToken struct {
	WrappedNativeTokenCaller     // Read-only binding to the contract
	WrappedNativeTokenTransactor // Write-only binding to the contract
	WrappedNativeTokenFilterer   // Log filterer for contract events
}

// WrappedNativeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedNativeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedNativeTokenSession struct {
	Contract     *WrappedNativeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WrappedNativeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedNativeTokenCallerSession struct {
	Contract *WrappedNativeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WrappedNativeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedNativeTokenTransactorSession struct {
	Contract     *WrappedNativeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WrappedNativeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedNativeTokenRaw struct {
	Contract *WrappedNativeToken // Generic contract binding to access the raw methods on
}

// WrappedNativeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCallerRaw struct {
	Contract *WrappedNativeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedNativeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactorRaw struct {
	Contract *WrappedNativeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedNativeToken creates a new instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeToken(address common.Address, backend bind.ContractBackend) (*WrappedNativeToken, error) {
	contract, err := bindWrappedNativeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// NewWrappedNativeTokenCaller creates a new read-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenCaller(address common.Address, caller bind.ContractCaller) (*WrappedNativeTokenCaller, error) {
	contract, err := bindWrappedNativeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenCaller{contract: contract}, nil
}

// NewWrappedNativeTokenTransactor creates a new write-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedNativeTokenTransactor, error) {
	contract, err := bindWrappedNativeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransactor{contract: contract}, nil
}

// NewWrappedNativeTokenFilterer creates a new log filterer instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedNativeTokenFilterer, error) {
	contract, err := bindWrappedNativeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenFilterer{contract: contract}, nil
}

// bindWrappedNativeToken binds a generic wrapper to an already deployed contract.
func bindWrappedNativeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.WrappedNativeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// WrappedNativeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedNativeToken contract.
type WrappedNativeTokenApprovalIterator struct {
	Event *WrappedNativeTokenApproval // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenApproval)
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
		it.Event = new(WrappedNativeTokenApproval)
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
func (it *WrappedNativeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenApproval represents a Approval event raised by the WrappedNativeToken contract.
type WrappedNativeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedNativeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenApprovalIterator{contract: _WrappedNativeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenApproval)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseApproval(log types.Log) (*WrappedNativeTokenApproval, error) {
	event := new(WrappedNativeTokenApproval)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WrappedNativeToken contract.
type WrappedNativeTokenDepositIterator struct {
	Event *WrappedNativeTokenDeposit // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenDeposit)
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
		it.Event = new(WrappedNativeTokenDeposit)
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
func (it *WrappedNativeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenDeposit represents a Deposit event raised by the WrappedNativeToken contract.
type WrappedNativeTokenDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenDepositIterator{contract: _WrappedNativeToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenDeposit)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseDeposit(log types.Log) (*WrappedNativeTokenDeposit, error) {
	event := new(WrappedNativeTokenDeposit)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransferIterator struct {
	Event *WrappedNativeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenTransfer)
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
		it.Event = new(WrappedNativeTokenTransfer)
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
func (it *WrappedNativeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenTransfer represents a Transfer event raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedNativeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransferIterator{contract: _WrappedNativeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenTransfer)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseTransfer(log types.Log) (*WrappedNativeTokenTransfer, error) {
	event := new(WrappedNativeTokenTransfer)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawalIterator struct {
	Event *WrappedNativeTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenWithdrawal)
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
		it.Event = new(WrappedNativeTokenWithdrawal)
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
func (it *WrappedNativeTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenWithdrawal represents a Withdrawal event raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenWithdrawalIterator{contract: _WrappedNativeToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenWithdrawal)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseWithdrawal(log types.Log) (*WrappedNativeTokenWithdrawal, error) {
	event := new(WrappedNativeTokenWithdrawal)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
