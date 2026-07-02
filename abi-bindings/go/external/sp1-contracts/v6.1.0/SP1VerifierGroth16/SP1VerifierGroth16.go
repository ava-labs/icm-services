// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sp1verifiergroth16

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

// SP1VerifierGroth16MetaData contains all meta data concerning the SP1VerifierGroth16 contract.
var SP1VerifierGroth16MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidExitCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVkRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"received\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"}],\"name\":\"WrongVerifierSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VERIFIER_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VK_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"}],\"name\":\"hashPublicValues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[5]\",\"name\":\"input\",\"type\":\"uint256[5]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[5]\",\"name\":\"input\",\"type\":\"uint256[5]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611a488061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c80636b61d8e7116100585780636b61d8e7146101065780637cad4e1314610119578063a67086041461013e578063ffa1ad7414610151575f5ffd5b80632a07d99a146100895780632a5104361461009e57806341493c60146100d357806344f63692146100e6575b5f5ffd5b61009c610097366004611621565b610179565b005b7f4388a21c687fdd5f218d7e3d13190cac4c5355818d3605fd5fb811df468ee6965b6040519081526020015b60405180910390f35b61009c6100e136600461169b565b6103f7565b6100f96100f4366004611714565b610594565b6040516100ca9190611736565b6100c0610114366004611766565b6105f1565b7e2f850ee998974d6cc00e50cd0814b098c05bfade466d28573240d057f253526100c0565b61009c61014c3660046117a5565b610656565b6040805180820182526006815265076362e312e360d41b602082015290516100ca91906117de565b5f5f61018483610970565b915091505f6040516101008682377f11b7e9276171bb0efd647fc63e38bbfba3076f20daca8cd52bcc7284d9b1c6eb6101008201527f1723616533dd6ae53502c9c506a81f23f543d68750b5133ebfbe1f4746b3b0116101208201527f241dc9f1df2ecb0c7e6d9522e0d83fe29cb5f4af6571c91c49ea04640c0492216101408201527f2addf202c6b971af36fbafa29888d52e2d13e3e95ea7513d6436f0c1f256f53b6101608201527f21c7d728a5fd961fc179ec5eab938f564deba5b271e1c90c2c29a79648418fc16101808201527f2317bd9b644830ef5ffec1c2d46b5442aa8029d41eb58fb8fc939fb03365e00e6101a08201527f1c3c9339849225980c7d3f824f80d19e2a9c2554b6ab2160fa9635528f693fc06101c08201527f0d964538da2653f2e62499571e6c78afb8909d3ea8107f306bd6928253680a3a6101e08201527f03d612368fddb0f3f83497dbc0ae760f98978cff3af02fdee725fb3562839ae96102008201527f09ec9e44c7989c61a534b716d1d6db028893c68a9934b17f3a50af2e56d1cc7d61022082015283610240820152826102608201527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102808201527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102a08201527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec6102c08201527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d6102e08201526020816103008360085afa9051169050806103f057604051631ff3747d60e21b815260040160405180910390fd5b5050505050565b5f6104056004828486611813565b61040e9161183a565b90507f4388a21c687fdd5f218d7e3d13190cac4c5355818d3605fd5fb811df468ee6966310e2288760e21b6001600160e01b031983161461047a5760405163988066a160e01b81526001600160e01b031980841660048301528216602482015260440160405180910390fd5b7e2f850ee998974d6cc00e50cd0814b098c05bfade466d28573240d057f253525f6104a588886105f1565b90505f8080806104b8896004818d611813565b8101906104c59190611884565b93509350935093506104d5611587565b8d815260208101869052604081018590526060810184905260808101839052841561051357604051631fcf917760e01b815260040160405180910390fd5b86841461053357604051631ab15d8b60e31b815260040160405180910390fd5b604051631503eccd60e11b81523090632a07d99a90610558908590859060040161191d565b5f6040518083038186803b15801561056e575f5ffd5b505afa158015610580573d5f5f3e3d5ffd5b505050505050505050505050505050505050565b61059c6115a5565b6105af82358360015b6020020135610c91565b81526105cd6060830135604084013560a08501356080860135610d7d565b602083015260408201526105e760c08301358360076105a5565b6060820152919050565b5f6001600160fd1b035f1b6002848460405161060e929190611979565b602060405180830381855afa158015610629573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061064c9190611988565b1690505b92915050565b61065e6115c3565b5f8061067085825b6020020135611056565b90925090505f80808061068b60408a013560208b01356110f4565b929650909450925090505f806106a28b6003610666565b915091505f5f6106b18c610970565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f11b7e9276171bb0efd647fc63e38bbfba3076f20daca8cd52bcc7284d9b1c6eb6101008e01527f1723616533dd6ae53502c9c506a81f23f543d68750b5133ebfbe1f4746b3b0116101208e01527f241dc9f1df2ecb0c7e6d9522e0d83fe29cb5f4af6571c91c49ea04640c0492216101408e01527f2addf202c6b971af36fbafa29888d52e2d13e3e95ea7513d6436f0c1f256f53b6101608e01527f21c7d728a5fd961fc179ec5eab938f564deba5b271e1c90c2c29a79648418fc16101808e01527f2317bd9b644830ef5ffec1c2d46b5442aa8029d41eb58fb8fc939fb03365e00e6101a08e01527f1c3c9339849225980c7d3f824f80d19e2a9c2554b6ab2160fa9635528f693fc06101c08e01527f0d964538da2653f2e62499571e6c78afb8909d3ea8107f306bd6928253680a3a6101e08e01527f03d612368fddb0f3f83497dbc0ae760f98978cff3af02fdee725fb3562839ae96102008e01527f09ec9e44c7989c61a534b716d1d6db028893c68a9934b17f3a50af2e56d1cc7d6102208e01526102408d018290526102608d018190527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26102808e01527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6102a08e01527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec6102c08e01527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d6102e08e015290925090505f6109256115e2565b6020816103008f60085afa915081158061094157508051600114155b1561095f57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050505050505050565b5f5f5f60019050604051604081015f7f2cd6bf7f164af0b6b0bbbe0fdcb06ee0c1ba07f8e6eb2f9f3943a90cb1d4029083527f0a76f7e6e8a78d71649eb7e3132018885e2f57ccca64394738e746c63c9fa83b60208401527f0f5460f3b7221705435e745da21e276536379c0113c13c4255e7ae101f1e90bf82527f022b4d6c4b7dbdeed528d5abe6eddd6f2bd4ae4b9b0450dc18901fc727da799c6020830152863590508060408301525f5160206119f35f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0b0ae6e491bc04c544da9e8cd4857d201b4cfa0222dbe96aac97f044fdf1c92282527f123e064d2ffdb7f8094b727cde1c0beea6d4d551e5743a336804f8effb7f07ac6020830152602087013590508060408301525f5160206119f35f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f097c875a6ebd0999b06e7267ff3d8a6bf859bb9635abae07cb6b3534ba409a8382527f2d84b73fb0f397e356b524f2d0e9ddc9d0f6ac75e566041147ee8554879183ec6020830152604087013590508060408301525f5160206119f35f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1807204ddcd27506ba72e17b55227b0bf310136ecb40c74acd52f3ccfbcba9f782527f011078b5c7767fe73bb54241dcd143ee805ea4383d48d46bced129482c8c20216020830152606087013590508060408301525f5160206119f35f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7e8c7b7c98d78c07a2c4be5f6be7082ba41021611f9a2dfc016f8bbb37d36bee83527f09ce6d246baeac24b5a404ca175c2cc4a974e217176fe2ec2a6c22eb60fd65c16020840152608088013560408085018290525f5160206119f35f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa81516020909201519194509092501680610c8b5760405163a54f8e2760e01b815260040160405180910390fd5b50915091565b5f5f5160206119d35f395f51905f5283101580610cbb57505f5160206119d35f395f51905f528210155b15610cd957604051631ff3747d60e21b815260040160405180910390fd5b82158015610ce5575081155b15610cf157505f610650565b5f610d2c5f5160206119d35f395f51905f5260035f5160206119d35f395f51905f52875f5160206119d35f395f51905f52898a0909086112c8565b9050808303610d41575050600182901b610650565b610d4a8161132a565b8303610d5d575050600182811b17610650565b604051631ff3747d60e21b815260040160405180910390fd5b5092915050565b5f5f5f5160206119d35f395f51905f5286101580610da857505f5160206119d35f395f51905f528510155b80610dc057505f5160206119d35f395f51905f528410155b80610dd857505f5160206119d35f395f51905f528310155b15610df657604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610e0d57505f90508061104d565b5f80805f5160206119d35f395f51905f52610e3660035f5160206119d35f395f51905f526119b3565b5f5160206119d35f395f51905f528a8c090990505f5f5160206119d35f395f51905f528a5f5160206119d35f395f51905f528c8d090990505f5f5160206119d35f395f51905f528a5f5160206119d35f395f51905f528c8d090990505f5160206119d35f395f51905f52805f5160206119d35f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610f245f5160206119d35f395f51905f52805f5160206119d35f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861132a565b93505050505f5f610f715f5160206119d35f395f51905f5280610f4957610f4961199f565b5f5160206119d35f395f51905f528586095f5160206119d35f395f51905f52878809086112c8565b9050610fbc5f5160206119d35f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206119d35f395f51905f5284880809611342565b15915050610fcb83838361138a565b90935091508683148015610fde57508186145b156110065780610fee575f610ff1565b60025b60ff1660028a901b175f179450879350611049565b61100f8361132a565b8714801561102457506110218261132a565b86145b15610d5d5780611034575f611037565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b5f5f825f0361106957505f928392509050565b600183811c9250808416145f5160206119d35f395f51905f5283106110a157604051631ff3747d60e21b815260040160405180910390fd5b6110db5f5160206119d35f395f51905f5260035f5160206119d35f395f51905f52865f5160206119d35f395f51905f5288890909086112c8565b91508015610c8b576110ec8261132a565b915050915091565b5f80808085158015611104575084155b1561111957505f9250829150819050806112bf565b600286811c945085935060018088161490808816145f5160206119d35f395f51905f528610158061115757505f5160206119d35f395f51905f528510155b1561117557604051631ff3747d60e21b815260040160405180910390fd5b5f5f5160206119d35f395f51905f5261119c60035f5160206119d35f395f51905f526119b3565b5f5160206119d35f395f51905f52888a090990505f5f5160206119d35f395f51905f52885f5160206119d35f395f51905f528a8b090990505f5f5160206119d35f395f51905f52885f5160206119d35f395f51905f528a8b090990505f5160206119d35f395f51905f52805f5160206119d35f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061128a5f5160206119d35f395f51905f52805f5160206119d35f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861132a565b955061129787878661138a565b909750955084156112b9576112ab8761132a565b96506112b68661132a565b95505b50505050505b92959194509250565b5f6112f3827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526114c6565b9050815f5160206119d35f395f51905f528283091461132557604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f5160206119d35f395f51905f529081900681030690565b5f5f61136e837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526114c6565b9050825f5160206119d35f395f51905f52828309149392505050565b5f80806113b95f5160206119d35f395f51905f52808788095f5160206119d35f395f51905f52898a09086112c8565b905083156113cd576113ca8161132a565b90505b6114165f5160206119d35f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206119d35f395f51905f52848a08096112c8565b92505f5160206119d35f395f51905f526114405f5160206119d35f395f51905f5260028609611529565b860991505f5160206119d35f395f51905f5261146b5f5160206119d35f395f51905f5284850961132a565b5f5160206119d35f395f51905f52858609088614158061149f57505f5160206119d35f395f51905f52808385096002098514155b156114bd57604051631ff3747d60e21b815260040160405180910390fd5b50935093915050565b5f5f60405160208152602080820152602060408201528460608201528360808201525f5160206119d35f395f51905f5260a082015260208160c08360055afa90519250905080610d7657604051631ff3747d60e21b815260040160405180910390fd5b5f611554827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456114c6565b90505f5160206119d35f395f51905f5281830960011461132557604051631ff3747d60e21b815260040160405180910390fd5b6040518060a001604052806005906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b806101008101831015610650575f5ffd5b8060a08101831015610650575f5ffd5b5f5f6101a08385031215611633575f5ffd5b61163d8484611600565b915061164d846101008501611611565b90509250929050565b5f5f83601f840112611666575f5ffd5b50813567ffffffffffffffff81111561167d575f5ffd5b602083019150836020828501011115611694575f5ffd5b9250929050565b5f5f5f5f5f606086880312156116af575f5ffd5b85359450602086013567ffffffffffffffff8111156116cc575f5ffd5b6116d888828901611656565b909550935050604086013567ffffffffffffffff8111156116f7575f5ffd5b61170388828901611656565b969995985093965092949392505050565b5f6101008284031215611725575f5ffd5b61172f8383611600565b9392505050565b6080810181835f5b600481101561175d57815183526020928301929091019060010161173e565b50505092915050565b5f5f60208385031215611777575f5ffd5b823567ffffffffffffffff81111561178d575f5ffd5b61179985828601611656565b90969095509350505050565b5f5f61012083850312156117b7575f5ffd5b60808301848111156117c7575f5ffd5b8392506117d48582611611565b9150509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f85851115611821575f5ffd5b8386111561182d575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610d76576001600160e01b031960049490940360031b84901b1690921692915050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f6101608587031215611898575f5ffd5b843593506020850135925060408501359150607f850186136118b8575f5ffd5b604051610100810181811067ffffffffffffffff821117156118dc576118dc611870565b604052806101608701888111156118f1575f5ffd5b606088015b8181101561190e5780358352602092830192016118f6565b50959894975092955093505050565b6101a0810181845f5b6008811015611945578151835260209283019290910190600101611926565b5050506101008201835f5b600581101561196f578151835260209283019290910190600101611950565b5050509392505050565b818382375f9101908152919050565b5f60208284031215611998575f5ffd5b5051919050565b634e487b7160e01b5f52601260045260245ffd5b8181038181111561065057634e487b7160e01b5f52601160045260245ffdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122004a72a5eae9107ee0d534f3e4a8d241807f93bda29e8e21f16d7d90db57c6fdc64736f6c634300081e0033",
}

