// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package announce

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

// AnnounceMetaData contains all meta data concerning the Announce contract.
var AnnounceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mailbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageLocation\",\"type\":\"string\"}],\"name\":\"ValidatorAnnouncement\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_storageLocation\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"announce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getAnnouncedStorageLocations\",\"outputs\":[{\"internalType\":\"string[][]\",\"name\":\"\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAnnouncedValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610fc3380380610fc383398101604081905261002f916100ac565b6001600160a01b03811660808190526040805163234d8e3d60e21b81529051638d3638f4916004808201926020929091908290030181865afa158015610079573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061009d91906100dc565b63ffffffff1660a05250610102565b6000602082840312156100be57600080fd5b81516001600160a01b03811681146100d557600080fd5b9392505050565b6000602082840312156100ee57600080fd5b815163ffffffff811681146100d557600080fd5b60805160a051610e906101336000396000818160be015261020b01526000818160fa01526101e90152610e906000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806321f717811461005c57806351abe7cc14610084578063690cb786146100a45780638d3638f4146100b9578063d5438eae146100f5575b600080fd5b61006f61006a3660046109eb565b610134565b60405190151581526020015b60405180910390f35b610097610092366004610a6c565b610381565b60405161007b9190610b05565b6100ac610534565b60405161007b9190610bca565b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161007b565b61011c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161007b565b60008086868660405160200161014c93929190610c0b565b60408051601f1981840301815291815281516020928301206000818152600390935291205490915060ff16156101b25760405162461bcd60e51b81526020600482015260066024820152657265706c617960d01b60448201526064015b60405180910390fd5b6000818152600360209081526040808320805460ff191660011790558051601f890183900483028101830190915287815261024b917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000918b908b908190840183828082843760009201919091525061054592505050565b9050600061028f8287878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061062992505050565b9050886001600160a01b0316816001600160a01b0316146102df5760405162461bcd60e51b815260206004820152600a602482015269217369676e617475726560b01b60448201526064016101a9565b6102ea60008a610645565b6102fb576102f960008a61066a565b505b6001600160a01b03891660009081526002602090815260408220805460018101825590835291200161032e888a83610cd6565b50886001600160a01b03167f78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e898960405161036a929190610d97565b60405180910390a250600198975050505050505050565b606060008267ffffffffffffffff81111561039e5761039e610c37565b6040519080825280602002602001820160405280156103d157816020015b60608152602001906001900390816103bc5790505b50905060005b8381101561052a57600260008686848181106103f5576103f5610dc6565b905060200201602081019061040a9190610ddc565b6001600160a01b03166001600160a01b03168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156104f657838290600052602060002001805461046990610c4d565b80601f016020809104026020016040519081016040528092919081815260200182805461049590610c4d565b80156104e25780601f106104b7576101008083540402835291602001916104e2565b820191906000526020600020905b8154815290600101906020018083116104c557829003601f168201915b50505050508152602001906001019061044a565b5050505082828151811061050c5761050c610dc6565b6020026020010181905250808061052290610df7565b9150506103d7565b5090505b92915050565b6060610540600061067f565b905090565b600080836001600160a01b03861660405160e09290921b6001600160e01b031916602083015260248201527512165411549310539157d0539393d55390d15351539560521b6044820152605a0160405160208183030381529060405280519060200120905061062081846040516020016105c0929190610e1e565b60408051601f1981840301815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000084830152603c8085019190915282518085039091018152605c909301909152815191012090565b95945050505050565b6000806000610638858561068c565b9150915061052a816106d1565b6001600160a01b038116600090815260018301602052604081205415155b9392505050565b6000610663836001600160a01b03841661081e565b606060006106638361086d565b60008082516041036106c25760208301516040840151606085015160001a6106b6878285856108c9565b945094505050506106ca565b506000905060025b9250929050565b60008160048111156106e5576106e5610e44565b036106ed5750565b600181600481111561070157610701610e44565b0361074e5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016101a9565b600281600481111561076257610762610e44565b036107af5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016101a9565b60038160048111156107c3576107c3610e44565b0361081b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016101a9565b50565b60008181526001830160205260408120546108655750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561052e565b50600061052e565b6060816000018054806020026020016040519081016040528092919081815260200182805480156108bd57602002820191906000526020600020905b8154815260200190600101908083116108a9575b50505050509050919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156109005750600090506003610984565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610954573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661097d57600060019250925050610984565b9150600090505b94509492505050565b80356001600160a01b03811681146109a457600080fd5b919050565b60008083601f8401126109bb57600080fd5b50813567ffffffffffffffff8111156109d357600080fd5b6020830191508360208285010111156106ca57600080fd5b600080600080600060608688031215610a0357600080fd5b610a0c8661098d565b9450602086013567ffffffffffffffff80821115610a2957600080fd5b610a3589838a016109a9565b90965094506040880135915080821115610a4e57600080fd5b50610a5b888289016109a9565b969995985093965092949392505050565b60008060208385031215610a7f57600080fd5b823567ffffffffffffffff80821115610a9757600080fd5b818501915085601f830112610aab57600080fd5b813581811115610aba57600080fd5b8660208260051b8501011115610acf57600080fd5b60209290920196919550909350505050565b60005b83811015610afc578181015183820152602001610ae4565b50506000910152565b60208152600060208201835180825260408401915060408160051b8501016020860160005b83811015610bbe57868303603f19018552815180518085526020918201918086019190600582901b87010160005b82811015610ba457601f198089840301855285518051808552610b82816020870160208501610ae1565b60209788019796870196601f9190910190921693909301019150600101610b58565b506020988901989096509490940193505050600101610b2a565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610bbe5783516001600160a01b031683529284019291840191600101610be6565b6bffffffffffffffffffffffff198460601b168152818360148301376000910160140190815292915050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680610c6157607f821691505b602082108103610c8157634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610cd157600081815260208120601f850160051c81016020861015610cae5750805b601f850160051c820191505b81811015610ccd57828155600101610cba565b5050505b505050565b67ffffffffffffffff831115610cee57610cee610c37565b610d0283610cfc8354610c4d565b83610c87565b6000601f841160018114610d365760008515610d1e5750838201355b600019600387901b1c1916600186901b178355610d90565b600083815260209020601f19861690835b82811015610d675786850135825560209485019460019092019101610d47565b5086821015610d845760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215610dee57600080fd5b6106638261098d565b600060018201610e1757634e487b7160e01b600052601160045260246000fd5b5060010190565b82815260008251610e36816020850160208701610ae1565b919091016020019392505050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212206c0f7b69e2d0a29e7e9a737e37991375a3825ea9bdaaee99a9f01bf28d17102a64736f6c63430008130033",
}

