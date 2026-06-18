// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenremote

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

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	RecipientContract                  common.Address
	RecipientPayload                   []byte
	RequiredGasLimit                   *big.Int
	RecipientGasLimit                  *big.Int
	MultiHopFallback                   common.Address
	FallbackRecipient                  common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenRemoteSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenRemoteSettings struct {
	TeleporterRegistryAddress common.Address
	TeleporterManager         common.Address
	MinTeleporterVersion      *big.Int
	TokenHomeBlockchainID     [32]byte
	TokenHomeAddress          common.Address
	TokenHomeDecimals         uint8
}

// NativeTokenRemoteMetaData contains all meta data concerning the NativeTokenRemote contract.
var NativeTokenRemoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHomeAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenRemoteSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_TRANSFER_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOME_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMultiplyOnRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHomeAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenRemoteSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461090057615c55803803809161001c8285610df9565b833981019080820361012081126109005760c013610900576040519060c082016001600160401b03811183821017610c1e5760405261005a81610e1c565b825261006860208201610e1c565b906020830191825260408101516040840190815260608201516060850190815261009460808401610e1c565b916080860192835260a08401519460ff861686036109005760a0870195865260c08501516001600160401b03811161090057850188601f82011215610900578051906001600160401b038211610c1e57604051996100fc601f8401601f19166020018c610df9565b828b5260208383010111610900575f5b828110610de557505060205f918a01015261010060e0860151950151965f516020615c155f395f51905f52549860ff8a60401c16159960018060401b03811680159081610ddd575b6001149081610dd3575b159081610dca575b50610dbb576001600160401b031981166001175f516020615c155f395f51905f52558a610d91575b50610197610e30565b8615610d32576101a5610e30565b6101ad610e30565b80516001600160401b038111610c1e575f516020615b955f395f51905f5254600181811c91168015610d28575b6020821014610c0057601f8111610cb9575b50806020601f8211600114610c3d575f91610c32575b508160011b915f199060031b1c1916175f516020615b955f395f51905f52555b8051906001600160401b038211610c1e575f516020615bb55f395f51905f525490600182811c92168015610c14575b6020831014610c005781601f849311610b86575b50602090601f8311600114610b08575f92610afd575b50508160011b915f199060031b1c1916175f516020615bb55f395f51905f52555b6102a4610e30565b51905191516001600160a01b03928316929091166102c0610e30565b6102c8610e30565b6102d0610e30565b6102d8610e30565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055610304610e30565b61030c610e30565b8015610a925760405163301fd1f560e21b8152602081600481855afa90811561090c575f91610a60575b5015610a13577fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0080546001600160a01b0319168217905560405163301fd1f560e21b815290602090829060049082905afa90811561090c575f916109e1575b505f516020615bd55f395f51905f5254908211610995578082111561092a57815f516020615bd55f395f51905f52557fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d5f80a36103f0610e30565b6103f8610e30565b8015610917577f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3610466610e30565b60017fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750055610492610e30565b60405163084279ef60e31b81526020816004817302000000000000000000000000000000000000055afa90811561090c575f916108d6575b50807f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef005581511561087e578151146108135781516001600160a01b0316156107c257601260ff8551161161076b57517f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0155517f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0280547f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef05939093557f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef06805460ff1916905592516001600160b01b03199092166001600160a01b03919091161760a082901b60ff60a01b1617600960a91b1790915560ff16601281109081156107565760120360ff81116107425760ff16905b604d82116107425760ff80195f516020615bf55f395f51905f5254169116175f516020615bf55f395f51905f5255600a0a7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0355610656610e30565b60648110156106ef577f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50055610695575b604051614d399081610e5c8239f35b68ff0000000000000000195f516020615c155f395f51905f5254165f516020615c155f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1610686565b60405162461bcd60e51b815260206004820152602560248201527f4e6174697665546f6b656e52656d6f74653a20696e76616c69642070657263656044820152646e7461676560d81b6064820152608490fd5b634e487b7160e01b5f52601160045260245ffd5b6011190160ff81116107425760ff16906105fb565b60405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b6064820152608490fd5b60405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d6500000000006064820152608490fd5b60405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b6064820152608490fd5b90506020813d602011610904575b816108f160209383610df9565b8101031261090057515f6104ca565b5f80fd5b3d91506108e4565b6040513d5f823e3d90fd5b631e4fbdf760e01b5f525f60045260245ffd5b60405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608490fd5b60405162461bcd60e51b815260206004820152603160248201525f516020615c355f395f51905f5260448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608490fd5b90506020813d602011610a0b575b816109fc60209383610df9565b8101031261090057515f610395565b3d91506109ef565b60405162461bcd60e51b815260206004820152603260248201525f516020615c355f395f51905f52604482015271656c65706f7274657220726567697374727960701b6064820152608490fd5b90506020813d602011610a8a575b81610a7b60209383610df9565b8101031261090057515f610336565b3d9150610a6e565b60405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608490fd5b015190505f8061027b565b5f516020615bb55f395f51905f525f9081528281209350601f198516905b818110610b6e5750908460019594939210610b56575b505050811b015f516020615bb55f395f51905f525561029c565b01515f1960f88460031b161c191690555f8080610b3c565b92936020600181928786015181550195019301610b26565b5f516020615bb55f395f51905f525f529091507f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa601f840160051c81019160208510610bf6575b90601f859493920160051c01905b818110610be85750610265565b5f8155849350600101610bdb565b9091508190610bcd565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610251565b634e487b7160e01b5f52604160045260245ffd5b90508201515f610202565b5f516020615b955f395f51905f525f9081528181209250601f198416905b818110610ca157509083600194939210610c89575b5050811b015f516020615b955f395f51905f5255610222565b8401515f1960f88460031b161c191690555f80610c70565b91926020600181928689015181550194019201610c5b565b5f516020615b955f395f51905f525f527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0601f830160051c81019160208410610d1e575b601f0160051c01905b818110610d1357506101ec565b5f8155600101610d06565b9091508190610cfd565b90607f16906101da565b60405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e52656d6f74653a207a65726f20696e697469616c206044820152707265736572766520696d62616c616e636560781b6064820152608490fd5b6001600160481b03191668010000000000000001175f516020615c155f395f51905f52555f61018e565b63f92ee8a960e01b5f5260045ffd5b9050155f610166565b303b15915061015e565b8c9150610154565b60208183018101518c83018201520161010c565b601f909101601f19168101906001600160401b03821190821017610c1e57604052565b51906001600160a01b038216820361090057565b60ff5f516020615c155f395f51905f525460401c1615610e4c57565b631afcd79f60e31b5f5260045ffdfe60806040526004361015610026575b361561001e5761001c611e73565b005b61001c611e73565b5f3560e01c806302a30c7d1461030057806306fdde03146102fb5780630733c8c8146102f6578063095ea7b3146102f15780630ca1c5c9146102ec57806315beb59f146102e757806318160ddd146102e25780631906529c146102dd57806323b872dd146102d8578063254ac160146102d35780632b0d8f18146102ce5780632e1a7d4d146102c9578063313ce567146102c4578063329c3e12146102bf578063347212c41461023357806335cac159146102ba5780634213cf78146102b55780634511243e146102b05780635507f3d1146102ab57806355538c8b146102a65780635eb99514146102a15780636e6eef8d1461029c57806370a0823114610297578063715018a61461029257806371717c181461028d5780637ee3779a146102885780638bf2fa94146102835780638da5cb5b1461027e5780638f6cec8814610279578063909a6ac01461027457806395d89b411461026f578063973142971461026a578063a5717bc014610265578063a9059cbb14610260578063b8a46d021461025b578063c3cd692714610256578063c452165e14610251578063c868efaa1461024c578063d0e30db014610247578063d2cc7a7014610242578063dd62ed3e1461023d578063e0fd9cb814610238578063ed0ae4b014610233578063ef793e2a1461022e578063f2fde38b146102295763f3f981d80361000e576116db565b6116ae565b611685565b6108d3565b61165c565b61161c565b6115f3565b6115e0565b611576565b611557565b611523565b611362565b611338565b6112fe565b6112c9565b61120c565b6111e5565b61113c565b610fd0565b610e90565b610e62565b610e45565b610dde565b610d84565b610cf2565b610cce565b610a43565b610a26565b610946565b61091d565b6108f6565b6108b1565b610896565b6107a4565b6106be565b6106a1565b610656565b6105eb565b6105c2565b6105a6565b61057d565b6104c5565b61047e565b61038c565b3461032e575f36600319011261032e57602060ff5f516020614b6d5f395f51905f5254166040519015158152f35b5f80fd5b5f5b8381106103435750505f910152565b8181015183820152602001610334565b9060209161036c81518092818552858086019101610332565b601f01601f1916010190565b906020610389928181520190610353565b90565b3461032e575f36600319011261032e576040515f5f516020614b2d5f395f51905f52546103b881611701565b808452906001811690811561045a57506001146103f0575b6103ec836103e081850382611053565b60405191829182610378565b0390f35b5f516020614b2d5f395f51905f525f9081527f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0939250905b808210610440575090915081016020016103e06103d0565b919260018160209254838588010152019101909291610428565b60ff191660208086019190915291151560051b840190910191506103e090506103d0565b3461032e575f36600319011261032e5760205f516020614bcd5f395f51905f5254604051908152f35b6001600160a01b0381160361032e57565b35906104c3826104a7565b565b3461032e57604036600319011261032e576004356104e2816104a7565b602435331561056a576001600160a01b0382169182156105575761051f829161050a33611807565b9060018060a01b03165f5260205260405f2090565b556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b3461032e575f36600319011261032e5760205f516020614b4d5f395f51905f5254604051908152f35b3461032e575f36600319011261032e5760206040516105dc8152f35b3461032e575f36600319011261032e5760205f516020614bed5f395f51905f5254604051908152f35b3461032e575f36600319011261032e57600160981b3162010203600160981b01318101809111610651575f516020614b4d5f395f51905f5254905f516020614ccd5f395f51905f5254820180921161065157810390811161065157602090604051908152f35b611739565b3461032e57606036600319011261032e57610696600435610676816104a7565b602435610682816104a7565b60443591610691833383611f2b565b611fa9565b602060405160018152f35b3461032e575f36600319011261032e5760206040516201fbd08152f35b3461032e57602036600319011261032e576004356106db816104a7565b6106e36125df565b6001600160a01b038116906106f982151561176c565b60ff610704826117cf565b541661074957610716610723916117cf565b805460ff19166001179055565b7f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c5f80a2005b60405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608490fd5b3461032e57602036600319011261032e576004356040518181527f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b6560203392a23315610883576107f33361183f565b5490808210610867578061001c920361080b3361183f565b55610831815f516020614bed5f395f51905f5254035f516020614bed5f395f51905f5255565b6040518181525f9033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602090a333612060565b63391434e360e21b5f523360045260249190915260445260645ffd5b634b637e8f60e11b5f525f60045260245ffd5b3461032e575f36600319011261032e57602060405160128152f35b3461032e575f36600319011261032e576040516001600160991b018152602090f35b3461032e575f36600319011261032e57602060405162010203600160981b018152f35b3461032e575f36600319011261032e5760206040515f516020614c6d5f395f51905f528152f35b3461032e575f36600319011261032e5760205f516020614c6d5f395f51905f5254604051908152f35b3461032e57602036600319011261032e57600435610963816104a7565b61096b6125df565b6001600160a01b0381169061098182151561176c565b60ff61098c826117cf565b5416156109cf5761099f6109a9916117cf565b805460ff19169055565b7f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c35f80a2005b60405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608490fd5b3461032e575f36600319011261032e576020604051620530208152f35b3461032e57602036600319011261032e57600435610a7160015f516020614c0d5f395f51905f525414611877565b60025f516020614c0d5f395f51905f52557f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5610c9c610c8b600160981b3193610c49610bcf610aec7f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50254610ae6818a116118cf565b8861175f565b92610b54610b2f610b28610b217f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf5005488611965565b6064900490565b809661175f565b987f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50255565b83610cb5575b610b95610b8e89610b765f516020614bcd5f395f51905f525490565b5f516020614c4d5f395f51905f525460ff1690612cd8565b1515611978565b610bdd610ba0611074565b62010203600160981b0180825260209182018b8152604080519384019290925251908201529283906060820190565b03601f198101845283611053565b610be5611074565b600181529160208301525f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f5254909490610c72906001600160a01b0316610c25611074565b308152926020840152610c57610c396119ff565b9560405197889160208301611a1a565b03601f198101885287611053565b610c5f611083565b9687526001600160a01b03166020870152565b60408501526060840152608083015260a08201526122de565b604051938452929081906020820190565b0390a261001c60015f516020614c0d5f395f51905f5255565b610cbf84306120d4565b610cc98430612157565b610b5a565b3461032e57602036600319011261032e5761001c600435610ced6125df565b612412565b602036600319011261032e576004356001600160401b03811161032e57610160600319823603011261032e57610d7190610d3c60ff5f516020614b6d5f395f51905f525416611a47565b610d5660015f516020614c0d5f395f51905f525414611877565b60025f516020614c0d5f395f51905f5255349060040161252d565b60015f516020614c0d5f395f51905f5255005b3461032e57602036600319011261032e57600435610da1816104a7565b60018060a01b03165f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00602052602060405f2054604051908152f35b3461032e575f36600319011261032e57610df66125df565b5f516020614bad5f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461032e575f36600319011261032e576020604051620557308152f35b3461032e575f36600319011261032e57602060ff5f516020614c4d5f395f51905f5254166040519015158152f35b61010036600319011261032e57610ebd610eb860ff5f516020614b6d5f395f51905f52541690565b611a47565b610ed760015f516020614c0d5f395f51905f525414611877565b60025f516020614c0d5f395f51905f5255604435610ef4816104a7565b6001600160a01b031615610f7f57610f0f60c4351515613075565b600435610f1d811515612f42565b610f3e610f37610f2b611c4e565b6001600160a01b031690565b1515612fa2565b5f516020614b0d5f395f51905f525403610f7157610f5b34613b04565b61001c60015f516020614c0d5f395f51905f5255565b610f7a3461393a565b610f5b565b60405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b6064820152608490fd5b3461032e575f36600319011261032e575f516020614bad5f395f51905f52546040516001600160a01b039091168152602090f35b634e487b7160e01b5f52604160045260245ffd5b60c081019081106001600160401b0382111761103357604052565b611004565b604081019081106001600160401b0382111761103357604052565b90601f801991011681019081106001600160401b0382111761103357604052565b604051906104c3604083611053565b604051906104c360c083611053565b604051906104c3606083611053565b604051906104c361010083611053565b604051906104c361016083611053565b604051906104c360e083611053565b6001600160401b03811161103357601f01601f191660200190565b9291926110f7826110d0565b916111056040519384611053565b82948184528183011161032e578281602093845f960137010152565b9080601f8301121561032e57816020610389933591016110eb565b3461032e573660031901610120811261032e5760c01361032e5760405161116281611018565b60043561116e816104a7565b815260243561117c816104a7565b60208201526044356040820152606435606082015260843561119d816104a7565b608082015260a43560ff8116810361032e5760a082015260c435906001600160401b03821161032e576111d761001c923690600401611121565b60e435906101043592611aab565b3461032e575f36600319011261032e5760206040515f516020614ced5f395f51905f528152f35b3461032e575f36600319011261032e576040515f5f516020614b8d5f395f51905f525461123881611701565b808452906001811690811561045a575060011461125f576103ec836103e081850382611053565b5f516020614b8d5f395f51905f525f9081527f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa939250905b8082106112af575090915081016020016103e06103d0565b919260018160209254838588010152019101909291611297565b3461032e57602036600319011261032e57602060ff6112f26004356112ed816104a7565b6117cf565b54166040519015158152f35b3461032e575f36600319011261032e5760206040517f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf5008152f35b3461032e57604036600319011261032e57610696600435611358816104a7565b6024359033611fa9565b3461032e57604036600319011261032e5761139e6113996113955f516020614b6d5f395f51905f525460ff9060081c1690565b1590565b611c02565b61001c6114356114435f516020614ccd5f395f51905f52546114065f516020614c8d5f395f51905f52546113fc6113e66113dc8360ff9060a01c1690565b9260a81c60ff1690565b916113ef611092565b94855260ff166020850152565b60ff166040830152565b60405192839160208301919091604060ff81606084019580518552826020820151166020860152015116910152565b03601f198101835282611053565b61144b611074565b905f825260208201526114dd61146b611462611c5a565b602435906127c8565b5f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f525490939190611506906001600160a01b03166114a2611c5a565b926114bd6114ae611074565b6001600160a01b039095168552565b60208401526114eb6114cd6119ff565b9460405196879160208301611a1a565b03601f198101875286611053565b6114f3611083565b9586526001600160a01b03166020860152565b60408401526201fbd06060840152608083015260a08201526122de565b3461032e575f36600319011261032e575f516020614c8d5f395f51905f52546040516001600160a01b039091168152602090f35b3461032e575f36600319011261032e57604051600160981b8152602090f35b3461032e57606036600319011261032e57600435602435611596816104a7565b604435916001600160401b03831161032e573660238401121561032e578260040135916001600160401b03831161032e57366024848601011161032e57602461001c940191611c70565b5f36600319011261032e5761001c611e73565b3461032e575f36600319011261032e5760205f516020614c2d5f395f51905f5254604051908152f35b3461032e57604036600319011261032e57602061165360043561163e816104a7565b61050a6024359161164e836104a7565b611807565b54604051908152f35b3461032e575f36600319011261032e5760205f516020614b0d5f395f51905f5254604051908152f35b3461032e575f36600319011261032e5760205f516020614ccd5f395f51905f5254604051908152f35b3461032e57602036600319011261032e5761001c6004356116ce816104a7565b6116d66125df565b611ea9565b3461032e57602036600319011261032e5760206116f9600435611f1a565b604051908152f35b90600182811c9216801561172f575b602083101461171b57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611710565b634e487b7160e01b5f52601160045260245ffd5b62055730019081620557301161065157565b9190820391821161065157565b1561177357565b60405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201526d706f72746572206164647265737360901b6064820152608490fd5b6001600160a01b03165f9081527fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace006020526040902090565b1561187e57565b60405162461bcd60e51b8152602060048201526024808201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604482015263616e637960e01b6064820152608490fd5b156118d657565b60a460405162461bcd60e51b815260206004820152604460248201527f4e6174697665546f6b656e52656d6f74653a206275726e20616464726573732060448201527f62616c616e6365206e6f742067726561746572207468616e206c6173742072656064820152631c1bdc9d60e21b6084820152fd5b906105dc8202918083046105dc149015171561065157565b8181029291811591840414171561065157565b1561197f57565b60405162461bcd60e51b815260206004820152603460248201527f4e6174697665546f6b656e52656d6f74653a207a65726f207363616c6564206160448201527336b7bab73a103a37903932b837b93a10313ab93760611b6064820152608490fd5b600511156119eb57565b634e487b7160e01b5f52602160045260245ffd5b60405190611a0e602083611053565b5f808352366020840137565b6020815281519160058310156119eb57602060609161038994828501520151916040808201520190610353565b15611a4e57565b60405162461bcd60e51b815260206004820152602f60248201527f4e6174697665546f6b656e52656d6f74653a20636f6e747261637420756e646560448201526e1c98dbdb1b185d195c985b1a5e9959608a1b6064820152608490fd5b5f516020614d0d5f395f51905f525493909290916001600160401b03611ae0604087901c60ff1615966001600160401b031690565b1680159081611bfa575b6001149081611bf0575b159081611be7575b50611bd857611b3f9385611b3660016001600160401b03195f516020614d0d5f395f51905f525416175f516020614d0d5f395f51905f5255565b611b9e57612612565b611b4557565b611b6f60ff60401b195f516020614d0d5f395f51905f5254165f516020614d0d5f395f51905f5255565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1565b611bd36801000000000000000060ff60401b195f516020614d0d5f395f51905f525416175f516020614d0d5f395f51905f5255565b612612565b63f92ee8a960e01b5f5260045ffd5b9050155f611afc565b303b159150611af4565b869150611aea565b15611c0957565b60405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c72656164792072656769737465726564006044820152606490fd5b602435610389816104a7565b600435610389816104a7565b35610389816104a7565b919260025f516020614cad5f395f51905f525414611d805760025f516020614cad5f395f51905f52555f516020614ced5f395f51905f5254611cba906001600160a01b0316610f2b565b60405163260f846760e11b81523360048201529490602090869060249082905afa918215611d7b57611d0b611d3093611d36975f91611d4c575b505f516020614c2d5f395f51905f52541115611da9565b611d29611d24611395611d1d336117cf565b5460ff1690565b611e0e565b36916110eb565b91612b0e565b6104c360015f516020614cad5f395f51905f5255565b611d6e915060203d602011611d74575b611d668183611053565b810190611d8f565b5f611cf4565b503d611d5c565b611d9e565b633ee5aeb560e01b5f5260045ffd5b9081602091031261032e575190565b6040513d5f823e3d90fd5b15611db057565b60405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608490fd5b15611e1557565b60405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608490fd5b6040513481527fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c60203392a26104c33433612157565b6001600160a01b03168015611f07575f516020614bad5f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b601f81018091116106515760051c90565b9190611f3a8161050a85611807565b5460018101611f4a575b50505050565b828110611f88576001600160a01b0384161561056a576001600160a01b0382161561055757611f7e9261050a910393611807565b555f808080611f44565b90637dc7a0d960e11b5f5260018060a01b031660045260245260445260645ffd5b916001600160a01b038316918215610883576001600160a01b03811693841561204d57611fd58161183f565b54838110612028579161201691612010857fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9695039161183f565b5561183f565b805482019055604051908152602090a3565b63391434e360e21b5f526001600160a01b03909116600452602452604482905260645ffd5b63ec442f0560e01b5f525f60045260245ffd5b8147106120bd575f918291829182916001600160a01b03165af13d156120b8573d61208a816110d0565b906120986040519283611053565b81525f60203d92013e5b156120a957565b63d6bda27560e01b5f5260045ffd5b6120a2565b504763cf47918160e01b5f5260045260245260445ffd5b5f516020614b4d5f395f51905f5254828101809111610651575f516020614b4d5f395f51905f52556001600160991b013b1561032e576040516327ad555d60e11b81526001600160a01b0391909116600482015260248101919091525f81604481836001600160991b015af18015611d7b5761214d5750565b5f6104c391611053565b6001600160a01b0381169190821561204d575f516020614bed5f395f51905f525490828201809211610651575f926121c16020927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef945f516020614bed5f395f51905f525561183f565b818154019055604051908152a3565b156121d757565b60405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608490fd5b6020808252825181830152808301516001600160a01b039081166040808501919091528401518051909116606084015201516080820152909190610100810190606084015160a082015260808401519160e060c08301528251809152602061012083019301905f5b8181106122bf5750505060a0610389939401519060e0601f1982850301910152610353565b82516001600160a01b031685526020948501949092019160010161229a565b602061231c5f926122ed612d6c565b906040810180519085820151612353575b5050604051948580948193630624488560e41b835260048301612232565b03926001600160a01b03165af1908115611d7b575f9161233a575090565b610389915060203d602011611d7457611d668183611053565b9051612399919061236e906001600160a01b031615156121d0565b5180518690612385906001600160a01b0316610f2b565b910151906001600160a01b03851690612e07565b5f806122fe565b156123a757565b60405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608490fd5b5f516020614ced5f395f51905f525460405163301fd1f560e21b815290602090829060049082906001600160a01b03165afa908115611d7b575f9161250e575b505f516020614c2d5f395f51905f52549082116124af576124748183116123a0565b612489825f516020614c2d5f395f51905f5255565b7fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d5f80a3565b60405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608490fd5b612527915060203d602011611d7457611d668183611053565b5f612452565b90813561253b811515612f42565b61255c602084013561254c816104a7565b6001600160a01b03161515612fa2565b6125776125706040850135610f2b816104a7565b1515613014565b6125a0608084013561258a811515613075565b60a085013561259a8115156130cd565b10613127565b6125b96125b2610f2b60e08601611c66565b1515613184565b5f516020614b0d5f395f51905f5254036125d6576104c3916136f5565b6104c391613421565b5f516020614bad5f395f51905f52546001600160a01b031633036125ff57565b63118cdaa760e01b5f523360045260245ffd5b9061261b613c02565b821561276957612629613c02565b612631613c02565b80516001600160401b038111611033576126618161265c5f516020614b2d5f395f51905f5254611701565b614310565b6020601f82116001146126c4576104c395949261269b836126b496946126af945f916126b9575b508160011b915f199060031b1c19161790565b5f516020614b2d5f395f51905f52556143bb565b613c2d565b613db2565b90508301515f612688565b5f516020614b2d5f395f51905f525f52601f198216907f2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0915f5b8181106127515750836126af936104c3999896936126b4989660019410612739575b5050811b015f516020614b2d5f395f51905f52556143bb565b8401515f1960f88460031b161c191690555f80612720565b919260206001819286890151815501940192016126fe565b60405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e52656d6f74653a207a65726f20696e697469616c206044820152707265736572766520696d62616c616e636560781b6064820152608490fd5b81156128b5576001600160a01b03169030821461289d576040516370a0823160e01b815230600482015290602082602481865afa918215611d7b575f92612878575b50612817903033856148ea565b6040516370a0823160e01b815230600482015291602090839060249082905afa8015611d7b57610389925f91612859575b50612854828211613e3a565b61175f565b612872915060203d602011611d7457611d668183611053565b5f612848565b6128179192506128969060203d602011611d7457611d668183611053565b919061280a565b90506128aa813033611f2b565b610389813033611fa9565b50505f90565b156128c257565b60405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608490fd5b1561292057565b60405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608490fd5b81601f8201121561032e57805161298e816110d0565b9261299c6040519485611053565b8184526020828401011161032e576103899160208085019101610332565b60208183031261032e578051906001600160401b03821161032e570160408183031261032e57604051916129ed83611038565b8151600581101561032e57835260208201516001600160401b03811161032e57612a179201612978565b602082015290565b51906104c3826104a7565b60208183031261032e578051906001600160401b03821161032e57016101008183031261032e57612a596110a1565b9181518352612a6a60208301612a1f565b6020840152612a7b60408301612a1f565b6040840152612a8c60608301612a1f565b60608401526080820151608084015260a0820151916001600160401b03831161032e57612ac060e092612ad5948301612978565b60a085015260c081015160c085015201612a1f565b60e082015290565b9081604091031261032e57602060405191612af783611038565b8051612b02816104a7565b83520151602082015290565b612b709291612b2f612b61925f516020614b0d5f395f51905f5254146128bb565b5f516020614c8d5f395f51905f5254612b50906001600160a01b0316610f2b565b6001600160a01b0390911614612919565b602080825183010191016129ba565b5f516020614b6d5f395f51905f525460ff600882901c1615908115612ccc575b50612c71575b60018151612ba3816119e1565b612bac816119e1565b03612be257612bcb60206104c392015160208082518301019101612add565b80516020906001600160a01b031691015190613fe1565b60028151612bef816119e1565b612bf8816119e1565b03612c2257612c1760206104c392015160208082518301019101612a2a565b608081015190613eca565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b6064820152608490fd5b612c9d61010061ff00195f516020614b6d5f395f51905f525416175f516020614b6d5f395f51905f5255565b612cc7600160ff195f516020614b6d5f395f51905f525416175f516020614b6d5f395f51905f5255565b612b96565b60ff161590505f612b90565b919015612cfd578115612ce9570490565b634e487b7160e01b5f52601260045260245ffd5b9061038991611965565b15612d0e57565b60405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608490fd5b5f516020614ced5f395f51905f525460405163d820e64f60e01b815290602090829060049082906001600160a01b03165afa908115611d7b575f91612dcc575b5061038960ff612dc46001600160a01b0384166117cf565b541615612d07565b90506020813d602011612dff575b81612de760209383611053565b8101031261032e5751612df9816104a7565b5f612dac565b3d9150612dda565b604051636eb1769f60e11b81523060048201526001600160a01b038084166024830152929390926020908490604490829086165afa928315611d7b575f93612f21575b5082018092116106515760405163095ea7b360e01b60208281019182526001600160a01b038616602484015260448301949094529092905f90612e9085606481016114dd565b84519082855af15f51903d81612ef5575b501590505b612eaf57505050565b60405163095ea7b360e01b60208201526001600160a01b0390931660248401525f60448401526104c392612ef090612eea8160648101611435565b826149e7565b6149e7565b15159050612f155750612ea66001600160a01b0382163b15155b5f612ea1565b6001612ea69114612f0f565b612f3b91935060203d602011611d7457611d668183611053565b915f612e4a565b15612f4957565b60405162461bcd60e51b815260206004820152602b60248201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60448201526a1bd8dad8da185a5b88125160aa1b6064820152608490fd5b15612fa957565b60405162461bcd60e51b815260206004820152603760248201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f60448201527f6b656e207472616e7366657272657220616464726573730000000000000000006064820152608490fd5b1561301b57565b60405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b6064820152608490fd5b1561307c57565b60405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6044820152631a5b5a5d60e21b6064820152608490fd5b156130d457565b60405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b6064820152608490fd5b1561312e57565b60405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b6064820152608490fd5b1561318b57565b60405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b6064820152608490fd5b903590601e198136030182121561032e57018035906001600160401b03821161032e5760200191813603831361032e57565b602080825282516001600160a01b03169082015260208201516040828101919091528201516001600160a01b0316606082015260608201516001600160a01b03166080820152608082015160a082015261016061014061328660a08501518360c0860152610180850190610353565b60c085015160e0858101919091528501516001600160a01b031661010085015293610100810151610120858101919091528101516001600160a01b031682850152015191015290565b9035601e198236030181121561032e5701602081359101916001600160401b03821161032e57813603831361032e57565b908060209392818452848401375f828201840152601f01601f1916010190565b929190602090604085528035604086015261334f61333f8383016104b8565b6001600160a01b03166060870152565b61336e61335e604083016104b8565b6001600160a01b03166080870152565b61014061339461338160608401846132cf565b61016060a08a01526101a0890191613300565b91608081013560c088015260a081013560e08801526133c96133b860c083016104b8565b6001600160a01b0316610100890152565b6133e96133d860e083016104b8565b6001600160a01b0316610120890152565b6134086133f961010083016104b8565b6001600160a01b031688840152565b6101208101356101608801520135610180860152930152565b6020810191813561343184611c66565b60c084019061343f82611c66565b6134499184614095565b61010084019061345882611c66565b6101208601359461014087013595613471928792614127565b94909661347d90611c66565b61348960408801611c66565b906060880193613499858a6131e5565b60a08b013591906134ac60e08d01611c66565b936134b690611c66565b9560808d0135956134c56110b1565b3381529b60208d01526001600160a01b031660408c01526001600160a01b031660608b01528c60808b015236906134fb926110eb565b60a089015260c08801526001600160a01b031660e08701526101008601526001600160a01b03166101208501526101408401526040518093602082019061354191613217565b03601f19810184526135539084611053565b61355b611074565b6004815292602084015261356f90856131e5565b6135799150611f1a565b6135829061194d565b61358b9061174d565b5f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f52549094906001600160a01b0316926135bf90611c66565b906135c8611074565b6001600160a01b03909216825260208201526135e26119ff565b92604051809560208201906135f691611a1a565b03601f19810186526136089086611053565b613610611083565b9586526001600160a01b0316602086015260408501526060840152608083015260a082015261363e906122de565b90604051809133946136509183613320565b037f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1691a3565b60208082528251818301528201516001600160a01b03908116604080840191909152830151811660608084019190915283015116608082015261038990608083015160a082015261010060e06136db60a08601518360c0860152610120850190610353565b60c0860151848301529401516001600160a01b0316910152565b907f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1661389c61388d610bcf610c4961376a61373260208901611c66565b61374f6101408a0135918261374960c08d01611c66565b91614230565b61010089019761375e89611c66565b6101208b013591614127565b9290966138135f516020614c6d5f395f51905f525461380461378e60408d01611c66565b8c6137ea8d6137a060608401846131e5565b90916137de6137b660e060a08801359701611c66565b966137bf6110a1565b998a523060208b01523360408b01526001600160a01b031660608a0152565b608088015236916110eb565b60a085015260c08401526001600160a01b031660e0830152565b60405193849160208301613676565b61381b611074565b600281529160208301525f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f5254909490613859906001600160a01b031692611c66565b90613874613865611074565b6001600160a01b039093168352565b6020820152610c7260808b013592610c57610c396119ff565b92604051918291339683613320565b0390a3565b610100909392919361393661392660e06101208401978035855260208101356138c9816104a7565b6001600160a01b0316602086015260408101356138e5816104a7565b6001600160a01b0316604086015261390261333f606083016104b8565b6080810135608086015260a081013560a086015260c081013560c0860152016104b8565b6001600160a01b031660e0830152565b0152565b613ad2613a006114dd611435613a6b60043561396a6139596024611c66565b61396360e4611c66565b9083614095565b6139866139776064611c66565b9760843560a435998a92614127565b9590976139936024611c66565b9061399e6044611c66565b906139df60c435926139cf6139b360e4611c66565b956139bc6110c1565b9889526001600160a01b03166020890152565b6001600160a01b03166040870152565b606085018b9052608085015260a08401526001600160a01b031660c0830152565b6040519283916020830191909160c060e08201938051835260018060a01b03602082015116602084015260018060a01b036040820151166040840152606081015160608401526080810151608084015260a081015160a08401528160018060a01b0391015116910152565b613a73611074565b600381529060208201525f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f5254909390613ab5906001600160a01b03166114a26064611c66565b6040840152620530206060840152608083015260a08201526122de565b7f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb526040518061389c33956004836138a1565b613b3b613ad291613b156024611c66565b90613b2860a435928361374960e4611c66565b613b326064611c66565b60843591614127565b610c49611435613b98613b516044969596611c66565b613b6b613b5c611074565b6001600160a01b039092168252565b60208181018881526040805193516001600160a01b03169284019290925251908201529182906060820190565b613ba0611074565b600181529060208201525f516020614b0d5f395f51905f52545f516020614c8d5f395f51905f52549093906001600160a01b031690613bdf6064611c66565b90613beb613865611074565b6020820152610c7260c43592610c57610c396119ff565b60ff5f516020614d0d5f395f51905f525460401c1615613c1e57565b631afcd79f60e31b5f5260045ffd5b613c35613c02565b805160208201516040830151936001600160a01b03918216939092909116613c5b613c02565b613c63613c02565b613c6b613c02565b613c73613c02565b60015f516020614cad5f395f51905f5255613c8c613c02565b613c94613c02565b8015613d475760405163301fd1f560e21b815293602085600481855afa948515611d7b576104c396610ced613d1b94613cdc601299613d16955f91613d28575b501515614aa5565b60018060a01b03166bffffffffffffffffffffffff60a01b5f516020614ced5f395f51905f525416175f516020614ced5f395f51905f5255565b614a3f565b613d236144da565b6144f5565b613d41915060203d602011611d7457611d668183611053565b5f613cd4565b60405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608490fd5b613dba613c02565b6064811015613de7577f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50055565b60405162461bcd60e51b815260206004820152602560248201527f4e6174697665546f6b656e52656d6f74653a20696e76616c69642070657263656044820152646e7461676560d81b6064820152608490fd5b15613e4157565b60405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608490fd5b9081526001600160a01b0391821660208201529116604082015260806060820181905261038992910190610353565b613ed482306120d4565b80516020820151909190613f22906001600160a01b03166040830151909390610bcf906001600160a01b031660a08501519060405196879463161b12ff60e11b602087015260248601613e9b565b60c08101516060820180519093613f4692909186906001600160a01b03169161492e565b15613f855750516040519182526001600160a01b0316907f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff490602090a2565b90516040518381526104c39392613fdc92610f2b9260e092916001600160a01b0316907fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb090602090a201516001600160a01b031690565b612060565b6040518281526104c39291906001600160a01b038216907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b90602090a26120d4565b1561402a57565b60405162461bcd60e51b815260206004820152603a60248201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60448201527f20746f6b656e207472616e7366657272657220616464726573730000000000006064820152608490fd5b5f516020614c6d5f395f51905f52541461410c575b506001600160a01b0316156140bb57565b60405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b6064820152608490fd5b614121906001600160a01b0316301415614023565b5f6140aa565b91939293824710614219575f8080808662010203600160981b015af13d15614214573d614153816110d0565b906141616040519283611053565b81525f60203d92013e5b156120a957614179916127c8565b926141b05f516020614bcd5f395f51905f52549160ff5f516020614c4d5f395f51905f5254166141aa858286612cd8565b93612cd8565b10156141ba579190565b60405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b6064820152608490fd5b61416b565b824763cf47918160e01b5f5260045260245260445ffd5b5f516020614c8d5f395f51905f5254614256916001600160a01b03908116911614614023565b6142bf576001600160a01b031661426957565b60405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b6064820152608490fd5b601f811161431c575050565b5f516020614b2d5f395f51905f525f5260205f20906020601f840160051c83019310614362575b601f0160051c01905b818110614357575050565b5f815560010161434c565b9091508190614343565b601f821161437957505050565b5f5260205f20906020601f840160051c830193106143b1575b601f0160051c01905b8181106143a6575050565b5f815560010161439b565b9091508190614392565b9081516001600160401b038111611033576143fa816143e75f516020614b8d5f395f51905f5254611701565b5f516020614b8d5f395f51905f5261436c565b602092601f82116001146144465761442a929382915f9261443b575b50508160011b915f199060031b1c19161790565b5f516020614b8d5f395f51905f5255565b015190505f80614416565b5f516020614b8d5f395f51905f525f52601f198216937f46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aa915f5b8681106144c257508360019596106144aa575b505050811b015f516020614b8d5f395f51905f5255565b01515f1960f88460031b161c191690555f8080614493565b91926020600181928685015181550194019201614480565b6144e2613c02565b60015f516020614c0d5f395f51905f5255565b91906144ff613c02565b60405163084279ef60e31b8152916020836004816005600160991b015afa928315611d7b576104c39461466e6146ae946145536146db97614675955f916146ec575b505f516020614c6d5f395f51905f5255565b61462e6145f4606085019461456a8651151561470b565b61458586515f516020614c6d5f395f51905f5254141561476a565b60808101805190916145e79160a091906145a9906001600160a01b031615156147dc565b01966145c5601260ff6145bd8b5160ff1690565b161115614834565b6145d5601260ff8c161115614892565b515f516020614b0d5f395f51905f5255565b516001600160a01b031690565b60018060a01b03166bffffffffffffffffffffffff60a01b5f516020614c8d5f395f51905f525416175f516020614c8d5f395f51905f5255565b614643815f516020614ccd5f395f51905f5255565b1560ff80195f516020614b6d5f395f51905f52541691151516175f516020614b6d5f395f51905f5255565b5160ff1690565b5f516020614c8d5f395f51905f52805460ff60a81b60a885901b1661ffff60a01b1990911660ff60a01b60a085901b1617179055614a67565b91909160ff80195f516020614c4d5f395f51905f52541691151516175f516020614c4d5f395f51905f5255565b5f516020614bcd5f395f51905f5255565b614705915060203d602011611d7457611d668183611053565b5f614541565b1561471257565b60405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b6064820152608490fd5b1561477157565b60405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d6500000000006064820152608490fd5b156147e357565b60405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b6064820152608490fd5b1561483b57565b60405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608490fd5b1561489957565b60405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b6064820152608490fd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526104c391612ef0608483611053565b929092805a106149a25783471061495d57823b15614955575f93849360208451940192f190565b505050505f90565b60405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606490fd5b60405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606490fd5b905f602091828151910182855af115611d9e575f513d614a3657506001600160a01b0381163b155b614a165750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415614a0f565b6104c390614a4b613c02565b6116d6613c02565b9060ff8091169116039060ff821161065157565b60ff8181168382161193929091908415614a965790614a8591614a53565b165b604d811161065157600a0a9190565b614a9f91614a53565b16614a87565b15614aac57565b60405162461bcd60e51b815260206004820152603260248201527f54656c65706f7274657252656769737472794170703a20696e76616c69642054604482015271656c65706f7274657220726567697374727960701b6064820152608490fdfe600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0152c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0369a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf501600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0652c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace049016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d02600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef04600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef00600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef029b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef05de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a164736f6c634300081e000a52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace04de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d02600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef04f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054656c65706f7274657252656769737472794170703a20696e76616c69642054",
}

// NativeTokenRemoteABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenRemoteMetaData.ABI instead.
var NativeTokenRemoteABI = NativeTokenRemoteMetaData.ABI

// NativeTokenRemoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenRemoteMetaData.Bin instead.
var NativeTokenRemoteBin = NativeTokenRemoteMetaData.Bin

// DeployNativeTokenRemote deploys a new Ethereum contract, binding an instance of NativeTokenRemote to it.
func DeployNativeTokenRemote(auth *bind.TransactOpts, backend bind.ContractBackend, settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (common.Address, *types.Transaction, *NativeTokenRemote, error) {
	parsed, err := NativeTokenRemoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenRemoteBin), backend, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenRemote{NativeTokenRemoteCaller: NativeTokenRemoteCaller{contract: contract}, NativeTokenRemoteTransactor: NativeTokenRemoteTransactor{contract: contract}, NativeTokenRemoteFilterer: NativeTokenRemoteFilterer{contract: contract}}, nil
}

// NativeTokenRemote is an auto generated Go binding around an Ethereum contract.
type NativeTokenRemote struct {
	NativeTokenRemoteCaller     // Read-only binding to the contract
	NativeTokenRemoteTransactor // Write-only binding to the contract
	NativeTokenRemoteFilterer   // Log filterer for contract events
}

// NativeTokenRemoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenRemoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenRemoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenRemoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenRemoteSession struct {
	Contract     *NativeTokenRemote // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativeTokenRemoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenRemoteCallerSession struct {
	Contract *NativeTokenRemoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativeTokenRemoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenRemoteTransactorSession struct {
	Contract     *NativeTokenRemoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativeTokenRemoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenRemoteRaw struct {
	Contract *NativeTokenRemote // Generic contract binding to access the raw methods on
}

// NativeTokenRemoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenRemoteCallerRaw struct {
	Contract *NativeTokenRemoteCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenRemoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenRemoteTransactorRaw struct {
	Contract *NativeTokenRemoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenRemote creates a new instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemote(address common.Address, backend bind.ContractBackend) (*NativeTokenRemote, error) {
	contract, err := bindNativeTokenRemote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemote{NativeTokenRemoteCaller: NativeTokenRemoteCaller{contract: contract}, NativeTokenRemoteTransactor: NativeTokenRemoteTransactor{contract: contract}, NativeTokenRemoteFilterer: NativeTokenRemoteFilterer{contract: contract}}, nil
}

// NewNativeTokenRemoteCaller creates a new read-only instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenRemoteCaller, error) {
	contract, err := bindNativeTokenRemote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCaller{contract: contract}, nil
}

// NewNativeTokenRemoteTransactor creates a new write-only instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenRemoteTransactor, error) {
	contract, err := bindNativeTokenRemote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTransactor{contract: contract}, nil
}

