// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bandchain

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IStdReferenceReferenceData is an auto generated low-level Go binding around an user-defined struct.
type IStdReferenceReferenceData struct {
	Rate             *big.Int
	LastUpdatedBase  *big.Int
	LastUpdatedQuote *big.Int
}

// BandchainABI is the input ABI used to generate the binding from.
const BandchainABI = "[{\"inputs\":[{\"internalType\":\"contractIStdReference\",\"name\":\"_ref\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_quote\",\"type\":\"string\"}],\"name\":\"getReferenceData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedQuote\",\"type\":\"uint256\"}],\"internalType\":\"structIStdReference.ReferenceData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_bases\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_quotes\",\"type\":\"string[]\"}],\"name\":\"getReferenceDataBulk\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedQuote\",\"type\":\"uint256\"}],\"internalType\":\"structIStdReference.ReferenceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ref\",\"outputs\":[{\"internalType\":\"contractIStdReference\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStdReference\",\"name\":\"_ref\",\"type\":\"address\"}],\"name\":\"setRef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Bandchain is an auto generated Go binding around an Ethereum contract.
type Bandchain struct {
	BandchainCaller     // Read-only binding to the contract
	BandchainTransactor // Write-only binding to the contract
	BandchainFilterer   // Log filterer for contract events
}

// BandchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BandchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BandchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BandchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BandchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BandchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BandchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BandchainSession struct {
	Contract     *Bandchain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BandchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BandchainCallerSession struct {
	Contract *BandchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BandchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BandchainTransactorSession struct {
	Contract     *BandchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BandchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BandchainRaw struct {
	Contract *Bandchain // Generic contract binding to access the raw methods on
}

// BandchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BandchainCallerRaw struct {
	Contract *BandchainCaller // Generic read-only contract binding to access the raw methods on
}

// BandchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BandchainTransactorRaw struct {
	Contract *BandchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBandchain creates a new instance of Bandchain, bound to a specific deployed contract.
func NewBandchain(address common.Address, backend bind.ContractBackend) (*Bandchain, error) {
	contract, err := bindBandchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bandchain{BandchainCaller: BandchainCaller{contract: contract}, BandchainTransactor: BandchainTransactor{contract: contract}, BandchainFilterer: BandchainFilterer{contract: contract}}, nil
}

// NewBandchainCaller creates a new read-only instance of Bandchain, bound to a specific deployed contract.
func NewBandchainCaller(address common.Address, caller bind.ContractCaller) (*BandchainCaller, error) {
	contract, err := bindBandchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BandchainCaller{contract: contract}, nil
}

// NewBandchainTransactor creates a new write-only instance of Bandchain, bound to a specific deployed contract.
func NewBandchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BandchainTransactor, error) {
	contract, err := bindBandchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BandchainTransactor{contract: contract}, nil
}

// NewBandchainFilterer creates a new log filterer instance of Bandchain, bound to a specific deployed contract.
func NewBandchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BandchainFilterer, error) {
	contract, err := bindBandchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BandchainFilterer{contract: contract}, nil
}

// bindBandchain binds a generic wrapper to an already deployed contract.
func bindBandchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BandchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bandchain *BandchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bandchain.Contract.BandchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bandchain *BandchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bandchain.Contract.BandchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bandchain *BandchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bandchain.Contract.BandchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bandchain *BandchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bandchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bandchain *BandchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bandchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bandchain *BandchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bandchain.Contract.contract.Transact(opts, method, params...)
}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Bandchain *BandchainCaller) GetReferenceData(opts *bind.CallOpts, _base string, _quote string) (IStdReferenceReferenceData, error) {
	var out []interface{}
	err := _Bandchain.contract.Call(opts, &out, "getReferenceData", _base, _quote)

	if err != nil {
		return *new(IStdReferenceReferenceData), err
	}

	out0 := *abi.ConvertType(out[0], new(IStdReferenceReferenceData)).(*IStdReferenceReferenceData)

	return out0, err

}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Bandchain *BandchainSession) GetReferenceData(_base string, _quote string) (IStdReferenceReferenceData, error) {
	return _Bandchain.Contract.GetReferenceData(&_Bandchain.CallOpts, _base, _quote)
}

