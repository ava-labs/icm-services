// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessenger

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
	UnsignedMessage    TeleporterMessage
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
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

// TeleporterMessengerMetaData contains all meta data concerning the TeleporterMessenger contract.
var TeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierSender\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"BlockchainIDInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"contractIMessageSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVerifier\",\"outputs\":[{\"internalType\":\"contractIMessageVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"unsignedMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161321b38038061321b83398101604081905261002e916100c1565b60015f81905580556040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015610084573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a891906100ee565b6002556001600160a01b031660a0819052608052610105565b5f602082840312156100d1575f5ffd5b81516001600160a01b03811681146100e7575f5ffd5b9392505050565b5f602082840312156100fe575f5ffd5b5051919050565b60805160a0516130e061013b5f395f81816102d7015261071201525f818161037301528181610dd70152611b5c01526130e05ff3fe608060405234801561000f575f5ffd5b5060043610610153575f3560e01c80638ae9753d116100bf578063d67bdd2511610079578063d67bdd251461036e578063df20e8bc14610395578063e69d606a146103a8578063e6e67bd51461040f578063ebc3b1ba1461044a578063ecc704281461046d575f5ffd5b80638ae9753d146102d2578063a8898181146102f9578063a9a856141461030c578063b771b3bc1461031f578063c473eef81461032d578063d127dc9b14610365575f5ffd5b8063399b77da11610110578063399b77da1461023b5780633ba2b9831461025a578063624488501461026d578063860a3b0614610280578063892bf4121461029f5780638ac0fd04146102bf575f5ffd5b80630f635b6c146101575780631c54d5351461016c57806322296c3a1461017f5780632bc8b0bf146101925780632ca40f55146101b85780632e27c22314610210575b5f5ffd5b61016a610165366004612121565b610476565b005b61016a61017a36600461217f565b6106d5565b61016a61018d3660046121cd565b610afc565b6101a56101a03660046121e6565b610bef565b6040519081526020015b60405180910390f35b6102026101c63660046121e6565b600560209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b6040516101af9291906121fd565b61022361021e3660046121e6565b610c0b565b6040516001600160a01b0390911681526020016101af565b6101a56102493660046121e6565b5f9081526005602052604090205490565b61016a610268366004612224565b610c92565b6101a561027b366004612255565b610e43565b6101a561028e3660046121e6565b60066020525f908152604090205481565b6102b26102ad36600461228b565b610e9c565b6040516101af91906122ab565b61016a6102cd3660046122cb565b610ecd565b6102237f000000000000000000000000000000000000000000000000000000000000000081565b6101a56103073660046122fe565b611104565b6101a561031a36600461236e565b611146565b6102236005600160991b0181565b6101a561033b366004612400565b6001600160a01b039182165f90815260096020908152604080832093909416825291909152205490565b6101a560025481565b6102237f000000000000000000000000000000000000000000000000000000000000000081565b6101a56103a33660046121e6565b6113d8565b6103f06103b63660046121e6565b5f90815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b0390931683526020830191909152016101af565b61043561041d3660046121e6565b60046020525f90815260409020805460019091015482565b604080519283526020830191909152016101af565b61045d6104583660046121e6565b61145f565b60405190151581526020016101af565b6101a560035481565b60018054146104a05760405162461bcd60e51b81526004016104979061241a565b60405180910390fd5b60026001819055545f906104b79084908435611104565b5f81815260066020526040902054909150806104e55760405162461bcd60e51b81526004016104979061245f565b80836040516020016104f791906126f6565b604051602081830303815290604052805190602001201461052a5760405162461bcd60e51b815260040161049790612708565b5f61053b60a08501608086016121cd565b6001600160a01b03163b116105af5760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b6064820152608401610497565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f82815260066020908152604080832083905586916105fe9187019087016121cd565b61060c610100870187612751565b60405160240161061f9493929190612793565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f61066661065f60a08701608088016121cd565b5a84611474565b9050806106c95760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b6064820152608401610497565b50506001805550505050565b60018054146106f65760405162461bcd60e51b81526004016104979061241a565b600260015560405163b85958cf60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b85958cf906107479085906004016127bd565b6020604051808303815f875af1158015610763573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107879190612821565b6107ec5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465724d657373656e6765723a206d6573736167652076657260448201526f1a599a58d85d1a5bdb8819985a5b195960821b6064820152608401610497565b366107f78380612840565b90506002548160600135146108685760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b6064820152608401610497565b5f61087d8460200135600254845f0135611104565b5f81815260076020526040902054909150156108f15760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b6064820152608401610497565b6109383361090260c085018561285f565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061148b92505050565b6109965760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b6064820152608401610497565b6109a18183356114f7565b6001600160a01b038316156109d7575f81815260086020526040902080546001600160a01b0319166001600160a01b0385161790555b5f6109e560e08401846128a4565b905090505f5b81811015610a3f57600254610a37906020880135610a0c60e08801886128a4565b85818110610a1c57610a1c6128e9565b905060400201803603810190610a3291906129e3565b611567565b6001016109eb565b50604080518082018252843581526001600160a01b038616602080830191909152878101355f908152600490915291909120610a7a9161168b565b336001600160a01b03168560200135837f51f9dcb4f762bab1ca8136939fec0e5a257efbd0f69dea14b60d6d1ae0750f808787604051610abb9291906129fd565b60405180910390a45f610ad2610100850185612751565b90501115610af157610af1826020870135610aec86612b83565b6116e5565b505060018055505050565b335f9081526009602090815260408083206001600160a01b038516845290915290205480610b7d5760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b6064820152608401610497565b335f8181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a3610beb6001600160a01b038316338361181a565b5050565b5f818152600460205260408120610c0590611879565b92915050565b5f81815260076020526040812054610c775760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b6064820152608401610497565b505f908152600860205260409020546001600160a01b031690565b60015f5414610cb35760405162461bcd60e51b815260040161049790612c6d565b60025f8181559054610ccb9060608401358435611104565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b031684526002909101548383015290810191909152805191925090610d325760405162461bcd60e51b81526004016104979061245f565b5f83604051602001610d4491906126f6565b60408051601f1981840301815291905282518151602083012091925014610d7d5760405162461bcd60e51b815260040161049790612708565b8360600135837f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1868560200151604051610db8929190612cb0565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90610e0c9087906004016126f6565b5f604051808303815f87803b158015610e23575f5ffd5b505af1158015610e35573d5f5f3e3d5ffd5b505060015f55505050505050565b5f60015f5414610e655760405162461bcd60e51b815260040161049790612c6d565b60025f55610e92610e7583612d1b565b83355f908152600460205260409020610e8d9061188b565b611985565b60015f5592915050565b604080518082019091525f80825260208201525f838152600460205260409020610ec69083611bcb565b9392505050565b60015f5414610eee5760405162461bcd60e51b815260040161049790612c6d565b60025f556001805414610f135760405162461bcd60e51b81526004016104979061241a565b600260015580610f7d5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b6064820152608401610497565b6001600160a01b038216610fa35760405162461bcd60e51b815260040161049790612dbe565b5f83815260056020526040902054610fcd5760405162461bcd60e51b81526004016104979061245f565b5f838152600560205260409020600101546001600160a01b0383811691161461105e5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e747261637420616464726573730000000000000000006064820152608401610497565b5f6110698383611c8c565b5f8581526005602052604081206002018054929350839290919061108e908490612e26565b90915550505f8481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1916110ef916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b6040805130602082015290810184905260608101839052608081018290525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f60015f54146111685760405162461bcd60e51b815260040161049790612c6d565b60025f818155905490866001600160401b03811115611189576111896128fd565b6040519080825280602002602001820160405280156111cd57816020015b604080518082019091525f80825260208201528152602001906001900390816111a75790505b509050865f5b81811015611345575f8a8a838181106111ee576111ee6128e9565b9050602002013590505f60075f8381526020019081526020015f20549050805f0361126a5760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b6064820152608401610497565b6112758d8783611104565b82146112e95760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e0000000000006064820152608401610497565b5f828152600860209081526040918290205482518084019093528383526001600160a01b0316908201819052865190919087908690811061132c5761132c6128e9565b60200260200101819052505050508060010190506111d3565b506040805160c0810182528b81525f60208201526113c691810161136e368b90038b018b612e39565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f9201829052509385525050604080519283526020808401909152909201525083611985565b60015f559a9950505050505050505050565b6002545f908061143a5760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b6064820152608401610497565b5f600354600161144a9190612e26565b9050611457828583611104565b949350505050565b5f818152600760205260408120541515610c05565b5f5f5f5f8451602086015f8989f195945050505050565b5f81515f0361149c57506001610c05565b81515f5b818110156114ed57846001600160a01b03168482815181106114c4576114c46128e9565b60200260200101516001600160a01b0316036114e557600192505050610c05565b6001016114a0565b505f949350505050565b805f036115565760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b6064820152608401610497565b5f9182526007602052604090912055565b5f6115768484845f0151611104565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906115cc575050505050565b5f8281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b0390811686526009855283862092515116855292528220805491929091611632908490612e26565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57846020015160405161167c9190612e53565b60405180910390a45050505050565b600182018054829160028501915f91826116a483612e73565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b8060a001515a10156117475760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b6064820152608401610497565b80608001516001600160a01b03163b5f0361176c57611767838383611c98565b505050565b60208101516101008201516040515f9261178a928692602401612ed8565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052608083015160a08401519192505f916117ce919084611474565b9050806117e7576117e0858585611c98565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b6040516001600160a01b0383811660248301526044820183905261176791859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611d0c565b805460018201545f91610c0591612f0a565b60605f6118a1600561189c85611879565b611d6d565b9050805f036118ed57604080515f80825260208201909252906118e5565b604080518082019091525f80825260208201528152602001906001900390816118bf5790505b509392505050565b5f816001600160401b03811115611906576119066128fd565b60405190808252806020026020018201604052801561194a57816020015b604080518082019091525f80825260208201528152602001906001900390816119245790505b5090505f5b828110156118e55761196085611d82565b828281518110611972576119726128e9565b602090810291909101015260010161194f565b5f5f60035f815461199590612e73565b91905081905590505f6119ae600254865f015184611104565b90505f604051806101200160405280848152602001336001600160a01b03168152602001306001600160a01b03168152602001875f0151815260200187602001516001600160a01b0316815260200187606001518152602001876080015181526020018681526020018760a0015181525090505f81604051602001611a33919061305e565b60405160208183030381529060405290505f5f8860400151602001511115611a9a576040880151516001600160a01b0316611a805760405162461bcd60e51b815260040161049790612dbe565b60408801518051602090910151611a979190611c8c565b90505b60408051808201825289820151516001600160a01b03908116825260208083018590528351808501855286518783012081528082018481525f8a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558951915190919086907f14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d190611b3d9088908690613070565b60405180910390a3604051633ae5f34b60e21b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb97cd2c90611b9190879060040161305e565b5f604051808303815f87803b158015611ba8575f5ffd5b505af1158015611bba573d5f5f3e3d5ffd5b50969b9a5050505050505050505050565b604080518082019091525f8082526020820152611be783611879565b8210611c3f5760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b6064820152608401610497565b826002015f83855f0154611c539190612e26565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b5f610ec6833384611e4c565b80604051602001611ca9919061305e565b60408051601f1981840301815282825280516020918201205f878152600690925291902055829084907fd1985b49432b2cff70033f5d03d555fdb512888d868400d49197a4a355c0245990611cff90859061305e565b60405180910390a3505050565b5f611d206001600160a01b03841683611faf565b905080515f14158015611d44575080806020019051810190611d429190612821565b155b1561176757604051635274afe760e01b81526001600160a01b0384166004820152602401610497565b5f818310611d7b5781610ec6565b5090919050565b604080518082019091525f808252602082015281546001830154819003611deb5760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d707479207175657565000000000000006044820152606401610497565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b0319909216905590611e42908390612e26565b9093555090919050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015611e92573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611eb69190613082565b9050611ecd6001600160a01b038616853086611fbc565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015611f11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f359190613082565b9050818111611f9b5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610497565b611fa58282612f0a565b9695505050505050565b6060610ec683835f611ffb565b6040516001600160a01b038481166024830152838116604483015260648201839052611ff59186918216906323b872dd90608401611847565b50505050565b6060814710156120205760405163cd78605960e01b8152306004820152602401610497565b5f5f856001600160a01b0316848660405161203b9190613099565b5f6040518083038185875af1925050503d805f8114612075576040519150601f19603f3d011682016040523d82523d5f602084013e61207a565b606091505b5091509150611fa586838360608261209a57612095826120e1565b610ec6565b81511580156120b157506001600160a01b0384163b155b156120da57604051639996b31560e01b81526001600160a01b0385166004820152602401610497565b5080610ec6565b8051156120f15780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f610120828403121561211b575f5ffd5b50919050565b5f5f60408385031215612132575f5ffd5b8235915060208301356001600160401b0381111561214e575f5ffd5b61215a8582860161210a565b9150509250929050565b80356001600160a01b038116811461217a575f5ffd5b919050565b5f5f60408385031215612190575f5ffd5b82356001600160401b038111156121a5575f5ffd5b8301606081860312156121b6575f5ffd5b91506121c460208401612164565b90509250929050565b5f602082840312156121dd575f5ffd5b610ec682612164565b5f602082840312156121f6575f5ffd5b5035919050565b82815260608101610ec6602083018480516001600160a01b03168252602090810151910152565b5f60208284031215612234575f5ffd5b81356001600160401b03811115612249575f5ffd5b6114578482850161210a565b5f60208284031215612265575f5ffd5b81356001600160401b0381111561227a575f5ffd5b820160e08185031215610ec6575f5ffd5b5f5f6040838503121561229c575f5ffd5b50508035926020909101359150565b815181526020808301516001600160a01b03169082015260408101610c05565b5f5f5f606084860312156122dd575f5ffd5b833592506122ed60208501612164565b929592945050506040919091013590565b5f5f5f60608486031215612310575f5ffd5b505081359360208301359350604090920135919050565b5f5f83601f840112612337575f5ffd5b5081356001600160401b0381111561234d575f5ffd5b6020830191508360208260051b8501011115612367575f5ffd5b9250929050565b5f5f5f5f5f5f86880360a0811215612384575f5ffd5b8735965060208801356001600160401b038111156123a0575f5ffd5b6123ac8a828b01612327565b9097509550506040603f19820112156123c3575f5ffd5b5060408701925060808701356001600160401b038111156123e2575f5ffd5b6123ee89828a01612327565b979a9699509497509295939492505050565b5f5f60408385031215612411575f5ffd5b6121b683612164565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f5f8335601e198436030181126124ba575f5ffd5b83016020810192503590506001600160401b038111156124d8575f5ffd5b8060051b3603821315612367575f5ffd5b8183526020830192505f815f5b84811015612525576001600160a01b0361250f83612164565b16865260209586019591909101906001016124f6565b5093949350505050565b5f5f8335601e19843603018112612544575f5ffd5b83016020810192503590506001600160401b03811115612562575f5ffd5b8060061b3603821315612367575f5ffd5b8183526020830192505f815f5b8481101561252557813586526001600160a01b036125a060208401612164565b1660208701526040958601959190910190600101612580565b5f5f8335601e198436030181126125ce575f5ffd5b83016020810192503590506001600160401b038111156125ec575f5ffd5b803603821315612367575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b803582525f61263360208301612164565b6001600160a01b0316602084015261264d60408301612164565b6001600160a01b031660408401526060828101359084015261267160808301612164565b6001600160a01b0316608084015260a0828101359084015261269660c08301836124a5565b61012060c08601526126ad610120860182846124e9565b9150506126bd60e084018461252f565b85830360e08701526126d0838284612573565b925050506126e26101008401846125b9565b858303610100870152611fa58382846125fa565b602081525f610ec66020830184612622565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b5f5f8335601e19843603018112612766575f5ffd5b8301803591506001600160401b0382111561277f575f5ffd5b602001915036819003821315612367575f5ffd5b8481526001600160a01b03841660208201526060604082018190525f90611fa590830184866125fa565b602081525f823561011e198436030181126127d6575f5ffd5b606060208401526127ec60808401858301612622565b90505f602085013590508060408501525061280a60408501856125b9565b848303601f19016060860152611fa58382846125fa565b5f60208284031215612831575f5ffd5b81518015158114610ec6575f5ffd5b5f823561011e19833603018112612855575f5ffd5b9190910192915050565b5f5f8335601e19843603018112612874575f5ffd5b8301803591506001600160401b0382111561288d575f5ffd5b6020019150600581901b3603821315612367575f5ffd5b5f5f8335601e198436030181126128b9575f5ffd5b8301803591506001600160401b038211156128d2575f5ffd5b6020019150600681901b3603821315612367575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715612933576129336128fd565b60405290565b60405161012081016001600160401b0381118282101715612933576129336128fd565b60405160c081016001600160401b0381118282101715612933576129336128fd565b604051601f8201601f191681016001600160401b03811182821017156129a6576129a66128fd565b604052919050565b5f604082840312156129be575f5ffd5b6129c6612911565b8235815290506129d860208301612164565b602082015292915050565b5f604082840312156129f3575f5ffd5b610ec683836129ae565b6001600160a01b03831681526040602082018190525f9061145790830184612622565b5f6001600160401b03821115612a3857612a386128fd565b5060051b60200190565b5f82601f830112612a51575f5ffd5b8135612a64612a5f82612a20565b61297e565b8082825260208201915060208360051b860101925085831115612a85575f5ffd5b602085015b83811015612aa957612a9b81612164565b835260209283019201612a8a565b5095945050505050565b5f82601f830112612ac2575f5ffd5b8135612ad0612a5f82612a20565b8082825260208201915060208360061b860101925085831115612af1575f5ffd5b602085015b83811015612aa957612b0887826129ae565b8352602090920191604001612af6565b5f82601f830112612b27575f5ffd5b81356001600160401b03811115612b4057612b406128fd565b612b53601f8201601f191660200161297e565b818152846020838601011115612b67575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6101208236031215612b94575f5ffd5b612b9c612939565b82358152612bac60208401612164565b6020820152612bbd60408401612164565b604082015260608381013590820152612bd860808401612164565b608082015260a0838101359082015260c08301356001600160401b03811115612bff575f5ffd5b612c0b36828601612a42565b60c08301525060e08301356001600160401b03811115612c29575f5ffd5b612c3536828601612ab3565b60e0830152506101008301356001600160401b03811115612c54575f5ffd5b612c6036828601612b18565b6101008301525092915050565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b606081525f612cc26060830185612622565b9050610ec6602083018480516001600160a01b03168252602090810151910152565b5f60408284031215612cf4575f5ffd5b612cfc612911565b9050612d0782612164565b815260209182013591810191909152919050565b5f60e08236031215612d2b575f5ffd5b612d3361295c565b82358152612d4360208401612164565b6020820152612d553660408501612ce4565b60408201526080830135606082015260a08301356001600160401b03811115612d7c575f5ffd5b612d8836828601612a42565b60808301525060c08301356001600160401b03811115612da6575f5ffd5b612db236828601612b18565b60a08301525092915050565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610c0557610c05612e12565b5f60408284031215612e49575f5ffd5b610ec68383612ce4565b81516001600160a01b031681526020808301519082015260408101610c05565b5f60018201612e8457612e84612e12565b5060010190565b5f5b83811015612ea5578181015183820152602001612e8d565b50505f910152565b5f8151808452612ec4816020860160208601612e8b565b601f01601f19169290920160200192915050565b8381526001600160a01b03831660208201526060604082018190525f90612f0190830184612ead565b95945050505050565b81810381811115610c0557610c05612e12565b5f8151808452602084019350602083015f5b828110156125255781516001600160a01b0316865260209586019590910190600101612f2f565b5f8151808452602084019350602083015f5b8281101561252557612f8e868351805182526020908101516001600160a01b0316910152565b6040959095019460209190910190600101612f68565b805182525f6020820151612fc360208501826001600160a01b03169052565b506040820151612fde60408501826001600160a01b03169052565b5060608201516060840152608082015161300360808501826001600160a01b03169052565b5060a082015160a084015260c082015161012060c0850152613029610120850182612f1d565b905060e083015184820360e08601526130428282612f56565b915050610100830151848203610100860152612f018282612ead565b602081525f610ec66020830184612fa4565b606081525f612cc26060830185612fa4565b5f60208284031215613092575f5ffd5b5051919050565b5f8251612855818460208701612e8b56fea264697066735822122088a94cba91914b9e9c6a7a567331b66525d67e64777ca2ba7a6be335c39a029264736f6c634300081e0033",
}

// TeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerMetaData.ABI instead.
var TeleporterMessengerABI = TeleporterMessengerMetaData.ABI

// TeleporterMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterMessengerMetaData.Bin instead.
var TeleporterMessengerBin = TeleporterMessengerMetaData.Bin

// DeployTeleporterMessenger deploys a new Ethereum contract, binding an instance of TeleporterMessenger to it.
func DeployTeleporterMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, verifierSender common.Address) (common.Address, *types.Transaction, *TeleporterMessenger, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterMessengerBin), backend, verifierSender)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// TeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type TeleporterMessenger struct {
	TeleporterMessengerCaller     // Read-only binding to the contract
	TeleporterMessengerTransactor // Write-only binding to the contract
	TeleporterMessengerFilterer   // Log filterer for contract events
}

// TeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterMessengerSession struct {
	Contract     *TeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterMessengerCallerSession struct {
	Contract *TeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterMessengerTransactorSession struct {
	Contract     *TeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterMessengerRaw struct {
	Contract *TeleporterMessenger // Generic contract binding to access the raw methods on
}

// TeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterMessengerCallerRaw struct {
	Contract *TeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactorRaw struct {
	Contract *TeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterMessenger creates a new instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessenger(address common.Address, backend bind.ContractBackend) (*TeleporterMessenger, error) {
	contract, err := bindTeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// NewTeleporterMessengerCaller creates a new read-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleporterMessengerCaller, error) {
	contract, err := bindTeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerCaller{contract: contract}, nil
}

// NewTeleporterMessengerTransactor creates a new write-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterMessengerTransactor, error) {
	contract, err := bindTeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerTransactor{contract: contract}, nil
}

// NewTeleporterMessengerFilterer creates a new log filterer instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterMessengerFilterer, error) {
	contract, err := bindTeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerFilterer{contract: contract}, nil
}

// bindTeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindTeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.TeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) CalculateMessageID(opts *bind.CallOpts, sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "calculateMessageID", sourceBlockchainID, destinationBlockchainID, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", messageID)

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
func (_TeleporterMessenger *TeleporterMessengerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageSender() (common.Address, error) {
	return _TeleporterMessenger.Contract.MessageSender(&_TeleporterMessenger.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageSender() (common.Address, error) {
	return _TeleporterMessenger.Contract.MessageSender(&_TeleporterMessenger.CallOpts)
}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageVerifier() (common.Address, error) {
	return _TeleporterMessenger.Contract.MessageVerifier(&_TeleporterMessenger.CallOpts)
}

// MessageVerifier is a free data retrieval call binding the contract method 0x8ae9753d.
//
// Solidity: function messageVerifier() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageVerifier() (common.Address, error) {
	return _TeleporterMessenger.Contract.MessageVerifier(&_TeleporterMessenger.CallOpts)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceiptQueues(opts *bind.CallOpts, sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receiptQueues", sourceBlockchainID)

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
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCaller) SentMessageInfo(opts *bind.CallOpts, messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "sentMessageInfo", messageID)

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
func (_TeleporterMessenger *TeleporterMessengerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x1c54d535.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, message ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", message, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x1c54d535.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiveCrossChainMessage(message ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, message, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x1c54d535.
//
// Solidity: function receiveCrossChainMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),bytes32,bytes) message, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) ReceiveCrossChainMessage(message ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, message, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0x0f635b6c.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x3ba2b983.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// TeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmountIterator struct {
	Event *TeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerAddFeeAmount)
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
		it.Event = new(TeleporterMessengerAddFeeAmount)
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
func (it *TeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*TeleporterMessengerAddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerAddFeeAmountIterator{contract: _TeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerAddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerAddFeeAmount)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleporterMessengerAddFeeAmount, error) {
	event := new(TeleporterMessengerAddFeeAmount)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerBlockchainIDInitializedIterator is returned from FilterBlockchainIDInitialized and is used to iterate over the raw logs and unpacked data for BlockchainIDInitialized events raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitializedIterator struct {
	Event *TeleporterMessengerBlockchainIDInitialized // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
		it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerBlockchainIDInitialized represents a BlockchainIDInitialized event raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitialized struct {
	BlockchainID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockchainIDInitialized is a free log retrieval operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterBlockchainIDInitialized(opts *bind.FilterOpts, blockchainID [][32]byte) (*TeleporterMessengerBlockchainIDInitializedIterator, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerBlockchainIDInitializedIterator{contract: _TeleporterMessenger.contract, event: "BlockchainIDInitialized", logs: logs, sub: sub}, nil
}

// WatchBlockchainIDInitialized is a free log subscription operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchBlockchainIDInitialized(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerBlockchainIDInitialized, blockchainID [][32]byte) (event.Subscription, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerBlockchainIDInitialized)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
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

// ParseBlockchainIDInitialized is a log parse operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseBlockchainIDInitialized(log types.Log) (*TeleporterMessengerBlockchainIDInitialized, error) {
	event := new(TeleporterMessengerBlockchainIDInitialized)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutedIterator struct {
	Event *TeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecuted)
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
		it.Event = new(TeleporterMessengerMessageExecuted)
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
func (it *TeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecuted)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*TeleporterMessengerMessageExecuted, error) {
	event := new(TeleporterMessengerMessageExecuted)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailedIterator struct {
	Event *TeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
		it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
func (it *TeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0xd1985b49432b2cff70033f5d03d555fdb512888d868400d49197a4a355c02459.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutionFailedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0xd1985b49432b2cff70033f5d03d555fdb512888d868400d49197a4a355c02459.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecutionFailed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0xd1985b49432b2cff70033f5d03d555fdb512888d868400d49197a4a355c02459.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*TeleporterMessengerMessageExecutionFailed, error) {
	event := new(TeleporterMessengerMessageExecutionFailed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceivedIterator struct {
	Event *TeleporterMessengerReceiptReceived // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiptReceived)
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
		it.Event = new(TeleporterMessengerReceiptReceived)
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
func (it *TeleporterMessengerReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiptReceived represents a ReceiptReceived event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*TeleporterMessengerReceiptReceivedIterator, error) {

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

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiptReceivedIterator{contract: _TeleporterMessenger.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiptReceived)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiptReceived(log types.Log) (*TeleporterMessengerReceiptReceived, error) {
	event := new(TeleporterMessengerReceiptReceived)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *TeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
		it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x51f9dcb4f762bab1ca8136939fec0e5a257efbd0f69dea14b60d6d1ae0750f80.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*TeleporterMessengerReceiveCrossChainMessageIterator, error) {

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

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiveCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x51f9dcb4f762bab1ca8136939fec0e5a257efbd0f69dea14b60d6d1ae0750f80.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiveCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x51f9dcb4f762bab1ca8136939fec0e5a257efbd0f69dea14b60d6d1ae0750f80.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(TeleporterMessengerReceiveCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *TeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
		it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*TeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerRelayerRewardsRedeemedIterator{contract: _TeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerRelayerRewardsRedeemed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*TeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(TeleporterMessengerRelayerRewardsRedeemed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessageIterator struct {
	Event *TeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
		it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
func (it *TeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*TeleporterMessengerSendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerSendCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x14d9f96bda232881a6460c5fca942a64c8f5babfc75ae626b27caff9fa7ce7d1.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerSendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerSendCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleporterMessengerSendCrossChainMessage, error) {
	event := new(TeleporterMessengerSendCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