// SP1VerifierGroth16ABI is the input ABI used to generate the binding from.
// Deprecated: Use SP1VerifierGroth16MetaData.ABI instead.
var SP1VerifierGroth16ABI = SP1VerifierGroth16MetaData.ABI

// SP1VerifierGroth16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SP1VerifierGroth16MetaData.Bin instead.
var SP1VerifierGroth16Bin = SP1VerifierGroth16MetaData.Bin

// DeploySP1VerifierGroth16 deploys a new Ethereum contract, binding an instance of SP1VerifierGroth16 to it.
func DeploySP1VerifierGroth16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SP1VerifierGroth16, error) {
	parsed, err := SP1VerifierGroth16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SP1VerifierGroth16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SP1VerifierGroth16{SP1VerifierGroth16Caller: SP1VerifierGroth16Caller{contract: contract}, SP1VerifierGroth16Transactor: SP1VerifierGroth16Transactor{contract: contract}, SP1VerifierGroth16Filterer: SP1VerifierGroth16Filterer{contract: contract}}, nil
}

// SP1VerifierGroth16 is an auto generated Go binding around an Ethereum contract.
type SP1VerifierGroth16 struct {
	SP1VerifierGroth16Caller     // Read-only binding to the contract
	SP1VerifierGroth16Transactor // Write-only binding to the contract
	SP1VerifierGroth16Filterer   // Log filterer for contract events
}

