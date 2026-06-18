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
	Bin: "0x610120604052348015610010575f5ffd5b50604051611bd9380380611bd983398101604081905261002f91610119565b63ffffffff909816608090815260a088815292151560c0526001600160a01b0390911660e0526101009790975260408051918201815286825260208083019687526001600160401b03958616838301908152948616606084019081529386169883019889525f97885287905290952094518555925160018501555160029093018054925194518216600160801b02600160801b600160c01b031995831668010000000000000000026001600160801b0319909416949092169390931791909117929092169190911790556101cc565b80516001600160401b0381168114610114575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f6101208a8c031215610132575f5ffd5b895163ffffffff81168114610145575f5ffd5b60208b015160408c0151919a509850965061016260608b016100fe565b955061017060808b016100fe565b945061017e60a08b016100fe565b935060c08a01518015158114610192575f5ffd5b60e08b01519093506001600160a01b03811681146101ae575f5ffd5b809250505f6101008b01519050809150509295985092959850929598565b60805160a05160c05160e051610100516119996102405f395f818160fc0152610bdf01525f81816101310152610bb201525f8181610170015261064201525f8181610197015281816104a3015261057c01525f81816101db015281816103d30152818161050c015261083e01526119995ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c8063580d632b1161006e578063580d632b146101ce57806368531ed0146101d65780637abab0f3146102125780638d8d23a714610225578063eb97cd2c14610317578063f1faff001461032a575f5ffd5b806327258b22146100b55780633ee25518146100f757806352a07fa31461012c57806353b4c6491461016b578063541dcba41461019257806357262e7f146101b9575b5f5ffd5b6100e26100c3366004610ff8565b5f908152602081905260409020600201546001600160401b0316151590565b60405190151581526020015b60405180910390f35b61011e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100ee565b6101537f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ee565b6100e27f000000000000000000000000000000000000000000000000000000000000000081565b61011e7f000000000000000000000000000000000000000000000000000000000000000081565b6101cc6101c7366004611025565b61033d565b005b6100e26104a1565b6101fd7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100ee565b6101cc610220366004611025565b6104e1565b6102bb610233366004610ff8565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152505f9081526020818152604091829020825160a08101845281548152600182015492810192909252600201546001600160401b0380821693830193909352600160401b810483166060830152600160801b9004909116608082015290565b6040516100ee91905f60a08201905082518252602083015160208301526001600160401b0360408401511660408301526001600160401b0360608401511660608301526001600160401b03608084015116608083015292915050565b6101cc610325366004611066565b61075b565b6100e26103383660046110a4565b6107d9565b6103456104a1565b6103965760405162461bcd60e51b815260206004820152601760248201527f502d636861696e206e6f7420696e697469616c697a656400000000000000000060448201526064015b60405180910390fd5b5f818152602081905260409020600201546001600160401b03166103cc5760405162461bcd60e51b815260040161038d906110d5565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610403604084016020850161111c565b63ffffffff16146104265760405162461bcd60e51b815260040161038d9061113f565b5f61048361043a604085016020860161111c565b60408501355f61044a878061116c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506108ec92505050565b905061049c610495606085018561116c565b8385610998565b505050565b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b0316151590565b6104e96104a1565b6105055760405162461bcd60e51b815260040161038d906111b5565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661053c604084016020850161111c565b63ffffffff161461055f5760405162461bcd60e51b815260040161038d9061113f565b5f61057261056d848061116c565b610c41565b80519091508281147f000000000000000000000000000000000000000000000000000000000000000084146105c0835f908152602081905260409020600201546001600160401b0316151590565b610631578061062c5760405162461bcd60e51b815260206004820152603260248201527f496e697469616c20726567697374726174696f6e206d757374206265207369676044820152713732b210313c903a3432902816a1b430b4b760711b606482015260840161038d565b6106a6565b8180610662575080801561066257507f00000000000000000000000000000000000000000000000000000000000000005b6106a65760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b21039b4b3b734b7339031b430b4b760591b604482015260640161038d565b6106b0868661033d565b5f83815260208181526040808320875181559187015160018301558087015160029092018054606089015160808a01516001600160401b03908116600160801b0267ffffffffffffffff60801b19928216600160401b026fffffffffffffffffffffffffffffffff1990941691909616179190911716929092179091555184917f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce9391a2505050505050565b6005600160991b0163ee5b48eb6107796107748461142d565b610e2b565b6040518263ffffffff1660e01b81526004016107959190611545565b6020604051808303815f875af11580156107b1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107d59190611557565b5050565b5f6107e26104a1565b6107fe5760405162461bcd60e51b815260040161038d906111b5565b6040808301355f908152602081905220600201546001600160401b03166108375760405162461bcd60e51b815260040161038d906110d5565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661086e604084016020850161111c565b63ffffffff16146108915760405162461bcd60e51b815260040161038d9061113f565b5f6108c66108a5604085016020860161111c565b6040850135306108c16108b8888061156e565b6107749061142d565b6108ec565b90506108e36108d8606085018561116c565b838660400135610998565b50600192915050565b60606001600160a01b038316610944575f60f01b85858451600e61091091906115a1565b855160405161092e95949392915f9160019183918b906020016115cb565b6040516020818303038152906040529050610990565b5f60f01b85858451602261095891906115a1565b855160405161097e95949392915f9160019160149160608d901b91908c90602001611642565b60405160208183030381529060405290505b949350505050565b5f806109a6858701876116cf565b915091505f828060200190518101906109bf9190611732565b5f8581526020819052604090208151919250908514610a205760405162461bcd60e51b815260206004820152601e60248201527f7075626c69632076616c756520636861696e204944206d69736d617463680000604482015260640161038d565b8060010154826020015114610a775760405162461bcd60e51b815260206004820152601a60248201527f7075626c69632076616c756520726f6f74206d69736d61746368000000000000604482015260640161038d565b600286604051610a8791906117a6565b602060405180830381855afa158015610aa2573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610ac59190611557565b826040015114610b225760405162461bcd60e51b815260206004820152602260248201527f7075626c69632076616c7565206d6573736167652068617368206d69736d61746044820152610c6d60f31b606482015260840161038d565b60608201516002820154610b3f91906001600160401b0316610ea4565b610b9b5760405162461bcd60e51b815260206004820152602760248201527f7374616b652d77656967687465642071756f72756d207468726573686f6c64206044820152661b9bdd081b595d60ca1b606482015260840161038d565b60405163020a49e360e51b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906341493c6090610c0b907f000000000000000000000000000000000000000000000000000000000000000090889088906004016117b1565b5f6040518083038186803b158015610c21575f5ffd5b505afa158015610c33573d5f5f3e3d5ffd5b505050505050505050505050565b6040805160a0810182525f8082526020820181905291810182905260608101829052608081019190915282825f818110610c7d57610c7d6117e5565b909101356001600160f81b0319161590508015610cba575082826001818110610ca857610ca86117e5565b909101356001600160f81b0319161590505b610cf95760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818dbd91958c8125160821b604482015260640161038d565b5f610d086006600285876117f9565b610d1191611820565b60e01c905060068114610d835760405162461bcd60e51b815260206004820152603460248201527f496e76616c69642056616c696461746f725365744d65726b6c65436f6d6d69746044820152731b595b9d081c185e5b1bd859081d1e5c1948125160621b606482015260840161038d565b6040805160a0810190915280610d9d6026600687896117f9565b610da691611858565b8152602001610db96046602687896117f9565b610dc291611858565b8152602001610dd5604e604687896117f9565b610dde91611875565b60c01c8152602001610df46056604e87896117f9565b610dfd91611875565b60c01c8152602001610e13605e605687896117f9565b610e1c91611875565b60c01c90529150505b92915050565b6060815f015182602001518360400151846060015185608001518660a001518760c0015151610e5d8960c00151610edd565b60e08a0151805190610e6e90610f69565b8b6101000151604051602001610e8e9b9a999897969594939291906118ab565b6040516020818303038152906040529050919050565b5f80610eba6001600160401b038416604361194c565b90505f610ed16001600160401b038616606461194c565b90911115949350505050565b60605f82516014610eee919061194c565b6001600160401b03811115610f0557610f056111f9565b6040519080825280601f01601f191660200182016040528015610f2f576020820181803683370190505b50905060208101602c840184515f5b81811015610f5e5782518452601490930192602090920191600101610f3e565b509295945050505050565b60605f82516034610f7a919061194c565b6001600160401b03811115610f9157610f916111f9565b6040519080825280601f01601f191660200182016040528015610fbb576020820181803683370190505b50835190915060208083019085015f5b83811015610f5e5781518051845260209081015160601b8185015260349093019290910190600101610fcb565b5f60208284031215611008575f5ffd5b5035919050565b5f6080828403121561101f575f5ffd5b50919050565b5f5f60408385031215611036575f5ffd5b82356001600160401b0381111561104b575f5ffd5b6110578582860161100f565b95602094909401359450505050565b5f60208284031215611076575f5ffd5b81356001600160401b0381111561108b575f5ffd5b8201610120818503121561109d575f5ffd5b9392505050565b5f602082840312156110b4575f5ffd5b81356001600160401b038111156110c9575f5ffd5b6109908482850161100f565b60208082526027908201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676040820152661a5d995b88125160ca1b606082015260800190565b5f6020828403121561112c575f5ffd5b813563ffffffff8116811461109d575f5ffd5b60208082526013908201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b604082015260600190565b5f5f8335601e19843603018112611181575f5ffd5b8301803591506001600160401b0382111561119a575f5ffd5b6020019150368190038213156111ae575f5ffd5b9250929050565b60208082526024908201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656040820152633932b21760e11b606082015260800190565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561122f5761122f6111f9565b60405290565b60405161012081016001600160401b038111828210171561122f5761122f6111f9565b604051601f8201601f191681016001600160401b0381118282101715611280576112806111f9565b604052919050565b80356001600160a01b038116811461129e575f5ffd5b919050565b5f6001600160401b038211156112bb576112bb6111f9565b5060051b60200190565b5f82601f8301126112d4575f5ffd5b81356112e76112e2826112a3565b611258565b8082825260208201915060208360051b860101925085831115611308575f5ffd5b602085015b8381101561132c5761131e81611288565b83526020928301920161130d565b5095945050505050565b5f82601f830112611345575f5ffd5b81356113536112e2826112a3565b8082825260208201915060208360061b860101925085831115611374575f5ffd5b602085015b8381101561132c5760408188031215611390575f5ffd5b61139861120d565b813581526113a860208301611288565b602082015280845250602083019250604081019050611379565b5f82601f8301126113d1575f5ffd5b81356001600160401b038111156113ea576113ea6111f9565b6113fd601f8201601f1916602001611258565b818152846020838601011115611411575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f610120823603121561143e575f5ffd5b611446611235565b8235815261145660208401611288565b602082015261146760408401611288565b60408201526060838101359082015261148260808401611288565b608082015260a0838101359082015260c08301356001600160401b038111156114a9575f5ffd5b6114b5368286016112c5565b60c08301525060e08301356001600160401b038111156114d3575f5ffd5b6114df36828601611336565b60e0830152506101008301356001600160401b038111156114fe575f5ffd5b61150a368286016113c2565b6101008301525092915050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61109d6020830184611517565b5f60208284031215611567575f5ffd5b5051919050565b5f823561011e19833603018112611583575f5ffd5b9190910192915050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610e2557610e2561158d565b5f81518060208401855e5f93019283525090919050565b6001600160f01b03198a811682526001600160e01b031960e08b811b82166002850152600684018b905289811b82166026850152918816602a84015286821b8116602c84015285821b811660308401529084901b1660348201525f61163360388301846115b4565b9b9a5050505050505050505050565b6001600160f01b03198b8116825260e08b811b6001600160e01b03199081166002850152600684018c90528a821b81166026850152918916602a84015287811b8216602c84015286811b821660308401526bffffffffffffffffffffffff198616603484015284901b1660488201525f6116bf604c8301846115b4565b9c9b505050505050505050505050565b5f5f604083850312156116e0575f5ffd5b82356001600160401b038111156116f5575f5ffd5b611701858286016113c2565b92505060208301356001600160401b0381111561171c575f5ffd5b611728858286016113c2565b9150509250929050565b5f6080828403128015611743575f5ffd5b50604051608081016001600160401b0381118282101715611766576117666111f9565b60409081528351825260208085015190830152838101519082015260608301516001600160401b038116811461179a575f5ffd5b60608201529392505050565b5f61109d82846115b4565b838152606060208201525f6117c96060830185611517565b82810360408401526117db8185611517565b9695505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f5f85851115611807575f5ffd5b83861115611813575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015611851576001600160e01b0319600485900360031b81901b82161691505b5092915050565b80356020831015610e25575f19602084900360031b1b1692915050565b80356001600160c01b03198116906008841015611851576001600160c01b031960089490940360031b84901b1690921692915050565b8b81526001600160601b03198b60601b1660208201526001600160601b03198a60601b1660348201528860488201526001600160601b03198860601b16606882015286607c82015263ffffffff60e01b8660e01b16609c8201525f61191360a08301876115b4565b60e086901b6001600160e01b031916815261193a61193460048301876115b4565b856115b4565b9e9d5050505050505050505050505050565b8082028115828204841417610e2557610e2561158d56fea26469706673582212204c087c7caae06b18815277d4bfec955e2e6db364b464e59e6abab186a99a9e0564736f6c634300081e0033",
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
