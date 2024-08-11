// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package olympic

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

// CollectionSecurityPolicy is an auto generated low-level Go binding around an user-defined struct.
type CollectionSecurityPolicy struct {
	TransferSecurityLevel        uint8
	OperatorWhitelistId          *big.Int
	PermittedContractReceiversId *big.Int
}

// IERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// IERC721MMintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type IERC721MMintStageInfo struct {
	Price                *big.Int
	MintFee              *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
}

// OlympicMetaData contains all meta data concerning the Olympic contract.
var OlympicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collectionSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURISuffix\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestampExpirySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyFeeNumerator\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotIncreaseMaxMintableSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorTokenBase__InvalidTransferValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorTokenBase__SetTransferValidatorFirst\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintAddressNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalWalletLimitOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStageTimeGap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStageArgsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSupplyLeft\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShouldNotMintToBurnAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StageSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletGlobalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletStageLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongMintCurrency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"DefaultRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeStage\",\"type\":\"uint256\"}],\"name\":\"SetActiveStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"SetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"SetCosigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"SetCrossmintAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"SetGlobalWalletLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"SetMaxMintableSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"}],\"name\":\"SetMintCurrency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"SetMintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"SetTimestampExpirySeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"TokenRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"TransferValidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"name\":\"UpdateStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_OPERATOR_WHITELIST_ID\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSFER_SECURITY_LEVEL\",\"outputs\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSFER_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUND_RECEIVER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"assertValidCosign\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"crossmint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getActiveStageFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getCosignDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getCosignNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalWalletLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberStages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermittedContractReceivers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecurityPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"transferSecurityLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversId\",\"type\":\"uint120\"}],\"internalType\":\"structCollectionSecurityPolicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStageInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferValidator\",\"outputs\":[{\"internalType\":\"contractICreatorTokenTransferValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"isContractReceiverPermitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintWithLimit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"setCrossmintAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"setGlobalWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxMintableSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo[]\",\"name\":\"newStages\",\"type\":\"tuple[]\"}],\"name\":\"setStages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversAllowlistId\",\"type\":\"uint120\"}],\"name\":\"setToCustomSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversAllowlistId\",\"type\":\"uint120\"}],\"name\":\"setToCustomValidatorAndSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setToDefaultSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferValidator_\",\"type\":\"address\"}],\"name\":\"setTransferValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"totalMintedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436101561001257600080fd5b60003560e01c8063014635461461044257806301ffc9a71461043d578063020451381461043857806304634d8d1461043357806306fdde031461042e578063081812fc14610429578063095ea7b314610424578063098144d41461041f57806318160ddd1461041a5780631b25b077146104155780631c33b328146104105780631ce03eed1461040b57806323b872dd14610406578063285d70d4146104015780632a55205a146103fc5780632e8da829146103f75780632ed6d5e8146103f2578063372992e4146103ed5780633ccfd60b146103e85780633d6375b2146103e3578063424aa884146103de57806342842e0e146103d9578063495c8bf9146103d45780634b1c53b4146103cf57806355f804b3146103ca5780635944c753146103c55780635bbb2177146103c05780635d4c1d46146103bb57806361347162146103b657806362acbd9a146103b15780636352211e146103ac57806367808a34146103a75780636c3b8699146103a2578063700d19f21461039d57806370a082311461039857806370da24ee14610393578063715018a61461038e5780638462151c146103895780638da5cb5b14610384578063938e3d7b1461037f57806395d89b411461037a57806397cf84fc14610366578063997556241461037557806399a2557a146103705780639d645a441461036b578063a06c492f14610366578063a22cb46514610361578063a3759f601461035c578063a9852bfb14610357578063a9fc664e14610352578063aac5ab1f1461034d578063b50248e714610348578063b88d4fde14610343578063be537f431461033e578063c23dc68f14610339578063c87b56dd14610334578063ce2b0ec01461032f578063d007af5c1461032a578063e8a3d48514610325578063e985e9c514610320578063efb6b11f1461031b578063efdaa2ec14610316578063f2fde38b14610311578063f698bceb1461030c578063f830e8b814610307578063f8d09696146103025763fd762d92146102fd57600080fd5b6133d3565b61336b565b613030565b61300d565b612f80565b612f62565b612ca1565b612c39565b612b92565b612b76565b612af9565b612a24565b6129c1565b612976565b612925565b6128c6565b6126f2565b6126cd565b6125dc565b61252e565b61240a565b6122f0565b6123e3565b6123a7565b612337565b612249565b612158565b61212f565b612071565b611fd8565b611fba565b611f93565b611f4e565b611e98565b611e71565b611e42565b611aa2565b611938565b611901565b611830565b611692565b611547565b6114fb565b6114d3565b61146c565b611443565b610f5a565b610e1c565b610db4565b610c56565b610c2f565b610b88565b610b26565b610b08565b610a94565b610a59565b6109e7565b6109c4565b61099b565b6108e0565b61088c565b6107ab565b610634565b61056d565b610498565b610457565b600091031261045257565b600080fd5b3461045257600036600319011261045257602060405173721c00182a990771244d7a71b9fa2ea789a3b4338152f35b6001600160e01b031981160361045257565b34610452576020366003190112610452576104ed6004356104b881610486565b63ffffffff60e01b166310c8aba560e31b811490811561051c575b81156104f1575b5060405190151581529081906020820190565b0390f35b63152a902d60e11b81149150811561050b575b50386104da565b6301ffc9a760e01b14905038610504565b90506301ffc9a760e01b8114801561054c575b801561053c575b906104d3565b50635b5e139f60e01b8114610536565b506380ac58cd60e01b811461052f565b6001600160a01b0381160361045257565b34610452576020366003190112610452577faea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd960206004356105ad8161055c565b6105b561345d565b600d80547fffffff0000000000000000000000000000000000000000ffffffffffffffffff16604883901b6901000000000000000000600160e81b03161790556040516001600160a01b039091168152a1005b602435906001600160601b038216820361045257565b604435906001600160601b038216820361045257565b34610452576040366003190112610452576004356106518161055c565b610659610608565b61066161345d565b6001600160601b0381166127108082116107315750506001600160a01b038216918215610718577f8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef916106f9610713926106cb6106bc612830565b6001600160a01b039092168252565b6001600160601b0383166020820152516001600160a01b03166001600160a01b031960a084901b1617600855565b6040516001600160601b0390911681529081906020820190565b0390a2005b604051635b6cc80560e11b815260006004820152602490fd5b6044925060405191636f483d0960e01b835260048301526024820152fd5b60005b8381106107625750506000910152565b8181015183820152602001610752565b9060209161078b8151809281855285808601910161074f565b601f01601f1916010190565b9060206107a8928181520190610772565b90565b34610452576000806003193601126108895760405190806002546107ce81613bb3565b8085529160019180831690811561085f5750600114610804575b6104ed856107f88187038261280f565b60405191829182610797565b9250600283527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b8284106108475750505081016020016107f8826104ed6107e8565b8054602085870181019190915290930192810161082c565b8695506104ed969350602092506107f894915060ff191682840152151560051b82010192936107e8565b80fd5b34610452576020366003190112610452576004356108a981614a93565b156108ce576000526006602052602060018060a01b0360406000205416604051908152f35b6040516333d1c03960e21b8152600490fd5b6040366003190112610452576004356108f88161055c565b6024356001600160a01b038061090d83614a24565b1690813303610968575b600083815260066020526040812080546001600160a01b0319166001600160a01b0387161790559316907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258480a480f35b600082815260076020908152604080832033845290915290205460ff16610917576040516367d9dca160e11b8152600490fd5b3461045257600036600319011261045257600a546040516001600160a01b039091168152602090f35b346104525760003660031901126104525760206000546001549003604051908152f35b34610452576060366003190112610452576020610a27600435610a098161055c565b602435610a158161055c565b60443591610a228361055c565b614967565b6040519015158152f35b634e487b7160e01b600052602160045260246000fd5b906009821015610a545752565b610a31565b3461045257600036600319011261045257602060405160028152f35b63ffffffff81160361045257565b6001600160401b0381160361045257565b34610452576060366003190112610452576020610ad4600435610ab68161055c565b602435610ac281610a75565b60443591610acf83610a83565b613eca565b604051908152f35b606090600319011261045257600435610af48161055c565b90602435610b018161055c565b9060443590565b610b1a610b1436610adc565b91614abc565b005b8015150361045257565b34610452576020366003190112610452577fe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf6020600435610b6681610b1c565b610b6e61345d565b151560ff19600d541660ff821617600d55604051908152a1005b34610452576040366003190112610452576004356000526009602052604060002060405190610bb682612788565b546001600160a01b03811680835260a09190911c602083015215610c21575b610c05612710610bf46001600160601b036020850151166024356134c5565b92519204916001600160a01b031690565b604080516001600160a01b039290921682526020820192909252f35b50610c2a613489565b610bd5565b34610452576020366003190112610452576020610a27600435610c518161055c565b6147de565b346104525760008060031936011261088957610c7061345d565b6017546001600160a01b039081168015610da257601854610c9091613b1a565b601754610cad906001600160a01b03165b6001600160a01b031690565b6040516370a0823160e01b8152306004820152916020908390602490829085165afa908115610d9d577fbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22928492610d69575b5081610d2c917f00000000000000000000000044ca337e4a925b5d70c5bc2c1cc3786f654bcdd590613b72565b601754610d45906001600160a01b03169160185461385e565b604080516001600160a01b0390931683526020830191909152819081015b0390a180f35b610d2c919250610d8f9060203d8111610d96575b610d87818361280f565b810190613aff565b9190610cff565b503d610d7d565b613b0e565b60405163a47ca0b760e01b8152600490fd5b3461045257602036600319011261045257600435610dd061345d565b600f548111610e0a576020817f5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b3592601055604051908152a1005b604051630590c51360e01b8152600490fd5b346104525760008060031936011261088957610e3661345d565b80808080601854730b98151bedee73f9ba5f2c7b72dea02d38ce49fc5af1610e5c613a00565b5015610eeb574781808080847f00000000000000000000000044ca337e4a925b5d70c5bc2c1cc3786f654bcdd55af1610e93613a00565b5015610ed957610d63610ec97f5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d9260185461385e565b6040519081529081906020820190565b604051631d42c86760e21b8152600490fd5b6040516312171d8360e31b8152600490fd5b9181601f84011215610452578235916001600160401b038311610452576020808501948460051b01011161045257565b9181601f84011215610452578235916001600160401b038311610452576020838186019501011161045257565b60a0366003190112610452576004803590610f7482610a75565b60243591610f8183610a75565b6001600160401b039160443583811161045257610fa19036908301610efd565b60643592610fae84610a83565b60843586811161045257610fc59036908301610f2d565b610fcd61383b565b600d549560ff8716156114325763ffffffff928389169960005460015490038b810180911161142d57600f541061141c579061102c939291421692611010613795565b506001600160a01b039960481c8a166113f1575b505050614170565b9761103f6110398a6135e6565b506137cd565b6017549096166001600160a01b03169485159485806113b9575b6113a857608088018b8b62ffffff611074845162ffffff1690565b16611355575b5050506010548061130c575b5060408801848c8b8261109d855163ffffffff1690565b166112ba575b50505050606088019182516111d6575b5050505050509261110c6111076110ff6110f3886110ee60206111899c99611166996111769c15611193575b5001516001600160501b031690565b613886565b6001600160501b031690565b60185461385e565b601855565b61115261113e33611127866000526015602052604060002090565b9060018060a01b0316600052602052604060002090565b9161114d835463ffffffff1690565b6138a4565b63ffffffff1663ffffffff19825416179055565b6000526016602052604060002090565b61118182825461385e565b905533614f6b565b610b1a6001600c55565b6111d0906111c66110f3876110ee6111b287516001600160501b031690565b878901516001600160501b03165b9061386b565b9030903390613921565b386110df565b6040513360601b6bffffffffffffffffffffffff19166020820190815260e086901b6001600160e01b031916603483015261123a93926112359290919061122a81603881015b03601f19810183528261280f565b5190209236916138d3565b613aa7565b9051036112a95716801515908161126a575b5061125b5780808080806110b3565b60405163b4f3729b60e01b8152fd5b90506112a26112998761114d61128f8c61112733916000526015602052604060002090565b5463ffffffff1690565b63ffffffff1690565b113861124c565b6040516309bde33960e01b81528390fd5b6112e06112999161114d61128f6112eb9561112733916000526015602052604060002090565b935163ffffffff1690565b9116116112fb5738848c8b6110a3565b60405163b4f3729b60e01b81528590fd5b336000908152600560205260409081902054611339918d91901c6001600160401b031661385e565b61385e565b116113445738611086565b60405163751304ed60e11b81528590fd5b6113796113839161137361138b946000526016602052604060002090565b5461385e565b925162ffffff1690565b62ffffff1690565b1061139857388b8b61107a565b60405162d0844960e21b81528590fd5b604051630717c22560e51b81528590fd5b506113ea6110f38b6110ee8b6111c060206113db83516001600160501b031690565b9201516001600160501b031690565b3410611059565b61140b92935090611403913691612874565b828a33613fbb565b611414816141e8565b388080611024565b60405163800113cb60e01b81528690fd5b6134af565b604051630952c8a960e11b81528490fd5b34610452576000366003190112610452576017546040516001600160a01b039091168152602090f35b610b1a61147836610adc565b9060405192611486846127d9565b60008452614de4565b6020908160408183019282815285518094520193019160005b8281106114b6575050505090565b83516001600160a01b0316855293810193928101926001016114a8565b34610452576000366003190112610452576104ed6114ef614656565b6040519182918261148f565b34610452576000366003190112610452576020600f54604051908152f35b602060031982011261045257600435906001600160401b0382116104525761154391600401610f2d565b9091565b346104525761155536611519565b61155d61345d565b6001600160401b03811161168d5761157f8161157a601154613bb3565b613bed565b600091601f82116001146115ec57817f23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa936000916115e1575b508260011b906000198460031b1c1916176011555b6115dc60405192839283613d40565b0390a1005b9050810135386115b8565b60116000527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68601f198316845b818110611675575093837f23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa951061165b575b5050600182811b016011556115cd565b820135600019600385901b60f8161c19169055388061164b565b83860135835560209586019560019093019201611619565b61275f565b34610452576060366003190112610452576024356004356116b28261055c565b6116ba61061e565b6116c261345d565b6127106001600160601b03821681811161178d5750506001600160a01b03831692831561176d577f7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c916106f96117689261171d6106bc612830565b6001600160601b0383166020820152611740866000526009602052604060002090565b815160209092015160a01b6001600160a01b0319166001600160a01b03909216919091179055565b0390a3005b604051634b4f842960e11b81526004810184905260006024820152604490fd5b60405163dfd1fc1b60e01b815260048101949094526024840152604483015250606490fd5b6020908160408183019282815285518094520193019160005b8281106117d9575050505090565b9091929382608082611824600194895162ffffff6060809260018060a01b0381511685526001600160401b036020820151166020860152604081015115156040860152015116910152565b019501939291016117cb565b3461045257602080600319360112610452576004356001600160401b03811161045257611861903690600401610efd565b61186d819392936138bc565b9161187b604051938461280f565b818352601f1961188a836138bc565b0160005b8181106118eb5750505060005b8181036118b057604051806104ed85826117b2565b818110156118e657806118ca60019260051b86013561510d565b6118d48286613a93565b526118df8185613a93565b500161189b565b61355e565b82906118f56150d8565b8282880101520161188e565b3461045257600036600319011261045257602060405160018152f35b6009111561045257565b6001600160781b0381160361045257565b34610452576060366003190112610452576004356119558161191d565b60243561196181611927565b60443561196d81611927565b61197561345d565b600a546001600160a01b0316908115611a9057813b1561045257604051630368065360e61b81526000948590829081906119b3903060048401614243565b038183875af18015610d9d57611a7d575b5083823b1561088957604051631182550160e11b81523060048201526001600160781b039490941660248501528360448183865af1928315610d9d578493611a6a575b50813b15611a665760405163235d10c560e21b81523060048201526001600160781b03909116602482015291908290818381604481015b03925af18015610d9d57611a50575080f35b80611a5d611a6392612775565b80610447565b80f35b5050fd5b80611a5d611a7792612775565b38611a07565b80611a5d611a8a92612775565b386119c4565b604051631cffe3dd60e11b8152600490fd5b600060a0366003190112610889576004908135611abe81610a75565b602435611aca8161055c565b6001600160401b0391604435838111611e3e57611aea9036908701610efd565b60649691963591611afa83610a83565b608435868111611e3a57611b119036908301610f2d565b9190611b1b61383b565b600e546001600160a01b03959086168015611e29573303611e1857600d549160ff8316156114325763ffffffff948589169a8c5460015490038c810180911161142d57600f541061141c579088611b8b95949392421694611b7a613795565b5060481c16611dfe57505050614170565b98611b986110398b6135e6565b6017549095166001600160a01b0316938415938480611dc8575b611db7578b90608088018b62ffffff611bce835162ffffff1690565b16611d79575b505060105480611d2b575b50888a60408a019383611bf6865163ffffffff1690565b16611ce4575b505050505060608601918251611c7d575b5050505092611c466111076110ff6110f3896110ee6020611c739d9e9961116699611c619c15611193575001516001600160501b031690565b61115261113e86611127866000526015602052604060002090565b611c6c83825461385e565b9055614f6b565b611a636001600c55565b90611235611cbb928a60409e999a979d969e5161122a8161121c6020820194856018916001600160601b03199060601b168152600060148201520190565b905103611cd357929750919591929185388080611c0d565b6040516309bde33960e01b81528990fd5b611d099261114d61128f611299946111276112e0956000526015602052604060002090565b911611611d1a573880888a8e611bfc565b60405163b4f3729b60e01b81528390fd5b909150611d5b8b6113348c60018060a01b031660005260056020526001600160401b0360406000205460401c1690565b11611d68578b9038611bdf565b60405163751304ed60e11b81528490fd5b61137961138391611373611d999495966000526016602052604060002090565b10611da7578b90388b611bd4565b60405162d0844960e21b81528490fd5b604051630717c22560e51b81528490fd5b50611df76110f38b6110ee611de48b516001600160501b031690565b60208c01516001600160501b03166111c0565b3410611bb2565b61140b92935090611e10913691612874565b828933613fbb565b60405163f46fd68360e01b81528390fd5b604051637e9f68eb60e11b81528490fd5b8780fd5b8480fd5b346104525760203660031901126104525760206001600160a01b03611e68600435614a24565b16604051908152f35b34610452576020366003190112610452576020610ad4600435611e9381610a83565b614170565b346104525760008060031936011261088957611eb261345d565b611eba614263565b73721c00182a990771244d7a71b9fa2ea789a3b433803b15611f3757604051630368065360e61b815230600482015260026024820152828160448183865af18015610d9d57611f3b575b50803b15611f3757604051631182550160e11b815230600482015260016024820152908290829081838160448101611a3e565b5080fd5b80611a5d611f4892612775565b38611f04565b34610452576000366003190112610452576040517f00000000000000000000000044ca337e4a925b5d70c5bc2c1cc3786f654bcdd56001600160a01b03168152602090f35b34610452576020366003190112610452576020610ad4600435611fb58161055c565b6149e9565b34610452576000366003190112610452576020601454604051908152f35b346104525760008060031936011261088957611ff261345d565b600b80546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b6020908160408183019282815285518094520193019160005b82811061205d575050505090565b83518552938101939281019260010161204f565b346104525760203660031901126104525760043561208e8161055c565b60008061209a836149e9565b916120a4836151b9565b936120ad6150d8565b506001600160a01b0390811691835b8585036120d157604051806104ed8982612036565b6120da8161515b565b604081015161212657516001600160a01b031683811661211d575b506001908484841614612109575b016120bc565b80612117838801978a613a93565b52612103565b915060016120f5565b50600190612103565b3461045257600036600319011261045257600b546040516001600160a01b039091168152602090f35b346104525761216636611519565b61216e61345d565b6001600160401b03811161168d576121908161218b601354613bb3565b613c5e565b6000601f82116001146121cb5781926000926121c0575b5050600019600383901b1c191660019190911b17601355005b0135905038806121a7565b6013600052601f198216927f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09091805b85811061223157508360019510612217575b505050811b01601355005b0135600019600384901b60f8161c1916905538808061220c565b909260206001819286860135815501940191016121fa565b346104525760008060031936011261088957604051908060035461226c81613bb3565b8085529160019180831690811561085f5750600114612295576104ed856107f88187038261280f565b9250600383527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8284106122d85750505081016020016107f8826104ed6107e8565b805460208587018101919091529093019281016122bd565b34610452576020366003190112610452576020610ad46004356123128161055c565b60018060a01b031660005260056020526001600160401b0360406000205460401c1690565b34610452576020366003190112610452577ff477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec460206004356123778161055c565b61237f61345d565b600e80546001600160a01b0319166001600160a01b03929092169182179055604051908152a1005b34610452576060366003190112610452576104ed6123d76004356123ca8161055c565b60443590602435906151eb565b60405191829182612036565b34610452576020366003190112610452576020610a276004356124058161055c565b6148b9565b34610452576040366003190112610452576004356124278161055c565b6024359061243482610b1c565b3360009081526007602090815260408083206001600160a01b038516845290915290209115159160ff1981541660ff841617905560405191825260018060a01b0316907f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b61252a6101009295949361251d60c06101208501986001600160501b03808251168752602082015116602087015263ffffffff60408201511660408701526060810151606087015262ffffff60808201511660808701526001600160401b0360a08201511660a0870152015160c08501906001600160401b03169052565b63ffffffff1660e0830152565b0152565b346104525760203660031901126104525760043561254a613795565b506014548110156125a8576000818152601560209081526040808320338452909152902063ffffffff905416906104ed61259b611039612594846000526016602052604060002090565b54936135e6565b916040519384938461249f565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964537461676560a01b6044820152606490fd5b34610452576125ea36611519565b6125f261345d565b6001600160401b03811161168d576126148161260f601254613bb3565b613ccf565b6000601f821160011461264f578192600092612644575b5050600019600383901b1c191660019190911b17601255005b01359050388061262b565b6012600052601f198216927fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec344491805b8581106126b55750836001951061269b575b505050811b01601255005b0135600019600384901b60f8161c19169055388080612690565b9092602060018192868601358155019401910161267e565b3461045257602036600319011261045257610b1a6004356126ed8161055c565b6143a2565b346104525760403660031901126104525760043561270f81610a75565b63ffffffff602435916127218361055c565b61272961345d565b16600054600154900381810180911161142d57600f541061274d57610b1a91614f6b565b60405163800113cb60e01b8152600490fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b03811161168d57604052565b604081019081106001600160401b0382111761168d57604052565b60e081019081106001600160401b0382111761168d57604052565b606081019081106001600160401b0382111761168d57604052565b602081019081106001600160401b0382111761168d57604052565b60c081019081106001600160401b0382111761168d57604052565b90601f801991011681019081106001600160401b0382111761168d57604052565b6040519061283d82612788565b565b6040519061283d826127a3565b6040519061283d826127be565b6001600160401b03811161168d57601f01601f191660200190565b92919261288082612859565b9161288e604051938461280f565b829481845281830111610452578281602093846000960137010152565b9080601f83011215610452578160206107a893359101612874565b34610452576080366003190112610452576004356128e38161055c565b602435906128f082610a75565b6044356128fc81610a83565b606435926001600160401b0384116104525761291f610b1a9436906004016128ab565b92613fbb565b60803660031901126104525760043561293d8161055c565b6024356129498161055c565b606435916001600160401b0383116104525761296c610b1a9336906004016128ab565b9160443591614de4565b34610452576000366003190112610452576060612991614517565b604051906129a0828251610a47565b60406001600160781b03918260208201511660208501520151166040820152f35b346104525760203660031901126104525760806129df60043561510d565b612a22604051809262ffffff6060809260018060a01b0381511685526001600160401b036020820151166020860152604081015115156040860152015116910152565bf35b3461045257602036600319011261045257600435612a4181614a93565b15612ae757612a4e613d68565b805160009015612acd575060405160a08101604052608081019260008452925b6000190192600a906030828206018553049283612a6e576104ed9350612ab692612ac1612abc6107f8946080601f199586810192030181526040519687946020860190613e24565b90613e24565b613e3b565b0390810183528261280f565b6040516104ed93509150612ae0826127d9565b81526107f8565b604051630a14c4b560e41b8152600490fd5b34610452576020366003190112610452577f41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f6020600435612b3981610a83565b612b4161345d565b600d5468ffffffffffffffff008260081b169068ffffffffffffffff00191617600d556001600160401b0360405191168152a1005b34610452576000366003190112610452576104ed6114ef614742565b3461045257600080600319360112610889576040519080601354612bb581613bb3565b8085529160019180831690811561085f5750600114612bde576104ed856107f88187038261280f565b9250601383527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0905b828410612c215750505081016020016107f8826104ed6107e8565b80546020858701810191909152909301928101612c06565b3461045257604036600319011261045257602060ff612c95600435612c5d8161055c565b60243590612c6a8261055c565b60018060a01b03166000526007845260406000209060018060a01b0316600052602052604060002090565b54166040519015158152f35b60006080366003190112610889576004803590612cbd82610a75565b6001600160401b0391602435838111611e3e57612cdd9036908401610efd565b60449391933591612ced83610a83565b606435868111611e3a57612d049036908301610f2d565b612d0f92919261383b565b600d549460ff861615612f515763ffffffff93848816998b5460015490038b810180911161142d57600f5410612f405790612d6c939291421692612d51613795565b506001600160a01b039860481c8916611dfe57505050614170565b95612d79611039886135e6565b6017549095166001600160a01b0316938415938480612f1d575b611db75760808701898b62ffffff612dae845162ffffff1690565b16612ef2575b50505060105480612ebf575b506040870190898982612dd7855163ffffffff1690565b16612e89575b5050505060608601918251612e26575b505050509261110c6111076110ff6110f3896110ee6020611c739c9b99611166996111769c15611193575001516001600160501b031690565b90611235612e609260409c969c51602081019061122a8161121c33856018916001600160601b03199060601b168152600060148201520190565b905103612e7a57509561110c6111076110ff6110f3612ded565b6040516309bde33960e01b8152fd5b6112e06112999161114d61128f612eaf9561112733916000526015602052604060002090565b911611611d1a5738808989612ddd565b336000908152600560205260409081902054612ee7918d91901c6001600160401b031661385e565b11611d685738612dc0565b61137961138391611373612f10946000526016602052604060002090565b10611da75738898b612db4565b50612f396110f38b6110ee611de48b516001600160501b031690565b3410612d93565b60405163800113cb60e01b81528590fd5b604051630952c8a960e11b81528390fd5b34610452576000366003190112610452576020601054604051908152f35b3461045257602036600319011261045257600435612f9d8161055c565b612fa561345d565b6001600160a01b03908116908115612ff457600b54826001600160601b0360a01b821617600b55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b604051631e4fbdf760e01b815260006004820152602490fd5b3461045257600036600319011261045257602060ff600d54166040519015158152f35b3461045257602080600319360112610452576004906001600160401b0390823582811161045257366023820112156104525780840135918383116104525760248201916024369160e086020101116104525761308a61345d565b6130926134ee565b60005b83811061309e57005b600190818110156132fc575b60a0806130b8838888613574565b016130c290613584565b60c090816130d1858a8a613574565b016130db90613584565b6130e49161421c565b866130f0848289613574565b6130f9906135b8565b818761310687838c613574565b01613110906135b8565b94896040938a848a87613124828488613574565b0161312e906135cc565b9a60609b8c61313e84868a613574565b0135918c60809889613151878984613574565b0161315b906135d6565b958b613168828a85613574565b0161317290613584565b9761317c92613574565b0161318690613584565b9561318f61283f565b6001600160501b0390981688526001600160501b039091169087015263ffffffff16858a0152848c015262ffffff16848401526001600160401b0316828501526001600160401b0316818801526131e590613621565b898b896131f3818784613574565b6131fc906135b8565b9761320692613574565b01613210906135b8565b958b898661321f828885613574565b01613229906135cc565b99613235828885613574565b013593613243828885613574565b0161324d906135d6565b94613259828885613574565b0161326390613584565b9561326d92613574565b0161327790613584565b93518881526001600160501b03958616602082015295909416604086015263ffffffff959095166060850152608084019290925262ffffff90931660a08301526001600160401b0390811660c083015290911660e08201527fc4737822c84fe15fce8213ef237bb06d7d6c1603adfa65bf6d3a6531959790929061010090a101613095565b61331260a061330c838888613574565b01613584565b8661335261334661333160c061330c61332a8861358e565b8c8c613574565b600d5460081c6001600160401b03169061359d565b6001600160401b031690565b911610156130aa57604051636bc1af9360e01b81528790fd5b346104525760203660031901126104525760043561338761345d565b600f5481116133c1576020817fc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a92600f55604051908152a1005b60405163430b83b160e11b8152600490fd5b34610452576080366003190112610452576004356133f08161055c565b602435906133fd8261191d565b6044359061340a82611927565b6064359061341782611927565b61341f61345d565b613428816143a2565b6001600160a01b031690813b1561045257604051630368065360e61b81526000948590829081906119b3903060048401614243565b600b546001600160a01b0316330361347157565b60405163118cdaa760e01b8152336004820152602490fd5b6040519061349682612788565b6008546001600160a01b038116835260a01c6020830152565b634e487b7160e01b600052601160045260246000fd5b8181029291811591840414171561142d57565b634e487b7160e01b600052600060045260246000fd5b6014546000908160145580613501575050565b600391818302918383040361142d57601481527fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec918201915b8281106135475750505050565b80828592558260018201558260028201550161353a565b634e487b7160e01b600052603260045260246000fd5b91908110156118e65760e0020190565b356107a881610a83565b60001981019190821161142d57565b9190916001600160401b038080941691160191821161142d57565b356001600160501b03811681036104525790565b356107a881610a75565b3562ffffff811681036104525790565b6014548110156118e6576003906014600052027fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec0190600090565b6014546801000000000000000081101561168d5780600161364592016014556135e6565b9190916137905761376860c0600261283d946001600160501b038551166001600160501b03198254161781556136b461368860208701516001600160501b031690565b825469ffffffffffffffffffff60501b191660509190911b69ffffffffffffffffffff60501b16178255565b6136e86136c8604087015163ffffffff1690565b825463ffffffff60a01b191660a09190911b63ffffffff60a01b16178255565b60608501516001820155019261371b613707608083015162ffffff1690565b855462ffffff191662ffffff909116178555565b61375a61373260a08301516001600160401b031690565b85546affffffffffffffff000000191660189190911b6affffffffffffffff00000016178555565b01516001600160401b031690565b815467ffffffffffffffff60581b191660589190911b67ffffffffffffffff60581b16179055565b6134d8565b604051906137a2826127a3565b8160c06000918281528260208201528260408201528260608201528260808201528260a08201520152565b906040516137da816127a3565b60c06002829463ffffffff81546001600160501b0380821687528160501c16602087015260a01c16604085015260018101546060850152015462ffffff811660808401526001600160401b0390818160181c1660a085015260581c16910152565b6002600c541461384c576002600c55565b604051633ee5aeb560e01b8152600490fd5b9190820180921161142d57565b9190916001600160501b038080941691160191821161142d57565b9190916001600160501b038080941691160291821691820361142d57565b91909163ffffffff8080941691160191821161142d57565b6001600160401b03811161168d5760051b60200190565b92916138de826138bc565b916138ec604051938461280f565b829481845260208094019160051b810192831161045257905b8282106139125750505050565b81358152908301908301613905565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815260a08101918183106001600160401b0384111761168d5761283d9260405261398e565b9081602091031261045257516107a881610b1c565b6000806139b79260018060a01b03169360208151910182865af16139b0613a00565b9083613a30565b80519081151591826139e5575b50506139cd5750565b60249060405190635274afe760e01b82526004820152fd5b6139f89250602080918301019101613979565b1538806139c4565b3d15613a2b573d90613a1182612859565b91613a1f604051938461280f565b82523d6000602084013e565b606090565b90613a575750805115613a4557805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580613a8a575b613a68575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15613a60565b80518210156118e65760209160051b010190565b6000915b8151831015613af957613abe8383613a93565b5190600082821015613aea575060005260205260406000205b91600019811461142d5760010191613aab565b90604092825260205220613ad7565b91505090565b90816020910312610452575190565b6040513d6000823e3d90fd5b906040519063a9059cbb60e01b6020830152730b98151bedee73f9ba5f2c7b72dea02d38ce49fc602483015260448201526044815260808101918183106001600160401b0384111761168d5761283d9260405261398e565b60405163a9059cbb60e01b60208201526001600160a01b0392909216602483015260448083019390935291815261283d91613bae60648361280f565b61398e565b90600182811c92168015613be3575b6020831014613bcd57565b634e487b7160e01b600052602260045260246000fd5b91607f1691613bc2565b601f8111613bf9575050565b600090601182527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c68906020601f850160051c83019410613c54575b601f0160051c01915b828110613c4957505050565b818155600101613c3d565b9092508290613c34565b601f8111613c6a575050565b600090601382527f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a090906020601f850160051c83019410613cc5575b601f0160051c01915b828110613cba57505050565b818155600101613cae565b9092508290613ca5565b601f8111613cdb575050565b600090601282527fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec3444906020601f850160051c83019410613d36575b601f0160051c01915b828110613d2b57505050565b818155600101613d1f565b9092508290613d16565b90918060409360208452816020850152848401376000828201840152601f01601f1916010190565b6040519060008260115491613d7c83613bb3565b80835292600190818116908115613e025750600114613da3575b5061283d9250038361280f565b6011600090815291507f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c685b848310613de7575061283d935050810160200138613d96565b81935090816020925483858a01015201910190918592613dce565b90506020925061283d94915060ff191682840152151560051b82010138613d96565b90613e376020928281519485920161074f565b0190565b60125460009291613e4b82613bb3565b91600190818116908115613eb75750600114613e6657505050565b909192935060126000527fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec3444906000915b848310613ea4575050500190565b8181602092548587015201920191613e96565b60ff191683525050811515909102019150565b600d5490604882901c6001600160a01b031615613fa9576001600160a01b03811660009081526005602052604090819020546107a895911c6001600160401b0316916040519360208501953060601b87526001600160601b0319809360601b16603487015263ffffffff60e01b9060e01b16604886015260181b16604c8401526001600160401b0360c01b9060c01b166060830152466068830152608882015260888152613f77816127f4565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002090565b6040516353bd4fb360e11b8152600490fd5b600d546001600160a01b03949360489190911c851692613fda92613eca565b91613fe581846140b0565b506004819692961015610a54571594856140a4575b5050831561401e575b5050501561400d57565b60405162b7fad960e11b8152600490fd5b60009293509082916040516140578161121c6020820194630b135d3f60e11b998a87526024840152604060448401526064830190610772565b51915afa90614064613a00565b82614096575b8261407a575b5050388080614003565b9091506020818051810103126104525760200151143880614070565b91506020825110159161406a565b16821493503880613ffa565b81519190604183036140e1576140da92506020820151906060604084015193015160001a906140ec565b9192909190565b505060009160029190565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841161416457926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15610d9d5780516001600160a01b0381161561415b57918190565b50809160019190565b50505060009160039190565b6014549060005b8281106141905760405163e82a532960e01b8152600490fd5b614199816135e6565b506001600160401b038060028093015460181c16908085169182101592836141cb575b505050613af957600101614177565b909192506141d8846135e6565b50015460581c16113880806141bc565b6001600160401b0380600d5460081c1642039142831161142d57161061420a57565b6040516313634e8d60e11b8152600490fd5b6001600160401b039182169116101561423157565b604051631750215560e11b8152600490fd5b6001600160a01b03909116815260408101929161283d9160200190610a47565b61426b61345d565b600073721c00182a990771244d7a71b9fa2ea789a3b433803b61433b575b501561432957600a547fcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac906142fe906142ca906001600160a01b0316610ca1565b604080516001600160a01b03909216825273721c00182a990771244d7a71b9fa2ea789a3b433602083015290918291820190565b0390a1600a80546001600160a01b03191673721c00182a990771244d7a71b9fa2ea789a3b433179055565b6040516332483afb60e01b8152600490fd5b6040516301ffc9a760e01b81526000600482015290602090829060249082905afa829181614372575b501561428957905038614289565b61439491925060203d811161439b575b61438c818361280f565b810190613979565b9038614364565b503d614382565b6143aa61345d565b6000813b61444a575b6001600160a01b038216908115159081614441575b5061432957600a5461283d927fcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac91614408906001600160a01b0316610ca1565b604080516001600160a01b03928316815292909116602083015290a160018060a01b03166001600160601b0360a01b600a541617600a55565b905015386143c8565b6040516301ffc9a760e01b8152600060048201526020816024816001600160a01b0387165afa82918161448c575b50614484575b506143b3565b90503861447e565b6144a591925060203d811161439b5761438c818361280f565b9038614478565b604051906144b9826127be565b60006040838281528260208201520152565b908160609103126104525760408051916144e4836127be565b80516144ef8161191d565b835260208101516144ff81611927565b6020840152015161450f81611927565b604082015290565b61451f6144ac565b50600a54614535906001600160a01b0316610ca1565b6001600160a01b038116614561575061454c61284c565b60008152600060208201526000604082015290565b604051635caaa2a960e11b815230600482015290606090829060249082905afa908115610d9d57600091614593575090565b6107a8915060603d81116145b4575b6145ac818361280f565b8101906144cb565b503d6145a2565b6020908181840312610452578051906001600160401b03821161045257019180601f840112156104525782516145f0816138bc565b936145fe604051958661280f565b818552838086019260051b820101928311610452578301905b828210614625575050505090565b83809183516146338161055c565b815201910190614617565b60405161464a816127d9565b60008152906000368137565b600a5461466b906001600160a01b0316610ca1565b6001600160a01b03811661468257506107a861463e565b604051635caaa2a960e11b8152306004820152606081602481855afa918215610d9d576146c760206146f3946000948591614724575b5001516001600160781b031690565b604051633fe5df9960e01b81526001600160781b03909116600482015292839190829081906024820190565b03915afa908115610d9d57600091614709575090565b6107a8913d8091833e61471c818361280f565b8101906145bb565b61473c915060603d81116145b4576145ac818361280f565b386146b8565b600a54614757906001600160a01b0316610ca1565b6001600160a01b03811661476e57506107a861463e565b604051635caaa2a960e11b8152306004820152606081602481855afa918215610d9d576147b260406146f3946000948591614724575001516001600160781b031690565b6040516305fa529b60e21b81526001600160781b03909116600482015292839190829081906024820190565b600a546147f3906001600160a01b0316610ca1565b906001600160a01b038216614809575050600090565b604051635caaa2a960e11b815230600482015290606082602481865afa928315610d9d5761484e602061488b958195600091614724575001516001600160781b031690565b604051636b96ef2f60e11b81526001600160781b0390911660048201526001600160a01b0390921660248301529092839190829081906044820190565b03915afa908115610d9d576000916148a1575090565b6107a8915060203d811161439b5761438c818361280f565b600a546148ce906001600160a01b0316610ca1565b906001600160a01b0382166148e4575050600090565b604051635caaa2a960e11b815230600482015290606082602481865afa928315610d9d5761492a604061488b95602095600091614724575001516001600160781b031690565b6040516309445f5360e41b81526001600160781b0390911660048201526001600160a01b0390921660248301529092839190829081906044820190565b600a54919290916001600160a01b0316806149855750505050600190565b803b156104525760405163050bf71960e31b81526001600160a01b039384166004820152938316602485015291166044830152600090829060649082905afa90816149da575b506149d557600090565b600190565b6149e390612775565b386149cb565b6001600160a01b03168015614a125760005260056020526001600160401b036040600020541690565b6040516323d3ad8160e21b8152600490fd5b8060009081548110614a43575b604051636f96cda160e11b8152600490fd5b81526004906020918083526040928383205494600160e01b861615614a6a57505050614a31565b93929190935b8515614a7e57505050505090565b60001901808352818552838320549550614a70565b60005481109081614aa2575090565b90506000526004602052600160e01b604060002054161590565b9190614ac782614a24565b6001600160a01b03939084821690858116829003614c5e57600085815260066020526040902080549096909290614b116001600160a01b03861633908114908614171590565b1590565b614c1a575b8516918215614c085761283d978794614b30868989614c6f565b614bfe575b506001600160a01b03858116600090815260056020908152604080832080546000190190559289168252828220805460010190558682526004905220600160e11b904260a01b851782179055811615614bb4575b507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600080a4614d8a565b60018401614bcc816000526004602052604060002090565b5415614bd9575b50614b89565b6000548114614bd357614bf6906000526004602052604060002090565b553880614bd3565b6000905538614b35565b604051633a954ecd60e21b8152600490fd5b614c47614b0d614c40336111278960018060a01b03166000526007602052604060002090565b5460ff1690565b15614b1657604051632ce44b5f60e11b8152600490fd5b60405162a1148160e81b8152600490fd5b929190600090815b60019081811015614d3f57808501851161142d576001600160a01b0383811615888216158080614d38575b15614cb957604051635cbd944160e01b8152600490fd5b15614cc7575b505001614c77565b15614cd3575b80614cbf565b600a54168015614ccd57803b15611e3e5760405163050bf71960e31b81523360048201526001600160a01b03898116602483015285166044820152908590829060649082905afa8015610d9d5715614ccd5780611a5d614d3292612775565b38614ccd565b5081614ca2565b50505050509050565b9060005b838110614d595750505050565b808201821161142d576001600160a01b038316614d8257604051635cbd944160e01b8152600490fd5b600101614d4c565b60005b60019081811015614ddd57808501851161142d576001600160a01b03838116159081614dd2575b5015614dcc57604051635cbd944160e01b8152600490fd5b01614d8d565b905084161538614db4565b5050505050565b929190614df2828286614abc565b803b614dff575b50505050565b614e0893614f42565b15614e165738808080614df9565b6040516368d2bf6b60e11b8152600490fd5b9081602091031261045257516107a881610486565b6107a8939260809260018060a01b031682526000602083015260408201528160608201520190610772565b6001600160a01b0391821681529116602082015260408101919091526080606082018190526107a892910190610772565b614ec260209160009394604051948580948193630a85bd0160e11b998a84523360048501614e3d565b03926001600160a01b03165af160009181614f12575b50614f0457614ee5613a00565b80519081614eff576040516368d2bf6b60e11b8152600490fd5b602001fd5b6001600160e01b0319161490565b614f3491925060203d8111614f3b575b614f2c818361280f565b810190614e28565b9038614ed8565b503d614f22565b92602091614ec2936000604051809681958294630a85bd0160e11b9a8b85523360048601614e68565b60408051614f78816127d9565b60009384825284549381156150c757614f92828683614d48565b6001600160a01b0381166000908152600560205260409020805468010000000000000001840201905560008581526004602052604090206001600160a01b03821695600193914260a01b85841460e11b1788179055818101967fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef82828b838180a4858084015b8a81036150b857505050156150a8578161503591888a5584614d48565b813b615045575b50505050505050565b85039180805b615068575b5050505050508154036108895780808080808061503c565b1561509b575b86615080614b0d868487019686614e99565b61508a578161504b565b85516368d2bf6b60e11b8152600490fd5b85831061506e5780615050565b8551622e076360e81b8152600490fd5b80848d858180a4018690615018565b835163b562e8dd60e01b8152600490fd5b60405190608082018281106001600160401b0382111761168d5760405260006060838281528260208201528260408201520152565b6151156150d8565b5061511e6150d8565b60005482101561515657506151328161515b565b604081015161515657506151516107a89161514b6150d8565b50614a24565b615176565b905090565b6151636150d8565b5060005260046020526107a86040600020545b9061517f6150d8565b6001600160a01b038316815260a083901c6001600160401b03166020820152600160e01b83161515604082015260e89290921c6060830152565b906151c3826138bc565b6151d0604051918261280f565b82815280926151e1601f19916138bc565b0190602036910137565b908281101561530c576000918254808511615304575b5061520b816149e9565b848310156152fd578285038181106152f5575b505b615229816151b9565b9581156152ed576152398461510d565b91859460409361524e614b0d86830151151590565b6152db575b505b87811415806152d1575b156152c45761526d8161515b565b808501516152bb57516001600160a01b03908116806152b2575b50908160019287169088161461529e575b01615255565b806152ac838a01998c613a93565b52615298565b96506001615287565b50600190615298565b5050959450505050815290565b508187141561525f565b516001600160a01b0316955038615253565b945050505050565b90503861521e565b5082615220565b935038615201565b604051631960ccad60e11b8152600490fdfea264697066735822122058bf64a0ee6ac10b17b948abbd80ad044c6e0d0b24dbaf83404d5376d4b2acca64736f6c63430008140033",
}

