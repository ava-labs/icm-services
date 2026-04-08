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
	Bin: "0x610100604052348015610010575f5ffd5b5060405161247038038061247083398101604081905261002f91611047565b610038826102b9565b6001600160801b031990811660a0521660805260c08190526040517f72697363302e47726f74683136526563656970745665726966696572506172618152656d657465727360d01b60208201526002908190602601602060405180830381855afa1580156100a8573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906100cb9190611069565b8361021b845f8190506008817eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff16901b6008827fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0016901c1790506010817dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff16901b6010827fffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff000016901c1790506020817bffffffff00000000ffffffff00000000ffffffff00000000ffffffff16901b6020827fffffffff00000000ffffffff00000000ffffffff00000000ffffffff0000000016901c17905060408177ffffffffffffffff0000000000000000ffffffffffffffff16901b6040827fffffffffffffffff0000000000000000ffffffffffffffff000000000000000016901c179050608081901b608082901c179050919050565b610223610425565b60408051602081019590955284019290925260608301526080820152600360f81b60a082015260a20160408051601f1981840301815290829052610266916110ad565b602060405180830381855afa158015610281573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906102a49190611069565b6001600160e01b03191660e0525061114b9050565b5f808061040b845f8190506008817eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff16901b6008827fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0016901c1790506010817dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff16901b6010827fffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff000016901c1790506020817bffffffff00000000ffffffff00000000ffffffff00000000ffffffff16901b6020827fffffffff00000000ffffffff00000000ffffffff00000000ffffffff0000000016901c17905060408177ffffffffffffffff0000000000000000ffffffffffffffff16901b6040827fffffffffffffffff0000000000000000ffffffffffffffff000000000000000016901c179050608081901b608082901c179050919050565b608081901b956001600160801b0319909116945092505050565b60408051600680825260e082019092525f918291906020820160c08036833701905050905060027f12ac9a25dcd5e1a832a9061a082c15dd1d61aa9c4d553505739d0f5d65dc3be47f025aa744581ebe7ad91731911c898569106ff5a2d30f3eee2b23c60ee980acd46040516020016104a8929190918252602082015260400190565b60408051601f19818403018152908290526104c2916110ad565b602060405180830381855afa1580156104dd573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906105009190611069565b815f81518110610512576105126110bf565b60200260200101818152505060027f0707b920bc978c02f292fae2036e057be54294114ccc3c8769d883f688a1423f7f2e32a094b7589554f7bc357bf63481acd2d55555c203383782a4650787ff664260405160200161057c929190918252602082015260400190565b60408051601f1981840301815290829052610596916110ad565b602060405180830381855afa1580156105b1573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906105d49190611069565b816001815181106105e7576105e76110bf565b60200260200101818152505060027f0bca36e2cbe6394b3e249751853f961511011c7148e336f4fd974644850fc3477f2ede7c9acf48cf3a3729fa3d68714e2a8435d4fa6db8f7f409c153b1fcdf9b8b604051602001610651929190918252602082015260400190565b60408051601f198184030181529082905261066b916110ad565b602060405180830381855afa158015610686573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906106a99190611069565b816002815181106106bc576106bc6110bf565b60200260200101818152505060027f1b8af999dbfbb3927c091cc2aaf201e488cbacc3e2c6b6fb5a25f9112e04f2a77f2b91a26aa92e1b6f5722949f192a81c850d586d81a60157f3e9cf04f679cccd6604051602001610726929190918252602082015260400190565b60408051601f1981840301815290829052610740916110ad565b602060405180830381855afa15801561075b573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061077e9190611069565b81600381518110610791576107916110bf565b60200260200101818152505060027f2b5f494ed674235b8ac1750bdfd5a7615f002d4a1dcefeddd06eda5a076ccd0d7f2fe520ad2020aab9cbba817fcbb9a863b8a76ff88f14f912c5e71665b2ad5e826040516020016107fb929190918252602082015260400190565b60408051601f1981840301815290829052610815916110ad565b602060405180830381855afa158015610830573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906108539190611069565b81600481518110610866576108666110bf565b60200260200101818152505060027f0f1c3c0d5d9da0fa03666843cde4e82e869ba5252fce3c25d5940320b1c4d4937f214bfcff74f425f6fe8c0d07b307482d8bc8bb2f3608f68287aa01bd0b69e8096040516020016108d0929190918252602082015260400190565b60408051601f19818403018152908290526108ea916110ad565b602060405180830381855afa158015610905573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906109289190611069565b8160058151811061093b5761093b6110bf565b60200260200101818152505060028060405161097a907f72697363305f67726f746831362e566572696679696e674b65790000000000008152601a0190565b602060405180830381855afa158015610995573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906109b89190611069565b60027f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e27f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d1926604051602001610a16929190918252602082015260400190565b60408051601f1981840301815290829052610a30916110ad565b602060405180830381855afa158015610a4b573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610a6e9190611069565b604080517f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c60208201527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab918101919091527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a760608201527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec8608082015260029060a00160408051601f1981840301815290829052610b2c916110ad565b602060405180830381855afa158015610b47573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610b6a9190611069565b604080517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c260208201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed918101919091527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60608201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa608082015260029060a00160408051601f1981840301815290829052610c28916110ad565b602060405180830381855afa158015610c43573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610c669190611069565b604080517f03b03cd5effa95ac9bee94f1f5ef907157bda4812ccf0b4c91f42bb629f83a1c60208201527f1aa085ff28179a12d922dba0547057ccaae94b9d69cfaa4e60401fea7f3e0333918101919091527f110c10134f200b19f6490846d518c9aea868366efb7228ca5c91d2940d03076260608201527f1e60f31fcbf757e837e867178318832d0b2d74d59e2fea1c7142df187d3fc6d3608082015260029060a00160408051601f1981840301815290829052610d24916110ad565b602060405180830381855afa158015610d3f573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d629190611069565b610ddb6002604051610d97907f72697363305f67726f746831362e566572696679696e674b65792e49430000008152601d0190565b602060405180830381855afa158015610db2573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610dd59190611069565b88610e6f565b6040805160208101979097528601949094526060850192909252608084015260a083015260c0820152600560f81b60e082015260e20160408051601f1981840301815290829052610e2b916110ad565b602060405180830381855afa158015610e46573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610e699190611069565b91505090565b5f80805b8351811015610ec857610ebe85858360018851610e9091906110d3565b610e9a91906110d3565b81518110610eaa57610eaa6110bf565b602002602001015184610ed260201b60201c565b9150600101610e73565b5090505b92915050565b6040805160028082526060820183525f928392919060208301908036833701905050905083815f81518110610f0957610f096110bf565b6020026020010181815250508281600181518110610f2957610f296110bf565b6020908102919091010152610f3e8582610f47565b95945050505050565b604080515f80825260208201909252610f61848483610f69565b949350505050565b5f5f610f7b845161101160201b60201c565b90505f60088261ffff16901c60088361ffff16901b1760f01b9050600286868684604051602001610faf94939291906110f2565b60408051601f1981840301815290829052610fc9916110ad565b602060405180830381855afa158015610fe4573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906110079190611069565b9695505050505050565b5f61ffff821115611043576040516306dfcc6560e41b8152601060048201526024810183905260440160405180910390fd5b5090565b5f5f60408385031215611058575f5ffd5b505080516020909101519092909150565b5f60208284031215611079575f5ffd5b5051919050565b5f81515f5b8181101561109f5760208185018101518683015201611085565b505f93019283525090919050565b5f6110b88284611080565b9392505050565b634e487b7160e01b5f52603260045260245ffd5b81810381811115610ecc57634e487b7160e01b5f52601160045260245ffd5b8481525f602082018551602087015f5b82811015611120578151845260209384019390910190600101611102565b50505061112d8186611080565b6001600160f01b031994909416845250506002909101949350505050565b60805160a05160c05160e0516112d161119f5f395f8181608e01528181610800015261086401525f818160e8015261093b01525f8181610140015261090201525f818161018001526108da01526112d15ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80638989fa2e116100585780638989fa2e1461013b5780639181e4b11461017b578063ab750e75146101a2578063ffa1ad74146101b5575f5ffd5b8063053c238d146100895780631599ead5146100ce578063258038e2146100e357806334baeab914610118575b5f5ffd5b6100b07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160e01b031990911681526020015b60405180910390f35b6100e16100dc366004610e14565b6101eb565b005b61010a7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100c5565b61012b610126366004610e62565b610205565b60405190151581526020016100c5565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160801b031990911681526020016100c5565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6100e16101b0366004610ec4565b6107bc565b6101de6040518060400160405280600a815260200169352e302e302d72632e3160b01b81525081565b6040516100c59190610f60565b6102026101f88280610f92565b83602001356107de565b50565b5f61075a565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018110610202575f5f5260205ff35b5f60405183815284602082015285604082015260408160608360076107d05a03fa91508161026a575f5f5260205ff35b825160408201526020830151606082015260408360808360066107d05a03fa91505080610299575f5f5260205ff35b5050505050565b7f12ac9a25dcd5e1a832a9061a082c15dd1d61aa9c4d553505739d0f5d65dc3be485527f025aa744581ebe7ad91731911c898569106ff5a2d30f3eee2b23c60ee980acd460208601525f608086018661033c87357f2e32a094b7589554f7bc357bf63481acd2d55555c203383782a4650787ff66427f0707b920bc978c02f292fae2036e057be54294114ccc3c8769d883f688a1423f8461023a565b61038c60208801357f2ede7c9acf48cf3a3729fa3d68714e2a8435d4fa6db8f7f409c153b1fcdf9b8b7f0bca36e2cbe6394b3e249751853f961511011c7148e336f4fd974644850fc3478461023a565b6103dc60408801357f2b91a26aa92e1b6f5722949f192a81c850d586d81a60157f3e9cf04f679cccd67f1b8af999dbfbb3927c091cc2aaf201e488cbacc3e2c6b6fb5a25f9112e04f2a78461023a565b61042c60608801357f2fe520ad2020aab9cbba817fcbb9a863b8a76ff88f14f912c5e71665b2ad5e827f2b5f494ed674235b8ac1750bdfd5a7615f002d4a1dcefeddd06eda5a076ccd0d8461023a565b61047c60808801357f214bfcff74f425f6fe8c0d07b307482d8bc8bb2f3608f68287aa01bd0b69e8097f0f1c3c0d5d9da0fa03666843cde4e82e869ba5252fce3c25d5940320b1c4d4938461023a565b50823581527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208401357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020820152833560408201526020840135606082015260408401356080820152606084013560a08201527f2d4d9aa7e302d9df41749d5507949d05dbea33fbb16c643b22f599a2be6df2e260c08201527f14bedd503c37ceb061d8ec60209fe345ce89830a19230301f076caff004d192660e08201527f0967032fcbf776d1afc985f88877f182d38480a653f2decaa9794cbc3bf3060c6101008201527f0e187847ad4c798374d0d6732bf501847dd68bc0e071241e0213bc7fc13db7ab6101208201527f304cfbd1e08a704a99f5e847d93f8c3caafddec46b7a0d379da69a4d112346a76101408201527f1739c1b1a457a8c7313123d24d2f9192f896b7c63eea05a9d57f06547ad0cec86101608201525f87015161018082015260205f018701516101a08201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08201527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008201527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220820152843561024082015260208501356102608201527f03b03cd5effa95ac9bee94f1f5ef907157bda4812ccf0b4c91f42bb629f83a1c6102808201527f1aa085ff28179a12d922dba0547057ccaae94b9d69cfaa4e60401fea7f3e03336102a08201527f110c10134f200b19f6490846d518c9aea868366efb7228ca5c91d2940d0307626102c08201527f1e60f31fcbf757e837e867178318832d0b2d74d59e2fea1c7142df187d3fc6d36102e08201526020816103008360086107d05a03fa9051169695505050505050565b60405161038081016040526107715f84013561020b565b61077e602084013561020b565b61078b604084013561020b565b610798606084013561020b565b6107a5608084013561020b565b6107b2818486888a6102a0565b9050805f5260205ff35b6107d884846107d36107ce86866109f2565b610a7e565b6107de565b50505050565b6107eb60045f8486610fdc565b6107f491611003565b6001600160e01b0319167f00000000000000000000000000000000000000000000000000000000000000006001600160e01b031916146108975761083b60045f8486610fdc565b61084491611003565b604051632e2ce35360e21b81526001600160e01b031991821660048201527f0000000000000000000000000000000000000000000000000000000000000000909116602482015260440160405180910390fd5b5f5f6108a283610bd8565b90925090505f6108b58560048189610fdc565b8101906108c291906110f8565b8051602080830151604080850151815160a0810183527f0000000000000000000000000000000000000000000000000000000000000000608090811c82527f0000000000000000000000000000000000000000000000000000000000000000811c9582019590955289851c8184015288851c60608201527f00000000000000000000000000000000000000000000000000000000000000009481019490945290516334baeab960e01b81529495505f9430946334baeab99461098a94919391926004016111b3565b602060405180830381865afa1580156109a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109c99190611236565b9050806109e95760405163439cc0cd60e01b815260040160405180910390fd5b50505050505050565b6109fa610dd2565b6040805160a0810182528481527fa3acc27117418996340b84e5a90f3ef4c49d22c79e44aad822ec9c313e1eb8e26020820152815180830183529091820190805f81526020015f60ff1681525081526020015f5f1b8152602001610a7360405180604001604052808681526020015f5f1b815250610d3d565b905290505b92915050565b5f600280604051610aa7907172697363302e52656365697074436c61696d60701b815260120190565b602060405180830381855afa158015610ac2573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610ae59190611269565b60608401518451602086015160808701516040880151516018906002811115610b1057610b10611255565b60408a810151602090810151825191820199909952908101969096526060860194909452608085019290925260a084015263ffffffff909116901b60e01b6001600160e01b03191660c082015260f89190911b6001600160f81b03191660c4820152600160fa1b60c882015260ca015b60408051601f1981840301815290829052610b9a91611280565b602060405180830381855afa158015610bb5573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610a789190611269565b5f8080610d23845f8190506008817eff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff16901b6008827fff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff0016901c1790506010817dffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff16901b6010827fffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff0000ffff000016901c1790506020817bffffffff00000000ffffffff00000000ffffffff00000000ffffffff16901b6020827fffffffff00000000ffffffff00000000ffffffff00000000ffffffff0000000016901c17905060408177ffffffffffffffff0000000000000000ffffffffffffffff16901b60408277ffffffffffffffff0000000000000000ffffffffffffffff1916901c179050608081901b608082901c179050919050565b608081901b956001600160801b0319909116945092505050565b5f600280604051610d60906b1c9a5cd8cc0b93dd5d1c1d5d60a21b8152600c0190565b602060405180830381855afa158015610d7b573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d9e9190611269565b83516020808601516040805192830194909452928101919091526060810191909152600160f91b6080820152608201610b80565b6040518060a001604052805f81526020015f8152602001610e02604080518082019091525f808252602082015290565b81526020015f81526020015f81525090565b5f60208284031215610e24575f5ffd5b813567ffffffffffffffff811115610e3a575f5ffd5b820160408185031215610e4b575f5ffd5b9392505050565b8060408101831015610a78575f5ffd5b5f5f5f5f6101a08587031215610e76575f5ffd5b610e808686610e52565b935060c0850186811115610e92575f5ffd5b604086019350610ea28782610e52565b925050856101a086011115610eb5575f5ffd5b50919490935090916101000190565b5f5f5f5f60608587031215610ed7575f5ffd5b843567ffffffffffffffff811115610eed575f5ffd5b8501601f81018713610efd575f5ffd5b803567ffffffffffffffff811115610f13575f5ffd5b876020828401011115610f24575f5ffd5b602091820198909750908601359560400135945092505050565b5f5b83811015610f58578181015183820152602001610f40565b50505f910152565b602081525f8251806020840152610f7e816040850160208701610f3e565b601f01601f19169190910160400192915050565b5f5f8335601e19843603018112610fa7575f5ffd5b83018035915067ffffffffffffffff821115610fc1575f5ffd5b602001915036819003821315610fd5575f5ffd5b9250929050565b5f5f85851115610fea575f5ffd5b83861115610ff6575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015611034576001600160e01b0319600485900360031b81901b82161691505b5092915050565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156110725761107261103b565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156110a1576110a161103b565b604052919050565b5f82601f8301126110b8575f5ffd5b6110c26040611078565b8060408401858111156110d3575f5ffd5b845b818110156110ed5780358452602093840193016110d5565b509095945050505050565b5f61010082840312801561110a575f5ffd5b5061111361104f565b61111d84846110a9565b815283605f84011261112d575f5ffd5b604061113881611078565b8060c0860187811115611149575f5ffd5b604087015b8181101561116f5761116089826110a9565b8452602090930192840161114e565b5081602086015261118088826110a9565b604086015250929695505050505050565b805f5b60028110156107d8578151845260209384019390910190600101611194565b6101a081016111c28287611191565b60408201855f5b60028110156111f3576111dd838351611191565b60409290920191602091909101906001016111c9565b50505061120360c0830185611191565b6101008201835f5b600581101561122a57815183526020928301929091019060010161120b565b50505095945050505050565b5f60208284031215611246575f5ffd5b81518015158114610e4b575f5ffd5b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215611279575f5ffd5b5051919050565b5f8251611291818460208701610f3e565b919091019291505056fea2646970667358221220d1a0e9d547772bdc446754c9f8e19fc6f62a02c6f14c839524e26dca861a85bf64736f6c634300081e0033",
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
