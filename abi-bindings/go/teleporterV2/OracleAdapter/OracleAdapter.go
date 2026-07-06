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

// OracleAdapterMetaData contains all meta data concerning the OracleAdapter contract.
var OracleAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowed\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProcessed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structTeleporterMessageV2\",\"components\":[{\"name\":\"messageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originTeleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"internalType\":\"structTeleporterMessageReceipt[]\",\"components\":[{\"name\":\"receivedMessageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"relayerRewardAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedSource\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structTeleporterICMMessage\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structTeleporterMessageV2\",\"components\":[{\"name\":\"messageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originTeleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"internalType\":\"structTeleporterMessageReceipt[]\",\"components\":[{\"name\":\"receivedMessageNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"relayerRewardAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sourceNetworkID\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSourceUpdated\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleMessageVerified\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyProcessed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceNotAllowed\",\"inputs\":[{\"name\":\"sourceType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"want\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051611002380380611002833981016040819052602b91609d565b6001600160a01b03811660515760405163d92e233d60e01b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b03831690811782556040519091907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060c8565b5f6020828403121560ac575f5ffd5b81516001600160a01b038116811460c1575f5ffd5b9392505050565b610f2d806100d55f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c8063b771b3bc11610058578063b771b3bc1461011b578063eb97cd2c14610129578063f1faff001461013c578063f2fde38b1461014f575f5ffd5b806354f12a2c146100895780637a22977c146100b15780638da5cb5b146100dc5780639b46d5e414610106575b5f5ffd5b61009c61009736600461073c565b610162565b60405190151581526020015b60405180910390f35b61009c6100bf3660046107bd565b6001600160401b03165f9081526002602052604090205460ff1690565b5f546100ee906001600160a01b031681565b6040516001600160a01b0390911681526020016100a8565b6101196101143660046107ec565b6101b0565b005b6100ee6005600160991b0181565b61011961013736600461086d565b61026f565b61009c61014a3660046108a4565b610305565b61011961015d3660046108fe565b61064f565b5f60018585604051610175929190610919565b90815260200160405180910390208383604051610193929190610919565b9081526040519081900360200190205460ff169050949350505050565b5f546001600160a01b031633146101d9576040516282b42960e81b815260040160405180910390fd5b80600186866040516101ec929190610919565b9081526020016040518091039020848460405161020a929190610919565b908152604051908190036020018120805492151560ff19909316929092179091557f394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee906102609087908790879087908790610950565b60405180910390a15050505050565b6005600160991b016001600160a01b031663ee5b48eb826040516020016102969190610ae4565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016102c19190610c2b565b6020604051808303815f875af11580156102dd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103019190610c3d565b5050565b5f806103146060840184610c54565b8101906103219190610c96565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa15801561036d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103949190810190610d76565b91509150806103b657604051636b2f19e960e01b815260040160405180910390fd5b5f6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061041e9190610c3d565b8351909150811461045457825160405163344de78960e01b81526004810191909152602481018290526044015b60405180910390fd5b6104a56040518060c0016040528060608152602001606081526020015f6001600160a01b031681526020015f6001600160401b031681526020015f6001600160401b03168152602001606081525090565b5f5f5f5f5f5f89604001518060200190518101906104c39190610e14565b6040805160c08101825287815260208101969096526001600160a01b03909416858501526001600160401b0392831660608601529116608084015260a0830152519098506001975061051d96509094509250610ed8915050565b9081526020016040518091039020816020015160405161053d9190610ed8565b9081526040519081900360200190205460ff166105765780516020820151604051630d708e6360e41b815261044b929190600401610ef3565b60808101516001600160401b03165f9081526002602052604090205460ff16156105c4576080810151604051633d16f9d960e01b81526001600160401b03909116600482015260240161044b565b6080810180516001600160401b039081165f90815260026020908152604091829020805460ff1916600117905581850151935185519186015192516001600160a01b03909516949316927fce36b341d6a7eec024aa12775c1c08fe81117add41f1b0f0425eb04dc27beb569261063a9291610ef3565b60405180910390a35060019695505050505050565b5f546001600160a01b03163314610678576040516282b42960e81b815260040160405180910390fd5b6001600160a01b03811661069f5760405163d92e233d60e01b815260040160405180910390fd5b5f80546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f83601f840112610708575f5ffd5b5081356001600160401b0381111561071e575f5ffd5b602083019150836020828501011115610735575f5ffd5b9250929050565b5f5f5f5f6040858703121561074f575f5ffd5b84356001600160401b03811115610764575f5ffd5b610770878288016106f8565b90955093505060208501356001600160401b0381111561078e575f5ffd5b61079a878288016106f8565b95989497509550505050565b6001600160401b03811681146107ba575f5ffd5b50565b5f602082840312156107cd575f5ffd5b81356107d8816107a6565b9392505050565b80151581146107ba575f5ffd5b5f5f5f5f5f60608688031215610800575f5ffd5b85356001600160401b03811115610815575f5ffd5b610821888289016106f8565b90965094505060208601356001600160401b0381111561083f575f5ffd5b61084b888289016106f8565b909450925050604086013561085f816107df565b809150509295509295909350565b5f6020828403121561087d575f5ffd5b81356001600160401b03811115610892575f5ffd5b820161012081850312156107d8575f5ffd5b5f602082840312156108b4575f5ffd5b81356001600160401b038111156108c9575f5ffd5b8201608081850312156107d8575f5ffd5b6001600160a01b03811681146107ba575f5ffd5b80356108f9816108da565b919050565b5f6020828403121561090e575f5ffd5b81356107d8816108da565b818382375f9101908152919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f610963606083018789610928565b8281036020840152610976818688610928565b91505082151560408301529695505050505050565b5f5f8335601e198436030181126109a0575f5ffd5b83016020810192503590506001600160401b038111156109be575f5ffd5b8060051b3603821315610735575f5ffd5b8183526020830192505f815f5b84811015610a0d5781356109ef816108da565b6001600160a01b0316865260209586019591909101906001016109dc565b5093949350505050565b5f5f8335601e19843603018112610a2c575f5ffd5b83016020810192503590506001600160401b03811115610a4a575f5ffd5b8060061b3603821315610735575f5ffd5b8183526020830192505f815f5b84811015610a0d57813586526020820135610a82816108da565b6001600160a01b031660208701526040958601959190910190600101610a68565b5f5f8335601e19843603018112610ab8575f5ffd5b83016020810192503590506001600160401b03811115610ad6575f5ffd5b803603821315610735575f5ffd5b60208082528235828201525f90610afc9084016108ee565b6001600160a01b038116604084015250610b18604084016108ee565b6001600160a01b0381166060840152506060830135608083810191909152610b419084016108ee565b6001600160a01b03811660a08401525060a083013560c083810191909152610b6b9084018461098b565b61012060e0850152610b82610140850182846109cf565b915050610b9260e0850185610a17565b848303601f1901610100860152610baa838284610a5b565b92505050610bbc610100850185610aa3565b848303601f1901610120860152610bd4838284610928565b9695505050505050565b5f5b83811015610bf8578181015183820152602001610be0565b50505f910152565b5f8151808452610c17816020860160208601610bde565b601f01601f19169290920160200192915050565b602081525f6107d86020830184610c00565b5f60208284031215610c4d575f5ffd5b5051919050565b5f5f8335601e19843603018112610c69575f5ffd5b8301803591506001600160401b03821115610c82575f5ffd5b602001915036819003821315610735575f5ffd5b5f60208284031215610ca6575f5ffd5b813563ffffffff811681146107d8575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715610cef57610cef610cb9565b60405290565b5f82601f830112610d04575f5ffd5b8151602083015f5f6001600160401b03841115610d2357610d23610cb9565b50604051601f19601f85018116603f011681018181106001600160401b0382111715610d5157610d51610cb9565b604052838152905080828401871015610d68575f5ffd5b610bd4846020830185610bde565b5f5f60408385031215610d87575f5ffd5b82516001600160401b03811115610d9c575f5ffd5b830160608186031215610dad575f5ffd5b610db5610ccd565b815181526020820151610dc7816108da565b602082015260408201516001600160401b03811115610de4575f5ffd5b610df087828501610cf5565b60408301525080935050506020830151610e09816107df565b809150509250929050565b5f5f5f5f5f5f60c08789031215610e29575f5ffd5b86516001600160401b03811115610e3e575f5ffd5b610e4a89828a01610cf5565b96505060208701516001600160401b03811115610e65575f5ffd5b610e7189828a01610cf5565b9550506040870151610e82816108da565b6060880151909450610e93816107a6565b6080880151909350610ea4816107a6565b60a08801519092506001600160401b03811115610ebf575f5ffd5b610ecb89828a01610cf5565b9150509295509295509295565b5f8251610ee9818460208701610bde565b9190910192915050565b604081525f610f056040830185610c00565b8281036020840152610f178185610c00565b9594505050505056fea164736f6c634300081e000a",
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

