package basename

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getInstance(client *ethclient.Client, contractAddress string) (instance *Basename) {
	instance, err := NewBasename(
		common.HexToAddress(contractAddress),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

//func BalanceOf(account *account.Account, contractAddress string) *big.Int {
//	instance := getInstance(account.Client, contractAddress)
//	balance, err := instance.BalanceOf(&bind.CallOpts{}, account.Address())
//	if err != nil {
//		panic(fmt.Sprintf("Get balance error - %s", err))
//	}
//	return balance
//}

//func Mint(account *account.Account, contractAddress string, value int, payload claimPayload, badgeId string) {
//	instance := getInstance(account.Client, contractAddress)
//	balance := BalanceOf(account, contractAddress)
//	if balance.Int64() == int64(0) {
//		txData := account.TxData(int64(value))
//		tx, err := instance.Regis()
//		if err != nil {
//			panic(fmt.Sprintf("Mint error - %s", err))
//		}
//		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
//		err = account.ClaimBadge(badgeId)
//		if err != nil {
//			log.Printf("[%s] [%s] Claim Badge error - %s", strconv.Itoa(account.Id), account.Address(), err)
//		}
//	} else {
//		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
//	}
//}