// GetReferenceData is a free data retrieval call binding the contract method 0x65555bcc.
//
// Solidity: function getReferenceData(string _base, string _quote) view returns((uint256,uint256,uint256))
func (_Bandchain *BandchainCallerSession) GetReferenceData(_base string, _quote string) (IStdReferenceReferenceData, error) {
	return _Bandchain.Contract.GetReferenceData(&_Bandchain.CallOpts, _base, _quote)
}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Bandchain *BandchainCaller) GetReferenceDataBulk(opts *bind.CallOpts, _bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	var out []interface{}
	err := _Bandchain.contract.Call(opts, &out, "getReferenceDataBulk", _bases, _quotes)

	if err != nil {
		return *new([]IStdReferenceReferenceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStdReferenceReferenceData)).(*[]IStdReferenceReferenceData)

	return out0, err

}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Bandchain *BandchainSession) GetReferenceDataBulk(_bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	return _Bandchain.Contract.GetReferenceDataBulk(&_Bandchain.CallOpts, _bases, _quotes)
}

// GetReferenceDataBulk is a free data retrieval call binding the contract method 0xe42a071b.
//
// Solidity: function getReferenceDataBulk(string[] _bases, string[] _quotes) view returns((uint256,uint256,uint256)[])
func (_Bandchain *BandchainCallerSession) GetReferenceDataBulk(_bases []string, _quotes []string) ([]IStdReferenceReferenceData, error) {
	return _Bandchain.Contract.GetReferenceDataBulk(&_Bandchain.CallOpts, _bases, _quotes)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bandchain *BandchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bandchain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bandchain *BandchainSession) Owner() (common.Address, error) {
	return _Bandchain.Contract.Owner(&_Bandchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bandchain *BandchainCallerSession) Owner() (common.Address, error) {
	return _Bandchain.Contract.Owner(&_Bandchain.CallOpts)
}

// Ref is a free data retrieval call binding the contract method 0x21a78f68.
//
// Solidity: function ref() view returns(address)
func (_Bandchain *BandchainCaller) Ref(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bandchain.contract.Call(opts, &out, "ref")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ref is a free data retrieval call binding the contract method 0x21a78f68.
//
// Solidity: function ref() view returns(address)
func (_Bandchain *BandchainSession) Ref() (common.Address, error) {
	return _Bandchain.Contract.Ref(&_Bandchain.CallOpts)
}

// Ref is a free data retrieval call binding the contract method 0x21a78f68.
//
// Solidity: function ref() view returns(address)
func (_Bandchain *BandchainCallerSession) Ref() (common.Address, error) {
	return _Bandchain.Contract.Ref(&_Bandchain.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bandchain *BandchainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bandchain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bandchain *BandchainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bandchain.Contract.RenounceOwnership(&_Bandchain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bandchain *BandchainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bandchain.Contract.RenounceOwnership(&_Bandchain.TransactOpts)
}

// SetRef is a paid mutator transaction binding the contract method 0x6bc855cc.
//
// Solidity: function setRef(address _ref) returns()
func (_Bandchain *BandchainTransactor) SetRef(opts *bind.TransactOpts, _ref common.Address) (*types.Transaction, error) {
	return _Bandchain.contract.Transact(opts, "setRef", _ref)
}

// SetRef is a paid mutator transaction binding the contract method 0x6bc855cc.
//
// Solidity: function setRef(address _ref) returns()
func (_Bandchain *BandchainSession) SetRef(_ref common.Address) (*types.Transaction, error) {
	return _Bandchain.Contract.SetRef(&_Bandchain.TransactOpts, _ref)
}

// SetRef is a paid mutator transaction binding the contract method 0x6bc855cc.
//
// Solidity: function setRef(address _ref) returns()
func (_Bandchain *BandchainTransactorSession) SetRef(_ref common.Address) (*types.Transaction, error) {
	return _Bandchain.Contract.SetRef(&_Bandchain.TransactOpts, _ref)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bandchain *BandchainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bandchain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bandchain *BandchainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bandchain.Contract.TransferOwnership(&_Bandchain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bandchain *BandchainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bandchain.Contract.TransferOwnership(&_Bandchain.TransactOpts, newOwner)
}

// BandchainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bandchain contract.
type BandchainOwnershipTransferredIterator struct {
	Event *BandchainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BandchainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BandchainOwnershipTransferred)
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
		it.Event = new(BandchainOwnershipTransferred)
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
func (it *BandchainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BandchainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BandchainOwnershipTransferred represents a OwnershipTransferred event raised by the Bandchain contract.
type BandchainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bandchain *BandchainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BandchainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bandchain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BandchainOwnershipTransferredIterator{contract: _Bandchain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bandchain *BandchainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BandchainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bandchain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BandchainOwnershipTransferred)
				if err := _Bandchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bandchain *BandchainFilterer) ParseOwnershipTransferred(log types.Log) (*BandchainOwnershipTransferred, error) {
	event := new(BandchainOwnershipTransferred)
	if err := _Bandchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