// OlympicABI is the input ABI used to generate the binding from.
// Deprecated: Use OlympicMetaData.ABI instead.
var OlympicABI = OlympicMetaData.ABI

// OlympicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OlympicMetaData.Bin instead.
var OlympicBin = OlympicMetaData.Bin

// DeployOlympic deploys a new Ethereum contract, binding an instance of Olympic to it.
func DeployOlympic(auth *bind.TransactOpts, backend bind.ContractBackend, collectionName string, collectionSymbol string, tokenURISuffix string, maxMintableSupply *big.Int, globalWalletLimit *big.Int, cosigner common.Address, timestampExpirySeconds uint64, mintCurrency common.Address, fundReceiver common.Address, royaltyReceiver common.Address, royaltyFeeNumerator *big.Int) (common.Address, *types.Transaction, *Olympic, error) {
	parsed, err := OlympicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OlympicBin), backend, collectionName, collectionSymbol, tokenURISuffix, maxMintableSupply, globalWalletLimit, cosigner, timestampExpirySeconds, mintCurrency, fundReceiver, royaltyReceiver, royaltyFeeNumerator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Olympic{OlympicCaller: OlympicCaller{contract: contract}, OlympicTransactor: OlympicTransactor{contract: contract}, OlympicFilterer: OlympicFilterer{contract: contract}}, nil
}

