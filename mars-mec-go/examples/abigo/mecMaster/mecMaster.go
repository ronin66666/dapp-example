// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mecMaster

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

// MecMasterMetaData contains all meta data concerning the MecMaster contract.
var MecMasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mecToken\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount_\",\"type\":\"uint256[]\"}],\"name\":\"multiSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"pullFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"resetAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setMecAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611724806100206000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063a217fddf116100a2578063c653fe2311610071578063c653fe2314610268578063ca15c8731461027b578063d547741f1461028e578063de242ff4146102a1578063eee4f11b146102aa57600080fd5b8063a217fddf14610227578063b6b55f251461022f578063bb4c9f0b14610242578063c4d66de81461025557600080fd5b806336568abe116100e957806336568abe146101b45780635d5a92fb146101c757806375b238fc146101da5780639010d07c1461020157806391d148541461021457600080fd5b806301ffc9a71461011b578063248a9ca3146101435780632f2ff15d1461017457806334b412b914610189575b600080fd5b61012e6101293660046112a1565b6102bd565b60405190151581526020015b60405180910390f35b6101666101513660046112cb565b60009081526065602052604090206001015490565b60405190815260200161013a565b6101876101823660046112fc565b6102e8565b005b60c95461019c906001600160a01b031681565b6040516001600160a01b03909116815260200161013a565b6101876101c23660046112fc565b610313565b6101876101d536600461132c565b610396565b6101667fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b61019c61020f366004611349565b6104ca565b61012e6102223660046112fc565b6104e9565b610166600081565b61018761023d3660046112cb565b610514565b610187610250366004611441565b6105b8565b61018761026336600461132c565b610832565b61018761027636600461132c565b610956565b6101666102893660046112cb565b6109db565b61018761029c3660046112fc565b6109f2565b61016660ca5481565b6101876102b83660046112cb565b610a18565b60006001600160e01b03198216635a05180f60e01b14806102e257506102e282610a2a565b92915050565b6000828152606560205260409020600101546103048133610a5f565b61030e8383610ac3565b505050565b6001600160a01b03811633146103885760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6103928282610ae5565b5050565b60006103a28133610a5f565b6001600160a01b0382166103dc5760405133904780156108fc02916000818181858888f1935050505015801561030e573d6000803e3d6000fd5b816001600160a01b03811663a9059cbb336040516370a0823160e01b81523060048201526001600160a01b038516906370a0823190602401602060405180830381865afa158015610431573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104559190611503565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156104a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c4919061151c565b50505050565b60008281526097602052604081206104e29083610b07565b9392505050565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3332146105635760405162461bcd60e51b815260206004820152601760248201527f466f7263654e46544d61726b65743a206e6f7420656f61000000000000000000604482015260640161037f565b8060000361057057600080fd5b6105883360c9546001600160a01b0316903084610b13565b604051819033907f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c490600090a350565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756105e38133610a5f565b81518351146106345760405162461bcd60e51b815260206004820152601b60248201527f56616c7565206c656e6774687320646f206e6f74206d617463682e0000000000604482015260640161037f565b60008351116106775760405162461bcd60e51b815260206004820152600f60248201526e0546865206c656e677468206973203608c1b604482015260640161037f565b6000805b83518110156106bd578381815181106106965761069661153e565b6020026020010151826106a9919061156a565b9150806106b581611582565b91505061067b565b5060ca548111156107075760405162461bcd60e51b81526020600482015260146024820152736e6f7420656e6f75676820616c6c6f77616e636560601b604482015260640161037f565b8060ca6000828254610719919061159b565b90915550600090505b845181101561082b5760006001600160a01b03168582815181106107485761074861153e565b60200260200101516001600160a01b03160361076357600080fd5b6107b08582815181106107785761077861153e565b60200260200101518583815181106107925761079261153e565b602090810291909101015160c9546001600160a01b03169190610b7e565b8381815181106107c2576107c261153e565b60200260200101518582815181106107dc576107dc61153e565b60200260200101516001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436460405160405180910390a38061082381611582565b915050610722565b5050505050565b600054610100900460ff1661084d5760005460ff1615610851565b303b155b6108b45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161037f565b600054610100900460ff161580156108d6576000805461ffff19166101011790555b6108de610bae565b6001600160a01b0382166108f157600080fd5b60c980546001600160a01b0319166001600160a01b038416179055610917600033610c1b565b6109417fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533610c1b565b8015610392576000805461ff00191690555050565b60006109628133610a5f565b6001600160a01b0382166109b85760405162461bcd60e51b815260206004820152601c60248201527f5468652061646472657373206f6620746f6b656e206973206e756c6c00000000604482015260640161037f565b5060c980546001600160a01b0319166001600160a01b0392909216919091179055565b60008181526097602052604081206102e290610c25565b600082815260656020526040902060010154610a0e8133610a5f565b61030e8383610ae5565b6000610a248133610a5f565b5060ca55565b60006001600160e01b03198216637965db0b60e01b14806102e257506301ffc9a760e01b6001600160e01b03198316146102e2565b610a6982826104e9565b61039257610a81816001600160a01b03166014610c2f565b610a8c836020610c2f565b604051602001610a9d9291906115de565b60408051601f198184030181529082905262461bcd60e51b825261037f91600401611653565b610acd8282610dcb565b600082815260976020526040902061030e9082610e51565b610aef8282610e66565b600082815260976020526040902061030e9082610ecd565b60006104e28383610ee2565b6040516001600160a01b03808516602483015283166044820152606481018290526104c49085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f0c565b6040516001600160a01b03831660248201526044810182905261030e90849063a9059cbb60e01b90606401610b47565b600054610100900460ff16610c195760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161037f565b565b6103928282610ac3565b60006102e2825490565b60606000610c3e836002611686565b610c4990600261156a565b67ffffffffffffffff811115610c6157610c6161136b565b6040519080825280601f01601f191660200182016040528015610c8b576020820181803683370190505b509050600360fc1b81600081518110610ca657610ca661153e565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110610cd557610cd561153e565b60200101906001600160f81b031916908160001a9053506000610cf9846002611686565b610d0490600161156a565b90505b6001811115610d7c576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110610d3857610d3861153e565b1a60f81b828281518110610d4e57610d4e61153e565b60200101906001600160f81b031916908160001a90535060049490941c93610d75816116a5565b9050610d07565b5083156104e25760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161037f565b610dd582826104e9565b6103925760008281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610e0d3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60006104e2836001600160a01b038416610fde565b610e7082826104e9565b156103925760008281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60006104e2836001600160a01b03841661102d565b6000826000018281548110610ef957610ef961153e565b9060005260206000200154905092915050565b6000610f61826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111209092919063ffffffff16565b80519091501561030e5780806020019051810190610f7f919061151c565b61030e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161037f565b6000818152600183016020526040812054611025575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556102e2565b5060006102e2565b6000818152600183016020526040812054801561111657600061105160018361159b565b85549091506000906110659060019061159b565b90508181146110ca5760008660000182815481106110855761108561153e565b90600052602060002001549050808760000184815481106110a8576110a861153e565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806110db576110db6116bc565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506102e2565b60009150506102e2565b606061112f8484600085611137565b949350505050565b6060824710156111985760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161037f565b6001600160a01b0385163b6111ef5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161037f565b600080866001600160a01b0316858760405161120b91906116d2565b60006040518083038185875af1925050503d8060008114611248576040519150601f19603f3d011682016040523d82523d6000602084013e61124d565b606091505b509150915061125d828286611268565b979650505050505050565b606083156112775750816104e2565b8251156112875782518084602001fd5b8160405162461bcd60e51b815260040161037f9190611653565b6000602082840312156112b357600080fd5b81356001600160e01b0319811681146104e257600080fd5b6000602082840312156112dd57600080fd5b5035919050565b6001600160a01b03811681146112f957600080fd5b50565b6000806040838503121561130f57600080fd5b823591506020830135611321816112e4565b809150509250929050565b60006020828403121561133e57600080fd5b81356104e2816112e4565b6000806040838503121561135c57600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156113aa576113aa61136b565b604052919050565b600067ffffffffffffffff8211156113cc576113cc61136b565b5060051b60200190565b600082601f8301126113e757600080fd5b813560206113fc6113f7836113b2565b611381565b82815260059290921b8401810191818101908684111561141b57600080fd5b8286015b84811015611436578035835291830191830161141f565b509695505050505050565b6000806040838503121561145457600080fd5b823567ffffffffffffffff8082111561146c57600080fd5b818501915085601f83011261148057600080fd5b813560206114906113f7836113b2565b82815260059290921b840181019181810190898411156114af57600080fd5b948201945b838610156114d65785356114c7816112e4565b825294820194908201906114b4565b965050860135925050808211156114ec57600080fd5b506114f9858286016113d6565b9150509250929050565b60006020828403121561151557600080fd5b5051919050565b60006020828403121561152e57600080fd5b815180151581146104e257600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000821982111561157d5761157d611554565b500190565b60006001820161159457611594611554565b5060010190565b6000828210156115ad576115ad611554565b500390565b60005b838110156115cd5781810151838201526020016115b5565b838111156104c45750506000910152565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516116168160178501602088016115b2565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516116478160288401602088016115b2565b01602801949350505050565b60208152600082518060208401526116728160408501602087016115b2565b601f01601f19169190910160400192915050565b60008160001904831182151516156116a0576116a0611554565b500290565b6000816116b4576116b4611554565b506000190190565b634e487b7160e01b600052603160045260246000fd5b600082516116e48184602087016115b2565b919091019291505056fea2646970667358221220f17350a0154ae587068944f8e9d5aa894ebd7939821c8140ec751be1684379ce64736f6c634300080d0033",
}

