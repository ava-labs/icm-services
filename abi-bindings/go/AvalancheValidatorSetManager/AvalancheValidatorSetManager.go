// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avalanchevalidatorsetmanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
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
	Message         ICMRawMessage
	RawMessageBytes []byte
	Attestation     []byte
}

// ICMRawMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMRawMessage struct {
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	SourceAddress      common.Address
	VerifierAddress    common.Address
	Payload            []byte
}

// PartialValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type PartialValidatorSet struct {
	PChainHeight    uint64
	PChainTimestamp uint64
	ShardHashes     [][32]byte
	ShardsReceived  uint64
	Validators      []Validator
	NumValidators   *big.Int
	PartialWeight   uint64
	InProgress      bool
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
	TotalValidators       uint64
	ShardHashes           [][32]byte
}

// ValidatorSetShard is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetShard struct {
	ShardNumber           uint64
	AvalancheBlockchainID [32]byte
}

// AvalancheValidatorSetManagerMetaData contains all meta data concerning the AvalancheValidatorSetManager contract.
var AvalancheValidatorSetManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getShardHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getShardsReceived\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"shardsReceived\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"numValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"partialWeight\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"inProgress\",\"type\":\"bool\"}],\"internalType\":\"structPartialValidatorSet\",\"name\":\"partialValidatorSet\",\"type\":\"tuple\"}],\"name\":\"setPartialValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"validatorSet\",\"type\":\"tuple\"}],\"name\":\"setValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611a45908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816308e64d7414610ca957508063141eaf4a14610c6857806327258b2214610c3057806349e4db9c146109e95780638457eaa7146109b4578063893d20e81461098d57806393356840146106ae578063c49abdfa146103a5578063c4d66de814610253578063e8a6c940146100cf5763eb3b597d14610095575f80fd5b346100cb5760203660031901126100cb576004355f52600260205260206001600160401b03600260405f20015416604051908152f35b5f80fd5b346100cb5760403660031901126100cb576024356001600160401b0381116100cb5760a060031982360301126100cb576040519061010c8261120c565b8060040135825260248101356001600160401b0381116100cb5761013690600436918401016112ef565b906020830191825261014a604482016112c4565b92604081019384526101716084610163606485016112c4565b9360608401948552016112c4565b926080820193845261018d60018060a01b035f541633146115f0565b6004355f52600160205260405f20915182556001820190519060208251926101b584846117cb565b01905f5260205f205f915b838310610235578751600286018054885167ffffffffffffffff60401b60409190911b166001600160401b039093166fffffffffffffffffffffffffffffffff19909116179190911781558751815467ffffffffffffffff60801b191660809190911b67ffffffffffffffff60801b16179055005b600260208261024760019451866116ca565b019201920191906101c0565b346100cb5760203660031901126100cb576004356001600160a01b038116908190036100cb575f5160206119f05f395f51905f52549060ff8260401c1615916001600160401b0381168015908161039d575b6001149081610393575b15908161038a575b5061037b5767ffffffffffffffff1981166001175f5160206119f05f395f51905f52558261034f575b506bffffffffffffffffffffffff60a01b5f5416175f556102fd57005b60ff60401b195f5160206119f05f395f51905f5254165f5160206119f05f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b68ffffffffffffffffff191668010000000000000001175f5160206119f05f395f51905f5255826102e0565b63f92ee8a960e01b5f5260045ffd5b905015846102b7565b303b1591506102af565b8491506102a5565b346100cb5760403660031901126100cb576024356001600160401b0381116100cb5761010060031982360301126100cb5760405161010081018181106001600160401b0382111761068a576040526103ff826004016112c4565b815261040d602483016112c4565b906020810191825260448301356001600160401b0381116100cb578301366023820112156100cb57600481013590610444826112d8565b916104526040519384611242565b808352602060048185019260051b84010101913683116100cb57602401905b82821061069e575050506040820190815261048e606485016112c4565b6060830190815260848501356001600160401b0381116100cb576104b890600436918801016112ef565b906080840191825260a084019260a4870135845260e46104da60c489016112c4565b9760c0870198895201359586151587036100cb576001600160401b0361055c9160e0880198895261051560018060a01b035f541633146115f0565b6004355f526002602052818060405f209951161682198954161788555116869067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b5180519060018601906001600160401b03831161068a57600160401b831161068a578154838355808410610664575b50602001905f5260205f205f5b83811061065057505050506001600160401b039051166001600160401b036002850191166001600160401b03198254161790556003830190519060208251926105e184846117cb565b01905f5260205f205f915b83831061063257845160048701558751600587018054895160ff60401b90151560401b166001600160401b0390931668ffffffffffffffffff1990911617919091179055005b600260208261064460019451866116ca565b019201920191906105ec565b600190602084519401938184015501610598565b825f528360205f2091820191015b81811061067f575061058b565b5f8155600101610672565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610471565b346100cb573660031901606081126100cb576040136100cb576044356001600160401b0381116100cb576106e690369060040161127e565b6106fa60018060a01b035f541633146115f0565b6107245f602435926040518093819263b9a1525960e01b83526020600484015260248301906111e8565b038173__$aaf4ae346b84a712cc43f25bb66199d6fb$__5af48015610982575f915f91610955575b506001600160401b039061076283511515611514565b169061076f821515611560565b825f526002602052600460405f200154905f5b81518110156107f25781518110156107de5760208160051b8301015190855f526002602052600360405f20016107b88286611662565b9281548410156107de576001936107d8925f5260205f2090851b016116ca565b01610782565b634e487b7160e01b5f52603260045260245ffd5b50905051825f526002602052610810600460405f2001918254611662565b9055815f526002602052600560405f2001906001600160401b03825416016001600160401b038111610941576001600160401b03166001600160401b0319825416179055805f526002602052600260405f200160016001600160401b03825416016001600160401b038111610941576001600160401b03166001600160401b03198254161790556004356001600160401b0381168091036100cb57815f526002602052600160405f200154146108c257005b5f81815260026020908152604080832060058101805468ff00000000000000001916905560019283905292206108fc926003019101611968565b805f5260026020526001600160401b03600560405f20015416905f5260016020526001600160401b03600260405f200191166001600160401b03198254161790555f80f35b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03925061097b91503d805f833e6109738183611242565b8101906113ed565b909161074c565b6040513d5f823e3d90fd5b346100cb575f3660031901126100cb575f546040516001600160a01b039091168152602090f35b346100cb5760203660031901126100cb576004355f526002602052602060ff600560405f20015460401c166040519015158152f35b346100cb5760203660031901126100cb575f6080604051610a098161120c565b8281526060602082015282604082015282606082015201526004355f52600160205260405f2060405190610a3c8261120c565b805482526001810190815491610a51836112d8565b92610a5f6040519485611242565b80845260208401915f5260205f205f925b828410610b6c57868660028760208401928352015491604081016001600160401b038416815260608201906001600160401b038560401c1682526001600160401b03608084019560801c168552604051936020855260c0850193516020860152519260a06040860152835180915260e0850190602060e08260051b8801019501915f905b828210610b24578780886001600160401b038c818b818c5116606087015251166080850152511660a08301520390f35b9091929560208060019260df198b8203018552895190826001600160401b0381610b5785516040865260408601906111e8565b94015116910152980192019201909291610af4565b604051610b7881611227565b6040515f8454610b87816115b8565b8084529060018116908115610c0d5750600114610bd6575b509260029282610bb56020946001970382611242565b81526001600160401b03858701541683820152815201920193019290610a70565b5f868152602081209092505b818310610bf757505081016020016002610b9f565b6001816020925483868801015201920191610be2565b60ff191660208086019190915291151560051b8401909101915060029050610b9f565b346100cb5760203660031901126100cb576004355f52600160205260206001600160401b03600260405f200154161515604051908152f35b346100cb5760403660031901126100cb576024356004355f526002602052600160405f20019081548110156107de576020915f52815f200154604051908152f35b346100cb5760403660031901126100cb57600435906001600160401b0382116100cb578136039160606003198401126100cb57602435926001600160401b0384116100cb57366023850112156100cb578360040135926001600160401b0384116100cb57602485019460248536920101116100cb57608081610d2c60609361120c565b5f81525f60208201525f60408201525f83820152015273__$aaf4ae346b84a712cc43f25bb66199d6fb$__9180600401359160a219018212156100cb5701608481013560221936839003018112156100cb5701600401803593906001600160401b0385116100cb5760200184360381136100cb5760405163b70e3f0360e01b8152945f9186918291610dc29190600484016113c6565b0381855af4938415610982575f946110ed575b50608084019283518051156107de576020015160205f6040518486823780858101838152039060025afa15610982575f51036110a85760405163b9a1525960e01b8152925f92849283918291610e2e91600484016113c6565b03915af4918215610982575f915f9361108a575b50835190815f5260016020526001600160401b03600260405f20015460401c169160208601926001600160401b03845116111561101f575f5260016020526001600160401b03600260405f20015460801c169360408601946001600160401b038651161115610fb55760606001600160401b038092610ec387511515611514565b1696610ed0881515611560565b8260405197838952816101008a01978451868c0152511660808a0152511660a088015201511660c0850152519060a060e08501528151809152602061012085019201905f5b818110610f9f575050508281036020840152815180825260208201916020808360051b8301019401925f915b838310610f55578680878a60408301520390f35b9091929394602080600192601f19858203018652885190826001600160401b0381610f8985516040865260408601906111e8565b9401511691015297019301930191939290610f41565b8251845260209384019390920191600101610f15565b608460405162461bcd60e51b815260206004820152604060248201527f502d436861696e2074696d657374616d70206d7573742062652067726561746560448201527f72207468616e207468652063757272656e742076616c696461746f72207365746064820152fd5b60405162461bcd60e51b815260206004820152603d60248201527f502d436861696e20686569676874206d7573742062652067726561746572207460448201527f68616e207468652063757272656e742076616c696461746f72207365740000006064820152608490fd5b9092506110a191503d805f833e6109738183611242565b9184610e42565b60405162461bcd60e51b815260206004820152601b60248201527f56616c696461746f72207365742068617368206d69736d6174636800000000006044820152606490fd5b9093503d805f833e6110ff8183611242565b8101906020818303126100cb578051906001600160401b0382116100cb57019060a0828203126100cb57604051916111368361120c565b80518352611146602082016113b2565b6020840152611157604082016113b2565b6040840152611168606082016113b2565b60608401526080810151906001600160401b0382116100cb57019080601f830112156100cb578151611199816112d8565b926111a76040519485611242565b81845260208085019260051b8201019283116100cb57602001905b8282106111d85750505060808201529284610dd5565b81518152602091820191016111c2565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60a081019081106001600160401b0382111761068a57604052565b604081019081106001600160401b0382111761068a57604052565b90601f801991011681019081106001600160401b0382111761068a57604052565b6001600160401b03811161068a57601f01601f191660200190565b81601f820112156100cb5780359061129582611263565b926112a36040519485611242565b828452602083830101116100cb57815f926020809301838601378301015290565b35906001600160401b03821682036100cb57565b6001600160401b03811161068a5760051b60200190565b9080601f830112156100cb57813591611307836112d8565b926113156040519485611242565b80845260208085019160051b830101918383116100cb5760208101915b83831061134157505050505090565b82356001600160401b0381116100cb578201906040828703601f1901126100cb576040519061136f82611227565b6020830135916001600160401b0383116100cb576113a360408561139b8b60208099988199010161127e565b8452016112c4565b83820152815201920191611332565b51906001600160401b03821682036100cb57565b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b91906040838203126100cb5782516001600160401b0381116100cb57830181601f820112156100cb57805190611422826112d8565b926114306040519485611242565b82845260208085019360051b830101918183116100cb5760208101935b83851061146957505050505060206114669193016113b2565b90565b84516001600160401b0381116100cb578201906040828503601f1901126100cb576040519061149782611227565b60208301516001600160401b0381116100cb576020908401019185601f840112156100cb5782516114c781611263565b936114d56040519586611242565b81855287602083830101116100cb5760209586955f87856115059682604097018386015e830101528452016113b2565b8382015281520194019361144d565b1561151b57565b60405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d7074790000006044820152606490fd5b1561156757565b60405162461bcd60e51b815260206004820152602360248201527f546f74616c20776569676874206d75737420626520677265617465722074686160448201526206e20360ec1b6064820152608490fd5b90600182811c921680156115e6575b60208310146115d257565b634e487b7160e01b5f52602260045260245ffd5b91607f16916115c7565b156115f757565b60405162461bcd60e51b815260206004820152603760248201527f416e20756e617574686f72697a6564206164647265737320617474656d70746560448201527f6420746f2063616c6c20746869732066756e6374696f6e0000000000000000006064820152608490fd5b9190820180921161094157565b81811061167a575050565b5f815560010161166f565b9190601f811161169457505050565b6116be925f5260205f20906020601f840160051c830193106116c0575b601f0160051c019061166f565b565b90915081906116b1565b91909182519283516001600160401b03811161068a576116f4816116ee85546115b8565b85611685565b6020601f821160011461175a57600192611734836001600160401b039697989460209488965f9261174f575b50508160011b915f199060031b1c19161790565b86555b015116920191166001600160401b0319825416179055565b015190505f80611720565b601f19821695845f52815f20965f5b8181106117b35750836001600160401b0396979860209460019794899789951061179b575b505050811b018655611737565b01515f1960f88460031b161c191690555f808061178e565b83830151895560019098019760209384019301611769565b600160401b821161068a578054908281558183106117e857505050565b6001600160ff1b0382168203610941576001600160ff1b0383168303610941575f5260205f209060011b81019160011b015b818110611825575050565b80611832600292546115b8565b80611845575b505f60018201550161181a565b601f811160011461185b57505f81555b5f611838565b61187890825f526001601f60205f20920160051c8201910161166f565b805f525f6020812081835555611855565b9190918281146119635761189d83546115b8565b6001600160401b03811161068a576118bf816118b984546115b8565b84611685565b5f93601f82116001146118fd576118ee92939482915f926118f25750508160011b915f199060031b1c19161790565b9055565b015490505f80611720565b601f198216905f5260205f2094835f5260205f20915f5b81811061194b57509583600195969710611933575b505050811b019055565b01545f1960f88460031b161c191690555f8080611929565b9192600180602092868b015481550194019201611914565b509050565b8181146119eb5781549161197c83836117cb565b5f5260205f20905f5260205f205f915b8383106119995750505050565b6002808260019385036119b3575b0192019201919061198c565b6119bd8186611889565b6001600160401b0384820154166001600160401b0385870191166001600160401b03198254161790556119a7565b505056fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220782e206c0a39eaa0b8dc3c1fb117de4ca8c1f80cd1944f29d40667b3ccfdd55a64736f6c634300081e0033",
}

// AvalancheValidatorSetManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AvalancheValidatorSetManagerMetaData.ABI instead.
var AvalancheValidatorSetManagerABI = AvalancheValidatorSetManagerMetaData.ABI

// AvalancheValidatorSetManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AvalancheValidatorSetManagerMetaData.Bin instead.
var AvalancheValidatorSetManagerBin = AvalancheValidatorSetManagerMetaData.Bin

// DeployAvalancheValidatorSetManager deploys a new Ethereum contract, binding an instance of AvalancheValidatorSetManager to it.
func DeployAvalancheValidatorSetManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AvalancheValidatorSetManager, error) {
	parsed, err := AvalancheValidatorSetManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AvalancheValidatorSetManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AvalancheValidatorSetManager{AvalancheValidatorSetManagerCaller: AvalancheValidatorSetManagerCaller{contract: contract}, AvalancheValidatorSetManagerTransactor: AvalancheValidatorSetManagerTransactor{contract: contract}, AvalancheValidatorSetManagerFilterer: AvalancheValidatorSetManagerFilterer{contract: contract}}, nil
}

// AvalancheValidatorSetManager is an auto generated Go binding around an Ethereum contract.
type AvalancheValidatorSetManager struct {
	AvalancheValidatorSetManagerCaller     // Read-only binding to the contract
	AvalancheValidatorSetManagerTransactor // Write-only binding to the contract
	AvalancheValidatorSetManagerFilterer   // Log filterer for contract events
}

// AvalancheValidatorSetManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvalancheValidatorSetManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalancheValidatorSetManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvalancheValidatorSetManagerSession struct {
	Contract     *AvalancheValidatorSetManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvalancheValidatorSetManagerCallerSession struct {
	Contract *AvalancheValidatorSetManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// AvalancheValidatorSetManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvalancheValidatorSetManagerTransactorSession struct {
	Contract     *AvalancheValidatorSetManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// AvalancheValidatorSetManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvalancheValidatorSetManagerRaw struct {
	Contract *AvalancheValidatorSetManager // Generic contract binding to access the raw methods on
}

// AvalancheValidatorSetManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvalancheValidatorSetManagerCallerRaw struct {
	Contract *AvalancheValidatorSetManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AvalancheValidatorSetManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvalancheValidatorSetManagerTransactorRaw struct {
	Contract *AvalancheValidatorSetManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvalancheValidatorSetManager creates a new instance of AvalancheValidatorSetManager, bound to a specific deployed contract.
func NewAvalancheValidatorSetManager(address common.Address, backend bind.ContractBackend) (*AvalancheValidatorSetManager, error) {
	contract, err := bindAvalancheValidatorSetManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetManager{AvalancheValidatorSetManagerCaller: AvalancheValidatorSetManagerCaller{contract: contract}, AvalancheValidatorSetManagerTransactor: AvalancheValidatorSetManagerTransactor{contract: contract}, AvalancheValidatorSetManagerFilterer: AvalancheValidatorSetManagerFilterer{contract: contract}}, nil
}

// NewAvalancheValidatorSetManagerCaller creates a new read-only instance of AvalancheValidatorSetManager, bound to a specific deployed contract.
func NewAvalancheValidatorSetManagerCaller(address common.Address, caller bind.ContractCaller) (*AvalancheValidatorSetManagerCaller, error) {
	contract, err := bindAvalancheValidatorSetManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetManagerCaller{contract: contract}, nil
}

// NewAvalancheValidatorSetManagerTransactor creates a new write-only instance of AvalancheValidatorSetManager, bound to a specific deployed contract.
func NewAvalancheValidatorSetManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AvalancheValidatorSetManagerTransactor, error) {
	contract, err := bindAvalancheValidatorSetManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetManagerTransactor{contract: contract}, nil
}

