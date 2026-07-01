// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracleadapter

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

// OracleMessage is an auto generated low-level Go binding around an user-defined struct.
type OracleMessage struct {
	SourceType        string
	SourceAddress     string
	DestContract      common.Address
	SourceBlockHeight uint64
	Nonce             uint64
	Payload           []byte
}

// OracleAdapterMetaData contains all meta data concerning the OracleAdapter contract.
var OracleAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowed\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProcessed\",\"inputs\":[{\"name\":\"messageID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveOracleMessage\",\"inputs\":[{\"name\":\"warpIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"oracleMsg\",\"type\":\"tuple\",\"internalType\":\"structOracleMessage\",\"components\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedSource\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSourceUpdated\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleMessageReceived\",\"inputs\":[{\"name\":\"messageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyProcessed\",\"inputs\":[{\"name\":\"messageID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceNotAllowed\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"want\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051610d90380380610d90833981016040819052602b91609d565b6001600160a01b03811660515760405163d92e233d60e01b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b03831690811782556040519091907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060c8565b5f6020828403121560ac575f5ffd5b81516001600160a01b038116811460c1575f5ffd5b9392505050565b610cbb806100d55f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80638da5cb5b116100585780638da5cb5b146100dd5780639b46d5e414610107578063b771b3bc1461011a578063f2fde38b14610128575f5ffd5b806311c168961461007e5780631b6a3757146100b557806354f12a2c146100ca575b5f5ffd5b6100a061008c36600461071c565b5f9081526002602052604090205460ff1690565b60405190151581526020015b60405180910390f35b6100c86100c3366004610733565b61013b565b005b6100a06100d83660046107d1565b61056f565b5f546100ef906001600160a01b031681565b6040516001600160a01b0390911681526020016100ac565b6100c861011536600461084b565b6105bb565b6100ef6005600160991b0181565b6100c86101363660046108e0565b610673565b6040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610184573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101ab919081019061097e565b91509150806101cd57604051636b2f19e960e01b815260040160405180910390fd5b5f6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610211573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102359190610a89565b8351909150811461026b57825160405163344de78960e01b81526004810191909152602481018290526044015b60405180910390fd5b604083015180516020909101205f6102838680610aa0565b6102906020890189610aa0565b6102a060608b0160408c016108e0565b6102b060808c0160608d01610ae2565b6102c060a08d0160808e01610ae2565b6102cd60a08e018e610aa0565b6040516020016102e599989796959493929190610b30565b60405160208183030381529060405280519060200120905080821461031d57604051631d6e22b960e01b815260040160405180910390fd5b5f6103288780610aa0565b61033560208a018a610aa0565b6040516020016103489493929190610ba1565b60408051601f1981840301815291815281516020928301205f818152600190935291205490915060ff166103aa576103808780610aa0565b61038d60208a018a610aa0565b604051630d708e6360e41b81526004016102629493929190610ba1565b5f6103b58880610aa0565b6103c260208b018b610aa0565b6103d260a08d0160808e01610ae2565b6040516020016103e6959493929190610bd2565b60408051601f1981840301815291815281516020928301205f818152600290935291205490915060ff161561043157604051630d1069f360e11b815260048101829052602401610262565b5f8181526002602052604090819020805460ff1916600117905561045b9060608a01908a016108e0565b6001600160a01b0316817f30bdff9e103b09743047ceadaf3b56b42849e322b0ffb9bd4af0c30d4bf28ff86104908b80610aa0565b61049d60208e018e610aa0565b6040516104ad9493929190610ba1565b60405180910390a36104c56060890160408a016108e0565b87516001600160a01b039190911690631c67934b906104e48b80610aa0565b6104f160208e018e610aa0565b8e60800160208101906105049190610ae2565b8f8060a001906105149190610aa0565b6040518963ffffffff1660e01b8152600401610537989796959493929190610c14565b5f604051808303815f87803b15801561054e575f5ffd5b505af1158015610560573d5f5f3e3d5ffd5b50505050505050505050505050565b5f60015f8686868660405160200161058a9493929190610ba1565b60408051808303601f190181529181528151602092830120835290820192909252015f205460ff1695945050505050565b5f546001600160a01b031633146105e4576040516282b42960e81b815260040160405180910390fd5b5f858585856040516020016105fc9493929190610ba1565b60408051808303601f1901815282825280516020918201205f8181526001909252919020805460ff191685151517905591507f394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee906106639088908890889088908890610c73565b60405180910390a1505050505050565b5f546001600160a01b0316331461069c576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0381166106c35760405163d92e233d60e01b815260040160405180910390fd5b5f80546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f6020828403121561072c575f5ffd5b5035919050565b5f5f60408385031215610744575f5ffd5b823563ffffffff81168114610757575f5ffd5b915060208301356001600160401b03811115610771575f5ffd5b830160c08186031215610782575f5ffd5b809150509250929050565b5f5f83601f84011261079d575f5ffd5b5081356001600160401b038111156107b3575f5ffd5b6020830191508360208285010111156107ca575f5ffd5b9250929050565b5f5f5f5f604085870312156107e4575f5ffd5b84356001600160401b038111156107f9575f5ffd5b6108058782880161078d565b90955093505060208501356001600160401b03811115610823575f5ffd5b61082f8782880161078d565b95989497509550505050565b8015158114610848575f5ffd5b50565b5f5f5f5f5f6060868803121561085f575f5ffd5b85356001600160401b03811115610874575f5ffd5b6108808882890161078d565b90965094505060208601356001600160401b0381111561089e575f5ffd5b6108aa8882890161078d565b90945092505060408601356108be8161083b565b809150509295509295909350565b6001600160a01b0381168114610848575f5ffd5b5f602082840312156108f0575f5ffd5b81356108fb816108cc565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561093857610938610902565b60405290565b604051601f8201601f191681016001600160401b038111828210171561096657610966610902565b604052919050565b80516109798161083b565b919050565b5f5f6040838503121561098f575f5ffd5b82516001600160401b038111156109a4575f5ffd5b8301606081860312156109b5575f5ffd5b6109bd610916565b8151815260208201516109cf816108cc565b602082015260408201516001600160401b038111156109ec575f5ffd5b80830192505085601f830112610a00575f5ffd5b81516001600160401b03811115610a1957610a19610902565b610a2c601f8201601f191660200161093e565b818152876020838601011115610a40575f5ffd5b5f5b82811015610a5e57602081860181015183830182015201610a42565b505f6020838301015280604084015250508093505050610a806020840161096e565b90509250929050565b5f60208284031215610a99575f5ffd5b5051919050565b5f5f8335601e19843603018112610ab5575f5ffd5b8301803591506001600160401b03821115610ace575f5ffd5b6020019150368190038213156107ca575f5ffd5b5f60208284031215610af2575f5ffd5b81356001600160401b03811681146108fb575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60c081525f610b4360c083018b8d610b08565b8281036020840152610b56818a8c610b08565b6001600160a01b03891660408501526001600160401b0388811660608601528716608085015283810360a08501529050610b91818587610b08565b9c9b505050505050505050505050565b604081525f610bb4604083018688610b08565b8281036020840152610bc7818587610b08565b979650505050505050565b606081525f610be5606083018789610b08565b8281036020840152610bf8818688610b08565b9150506001600160401b03831660408301529695505050505050565b88815260a060208201525f610c2d60a08301898b610b08565b8281036040840152610c4081888a610b08565b90506001600160401b03861660608401528281036080840152610c64818587610b08565b9b9a5050505050505050505050565b606081525f610c86606083018789610b08565b8281036020840152610c99818688610b08565b9150508215156040830152969550505050505056fea164736f6c634300081e000a",
}

// OracleAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleAdapterMetaData.ABI instead.
var OracleAdapterABI = OracleAdapterMetaData.ABI

// OracleAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleAdapterMetaData.Bin instead.
var OracleAdapterBin = OracleAdapterMetaData.Bin

// DeployOracleAdapter deploys a new Ethereum contract, binding an instance of OracleAdapter to it.
func DeployOracleAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *OracleAdapter, error) {
	parsed, err := OracleAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleAdapterBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleAdapter{OracleAdapterCaller: OracleAdapterCaller{contract: contract}, OracleAdapterTransactor: OracleAdapterTransactor{contract: contract}, OracleAdapterFilterer: OracleAdapterFilterer{contract: contract}}, nil
}

// OracleAdapter is an auto generated Go binding around an Ethereum contract.
type OracleAdapter struct {
	OracleAdapterCaller     // Read-only binding to the contract
	OracleAdapterTransactor // Write-only binding to the contract
	OracleAdapterFilterer   // Log filterer for contract events
}

// OracleAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleAdapterSession struct {
	Contract     *OracleAdapter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleAdapterCallerSession struct {
	Contract *OracleAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OracleAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleAdapterTransactorSession struct {
	Contract     *OracleAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OracleAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleAdapterRaw struct {
	Contract *OracleAdapter // Generic contract binding to access the raw methods on
}

// OracleAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleAdapterCallerRaw struct {
	Contract *OracleAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// OracleAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleAdapterTransactorRaw struct {
	Contract *OracleAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleAdapter creates a new instance of OracleAdapter, bound to a specific deployed contract.
func NewOracleAdapter(address common.Address, backend bind.ContractBackend) (*OracleAdapter, error) {
	contract, err := bindOracleAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleAdapter{OracleAdapterCaller: OracleAdapterCaller{contract: contract}, OracleAdapterTransactor: OracleAdapterTransactor{contract: contract}, OracleAdapterFilterer: OracleAdapterFilterer{contract: contract}}, nil
}

// NewOracleAdapterCaller creates a new read-only instance of OracleAdapter, bound to a specific deployed contract.
func NewOracleAdapterCaller(address common.Address, caller bind.ContractCaller) (*OracleAdapterCaller, error) {
	contract, err := bindOracleAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterCaller{contract: contract}, nil
}

// NewOracleAdapterTransactor creates a new write-only instance of OracleAdapter, bound to a specific deployed contract.
func NewOracleAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleAdapterTransactor, error) {
	contract, err := bindOracleAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterTransactor{contract: contract}, nil
}

// NewOracleAdapterFilterer creates a new log filterer instance of OracleAdapter, bound to a specific deployed contract.
func NewOracleAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleAdapterFilterer, error) {
	contract, err := bindOracleAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterFilterer{contract: contract}, nil
}

