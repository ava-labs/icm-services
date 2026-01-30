// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avalanchevalidatorsetregistry

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

// AvalancheValidatorSetRegistryMetaData contains all meta data concerning the AvalancheValidatorSetRegistry contract.
var AvalancheValidatorSetRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"validatorSetUpdaterContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSetManagerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e08060405234610465575f611f5b803803809161001d82866104a8565b843982019160608184031261046557805163ffffffff811681036104655760208201516001600160401b0381116104655782019160a08386031261046557604051946100688661048d565b83518652610078602085016104cb565b926020870193845261008c604086016104cb565b92604088019384526100a0606087016104cb565b60608901908152608087015190966001600160401b03821161046557019280601f850112156104655783516100d4816104df565b946100e260405196876104a8565b81865260208087019260051b82010192831161046557602001905b82821061047d575050506080880192835260400151906001600160a01b038216820361046557608052865160a05260c05282518251915185516001600160401b0393841693919282169161015191166104f6565b6040519161010083016001600160401b03811184821017610469576040528252602082019384526040820192835260608201905f82526080830190815260a08301925f845260c08101915f835260e08201936001855260018060a01b0360c051169660a05198883b15610465576040805163624d5efd60e11b8152600481019b909b5260248b015293516001600160401b0390811660448b0152935190931660648901529151610100608489015280516101448901819052939688969095909490936101648801939290916020909101905f5b81811061044c57505090516001600160401b031660a488015250518582036043190160c48701525f96869488948694919390929161026191610563565b925160e4850152516001600160401b031661010484015251151561012483015203925af1801561044157610427575b50935190519351915192936001600160401b03928316938593918216916102b791166104f6565b91604051906102c58261048d565b81526020810192835260408101848152606082019283526080820195865260c05160a0516001600160a01b039091169391843b156104235761033b978793604051998a98899788966303a29b2560e61b88526004880152604060248801525160448701525160a0606487015260e4860190610563565b92516001600160401b0390811660848601529051811660a485015290511660c483015203925af1801561041857610400575b60405161196e90816105ed82396080518181816101160152610f20015260a0518181816101fd01528181610dfb01526116e4015260c051818181610360015281816104900152818161056b015281816106320152818161068501528181610782015281816107d501528181610a9501528181610db5015281816112140152818161167e0152818161171001526117680152f35b61040b8280926104a8565b610415578061036d565b80fd5b6040513d84823e3d90fd5b8680fd5b6102b7959394505f610438916104a8565b5f939294610290565b6040513d5f823e3d90fd5b825186528b995060209586019590920191600101610224565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b81518152602091820191016100fd565b60a081019081106001600160401b0382111761046957604052565b601f909101601f19168101906001600160401b0382119082101761046957604052565b51906001600160401b038216820361046557565b6001600160401b0381116104695760051b60200190565b90610500826104df565b61050d60405191826104a8565b828152809261051e601f19916104df565b015f5b81811061052d57505050565b6040805190810191906001600160401b0383118184101761046957602092604052606081525f8382015282828601015201610521565b9080602083519182815201916020808360051b8301019401925f915b83831061058e57505050505090565b90919293946020806060600193601f19868203018752828a518051604084528051928391826040870152018585015e5f838301850152840151604087901b8790031684830152601f01601f191601019701959490910192019061057f56fe60806040526004361015610011575f80fd5b5f5f3560e01c80631c0a98ae14610e3c57806327258b2214610e1e578063541dcba414610de457806357f1617714610da0578063580d632b14610d865780636766233d146109f857806368531ed0146109f357806382366d05146109f35780638457eaa7146109ca5763e82d170914610088575f80fd5b3461095c57604036600319011261095c57600435906001600160401b03821161095c578136039160606003198401126106d6576024356001600160401b03811161094257366023820112156109425780600401356001600160401b03811161095a57366024828401011161095a5761013e61010e6101096004860180610f44565b610f59565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614610f6a565b61015760206101506004860180610f44565b013561174b565b156101db5760405162461bcd60e51b815260206004820152604b60248201527f43616e277420726567697374657220746f206120626c6f636b636861696e204960448201527f44207768696c6520616e6f7468657220726567697374726174696f6e2069732060648201526a696e2070726f677265737360a81b608482015260a490fd5b6101f460206101ed6004860180610f44565b0135611661565b6109a8576102257f00000000000000000000000000000000000000000000000000000000000000008460040161116e565b60406102346004850180610f44565b01356001600160a01b038116908190036106fa576109635760405194630239935d60e21b86526040600487015283600401359060a219018112156106fa578301600481019060606044880152813563ffffffff811680910361095f57938793602461034a6103286103076102f48998608461035c998f9d60a48d01528781013560c48d015260018060a01b036102cc604483016117a3565b1660e48d01526001600160a01b036102e6606483016117a3565b166101048d015201906117b7565b60a06101248b01526101448a0191611024565b610316848c018c6004016117b7565b8983036043190160648b015290611024565b61033860448b018b6004016117b7565b8883036043190160848a015290611024565b85810360031901828701529201611024565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156106da578280948192610826575b50835160808501946001865151115f1461073b57606081016103c56001600160401b038251166118b4565b95845b89518110156103fb57806103de6001928c611910565b516103e9828b611910565b526103f4818a611910565b50016103c8565b5091949792959690939660208601936001600160401b03855116996001600160401b0360408901511693519151926040519b6101008d018d81106001600160401b03821117610727576040528c5260208c0194855260408c019283528b6001600160401b0360c06060830192600184526080810194855260a08101978852019316835260e08d01936001855260018060a01b037f0000000000000000000000000000000000000000000000000000000000000000163b15610723576040519663624d5efd60e11b88528c6004890152602488016040905261014488019e516001600160401b03166044890152516001600160401b03166064880152519c6084870161010090528d5180915261016487019d602001908d5b81811061070957505050926001600160401b0386949361054e8e9f8f9e9f9895848998511660a4890152516043198883030160c48901526117e8565b935160e486015251166101048401525115156101248301520381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19081156106fe5786916106e5575b506105b991602091506101ed9060040180610f44565b156105eb575b505050505b7f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce938280a280f35b6001600160401b0380604081610609945116950151169251166118b4565b916040519261061784610e7e565b858452602084015260408301849052606083015260808201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163b156106d6578161068091604051809381926303a29b2560e61b8352876004840161185d565b0381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156106da576106c1575b80806105bf565b816106cb91610ec8565b6106d657815f6106ba565b5080fd5b6040513d84823e3d90fd5b816106ef91610ec8565b6106fa57845f6105a3565b8480fd5b6040513d88823e3d90fd5b90919e8f5181526020019e60200190600101919091610512565b8c80fd5b634e487b7160e01b8d52604160045260248dfd5b9094508195929193506001600160401b038060408160208501511693015116926040519461076886610e7e565b8786526020860152166040840152606083015260808201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163b156106d657816107d091604051809381926303a29b2560e61b8352876004840161185d565b0381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156106da57610811575b50506105c4565b8161081b91610ec8565b6106d657815f61080a565b93509350503d8084843e61083a8184610ec8565b82019260608385031261095c5782516001600160401b0381116106d65783019360a0858203126106d6576040519461087186610e7e565b805186526108816020820161105b565b60208701526108926040820161105b565b60408701526108a36060820161105b565b60608701526080810151906001600160401b03821161095a570181601f82011215610942578051906108d482611044565b916108e26040519384610ec8565b80835260208084019160051b8301019184831161095657602001905b82821061094657505050608086015260208401516001600160401b0381116109425761093161093892604092870161106f565b940161105b565b939293905f61039a565b8280fd5b81518152602091820191016108fe565b8580fd5b835b80fd5b8680fd5b60405162461bcd60e51b815260206004820152601c60248201527f536f757263652061646472657373206d75737420626520656d707479000000006044820152606490fd5b6109c560206109ba6004860180610f44565b01358460040161116e565b610225565b503461095c57602036600319011261095c5760206109e960043561174b565b6040519015158152f35b610f04565b5034610c1557366003190160608112610c1557604013610c15576044356001600160401b038111610c155736602382011215610c15578060040135610a3c81610ee9565b90610a4a6040519283610ec8565b80825260208201923660248383010111610c1557815f9260246020930186378301015260243591610a7a8361174b565b15610d235760405163eb3b597d60e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169290602081602481875afa8015610c0a575f90610ce3575b6001600160401b0391501660018101926001600160401b038411610ccf57600435936001600160401b038516809503610c15576001600160401b0385911603610c8a5760405191630a0f57a560e11b83528660048401526024830152602082604481885afa918215610c0a575f92610c56575b505f602091604051809186518091835e8101838152039060025afa15610c0a575f5103610c1957823b15610c1557610bac925f928360405180968195829463024cd5a160e61b8452600484015289602484015260606044840152606483019061114a565b03925af18015610c0a57610bf5575b50610bc58161174b565b15610bce575080f35b7f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe28280a280f35b610c029192505f90610ec8565b5f905f610bbb565b6040513d5f823e3d90fd5b5f80fd5b60405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b6044820152606490fd5b9091506020813d602011610c82575b81610c7260209383610ec8565b81010312610c155751905f610b48565b3d9150610c65565b60405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f7264657200000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b506020813d602011610d1b575b81610cfd60209383610ec8565b81010312610c1557610d166001600160401b039161105b565b610ad5565b3d9150610cf0565b60405162461bcd60e51b815260206004820152603560248201527f43616e6e6f74206170706c7920736861726420696620726567697374726174696044820152746f6e206973206e6f7420696e2070726f677265737360581b6064820152608490fd5b34610c15575f366003190112610c155760206109e96116d5565b34610c15575f366003190112610c15576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610c15575f366003190112610c155760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34610c15576020366003190112610c155760206109e9600435611661565b34610c15576040366003190112610c15576004356001600160401b038111610c155760606003198236030112610c1557610e7c906024359060040161116e565b005b60a081019081106001600160401b03821117610e9957604052565b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b03821117610e9957604052565b90601f801991011681019081106001600160401b03821117610e9957604052565b6001600160401b038111610e9957601f01601f191660200190565b34610c15575f366003190112610c1557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b903590609e1981360301821215610c15570190565b3563ffffffff81168103610c155790565b15610f7157565b60405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606490fd5b903590601e1981360301821215610c1557018035906001600160401b038211610c1557602001918136038313610c1557565b81601f82011215610c1557805190610ff582610ee9565b926110036040519485610ec8565b82845260208383010111610c1557815f9260208093018386015e8301015290565b908060209392818452848401375f828201840152601f01601f1916010190565b6001600160401b038111610e995760051b60200190565b51906001600160401b0382168203610c1557565b9080601f83011215610c155781519161108783611044565b926110956040519485610ec8565b80845260208085019160051b83010191838311610c155760208101915b8383106110c157505050505090565b82516001600160401b038111610c15578201906040828703601f190112610c1557604051906110ef82610ead565b6020830151916001600160401b038311610c155761112360408561111b8b602080999881990101610fde565b84520161105b565b838201528152019201916110b2565b90816020910312610c1557518015158103610c155790565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b6111766116d5565b156115e75761118482611661565b1561156c576111e49161119d61010e6101098480610f44565b73__$aaf4ae346b84a712cc43f25bb66199d6fb$__915f6111c16040830183610fac565b6040516310b15a7360e31b81526020600482015296879283926024840191611024565b0381865af4938415610c0a575f946114dd575b5060405163127936e760e21b815260048101929092525f826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610c0a575f9261142d575b50906112b86112638360206112a697950190610fac565b604051968795630161c9f960e61b87526060600488015260206112928251604060648b015260a48a019061114a565b91015187820360631901608489015261114a565b85810360031901602487015291611024565b60031983820301604484015260a08101918051825260208101519260a06020840152835180915260c0830190602060c08260051b8601019501915f905b8282106113dd575050505060209492849260806001600160401b0381858260408998015116604086015282606082015116606086015201511691015203915af4908115610c0a575f916113ae575b501561134b57565b60405162461bcd60e51b815260206004820152603560248201527f436f756c64206e6f74207665726966792049434d206d6573736167653a20536960448201527419db985d1d5c994818da1958dadcc819985a5b1959605a1b6064820152608490fd5b6113d0915060203d6020116113d6575b6113c88183610ec8565b810190611132565b5f611343565b503d6113be565b919394956001919397506020809160bf19898203018552895190826001600160401b0381611414855160408652604086019061114a565b94015116910152980192019201889695949391926112f5565b9091503d805f833e61143f8183610ec8565b810190602081830312610c15578051906001600160401b038211610c15570160a081830312610c155760405161147481610e7e565b8151815260208201516001600160401b038111610c15576080836114a36112b8966114d094611263970161106f565b60208501526114b46040820161105b565b60408501526114c56060820161105b565b60608501520161105b565b608082015292915061124c565b9093503d805f833e6114ef8183610ec8565b810190602081830312610c15578051906001600160401b038211610c155701604081830312610c15576040519161152583610ead565b81516001600160401b038111610c155781611541918401610fde565b835260208201516001600160401b038111610c15576115609201610fde565b6020820152925f6111f7565b60405162461bcd60e51b815260206004820152604760248201527f4e6f2076616c696461746f72207365742069732072656769737465726564206660448201527f6f72207468652070726f7669646564204176616c616e63686520626c6f636b636064820152661a185a5b88125160ca1b608482015260a490fd5b60405162461bcd60e51b815260206004820152604660248201527f4120636f6d706c65746520502d636861696e2076616c696461746f72206d757360448201527f74206265207265676973746572656420746f207665726966792049434d206d6560648201526573736167657360d01b608482015260a490fd5b604051631392c59160e11b815260048101919091526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c0a575f916116b9575090565b6116d2915060203d6020116113d6576113c88183610ec8565b90565b604051631392c59160e11b81527f000000000000000000000000000000000000000000000000000000000000000060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c0a575f916116b9575090565b604051638457eaa760e01b815260048101919091526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c0a575f916116b9575090565b35906001600160a01b0382168203610c1557565b9035601e1982360301811215610c155701602081359101916001600160401b038211610c15578136038313610c1557565b9080602083519182815201916020808360051b8301019401925f915b83831061181357505050505090565b9091929394602080600192601f19858203018652885190826001600160401b0381611847855160408652604086019061114a565b9401511691015297019301930191939290611804565b908152604060208201528151604082015260c06001600160401b036080611893602086015160a0606087015260e08601906117e8565b9482604082015116828601528260608201511660a086015201511691015290565b906118be82611044565b6118cb6040519182610ec8565b82815280926118dc601f1991611044565b01905f5b8281106118ec57505050565b6020906040516118fb81610ead565b606081525f83820152828285010152016118e0565b80518210156119245760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfea2646970667358221220420427085b1dd6673b2379cb14f8ef9af69d6fd3314dc689a5ce8f8f298d712864736f6c634300081e0033",
}

// AvalancheValidatorSetRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.ABI instead.
var AvalancheValidatorSetRegistryABI = AvalancheValidatorSetRegistryMetaData.ABI

// AvalancheValidatorSetRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AvalancheValidatorSetRegistryMetaData.Bin instead.
var AvalancheValidatorSetRegistryBin = AvalancheValidatorSetRegistryMetaData.Bin

// DeployAvalancheValidatorSetRegistry deploys a new Ethereum contract, binding an instance of AvalancheValidatorSetRegistry to it.
func DeployAvalancheValidatorSetRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, avalancheNetworkID_ uint32, initialValidatorSetData ValidatorSetMetadata, validatorSetUpdaterContract_ common.Address) (common.Address, *types.Transaction, *AvalancheValidatorSetRegistry, error) {
	parsed, err := AvalancheValidatorSetRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AvalancheValidatorSetRegistryBin), backend, avalancheNetworkID_, initialValidatorSetData, validatorSetUpdaterContract_)
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

// ValidatorSetManagerContract is a free data retrieval call binding the contract method 0x57f16177.
//
// Solidity: function validatorSetManagerContract() view returns(address)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) ValidatorSetManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "validatorSetManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSetManagerContract is a free data retrieval call binding the contract method 0x57f16177.
//
// Solidity: function validatorSetManagerContract() view returns(address)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) ValidatorSetManagerContract() (common.Address, error) {
	return _AvalancheValidatorSetRegistry.Contract.ValidatorSetManagerContract(&_AvalancheValidatorSetRegistry.CallOpts)
}

