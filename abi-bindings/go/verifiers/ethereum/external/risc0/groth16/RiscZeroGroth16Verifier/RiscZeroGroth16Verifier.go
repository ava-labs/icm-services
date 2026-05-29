// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package risczerogroth16verifier

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

// Receipt is an auto generated low-level Go binding around an user-defined struct.
type Receipt struct {
	Seal        []byte
	ClaimDigest [32]byte
}

// RiscZeroGroth16VerifierMetaData contains all meta data concerning the RiscZeroGroth16Verifier contract.
var RiscZeroGroth16VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"control_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bn254_control_id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"received\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"}],\"name\":\"SelectorMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerificationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BN254_CONTROL_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTROL_ROOT_0\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTROL_ROOT_1\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"seal\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"journalDigest\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"seal\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"claimDigest\",\"type\":\"bytes32\"}],\"internalType\":\"structReceipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"verifyIntegrity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_pubSignals\",\"type\":\"uint256[5]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101808060405234610bf157604081611eb580380380916100208285610bf5565b833981010312610bf15780516020918201519091600883811c7eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff169084901b7fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff001617601081811c7dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff1691901b7fffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000161780821c7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff16911b7fffffffff00000000ffffffff00000000ffffffff00000000ffffffff000000001617604081811c77ffffffffffffffff0000000000000000ffffffffffffffff1691901b7fffffffffffffffff0000000000000000ffffffffffffffff00000000000000001617608081811c91901b176001600160801b031981811660a052608091821b16905260c08190526040517f72697363302e47726f74683136526563656970745665726966696572506172618152656d657465727360d01b602082810191909152905f9060269060025afa15610a7b575f5190600881811c7eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff1691901b7fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff001617601081811c7dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff1691901b7fffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff00001617602081811c7bffffffff00000000ffffffff00000000ffffffff00000000ffffffff1691901b7fffffffff00000000ffffffff00000000ffffffff00000000ffffffff000000001617604081811c77ffffffffffffffff0000000000000000ffffffffffffffff1691901b7fffffffffffffffff0000000000000000ffffffffffffffff00000000000000001617608081811c91901b179060e0604051916103068284610bf5565b60068352601f19820136602085013760205f6103846040517f12ac9a25dcd5e1a832a9061a082c15dd1d61aa9c4d553505739d0f5d65dc3be4848201527f025aa744581ebe7ad91731911c898569106ff5a2d30f3eee2b23c60ee980acd4604082015260408152610378606082610bf5565b60405191828092610c2c565b039060025afa15610a7b575f5161039a84610c55565b5260205f6103fe6040517f0707b920bc978c02f292fae2036e057be54294114ccc3c8769d883f688a1423f848201527f2e32a094b7589554f7bc357bf63481acd2d55555c203383782a4650787ff6642604082015260408152610378606082610bf5565b039060025afa15610a7b575f5161041484610c62565b5260205f6104786040517f0bca36e2cbe6394b3e249751853f961511011c7148e336f4fd974644850fc347848201527f2ede7c9acf48cf3a3729fa3d68714e2a8435d4fa6db8f7f409c153b1fcdf9b8b604082015260408152610378606082610bf5565b039060025afa15610a7b575f51835160021015610abc57606084015260205f6104f76040517f1b8af999dbfbb3927c091cc2aaf201e488cbacc3e2c6b6fb5a25f9112e04f2a7848201527f2b91a26aa92e1b6f5722949f192a81c850d586d81a60157f3e9cf04f679cccd6604082015260408152610378606082610bf5565b039060025afa15610a7b575f51835160031015610abc57608084015260205f6105766040517f2b5f494ed674235b8ac1750bdfd5a7615f002d4a1dcefeddd06eda5a076ccd0d848201527f2fe520ad2020aab9cbba817fcbb9a863b8a76ff88f14f912c5e71665b2ad5e82604082015260408152610378606082610bf5565b039060025afa15610a7b575f51835160041015610abc5760a084015260205f6105f56040517f0f1c3c0d5d9da0fa03666843cde4e82e869ba5252fce3c25d5940320b1c4d493848201527f214bfcff74f425f6fe8c0d07b307482d8bc8bb2f3608f68287aa01bd0b69e809604082015260408152610378606082610bf5565b039060025afa15610a7b575f51835160051015610abc5760c084015260205f601a6040517f72697363305f67726f746831362e566572696679696e674b6579000000000000815260025afa15610a7b575f519460205f6106ab6040517f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e2848201527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d1926604082015260408152610378606082610bf5565b039060025afa15610a7b575f519460205f6107686040517f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c848201527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab60408201527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a760608201527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec860808201526080815261037860a082610bf5565b039060025afa15610a7b575f519660205f6108256040517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2848201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60408201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60608201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60808201526080815261037860a082610bf5565b039060025afa15610a7b575f5160205f6108e16040517f03b03cd5effa95ac9bee94f1f5ef907157bda4812ccf0b4c91f42bb629f83a1c848201527f1aa085ff28179a12d922dba0547057ccaae94b9d69cfaa4e60401fea7f3e033360408201527f110c10134f200b19f6490846d518c9aea868366efb7228ca5c91d2940d03076260608201527f1e60f31fcbf757e837e867178318832d0b2d74d59e2fea1c7142df187d3fc6d360808201526080815261037860a082610bf5565b039060025afa15610a7b575f519060205f601d6040517f72697363305f67726f746831362e566572696679696e674b65792e4943000000815260025afa15610a7b575f8051610140526101008190526060610120526020610160525b885180610100511015610ae4575f19810190808211610ad0576101005190035f1901908111610ad0578951811015610abc57610160519060051b8a010151906040519161098d6101205184610bf5565b600283526101605160409036908501376109a683610c55565b526109b082610c62565b52604051906109c26101605183610bf5565b5f8252601f196101605101366101605184013780519061ffff8211610aa45791604051928391610140516101605184015260408301825190926101605101905f905b808210610a865750505092610a1f60029392610a4c95610c2c565b9061ffff60f01b9061ff0060ff8260081c169160081b161760f01b16815203601d19810184520182610bf5565b5f60405180610a5f816101605195610c2c565b039060025afa15610a7b575f516101008051600101905261093d565b6040513d5f823e3d90fd5b82518552610160518896509485019490920191600190910190610a04565b506306dfcc6560e41b5f52601060045260245260445ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5091908a8a604051956101605187015260408601526060850152608084015260a083015260c0820152600560f81b8582015260c28152610b2560e282610bf5565b5f60405180610b38816101605195610c2c565b039060025afa15610a7b575f51916040519361016051850152604084015260608301526080820152600360f81b60a082015260828152610b7960a282610bf5565b5f60405180610b8c816101605195610c2c565b039060025afa15610a7b575f516001600160e01b03191681526040516112429182610c7383396080518281816105c50152610dc5015260a0518281816105810152610deb015260c0518281816101690152610e2301525181818160ae0152610d310152f35b5f80fd5b601f909101601f19168101906001600160401b03821190821017610c1857604052565b634e487b7160e01b5f52604160045260245ffd5b908151915f5b838110610c42575050015f815290565b8060208092840101518185015201610c32565b805115610abc5760200190565b805160011015610abc576040019056fe60806040526004361015610011575f80fd5b5f3560e01c8063053c238d146100945780631599ead51461008f578063258038e21461008a57806334baeab9146100855780638989fa2e146100805780639181e4b11461007b578063ab750e75146100765763ffa1ad7414610071575f80fd5b61072a565b6105e9565b6105a5565b610561565b6101a7565b610152565b6100db565b346100d7575f3660031901126100d75763ffffffff60e01b7f00000000000000000000000000000000000000000000000000000000000000001660805260206080f35b5f80fd5b346100d75760203660031901126100d75760043567ffffffffffffffff81116100d75780360360406003198201126100d757600482013590602219018112156100d757810160048101359067ffffffffffffffff82116100d7576024019080360382136100d757602461015093013591610d2d565b005b346100d7575f3660031901126100d75760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b906004916044116100d757565b9060c491610104116100d757565b346100d7576101a03660031901126100d7576101c23661018c565b3660c4116100d7576101d336610199565b366101a4116100d757604051906103808201604052610104356101f58161078b565b61012435936102038561078b565b610144356102108161078b565b6101643561021d8161078b565b610184359161022b8361078b565b60808701977f12ac9a25dcd5e1a832a9061a082c15dd1d61aa9c4d553505739d0f5d65dc3be4885260208801957f025aa744581ebe7ad91731911c898569106ff5a2d30f3eee2b23c60ee980acd4875261028590896107bc565b61028f9088610848565b61029990876108d4565b6102a39086610960565b6102ad90856109ec565b803585527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760209182013581030660a085015260443560c085015260643560e085015260843561010085015260a4356101208501527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e26101408501527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d19266101608501527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101808501527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101a08501527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101c08501527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec86101e0850152835161020085015290516102208401527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102408401527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102608401527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102808401527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6102a084015281356102c084015201356102e08201527f03b03cd5effa95ac9bee94f1f5ef907157bda4812ccf0b4c91f42bb629f83a1c6103008201527f1aa085ff28179a12d922dba0547057ccaae94b9d69cfaa4e60401fea7f3e03336103208201527f110c10134f200b19f6490846d518c9aea868366efb7228ca5c91d2940d0307626103408201527f1e60f31fcbf757e837e867178318832d0b2d74d59e2fea1c7142df187d3fc6d36103609091015280806107cf195a01602092600861030092fa9051165f5260205ff35b346100d7575f3660031901126100d75760206040516001600160801b03197f0000000000000000000000000000000000000000000000000000000000000000168152f35b346100d7575f3660031901126100d75760206040516001600160801b03197f0000000000000000000000000000000000000000000000000000000000000000168152f35b346100d75760603660031901126100d75760043567ffffffffffffffff81116100d757366023820112156100d75780600401359067ffffffffffffffff82116100d75736602483830101116100d757610150916024359060246044359301610a78565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff82111761067c57604052565b61064c565b60a0810190811067ffffffffffffffff82111761067c57604052565b6060810190811067ffffffffffffffff82111761067c57604052565b90601f8019910116810190811067ffffffffffffffff82111761067c57604052565b604051906106ea6040836106b9565b565b604051906106ea60a0836106b9565b906106ea60405192836106b9565b5f5b83811061071a5750505f910152565b818101518382015260200161070b565b346100d7575f3660031901126100d75761077d6040805161074a81610660565b600a81526020810169352e302e302d72632e3160b01b815282519384926020845251809281602086015285850190610709565b601f01601f19168101030190f35b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000111156107b457565b5f805260205ff35b604051917f0707b920bc978c02f292fae2036e057be54294114ccc3c8769d883f688a1423f83527f2e32a094b7589554f7bc357bf63481acd2d55555c203383782a4650787ff664260208401526040830190815260408360608160076107cf195a01fa156107b457815190526020810151606083015260409160809060066107cf195a01fa156107b457565b604051917f0bca36e2cbe6394b3e249751853f961511011c7148e336f4fd974644850fc34783527f2ede7c9acf48cf3a3729fa3d68714e2a8435d4fa6db8f7f409c153b1fcdf9b8b60208401526040830190815260408360608160076107cf195a01fa156107b457815190526020810151606083015260409160809060066107cf195a01fa156107b457565b604051917f1b8af999dbfbb3927c091cc2aaf201e488cbacc3e2c6b6fb5a25f9112e04f2a783527f2b91a26aa92e1b6f5722949f192a81c850d586d81a60157f3e9cf04f679cccd660208401526040830190815260408360608160076107cf195a01fa156107b457815190526020810151606083015260409160809060066107cf195a01fa156107b457565b604051917f2b5f494ed674235b8ac1750bdfd5a7615f002d4a1dcefeddd06eda5a076ccd0d83527f2fe520ad2020aab9cbba817fcbb9a863b8a76ff88f14f912c5e71665b2ad5e8260208401526040830190815260408360608160076107cf195a01fa156107b457815190526020810151606083015260409160809060066107cf195a01fa156107b457565b604051917f0f1c3c0d5d9da0fa03666843cde4e82e869ba5252fce3c25d5940320b1c4d49383527f214bfcff74f425f6fe8c0d07b307482d8bc8bb2f3608f68287aa01bd0b69e80960208401526040830190815260408360608160076107cf195a01fa156107b457815190526020810151606083015260409160809060066107cf195a01fa156107b457565b91610b2d906106ea945f6080604051610a9081610681565b828152826020820152604051610aa581610660565b83815283602082015260408201528260608201520152610ae6610ac66106db565b915f83525f6020840152610ad86106db565b9081525f60208201526111d4565b90610aef6106ec565b9283527fa3acc27117418996340b84e5a90f3ef4c49d22c79e44aad822ec9c313e1eb8e2602084015260408301525f60608301526080820152610f66565b91610d2d565b906004116100d75790600490565b90929192836004116100d75783116100d757600401916003190190565b356001600160e01b0319811692919060048210610b79575050565b6001600160e01b031960049290920360031b82901b16169150565b9080601f830112156100d75760405191610baf6040846106b9565b8290604081019283116100d757905b828210610bcb5750505090565b8135815260209182019101610bbe565b610100818303126100d75760405191610bf38361069d565b610bfd8183610b94565b835280605f830112156100d7576040918251610c1984826106b9565b8060c08301928484116100d75785809101915b848310610c4c575050506020850152610c459190610b94565b9082015290565b602090610c598785610b94565b8152019101908590610c2c565b908160209103126100d7575180151581036100d75790565b905f905b60028210610c8f57505050565b6020806001928551815201930191019091610c82565b919493929094610cba836101a0810197610c7e565b5f604084015b60028210610d045750505090610cdd6101009260c0830190610c7e565b015f905b60058210610cee57505050565b6020806001928551815201930191019091610ce1565b6020604082610d166001948751610c7e565b01930191019091610cc0565b6040513d5f823e3d90fd5b90917f0000000000000000000000000000000000000000000000000000000000000000610d73610d66610d608686610b33565b90610b5e565b6001600160e01b03191690565b6001600160e01b0319821603610ec0575090610da7610d9f84610d97602095611051565b969094610b41565b810190610bdb565b90610e6282519160408585015194015195610dc260a06106fb565b917f000000000000000000000000000000000000000000000000000000000000000060801c83527f000000000000000000000000000000000000000000000000000000000000000060801c8784015260801c604083015260801c60608201527f0000000000000000000000000000000000000000000000000000000000000000608082015260405195869485946334baeab960e01b865260048601610ca5565b0381305afa908115610ebb575f91610e8c575b5015610e7d57565b63439cc0cd60e01b5f5260045ffd5b610eae915060203d602011610eb4575b610ea681836106b9565b810190610c66565b5f610e75565b503d610e9c565b610d22565b610ef390610ed1610d608686610b33565b632e2ce35360e21b5f526001600160e01b031990811660045216602452604490565b5ffd5b60031115610f0057565b634e487b7160e01b5f52602160045260245ffd5b60205f60126040517172697363302e52656365697074436c61696d60701b815260025afa15610ebb575f5190565b516003811015610f005790565b90610f6260209282815194859201610709565b0190565b5f611041602092611035610f78610f14565b611027606084015193805190888101519060406080820151910190610fcf610fb3610fc98d610fbf610faa8751610f42565b610fb381610ef6565b60181b63ff0000001690565b9551015160ff1690565b60ff1690565b604080518d8101988952602089019a909a52870194909452606086019290925260808501919091526001600160e01b031960e091821b811660a086015291901b1660a4830152600160fa1b60a8830152839160aa0190565b03601f1981018352826106b9565b60405191828092610f4f565b039060025afa15610ebb575f5190565b8060081c9060081b907cff000000ff000000ff000000ff000000ff000000ff000000ff000000ff7dff000000ff000000ff000000ff000000ff000000ff000000ff000000ff007fff000000ff000000ff000000ff000000ff000000ff000000ff000000ff00000084167eff000000ff000000ff000000ff000000ff000000ff000000ff000000ff000084161760101c931691161760101b176111377bffffffff00000000ffffffff00000000ffffffff00000000ffffffff7fffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000831660201c921660201b90565b17604081811c77ffffffffffffffff0000000000000000ffffffffffffffff169177ffffffffffffffff0000000000000000ffffffffffffffff19911b161761118a6111838260801c90565b9160801b90565b17906111c16111a861119c8460801c90565b6001600160801b031690565b60801b6fffffffffffffffffffffffffffffffff191690565b916001600160801b03199060801b169190565b60205f600c6040516b1c9a5cd8cc0b93dd5d1c1d5d60a21b815260025afa15610ebb575f8051825160209384015160408051808701949094528301919091526060820152600160f91b608082015260628152611041906110356082826106b956fea164736f6c634300081e000a",
}