// NewAvalancheValidatorSetManagerFilterer creates a new log filterer instance of AvalancheValidatorSetManager, bound to a specific deployed contract.
func NewAvalancheValidatorSetManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AvalancheValidatorSetManagerFilterer, error) {
	contract, err := bindAvalancheValidatorSetManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetManagerFilterer{contract: contract}, nil
}

// bindAvalancheValidatorSetManager binds a generic wrapper to an already deployed contract.
func bindAvalancheValidatorSetManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvalancheValidatorSetManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetManager.Contract.AvalancheValidatorSetManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.AvalancheValidatorSetManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.AvalancheValidatorSetManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalancheValidatorSetManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) GetOwner() (common.Address, error) {
	return _AvalancheValidatorSetManager.Contract.GetOwner(&_AvalancheValidatorSetManager.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) GetOwner() (common.Address, error) {
	return _AvalancheValidatorSetManager.Contract.GetOwner(&_AvalancheValidatorSetManager.CallOpts)
}

// GetShardHash is a free data retrieval call binding the contract method 0x141eaf4a.
//
// Solidity: function getShardHash(bytes32 avalancheBlockchainID, uint256 index) view returns(bytes32)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) GetShardHash(opts *bind.CallOpts, avalancheBlockchainID [32]byte, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "getShardHash", avalancheBlockchainID, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetShardHash is a free data retrieval call binding the contract method 0x141eaf4a.
//
// Solidity: function getShardHash(bytes32 avalancheBlockchainID, uint256 index) view returns(bytes32)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) GetShardHash(avalancheBlockchainID [32]byte, index *big.Int) ([32]byte, error) {
	return _AvalancheValidatorSetManager.Contract.GetShardHash(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID, index)
}

