// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legacy_multisig

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

// LegacyMultisigMetaData contains all meta data concerning the LegacyMultisig contract.
var LegacyMultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorEnrolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorCount\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnenrolled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"commitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"enrollValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"address[][]\",\"name\":\"_validators\",\"type\":\"address[][]\"}],\"name\":\"enrollValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isEnrolled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moduleType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_thresholds\",\"type\":\"uint8[]\"}],\"name\":\"setThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"unenrollValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"validatorsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611bc78061007e6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063f211b73211610066578063f211b7321461023a578063f2fde38b1461025d578063f7e83aee14610270578063fe9f5ab11461028357600080fd5b8063715018a6146101e45780638da5cb5b146101ec5780639a6e9a2414610207578063e7d5d3a11461022757600080fd5b8063414c676f116100d3578063414c676f146101845780634422cc2a146101975780635d94f2b9146101aa5780636465e69f146101ca57600080fd5b8063020addf114610105578063174769dc1461012b5780632e0ed234146101405780634098991214610161575b600080fd5b61011861011336600461159d565b610296565b6040519081526020015b60405180910390f35b61013e6101393660046115fd565b6102ba565b005b61015361014e3660046116ab565b6103fc565b604051610122929190611731565b61017461016f36600461176d565b610440565b6040519015158152602001610122565b61013e6101923660046115fd565b61046e565b61013e6101a53660046117b1565b61051e565b6101186101b836600461159d565b60036020526000908152604090205481565b6101d2600381565b60405160ff9091168152602001610122565b61013e6105e0565b6000546040516001600160a01b039091168152602001610122565b61021a61021536600461159d565b6105f4565b60405161012291906117db565b61013e61023536600461176d565b6106b8565b6101d261024836600461159d565b60016020526000908152604090205460ff1681565b61013e61026b3660046117ee565b6107e9565b61017461027e366004611809565b610862565b61013e61029136600461176d565b6108f1565b63ffffffff811660009081526002602052604081206102b49061090c565b92915050565b6102c2610916565b828181146103015760405162461bcd60e51b8152602060048201526007602482015266042d8cadccee8d60cb1b60448201526064015b60405180910390fd5b60005b818110156103f45736600085858481811061032157610321611869565b9050602002810190610333919061187f565b90925090508060005b818110156103ac5761039a8a8a8781811061035957610359611869565b905060200201602081019061036e919061159d565b85858481811061038057610380611869565b905060200201602081019061039591906117ee565b610970565b6103a56001826118df565b905061033c565b506103dc8989868181106103c2576103c2611869565b90506020020160208101906103d7919061159d565b610a66565b505050506001816103ed91906118df565b9050610304565b505050505050565b606060008061040b8585610b0d565b90506000610418826105f4565b63ffffffff9092166000908152600160205260409020549193505060ff1690505b9250929050565b63ffffffff808316600090815260026020526040812090916104669082908590610b3016565b949350505050565b610476610916565b828181146104b05760405162461bcd60e51b8152602060048201526007602482015266042d8cadccee8d60cb1b60448201526064016102f8565b60005b818110156103f45761050c8686838181106104d0576104d0611869565b90506020020160208101906104e5919061159d565b8585848181106104f7576104f7611869565b90506020020160208101906101a591906118f2565b6105176001826118df565b90506104b3565b610526610916565b60008160ff16118015610544575061053d82610296565b8160ff1611155b6105795760405162461bcd60e51b81526020600482015260066024820152652172616e676560d01b60448201526064016102f8565b63ffffffff8216600081815260016020908152604091829020805460ff191660ff861690811790915591519182527ff25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a910160405180910390a26105db82610a66565b505050565b6105e8610916565b6105f26000610b55565b565b63ffffffff811660009081526002602052604081206060916106158261090c565b905060008167ffffffffffffffff81111561063257610632611923565b60405190808252806020026020018201604052801561065b578160200160208202803683370190505b50905060005b828110156106af576106738482610ba5565b82828151811061068557610685611869565b6001600160a01b0390921660209283029190910190910152806106a781611939565b915050610661565b50949350505050565b6106c0610916565b63ffffffff80831660009081526002602052604090206106e2918390610bb116565b61071a5760405162461bcd60e51b815260206004820152600960248201526808595b9c9bdb1b195960ba1b60448201526064016102f8565b600061072583610296565b63ffffffff841660009081526001602052604090205490915060ff168110156107905760405162461bcd60e51b815260206004820152601960248201527f76696f6c617465732071756f72756d207468726573686f6c640000000000000060448201526064016102f8565b61079983610a66565b50816001600160a01b03168363ffffffff167fa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c836040516107dc91815260200190565b60405180910390a3505050565b6107f1610916565b6001600160a01b0381166108565760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102f8565b61085f81610b55565b50565b600061087085858585610bc6565b6108a65760405162461bcd60e51b8152602060048201526007602482015266216d65726b6c6560c81b60448201526064016102f8565b6108b285858585610c40565b6108e65760405162461bcd60e51b8152602060048201526005602482015264217369677360d81b60448201526064016102f8565b506001949350505050565b6108f9610916565b6109038282610970565b6105db82610a66565b60006102b4825490565b6000546001600160a01b031633146105f25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102f8565b6001600160a01b0381166109b55760405162461bcd60e51b815260206004820152600c60248201526b7a65726f206164647265737360a01b60448201526064016102f8565b63ffffffff80831660009081526002602052604090206109d7918390610e2716565b610a165760405162461bcd60e51b815260206004820152601060248201526f185b1c9958591e48195b9c9bdb1b195960821b60448201526064016102f8565b806001600160a01b03168263ffffffff167f890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a610a5185610296565b60405190815260200160405180910390a35050565b600080610a72836105f4565b63ffffffff8416600090815260016020908152604080832054905193945060ff1692610aa2918491869101611952565b60408051808303601f19018152828252805160209182012063ffffffff8916600081815260038452849020829055845290830181905292507f6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3910160405180910390a1949350505050565b6000610b1d6009600584866119a1565b610b26916119cb565b60e01c9392505050565b6001600160a01b038116600090815260018301602052604081205415155b9392505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610b4e8383610e3c565b6000610b4e836001600160a01b038416610e66565b600080610c2a610c0b85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f5992505050565b610c158888610f64565b610c1f8787610f88565b63ffffffff16610f98565b9050610c368686611049565b1495945050505050565b600080610c4d8686611061565b9050600080610c5c8686610b0d565b9050600083610c6b8a8a611086565b604051602001610c7d939291906119fb565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260039093529120549091508114610ce85760405162461bcd60e51b815260206004820152600b60248201526a0858dbdb5b5a5d1b595b9d60aa1b60448201526064016102f8565b610d0f82610cf68b8b6110ab565b610d008c8c611049565b610d0a8d8d6110bb565b6110cb565b925050506000610d1f8888611167565b90506000805b8460ff16811015610e17576000610d7b85610d418d8d86611189565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506111cf92505050565b90505b8383108015610da95750610d938b8b856111eb565b6001600160a01b0316816001600160a01b031614155b15610dbe57610db783611939565b9250610d7e565b838310610dfa5760405162461bcd60e51b815260206004820152600a602482015269085d1a1c995cda1bdb1960b21b60448201526064016102f8565b610e0383611939565b92505080610e1090611939565b9050610d25565b5060019998505050505050505050565b6000610b4e836001600160a01b03841661124b565b6000826000018281548110610e5357610e53611869565b9060005260206000200154905092915050565b60008181526001830160205260408120548015610f4f576000610e8a600183611a23565b8554909150600090610e9e90600190611a23565b9050818114610f03576000866000018281548110610ebe57610ebe611869565b9060005260206000200154905080876000018481548110610ee157610ee1611869565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610f1457610f14611a36565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506102b4565b60009150506102b4565b805160209091012090565b610f6c611565565b610f7b610444604484866119a1565b810190610b4e9190611a4c565b6000610b1d6005600184866119a1565b8260005b602081101561104157600183821c166000858360208110610fbf57610fbf611869565b6020020151905081600103610fff57604080516020810183905290810185905260600160405160208183030381529060405280519060200120935061102c565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050808061103990611939565b915050610f9c565b509392505050565b600061105860208284866119a1565b610b4e91611ad9565b600061107361044561044484866119a1565b61107c91611af7565b60f81c9392505050565b3660008383611095868661129a565b6110a09282906119a1565b915091509250929050565b60006110586044602484866119a1565b6000610b1d6024602084866119a1565b6000806110d886866112c1565b60408051602080820184905281830188905260e087901b6001600160e01b0319166060830152825160448184030181526064830184528051908201207f19457468657265756d205369676e6564204d6573736167653a0a333200000000608484015260a0808401919091528351808403909101815260c090920190925280519101209091509695505050505050565b60006020611175848461129a565b61117f9084611a23565b610b4e9190611b25565b36600080611198604185611b47565b6111a4906104456118df565b905060006111b36041836118df565b90506111c18183888a6119a1565b935093505050935093915050565b60008060006111de8585611315565b9150915061104181611357565b6000806111f9836020611b47565b611203868661129a565b61120d91906118df565b61121890600c6118df565b905060006112278260146118df565b9050611235818387896119a1565b61123e91611b5e565b60601c9695505050505050565b6000818152600183016020526040812054611292575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556102b4565b5060006102b4565b600060416112a88484611061565b60ff166112b59190611b47565b610b4e906104456118df565b6040805160e084901b6001600160e01b031916602080830191909152602482018490526848595045524c414e4560b81b60448301528251602d818403018152604d9092019092528051910120600090610b4e565b600080825160410361134b5760208301516040840151606085015160001a61133f878285856114a1565b94509450505050610439565b50600090506002610439565b600081600481111561136b5761136b61190d565b036113735750565b60018160048111156113875761138761190d565b036113d45760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016102f8565b60028160048111156113e8576113e861190d565b036114355760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016102f8565b60038160048111156114495761144961190d565b0361085f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016102f8565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114d8575060009050600361155c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561152c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166115555760006001925092505061155c565b9150600090505b94509492505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff8116811461159857600080fd5b919050565b6000602082840312156115af57600080fd5b610b4e82611584565b60008083601f8401126115ca57600080fd5b50813567ffffffffffffffff8111156115e257600080fd5b6020830191508360208260051b850101111561043957600080fd5b6000806000806040858703121561161357600080fd5b843567ffffffffffffffff8082111561162b57600080fd5b611637888389016115b8565b9096509450602087013591508082111561165057600080fd5b5061165d878288016115b8565b95989497509550505050565b60008083601f84011261167b57600080fd5b50813567ffffffffffffffff81111561169357600080fd5b60208301915083602082850101111561043957600080fd5b600080602083850312156116be57600080fd5b823567ffffffffffffffff8111156116d557600080fd5b6116e185828601611669565b90969095509350505050565b600081518084526020808501945080840160005b838110156117265781516001600160a01b031687529582019590820190600101611701565b509495945050505050565b60408152600061174460408301856116ed565b905060ff831660208301529392505050565b80356001600160a01b038116811461159857600080fd5b6000806040838503121561178057600080fd5b61178983611584565b915061179760208401611756565b90509250929050565b803560ff8116811461159857600080fd5b600080604083850312156117c457600080fd5b6117cd83611584565b9150611797602084016117a0565b602081526000610b4e60208301846116ed565b60006020828403121561180057600080fd5b610b4e82611756565b6000806000806040858703121561181f57600080fd5b843567ffffffffffffffff8082111561183757600080fd5b61184388838901611669565b9096509450602087013591508082111561185c57600080fd5b5061165d87828801611669565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261189657600080fd5b83018035915067ffffffffffffffff8211156118b157600080fd5b6020019150600581901b360382131561043957600080fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156102b4576102b46118c9565b60006020828403121561190457600080fd5b610b4e826117a0565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b60006001820161194b5761194b6118c9565b5060010190565b60ff60f81b8360f81b168152600060018083018451602080870160005b838110156119935781516001600160a01b031685529382019390820190850161196f565b509298975050505050505050565b600080858511156119b157600080fd5b838611156119be57600080fd5b5050820193919092039150565b6001600160e01b031981358181169160048510156119f35780818660040360031b1b83161692505b505092915050565b60f884901b6001600160f81b0319168152818360018301376000910160010190815292915050565b818103818111156102b4576102b46118c9565b634e487b7160e01b600052603160045260246000fd5b6000610400808385031215611a6057600080fd5b83601f840112611a6f57600080fd5b60405181810181811067ffffffffffffffff82111715611a9f57634e487b7160e01b600052604160045260246000fd5b604052908301908085831115611ab457600080fd5b845b83811015611ace578035825260209182019101611ab6565b509095945050505050565b803560208310156102b457600019602084900360031b1b1692915050565b6001600160f81b031981358181169160018510156119f35760019490940360031b84901b1690921692915050565b600082611b4257634e487b7160e01b600052601260045260246000fd5b500490565b80820281158282048414176102b4576102b46118c9565b6bffffffffffffffffffffffff1981358181169160148510156119f35760149490940360031b84901b169092169291505056fea26469706673582212205be5f802b350ff52ba06752c41cfad8213d132cb8d30ed2d7b42f903a030457164736f6c63430008130033",
}

// LegacyMultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyMultisigMetaData.ABI instead.
var LegacyMultisigABI = LegacyMultisigMetaData.ABI

// LegacyMultisigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyMultisigMetaData.Bin instead.
var LegacyMultisigBin = LegacyMultisigMetaData.Bin

// DeployLegacyMultisig deploys a new Ethereum contract, binding an instance of LegacyMultisig to it.
func DeployLegacyMultisig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LegacyMultisig, error) {
	parsed, err := LegacyMultisigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyMultisigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyMultisig{LegacyMultisigCaller: LegacyMultisigCaller{contract: contract}, LegacyMultisigTransactor: LegacyMultisigTransactor{contract: contract}, LegacyMultisigFilterer: LegacyMultisigFilterer{contract: contract}}, nil
}

// LegacyMultisig is an auto generated Go binding around an Ethereum contract.
type LegacyMultisig struct {
	LegacyMultisigCaller     // Read-only binding to the contract
	LegacyMultisigTransactor // Write-only binding to the contract
	LegacyMultisigFilterer   // Log filterer for contract events
}

// LegacyMultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyMultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyMultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyMultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyMultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyMultisigSession struct {
	Contract     *LegacyMultisig   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyMultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyMultisigCallerSession struct {
	Contract *LegacyMultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LegacyMultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyMultisigTransactorSession struct {
	Contract     *LegacyMultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LegacyMultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyMultisigRaw struct {
	Contract *LegacyMultisig // Generic contract binding to access the raw methods on
}

// LegacyMultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyMultisigCallerRaw struct {
	Contract *LegacyMultisigCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyMultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyMultisigTransactorRaw struct {
	Contract *LegacyMultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyMultisig creates a new instance of LegacyMultisig, bound to a specific deployed contract.
func NewLegacyMultisig(address common.Address, backend bind.ContractBackend) (*LegacyMultisig, error) {
	contract, err := bindLegacyMultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisig{LegacyMultisigCaller: LegacyMultisigCaller{contract: contract}, LegacyMultisigTransactor: LegacyMultisigTransactor{contract: contract}, LegacyMultisigFilterer: LegacyMultisigFilterer{contract: contract}}, nil
}

// NewLegacyMultisigCaller creates a new read-only instance of LegacyMultisig, bound to a specific deployed contract.
func NewLegacyMultisigCaller(address common.Address, caller bind.ContractCaller) (*LegacyMultisigCaller, error) {
	contract, err := bindLegacyMultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigCaller{contract: contract}, nil
}

// NewLegacyMultisigTransactor creates a new write-only instance of LegacyMultisig, bound to a specific deployed contract.
func NewLegacyMultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyMultisigTransactor, error) {
	contract, err := bindLegacyMultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigTransactor{contract: contract}, nil
}

// NewLegacyMultisigFilterer creates a new log filterer instance of LegacyMultisig, bound to a specific deployed contract.
func NewLegacyMultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyMultisigFilterer, error) {
	contract, err := bindLegacyMultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigFilterer{contract: contract}, nil
}

// bindLegacyMultisig binds a generic wrapper to an already deployed contract.
func bindLegacyMultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyMultisigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMultisig *LegacyMultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMultisig.Contract.LegacyMultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMultisig *LegacyMultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.LegacyMultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMultisig *LegacyMultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.LegacyMultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyMultisig *LegacyMultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyMultisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyMultisig *LegacyMultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyMultisig *LegacyMultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.contract.Transact(opts, method, params...)
}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisig *LegacyMultisigCaller) Commitment(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "commitment", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisig *LegacyMultisigSession) Commitment(arg0 uint32) ([32]byte, error) {
	return _LegacyMultisig.Contract.Commitment(&_LegacyMultisig.CallOpts, arg0)
}

