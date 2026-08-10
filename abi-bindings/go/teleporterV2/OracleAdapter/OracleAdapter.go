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
	SourceBlockHeight *big.Int
	Nonce             *big.Int
	Payload           []byte
}

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"AlreadyProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"}],\"name\":\"SourceNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"got\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"want\",\"type\":\"bytes32\"}],\"name\":\"WrongSourceChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedSourceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"OracleMessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"}],\"name\":\"OracleMessageVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"decodeOracleMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structOracleMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"warpIndex\",\"type\":\"uint32\"}],\"name\":\"encodeOracleAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"encodeOracleMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sourceBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"encodeReceiverPayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setAllowedSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051611404380380611404833981016040819052602b9160b4565b806001600160a01b038116605857604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b605f816065565b505060df565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020828403121560c3575f5ffd5b81516001600160a01b038116811460d8575f5ffd5b9392505050565b611318806100ec5f395ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c80639b46d5e411610088578063bfe7edd711610063578063bfe7edd7146101d5578063eb97cd2c146101f7578063f1faff001461020a578063f2fde38b1461021d575f5ffd5b80639b46d5e414610194578063b771b3bc146101a7578063b870ba53146101b5575f5ffd5b80630ac675f1146100cf5780631e028233146100f857806354f12a2c14610130578063715018a6146101535780637a9310fd1461015d5780638da5cb5b14610170575b5f5ffd5b6100e26100dd3660046108df565b610230565b6040516100ef91906109db565b60405180910390f35b6100e26101063660046109f4565b6040805163ffffffff83166020820152606091016040516020818303038152906040529050919050565b61014361013e366004610a5b565b610268565b60405190151581526020016100ef565b61015b6102b6565b005b6100e261016b366004610ac5565b6102c9565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100ef565b61015b6101a2366004610b70565b6102fe565b61017c6005600160991b0181565b6101c86101c3366004610bf1565b61039c565b6040516100ef9190610c2a565b6101436101e3366004610caf565b5f9081526002602052604090205460ff1690565b61015b610205366004610cc6565b61043a565b610143610218366004610cfd565b610474565b61015b61022b366004610d33565b61071d565b606086868686868660405160200161024d96959493929190610d4e565b60405160208183030381529060405290509695505050505050565b5f6001858560405161027b929190610db0565b90815260200160405180910390208383604051610299929190610db0565b9081526040519081900360200190205460ff169050949350505050565b6102be61075a565b6102c75f610786565b565b606085858585856040516020016102e4959493929190610dbf565b604051602081830303815290604052905095945050505050565b61030661075a565b8060018686604051610319929190610db0565b90815260200160405180910390208484604051610337929190610db0565b908152604051908190036020018120805492151560ff19909316929092179091557f394722fe385073e63adaca9c5034df77a80030d78984a16eab5b57556609b0ee9061038d9087908790879087908790610e37565b60405180910390a15050505050565b6103db6040518060c0016040528060608152602001606081526020015f6001600160a01b031681526020015f81526020015f8152602001606081525090565b5f5f5f5f5f5f878060200190518101906103f59190610eb5565b6040805160c08101825296875260208701959095526001600160a01b03909316938501939093526060840152608083019190915260a082015298975050505050505050565b7fe036eca82262a249913f85dc821a7d3c7d08e90a1fdbd6d2ddc80c474b226dd48160405161046991906110b2565b60405180910390a150565b5f8061048360608401846111ac565b81019061049091906109f4565b6040516306f8253560e41b815263ffffffff821660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156104dc573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261050391908101906111ee565b915091508061052557604051636b2f19e960e01b815260040160405180910390fd5b5f6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610569573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058d919061128c565b835190915081146105c357825160405163344de78960e01b81526004810191909152602481018290526044015b60405180910390fd5b5f6105d1846040015161039c565b90506001815f01516040516105e691906112a3565b9081526020016040518091039020816020015160405161060691906112a3565b9081526040519081900360200190205460ff1661063f5780516020820151604051630d708e6360e41b81526105ba9291906004016112be565b60808101515f9081526002602052604090205460ff161561067b578060800151604051635ad313b960e01b81526004016105ba91815260200190565b60808101515f90815260026020908152604091829020805460ff19166001179055818301519083015191516001600160a01b03909116916106bb916112a3565b604051908190038120835190916106d291906112a3565b60405190819003812060808501518252907f2d99b811a0eb7abb06c26b6a03250e336194b44578a56313b8002d7b1bbc8f499060200160405180910390a45060019695505050505050565b61072561075a565b6001600160a01b03811661074e57604051631e4fbdf760e01b81525f60048201526024016105ba565b61075781610786565b50565b5f546001600160a01b031633146102c75760405163118cdaa760e01b81523360048201526024016105ba565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561080b5761080b6107d5565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610839576108396107d5565b604052919050565b5f6001600160401b03821115610859576108596107d5565b50601f01601f191660200190565b5f82601f830112610876575f5ffd5b8135602083015f61088e61088984610841565b610811565b90508281528583830111156108a1575f5ffd5b828260208301375f92810160200192909252509392505050565b6001600160a01b0381168114610757575f5ffd5b80356108da816108bb565b919050565b5f5f5f5f5f5f60c087890312156108f4575f5ffd5b86356001600160401b03811115610909575f5ffd5b61091589828a01610867565b96505060208701356001600160401b03811115610930575f5ffd5b61093c89828a01610867565b955050604087013561094d816108bb565b9350606087013592506080870135915060a08701356001600160401b03811115610975575f5ffd5b61098189828a01610867565b9150509295509295509295565b5f5b838110156109a8578181015183820152602001610990565b50505f910152565b5f81518084526109c781602086016020860161098e565b601f01601f19169290920160200192915050565b602081525f6109ed60208301846109b0565b9392505050565b5f60208284031215610a04575f5ffd5b813563ffffffff811681146109ed575f5ffd5b5f5f83601f840112610a27575f5ffd5b5081356001600160401b03811115610a3d575f5ffd5b602083019150836020828501011115610a54575f5ffd5b9250929050565b5f5f5f5f60408587031215610a6e575f5ffd5b84356001600160401b03811115610a83575f5ffd5b610a8f87828801610a17565b90955093505060208501356001600160401b03811115610aad575f5ffd5b610ab987828801610a17565b95989497509550505050565b5f5f5f5f5f60a08688031215610ad9575f5ffd5b85356001600160401b03811115610aee575f5ffd5b610afa88828901610867565b95505060208601356001600160401b03811115610b15575f5ffd5b610b2188828901610867565b945050604086013592506060860135915060808601356001600160401b03811115610b4a575f5ffd5b610b5688828901610867565b9150509295509295909350565b8015158114610757575f5ffd5b5f5f5f5f5f60608688031215610b84575f5ffd5b85356001600160401b03811115610b99575f5ffd5b610ba588828901610a17565b90965094505060208601356001600160401b03811115610bc3575f5ffd5b610bcf88828901610a17565b9094509250506040860135610be381610b63565b809150509295509295909350565b5f60208284031215610c01575f5ffd5b81356001600160401b03811115610c16575f5ffd5b610c2284828501610867565b949350505050565b602081525f825160c06020840152610c4560e08401826109b0565b90506020840151601f19848303016040850152610c6282826109b0565b91505060018060a01b03604085015116606084015260608401516080840152608084015160a084015260a0840151601f198483030160c0850152610ca682826109b0565b95945050505050565b5f60208284031215610cbf575f5ffd5b5035919050565b5f60208284031215610cd6575f5ffd5b81356001600160401b03811115610ceb575f5ffd5b820161012081850312156109ed575f5ffd5b5f60208284031215610d0d575f5ffd5b81356001600160401b03811115610d22575f5ffd5b8201608081850312156109ed575f5ffd5b5f60208284031215610d43575f5ffd5b81356109ed816108bb565b60c081525f610d6060c08301896109b0565b8281036020840152610d7281896109b0565b6001600160a01b0388166040850152606084018790526080840186905283810360a08501529050610da381856109b0565b9998505050505050505050565b818382375f9101908152919050565b60a081525f610dd160a08301886109b0565b8281036020840152610de381886109b0565b90508560408401528460608401528281036080840152610e0381856109b0565b98975050505050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b606081525f610e4a606083018789610e0f565b8281036020840152610e5d818688610e0f565b91505082151560408301529695505050505050565b5f82601f830112610e81575f5ffd5b8151602083015f610e9461088984610841565b9050828152858383011115610ea7575f5ffd5b610ca683602083018461098e565b5f5f5f5f5f5f60c08789031215610eca575f5ffd5b86516001600160401b03811115610edf575f5ffd5b610eeb89828a01610e72565b96505060208701516001600160401b03811115610f06575f5ffd5b610f1289828a01610e72565b9550506040870151610f23816108bb565b6060880151608089015160a08a015192965090945092506001600160401b03811115610f4d575f5ffd5b61098189828a01610e72565b5f5f8335601e19843603018112610f6e575f5ffd5b83016020810192503590506001600160401b03811115610f8c575f5ffd5b8060051b3603821315610a54575f5ffd5b8183526020830192505f815f5b84811015610fdb578135610fbd816108bb565b6001600160a01b031686526020958601959190910190600101610faa565b5093949350505050565b5f5f8335601e19843603018112610ffa575f5ffd5b83016020810192503590506001600160401b03811115611018575f5ffd5b8060061b3603821315610a54575f5ffd5b8183526020830192505f815f5b84811015610fdb57813586526020820135611050816108bb565b6001600160a01b031660208701526040958601959190910190600101611036565b5f5f8335601e19843603018112611086575f5ffd5b83016020810192503590506001600160401b038111156110a4575f5ffd5b803603821315610a54575f5ffd5b60208082528235828201525f906110ca9084016108cf565b6001600160a01b0381166040840152506110e6604084016108cf565b6001600160a01b038116606084015250606083013560808381019190915261110f9084016108cf565b6001600160a01b03811660a08401525060a083013560c08381019190915261113990840184610f59565b61012060e085015261115061014085018284610f9d565b91505061116060e0850185610fe5565b848303601f1901610100860152611178838284611029565b9250505061118a610100850185611071565b848303601f19016101208601526111a2838284610e0f565b9695505050505050565b5f5f8335601e198436030181126111c1575f5ffd5b8301803591506001600160401b038211156111da575f5ffd5b602001915036819003821315610a54575f5ffd5b5f5f604083850312156111ff575f5ffd5b82516001600160401b03811115611214575f5ffd5b830160608186031215611225575f5ffd5b61122d6107e9565b81518152602082015161123f816108bb565b602082015260408201516001600160401b0381111561125c575f5ffd5b61126887828501610e72565b6040830152508093505050602083015161128181610b63565b809150509250929050565b5f6020828403121561129c575f5ffd5b5051919050565b5f82516112b481846020870161098e565b9190910192915050565b604081525f6112d060408301856109b0565b8281036020840152610ca681856109b056fea264697066735822122038fe81f8811439900e25bd18abd7cba483d7bb59814275498a8a7183eb74e13064736f6c634300081e0033",
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

