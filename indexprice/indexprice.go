// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package indexprice

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

// IndexpriceMetaData contains all meta data concerning the Indexprice contract.
var IndexpriceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"indexPrice\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IndexpriceABI is the input ABI used to generate the binding from.
// Deprecated: Use IndexpriceMetaData.ABI instead.
var IndexpriceABI = IndexpriceMetaData.ABI

// Indexprice is an auto generated Go binding around an Ethereum contract.
type Indexprice struct {
	IndexpriceCaller     // Read-only binding to the contract
	IndexpriceTransactor // Write-only binding to the contract
	IndexpriceFilterer   // Log filterer for contract events
}

// IndexpriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexpriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexpriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexpriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexpriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexpriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexpriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexpriceSession struct {
	Contract     *Indexprice       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexpriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexpriceCallerSession struct {
	Contract *IndexpriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IndexpriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexpriceTransactorSession struct {
	Contract     *IndexpriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IndexpriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexpriceRaw struct {
	Contract *Indexprice // Generic contract binding to access the raw methods on
}

// IndexpriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexpriceCallerRaw struct {
	Contract *IndexpriceCaller // Generic read-only contract binding to access the raw methods on
}

// IndexpriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexpriceTransactorRaw struct {
	Contract *IndexpriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndexprice creates a new instance of Indexprice, bound to a specific deployed contract.
func NewIndexprice(address common.Address, backend bind.ContractBackend) (*Indexprice, error) {
	contract, err := bindIndexprice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Indexprice{IndexpriceCaller: IndexpriceCaller{contract: contract}, IndexpriceTransactor: IndexpriceTransactor{contract: contract}, IndexpriceFilterer: IndexpriceFilterer{contract: contract}}, nil
}

// NewIndexpriceCaller creates a new read-only instance of Indexprice, bound to a specific deployed contract.
func NewIndexpriceCaller(address common.Address, caller bind.ContractCaller) (*IndexpriceCaller, error) {
	contract, err := bindIndexprice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexpriceCaller{contract: contract}, nil
}

// NewIndexpriceTransactor creates a new write-only instance of Indexprice, bound to a specific deployed contract.
func NewIndexpriceTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexpriceTransactor, error) {
	contract, err := bindIndexprice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexpriceTransactor{contract: contract}, nil
}

// NewIndexpriceFilterer creates a new log filterer instance of Indexprice, bound to a specific deployed contract.
func NewIndexpriceFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexpriceFilterer, error) {
	contract, err := bindIndexprice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexpriceFilterer{contract: contract}, nil
}

// bindIndexprice binds a generic wrapper to an already deployed contract.
func bindIndexprice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndexpriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Indexprice *IndexpriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Indexprice.Contract.IndexpriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Indexprice *IndexpriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Indexprice.Contract.IndexpriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Indexprice *IndexpriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Indexprice.Contract.IndexpriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Indexprice *IndexpriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Indexprice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Indexprice *IndexpriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Indexprice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Indexprice *IndexpriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Indexprice.Contract.contract.Transact(opts, method, params...)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4eda9b42.
//
// Solidity: function indexPrice(address[] addresses) view returns(uint256[])
func (_Indexprice *IndexpriceCaller) IndexPrice(opts *bind.CallOpts, addresses []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Indexprice.contract.Call(opts, &out, "indexPrice", addresses)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// IndexPrice is a free data retrieval call binding the contract method 0x4eda9b42.
//
// Solidity: function indexPrice(address[] addresses) view returns(uint256[])
func (_Indexprice *IndexpriceSession) IndexPrice(addresses []common.Address) ([]*big.Int, error) {
	return _Indexprice.Contract.IndexPrice(&_Indexprice.CallOpts, addresses)
}

// IndexPrice is a free data retrieval call binding the contract method 0x4eda9b42.
//
// Solidity: function indexPrice(address[] addresses) view returns(uint256[])
func (_Indexprice *IndexpriceCallerSession) IndexPrice(addresses []common.Address) ([]*big.Int, error) {
	return _Indexprice.Contract.IndexPrice(&_Indexprice.CallOpts, addresses)
}
