// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkvalidatorsetregistry

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

// ICMMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMMessage struct {
	RawMessage         []byte
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	Attestation        []byte
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

// ValidatorSetMerkleCommitment is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetMerkleCommitment struct {
	AvalancheBlockchainID [32]byte
	Root                  [32]byte
	TotalWeight           uint64
	PChainHeight          uint64
	PChainTimestamp       uint64
}

// ZKValidatorSetRegistryMetaData contains all meta data concerning the ZKValidatorSetRegistry contract.
var ZKValidatorSetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"pChainID_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pChainGenesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainTotalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowPChainFallback_\",\"type\":\"bool\"},{\"internalType\":\"contractISP1Verifier\",\"name\":\"sp1Verifier_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"attestationProgramVKey_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowPChainFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationProgramVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSetCommitment\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSetMerkleCommitment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"signingChainID\",\"type\":\"bytes32\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sp1Verifier\",\"outputs\":[{\"internalType\":\"contractISP1Verifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052346101f257604051601f61197a38819003918201601f19168301916001600160401b038311848410176101de57808492610120946040528339810103126101f257805163ffffffff811681036101f25760208201519160408101519261006d606083016101f6565b9361007a608084016101f6565b9361008760a085016101f6565b9360c081015180151581036101f25760e0820151916001600160a01b03831683036101f2576101000151926080528460a05260c05260e052610100526040519060a082019282841060018060401b038511176101de57604093845280835260208084019283526001600160401b039788168486019081529688166060850190815295881660808086019182525f9384529183905291859020935184559151600184015594516002929092018054945195516001600160c01b0319909516929096169190911793821b6fffffffffffffffff000000000000000016939093179190921b600160801b600160c01b0316179091555161176f908161020b823960805181818161012d0152610854015260a051818181610508015281816108cc0152610c59015260c0518181816107330152610906015260e051818181610940015261126b01526101005181818161098601526112aa0152f35b634e487b7160e01b5f52604160045260245ffd5b5f80fd5b51906001600160401b03821682036101f25756fe60806040526004361015610011575f80fd5b5f3560e01c806327258b22146109a95780633ee255181461096f57806352a07fa31461092b57806353b4c649146108ef578063541dcba4146108b557806357262e7f1461089c578063580d632b1461087857806368531ed0146108385780637abab0f3146103f45780638d8d23a714610357578063eb97cd2c146101b45763f1faff001461009d575f80fd5b346101b05760203660031901126101b0576004356001600160401b0381116101b05780600401908036039060806003198301126101b0576100e46100df610c57565b610c94565b60448101359161011261010d845f525f6020526001600160401b03600260405f20015416151590565b610a15565b61015a6024830161015561012582610a71565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614610a82565b610a71565b9380359161012219018212156101b05760646101976101a59661018e61018961019f96600436918a0101610d73565b61144c565b90873091610ff1565b930190610ac4565b9061113c565b602060405160018152f35b5f80fd5b346101b05760203660031901126101b0576004356001600160401b0381116101b05761012060031982360301126101b057604481016001600160a01b036101fa82610d16565b1633149081156102c8575b501561028d576020610221610189610242933690600401610d73565b6040518093819263ee5b48eb60e01b83528460048401526024830190610f3f565b03815f6005600160991b015af180156102825761025b57005b602090813d831161027b575b6102718183610b25565b810103126101b057005b503d610267565b6040513d5f823e3d90fd5b60405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b6044820152606490fd5b600491506020906001600160a01b03906102e190610d16565b166040519283809263d67bdd2560e01b82525afa908115610282575f91610315575b506001600160a01b0316331482610205565b90506020813d60201161034f575b8161033060209383610b25565b810103126101b057516001600160a01b03811681036101b05782610303565b3d9150610323565b346101b05760203660031901126101b057610370610cec565b506004355f525f60205260a060405f206001600160401b036040519161039583610af6565b8181549384815281600260018501549460208401958652015493604083019082861682528260806060860195828960401c168752019660801c168652604051978852516020880152511660408601525116606084015251166080820152f35b346101b057610402366109de565b9061040e6100df610c57565b61041d61012560208301610a71565b6104278180610ac4565b9091610431610cec565b5081156108245782356001600160f81b0319161580610805575b156107cd57816006116101b0576006600284013560e01c0361076b57600519820190610475610cec565b5061047e610cec565b94826020116101b0576006850135865260251984016020116101b0576026850135602087019081529560451985016008116101b057604686013560c01c6040820190815295604d1986016008116101b057604e81013560c01c6060830190815295605519016008116101b0576056013560c01c6080820190815293605811610757578051968783147f000000000000000000000000000000000000000000000000000000000000000084149061054a8a5f525f6020526001600160401b03600260405f20015416151590565b6106445750156105e4576105686002936001600160401b0395610b97565b875f525f60205260405f2091518255516001820155019351166fffffffffffffffff00000000000000008454935160401b16916001600160401b0360801b905160801b16926001600160401b0360c01b1617171790557f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce935f80a2005b60405162461bcd60e51b815260206004820152603260248201527f496e697469616c20726567697374726174696f6e206d757374206265207369676044820152713732b210313c903a3432902816a1b430b4b760711b6064820152608490fd5b908115610722575b50156106e557875f525f6020526001600160401b03600260405f20015460401c166001600160401b038751161115610692576105686002936001600160401b0395610b97565b60405162461bcd60e51b815260206004820152602560248201527f426c6f636b20686569676874206f6620636f6d6d69746d656e7420697320746f6044820152641bc81bdb1960da1b6064820152608490fd5b60405162461bcd60e51b815260206004820152601560248201527424b73b30b634b21039b4b3b734b7339031b430b4b760591b6044820152606490fd5b905080610730575b8961064c565b507f000000000000000000000000000000000000000000000000000000000000000061072a565b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152603460248201527f496e76616c69642056616c696461746f725365744d65726b6c65436f6d6d69746044820152731b595b9d081c185e5b1bd859081d1e5c1948125160621b6064820152608490fd5b60405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818dbd91958c8125160821b6044820152606490fd5b5081600110156108245760018301356001600160f81b0319161561044b565b634e487b7160e01b5f52603260045260245ffd5b346101b0575f3660031901126101b057602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101b0575f3660031901126101b0576020610892610c57565b6040519015158152f35b346101b0576108b36108ad366109de565b90610b97565b005b346101b0575f3660031901126101b05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101b0575f3660031901126101b05760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346101b0575f3660031901126101b0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101b0575f3660031901126101b05760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101b05760203660031901126101b05760206108926004355f525f6020526001600160401b03600260405f20015416151590565b60406003198201126101b057600435906001600160401b0382116101b05760809082900360031901126101b0576004019060243590565b15610a1c57565b60405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b6064820152608490fd5b3563ffffffff811681036101b05790565b15610a8957565b60405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606490fd5b903590601e19813603018212156101b057018035906001600160401b0382116101b0576020019181360383136101b057565b60a081019081106001600160401b03821117610b1157604052565b634e487b7160e01b5f52604160045260245ffd5b90601f801991011681019081106001600160401b03821117610b1157604052565b6001600160401b038111610b1157601f01601f191660200190565b929192610b6d82610b46565b91610b7b6040519384610b25565b8294818452818301116101b0578281602093845f960137010152565b90610ba0610c57565b15610c1257610c1091610bcc61010d835f525f6020526001600160401b03600260405f20015416151590565b61019f610c05610be56020840161015561012582610a71565b610bf9610bf28580610ac4565b3691610b61565b90604085013590610f63565b916060810190610ac4565b565b60405162461bcd60e51b815260206004820152601760248201527f502d636861696e206e6f7420696e697469616c697a65640000000000000000006044820152606490fd5b7f00000000000000000000000000000000000000000000000000000000000000005f525f6020526001600160401b03600260405f20015416151590565b15610c9b57565b60405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b6064820152608490fd5b60405190610cf982610af6565b5f6080838281528260208201528260408201528260608201520152565b356001600160a01b03811681036101b05790565b35906001600160a01b03821682036101b057565b6001600160401b038111610b115760051b60200190565b9080601f830112156101b057816020610d7093359101610b61565b90565b9190610120838203126101b0576040519061012082018281106001600160401b03821117610b1157604052819380358352610db060208201610d2a565b6020840152610dc160408201610d2a565b604084015260608101356060840152610ddc60808201610d2a565b608084015260a081013560a084015260c08101356001600160401b0381116101b057810182601f820112156101b057803590610e1782610d3e565b91610e256040519384610b25565b80835260208084019160051b830101918583116101b057602001905b828210610f275750505060c084015260e08101356001600160401b0381116101b057810182601f820112156101b0578035610e7b81610d3e565b91610e896040519384610b25565b81835260208084019260061b820101908582116101b057602001915b818310610edc5750505060e0840152610100810135916001600160401b0383116101b05761010092610ed79201610d55565b910152565b6040838703126101b0576040519060408201908282106001600160401b03831117610b1157604092602092845285358152610f18838701610d2a565b83820152815201920191610ea5565b60208091610f3484610d2a565b815201910190610e41565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9091805191600e83018093116107575781516040515f60208083018290526001600160e01b031960e095861b81166022850152602684019890985295841b87166046830152604a8201819052600160e01b604c8301526050820152911b90931660548401528051610d70926058928592918291018484015e81015f838201520301601f198101835282610b25565b909291906001600160a01b038216156110af57825191602283018093116107575783516040515f602082810182905260e095861b6001600160e01b03199081166022850152602684019990995295851b88166046830152604a820152600160e01b604c820152600560e21b605082015260609290921b6bffffffffffffffffffffffff1916605483015290911b90931660688401528151610d7092606c928592918291018385015e81015f838201520301601f198101835282610b25565b8251600e810193925083106107575781516040515f60208083018290526001600160e01b031960e095861b81166022850152602684019890985295841b87166046830152604a8201819052600160e01b604c8301526050820152911b90931660548401528051610d70926058928592918291018484015e81015f838201520301601f198101835282610b25565b91908201916040818403126101b05780356001600160401b0381116101b05783611167918301610d55565b9260208201356001600160401b0381116101b0576111859201610d55565b926080838051810103126101b05760405191608083018381106001600160401b03821117610b115760405260208401518352604084015160208401908152606085015190604085019182526080860151936001600160401b03851685036101b05760608601948552805f525f60205260405f2095510361140757516001850154036113c2575f602080925193604051918183925191829101835e8101838152039060025afa15610282575f51036113725760026001600160401b038092511692015416806043029060438204036107575781606402916064830403610757571161131d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691823b156101b05760405163020a49e360e51b81527f0000000000000000000000000000000000000000000000000000000000000000600482015260606024820152925f92849283918291611302916112f0906064850190610f3f565b83810360031901604485015290610f3f565b03915afa8015610282576113135750565b5f610c1091610b25565b60405162461bcd60e51b815260206004820152602760248201527f7374616b652d77656967687465642071756f72756d207468726573686f6c64206044820152661b9bdd081b595d60ca1b6064820152608490fd5b60405162461bcd60e51b815260206004820152602260248201527f7075626c69632076616c7565206d6573736167652068617368206d69736d61746044820152610c6d60f31b6064820152608490fd5b60405162461bcd60e51b815260206004820152601a60248201527f7075626c69632076616c756520726f6f74206d69736d617463680000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601e60248201527f7075626c69632076616c756520636861696e204944206d69736d6174636800006044820152606490fd5b805160405190602082015260208152611466604082610b25565b602082015191604051926001600160601b03199060601b16602084015260148352611492603484610b25565b604081015190604051916001600160601b03199060601b166020830152601482526114be603483610b25565b6060810151604051906020820152602081526114db604082610b25565b6080820151604051906001600160601b03199060601b16602082015260148152611506603482610b25565b60a08301519060405191602083015260208252611524604083610b25565b60c08401519283516014810290808204601414901517156107575761154890611709565b84519094905f602c8301602088015b8383106116f2575050505063ffffffff61158d91511661157f6040519687926020840161173b565b03601f198101865285610b25565b60e0850151948551603481029080820460341490151715610757576115b190611709565b95805160208801602083015f915b8383106116cf57505050505163ffffffff169560405196879160208301916115e69261173b565b03601f19810187526115f89087610b25565b6101000151958651604051978891602083015280516020819201604084015e8101604081015f905203604001601f19810188526116359088610b25565b60405198888a9951602081920160208c015e890160208101915f83528051926020849201905e016020015f815281516020819301825e015f815281516020819301825e015f815281516020819301825e015f815281516020819301825e015f815281516020819301825e015f815281516020819301825e015f815281516020819301825e015f815203601f1981018252610d709082610b25565b6020603460019282855180518352015160601b83820152019201920191906115bf565b602060146001928451815201920192019190611557565b9061171382610b46565b6117206040519182610b25565b8281528092611731601f1991610b46565b0190602036910137565b60049260209263ffffffff60e01b9060e01b1682528051928391018483015e01015f81529056fea164736f6c634300081e000a",
}

// ZKValidatorSetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKValidatorSetRegistryMetaData.ABI instead.
var ZKValidatorSetRegistryABI = ZKValidatorSetRegistryMetaData.ABI

// ZKValidatorSetRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKValidatorSetRegistryMetaData.Bin instead.
var ZKValidatorSetRegistryBin = ZKValidatorSetRegistryMetaData.Bin

// DeployZKValidatorSetRegistry deploys a new Ethereum contract, binding an instance of ZKValidatorSetRegistry to it.
func DeployZKValidatorSetRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, avalancheNetworkID_ uint32, pChainID_ [32]byte, pChainGenesisRoot [32]byte, pChainTotalWeight uint64, pChainHeight uint64, pChainTimestamp uint64, allowPChainFallback_ bool, sp1Verifier_ common.Address, attestationProgramVKey_ [32]byte) (common.Address, *types.Transaction, *ZKValidatorSetRegistry, error) {
	parsed, err := ZKValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKValidatorSetRegistryBin), backend, avalancheNetworkID_, pChainID_, pChainGenesisRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, allowPChainFallback_, sp1Verifier_, attestationProgramVKey_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKValidatorSetRegistry{ZKValidatorSetRegistryCaller: ZKValidatorSetRegistryCaller{contract: contract}, ZKValidatorSetRegistryTransactor: ZKValidatorSetRegistryTransactor{contract: contract}, ZKValidatorSetRegistryFilterer: ZKValidatorSetRegistryFilterer{contract: contract}}, nil
}

