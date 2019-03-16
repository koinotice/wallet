// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package a

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HashDiceABI is the input ABI used to generate the binding from.
const HashDiceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"gRoundWait\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gRoundPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gRoomId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gProfitRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"decimical\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastLockedValue\",\"type\":\"uint256\"}],\"name\":\"RoomOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"}],\"name\":\"RoomClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"remained\",\"type\":\"uint256\"}],\"name\":\"Withdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"remained\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NewBetOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlock\",\"type\":\"uint256\"}],\"name\":\"CloseRoundTooLate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roomId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PayBetOwner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"GetBetOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"}],\"name\":\"GetRoomInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256[4]\"},{\"name\":\"\",\"type\":\"uint256[4]\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetRooms\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetProfitRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"SetProfitRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetRoundWait\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetRoundWait\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetRoundPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetRoundPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"erc20Addr\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"initLockedValue\",\"type\":\"uint256\"},{\"name\":\"decimical\",\"type\":\"uint256\"},{\"name\":\"roomName\",\"type\":\"string\"},{\"name\":\"nominator\",\"type\":\"uint256[4]\"},{\"name\":\"denominator\",\"type\":\"uint256[4]\"}],\"name\":\"OpenRoom\",\"outputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"}],\"name\":\"CloseRoom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"GetProfit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"erc20Addr\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WithdrawProfit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"}],\"name\":\"CloseBetOrders\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"CloseBetOrder\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"betType\",\"type\":\"uint8\"},{\"name\":\"betValue\",\"type\":\"uint32[]\"}],\"name\":\"SubmitBetOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"betValue\",\"type\":\"uint32[]\"}],\"name\":\"SubmitBetOrderType0\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// HashDice is an auto generated Go binding around an Ethereum contract.
type HashDice struct {
	HashDiceCaller     // Read-only binding to the contract
	HashDiceTransactor // Write-only binding to the contract
	HashDiceFilterer   // Log filterer for contract events
}

// HashDiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashDiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashDiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashDiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashDiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashDiceSession struct {
	Contract     *HashDice         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashDiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashDiceCallerSession struct {
	Contract *HashDiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HashDiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashDiceTransactorSession struct {
	Contract     *HashDiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HashDiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashDiceRaw struct {
	Contract *HashDice // Generic contract binding to access the raw methods on
}

// HashDiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashDiceCallerRaw struct {
	Contract *HashDiceCaller // Generic read-only contract binding to access the raw methods on
}

// HashDiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashDiceTransactorRaw struct {
	Contract *HashDiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashDice creates a new instance of HashDice, bound to a specific deployed contract.
func NewHashDice(address common.Address, backend bind.ContractBackend) (*HashDice, error) {
	contract, err := bindHashDice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashDice{HashDiceCaller: HashDiceCaller{contract: contract}, HashDiceTransactor: HashDiceTransactor{contract: contract}, HashDiceFilterer: HashDiceFilterer{contract: contract}}, nil
}

// NewHashDiceCaller creates a new read-only instance of HashDice, bound to a specific deployed contract.
func NewHashDiceCaller(address common.Address, caller bind.ContractCaller) (*HashDiceCaller, error) {
	contract, err := bindHashDice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashDiceCaller{contract: contract}, nil
}

// NewHashDiceTransactor creates a new write-only instance of HashDice, bound to a specific deployed contract.
func NewHashDiceTransactor(address common.Address, transactor bind.ContractTransactor) (*HashDiceTransactor, error) {
	contract, err := bindHashDice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashDiceTransactor{contract: contract}, nil
}

// NewHashDiceFilterer creates a new log filterer instance of HashDice, bound to a specific deployed contract.
func NewHashDiceFilterer(address common.Address, filterer bind.ContractFilterer) (*HashDiceFilterer, error) {
	contract, err := bindHashDice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashDiceFilterer{contract: contract}, nil
}

// bindHashDice binds a generic wrapper to an already deployed contract.
func bindHashDice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashDiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashDice *HashDiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashDice.Contract.HashDiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashDice *HashDiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashDice.Contract.HashDiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashDice *HashDiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashDice.Contract.HashDiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashDice *HashDiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HashDice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashDice *HashDiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashDice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashDice *HashDiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashDice.Contract.contract.Transact(opts, method, params...)
}