// GetShardHash is a free data retrieval call binding the contract method 0x141eaf4a.
//
// Solidity: function getShardHash(bytes32 avalancheBlockchainID, uint256 index) view returns(bytes32)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) GetShardHash(avalancheBlockchainID [32]byte, index *big.Int) ([32]byte, error) {
	return _AvalancheValidatorSetManager.Contract.GetShardHash(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID, index)
}

// GetShardsReceived is a free data retrieval call binding the contract method 0xeb3b597d.
//
// Solidity: function getShardsReceived(bytes32 avalancheBlockchainID) view returns(uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) GetShardsReceived(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (uint64, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "getShardsReceived", avalancheBlockchainID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetShardsReceived is a free data retrieval call binding the contract method 0xeb3b597d.
//
// Solidity: function getShardsReceived(bytes32 avalancheBlockchainID) view returns(uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) GetShardsReceived(avalancheBlockchainID [32]byte) (uint64, error) {
	return _AvalancheValidatorSetManager.Contract.GetShardsReceived(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// GetShardsReceived is a free data retrieval call binding the contract method 0xeb3b597d.
//
// Solidity: function getShardsReceived(bytes32 avalancheBlockchainID) view returns(uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) GetShardsReceived(avalancheBlockchainID [32]byte) (uint64, error) {
	return _AvalancheValidatorSetManager.Contract.GetShardsReceived(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) GetValidatorSet(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "getValidatorSet", avalancheBlockchainID)

	if err != nil {
		return *new(ValidatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSet)).(*ValidatorSet)

	return out0, err

}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _AvalancheValidatorSetManager.Contract.GetValidatorSet(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// GetValidatorSet is a free data retrieval call binding the contract method 0x49e4db9c.
//
// Solidity: function getValidatorSet(bytes32 avalancheBlockchainID) view returns((bytes32,(bytes,uint64)[],uint64,uint64,uint64))
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) GetValidatorSet(avalancheBlockchainID [32]byte) (ValidatorSet, error) {
	return _AvalancheValidatorSetManager.Contract.GetValidatorSet(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) IsRegistered(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "isRegistered", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetManager.Contract.IsRegistered(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) IsRegistered(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetManager.Contract.IsRegistered(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) IsRegistrationInProgress(opts *bind.CallOpts, avalancheBlockchainID [32]byte) (bool, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "isRegistrationInProgress", avalancheBlockchainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetManager.Contract.IsRegistrationInProgress(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// IsRegistrationInProgress is a free data retrieval call binding the contract method 0x8457eaa7.
//
// Solidity: function isRegistrationInProgress(bytes32 avalancheBlockchainID) view returns(bool)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) IsRegistrationInProgress(avalancheBlockchainID [32]byte) (bool, error) {
	return _AvalancheValidatorSetManager.Contract.IsRegistrationInProgress(&_AvalancheValidatorSetManager.CallOpts, avalancheBlockchainID)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x08e64d74.
//
// Solidity: function parseValidatorSetMetadata(((uint32,bytes32,address,address,bytes),bytes,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCaller) ParseValidatorSetMetadata(opts *bind.CallOpts, icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	var out []interface{}
	err := _AvalancheValidatorSetManager.contract.Call(opts, &out, "parseValidatorSetMetadata", icmMessage, shardBytes)

	if err != nil {
		return *new(ValidatorSetMetadata), *new([]Validator), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorSetMetadata)).(*ValidatorSetMetadata)
	out1 := *abi.ConvertType(out[1], new([]Validator)).(*[]Validator)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x08e64d74.
//
// Solidity: function parseValidatorSetMetadata(((uint32,bytes32,address,address,bytes),bytes,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetManager.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetManager.CallOpts, icmMessage, shardBytes)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x08e64d74.
//
// Solidity: function parseValidatorSetMetadata(((uint32,bytes32,address,address,bytes),bytes,bytes) icmMessage, bytes shardBytes) view returns((bytes32,uint64,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerCallerSession) ParseValidatorSetMetadata(icmMessage ICMMessage, shardBytes []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetManager.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetManager.CallOpts, icmMessage, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactor) ApplyShard(opts *bind.TransactOpts, shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.contract.Transact(opts, "applyShard", shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.ApplyShard(&_AvalancheValidatorSetManager.TransactOpts, shard, shardBytes)
}

// ApplyShard is a paid mutator transaction binding the contract method 0x93356840.
//
// Solidity: function applyShard((uint64,bytes32) shard, bytes shardBytes) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorSession) ApplyShard(shard ValidatorSetShard, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.ApplyShard(&_AvalancheValidatorSetManager.TransactOpts, shard, shardBytes)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.Initialize(&_AvalancheValidatorSetManager.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.Initialize(&_AvalancheValidatorSetManager.TransactOpts, owner)
}

// SetPartialValidatorSet is a paid mutator transaction binding the contract method 0xc49abdfa.
//
// Solidity: function setPartialValidatorSet(bytes32 avalancheBlockchainID, (uint64,uint64,bytes32[],uint64,(bytes,uint64)[],uint256,uint64,bool) partialValidatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactor) SetPartialValidatorSet(opts *bind.TransactOpts, avalancheBlockchainID [32]byte, partialValidatorSet PartialValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.contract.Transact(opts, "setPartialValidatorSet", avalancheBlockchainID, partialValidatorSet)
}

// SetPartialValidatorSet is a paid mutator transaction binding the contract method 0xc49abdfa.
//
// Solidity: function setPartialValidatorSet(bytes32 avalancheBlockchainID, (uint64,uint64,bytes32[],uint64,(bytes,uint64)[],uint256,uint64,bool) partialValidatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) SetPartialValidatorSet(avalancheBlockchainID [32]byte, partialValidatorSet PartialValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.SetPartialValidatorSet(&_AvalancheValidatorSetManager.TransactOpts, avalancheBlockchainID, partialValidatorSet)
}