// MecMasterABI is the input ABI used to generate the binding from.
// Deprecated: Use MecMasterMetaData.ABI instead.
var MecMasterABI = MecMasterMetaData.ABI

// MecMasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MecMasterMetaData.Bin instead.
var MecMasterBin = MecMasterMetaData.Bin

// DeployMecMaster deploys a new Ethereum contract, binding an instance of MecMaster to it.
func DeployMecMaster(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MecMaster, error) {
	parsed, err := MecMasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MecMasterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MecMaster{MecMasterCaller: MecMasterCaller{contract: contract}, MecMasterTransactor: MecMasterTransactor{contract: contract}, MecMasterFilterer: MecMasterFilterer{contract: contract}}, nil
}

// MecMaster is an auto generated Go binding around an Ethereum contract.
type MecMaster struct {
	MecMasterCaller     // Read-only binding to the contract
	MecMasterTransactor // Write-only binding to the contract
	MecMasterFilterer   // Log filterer for contract events
}

// MecMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MecMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MecMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MecMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MecMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MecMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MecMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MecMasterSession struct {
	Contract     *MecMaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MecMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MecMasterCallerSession struct {
	Contract *MecMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MecMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MecMasterTransactorSession struct {
	Contract     *MecMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MecMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MecMasterRaw struct {
	Contract *MecMaster // Generic contract binding to access the raw methods on
}

// MecMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MecMasterCallerRaw struct {
	Contract *MecMasterCaller // Generic read-only contract binding to access the raw methods on
}

// MecMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MecMasterTransactorRaw struct {
	Contract *MecMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMecMaster creates a new instance of MecMaster, bound to a specific deployed contract.
func NewMecMaster(address common.Address, backend bind.ContractBackend) (*MecMaster, error) {
	contract, err := bindMecMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MecMaster{MecMasterCaller: MecMasterCaller{contract: contract}, MecMasterTransactor: MecMasterTransactor{contract: contract}, MecMasterFilterer: MecMasterFilterer{contract: contract}}, nil
}

// NewMecMasterCaller creates a new read-only instance of MecMaster, bound to a specific deployed contract.
func NewMecMasterCaller(address common.Address, caller bind.ContractCaller) (*MecMasterCaller, error) {
	contract, err := bindMecMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MecMasterCaller{contract: contract}, nil
}

// NewMecMasterTransactor creates a new write-only instance of MecMaster, bound to a specific deployed contract.
func NewMecMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*MecMasterTransactor, error) {
	contract, err := bindMecMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MecMasterTransactor{contract: contract}, nil
}