// ZKValidatorSetRegistry is an auto generated Go binding around an Ethereum contract.
type ZKValidatorSetRegistry struct {
	ZKValidatorSetRegistryCaller     // Read-only binding to the contract
	ZKValidatorSetRegistryTransactor // Write-only binding to the contract
	ZKValidatorSetRegistryFilterer   // Log filterer for contract events
}

// ZKValidatorSetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKValidatorSetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKValidatorSetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKValidatorSetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKValidatorSetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKValidatorSetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKValidatorSetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKValidatorSetRegistrySession struct {
	Contract     *ZKValidatorSetRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZKValidatorSetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKValidatorSetRegistryCallerSession struct {
	Contract *ZKValidatorSetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZKValidatorSetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKValidatorSetRegistryTransactorSession struct {
	Contract     *ZKValidatorSetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZKValidatorSetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKValidatorSetRegistryRaw struct {
	Contract *ZKValidatorSetRegistry // Generic contract binding to access the raw methods on
}

// ZKValidatorSetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKValidatorSetRegistryCallerRaw struct {
	Contract *ZKValidatorSetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZKValidatorSetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKValidatorSetRegistryTransactorRaw struct {
	Contract *ZKValidatorSetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKValidatorSetRegistry creates a new instance of ZKValidatorSetRegistry, bound to a specific deployed contract.
func NewZKValidatorSetRegistry(address common.Address, backend bind.ContractBackend) (*ZKValidatorSetRegistry, error) {
	contract, err := bindZKValidatorSetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKValidatorSetRegistry{ZKValidatorSetRegistryCaller: ZKValidatorSetRegistryCaller{contract: contract}, ZKValidatorSetRegistryTransactor: ZKValidatorSetRegistryTransactor{contract: contract}, ZKValidatorSetRegistryFilterer: ZKValidatorSetRegistryFilterer{contract: contract}}, nil
}

// NewZKValidatorSetRegistryCaller creates a new read-only instance of ZKValidatorSetRegistry, bound to a specific deployed contract.
func NewZKValidatorSetRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZKValidatorSetRegistryCaller, error) {
	contract, err := bindZKValidatorSetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKValidatorSetRegistryCaller{contract: contract}, nil
}

// NewZKValidatorSetRegistryTransactor creates a new write-only instance of ZKValidatorSetRegistry, bound to a specific deployed contract.
func NewZKValidatorSetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKValidatorSetRegistryTransactor, error) {
	contract, err := bindZKValidatorSetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKValidatorSetRegistryTransactor{contract: contract}, nil
}

// NewZKValidatorSetRegistryFilterer creates a new log filterer instance of ZKValidatorSetRegistry, bound to a specific deployed contract.
func NewZKValidatorSetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKValidatorSetRegistryFilterer, error) {
	contract, err := bindZKValidatorSetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKValidatorSetRegistryFilterer{contract: contract}, nil
}

