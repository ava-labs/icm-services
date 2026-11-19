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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"}],\"name\":\"SendMessages\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessages\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"}],\"name\":\"sendMessages\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346103d957606081611c4a803803809161001f828561052e565b8339810103126103d95761003281610565565b90604061004160208301610565565b9101515f516020611c0a5f395f51905f5254604081901c60ff161593906001600160401b03811680159081610526575b600114908161051c575b159081610513575b50610504576001600160401b031981166001175f516020611c0a5f395f51905f5255846104da575b506100b4610579565b6100bc610579565b60015f516020611bea5f395f51905f52556100d5610579565b6100dd610579565b6100e5610579565b6100ed610579565b60015f516020611bea5f395f51905f5255610106610579565b61010e610579565b6001600160a01b0316801561046f5760405163301fd1f560e21b8152602081600481855afa9081156103e5575f9161043d575b50156103f0577fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0080546001600160a01b0319168217905560405163301fd1f560e21b815290602090829060049082905afa9081156103e5575f916103af575b505f516020611bca5f395f51905f525490821161036357808211156102f857815f516020611bca5f395f51905f52557fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d5f80a36101fb610579565b610203610579565b6001600160a01b03169081156102e5577f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b03198116841790915560405192906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a361028c575b60405161162590816105a58239f35b60207fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29168ff0000000000000000195f516020611c0a5f395f51905f5254165f516020611c0a5f395f51905f525560018152a15f61027d565b631e4fbdf760e01b5f525f60045260245ffd5b60405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608490fd5b60405162461bcd60e51b815260206004820152603160248201525f516020611c2a5f395f51905f5260448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608490fd5b90506020813d6020116103dd575b816103ca6020938361052e565b810103126103d957515f6101a0565b5f80fd5b3d91506103bd565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152603260248201525f516020611c2a5f395f51905f52604482015271656c65706f7274657220726567697374727960701b6064820152608490fd5b90506020813d602011610467575b816104586020938361052e565b810103126103d957515f610141565b3d915061044b565b60405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608490fd5b6001600160481b03191668010000000000000001175f516020611c0a5f395f51905f52555f6100ab565b63f92ee8a960e01b5f5260045ffd5b9050155f610083565b303b15915061007b565b869150610071565b601f909101601f19168101906001600160401b0382119082101761055157604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036103d957565b60ff5f516020611c0a5f395f51905f525460401c161561059557565b631afcd79f60e31b5f5260045ffdfe60806040526004361015610011575f80fd5b5f3560e01c80632b0d8f18146111bb5780633902970c146109ff5780634511243e146109285780635eb995141461077b578063715018a6146107145780638da5cb5b146106e0578063909a6ac0146106b95780639731429714610688578063c1329fcb1461057b578063c868efaa1461015c578063d2cc7a70146101335763f2fde38b1461009d575f80fd5b3461012f57602036600319011261012f576100b6611298565b6100be61150d565b6001600160a01b0316801561011c575f5160206115995f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b3461012f575f36600319011261012f5760205f5160206115b95f395f51905f5254604051908152f35b3461012f57606036600319011261012f576004356101786112ae565b9060443567ffffffffffffffff811161012f573660238201121561012f5780600401359067ffffffffffffffff821161012f57366024838301011161012f576101bf6114d5565b5f5160206115f95f395f51905f525460405163260f846760e11b815233600482015290602090829060249082906001600160a01b03165afa908115610570575f9161053e575b505f5160206115b95f395f51905f5254116104e05760ff61022533611451565b54166104825761023991602436920161131a565b91825183019260208181860195031261012f5760208101519067ffffffffffffffff821161012f570183603f8201121561012f57602081015161027b816112fe565b9461028960405196876112c4565b8186526040838301011161012f576102a8916040602087019101611350565b815f525f60205260405f208054906801000000000000000082101561045a576001820180825582101561046e575f5260205f2001835167ffffffffffffffff811161045a576102f7825461149d565b601f8111610415575b506020601f821160011461038f5791817f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa949261036e945f91610384575b508160011b915f199060031b1c19161790555b6040519182916020835260018060a01b0316956020830190611371565b0390a360015f5160206115d95f395f51905f5255005b90508701518861033e565b601f19821690835f52805f20915f5b8181106103fd57509261036e9492600192827f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa9896106103e5575b5050811b019055610351565b8901515f1960f88460031b161c1916905588806103d9565b9192602060018192868c01518155019401920161039e565b825f5260205f20601f830160051c81019160208410610450575b601f0160051c01905b8181106104455750610300565b5f8155600101610438565b909150819061042f565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b60405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608490fd5b60405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608490fd5b90506020813d602011610568575b81610559602093836112c4565b8101031261012f575185610205565b3d915061054c565b6040513d5f823e3d90fd5b3461012f57602036600319011261012f576004355f525f60205260405f2080546105a4816112e6565b906105b260405192836112c4565b80825260208201925f5260205f20925f905b8282106105e557604051602080825281906105e190820187611396565b0390f35b6040515f86546105f48161149d565b8084529060018116908115610665575060011461062e575b5060019282610620859460209403826112c4565b8152019501910190936105c4565b5f888152602081209092505b81831061064f5750508101602001600161060c565b600181602092548386880101520192019161063a565b60ff191660208581019190915291151560051b840190910191506001905061060c565b3461012f57602036600319011261012f57602060ff6106ad6106a8611298565b611451565b54166040519015158152f35b3461012f575f36600319011261012f5760206040515f5160206115f95f395f51905f528152f35b3461012f575f36600319011261012f575f5160206115995f395f51905f52546040516001600160a01b039091168152602090f35b3461012f575f36600319011261012f5761072c61150d565b5f5160206115995f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461012f57602036600319011261012f5760043561079761150d565b5f5160206115f95f395f51905f525460405163301fd1f560e21b815290602090829060049082906001600160a01b03165afa908115610570575f916108f6575b505f5160206115b95f395f51905f5254908211610897578082111561082c57815f5160206115b95f395f51905f52557fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d5f80a3005b60405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608490fd5b60405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608490fd5b90506020813d602011610920575b81610911602093836112c4565b8101031261012f5751826107d7565b3d9150610904565b3461012f57602036600319011261012f57610941611298565b61094961150d565b6001600160a01b0381169061095f8215156113ee565b60ff61096a82611451565b5416156109a85761097a90611451565b805460ff191690557f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c35f80a2005b60405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608490fd5b3461012f5760c036600319011261012f57600435610a1b6112ae565b6044356001600160a01b0381169081900361012f57606435926084359360a4359367ffffffffffffffff851161012f573660238601121561012f578460040135610a64816112e6565b95610a7260405197886112c4565b8187526024602088019260051b8201019036821161012f5760248101925b82841061117a5750505050610aa36114d5565b5f9180611009575b5060018060a01b03169182817f430d1906813fdb2129a19139f4112a1396804605501a798df3a4042590ba20d56040518781528560208201528960408201526080606082015280610aff608082018b611396565b0390a3845193610b27610b11866112e6565b95610b1f60405197886112c4565b8087526112e6565b602086019690601f19013688375f5b8151811015610fb457604051906040820182811067ffffffffffffffff82111761045a57604052838252856020830152602060405192610b7682856112c4565b5f84525f368137610ba3610bb1610b8d8588611489565b5160405192839186808401526040830190611371565b03601f1981018352826112c4565b6040519060c082019582871067ffffffffffffffff88111761045a578e96604052898352848301968c885260408401948552606084019081526080840191825260a0840192835260048660018060a01b035f5160206115f95f395f51905f5254166040519283809263d820e64f60e01b82525afa908115610570575f91610f77575b506001600160a01b03169460ff610c4987611451565b5416610f19578087915182810151610d6e575b50604051630624488560e41b8152600481018390529551602487015298516001600160a01b0390811660448701529851805190991660648601529788015160848501525160a48401525160e060c484015280516101048401819052929693879390926101248501929091908601905f905b808210610d485750505083610cf181935f93516023198483030160e4850152611371565b03925af1908115610570575f91610d19575b5060019250610d12828a611489565b5201610b36565b905082813d8311610d41575b610d2f81836112c4565b8101031261012f57600191518b610d03565b503d610d25565b82516001600160a01b0316855288978b9750948501949092019160019190910190610ccd565b519091506001600160a01b031615610ebe578051805190880151604051636eb1769f60e11b81523060048201526024810189905292916001600160a01b0316908984604481855afa938415610570575f94610e8f575b508301809311610e7b578892835f604051928284019063095ea7b360e01b82528c6024860152604485015260448452610dfe6064856112c4565b83519082865af15f513d82610e5f575b505015610e1c575b50610c5c565b610e5891610e5360405163095ea7b360e01b878201528b60248201525f604482015260448152610e4d6064826112c4565b82611540565b611540565b5f80610e16565b909150610e735750813b15155b5f80610e0e565b600114610e6c565b634e487b7160e01b5f52601160045260245ffd5b9093508981813d8311610eb7575b610ea781836112c4565b8101031261012f5751925f610dc4565b503d610e9d565b60405162461bcd60e51b815260048101889052602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608490fd5b60405162461bcd60e51b815260048101889052603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608490fd5b90508681813d8311610fad575b610f8e81836112c4565b8101031261012f57516001600160a01b038116810361012f575f610c33565b503d610f84565b878760015f5160206115d95f395f51905f5255604051918291602083019060208452518091526040830191905f5b818110610ff0575050500390f35b8251845285945060209384019390920191600101610fe2565b6040516370a0823160e01b8152306004820152919250602082602481885afa918215610570575f92611144575b5061107290604051906323b872dd60e01b602083015233602483015230604483015260648201526064815261106c6084826112c4565b85611540565b6040516370a0823160e01b815230600482015290602082602481885afa918215610570575f92611110575b50808211156110b6578103908111610e7b579086610aab565b60405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608490fd5b9091506020813d60201161113c575b8161112c602093836112c4565b8101031261012f5751908761109d565b3d915061111f565b9091506020813d602011611172575b81611160602093836112c4565b8101031261012f575190611072611036565b3d9150611153565b833567ffffffffffffffff811161012f5782013660438201121561012f576020916111b08392369060446024820135910161131a565b815201930192610a90565b3461012f57602036600319011261012f576111d4611298565b6111dc61150d565b6001600160a01b038116906111f28215156113ee565b60ff6111fd82611451565b541661123d5761120c90611451565b805460ff191660011790557f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c5f80a2005b60405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608490fd5b600435906001600160a01b038216820361012f57565b602435906001600160a01b038216820361012f57565b90601f8019910116810190811067ffffffffffffffff82111761045a57604052565b67ffffffffffffffff811161045a5760051b60200190565b67ffffffffffffffff811161045a57601f01601f191660200190565b929192611326826112fe565b9161133460405193846112c4565b82948184528183011161012f578281602093845f960137010152565b5f5b8381106113615750505f910152565b8181015183820152602001611352565b9060209161138a81518092818552858086019101611350565b601f01601f1916010190565b9080602083519182815201916020808360051b8301019401925f915b8383106113c157505050505090565b90919293946020806113df600193601f198682030187528951611371565b970193019301919392906113b2565b156113f557565b60405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201526d706f72746572206164647265737360901b6064820152608490fd5b6001600160a01b03165f9081527fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d016020526040902090565b805182101561046e5760209160051b010190565b90600182811c921680156114cb575b60208310146114b757565b634e487b7160e01b5f52602260045260245ffd5b91607f16916114ac565b60025f5160206115d95f395f51905f5254146114fe5760025f5160206115d95f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b5f5160206115995f395f51905f52546001600160a01b0316330361152d57565b63118cdaa760e01b5f523360045260245ffd5b905f602091828151910182855af115610570575f513d61158f57506001600160a01b0381163b155b61156f5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b6001141561156856fe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d029b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a164736f6c634300081e000ade77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d029b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054656c65706f7274657252656769737472794170703a20696e76616c69642054",
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
