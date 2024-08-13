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

func MintWithComment(account *account.Account, contractAddress string, value int, quantity int) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress)
	if balance != big.NewInt(0) {
		txData := account.TxData(int64(value))
		tx, err := instance.MintWithComment(txData,
			account.Address(),
			big.NewInt(int64(quantity)),
			"")

		if err != nil {
			panic(fmt.Sprintf("Tx error: %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), strconv.Itoa(int(balance.Int64())))
	}
}

func MintNounsSummer(account *account.Account) {
	MintWithComment(account, "0x01F0D0D40D4BB499e7CB35940E908b74D08BA412", 100000000000000, 1)
}

func MintWalkingOnSunshine(account *account.Account) {
	MintWithComment(account, "0x4E4c12451A6e2473Fc4C63f84E175C3D31555F47", 100000000000000, 1)
}

func MintConsolationOfChromaConvergence(account *account.Account) {
	MintWithComment(account, "0x7d22F7EC034bB1060E033f132b0b23aA45b2B9B4", 100000000000000, 1)
}
