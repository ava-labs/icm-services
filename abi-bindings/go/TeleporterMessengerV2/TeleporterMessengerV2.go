// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessengerv2

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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TeleporterICMMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterICMMessage struct {
	Message            TeleporterMessageV2
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessageV2 is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageV2 struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	OriginTeleporterAddress common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessengerV2MetaData contains all meta data concerning the TeleporterMessengerV2 contract.
var TeleporterMessengerV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initializerAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID_\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"contractIMessageSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVerifier\",\"outputs\":[{\"internalType\":\"contractIMessageVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161342d38038061342d83398101604081905261002e91610146565b60015f81905580556001600160a01b0382166100a35760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646170746560448201526872206164647265737360b81b60648201526084015b60405180910390fd5b6001600160a01b03811661010f5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a207a65726f20696e6974696160448201526c6c697a6572206164647265737360981b606482015260840161009a565b6001600160a01b0391821660a08190526080521660c052610177565b80516001600160a01b0381168114610141575f5ffd5b919050565b5f5f60408385031215610157575f5ffd5b6101608361012b565b915061016e6020840161012b565b90509250929050565b60805160a05160c0516132706101bd5f395f81816103b901526113a101525f81816102f10152610a9e01525f8181610392015281816109f60152611eab01526132705ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c80638ae9753d116100d9578063d67bdd2511610093578063e69d606a1161006e578063e69d606a146103ee578063e6e67bd514610455578063ebc3b1ba14610490578063ecc70428146104b3575f5ffd5b8063d67bdd251461038d578063d6dc0b0d146103b4578063df20e8bc146103db575f5ffd5b80638ae9753d146102ec5780639498bd7114610313578063a889818114610326578063a9a8561414610339578063c473eef81461034c578063d127dc9b14610384575f5ffd5b80633ba2b9831161012a5780633ba2b983146102615780633e249b56146102745780636244885014610287578063860a3b061461029a578063892bf412146102b95780638ac0fd04146102d9575f5ffd5b80630f635b6c1461017157806322296c3a146101865780632bc8b0bf146101995780632ca40f55146101bf5780632e27c22314610217578063399b77da14610242575b5f5ffd5b61018461017f366004612359565b6104bc565b005b6101846101943660046123b7565b61071b565b6101ac6101a73660046123d0565b61080e565b6040519081526020015b60405180910390f35b6102096101cd3660046123d0565b600560209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b6040516101b69291906123e7565b61022a6102253660046123d0565b61082a565b6040516001600160a01b0390911681526020016101b6565b6101ac6102503660046123d0565b5f9081526005602052604090205490565b61018461026f36600461240e565b6108b1565b61018461028236600461243f565b610a62565b6101ac61029536600461248d565b611020565b6101ac6102a83660046123d0565b60066020525f908152604090205481565b6102cc6102c73660046124c3565b611079565b6040516101b691906124e3565b6101846102e7366004612503565b6110aa565b61022a7f000000000000000000000000000000000000000000000000000000000000000081565b6101846103213660046123d0565b6112e1565b6101ac610334366004612536565b611494565b6101ac6103473660046125a6565b6114d6565b6101ac61035a366004612638565b6001600160a01b039182165f90815260096020908152604080832093909416825291909152205490565b6101ac60025481565b61022a7f000000000000000000000000000000000000000000000000000000000000000081565b61022a7f000000000000000000000000000000000000000000000000000000000000000081565b6101ac6103e93660046123d0565b611768565b6104366103fc3660046123d0565b5f90815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b0390931683526020830191909152016101b6565b61047b6104633660046123d0565b60046020525f90815260409020805460019091015482565b604080519283526020830191909152016101b6565b6104a361049e3660046123d0565b6117af565b60405190151581526020016101b6565b6101ac60035481565b60018054146104e65760405162461bcd60e51b81526004016104dd90612652565b60405180910390fd5b60026001819055545f906104fd9084908435611494565b5f818152600660205260409020549091508061052b5760405162461bcd60e51b81526004016104dd90612697565b808360405160200161053d919061292e565b60405160208183030381529060405280519060200120146105705760405162461bcd60e51b81526004016104dd90612940565b5f61058160a08501608086016123b7565b6001600160a01b03163b116105f55760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b60648201526084016104dd565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f82815260066020908152604080832083905586916106449187019087016123b7565b610652610100870187612989565b60405160240161066594939291906129cb565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f6106ac6106a560a08701608088016123b7565b5a846117c4565b90508061070f5760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b60648201526084016104dd565b50506001805550505050565b335f9081526009602090815260408083206001600160a01b03851684529091529020548061079c5760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b60648201526084016104dd565b335f8181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a361080a6001600160a01b03831633836117db565b5050565b5f8181526004602052604081206108249061183f565b92915050565b5f818152600760205260408120546108965760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b60648201526084016104dd565b505f908152600860205260409020546001600160a01b031690565b60015f54146108d25760405162461bcd60e51b81526004016104dd906129f5565b60025f81815590546108ea9060608401358435611494565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906109515760405162461bcd60e51b81526004016104dd90612697565b5f83604051602001610963919061292e565b60408051601f198184030181529190528251815160208301209192501461099c5760405162461bcd60e51b81526004016104dd90612940565b8360600135837f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d18685602001516040516109d7929190612a38565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90610a2b90879060040161292e565b5f604051808303815f87803b158015610a42575f5ffd5b505af1158015610a54573d5f5f3e3d5ffd5b505060015f55505050505050565b6001805414610a835760405162461bcd60e51b81526004016104dd90612652565b600260015560405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff0090610ad3908590600401612a6c565b6020604051808303815f875af1158015610aef573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b139190612aef565b610b785760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465724d657373656e6765723a206d6573736167652076657260448201526f1a599a58d85d1a5bdb8819985a5b195960821b60648201526084016104dd565b36610b838380612b0e565b905030610b9660608301604084016123b7565b6001600160a01b031614610c125760405162461bcd60e51b815260206004820152603860248201527f54656c65706f727465724d657373656e67657256323a20696e76616c6964206f60448201527f726967696e2074656c65706f727465722061646472657373000000000000000060648201526084016104dd565b5f610c2060e0830183612b2d565b808060200260200160405190810160405280939291908181526020015f905b82821015610c6b57610c5c60408302860136819003810190612c00565b81526020019060010190610c3f565b505050505090505f604051806101000160405280845f01358152602001846020016020810190610c9b91906123b7565b6001600160a01b0316815260608501356020820152604001610cc360a08601608087016123b7565b6001600160a01b0316815260a08501356020820152604001610ce860c0860186612c36565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050509082525060208101849052604001610d33610100860186612989565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505091525060025490915080610d8b5760405162461bcd60e51b81526004016104dd90612c7b565b80826040015114610df85760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b60648201526084016104dd565b5f610e0b876040013583855f0151611494565b5f8181526007602052604090205490915015610e7f5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b60648201526084016104dd565b610e8d338460a00151611851565b610eeb5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b60648201526084016104dd565b610ef881845f01516118bd565b6001600160a01b03861615610f2e575f81815260086020526040902080546001600160a01b0319166001600160a01b0388161790555b60c0830151515f5b81811015610f7457610f6c6002548a604001358760c001518481518110610f5f57610f5f612cc2565b602002602001015161192d565b600101610f36565b50604080518082018252855181526001600160a01b0389166020808301919091528a8301355f908152600490915291909120610faf91611a51565b336001600160a01b03168860400135837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e348a88604051610ff0929190612e46565b60405180910390a460e084015151156110125761101282896040013586611aab565b505060018055505050505050565b5f60015f54146110425760405162461bcd60e51b81526004016104dd906129f5565b60025f5561106f61105283612f87565b83355f90815260046020526040902061106a90611bda565b611cd4565b60015f5592915050565b604080518082019091525f80825260208201525f8381526004602052604090206110a39083611f1a565b9392505050565b60015f54146110cb5760405162461bcd60e51b81526004016104dd906129f5565b60025f5560018054146110f05760405162461bcd60e51b81526004016104dd90612652565b60026001558061115a5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b60648201526084016104dd565b6001600160a01b0382166111805760405162461bcd60e51b81526004016104dd9061302a565b5f838152600560205260409020546111aa5760405162461bcd60e51b81526004016104dd90612697565b5f838152600560205260409020600101546001600160a01b0383811691161461123b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e7472616374206164647265737300000000000000000060648201526084016104dd565b5f6112468383611fdb565b5f8581526005602052604081206002018054929350839290919061126b908490613092565b90915550505f8481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1916112cc916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156113255750825b90505f826001600160401b031660011480156113405750303b155b90508115801561134e575080155b1561136c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561139657845460ff60401b1916600160401b1785555b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146114245760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526c321034b734ba34b0b634bd32b960991b60648201526084016104dd565b856114415760405162461bcd60e51b81526004016104dd90612c7b565b6002869055831561148c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6040805130602082015290810184905260608101839052608081018290525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f60015f54146114f85760405162461bcd60e51b81526004016104dd906129f5565b60025f818155905490866001600160401b0381111561151957611519612b72565b60405190808252806020026020018201604052801561155d57816020015b604080518082019091525f80825260208201528152602001906001900390816115375790505b509050865f5b818110156116d5575f8a8a8381811061157e5761157e612cc2565b9050602002013590505f60075f8381526020019081526020015f20549050805f036115fa5760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b60648201526084016104dd565b6116058d8783611494565b82146116795760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e00000000000060648201526084016104dd565b5f828152600860209081526040918290205482518084019093528383526001600160a01b031690820181905286519091908790869081106116bc576116bc612cc2565b6020026020010181905250505050806001019050611563565b506040805160c0810182528b81525f60208201526117569181016116fe368b90038b018b6130a5565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f9201829052509385525050604080519283526020808401909152909201525083611cd4565b60015f559a9950505050505050505050565b6002545f908061178a5760405162461bcd60e51b81526004016104dd90612c7b565b5f600354600161179a9190613092565b90506117a7828583611494565b949350505050565b5f818152600760205260408120541515610824565b5f5f5f5f8451602086015f8989f195945050505050565b6040516001600160a01b0383811660248301526044820183905261183a91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611fe7565b505050565b805460018201545f91610824916130bf565b5f81515f0361186257506001610824565b81515f5b818110156118b357846001600160a01b031684828151811061188a5761188a612cc2565b60200260200101516001600160a01b0316036118ab57600192505050610824565b600101611866565b505f949350505050565b805f0361191c5760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b60648201526084016104dd565b5f9182526007602052604090912055565b5f61193c8484845f0151611494565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b031684526002909101548383015290810191909152805191925090611992575050505050565b5f8281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b03908116865260098552838620925151168552925282208054919290916119f8908490613092565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf578460200151604051611a4291906130d2565b60405180910390a45050505050565b600182018054829160028501915f9182611a6a836130f2565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a1015611b0d5760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b60648201526084016104dd565b80606001516001600160a01b03163b5f03611b2d5761183a838383612059565b602081015160e08201516040515f92611b4a92869260240161310a565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052606083015160808401519192505f91611b8e9190846117c4565b905080611ba757611ba0858585612059565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b60605f611bf06005611beb8561183f565b6120cd565b9050805f03611c3c57604080515f8082526020820190925290611c34565b604080518082019091525f8082526020820152815260200190600190039081611c0e5790505b509392505050565b5f816001600160401b03811115611c5557611c55612b72565b604051908082528060200260200182016040528015611c9957816020015b604080518082019091525f8082526020820152815260200190600190039081611c735790505b5090505f5b82811015611c3457611caf856120dc565b828281518110611cc157611cc1612cc2565b6020908102919091010152600101611c9e565b5f5f60035f8154611ce4906130f2565b91905081905590505f611cfd600254865f015184611494565b90505f604051806101200160405280848152602001336001600160a01b03168152602001306001600160a01b03168152602001875f0151815260200187602001516001600160a01b0316815260200187606001518152602001876080015181526020018681526020018760a0015181525090505f81604051602001611d8291906131ed565b60405160208183030381529060405290505f5f8860400151602001511115611de9576040880151516001600160a01b0316611dcf5760405162461bcd60e51b81526004016104dd9061302a565b60408801518051602090910151611de69190611fdb565b90505b60408051808201825289820151516001600160a01b03908116825260208083018590528351808501855286518783012081528082018481525f8a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558951915190919086907f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d190611e8c90889086906131ff565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90611ee09087906004016131ed565b5f604051808303815f87803b158015611ef7575f5ffd5b505af1158015611f09573d5f5f3e3d5ffd5b50969b9a5050505050505050505050565b604080518082019091525f8082526020820152611f368361183f565b8210611f8e5760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b60648201526084016104dd565b826002015f83855f0154611fa29190613092565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b5f6110a38333846121a6565b5f5f60205f8451602086015f885af180612006576040513d5f823e3d81fd5b50505f513d9150811561201d57806001141561202a565b6001600160a01b0384163b155b1561205357604051635274afe760e01b81526001600160a01b03851660048201526024016104dd565b50505050565b8060405160200161206a9190613211565b60408051601f1981840301815282825280516020918201205f878152600690925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985906120c0908590613211565b60405180910390a3505050565b5f8282188284100282186110a3565b604080518082019091525f8082526020820152815460018301548190036121455760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d7074792071756575650000000000000060448201526064016104dd565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b031990921690559061219c908390613092565b9093555090919050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156121ec573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122109190613223565b90506122276001600160a01b038616853086612309565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa15801561226b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061228f9190613223565b90508181116122f55760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016104dd565b6122ff82826130bf565b9695505050505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526120539186918216906323b872dd90608401611808565b5f6101208284031215612353575f5ffd5b50919050565b5f5f6040838503121561236a575f5ffd5b8235915060208301356001600160401b03811115612386575f5ffd5b61239285828601612342565b9150509250929050565b80356001600160a01b03811681146123b2575f5ffd5b919050565b5f602082840312156123c7575f5ffd5b6110a38261239c565b5f602082840312156123e0575f5ffd5b5035919050565b828152606081016110a3602083018480516001600160a01b03168252602090810151910152565b5f6020828403121561241e575f5ffd5b81356001600160401b03811115612433575f5ffd5b6117a784828501612342565b5f5f60408385031215612450575f5ffd5b82356001600160401b03811115612465575f5ffd5b830160808186031215612476575f5ffd5b91506124846020840161239c565b90509250929050565b5f6020828403121561249d575f5ffd5b81356001600160401b038111156124b2575f5ffd5b820160e081850312156110a3575f5ffd5b5f5f604083850312156124d4575f5ffd5b50508035926020909101359150565b815181526020808301516001600160a01b03169082015260408101610824565b5f5f5f60608486031215612515575f5ffd5b833592506125256020850161239c565b929592945050506040919091013590565b5f5f5f60608486031215612548575f5ffd5b505081359360208301359350604090920135919050565b5f5f83601f84011261256f575f5ffd5b5081356001600160401b03811115612585575f5ffd5b6020830191508360208260051b850101111561259f575f5ffd5b9250929050565b5f5f5f5f5f5f86880360a08112156125bc575f5ffd5b8735965060208801356001600160401b038111156125d8575f5ffd5b6125e48a828b0161255f565b9097509550506040603f19820112156125fb575f5ffd5b5060408701925060808701356001600160401b0381111561261a575f5ffd5b61262689828a0161255f565b979a9699509497509295939492505050565b5f5f60408385031215612649575f5ffd5b6124768361239c565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f5f8335601e198436030181126126f2575f5ffd5b83016020810192503590506001600160401b03811115612710575f5ffd5b8060051b360382131561259f575f5ffd5b8183526020830192505f815f5b8481101561275d576001600160a01b036127478361239c565b168652602095860195919091019060010161272e565b5093949350505050565b5f5f8335601e1984360301811261277c575f5ffd5b83016020810192503590506001600160401b0381111561279a575f5ffd5b8060061b360382131561259f575f5ffd5b8183526020830192505f815f5b8481101561275d57813586526001600160a01b036127d86020840161239c565b16602087015260409586019591909101906001016127b8565b5f5f8335601e19843603018112612806575f5ffd5b83016020810192503590506001600160401b03811115612824575f5ffd5b80360382131561259f575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f61286b6020830161239c565b6001600160a01b031660208401526128856040830161239c565b6001600160a01b03166040840152606082810135908401526128a96080830161239c565b6001600160a01b0316608084015260a082810135908401526128ce60c08301836126dd565b61012060c08601526128e561012086018284612721565b9150506128f560e0840184612767565b85830360e08701526129088382846127ab565b9250505061291a6101008401846127f1565b8583036101008701526122ff838284612832565b602081525f6110a3602083018461285a565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b5f5f8335601e1984360301811261299e575f5ffd5b8301803591506001600160401b038211156129b7575f5ffd5b60200191503681900382131561259f575f5ffd5b8481526001600160a01b03841660208201526060604082018190525f906122ff9083018486612832565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b606081525f612a4a606083018561285a565b90506110a3602083018480516001600160a01b03168252602090810151910152565b602081525f823561011e19843603018112612a85575f5ffd5b60806020840152612a9b60a0840185830161285a565b9050602084013563ffffffff8116808214612ab4575f5ffd5b80604086015250505f6040850135905080606085015250612ad860608501856127f1565b848303601f190160808601526122ff838284612832565b5f60208284031215612aff575f5ffd5b815180151581146110a3575f5ffd5b5f823561011e19833603018112612b23575f5ffd5b9190910192915050565b5f5f8335601e19843603018112612b42575f5ffd5b8301803591506001600160401b03821115612b5b575f5ffd5b6020019150600681901b360382131561259f575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715612ba857612ba8612b72565b60405290565b60405160c081016001600160401b0381118282101715612ba857612ba8612b72565b604051601f8201601f191681016001600160401b0381118282101715612bf857612bf8612b72565b604052919050565b5f6040828403128015612c11575f5ffd5b50612c1a612b86565b82358152612c2a6020840161239c565b60208201529392505050565b5f5f8335601e19843603018112612c4b575f5ffd5b8301803591506001600160401b03821115612c64575f5ffd5b6020019150600581901b360382131561259f575f5ffd5b60208082526027908201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636040820152661a185a5b88125160ca1b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b8281101561275d5781516001600160a01b0316865260209586019590910190600101612ce8565b5f8151808452602084019350602083015f5b8281101561275d57612d47868351805182526020908101516001600160a01b0316910152565b6040959095019460209190910190600101612d21565b5f81518084525f5b81811015612d8157602081850181015186830182015201612d65565b505f602082860101526020601f19601f83011685010191505092915050565b805182525f6020820151612dbf60208501826001600160a01b03169052565b50604082015160408401526060820151612de460608501826001600160a01b03169052565b506080820151608084015260a082015161010060a0850152612e0a610100850182612cd6565b905060c083015184820360c0860152612e238282612d0f565b91505060e083015184820360e0860152612e3d8282612d5d565b95945050505050565b6001600160a01b03831681526040602082018190525f906117a790830184612da0565b5f60408284031215612e79575f5ffd5b612e81612b86565b9050612e8c8261239c565b815260209182013591810191909152919050565b5f82601f830112612eaf575f5ffd5b81356001600160401b03811115612ec857612ec8612b72565b8060051b612ed860208201612bd0565b91825260208185018101929081019086841115612ef3575f5ffd5b6020860192505b838310156122ff57612f0b8361239c565b825260209283019290910190612efa565b5f82601f830112612f2b575f5ffd5b81356001600160401b03811115612f4457612f44612b72565b612f57601f8201601f1916602001612bd0565b818152846020838601011115612f6b575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60e08236031215612f97575f5ffd5b612f9f612bae565b82358152612faf6020840161239c565b6020820152612fc13660408501612e69565b60408201526080830135606082015260a08301356001600160401b03811115612fe8575f5ffd5b612ff436828601612ea0565b60808301525060c08301356001600160401b03811115613012575f5ffd5b61301e36828601612f1c565b60a08301525092915050565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156108245761082461307e565b5f604082840312156130b5575f5ffd5b6110a38383612e69565b818103818111156108245761082461307e565b81516001600160a01b031681526020808301519082015260408101610824565b5f600182016131035761310361307e565b5060010190565b8381526001600160a01b03831660208201526060604082018190525f90612e3d90830184612d5d565b805182525f602082015161315260208501826001600160a01b03169052565b50604082015161316d60408501826001600160a01b03169052565b5060608201516060840152608082015161319260808501826001600160a01b03169052565b5060a082015160a084015260c082015161012060c08501526131b8610120850182612cd6565b905060e083015184820360e08601526131d18282612d0f565b915050610100830151848203610100860152612e3d8282612d5d565b602081525f6110a36020830184613133565b606081525f612a4a6060830185613133565b602081525f6110a36020830184612da0565b5f60208284031215613233575f5ffd5b505191905056fea2646970667358221220ed26f0530a70b90d9f71debe1139c974e8fb57e60d8cc3ffd10ece815412e4d964736f6c634300081e0033",
}