// RiscZeroGroth16VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use RiscZeroGroth16VerifierMetaData.ABI instead.
var RiscZeroGroth16VerifierABI = RiscZeroGroth16VerifierMetaData.ABI

// RiscZeroGroth16VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RiscZeroGroth16VerifierMetaData.Bin instead.
var RiscZeroGroth16VerifierBin = RiscZeroGroth16VerifierMetaData.Bin

// DeployRiscZeroGroth16Verifier deploys a new Ethereum contract, binding an instance of RiscZeroGroth16Verifier to it.
func DeployRiscZeroGroth16Verifier(auth *bind.TransactOpts, backend bind.ContractBackend, control_root [32]byte, bn254_control_id [32]byte) (common.Address, *types.Transaction, *RiscZeroGroth16Verifier, error) {
	parsed, err := RiscZeroGroth16VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RiscZeroGroth16VerifierBin), backend, control_root, bn254_control_id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RiscZeroGroth16Verifier{RiscZeroGroth16VerifierCaller: RiscZeroGroth16VerifierCaller{contract: contract}, RiscZeroGroth16VerifierTransactor: RiscZeroGroth16VerifierTransactor{contract: contract}, RiscZeroGroth16VerifierFilterer: RiscZeroGroth16VerifierFilterer{contract: contract}}, nil
}