// NewMecMasterFilterer creates a new log filterer instance of MecMaster, bound to a specific deployed contract.
func NewMecMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*MecMasterFilterer, error) {
	contract, err := bindMecMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MecMasterFilterer{contract: contract}, nil
}

// bindMecMaster binds a generic wrapper to an already deployed contract.
func bindMecMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MecMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MecMaster *MecMasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MecMaster.Contract.MecMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MecMaster *MecMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MecMaster.Contract.MecMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MecMaster *MecMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MecMaster.Contract.MecMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MecMaster *MecMasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MecMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MecMaster *MecMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MecMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MecMaster *MecMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MecMaster.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterSession) ADMINROLE() ([32]byte, error) {
	return _MecMaster.Contract.ADMINROLE(&_MecMaster.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterCallerSession) ADMINROLE() ([32]byte, error) {
	return _MecMaster.Contract.ADMINROLE(&_MecMaster.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MecMaster.Contract.DEFAULTADMINROLE(&_MecMaster.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MecMaster *MecMasterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MecMaster.Contract.DEFAULTADMINROLE(&_MecMaster.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xde242ff4.
//
// Solidity: function allowance() view returns(uint256)
func (_MecMaster *MecMasterCaller) Allowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "allowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xde242ff4.
//
// Solidity: function allowance() view returns(uint256)
func (_MecMaster *MecMasterSession) Allowance() (*big.Int, error) {
	return _MecMaster.Contract.Allowance(&_MecMaster.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xde242ff4.
//
// Solidity: function allowance() view returns(uint256)
func (_MecMaster *MecMasterCallerSession) Allowance() (*big.Int, error) {
	return _MecMaster.Contract.Allowance(&_MecMaster.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MecMaster *MecMasterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MecMaster *MecMasterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MecMaster.Contract.GetRoleAdmin(&_MecMaster.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MecMaster *MecMasterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MecMaster.Contract.GetRoleAdmin(&_MecMaster.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MecMaster *MecMasterCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MecMaster *MecMasterSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MecMaster.Contract.GetRoleMember(&_MecMaster.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MecMaster *MecMasterCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MecMaster.Contract.GetRoleMember(&_MecMaster.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MecMaster *MecMasterCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MecMaster *MecMasterSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MecMaster.Contract.GetRoleMemberCount(&_MecMaster.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MecMaster *MecMasterCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MecMaster.Contract.GetRoleMemberCount(&_MecMaster.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MecMaster *MecMasterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MecMaster *MecMasterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MecMaster.Contract.HasRole(&_MecMaster.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MecMaster *MecMasterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MecMaster.Contract.HasRole(&_MecMaster.CallOpts, role, account)
}

// MecToken is a free data retrieval call binding the contract method 0x34b412b9.
//
// Solidity: function mecToken() view returns(address)
func (_MecMaster *MecMasterCaller) MecToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "mecToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MecToken is a free data retrieval call binding the contract method 0x34b412b9.
//
// Solidity: function mecToken() view returns(address)
func (_MecMaster *MecMasterSession) MecToken() (common.Address, error) {
	return _MecMaster.Contract.MecToken(&_MecMaster.CallOpts)
}

// MecToken is a free data retrieval call binding the contract method 0x34b412b9.
//
// Solidity: function mecToken() view returns(address)
func (_MecMaster *MecMasterCallerSession) MecToken() (common.Address, error) {
	return _MecMaster.Contract.MecToken(&_MecMaster.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MecMaster *MecMasterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MecMaster.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MecMaster *MecMasterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MecMaster.Contract.SupportsInterface(&_MecMaster.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MecMaster *MecMasterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MecMaster.Contract.SupportsInterface(&_MecMaster.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) returns()
func (_MecMaster *MecMasterTransactor) Deposit(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "deposit", amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) returns()
func (_MecMaster *MecMasterSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.Deposit(&_MecMaster.TransactOpts, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) returns()
func (_MecMaster *MecMasterTransactorSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.Deposit(&_MecMaster.TransactOpts, amount_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.GrantRole(&_MecMaster.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.GrantRole(&_MecMaster.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_MecMaster *MecMasterTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "initialize", token_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_MecMaster *MecMasterSession) Initialize(token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.Initialize(&_MecMaster.TransactOpts, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_MecMaster *MecMasterTransactorSession) Initialize(token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.Initialize(&_MecMaster.TransactOpts, token_)
}

// MultiSend is a paid mutator transaction binding the contract method 0xbb4c9f0b.
//
// Solidity: function multiSend(address[] recipients_, uint256[] amount_) returns()
func (_MecMaster *MecMasterTransactor) MultiSend(opts *bind.TransactOpts, recipients_ []common.Address, amount_ []*big.Int) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "multiSend", recipients_, amount_)
}

// MultiSend is a paid mutator transaction binding the contract method 0xbb4c9f0b.
//
// Solidity: function multiSend(address[] recipients_, uint256[] amount_) returns()
func (_MecMaster *MecMasterSession) MultiSend(recipients_ []common.Address, amount_ []*big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.MultiSend(&_MecMaster.TransactOpts, recipients_, amount_)
}

// MultiSend is a paid mutator transaction binding the contract method 0xbb4c9f0b.
//
// Solidity: function multiSend(address[] recipients_, uint256[] amount_) returns()
func (_MecMaster *MecMasterTransactorSession) MultiSend(recipients_ []common.Address, amount_ []*big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.MultiSend(&_MecMaster.TransactOpts, recipients_, amount_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_MecMaster *MecMasterTransactor) PullFunds(opts *bind.TransactOpts, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "pullFunds", tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_MecMaster *MecMasterSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.PullFunds(&_MecMaster.TransactOpts, tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_MecMaster *MecMasterTransactorSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.PullFunds(&_MecMaster.TransactOpts, tokenAddress_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.RenounceRole(&_MecMaster.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.RenounceRole(&_MecMaster.TransactOpts, role, account)
}

// ResetAllowance is a paid mutator transaction binding the contract method 0xeee4f11b.
//
// Solidity: function resetAllowance(uint256 amount_) returns()
func (_MecMaster *MecMasterTransactor) ResetAllowance(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "resetAllowance", amount_)
}

// ResetAllowance is a paid mutator transaction binding the contract method 0xeee4f11b.
//
// Solidity: function resetAllowance(uint256 amount_) returns()
func (_MecMaster *MecMasterSession) ResetAllowance(amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.ResetAllowance(&_MecMaster.TransactOpts, amount_)
}

// ResetAllowance is a paid mutator transaction binding the contract method 0xeee4f11b.
//
// Solidity: function resetAllowance(uint256 amount_) returns()
func (_MecMaster *MecMasterTransactorSession) ResetAllowance(amount_ *big.Int) (*types.Transaction, error) {
	return _MecMaster.Contract.ResetAllowance(&_MecMaster.TransactOpts, amount_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.RevokeRole(&_MecMaster.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MecMaster *MecMasterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.RevokeRole(&_MecMaster.TransactOpts, role, account)
}

// SetMecAddress is a paid mutator transaction binding the contract method 0xc653fe23.
//
// Solidity: function setMecAddress(address token_) returns()
func (_MecMaster *MecMasterTransactor) SetMecAddress(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.contract.Transact(opts, "setMecAddress", token_)
}

// SetMecAddress is a paid mutator transaction binding the contract method 0xc653fe23.
//
// Solidity: function setMecAddress(address token_) returns()
func (_MecMaster *MecMasterSession) SetMecAddress(token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.SetMecAddress(&_MecMaster.TransactOpts, token_)
}

// SetMecAddress is a paid mutator transaction binding the contract method 0xc653fe23.
//
// Solidity: function setMecAddress(address token_) returns()
func (_MecMaster *MecMasterTransactorSession) SetMecAddress(token_ common.Address) (*types.Transaction, error) {
	return _MecMaster.Contract.SetMecAddress(&_MecMaster.TransactOpts, token_)
}

// MecMasterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the MecMaster contract.
type MecMasterDepositedIterator struct {
	Event *MecMasterDeposited // Event containing the contract specifics and raw log

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
func (it *MecMasterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MecMasterDeposited)
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
		it.Event = new(MecMasterDeposited)
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
func (it *MecMasterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MecMasterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MecMasterDeposited represents a Deposited event raised by the MecMaster contract.
type MecMasterDeposited struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) FilterDeposited(opts *bind.FilterOpts, who []common.Address, amount []*big.Int) (*MecMasterDepositedIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MecMaster.contract.FilterLogs(opts, "Deposited", whoRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MecMasterDepositedIterator{contract: _MecMaster.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MecMasterDeposited, who []common.Address, amount []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MecMaster.contract.WatchLogs(opts, "Deposited", whoRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MecMasterDeposited)
				if err := _MecMaster.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) ParseDeposited(log types.Log) (*MecMasterDeposited, error) {
	event := new(MecMasterDeposited)
	if err := _MecMaster.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MecMasterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MecMaster contract.
type MecMasterRoleAdminChangedIterator struct {
	Event *MecMasterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MecMasterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MecMasterRoleAdminChanged)
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
		it.Event = new(MecMasterRoleAdminChanged)
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
func (it *MecMasterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MecMasterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MecMasterRoleAdminChanged represents a RoleAdminChanged event raised by the MecMaster contract.
type MecMasterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MecMaster *MecMasterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MecMasterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MecMaster.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MecMasterRoleAdminChangedIterator{contract: _MecMaster.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MecMaster *MecMasterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MecMasterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MecMaster.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MecMasterRoleAdminChanged)
				if err := _MecMaster.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MecMaster *MecMasterFilterer) ParseRoleAdminChanged(log types.Log) (*MecMasterRoleAdminChanged, error) {
	event := new(MecMasterRoleAdminChanged)
	if err := _MecMaster.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MecMasterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MecMaster contract.
type MecMasterRoleGrantedIterator struct {
	Event *MecMasterRoleGranted // Event containing the contract specifics and raw log

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
func (it *MecMasterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MecMasterRoleGranted)
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
		it.Event = new(MecMasterRoleGranted)
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
func (it *MecMasterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MecMasterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MecMasterRoleGranted represents a RoleGranted event raised by the MecMaster contract.
type MecMasterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MecMaster *MecMasterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MecMasterRoleGrantedIterator, error) {

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

	logs, sub, err := _MecMaster.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MecMasterRoleGrantedIterator{contract: _MecMaster.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MecMaster *MecMasterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MecMasterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MecMaster.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MecMasterRoleGranted)
				if err := _MecMaster.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MecMaster *MecMasterFilterer) ParseRoleGranted(log types.Log) (*MecMasterRoleGranted, error) {
	event := new(MecMasterRoleGranted)
	if err := _MecMaster.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MecMasterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MecMaster contract.
type MecMasterRoleRevokedIterator struct {
	Event *MecMasterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MecMasterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MecMasterRoleRevoked)
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
		it.Event = new(MecMasterRoleRevoked)
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
func (it *MecMasterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MecMasterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MecMasterRoleRevoked represents a RoleRevoked event raised by the MecMaster contract.
type MecMasterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MecMaster *MecMasterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MecMasterRoleRevokedIterator, error) {

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

	logs, sub, err := _MecMaster.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MecMasterRoleRevokedIterator{contract: _MecMaster.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MecMaster *MecMasterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MecMasterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MecMaster.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MecMasterRoleRevoked)
				if err := _MecMaster.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MecMaster *MecMasterFilterer) ParseRoleRevoked(log types.Log) (*MecMasterRoleRevoked, error) {
	event := new(MecMasterRoleRevoked)
	if err := _MecMaster.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MecMasterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MecMaster contract.
type MecMasterWithdrawIterator struct {
	Event *MecMasterWithdraw // Event containing the contract specifics and raw log

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
func (it *MecMasterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MecMasterWithdraw)
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
		it.Event = new(MecMasterWithdraw)
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
func (it *MecMasterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MecMasterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MecMasterWithdraw represents a Withdraw event raised by the MecMaster contract.
type MecMasterWithdraw struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) FilterWithdraw(opts *bind.FilterOpts, who []common.Address, amount []*big.Int) (*MecMasterWithdrawIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MecMaster.contract.FilterLogs(opts, "Withdraw", whoRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MecMasterWithdrawIterator{contract: _MecMaster.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MecMasterWithdraw, who []common.Address, amount []*big.Int) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MecMaster.contract.WatchLogs(opts, "Withdraw", whoRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MecMasterWithdraw)
				if err := _MecMaster.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed who, uint256 indexed amount)
func (_MecMaster *MecMasterFilterer) ParseWithdraw(log types.Log) (*MecMasterWithdraw, error) {
	event := new(MecMasterWithdraw)
	if err := _MecMaster.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