// ValidatorSetManagerContract is a free data retrieval call binding the contract method 0x57f16177.
//
// Solidity: function validatorSetManagerContract() view returns(address)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) ValidatorSetManagerContract() (common.Address, error) {
	return _AvalancheValidatorSetRegistry.Contract.ValidatorSetManagerContract(&_AvalancheValidatorSetRegistry.CallOpts)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x1c0a98ae.
//
// Solidity: function verifyICMMessage(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) VerifyICMMessage(opts *bind.CallOpts, message ICMMessage, avalancheBlockchainID [32]byte) error {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "verifyICMMessage", message, avalancheBlockchainID)

	if err != nil {
		return err
	}

	return err

}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x1c0a98ae.
//
// Solidity: function verifyICMMessage(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _AvalancheValidatorSetRegistry.Contract.VerifyICMMessage(&_AvalancheValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// VerifyICMMessage is a free data retrieval call binding the contract method 0x1c0a98ae.
//
// Solidity: function verifyICMMessage(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes32 avalancheBlockchainID) view returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) VerifyICMMessage(message ICMMessage, avalancheBlockchainID [32]byte) error {
	return _AvalancheValidatorSetRegistry.Contract.VerifyICMMessage(&_AvalancheValidatorSetRegistry.CallOpts, message, avalancheBlockchainID)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0xe82d1709.
//
// Solidity: function registerValidatorSet(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryTransactor) RegisterValidatorSet(opts *bind.TransactOpts, message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.contract.Transact(opts, "registerValidatorSet", message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0xe82d1709.
//
// Solidity: function registerValidatorSet(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes shardBytes) returns()
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) RegisterValidatorSet(message ICMMessage, shardBytes []byte) (*types.Transaction, error) {
	return _AvalancheValidatorSetRegistry.Contract.RegisterValidatorSet(&_AvalancheValidatorSetRegistry.TransactOpts, message, shardBytes)
}

// RegisterValidatorSet is a paid mutator transaction binding the contract method 0xe82d1709.
//
// Solidity: function registerValidatorSet(((uint32,bytes32,address,address,bytes),bytes,bytes) message, bytes shardBytes) returns()
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
