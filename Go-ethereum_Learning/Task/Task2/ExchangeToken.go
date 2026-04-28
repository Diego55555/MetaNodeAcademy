// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ExchangeTokenMetaData contains all meta data concerning the ExchangeToken contract.
var ExchangeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeETHToETK\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"exchangeETKToETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526103e8600555348015610015575f5ffd5b50604051611e67380380611e67833981810160405281019061003791906104f6565b6040518060400160405280600d81526020017f45786368616e6765546f6b656e000000000000000000000000000000000000008152506040518060400160405280600381526020017f45544b000000000000000000000000000000000000000000000000000000000081525081600390816100b29190610766565b5080600490816100c29190610766565b50505061010a6040518060400160405280600e81526020017f696e697469616c537570706c793a0000000000000000000000000000000000008152508261012060201b60201c565b61011a30826101c260201b60201c565b50610a0b565b6101be82826040516024016101369291906108aa565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061024760201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610232575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016102299190610917565b60405180910390fd5b6102435f838361026b60201b60201c565b5050565b6102688161026361048460201b610768176104a360201b60201c565b60201c565b50565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036102bb578060025f8282546102af919061095d565b92505081905550610389565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610344578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161033b93929190610990565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036103d0578060025f828254039250508190555061041a565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161047791906109c5565b60405180910390a3505050565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b6104b560201b610dcc17819050919050565b6104bd6109de565b565b5f5ffd5b5f819050919050565b6104d5816104c3565b81146104df575f5ffd5b50565b5f815190506104f0816104cc565b92915050565b5f6020828403121561050b5761050a6104bf565b5b5f610518848285016104e2565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061059c57607f821691505b6020821081036105af576105ae610558565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106117fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105d6565b61061b86836105d6565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61065661065161064c846104c3565b610633565b6104c3565b9050919050565b5f819050919050565b61066f8361063c565b61068361067b8261065d565b8484546105e2565b825550505050565b5f5f905090565b61069a61068b565b6106a5818484610666565b505050565b5f5b828110156106cb576106c05f828401610692565b6001810190506106ac565b505050565b601f82111561071e578282111561071d576106ea816105b5565b6106f3836105c7565b6106fc856105c7565b6020861015610709575f90505b808301610718828403826106aa565b505050505b5b505050565b5f82821c905092915050565b5f61073e5f1984600802610723565b1980831691505092915050565b5f610756838361072f565b9150826002028217905092915050565b61076f82610521565b67ffffffffffffffff8111156107885761078761052b565b5b6107928254610585565b61079d8282856106d0565b5f60209050601f8311600181146107ce575f84156107bc578287015190505b6107c6858261074b565b86555061082d565b601f1984166107dc866105b5565b5f5b82811015610803578489015182556001820191506020850194506020810190506107de565b86831015610820578489015161081c601f89168261072f565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61086d82610521565b6108778185610835565b9350610887818560208601610845565b61089081610853565b840191505092915050565b6108a4816104c3565b82525050565b5f6040820190508181035f8301526108c28185610863565b90506108d1602083018461089b565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610901826108d8565b9050919050565b610911816108f7565b82525050565b5f60208201905061092a5f830184610908565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610967826104c3565b9150610972836104c3565b925082820190508082111561098a57610989610930565b5b92915050565b5f6060820190506109a35f830186610908565b6109b0602083018561089b565b6109bd604083018461089b565b949350505050565b5f6020820190506109d85f83018461089b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffd5b61144f80610a185f395ff3fe6080604052600436106100a6575f3560e01c80635bd56523116100635780635bd56523146101ca57806370a08231146101d457806395d89b4114610210578063a9059cbb1461023a578063da4fe10f14610276578063dd62ed3e1461029e576100a6565b806306fdde03146100aa578063095ea7b3146100d457806318160ddd1461011057806323b872dd1461013a578063313ce567146101765780633ba0b9a9146101a0575b5f5ffd5b3480156100b5575f5ffd5b506100be6102da565b6040516100cb9190610e46565b60405180910390f35b3480156100df575f5ffd5b506100fa60048036038101906100f59190610ef7565b61036a565b6040516101079190610f4f565b60405180910390f35b34801561011b575f5ffd5b5061012461038c565b6040516101319190610f77565b60405180910390f35b348015610145575f5ffd5b50610160600480360381019061015b9190610f90565b610395565b60405161016d9190610f4f565b60405180910390f35b348015610181575f5ffd5b5061018a6103c3565b6040516101979190610ffb565b60405180910390f35b3480156101ab575f5ffd5b506101b46103cb565b6040516101c19190610f77565b60405180910390f35b6101d26103d1565b005b3480156101df575f5ffd5b506101fa60048036038101906101f59190611014565b610500565b6040516102079190610f77565b60405180910390f35b34801561021b575f5ffd5b50610224610545565b6040516102319190610e46565b60405180910390f35b348015610245575f5ffd5b50610260600480360381019061025b9190610ef7565b6105d5565b60405161026d9190610f4f565b60405180910390f35b348015610281575f5ffd5b5061029c6004803603810190610297919061103f565b6105f7565b005b3480156102a9575f5ffd5b506102c460048036038101906102bf919061106a565b6106e6565b6040516102d19190610f77565b60405180910390f35b6060600380546102e9906110d5565b80601f0160208091040260200160405190810160405280929190818152602001828054610315906110d5565b80156103605780601f1061033757610100808354040283529160200191610360565b820191905f5260205f20905b81548152906001019060200180831161034357829003601f168201915b5050505050905090565b5f5f610374610787565b905061038181858561078e565b600191505092915050565b5f600254905090565b5f5f61039f610787565b90506103ac8582856107a0565b6103b7858585610833565b60019150509392505050565b5f6012905090565b60055481565b5f3411610413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040a9061114f565b60405180910390fd5b5f61041d30610500565b90505f6005543461042e919061119a565b905061046f6040518060400160405280600c81526020017f746f6b656e52656d61696e3a000000000000000000000000000000000000000081525083610923565b6104ae6040518060400160405280600c81526020017f746f6b656e416d6f756e743a000000000000000000000000000000000000000081525082610923565b808210156104f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e890611225565b60405180910390fd5b6104fc303383610833565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b606060048054610554906110d5565b80601f0160208091040260200160405190810160405280929190818152602001828054610580906110d5565b80156105cb5780601f106105a2576101008083540402835291602001916105cb565b820191905f5260205f20905b8154815290600101906020018083116105ae57829003601f168201915b5050505050905090565b5f5f6105df610787565b90506105ec818585610833565b600191505092915050565b6103e881101561063c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610633906112b3565b60405180910390fd5b8061064633610500565b1015610687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067e90611225565b60405180910390fd5b610692333083610833565b3373ffffffffffffffffffffffffffffffffffffffff166108fc600554836106ba91906112fe565b90811502906040515f60405180830381858888f193505050501580156106e2573d5f5f3e3d5ffd5b5050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b5f33905090565b61079b83838360016109bf565b505050565b5f6107ab84846106e6565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561082d578181101561081e578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016108159392919061133d565b60405180910390fd5b61082c84848484035f6109bf565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108a3575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161089a9190611372565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610913575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161090a9190611372565b60405180910390fd5b61091e838383610b8e565b505050565b6109bb828260405160240161093992919061138b565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610da7565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a2f575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610a269190611372565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a9f575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610a969190611372565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610b88578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610b7f9190610f77565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bde578060025f828254610bd291906113b9565b92505081905550610cac565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610c67578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610c5e9392919061133d565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cf3578060025f8282540392505081905550610d3d565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d9a9190610f77565b60405180910390a3505050565b610dbe81610db6610768610dc1565b63ffffffff16565b50565b610dcc819050919050565b610dd46113ec565b565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e1882610dd6565b610e228185610de0565b9350610e32818560208601610df0565b610e3b81610dfe565b840191505092915050565b5f6020820190508181035f830152610e5e8184610e0e565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e9382610e6a565b9050919050565b610ea381610e89565b8114610ead575f5ffd5b50565b5f81359050610ebe81610e9a565b92915050565b5f819050919050565b610ed681610ec4565b8114610ee0575f5ffd5b50565b5f81359050610ef181610ecd565b92915050565b5f5f60408385031215610f0d57610f0c610e66565b5b5f610f1a85828601610eb0565b9250506020610f2b85828601610ee3565b9150509250929050565b5f8115159050919050565b610f4981610f35565b82525050565b5f602082019050610f625f830184610f40565b92915050565b610f7181610ec4565b82525050565b5f602082019050610f8a5f830184610f68565b92915050565b5f5f5f60608486031215610fa757610fa6610e66565b5b5f610fb486828701610eb0565b9350506020610fc586828701610eb0565b9250506040610fd686828701610ee3565b9150509250925092565b5f60ff82169050919050565b610ff581610fe0565b82525050565b5f60208201905061100e5f830184610fec565b92915050565b5f6020828403121561102957611028610e66565b5b5f61103684828501610eb0565b91505092915050565b5f6020828403121561105457611053610e66565b5b5f61106184828501610ee3565b91505092915050565b5f5f604083850312156110805761107f610e66565b5b5f61108d85828601610eb0565b925050602061109e85828601610eb0565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806110ec57607f821691505b6020821081036110ff576110fe6110a8565b5b50919050565b7f56616c7565206d7573742062652067726561746572207468616e2030000000005f82015250565b5f611139601c83610de0565b915061114482611105565b602082019050919050565b5f6020820190508181035f8301526111668161112d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111a482610ec4565b91506111af83610ec4565b92508282026111bd81610ec4565b915082820484148315176111d4576111d361116d565b5b5092915050565b7f496e73756666696369656e742045544b000000000000000000000000000000005f82015250565b5f61120f601083610de0565b915061121a826111db565b602082019050919050565b5f6020820190508181035f83015261123c81611203565b9050919050565b7f416d6f756e74206d7573742062652067726561746572207468616e206f7220655f8201527f7175616c20746f20313030300000000000000000000000000000000000000000602082015250565b5f61129d602c83610de0565b91506112a882611243565b604082019050919050565b5f6020820190508181035f8301526112ca81611291565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61130882610ec4565b915061131383610ec4565b925082611323576113226112d1565b5b828204905092915050565b61133781610e89565b82525050565b5f6060820190506113505f83018661132e565b61135d6020830185610f68565b61136a6040830184610f68565b949350505050565b5f6020820190506113855f83018461132e565b92915050565b5f6040820190508181035f8301526113a38185610e0e565b90506113b26020830184610f68565b9392505050565b5f6113c382610ec4565b91506113ce83610ec4565b92508282019050808211156113e6576113e561116d565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220ae7ef47e1733c51156a3fbef1d9641209eda5636ecae3ee6911aeebae90bc8dc64736f6c63430008220033",
}

// ExchangeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeTokenMetaData.ABI instead.
var ExchangeTokenABI = ExchangeTokenMetaData.ABI

// ExchangeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExchangeTokenMetaData.Bin instead.
var ExchangeTokenBin = ExchangeTokenMetaData.Bin

// DeployExchangeToken deploys a new Ethereum contract, binding an instance of ExchangeToken to it.
func DeployExchangeToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *ExchangeToken, error) {
	parsed, err := ExchangeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExchangeTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExchangeToken{ExchangeTokenCaller: ExchangeTokenCaller{contract: contract}, ExchangeTokenTransactor: ExchangeTokenTransactor{contract: contract}, ExchangeTokenFilterer: ExchangeTokenFilterer{contract: contract}}, nil
}

// ExchangeToken is an auto generated Go binding around an Ethereum contract.
type ExchangeToken struct {
	ExchangeTokenCaller     // Read-only binding to the contract
	ExchangeTokenTransactor // Write-only binding to the contract
	ExchangeTokenFilterer   // Log filterer for contract events
}

// ExchangeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeTokenSession struct {
	Contract     *ExchangeToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeTokenCallerSession struct {
	Contract *ExchangeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ExchangeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTokenTransactorSession struct {
	Contract     *ExchangeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExchangeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeTokenRaw struct {
	Contract *ExchangeToken // Generic contract binding to access the raw methods on
}

// ExchangeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeTokenCallerRaw struct {
	Contract *ExchangeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTokenTransactorRaw struct {
	Contract *ExchangeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeToken creates a new instance of ExchangeToken, bound to a specific deployed contract.
func NewExchangeToken(address common.Address, backend bind.ContractBackend) (*ExchangeToken, error) {
	contract, err := bindExchangeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeToken{ExchangeTokenCaller: ExchangeTokenCaller{contract: contract}, ExchangeTokenTransactor: ExchangeTokenTransactor{contract: contract}, ExchangeTokenFilterer: ExchangeTokenFilterer{contract: contract}}, nil
}

// NewExchangeTokenCaller creates a new read-only instance of ExchangeToken, bound to a specific deployed contract.
func NewExchangeTokenCaller(address common.Address, caller bind.ContractCaller) (*ExchangeTokenCaller, error) {
	contract, err := bindExchangeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenCaller{contract: contract}, nil
}

// NewExchangeTokenTransactor creates a new write-only instance of ExchangeToken, bound to a specific deployed contract.
func NewExchangeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTokenTransactor, error) {
	contract, err := bindExchangeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenTransactor{contract: contract}, nil
}

// NewExchangeTokenFilterer creates a new log filterer instance of ExchangeToken, bound to a specific deployed contract.
func NewExchangeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeTokenFilterer, error) {
	contract, err := bindExchangeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenFilterer{contract: contract}, nil
}

// bindExchangeToken binds a generic wrapper to an already deployed contract.
func bindExchangeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeToken *ExchangeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeToken.Contract.ExchangeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeToken *ExchangeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeToken *ExchangeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeToken *ExchangeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeToken *ExchangeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeToken *ExchangeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExchangeToken *ExchangeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExchangeToken *ExchangeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExchangeToken.Contract.Allowance(&_ExchangeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExchangeToken *ExchangeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExchangeToken.Contract.Allowance(&_ExchangeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExchangeToken *ExchangeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExchangeToken *ExchangeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExchangeToken.Contract.BalanceOf(&_ExchangeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExchangeToken *ExchangeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExchangeToken.Contract.BalanceOf(&_ExchangeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExchangeToken *ExchangeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExchangeToken *ExchangeTokenSession) Decimals() (uint8, error) {
	return _ExchangeToken.Contract.Decimals(&_ExchangeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExchangeToken *ExchangeTokenCallerSession) Decimals() (uint8, error) {
	return _ExchangeToken.Contract.Decimals(&_ExchangeToken.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ExchangeToken *ExchangeTokenCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ExchangeToken *ExchangeTokenSession) ExchangeRate() (*big.Int, error) {
	return _ExchangeToken.Contract.ExchangeRate(&_ExchangeToken.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ExchangeToken *ExchangeTokenCallerSession) ExchangeRate() (*big.Int, error) {
	return _ExchangeToken.Contract.ExchangeRate(&_ExchangeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExchangeToken *ExchangeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExchangeToken *ExchangeTokenSession) Name() (string, error) {
	return _ExchangeToken.Contract.Name(&_ExchangeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExchangeToken *ExchangeTokenCallerSession) Name() (string, error) {
	return _ExchangeToken.Contract.Name(&_ExchangeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExchangeToken *ExchangeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExchangeToken *ExchangeTokenSession) Symbol() (string, error) {
	return _ExchangeToken.Contract.Symbol(&_ExchangeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExchangeToken *ExchangeTokenCallerSession) Symbol() (string, error) {
	return _ExchangeToken.Contract.Symbol(&_ExchangeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExchangeToken *ExchangeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExchangeToken *ExchangeTokenSession) TotalSupply() (*big.Int, error) {
	return _ExchangeToken.Contract.TotalSupply(&_ExchangeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExchangeToken *ExchangeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ExchangeToken.Contract.TotalSupply(&_ExchangeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.Approve(&_ExchangeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.Approve(&_ExchangeToken.TransactOpts, spender, value)
}

// ExchangeETHToETK is a paid mutator transaction binding the contract method 0x5bd56523.
//
// Solidity: function exchangeETHToETK() payable returns()
func (_ExchangeToken *ExchangeTokenTransactor) ExchangeETHToETK(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeToken.contract.Transact(opts, "exchangeETHToETK")
}

// ExchangeETHToETK is a paid mutator transaction binding the contract method 0x5bd56523.
//
// Solidity: function exchangeETHToETK() payable returns()
func (_ExchangeToken *ExchangeTokenSession) ExchangeETHToETK() (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeETHToETK(&_ExchangeToken.TransactOpts)
}

// ExchangeETHToETK is a paid mutator transaction binding the contract method 0x5bd56523.
//
// Solidity: function exchangeETHToETK() payable returns()
func (_ExchangeToken *ExchangeTokenTransactorSession) ExchangeETHToETK() (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeETHToETK(&_ExchangeToken.TransactOpts)
}

// ExchangeETKToETH is a paid mutator transaction binding the contract method 0xda4fe10f.
//
// Solidity: function exchangeETKToETH(uint256 tokenAmount) returns()
func (_ExchangeToken *ExchangeTokenTransactor) ExchangeETKToETH(opts *bind.TransactOpts, tokenAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.contract.Transact(opts, "exchangeETKToETH", tokenAmount)
}

// ExchangeETKToETH is a paid mutator transaction binding the contract method 0xda4fe10f.
//
// Solidity: function exchangeETKToETH(uint256 tokenAmount) returns()
func (_ExchangeToken *ExchangeTokenSession) ExchangeETKToETH(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeETKToETH(&_ExchangeToken.TransactOpts, tokenAmount)
}

// ExchangeETKToETH is a paid mutator transaction binding the contract method 0xda4fe10f.
//
// Solidity: function exchangeETKToETH(uint256 tokenAmount) returns()
func (_ExchangeToken *ExchangeTokenTransactorSession) ExchangeETKToETH(tokenAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.ExchangeETKToETH(&_ExchangeToken.TransactOpts, tokenAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.Transfer(&_ExchangeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.Transfer(&_ExchangeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.TransferFrom(&_ExchangeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExchangeToken *ExchangeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExchangeToken.Contract.TransferFrom(&_ExchangeToken.TransactOpts, from, to, value)
}

// ExchangeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExchangeToken contract.
type ExchangeTokenApprovalIterator struct {
	Event *ExchangeTokenApproval // Event containing the contract specifics and raw log

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
func (it *ExchangeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTokenApproval)
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
		it.Event = new(ExchangeTokenApproval)
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
func (it *ExchangeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTokenApproval represents a Approval event raised by the ExchangeToken contract.
type ExchangeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExchangeToken *ExchangeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExchangeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExchangeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenApprovalIterator{contract: _ExchangeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExchangeToken *ExchangeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExchangeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExchangeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTokenApproval)
				if err := _ExchangeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ExchangeToken *ExchangeTokenFilterer) ParseApproval(log types.Log) (*ExchangeTokenApproval, error) {
	event := new(ExchangeTokenApproval)
	if err := _ExchangeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExchangeToken contract.
type ExchangeTokenTransferIterator struct {
	Event *ExchangeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTokenTransfer)
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
		it.Event = new(ExchangeTokenTransfer)
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
func (it *ExchangeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTokenTransfer represents a Transfer event raised by the ExchangeToken contract.
type ExchangeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExchangeToken *ExchangeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExchangeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExchangeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenTransferIterator{contract: _ExchangeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExchangeToken *ExchangeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExchangeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTokenTransfer)
				if err := _ExchangeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ExchangeToken *ExchangeTokenFilterer) ParseTransfer(log types.Log) (*ExchangeTokenTransfer, error) {
	event := new(ExchangeTokenTransfer)
	if err := _ExchangeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