// GetBetOrder is a free data retrieval call binding the contract method 0xb076650a.
//
// Solidity: function GetBetOrder(uint256 roomId, uint256 orderId) constant returns(address, uint256, uint256, uint256, uint256, bool, uint32[])
func (_HashDice *HashDiceCaller) GetBetOrder(opts *bind.CallOpts, roomId *big.Int, orderId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, bool, []uint32, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(bool)
		ret6 = new([]uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _HashDice.contract.Call(opts, out, "GetBetOrder", roomId, orderId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetBetOrder is a free data retrieval call binding the contract method 0xb076650a.
//
// Solidity: function GetBetOrder(uint256 roomId, uint256 orderId) constant returns(address, uint256, uint256, uint256, uint256, bool, uint32[])
func (_HashDice *HashDiceSession) GetBetOrder(roomId *big.Int, orderId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, bool, []uint32, error) {
	return _HashDice.Contract.GetBetOrder(&_HashDice.CallOpts, roomId, orderId)
}

// GetBetOrder is a free data retrieval call binding the contract method 0xb076650a.
//
// Solidity: function GetBetOrder(uint256 roomId, uint256 orderId) constant returns(address, uint256, uint256, uint256, uint256, bool, uint32[])
func (_HashDice *HashDiceCallerSession) GetBetOrder(roomId *big.Int, orderId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, bool, []uint32, error) {
	return _HashDice.Contract.GetBetOrder(&_HashDice.CallOpts, roomId, orderId)
}

// GetProfit is a free data retrieval call binding the contract method 0xccc8efcc.
//
// Solidity: function GetProfit(address erc20Addr) constant returns(uint256)
func (_HashDice *HashDiceCaller) GetProfit(opts *bind.CallOpts, erc20Addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "GetProfit", erc20Addr)
	return *ret0, err
}

// GetProfit is a free data retrieval call binding the contract method 0xccc8efcc.
//
// Solidity: function GetProfit(address erc20Addr) constant returns(uint256)
func (_HashDice *HashDiceSession) GetProfit(erc20Addr common.Address) (*big.Int, error) {
	return _HashDice.Contract.GetProfit(&_HashDice.CallOpts, erc20Addr)
}

// GetProfit is a free data retrieval call binding the contract method 0xccc8efcc.
//
// Solidity: function GetProfit(address erc20Addr) constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GetProfit(erc20Addr common.Address) (*big.Int, error) {
	return _HashDice.Contract.GetProfit(&_HashDice.CallOpts, erc20Addr)
}

// GetProfitRatio is a free data retrieval call binding the contract method 0x9b458bd3.
//
// Solidity: function GetProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceCaller) GetProfitRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "GetProfitRatio")
	return *ret0, err
}

// GetProfitRatio is a free data retrieval call binding the contract method 0x9b458bd3.
//
// Solidity: function GetProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceSession) GetProfitRatio() (*big.Int, error) {
	return _HashDice.Contract.GetProfitRatio(&_HashDice.CallOpts)
}

// GetProfitRatio is a free data retrieval call binding the contract method 0x9b458bd3.
//
// Solidity: function GetProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GetProfitRatio() (*big.Int, error) {
	return _HashDice.Contract.GetProfitRatio(&_HashDice.CallOpts)
}

// GetRoomInfo is a free data retrieval call binding the contract method 0x3c3419e6.
//
// Solidity: function GetRoomInfo(uint256 roomId) constant returns(address, address, string, string, uint256[4], uint256[4], bool, uint256, uint256, uint256, uint256, uint256)
func (_HashDice *HashDiceCaller) GetRoomInfo(opts *bind.CallOpts, roomId *big.Int) (common.Address, common.Address, string, string, [4]*big.Int, [4]*big.Int, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0  = new(common.Address)
		ret1  = new(common.Address)
		ret2  = new(string)
		ret3  = new(string)
		ret4  = new([4]*big.Int)
		ret5  = new([4]*big.Int)
		ret6  = new(bool)
		ret7  = new(*big.Int)
		ret8  = new(*big.Int)
		ret9  = new(*big.Int)
		ret10 = new(*big.Int)
		ret11 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
		ret10,
		ret11,
	}
	err := _HashDice.contract.Call(opts, out, "GetRoomInfo", roomId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, *ret11, err
}

// GetRoomInfo is a free data retrieval call binding the contract method 0x3c3419e6.
//
// Solidity: function GetRoomInfo(uint256 roomId) constant returns(address, address, string, string, uint256[4], uint256[4], bool, uint256, uint256, uint256, uint256, uint256)
func (_HashDice *HashDiceSession) GetRoomInfo(roomId *big.Int) (common.Address, common.Address, string, string, [4]*big.Int, [4]*big.Int, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _HashDice.Contract.GetRoomInfo(&_HashDice.CallOpts, roomId)
}

// GetRoomInfo is a free data retrieval call binding the contract method 0x3c3419e6.
//
// Solidity: function GetRoomInfo(uint256 roomId) constant returns(address, address, string, string, uint256[4], uint256[4], bool, uint256, uint256, uint256, uint256, uint256)
func (_HashDice *HashDiceCallerSession) GetRoomInfo(roomId *big.Int) (common.Address, common.Address, string, string, [4]*big.Int, [4]*big.Int, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _HashDice.Contract.GetRoomInfo(&_HashDice.CallOpts, roomId)
}

// GetRooms is a free data retrieval call binding the contract method 0xdd840951.
//
// Solidity: function GetRooms() constant returns(uint256)
func (_HashDice *HashDiceCaller) GetRooms(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "GetRooms")
	return *ret0, err
}

// GetRooms is a free data retrieval call binding the contract method 0xdd840951.
//
// Solidity: function GetRooms() constant returns(uint256)
func (_HashDice *HashDiceSession) GetRooms() (*big.Int, error) {
	return _HashDice.Contract.GetRooms(&_HashDice.CallOpts)
}