// bindZKValidatorSetRegistry binds a generic wrapper to an already deployed contract.
func bindZKValidatorSetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKValidatorSetRegistry.Contract.ZKValidatorSetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.ZKValidatorSetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.ZKValidatorSetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKValidatorSetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllowPChainFallback is a free data retrieval call binding the contract method 0x53b4c649.
//
// Solidity: function allowPChainFallback() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) AllowPChainFallback(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "allowPChainFallback")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowPChainFallback is a free data retrieval call binding the contract method 0x53b4c649.
//
// Solidity: function allowPChainFallback() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) AllowPChainFallback() (bool, error) {
	return _ZKValidatorSetRegistry.Contract.AllowPChainFallback(&_ZKValidatorSetRegistry.CallOpts)
}

// AllowPChainFallback is a free data retrieval call binding the contract method 0x53b4c649.
//
// Solidity: function allowPChainFallback() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) AllowPChainFallback() (bool, error) {
	return _ZKValidatorSetRegistry.Contract.AllowPChainFallback(&_ZKValidatorSetRegistry.CallOpts)
}

// AttestationProgramVKey is a free data retrieval call binding the contract method 0x3ee25518.
//
// Solidity: function attestationProgramVKey() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) AttestationProgramVKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "attestationProgramVKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AttestationProgramVKey is a free data retrieval call binding the contract method 0x3ee25518.
//
// Solidity: function attestationProgramVKey() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) AttestationProgramVKey() ([32]byte, error) {
	return _ZKValidatorSetRegistry.Contract.AttestationProgramVKey(&_ZKValidatorSetRegistry.CallOpts)
}

