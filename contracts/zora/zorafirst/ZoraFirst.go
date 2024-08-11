// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zorafirst

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

// ICreatorCommandsCommand is an auto generated low-level Go binding around an user-defined struct.
type ICreatorCommandsCommand struct {
	Method uint8
	Args   []byte
}

// ICreatorCommandsCommandSet is an auto generated low-level Go binding around an user-defined struct.
type ICreatorCommandsCommandSet struct {
	Commands []ICreatorCommandsCommand
	At       *big.Int
}

// IZoraTimedSaleStrategyERC20zActivate is an auto generated low-level Go binding around an user-defined struct.
type IZoraTimedSaleStrategyERC20zActivate struct {
	FinalTotalERC20ZSupply  *big.Int
	Erc20Reserve            *big.Int
	Erc20Liquidity          *big.Int
	ExcessERC20             *big.Int
	ExcessERC1155           *big.Int
	AdditionalERC1155ToMint *big.Int
	Final1155Supply         *big.Int
}

// IZoraTimedSaleStrategyRewardsSettings is an auto generated low-level Go binding around an user-defined struct.
type IZoraTimedSaleStrategyRewardsSettings struct {
	TotalReward          *big.Int
	CreatorReward        *big.Int
	CreateReferralReward *big.Int
	MintReferralReward   *big.Int
	MarketReward         *big.Int
	ZoraReward           *big.Int
}

// IZoraTimedSaleStrategySaleStorage is an auto generated low-level Go binding around an user-defined struct.
type IZoraTimedSaleStrategySaleStorage struct {
	Erc20zAddress      common.Address
	SaleStart          uint64
	PoolAddress        common.Address
	SaleEnd            uint64
	SecondaryActivated bool
}

// IZoraTimedSaleStrategySalesConfig is an auto generated low-level Go binding around an user-defined struct.
type IZoraTimedSaleStrategySalesConfig struct {
	SaleStart uint64
	SaleEnd   uint64
	Name      string
	Symbol    string
}

// ZorafirstMetaData contains all meta data concerning the Zorafirst contract.
var ZorafirstMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1167FailedCreateClone\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EndTimeCannotBeInThePast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketAlreadyLaunched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NeedsToBeAtLeastOneSaleToStartMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyZoraRewardRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestMintInvalidUseMint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResetSaleNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleHasNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartTimeCannotBeAfterEndTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongValueSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZoraCreator1155ContractNeedsToSupportReduceSupply\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20zAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"MarketLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"MintComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"saleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleEnd\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIZoraTimedSaleStrategy.SalesConfig\",\"name\":\"salesConfig\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20zAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFee\",\"type\":\"uint256\"}],\"name\":\"SaleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"ZoraRewardRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zoraRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"name\":\"ZoraTimedSaleStrategyRewards\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20zAddress\",\"type\":\"address\"}],\"name\":\"calculateERC20zActivate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"finalTotalERC20ZSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"erc20Reserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"erc20Liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessERC20\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"excessERC1155\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalERC1155ToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"final1155Supply\",\"type\":\"uint256\"}],\"internalType\":\"structIZoraTimedSaleStrategy.ERC20zActivate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"computeRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"internalType\":\"structIZoraTimedSaleStrategy.RewardsSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20zImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCreateReferral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zoraRewardRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20zImpl\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolRewards\",\"name\":\"_protocolRewards\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"launchMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolRewards\",\"outputs\":[{\"internalType\":\"contractIProtocolRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"requestMint\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumICreatorCommands.CreatorActions\",\"name\":\"method\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"internalType\":\"structICreatorCommands.Command[]\",\"name\":\"commands\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"}],\"internalType\":\"structICreatorCommands.CommandSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"sale\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"erc20zAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"saleStart\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"saleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"secondaryActivated\",\"type\":\"bool\"}],\"internalType\":\"structIZoraTimedSaleStrategy.SaleStorage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"saleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleEnd\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structIZoraTimedSaleStrategy.SalesConfig\",\"name\":\"salesConfig\",\"type\":\"tuple\"}],\"name\":\"setSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setZoraRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"newStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newEndTime\",\"type\":\"uint64\"}],\"name\":\"updateSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ZorafirstABI is the input ABI used to generate the binding from.
// Deprecated: Use ZorafirstMetaData.ABI instead.
var ZorafirstABI = ZorafirstMetaData.ABI

// Zorafirst is an auto generated Go binding around an Ethereum contract.
type Zorafirst struct {
	ZorafirstCaller     // Read-only binding to the contract
	ZorafirstTransactor // Write-only binding to the contract
	ZorafirstFilterer   // Log filterer for contract events
}

// ZorafirstCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZorafirstCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorafirstTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZorafirstTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorafirstFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZorafirstFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorafirstSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZorafirstSession struct {
	Contract     *Zorafirst        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZorafirstCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZorafirstCallerSession struct {
	Contract *ZorafirstCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZorafirstTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZorafirstTransactorSession struct {
	Contract     *ZorafirstTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZorafirstRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZorafirstRaw struct {
	Contract *Zorafirst // Generic contract binding to access the raw methods on
}

// ZorafirstCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZorafirstCallerRaw struct {
	Contract *ZorafirstCaller // Generic read-only contract binding to access the raw methods on
}

// ZorafirstTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZorafirstTransactorRaw struct {
	Contract *ZorafirstTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZorafirst creates a new instance of Zorafirst, bound to a specific deployed contract.
func NewZorafirst(address common.Address, backend bind.ContractBackend) (*Zorafirst, error) {
	contract, err := bindZorafirst(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zorafirst{ZorafirstCaller: ZorafirstCaller{contract: contract}, ZorafirstTransactor: ZorafirstTransactor{contract: contract}, ZorafirstFilterer: ZorafirstFilterer{contract: contract}}, nil
}

// NewZorafirstCaller creates a new read-only instance of Zorafirst, bound to a specific deployed contract.
func NewZorafirstCaller(address common.Address, caller bind.ContractCaller) (*ZorafirstCaller, error) {
	contract, err := bindZorafirst(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZorafirstCaller{contract: contract}, nil
}

// NewZorafirstTransactor creates a new write-only instance of Zorafirst, bound to a specific deployed contract.
func NewZorafirstTransactor(address common.Address, transactor bind.ContractTransactor) (*ZorafirstTransactor, error) {
	contract, err := bindZorafirst(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZorafirstTransactor{contract: contract}, nil
}

// NewZorafirstFilterer creates a new log filterer instance of Zorafirst, bound to a specific deployed contract.
func NewZorafirstFilterer(address common.Address, filterer bind.ContractFilterer) (*ZorafirstFilterer, error) {
	contract, err := bindZorafirst(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZorafirstFilterer{contract: contract}, nil
}

// bindZorafirst binds a generic wrapper to an already deployed contract.
func bindZorafirst(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZorafirstMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zorafirst *ZorafirstRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zorafirst.Contract.ZorafirstCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zorafirst *ZorafirstRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorafirst.Contract.ZorafirstTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zorafirst *ZorafirstRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zorafirst.Contract.ZorafirstTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zorafirst *ZorafirstCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zorafirst.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zorafirst *ZorafirstTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorafirst.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zorafirst *ZorafirstTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zorafirst.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Zorafirst *ZorafirstCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Zorafirst *ZorafirstSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Zorafirst.Contract.UPGRADEINTERFACEVERSION(&_Zorafirst.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Zorafirst *ZorafirstCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Zorafirst.Contract.UPGRADEINTERFACEVERSION(&_Zorafirst.CallOpts)
}

// CalculateERC20zActivate is a free data retrieval call binding the contract method 0xd8d474a0.
//
// Solidity: function calculateERC20zActivate(address collection, uint256 tokenId, address erc20zAddress) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstCaller) CalculateERC20zActivate(opts *bind.CallOpts, collection common.Address, tokenId *big.Int, erc20zAddress common.Address) (IZoraTimedSaleStrategyERC20zActivate, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "calculateERC20zActivate", collection, tokenId, erc20zAddress)

	if err != nil {
		return *new(IZoraTimedSaleStrategyERC20zActivate), err
	}

	out0 := *abi.ConvertType(out[0], new(IZoraTimedSaleStrategyERC20zActivate)).(*IZoraTimedSaleStrategyERC20zActivate)

	return out0, err

}

// CalculateERC20zActivate is a free data retrieval call binding the contract method 0xd8d474a0.
//
// Solidity: function calculateERC20zActivate(address collection, uint256 tokenId, address erc20zAddress) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstSession) CalculateERC20zActivate(collection common.Address, tokenId *big.Int, erc20zAddress common.Address) (IZoraTimedSaleStrategyERC20zActivate, error) {
	return _Zorafirst.Contract.CalculateERC20zActivate(&_Zorafirst.CallOpts, collection, tokenId, erc20zAddress)
}

// CalculateERC20zActivate is a free data retrieval call binding the contract method 0xd8d474a0.
//
// Solidity: function calculateERC20zActivate(address collection, uint256 tokenId, address erc20zAddress) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstCallerSession) CalculateERC20zActivate(collection common.Address, tokenId *big.Int, erc20zAddress common.Address) (IZoraTimedSaleStrategyERC20zActivate, error) {
	return _Zorafirst.Contract.CalculateERC20zActivate(&_Zorafirst.CallOpts, collection, tokenId, erc20zAddress)
}

// ComputeRewards is a free data retrieval call binding the contract method 0x7a195ce0.
//
// Solidity: function computeRewards(uint256 quantity) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstCaller) ComputeRewards(opts *bind.CallOpts, quantity *big.Int) (IZoraTimedSaleStrategyRewardsSettings, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "computeRewards", quantity)

	if err != nil {
		return *new(IZoraTimedSaleStrategyRewardsSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(IZoraTimedSaleStrategyRewardsSettings)).(*IZoraTimedSaleStrategyRewardsSettings)

	return out0, err

}

// ComputeRewards is a free data retrieval call binding the contract method 0x7a195ce0.
//
// Solidity: function computeRewards(uint256 quantity) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstSession) ComputeRewards(quantity *big.Int) (IZoraTimedSaleStrategyRewardsSettings, error) {
	return _Zorafirst.Contract.ComputeRewards(&_Zorafirst.CallOpts, quantity)
}

// ComputeRewards is a free data retrieval call binding the contract method 0x7a195ce0.
//
// Solidity: function computeRewards(uint256 quantity) pure returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Zorafirst *ZorafirstCallerSession) ComputeRewards(quantity *big.Int) (IZoraTimedSaleStrategyRewardsSettings, error) {
	return _Zorafirst.Contract.ComputeRewards(&_Zorafirst.CallOpts, quantity)
}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() pure returns(string)
func (_Zorafirst *ZorafirstCaller) ContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "contractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() pure returns(string)
func (_Zorafirst *ZorafirstSession) ContractName() (string, error) {
	return _Zorafirst.Contract.ContractName(&_Zorafirst.CallOpts)
}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() pure returns(string)
func (_Zorafirst *ZorafirstCallerSession) ContractName() (string, error) {
	return _Zorafirst.Contract.ContractName(&_Zorafirst.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Zorafirst *ZorafirstCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Zorafirst *ZorafirstSession) ContractURI() (string, error) {
	return _Zorafirst.Contract.ContractURI(&_Zorafirst.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Zorafirst *ZorafirstCallerSession) ContractURI() (string, error) {
	return _Zorafirst.Contract.ContractURI(&_Zorafirst.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorafirst *ZorafirstCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorafirst *ZorafirstSession) ContractVersion() (string, error) {
	return _Zorafirst.Contract.ContractVersion(&_Zorafirst.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorafirst *ZorafirstCallerSession) ContractVersion() (string, error) {
	return _Zorafirst.Contract.ContractVersion(&_Zorafirst.CallOpts)
}

// Erc20zImpl is a free data retrieval call binding the contract method 0x7784cc4f.
//
// Solidity: function erc20zImpl() view returns(address)
func (_Zorafirst *ZorafirstCaller) Erc20zImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "erc20zImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20zImpl is a free data retrieval call binding the contract method 0x7784cc4f.
//
// Solidity: function erc20zImpl() view returns(address)
func (_Zorafirst *ZorafirstSession) Erc20zImpl() (common.Address, error) {
	return _Zorafirst.Contract.Erc20zImpl(&_Zorafirst.CallOpts)
}

// Erc20zImpl is a free data retrieval call binding the contract method 0x7784cc4f.
//
// Solidity: function erc20zImpl() view returns(address)
func (_Zorafirst *ZorafirstCallerSession) Erc20zImpl() (common.Address, error) {
	return _Zorafirst.Contract.Erc20zImpl(&_Zorafirst.CallOpts)
}

// GetCreateReferral is a free data retrieval call binding the contract method 0x5d759bd8.
//
// Solidity: function getCreateReferral(address collection, uint256 tokenId) view returns(address createReferral)
func (_Zorafirst *ZorafirstCaller) GetCreateReferral(opts *bind.CallOpts, collection common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "getCreateReferral", collection, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreateReferral is a free data retrieval call binding the contract method 0x5d759bd8.
//
// Solidity: function getCreateReferral(address collection, uint256 tokenId) view returns(address createReferral)
func (_Zorafirst *ZorafirstSession) GetCreateReferral(collection common.Address, tokenId *big.Int) (common.Address, error) {
	return _Zorafirst.Contract.GetCreateReferral(&_Zorafirst.CallOpts, collection, tokenId)
}

// GetCreateReferral is a free data retrieval call binding the contract method 0x5d759bd8.
//
// Solidity: function getCreateReferral(address collection, uint256 tokenId) view returns(address createReferral)
func (_Zorafirst *ZorafirstCallerSession) GetCreateReferral(collection common.Address, tokenId *big.Int) (common.Address, error) {
	return _Zorafirst.Contract.GetCreateReferral(&_Zorafirst.CallOpts, collection, tokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorafirst *ZorafirstCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorafirst *ZorafirstSession) Owner() (common.Address, error) {
	return _Zorafirst.Contract.Owner(&_Zorafirst.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorafirst *ZorafirstCallerSession) Owner() (common.Address, error) {
	return _Zorafirst.Contract.Owner(&_Zorafirst.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Zorafirst *ZorafirstCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Zorafirst *ZorafirstSession) PendingOwner() (common.Address, error) {
	return _Zorafirst.Contract.PendingOwner(&_Zorafirst.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Zorafirst *ZorafirstCallerSession) PendingOwner() (common.Address, error) {
	return _Zorafirst.Contract.PendingOwner(&_Zorafirst.CallOpts)
}

// ProtocolRewards is a free data retrieval call binding the contract method 0x29df6479.
//
// Solidity: function protocolRewards() view returns(address)
func (_Zorafirst *ZorafirstCaller) ProtocolRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "protocolRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolRewards is a free data retrieval call binding the contract method 0x29df6479.
//
// Solidity: function protocolRewards() view returns(address)
func (_Zorafirst *ZorafirstSession) ProtocolRewards() (common.Address, error) {
	return _Zorafirst.Contract.ProtocolRewards(&_Zorafirst.CallOpts)
}

// ProtocolRewards is a free data retrieval call binding the contract method 0x29df6479.
//
// Solidity: function protocolRewards() view returns(address)
func (_Zorafirst *ZorafirstCallerSession) ProtocolRewards() (common.Address, error) {
	return _Zorafirst.Contract.ProtocolRewards(&_Zorafirst.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorafirst *ZorafirstCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorafirst *ZorafirstSession) ProxiableUUID() ([32]byte, error) {
	return _Zorafirst.Contract.ProxiableUUID(&_Zorafirst.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorafirst *ZorafirstCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Zorafirst.Contract.ProxiableUUID(&_Zorafirst.CallOpts)
}

// RequestMint is a free data retrieval call binding the contract method 0x6890e5b3.
//
// Solidity: function requestMint(address , uint256 , uint256 , uint256 , bytes ) pure returns(((uint8,bytes)[],uint256))
func (_Zorafirst *ZorafirstCaller) RequestMint(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (ICreatorCommandsCommandSet, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "requestMint", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(ICreatorCommandsCommandSet), err
	}

	out0 := *abi.ConvertType(out[0], new(ICreatorCommandsCommandSet)).(*ICreatorCommandsCommandSet)

	return out0, err

}

// RequestMint is a free data retrieval call binding the contract method 0x6890e5b3.
//
// Solidity: function requestMint(address , uint256 , uint256 , uint256 , bytes ) pure returns(((uint8,bytes)[],uint256))
func (_Zorafirst *ZorafirstSession) RequestMint(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (ICreatorCommandsCommandSet, error) {
	return _Zorafirst.Contract.RequestMint(&_Zorafirst.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// RequestMint is a free data retrieval call binding the contract method 0x6890e5b3.
//
// Solidity: function requestMint(address , uint256 , uint256 , uint256 , bytes ) pure returns(((uint8,bytes)[],uint256))
func (_Zorafirst *ZorafirstCallerSession) RequestMint(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (ICreatorCommandsCommandSet, error) {
	return _Zorafirst.Contract.RequestMint(&_Zorafirst.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Sale is a free data retrieval call binding the contract method 0x611efc09.
//
// Solidity: function sale(address collection, uint256 tokenId) view returns((address,uint64,address,uint64,bool))
func (_Zorafirst *ZorafirstCaller) Sale(opts *bind.CallOpts, collection common.Address, tokenId *big.Int) (IZoraTimedSaleStrategySaleStorage, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "sale", collection, tokenId)

	if err != nil {
		return *new(IZoraTimedSaleStrategySaleStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(IZoraTimedSaleStrategySaleStorage)).(*IZoraTimedSaleStrategySaleStorage)

	return out0, err

}

// Sale is a free data retrieval call binding the contract method 0x611efc09.
//
// Solidity: function sale(address collection, uint256 tokenId) view returns((address,uint64,address,uint64,bool))
func (_Zorafirst *ZorafirstSession) Sale(collection common.Address, tokenId *big.Int) (IZoraTimedSaleStrategySaleStorage, error) {
	return _Zorafirst.Contract.Sale(&_Zorafirst.CallOpts, collection, tokenId)
}

// Sale is a free data retrieval call binding the contract method 0x611efc09.
//
// Solidity: function sale(address collection, uint256 tokenId) view returns((address,uint64,address,uint64,bool))
func (_Zorafirst *ZorafirstCallerSession) Sale(collection common.Address, tokenId *big.Int) (IZoraTimedSaleStrategySaleStorage, error) {
	return _Zorafirst.Contract.Sale(&_Zorafirst.CallOpts, collection, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Zorafirst *ZorafirstCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Zorafirst.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Zorafirst *ZorafirstSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zorafirst.Contract.SupportsInterface(&_Zorafirst.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Zorafirst *ZorafirstCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zorafirst.Contract.SupportsInterface(&_Zorafirst.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Zorafirst *ZorafirstTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Zorafirst *ZorafirstSession) AcceptOwnership() (*types.Transaction, error) {
	return _Zorafirst.Contract.AcceptOwnership(&_Zorafirst.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Zorafirst *ZorafirstTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Zorafirst.Contract.AcceptOwnership(&_Zorafirst.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultOwner, address _zoraRewardRecipient, address _erc20zImpl, address _protocolRewards) returns()
func (_Zorafirst *ZorafirstTransactor) Initialize(opts *bind.TransactOpts, _defaultOwner common.Address, _zoraRewardRecipient common.Address, _erc20zImpl common.Address, _protocolRewards common.Address) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "initialize", _defaultOwner, _zoraRewardRecipient, _erc20zImpl, _protocolRewards)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultOwner, address _zoraRewardRecipient, address _erc20zImpl, address _protocolRewards) returns()
func (_Zorafirst *ZorafirstSession) Initialize(_defaultOwner common.Address, _zoraRewardRecipient common.Address, _erc20zImpl common.Address, _protocolRewards common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.Initialize(&_Zorafirst.TransactOpts, _defaultOwner, _zoraRewardRecipient, _erc20zImpl, _protocolRewards)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultOwner, address _zoraRewardRecipient, address _erc20zImpl, address _protocolRewards) returns()
func (_Zorafirst *ZorafirstTransactorSession) Initialize(_defaultOwner common.Address, _zoraRewardRecipient common.Address, _erc20zImpl common.Address, _protocolRewards common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.Initialize(&_Zorafirst.TransactOpts, _defaultOwner, _zoraRewardRecipient, _erc20zImpl, _protocolRewards)
}

// LaunchMarket is a paid mutator transaction binding the contract method 0x3f4e74b7.
//
// Solidity: function launchMarket(address collection, uint256 tokenId) returns()
func (_Zorafirst *ZorafirstTransactor) LaunchMarket(opts *bind.TransactOpts, collection common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "launchMarket", collection, tokenId)
}

// LaunchMarket is a paid mutator transaction binding the contract method 0x3f4e74b7.
//
// Solidity: function launchMarket(address collection, uint256 tokenId) returns()
func (_Zorafirst *ZorafirstSession) LaunchMarket(collection common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zorafirst.Contract.LaunchMarket(&_Zorafirst.TransactOpts, collection, tokenId)
}

// LaunchMarket is a paid mutator transaction binding the contract method 0x3f4e74b7.
//
// Solidity: function launchMarket(address collection, uint256 tokenId) returns()
func (_Zorafirst *ZorafirstTransactorSession) LaunchMarket(collection common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Zorafirst.Contract.LaunchMarket(&_Zorafirst.TransactOpts, collection, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa836f32f.
//
// Solidity: function mint(address mintTo, uint256 quantity, address collection, uint256 tokenId, address mintReferral, string comment) payable returns()
func (_Zorafirst *ZorafirstTransactor) Mint(opts *bind.TransactOpts, mintTo common.Address, quantity *big.Int, collection common.Address, tokenId *big.Int, mintReferral common.Address, comment string) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "mint", mintTo, quantity, collection, tokenId, mintReferral, comment)
}

// Mint is a paid mutator transaction binding the contract method 0xa836f32f.
//
// Solidity: function mint(address mintTo, uint256 quantity, address collection, uint256 tokenId, address mintReferral, string comment) payable returns()
func (_Zorafirst *ZorafirstSession) Mint(mintTo common.Address, quantity *big.Int, collection common.Address, tokenId *big.Int, mintReferral common.Address, comment string) (*types.Transaction, error) {
	return _Zorafirst.Contract.Mint(&_Zorafirst.TransactOpts, mintTo, quantity, collection, tokenId, mintReferral, comment)
}

// Mint is a paid mutator transaction binding the contract method 0xa836f32f.
//
// Solidity: function mint(address mintTo, uint256 quantity, address collection, uint256 tokenId, address mintReferral, string comment) payable returns()
func (_Zorafirst *ZorafirstTransactorSession) Mint(mintTo common.Address, quantity *big.Int, collection common.Address, tokenId *big.Int, mintReferral common.Address, comment string) (*types.Transaction, error) {
	return _Zorafirst.Contract.Mint(&_Zorafirst.TransactOpts, mintTo, quantity, collection, tokenId, mintReferral, comment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zorafirst *ZorafirstTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zorafirst *ZorafirstSession) RenounceOwnership() (*types.Transaction, error) {
	return _Zorafirst.Contract.RenounceOwnership(&_Zorafirst.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Zorafirst *ZorafirstTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Zorafirst.Contract.RenounceOwnership(&_Zorafirst.TransactOpts)
}

// SetSale is a paid mutator transaction binding the contract method 0x642436e3.
//
// Solidity: function setSale(uint256 tokenId, (uint64,uint64,string,string) salesConfig) returns()
func (_Zorafirst *ZorafirstTransactor) SetSale(opts *bind.TransactOpts, tokenId *big.Int, salesConfig IZoraTimedSaleStrategySalesConfig) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "setSale", tokenId, salesConfig)
}

// SetSale is a paid mutator transaction binding the contract method 0x642436e3.
//
// Solidity: function setSale(uint256 tokenId, (uint64,uint64,string,string) salesConfig) returns()
func (_Zorafirst *ZorafirstSession) SetSale(tokenId *big.Int, salesConfig IZoraTimedSaleStrategySalesConfig) (*types.Transaction, error) {
	return _Zorafirst.Contract.SetSale(&_Zorafirst.TransactOpts, tokenId, salesConfig)
}

// SetSale is a paid mutator transaction binding the contract method 0x642436e3.
//
// Solidity: function setSale(uint256 tokenId, (uint64,uint64,string,string) salesConfig) returns()
func (_Zorafirst *ZorafirstTransactorSession) SetSale(tokenId *big.Int, salesConfig IZoraTimedSaleStrategySalesConfig) (*types.Transaction, error) {
	return _Zorafirst.Contract.SetSale(&_Zorafirst.TransactOpts, tokenId, salesConfig)
}

// SetZoraRewardRecipient is a paid mutator transaction binding the contract method 0x0deab73c.
//
// Solidity: function setZoraRewardRecipient(address recipient) returns()
func (_Zorafirst *ZorafirstTransactor) SetZoraRewardRecipient(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "setZoraRewardRecipient", recipient)
}

// SetZoraRewardRecipient is a paid mutator transaction binding the contract method 0x0deab73c.
//
// Solidity: function setZoraRewardRecipient(address recipient) returns()
func (_Zorafirst *ZorafirstSession) SetZoraRewardRecipient(recipient common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.SetZoraRewardRecipient(&_Zorafirst.TransactOpts, recipient)
}

// SetZoraRewardRecipient is a paid mutator transaction binding the contract method 0x0deab73c.
//
// Solidity: function setZoraRewardRecipient(address recipient) returns()
func (_Zorafirst *ZorafirstTransactorSession) SetZoraRewardRecipient(recipient common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.SetZoraRewardRecipient(&_Zorafirst.TransactOpts, recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zorafirst *ZorafirstTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zorafirst *ZorafirstSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.TransferOwnership(&_Zorafirst.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Zorafirst *ZorafirstTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Zorafirst.Contract.TransferOwnership(&_Zorafirst.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Zorafirst *ZorafirstTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Zorafirst *ZorafirstSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Zorafirst.Contract.UniswapV3SwapCallback(&_Zorafirst.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Zorafirst *ZorafirstTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Zorafirst.Contract.UniswapV3SwapCallback(&_Zorafirst.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x8842d6d9.
//
// Solidity: function updateSale(uint256 tokenId, uint64 newStartTime, uint64 newEndTime) returns()
func (_Zorafirst *ZorafirstTransactor) UpdateSale(opts *bind.TransactOpts, tokenId *big.Int, newStartTime uint64, newEndTime uint64) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "updateSale", tokenId, newStartTime, newEndTime)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x8842d6d9.
//
// Solidity: function updateSale(uint256 tokenId, uint64 newStartTime, uint64 newEndTime) returns()
func (_Zorafirst *ZorafirstSession) UpdateSale(tokenId *big.Int, newStartTime uint64, newEndTime uint64) (*types.Transaction, error) {
	return _Zorafirst.Contract.UpdateSale(&_Zorafirst.TransactOpts, tokenId, newStartTime, newEndTime)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x8842d6d9.
//
// Solidity: function updateSale(uint256 tokenId, uint64 newStartTime, uint64 newEndTime) returns()
func (_Zorafirst *ZorafirstTransactorSession) UpdateSale(tokenId *big.Int, newStartTime uint64, newEndTime uint64) (*types.Transaction, error) {
	return _Zorafirst.Contract.UpdateSale(&_Zorafirst.TransactOpts, tokenId, newStartTime, newEndTime)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorafirst *ZorafirstTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorafirst.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorafirst *ZorafirstSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorafirst.Contract.UpgradeToAndCall(&_Zorafirst.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorafirst *ZorafirstTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorafirst.Contract.UpgradeToAndCall(&_Zorafirst.TransactOpts, newImplementation, data)
}

// ZorafirstInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Zorafirst contract.
type ZorafirstInitializedIterator struct {
	Event *ZorafirstInitialized // Event containing the contract specifics and raw log

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
func (it *ZorafirstInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstInitialized)
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
		it.Event = new(ZorafirstInitialized)
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
func (it *ZorafirstInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstInitialized represents a Initialized event raised by the Zorafirst contract.
type ZorafirstInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Zorafirst *ZorafirstFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZorafirstInitializedIterator, error) {

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZorafirstInitializedIterator{contract: _Zorafirst.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Zorafirst *ZorafirstFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZorafirstInitialized) (event.Subscription, error) {

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstInitialized)
				if err := _Zorafirst.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Zorafirst *ZorafirstFilterer) ParseInitialized(log types.Log) (*ZorafirstInitialized, error) {
	event := new(ZorafirstInitialized)
	if err := _Zorafirst.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstMarketLaunchedIterator is returned from FilterMarketLaunched and is used to iterate over the raw logs and unpacked data for MarketLaunched events raised by the Zorafirst contract.
type ZorafirstMarketLaunchedIterator struct {
	Event *ZorafirstMarketLaunched // Event containing the contract specifics and raw log

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
func (it *ZorafirstMarketLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstMarketLaunched)
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
		it.Event = new(ZorafirstMarketLaunched)
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
func (it *ZorafirstMarketLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstMarketLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstMarketLaunched represents a MarketLaunched event raised by the Zorafirst contract.
type ZorafirstMarketLaunched struct {
	Collection    common.Address
	TokenId       *big.Int
	Erc20zAddress common.Address
	PoolAddress   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketLaunched is a free log retrieval operation binding the contract event 0xc6c674475fa7ff9d2766a2efde8df64fa54848e42139f88a888916b07bcfcdbe.
//
// Solidity: event MarketLaunched(address indexed collection, uint256 indexed tokenId, address erc20zAddress, address poolAddress)
func (_Zorafirst *ZorafirstFilterer) FilterMarketLaunched(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int) (*ZorafirstMarketLaunchedIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "MarketLaunched", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstMarketLaunchedIterator{contract: _Zorafirst.contract, event: "MarketLaunched", logs: logs, sub: sub}, nil
}

// WatchMarketLaunched is a free log subscription operation binding the contract event 0xc6c674475fa7ff9d2766a2efde8df64fa54848e42139f88a888916b07bcfcdbe.
//
// Solidity: event MarketLaunched(address indexed collection, uint256 indexed tokenId, address erc20zAddress, address poolAddress)
func (_Zorafirst *ZorafirstFilterer) WatchMarketLaunched(opts *bind.WatchOpts, sink chan<- *ZorafirstMarketLaunched, collection []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "MarketLaunched", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstMarketLaunched)
				if err := _Zorafirst.contract.UnpackLog(event, "MarketLaunched", log); err != nil {
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

// ParseMarketLaunched is a log parse operation binding the contract event 0xc6c674475fa7ff9d2766a2efde8df64fa54848e42139f88a888916b07bcfcdbe.
//
// Solidity: event MarketLaunched(address indexed collection, uint256 indexed tokenId, address erc20zAddress, address poolAddress)
func (_Zorafirst *ZorafirstFilterer) ParseMarketLaunched(log types.Log) (*ZorafirstMarketLaunched, error) {
	event := new(ZorafirstMarketLaunched)
	if err := _Zorafirst.contract.UnpackLog(event, "MarketLaunched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstMintCommentIterator is returned from FilterMintComment and is used to iterate over the raw logs and unpacked data for MintComment events raised by the Zorafirst contract.
type ZorafirstMintCommentIterator struct {
	Event *ZorafirstMintComment // Event containing the contract specifics and raw log

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
func (it *ZorafirstMintCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstMintComment)
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
		it.Event = new(ZorafirstMintComment)
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
func (it *ZorafirstMintCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstMintCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstMintComment represents a MintComment event raised by the Zorafirst contract.
type ZorafirstMintComment struct {
	Sender     common.Address
	Collection common.Address
	TokenId    *big.Int
	Quantity   *big.Int
	Comment    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintComment is a free log retrieval operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed collection, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Zorafirst *ZorafirstFilterer) FilterMintComment(opts *bind.FilterOpts, sender []common.Address, collection []common.Address, tokenId []*big.Int) (*ZorafirstMintCommentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "MintComment", senderRule, collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstMintCommentIterator{contract: _Zorafirst.contract, event: "MintComment", logs: logs, sub: sub}, nil
}

// WatchMintComment is a free log subscription operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed collection, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Zorafirst *ZorafirstFilterer) WatchMintComment(opts *bind.WatchOpts, sink chan<- *ZorafirstMintComment, sender []common.Address, collection []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "MintComment", senderRule, collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstMintComment)
				if err := _Zorafirst.contract.UnpackLog(event, "MintComment", log); err != nil {
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

// ParseMintComment is a log parse operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed collection, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Zorafirst *ZorafirstFilterer) ParseMintComment(log types.Log) (*ZorafirstMintComment, error) {
	event := new(ZorafirstMintComment)
	if err := _Zorafirst.contract.UnpackLog(event, "MintComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Zorafirst contract.
type ZorafirstOwnershipTransferStartedIterator struct {
	Event *ZorafirstOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ZorafirstOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstOwnershipTransferStarted)
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
		it.Event = new(ZorafirstOwnershipTransferStarted)
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
func (it *ZorafirstOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Zorafirst contract.
type ZorafirstOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Zorafirst *ZorafirstFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZorafirstOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstOwnershipTransferStartedIterator{contract: _Zorafirst.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Zorafirst *ZorafirstFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ZorafirstOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstOwnershipTransferStarted)
				if err := _Zorafirst.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Zorafirst *ZorafirstFilterer) ParseOwnershipTransferStarted(log types.Log) (*ZorafirstOwnershipTransferStarted, error) {
	event := new(ZorafirstOwnershipTransferStarted)
	if err := _Zorafirst.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zorafirst contract.
type ZorafirstOwnershipTransferredIterator struct {
	Event *ZorafirstOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZorafirstOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstOwnershipTransferred)
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
		it.Event = new(ZorafirstOwnershipTransferred)
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
func (it *ZorafirstOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstOwnershipTransferred represents a OwnershipTransferred event raised by the Zorafirst contract.
type ZorafirstOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zorafirst *ZorafirstFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZorafirstOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstOwnershipTransferredIterator{contract: _Zorafirst.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Zorafirst *ZorafirstFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZorafirstOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstOwnershipTransferred)
				if err := _Zorafirst.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Zorafirst *ZorafirstFilterer) ParseOwnershipTransferred(log types.Log) (*ZorafirstOwnershipTransferred, error) {
	event := new(ZorafirstOwnershipTransferred)
	if err := _Zorafirst.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstSaleSetIterator is returned from FilterSaleSet and is used to iterate over the raw logs and unpacked data for SaleSet events raised by the Zorafirst contract.
type ZorafirstSaleSetIterator struct {
	Event *ZorafirstSaleSet // Event containing the contract specifics and raw log

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
func (it *ZorafirstSaleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstSaleSet)
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
		it.Event = new(ZorafirstSaleSet)
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
func (it *ZorafirstSaleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstSaleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstSaleSet represents a SaleSet event raised by the Zorafirst contract.
type ZorafirstSaleSet struct {
	Collection    common.Address
	TokenId       *big.Int
	SalesConfig   IZoraTimedSaleStrategySalesConfig
	Erc20zAddress common.Address
	PoolAddress   common.Address
	MintFee       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSaleSet is a free log retrieval operation binding the contract event 0x0eb75baee9035d7d1cf49fafe6f85bda3c71c4ec61a1d3c6a27877e58b55c73b.
//
// Solidity: event SaleSet(address indexed collection, uint256 indexed tokenId, (uint64,uint64,string,string) salesConfig, address erc20zAddress, address poolAddress, uint256 mintFee)
func (_Zorafirst *ZorafirstFilterer) FilterSaleSet(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int) (*ZorafirstSaleSetIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "SaleSet", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstSaleSetIterator{contract: _Zorafirst.contract, event: "SaleSet", logs: logs, sub: sub}, nil
}

// WatchSaleSet is a free log subscription operation binding the contract event 0x0eb75baee9035d7d1cf49fafe6f85bda3c71c4ec61a1d3c6a27877e58b55c73b.
//
// Solidity: event SaleSet(address indexed collection, uint256 indexed tokenId, (uint64,uint64,string,string) salesConfig, address erc20zAddress, address poolAddress, uint256 mintFee)
func (_Zorafirst *ZorafirstFilterer) WatchSaleSet(opts *bind.WatchOpts, sink chan<- *ZorafirstSaleSet, collection []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "SaleSet", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstSaleSet)
				if err := _Zorafirst.contract.UnpackLog(event, "SaleSet", log); err != nil {
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

// ParseSaleSet is a log parse operation binding the contract event 0x0eb75baee9035d7d1cf49fafe6f85bda3c71c4ec61a1d3c6a27877e58b55c73b.
//
// Solidity: event SaleSet(address indexed collection, uint256 indexed tokenId, (uint64,uint64,string,string) salesConfig, address erc20zAddress, address poolAddress, uint256 mintFee)
func (_Zorafirst *ZorafirstFilterer) ParseSaleSet(log types.Log) (*ZorafirstSaleSet, error) {
	event := new(ZorafirstSaleSet)
	if err := _Zorafirst.contract.UnpackLog(event, "SaleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Zorafirst contract.
type ZorafirstUpgradedIterator struct {
	Event *ZorafirstUpgraded // Event containing the contract specifics and raw log

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
func (it *ZorafirstUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstUpgraded)
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
		it.Event = new(ZorafirstUpgraded)
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
func (it *ZorafirstUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstUpgraded represents a Upgraded event raised by the Zorafirst contract.
type ZorafirstUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zorafirst *ZorafirstFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZorafirstUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstUpgradedIterator{contract: _Zorafirst.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zorafirst *ZorafirstFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZorafirstUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstUpgraded)
				if err := _Zorafirst.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zorafirst *ZorafirstFilterer) ParseUpgraded(log types.Log) (*ZorafirstUpgraded, error) {
	event := new(ZorafirstUpgraded)
	if err := _Zorafirst.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstZoraRewardRecipientUpdatedIterator is returned from FilterZoraRewardRecipientUpdated and is used to iterate over the raw logs and unpacked data for ZoraRewardRecipientUpdated events raised by the Zorafirst contract.
type ZorafirstZoraRewardRecipientUpdatedIterator struct {
	Event *ZorafirstZoraRewardRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ZorafirstZoraRewardRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstZoraRewardRecipientUpdated)
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
		it.Event = new(ZorafirstZoraRewardRecipientUpdated)
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
func (it *ZorafirstZoraRewardRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstZoraRewardRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstZoraRewardRecipientUpdated represents a ZoraRewardRecipientUpdated event raised by the Zorafirst contract.
type ZorafirstZoraRewardRecipientUpdated struct {
	PrevRecipient common.Address
	NewRecipient  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZoraRewardRecipientUpdated is a free log retrieval operation binding the contract event 0x6aca4ad70d40a81965e7a68e8c84c9edbe9946296012177f423fd0e3e4687827.
//
// Solidity: event ZoraRewardRecipientUpdated(address indexed prevRecipient, address indexed newRecipient)
func (_Zorafirst *ZorafirstFilterer) FilterZoraRewardRecipientUpdated(opts *bind.FilterOpts, prevRecipient []common.Address, newRecipient []common.Address) (*ZorafirstZoraRewardRecipientUpdatedIterator, error) {

	var prevRecipientRule []interface{}
	for _, prevRecipientItem := range prevRecipient {
		prevRecipientRule = append(prevRecipientRule, prevRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "ZoraRewardRecipientUpdated", prevRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstZoraRewardRecipientUpdatedIterator{contract: _Zorafirst.contract, event: "ZoraRewardRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchZoraRewardRecipientUpdated is a free log subscription operation binding the contract event 0x6aca4ad70d40a81965e7a68e8c84c9edbe9946296012177f423fd0e3e4687827.
//
// Solidity: event ZoraRewardRecipientUpdated(address indexed prevRecipient, address indexed newRecipient)
func (_Zorafirst *ZorafirstFilterer) WatchZoraRewardRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ZorafirstZoraRewardRecipientUpdated, prevRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var prevRecipientRule []interface{}
	for _, prevRecipientItem := range prevRecipient {
		prevRecipientRule = append(prevRecipientRule, prevRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "ZoraRewardRecipientUpdated", prevRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstZoraRewardRecipientUpdated)
				if err := _Zorafirst.contract.UnpackLog(event, "ZoraRewardRecipientUpdated", log); err != nil {
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

// ParseZoraRewardRecipientUpdated is a log parse operation binding the contract event 0x6aca4ad70d40a81965e7a68e8c84c9edbe9946296012177f423fd0e3e4687827.
//
// Solidity: event ZoraRewardRecipientUpdated(address indexed prevRecipient, address indexed newRecipient)
func (_Zorafirst *ZorafirstFilterer) ParseZoraRewardRecipientUpdated(log types.Log) (*ZorafirstZoraRewardRecipientUpdated, error) {
	event := new(ZorafirstZoraRewardRecipientUpdated)
	if err := _Zorafirst.contract.UnpackLog(event, "ZoraRewardRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorafirstZoraTimedSaleStrategyRewardsIterator is returned from FilterZoraTimedSaleStrategyRewards and is used to iterate over the raw logs and unpacked data for ZoraTimedSaleStrategyRewards events raised by the Zorafirst contract.
type ZorafirstZoraTimedSaleStrategyRewardsIterator struct {
	Event *ZorafirstZoraTimedSaleStrategyRewards // Event containing the contract specifics and raw log

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
func (it *ZorafirstZoraTimedSaleStrategyRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorafirstZoraTimedSaleStrategyRewards)
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
		it.Event = new(ZorafirstZoraTimedSaleStrategyRewards)
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
func (it *ZorafirstZoraTimedSaleStrategyRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorafirstZoraTimedSaleStrategyRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorafirstZoraTimedSaleStrategyRewards represents a ZoraTimedSaleStrategyRewards event raised by the Zorafirst contract.
type ZorafirstZoraTimedSaleStrategyRewards struct {
	Collection           common.Address
	TokenId              *big.Int
	Creator              common.Address
	CreatorReward        *big.Int
	CreateReferral       common.Address
	CreateReferralReward *big.Int
	MintReferral         common.Address
	MintReferralReward   *big.Int
	Market               common.Address
	MarketReward         *big.Int
	ZoraRecipient        common.Address
	ZoraReward           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterZoraTimedSaleStrategyRewards is a free log retrieval operation binding the contract event 0xc773e203af3f3079b18c21f98bb8d8ccd2fea097d631d448df89de4edbe7a2a8.
//
// Solidity: event ZoraTimedSaleStrategyRewards(address indexed collection, uint256 indexed tokenId, address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address market, uint256 marketReward, address zoraRecipient, uint256 zoraReward)
func (_Zorafirst *ZorafirstFilterer) FilterZoraTimedSaleStrategyRewards(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int) (*ZorafirstZoraTimedSaleStrategyRewardsIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.FilterLogs(opts, "ZoraTimedSaleStrategyRewards", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorafirstZoraTimedSaleStrategyRewardsIterator{contract: _Zorafirst.contract, event: "ZoraTimedSaleStrategyRewards", logs: logs, sub: sub}, nil
}

// WatchZoraTimedSaleStrategyRewards is a free log subscription operation binding the contract event 0xc773e203af3f3079b18c21f98bb8d8ccd2fea097d631d448df89de4edbe7a2a8.
//
// Solidity: event ZoraTimedSaleStrategyRewards(address indexed collection, uint256 indexed tokenId, address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address market, uint256 marketReward, address zoraRecipient, uint256 zoraReward)
func (_Zorafirst *ZorafirstFilterer) WatchZoraTimedSaleStrategyRewards(opts *bind.WatchOpts, sink chan<- *ZorafirstZoraTimedSaleStrategyRewards, collection []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorafirst.contract.WatchLogs(opts, "ZoraTimedSaleStrategyRewards", collectionRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorafirstZoraTimedSaleStrategyRewards)
				if err := _Zorafirst.contract.UnpackLog(event, "ZoraTimedSaleStrategyRewards", log); err != nil {
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

// ParseZoraTimedSaleStrategyRewards is a log parse operation binding the contract event 0xc773e203af3f3079b18c21f98bb8d8ccd2fea097d631d448df89de4edbe7a2a8.
//
// Solidity: event ZoraTimedSaleStrategyRewards(address indexed collection, uint256 indexed tokenId, address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address market, uint256 marketReward, address zoraRecipient, uint256 zoraReward)
func (_Zorafirst *ZorafirstFilterer) ParseZoraTimedSaleStrategyRewards(log types.Log) (*ZorafirstZoraTimedSaleStrategyRewards, error) {
	event := new(ZorafirstZoraTimedSaleStrategyRewards)
	if err := _Zorafirst.contract.UnpackLog(event, "ZoraTimedSaleStrategyRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
