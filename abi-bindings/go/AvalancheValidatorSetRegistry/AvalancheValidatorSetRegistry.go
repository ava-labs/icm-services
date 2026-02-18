// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avalanchevalidatorsetregistry

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

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	BlsPublicKey []byte
	Weight       uint64
}

// ValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSet struct {
	AvalancheBlockchainID [32]byte
	Validators            []Validator
	TotalWeight           uint64
	PChainHeight          uint64
	PChainTimestamp       uint64
}

// ValidatorSetMetadata is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetMetadata struct {
	AvalancheBlockchainID [32]byte
	PChainHeight          uint64
	PChainTimestamp       uint64
	ShardHashes           [][32]byte
}

// ValidatorSetShard is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetShard struct {
	ShardNumber           uint64
	AvalancheBlockchainID [32]byte
}

// AvalancheValidatorSetRegistryMetaData contains all meta data concerning the AvalancheValidatorSetRegistry contract.
var AvalancheValidatorSetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"updateValidatorSetWithDiff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b50604051611cc1380380611cc183398101604081905261002e91610208565b63ffffffff8216608052805160a08190525f90815260016020818152604092839020818501518154948601516001600160401b0390811668010000000000000000026001600160801b03199096169116179390931783556060840151805161009d938501929190910190610124565b50600401805460ff60401b19166801000000000000000090811790915560a0515f908152602081815260409182902084518155908401516002909101805494909201516001600160401b03908116600160801b02600160801b600160c01b03199190921690930292909216600160401b600160c01b03199093169290921717905550610330565b828054828255905f5260205f2090810192821561015d579160200282015b8281111561015d578251825591602001919060010190610142565b5061016992915061016d565b5090565b5b80821115610169575f815560010161016e565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156101b7576101b7610181565b60405290565b604051601f8201601f191681016001600160401b03811182821017156101e5576101e5610181565b604052919050565b80516001600160401b0381168114610203575f5ffd5b919050565b5f5f60408385031215610219575f5ffd5b825163ffffffff8116811461022c575f5ffd5b60208401519092506001600160401b03811115610247575f5ffd5b830160808186031215610258575f5ffd5b610260610195565b81518152610270602083016101ed565b6020820152610281604083016101ed565b604082015260608201516001600160401b0381111561029e575f5ffd5b80830192505085601f8301126102b2575f5ffd5b81516001600160401b038111156102cb576102cb610181565b8060051b6102db602082016101bd565b918252602081850181019290810190898411156102f6575f5ffd5b6020860195505b8386101561031c578551808352602096870196909350909101906102fd565b606085015250949791965090945050505050565b60805160a05161194661037b5f395f81816101360152818161017d015281816104330152610ace01525f81816101d50152818161020e0152818161054601526109b501526119465ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c806368531ed0116100885780638e91cb43116100635780638e91cb431461025e57806393356840146102715780639def1e78146102845780639fd530d4146102a6575f5ffd5b806368531ed0146101d057806382366d051461020c5780638457eaa714610232575f5ffd5b806327258b22146100cf57806349e4db9c14610111578063541dcba41461013157806357262e7f14610166578063580d632b1461017b5780636766233d146101bd575b5f5ffd5b6100fc6100dd366004610f4e565b5f908152602081905260409020600201546001600160401b0316151590565b60405190151581526020015b60405180910390f35b61012461011f366004610f4e565b6102b4565b6040516101089190610f93565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610108565b61017961017436600461107c565b610431565b005b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b031615156100fc565b6101796101cb366004611127565b610761565b6101f77f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610108565b7f00000000000000000000000000000000000000000000000000000000000000006101f7565b6100fc610240366004610f4e565b5f90815260016020526040902060040154600160401b900460ff1690565b61017961026c3660046111b8565b6109ae565b61017961027f366004611127565b610e56565b6102976102923660046111b8565b610e90565b604051610108939291906112c3565b61017961027f366004611374565b6040805160a0810182525f80825260606020830181905292820181905291810182905260808101919091525f82815260208181526040808320815160a081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b828210156103f0578382905f5260205f2090600202016040518060400160405290815f8201805461034e906113a5565b80601f016020809104026020016040519081016040528092919081815260200182805461037a906113a5565b80156103c55780601f1061039c576101008083540402835291602001916103c5565b820191905f5260205f20905b8154815290600101906020018083116103a857829003601f168201915b50505091835250506001918201546001600160401b031660209182015291835292909201910161031e565b50505090825250600291909101546001600160401b038082166020840152600160401b820481166040840152600160801b9091041660609091015292915050565b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b03166104c95760405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b60648201526084015b60405180910390fd5b5f818152602081905260409020600201546001600160401b031661053f5760405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b60648201526084016104c0565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661057660408401602085016113d7565b63ffffffff16146105bf5760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b60448201526064016104c0565b5f73__$aaf4ae346b84a712cc43f25bb66199d6fb$__63858ad3986105e76060860186611401565b6040518363ffffffff1660e01b815260040161060492919061144a565b5f60405180830381865af415801561061e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261064591908101906114c5565b90505f61065860408501602086016113d7565b60408501356106678680611401565b60405160200161067a9493929190611576565b60408051601f198184030181528282525f868152602081905291909120630161c9f960e61b835290925073__$aaf4ae346b84a712cc43f25bb66199d6fb$__916358727e40916106d091869186916004016115a2565b602060405180830381865af41580156106eb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061070f919061172d565b61075b5760405162461bcd60e51b815260206004820152601b60248201527f4661696c656420746f20766572696679207369676e617475726573000000000060448201526064016104c0565b50505050565b61078882602001355f9081526001602052604090206004015460ff600160401b9091041690565b6107d45760405162461bcd60e51b815260206004820152601f60248201527f526567697374726174696f6e206973206e6f7420696e2070726f67726573730060448201526064016104c0565b602082018035906107e5908461174c565b5f828152600160208190526040909120600201546001600160401b039283169261081192911690611786565b6001600160401b0316146108675760405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f72646572000000000060448201526064016104c0565b60028260405161087791906117ab565b602060405180830381855afa158015610892573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906108b591906117c1565b5f8281526001602081815260409092208101916108d49087018761174c565b6108de91906117d8565b6001600160401b0316815481106108f7576108f76117f7565b905f5260205f200154146109455760405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b60448201526064016104c0565b61094f8383610e56565b61097683602001355f9081526001602052604090206004015460ff600160401b9091041690565b6109a9576040516020840135907f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2905f90a25b505050565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166109e560408501602086016113d7565b63ffffffff1614610a2e5760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b60448201526064016104c0565b6040808401355f90815260016020522060040154600160401b900460ff1615610aa75760405162461bcd60e51b815260206004820152602560248201527f4120726567697374726174696f6e20697320616c726561647920696e2070726f604482015264677265737360d81b60648201526084016104c0565b6040808401355f908152602081905220600201546001600160401b0316610af757610af2837f0000000000000000000000000000000000000000000000000000000000000000610431565b610b05565b610b05838460400135610431565b5f5f5f610b13868686610e90565b8251929550909350915060408701358114610b705760405162461bcd60e51b815260206004820152601860248201527f536f7572636520636861696e204944206d69736d61746368000000000000000060448201526064016104c0565b825160608501515160011015610d335784515f90815260016020818152604092839020818901518154948a01516001600160401b03908116600160401b026001600160801b031990961691161793909317835560608801518051610bdb938501929190910190610ef1565b5060028101805467ffffffffffffffff1916600117905560048101805468ffffffffffffffffff19166001600160401b03861617600160401b1790555f5b82811015610c995781600301868281518110610c3757610c376117f7565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610c689082611856565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610c19565b506040808a01355f908152602081905220600201546001600160401b0316610d2d5785515f9081526020818152604091829020885181559088015160029091018054928901516001600160401b03908116600160801b0267ffffffffffffffff60801b1991909316600160401b021677ffffffffffffffffffffffffffffffff000000000000000019909316929092171790555b50610e22565b84515f9081526020818152604080832088518155600281018054938a0151928a01516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909616918a16919091179490941792909216929092179055905b82811015610e1f5781600101868281518110610dbd57610dbd6117f7565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610dee9082611856565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610d9f565b50505b60405182907f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93905f90a25050505050505050565b60405162461bcd60e51b815260206004820152600f60248201526e139bdd081a5b5c1b195b595b9d1959608a1b60448201526064016104c0565b604080516080810182525f808252602082018190529181019190915260608082015260605f60405162461bcd60e51b81526004016104c0906020808252600f908201526e139bdd081a5b5c1b195b595b9d1959608a1b604082015260600190565b828054828255905f5260205f20908101928215610f2a579160200282015b82811115610f2a578251825591602001919060010190610f0f565b50610f36929150610f3a565b5090565b5b80821115610f36575f8155600101610f3b565b5f60208284031215610f5e575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c0820183516020840152602084015160a0604085015281815180845260e08601915060e08160051b87010193506020830192505f5b8181101561101c5760df198786030183528351805160408752610ff46040880182610f65565b6020928301516001600160401b03169783019790975250938401939290920191600101610fce565b5050505060408401516001600160401b03811660608501525060608401516001600160401b03811660808501525060808401516001600160401b03811660a0850152509392505050565b5f60808284031215611076575f5ffd5b50919050565b5f5f6040838503121561108d575f5ffd5b82356001600160401b038111156110a2575f5ffd5b6110ae85828601611066565b95602094909401359450505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156110f9576110f96110bd565b604052919050565b5f6001600160401b03821115611119576111196110bd565b50601f01601f191660200190565b5f5f8284036060811215611139575f5ffd5b6040811215611146575f5ffd5b5082915060408301356001600160401b03811115611162575f5ffd5b8301601f81018513611172575f5ffd5b803561118561118082611101565b6110d1565b818152866020838501011115611199575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f604084860312156111ca575f5ffd5b83356001600160401b038111156111df575f5ffd5b6111eb86828701611066565b93505060208401356001600160401b03811115611206575f5ffd5b8401601f81018613611216575f5ffd5b80356001600160401b0381111561122b575f5ffd5b86602082840101111561123c575f5ffd5b939660209190910195509293505050565b5f82825180855260208501945060208160051b830101602085015f5b838110156112b757601f19858403018852815180516040855261128f6040860182610f65565b6020928301516001600160401b03169583019590955250978801979190910190600101611269565b50909695505050505050565b606081525f60e08201855160608401526001600160401b0360208701511660808401526001600160401b0360408701511660a08401526060860151608060c0850152818151808452610100860191506020830193505f92505b8083101561133f578351825260208201915060208401935060018301925061131c565b508481036020860152611352818861124d565b935050505061136c60408301846001600160401b03169052565b949350505050565b5f60208284031215611384575f5ffd5b81356001600160401b03811115611399575f5ffd5b61136c84828501611066565b600181811c908216806113b957607f821691505b60208210810361107657634e487b7160e01b5f52602260045260245ffd5b5f602082840312156113e7575f5ffd5b813563ffffffff811681146113fa575f5ffd5b9392505050565b5f5f8335601e19843603018112611416575f5ffd5b8301803591506001600160401b0382111561142f575f5ffd5b602001915036819003821315611443575f5ffd5b9250929050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f82601f830112611487575f5ffd5b815161149561118082611101565b8181528460208386010111156114a9575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f602082840312156114d5575f5ffd5b81516001600160401b038111156114ea575f5ffd5b8201604081850312156114fb575f5ffd5b604080519081016001600160401b038111828210171561151d5761151d6110bd565b60405281516001600160401b03811115611535575f5ffd5b61154186828501611478565b82525060208201516001600160401b0381111561155c575f5ffd5b61156886828501611478565b602083015250949350505050565b63ffffffff60e01b8560e01b168152836004820152818360248301375f91016024019081529392505050565b606081525f8451604060608401526115bd60a0840182610f65565b90506020860151605f198483030160808501526115da8282610f65565b91505082810360208401526115ef8186610f65565b9050828103604084015260a08101845482526001850160a0602084015281815480845260c08501915060c08160051b8601019350825f5260205f2092505f5b818110156116e55760bf19868603018352604085525f845461164f816113a5565b806040890152600182165f811461166d5760018114611689576116ba565b60ff19831660608a0152606082151560051b8a010193506116ba565b875f5260205f205f5b838110156116b15781548b820160600152600190910190602001611692565b8a016060019450505b5050506001858101546001600160401b0316602097880152909560029095019493909301920161162e565b5050505060028501546001600160401b0381166040840152604081901c6001600160401b03166060840152608081811c6001600160401b031690840152509695505050505050565b5f6020828403121561173d575f5ffd5b815180151581146113fa575f5ffd5b5f6020828403121561175c575f5ffd5b81356001600160401b03811681146113fa575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b6001600160401b0381811683821601908111156117a5576117a5611772565b92915050565b5f82518060208501845e5f920191825250919050565b5f602082840312156117d1575f5ffd5b5051919050565b6001600160401b0382811682821603908111156117a5576117a5611772565b634e487b7160e01b5f52603260045260245ffd5b601f8211156109a957805f5260205f20601f840160051c810160208510156118305750805b601f840160051c820191505b8181101561184f575f815560010161183c565b5050505050565b81516001600160401b0381111561186f5761186f6110bd565b6118838161187d84546113a5565b8461180b565b6020601f8211600181146118b5575f831561189e5750848201515b5f19600385901b1c1916600184901b17845561184f565b5f84815260208120601f198516915b828110156118e457878501518255602094850194600190920191016118c4565b508482101561190157868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea2646970667358221220fc19b16d0db1549b278f0e2597de32afb5634deade973f673111fbfaf45dd16c64736f6c634300081e0033",
}

// AvalancheValidatorSetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.ABI instead.
var AvalancheValidatorSetRegistryABI = AvalancheValidatorSetRegistryMetaData.ABI

// AvalancheValidatorSetRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.Bin instead.
var AvalancheValidatorSetRegistryBin = AvalancheValidatorSetRegistryMetaData.Bin

// DeployAvalancheValidatorSetRegistry deploys a new Ethereum contract, binding an instance of AvalancheValidatorSetRegistry to it.
func DeployAvalancheValidatorSetRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, avalancheNetworkID_ uint32, initialValidatorSetData ValidatorSetMetadata) (common.Address, *types.Transaction, *AvalancheValidatorSetRegistry, error) {
	parsed, err := AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AvalancheValidatorSetRegistryBin), backend, avalancheNetworkID_, initialValidatorSetData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AvalancheValidatorSetRegistry{AvalancheValidatorSetRegistryCaller: AvalancheValidatorSetRegistryCaller{contract: contract}, AvalancheValidatorSetRegistryTransactor: AvalancheValidatorSetRegistryTransactor{contract: contract}, AvalancheValidatorSetRegistryFilterer: AvalancheValidatorSetRegistryFilterer{contract: contract}}, nil
}

// AvalancheValidatorSetRegistry is an auto generated Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistry struct {
	AvalancheValidatorSetRegistryCaller     // Read-only binding to the contract
	AvalancheValidatorSetRegistryTransactor // Write-only binding to the contract
	AvalancheValidatorSetRegistryFilterer   // Log filterer for contract events
}

// AvalancheValidatorSetRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvalancheValidatorSetRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvalancheValidatorSetRegistrySession struct {
	Contract     *AvalancheValidatorSetRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvalancheValidatorSetRegistryCallerSession struct {
	Contract *AvalancheValidatorSetRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// AvalancheValidatorSetRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvalancheValidatorSetRegistryTransactorSession struct {
	Contract     *AvalancheValidatorSetRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryRaw struct {
	Contract *AvalancheValidatorSetRegistry // Generic contract binding to access the raw methods on
}

// AvalancheValidatorSetRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryCallerRaw struct {
	Contract *AvalancheValidatorSetRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AvalancheValidatorSetRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetRegistryTransactorRaw struct {
	Contract *AvalancheValidatorSetRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvalancheValidatorSetRegistry creates a new instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistry(address common.Address, backend bind.ContractBackend) (*AvalancheValidatorSetRegistry, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistry{AvalancheValidatorSetRegistryCaller: AvalancheValidatorSetRegistryCaller{contract: contract}, AvalancheValidatorSetRegistryTransactor: AvalancheValidatorSetRegistryTransactor{contract: contract}, AvalancheValidatorSetRegistryFilterer: AvalancheValidatorSetRegistryFilterer{contract: contract}}, nil
}

// NewAvalancheValidatorSetRegistryCaller creates a new read-only instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryCaller(address common.Address, caller bind.ContractCaller) (*AvalancheValidatorSetRegistryCaller, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryCaller{contract: contract}, nil
}

// NewAvalancheValidatorSetRegistryTransactor creates a new write-only instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AvalancheValidatorSetRegistryTransactor, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryTransactor{contract: contract}, nil
}

