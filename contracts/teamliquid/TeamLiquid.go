// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teamliquid

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

// ReservoirV601AmountCheckInfo is an auto generated low-level Go binding around an user-defined struct.
type ReservoirV601AmountCheckInfo struct {
	Target    common.Address
	Data      []byte
	Threshold *big.Int
}

// ReservoirV601ExecutionInfo is an auto generated low-level Go binding around an user-defined struct.
type ReservoirV601ExecutionInfo struct {
	Module common.Address
	Data   []byte
	Value  *big.Int
}

// TeamliquidMetaData contains all meta data concerning the Teamliquid contract.
var TeamliquidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UnsuccessfulExecution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsuccessfulPayment\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structReservoirV6_0_1.ExecutionInfo[]\",\"name\":\"executionInfos\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structReservoirV6_0_1.ExecutionInfo[]\",\"name\":\"executionInfos\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"internalType\":\"structReservoirV6_0_1.AmountCheckInfo\",\"name\":\"amountCheckInfo\",\"type\":\"tuple\"}],\"name\":\"executeWithAmountCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60406080815260048036101561001f575b5050361561001d57600080fd5b005b600091823560e01c806304871891146100dd5763760f2a0b146100425750610010565b60203660031901126100d957813567ffffffffffffffff81116100d55761006c90369084016101f4565b6100746102c3565b845b8181106100b957505050824780610090575b506001815580f35b81808092335af161009f61022a565b50156100ac578281610088565b5163d2dcf4f360e01b8152fd5b806100cf6100ca600193858761028b565b610360565b01610076565b8380fd5b8280fd5b5060031981813601126100d55767ffffffffffffffff83358181116101f05761010990369086016101f4565b926024359283116101ec5760608387019184360301126101ec5760449061012e6102c3565b61014561013a82610319565b91602486019061032d565b929094013591885b86811061016b575b5050505050505082478061009057506001815580f35b823b156101dc578980895184898237808581018381520390865afa61018e61022a565b90156101cc57602080828051810103126101c8578591015110156101c357806101bd6100ca6001938a8961028b565b0161014d565b610155565b8b80fd5b8851635589343b60e11b81528a90fd5b8751635589343b60e11b81528990fd5b8680fd5b8580fd5b9181601f840112156102255782359167ffffffffffffffff8311610225576020808501948460051b01011161022557565b600080fd5b3d156102865767ffffffffffffffff903d8281116102705760405192601f8201601f19908116603f01168401908111848210176102705760405282523d6000602084013e565b634e487b7160e01b600052604160045260246000fd5b606090565b91908110156102ad5760051b81013590605e1981360301821215610225570190565b634e487b7160e01b600052603260045260246000fd5b6002600054146102d4576002600055565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b356001600160a01b03811681036102255790565b903590601e1981360301821215610225570180359067ffffffffffffffff82116102255760200191813603831361022557565b61036981610319565b803b156103ab578160406000939261038560208695018461032d565b92908382519485928337810186815203930135905af16103a361022a565b50156103ab57565b604051635589343b60e11b8152600490fdfea2646970667358221220ce1d7e58645bc5b9b56d7d97491dc6b1dac67f3bc9bc72e2cd48e0f078f3392764736f6c63430008110033",
}

// TeamliquidABI is the input ABI used to generate the binding from.
// Deprecated: Use TeamliquidMetaData.ABI instead.
var TeamliquidABI = TeamliquidMetaData.ABI

// TeamliquidBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeamliquidMetaData.Bin instead.
var TeamliquidBin = TeamliquidMetaData.Bin

// DeployTeamliquid deploys a new Ethereum contract, binding an instance of Teamliquid to it.
func DeployTeamliquid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Teamliquid, error) {
	parsed, err := TeamliquidMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeamliquidBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Teamliquid{TeamliquidCaller: TeamliquidCaller{contract: contract}, TeamliquidTransactor: TeamliquidTransactor{contract: contract}, TeamliquidFilterer: TeamliquidFilterer{contract: contract}}, nil
}

// Teamliquid is an auto generated Go binding around an Ethereum contract.
type Teamliquid struct {
	TeamliquidCaller     // Read-only binding to the contract
	TeamliquidTransactor // Write-only binding to the contract
	TeamliquidFilterer   // Log filterer for contract events
}

// TeamliquidCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamliquidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamliquidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamliquidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamliquidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamliquidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamliquidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamliquidSession struct {
	Contract     *Teamliquid       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeamliquidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamliquidCallerSession struct {
	Contract *TeamliquidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TeamliquidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamliquidTransactorSession struct {
	Contract     *TeamliquidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeamliquidRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamliquidRaw struct {
	Contract *Teamliquid // Generic contract binding to access the raw methods on
}

// TeamliquidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamliquidCallerRaw struct {
	Contract *TeamliquidCaller // Generic read-only contract binding to access the raw methods on
}

// TeamliquidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamliquidTransactorRaw struct {
	Contract *TeamliquidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamliquid creates a new instance of Teamliquid, bound to a specific deployed contract.
func NewTeamliquid(address common.Address, backend bind.ContractBackend) (*Teamliquid, error) {
	contract, err := bindTeamliquid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Teamliquid{TeamliquidCaller: TeamliquidCaller{contract: contract}, TeamliquidTransactor: TeamliquidTransactor{contract: contract}, TeamliquidFilterer: TeamliquidFilterer{contract: contract}}, nil
}

// NewTeamliquidCaller creates a new read-only instance of Teamliquid, bound to a specific deployed contract.
func NewTeamliquidCaller(address common.Address, caller bind.ContractCaller) (*TeamliquidCaller, error) {
	contract, err := bindTeamliquid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamliquidCaller{contract: contract}, nil
}

// NewTeamliquidTransactor creates a new write-only instance of Teamliquid, bound to a specific deployed contract.
func NewTeamliquidTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamliquidTransactor, error) {
	contract, err := bindTeamliquid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamliquidTransactor{contract: contract}, nil
}