// RiscZeroGroth16Verifier is an auto generated Go binding around an Ethereum contract.
type RiscZeroGroth16Verifier struct {
	RiscZeroGroth16VerifierCaller     // Read-only binding to the contract
	RiscZeroGroth16VerifierTransactor // Write-only binding to the contract
	RiscZeroGroth16VerifierFilterer   // Log filterer for contract events
}

// RiscZeroGroth16VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type RiscZeroGroth16VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RiscZeroGroth16VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RiscZeroGroth16VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RiscZeroGroth16VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RiscZeroGroth16VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RiscZeroGroth16VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RiscZeroGroth16VerifierSession struct {
	Contract     *RiscZeroGroth16Verifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RiscZeroGroth16VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RiscZeroGroth16VerifierCallerSession struct {
	Contract *RiscZeroGroth16VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RiscZeroGroth16VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RiscZeroGroth16VerifierTransactorSession struct {
	Contract     *RiscZeroGroth16VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RiscZeroGroth16VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type RiscZeroGroth16VerifierRaw struct {
	Contract *RiscZeroGroth16Verifier // Generic contract binding to access the raw methods on
}

// RiscZeroGroth16VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RiscZeroGroth16VerifierCallerRaw struct {
	Contract *RiscZeroGroth16VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// RiscZeroGroth16VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RiscZeroGroth16VerifierTransactorRaw struct {
	Contract *RiscZeroGroth16VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRiscZeroGroth16Verifier creates a new instance of RiscZeroGroth16Verifier, bound to a specific deployed contract.
func NewRiscZeroGroth16Verifier(address common.Address, backend bind.ContractBackend) (*RiscZeroGroth16Verifier, error) {
	contract, err := bindRiscZeroGroth16Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RiscZeroGroth16Verifier{RiscZeroGroth16VerifierCaller: RiscZeroGroth16VerifierCaller{contract: contract}, RiscZeroGroth16VerifierTransactor: RiscZeroGroth16VerifierTransactor{contract: contract}, RiscZeroGroth16VerifierFilterer: RiscZeroGroth16VerifierFilterer{contract: contract}}, nil
}

// NewRiscZeroGroth16VerifierCaller creates a new read-only instance of RiscZeroGroth16Verifier, bound to a specific deployed contract.
func NewRiscZeroGroth16VerifierCaller(address common.Address, caller bind.ContractCaller) (*RiscZeroGroth16VerifierCaller, error) {
	contract, err := bindRiscZeroGroth16Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RiscZeroGroth16VerifierCaller{contract: contract}, nil
}

// NewRiscZeroGroth16VerifierTransactor creates a new write-only instance of RiscZeroGroth16Verifier, bound to a specific deployed contract.
func NewRiscZeroGroth16VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*RiscZeroGroth16VerifierTransactor, error) {
	contract, err := bindRiscZeroGroth16Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RiscZeroGroth16VerifierTransactor{contract: contract}, nil
}

// NewRiscZeroGroth16VerifierFilterer creates a new log filterer instance of RiscZeroGroth16Verifier, bound to a specific deployed contract.
func NewRiscZeroGroth16VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*RiscZeroGroth16VerifierFilterer, error) {
	contract, err := bindRiscZeroGroth16Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RiscZeroGroth16VerifierFilterer{contract: contract}, nil
}

