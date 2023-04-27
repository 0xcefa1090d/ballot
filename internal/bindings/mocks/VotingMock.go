// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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
	_ = abi.ConvertType
)

// VotingMockMetaData contains all meta data concerning the VotingMock contract.
var VotingMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vote_id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"min_balance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"min_time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total_supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creator_voting_power\",\"type\":\"uint256\"}],\"name\":\"StartVote\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"_metadata\",\"type\":\"string\"},{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"new_vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vote_id\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x0x61022361001161000039610223610000f36003361161000c5761020b565b60003560e01c34610211576322c44e8381186101665760843610610211576004356004016104008135116102115780358060405260208201818160603750505060243560040161040081351161021157803580610460526020820181816104803750505060005461088052610880516001810181811061021157905060005561046051806001610880516020526000526040600020556001600161088051602052600052604060002001600082601f0160051c602081116102115780156100e757905b8060051b6104800151818401556001018181186100cf575b5050505033610880517f0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e39460a0806108a052806108a00160405180825260208201818183606060045afa5050508051806020830101601f82600003163682375050601f19601f825160200101169050810190506080366108c0376108a0a3005b635a55c1f08118610209576024361061021157610140610120366040378061016052600160043560205260005260406000208160400181548082526001830160208301600083601f0160051c602081116102115780156101d857905b808401548160051b8401526001018181186101c2575b50505050508051806020830101601f82600003163682375050601f19601f8251602001011690509050810190506040f35b505b60006000fd5b600080fda165767970657283000307000b",
}

// VotingMockABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMockMetaData.ABI instead.
var VotingMockABI = VotingMockMetaData.ABI

// VotingMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMockMetaData.Bin instead.
var VotingMockBin = VotingMockMetaData.Bin

// DeployVotingMock deploys a new Ethereum contract, binding an instance of VotingMock to it.
func DeployVotingMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VotingMock, error) {
	parsed, err := VotingMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VotingMock{VotingMockCaller: VotingMockCaller{contract: contract}, VotingMockTransactor: VotingMockTransactor{contract: contract}, VotingMockFilterer: VotingMockFilterer{contract: contract}}, nil
}

// VotingMock is an auto generated Go binding around an Ethereum contract.
type VotingMock struct {
	VotingMockCaller     // Read-only binding to the contract
	VotingMockTransactor // Write-only binding to the contract
	VotingMockFilterer   // Log filterer for contract events
}

// VotingMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingMockSession struct {
	Contract     *VotingMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingMockCallerSession struct {
	Contract *VotingMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VotingMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingMockTransactorSession struct {
	Contract     *VotingMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VotingMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingMockRaw struct {
	Contract *VotingMock // Generic contract binding to access the raw methods on
}

// VotingMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingMockCallerRaw struct {
	Contract *VotingMockCaller // Generic read-only contract binding to access the raw methods on
}

// VotingMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingMockTransactorRaw struct {
	Contract *VotingMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingMock creates a new instance of VotingMock, bound to a specific deployed contract.
func NewVotingMock(address common.Address, backend bind.ContractBackend) (*VotingMock, error) {
	contract, err := bindVotingMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingMock{VotingMockCaller: VotingMockCaller{contract: contract}, VotingMockTransactor: VotingMockTransactor{contract: contract}, VotingMockFilterer: VotingMockFilterer{contract: contract}}, nil
}

// NewVotingMockCaller creates a new read-only instance of VotingMock, bound to a specific deployed contract.
func NewVotingMockCaller(address common.Address, caller bind.ContractCaller) (*VotingMockCaller, error) {
	contract, err := bindVotingMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingMockCaller{contract: contract}, nil
}

// NewVotingMockTransactor creates a new write-only instance of VotingMock, bound to a specific deployed contract.
func NewVotingMockTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingMockTransactor, error) {
	contract, err := bindVotingMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingMockTransactor{contract: contract}, nil
}

// NewVotingMockFilterer creates a new log filterer instance of VotingMock, bound to a specific deployed contract.
func NewVotingMockFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingMockFilterer, error) {
	contract, err := bindVotingMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingMockFilterer{contract: contract}, nil
}

