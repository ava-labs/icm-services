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
	ValidatorSetHash      [32]byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"totalValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"validatorSetUpdaterContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSetManagerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMRawMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rawMessageBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e080604052346104b5575f611fd9803803809161001d82866104dd565b84398201916060818403126104b557805163ffffffff811681036104b55760208201516001600160401b0381116104b55782019160c0838603126104b5576040519060c082016001600160401b038111838210176104b9576040528351825261008860208501610500565b956020830196875261009c60408601610500565b9360408401948552606086015192606085019384526100bd60808801610500565b6080860190815260a088015190976001600160401b0382116104b557019280601f850112156104b55783516100f181610514565b946100ff60405196876104dd565b81865260208087019260051b8201019283116104b557602001905b8282106104cd5750505060a0850192835260400151906001600160a01b03821682036104b557608052835160a05260c052865184519251915186516001600160401b03948516949283169261016f911661052b565b604051909261012082016001600160401b038111838210176104b957604052815260208101948552604081019384526060810191825260808101905f825260a0810193845260c08101935f855260e08201925f84526101008301946001865260018060a01b0360c051169760a05199893b156104b55760408051632bd41b5360e01b8152600481019c909c5260248c015294516001600160401b0390811660448c0152945190941660648a015292516084890152915161012060a489015280516101648901819052939688969095909490936101848801939290916020909101905f5b81811061049c57505090516001600160401b031660c488015250518582036043190160e48701525f96869488948694919390929161028f91610598565b9251610104850152516001600160401b031661012484015251151561014483015203925af180156104915761047a575b50519351905191516001600160401b0392831692918216916102e1911661052b565b60405160a08101959192916001600160401b038711828810176104665785966040969596528152602081019283526040810191848352606082019081526080820195865260018060a01b0360c051169260a05190843b156104625761037a978793604051998a98899788966303a29b2560e61b88526004880152604060248801525160448701525160a0606487015260e4860190610598565b92516001600160401b0390811660848601529051811660a485015290511660c483015203925af180156104575761043f575b6040516119b7908161062282396080518181816101150152610f69015260a0518181816101fc01528181610e44015261172d015260c05181818161035f0152818161049d0152818161057e015281816106450152818161069801528181610791015281816107e401528181610ade01528181610dfe0152818161125d015281816116c70152818161175901526117b10152f35b61044a8280926104dd565b61045457806103ac565b80fd5b6040513d84823e3d90fd5b8680fd5b634e487b7160e01b86526041600452602486fd5b6104879194505f906104dd565b5f926102e16102bf565b6040513d5f823e3d90fd5b825186528b995060209586019590920191600101610252565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b815181526020918201910161011a565b601f909101601f19168101906001600160401b038211908210176104b957604052565b51906001600160401b03821682036104b557565b6001600160401b0381116104b95760051b60200190565b9061053582610514565b61054260405191826104dd565b8281528092610553601f1991610514565b015f5b81811061056257505050565b6040805190810191906001600160401b038311818410176104b957602092604052606081525f8382015282828601015201610556565b9080602083519182815201916020808360051b8301019401925f915b8383106105c357505050505090565b90919293946020806060600193601f19868203018752828a518051604084528051928391826040870152018585015e5f838301850152840151604087901b8790031684830152601f01601f19160101970195949091019201906105b456fe60806040526004361015610011575f80fd5b5f5f3560e01c80631c0a98ae14610e8557806327258b2214610e67578063541dcba414610e2d57806357f1617714610de9578063580d632b14610dcf5780636766233d14610a4157806368531ed014610a3c57806382366d0514610a3c5780638457eaa714610a135763e82d170914610088575f80fd5b3461099957604036600319011261099957600435906001600160401b0382116109995781360360606003198201126106e9576024356001600160401b03811161096957366023820112156109695780600401356001600160401b0381116109815736602482840101116109815761013d61010d6101086004880180610f8d565b610fa2565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614610fb3565b610156602061014f6004880180610f8d565b0135611794565b156101da5760405162461bcd60e51b815260206004820152604b60248201527f43616e277420726567697374657220746f206120626c6f636b636861696e204960448201527f44207768696c6520616e6f7468657220726567697374726174696f6e2069732060648201526a696e2070726f677265737360a81b608482015260a490fd5b6101f360206101ec6004880180610f8d565b01356116aa565b6109f1576102247f0000000000000000000000000000000000000000000000000000000000000000866004016111b7565b60406102336004870180610f8d565b01356001600160a01b0381169081900361070d576109ac5760405192630239935d60e21b84526040600485015285600401359060a2190181121561070d578501600481019060606044860152813563ffffffff81168091036109a85793859360246103496103276103066102f38998608461035b998f9d60a48d01528781013560c48d015260018060a01b036102cb604483016117ec565b1660e48d01526001600160a01b036102e5606483016117ec565b166101048d01520190611800565b60a06101248b01526101448a019161106d565b610315848e0160048f01611800565b8983036043190160648b01529061106d565b61033760448d018d600401611800565b8883036043190160848a01529061106d565b8581036003190182870152920161106d565b03817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa91821561099c578081928294610835575b5080519060a08101946001865151115f1461074a5760808201906103c76001600160401b038351166118fd565b97855b87518110156103fd57806103e06001928a611959565b516103eb828d611959565b526103f6818c611959565b50016103ca565b50919395969092949660208601936001600160401b03855116906001600160401b036040890151169a606089015194519351936040519361012085018581106001600160401b03821117610736576040528452602084019c8d52604084019586526060840190815260808401906001825260a0850192835260c085019586526001600160401b0360e0860194168452600161010086015260018060a01b037f0000000000000000000000000000000000000000000000000000000000000000163b156107325760409d939d5196632bd41b5360e01b88528c6004890152604060248901526001600160401b036101648901958188511660448b015251166064890152516084880152519261012060a48801528351809152602061018488019401908d5b81811061071c5750505085936001600160401b038d9e61055f8f9e9f98959688978461010097511660c48a0152516043198983030160e48a0152611831565b94516101048701525116610124850152015115156101448301520381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af19081156107115786916106f8575b506105cc91602091506101ec9060040180610f8d565b156105fe575b505050505b7f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce938280a280f35b6001600160401b038060408161061c945116950151169251166118fd565b916040519261062a84610ec7565b858452602084015260408301849052606083015260808201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163b156106e9578161069391604051809381926303a29b2560e61b835287600484016118a6565b0381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156106ed576106d4575b80806105d2565b816106de91610f11565b6106e957815f6106cd565b5080fd5b6040513d84823e3d90fd5b8161070291610f11565b61070d57845f6105b6565b8480fd5b6040513d88823e3d90fd5b8251865260209586019590920191600101610520565b8c80fd5b634e487b7160e01b8e52604160045260248efd5b9194508295506001600160401b038060408160208598979801511693015116926040519461077786610ec7565b8786526020860152166040840152606083015260808201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163b156106e957816107df91604051809381926303a29b2560e61b835287600484016118a6565b0381837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af180156106ed57610820575b50506105d7565b8161082a91610f11565b6106e957815f610819565b93505090503d8082843e6108498184610f11565b8201906060838303126109995782516001600160401b0381116106e95783019160c0838203126106e9576040519260c084018481106001600160401b0382111761098557604052805184526108a0602082016110a4565b60208501526108b1604082016110a4565b6040850152606081015160608501526108cc608082016110a4565b608085015260a0810151906001600160401b038211610981570181601f82011215610969578051906108fd8261108d565b9161090b6040519384610f11565b80835260208084019160051b8301019184831161097d57602001905b82821061096d5750505060a084015260208401516001600160401b0381116109695761095a6109619260409287016110b8565b94016110a4565b92915f61039a565b8280fd5b8151815260209182019101610927565b8580fd5b8380fd5b634e487b7160e01b84526041600452602484fd5b80fd5b604051903d90823e3d90fd5b8680fd5b60405162461bcd60e51b815260206004820152601c60248201527f536f757263652061646472657373206d75737420626520656d707479000000006044820152606490fd5b610a0e6020610a036004880180610f8d565b0135866004016111b7565b610224565b5034610999576020366003190112610999576020610a32600435611794565b6040519015158152f35b610f4d565b5034610c5e57366003190160608112610c5e57604013610c5e576044356001600160401b038111610c5e5736602382011215610c5e578060040135610a8581610f32565b90610a936040519283610f11565b80825260208201923660248383010111610c5e57815f9260246020930186378301015260243591610ac383611794565b15610d6c5760405163eb3b597d60e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169290602081602481875afa8015610c53575f90610d2c575b6001600160401b0391501660018101926001600160401b038411610d1857600435936001600160401b038516809503610c5e576001600160401b0385911603610cd35760405191630a0f57a560e11b83528660048401526024830152602082604481885afa918215610c53575f92610c9f575b505f602091604051809186518091835e8101838152039060025afa15610c53575f5103610c6257823b15610c5e57610bf5925f928360405180968195829463024cd5a160e61b84526004840152896024840152606060448401526064830190611193565b03925af18015610c5357610c3e575b50610c0e81611794565b15610c17575080f35b7f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe28280a280f35b610c4b9192505f90610f11565b5f905f610c04565b6040513d5f823e3d90fd5b5f80fd5b60405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b6044820152606490fd5b9091506020813d602011610ccb575b81610cbb60209383610f11565b81010312610c5e5751905f610b91565b3d9150610cae565b60405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f7264657200000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b506020813d602011610d64575b81610d4660209383610f11565b81010312610c5e57610d5f6001600160401b03916110a4565b610b1e565b3d9150610d39565b60405162461bcd60e51b815260206004820152603560248201527f43616e6e6f74206170706c7920736861726420696620726567697374726174696044820152746f6e206973206e6f7420696e2070726f677265737360581b6064820152608490fd5b34610c5e575f366003190112610c5e576020610a3261171e565b34610c5e575f366003190112610c5e576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610c5e575f366003190112610c5e5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34610c5e576020366003190112610c5e576020610a326004356116aa565b34610c5e576040366003190112610c5e576004356001600160401b038111610c5e5760606003198236030112610c5e57610ec590602435906004016111b7565b005b60a081019081106001600160401b03821117610ee257604052565b634e487b7160e01b5f52604160045260245ffd5b604081019081106001600160401b03821117610ee257604052565b90601f801991011681019081106001600160401b03821117610ee257604052565b6001600160401b038111610ee257601f01601f191660200190565b34610c5e575f366003190112610c5e57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b903590609e1981360301821215610c5e570190565b3563ffffffff81168103610c5e5790565b15610fba57565b60405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606490fd5b903590601e1981360301821215610c5e57018035906001600160401b038211610c5e57602001918136038313610c5e57565b81601f82011215610c5e5780519061103e82610f32565b9261104c6040519485610f11565b82845260208383010111610c5e57815f9260208093018386015e8301015290565b908060209392818452848401375f828201840152601f01601f1916010190565b6001600160401b038111610ee25760051b60200190565b51906001600160401b0382168203610c5e57565b9080601f83011215610c5e578151916110d08361108d565b926110de6040519485610f11565b80845260208085019160051b83010191838311610c5e5760208101915b83831061110a57505050505090565b82516001600160401b038111610c5e578201906040828703601f190112610c5e576040519061113882610ef6565b6020830151916001600160401b038311610c5e5761116c6040856111648b602080999881990101611027565b8452016110a4565b838201528152019201916110fb565b90816020910312610c5e57518015158103610c5e5790565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b6111bf61171e565b15611630576111cd826116aa565b156115b55761122d916111e661010d6101088480610f8d565b73__$aaf4ae346b84a712cc43f25bb66199d6fb$__915f61120a6040830183610ff5565b6040516310b15a7360e31b8152602060048201529687928392602484019161106d565b0381865af4938415610c53575f94611526575b5060405163127936e760e21b815260048101929092525f826024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa918215610c53575f92611476575b50906113016112ac8360206112ef97950190610ff5565b604051968795630161c9f960e61b87526060600488015260206112db8251604060648b015260a48a0190611193565b910151878203606319016084890152611193565b8581036003190160248701529161106d565b60031983820301604484015260a08101918051825260208101519260a06020840152835180915260c0830190602060c08260051b8601019501915f905b828210611426575050505060209492849260806001600160401b0381858260408998015116604086015282606082015116606086015201511691015203915af4908115610c53575f916113f7575b501561139457565b60405162461bcd60e51b815260206004820152603560248201527f436f756c64206e6f74207665726966792049434d206d6573736167653a20536960448201527419db985d1d5c994818da1958dadcc819985a5b1959605a1b6064820152608490fd5b611419915060203d60201161141f575b6114118183610f11565b81019061117b565b5f61138c565b503d611407565b919394956001919397506020809160bf19898203018552895190826001600160401b038161145d8551604086526040860190611193565b940151169101529801920192018896959493919261133e565b9091503d805f833e6114888183610f11565b810190602081830312610c5e578051906001600160401b038211610c5e570160a081830312610c5e576040516114bd81610ec7565b8151815260208201516001600160401b038111610c5e576080836114ec61130196611519946112ac97016110b8565b60208501526114fd604082016110a4565b604085015261150e606082016110a4565b6060850152016110a4565b6080820152929150611295565b9093503d805f833e6115388183610f11565b810190602081830312610c5e578051906001600160401b038211610c5e5701604081830312610c5e576040519161156e83610ef6565b81516001600160401b038111610c5e578161158a918401611027565b835260208201516001600160401b038111610c5e576115a99201611027565b6020820152925f611240565b60405162461bcd60e51b815260206004820152604760248201527f4e6f2076616c696461746f72207365742069732072656769737465726564206660448201527f6f72207468652070726f7669646564204176616c616e63686520626c6f636b636064820152661a185a5b88125160ca1b608482015260a490fd5b60405162461bcd60e51b815260206004820152604660248201527f4120636f6d706c65746520502d636861696e2076616c696461746f72206d757360448201527f74206265207265676973746572656420746f207665726966792049434d206d6560648201526573736167657360d01b608482015260a490fd5b604051631392c59160e11b815260048101919091526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c53575f91611702575090565b61171b915060203d60201161141f576114118183610f11565b90565b604051631392c59160e11b81527f000000000000000000000000000000000000000000000000000000000000000060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c53575f91611702575090565b604051638457eaa760e01b815260048101919091526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa908115610c53575f91611702575090565b35906001600160a01b0382168203610c5e57565b9035601e1982360301811215610c5e5701602081359101916001600160401b038211610c5e578136038313610c5e57565b9080602083519182815201916020808360051b8301019401925f915b83831061185c57505050505090565b9091929394602080600192601f19858203018652885190826001600160401b03816118908551604086526040860190611193565b940151169101529701930193019193929061184d565b908152604060208201528151604082015260c06001600160401b0360806118dc602086015160a0606087015260e0860190611831565b9482604082015116828601528260608201511660a086015201511691015290565b906119078261108d565b6119146040519182610f11565b8281528092611925601f199161108d565b01905f5b82811061193557505050565b60209060405161194481610ef6565b606081525f8382015282828501015201611929565b805182101561196d5760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212207d2af914b49b221e45713d3931f8e96aba94cf065b0d3b36cd0f1f014ab12dc764736f6c634300081e0033",
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
