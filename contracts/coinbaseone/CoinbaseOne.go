// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coinbaseone

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

// CollectionCreationRequest is an auto generated low-level Go binding around an user-defined struct.
type CollectionCreationRequest struct {
	Creator      common.Address
	Name         string
	Description  string
	Symbol       string
	Image        string
	AnimationUrl string
	MintType     string
	MaxSupply    *big.Int
	MaxPerWallet *big.Int
	Cost         *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Royalties    []Royalty
	Nonce        *big.Int
}

// Royalty is an auto generated low-level Go binding around an user-defined struct.
type Royalty struct {
	To  common.Address
	Amt *big.Int
}

// CoinbaseoneMetaData contains all meta data concerning the Coinbaseone contract.
var CoinbaseoneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedToSendEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"IncorrectETHAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidCollectionRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverClaimLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"TokenForgeMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"TokenForgeMintComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"mintingContract_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"mintWithComment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"metadata_\",\"type\":\"tuple\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoinbaseoneABI is the input ABI used to generate the binding from.
// Deprecated: Use CoinbaseoneMetaData.ABI instead.
var CoinbaseoneABI = CoinbaseoneMetaData.ABI

// Coinbaseone is an auto generated Go binding around an Ethereum contract.
type Coinbaseone struct {
	CoinbaseoneCaller     // Read-only binding to the contract
	CoinbaseoneTransactor // Write-only binding to the contract
	CoinbaseoneFilterer   // Log filterer for contract events
}

// CoinbaseoneCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinbaseoneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseoneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinbaseoneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseoneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinbaseoneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseoneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinbaseoneSession struct {
	Contract     *Coinbaseone      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinbaseoneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinbaseoneCallerSession struct {
	Contract *CoinbaseoneCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CoinbaseoneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinbaseoneTransactorSession struct {
	Contract     *CoinbaseoneTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CoinbaseoneRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinbaseoneRaw struct {
	Contract *Coinbaseone // Generic contract binding to access the raw methods on
}

// CoinbaseoneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinbaseoneCallerRaw struct {
	Contract *CoinbaseoneCaller // Generic read-only contract binding to access the raw methods on
}

// CoinbaseoneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinbaseoneTransactorRaw struct {
	Contract *CoinbaseoneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinbaseone creates a new instance of Coinbaseone, bound to a specific deployed contract.
func NewCoinbaseone(address common.Address, backend bind.ContractBackend) (*Coinbaseone, error) {
	contract, err := bindCoinbaseone(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coinbaseone{CoinbaseoneCaller: CoinbaseoneCaller{contract: contract}, CoinbaseoneTransactor: CoinbaseoneTransactor{contract: contract}, CoinbaseoneFilterer: CoinbaseoneFilterer{contract: contract}}, nil
}

// NewCoinbaseoneCaller creates a new read-only instance of Coinbaseone, bound to a specific deployed contract.
func NewCoinbaseoneCaller(address common.Address, caller bind.ContractCaller) (*CoinbaseoneCaller, error) {
	contract, err := bindCoinbaseone(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneCaller{contract: contract}, nil
}

// NewCoinbaseoneTransactor creates a new write-only instance of Coinbaseone, bound to a specific deployed contract.
func NewCoinbaseoneTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinbaseoneTransactor, error) {
	contract, err := bindCoinbaseone(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneTransactor{contract: contract}, nil
}

// NewCoinbaseoneFilterer creates a new log filterer instance of Coinbaseone, bound to a specific deployed contract.
func NewCoinbaseoneFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinbaseoneFilterer, error) {
	contract, err := bindCoinbaseone(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneFilterer{contract: contract}, nil
}

// bindCoinbaseone binds a generic wrapper to an already deployed contract.
func bindCoinbaseone(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoinbaseoneMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coinbaseone *CoinbaseoneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coinbaseone.Contract.CoinbaseoneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coinbaseone *CoinbaseoneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseone.Contract.CoinbaseoneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coinbaseone *CoinbaseoneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coinbaseone.Contract.CoinbaseoneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coinbaseone *CoinbaseoneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coinbaseone.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coinbaseone *CoinbaseoneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseone.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coinbaseone *CoinbaseoneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coinbaseone.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseone *CoinbaseoneCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseone *CoinbaseoneSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coinbaseone.Contract.BalanceOf(&_Coinbaseone.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseone *CoinbaseoneCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coinbaseone.Contract.BalanceOf(&_Coinbaseone.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseone *CoinbaseoneCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseone *CoinbaseoneSession) ContractURI() (string, error) {
	return _Coinbaseone.Contract.ContractURI(&_Coinbaseone.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseone *CoinbaseoneCallerSession) ContractURI() (string, error) {
	return _Coinbaseone.Contract.ContractURI(&_Coinbaseone.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseone *CoinbaseoneCaller) Cost(opts *bind.CallOpts, quantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "cost", quantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseone *CoinbaseoneSession) Cost(quantity *big.Int) (*big.Int, error) {
	return _Coinbaseone.Contract.Cost(&_Coinbaseone.CallOpts, quantity)
}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseone *CoinbaseoneCallerSession) Cost(quantity *big.Int) (*big.Int, error) {
	return _Coinbaseone.Contract.Cost(&_Coinbaseone.CallOpts, quantity)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseone *CoinbaseoneCaller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseone *CoinbaseoneSession) CurrentTokenId() (*big.Int, error) {
	return _Coinbaseone.Contract.CurrentTokenId(&_Coinbaseone.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseone *CoinbaseoneCallerSession) CurrentTokenId() (*big.Int, error) {
	return _Coinbaseone.Contract.CurrentTokenId(&_Coinbaseone.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseone *CoinbaseoneCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseone *CoinbaseoneSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Coinbaseone.Contract.GetApproved(&_Coinbaseone.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseone *CoinbaseoneCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Coinbaseone.Contract.GetApproved(&_Coinbaseone.CallOpts, arg0)
}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseone *CoinbaseoneCaller) GetMetadata(opts *bind.CallOpts) (CollectionCreationRequest, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "getMetadata")

	if err != nil {
		return *new(CollectionCreationRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(CollectionCreationRequest)).(*CollectionCreationRequest)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseone *CoinbaseoneSession) GetMetadata() (CollectionCreationRequest, error) {
	return _Coinbaseone.Contract.GetMetadata(&_Coinbaseone.CallOpts)
}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseone *CoinbaseoneCallerSession) GetMetadata() (CollectionCreationRequest, error) {
	return _Coinbaseone.Contract.GetMetadata(&_Coinbaseone.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseone *CoinbaseoneCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseone *CoinbaseoneSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Coinbaseone.Contract.IsApprovedForAll(&_Coinbaseone.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseone *CoinbaseoneCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Coinbaseone.Contract.IsApprovedForAll(&_Coinbaseone.CallOpts, arg0, arg1)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(address creator, string name, string description, string symbol, string image, string animation_url, string mintType, uint128 maxSupply, uint128 maxPerWallet, uint256 cost, uint256 startTime, uint256 endTime, uint256 nonce)
func (_Coinbaseone *CoinbaseoneCaller) Metadata(opts *bind.CallOpts) (struct {
	Creator      common.Address
	Name         string
	Description  string
	Symbol       string
	Image        string
	AnimationUrl string
	MintType     string
	MaxSupply    *big.Int
	MaxPerWallet *big.Int
	Cost         *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Nonce        *big.Int
}, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "metadata")

	outstruct := new(struct {
		Creator      common.Address
		Name         string
		Description  string
		Symbol       string
		Image        string
		AnimationUrl string
		MintType     string
		MaxSupply    *big.Int
		MaxPerWallet *big.Int
		Cost         *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Nonce        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Image = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.AnimationUrl = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.MintType = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.MaxSupply = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MaxPerWallet = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Cost = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(address creator, string name, string description, string symbol, string image, string animation_url, string mintType, uint128 maxSupply, uint128 maxPerWallet, uint256 cost, uint256 startTime, uint256 endTime, uint256 nonce)
func (_Coinbaseone *CoinbaseoneSession) Metadata() (struct {
	Creator      common.Address
	Name         string
	Description  string
	Symbol       string
	Image        string
	AnimationUrl string
	MintType     string
	MaxSupply    *big.Int
	MaxPerWallet *big.Int
	Cost         *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Nonce        *big.Int
}, error) {
	return _Coinbaseone.Contract.Metadata(&_Coinbaseone.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(address creator, string name, string description, string symbol, string image, string animation_url, string mintType, uint128 maxSupply, uint128 maxPerWallet, uint256 cost, uint256 startTime, uint256 endTime, uint256 nonce)
func (_Coinbaseone *CoinbaseoneCallerSession) Metadata() (struct {
	Creator      common.Address
	Name         string
	Description  string
	Symbol       string
	Image        string
	AnimationUrl string
	MintType     string
	MaxSupply    *big.Int
	MaxPerWallet *big.Int
	Cost         *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Nonce        *big.Int
}, error) {
	return _Coinbaseone.Contract.Metadata(&_Coinbaseone.CallOpts)
}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseone *CoinbaseoneCaller) MintingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "mintingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseone *CoinbaseoneSession) MintingContract() (common.Address, error) {
	return _Coinbaseone.Contract.MintingContract(&_Coinbaseone.CallOpts)
}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseone *CoinbaseoneCallerSession) MintingContract() (common.Address, error) {
	return _Coinbaseone.Contract.MintingContract(&_Coinbaseone.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseone *CoinbaseoneCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseone *CoinbaseoneSession) Name() (string, error) {
	return _Coinbaseone.Contract.Name(&_Coinbaseone.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseone *CoinbaseoneCallerSession) Name() (string, error) {
	return _Coinbaseone.Contract.Name(&_Coinbaseone.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseone *CoinbaseoneCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseone *CoinbaseoneSession) Owner() (common.Address, error) {
	return _Coinbaseone.Contract.Owner(&_Coinbaseone.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseone *CoinbaseoneCallerSession) Owner() (common.Address, error) {
	return _Coinbaseone.Contract.Owner(&_Coinbaseone.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseone *CoinbaseoneCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseone *CoinbaseoneSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Coinbaseone.Contract.OwnerOf(&_Coinbaseone.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseone *CoinbaseoneCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Coinbaseone.Contract.OwnerOf(&_Coinbaseone.CallOpts, id)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseone *CoinbaseoneCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseone *CoinbaseoneSession) PendingOwner() (common.Address, error) {
	return _Coinbaseone.Contract.PendingOwner(&_Coinbaseone.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseone *CoinbaseoneCallerSession) PendingOwner() (common.Address, error) {
	return _Coinbaseone.Contract.PendingOwner(&_Coinbaseone.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseone *CoinbaseoneCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseone *CoinbaseoneSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coinbaseone.Contract.SupportsInterface(&_Coinbaseone.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseone *CoinbaseoneCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coinbaseone.Contract.SupportsInterface(&_Coinbaseone.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseone *CoinbaseoneCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseone *CoinbaseoneSession) Symbol() (string, error) {
	return _Coinbaseone.Contract.Symbol(&_Coinbaseone.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseone *CoinbaseoneCallerSession) Symbol() (string, error) {
	return _Coinbaseone.Contract.Symbol(&_Coinbaseone.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseone *CoinbaseoneCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseone *CoinbaseoneSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coinbaseone.Contract.TokenURI(&_Coinbaseone.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseone *CoinbaseoneCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coinbaseone.Contract.TokenURI(&_Coinbaseone.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseone *CoinbaseoneCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseone.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseone *CoinbaseoneSession) TotalSupply() (*big.Int, error) {
	return _Coinbaseone.Contract.TotalSupply(&_Coinbaseone.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseone *CoinbaseoneCallerSession) TotalSupply() (*big.Int, error) {
	return _Coinbaseone.Contract.TotalSupply(&_Coinbaseone.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseone *CoinbaseoneTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseone *CoinbaseoneSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coinbaseone.Contract.AcceptOwnership(&_Coinbaseone.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coinbaseone.Contract.AcceptOwnership(&_Coinbaseone.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Approve(&_Coinbaseone.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Approve(&_Coinbaseone.TransactOpts, spender, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseone *CoinbaseoneTransactor) Initialize(opts *bind.TransactOpts, request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "initialize", request, mintingContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseone *CoinbaseoneSession) Initialize(request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Initialize(&_Coinbaseone.TransactOpts, request, mintingContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) Initialize(request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Initialize(&_Coinbaseone.TransactOpts, request, mintingContract_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseone *CoinbaseoneTransactor) Mint(opts *bind.TransactOpts, to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "mint", to, quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseone *CoinbaseoneSession) Mint(to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Mint(&_Coinbaseone.TransactOpts, to, quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) Mint(to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.Mint(&_Coinbaseone.TransactOpts, to, quantity)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseone *CoinbaseoneTransactor) MintWithComment(opts *bind.TransactOpts, to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "mintWithComment", to, quantity, comment)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseone *CoinbaseoneSession) MintWithComment(to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseone.Contract.MintWithComment(&_Coinbaseone.TransactOpts, to, quantity, comment)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) MintWithComment(to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseone.Contract.MintWithComment(&_Coinbaseone.TransactOpts, to, quantity, comment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseone *CoinbaseoneTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseone *CoinbaseoneSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coinbaseone.Contract.RenounceOwnership(&_Coinbaseone.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coinbaseone.Contract.RenounceOwnership(&_Coinbaseone.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SafeTransferFrom(&_Coinbaseone.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SafeTransferFrom(&_Coinbaseone.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseone *CoinbaseoneTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseone *CoinbaseoneSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SafeTransferFrom0(&_Coinbaseone.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SafeTransferFrom0(&_Coinbaseone.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseone *CoinbaseoneTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseone *CoinbaseoneSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SetApprovalForAll(&_Coinbaseone.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SetApprovalForAll(&_Coinbaseone.TransactOpts, operator, approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseone *CoinbaseoneTransactor) SetMetadata(opts *bind.TransactOpts, metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "setMetadata", metadata_)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseone *CoinbaseoneSession) SetMetadata(metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SetMetadata(&_Coinbaseone.TransactOpts, metadata_)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) SetMetadata(metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseone.Contract.SetMetadata(&_Coinbaseone.TransactOpts, metadata_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.TransferFrom(&_Coinbaseone.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseone.Contract.TransferFrom(&_Coinbaseone.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseone *CoinbaseoneTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseone *CoinbaseoneSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseone.Contract.TransferOwnership(&_Coinbaseone.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseone.Contract.TransferOwnership(&_Coinbaseone.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseone *CoinbaseoneTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseone.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseone *CoinbaseoneSession) Withdraw() (*types.Transaction, error) {
	return _Coinbaseone.Contract.Withdraw(&_Coinbaseone.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseone *CoinbaseoneTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Coinbaseone.Contract.Withdraw(&_Coinbaseone.TransactOpts)
}

// CoinbaseoneApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Coinbaseone contract.
type CoinbaseoneApprovalIterator struct {
	Event *CoinbaseoneApproval // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneApproval)
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
		it.Event = new(CoinbaseoneApproval)
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
func (it *CoinbaseoneApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneApproval represents a Approval event raised by the Coinbaseone contract.
type CoinbaseoneApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*CoinbaseoneApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneApprovalIterator{contract: _Coinbaseone.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CoinbaseoneApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneApproval)
				if err := _Coinbaseone.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) ParseApproval(log types.Log) (*CoinbaseoneApproval, error) {
	event := new(CoinbaseoneApproval)
	if err := _Coinbaseone.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Coinbaseone contract.
type CoinbaseoneApprovalForAllIterator struct {
	Event *CoinbaseoneApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneApprovalForAll)
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
		it.Event = new(CoinbaseoneApprovalForAll)
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
func (it *CoinbaseoneApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneApprovalForAll represents a ApprovalForAll event raised by the Coinbaseone contract.
type CoinbaseoneApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coinbaseone *CoinbaseoneFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CoinbaseoneApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneApprovalForAllIterator{contract: _Coinbaseone.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coinbaseone *CoinbaseoneFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CoinbaseoneApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneApprovalForAll)
				if err := _Coinbaseone.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coinbaseone *CoinbaseoneFilterer) ParseApprovalForAll(log types.Log) (*CoinbaseoneApprovalForAll, error) {
	event := new(CoinbaseoneApprovalForAll)
	if err := _Coinbaseone.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Coinbaseone contract.
type CoinbaseoneInitializedIterator struct {
	Event *CoinbaseoneInitialized // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneInitialized)
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
		it.Event = new(CoinbaseoneInitialized)
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
func (it *CoinbaseoneInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneInitialized represents a Initialized event raised by the Coinbaseone contract.
type CoinbaseoneInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Coinbaseone *CoinbaseoneFilterer) FilterInitialized(opts *bind.FilterOpts) (*CoinbaseoneInitializedIterator, error) {

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneInitializedIterator{contract: _Coinbaseone.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Coinbaseone *CoinbaseoneFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CoinbaseoneInitialized) (event.Subscription, error) {

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneInitialized)
				if err := _Coinbaseone.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Coinbaseone *CoinbaseoneFilterer) ParseInitialized(log types.Log) (*CoinbaseoneInitialized, error) {
	event := new(CoinbaseoneInitialized)
	if err := _Coinbaseone.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneMintConfigChangedIterator is returned from FilterMintConfigChanged and is used to iterate over the raw logs and unpacked data for MintConfigChanged events raised by the Coinbaseone contract.
type CoinbaseoneMintConfigChangedIterator struct {
	Event *CoinbaseoneMintConfigChanged // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneMintConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneMintConfigChanged)
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
		it.Event = new(CoinbaseoneMintConfigChanged)
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
func (it *CoinbaseoneMintConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneMintConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneMintConfigChanged represents a MintConfigChanged event raised by the Coinbaseone contract.
type CoinbaseoneMintConfigChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintConfigChanged is a free log retrieval operation binding the contract event 0xa703f5371c9a5519d27a0ab98ff81ca400a4adb7bf05d607347bfffc0efabe8f.
//
// Solidity: event MintConfigChanged()
func (_Coinbaseone *CoinbaseoneFilterer) FilterMintConfigChanged(opts *bind.FilterOpts) (*CoinbaseoneMintConfigChangedIterator, error) {

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "MintConfigChanged")
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneMintConfigChangedIterator{contract: _Coinbaseone.contract, event: "MintConfigChanged", logs: logs, sub: sub}, nil
}

// WatchMintConfigChanged is a free log subscription operation binding the contract event 0xa703f5371c9a5519d27a0ab98ff81ca400a4adb7bf05d607347bfffc0efabe8f.
//
// Solidity: event MintConfigChanged()
func (_Coinbaseone *CoinbaseoneFilterer) WatchMintConfigChanged(opts *bind.WatchOpts, sink chan<- *CoinbaseoneMintConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "MintConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneMintConfigChanged)
				if err := _Coinbaseone.contract.UnpackLog(event, "MintConfigChanged", log); err != nil {
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

// ParseMintConfigChanged is a log parse operation binding the contract event 0xa703f5371c9a5519d27a0ab98ff81ca400a4adb7bf05d607347bfffc0efabe8f.
//
// Solidity: event MintConfigChanged()
func (_Coinbaseone *CoinbaseoneFilterer) ParseMintConfigChanged(log types.Log) (*CoinbaseoneMintConfigChanged, error) {
	event := new(CoinbaseoneMintConfigChanged)
	if err := _Coinbaseone.contract.UnpackLog(event, "MintConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Coinbaseone contract.
type CoinbaseoneOwnershipTransferStartedIterator struct {
	Event *CoinbaseoneOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneOwnershipTransferStarted)
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
		it.Event = new(CoinbaseoneOwnershipTransferStarted)
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
func (it *CoinbaseoneOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Coinbaseone contract.
type CoinbaseoneOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseone *CoinbaseoneFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinbaseoneOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneOwnershipTransferStartedIterator{contract: _Coinbaseone.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseone *CoinbaseoneFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *CoinbaseoneOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneOwnershipTransferStarted)
				if err := _Coinbaseone.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseone *CoinbaseoneFilterer) ParseOwnershipTransferStarted(log types.Log) (*CoinbaseoneOwnershipTransferStarted, error) {
	event := new(CoinbaseoneOwnershipTransferStarted)
	if err := _Coinbaseone.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Coinbaseone contract.
type CoinbaseoneOwnershipTransferredIterator struct {
	Event *CoinbaseoneOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneOwnershipTransferred)
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
		it.Event = new(CoinbaseoneOwnershipTransferred)
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
func (it *CoinbaseoneOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneOwnershipTransferred represents a OwnershipTransferred event raised by the Coinbaseone contract.
type CoinbaseoneOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseone *CoinbaseoneFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinbaseoneOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneOwnershipTransferredIterator{contract: _Coinbaseone.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseone *CoinbaseoneFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoinbaseoneOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneOwnershipTransferred)
				if err := _Coinbaseone.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Coinbaseone *CoinbaseoneFilterer) ParseOwnershipTransferred(log types.Log) (*CoinbaseoneOwnershipTransferred, error) {
	event := new(CoinbaseoneOwnershipTransferred)
	if err := _Coinbaseone.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneTokenForgeMintIterator is returned from FilterTokenForgeMint and is used to iterate over the raw logs and unpacked data for TokenForgeMint events raised by the Coinbaseone contract.
type CoinbaseoneTokenForgeMintIterator struct {
	Event *CoinbaseoneTokenForgeMint // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneTokenForgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneTokenForgeMint)
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
		it.Event = new(CoinbaseoneTokenForgeMint)
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
func (it *CoinbaseoneTokenForgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneTokenForgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneTokenForgeMint represents a TokenForgeMint event raised by the Coinbaseone contract.
type CoinbaseoneTokenForgeMint struct {
	To       common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenForgeMint is a free log retrieval operation binding the contract event 0x02c36b548faac112a24e09f132ea830e930cc215c0a74a678b70e43aede11edd.
//
// Solidity: event TokenForgeMint(address indexed to, uint256 quantity)
func (_Coinbaseone *CoinbaseoneFilterer) FilterTokenForgeMint(opts *bind.FilterOpts, to []common.Address) (*CoinbaseoneTokenForgeMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "TokenForgeMint", toRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneTokenForgeMintIterator{contract: _Coinbaseone.contract, event: "TokenForgeMint", logs: logs, sub: sub}, nil
}

// WatchTokenForgeMint is a free log subscription operation binding the contract event 0x02c36b548faac112a24e09f132ea830e930cc215c0a74a678b70e43aede11edd.
//
// Solidity: event TokenForgeMint(address indexed to, uint256 quantity)
func (_Coinbaseone *CoinbaseoneFilterer) WatchTokenForgeMint(opts *bind.WatchOpts, sink chan<- *CoinbaseoneTokenForgeMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "TokenForgeMint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneTokenForgeMint)
				if err := _Coinbaseone.contract.UnpackLog(event, "TokenForgeMint", log); err != nil {
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

// ParseTokenForgeMint is a log parse operation binding the contract event 0x02c36b548faac112a24e09f132ea830e930cc215c0a74a678b70e43aede11edd.
//
// Solidity: event TokenForgeMint(address indexed to, uint256 quantity)
func (_Coinbaseone *CoinbaseoneFilterer) ParseTokenForgeMint(log types.Log) (*CoinbaseoneTokenForgeMint, error) {
	event := new(CoinbaseoneTokenForgeMint)
	if err := _Coinbaseone.contract.UnpackLog(event, "TokenForgeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneTokenForgeMintCommentIterator is returned from FilterTokenForgeMintComment and is used to iterate over the raw logs and unpacked data for TokenForgeMintComment events raised by the Coinbaseone contract.
type CoinbaseoneTokenForgeMintCommentIterator struct {
	Event *CoinbaseoneTokenForgeMintComment // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneTokenForgeMintCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneTokenForgeMintComment)
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
		it.Event = new(CoinbaseoneTokenForgeMintComment)
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
func (it *CoinbaseoneTokenForgeMintCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneTokenForgeMintCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneTokenForgeMintComment represents a TokenForgeMintComment event raised by the Coinbaseone contract.
type CoinbaseoneTokenForgeMintComment struct {
	To       common.Address
	Quantity *big.Int
	Comment  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenForgeMintComment is a free log retrieval operation binding the contract event 0x203498aecb28c99d51721d218d93e378293b86eacb26b42246dae394840ae756.
//
// Solidity: event TokenForgeMintComment(address indexed to, uint256 quantity, string comment)
func (_Coinbaseone *CoinbaseoneFilterer) FilterTokenForgeMintComment(opts *bind.FilterOpts, to []common.Address) (*CoinbaseoneTokenForgeMintCommentIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "TokenForgeMintComment", toRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneTokenForgeMintCommentIterator{contract: _Coinbaseone.contract, event: "TokenForgeMintComment", logs: logs, sub: sub}, nil
}

// WatchTokenForgeMintComment is a free log subscription operation binding the contract event 0x203498aecb28c99d51721d218d93e378293b86eacb26b42246dae394840ae756.
//
// Solidity: event TokenForgeMintComment(address indexed to, uint256 quantity, string comment)
func (_Coinbaseone *CoinbaseoneFilterer) WatchTokenForgeMintComment(opts *bind.WatchOpts, sink chan<- *CoinbaseoneTokenForgeMintComment, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "TokenForgeMintComment", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneTokenForgeMintComment)
				if err := _Coinbaseone.contract.UnpackLog(event, "TokenForgeMintComment", log); err != nil {
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

// ParseTokenForgeMintComment is a log parse operation binding the contract event 0x203498aecb28c99d51721d218d93e378293b86eacb26b42246dae394840ae756.
//
// Solidity: event TokenForgeMintComment(address indexed to, uint256 quantity, string comment)
func (_Coinbaseone *CoinbaseoneFilterer) ParseTokenForgeMintComment(log types.Log) (*CoinbaseoneTokenForgeMintComment, error) {
	event := new(CoinbaseoneTokenForgeMintComment)
	if err := _Coinbaseone.contract.UnpackLog(event, "TokenForgeMintComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Coinbaseone contract.
type CoinbaseoneTransferIterator struct {
	Event *CoinbaseoneTransfer // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneTransfer)
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
		it.Event = new(CoinbaseoneTransfer)
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
func (it *CoinbaseoneTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneTransfer represents a Transfer event raised by the Coinbaseone contract.
type CoinbaseoneTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*CoinbaseoneTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneTransferIterator{contract: _Coinbaseone.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CoinbaseoneTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneTransfer)
				if err := _Coinbaseone.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Coinbaseone *CoinbaseoneFilterer) ParseTransfer(log types.Log) (*CoinbaseoneTransfer, error) {
	event := new(CoinbaseoneTransfer)
	if err := _Coinbaseone.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseoneWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Coinbaseone contract.
type CoinbaseoneWithdrawnIterator struct {
	Event *CoinbaseoneWithdrawn // Event containing the contract specifics and raw log

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
func (it *CoinbaseoneWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseoneWithdrawn)
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
		it.Event = new(CoinbaseoneWithdrawn)
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
func (it *CoinbaseoneWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseoneWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseoneWithdrawn represents a Withdrawn event raised by the Coinbaseone contract.
type CoinbaseoneWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 amount)
func (_Coinbaseone *CoinbaseoneFilterer) FilterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*CoinbaseoneWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Coinbaseone.contract.FilterLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseoneWithdrawnIterator{contract: _Coinbaseone.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 amount)
func (_Coinbaseone *CoinbaseoneFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CoinbaseoneWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Coinbaseone.contract.WatchLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseoneWithdrawn)
				if err := _Coinbaseone.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 amount)
func (_Coinbaseone *CoinbaseoneFilterer) ParseWithdrawn(log types.Log) (*CoinbaseoneWithdrawn, error) {
	event := new(CoinbaseoneWithdrawn)
	if err := _Coinbaseone.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