// GetRooms is a free data retrieval call binding the contract method 0xdd840951.
//
// Solidity: function GetRooms() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GetRooms() (*big.Int, error) {
	return _HashDice.Contract.GetRooms(&_HashDice.CallOpts)
}

// GetRoundPeriod is a free data retrieval call binding the contract method 0x78463323.
//
// Solidity: function GetRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceCaller) GetRoundPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "GetRoundPeriod")
	return *ret0, err
}

// GetRoundPeriod is a free data retrieval call binding the contract method 0x78463323.
//
// Solidity: function GetRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceSession) GetRoundPeriod() (*big.Int, error) {
	return _HashDice.Contract.GetRoundPeriod(&_HashDice.CallOpts)
}

// GetRoundPeriod is a free data retrieval call binding the contract method 0x78463323.
//
// Solidity: function GetRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GetRoundPeriod() (*big.Int, error) {
	return _HashDice.Contract.GetRoundPeriod(&_HashDice.CallOpts)
}

// GetRoundWait is a free data retrieval call binding the contract method 0x33d47a5b.
//
// Solidity: function GetRoundWait() constant returns(uint256)
func (_HashDice *HashDiceCaller) GetRoundWait(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "GetRoundWait")
	return *ret0, err
}

// GetRoundWait is a free data retrieval call binding the contract method 0x33d47a5b.
//
// Solidity: function GetRoundWait() constant returns(uint256)
func (_HashDice *HashDiceSession) GetRoundWait() (*big.Int, error) {
	return _HashDice.Contract.GetRoundWait(&_HashDice.CallOpts)
}

// GetRoundWait is a free data retrieval call binding the contract method 0x33d47a5b.
//
// Solidity: function GetRoundWait() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GetRoundWait() (*big.Int, error) {
	return _HashDice.Contract.GetRoundWait(&_HashDice.CallOpts)
}

// GOwner is a free data retrieval call binding the contract method 0x58fb58d8.
//
// Solidity: function gOwner() constant returns(address)
func (_HashDice *HashDiceCaller) GOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "gOwner")
	return *ret0, err
}

// GOwner is a free data retrieval call binding the contract method 0x58fb58d8.
//
// Solidity: function gOwner() constant returns(address)
func (_HashDice *HashDiceSession) GOwner() (common.Address, error) {
	return _HashDice.Contract.GOwner(&_HashDice.CallOpts)
}

// GOwner is a free data retrieval call binding the contract method 0x58fb58d8.
//
// Solidity: function gOwner() constant returns(address)
func (_HashDice *HashDiceCallerSession) GOwner() (common.Address, error) {
	return _HashDice.Contract.GOwner(&_HashDice.CallOpts)
}

// GProfitRatio is a free data retrieval call binding the contract method 0x668b0daf.
//
// Solidity: function gProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceCaller) GProfitRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "gProfitRatio")
	return *ret0, err
}

// GProfitRatio is a free data retrieval call binding the contract method 0x668b0daf.
//
// Solidity: function gProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceSession) GProfitRatio() (*big.Int, error) {
	return _HashDice.Contract.GProfitRatio(&_HashDice.CallOpts)
}

// GProfitRatio is a free data retrieval call binding the contract method 0x668b0daf.
//
// Solidity: function gProfitRatio() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GProfitRatio() (*big.Int, error) {
	return _HashDice.Contract.GProfitRatio(&_HashDice.CallOpts)
}

// GRoomId is a free data retrieval call binding the contract method 0x5ea431b2.
//
// Solidity: function gRoomId() constant returns(uint256)
func (_HashDice *HashDiceCaller) GRoomId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "gRoomId")
	return *ret0, err
}

// GRoomId is a free data retrieval call binding the contract method 0x5ea431b2.
//
// Solidity: function gRoomId() constant returns(uint256)
func (_HashDice *HashDiceSession) GRoomId() (*big.Int, error) {
	return _HashDice.Contract.GRoomId(&_HashDice.CallOpts)
}

// GRoomId is a free data retrieval call binding the contract method 0x5ea431b2.
//
// Solidity: function gRoomId() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GRoomId() (*big.Int, error) {
	return _HashDice.Contract.GRoomId(&_HashDice.CallOpts)
}

// GRoundPeriod is a free data retrieval call binding the contract method 0x5b1c48cb.
//
// Solidity: function gRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceCaller) GRoundPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "gRoundPeriod")
	return *ret0, err
}

// GRoundPeriod is a free data retrieval call binding the contract method 0x5b1c48cb.
//
// Solidity: function gRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceSession) GRoundPeriod() (*big.Int, error) {
	return _HashDice.Contract.GRoundPeriod(&_HashDice.CallOpts)
}

// GRoundPeriod is a free data retrieval call binding the contract method 0x5b1c48cb.
//
// Solidity: function gRoundPeriod() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GRoundPeriod() (*big.Int, error) {
	return _HashDice.Contract.GRoundPeriod(&_HashDice.CallOpts)
}

// GRoundWait is a free data retrieval call binding the contract method 0x0b543372.
//
// Solidity: function gRoundWait() constant returns(uint256)
func (_HashDice *HashDiceCaller) GRoundWait(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HashDice.contract.Call(opts, out, "gRoundWait")
	return *ret0, err
}

