// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recipient

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RecipientMetaData contains all meta data concerning the Recipient contract.
var RecipientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceivedCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceivedMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"fooBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainSecurityModule\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCallMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"setInterchainSecurityModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107d48061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c8063715018a611610071578063715018a6146101355780638da5cb5b1461013d578063a4982fde1461014e578063de523cf314610156578063f07c1f4714610169578063f2fde38b1461017c57600080fd5b80626e75ec146100ad5780630e72cc06146100cb5780632113522a146100e0578063256fec881461010b57806356d5d47514610122575b600080fd5b6100b561018f565b6040516100c2919061049a565b60405180910390f35b6100de6100d93660046104b4565b61021d565b005b6004546100f3906001600160a01b031681565b6040516001600160a01b0390911681526020016100c2565b61011460025481565b6040519081526020016100c2565b6100de610130366004610526565b610247565b6100de6102a1565b6000546001600160a01b03166100f3565b6100b56102b5565b6001546100f3906001600160a01b031681565b6100de61017736600461058b565b6102c2565b6100de61018a3660046104b4565b61032c565b6003805461019c906105d7565b80601f01602080910402602001604051908101604052809291908181526020018280546101c8906105d7565b80156102155780601f106101ea57610100808354040283529160200191610215565b820191906000526020600020905b8154815290600101906020018083116101f857829003601f168201915b505050505081565b6102256103aa565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b828463ffffffff167fba67744c899113a84f615d51af5d82f5fedcf26c9a474d9363c3ad9b0bd501ac848460405161028092919061063a565b60405180910390a36002839055600361029a8284836106bb565b5050505050565b6102a96103aa565b6102b36000610404565b565b6005805461019c906105d7565b336001600160a01b03167f97d8367a1f39eb9e97f262fafbb05925c0bcfe120aaad7b9737cae34f749c2068484846040516102ff9392919061077b565b60405180910390a2600480546001600160a01b0319163317905560056103268284836106bb565b50505050565b6103346103aa565b6001600160a01b03811661039e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6103a781610404565b50565b6000546001600160a01b031633146102b35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610395565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000815180845260005b8181101561047a5760208185018101518683018201520161045e565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006104ad6020830184610454565b9392505050565b6000602082840312156104c657600080fd5b81356001600160a01b03811681146104ad57600080fd5b60008083601f8401126104ef57600080fd5b50813567ffffffffffffffff81111561050757600080fd5b60208301915083602082850101111561051f57600080fd5b9250929050565b6000806000806060858703121561053c57600080fd5b843563ffffffff8116811461055057600080fd5b935060208501359250604085013567ffffffffffffffff81111561057357600080fd5b61057f878288016104dd565b95989497509550505050565b6000806000604084860312156105a057600080fd5b83359250602084013567ffffffffffffffff8111156105be57600080fd5b6105ca868287016104dd565b9497909650939450505050565b600181811c908216806105eb57607f821691505b60208210810361060b57634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061064e602083018486610611565b949350505050565b634e487b7160e01b600052604160045260246000fd5b601f8211156106b657600081815260208120601f850160051c810160208610156106935750805b601f850160051c820191505b818110156106b25782815560010161069f565b5050505b505050565b67ffffffffffffffff8311156106d3576106d3610656565b6106e7836106e183546105d7565b8361066c565b6000601f84116001811461071b57600085156107035750838201355b600019600387901b1c1916600186901b17835561029a565b600083815260209020601f19861690835b8281101561074c578685013582556020948501946001909201910161072c565b50868210156107695760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b838152604060208201526000610795604083018486610611565b9594505050505056fea26469706673582212201fdad42f1155ab88eeb040c2c6d8bb3eb7f9ca7d9eca723cda08c1e7ecc4e58264736f6c63430008130033",
}

// RecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use RecipientMetaData.ABI instead.
var RecipientABI = RecipientMetaData.ABI

// RecipientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecipientMetaData.Bin instead.
var RecipientBin = RecipientMetaData.Bin

// DeployRecipient deploys a new Ethereum contract, binding an instance of Recipient to it.
func DeployRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recipient, error) {
	parsed, err := RecipientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recipient{RecipientCaller: RecipientCaller{contract: contract}, RecipientTransactor: RecipientTransactor{contract: contract}, RecipientFilterer: RecipientFilterer{contract: contract}}, nil
}

// Recipient is an auto generated Go binding around an Ethereum contract.
type Recipient struct {
	RecipientCaller     // Read-only binding to the contract
	RecipientTransactor // Write-only binding to the contract
	RecipientFilterer   // Log filterer for contract events
}

// RecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecipientSession struct {
	Contract     *Recipient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecipientCallerSession struct {
	Contract *RecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecipientTransactorSession struct {
	Contract     *RecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecipientRaw struct {
	Contract *Recipient // Generic contract binding to access the raw methods on
}

// RecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecipientCallerRaw struct {
	Contract *RecipientCaller // Generic read-only contract binding to access the raw methods on
}

// RecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecipientTransactorRaw struct {
	Contract *RecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecipient creates a new instance of Recipient, bound to a specific deployed contract.
func NewRecipient(address common.Address, backend bind.ContractBackend) (*Recipient, error) {
	contract, err := bindRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recipient{RecipientCaller: RecipientCaller{contract: contract}, RecipientTransactor: RecipientTransactor{contract: contract}, RecipientFilterer: RecipientFilterer{contract: contract}}, nil
}

// NewRecipientCaller creates a new read-only instance of Recipient, bound to a specific deployed contract.
func NewRecipientCaller(address common.Address, caller bind.ContractCaller) (*RecipientCaller, error) {
	contract, err := bindRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecipientCaller{contract: contract}, nil
}

// NewRecipientTransactor creates a new write-only instance of Recipient, bound to a specific deployed contract.
func NewRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*RecipientTransactor, error) {
	contract, err := bindRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecipientTransactor{contract: contract}, nil
}

// NewRecipientFilterer creates a new log filterer instance of Recipient, bound to a specific deployed contract.
func NewRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*RecipientFilterer, error) {
	contract, err := bindRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecipientFilterer{contract: contract}, nil
}