// DecodeOracleMessage is a free data retrieval call binding the contract method 0xb870ba53.
//
// Solidity: function decodeOracleMessage(bytes payload) pure returns((string,string,address,uint256,uint256,bytes))
func (_OracleAdapter *OracleAdapterCaller) DecodeOracleMessage(opts *bind.CallOpts, payload []byte) (OracleMessage, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "decodeOracleMessage", payload)

	if err != nil {
		return *new(OracleMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleMessage)).(*OracleMessage)

	return out0, err

}

// DecodeOracleMessage is a free data retrieval call binding the contract method 0xb870ba53.
//
// Solidity: function decodeOracleMessage(bytes payload) pure returns((string,string,address,uint256,uint256,bytes))
func (_OracleAdapter *OracleAdapterSession) DecodeOracleMessage(payload []byte) (OracleMessage, error) {
	return _OracleAdapter.Contract.DecodeOracleMessage(&_OracleAdapter.CallOpts, payload)
}

// DecodeOracleMessage is a free data retrieval call binding the contract method 0xb870ba53.
//
// Solidity: function decodeOracleMessage(bytes payload) pure returns((string,string,address,uint256,uint256,bytes))
func (_OracleAdapter *OracleAdapterCallerSession) DecodeOracleMessage(payload []byte) (OracleMessage, error) {
	return _OracleAdapter.Contract.DecodeOracleMessage(&_OracleAdapter.CallOpts, payload)
}

