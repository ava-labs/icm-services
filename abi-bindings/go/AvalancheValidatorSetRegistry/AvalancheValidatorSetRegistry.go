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
	Payload            []byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID_\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"initialValidatorSetData\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"applyShard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvalancheNetworkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"totalWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"isRegistrationInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pChainInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"parseValidatorSetMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"pChainHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pChainTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"shardHashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structValidatorSetMetadata\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"registerValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"shardNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorSetShard\",\"name\":\"shard\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"shardBytes\",\"type\":\"bytes\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"updateValidatorSetWithDiff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"verifyICMMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052346102c9576121d980380380610019816102cd565b9283398101906040818303126102c95780519063ffffffff821682036102c9576020810151906001600160401b0382116102c95701906080828403126102c95760405192608084016001600160401b038111858210176102a55760405282518452610086602084016102f2565b916020850192835261009a604085016102f2565b60408601908152606085015190946001600160401b0382116102c957019180601f840112156102c9578251926001600160401b0384116102a5578360051b906020806100e78185016102cd565b8097815201928201019283116102c957602001905b8282106102b95750505060608501918252608052835160a08190525f81815260016020526040908190208451815487516001600160801b03199091166001600160401b039290921691909117921b6fffffffffffffffff00000000000000001691909117815590915180519060018301906001600160401b0383116102a5576801000000000000000083116102a557815483835580841061027f575b50602001905f5260205f205f5b83811061026b5760048501805460ff60401b1916680100000000000000001790555f86815260208190526040908190208a51815588516002919091018054600160401b600160801b0319169190921b6fffffffffffffffff00000000000000001617815588518154600160801b600160c01b03191660809190911b600160801b600160c01b0316179055604051611ed29081610307823960805181818160f701526113ef015260a051818181610a4e01528181610ff80152818161108a01526116440152f35b6001906020845194019381840155016101a5565b825f528360205f2091820191015b81811061029a5750610198565b5f815560010161028d565b634e487b7160e01b5f52604160045260245ffd5b81518152602091820191016100fc565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176102a557604052565b51906001600160401b03821682036102c95756fe60806040526004361015610011575f80fd5b5f3560e01c806327258b221461127a57806349e4db9c146110ad578063541dcba41461107357806357262e7f14611033578063580d632b14610fe15780636766233d14610dda57806368531ed014610dd557806382366d0514610dd55780638457eaa714610d9a5780638e91cb43146109dd57806393356840146109cd5780639def1e78146108aa57639fd530d4146100a8575f80fd5b34610583576020366003190112610583576004356001600160401b03811161058357806004019060806003198236030112610583578061011f6100ef602460449401611554565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614611565565b0135610141815f525f6020526001600160401b03600260405f20015416151590565b1561086557610162815f52600160205260ff600460405f20015460401c1690565b6107fa5781610174826101b79461163f565b5f61019473__$aaf4ae346b84a712cc43f25bb66199d6fb$__92806115a7565b60405163acac250f60e01b8152602060048201529586928392602484019161161f565b0381845af4928315610587575f936106cf575b5060405163127936e760e21b815260048101839052905f82602481305afa918215610587575f92610625575b5060808401916001600160401b038351166001600160401b0360608301511610156105e057602090949194015190604051948591630978573760e21b8352604483016040600485015284518091526064840190602060648260051b8701019601915f905b82821061059257505050509082806103245f9594600319838903016024840152805188526001600160401b0360208201511660208901526001600160401b036040820151166040890152606081015160608901526001600160401b03895116608089015260a08101976001600160401b0389511660a082015260c082015160c08201526101206103126102fe60e085015161014060e0860152610140850190611e03565b610100850151848203610100860152611e03565b92015190610120818403910152611e03565b03915af4918215610587575f945f93610529575b506001600160401b038091511691511690604051610355816112d3565b848152602081019586526001600160401b0360408201941684526060810191825260808101928352845f525f60205260405f2090518155600181019551805190600160401b821161051557875482895580831061045c575b50602001965f5260205f20965f905b82821061043e5787610418886001600160401b03896103f0828b818060028e01975116168219875416178655511684611a5a565b51825467ffffffffffffffff60801b1916911660801b67ffffffffffffffff60801b16179055565b7f3eb200e50e17828341d0b21af4671d123979b6e0e84ed7e47d43227a4fb52fe25f80a2005b6002602082610450600194518d611aac565b019901910190976103bc565b6001600160ff1b0381168103610501576001600160ff1b038316830361050157885f5260205f209060011b8101908360011b015b81811061049d57506103ad565b806104aa6002925461149b565b806104bd575b505f600182015501610490565b601f81116001146104d357505f81555b5f6104b0565b6104f090825f526001601f60205f20920160051c82019101611a96565b805f525f60208120818355556104cd565b634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b945091503d805f863e61053c8186611324565b8401604085820312610583578451906001600160401b0382116105835761057b60206105746001600160401b03949385948a01611d40565b9701611c38565b939150610338565b5f80fd5b6040513d5f823e3d90fd5b919394955091956020806001926063198d8203018552895190826001600160401b03816105c885516040865260408601906112af565b9401511691015298019201920189959493919261025a565b60405162461bcd60e51b815260206004820152601960248201527f496e76616c696420626c6f636b636861696e20686569676874000000000000006044820152606490fd5b9091503d805f833e6106378183611324565b810190602081830312610583578051906001600160401b038211610583570160a081830312610583576040519161066d836112d3565b815183526020820151916001600160401b038311610583576106966080926106c3948301611d40565b60208501526106a760408201611c38565b60408501526106b860608201611c38565b606085015201611c38565b6080820152905f6101f6565b9092503d805f833e6106e18183611324565b810190602081830312610583578051906001600160401b038211610583570161014081830312610583576040519161014083018381106001600160401b03821117610515576040528151835261073960208301611c38565b602084015261074a60408301611c38565b60408401526060820151606084015261076560808301611c38565b608084015261077660a08301611c38565b60a084015260c082015160c084015260e08201516001600160401b03811161058357816107a4918401611c4c565b60e08401526101008201516001600160401b03811161058357816107c9918401611c4c565b6101008401526101208201516001600160401b038111610583576107ed9201611c4c565b610120820152915f6101ca565b60405162461bcd60e51b815260206004820152603760248201527f43616e6e6f74206170706c792064696666207769746820616e6f74686572207260448201527f6567697374726174696f6e20696e2070726f67726573730000000000000000006064820152608490fd5b60405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f7220736574206e6f742072656769737465726564000000006044820152606490fd5b34610583576108b836611413565b5050506108c3611c14565b506108cc611c14565b60206040516108db8282611324565b5f81526040519160608352606060e08401948051828601526001600160401b03838201511660808601526001600160401b0360408201511660a0860152015193608060c085015284518091528161010085019501905f5b8181106109b9575050508284038184015281518085528185019180808360051b8801019401925f965b83881061096e575f604088015286860387f35b90919293948380600192601f19858203018652885190826001600160401b03816109a185516040865260408601906112af565b9401511691015297019301970196909392919361095b565b825187529583019591830191600101610932565b34610583576109db36611360565b005b34610583576109eb36611413565b50506109fc6100ef60208301611554565b604081013590610a1e825f52600160205260ff600460405f20015460401c1690565b610d475781610a43815f525f6020526001600160401b03600260405f20015416151590565b610d395750610a73907f00000000000000000000000000000000000000000000000000000000000000009061163f565b610a7b611c14565b50610a84611c14565b602090604051610a948382611324565b5f8152815193848103610cf45781519060608401916001835151115f14610c635784515f526001865260405f2090868601946001600160401b0380875116166001600160401b03198454161783556040870194610afb6001600160401b0387511685611a5a565b51600184018151916001600160401b03831161051557600160401b8311610515578a908254848455808510610c3b575b5095949501905f52895f205f5b838110610c29575050505060028201805467ffffffffffffffff1916600117905560048201805468ffffffffffffffffff1916600160401b1790556003909101905f5b838110610c0c5750505050610ba6905f525f6020526001600160401b03600260405f20015416151590565b15610bd7575b505050505b7f715216b8fb094b002b3a62b413e8a3d36b5af37f18205d2d08926df7fcb4ce935f80a2005b6103f060026001600160401b0380945f610c0398885182525260405f2096518755511694019384611a5a565b81808080610bac565b80610c23610c1c60019385611a82565b5185611be1565b01610b7b565b825182820155918b0191600101610b38565b835f5284835f2091820191015b818110610c555750610b2b565b5f81558d9350600101610c48565b949050610cc7915083515f525f81526001600160401b036040805f209580518755610c9e836002890195811987541687558301511685611a5a565b0151825467ffffffffffffffff60801b1916911660801b67ffffffffffffffff60801b16179055565b60015f9201915b838110610cde5750505050610bb1565b80610cee610c1c60019385611a82565b01610cce565b60405162461bcd60e51b815260048101859052601860248201527f536f7572636520636861696e204944206d69736d6174636800000000000000006044820152606490fd5b610d429161163f565b610a73565b60405162461bcd60e51b815260206004820152602560248201527f4120726567697374726174696f6e20697320616c726561647920696e2070726f604482015264677265737360d81b6064820152608490fd5b34610583576020366003190112610583576020610dcb6004355f52600160205260ff600460405f20015460401c1690565b6040519015158152f35b6113d3565b3461058357610de836611360565b602082013591610e0a835f52600160205260ff600460405f20015460401c1690565b15610f7e57825f52600160205260016001600160401b03600260405f20015416016001600160401b038111610501576001600160401b0380610e4b84611a46565b16911603610f3957825f5260016020526001600160401b03610e73600160405f200192611a46565b165f1981016001600160401b0381116105015782541115610f255760205f918193835282199082842001015493604051918183925191829101835e8101838152039060025afa15610587575f5103610ee857610ee1815f52600160205260ff600460405f20015460401c1690565b1561041857005b60405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b6044820152606490fd5b634e487b7160e01b5f52603260045260245ffd5b60405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f7264657200000000006044820152606490fd5b60405162461bcd60e51b815260206004820152603560248201527f43616e6e6f74206170706c7920736861726420696620726567697374726174696044820152746f6e206973206e6f7420696e2070726f677265737360581b6064820152608490fd5b34610583575f366003190112610583576020610dcb7f00000000000000000000000000000000000000000000000000000000000000005f525f6020526001600160401b03600260405f20015416151590565b34610583576040366003190112610583576004356001600160401b0381116105835760806003198236030112610583576109db906024359060040161163f565b34610583575f3660031901126105835760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b34610583576020366003190112610583575f60806040516110cd816112d3565b8281526060602082015282604082015282606082015201526004355f525f60205260405f20604051906110ff826112d3565b805482526001810190815461111381611484565b926111216040519485611324565b81845260208401905f5260205f205f915b83831061122e57868660028760208401928352015491604081016001600160401b038416815260608201906001600160401b038560401c1682526001600160401b03608084019560801c168552604051936020855260c0850193516020860152519260a06040860152835180915260e0850190602060e08260051b8801019501915f905b8282106111e6578780886001600160401b038c818b818c5116606087015251166080850152511660a08301520390f35b9091929560208060019260df198b8203018552895190826001600160401b038161121985516040865260408601906112af565b940151169101529801920192019092916111b6565b60026020600192604051611241816112ee565b60405161125981611252818a6114d3565b0382611324565b81526001600160401b03858701541683820152815201920192019190611132565b34610583576020366003190112610583576020610dcb6004355f525f6020526001600160401b03600260405f20015416151590565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60a081019081106001600160401b0382111761051557604052565b604081019081106001600160401b0382111761051557604052565b608081019081106001600160401b0382111761051557604052565b90601f801991011681019081106001600160401b0382111761051557604052565b6001600160401b03811161051557601f01601f191660200190565b9060031982016060811261058357604013610583576004916044356001600160401b0381116105835781602382011215610583578060040135906113a382611345565b926113b16040519485611324565b8284526024838301011161058357815f92602460209301838601378301015290565b34610583575f36600319011261058357602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b6040600319820112610583576004356001600160401b038111610583576080818303600319011261058357600401916024356001600160401b0381116105835782602382011215610583578060040135926001600160401b0384116105835760248483010111610583576024019190565b6001600160401b0381116105155760051b60200190565b90600182811c921680156114c9575b60208310146114b557565b634e487b7160e01b5f52602260045260245ffd5b91607f16916114aa565b5f92918154916114e28361149b565b808352926001811690811561153757506001146114fe57505050565b5f9081526020812093945091925b83831061151d575060209250010190565b60018160209294939454838587010152019101919061150c565b915050602093945060ff929192191683830152151560051b010190565b3563ffffffff811681036105835790565b1561156c57565b60405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606490fd5b903590601e198136030182121561058357018035906001600160401b0382116105835760200191813603831361058357565b81601f82011215610583578051906115f082611345565b926115fe6040519485611324565b8284526020838301011161058357815f9260208093018386015e8301015290565b908060209392818452848401375f828201840152601f01601f1916010190565b61167f7f00000000000000000000000000000000000000000000000000000000000000005f525f6020526001600160401b03600260405f20015416151590565b156119f5576116a4825f525f6020526001600160401b03600260405f20015416151590565b156119a057611703916116bc6100ef60208401611554565b73__$aaf4ae346b84a712cc43f25bb66199d6fb$__915f6116e060608301836115a7565b6040516310b15a7360e31b8152602060048201529687928392602484019161161f565b0381865af4938415610587575f94611908575b509061178a6117298361177896946115a7565b925f525f60205260405f2092604051968795630161c9f960e61b87526060600488015260206117648251604060648b015260a48a01906112af565b9101518782036063190160848901526112af565b8581036003190160248701529161161f565b60031983820301604484015260a081019180548252600181019260a06020840152835480915260c083019060c08160051b850101945f5260205f20915f905b8282106118bb575050505060209492849260806001600160401b036002869501548181166040850152818160401c166060850152821c1691015203915af4908115610587575f91611880575b501561181d57565b60405162461bcd60e51b815260206004820152603560248201527f436f756c64206e6f74207665726966792049434d206d6573736167653a20536960448201527419db985d1d5c994818da1958dadcc819985a5b1959605a1b6064820152608490fd5b90506020813d6020116118b3575b8161189b60209383611324565b8101031261058357518015158103610583575f611815565b3d915061188e565b91939495600191939750600260209160bf19898203018552604081526118e4604082018b6114d3565b90836001600160401b03868d015416910152980192019201889695949391926117c9565b9093503d805f833e61191a8183611324565b810190602081830312610583578051906001600160401b03821161058357016040818303126105835760405161194f816112ee565b81516001600160401b038111610583578361196b9184016115d9565b81526020820151916001600160401b0383116105835761178a936117299361199392016115d9565b6020820152949150611716565b60405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b6064820152608490fd5b356001600160401b03811681036105835790565b9067ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b8051821015610f255760209160051b010190565b818110611aa1575050565b5f8155600101611a96565b91909182519283516001600160401b03811161051557611acc835461149b565b601f8111611ba6575b506020601f8211600114611b3557600192826001600160401b039596979360209387955f92611b2a575b50505f19600383901b1c191690851b1786555b015116920191166001600160401b0319825416179055565b015190505f80611aff565b601f19821695845f52815f20965f5b818110611b8e5750836001600160401b03969798602094600197948997899510611b76575b505050811b018655611b12565b01515f1960f88460031b161c191690555f8080611b69565b83830151895560019098019760209384019301611b44565b611bd190845f5260205f20601f840160051c81019160208510611bd7575b601f0160051c0190611a96565b5f611ad5565b9091508190611bc4565b908154600160401b8110156105155760018101808455811015610f2557611c12925f5260205f209060011b01611aac565b565b60405190611c2182611309565b606080835f81525f60208201525f60408201520152565b51906001600160401b038216820361058357565b9080601f8301121561058357815191611c6483611484565b92611c726040519485611324565b80845260208085019160051b830101918383116105835760208101915b838310611c9e57505050505090565b82516001600160401b038111610583578201906080828703601f1901126105835760405190611ccc82611309565b60208301516bffffffffffffffffffffffff19811681036105835782526040830151916001600160401b03831161058357611d30608085611d158b6020809998819901016115d9565b86850152611d2560608201611c38565b604085015201611c38565b6060820152815201920191611c8f565b9080601f8301121561058357815191611d5883611484565b92611d666040519485611324565b80845260208085019160051b830101918383116105835760208101915b838310611d9257505050505090565b82516001600160401b038111610583578201906040828703601f1901126105835760405190611dc0826112ee565b6020830151916001600160401b03831161058357611df4604085611dec8b6020809998819901016115d9565b845201611c38565b83820152815201920191611d83565b9080602083519182815201916020808360051b8301019401925f915b838310611e2e57505050505090565b9091929394602080600192601f198582030186528851906bffffffffffffffffffffffff19825116815260606001600160401b0381611e7a8686015160808887015260808601906112af565b9482604082015116604086015201511691015297019301930191939290611e1f56fea26469706673582212207f5c44335d39afa258e84076db5555829def32ca51f5427451ab0caf2d9ddb2764736f6c634300081e0033",
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
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) , bytes ) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCaller) ParseValidatorSetMetadata(opts *bind.CallOpts, arg0 ICMMessage, arg1 []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	var out []interface{}
	err := _AvalancheValidatorSetRegistry.contract.Call(opts, &out, "parseValidatorSetMetadata", arg0, arg1)

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
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) , bytes ) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistrySession) ParseValidatorSetMetadata(arg0 ICMMessage, arg1 []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetRegistry.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetRegistry.CallOpts, arg0, arg1)
}

// ParseValidatorSetMetadata is a free data retrieval call binding the contract method 0x9def1e78.
//
// Solidity: function parseValidatorSetMetadata((bytes,uint32,bytes32,bytes) , bytes ) view returns((bytes32,uint64,uint64,bytes32[]), (bytes,uint64)[], uint64)
func (_AvalancheValidatorSetRegistry *AvalancheValidatorSetRegistryCallerSession) ParseValidatorSetMetadata(arg0 ICMMessage, arg1 []byte) (ValidatorSetMetadata, []Validator, uint64, error) {
	return _AvalancheValidatorSetRegistry.Contract.ParseValidatorSetMetadata(&_AvalancheValidatorSetRegistry.CallOpts, arg0, arg1)
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
