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
	Bin: "0x60c0604052346102ca576112d080380380610019816102ce565b9283398101906040818303126102ca5780519063ffffffff821682036102ca576020810151906001600160401b0382116102ca5701906080828403126102ca5760405192608084016001600160401b038111858210176102a65760405282518452610086602084016102f3565b916020850192835261009a604085016102f3565b60408601908152606085015190946001600160401b0382116102ca57019180601f840112156102ca578251926001600160401b0384116102a6578360051b906020806100e78185016102ce565b8097815201928201019283116102ca57602001905b8282106102ba5750505060608501918252608052835160a08190525f81815260016020526040908190208451815487516001600160801b03199091166001600160401b039290921691909117921b6fffffffffffffffff00000000000000001691909117815590915180519060018301906001600160401b0383116102a6576801000000000000000083116102a6578154838355808410610280575b50602001905f5260205f205f5b83811061026c5760048501805460ff60401b1916680100000000000000001790555f86815260208190526040908190208a51815588516002919091018054600160401b600160801b0319169190921b6fffffffffffffffff00000000000000001617815588518154600160801b600160c01b03191660809190911b600160801b600160c01b0316179055604051610fc89081610308823960805181818161018e01526108ee015260a051818181610206015281816104d60152818161056a0152610aff0152f35b6001906020845194019381840155016101a5565b825f528360205f2091820191015b81811061029b5750610198565b5f815560010161028e565b634e487b7160e01b5f52604160045260245ffd5b81518152602091820191016100fc565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176102a657604052565b51906001600160401b03821682036102ca5756fe60806040526004361015610011575f80fd5b5f3560e01c806327258b221461077957806349e4db9c1461058d578063541dcba41461055357806357262e7f14610511578063580d632b146104bf5780636766233d146102db57806368531ed0146102d657806382366d05146102d65780638457eaa71461029b5780638e91cb431461016757806393356840146101525780639def1e78146100df57639fd530d4146100a8575f80fd5b346100db5760203660031901126100db576004356001600160401b0381116100db57608090600319903603011215610f5b575b5f80fd5b346100db576100ed36610912565b5050506060806040516100ff81610808565b5f81525f60208201525f6040820152015260405162461bcd60e51b81528061014e6004820160609060208152600f60208201526e139bdd081a5b5c1b195b595b9d1959608a1b60408201520190565b0390fd5b346100db576101603661085f565b5050610f5b565b346100db5761017536610912565b50506101b661018660208301610a2e565b63ffffffff807f000000000000000000000000000000000000000000000000000000000000000016911614610a3f565b60408101356101d7815f52600160205260ff600460405f20015460401c1690565b610248576101fb815f525f6020526001600160401b03600260405f20015416151590565b61023a575061022b907f000000000000000000000000000000000000000000000000000000000000000090610af9565b6060806040516100ff81610808565b61024391610af9565b61022b565b60405162461bcd60e51b815260206004820152602560248201527f4120726567697374726174696f6e20697320616c726561647920696e2070726f604482015264677265737360d81b6064820152608490fd5b346100db5760203660031901126100db5760206102cc6004355f52600160205260ff600460405f20015460401c1690565b6040519015158152f35b6108d2565b346100db576102e93661085f565b90602081013561030b815f52600160205260ff600460405f20015460401c1690565b1561047a57805f52600160205260016001600160401b03600260405f20015416016001600160401b038111610421576001600160401b038061034c85610f47565b16911603610435575f5260016020526001600160401b03610373600160405f200192610f47565b165f1981016001600160401b038111610421578254111561040d5760205f918193835282199082842001015493604051918183925191829101835e8101838152039060025afa15610402575f5114610f5b5760405162461bcd60e51b81526020600482015260156024820152740aadccaf0e0cac6e8cac840e6d0c2e4c840d0c2e6d605b1b6044820152606490fd5b6040513d5f823e3d90fd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152601b60248201527f5265636569766564207368617264206f7574206f66206f7264657200000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601f60248201527f526567697374726174696f6e206973206e6f7420696e2070726f6772657373006044820152606490fd5b346100db575f3660031901126100db5760206102cc7f00000000000000000000000000000000000000000000000000000000000000005f525f6020526001600160401b03600260405f20015416151590565b346100db5760403660031901126100db576004356001600160401b0381116100db57608060031982360301126100db576105519060243590600401610af9565b005b346100db575f3660031901126100db5760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346100db5760203660031901126100db575f60806040516105ad816107d2565b8281526060602082015282604082015282606082015201526004355f525f60205260405f20604051906105df826107d2565b80548252600181019081546001600160401b038111610765576040519261060c60208360051b0185610823565b81845260208401905f5260205f205f915b83831061071957868660028760208401928352015491604081016001600160401b038416815260608201906001600160401b038560401c1682526001600160401b03608084019560801c168552604051936020855260c0850193516020860152519260a06040860152835180915260e0850190602060e08260051b8801019501915f905b8282106106d1578780886001600160401b038c818b818c5116606087015251166080850152511660a08301520390f35b9091929560208060019260df198b8203018552895190826001600160401b038161070485516040865260408601906107ae565b940151169101529801920192019092916106a1565b6002602060019260405161072c816107ed565b6040516107448161073d818a610983565b0382610823565b81526001600160401b0385870154168382015281520192019201919061061d565b634e487b7160e01b5f52604160045260245ffd5b346100db5760203660031901126100db5760206102cc6004355f525f6020526001600160401b03600260405f20015416151590565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b60a081019081106001600160401b0382111761076557604052565b604081019081106001600160401b0382111761076557604052565b608081019081106001600160401b0382111761076557604052565b90601f801991011681019081106001600160401b0382111761076557604052565b6001600160401b03811161076557601f01601f191660200190565b906003198201606081126100db576040136100db576004916044356001600160401b0381116100db57816023820112156100db578060040135906108a282610844565b926108b06040519485610823565b828452602483830101116100db57815f92602460209301838601378301015290565b346100db575f3660031901126100db57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60406003198201126100db576004356001600160401b0381116100db57608081830360031901126100db57600401916024356001600160401b0381116100db57826023820112156100db578060040135926001600160401b0384116100db57602484830101116100db576024019190565b5f92918154918260011c92600181168015610a24575b602085108114610a10578484529081156109f357506001146109ba57505050565b5f9081526020812093945091925b8383106109d9575060209250010190565b6001816020929493945483858701015201910191906109c8565b915050602093945060ff929192191683830152151560051b010190565b634e487b7160e01b5f52602260045260245ffd5b93607f1693610999565b3563ffffffff811681036100db5790565b15610a4657565b60405162461bcd60e51b815260206004820152601360248201527209ccae8eedee4d640928840dad2e6dac2e8c6d606b1b6044820152606490fd5b903590601e19813603018212156100db57018035906001600160401b0382116100db576020019181360383136100db57565b81601f820112156100db57805190610aca82610844565b92610ad86040519485610823565b828452602083830101116100db57815f9260208093018386015e8301015290565b90610b3a7f00000000000000000000000000000000000000000000000000000000000000005f525f6020526001600160401b03600260405f20015416151590565b15610ef657610b5f815f525f6020526001600160401b03600260405f20015416151590565b15610ea1578160206044930191610b7861018684610a2e565b73__$aaf4ae346b84a712cc43f25bb66199d6fb$__925f610b9c6060850185610a81565b96879160405198899384926310b15a7360e31b845260206004850152816024850152848401378181018301859052601f01601f19168101030181875af4948515610402575f95610e05575b5090610c426044610bfc610c93979594610a2e565b610c068580610a81565b80916040805198899563ffffffff60e01b9060e01b166020870152013560248501528484013781015f838201520301601f198101845283610823565b5f525f602052610ca560405f2091604051958694630161c9f960e61b8652606060048701526020610c7f8251604060648a015260a48901906107ae565b9101518682036063190160848801526107ae565b848103600319016024860152906107ae565b60031983820301604484015260a081019180548252600181019260a06020840152835480915260c083019060c08160051b850101945f5260205f20915f905b828210610db8575050505060209492849260806001600160401b036002869501548181166040850152818160401c166060850152821c1691015203915af4908115610402575f91610d7d575b5015610d3857565b60405162461bcd60e51b815260206004820152601b60248201527f4661696c656420746f20766572696679207369676e61747572657300000000006044820152606490fd5b90506020813d602011610db0575b81610d9860209383610823565b810103126100db575180151581036100db575f610d30565b3d9150610d8b565b91939495600191939750600260209160bf1989820301855260408152610de1604082018b610983565b90836001600160401b03868d01541691015298019201920188969594939192610ce4565b9094503d805f833e610e178183610823565b8101906020818303126100db578051906001600160401b0382116100db5701906040828203126100db57604051610e4d816107ed565b82516001600160401b0381116100db5782610e69918501610ab3565b81526020830151906001600160401b0382116100db57610e93604493610bfc93610c429601610ab3565b602082015296925050610be7565b60405162461bcd60e51b815260206004820152602760248201527f4e6f2076616c696461746f7220736574207265676973746572656420746f20676044820152661a5d995b88125160ca1b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f4e6f20502d636861696e2076616c696461746f722073657420726567697374656044820152633932b21760e11b6064820152608490fd5b356001600160401b03811681036100db5790565b60405162461bcd60e51b815260206004820152600f60248201526e139bdd081a5b5c1b195b595b9d1959608a1b6044820152606490fdfea26469706673582212204fffec5b622cd0e1a6e88290d2fcf94bb534ee7fc6ebf87ec2de7200e51ed8f964736f6c634300081e0033",
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