// TeleporterMessengerV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerV2MetaData.ABI instead.
var TeleporterMessengerV2ABI = TeleporterMessengerV2MetaData.ABI

// TeleporterMessengerV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterMessengerV2MetaData.Bin instead.
var TeleporterMessengerV2Bin = TeleporterMessengerV2MetaData.Bin

// DeployTeleporterMessengerV2 deploys a new Ethereum contract, binding an instance of TeleporterMessengerV2 to it.
func DeployTeleporterMessengerV2(auth *bind.TransactOpts, backend bind.ContractBackend, verifierSender common.Address, initializerAddress_ common.Address) (common.Address, *types.Transaction, *TeleporterMessengerV2, error) {
	parsed, err := TeleporterMessengerV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterMessengerV2Bin), backend, verifierSender, initializerAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterMessengerV2{TeleporterMessengerV2Caller: TeleporterMessengerV2Caller{contract: contract}, TeleporterMessengerV2Transactor: TeleporterMessengerV2Transactor{contract: contract}, TeleporterMessengerV2Filterer: TeleporterMessengerV2Filterer{contract: contract}}, nil
}

// TeleporterMessengerV2 is an auto generated Go binding around an Ethereum contract.
type TeleporterMessengerV2 struct {
	TeleporterMessengerV2Caller     // Read-only binding to the contract
	TeleporterMessengerV2Transactor // Write-only binding to the contract
	TeleporterMessengerV2Filterer   // Log filterer for contract events
}