// bindRiscZeroGroth16Verifier binds a generic wrapper to an already deployed contract.
func bindRiscZeroGroth16Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RiscZeroGroth16VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RiscZeroGroth16Verifier.Contract.RiscZeroGroth16VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RiscZeroGroth16Verifier.Contract.RiscZeroGroth16VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RiscZeroGroth16Verifier.Contract.RiscZeroGroth16VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RiscZeroGroth16Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RiscZeroGroth16Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RiscZeroGroth16Verifier.Contract.contract.Transact(opts, method, params...)
}

// BN254CONTROLID is a free data retrieval call binding the contract method 0x258038e2.
//
// Solidity: function BN254_CONTROL_ID() view returns(bytes32)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) BN254CONTROLID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "BN254_CONTROL_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BN254CONTROLID is a free data retrieval call binding the contract method 0x258038e2.
//
// Solidity: function BN254_CONTROL_ID() view returns(bytes32)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) BN254CONTROLID() ([32]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.BN254CONTROLID(&_RiscZeroGroth16Verifier.CallOpts)
}

// BN254CONTROLID is a free data retrieval call binding the contract method 0x258038e2.
//
// Solidity: function BN254_CONTROL_ID() view returns(bytes32)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) BN254CONTROLID() ([32]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.BN254CONTROLID(&_RiscZeroGroth16Verifier.CallOpts)
}