// EncodeOracleAttestation is a free data retrieval call binding the contract method 0x1e028233.
//
// Solidity: function encodeOracleAttestation(uint32 warpIndex) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCaller) EncodeOracleAttestation(opts *bind.CallOpts, warpIndex uint32) ([]byte, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "encodeOracleAttestation", warpIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeOracleAttestation is a free data retrieval call binding the contract method 0x1e028233.
//
// Solidity: function encodeOracleAttestation(uint32 warpIndex) pure returns(bytes)
func (_OracleAdapter *OracleAdapterSession) EncodeOracleAttestation(warpIndex uint32) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeOracleAttestation(&_OracleAdapter.CallOpts, warpIndex)
}

// EncodeOracleAttestation is a free data retrieval call binding the contract method 0x1e028233.
//
// Solidity: function encodeOracleAttestation(uint32 warpIndex) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCallerSession) EncodeOracleAttestation(warpIndex uint32) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeOracleAttestation(&_OracleAdapter.CallOpts, warpIndex)
}

// EncodeOracleMessage is a free data retrieval call binding the contract method 0x0ac675f1.
//
// Solidity: function encodeOracleMessage(string sourceType, string sourceAddress, address destContract, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCaller) EncodeOracleMessage(opts *bind.CallOpts, sourceType string, sourceAddress string, destContract common.Address, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "encodeOracleMessage", sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeOracleMessage is a free data retrieval call binding the contract method 0x0ac675f1.
//
// Solidity: function encodeOracleMessage(string sourceType, string sourceAddress, address destContract, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterSession) EncodeOracleMessage(sourceType string, sourceAddress string, destContract common.Address, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeOracleMessage(&_OracleAdapter.CallOpts, sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload)
}

