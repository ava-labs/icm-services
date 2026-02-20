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

// SubsetUpdaterMetaData contains all meta data concerning the SubsetUpdater contract.
var SubsetUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"rawMessage\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405260405161269838038061269883398101604081905261002291610204565b63ffffffff8216608052805160a08190525f90815260016020818152604092839020818501518154948601516001600160401b0390811668010000000000000000026001600160801b0319909616911617939093178355606084015180518694869490936100969391850192910190610120565b50600401805460ff60401b19166801000000000000000090811790915560a0515f908152602081815260409182902084518155908401516002909101805494909201516001600160401b03908116600160801b02600160801b600160c01b03199190921690930292909216600160401b600160c01b0319909316929092171790555061032c915050565b828054828255905f5260205f20908101928215610159579160200282015b8281111561015957825182559160200191906001019061013e565b50610165929150610169565b5090565b5b80821115610165575f815560010161016a565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156101b3576101b361017d565b60405290565b604051601f8201601f191681016001600160401b03811182821017156101e1576101e161017d565b604052919050565b80516001600160401b03811681146101ff575f5ffd5b919050565b5f5f60408385031215610215575f5ffd5b825163ffffffff81168114610228575f5ffd5b60208401519092506001600160401b03811115610243575f5ffd5b830160808186031215610254575f5ffd5b61025c610191565b8151815261026c602083016101e9565b602082015261027d604083016101e9565b604082015260608201516001600160401b0381111561029a575f5ffd5b80830192505085601f8301126102ae575f5ffd5b81516001600160401b038111156102c7576102c761017d565b8060051b6102d7602082016101b9565b918252602081850181019290810190898411156102f2575f5ffd5b6020860195505b83861015610318578551808352602096870196909350909101906102f9565b606085015250949791965090945050505050565b60805160a0516123216103775f395f818161011c015281816101630152818161040b0152610aa601525f81816101bb015281816101f40152818161051e015261098d01526123215ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c806368531ed01161006e57806368531ed0146101b657806382366d05146101f25780638457eaa7146102185780638e91cb431461024457806393356840146102575780639def1e781461026a575f5ffd5b806327258b22146100b557806349e4db9c146100f7578063541dcba41461011757806357262e7f1461014c578063580d632b146101615780636766233d146101a3575b5f5ffd5b6100e26100c3366004611605565b5f908152602081905260409020600201546001600160401b0316151590565b60405190151581526020015b60405180910390f35b61010a610105366004611605565b61028c565b6040516100ee919061164a565b61013e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100ee565b61015f61015a366004611733565b610409565b005b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b031615156100e2565b61015f6101b1366004611828565b610739565b6101dd7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100ee565b7f00000000000000000000000000000000000000000000000000000000000000006101dd565b6100e2610226366004611605565b5f90815260016020526040902060040154600160401b900460ff1690565b61015f6102523660046118b9565b610986565b61015f610265366004611828565b610e2e565b61027d6102783660046118b9565b611131565b6040516100ee939291906119c4565b6040805160a0810182525f80825260606020830181905292820181905291810182905260808101919091525f82815260208181526040808320815160a081018352815481526001820180548451818702810187019095528085529195929486810194939192919084015b828210156103c8578382905f5260205f2090600202016040518060400160405290815f8201805461032690611a75565b80601f016020809104026020016040519081016040528092919081815260200182805461035290611a75565b801561039d5780601f106103745761010080835404028352916020019161039d565b820191905f5260205f20905b81548152906001019060200180831161038057829003601f168201915b50505091835250506001918201546001600160401b03166020918201529183529290920191016102f6565b50505090825250600291909101546001600160401b038082166020840152600160401b820481166040840152600160801b9091041660609091015292915050565b7f00000000000000000000000000000000000000000000000000000000000000005f908152602081905260409020600201546001600160401b03166104a15760405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b60648201526084015b60405180910390fd5b5f818152602081905260409020600201546001600160401b03166105175760405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b6064820152608401610498565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661054e6040840160208501611aa7565b63ffffffff16146105975760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606401610498565b5f73__$aaf4ae346b84a712cc43f25bb66199d6fb$__63858ad3986105bf6060860186611ad1565b6040518363ffffffff1660e01b81526004016105dc929190611b1a565b5f60405180830381865af41580156105f6573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261061d9190810190611b95565b90505f6106306040850160208601611aa7565b604085013561063f8680611ad1565b6040516020016106529493929190611c29565b60408051601f198184030181528282525f868152602081905291909120630161c9f960e61b835290925073__$aaf4ae346b84a712cc43f25bb66199d6fb$__916358727e40916106a89186918691600401611c55565b602060405180830381865af41580156106c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e79190611de0565b6107335760405162461bcd60e51b815260206004820152601b60248201527f4661696c656420746f20766572696679207369676e61747572657300000000006044820152606401610498565b50505050565b61076082602001355f9081526001602052604090206004015460ff600160401b9091041690565b6107ac5760405162461bcd60e51b815260206004820152601f60248201527f526567697374726174696f6e206973206e6f7420696e2070726f6772657373006044820152606401610498565b602082018035906107bd9084611e13565b5f828152600160208190526040909120600201546001600160401b03928316926107e992911690611e42565b6001600160401b03161461083f5760405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f7264657200000000006044820152606401610498565b60028260405161084f9190611e67565b602060405180830381855afa15801561086a573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061088d9190611e7d565b5f8281526001602081815260409092208101916108ac90870187611e13565b6108b69190611e94565b6001600160401b0316815481106108cf576108cf611eb3565b905f5260205f2001541461091d5760405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b6044820152606401610498565b6109278383610e2e565b61094e83602001355f9081526001602052604090206004015460ff600160401b9091041690565b610981576040516020840135907f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe2905f90a25b505050565b63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166109bd6040850160208601611aa7565b63ffffffff1614610a065760405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606401610498565b6040808401355f90815260016020522060040154600160401b900460ff1615610a7f5760405162461bcd60e51b815260206004820152602560248201527f4120726567697374726174696f6e20697320616c726561647920696e2070726f604482015264677265737360d81b6064820152608401610498565b6040808401355f908152602081905220600201546001600160401b0316610acf57610aca837f0000000000000000000000000000000000000000000000000000000000000000610409565b610add565b610add838460400135610409565b5f5f5f610aeb868686611131565b8251929550909350915060408701358114610b485760405162461bcd60e51b815260206004820152601860248201527f536f7572636520636861696e204944206d69736d6174636800000000000000006044820152606401610498565b825160608501515160011015610d0b5784515f90815260016020818152604092839020818901518154948a01516001600160401b03908116600160401b026001600160801b031990961691161793909317835560608801518051610bb39385019291909101906114c0565b5060028101805467ffffffffffffffff1916600117905560048101805468ffffffffffffffffff19166001600160401b03861617600160401b1790555f5b82811015610c715781600301868281518110610c0f57610c0f611eb3565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610c409082611f0b565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610bf1565b506040808a01355f908152602081905220600201546001600160401b0316610d055785515f9081526020818152604091829020885181559088015160029091018054928901516001600160401b03908116600160801b0267ffffffffffffffff60801b1991909316600160401b021677ffffffffffffffffffffffffffffffff000000000000000019909316929092171790555b50610dfa565b84515f9081526020818152604080832088518155600281018054938a0151928a01516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909616918a16919091179490941792909216929092179055905b82811015610df75781600101868281518110610d9557610d95611eb3565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610dc69082611f0b565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610d77565b50505b60405182907f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce93905f90a25050505050505050565b60405163b9a1525960e01b81526020830135905f90819073__$aaf4ae346b84a712cc43f25bb66199d6fb$__9063b9a1525990610e6f908790600401611fc8565b5f60405180830381865af4158015610e89573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610eb0919081019061200c565b915091505f825111610f045760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d7074790000006044820152606401610498565b5f816001600160401b031611610f5c5760405162461bcd60e51b815260206004820152601a60248201527f546f74616c20776569676874206d7573742065786365656420300000000000006044820152606401610498565b5f5b8251811015610fef5760015f8581526020019081526020015f20600301838281518110610f8d57610f8d611eb3565b60209081029190910181015182546001810184555f938452919092208251600290920201908190610fbe9082611f0b565b50602091909101516001918201805467ffffffffffffffff19166001600160401b0390921691909117905501610f5e565b505f83815260016020526040812060040180548392906110199084906001600160401b0316611e42565b82546101009290920a6001600160401b038181021990931691831602179091555f858152600160208190526040822060020180549194509261105d91859116611e42565b82546001600160401b039182166101009390930a9283029190920219909116179055505f83815260016020818152604090922001549061109f90870187611e13565b6001600160401b03160361112a575f83815260016020818152604080842060048101805468ff00000000000000001916905591849052909220600390920180546110ec9390920191611509565b505f8381526001602090815260408083206004015491839052909120600201805467ffffffffffffffff19166001600160401b039092169190911790555b5050505050565b604080516080810182525f808252602082018190529181019190915260608082015260605f8073__$aaf4ae346b84a712cc43f25bb66199d6fb$__63b70e3f0361117b8980611ad1565b6040518363ffffffff1660e01b8152600401611198929190611b1a565b5f60405180830381865af41580156111b2573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526111d9919081019061211b565b9050600286866040516111ed929190612210565b602060405180830381855afa158015611208573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061122b9190611e7d565b81606001515f8151811061124157611241611eb3565b6020026020010151146112965760405162461bcd60e51b815260206004820152601b60248201527f56616c696461746f72207365742068617368206d69736d6174636800000000006044820152606401610498565b5f5f73__$aaf4ae346b84a712cc43f25bb66199d6fb$__63b9a1525989896040518363ffffffff1660e01b81526004016112d1929190611b1a565b5f60405180830381865af41580156112eb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611312919081019061200c565b84516020808701515f83815291829052604090912060020154939550919350916001600160401b03918216600160401b9091049091161061138e5760405162461bcd60e51b8152602060048201526016602482015275502d436861696e2068656967687420746f6f206c6f7760501b6044820152606401610498565b6040808501515f838152602081905291909120600201546001600160401b03918216600160801b909104909116106114085760405162461bcd60e51b815260206004820152601960248201527f502d436861696e2074696d657374616d7020746f6f206c6f77000000000000006044820152606401610498565b5f8351116114585760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d7074790000006044820152606401610498565b5f826001600160401b0316116114b05760405162461bcd60e51b815260206004820152601a60248201527f546f74616c20776569676874206d7573742065786365656420300000000000006044820152606401610498565b5091989097509095509350505050565b828054828255905f5260205f209081019282156114f9579160200282015b828111156114f95782518255916020019190600101906114de565b50611505929150611588565b5090565b828054828255905f5260205f2090600202810192821561157c575f5260205f209160020282015b8281111561157c57828280611545838261221f565b506001918201549101805467ffffffffffffffff19166001600160401b039092169190911790556002928301929190910190611530565b5061150592915061159c565b5b80821115611505575f8155600101611589565b80821115611505575f6115af82826115cb565b5060018101805467ffffffffffffffff1916905560020161159c565b5080546115d790611a75565b5f825580601f106115e6575050565b601f0160209004905f5260205f20908101906116029190611588565b50565b5f60208284031215611615575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f60c0820183516020840152602084015160a0604085015281815180845260e08601915060e08160051b87010193506020830192505f5b818110156116d35760df1987860301835283518051604087526116ab604088018261161c565b6020928301516001600160401b03169783019790975250938401939290920191600101611685565b5050505060408401516001600160401b03811660608501525060608401516001600160401b03811660808501525060808401516001600160401b03811660a0850152509392505050565b5f6080828403121561172d575f5ffd5b50919050565b5f5f60408385031215611744575f5ffd5b82356001600160401b03811115611759575f5ffd5b6117658582860161171d565b95602094909401359450505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156117aa576117aa611774565b60405290565b604051608081016001600160401b03811182821017156117aa576117aa611774565b604051601f8201601f191681016001600160401b03811182821017156117fa576117fa611774565b604052919050565b5f6001600160401b0382111561181a5761181a611774565b50601f01601f191660200190565b5f5f828403606081121561183a575f5ffd5b6040811215611847575f5ffd5b5082915060408301356001600160401b03811115611863575f5ffd5b8301601f81018513611873575f5ffd5b803561188661188182611802565b6117d2565b81815286602083850101111561189a575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f604084860312156118cb575f5ffd5b83356001600160401b038111156118e0575f5ffd5b6118ec8682870161171d565b93505060208401356001600160401b03811115611907575f5ffd5b8401601f81018613611917575f5ffd5b80356001600160401b0381111561192c575f5ffd5b86602082840101111561193d575f5ffd5b939660209190910195509293505050565b5f82825180855260208501945060208160051b830101602085015f5b838110156119b857601f198584030188528151805160408552611990604086018261161c565b6020928301516001600160401b0316958301959095525097880197919091019060010161196a565b50909695505050505050565b606081525f60e08201855160608401526001600160401b0360208701511660808401526001600160401b0360408701511660a08401526060860151608060c0850152818151808452610100860191506020830193505f92505b80831015611a405783518252602082019150602084019350600183019250611a1d565b508481036020860152611a53818861194e565b9350505050611a6d60408301846001600160401b03169052565b949350505050565b600181811c90821680611a8957607f821691505b60208210810361172d57634e487b7160e01b5f52602260045260245ffd5b5f60208284031215611ab7575f5ffd5b813563ffffffff81168114611aca575f5ffd5b9392505050565b5f5f8335601e19843603018112611ae6575f5ffd5b8301803591506001600160401b03821115611aff575f5ffd5b602001915036819003821315611b13575f5ffd5b9250929050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f82601f830112611b57575f5ffd5b8151611b6561188182611802565b818152846020838601011115611b79575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f60208284031215611ba5575f5ffd5b81516001600160401b03811115611bba575f5ffd5b820160408185031215611bcb575f5ffd5b611bd3611788565b81516001600160401b03811115611be8575f5ffd5b611bf486828501611b48565b82525060208201516001600160401b03811115611c0f575f5ffd5b611c1b86828501611b48565b602083015250949350505050565b63ffffffff60e01b8560e01b168152836004820152818360248301375f91016024019081529392505050565b606081525f845160406060840152611c7060a084018261161c565b90506020860151605f19848303016080850152611c8d828261161c565b9150508281036020840152611ca2818661161c565b9050828103604084015260a08101845482526001850160a0602084015281815480845260c08501915060c08160051b8601019350825f5260205f2092505f5b81811015611d985760bf19868603018352604085525f8454611d0281611a75565b806040890152600182165f8114611d205760018114611d3c57611d6d565b60ff19831660608a0152606082151560051b8a01019350611d6d565b875f5260205f205f5b83811015611d645781548b820160600152600190910190602001611d45565b8a016060019450505b5050506001858101546001600160401b03166020978801529095600290950194939093019201611ce1565b5050505060028501546001600160401b0381166040840152604081901c6001600160401b03166060840152608081811c6001600160401b031690840152509695505050505050565b5f60208284031215611df0575f5ffd5b81518015158114611aca575f5ffd5b6001600160401b0381168114611602575f5ffd5b5f60208284031215611e23575f5ffd5b8135611aca81611dff565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190811115611e6157611e61611e2e565b92915050565b5f82518060208501845e5f920191825250919050565b5f60208284031215611e8d575f5ffd5b5051919050565b6001600160401b038281168282160390811115611e6157611e61611e2e565b634e487b7160e01b5f52603260045260245ffd5b601f82111561098157805f5260205f20601f840160051c81016020851015611eec5750805b601f840160051c820191505b8181101561112a575f8155600101611ef8565b81516001600160401b03811115611f2457611f24611774565b611f3881611f328454611a75565b84611ec7565b6020601f821160018114611f6d575f8315611f535750848201515b600184901b5f19600386901b1c198216175b85555061112a565b5f84815260208120601f198516915b82811015611f9c5787850151825560209485019460019092019101611f7c565b5084821015611fb957868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f611aca602083018461161c565b5f6001600160401b03821115611ff257611ff2611774565b5060051b60200190565b805161200781611dff565b919050565b5f5f6040838503121561201d575f5ffd5b82516001600160401b03811115612032575f5ffd5b8301601f81018513612042575f5ffd5b805161205061188182611fda565b8082825260208201915060208360051b850101925087831115612071575f5ffd5b602084015b838110156120ff5780516001600160401b03811115612093575f5ffd5b85016040818b03601f190112156120a8575f5ffd5b6120b0611788565b60208201516001600160401b038111156120c8575f5ffd5b6120d78c602083860101611b48565b825250604082015191506120ea82611dff565b60208181019290925284529283019201612076565b5094506121129250505060208401611ffc565b90509250929050565b5f6020828403121561212b575f5ffd5b81516001600160401b03811115612140575f5ffd5b820160808185031215612151575f5ffd5b6121596117b0565b81518152602082015161216b81611dff565b6020820152604082015161217e81611dff565b604082015260608201516001600160401b0381111561219b575f5ffd5b80830192505084601f8301126121af575f5ffd5b81516121bd61188182611fda565b8082825260208201915060208360051b8601019250878311156121de575f5ffd5b6020850194505b828510156122005784518252602094850194909101906121e5565b6060840152509095945050505050565b818382375f9101908152919050565b81810361222a575050565b6122348254611a75565b6001600160401b0381111561224b5761224b611774565b61225981611f328454611a75565b5f601f821160018114612288575f8315611f53575081850154600184901b5f19600386901b1c19821617611f65565b5f8581526020808220868352908220601f198616925b838110156122be578286015482556001958601959091019060200161229e565b50858310156122db57818501545f19600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122099864462d7b82b9ac4ed20355317ed2994b8dacf9a365ddfaaa00f4aca7dd3d864736f6c634300081e0033",
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

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_SubsetUpdater *SubsetUpdaterCaller) GetValidatorSet(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	var out []interface{}
	err := _SubsetUpdater.contract.Call(opts, &out, "getValidatorSet", avalancheBlockchainID)

	if err != nil {
		return *new(ValidatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSet)).(*ValidatorSet)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_SubsetUpdater *SubsetUpdaterSession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _SubsetUpdater.Contract.GetValidatorSet(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_SubsetUpdater *SubsetUpdaterCallerSession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _SubsetUpdater.Contract.GetValidatorSet(&_SubsetUpdater.CallOpts, avalancheBlockchainID)
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
