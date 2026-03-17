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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierSender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID_\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"contractIMessageSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVerifier\",\"outputs\":[{\"internalType\":\"contractIMessageVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161334138038061334183398101604081905261002e9161004c565b60015f81905580556001600160a01b031660a0819052608052610079565b5f6020828403121561005c575f5ffd5b81516001600160a01b0381168114610072575f5ffd5b9392505050565b60805160a0516132926100af5f395f81816102d70152610a5d01525f8181610378015281816109b50152611de201526132925ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c80638ae9753d116100bf578063d67bdd2511610079578063d67bdd2514610373578063df20e8bc1461039a578063e69d606a146103ad578063e6e67bd514610414578063ebc3b1ba1461044f578063ecc7042814610472575f5ffd5b80638ae9753d146102d25780639498bd71146102f9578063a88981811461030c578063a9a856141461031f578063c473eef814610332578063d127dc9b1461036a575f5ffd5b80633ba2b983116101105780633ba2b983146102475780633e249b561461025a578063624488501461026d578063860a3b0614610280578063892bf4121461029f5780638ac0fd04146102bf575f5ffd5b80630f635b6c1461015757806322296c3a1461016c5780632bc8b0bf1461017f5780632ca40f55146101a55780632e27c223146101fd578063399b77da14610228575b5f5ffd5b61016a6101653660046123a7565b61047b565b005b61016a61017a366004612405565b6106da565b61019261018d36600461241e565b6107cd565b6040519081526020015b60405180910390f35b6101ef6101b336600461241e565b600560209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b60405161019c929190612435565b61021061020b36600461241e565b6107e9565b6040516001600160a01b03909116815260200161019c565b61019261023636600461241e565b5f9081526005602052604090205490565b61016a61025536600461245c565b610870565b61016a61026836600461248d565b610a21565b61019261027b3660046124db565b610fc2565b61019261028e36600461241e565b60066020525f908152604090205481565b6102b26102ad366004612511565b61101b565b60405161019c9190612531565b61016a6102cd366004612551565b61104c565b6102107f000000000000000000000000000000000000000000000000000000000000000081565b61016a61030736600461241e565b611283565b61019261031a366004612584565b61138b565b61019261032d3660046125f4565b6113cd565b610192610340366004612686565b6001600160a01b039182165f90815260096020908152604080832093909416825291909152205490565b61019260025481565b6102107f000000000000000000000000000000000000000000000000000000000000000081565b6101926103a836600461241e565b61165f565b6103f56103bb36600461241e565b5f90815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b03909316835260208301919091520161019c565b61043a61042236600461241e565b60046020525f90815260409020805460019091015482565b6040805192835260208301919091520161019c565b61046261045d36600461241e565b6116e6565b604051901515815260200161019c565b61019260035481565b60018054146104a55760405162461bcd60e51b815260040161049c906126a0565b60405180910390fd5b60026001819055545f906104bc908490843561138b565b5f81815260066020526040902054909150806104ea5760405162461bcd60e51b815260040161049c906126e5565b80836040516020016104fc919061297c565b604051602081830303815290604052805190602001201461052f5760405162461bcd60e51b815260040161049c9061298e565b5f61054060a0850160808601612405565b6001600160a01b03163b116105b45760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b606482015260840161049c565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f8281526006602090815260408083208390558691610603918701908701612405565b6106116101008701876129d7565b6040516024016106249493929190612a19565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f61066b61066460a0870160808801612405565b5a846116fb565b9050806106ce5760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b606482015260840161049c565b50506001805550505050565b335f9081526009602090815260408083206001600160a01b03851684529091529020548061075b5760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b606482015260840161049c565b335f8181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a36107c96001600160a01b0383163383611712565b5050565b5f8181526004602052604081206107e390611776565b92915050565b5f818152600760205260408120546108555760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b606482015260840161049c565b505f908152600860205260409020546001600160a01b031690565b60015f54146108915760405162461bcd60e51b815260040161049c90612a43565b60025f81815590546108a9906060840135843561138b565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906109105760405162461bcd60e51b815260040161049c906126e5565b5f83604051602001610922919061297c565b60408051601f198184030181529190528251815160208301209192501461095b5760405162461bcd60e51b815260040161049c9061298e565b8360600135837f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1868560200151604051610996929190612a86565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c906109ea90879060040161297c565b5f604051808303815f87803b158015610a01575f5ffd5b505af1158015610a13573d5f5f3e3d5ffd5b505060015f55505050505050565b6001805414610a425760405162461bcd60e51b815260040161049c906126a0565b600260015560405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff0090610a92908590600401612aba565b6020604051808303815f875af1158015610aae573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad29190612b3d565b610b375760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465724d657373656e6765723a206d6573736167652076657260448201526f1a599a58d85d1a5bdb8819985a5b195960821b606482015260840161049c565b36610b428380612b5c565b905030610b556060830160408401612405565b6001600160a01b031614610bd15760405162461bcd60e51b815260206004820152603860248201527f54656c65706f727465724d657373656e67657256323a20696e76616c6964206f60448201527f726967696e2074656c65706f7274657220616464726573730000000000000000606482015260840161049c565b5f610bdf60e0830183612b7b565b808060200260200160405190810160405280939291908181526020015f905b82821015610c2a57610c1b60408302860136819003810190612c4e565b81526020019060010190610bfe565b505050505090505f604051806101000160405280845f01358152602001846020016020810190610c5a9190612405565b6001600160a01b0316815260608501356020820152604001610c8260a0860160808701612405565b6001600160a01b0316815260a08501356020820152604001610ca760c0860186612c84565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050509082525060208101849052604001610cf26101008601866129d7565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050915250600254604082015191925014610d995760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b606482015260840161049c565b5f610dae8660400135600254845f015161138b565b5f8181526007602052604090205490915015610e225760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b606482015260840161049c565b610e30338360a00151611788565b610e8e5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b606482015260840161049c565b610e9b81835f01516117f4565b6001600160a01b03851615610ed1575f81815260086020526040902080546001600160a01b0319166001600160a01b0387161790555b60c0820151515f5b81811015610f1757610f0f60025489604001358660c001518481518110610f0257610f02612cc9565b6020026020010151611864565b600101610ed9565b50604080518082018252845181526001600160a01b038816602080830191909152898301355f908152600490915291909120610f5291611988565b336001600160a01b03168760400135837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e348987604051610f93929190612e57565b60405180910390a460e08301515115610fb557610fb5828860400135856119e2565b5050600180555050505050565b5f60015f5414610fe45760405162461bcd60e51b815260040161049c90612a43565b60025f55611011610ff483612f98565b83355f90815260046020526040902061100c90611b11565b611c0b565b60015f5592915050565b604080518082019091525f80825260208201525f8381526004602052604090206110459083611e51565b9392505050565b60015f541461106d5760405162461bcd60e51b815260040161049c90612a43565b60025f5560018054146110925760405162461bcd60e51b815260040161049c906126a0565b6002600155806110fc5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b606482015260840161049c565b6001600160a01b0382166111225760405162461bcd60e51b815260040161049c9061303b565b5f8381526005602052604090205461114c5760405162461bcd60e51b815260040161049c906126e5565b5f838152600560205260409020600101546001600160a01b038381169116146111dd5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e74726163742061646472657373000000000000000000606482015260840161049c565b5f6111e88383611f12565b5f8581526005602052604081206002018054929350839290919061120d9084906130a3565b90915550505f8481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e19161126e916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156112c75750825b90505f826001600160401b031660011480156112e25750303b155b9050811580156112f0575080155b1561130e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561133857845460ff60401b1916600160401b1785555b6002869055831561138357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6040805130602082015290810184905260608101839052608081018290525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f60015f54146113ef5760405162461bcd60e51b815260040161049c90612a43565b60025f818155905490866001600160401b0381111561141057611410612bc0565b60405190808252806020026020018201604052801561145457816020015b604080518082019091525f808252602082015281526020019060019003908161142e5790505b509050865f5b818110156115cc575f8a8a8381811061147557611475612cc9565b9050602002013590505f60075f8381526020019081526020015f20549050805f036114f15760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b606482015260840161049c565b6114fc8d878361138b565b82146115705760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e000000000000606482015260840161049c565b5f828152600860209081526040918290205482518084019093528383526001600160a01b031690820181905286519091908790869081106115b3576115b3612cc9565b602002602001018190525050505080600101905061145a565b506040805160c0810182528b81525f602082015261164d9181016115f5368b90038b018b6130b6565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f9201829052509385525050604080519283526020808401909152909201525083611c0b565b60015f559a9950505050505050505050565b6002545f90806116c15760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b606482015260840161049c565b5f60035460016116d191906130a3565b90506116de82858361138b565b949350505050565b5f8181526007602052604081205415156107e3565b5f5f5f5f8451602086015f8989f195945050505050565b6040516001600160a01b0383811660248301526044820183905261177191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611f1e565b505050565b805460018201545f916107e3916130d0565b5f81515f03611799575060016107e3565b81515f5b818110156117ea57846001600160a01b03168482815181106117c1576117c1612cc9565b60200260200101516001600160a01b0316036117e2576001925050506107e3565b60010161179d565b505f949350505050565b805f036118535760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b606482015260840161049c565b5f9182526007602052604090912055565b5f6118738484845f015161138b565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906118c9575050505050565b5f8281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b039081168652600985528386209251511685529252822080549192909161192f9084906130a3565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57846020015160405161197991906130e3565b60405180910390a45050505050565b600182018054829160028501915f91826119a183613103565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a1015611a445760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b606482015260840161049c565b80606001516001600160a01b03163b5f03611a6457611771838383611f7f565b602081015160e08201516040515f92611a8192869260240161311b565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052606083015160808401519192505f91611ac59190846116fb565b905080611ade57611ad7858585611f7f565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b60605f611b276005611b2285611776565b611ff3565b9050805f03611b7357604080515f8082526020820190925290611b6b565b604080518082019091525f8082526020820152815260200190600190039081611b455790505b509392505050565b5f816001600160401b03811115611b8c57611b8c612bc0565b604051908082528060200260200182016040528015611bd057816020015b604080518082019091525f8082526020820152815260200190600190039081611baa5790505b5090505f5b82811015611b6b57611be685612008565b828281518110611bf857611bf8612cc9565b6020908102919091010152600101611bd5565b5f5f60035f8154611c1b90613103565b91905081905590505f611c34600254865f01518461138b565b90505f604051806101200160405280848152602001336001600160a01b03168152602001306001600160a01b03168152602001875f0151815260200187602001516001600160a01b0316815260200187606001518152602001876080015181526020018681526020018760a0015181525090505f81604051602001611cb991906131fe565b60405160208183030381529060405290505f5f8860400151602001511115611d20576040880151516001600160a01b0316611d065760405162461bcd60e51b815260040161049c9061303b565b60408801518051602090910151611d1d9190611f12565b90505b60408051808201825289820151516001600160a01b03908116825260208083018590528351808501855286518783012081528082018481525f8a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558951915190919086907f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d190611dc39088908690613210565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90611e179087906004016131fe565b5f604051808303815f87803b158015611e2e575f5ffd5b505af1158015611e40573d5f5f3e3d5ffd5b50969b9a5050505050505050505050565b604080518082019091525f8082526020820152611e6d83611776565b8210611ec55760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b606482015260840161049c565b826002015f83855f0154611ed991906130a3565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b5f6110458333846120d2565b5f611f326001600160a01b03841683612235565b905080515f14158015611f56575080806020019051810190611f549190612b3d565b155b1561177157604051635274afe760e01b81526001600160a01b038416600482015260240161049c565b80604051602001611f909190613222565b60408051601f1981840301815282825280516020918201205f878152600690925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c98590611fe6908590613222565b60405180910390a3505050565b5f8183106120015781611045565b5090919050565b604080518082019091525f8082526020820152815460018301548190036120715760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d70747920717565756500000000000000604482015260640161049c565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b03199092169055906120c89083906130a3565b9093555090919050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015612118573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061213c9190613234565b90506121536001600160a01b038616853086612242565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015612197573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121bb9190613234565b90508181116122215760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161049c565b61222b82826130d0565b9695505050505050565b606061104583835f612281565b6040516001600160a01b03848116602483015283811660448301526064820183905261227b9186918216906323b872dd9060840161173f565b50505050565b6060814710156122a65760405163cd78605960e01b815230600482015260240161049c565b5f5f856001600160a01b031684866040516122c1919061324b565b5f6040518083038185875af1925050503d805f81146122fb576040519150601f19603f3d011682016040523d82523d5f602084013e612300565b606091505b509150915061222b8683836060826123205761231b82612367565b611045565b815115801561233757506001600160a01b0384163b155b1561236057604051639996b31560e01b81526001600160a01b038516600482015260240161049c565b5080611045565b8051156123775780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f61012082840312156123a1575f5ffd5b50919050565b5f5f604083850312156123b8575f5ffd5b8235915060208301356001600160401b038111156123d4575f5ffd5b6123e085828601612390565b9150509250929050565b80356001600160a01b0381168114612400575f5ffd5b919050565b5f60208284031215612415575f5ffd5b611045826123ea565b5f6020828403121561242e575f5ffd5b5035919050565b82815260608101611045602083018480516001600160a01b03168252602090810151910152565b5f6020828403121561246c575f5ffd5b81356001600160401b03811115612481575f5ffd5b6116de84828501612390565b5f5f6040838503121561249e575f5ffd5b82356001600160401b038111156124b3575f5ffd5b8301608081860312156124c4575f5ffd5b91506124d2602084016123ea565b90509250929050565b5f602082840312156124eb575f5ffd5b81356001600160401b03811115612500575f5ffd5b820160e08185031215611045575f5ffd5b5f5f60408385031215612522575f5ffd5b50508035926020909101359150565b815181526020808301516001600160a01b031690820152604081016107e3565b5f5f5f60608486031215612563575f5ffd5b83359250612573602085016123ea565b929592945050506040919091013590565b5f5f5f60608486031215612596575f5ffd5b505081359360208301359350604090920135919050565b5f5f83601f8401126125bd575f5ffd5b5081356001600160401b038111156125d3575f5ffd5b6020830191508360208260051b85010111156125ed575f5ffd5b9250929050565b5f5f5f5f5f5f86880360a081121561260a575f5ffd5b8735965060208801356001600160401b03811115612626575f5ffd5b6126328a828b016125ad565b9097509550506040603f1982011215612649575f5ffd5b5060408701925060808701356001600160401b03811115612668575f5ffd5b61267489828a016125ad565b979a9699509497509295939492505050565b5f5f60408385031215612697575f5ffd5b6124c4836123ea565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f5f8335601e19843603018112612740575f5ffd5b83016020810192503590506001600160401b0381111561275e575f5ffd5b8060051b36038213156125ed575f5ffd5b8183526020830192505f815f5b848110156127ab576001600160a01b03612795836123ea565b168652602095860195919091019060010161277c565b5093949350505050565b5f5f8335601e198436030181126127ca575f5ffd5b83016020810192503590506001600160401b038111156127e8575f5ffd5b8060061b36038213156125ed575f5ffd5b8183526020830192505f815f5b848110156127ab57813586526001600160a01b03612826602084016123ea565b1660208701526040958601959190910190600101612806565b5f5f8335601e19843603018112612854575f5ffd5b83016020810192503590506001600160401b03811115612872575f5ffd5b8036038213156125ed575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f6128b9602083016123ea565b6001600160a01b031660208401526128d3604083016123ea565b6001600160a01b03166040840152606082810135908401526128f7608083016123ea565b6001600160a01b0316608084015260a0828101359084015261291c60c083018361272b565b61012060c08601526129336101208601828461276f565b91505061294360e08401846127b5565b85830360e08701526129568382846127f9565b9250505061296861010084018461283f565b85830361010087015261222b838284612880565b602081525f61104560208301846128a8565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b5f5f8335601e198436030181126129ec575f5ffd5b8301803591506001600160401b03821115612a05575f5ffd5b6020019150368190038213156125ed575f5ffd5b8481526001600160a01b03841660208201526060604082018190525f9061222b9083018486612880565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b606081525f612a9860608301856128a8565b9050611045602083018480516001600160a01b03168252602090810151910152565b602081525f823561011e19843603018112612ad3575f5ffd5b60806020840152612ae960a084018583016128a8565b9050602084013563ffffffff8116808214612b02575f5ffd5b80604086015250505f6040850135905080606085015250612b26606085018561283f565b848303601f1901608086015261222b838284612880565b5f60208284031215612b4d575f5ffd5b81518015158114611045575f5ffd5b5f823561011e19833603018112612b71575f5ffd5b9190910192915050565b5f5f8335601e19843603018112612b90575f5ffd5b8301803591506001600160401b03821115612ba9575f5ffd5b6020019150600681901b36038213156125ed575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715612bf657612bf6612bc0565b60405290565b60405160c081016001600160401b0381118282101715612bf657612bf6612bc0565b604051601f8201601f191681016001600160401b0381118282101715612c4657612c46612bc0565b604052919050565b5f6040828403128015612c5f575f5ffd5b50612c68612bd4565b82358152612c78602084016123ea565b60208201529392505050565b5f5f8335601e19843603018112612c99575f5ffd5b8301803591506001600160401b03821115612cb2575f5ffd5b6020019150600581901b36038213156125ed575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b828110156127ab5781516001600160a01b0316865260209586019590910190600101612cef565b5f8151808452602084019350602083015f5b828110156127ab57612d4e868351805182526020908101516001600160a01b0316910152565b6040959095019460209190910190600101612d28565b5f5b83811015612d7e578181015183820152602001612d66565b50505f910152565b5f8151808452612d9d816020860160208601612d64565b601f01601f19169290920160200192915050565b805182525f6020820151612dd060208501826001600160a01b03169052565b50604082015160408401526060820151612df560608501826001600160a01b03169052565b506080820151608084015260a082015161010060a0850152612e1b610100850182612cdd565b905060c083015184820360c0860152612e348282612d16565b91505060e083015184820360e0860152612e4e8282612d86565b95945050505050565b6001600160a01b03831681526040602082018190525f906116de90830184612db1565b5f60408284031215612e8a575f5ffd5b612e92612bd4565b9050612e9d826123ea565b815260209182013591810191909152919050565b5f82601f830112612ec0575f5ffd5b81356001600160401b03811115612ed957612ed9612bc0565b8060051b612ee960208201612c1e565b91825260208185018101929081019086841115612f04575f5ffd5b6020860192505b8383101561222b57612f1c836123ea565b825260209283019290910190612f0b565b5f82601f830112612f3c575f5ffd5b81356001600160401b03811115612f5557612f55612bc0565b612f68601f8201601f1916602001612c1e565b818152846020838601011115612f7c575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60e08236031215612fa8575f5ffd5b612fb0612bfc565b82358152612fc0602084016123ea565b6020820152612fd23660408501612e7a565b60408201526080830135606082015260a08301356001600160401b03811115612ff9575f5ffd5b61300536828601612eb1565b60808301525060c08301356001600160401b03811115613023575f5ffd5b61302f36828601612f2d565b60a08301525092915050565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156107e3576107e361308f565b5f604082840312156130c6575f5ffd5b6110458383612e7a565b818103818111156107e3576107e361308f565b81516001600160a01b0316815260208083015190820152604081016107e3565b5f600182016131145761311461308f565b5060010190565b8381526001600160a01b03831660208201526060604082018190525f90612e4e90830184612d86565b805182525f602082015161316360208501826001600160a01b03169052565b50604082015161317e60408501826001600160a01b03169052565b506060820151606084015260808201516131a360808501826001600160a01b03169052565b5060a082015160a084015260c082015161012060c08501526131c9610120850182612cdd565b905060e083015184820360e08601526131e28282612d16565b915050610100830151848203610100860152612e4e8282612d86565b602081525f6110456020830184613144565b606081525f612a986060830185613144565b602081525f6110456020830184612db1565b5f60208284031215613244575f5ffd5b5051919050565b5f8251612b71818460208701612d6456fea2646970667358221220fa785d37e568298b3cf5a1039a9fbf92f6d0ee5e16597fd0ba4f35f4bf581c7e64736f6c634300081e0033",
}

// TeleporterMessengerV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerV2MetaData.ABI instead.
var TeleporterMessengerV2ABI = TeleporterMessengerV2MetaData.ABI

// TeleporterMessengerV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterMessengerV2MetaData.Bin instead.
var TeleporterMessengerV2Bin = TeleporterMessengerV2MetaData.Bin

// DeployTeleporterMessengerV2 deploys a new Ethereum contract, binding an instance of TeleporterMessengerV2 to it.
func DeployTeleporterMessengerV2(auth *bind.TransactOpts, backend bind.ContractBackend, verifierSender common.Address) (common.Address, *types.Transaction, *TeleporterMessengerV2, error) {
	parsed, err := TeleporterMessengerV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterMessengerV2Bin), backend, verifierSender)
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