// Commitment is a free data retrieval call binding the contract method 0x5d94f2b9.
//
// Solidity: function commitment(uint32 ) view returns(bytes32)
func (_LegacyMultisig *LegacyMultisigCallerSession) Commitment(arg0 uint32) ([32]byte, error) {
	return _LegacyMultisig.Contract.Commitment(&_LegacyMultisig.CallOpts, arg0)
}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisig *LegacyMultisigCaller) IsEnrolled(opts *bind.CallOpts, _domain uint32, _address common.Address) (bool, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "isEnrolled", _domain, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisig *LegacyMultisigSession) IsEnrolled(_domain uint32, _address common.Address) (bool, error) {
	return _LegacyMultisig.Contract.IsEnrolled(&_LegacyMultisig.CallOpts, _domain, _address)
}

// IsEnrolled is a free data retrieval call binding the contract method 0x40989912.
//
// Solidity: function isEnrolled(uint32 _domain, address _address) view returns(bool)
func (_LegacyMultisig *LegacyMultisigCallerSession) IsEnrolled(_domain uint32, _address common.Address) (bool, error) {
	return _LegacyMultisig.Contract.IsEnrolled(&_LegacyMultisig.CallOpts, _domain, _address)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisig *LegacyMultisigCaller) ModuleType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "moduleType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisig *LegacyMultisigSession) ModuleType() (uint8, error) {
	return _LegacyMultisig.Contract.ModuleType(&_LegacyMultisig.CallOpts)
}

