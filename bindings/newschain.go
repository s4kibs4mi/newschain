// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"createdAt\",\"type\":\"int256\"}],\"name\":\"PostCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"name\":\"PostUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"created_at\",\"type\":\"int256\"}],\"name\":\"createAuthor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"created_at\",\"type\":\"int256\"}],\"name\":\"createPost\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"created_at\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getPost\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"created_at\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"updated_at\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAuthorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"updateAuthor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"updated_at\",\"type\":\"int256\"}],\"name\":\"updatePost\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetAuthorProfile is a free data retrieval call binding the contract method 0xfbea1472.
//
// Solidity: function getAuthorProfile() view returns(string name, string title, string email, int256 created_at)
func (_Contracts *ContractsCaller) GetAuthorProfile(opts *bind.CallOpts) (struct {
	Name      string
	Title     string
	Email     string
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAuthorProfile")

	outstruct := new(struct {
		Name      string
		Title     string
		Email     string
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = out[0].(string)
	outstruct.Title = out[1].(string)
	outstruct.Email = out[2].(string)
	outstruct.CreatedAt = out[3].(*big.Int)

	return *outstruct, err

}

// GetAuthorProfile is a free data retrieval call binding the contract method 0xfbea1472.
//
// Solidity: function getAuthorProfile() view returns(string name, string title, string email, int256 created_at)
func (_Contracts *ContractsSession) GetAuthorProfile() (struct {
	Name      string
	Title     string
	Email     string
	CreatedAt *big.Int
}, error) {
	return _Contracts.Contract.GetAuthorProfile(&_Contracts.CallOpts)
}

// GetAuthorProfile is a free data retrieval call binding the contract method 0xfbea1472.
//
// Solidity: function getAuthorProfile() view returns(string name, string title, string email, int256 created_at)
func (_Contracts *ContractsCallerSession) GetAuthorProfile() (struct {
	Name      string
	Title     string
	Email     string
	CreatedAt *big.Int
}, error) {
	return _Contracts.Contract.GetAuthorProfile(&_Contracts.CallOpts)
}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string id) view returns(string title, address author, string details, int256 created_at, int256 updated_at)
func (_Contracts *ContractsCaller) GetPost(opts *bind.CallOpts, id string) (struct {
	Title     string
	Author    common.Address
	Details   string
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPost", id)

	outstruct := new(struct {
		Title     string
		Author    common.Address
		Details   string
		CreatedAt *big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = out[0].(string)
	outstruct.Author = out[1].(common.Address)
	outstruct.Details = out[2].(string)
	outstruct.CreatedAt = out[3].(*big.Int)
	outstruct.UpdatedAt = out[4].(*big.Int)

	return *outstruct, err

}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string id) view returns(string title, address author, string details, int256 created_at, int256 updated_at)
func (_Contracts *ContractsSession) GetPost(id string) (struct {
	Title     string
	Author    common.Address
	Details   string
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	return _Contracts.Contract.GetPost(&_Contracts.CallOpts, id)
}

// GetPost is a free data retrieval call binding the contract method 0xdf4a259b.
//
// Solidity: function getPost(string id) view returns(string title, address author, string details, int256 created_at, int256 updated_at)
func (_Contracts *ContractsCallerSession) GetPost(id string) (struct {
	Title     string
	Author    common.Address
	Details   string
	CreatedAt *big.Int
	UpdatedAt *big.Int
}, error) {
	return _Contracts.Contract.GetPost(&_Contracts.CallOpts, id)
}

// IsAuthorRegistered is a free data retrieval call binding the contract method 0x4efbd806.
//
// Solidity: function isAuthorRegistered() view returns(bool)
func (_Contracts *ContractsCaller) IsAuthorRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isAuthorRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorRegistered is a free data retrieval call binding the contract method 0x4efbd806.
//
// Solidity: function isAuthorRegistered() view returns(bool)
func (_Contracts *ContractsSession) IsAuthorRegistered() (bool, error) {
	return _Contracts.Contract.IsAuthorRegistered(&_Contracts.CallOpts)
}

// IsAuthorRegistered is a free data retrieval call binding the contract method 0x4efbd806.
//
// Solidity: function isAuthorRegistered() view returns(bool)
func (_Contracts *ContractsCallerSession) IsAuthorRegistered() (bool, error) {
	return _Contracts.Contract.IsAuthorRegistered(&_Contracts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contracts *ContractsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contracts *ContractsSession) IsOwner() (bool, error) {
	return _Contracts.Contract.IsOwner(&_Contracts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contracts *ContractsCallerSession) IsOwner() (bool, error) {
	return _Contracts.Contract.IsOwner(&_Contracts.CallOpts)
}

// CreateAuthor is a paid mutator transaction binding the contract method 0xf2d601e0.
//
// Solidity: function createAuthor(string name, string title, string email, int256 created_at) payable returns()
func (_Contracts *ContractsTransactor) CreateAuthor(opts *bind.TransactOpts, name string, title string, email string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createAuthor", name, title, email, created_at)
}

// CreateAuthor is a paid mutator transaction binding the contract method 0xf2d601e0.
//
// Solidity: function createAuthor(string name, string title, string email, int256 created_at) payable returns()
func (_Contracts *ContractsSession) CreateAuthor(name string, title string, email string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAuthor(&_Contracts.TransactOpts, name, title, email, created_at)
}

// CreateAuthor is a paid mutator transaction binding the contract method 0xf2d601e0.
//
// Solidity: function createAuthor(string name, string title, string email, int256 created_at) payable returns()
func (_Contracts *ContractsTransactorSession) CreateAuthor(name string, title string, email string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAuthor(&_Contracts.TransactOpts, name, title, email, created_at)
}

// CreatePost is a paid mutator transaction binding the contract method 0xa108fd44.
//
// Solidity: function createPost(string id, string title, string details, int256 created_at) payable returns()
func (_Contracts *ContractsTransactor) CreatePost(opts *bind.TransactOpts, id string, title string, details string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createPost", id, title, details, created_at)
}

// CreatePost is a paid mutator transaction binding the contract method 0xa108fd44.
//
// Solidity: function createPost(string id, string title, string details, int256 created_at) payable returns()
func (_Contracts *ContractsSession) CreatePost(id string, title string, details string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePost(&_Contracts.TransactOpts, id, title, details, created_at)
}

// CreatePost is a paid mutator transaction binding the contract method 0xa108fd44.
//
// Solidity: function createPost(string id, string title, string details, int256 created_at) payable returns()
func (_Contracts *ContractsTransactorSession) CreatePost(id string, title string, details string, created_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePost(&_Contracts.TransactOpts, id, title, details, created_at)
}

// UpdateAuthor is a paid mutator transaction binding the contract method 0xf95faeaa.
//
// Solidity: function updateAuthor(string name, string title, string email) payable returns()
func (_Contracts *ContractsTransactor) UpdateAuthor(opts *bind.TransactOpts, name string, title string, email string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateAuthor", name, title, email)
}

// UpdateAuthor is a paid mutator transaction binding the contract method 0xf95faeaa.
//
// Solidity: function updateAuthor(string name, string title, string email) payable returns()
func (_Contracts *ContractsSession) UpdateAuthor(name string, title string, email string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAuthor(&_Contracts.TransactOpts, name, title, email)
}

// UpdateAuthor is a paid mutator transaction binding the contract method 0xf95faeaa.
//
// Solidity: function updateAuthor(string name, string title, string email) payable returns()
func (_Contracts *ContractsTransactorSession) UpdateAuthor(name string, title string, email string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateAuthor(&_Contracts.TransactOpts, name, title, email)
}

// UpdatePost is a paid mutator transaction binding the contract method 0xe1650361.
//
// Solidity: function updatePost(string id, string title, string details, int256 updated_at) payable returns()
func (_Contracts *ContractsTransactor) UpdatePost(opts *bind.TransactOpts, id string, title string, details string, updated_at *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updatePost", id, title, details, updated_at)
}

// UpdatePost is a paid mutator transaction binding the contract method 0xe1650361.
//
// Solidity: function updatePost(string id, string title, string details, int256 updated_at) payable returns()
func (_Contracts *ContractsSession) UpdatePost(id string, title string, details string, updated_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePost(&_Contracts.TransactOpts, id, title, details, updated_at)
}

// UpdatePost is a paid mutator transaction binding the contract method 0xe1650361.
//
// Solidity: function updatePost(string id, string title, string details, int256 updated_at) payable returns()
func (_Contracts *ContractsTransactorSession) UpdatePost(id string, title string, details string, updated_at *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePost(&_Contracts.TransactOpts, id, title, details, updated_at)
}

// ContractsPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the Contracts contract.
type ContractsPostCreatedIterator struct {
	Event *ContractsPostCreated // Event containing the contract specifics and raw log

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
func (it *ContractsPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPostCreated)
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
		it.Event = new(ContractsPostCreated)
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
func (it *ContractsPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPostCreated represents a PostCreated event raised by the Contracts contract.
type ContractsPostCreated struct {
	Id        string
	Title     string
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0xc099d9374b89a291eea81637ca8b5fb75897e5022789342cc21df63b0f4d26b0.
//
// Solidity: event PostCreated(string id, string title, int256 createdAt)
func (_Contracts *ContractsFilterer) FilterPostCreated(opts *bind.FilterOpts) (*ContractsPostCreatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PostCreated")
	if err != nil {
		return nil, err
	}
	return &ContractsPostCreatedIterator{contract: _Contracts.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0xc099d9374b89a291eea81637ca8b5fb75897e5022789342cc21df63b0f4d26b0.
//
// Solidity: event PostCreated(string id, string title, int256 createdAt)
func (_Contracts *ContractsFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *ContractsPostCreated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PostCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPostCreated)
				if err := _Contracts.contract.UnpackLog(event, "PostCreated", log); err != nil {
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

// ParsePostCreated is a log parse operation binding the contract event 0xc099d9374b89a291eea81637ca8b5fb75897e5022789342cc21df63b0f4d26b0.
//
// Solidity: event PostCreated(string id, string title, int256 createdAt)
func (_Contracts *ContractsFilterer) ParsePostCreated(log types.Log) (*ContractsPostCreated, error) {
	event := new(ContractsPostCreated)
	if err := _Contracts.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPostUpdatedIterator is returned from FilterPostUpdated and is used to iterate over the raw logs and unpacked data for PostUpdated events raised by the Contracts contract.
type ContractsPostUpdatedIterator struct {
	Event *ContractsPostUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsPostUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPostUpdated)
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
		it.Event = new(ContractsPostUpdated)
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
func (it *ContractsPostUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPostUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPostUpdated represents a PostUpdated event raised by the Contracts contract.
type ContractsPostUpdated struct {
	Id    string
	Title string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPostUpdated is a free log retrieval operation binding the contract event 0x40600307ff12cf80070001dc492e22e5f6d6f3961ce807f013a4a139edd3ab14.
//
// Solidity: event PostUpdated(string id, string title)
func (_Contracts *ContractsFilterer) FilterPostUpdated(opts *bind.FilterOpts) (*ContractsPostUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PostUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsPostUpdatedIterator{contract: _Contracts.contract, event: "PostUpdated", logs: logs, sub: sub}, nil
}

// WatchPostUpdated is a free log subscription operation binding the contract event 0x40600307ff12cf80070001dc492e22e5f6d6f3961ce807f013a4a139edd3ab14.
//
// Solidity: event PostUpdated(string id, string title)
func (_Contracts *ContractsFilterer) WatchPostUpdated(opts *bind.WatchOpts, sink chan<- *ContractsPostUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PostUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPostUpdated)
				if err := _Contracts.contract.UnpackLog(event, "PostUpdated", log); err != nil {
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

// ParsePostUpdated is a log parse operation binding the contract event 0x40600307ff12cf80070001dc492e22e5f6d6f3961ce807f013a4a139edd3ab14.
//
// Solidity: event PostUpdated(string id, string title)
func (_Contracts *ContractsFilterer) ParsePostUpdated(log types.Log) (*ContractsPostUpdated, error) {
	event := new(ContractsPostUpdated)
	if err := _Contracts.contract.UnpackLog(event, "PostUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
