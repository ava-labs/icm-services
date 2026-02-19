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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"trustedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161080538038061080583398101604081905261002e91610099565b6001600160a01b0381166100885760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964207369676e6572206164647265737300000000000000000000604482015260640160405180910390fd5b6001600160a01b03166080526100c6565b5f602082840312156100a9575f5ffd5b81516001600160a01b03811681146100bf575f5ffd5b9392505050565b6080516107216100e45f395f81816065015261016301526107215ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063f1faff0014610038578063f74d548014610060575b5f5ffd5b61004b61004636600461039a565b61009f565b60405190151581526020015b60405180910390f35b6100877f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610057565b5f806100ab83806103d8565b83604001356040516020016100c1929190610599565b6040516020818303038152906040528051906020012090505f610110827f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f908152601c91909152603c902090565b90505f61015f6101236060870187610694565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525086939250506101a09050565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316149350505050919050565b5f5f5f5f6101ae86866101c8565b9250925092506101be8282610211565b5090949350505050565b5f5f5f83516041036101ff576020840151604085015160608601515f1a6101f1888285856102d2565b95509550955050505061020a565b505081515f91506002905b9250925092565b5f826003811115610224576102246106d7565b0361022d575050565b6001826003811115610241576102416106d7565b0361025f5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610273576102736106d7565b036102995760405163fce698f760e01b8152600481018290526024015b60405180910390fd5b60038260038111156102ad576102ad6106d7565b036102ce576040516335e2f38360e21b815260048101829052602401610290565b5050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561030b57505f91506003905082610390565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561035c573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661038757505f925060019150829050610390565b92505f91508190505b9450945094915050565b5f602082840312156103aa575f5ffd5b813567ffffffffffffffff8111156103c0575f5ffd5b8201608081850312156103d1575f5ffd5b9392505050565b5f823561011e198336030181126103ed575f5ffd5b9190910192915050565b80356001600160a01b038116811461040d575f5ffd5b919050565b5f5f8335601e19843603018112610427575f5ffd5b830160208101925035905067ffffffffffffffff811115610446575f5ffd5b8060051b3603821315610457575f5ffd5b9250929050565b8183526020830192505f815f5b8481101561049a576001600160a01b03610484836103f7565b168652602095860195919091019060010161046b565b5093949350505050565b5f5f8335601e198436030181126104b9575f5ffd5b830160208101925035905067ffffffffffffffff8111156104d8575f5ffd5b8060061b3603821315610457575f5ffd5b8183526020830192505f815f5b8481101561049a57813586526001600160a01b03610516602084016103f7565b16602087015260409586019591909101906001016104f6565b5f5f8335601e19843603018112610544575f5ffd5b830160208101925035905067ffffffffffffffff811115610563575f5ffd5b803603821315610457575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408082528335908201525f6105b1602085016103f7565b6001600160a01b031660608301526105cb604085016103f7565b6001600160a01b038116608084015250606084013560a08301526105f1608085016103f7565b6001600160a01b03811660c08401525060a084013560e083015261061860c0850185610412565b6101206101008501526106306101608501828461045e565b91505061064060e08601866104a4565b848303603f19016101208601526106588382846104e9565b9250505061066a61010086018661052f565b848303603f1901610140860152610682838284610571565b93505050508260208301529392505050565b5f5f8335601e198436030181126106a9575f5ffd5b83018035915067ffffffffffffffff8211156106c3575f5ffd5b602001915036819003821315610457575f5ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220539e4937f23897725129ce16388a6b09aa10c1097db6fab899a9d5e7ffb6700664736f6c634300081e0033",
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