// CONTROLROOT0 is a free data retrieval call binding the contract method 0x9181e4b1.
//
// Solidity: function CONTROL_ROOT_0() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) CONTROLROOT0(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "CONTROL_ROOT_0")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// CONTROLROOT0 is a free data retrieval call binding the contract method 0x9181e4b1.
//
// Solidity: function CONTROL_ROOT_0() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) CONTROLROOT0() ([16]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.CONTROLROOT0(&_RiscZeroGroth16Verifier.CallOpts)
}

// CONTROLROOT0 is a free data retrieval call binding the contract method 0x9181e4b1.
//
// Solidity: function CONTROL_ROOT_0() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) CONTROLROOT0() ([16]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.CONTROLROOT0(&_RiscZeroGroth16Verifier.CallOpts)
}

// CONTROLROOT1 is a free data retrieval call binding the contract method 0x8989fa2e.
//
// Solidity: function CONTROL_ROOT_1() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) CONTROLROOT1(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "CONTROL_ROOT_1")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// CONTROLROOT1 is a free data retrieval call binding the contract method 0x8989fa2e.
//
// Solidity: function CONTROL_ROOT_1() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) CONTROLROOT1() ([16]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.CONTROLROOT1(&_RiscZeroGroth16Verifier.CallOpts)
}

