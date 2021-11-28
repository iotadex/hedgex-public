// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hedgex

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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
)

// HedgexMetaData contains all meta data concerning the Hedgex contract.
var HedgexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feedPrice\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feedPriceDecimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStartPool\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_leverage\",\"type\":\"uint8\"},{\"internalType\":\"int8\",\"name\":\"_amountDecimal\",\"type\":\"int8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"direction\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Explosive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ForceClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"direction\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakeInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"direction\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDecimal\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"closeLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"closeShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyInterestRateBase\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"detectSlide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divConst\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"explosive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explosivePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDivide\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feedPrice\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feedPriceDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foceCloseRewardGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"forceCloseAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasUsdPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsdPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolNet\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feedPrice\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feedPriceDecimal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStartPool\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_leverage\",\"type\":\"uint8\"},{\"internalType\":\"int8\",\"name\":\"_amountDecimal\",\"type\":\"int8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRewardRate\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStart\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leverage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"openShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolExplosivePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLeftAmountRate\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLongAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLongPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolNetAmountRateLimitOpen\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolNetAmountRateLimitPrice\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolShortAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolShortPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolState\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rechargeMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"singleCloseLimitRate\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"singleOpenLimitRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slideP\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sumFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0Decimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPool\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"traders\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"margin\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"longAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"interestDay\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HedgexABI is the input ABI used to generate the binding from.
// Deprecated: Use HedgexMetaData.ABI instead.
var HedgexABI = HedgexMetaData.ABI

// Hedgex is an auto generated Go binding around an Ethereum contract.
type Hedgex struct {
	HedgexCaller     // Read-only binding to the contract
	HedgexTransactor // Write-only binding to the contract
	HedgexFilterer   // Log filterer for contract events
}