// IsProcessed is a free data retrieval call binding the contract method 0x7a22977c.
//
// Solidity: function isProcessed(uint64 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterCaller) IsProcessed(opts *bind.CallOpts, nonce uint64) (bool, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "isProcessed", nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProcessed is a free data retrieval call binding the contract method 0x7a22977c.
//
// Solidity: function isProcessed(uint64 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterSession) IsProcessed(nonce uint64) (bool, error) {
	return _OracleAdapter.Contract.IsProcessed(&_OracleAdapter.CallOpts, nonce)
}

// IsProcessed is a free data retrieval call binding the contract method 0x7a22977c.
//
// Solidity: function isProcessed(uint64 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterCallerSession) IsProcessed(nonce uint64) (bool, error) {
	return _OracleAdapter.Contract.IsProcessed(&_OracleAdapter.CallOpts, nonce)
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

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_OracleAdapter *OracleAdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_OracleAdapter *OracleAdapterSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _OracleAdapter.Contract.SendMessage(&_OracleAdapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_OracleAdapter *OracleAdapterTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _OracleAdapter.Contract.SendMessage(&_OracleAdapter.TransactOpts, message)
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

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_OracleAdapter *OracleAdapterTransactor) VerifyMessage(opts *bind.TransactOpts, message TeleporterICMMessage) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "verifyMessage", message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_OracleAdapter *OracleAdapterSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _OracleAdapter.Contract.VerifyMessage(&_OracleAdapter.TransactOpts, message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_OracleAdapter *OracleAdapterTransactorSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _OracleAdapter.Contract.VerifyMessage(&_OracleAdapter.TransactOpts, message)
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

// OracleAdapterOracleMessageVerifiedIterator is returned from FilterOracleMessageVerified and is used to iterate over the raw logs and unpacked data for OracleMessageVerified events raised by the OracleAdapter contract.
type OracleAdapterOracleMessageVerifiedIterator struct {
	Event *OracleAdapterOracleMessageVerified // Event containing the contract specifics and raw log

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
func (it *OracleAdapterOracleMessageVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdapterOracleMessageVerified)
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
		it.Event = new(OracleAdapterOracleMessageVerified)
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
func (it *OracleAdapterOracleMessageVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdapterOracleMessageVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdapterOracleMessageVerified represents a OracleMessageVerified event raised by the OracleAdapter contract.
type OracleAdapterOracleMessageVerified struct {
	Nonce         uint64
	SourceType    string
	SourceAddress string
	DestContract  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleMessageVerified is a free log retrieval operation binding the contract event 0xce36b341d6a7eec024aa12775c1c08fe81117add41f1b0f0425eb04dc27beb56.
//
// Solidity: event OracleMessageVerified(uint64 indexed nonce, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) FilterOracleMessageVerified(opts *bind.FilterOpts, nonce []uint64, destContract []common.Address) (*OracleAdapterOracleMessageVerifiedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "OracleMessageVerified", nonceRule, destContractRule)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterOracleMessageVerifiedIterator{contract: _OracleAdapter.contract, event: "OracleMessageVerified", logs: logs, sub: sub}, nil
}

// WatchOracleMessageVerified is a free log subscription operation binding the contract event 0xce36b341d6a7eec024aa12775c1c08fe81117add41f1b0f0425eb04dc27beb56.
//
// Solidity: event OracleMessageVerified(uint64 indexed nonce, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) WatchOracleMessageVerified(opts *bind.WatchOpts, sink chan<- *OracleAdapterOracleMessageVerified, nonce []uint64, destContract []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "OracleMessageVerified", nonceRule, destContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdapterOracleMessageVerified)
				if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageVerified", log); err != nil {
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

// ParseOracleMessageVerified is a log parse operation binding the contract event 0xce36b341d6a7eec024aa12775c1c08fe81117add41f1b0f0425eb04dc27beb56.
//
// Solidity: event OracleMessageVerified(uint64 indexed nonce, string sourceType, string sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) ParseOracleMessageVerified(log types.Log) (*OracleAdapterOracleMessageVerified, error) {
	event := new(OracleAdapterOracleMessageVerified)
	if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageVerified", log); err != nil {
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