// CONTROLROOT1 is a free data retrieval call binding the contract method 0x8989fa2e.
//
// Solidity: function CONTROL_ROOT_1() view returns(bytes16)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) CONTROLROOT1() ([16]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.CONTROLROOT1(&_RiscZeroGroth16Verifier.CallOpts)
}

// SELECTOR is a free data retrieval call binding the contract method 0x053c238d.
//
// Solidity: function SELECTOR() view returns(bytes4)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) SELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// SELECTOR is a free data retrieval call binding the contract method 0x053c238d.
//
// Solidity: function SELECTOR() view returns(bytes4)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) SELECTOR() ([4]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.SELECTOR(&_RiscZeroGroth16Verifier.CallOpts)
}

// SELECTOR is a free data retrieval call binding the contract method 0x053c238d.
//
// Solidity: function SELECTOR() view returns(bytes4)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) SELECTOR() ([4]byte, error) {
	return _RiscZeroGroth16Verifier.Contract.SELECTOR(&_RiscZeroGroth16Verifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) VERSION() (string, error) {
	return _RiscZeroGroth16Verifier.Contract.VERSION(&_RiscZeroGroth16Verifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) VERSION() (string, error) {
	return _RiscZeroGroth16Verifier.Contract.VERSION(&_RiscZeroGroth16Verifier.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xab750e75.
//
// Solidity: function verify(bytes seal, bytes32 imageId, bytes32 journalDigest) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) Verify(opts *bind.CallOpts, seal []byte, imageId [32]byte, journalDigest [32]byte) error {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "verify", seal, imageId, journalDigest)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0xab750e75.
//
// Solidity: function verify(bytes seal, bytes32 imageId, bytes32 journalDigest) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) Verify(seal []byte, imageId [32]byte, journalDigest [32]byte) error {
	return _RiscZeroGroth16Verifier.Contract.Verify(&_RiscZeroGroth16Verifier.CallOpts, seal, imageId, journalDigest)
}

// Verify is a free data retrieval call binding the contract method 0xab750e75.
//
// Solidity: function verify(bytes seal, bytes32 imageId, bytes32 journalDigest) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) Verify(seal []byte, imageId [32]byte, journalDigest [32]byte) error {
	return _RiscZeroGroth16Verifier.Contract.Verify(&_RiscZeroGroth16Verifier.CallOpts, seal, imageId, journalDigest)
}

// VerifyIntegrity is a free data retrieval call binding the contract method 0x1599ead5.
//
// Solidity: function verifyIntegrity((bytes,bytes32) receipt) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) VerifyIntegrity(opts *bind.CallOpts, receipt Receipt) error {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "verifyIntegrity", receipt)

	if err != nil {
		return err
	}

	return err

}