// NewAvalancheValidatorSetRegistryFilterer creates a new log filterer instance of AvalancheValidatorSetRegistry, bound to a specific deployed contract.
func NewAvalancheValidatorSetRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AvalancheValidatorSetRegistryFilterer, error) {
	contract, err := bindAvalancheValidatorSetRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryFilterer{contract: contract}, nil
}

// bindAvalancheValidatorSetRegistry binds a generic wrapper to an already deployed contract.
func bindAvalancheValidatorSetRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheValidatorSetRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.contract.Transact(opts, method, params...)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) AvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "avalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) AvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) AvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.AvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) GetAvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "getAvalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) GetAvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetAvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) GetAvalancheNetworkID() (uint32, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetAvalancheNetworkID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) GetValidatorSet(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "getValidatorSet", avalancheBlockchainID)

	if err != nil {
		return *new(ValidatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSet)).(*ValidatorSet)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetValidatorSet(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _AvalancheValidatorSetRegistry.Contract.GetValidatorSet(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) IsRegistered(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "isRegistered", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.IsRegistered(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.IsRegistered(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) IsRegistrationInProgress(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "isRegistrationInProgress", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.IsRegistrationInProgress(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.IsRegistrationInProgress(&_AvalancheValidatorSetRegistry.CallOpts, avalancheBlockchainID)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) PChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "pChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) PChainID() ([32]byte, error) {
	return _AvalancheValidatorSetRegistry.Contract.PChainID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) PChainID() ([32]byte, error) {
	return _AvalancheValidatorSetRegistry.Contract.PChainID(&_AvalancheValidatorSetRegistry.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) PChainInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "pChainInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) PChainInitialized() (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.PChainInitialized(&_AvalancheValidatorSetRegistry.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) PChainInitialized() (bool, error) {
	return _AvalancheValidatorSetRegistry.Contract.PChainInitialized(&_AvalancheValidatorSetRegistry.CallOpts)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) ParseValidatorSetMetadata(opts *bind.CallOpts, icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "parseValidatorSetMetadata", icmMessage, shardBytes)

	if err != nil {
		return *new(ValidatorSetMetadata), *new([]Validator), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSetMetadata)).(*ValidatorSetMetadata)
	out1 := *abi.ConvertType(out[1], new([]Validator)).(*[]Validator)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetRegistry.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetRegistry.CallOpts, icmMessage, shardBytes)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetRegistry.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetRegistry.CallOpts, icmMessage, shardBytes)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) VerifyICMMessage(opts *bind.CallOpts, message ICMMessage, avalancheBlockchainID [32]byte) error {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "verifyICMMessage", message, avalancheBlockchainID)

	if err != nil {
		return err
	}

	return err

}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _AvalancheValidatorSetRegistry.Contract.VerifyICMMessage(&_AvalancheValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _AvalancheValidatorSetRegistry.Contract.VerifyICMMessage(&_AvalancheValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactor) ApplyShard(opts *bind.TransactOpts, shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.contract.Transact(opts, "applyShard", shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.ApplyShard(&_AvalancheValidatorSetRegistry.TransactOpts, shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorSession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.ApplyShard(&_AvalancheValidatorSetRegistry.TransactOpts, shard, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactor) RegisterValidatorSet(opts *bind.TransactOpts, message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.contract.Transact(opts, "registerValidatorSet", message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) RegisterValidatorSet(message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.RegisterValidatorSet(&_AvalancheValidatorSetRegistry.TransactOpts, message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorSession) RegisterValidatorSet(message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.RegisterValidatorSet(&_AvalancheValidatorSetRegistry.TransactOpts, message, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactor) UpdateValidatorSet(opts *bind.TransactOpts, shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.contract.Transact(opts, "updateValidatorSet", shard, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) UpdateValidatorSet(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.UpdateValidatorSet(&_AvalancheValidatorSetRegistry.TransactOpts, shard, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorSession) UpdateValidatorSet(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.UpdateValidatorSet(&_AvalancheValidatorSetRegistry.TransactOpts, shard, shardBytes)
}

// UpdateValidatorSetWithDiff is a paid mutator transaction binding the contract method 0x9fd530d4.
//
// Solidity: function updateValidatorSetWithDiff((bytes,uint32,bytes32,bytes) message) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactor) UpdateValidatorSetWithDiff(opts *bind.TransactOpts, message ICMMessage) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.contract.Transact(opts, "updateValidatorSetWithDiff", message)
}

// UpdateValidatorSetWithDiff is a paid mutator transaction binding the contract method 0x9fd530d4.
//
// Solidity: function updateValidatorSetWithDiff((bytes,uint32,bytes32,bytes) message) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) UpdateValidatorSetWithDiff(message ICMMessage) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.UpdateValidatorSetWithDiff(&_AvalancheValidatorSetRegistry.TransactOpts, message)
}

// UpdateValidatorSetWithDiff is a paid mutator transaction binding the contract method 0x9fd530d4.
//
// Solidity: function updateValidatorSetWithDiff((bytes,uint32,bytes32,bytes) message) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactorSession) UpdateValidatorSetWithDiff(message ICMMessage) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.UpdateValidatorSetWithDiff(&_AvalancheValidatorSetRegistry.TransactOpts, message)
}

// AvalancheValidatorSetRegistryValidatorSetRegisteredIterator is returned from FilterValidatorSetRegistered and is used to iterate over the raw logs and unpacked data for ValidatorSetRegistered events raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetRegisteredIterator struct {
	Event *AvalancheValidatorSetRegistryValidatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalancheValidatorSetRegistryValidatorSetRegistered)
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
		it.Event = new(AvalancheValidatorSetRegistryValidatorSetRegistered)
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
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalancheValidatorSetRegistryValidatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalancheValidatorSetRegistryValidatorSetRegistered represents a ValidatorSetRegistered event raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetRegistered struct {
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetRegistered is a free log retrieval operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) FilterValidatorSetRegistered(opts *bind.FilterOpts, avalancheBlockchainID [][32]byte) (*AvalancheValidatorSetRegistryValidatorSetRegisteredIterator, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.FilterLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryValidatorSetRegisteredIterator{contract: _AvalancheValidatorSetRegistry.contract, event: "ValidatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorSetRegistered is a free log subscription operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) WatchValidatorSetRegistered(opts *bind.WatchOpts, sink chan<- *AvalancheValidatorSetRegistryValidatorSetRegistered, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.WatchLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalancheValidatorSetRegistryValidatorSetRegistered)
				if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
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
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) ParseValidatorSetRegistered(log types.Log) (*AvalancheValidatorSetRegistryValidatorSetRegistered, error) {
	event := new(AvalancheValidatorSetRegistryValidatorSetRegistered)
	if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvalancheValidatorSetRegistryValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetUpdatedIterator struct {
	Event *AvalancheValidatorSetRegistryValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalancheValidatorSetRegistryValidatorSetUpdated)
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
		it.Event = new(AvalancheValidatorSetRegistryValidatorSetUpdated)
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
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalancheValidatorSetRegistryValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalancheValidatorSetRegistryValidatorSetUpdated represents a ValidatorSetUpdated event raised by the AvalancheValidatorSetRegistry contract.
type AvalancheValidatorSetRegistryValidatorSetUpdated struct {
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2.
//
// Solidity: event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, avalancheBlockchainID [][32]byte) (*AvalancheValidatorSetRegistryValidatorSetUpdatedIterator, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.FilterLogs(opts, "ValidatorSetUpdated", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetRegistryValidatorSetUpdatedIterator{contract: _AvalancheValidatorSetRegistry.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2.
//
// Solidity: event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *AvalancheValidatorSetRegistryValidatorSetUpdated, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _AvalancheValidatorSetRegistry.contract.WatchLogs(opts, "ValidatorSetUpdated", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalancheValidatorSetRegistryValidatorSetUpdated)
				if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2.
//
// Solidity: event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryFilterer) ParseValidatorSetUpdated(log types.Log) (*AvalancheValidatorSetRegistryValidatorSetUpdated, error) {
	event := new(AvalancheValidatorSetRegistryValidatorSetUpdated)
	if err := _AvalancheValidatorSetRegistry.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
