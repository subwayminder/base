package reservoir

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

func getInstance(client *ethclient.Client, contractAddress string) (instance *Msmiggles) {
	instance, err := NewMsmiggles(
		common.HexToAddress(contractAddress),
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

func execute(account *account.Account, contractAddress string, module string, value int64, data []byte) {
	instance := getInstance(account.Client, contractAddress)
	txData := account.TxData(value)
	var infos []ReservoirV601ExecutionInfo
	infos = append(infos, ReservoirV601ExecutionInfo{
		Module: common.HexToAddress(module),
		Data:   data,
		Value:  big.NewInt(value),
	})
	tx, err := instance.Execute(txData, infos)
	if err != nil {
		panic(err)
	}
	log.Printf("[%s] Tx is sent: %s", account.Address(), tx.Hash().Hex())
}

func MintBuildation(account *account.Account, badgeId string) {
	address := strings.ToLower(account.Address().Hex())[2:]
	data := fmt.Sprintf("0xb510391f000000000000000000000000"+
		"%s"+
		"000000000000000000000000000000000000000000000000000000000000004"+
		"00000000000000000000000000000000000000000000000000000000000000"+
		"3e4e8d51ef5000000000000000000000000000000000000000000000000000000"+
		"0000000060000000000000000000000000"+
		"%s"+
		"00000000000000000000000000000000000000000000000000000000000000"+
		"0100000000000000000000000000000000000000000000000000000000000000"+
		"010000000000000000000000000000000000000000000000000000000000000"+
		"0200000000000000000000000000"+
		"c45ca58cfa181b038e06dd65eabbd1a68d3ccf300000000000000000000000000"+
		"000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000"+
		"5af3107a400000000000000000000000000000000000000000000000000000000000000002c000000000000000000"+
		"00000000c45ca58cfa181b038e06dd65eabbd1a68d3ccf300000000000000000000000000000000000000000000000000"+
		"000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000"+
		"00000000000000000000000000000000000000001a484bb1e42000000000000000000000000"+
		"%s"+
		"0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000"+
		"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"+
		"0000000000000000000000000000000000000000000000000000"+
		"5af3107a400000000000000000000000000000000000000000000000000000000000000000"+
		"c00000000000000000000000000000000000000000000000000000000000000"+
		"1800000000000000000000000000000000000000000000000000000000000000080"+
		"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"+
		"00000000000000000000000000000000000000000000000000005af3107a4000000000000000000000000000"+
		"eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000000000000000000000"+
		"1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"+
		"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"+
		"000000000000000000000000000000000000000001000000000000000000000000"+
		"55c88bb05602da94fce8feadc1cbebf5b72c2453"+
		"00000000000000000000000000000000000000000000000000005af3107a4000000000000000000000000000000000000000000"+
		"000000000000000000000000000000000000000000000000000000000000000000000000000000000", address, address, address)

	execute(
		account,
		"0x1aeD60A97192157fDA7fb26267A439d523d09c5e",
		"0x849Ef788b40Af342e2883C3112Dd636f03a4203E",
		200000000000000,
		[]byte(data),
	)
}