// bindRecipient binds a generic wrapper to an already deployed contract.
func bindRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecipientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recipient *RecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recipient.Contract.RecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recipient *RecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recipient.Contract.RecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recipient *RecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recipient.Contract.RecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recipient *RecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recipient *RecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recipient *RecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recipient.Contract.contract.Transact(opts, method, params...)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Recipient *RecipientCaller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Recipient *RecipientSession) InterchainSecurityModule() (common.Address, error) {
	return _Recipient.Contract.InterchainSecurityModule(&_Recipient.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Recipient *RecipientCallerSession) InterchainSecurityModule() (common.Address, error) {
	return _Recipient.Contract.InterchainSecurityModule(&_Recipient.CallOpts)
}

// LastCallMessage is a free data retrieval call binding the contract method 0xa4982fde.
//
// Solidity: function lastCallMessage() view returns(string)
func (_Recipient *RecipientCaller) LastCallMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "lastCallMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LastCallMessage is a free data retrieval call binding the contract method 0xa4982fde.
//
// Solidity: function lastCallMessage() view returns(string)
func (_Recipient *RecipientSession) LastCallMessage() (string, error) {
	return _Recipient.Contract.LastCallMessage(&_Recipient.CallOpts)
}

// LastCallMessage is a free data retrieval call binding the contract method 0xa4982fde.
//
// Solidity: function lastCallMessage() view returns(string)
func (_Recipient *RecipientCallerSession) LastCallMessage() (string, error) {
	return _Recipient.Contract.LastCallMessage(&_Recipient.CallOpts)
}

// LastCaller is a free data retrieval call binding the contract method 0x2113522a.
//
// Solidity: function lastCaller() view returns(address)
func (_Recipient *RecipientCaller) LastCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "lastCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastCaller is a free data retrieval call binding the contract method 0x2113522a.
//
// Solidity: function lastCaller() view returns(address)
func (_Recipient *RecipientSession) LastCaller() (common.Address, error) {
	return _Recipient.Contract.LastCaller(&_Recipient.CallOpts)
}

// LastCaller is a free data retrieval call binding the contract method 0x2113522a.
//
// Solidity: function lastCaller() view returns(address)
func (_Recipient *RecipientCallerSession) LastCaller() (common.Address, error) {
	return _Recipient.Contract.LastCaller(&_Recipient.CallOpts)
}

// LastData is a free data retrieval call binding the contract method 0x006e75ec.
//
// Solidity: function lastData() view returns(bytes)
func (_Recipient *RecipientCaller) LastData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "lastData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LastData is a free data retrieval call binding the contract method 0x006e75ec.
//
// Solidity: function lastData() view returns(bytes)
func (_Recipient *RecipientSession) LastData() ([]byte, error) {
	return _Recipient.Contract.LastData(&_Recipient.CallOpts)
}

// LastData is a free data retrieval call binding the contract method 0x006e75ec.
//
// Solidity: function lastData() view returns(bytes)
func (_Recipient *RecipientCallerSession) LastData() ([]byte, error) {
	return _Recipient.Contract.LastData(&_Recipient.CallOpts)
}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(bytes32)
func (_Recipient *RecipientCaller) LastSender(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "lastSender")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(bytes32)
func (_Recipient *RecipientSession) LastSender() ([32]byte, error) {
	return _Recipient.Contract.LastSender(&_Recipient.CallOpts)
}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(bytes32)
func (_Recipient *RecipientCallerSession) LastSender() ([32]byte, error) {
	return _Recipient.Contract.LastSender(&_Recipient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recipient *RecipientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recipient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recipient *RecipientSession) Owner() (common.Address, error) {
	return _Recipient.Contract.Owner(&_Recipient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recipient *RecipientCallerSession) Owner() (common.Address, error) {
	return _Recipient.Contract.Owner(&_Recipient.CallOpts)
}

// FooBar is a paid mutator transaction binding the contract method 0xf07c1f47.
//
// Solidity: function fooBar(uint256 amount, string message) returns()
func (_Recipient *RecipientTransactor) FooBar(opts *bind.TransactOpts, amount *big.Int, message string) (*types.Transaction, error) {
	return _Recipient.contract.Transact(opts, "fooBar", amount, message)
}

// FooBar is a paid mutator transaction binding the contract method 0xf07c1f47.
//
// Solidity: function fooBar(uint256 amount, string message) returns()
func (_Recipient *RecipientSession) FooBar(amount *big.Int, message string) (*types.Transaction, error) {
	return _Recipient.Contract.FooBar(&_Recipient.TransactOpts, amount, message)
}

// FooBar is a paid mutator transaction binding the contract method 0xf07c1f47.
//
// Solidity: function fooBar(uint256 amount, string message) returns()
func (_Recipient *RecipientTransactorSession) FooBar(amount *big.Int, message string) (*types.Transaction, error) {
	return _Recipient.Contract.FooBar(&_Recipient.TransactOpts, amount, message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _data) returns()
func (_Recipient *RecipientTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _sender [32]byte, _data []byte) (*types.Transaction, error) {
	return _Recipient.contract.Transact(opts, "handle", _origin, _sender, _data)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _data) returns()
func (_Recipient *RecipientSession) Handle(_origin uint32, _sender [32]byte, _data []byte) (*types.Transaction, error) {
	return _Recipient.Contract.Handle(&_Recipient.TransactOpts, _origin, _sender, _data)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _data) returns()
func (_Recipient *RecipientTransactorSession) Handle(_origin uint32, _sender [32]byte, _data []byte) (*types.Transaction, error) {
	return _Recipient.Contract.Handle(&_Recipient.TransactOpts, _origin, _sender, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Recipient *RecipientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recipient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Recipient *RecipientSession) RenounceOwnership() (*types.Transaction, error) {
	return _Recipient.Contract.RenounceOwnership(&_Recipient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Recipient *RecipientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Recipient.Contract.RenounceOwnership(&_Recipient.TransactOpts)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _ism) returns()
func (_Recipient *RecipientTransactor) SetInterchainSecurityModule(opts *bind.TransactOpts, _ism common.Address) (*types.Transaction, error) {
	return _Recipient.contract.Transact(opts, "setInterchainSecurityModule", _ism)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _ism) returns()
func (_Recipient *RecipientSession) SetInterchainSecurityModule(_ism common.Address) (*types.Transaction, error) {
	return _Recipient.Contract.SetInterchainSecurityModule(&_Recipient.TransactOpts, _ism)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _ism) returns()
func (_Recipient *RecipientTransactorSession) SetInterchainSecurityModule(_ism common.Address) (*types.Transaction, error) {
	return _Recipient.Contract.SetInterchainSecurityModule(&_Recipient.TransactOpts, _ism)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recipient *RecipientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Recipient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recipient *RecipientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Recipient.Contract.TransferOwnership(&_Recipient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Recipient *RecipientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Recipient.Contract.TransferOwnership(&_Recipient.TransactOpts, newOwner)
}

// RecipientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Recipient contract.
type RecipientOwnershipTransferredIterator struct {
	Event *RecipientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RecipientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecipientOwnershipTransferred)
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
		it.Event = new(RecipientOwnershipTransferred)
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
func (it *RecipientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecipientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecipientOwnershipTransferred represents a OwnershipTransferred event raised by the Recipient contract.
type RecipientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Recipient *RecipientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RecipientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Recipient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RecipientOwnershipTransferredIterator{contract: _Recipient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Recipient *RecipientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecipientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Recipient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecipientOwnershipTransferred)
				if err := _Recipient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Recipient *RecipientFilterer) ParseOwnershipTransferred(log types.Log) (*RecipientOwnershipTransferred, error) {
	event := new(RecipientOwnershipTransferred)
	if err := _Recipient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecipientReceivedCallIterator is returned from FilterReceivedCall and is used to iterate over the raw logs and unpacked data for ReceivedCall events raised by the Recipient contract.
type RecipientReceivedCallIterator struct {
	Event *RecipientReceivedCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RecipientReceivedCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecipientReceivedCall)
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
		it.Event = new(RecipientReceivedCall)
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
func (it *RecipientReceivedCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecipientReceivedCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecipientReceivedCall represents a ReceivedCall event raised by the Recipient contract.
type RecipientReceivedCall struct {
	Caller  common.Address
	Amount  *big.Int
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedCall is a free log retrieval operation binding the contract event 0x97d8367a1f39eb9e97f262fafbb05925c0bcfe120aaad7b9737cae34f749c206.
//
// Solidity: event ReceivedCall(address indexed caller, uint256 amount, string message)
func (_Recipient *RecipientFilterer) FilterReceivedCall(opts *bind.FilterOpts, caller []common.Address) (*RecipientReceivedCallIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Recipient.contract.FilterLogs(opts, "ReceivedCall", callerRule)
	if err != nil {
		return nil, err
	}
	return &RecipientReceivedCallIterator{contract: _Recipient.contract, event: "ReceivedCall", logs: logs, sub: sub}, nil
}

// WatchReceivedCall is a free log subscription operation binding the contract event 0x97d8367a1f39eb9e97f262fafbb05925c0bcfe120aaad7b9737cae34f749c206.
//
// Solidity: event ReceivedCall(address indexed caller, uint256 amount, string message)
func (_Recipient *RecipientFilterer) WatchReceivedCall(opts *bind.WatchOpts, sink chan<- *RecipientReceivedCall, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Recipient.contract.WatchLogs(opts, "ReceivedCall", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecipientReceivedCall)
				if err := _Recipient.contract.UnpackLog(event, "ReceivedCall", log); err != nil {
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

// ParseReceivedCall is a log parse operation binding the contract event 0x97d8367a1f39eb9e97f262fafbb05925c0bcfe120aaad7b9737cae34f749c206.
//
// Solidity: event ReceivedCall(address indexed caller, uint256 amount, string message)
func (_Recipient *RecipientFilterer) ParseReceivedCall(log types.Log) (*RecipientReceivedCall, error) {
	event := new(RecipientReceivedCall)
	if err := _Recipient.contract.UnpackLog(event, "ReceivedCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecipientReceivedMessageIterator is returned from FilterReceivedMessage and is used to iterate over the raw logs and unpacked data for ReceivedMessage events raised by the Recipient contract.
type RecipientReceivedMessageIterator struct {
	Event *RecipientReceivedMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RecipientReceivedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecipientReceivedMessage)
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
		it.Event = new(RecipientReceivedMessage)
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
func (it *RecipientReceivedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecipientReceivedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecipientReceivedMessage represents a ReceivedMessage event raised by the Recipient contract.
type RecipientReceivedMessage struct {
	Origin  uint32
	Sender  [32]byte
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedMessage is a free log retrieval operation binding the contract event 0xba67744c899113a84f615d51af5d82f5fedcf26c9a474d9363c3ad9b0bd501ac.
//
// Solidity: event ReceivedMessage(uint32 indexed origin, bytes32 indexed sender, string message)
func (_Recipient *RecipientFilterer) FilterReceivedMessage(opts *bind.FilterOpts, origin []uint32, sender [][32]byte) (*RecipientReceivedMessageIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Recipient.contract.FilterLogs(opts, "ReceivedMessage", originRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RecipientReceivedMessageIterator{contract: _Recipient.contract, event: "ReceivedMessage", logs: logs, sub: sub}, nil
}

// WatchReceivedMessage is a free log subscription operation binding the contract event 0xba67744c899113a84f615d51af5d82f5fedcf26c9a474d9363c3ad9b0bd501ac.
//
// Solidity: event ReceivedMessage(uint32 indexed origin, bytes32 indexed sender, string message)
func (_Recipient *RecipientFilterer) WatchReceivedMessage(opts *bind.WatchOpts, sink chan<- *RecipientReceivedMessage, origin []uint32, sender [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Recipient.contract.WatchLogs(opts, "ReceivedMessage", originRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecipientReceivedMessage)
				if err := _Recipient.contract.UnpackLog(event, "ReceivedMessage", log); err != nil {
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

// ParseReceivedMessage is a log parse operation binding the contract event 0xba67744c899113a84f615d51af5d82f5fedcf26c9a474d9363c3ad9b0bd501ac.
//
// Solidity: event ReceivedMessage(uint32 indexed origin, bytes32 indexed sender, string message)
func (_Recipient *RecipientFilterer) ParseReceivedMessage(log types.Log) (*RecipientReceivedMessage, error) {
	event := new(RecipientReceivedMessage)
	if err := _Recipient.contract.UnpackLog(event, "ReceivedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
