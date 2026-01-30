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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getShardHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getShardsReceived\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"shardsReceived\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"numValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"partialWeight\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"inProgress\",\"type\":\"bool\"}],\"internalType\":\"structPartialValidatorSet\",\"name\":\"partialValidatorSet\",\"type\":\"tuple\"}],\"name\":\"setPartialValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"validatorSet\",\"type\":\"tuple\"}],\"name\":\"setValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576117ce908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816308e64d7414610ac657508063141eaf4a14610a8557806327258b2214610a4e57806349e4db9c146108085780638457eaa7146107d3578063933568401461050a578063c49abdfa1461021f578063e8a6c940146100b95763eb3b597d1461007f575f80fd5b346100b55760203660031901126100b5576004355f52600160205260206001600160401b03600260405f20015416604051908152f35b5f80fd5b346100b55760403660031901126100b5576024356001600160401b0381116100b55760a060031982360301126100b557604051906100f682611027565b8060040135825260248101356001600160401b0381116100b557610120906004369184010161110a565b9060208301918252610134604482016110df565b926040810193845261015b608461014d606485016110df565b9360608401948552016110df565b92608082019384526004355f525f60205260405f209151825560018201905190602082519261018a8484611574565b01905f5260205f205f915b838310610201578751600286018054885167ffffffffffffffff60401b60409190911b166001600160401b039093166001600160801b0319909116179190911781558751815467ffffffffffffffff60801b191660809190911b67ffffffffffffffff60801b16179055005b60026020826102136001945186611473565b01920192019190610195565b346100b55760403660031901126100b5576024356001600160401b0381116100b55761010060031982360301126100b55760405161010081018181106001600160401b038211176104e657604052610279826004016110df565b8152610287602483016110df565b906020810191825260448301356001600160401b0381116100b5578301366023820112156100b5576004810135906102be826110f3565b916102cc604051938461105d565b808352602060048185019260051b84010101913683116100b557602401905b8282106104fa5750505060408201908152610308606485016110df565b6060830190815260848501356001600160401b0381116100b557610332906004369188010161110a565b906080840191825260a084019260a4870135845260e461035460c489016110df565b9760c0870198895201359586151587036100b55760e086019687526004355f90815260016020526040908190209651875492516001600160801b03199093166001600160401b03919091161791901b67ffffffffffffffff60401b161785555180519060018601906001600160401b0383116104e657600160401b83116104e65781548383558084106104c0575b50602001905f5260205f205f5b8381106104ac57505050506001600160401b039051166001600160401b036002850191166001600160401b03198254161790556003830190519060208251926104388484611574565b01905f5260205f205f915b83831061048e57845160048701558751600587018054895168ff000000000000000090151560401b166001600160401b0390931668ffffffffffffffffff1990911617919091179055005b60026020826104a06001945186611473565b01920192019190610443565b6001906020845194019381840155016103ef565b825f528360205f2091820191015b8181106104db57506103e2565b5f81556001016104ce565b634e487b7160e01b5f52604160045260245ffd5b81358152602091820191016102eb565b346100b5573660031901606081126100b5576040136100b5576044356001600160401b0381116100b557610542903690600401611099565b61056c5f602435926040518093819263b9a1525960e01b8352602060048401526024830190611003565b038173__$aaf4ae346b84a712cc43f25bb66199d6fb$__5af480156107c8575f915f9161079b575b506001600160401b03906105aa8351151561132f565b16906105b782151561137b565b825f526001602052600460405f200154905f5b815181101561063a5781518110156106265760208160051b8301015190855f526001602052600360405f2001610600828661140b565b92815484101561062657600193610620925f5260205f2090851b01611473565b016105ca565b634e487b7160e01b5f52603260045260245ffd5b50905051825f526001602052610658600460405f200191825461140b565b9055815f526001602052600560405f2001906001600160401b03825416016001600160401b038111610787576001600160401b03166001600160401b0319825416179055805f526001602052600260405f200160016001600160401b03825416016001600160401b038111610787576001600160401b03166001600160401b03198254161790556004356001600160401b0381168091036100b557815f526001602052600160405f2001541461070a57005b5f81815260016020818152604080842060058101805468ff00000000000000001916905591849052909220610743926003019101611711565b805f5260016020526001600160401b03600560405f20015416905f525f6020526001600160401b03600260405f200191166001600160401b03198254161790555f80f35b634e487b7160e01b5f52601160045260245ffd5b6001600160401b0392506107c191503d805f833e6107b9818361105d565b810190611208565b9091610594565b6040513d5f823e3d90fd5b346100b55760203660031901126100b5576004355f526001602052602060ff600560405f20015460401c166040519015158152f35b346100b55760203660031901126100b5575f608060405161082881611027565b8281526060602082015282604082015282606082015201526004355f525f60205260405f206040519061085a82611027565b80548252600181019081549161086f836110f3565b9261087d604051948561105d565b80845260208401915f5260205f205f925b82841061098a57868660028760208401928352015491604081016001600160401b038416815260608201906001600160401b038560401c1682526001600160401b03608084019560801c168552604051936020855260c0850193516020860152519260a06040860152835180915260e0850190602060e08260051b8801019501915f905b828210610942578780886001600160401b038c818b818c5116606087015251166080850152511660a08301520390f35b9091929560208060019260df198b8203018552895190826001600160401b03816109758551604086526040860190611003565b94015116910152980192019201909291610912565b60405161099681611042565b6040515f84546109a5816113d3565b8084529060018116908115610a2b57506001146109f4575b5092600292826109d3602094600197038261105d565b81526001600160401b0385870154168382015281520192019301929061088e565b5f868152602081209092505b818310610a15575050810160200160026109bd565b6001816020925483868801015201920191610a00565b60ff191660208086019190915291151560051b84019091019150600290506109bd565b346100b55760203660031901126100b5576004355f525f60205260206001600160401b03600260405f200154161515604051908152f35b346100b55760403660031901126100b5576024356004355f526001602052600160405f2001908154811015610626576020915f52815f200154604051908152f35b346100b55760403660031901126100b557600435906001600160401b0382116100b5578136039160606003198401126100b557602435926001600160401b0384116100b557366023850112156100b5578360040135926001600160401b0384116100b557602485019460248536920101116100b557608081610b49606093611027565b5f81525f60208201525f60408201525f83820152015273__$aaf4ae346b84a712cc43f25bb66199d6fb$__9180600401359160a219018212156100b55701608481013560221936839003018112156100b55701600401803593906001600160401b0385116100b55760200184360381136100b55760405163b70e3f0360e01b8152945f9186918291610bdf9190600484016111e1565b0381855af49384156107c8575f94610f08575b5060808401928351805115610626576020015160205f6040518486823780858101838152039060025afa156107c8575f5103610ec35760405163b9a1525960e01b8152925f92849283918291610c4b91600484016111e1565b03915af49182156107c8575f915f93610ea5575b50835190815f525f6020526001600160401b03600260405f20015460401c169160208601926001600160401b038451161115610e3a575f525f6020526001600160401b03600260405f20015460801c169360408601946001600160401b038651161115610dd05760606001600160401b038092610cde8751151561132f565b1696610ceb88151561137b565b8260405197838952816101008a01978451868c0152511660808a0152511660a088015201511660c0850152519060a060e08501528151809152602061012085019201905f5b818110610dba575050508281036020840152815180825260208201916020808360051b8301019401925f915b838310610d70578680878a60408301520390f35b9091929394602080600192601f19858203018652885190826001600160401b0381610da48551604086526040860190611003565b9401511691015297019301930191939290610d5c565b8251845260209384019390920191600101610d30565b608460405162461bcd60e51b815260206004820152604060248201527f502d436861696e2074696d657374616d70206d7573742062652067726561746560448201527f72207468616e207468652063757272656e742076616c696461746f72207365746064820152fd5b60405162461bcd60e51b815260206004820152603d60248201527f502d436861696e20686569676874206d7573742062652067726561746572207460448201527f68616e207468652063757272656e742076616c696461746f72207365740000006064820152608490fd5b909250610ebc91503d805f833e6107b9818361105d565b9184610c5f565b60405162461bcd60e51b815260206004820152601b60248201527f56616c696461746f72207365742068617368206d69736d6174636800000000006044820152606490fd5b9093503d805f833e610f1a818361105d565b8101906020818303126100b5578051906001600160401b0382116100b557019060a0828203126100b55760405191610f5183611027565b80518352610f61602082016111cd565b6020840152610f72604082016111cd565b6040840152610f83606082016111cd565b60608401526080810151906001600160401b0382116100b557019080601f830112156100b5578151610fb4816110f3565b92610fc2604051948561105d565b81845260208085019260051b8201019283116100b557602001905b828210610ff35750505060808201529284610bf2565b8151815260209182019101610fdd565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60a081019081106001600160401b038211176104e657604052565b604081019081106001600160401b038211176104e657604052565b90601f801991011681019081106001600160401b038211176104e657604052565b6001600160401b0381116104e657601f01601f191660200190565b81601f820112156100b5578035906110b08261107e565b926110be604051948561105d565b828452602083830101116100b557815f926020809301838601378301015290565b35906001600160401b03821682036100b557565b6001600160401b0381116104e65760051b60200190565b9080601f830112156100b557813591611122836110f3565b92611130604051948561105d565b80845260208085019160051b830101918383116100b55760208101915b83831061115c57505050505090565b82356001600160401b0381116100b5578201906040828703601f1901126100b5576040519061118a82611042565b6020830135916001600160401b0383116100b5576111be6040856111b68b602080999881990101611099565b8452016110df565b8382015281520192019161114d565b51906001600160401b03821682036100b557565b90918060409360208452816020850152848401375f828201840152601f01601f1916010190565b91906040838203126100b55782516001600160401b0381116100b557830181601f820112156100b55780519061123d826110f3565b9261124b604051948561105d565b82845260208085019360051b830101918183116100b55760208101935b83851061128457505050505060206112819193016111cd565b90565b84516001600160401b0381116100b5578201906040828503601f1901126100b557604051906112b282611042565b60208301516001600160401b0381116100b5576020908401019185601f840112156100b55782516112e28161107e565b936112f0604051958661105d565b81855287602083830101116100b55760209586955f87856113209682604097018386015e830101528452016111cd565b83820152815201940193611268565b1561133657565b60405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f72207365742063616e6e6f7420626520656d7074790000006044820152606490fd5b1561138257565b60405162461bcd60e51b815260206004820152602360248201527f546f74616c20776569676874206d75737420626520677265617465722074686160448201526206e20360ec1b6064820152608490fd5b90600182811c92168015611401575b60208310146113ed57565b634e487b7160e01b5f52602260045260245ffd5b91607f16916113e2565b9190820180921161078757565b818110611423575050565b5f8155600101611418565b9190601f811161143d57505050565b611467925f5260205f20906020601f840160051c83019310611469575b601f0160051c0190611418565b565b909150819061145a565b91909182519283516001600160401b0381116104e65761149d8161149785546113d3565b8561142e565b6020601f8211600114611503576001926114dd836001600160401b039697989460209488965f926114f8575b50508160011b915f199060031b1c19161790565b86555b015116920191166001600160401b0319825416179055565b015190505f806114c9565b601f19821695845f52815f20965f5b81811061155c5750836001600160401b03969798602094600197948997899510611544575b505050811b0186556114e0565b01515f1960f88460031b161c191690555f8080611537565b83830151895560019098019760209384019301611512565b600160401b82116104e65780549082815581831061159157505050565b6001600160ff1b0382168203610787576001600160ff1b0383168303610787575f5260205f209060011b81019160011b015b8181106115ce575050565b806115db600292546113d3565b806115ee575b505f6001820155016115c3565b601f811160011461160457505f81555b5f6115e1565b61162190825f526001601f60205f20920160051c82019101611418565b805f525f60208120818355556115fe565b91909182811461170c5761164683546113d3565b6001600160401b0381116104e6576116688161166284546113d3565b8461142e565b5f93601f82116001146116a65761169792939482915f9261169b5750508160011b915f199060031b1c19161790565b9055565b015490505f806114c9565b601f198216905f5260205f2094835f5260205f20915f5b8181106116f4575095836001959697106116dc575b505050811b019055565b01545f1960f88460031b161c191690555f80806116d2565b9192600180602092868b0154815501940192016116bd565b509050565b818114611794578154916117258383611574565b5f5260205f20905f5260205f205f915b8383106117425750505050565b60028082600193850361175c575b01920192019190611735565b6117668186611632565b6001600160401b0384820154166001600160401b0385870191166001600160401b0319825416179055611750565b505056fea264697066735822122091713ceb41392e0f016043e28a9c7e09d34db85ee5d4431405c263a2a63ff3c264736f6c634300081e0033",
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
