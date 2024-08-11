package teamliquid

import (
	"base/account"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
	"strings"
)

type CallData struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Data  string `json:"data"`
	Value string `json:"value"`
}

type DataResponse struct {
	CallData CallData `json:"callData"`
}

func getInstance(client *ethclient.Client) (instance *Teamliquid) {
	instance, err := NewTeamliquid(
		common.HexToAddress("0x1aeD60A97192157fDA7fb26267A439d523d09c5e"),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

func getData(account *account.Account) (data string) {
	body := strings.NewReader(fmt.Sprintf(
		`{"bypassSimulation":true,"mintAddress":"%s","network":"networks/base-mainnet","quantity":"%s","takerAddress":"%s"}`,
		"0x1b9ac8580d2e81d7322f163362831448e7fcad1b",
		"1",
		account.Address()))
	dataRequest, err := http.NewRequest(
		"POST",
		"https://api.wallet.coinbase.com/rpc/v3/creators/mintToken",
		body)
	if err != nil {
		panic(fmt.Sprintf("Error init http request: %s", err))
	}
	//dataRequest.Header.Set("Proxy", account.Proxy)
	dataRequest.Header.Add("Content-Type", "application/json")
	dataRequest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	dataRequest.Header.Add("Cookie", "cb_dm=ffece33e-6ae3-4e3b-bd6f-4f18b9c4dd95; __cf_bm=e47z1mqCf.0HOfB8N6KVw.6_I9t4LvS6idKukTV4Q_w-1723148635-1.0.1.1-76GbGlvEcHzPmBCr_gbOZ2BfoH63RpOyRsUsJeiBQGCZRGy9FJzxeVIC1Ivlr9h3jP158fwDrUzG6MPcjvYFCw")
	dataResponse, err := http.DefaultClient.Do(dataRequest)

	if err != nil {
		panic(fmt.Sprintf("Error executing http request: %s\n", err))
	}
	if dataResponse.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Error executing http request: %s\n", dataResponse.Status))
	}
	var dataResponseBody DataResponse
	err = json.NewDecoder(dataResponse.Body).Decode(&dataResponseBody)
	if err != nil {
		panic(fmt.Sprintf("Data request failed - %s", err))
	}
	return dataResponseBody.CallData.Data
}

func MintTeamLiquid(account *account.Account) {
	instance := getInstance(account.Client)
	txData := account.TxData(0)
	var infos []ReservoirV601ExecutionInfo
	infos = append(infos, ReservoirV601ExecutionInfo{
		Module: account.Address(),
		Data:   []byte(getData(account)),
		Value:  big.NewInt(100000000000000),
	})
	tx, err := instance.Execute(txData, infos)
	if err != nil {
		panic(err)
	}
	log.Printf("[%s] Tx is sent: %s", account.Address(), tx.Hash().Hex())
}
