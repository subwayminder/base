package msmiggles

import (
	"base/account"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
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

func getInstance(client *ethclient.Client) (instance *Msmiggles) {
	instance, err := NewMsmiggles(
		common.HexToAddress("0x1aeD60A97192157fDA7fb26267A439d523d09c5e"),
		client)
	if err != nil {
		panic(err)
	}
	return instance
}

func getData(account *account.Account) (data string) {
	body := []byte(fmt.Sprintf(
		`{"bypassSimulation":true,"mintAddress":"%s","network":"networks/base-mainnet","quantity":"%s","takerAddress":"%s"}`,
		"0xDc03a75F96f38615B3eB55F0F289d36E7A706660",
		"1",
		account.Address()))
	dataRequest, err := http.NewRequest(
		"POST",
		"https://api.wallet.coinbase.com/rpc/v3/creators/mintToken",
		bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Sprintf("Error init http request: %s", err))
	}
	dataRequest.Header.Set("Proxy", account.Proxy)
	dataRequest.Header.Add("Content-Type", "application/json")
	dataRequest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	dataRequest.Header.Add("Accept", "*/*")
	dataRequest.Header.Add("Accept-Encoding", "gzip, deflate, br")
	dataRequest.Header.Add("Connection", "keep-alive")
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

func MintMisterMiggles(account *account.Account) {
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
