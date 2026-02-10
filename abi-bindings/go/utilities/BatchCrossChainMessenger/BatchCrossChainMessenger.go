// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchcrosschainmessenger

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

// BatchCrossChainMessengerMetaData contains all meta data concerning the BatchCrossChainMessenger contract.
var BatchCrossChainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"}],\"name\":\"SendMessages\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessages\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"}],\"name\":\"sendMessages\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161217a38038061217a83398101604081905261002e916105e2565b5f51602061213a5f395f51905f52805468010000000000000000810460ff1615906001600160401b03165f811580156100645750825b90505f826001600160401b0316600114801561007f5750303b155b90508115801561008d575080155b156100ab5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b031916600117855583156100d957845460ff60401b1916680100000000000000001785555b6100e161013f565b6100ec888888610151565b831561013257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050610635565b610147610171565b61014f6101ac565b565b610159610171565b61016383826101da565b61016c82610200565b505050565b5f51602061213a5f395f51905f525468010000000000000000900460ff1661014f57604051631afcd79f60e31b815260040160405180910390fd5b6101b4610171565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6101e2610171565b6101ea61013f565b6101f2610214565b6101fc828261021c565b5050565b610208610171565b61021181610398565b50565b61014f610171565b610224610171565b6001600160a01b0382166102a55760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b5f5f51602061211a5f395f51905f5290505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102f7573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031b919061061e565b116103705760405162461bcd60e51b815260206004820152603260248201525f51602061215a5f395f51905f52604482015271656c65706f7274657220726567697374727960701b606482015260840161029c565b81546001600160a01b0319166001600160a01b038216178255610392836103d2565b50505050565b6103a0610171565b6001600160a01b0381166103c957604051631e4fbdf760e01b81525f600482015260240161029c565b61021181610557565b5f51602061211a5f395f51905f5280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610426573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061044a919061061e565b6002830154909150818411156104a95760405162461bcd60e51b815260206004820152603160248201525f51602061215a5f395f51905f5260448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161029c565b80841161051e5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161029c565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b80516001600160a01b03811681146105dd575f5ffd5b919050565b5f5f5f606084860312156105f4575f5ffd5b6105fd846105c7565b925061060b602085016105c7565b6040949094015192959394509192915050565b5f6020828403121561062e575f5ffd5b5051919050565b611ad8806106425f395ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c8063909a6ac01161006e578063909a6ac01461015b578063973142971461017d578063c1329fcb146101a0578063c868efaa146101c0578063d2cc7a70146101d3578063f2fde38b146101fa575f5ffd5b80632b0d8f18146100b55780633902970c146100ca5780634511243e146100f35780635eb9951414610106578063715018a6146101195780638da5cb5b14610121575b5f5ffd5b6100c86100c336600461137d565b61020d565b005b6100dd6100d8366004611404565b61030f565b6040516100ea9190611555565b60405180910390f35b6100c861010136600461137d565b6104fd565b6100c8610114366004611597565b6105ec565b6100c8610600565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020016100ea565b61016f5f516020611a835f395f51905f5281565b6040519081526020016100ea565b61019061018b36600461137d565b610613565b60405190151581526020016100ea565b6101b36101ae366004611597565b610633565b6040516100ea9190611636565b6100c86101ce366004611648565b61070f565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d025461016f565b6100c861020836600461137d565b6108ed565b5f516020611a835f395f51905f52610223610927565b6001600160a01b0382166102525760405162461bcd60e51b8152600401610249906116cd565b60405180910390fd5b61025c818361092f565b156102bf5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610249565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b6060610319610953565b5f841561032d5761032a868661099d565b90505b866001600160a01b0316887f430d1906813fdb2129a19139f4112a1396804605501a798df3a4042590ba20d58884888860405161036d949392919061171b565b60405180910390a35f835167ffffffffffffffff81111561039057610390611398565b6040519080825280602002602001820160405280156103b9578160200160208202803683370190505b5090505f5b84518110156104c6575f61049d6040518060c001604052808d81526020018c6001600160a01b0316815260200160405180604001604052808d6001600160a01b031681526020018881525081526020018981526020015f67ffffffffffffffff81111561042d5761042d611398565b604051908082528060200260200182016040528015610456578160200160208202803683370190505b50815260200188858151811061046e5761046e611747565b6020026020010151604051602001610486919061175b565b6040516020818303038152906040528152506109a9565b9050808383815181106104b2576104b2611747565b6020908102919091010152506001016103be565b509150506104f360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b9695505050505050565b5f516020611a835f395f51905f52610513610927565b6001600160a01b0382166105395760405162461bcd60e51b8152600401610249906116cd565b610543818361092f565b6105a15760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610249565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b6105f4610927565b6105fd81610ac4565b50565b610608610c5c565b6106115f610cb7565b565b5f5f516020611a835f395f51905f5261062c818461092f565b9392505050565b5f81815260208181526040808320805482518185028101850190935280835260609493849084015b82821015610703578382905f5260205f200180546106789061176d565b80601f01602080910402602001604051908101604052809291908181526020018280546106a49061176d565b80156106ef5780601f106106c6576101008083540402835291602001916106ef565b820191905f5260205f20905b8154815290600101906020018083116106d257829003601f168201915b50505050508152602001906001019061065b565b50929695505050505050565b610717610953565b5f5f516020611a835f395f51905f5260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610782573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a691906117a5565b101561080d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610249565b610817813361092f565b1561087d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610249565b6108bd858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610d2792505050565b506108e760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b6108f5610c5c565b6001600160a01b03811661091e57604051631e4fbdf760e01b81525f6004820152602401610249565b6105fd81610cb7565b610611610c5c565b6001600160a01b0381165f90815260018301602052604090205460ff165b92915050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161099757604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f61062c833384610dad565b5f5f6109b3610f06565b60408401516020015190915015610a58576040830151516001600160a01b0316610a355760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610249565b604083015160208101519051610a58916001600160a01b03909116908390610ff6565b604051630624488560e41b81526001600160a01b03821690636244885090610a849086906004016117ff565b6020604051808303815f875af1158015610aa0573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061062c91906117a5565b5f516020611a835f395f51905f5280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610b18573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b3c91906117a5565b600283015490915081841115610bae5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610249565b808411610c235760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610249565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b33610c8e7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106115760405163118cdaa760e01b8152336004820152602401610249565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f81806020019051810190610d3c919061187d565b5f8581526020818152604082208054600181018255908352912091925001610d64828261193d565b50826001600160a01b0316847f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa83604051610d9f919061175b565b60405180910390a350505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015610df3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e1791906117a5565b9050610e2e6001600160a01b03861685308661107d565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015610e72573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e9691906117a5565b9050818111610efc5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610249565b6104f38282611a0c565b5f516020611a835f395f51905f5280546040805163d820e64f60e01b815290515f939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015610f60573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f849190611a1f565b9050610f90828261092f565b1561094d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610249565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015611043573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106791906117a5565b90506108e784846110788585611a3a565b6110e4565b6040516001600160a01b0384811660248301528381166044830152606482018390526108e79186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b03838183161783525050505061116f565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b17905261113584826111d5565b6108e7576040516001600160a01b0384811660248301525f604483015261116991869182169063095ea7b3906064016110b2565b6108e784825b5f6111836001600160a01b03841683611276565b905080515f141580156111a75750808060200190518101906111a59190611a4d565b155b156111d057604051635274afe760e01b81526001600160a01b0384166004820152602401610249565b505050565b5f5f5f846001600160a01b0316846040516111f09190611a6c565b5f604051808303815f865af19150503d805f8114611229576040519150601f19603f3d011682016040523d82523d5f602084013e61122e565b606091505b50915091508180156112585750805115806112585750808060200190518101906112589190611a4d565b801561126d57505f856001600160a01b03163b115b95945050505050565b606061062c83835f845f5f856001600160a01b0316848660405161129a9190611a6c565b5f6040518083038185875af1925050503d805f81146112d4576040519150601f19603f3d011682016040523d82523d5f602084013e6112d9565b606091505b50915091506104f38683836060826112f9576112f482611340565b61062c565b815115801561131057506001600160a01b0384163b155b1561133957604051639996b31560e01b81526001600160a01b0385166004820152602401610249565b508061062c565b8051156113505780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b03811681146105fd575f5ffd5b5f6020828403121561138d575f5ffd5b813561062c81611369565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156113d5576113d5611398565b604052919050565b5f67ffffffffffffffff8211156113f6576113f6611398565b50601f01601f191660200190565b5f5f5f5f5f5f60c08789031215611419575f5ffd5b86359550602087013561142b81611369565b9450604087013561143b81611369565b9350606087013592506080870135915060a087013567ffffffffffffffff811115611464575f5ffd5b8701601f81018913611474575f5ffd5b803567ffffffffffffffff81111561148e5761148e611398565b8060051b61149e602082016113ac565b9182526020818401810192908101908c8411156114b9575f5ffd5b6020850192505b8383101561154357823567ffffffffffffffff8111156114de575f5ffd5b8501603f81018e136114ee575f5ffd5b60208101356115046114ff826113dd565b6113ac565b8181528f6020808486010101111561151a575f5ffd5b816040840160208301375f602083830101528085525050506020820191506020830192506114c0565b80955050505050509295509295509295565b602080825282518282018190525f918401906040840190835b8181101561158c57835183526020938401939092019160010161156e565b509095945050505050565b5f602082840312156115a7575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208501945060208160051b830101602085015f5b8381101561162a57601f198584030188526116148383516115ae565b60209889019890935091909101906001016115f8565b50909695505050505050565b602081525f61062c60208301846115dc565b5f5f5f5f6060858703121561165b575f5ffd5b84359350602085013561166d81611369565b9250604085013567ffffffffffffffff811115611688575f5ffd5b8501601f81018713611698575f5ffd5b803567ffffffffffffffff8111156116ae575f5ffd5b8760208284010111156116bf575f5ffd5b949793965060200194505050565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60018060a01b0385168152836020820152826040820152608060608201525f6104f360808301846115dc565b634e487b7160e01b5f52603260045260245ffd5b602081525f61062c60208301846115ae565b600181811c9082168061178157607f821691505b60208210810361179f57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156117b5575f5ffd5b5051919050565b5f8151808452602084019350602083015f5b828110156117f55781516001600160a01b03168652602095860195909101906001016117ce565b5093949350505050565b602081528151602082015260018060a01b0360208301511660408201525f604083015160018060a01b0381511660608401526020810151608084015250606083015160a0830152608083015160e060c08401526118606101008401826117bc565b905060a0840151601f198483030160e085015261126d82826115ae565b5f6020828403121561188d575f5ffd5b815167ffffffffffffffff8111156118a3575f5ffd5b8201601f810184136118b3575f5ffd5b80516118c16114ff826113dd565b8181528560208385010111156118d5575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b601f8211156111d057805f5260205f20601f840160051c810160208510156119175750805b601f840160051c820191505b81811015611936575f8155600101611923565b5050505050565b815167ffffffffffffffff81111561195757611957611398565b61196b81611965845461176d565b846118f2565b6020601f82116001811461199d575f83156119865750848201515b5f19600385901b1c1916600184901b178455611936565b5f84815260208120601f198516915b828110156119cc57878501518255602094850194600190920191016119ac565b50848210156119e957868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561094d5761094d6119f8565b5f60208284031215611a2f575f5ffd5b815161062c81611369565b8082018082111561094d5761094d6119f8565b5f60208284031215611a5d575f5ffd5b8151801515811461062c575f5ffd5b5f82518060208501845e5f92019182525091905056fede77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a2646970667358221220acac88f2a7561440ca4f0b0748b76870a51172cb05101f37fe4d155c1035e7e764736f6c634300081e0033de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054656c65706f7274657252656769737472794170703a20696e76616c69642054",
}

// BatchCrossChainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchCrossChainMessengerMetaData.ABI instead.
var BatchCrossChainMessengerABI = BatchCrossChainMessengerMetaData.ABI

// BatchCrossChainMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchCrossChainMessengerMetaData.Bin instead.
var BatchCrossChainMessengerBin = BatchCrossChainMessengerMetaData.Bin

// DeployBatchCrossChainMessenger deploys a new Ethereum contract, binding an instance of BatchCrossChainMessenger to it.
func DeployBatchCrossChainMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int) (common.Address, *types.Transaction, *BatchCrossChainMessenger, error) {
	parsed, err := BatchCrossChainMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchCrossChainMessengerBin), backend, teleporterRegistryAddress, teleporterManager, minTeleporterVersion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchCrossChainMessenger{BatchCrossChainMessengerCaller: BatchCrossChainMessengerCaller{contract: contract}, BatchCrossChainMessengerTransactor: BatchCrossChainMessengerTransactor{contract: contract}, BatchCrossChainMessengerFilterer: BatchCrossChainMessengerFilterer{contract: contract}}, nil
}

// BatchCrossChainMessenger is an auto generated Go binding around an Ethereum contract.
type BatchCrossChainMessenger struct {
	BatchCrossChainMessengerCaller     // Read-only binding to the contract
	BatchCrossChainMessengerTransactor // Write-only binding to the contract
	BatchCrossChainMessengerFilterer   // Log filterer for contract events
}

// BatchCrossChainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchCrossChainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCrossChainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchCrossChainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCrossChainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchCrossChainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCrossChainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchCrossChainMessengerSession struct {
	Contract     *BatchCrossChainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BatchCrossChainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchCrossChainMessengerCallerSession struct {
	Contract *BatchCrossChainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// BatchCrossChainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchCrossChainMessengerTransactorSession struct {
	Contract     *BatchCrossChainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// BatchCrossChainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchCrossChainMessengerRaw struct {
	Contract *BatchCrossChainMessenger // Generic contract binding to access the raw methods on
}

// BatchCrossChainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchCrossChainMessengerCallerRaw struct {
	Contract *BatchCrossChainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// BatchCrossChainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchCrossChainMessengerTransactorRaw struct {
	Contract *BatchCrossChainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchCrossChainMessenger creates a new instance of BatchCrossChainMessenger, bound to a specific deployed contract.
func NewBatchCrossChainMessenger(address common.Address, backend bind.ContractBackend) (*BatchCrossChainMessenger, error) {
	contract, err := bindBatchCrossChainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessenger{BatchCrossChainMessengerCaller: BatchCrossChainMessengerCaller{contract: contract}, BatchCrossChainMessengerTransactor: BatchCrossChainMessengerTransactor{contract: contract}, BatchCrossChainMessengerFilterer: BatchCrossChainMessengerFilterer{contract: contract}}, nil
}