// SP1VerifierGroth16Caller is an auto generated read-only Go binding around an Ethereum contract.
type SP1VerifierGroth16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SP1VerifierGroth16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SP1VerifierGroth16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SP1VerifierGroth16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SP1VerifierGroth16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SP1VerifierGroth16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SP1VerifierGroth16Session struct {
	Contract     *SP1VerifierGroth16 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SP1VerifierGroth16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SP1VerifierGroth16CallerSession struct {
	Contract *SP1VerifierGroth16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SP1VerifierGroth16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SP1VerifierGroth16TransactorSession struct {
	Contract     *SP1VerifierGroth16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SP1VerifierGroth16Raw is an auto generated low-level Go binding around an Ethereum contract.
type SP1VerifierGroth16Raw struct {
	Contract *SP1VerifierGroth16 // Generic contract binding to access the raw methods on
}

// SP1VerifierGroth16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SP1VerifierGroth16CallerRaw struct {
	Contract *SP1VerifierGroth16Caller // Generic read-only contract binding to access the raw methods on
}

// SP1VerifierGroth16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SP1VerifierGroth16TransactorRaw struct {
	Contract *SP1VerifierGroth16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSP1VerifierGroth16 creates a new instance of SP1VerifierGroth16, bound to a specific deployed contract.
func NewSP1VerifierGroth16(address common.Address, backend bind.ContractBackend) (*SP1VerifierGroth16, error) {
	contract, err := bindSP1VerifierGroth16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SP1VerifierGroth16{SP1VerifierGroth16Caller: SP1VerifierGroth16Caller{contract: contract}, SP1VerifierGroth16Transactor: SP1VerifierGroth16Transactor{contract: contract}, SP1VerifierGroth16Filterer: SP1VerifierGroth16Filterer{contract: contract}}, nil
}

// NewSP1VerifierGroth16Caller creates a new read-only instance of SP1VerifierGroth16, bound to a specific deployed contract.
func NewSP1VerifierGroth16Caller(address common.Address, caller bind.ContractCaller) (*SP1VerifierGroth16Caller, error) {
	contract, err := bindSP1VerifierGroth16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SP1VerifierGroth16Caller{contract: contract}, nil
}

// NewSP1VerifierGroth16Transactor creates a new write-only instance of SP1VerifierGroth16, bound to a specific deployed contract.
func NewSP1VerifierGroth16Transactor(address common.Address, transactor bind.ContractTransactor) (*SP1VerifierGroth16Transactor, error) {
	contract, err := bindSP1VerifierGroth16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SP1VerifierGroth16Transactor{contract: contract}, nil
}

// NewSP1VerifierGroth16Filterer creates a new log filterer instance of SP1VerifierGroth16, bound to a specific deployed contract.
func NewSP1VerifierGroth16Filterer(address common.Address, filterer bind.ContractFilterer) (*SP1VerifierGroth16Filterer, error) {
	contract, err := bindSP1VerifierGroth16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SP1VerifierGroth16Filterer{contract: contract}, nil
}

// bindSP1VerifierGroth16 binds a generic wrapper to an already deployed contract.
func bindSP1VerifierGroth16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SP1VerifierGroth16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SP1VerifierGroth16 *SP1VerifierGroth16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SP1VerifierGroth16.Contract.SP1VerifierGroth16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SP1VerifierGroth16 *SP1VerifierGroth16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SP1VerifierGroth16.Contract.SP1VerifierGroth16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SP1VerifierGroth16 *SP1VerifierGroth16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SP1VerifierGroth16.Contract.SP1VerifierGroth16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SP1VerifierGroth16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SP1VerifierGroth16 *SP1VerifierGroth16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SP1VerifierGroth16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SP1VerifierGroth16 *SP1VerifierGroth16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SP1VerifierGroth16.Contract.contract.Transact(opts, method, params...)
}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VERIFIERHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "VERIFIER_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VERIFIERHASH() ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.VERIFIERHASH(&_SP1VerifierGroth16.CallOpts)
}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VERIFIERHASH() ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.VERIFIERHASH(&_SP1VerifierGroth16.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VERSION() (string, error) {
	return _SP1VerifierGroth16.Contract.VERSION(&_SP1VerifierGroth16.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VERSION() (string, error) {
	return _SP1VerifierGroth16.Contract.VERSION(&_SP1VerifierGroth16.CallOpts)
}

// VKROOT is a free data retrieval call binding the contract method 0x7cad4e13.
//
// Solidity: function VK_ROOT() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VKROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "VK_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VKROOT is a free data retrieval call binding the contract method 0x7cad4e13.
//
// Solidity: function VK_ROOT() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VKROOT() ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.VKROOT(&_SP1VerifierGroth16.CallOpts)
}

// VKROOT is a free data retrieval call binding the contract method 0x7cad4e13.
//
// Solidity: function VK_ROOT() pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VKROOT() ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.VKROOT(&_SP1VerifierGroth16.CallOpts)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _SP1VerifierGroth16.Contract.CompressProof(&_SP1VerifierGroth16.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _SP1VerifierGroth16.Contract.CompressProof(&_SP1VerifierGroth16.CallOpts, proof)
}

// HashPublicValues is a free data retrieval call binding the contract method 0x6b61d8e7.
//
// Solidity: function hashPublicValues(bytes publicValues) pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) HashPublicValues(opts *bind.CallOpts, publicValues []byte) ([32]byte, error) {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "hashPublicValues", publicValues)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashPublicValues is a free data retrieval call binding the contract method 0x6b61d8e7.
//
// Solidity: function hashPublicValues(bytes publicValues) pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) HashPublicValues(publicValues []byte) ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.HashPublicValues(&_SP1VerifierGroth16.CallOpts, publicValues)
}

