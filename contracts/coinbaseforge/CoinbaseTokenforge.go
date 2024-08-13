// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coinbaseforge

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

// CoinbaseforgeMetaData contains all meta data concerning the Coinbaseforge contract.
var CoinbaseforgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedToSendEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"IncorrectETHAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidCollectionRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverClaimLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"TokenForgeMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"TokenForgeMintComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"mintingContract_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"mintWithComment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"animation_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mintType\",\"type\":\"string\"},{\"internalType\":\"uint128\",\"name\":\"maxSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"internalType\":\"structRoyalty[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structCollectionCreationRequest\",\"name\":\"metadata_\",\"type\":\"tuple\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CoinbaseforgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CoinbaseforgeMetaData.ABI instead.
var CoinbaseforgeABI = CoinbaseforgeMetaData.ABI

// Coinbaseforge is an auto generated Go binding around an Ethereum contract.
type Coinbaseforge struct {
	CoinbaseforgeCaller     // Read-only binding to the contract
	CoinbaseforgeTransactor // Write-only binding to the contract
	CoinbaseforgeFilterer   // Log filterer for contract events
}

// CoinbaseforgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinbaseforgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseforgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinbaseforgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseforgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinbaseforgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinbaseforgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinbaseforgeSession struct {
	Contract     *Coinbaseforge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinbaseforgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinbaseforgeCallerSession struct {
	Contract *CoinbaseforgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CoinbaseforgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinbaseforgeTransactorSession struct {
	Contract     *CoinbaseforgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CoinbaseforgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinbaseforgeRaw struct {
	Contract *Coinbaseforge // Generic contract binding to access the raw methods on
}

// CoinbaseforgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinbaseforgeCallerRaw struct {
	Contract *CoinbaseforgeCaller // Generic read-only contract binding to access the raw methods on
}

// CoinbaseforgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinbaseforgeTransactorRaw struct {
	Contract *CoinbaseforgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinbaseforge creates a new instance of Coinbaseforge, bound to a specific deployed contract.
func NewCoinbaseforge(address common.Address, backend bind.ContractBackend) (*Coinbaseforge, error) {
	contract, err := bindCoinbaseforge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coinbaseforge{CoinbaseforgeCaller: CoinbaseforgeCaller{contract: contract}, CoinbaseforgeTransactor: CoinbaseforgeTransactor{contract: contract}, CoinbaseforgeFilterer: CoinbaseforgeFilterer{contract: contract}}, nil
}

// NewCoinbaseforgeCaller creates a new read-only instance of Coinbaseforge, bound to a specific deployed contract.
func NewCoinbaseforgeCaller(address common.Address, caller bind.ContractCaller) (*CoinbaseforgeCaller, error) {
	contract, err := bindCoinbaseforge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeCaller{contract: contract}, nil
}

// NewCoinbaseforgeTransactor creates a new write-only instance of Coinbaseforge, bound to a specific deployed contract.
func NewCoinbaseforgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinbaseforgeTransactor, error) {
	contract, err := bindCoinbaseforge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeTransactor{contract: contract}, nil
}

// NewCoinbaseforgeFilterer creates a new log filterer instance of Coinbaseforge, bound to a specific deployed contract.
func NewCoinbaseforgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinbaseforgeFilterer, error) {
	contract, err := bindCoinbaseforge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeFilterer{contract: contract}, nil
}