// AnnounceABI is the input ABI used to generate the binding from.
// Deprecated: Use AnnounceMetaData.ABI instead.
var AnnounceABI = AnnounceMetaData.ABI

// AnnounceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AnnounceMetaData.Bin instead.
var AnnounceBin = AnnounceMetaData.Bin

// DeployAnnounce deploys a new Ethereum contract, binding an instance of Announce to it.
func DeployAnnounce(auth *bind.TransactOpts, backend bind.ContractBackend, _mailbox common.Address) (common.Address, *types.Transaction, *Announce, error) {
	parsed, err := AnnounceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AnnounceBin), backend, _mailbox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Announce{AnnounceCaller: AnnounceCaller{contract: contract}, AnnounceTransactor: AnnounceTransactor{contract: contract}, AnnounceFilterer: AnnounceFilterer{contract: contract}}, nil
}

// Announce is an auto generated Go binding around an Ethereum contract.
type Announce struct {
	AnnounceCaller     // Read-only binding to the contract
	AnnounceTransactor // Write-only binding to the contract
	AnnounceFilterer   // Log filterer for contract events
}

// AnnounceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnnounceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnnounceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnnounceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnnounceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnnounceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnnounceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnnounceSession struct {
	Contract     *Announce         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnnounceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnnounceCallerSession struct {
	Contract *AnnounceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AnnounceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnnounceTransactorSession struct {
	Contract     *AnnounceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AnnounceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnnounceRaw struct {
	Contract *Announce // Generic contract binding to access the raw methods on
}

// AnnounceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnnounceCallerRaw struct {
	Contract *AnnounceCaller // Generic read-only contract binding to access the raw methods on
}

// AnnounceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnnounceTransactorRaw struct {
	Contract *AnnounceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnnounce creates a new instance of Announce, bound to a specific deployed contract.
func NewAnnounce(address common.Address, backend bind.ContractBackend) (*Announce, error) {
	contract, err := bindAnnounce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Announce{AnnounceCaller: AnnounceCaller{contract: contract}, AnnounceTransactor: AnnounceTransactor{contract: contract}, AnnounceFilterer: AnnounceFilterer{contract: contract}}, nil
}

// NewAnnounceCaller creates a new read-only instance of Announce, bound to a specific deployed contract.
func NewAnnounceCaller(address common.Address, caller bind.ContractCaller) (*AnnounceCaller, error) {
	contract, err := bindAnnounce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnnounceCaller{contract: contract}, nil
}

// NewAnnounceTransactor creates a new write-only instance of Announce, bound to a specific deployed contract.
func NewAnnounceTransactor(address common.Address, transactor bind.ContractTransactor) (*AnnounceTransactor, error) {
	contract, err := bindAnnounce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnnounceTransactor{contract: contract}, nil
}

// NewAnnounceFilterer creates a new log filterer instance of Announce, bound to a specific deployed contract.
func NewAnnounceFilterer(address common.Address, filterer bind.ContractFilterer) (*AnnounceFilterer, error) {
	contract, err := bindAnnounce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnnounceFilterer{contract: contract}, nil
}

// bindAnnounce binds a generic wrapper to an already deployed contract.
func bindAnnounce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnnounceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Announce *AnnounceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Announce.Contract.AnnounceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Announce *AnnounceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Announce.Contract.AnnounceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Announce *AnnounceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Announce.Contract.AnnounceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Announce *AnnounceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Announce.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Announce *AnnounceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Announce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Announce *AnnounceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Announce.Contract.contract.Transact(opts, method, params...)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_Announce *AnnounceCaller) GetAnnouncedStorageLocations(opts *bind.CallOpts, _validators []common.Address) ([][]string, error) {
	var out []interface{}
	err := _Announce.contract.Call(opts, &out, "getAnnouncedStorageLocations", _validators)

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_Announce *AnnounceSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _Announce.Contract.GetAnnouncedStorageLocations(&_Announce.CallOpts, _validators)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_Announce *AnnounceCallerSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _Announce.Contract.GetAnnouncedStorageLocations(&_Announce.CallOpts, _validators)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_Announce *AnnounceCaller) GetAnnouncedValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Announce.contract.Call(opts, &out, "getAnnouncedValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_Announce *AnnounceSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _Announce.Contract.GetAnnouncedValidators(&_Announce.CallOpts)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_Announce *AnnounceCallerSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _Announce.Contract.GetAnnouncedValidators(&_Announce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Announce *AnnounceCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Announce.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Announce *AnnounceSession) LocalDomain() (uint32, error) {
	return _Announce.Contract.LocalDomain(&_Announce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Announce *AnnounceCallerSession) LocalDomain() (uint32, error) {
	return _Announce.Contract.LocalDomain(&_Announce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Announce *AnnounceCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Announce.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Announce *AnnounceSession) Mailbox() (common.Address, error) {
	return _Announce.Contract.Mailbox(&_Announce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Announce *AnnounceCallerSession) Mailbox() (common.Address, error) {
	return _Announce.Contract.Mailbox(&_Announce.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_Announce *AnnounceTransactor) Announce(opts *bind.TransactOpts, _validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _Announce.contract.Transact(opts, "announce", _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_Announce *AnnounceSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _Announce.Contract.Announce(&_Announce.TransactOpts, _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_Announce *AnnounceTransactorSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _Announce.Contract.Announce(&_Announce.TransactOpts, _validator, _storageLocation, _signature)
}

// AnnounceValidatorAnnouncementIterator is returned from FilterValidatorAnnouncement and is used to iterate over the raw logs and unpacked data for ValidatorAnnouncement events raised by the Announce contract.
type AnnounceValidatorAnnouncementIterator struct {
	Event *AnnounceValidatorAnnouncement // Event containing the contract specifics and raw log

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
func (it *AnnounceValidatorAnnouncementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnnounceValidatorAnnouncement)
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
		it.Event = new(AnnounceValidatorAnnouncement)
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
func (it *AnnounceValidatorAnnouncementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnnounceValidatorAnnouncementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnnounceValidatorAnnouncement represents a ValidatorAnnouncement event raised by the Announce contract.
type AnnounceValidatorAnnouncement struct {
	Validator       common.Address
	StorageLocation string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorAnnouncement is a free log retrieval operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_Announce *AnnounceFilterer) FilterValidatorAnnouncement(opts *bind.FilterOpts, validator []common.Address) (*AnnounceValidatorAnnouncementIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Announce.contract.FilterLogs(opts, "ValidatorAnnouncement", validatorRule)
	if err != nil {
		return nil, err
	}
	return &AnnounceValidatorAnnouncementIterator{contract: _Announce.contract, event: "ValidatorAnnouncement", logs: logs, sub: sub}, nil
}

// WatchValidatorAnnouncement is a free log subscription operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_Announce *AnnounceFilterer) WatchValidatorAnnouncement(opts *bind.WatchOpts, sink chan<- *AnnounceValidatorAnnouncement, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Announce.contract.WatchLogs(opts, "ValidatorAnnouncement", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnnounceValidatorAnnouncement)
				if err := _Announce.contract.UnpackLog(event, "ValidatorAnnouncement", log); err != nil {
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

// ParseValidatorAnnouncement is a log parse operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_Announce *AnnounceFilterer) ParseValidatorAnnouncement(log types.Log) (*AnnounceValidatorAnnouncement, error) {
	event := new(AnnounceValidatorAnnouncement)
	if err := _Announce.contract.UnpackLog(event, "ValidatorAnnouncement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
