// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basename

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

// IPriceOraclePrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceOraclePrice struct {
	Base    *big.Int
	Premium *big.Int
}

// RegistrarControllerDiscountDetails is an auto generated low-level Go binding around an user-defined struct.
type RegistrarControllerDiscountDetails struct {
	Active            bool
	DiscountValidator common.Address
	Key               [32]byte
	Discount          *big.Int
}

// RegistrarControllerRegisterRequest is an auto generated low-level Go binding around an user-defined struct.
type RegistrarControllerRegisterRequest struct {
	Name          string
	Owner         common.Address
	Duration      *big.Int
	Resolver      common.Address
	Data          [][]byte
	ReverseRecord bool
}

// BasenameMetaData contains all meta data concerning the Basename contract.
var BasenameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"base_\",\"type\":\"address\"},{\"internalType\":\"contractIPriceOracle\",\"name\":\"prices_\",\"type\":\"address\"},{\"internalType\":\"contractIReverseRegistrar\",\"name\":\"reverseRegistrar_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootNode_\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"rootName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AlreadyRegisteredWithDiscount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"InactiveDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InvalidDiscount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"InvalidDiscountAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResolverRequiredWhenDataSupplied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"discountKey\",\"type\":\"bytes32\"}],\"name\":\"DiscountApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"discountKey\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"discountValidator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRegistrarController.DiscountDetails\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"DiscountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ETHPaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPaymentReceiver\",\"type\":\"address\"}],\"name\":\"PaymentReceiverUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPrices\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReverseRegistrar\",\"type\":\"address\"}],\"name\":\"ReverseRegistrarUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"}],\"internalType\":\"structRegistrarController.RegisterRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"discountKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"validationData\",\"type\":\"bytes\"}],\"name\":\"discountedRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"discountKey\",\"type\":\"bytes32\"}],\"name\":\"discountedRegisterPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"}],\"name\":\"discountedRegistrants\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasRegisteredWithDiscount\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"discounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"discountValidator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveDiscounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"discountValidator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"internalType\":\"structRegistrarController.DiscountDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"hasRegisteredWithDiscount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"}],\"internalType\":\"structRegistrarController.RegisterRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"registerPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"internalType\":\"contractIReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"discountValidator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"internalType\":\"structRegistrarController.DiscountDetails\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"setDiscountDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"launchTime_\",\"type\":\"uint256\"}],\"name\":\"setLaunchTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentReceiver_\",\"type\":\"address\"}],\"name\":\"setPaymentReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"prices_\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIReverseRegistrar\",\"name\":\"reverse_\",\"type\":\"address\"}],\"name\":\"setReverseRegistrar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasenameABI is the input ABI used to generate the binding from.
// Deprecated: Use BasenameMetaData.ABI instead.
var BasenameABI = BasenameMetaData.ABI

// Basename is an auto generated Go binding around an Ethereum contract.
type Basename struct {
	BasenameCaller     // Read-only binding to the contract
	BasenameTransactor // Write-only binding to the contract
	BasenameFilterer   // Log filterer for contract events
}

// BasenameCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasenameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasenameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasenameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasenameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasenameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasenameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasenameSession struct {
	Contract     *Basename         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasenameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasenameCallerSession struct {
	Contract *BasenameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasenameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasenameTransactorSession struct {
	Contract     *BasenameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasenameRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasenameRaw struct {
	Contract *Basename // Generic contract binding to access the raw methods on
}

// BasenameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasenameCallerRaw struct {
	Contract *BasenameCaller // Generic read-only contract binding to access the raw methods on
}

// BasenameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasenameTransactorRaw struct {
	Contract *BasenameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasename creates a new instance of Basename, bound to a specific deployed contract.
func NewBasename(address common.Address, backend bind.ContractBackend) (*Basename, error) {
	contract, err := bindBasename(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Basename{BasenameCaller: BasenameCaller{contract: contract}, BasenameTransactor: BasenameTransactor{contract: contract}, BasenameFilterer: BasenameFilterer{contract: contract}}, nil
}

// NewBasenameCaller creates a new read-only instance of Basename, bound to a specific deployed contract.
func NewBasenameCaller(address common.Address, caller bind.ContractCaller) (*BasenameCaller, error) {
	contract, err := bindBasename(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasenameCaller{contract: contract}, nil
}

// NewBasenameTransactor creates a new write-only instance of Basename, bound to a specific deployed contract.
func NewBasenameTransactor(address common.Address, transactor bind.ContractTransactor) (*BasenameTransactor, error) {
	contract, err := bindBasename(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasenameTransactor{contract: contract}, nil
}

// NewBasenameFilterer creates a new log filterer instance of Basename, bound to a specific deployed contract.
func NewBasenameFilterer(address common.Address, filterer bind.ContractFilterer) (*BasenameFilterer, error) {
	contract, err := bindBasename(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasenameFilterer{contract: contract}, nil
}

// bindBasename binds a generic wrapper to an already deployed contract.
func bindBasename(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasenameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basename *BasenameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basename.Contract.BasenameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basename *BasenameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.Contract.BasenameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basename *BasenameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basename.Contract.BasenameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basename *BasenameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basename.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basename *BasenameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basename *BasenameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basename.Contract.contract.Transact(opts, method, params...)
}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Basename *BasenameCaller) MINNAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "MIN_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Basename *BasenameSession) MINNAMELENGTH() (*big.Int, error) {
	return _Basename.Contract.MINNAMELENGTH(&_Basename.CallOpts)
}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Basename *BasenameCallerSession) MINNAMELENGTH() (*big.Int, error) {
	return _Basename.Contract.MINNAMELENGTH(&_Basename.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Basename *BasenameCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Basename *BasenameSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Basename.Contract.MINREGISTRATIONDURATION(&_Basename.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Basename *BasenameCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Basename.Contract.MINREGISTRATIONDURATION(&_Basename.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Basename *BasenameCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Basename *BasenameSession) Available(name string) (bool, error) {
	return _Basename.Contract.Available(&_Basename.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Basename *BasenameCallerSession) Available(name string) (bool, error) {
	return _Basename.Contract.Available(&_Basename.CallOpts, name)
}

// DiscountedRegisterPrice is a free data retrieval call binding the contract method 0xedd7f501.
//
// Solidity: function discountedRegisterPrice(string name, uint256 duration, bytes32 discountKey) view returns(uint256 price)
func (_Basename *BasenameCaller) DiscountedRegisterPrice(opts *bind.CallOpts, name string, duration *big.Int, discountKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "discountedRegisterPrice", name, duration, discountKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiscountedRegisterPrice is a free data retrieval call binding the contract method 0xedd7f501.
//
// Solidity: function discountedRegisterPrice(string name, uint256 duration, bytes32 discountKey) view returns(uint256 price)
func (_Basename *BasenameSession) DiscountedRegisterPrice(name string, duration *big.Int, discountKey [32]byte) (*big.Int, error) {
	return _Basename.Contract.DiscountedRegisterPrice(&_Basename.CallOpts, name, duration, discountKey)
}

// DiscountedRegisterPrice is a free data retrieval call binding the contract method 0xedd7f501.
//
// Solidity: function discountedRegisterPrice(string name, uint256 duration, bytes32 discountKey) view returns(uint256 price)
func (_Basename *BasenameCallerSession) DiscountedRegisterPrice(name string, duration *big.Int, discountKey [32]byte) (*big.Int, error) {
	return _Basename.Contract.DiscountedRegisterPrice(&_Basename.CallOpts, name, duration, discountKey)
}

// DiscountedRegistrants is a free data retrieval call binding the contract method 0x06aa55b2.
//
// Solidity: function discountedRegistrants(address registrant) view returns(bool hasRegisteredWithDiscount)
func (_Basename *BasenameCaller) DiscountedRegistrants(opts *bind.CallOpts, registrant common.Address) (bool, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "discountedRegistrants", registrant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DiscountedRegistrants is a free data retrieval call binding the contract method 0x06aa55b2.
//
// Solidity: function discountedRegistrants(address registrant) view returns(bool hasRegisteredWithDiscount)
func (_Basename *BasenameSession) DiscountedRegistrants(registrant common.Address) (bool, error) {
	return _Basename.Contract.DiscountedRegistrants(&_Basename.CallOpts, registrant)
}

// DiscountedRegistrants is a free data retrieval call binding the contract method 0x06aa55b2.
//
// Solidity: function discountedRegistrants(address registrant) view returns(bool hasRegisteredWithDiscount)
func (_Basename *BasenameCallerSession) DiscountedRegistrants(registrant common.Address) (bool, error) {
	return _Basename.Contract.DiscountedRegistrants(&_Basename.CallOpts, registrant)
}

// Discounts is a free data retrieval call binding the contract method 0xbb3cc56f.
//
// Solidity: function discounts(bytes32 key) view returns(bool active, address discountValidator, bytes32 key, uint256 discount)
func (_Basename *BasenameCaller) Discounts(opts *bind.CallOpts, key [32]byte) (struct {
	Active            bool
	DiscountValidator common.Address
	Key               [32]byte
	Discount          *big.Int
}, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "discounts", key)

	outstruct := new(struct {
		Active            bool
		DiscountValidator common.Address
		Key               [32]byte
		Discount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DiscountValidator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Key = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Discount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Discounts is a free data retrieval call binding the contract method 0xbb3cc56f.
//
// Solidity: function discounts(bytes32 key) view returns(bool active, address discountValidator, bytes32 key, uint256 discount)
func (_Basename *BasenameSession) Discounts(key [32]byte) (struct {
	Active            bool
	DiscountValidator common.Address
	Key               [32]byte
	Discount          *big.Int
}, error) {
	return _Basename.Contract.Discounts(&_Basename.CallOpts, key)
}

// Discounts is a free data retrieval call binding the contract method 0xbb3cc56f.
//
// Solidity: function discounts(bytes32 key) view returns(bool active, address discountValidator, bytes32 key, uint256 discount)
func (_Basename *BasenameCallerSession) Discounts(key [32]byte) (struct {
	Active            bool
	DiscountValidator common.Address
	Key               [32]byte
	Discount          *big.Int
}, error) {
	return _Basename.Contract.Discounts(&_Basename.CallOpts, key)
}

// GetActiveDiscounts is a free data retrieval call binding the contract method 0xb053bc17.
//
// Solidity: function getActiveDiscounts() view returns((bool,address,bytes32,uint256)[])
func (_Basename *BasenameCaller) GetActiveDiscounts(opts *bind.CallOpts) ([]RegistrarControllerDiscountDetails, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "getActiveDiscounts")

	if err != nil {
		return *new([]RegistrarControllerDiscountDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegistrarControllerDiscountDetails)).(*[]RegistrarControllerDiscountDetails)

	return out0, err

}

// GetActiveDiscounts is a free data retrieval call binding the contract method 0xb053bc17.
//
// Solidity: function getActiveDiscounts() view returns((bool,address,bytes32,uint256)[])
func (_Basename *BasenameSession) GetActiveDiscounts() ([]RegistrarControllerDiscountDetails, error) {
	return _Basename.Contract.GetActiveDiscounts(&_Basename.CallOpts)
}

// GetActiveDiscounts is a free data retrieval call binding the contract method 0xb053bc17.
//
// Solidity: function getActiveDiscounts() view returns((bool,address,bytes32,uint256)[])
func (_Basename *BasenameCallerSession) GetActiveDiscounts() ([]RegistrarControllerDiscountDetails, error) {
	return _Basename.Contract.GetActiveDiscounts(&_Basename.CallOpts)
}

// HasRegisteredWithDiscount is a free data retrieval call binding the contract method 0x8e81d4f0.
//
// Solidity: function hasRegisteredWithDiscount(address[] addresses) view returns(bool)
func (_Basename *BasenameCaller) HasRegisteredWithDiscount(opts *bind.CallOpts, addresses []common.Address) (bool, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "hasRegisteredWithDiscount", addresses)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRegisteredWithDiscount is a free data retrieval call binding the contract method 0x8e81d4f0.
//
// Solidity: function hasRegisteredWithDiscount(address[] addresses) view returns(bool)
func (_Basename *BasenameSession) HasRegisteredWithDiscount(addresses []common.Address) (bool, error) {
	return _Basename.Contract.HasRegisteredWithDiscount(&_Basename.CallOpts, addresses)
}

// HasRegisteredWithDiscount is a free data retrieval call binding the contract method 0x8e81d4f0.
//
// Solidity: function hasRegisteredWithDiscount(address[] addresses) view returns(bool)
func (_Basename *BasenameCallerSession) HasRegisteredWithDiscount(addresses []common.Address) (bool, error) {
	return _Basename.Contract.HasRegisteredWithDiscount(&_Basename.CallOpts, addresses)
}

// LaunchTime is a free data retrieval call binding the contract method 0x790ca413.
//
// Solidity: function launchTime() view returns(uint256)
func (_Basename *BasenameCaller) LaunchTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "launchTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchTime is a free data retrieval call binding the contract method 0x790ca413.
//
// Solidity: function launchTime() view returns(uint256)
func (_Basename *BasenameSession) LaunchTime() (*big.Int, error) {
	return _Basename.Contract.LaunchTime(&_Basename.CallOpts)
}

// LaunchTime is a free data retrieval call binding the contract method 0x790ca413.
//
// Solidity: function launchTime() view returns(uint256)
func (_Basename *BasenameCallerSession) LaunchTime() (*big.Int, error) {
	return _Basename.Contract.LaunchTime(&_Basename.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Basename *BasenameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Basename *BasenameSession) Owner() (common.Address, error) {
	return _Basename.Contract.Owner(&_Basename.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Basename *BasenameCallerSession) Owner() (common.Address, error) {
	return _Basename.Contract.Owner(&_Basename.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Basename *BasenameCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Basename *BasenameSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Basename.Contract.OwnershipHandoverExpiresAt(&_Basename.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Basename *BasenameCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Basename.Contract.OwnershipHandoverExpiresAt(&_Basename.CallOpts, pendingOwner)
}

// PaymentReceiver is a free data retrieval call binding the contract method 0xcb37f3b2.
//
// Solidity: function paymentReceiver() view returns(address)
func (_Basename *BasenameCaller) PaymentReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "paymentReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentReceiver is a free data retrieval call binding the contract method 0xcb37f3b2.
//
// Solidity: function paymentReceiver() view returns(address)
func (_Basename *BasenameSession) PaymentReceiver() (common.Address, error) {
	return _Basename.Contract.PaymentReceiver(&_Basename.CallOpts)
}

// PaymentReceiver is a free data retrieval call binding the contract method 0xcb37f3b2.
//
// Solidity: function paymentReceiver() view returns(address)
func (_Basename *BasenameCallerSession) PaymentReceiver() (common.Address, error) {
	return _Basename.Contract.PaymentReceiver(&_Basename.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Basename *BasenameCaller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "prices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Basename *BasenameSession) Prices() (common.Address, error) {
	return _Basename.Contract.Prices(&_Basename.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Basename *BasenameCallerSession) Prices() (common.Address, error) {
	return _Basename.Contract.Prices(&_Basename.CallOpts)
}

// RegisterPrice is a free data retrieval call binding the contract method 0xe72c1e55.
//
// Solidity: function registerPrice(string name, uint256 duration) view returns(uint256)
func (_Basename *BasenameCaller) RegisterPrice(opts *bind.CallOpts, name string, duration *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "registerPrice", name, duration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegisterPrice is a free data retrieval call binding the contract method 0xe72c1e55.
//
// Solidity: function registerPrice(string name, uint256 duration) view returns(uint256)
func (_Basename *BasenameSession) RegisterPrice(name string, duration *big.Int) (*big.Int, error) {
	return _Basename.Contract.RegisterPrice(&_Basename.CallOpts, name, duration)
}

// RegisterPrice is a free data retrieval call binding the contract method 0xe72c1e55.
//
// Solidity: function registerPrice(string name, uint256 duration) view returns(uint256)
func (_Basename *BasenameCallerSession) RegisterPrice(name string, duration *big.Int) (*big.Int, error) {
	return _Basename.Contract.RegisterPrice(&_Basename.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Basename *BasenameCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(IPriceOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Basename *BasenameSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Basename.Contract.RentPrice(&_Basename.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Basename *BasenameCallerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Basename.Contract.RentPrice(&_Basename.CallOpts, name, duration)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Basename *BasenameCaller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "reverseRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Basename *BasenameSession) ReverseRegistrar() (common.Address, error) {
	return _Basename.Contract.ReverseRegistrar(&_Basename.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Basename *BasenameCallerSession) ReverseRegistrar() (common.Address, error) {
	return _Basename.Contract.ReverseRegistrar(&_Basename.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() view returns(string)
func (_Basename *BasenameCaller) RootName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "rootName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() view returns(string)
func (_Basename *BasenameSession) RootName() (string, error) {
	return _Basename.Contract.RootName(&_Basename.CallOpts)
}

// RootName is a free data retrieval call binding the contract method 0xf20387df.
//
// Solidity: function rootName() view returns(string)
func (_Basename *BasenameCallerSession) RootName() (string, error) {
	return _Basename.Contract.RootName(&_Basename.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() view returns(bytes32)
func (_Basename *BasenameCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "rootNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() view returns(bytes32)
func (_Basename *BasenameSession) RootNode() ([32]byte, error) {
	return _Basename.Contract.RootNode(&_Basename.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() view returns(bytes32)
func (_Basename *BasenameCallerSession) RootNode() ([32]byte, error) {
	return _Basename.Contract.RootNode(&_Basename.CallOpts)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Basename *BasenameCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Basename.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Basename *BasenameSession) Valid(name string) (bool, error) {
	return _Basename.Contract.Valid(&_Basename.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_Basename *BasenameCallerSession) Valid(name string) (bool, error) {
	return _Basename.Contract.Valid(&_Basename.CallOpts, name)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Basename *BasenameTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Basename *BasenameSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Basename.Contract.CancelOwnershipHandover(&_Basename.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Basename *BasenameTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Basename.Contract.CancelOwnershipHandover(&_Basename.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Basename *BasenameTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Basename *BasenameSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Basename.Contract.CompleteOwnershipHandover(&_Basename.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Basename *BasenameTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Basename.Contract.CompleteOwnershipHandover(&_Basename.TransactOpts, pendingOwner)
}

// DiscountedRegister is a paid mutator transaction binding the contract method 0xe0093eda.
//
// Solidity: function discountedRegister((string,address,uint256,address,bytes[],bool) request, bytes32 discountKey, bytes validationData) payable returns()
func (_Basename *BasenameTransactor) DiscountedRegister(opts *bind.TransactOpts, request RegistrarControllerRegisterRequest, discountKey [32]byte, validationData []byte) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "discountedRegister", request, discountKey, validationData)
}

// DiscountedRegister is a paid mutator transaction binding the contract method 0xe0093eda.
//
// Solidity: function discountedRegister((string,address,uint256,address,bytes[],bool) request, bytes32 discountKey, bytes validationData) payable returns()
func (_Basename *BasenameSession) DiscountedRegister(request RegistrarControllerRegisterRequest, discountKey [32]byte, validationData []byte) (*types.Transaction, error) {
	return _Basename.Contract.DiscountedRegister(&_Basename.TransactOpts, request, discountKey, validationData)
}

// DiscountedRegister is a paid mutator transaction binding the contract method 0xe0093eda.
//
// Solidity: function discountedRegister((string,address,uint256,address,bytes[],bool) request, bytes32 discountKey, bytes validationData) payable returns()
func (_Basename *BasenameTransactorSession) DiscountedRegister(request RegistrarControllerRegisterRequest, discountKey [32]byte, validationData []byte) (*types.Transaction, error) {
	return _Basename.Contract.DiscountedRegister(&_Basename.TransactOpts, request, discountKey, validationData)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Basename *BasenameTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Basename *BasenameSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.RecoverFunds(&_Basename.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Basename *BasenameTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.RecoverFunds(&_Basename.TransactOpts, _token, _to, _amount)
}

// Register is a paid mutator transaction binding the contract method 0xc7c79676.
//
// Solidity: function register((string,address,uint256,address,bytes[],bool) request) payable returns()
func (_Basename *BasenameTransactor) Register(opts *bind.TransactOpts, request RegistrarControllerRegisterRequest) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "register", request)
}

// Register is a paid mutator transaction binding the contract method 0xc7c79676.
//
// Solidity: function register((string,address,uint256,address,bytes[],bool) request) payable returns()
func (_Basename *BasenameSession) Register(request RegistrarControllerRegisterRequest) (*types.Transaction, error) {
	return _Basename.Contract.Register(&_Basename.TransactOpts, request)
}

// Register is a paid mutator transaction binding the contract method 0xc7c79676.
//
// Solidity: function register((string,address,uint256,address,bytes[],bool) request) payable returns()
func (_Basename *BasenameTransactorSession) Register(request RegistrarControllerRegisterRequest) (*types.Transaction, error) {
	return _Basename.Contract.Register(&_Basename.TransactOpts, request)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Basename *BasenameTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Basename *BasenameSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.Renew(&_Basename.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Basename *BasenameTransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.Renew(&_Basename.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Basename *BasenameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Basename *BasenameSession) RenounceOwnership() (*types.Transaction, error) {
	return _Basename.Contract.RenounceOwnership(&_Basename.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Basename *BasenameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Basename.Contract.RenounceOwnership(&_Basename.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Basename *BasenameTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Basename *BasenameSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Basename.Contract.RequestOwnershipHandover(&_Basename.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Basename *BasenameTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Basename.Contract.RequestOwnershipHandover(&_Basename.TransactOpts)
}

// SetDiscountDetails is a paid mutator transaction binding the contract method 0x682baa90.
//
// Solidity: function setDiscountDetails((bool,address,bytes32,uint256) details) returns()
func (_Basename *BasenameTransactor) SetDiscountDetails(opts *bind.TransactOpts, details RegistrarControllerDiscountDetails) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "setDiscountDetails", details)
}

// SetDiscountDetails is a paid mutator transaction binding the contract method 0x682baa90.
//
// Solidity: function setDiscountDetails((bool,address,bytes32,uint256) details) returns()
func (_Basename *BasenameSession) SetDiscountDetails(details RegistrarControllerDiscountDetails) (*types.Transaction, error) {
	return _Basename.Contract.SetDiscountDetails(&_Basename.TransactOpts, details)
}

// SetDiscountDetails is a paid mutator transaction binding the contract method 0x682baa90.
//
// Solidity: function setDiscountDetails((bool,address,bytes32,uint256) details) returns()
func (_Basename *BasenameTransactorSession) SetDiscountDetails(details RegistrarControllerDiscountDetails) (*types.Transaction, error) {
	return _Basename.Contract.SetDiscountDetails(&_Basename.TransactOpts, details)
}

// SetLaunchTime is a paid mutator transaction binding the contract method 0x9ff46e74.
//
// Solidity: function setLaunchTime(uint256 launchTime_) returns()
func (_Basename *BasenameTransactor) SetLaunchTime(opts *bind.TransactOpts, launchTime_ *big.Int) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "setLaunchTime", launchTime_)
}

// SetLaunchTime is a paid mutator transaction binding the contract method 0x9ff46e74.
//
// Solidity: function setLaunchTime(uint256 launchTime_) returns()
func (_Basename *BasenameSession) SetLaunchTime(launchTime_ *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.SetLaunchTime(&_Basename.TransactOpts, launchTime_)
}

// SetLaunchTime is a paid mutator transaction binding the contract method 0x9ff46e74.
//
// Solidity: function setLaunchTime(uint256 launchTime_) returns()
func (_Basename *BasenameTransactorSession) SetLaunchTime(launchTime_ *big.Int) (*types.Transaction, error) {
	return _Basename.Contract.SetLaunchTime(&_Basename.TransactOpts, launchTime_)
}

// SetPaymentReceiver is a paid mutator transaction binding the contract method 0x65ebf99a.
//
// Solidity: function setPaymentReceiver(address paymentReceiver_) returns()
func (_Basename *BasenameTransactor) SetPaymentReceiver(opts *bind.TransactOpts, paymentReceiver_ common.Address) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "setPaymentReceiver", paymentReceiver_)
}

// SetPaymentReceiver is a paid mutator transaction binding the contract method 0x65ebf99a.
//
// Solidity: function setPaymentReceiver(address paymentReceiver_) returns()
func (_Basename *BasenameSession) SetPaymentReceiver(paymentReceiver_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetPaymentReceiver(&_Basename.TransactOpts, paymentReceiver_)
}

// SetPaymentReceiver is a paid mutator transaction binding the contract method 0x65ebf99a.
//
// Solidity: function setPaymentReceiver(address paymentReceiver_) returns()
func (_Basename *BasenameTransactorSession) SetPaymentReceiver(paymentReceiver_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetPaymentReceiver(&_Basename.TransactOpts, paymentReceiver_)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address prices_) returns()
func (_Basename *BasenameTransactor) SetPriceOracle(opts *bind.TransactOpts, prices_ common.Address) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "setPriceOracle", prices_)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address prices_) returns()
func (_Basename *BasenameSession) SetPriceOracle(prices_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetPriceOracle(&_Basename.TransactOpts, prices_)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address prices_) returns()
func (_Basename *BasenameTransactorSession) SetPriceOracle(prices_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetPriceOracle(&_Basename.TransactOpts, prices_)
}

// SetReverseRegistrar is a paid mutator transaction binding the contract method 0x557499ba.
//
// Solidity: function setReverseRegistrar(address reverse_) returns()
func (_Basename *BasenameTransactor) SetReverseRegistrar(opts *bind.TransactOpts, reverse_ common.Address) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "setReverseRegistrar", reverse_)
}

// SetReverseRegistrar is a paid mutator transaction binding the contract method 0x557499ba.
//
// Solidity: function setReverseRegistrar(address reverse_) returns()
func (_Basename *BasenameSession) SetReverseRegistrar(reverse_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetReverseRegistrar(&_Basename.TransactOpts, reverse_)
}

// SetReverseRegistrar is a paid mutator transaction binding the contract method 0x557499ba.
//
// Solidity: function setReverseRegistrar(address reverse_) returns()
func (_Basename *BasenameTransactorSession) SetReverseRegistrar(reverse_ common.Address) (*types.Transaction, error) {
	return _Basename.Contract.SetReverseRegistrar(&_Basename.TransactOpts, reverse_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Basename *BasenameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Basename *BasenameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Basename.Contract.TransferOwnership(&_Basename.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Basename *BasenameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Basename.Contract.TransferOwnership(&_Basename.TransactOpts, newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Basename *BasenameTransactor) WithdrawETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basename.contract.Transact(opts, "withdrawETH")
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Basename *BasenameSession) WithdrawETH() (*types.Transaction, error) {
	return _Basename.Contract.WithdrawETH(&_Basename.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Basename *BasenameTransactorSession) WithdrawETH() (*types.Transaction, error) {
	return _Basename.Contract.WithdrawETH(&_Basename.TransactOpts)
}

// BasenameDiscountAppliedIterator is returned from FilterDiscountApplied and is used to iterate over the raw logs and unpacked data for DiscountApplied events raised by the Basename contract.
type BasenameDiscountAppliedIterator struct {
	Event *BasenameDiscountApplied // Event containing the contract specifics and raw log

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
func (it *BasenameDiscountAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameDiscountApplied)
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
		it.Event = new(BasenameDiscountApplied)
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
func (it *BasenameDiscountAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameDiscountAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameDiscountApplied represents a DiscountApplied event raised by the Basename contract.
type BasenameDiscountApplied struct {
	Registrant  common.Address
	DiscountKey [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDiscountApplied is a free log retrieval operation binding the contract event 0xfe82878a5987cea7129c337d7aaa6a49585236fc104b066223bc5b5e49510e2b.
//
// Solidity: event DiscountApplied(address indexed registrant, bytes32 indexed discountKey)
func (_Basename *BasenameFilterer) FilterDiscountApplied(opts *bind.FilterOpts, registrant []common.Address, discountKey [][32]byte) (*BasenameDiscountAppliedIterator, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var discountKeyRule []interface{}
	for _, discountKeyItem := range discountKey {
		discountKeyRule = append(discountKeyRule, discountKeyItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "DiscountApplied", registrantRule, discountKeyRule)
	if err != nil {
		return nil, err
	}
	return &BasenameDiscountAppliedIterator{contract: _Basename.contract, event: "DiscountApplied", logs: logs, sub: sub}, nil
}

// WatchDiscountApplied is a free log subscription operation binding the contract event 0xfe82878a5987cea7129c337d7aaa6a49585236fc104b066223bc5b5e49510e2b.
//
// Solidity: event DiscountApplied(address indexed registrant, bytes32 indexed discountKey)
func (_Basename *BasenameFilterer) WatchDiscountApplied(opts *bind.WatchOpts, sink chan<- *BasenameDiscountApplied, registrant []common.Address, discountKey [][32]byte) (event.Subscription, error) {

	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}
	var discountKeyRule []interface{}
	for _, discountKeyItem := range discountKey {
		discountKeyRule = append(discountKeyRule, discountKeyItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "DiscountApplied", registrantRule, discountKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameDiscountApplied)
				if err := _Basename.contract.UnpackLog(event, "DiscountApplied", log); err != nil {
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

// ParseDiscountApplied is a log parse operation binding the contract event 0xfe82878a5987cea7129c337d7aaa6a49585236fc104b066223bc5b5e49510e2b.
//
// Solidity: event DiscountApplied(address indexed registrant, bytes32 indexed discountKey)
func (_Basename *BasenameFilterer) ParseDiscountApplied(log types.Log) (*BasenameDiscountApplied, error) {
	event := new(BasenameDiscountApplied)
	if err := _Basename.contract.UnpackLog(event, "DiscountApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameDiscountUpdatedIterator is returned from FilterDiscountUpdated and is used to iterate over the raw logs and unpacked data for DiscountUpdated events raised by the Basename contract.
type BasenameDiscountUpdatedIterator struct {
	Event *BasenameDiscountUpdated // Event containing the contract specifics and raw log

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
func (it *BasenameDiscountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameDiscountUpdated)
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
		it.Event = new(BasenameDiscountUpdated)
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
func (it *BasenameDiscountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameDiscountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameDiscountUpdated represents a DiscountUpdated event raised by the Basename contract.
type BasenameDiscountUpdated struct {
	DiscountKey [32]byte
	Details     RegistrarControllerDiscountDetails
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDiscountUpdated is a free log retrieval operation binding the contract event 0x8e8593a48f626f59b50ae7022bc900a089280e27a89c5d7c1a4ca9a4338b47d1.
//
// Solidity: event DiscountUpdated(bytes32 indexed discountKey, (bool,address,bytes32,uint256) details)
func (_Basename *BasenameFilterer) FilterDiscountUpdated(opts *bind.FilterOpts, discountKey [][32]byte) (*BasenameDiscountUpdatedIterator, error) {

	var discountKeyRule []interface{}
	for _, discountKeyItem := range discountKey {
		discountKeyRule = append(discountKeyRule, discountKeyItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "DiscountUpdated", discountKeyRule)
	if err != nil {
		return nil, err
	}
	return &BasenameDiscountUpdatedIterator{contract: _Basename.contract, event: "DiscountUpdated", logs: logs, sub: sub}, nil
}

// WatchDiscountUpdated is a free log subscription operation binding the contract event 0x8e8593a48f626f59b50ae7022bc900a089280e27a89c5d7c1a4ca9a4338b47d1.
//
// Solidity: event DiscountUpdated(bytes32 indexed discountKey, (bool,address,bytes32,uint256) details)
func (_Basename *BasenameFilterer) WatchDiscountUpdated(opts *bind.WatchOpts, sink chan<- *BasenameDiscountUpdated, discountKey [][32]byte) (event.Subscription, error) {

	var discountKeyRule []interface{}
	for _, discountKeyItem := range discountKey {
		discountKeyRule = append(discountKeyRule, discountKeyItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "DiscountUpdated", discountKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameDiscountUpdated)
				if err := _Basename.contract.UnpackLog(event, "DiscountUpdated", log); err != nil {
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

// ParseDiscountUpdated is a log parse operation binding the contract event 0x8e8593a48f626f59b50ae7022bc900a089280e27a89c5d7c1a4ca9a4338b47d1.
//
// Solidity: event DiscountUpdated(bytes32 indexed discountKey, (bool,address,bytes32,uint256) details)
func (_Basename *BasenameFilterer) ParseDiscountUpdated(log types.Log) (*BasenameDiscountUpdated, error) {
	event := new(BasenameDiscountUpdated)
	if err := _Basename.contract.UnpackLog(event, "DiscountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameETHPaymentProcessedIterator is returned from FilterETHPaymentProcessed and is used to iterate over the raw logs and unpacked data for ETHPaymentProcessed events raised by the Basename contract.
type BasenameETHPaymentProcessedIterator struct {
	Event *BasenameETHPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *BasenameETHPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameETHPaymentProcessed)
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
		it.Event = new(BasenameETHPaymentProcessed)
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
func (it *BasenameETHPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameETHPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameETHPaymentProcessed represents a ETHPaymentProcessed event raised by the Basename contract.
type BasenameETHPaymentProcessed struct {
	Payee common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterETHPaymentProcessed is a free log retrieval operation binding the contract event 0xbc769889246686134856b409155bb87630ea5797a705fa98b61f576d316aab9b.
//
// Solidity: event ETHPaymentProcessed(address indexed payee, uint256 price)
func (_Basename *BasenameFilterer) FilterETHPaymentProcessed(opts *bind.FilterOpts, payee []common.Address) (*BasenameETHPaymentProcessedIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "ETHPaymentProcessed", payeeRule)
	if err != nil {
		return nil, err
	}
	return &BasenameETHPaymentProcessedIterator{contract: _Basename.contract, event: "ETHPaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchETHPaymentProcessed is a free log subscription operation binding the contract event 0xbc769889246686134856b409155bb87630ea5797a705fa98b61f576d316aab9b.
//
// Solidity: event ETHPaymentProcessed(address indexed payee, uint256 price)
func (_Basename *BasenameFilterer) WatchETHPaymentProcessed(opts *bind.WatchOpts, sink chan<- *BasenameETHPaymentProcessed, payee []common.Address) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "ETHPaymentProcessed", payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameETHPaymentProcessed)
				if err := _Basename.contract.UnpackLog(event, "ETHPaymentProcessed", log); err != nil {
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

// ParseETHPaymentProcessed is a log parse operation binding the contract event 0xbc769889246686134856b409155bb87630ea5797a705fa98b61f576d316aab9b.
//
// Solidity: event ETHPaymentProcessed(address indexed payee, uint256 price)
func (_Basename *BasenameFilterer) ParseETHPaymentProcessed(log types.Log) (*BasenameETHPaymentProcessed, error) {
	event := new(BasenameETHPaymentProcessed)
	if err := _Basename.contract.UnpackLog(event, "ETHPaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Basename contract.
type BasenameNameRegisteredIterator struct {
	Event *BasenameNameRegistered // Event containing the contract specifics and raw log

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
func (it *BasenameNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameNameRegistered)
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
		it.Event = new(BasenameNameRegistered)
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
func (it *BasenameNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameNameRegistered represents a NameRegistered event raised by the Basename contract.
type BasenameNameRegistered struct {
	Name    string
	Label   [32]byte
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x0667086d08417333ce63f40d5bc2ef6fd330e25aaaf317b7c489541f8fe600fa.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 expires)
func (_Basename *BasenameFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*BasenameNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BasenameNameRegisteredIterator{contract: _Basename.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x0667086d08417333ce63f40d5bc2ef6fd330e25aaaf317b7c489541f8fe600fa.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 expires)
func (_Basename *BasenameFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *BasenameNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameNameRegistered)
				if err := _Basename.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x0667086d08417333ce63f40d5bc2ef6fd330e25aaaf317b7c489541f8fe600fa.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 expires)
func (_Basename *BasenameFilterer) ParseNameRegistered(log types.Log) (*BasenameNameRegistered, error) {
	event := new(BasenameNameRegistered)
	if err := _Basename.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Basename contract.
type BasenameNameRenewedIterator struct {
	Event *BasenameNameRenewed // Event containing the contract specifics and raw log

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
func (it *BasenameNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameNameRenewed)
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
		it.Event = new(BasenameNameRenewed)
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
func (it *BasenameNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameNameRenewed represents a NameRenewed event raised by the Basename contract.
type BasenameNameRenewed struct {
	Name    string
	Label   [32]byte
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x93bc1a84707231b1d9552157299797c64a1a8c5bc79f05153716630c9c4936fc.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 expires)
func (_Basename *BasenameFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*BasenameNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &BasenameNameRenewedIterator{contract: _Basename.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x93bc1a84707231b1d9552157299797c64a1a8c5bc79f05153716630c9c4936fc.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 expires)
func (_Basename *BasenameFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *BasenameNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameNameRenewed)
				if err := _Basename.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x93bc1a84707231b1d9552157299797c64a1a8c5bc79f05153716630c9c4936fc.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 expires)
func (_Basename *BasenameFilterer) ParseNameRenewed(log types.Log) (*BasenameNameRenewed, error) {
	event := new(BasenameNameRenewed)
	if err := _Basename.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Basename contract.
type BasenameOwnershipHandoverCanceledIterator struct {
	Event *BasenameOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *BasenameOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameOwnershipHandoverCanceled)
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
		it.Event = new(BasenameOwnershipHandoverCanceled)
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
func (it *BasenameOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Basename contract.
type BasenameOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Basename *BasenameFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*BasenameOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasenameOwnershipHandoverCanceledIterator{contract: _Basename.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Basename *BasenameFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *BasenameOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameOwnershipHandoverCanceled)
				if err := _Basename.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Basename *BasenameFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*BasenameOwnershipHandoverCanceled, error) {
	event := new(BasenameOwnershipHandoverCanceled)
	if err := _Basename.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Basename contract.
type BasenameOwnershipHandoverRequestedIterator struct {
	Event *BasenameOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *BasenameOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameOwnershipHandoverRequested)
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
		it.Event = new(BasenameOwnershipHandoverRequested)
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
func (it *BasenameOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Basename contract.
type BasenameOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Basename *BasenameFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*BasenameOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasenameOwnershipHandoverRequestedIterator{contract: _Basename.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Basename *BasenameFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *BasenameOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameOwnershipHandoverRequested)
				if err := _Basename.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Basename *BasenameFilterer) ParseOwnershipHandoverRequested(log types.Log) (*BasenameOwnershipHandoverRequested, error) {
	event := new(BasenameOwnershipHandoverRequested)
	if err := _Basename.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Basename contract.
type BasenameOwnershipTransferredIterator struct {
	Event *BasenameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BasenameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameOwnershipTransferred)
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
		it.Event = new(BasenameOwnershipTransferred)
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
func (it *BasenameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameOwnershipTransferred represents a OwnershipTransferred event raised by the Basename contract.
type BasenameOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Basename *BasenameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BasenameOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Basename.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasenameOwnershipTransferredIterator{contract: _Basename.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Basename *BasenameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BasenameOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Basename.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameOwnershipTransferred)
				if err := _Basename.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Basename *BasenameFilterer) ParseOwnershipTransferred(log types.Log) (*BasenameOwnershipTransferred, error) {
	event := new(BasenameOwnershipTransferred)
	if err := _Basename.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenamePaymentReceiverUpdatedIterator is returned from FilterPaymentReceiverUpdated and is used to iterate over the raw logs and unpacked data for PaymentReceiverUpdated events raised by the Basename contract.
type BasenamePaymentReceiverUpdatedIterator struct {
	Event *BasenamePaymentReceiverUpdated // Event containing the contract specifics and raw log

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
func (it *BasenamePaymentReceiverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenamePaymentReceiverUpdated)
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
		it.Event = new(BasenamePaymentReceiverUpdated)
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
func (it *BasenamePaymentReceiverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenamePaymentReceiverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenamePaymentReceiverUpdated represents a PaymentReceiverUpdated event raised by the Basename contract.
type BasenamePaymentReceiverUpdated struct {
	NewPaymentReceiver common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceiverUpdated is a free log retrieval operation binding the contract event 0x6cfddd24d2afc1b9f31d51f0ef77029fdde044799f0a87a2a09b7673b6097422.
//
// Solidity: event PaymentReceiverUpdated(address newPaymentReceiver)
func (_Basename *BasenameFilterer) FilterPaymentReceiverUpdated(opts *bind.FilterOpts) (*BasenamePaymentReceiverUpdatedIterator, error) {

	logs, sub, err := _Basename.contract.FilterLogs(opts, "PaymentReceiverUpdated")
	if err != nil {
		return nil, err
	}
	return &BasenamePaymentReceiverUpdatedIterator{contract: _Basename.contract, event: "PaymentReceiverUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentReceiverUpdated is a free log subscription operation binding the contract event 0x6cfddd24d2afc1b9f31d51f0ef77029fdde044799f0a87a2a09b7673b6097422.
//
// Solidity: event PaymentReceiverUpdated(address newPaymentReceiver)
func (_Basename *BasenameFilterer) WatchPaymentReceiverUpdated(opts *bind.WatchOpts, sink chan<- *BasenamePaymentReceiverUpdated) (event.Subscription, error) {

	logs, sub, err := _Basename.contract.WatchLogs(opts, "PaymentReceiverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenamePaymentReceiverUpdated)
				if err := _Basename.contract.UnpackLog(event, "PaymentReceiverUpdated", log); err != nil {
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

// ParsePaymentReceiverUpdated is a log parse operation binding the contract event 0x6cfddd24d2afc1b9f31d51f0ef77029fdde044799f0a87a2a09b7673b6097422.
//
// Solidity: event PaymentReceiverUpdated(address newPaymentReceiver)
func (_Basename *BasenameFilterer) ParsePaymentReceiverUpdated(log types.Log) (*BasenamePaymentReceiverUpdated, error) {
	event := new(BasenamePaymentReceiverUpdated)
	if err := _Basename.contract.UnpackLog(event, "PaymentReceiverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenamePriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the Basename contract.
type BasenamePriceOracleUpdatedIterator struct {
	Event *BasenamePriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *BasenamePriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenamePriceOracleUpdated)
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
		it.Event = new(BasenamePriceOracleUpdated)
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
func (it *BasenamePriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenamePriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenamePriceOracleUpdated represents a PriceOracleUpdated event raised by the Basename contract.
type BasenamePriceOracleUpdated struct {
	NewPrices common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address newPrices)
func (_Basename *BasenameFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts) (*BasenamePriceOracleUpdatedIterator, error) {

	logs, sub, err := _Basename.contract.FilterLogs(opts, "PriceOracleUpdated")
	if err != nil {
		return nil, err
	}
	return &BasenamePriceOracleUpdatedIterator{contract: _Basename.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address newPrices)
func (_Basename *BasenameFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *BasenamePriceOracleUpdated) (event.Subscription, error) {

	logs, sub, err := _Basename.contract.WatchLogs(opts, "PriceOracleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenamePriceOracleUpdated)
				if err := _Basename.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address newPrices)
func (_Basename *BasenameFilterer) ParsePriceOracleUpdated(log types.Log) (*BasenamePriceOracleUpdated, error) {
	event := new(BasenamePriceOracleUpdated)
	if err := _Basename.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasenameReverseRegistrarUpdatedIterator is returned from FilterReverseRegistrarUpdated and is used to iterate over the raw logs and unpacked data for ReverseRegistrarUpdated events raised by the Basename contract.
type BasenameReverseRegistrarUpdatedIterator struct {
	Event *BasenameReverseRegistrarUpdated // Event containing the contract specifics and raw log

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
func (it *BasenameReverseRegistrarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasenameReverseRegistrarUpdated)
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
		it.Event = new(BasenameReverseRegistrarUpdated)
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
func (it *BasenameReverseRegistrarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasenameReverseRegistrarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasenameReverseRegistrarUpdated represents a ReverseRegistrarUpdated event raised by the Basename contract.
type BasenameReverseRegistrarUpdated struct {
	NewReverseRegistrar common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReverseRegistrarUpdated is a free log retrieval operation binding the contract event 0xd192c0b229b00473ccb6ccfebf6642805bca1dcdf2d9fb4fd102c7dc7ea4ce23.
//
// Solidity: event ReverseRegistrarUpdated(address newReverseRegistrar)
func (_Basename *BasenameFilterer) FilterReverseRegistrarUpdated(opts *bind.FilterOpts) (*BasenameReverseRegistrarUpdatedIterator, error) {

	logs, sub, err := _Basename.contract.FilterLogs(opts, "ReverseRegistrarUpdated")
	if err != nil {
		return nil, err
	}
	return &BasenameReverseRegistrarUpdatedIterator{contract: _Basename.contract, event: "ReverseRegistrarUpdated", logs: logs, sub: sub}, nil
}

// WatchReverseRegistrarUpdated is a free log subscription operation binding the contract event 0xd192c0b229b00473ccb6ccfebf6642805bca1dcdf2d9fb4fd102c7dc7ea4ce23.
//
// Solidity: event ReverseRegistrarUpdated(address newReverseRegistrar)
func (_Basename *BasenameFilterer) WatchReverseRegistrarUpdated(opts *bind.WatchOpts, sink chan<- *BasenameReverseRegistrarUpdated) (event.Subscription, error) {

	logs, sub, err := _Basename.contract.WatchLogs(opts, "ReverseRegistrarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasenameReverseRegistrarUpdated)
				if err := _Basename.contract.UnpackLog(event, "ReverseRegistrarUpdated", log); err != nil {
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

// ParseReverseRegistrarUpdated is a log parse operation binding the contract event 0xd192c0b229b00473ccb6ccfebf6642805bca1dcdf2d9fb4fd102c7dc7ea4ce23.
//
// Solidity: event ReverseRegistrarUpdated(address newReverseRegistrar)
func (_Basename *BasenameFilterer) ParseReverseRegistrarUpdated(log types.Log) (*BasenameReverseRegistrarUpdated, error) {
	event := new(BasenameReverseRegistrarUpdated)
	if err := _Basename.contract.UnpackLog(event, "ReverseRegistrarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