// SetPartialValidatorSet is a paid mutator transaction binding the contract method 0xc49abdfa.
//
// Solidity: function setPartialValidatorSet(bytes32 avalancheBlockchainID, (uint64,uint64,bytes32[],uint64,(bytes,uint64)[],uint256,uint64,bool) partialValidatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorSession) SetPartialValidatorSet(avalancheBlockchainID [32]byte, partialValidatorSet PartialValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.SetPartialValidatorSet(&_AvalancheValidatorSetManager.TransactOpts, avalancheBlockchainID, partialValidatorSet)
}

// SetValidatorSet is a paid mutator transaction binding the contract method 0xe8a6c940.
//
// Solidity: function setValidatorSet(bytes32 avalancheBlockchainID, (bytes32,(bytes,uint64)[],uint64,uint64,uint64) validatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactor) SetValidatorSet(opts *bind.TransactOpts, avalancheBlockchainID [32]byte, validatorSet ValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.contract.Transact(opts, "setValidatorSet", avalancheBlockchainID, validatorSet)
}

// SetValidatorSet is a paid mutator transaction binding the contract method 0xe8a6c940.
//
// Solidity: function setValidatorSet(bytes32 avalancheBlockchainID, (bytes32,(bytes,uint64)[],uint64,uint64,uint64) validatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerSession) SetValidatorSet(avalancheBlockchainID [32]byte, validatorSet ValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.SetValidatorSet(&_AvalancheValidatorSetManager.TransactOpts, avalancheBlockchainID, validatorSet)
}