// TeleporterMessengerV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterMessengerV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterMessengerV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterMessengerV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterMessengerV2Session struct {
	Contract     *TeleporterMessengerV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeleporterMessengerV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterMessengerV2CallerSession struct {
	Contract *TeleporterMessengerV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TeleporterMessengerV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterMessengerV2TransactorSession struct {
	Contract     *TeleporterMessengerV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TeleporterMessengerV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterMessengerV2Raw struct {
	Contract *TeleporterMessengerV2 // Generic contract binding to access the raw methods on
}

// TeleporterMessengerV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterMessengerV2CallerRaw struct {
	Contract *TeleporterMessengerV2Caller // Generic read-only contract binding to access the raw methods on
}

// TeleporterMessengerV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterMessengerV2TransactorRaw struct {
	Contract *TeleporterMessengerV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterMessengerV2 creates a new instance of TeleporterMessengerV2, bound to a specific deployed contract.
func NewTeleporterMessengerV2(address common.Address, backend bind.ContractBackend) (*TeleporterMessengerV2, error) {
	contract, err := bindTeleporterMessengerV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2{TeleporterMessengerV2Caller: TeleporterMessengerV2Caller{contract: contract}, TeleporterMessengerV2Transactor: TeleporterMessengerV2Transactor{contract: contract}, TeleporterMessengerV2Filterer: TeleporterMessengerV2Filterer{contract: contract}}, nil
}

// NewTeleporterMessengerV2Caller creates a new read-only instance of TeleporterMessengerV2, bound to a specific deployed contract.
func NewTeleporterMessengerV2Caller(address common.Address, caller bind.ContractCaller) (*TeleporterMessengerV2Caller, error) {
	contract, err := bindTeleporterMessengerV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2Caller{contract: contract}, nil
}

// NewTeleporterMessengerV2Transactor creates a new write-only instance of TeleporterMessengerV2, bound to a specific deployed contract.
func NewTeleporterMessengerV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterMessengerV2Transactor, error) {
	contract, err := bindTeleporterMessengerV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2Transactor{contract: contract}, nil
}

