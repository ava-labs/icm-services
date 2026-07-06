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
	Bin: "0x6080604052348015600e575f5ffd5b50604051610ffc380380610ffc833981016040819052602b91609d565b6001600160a01b03811660515760405163d92e233d60e01b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b03831690811782556040519091907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060c8565b5f6020828403121560ac575f5ffd5b81516001600160a01b038116811460c1575f5ffd5b9392505050565b610f27806100d55f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c8063b771b3bc11610058578063b771b3bc1461011b578063eb97cd2c14610129578063f1faff001461013c578063f2fde38b1461014f575f5ffd5b806354f12a2c146100895780637a22977c146100b15780638da5cb5b146100dc5780639b46d5e414610106575b5f5ffd5b61009c61009736600461072f565b610162565b60405190151581526020015b60405180910390f35b61009c6100bf3660046107b0565b6001600160401b03165f9081526002602052604090205460ff1690565b5f546100ee906001600160a01b031681565b6040516001600160a01b0390911681526020016100a8565b6101196101143660046107df565b6101ae565b005b6100ee6005600160991b0181565b610119610137366004610860565b610266565b61009c61014a366004610897565b6102fc565b61011961015d3660046108f1565b610642565b5f60015f8686868660405160200161017d9493929190610934565b60408051808303601f190181529181528151602092830120835290820192909252015f205460ff1695945050505050565b5f546001600160a01b031633146101d7576040516282b42960e81b815260040160405180910390fd5b5f858585856040516020016101ef9493929190610934565b60408051808303601f1901815282825280516020918201205f8181526001909252919020805460ff191685151517905591507f394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee906102569088908890889088908890610965565b60405180910390a1505050505050565b6005600160991b016001600160a01b031663ee5b48eb8260405160200161028d9190610af9565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016102b89190610c40565b6020604051808303815f875af11580156102d4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102f89190610c52565b5050565b5f8061030b6060840184610c69565b8101906103189190610cab565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610364573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261038b9190810190610d8b565b91509150806103ad57604051636b2f19e960e01b815260040160405180910390fd5b5f6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103f1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104159190610c52565b8351909150811461044b57825160405163344de78960e01b81526004810191909152602481018290526044015b60405180910390fd5b61049c6040518060c0016040528060608152602001606081526020015f6001600160a01b031681526020015f6001600160401b031681526020015f6001600160401b03168152602001606081525090565b5f5f5f5f5f5f89604001518060200190518101906104ba9190610e29565b6040805160c08101825287815260208082018890526001600160a01b03909616818301526001600160401b03948516606082015293909216608084015260a083015251909a505f99506105189850929650909450019150610eed9050565b60408051601f1981840301815291815281516020928301205f818152600190935291205490915060ff166105685781516020830151604051630d708e6360e41b8152610442929190600401610eed565b60808201516001600160401b03165f9081526002602052604090205460ff16156105b6576080820151604051633d16f9d960e01b81526001600160401b039091166004820152602401610442565b6080820180516001600160401b039081165f90815260026020908152604091829020805460ff1916600117905581860151935186519187015192516001600160a01b03909516949316927fce36b341d6a7eec024aa12775c1c08fe81117add41f1b0f0425eb04dc27beb569261062c9291610eed565b60405180910390a3506001979650505050505050565b5f546001600160a01b0316331461066b576040516282b42960e81b815260040160405180910390fd5b6001600160a01b0381166106925760405163d92e233d60e01b815260040160405180910390fd5b5f80546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35f80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f83601f8401126106fb575f5ffd5b5081356001600160401b03811115610711575f5ffd5b602083019150836020828501011115610728575f5ffd5b9250929050565b5f5f5f5f60408587031215610742575f5ffd5b84356001600160401b03811115610757575f5ffd5b610763878288016106eb565b90955093505060208501356001600160401b03811115610781575f5ffd5b61078d878288016106eb565b95989497509550505050565b6001600160401b03811681146107ad575f5ffd5b50565b5f602082840312156107c0575f5ffd5b81356107cb81610799565b9392505050565b80151581146107ad575f5ffd5b5f5f5f5f5f606086880312156107f3575f5ffd5b85356001600160401b03811115610808575f5ffd5b610814888289016106eb565b90965094505060208601356001600160401b03811115610832575f5ffd5b61083e888289016106eb565b9094509250506040860135610852816107d2565b809150509295509295909350565b5f60208284031215610870575f5ffd5b81356001600160401b03811115610885575f5ffd5b820161012081850312156107cb575f5ffd5b5f602082840312156108a7575f5ffd5b81356001600160401b038111156108bc575f5ffd5b8201608081850312156107cb575f5ffd5b6001600160a01b03811681146107ad575f5ffd5b80356108ec816108cd565b919050565b5f60208284031215610901575f5ffd5b81356107cb816108cd565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f61094760408301868861090c565b828103602084015261095a81858761090c565b979650505050505050565b606081525f61097860608301878961090c565b828103602084015261098b81868861090c565b91505082151560408301529695505050505050565b5f5f8335601e198436030181126109b5575f5ffd5b83016020810192503590506001600160401b038111156109d3575f5ffd5b8060051b3603821315610728575f5ffd5b8183526020830192505f815f5b84811015610a22578135610a04816108cd565b6001600160a01b0316865260209586019591909101906001016109f1565b5093949350505050565b5f5f8335601e19843603018112610a41575f5ffd5b83016020810192503590506001600160401b03811115610a5f575f5ffd5b8060061b3603821315610728575f5ffd5b8183526020830192505f815f5b84811015610a2257813586526020820135610a97816108cd565b6001600160a01b031660208701526040958601959190910190600101610a7d565b5f5f8335601e19843603018112610acd575f5ffd5b83016020810192503590506001600160401b03811115610aeb575f5ffd5b803603821315610728575f5ffd5b60208082528235828201525f90610b119084016108e1565b6001600160a01b038116604084015250610b2d604084016108e1565b6001600160a01b0381166060840152506060830135608083810191909152610b569084016108e1565b6001600160a01b03811660a08401525060a083013560c083810191909152610b80908401846109a0565b61012060e0850152610b97610140850182846109e4565b915050610ba760e0850185610a2c565b848303601f1901610100860152610bbf838284610a70565b92505050610bd1610100850185610ab8565b848303601f1901610120860152610be983828461090c565b9695505050505050565b5f5b83811015610c0d578181015183820152602001610bf5565b50505f910152565b5f8151808452610c2c816020860160208601610bf3565b601f01601f19169290920160200192915050565b602081525f6107cb6020830184610c15565b5f60208284031215610c62575f5ffd5b5051919050565b5f5f8335601e19843603018112610c7e575f5ffd5b8301803591506001600160401b03821115610c97575f5ffd5b602001915036819003821315610728575f5ffd5b5f60208284031215610cbb575f5ffd5b813563ffffffff811681146107cb575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715610d0457610d04610cce565b60405290565b5f82601f830112610d19575f5ffd5b8151602083015f5f6001600160401b03841115610d3857610d38610cce565b50604051601f19601f85018116603f011681018181106001600160401b0382111715610d6657610d66610cce565b604052838152905080828401871015610d7d575f5ffd5b610be9846020830185610bf3565b5f5f60408385031215610d9c575f5ffd5b82516001600160401b03811115610db1575f5ffd5b830160608186031215610dc2575f5ffd5b610dca610ce2565b815181526020820151610ddc816108cd565b602082015260408201516001600160401b03811115610df9575f5ffd5b610e0587828501610d0a565b60408301525080935050506020830151610e1e816107d2565b809150509250929050565b5f5f5f5f5f5f60c08789031215610e3e575f5ffd5b86516001600160401b03811115610e53575f5ffd5b610e5f89828a01610d0a565b96505060208701516001600160401b03811115610e7a575f5ffd5b610e8689828a01610d0a565b9550506040870151610e97816108cd565b6060880151909450610ea881610799565b6080880151909350610eb981610799565b60a08801519092506001600160401b03811115610ed4575f5ffd5b610ee089828a01610d0a565b9150509295509295509295565b604081525f610eff6040830185610c15565b8281036020840152610f118185610c15565b9594505050505056fea164736f6c634300081e000a",
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