// EncodeOracleMessage is a free data retrieval call binding the contract method 0x0ac675f1.
//
// Solidity: function encodeOracleMessage(string sourceType, string sourceAddress, address destContract, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCallerSession) EncodeOracleMessage(sourceType string, sourceAddress string, destContract common.Address, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeOracleMessage(&_OracleAdapter.CallOpts, sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload)
}

// EncodeReceiverPayload is a free data retrieval call binding the contract method 0x7a9310fd.
//
// Solidity: function encodeReceiverPayload(string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCaller) EncodeReceiverPayload(opts *bind.CallOpts, sourceType string, sourceAddress string, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "encodeReceiverPayload", sourceType, sourceAddress, sourceBlockHeight, nonce, payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeReceiverPayload is a free data retrieval call binding the contract method 0x7a9310fd.
//
// Solidity: function encodeReceiverPayload(string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterSession) EncodeReceiverPayload(sourceType string, sourceAddress string, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeReceiverPayload(&_OracleAdapter.CallOpts, sourceType, sourceAddress, sourceBlockHeight, nonce, payload)
}

// EncodeReceiverPayload is a free data retrieval call binding the contract method 0x7a9310fd.
//
// Solidity: function encodeReceiverPayload(string sourceType, string sourceAddress, uint256 sourceBlockHeight, uint256 nonce, bytes payload) pure returns(bytes)
func (_OracleAdapter *OracleAdapterCallerSession) EncodeReceiverPayload(sourceType string, sourceAddress string, sourceBlockHeight *big.Int, nonce *big.Int, payload []byte) ([]byte, error) {
	return _OracleAdapter.Contract.EncodeReceiverPayload(&_OracleAdapter.CallOpts, sourceType, sourceAddress, sourceBlockHeight, nonce, payload)
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

// IsProcessed is a free data retrieval call binding the contract method 0xbfe7edd7.
//
// Solidity: function isProcessed(uint256 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterCaller) IsProcessed(opts *bind.CallOpts, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleAdapter.contract.Call(opts, &out, "isProcessed", nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProcessed is a free data retrieval call binding the contract method 0xbfe7edd7.
//
// Solidity: function isProcessed(uint256 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterSession) IsProcessed(nonce *big.Int) (bool, error) {
	return _OracleAdapter.Contract.IsProcessed(&_OracleAdapter.CallOpts, nonce)
}

