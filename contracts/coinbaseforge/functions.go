package coinbaseforge

import (
	"base/account"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
)

func getInstance(client *ethclient.Client, contractAddress string) (instance *Coinbaseforge) {
	instance, err := NewCoinbaseforge(
		common.HexToAddress(contractAddress),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

func BalanceOf(account *account.Account, contractAddress string) *big.Int {
	instance := getInstance(account.Client, contractAddress)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account.Address())
	if err != nil {
		panic(fmt.Sprintf("Get balance error - %s", err))
	}
	return balance
}

func MintWithComment(account *account.Account, contractAddress string, value int, quantity int, comment string) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress)
	if balance != big.NewInt(0) {
		txData := account.TxData(int64(value))
		tx, err := instance.MintWithComment(txData,
			account.Address(),
			big.NewInt(int64(quantity)),
			comment,
		)

		if err != nil {
			panic(fmt.Sprintf("Tx error: %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), strconv.Itoa(int(balance.Int64())))
	}
}

func MintNounsSummer(account *account.Account) {
	MintWithComment(account, "0x01F0D0D40D4BB499e7CB35940E908b74D08BA412", 100000000000000, 1, "")
}

func MintWalkingOnSunshine(account *account.Account) {
	MintWithComment(account, "0x4E4c12451A6e2473Fc4C63f84E175C3D31555F47", 100000000000000, 1, "")
}

func MintConsolationOfChromaConvergence(account *account.Account) {
	MintWithComment(account, "0x7d22F7EC034bB1060E033f132b0b23aA45b2B9B4", 100000000000000, 1, "")
}

func MintStrut(account *account.Account) {
	MintWithComment(account, "0x9FF8Fd82c0ce09caE76e777f47d536579AF2Fe7C", 100000000000000, 1, "")
}

func MintBaseNft(account *account.Account) {
	MintWithComment(account, "0x3b4B32a5c9A01763A0945A8a4a4269052DC3DE2F", 100000000000000, 1, "")
}

func MintToshiOnchainSummer(account *account.Account) {
	MintWithComment(account, "0x44d461Da8A451f05b6b75EdD5C4a2d2f3C14aaB4", 100000000000000, 1, "")
}

func Mintᐊᔮᑎᓯᐏᐣayâtisiwin(account *account.Account) {
	MintWithComment(account, "0x7B791EdF061Df65bAC7a9d47668F61F1a9A998C0", 100000000000000, 1, "")
}

func MintNounisvary(account *account.Account) {
	MintWithComment(account, "0xC8f93Ce7A12960466a2e13E70dE5CA41B652e4E6", 100000000000000, 1, "")
}

func MintSummerSerenity(account *account.Account) {
	MintWithComment(account, "0x777777722D078c97c6ad07d9f36801e653E356Ae", 100000000000000, 1, "")
}

func MintThinkBig(account *account.Account) {
	MintWithComment(account, "0x752d593b3B8aD1c5d827F5B9AA9b653eE7134ea0", 100000000000000, 1, "")
}
