// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gameAccount

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

// GameAccountMetaData contains all meta data concerning the GameAccount contract.
var GameAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"account\",\"type\":\"bytes32\"}],\"name\":\"Binding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delegatedAccount\",\"type\":\"bytes32\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UnBinding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delegatedAccount\",\"type\":\"bytes32\"}],\"name\":\"UnDelegated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"email_\",\"type\":\"string\"}],\"name\":\"bindAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"email_\",\"type\":\"string\"}],\"name\":\"calcEmailHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"email_\",\"type\":\"string\"}],\"name\":\"delegateAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getDelegatedEmailHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"getEmailHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"pullFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"email_\",\"type\":\"string\"}],\"name\":\"undelegateAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611757806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80639010d07c116100ad578063ca15c87311610071578063ca15c8731461027c578063d547741f1461028f578063ec079caa146102a2578063fdf6e47d146102b5578063fe84244f146102c857600080fd5b80639010d07c1461021b57806390b1cb271461024657806391d148541461024e578063a217fddf14610261578063c4740a951461026957600080fd5b80633adde5bf116100f45780633adde5bf146101c55780635d5a92fb146101d85780636239b7b1146101eb57806375b238fc146101fe5780638129fc1c1461021357600080fd5b806301ffc9a7146101315780630c926d1f14610159578063248a9ca31461016e5780632f2ff15d1461019f57806336568abe146101b2575b600080fd5b61014461013f366004611321565b6102db565b60405190151581526020015b60405180910390f35b61016c610167366004611361565b610306565b005b61019161017c366004611412565b60009081526065602052604090206001015490565b604051908152602001610150565b61016c6101ad366004611447565b610480565b61016c6101c0366004611447565b6104ab565b6101916101d3366004611473565b610529565b61016c6101e6366004611473565b61055a565b61016c6101f9366004611361565b61069b565b61019160008051602061170283398151915281565b61016c610800565b61022e61022936600461148e565b6108e4565b6040516001600160a01b039091168152602001610150565b61016c610903565b61014461025c366004611447565b61095b565b610191600081565b61016c610277366004611473565b610986565b61019161028a366004611412565b6109a8565b61016c61029d366004611447565b6109bf565b6101916102b0366004611361565b6109e5565b61016c6102c3366004611361565b6109f0565b6101916102d6366004611473565b610b68565b60006001600160e01b03198216635a05180f60e01b1480610300575061030082610b99565b92915050565b600081511161031457600080fd5b600061031f82610bce565b33600090815260ca6020526040902054909150801561035557600081815260cc6020526040902080546001600160a01b03191690555b600082815260cb60205260409020546001600160a01b0316156103e55760405162461bcd60e51b815260206004820152603760248201527f54686520656d61696c2068617320616c7265616479206265656e20726567697360448201527f7465726564206173206d6173746572206163636f756e7400000000000000000060648201526084015b60405180910390fd5b600082815260cc60205260409020546001600160a01b03161561041a5760405162461bcd60e51b81526004016103dc906114b0565b33600081815260ca6020908152604080832086905585835260cc82529182902080546001600160a01b0319168417905590518481527fa6ca69be1634c9486160d4fa9f11c9bf604a6a4b1fd23c8336ffc5889ef4b5ab91015b60405180910390a2505050565b60008281526065602052604090206001015461049c8133610c13565b6104a68383610c77565b505050565b6001600160a01b038116331461051b5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016103dc565b6105258282610c99565b5050565b60006001600160a01b03821661053e57600080fd5b506001600160a01b0316600090815260c9602052604090205490565b6000805160206117028339815191526105738133610c13565b6001600160a01b0382166105ad5760405133904780156108fc02916000818181858888f193505050501580156104a6573d6000803e3d6000fd5b816001600160a01b03811663a9059cbb336040516370a0823160e01b81523060048201526001600160a01b038516906370a0823190602401602060405180830381865afa158015610602573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610626919061150d565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015610671573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106959190611526565b50505050565b60008151116106a957600080fd5b60006106b482610bce565b33600090815260c9602052604090205490915080156106ea57600081815260cb6020526040902080546001600160a01b03191690555b600082815260cb60205260409020546001600160a01b03161561075d5760405162461bcd60e51b815260206004820152602560248201527f54686520656d61696c2068617320616c7265616479206265656e2072656769736044820152641d195c995960da1b60648201526084016103dc565b600082815260cc60205260409020546001600160a01b0316156107925760405162461bcd60e51b81526004016103dc906114b0565b33600081815260c96020908152604080832086905585835260cb909152902080546001600160a01b031916821790556107cd9060cd90610cbb565b5060405182815233907fcd73894d1869cdb933f1b053a0babed16ab19fd06a735373f9ba1792bc0da2a290602001610473565b600054610100900460ff1661081b5760005460ff161561081f565b303b155b6108825760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103dc565b600054610100900460ff161580156108a4576000805461ffff19166101011790555b6108ac610cd0565b6108b7600033610d3d565b6108cf60008051602061170283398151915233610d3d565b80156108e1576000805461ff00191690555b50565b60008281526097602052604081206108fc9083610d47565b9392505050565b60008051602061170283398151915261091c8133610c13565b600061092860cd610d53565b905060005b818110156104a65761094961094460cd6000610d47565b610d5d565b806109538161155e565b91505061092d565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60008051602061170283398151915261099f8133610c13565b61052582610d5d565b600081815260976020526040812061030090610d53565b6000828152606560205260409020600101546109db8133610c13565b6104a68383610c99565b600061030082610bce565b60008151116109fe57600080fd5b6000610a0982610bce565b33600090815260ca60205260409020549091508114610a885760405162461bcd60e51b815260206004820152603560248201527f54686520656d61696c20776173206e6f742072656769737465726564206173206044820152741d1a194819195b1959d85d1959081858d8dbdd5b9d605a1b60648201526084016103dc565b600081815260cc60205260409020546001600160a01b0316610b065760405162461bcd60e51b815260206004820152603160248201527f54686520656d61696c20776173206e6f7420726567697374657265642061732060448201527019195b1959d85d1959081858d8dbdd5b9d607a1b60648201526084016103dc565b33600081815260ca6020908152604080832083905584835260cc82529182902080546001600160a01b031916905590518381527fb5261b10ff84e4fced2848e663a8a79ba863dbf0f48fc6f07f8204c8de7a4274910160405180910390a25050565b60006001600160a01b038216610b7d57600080fd5b506001600160a01b0316600090815260ca602052604090205490565b60006001600160e01b03198216637965db0b60e01b148061030057506301ffc9a760e01b6001600160e01b0319831614610300565b600080825111610bdd57600080fd5b610be682610e4e565b604051602001610bf691906115a3565b604051602081830303815290604052805190602001209050919050565b610c1d828261095b565b61052557610c35816001600160a01b03166014610ec8565b610c40836020610ec8565b604051602001610c519291906115d6565b60408051601f198184030181529082905262461bcd60e51b82526103dc916004016115a3565b610c818282611064565b60008281526097602052604090206104a69082610cbb565b610ca382826110ea565b60008281526097602052604090206104a69082611151565b60006108fc836001600160a01b038416611166565b600054610100900460ff16610d3b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016103dc565b565b6105258282610c77565b60006108fc83836111b5565b6000610300825490565b6001600160a01b038116600090815260c960205260409020548015610db3576001600160a01b038216600090815260c96020908152604080832083905583835260cb909152902080546001600160a01b03191690555b6001600160a01b038216600090815260ca60205260409020548015610e09576001600160a01b038316600090815260ca6020908152604080832083905583835260cc909152902080546001600160a01b03191690555b610e1460cd84611151565b506040516001600160a01b038416907f0f8f522aef36d3c81550e748a21b2640596a14aa087111ab9df0f22bf8fb2c2790600090a2505050565b60608160005b8151811015610ec157610e86828281518110610e7257610e7261164b565b01602001516001600160f81b0319166111df565b828281518110610e9857610e9861164b565b60200101906001600160f81b031916908160001a90535080610eb98161155e565b915050610e54565b5092915050565b60606000610ed7836002611661565b610ee2906002611680565b67ffffffffffffffff811115610efa57610efa61134b565b6040519080825280601f01601f191660200182016040528015610f24576020820181803683370190505b509050600360fc1b81600081518110610f3f57610f3f61164b565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110610f6e57610f6e61164b565b60200101906001600160f81b031916908160001a9053506000610f92846002611661565b610f9d906001611680565b90505b6001811115611015576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110610fd157610fd161164b565b1a60f81b828281518110610fe757610fe761164b565b60200101906001600160f81b031916908160001a90535060049490941c9361100e81611698565b9050610fa0565b5083156108fc5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103dc565b61106e828261095b565b6105255760008281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556110a63390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6110f4828261095b565b156105255760008281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60006108fc836001600160a01b03841661122e565b60008181526001830160205260408120546111ad57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610300565b506000610300565b60008260000182815481106111cc576111cc61164b565b9060005260206000200154905092915050565b6000604160f81b6001600160f81b031983161080159061120d5750602d60f91b6001600160f81b0319831611155b1561122a5761122160f883901c60206116af565b60f81b92915050565b5090565b600081815260018301602052604081205480156113175760006112526001836116d4565b8554909150600090611266906001906116d4565b90508181146112cb5760008660000182815481106112865761128661164b565b90600052602060002001549050808760000184815481106112a9576112a961164b565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806112dc576112dc6116eb565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610300565b6000915050610300565b60006020828403121561133357600080fd5b81356001600160e01b0319811681146108fc57600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561137357600080fd5b813567ffffffffffffffff8082111561138b57600080fd5b818401915084601f83011261139f57600080fd5b8135818111156113b1576113b161134b565b604051601f8201601f19908116603f011681019083821181831017156113d9576113d961134b565b816040528281528760208487010111156113f257600080fd5b826020860160208301376000928101602001929092525095945050505050565b60006020828403121561142457600080fd5b5035919050565b80356001600160a01b038116811461144257600080fd5b919050565b6000806040838503121561145a57600080fd5b8235915061146a6020840161142b565b90509250929050565b60006020828403121561148557600080fd5b6108fc8261142b565b600080604083850312156114a157600080fd5b50508035926020909101359150565b6020808252603a908201527f54686520656d61696c2068617320616c7265616479206265656e20726567697360408201527f74657265642061732064656c656761746564206163636f756e74000000000000606082015260800190565b60006020828403121561151f57600080fd5b5051919050565b60006020828403121561153857600080fd5b815180151581146108fc57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161157057611570611548565b5060010190565b60005b8381101561159257818101518382015260200161157a565b838111156106955750506000910152565b60208152600082518060208401526115c2816040850160208701611577565b601f01601f19169190910160400192915050565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081526000835161160e816017850160208801611577565b7001034b99036b4b9b9b4b733903937b6329607d1b601791840191820152835161163f816028840160208801611577565b01602801949350505050565b634e487b7160e01b600052603260045260246000fd5b600081600019048311821515161561167b5761167b611548565b500290565b6000821982111561169357611693611548565b500190565b6000816116a7576116a7611548565b506000190190565b600060ff821660ff84168060ff038211156116cc576116cc611548565b019392505050565b6000828210156116e6576116e6611548565b500390565b634e487b7160e01b600052603160045260246000fdfea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220c8d8c2e85cea9681465c3698110044d724f62cdb72815854e5da9869421e02fd64736f6c634300080d0033",
}

// GameAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use GameAccountMetaData.ABI instead.
var GameAccountABI = GameAccountMetaData.ABI

// GameAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GameAccountMetaData.Bin instead.
var GameAccountBin = GameAccountMetaData.Bin

// DeployGameAccount deploys a new Ethereum contract, binding an instance of GameAccount to it.
func DeployGameAccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameAccount, error) {
	parsed, err := GameAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GameAccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameAccount{GameAccountCaller: GameAccountCaller{contract: contract}, GameAccountTransactor: GameAccountTransactor{contract: contract}, GameAccountFilterer: GameAccountFilterer{contract: contract}}, nil
}

// GameAccount is an auto generated Go binding around an Ethereum contract.
type GameAccount struct {
	GameAccountCaller     // Read-only binding to the contract
	GameAccountTransactor // Write-only binding to the contract
	GameAccountFilterer   // Log filterer for contract events
}

// GameAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameAccountSession struct {
	Contract     *GameAccount      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameAccountCallerSession struct {
	Contract *GameAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GameAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameAccountTransactorSession struct {
	Contract     *GameAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GameAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameAccountRaw struct {
	Contract *GameAccount // Generic contract binding to access the raw methods on
}

// GameAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameAccountCallerRaw struct {
	Contract *GameAccountCaller // Generic read-only contract binding to access the raw methods on
}

// GameAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameAccountTransactorRaw struct {
	Contract *GameAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameAccount creates a new instance of GameAccount, bound to a specific deployed contract.
func NewGameAccount(address common.Address, backend bind.ContractBackend) (*GameAccount, error) {
	contract, err := bindGameAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameAccount{GameAccountCaller: GameAccountCaller{contract: contract}, GameAccountTransactor: GameAccountTransactor{contract: contract}, GameAccountFilterer: GameAccountFilterer{contract: contract}}, nil
}

// NewGameAccountCaller creates a new read-only instance of GameAccount, bound to a specific deployed contract.
func NewGameAccountCaller(address common.Address, caller bind.ContractCaller) (*GameAccountCaller, error) {
	contract, err := bindGameAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameAccountCaller{contract: contract}, nil
}

// NewGameAccountTransactor creates a new write-only instance of GameAccount, bound to a specific deployed contract.
func NewGameAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*GameAccountTransactor, error) {
	contract, err := bindGameAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameAccountTransactor{contract: contract}, nil
}