// ModuleType is a free data retrieval call binding the contract method 0x6465e69f.
//
// Solidity: function moduleType() view returns(uint8)
func (_LegacyMultisig *LegacyMultisigCallerSession) ModuleType() (uint8, error) {
	return _LegacyMultisig.Contract.ModuleType(&_LegacyMultisig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisig *LegacyMultisigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisig *LegacyMultisigSession) Owner() (common.Address, error) {
	return _LegacyMultisig.Contract.Owner(&_LegacyMultisig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LegacyMultisig *LegacyMultisigCallerSession) Owner() (common.Address, error) {
	return _LegacyMultisig.Contract.Owner(&_LegacyMultisig.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisig *LegacyMultisigCaller) Threshold(opts *bind.CallOpts, arg0 uint32) (uint8, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisig *LegacyMultisigSession) Threshold(arg0 uint32) (uint8, error) {
	return _LegacyMultisig.Contract.Threshold(&_LegacyMultisig.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xf211b732.
//
// Solidity: function threshold(uint32 ) view returns(uint8)
func (_LegacyMultisig *LegacyMultisigCallerSession) Threshold(arg0 uint32) (uint8, error) {
	return _LegacyMultisig.Contract.Threshold(&_LegacyMultisig.CallOpts, arg0)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisig *LegacyMultisigCaller) ValidatorCount(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "validatorCount", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisig *LegacyMultisigSession) ValidatorCount(_domain uint32) (*big.Int, error) {
	return _LegacyMultisig.Contract.ValidatorCount(&_LegacyMultisig.CallOpts, _domain)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x020addf1.
//
// Solidity: function validatorCount(uint32 _domain) view returns(uint256)
func (_LegacyMultisig *LegacyMultisigCallerSession) ValidatorCount(_domain uint32) (*big.Int, error) {
	return _LegacyMultisig.Contract.ValidatorCount(&_LegacyMultisig.CallOpts, _domain)
}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisig *LegacyMultisigCaller) Validators(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "validators", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisig *LegacyMultisigSession) Validators(_domain uint32) ([]common.Address, error) {
	return _LegacyMultisig.Contract.Validators(&_LegacyMultisig.CallOpts, _domain)
}

// Validators is a free data retrieval call binding the contract method 0x9a6e9a24.
//
// Solidity: function validators(uint32 _domain) view returns(address[])
func (_LegacyMultisig *LegacyMultisigCallerSession) Validators(_domain uint32) ([]common.Address, error) {
	return _LegacyMultisig.Contract.Validators(&_LegacyMultisig.CallOpts, _domain)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisig *LegacyMultisigCaller) ValidatorsAndThreshold(opts *bind.CallOpts, _message []byte) ([]common.Address, uint8, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "validatorsAndThreshold", _message)

	if err != nil {
		return *new([]common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisig *LegacyMultisigSession) ValidatorsAndThreshold(_message []byte) ([]common.Address, uint8, error) {
	return _LegacyMultisig.Contract.ValidatorsAndThreshold(&_LegacyMultisig.CallOpts, _message)
}

// ValidatorsAndThreshold is a free data retrieval call binding the contract method 0x2e0ed234.
//
// Solidity: function validatorsAndThreshold(bytes _message) view returns(address[], uint8)
func (_LegacyMultisig *LegacyMultisigCallerSession) ValidatorsAndThreshold(_message []byte) ([]common.Address, uint8, error) {
	return _LegacyMultisig.Contract.ValidatorsAndThreshold(&_LegacyMultisig.CallOpts, _message)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisig *LegacyMultisigCaller) Verify(opts *bind.CallOpts, _metadata []byte, _message []byte) (bool, error) {
	var out []interface{}
	err := _LegacyMultisig.contract.Call(opts, &out, "verify", _metadata, _message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisig *LegacyMultisigSession) Verify(_metadata []byte, _message []byte) (bool, error) {
	return _LegacyMultisig.Contract.Verify(&_LegacyMultisig.CallOpts, _metadata, _message)
}

// Verify is a free data retrieval call binding the contract method 0xf7e83aee.
//
// Solidity: function verify(bytes _metadata, bytes _message) view returns(bool)
func (_LegacyMultisig *LegacyMultisigCallerSession) Verify(_metadata []byte, _message []byte) (bool, error) {
	return _LegacyMultisig.Contract.Verify(&_LegacyMultisig.CallOpts, _metadata, _message)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) EnrollValidator(opts *bind.TransactOpts, _domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "enrollValidator", _domain, _validator)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigSession) EnrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.EnrollValidator(&_LegacyMultisig.TransactOpts, _domain, _validator)
}

// EnrollValidator is a paid mutator transaction binding the contract method 0xfe9f5ab1.
//
// Solidity: function enrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) EnrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.EnrollValidator(&_LegacyMultisig.TransactOpts, _domain, _validator)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) EnrollValidators(opts *bind.TransactOpts, _domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "enrollValidators", _domains, _validators)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisig *LegacyMultisigSession) EnrollValidators(_domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.EnrollValidators(&_LegacyMultisig.TransactOpts, _domains, _validators)
}

// EnrollValidators is a paid mutator transaction binding the contract method 0x174769dc.
//
// Solidity: function enrollValidators(uint32[] _domains, address[][] _validators) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) EnrollValidators(_domains []uint32, _validators [][]common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.EnrollValidators(&_LegacyMultisig.TransactOpts, _domains, _validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisig *LegacyMultisigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisig *LegacyMultisigSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyMultisig.Contract.RenounceOwnership(&_LegacyMultisig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LegacyMultisig.Contract.RenounceOwnership(&_LegacyMultisig.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) SetThreshold(opts *bind.TransactOpts, _domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "setThreshold", _domain, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisig *LegacyMultisigSession) SetThreshold(_domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.SetThreshold(&_LegacyMultisig.TransactOpts, _domain, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x4422cc2a.
//
// Solidity: function setThreshold(uint32 _domain, uint8 _threshold) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) SetThreshold(_domain uint32, _threshold uint8) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.SetThreshold(&_LegacyMultisig.TransactOpts, _domain, _threshold)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) SetThresholds(opts *bind.TransactOpts, _domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "setThresholds", _domains, _thresholds)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisig *LegacyMultisigSession) SetThresholds(_domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.SetThresholds(&_LegacyMultisig.TransactOpts, _domains, _thresholds)
}

// SetThresholds is a paid mutator transaction binding the contract method 0x414c676f.
//
// Solidity: function setThresholds(uint32[] _domains, uint8[] _thresholds) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) SetThresholds(_domains []uint32, _thresholds []uint8) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.SetThresholds(&_LegacyMultisig.TransactOpts, _domains, _thresholds)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisig *LegacyMultisigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.TransferOwnership(&_LegacyMultisig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.TransferOwnership(&_LegacyMultisig.TransactOpts, newOwner)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigTransactor) UnenrollValidator(opts *bind.TransactOpts, _domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.contract.Transact(opts, "unenrollValidator", _domain, _validator)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigSession) UnenrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.UnenrollValidator(&_LegacyMultisig.TransactOpts, _domain, _validator)
}

// UnenrollValidator is a paid mutator transaction binding the contract method 0xe7d5d3a1.
//
// Solidity: function unenrollValidator(uint32 _domain, address _validator) returns()
func (_LegacyMultisig *LegacyMultisigTransactorSession) UnenrollValidator(_domain uint32, _validator common.Address) (*types.Transaction, error) {
	return _LegacyMultisig.Contract.UnenrollValidator(&_LegacyMultisig.TransactOpts, _domain, _validator)
}

// LegacyMultisigCommitmentUpdatedIterator is returned from FilterCommitmentUpdated and is used to iterate over the raw logs and unpacked data for CommitmentUpdated events raised by the LegacyMultisig contract.
type LegacyMultisigCommitmentUpdatedIterator struct {
	Event *LegacyMultisigCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigCommitmentUpdated)
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
		it.Event = new(LegacyMultisigCommitmentUpdated)
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
func (it *LegacyMultisigCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigCommitmentUpdated represents a CommitmentUpdated event raised by the LegacyMultisig contract.
type LegacyMultisigCommitmentUpdated struct {
	Domain     uint32
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitmentUpdated is a free log retrieval operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisig *LegacyMultisigFilterer) FilterCommitmentUpdated(opts *bind.FilterOpts) (*LegacyMultisigCommitmentUpdatedIterator, error) {

	logs, sub, err := _LegacyMultisig.contract.FilterLogs(opts, "CommitmentUpdated")
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigCommitmentUpdatedIterator{contract: _LegacyMultisig.contract, event: "CommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitmentUpdated is a free log subscription operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisig *LegacyMultisigFilterer) WatchCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *LegacyMultisigCommitmentUpdated) (event.Subscription, error) {

	logs, sub, err := _LegacyMultisig.contract.WatchLogs(opts, "CommitmentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigCommitmentUpdated)
				if err := _LegacyMultisig.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
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

// ParseCommitmentUpdated is a log parse operation binding the contract event 0x6f17fabbcaf1b0e8cfbc153d8884a27de46c3076c2de8ac2aef0094a134909e3.
//
// Solidity: event CommitmentUpdated(uint32 domain, bytes32 commitment)
func (_LegacyMultisig *LegacyMultisigFilterer) ParseCommitmentUpdated(log types.Log) (*LegacyMultisigCommitmentUpdated, error) {
	event := new(LegacyMultisigCommitmentUpdated)
	if err := _LegacyMultisig.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LegacyMultisig contract.
type LegacyMultisigOwnershipTransferredIterator struct {
	Event *LegacyMultisigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigOwnershipTransferred)
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
		it.Event = new(LegacyMultisigOwnershipTransferred)
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
func (it *LegacyMultisigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigOwnershipTransferred represents a OwnershipTransferred event raised by the LegacyMultisig contract.
type LegacyMultisigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyMultisig *LegacyMultisigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LegacyMultisigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyMultisig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigOwnershipTransferredIterator{contract: _LegacyMultisig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LegacyMultisig *LegacyMultisigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LegacyMultisigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LegacyMultisig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigOwnershipTransferred)
				if err := _LegacyMultisig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LegacyMultisig *LegacyMultisigFilterer) ParseOwnershipTransferred(log types.Log) (*LegacyMultisigOwnershipTransferred, error) {
	event := new(LegacyMultisigOwnershipTransferred)
	if err := _LegacyMultisig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigThresholdSetIterator is returned from FilterThresholdSet and is used to iterate over the raw logs and unpacked data for ThresholdSet events raised by the LegacyMultisig contract.
type LegacyMultisigThresholdSetIterator struct {
	Event *LegacyMultisigThresholdSet // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigThresholdSet)
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
		it.Event = new(LegacyMultisigThresholdSet)
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
func (it *LegacyMultisigThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigThresholdSet represents a ThresholdSet event raised by the LegacyMultisig contract.
type LegacyMultisigThresholdSet struct {
	Domain    uint32
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdSet is a free log retrieval operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisig *LegacyMultisigFilterer) FilterThresholdSet(opts *bind.FilterOpts, domain []uint32) (*LegacyMultisigThresholdSetIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _LegacyMultisig.contract.FilterLogs(opts, "ThresholdSet", domainRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigThresholdSetIterator{contract: _LegacyMultisig.contract, event: "ThresholdSet", logs: logs, sub: sub}, nil
}

// WatchThresholdSet is a free log subscription operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisig *LegacyMultisigFilterer) WatchThresholdSet(opts *bind.WatchOpts, sink chan<- *LegacyMultisigThresholdSet, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _LegacyMultisig.contract.WatchLogs(opts, "ThresholdSet", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigThresholdSet)
				if err := _LegacyMultisig.contract.UnpackLog(event, "ThresholdSet", log); err != nil {
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

// ParseThresholdSet is a log parse operation binding the contract event 0xf25cfff98c95cf069df801752174d854732576e4b283bc4299386f65676e386a.
//
// Solidity: event ThresholdSet(uint32 indexed domain, uint8 threshold)
func (_LegacyMultisig *LegacyMultisigFilterer) ParseThresholdSet(log types.Log) (*LegacyMultisigThresholdSet, error) {
	event := new(LegacyMultisigThresholdSet)
	if err := _LegacyMultisig.contract.UnpackLog(event, "ThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigValidatorEnrolledIterator is returned from FilterValidatorEnrolled and is used to iterate over the raw logs and unpacked data for ValidatorEnrolled events raised by the LegacyMultisig contract.
type LegacyMultisigValidatorEnrolledIterator struct {
	Event *LegacyMultisigValidatorEnrolled // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigValidatorEnrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigValidatorEnrolled)
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
		it.Event = new(LegacyMultisigValidatorEnrolled)
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
func (it *LegacyMultisigValidatorEnrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigValidatorEnrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigValidatorEnrolled represents a ValidatorEnrolled event raised by the LegacyMultisig contract.
type LegacyMultisigValidatorEnrolled struct {
	Domain         uint32
	Validator      common.Address
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorEnrolled is a free log retrieval operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) FilterValidatorEnrolled(opts *bind.FilterOpts, domain []uint32, validator []common.Address) (*LegacyMultisigValidatorEnrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisig.contract.FilterLogs(opts, "ValidatorEnrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigValidatorEnrolledIterator{contract: _LegacyMultisig.contract, event: "ValidatorEnrolled", logs: logs, sub: sub}, nil
}

// WatchValidatorEnrolled is a free log subscription operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) WatchValidatorEnrolled(opts *bind.WatchOpts, sink chan<- *LegacyMultisigValidatorEnrolled, domain []uint32, validator []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisig.contract.WatchLogs(opts, "ValidatorEnrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigValidatorEnrolled)
				if err := _LegacyMultisig.contract.UnpackLog(event, "ValidatorEnrolled", log); err != nil {
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

// ParseValidatorEnrolled is a log parse operation binding the contract event 0x890997ec79a2b993cbc6e69433ed0ffa6303de16e3c21e315ba91c21ab5ec15a.
//
// Solidity: event ValidatorEnrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) ParseValidatorEnrolled(log types.Log) (*LegacyMultisigValidatorEnrolled, error) {
	event := new(LegacyMultisigValidatorEnrolled)
	if err := _LegacyMultisig.contract.UnpackLog(event, "ValidatorEnrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyMultisigValidatorUnenrolledIterator is returned from FilterValidatorUnenrolled and is used to iterate over the raw logs and unpacked data for ValidatorUnenrolled events raised by the LegacyMultisig contract.
type LegacyMultisigValidatorUnenrolledIterator struct {
	Event *LegacyMultisigValidatorUnenrolled // Event containing the contract specifics and raw log

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
func (it *LegacyMultisigValidatorUnenrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyMultisigValidatorUnenrolled)
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
		it.Event = new(LegacyMultisigValidatorUnenrolled)
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
func (it *LegacyMultisigValidatorUnenrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyMultisigValidatorUnenrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyMultisigValidatorUnenrolled represents a ValidatorUnenrolled event raised by the LegacyMultisig contract.
type LegacyMultisigValidatorUnenrolled struct {
	Domain         uint32
	Validator      common.Address
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnenrolled is a free log retrieval operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) FilterValidatorUnenrolled(opts *bind.FilterOpts, domain []uint32, validator []common.Address) (*LegacyMultisigValidatorUnenrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisig.contract.FilterLogs(opts, "ValidatorUnenrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LegacyMultisigValidatorUnenrolledIterator{contract: _LegacyMultisig.contract, event: "ValidatorUnenrolled", logs: logs, sub: sub}, nil
}

// WatchValidatorUnenrolled is a free log subscription operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) WatchValidatorUnenrolled(opts *bind.WatchOpts, sink chan<- *LegacyMultisigValidatorUnenrolled, domain []uint32, validator []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LegacyMultisig.contract.WatchLogs(opts, "ValidatorUnenrolled", domainRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyMultisigValidatorUnenrolled)
				if err := _LegacyMultisig.contract.UnpackLog(event, "ValidatorUnenrolled", log); err != nil {
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

// ParseValidatorUnenrolled is a log parse operation binding the contract event 0xa4c7a7b783c9afd72ed0b93a7e67ca063acdaa9c3b3268bd43abe1199bdad27c.
//
// Solidity: event ValidatorUnenrolled(uint32 indexed domain, address indexed validator, uint256 validatorCount)
func (_LegacyMultisig *LegacyMultisigFilterer) ParseValidatorUnenrolled(log types.Log) (*LegacyMultisigValidatorUnenrolled, error) {
	event := new(LegacyMultisigValidatorUnenrolled)
	if err := _LegacyMultisig.contract.UnpackLog(event, "ValidatorUnenrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