// HedgexCaller is an auto generated read-only Go binding around an Ethereum contract.
type HedgexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HedgexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HedgexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HedgexSession struct {
	Contract     *Hedgex           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HedgexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HedgexCallerSession struct {
	Contract *HedgexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HedgexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HedgexTransactorSession struct {
	Contract     *HedgexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HedgexRaw is an auto generated low-level Go binding around an Ethereum contract.
type HedgexRaw struct {
	Contract *Hedgex // Generic contract binding to access the raw methods on
}

// HedgexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HedgexCallerRaw struct {
	Contract *HedgexCaller // Generic read-only contract binding to access the raw methods on
}

// HedgexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HedgexTransactorRaw struct {
	Contract *HedgexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHedgex creates a new instance of Hedgex, bound to a specific deployed contract.
func NewHedgex(address common.Address, backend bind.ContractBackend) (*Hedgex, error) {
	contract, err := bindHedgex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hedgex{HedgexCaller: HedgexCaller{contract: contract}, HedgexTransactor: HedgexTransactor{contract: contract}, HedgexFilterer: HedgexFilterer{contract: contract}}, nil
}

// NewHedgexCaller creates a new read-only instance of Hedgex, bound to a specific deployed contract.
func NewHedgexCaller(address common.Address, caller bind.ContractCaller) (*HedgexCaller, error) {
	contract, err := bindHedgex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HedgexCaller{contract: contract}, nil
}

// NewHedgexTransactor creates a new write-only instance of Hedgex, bound to a specific deployed contract.
func NewHedgexTransactor(address common.Address, transactor bind.ContractTransactor) (*HedgexTransactor, error) {
	contract, err := bindHedgex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HedgexTransactor{contract: contract}, nil
}

// NewHedgexFilterer creates a new log filterer instance of Hedgex, bound to a specific deployed contract.
func NewHedgexFilterer(address common.Address, filterer bind.ContractFilterer) (*HedgexFilterer, error) {
	contract, err := bindHedgex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HedgexFilterer{contract: contract}, nil
}

// bindHedgex binds a generic wrapper to an already deployed contract.
func bindHedgex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HedgexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hedgex *HedgexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hedgex.Contract.HedgexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hedgex *HedgexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgex.Contract.HedgexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hedgex *HedgexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hedgex.Contract.HedgexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hedgex *HedgexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hedgex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hedgex *HedgexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hedgex *HedgexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hedgex.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hedgex *HedgexCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hedgex *HedgexSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hedgex.Contract.DOMAINSEPARATOR(&_Hedgex.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Hedgex *HedgexCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Hedgex.Contract.DOMAINSEPARATOR(&_Hedgex.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Hedgex *HedgexCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Hedgex *HedgexSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Hedgex.Contract.PERMITTYPEHASH(&_Hedgex.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Hedgex *HedgexCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Hedgex.Contract.PERMITTYPEHASH(&_Hedgex.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Hedgex *HedgexCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Hedgex *HedgexSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.Allowance(&_Hedgex.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Hedgex *HedgexCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.Allowance(&_Hedgex.CallOpts, arg0, arg1)
}

// AmountDecimal is a free data retrieval call binding the contract method 0x2be90282.
//
// Solidity: function amountDecimal() view returns(int8)
func (_Hedgex *HedgexCaller) AmountDecimal(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "amountDecimal")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// AmountDecimal is a free data retrieval call binding the contract method 0x2be90282.
//
// Solidity: function amountDecimal() view returns(int8)
func (_Hedgex *HedgexSession) AmountDecimal() (int8, error) {
	return _Hedgex.Contract.AmountDecimal(&_Hedgex.CallOpts)
}

// AmountDecimal is a free data retrieval call binding the contract method 0x2be90282.
//
// Solidity: function amountDecimal() view returns(int8)
func (_Hedgex *HedgexCallerSession) AmountDecimal() (int8, error) {
	return _Hedgex.Contract.AmountDecimal(&_Hedgex.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Hedgex *HedgexCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Hedgex *HedgexSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.BalanceOf(&_Hedgex.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Hedgex *HedgexCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.BalanceOf(&_Hedgex.CallOpts, arg0)
}

// DailyInterestRateBase is a free data retrieval call binding the contract method 0x8967b694.
//
// Solidity: function dailyInterestRateBase() view returns(uint16)
func (_Hedgex *HedgexCaller) DailyInterestRateBase(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "dailyInterestRateBase")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyInterestRateBase is a free data retrieval call binding the contract method 0x8967b694.
//
// Solidity: function dailyInterestRateBase() view returns(uint16)
func (_Hedgex *HedgexSession) DailyInterestRateBase() (uint16, error) {
	return _Hedgex.Contract.DailyInterestRateBase(&_Hedgex.CallOpts)
}

// DailyInterestRateBase is a free data retrieval call binding the contract method 0x8967b694.
//
// Solidity: function dailyInterestRateBase() view returns(uint16)
func (_Hedgex *HedgexCallerSession) DailyInterestRateBase() (uint16, error) {
	return _Hedgex.Contract.DailyInterestRateBase(&_Hedgex.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hedgex *HedgexCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hedgex *HedgexSession) Decimals() (uint8, error) {
	return _Hedgex.Contract.Decimals(&_Hedgex.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hedgex *HedgexCallerSession) Decimals() (uint8, error) {
	return _Hedgex.Contract.Decimals(&_Hedgex.CallOpts)
}

// DivConst is a free data retrieval call binding the contract method 0x7e4856c4.
//
// Solidity: function divConst() view returns(uint24)
func (_Hedgex *HedgexCaller) DivConst(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "divConst")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DivConst is a free data retrieval call binding the contract method 0x7e4856c4.
//
// Solidity: function divConst() view returns(uint24)
func (_Hedgex *HedgexSession) DivConst() (*big.Int, error) {
	return _Hedgex.Contract.DivConst(&_Hedgex.CallOpts)
}

// DivConst is a free data retrieval call binding the contract method 0x7e4856c4.
//
// Solidity: function divConst() view returns(uint24)
func (_Hedgex *HedgexCallerSession) DivConst() (*big.Int, error) {
	return _Hedgex.Contract.DivConst(&_Hedgex.CallOpts)
}

// FeeDivide is a free data retrieval call binding the contract method 0xd375454f.
//
// Solidity: function feeDivide() view returns(uint24)
func (_Hedgex *HedgexCaller) FeeDivide(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "feeDivide")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeDivide is a free data retrieval call binding the contract method 0xd375454f.
//
// Solidity: function feeDivide() view returns(uint24)
func (_Hedgex *HedgexSession) FeeDivide() (*big.Int, error) {
	return _Hedgex.Contract.FeeDivide(&_Hedgex.CallOpts)
}

// FeeDivide is a free data retrieval call binding the contract method 0xd375454f.
//
// Solidity: function feeDivide() view returns(uint24)
func (_Hedgex *HedgexCallerSession) FeeDivide() (*big.Int, error) {
	return _Hedgex.Contract.FeeDivide(&_Hedgex.CallOpts)
}

// FeeOn is a free data retrieval call binding the contract method 0x4f335d0a.
//
// Solidity: function feeOn() view returns(bool)
func (_Hedgex *HedgexCaller) FeeOn(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "feeOn")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FeeOn is a free data retrieval call binding the contract method 0x4f335d0a.
//
// Solidity: function feeOn() view returns(bool)
func (_Hedgex *HedgexSession) FeeOn() (bool, error) {
	return _Hedgex.Contract.FeeOn(&_Hedgex.CallOpts)
}

// FeeOn is a free data retrieval call binding the contract method 0x4f335d0a.
//
// Solidity: function feeOn() view returns(bool)
func (_Hedgex *HedgexCallerSession) FeeOn() (bool, error) {
	return _Hedgex.Contract.FeeOn(&_Hedgex.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint16)
func (_Hedgex *HedgexCaller) FeeRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint16)
func (_Hedgex *HedgexSession) FeeRate() (uint16, error) {
	return _Hedgex.Contract.FeeRate(&_Hedgex.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint16)
func (_Hedgex *HedgexCallerSession) FeeRate() (uint16, error) {
	return _Hedgex.Contract.FeeRate(&_Hedgex.CallOpts)
}

// FeedPrice is a free data retrieval call binding the contract method 0xa04eecae.
//
// Solidity: function feedPrice() view returns(address)
func (_Hedgex *HedgexCaller) FeedPrice(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "feedPrice")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeedPrice is a free data retrieval call binding the contract method 0xa04eecae.
//
// Solidity: function feedPrice() view returns(address)
func (_Hedgex *HedgexSession) FeedPrice() (common.Address, error) {
	return _Hedgex.Contract.FeedPrice(&_Hedgex.CallOpts)
}

// FeedPrice is a free data retrieval call binding the contract method 0xa04eecae.
//
// Solidity: function feedPrice() view returns(address)
func (_Hedgex *HedgexCallerSession) FeedPrice() (common.Address, error) {
	return _Hedgex.Contract.FeedPrice(&_Hedgex.CallOpts)
}

// FeedPriceDecimal is a free data retrieval call binding the contract method 0x08c73e74.
//
// Solidity: function feedPriceDecimal() view returns(uint256)
func (_Hedgex *HedgexCaller) FeedPriceDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "feedPriceDecimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeedPriceDecimal is a free data retrieval call binding the contract method 0x08c73e74.
//
// Solidity: function feedPriceDecimal() view returns(uint256)
func (_Hedgex *HedgexSession) FeedPriceDecimal() (*big.Int, error) {
	return _Hedgex.Contract.FeedPriceDecimal(&_Hedgex.CallOpts)
}

// FeedPriceDecimal is a free data retrieval call binding the contract method 0x08c73e74.
//
// Solidity: function feedPriceDecimal() view returns(uint256)
func (_Hedgex *HedgexCallerSession) FeedPriceDecimal() (*big.Int, error) {
	return _Hedgex.Contract.FeedPriceDecimal(&_Hedgex.CallOpts)
}

// FoceCloseRewardGas is a free data retrieval call binding the contract method 0x3d5f8e7a.
//
// Solidity: function foceCloseRewardGas() view returns(uint256)
func (_Hedgex *HedgexCaller) FoceCloseRewardGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "foceCloseRewardGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoceCloseRewardGas is a free data retrieval call binding the contract method 0x3d5f8e7a.
//
// Solidity: function foceCloseRewardGas() view returns(uint256)
func (_Hedgex *HedgexSession) FoceCloseRewardGas() (*big.Int, error) {
	return _Hedgex.Contract.FoceCloseRewardGas(&_Hedgex.CallOpts)
}

// FoceCloseRewardGas is a free data retrieval call binding the contract method 0x3d5f8e7a.
//
// Solidity: function foceCloseRewardGas() view returns(uint256)
func (_Hedgex *HedgexCallerSession) FoceCloseRewardGas() (*big.Int, error) {
	return _Hedgex.Contract.FoceCloseRewardGas(&_Hedgex.CallOpts)
}

// GasUsdPriceFeed is a free data retrieval call binding the contract method 0x6afb3760.
//
// Solidity: function gasUsdPriceFeed() view returns(address)
func (_Hedgex *HedgexCaller) GasUsdPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "gasUsdPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasUsdPriceFeed is a free data retrieval call binding the contract method 0x6afb3760.
//
// Solidity: function gasUsdPriceFeed() view returns(address)
func (_Hedgex *HedgexSession) GasUsdPriceFeed() (common.Address, error) {
	return _Hedgex.Contract.GasUsdPriceFeed(&_Hedgex.CallOpts)
}

// GasUsdPriceFeed is a free data retrieval call binding the contract method 0x6afb3760.
//
// Solidity: function gasUsdPriceFeed() view returns(address)
func (_Hedgex *HedgexCallerSession) GasUsdPriceFeed() (common.Address, error) {
	return _Hedgex.Contract.GasUsdPriceFeed(&_Hedgex.CallOpts)
}

// GetGasUsdPrice is a free data retrieval call binding the contract method 0x8ed3896c.
//
// Solidity: function getGasUsdPrice() view returns(uint256)
func (_Hedgex *HedgexCaller) GetGasUsdPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "getGasUsdPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasUsdPrice is a free data retrieval call binding the contract method 0x8ed3896c.
//
// Solidity: function getGasUsdPrice() view returns(uint256)
func (_Hedgex *HedgexSession) GetGasUsdPrice() (*big.Int, error) {
	return _Hedgex.Contract.GetGasUsdPrice(&_Hedgex.CallOpts)
}

// GetGasUsdPrice is a free data retrieval call binding the contract method 0x8ed3896c.
//
// Solidity: function getGasUsdPrice() view returns(uint256)
func (_Hedgex *HedgexCallerSession) GetGasUsdPrice() (*big.Int, error) {
	return _Hedgex.Contract.GetGasUsdPrice(&_Hedgex.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_Hedgex *HedgexCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_Hedgex *HedgexSession) GetLatestPrice() (*big.Int, error) {
	return _Hedgex.Contract.GetLatestPrice(&_Hedgex.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(uint256)
func (_Hedgex *HedgexCallerSession) GetLatestPrice() (*big.Int, error) {
	return _Hedgex.Contract.GetLatestPrice(&_Hedgex.CallOpts)
}

// GetPoolNet is a free data retrieval call binding the contract method 0x7bb71a97.
//
// Solidity: function getPoolNet() view returns(int256)
func (_Hedgex *HedgexCaller) GetPoolNet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "getPoolNet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolNet is a free data retrieval call binding the contract method 0x7bb71a97.
//
// Solidity: function getPoolNet() view returns(int256)
func (_Hedgex *HedgexSession) GetPoolNet() (*big.Int, error) {
	return _Hedgex.Contract.GetPoolNet(&_Hedgex.CallOpts)
}

// GetPoolNet is a free data retrieval call binding the contract method 0x7bb71a97.
//
// Solidity: function getPoolNet() view returns(int256)
func (_Hedgex *HedgexCallerSession) GetPoolNet() (*big.Int, error) {
	return _Hedgex.Contract.GetPoolNet(&_Hedgex.CallOpts)
}

// InterestRewardRate is a free data retrieval call binding the contract method 0x01bcc908.
//
// Solidity: function interestRewardRate() view returns(uint24)
func (_Hedgex *HedgexCaller) InterestRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "interestRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRewardRate is a free data retrieval call binding the contract method 0x01bcc908.
//
// Solidity: function interestRewardRate() view returns(uint24)
func (_Hedgex *HedgexSession) InterestRewardRate() (*big.Int, error) {
	return _Hedgex.Contract.InterestRewardRate(&_Hedgex.CallOpts)
}

// InterestRewardRate is a free data retrieval call binding the contract method 0x01bcc908.
//
// Solidity: function interestRewardRate() view returns(uint24)
func (_Hedgex *HedgexCallerSession) InterestRewardRate() (*big.Int, error) {
	return _Hedgex.Contract.InterestRewardRate(&_Hedgex.CallOpts)
}

// IsStart is a free data retrieval call binding the contract method 0x8a55d36e.
//
// Solidity: function isStart() view returns(bool)
func (_Hedgex *HedgexCaller) IsStart(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "isStart")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStart is a free data retrieval call binding the contract method 0x8a55d36e.
//
// Solidity: function isStart() view returns(bool)
func (_Hedgex *HedgexSession) IsStart() (bool, error) {
	return _Hedgex.Contract.IsStart(&_Hedgex.CallOpts)
}

// IsStart is a free data retrieval call binding the contract method 0x8a55d36e.
//
// Solidity: function isStart() view returns(bool)
func (_Hedgex *HedgexCallerSession) IsStart() (bool, error) {
	return _Hedgex.Contract.IsStart(&_Hedgex.CallOpts)
}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Hedgex *HedgexCaller) Leverage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "leverage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Hedgex *HedgexSession) Leverage() (uint8, error) {
	return _Hedgex.Contract.Leverage(&_Hedgex.CallOpts)
}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Hedgex *HedgexCallerSession) Leverage() (uint8, error) {
	return _Hedgex.Contract.Leverage(&_Hedgex.CallOpts)
}

// MinPool is a free data retrieval call binding the contract method 0x2678b8fd.
//
// Solidity: function minPool() view returns(uint256)
func (_Hedgex *HedgexCaller) MinPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "minPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPool is a free data retrieval call binding the contract method 0x2678b8fd.
//
// Solidity: function minPool() view returns(uint256)
func (_Hedgex *HedgexSession) MinPool() (*big.Int, error) {
	return _Hedgex.Contract.MinPool(&_Hedgex.CallOpts)
}

// MinPool is a free data retrieval call binding the contract method 0x2678b8fd.
//
// Solidity: function minPool() view returns(uint256)
func (_Hedgex *HedgexCallerSession) MinPool() (*big.Int, error) {
	return _Hedgex.Contract.MinPool(&_Hedgex.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgex *HedgexCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgex *HedgexSession) Name() (string, error) {
	return _Hedgex.Contract.Name(&_Hedgex.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgex *HedgexCallerSession) Name() (string, error) {
	return _Hedgex.Contract.Name(&_Hedgex.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Hedgex *HedgexCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Hedgex *HedgexSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.Nonces(&_Hedgex.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Hedgex *HedgexCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Hedgex.Contract.Nonces(&_Hedgex.CallOpts, arg0)
}

// PoolExplosivePrice is a free data retrieval call binding the contract method 0x33d6f1f7.
//
// Solidity: function poolExplosivePrice() view returns(uint256)
func (_Hedgex *HedgexCaller) PoolExplosivePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolExplosivePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolExplosivePrice is a free data retrieval call binding the contract method 0x33d6f1f7.
//
// Solidity: function poolExplosivePrice() view returns(uint256)
func (_Hedgex *HedgexSession) PoolExplosivePrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolExplosivePrice(&_Hedgex.CallOpts)
}

// PoolExplosivePrice is a free data retrieval call binding the contract method 0x33d6f1f7.
//
// Solidity: function poolExplosivePrice() view returns(uint256)
func (_Hedgex *HedgexCallerSession) PoolExplosivePrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolExplosivePrice(&_Hedgex.CallOpts)
}

// PoolLeftAmountRate is a free data retrieval call binding the contract method 0xce8ad1d2.
//
// Solidity: function poolLeftAmountRate() view returns(uint24)
func (_Hedgex *HedgexCaller) PoolLeftAmountRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolLeftAmountRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLeftAmountRate is a free data retrieval call binding the contract method 0xce8ad1d2.
//
// Solidity: function poolLeftAmountRate() view returns(uint24)
func (_Hedgex *HedgexSession) PoolLeftAmountRate() (*big.Int, error) {
	return _Hedgex.Contract.PoolLeftAmountRate(&_Hedgex.CallOpts)
}

// PoolLeftAmountRate is a free data retrieval call binding the contract method 0xce8ad1d2.
//
// Solidity: function poolLeftAmountRate() view returns(uint24)
func (_Hedgex *HedgexCallerSession) PoolLeftAmountRate() (*big.Int, error) {
	return _Hedgex.Contract.PoolLeftAmountRate(&_Hedgex.CallOpts)
}

// PoolLongAmount is a free data retrieval call binding the contract method 0xe5984319.
//
// Solidity: function poolLongAmount() view returns(uint256)
func (_Hedgex *HedgexCaller) PoolLongAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolLongAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLongAmount is a free data retrieval call binding the contract method 0xe5984319.
//
// Solidity: function poolLongAmount() view returns(uint256)
func (_Hedgex *HedgexSession) PoolLongAmount() (*big.Int, error) {
	return _Hedgex.Contract.PoolLongAmount(&_Hedgex.CallOpts)
}

// PoolLongAmount is a free data retrieval call binding the contract method 0xe5984319.
//
// Solidity: function poolLongAmount() view returns(uint256)
func (_Hedgex *HedgexCallerSession) PoolLongAmount() (*big.Int, error) {
	return _Hedgex.Contract.PoolLongAmount(&_Hedgex.CallOpts)
}

// PoolLongPrice is a free data retrieval call binding the contract method 0x994af3ae.
//
// Solidity: function poolLongPrice() view returns(uint256)
func (_Hedgex *HedgexCaller) PoolLongPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolLongPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLongPrice is a free data retrieval call binding the contract method 0x994af3ae.
//
// Solidity: function poolLongPrice() view returns(uint256)
func (_Hedgex *HedgexSession) PoolLongPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolLongPrice(&_Hedgex.CallOpts)
}

// PoolLongPrice is a free data retrieval call binding the contract method 0x994af3ae.
//
// Solidity: function poolLongPrice() view returns(uint256)
func (_Hedgex *HedgexCallerSession) PoolLongPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolLongPrice(&_Hedgex.CallOpts)
}

// PoolNetAmountRateLimitOpen is a free data retrieval call binding the contract method 0xcd8bb6cd.
//
// Solidity: function poolNetAmountRateLimitOpen() view returns(int24)
func (_Hedgex *HedgexCaller) PoolNetAmountRateLimitOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolNetAmountRateLimitOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolNetAmountRateLimitOpen is a free data retrieval call binding the contract method 0xcd8bb6cd.
//
// Solidity: function poolNetAmountRateLimitOpen() view returns(int24)
func (_Hedgex *HedgexSession) PoolNetAmountRateLimitOpen() (*big.Int, error) {
	return _Hedgex.Contract.PoolNetAmountRateLimitOpen(&_Hedgex.CallOpts)
}

// PoolNetAmountRateLimitOpen is a free data retrieval call binding the contract method 0xcd8bb6cd.
//
// Solidity: function poolNetAmountRateLimitOpen() view returns(int24)
func (_Hedgex *HedgexCallerSession) PoolNetAmountRateLimitOpen() (*big.Int, error) {
	return _Hedgex.Contract.PoolNetAmountRateLimitOpen(&_Hedgex.CallOpts)
}

// PoolNetAmountRateLimitPrice is a free data retrieval call binding the contract method 0x1cce9103.
//
// Solidity: function poolNetAmountRateLimitPrice() view returns(int24)
func (_Hedgex *HedgexCaller) PoolNetAmountRateLimitPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolNetAmountRateLimitPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolNetAmountRateLimitPrice is a free data retrieval call binding the contract method 0x1cce9103.
//
// Solidity: function poolNetAmountRateLimitPrice() view returns(int24)
func (_Hedgex *HedgexSession) PoolNetAmountRateLimitPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolNetAmountRateLimitPrice(&_Hedgex.CallOpts)
}

// PoolNetAmountRateLimitPrice is a free data retrieval call binding the contract method 0x1cce9103.
//
// Solidity: function poolNetAmountRateLimitPrice() view returns(int24)
func (_Hedgex *HedgexCallerSession) PoolNetAmountRateLimitPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolNetAmountRateLimitPrice(&_Hedgex.CallOpts)
}

// PoolShortAmount is a free data retrieval call binding the contract method 0x8e406892.
//
// Solidity: function poolShortAmount() view returns(uint256)
func (_Hedgex *HedgexCaller) PoolShortAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolShortAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolShortAmount is a free data retrieval call binding the contract method 0x8e406892.
//
// Solidity: function poolShortAmount() view returns(uint256)
func (_Hedgex *HedgexSession) PoolShortAmount() (*big.Int, error) {
	return _Hedgex.Contract.PoolShortAmount(&_Hedgex.CallOpts)
}

// PoolShortAmount is a free data retrieval call binding the contract method 0x8e406892.
//
// Solidity: function poolShortAmount() view returns(uint256)
func (_Hedgex *HedgexCallerSession) PoolShortAmount() (*big.Int, error) {
	return _Hedgex.Contract.PoolShortAmount(&_Hedgex.CallOpts)
}

// PoolShortPrice is a free data retrieval call binding the contract method 0x971059fb.
//
// Solidity: function poolShortPrice() view returns(uint256)
func (_Hedgex *HedgexCaller) PoolShortPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolShortPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolShortPrice is a free data retrieval call binding the contract method 0x971059fb.
//
// Solidity: function poolShortPrice() view returns(uint256)
func (_Hedgex *HedgexSession) PoolShortPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolShortPrice(&_Hedgex.CallOpts)
}

// PoolShortPrice is a free data retrieval call binding the contract method 0x971059fb.
//
// Solidity: function poolShortPrice() view returns(uint256)
func (_Hedgex *HedgexCallerSession) PoolShortPrice() (*big.Int, error) {
	return _Hedgex.Contract.PoolShortPrice(&_Hedgex.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_Hedgex *HedgexCaller) PoolState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "poolState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_Hedgex *HedgexSession) PoolState() (uint8, error) {
	return _Hedgex.Contract.PoolState(&_Hedgex.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_Hedgex *HedgexCallerSession) PoolState() (uint8, error) {
	return _Hedgex.Contract.PoolState(&_Hedgex.CallOpts)
}

// SingleCloseLimitRate is a free data retrieval call binding the contract method 0x7aaae950.
//
// Solidity: function singleCloseLimitRate() view returns(uint24)
func (_Hedgex *HedgexCaller) SingleCloseLimitRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "singleCloseLimitRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleCloseLimitRate is a free data retrieval call binding the contract method 0x7aaae950.
//
// Solidity: function singleCloseLimitRate() view returns(uint24)
func (_Hedgex *HedgexSession) SingleCloseLimitRate() (*big.Int, error) {
	return _Hedgex.Contract.SingleCloseLimitRate(&_Hedgex.CallOpts)
}

// SingleCloseLimitRate is a free data retrieval call binding the contract method 0x7aaae950.
//
// Solidity: function singleCloseLimitRate() view returns(uint24)
func (_Hedgex *HedgexCallerSession) SingleCloseLimitRate() (*big.Int, error) {
	return _Hedgex.Contract.SingleCloseLimitRate(&_Hedgex.CallOpts)
}

// SingleOpenLimitRate is a free data retrieval call binding the contract method 0x5d0291ac.
//
// Solidity: function singleOpenLimitRate() view returns(uint16)
func (_Hedgex *HedgexCaller) SingleOpenLimitRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "singleOpenLimitRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SingleOpenLimitRate is a free data retrieval call binding the contract method 0x5d0291ac.
//
// Solidity: function singleOpenLimitRate() view returns(uint16)
func (_Hedgex *HedgexSession) SingleOpenLimitRate() (uint16, error) {
	return _Hedgex.Contract.SingleOpenLimitRate(&_Hedgex.CallOpts)
}

// SingleOpenLimitRate is a free data retrieval call binding the contract method 0x5d0291ac.
//
// Solidity: function singleOpenLimitRate() view returns(uint16)
func (_Hedgex *HedgexCallerSession) SingleOpenLimitRate() (uint16, error) {
	return _Hedgex.Contract.SingleOpenLimitRate(&_Hedgex.CallOpts)
}

// SlideP is a free data retrieval call binding the contract method 0x2c204fa8.
//
// Solidity: function slideP() view returns(uint8)
func (_Hedgex *HedgexCaller) SlideP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "slideP")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SlideP is a free data retrieval call binding the contract method 0x2c204fa8.
//
// Solidity: function slideP() view returns(uint8)
func (_Hedgex *HedgexSession) SlideP() (uint8, error) {
	return _Hedgex.Contract.SlideP(&_Hedgex.CallOpts)
}

// SlideP is a free data retrieval call binding the contract method 0x2c204fa8.
//
// Solidity: function slideP() view returns(uint8)
func (_Hedgex *HedgexCallerSession) SlideP() (uint8, error) {
	return _Hedgex.Contract.SlideP(&_Hedgex.CallOpts)
}

// SumFee is a free data retrieval call binding the contract method 0x190b065f.
//
// Solidity: function sumFee() view returns(uint256)
func (_Hedgex *HedgexCaller) SumFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "sumFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumFee is a free data retrieval call binding the contract method 0x190b065f.
//
// Solidity: function sumFee() view returns(uint256)
func (_Hedgex *HedgexSession) SumFee() (*big.Int, error) {
	return _Hedgex.Contract.SumFee(&_Hedgex.CallOpts)
}

// SumFee is a free data retrieval call binding the contract method 0x190b065f.
//
// Solidity: function sumFee() view returns(uint256)
func (_Hedgex *HedgexCallerSession) SumFee() (*big.Int, error) {
	return _Hedgex.Contract.SumFee(&_Hedgex.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgex *HedgexCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgex *HedgexSession) Symbol() (string, error) {
	return _Hedgex.Contract.Symbol(&_Hedgex.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgex *HedgexCallerSession) Symbol() (string, error) {
	return _Hedgex.Contract.Symbol(&_Hedgex.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Hedgex *HedgexCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Hedgex *HedgexSession) Token0() (common.Address, error) {
	return _Hedgex.Contract.Token0(&_Hedgex.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Hedgex *HedgexCallerSession) Token0() (common.Address, error) {
	return _Hedgex.Contract.Token0(&_Hedgex.CallOpts)
}

// Token0Decimal is a free data retrieval call binding the contract method 0x9ddc230a.
//
// Solidity: function token0Decimal() view returns(uint256)
func (_Hedgex *HedgexCaller) Token0Decimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "token0Decimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Token0Decimal is a free data retrieval call binding the contract method 0x9ddc230a.
//
// Solidity: function token0Decimal() view returns(uint256)
func (_Hedgex *HedgexSession) Token0Decimal() (*big.Int, error) {
	return _Hedgex.Contract.Token0Decimal(&_Hedgex.CallOpts)
}

// Token0Decimal is a free data retrieval call binding the contract method 0x9ddc230a.
//
// Solidity: function token0Decimal() view returns(uint256)
func (_Hedgex *HedgexCallerSession) Token0Decimal() (*big.Int, error) {
	return _Hedgex.Contract.Token0Decimal(&_Hedgex.CallOpts)
}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(int256)
func (_Hedgex *HedgexCaller) TotalPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "totalPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(int256)
func (_Hedgex *HedgexSession) TotalPool() (*big.Int, error) {
	return _Hedgex.Contract.TotalPool(&_Hedgex.CallOpts)
}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(int256)
func (_Hedgex *HedgexCallerSession) TotalPool() (*big.Int, error) {
	return _Hedgex.Contract.TotalPool(&_Hedgex.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgex *HedgexCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgex *HedgexSession) TotalSupply() (*big.Int, error) {
	return _Hedgex.Contract.TotalSupply(&_Hedgex.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgex *HedgexCallerSession) TotalSupply() (*big.Int, error) {
	return _Hedgex.Contract.TotalSupply(&_Hedgex.CallOpts)
}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(int256 margin, uint256 longAmount, uint256 longPrice, uint256 shortAmount, uint256 shortPrice, uint32 interestDay)
func (_Hedgex *HedgexCaller) Traders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Margin      *big.Int
	LongAmount  *big.Int
	LongPrice   *big.Int
	ShortAmount *big.Int
	ShortPrice  *big.Int
	InterestDay uint32
}, error) {
	var out []interface{}
	err := _Hedgex.contract.Call(opts, &out, "traders", arg0)

	outstruct := new(struct {
		Margin      *big.Int
		LongAmount  *big.Int
		LongPrice   *big.Int
		ShortAmount *big.Int
		ShortPrice  *big.Int
		InterestDay uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Margin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LongAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LongPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ShortPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.InterestDay = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(int256 margin, uint256 longAmount, uint256 longPrice, uint256 shortAmount, uint256 shortPrice, uint32 interestDay)
func (_Hedgex *HedgexSession) Traders(arg0 common.Address) (struct {
	Margin      *big.Int
	LongAmount  *big.Int
	LongPrice   *big.Int
	ShortAmount *big.Int
	ShortPrice  *big.Int
	InterestDay uint32
}, error) {
	return _Hedgex.Contract.Traders(&_Hedgex.CallOpts, arg0)
}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(int256 margin, uint256 longAmount, uint256 longPrice, uint256 shortAmount, uint256 shortPrice, uint32 interestDay)
func (_Hedgex *HedgexCallerSession) Traders(arg0 common.Address) (struct {
	Margin      *big.Int
	LongAmount  *big.Int
	LongPrice   *big.Int
	ShortAmount *big.Int
	ShortPrice  *big.Int
	InterestDay uint32
}, error) {
	return _Hedgex.Contract.Traders(&_Hedgex.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xc95f9d0e.
//
// Solidity: function addLiquidity(uint256 amount, address to) returns()
func (_Hedgex *HedgexTransactor) AddLiquidity(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "addLiquidity", amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xc95f9d0e.
//
// Solidity: function addLiquidity(uint256 amount, address to) returns()
func (_Hedgex *HedgexSession) AddLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.AddLiquidity(&_Hedgex.TransactOpts, amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xc95f9d0e.
//
// Solidity: function addLiquidity(uint256 amount, address to) returns()
func (_Hedgex *HedgexTransactorSession) AddLiquidity(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.AddLiquidity(&_Hedgex.TransactOpts, amount, to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hedgex *HedgexSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.Approve(&_Hedgex.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.Approve(&_Hedgex.TransactOpts, spender, value)
}

// CloseLong is a paid mutator transaction binding the contract method 0x90edbeb0.
//
// Solidity: function closeLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactor) CloseLong(opts *bind.TransactOpts, priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "closeLong", priceExp, amount, deadline)
}

// CloseLong is a paid mutator transaction binding the contract method 0x90edbeb0.
//
// Solidity: function closeLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexSession) CloseLong(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.CloseLong(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// CloseLong is a paid mutator transaction binding the contract method 0x90edbeb0.
//
// Solidity: function closeLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactorSession) CloseLong(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.CloseLong(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// CloseShort is a paid mutator transaction binding the contract method 0x807ca353.
//
// Solidity: function closeShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactor) CloseShort(opts *bind.TransactOpts, priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "closeShort", priceExp, amount, deadline)
}

// CloseShort is a paid mutator transaction binding the contract method 0x807ca353.
//
// Solidity: function closeShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexSession) CloseShort(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.CloseShort(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// CloseShort is a paid mutator transaction binding the contract method 0x807ca353.
//
// Solidity: function closeShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactorSession) CloseShort(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.CloseShort(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// DetectSlide is a paid mutator transaction binding the contract method 0x3d07bfb7.
//
// Solidity: function detectSlide(address account) returns()
func (_Hedgex *HedgexTransactor) DetectSlide(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "detectSlide", account)
}

// DetectSlide is a paid mutator transaction binding the contract method 0x3d07bfb7.
//
// Solidity: function detectSlide(address account) returns()
func (_Hedgex *HedgexSession) DetectSlide(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.DetectSlide(&_Hedgex.TransactOpts, account)
}

// DetectSlide is a paid mutator transaction binding the contract method 0x3d07bfb7.
//
// Solidity: function detectSlide(address account) returns()
func (_Hedgex *HedgexTransactorSession) DetectSlide(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.DetectSlide(&_Hedgex.TransactOpts, account)
}

// Explosive is a paid mutator transaction binding the contract method 0xfc1e2c45.
//
// Solidity: function explosive(address account) returns()
func (_Hedgex *HedgexTransactor) Explosive(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "explosive", account)
}

// Explosive is a paid mutator transaction binding the contract method 0xfc1e2c45.
//
// Solidity: function explosive(address account) returns()
func (_Hedgex *HedgexSession) Explosive(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.Explosive(&_Hedgex.TransactOpts, account)
}

// Explosive is a paid mutator transaction binding the contract method 0xfc1e2c45.
//
// Solidity: function explosive(address account) returns()
func (_Hedgex *HedgexTransactorSession) Explosive(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.Explosive(&_Hedgex.TransactOpts, account)
}

// ExplosivePool is a paid mutator transaction binding the contract method 0x3e31fb45.
//
// Solidity: function explosivePool() returns()
func (_Hedgex *HedgexTransactor) ExplosivePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "explosivePool")
}

// ExplosivePool is a paid mutator transaction binding the contract method 0x3e31fb45.
//
// Solidity: function explosivePool() returns()
func (_Hedgex *HedgexSession) ExplosivePool() (*types.Transaction, error) {
	return _Hedgex.Contract.ExplosivePool(&_Hedgex.TransactOpts)
}

// ExplosivePool is a paid mutator transaction binding the contract method 0x3e31fb45.
//
// Solidity: function explosivePool() returns()
func (_Hedgex *HedgexTransactorSession) ExplosivePool() (*types.Transaction, error) {
	return _Hedgex.Contract.ExplosivePool(&_Hedgex.TransactOpts)
}

// ForceCloseAccount is a paid mutator transaction binding the contract method 0x93f8d812.
//
// Solidity: function forceCloseAccount(address account) returns()
func (_Hedgex *HedgexTransactor) ForceCloseAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "forceCloseAccount", account)
}

// ForceCloseAccount is a paid mutator transaction binding the contract method 0x93f8d812.
//
// Solidity: function forceCloseAccount(address account) returns()
func (_Hedgex *HedgexSession) ForceCloseAccount(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.ForceCloseAccount(&_Hedgex.TransactOpts, account)
}

// ForceCloseAccount is a paid mutator transaction binding the contract method 0x93f8d812.
//
// Solidity: function forceCloseAccount(address account) returns()
func (_Hedgex *HedgexTransactorSession) ForceCloseAccount(account common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.ForceCloseAccount(&_Hedgex.TransactOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x251c073a.
//
// Solidity: function initialize(address _token0, address _feedPrice, uint256 _feedPriceDecimal, uint256 _minStartPool, uint8 _leverage, int8 _amountDecimal) returns()
func (_Hedgex *HedgexTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _feedPrice common.Address, _feedPriceDecimal *big.Int, _minStartPool *big.Int, _leverage uint8, _amountDecimal int8) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "initialize", _token0, _feedPrice, _feedPriceDecimal, _minStartPool, _leverage, _amountDecimal)
}

// Initialize is a paid mutator transaction binding the contract method 0x251c073a.
//
// Solidity: function initialize(address _token0, address _feedPrice, uint256 _feedPriceDecimal, uint256 _minStartPool, uint8 _leverage, int8 _amountDecimal) returns()
func (_Hedgex *HedgexSession) Initialize(_token0 common.Address, _feedPrice common.Address, _feedPriceDecimal *big.Int, _minStartPool *big.Int, _leverage uint8, _amountDecimal int8) (*types.Transaction, error) {
	return _Hedgex.Contract.Initialize(&_Hedgex.TransactOpts, _token0, _feedPrice, _feedPriceDecimal, _minStartPool, _leverage, _amountDecimal)
}

// Initialize is a paid mutator transaction binding the contract method 0x251c073a.
//
// Solidity: function initialize(address _token0, address _feedPrice, uint256 _feedPriceDecimal, uint256 _minStartPool, uint8 _leverage, int8 _amountDecimal) returns()
func (_Hedgex *HedgexTransactorSession) Initialize(_token0 common.Address, _feedPrice common.Address, _feedPriceDecimal *big.Int, _minStartPool *big.Int, _leverage uint8, _amountDecimal int8) (*types.Transaction, error) {
	return _Hedgex.Contract.Initialize(&_Hedgex.TransactOpts, _token0, _feedPrice, _feedPriceDecimal, _minStartPool, _leverage, _amountDecimal)
}

// OpenLong is a paid mutator transaction binding the contract method 0x6c2ee359.
//
// Solidity: function openLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactor) OpenLong(opts *bind.TransactOpts, priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "openLong", priceExp, amount, deadline)
}

// OpenLong is a paid mutator transaction binding the contract method 0x6c2ee359.
//
// Solidity: function openLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexSession) OpenLong(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.OpenLong(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// OpenLong is a paid mutator transaction binding the contract method 0x6c2ee359.
//
// Solidity: function openLong(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactorSession) OpenLong(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.OpenLong(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// OpenShort is a paid mutator transaction binding the contract method 0x71b31707.
//
// Solidity: function openShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactor) OpenShort(opts *bind.TransactOpts, priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "openShort", priceExp, amount, deadline)
}

// OpenShort is a paid mutator transaction binding the contract method 0x71b31707.
//
// Solidity: function openShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexSession) OpenShort(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.OpenShort(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// OpenShort is a paid mutator transaction binding the contract method 0x71b31707.
//
// Solidity: function openShort(uint256 priceExp, uint256 amount, uint256 deadline) returns()
func (_Hedgex *HedgexTransactorSession) OpenShort(priceExp *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.OpenShort(&_Hedgex.TransactOpts, priceExp, amount, deadline)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hedgex *HedgexTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hedgex *HedgexSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hedgex.Contract.Permit(&_Hedgex.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Hedgex *HedgexTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Hedgex.Contract.Permit(&_Hedgex.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RechargeMargin is a paid mutator transaction binding the contract method 0x11ce83f8.
//
// Solidity: function rechargeMargin(uint256 amount) returns()
func (_Hedgex *HedgexTransactor) RechargeMargin(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "rechargeMargin", amount)
}

// RechargeMargin is a paid mutator transaction binding the contract method 0x11ce83f8.
//
// Solidity: function rechargeMargin(uint256 amount) returns()
func (_Hedgex *HedgexSession) RechargeMargin(amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.RechargeMargin(&_Hedgex.TransactOpts, amount)
}

// RechargeMargin is a paid mutator transaction binding the contract method 0x11ce83f8.
//
// Solidity: function rechargeMargin(uint256 amount) returns()
func (_Hedgex *HedgexTransactorSession) RechargeMargin(amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.RechargeMargin(&_Hedgex.TransactOpts, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 liquidity, address to) returns()
func (_Hedgex *HedgexTransactor) RemoveLiquidity(opts *bind.TransactOpts, liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "removeLiquidity", liquidity, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 liquidity, address to) returns()
func (_Hedgex *HedgexSession) RemoveLiquidity(liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.RemoveLiquidity(&_Hedgex.TransactOpts, liquidity, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x05fe138b.
//
// Solidity: function removeLiquidity(uint256 liquidity, address to) returns()
func (_Hedgex *HedgexTransactorSession) RemoveLiquidity(liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _Hedgex.Contract.RemoveLiquidity(&_Hedgex.TransactOpts, liquidity, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hedgex *HedgexSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.Transfer(&_Hedgex.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.Transfer(&_Hedgex.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hedgex *HedgexSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.TransferFrom(&_Hedgex.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Hedgex *HedgexTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.TransferFrom(&_Hedgex.TransactOpts, from, to, value)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x0cea7534.
//
// Solidity: function withdrawMargin(uint256 amount) returns()
func (_Hedgex *HedgexTransactor) WithdrawMargin(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.contract.Transact(opts, "withdrawMargin", amount)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x0cea7534.
//
// Solidity: function withdrawMargin(uint256 amount) returns()
func (_Hedgex *HedgexSession) WithdrawMargin(amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.WithdrawMargin(&_Hedgex.TransactOpts, amount)
}

// WithdrawMargin is a paid mutator transaction binding the contract method 0x0cea7534.
//
// Solidity: function withdrawMargin(uint256 amount) returns()
func (_Hedgex *HedgexTransactorSession) WithdrawMargin(amount *big.Int) (*types.Transaction, error) {
	return _Hedgex.Contract.WithdrawMargin(&_Hedgex.TransactOpts, amount)
}

// HedgexApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hedgex contract.
type HedgexApprovalIterator struct {
	Event *HedgexApproval // Event containing the contract specifics and raw log

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
func (it *HedgexApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexApproval)
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
		it.Event = new(HedgexApproval)
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
func (it *HedgexApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexApproval represents a Approval event raised by the Hedgex contract.
type HedgexApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hedgex *HedgexFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HedgexApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexApprovalIterator{contract: _Hedgex.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Hedgex *HedgexFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HedgexApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexApproval)
				if err := _Hedgex.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Hedgex *HedgexFilterer) ParseApproval(log types.Log) (*HedgexApproval, error) {
	event := new(HedgexApproval)
	if err := _Hedgex.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Hedgex contract.
type HedgexBurnIterator struct {
	Event *HedgexBurn // Event containing the contract specifics and raw log

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
func (it *HedgexBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexBurn)
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
		it.Event = new(HedgexBurn)
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
func (it *HedgexBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexBurn represents a Burn event raised by the Hedgex contract.
type HedgexBurn struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address) (*HedgexBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Burn", senderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexBurnIterator{contract: _Hedgex.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *HedgexBurn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Burn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexBurn)
				if err := _Hedgex.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) ParseBurn(log types.Log) (*HedgexBurn, error) {
	event := new(HedgexBurn)
	if err := _Hedgex.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexExplosiveIterator is returned from FilterExplosive and is used to iterate over the raw logs and unpacked data for Explosive events raised by the Hedgex contract.
type HedgexExplosiveIterator struct {
	Event *HedgexExplosive // Event containing the contract specifics and raw log

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
func (it *HedgexExplosiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexExplosive)
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
		it.Event = new(HedgexExplosive)
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
func (it *HedgexExplosiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexExplosiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexExplosive represents a Explosive event raised by the Hedgex contract.
type HedgexExplosive struct {
	User      common.Address
	Direction int8
	Amount    *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExplosive is a free log retrieval operation binding the contract event 0x1cf7369c3767c976b447baaae4eed8217e2d190d1750c589e62bf5f7ecc4aaff.
//
// Solidity: event Explosive(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) FilterExplosive(opts *bind.FilterOpts, user []common.Address) (*HedgexExplosiveIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Explosive", userRule)
	if err != nil {
		return nil, err
	}
	return &HedgexExplosiveIterator{contract: _Hedgex.contract, event: "Explosive", logs: logs, sub: sub}, nil
}

// WatchExplosive is a free log subscription operation binding the contract event 0x1cf7369c3767c976b447baaae4eed8217e2d190d1750c589e62bf5f7ecc4aaff.
//
// Solidity: event Explosive(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) WatchExplosive(opts *bind.WatchOpts, sink chan<- *HedgexExplosive, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Explosive", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexExplosive)
				if err := _Hedgex.contract.UnpackLog(event, "Explosive", log); err != nil {
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

// ParseExplosive is a log parse operation binding the contract event 0x1cf7369c3767c976b447baaae4eed8217e2d190d1750c589e62bf5f7ecc4aaff.
//
// Solidity: event Explosive(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) ParseExplosive(log types.Log) (*HedgexExplosive, error) {
	event := new(HedgexExplosive)
	if err := _Hedgex.contract.UnpackLog(event, "Explosive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexForceCloseIterator is returned from FilterForceClose and is used to iterate over the raw logs and unpacked data for ForceClose events raised by the Hedgex contract.
type HedgexForceCloseIterator struct {
	Event *HedgexForceClose // Event containing the contract specifics and raw log

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
func (it *HedgexForceCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexForceClose)
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
		it.Event = new(HedgexForceClose)
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
func (it *HedgexForceCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexForceCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexForceClose represents a ForceClose event raised by the Hedgex contract.
type HedgexForceClose struct {
	Account common.Address
	Long    *big.Int
	Short   *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterForceClose is a free log retrieval operation binding the contract event 0xfca559ae7a3bea5c41671f27049d90b84fbf98be965e6d7f0259ee6e5caa980b.
//
// Solidity: event ForceClose(address indexed account, uint256 long, uint256 short, uint256 price)
func (_Hedgex *HedgexFilterer) FilterForceClose(opts *bind.FilterOpts, account []common.Address) (*HedgexForceCloseIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "ForceClose", accountRule)
	if err != nil {
		return nil, err
	}
	return &HedgexForceCloseIterator{contract: _Hedgex.contract, event: "ForceClose", logs: logs, sub: sub}, nil
}

// WatchForceClose is a free log subscription operation binding the contract event 0xfca559ae7a3bea5c41671f27049d90b84fbf98be965e6d7f0259ee6e5caa980b.
//
// Solidity: event ForceClose(address indexed account, uint256 long, uint256 short, uint256 price)
func (_Hedgex *HedgexFilterer) WatchForceClose(opts *bind.WatchOpts, sink chan<- *HedgexForceClose, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "ForceClose", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexForceClose)
				if err := _Hedgex.contract.UnpackLog(event, "ForceClose", log); err != nil {
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

// ParseForceClose is a log parse operation binding the contract event 0xfca559ae7a3bea5c41671f27049d90b84fbf98be965e6d7f0259ee6e5caa980b.
//
// Solidity: event ForceClose(address indexed account, uint256 long, uint256 short, uint256 price)
func (_Hedgex *HedgexFilterer) ParseForceClose(log types.Log) (*HedgexForceClose, error) {
	event := new(HedgexForceClose)
	if err := _Hedgex.contract.UnpackLog(event, "ForceClose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Hedgex contract.
type HedgexMintIterator struct {
	Event *HedgexMint // Event containing the contract specifics and raw log

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
func (it *HedgexMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexMint)
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
		it.Event = new(HedgexMint)
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
func (it *HedgexMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexMint represents a Mint event raised by the Hedgex contract.
type HedgexMint struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*HedgexMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexMintIterator{contract: _Hedgex.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *HedgexMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexMint)
				if err := _Hedgex.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) ParseMint(log types.Log) (*HedgexMint, error) {
	event := new(HedgexMint)
	if err := _Hedgex.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the Hedgex contract.
type HedgexRechargeIterator struct {
	Event *HedgexRecharge // Event containing the contract specifics and raw log

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
func (it *HedgexRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexRecharge)
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
		it.Event = new(HedgexRecharge)
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
func (it *HedgexRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexRecharge represents a Recharge event raised by the Hedgex contract.
type HedgexRecharge struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x78d9de6f3ad1cae9e0cbcbaa5267fd90f6a6728831eec42b7c147b398b226924.
//
// Solidity: event Recharge(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) FilterRecharge(opts *bind.FilterOpts, sender []common.Address) (*HedgexRechargeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Recharge", senderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexRechargeIterator{contract: _Hedgex.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x78d9de6f3ad1cae9e0cbcbaa5267fd90f6a6728831eec42b7c147b398b226924.
//
// Solidity: event Recharge(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *HedgexRecharge, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Recharge", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexRecharge)
				if err := _Hedgex.contract.UnpackLog(event, "Recharge", log); err != nil {
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

// ParseRecharge is a log parse operation binding the contract event 0x78d9de6f3ad1cae9e0cbcbaa5267fd90f6a6728831eec42b7c147b398b226924.
//
// Solidity: event Recharge(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) ParseRecharge(log types.Log) (*HedgexRecharge, error) {
	event := new(HedgexRecharge)
	if err := _Hedgex.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexTakeInterestIterator is returned from FilterTakeInterest and is used to iterate over the raw logs and unpacked data for TakeInterest events raised by the Hedgex contract.
type HedgexTakeInterestIterator struct {
	Event *HedgexTakeInterest // Event containing the contract specifics and raw log

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
func (it *HedgexTakeInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexTakeInterest)
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
		it.Event = new(HedgexTakeInterest)
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
func (it *HedgexTakeInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexTakeInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexTakeInterest represents a TakeInterest event raised by the Hedgex contract.
type HedgexTakeInterest struct {
	User      common.Address
	Direction int8
	Amount    *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTakeInterest is a free log retrieval operation binding the contract event 0x78213dfb97e98b601b3fee012a0f6bddbf4e63c293c77b5186137af1e7cf62ae.
//
// Solidity: event TakeInterest(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) FilterTakeInterest(opts *bind.FilterOpts, user []common.Address) (*HedgexTakeInterestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "TakeInterest", userRule)
	if err != nil {
		return nil, err
	}
	return &HedgexTakeInterestIterator{contract: _Hedgex.contract, event: "TakeInterest", logs: logs, sub: sub}, nil
}

// WatchTakeInterest is a free log subscription operation binding the contract event 0x78213dfb97e98b601b3fee012a0f6bddbf4e63c293c77b5186137af1e7cf62ae.
//
// Solidity: event TakeInterest(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) WatchTakeInterest(opts *bind.WatchOpts, sink chan<- *HedgexTakeInterest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "TakeInterest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexTakeInterest)
				if err := _Hedgex.contract.UnpackLog(event, "TakeInterest", log); err != nil {
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

// ParseTakeInterest is a log parse operation binding the contract event 0x78213dfb97e98b601b3fee012a0f6bddbf4e63c293c77b5186137af1e7cf62ae.
//
// Solidity: event TakeInterest(address indexed user, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) ParseTakeInterest(log types.Log) (*HedgexTakeInterest, error) {
	event := new(HedgexTakeInterest)
	if err := _Hedgex.contract.UnpackLog(event, "TakeInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Hedgex contract.
type HedgexTradeIterator struct {
	Event *HedgexTrade // Event containing the contract specifics and raw log

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
func (it *HedgexTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexTrade)
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
		it.Event = new(HedgexTrade)
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
func (it *HedgexTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexTrade represents a Trade event raised by the Hedgex contract.
type HedgexTrade struct {
	Sender    common.Address
	Direction int8
	Amount    *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x303db7636b0e071f088d2cd5e8463efe513531272d61a18739b21053f6d7b430.
//
// Solidity: event Trade(address indexed sender, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) FilterTrade(opts *bind.FilterOpts, sender []common.Address) (*HedgexTradeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Trade", senderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexTradeIterator{contract: _Hedgex.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x303db7636b0e071f088d2cd5e8463efe513531272d61a18739b21053f6d7b430.
//
// Solidity: event Trade(address indexed sender, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *HedgexTrade, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Trade", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexTrade)
				if err := _Hedgex.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x303db7636b0e071f088d2cd5e8463efe513531272d61a18739b21053f6d7b430.
//
// Solidity: event Trade(address indexed sender, int8 direction, uint256 amount, uint256 price)
func (_Hedgex *HedgexFilterer) ParseTrade(log types.Log) (*HedgexTrade, error) {
	event := new(HedgexTrade)
	if err := _Hedgex.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hedgex contract.
type HedgexTransferIterator struct {
	Event *HedgexTransfer // Event containing the contract specifics and raw log

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
func (it *HedgexTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexTransfer)
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
		it.Event = new(HedgexTransfer)
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
func (it *HedgexTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexTransfer represents a Transfer event raised by the Hedgex contract.
type HedgexTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hedgex *HedgexFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HedgexTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HedgexTransferIterator{contract: _Hedgex.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hedgex *HedgexFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HedgexTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexTransfer)
				if err := _Hedgex.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hedgex *HedgexFilterer) ParseTransfer(log types.Log) (*HedgexTransfer, error) {
	event := new(HedgexTransfer)
	if err := _Hedgex.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgexWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Hedgex contract.
type HedgexWithdrawIterator struct {
	Event *HedgexWithdraw // Event containing the contract specifics and raw log

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
func (it *HedgexWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgexWithdraw)
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
		it.Event = new(HedgexWithdraw)
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
func (it *HedgexWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgexWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgexWithdraw represents a Withdraw event raised by the Hedgex contract.
type HedgexWithdraw struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*HedgexWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &HedgexWithdrawIterator{contract: _Hedgex.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *HedgexWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Hedgex.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgexWithdraw)
				if err := _Hedgex.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount)
func (_Hedgex *HedgexFilterer) ParseWithdraw(log types.Log) (*HedgexWithdraw, error) {
	event := new(HedgexWithdraw)
	if err := _Hedgex.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
