// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zorasecond

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

// ICreatorRoyaltiesControlRoyaltyConfiguration is an auto generated low-level Go binding around an user-defined struct.
type ICreatorRoyaltiesControlRoyaltyConfiguration struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}

// IZoraCreator1155TypesV1ContractConfig is an auto generated low-level Go binding around an user-defined struct.
type IZoraCreator1155TypesV1ContractConfig struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}

// IZoraCreator1155TypesV1TokenData is an auto generated low-level Go binding around an user-defined struct.
type IZoraCreator1155TypesV1TokenData struct {
	Uri         string
	MaxSupply   *big.Int
	TotalMinted *big.Int
}

// ZorasecondMetaData contains all meta data concerning the Zorasecond contract.
var ZorasecondMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mintFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_upgradeGate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mints\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESS_DELEGATECALL_TO_NON_CONTRACT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ADDRESS_LOW_LEVEL_CALL_FAILED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Burn_NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CREATOR_FUNDS_RECIPIENT_NOT_SET\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Call_TokenIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotZoraCreator1155\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotMintMoreTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"Config_TransferHookNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ACCOUNTS_AND_IDS_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ADDRESS_ZERO_IS_NOT_A_VALID_OWNER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_BURN_AMOUNT_EXCEEDS_BALANCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_BURN_FROM_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_CALLER_IS_NOT_TOKEN_OWNER_OR_APPROVED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ERC1155RECEIVER_REJECTED_TOKENS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_IDS_AND_AMOUNTS_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_INSUFFICIENT_BALANCE_FOR_TRANSFER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_MINT_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_MINT_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_SETTING_APPROVAL_FOR_SELF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_TRANSFER_TO_NON_ERC1155RECEIVER_IMPLEMENTER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_TRANSFER_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_NEW_IMPL_NOT_CONTRACT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_NEW_IMPL_NOT_UUPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_UNSUPPORTED_PROXIABLEUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHWithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FUNCTION_MUST_BE_CALLED_THROUGH_ACTIVE_PROXY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FUNCTION_MUST_BE_CALLED_THROUGH_DELEGATECALL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstMinterAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractValue\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawInsolvent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INITIALIZABLE_CONTRACT_ALREADY_INITIALIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INITIALIZABLE_CONTRACT_IS_NOT_INITIALIZING\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ETH_AMOUNT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintTo\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintSchedule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintSchedule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPremintVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintNotYetStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_InsolventSaleTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_InvalidMintArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_TokenIDMintNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_UnknownCommand\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_ValueTransferFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinterContractAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinterContractDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerNeedsToBeAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoRendererForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonEthRedemption\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_CREATE_REFERRAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTransfersFromZoraMints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PremintDeleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolRewardsWithdrawFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"RendererNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Renderer_NotValidRendererContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SaleHasNotStarted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"Sale_CannotCallNonSalesContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TokenIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPS_UPGRADEABLE_MUST_NOT_BE_CALLED_THROUGH_DELEGATECALL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedAmount\",\"type\":\"uint256\"}],\"name\":\"UserExceedsMintLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"UserMissingRoleForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongValueSent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"premintSignerContractFailedToRecoverSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"premintSignerContractNotAContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIZoraCreator1155.ConfigUpdate\",\"name\":\"updateType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap1\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap2\",\"type\":\"uint96\"},{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap3\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structIZoraCreator1155TypesV1.ContractConfig\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ContractMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRenderer1155\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"ContractRendererUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"structHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domainName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"CreatorAttribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lastOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Purchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RendererUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"SetupNewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"}],\"name\":\"UpdatedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"UpdatedRoyalties\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIZoraCreator1155TypesV1.TokenData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"name\":\"UpdatedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRACT_BASE_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_FUNDS_MANAGER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_METADATA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_MINTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_SALES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBits\",\"type\":\"uint256\"}],\"name\":\"addPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastTokenId\",\"type\":\"uint256\"}],\"name\":\"assumeLastTokenIdMatches\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"batchBalances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIMinter1155\",\"name\":\"salesConfig\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"computeTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap1\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap2\",\"type\":\"uint96\"},{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap3\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"createReferrals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customRenderers\",\"outputs\":[{\"internalType\":\"contractIRenderer1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"premintConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"premintVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"firstMinter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"premintSignerContract\",\"type\":\"address\"}],\"name\":\"delegateSetupNewToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"delegatedTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstMinters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCreatorRewardRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCustomRenderer\",\"outputs\":[{\"internalType\":\"contractIRenderer1155\",\"name\":\"customRenderer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"}],\"internalType\":\"structIZoraCreator1155TypesV1.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"defaultRoyaltyConfiguration\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"setupActions\",\"type\":\"bytes[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"isAdminOrRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"metadataRendererContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMinter1155\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"rewardsRecipients\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"minterArguments\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"mintTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIMinter1155\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"rewardsRecipients\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"minterArguments\",\"type\":\"bytes\"}],\"name\":\"mintWithMints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMinter1155\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"minterArguments\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"}],\"name\":\"mintWithRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBits\",\"type\":\"uint256\"}],\"name\":\"removePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"name\":\"setFundsRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIRenderer1155\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"setTokenMetadataRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"}],\"name\":\"setTransferHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"setupNewToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"}],\"name\":\"setupNewTokenWithCreateReferral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportedPremintSignatureVersions\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"}],\"name\":\"updateContractMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"updateCreateReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"newConfiguration\",\"type\":\"tuple\"}],\"name\":\"updateRoyaltiesForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ZorasecondABI is the input ABI used to generate the binding from.
// Deprecated: Use ZorasecondMetaData.ABI instead.
var ZorasecondABI = ZorasecondMetaData.ABI

// Zorasecond is an auto generated Go binding around an Ethereum contract.
type Zorasecond struct {
	ZorasecondCaller     // Read-only binding to the contract
	ZorasecondTransactor // Write-only binding to the contract
	ZorasecondFilterer   // Log filterer for contract events
}

// ZorasecondCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZorasecondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorasecondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZorasecondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorasecondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZorasecondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZorasecondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZorasecondSession struct {
	Contract     *Zorasecond       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZorasecondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZorasecondCallerSession struct {
	Contract *ZorasecondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZorasecondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZorasecondTransactorSession struct {
	Contract     *ZorasecondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZorasecondRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZorasecondRaw struct {
	Contract *Zorasecond // Generic contract binding to access the raw methods on
}

// ZorasecondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZorasecondCallerRaw struct {
	Contract *ZorasecondCaller // Generic read-only contract binding to access the raw methods on
}

// ZorasecondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZorasecondTransactorRaw struct {
	Contract *ZorasecondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZorasecond creates a new instance of Zorasecond, bound to a specific deployed contract.
func NewZorasecond(address common.Address, backend bind.ContractBackend) (*Zorasecond, error) {
	contract, err := bindZorasecond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zorasecond{ZorasecondCaller: ZorasecondCaller{contract: contract}, ZorasecondTransactor: ZorasecondTransactor{contract: contract}, ZorasecondFilterer: ZorasecondFilterer{contract: contract}}, nil
}

// NewZorasecondCaller creates a new read-only instance of Zorasecond, bound to a specific deployed contract.
func NewZorasecondCaller(address common.Address, caller bind.ContractCaller) (*ZorasecondCaller, error) {
	contract, err := bindZorasecond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZorasecondCaller{contract: contract}, nil
}

// NewZorasecondTransactor creates a new write-only instance of Zorasecond, bound to a specific deployed contract.
func NewZorasecondTransactor(address common.Address, transactor bind.ContractTransactor) (*ZorasecondTransactor, error) {
	contract, err := bindZorasecond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZorasecondTransactor{contract: contract}, nil
}

// NewZorasecondFilterer creates a new log filterer instance of Zorasecond, bound to a specific deployed contract.
func NewZorasecondFilterer(address common.Address, filterer bind.ContractFilterer) (*ZorasecondFilterer, error) {
	contract, err := bindZorasecond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZorasecondFilterer{contract: contract}, nil
}

// bindZorasecond binds a generic wrapper to an already deployed contract.
func bindZorasecond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZorasecondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zorasecond *ZorasecondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zorasecond.Contract.ZorasecondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zorasecond *ZorasecondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorasecond.Contract.ZorasecondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zorasecond *ZorasecondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zorasecond.Contract.ZorasecondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zorasecond *ZorasecondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zorasecond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zorasecond *ZorasecondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorasecond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zorasecond *ZorasecondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zorasecond.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) CONTRACTBASEID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "CONTRACT_BASE_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zorasecond *ZorasecondSession) CONTRACTBASEID() (*big.Int, error) {
	return _Zorasecond.Contract.CONTRACTBASEID(&_Zorasecond.CallOpts)
}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) CONTRACTBASEID() (*big.Int, error) {
	return _Zorasecond.Contract.CONTRACTBASEID(&_Zorasecond.CallOpts)
}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) PERMISSIONBITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "PERMISSION_BIT_ADMIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zorasecond *ZorasecondSession) PERMISSIONBITADMIN() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITADMIN(&_Zorasecond.CallOpts)
}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) PERMISSIONBITADMIN() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITADMIN(&_Zorasecond.CallOpts)
}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) PERMISSIONBITFUNDSMANAGER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "PERMISSION_BIT_FUNDS_MANAGER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zorasecond *ZorasecondSession) PERMISSIONBITFUNDSMANAGER() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITFUNDSMANAGER(&_Zorasecond.CallOpts)
}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) PERMISSIONBITFUNDSMANAGER() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITFUNDSMANAGER(&_Zorasecond.CallOpts)
}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) PERMISSIONBITMETADATA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "PERMISSION_BIT_METADATA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zorasecond *ZorasecondSession) PERMISSIONBITMETADATA() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITMETADATA(&_Zorasecond.CallOpts)
}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) PERMISSIONBITMETADATA() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITMETADATA(&_Zorasecond.CallOpts)
}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) PERMISSIONBITMINTER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "PERMISSION_BIT_MINTER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zorasecond *ZorasecondSession) PERMISSIONBITMINTER() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITMINTER(&_Zorasecond.CallOpts)
}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) PERMISSIONBITMINTER() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITMINTER(&_Zorasecond.CallOpts)
}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) PERMISSIONBITSALES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "PERMISSION_BIT_SALES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zorasecond *ZorasecondSession) PERMISSIONBITSALES() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITSALES(&_Zorasecond.CallOpts)
}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) PERMISSIONBITSALES() (*big.Int, error) {
	return _Zorasecond.Contract.PERMISSIONBITSALES(&_Zorasecond.CallOpts)
}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zorasecond *ZorasecondCaller) AssumeLastTokenIdMatches(opts *bind.CallOpts, lastTokenId *big.Int) error {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "assumeLastTokenIdMatches", lastTokenId)

	if err != nil {
		return err
	}

	return err

}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zorasecond *ZorasecondSession) AssumeLastTokenIdMatches(lastTokenId *big.Int) error {
	return _Zorasecond.Contract.AssumeLastTokenIdMatches(&_Zorasecond.CallOpts, lastTokenId)
}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zorasecond *ZorasecondCallerSession) AssumeLastTokenIdMatches(lastTokenId *big.Int) error {
	return _Zorasecond.Contract.AssumeLastTokenIdMatches(&_Zorasecond.CallOpts, lastTokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zorasecond *ZorasecondCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zorasecond *ZorasecondSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Zorasecond.Contract.BalanceOf(&_Zorasecond.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Zorasecond.Contract.BalanceOf(&_Zorasecond.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zorasecond *ZorasecondCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zorasecond *ZorasecondSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Zorasecond.Contract.BalanceOfBatch(&_Zorasecond.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zorasecond *ZorasecondCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Zorasecond.Contract.BalanceOfBatch(&_Zorasecond.CallOpts, accounts, ids)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0xa457c673.
//
// Solidity: function computeTotalReward(uint256 mintPrice, uint256 quantity) pure returns(uint256)
func (_Zorasecond *ZorasecondCaller) ComputeTotalReward(opts *bind.CallOpts, mintPrice *big.Int, quantity *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "computeTotalReward", mintPrice, quantity)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeTotalReward is a free data retrieval call binding the contract method 0xa457c673.
//
// Solidity: function computeTotalReward(uint256 mintPrice, uint256 quantity) pure returns(uint256)
func (_Zorasecond *ZorasecondSession) ComputeTotalReward(mintPrice *big.Int, quantity *big.Int) (*big.Int, error) {
	return _Zorasecond.Contract.ComputeTotalReward(&_Zorasecond.CallOpts, mintPrice, quantity)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0xa457c673.
//
// Solidity: function computeTotalReward(uint256 mintPrice, uint256 quantity) pure returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) ComputeTotalReward(mintPrice *big.Int, quantity *big.Int) (*big.Int, error) {
	return _Zorasecond.Contract.ComputeTotalReward(&_Zorasecond.CallOpts, mintPrice, quantity)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zorasecond *ZorasecondCaller) Config(opts *bind.CallOpts) (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner          common.Address
		Gap1           *big.Int
		FundsRecipient common.Address
		Gap2           *big.Int
		TransferHook   common.Address
		Gap3           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Gap1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FundsRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Gap2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TransferHook = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Gap3 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zorasecond *ZorasecondSession) Config() (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	return _Zorasecond.Contract.Config(&_Zorasecond.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zorasecond *ZorasecondCallerSession) Config() (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	return _Zorasecond.Contract.Config(&_Zorasecond.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zorasecond *ZorasecondCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zorasecond *ZorasecondSession) ContractURI() (string, error) {
	return _Zorasecond.Contract.ContractURI(&_Zorasecond.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zorasecond *ZorasecondCallerSession) ContractURI() (string, error) {
	return _Zorasecond.Contract.ContractURI(&_Zorasecond.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorasecond *ZorasecondCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorasecond *ZorasecondSession) ContractVersion() (string, error) {
	return _Zorasecond.Contract.ContractVersion(&_Zorasecond.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zorasecond *ZorasecondCallerSession) ContractVersion() (string, error) {
	return _Zorasecond.Contract.ContractVersion(&_Zorasecond.CallOpts)
}

// CreateReferrals is a free data retrieval call binding the contract method 0x7dafae4d.
//
// Solidity: function createReferrals(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCaller) CreateReferrals(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "createReferrals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreateReferrals is a free data retrieval call binding the contract method 0x7dafae4d.
//
// Solidity: function createReferrals(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondSession) CreateReferrals(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.CreateReferrals(&_Zorasecond.CallOpts, arg0)
}

// CreateReferrals is a free data retrieval call binding the contract method 0x7dafae4d.
//
// Solidity: function createReferrals(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCallerSession) CreateReferrals(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.CreateReferrals(&_Zorasecond.CallOpts, arg0)
}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCaller) CustomRenderers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "customRenderers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondSession) CustomRenderers(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.CustomRenderers(&_Zorasecond.CallOpts, arg0)
}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCallerSession) CustomRenderers(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.CustomRenderers(&_Zorasecond.CallOpts, arg0)
}

// DelegatedTokenId is a free data retrieval call binding the contract method 0xbdd864f2.
//
// Solidity: function delegatedTokenId(uint32 ) view returns(uint256)
func (_Zorasecond *ZorasecondCaller) DelegatedTokenId(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "delegatedTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedTokenId is a free data retrieval call binding the contract method 0xbdd864f2.
//
// Solidity: function delegatedTokenId(uint32 ) view returns(uint256)
func (_Zorasecond *ZorasecondSession) DelegatedTokenId(arg0 uint32) (*big.Int, error) {
	return _Zorasecond.Contract.DelegatedTokenId(&_Zorasecond.CallOpts, arg0)
}

// DelegatedTokenId is a free data retrieval call binding the contract method 0xbdd864f2.
//
// Solidity: function delegatedTokenId(uint32 ) view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) DelegatedTokenId(arg0 uint32) (*big.Int, error) {
	return _Zorasecond.Contract.DelegatedTokenId(&_Zorasecond.CallOpts, arg0)
}

// FirstMinters is a free data retrieval call binding the contract method 0x9ebb8324.
//
// Solidity: function firstMinters(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCaller) FirstMinters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "firstMinters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstMinters is a free data retrieval call binding the contract method 0x9ebb8324.
//
// Solidity: function firstMinters(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondSession) FirstMinters(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.FirstMinters(&_Zorasecond.CallOpts, arg0)
}

// FirstMinters is a free data retrieval call binding the contract method 0x9ebb8324.
//
// Solidity: function firstMinters(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCallerSession) FirstMinters(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.FirstMinters(&_Zorasecond.CallOpts, arg0)
}

// GetCreatorRewardRecipient is a free data retrieval call binding the contract method 0x5e4e0404.
//
// Solidity: function getCreatorRewardRecipient(uint256 tokenId) view returns(address)
func (_Zorasecond *ZorasecondCaller) GetCreatorRewardRecipient(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "getCreatorRewardRecipient", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreatorRewardRecipient is a free data retrieval call binding the contract method 0x5e4e0404.
//
// Solidity: function getCreatorRewardRecipient(uint256 tokenId) view returns(address)
func (_Zorasecond *ZorasecondSession) GetCreatorRewardRecipient(tokenId *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.GetCreatorRewardRecipient(&_Zorasecond.CallOpts, tokenId)
}

// GetCreatorRewardRecipient is a free data retrieval call binding the contract method 0x5e4e0404.
//
// Solidity: function getCreatorRewardRecipient(uint256 tokenId) view returns(address)
func (_Zorasecond *ZorasecondCallerSession) GetCreatorRewardRecipient(tokenId *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.GetCreatorRewardRecipient(&_Zorasecond.CallOpts, tokenId)
}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zorasecond *ZorasecondCaller) GetCustomRenderer(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "getCustomRenderer", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zorasecond *ZorasecondSession) GetCustomRenderer(tokenId *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.GetCustomRenderer(&_Zorasecond.CallOpts, tokenId)
}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zorasecond *ZorasecondCallerSession) GetCustomRenderer(tokenId *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.GetCustomRenderer(&_Zorasecond.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zorasecond *ZorasecondCaller) GetRoyalties(opts *bind.CallOpts, tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "getRoyalties", tokenId)

	if err != nil {
		return *new(ICreatorRoyaltiesControlRoyaltyConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new(ICreatorRoyaltiesControlRoyaltyConfiguration)).(*ICreatorRoyaltiesControlRoyaltyConfiguration)

	return out0, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zorasecond *ZorasecondSession) GetRoyalties(tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	return _Zorasecond.Contract.GetRoyalties(&_Zorasecond.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zorasecond *ZorasecondCallerSession) GetRoyalties(tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	return _Zorasecond.Contract.GetRoyalties(&_Zorasecond.CallOpts, tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zorasecond *ZorasecondCaller) GetTokenInfo(opts *bind.CallOpts, tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "getTokenInfo", tokenId)

	if err != nil {
		return *new(IZoraCreator1155TypesV1TokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(IZoraCreator1155TypesV1TokenData)).(*IZoraCreator1155TypesV1TokenData)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zorasecond *ZorasecondSession) GetTokenInfo(tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	return _Zorasecond.Contract.GetTokenInfo(&_Zorasecond.CallOpts, tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zorasecond *ZorasecondCallerSession) GetTokenInfo(tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	return _Zorasecond.Contract.GetTokenInfo(&_Zorasecond.CallOpts, tokenId)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Zorasecond *ZorasecondCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Zorasecond *ZorasecondSession) Implementation() (common.Address, error) {
	return _Zorasecond.Contract.Implementation(&_Zorasecond.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Zorasecond *ZorasecondCallerSession) Implementation() (common.Address, error) {
	return _Zorasecond.Contract.Implementation(&_Zorasecond.CallOpts)
}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zorasecond *ZorasecondCaller) IsAdminOrRole(opts *bind.CallOpts, user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "isAdminOrRole", user, tokenId, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zorasecond *ZorasecondSession) IsAdminOrRole(user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	return _Zorasecond.Contract.IsAdminOrRole(&_Zorasecond.CallOpts, user, tokenId, role)
}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zorasecond *ZorasecondCallerSession) IsAdminOrRole(user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	return _Zorasecond.Contract.IsAdminOrRole(&_Zorasecond.CallOpts, user, tokenId, role)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zorasecond *ZorasecondCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zorasecond *ZorasecondSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Zorasecond.Contract.IsApprovedForAll(&_Zorasecond.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zorasecond *ZorasecondCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Zorasecond.Contract.IsApprovedForAll(&_Zorasecond.CallOpts, account, operator)
}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCaller) MetadataRendererContract(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "metadataRendererContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondSession) MetadataRendererContract(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.MetadataRendererContract(&_Zorasecond.CallOpts, arg0)
}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zorasecond *ZorasecondCallerSession) MetadataRendererContract(arg0 *big.Int) (common.Address, error) {
	return _Zorasecond.Contract.MetadataRendererContract(&_Zorasecond.CallOpts, arg0)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) MintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "mintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zorasecond *ZorasecondSession) MintFee() (*big.Int, error) {
	return _Zorasecond.Contract.MintFee(&_Zorasecond.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) MintFee() (*big.Int, error) {
	return _Zorasecond.Contract.MintFee(&_Zorasecond.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zorasecond *ZorasecondCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zorasecond *ZorasecondSession) Name() (string, error) {
	return _Zorasecond.Contract.Name(&_Zorasecond.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zorasecond *ZorasecondCallerSession) Name() (string, error) {
	return _Zorasecond.Contract.Name(&_Zorasecond.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zorasecond *ZorasecondCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zorasecond *ZorasecondSession) NextTokenId() (*big.Int, error) {
	return _Zorasecond.Contract.NextTokenId(&_Zorasecond.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) NextTokenId() (*big.Int, error) {
	return _Zorasecond.Contract.NextTokenId(&_Zorasecond.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Zorasecond.Contract.OnERC1155BatchReceived(&_Zorasecond.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Zorasecond.Contract.OnERC1155BatchReceived(&_Zorasecond.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Zorasecond.Contract.OnERC1155Received(&_Zorasecond.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_Zorasecond *ZorasecondCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Zorasecond.Contract.OnERC1155Received(&_Zorasecond.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorasecond *ZorasecondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorasecond *ZorasecondSession) Owner() (common.Address, error) {
	return _Zorasecond.Contract.Owner(&_Zorasecond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zorasecond *ZorasecondCallerSession) Owner() (common.Address, error) {
	return _Zorasecond.Contract.Owner(&_Zorasecond.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zorasecond *ZorasecondCaller) Permissions(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "permissions", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zorasecond *ZorasecondSession) Permissions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Zorasecond.Contract.Permissions(&_Zorasecond.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zorasecond *ZorasecondCallerSession) Permissions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Zorasecond.Contract.Permissions(&_Zorasecond.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorasecond *ZorasecondCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorasecond *ZorasecondSession) ProxiableUUID() ([32]byte, error) {
	return _Zorasecond.Contract.ProxiableUUID(&_Zorasecond.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zorasecond *ZorasecondCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Zorasecond.Contract.ProxiableUUID(&_Zorasecond.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zorasecond *ZorasecondCaller) Royalties(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "royalties", arg0)

	outstruct := new(struct {
		RoyaltyMintSchedule uint32
		RoyaltyBPS          uint32
		RoyaltyRecipient    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoyaltyMintSchedule = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RoyaltyBPS = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RoyaltyRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zorasecond *ZorasecondSession) Royalties(arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	return _Zorasecond.Contract.Royalties(&_Zorasecond.CallOpts, arg0)
}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zorasecond *ZorasecondCallerSession) Royalties(arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	return _Zorasecond.Contract.Royalties(&_Zorasecond.CallOpts, arg0)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zorasecond *ZorasecondCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zorasecond *ZorasecondSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zorasecond.Contract.RoyaltyInfo(&_Zorasecond.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zorasecond *ZorasecondCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zorasecond.Contract.RoyaltyInfo(&_Zorasecond.CallOpts, tokenId, salePrice)
}

// SupportedPremintSignatureVersions is a free data retrieval call binding the contract method 0xed788913.
//
// Solidity: function supportedPremintSignatureVersions() pure returns(string[])
func (_Zorasecond *ZorasecondCaller) SupportedPremintSignatureVersions(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "supportedPremintSignatureVersions")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// SupportedPremintSignatureVersions is a free data retrieval call binding the contract method 0xed788913.
//
// Solidity: function supportedPremintSignatureVersions() pure returns(string[])
func (_Zorasecond *ZorasecondSession) SupportedPremintSignatureVersions() ([]string, error) {
	return _Zorasecond.Contract.SupportedPremintSignatureVersions(&_Zorasecond.CallOpts)
}

// SupportedPremintSignatureVersions is a free data retrieval call binding the contract method 0xed788913.
//
// Solidity: function supportedPremintSignatureVersions() pure returns(string[])
func (_Zorasecond *ZorasecondCallerSession) SupportedPremintSignatureVersions() ([]string, error) {
	return _Zorasecond.Contract.SupportedPremintSignatureVersions(&_Zorasecond.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zorasecond *ZorasecondCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zorasecond *ZorasecondSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zorasecond.Contract.SupportsInterface(&_Zorasecond.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zorasecond *ZorasecondCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zorasecond.Contract.SupportsInterface(&_Zorasecond.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zorasecond *ZorasecondCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zorasecond *ZorasecondSession) Symbol() (string, error) {
	return _Zorasecond.Contract.Symbol(&_Zorasecond.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zorasecond *ZorasecondCallerSession) Symbol() (string, error) {
	return _Zorasecond.Contract.Symbol(&_Zorasecond.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zorasecond *ZorasecondCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Zorasecond.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zorasecond *ZorasecondSession) Uri(tokenId *big.Int) (string, error) {
	return _Zorasecond.Contract.Uri(&_Zorasecond.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zorasecond *ZorasecondCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Zorasecond.Contract.Uri(&_Zorasecond.CallOpts, tokenId)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondTransactor) AddPermission(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "addPermission", tokenId, user, permissionBits)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondSession) AddPermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.AddPermission(&_Zorasecond.TransactOpts, tokenId, user, permissionBits)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondTransactorSession) AddPermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.AddPermission(&_Zorasecond.TransactOpts, tokenId, user, permissionBits)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zorasecond *ZorasecondTransactor) AdminMint(opts *bind.TransactOpts, recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "adminMint", recipient, tokenId, quantity, data)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zorasecond *ZorasecondSession) AdminMint(recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.AdminMint(&_Zorasecond.TransactOpts, recipient, tokenId, quantity, data)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zorasecond *ZorasecondTransactorSession) AdminMint(recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.AdminMint(&_Zorasecond.TransactOpts, recipient, tokenId, quantity, data)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zorasecond *ZorasecondTransactor) BurnBatch(opts *bind.TransactOpts, from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "burnBatch", from, tokenIds, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zorasecond *ZorasecondSession) BurnBatch(from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.BurnBatch(&_Zorasecond.TransactOpts, from, tokenIds, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zorasecond *ZorasecondTransactorSession) BurnBatch(from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.BurnBatch(&_Zorasecond.TransactOpts, from, tokenIds, amounts)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zorasecond *ZorasecondTransactor) CallRenderer(opts *bind.TransactOpts, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "callRenderer", tokenId, data)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zorasecond *ZorasecondSession) CallRenderer(tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.CallRenderer(&_Zorasecond.TransactOpts, tokenId, data)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zorasecond *ZorasecondTransactorSession) CallRenderer(tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.CallRenderer(&_Zorasecond.TransactOpts, tokenId, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zorasecond *ZorasecondTransactor) CallSale(opts *bind.TransactOpts, tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "callSale", tokenId, salesConfig, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zorasecond *ZorasecondSession) CallSale(tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.CallSale(&_Zorasecond.TransactOpts, tokenId, salesConfig, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zorasecond *ZorasecondTransactorSession) CallSale(tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.CallSale(&_Zorasecond.TransactOpts, tokenId, salesConfig, data)
}

// DelegateSetupNewToken is a paid mutator transaction binding the contract method 0x709e537f.
//
// Solidity: function delegateSetupNewToken(bytes premintConfig, bytes32 premintVersion, bytes signature, address firstMinter, address premintSignerContract) returns(uint256 newTokenId)
func (_Zorasecond *ZorasecondTransactor) DelegateSetupNewToken(opts *bind.TransactOpts, premintConfig []byte, premintVersion [32]byte, signature []byte, firstMinter common.Address, premintSignerContract common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "delegateSetupNewToken", premintConfig, premintVersion, signature, firstMinter, premintSignerContract)
}

// DelegateSetupNewToken is a paid mutator transaction binding the contract method 0x709e537f.
//
// Solidity: function delegateSetupNewToken(bytes premintConfig, bytes32 premintVersion, bytes signature, address firstMinter, address premintSignerContract) returns(uint256 newTokenId)
func (_Zorasecond *ZorasecondSession) DelegateSetupNewToken(premintConfig []byte, premintVersion [32]byte, signature []byte, firstMinter common.Address, premintSignerContract common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.DelegateSetupNewToken(&_Zorasecond.TransactOpts, premintConfig, premintVersion, signature, firstMinter, premintSignerContract)
}

// DelegateSetupNewToken is a paid mutator transaction binding the contract method 0x709e537f.
//
// Solidity: function delegateSetupNewToken(bytes premintConfig, bytes32 premintVersion, bytes signature, address firstMinter, address premintSignerContract) returns(uint256 newTokenId)
func (_Zorasecond *ZorasecondTransactorSession) DelegateSetupNewToken(premintConfig []byte, premintVersion [32]byte, signature []byte, firstMinter common.Address, premintSignerContract common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.DelegateSetupNewToken(&_Zorasecond.TransactOpts, premintConfig, premintVersion, signature, firstMinter, premintSignerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zorasecond *ZorasecondTransactor) Initialize(opts *bind.TransactOpts, contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "initialize", contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zorasecond *ZorasecondSession) Initialize(contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Initialize(&_Zorasecond.TransactOpts, contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zorasecond *ZorasecondTransactorSession) Initialize(contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Initialize(&_Zorasecond.TransactOpts, contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Mint is a paid mutator transaction binding the contract method 0x359f1302.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, address[] rewardsRecipients, bytes minterArguments) payable returns()
func (_Zorasecond *ZorasecondTransactor) Mint(opts *bind.TransactOpts, minter common.Address, tokenId *big.Int, quantity *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "mint", minter, tokenId, quantity, rewardsRecipients, minterArguments)
}

// Mint is a paid mutator transaction binding the contract method 0x359f1302.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, address[] rewardsRecipients, bytes minterArguments) payable returns()
func (_Zorasecond *ZorasecondSession) Mint(minter common.Address, tokenId *big.Int, quantity *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Mint(&_Zorasecond.TransactOpts, minter, tokenId, quantity, rewardsRecipients, minterArguments)
}

// Mint is a paid mutator transaction binding the contract method 0x359f1302.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, address[] rewardsRecipients, bytes minterArguments) payable returns()
func (_Zorasecond *ZorasecondTransactorSession) Mint(minter common.Address, tokenId *big.Int, quantity *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Mint(&_Zorasecond.TransactOpts, minter, tokenId, quantity, rewardsRecipients, minterArguments)
}

// MintWithMints is a paid mutator transaction binding the contract method 0x086bbace.
//
// Solidity: function mintWithMints(uint256[] mintTokenIds, uint256[] quantities, address minter, uint256 tokenId, address[] rewardsRecipients, bytes minterArguments) payable returns(uint256 quantityMinted)
func (_Zorasecond *ZorasecondTransactor) MintWithMints(opts *bind.TransactOpts, mintTokenIds []*big.Int, quantities []*big.Int, minter common.Address, tokenId *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "mintWithMints", mintTokenIds, quantities, minter, tokenId, rewardsRecipients, minterArguments)
}

// MintWithMints is a paid mutator transaction binding the contract method 0x086bbace.
//
// Solidity: function mintWithMints(uint256[] mintTokenIds, uint256[] quantities, address minter, uint256 tokenId, address[] rewardsRecipients, bytes minterArguments) payable returns(uint256 quantityMinted)
func (_Zorasecond *ZorasecondSession) MintWithMints(mintTokenIds []*big.Int, quantities []*big.Int, minter common.Address, tokenId *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.MintWithMints(&_Zorasecond.TransactOpts, mintTokenIds, quantities, minter, tokenId, rewardsRecipients, minterArguments)
}

// MintWithMints is a paid mutator transaction binding the contract method 0x086bbace.
//
// Solidity: function mintWithMints(uint256[] mintTokenIds, uint256[] quantities, address minter, uint256 tokenId, address[] rewardsRecipients, bytes minterArguments) payable returns(uint256 quantityMinted)
func (_Zorasecond *ZorasecondTransactorSession) MintWithMints(mintTokenIds []*big.Int, quantities []*big.Int, minter common.Address, tokenId *big.Int, rewardsRecipients []common.Address, minterArguments []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.MintWithMints(&_Zorasecond.TransactOpts, mintTokenIds, quantities, minter, tokenId, rewardsRecipients, minterArguments)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x9dbb844d.
//
// Solidity: function mintWithRewards(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments, address mintReferral) payable returns()
func (_Zorasecond *ZorasecondTransactor) MintWithRewards(opts *bind.TransactOpts, minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte, mintReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "mintWithRewards", minter, tokenId, quantity, minterArguments, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x9dbb844d.
//
// Solidity: function mintWithRewards(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments, address mintReferral) payable returns()
func (_Zorasecond *ZorasecondSession) MintWithRewards(minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte, mintReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.MintWithRewards(&_Zorasecond.TransactOpts, minter, tokenId, quantity, minterArguments, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x9dbb844d.
//
// Solidity: function mintWithRewards(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments, address mintReferral) payable returns()
func (_Zorasecond *ZorasecondTransactorSession) MintWithRewards(minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte, mintReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.MintWithRewards(&_Zorasecond.TransactOpts, minter, tokenId, quantity, minterArguments, mintReferral)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zorasecond *ZorasecondTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zorasecond *ZorasecondSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Multicall(&_Zorasecond.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zorasecond *ZorasecondTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.Multicall(&_Zorasecond.TransactOpts, data)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondTransactor) RemovePermission(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "removePermission", tokenId, user, permissionBits)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondSession) RemovePermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.RemovePermission(&_Zorasecond.TransactOpts, tokenId, user, permissionBits)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zorasecond *ZorasecondTransactorSession) RemovePermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.RemovePermission(&_Zorasecond.TransactOpts, tokenId, user, permissionBits)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zorasecond *ZorasecondTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zorasecond *ZorasecondSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.SafeBatchTransferFrom(&_Zorasecond.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zorasecond *ZorasecondTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.SafeBatchTransferFrom(&_Zorasecond.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zorasecond *ZorasecondTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zorasecond *ZorasecondSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.SafeTransferFrom(&_Zorasecond.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zorasecond *ZorasecondTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.SafeTransferFrom(&_Zorasecond.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zorasecond *ZorasecondTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zorasecond *ZorasecondSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetApprovalForAll(&_Zorasecond.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zorasecond *ZorasecondTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetApprovalForAll(&_Zorasecond.TransactOpts, operator, approved)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zorasecond *ZorasecondTransactor) SetFundsRecipient(opts *bind.TransactOpts, fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setFundsRecipient", fundsRecipient)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zorasecond *ZorasecondSession) SetFundsRecipient(fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetFundsRecipient(&_Zorasecond.TransactOpts, fundsRecipient)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zorasecond *ZorasecondTransactorSession) SetFundsRecipient(fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetFundsRecipient(&_Zorasecond.TransactOpts, fundsRecipient)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zorasecond *ZorasecondTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zorasecond *ZorasecondSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetOwner(&_Zorasecond.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zorasecond *ZorasecondTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetOwner(&_Zorasecond.TransactOpts, newOwner)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zorasecond *ZorasecondTransactor) SetTokenMetadataRenderer(opts *bind.TransactOpts, tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setTokenMetadataRenderer", tokenId, renderer)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zorasecond *ZorasecondSession) SetTokenMetadataRenderer(tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetTokenMetadataRenderer(&_Zorasecond.TransactOpts, tokenId, renderer)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zorasecond *ZorasecondTransactorSession) SetTokenMetadataRenderer(tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetTokenMetadataRenderer(&_Zorasecond.TransactOpts, tokenId, renderer)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zorasecond *ZorasecondTransactor) SetTransferHook(opts *bind.TransactOpts, transferHook common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setTransferHook", transferHook)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zorasecond *ZorasecondSession) SetTransferHook(transferHook common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetTransferHook(&_Zorasecond.TransactOpts, transferHook)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zorasecond *ZorasecondTransactorSession) SetTransferHook(transferHook common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetTransferHook(&_Zorasecond.TransactOpts, transferHook)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zorasecond *ZorasecondTransactor) SetupNewToken(opts *bind.TransactOpts, newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setupNewToken", newURI, maxSupply)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zorasecond *ZorasecondSession) SetupNewToken(newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetupNewToken(&_Zorasecond.TransactOpts, newURI, maxSupply)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zorasecond *ZorasecondTransactorSession) SetupNewToken(newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetupNewToken(&_Zorasecond.TransactOpts, newURI, maxSupply)
}

// SetupNewTokenWithCreateReferral is a paid mutator transaction binding the contract method 0x674cbae6.
//
// Solidity: function setupNewTokenWithCreateReferral(string newURI, uint256 maxSupply, address createReferral) returns(uint256)
func (_Zorasecond *ZorasecondTransactor) SetupNewTokenWithCreateReferral(opts *bind.TransactOpts, newURI string, maxSupply *big.Int, createReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "setupNewTokenWithCreateReferral", newURI, maxSupply, createReferral)
}

// SetupNewTokenWithCreateReferral is a paid mutator transaction binding the contract method 0x674cbae6.
//
// Solidity: function setupNewTokenWithCreateReferral(string newURI, uint256 maxSupply, address createReferral) returns(uint256)
func (_Zorasecond *ZorasecondSession) SetupNewTokenWithCreateReferral(newURI string, maxSupply *big.Int, createReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetupNewTokenWithCreateReferral(&_Zorasecond.TransactOpts, newURI, maxSupply, createReferral)
}

// SetupNewTokenWithCreateReferral is a paid mutator transaction binding the contract method 0x674cbae6.
//
// Solidity: function setupNewTokenWithCreateReferral(string newURI, uint256 maxSupply, address createReferral) returns(uint256)
func (_Zorasecond *ZorasecondTransactorSession) SetupNewTokenWithCreateReferral(newURI string, maxSupply *big.Int, createReferral common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.SetupNewTokenWithCreateReferral(&_Zorasecond.TransactOpts, newURI, maxSupply, createReferral)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zorasecond *ZorasecondTransactor) UpdateContractMetadata(opts *bind.TransactOpts, _newURI string, _newName string) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "updateContractMetadata", _newURI, _newName)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zorasecond *ZorasecondSession) UpdateContractMetadata(_newURI string, _newName string) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateContractMetadata(&_Zorasecond.TransactOpts, _newURI, _newName)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zorasecond *ZorasecondTransactorSession) UpdateContractMetadata(_newURI string, _newName string) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateContractMetadata(&_Zorasecond.TransactOpts, _newURI, _newName)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x17bd48bb.
//
// Solidity: function updateCreateReferral(uint256 tokenId, address recipient) returns()
func (_Zorasecond *ZorasecondTransactor) UpdateCreateReferral(opts *bind.TransactOpts, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "updateCreateReferral", tokenId, recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x17bd48bb.
//
// Solidity: function updateCreateReferral(uint256 tokenId, address recipient) returns()
func (_Zorasecond *ZorasecondSession) UpdateCreateReferral(tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateCreateReferral(&_Zorasecond.TransactOpts, tokenId, recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x17bd48bb.
//
// Solidity: function updateCreateReferral(uint256 tokenId, address recipient) returns()
func (_Zorasecond *ZorasecondTransactorSession) UpdateCreateReferral(tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateCreateReferral(&_Zorasecond.TransactOpts, tokenId, recipient)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zorasecond *ZorasecondTransactor) UpdateRoyaltiesForToken(opts *bind.TransactOpts, tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "updateRoyaltiesForToken", tokenId, newConfiguration)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zorasecond *ZorasecondSession) UpdateRoyaltiesForToken(tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateRoyaltiesForToken(&_Zorasecond.TransactOpts, tokenId, newConfiguration)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zorasecond *ZorasecondTransactorSession) UpdateRoyaltiesForToken(tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateRoyaltiesForToken(&_Zorasecond.TransactOpts, tokenId, newConfiguration)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zorasecond *ZorasecondTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "updateTokenURI", tokenId, _newURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zorasecond *ZorasecondSession) UpdateTokenURI(tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateTokenURI(&_Zorasecond.TransactOpts, tokenId, _newURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zorasecond *ZorasecondTransactorSession) UpdateTokenURI(tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpdateTokenURI(&_Zorasecond.TransactOpts, tokenId, _newURI)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zorasecond *ZorasecondTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zorasecond *ZorasecondSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpgradeTo(&_Zorasecond.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zorasecond *ZorasecondTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpgradeTo(&_Zorasecond.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorasecond *ZorasecondTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorasecond *ZorasecondSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpgradeToAndCall(&_Zorasecond.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zorasecond *ZorasecondTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zorasecond.Contract.UpgradeToAndCall(&_Zorasecond.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zorasecond *ZorasecondTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorasecond.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zorasecond *ZorasecondSession) Withdraw() (*types.Transaction, error) {
	return _Zorasecond.Contract.Withdraw(&_Zorasecond.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zorasecond *ZorasecondTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Zorasecond.Contract.Withdraw(&_Zorasecond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zorasecond *ZorasecondTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zorasecond.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zorasecond *ZorasecondSession) Receive() (*types.Transaction, error) {
	return _Zorasecond.Contract.Receive(&_Zorasecond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Zorasecond *ZorasecondTransactorSession) Receive() (*types.Transaction, error) {
	return _Zorasecond.Contract.Receive(&_Zorasecond.TransactOpts)
}

// ZorasecondAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Zorasecond contract.
type ZorasecondAdminChangedIterator struct {
	Event *ZorasecondAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZorasecondAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondAdminChanged)
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
		it.Event = new(ZorasecondAdminChanged)
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
func (it *ZorasecondAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondAdminChanged represents a AdminChanged event raised by the Zorasecond contract.
type ZorasecondAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zorasecond *ZorasecondFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ZorasecondAdminChangedIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ZorasecondAdminChangedIterator{contract: _Zorasecond.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zorasecond *ZorasecondFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ZorasecondAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondAdminChanged)
				if err := _Zorasecond.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zorasecond *ZorasecondFilterer) ParseAdminChanged(log types.Log) (*ZorasecondAdminChanged, error) {
	event := new(ZorasecondAdminChanged)
	if err := _Zorasecond.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Zorasecond contract.
type ZorasecondApprovalForAllIterator struct {
	Event *ZorasecondApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ZorasecondApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondApprovalForAll)
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
		it.Event = new(ZorasecondApprovalForAll)
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
func (it *ZorasecondApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondApprovalForAll represents a ApprovalForAll event raised by the Zorasecond contract.
type ZorasecondApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zorasecond *ZorasecondFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ZorasecondApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondApprovalForAllIterator{contract: _Zorasecond.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zorasecond *ZorasecondFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZorasecondApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondApprovalForAll)
				if err := _Zorasecond.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zorasecond *ZorasecondFilterer) ParseApprovalForAll(log types.Log) (*ZorasecondApprovalForAll, error) {
	event := new(ZorasecondApprovalForAll)
	if err := _Zorasecond.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Zorasecond contract.
type ZorasecondBeaconUpgradedIterator struct {
	Event *ZorasecondBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ZorasecondBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondBeaconUpgraded)
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
		it.Event = new(ZorasecondBeaconUpgraded)
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
func (it *ZorasecondBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondBeaconUpgraded represents a BeaconUpgraded event raised by the Zorasecond contract.
type ZorasecondBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zorasecond *ZorasecondFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ZorasecondBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondBeaconUpgradedIterator{contract: _Zorasecond.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zorasecond *ZorasecondFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ZorasecondBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondBeaconUpgraded)
				if err := _Zorasecond.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zorasecond *ZorasecondFilterer) ParseBeaconUpgraded(log types.Log) (*ZorasecondBeaconUpgraded, error) {
	event := new(ZorasecondBeaconUpgraded)
	if err := _Zorasecond.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the Zorasecond contract.
type ZorasecondConfigUpdatedIterator struct {
	Event *ZorasecondConfigUpdated // Event containing the contract specifics and raw log

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
func (it *ZorasecondConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondConfigUpdated)
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
		it.Event = new(ZorasecondConfigUpdated)
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
func (it *ZorasecondConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondConfigUpdated represents a ConfigUpdated event raised by the Zorasecond contract.
type ZorasecondConfigUpdated struct {
	Updater    common.Address
	UpdateType uint8
	NewConfig  IZoraCreator1155TypesV1ContractConfig
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zorasecond *ZorasecondFilterer) FilterConfigUpdated(opts *bind.FilterOpts, updater []common.Address, updateType []uint8) (*ZorasecondConfigUpdatedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "ConfigUpdated", updaterRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondConfigUpdatedIterator{contract: _Zorasecond.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zorasecond *ZorasecondFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *ZorasecondConfigUpdated, updater []common.Address, updateType []uint8) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "ConfigUpdated", updaterRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondConfigUpdated)
				if err := _Zorasecond.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zorasecond *ZorasecondFilterer) ParseConfigUpdated(log types.Log) (*ZorasecondConfigUpdated, error) {
	event := new(ZorasecondConfigUpdated)
	if err := _Zorasecond.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondContractMetadataUpdatedIterator is returned from FilterContractMetadataUpdated and is used to iterate over the raw logs and unpacked data for ContractMetadataUpdated events raised by the Zorasecond contract.
type ZorasecondContractMetadataUpdatedIterator struct {
	Event *ZorasecondContractMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *ZorasecondContractMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondContractMetadataUpdated)
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
		it.Event = new(ZorasecondContractMetadataUpdated)
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
func (it *ZorasecondContractMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondContractMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondContractMetadataUpdated represents a ContractMetadataUpdated event raised by the Zorasecond contract.
type ZorasecondContractMetadataUpdated struct {
	Updater common.Address
	Uri     string
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractMetadataUpdated is a free log retrieval operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zorasecond *ZorasecondFilterer) FilterContractMetadataUpdated(opts *bind.FilterOpts, updater []common.Address) (*ZorasecondContractMetadataUpdatedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "ContractMetadataUpdated", updaterRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondContractMetadataUpdatedIterator{contract: _Zorasecond.contract, event: "ContractMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchContractMetadataUpdated is a free log subscription operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zorasecond *ZorasecondFilterer) WatchContractMetadataUpdated(opts *bind.WatchOpts, sink chan<- *ZorasecondContractMetadataUpdated, updater []common.Address) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "ContractMetadataUpdated", updaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondContractMetadataUpdated)
				if err := _Zorasecond.contract.UnpackLog(event, "ContractMetadataUpdated", log); err != nil {
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

// ParseContractMetadataUpdated is a log parse operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zorasecond *ZorasecondFilterer) ParseContractMetadataUpdated(log types.Log) (*ZorasecondContractMetadataUpdated, error) {
	event := new(ZorasecondContractMetadataUpdated)
	if err := _Zorasecond.contract.UnpackLog(event, "ContractMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondContractRendererUpdatedIterator is returned from FilterContractRendererUpdated and is used to iterate over the raw logs and unpacked data for ContractRendererUpdated events raised by the Zorasecond contract.
type ZorasecondContractRendererUpdatedIterator struct {
	Event *ZorasecondContractRendererUpdated // Event containing the contract specifics and raw log

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
func (it *ZorasecondContractRendererUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondContractRendererUpdated)
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
		it.Event = new(ZorasecondContractRendererUpdated)
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
func (it *ZorasecondContractRendererUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondContractRendererUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondContractRendererUpdated represents a ContractRendererUpdated event raised by the Zorasecond contract.
type ZorasecondContractRendererUpdated struct {
	Renderer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractRendererUpdated is a free log retrieval operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zorasecond *ZorasecondFilterer) FilterContractRendererUpdated(opts *bind.FilterOpts) (*ZorasecondContractRendererUpdatedIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "ContractRendererUpdated")
	if err != nil {
		return nil, err
	}
	return &ZorasecondContractRendererUpdatedIterator{contract: _Zorasecond.contract, event: "ContractRendererUpdated", logs: logs, sub: sub}, nil
}

// WatchContractRendererUpdated is a free log subscription operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zorasecond *ZorasecondFilterer) WatchContractRendererUpdated(opts *bind.WatchOpts, sink chan<- *ZorasecondContractRendererUpdated) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "ContractRendererUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondContractRendererUpdated)
				if err := _Zorasecond.contract.UnpackLog(event, "ContractRendererUpdated", log); err != nil {
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

// ParseContractRendererUpdated is a log parse operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zorasecond *ZorasecondFilterer) ParseContractRendererUpdated(log types.Log) (*ZorasecondContractRendererUpdated, error) {
	event := new(ZorasecondContractRendererUpdated)
	if err := _Zorasecond.contract.UnpackLog(event, "ContractRendererUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Zorasecond contract.
type ZorasecondContractURIUpdatedIterator struct {
	Event *ZorasecondContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *ZorasecondContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondContractURIUpdated)
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
		it.Event = new(ZorasecondContractURIUpdated)
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
func (it *ZorasecondContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondContractURIUpdated represents a ContractURIUpdated event raised by the Zorasecond contract.
type ZorasecondContractURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_Zorasecond *ZorasecondFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*ZorasecondContractURIUpdatedIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &ZorasecondContractURIUpdatedIterator{contract: _Zorasecond.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_Zorasecond *ZorasecondFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *ZorasecondContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondContractURIUpdated)
				if err := _Zorasecond.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xa5d4097edda6d87cb9329af83fb3712ef77eeb13738ffe43cc35a4ce305ad962.
//
// Solidity: event ContractURIUpdated()
func (_Zorasecond *ZorasecondFilterer) ParseContractURIUpdated(log types.Log) (*ZorasecondContractURIUpdated, error) {
	event := new(ZorasecondContractURIUpdated)
	if err := _Zorasecond.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondCreatorAttributionIterator is returned from FilterCreatorAttribution and is used to iterate over the raw logs and unpacked data for CreatorAttribution events raised by the Zorasecond contract.
type ZorasecondCreatorAttributionIterator struct {
	Event *ZorasecondCreatorAttribution // Event containing the contract specifics and raw log

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
func (it *ZorasecondCreatorAttributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondCreatorAttribution)
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
		it.Event = new(ZorasecondCreatorAttribution)
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
func (it *ZorasecondCreatorAttributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondCreatorAttributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondCreatorAttribution represents a CreatorAttribution event raised by the Zorasecond contract.
type ZorasecondCreatorAttribution struct {
	StructHash [32]byte
	DomainName string
	Version    string
	Creator    common.Address
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatorAttribution is a free log retrieval operation binding the contract event 0x06c5a80e592816bd4f60093568e69affa68b5e378a189b2f59a1121703de47de.
//
// Solidity: event CreatorAttribution(bytes32 structHash, string domainName, string version, address creator, bytes signature)
func (_Zorasecond *ZorasecondFilterer) FilterCreatorAttribution(opts *bind.FilterOpts) (*ZorasecondCreatorAttributionIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "CreatorAttribution")
	if err != nil {
		return nil, err
	}
	return &ZorasecondCreatorAttributionIterator{contract: _Zorasecond.contract, event: "CreatorAttribution", logs: logs, sub: sub}, nil
}

// WatchCreatorAttribution is a free log subscription operation binding the contract event 0x06c5a80e592816bd4f60093568e69affa68b5e378a189b2f59a1121703de47de.
//
// Solidity: event CreatorAttribution(bytes32 structHash, string domainName, string version, address creator, bytes signature)
func (_Zorasecond *ZorasecondFilterer) WatchCreatorAttribution(opts *bind.WatchOpts, sink chan<- *ZorasecondCreatorAttribution) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "CreatorAttribution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondCreatorAttribution)
				if err := _Zorasecond.contract.UnpackLog(event, "CreatorAttribution", log); err != nil {
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

// ParseCreatorAttribution is a log parse operation binding the contract event 0x06c5a80e592816bd4f60093568e69affa68b5e378a189b2f59a1121703de47de.
//
// Solidity: event CreatorAttribution(bytes32 structHash, string domainName, string version, address creator, bytes signature)
func (_Zorasecond *ZorasecondFilterer) ParseCreatorAttribution(log types.Log) (*ZorasecondCreatorAttribution, error) {
	event := new(ZorasecondCreatorAttribution)
	if err := _Zorasecond.contract.UnpackLog(event, "CreatorAttribution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Zorasecond contract.
type ZorasecondInitializedIterator struct {
	Event *ZorasecondInitialized // Event containing the contract specifics and raw log

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
func (it *ZorasecondInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondInitialized)
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
		it.Event = new(ZorasecondInitialized)
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
func (it *ZorasecondInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondInitialized represents a Initialized event raised by the Zorasecond contract.
type ZorasecondInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Zorasecond *ZorasecondFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZorasecondInitializedIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZorasecondInitializedIterator{contract: _Zorasecond.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Zorasecond *ZorasecondFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZorasecondInitialized) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondInitialized)
				if err := _Zorasecond.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Zorasecond *ZorasecondFilterer) ParseInitialized(log types.Log) (*ZorasecondInitialized, error) {
	event := new(ZorasecondInitialized)
	if err := _Zorasecond.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zorasecond contract.
type ZorasecondOwnershipTransferredIterator struct {
	Event *ZorasecondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZorasecondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondOwnershipTransferred)
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
		it.Event = new(ZorasecondOwnershipTransferred)
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
func (it *ZorasecondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondOwnershipTransferred represents a OwnershipTransferred event raised by the Zorasecond contract.
type ZorasecondOwnershipTransferred struct {
	LastOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zorasecond *ZorasecondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*ZorasecondOwnershipTransferredIterator, error) {

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &ZorasecondOwnershipTransferredIterator{contract: _Zorasecond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zorasecond *ZorasecondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZorasecondOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondOwnershipTransferred)
				if err := _Zorasecond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zorasecond *ZorasecondFilterer) ParseOwnershipTransferred(log types.Log) (*ZorasecondOwnershipTransferred, error) {
	event := new(ZorasecondOwnershipTransferred)
	if err := _Zorasecond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondPurchasedIterator is returned from FilterPurchased and is used to iterate over the raw logs and unpacked data for Purchased events raised by the Zorasecond contract.
type ZorasecondPurchasedIterator struct {
	Event *ZorasecondPurchased // Event containing the contract specifics and raw log

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
func (it *ZorasecondPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondPurchased)
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
		it.Event = new(ZorasecondPurchased)
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
func (it *ZorasecondPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondPurchased represents a Purchased event raised by the Zorasecond contract.
type ZorasecondPurchased struct {
	Sender   common.Address
	Minter   common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchased is a free log retrieval operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zorasecond *ZorasecondFilterer) FilterPurchased(opts *bind.FilterOpts, sender []common.Address, minter []common.Address, tokenId []*big.Int) (*ZorasecondPurchasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "Purchased", senderRule, minterRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondPurchasedIterator{contract: _Zorasecond.contract, event: "Purchased", logs: logs, sub: sub}, nil
}

// WatchPurchased is a free log subscription operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zorasecond *ZorasecondFilterer) WatchPurchased(opts *bind.WatchOpts, sink chan<- *ZorasecondPurchased, sender []common.Address, minter []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "Purchased", senderRule, minterRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondPurchased)
				if err := _Zorasecond.contract.UnpackLog(event, "Purchased", log); err != nil {
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

// ParsePurchased is a log parse operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zorasecond *ZorasecondFilterer) ParsePurchased(log types.Log) (*ZorasecondPurchased, error) {
	event := new(ZorasecondPurchased)
	if err := _Zorasecond.contract.UnpackLog(event, "Purchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondRendererUpdatedIterator is returned from FilterRendererUpdated and is used to iterate over the raw logs and unpacked data for RendererUpdated events raised by the Zorasecond contract.
type ZorasecondRendererUpdatedIterator struct {
	Event *ZorasecondRendererUpdated // Event containing the contract specifics and raw log

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
func (it *ZorasecondRendererUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondRendererUpdated)
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
		it.Event = new(ZorasecondRendererUpdated)
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
func (it *ZorasecondRendererUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondRendererUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondRendererUpdated represents a RendererUpdated event raised by the Zorasecond contract.
type ZorasecondRendererUpdated struct {
	TokenId  *big.Int
	Renderer common.Address
	User     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRendererUpdated is a free log retrieval operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zorasecond *ZorasecondFilterer) FilterRendererUpdated(opts *bind.FilterOpts, tokenId []*big.Int, renderer []common.Address, user []common.Address) (*ZorasecondRendererUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rendererRule []interface{}
	for _, rendererItem := range renderer {
		rendererRule = append(rendererRule, rendererItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "RendererUpdated", tokenIdRule, rendererRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondRendererUpdatedIterator{contract: _Zorasecond.contract, event: "RendererUpdated", logs: logs, sub: sub}, nil
}

// WatchRendererUpdated is a free log subscription operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zorasecond *ZorasecondFilterer) WatchRendererUpdated(opts *bind.WatchOpts, sink chan<- *ZorasecondRendererUpdated, tokenId []*big.Int, renderer []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rendererRule []interface{}
	for _, rendererItem := range renderer {
		rendererRule = append(rendererRule, rendererItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "RendererUpdated", tokenIdRule, rendererRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondRendererUpdated)
				if err := _Zorasecond.contract.UnpackLog(event, "RendererUpdated", log); err != nil {
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

// ParseRendererUpdated is a log parse operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zorasecond *ZorasecondFilterer) ParseRendererUpdated(log types.Log) (*ZorasecondRendererUpdated, error) {
	event := new(ZorasecondRendererUpdated)
	if err := _Zorasecond.contract.UnpackLog(event, "RendererUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondSetupNewTokenIterator is returned from FilterSetupNewToken and is used to iterate over the raw logs and unpacked data for SetupNewToken events raised by the Zorasecond contract.
type ZorasecondSetupNewTokenIterator struct {
	Event *ZorasecondSetupNewToken // Event containing the contract specifics and raw log

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
func (it *ZorasecondSetupNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondSetupNewToken)
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
		it.Event = new(ZorasecondSetupNewToken)
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
func (it *ZorasecondSetupNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondSetupNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondSetupNewToken represents a SetupNewToken event raised by the Zorasecond contract.
type ZorasecondSetupNewToken struct {
	TokenId   *big.Int
	Sender    common.Address
	NewURI    string
	MaxSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetupNewToken is a free log retrieval operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zorasecond *ZorasecondFilterer) FilterSetupNewToken(opts *bind.FilterOpts, tokenId []*big.Int, sender []common.Address) (*ZorasecondSetupNewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "SetupNewToken", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondSetupNewTokenIterator{contract: _Zorasecond.contract, event: "SetupNewToken", logs: logs, sub: sub}, nil
}

// WatchSetupNewToken is a free log subscription operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zorasecond *ZorasecondFilterer) WatchSetupNewToken(opts *bind.WatchOpts, sink chan<- *ZorasecondSetupNewToken, tokenId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "SetupNewToken", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondSetupNewToken)
				if err := _Zorasecond.contract.UnpackLog(event, "SetupNewToken", log); err != nil {
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

// ParseSetupNewToken is a log parse operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zorasecond *ZorasecondFilterer) ParseSetupNewToken(log types.Log) (*ZorasecondSetupNewToken, error) {
	event := new(ZorasecondSetupNewToken)
	if err := _Zorasecond.contract.UnpackLog(event, "SetupNewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Zorasecond contract.
type ZorasecondTransferBatchIterator struct {
	Event *ZorasecondTransferBatch // Event containing the contract specifics and raw log

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
func (it *ZorasecondTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondTransferBatch)
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
		it.Event = new(ZorasecondTransferBatch)
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
func (it *ZorasecondTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondTransferBatch represents a TransferBatch event raised by the Zorasecond contract.
type ZorasecondTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zorasecond *ZorasecondFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ZorasecondTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondTransferBatchIterator{contract: _Zorasecond.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zorasecond *ZorasecondFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ZorasecondTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondTransferBatch)
				if err := _Zorasecond.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zorasecond *ZorasecondFilterer) ParseTransferBatch(log types.Log) (*ZorasecondTransferBatch, error) {
	event := new(ZorasecondTransferBatch)
	if err := _Zorasecond.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Zorasecond contract.
type ZorasecondTransferSingleIterator struct {
	Event *ZorasecondTransferSingle // Event containing the contract specifics and raw log

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
func (it *ZorasecondTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondTransferSingle)
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
		it.Event = new(ZorasecondTransferSingle)
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
func (it *ZorasecondTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondTransferSingle represents a TransferSingle event raised by the Zorasecond contract.
type ZorasecondTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zorasecond *ZorasecondFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ZorasecondTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondTransferSingleIterator{contract: _Zorasecond.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zorasecond *ZorasecondFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ZorasecondTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondTransferSingle)
				if err := _Zorasecond.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zorasecond *ZorasecondFilterer) ParseTransferSingle(log types.Log) (*ZorasecondTransferSingle, error) {
	event := new(ZorasecondTransferSingle)
	if err := _Zorasecond.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Zorasecond contract.
type ZorasecondURIIterator struct {
	Event *ZorasecondURI // Event containing the contract specifics and raw log

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
func (it *ZorasecondURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondURI)
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
		it.Event = new(ZorasecondURI)
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
func (it *ZorasecondURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondURI represents a URI event raised by the Zorasecond contract.
type ZorasecondURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zorasecond *ZorasecondFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ZorasecondURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondURIIterator{contract: _Zorasecond.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zorasecond *ZorasecondFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ZorasecondURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondURI)
				if err := _Zorasecond.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zorasecond *ZorasecondFilterer) ParseURI(log types.Log) (*ZorasecondURI, error) {
	event := new(ZorasecondURI)
	if err := _Zorasecond.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondUpdatedPermissionsIterator is returned from FilterUpdatedPermissions and is used to iterate over the raw logs and unpacked data for UpdatedPermissions events raised by the Zorasecond contract.
type ZorasecondUpdatedPermissionsIterator struct {
	Event *ZorasecondUpdatedPermissions // Event containing the contract specifics and raw log

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
func (it *ZorasecondUpdatedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondUpdatedPermissions)
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
		it.Event = new(ZorasecondUpdatedPermissions)
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
func (it *ZorasecondUpdatedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondUpdatedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondUpdatedPermissions represents a UpdatedPermissions event raised by the Zorasecond contract.
type ZorasecondUpdatedPermissions struct {
	TokenId     *big.Int
	User        common.Address
	Permissions *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPermissions is a free log retrieval operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zorasecond *ZorasecondFilterer) FilterUpdatedPermissions(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address, permissions []*big.Int) (*ZorasecondUpdatedPermissionsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "UpdatedPermissions", tokenIdRule, userRule, permissionsRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondUpdatedPermissionsIterator{contract: _Zorasecond.contract, event: "UpdatedPermissions", logs: logs, sub: sub}, nil
}

// WatchUpdatedPermissions is a free log subscription operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zorasecond *ZorasecondFilterer) WatchUpdatedPermissions(opts *bind.WatchOpts, sink chan<- *ZorasecondUpdatedPermissions, tokenId []*big.Int, user []common.Address, permissions []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "UpdatedPermissions", tokenIdRule, userRule, permissionsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondUpdatedPermissions)
				if err := _Zorasecond.contract.UnpackLog(event, "UpdatedPermissions", log); err != nil {
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

// ParseUpdatedPermissions is a log parse operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zorasecond *ZorasecondFilterer) ParseUpdatedPermissions(log types.Log) (*ZorasecondUpdatedPermissions, error) {
	event := new(ZorasecondUpdatedPermissions)
	if err := _Zorasecond.contract.UnpackLog(event, "UpdatedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondUpdatedRoyaltiesIterator is returned from FilterUpdatedRoyalties and is used to iterate over the raw logs and unpacked data for UpdatedRoyalties events raised by the Zorasecond contract.
type ZorasecondUpdatedRoyaltiesIterator struct {
	Event *ZorasecondUpdatedRoyalties // Event containing the contract specifics and raw log

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
func (it *ZorasecondUpdatedRoyaltiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondUpdatedRoyalties)
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
		it.Event = new(ZorasecondUpdatedRoyalties)
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
func (it *ZorasecondUpdatedRoyaltiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondUpdatedRoyaltiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondUpdatedRoyalties represents a UpdatedRoyalties event raised by the Zorasecond contract.
type ZorasecondUpdatedRoyalties struct {
	TokenId       *big.Int
	User          common.Address
	Configuration ICreatorRoyaltiesControlRoyaltyConfiguration
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRoyalties is a free log retrieval operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zorasecond *ZorasecondFilterer) FilterUpdatedRoyalties(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*ZorasecondUpdatedRoyaltiesIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "UpdatedRoyalties", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondUpdatedRoyaltiesIterator{contract: _Zorasecond.contract, event: "UpdatedRoyalties", logs: logs, sub: sub}, nil
}

// WatchUpdatedRoyalties is a free log subscription operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zorasecond *ZorasecondFilterer) WatchUpdatedRoyalties(opts *bind.WatchOpts, sink chan<- *ZorasecondUpdatedRoyalties, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "UpdatedRoyalties", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondUpdatedRoyalties)
				if err := _Zorasecond.contract.UnpackLog(event, "UpdatedRoyalties", log); err != nil {
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

// ParseUpdatedRoyalties is a log parse operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zorasecond *ZorasecondFilterer) ParseUpdatedRoyalties(log types.Log) (*ZorasecondUpdatedRoyalties, error) {
	event := new(ZorasecondUpdatedRoyalties)
	if err := _Zorasecond.contract.UnpackLog(event, "UpdatedRoyalties", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondUpdatedTokenIterator is returned from FilterUpdatedToken and is used to iterate over the raw logs and unpacked data for UpdatedToken events raised by the Zorasecond contract.
type ZorasecondUpdatedTokenIterator struct {
	Event *ZorasecondUpdatedToken // Event containing the contract specifics and raw log

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
func (it *ZorasecondUpdatedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondUpdatedToken)
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
		it.Event = new(ZorasecondUpdatedToken)
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
func (it *ZorasecondUpdatedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondUpdatedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondUpdatedToken represents a UpdatedToken event raised by the Zorasecond contract.
type ZorasecondUpdatedToken struct {
	From      common.Address
	TokenId   *big.Int
	TokenData IZoraCreator1155TypesV1TokenData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedToken is a free log retrieval operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zorasecond *ZorasecondFilterer) FilterUpdatedToken(opts *bind.FilterOpts, from []common.Address, tokenId []*big.Int) (*ZorasecondUpdatedTokenIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "UpdatedToken", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondUpdatedTokenIterator{contract: _Zorasecond.contract, event: "UpdatedToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedToken is a free log subscription operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zorasecond *ZorasecondFilterer) WatchUpdatedToken(opts *bind.WatchOpts, sink chan<- *ZorasecondUpdatedToken, from []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "UpdatedToken", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondUpdatedToken)
				if err := _Zorasecond.contract.UnpackLog(event, "UpdatedToken", log); err != nil {
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

// ParseUpdatedToken is a log parse operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zorasecond *ZorasecondFilterer) ParseUpdatedToken(log types.Log) (*ZorasecondUpdatedToken, error) {
	event := new(ZorasecondUpdatedToken)
	if err := _Zorasecond.contract.UnpackLog(event, "UpdatedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZorasecondUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Zorasecond contract.
type ZorasecondUpgradedIterator struct {
	Event *ZorasecondUpgraded // Event containing the contract specifics and raw log

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
func (it *ZorasecondUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZorasecondUpgraded)
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
		it.Event = new(ZorasecondUpgraded)
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
func (it *ZorasecondUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZorasecondUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZorasecondUpgraded represents a Upgraded event raised by the Zorasecond contract.
type ZorasecondUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zorasecond *ZorasecondFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZorasecondUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zorasecond.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZorasecondUpgradedIterator{contract: _Zorasecond.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zorasecond *ZorasecondFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZorasecondUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zorasecond.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZorasecondUpgraded)
				if err := _Zorasecond.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Zorasecond *ZorasecondFilterer) ParseUpgraded(log types.Log) (*ZorasecondUpgraded, error) {
	event := new(ZorasecondUpgraded)
	if err := _Zorasecond.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
