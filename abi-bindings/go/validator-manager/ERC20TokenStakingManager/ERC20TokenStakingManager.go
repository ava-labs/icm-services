// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenstakingmanager

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

// ERC20TokenStakingManagerMetaData contains all meta data concerning the ERC20TokenStakingManager contract.
var ERC20TokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InvalidRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidUptimeBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedValidationID\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWeightToValueFactor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"CompletedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"}],\"name\":\"DelegatorRewardRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InitiatedDelegatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InitiatedDelegatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"InitiatedStakingValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"}],\"name\":\"ValidatorRewardRecipientChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"contractIERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startingNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endingNonce\",\"type\":\"uint64\"}],\"internalType\":\"structDelegator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"getDelegatorRewardInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingManagerSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structStakingManagerSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getStakingValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uptimeSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structPoSValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidatorRewardInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"weightToValueFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\"}],\"internalType\":\"structStakingManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateDelegatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100fd57601f613bb338819003918201601f19168301916001600160401b03831184841017610101578084926020946040528339810103126100fd575160028110156100fd5760011461005f575b604051613a7d90816101168239f35b5f516020613b935f395f51905f525460ff8160401c166100ee576002600160401b03196001600160401b03821601610098575b50610050565b6001600160401b0319166001600160401b039081175f516020613b935f395f51905f52556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f610092565b63f92ee8a960e01b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80631340964514611aea578063151d30d114611acf5780631667956414611ab8578063245dafcb1461183857806325e1c7761461176a5780632674874b146116b357806327bf60cd1461169c5780632aa56638146116695780632e2194d81461163a57806335455ded14610dc457806353a133381461152257806360ad7784146114fe57806362065856146114b6578063785e9e86146114825780637a63ad851461145b5780637e65f4da146110975780638ef34c9814610fe757806393e2459814610eb25780639681d94014610e6e578063a3a65e4814610dc9578063a9778a7a14610dc4578063af1dd66c14610d65578063b2c1712e14610d32578063b5c9349814610880578063b771b3bc1461085e578063caa71874146106af578063cd236e7f146102c5578063e24b268014610266578063e4a63c401461023f5763fb8b11dd14610163575f80fd5b3461023b57604036600319011261023b576004356001600160a01b03610187611ecd565b168015610229575f8281525f5160206139b15f395f51905f52602052604090205460081c6001600160a01b03163303610216575f8281525f5160206139515f395f51905f526020526040812080546001600160a01b0319811684179091556001600160a01b031692907f6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e0059080a4005b636e2ccd7560e11b5f523360045260245ffd5b63caa903f960e01b5f5260045260245ffd5b5f80fd5b3461023b575f36600319011261023b5760206040515f5160206139715f395f51905f528152f35b3461023b57602036600319011261023b576004355f9081525f5160206138f15f395f51905f5260209081526040808320545f5160206139315f395f51905f528352928190205481516001600160a01b0390941684529183019190915290f35b3461023b57606036600319011261023b576004356044356001600160a01b0381169081900361023b576102f66125ee565b5f5160206139715f395f51905f52546103249061031f90602435906001600160a01b0316613707565b612153565b9161032e81612c51565b1561069d57811561068a575f5160206139f15f395f51905f5254604051636af907fb60e11b81526004810183905292906001600160a01b03165f84602481845afa938415610634575f94610666575b506001600160401b036040610398878360a089015116612626565b950151166001600160401b035f5160206138d15f395f51905f525460501c1602936001600160401b038516948503610652576001600160401b03811694851161063f5760408051636610966960e01b8152600481018690526001600160401b0392909216602483015290939291849060449082905f905af18015610634576020955f945f926105f6575b506001600160401b0390604051888101908682528360c01b8860c01b16604082015260288152610453604882611e4b565b51902096875f525f5160206139b15f395f51905f52895260405f20600160ff19825416179055875f525f5160206139b15f395f51905f52895260405f208054610100600160a81b033360081b1690610100600160a81b031916179055875f525f5160206139b15f395f51905f52895285600160405f200155875f525f5160206139b15f395f51905f528952600260405f20018383168419825416179055875f525f5160206139b15f395f51905f528952600260405f200167ffffffffffffffff60401b198154169055875f525f5160206139b15f395f51905f528952600260405f200180548460801b8960801b16908560801b1916179055875f525f5160206139b15f395f51905f528952600260405f200160018060c01b038154169055875f525f5160206139515f395f51905f52895260405f20856001600160601b0360a01b82541617905582604051971687528887015216604085015260608401526080830152827f77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d7760a03394a460015f516020613a115f395f51905f5255604051908152f35b6001600160401b03955061062391925060403d60401161062d575b61061b8183611e4b565b810190612f34565b9490949190610422565b503d610611565b6040513d5f823e3d90fd5b84636d51fe0560e11b5f5260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b6106839194503d805f833e61067b8183611e4b565b810190612026565b928561037d565b5063caa903f960e01b5f5260045260245ffd5b6330efa98b60e01b5f5260045260245ffd5b3461023b575f36600319011261023b575f6101006040516106cf81611e2f565b8281528260208201528260408201528260608201528260808201528260a08201528260c08201528260e0820152015261012060018060a01b035f5160206139f15f395f51905f5254167fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab601547fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab602545f5160206138d15f395f51905f525460ff5f5160206139d15f395f51905f52549161ffff60018060a01b035f516020613a515f395f51905f525416946001600160401b037fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6065497604051926107d084611e2f565b8a84526020840191825260408401908152606084019183871683526101006080860195878960401c1687528960a082019960501c16895260c081019a8b5260e081019b8c52019a8b526040519b8c525160208c01525160408b01525116606089015251166080870152511660a08501525160c084015260018060a01b0390511660e083015251610100820152f35b3461023b575f36600319011261023b576040516005600160991b018152602090f35b3461023b573660031901610140811261023b576101201361023b57610124356001600160a01b0381169081900361023b575f516020613a315f395f51905f52549060ff8260401c1615916001600160401b03811680159081610d2a575b6001149081610d20575b159081610d17575b50610d085767ffffffffffffffff1981166001175f516020613a315f395f51905f525582610cdc575b506109216136dc565b6109296136dc565b6109316136dc565b6109396136dc565b60015f516020613a115f395f51905f52556004356001600160a01b0381169081900361023b5760243560443591606435926001600160401b03841680940361023b576084359061ffff82169485830361023b5760a4359160ff83169586840361023b5760e4356001600160a01b0381169760c435979189900361023b5761010435996109c36136dc565b80158015610cd1575b610cbf5750838311610cac5780158015610ca2575b610c9057508015610c0d578715610c0d576040516304e0efb360e11b8152602081600481855afa8015610634575f90610c50575b6001600160401b039150168410610c3e578615610c2f578815610c1c576001600160601b0360a01b5f5160206139f15f395f51905f525416175f5160206139f15f395f51905f52557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab601557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab602556001600160401b03195f5160206138d15f395f51905f525416175f5160206138d15f395f51905f525569ffff00000000000000005f5160206138d15f395f51905f52549160ff60501b9060501b169260401b169071ffffffffffffffffffff0000000000000000191617175f5160206138d15f395f51905f52555f5160206139d15f395f51905f52556001600160601b0360a01b5f516020613a515f395f51905f525416175f516020613a515f395f51905f52557fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60655610b7f6136dc565b8015610c0d576001600160601b0360a01b5f5160206139715f395f51905f525416175f5160206139715f395f51905f5255610bb657005b68ff0000000000000000195f516020613a315f395f51905f5254165f516020613a315f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b63d92e233d60e01b5f5260045ffd5b88632f6bd1db60e01b5f5260045260245ffd5b63a733007160e01b5f5260045ffd5b836202a06d60e11b5f5260045260245ffd5b506020813d602011610c88575b81610c6a60209383611e4b565b8101031261023b57610c836001600160401b0391612012565b610a15565b3d9150610c5d565b63170db35960e31b5f5260045260245ffd5b50600a81116109e1565b8263222d164360e21b5f5260045260245ffd5b635f12e6c360e11b5f5260045260245ffd5b5061271081116109cc565b68ffffffffffffffffff191668010000000000000001175f516020613a315f395f51905f525582610918565b63f92ee8a960e01b5f5260045ffd5b905015846108ef565b303b1591506108e7565b8491506108dd565b3461023b57610d4c610d4336611d8b565b90829392612971565b15610d5357005b635bff683f60e11b5f5260045260245ffd5b3461023b57602036600319011261023b576004355f9081525f5160206139515f395f51905f5260209081526040808320545f5160206139115f395f51905f528352928190205481516001600160a01b0390941684529183019190915290f35b611dbb565b3461023b57602036600319011261023b57610de2611d78565b5f5160206139f15f395f51905f5254604051631474cbc960e31b815263ffffffff929092166004830152602090829060249082905f906001600160a01b03165af18015610634575f90610e3b575b602090604051908152f35b506020813d602011610e66575b81610e5560209383611e4b565b8101031261023b5760209051610e30565b3d9150610e48565b3461023b57602036600319011261023b576020610e99610e8c611d78565b610e946125ee565b612450565b60015f516020613a115f395f51905f5255604051908152f35b3461023b57602036600319011261023b575f5160206139f15f395f51905f5254604051636af907fb60e11b815260048035908201819052915f90829060249082906001600160a01b03165afa908115610634575f91610fcd575b50516006811015610fb95760048103610fa757505f8181525f5160206139915f395f51905f5260205260409020546001600160a01b03163303610216575f8181525f5160206138f15f395f51905f526020526040902054610f7c91906001600160a01b03168015610f7e5761335a565b005b505f8181525f5160206139915f395f51905f5260205260409020546001600160a01b031661335a565b63170cc93360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b610fe191503d805f833e61067b8183611e4b565b82610f0c565b3461023b57604036600319011261023b576004356001600160a01b0361100b611ecd565b168015610229575f8281525f5160206139915f395f51905f5260205260409020546001600160a01b03163303610216575f8281525f5160206138f15f395f51905f526020526040812080546001600160a01b0319811684179091556001600160a01b031692907f28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c339080a4005b3461023b5761010036600319011261023b576004356001600160401b03811161023b576110c8903690600401611e87565b6024356001600160401b03811161023b576110e7903690600401611e87565b6044356001600160401b03811161023b57611106903690600401611ee3565b916064356001600160401b03811161023b57611126903690600401611ee3565b906084359261ffff84169283850361023b5760a435906001600160401b0382169283830361023b5760e4356001600160a01b038116959060c4359087900361023b576111706125ee565b5f5160206138d15f395f51905f525461ffff8160401c1689108015611450575b61143d576001600160401b0316861061142b577fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6015481108015611401575b6113ef5786156113dc575f5160206139715f395f51905f52546111fa91906001600160a01b0316613707565b61120390612153565b5f5160206139f15f395f51905f5254604051634e5bb12760e11b815260a060048201529a6001600160a01b03909116948b9485949192916112489060a487019061212e565b85810360031901602487015261125d9161212e565b84810360031901604486015261127291613306565b83810360031901606485015261128791613306565b906001600160401b0316608483015203815a6020945f91f1948515610634575f956113a7575b5f8681525f5160206139915f395f51905f5260209081526040808320805469ffffffffffffffffffff60a01b19339081166001600160f01b03199092169190911760a09690961b61ffff60a01b169590951760b09690961b67ffffffffffffffff60b01b169590951785556001909401805467ffffffffffffffff191690555f5160206138f15f395f51905f5281529083902080546001600160a01b031916861790558251958652858101939093529084019290925293509082907ff51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f90606090a360015f516020613a115f395f51905f5255604051908152f35b94506020863d6020116113d4575b816113c260209383611e4b565b8101031261023b5760209551946112ad565b3d91506113b5565b8663caa903f960e01b5f5260045260245ffd5b63222d164360e21b5f5260045260245ffd5b507fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6025481116111ce565b856202a06d60e11b5f5260045260245ffd5b88635f12e6c360e11b5f5260045260245ffd5b506127108911611190565b3461023b575f36600319011261023b5760206040515f5160206139f15f395f51905f528152f35b3461023b575f36600319011261023b575f5160206139715f395f51905f52546040516001600160a01b039091168152602090f35b3461023b57602036600319011261023b576004356001600160401b03811680910361023b576114f66020915f5160206139d15f395f51905f52549061243d565b604051908152f35b3461023b57604036600319011261023b57610f7c61151a611d65565b6004356121a9565b3461023b57602036600319011261023b575f60c060405161154281611de5565b8281528260208201528260408201528260608201528260808201528260a082015201526004355f525f5160206139b15f395f51905f5260205260405f206040519061158c82611de5565b80549160ff8316916004831015610fb9576001600160401b03828160e09686829652602083019060018060a01b039060081c16815281600260018701549660408601978852015495606085019082881682526080860193838960401c16855260c060a0880197858b60801c168952019860c01c89526040519a8b5260018060a01b0390511660208b01525160408a01525116606088015251166080860152511660a0840152511660c0820152f35b3461023b57602036600319011261023b576020611658600435612153565b6001600160401b0360405191168152f35b3461023b5761168361167a36611d8b565b90829392612f51565b1561168a57005b631036cf9160e11b5f5260045260245ffd5b3461023b57610f7c6116ad36611d8b565b91612f51565b3461023b57602036600319011261023b575f60606040516116d381611e14565b82815282602082015282604082015201526004355f525f5160206139915f395f51905f52602052608060405f206001600160401b0360405161171481611e14565b81835461ffff60018060a01b0382169586855260608460016020880193858760a01c1685528260408a019760b01c1687520154169501948552604051968752511660208601525116604084015251166060820152f35b3461023b57604036600319011261023b57600435611786611d65565b61178f82612c51565b15611825575f5160206139f15f395f51905f5254604051636af907fb60e11b815260048101849052905f90829060249082906001600160a01b03165afa908115610634575f9161180b575b5051916006831015610fb957600283036117f857610f7c9250612c79565b8263170cc93360e21b5f5260045260245ffd5b61181f91503d805f833e61067b8183611e4b565b836117da565b506330efa98b60e01b5f5260045260245ffd5b3461023b57602036600319011261023b57600435805f525f5160206139b15f395f51905f5260205260405f2060405161187081611de5565b815460ff81166004811015610fb957600191818452828060a01b039060081c166020840152600282850154946040850195865201546001600160401b03811660608501526001600160401b038160401c1660808501526001600160401b038160801c1660a085015260c01c60c0840152141580611aa3575b611a81575060018060a01b035f5160206139f15f395f51905f525416905f8151602460405180958193636af907fb60e11b835260048301525afa918215610634575f92611a65575b5060608201926001600160401b0384511615611a53575060a06001600160401b0380925194511692015116906040519263854a893f60e01b84526004840152602483015260448201525f8160648173__$e8433aec8a689eedbe2673465b6d21420c$__5af4908115610634575f91611a0c575b5060206119cb916040518093819263ee5b48eb60e01b8352846004840152602483019061212e565b03815f6005600160991b015af18015610634576119e457005b610f7c9060203d602011611a05575b6119fd8183611e4b565b81019061211f565b503d6119f3565b90503d805f833e611a1d8183611e4b565b81019060208183031261023b578051906001600160401b03821161023b576119cb92602092611a4c9201611fcd565b91506119a3565b6339b894f960e21b5f5260045260245ffd5b611a7a9192503d805f833e61067b8183611e4b565b9083611930565b516004811015610fb957633b0d540d60e21b5f52611a9e90611dd7565b60245ffd5b5080516004811015610fb957600314156118e8565b3461023b57610f7c611ac936611d8b565b91612971565b3461023b575f36600319011261023b576020604051600a8152f35b3461023b57604036600319011261023b57600435611b06611d65565b90611b0f6125ee565b805f525f5160206139b15f395f51905f5260205260405f209160405190611b3582611de5565b83549160ff8316926004841015610fb95783825260018060a01b039060081c16602082015260036002600187015496604084019788520154936001600160401b03851660608401526001600160401b038560401c1660808401526001600160401b038560801c1660a084015260c083019460c01c855203611a81575060018060a01b035f5160206139f15f395f51905f52541690845160405190636af907fb60e11b825260048201525f81602481865afa908115610634575f91611d4b575b50855160405190636af907fb60e11b825260048201525f81602481875afa908115610634575f91611d31575b50516006811015610fb957600414159081611d10575b50611c58575b611c4584612646565b60015f516020613a115f395f51905f5255005b63ffffffff60246040925f8451958694859363338587c560e21b85521660048401525af1938415610634575f915f95611cdd575b505190808203611cc85750506001600160401b03809151169216809211611cb557818080611c3c565b50632e19bc2d60e11b5f5260045260245ffd5b63fee3144560e01b5f5260045260245260445ffd5b909450611d02915060403d604011611d09575b611cfa8183611e4b565b810190612105565b9385611c8c565b503d611cf0565b6001600160401b03915060800151166001600160401b038451161186611c36565b611d4591503d805f833e61067b8183611e4b565b87611c20565b611d5f91503d805f833e61067b8183611e4b565b86611bf4565b6024359063ffffffff8216820361023b57565b6004359063ffffffff8216820361023b57565b606090600319011261023b5760043590602435801515810361023b579060443563ffffffff8116810361023b5790565b3461023b575f36600319011261023b5760206040516127108152f35b6004811015610fb957600452565b60e081019081106001600160401b03821117611e0057604052565b634e487b7160e01b5f52604160045260245ffd5b608081019081106001600160401b03821117611e0057604052565b61012081019081106001600160401b03821117611e0057604052565b90601f801991011681019081106001600160401b03821117611e0057604052565b6001600160401b038111611e0057601f01601f191660200190565b81601f8201121561023b57803590611e9e82611e6c565b92611eac6040519485611e4b565b8284526020838301011161023b57815f926020809301838601378301015290565b602435906001600160a01b038216820361023b57565b919060408382031261023b57604051604081018181106001600160401b03821117611e00576040528093803563ffffffff8116810361023b5782526020810135906001600160401b03821161023b570182601f8201121561023b578035926001600160401b038411611e00578360051b9160405194611f656020850187611e4b565b855260208086019382010191821161023b57602001915b818310611f8c5750505060200152565b82356001600160a01b038116810361023b57815260209283019201611f7c565b5f5b838110611fbd5750505f910152565b8181015183820152602001611fae565b81601f8201121561023b578051611fe381611e6c565b92611ff16040519485611e4b565b8184526020828401011161023b5761200f9160208085019101611fac565b90565b51906001600160401b038216820361023b57565b60208183031261023b578051906001600160401b03821161023b57016101008183031261023b576040519161010083018381106001600160401b03821117611e00576040528151600681101561023b5783526020820151916001600160401b03831161023b5761209d60e0926120fd948301611fcd565b60208501526120ae60408201612012565b60408501526120bf60608201612012565b60608501526120d060808201612012565b60808501526120e160a08201612012565b60a08501526120f260c08201612012565b60c085015201612012565b60e082015290565b919082604091031261023b5761200f602083519301612012565b9081602091031261023b575190565b9060209161214781518092818552858086019101611fac565b601f01601f1916010190565b5f5160206139d15f395f51905f525480156121955781049081158015612185575b6113ef57506001600160401b031690565b506001600160401b038211612174565b634e487b7160e01b5f52601260045260245ffd5b805f525f5160206139b15f395f51905f5260205260405f20916040516121ce81611de5565b835460ff81166004811015610fb957825260018060a01b039060081c166020820152602460026001860154958660408501520154916001600160401b03831660608201526001600160401b038360401c16608082015260a08101926001600160401b038160801c16845260c01c60c08201525f60018060a01b035f5160206139f15f395f51905f52541660405193848092636af907fb60e11b82528a60048301525afa918215610634575f92612421575b5080516004811015610fb9575f1901611a81575080516006811015610fb9576004146124125760806001600160401b03910151166001600160401b0382511611612352575b50505f8181525f5160206139b15f395f51905f526020908152604091829020805460ff1916600290811782550180546fffffffffffffffff000000000000000019164280851b67ffffffffffffffff60401b169190911790915591516001600160401b0390921682527f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa91a3565b5f5160206139f15f395f51905f52546040805163338587c560e21b815263ffffffff94909416600485015290839060249082905f906001600160a01b03165af1918215610634575f905f936123ee575b508085036123d75750516001600160401b03918216911681106123c557806122c4565b632e19bc2d60e11b5f5260045260245ffd5b849063fee3144560e01b5f5260045260245260445ffd5b905061240a91925060403d604011611d0957611cfa8183611e4b565b91905f6123a2565b50505061241f9150612646565b565b6124369192503d805f833e61067b8183611e4b565b905f61227f565b8181029291811591840414171561065257565b5f5160206139f15f395f51905f525460405163025a076560e61b815263ffffffff929092166004830152602090829060249082905f906001600160a01b03165af1908115610634575f916125bc575b505f5160206139f15f395f51905f5254604051636af907fb60e11b815260048101839052905f90829060249082906001600160a01b03165afa908115610634575f916125a2575b506124f082612c51565b1561259e575f8281525f5160206139915f395f51905f5260209081526040808320545f5160206138f15f395f51905f52909252909120546001600160a01b03908116929116908215612595575b8051926006841015610fb957604061257f926001600160401b039287600461200f9814612585575b50500151165f5160206139d15f395f51905f52549061243d565b906133bd565b61258e9161335a565b5f87612565565b9150809161253d565b5090565b6125b691503d805f833e61067b8183611e4b565b5f6124e6565b90506020813d6020116125e6575b816125d760209383611e4b565b8101031261023b57515f61249f565b3d91506125ca565b60025f516020613a115f395f51905f5254146126175760025f516020613a115f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b906001600160401b03809116911601906001600160401b03821161065257565b805f525f5160206139b15f395f51905f5260205260405f209060405161266b81611de5565b82549060ff82166004811015610fb9578152602081019160018060a01b039060081c1682526002600185015494856040840152015460608201916001600160401b038216835260c06001600160401b038360401c16928360808401526001600160401b038160801c1660a0840152811c9101526004602060018060a01b035f5160206139f15f395f51905f525416604051928380926304e0efb360e11b82525afa908115610634575f9161291c575b5061272d906001600160401b0392612626565b164210612900575f8381525f5160206139b15f395f51905f5260209081526040808320838155600181018490556002018390555f5160206139515f395f51905f529091529081902080546001600160a01b031981169091557f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e9391926001600160a01b039091169081156128ee575b5f925f92875f525f5160206139115f395f51905f52602052855f205480612820575b50505190515f5160206139d15f395f51905f5254612813926001600160a01b03169161257f91906001600160401b031661243d565b82519182526020820152a3565b935093509061257f6001600160401b0361281393895f525f5160206139115f395f51905f526020525f888120558a5f525f5160206139915f395f51905f526020526128a761271061287b61ffff8b5f205460a01c168961243d565b0480978d5f525f5160206139315f395f51905f526020528a5f206128a0838254612964565b905561340e565b966128b28882613674565b8a7f3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf6360208b51938b855260018060a01b031693a39293506127de565b82516001600160a01b031691506127bc565b63fb6ce63f60e01b5f526001600160401b03421660045260245ffd5b90506020813d60201161295c575b8161293760209383611e4b565b8101031261023b576001600160401b039161295461272d92612012565b91509161271a565b3d915061292a565b9190820180921161065257565b5f5160206139f15f395f51905f525490925f916001600160a01b0316803b1561023b575f8091602460405180948193635b73516560e11b83528a60048401525af1801561063457612c3b575b505f5160206139f15f395f51905f5254604051636af907fb60e11b81526004810186905291908390839060249082906001600160a01b03165afa918215612c30578392612c14575b50612a0f85612c51565b15612c0a578483525f5160206139915f395f51905f5260205260408320546001600160a01b03163303612bf75760e08201906001600160401b038251169460c08401956001600160401b03612a85818951168a89525f5160206139915f395f51905f526020528260408a205460b01c1690612626565b1611612bda5791602093916001600160401b03935f14612bb657612aa99088612c79565b955b612b3784612ae581604060018060a01b035f516020613a515f395f51905f525416970151165f5160206139d15f395f51905f52549061243d565b92519351604051634f22429f60e01b815260048101949094526001600160401b0391909416811660248401819052604484015292909416821660648201529516608486015284918290819060a4820190565b03915afa918215612ba9578192612b74575b506040919281525f5160206139315f395f51905f5260205220612b6d828254612964565b9055151590565b91506020823d602011612ba1575b81612b8f60209383611e4b565b8101031261023b576040915191612b49565b3d9150612b82565b50604051903d90823e3d90fd5b508685525f5160206139915f395f51905f5284528260016040872001541695612aab565b825163fb6ce63f60e01b86526001600160401b0316600452602485fd5b636e2ccd7560e11b835233600452602483fd5b5050505050600190565b612c299192503d8085833e61067b8183611e4b565b905f612a05565b6040513d85823e3d90fd5b612c489192505f90611e4b565b5f9060246129bd565b5f9081525f5160206139915f395f51905f5260205260409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff9092166004830152905f816024816005600160991b015afa8015610634575f915f91612e85575b5015612e765780517fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab606548103612e64575060208101516001600160a01b031680612e53575090604080612d2293015181518094819263088c246360e01b835260206004840152602483019061212e565b038173__$e8433aec8a689eedbe2673465b6d21420c$__5af48015610634575f925f91612e2e575b508092808303612e175750815f525f5160206139915f395f51905f526020526001600160401b0380600160405f200154169116115f14612dee57805f525f5160206139915f395f51905f52602052600160405f20016001600160401b0383166001600160401b03198254161790557fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d043560206040516001600160401b0385168152a290565b90505f525f5160206139915f395f51905f526020526001600160401b03600160405f2001541690565b905063fee3144560e01b5f5260045260245260445ffd5b9050612e4a91925060403d604011611d0957611cfa8183611e4b565b9190915f612d4a565b624de75d60e31b5f5260045260245ffd5b636ba589a560e01b5f5260045260245ffd5b636b2f19e960e01b5f5260045ffd5b9150503d805f833e612e978183611e4b565b810160408282031261023b5781516001600160401b03811161023b5782019160608383031261023b5760405192606084018481106001600160401b03821117611e00576040528051845260208101516001600160a01b038116810361023b5760208501526040810151926001600160401b03841161023b57602093612f1c9201611fcd565b6040840152015190811515820361023b57905f612cb2565b919082604091031261023b576020612f4b83612012565b92015190565b9190825f525f5160206139b15f395f51905f5260205260405f209160405191612f7983611de5565b835460ff8116926004841015610fb9576024938552602085019160018060a01b039060081c168252600260018701549687604088015201549260608601926001600160401b038516845260808701946001600160401b038160401c1686526001600160401b038160801c1660a089015260c01c60c08801525f60018060a01b035f5160206139f15f395f51905f52541660405197888092636af907fb60e11b82528c60048301525afa958615610634575f966132ea575b5086516004811015610fb957600119016132cc57516001600160a01b0316330361324b575b875f525f5160206139515f395f51905f5260205260018060a01b0360405f2054169385516006811015610fb9576002036131fc576130af6001600160401b0380925116825f5160206138d15f395f51905f52541690612626565b164210612900576131eb575b50855f525f5160206139b15f395f51905f5260205260405f20600360ff198254161790556001600160401b038060a0600180821b035f5160206139f15f395f51905f525416950151169151169003916001600160401b0383116106525760408051636610966960e01b8152600481018790526001600160401b0394909416602485015290839060449082905f905af1928315610634576131a19386935f916131cb575b505f8481525f5160206139b15f395f51905f526020526040902060020180546001600160c01b031660c09290921b6001600160c01b03191691909117905561341b565b917f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf5f80a3151590565b6131e4915060403d60401161062d5761061b8183611e4b565b505f61315e565b6131f59086612c79565b505f6130bb565b50505050925080516006811015610fb95760040361322f57508261322a93926132249261341b565b50612646565b600190565b516006811015610fb95763170cc93360e21b5f5260045260245ffd5b865f525f5160206139915f395f51905f526020523360018060a01b0360405f20541603610216576001600160401b036132a88160c088015116895f525f5160206139915f395f51905f526020528260405f205460b01c1690612626565b164210156130555763fb6ce63f60e01b5f526001600160401b03421660045260245ffd5b86516004811015610fb957633b0d540d60e21b5f52611a9e90611dd7565b6132ff9196503d805f833e61067b8183611e4b565b945f613030565b6020606081604085019363ffffffff81511686520151936040838201528451809452019201905f5b81811061333b5750505090565b82516001600160a01b031684526020938401939092019160010161332e565b5f8281525f5160206139315f395f51905f5260209081526040822080549290559092917f875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e491906133aa8582613674565b6040519485526001600160a01b031693a3565b5f5160206139715f395f51905f525460405163a9059cbb60e01b60208201526001600160a01b03928316602482015260448082019490945292835261241f929116613409606483611e4b565b613878565b9190820391821161065257565b9160018060a01b035f5160206139f15f395f51905f5254169260408101935f8551602460405180948193636af907fb60e11b835260048301525afa908115610634575f9161365a575b5080516006811015610fb9576003148015613646575b15613624576001600160401b0360e082015116945b60808301916001600160401b038351166001600160401b03881611156136195760209260018060a01b035f516020613a515f395f51905f525416916001600160401b038060c06134f48260608b0151165f5160206139d15f395f51905f52549061243d565b930151935195515f9081525f5160206139915f395f51905f528852604090819020600101549051634f22429f60e01b8152600481019490945293166001600160401b0390811660248401529416841660448201529783166064890152919091166084870152859060a49082905afa938415610634575f946135e5575b506001600160a01b038316156135d0575b50805f525f5160206139115f395f51905f526020528260405f20555f525f5160206139515f395f51905f5260205260405f209060018060a01b03166001600160601b0360a01b82541617905590565b602001516001600160a01b031691505f613581565b9093506020813d602011613611575b8161360160209383611e4b565b8101031261023b5751925f613570565b3d91506135f4565b505050505050505f90565b80516006811015610fb95760020361322f576001600160401b0342169461348f565b5080516006811015610fb95760041461347a565b61366e91503d805f833e61067b8183611e4b565b5f613464565b5f5160206139715f395f51905f52546001600160a01b031691823b1561023b576040516340c10f1960e01b81526001600160a01b039290921660048301526024820152905f908290604490829084905af18015610634576136d25750565b5f61241f91611e4b565b60ff5f516020613a315f395f51905f525460401c16156136f857565b631afcd79f60e31b5f5260045ffd5b6040516370a0823160e01b8152306004820152916001600160a01b03821691602084602481865afa938415610634575f94613840575b509161377860209260249594604051916323b872dd60e01b868401523388840152306044840152606483015260648252613409608483611e4b565b6040516370a0823160e01b815230600482015293849182905afa918215610634575f9261380c575b50808211156137b25761200f9161340e565b60405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608490fd5b9091506020813d602011613838575b8161382860209383611e4b565b8101031261023b5751905f6137a0565b3d915061381b565b9350916020843d602011613870575b8161385c60209383611e4b565b8101031261023b579251929161377861373d565b3d915061384f565b905f602091828151910182855af115610634575f513d6138c757506001600160a01b0381163b155b6138a75750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156138a056feafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab603afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60cafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab609afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60bafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60a6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab00afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab607afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab608afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab604afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00afe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab605a164736f6c634300081e000af0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// ERC20TokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenStakingManagerMetaData.ABI instead.
var ERC20TokenStakingManagerABI = ERC20TokenStakingManagerMetaData.ABI

