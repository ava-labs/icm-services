// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subsetupdater

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

// SubsetUpdaterMetaData contains all meta data concerning the SubsetUpdater contract.
var SubsetUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405260405161241c38038061241c83398101604081905261002291610204565b63ffffffff8216608052805160a08190525f90815260016020818152604092839020818501518154948601516001600160401b0390811668010000000000000000026001600160801b0319909616911617939093178355606084015180518694869490936100969391850192910190610120565b50600401805460ff60401b19166801000000000000000090811790915560a0515f908152602081815260409182902084518155908401516002909101805494909201516001600160401b03908116600160801b02600160801b600160c01b03199190921690930292909216600160401b600160c01b0319909316929092171790555061032c915050565b828054828255905f5260205f20908101928215610159579160200282015b8281111561015957825182559160200191906001019061013e565b50610165929150610169565b5090565b5b80821115610165575f815560010161016a565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156101b3576101b361017d565b60405290565b604051601f8201601f191681016001600160401b03811182821017156101e1576101e161017d565b604052919050565b80516001600160401b03811681146101ff575f5ffd5b919050565b5f5f60408385031215610215575f5ffd5b825163ffffffff81168114610228575f5ffd5b60208401519092506001600160401b03811115610243575f5ffd5b830160808186031215610254575f5ffd5b61025c610191565b8151815261026c602083016101e9565b602082015261027d604083016101e9565b604082015260608201516001600160401b0381111561029a575f5ffd5b80830192505085601f8301126102ae575f5ffd5b81516001600160401b038111156102c7576102c761017d565b8060051b6102d7602082016101b9565b918252602081850181019290810190898411156102f2575f5ffd5b6020860195505b83861015610318578551808352602096870196909350909101906102f9565b606085015250949791965090945050505050565b60805160a0516120a66103765f395f818160f1015281816101380152818161026301526108fe01525f8181610190015281816101c90152818161037601526107e501526120a65ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c806368531ed01161006e57806368531ed01461018b57806382366d05146101c75780638457eaa7146101ed5780638e91cb4314610219578063933568401461022c5780639def1e781461023f575f5ffd5b806327258b22146100aa578063541dcba4146100ec57806357262e7f14610121578063580d632b146101365780636766233d14610178575b5f5ffd5b6100d76100b836600461145d565b5f908152602081905260409020600201546001600160401b0316151590565b60405190151581526020015b60405180910390f35b6101137f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100e3565b61013461012f36600461148a565b610261565b005b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b031615156100d7565b61013461018636600461157f565b610591565b6101b27f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100e3565b7f00000000000000000000000000000000000000000000000000000000000000006101b2565b6100d76101fb36600461145d565b5f90815260016020526040902060040154600160401b900460ff1690565b610134610227366004611610565b6107de565b61013461023a36600461157f565b610c86565b61025261024d366004611610565b610f89565b6040516100e393929190611749565b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b03166102f95760405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b60648201526084015b60405180910390fd5b5f818152602081905260409020600201546001600160401b031661036f5760405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b60648201526084016102f0565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166103a660408401602085016117fa565b63ffffffff16146103ef5760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b60448201526064016102f0565b5f73__$aaf4ae346b84a712cc43f25bb66199d6fb$__63858ad3986104176060860186611824565b6040518363ffffffff1660e01b815260040161043492919061186d565b5f60405180830381865af415801561044e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261047591908101906118e8565b90505f61048860408501602086016117fa565b60408501356104978680611824565b6040516020016104aa949392919061197c565b60408051601f198184030181528282525f868152602081905291909120630161c9f960e61b835290925073__$aaf4ae346b84a712cc43f25bb66199d6fb$__916358727e409161050091869186916004016119da565b602060405180830381865af415801561051b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061053f9190611b65565b61058b5760405162461bcd60e51b815260206004820152601b60248201527f4661696c656420746f20766572696679207369676e617475726573000000000060448201526064016102f0565b50505050565b6105b882602001355f9081526001602052604090206004015460ff600160401b9091041690565b6106045760405162461bcd60e51b815260206004820152601f60248201527f526567697374726174696f6e206973206e6f7420696e2070726f67726573730060448201526064016102f0565b602082018035906106159084611b98565b5f828152600160208190526040909120600201546001600160401b039283169261064192911690611bc7565b6001600160401b0316146106975760405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f72646572000000000060448201526064016102f0565b6002826040516106a79190611bec565b602060405180830381855afa1580156106c2573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906106e59190611c02565b5f82815260016020818152604090922081019161070490870187611b98565b61070e9190611c19565b6001600160401b03168154811061072757610727611c38565b905f5260205f200154146107755760405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b60448201526064016102f0565b61077f8383610c86565b6107a683602001355f9081526001602052604090206004015460ff600160401b9091041690565b6107d9576040516020840135907f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2905f90a25b505050565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661081560408501602086016117fa565b63ffffffff161461085e5760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b60448201526064016102f0565b6040808401355f90815260016020522060040154600160401b900460ff16156108d75760405162461bcd60e51b815260206004820152602560248201527f4120726567697374726174696f6e20697320616c726561647920696e2070726f604482015264677265737360d81b60648201526084016102f0565b6040808401355f908152602081905220600201546001600160401b031661092757610922837f0000000000000000000000000000000000000000000000000000000000000000610261565b610935565b610935838460400135610261565b5f5f5f610943868686610f89565b82519295509093509150604087013581146109a05760405162461bcd60e51b815260206004820152601860248201527f536f7572636520636861696e204944206d69736d61746368000000000000000060448201526064016102f0565b825160608501515160011015610b635784515f90815260016020818152604092839020818901518154948a01516001600160401b03908116600160401b026001600160801b031990961691161793909317835560608801518051610a0b938501929190910190611318565b5060028101805467ffffffffffffffff1916600117905560048101805468ffffffffffffffffff19166001600160401b03861617600160401b1790555f5b82811015610ac95781600301868281518110610a6757610a67611c38565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610a989082611c90565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610a49565b506040808a01355f908152602081905220600201546001600160401b0316610b5d5785515f9081526020818152604091829020885181559088015160029091018054928901516001600160401b03908116600160801b0267ffffffffffffffff60801b1991909316600160401b021677ffffffffffffffffffffffffffffffff000000000000000019909316929092171790555b50610c52565b84515f9081526020818152604080832088518155600281018054938a0151928a01516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909616918a16919091179490941792909216929092179055905b82811015610c4f5781600101868281518110610bed57610bed611c38565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610c1e9082611c90565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610bcf565b50505b60405182907f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93905f90a25050505050505050565b60405163b9a1525960e01b81526020830135905f90819073__$aaf4ae346b84a712cc43f25bb66199d6fb$__9063b9a1525990610cc7908790600401611d4d565b5f60405180830381865af4158015610ce1573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d089190810190611d91565b915091505f825111610d5c5760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d70747900000060448201526064016102f0565b5f816001600160401b031611610db45760405162461bcd60e51b815260206004820152601a60248201527f546f74616c20776569676874206d75737420657863656564203000000000000060448201526064016102f0565b5f5b8251811015610e475760015f8581526020019081526020015f20600301838281518110610de557610de5611c38565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610e169082611c90565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610db6565b505f8381526001602052604081206004018054839290610e719084906001600160401b0316611bc7565b82546101009290920a6001600160401b038181021990931691831602179091555f8581526001602081905260408220600201805491945092610eb591859116611bc7565b82546001600160401b039182166101009390930a9283029190920219909116179055505f838152600160208181526040909220015490610ef790870187611b98565b6001600160401b031603610f82575f83815260016020818152604080842060048101805468ff0000000000000000191690559184905290922060039092018054610f449390920191611361565b505f8381526001602090815260408083206004015491839052909120600201805467ffffffffffffffff19166001600160401b039092169190911790555b5050505050565b604080516080810182525f808252602082018190529181019190915260608082015260605f8073__$aaf4ae346b84a712cc43f25bb66199d6fb$__63b70e3f03610fd38980611824565b6040518363ffffffff1660e01b8152600401610ff092919061186d565b5f60405180830381865af415801561100a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110319190810190611ea0565b905060028686604051611045929190611f95565b602060405180830381855afa158015611060573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906110839190611c02565b81606001515f8151811061109957611099611c38565b6020026020010151146110ee5760405162461bcd60e51b815260206004820152601b60248201527f56616c696461746f72207365742068617368206d69736d61746368000000000060448201526064016102f0565b5f5f73__$aaf4ae346b84a712cc43f25bb66199d6fb$__63b9a1525989896040518363ffffffff1660e01b815260040161112992919061186d565b5f60405180830381865af4158015611143573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261116a9190810190611d91565b84516020808701515f83815291829052604090912060020154939550919350916001600160401b03918216600160401b909104909116106111e65760405162461bcd60e51b8152602060048201526016602482015275502d436861696e2068656967687420746f6f206c6f7760501b60448201526064016102f0565b6040808501515f838152602081905291909120600201546001600160401b03918216600160801b909104909116106112605760405162461bcd60e51b815260206004820152601960248201527f502d436861696e2074696d657374616d7020746f6f206c6f770000000000000060448201526064016102f0565b5f8351116112b05760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d70747900000060448201526064016102f0565b5f826001600160401b0316116113085760405162461bcd60e51b815260206004820152601a60248201527f546f74616c20776569676874206d75737420657863656564203000000000000060448201526064016102f0565b5091989097509095509350505050565b828054828255905f5260205f20908101928215611351579160200282015b82811115611351578251825591602001919060010190611336565b5061135d9291506113e0565b5090565b828054828255905f5260205f209060020281019282156113d4575f5260205f209160020282015b828111156113d45782828061139d8382611fa4565b506001918201549101805467ffffffffffffffff19166001600160401b039092169190911790556002928301929190910190611388565b5061135d9291506113f4565b5b8082111561135d575f81556001016113e1565b8082111561135d575f6114078282611423565b5060018101805467ffffffffffffffff191690556002016113f4565b50805461142f906119a8565b5f825580601f1061143e575050565b601f0160209004905f5260205f209081019061145a91906113e0565b50565b5f6020828403121561146d575f5ffd5b5035919050565b5f60808284031215611484575f5ffd5b50919050565b5f5f6040838503121561149b575f5ffd5b82356001600160401b038111156114b0575f5ffd5b6114bc85828601611474565b95602094909401359450505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611501576115016114cb565b60405290565b604051608081016001600160401b0381118282101715611501576115016114cb565b604051601f8201601f191681016001600160401b0381118282101715611551576115516114cb565b604052919050565b5f6001600160401b03821115611571576115716114cb565b50601f01601f191660200190565b5f5f8284036060811215611591575f5ffd5b604081121561159e575f5ffd5b5082915060408301356001600160401b038111156115ba575f5ffd5b8301601f810185136115ca575f5ffd5b80356115dd6115d882611559565b611529565b8181528660208385010111156115f1575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f60408486031215611622575f5ffd5b83356001600160401b03811115611637575f5ffd5b61164386828701611474565b93505060208401356001600160401b0381111561165e575f5ffd5b8401601f8101861361166e575f5ffd5b80356001600160401b03811115611683575f5ffd5b866020828401011115611694575f5ffd5b939660209190910195509293505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208501945060208160051b830101602085015f5b8381101561173d57601f19858403018852815180516040855261171560408601826116a5565b6020928301516001600160401b031695830195909552509788019791909101906001016116ef565b50909695505050505050565b606081525f60e08201855160608401526001600160401b0360208701511660808401526001600160401b0360408701511660a08401526060860151608060c0850152818151808452610100860191506020830193505f92505b808310156117c557835182526020820191506020840193506001830192506117a2565b5084810360208601526117d881886116d3565b93505050506117f260408301846001600160401b03169052565b949350505050565b5f6020828403121561180a575f5ffd5b813563ffffffff8116811461181d575f5ffd5b9392505050565b5f5f8335601e19843603018112611839575f5ffd5b8301803591506001600160401b03821115611852575f5ffd5b602001915036819003821315611866575f5ffd5b9250929050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f82601f8301126118aa575f5ffd5b81516118b86115d882611559565b8181528460208386010111156118cc575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f602082840312156118f8575f5ffd5b81516001600160401b0381111561190d575f5ffd5b82016040818503121561191e575f5ffd5b6119266114df565b81516001600160401b0381111561193b575f5ffd5b6119478682850161189b565b82525060208201516001600160401b03811115611962575f5ffd5b61196e8682850161189b565b602083015250949350505050565b63ffffffff60e01b8560e01b168152836004820152818360248301375f91016024019081529392505050565b600181811c908216806119bc57607f821691505b60208210810361148457634e487b7160e01b5f52602260045260245ffd5b606081525f8451604060608401526119f560a08401826116a5565b90506020860151605f19848303016080850152611a1282826116a5565b9150508281036020840152611a2781866116a5565b9050828103604084015260a08101845482526001850160a0602084015281815480845260c08501915060c08160051b8601019350825f5260205f2092505f5b81811015611b1d5760bf19868603018352604085525f8454611a87816119a8565b806040890152600182165f8114611aa55760018114611ac157611af2565b60ff19831660608a0152606082151560051b8a01019350611af2565b875f5260205f205f5b83811015611ae95781548b820160600152600190910190602001611aca565b8a016060019450505b5050506001858101546001600160401b03166020978801529095600290950194939093019201611a66565b5050505060028501546001600160401b0381166040840152604081901c6001600160401b03166060840152608081811c6001600160401b031690840152509695505050505050565b5f60208284031215611b75575f5ffd5b8151801515811461181d575f5ffd5b6001600160401b038116811461145a575f5ffd5b5f60208284031215611ba8575f5ffd5b813561181d81611b84565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190811115611be657611be6611bb3565b92915050565b5f82518060208501845e5f920191825250919050565b5f60208284031215611c12575f5ffd5b5051919050565b6001600160401b038281168282160390811115611be657611be6611bb3565b634e487b7160e01b5f52603260045260245ffd5b601f8211156107d957805f5260205f20601f840160051c81016020851015611c715750805b601f840160051c820191505b81811015610f82575f8155600101611c7d565b81516001600160401b03811115611ca957611ca96114cb565b611cbd81611cb784546119a8565b84611c4c565b6020601f821160018114611cf2575f8315611cd85750848201515b600184901b5f19600386901b1c198216175b855550610f82565b5f84815260208120601f198516915b82811015611d215787850151825560209485019460019092019101611d01565b5084821015611d3e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f61181d60208301846116a5565b5f6001600160401b03821115611d7757611d776114cb565b5060051b60200190565b8051611d8c81611b84565b919050565b5f5f60408385031215611da2575f5ffd5b82516001600160401b03811115611db7575f5ffd5b8301601f81018513611dc7575f5ffd5b8051611dd56115d882611d5f565b8082825260208201915060208360051b850101925087831115611df6575f5ffd5b602084015b83811015611e845780516001600160401b03811115611e18575f5ffd5b85016040818b03601f19011215611e2d575f5ffd5b611e356114df565b60208201516001600160401b03811115611e4d575f5ffd5b611e5c8c60208386010161189b565b82525060408201519150611e6f82611b84565b60208181019290925284529283019201611dfb565b509450611e979250505060208401611d81565b90509250929050565b5f60208284031215611eb0575f5ffd5b81516001600160401b03811115611ec5575f5ffd5b820160808185031215611ed6575f5ffd5b611ede611507565b815181526020820151611ef081611b84565b60208201526040820151611f0381611b84565b604082015260608201516001600160401b03811115611f20575f5ffd5b80830192505084601f830112611f34575f5ffd5b8151611f426115d882611d5f565b8082825260208201915060208360051b860101925087831115611f63575f5ffd5b6020850194505b82851015611f85578451825260209485019490910190611f6a565b6060840152509095945050505050565b818382375f9101908152919050565b818103611faf575050565b611fb982546119a8565b6001600160401b03811115611fd057611fd06114cb565b611fde81611cb784546119a8565b5f601f82116001811461200d575f8315611cd8575081850154600184901b5f19600386901b1c19821617611cea565b5f8581526020808220868352908220601f198616925b838110156120435782860154825560019586019590910190602001612023565b508583101561206057818501545f19600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220735c17828e1070cf18dc647e1fd51a59e63df552c9b38bbd1d03fb220746299464736f6c634300081e0033",
}

// SubsetUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use SubsetUpdaterMetaData.ABI instead.
var SubsetUpdaterABI = SubsetUpdaterMetaData.ABI

// SubsetUpdaterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubsetUpdaterMetaData.Bin instead.
var SubsetUpdaterBin = SubsetUpdaterMetaData.Bin

// DeploySubsetUpdater deploys a new Ethereum contract, binding an instance of SubsetUpdater to it.
func DeploySubsetUpdater(auth *bind.TransactOpts, backend bind.ContractBackend, avalancheNetworkID_ uint32, initialValidatorSetData ValidatorSetMetadata) (common.Address, *types.Transaction, *SubsetUpdater, error) {
	parsed, err := SubsetUpdaterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubsetUpdaterBin), backend, avalancheNetworkID_, initialValidatorSetData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubsetUpdater{SubsetUpdaterCaller: SubsetUpdaterCaller{contract: contract}, SubsetUpdaterTransactor: SubsetUpdaterTransactor{contract: contract}, SubsetUpdaterFilterer: SubsetUpdaterFilterer{contract: contract}}, nil
}

// SubsetUpdater is an auto generated Go binding around an Ethereum contract.
type SubsetUpdater struct {
	SubsetUpdaterCaller     // Read-only binding to the contract
	SubsetUpdaterTransactor // Write-only binding to the contract
	SubsetUpdaterFilterer   // Log filterer for contract events
}

// SubsetUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubsetUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubsetUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubsetUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubsetUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubsetUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubsetUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubsetUpdaterSession struct {
	Contract     *SubsetUpdater    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubsetUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubsetUpdaterCallerSession struct {
	Contract *SubsetUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SubsetUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubsetUpdaterTransactorSession struct {
	Contract     *SubsetUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubsetUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubsetUpdaterRaw struct {
	Contract *SubsetUpdater // Generic contract binding to access the raw methods on
}

// SubsetUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubsetUpdaterCallerRaw struct {
	Contract *SubsetUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// SubsetUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubsetUpdaterTransactorRaw struct {
	Contract *SubsetUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubsetUpdater creates a new instance of SubsetUpdater, bound to a specific deployed contract.
func NewSubsetUpdater(address common.Address, backend bind.ContractBackend) (*SubsetUpdater, error) {
	contract, err := bindSubsetUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdater{SubsetUpdaterCaller: SubsetUpdaterCaller{contract: contract}, SubsetUpdaterTransactor: SubsetUpdaterTransactor{contract: contract}, SubsetUpdaterFilterer: SubsetUpdaterFilterer{contract: contract}}, nil
}

// NewSubsetUpdaterCaller creates a new read-only instance of SubsetUpdater, bound to a specific deployed contract.
func NewSubsetUpdaterCaller(address common.Address, caller bind.ContractCaller) (*SubsetUpdaterCaller, error) {
	contract, err := bindSubsetUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdaterCaller{contract: contract}, nil
}

// NewSubsetUpdaterTransactor creates a new write-only instance of SubsetUpdater, bound to a specific deployed contract.
func NewSubsetUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*SubsetUpdaterTransactor, error) {
	contract, err := bindSubsetUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdaterTransactor{contract: contract}, nil
}

// NewSubsetUpdaterFilterer creates a new log filterer instance of SubsetUpdater, bound to a specific deployed contract.
func NewSubsetUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*SubsetUpdaterFilterer, error) {
	contract, err := bindSubsetUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdaterFilterer{contract: contract}, nil
}

