package zorasecond

import (
	"base/account"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
	"time"
)

func getInstance(client *ethclient.Client, contractAddress string) (instance *Zorasecond) {
	instance, err := NewZorasecond(
		common.HexToAddress(contractAddress),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

func BalanceOf(account *account.Account, contractAddress string, tokenId int) *big.Int {
	instance := getInstance(account.Client, contractAddress)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account.Address(), big.NewInt(int64(tokenId)))
	if err != nil {
		panic(fmt.Sprintf("Get balance error - %s", err))
	}
	log.Printf(balance.String())
	return balance
}

func Mint(
	account *account.Account,
	contractAddress string,
	value int,
	tokenId int,
	minter string,
	rewardAddress string,
	badgeId string,
) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress, 1)
	if balance.Int64() == 0 {
		txData := account.TxData(int64(value))
		var rewardRecipients []common.Address
		rewardRecipients = append(rewardRecipients, common.HexToAddress(rewardAddress))
		tx, err := instance.Mint(txData,
			common.HexToAddress(minter),
			big.NewInt(int64(tokenId)),
			big.NewInt(1),
			rewardRecipients,
			common.FromHex(fmt.Sprintf("0x000000000000000000000000%s", account.Address().String()[2:])),
		)
		if err != nil {
			panic("Mint error: " + err.Error())
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
		err = account.ClaimBadge(badgeId)
		if err != nil {
			log.Printf("[%s] [%s] Claim Badge error - %s", strconv.Itoa(account.Id), account.Address(), err)
		} else {
			log.Printf("[%s] [%s] Claim Badge success", strconv.Itoa(account.Id), account.Address())
		}
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
	}
}

func MintWithRewards(
	account *account.Account,
	contractAddress string,
	value int, tokenId int,
	minter string,
	mintReferral string,
	badgeId string,
) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress, 1)
	if balance.Int64() == 0 {
		txData := account.TxData(int64(value))
		tx, err := instance.MintWithRewards(txData,
			common.HexToAddress(minter),
			big.NewInt(int64(tokenId)),
			big.NewInt(1),
			common.FromHex(fmt.Sprintf("0x000000000000000000000000%s", account.Address().String()[2:])),
			common.HexToAddress(mintReferral),
		)
		if err != nil {
			panic("Mint error: " + err.Error())
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
		time.Sleep(3 * time.Second)
		err = account.ClaimBadge(badgeId)
		if err != nil {
			log.Printf("[%s] [%s] Claim Badge error - %s", strconv.Itoa(account.Id), account.Address(), err)
		}
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
	}
}

func MintHigher(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x9aa1cb37e9fE475D21bc2785AE402Ca4e9fC4fCD",
		777000000000000,
		49,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintBouquet(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x54535B6e8eFfDB89D54Ca58BFc774d04Bf5Af542",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintHodlhq(account *account.Account, badgeId string) {
	MintWithRewards(
		account,
		"0x1b0475c656A47403Ba9f23C1ED7f695b1914d81C",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintFullmetaldroid(account *account.Account, badgeId string) {
	Mint(
		account,
		"0xa902601Ece8B81D906b7dECeB67F5bADcbdfF7dF",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintEssence(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x364065FE510A2d097fbE758Fa78Ca4B7a55F9B84",
		777000000000000,
		3,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintBaseSummerGenesis(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x28510c1D60595349C2fa4d4565eec3A53e709c7a",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintOnchainomics(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x36A16070Fe69Fe6A8257d9ce77b3aD04B36670A9",
		777000000000000,
		2,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintStayBased(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x08f83bC62115c7808009cBA5766517C71d5F774c",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintPalomarGroupArchive(account *account.Account, badgeId string) {
	Mint(
		account,
		"0xbCbEd193Fbc6bBA740607E64CC26042d052EBE85",
		777000000000000,
		16,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintBuildingOnchainPodcast(account *account.Account, badgeId string) {
	Mint(
		account,
		"0xE90577c18a1abd0F748dFB542Cd55a81EE0dDA1E",
		777000000000000,
		2,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintOnchainomicsEssay(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x36A16070Fe69Fe6A8257d9ce77b3aD04B36670A9",
		777000000000000,
		3,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintBasedOnHappines(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x364065FE510A2d097fbE758Fa78Ca4B7a55F9B84",
		777000000000000,
		4,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintSummerBasecamp(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x352efBB53972d962A01A9d708C43F5B4510151b6",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintPersonaOnBase(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x607e2e892dD7ae0c9B8C92F518196D254FE2B071",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintGirlOnchain(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x61AF473071F4c3Ab7d8E02E5083068bCe586C56F",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		badgeId,
	)
}

func MintZeropod(account *account.Account, badgeId string) {
	Mint(
		account,
		"0xF85fdc5728439105CEA9c52baabe3e7Cbac36934",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintSpiral360(account *account.Account, badgeId string) {
	Mint(
		account,
		"0xE201626aea546dB793018131ecf131ECdAba9c07",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}

func MintMemloop(account *account.Account, badgeId string) {
	Mint(
		account,
		"0x3397E9911E4C3c87dC17890ddd306f01c0F63E73",
		777000000000000,
		1,
		"0x04E2516A2c207E84a1839755675dfd8eF6302F0a",
		"0x0000000000000000000000000000000000000000",
		badgeId,
	)
}