// ERC20TokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenStakingManagerMetaData.Bin instead.
var ERC20TokenStakingManagerBin = ERC20TokenStakingManagerMetaData.Bin

// DeployERC20TokenStakingManager deploys a new Ethereum contract, binding an instance of ERC20TokenStakingManager to it.
func DeployERC20TokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenStakingManager, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	ERC20TokenStakingManagerBin = strings.ReplaceAll(ERC20TokenStakingManagerBin, "__$e8433aec8a689eedbe2673465b6d21420c$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// ERC20TokenStakingManager is an auto generated Go binding around an Ethereum contract.
type ERC20TokenStakingManager struct {
	ERC20TokenStakingManagerCaller     // Read-only binding to the contract
	ERC20TokenStakingManagerTransactor // Write-only binding to the contract
	ERC20TokenStakingManagerFilterer   // Log filterer for contract events
}

// ERC20TokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenStakingManagerSession struct {
	Contract     *ERC20TokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenStakingManagerCallerSession struct {
	Contract *ERC20TokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20TokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenStakingManagerTransactorSession struct {
	Contract     *ERC20TokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenStakingManagerRaw struct {
	Contract *ERC20TokenStakingManager // Generic contract binding to access the raw methods on
}

// ERC20TokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCallerRaw struct {
	Contract *ERC20TokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactorRaw struct {
	Contract *ERC20TokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenStakingManager creates a new instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManager(address common.Address, backend bind.ContractBackend) (*ERC20TokenStakingManager, error) {
	contract, err := bindERC20TokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// NewERC20TokenStakingManagerCaller creates a new read-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenStakingManagerCaller, error) {
	contract, err := bindERC20TokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCaller{contract: contract}, nil
}

// NewERC20TokenStakingManagerTransactor creates a new write-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenStakingManagerTransactor, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerTransactor{contract: contract}, nil
}

// NewERC20TokenStakingManagerFilterer creates a new log filterer instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenStakingManagerFilterer, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerFilterer{contract: contract}, nil
}

// bindERC20TokenStakingManager binds a generic wrapper to an already deployed contract.
func bindERC20TokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ERC20STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ERC20_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Erc20() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.Erc20(&_ERC20TokenStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) Erc20() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.Erc20(&_ERC20TokenStakingManager.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetDelegatorInfo(opts *bind.CallOpts, delegationID [32]byte) (Delegator, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getDelegatorInfo", delegationID)

	if err != nil {
		return *new(Delegator), err
	}

	out0 := *abi.ConvertType(out[0], new(Delegator)).(*Delegator)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _ERC20TokenStakingManager.Contract.GetDelegatorInfo(&_ERC20TokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _ERC20TokenStakingManager.Contract.GetDelegatorInfo(&_ERC20TokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetDelegatorRewardInfo(opts *bind.CallOpts, delegationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getDelegatorRewardInfo", delegationID)

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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20TokenStakingManager.Contract.GetDelegatorRewardInfo(&_ERC20TokenStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20TokenStakingManager.Contract.GetDelegatorRewardInfo(&_ERC20TokenStakingManager.CallOpts, delegationID)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetStakingManagerSettings(opts *bind.CallOpts) (StakingManagerSettings, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getStakingManagerSettings")

	if err != nil {
		return *new(StakingManagerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerSettings)).(*StakingManagerSettings)

	return out0, err

}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _ERC20TokenStakingManager.Contract.GetStakingManagerSettings(&_ERC20TokenStakingManager.CallOpts)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _ERC20TokenStakingManager.Contract.GetStakingManagerSettings(&_ERC20TokenStakingManager.CallOpts)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetStakingValidator(opts *bind.CallOpts, validationID [32]byte) (PoSValidatorInfo, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getStakingValidator", validationID)

	if err != nil {
		return *new(PoSValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoSValidatorInfo)).(*PoSValidatorInfo)

	return out0, err

}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _ERC20TokenStakingManager.Contract.GetStakingValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _ERC20TokenStakingManager.Contract.GetStakingValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetValidatorRewardInfo(opts *bind.CallOpts, validationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getValidatorRewardInfo", validationID)

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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20TokenStakingManager.Contract.GetValidatorRewardInfo(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20TokenStakingManager.Contract.GetValidatorRewardInfo(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20TokenStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5c93498.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings StakingManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initialize", settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5c93498.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Initialize(settings StakingManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xb5c93498.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) Initialize(settings StakingManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xcd236e7f.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, delegationAmount, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xcd236e7f.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xcd236e7f.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount, rewardRecipient)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x7e65f4da.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x7e65f4da.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x7e65f4da.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, address rewardRecipient) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, rewardRecipient)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitiateValidatorRemoval(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegator(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegator(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
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
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedDelegatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*ERC20TokenStakingManagerCompletedDelegatorRegistration, error) {
	event := new(ERC20TokenStakingManagerCompletedDelegatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
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
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCompletedDelegatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*ERC20TokenStakingManagerCompletedDelegatorRemoval, error) {
	event := new(ERC20TokenStakingManagerCompletedDelegatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRewardClaimedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRewardClaimed)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRewardClaimed)
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
func (it *ERC20TokenStakingManagerDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRewardClaimed struct {
	DelegationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address) (*ERC20TokenStakingManagerDelegatorRewardClaimedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRewardClaimedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRewardClaimed, delegationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRewardClaimed)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRewardClaimed(log types.Log) (*ERC20TokenStakingManagerDelegatorRewardClaimed, error) {
	event := new(ERC20TokenStakingManagerDelegatorRewardClaimed)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator is returned from FilterDelegatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for DelegatorRewardRecipientChanged events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRewardRecipientChanged)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRewardRecipientChanged)
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
func (it *ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRewardRecipientChanged represents a DelegatorRewardRecipientChanged event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRewardRecipientChanged struct {
	DelegationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRewardRecipientChanged(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRewardRecipientChangedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRewardRecipientChanged, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRewardRecipientChanged)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRewardRecipientChanged(log types.Log) (*ERC20TokenStakingManagerDelegatorRewardRecipientChanged, error) {
	event := new(ERC20TokenStakingManagerDelegatorRewardRecipientChanged)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitializedIterator struct {
	Event *ERC20TokenStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialized)
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
		it.Event = new(ERC20TokenStakingManagerInitialized)
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
func (it *ERC20TokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialized represents a Initialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialized(log types.Log) (*ERC20TokenStakingManagerInitialized, error) {
	event := new(ERC20TokenStakingManagerInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRegistration struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedDelegatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*ERC20TokenStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(ERC20TokenStakingManagerInitiatedDelegatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *ERC20TokenStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
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
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedDelegatorRemovalIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*ERC20TokenStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(ERC20TokenStakingManagerInitiatedDelegatorRemoval)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator is returned from FilterInitiatedStakingValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedStakingValidatorRegistration events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator struct {
	Event *ERC20TokenStakingManagerInitiatedStakingValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitiatedStakingValidatorRegistration)
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
		it.Event = new(ERC20TokenStakingManagerInitiatedStakingValidatorRegistration)
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
func (it *ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitiatedStakingValidatorRegistration represents a InitiatedStakingValidatorRegistration event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitiatedStakingValidatorRegistration struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitiatedStakingValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, owner []common.Address) (*ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitiatedStakingValidatorRegistrationIterator{contract: _ERC20TokenStakingManager.contract, event: "InitiatedStakingValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedStakingValidatorRegistration is a free log subscription operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitiatedStakingValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitiatedStakingValidatorRegistration, validationID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitiatedStakingValidatorRegistration)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitiatedStakingValidatorRegistration(log types.Log) (*ERC20TokenStakingManagerInitiatedStakingValidatorRegistration, error) {
	event := new(ERC20TokenStakingManagerInitiatedStakingValidatorRegistration)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdatedIterator struct {
	Event *ERC20TokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
		it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerUptimeUpdatedIterator{contract: _ERC20TokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerUptimeUpdated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*ERC20TokenStakingManagerUptimeUpdated, error) {
	event := new(ERC20TokenStakingManagerUptimeUpdated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorRewardClaimedIterator is returned from FilterValidatorRewardClaimed and is used to iterate over the raw logs and unpacked data for ValidatorRewardClaimed events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRewardClaimedIterator struct {
	Event *ERC20TokenStakingManagerValidatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorRewardClaimed)
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
		it.Event = new(ERC20TokenStakingManagerValidatorRewardClaimed)
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
func (it *ERC20TokenStakingManagerValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorRewardClaimed represents a ValidatorRewardClaimed event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRewardClaimed struct {
	ValidationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardClaimed is a free log retrieval operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorRewardClaimed(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address) (*ERC20TokenStakingManagerValidatorRewardClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorRewardClaimedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardClaimed is a free log subscription operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorRewardClaimed, validationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorRewardClaimed)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorRewardClaimed(log types.Log) (*ERC20TokenStakingManagerValidatorRewardClaimed, error) {
	event := new(ERC20TokenStakingManagerValidatorRewardClaimed)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator is returned from FilterValidatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for ValidatorRewardRecipientChanged events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator struct {
	Event *ERC20TokenStakingManagerValidatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorRewardRecipientChanged)
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
		it.Event = new(ERC20TokenStakingManagerValidatorRewardRecipientChanged)
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
func (it *ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorRewardRecipientChanged represents a ValidatorRewardRecipientChanged event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRewardRecipientChanged struct {
	ValidationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorRewardRecipientChanged(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorRewardRecipientChangedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorRewardRecipientChanged, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorRewardRecipientChanged)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorRewardRecipientChanged(log types.Log) (*ERC20TokenStakingManagerValidatorRewardRecipientChanged, error) {
	event := new(ERC20TokenStakingManagerValidatorRewardRecipientChanged)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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