// SetValidatorSet is a paid mutator transaction binding the contract method 0xe8a6c940.
//
// Solidity: function setValidatorSet(bytes32 avalancheBlockchainID, (bytes32,(bytes,uint64)[],uint64,uint64,uint64) validatorSet) returns()
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerTransactorSession) SetValidatorSet(avalancheBlockchainID [32]byte, validatorSet ValidatorSet) (*types.Transaction, error) {
	return _AvalancheValidatorSetManager.Contract.SetValidatorSet(&_AvalancheValidatorSetManager.TransactOpts, avalancheBlockchainID, validatorSet)
}

// AvalancheValidatorSetManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AvalancheValidatorSetManager contract.
type AvalancheValidatorSetManagerInitializedIterator struct {
	Event *AvalancheValidatorSetManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AvalancheValidatorSetManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalancheValidatorSetManagerInitialized)
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
		it.Event = new(AvalancheValidatorSetManagerInitialized)
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
func (it *AvalancheValidatorSetManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalancheValidatorSetManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalancheValidatorSetManagerInitialized represents a Initialized event raised by the AvalancheValidatorSetManager contract.
type AvalancheValidatorSetManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvalancheValidatorSetManagerInitializedIterator, error) {

	logs, sub, err := _AvalancheValidatorSetManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvalancheValidatorSetManagerInitializedIterator{contract: _AvalancheValidatorSetManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvalancheValidatorSetManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AvalancheValidatorSetManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalancheValidatorSetManagerInitialized)
				if err := _AvalancheValidatorSetManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AvalancheValidatorSetManager *AvalancheValidatorSetManagerFilterer) ParseInitialized(log types.Log) (*AvalancheValidatorSetManagerInitialized, error) {
	event := new(AvalancheValidatorSetManagerInitialized)
	if err := _AvalancheValidatorSetManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