// GRoundWait is a free data retrieval call binding the contract method 0x0b543372.
//
// Solidity: function gRoundWait() constant returns(uint256)
func (_HashDice *HashDiceSession) GRoundWait() (*big.Int, error) {
	return _HashDice.Contract.GRoundWait(&_HashDice.CallOpts)
}

// GRoundWait is a free data retrieval call binding the contract method 0x0b543372.
//
// Solidity: function gRoundWait() constant returns(uint256)
func (_HashDice *HashDiceCallerSession) GRoundWait() (*big.Int, error) {
	return _HashDice.Contract.GRoundWait(&_HashDice.CallOpts)
}

// CloseBetOrder is a paid mutator transaction binding the contract method 0x5b65552f.
//
// Solidity: function CloseBetOrder(uint256 roomId, uint256 orderId) returns()
func (_HashDice *HashDiceTransactor) CloseBetOrder(opts *bind.TransactOpts, roomId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "CloseBetOrder", roomId, orderId)
}

// CloseBetOrder is a paid mutator transaction binding the contract method 0x5b65552f.
//
// Solidity: function CloseBetOrder(uint256 roomId, uint256 orderId) returns()
func (_HashDice *HashDiceSession) CloseBetOrder(roomId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseBetOrder(&_HashDice.TransactOpts, roomId, orderId)
}

// CloseBetOrder is a paid mutator transaction binding the contract method 0x5b65552f.
//
// Solidity: function CloseBetOrder(uint256 roomId, uint256 orderId) returns()
func (_HashDice *HashDiceTransactorSession) CloseBetOrder(roomId *big.Int, orderId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseBetOrder(&_HashDice.TransactOpts, roomId, orderId)
}

// CloseBetOrders is a paid mutator transaction binding the contract method 0x34e4aa49.
//
// Solidity: function CloseBetOrders(uint256 roomId) returns()
func (_HashDice *HashDiceTransactor) CloseBetOrders(opts *bind.TransactOpts, roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "CloseBetOrders", roomId)
}

// CloseBetOrders is a paid mutator transaction binding the contract method 0x34e4aa49.
//
// Solidity: function CloseBetOrders(uint256 roomId) returns()
func (_HashDice *HashDiceSession) CloseBetOrders(roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseBetOrders(&_HashDice.TransactOpts, roomId)
}

// CloseBetOrders is a paid mutator transaction binding the contract method 0x34e4aa49.
//
// Solidity: function CloseBetOrders(uint256 roomId) returns()
func (_HashDice *HashDiceTransactorSession) CloseBetOrders(roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseBetOrders(&_HashDice.TransactOpts, roomId)
}

// CloseRoom is a paid mutator transaction binding the contract method 0x1d424ea1.
//
// Solidity: function CloseRoom(uint256 roomId) returns(bool)
func (_HashDice *HashDiceTransactor) CloseRoom(opts *bind.TransactOpts, roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "CloseRoom", roomId)
}

// CloseRoom is a paid mutator transaction binding the contract method 0x1d424ea1.
//
// Solidity: function CloseRoom(uint256 roomId) returns(bool)
func (_HashDice *HashDiceSession) CloseRoom(roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseRoom(&_HashDice.TransactOpts, roomId)
}