// NewTeleporterMessengerV2Filterer creates a new log filterer instance of TeleporterMessengerV2, bound to a specific deployed contract.
func NewTeleporterMessengerV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterMessengerV2Filterer, error) {
	contract, err := bindTeleporterMessengerV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2Filterer{contract: contract}, nil
}

// bindTeleporterMessengerV2 binds a generic wrapper to an already deployed contract.
func bindTeleporterMessengerV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterMessengerV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessengerV2 *TeleporterMessengerV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessengerV2.Contract.TeleporterMessengerV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessengerV2 *TeleporterMessengerV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.TeleporterMessengerV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessengerV2 *TeleporterMessengerV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.TeleporterMessengerV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessengerV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.contract.Transact(opts, method, params...)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) BlockchainID() ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.BlockchainID(&_TeleporterMessengerV2.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.BlockchainID(&_TeleporterMessengerV2.CallOpts)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) CalculateMessageID(opts *bind.CallOpts, sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "calculateMessageID", sourceBlockchainID, destinationBlockchainID, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.CalculateMessageID(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.CalculateMessageID(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.CheckRelayerRewardAmount(&_TeleporterMessengerV2.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.CheckRelayerRewardAmount(&_TeleporterMessengerV2.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getFeeInfo", messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessengerV2.Contract.GetFeeInfo(&_TeleporterMessengerV2.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessengerV2.Contract.GetFeeInfo(&_TeleporterMessengerV2.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.GetMessageHash(&_TeleporterMessengerV2.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.GetMessageHash(&_TeleporterMessengerV2.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.GetNextMessageID(&_TeleporterMessengerV2.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.GetNextMessageID(&_TeleporterMessengerV2.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessengerV2.Contract.GetReceiptAtIndex(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessengerV2.Contract.GetReceiptAtIndex(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.GetReceiptQueueSize(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.GetReceiptQueueSize(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessengerV2.Contract.GetRelayerRewardAddress(&_TeleporterMessengerV2.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessengerV2.Contract.GetRelayerRewardAddress(&_TeleporterMessengerV2.CallOpts, messageID)
}

// InitializerAddress is a free data retrieval call binding the contract method 0xd6dc0b0d.
//
// Solidity: function initializerAddress() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) InitializerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "initializerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitializerAddress is a free data retrieval call binding the contract method 0xd6dc0b0d.
//
// Solidity: function initializerAddress() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) InitializerAddress() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.InitializerAddress(&_TeleporterMessengerV2.CallOpts)
}

// InitializerAddress is a free data retrieval call binding the contract method 0xd6dc0b0d.
//
// Solidity: function initializerAddress() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) InitializerAddress() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.InitializerAddress(&_TeleporterMessengerV2.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) MessageNonce() (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.MessageNonce(&_TeleporterMessengerV2.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessengerV2.Contract.MessageNonce(&_TeleporterMessengerV2.CallOpts)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessengerV2.Contract.MessageReceived(&_TeleporterMessengerV2.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessengerV2.Contract.MessageReceived(&_TeleporterMessengerV2.CallOpts, messageID)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) MessageSender() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.MessageSender(&_TeleporterMessengerV2.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) MessageSender() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.MessageSender(&_TeleporterMessengerV2.CallOpts)
}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) MessageVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "messageVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) MessageVerifier() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.MessageVerifier(&_TeleporterMessengerV2.CallOpts)
}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) MessageVerifier() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.MessageVerifier(&_TeleporterMessengerV2.CallOpts)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) ReceiptQueues(opts *bind.CallOpts, sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "receiptQueues", sourceBlockchainID)

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessengerV2.Contract.ReceiptQueues(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessengerV2.Contract.ReceiptQueues(&_TeleporterMessengerV2.CallOpts, sourceBlockchainID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) ReceivedFailedMessageHashes(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "receivedFailedMessageHashes", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.ReceivedFailedMessageHashes(&_TeleporterMessengerV2.CallOpts, messageID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessengerV2.Contract.ReceivedFailedMessageHashes(&_TeleporterMessengerV2.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) SentMessageInfo(opts *bind.CallOpts, messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "sentMessageInfo", messageID)

	outstruct := new(struct {
		MessageHash [32]byte
		FeeInfo     TeleporterFeeInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.FeeInfo = *abi.ConvertType(out[1], new(TeleporterFeeInfo)).(*TeleporterFeeInfo)

	return *outstruct, err

}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessengerV2.Contract.SentMessageInfo(&_TeleporterMessengerV2.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessengerV2.Contract.SentMessageInfo(&_TeleporterMessengerV2.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.AddFeeAmount(&_TeleporterMessengerV2.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.AddFeeAmount(&_TeleporterMessengerV2.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 blockchainID_) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) Initialize(opts *bind.TransactOpts, blockchainID_ [32]byte) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "initialize", blockchainID_)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 blockchainID_) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) Initialize(blockchainID_ [32]byte) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.Initialize(&_TeleporterMessengerV2.TransactOpts, blockchainID_)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 blockchainID_) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) Initialize(blockchainID_ [32]byte) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.Initialize(&_TeleporterMessengerV2.TransactOpts, blockchainID_)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x3e249b56.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, message TeleporterICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "receiveCrossChainMessage", message, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x3e249b56.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) ReceiveCrossChainMessage(message TeleporterICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.ReceiveCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, message, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x3e249b56.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) ReceiveCrossChainMessage(message TeleporterICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.ReceiveCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, message, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RedeemRelayerRewards(&_TeleporterMessengerV2.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RedeemRelayerRewards(&_TeleporterMessengerV2.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RetryMessageExecution(&_TeleporterMessengerV2.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RetryMessageExecution(&_TeleporterMessengerV2.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) RetrySendCrossChainMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RetrySendCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) RetrySendCrossChainMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.RetrySendCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.SendCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.SendCrossChainMessage(&_TeleporterMessengerV2.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Transactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.SendSpecifiedReceipts(&_TeleporterMessengerV2.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessengerV2 *TeleporterMessengerV2TransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessengerV2.Contract.SendSpecifiedReceipts(&_TeleporterMessengerV2.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// TeleporterMessengerV2AddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2AddFeeAmountIterator struct {
	Event *TeleporterMessengerV2AddFeeAmount // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2AddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2AddFeeAmount)
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
		it.Event = new(TeleporterMessengerV2AddFeeAmount)
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
func (it *TeleporterMessengerV2AddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2AddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2AddFeeAmount represents a AddFeeAmount event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2AddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*TeleporterMessengerV2AddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2AddFeeAmountIterator{contract: _TeleporterMessengerV2.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2AddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2AddFeeAmount)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseAddFeeAmount(log types.Log) (*TeleporterMessengerV2AddFeeAmount, error) {
	event := new(TeleporterMessengerV2AddFeeAmount)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2InitializedIterator struct {
	Event *TeleporterMessengerV2Initialized // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2Initialized)
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
		it.Event = new(TeleporterMessengerV2Initialized)
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
func (it *TeleporterMessengerV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2Initialized represents a Initialized event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*TeleporterMessengerV2InitializedIterator, error) {

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2InitializedIterator{contract: _TeleporterMessengerV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2Initialized) (event.Subscription, error) {

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2Initialized)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseInitialized(log types.Log) (*TeleporterMessengerV2Initialized, error) {
	event := new(TeleporterMessengerV2Initialized)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2MessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2MessageExecutedIterator struct {
	Event *TeleporterMessengerV2MessageExecuted // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2MessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2MessageExecuted)
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
		it.Event = new(TeleporterMessengerV2MessageExecuted)
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
func (it *TeleporterMessengerV2MessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2MessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2MessageExecuted represents a MessageExecuted event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2MessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerV2MessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2MessageExecutedIterator{contract: _TeleporterMessengerV2.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2MessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2MessageExecuted)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

// ParseMessageExecuted is a log parse operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseMessageExecuted(log types.Log) (*TeleporterMessengerV2MessageExecuted, error) {
	event := new(TeleporterMessengerV2MessageExecuted)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2MessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2MessageExecutionFailedIterator struct {
	Event *TeleporterMessengerV2MessageExecutionFailed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2MessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2MessageExecutionFailed)
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
		it.Event = new(TeleporterMessengerV2MessageExecutionFailed)
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
func (it *TeleporterMessengerV2MessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2MessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2MessageExecutionFailed represents a MessageExecutionFailed event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2MessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerV2MessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2MessageExecutionFailedIterator{contract: _TeleporterMessengerV2.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2MessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2MessageExecutionFailed)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseMessageExecutionFailed(log types.Log) (*TeleporterMessengerV2MessageExecutionFailed, error) {
	event := new(TeleporterMessengerV2MessageExecutionFailed)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2ReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2ReceiptReceivedIterator struct {
	Event *TeleporterMessengerV2ReceiptReceived // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2ReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2ReceiptReceived)
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
		it.Event = new(TeleporterMessengerV2ReceiptReceived)
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
func (it *TeleporterMessengerV2ReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2ReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2ReceiptReceived represents a ReceiptReceived event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2ReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*TeleporterMessengerV2ReceiptReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2ReceiptReceivedIterator{contract: _TeleporterMessengerV2.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2ReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2ReceiptReceived)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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

// ParseReceiptReceived is a log parse operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseReceiptReceived(log types.Log) (*TeleporterMessengerV2ReceiptReceived, error) {
	event := new(TeleporterMessengerV2ReceiptReceived)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2ReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2ReceiveCrossChainMessageIterator struct {
	Event *TeleporterMessengerV2ReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2ReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2ReceiveCrossChainMessage)
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
		it.Event = new(TeleporterMessengerV2ReceiveCrossChainMessage)
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
func (it *TeleporterMessengerV2ReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2ReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2ReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2ReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*TeleporterMessengerV2ReceiveCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2ReceiveCrossChainMessageIterator{contract: _TeleporterMessengerV2.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2ReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2ReceiveCrossChainMessage)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleporterMessengerV2ReceiveCrossChainMessage, error) {
	event := new(TeleporterMessengerV2ReceiveCrossChainMessage)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2RelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2RelayerRewardsRedeemedIterator struct {
	Event *TeleporterMessengerV2RelayerRewardsRedeemed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2RelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2RelayerRewardsRedeemed)
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
		it.Event = new(TeleporterMessengerV2RelayerRewardsRedeemed)
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
func (it *TeleporterMessengerV2RelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2RelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2RelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2RelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*TeleporterMessengerV2RelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2RelayerRewardsRedeemedIterator{contract: _TeleporterMessengerV2.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2RelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2RelayerRewardsRedeemed)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseRelayerRewardsRedeemed(log types.Log) (*TeleporterMessengerV2RelayerRewardsRedeemed, error) {
	event := new(TeleporterMessengerV2RelayerRewardsRedeemed)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerV2SendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2SendCrossChainMessageIterator struct {
	Event *TeleporterMessengerV2SendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerV2SendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerV2SendCrossChainMessage)
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
		it.Event = new(TeleporterMessengerV2SendCrossChainMessage)
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
func (it *TeleporterMessengerV2SendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerV2SendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerV2SendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleporterMessengerV2 contract.
type TeleporterMessengerV2SendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessageV2
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*TeleporterMessengerV2SendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerV2SendCrossChainMessageIterator{contract: _TeleporterMessengerV2.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerV2SendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessengerV2.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerV2SendCrossChainMessage)
				if err := _TeleporterMessengerV2.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Filterer) ParseSendCrossChainMessage(log types.Log) (*TeleporterMessengerV2SendCrossChainMessage, error) {
	event := new(TeleporterMessengerV2SendCrossChainMessage)
	if err := _TeleporterMessengerV2.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