// VerifyIntegrity is a free data retrieval call binding the contract method 0x1599ead5.
//
// Solidity: function verifyIntegrity((bytes,bytes32) receipt) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) VerifyIntegrity(receipt Receipt) error {
	return _RiscZeroGroth16Verifier.Contract.VerifyIntegrity(&_RiscZeroGroth16Verifier.CallOpts, receipt)
}

// VerifyIntegrity is a free data retrieval call binding the contract method 0x1599ead5.
//
// Solidity: function verifyIntegrity((bytes,bytes32) receipt) view returns()
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) VerifyIntegrity(receipt Receipt) error {
	return _RiscZeroGroth16Verifier.Contract.VerifyIntegrity(&_RiscZeroGroth16Verifier.CallOpts, receipt)
}

// VerifyProof is a free data retrieval call binding the contract method 0x34baeab9.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[5] _pubSignals) view returns(bool)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [5]*big.Int) (bool, error) {
	var out []interface{}
	err := _RiscZeroGroth16Verifier.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x34baeab9.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[5] _pubSignals) view returns(bool)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [5]*big.Int) (bool, error) {
	return _RiscZeroGroth16Verifier.Contract.VerifyProof(&_RiscZeroGroth16Verifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x34baeab9.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[5] _pubSignals) view returns(bool)
func (_RiscZeroGroth16Verifier *RiscZeroGroth16VerifierCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [5]*big.Int) (bool, error) {
	return _RiscZeroGroth16Verifier.Contract.VerifyProof(&_RiscZeroGroth16Verifier.CallOpts, _pA, _pB, _pC, _pubSignals)
}