// NewTeamliquidFilterer creates a new log filterer instance of Teamliquid, bound to a specific deployed contract.
func NewTeamliquidFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamliquidFilterer, error) {
	contract, err := bindTeamliquid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamliquidFilterer{contract: contract}, nil
}

// bindTeamliquid binds a generic wrapper to an already deployed contract.
func bindTeamliquid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeamliquidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teamliquid *TeamliquidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teamliquid.Contract.TeamliquidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teamliquid *TeamliquidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teamliquid.Contract.TeamliquidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teamliquid *TeamliquidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teamliquid.Contract.TeamliquidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teamliquid *TeamliquidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teamliquid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teamliquid *TeamliquidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teamliquid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teamliquid *TeamliquidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teamliquid.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x760f2a0b.
//
// Solidity: function execute((address,bytes,uint256)[] executionInfos) payable returns()
func (_Teamliquid *TeamliquidTransactor) Execute(opts *bind.TransactOpts, executionInfos []ReservoirV601ExecutionInfo) (*types.Transaction, error) {
	return _Teamliquid.contract.Transact(opts, "execute", executionInfos)
}

// Execute is a paid mutator transaction binding the contract method 0x760f2a0b.
//
// Solidity: function execute((address,bytes,uint256)[] executionInfos) payable returns()
func (_Teamliquid *TeamliquidSession) Execute(executionInfos []ReservoirV601ExecutionInfo) (*types.Transaction, error) {
	return _Teamliquid.Contract.Execute(&_Teamliquid.TransactOpts, executionInfos)
}

// Execute is a paid mutator transaction binding the contract method 0x760f2a0b.
//
// Solidity: function execute((address,bytes,uint256)[] executionInfos) payable returns()
func (_Teamliquid *TeamliquidTransactorSession) Execute(executionInfos []ReservoirV601ExecutionInfo) (*types.Transaction, error) {
	return _Teamliquid.Contract.Execute(&_Teamliquid.TransactOpts, executionInfos)
}

// ExecuteWithAmountCheck is a paid mutator transaction binding the contract method 0x04871891.
//
// Solidity: function executeWithAmountCheck((address,bytes,uint256)[] executionInfos, (address,bytes,uint256) amountCheckInfo) payable returns()
func (_Teamliquid *TeamliquidTransactor) ExecuteWithAmountCheck(opts *bind.TransactOpts, executionInfos []ReservoirV601ExecutionInfo, amountCheckInfo ReservoirV601AmountCheckInfo) (*types.Transaction, error) {
	return _Teamliquid.contract.Transact(opts, "executeWithAmountCheck", executionInfos, amountCheckInfo)
}

// ExecuteWithAmountCheck is a paid mutator transaction binding the contract method 0x04871891.
//
// Solidity: function executeWithAmountCheck((address,bytes,uint256)[] executionInfos, (address,bytes,uint256) amountCheckInfo) payable returns()
func (_Teamliquid *TeamliquidSession) ExecuteWithAmountCheck(executionInfos []ReservoirV601ExecutionInfo, amountCheckInfo ReservoirV601AmountCheckInfo) (*types.Transaction, error) {
	return _Teamliquid.Contract.ExecuteWithAmountCheck(&_Teamliquid.TransactOpts, executionInfos, amountCheckInfo)
}

// ExecuteWithAmountCheck is a paid mutator transaction binding the contract method 0x04871891.
//
// Solidity: function executeWithAmountCheck((address,bytes,uint256)[] executionInfos, (address,bytes,uint256) amountCheckInfo) payable returns()
func (_Teamliquid *TeamliquidTransactorSession) ExecuteWithAmountCheck(executionInfos []ReservoirV601ExecutionInfo, amountCheckInfo ReservoirV601AmountCheckInfo) (*types.Transaction, error) {
	return _Teamliquid.Contract.ExecuteWithAmountCheck(&_Teamliquid.TransactOpts, executionInfos, amountCheckInfo)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Teamliquid *TeamliquidTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teamliquid.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Teamliquid *TeamliquidSession) Receive() (*types.Transaction, error) {
	return _Teamliquid.Contract.Receive(&_Teamliquid.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Teamliquid *TeamliquidTransactorSession) Receive() (*types.Transaction, error) {
	return _Teamliquid.Contract.Receive(&_Teamliquid.TransactOpts)
}