// NewBatchCrossChainMessengerCaller creates a new read-only instance of BatchCrossChainMessenger, bound to a specific deployed contract.
func NewBatchCrossChainMessengerCaller(address common.Address, caller bind.ContractCaller) (*BatchCrossChainMessengerCaller, error) {
	contract, err := bindBatchCrossChainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerCaller{contract: contract}, nil
}

// NewBatchCrossChainMessengerTransactor creates a new write-only instance of BatchCrossChainMessenger, bound to a specific deployed contract.
func NewBatchCrossChainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchCrossChainMessengerTransactor, error) {
	contract, err := bindBatchCrossChainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerTransactor{contract: contract}, nil
}

// NewBatchCrossChainMessengerFilterer creates a new log filterer instance of BatchCrossChainMessenger, bound to a specific deployed contract.
func NewBatchCrossChainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchCrossChainMessengerFilterer, error) {
	contract, err := bindBatchCrossChainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerFilterer{contract: contract}, nil
}

// bindBatchCrossChainMessenger binds a generic wrapper to an already deployed contract.
func bindBatchCrossChainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchCrossChainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchCrossChainMessenger.Contract.BatchCrossChainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.BatchCrossChainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.BatchCrossChainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchCrossChainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.contract.Transact(opts, method, params...)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BatchCrossChainMessenger.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _BatchCrossChainMessenger.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_BatchCrossChainMessenger.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _BatchCrossChainMessenger.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_BatchCrossChainMessenger.CallOpts)
}

// GetCurrentMessages is a free data retrieval call binding the contract method 0xc1329fcb.
//
// Solidity: function getCurrentMessages(bytes32 sourceBlockchainID) view returns(string[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCaller) GetCurrentMessages(opts *bind.CallOpts, sourceBlockchainID [32]byte) ([]string, error) {
	var out []interface{}
	err := _BatchCrossChainMessenger.contract.Call(opts, &out, "getCurrentMessages", sourceBlockchainID)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCurrentMessages is a free data retrieval call binding the contract method 0xc1329fcb.
//
// Solidity: function getCurrentMessages(bytes32 sourceBlockchainID) view returns(string[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) GetCurrentMessages(sourceBlockchainID [32]byte) ([]string, error) {
	return _BatchCrossChainMessenger.Contract.GetCurrentMessages(&_BatchCrossChainMessenger.CallOpts, sourceBlockchainID)
}