// CloseRoom is a paid mutator transaction binding the contract method 0x1d424ea1.
//
// Solidity: function CloseRoom(uint256 roomId) returns(bool)
func (_HashDice *HashDiceTransactorSession) CloseRoom(roomId *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.CloseRoom(&_HashDice.TransactOpts, roomId)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3af609b.
//
// Solidity: function Deposit(uint256 roomId, uint256 value) returns()
func (_HashDice *HashDiceTransactor) Deposit(opts *bind.TransactOpts, roomId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "Deposit", roomId, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3af609b.
//
// Solidity: function Deposit(uint256 roomId, uint256 value) returns()
func (_HashDice *HashDiceSession) Deposit(roomId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.Deposit(&_HashDice.TransactOpts, roomId, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3af609b.
//
// Solidity: function Deposit(uint256 roomId, uint256 value) returns()
func (_HashDice *HashDiceTransactorSession) Deposit(roomId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.Deposit(&_HashDice.TransactOpts, roomId, value)
}

// OpenRoom is a paid mutator transaction binding the contract method 0xf39e5af9.
//
// Solidity: function OpenRoom(address erc20Addr, string symbol, uint256 initLockedValue, uint256 decimical, string roomName, uint256[4] nominator, uint256[4] denominator) returns(uint256 roomId)
func (_HashDice *HashDiceTransactor) OpenRoom(opts *bind.TransactOpts, erc20Addr common.Address, symbol string, initLockedValue *big.Int, decimical *big.Int, roomName string, nominator [4]*big.Int, denominator [4]*big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "OpenRoom", erc20Addr, symbol, initLockedValue, decimical, roomName, nominator, denominator)
}

// OpenRoom is a paid mutator transaction binding the contract method 0xf39e5af9.
//
// Solidity: function OpenRoom(address erc20Addr, string symbol, uint256 initLockedValue, uint256 decimical, string roomName, uint256[4] nominator, uint256[4] denominator) returns(uint256 roomId)
func (_HashDice *HashDiceSession) OpenRoom(erc20Addr common.Address, symbol string, initLockedValue *big.Int, decimical *big.Int, roomName string, nominator [4]*big.Int, denominator [4]*big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.OpenRoom(&_HashDice.TransactOpts, erc20Addr, symbol, initLockedValue, decimical, roomName, nominator, denominator)
}

// OpenRoom is a paid mutator transaction binding the contract method 0xf39e5af9.
//
// Solidity: function OpenRoom(address erc20Addr, string symbol, uint256 initLockedValue, uint256 decimical, string roomName, uint256[4] nominator, uint256[4] denominator) returns(uint256 roomId)
func (_HashDice *HashDiceTransactorSession) OpenRoom(erc20Addr common.Address, symbol string, initLockedValue *big.Int, decimical *big.Int, roomName string, nominator [4]*big.Int, denominator [4]*big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.OpenRoom(&_HashDice.TransactOpts, erc20Addr, symbol, initLockedValue, decimical, roomName, nominator, denominator)
}

// SetProfitRatio is a paid mutator transaction binding the contract method 0x8a427002.
//
// Solidity: function SetProfitRatio(uint256 ratio) returns()
func (_HashDice *HashDiceTransactor) SetProfitRatio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "SetProfitRatio", ratio)
}

// SetProfitRatio is a paid mutator transaction binding the contract method 0x8a427002.
//
// Solidity: function SetProfitRatio(uint256 ratio) returns()
func (_HashDice *HashDiceSession) SetProfitRatio(ratio *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetProfitRatio(&_HashDice.TransactOpts, ratio)
}

// SetProfitRatio is a paid mutator transaction binding the contract method 0x8a427002.
//
// Solidity: function SetProfitRatio(uint256 ratio) returns()
func (_HashDice *HashDiceTransactorSession) SetProfitRatio(ratio *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetProfitRatio(&_HashDice.TransactOpts, ratio)
}

// SetRoundPeriod is a paid mutator transaction binding the contract method 0x0df298d1.
//
// Solidity: function SetRoundPeriod(uint256 value) returns()
func (_HashDice *HashDiceTransactor) SetRoundPeriod(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "SetRoundPeriod", value)
}

// SetRoundPeriod is a paid mutator transaction binding the contract method 0x0df298d1.
//
// Solidity: function SetRoundPeriod(uint256 value) returns()
func (_HashDice *HashDiceSession) SetRoundPeriod(value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetRoundPeriod(&_HashDice.TransactOpts, value)
}

// SetRoundPeriod is a paid mutator transaction binding the contract method 0x0df298d1.
//
// Solidity: function SetRoundPeriod(uint256 value) returns()
func (_HashDice *HashDiceTransactorSession) SetRoundPeriod(value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetRoundPeriod(&_HashDice.TransactOpts, value)
}

// SetRoundWait is a paid mutator transaction binding the contract method 0x876148ac.
//
// Solidity: function SetRoundWait(uint256 value) returns()
func (_HashDice *HashDiceTransactor) SetRoundWait(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "SetRoundWait", value)
}

// SetRoundWait is a paid mutator transaction binding the contract method 0x876148ac.
//
// Solidity: function SetRoundWait(uint256 value) returns()
func (_HashDice *HashDiceSession) SetRoundWait(value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetRoundWait(&_HashDice.TransactOpts, value)
}

// SetRoundWait is a paid mutator transaction binding the contract method 0x876148ac.
//
// Solidity: function SetRoundWait(uint256 value) returns()
func (_HashDice *HashDiceTransactorSession) SetRoundWait(value *big.Int) (*types.Transaction, error) {
	return _HashDice.Contract.SetRoundWait(&_HashDice.TransactOpts, value)
}

// SubmitBetOrder is a paid mutator transaction binding the contract method 0x2c412f6f.
//
// Solidity: function SubmitBetOrder(uint256 roomId, uint256 startBlock, uint8 betType, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceTransactor) SubmitBetOrder(opts *bind.TransactOpts, roomId *big.Int, startBlock *big.Int, betType uint8, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "SubmitBetOrder", roomId, startBlock, betType, betValue)
}

// SubmitBetOrder is a paid mutator transaction binding the contract method 0x2c412f6f.
//
// Solidity: function SubmitBetOrder(uint256 roomId, uint256 startBlock, uint8 betType, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceSession) SubmitBetOrder(roomId *big.Int, startBlock *big.Int, betType uint8, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.Contract.SubmitBetOrder(&_HashDice.TransactOpts, roomId, startBlock, betType, betValue)
}

// SubmitBetOrder is a paid mutator transaction binding the contract method 0x2c412f6f.
//
// Solidity: function SubmitBetOrder(uint256 roomId, uint256 startBlock, uint8 betType, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceTransactorSession) SubmitBetOrder(roomId *big.Int, startBlock *big.Int, betType uint8, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.Contract.SubmitBetOrder(&_HashDice.TransactOpts, roomId, startBlock, betType, betValue)
}

// SubmitBetOrderType0 is a paid mutator transaction binding the contract method 0x722dd1c6.
//
// Solidity: function SubmitBetOrderType0(uint256 roomId, uint256 startBlock, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceTransactor) SubmitBetOrderType0(opts *bind.TransactOpts, roomId *big.Int, startBlock *big.Int, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "SubmitBetOrderType0", roomId, startBlock, betValue)
}

