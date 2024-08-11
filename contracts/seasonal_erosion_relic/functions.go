package seasonal_erosion_relic

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

func getInstance(client *ethclient.Client) (instance *SeasonalErosionRelic) {
	instance, err := NewSeasonalErosionRelic(
		common.HexToAddress("0x2aa80a13395425EF3897c9684a0249a5226eA779"),
		client)
	if err != nil {
		panic(fmt.Sprintf("Get contact instance error - %s", err))
	}
	return instance
}

func BalanceOf(account *account.Account) *big.Int {
	instance := getInstance(account.Client)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, account.Address(), big.NewInt(4))
	if err != nil {
		panic(fmt.Sprintf("Get balance error - %s", err))
	}
	return balance
}

func MintSeasonalErosion(account *account.Account) {
	instance := getInstance(account.Client)
	balance := BalanceOf(account)
	if balance.Int64() == int64(0) {
		txData := account.TxData(0)
		tx, err := instance.Mint(txData, big.NewInt(4))
		if err != nil {
			panic(err)
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
	}
}
