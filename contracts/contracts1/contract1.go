// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts1

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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasicTokenFuncSigs maps the 4-byte function signature to its string representation.
var BasicTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
var BasicTokenBin = "0x60806040526a52b7d2dcc80cd2e400000060005534801561001f57600080fd5b506102098061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806318160ddd1461004657806370a0823114610060578063a9059cbb14610086575b600080fd5b61004e6100c6565b60408051918252519081900360200190f35b61004e6004803603602081101561007657600080fd5b50356001600160a01b03166100cc565b6100b26004803603604081101561009c57600080fd5b506001600160a01b0381351690602001356100e7565b604080519115158252519081900360200190f35b60005481565b6001600160a01b031660009081526001602052604090205490565b60006001600160a01b0383166100fc57600080fd5b3360009081526001602052604090205461011c908363ffffffff6101ac16565b33600090815260016020526040808220929092556001600160a01b0385168152205461014e908363ffffffff6101be16565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6000828211156101b857fe5b50900390565b6000828201838110156101cd57fe5b939250505056fea265627a7a723158204680e999e9a2c5860832ffb5eaa178ad4fec6b5f506e59c3448054a8fc632a1264736f6c63430005100032"

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasicToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BasicToken *BasicTokenFilterer) ParseTransfer(log types.Log) (*BasicTokenTransfer, error) {
	event := new(BasicTokenTransfer)
	if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Basic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenABI is the input ABI used to generate the binding from.
const MintableTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MintableTokenFuncSigs maps the 4-byte function signature to its string representation.
var MintableTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"7d64bcb4": "finishMinting()",
	"893d20e8": "getOwner()",
	"d73dd623": "increaseApproval(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"05d2035b": "mintingFinished()",
	"8da5cb5b": "owner()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// MintableTokenBin is the compiled bytecode used for deploying new contracts.
var MintableTokenBin = "0x60806040526a52b7d2dcc80cd2e4000000600055600380546001600160a81b0319163317905561097d806100346000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637d64bcb41161008c578063a9059cbb11610066578063a9059cbb14610239578063d73dd62314610265578063dd62ed3e14610291578063f2fde38b146102bf576100ea565b80637d64bcb414610205578063893d20e81461020d5780638da5cb5b14610231576100ea565b806323b872dd116100c857806323b872dd1461015157806340c10f191461018757806366188463146101b357806370a08231146101df576100ea565b806305d2035b146100ef578063095ea7b31461010b57806318160ddd14610137575b600080fd5b6100f76102e7565b604080519115158252519081900360200190f35b6100f76004803603604081101561012157600080fd5b506001600160a01b0381351690602001356102f7565b61013f61035d565b60408051918252519081900360200190f35b6100f76004803603606081101561016757600080fd5b506001600160a01b03813581169160208101359091169060400135610363565b6100f76004803603604081101561019d57600080fd5b506001600160a01b038135169060200135610481565b6100f7600480360360408110156101c957600080fd5b506001600160a01b03813516906020013561058c565b61013f600480360360208110156101f557600080fd5b50356001600160a01b031661067c565b6100f7610697565b6102156106f3565b604080516001600160a01b039092168252519081900360200190f35b610215610702565b6100f76004803603604081101561024f57600080fd5b506001600160a01b038135169060200135610711565b6100f76004803603604081101561027b57600080fd5b506001600160a01b0381351690602001356107d6565b61013f600480360360408110156102a757600080fd5b506001600160a01b038135811691602001351661086f565b6102e5600480360360208110156102d557600080fd5b50356001600160a01b031661089a565b005b600354600160a01b900460ff1681565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661037857600080fd5b6001600160a01b038416600081815260026020908152604080832033845282528083205493835260019091529020546103b7908463ffffffff61092016565b6001600160a01b0380871660009081526001602052604080822093909355908616815220546103ec908463ffffffff61093216565b6001600160a01b038516600090815260016020526040902055610415818463ffffffff61092016565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b6003546000906001600160a01b0316331461049b57600080fd5b600354600160a01b900460ff16156104b257600080fd5b6000546104c5908363ffffffff61093216565b60009081556001600160a01b0384168152600160205260409020546104f0908363ffffffff61093216565b6001600160a01b038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a26040805183815290516001600160a01b038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808311156105e1573360009081526002602090815260408083206001600160a01b0388168452909152812055610616565b6105f1818463ffffffff61092016565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b6003546000906001600160a01b031633146106b157600080fd5b6003805460ff60a01b1916600160a01b1790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b6003546001600160a01b031690565b6003546001600160a01b031681565b60006001600160a01b03831661072657600080fd5b33600090815260016020526040902054610746908363ffffffff61092016565b33600090815260016020526040808220929092556001600160a01b03851681522054610778908363ffffffff61093216565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b038616845290915281205461080a908363ffffffff61093216565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b031633146108b157600080fd5b6001600160a01b0381166108c457600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b60008282111561092c57fe5b50900390565b60008282018381101561094157fe5b939250505056fea265627a7a72315820a4fb99580d6c2f89bd4ac550f7e6aa72aafab93027b4e529df3d8f8e78133f9464736f6c63430005100032"

// DeployMintableToken deploys a new Ethereum contract, binding an instance of MintableToken to it.
func DeployMintableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MintableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// MintableToken is an auto generated Go binding around an Ethereum contract.
type MintableToken struct {
	MintableTokenCaller     // Read-only binding to the contract
	MintableTokenTransactor // Write-only binding to the contract
	MintableTokenFilterer   // Log filterer for contract events
}

// MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableTokenSession struct {
	Contract     *MintableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableTokenCallerSession struct {
	Contract *MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableTokenTransactorSession struct {
	Contract     *MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableTokenRaw struct {
	Contract *MintableToken // Generic contract binding to access the raw methods on
}

// MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableTokenCallerRaw struct {
	Contract *MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableTokenTransactorRaw struct {
	Contract *MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableToken creates a new instance of MintableToken, bound to a specific deployed contract.
func NewMintableToken(address common.Address, backend bind.ContractBackend) (*MintableToken, error) {
	contract, err := bindMintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// NewMintableTokenCaller creates a new read-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableTokenCaller, error) {
	contract, err := bindMintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenCaller{contract: contract}, nil
}

// NewMintableTokenTransactor creates a new write-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableTokenTransactor, error) {
	contract, err := bindMintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransactor{contract: contract}, nil
}

// NewMintableTokenFilterer creates a new log filterer instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableTokenFilterer, error) {
	contract, err := bindMintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableTokenFilterer{contract: contract}, nil
}

// bindMintableToken binds a generic wrapper to an already deployed contract.
func bindMintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_MintableToken *MintableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_MintableToken *MintableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenSession) GetOwner() (common.Address, error) {
	return _MintableToken.Contract.GetOwner(&_MintableToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MintableToken *MintableTokenCallerSession) GetOwner() (common.Address, error) {
	return _MintableToken.Contract.GetOwner(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "mintingFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_MintableToken *MintableTokenCallerSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintableToken *MintableTokenCallerSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintableToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableToken *MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_MintableToken *MintableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintableToken *MintableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// MintableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintableToken contract.
type MintableTokenApprovalIterator struct {
	Event *MintableTokenApproval // Event containing the contract specifics and raw log

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
func (it *MintableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenApproval)
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
		it.Event = new(MintableTokenApproval)
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
func (it *MintableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenApproval represents a Approval event raised by the MintableToken contract.
type MintableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableToken *MintableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MintableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenApprovalIterator{contract: _MintableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableToken *MintableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenApproval)
				if err := _MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseApproval(log types.Log) (*MintableTokenApproval, error) {
	event := new(MintableTokenApproval)
	if err := _MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MintableToken contract.
type MintableTokenMintIterator struct {
	Event *MintableTokenMint // Event containing the contract specifics and raw log

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
func (it *MintableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMint)
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
		it.Event = new(MintableTokenMint)
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
func (it *MintableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMint represents a Mint event raised by the MintableToken contract.
type MintableTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MintableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintIterator{contract: _MintableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MintableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMint)
				if err := _MintableToken.contract.UnpackLog(event, "Mint", log); err != nil {
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
// Solidity: event Mint(address indexed to, uint256 amount)
func (_MintableToken *MintableTokenFilterer) ParseMint(log types.Log) (*MintableTokenMint, error) {
	event := new(MintableTokenMint)
	if err := _MintableToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the MintableToken contract.
type MintableTokenMintFinishedIterator struct {
	Event *MintableTokenMintFinished // Event containing the contract specifics and raw log

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
func (it *MintableTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMintFinished)
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
		it.Event = new(MintableTokenMintFinished)
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
func (it *MintableTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMintFinished represents a MintFinished event raised by the MintableToken contract.
type MintableTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*MintableTokenMintFinishedIterator, error) {

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintFinishedIterator{contract: _MintableToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *MintableTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMintFinished)
				if err := _MintableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// ParseMintFinished is a log parse operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_MintableToken *MintableTokenFilterer) ParseMintFinished(log types.Log) (*MintableTokenMintFinished, error) {
	event := new(MintableTokenMintFinished)
	if err := _MintableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintableToken contract.
type MintableTokenOwnershipTransferredIterator struct {
	Event *MintableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenOwnershipTransferred)
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
		it.Event = new(MintableTokenOwnershipTransferred)
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
func (it *MintableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MintableToken contract.
type MintableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintableToken *MintableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenOwnershipTransferredIterator{contract: _MintableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintableToken *MintableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenOwnershipTransferred)
				if err := _MintableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MintableTokenOwnershipTransferred, error) {
	event := new(MintableTokenOwnershipTransferred)
	if err := _MintableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintableToken contract.
type MintableTokenTransferIterator struct {
	Event *MintableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MintableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenTransfer)
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
		it.Event = new(MintableTokenTransfer)
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
func (it *MintableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenTransfer represents a Transfer event raised by the MintableToken contract.
type MintableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableToken *MintableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MintableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransferIterator{contract: _MintableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableToken *MintableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenTransfer)
				if err := _MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MintableToken *MintableTokenFilterer) ParseTransfer(log types.Log) (*MintableTokenTransfer, error) {
	event := new(MintableTokenTransfer)
	if err := _MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"893d20e8": "getOwner()",
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610172806100326000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063893d20e8146100465780638da5cb5b1461006a578063f2fde38b14610072575b600080fd5b61004e61009a565b604080516001600160a01b039092168252519081900360200190f35b61004e6100a9565b6100986004803603602081101561008857600080fd5b50356001600160a01b03166100b8565b005b6000546001600160a01b031690565b6000546001600160a01b031681565b6000546001600160a01b031633146100cf57600080fd5b6001600160a01b0381166100e257600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820f913644562749c8b0264710b8924c03bc0041818cf9f178e8c6b8fba9805320f64736f6c63430005100032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Ownable *OwnableCallerSession) GetOwner() (common.Address, error) {
	return _Ownable.Contract.GetOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820044af1e00e696012416fa223007c0d20928982f87fdb2b17892ce6ab3893fb5564736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMABI is the input ABI used to generate the binding from.
const SafeMathMABI = "[]"

// SafeMathMBin is the compiled bytecode used for deploying new contracts.
var SafeMathMBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582052e59e6f6f57266366f14c0ad99ade796e89d6cfa2bd47f13d160d1e228c890564736f6c63430005100032"

// DeploySafeMathM deploys a new Ethereum contract, binding an instance of SafeMathM to it.
func DeploySafeMathM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathM, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathM{SafeMathMCaller: SafeMathMCaller{contract: contract}, SafeMathMTransactor: SafeMathMTransactor{contract: contract}, SafeMathMFilterer: SafeMathMFilterer{contract: contract}}, nil
}

// SafeMathM is an auto generated Go binding around an Ethereum contract.
type SafeMathM struct {
	SafeMathMCaller     // Read-only binding to the contract
	SafeMathMTransactor // Write-only binding to the contract
	SafeMathMFilterer   // Log filterer for contract events
}

// SafeMathMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathMSession struct {
	Contract     *SafeMathM        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathMCallerSession struct {
	Contract *SafeMathMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeMathMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathMTransactorSession struct {
	Contract     *SafeMathMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeMathMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathMRaw struct {
	Contract *SafeMathM // Generic contract binding to access the raw methods on
}

// SafeMathMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathMCallerRaw struct {
	Contract *SafeMathMCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathMTransactorRaw struct {
	Contract *SafeMathMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathM creates a new instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathM(address common.Address, backend bind.ContractBackend) (*SafeMathM, error) {
	contract, err := bindSafeMathM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathM{SafeMathMCaller: SafeMathMCaller{contract: contract}, SafeMathMTransactor: SafeMathMTransactor{contract: contract}, SafeMathMFilterer: SafeMathMFilterer{contract: contract}}, nil
}

// NewSafeMathMCaller creates a new read-only instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMCaller(address common.Address, caller bind.ContractCaller) (*SafeMathMCaller, error) {
	contract, err := bindSafeMathM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathMCaller{contract: contract}, nil
}

// NewSafeMathMTransactor creates a new write-only instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathMTransactor, error) {
	contract, err := bindSafeMathM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathMTransactor{contract: contract}, nil
}

// NewSafeMathMFilterer creates a new log filterer instance of SafeMathM, bound to a specific deployed contract.
func NewSafeMathMFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathMFilterer, error) {
	contract, err := bindSafeMathM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathMFilterer{contract: contract}, nil
}

// bindSafeMathM binds a generic wrapper to an already deployed contract.
func bindSafeMathM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathM *SafeMathMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathM.Contract.SafeMathMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathM *SafeMathMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathM.Contract.SafeMathMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathM *SafeMathMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathM.Contract.SafeMathMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathM *SafeMathMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathM *SafeMathMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathM *SafeMathMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathM.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StandardTokenFuncSigs maps the 4-byte function signature to its string representation.
var StandardTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"66188463": "decreaseApproval(address,uint256)",
	"d73dd623": "increaseApproval(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
var StandardTokenBin = "0x60806040526a52b7d2dcc80cd2e400000060005534801561001f57600080fd5b506106708061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a9059cbb1461016f578063d73dd6231461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063661884631461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d561025b565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610261565b6100b96004803603604081101561013357600080fd5b506001600160a01b03813516906020013561037f565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b031661046f565b6100b96004803603604081101561018557600080fd5b506001600160a01b03813516906020013561048a565b6100b9600480360360408110156101b157600080fd5b506001600160a01b03813516906020013561054f565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b03813581169160200135166105e8565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661027657600080fd5b6001600160a01b038416600081815260026020908152604080832033845282528083205493835260019091529020546102b5908463ffffffff61061316565b6001600160a01b0380871660009081526001602052604080822093909355908616815220546102ea908463ffffffff61062516565b6001600160a01b038516600090815260016020526040902055610313818463ffffffff61061316565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054808311156103d4573360009081526002602090815260408083206001600160a01b0388168452909152812055610409565b6103e4818463ffffffff61061316565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b60006001600160a01b03831661049f57600080fd5b336000908152600160205260409020546104bf908363ffffffff61061316565b33600090815260016020526040808220929092556001600160a01b038516815220546104f1908363ffffffff61062516565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610583908363ffffffff61062516565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561061f57fe5b50900390565b60008282018381101561063457fe5b939250505056fea265627a7a72315820fbf024da5c4cb7201e2105ebd678f8ccd2c9c18af1c25206a9aa7a6716a5f73864736f6c63430005100032"

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
var UtilitiesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820172e345f65a42a054f2e5116eb7fb7d39de6cef4e9fc5c12b39d3ba9a9c7510b64736f6c63430005100032"

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"miniVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"authorizedUser\",\"type\":\"address\"}],\"name\":\"lockSmith\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zapToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = map[string]string{
	"47e7ef24": "deposit(address,uint256)",
	"1709ef07": "hasAccess(address,address)",
	"cd8f942c": "increaseApproval()",
	"7a47328b": "lockSmith(address,address)",
	"0103c92b": "userBalance(address)",
	"f3fef3a3": "withdraw(address,uint256)",
	"2b56eb56": "zapToken()",
}

// VaultBin is the compiled bytecode used for deploying new contracts.
var VaultBin = "0x608060405234801561001057600080fd5b50604051610a0d380380610a0d8339818101604052604081101561003357600080fd5b508051602091820151600080546001600160a01b038085166001600160a01b0319928316811790935560018054918516919092168117909155604080516024810192909252600019604480840191909152815180840390910181526064909201815294810180516001600160e01b031663095ea7b360e01b178152945181519495939492939192909182918083835b602083106100e15780518252601f1990920191602091820191016100c2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610143576040519150601f19603f3d011682016040523d82523d6000602084013e610148565b606091505b50505050506108b18061015c6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806347e7ef241161005b57806347e7ef24146101205780637a47328b1461014e578063cd8f942c1461017c578063f3fef3a3146101845761007d565b80630103c92b146100825780631709ef07146100ba5780632b56eb56146100fc575b600080fd5b6100a86004803603602081101561009857600080fd5b50356001600160a01b03166101b0565b60408051918252519081900360200190f35b6100e8600480360360408110156100d057600080fd5b506001600160a01b03813581169160200135166101cb565b604080519115158252519081900360200190f35b61010461024f565b604080516001600160a01b039092168252519081900360200190f35b61014c6004803603604081101561013657600080fd5b506001600160a01b03813516906020013561025e565b005b61014c6004803603604081101561016457600080fd5b506001600160a01b0381358116916020013516610331565b6100e861046b565b61014c6004803603604081101561019a57600080fd5b506001600160a01b038135169060200135610671565b6001600160a01b031660009081526002602052604090205490565b6000331515806101e457506001600160a01b0382163314155b61021f5760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b506001600160a01b0380821660009081526003602090815260408083209386168352929052205460ff1692915050565b6000546001600160a01b031681565b6001600160a01b0382166102a35760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b6102ad33836101cb565b6102e85760405162461bcd60e51b815260040180806020018281038252602c815260200180610806602c913960400191505060405180910390fd5b6001600160a01b038216600090815260026020526040902054610311908263ffffffff61078116565b6001600160a01b0390921660009081526002602052604090209190915550565b336001600160a01b0383161461038e576040805162461bcd60e51b815260206004820152601a60248201527f596f7520646f206e6f74206f776e2074686973207661756c742e000000000000604482015290519081900360640190fd5b331515806103a557506001600160a01b0382163314155b6103e05760405162461bcd60e51b81526004018080602001828103825260258152602001806108326025913960400191505060405180910390fd5b6001600160a01b038216600090815260036020908152604080832033845290915290205460ff16610437576001600160a01b03821660009081526003602090815260408083209091529020805460ff191660011790555b6001600160a01b0391821660009081526003602090815260408083209390941682529190915220805460ff19166001179055565b60008054600154604080513060248201526001600160a01b0392831660448083019190915282518083039091018152606490910182526020810180516001600160e01b0316636eb1769f60e11b178152915181518695606095169382918083835b602083106104eb5780518252601f1990920191602091820191016104cc565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461054d576040519150601f19603f3d011682016040523d82523d6000602084013e610552565b606091505b50915091506000610576610567836000610797565b6000199063ffffffff6107f316565b60008054600154604080516001600160a01b039283166024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b031663d73dd62360e01b178152915181519697509495606095939094169390929182918083835b602083106105fd5780518252601f1990920191602091820191016105de565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461065f576040519150601f19603f3d011682016040523d82523d6000602084013e610664565b606091505b5090965050505050505090565b6001600160a01b0382166106b65760405162461bcd60e51b81526004018080602001828103825260268152602001806108576026913960400191505060405180910390fd5b6106c033836101cb565b6106fb5760405162461bcd60e51b815260040180806020018281038252602c815260200180610806602c913960400191505060405180910390fd5b80610705836101b0565b1015610758576040805162461bcd60e51b815260206004820152601d60248201527f596f75722062616c616e636520697320696e73756666696369656e742e000000604482015290519081900360640190fd5b6001600160a01b038216600090815260026020526040902054610311908263ffffffff6107f316565b60008282018381101561079057fe5b9392505050565b600081602001835110156107ea576040805162461bcd60e51b8152602060048201526015602482015274746f55696e743235365f6f75744f66426f756e647360581b604482015290519081900360640190fd5b50016020015190565b6000828211156107ff57fe5b5090039056fe596f7520617265206e6f7420617574686f72697a656420746f206163636573732074686973207661756c742e546865207a65726f20616464726573732063616e206e6f74206f776e2061207661756c742e546865207a65726f206164647265737320646f6573206e6f74206f776e2061207661756c742ea265627a7a723158206176e431286868f4c78cf3dc1f98c589f24b719a70167ddb916c7b94b6cf113664736f6c63430005100032"

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, master common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, token, master)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultCaller) HasAccess(opts *bind.CallOpts, user common.Address, miniVault common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "hasAccess", user, miniVault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultSession) HasAccess(user common.Address, miniVault common.Address) (bool, error) {
	return _Vault.Contract.HasAccess(&_Vault.CallOpts, user, miniVault)
}

// HasAccess is a free data retrieval call binding the contract method 0x1709ef07.
//
// Solidity: function hasAccess(address user, address miniVault) view returns(bool)
func (_Vault *VaultCallerSession) HasAccess(user common.Address, miniVault common.Address) (bool, error) {
	return _Vault.Contract.HasAccess(&_Vault.CallOpts, user, miniVault)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultCaller) UserBalance(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "userBalance", userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultSession) UserBalance(userAddress common.Address) (*big.Int, error) {
	return _Vault.Contract.UserBalance(&_Vault.CallOpts, userAddress)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address userAddress) view returns(uint256 balance)
func (_Vault *VaultCallerSession) UserBalance(userAddress common.Address) (*big.Int, error) {
	return _Vault.Contract.UserBalance(&_Vault.CallOpts, userAddress)
}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultCaller) ZapToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "zapToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultSession) ZapToken() (common.Address, error) {
	return _Vault.Contract.ZapToken(&_Vault.CallOpts)
}

// ZapToken is a free data retrieval call binding the contract method 0x2b56eb56.
//
// Solidity: function zapToken() view returns(address)
func (_Vault *VaultCallerSession) ZapToken() (common.Address, error) {
	return _Vault.Contract.ZapToken(&_Vault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", userAddress, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultSession) Deposit(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, userAddress, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactorSession) Deposit(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, userAddress, value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultTransactor) IncreaseApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "increaseApproval")
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultSession) IncreaseApproval() (*types.Transaction, error) {
	return _Vault.Contract.IncreaseApproval(&_Vault.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xcd8f942c.
//
// Solidity: function increaseApproval() returns(bool)
func (_Vault *VaultTransactorSession) IncreaseApproval() (*types.Transaction, error) {
	return _Vault.Contract.IncreaseApproval(&_Vault.TransactOpts)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultTransactor) LockSmith(opts *bind.TransactOpts, miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "lockSmith", miniVault, authorizedUser)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultSession) LockSmith(miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LockSmith(&_Vault.TransactOpts, miniVault, authorizedUser)
}

// LockSmith is a paid mutator transaction binding the contract method 0x7a47328b.
//
// Solidity: function lockSmith(address miniVault, address authorizedUser) returns()
func (_Vault *VaultTransactorSession) LockSmith(miniVault common.Address, authorizedUser common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LockSmith(&_Vault.TransactOpts, miniVault, authorizedUser)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", userAddress, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultSession) Withdraw(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, userAddress, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address userAddress, uint256 value) returns()
func (_Vault *VaultTransactorSession) Withdraw(userAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, userAddress, value)
}

// ZapABI is the input ABI used to generate the binding from.
const ZapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"increaseVaultApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewZapAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_c_sapi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_c_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractZapTokenBSC\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZapFuncSigs maps the 4-byte function signature to its string representation.
var ZapFuncSigs = map[string]string{
	"752d49a1": "addTip(uint256,uint256)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"8581af19": "beginDispute(uint256,uint256,uint256)",
	"0d2d76a2": "depositStake()",
	"4049f198": "getNewCurrentVariables()",
	"7cf26110": "increaseVaultApproval(address)",
	"8da5cb5b": "owner()",
	"26b7d9f6": "proposeFork(address)",
	"3fff2816": "requestData(string,string,uint256,uint256)",
	"28449c3a": "requestStakingWithdraw()",
	"68c180d5": "submitMiningSolution(string,uint256,uint256)",
	"4d318b0e": "tallyVotes(uint256)",
	"fc0c546a": "token()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"c9d27afe": "vote(uint256,bool)",
	"bed9d861": "withdrawStake()",
}

// ZapBin is the compiled bytecode used for deploying new contracts.
var ZapBin = "0x608060405234801561001057600080fd5b506040516127bf3803806127bf8339818101604052602081101561003357600080fd5b5051604b80546001600160a01b039092166001600160a01b0319928316179055604c8054909116331790556127528061006d6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a25780638da5cb5b116100715780638da5cb5b14610423578063a9059cbb14610447578063bed9d86114610473578063c9d27afe1461047b578063fc0c546a146104a057610116565b806370a0823114610379578063752d49a1146103b15780637cf26110146103d45780638581af19146103fa57610116565b806328449c3a116100e957806328449c3a146101c15780633fff2816146101c95780634049f198146102915780634d318b0e146102e657806368c180d51461030357610116565b8063095ea7b31461011b5780630d2d76a21461015b57806323b872dd1461016557806326b7d9f61461019b575b600080fd5b6101476004803603604081101561013157600080fd5b506001600160a01b0381351690602001356104a8565b604080519115158252519081900360200190f35b610163610534565b005b6101476004803603606081101561017b57600080fd5b506001600160a01b038135811691602081013590911690604001356107e1565b610163600480360360208110156101b157600080fd5b50356001600160a01b0316610876565b610163610998565b610163600480360360808110156101df57600080fd5b8101906020810181356401000000008111156101fa57600080fd5b82018360208201111561020c57600080fd5b8035906020019184600183028401116401000000008311171561022e57600080fd5b91939092909160208101903564010000000081111561024c57600080fd5b82018360208201111561025e57600080fd5b8035906020019184600183028401116401000000008311171561028057600080fd5b919350915080359060200135610a02565b610299610ec7565b604051848152602081018460a080838360005b838110156102c45781810151838201526020016102ac565b5050505090500183815260200182815260200194505050505060405180910390f35b610163600480360360208110156102fc57600080fd5b5035610eec565b6101636004803603606081101561031957600080fd5b81019060208101813564010000000081111561033457600080fd5b82018360208201111561034657600080fd5b8035906020019184600183028401116401000000008311171561036857600080fd5b919350915080359060200135611025565b61039f6004803603602081101561038f57600080fd5b50356001600160a01b0316611521565b60408051918252519081900360200190f35b610163600480360360408110156103c757600080fd5b50803590602001356115a4565b610147600480360360208110156103ea57600080fd5b50356001600160a01b03166116bf565b6101636004803603606081101561041057600080fd5b5080359060208101359060400135611700565b61042b611b90565b604080516001600160a01b039092168252519081900360200190f35b6101476004803603604081101561045d57600080fd5b506001600160a01b038135169060200135611b9f565b610163611bf8565b6101636004803603604081101561049157600080fd5b50803590602001351515611da8565b61042b611e9e565b604b546040805163095ea7b360e01b81526001600160a01b038581166004830152602482018590529151600093929092169163095ea7b39160448082019260209290919082900301818787803b15801561050157600080fd5b505af1158015610515573d6000803e3d6000fd5b505050506040513d602081101561052b57600080fd5b50519392505050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b018120600090815260208381529083902054604b546370a0823160e01b84523360048501529351909384936001600160a01b03909116926370a0823192602480840193829003018186803b1580156105a857600080fd5b505afa1580156105bc573d6000803e3d6000fd5b505050506040513d60208110156105d257600080fd5b505110156105df57600080fd5b6040805163326991a360e01b8152600060048201819052915173__$ceb17272379ced45263b3cfc641f24c214$__9263326991a39260248082019391829003018186803b15801561062f57600080fd5b505af4158015610643573d6000803e3d6000fd5b5050604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f602090815283822054604b5463095ea7b360e01b85523060048601526024850189905294516001600160a01b03918216975087965094169363095ea7b393604480820194918390030190829087803b1580156106c257600080fd5b505af11580156106d6573d6000803e3d6000fd5b505050506040513d60208110156106ec57600080fd5b5050604b54604080516323b872dd60e01b81523360048201526001600160a01b03858116602483015260448201879052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561074a57600080fd5b505af115801561075e573d6000803e3d6000fd5b505050506040513d602081101561077457600080fd5b5050604080516311f9fbc960e21b81523360048201526024810185905290516001600160a01b038316916347e7ef2491604480830192600092919082900301818387803b1580156107c457600080fd5b505af11580156107d8573d6000803e3d6000fd5b50505050505050565b604b54604080516323b872dd60e01b81526001600160a01b038681166004830152858116602483015260448201859052915160009392909216916323b872dd9160648082019260209290919082900301818787803b15801561084257600080fd5b505af1158015610856573d6000803e3d6000fd5b505050506040513d602081101561086c57600080fd5b5051949350505050565b6040805163804b893160e01b81526000600482018190526001600160a01b0384166024830152915173__$78be728ae59efcf7bc760a8cdd86b21178$__9263804b89319260448082019391829003018186803b1580156108d557600080fd5b505af41580156108e9573d6000803e3d6000fd5b5050604b5460408051696469737075746546656560b01b8152815190819003600a01812060009081526020838152838220546323b872dd60e01b8452336004850152306024850152604484015292516001600160a01b0390941695506323b872dd9450606480830194928390030190829087803b15801561096957600080fd5b505af115801561097d573d6000803e3d6000fd5b505050506040513d602081101561099357600080fd5b505050565b60408051633c73482760e01b8152600060048201819052915173__$ceb17272379ced45263b3cfc641f24c214$__92633c7348279260248082019391829003018186803b1580156109e857600080fd5b505af41580156109fc573d6000803e3d6000fd5b50505050565b60008211610a0f57600080fd5b670de0b6b3a7640000821115610a2457600080fd5b606086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8a0181900481028201810190925288815293945060609392508891508790819084018382808284376000920191909152505084519293505050610aa457600080fd5b6040815110610ab257600080fd5b600082856040516020018083805190602001908083835b60208310610ae85780518252601f199092019160209182019101610ac9565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201815283519382019390932060008181526049909252929020549193505015159050610ea357604080516b1c995c5d595cdd10dbdd5b9d60a21b8082528251600c9281900383018120600090815260208581528582208054600101905592825284519182900390930181208352838252838320546080820185528782528183018790528185018690528451848152808401865260608301528084526048835293909220825180519192610bcb928492909101906125e6565b506020828101518051610be492600185019201906125e6565b506040820151600282015560608201518051610c0a916003840191602090910190612664565b505050600081815260486020818152604080842081516a6772616e756c617269747960a81b8152825190819003600b018120865260049091018084528286208c90558686528484526f3932b8bab2b9ba28a837b9b4ba34b7b760811b825282519182900360100182208652808452828620869055868652938352670746f74616c5469760c41b8152815190819003600801902084529181528183208390558483526049905290208190558415610d4157604b54604080516323b872dd60e01b81523360048201523060248201526044810188905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b158015610d1457600080fd5b505af1158015610d28573d6000803e3d6000fd5b505050506040513d6020811015610d3e57600080fd5b50505b610d4b8186611ead565b600081815260486020908152604091829020825192830189905260608301889052608080845281546002600180831615610100026000190190921604918501829052859433947f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc99493928401928d928d929091829182019060a083019088908015610e175780601f10610dec57610100808354040283529160200191610e17565b820191906000526020600020905b815481529060010190602001808311610dfa57829003601f168201915b5050838103825286546002600019610100600184161502019091160480825260209091019087908015610e8b5780601f10610e6057610100808354040283529160200191610e8b565b820191906000526020600020905b815481529060010190602001808311610e6e57829003601f168201915b5050965050505050505060405180910390a350610ebc565b600081815260496020526040902054610ebc90856115a4565b505050505050505050565b6000610ed161269e565b600080610ede6000612460565b935093509350935090919293565b60008060008073__$78be728ae59efcf7bc760a8cdd86b21178$__633cedafe59091866040518363ffffffff1660e01b8152600401808381526020018281526020019250505060606040518083038186803b158015610f4a57600080fd5b505af4158015610f5e573d6000803e3d6000fd5b505050506040513d6060811015610f7457600080fd5b50805160208201516040909201519094509092509050610f9483826104a8565b50604b54604080516323b872dd60e01b81526001600160a01b038681166004830152858116602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610ff357600080fd5b505af1158015611007573d6000803e3d6000fd5b505050506040513d602081101561101d57600080fd5b505050505050565b600073__$42ed2a11598b918ff241a0199a0775a312$__63d723552b9091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b1580156110bb57600080fd5b505af41580156110cf573d6000803e3d6000fd5b505050506110db6126bc565b6040805160a08101909152603560056000835b8282101561112f5760408051808201909152600283028501805482526001908101546001600160a01b031660208084019190915291835290920191016110ee565b5050604080516517dd985d5b1d60d21b815281519081900360060181206000908152603f6020908152838220547118dd5c9c995b9d135a5b995c94995dd85c9960721b8452845193849003601201909320825283905291909120549394506001600160a01b031692839250905080156114e25760005b60058110156113705760008582600581106111bc57fe5b6020020151602001516001600160a01b03161461136857604b546040805163095ea7b360e01b81523060048201526024810185905290516001600160a01b039092169163095ea7b3916044808201926020929091908290030181600087803b15801561122757600080fd5b505af115801561123b573d6000803e3d6000fd5b505050506040513d602081101561125157600080fd5b5050604b54604080516323b872dd60e01b81523060048201526001600160a01b03868116602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b1580156112af57600080fd5b505af11580156112c3573d6000803e3d6000fd5b505050506040513d60208110156112d957600080fd5b50506001600160a01b0383166347e7ef248683600581106112f657fe5b602002015160200151846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b15801561134f57600080fd5b505af1158015611363573d6000803e3d6000fd5b505050505b6001016111a5565b50604b546040805167646576536861726560c01b81528151908190036008018120600090815260208381528382205463095ea7b360e01b8452306004850152602484015292516001600160a01b039094169363095ea7b3936044808501949193918390030190829087803b1580156113e757600080fd5b505af11580156113fb573d6000803e3d6000fd5b505050506040513d602081101561141157600080fd5b5050604b5460408051652fb7bbb732b960d11b815281519081900360060181206000908152603f60209081528382205467646576536861726560c01b845284519384900360080184208352848252848320546323b872dd60e01b85523060048601526001600160a01b0391821660248601526044850152935193909416936323b872dd93606480850194929391928390030190829087803b1580156114b557600080fd5b505af11580156114c9573d6000803e3d6000fd5b505050506040513d60208110156114df57600080fd5b50505b5050604080517118dd5c9c995b9d135a5b995c94995dd85c9960721b815281519081900360120190206000908152602082905290812055505050505050565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b15801561157257600080fd5b505afa158015611586573d6000803e3d6000fd5b505050506040513d602081101561159c57600080fd5b505192915050565b600082116115b157600080fd5b801561163e57604b54604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561161157600080fd5b505af1158015611625573d6000803e3d6000fd5b505050506040513d602081101561163b57600080fd5b50505b6116488282611ead565b60008281526048602090815260408083208151670746f74616c5469760c41b815282519081900360080181208552600490910183529281902054848452918301919091528051849233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e882092918290030190a35050565b600080829050806001600160a01b031663cd8f942c6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561050157600080fd5b6000838152604860209081526040808320858452600581019092529091205460904391909103111561173157600080fd5b600083815260058201602052604090205461174b57600080fd5b6005821061175857600080fd5b60008381526008820160205260408120836005811061177357fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091206000818152604a909252919020546001600160a01b03909216925090156117e557600080fd5b60408051696469737075746546656560b01b8152815190819003600a01902060009081526020829052205461181d90339030906107e1565b50604080516b191a5cdc1d5d1950dbdd5b9d60a21b808252825191829003600c90810183206000908152602085815285822054848652865195869003840186208352868252868320600191820190559385528551948590039092018420815284825284812054868252604a8352858220819055610100808601875287865285840183815286880184815260608801858152608089018681526001600160a01b03808e1660a08c019081523360c08d0190815260e08d018a8152898b526044808d528f8c209e518f5597519c8e019c909c55945160028d01805495519451925184166301000000026301000000600160b81b0319931515620100000262ff000019961515909a0261ff001993151560ff19909816979097179290921695909517939093169690961795909516179055516003880180549184166001600160a01b031992831617905595516004880180549190931696169590951790558551681c995c5d595cdd125960ba1b815286519081900360099081018220845260059687018086528885208f905583855295855268074696d657374616d760bc1b8252875191829003810190912083529383528582208b90558a82529288019091529290922090869081106119e957fe5b0154600082815260446020818152604080842081516476616c756560d81b8152825190819003600590810182208752909101808452828620969096558685528383526f6d696e457865637574696f6e4461746560801b81528151908190036010018120855285835281852062093a80420190558685528383526a313637b1b5a73ab6b132b960a91b8152815190819003600b0181208552858352818520439055868552838352681b5a5b995c94db1bdd60ba1b8152815190819003600901812085528583528185208b9055696469737075746546656560b01b8152815190819003600a0181208552818352818520548786529383526266656560e81b815281519081900360030190208452939052919020556002851415611b2c5760008781526048602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260476020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a350505050505050565b604c546001600160a01b031681565b604b546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151600093929092169163a9059cbb9160448082019260209290919082900301818787803b15801561050157600080fd5b604080516378bfa27760e01b8152600060048201819052915173__$ceb17272379ced45263b3cfc641f24c214$__926378bfa2779260248082019391829003018186803b158015611c4857600080fd5b505af4158015611c5c573d6000803e3d6000fd5b5050604080516517dd985d5b1d60d21b8082528251600692819003830181206000908152603f602081815286832054604b54958552875194859003909601842083529081529085902054630103c92b60e01b8352336004840181905295516001600160a01b039586169850889750938516956323b872dd959091169390928792630103c92b92602480840193829003018186803b158015611cfc57600080fd5b505afa158015611d10573d6000803e3d6000fd5b505050506040513d6020811015611d2657600080fd5b5051604080516001600160e01b031960e087901b1681526001600160a01b0394851660048201529290931660248301526044820152905160648083019260209291908290030181600087803b158015611d7e57600080fd5b505af1158015611d92573d6000803e3d6000fd5b505050506040513d60208110156109fc57600080fd5b604b54604080516370a0823160e01b815233600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015611df357600080fd5b505afa158015611e07573d6000803e3d6000fd5b505050506040513d6020811015611e1d57600080fd5b50516040805163f9f6af1760e01b815260006004820181905260248201879052851515604483015260648201849052915192935073__$78be728ae59efcf7bc760a8cdd86b21178$__9263f9f6af1792608480840193919291829003018186803b158015611e8a57600080fd5b505af41580156107d8573d6000803e3d6000fd5b604b546001600160a01b031681565b600082815260486020526040812090611ec58161250f565b90508215611f345760408051670746f74616c5469760c41b815281519081900360080190206000908152600484016020522054611f08908463ffffffff61256b16565b60408051670746f74616c5469760c41b8152815190819003600801902060009081526004850160205220555b60408051670746f74616c5469760c41b815281519081900360080181206000908152600485016020908152838220546f18dd5c9c995b9d14995c5d595cdd125960821b8452845193849003601001909320825283905291909120546121e35760408051670746f74616c5469760c41b815281519081900360080181206000908152600480870160209081528483208390556f18dd5c9c995b9d14995c5d595cdd125960821b8085528551948590036010908101862085528683528685208c90556f63757272656e74546f74616c5469707360801b80875287519687900382018720865287845287862089905585548785018a905287890152600019438101406060808a01919091528951808a03820181526080808b01808d52825192890192909220808b55918790528b519a8b90036090018b208a528b88528b8a205469646966666963756c747960b01b8c528c519b8c9003600a018c208b528c89528c8b2054888d528d519c8d900388018d208c528d8a528d8c20548c526048808b528e8d206a6772616e756c617269747960a81b8f528f519e8f9003600b018f208e52909b018a528d8c2054988d528d519c8d900388018d208c528d8a528d8c20548c529989528c8b20958c528c519b8c90039096018b208a528b8852988b902054818b52968a0188905299890185905296880185905260a096880187815282546002610100600183161502909301169190910496880187905291977f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42979096939491939192909160c0830190859080156121cc5780601f106121a1576101008083540402835291602001916121cc565b820191906000526020600020905b8154815290600101906020018083116121af57829003601f168201915b5050965050505050505060405180910390a2612459565b60008281526048602090815260408083208151670746f74616c5469760c41b815282519081900360080190208452600401909152902054811180612225575081155b156122f25760028084015460408051602081018390529081018490526060808252865460001961010060018316150201169390930492810183905287927f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c928792909186919081906080820190869080156122e15780601f106122b6576101008083540402835291602001916122e1565b820191906000526020600020905b8154815290600101906020018083116122c457829003601f168201915b505094505050505060405180910390a25b604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054612408576040805161066081019182905260009182916123659160019060339082845b815481526020019060010190808311612348575050505050612581565b909250905081831180612376575081155b1561240157826001826033811061238957fe5b01556000818152604360208181526040808420805485526048835281852082516f3932b8bab2b9ba28a837b9b4ba34b7b760811b80825284519182900360109081018320895260049384018752858920899055898952968652928e905591825282519182900390940190208452918801905290208190555b5050612459565b831561245957604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b81528151908190036010019020600090815260048501602052205484906001906033811061245057fe5b01805490910190555b5050505050565b600061246a61269e565b600080805b60058110156124a55785603501816005811061248757fe5b600202015484826005811061249857fe5b602002015260010161246f565b505083546040805169646966666963756c747960b01b8152815190819003600a01812060009081528288016020818152848320546f63757272656e74546f74616c5469707360801b8552855194859003601001909420835252919091205491945091509193509193565b60408051610660810191829052600091829182916125509190600187019060339082845b8154815260200190600101908083116125335750505050506125b8565b60009081526043909501602052505060409092205492915050565b60008282018381101561257a57fe5b9392505050565b610640810151603260315b80156125b25760208102840151838110156125a8578093508192505b506000190161258c565b50915091565b60008060005b60338110156125b25760208102840151808410156125dd578093508192505b506001016125be565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061262757805160ff1916838001178555612654565b82800160010185558215612654579182015b82811115612654578251825591602001919060010190612639565b506126609291506126e9565b5090565b8280548282559060005260206000209081019282156126545791602002820182811115612654578251825591602001919060010190612639565b6040518060a001604052806005906020820280388339509192915050565b6040518060a001604052806005905b6126d3612706565b8152602001906001900390816126cb5790505090565b61270391905b8082111561266057600081556001016126ef565b90565b60408051808201909152600080825260208201529056fea265627a7a72315820b7c99f3f84452d3388f87d867af33c1b7a3daeb68dd01943c689bddb4c2b767664736f6c63430005100032"

// DeployZap deploys a new Ethereum contract, binding an instance of Zap to it.
func DeployZap(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *Zap, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapLibraryAddr, _, _, _ := DeployZapLibrary(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$42ed2a11598b918ff241a0199a0775a312$__", zapLibraryAddr.String()[2:], -1)

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$78be728ae59efcf7bc760a8cdd86b21178$__", zapDisputeAddr.String()[2:], -1)

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapBin = strings.Replace(ZapBin, "__$ceb17272379ced45263b3cfc641f24c214$__", zapStakeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapBin), backend, zapTokenBsc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// Zap is an auto generated Go binding around an Ethereum contract.
type Zap struct {
	ZapCaller     // Read-only binding to the contract
	ZapTransactor // Write-only binding to the contract
	ZapFilterer   // Log filterer for contract events
}

// ZapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapSession struct {
	Contract     *Zap              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapCallerSession struct {
	Contract *ZapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransactorSession struct {
	Contract     *ZapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapRaw struct {
	Contract *Zap // Generic contract binding to access the raw methods on
}

// ZapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapCallerRaw struct {
	Contract *ZapCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransactorRaw struct {
	Contract *ZapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZap creates a new instance of Zap, bound to a specific deployed contract.
func NewZap(address common.Address, backend bind.ContractBackend) (*Zap, error) {
	contract, err := bindZap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// NewZapCaller creates a new read-only instance of Zap, bound to a specific deployed contract.
func NewZapCaller(address common.Address, caller bind.ContractCaller) (*ZapCaller, error) {
	contract, err := bindZap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapCaller{contract: contract}, nil
}

// NewZapTransactor creates a new write-only instance of Zap, bound to a specific deployed contract.
func NewZapTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransactor, error) {
	contract, err := bindZap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransactor{contract: contract}, nil
}

// NewZapFilterer creates a new log filterer instance of Zap, bound to a specific deployed contract.
func NewZapFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapFilterer, error) {
	contract, err := bindZap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapFilterer{contract: contract}, nil
}

// bindZap binds a generic wrapper to an already deployed contract.
func bindZap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.ZapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Zap.Contract.BalanceOf(&_Zap.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 balance)
func (_Zap *ZapCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Zap.Contract.BalanceOf(&_Zap.CallOpts, _user)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RequestIds = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Difficutly = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tip = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Zap *ZapCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Zap.Contract.GetNewCurrentVariables(&_Zap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapSession) Owner() (common.Address, error) {
	return _Zap.Contract.Owner(&_Zap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zap *ZapCallerSession) Owner() (common.Address, error) {
	return _Zap.Contract.Owner(&_Zap.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zap.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapSession) Token() (common.Address, error) {
	return _Zap.Contract.Token(&_Zap.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Zap *ZapCallerSession) Token() (common.Address, error) {
	return _Zap.Contract.Token(&_Zap.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Zap *ZapTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Zap *ZapTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapTransactor) IncreaseVaultApproval(opts *bind.TransactOpts, vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "increaseVaultApproval", vaultAddress)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapSession) IncreaseVaultApproval(vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts, vaultAddress)
}

// IncreaseVaultApproval is a paid mutator transaction binding the contract method 0x7cf26110.
//
// Solidity: function increaseVaultApproval(address vaultAddress) returns(bool)
func (_Zap *ZapTransactorSession) IncreaseVaultApproval(vaultAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.IncreaseVaultApproval(&_Zap.TransactOpts, vaultAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapTransactor) ProposeFork(opts *bind.TransactOpts, _propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "proposeFork", _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewZapAddress) returns()
func (_Zap *ZapTransactorSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(string _c_sapi, string _c_symbol, uint256 _granularity, uint256 _tip) returns()
func (_Zap *ZapTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(string _nonce, uint256 _requestId, uint256 _value) returns()
func (_Zap *ZapTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Zap *ZapTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Zap *ZapTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Zap *ZapTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// ZapApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Zap contract.
type ZapApprovalIterator struct {
	Event *ZapApproval // Event containing the contract specifics and raw log

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
func (it *ZapApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapApproval)
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
		it.Event = new(ZapApproval)
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
func (it *ZapApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapApproval represents a Approval event raised by the Zap contract.
type ZapApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ZapApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapApprovalIterator{contract: _Zap.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapApproval)
				if err := _Zap.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Zap *ZapFilterer) ParseApproval(log types.Log) (*ZapApproval, error) {
	event := new(ZapApproval)
	if err := _Zap.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the Zap contract.
type ZapDataRequestedIterator struct {
	Event *ZapDataRequested // Event containing the contract specifics and raw log

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
func (it *ZapDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDataRequested)
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
		it.Event = new(ZapDataRequested)
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
func (it *ZapDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDataRequested represents a DataRequested event raised by the Zap contract.
type ZapDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapDataRequestedIterator{contract: _Zap.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *ZapDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDataRequested)
				if err := _Zap.contract.UnpackLog(event, "DataRequested", log); err != nil {
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

// ParseDataRequested is a log parse operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseDataRequested(log types.Log) (*ZapDataRequested, error) {
	event := new(ZapDataRequested)
	if err := _Zap.contract.UnpackLog(event, "DataRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the Zap contract.
type ZapNewChallengeIterator struct {
	Event *ZapNewChallenge // Event containing the contract specifics and raw log

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
func (it *ZapNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewChallenge)
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
		it.Event = new(ZapNewChallenge)
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
func (it *ZapNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewChallenge represents a NewChallenge event raised by the Zap contract.
type ZapNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*ZapNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewChallengeIterator{contract: _Zap.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ZapNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewChallenge)
				if err := _Zap.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ParseNewChallenge is a log parse operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseNewChallenge(log types.Log) (*ZapNewChallenge, error) {
	event := new(ZapNewChallenge)
	if err := _Zap.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the Zap contract.
type ZapNewDisputeIterator struct {
	Event *ZapNewDispute // Event containing the contract specifics and raw log

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
func (it *ZapNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewDispute)
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
		it.Event = new(ZapNewDispute)
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
func (it *ZapNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewDispute represents a NewDispute event raised by the Zap contract.
type ZapNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ZapNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewDisputeIterator{contract: _Zap.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ZapNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewDispute)
				if err := _Zap.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Zap *ZapFilterer) ParseNewDispute(log types.Log) (*ZapNewDispute, error) {
	event := new(ZapNewDispute)
	if err := _Zap.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the Zap contract.
type ZapNewRequestOnDeckIterator struct {
	Event *ZapNewRequestOnDeck // Event containing the contract specifics and raw log

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
func (it *ZapNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapNewRequestOnDeck)
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
		it.Event = new(ZapNewRequestOnDeck)
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
func (it *ZapNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapNewRequestOnDeck represents a NewRequestOnDeck event raised by the Zap contract.
type ZapNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapNewRequestOnDeckIterator{contract: _Zap.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *ZapNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapNewRequestOnDeck)
				if err := _Zap.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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

// ParseNewRequestOnDeck is a log parse operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_Zap *ZapFilterer) ParseNewRequestOnDeck(log types.Log) (*ZapNewRequestOnDeck, error) {
	event := new(ZapNewRequestOnDeck)
	if err := _Zap.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zap contract.
type ZapOwnershipTransferredIterator struct {
	Event *ZapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapOwnershipTransferred)
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
		it.Event = new(ZapOwnershipTransferred)
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
func (it *ZapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapOwnershipTransferred represents a OwnershipTransferred event raised by the Zap contract.
type ZapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zap *ZapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapOwnershipTransferredIterator{contract: _Zap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zap *ZapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapOwnershipTransferred)
				if err := _Zap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Zap *ZapFilterer) ParseOwnershipTransferred(log types.Log) (*ZapOwnershipTransferred, error) {
	event := new(ZapOwnershipTransferred)
	if err := _Zap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the Zap contract.
type ZapTipAddedIterator struct {
	Event *ZapTipAdded // Event containing the contract specifics and raw log

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
func (it *ZapTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTipAdded)
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
		it.Event = new(ZapTipAdded)
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
func (it *ZapTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTipAdded represents a TipAdded event raised by the Zap contract.
type ZapTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapTipAddedIterator{contract: _Zap.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ZapTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Zap.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTipAdded)
				if err := _Zap.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Zap *ZapFilterer) ParseTipAdded(log types.Log) (*ZapTipAdded, error) {
	event := new(ZapTipAdded)
	if err := _Zap.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeABI is the input ABI used to generate the binding from.
const ZapDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"}]"

// ZapDisputeFuncSigs maps the 4-byte function signature to its string representation.
var ZapDisputeFuncSigs = map[string]string{
	"804b8931": "proposeFork(ZapStorage.ZapStorageStruct storage,address)",
	"3cedafe5": "tallyVotes(ZapStorage.ZapStorageStruct storage,uint256)",
	"9264b888": "updateDisputeFee(ZapStorage.ZapStorageStruct storage)",
	"f9f6af17": "vote(ZapStorage.ZapStorageStruct storage,uint256,bool,uint256)",
}

// ZapDisputeBin is the compiled bytecode used for deploying new contracts.
var ZapDisputeBin = "0x610c65610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c80633cedafe51461005b578063804b8931146100b55780639264b888146100f0578063f9f6af171461011a575b600080fd5b81801561006757600080fd5b5061008b6004803603604081101561007e57600080fd5b5080359060200135610158565b604080516001600160a01b0394851681529290931660208301528183015290519081900360600190f35b8180156100c157600080fd5b506100ee600480360360408110156100d857600080fd5b50803590602001356001600160a01b0316610592565b005b8180156100fc57600080fd5b506100ee6004803603602081101561011357600080fd5b503561092c565b81801561012657600080fd5b506100ee6004803603608081101561013d57600080fd5b50803590602081013590604081013515159060600135610ab7565b600081815260448301602090815260408083208151681c995c5d595cdd125960ba1b815282519081900360090181208552600582018085528386205486526048880185528386206266656560e81b8352845192839003600301909220865290935290832054600282015484938493929091849060ff16156101d857600080fd5b604080516f6d696e457865637574696f6e4461746560801b815281519081900360100190206000908152600586016020522054421161021657600080fd5b600284015462010000900460ff166103fc576002840154630100000090046001600160a01b0316600090815260478a0160205260408120600186015490911215610352576000815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818c016020522080546000190190556102a88a61092c565b60028501805461ff0019166101001790556040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff1615156001141561033d576040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600687019052908120555b60038501546001600160a01b031691506103f6565b60018082556040805168074696d657374616d760bc1b815281519081900360090190206000908152600588016020908152828220548252600788019052205460ff16151514156103de576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058701602090815282822054825260078701905220805460ff191690555b6002850154630100000090046001600160a01b031691505b5061050a565b60008460010154131561050a57604080516b746f74616c5f737570706c7960a01b8152815190819003600c0190206000908152818b016020522054606490601402604080516571756f72756d60d01b8152815190819003600601902060009081526005880160205220549190041061047357600080fd5b600484018054604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f8e0160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028901805461ff00191661010017905593549092168252517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270929181900390910190a15b60028401805460ff19166001908117918290558501546003860154604080519283526001600160a01b039182166020840152610100840460ff1615158382015251630100000090930416918a917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a33096509450925050509250925092565b604080516bffffffffffffffffffffffff19606084901b1660208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156105e757600080fd5b82604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002060008154809291906001019190505550600083604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505043846044016000838152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080696469737075746546656560b01b815250600a0190506040518091039020815260200190815260200160002054846044016000838152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020819055504262093a8001846044016000838152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b8152506010019050604051809103902081526020019081526020016000208190555050505050565b604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546103e8919082028161099157fe5b041015610a8457604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b019094208352529190912054610a5391600f916103e891610a4691830281610a0757fe5b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b019020600090815281890160205220549190046103e80363ffffffff610bf316565b81610a4d57fe5b04610c1a565b60408051696469737075746546656560b01b8152815190819003600a01902060009081528184016020522055610ab4565b60408051696469737075746546656560b01b8152815190819003600a019020600090815281830160205220600f90555b50565b60008381526044850160209081526040808320338452600681019092529091205460ff16151560011415610aea57600080fd5b60008211610af757600080fd5b33600090815260478601602052604090205460031415610b1657600080fd5b336000908152600680830160209081526040808420805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d01812086526005870180855283872080549093019092556571756f72756d60d01b8152825190819003909401909320845291905290208054830190558215610ba75760018101805483019055610bb3565b60018101805483900390555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b6000828202831580610c0d575082848281610c0a57fe5b04145b610c1357fe5b9392505050565b6000818311610c295781610c13565b509091905056fea265627a7a72315820617a54cc2f133639f1a6a78f52adc8b78f50146711f593fe5541cbc1c31277f464736f6c63430005100032"

// DeployZapDispute deploys a new Ethereum contract, binding an instance of ZapDispute to it.
func DeployZapDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapDisputeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// ZapDispute is an auto generated Go binding around an Ethereum contract.
type ZapDispute struct {
	ZapDisputeCaller     // Read-only binding to the contract
	ZapDisputeTransactor // Write-only binding to the contract
	ZapDisputeFilterer   // Log filterer for contract events
}

// ZapDisputeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapDisputeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapDisputeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapDisputeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapDisputeSession struct {
	Contract     *ZapDispute       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapDisputeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapDisputeCallerSession struct {
	Contract *ZapDisputeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapDisputeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapDisputeTransactorSession struct {
	Contract     *ZapDisputeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapDisputeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapDisputeRaw struct {
	Contract *ZapDispute // Generic contract binding to access the raw methods on
}

// ZapDisputeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapDisputeCallerRaw struct {
	Contract *ZapDisputeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapDisputeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapDisputeTransactorRaw struct {
	Contract *ZapDisputeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapDispute creates a new instance of ZapDispute, bound to a specific deployed contract.
func NewZapDispute(address common.Address, backend bind.ContractBackend) (*ZapDispute, error) {
	contract, err := bindZapDispute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// NewZapDisputeCaller creates a new read-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeCaller(address common.Address, caller bind.ContractCaller) (*ZapDisputeCaller, error) {
	contract, err := bindZapDispute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeCaller{contract: contract}, nil
}

// NewZapDisputeTransactor creates a new write-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapDisputeTransactor, error) {
	contract, err := bindZapDispute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeTransactor{contract: contract}, nil
}

// NewZapDisputeFilterer creates a new log filterer instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapDisputeFilterer, error) {
	contract, err := bindZapDispute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeFilterer{contract: contract}, nil
}

// bindZapDispute binds a generic wrapper to an already deployed contract.
func bindZapDispute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.ZapDisputeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transact(opts, method, params...)
}

// ZapDisputeDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTalliedIterator struct {
	Event *ZapDisputeDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ZapDisputeDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeDisputeVoteTallied)
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
		it.Event = new(ZapDisputeDisputeVoteTallied)
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
func (it *ZapDisputeDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeDisputeVoteTallied represents a DisputeVoteTallied event raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ZapDisputeDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeDisputeVoteTalliedIterator{contract: _ZapDispute.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ZapDisputeDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeDisputeVoteTallied)
				if err := _ZapDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// ParseDisputeVoteTallied is a log parse operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _active)
func (_ZapDispute *ZapDisputeFilterer) ParseDisputeVoteTallied(log types.Log) (*ZapDisputeDisputeVoteTallied, error) {
	event := new(ZapDisputeDisputeVoteTallied)
	if err := _ZapDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the ZapDispute contract.
type ZapDisputeNewDisputeIterator struct {
	Event *ZapDisputeNewDispute // Event containing the contract specifics and raw log

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
func (it *ZapDisputeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewDispute)
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
		it.Event = new(ZapDisputeNewDispute)
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
func (it *ZapDisputeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewDispute represents a NewDispute event raised by the ZapDispute contract.
type ZapDisputeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ZapDisputeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewDisputeIterator{contract: _ZapDispute.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewDispute)
				if err := _ZapDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ZapDispute *ZapDisputeFilterer) ParseNewDispute(log types.Log) (*ZapDisputeNewDispute, error) {
	event := new(ZapDisputeNewDispute)
	if err := _ZapDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapDispute contract.
type ZapDisputeNewZapAddressIterator struct {
	Event *ZapDisputeNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapDisputeNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewZapAddress)
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
		it.Event = new(ZapDisputeNewZapAddress)
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
func (it *ZapDisputeNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewZapAddress represents a NewZapAddress event raised by the ZapDispute contract.
type ZapDisputeNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapDisputeNewZapAddressIterator, error) {

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewZapAddressIterator{contract: _ZapDispute.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewZapAddress)
				if err := _ZapDispute.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapDispute *ZapDisputeFilterer) ParseNewZapAddress(log types.Log) (*ZapDisputeNewZapAddress, error) {
	event := new(ZapDisputeNewZapAddress)
	if err := _ZapDispute.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapDisputeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ZapDispute contract.
type ZapDisputeVotedIterator struct {
	Event *ZapDisputeVoted // Event containing the contract specifics and raw log

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
func (it *ZapDisputeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeVoted)
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
		it.Event = new(ZapDisputeVoted)
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
func (it *ZapDisputeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeVoted represents a Voted event raised by the ZapDispute contract.
type ZapDisputeVoted struct {
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address) (*ZapDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeVotedIterator{contract: _ZapDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ZapDisputeVoted, _disputeID []*big.Int, _voter []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeVoted)
				if err := _ZapDispute.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter)
func (_ZapDispute *ZapDisputeFilterer) ParseVoted(log types.Log) (*ZapDisputeVoted, error) {
	event := new(ZapDisputeVoted)
	if err := _ZapDispute.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapGettersABI is the input ABI used to generate the binding from.
const ZapGettersABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zapTokenBsc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZapGettersFuncSigs maps the 4-byte function signature to its string representation.
var ZapGettersFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"70a08231": "balanceOf(address)",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"a22e407a": "getCurrentVariables()",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"c775b542": "getMinedBlockNum(uint256,uint256)",
	"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
	"17d7de7c": "getName()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"15070401": "getSymbol()",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"1ca8b6cb": "totalTokenSupply()",
}

// ZapGettersBin is the compiled bytecode used for deploying new contracts.
var ZapGettersBin = "0x608060405234801561001057600080fd5b506040516119233803806119238339818101604052602081101561003357600080fd5b5051604b80546001600160a01b0319166001600160a01b039092169190911790556118c0806100636000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806370a0823111610104578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e1461078c578063e0ae93c1146107ba578063e1eee6d6146107dd578063fc7cf0a0146108f4576101da565b8063af0b13271461068a578063b54130291461072e578063c775b5421461074c578063da3799411461076f576101da565b80637f6fd5d9116100de5780637f6fd5d91461057557806393fa491514610598578063a22e407a146105bb578063a7c438bc1461065e576101da565b806370a08231146104ed578063733bdef01461051357806377fbb66314610552576101da565b80631db842f01161017c578063612c8f7f1161014b578063612c8f7f146104645780636173c0b81461048157806363bb82ad1461049e57806369026d63146104ca576101da565b80631db842f0146103bd5780633180f8df146103da5780633df0777b1461041057806346eee1c414610447576101da565b806315070401116101b857806315070401146102a257806317d7de7c1461031f57806319e8e03b146103275780631ca8b6cb146103b5576101da565b80630f0b424d146101df57806311c985121461020e578063133bee5e14610269575b600080fd5b6101fc600480360360208110156101f557600080fd5b50356108fc565b60408051918252519081900360200190f35b6102316004803603604081101561022457600080fd5b5080359060200135610914565b604051808260a080838360005b8381101561025657818101518382015260200161023e565b5050505090500191505060405180910390f35b6102866004803603602081101561027f57600080fd5b5035610935565b604080516001600160a01b039092168252519081900360200190f35b6102aa610947565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102e45781810151838201526020016102cc565b50505050905090810190601f1680156103115780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102aa610958565b61032f610964565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610378578181015183820152602001610360565b50505050905090810190601f1680156103a55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6101fc61097e565b6101fc600480360360208110156103d357600080fd5b503561098a565b6103f7600480360360208110156103f057600080fd5b503561099c565b6040805192835290151560208301528051918290030190f35b6104336004803603604081101561042657600080fd5b50803590602001356109b8565b604080519115158252519081900360200190f35b6101fc6004803603602081101561045d57600080fd5b50356109cb565b6101fc6004803603602081101561047a57600080fd5b50356109dd565b6101fc6004803603602081101561049757600080fd5b50356109ef565b610433600480360360408110156104b457600080fd5b50803590602001356001600160a01b0316610a01565b610231600480360360408110156104e057600080fd5b5080359060200135610a14565b6101fc6004803603602081101561050357600080fd5b50356001600160a01b0316610a2e565b6105396004803603602081101561052957600080fd5b50356001600160a01b0316610ab1565b6040805192835260208301919091528051918290030190f35b6101fc6004803603604081101561056857600080fd5b5080359060200135610ac4565b6101fc6004803603604081101561058b57600080fd5b5080359060200135610ad7565b6101fc600480360360408110156105ae57600080fd5b5080359060200135610aea565b6105c3610afd565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561061e578181015183820152602001610606565b50505050905090810190601f16801561064b5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6104336004803603604081101561067457600080fd5b50803590602001356001600160a01b0316610b24565b6106a7600480360360208110156106a057600080fd5b5035610b37565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561070d5781810151838201526020016106f5565b50505050905001828152602001995050505050505050505060405180910390f35b610736610b7b565b604051815181528082610660808383602061023e565b6101fc6004803603604081101561076257600080fd5b5080359060200135610b8d565b6101fc6004803603602081101561078557600080fd5b5035610ba0565b6101fc600480360360408110156107a257600080fd5b506001600160a01b0381358116916020013516610bb2565b6101fc600480360360408110156107d057600080fd5b5080359060200135610c3e565b6107fa600480360360208110156107f357600080fd5b5035610c51565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b8381101561085357818101518382015260200161083b565b50505050905090810190601f1680156108805780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b838110156108b357818101518382015260200161089b565b50505050905090810190601f1680156108e05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6103f7610c7d565b600061090e818363ffffffff610c9216565b92915050565b61091c61182f565b61092e6000848463ffffffff610ca816565b9392505050565b600061090e818363ffffffff610d0216565b60606109536000610d21565b905090565b60606109536000610d40565b60008060606109736000610d64565b925092509250909192565b60006109536000610e50565b600061090e818363ffffffff610e8316565b6000806109af818463ffffffff610e9916565b91509150915091565b600061092e81848463ffffffff610eff16565b600061090e818363ffffffff610f2616565b600061090e818363ffffffff610f3f16565b600061090e818363ffffffff610f5116565b600061092e81848463ffffffff610f7816565b610a1c61182f565b61092e6000848463ffffffff610fa516565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610a7f57600080fd5b505afa158015610a93573d6000803e3d6000fd5b505050506040513d6020811015610aa957600080fd5b505192915050565b6000806109af818463ffffffff61100916565b600061092e81848463ffffffff61103016565b600061092e81848463ffffffff61106316565b600061092e81848463ffffffff61108716565b60008060006060600080610b1160006110ab565b949b939a50919850965094509092509050565b600061092e81848463ffffffff61126916565b6000806000806000806000610b4a61184d565b6000610b5c818b63ffffffff61129b16565b9850985098509850985098509850985098509193959799909294969850565b610b8361186c565b61095360006114b0565b600061092e81848463ffffffff6114f016565b600061090e818363ffffffff61151416565b604b5460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b158015610c0b57600080fd5b505afa158015610c1f573d6000803e3d6000fd5b505050506040513d6020811015610c3557600080fd5b50519392505050565b600061092e81848463ffffffff61152a16565b6060806000808080610c69818863ffffffff61154e16565b949c939b5091995097509550909350915050565b600080610c8a6000611729565b915091509091565b6000908152604291909101602052604090205490565b610cb061182f565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610ce157505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b506040805180820190915260048152632d20a82160e11b602082015290565b5060408051808201909152600981526805a61702042455032360bc1b602082015290565b60008060606000610d748561179f565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610e3b5780601f10610e1057610100808354040283529160200191610e3b565b820191906000526020600020905b815481529060010190602001808311610e1e57829003601f168201915b50505050509050935093509350509193909250565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b6000908152604991909101602052604090205490565b60008181526048830160205260408120600381015482919015610eef57600381018054610ee39187918791906000198101908110610ed357fe5b9060005260206000200154611087565b60019250925050610ef8565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b60006032821115610f6157600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b610fad61182f565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b03168152600190910190602001808311610fde57505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061104f57fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a988998949795969095949193919291859183018282801561124d5780601f106112225761010080835404028352916020019161124d565b820191906000526020600020905b81548152906001019060200180831161123057829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60008060008060008060006112ae61184d565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b6114b861186c565b6040805161066081019182905290600184019060339082845b8154815260200190600101908083116114d15750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a9990988998899889989197889793880196929592949091889183018282801561167d5780601f106116525761010080835404028352916020019161167d565b820191906000526020600020905b81548152906001019060200180831161166057829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a94509250840190508282801561170b5780601f106116e05761010080835404028352916020019161170b565b820191906000526020600020905b8154815290600101906020018083116116ee57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611795918591611087565b9360019350915050565b60408051610660810191829052600091829182916117e09190600187019060339082845b8154815260200190600101908083116117c35750505050506117fb565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611829576020810284015180841015611820578093508192505b50600101611801565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a7231582049ef0739b488fbf87d0240538b78a0c31233af31d81006239bec62260f95ae5264736f6c63430005100032"

// DeployZapGetters deploys a new Ethereum contract, binding an instance of ZapGetters to it.
func DeployZapGetters(auth *bind.TransactOpts, backend bind.ContractBackend, zapTokenBsc common.Address) (common.Address, *types.Transaction, *ZapGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersBin), backend, zapTokenBsc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// ZapGetters is an auto generated Go binding around an Ethereum contract.
type ZapGetters struct {
	ZapGettersCaller     // Read-only binding to the contract
	ZapGettersTransactor // Write-only binding to the contract
	ZapGettersFilterer   // Log filterer for contract events
}

// ZapGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersSession struct {
	Contract     *ZapGetters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersCallerSession struct {
	Contract *ZapGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersTransactorSession struct {
	Contract     *ZapGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersRaw struct {
	Contract *ZapGetters // Generic contract binding to access the raw methods on
}

// ZapGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersCallerRaw struct {
	Contract *ZapGettersCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersTransactorRaw struct {
	Contract *ZapGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGetters creates a new instance of ZapGetters, bound to a specific deployed contract.
func NewZapGetters(address common.Address, backend bind.ContractBackend) (*ZapGetters, error) {
	contract, err := bindZapGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// NewZapGettersCaller creates a new read-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersCaller, error) {
	contract, err := bindZapGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersCaller{contract: contract}, nil
}

// NewZapGettersTransactor creates a new write-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersTransactor, error) {
	contract, err := bindZapGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersTransactor{contract: contract}, nil
}

// NewZapGettersFilterer creates a new log filterer instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersFilterer, error) {
	contract, err := bindZapGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersFilterer{contract: contract}, nil
}

// bindZapGetters binds a generic wrapper to an already deployed contract.
func bindZapGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.ZapGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapGetters *ZapGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getCurrentVariables")

	if err != nil {
		return *new([32]byte), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapGetters *ZapGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapGetters *ZapGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapGetters *ZapGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "getVariablesOnDeck")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapGetters *ZapGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersCaller) TotalTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZapGetters.contract.Call(opts, &out, "totalTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalTokenSupply(&_ZapGetters.CallOpts)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalTokenSupply(&_ZapGetters.CallOpts)
}

// ZapGettersLibraryABI is the input ABI used to generate the binding from.
const ZapGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]"

// ZapGettersLibraryBin is the compiled bytecode used for deploying new contracts.
var ZapGettersLibraryBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820dbafeb24e1139a9c278e338f7293e0f06d0fc04da18189bca871b491089552a364736f6c63430005100032"

// DeployZapGettersLibrary deploys a new Ethereum contract, binding an instance of ZapGettersLibrary to it.
func DeployZapGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapGettersLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// ZapGettersLibrary is an auto generated Go binding around an Ethereum contract.
type ZapGettersLibrary struct {
	ZapGettersLibraryCaller     // Read-only binding to the contract
	ZapGettersLibraryTransactor // Write-only binding to the contract
	ZapGettersLibraryFilterer   // Log filterer for contract events
}

// ZapGettersLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersLibrarySession struct {
	Contract     *ZapGettersLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZapGettersLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersLibraryCallerSession struct {
	Contract *ZapGettersLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZapGettersLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersLibraryTransactorSession struct {
	Contract     *ZapGettersLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZapGettersLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersLibraryRaw struct {
	Contract *ZapGettersLibrary // Generic contract binding to access the raw methods on
}

// ZapGettersLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCallerRaw struct {
	Contract *ZapGettersLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactorRaw struct {
	Contract *ZapGettersLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGettersLibrary creates a new instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibrary(address common.Address, backend bind.ContractBackend) (*ZapGettersLibrary, error) {
	contract, err := bindZapGettersLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// NewZapGettersLibraryCaller creates a new read-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersLibraryCaller, error) {
	contract, err := bindZapGettersLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryCaller{contract: contract}, nil
}

// NewZapGettersLibraryTransactor creates a new write-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersLibraryTransactor, error) {
	contract, err := bindZapGettersLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryTransactor{contract: contract}, nil
}

// NewZapGettersLibraryFilterer creates a new log filterer instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersLibraryFilterer, error) {
	contract, err := bindZapGettersLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryFilterer{contract: contract}, nil
}

// bindZapGettersLibrary binds a generic wrapper to an already deployed contract.
func bindZapGettersLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapGettersLibraryNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddressIterator struct {
	Event *ZapGettersLibraryNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapGettersLibraryNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapGettersLibraryNewZapAddress)
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
		it.Event = new(ZapGettersLibraryNewZapAddress)
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
func (it *ZapGettersLibraryNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapGettersLibraryNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapGettersLibraryNewZapAddress represents a NewZapAddress event raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapGettersLibraryNewZapAddressIterator, error) {

	logs, sub, err := _ZapGettersLibrary.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryNewZapAddressIterator{contract: _ZapGettersLibrary.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapGettersLibraryNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapGettersLibrary.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapGettersLibraryNewZapAddress)
				if err := _ZapGettersLibrary.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) ParseNewZapAddress(log types.Log) (*ZapGettersLibraryNewZapAddress, error) {
	event := new(ZapGettersLibraryNewZapAddress)
	if err := _ZapGettersLibrary.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryABI is the input ABI used to generate the binding from.
const ZapLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"}]"

// ZapLibraryFuncSigs maps the 4-byte function signature to its string representation.
var ZapLibraryFuncSigs = map[string]string{
	"d723552b": "submitMiningSolution(ZapStorage.ZapStorageStruct storage,string,uint256,uint256)",
}

// ZapLibraryBin is the compiled bytecode used for deploying new contracts.
var ZapLibraryBin = "0x611783610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063d723552b1461003a575b600080fd5b81801561004657600080fd5b506100f96004803603608081101561005d57600080fd5b8135919081019060408101602082013564010000000081111561007f57600080fd5b82018360208201111561009157600080fd5b803590602001918460018302840111640100000000831117156100b357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356100fb565b005b33600090815260478501602052604090205460011461011957600080fd5b604080516f18dd5c9c995b9d14995c5d595cdd125960821b8152815190819003601001902060009081528186016020522054821461015657600080fd5b836040016000604051808069646966666963756c747960b01b815250600a0190506040518091039020815260200190815260200160002054600260038660000154338760405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b602083106101ec5780518252601f1990920191602091820191016101cd565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106102775780518252601f199092019160209182019101610258565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156102b6573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b6020831061032c5780518252601f19909201916020918201910161030d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561036b573d6000803e3d6000fd5b5050506040513d602081101561038057600080fd5b50518161038957fe5b061561039457600080fd5b83546000908152604185016020908152604080832033845290915290205460ff16156103bf57600080fd5b604080517118dd5c9c995b9d135a5b995c94995dd85c9960721b81528151908190036012018120600090815282870160208181528483208390556b736c6f7450726f677265737360a01b8452845193849003600c0190932082529091522054819060358601906005811061042f57fe5b6002020155604080516b736c6f7450726f677265737360a01b8152815190819003600c01902060009081528186016020522054339060358601906005811061047357fe5b60020201600190810180546001600160a01b0319166001600160a01b039390931692909217909155604080516b736c6f7450726f677265737360a01b81528151600c91819003919091018120600090815287830160209081528382208054860190558854825260418901815283822033808452908252848320805460ff19169096179095558854818401879052938301849052606080845288519084015287518795947fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4948a9489949293919283926080840192918801918190849084905b8381101561056a578181015183820152602001610552565b50505050905090810190601f1680156105975780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3604080516b736c6f7450726f677265737360a01b8152815190819003600c01902060009081528186016020522054600514156105e7576105e78484846105ed565b50505050565b6000818152604884016020908152604080832081517174696d654f664c6173744e657756616c756560701b81528251601291819003919091018120855282880180855283862054691d1a5b5955185c99d95d60b21b83528451600a938190038401812088528287528588205469646966666963756c747960b01b808352875192839003860183208a52848952878a20549083528751928390039095019091208852919095529285205491946064429590950390930302929092059091019081136106e5576040805169646966666963756c747960b01b8152815190819003600a01902060009081528187016020522060019055610714565b6040805169646966666963756c747960b01b8152815190819003600a0190206000908152818701602052208190555b603c42604080517174696d654f664c6173744e657756616c756560701b8152815190819003601201902060009081528189016020522091900642039055610759611637565b6040805160a081019091526035870160056000835b828210156107af5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161076e565b509293506001925050505b60058110156108d95760008282600581106107d157fe5b602002015151905060008383600581106107e757fe5b602002015160200151905060008390505b60008111801561081b575084600182036005811061081257fe5b60200201515183105b1561088d5784600182036005811061082f57fe5b60200201515185826005811061084157fe5b6020020151528460001982016005811061085757fe5b60200201516020015185826005811061086c57fe5b602090810291909101516001600160a01b03909216910152600019016107f8565b838110156108ce57828582600581106108a257fe5b602002015152818582600581106108b557fe5b602090810291909101516001600160a01b039092169101525b5050506001016107ba565b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d0190206000908152818901602052205461094657604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d0190206000908152818901602052206753444835ec58000090555b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d01902060009081528189016020522054670de0b6b3a76400001015610a5157604080516c18dd5c9c995b9d14995dd85c99609a1b8082528251600d928190038301812060009081528b850160208181528683205485855287519485900387018520845282825287842054868652885195869003880186208552838352888520670de0b6b3a7640000651bd78f205bc69093028390049091039055948452865193849003909501909220815292529190205460649160329104026040805167646576536861726560c01b815281519081900360080190206000908152818b01602052209190049055610a8b565b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d019020600090815281890160205220670de0b6b3a764000090555b604080516c18dd5c9c995b9d14995dd85c99609a1b8152815190819003600d01812060009081528983016020818152848320546f63757272656e74546f74616c5469707360801b808652865195869003601090810187208652848452878620547118dd5c9c995b9d135a5b995c94995dd85c9960721b88528851601298819003890181208852868652898820670de0b6b3a764000090950460059283900481019095557174696d654f664c6173744e657756616c756560701b815289519081900390980188208752858552888720548b8a015151848a528a51998a900384018a2089528787528a892054948a528a51998a90039093018920885295855295889020548f5495885293870152939093069003838501526060830152915187917fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16916080918190039190910190a2604080516b746f74616c5f737570706c7960a01b8152815190819003600c01812060009081528a8301602081815284832080546101130190558785018051517174696d654f664c6173744e657756616c756560701b80875287519687900360129081018820875285855288872054875260068e018552888720929092558087528751968790038201872086528484528786205460038e018054600181018255908852858820015560a0870188528a518401516001600160a01b0390811688528b85015185015181168886015292518401518316878901526060808c01518501518416908801526080808c0151850151909316928701929092528651918252865191829003019020835290815283822054825260088901905291909120610cfe916005611664565b506040805160a081018252845151815260208086015151818301528583015151828401526060808701515190830152608080870151519083015282517174696d654f664c6173744e657756616c756560701b8152835190819003601201902060009081528b840182528381205481526009890190915291909120610d839160056116bc565b50604080517174696d654f664c6173744e657756616c756560701b808252825191829003601290810183206000908152848d01602081815286832054835260058c01815286832043905584865286519586900384018620835281815286832054835260428f0181528683208d9055938552855194859003909201842081528183528481205460348e01805460018101825590835284832001556b736c6f7450726f677265737360a01b8452845193849003600c019093208352905290812055610e4b886115a7565b604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282516010928190038301812060009081528d8501602081815286832097909755928252845191829003909301902082529092529020541561156457604080516f18dd5c9c995b9d14995c5d595cdd125960821b808252825191829003601090810183206000908152848d01602081815286832054835260488f01808252878420670746f74616c5469760c41b88528851978890036008018820855260049081018352888520546f63757272656e74546f74616c5469707360801b8952895198899003870189208652848452898620559587528751968790038501872084529181528683205483529081528582206f3932b8bab2b9ba28a837b9b4ba34b7b760811b8652865195869003909301909420815291019091529081205460018a019060338110610f8f57fe5b018190555060008860430160008a60480160008c604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000208190555060008860480160008a604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b8152506010019050604051809103902081526020019081526020016000208190555060008860480160008a604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002060040160006040518080670746f74616c5469760c41b815250600801905060405180910390208152602001908152602001600020819055506000611168896115a7565b905087896000015460014303406040516020018084805190602001908083835b602083106111a75780518252601f199092019160209182019101611188565b6001836020036101000a038019825116818451168082178552505050505050905001838152602001828152602001935050505060405160208183030381529060405280519060200120896000018190555088604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428a600001548b6040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548c60480160008e604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548d60480160008f604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018e604001600060405180806f63757272656e74546f74616c5469707360801b8152506010019050604051809103902081526020019081526020016000205460405180868152602001858152602001848152602001806020018381526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156114525780601f1061142757610100808354040283529160200191611452565b820191906000526020600020905b81548152906001019060200180831161143557829003601f168201915b5050965050505050505060405180910390a2600081815260488a01602090815260408083206002808201548351670746f74616c5469760c41b81528451908190036008018120875260048401865295849020549486018190529285018490526060808652825460001961010060018316150201169190910490850181905285947f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c9492939291819060808201908690801561154e5780601f106115235761010080835404028352916020019161154e565b820191906000526020600020905b81548152906001019060200180831161153157829003601f168201915b505094505050505060405180910390a25061159d565b604080516f63757272656e74546f74616c5469707360801b815281519081900360100190206000908152818a0160205290812081905588555b5050505050505050565b60408051610660810191829052600091829182916115e89190600187019060339082845b8154815260200190600101908083116115cb575050505050611603565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611631576020810284015180841015611628578093508192505b50600101611609565b50915091565b6040518060a001604052806005905b61164e6116f6565b8152602001906001900390816116465790505090565b82600581019282156116ac579160200282015b828111156116ac57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611677565b506116b892915061170d565b5090565b82600581019282156116ea579160200282015b828111156116ea5782518255916020019190600101906116cf565b506116b8929150611734565b604080518082019091526000808252602082015290565b61173191905b808211156116b85780546001600160a01b0319168155600101611713565b90565b61173191905b808211156116b8576000815560010161173a56fea265627a7a72315820682589045eb8598b86732271a92c6be45e347c564f4499436149af621d9e64b464736f6c63430005100032"

// DeployZapLibrary deploys a new Ethereum contract, binding an instance of ZapLibrary to it.
func DeployZapLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// ZapLibrary is an auto generated Go binding around an Ethereum contract.
type ZapLibrary struct {
	ZapLibraryCaller     // Read-only binding to the contract
	ZapLibraryTransactor // Write-only binding to the contract
	ZapLibraryFilterer   // Log filterer for contract events
}

// ZapLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapLibrarySession struct {
	Contract     *ZapLibrary       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapLibraryCallerSession struct {
	Contract *ZapLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapLibraryTransactorSession struct {
	Contract     *ZapLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapLibraryRaw struct {
	Contract *ZapLibrary // Generic contract binding to access the raw methods on
}

// ZapLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapLibraryCallerRaw struct {
	Contract *ZapLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapLibraryTransactorRaw struct {
	Contract *ZapLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapLibrary creates a new instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibrary(address common.Address, backend bind.ContractBackend) (*ZapLibrary, error) {
	contract, err := bindZapLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// NewZapLibraryCaller creates a new read-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapLibraryCaller, error) {
	contract, err := bindZapLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryCaller{contract: contract}, nil
}

// NewZapLibraryTransactor creates a new write-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapLibraryTransactor, error) {
	contract, err := bindZapLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTransactor{contract: contract}, nil
}

// NewZapLibraryFilterer creates a new log filterer instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapLibraryFilterer, error) {
	contract, err := bindZapLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryFilterer{contract: contract}, nil
}

// bindZapLibrary binds a generic wrapper to an already deployed contract.
func bindZapLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.ZapLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapLibraryDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the ZapLibrary contract.
type ZapLibraryDataRequestedIterator struct {
	Event *ZapLibraryDataRequested // Event containing the contract specifics and raw log

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
func (it *ZapLibraryDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryDataRequested)
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
		it.Event = new(ZapLibraryDataRequested)
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
func (it *ZapLibraryDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryDataRequested represents a DataRequested event raised by the ZapLibrary contract.
type ZapLibraryDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryDataRequestedIterator{contract: _ZapLibrary.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *ZapLibraryDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryDataRequested)
				if err := _ZapLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
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

// ParseDataRequested is a log parse operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: event DataRequested(address indexed _sender, string _query, string _querySymbol, uint256 _granularity, uint256 indexed _requestId, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseDataRequested(log types.Log) (*ZapLibraryDataRequested, error) {
	event := new(ZapLibraryDataRequested)
	if err := _ZapLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the ZapLibrary contract.
type ZapLibraryNewChallengeIterator struct {
	Event *ZapLibraryNewChallenge // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewChallenge)
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
		it.Event = new(ZapLibraryNewChallenge)
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
func (it *ZapLibraryNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewChallenge represents a NewChallenge event raised by the ZapLibrary contract.
type ZapLibraryNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*ZapLibraryNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewChallengeIterator{contract: _ZapLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewChallenge)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ParseNewChallenge is a log parse operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: event NewChallenge(bytes32 _currentChallenge, uint256 indexed _currentRequestId, uint256 _difficulty, uint256 _multiplier, string _query, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewChallenge(log types.Log) (*ZapLibraryNewChallenge, error) {
	event := new(ZapLibraryNewChallenge)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeckIterator struct {
	Event *ZapLibraryNewRequestOnDeck // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewRequestOnDeck)
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
		it.Event = new(ZapLibraryNewRequestOnDeck)
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
func (it *ZapLibraryNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewRequestOnDeck represents a NewRequestOnDeck event raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewRequestOnDeckIterator{contract: _ZapLibrary.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewRequestOnDeck)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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

// ParseNewRequestOnDeck is a log parse operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: event NewRequestOnDeck(uint256 indexed _requestId, string _query, bytes32 _onDeckQueryHash, uint256 _onDeckTotalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewRequestOnDeck(log types.Log) (*ZapLibraryNewRequestOnDeck, error) {
	event := new(ZapLibraryNewRequestOnDeck)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the ZapLibrary contract.
type ZapLibraryNewValueIterator struct {
	Event *ZapLibraryNewValue // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewValue)
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
		it.Event = new(ZapLibraryNewValue)
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
func (it *ZapLibraryNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewValue represents a NewValue event raised by the ZapLibrary contract.
type ZapLibraryNewValue struct {
	RequestId        *big.Int
	Time             *big.Int
	Value            *big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewValueIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewValueIterator{contract: _ZapLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewValue, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewValue)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: event NewValue(uint256 indexed _requestId, uint256 _time, uint256 _value, uint256 _totalTips, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) ParseNewValue(log types.Log) (*ZapLibraryNewValue, error) {
	event := new(ZapLibraryNewValue)
	if err := _ZapLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the ZapLibrary contract.
type ZapLibraryNonceSubmittedIterator struct {
	Event *ZapLibraryNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNonceSubmitted)
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
		it.Event = new(ZapLibraryNonceSubmitted)
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
func (it *ZapLibraryNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNonceSubmitted represents a NonceSubmitted event raised by the ZapLibrary contract.
type ZapLibraryNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        *big.Int
	Value            *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _requestId []*big.Int) (*ZapLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNonceSubmittedIterator{contract: _ZapLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *ZapLibraryNonceSubmitted, _miner []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNonceSubmitted)
				if err := _ZapLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256 indexed _requestId, uint256 _value, bytes32 _currentChallenge)
func (_ZapLibrary *ZapLibraryFilterer) ParseNonceSubmitted(log types.Log) (*ZapLibraryNonceSubmitted, error) {
	event := new(ZapLibraryNonceSubmitted)
	if err := _ZapLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferredIterator struct {
	Event *ZapLibraryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZapLibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryOwnershipTransferred)
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
		it.Event = new(ZapLibraryOwnershipTransferred)
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
func (it *ZapLibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryOwnershipTransferred represents a OwnershipTransferred event raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*ZapLibraryOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryOwnershipTransferredIterator{contract: _ZapLibrary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapLibraryOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryOwnershipTransferred)
				if err := _ZapLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner)
func (_ZapLibrary *ZapLibraryFilterer) ParseOwnershipTransferred(log types.Log) (*ZapLibraryOwnershipTransferred, error) {
	event := new(ZapLibraryOwnershipTransferred)
	if err := _ZapLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapLibraryTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the ZapLibrary contract.
type ZapLibraryTipAddedIterator struct {
	Event *ZapLibraryTipAdded // Event containing the contract specifics and raw log

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
func (it *ZapLibraryTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryTipAdded)
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
		it.Event = new(ZapLibraryTipAdded)
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
func (it *ZapLibraryTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryTipAdded represents a TipAdded event raised by the ZapLibrary contract.
type ZapLibraryTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTipAddedIterator{contract: _ZapLibrary.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ZapLibraryTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryTipAdded)
				if err := _ZapLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ZapLibrary *ZapLibraryFilterer) ParseTipAdded(log types.Log) (*ZapLibraryTipAdded, error) {
	event := new(ZapLibraryTipAdded)
	if err := _ZapLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapMasterABI is the input ABI used to generate the binding from.
const ZapMasterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultContract\",\"type\":\"address\"}],\"name\":\"changeVaultContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"changeZapContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZapMasterFuncSigs maps the 4-byte function signature to its string representation.
var ZapMasterFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"70a08231": "balanceOf(address)",
	"47abd7f1": "changeDeity(address)",
	"af2a5102": "changeVaultContract(address)",
	"e4203f66": "changeZapContract(address)",
	"63bb82ad": "didMine(bytes32,address)",
	"a7c438bc": "didVote(uint256,address)",
	"133bee5e": "getAddressVars(bytes32)",
	"af0b1327": "getAllDisputeVars(uint256)",
	"a22e407a": "getCurrentVariables()",
	"da379941": "getDisputeIdByDisputeHash(bytes32)",
	"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
	"fc7cf0a0": "getLastNewValue()",
	"3180f8df": "getLastNewValueById(uint256)",
	"c775b542": "getMinedBlockNum(uint256,uint256)",
	"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
	"17d7de7c": "getName()",
	"46eee1c4": "getNewValueCountbyRequestId(uint256)",
	"1db842f0": "getRequestIdByQueryHash(bytes32)",
	"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
	"0f0b424d": "getRequestIdByTimestamp(uint256)",
	"b5413029": "getRequestQ()",
	"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
	"e1eee6d6": "getRequestVars(uint256)",
	"733bdef0": "getStakerInfo(address)",
	"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
	"15070401": "getSymbol()",
	"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
	"612c8f7f": "getUintVar(bytes32)",
	"19e8e03b": "getVariablesOnDeck()",
	"3df0777b": "isInDispute(uint256,uint256)",
	"93fa4915": "retrieveData(uint256,uint256)",
	"1ca8b6cb": "totalTokenSupply()",
}

// ZapMasterBin is the compiled bytecode used for deploying new contracts.
var ZapMasterBin = "0x608060405234801561001057600080fd5b50604051611f66380380611f668339818101604052604081101561003357600080fd5b508051602090910151604b80546001600160a01b0319166001600160a01b038316179055604080516347b024eb60e01b8152600060048201819052915173__$ceb17272379ced45263b3cfc641f24c214$__926347b024eb9260248082019391829003018186803b1580156100a757600080fd5b505af41580156100bb573d6000803e3d6000fd5b505060408051652fb7bbb732b960d11b8152815190819003600690810182206000908152603f602081815285832080546001600160a01b031990811633908117909255655f646569747960d01b8752875196879003909501862084528282528684208054861690911790556a1e985c10dbdb9d1c9858dd60aa1b8552855194859003600b018520835281815285832080546001600160a01b038c811691871682179092556f1e985c151bdad95b90dbdb9d1c9858dd60821b875287519687900360100187208552928252928690208054938a16939094169290921790925590825291517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab2709450908190039091019150a15050611d8a806101dc6000396000f3fe6080604052600436106101ee5760003560e01c806370a082311161010d578063af2a5102116100a0578063dd62ed3e1161006f578063dd62ed3e146109e8578063e0ae93c114610a23578063e1eee6d614610a53578063e4203f6614610b77578063fc7cf0a014610baa576101ee565b8063af2a510214610930578063b541302914610963578063c775b5421461098e578063da379941146109be576101ee565b806393fa4915116100dc57806393fa491514610766578063a22e407a14610796578063a7c438bc14610846578063af0b13271461087f576101ee565b806370a0823114610687578063733bdef0146106ba57806377fbb663146107065780637f6fd5d914610736576101ee565b80633180f8df11610185578063612c8f7f11610154578063612c8f7f146105ca5780636173c0b8146105f457806363bb82ad1461061e57806369026d6314610657576101ee565b80633180f8df146104e45780633df0777b1461052757806346eee1c41461056b57806347abd7f114610595576101ee565b806317d7de7c116101c157806317d7de7c146103f557806319e8e03b1461040a5780631ca8b6cb146104a55780631db842f0146104ba576101ee565b80630f0b424d1461028157806311c98512146102bd578063133bee5e14610325578063150704011461036b575b604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f602090815283822054601f369081018390048302850183019095528484526001600160a01b03169360609392918190840183828082843760009201829052508451949550938493509150506020840185600019f43d604051816000823e82801561027d578282f35b8282fd5b34801561028d57600080fd5b506102ab600480360360208110156102a457600080fd5b5035610bbf565b60408051918252519081900360200190f35b3480156102c957600080fd5b506102ed600480360360408110156102e057600080fd5b5080359060200135610bd7565b604051808260a080838360005b838110156103125781810151838201526020016102fa565b5050505090500191505060405180910390f35b34801561033157600080fd5b5061034f6004803603602081101561034857600080fd5b5035610bf8565b604080516001600160a01b039092168252519081900360200190f35b34801561037757600080fd5b50610380610c0a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103ba5781810151838201526020016103a2565b50505050905090810190601f1680156103e75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040157600080fd5b50610380610c1b565b34801561041657600080fd5b5061041f610c27565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610468578181015183820152602001610450565b50505050905090810190601f1680156104955780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104b157600080fd5b506102ab610c41565b3480156104c657600080fd5b506102ab600480360360208110156104dd57600080fd5b5035610c4d565b3480156104f057600080fd5b5061050e6004803603602081101561050757600080fd5b5035610c5f565b6040805192835290151560208301528051918290030190f35b34801561053357600080fd5b506105576004803603604081101561054a57600080fd5b5080359060200135610c7b565b604080519115158252519081900360200190f35b34801561057757600080fd5b506102ab6004803603602081101561058e57600080fd5b5035610c8e565b3480156105a157600080fd5b506105c8600480360360208110156105b857600080fd5b50356001600160a01b0316610ca0565b005b3480156105d657600080fd5b506102ab600480360360208110156105ed57600080fd5b5035610cb4565b34801561060057600080fd5b506102ab6004803603602081101561061757600080fd5b5035610cc6565b34801561062a57600080fd5b506105576004803603604081101561064157600080fd5b50803590602001356001600160a01b0316610cd8565b34801561066357600080fd5b506102ed6004803603604081101561067a57600080fd5b5080359060200135610ceb565b34801561069357600080fd5b506102ab600480360360208110156106aa57600080fd5b50356001600160a01b0316610d05565b3480156106c657600080fd5b506106ed600480360360208110156106dd57600080fd5b50356001600160a01b0316610d88565b6040805192835260208301919091528051918290030190f35b34801561071257600080fd5b506102ab6004803603604081101561072957600080fd5b5080359060200135610d9b565b34801561074257600080fd5b506102ab6004803603604081101561075957600080fd5b5080359060200135610dae565b34801561077257600080fd5b506102ab6004803603604081101561078957600080fd5b5080359060200135610dc1565b3480156107a257600080fd5b506107ab610dd4565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156108065781810151838201526020016107ee565b50505050905090810190601f1680156108335780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b34801561085257600080fd5b506105576004803603604081101561086957600080fd5b50803590602001356001600160a01b0316610dfb565b34801561088b57600080fd5b506108a9600480360360208110156108a257600080fd5b5035610e0e565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561090f5781810151838201526020016108f7565b50505050905001828152602001995050505050505050505060405180910390f35b34801561093c57600080fd5b506105c86004803603602081101561095357600080fd5b50356001600160a01b0316610e52565b34801561096f57600080fd5b50610978610e63565b60405181518152808261066080838360206102fa565b34801561099a57600080fd5b506102ab600480360360408110156109b157600080fd5b5080359060200135610e75565b3480156109ca57600080fd5b506102ab600480360360208110156109e157600080fd5b5035610e88565b3480156109f457600080fd5b506102ab60048036036040811015610a0b57600080fd5b506001600160a01b0381358116916020013516610e9a565b348015610a2f57600080fd5b506102ab60048036036040811015610a4657600080fd5b5080359060200135610f26565b348015610a5f57600080fd5b50610a7d60048036036020811015610a7657600080fd5b5035610f39565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610ad6578181015183820152602001610abe565b50505050905090810190601f168015610b035780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610b36578181015183820152602001610b1e565b50505050905090810190601f168015610b635780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610b8357600080fd5b506105c860048036036020811015610b9a57600080fd5b50356001600160a01b0316610f65565b348015610bb657600080fd5b5061050e610f76565b6000610bd1818363ffffffff610f8b16565b92915050565b610bdf611cf9565b610bf16000848463ffffffff610fa116565b9392505050565b6000610bd1818363ffffffff610ffb16565b6060610c16600061101a565b905090565b6060610c166000611039565b6000806060610c36600061105d565b925092509250909192565b6000610c166000611149565b6000610bd1818363ffffffff61117c16565b600080610c72818463ffffffff61119216565b91509150915091565b6000610bf181848463ffffffff6111f816565b6000610bd1818363ffffffff61121f16565b610cb160008263ffffffff61123816565b50565b6000610bd1818363ffffffff6112c116565b6000610bd1818363ffffffff6112d316565b6000610bf181848463ffffffff6112fa16565b610cf3611cf9565b610bf16000848463ffffffff61132716565b604b54604080516370a0823160e01b81526001600160a01b038481166004830152915160009392909216916370a0823191602480820192602092909190829003018186803b158015610d5657600080fd5b505afa158015610d6a573d6000803e3d6000fd5b505050506040513d6020811015610d8057600080fd5b505192915050565b600080610c72818463ffffffff61138b16565b6000610bf181848463ffffffff6113b216565b6000610bf181848463ffffffff6113e516565b6000610bf181848463ffffffff61140916565b60008060006060600080610de8600061142d565b949b939a50919850965094509092509050565b6000610bf181848463ffffffff6115eb16565b6000806000806000806000610e21611d17565b6000610e33818b63ffffffff61161d16565b9850985098509850985098509850985098509193959799909294969850565b610cb160008263ffffffff61183216565b610e6b611d36565b610c1660006118bb565b6000610bf181848463ffffffff6118fb16565b6000610bd1818363ffffffff61191f16565b604b5460408051636eb1769f60e11b81526001600160a01b03858116600483015284811660248301529151600093929092169163dd62ed3e91604480820192602092909190829003018186803b158015610ef357600080fd5b505afa158015610f07573d6000803e3d6000fd5b505050506040513d6020811015610f1d57600080fd5b50519392505050565b6000610bf181848463ffffffff61193516565b6060806000808080610f51818863ffffffff61195916565b949c939b5091995097509550909350915050565b610cb160008263ffffffff611b3416565b600080610f836000611bf3565b915091509091565b6000908152604291909101602052604090205490565b610fa9611cf9565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610fda57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b506040805180820190915260048152632d20a82160e11b602082015290565b5060408051808201909152600981526805a61702042455032360bc1b602082015290565b6000806060600061106d85611c69565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f8101849004840285018401909252818452949550859492918391908301828280156111345780601f1061110957610100808354040283529160200191611134565b820191906000526020600020905b81548152906001019060200180831161111757829003601f168201915b50505050509050935093509350509193909250565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b6000908152604991909101602052604090205490565b600081815260488301602052604081206003810154829190156111e8576003810180546111dc91879187919060001981019081106111cc57fe5b9060005260206000200154611409565b600192509250506111f1565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461127557600080fd5b60408051655f646569747960d01b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b60009081526040918201602052205490565b600060328211156112e357600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b61132f611cf9565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161136057505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b600082815260488401602052604081206003018054839081106113d157fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156115cf5780601f106115a4576101008083540402835291602001916115cf565b820191906000526020600020905b8154815290600101906020018083116115b257829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b6000806000806000806000611630611d17565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b60408051652fb7bbb732b960d11b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461186f57600080fd5b604080516517dd985d5b1d60d21b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b6118c3611d36565b6040805161066081019182905290600184019060339082845b8154815260200190600101908083116118dc5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a99909889988998899891978897938801969295929490918891830182828015611a885780601f10611a5d57610100808354040283529160200191611a88565b820191906000526020600020905b815481529060010190602001808311611a6b57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015611b165780601f10611aeb57610100808354040283529160200191611b16565b820191906000526020600020905b815481529060010190602001808311611af957829003601f168201915b50505050509450965096509650965096509650509295509295509295565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b03163314611b7157600080fd5b604080516a1e985c10dbdb9d1c9858dd60aa1b8152815190819003600b0181206000908152603f850160209081529083902080546001600160a01b0386166001600160a01b03199091168117909155825291517f4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270929181900390910190a15050565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611c5f918591611409565b9360019350915050565b6040805161066081019182905260009182918291611caa9190600187019060339082845b815481526020019060010190808311611c8d575050505050611cc5565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611cf3576020810284015180841015611cea578093508192505b50600101611ccb565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a723158202e15cf4bea8d9d62d3f25c3ab70eb7cbb98c65a7ec70b2d6a63d3272677d257d64736f6c63430005100032"

// DeployZapMaster deploys a new Ethereum contract, binding an instance of ZapMaster to it.
func DeployZapMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _zapContract common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *ZapMaster, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapStakeAddr, _, _, _ := DeployZapStake(auth, backend)
	ZapMasterBin = strings.Replace(ZapMasterBin, "__$ceb17272379ced45263b3cfc641f24c214$__", zapStakeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapMasterBin), backend, _zapContract, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// ZapMaster is an auto generated Go binding around an Ethereum contract.
type ZapMaster struct {
	ZapMasterCaller     // Read-only binding to the contract
	ZapMasterTransactor // Write-only binding to the contract
	ZapMasterFilterer   // Log filterer for contract events
}

// ZapMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapMasterSession struct {
	Contract     *ZapMaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapMasterCallerSession struct {
	Contract *ZapMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZapMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapMasterTransactorSession struct {
	Contract     *ZapMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZapMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapMasterRaw struct {
	Contract *ZapMaster // Generic contract binding to access the raw methods on
}

// ZapMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapMasterCallerRaw struct {
	Contract *ZapMasterCaller // Generic read-only contract binding to access the raw methods on
}

// ZapMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapMasterTransactorRaw struct {
	Contract *ZapMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapMaster creates a new instance of ZapMaster, bound to a specific deployed contract.
func NewZapMaster(address common.Address, backend bind.ContractBackend) (*ZapMaster, error) {
	contract, err := bindZapMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// NewZapMasterCaller creates a new read-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterCaller(address common.Address, caller bind.ContractCaller) (*ZapMasterCaller, error) {
	contract, err := bindZapMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterCaller{contract: contract}, nil
}

// NewZapMasterTransactor creates a new write-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapMasterTransactor, error) {
	contract, err := bindZapMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterTransactor{contract: contract}, nil
}

// NewZapMasterFilterer creates a new log filterer instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapMasterFilterer, error) {
	contract, err := bindZapMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapMasterFilterer{contract: contract}, nil
}

// bindZapMaster binds a generic wrapper to an already deployed contract.
func bindZapMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.ZapMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ZapMaster *ZapMasterCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getCurrentVariables")

	if err != nil {
		return *new([32]byte), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() view returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ZapMaster *ZapMasterCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByQueryHash", _request)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(bytes32 _request) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ZapMaster *ZapMasterCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(string), *new(string), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ZapMaster *ZapMasterCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "getVariablesOnDeck")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() view returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ZapMaster *ZapMasterCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterCaller) TotalTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZapMaster.contract.Call(opts, &out, "totalTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalTokenSupply(&_ZapMaster.CallOpts)
}

// TotalTokenSupply is a free data retrieval call binding the contract method 0x1ca8b6cb.
//
// Solidity: function totalTokenSupply() view returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) TotalTokenSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalTokenSupply(&_ZapMaster.CallOpts)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeVaultContract(opts *bind.TransactOpts, _vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeVaultContract", _vaultContract)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterSession) ChangeVaultContract(_vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeVaultContract(&_ZapMaster.TransactOpts, _vaultContract)
}

// ChangeVaultContract is a paid mutator transaction binding the contract method 0xaf2a5102.
//
// Solidity: function changeVaultContract(address _vaultContract) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeVaultContract(_vaultContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeVaultContract(&_ZapMaster.TransactOpts, _vaultContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeZapContract", _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xe4203f66.
//
// Solidity: function changeZapContract(address _zapContract) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.Contract.Fallback(&_ZapMaster.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZapMaster *ZapMasterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZapMaster.Contract.Fallback(&_ZapMaster.TransactOpts, calldata)
}

// ZapMasterNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapMaster contract.
type ZapMasterNewZapAddressIterator struct {
	Event *ZapMasterNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapMasterNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapMasterNewZapAddress)
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
		it.Event = new(ZapMasterNewZapAddress)
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
func (it *ZapMasterNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapMasterNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapMasterNewZapAddress represents a NewZapAddress event raised by the ZapMaster contract.
type ZapMasterNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapMasterNewZapAddressIterator, error) {

	logs, sub, err := _ZapMaster.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapMasterNewZapAddressIterator{contract: _ZapMaster.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapMasterNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapMaster.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapMasterNewZapAddress)
				if err := _ZapMaster.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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

// ParseNewZapAddress is a log parse operation binding the contract event 0x4a9a276906262ed9ed5e1fd15850a5f2b951b97198cc2fc0d32625f1bf3ab270.
//
// Solidity: event NewZapAddress(address _newZap)
func (_ZapMaster *ZapMasterFilterer) ParseNewZapAddress(log types.Log) (*ZapMasterNewZapAddress, error) {
	event := new(ZapMasterNewZapAddress)
	if err := _ZapMaster.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeABI is the input ABI used to generate the binding from.
const ZapStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"}]"

// ZapStakeFuncSigs maps the 4-byte function signature to its string representation.
var ZapStakeFuncSigs = map[string]string{
	"326991a3": "depositStake(ZapStorage.ZapStorageStruct storage)",
	"47b024eb": "init(ZapStorage.ZapStorageStruct storage)",
	"3c734827": "requestStakingWithdraw(ZapStorage.ZapStorageStruct storage)",
	"78bfa277": "withdrawStake(ZapStorage.ZapStorageStruct storage)",
}

// ZapStakeBin is the compiled bytecode used for deploying new contracts.
var ZapStakeBin = "0x61056b610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c8063326991a31461005b5780633c7348271461008757806347b024eb146100b157806378bfa277146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610179565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610270565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b50356103fd565b61010f813361044d565b73__$78be728ae59efcf7bc760a8cdd86b21178$__639264b888826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561015e57600080fd5b505af4158015610172573d6000803e3d6000fd5b5050505050565b3360009081526047820160205260409020805460011461019857600080fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0181206000908152828501602052828120805460001901905563124c971160e31b825260048201859052915173__$78be728ae59efcf7bc760a8cdd86b21178$__92639264b8889260248082019391829003018186803b15801561022957600080fd5b505af415801561023d573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b6040805167646563696d616c7360c01b8152815190819003600801902060009081528183016020522054156102a457600080fd5b6040805167646563696d616c7360c01b8152815190819003600801812060009081528284016020818152848320601290556b7461726765744d696e65727360a01b8452845193849003600c018420835281815284832060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b01842083528181528483206207a1209055696469737075746546656560b01b8452845193849003600a908101852084528282528584206103ca9055691d1a5b5955185c99d95d60b21b808652865195869003820186208552838352868520610258905585528551948590030190932082529091522054428161039557fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152948201602081815283872095909406420390945569646966666963756c747960b01b8152815190819003600a0190208452919052902060019055565b3360009081526047820160205260409020805460021461041c57600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6001600160a01b0381166000908152604783016020526040902054158061048e57506001600160a01b03811660009081526047830160205260409020546002145b61049757600080fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a2505056fea265627a7a72315820861520f0a0bc9c2c6ae5dc00316dd2db3f8eaeec1997534e12713274475b8c2964736f6c63430005100032"

// DeployZapStake deploys a new Ethereum contract, binding an instance of ZapStake to it.
func DeployZapStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStake, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	zapDisputeAddr, _, _, _ := DeployZapDispute(auth, backend)
	ZapStakeBin = strings.Replace(ZapStakeBin, "__$78be728ae59efcf7bc760a8cdd86b21178$__", zapDisputeAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// ZapStake is an auto generated Go binding around an Ethereum contract.
type ZapStake struct {
	ZapStakeCaller     // Read-only binding to the contract
	ZapStakeTransactor // Write-only binding to the contract
	ZapStakeFilterer   // Log filterer for contract events
}

// ZapStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStakeSession struct {
	Contract     *ZapStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStakeCallerSession struct {
	Contract *ZapStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZapStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStakeTransactorSession struct {
	Contract     *ZapStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZapStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStakeRaw struct {
	Contract *ZapStake // Generic contract binding to access the raw methods on
}

// ZapStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStakeCallerRaw struct {
	Contract *ZapStakeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStakeTransactorRaw struct {
	Contract *ZapStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStake creates a new instance of ZapStake, bound to a specific deployed contract.
func NewZapStake(address common.Address, backend bind.ContractBackend) (*ZapStake, error) {
	contract, err := bindZapStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// NewZapStakeCaller creates a new read-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeCaller(address common.Address, caller bind.ContractCaller) (*ZapStakeCaller, error) {
	contract, err := bindZapStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeCaller{contract: contract}, nil
}

// NewZapStakeTransactor creates a new write-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStakeTransactor, error) {
	contract, err := bindZapStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeTransactor{contract: contract}, nil
}

// NewZapStakeFilterer creates a new log filterer instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStakeFilterer, error) {
	contract, err := bindZapStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStakeFilterer{contract: contract}, nil
}

// bindZapStake binds a generic wrapper to an already deployed contract.
func bindZapStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.ZapStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transact(opts, method, params...)
}

// ZapStakeNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the ZapStake contract.
type ZapStakeNewStakeIterator struct {
	Event *ZapStakeNewStake // Event containing the contract specifics and raw log

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
func (it *ZapStakeNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeNewStake)
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
		it.Event = new(ZapStakeNewStake)
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
func (it *ZapStakeNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeNewStake represents a NewStake event raised by the ZapStake contract.
type ZapStakeNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeNewStakeIterator{contract: _ZapStake.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ZapStakeNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeNewStake)
				if err := _ZapStake.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// ParseNewStake is a log parse operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseNewStake(log types.Log) (*ZapStakeNewStake, error) {
	event := new(ZapStakeNewStake)
	if err := _ZapStake.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequestedIterator struct {
	Event *ZapStakeStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ZapStakeStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawRequested)
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
		it.Event = new(ZapStakeStakeWithdrawRequested)
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
func (it *ZapStakeStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawRequestedIterator{contract: _ZapStake.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawRequested)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// ParseStakeWithdrawRequested is a log parse operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseStakeWithdrawRequested(log types.Log) (*ZapStakeStakeWithdrawRequested, error) {
	event := new(ZapStakeStakeWithdrawRequested)
	if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ZapStake contract.
type ZapStakeStakeWithdrawnIterator struct {
	Event *ZapStakeStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZapStakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawn)
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
		it.Event = new(ZapStakeStakeWithdrawn)
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
func (it *ZapStakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawn represents a StakeWithdrawn event raised by the ZapStake contract.
type ZapStakeStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawnIterator{contract: _ZapStake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawn)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ZapStake *ZapStakeFilterer) ParseStakeWithdrawn(log types.Log) (*ZapStakeStakeWithdrawn, error) {
	event := new(ZapStakeStakeWithdrawn)
	if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapStorageABI is the input ABI used to generate the binding from.
const ZapStorageABI = "[]"

// ZapStorageBin is the compiled bytecode used for deploying new contracts.
var ZapStorageBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158206d586078f1a88031d927514e195cbb9c4813080956442ecc1dc386986a2fa09c64736f6c63430005100032"

// DeployZapStorage deploys a new Ethereum contract, binding an instance of ZapStorage to it.
func DeployZapStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// ZapStorage is an auto generated Go binding around an Ethereum contract.
type ZapStorage struct {
	ZapStorageCaller     // Read-only binding to the contract
	ZapStorageTransactor // Write-only binding to the contract
	ZapStorageFilterer   // Log filterer for contract events
}

// ZapStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStorageSession struct {
	Contract     *ZapStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStorageCallerSession struct {
	Contract *ZapStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStorageTransactorSession struct {
	Contract     *ZapStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStorageRaw struct {
	Contract *ZapStorage // Generic contract binding to access the raw methods on
}

// ZapStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStorageCallerRaw struct {
	Contract *ZapStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStorageTransactorRaw struct {
	Contract *ZapStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStorage creates a new instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorage(address common.Address, backend bind.ContractBackend) (*ZapStorage, error) {
	contract, err := bindZapStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// NewZapStorageCaller creates a new read-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageCaller(address common.Address, caller bind.ContractCaller) (*ZapStorageCaller, error) {
	contract, err := bindZapStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageCaller{contract: contract}, nil
}

// NewZapStorageTransactor creates a new write-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStorageTransactor, error) {
	contract, err := bindZapStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageTransactor{contract: contract}, nil
}

// NewZapStorageFilterer creates a new log filterer instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStorageFilterer, error) {
	contract, err := bindZapStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStorageFilterer{contract: contract}, nil
}

// bindZapStorage binds a generic wrapper to an already deployed contract.
func bindZapStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.ZapStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transact(opts, method, params...)
}

// ZapTokenBSCABI is the input ABI used to generate the binding from.
const ZapTokenBSCABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"allocate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZapTokenBSCFuncSigs maps the 4-byte function signature to its string representation.
var ZapTokenBSCFuncSigs = map[string]string{
	"b78b52df": "allocate(address,uint256)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"66188463": "decreaseApproval(address,uint256)",
	"7d64bcb4": "finishMinting()",
	"893d20e8": "getOwner()",
	"d73dd623": "increaseApproval(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"05d2035b": "mintingFinished()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// ZapTokenBSCBin is the compiled bytecode used for deploying new contracts.
var ZapTokenBSCBin = "0x6a52b7d2dcc80cd2e40000006000556003805460ff60a01b1916905560c0604052600960808190526805a61702042455032360bc1b60a090815261004691600491906100ed565b50604080518082019091526004808252632d20a82160e11b6020909201918252610072916005916100ed565b506006805460ff1916601217905534801561008c57600080fd5b50600380546001600160a01b03191633908117909155600080548282526001602090815260408084208390558051928352517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3610188565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012e57805160ff191683800117855561015b565b8280016001018555821561015b579182015b8281111561015b578251825591602001919060010190610140565b5061016792915061016b565b5090565b61018591905b808211156101675760008155600101610171565b90565b610b79806101976000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80637d64bcb4116100a2578063a9059cbb11610071578063a9059cbb14610308578063b78b52df14610334578063d73dd62314610362578063dd62ed3e1461038e578063f2fde38b146103bc57610116565b80637d64bcb4146102cc578063893d20e8146102d45780638da5cb5b146102f857806395d89b411461030057610116565b806323b872dd116100e957806323b872dd146101fa578063313ce5671461023057806340c10f191461024e578063661884631461027a57806370a08231146102a657610116565b806305d2035b1461011b57806306fdde0314610137578063095ea7b3146101b457806318160ddd146101e0575b600080fd5b6101236103e2565b604080519115158252519081900360200190f35b61013f6103f2565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610179578181015183820152602001610161565b50505050905090810190601f1680156101a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610123600480360360408110156101ca57600080fd5b506001600160a01b038135169060200135610480565b6101e86104e6565b60408051918252519081900360200190f35b6101236004803603606081101561021057600080fd5b506001600160a01b038135811691602081013590911690604001356104ec565b61023861060a565b6040805160ff9092168252519081900360200190f35b6101236004803603604081101561026457600080fd5b506001600160a01b038135169060200135610613565b6101236004803603604081101561029057600080fd5b506001600160a01b03813516906020013561071e565b6101e8600480360360208110156102bc57600080fd5b50356001600160a01b031661080e565b610123610829565b6102dc610885565b604080516001600160a01b039092168252519081900360200190f35b6102dc610894565b61013f6108a3565b6101236004803603604081101561031e57600080fd5b506001600160a01b0381351690602001356108fe565b6103606004803603604081101561034a57600080fd5b506001600160a01b0381351690602001356109c3565b005b6101236004803603604081101561037857600080fd5b506001600160a01b0381351690602001356109d2565b6101e8600480360360408110156103a457600080fd5b506001600160a01b0381358116916020013516610a6b565b610360600480360360208110156103d257600080fd5b50356001600160a01b0316610a96565b600354600160a01b900460ff1681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104785780601f1061044d57610100808354040283529160200191610478565b820191906000526020600020905b81548152906001019060200180831161045b57829003601f168201915b505050505081565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60006001600160a01b03831661050157600080fd5b6001600160a01b03841660008181526002602090815260408083203384528252808320549383526001909152902054610540908463ffffffff610b1c16565b6001600160a01b038087166000908152600160205260408082209390935590861681522054610575908463ffffffff610b2e16565b6001600160a01b03851660009081526001602052604090205561059e818463ffffffff610b1c16565b6001600160a01b03808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b60065460ff1681565b6003546000906001600160a01b0316331461062d57600080fd5b600354600160a01b900460ff161561064457600080fd5b600054610657908363ffffffff610b2e16565b60009081556001600160a01b038416815260016020526040902054610682908363ffffffff610b2e16565b6001600160a01b038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a26040805183815290516001600160a01b038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b3360009081526002602090815260408083206001600160a01b038616845290915281205480831115610773573360009081526002602090815260408083206001600160a01b03881684529091528120556107a8565b610783818463ffffffff610b1c16565b3360009081526002602090815260408083206001600160a01b03891684529091529020555b3360008181526002602090815260408083206001600160a01b0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b6003546000906001600160a01b0316331461084357600080fd5b6003805460ff60a01b1916600160a01b1790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b6003546001600160a01b031690565b6003546001600160a01b031681565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104785780601f1061044d57610100808354040283529160200191610478565b60006001600160a01b03831661091357600080fd5b33600090815260016020526040902054610933908363ffffffff610b1c16565b33600090815260016020526040808220929092556001600160a01b03851681522054610965908363ffffffff610b2e16565b6001600160a01b0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b6109cd8282610613565b505050565b3360009081526002602090815260408083206001600160a01b0386168452909152812054610a06908363ffffffff610b2e16565b3360008181526002602090815260408083206001600160a01b0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6003546001600160a01b03163314610aad57600080fd5b6001600160a01b038116610ac057600080fd5b6003546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b600082821115610b2857fe5b50900390565b600082820183811015610b3d57fe5b939250505056fea265627a7a72315820332fc95d78eda577cc43047718591fbd51df488f31c1fb14d69a949bc0b9909464736f6c63430005100032"

// DeployZapTokenBSC deploys a new Ethereum contract, binding an instance of ZapTokenBSC to it.
func DeployZapTokenBSC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapTokenBSC, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTokenBSCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapTokenBSCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapTokenBSC{ZapTokenBSCCaller: ZapTokenBSCCaller{contract: contract}, ZapTokenBSCTransactor: ZapTokenBSCTransactor{contract: contract}, ZapTokenBSCFilterer: ZapTokenBSCFilterer{contract: contract}}, nil
}

// ZapTokenBSC is an auto generated Go binding around an Ethereum contract.
type ZapTokenBSC struct {
	ZapTokenBSCCaller     // Read-only binding to the contract
	ZapTokenBSCTransactor // Write-only binding to the contract
	ZapTokenBSCFilterer   // Log filterer for contract events
}

// ZapTokenBSCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapTokenBSCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTokenBSCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapTokenBSCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTokenBSCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapTokenBSCSession struct {
	Contract     *ZapTokenBSC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapTokenBSCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapTokenBSCCallerSession struct {
	Contract *ZapTokenBSCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZapTokenBSCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTokenBSCTransactorSession struct {
	Contract     *ZapTokenBSCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZapTokenBSCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapTokenBSCRaw struct {
	Contract *ZapTokenBSC // Generic contract binding to access the raw methods on
}

// ZapTokenBSCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapTokenBSCCallerRaw struct {
	Contract *ZapTokenBSCCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTokenBSCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTokenBSCTransactorRaw struct {
	Contract *ZapTokenBSCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapTokenBSC creates a new instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSC(address common.Address, backend bind.ContractBackend) (*ZapTokenBSC, error) {
	contract, err := bindZapTokenBSC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSC{ZapTokenBSCCaller: ZapTokenBSCCaller{contract: contract}, ZapTokenBSCTransactor: ZapTokenBSCTransactor{contract: contract}, ZapTokenBSCFilterer: ZapTokenBSCFilterer{contract: contract}}, nil
}

// NewZapTokenBSCCaller creates a new read-only instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCCaller(address common.Address, caller bind.ContractCaller) (*ZapTokenBSCCaller, error) {
	contract, err := bindZapTokenBSC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCCaller{contract: contract}, nil
}

// NewZapTokenBSCTransactor creates a new write-only instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTokenBSCTransactor, error) {
	contract, err := bindZapTokenBSC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCTransactor{contract: contract}, nil
}

// NewZapTokenBSCFilterer creates a new log filterer instance of ZapTokenBSC, bound to a specific deployed contract.
func NewZapTokenBSCFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapTokenBSCFilterer, error) {
	contract, err := bindZapTokenBSC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCFilterer{contract: contract}, nil
}

// bindZapTokenBSC binds a generic wrapper to an already deployed contract.
func bindZapTokenBSC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTokenBSCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTokenBSC *ZapTokenBSCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTokenBSC.Contract.ZapTokenBSCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTokenBSC *ZapTokenBSCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.ZapTokenBSCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTokenBSC *ZapTokenBSCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.ZapTokenBSCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTokenBSC *ZapTokenBSCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTokenBSC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTokenBSC *ZapTokenBSCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTokenBSC *ZapTokenBSCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.Allowance(&_ZapTokenBSC.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.Allowance(&_ZapTokenBSC.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.BalanceOf(&_ZapTokenBSC.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ZapTokenBSC.Contract.BalanceOf(&_ZapTokenBSC.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCSession) Decimals() (uint8, error) {
	return _ZapTokenBSC.Contract.Decimals(&_ZapTokenBSC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Decimals() (uint8, error) {
	return _ZapTokenBSC.Contract.Decimals(&_ZapTokenBSC.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCSession) GetOwner() (common.Address, error) {
	return _ZapTokenBSC.Contract.GetOwner(&_ZapTokenBSC.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) GetOwner() (common.Address, error) {
	return _ZapTokenBSC.Contract.GetOwner(&_ZapTokenBSC.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "mintingFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) MintingFinished() (bool, error) {
	return _ZapTokenBSC.Contract.MintingFinished(&_ZapTokenBSC.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() view returns(bool)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) MintingFinished() (bool, error) {
	return _ZapTokenBSC.Contract.MintingFinished(&_ZapTokenBSC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCSession) Name() (string, error) {
	return _ZapTokenBSC.Contract.Name(&_ZapTokenBSC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Name() (string, error) {
	return _ZapTokenBSC.Contract.Name(&_ZapTokenBSC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCSession) Owner() (common.Address, error) {
	return _ZapTokenBSC.Contract.Owner(&_ZapTokenBSC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Owner() (common.Address, error) {
	return _ZapTokenBSC.Contract.Owner(&_ZapTokenBSC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCSession) Symbol() (string, error) {
	return _ZapTokenBSC.Contract.Symbol(&_ZapTokenBSC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) Symbol() (string, error) {
	return _ZapTokenBSC.Contract.Symbol(&_ZapTokenBSC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZapTokenBSC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCSession) TotalSupply() (*big.Int, error) {
	return _ZapTokenBSC.Contract.TotalSupply(&_ZapTokenBSC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZapTokenBSC *ZapTokenBSCCallerSession) TotalSupply() (*big.Int, error) {
	return _ZapTokenBSC.Contract.TotalSupply(&_ZapTokenBSC.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactor) Allocate(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "allocate", to, amount)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCSession) Allocate(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Allocate(&_ZapTokenBSC.TransactOpts, to, amount)
}

// Allocate is a paid mutator transaction binding the contract method 0xb78b52df.
//
// Solidity: function allocate(address to, uint256 amount) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Allocate(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Allocate(&_ZapTokenBSC.TransactOpts, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Approve(&_ZapTokenBSC.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Approve(&_ZapTokenBSC.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.DecreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address _spender, uint256 _subtractedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.DecreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) FinishMinting() (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.FinishMinting(&_ZapTokenBSC.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.FinishMinting(&_ZapTokenBSC.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.IncreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address _spender, uint256 _addedValue) returns(bool success)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.IncreaseApproval(&_ZapTokenBSC.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Mint(&_ZapTokenBSC.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Mint(&_ZapTokenBSC.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Transfer(&_ZapTokenBSC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.Transfer(&_ZapTokenBSC.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferFrom(&_ZapTokenBSC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferFrom(&_ZapTokenBSC.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferOwnership(&_ZapTokenBSC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZapTokenBSC *ZapTokenBSCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZapTokenBSC.Contract.TransferOwnership(&_ZapTokenBSC.TransactOpts, newOwner)
}

// ZapTokenBSCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZapTokenBSC contract.
type ZapTokenBSCApprovalIterator struct {
	Event *ZapTokenBSCApproval // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCApproval)
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
		it.Event = new(ZapTokenBSCApproval)
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
func (it *ZapTokenBSCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCApproval represents a Approval event raised by the ZapTokenBSC contract.
type ZapTokenBSCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZapTokenBSCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCApprovalIterator{contract: _ZapTokenBSC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCApproval)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseApproval(log types.Log) (*ZapTokenBSCApproval, error) {
	event := new(ZapTokenBSCApproval)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ZapTokenBSC contract.
type ZapTokenBSCMintIterator struct {
	Event *ZapTokenBSCMint // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCMint)
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
		it.Event = new(ZapTokenBSCMint)
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
func (it *ZapTokenBSCMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCMint represents a Mint event raised by the ZapTokenBSC contract.
type ZapTokenBSCMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*ZapTokenBSCMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCMintIterator{contract: _ZapTokenBSC.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCMint)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Mint", log); err != nil {
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
// Solidity: event Mint(address indexed to, uint256 amount)
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseMint(log types.Log) (*ZapTokenBSCMint, error) {
	event := new(ZapTokenBSCMint)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the ZapTokenBSC contract.
type ZapTokenBSCMintFinishedIterator struct {
	Event *ZapTokenBSCMintFinished // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCMintFinished)
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
		it.Event = new(ZapTokenBSCMintFinished)
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
func (it *ZapTokenBSCMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCMintFinished represents a MintFinished event raised by the ZapTokenBSC contract.
type ZapTokenBSCMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterMintFinished(opts *bind.FilterOpts) (*ZapTokenBSCMintFinishedIterator, error) {

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCMintFinishedIterator{contract: _ZapTokenBSC.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCMintFinished) (event.Subscription, error) {

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCMintFinished)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// ParseMintFinished is a log parse operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseMintFinished(log types.Log) (*ZapTokenBSCMintFinished, error) {
	event := new(ZapTokenBSCMintFinished)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "MintFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZapTokenBSC contract.
type ZapTokenBSCOwnershipTransferredIterator struct {
	Event *ZapTokenBSCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCOwnershipTransferred)
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
		it.Event = new(ZapTokenBSCOwnershipTransferred)
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
func (it *ZapTokenBSCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCOwnershipTransferred represents a OwnershipTransferred event raised by the ZapTokenBSC contract.
type ZapTokenBSCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZapTokenBSCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCOwnershipTransferredIterator{contract: _ZapTokenBSC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCOwnershipTransferred)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseOwnershipTransferred(log types.Log) (*ZapTokenBSCOwnershipTransferred, error) {
	event := new(ZapTokenBSCOwnershipTransferred)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTokenBSCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZapTokenBSC contract.
type ZapTokenBSCTransferIterator struct {
	Event *ZapTokenBSCTransfer // Event containing the contract specifics and raw log

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
func (it *ZapTokenBSCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTokenBSCTransfer)
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
		it.Event = new(ZapTokenBSCTransfer)
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
func (it *ZapTokenBSCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTokenBSCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTokenBSCTransfer represents a Transfer event raised by the ZapTokenBSC contract.
type ZapTokenBSCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZapTokenBSCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTokenBSCTransferIterator{contract: _ZapTokenBSC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZapTokenBSC *ZapTokenBSCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZapTokenBSCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZapTokenBSC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTokenBSCTransfer)
				if err := _ZapTokenBSC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ZapTokenBSC *ZapTokenBSCFilterer) ParseTransfer(log types.Log) (*ZapTokenBSCTransfer, error) {
	event := new(ZapTokenBSCTransfer)
	if err := _ZapTokenBSC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZapTransferABI is the input ABI used to generate the binding from.
const ZapTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ZapTransferFuncSigs maps the 4-byte function signature to its string representation.
var ZapTransferFuncSigs = map[string]string{
	"fade3342": "allowance(ZapStorage.ZapStorageStruct storage,address,address)",
	"b9290ca5": "allowedToTrade(ZapStorage.ZapStorageStruct storage,address,uint256)",
	"c6f7efe0": "balanceOfAt(ZapStorage.ZapStorageStruct storage,address,uint256)",
	"a93a4d03": "doTransfer(ZapStorage.ZapStorageStruct storage,address,address,uint256)",
	"82c1fecb": "getBalanceAt(ZapStorage.Checkpoint[] storage,uint256)",
	"2dfd8908": "updateBalanceAtNow(ZapStorage.Checkpoint[] storage,uint256)",
}

// ZapTransferBin is the compiled bytecode used for deploying new contracts.
var ZapTransferBin = "0x610663610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c80632dfd89081461007157806382c1fecb146100a3578063a93a4d03146100d8578063b9290ca514610121578063c6f7efe014610167578063fade334214610199575b600080fd5b81801561007d57600080fd5b506100a16004803603604081101561009457600080fd5b50803590602001356101cd565b005b6100c6600480360360408110156100b957600080fd5b50803590602001356102a6565b60408051918252519081900360200190f35b8180156100e457600080fd5b506100a1600480360360808110156100fb57600080fd5b508035906001600160a01b036020820135811691604081013590911690606001356103d8565b6101536004803603606081101561013757600080fd5b508035906001600160a01b036020820135169060400135610497565b604080519115158252519081900360200190f35b6100c66004803603606081101561017d57600080fd5b508035906001600160a01b03602082013516906040013561050f565b6100c6600480360360608110156101af57600080fd5b508035906001600160a01b03602082013581169160400135166105a5565b81541580610201575081544390839060001981019081106101ea57fe5b6000918252602090912001546001600160801b0316105b15610268578154600090839061021a82600183016105e4565b8154811061022457fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506102a2565b81546000908390600019810190811061027d57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b81546000906102b7575060006103d2565b8254839060001981019081106102c957fe5b6000918252602090912001546001600160801b03168210610319578254839060001981019081106102f657fe5b600091825260209091200154600160801b90046001600160801b031690506103d2565b8260008154811061032657fe5b6000918252602090912001546001600160801b031682101561034a575060006103d2565b8254600090600019015b818111156103a557600060026001838501010490508486828154811061037657fe5b6000918252602090912001546001600160801b0316116103985780925061039f565b6001810391505b50610354565b8482815481106103b157fe5b600091825260209091200154600160801b90046001600160801b0316925050505b92915050565b600081116103e557600080fd5b6001600160a01b0382166103f857600080fd5b610403848483610497565b61040c57600080fd5b600061041985854361050f565b6001600160a01b03851660009081526045870160205260409020909150610442908383036101cd565b61044d85844361050f565b905080828201101561045e57600080fd5b6001600160a01b03831660009081526045860160205260409020610484908284016101cd565b61048f85844361050f565b505050505050565b6001600160a01b0382166000908152604784016020526040812054156104e65760006104d4836104c887874361050f565b9063ffffffff6105d216565b106104e157506001610508565b610504565b60006104f7836104c887874361050f565b1061050457506001610508565b5060005b9392505050565b6001600160a01b0382166000908152604584016020526040812054158061056d57506001600160a01b03831660009081526045850160205260408120805484929061055657fe5b6000918252602090912001546001600160801b0316115b1561057a57506000610508565b6001600160a01b0383166000908152604585016020526040902061059e90836102a6565b9050610508565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b6000828211156105de57fe5b50900390565b8154818355818111156106085760008381526020902061060891810190830161060d565b505050565b61062b91905b808211156106275760008155600101610613565b5090565b9056fea265627a7a723158207d03224e02e9d6b060f68a4ea102f7f112a027731fe0a403366a0db0d016738c64736f6c63430005100032"

// DeployZapTransfer deploys a new Ethereum contract, binding an instance of ZapTransfer to it.
func DeployZapTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// ZapTransfer is an auto generated Go binding around an Ethereum contract.
type ZapTransfer struct {
	ZapTransferCaller     // Read-only binding to the contract
	ZapTransferTransactor // Write-only binding to the contract
	ZapTransferFilterer   // Log filterer for contract events
}

// ZapTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapTransferSession struct {
	Contract     *ZapTransfer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapTransferCallerSession struct {
	Contract *ZapTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZapTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransferTransactorSession struct {
	Contract     *ZapTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZapTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapTransferRaw struct {
	Contract *ZapTransfer // Generic contract binding to access the raw methods on
}

// ZapTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapTransferCallerRaw struct {
	Contract *ZapTransferCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransferTransactorRaw struct {
	Contract *ZapTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapTransfer creates a new instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransfer(address common.Address, backend bind.ContractBackend) (*ZapTransfer, error) {
	contract, err := bindZapTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// NewZapTransferCaller creates a new read-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferCaller(address common.Address, caller bind.ContractCaller) (*ZapTransferCaller, error) {
	contract, err := bindZapTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferCaller{contract: contract}, nil
}

// NewZapTransferTransactor creates a new write-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransferTransactor, error) {
	contract, err := bindZapTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferTransactor{contract: contract}, nil
}

// NewZapTransferFilterer creates a new log filterer instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapTransferFilterer, error) {
	contract, err := bindZapTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapTransferFilterer{contract: contract}, nil
}

// bindZapTransfer binds a generic wrapper to an already deployed contract.
func bindZapTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.ZapTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transact(opts, method, params...)
}

// ZapTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZapTransfer contract.
type ZapTransferApprovalIterator struct {
	Event *ZapTransferApproval // Event containing the contract specifics and raw log

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
func (it *ZapTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferApproval)
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
		it.Event = new(ZapTransferApproval)
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
func (it *ZapTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransferApproval represents a Approval event raised by the ZapTransfer contract.
type ZapTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ZapTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferApprovalIterator{contract: _ZapTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferApproval)
				if err := _ZapTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ZapTransfer *ZapTransferFilterer) ParseApproval(log types.Log) (*ZapTransferApproval, error) {
	event := new(ZapTransferApproval)
	if err := _ZapTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}