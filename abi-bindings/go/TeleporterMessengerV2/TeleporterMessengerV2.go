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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierSender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"contractIMessageSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVerifier\",\"outputs\":[{\"internalType\":\"contractIMessageVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516132c03803806132c083398101604081905261002e916100c1565b60015f81905580556040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015610084573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a891906100ee565b6002556001600160a01b031660a0819052608052610105565b5f602082840312156100d1575f5ffd5b81516001600160a01b03811681146100e7575f5ffd5b9392505050565b5f602082840312156100fe575f5ffd5b5051919050565b60805160a05161318561013b5f395f81816102d70152610a5801525f8181610373015281816109b00152611cd501526131855ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c80638ae9753d116100bf578063d67bdd2511610079578063d67bdd251461036e578063df20e8bc14610395578063e69d606a146103a8578063e6e67bd51461040f578063ebc3b1ba1461044a578063ecc704281461046d575f5ffd5b80638ae9753d146102d2578063a8898181146102f9578063a9a856141461030c578063b771b3bc1461031f578063c473eef81461032d578063d127dc9b14610365575f5ffd5b80633ba2b983116101105780633ba2b983146102475780633e249b561461025a578063624488501461026d578063860a3b0614610280578063892bf4121461029f5780638ac0fd04146102bf575f5ffd5b80630f635b6c1461015757806322296c3a1461016c5780632bc8b0bf1461017f5780632ca40f55146101a55780632e27c223146101fd578063399b77da14610228575b5f5ffd5b61016a61016536600461229a565b610476565b005b61016a61017a3660046122f8565b6106d5565b61019261018d366004612311565b6107c8565b6040519081526020015b60405180910390f35b6101ef6101b3366004612311565b600560209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b60405161019c929190612328565b61021061020b366004612311565b6107e4565b6040516001600160a01b03909116815260200161019c565b610192610236366004612311565b5f9081526005602052604090205490565b61016a61025536600461234f565b61086b565b61016a610268366004612380565b610a1c565b61019261027b3660046123ce565b610fbd565b61019261028e366004612311565b60066020525f908152604090205481565b6102b26102ad366004612404565b611016565b60405161019c9190612424565b61016a6102cd366004612444565b611047565b6102107f000000000000000000000000000000000000000000000000000000000000000081565b610192610307366004612477565b61127e565b61019261031a3660046124e7565b6112c0565b6102106005600160991b0181565b61019261033b366004612579565b6001600160a01b039182165f90815260096020908152604080832093909416825291909152205490565b61019260025481565b6102107f000000000000000000000000000000000000000000000000000000000000000081565b6101926103a3366004612311565b611552565b6103f06103b6366004612311565b5f90815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b03909316835260208301919091520161019c565b61043561041d366004612311565b60046020525f90815260409020805460019091015482565b6040805192835260208301919091520161019c565b61045d610458366004612311565b6115d9565b604051901515815260200161019c565b61019260035481565b60018054146104a05760405162461bcd60e51b815260040161049790612593565b60405180910390fd5b60026001819055545f906104b7908490843561127e565b5f81815260066020526040902054909150806104e55760405162461bcd60e51b8152600401610497906125d8565b80836040516020016104f7919061286f565b604051602081830303815290604052805190602001201461052a5760405162461bcd60e51b815260040161049790612881565b5f61053b60a08501608086016122f8565b6001600160a01b03163b116105af5760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b6064820152608401610497565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f82815260066020908152604080832083905586916105fe9187019087016122f8565b61060c6101008701876128ca565b60405160240161061f949392919061290c565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f61066661065f60a08701608088016122f8565b5a846115ee565b9050806106c95760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b6064820152608401610497565b50506001805550505050565b335f9081526009602090815260408083206001600160a01b0385168452909152902054806107565760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b6064820152608401610497565b335f8181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a36107c46001600160a01b0383163383611605565b5050565b5f8181526004602052604081206107de90611669565b92915050565b5f818152600760205260408120546108505760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b6064820152608401610497565b505f908152600860205260409020546001600160a01b031690565b60015f541461088c5760405162461bcd60e51b815260040161049790612936565b60025f81815590546108a4906060840135843561127e565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b03168452600290910154838301529081019190915280519192509061090b5760405162461bcd60e51b8152600401610497906125d8565b5f8360405160200161091d919061286f565b60408051601f19818403018152919052825181516020830120919250146109565760405162461bcd60e51b815260040161049790612881565b8360600135837f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1868560200151604051610991929190612979565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c906109e590879060040161286f565b5f604051808303815f87803b1580156109fc575f5ffd5b505af1158015610a0e573d5f5f3e3d5ffd5b505060015f55505050505050565b6001805414610a3d5760405162461bcd60e51b815260040161049790612593565b600260015560405162f1faff60e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f1faff0090610a8d9085906004016129ad565b6020604051808303815f875af1158015610aa9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610acd9190612a30565b610b325760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465724d657373656e6765723a206d6573736167652076657260448201526f1a599a58d85d1a5bdb8819985a5b195960821b6064820152608401610497565b36610b3d8380612a4f565b905030610b5060608301604084016122f8565b6001600160a01b031614610bcc5760405162461bcd60e51b815260206004820152603860248201527f54656c65706f727465724d657373656e67657256323a20696e76616c6964206f60448201527f726967696e2074656c65706f72746572206164647265737300000000000000006064820152608401610497565b5f610bda60e0830183612a6e565b808060200260200160405190810160405280939291908181526020015f905b82821015610c2557610c1660408302860136819003810190612b41565b81526020019060010190610bf9565b505050505090505f604051806101000160405280845f01358152602001846020016020810190610c5591906122f8565b6001600160a01b0316815260608501356020820152604001610c7d60a08601608087016122f8565b6001600160a01b0316815260a08501356020820152604001610ca260c0860186612b77565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050509082525060208101849052604001610ced6101008601866128ca565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050915250600254604082015191925014610d945760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b6064820152608401610497565b5f610da98660400135600254845f015161127e565b5f8181526007602052604090205490915015610e1d5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b6064820152608401610497565b610e2b338360a0015161167b565b610e895760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b6064820152608401610497565b610e9681835f01516116e7565b6001600160a01b03851615610ecc575f81815260086020526040902080546001600160a01b0319166001600160a01b0387161790555b60c0820151515f5b81811015610f1257610f0a60025489604001358660c001518481518110610efd57610efd612bbc565b6020026020010151611757565b600101610ed4565b50604080518082018252845181526001600160a01b038816602080830191909152898301355f908152600490915291909120610f4d9161187b565b336001600160a01b03168760400135837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e348987604051610f8e929190612d4a565b60405180910390a460e08301515115610fb057610fb0828860400135856118d5565b5050600180555050505050565b5f60015f5414610fdf5760405162461bcd60e51b815260040161049790612936565b60025f5561100c610fef83612e8b565b83355f90815260046020526040902061100790611a04565b611afe565b60015f5592915050565b604080518082019091525f80825260208201525f8381526004602052604090206110409083611d44565b9392505050565b60015f54146110685760405162461bcd60e51b815260040161049790612936565b60025f55600180541461108d5760405162461bcd60e51b815260040161049790612593565b6002600155806110f75760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b6064820152608401610497565b6001600160a01b03821661111d5760405162461bcd60e51b815260040161049790612f2e565b5f838152600560205260409020546111475760405162461bcd60e51b8152600401610497906125d8565b5f838152600560205260409020600101546001600160a01b038381169116146111d85760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e747261637420616464726573730000000000000000006064820152608401610497565b5f6111e38383611e05565b5f85815260056020526040812060020180549293508392909190611208908490612f96565b90915550505f8481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e191611269916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b6040805130602082015290810184905260608101839052608081018290525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f60015f54146112e25760405162461bcd60e51b815260040161049790612936565b60025f818155905490866001600160401b0381111561130357611303612ab3565b60405190808252806020026020018201604052801561134757816020015b604080518082019091525f80825260208201528152602001906001900390816113215790505b509050865f5b818110156114bf575f8a8a8381811061136857611368612bbc565b9050602002013590505f60075f8381526020019081526020015f20549050805f036113e45760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b6064820152608401610497565b6113ef8d878361127e565b82146114635760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e0000000000006064820152608401610497565b5f828152600860209081526040918290205482518084019093528383526001600160a01b031690820181905286519091908790869081106114a6576114a6612bbc565b602002602001018190525050505080600101905061134d565b506040805160c0810182528b81525f60208201526115409181016114e8368b90038b018b612fa9565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f9201829052509385525050604080519283526020808401909152909201525083611afe565b60015f559a9950505050505050505050565b6002545f90806115b45760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b6064820152608401610497565b5f60035460016115c49190612f96565b90506115d182858361127e565b949350505050565b5f8181526007602052604081205415156107de565b5f5f5f5f8451602086015f8989f195945050505050565b6040516001600160a01b0383811660248301526044820183905261166491859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611e11565b505050565b805460018201545f916107de91612fc3565b5f81515f0361168c575060016107de565b81515f5b818110156116dd57846001600160a01b03168482815181106116b4576116b4612bbc565b60200260200101516001600160a01b0316036116d5576001925050506107de565b600101611690565b505f949350505050565b805f036117465760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b6064820152608401610497565b5f9182526007602052604090912055565b5f6117668484845f015161127e565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906117bc575050505050565b5f8281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b0390811686526009855283862092515116855292528220805491929091611822908490612f96565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57846020015160405161186c9190612fd6565b60405180910390a45050505050565b600182018054829160028501915f918261189483612ff6565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a10156119375760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b6064820152608401610497565b80606001516001600160a01b03163b5f0361195757611664838383611e72565b602081015160e08201516040515f9261197492869260240161300e565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052606083015160808401519192505f916119b89190846115ee565b9050806119d1576119ca858585611e72565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b60605f611a1a6005611a1585611669565b611ee6565b9050805f03611a6657604080515f8082526020820190925290611a5e565b604080518082019091525f8082526020820152815260200190600190039081611a385790505b509392505050565b5f816001600160401b03811115611a7f57611a7f612ab3565b604051908082528060200260200182016040528015611ac357816020015b604080518082019091525f8082526020820152815260200190600190039081611a9d5790505b5090505f5b82811015611a5e57611ad985611efb565b828281518110611aeb57611aeb612bbc565b6020908102919091010152600101611ac8565b5f5f60035f8154611b0e90612ff6565b91905081905590505f611b27600254865f01518461127e565b90505f604051806101200160405280848152602001336001600160a01b03168152602001306001600160a01b03168152602001875f0151815260200187602001516001600160a01b0316815260200187606001518152602001876080015181526020018681526020018760a0015181525090505f81604051602001611bac91906130f1565b60405160208183030381529060405290505f5f8860400151602001511115611c13576040880151516001600160a01b0316611bf95760405162461bcd60e51b815260040161049790612f2e565b60408801518051602090910151611c109190611e05565b90505b60408051808201825289820151516001600160a01b03908116825260208083018590528351808501855286518783012081528082018481525f8a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558951915190919086907f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d190611cb69088908690613103565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90611d0a9087906004016130f1565b5f604051808303815f87803b158015611d21575f5ffd5b505af1158015611d33573d5f5f3e3d5ffd5b50969b9a5050505050505050505050565b604080518082019091525f8082526020820152611d6083611669565b8210611db85760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b6064820152608401610497565b826002015f83855f0154611dcc9190612f96565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b5f611040833384611fc5565b5f611e256001600160a01b03841683612128565b905080515f14158015611e49575080806020019051810190611e479190612a30565b155b1561166457604051635274afe760e01b81526001600160a01b0384166004820152602401610497565b80604051602001611e839190613115565b60408051601f1981840301815282825280516020918201205f878152600690925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c98590611ed9908590613115565b60405180910390a3505050565b5f818310611ef45781611040565b5090919050565b604080518082019091525f808252602082015281546001830154819003611f645760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d707479207175657565000000000000006044820152606401610497565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b0319909216905590611fbb908390612f96565b9093555090919050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa15801561200b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061202f9190613127565b90506120466001600160a01b038616853086612135565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa15801561208a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120ae9190613127565b90508181116121145760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610497565b61211e8282612fc3565b9695505050505050565b606061104083835f612174565b6040516001600160a01b03848116602483015283811660448301526064820183905261216e9186918216906323b872dd90608401611632565b50505050565b6060814710156121995760405163cd78605960e01b8152306004820152602401610497565b5f5f856001600160a01b031684866040516121b4919061313e565b5f6040518083038185875af1925050503d805f81146121ee576040519150601f19603f3d011682016040523d82523d5f602084013e6121f3565b606091505b509150915061211e8683836060826122135761220e8261225a565b611040565b815115801561222a57506001600160a01b0384163b155b1561225357604051639996b31560e01b81526001600160a01b0385166004820152602401610497565b5080611040565b80511561226a5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f6101208284031215612294575f5ffd5b50919050565b5f5f604083850312156122ab575f5ffd5b8235915060208301356001600160401b038111156122c7575f5ffd5b6122d385828601612283565b9150509250929050565b80356001600160a01b03811681146122f3575f5ffd5b919050565b5f60208284031215612308575f5ffd5b611040826122dd565b5f60208284031215612321575f5ffd5b5035919050565b82815260608101611040602083018480516001600160a01b03168252602090810151910152565b5f6020828403121561235f575f5ffd5b81356001600160401b03811115612374575f5ffd5b6115d184828501612283565b5f5f60408385031215612391575f5ffd5b82356001600160401b038111156123a6575f5ffd5b8301608081860312156123b7575f5ffd5b91506123c5602084016122dd565b90509250929050565b5f602082840312156123de575f5ffd5b81356001600160401b038111156123f3575f5ffd5b820160e08185031215611040575f5ffd5b5f5f60408385031215612415575f5ffd5b50508035926020909101359150565b815181526020808301516001600160a01b031690820152604081016107de565b5f5f5f60608486031215612456575f5ffd5b83359250612466602085016122dd565b929592945050506040919091013590565b5f5f5f60608486031215612489575f5ffd5b505081359360208301359350604090920135919050565b5f5f83601f8401126124b0575f5ffd5b5081356001600160401b038111156124c6575f5ffd5b6020830191508360208260051b85010111156124e0575f5ffd5b9250929050565b5f5f5f5f5f5f86880360a08112156124fd575f5ffd5b8735965060208801356001600160401b03811115612519575f5ffd5b6125258a828b016124a0565b9097509550506040603f198201121561253c575f5ffd5b5060408701925060808701356001600160401b0381111561255b575f5ffd5b61256789828a016124a0565b979a9699509497509295939492505050565b5f5f6040838503121561258a575f5ffd5b6123b7836122dd565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f5f8335601e19843603018112612633575f5ffd5b83016020810192503590506001600160401b03811115612651575f5ffd5b8060051b36038213156124e0575f5ffd5b8183526020830192505f815f5b8481101561269e576001600160a01b03612688836122dd565b168652602095860195919091019060010161266f565b5093949350505050565b5f5f8335601e198436030181126126bd575f5ffd5b83016020810192503590506001600160401b038111156126db575f5ffd5b8060061b36038213156124e0575f5ffd5b8183526020830192505f815f5b8481101561269e57813586526001600160a01b03612719602084016122dd565b16602087015260409586019591909101906001016126f9565b5f5f8335601e19843603018112612747575f5ffd5b83016020810192503590506001600160401b03811115612765575f5ffd5b8036038213156124e0575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f6127ac602083016122dd565b6001600160a01b031660208401526127c6604083016122dd565b6001600160a01b03166040840152606082810135908401526127ea608083016122dd565b6001600160a01b0316608084015260a0828101359084015261280f60c083018361261e565b61012060c086015261282661012086018284612662565b91505061283660e08401846126a8565b85830360e08701526128498382846126ec565b9250505061285b610100840184612732565b85830361010087015261211e838284612773565b602081525f611040602083018461279b565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b5f5f8335601e198436030181126128df575f5ffd5b8301803591506001600160401b038211156128f8575f5ffd5b6020019150368190038213156124e0575f5ffd5b8481526001600160a01b03841660208201526060604082018190525f9061211e9083018486612773565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b606081525f61298b606083018561279b565b9050611040602083018480516001600160a01b03168252602090810151910152565b602081525f823561011e198436030181126129c6575f5ffd5b608060208401526129dc60a0840185830161279b565b9050602084013563ffffffff81168082146129f5575f5ffd5b80604086015250505f6040850135905080606085015250612a196060850185612732565b848303601f1901608086015261211e838284612773565b5f60208284031215612a40575f5ffd5b81518015158114611040575f5ffd5b5f823561011e19833603018112612a64575f5ffd5b9190910192915050565b5f5f8335601e19843603018112612a83575f5ffd5b8301803591506001600160401b03821115612a9c575f5ffd5b6020019150600681901b36038213156124e0575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715612ae957612ae9612ab3565b60405290565b60405160c081016001600160401b0381118282101715612ae957612ae9612ab3565b604051601f8201601f191681016001600160401b0381118282101715612b3957612b39612ab3565b604052919050565b5f6040828403128015612b52575f5ffd5b50612b5b612ac7565b82358152612b6b602084016122dd565b60208201529392505050565b5f5f8335601e19843603018112612b8c575f5ffd5b8301803591506001600160401b03821115612ba5575f5ffd5b6020019150600581901b36038213156124e0575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f8151808452602084019350602083015f5b8281101561269e5781516001600160a01b0316865260209586019590910190600101612be2565b5f8151808452602084019350602083015f5b8281101561269e57612c41868351805182526020908101516001600160a01b0316910152565b6040959095019460209190910190600101612c1b565b5f5b83811015612c71578181015183820152602001612c59565b50505f910152565b5f8151808452612c90816020860160208601612c57565b601f01601f19169290920160200192915050565b805182525f6020820151612cc360208501826001600160a01b03169052565b50604082015160408401526060820151612ce860608501826001600160a01b03169052565b506080820151608084015260a082015161010060a0850152612d0e610100850182612bd0565b905060c083015184820360c0860152612d278282612c09565b91505060e083015184820360e0860152612d418282612c79565b95945050505050565b6001600160a01b03831681526040602082018190525f906115d190830184612ca4565b5f60408284031215612d7d575f5ffd5b612d85612ac7565b9050612d90826122dd565b815260209182013591810191909152919050565b5f82601f830112612db3575f5ffd5b81356001600160401b03811115612dcc57612dcc612ab3565b8060051b612ddc60208201612b11565b91825260208185018101929081019086841115612df7575f5ffd5b6020860192505b8383101561211e57612e0f836122dd565b825260209283019290910190612dfe565b5f82601f830112612e2f575f5ffd5b81356001600160401b03811115612e4857612e48612ab3565b612e5b601f8201601f1916602001612b11565b818152846020838601011115612e6f575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f60e08236031215612e9b575f5ffd5b612ea3612aef565b82358152612eb3602084016122dd565b6020820152612ec53660408501612d6d565b60408201526080830135606082015260a08301356001600160401b03811115612eec575f5ffd5b612ef836828601612da4565b60808301525060c08301356001600160401b03811115612f16575f5ffd5b612f2236828601612e20565b60a08301525092915050565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156107de576107de612f82565b5f60408284031215612fb9575f5ffd5b6110408383612d6d565b818103818111156107de576107de612f82565b81516001600160a01b0316815260208083015190820152604081016107de565b5f6001820161300757613007612f82565b5060010190565b8381526001600160a01b03831660208201526060604082018190525f90612d4190830184612c79565b805182525f602082015161305660208501826001600160a01b03169052565b50604082015161307160408501826001600160a01b03169052565b5060608201516060840152608082015161309660808501826001600160a01b03169052565b5060a082015160a084015260c082015161012060c08501526130bc610120850182612bd0565b905060e083015184820360e08601526130d58282612c09565b915050610100830151848203610100860152612d418282612c79565b602081525f6110406020830184613037565b606081525f61298b6060830185613037565b602081525f6110406020830184612ca4565b5f60208284031215613137575f5ffd5b5051919050565b5f8251612a64818460208701612c5756fea264697066735822122089170ecdf5087a0b63fb68aeb060166115a3e445845d2bd886e193c3db44354564736f6c634300081e0033",
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

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Caller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessengerV2.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2Session) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.WARPMESSENGER(&_TeleporterMessengerV2.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessengerV2 *TeleporterMessengerV2CallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessengerV2.Contract.WARPMESSENGER(&_TeleporterMessengerV2.CallOpts)
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