// AttestationProgramVKey is a free data retrieval call binding the contract method 0x3ee25518.
//
// Solidity: function attestationProgramVKey() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) AttestationProgramVKey() ([32]byte, error) {
	return _ZKValidatorSetRegistry.Contract.AttestationProgramVKey(&_ZKValidatorSetRegistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) AvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "avalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) AvalancheNetworkID() (uint32, error) {
	return _ZKValidatorSetRegistry.Contract.AvalancheNetworkID(&_ZKValidatorSetRegistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) AvalancheNetworkID() (uint32, error) {
	return _ZKValidatorSetRegistry.Contract.AvalancheNetworkID(&_ZKValidatorSetRegistry.CallOpts)
}

// GetValidatorSetCommitment is a free data retrieval call binding the contract method 0x8d8d23a7.
//
// Solidity: function getValidatorSetCommitment(bytes32 avalancheBlockchainID) view returns((bytes32,bytes32,uint64,uint64,uint64))
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) GetValidatorSetCommitment(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (ValidatorSetMerkleCommitment, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "getValidatorSetCommitment", avalancheBlockchainID)

	if err != nil {
		return *new(ValidatorSetMerkleCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSetMerkleCommitment)).(*ValidatorSetMerkleCommitment)

	return out0, err

}

// GetValidatorSetCommitment is a free data retrieval call binding the contract method 0x8d8d23a7.
//
// Solidity: function getValidatorSetCommitment(bytes32 avalancheBlockchainID) view returns((bytes32,bytes32,uint64,uint64,uint64))
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) GetValidatorSetCommitment(avalancheBlockchainID [32]byte) (ValidatorSetMerkleCommitment, error) {
	return _ZKValidatorSetRegistry.Contract.GetValidatorSetCommitment(&_ZKValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// GetValidatorSetCommitment is a free data retrieval call binding the contract method 0x8d8d23a7.
//
// Solidity: function getValidatorSetCommitment(bytes32 avalancheBlockchainID) view returns((bytes32,bytes32,uint64,uint64,uint64))
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) GetValidatorSetCommitment(avalancheBlockchainID [32]byte) (ValidatorSetMerkleCommitment, error) {
	return _ZKValidatorSetRegistry.Contract.GetValidatorSetCommitment(&_ZKValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) IsRegistered(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "isRegistered", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _ZKValidatorSetRegistry.Contract.IsRegistered(&_ZKValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _ZKValidatorSetRegistry.Contract.IsRegistered(&_ZKValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) PChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "pChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) PChainID() ([32]byte, error) {
	return _ZKValidatorSetRegistry.Contract.PChainID(&_ZKValidatorSetRegistry.CallOpts)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) PChainID() ([32]byte, error) {
	return _ZKValidatorSetRegistry.Contract.PChainID(&_ZKValidatorSetRegistry.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) PChainInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "pChainInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) PChainInitialized() (bool, error) {
	return _ZKValidatorSetRegistry.Contract.PChainInitialized(&_ZKValidatorSetRegistry.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) PChainInitialized() (bool, error) {
	return _ZKValidatorSetRegistry.Contract.PChainInitialized(&_ZKValidatorSetRegistry.CallOpts)
}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) Sp1Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "sp1Verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) Sp1Verifier() (common.Address, error) {
	return _ZKValidatorSetRegistry.Contract.Sp1Verifier(&_ZKValidatorSetRegistry.CallOpts)
}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) Sp1Verifier() (common.Address, error) {
	return _ZKValidatorSetRegistry.Contract.Sp1Verifier(&_ZKValidatorSetRegistry.CallOpts)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) VerifyICMMessage(opts *bind.CallOpts, message ICMMessage, avalancheBlockchainID [32]byte) error {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "verifyICMMessage", message, avalancheBlockchainID)

	if err != nil {
		return err
	}

	return err

}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _ZKValidatorSetRegistry.Contract.VerifyICMMessage(&_ZKValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _ZKValidatorSetRegistry.Contract.VerifyICMMessage(&_ZKValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCaller) VerifyMessage(opts *bind.CallOpts, message TeleporterICMMessage) (bool, error) {
	var out []interface{}
	err := _ZKValidatorSetRegistry.contract.Call(opts, &out, "verifyMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _ZKValidatorSetRegistry.Contract.VerifyMessage(&_ZKValidatorSetRegistry.CallOpts, message)
}

// VerifyMessage is a free data retrieval call binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) view returns(bool)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryCallerSession) VerifyMessage(message TeleporterICMMessage) (bool, error) {
	return _ZKValidatorSetRegistry.Contract.VerifyMessage(&_ZKValidatorSetRegistry.CallOpts, message)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x7abab0f3.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes32 signingChainID) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactor) RegisterValidatorSet(opts *bind.TransactOpts, message ICMMessage, signingChainID [32]byte) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.contract.Transact(opts, "registerValidatorSet", message, signingChainID)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x7abab0f3.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes32 signingChainID) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) RegisterValidatorSet(message ICMMessage, signingChainID [32]byte) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.RegisterValidatorSet(&_ZKValidatorSetRegistry.TransactOpts, message, signingChainID)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x7abab0f3.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes32 signingChainID) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactorSession) RegisterValidatorSet(message ICMMessage, signingChainID [32]byte) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.RegisterValidatorSet(&_ZKValidatorSetRegistry.TransactOpts, message, signingChainID)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistrySession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.SendMessage(&_ZKValidatorSetRegistry.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKValidatorSetRegistry.Contract.SendMessage(&_ZKValidatorSetRegistry.TransactOpts, message)
}