// Olympic is an auto generated Go binding around an Ethereum contract.
type Olympic struct {
	OlympicCaller     // Read-only binding to the contract
	OlympicTransactor // Write-only binding to the contract
	OlympicFilterer   // Log filterer for contract events
}

// OlympicCaller is an auto generated read-only Go binding around an Ethereum contract.
type OlympicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OlympicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OlympicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OlympicSession struct {
	Contract     *Olympic          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OlympicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OlympicCallerSession struct {
	Contract *OlympicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OlympicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OlympicTransactorSession struct {
	Contract     *OlympicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OlympicRaw is an auto generated low-level Go binding around an Ethereum contract.
type OlympicRaw struct {
	Contract *Olympic // Generic contract binding to access the raw methods on
}

// OlympicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OlympicCallerRaw struct {
	Contract *OlympicCaller // Generic read-only contract binding to access the raw methods on
}

// OlympicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OlympicTransactorRaw struct {
	Contract *OlympicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOlympic creates a new instance of Olympic, bound to a specific deployed contract.
func NewOlympic(address common.Address, backend bind.ContractBackend) (*Olympic, error) {
	contract, err := bindOlympic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Olympic{OlympicCaller: OlympicCaller{contract: contract}, OlympicTransactor: OlympicTransactor{contract: contract}, OlympicFilterer: OlympicFilterer{contract: contract}}, nil
}

// NewOlympicCaller creates a new read-only instance of Olympic, bound to a specific deployed contract.
func NewOlympicCaller(address common.Address, caller bind.ContractCaller) (*OlympicCaller, error) {
	contract, err := bindOlympic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OlympicCaller{contract: contract}, nil
}

// NewOlympicTransactor creates a new write-only instance of Olympic, bound to a specific deployed contract.
func NewOlympicTransactor(address common.Address, transactor bind.ContractTransactor) (*OlympicTransactor, error) {
	contract, err := bindOlympic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OlympicTransactor{contract: contract}, nil
}

// NewOlympicFilterer creates a new log filterer instance of Olympic, bound to a specific deployed contract.
func NewOlympicFilterer(address common.Address, filterer bind.ContractFilterer) (*OlympicFilterer, error) {
	contract, err := bindOlympic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OlympicFilterer{contract: contract}, nil
}

// bindOlympic binds a generic wrapper to an already deployed contract.
func bindOlympic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OlympicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Olympic *OlympicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Olympic.Contract.OlympicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Olympic *OlympicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.Contract.OlympicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Olympic *OlympicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Olympic.Contract.OlympicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Olympic *OlympicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Olympic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Olympic *OlympicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Olympic *OlympicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Olympic.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Olympic *OlympicCaller) DEFAULTOPERATORWHITELISTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "DEFAULT_OPERATOR_WHITELIST_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Olympic *OlympicSession) DEFAULTOPERATORWHITELISTID() (*big.Int, error) {
	return _Olympic.Contract.DEFAULTOPERATORWHITELISTID(&_Olympic.CallOpts)
}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Olympic *OlympicCallerSession) DEFAULTOPERATORWHITELISTID() (*big.Int, error) {
	return _Olympic.Contract.DEFAULTOPERATORWHITELISTID(&_Olympic.CallOpts)
}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Olympic *OlympicCaller) DEFAULTTRANSFERSECURITYLEVEL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "DEFAULT_TRANSFER_SECURITY_LEVEL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Olympic *OlympicSession) DEFAULTTRANSFERSECURITYLEVEL() (uint8, error) {
	return _Olympic.Contract.DEFAULTTRANSFERSECURITYLEVEL(&_Olympic.CallOpts)
}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Olympic *OlympicCallerSession) DEFAULTTRANSFERSECURITYLEVEL() (uint8, error) {
	return _Olympic.Contract.DEFAULTTRANSFERSECURITYLEVEL(&_Olympic.CallOpts)
}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Olympic *OlympicCaller) DEFAULTTRANSFERVALIDATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "DEFAULT_TRANSFER_VALIDATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Olympic *OlympicSession) DEFAULTTRANSFERVALIDATOR() (common.Address, error) {
	return _Olympic.Contract.DEFAULTTRANSFERVALIDATOR(&_Olympic.CallOpts)
}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Olympic *OlympicCallerSession) DEFAULTTRANSFERVALIDATOR() (common.Address, error) {
	return _Olympic.Contract.DEFAULTTRANSFERVALIDATOR(&_Olympic.CallOpts)
}