// NewNativeTokenRemoteFilterer creates a new log filterer instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenRemoteFilterer, error) {
	contract, err := bindNativeTokenRemote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteFilterer{contract: contract}, nil
}

// bindNativeTokenRemote binds a generic wrapper to an already deployed contract.
func bindNativeTokenRemote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenRemoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemote.Contract.NativeTokenRemoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.NativeTokenRemoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.NativeTokenRemoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemote *NativeTokenRemoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemote *NativeTokenRemoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemote *NativeTokenRemoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BURNEDFORTRANSFERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "BURNED_FOR_TRANSFER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemote.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) HOMECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "HOME_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemote.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemote.Contract.NATIVEMINTER(&_NativeTokenRemote.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemote.Contract.NATIVEMINTER(&_NativeTokenRemote.CallOpts)
}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) NATIVETOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "NATIVE_TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) NATIVETOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.NATIVETOKENREMOTESTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) NATIVETOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.NATIVETOKENREMOTESTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.TOKENREMOTESTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemote.Contract.TOKENREMOTESTORAGELOCATION(&_NativeTokenRemote.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.Allowance(&_NativeTokenRemote.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.Allowance(&_NativeTokenRemote.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.BalanceOf(&_NativeTokenRemote.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.BalanceOf(&_NativeTokenRemote.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemote.Contract.CalculateNumWords(&_NativeTokenRemote.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemote.Contract.CalculateNumWords(&_NativeTokenRemote.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteSession) Decimals() (uint8, error) {
	return _NativeTokenRemote.Contract.Decimals(&_NativeTokenRemote.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Decimals() (uint8, error) {
	return _NativeTokenRemote.Contract.Decimals(&_NativeTokenRemote.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetInitialReserveImbalance(&_NativeTokenRemote.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetInitialReserveImbalance(&_NativeTokenRemote.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemote.Contract.GetIsCollateralized(&_NativeTokenRemote.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemote.Contract.GetIsCollateralized(&_NativeTokenRemote.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetMinTeleporterVersion(&_NativeTokenRemote.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetMinTeleporterVersion(&_NativeTokenRemote.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemote.Contract.GetMultiplyOnRemote(&_NativeTokenRemote.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemote.Contract.GetMultiplyOnRemote(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeAddress(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeAddress(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTokenMultiplier(&_NativeTokenRemote.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTokenMultiplier(&_NativeTokenRemote.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTotalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTotalMinted(&_NativeTokenRemote.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTotalMinted(&_NativeTokenRemote.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemote.Contract.IsTeleporterAddressPaused(&_NativeTokenRemote.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemote.Contract.IsTeleporterAddressPaused(&_NativeTokenRemote.CallOpts, teleporterAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteSession) Name() (string, error) {
	return _NativeTokenRemote.Contract.Name(&_NativeTokenRemote.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Name() (string, error) {
	return _NativeTokenRemote.Contract.Name(&_NativeTokenRemote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) Owner() (common.Address, error) {
	return _NativeTokenRemote.Contract.Owner(&_NativeTokenRemote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Owner() (common.Address, error) {
	return _NativeTokenRemote.Contract.Owner(&_NativeTokenRemote.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteSession) Symbol() (string, error) {
	return _NativeTokenRemote.Contract.Symbol(&_NativeTokenRemote.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Symbol() (string, error) {
	return _NativeTokenRemote.Contract.Symbol(&_NativeTokenRemote.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalNativeAssetSupply(&_NativeTokenRemote.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalNativeAssetSupply(&_NativeTokenRemote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalSupply(&_NativeTokenRemote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalSupply(&_NativeTokenRemote.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Approve(&_NativeTokenRemote.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Approve(&_NativeTokenRemote.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Deposit(&_NativeTokenRemote.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Deposit(&_NativeTokenRemote.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "initialize", settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Initialize(&_NativeTokenRemote.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Initialize(&_NativeTokenRemote.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.PauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.PauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReceiveTeleporterMessage(&_NativeTokenRemote.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReceiveTeleporterMessage(&_NativeTokenRemote.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RegisterWithHome(&_NativeTokenRemote.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RegisterWithHome(&_NativeTokenRemote.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RenounceOwnership(&_NativeTokenRemote.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RenounceOwnership(&_NativeTokenRemote.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReportBurnedTxFees(&_NativeTokenRemote.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReportBurnedTxFees(&_NativeTokenRemote.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Send(&_NativeTokenRemote.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Send(&_NativeTokenRemote.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.SendAndCall(&_NativeTokenRemote.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.SendAndCall(&_NativeTokenRemote.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Transfer(&_NativeTokenRemote.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Transfer(&_NativeTokenRemote.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferFrom(&_NativeTokenRemote.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferFrom(&_NativeTokenRemote.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferOwnership(&_NativeTokenRemote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferOwnership(&_NativeTokenRemote.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UnpauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UnpauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemote.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemote.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Withdraw(&_NativeTokenRemote.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Withdraw(&_NativeTokenRemote.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Fallback(&_NativeTokenRemote.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Fallback(&_NativeTokenRemote.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Receive(&_NativeTokenRemote.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Receive(&_NativeTokenRemote.TransactOpts)
}

// NativeTokenRemoteApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenRemote contract.
type NativeTokenRemoteApprovalIterator struct {
	Event *NativeTokenRemoteApproval // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteApproval)
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
		it.Event = new(NativeTokenRemoteApproval)
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
func (it *NativeTokenRemoteApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteApproval represents a Approval event raised by the NativeTokenRemote contract.
type NativeTokenRemoteApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenRemoteApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteApprovalIterator{contract: _NativeTokenRemote.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteApproval)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseApproval(log types.Log) (*NativeTokenRemoteApproval, error) {
	event := new(NativeTokenRemoteApproval)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallFailedIterator struct {
	Event *NativeTokenRemoteCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteCallFailed)
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
		it.Event = new(NativeTokenRemoteCallFailed)
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
func (it *NativeTokenRemoteCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteCallFailed represents a CallFailed event raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCallFailedIterator{contract: _NativeTokenRemote.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteCallFailed)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "CallFailed", log); err != nil {
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

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseCallFailed(log types.Log) (*NativeTokenRemoteCallFailed, error) {
	event := new(NativeTokenRemoteCallFailed)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallSucceededIterator struct {
	Event *NativeTokenRemoteCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteCallSucceeded)
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
		it.Event = new(NativeTokenRemoteCallSucceeded)
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
func (it *NativeTokenRemoteCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteCallSucceeded represents a CallSucceeded event raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCallSucceededIterator{contract: _NativeTokenRemote.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteCallSucceeded)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenRemoteCallSucceeded, error) {
	event := new(NativeTokenRemoteCallSucceeded)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenRemote contract.
type NativeTokenRemoteDepositIterator struct {
	Event *NativeTokenRemoteDeposit // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteDeposit)
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
		it.Event = new(NativeTokenRemoteDeposit)
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
func (it *NativeTokenRemoteDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteDeposit represents a Deposit event raised by the NativeTokenRemote contract.
type NativeTokenRemoteDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteDepositIterator{contract: _NativeTokenRemote.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteDeposit)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseDeposit(log types.Log) (*NativeTokenRemoteDeposit, error) {
	event := new(NativeTokenRemoteDeposit)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenRemote contract.
type NativeTokenRemoteInitializedIterator struct {
	Event *NativeTokenRemoteInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteInitialized)
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
		it.Event = new(NativeTokenRemoteInitialized)
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
func (it *NativeTokenRemoteInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteInitialized represents a Initialized event raised by the NativeTokenRemote contract.
type NativeTokenRemoteInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenRemoteInitializedIterator, error) {

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteInitializedIterator{contract: _NativeTokenRemote.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteInitialized)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseInitialized(log types.Log) (*NativeTokenRemoteInitialized, error) {
	event := new(NativeTokenRemoteInitialized)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenRemote contract.
type NativeTokenRemoteMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenRemoteMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenRemoteMinTeleporterVersionUpdated)
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
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenRemote contract.
type NativeTokenRemoteMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenRemoteMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteMinTeleporterVersionUpdatedIterator{contract: _NativeTokenRemote.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteMinTeleporterVersionUpdated)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenRemoteMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenRemoteMinTeleporterVersionUpdated)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenRemote contract.
type NativeTokenRemoteOwnershipTransferredIterator struct {
	Event *NativeTokenRemoteOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteOwnershipTransferred)
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
		it.Event = new(NativeTokenRemoteOwnershipTransferred)
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
func (it *NativeTokenRemoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenRemote contract.
type NativeTokenRemoteOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenRemoteOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteOwnershipTransferredIterator{contract: _NativeTokenRemote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteOwnershipTransferred)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenRemoteOwnershipTransferred, error) {
	event := new(NativeTokenRemoteOwnershipTransferred)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenRemote contract.
type NativeTokenRemoteReportBurnedTxFeesIterator struct {
	Event *NativeTokenRemoteReportBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteReportBurnedTxFees)
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
		it.Event = new(NativeTokenRemoteReportBurnedTxFees)
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
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenRemote contract.
type NativeTokenRemoteReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenRemoteReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteReportBurnedTxFeesIterator{contract: _NativeTokenRemote.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteReportBurnedTxFees)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
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

// ParseReportBurnedTxFees is a log parse operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenRemoteReportBurnedTxFees, error) {
	event := new(NativeTokenRemoteReportBurnedTxFees)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressPausedIterator struct {
	Event *NativeTokenRemoteTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTeleporterAddressPaused)
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
		it.Event = new(NativeTokenRemoteTeleporterAddressPaused)
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
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTeleporterAddressPausedIterator{contract: _NativeTokenRemote.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTeleporterAddressPaused)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenRemoteTeleporterAddressPaused, error) {
	event := new(NativeTokenRemoteTeleporterAddressPaused)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenRemoteTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenRemoteTeleporterAddressUnpaused)
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
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTeleporterAddressUnpausedIterator{contract: _NativeTokenRemote.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTeleporterAddressUnpaused)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenRemoteTeleporterAddressUnpaused, error) {
	event := new(NativeTokenRemoteTeleporterAddressUnpaused)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensAndCallSentIterator struct {
	Event *NativeTokenRemoteTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensAndCallSent)
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
		it.Event = new(NativeTokenRemoteTokensAndCallSent)
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
func (it *NativeTokenRemoteTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensAndCallSentIterator{contract: _NativeTokenRemote.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensAndCallSent)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenRemoteTokensAndCallSent, error) {
	event := new(NativeTokenRemoteTokensAndCallSent)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensSentIterator struct {
	Event *NativeTokenRemoteTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensSent)
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
		it.Event = new(NativeTokenRemoteTokensSent)
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
func (it *NativeTokenRemoteTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensSent represents a TokensSent event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensSentIterator{contract: _NativeTokenRemote.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensSent)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensSent", log); err != nil {
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

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensSent(log types.Log) (*NativeTokenRemoteTokensSent, error) {
	event := new(NativeTokenRemoteTokensSent)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensWithdrawnIterator struct {
	Event *NativeTokenRemoteTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensWithdrawn)
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
		it.Event = new(NativeTokenRemoteTokensWithdrawn)
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
func (it *NativeTokenRemoteTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenRemoteTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensWithdrawnIterator{contract: _NativeTokenRemote.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensWithdrawn)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenRemoteTokensWithdrawn, error) {
	event := new(NativeTokenRemoteTokensWithdrawn)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTransferIterator struct {
	Event *NativeTokenRemoteTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTransfer)
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
		it.Event = new(NativeTokenRemoteTransfer)
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
func (it *NativeTokenRemoteTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTransfer represents a Transfer event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenRemoteTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTransferIterator{contract: _NativeTokenRemote.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTransfer)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTransfer(log types.Log) (*NativeTokenRemoteTransfer, error) {
	event := new(NativeTokenRemoteTransfer)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenRemote contract.
type NativeTokenRemoteWithdrawalIterator struct {
	Event *NativeTokenRemoteWithdrawal // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteWithdrawal)
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
		it.Event = new(NativeTokenRemoteWithdrawal)
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
func (it *NativeTokenRemoteWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteWithdrawal represents a Withdrawal event raised by the NativeTokenRemote contract.
type NativeTokenRemoteWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteWithdrawalIterator{contract: _NativeTokenRemote.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteWithdrawal)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseWithdrawal(log types.Log) (*NativeTokenRemoteWithdrawal, error) {
	event := new(NativeTokenRemoteWithdrawal)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