// GetCurrentMessages is a free data retrieval call binding the contract method 0xc1329fcb.
//
// Solidity: function getCurrentMessages(bytes32 sourceBlockchainID) view returns(string[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerSession) GetCurrentMessages(sourceBlockchainID [32]byte) ([]string, error) {
	return _BatchCrossChainMessenger.Contract.GetCurrentMessages(&_BatchCrossChainMessenger.CallOpts, sourceBlockchainID)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BatchCrossChainMessenger.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _BatchCrossChainMessenger.Contract.GetMinTeleporterVersion(&_BatchCrossChainMessenger.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _BatchCrossChainMessenger.Contract.GetMinTeleporterVersion(&_BatchCrossChainMessenger.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BatchCrossChainMessenger.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _BatchCrossChainMessenger.Contract.IsTeleporterAddressPaused(&_BatchCrossChainMessenger.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _BatchCrossChainMessenger.Contract.IsTeleporterAddressPaused(&_BatchCrossChainMessenger.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchCrossChainMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) Owner() (common.Address, error) {
	return _BatchCrossChainMessenger.Contract.Owner(&_BatchCrossChainMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerCallerSession) Owner() (common.Address, error) {
	return _BatchCrossChainMessenger.Contract.Owner(&_BatchCrossChainMessenger.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.PauseTeleporterAddress(&_BatchCrossChainMessenger.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.PauseTeleporterAddress(&_BatchCrossChainMessenger.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_BatchCrossChainMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_BatchCrossChainMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.RenounceOwnership(&_BatchCrossChainMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.RenounceOwnership(&_BatchCrossChainMessenger.TransactOpts)
}

// SendMessages is a paid mutator transaction binding the contract method 0x3902970c.
//
// Solidity: function sendMessages(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages) returns(bytes32[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) SendMessages(opts *bind.TransactOpts, destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, messages []string) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "sendMessages", destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0x3902970c.
//
// Solidity: function sendMessages(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages) returns(bytes32[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) SendMessages(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, messages []string) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.SendMessages(&_BatchCrossChainMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0x3902970c.
//
// Solidity: function sendMessages(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages) returns(bytes32[])
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) SendMessages(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, messages []string) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.SendMessages(&_BatchCrossChainMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, messages)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.TransferOwnership(&_BatchCrossChainMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.TransferOwnership(&_BatchCrossChainMessenger.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.UnpauseTeleporterAddress(&_BatchCrossChainMessenger.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.UnpauseTeleporterAddress(&_BatchCrossChainMessenger.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_BatchCrossChainMessenger.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BatchCrossChainMessenger *BatchCrossChainMessengerTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _BatchCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_BatchCrossChainMessenger.TransactOpts, version)
}

// BatchCrossChainMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerInitializedIterator struct {
	Event *BatchCrossChainMessengerInitialized // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerInitialized)
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
		it.Event = new(BatchCrossChainMessengerInitialized)
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
func (it *BatchCrossChainMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerInitialized represents a Initialized event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*BatchCrossChainMessengerInitializedIterator, error) {

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerInitializedIterator{contract: _BatchCrossChainMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerInitialized)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseInitialized(log types.Log) (*BatchCrossChainMessengerInitialized, error) {
	event := new(BatchCrossChainMessengerInitialized)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator struct {
	Event *BatchCrossChainMessengerMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerMinTeleporterVersionUpdated)
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
		it.Event = new(BatchCrossChainMessengerMinTeleporterVersionUpdated)
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
func (it *BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerMinTeleporterVersionUpdatedIterator{contract: _BatchCrossChainMessenger.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerMinTeleporterVersionUpdated)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*BatchCrossChainMessengerMinTeleporterVersionUpdated, error) {
	event := new(BatchCrossChainMessengerMinTeleporterVersionUpdated)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerOwnershipTransferredIterator struct {
	Event *BatchCrossChainMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerOwnershipTransferred)
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
		it.Event = new(BatchCrossChainMessengerOwnershipTransferred)
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
func (it *BatchCrossChainMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchCrossChainMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerOwnershipTransferredIterator{contract: _BatchCrossChainMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerOwnershipTransferred)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*BatchCrossChainMessengerOwnershipTransferred, error) {
	event := new(BatchCrossChainMessengerOwnershipTransferred)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerReceiveMessageIterator struct {
	Event *BatchCrossChainMessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerReceiveMessage)
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
		it.Event = new(BatchCrossChainMessengerReceiveMessage)
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
func (it *BatchCrossChainMessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerReceiveMessage represents a ReceiveMessage event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerReceiveMessage struct {
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (*BatchCrossChainMessengerReceiveMessageIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerReceiveMessageIterator{contract: _BatchCrossChainMessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerReceiveMessage, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerReceiveMessage)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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

// ParseReceiveMessage is a log parse operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseReceiveMessage(log types.Log) (*BatchCrossChainMessengerReceiveMessage, error) {
	event := new(BatchCrossChainMessengerReceiveMessage)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerSendMessagesIterator is returned from FilterSendMessages and is used to iterate over the raw logs and unpacked data for SendMessages events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerSendMessagesIterator struct {
	Event *BatchCrossChainMessengerSendMessages // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerSendMessagesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerSendMessages)
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
		it.Event = new(BatchCrossChainMessengerSendMessages)
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
func (it *BatchCrossChainMessengerSendMessagesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerSendMessagesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerSendMessages represents a SendMessages event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerSendMessages struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeTokenAddress         common.Address
	FeeAmount               *big.Int
	RequiredGasLimit        *big.Int
	Messages                []string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendMessages is a free log retrieval operation binding the contract event 0x430d1906813fdb2129a19139f4112a1396804605501a798df3a4042590ba20d5.
//
// Solidity: event SendMessages(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterSendMessages(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (*BatchCrossChainMessengerSendMessagesIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "SendMessages", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerSendMessagesIterator{contract: _BatchCrossChainMessenger.contract, event: "SendMessages", logs: logs, sub: sub}, nil
}

// WatchSendMessages is a free log subscription operation binding the contract event 0x430d1906813fdb2129a19139f4112a1396804605501a798df3a4042590ba20d5.
//
// Solidity: event SendMessages(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchSendMessages(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerSendMessages, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "SendMessages", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerSendMessages)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "SendMessages", log); err != nil {
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

// ParseSendMessages is a log parse operation binding the contract event 0x430d1906813fdb2129a19139f4112a1396804605501a798df3a4042590ba20d5.
//
// Solidity: event SendMessages(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string[] messages)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseSendMessages(log types.Log) (*BatchCrossChainMessengerSendMessages, error) {
	event := new(BatchCrossChainMessengerSendMessages)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "SendMessages", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerTeleporterAddressPausedIterator struct {
	Event *BatchCrossChainMessengerTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerTeleporterAddressPaused)
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
		it.Event = new(BatchCrossChainMessengerTeleporterAddressPaused)
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
func (it *BatchCrossChainMessengerTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*BatchCrossChainMessengerTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerTeleporterAddressPausedIterator{contract: _BatchCrossChainMessenger.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerTeleporterAddressPaused)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseTeleporterAddressPaused(log types.Log) (*BatchCrossChainMessengerTeleporterAddressPaused, error) {
	event := new(BatchCrossChainMessengerTeleporterAddressPaused)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchCrossChainMessengerTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerTeleporterAddressUnpausedIterator struct {
	Event *BatchCrossChainMessengerTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *BatchCrossChainMessengerTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCrossChainMessengerTeleporterAddressUnpaused)
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
		it.Event = new(BatchCrossChainMessengerTeleporterAddressUnpaused)
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
func (it *BatchCrossChainMessengerTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCrossChainMessengerTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCrossChainMessengerTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the BatchCrossChainMessenger contract.
type BatchCrossChainMessengerTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*BatchCrossChainMessengerTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &BatchCrossChainMessengerTeleporterAddressUnpausedIterator{contract: _BatchCrossChainMessenger.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *BatchCrossChainMessengerTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BatchCrossChainMessenger.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCrossChainMessengerTeleporterAddressUnpaused)
				if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BatchCrossChainMessenger *BatchCrossChainMessengerFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*BatchCrossChainMessengerTeleporterAddressUnpaused, error) {
	event := new(BatchCrossChainMessengerTeleporterAddressUnpaused)
	if err := _BatchCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
