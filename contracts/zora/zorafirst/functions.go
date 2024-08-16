package zorafirst

import (
	"base/account"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BalanceData struct {
	Balance int `json:"balance"`
}

func getInstance(client *ethclient.Client, contractAddress string) (instance *Zorafirst) {
	instance, err := NewZorafirst(
		common.HexToAddress(contractAddress),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

func BalanceOf(account *account.Account, collectionAddress string, tokenId int) int {
	urlBalance := fmt.Sprintf("https://zora.co/api/personalize/collection/base:%s/balance?user=%s&tokenId=%s",
		collectionAddress,
		account.Address(),
		strconv.Itoa(tokenId),
	)
	urlI := url.URL{}
	proxyUrl, err := urlI.Parse(account.Proxy)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse proxy url: %s", err))
	}
	transport := http.Transport{}
	transport.Proxy = http.ProxyURL(proxyUrl)
	proxyClient := &http.Client{Transport: &transport, Timeout: 5 * time.Second}

	balanceResponse, err := proxyClient.Get(urlBalance)
	if err != nil && balanceResponse.StatusCode != http.StatusOK {
		log.Fatalf("Error init http request: %s\n", err)
	}
	var balanceResponseBody BalanceData
	byteBody, _ := io.ReadAll(balanceResponse.Body)
	if err := json.Unmarshal(byteBody, &balanceResponseBody); err != nil {
		panic(fmt.Sprintf("Can not unmarshal JSON - %s", err))
	}
	return balanceResponseBody.Balance
}

func Mint(account *account.Account, contractAddress string, collectionAddress string, value int, tokenId int, minterReferal string, comment string) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, collectionAddress, tokenId)
	if balance == 0 {
		txData := account.TxData(int64(value))
		tx, err := instance.Mint(txData,
			account.Address(),
			big.NewInt(1),
			common.HexToAddress(collectionAddress),
			big.NewInt(int64(tokenId)),
			common.HexToAddress(minterReferal),
			comment)

		if err != nil {
			panic(fmt.Sprintf("Tx error: %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), strconv.Itoa(balance))
	}
}

func MintOneYearOnBase(account *account.Account) {
	Mint(account,
		"0x777777722D078c97c6ad07d9f36801e653E356Ae",
		"0xb4703a3a73Aec16E764CBd210b0Fde9EFdAB8941",
		111000000000000,
		1,
		"0x0000000000000000000000000000000000000000",
		"",
	)
}

func MintEicSummer(account *account.Account) {
	Mint(
		account,
		"0x777777722D078c97c6ad07d9f36801e653E356Ae",
		"0x9D2FC5fFE5939Efd1d573f975BC5EEFd364779ae",
		111000000000000,
		5,
		"0x55C88bB05602Da94fCe8FEadc1cbebF5B72c2453",
		"",
	)
}