// FUNDRECEIVER is a free data retrieval call binding the contract method 0x700d19f2.
//
// Solidity: function FUND_RECEIVER() view returns(address)
func (_Olympic *OlympicCaller) FUNDRECEIVER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "FUND_RECEIVER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNDRECEIVER is a free data retrieval call binding the contract method 0x700d19f2.
//
// Solidity: function FUND_RECEIVER() view returns(address)
func (_Olympic *OlympicSession) FUNDRECEIVER() (common.Address, error) {
	return _Olympic.Contract.FUNDRECEIVER(&_Olympic.CallOpts)
}

// FUNDRECEIVER is a free data retrieval call binding the contract method 0x700d19f2.
//
// Solidity: function FUND_RECEIVER() view returns(address)
func (_Olympic *OlympicCallerSession) FUNDRECEIVER() (common.Address, error) {
	return _Olympic.Contract.FUNDRECEIVER(&_Olympic.CallOpts)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Olympic *OlympicCaller) AssertValidCosign(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "assertValidCosign", minter, qty, timestamp, signature)

	if err != nil {
		return err
	}

	return err

}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Olympic *OlympicSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Olympic.Contract.AssertValidCosign(&_Olympic.CallOpts, minter, qty, timestamp, signature)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Olympic *OlympicCallerSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Olympic.Contract.AssertValidCosign(&_Olympic.CallOpts, minter, qty, timestamp, signature)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Olympic *OlympicCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Olympic *OlympicSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Olympic.Contract.BalanceOf(&_Olympic.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Olympic *OlympicCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Olympic.Contract.BalanceOf(&_Olympic.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Olympic *OlympicCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Olympic *OlympicSession) ContractURI() (string, error) {
	return _Olympic.Contract.ContractURI(&_Olympic.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Olympic *OlympicCallerSession) ContractURI() (string, error) {
	return _Olympic.Contract.ContractURI(&_Olympic.CallOpts)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Olympic *OlympicCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721ATokenOwnership)).(*IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Olympic *OlympicSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Olympic.Contract.ExplicitOwnershipOf(&_Olympic.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Olympic *OlympicCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Olympic.Contract.ExplicitOwnershipOf(&_Olympic.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Olympic *OlympicCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721ATokenOwnership)).(*[]IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Olympic *OlympicSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Olympic.Contract.ExplicitOwnershipsOf(&_Olympic.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Olympic *OlympicCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Olympic.Contract.ExplicitOwnershipsOf(&_Olympic.CallOpts, tokenIds)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Olympic *OlympicCaller) GetActiveStageFromTimestamp(opts *bind.CallOpts, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getActiveStageFromTimestamp", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Olympic *OlympicSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Olympic.Contract.GetActiveStageFromTimestamp(&_Olympic.CallOpts, timestamp)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Olympic *OlympicCallerSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Olympic.Contract.GetActiveStageFromTimestamp(&_Olympic.CallOpts, timestamp)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Olympic *OlympicCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Olympic *OlympicSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Olympic.Contract.GetApproved(&_Olympic.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Olympic *OlympicCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Olympic.Contract.GetApproved(&_Olympic.CallOpts, tokenId)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Olympic *OlympicCaller) GetCosignDigest(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getCosignDigest", minter, qty, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Olympic *OlympicSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Olympic.Contract.GetCosignDigest(&_Olympic.CallOpts, minter, qty, timestamp)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Olympic *OlympicCallerSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Olympic.Contract.GetCosignDigest(&_Olympic.CallOpts, minter, qty, timestamp)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Olympic *OlympicCaller) GetCosignNonce(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getCosignNonce", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Olympic *OlympicSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Olympic.Contract.GetCosignNonce(&_Olympic.CallOpts, minter)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Olympic *OlympicCallerSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Olympic.Contract.GetCosignNonce(&_Olympic.CallOpts, minter)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Olympic *OlympicCaller) GetGlobalWalletLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getGlobalWalletLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Olympic *OlympicSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Olympic.Contract.GetGlobalWalletLimit(&_Olympic.CallOpts)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Olympic *OlympicCallerSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Olympic.Contract.GetGlobalWalletLimit(&_Olympic.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Olympic *OlympicCaller) GetMaxMintableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getMaxMintableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Olympic *OlympicSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Olympic.Contract.GetMaxMintableSupply(&_Olympic.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Olympic *OlympicCallerSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Olympic.Contract.GetMaxMintableSupply(&_Olympic.CallOpts)
}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Olympic *OlympicCaller) GetMintCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getMintCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Olympic *OlympicSession) GetMintCurrency() (common.Address, error) {
	return _Olympic.Contract.GetMintCurrency(&_Olympic.CallOpts)
}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Olympic *OlympicCallerSession) GetMintCurrency() (common.Address, error) {
	return _Olympic.Contract.GetMintCurrency(&_Olympic.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Olympic *OlympicCaller) GetMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Olympic *OlympicSession) GetMintable() (bool, error) {
	return _Olympic.Contract.GetMintable(&_Olympic.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Olympic *OlympicCallerSession) GetMintable() (bool, error) {
	return _Olympic.Contract.GetMintable(&_Olympic.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Olympic *OlympicCaller) GetNumberStages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getNumberStages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Olympic *OlympicSession) GetNumberStages() (*big.Int, error) {
	return _Olympic.Contract.GetNumberStages(&_Olympic.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Olympic *OlympicCallerSession) GetNumberStages() (*big.Int, error) {
	return _Olympic.Contract.GetNumberStages(&_Olympic.CallOpts)
}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Olympic *OlympicCaller) GetPermittedContractReceivers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getPermittedContractReceivers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Olympic *OlympicSession) GetPermittedContractReceivers() ([]common.Address, error) {
	return _Olympic.Contract.GetPermittedContractReceivers(&_Olympic.CallOpts)
}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Olympic *OlympicCallerSession) GetPermittedContractReceivers() ([]common.Address, error) {
	return _Olympic.Contract.GetPermittedContractReceivers(&_Olympic.CallOpts)
}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Olympic *OlympicCaller) GetSecurityPolicy(opts *bind.CallOpts) (CollectionSecurityPolicy, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getSecurityPolicy")

	if err != nil {
		return *new(CollectionSecurityPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(CollectionSecurityPolicy)).(*CollectionSecurityPolicy)

	return out0, err

}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Olympic *OlympicSession) GetSecurityPolicy() (CollectionSecurityPolicy, error) {
	return _Olympic.Contract.GetSecurityPolicy(&_Olympic.CallOpts)
}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Olympic *OlympicCallerSession) GetSecurityPolicy() (CollectionSecurityPolicy, error) {
	return _Olympic.Contract.GetSecurityPolicy(&_Olympic.CallOpts)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Olympic *OlympicCaller) GetStageInfo(opts *bind.CallOpts, index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getStageInfo", index)

	if err != nil {
		return *new(IERC721MMintStageInfo), *new(uint32), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721MMintStageInfo)).(*IERC721MMintStageInfo)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Olympic *OlympicSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Olympic.Contract.GetStageInfo(&_Olympic.CallOpts, index)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Olympic *OlympicCallerSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Olympic.Contract.GetStageInfo(&_Olympic.CallOpts, index)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Olympic *OlympicCaller) GetTransferValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getTransferValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Olympic *OlympicSession) GetTransferValidator() (common.Address, error) {
	return _Olympic.Contract.GetTransferValidator(&_Olympic.CallOpts)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Olympic *OlympicCallerSession) GetTransferValidator() (common.Address, error) {
	return _Olympic.Contract.GetTransferValidator(&_Olympic.CallOpts)
}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Olympic *OlympicCaller) GetWhitelistedOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "getWhitelistedOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Olympic *OlympicSession) GetWhitelistedOperators() ([]common.Address, error) {
	return _Olympic.Contract.GetWhitelistedOperators(&_Olympic.CallOpts)
}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Olympic *OlympicCallerSession) GetWhitelistedOperators() ([]common.Address, error) {
	return _Olympic.Contract.GetWhitelistedOperators(&_Olympic.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Olympic *OlympicCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Olympic *OlympicSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Olympic.Contract.IsApprovedForAll(&_Olympic.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Olympic *OlympicCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Olympic.Contract.IsApprovedForAll(&_Olympic.CallOpts, owner, operator)
}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Olympic *OlympicCaller) IsContractReceiverPermitted(opts *bind.CallOpts, receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "isContractReceiverPermitted", receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Olympic *OlympicSession) IsContractReceiverPermitted(receiver common.Address) (bool, error) {
	return _Olympic.Contract.IsContractReceiverPermitted(&_Olympic.CallOpts, receiver)
}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Olympic *OlympicCallerSession) IsContractReceiverPermitted(receiver common.Address) (bool, error) {
	return _Olympic.Contract.IsContractReceiverPermitted(&_Olympic.CallOpts, receiver)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Olympic *OlympicCaller) IsOperatorWhitelisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "isOperatorWhitelisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Olympic *OlympicSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _Olympic.Contract.IsOperatorWhitelisted(&_Olympic.CallOpts, operator)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Olympic *OlympicCallerSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _Olympic.Contract.IsOperatorWhitelisted(&_Olympic.CallOpts, operator)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Olympic *OlympicCaller) IsTransferAllowed(opts *bind.CallOpts, caller common.Address, from common.Address, to common.Address) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "isTransferAllowed", caller, from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Olympic *OlympicSession) IsTransferAllowed(caller common.Address, from common.Address, to common.Address) (bool, error) {
	return _Olympic.Contract.IsTransferAllowed(&_Olympic.CallOpts, caller, from, to)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Olympic *OlympicCallerSession) IsTransferAllowed(caller common.Address, from common.Address, to common.Address) (bool, error) {
	return _Olympic.Contract.IsTransferAllowed(&_Olympic.CallOpts, caller, from, to)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Olympic *OlympicCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Olympic *OlympicSession) Name() (string, error) {
	return _Olympic.Contract.Name(&_Olympic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Olympic *OlympicCallerSession) Name() (string, error) {
	return _Olympic.Contract.Name(&_Olympic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Olympic *OlympicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Olympic *OlympicSession) Owner() (common.Address, error) {
	return _Olympic.Contract.Owner(&_Olympic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Olympic *OlympicCallerSession) Owner() (common.Address, error) {
	return _Olympic.Contract.Owner(&_Olympic.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Olympic *OlympicCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Olympic *OlympicSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Olympic.Contract.OwnerOf(&_Olympic.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Olympic *OlympicCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Olympic.Contract.OwnerOf(&_Olympic.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Olympic *OlympicCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Olympic *OlympicSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Olympic.Contract.RoyaltyInfo(&_Olympic.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Olympic *OlympicCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Olympic.Contract.RoyaltyInfo(&_Olympic.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Olympic *OlympicCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Olympic *OlympicSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Olympic.Contract.SupportsInterface(&_Olympic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Olympic *OlympicCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Olympic.Contract.SupportsInterface(&_Olympic.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Olympic *OlympicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Olympic *OlympicSession) Symbol() (string, error) {
	return _Olympic.Contract.Symbol(&_Olympic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Olympic *OlympicCallerSession) Symbol() (string, error) {
	return _Olympic.Contract.Symbol(&_Olympic.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Olympic *OlympicCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Olympic *OlympicSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Olympic.Contract.TokenURI(&_Olympic.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Olympic *OlympicCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Olympic.Contract.TokenURI(&_Olympic.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Olympic *OlympicCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Olympic *OlympicSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Olympic.Contract.TokensOfOwner(&_Olympic.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Olympic *OlympicCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Olympic.Contract.TokensOfOwner(&_Olympic.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Olympic *OlympicCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Olympic *OlympicSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Olympic.Contract.TokensOfOwnerIn(&_Olympic.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Olympic *OlympicCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Olympic.Contract.TokensOfOwnerIn(&_Olympic.CallOpts, owner, start, stop)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Olympic *OlympicCaller) TotalMintedByAddress(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "totalMintedByAddress", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Olympic *OlympicSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Olympic.Contract.TotalMintedByAddress(&_Olympic.CallOpts, a)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Olympic *OlympicCallerSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Olympic.Contract.TotalMintedByAddress(&_Olympic.CallOpts, a)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Olympic *OlympicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Olympic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Olympic *OlympicSession) TotalSupply() (*big.Int, error) {
	return _Olympic.Contract.TotalSupply(&_Olympic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Olympic *OlympicCallerSession) TotalSupply() (*big.Int, error) {
	return _Olympic.Contract.TotalSupply(&_Olympic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.Approve(&_Olympic.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.Approve(&_Olympic.TransactOpts, to, tokenId)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactor) Crossmint(opts *bind.TransactOpts, qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "crossmint", qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.Crossmint(&_Olympic.TransactOpts, qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactorSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.Crossmint(&_Olympic.TransactOpts, qty, to, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactor) Mint(opts *bind.TransactOpts, qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "mint", qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.Mint(&_Olympic.TransactOpts, qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactorSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.Mint(&_Olympic.TransactOpts, qty, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactor) MintWithLimit(opts *bind.TransactOpts, qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "mintWithLimit", qty, limit, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicSession) MintWithLimit(qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.MintWithLimit(&_Olympic.TransactOpts, qty, limit, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Olympic *OlympicTransactorSession) MintWithLimit(qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Olympic.Contract.MintWithLimit(&_Olympic.TransactOpts, qty, limit, proof, timestamp, signature)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Olympic *OlympicTransactor) OwnerMint(opts *bind.TransactOpts, qty uint32, to common.Address) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "ownerMint", qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Olympic *OlympicSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.OwnerMint(&_Olympic.TransactOpts, qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Olympic *OlympicTransactorSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.OwnerMint(&_Olympic.TransactOpts, qty, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Olympic *OlympicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Olympic *OlympicSession) RenounceOwnership() (*types.Transaction, error) {
	return _Olympic.Contract.RenounceOwnership(&_Olympic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Olympic *OlympicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Olympic.Contract.RenounceOwnership(&_Olympic.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SafeTransferFrom(&_Olympic.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SafeTransferFrom(&_Olympic.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Olympic *OlympicTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Olympic *OlympicSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Olympic.Contract.SafeTransferFrom0(&_Olympic.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Olympic *OlympicTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Olympic.Contract.SafeTransferFrom0(&_Olympic.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Olympic *OlympicTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Olympic *OlympicSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Olympic.Contract.SetApprovalForAll(&_Olympic.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Olympic *OlympicTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Olympic.Contract.SetApprovalForAll(&_Olympic.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Olympic *OlympicTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Olympic *OlympicSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Olympic.Contract.SetBaseURI(&_Olympic.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Olympic *OlympicTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Olympic.Contract.SetBaseURI(&_Olympic.TransactOpts, baseURI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Olympic *OlympicTransactor) SetContractURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setContractURI", uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Olympic *OlympicSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Olympic.Contract.SetContractURI(&_Olympic.TransactOpts, uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Olympic *OlympicTransactorSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Olympic.Contract.SetContractURI(&_Olympic.TransactOpts, uri)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Olympic *OlympicTransactor) SetCosigner(opts *bind.TransactOpts, cosigner common.Address) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setCosigner", cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Olympic *OlympicSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetCosigner(&_Olympic.TransactOpts, cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Olympic *OlympicTransactorSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetCosigner(&_Olympic.TransactOpts, cosigner)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Olympic *OlympicTransactor) SetCrossmintAddress(opts *bind.TransactOpts, crossmintAddress common.Address) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setCrossmintAddress", crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Olympic *OlympicSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetCrossmintAddress(&_Olympic.TransactOpts, crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Olympic *OlympicTransactorSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetCrossmintAddress(&_Olympic.TransactOpts, crossmintAddress)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setDefaultRoyalty", receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetDefaultRoyalty(&_Olympic.TransactOpts, receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicTransactorSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetDefaultRoyalty(&_Olympic.TransactOpts, receiver, feeNumerator)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Olympic *OlympicTransactor) SetGlobalWalletLimit(opts *bind.TransactOpts, globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setGlobalWalletLimit", globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Olympic *OlympicSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetGlobalWalletLimit(&_Olympic.TransactOpts, globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Olympic *OlympicTransactorSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetGlobalWalletLimit(&_Olympic.TransactOpts, globalWalletLimit)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Olympic *OlympicTransactor) SetMaxMintableSupply(opts *bind.TransactOpts, maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setMaxMintableSupply", maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Olympic *OlympicSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetMaxMintableSupply(&_Olympic.TransactOpts, maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Olympic *OlympicTransactorSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetMaxMintableSupply(&_Olympic.TransactOpts, maxMintableSupply)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Olympic *OlympicTransactor) SetMintable(opts *bind.TransactOpts, mintable bool) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setMintable", mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Olympic *OlympicSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Olympic.Contract.SetMintable(&_Olympic.TransactOpts, mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Olympic *OlympicTransactorSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Olympic.Contract.SetMintable(&_Olympic.TransactOpts, mintable)
}

// SetStages is a paid mutator transaction binding the contract method 0xf830e8b8.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Olympic *OlympicTransactor) SetStages(opts *bind.TransactOpts, newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setStages", newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0xf830e8b8.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Olympic *OlympicSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Olympic.Contract.SetStages(&_Olympic.TransactOpts, newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0xf830e8b8.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Olympic *OlympicTransactorSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Olympic.Contract.SetStages(&_Olympic.TransactOpts, newStages)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Olympic *OlympicTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, expiry uint64) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setTimestampExpirySeconds", expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Olympic *OlympicSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Olympic.Contract.SetTimestampExpirySeconds(&_Olympic.TransactOpts, expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Olympic *OlympicTransactorSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Olympic.Contract.SetTimestampExpirySeconds(&_Olympic.TransactOpts, expiry)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicTransactor) SetToCustomSecurityPolicy(opts *bind.TransactOpts, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setToCustomSecurityPolicy", level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicSession) SetToCustomSecurityPolicy(level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetToCustomSecurityPolicy(&_Olympic.TransactOpts, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicTransactorSession) SetToCustomSecurityPolicy(level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetToCustomSecurityPolicy(&_Olympic.TransactOpts, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicTransactor) SetToCustomValidatorAndSecurityPolicy(opts *bind.TransactOpts, validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setToCustomValidatorAndSecurityPolicy", validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicSession) SetToCustomValidatorAndSecurityPolicy(validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetToCustomValidatorAndSecurityPolicy(&_Olympic.TransactOpts, validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Olympic *OlympicTransactorSession) SetToCustomValidatorAndSecurityPolicy(validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetToCustomValidatorAndSecurityPolicy(&_Olympic.TransactOpts, validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Olympic *OlympicTransactor) SetToDefaultSecurityPolicy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setToDefaultSecurityPolicy")
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Olympic *OlympicSession) SetToDefaultSecurityPolicy() (*types.Transaction, error) {
	return _Olympic.Contract.SetToDefaultSecurityPolicy(&_Olympic.TransactOpts)
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Olympic *OlympicTransactorSession) SetToDefaultSecurityPolicy() (*types.Transaction, error) {
	return _Olympic.Contract.SetToDefaultSecurityPolicy(&_Olympic.TransactOpts)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicTransactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetTokenRoyalty(&_Olympic.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Olympic *OlympicTransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.SetTokenRoyalty(&_Olympic.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Olympic *OlympicTransactor) SetTokenURISuffix(opts *bind.TransactOpts, suffix string) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setTokenURISuffix", suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Olympic *OlympicSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Olympic.Contract.SetTokenURISuffix(&_Olympic.TransactOpts, suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Olympic *OlympicTransactorSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Olympic.Contract.SetTokenURISuffix(&_Olympic.TransactOpts, suffix)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Olympic *OlympicTransactor) SetTransferValidator(opts *bind.TransactOpts, transferValidator_ common.Address) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "setTransferValidator", transferValidator_)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Olympic *OlympicSession) SetTransferValidator(transferValidator_ common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetTransferValidator(&_Olympic.TransactOpts, transferValidator_)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Olympic *OlympicTransactorSession) SetTransferValidator(transferValidator_ common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.SetTransferValidator(&_Olympic.TransactOpts, transferValidator_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.TransferFrom(&_Olympic.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Olympic *OlympicTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Olympic.Contract.TransferFrom(&_Olympic.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Olympic *OlympicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Olympic *OlympicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.TransferOwnership(&_Olympic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Olympic *OlympicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Olympic.Contract.TransferOwnership(&_Olympic.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Olympic *OlympicTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Olympic *OlympicSession) Withdraw() (*types.Transaction, error) {
	return _Olympic.Contract.Withdraw(&_Olympic.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Olympic *OlympicTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Olympic.Contract.Withdraw(&_Olympic.TransactOpts)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Olympic *OlympicTransactor) WithdrawERC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Olympic.contract.Transact(opts, "withdrawERC20")
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Olympic *OlympicSession) WithdrawERC20() (*types.Transaction, error) {
	return _Olympic.Contract.WithdrawERC20(&_Olympic.TransactOpts)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Olympic *OlympicTransactorSession) WithdrawERC20() (*types.Transaction, error) {
	return _Olympic.Contract.WithdrawERC20(&_Olympic.TransactOpts)
}

// OlympicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Olympic contract.
type OlympicApprovalIterator struct {
	Event *OlympicApproval // Event containing the contract specifics and raw log

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
func (it *OlympicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicApproval)
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
		it.Event = new(OlympicApproval)
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
func (it *OlympicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicApproval represents a Approval event raised by the Olympic contract.
type OlympicApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*OlympicApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OlympicApprovalIterator{contract: _Olympic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OlympicApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicApproval)
				if err := _Olympic.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) ParseApproval(log types.Log) (*OlympicApproval, error) {
	event := new(OlympicApproval)
	if err := _Olympic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Olympic contract.
type OlympicApprovalForAllIterator struct {
	Event *OlympicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OlympicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicApprovalForAll)
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
		it.Event = new(OlympicApprovalForAll)
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
func (it *OlympicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicApprovalForAll represents a ApprovalForAll event raised by the Olympic contract.
type OlympicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Olympic *OlympicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*OlympicApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OlympicApprovalForAllIterator{contract: _Olympic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Olympic *OlympicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OlympicApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicApprovalForAll)
				if err := _Olympic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Olympic *OlympicFilterer) ParseApprovalForAll(log types.Log) (*OlympicApprovalForAll, error) {
	event := new(OlympicApprovalForAll)
	if err := _Olympic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Olympic contract.
type OlympicConsecutiveTransferIterator struct {
	Event *OlympicConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *OlympicConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicConsecutiveTransfer)
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
		it.Event = new(OlympicConsecutiveTransfer)
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
func (it *OlympicConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Olympic contract.
type OlympicConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Olympic *OlympicFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*OlympicConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OlympicConsecutiveTransferIterator{contract: _Olympic.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Olympic *OlympicFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *OlympicConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicConsecutiveTransfer)
				if err := _Olympic.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Olympic *OlympicFilterer) ParseConsecutiveTransfer(log types.Log) (*OlympicConsecutiveTransfer, error) {
	event := new(OlympicConsecutiveTransfer)
	if err := _Olympic.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicDefaultRoyaltySetIterator is returned from FilterDefaultRoyaltySet and is used to iterate over the raw logs and unpacked data for DefaultRoyaltySet events raised by the Olympic contract.
type OlympicDefaultRoyaltySetIterator struct {
	Event *OlympicDefaultRoyaltySet // Event containing the contract specifics and raw log

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
func (it *OlympicDefaultRoyaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicDefaultRoyaltySet)
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
		it.Event = new(OlympicDefaultRoyaltySet)
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
func (it *OlympicDefaultRoyaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicDefaultRoyaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicDefaultRoyaltySet represents a DefaultRoyaltySet event raised by the Olympic contract.
type OlympicDefaultRoyaltySet struct {
	Receiver     common.Address
	FeeNumerator *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyaltySet is a free log retrieval operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) FilterDefaultRoyaltySet(opts *bind.FilterOpts, receiver []common.Address) (*OlympicDefaultRoyaltySetIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "DefaultRoyaltySet", receiverRule)
	if err != nil {
		return nil, err
	}
	return &OlympicDefaultRoyaltySetIterator{contract: _Olympic.contract, event: "DefaultRoyaltySet", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyaltySet is a free log subscription operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) WatchDefaultRoyaltySet(opts *bind.WatchOpts, sink chan<- *OlympicDefaultRoyaltySet, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "DefaultRoyaltySet", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicDefaultRoyaltySet)
				if err := _Olympic.contract.UnpackLog(event, "DefaultRoyaltySet", log); err != nil {
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

// ParseDefaultRoyaltySet is a log parse operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) ParseDefaultRoyaltySet(log types.Log) (*OlympicDefaultRoyaltySet, error) {
	event := new(OlympicDefaultRoyaltySet)
	if err := _Olympic.contract.UnpackLog(event, "DefaultRoyaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Olympic contract.
type OlympicOwnershipTransferredIterator struct {
	Event *OlympicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OlympicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicOwnershipTransferred)
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
		it.Event = new(OlympicOwnershipTransferred)
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
func (it *OlympicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicOwnershipTransferred represents a OwnershipTransferred event raised by the Olympic contract.
type OlympicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Olympic *OlympicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OlympicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OlympicOwnershipTransferredIterator{contract: _Olympic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Olympic *OlympicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OlympicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicOwnershipTransferred)
				if err := _Olympic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Olympic *OlympicFilterer) ParseOwnershipTransferred(log types.Log) (*OlympicOwnershipTransferred, error) {
	event := new(OlympicOwnershipTransferred)
	if err := _Olympic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetActiveStageIterator is returned from FilterSetActiveStage and is used to iterate over the raw logs and unpacked data for SetActiveStage events raised by the Olympic contract.
type OlympicSetActiveStageIterator struct {
	Event *OlympicSetActiveStage // Event containing the contract specifics and raw log

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
func (it *OlympicSetActiveStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetActiveStage)
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
		it.Event = new(OlympicSetActiveStage)
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
func (it *OlympicSetActiveStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetActiveStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetActiveStage represents a SetActiveStage event raised by the Olympic contract.
type OlympicSetActiveStage struct {
	ActiveStage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetActiveStage is a free log retrieval operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Olympic *OlympicFilterer) FilterSetActiveStage(opts *bind.FilterOpts) (*OlympicSetActiveStageIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return &OlympicSetActiveStageIterator{contract: _Olympic.contract, event: "SetActiveStage", logs: logs, sub: sub}, nil
}

// WatchSetActiveStage is a free log subscription operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Olympic *OlympicFilterer) WatchSetActiveStage(opts *bind.WatchOpts, sink chan<- *OlympicSetActiveStage) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetActiveStage)
				if err := _Olympic.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
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

// ParseSetActiveStage is a log parse operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Olympic *OlympicFilterer) ParseSetActiveStage(log types.Log) (*OlympicSetActiveStage, error) {
	event := new(OlympicSetActiveStage)
	if err := _Olympic.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetBaseURIIterator is returned from FilterSetBaseURI and is used to iterate over the raw logs and unpacked data for SetBaseURI events raised by the Olympic contract.
type OlympicSetBaseURIIterator struct {
	Event *OlympicSetBaseURI // Event containing the contract specifics and raw log

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
func (it *OlympicSetBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetBaseURI)
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
		it.Event = new(OlympicSetBaseURI)
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
func (it *OlympicSetBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetBaseURI represents a SetBaseURI event raised by the Olympic contract.
type OlympicSetBaseURI struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBaseURI is a free log retrieval operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Olympic *OlympicFilterer) FilterSetBaseURI(opts *bind.FilterOpts) (*OlympicSetBaseURIIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return &OlympicSetBaseURIIterator{contract: _Olympic.contract, event: "SetBaseURI", logs: logs, sub: sub}, nil
}

// WatchSetBaseURI is a free log subscription operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Olympic *OlympicFilterer) WatchSetBaseURI(opts *bind.WatchOpts, sink chan<- *OlympicSetBaseURI) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetBaseURI)
				if err := _Olympic.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
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

// ParseSetBaseURI is a log parse operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Olympic *OlympicFilterer) ParseSetBaseURI(log types.Log) (*OlympicSetBaseURI, error) {
	event := new(OlympicSetBaseURI)
	if err := _Olympic.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetCosignerIterator is returned from FilterSetCosigner and is used to iterate over the raw logs and unpacked data for SetCosigner events raised by the Olympic contract.
type OlympicSetCosignerIterator struct {
	Event *OlympicSetCosigner // Event containing the contract specifics and raw log

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
func (it *OlympicSetCosignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetCosigner)
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
		it.Event = new(OlympicSetCosigner)
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
func (it *OlympicSetCosignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetCosignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetCosigner represents a SetCosigner event raised by the Olympic contract.
type OlympicSetCosigner struct {
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCosigner is a free log retrieval operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Olympic *OlympicFilterer) FilterSetCosigner(opts *bind.FilterOpts) (*OlympicSetCosignerIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return &OlympicSetCosignerIterator{contract: _Olympic.contract, event: "SetCosigner", logs: logs, sub: sub}, nil
}

// WatchSetCosigner is a free log subscription operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Olympic *OlympicFilterer) WatchSetCosigner(opts *bind.WatchOpts, sink chan<- *OlympicSetCosigner) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetCosigner)
				if err := _Olympic.contract.UnpackLog(event, "SetCosigner", log); err != nil {
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

// ParseSetCosigner is a log parse operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Olympic *OlympicFilterer) ParseSetCosigner(log types.Log) (*OlympicSetCosigner, error) {
	event := new(OlympicSetCosigner)
	if err := _Olympic.contract.UnpackLog(event, "SetCosigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetCrossmintAddressIterator is returned from FilterSetCrossmintAddress and is used to iterate over the raw logs and unpacked data for SetCrossmintAddress events raised by the Olympic contract.
type OlympicSetCrossmintAddressIterator struct {
	Event *OlympicSetCrossmintAddress // Event containing the contract specifics and raw log

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
func (it *OlympicSetCrossmintAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetCrossmintAddress)
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
		it.Event = new(OlympicSetCrossmintAddress)
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
func (it *OlympicSetCrossmintAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetCrossmintAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetCrossmintAddress represents a SetCrossmintAddress event raised by the Olympic contract.
type OlympicSetCrossmintAddress struct {
	CrossmintAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCrossmintAddress is a free log retrieval operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Olympic *OlympicFilterer) FilterSetCrossmintAddress(opts *bind.FilterOpts) (*OlympicSetCrossmintAddressIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return &OlympicSetCrossmintAddressIterator{contract: _Olympic.contract, event: "SetCrossmintAddress", logs: logs, sub: sub}, nil
}

// WatchSetCrossmintAddress is a free log subscription operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Olympic *OlympicFilterer) WatchSetCrossmintAddress(opts *bind.WatchOpts, sink chan<- *OlympicSetCrossmintAddress) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetCrossmintAddress)
				if err := _Olympic.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
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

// ParseSetCrossmintAddress is a log parse operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Olympic *OlympicFilterer) ParseSetCrossmintAddress(log types.Log) (*OlympicSetCrossmintAddress, error) {
	event := new(OlympicSetCrossmintAddress)
	if err := _Olympic.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetGlobalWalletLimitIterator is returned from FilterSetGlobalWalletLimit and is used to iterate over the raw logs and unpacked data for SetGlobalWalletLimit events raised by the Olympic contract.
type OlympicSetGlobalWalletLimitIterator struct {
	Event *OlympicSetGlobalWalletLimit // Event containing the contract specifics and raw log

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
func (it *OlympicSetGlobalWalletLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetGlobalWalletLimit)
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
		it.Event = new(OlympicSetGlobalWalletLimit)
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
func (it *OlympicSetGlobalWalletLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetGlobalWalletLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetGlobalWalletLimit represents a SetGlobalWalletLimit event raised by the Olympic contract.
type OlympicSetGlobalWalletLimit struct {
	GlobalWalletLimit *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalWalletLimit is a free log retrieval operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Olympic *OlympicFilterer) FilterSetGlobalWalletLimit(opts *bind.FilterOpts) (*OlympicSetGlobalWalletLimitIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return &OlympicSetGlobalWalletLimitIterator{contract: _Olympic.contract, event: "SetGlobalWalletLimit", logs: logs, sub: sub}, nil
}

// WatchSetGlobalWalletLimit is a free log subscription operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Olympic *OlympicFilterer) WatchSetGlobalWalletLimit(opts *bind.WatchOpts, sink chan<- *OlympicSetGlobalWalletLimit) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetGlobalWalletLimit)
				if err := _Olympic.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
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

// ParseSetGlobalWalletLimit is a log parse operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Olympic *OlympicFilterer) ParseSetGlobalWalletLimit(log types.Log) (*OlympicSetGlobalWalletLimit, error) {
	event := new(OlympicSetGlobalWalletLimit)
	if err := _Olympic.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetMaxMintableSupplyIterator is returned from FilterSetMaxMintableSupply and is used to iterate over the raw logs and unpacked data for SetMaxMintableSupply events raised by the Olympic contract.
type OlympicSetMaxMintableSupplyIterator struct {
	Event *OlympicSetMaxMintableSupply // Event containing the contract specifics and raw log

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
func (it *OlympicSetMaxMintableSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetMaxMintableSupply)
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
		it.Event = new(OlympicSetMaxMintableSupply)
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
func (it *OlympicSetMaxMintableSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetMaxMintableSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetMaxMintableSupply represents a SetMaxMintableSupply event raised by the Olympic contract.
type OlympicSetMaxMintableSupply struct {
	MaxMintableSupply *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMaxMintableSupply is a free log retrieval operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Olympic *OlympicFilterer) FilterSetMaxMintableSupply(opts *bind.FilterOpts) (*OlympicSetMaxMintableSupplyIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return &OlympicSetMaxMintableSupplyIterator{contract: _Olympic.contract, event: "SetMaxMintableSupply", logs: logs, sub: sub}, nil
}

// WatchSetMaxMintableSupply is a free log subscription operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Olympic *OlympicFilterer) WatchSetMaxMintableSupply(opts *bind.WatchOpts, sink chan<- *OlympicSetMaxMintableSupply) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetMaxMintableSupply)
				if err := _Olympic.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
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

// ParseSetMaxMintableSupply is a log parse operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Olympic *OlympicFilterer) ParseSetMaxMintableSupply(log types.Log) (*OlympicSetMaxMintableSupply, error) {
	event := new(OlympicSetMaxMintableSupply)
	if err := _Olympic.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetMintCurrencyIterator is returned from FilterSetMintCurrency and is used to iterate over the raw logs and unpacked data for SetMintCurrency events raised by the Olympic contract.
type OlympicSetMintCurrencyIterator struct {
	Event *OlympicSetMintCurrency // Event containing the contract specifics and raw log

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
func (it *OlympicSetMintCurrencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetMintCurrency)
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
		it.Event = new(OlympicSetMintCurrency)
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
func (it *OlympicSetMintCurrencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetMintCurrencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetMintCurrency represents a SetMintCurrency event raised by the Olympic contract.
type OlympicSetMintCurrency struct {
	MintCurrency common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetMintCurrency is a free log retrieval operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Olympic *OlympicFilterer) FilterSetMintCurrency(opts *bind.FilterOpts) (*OlympicSetMintCurrencyIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetMintCurrency")
	if err != nil {
		return nil, err
	}
	return &OlympicSetMintCurrencyIterator{contract: _Olympic.contract, event: "SetMintCurrency", logs: logs, sub: sub}, nil
}

// WatchSetMintCurrency is a free log subscription operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Olympic *OlympicFilterer) WatchSetMintCurrency(opts *bind.WatchOpts, sink chan<- *OlympicSetMintCurrency) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetMintCurrency")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetMintCurrency)
				if err := _Olympic.contract.UnpackLog(event, "SetMintCurrency", log); err != nil {
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

// ParseSetMintCurrency is a log parse operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Olympic *OlympicFilterer) ParseSetMintCurrency(log types.Log) (*OlympicSetMintCurrency, error) {
	event := new(OlympicSetMintCurrency)
	if err := _Olympic.contract.UnpackLog(event, "SetMintCurrency", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetMintableIterator is returned from FilterSetMintable and is used to iterate over the raw logs and unpacked data for SetMintable events raised by the Olympic contract.
type OlympicSetMintableIterator struct {
	Event *OlympicSetMintable // Event containing the contract specifics and raw log

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
func (it *OlympicSetMintableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetMintable)
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
		it.Event = new(OlympicSetMintable)
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
func (it *OlympicSetMintableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetMintableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetMintable represents a SetMintable event raised by the Olympic contract.
type OlympicSetMintable struct {
	Mintable bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMintable is a free log retrieval operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Olympic *OlympicFilterer) FilterSetMintable(opts *bind.FilterOpts) (*OlympicSetMintableIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return &OlympicSetMintableIterator{contract: _Olympic.contract, event: "SetMintable", logs: logs, sub: sub}, nil
}

// WatchSetMintable is a free log subscription operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Olympic *OlympicFilterer) WatchSetMintable(opts *bind.WatchOpts, sink chan<- *OlympicSetMintable) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetMintable)
				if err := _Olympic.contract.UnpackLog(event, "SetMintable", log); err != nil {
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

// ParseSetMintable is a log parse operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Olympic *OlympicFilterer) ParseSetMintable(log types.Log) (*OlympicSetMintable, error) {
	event := new(OlympicSetMintable)
	if err := _Olympic.contract.UnpackLog(event, "SetMintable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicSetTimestampExpirySecondsIterator is returned from FilterSetTimestampExpirySeconds and is used to iterate over the raw logs and unpacked data for SetTimestampExpirySeconds events raised by the Olympic contract.
type OlympicSetTimestampExpirySecondsIterator struct {
	Event *OlympicSetTimestampExpirySeconds // Event containing the contract specifics and raw log

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
func (it *OlympicSetTimestampExpirySecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicSetTimestampExpirySeconds)
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
		it.Event = new(OlympicSetTimestampExpirySeconds)
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
func (it *OlympicSetTimestampExpirySecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicSetTimestampExpirySecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicSetTimestampExpirySeconds represents a SetTimestampExpirySeconds event raised by the Olympic contract.
type OlympicSetTimestampExpirySeconds struct {
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTimestampExpirySeconds is a free log retrieval operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Olympic *OlympicFilterer) FilterSetTimestampExpirySeconds(opts *bind.FilterOpts) (*OlympicSetTimestampExpirySecondsIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return &OlympicSetTimestampExpirySecondsIterator{contract: _Olympic.contract, event: "SetTimestampExpirySeconds", logs: logs, sub: sub}, nil
}

// WatchSetTimestampExpirySeconds is a free log subscription operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Olympic *OlympicFilterer) WatchSetTimestampExpirySeconds(opts *bind.WatchOpts, sink chan<- *OlympicSetTimestampExpirySeconds) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicSetTimestampExpirySeconds)
				if err := _Olympic.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
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

// ParseSetTimestampExpirySeconds is a log parse operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Olympic *OlympicFilterer) ParseSetTimestampExpirySeconds(log types.Log) (*OlympicSetTimestampExpirySeconds, error) {
	event := new(OlympicSetTimestampExpirySeconds)
	if err := _Olympic.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicTokenRoyaltySetIterator is returned from FilterTokenRoyaltySet and is used to iterate over the raw logs and unpacked data for TokenRoyaltySet events raised by the Olympic contract.
type OlympicTokenRoyaltySetIterator struct {
	Event *OlympicTokenRoyaltySet // Event containing the contract specifics and raw log

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
func (it *OlympicTokenRoyaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicTokenRoyaltySet)
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
		it.Event = new(OlympicTokenRoyaltySet)
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
func (it *OlympicTokenRoyaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicTokenRoyaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicTokenRoyaltySet represents a TokenRoyaltySet event raised by the Olympic contract.
type OlympicTokenRoyaltySet struct {
	TokenId      *big.Int
	Receiver     common.Address
	FeeNumerator *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRoyaltySet is a free log retrieval operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) FilterTokenRoyaltySet(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*OlympicTokenRoyaltySetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "TokenRoyaltySet", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &OlympicTokenRoyaltySetIterator{contract: _Olympic.contract, event: "TokenRoyaltySet", logs: logs, sub: sub}, nil
}

// WatchTokenRoyaltySet is a free log subscription operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) WatchTokenRoyaltySet(opts *bind.WatchOpts, sink chan<- *OlympicTokenRoyaltySet, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "TokenRoyaltySet", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicTokenRoyaltySet)
				if err := _Olympic.contract.UnpackLog(event, "TokenRoyaltySet", log); err != nil {
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

// ParseTokenRoyaltySet is a log parse operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Olympic *OlympicFilterer) ParseTokenRoyaltySet(log types.Log) (*OlympicTokenRoyaltySet, error) {
	event := new(OlympicTokenRoyaltySet)
	if err := _Olympic.contract.UnpackLog(event, "TokenRoyaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Olympic contract.
type OlympicTransferIterator struct {
	Event *OlympicTransfer // Event containing the contract specifics and raw log

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
func (it *OlympicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicTransfer)
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
		it.Event = new(OlympicTransfer)
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
func (it *OlympicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicTransfer represents a Transfer event raised by the Olympic contract.
type OlympicTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*OlympicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OlympicTransferIterator{contract: _Olympic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OlympicTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicTransfer)
				if err := _Olympic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Olympic *OlympicFilterer) ParseTransfer(log types.Log) (*OlympicTransfer, error) {
	event := new(OlympicTransfer)
	if err := _Olympic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicTransferValidatorUpdatedIterator is returned from FilterTransferValidatorUpdated and is used to iterate over the raw logs and unpacked data for TransferValidatorUpdated events raised by the Olympic contract.
type OlympicTransferValidatorUpdatedIterator struct {
	Event *OlympicTransferValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *OlympicTransferValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicTransferValidatorUpdated)
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
		it.Event = new(OlympicTransferValidatorUpdated)
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
func (it *OlympicTransferValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicTransferValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicTransferValidatorUpdated represents a TransferValidatorUpdated event raised by the Olympic contract.
type OlympicTransferValidatorUpdated struct {
	OldValidator common.Address
	NewValidator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferValidatorUpdated is a free log retrieval operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Olympic *OlympicFilterer) FilterTransferValidatorUpdated(opts *bind.FilterOpts) (*OlympicTransferValidatorUpdatedIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "TransferValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return &OlympicTransferValidatorUpdatedIterator{contract: _Olympic.contract, event: "TransferValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchTransferValidatorUpdated is a free log subscription operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Olympic *OlympicFilterer) WatchTransferValidatorUpdated(opts *bind.WatchOpts, sink chan<- *OlympicTransferValidatorUpdated) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "TransferValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicTransferValidatorUpdated)
				if err := _Olympic.contract.UnpackLog(event, "TransferValidatorUpdated", log); err != nil {
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

// ParseTransferValidatorUpdated is a log parse operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Olympic *OlympicFilterer) ParseTransferValidatorUpdated(log types.Log) (*OlympicTransferValidatorUpdated, error) {
	event := new(OlympicTransferValidatorUpdated)
	if err := _Olympic.contract.UnpackLog(event, "TransferValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicUpdateStageIterator is returned from FilterUpdateStage and is used to iterate over the raw logs and unpacked data for UpdateStage events raised by the Olympic contract.
type OlympicUpdateStageIterator struct {
	Event *OlympicUpdateStage // Event containing the contract specifics and raw log

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
func (it *OlympicUpdateStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicUpdateStage)
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
		it.Event = new(OlympicUpdateStage)
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
func (it *OlympicUpdateStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicUpdateStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicUpdateStage represents a UpdateStage event raised by the Olympic contract.
type OlympicUpdateStage struct {
	Stage                *big.Int
	Price                *big.Int
	MintFee              *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateStage is a free log retrieval operation binding the contract event 0xc4737822c84fe15fce8213ef237bb06d7d6c1603adfa65bf6d3a653195979092.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Olympic *OlympicFilterer) FilterUpdateStage(opts *bind.FilterOpts) (*OlympicUpdateStageIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return &OlympicUpdateStageIterator{contract: _Olympic.contract, event: "UpdateStage", logs: logs, sub: sub}, nil
}

// WatchUpdateStage is a free log subscription operation binding the contract event 0xc4737822c84fe15fce8213ef237bb06d7d6c1603adfa65bf6d3a653195979092.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Olympic *OlympicFilterer) WatchUpdateStage(opts *bind.WatchOpts, sink chan<- *OlympicUpdateStage) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicUpdateStage)
				if err := _Olympic.contract.UnpackLog(event, "UpdateStage", log); err != nil {
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

// ParseUpdateStage is a log parse operation binding the contract event 0xc4737822c84fe15fce8213ef237bb06d7d6c1603adfa65bf6d3a653195979092.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Olympic *OlympicFilterer) ParseUpdateStage(log types.Log) (*OlympicUpdateStage, error) {
	event := new(OlympicUpdateStage)
	if err := _Olympic.contract.UnpackLog(event, "UpdateStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Olympic contract.
type OlympicWithdrawIterator struct {
	Event *OlympicWithdraw // Event containing the contract specifics and raw log

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
func (it *OlympicWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicWithdraw)
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
		it.Event = new(OlympicWithdraw)
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
func (it *OlympicWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicWithdraw represents a Withdraw event raised by the Olympic contract.
type OlympicWithdraw struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Olympic *OlympicFilterer) FilterWithdraw(opts *bind.FilterOpts) (*OlympicWithdrawIterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &OlympicWithdrawIterator{contract: _Olympic.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Olympic *OlympicFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OlympicWithdraw) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicWithdraw)
				if err := _Olympic.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Olympic *OlympicFilterer) ParseWithdraw(log types.Log) (*OlympicWithdraw, error) {
	event := new(OlympicWithdraw)
	if err := _Olympic.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympicWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the Olympic contract.
type OlympicWithdrawERC20Iterator struct {
	Event *OlympicWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *OlympicWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympicWithdrawERC20)
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
		it.Event = new(OlympicWithdrawERC20)
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
func (it *OlympicWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympicWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympicWithdrawERC20 represents a WithdrawERC20 event raised by the Olympic contract.
type OlympicWithdrawERC20 struct {
	MintCurrency common.Address
	Value        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Olympic *OlympicFilterer) FilterWithdrawERC20(opts *bind.FilterOpts) (*OlympicWithdrawERC20Iterator, error) {

	logs, sub, err := _Olympic.contract.FilterLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return &OlympicWithdrawERC20Iterator{contract: _Olympic.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Olympic *OlympicFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *OlympicWithdrawERC20) (event.Subscription, error) {

	logs, sub, err := _Olympic.contract.WatchLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympicWithdrawERC20)
				if err := _Olympic.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Olympic *OlympicFilterer) ParseWithdrawERC20(log types.Log) (*OlympicWithdrawERC20, error) {
	event := new(OlympicWithdrawERC20)
	if err := _Olympic.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