// bindCoinbaseforge binds a generic wrapper to an already deployed contract.
func bindCoinbaseforge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoinbaseforgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coinbaseforge *CoinbaseforgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coinbaseforge.Contract.CoinbaseforgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coinbaseforge *CoinbaseforgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.CoinbaseforgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coinbaseforge *CoinbaseforgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.CoinbaseforgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coinbaseforge *CoinbaseforgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coinbaseforge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coinbaseforge *CoinbaseforgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coinbaseforge *CoinbaseforgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coinbaseforge.Contract.BalanceOf(&_Coinbaseforge.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Coinbaseforge.Contract.BalanceOf(&_Coinbaseforge.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseforge *CoinbaseforgeSession) ContractURI() (string, error) {
	return _Coinbaseforge.Contract.ContractURI(&_Coinbaseforge.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCallerSession) ContractURI() (string, error) {
	return _Coinbaseforge.Contract.ContractURI(&_Coinbaseforge.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCaller) Cost(opts *bind.CallOpts, quantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "cost", quantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeSession) Cost(quantity *big.Int) (*big.Int, error) {
	return _Coinbaseforge.Contract.Cost(&_Coinbaseforge.CallOpts, quantity)
}

// Cost is a free data retrieval call binding the contract method 0x9097548d.
//
// Solidity: function cost(uint256 quantity) view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCallerSession) Cost(quantity *big.Int) (*big.Int, error) {
	return _Coinbaseforge.Contract.Cost(&_Coinbaseforge.CallOpts, quantity)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCaller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeSession) CurrentTokenId() (*big.Int, error) {
	return _Coinbaseforge.Contract.CurrentTokenId(&_Coinbaseforge.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCallerSession) CurrentTokenId() (*big.Int, error) {
	return _Coinbaseforge.Contract.CurrentTokenId(&_Coinbaseforge.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseforge *CoinbaseforgeCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseforge *CoinbaseforgeSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Coinbaseforge.Contract.GetApproved(&_Coinbaseforge.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Coinbaseforge *CoinbaseforgeCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Coinbaseforge.Contract.GetApproved(&_Coinbaseforge.CallOpts, arg0)
}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseforge *CoinbaseforgeCaller) GetMetadata(opts *bind.CallOpts) (CollectionCreationRequest, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "getMetadata")

	if err != nil {
		return *new(CollectionCreationRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(CollectionCreationRequest)).(*CollectionCreationRequest)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseforge *CoinbaseforgeSession) GetMetadata() (CollectionCreationRequest, error) {
	return _Coinbaseforge.Contract.GetMetadata(&_Coinbaseforge.CallOpts)
}

// GetMetadata is a free data retrieval call binding the contract method 0x7a5b4f59.
//
// Solidity: function getMetadata() view returns((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256))
func (_Coinbaseforge *CoinbaseforgeCallerSession) GetMetadata() (CollectionCreationRequest, error) {
	return _Coinbaseforge.Contract.GetMetadata(&_Coinbaseforge.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Coinbaseforge.Contract.IsApprovedForAll(&_Coinbaseforge.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Coinbaseforge.Contract.IsApprovedForAll(&_Coinbaseforge.CallOpts, arg0, arg1)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(address creator, string name, string description, string symbol, string image, string animation_url, string mintType, uint128 maxSupply, uint128 maxPerWallet, uint256 cost, uint256 startTime, uint256 endTime, uint256 nonce)
func (_Coinbaseforge *CoinbaseforgeCaller) Metadata(opts *bind.CallOpts) (struct {
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
	err := _Coinbaseforge.contract.Call(opts, &out, "metadata")

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
func (_Coinbaseforge *CoinbaseforgeSession) Metadata() (struct {
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
	return _Coinbaseforge.Contract.Metadata(&_Coinbaseforge.CallOpts)
}

// Metadata is a free data retrieval call binding the contract method 0x392f37e9.
//
// Solidity: function metadata() view returns(address creator, string name, string description, string symbol, string image, string animation_url, string mintType, uint128 maxSupply, uint128 maxPerWallet, uint256 cost, uint256 startTime, uint256 endTime, uint256 nonce)
func (_Coinbaseforge *CoinbaseforgeCallerSession) Metadata() (struct {
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
	return _Coinbaseforge.Contract.Metadata(&_Coinbaseforge.CallOpts)
}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCaller) MintingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "mintingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseforge *CoinbaseforgeSession) MintingContract() (common.Address, error) {
	return _Coinbaseforge.Contract.MintingContract(&_Coinbaseforge.CallOpts)
}

// MintingContract is a free data retrieval call binding the contract method 0xd2f6f67d.
//
// Solidity: function mintingContract() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCallerSession) MintingContract() (common.Address, error) {
	return _Coinbaseforge.Contract.MintingContract(&_Coinbaseforge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseforge *CoinbaseforgeSession) Name() (string, error) {
	return _Coinbaseforge.Contract.Name(&_Coinbaseforge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCallerSession) Name() (string, error) {
	return _Coinbaseforge.Contract.Name(&_Coinbaseforge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeSession) Owner() (common.Address, error) {
	return _Coinbaseforge.Contract.Owner(&_Coinbaseforge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCallerSession) Owner() (common.Address, error) {
	return _Coinbaseforge.Contract.Owner(&_Coinbaseforge.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseforge *CoinbaseforgeCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseforge *CoinbaseforgeSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Coinbaseforge.Contract.OwnerOf(&_Coinbaseforge.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Coinbaseforge *CoinbaseforgeCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Coinbaseforge.Contract.OwnerOf(&_Coinbaseforge.CallOpts, id)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeSession) PendingOwner() (common.Address, error) {
	return _Coinbaseforge.Contract.PendingOwner(&_Coinbaseforge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Coinbaseforge *CoinbaseforgeCallerSession) PendingOwner() (common.Address, error) {
	return _Coinbaseforge.Contract.PendingOwner(&_Coinbaseforge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coinbaseforge.Contract.SupportsInterface(&_Coinbaseforge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Coinbaseforge *CoinbaseforgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Coinbaseforge.Contract.SupportsInterface(&_Coinbaseforge.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseforge *CoinbaseforgeSession) Symbol() (string, error) {
	return _Coinbaseforge.Contract.Symbol(&_Coinbaseforge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Coinbaseforge *CoinbaseforgeCallerSession) Symbol() (string, error) {
	return _Coinbaseforge.Contract.Symbol(&_Coinbaseforge.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseforge *CoinbaseforgeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseforge *CoinbaseforgeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coinbaseforge.Contract.TokenURI(&_Coinbaseforge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Coinbaseforge *CoinbaseforgeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Coinbaseforge.Contract.TokenURI(&_Coinbaseforge.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Coinbaseforge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeSession) TotalSupply() (*big.Int, error) {
	return _Coinbaseforge.Contract.TotalSupply(&_Coinbaseforge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Coinbaseforge *CoinbaseforgeCallerSession) TotalSupply() (*big.Int, error) {
	return _Coinbaseforge.Contract.TotalSupply(&_Coinbaseforge.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.AcceptOwnership(&_Coinbaseforge.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.AcceptOwnership(&_Coinbaseforge.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Approve(&_Coinbaseforge.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Approve(&_Coinbaseforge.TransactOpts, spender, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) Initialize(opts *bind.TransactOpts, request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "initialize", request, mintingContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseforge *CoinbaseforgeSession) Initialize(request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Initialize(&_Coinbaseforge.TransactOpts, request, mintingContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x33f014f0.
//
// Solidity: function initialize((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) request, address mintingContract_) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) Initialize(request CollectionCreationRequest, mintingContract_ common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Initialize(&_Coinbaseforge.TransactOpts, request, mintingContract_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) Mint(opts *bind.TransactOpts, to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "mint", to, quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseforge *CoinbaseforgeSession) Mint(to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Mint(&_Coinbaseforge.TransactOpts, to, quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 quantity) payable returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) Mint(to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Mint(&_Coinbaseforge.TransactOpts, to, quantity)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) MintWithComment(opts *bind.TransactOpts, to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "mintWithComment", to, quantity, comment)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseforge *CoinbaseforgeSession) MintWithComment(to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.MintWithComment(&_Coinbaseforge.TransactOpts, to, quantity, comment)
}

// MintWithComment is a paid mutator transaction binding the contract method 0x574fed17.
//
// Solidity: function mintWithComment(address to, uint256 quantity, string comment) payable returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) MintWithComment(to common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.MintWithComment(&_Coinbaseforge.TransactOpts, to, quantity, comment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.RenounceOwnership(&_Coinbaseforge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.RenounceOwnership(&_Coinbaseforge.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SafeTransferFrom(&_Coinbaseforge.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SafeTransferFrom(&_Coinbaseforge.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseforge *CoinbaseforgeSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SafeTransferFrom0(&_Coinbaseforge.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SafeTransferFrom0(&_Coinbaseforge.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseforge *CoinbaseforgeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SetApprovalForAll(&_Coinbaseforge.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SetApprovalForAll(&_Coinbaseforge.TransactOpts, operator, approved)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) SetMetadata(opts *bind.TransactOpts, metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "setMetadata", metadata_)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseforge *CoinbaseforgeSession) SetMetadata(metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SetMetadata(&_Coinbaseforge.TransactOpts, metadata_)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x792a0932.
//
// Solidity: function setMetadata((address,string,string,string,string,string,string,uint128,uint128,uint256,uint256,uint256,(address,uint256)[],uint256) metadata_) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) SetMetadata(metadata_ CollectionCreationRequest) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.SetMetadata(&_Coinbaseforge.TransactOpts, metadata_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.TransferFrom(&_Coinbaseforge.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.TransferFrom(&_Coinbaseforge.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseforge *CoinbaseforgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.TransferOwnership(&_Coinbaseforge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Coinbaseforge.Contract.TransferOwnership(&_Coinbaseforge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseforge *CoinbaseforgeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coinbaseforge.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseforge *CoinbaseforgeSession) Withdraw() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Withdraw(&_Coinbaseforge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Coinbaseforge *CoinbaseforgeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Coinbaseforge.Contract.Withdraw(&_Coinbaseforge.TransactOpts)
}

// CoinbaseforgeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Coinbaseforge contract.
type CoinbaseforgeApprovalIterator struct {
	Event *CoinbaseforgeApproval // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeApproval)
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
		it.Event = new(CoinbaseforgeApproval)
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
func (it *CoinbaseforgeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeApproval represents a Approval event raised by the Coinbaseforge contract.
type CoinbaseforgeApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*CoinbaseforgeApprovalIterator, error) {

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

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeApprovalIterator{contract: _Coinbaseforge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeApproval)
				if err := _Coinbaseforge.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseApproval(log types.Log) (*CoinbaseforgeApproval, error) {
	event := new(CoinbaseforgeApproval)
	if err := _Coinbaseforge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Coinbaseforge contract.
type CoinbaseforgeApprovalForAllIterator struct {
	Event *CoinbaseforgeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeApprovalForAll)
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
		it.Event = new(CoinbaseforgeApprovalForAll)
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
func (it *CoinbaseforgeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeApprovalForAll represents a ApprovalForAll event raised by the Coinbaseforge contract.
type CoinbaseforgeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CoinbaseforgeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeApprovalForAllIterator{contract: _Coinbaseforge.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeApprovalForAll)
				if err := _Coinbaseforge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseApprovalForAll(log types.Log) (*CoinbaseforgeApprovalForAll, error) {
	event := new(CoinbaseforgeApprovalForAll)
	if err := _Coinbaseforge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Coinbaseforge contract.
type CoinbaseforgeInitializedIterator struct {
	Event *CoinbaseforgeInitialized // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeInitialized)
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
		it.Event = new(CoinbaseforgeInitialized)
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
func (it *CoinbaseforgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeInitialized represents a Initialized event raised by the Coinbaseforge contract.
type CoinbaseforgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CoinbaseforgeInitializedIterator, error) {

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeInitializedIterator{contract: _Coinbaseforge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeInitialized)
				if err := _Coinbaseforge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseInitialized(log types.Log) (*CoinbaseforgeInitialized, error) {
	event := new(CoinbaseforgeInitialized)
	if err := _Coinbaseforge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeMintConfigChangedIterator is returned from FilterMintConfigChanged and is used to iterate over the raw logs and unpacked data for MintConfigChanged events raised by the Coinbaseforge contract.
type CoinbaseforgeMintConfigChangedIterator struct {
	Event *CoinbaseforgeMintConfigChanged // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeMintConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeMintConfigChanged)
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
		it.Event = new(CoinbaseforgeMintConfigChanged)
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
func (it *CoinbaseforgeMintConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeMintConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeMintConfigChanged represents a MintConfigChanged event raised by the Coinbaseforge contract.
type CoinbaseforgeMintConfigChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintConfigChanged is a free log retrieval operation binding the contract event 0xa703f5371c9a5519d27a0ab98ff81ca400a4adb7bf05d607347bfffc0efabe8f.
//
// Solidity: event MintConfigChanged()
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterMintConfigChanged(opts *bind.FilterOpts) (*CoinbaseforgeMintConfigChangedIterator, error) {

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "MintConfigChanged")
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeMintConfigChangedIterator{contract: _Coinbaseforge.contract, event: "MintConfigChanged", logs: logs, sub: sub}, nil
}

// WatchMintConfigChanged is a free log subscription operation binding the contract event 0xa703f5371c9a5519d27a0ab98ff81ca400a4adb7bf05d607347bfffc0efabe8f.
//
// Solidity: event MintConfigChanged()
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchMintConfigChanged(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeMintConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "MintConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeMintConfigChanged)
				if err := _Coinbaseforge.contract.UnpackLog(event, "MintConfigChanged", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseMintConfigChanged(log types.Log) (*CoinbaseforgeMintConfigChanged, error) {
	event := new(CoinbaseforgeMintConfigChanged)
	if err := _Coinbaseforge.contract.UnpackLog(event, "MintConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Coinbaseforge contract.
type CoinbaseforgeOwnershipTransferStartedIterator struct {
	Event *CoinbaseforgeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeOwnershipTransferStarted)
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
		it.Event = new(CoinbaseforgeOwnershipTransferStarted)
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
func (it *CoinbaseforgeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Coinbaseforge contract.
type CoinbaseforgeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinbaseforgeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeOwnershipTransferStartedIterator{contract: _Coinbaseforge.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeOwnershipTransferStarted)
				if err := _Coinbaseforge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseOwnershipTransferStarted(log types.Log) (*CoinbaseforgeOwnershipTransferStarted, error) {
	event := new(CoinbaseforgeOwnershipTransferStarted)
	if err := _Coinbaseforge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Coinbaseforge contract.
type CoinbaseforgeOwnershipTransferredIterator struct {
	Event *CoinbaseforgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeOwnershipTransferred)
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
		it.Event = new(CoinbaseforgeOwnershipTransferred)
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
func (it *CoinbaseforgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeOwnershipTransferred represents a OwnershipTransferred event raised by the Coinbaseforge contract.
type CoinbaseforgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinbaseforgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeOwnershipTransferredIterator{contract: _Coinbaseforge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeOwnershipTransferred)
				if err := _Coinbaseforge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseOwnershipTransferred(log types.Log) (*CoinbaseforgeOwnershipTransferred, error) {
	event := new(CoinbaseforgeOwnershipTransferred)
	if err := _Coinbaseforge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeTokenForgeMintIterator is returned from FilterTokenForgeMint and is used to iterate over the raw logs and unpacked data for TokenForgeMint events raised by the Coinbaseforge contract.
type CoinbaseforgeTokenForgeMintIterator struct {
	Event *CoinbaseforgeTokenForgeMint // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeTokenForgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeTokenForgeMint)
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
		it.Event = new(CoinbaseforgeTokenForgeMint)
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
func (it *CoinbaseforgeTokenForgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeTokenForgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeTokenForgeMint represents a TokenForgeMint event raised by the Coinbaseforge contract.
type CoinbaseforgeTokenForgeMint struct {
	To       common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenForgeMint is a free log retrieval operation binding the contract event 0x02c36b548faac112a24e09f132ea830e930cc215c0a74a678b70e43aede11edd.
//
// Solidity: event TokenForgeMint(address indexed to, uint256 quantity)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterTokenForgeMint(opts *bind.FilterOpts, to []common.Address) (*CoinbaseforgeTokenForgeMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "TokenForgeMint", toRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeTokenForgeMintIterator{contract: _Coinbaseforge.contract, event: "TokenForgeMint", logs: logs, sub: sub}, nil
}

// WatchTokenForgeMint is a free log subscription operation binding the contract event 0x02c36b548faac112a24e09f132ea830e930cc215c0a74a678b70e43aede11edd.
//
// Solidity: event TokenForgeMint(address indexed to, uint256 quantity)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchTokenForgeMint(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeTokenForgeMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "TokenForgeMint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeTokenForgeMint)
				if err := _Coinbaseforge.contract.UnpackLog(event, "TokenForgeMint", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseTokenForgeMint(log types.Log) (*CoinbaseforgeTokenForgeMint, error) {
	event := new(CoinbaseforgeTokenForgeMint)
	if err := _Coinbaseforge.contract.UnpackLog(event, "TokenForgeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeTokenForgeMintCommentIterator is returned from FilterTokenForgeMintComment and is used to iterate over the raw logs and unpacked data for TokenForgeMintComment events raised by the Coinbaseforge contract.
type CoinbaseforgeTokenForgeMintCommentIterator struct {
	Event *CoinbaseforgeTokenForgeMintComment // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeTokenForgeMintCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeTokenForgeMintComment)
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
		it.Event = new(CoinbaseforgeTokenForgeMintComment)
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
func (it *CoinbaseforgeTokenForgeMintCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeTokenForgeMintCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeTokenForgeMintComment represents a TokenForgeMintComment event raised by the Coinbaseforge contract.
type CoinbaseforgeTokenForgeMintComment struct {
	To       common.Address
	Quantity *big.Int
	Comment  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenForgeMintComment is a free log retrieval operation binding the contract event 0x203498aecb28c99d51721d218d93e378293b86eacb26b42246dae394840ae756.
//
// Solidity: event TokenForgeMintComment(address indexed to, uint256 quantity, string comment)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterTokenForgeMintComment(opts *bind.FilterOpts, to []common.Address) (*CoinbaseforgeTokenForgeMintCommentIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "TokenForgeMintComment", toRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeTokenForgeMintCommentIterator{contract: _Coinbaseforge.contract, event: "TokenForgeMintComment", logs: logs, sub: sub}, nil
}

// WatchTokenForgeMintComment is a free log subscription operation binding the contract event 0x203498aecb28c99d51721d218d93e378293b86eacb26b42246dae394840ae756.
//
// Solidity: event TokenForgeMintComment(address indexed to, uint256 quantity, string comment)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchTokenForgeMintComment(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeTokenForgeMintComment, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "TokenForgeMintComment", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeTokenForgeMintComment)
				if err := _Coinbaseforge.contract.UnpackLog(event, "TokenForgeMintComment", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseTokenForgeMintComment(log types.Log) (*CoinbaseforgeTokenForgeMintComment, error) {
	event := new(CoinbaseforgeTokenForgeMintComment)
	if err := _Coinbaseforge.contract.UnpackLog(event, "TokenForgeMintComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Coinbaseforge contract.
type CoinbaseforgeTransferIterator struct {
	Event *CoinbaseforgeTransfer // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeTransfer)
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
		it.Event = new(CoinbaseforgeTransfer)
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
func (it *CoinbaseforgeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeTransfer represents a Transfer event raised by the Coinbaseforge contract.
type CoinbaseforgeTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*CoinbaseforgeTransferIterator, error) {

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

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeTransferIterator{contract: _Coinbaseforge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeTransfer)
				if err := _Coinbaseforge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseTransfer(log types.Log) (*CoinbaseforgeTransfer, error) {
	event := new(CoinbaseforgeTransfer)
	if err := _Coinbaseforge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinbaseforgeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Coinbaseforge contract.
type CoinbaseforgeWithdrawnIterator struct {
	Event *CoinbaseforgeWithdrawn // Event containing the contract specifics and raw log

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
func (it *CoinbaseforgeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinbaseforgeWithdrawn)
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
		it.Event = new(CoinbaseforgeWithdrawn)
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
func (it *CoinbaseforgeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinbaseforgeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinbaseforgeWithdrawn represents a Withdrawn event raised by the Coinbaseforge contract.
type CoinbaseforgeWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 amount)
func (_Coinbaseforge *CoinbaseforgeFilterer) FilterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*CoinbaseforgeWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.FilterLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CoinbaseforgeWithdrawnIterator{contract: _Coinbaseforge.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed owner, uint256 amount)
func (_Coinbaseforge *CoinbaseforgeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CoinbaseforgeWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Coinbaseforge.contract.WatchLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinbaseforgeWithdrawn)
				if err := _Coinbaseforge.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_Coinbaseforge *CoinbaseforgeFilterer) ParseWithdrawn(log types.Log) (*CoinbaseforgeWithdrawn, error) {
	event := new(CoinbaseforgeWithdrawn)
	if err := _Coinbaseforge.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