// IsProcessed is a free data retrieval call binding the contract method 0xbfe7edd7.
//
// Solidity: function isProcessed(uint256 nonce) view returns(bool)
func (_OracleAdapter *OracleAdapterCallerSession) IsProcessed(nonce *big.Int) (bool, error) {
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleAdapter *OracleAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleAdapter *OracleAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleAdapter.Contract.RenounceOwnership(&_OracleAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleAdapter *OracleAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleAdapter.Contract.RenounceOwnership(&_OracleAdapter.TransactOpts)
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

// OracleAdapterOracleMessageSentIterator is returned from FilterOracleMessageSent and is used to iterate over the raw logs and unpacked data for OracleMessageSent events raised by the OracleAdapter contract.
type OracleAdapterOracleMessageSentIterator struct {
	Event *OracleAdapterOracleMessageSent // Event containing the contract specifics and raw log

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
func (it *OracleAdapterOracleMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAdapterOracleMessageSent)
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
		it.Event = new(OracleAdapterOracleMessageSent)
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
func (it *OracleAdapterOracleMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAdapterOracleMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAdapterOracleMessageSent represents a OracleMessageSent event raised by the OracleAdapter contract.
type OracleAdapterOracleMessageSent struct {
	Message TeleporterMessageV2
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOracleMessageSent is a free log retrieval operation binding the contract event 0xe036eca82262a249913f85dc821a7d3c7d08e90a1fdbd6d2ddc80c474b226dd4.
//
// Solidity: event OracleMessageSent((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_OracleAdapter *OracleAdapterFilterer) FilterOracleMessageSent(opts *bind.FilterOpts) (*OracleAdapterOracleMessageSentIterator, error) {

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "OracleMessageSent")
	if err != nil {
		return nil, err
	}
	return &OracleAdapterOracleMessageSentIterator{contract: _OracleAdapter.contract, event: "OracleMessageSent", logs: logs, sub: sub}, nil
}

// WatchOracleMessageSent is a free log subscription operation binding the contract event 0xe036eca82262a249913f85dc821a7d3c7d08e90a1fdbd6d2ddc80c474b226dd4.
//
// Solidity: event OracleMessageSent((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_OracleAdapter *OracleAdapterFilterer) WatchOracleMessageSent(opts *bind.WatchOpts, sink chan<- *OracleAdapterOracleMessageSent) (event.Subscription, error) {

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "OracleMessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAdapterOracleMessageSent)
				if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageSent", log); err != nil {
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

// ParseOracleMessageSent is a log parse operation binding the contract event 0xe036eca82262a249913f85dc821a7d3c7d08e90a1fdbd6d2ddc80c474b226dd4.
//
// Solidity: event OracleMessageSent((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_OracleAdapter *OracleAdapterFilterer) ParseOracleMessageSent(log types.Log) (*OracleAdapterOracleMessageSent, error) {
	event := new(OracleAdapterOracleMessageSent)
	if err := _OracleAdapter.contract.UnpackLog(event, "OracleMessageSent", log); err != nil {
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
	Nonce         *big.Int
	SourceType    common.Hash
	SourceAddress common.Hash
	DestContract  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleMessageVerified is a free log retrieval operation binding the contract event 0x2d99b811a0eb7abb06c26b6a03250e336194b44578a56313b8002d7b1bbc8f49.
//
// Solidity: event OracleMessageVerified(uint256 nonce, string indexed sourceType, string indexed sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) FilterOracleMessageVerified(opts *bind.FilterOpts, sourceType []string, sourceAddress []string, destContract []common.Address) (*OracleAdapterOracleMessageVerifiedIterator, error) {

	var sourceTypeRule []interface{}
	for _, sourceTypeItem := range sourceType {
		sourceTypeRule = append(sourceTypeRule, sourceTypeItem)
	}
	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}
	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.FilterLogs(opts, "OracleMessageVerified", sourceTypeRule, sourceAddressRule, destContractRule)
	if err != nil {
		return nil, err
	}
	return &OracleAdapterOracleMessageVerifiedIterator{contract: _OracleAdapter.contract, event: "OracleMessageVerified", logs: logs, sub: sub}, nil
}

// WatchOracleMessageVerified is a free log subscription operation binding the contract event 0x2d99b811a0eb7abb06c26b6a03250e336194b44578a56313b8002d7b1bbc8f49.
//
// Solidity: event OracleMessageVerified(uint256 nonce, string indexed sourceType, string indexed sourceAddress, address indexed destContract)
func (_OracleAdapter *OracleAdapterFilterer) WatchOracleMessageVerified(opts *bind.WatchOpts, sink chan<- *OracleAdapterOracleMessageVerified, sourceType []string, sourceAddress []string, destContract []common.Address) (event.Subscription, error) {

	var sourceTypeRule []interface{}
	for _, sourceTypeItem := range sourceType {
		sourceTypeRule = append(sourceTypeRule, sourceTypeItem)
	}
	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}
	var destContractRule []interface{}
	for _, destContractItem := range destContract {
		destContractRule = append(destContractRule, destContractItem)
	}

	logs, sub, err := _OracleAdapter.contract.WatchLogs(opts, "OracleMessageVerified", sourceTypeRule, sourceAddressRule, destContractRule)
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

// ParseOracleMessageVerified is a log parse operation binding the contract event 0x2d99b811a0eb7abb06c26b6a03250e336194b44578a56313b8002d7b1bbc8f49.
//
// Solidity: event OracleMessageVerified(uint256 nonce, string indexed sourceType, string indexed sourceAddress, address indexed destContract)
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
