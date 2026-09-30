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
	Bin: "0x610120604052348015610010575f5ffd5b50604051611d0e380380611d0e83398101604081905261002f91610119565b63ffffffff909816608090815260a088815292151560c0526001600160a01b0390911660e0526101009790975260408051918201815286825260208083019687526001600160401b03958616838301908152948616606084019081529386169883019889525f97885287905290952094518555925160018501555160029093018054925194518216600160801b02600160801b600160c01b031995831668010000000000000000026001600160801b0319909416949092169390931791909117929092169190911790556101cc565b80516001600160401b0381168114610114575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610132575f5ffd5b895163ffffffff81168114610145575f5ffd5b60208b015160408c0151919a509850965061016260608b016100fe565b955061017060808b016100fe565b945061017e60a08b016100fe565b935060c08a01518015158114610192575f5ffd5b60e08b01519093506001600160a01b03811681146101ae575f5ffd5b809250505f6101008b01519050809150509295985092959850929598565b60805160a05160c05160e05161010051611ace6102405f395f818160fc0152610ccf01525f81816101310152610ca201525f8181610170015261064201525f8181610197015281816104a3015261057c01525f81816101db015281816103d30152818161050c015261092e0152611ace5ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c8063580d632b1161006e578063580d632b146101ce57806368531ed0146101d65780637abab0f3146102125780638d8d23a714610225578063eb97cd2c14610317578063f1faff001461032a575f5ffd5b806327258b22146100b55780633ee25518146100f757806352a07fa31461012c57806353b4c6491461016b578063541dcba41461019257806357262e7f146101b9575b5f5ffd5b6100e26100c33660046110e8565b5f908152602081905260409020600201546001600160401b0316151590565b60405190151581526020015b60405180910390f35b61011e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100ee565b6101537f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ee565b6100e27f000000000000000000000000000000000000000000000000000000000000000081565b61011e7f000000000000000000000000000000000000000000000000000000000000000081565b6101cc6101c7366004611115565b61033d565b005b6100e26104a1565b6101fd7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100ee565b6101cc610220366004611115565b6104e1565b6102bb6102333660046110e8565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9081526020818152604091829020825160a08101845281548152600182015492810192909252600201546001600160401b0380821693830193909352600160401b810483166060830152600160801b9004909116608082015290565b6040516100ee91905f60a08201905082518252602083015160208301526001600160401b0360408401511660408301526001600160401b0360608401511660608301526001600160401b03608084015116608083015292915050565b6101cc610325366004611156565b61075b565b6100e2610338366004611194565b6108c9565b6103456104a1565b6103965760405162461bcd60e51b815260206004820152601760248201527f502d636861696e206e6f7420696e697469616c697a656400000000000000000060448201526064015b60405180910390fd5b5f818152602081905260409020600201546001600160401b03166103cc5760405162461bcd60e51b815260040161038d906111c5565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610403604084016020850161120c565b63ffffffff16146104265760405162461bcd60e51b815260040161038d9061122f565b5f61048361043a604085016020860161120c565b60408501355f61044a878061125c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506109dc92505050565b905061049c610495606085018561125c565b8385610a88565b505050565b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b0316151590565b6104e96104a1565b6105055760405162461bcd60e51b815260040161038d906112a5565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661053c604084016020850161120c565b63ffffffff161461055f5760405162461bcd60e51b815260040161038d9061122f565b5f61057261056d848061125c565b610d31565b80519091508281147f000000000000000000000000000000000000000000000000000000000000000084146105c0835f908152602081905260409020600201546001600160401b0316151590565b610631578061062c5760405162461bcd60e51b815260206004820152603260248201527f496e697469616c20726567697374726174696f6e206d757374206265207369676044820152713732b210313c903a3432902816a1b430b4b760711b606482015260840161038d565b6106a6565b8180610662575080801561066257507f00000000000000000000000000000000000000000000000000000000000000005b6106a65760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b21039b4b3b734b7339031b430b4b760591b604482015260640161038d565b6106b0868661033d565b5f83815260208181526040808320875181559187015160018301558087015160029092018054606089015160808a01516001600160401b03908116600160801b0267ffffffffffffffff60801b19928216600160401b026fffffffffffffffffffffffffffffffff1990941691909616179190911716929092179091555184917f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce9391a2505050505050565b61076b6060820160408301611310565b6001600160a01b0316336001600160a01b0316148061080957506107956060820160408301611310565b6001600160a01b031663d67bdd256040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107f4919061132b565b6001600160a01b0316336001600160a01b0316145b61084b5760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b604482015260640161038d565b6005600160991b0163ee5b48eb61086961086484611562565b610f1b565b6040518263ffffffff1660e01b8152600401610885919061167a565b6020604051808303815f875af11580156108a1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c5919061168c565b5050565b5f6108d26104a1565b6108ee5760405162461bcd60e51b815260040161038d906112a5565b6040808301355f908152602081905220600201546001600160401b03166109275760405162461bcd60e51b815260040161038d906111c5565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661095e604084016020850161120c565b63ffffffff16146109815760405162461bcd60e51b815260040161038d9061122f565b5f6109b6610995604085016020860161120c565b6040850135306109b16109a888806116a3565b61086490611562565b6109dc565b90506109d36109c8606085018561125c565b838660400135610a88565b50600192915050565b60606001600160a01b038316610a34575f60f01b85858451600e610a0091906116d6565b8551604051610a1e95949392915f9160019183918b90602001611700565b6040516020818303038152906040529050610a80565b5f60f01b858584516022610a4891906116d6565b8551604051610a6e95949392915f9160019160149160608d901b91908c90602001611777565b60405160208183030381529060405290505b949350505050565b5f80610a9685870187611804565b915091505f82806020019051810190610aaf9190611867565b5f8581526020819052604090208151919250908514610b105760405162461bcd60e51b815260206004820152601e60248201527f7075626c69632076616c756520636861696e204944206d69736d617463680000604482015260640161038d565b8060010154826020015114610b675760405162461bcd60e51b815260206004820152601a60248201527f7075626c69632076616c756520726f6f74206d69736d61746368000000000000604482015260640161038d565b600286604051610b7791906118db565b602060405180830381855afa158015610b92573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610bb5919061168c565b826040015114610c125760405162461bcd60e51b815260206004820152602260248201527f7075626c69632076616c7565206d6573736167652068617368206d69736d61746044820152610c6d60f31b606482015260840161038d565b60608201516002820154610c2f91906001600160401b0316610f94565b610c8b5760405162461bcd60e51b815260206004820152602760248201527f7374616b652d77656967687465642071756f72756d207468726573686f6c64206044820152661b9bdd081b595d60ca1b606482015260840161038d565b60405163020a49e360e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906341493c6090610cfb907f000000000000000000000000000000000000000000000000000000000000000090889088906004016118e6565b5f6040518083038186803b158015610d11575f5ffd5b505afa158015610d23573d5f5f3e3d5ffd5b505050505050505050505050565b6040805160a0810182525f8082526020820181905291810182905260608101829052608081019190915282825f818110610d6d57610d6d61191a565b909101356001600160f81b0319161590508015610daa575082826001818110610d9857610d9861191a565b909101356001600160f81b0319161590505b610de95760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818dbd91958c8125160821b604482015260640161038d565b5f610df860066002858761192e565b610e0191611955565b60e01c905060068114610e735760405162461bcd60e51b815260206004820152603460248201527f496e76616c69642056616c696461746f725365744d65726b6c65436f6d6d69746044820152731b595b9d081c185e5b1bd859081d1e5c1948125160621b606482015260840161038d565b6040805160a0810190915280610e8d60266006878961192e565b610e969161198d565b8152602001610ea960466026878961192e565b610eb29161198d565b8152602001610ec5604e6046878961192e565b610ece916119aa565b60c01c8152602001610ee46056604e878961192e565b610eed916119aa565b60c01c8152602001610f03605e6056878961192e565b610f0c916119aa565b60c01c90529150505b92915050565b6060815f015182602001518360400151846060015185608001518660a001518760c0015151610f4d8960c00151610fcd565b60e08a0151805190610f5e90611059565b8b6101000151604051602001610f7e9b9a999897969594939291906119e0565b6040516020818303038152906040529050919050565b5f80610faa6001600160401b0384166043611a81565b90505f610fc16001600160401b0386166064611a81565b90911115949350505050565b60605f82516014610fde9190611a81565b6001600160401b03811115610ff557610ff5611346565b6040519080825280601f01601f19166020018201604052801561101f576020820181803683370190505b50905060208101602c840184515f5b8181101561104e578251845260149093019260209092019160010161102e565b509295945050505050565b60605f8251603461106a9190611a81565b6001600160401b0381111561108157611081611346565b6040519080825280601f01601f1916602001820160405280156110ab576020820181803683370190505b50835190915060208083019085015f5b8381101561104e5781518051845260209081015160601b81850152603490930192909101906001016110bb565b5f602082840312156110f8575f5ffd5b5035919050565b5f6080828403121561110f575f5ffd5b50919050565b5f5f60408385031215611126575f5ffd5b82356001600160401b0381111561113b575f5ffd5b611147858286016110ff565b95602094909401359450505050565b5f60208284031215611166575f5ffd5b81356001600160401b0381111561117b575f5ffd5b8201610120818503121561118d575f5ffd5b9392505050565b5f602082840312156111a4575f5ffd5b81356001600160401b038111156111b9575f5ffd5b610a80848285016110ff565b60208082526027908201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676040820152661a5d995b88125160ca1b606082015260800190565b5f6020828403121561121c575f5ffd5b813563ffffffff8116811461118d575f5ffd5b60208082526013908201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b604082015260600190565b5f5f8335601e19843603018112611271575f5ffd5b8301803591506001600160401b0382111561128a575f5ffd5b60200191503681900382131561129e575f5ffd5b9250929050565b60208082526024908201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656040820152633932b21760e11b606082015260800190565b6001600160a01b03811681146112fd575f5ffd5b50565b803561130b816112e9565b919050565b5f60208284031215611320575f5ffd5b813561118d816112e9565b5f6020828403121561133b575f5ffd5b815161118d816112e9565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561137c5761137c611346565b60405290565b60405161012081016001600160401b038111828210171561137c5761137c611346565b604051601f8201601f191681016001600160401b03811182821017156113cd576113cd611346565b604052919050565b5f6001600160401b038211156113ed576113ed611346565b5060051b60200190565b5f82601f830112611406575f5ffd5b8135611419611414826113d5565b6113a5565b8082825260208201915060208360051b86010192508583111561143a575f5ffd5b602085015b83811015611460578035611452816112e9565b83526020928301920161143f565b5095945050505050565b5f82601f830112611479575f5ffd5b8135611487611414826113d5565b8082825260208201915060208360061b8601019250858311156114a8575f5ffd5b602085015b8381101561146057604081880312156114c4575f5ffd5b6114cc61135a565b8135815260208201356114de816112e9565b60208281019190915290845292909201916040016114ad565b5f82601f830112611506575f5ffd5b81356001600160401b0381111561151f5761151f611346565b611532601f8201601f19166020016113a5565b818152846020838601011115611546575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6101208236031215611573575f5ffd5b61157b611382565b8235815261158b60208401611300565b602082015261159c60408401611300565b6040820152606083810135908201526115b760808401611300565b608082015260a0838101359082015260c08301356001600160401b038111156115de575f5ffd5b6115ea368286016113f7565b60c08301525060e08301356001600160401b03811115611608575f5ffd5b6116143682860161146a565b60e0830152506101008301356001600160401b03811115611633575f5ffd5b61163f368286016114f7565b6101008301525092915050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61118d602083018461164c565b5f6020828403121561169c575f5ffd5b5051919050565b5f823561011e198336030181126116b8575f5ffd5b9190910192915050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610f1557610f156116c2565b5f81518060208401855e5f93019283525090919050565b6001600160f01b03198a811682526001600160e01b031960e08b811b82166002850152600684018b905289811b82166026850152918816602a84015286821b8116602c84015285821b811660308401529084901b1660348201525f61176860388301846116e9565b9b9a5050505050505050505050565b6001600160f01b03198b8116825260e08b811b6001600160e01b03199081166002850152600684018c90528a821b81166026850152918916602a84015287811b8216602c84015286811b821660308401526bffffffffffffffffffffffff198616603484015284901b1660488201525f6117f4604c8301846116e9565b9c9b505050505050505050505050565b5f5f60408385031215611815575f5ffd5b82356001600160401b0381111561182a575f5ffd5b611836858286016114f7565b92505060208301356001600160401b03811115611851575f5ffd5b61185d858286016114f7565b9150509250929050565b5f6080828403128015611878575f5ffd5b50604051608081016001600160401b038111828210171561189b5761189b611346565b60409081528351825260208085015190830152838101519082015260608301516001600160401b03811681146118cf575f5ffd5b60608201529392505050565b5f61118d82846116e9565b838152606060208201525f6118fe606083018561164c565b8281036040840152611910818561164c565b9695505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f5f8585111561193c575f5ffd5b83861115611948575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015611986576001600160e01b0319600485900360031b81901b82161691505b5092915050565b80356020831015610f15575f19602084900360031b1b1692915050565b80356001600160c01b03198116906008841015611986576001600160c01b031960089490940360031b84901b1690921692915050565b8b81526001600160601b03198b60601b1660208201526001600160601b03198a60601b1660348201528860488201526001600160601b03198860601b16606882015286607c82015263ffffffff60e01b8660e01b16609c8201525f611a4860a08301876116e9565b60e086901b6001600160e01b0319168152611a6f611a6960048301876116e9565b856116e9565b9e9d5050505050505050505050505050565b8082028115828204841417610f1557610f156116c256fea26469706673582212203e7356e593bd059454a7dae55b6d12d840d1ba356d1b03c43974e54a2f0d65df64736f6c634300081e0033",
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