// HashPublicValues is a free data retrieval call binding the contract method 0x6b61d8e7.
//
// Solidity: function hashPublicValues(bytes publicValues) pure returns(bytes32)
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) HashPublicValues(publicValues []byte) ([32]byte, error) {
	return _SP1VerifierGroth16.Contract.HashPublicValues(&_SP1VerifierGroth16.CallOpts, publicValues)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [5]*big.Int) error {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, input [5]*big.Int) error {
	return _SP1VerifierGroth16.Contract.VerifyCompressedProof(&_SP1VerifierGroth16.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xa6708604.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [5]*big.Int) error {
	return _SP1VerifierGroth16.Contract.VerifyCompressedProof(&_SP1VerifierGroth16.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [5]*big.Int) error {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VerifyProof(proof [8]*big.Int, input [5]*big.Int) error {
	return _SP1VerifierGroth16.Contract.VerifyProof(&_SP1VerifierGroth16.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x2a07d99a.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[5] input) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, input [5]*big.Int) error {
	return _SP1VerifierGroth16.Contract.VerifyProof(&_SP1VerifierGroth16.CallOpts, proof, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _SP1VerifierGroth16.contract.Call(opts, &out, "verifyProof0", programVKey, publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16Session) VerifyProof0(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _SP1VerifierGroth16.Contract.VerifyProof0(&_SP1VerifierGroth16.CallOpts, programVKey, publicValues, proofBytes)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_SP1VerifierGroth16 *SP1VerifierGroth16CallerSession) VerifyProof0(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _SP1VerifierGroth16.Contract.VerifyProof0(&_SP1VerifierGroth16.CallOpts, programVKey, publicValues, proofBytes)
}