// bindVotingMock binds a generic wrapper to an already deployed contract.
func bindVotingMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingMock *VotingMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingMock.Contract.VotingMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingMock *VotingMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingMock.Contract.VotingMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingMock *VotingMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingMock.Contract.VotingMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingMock *VotingMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingMock *VotingMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingMock *VotingMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingMock.Contract.contract.Transact(opts, method, params...)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _vote_id) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_VotingMock *VotingMockCaller) GetVote(opts *bind.CallOpts, _vote_id *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	var out []interface{}
	err := _VotingMock.contract.Call(opts, &out, "getVote", _vote_id)

	if err != nil {
		return *new(bool), *new(bool), *new(uint64), *new(uint64), *new(uint64), *new(uint64), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(uint64)).(*uint64)
	out5 := *abi.ConvertType(out[5], new(uint64)).(*uint64)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _vote_id) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_VotingMock *VotingMockSession) GetVote(_vote_id *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	return _VotingMock.Contract.GetVote(&_VotingMock.CallOpts, _vote_id)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _vote_id) view returns(bool, bool, uint64, uint64, uint64, uint64, uint256, uint256, uint256, bytes)
func (_VotingMock *VotingMockCallerSession) GetVote(_vote_id *big.Int) (bool, bool, uint64, uint64, uint64, uint64, *big.Int, *big.Int, *big.Int, []byte, error) {
	return _VotingMock.Contract.GetVote(&_VotingMock.CallOpts, _vote_id)
}

// NewVote is a paid mutator transaction binding the contract method 0x22c44e83.
//
// Solidity: function new_vote(string _metadata, bytes _script) returns()
func (_VotingMock *VotingMockTransactor) NewVote(opts *bind.TransactOpts, _metadata string, _script []byte) (*types.Transaction, error) {
	return _VotingMock.contract.Transact(opts, "new_vote", _metadata, _script)
}

// NewVote is a paid mutator transaction binding the contract method 0x22c44e83.
//
// Solidity: function new_vote(string _metadata, bytes _script) returns()
func (_VotingMock *VotingMockSession) NewVote(_metadata string, _script []byte) (*types.Transaction, error) {
	return _VotingMock.Contract.NewVote(&_VotingMock.TransactOpts, _metadata, _script)
}

// NewVote is a paid mutator transaction binding the contract method 0x22c44e83.
//
// Solidity: function new_vote(string _metadata, bytes _script) returns()
func (_VotingMock *VotingMockTransactorSession) NewVote(_metadata string, _script []byte) (*types.Transaction, error) {
	return _VotingMock.Contract.NewVote(&_VotingMock.TransactOpts, _metadata, _script)
}

// VotingMockStartVoteIterator is returned from FilterStartVote and is used to iterate over the raw logs and unpacked data for StartVote events raised by the VotingMock contract.
type VotingMockStartVoteIterator struct {
	Event *VotingMockStartVote // Event containing the contract specifics and raw log

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
func (it *VotingMockStartVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingMockStartVote)
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
		it.Event = new(VotingMockStartVote)
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
func (it *VotingMockStartVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingMockStartVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingMockStartVote represents a StartVote event raised by the VotingMock contract.
type VotingMockStartVote struct {
	VoteId             *big.Int
	Creator            common.Address
	Metadata           string
	MinBalance         *big.Int
	MinTime            *big.Int
	TotalSupply        *big.Int
	CreatorVotingPower *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStartVote is a free log retrieval operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed vote_id, address indexed creator, string metadata, uint256 min_balance, uint256 min_time, uint256 total_supply, uint256 creator_voting_power)
func (_VotingMock *VotingMockFilterer) FilterStartVote(opts *bind.FilterOpts, vote_id []*big.Int, creator []common.Address) (*VotingMockStartVoteIterator, error) {

	var vote_idRule []interface{}
	for _, vote_idItem := range vote_id {
		vote_idRule = append(vote_idRule, vote_idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VotingMock.contract.FilterLogs(opts, "StartVote", vote_idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &VotingMockStartVoteIterator{contract: _VotingMock.contract, event: "StartVote", logs: logs, sub: sub}, nil
}

// WatchStartVote is a free log subscription operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed vote_id, address indexed creator, string metadata, uint256 min_balance, uint256 min_time, uint256 total_supply, uint256 creator_voting_power)
func (_VotingMock *VotingMockFilterer) WatchStartVote(opts *bind.WatchOpts, sink chan<- *VotingMockStartVote, vote_id []*big.Int, creator []common.Address) (event.Subscription, error) {

	var vote_idRule []interface{}
	for _, vote_idItem := range vote_id {
		vote_idRule = append(vote_idRule, vote_idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VotingMock.contract.WatchLogs(opts, "StartVote", vote_idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingMockStartVote)
				if err := _VotingMock.contract.UnpackLog(event, "StartVote", log); err != nil {
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

// ParseStartVote is a log parse operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed vote_id, address indexed creator, string metadata, uint256 min_balance, uint256 min_time, uint256 total_supply, uint256 creator_voting_power)
func (_VotingMock *VotingMockFilterer) ParseStartVote(log types.Log) (*VotingMockStartVote, error) {
	event := new(VotingMockStartVote)
	if err := _VotingMock.contract.UnpackLog(event, "StartVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