// ZKValidatorSetRegistryValidatorSetRegisteredIterator is returned from FilterValidatorSetRegistered and is used to iterate over the raw logs and unpacked data for ValidatorSetRegistered events raised by the ZKValidatorSetRegistry contract.
type ZKValidatorSetRegistryValidatorSetRegisteredIterator struct {
	Event *ZKValidatorSetRegistryValidatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *ZKValidatorSetRegistryValidatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKValidatorSetRegistryValidatorSetRegistered)
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
		it.Event = new(ZKValidatorSetRegistryValidatorSetRegistered)
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
func (it *ZKValidatorSetRegistryValidatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKValidatorSetRegistryValidatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKValidatorSetRegistryValidatorSetRegistered represents a ValidatorSetRegistered event raised by the ZKValidatorSetRegistry contract.
type ZKValidatorSetRegistryValidatorSetRegistered struct {
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetRegistered is a free log retrieval operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryFilterer) FilterValidatorSetRegistered(opts *bind.FilterOpts, avalancheBlockchainID [][32]byte) (*ZKValidatorSetRegistryValidatorSetRegisteredIterator, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _ZKValidatorSetRegistry.contract.FilterLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ZKValidatorSetRegistryValidatorSetRegisteredIterator{contract: _ZKValidatorSetRegistry.contract, event: "ValidatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorSetRegistered is a free log subscription operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryFilterer) WatchValidatorSetRegistered(opts *bind.WatchOpts, sink chan<- *ZKValidatorSetRegistryValidatorSetRegistered, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _ZKValidatorSetRegistry.contract.WatchLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKValidatorSetRegistryValidatorSetRegistered)
				if err := _ZKValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
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

// ParseValidatorSetRegistered is a log parse operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_ZKValidatorSetRegistry *ZKValidatorSetRegistryFilterer) ParseValidatorSetRegistered(log types.Log) (*ZKValidatorSetRegistryValidatorSetRegistered, error) {
	event := new(ZKValidatorSetRegistryValidatorSetRegistered)
	if err := _ZKValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