// NewGameAccountFilterer creates a new log filterer instance of GameAccount, bound to a specific deployed contract.
func NewGameAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*GameAccountFilterer, error) {
	contract, err := bindGameAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameAccountFilterer{contract: contract}, nil
}

// bindGameAccount binds a generic wrapper to an already deployed contract.
func bindGameAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameAccount *GameAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameAccount.Contract.GameAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameAccount *GameAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameAccount.Contract.GameAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameAccount *GameAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameAccount.Contract.GameAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameAccount *GameAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameAccount *GameAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameAccount *GameAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameAccount.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountSession) ADMINROLE() ([32]byte, error) {
	return _GameAccount.Contract.ADMINROLE(&_GameAccount.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountCallerSession) ADMINROLE() ([32]byte, error) {
	return _GameAccount.Contract.ADMINROLE(&_GameAccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameAccount.Contract.DEFAULTADMINROLE(&_GameAccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GameAccount *GameAccountCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GameAccount.Contract.DEFAULTADMINROLE(&_GameAccount.CallOpts)
}

// CalcEmailHash is a free data retrieval call binding the contract method 0xec079caa.
//
// Solidity: function calcEmailHash(string email_) pure returns(bytes32)
func (_GameAccount *GameAccountCaller) CalcEmailHash(opts *bind.CallOpts, email_ string) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "calcEmailHash", email_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalcEmailHash is a free data retrieval call binding the contract method 0xec079caa.
//
// Solidity: function calcEmailHash(string email_) pure returns(bytes32)
func (_GameAccount *GameAccountSession) CalcEmailHash(email_ string) ([32]byte, error) {
	return _GameAccount.Contract.CalcEmailHash(&_GameAccount.CallOpts, email_)
}

// CalcEmailHash is a free data retrieval call binding the contract method 0xec079caa.
//
// Solidity: function calcEmailHash(string email_) pure returns(bytes32)
func (_GameAccount *GameAccountCallerSession) CalcEmailHash(email_ string) ([32]byte, error) {
	return _GameAccount.Contract.CalcEmailHash(&_GameAccount.CallOpts, email_)
}

// GetDelegatedEmailHash is a free data retrieval call binding the contract method 0xfe84244f.
//
// Solidity: function getDelegatedEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountCaller) GetDelegatedEmailHash(opts *bind.CallOpts, account_ common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "getDelegatedEmailHash", account_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedEmailHash is a free data retrieval call binding the contract method 0xfe84244f.
//
// Solidity: function getDelegatedEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountSession) GetDelegatedEmailHash(account_ common.Address) ([32]byte, error) {
	return _GameAccount.Contract.GetDelegatedEmailHash(&_GameAccount.CallOpts, account_)
}

// GetDelegatedEmailHash is a free data retrieval call binding the contract method 0xfe84244f.
//
// Solidity: function getDelegatedEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountCallerSession) GetDelegatedEmailHash(account_ common.Address) ([32]byte, error) {
	return _GameAccount.Contract.GetDelegatedEmailHash(&_GameAccount.CallOpts, account_)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountCaller) GetEmailHash(opts *bind.CallOpts, account_ common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "getEmailHash", account_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountSession) GetEmailHash(account_ common.Address) ([32]byte, error) {
	return _GameAccount.Contract.GetEmailHash(&_GameAccount.CallOpts, account_)
}

// GetEmailHash is a free data retrieval call binding the contract method 0x3adde5bf.
//
// Solidity: function getEmailHash(address account_) view returns(bytes32)
func (_GameAccount *GameAccountCallerSession) GetEmailHash(account_ common.Address) ([32]byte, error) {
	return _GameAccount.Contract.GetEmailHash(&_GameAccount.CallOpts, account_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameAccount *GameAccountCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameAccount *GameAccountSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameAccount.Contract.GetRoleAdmin(&_GameAccount.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GameAccount *GameAccountCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GameAccount.Contract.GetRoleAdmin(&_GameAccount.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GameAccount *GameAccountCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GameAccount *GameAccountSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _GameAccount.Contract.GetRoleMember(&_GameAccount.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_GameAccount *GameAccountCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _GameAccount.Contract.GetRoleMember(&_GameAccount.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GameAccount *GameAccountCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GameAccount *GameAccountSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _GameAccount.Contract.GetRoleMemberCount(&_GameAccount.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_GameAccount *GameAccountCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _GameAccount.Contract.GetRoleMemberCount(&_GameAccount.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameAccount *GameAccountCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameAccount *GameAccountSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameAccount.Contract.HasRole(&_GameAccount.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GameAccount *GameAccountCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GameAccount.Contract.HasRole(&_GameAccount.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameAccount *GameAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GameAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameAccount *GameAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameAccount.Contract.SupportsInterface(&_GameAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GameAccount *GameAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GameAccount.Contract.SupportsInterface(&_GameAccount.CallOpts, interfaceId)
}

// BindAccount is a paid mutator transaction binding the contract method 0x6239b7b1.
//
// Solidity: function bindAccount(string email_) returns()
func (_GameAccount *GameAccountTransactor) BindAccount(opts *bind.TransactOpts, email_ string) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "bindAccount", email_)
}

// BindAccount is a paid mutator transaction binding the contract method 0x6239b7b1.
//
// Solidity: function bindAccount(string email_) returns()
func (_GameAccount *GameAccountSession) BindAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.BindAccount(&_GameAccount.TransactOpts, email_)
}

// BindAccount is a paid mutator transaction binding the contract method 0x6239b7b1.
//
// Solidity: function bindAccount(string email_) returns()
func (_GameAccount *GameAccountTransactorSession) BindAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.BindAccount(&_GameAccount.TransactOpts, email_)
}

// DelegateAccount is a paid mutator transaction binding the contract method 0x0c926d1f.
//
// Solidity: function delegateAccount(string email_) returns()
func (_GameAccount *GameAccountTransactor) DelegateAccount(opts *bind.TransactOpts, email_ string) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "delegateAccount", email_)
}

// DelegateAccount is a paid mutator transaction binding the contract method 0x0c926d1f.
//
// Solidity: function delegateAccount(string email_) returns()
func (_GameAccount *GameAccountSession) DelegateAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.DelegateAccount(&_GameAccount.TransactOpts, email_)
}

// DelegateAccount is a paid mutator transaction binding the contract method 0x0c926d1f.
//
// Solidity: function delegateAccount(string email_) returns()
func (_GameAccount *GameAccountTransactorSession) DelegateAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.DelegateAccount(&_GameAccount.TransactOpts, email_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.GrantRole(&_GameAccount.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.GrantRole(&_GameAccount.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameAccount *GameAccountTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameAccount *GameAccountSession) Initialize() (*types.Transaction, error) {
	return _GameAccount.Contract.Initialize(&_GameAccount.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GameAccount *GameAccountTransactorSession) Initialize() (*types.Transaction, error) {
	return _GameAccount.Contract.Initialize(&_GameAccount.TransactOpts)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_GameAccount *GameAccountTransactor) PullFunds(opts *bind.TransactOpts, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "pullFunds", tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_GameAccount *GameAccountSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.PullFunds(&_GameAccount.TransactOpts, tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_GameAccount *GameAccountTransactorSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.PullFunds(&_GameAccount.TransactOpts, tokenAddress_)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account_) returns()
func (_GameAccount *GameAccountTransactor) RemoveAccount(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "removeAccount", account_)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account_) returns()
func (_GameAccount *GameAccountSession) RemoveAccount(account_ common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RemoveAccount(&_GameAccount.TransactOpts, account_)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account_) returns()
func (_GameAccount *GameAccountTransactorSession) RemoveAccount(account_ common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RemoveAccount(&_GameAccount.TransactOpts, account_)
}

// RemoveAllAccounts is a paid mutator transaction binding the contract method 0x90b1cb27.
//
// Solidity: function removeAllAccounts() returns()
func (_GameAccount *GameAccountTransactor) RemoveAllAccounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "removeAllAccounts")
}

// RemoveAllAccounts is a paid mutator transaction binding the contract method 0x90b1cb27.
//
// Solidity: function removeAllAccounts() returns()
func (_GameAccount *GameAccountSession) RemoveAllAccounts() (*types.Transaction, error) {
	return _GameAccount.Contract.RemoveAllAccounts(&_GameAccount.TransactOpts)
}

// RemoveAllAccounts is a paid mutator transaction binding the contract method 0x90b1cb27.
//
// Solidity: function removeAllAccounts() returns()
func (_GameAccount *GameAccountTransactorSession) RemoveAllAccounts() (*types.Transaction, error) {
	return _GameAccount.Contract.RemoveAllAccounts(&_GameAccount.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RenounceRole(&_GameAccount.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RenounceRole(&_GameAccount.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RevokeRole(&_GameAccount.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GameAccount *GameAccountTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GameAccount.Contract.RevokeRole(&_GameAccount.TransactOpts, role, account)
}

// UndelegateAccount is a paid mutator transaction binding the contract method 0xfdf6e47d.
//
// Solidity: function undelegateAccount(string email_) returns()
func (_GameAccount *GameAccountTransactor) UndelegateAccount(opts *bind.TransactOpts, email_ string) (*types.Transaction, error) {
	return _GameAccount.contract.Transact(opts, "undelegateAccount", email_)
}

// UndelegateAccount is a paid mutator transaction binding the contract method 0xfdf6e47d.
//
// Solidity: function undelegateAccount(string email_) returns()
func (_GameAccount *GameAccountSession) UndelegateAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.UndelegateAccount(&_GameAccount.TransactOpts, email_)
}

// UndelegateAccount is a paid mutator transaction binding the contract method 0xfdf6e47d.
//
// Solidity: function undelegateAccount(string email_) returns()
func (_GameAccount *GameAccountTransactorSession) UndelegateAccount(email_ string) (*types.Transaction, error) {
	return _GameAccount.Contract.UndelegateAccount(&_GameAccount.TransactOpts, email_)
}

// GameAccountBindingIterator is returned from FilterBinding and is used to iterate over the raw logs and unpacked data for Binding events raised by the GameAccount contract.
type GameAccountBindingIterator struct {
	Event *GameAccountBinding // Event containing the contract specifics and raw log

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
func (it *GameAccountBindingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountBinding)
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
		it.Event = new(GameAccountBinding)
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
func (it *GameAccountBindingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountBindingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountBinding represents a Binding event raised by the GameAccount contract.
type GameAccountBinding struct {
	User    common.Address
	Account [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBinding is a free log retrieval operation binding the contract event 0xcd73894d1869cdb933f1b053a0babed16ab19fd06a735373f9ba1792bc0da2a2.
//
// Solidity: event Binding(address indexed user, bytes32 account)
func (_GameAccount *GameAccountFilterer) FilterBinding(opts *bind.FilterOpts, user []common.Address) (*GameAccountBindingIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "Binding", userRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountBindingIterator{contract: _GameAccount.contract, event: "Binding", logs: logs, sub: sub}, nil
}

// WatchBinding is a free log subscription operation binding the contract event 0xcd73894d1869cdb933f1b053a0babed16ab19fd06a735373f9ba1792bc0da2a2.
//
// Solidity: event Binding(address indexed user, bytes32 account)
func (_GameAccount *GameAccountFilterer) WatchBinding(opts *bind.WatchOpts, sink chan<- *GameAccountBinding, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "Binding", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountBinding)
				if err := _GameAccount.contract.UnpackLog(event, "Binding", log); err != nil {
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

// ParseBinding is a log parse operation binding the contract event 0xcd73894d1869cdb933f1b053a0babed16ab19fd06a735373f9ba1792bc0da2a2.
//
// Solidity: event Binding(address indexed user, bytes32 account)
func (_GameAccount *GameAccountFilterer) ParseBinding(log types.Log) (*GameAccountBinding, error) {
	event := new(GameAccountBinding)
	if err := _GameAccount.contract.UnpackLog(event, "Binding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the GameAccount contract.
type GameAccountDelegatedIterator struct {
	Event *GameAccountDelegated // Event containing the contract specifics and raw log

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
func (it *GameAccountDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountDelegated)
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
		it.Event = new(GameAccountDelegated)
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
func (it *GameAccountDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountDelegated represents a Delegated event raised by the GameAccount contract.
type GameAccountDelegated struct {
	User             common.Address
	DelegatedAccount [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xa6ca69be1634c9486160d4fa9f11c9bf604a6a4b1fd23c8336ffc5889ef4b5ab.
//
// Solidity: event Delegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) FilterDelegated(opts *bind.FilterOpts, user []common.Address) (*GameAccountDelegatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "Delegated", userRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountDelegatedIterator{contract: _GameAccount.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xa6ca69be1634c9486160d4fa9f11c9bf604a6a4b1fd23c8336ffc5889ef4b5ab.
//
// Solidity: event Delegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *GameAccountDelegated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "Delegated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountDelegated)
				if err := _GameAccount.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xa6ca69be1634c9486160d4fa9f11c9bf604a6a4b1fd23c8336ffc5889ef4b5ab.
//
// Solidity: event Delegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) ParseDelegated(log types.Log) (*GameAccountDelegated, error) {
	event := new(GameAccountDelegated)
	if err := _GameAccount.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GameAccount contract.
type GameAccountRoleAdminChangedIterator struct {
	Event *GameAccountRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GameAccountRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountRoleAdminChanged)
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
		it.Event = new(GameAccountRoleAdminChanged)
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
func (it *GameAccountRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountRoleAdminChanged represents a RoleAdminChanged event raised by the GameAccount contract.
type GameAccountRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameAccount *GameAccountFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GameAccountRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountRoleAdminChangedIterator{contract: _GameAccount.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameAccount *GameAccountFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GameAccountRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountRoleAdminChanged)
				if err := _GameAccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GameAccount *GameAccountFilterer) ParseRoleAdminChanged(log types.Log) (*GameAccountRoleAdminChanged, error) {
	event := new(GameAccountRoleAdminChanged)
	if err := _GameAccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GameAccount contract.
type GameAccountRoleGrantedIterator struct {
	Event *GameAccountRoleGranted // Event containing the contract specifics and raw log

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
func (it *GameAccountRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountRoleGranted)
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
		it.Event = new(GameAccountRoleGranted)
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
func (it *GameAccountRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountRoleGranted represents a RoleGranted event raised by the GameAccount contract.
type GameAccountRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameAccountRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountRoleGrantedIterator{contract: _GameAccount.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GameAccountRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountRoleGranted)
				if err := _GameAccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) ParseRoleGranted(log types.Log) (*GameAccountRoleGranted, error) {
	event := new(GameAccountRoleGranted)
	if err := _GameAccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GameAccount contract.
type GameAccountRoleRevokedIterator struct {
	Event *GameAccountRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GameAccountRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountRoleRevoked)
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
		it.Event = new(GameAccountRoleRevoked)
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
func (it *GameAccountRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountRoleRevoked represents a RoleRevoked event raised by the GameAccount contract.
type GameAccountRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GameAccountRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountRoleRevokedIterator{contract: _GameAccount.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GameAccountRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountRoleRevoked)
				if err := _GameAccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GameAccount *GameAccountFilterer) ParseRoleRevoked(log types.Log) (*GameAccountRoleRevoked, error) {
	event := new(GameAccountRoleRevoked)
	if err := _GameAccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountUnBindingIterator is returned from FilterUnBinding and is used to iterate over the raw logs and unpacked data for UnBinding events raised by the GameAccount contract.
type GameAccountUnBindingIterator struct {
	Event *GameAccountUnBinding // Event containing the contract specifics and raw log

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
func (it *GameAccountUnBindingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountUnBinding)
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
		it.Event = new(GameAccountUnBinding)
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
func (it *GameAccountUnBindingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountUnBindingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountUnBinding represents a UnBinding event raised by the GameAccount contract.
type GameAccountUnBinding struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnBinding is a free log retrieval operation binding the contract event 0x0f8f522aef36d3c81550e748a21b2640596a14aa087111ab9df0f22bf8fb2c27.
//
// Solidity: event UnBinding(address indexed user)
func (_GameAccount *GameAccountFilterer) FilterUnBinding(opts *bind.FilterOpts, user []common.Address) (*GameAccountUnBindingIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "UnBinding", userRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountUnBindingIterator{contract: _GameAccount.contract, event: "UnBinding", logs: logs, sub: sub}, nil
}

// WatchUnBinding is a free log subscription operation binding the contract event 0x0f8f522aef36d3c81550e748a21b2640596a14aa087111ab9df0f22bf8fb2c27.
//
// Solidity: event UnBinding(address indexed user)
func (_GameAccount *GameAccountFilterer) WatchUnBinding(opts *bind.WatchOpts, sink chan<- *GameAccountUnBinding, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "UnBinding", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountUnBinding)
				if err := _GameAccount.contract.UnpackLog(event, "UnBinding", log); err != nil {
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

// ParseUnBinding is a log parse operation binding the contract event 0x0f8f522aef36d3c81550e748a21b2640596a14aa087111ab9df0f22bf8fb2c27.
//
// Solidity: event UnBinding(address indexed user)
func (_GameAccount *GameAccountFilterer) ParseUnBinding(log types.Log) (*GameAccountUnBinding, error) {
	event := new(GameAccountUnBinding)
	if err := _GameAccount.contract.UnpackLog(event, "UnBinding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameAccountUnDelegatedIterator is returned from FilterUnDelegated and is used to iterate over the raw logs and unpacked data for UnDelegated events raised by the GameAccount contract.
type GameAccountUnDelegatedIterator struct {
	Event *GameAccountUnDelegated // Event containing the contract specifics and raw log

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
func (it *GameAccountUnDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameAccountUnDelegated)
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
		it.Event = new(GameAccountUnDelegated)
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
func (it *GameAccountUnDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameAccountUnDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameAccountUnDelegated represents a UnDelegated event raised by the GameAccount contract.
type GameAccountUnDelegated struct {
	User             common.Address
	DelegatedAccount [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnDelegated is a free log retrieval operation binding the contract event 0xb5261b10ff84e4fced2848e663a8a79ba863dbf0f48fc6f07f8204c8de7a4274.
//
// Solidity: event UnDelegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) FilterUnDelegated(opts *bind.FilterOpts, user []common.Address) (*GameAccountUnDelegatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.FilterLogs(opts, "UnDelegated", userRule)
	if err != nil {
		return nil, err
	}
	return &GameAccountUnDelegatedIterator{contract: _GameAccount.contract, event: "UnDelegated", logs: logs, sub: sub}, nil
}

// WatchUnDelegated is a free log subscription operation binding the contract event 0xb5261b10ff84e4fced2848e663a8a79ba863dbf0f48fc6f07f8204c8de7a4274.
//
// Solidity: event UnDelegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) WatchUnDelegated(opts *bind.WatchOpts, sink chan<- *GameAccountUnDelegated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GameAccount.contract.WatchLogs(opts, "UnDelegated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameAccountUnDelegated)
				if err := _GameAccount.contract.UnpackLog(event, "UnDelegated", log); err != nil {
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

// ParseUnDelegated is a log parse operation binding the contract event 0xb5261b10ff84e4fced2848e663a8a79ba863dbf0f48fc6f07f8204c8de7a4274.
//
// Solidity: event UnDelegated(address indexed user, bytes32 delegatedAccount)
func (_GameAccount *GameAccountFilterer) ParseUnDelegated(log types.Log) (*GameAccountUnDelegated, error) {
	event := new(GameAccountUnDelegated)
	if err := _GameAccount.contract.UnpackLog(event, "UnDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
