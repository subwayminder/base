package olympic

import (
	"base/account"
	"base/helpers"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type NonceResponseBody struct {
	Nonce string `json:"nonce"`
}

func getInstance(client *ethclient.Client) (instance *Olympic) {
	instance, err := NewOlympic(
		common.HexToAddress("0xEEadefc9Df7ed4995cb93f5b5D9b923a7Dff8599"),
		client)
	if err != nil {
		panic(fmt.Sprintf("Get contact instance error - %s", err))
	}
	return instance
}

func getSign(account *account.Account) (signature []byte) {
	nonceRequest, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://auth.magiceden.io/api/v0/sdk/%s", "c1314b5b-ece8-4b4f-a879-3894dda364e4"),
		nil)
	if err != nil {
		panic(fmt.Sprintf("Error init http request: %s\n", err))
	}
	nonceRequest.Header.Set("Proxy", account.Proxy)
	nonceResponse, err := http.DefaultClient.Do(nonceRequest)
	if err != nil && nonceResponse.StatusCode != http.StatusOK {
		log.Fatalf("Error init http request: %s\n", err)
	}
	var nonceResponseBody NonceResponseBody
	byteBody, _ := io.ReadAll(nonceResponse.Body)
	if err := json.Unmarshal(byteBody, &nonceResponseBody); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	messageForSign := "magiceden.io wants you to sign in with your Ethereum account:" +
		"\n0xd3B7Bf6305Ae522890BA52c3A1eadF6A3FBb4B28" +
		"\n" +
		"\nWelcome to Magic Eden. Signing is the only way we can truly know that you are the owner of the wallet you are connecting. " +
		"Signing is a safe, gas-less transaction that does not in any way give Magic Eden permission to perform any transactions with your wallet.\n\nURI: https://magiceden.io/launchpad/base/olympic2024" +
		"\nVersion: 1" +
		"\nChain ID: " + account.ChainId.String() +
		"\nNonce: " + nonceResponseBody.Nonce +
		"\nIssued At: " + time.Now().In(helpers.GetUtcTimezone()).Format("2006-01-02T15:04:05.000000Z") +
		"\nRequest ID: c1314b5b-ece8-4b4f-a879-3894dda364e4"

	signature, err = account.PersonalSign(messageForSign)
	if err != nil {
		panic(fmt.Sprintf("Sign message error - %s", err))
	}
	return signature
}

func BalanceOf(account *account.Account) *big.Int {
	instance := getInstance(account.Client)

	balance, err := instance.BalanceOf(&bind.CallOpts{}, account.Address())
	if err != nil {
		panic(fmt.Sprintf("Get balance error - %s", err))
	}
	return balance
}

func MintOlympic(account *account.Account, badgeId string) {
	instance := getInstance(account.Client)
	balance := BalanceOf(account)
	if balance.Int64() == int64(0) {
		sign := getSign(account)
		txData := account.TxData(0)
		var byteSlice [][32]byte
		tx, err := instance.Mint(txData, 1, byteSlice, uint64(time.Now().In(helpers.GetUtcTimezone()).Unix()), sign)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
		time.Sleep(3 * time.Second)
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
