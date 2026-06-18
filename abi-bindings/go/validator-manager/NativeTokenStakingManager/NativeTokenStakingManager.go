// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenstakingmanager

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

// ConversionData is an auto generated low-level Go binding around an user-defined struct.
type ConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// Delegator is an auto generated low-level Go binding around an user-defined struct.
type Delegator struct {
	Status        uint8
	Owner         common.Address
	ValidationID  [32]byte
	Weight        uint64
	StartTime     uint64
	StartingNonce uint64
	EndingNonce   uint64
}

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorInfo struct {
	Owner             common.Address
	DelegationFeeBips uint16
	MinStakeDuration  uint64
	UptimeSeconds     uint64
}

// StakingManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerSettings struct {
	Manager                  common.Address
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	WeightToValueFactor      *big.Int
	RewardCalculator         common.Address
	UptimeBlockchainID       [32]byte
}

// ValidatorMessagesValidationPeriod is an auto generated low-level Go binding around an user-defined struct.
type ValidatorMessagesValidationPeriod struct {
	SubnetID              [32]byte
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
	Weight                uint64
}

// NativeTokenStakingManagerMetaData contains all meta data concerning the NativeTokenStakingManager contract.
var NativeTokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedValidationID\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"}],\"name\":\"DelegatorRewardRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InitiatedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InitiatedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InitiatedStakingValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"}],\"name\":\"ValidatorRewardRecipientChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startingNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endingNonce\",\"type\":\"uint64\"}],\"internalType\":\"structDelegator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"getDelegatorRewardInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingManagerSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structStakingManagerSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getStakingValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uptimeSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structPoSValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorRewardInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structStakingManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100fd57601f61390738819003918201601f19168301916001600160401b03831184841017610101578084926020946040528339810103126100fd575160028110156100fd5760011461005f575b6040516137d190816101168239f35b5f5160206138e75f395f51905f525460ff8160401c166100ee576002600160401b03196001600160401b03821601610098575b50610050565b6001600160401b0319166001600160401b039081175f5160206138e75f395f51905f52556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f610092565b63f92ee8a960e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80630d43631714611806578063134096451461158b578063151d30d1146115705780631667956414611559578063243ca19a146111c0578063245dafcb14610f4057806325e1c77614610e725780632674874b14610dbb57806327bf60cd14610da45780632aa5663814610d715780632e2194d814610d42578063329c3e1214610d2057806335455ded1461089a57806353a1333814610c0857806360ad778414610be45780636206585614610b9c5780637a63ad8514610b755780638ef34c9814610ac557806393e24598146109885780639681d94014610944578063a3a65e481461089f578063a9778a7a1461089a578063ad93936d146104f6578063af1dd66c14610497578063b2c1712e14610464578063b771b3bc14610442578063caa7187414610293578063e24b2680146102345763fb8b11dd14610158575f80fd5b34610230576040366003190112610230576004356001600160a01b0361017c611cb5565b16801561021e575f8281525f5160206137055f395f51905f52602052604090205460081c6001600160a01b0316330361020b575f8281525f5160206136c55f395f51905f526020526040812080546001600160a01b0319811684179091556001600160a01b031692907f6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e0059080a4005b636e2ccd7560e11b5f523360045260245ffd5b63caa903f960e01b5f5260045260245ffd5b5f80fd5b34610230576020366003190112610230576004355f9081525f5160206136655f395f51905f5260209081526040808320545f5160206136a55f395f51905f528352928190205481516001600160a01b0390941684529183019190915290f35b34610230575f366003190112610230575f6101006040516102b381611d3f565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152015261012060018060a01b035f5160206137455f395f51905f5254167fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab601547fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab602545f5160206136455f395f51905f525460ff5f5160206137255f395f51905f52549161ffff60018060a01b035f5160206137a55f395f51905f525416946001600160401b037fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6065497604051926103b484611d3f565b8a84526020840191825260408401908152606084019183871683526101006080860195878960401c1687528960a082019960501c16895260c081019a8b5260e081019b8c52019a8b526040519b8c525160208c01525160408b01525116606089015251166080870152511660a08501525160c084015260018060a01b0390511660e083015251610100820152f35b34610230575f366003190112610230576040516005600160991b018152602090f35b346102305761047e61047536611c85565b908293926128a5565b1561048557005b635bff683f60e11b5f5260045260245ffd5b34610230576020366003190112610230576004355f9081525f5160206136c55f395f51905f5260209081526040808320545f5160206136855f395f51905f528352928190205481516001600160a01b0390941684529183019190915290f35b60e0366003190112610230576004356001600160401b03811161023057610521903690600401611d97565b6024356001600160401b03811161023057610540903690600401611d97565b6044356001600160401b0381116102305761055f903690600401611ddd565b916064356001600160401b0381116102305761057f903690600401611ddd565b906084359261ffff8416928385036102305760a435906001600160401b038216928383036102305760c4356001600160a01b0381169590869003610230576105c56124fa565b5f5160206136455f395f51905f525461ffff8160401c168810801561088f575b61087c576001600160401b0316851061086a577fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6015434108015610840575b61082d57851561081a576106363461204d565b5f5160206137455f395f51905f5254604051634e5bb12760e11b815260a060048201529a6001600160a01b03909116948b94859491929161067b9060a4870190612028565b85810360031901602487015261069091612028565b8481036003190160448601526106a591613311565b8381036003190160648501526106ba91613311565b906001600160401b0316608483015203815a6020945f91f194851561080f575f956107da575b5f8681525f5160206136e55f395f51905f5260209081526040808320805469ffffffffffffffffffff60a01b19339081166001600160f01b03199092169190911760a09690961b61ffff60a01b169590951760b09690961b67ffffffffffffffff60b01b169590951785556001909401805467ffffffffffffffff191690555f5160206136655f395f51905f5281529083902080546001600160a01b031916861790558251958652858101939093529084019290925293509082907ff51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f90606090a360015f5160206137655f395f51905f5255604051908152f35b94506020863d602011610807575b816107f560209383611d5b565b810103126102305760209551946106e0565b3d91506107e8565b6040513d5f823e3d90fd5b8563caa903f960e01b5f5260045260245ffd5b63222d164360e21b5f523460045260245ffd5b507fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab602543411610623565b846202a06d60e11b5f5260045260245ffd5b87635f12e6c360e11b5f5260045260245ffd5b5061271088116105e5565b611ccb565b34610230576020366003190112610230576108b8611c72565b5f5160206137455f395f51905f5254604051631474cbc960e31b815263ffffffff929092166004830152602090829060249082905f906001600160a01b03165af1801561080f575f90610911575b602090604051908152f35b506020813d60201161093c575b8161092b60209383611d5b565b810103126102305760209051610906565b3d915061091e565b3461023057602036600319011261023057602061096f610962611c72565b61096a6124fa565b61235c565b60015f5160206137655f395f51905f5255604051908152f35b34610230576020366003190112610230575f5160206137455f395f51905f5254604051636af907fb60e11b815260048035908201819052915f90829060249082906001600160a01b03165afa90811561080f575f91610aa3575b50516006811015610a8f5760048103610a7d57505f8181525f5160206136e55f395f51905f5260205260409020546001600160a01b0316330361020b575f8181525f5160206136655f395f51905f526020526040902054610a5291906001600160a01b03168015610a545761323a565b005b505f8181525f5160206136e55f395f51905f5260205260409020546001600160a01b031661323a565b63170cc93360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b610abf91503d805f833e610ab78183611d5b565b810190611f20565b826109e2565b34610230576040366003190112610230576004356001600160a01b03610ae9611cb5565b16801561021e575f8281525f5160206136e55f395f51905f5260205260409020546001600160a01b0316330361020b575f8281525f5160206136655f395f51905f526020526040812080546001600160a01b0319811684179091556001600160a01b031692907f28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c339080a4005b34610230575f3660031901126102305760206040515f5160206137455f395f51905f528152f35b34610230576020366003190112610230576004356001600160401b03811680910361023057610bdc6020915f5160206137255f395f51905f525490612349565b604051908152f35b3461023057604036600319011261023057610a52610c00611c5f565b6004356120b5565b34610230576020366003190112610230575f60c0604051610c2881611cf5565b8281528260208201528260408201528260608201528260808201528260a082015201526004355f525f5160206137055f395f51905f5260205260405f2060405190610c7282611cf5565b80549160ff8316916004831015610a8f576001600160401b03828160e09686829652602083019060018060a01b039060081c16815281600260018701549660408601978852015495606085019082881682526080860193838960401c16855260c060a0880197858b60801c168952019860c01c89526040519a8b5260018060a01b0390511660208b01525160408a01525116606088015251166080860152511660a0840152511660c0820152f35b34610230575f366003190112610230576040516001600160991b018152602090f35b34610230576020366003190112610230576020610d6060043561204d565b6001600160401b0360405191168152f35b3461023057610d8b610d8236611c85565b90829392612e85565b15610d9257005b631036cf9160e11b5f5260045260245ffd5b3461023057610a52610db536611c85565b91612e85565b34610230576020366003190112610230575f6060604051610ddb81611d24565b82815282602082015282604082015201526004355f525f5160206136e55f395f51905f52602052608060405f206001600160401b03604051610e1c81611d24565b81835461ffff60018060a01b0382169586855260608460016020880193858760a01c1685528260408a019760b01c1687520154169501948552604051968752511660208601525116604084015251166060820152f35b3461023057604036600319011261023057600435610e8e611c5f565b610e9782612ba2565b15610f2d575f5160206137455f395f51905f5254604051636af907fb60e11b815260048101849052905f90829060249082906001600160a01b03165afa90811561080f575f91610f13575b5051916006831015610a8f5760028303610f0057610a529250612bca565b8263170cc93360e21b5f5260045260245ffd5b610f2791503d805f833e610ab78183611d5b565b83610ee2565b506330efa98b60e01b5f5260045260245ffd5b3461023057602036600319011261023057600435805f525f5160206137055f395f51905f5260205260405f20604051610f7881611cf5565b815460ff81166004811015610a8f57600191818452828060a01b039060081c166020840152600282850154946040850195865201546001600160401b03811660608501526001600160401b038160401c1660808501526001600160401b038160801c1660a085015260c01c60c08401521415806111ab575b611189575060018060a01b035f5160206137455f395f51905f525416905f8151602460405180958193636af907fb60e11b835260048301525afa91821561080f575f9261116d575b5060608201926001600160401b038451161561115b575060a06001600160401b0380925194511692015116906040519263854a893f60e01b84526004840152602483015260448201525f8160648173__$e8433aec8a689eedbe2673465b6d21420c$__5af490811561080f575f91611114575b5060206110d3916040518093819263ee5b48eb60e01b83528460048401526024830190612028565b03815f6005600160991b015af1801561080f576110ec57005b610a529060203d60201161110d575b6111058183611d5b565b810190612019565b503d6110fb565b90503d805f833e6111258183611d5b565b810190602081830312610230578051906001600160401b038211610230576110d3926020926111549201611ec7565b91506110ab565b6339b894f960e21b5f5260045260245ffd5b6111829192503d805f833e610ab78183611d5b565b9083611038565b516004811015610a8f57633b0d540d60e21b5f526111a690611ce7565b60245ffd5b5080516004811015610a8f5760031415610ff0565b6040366003190112610230576004356111d7611cb5565b906111e06124fa565b6111e93461204d565b916111f382612ba2565b15610f2d576001600160a01b0316908115611546575f5160206137455f395f51905f5254604051636af907fb60e11b81526004810183905292906001600160a01b03165f84602481845afa93841561080f575f9461152a575b506001600160401b036040611267878360a089015116612532565b950151166001600160401b035f5160206136455f395f51905f525460501c1602936001600160401b038516948503611516576001600160401b0381169485116115035760408051636610966960e01b8152600481018690526001600160401b0392909216602483015290939291849060449082905f905af1801561080f576020955f945f926114c5575b506001600160401b0390604051888101908682528360c01b8860c01b16604082015260288152611322604882611d5b565b51902096875f525f5160206137055f395f51905f52895260405f20600160ff19825416179055875f525f5160206137055f395f51905f52895260405f208054610100600160a81b033360081b1690610100600160a81b031916179055875f525f5160206137055f395f51905f52895285600160405f200155875f525f5160206137055f395f51905f528952600260405f20018383168419825416179055875f525f5160206137055f395f51905f528952600260405f200167ffffffffffffffff60401b198154169055875f525f5160206137055f395f51905f528952600260405f200180548460801b8960801b16908560801b1916179055875f525f5160206137055f395f51905f528952600260405f200160018060c01b038154169055875f525f5160206136c55f395f51905f52895260405f20856001600160601b0360a01b82541617905582604051971687528887015216604085015260608401526080830152827f77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d7760a03394a460015f5160206137655f395f51905f5255604051908152f35b6001600160401b0395506114f291925060403d6040116114fc575b6114ea8183611d5b565b810190612b85565b94909491906112f1565b503d6114e0565b84636d51fe0560e11b5f5260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b61153f9194503d805f833e610ab78183611d5b565b928561124c565b5063caa903f960e01b5f5260045260245ffd5b3461023057610a5261156a36611c85565b916128a5565b34610230575f366003190112610230576020604051600a8152f35b34610230576040366003190112610230576004356115a7611c5f565b906115b06124fa565b805f525f5160206137055f395f51905f5260205260405f2091604051906115d682611cf5565b83549160ff8316926004841015610a8f5783825260018060a01b039060081c16602082015260036002600187015496604084019788520154936001600160401b03851660608401526001600160401b038560401c1660808401526001600160401b038560801c1660a084015260c083019460c01c855203611189575060018060a01b035f5160206137455f395f51905f52541690845160405190636af907fb60e11b825260048201525f81602481865afa90811561080f575f916117ec575b50855160405190636af907fb60e11b825260048201525f81602481875afa90811561080f575f916117d2575b50516006811015610a8f576004141590816117b1575b506116f9575b6116e684612552565b60015f5160206137655f395f51905f5255005b63ffffffff60246040925f8451958694859363338587c560e21b85521660048401525af193841561080f575f915f9561177e575b5051908082036117695750506001600160401b03809151169216809211611756578180806116dd565b50632e19bc2d60e11b5f5260045260245ffd5b63fee3144560e01b5f5260045260245260445ffd5b9094506117a3915060403d6040116117aa575b61179b8183611d5b565b810190611fff565b938561172d565b503d611791565b6001600160401b03915060800151166001600160401b0384511611866116d7565b6117e691503d805f833e610ab78183611d5b565b876116c1565b61180091503d805f833e610ab78183611d5b565b86611695565b3461023057610120366003190112610230575f5160206137855f395f51905f525460ff8160401c1615906001600160401b03811680159081611c57575b6001149081611c4d575b159081611c44575b50611c355767ffffffffffffffff1981166001175f5160206137855f395f51905f525581611c09575b50611887613365565b61188f613365565b611897613365565b61189f613365565b60015f5160206137655f395f51905f52556004356001600160a01b038116908190036102305760243560443591606435926001600160401b038416809403610230576084359061ffff8216948583036102305760a4359160ff8316958684036102305760e4356001600160a01b0381169760c4359791899003610230576101043599611929613365565b80158015611bfe575b611bec5750838311611bd95780158015611bcf575b611bbd57508015611bae578715611bae576040516304e0efb360e11b8152602081600481855afa801561080f575f90611b6e575b6001600160401b039150168410611b5c578615611b4d578815611b3a576001600160601b0360a01b5f5160206137455f395f51905f525416175f5160206137455f395f51905f52557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab601557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab602556001600160401b03195f5160206136455f395f51905f525416175f5160206136455f395f51905f525569ffff00000000000000005f5160206136455f395f51905f52549160ff60501b9060501b169260401b169071ffffffffffffffffffff0000000000000000191617175f5160206136455f395f51905f52555f5160206137255f395f51905f52556001600160601b0360a01b5f5160206137a55f395f51905f525416175f5160206137a55f395f51905f52557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60655611ae357005b68ff0000000000000000195f5160206137855f395f51905f5254165f5160206137855f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b88632f6bd1db60e01b5f5260045260245ffd5b63a733007160e01b5f5260045ffd5b836202a06d60e11b5f5260045260245ffd5b506020813d602011611ba6575b81611b8860209383611d5b565b8101031261023057611ba16001600160401b0391611f0c565b61197b565b3d9150611b7b565b63d92e233d60e01b5f5260045ffd5b63170db35960e31b5f5260045260245ffd5b50600a8111611947565b8263222d164360e21b5f5260045260245ffd5b635f12e6c360e11b5f5260045260245ffd5b506127108111611932565b68ffffffffffffffffff191668010000000000000001175f5160206137855f395f51905f52558161187e565b63f92ee8a960e01b5f5260045ffd5b90501583611855565b303b15915061184d565b839150611843565b6024359063ffffffff8216820361023057565b6004359063ffffffff8216820361023057565b606090600319011261023057600435906024358015158103610230579060443563ffffffff811681036102305790565b602435906001600160a01b038216820361023057565b34610230575f3660031901126102305760206040516127108152f35b6004811015610a8f57600452565b60e081019081106001600160401b03821117611d1057604052565b634e487b7160e01b5f52604160045260245ffd5b608081019081106001600160401b03821117611d1057604052565b61012081019081106001600160401b03821117611d1057604052565b90601f801991011681019081106001600160401b03821117611d1057604052565b6001600160401b038111611d1057601f01601f191660200190565b81601f8201121561023057803590611dae82611d7c565b92611dbc6040519485611d5b565b8284526020838301011161023057815f926020809301838601378301015290565b919060408382031261023057604051604081018181106001600160401b03821117611d10576040528093803563ffffffff811681036102305782526020810135906001600160401b038211610230570182601f82011215610230578035926001600160401b038411611d10578360051b9160405194611e5f6020850187611d5b565b855260208086019382010191821161023057602001915b818310611e865750505060200152565b82356001600160a01b038116810361023057815260209283019201611e76565b5f5b838110611eb75750505f910152565b8181015183820152602001611ea8565b81601f82011215610230578051611edd81611d7c565b92611eeb6040519485611d5b565b8184526020828401011161023057611f099160208085019101611ea6565b90565b51906001600160401b038216820361023057565b602081830312610230578051906001600160401b038211610230570161010081830312610230576040519161010083018381106001600160401b03821117611d1057604052815160068110156102305783526020820151916001600160401b03831161023057611f9760e092611ff7948301611ec7565b6020850152611fa860408201611f0c565b6040850152611fb960608201611f0c565b6060850152611fca60808201611f0c565b6080850152611fdb60a08201611f0c565b60a0850152611fec60c08201611f0c565b60c085015201611f0c565b60e082015290565b919082604091031261023057611f09602083519301611f0c565b90816020910312610230575190565b9060209161204181518092818552858086019101611ea6565b601f01601f1916010190565b5f5160206137255f395f51905f525480156120a15781049081158015612091575b61207f57506001600160401b031690565b63222d164360e21b5f5260045260245ffd5b506001600160401b03821161206e565b634e487b7160e01b5f52601260045260245ffd5b805f525f5160206137055f395f51905f5260205260405f20916040516120da81611cf5565b835460ff81166004811015610a8f57825260018060a01b039060081c166020820152602460026001860154958660408501520154916001600160401b03831660608201526001600160401b038360401c16608082015260a08101926001600160401b038160801c16845260c01c60c08201525f60018060a01b035f5160206137455f395f51905f52541660405193848092636af907fb60e11b82528a60048301525afa91821561080f575f9261232d575b5080516004811015610a8f575f1901611189575080516006811015610a8f5760041461231e5760806001600160401b03910151166001600160401b038251161161225e575b50505f8181525f5160206137055f395f51905f526020908152604091829020805460ff1916600290811782550180546fffffffffffffffff000000000000000019164280851b67ffffffffffffffff60401b169190911790915591516001600160401b0390921682527f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa91a3565b5f5160206137455f395f51905f52546040805163338587c560e21b815263ffffffff94909416600485015290839060249082905f906001600160a01b03165af191821561080f575f905f936122fa575b508085036122e35750516001600160401b03918216911681106122d157806121d0565b632e19bc2d60e11b5f5260045260245ffd5b849063fee3144560e01b5f5260045260245260445ffd5b905061231691925060403d6040116117aa5761179b8183611d5b565b91905f6122ae565b50505061232b9150612552565b565b6123429192503d805f833e610ab78183611d5b565b905f61218b565b8181029291811591840414171561151657565b5f5160206137455f395f51905f525460405163025a076560e61b815263ffffffff929092166004830152602090829060249082905f906001600160a01b03165af190811561080f575f916124c8575b505f5160206137455f395f51905f5254604051636af907fb60e11b815260048101839052905f90829060249082906001600160a01b03165afa90811561080f575f916124ae575b506123fc82612ba2565b156124aa575f8281525f5160206136e55f395f51905f5260209081526040808320545f5160206136655f395f51905f52909252909120546001600160a01b039081169291169082156124a1575b8051926006841015610a8f57604061248b926001600160401b0392876004611f099814612491575b50500151165f5160206137255f395f51905f525490612349565b9061329d565b61249a9161323a565b5f87612471565b91508091612449565b5090565b6124c291503d805f833e610ab78183611d5b565b5f6123f2565b90506020813d6020116124f2575b816124e360209383611d5b565b8101031261023057515f6123ab565b3d91506124d6565b60025f5160206137655f395f51905f5254146125235760025f5160206137655f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b906001600160401b03809116911601906001600160401b03821161151657565b805f525f5160206137055f395f51905f5260205260405f209060405161257781611cf5565b825460ff81166004811015610a8f578252602082019060018060a01b039060081c1681526002600185015494856040850152015460608301926001600160401b038216845260c06001600160401b038360401c16928360808401526001600160401b038160801c1660a0840152811c9101526004602060018060a01b035f5160206137455f395f51905f525416604051928380926304e0efb360e11b82525afa90811561080f575f91612850575b50612638906001600160401b0392612532565b164210612834575f8381525f5160206137055f395f51905f5260209081526040808320838155600181018490556002018390555f5160206136c55f395f51905f52909152902080546001600160a01b031981169091556001600160a01b03168015612823575b5f915f91855f525f5160206136855f395f51905f5260205260405f205480612727575b50509061271a7f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e9461248b6001600160401b036040969560018060a01b039051169251165f5160206137255f395f51905f525490612349565b82519182526020820152a3565b9193509150845f525f5160206136855f395f51905f526020525f6040812055855f525f5160206136e55f395f51905f5260205261271061277261ffff60405f205460a01c1683612349565b049081875f525f5160206136a55f395f51905f5260205260405f20612798828254612898565b90558103908111611516577f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e9461248b6001600160401b03604096848a7f3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63602061271a989a61280785826135e9565b8c519485526001600160a01b031693a3949596505050946126c1565b5080516001600160a01b031661269e565b63fb6ce63f60e01b5f526001600160401b03421660045260245ffd5b90506020813d602011612890575b8161286b60209383611d5b565b81010312610230576001600160401b039161288861263892611f0c565b915091612625565b3d915061285e565b9190820180921161151657565b5f5160206137455f395f51905f525490925f916001600160a01b0316803b15610230575f8091602460405180948193635b73516560e11b83528a60048401525af1801561080f57612b6f575b505f5160206137455f395f51905f5254604051636af907fb60e11b81526004810186905291908390839060249082906001600160a01b03165afa918215612b64578392612b48575b5061294385612ba2565b15612b3e578483525f5160206136e55f395f51905f5260205260408320546001600160a01b03163303612b2b5760e08201906001600160401b038251169460c08401956001600160401b036129b9818951168a89525f5160206136e55f395f51905f526020528260408a205460b01c1690612532565b1611612b0e5791602093916001600160401b03935f14612aea576129dd9088612bca565b955b612a6b84612a1981604060018060a01b035f5160206137a55f395f51905f525416970151165f5160206137255f395f51905f525490612349565b92519351604051634f22429f60e01b815260048101949094526001600160401b0391909416811660248401819052604484015292909416821660648201529516608486015284918290819060a4820190565b03915afa918215612add578192612aa8575b506040919281525f5160206136a55f395f51905f5260205220612aa1828254612898565b9055151590565b91506020823d602011612ad5575b81612ac360209383611d5b565b81010312610230576040915191612a7d565b3d9150612ab6565b50604051903d90823e3d90fd5b508685525f5160206136e55f395f51905f52845282600160408720015416956129df565b825163fb6ce63f60e01b86526001600160401b0316600452602485fd5b636e2ccd7560e11b835233600452602483fd5b5050505050600190565b612b5d9192503d8085833e610ab78183611d5b565b905f612939565b6040513d85823e3d90fd5b612b7c9192505f90611d5b565b5f9060246128f1565b9190826040910312610230576020612b9c83611f0c565b92015190565b5f9081525f5160206136e55f395f51905f5260205260409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff9092166004830152905f816024816005600160991b015afa801561080f575f915f91612dd6575b5015612dc75780517fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab606548103612db5575060208101516001600160a01b031680612da4575090604080612c7393015181518094819263088c246360e01b8352602060048401526024830190612028565b038173__$e8433aec8a689eedbe2673465b6d21420c$__5af4801561080f575f925f91612d7f575b508092808303612d685750815f525f5160206136e55f395f51905f526020526001600160401b0380600160405f200154169116115f14612d3f57805f525f5160206136e55f395f51905f52602052600160405f20016001600160401b0383166001600160401b03198254161790557fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d043560206040516001600160401b0385168152a290565b90505f525f5160206136e55f395f51905f526020526001600160401b03600160405f2001541690565b905063fee3144560e01b5f5260045260245260445ffd5b9050612d9b91925060403d6040116117aa5761179b8183611d5b565b9190915f612c9b565b624de75d60e31b5f5260045260245ffd5b636ba589a560e01b5f5260045260245ffd5b636b2f19e960e01b5f5260045ffd5b9150503d805f833e612de88183611d5b565b81016040828203126102305781516001600160401b038111610230578201916060838303126102305760405192606084018481106001600160401b03821117611d10576040528051845260208101516001600160a01b03811681036102305760208501526040810151926001600160401b03841161023057602093612e6d9201611ec7565b6040840152015190811515820361023057905f612c03565b9190825f525f5160206137055f395f51905f5260205260405f209160405191612ead83611cf5565b835460ff8116926004841015610a8f576024938552602085019160018060a01b039060081c168252600260018701549687604088015201549260608601926001600160401b038516845260808701946001600160401b038160401c1686526001600160401b038160801c1660a089015260c01c60c08801525f60018060a01b035f5160206137455f395f51905f52541660405197888092636af907fb60e11b82528c60048301525afa95861561080f575f9661321e575b5086516004811015610a8f576001190161320057516001600160a01b0316330361317f575b875f525f5160206136c55f395f51905f5260205260018060a01b0360405f2054169385516006811015610a8f5760020361313057612fe36001600160401b0380925116825f5160206136455f395f51905f52541690612532565b1642106128345761311f575b50855f525f5160206137055f395f51905f5260205260405f20600360ff198254161790556001600160401b038060a0600180821b035f5160206137455f395f51905f525416950151169151169003916001600160401b0383116115165760408051636610966960e01b8152600481018790526001600160401b0394909416602485015290839060449082905f905af192831561080f576130d59386935f916130ff575b505f8481525f5160206137055f395f51905f526020526040902060020180546001600160c01b031660c09290921b6001600160c01b031916919091179055613390565b917f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf5f80a3151590565b613118915060403d6040116114fc576114ea8183611d5b565b505f613092565b6131299086612bca565b505f612fef565b50505050925080516006811015610a8f5760040361316357508261315e939261315892613390565b50612552565b600190565b516006811015610a8f5763170cc93360e21b5f5260045260245ffd5b865f525f5160206136e55f395f51905f526020523360018060a01b0360405f2054160361020b576001600160401b036131dc8160c088015116895f525f5160206136e55f395f51905f526020528260405f205460b01c1690612532565b16421015612f895763fb6ce63f60e01b5f526001600160401b03421660045260245ffd5b86516004811015610a8f57633b0d540d60e21b5f526111a690611ce7565b6132339196503d805f833e610ab78183611d5b565b945f612f64565b5f8281525f5160206136a55f395f51905f5260209081526040822080549290559092917f875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4919061328a85826135e9565b6040519485526001600160a01b031693a3565b8147106132fa575f918291829182916001600160a01b03165af13d156132f5573d6132c781611d7c565b906132d56040519283611d5b565b81525f60203d92013e5b156132e657565b63d6bda27560e01b5f5260045ffd5b6132df565b504763cf47918160e01b5f5260045260245260445ffd5b6020606081604085019363ffffffff81511686520151936040838201528451809452019201905f5b8181106133465750505090565b82516001600160a01b0316845260209384019390920191600101613339565b60ff5f5160206137855f395f51905f525460401c161561338157565b631afcd79f60e31b5f5260045ffd5b9160018060a01b035f5160206137455f395f51905f5254169260408101935f8551602460405180948193636af907fb60e11b835260048301525afa90811561080f575f916135cf575b5080516006811015610a8f5760031480156135bb575b15613599576001600160401b0360e082015116945b60808301916001600160401b038351166001600160401b038816111561358e5760209260018060a01b035f5160206137a55f395f51905f525416916001600160401b038060c06134698260608b0151165f5160206137255f395f51905f525490612349565b930151935195515f9081525f5160206136e55f395f51905f528852604090819020600101549051634f22429f60e01b8152600481019490945293166001600160401b0390811660248401529416841660448201529783166064890152919091166084870152859060a49082905afa93841561080f575f9461355a575b506001600160a01b03831615613545575b50805f525f5160206136855f395f51905f526020528260405f20555f525f5160206136c55f395f51905f5260205260405f209060018060a01b03166001600160601b0360a01b82541617905590565b602001516001600160a01b031691505f6134f6565b9093506020813d602011613586575b8161357660209383611d5b565b810103126102305751925f6134e5565b3d9150613569565b505050505050505f90565b80516006811015610a8f57600203613163576001600160401b03421694613404565b5080516006811015610a8f576004146133ef565b6135e391503d805f833e610ab78183611d5b565b5f6133d9565b6001600160991b013b15610230576040516327ad555d60e11b81526001600160a01b0391909116600482015260248101919091525f81604481836001600160991b015af1801561080f5761363a5750565b5f61232b91611d5b56feafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab603afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60cafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab609afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60bafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60aafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab607afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab608afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab604afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab605a164736f6c634300081e000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// NativeTokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenStakingManagerMetaData.ABI instead.
var NativeTokenStakingManagerABI = NativeTokenStakingManagerMetaData.ABI

// NativeTokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenStakingManagerMetaData.Bin instead.
var NativeTokenStakingManagerBin = NativeTokenStakingManagerMetaData.Bin

// DeployNativeTokenStakingManager deploys a new Ethereum contract, binding an instance of NativeTokenStakingManager to it.
func DeployNativeTokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenStakingManager, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	NativeTokenStakingManagerBin = strings.ReplaceAll(NativeTokenStakingManagerBin, "__$e8433aec8a689eedbe2673465b6d21420c$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NativeTokenStakingManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenStakingManager struct {
	NativeTokenStakingManagerCaller     // Read-only binding to the contract
	NativeTokenStakingManagerTransactor // Write-only binding to the contract
	NativeTokenStakingManagerFilterer   // Log filterer for contract events
}

// NativeTokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenStakingManagerSession struct {
	Contract     *NativeTokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenStakingManagerCallerSession struct {
	Contract *NativeTokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// NativeTokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenStakingManagerTransactorSession struct {
	Contract     *NativeTokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenStakingManagerRaw struct {
	Contract *NativeTokenStakingManager // Generic contract binding to access the raw methods on
}

// NativeTokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCallerRaw struct {
	Contract *NativeTokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactorRaw struct {
	Contract *NativeTokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenStakingManager creates a new instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManager(address common.Address, backend bind.ContractBackend) (*NativeTokenStakingManager, error) {
	contract, err := bindNativeTokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenStakingManagerCaller creates a new read-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenStakingManagerCaller, error) {
	contract, err := bindNativeTokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCaller{contract: contract}, nil
}

// NewNativeTokenStakingManagerTransactor creates a new write-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenStakingManagerTransactor, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerTransactor{contract: contract}, nil
}

// NewNativeTokenStakingManagerFilterer creates a new log filterer instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenStakingManagerFilterer, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerFilterer{contract: contract}, nil
}

// bindNativeTokenStakingManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetDelegatorInfo(opts *bind.CallOpts, delegationID [32]byte) (Delegator, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getDelegatorInfo", delegationID)

	if err != nil {
		return *new(Delegator), err
	}

	out0 := *abi.ConvertType(out[0], new(Delegator)).(*Delegator)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _NativeTokenStakingManager.Contract.GetDelegatorInfo(&_NativeTokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _NativeTokenStakingManager.Contract.GetDelegatorInfo(&_NativeTokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetDelegatorRewardInfo(opts *bind.CallOpts, delegationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getDelegatorRewardInfo", delegationID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenStakingManager.Contract.GetDelegatorRewardInfo(&_NativeTokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenStakingManager.Contract.GetDelegatorRewardInfo(&_NativeTokenStakingManager.CallOpts, delegationID)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetStakingManagerSettings(opts *bind.CallOpts) (StakingManagerSettings, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getStakingManagerSettings")

	if err != nil {
		return *new(StakingManagerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerSettings)).(*StakingManagerSettings)

	return out0, err

}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _NativeTokenStakingManager.Contract.GetStakingManagerSettings(&_NativeTokenStakingManager.CallOpts)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _NativeTokenStakingManager.Contract.GetStakingManagerSettings(&_NativeTokenStakingManager.CallOpts)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetStakingValidator(opts *bind.CallOpts, validationID [32]byte) (PoSValidatorInfo, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getStakingValidator", validationID)

	if err != nil {
		return *new(PoSValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoSValidatorInfo)).(*PoSValidatorInfo)

	return out0, err

}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _NativeTokenStakingManager.Contract.GetStakingValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _NativeTokenStakingManager.Contract.GetStakingValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetValidatorRewardInfo(opts *bind.CallOpts, validationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getValidatorRewardInfo", validationID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenStakingManager.Contract.GetValidatorRewardInfo(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenStakingManager.Contract.GetValidatorRewardInfo(&_NativeTokenStakingManager.CallOpts, validationID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d436317.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings StakingManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initialize", settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d436317.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) Initialize(settings StakingManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d436317.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) Initialize(settings StakingManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x243ca19a.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x243ca19a.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x243ca19a.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateDelegatorRemoval(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xad93936d.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xad93936d.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xad93936d.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, rewardRecipient)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitiateValidatorRemoval(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegator(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegator(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// NativeTokenStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *NativeTokenStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(NativeTokenStakingManagerCompletedDelegatorRegistration)
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
func (it *NativeTokenStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCompletedDelegatorRegistrationIterator{contract: _NativeTokenStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerCompletedDelegatorRegistration)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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

// ParseCompletedDelegatorRegistration is a log parse operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*NativeTokenStakingManagerCompletedDelegatorRegistration, error) {
	event := new(NativeTokenStakingManagerCompletedDelegatorRegistration)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *NativeTokenStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(NativeTokenStakingManagerCompletedDelegatorRemoval)
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
func (it *NativeTokenStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCompletedDelegatorRemovalIterator{contract: _NativeTokenStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerCompletedDelegatorRemoval)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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

// ParseCompletedDelegatorRemoval is a log parse operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*NativeTokenStakingManagerCompletedDelegatorRemoval, error) {
	event := new(NativeTokenStakingManagerCompletedDelegatorRemoval)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRewardClaimedIterator struct {
	Event *NativeTokenStakingManagerDelegatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRewardClaimed)
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
		it.Event = new(NativeTokenStakingManagerDelegatorRewardClaimed)
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
func (it *NativeTokenStakingManagerDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRewardClaimed struct {
	DelegationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address) (*NativeTokenStakingManagerDelegatorRewardClaimedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRewardClaimedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRewardClaimed, delegationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRewardClaimed)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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

// ParseDelegatorRewardClaimed is a log parse operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRewardClaimed(log types.Log) (*NativeTokenStakingManagerDelegatorRewardClaimed, error) {
	event := new(NativeTokenStakingManagerDelegatorRewardClaimed)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator is returned from FilterDelegatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for DelegatorRewardRecipientChanged events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator struct {
	Event *NativeTokenStakingManagerDelegatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRewardRecipientChanged)
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
		it.Event = new(NativeTokenStakingManagerDelegatorRewardRecipientChanged)
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
func (it *NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRewardRecipientChanged represents a DelegatorRewardRecipientChanged event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRewardRecipientChanged struct {
	DelegationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRewardRecipientChanged(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRewardRecipientChangedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRewardRecipientChanged, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRewardRecipientChanged)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
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

// ParseDelegatorRewardRecipientChanged is a log parse operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRewardRecipientChanged(log types.Log) (*NativeTokenStakingManagerDelegatorRewardRecipientChanged, error) {
	event := new(NativeTokenStakingManagerDelegatorRewardRecipientChanged)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitializedIterator struct {
	Event *NativeTokenStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialized)
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
		it.Event = new(NativeTokenStakingManagerInitialized)
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
func (it *NativeTokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialized represents a Initialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenStakingManagerInitialized, error) {
	event := new(NativeTokenStakingManagerInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *NativeTokenStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(NativeTokenStakingManagerInitiatedDelegatorRegistration)
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
func (it *NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedDelegatorRegistration struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	RewardRecipient    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRegistration is a free log retrieval operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedDelegatorRegistrationIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedDelegatorRegistration)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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

// ParseInitiatedDelegatorRegistration is a log parse operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*NativeTokenStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(NativeTokenStakingManagerInitiatedDelegatorRegistration)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *NativeTokenStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(NativeTokenStakingManagerInitiatedDelegatorRemoval)
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
func (it *NativeTokenStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedDelegatorRemovalIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedDelegatorRemoval)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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

// ParseInitiatedDelegatorRemoval is a log parse operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*NativeTokenStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(NativeTokenStakingManagerInitiatedDelegatorRemoval)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator is returned from FilterInitiatedStakingValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedStakingValidatorRegistration events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator struct {
	Event *NativeTokenStakingManagerInitiatedStakingValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitiatedStakingValidatorRegistration)
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
		it.Event = new(NativeTokenStakingManagerInitiatedStakingValidatorRegistration)
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
func (it *NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitiatedStakingValidatorRegistration represents a InitiatedStakingValidatorRegistration event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitiatedStakingValidatorRegistration struct {
	ValidationID      [32]byte
	Owner             common.Address
	DelegationFeeBips uint16
	MinStakeDuration  uint64
	RewardRecipient   common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitiatedStakingValidatorRegistration is a free log retrieval operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitiatedStakingValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, owner []common.Address) (*NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitiatedStakingValidatorRegistrationIterator{contract: _NativeTokenStakingManager.contract, event: "InitiatedStakingValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedStakingValidatorRegistration is a free log subscription operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitiatedStakingValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitiatedStakingValidatorRegistration, validationID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitiatedStakingValidatorRegistration)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
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

// ParseInitiatedStakingValidatorRegistration is a log parse operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitiatedStakingValidatorRegistration(log types.Log) (*NativeTokenStakingManagerInitiatedStakingValidatorRegistration, error) {
	event := new(NativeTokenStakingManagerInitiatedStakingValidatorRegistration)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdatedIterator struct {
	Event *NativeTokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerUptimeUpdated)
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
		it.Event = new(NativeTokenStakingManagerUptimeUpdated)
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
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerUptimeUpdatedIterator{contract: _NativeTokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerUptimeUpdated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*NativeTokenStakingManagerUptimeUpdated, error) {
	event := new(NativeTokenStakingManagerUptimeUpdated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorRewardClaimedIterator is returned from FilterValidatorRewardClaimed and is used to iterate over the raw logs and unpacked data for ValidatorRewardClaimed events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRewardClaimedIterator struct {
	Event *NativeTokenStakingManagerValidatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorRewardClaimed)
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
		it.Event = new(NativeTokenStakingManagerValidatorRewardClaimed)
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
func (it *NativeTokenStakingManagerValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorRewardClaimed represents a ValidatorRewardClaimed event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRewardClaimed struct {
	ValidationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardClaimed is a free log retrieval operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorRewardClaimed(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address) (*NativeTokenStakingManagerValidatorRewardClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorRewardClaimedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardClaimed is a free log subscription operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorRewardClaimed, validationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorRewardClaimed)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
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

// ParseValidatorRewardClaimed is a log parse operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorRewardClaimed(log types.Log) (*NativeTokenStakingManagerValidatorRewardClaimed, error) {
	event := new(NativeTokenStakingManagerValidatorRewardClaimed)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorRewardRecipientChangedIterator is returned from FilterValidatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for ValidatorRewardRecipientChanged events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRewardRecipientChangedIterator struct {
	Event *NativeTokenStakingManagerValidatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorRewardRecipientChanged)
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
		it.Event = new(NativeTokenStakingManagerValidatorRewardRecipientChanged)
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
func (it *NativeTokenStakingManagerValidatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorRewardRecipientChanged represents a ValidatorRewardRecipientChanged event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRewardRecipientChanged struct {
	ValidationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorRewardRecipientChanged(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*NativeTokenStakingManagerValidatorRewardRecipientChangedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorRewardRecipientChangedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorRewardRecipientChanged, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorRewardRecipientChanged)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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

// ParseValidatorRewardRecipientChanged is a log parse operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorRewardRecipientChanged(log types.Log) (*NativeTokenStakingManagerValidatorRewardRecipientChanged, error) {
	event := new(NativeTokenStakingManagerValidatorRewardRecipientChanged)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMessagesMetaData contains all meta data concerning the ValidatorMessages contract.
var ValidatorMessagesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBLSPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"InvalidCodecID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actual\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"}],\"name\":\"packConversionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"packL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"packL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"validationPeriod\",\"type\":\"tuple\"}],\"name\":\"packRegisterL1ValidatorMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conversionID\",\"type\":\"bytes32\"}],\"name\":\"packSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"packValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackL1ValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackRegisterL1ValidatorMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetToL1ConversionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601957611af0908161001e823930815050f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063021de88f146100c4578063088c2463146100bf5780634d847884146100ba57806350782b0f146100b55780637f7c427a146100b0578063854a893f146100ab57806387418b8e146100a65780639b835465146100a1578063a699c1351461009c578063e1d68f30146100975763eb97ce5114610092575f80fd5b611267565b611155565b611104565b610b4f565b6108f5565b610855565b6107cb565b6105fa565b61050f565b6103ae565b610206565b634e487b7160e01b5f52604160045260245ffd5b608081019081106001600160401b038211176100f857604052565b6100c9565b606081019081106001600160401b038211176100f857604052565b604081019081106001600160401b038211176100f857604052565b90601f801991011681019081106001600160401b038211176100f857604052565b60405190610163604083610133565b565b6040519061016360e083610133565b6001600160401b0381116100f857601f01601f191660200190565b81601f820112156101d5578035906101a682610174565b926101b46040519485610133565b828452602083830101116101d557815f926020809301838601378301015290565b5f80fd5b60206003198201126101d557600435906001600160401b0382116101d5576102039160040161018f565b90565b61020f366101d9565b80516027810361039157505f5f5b6002811061034b575061ffff811661033557505f5f5b600481106102e3575063ffffffff60029116036102d4575f905f5b6020811061028f575061027661026860ff60f81b92611390565b516001600160f81b03191690565b60408051938452911615156020830152819081015b0390f35b916001906102ba6102b46102ae6102686102a88861142d565b876113a5565b60f81c90565b60ff1690565b6102cb6102c6866113eb565b611407565b1b17920161024e565b635b60892f60e01b5f5260045ffd5b9060019061032d6103056102b46102ae6102686102ff8861141f565b896113a5565b61031d6103146102c6876113dd565b63ffffffff1690565b63ffffffff8080931691161b1690565b179101610233565b63407b587360e01b5f5261ffff1660045260245ffd5b906001906103896103656102b46102ae61026887896113a5565b61037b6103746102c6876113ca565b61ffff1690565b61ffff8080931691161b1690565b17910161021d565b63cc92daa160e01b5f5263ffffffff16600452602760245260445ffd5b6103b7366101d9565b8051602e81036104f257505f5f5b600281106104d0575061ffff811661033557505f805b600482106104ac5763ffffffff9150166102d4575f5f5b6020811061047e57505f915f905b60088210610424575050604080519182526001600160401b03929092166020820152f35b90926001906104756104476102b46102ae6102686104418a61143b565b886113a5565b6104626104566102c6896113f9565b6001600160401b031690565b6001600160401b038080931691161b1690565b17930190610400565b906001906104976102b46102ae6102686104418761142d565b6104a36102c6856113eb565b1b1791016103f2565b6001906104c76103056102b46102ae6102686102ff8861141f565b179101906103db565b906001906104ea6103656102b46102ae61026887896113a5565b1791016103c5565b63cc92daa160e01b5f5263ffffffff16600452602e60245260445ffd5b610518366101d9565b8051602681036105dd57505f5f5b600281106105bb575061ffff811661033557505f805b600482106105975763ffffffff9150166102d4575f5f915b6020831061056757604051828152602090f35b906001906105806102b46102ae6102686102a88861142d565b61058c6102c6866113eb565b1b1792019190610554565b6001906105b26103056102b46102ae6102686102ff8861141f565b1791019061053c565b906001906105d56103656102b46102ae61026887896113a5565b179101610526565b63cc92daa160e01b5f5263ffffffff16600452602660245260445ffd5b610603366101d9565b80516036810361075757505f5f5b60028110610735575061ffff811661033557505f5f5b60048110610711575063ffffffff60039116036102d4575f5f5b602081106106e357505f915f5b600881106106bf57505f5f915b6008831061068a5750604080519384526001600160401b03948516602085015293169282019290925260609150f35b906001906106b56106a66102b46102ae61026861044189611449565b6104626104566102c6886113f9565b179201919061065b565b926001906106db6104476102b46102ae6102686104418a61143b565b17930161064e565b906001906106fc6102b46102ae6102686104418761142d565b6107086102c6856113eb565b1b179101610641565b9060019061072d6103056102b46102ae6102686102ff8861141f565b179101610627565b9060019061074f6103656102b46102ae61026887896113a5565b179101610611565b63cc92daa160e01b5f5263ffffffff16600452603660245260445ffd5b5f5b8381106107855750505f910152565b8181015183820152602001610776565b906020916107ae81518092818552858086019101610774565b601f01601f1916010190565b906020610203928181520190610795565b60203660031901126101d55761028b600435604051905f60208301525f6022830152602682015260268152610801604682610133565b604051918291602083526020830190610795565b602435906001600160401b03821682036101d557565b604435906001600160401b03821682036101d557565b35906001600160401b03821682036101d557565b60603660031901126101d55761028b60043561086f610815565b61087761082b565b6040515f6020820152600360e01b602282015260268101939093526001600160c01b031960c092831b81166046850152911b16604e820152603681526108be605682610133565b604051918291826107ba565b35906001600160a01b03821682036101d557565b6001600160401b0381116100f85760051b60200190565b60203660031901126101d5576004356001600160401b0381116101d557608060031982360301126101d5576040519061092d826100dd565b806004013582526024810135602083015261094a604482016108ca565b60408301526064810135906001600160401b0382116101d55701366023820112156101d557600481013561097d816108de565b9161098b6040519384610133565b818352602060048185019360051b83010101903682116101d55760248101925b8284106109c45761028b6108be878760608201526114fb565b83356001600160401b0381116101d557600490830101906060601f1983360301126101d557604051906109f6826100fd565b60208301356001600160401b0381116101d557610a19906020369186010161018f565b82526040830135916001600160401b0383116101d557610a516060602095610a47879687369184010161018f565b8685015201610841565b60408201528152019301926109ab565b6020606081604085019363ffffffff81511686520151936040838201528451809452019201905f5b818110610a965750505090565b82516001600160a01b0316845260209384019390920191600101610a89565b61020390602081528251602082015260e060c0610b3e610b29610b00610aea6020890151866040890152610100880190610795565b6040890151878203601f19016060890152610795565b6001600160401b0360608901511660808701526080880151601f198783030160a0880152610a61565b60a0870151858203601f190184870152610a61565b9401516001600160401b0316910152565b610b58366101d9565b610b6061163d565b50610b6961163d565b5f5f5b600281106110e2575061ffff811661033557505f8060025b600482106110ac57505063ffffffff60019116036102d457610ba66002611693565b5f5f63ffffffff83165b6020821061107b5750508252610bc5906116ab565b5f905f63ffffffff82165b6004821061103f575050610be390611693565b8163ffffffff8116610bf481611780565b9063ffffffff84165f5b82811061100a575050506020850152610c1691611722565b610c1e61173c565b63ffffffff82165f5b60308110610fde5750506040840152610c3f906116c3565b5f5f63ffffffff83165b60088210610fb15750506001600160401b03166060840152610c6a906116db565b5f9182918263ffffffff82165b8760048310610f8957505050610c8c90611693565b5f63ffffffff82165b60048210610f61575050610ca890611693565b9263ffffffff8316610cb9816117b2565b905f905b808210610eff575050610cdd610cd1610154565b63ffffffff9093168352565b602082015260808501525f905f935f5b60048110610ec45750610cff90611693565b5f63ffffffff82165b8860048310610e8d57505050610d1d90611693565b9363ffffffff8316610d2e816117b2565b905f905b808210610e0f57505092610d66610d7193610d6c93610d779796610d57610cd1610154565b602082015260a08a015261170b565b93611722565b6117da565b90611722565b835163ffffffff82168103610df1575050915f925f915b60088310610db3576001600160401b03851660c08501526040518061028b8682610ab5565b909193600190610de7610dd86102b46102ae61026861044163ffffffff8a168c611457565b6104626104566102c68a6113f9565b1794019190610d8e565b63cc92daa160e01b5f5263ffffffff9081166004521660245260445ffd5b9096610e1961175e565b5f5b8b60148210610e5a575050600191610e4d6014610e52930151610e3e8c88611464565b6001600160a01b039091169052565b6116f3565b970190610d32565b90610e7a610268600193610e7463ffffffff881685611457565b906113a5565b5f1a610e8682856113a5565b5301610e1b565b94610ebb610eac6102b46102ae610268879a610e74886001999a611457565b61031d6103146102c68a6113dd565b17940190610d08565b94600190610ef7610ee86102b46102ae6102688d610e7463ffffffff8a168e611457565b61031d6103146102c68b6113dd565b179501610ced565b9095610f0961175e565b5f5b60148110610f345750600191610e4d6014610f2c930151610e3e8b88611464565b960190610cbd565b600190610f4e6102688d610e7463ffffffff881685611457565b5f1a610f5a82856113a5565b5301610f0b565b9093600190610f80610eac6102b46102ae6102688d610e74898d611457565b17940190610c95565b95610fa8610ee86102b46102ae610268879b610e74886001999a611457565b17950190610c77565b9091600190610fd56106a66102b46102ae610268610fcf888a611457565b8d6113a5565b17920190610c49565b80610ff7610268610ff185600195611457565b8a6113a5565b5f1a61100382866113a5565b5301610c27565b819293945061102761026861102160019484611457565b8b6113a5565b5f1a61103382876113a5565b53019085939291610bfe565b90926001906110726110636102b46102ae61026861105d888b611457565b8c6113a5565b61031d6103146102c6896113dd565b17930190610bd0565b90916001906110966102b46102ae6102686110218789611457565b6110a26102c6866113eb565b1b17920190610bb0565b90916001906110d96110ca6102b46102ae610268611021888a611457565b61031d6103146102c6886113dd565b17920190610b84565b906001906110fc6103656102b46102ae610268878a6113a5565b179101610b6c565b60403660031901126101d557600435602435908115158092036101d55761028b91604051915f6020840152600160e11b6022840152602683015260f81b6046820152602781526108be604782610133565b60403660031901126101d55761028b60043561116f610815565b604051915f60208401525f602284015260268301526001600160401b0360c01b9060c01b166046820152602e8152610801604e82610133565b91906040838203126101d5576040516111c081610118565b8093803563ffffffff811681036101d55782526020810135906001600160401b0382116101d557019180601f840112156101d55782356111ff816108de565b9361120d6040519586610133565b81855260208086019260051b8201019283116101d557602001905b8282106112385750505060200152565b60208091611245846108ca565b815201910190611228565b604090610203939281528160208201520190610795565b60203660031901126101d5576004356001600160401b0381116101d55760e060031982360301126101d55761129a610165565b8160040135815260248201356001600160401b0381116101d5576112c4906004369185010161018f565b602082015260448201356001600160401b0381116101d5576112ec906004369185010161018f565b60408201526112fd60648301610841565b606082015260848201356001600160401b0381116101d55761132590600436918501016111a8565b608082015260a4820135916001600160401b0383116101d55760c46113629161135761136c95600436918401016111a8565b60a085015201610841565b60c082015261192d565b9061028b60405192839283611250565b634e487b7160e01b5f52603260045260245ffd5b8051602610156113a05760460190565b61137c565b9081518110156113a0570160200190565b634e487b7160e01b5f52601160045260245ffd5b60010390600182116113d857565b6113b6565b60030390600382116113d857565b601f0390601f82116113d857565b60070390600782116113d857565b600381901b91906001600160fd1b038116036113d857565b90600282018092116113d857565b90600682018092116113d857565b90602682018092116113d857565b90602e82018092116113d857565b919082018092116113d857565b80518210156113a05760209160051b010190565b9061148b60209282815194859201610774565b0190565b91602060089695936114aa6004969482815194859201610774565b019063ffffffff60e01b9060e01b1681526114ce8251809360208785019101610774565b016114e28251809360208685019101610774565b0101906001600160401b0360c01b9060c01b1681520190565b90815161159860208401519161158a606061151f604088015160018060a01b031690565b96019561153187515163ffffffff1690565b6040515f602082015260228101949094526042840195909552600560e21b606284015260601b6bffffffffffffffffffffffff1916606683015260e09390931b6001600160e01b031916607a820152918290607e820190565b03601f198101835282610133565b915f925b8151805185101561161e57906116166115c66115ba87600195611464565b51515163ffffffff1690565b9161158a6115d5888751611464565b515160206115e48a8951611464565b51015161160660406115f78c8b51611464565b5101516001600160401b031690565b916040519687956020870161148f565b93019261159c565b50925050565b6040519061163182610118565b60606020835f81520152565b6040519060e082018281106001600160401b038211176100f8576040525f60c083828152606060208201526060604082015282606082015261167d611624565b608082015261168a611624565b60a08201520152565b63ffffffff60049116019063ffffffff82116113d857565b63ffffffff60209116019063ffffffff82116113d857565b63ffffffff60309116019063ffffffff82116113d857565b63ffffffff60089116019063ffffffff82116113d857565b63ffffffff60149116019063ffffffff82116113d857565b63ffffffff16607a019063ffffffff82116113d857565b9063ffffffff8091169116019063ffffffff82116113d857565b6040516060919061174d8382610133565b6030815291601f1901366020840137565b6040805190919061176f8382610133565b6014815291601f1901366020840137565b9061178a82610174565b6117976040519182610133565b82815280926117a8601f1991610174565b0190602036910137565b906117bc826108de565b6117c96040519182610133565b82815280926117a8601f19916108de565b63ffffffff60149116029063ffffffff82169182036113d857565b9261148b959693602a9360109996935f8352600160e01b6002840152600683015263ffffffff60e01b9060e01b16602682015261183b8251809360208785019101610774565b0161184f8251809360208685019101610774565b60c09690961b6001600160c01b03191695010193845260e090811b6001600160e01b0319908116600886015291901b16600c830152565b60209061189c6014949382815194859201610774565b01906bffffffffffffffffffffffff199060601b1681520190565b6020906118cd6008959382815194859201610774565b019163ffffffff60e01b9060e01b16825263ffffffff60e01b9060e01b1660048201520190565b60209061190a6008949382815194859201610774565b01906001600160401b0360c01b9060c01b1681520190565b6040513d5f823e3d90fd5b6040810190603082515103611ad4578051916119a960208301519361158a611959865163ffffffff1690565b93519561197060608701516001600160401b031690565b966080870197885191611999602061198c855163ffffffff1690565b9401515163ffffffff1690565b93604051988997602089016117f5565b905f915b60208451015180518410156119fa576020916119f06119de6119d187600195611464565b516001600160a01b031690565b9161158a604051938492878401611886565b93019290506119ad565b50925092905060a0830191611a3d83519161158a611a2e6020611a21865163ffffffff1690565b9501515163ffffffff1690565b604051948593602085016118b7565b905f915b6020845101518051841015611a6f57602091611a656119de6119d187600195611464565b9301929050611a41565b5093611a8c919350611aad925060c001516001600160401b031690565b92611a9f604051948592602084016118f4565b03601f198101845283610133565b60205f60405180611abe8187611478565b039060025afa15611acf575f519190565b611922565b63180ffa0d60e01b5f5260045ffdfea164736f6c634300081e000a",
}

// ValidatorMessagesABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMessagesMetaData.ABI instead.
var ValidatorMessagesABI = ValidatorMessagesMetaData.ABI

// ValidatorMessagesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMessagesMetaData.Bin instead.
var ValidatorMessagesBin = ValidatorMessagesMetaData.Bin

// DeployValidatorMessages deploys a new Ethereum contract, binding an instance of ValidatorMessages to it.
func DeployValidatorMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorMessages, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorMessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// ValidatorMessages is an auto generated Go binding around an Ethereum contract.
type ValidatorMessages struct {
	ValidatorMessagesCaller     // Read-only binding to the contract
	ValidatorMessagesTransactor // Write-only binding to the contract
	ValidatorMessagesFilterer   // Log filterer for contract events
}

// ValidatorMessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorMessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMessagesSession struct {
	Contract     *ValidatorMessages // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorMessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMessagesCallerSession struct {
	Contract *ValidatorMessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorMessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMessagesTransactorSession struct {
	Contract     *ValidatorMessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorMessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMessagesRaw struct {
	Contract *ValidatorMessages // Generic contract binding to access the raw methods on
}

// ValidatorMessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMessagesCallerRaw struct {
	Contract *ValidatorMessagesCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactorRaw struct {
	Contract *ValidatorMessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMessages creates a new instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessages(address common.Address, backend bind.ContractBackend) (*ValidatorMessages, error) {
	contract, err := bindValidatorMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// NewValidatorMessagesCaller creates a new read-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMessagesCaller, error) {
	contract, err := bindValidatorMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesCaller{contract: contract}, nil
}

// NewValidatorMessagesTransactor creates a new write-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMessagesTransactor, error) {
	contract, err := bindValidatorMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesTransactor{contract: contract}, nil
}

// NewValidatorMessagesFilterer creates a new log filterer instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMessagesFilterer, error) {
	contract, err := bindValidatorMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesFilterer{contract: contract}, nil
}

// bindValidatorMessages binds a generic wrapper to an already deployed contract.
func bindValidatorMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.ValidatorMessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transact(opts, method, params...)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackConversionData(opts *bind.CallOpts, conversionData ConversionData) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packConversionData", conversionData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackConversionData is a free data retrieval call binding the contract method 0x51f48008.
//
// Solidity: function packConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackConversionData(conversionData ConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackConversionData(&_ValidatorMessages.CallOpts, conversionData)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool registered) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorRegistrationMessage(opts *bind.CallOpts, validationID [32]byte, registered bool) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorRegistrationMessage", validationID, registered)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool registered) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, registered bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, registered)
}

// PackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xa699c135.
//
// Solidity: function packL1ValidatorRegistrationMessage(bytes32 validationID, bool registered) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorRegistrationMessage(validationID [32]byte, registered bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, registered)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackL1ValidatorWeightMessage(opts *bind.CallOpts, validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packL1ValidatorWeightMessage", validationID, nonce, weight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x854a893f.
//
// Solidity: function packL1ValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackL1ValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackRegisterL1ValidatorMessage(opts *bind.CallOpts, validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packRegisterL1ValidatorMessage", validationPeriod)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0xe0d5478f.
//
// Solidity: function packRegisterL1ValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackRegisterL1ValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetToL1ConversionMessage(opts *bind.CallOpts, conversionID [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetToL1ConversionMessage", conversionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x7f7c427a.
//
// Solidity: function packSubnetToL1ConversionMessage(bytes32 conversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetToL1ConversionMessage(conversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, conversionID)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackValidationUptimeMessage(opts *bind.CallOpts, validationID [32]byte, uptime uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packValidationUptimeMessage", validationID, uptime)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorRegistrationMessage(opts *bind.CallOpts, input []byte) ([32]byte, bool, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorRegistrationMessage", input)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x021de88f.
//
// Solidity: function unpackL1ValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackL1ValidatorWeightMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackL1ValidatorWeightMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackL1ValidatorWeightMessage is a free data retrieval call binding the contract method 0x50782b0f.
//
// Solidity: function unpackL1ValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackL1ValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackL1ValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackRegisterL1ValidatorMessage(opts *bind.CallOpts, input []byte) (ValidatorMessagesValidationPeriod, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackRegisterL1ValidatorMessage", input)

	if err != nil {
		return *new(ValidatorMessagesValidationPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorMessagesValidationPeriod)).(*ValidatorMessagesValidationPeriod)

	return out0, err

}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterL1ValidatorMessage is a free data retrieval call binding the contract method 0x9b835465.
//
// Solidity: function unpackRegisterL1ValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackRegisterL1ValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterL1ValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetToL1ConversionMessage(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetToL1ConversionMessage", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetToL1ConversionMessage is a free data retrieval call binding the contract method 0x4d847884.
//
// Solidity: function unpackSubnetToL1ConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetToL1ConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetToL1ConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackValidationUptimeMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackValidationUptimeMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}
