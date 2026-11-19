// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplerewardcalculator

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

// ExampleRewardCalculatorMetaData contains all meta data concerning the ExampleRewardCalculator contract.
var ExampleRewardCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"rewardBasisPoints_\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validatorStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uptimeSeconds\",\"type\":\"uint64\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBasisPoints\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a034606c57601f6102f438819003918201601f19168301916001600160401b03831184841017607057808492602094604052833981010312606c57516001600160401b0381168103606c5760805260405161026f90816100858239608051818181607001526102230152f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c80634f22429f146100ed5780635dcc9391146100cf578063a9778a7a146100b3578063afba878a146100985763bb65b24214610050575f80fd5b34610094575f36600319011261009457602060405167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5f80fd5b34610094575f36600319011261009457602060405160508152f35b34610094575f3660031901126100945760206040516127108152f35b34610094575f3660031901126100945760206040516301e133808152f35b346100945760a03660031901126100945760243567ffffffffffffffff81168103610094576044359067ffffffffffffffff821682036100945760643567ffffffffffffffff81168103610094576084359067ffffffffffffffff821682036100945760209361015f936004356101b0565b604051908152f35b9067ffffffffffffffff8091169116039067ffffffffffffffff821161018957565b634e487b7160e01b5f52601160045260245ffd5b8181029291811591840414171561018957565b91929367ffffffffffffffff60649116029067ffffffffffffffff82169182036101895767ffffffffffffffff6101e960509287610167565b16029067ffffffffffffffff8216918203610189571061025b576127109267ffffffffffffffff61024f6301e133809461024961025695847f0000000000000000000000000000000000000000000000000000000000000000169061019d565b93610167565b169061019d565b040490565b5050505f9056fea164736f6c634300081e000a",
}

// ExampleRewardCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleRewardCalculatorMetaData.ABI instead.
var ExampleRewardCalculatorABI = ExampleRewardCalculatorMetaData.ABI

// ExampleRewardCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleRewardCalculatorMetaData.Bin instead.
var ExampleRewardCalculatorBin = ExampleRewardCalculatorMetaData.Bin

// DeployExampleRewardCalculator deploys a new Ethereum contract, binding an instance of ExampleRewardCalculator to it.
func DeployExampleRewardCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, rewardBasisPoints_ uint64) (common.Address, *types.Transaction, *ExampleRewardCalculator, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleRewardCalculatorBin), backend, rewardBasisPoints_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// ExampleRewardCalculator is an auto generated Go binding around an Ethereum contract.
type ExampleRewardCalculator struct {
	ExampleRewardCalculatorCaller     // Read-only binding to the contract
	ExampleRewardCalculatorTransactor // Write-only binding to the contract
	ExampleRewardCalculatorFilterer   // Log filterer for contract events
}

// ExampleRewardCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleRewardCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleRewardCalculatorSession struct {
	Contract     *ExampleRewardCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleRewardCalculatorCallerSession struct {
	Contract *ExampleRewardCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ExampleRewardCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleRewardCalculatorTransactorSession struct {
	Contract     *ExampleRewardCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRewardCalculatorRaw struct {
	Contract *ExampleRewardCalculator // Generic contract binding to access the raw methods on
}

// ExampleRewardCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCallerRaw struct {
	Contract *ExampleRewardCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleRewardCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactorRaw struct {
	Contract *ExampleRewardCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleRewardCalculator creates a new instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculator(address common.Address, backend bind.ContractBackend) (*ExampleRewardCalculator, error) {
	contract, err := bindExampleRewardCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// NewExampleRewardCalculatorCaller creates a new read-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorCaller(address common.Address, caller bind.ContractCaller) (*ExampleRewardCalculatorCaller, error) {
	contract, err := bindExampleRewardCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorCaller{contract: contract}, nil
}

// NewExampleRewardCalculatorTransactor creates a new write-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleRewardCalculatorTransactor, error) {
	contract, err := bindExampleRewardCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorTransactor{contract: contract}, nil
}

// NewExampleRewardCalculatorFilterer creates a new log filterer instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleRewardCalculatorFilterer, error) {
	contract, err := bindExampleRewardCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorFilterer{contract: contract}, nil
}

// bindExampleRewardCalculator binds a generic wrapper to an already deployed contract.
func bindExampleRewardCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) SECONDSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "SECONDS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ExampleRewardCalculator.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ExampleRewardCalculator.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ExampleRewardCalculator.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ExampleRewardCalculator.CallOpts)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) CalculateReward(opts *bind.CallOpts, stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "calculateReward", stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) RewardBasisPoints(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "rewardBasisPoints")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}
