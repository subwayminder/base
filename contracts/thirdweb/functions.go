package thirdweb

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

type claimPayload struct {
	quantity       *big.Int
	currency       common.Address
	pricePerToken  *big.Int
	allowListProof IDropAllowlistProof
	data           []byte
}

func getInstance(client *ethclient.Client, contractAddress string) (instance *Thirdweb) {
	instance, err := NewThirdweb(
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

func Mint(account *account.Account, contractAddress string, value int, payload claimPayload) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress)
	if balance.Int64() == int64(0) {
		txData := account.TxData(int64(value))
		tx, err := instance.Claim(
			txData,
			account.Address(),
			payload.quantity,
			payload.currency,
			payload.pricePerToken,
			payload.allowListProof,
			payload.data,
		)
		if err != nil {
			panic(fmt.Sprintf("Mint error - %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
	}
}

func MintStix(account *account.Account) {
	proof := [][32]byte{[32]byte(common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000"))}
	payload := claimPayload{
		quantity:      big.NewInt(1),
		currency:      common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		pricePerToken: big.NewInt(0),
		allowListProof: IDropAllowlistProof{
			Proof:                  proof,
			QuantityLimitPerWallet: big.NewInt(1),
			PricePerToken:          big.NewInt(0),
			Currency:               common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		},
		data: common.FromHex("0x"),
	}
	Mint(account, "0xa7891c87933BB99Db006b60D8Cb7cf68141B492f", 0, payload)
}