// SubmitBetOrderType0 is a paid mutator transaction binding the contract method 0x722dd1c6.
//
// Solidity: function SubmitBetOrderType0(uint256 roomId, uint256 startBlock, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceSession) SubmitBetOrderType0(roomId *big.Int, startBlock *big.Int, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.Contract.SubmitBetOrderType0(&_HashDice.TransactOpts, roomId, startBlock, betValue)
}

// SubmitBetOrderType0 is a paid mutator transaction binding the contract method 0x722dd1c6.
//
// Solidity: function SubmitBetOrderType0(uint256 roomId, uint256 startBlock, uint32[] betValue) returns(uint256)
func (_HashDice *HashDiceTransactorSession) SubmitBetOrderType0(roomId *big.Int, startBlock *big.Int, betValue []uint32) (*types.Transaction, error) {
	return _HashDice.Contract.SubmitBetOrderType0(&_HashDice.TransactOpts, roomId, startBlock, betValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71ef96c4.
//
// Solidity: function Withdraw(uint256 roomId, uint256 value, address receiver) returns()
func (_HashDice *HashDiceTransactor) Withdraw(opts *bind.TransactOpts, roomId *big.Int, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "Withdraw", roomId, value, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71ef96c4.
//
// Solidity: function Withdraw(uint256 roomId, uint256 value, address receiver) returns()
func (_HashDice *HashDiceSession) Withdraw(roomId *big.Int, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.Contract.Withdraw(&_HashDice.TransactOpts, roomId, value, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x71ef96c4.
//
// Solidity: function Withdraw(uint256 roomId, uint256 value, address receiver) returns()
func (_HashDice *HashDiceTransactorSession) Withdraw(roomId *big.Int, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.Contract.Withdraw(&_HashDice.TransactOpts, roomId, value, receiver)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0xc2ee9822.
//
// Solidity: function WithdrawProfit(address erc20Addr, uint256 value, address receiver) returns()
func (_HashDice *HashDiceTransactor) WithdrawProfit(opts *bind.TransactOpts, erc20Addr common.Address, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.contract.Transact(opts, "WithdrawProfit", erc20Addr, value, receiver)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0xc2ee9822.
//
// Solidity: function WithdrawProfit(address erc20Addr, uint256 value, address receiver) returns()
func (_HashDice *HashDiceSession) WithdrawProfit(erc20Addr common.Address, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.Contract.WithdrawProfit(&_HashDice.TransactOpts, erc20Addr, value, receiver)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0xc2ee9822.
//
// Solidity: function WithdrawProfit(address erc20Addr, uint256 value, address receiver) returns()
func (_HashDice *HashDiceTransactorSession) WithdrawProfit(erc20Addr common.Address, value *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _HashDice.Contract.WithdrawProfit(&_HashDice.TransactOpts, erc20Addr, value, receiver)
}

// HashDiceCloseRoundTooLateIterator is returned from FilterCloseRoundTooLate and is used to iterate over the raw logs and unpacked data for CloseRoundTooLate events raised by the HashDice contract.
type HashDiceCloseRoundTooLateIterator struct {
	Event *HashDiceCloseRoundTooLate // Event containing the contract specifics and raw log

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
func (it *HashDiceCloseRoundTooLateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceCloseRoundTooLate)
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
		it.Event = new(HashDiceCloseRoundTooLate)
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
func (it *HashDiceCloseRoundTooLateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceCloseRoundTooLateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceCloseRoundTooLate represents a CloseRoundTooLate event raised by the HashDice contract.
type HashDiceCloseRoundTooLate struct {
	RoomId     *big.Int
	OrderId    *big.Int
	BlockNum   *big.Int
	StartBlock *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCloseRoundTooLate is a free log retrieval operation binding the contract event 0x84842e1cd70dcda5b4f0df63bcd2e252233f6e5419d41793fd75645b14f12087.
//
// Solidity: event CloseRoundTooLate(uint256 roomId, uint256 orderId, uint256 blockNum, uint256 startBlock)
func (_HashDice *HashDiceFilterer) FilterCloseRoundTooLate(opts *bind.FilterOpts) (*HashDiceCloseRoundTooLateIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "CloseRoundTooLate")
	if err != nil {
		return nil, err
	}
	return &HashDiceCloseRoundTooLateIterator{contract: _HashDice.contract, event: "CloseRoundTooLate", logs: logs, sub: sub}, nil
}

// WatchCloseRoundTooLate is a free log subscription operation binding the contract event 0x84842e1cd70dcda5b4f0df63bcd2e252233f6e5419d41793fd75645b14f12087.
//
// Solidity: event CloseRoundTooLate(uint256 roomId, uint256 orderId, uint256 blockNum, uint256 startBlock)
func (_HashDice *HashDiceFilterer) WatchCloseRoundTooLate(opts *bind.WatchOpts, sink chan<- *HashDiceCloseRoundTooLate) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "CloseRoundTooLate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceCloseRoundTooLate)
				if err := _HashDice.contract.UnpackLog(event, "CloseRoundTooLate", log); err != nil {
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

// HashDiceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the HashDice contract.
type HashDiceDepositedIterator struct {
	Event *HashDiceDeposited // Event containing the contract specifics and raw log

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
func (it *HashDiceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceDeposited)
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
		it.Event = new(HashDiceDeposited)
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
func (it *HashDiceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceDeposited represents a Deposited event raised by the HashDice contract.
type HashDiceDeposited struct {
	RoomId   *big.Int
	Value    *big.Int
	Remained *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1ca606821992e3b34613b5b29c0bbade3a907b2969d7f9f2927f726fa4baccfb.
//
// Solidity: event Deposited(uint256 roomId, uint256 value, uint256 remained)
func (_HashDice *HashDiceFilterer) FilterDeposited(opts *bind.FilterOpts) (*HashDiceDepositedIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &HashDiceDepositedIterator{contract: _HashDice.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1ca606821992e3b34613b5b29c0bbade3a907b2969d7f9f2927f726fa4baccfb.
//
// Solidity: event Deposited(uint256 roomId, uint256 value, uint256 remained)
func (_HashDice *HashDiceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *HashDiceDeposited) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceDeposited)
				if err := _HashDice.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// HashDiceNewBetOrderIterator is returned from FilterNewBetOrder and is used to iterate over the raw logs and unpacked data for NewBetOrder events raised by the HashDice contract.
type HashDiceNewBetOrderIterator struct {
	Event *HashDiceNewBetOrder // Event containing the contract specifics and raw log

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
func (it *HashDiceNewBetOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceNewBetOrder)
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
		it.Event = new(HashDiceNewBetOrder)
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
func (it *HashDiceNewBetOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceNewBetOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceNewBetOrder represents a NewBetOrder event raised by the HashDice contract.
type HashDiceNewBetOrder struct {
	RoomId  *big.Int
	OrderId *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBetOrder is a free log retrieval operation binding the contract event 0x3bcc29ebfaf3c11e8448a4d3659f7e312a211f4d9979bcbdbb42923150a4918d.
//
// Solidity: event NewBetOrder(uint256 roomId, uint256 orderId, uint256 value)
func (_HashDice *HashDiceFilterer) FilterNewBetOrder(opts *bind.FilterOpts) (*HashDiceNewBetOrderIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "NewBetOrder")
	if err != nil {
		return nil, err
	}
	return &HashDiceNewBetOrderIterator{contract: _HashDice.contract, event: "NewBetOrder", logs: logs, sub: sub}, nil
}

// WatchNewBetOrder is a free log subscription operation binding the contract event 0x3bcc29ebfaf3c11e8448a4d3659f7e312a211f4d9979bcbdbb42923150a4918d.
//
// Solidity: event NewBetOrder(uint256 roomId, uint256 orderId, uint256 value)
func (_HashDice *HashDiceFilterer) WatchNewBetOrder(opts *bind.WatchOpts, sink chan<- *HashDiceNewBetOrder) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "NewBetOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceNewBetOrder)
				if err := _HashDice.contract.UnpackLog(event, "NewBetOrder", log); err != nil {
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

// HashDicePayBetOwnerIterator is returned from FilterPayBetOwner and is used to iterate over the raw logs and unpacked data for PayBetOwner events raised by the HashDice contract.
type HashDicePayBetOwnerIterator struct {
	Event *HashDicePayBetOwner // Event containing the contract specifics and raw log

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
func (it *HashDicePayBetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDicePayBetOwner)
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
		it.Event = new(HashDicePayBetOwner)
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
func (it *HashDicePayBetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDicePayBetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDicePayBetOwner represents a PayBetOwner event raised by the HashDice contract.
type HashDicePayBetOwner struct {
	RoomId  *big.Int
	OrderId *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayBetOwner is a free log retrieval operation binding the contract event 0xc6463bc3c0665fea9e0be72485dd8124a9b05c1632ffa17f1935722c5851cdfe.
//
// Solidity: event PayBetOwner(uint256 roomId, uint256 orderId, uint256 value)
func (_HashDice *HashDiceFilterer) FilterPayBetOwner(opts *bind.FilterOpts) (*HashDicePayBetOwnerIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "PayBetOwner")
	if err != nil {
		return nil, err
	}
	return &HashDicePayBetOwnerIterator{contract: _HashDice.contract, event: "PayBetOwner", logs: logs, sub: sub}, nil
}

// WatchPayBetOwner is a free log subscription operation binding the contract event 0xc6463bc3c0665fea9e0be72485dd8124a9b05c1632ffa17f1935722c5851cdfe.
//
// Solidity: event PayBetOwner(uint256 roomId, uint256 orderId, uint256 value)
func (_HashDice *HashDiceFilterer) WatchPayBetOwner(opts *bind.WatchOpts, sink chan<- *HashDicePayBetOwner) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "PayBetOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDicePayBetOwner)
				if err := _HashDice.contract.UnpackLog(event, "PayBetOwner", log); err != nil {
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

// HashDiceRoomClosedIterator is returned from FilterRoomClosed and is used to iterate over the raw logs and unpacked data for RoomClosed events raised by the HashDice contract.
type HashDiceRoomClosedIterator struct {
	Event *HashDiceRoomClosed // Event containing the contract specifics and raw log

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
func (it *HashDiceRoomClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceRoomClosed)
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
		it.Event = new(HashDiceRoomClosed)
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
func (it *HashDiceRoomClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceRoomClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceRoomClosed represents a RoomClosed event raised by the HashDice contract.
type HashDiceRoomClosed struct {
	RoomId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoomClosed is a free log retrieval operation binding the contract event 0x67acaceb58ff1ad2aa8b3069f69c4cd076824e06a09184b75fb197ead7040be7.
//
// Solidity: event RoomClosed(uint256 roomId)
func (_HashDice *HashDiceFilterer) FilterRoomClosed(opts *bind.FilterOpts) (*HashDiceRoomClosedIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "RoomClosed")
	if err != nil {
		return nil, err
	}
	return &HashDiceRoomClosedIterator{contract: _HashDice.contract, event: "RoomClosed", logs: logs, sub: sub}, nil
}

// WatchRoomClosed is a free log subscription operation binding the contract event 0x67acaceb58ff1ad2aa8b3069f69c4cd076824e06a09184b75fb197ead7040be7.
//
// Solidity: event RoomClosed(uint256 roomId)
func (_HashDice *HashDiceFilterer) WatchRoomClosed(opts *bind.WatchOpts, sink chan<- *HashDiceRoomClosed) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "RoomClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceRoomClosed)
				if err := _HashDice.contract.UnpackLog(event, "RoomClosed", log); err != nil {
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

// HashDiceRoomOpenedIterator is returned from FilterRoomOpened and is used to iterate over the raw logs and unpacked data for RoomOpened events raised by the HashDice contract.
type HashDiceRoomOpenedIterator struct {
	Event *HashDiceRoomOpened // Event containing the contract specifics and raw log

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
func (it *HashDiceRoomOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceRoomOpened)
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
		it.Event = new(HashDiceRoomOpened)
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
func (it *HashDiceRoomOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceRoomOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceRoomOpened represents a RoomOpened event raised by the HashDice contract.
type HashDiceRoomOpened struct {
	Creator         common.Address
	Erc20Addr       common.Address
	Symbol          string
	Name            string
	Id              *big.Int
	Decimical       *big.Int
	LastLockedValue *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRoomOpened is a free log retrieval operation binding the contract event 0x2000b9717d942bac323f385a26f6b20bddb613171cc467b2b73f6817e86846f9.
//
// Solidity: event RoomOpened(address creator, address erc20Addr, string symbol, string name, uint256 id, uint256 decimical, uint256 lastLockedValue)
func (_HashDice *HashDiceFilterer) FilterRoomOpened(opts *bind.FilterOpts) (*HashDiceRoomOpenedIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "RoomOpened")
	if err != nil {
		return nil, err
	}
	return &HashDiceRoomOpenedIterator{contract: _HashDice.contract, event: "RoomOpened", logs: logs, sub: sub}, nil
}

// WatchRoomOpened is a free log subscription operation binding the contract event 0x2000b9717d942bac323f385a26f6b20bddb613171cc467b2b73f6817e86846f9.
//
// Solidity: event RoomOpened(address creator, address erc20Addr, string symbol, string name, uint256 id, uint256 decimical, uint256 lastLockedValue)
func (_HashDice *HashDiceFilterer) WatchRoomOpened(opts *bind.WatchOpts, sink chan<- *HashDiceRoomOpened) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "RoomOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceRoomOpened)
				if err := _HashDice.contract.UnpackLog(event, "RoomOpened", log); err != nil {
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

// HashDiceWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the HashDice contract.
type HashDiceWithdrawedIterator struct {
	Event *HashDiceWithdrawed // Event containing the contract specifics and raw log

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
func (it *HashDiceWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashDiceWithdrawed)
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
		it.Event = new(HashDiceWithdrawed)
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
func (it *HashDiceWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashDiceWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashDiceWithdrawed represents a Withdrawed event raised by the HashDice contract.
type HashDiceWithdrawed struct {
	RoomId   *big.Int
	Value    *big.Int
	Remained *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0xe492f03a13f16d9ab77da04e06cc43e6eb0a495bab6bc3d0018b7d6714e6359f.
//
// Solidity: event Withdrawed(uint256 roomId, uint256 value, uint256 remained)
func (_HashDice *HashDiceFilterer) FilterWithdrawed(opts *bind.FilterOpts) (*HashDiceWithdrawedIterator, error) {

	logs, sub, err := _HashDice.contract.FilterLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return &HashDiceWithdrawedIterator{contract: _HashDice.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0xe492f03a13f16d9ab77da04e06cc43e6eb0a495bab6bc3d0018b7d6714e6359f.
//
// Solidity: event Withdrawed(uint256 roomId, uint256 value, uint256 remained)
func (_HashDice *HashDiceFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *HashDiceWithdrawed) (event.Subscription, error) {

	logs, sub, err := _HashDice.contract.WatchLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashDiceWithdrawed)
				if err := _HashDice.contract.UnpackLog(event, "Withdrawed", log); err != nil {
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
