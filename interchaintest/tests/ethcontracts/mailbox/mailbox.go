// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mailbox

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

// MailboxMetaData contains all meta data concerning the Mailbox contract.
var MailboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DefaultIsmSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"DispatchId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"ProcessId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delivered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultIsm\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"recipientIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setDefaultIsm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611b5c380380611b5c83398101604081905261002f9161003d565b63ffffffff1660805261006a565b60006020828403121561004f57600080fd5b815163ffffffff8116811461006357600080fd5b9392505050565b608051611ac9610093600039600081816101d6015281816104ff0152610a0e0152611ac96000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063907c0f92116100ad578063f2fde38b11610071578063f2fde38b14610281578063f794687a14610294578063fa31de01146102a7578063fd54b228146102ba578063ffa1ad74146102c457600080fd5b8063907c0f9214610209578063b187bd261461022b578063e495f1d414610243578063e70f48ac14610266578063ebf0c7171461027957600080fd5b8063715018a6116100f4578063715018a6146101ae5780637c39d130146101b65780638456cb59146101c95780638d3638f4146101d15780638da5cb5b146101f857600080fd5b806306661abd146101315780633f4ba83a1461014f578063485cc95514610159578063522ae0021461016c5780636e5f516e14610183575b600080fd5b60b8545b60405163ffffffff90911681526020015b60405180910390f35b6101576102de565b005b6101576101673660046114fe565b610319565b61017561080081565b604051908152602001610146565b609754610196906001600160a01b031681565b6040516001600160a01b039091168152602001610146565b61015761044b565b6101576101c4366004611579565b61045f565b6101576107e2565b6101357f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610196565b61021161081d565b6040805192835263ffffffff909116602083015201610146565b610233610845565b6040519015158152602001610146565b6102336102513660046115e5565b60b96020526000908152604090205460ff1681565b6101966102743660046115fe565b610858565b6101756108e1565b61015761028f3660046115fe565b6108ed565b6101576102a23660046115fe565b610966565b6101756102b536600461161b565b610977565b60b8546101759081565b6102cc600081565b60405160ff9091168152602001610146565b6102e6610aca565b6102ee610b24565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b600054610100900460ff16158080156103395750600054600160ff909116105b806103535750303b158015610353575060005460ff166001145b6103bb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156103de576000805461ff0019166101001790555b6103e6610b67565b6103ee610b8e565b6103f7836108ed565b61040082610bbd565b8015610446576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610453610aca565b61045d6000610c4a565b565b6001606554146104b15760405162461bcd60e51b815260206004820152601a60248201527f7265656e7472616e742063616c6c20286f72207061757365642900000000000060448201526064016103b2565b600060658190556104c28383610c9c565b60ff16146104fd5760405162461bcd60e51b815260206004820152600860248201526710bb32b939b4b7b760c11b60448201526064016103b2565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff1661052e8383610cc0565b63ffffffff16146105705760405162461bcd60e51b815260206004820152600c60248201526b10b232b9ba34b730ba34b7b760a11b60448201526064016103b2565b60006105b183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ce392505050565b600081815260b9602052604090205490915060ff16156105ff5760405162461bcd60e51b815260206004820152600960248201526819195b1a5d995c995960ba1b60448201526064016103b2565b600081815260b960205260408120805460ff191660011790556106256102748585610cee565b604051637bf41d7760e11b81529091506001600160a01b0382169063f7e83aee9061065a90899089908990899060040161169d565b6020604051808303816000875af1158015610679573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069d91906116cf565b6106d35760405162461bcd60e51b8152602060048201526007602482015266216d6f64756c6560c81b60448201526064016103b2565b60006106df8585610d07565b905060006106ed8686610d17565b905060006106fb8787610cee565b9050806001600160a01b03166356d5d47584846107188b8b610d30565b6040518563ffffffff1660e01b815260040161073794939291906116f1565b600060405180830381600087803b15801561075157600080fd5b505af1158015610765573d6000803e3d6000fd5b50505050806001600160a01b0316828463ffffffff167f0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca657460405160405180910390a460405185907f1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a990600090a25050600160655550505050505050565b6107ea610aca565b6107f2610d4c565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a1565b6000806108286108e1565b600161083360b85490565b61083d9190611737565b915091509091565b600061085360655460021490565b905090565b6000816001600160a01b031663de523cf36040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156108b4575060408051601f3d908101601f191682019092526108b19181019061175b565b60015b156108d0576001600160a01b038116156108ce5792915050565b505b50506097546001600160a01b031690565b60006108536098610d8e565b6108f5610aca565b6001600160a01b03811661095a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103b2565b61096381610c4a565b50565b61096e610aca565b61096381610bbd565b600061098560655460021490565b156109bb5760405162461bcd60e51b81526020600482015260066024820152651c185d5cd95960d21b60448201526064016103b2565b6108008211156109fc5760405162461bcd60e51b815260206004820152600c60248201526b6d736720746f6f206c6f6e6760a01b60448201526064016103b2565b6000610a376000610a0c60b85490565b7f0000000000000000000000000000000000000000000000000000000000000000338a8a8a8a610da1565b80516020820120909150610a4c609882610ddf565b858763ffffffff16336001600160a01b03167f769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c81485604051610a8d9190611778565b60405180910390a460405181907f788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a90600090a29695505050505050565b6033546001600160a01b0316331461045d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103b2565b606554600214610b605760405162461bcd60e51b8152602060048201526007602482015266085c185d5cd95960ca1b60448201526064016103b2565b6001606555565b600054610100900460ff16610b605760405162461bcd60e51b81526004016103b2906117c6565b600054610100900460ff16610bb55760405162461bcd60e51b81526004016103b2906117c6565b61045d610ef7565b6001600160a01b0381163b610c005760405162461bcd60e51b81526020600482015260096024820152680858dbdb9d1c9858dd60ba1b60448201526064016103b2565b609780546001600160a01b0319166001600160a01b0383169081179091556040517fa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa90600090a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610cab6001828486611811565b610cb49161183b565b60f81c90505b92915050565b6000610cd0602d60298486611811565b610cd991611869565b60e01c9392505050565b805160209091012090565b6000610d00610cfd8484610f27565b90565b9392505050565b6000610cd0600960058486611811565b6000610d27602960098486611811565b610d0091611897565b366000610d4083604d8187611811565b915091505b9250929050565b606554600203610d875760405162461bcd60e51b81526020600482015260066024820152651c185d5cd95960d21b60448201526064016103b2565b6002606555565b6000610cba82610d9c610f37565b6113f8565b60608888888888888888604051602001610dc29897969594939291906118b5565b604051602081830303815290604052905098975050505050505050565b6001610ded602060026119fa565b610df79190611a06565b826020015410610e3c5760405162461bcd60e51b815260206004820152601060248201526f1b595c9adb19481d1c995948199d5b1b60821b60448201526064016103b2565b6001826020016000828254610e519190611a19565b9091555050602082015460005b6020811015610eee5781600116600103610e8d5782848260208110610e8557610e85611a2c565b015550505050565b838160208110610e9f57610e9f611a2c565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250600282610eda9190611a42565b915080610ee681611a64565b915050610e5e565b50610446611a7d565b600054610100900460ff16610f1e5760405162461bcd60e51b81526004016103b2906117c6565b61045d33610c4a565b6000610d27604d602d8486611811565b610f3f6114ca565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b60208110156114c257600182821c16600086836020811061142557611425611a2c565b01549050816001036114625760408051602081018390529081018690526060016040516020818303038152906040528051906020012094506114ad565b8486846020811061147557611475611a2c565b6020020151604051602001611494929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b505080806114ba90611a64565b915050611402565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b6001600160a01b038116811461096357600080fd5b6000806040838503121561151157600080fd5b823561151c816114e9565b9150602083013561152c816114e9565b809150509250929050565b60008083601f84011261154957600080fd5b50813567ffffffffffffffff81111561156157600080fd5b602083019150836020828501011115610d4557600080fd5b6000806000806040858703121561158f57600080fd5b843567ffffffffffffffff808211156115a757600080fd5b6115b388838901611537565b909650945060208701359150808211156115cc57600080fd5b506115d987828801611537565b95989497509550505050565b6000602082840312156115f757600080fd5b5035919050565b60006020828403121561161057600080fd5b8135610d00816114e9565b6000806000806060858703121561163157600080fd5b843563ffffffff8116811461164557600080fd5b935060208501359250604085013567ffffffffffffffff81111561166857600080fd5b6115d987828801611537565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815260006116b1604083018688611674565b82810360208401526116c4818587611674565b979650505050505050565b6000602082840312156116e157600080fd5b81518015158114610d0057600080fd5b63ffffffff85168152836020820152606060408201526000611717606083018486611674565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b63ffffffff82811682821603908082111561175457611754611721565b5092915050565b60006020828403121561176d57600080fd5b8151610d00816114e9565b600060208083528351808285015260005b818110156117a557858101830151858201604001528201611789565b506000604082860101526040601f19601f8301168501019250505092915050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6000808585111561182157600080fd5b8386111561182e57600080fd5b5050820193919092039150565b6001600160f81b031981358181169160018510156114c25760019490940360031b84901b1690921692915050565b6001600160e01b031981358181169160048510156114c25760049490940360031b84901b1690921692915050565b80356020831015610cba57600019602084900360031b1b1692915050565b60ff60f81b8960f81b168152600063ffffffff60e01b808a60e01b166001840152808960e01b166005840152876009840152808760e01b1660298401525084602d8301528284604d8401375060009101604d01908152979650505050505050565b600181815b8085111561195157816000190482111561193757611937611721565b8085161561194457918102915b93841c939080029061191b565b509250929050565b60008261196857506001610cba565b8161197557506000610cba565b816001811461198b5760028114611995576119b1565b6001915050610cba565b60ff8411156119a6576119a6611721565b50506001821b610cba565b5060208310610133831016604e8410600b84101617156119d4575081810a610cba565b6119de8383611916565b80600019048211156119f2576119f2611721565b029392505050565b6000610d008383611959565b81810381811115610cba57610cba611721565b80820180821115610cba57610cba611721565b634e487b7160e01b600052603260045260246000fd5b600082611a5f57634e487b7160e01b600052601260045260246000fd5b500490565b600060018201611a7657611a76611721565b5060010190565b634e487b7160e01b600052600160045260246000fdfea26469706673582212206ce32aa78bef04073d4d4988c158a459a1bb808ade2504b687acbc0d7e865dcd64736f6c63430008130033",
}

// MailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MailboxMetaData.ABI instead.
var MailboxABI = MailboxMetaData.ABI

// MailboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MailboxMetaData.Bin instead.
var MailboxBin = MailboxMetaData.Bin

// DeployMailbox deploys a new Ethereum contract, binding an instance of Mailbox to it.
func DeployMailbox(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *Mailbox, error) {
	parsed, err := MailboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MailboxBin), backend, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MailboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxCaller) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "MAX_MESSAGE_BODY_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Mailbox.Contract.MAXMESSAGEBODYBYTES(&_Mailbox.CallOpts)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxCallerSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Mailbox.Contract.MAXMESSAGEBODYBYTES(&_Mailbox.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCallerSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxCaller) Count(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxSession) Count() (uint32, error) {
	return _Mailbox.Contract.Count(&_Mailbox.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxCallerSession) Count() (uint32, error) {
	return _Mailbox.Contract.Count(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCaller) DefaultIsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "defaultIsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCallerSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxCaller) Delivered(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "delivered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxSession) Delivered(arg0 [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, arg0)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxCallerSession) Delivered(arg0 [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxSession) IsPaused() (bool, error) {
	return _Mailbox.Contract.IsPaused(&_Mailbox.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxCallerSession) IsPaused() (bool, error) {
	return _Mailbox.Contract.IsPaused(&_Mailbox.CallOpts)
}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxCaller) LatestCheckpoint(opts *bind.CallOpts) ([32]byte, uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "latestCheckpoint")

	if err != nil {
		return *new([32]byte), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxSession) LatestCheckpoint() ([32]byte, uint32, error) {
	return _Mailbox.Contract.LatestCheckpoint(&_Mailbox.CallOpts)
}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxCallerSession) LatestCheckpoint() ([32]byte, uint32, error) {
	return _Mailbox.Contract.LatestCheckpoint(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCallerSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCallerSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCaller) RecipientIsm(opts *bind.CallOpts, _recipient common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "recipientIsm", _recipient)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCallerSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxSession) Root() ([32]byte, error) {
	return _Mailbox.Contract.Root(&_Mailbox.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxCallerSession) Root() ([32]byte, error) {
	return _Mailbox.Contract.Root(&_Mailbox.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxSession) Tree() (*big.Int, error) {
	return _Mailbox.Contract.Tree(&_Mailbox.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxCallerSession) Tree() (*big.Int, error) {
	return _Mailbox.Contract.Tree(&_Mailbox.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "initialize", _owner, _defaultIsm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxSession) Initialize(_owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxTransactorSession) Initialize(_owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxSession) Pause() (*types.Transaction, error) {
	return _Mailbox.Contract.Pause(&_Mailbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Mailbox.Contract.Pause(&_Mailbox.TransactOpts)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxTransactor) Process(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "process", _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxTransactorSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactor) SetDefaultIsm(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setDefaultIsm", _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactorSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxSession) Unpause() (*types.Transaction, error) {
	return _Mailbox.Contract.Unpause(&_Mailbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Mailbox.Contract.Unpause(&_Mailbox.TransactOpts)
}

// MailboxDefaultIsmSetIterator is returned from FilterDefaultIsmSet and is used to iterate over the raw logs and unpacked data for DefaultIsmSet events raised by the Mailbox contract.
type MailboxDefaultIsmSetIterator struct {
	Event *MailboxDefaultIsmSet // Event containing the contract specifics and raw log

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
func (it *MailboxDefaultIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDefaultIsmSet)
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
		it.Event = new(MailboxDefaultIsmSet)
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
func (it *MailboxDefaultIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDefaultIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDefaultIsmSet represents a DefaultIsmSet event raised by the Mailbox contract.
type MailboxDefaultIsmSet struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultIsmSet is a free log retrieval operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) FilterDefaultIsmSet(opts *bind.FilterOpts, module []common.Address) (*MailboxDefaultIsmSetIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDefaultIsmSetIterator{contract: _Mailbox.contract, event: "DefaultIsmSet", logs: logs, sub: sub}, nil
}

// WatchDefaultIsmSet is a free log subscription operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) WatchDefaultIsmSet(opts *bind.WatchOpts, sink chan<- *MailboxDefaultIsmSet, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDefaultIsmSet)
				if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
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

// ParseDefaultIsmSet is a log parse operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) ParseDefaultIsmSet(log types.Log) (*MailboxDefaultIsmSet, error) {
	event := new(MailboxDefaultIsmSet)
	if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Mailbox contract.
type MailboxDispatchIterator struct {
	Event *MailboxDispatch // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatch)
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
		it.Event = new(MailboxDispatch)
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
func (it *MailboxDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatch represents a Dispatch event raised by the Mailbox contract.
type MailboxDispatch struct {
	Sender      common.Address
	Destination uint32
	Recipient   [32]byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) FilterDispatch(opts *bind.FilterOpts, sender []common.Address, destination []uint32, recipient [][32]byte) (*MailboxDispatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIterator{contract: _Mailbox.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *MailboxDispatch, sender []common.Address, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatch)
				if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) ParseDispatch(log types.Log) (*MailboxDispatch, error) {
	event := new(MailboxDispatch)
	if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIdIterator is returned from FilterDispatchId and is used to iterate over the raw logs and unpacked data for DispatchId events raised by the Mailbox contract.
type MailboxDispatchIdIterator struct {
	Event *MailboxDispatchId // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatchId)
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
		it.Event = new(MailboxDispatchId)
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
func (it *MailboxDispatchIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatchId represents a DispatchId event raised by the Mailbox contract.
type MailboxDispatchId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDispatchId is a free log retrieval operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterDispatchId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxDispatchIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIdIterator{contract: _Mailbox.contract, event: "DispatchId", logs: logs, sub: sub}, nil
}

// WatchDispatchId is a free log subscription operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchDispatchId(opts *bind.WatchOpts, sink chan<- *MailboxDispatchId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatchId)
				if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
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

// ParseDispatchId is a log parse operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseDispatchId(log types.Log) (*MailboxDispatchId, error) {
	event := new(MailboxDispatchId)
	if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mailbox contract.
type MailboxInitializedIterator struct {
	Event *MailboxInitialized // Event containing the contract specifics and raw log

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
func (it *MailboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxInitialized)
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
		it.Event = new(MailboxInitialized)
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
func (it *MailboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxInitialized represents a Initialized event raised by the Mailbox contract.
type MailboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*MailboxInitializedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MailboxInitializedIterator{contract: _Mailbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MailboxInitialized) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxInitialized)
				if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) ParseInitialized(log types.Log) (*MailboxInitialized, error) {
	event := new(MailboxInitialized)
	if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mailbox contract.
type MailboxOwnershipTransferredIterator struct {
	Event *MailboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MailboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxOwnershipTransferred)
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
		it.Event = new(MailboxOwnershipTransferred)
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
func (it *MailboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxOwnershipTransferred represents a OwnershipTransferred event raised by the Mailbox contract.
type MailboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MailboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MailboxOwnershipTransferredIterator{contract: _Mailbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MailboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxOwnershipTransferred)
				if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mailbox *MailboxFilterer) ParseOwnershipTransferred(log types.Log) (*MailboxOwnershipTransferred, error) {
	event := new(MailboxOwnershipTransferred)
	if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Mailbox contract.
type MailboxPausedIterator struct {
	Event *MailboxPaused // Event containing the contract specifics and raw log

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
func (it *MailboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxPaused)
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
		it.Event = new(MailboxPaused)
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
func (it *MailboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxPaused represents a Paused event raised by the Mailbox contract.
type MailboxPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) FilterPaused(opts *bind.FilterOpts) (*MailboxPausedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MailboxPausedIterator{contract: _Mailbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MailboxPaused) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxPaused)
				if err := _Mailbox.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) ParsePaused(log types.Log) (*MailboxPaused, error) {
	event := new(MailboxPaused)
	if err := _Mailbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the Mailbox contract.
type MailboxProcessIterator struct {
	Event *MailboxProcess // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcess)
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
		it.Event = new(MailboxProcess)
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
func (it *MailboxProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcess represents a Process event raised by the Mailbox contract.
type MailboxProcess struct {
	Origin    uint32
	Sender    [32]byte
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) FilterProcess(opts *bind.FilterOpts, origin []uint32, sender [][32]byte, recipient []common.Address) (*MailboxProcessIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIterator{contract: _Mailbox.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *MailboxProcess, origin []uint32, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcess)
				if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) ParseProcess(log types.Log) (*MailboxProcess, error) {
	event := new(MailboxProcess)
	if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIdIterator is returned from FilterProcessId and is used to iterate over the raw logs and unpacked data for ProcessId events raised by the Mailbox contract.
type MailboxProcessIdIterator struct {
	Event *MailboxProcessId // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcessId)
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
		it.Event = new(MailboxProcessId)
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
func (it *MailboxProcessIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcessId represents a ProcessId event raised by the Mailbox contract.
type MailboxProcessId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessId is a free log retrieval operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterProcessId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxProcessIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIdIterator{contract: _Mailbox.contract, event: "ProcessId", logs: logs, sub: sub}, nil
}

// WatchProcessId is a free log subscription operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchProcessId(opts *bind.WatchOpts, sink chan<- *MailboxProcessId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcessId)
				if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
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

// ParseProcessId is a log parse operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseProcessId(log types.Log) (*MailboxProcessId, error) {
	event := new(MailboxProcessId)
	if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Mailbox contract.
type MailboxUnpausedIterator struct {
	Event *MailboxUnpaused // Event containing the contract specifics and raw log

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
func (it *MailboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxUnpaused)
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
		it.Event = new(MailboxUnpaused)
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
func (it *MailboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxUnpaused represents a Unpaused event raised by the Mailbox contract.
type MailboxUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MailboxUnpausedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MailboxUnpausedIterator{contract: _Mailbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MailboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxUnpaused)
				if err := _Mailbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) ParseUnpaused(log types.Log) (*MailboxUnpaused, error) {
	event := new(MailboxUnpaused)
	if err := _Mailbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