// bindOracleAdapter binds a generic wrapper to an already deployed contract.
func bindOracleAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleAdapter *OracleAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleAdapter.Contract.OracleAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleAdapter *OracleAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleAdapter.Contract.OracleAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleAdapter *OracleAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleAdapter.Contract.OracleAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleAdapter *OracleAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleAdapter *OracleAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleAdapter *OracleAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleAdapter.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_OracleAdapter *OracleAdapterCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_OracleAdapter *OracleAdapterSession) WARPMESSENGER() (common.Address, error) {
	return _OracleAdapter.Contract.WARPMESSENGER(&_OracleAdapter.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_OracleAdapter *OracleAdapterCallerSession) WARPMESSENGER() (common.Address, error) {
	return _OracleAdapter.Contract.WARPMESSENGER(&_OracleAdapter.CallOpts)
}

// IsAllowed is a free data retrieval call binding the contract method 0x54f12a2c.
//
// Solidity: function isAllowed(string sourceType, string sourceAddress) view returns(bool)
func (_OracleAdapter *OracleAdapterCaller) IsAllowed(opts *bind.CallOpts, sourceType string, sourceAddress string) (bool, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "isAllowed", sourceType, sourceAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0x54f12a2c.
//
// Solidity: function isAllowed(string sourceType, string sourceAddress) view returns(bool)
func (_OracleAdapter *OracleAdapterSession) IsAllowed(sourceType string, sourceAddress string) (bool, error) {
	return _OracleAdapter.Contract.IsAllowed(&_OracleAdapter.CallOpts, sourceType, sourceAddress)
}

// IsAllowed is a free data retrieval call binding the contract method 0x54f12a2c.
//
// Solidity: function isAllowed(string sourceType, string sourceAddress) view returns(bool)
func (_OracleAdapter *OracleAdapterCallerSession) IsAllowed(sourceType string, sourceAddress string) (bool, error) {
	return _OracleAdapter.Contract.IsAllowed(&_OracleAdapter.CallOpts, sourceType, sourceAddress)
}

// IsProcessed is a free data retrieval call binding the contract method 0x11c16896.
//
// Solidity: function isProcessed(bytes32 messageID) view returns(bool)
func (_OracleAdapter *OracleAdapterCaller) IsProcessed(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "isProcessed", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProcessed is a free data retrieval call binding the contract method 0x11c16896.
//
// Solidity: function isProcessed(bytes32 messageID) view returns(bool)
func (_OracleAdapter *OracleAdapterSession) IsProcessed(messageID [32]byte) (bool, error) {
	return _OracleAdapter.Contract.IsProcessed(&_OracleAdapter.CallOpts, messageID)
}

// IsProcessed is a free data retrieval call binding the contract method 0x11c16896.
//
// Solidity: function isProcessed(bytes32 messageID) view returns(bool)
func (_OracleAdapter *OracleAdapterCallerSession) IsProcessed(messageID [32]byte) (bool, error) {
	return _OracleAdapter.Contract.IsProcessed(&_OracleAdapter.CallOpts, messageID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleAdapter *OracleAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleAdapter *OracleAdapterSession) Owner() (common.Address, error) {
	return _OracleAdapter.Contract.Owner(&_OracleAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleAdapter *OracleAdapterCallerSession) Owner() (common.Address, error) {
	return _OracleAdapter.Contract.Owner(&_OracleAdapter.CallOpts)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1b6a3757.
//
// Solidity: function receiveOracleMessage(uint32 warpIndex, (string,string,address,uint64,uint64,bytes) oracleMsg) returns()
func (_OracleAdapter *OracleAdapterTransactor) ReceiveOracleMessage(opts *bind.TransactOpts, warpIndex uint32, oracleMsg OracleMessage) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "receiveOracleMessage", warpIndex, oracleMsg)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1b6a3757.
//
// Solidity: function receiveOracleMessage(uint32 warpIndex, (string,string,address,uint64,uint64,bytes) oracleMsg) returns()
func (_OracleAdapter *OracleAdapterSession) ReceiveOracleMessage(warpIndex uint32, oracleMsg OracleMessage) (*types.Transaction, error) {
	return _OracleAdapter.Contract.ReceiveOracleMessage(&_OracleAdapter.TransactOpts, warpIndex, oracleMsg)
}

// ReceiveOracleMessage is a paid mutator transaction binding the contract method 0x1b6a3757.
//
// Solidity: function receiveOracleMessage(uint32 warpIndex, (string,string,address,uint64,uint64,bytes) oracleMsg) returns()
func (_OracleAdapter *OracleAdapterTransactorSession) ReceiveOracleMessage(warpIndex uint32, oracleMsg OracleMessage) (*types.Transaction, error) {
	return _OracleAdapter.Contract.ReceiveOracleMessage(&_OracleAdapter.TransactOpts, warpIndex, oracleMsg)
}

// SetAllowedSource is a paid mutator transaction binding the contract method 0x9b46d5e4.
//
// Solidity: function setAllowedSource(string sourceType, string sourceAddress, bool allowed) returns()
func (_OracleAdapter *OracleAdapterTransactor) SetAllowedSource(opts *bind.TransactOpts, sourceType string, sourceAddress string, allowed bool) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "setAllowedSource", sourceType, sourceAddress, allowed)
}

// SetAllowedSource is a paid mutator transaction binding the contract method 0x9b46d5e4.
//
// Solidity: function setAllowedSource(string sourceType, string sourceAddress, bool allowed) returns()
func (_OracleAdapter *OracleAdapterSession) SetAllowedSource(sourceType string, sourceAddress string, allowed bool) (*types.Transaction, error) {
	return _OracleAdapter.Contract.SetAllowedSource(&_OracleAdapter.TransactOpts, sourceType, sourceAddress, allowed)
}

// SetAllowedSource is a paid mutator transaction binding the contract method 0x9b46d5e4.
//
// Solidity: function setAllowedSource(string sourceType, string sourceAddress, bool allowed) returns()
func (_OracleAdapter *OracleAdapterTransactorSession) SetAllowedSource(sourceType string, sourceAddress string, allowed bool) (*types.Transaction, error) {
	return _OracleAdapter.Contract.SetAllowedSource(&_OracleAdapter.TransactOpts, sourceType, sourceAddress, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleAdapter *OracleAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleAdapter *OracleAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleAdapter.Contract.TransferOwnership(&_OracleAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleAdapter *OracleAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleAdapter.Contract.TransferOwnership(&_OracleAdapter.TransactOpts, newOwner)
}

// OracleAdapterAllowedSourceUpdatedIterator is returned from FilterAllowedSourceUpdated and is used to iterate over the raw logs and unpacked data for AllowedSourceUpdated events raised by the OracleAdapter contract.
type OracleAdapterAllowedSourceUpdatedIterator struct {
	Event *OracleAdapterAllowedSourceUpdated // Event containing the contract specifics and raw log

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
func (it *OracleAdapterAllowedSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdapterAllowedSourceUpdated)
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
		it.Event = new(OracleAdapterAllowedSourceUpdated)
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
func (it *OracleAdapterAllowedSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdapterAllowedSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdapterAllowedSourceUpdated represents a AllowedSourceUpdated event raised by the OracleAdapter contract.
type OracleAdapterAllowedSourceUpdated struct {
	SourceType    string
	SourceAddress string
	Allowed       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAllowedSourceUpdated is a free log retrieval operation binding the contract event 0x394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee.
//
// Solidity: event AllowedSourceUpdated(string sourceType, string sourceAddress, bool allowed)
func (_OracleAdapter *OracleAdapterFilterer) FilterAllowedSourceUpdated(opts *bind.FilterOpts) (*OracleAdapterAllowedSourceUpdatedIterator, error) {

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "AllowedSourceUpdated")
	if err != nil {
		return nil, err
	}
	return &OracleAdapterAllowedSourceUpdatedIterator{contract: _OracleAdapter.contract, event: "AllowedSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedSourceUpdated is a free log subscription operation binding the contract event 0x394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee.
//
// Solidity: event AllowedSourceUpdated(string sourceType, string sourceAddress, bool allowed)
func (_OracleAdapter *OracleAdapterFilterer) WatchAllowedSourceUpdated(opts *bind.WatchOpts, sink chan<- *OracleAdapterAllowedSourceUpdated) (event.Subscription, error) {

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "AllowedSourceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdapterAllowedSourceUpdated)
				if err := _OracleAdapter.contract.UnpackLog(event, "AllowedSourceUpdated", log); err != nil {
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

// ParseAllowedSourceUpdated is a log parse operation binding the contract event 0x394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee.
//
// Solidity: event AllowedSourceUpdated(string sourceType, string sourceAddress, bool allowed)
func (_OracleAdapter *OracleAdapterFilterer) ParseAllowedSourceUpdated(log types.Log) (*OracleAdapterAllowedSourceUpdated, error) {
	event := new(OracleAdapterAllowedSourceUpdated)
	if err := _OracleAdapter.contract.UnpackLog(event, "AllowedSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleAdapterOracleMessageReceivedIterator is returned from FilterOracleMessageReceived and is used to iterate over the raw logs and unpacked data for OracleMessageReceived events raised by the OracleAdapter contract.
type OracleAdapterOracleMessageReceivedIterator struct {
	Event *OracleAdapterOracleMessageReceived // Event containing the contract specifics and raw log

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
func (it *OracleAdapterOracleMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdapterOracleMessageReceived)
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
		it.Event = new(OracleAdapterOracleMessageReceived)
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
func (it *OracleAdapterOracleMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdapterOracleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdapterOracleMessageReceived represents a OracleMessageReceived event raised by the OracleAdapter contract.
type OracleAdapterOracleMessageReceived struct {
	MessageID     [32]byte
	SourceType    string
	SourceAddress string
	DestContract  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleMessageReceived is a free log retrieval operation binding the contract event 0x30bdff9e103b09743047ceadaf3b56b42849e322b0ffb9bd4af0c30d4bf28ff8.
//
// Solidity: event OracleMessageReceived(bytes32 indexed messageID, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) FilterOracleMessageReceived(opts *bind.FilterOpts, messageID [][32]byte, destContract []common.Address) (*OracleAdapterOracleMessageReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "OracleMessageReceived", messageIDRule, destContractRule)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterOracleMessageReceivedIterator{contract: _OracleAdapter.contract, event: "OracleMessageReceived", logs: logs, sub: sub}, nil
}

// WatchOracleMessageReceived is a free log subscription operation binding the contract event 0x30bdff9e103b09743047ceadaf3b56b42849e322b0ffb9bd4af0c30d4bf28ff8.
//
// Solidity: event OracleMessageReceived(bytes32 indexed messageID, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) WatchOracleMessageReceived(opts *bind.WatchOpts, sink chan<- *OracleAdapterOracleMessageReceived, messageID [][32]byte, destContract []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "OracleMessageReceived", messageIDRule, destContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdapterOracleMessageReceived)
				if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageReceived", log); err != nil {
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

// ParseOracleMessageReceived is a log parse operation binding the contract event 0x30bdff9e103b09743047ceadaf3b56b42849e322b0ffb9bd4af0c30d4bf28ff8.
//
// Solidity: event OracleMessageReceived(bytes32 indexed messageID, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) ParseOracleMessageReceived(log types.Log) (*OracleAdapterOracleMessageReceived, error) {
	event := new(OracleAdapterOracleMessageReceived)
	if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleAdapter contract.
type OracleAdapterOwnershipTransferredIterator struct {
	Event *OracleAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdapterOwnershipTransferred)
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
		it.Event = new(OracleAdapterOwnershipTransferred)
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
func (it *OracleAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the OracleAdapter contract.
type OracleAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleAdapter *OracleAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterOwnershipTransferredIterator{contract: _OracleAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleAdapter *OracleAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdapterOwnershipTransferred)
				if err := _OracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleAdapter *OracleAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*OracleAdapterOwnershipTransferred, error) {
	event := new(OracleAdapterOwnershipTransferred)
	if err := _OracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