// bindSubsetUpdater binds a generic wrapper to an already deployed contract.
func bindSubsetUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubsetUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubsetUpdater *SubsetUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubsetUpdater.Contract.SubsetUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubsetUpdater *SubsetUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.SubsetUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubsetUpdater *SubsetUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.SubsetUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubsetUpdater *SubsetUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubsetUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubsetUpdater *SubsetUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubsetUpdater *SubsetUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.contract.Transact(opts, method, params...)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterCaller) AvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "avalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterSession) AvalancheNetworkID() (uint32, error) {
	return _SubsetUpdater.Contract.AvalancheNetworkID(&_SubsetUpdater.CallOpts)
}

// AvalancheNetworkID is a free data retrieval call binding the contract method 0x68531ed0.
//
// Solidity: function avalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterCallerSession) AvalancheNetworkID() (uint32, error) {
	return _SubsetUpdater.Contract.AvalancheNetworkID(&_SubsetUpdater.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterCaller) GetAvalancheNetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "getAvalancheNetworkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterSession) GetAvalancheNetworkID() (uint32, error) {
	return _SubsetUpdater.Contract.GetAvalancheNetworkID(&_SubsetUpdater.CallOpts)
}

// GetAvalancheNetworkID is a free data retrieval call binding the contract method 0x82366d05.
//
// Solidity: function getAvalancheNetworkID() view returns(uint32)
func (_SubsetUpdater *SubsetUpdaterCallerSession) GetAvalancheNetworkID() (uint32, error) {
	return _SubsetUpdater.Contract.GetAvalancheNetworkID(&_SubsetUpdater.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCaller) IsRegistered(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "isRegistered", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _SubsetUpdater.Contract.IsRegistered(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCallerSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _SubsetUpdater.Contract.IsRegistered(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCaller) IsRegistrationInProgress(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "isRegistrationInProgress", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterSession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _SubsetUpdater.Contract.IsRegistrationInProgress(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCallerSession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _SubsetUpdater.Contract.IsRegistrationInProgress(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_SubsetUpdater *SubsetUpdaterCaller) PChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "pChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_SubsetUpdater *SubsetUpdaterSession) PChainID() ([32]byte, error) {
	return _SubsetUpdater.Contract.PChainID(&_SubsetUpdater.CallOpts)
}

// PChainID is a free data retrieval call binding the contract method 0x541dcba4.
//
// Solidity: function pChainID() view returns(bytes32)
func (_SubsetUpdater *SubsetUpdaterCallerSession) PChainID() ([32]byte, error) {
	return _SubsetUpdater.Contract.PChainID(&_SubsetUpdater.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCaller) PChainInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "pChainInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_SubsetUpdater *SubsetUpdaterSession) PChainInitialized() (bool, error) {
	return _SubsetUpdater.Contract.PChainInitialized(&_SubsetUpdater.CallOpts)
}

// PChainInitialized is a free data retrieval call binding the contract method 0x580d632b.
//
// Solidity: function pChainInitialized() view returns(bool)
func (_SubsetUpdater *SubsetUpdaterCallerSession) PChainInitialized() (bool, error) {
	return _SubsetUpdater.Contract.PChainInitialized(&_SubsetUpdater.CallOpts)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_SubsetUpdater *SubsetUpdaterCaller) ParseValidatorSetMetadata(opts *bind.CallOpts, icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "parseValidatorSetMetadata", icmMessage, shardBytes)

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
func (_SubsetUpdater *SubsetUpdaterSession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _SubsetUpdater.Contract.ParseValidatorSetMetadata(&_SubsetUpdater.CallOpts, icmMessage, shardBytes)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_SubsetUpdater *SubsetUpdaterCallerSession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _SubsetUpdater.Contract.ParseValidatorSetMetadata(&_SubsetUpdater.CallOpts, icmMessage, shardBytes)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_SubsetUpdater *SubsetUpdaterCaller) VerifyICMMessage(opts *bind.CallOpts, message ICMMessage, avalancheBlockchainID [32]byte) error {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "verifyICMMessage", message, avalancheBlockchainID)

	if err != nil {
		return err
	}

	return err

}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_SubsetUpdater *SubsetUpdaterSession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _SubsetUpdater.Contract.VerifyICMMessage(&_SubsetUpdater.CallOpts, message, avalancheBlockchainID)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x57262e7f.
//
// Solidity: function verifyICMMessage((bytes,uint32,bytes32,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_SubsetUpdater *SubsetUpdaterCallerSession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _SubsetUpdater.Contract.VerifyICMMessage(&_SubsetUpdater.CallOpts, message, avalancheBlockchainID)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactor) ApplyShard(opts *bind.TransactOpts, shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.contract.Transact(opts, "applyShard", shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterSession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.ApplyShard(&_SubsetUpdater.TransactOpts, shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactorSession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.ApplyShard(&_SubsetUpdater.TransactOpts, shard, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactor) RegisterValidatorSet(opts *bind.TransactOpts, message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.contract.Transact(opts, "registerValidatorSet", message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterSession) RegisterValidatorSet(message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.RegisterValidatorSet(&_SubsetUpdater.TransactOpts, message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0x8e91cb43.
//
// Solidity: function registerValidatorSet((bytes,uint32,bytes32,bytes) message, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactorSession) RegisterValidatorSet(message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.RegisterValidatorSet(&_SubsetUpdater.TransactOpts, message, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactor) UpdateValidatorSet(opts *bind.TransactOpts, shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.contract.Transact(opts, "updateValidatorSet", shard, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterSession) UpdateValidatorSet(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.UpdateValidatorSet(&_SubsetUpdater.TransactOpts, shard, shardBytes)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x6766233d.
//
// Solidity: function updateValidatorSet((uint64,bytes32) shard, bytes shardBytes) returns()
func (_SubsetUpdater *SubsetUpdaterTransactorSession) UpdateValidatorSet(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _SubsetUpdater.Contract.UpdateValidatorSet(&_SubsetUpdater.TransactOpts, shard, shardBytes)
}

// SubsetUpdaterValidatorSetRegisteredIterator is returned from FilterValidatorSetRegistered and is used to iterate over the raw logs and unpacked data for ValidatorSetRegistered events raised by the SubsetUpdater contract.
type SubsetUpdaterValidatorSetRegisteredIterator struct {
	Event *SubsetUpdaterValidatorSetRegistered // Event containing the contract specifics and raw log

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
func (it *SubsetUpdaterValidatorSetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubsetUpdaterValidatorSetRegistered)
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
		it.Event = new(SubsetUpdaterValidatorSetRegistered)
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
func (it *SubsetUpdaterValidatorSetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubsetUpdaterValidatorSetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubsetUpdaterValidatorSetRegistered represents a ValidatorSetRegistered event raised by the SubsetUpdater contract.
type SubsetUpdaterValidatorSetRegistered struct {
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetRegistered is a free log retrieval operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_SubsetUpdater *SubsetUpdaterFilterer) FilterValidatorSetRegistered(opts *bind.FilterOpts, avalancheBlockchainID [][32]byte) (*SubsetUpdaterValidatorSetRegisteredIterator, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _SubsetUpdater.contract.FilterLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdaterValidatorSetRegisteredIterator{contract: _SubsetUpdater.contract, event: "ValidatorSetRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorSetRegistered is a free log subscription operation binding the contract event 0x715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93.
//
// Solidity: event ValidatorSetRegistered(bytes32 indexed avalancheBlockchainID)
func (_SubsetUpdater *SubsetUpdaterFilterer) WatchValidatorSetRegistered(opts *bind.WatchOpts, sink chan<- *SubsetUpdaterValidatorSetRegistered, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _SubsetUpdater.contract.WatchLogs(opts, "ValidatorSetRegistered", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubsetUpdaterValidatorSetRegistered)
				if err := _SubsetUpdater.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
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
func (_SubsetUpdater *SubsetUpdaterFilterer) ParseValidatorSetRegistered(log types.Log) (*SubsetUpdaterValidatorSetRegistered, error) {
	event := new(SubsetUpdaterValidatorSetRegistered)
	if err := _SubsetUpdater.contract.UnpackLog(event, "ValidatorSetRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubsetUpdaterValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the SubsetUpdater contract.
type SubsetUpdaterValidatorSetUpdatedIterator struct {
	Event *SubsetUpdaterValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *SubsetUpdaterValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubsetUpdaterValidatorSetUpdated)
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
		it.Event = new(SubsetUpdaterValidatorSetUpdated)
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
func (it *SubsetUpdaterValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubsetUpdaterValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubsetUpdaterValidatorSetUpdated represents a ValidatorSetUpdated event raised by the SubsetUpdater contract.
type SubsetUpdaterValidatorSetUpdated struct {
	AvalancheBlockchainID [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2.
//
// Solidity: event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID)
func (_SubsetUpdater *SubsetUpdaterFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, avalancheBlockchainID [][32]byte) (*SubsetUpdaterValidatorSetUpdatedIterator, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _SubsetUpdater.contract.FilterLogs(opts, "ValidatorSetUpdated", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &SubsetUpdaterValidatorSetUpdatedIterator{contract: _SubsetUpdater.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2.
//
// Solidity: event ValidatorSetUpdated(bytes32 indexed avalancheBlockchainID)
func (_SubsetUpdater *SubsetUpdaterFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *SubsetUpdaterValidatorSetUpdated, avalancheBlockchainID [][32]byte) (event.Subscription, error) {

	var avalancheBlockchainIDRule []interface{}
	for _, avalancheBlockchainIDItem := range avalancheBlockchainID {
		avalancheBlockchainIDRule = append(avalancheBlockchainIDRule, avalancheBlockchainIDItem)
	}

	logs, sub, err := _SubsetUpdater.contract.WatchLogs(opts, "ValidatorSetUpdated", avalancheBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubsetUpdaterValidatorSetUpdated)
				if err := _SubsetUpdater.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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
func (_SubsetUpdater *SubsetUpdaterFilterer) ParseValidatorSetUpdated(log types.Log) (*SubsetUpdaterValidatorSetUpdated, error) {
	event := new(SubsetUpdaterValidatorSetUpdated)
	if err := _SubsetUpdater.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
