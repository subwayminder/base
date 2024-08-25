package account

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

type Account struct {
	Id         int
	PrivateKey *ecdsa.PrivateKey
	Client     *ethclient.Client
	ChainId    *big.Int
	Proxy      string
	Active     bool
}

type ClaimResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Client(rpc string) (client *ethclient.Client, chainId *big.Int) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(fmt.Sprintf("Oops! There was a problem: %s", err))
	}
	chainId, err = client.ChainID(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Problem occured while getting Chain ID", err))
	}
	return client, chainId
}

func (account *Account) Address() (address common.Address) {
	publicKey := account.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("error casting public key to ECDSA"))
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func (account *Account) Nonce() (nonce uint64) {
	nonce, err := account.Client.PendingNonceAt(context.Background(), account.Address())
	if err != nil {
		panic(err)
	}
	return nonce
}

func (account *Account) PersonalSign(message string) ([]byte, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), account.PrivateKey)
	if err != nil {
		return []byte(""), err
	}
	signatureBytes[64] += 27
	return signatureBytes, nil
}

func (account *Account) Gas() (gasPrice *big.Int) {
	gasPrice, err := account.Client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	return gasPrice
}

func (account *Account) ClaimBadge(badgeId string) (err error) {
	if badgeId == "" {
		return nil
	}
	url := "https://basehunt.xyz/api/challenges/complete"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
		"gameId": 2,
		"userAddress": "%s",
		"challengeId": "%s"
	}`, account.Address().String(), badgeId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var claimResponseBody ClaimResponseData
	if err := json.Unmarshal(body, &claimResponseBody); err != nil {
		log.Printf(fmt.Sprintf("Response - %s", string(body)))
		panic(fmt.Sprintf("Can not unmarshal JSON - %s", err))
	}
	log.Printf(fmt.Sprintf("Result - %s", string(body)))
	return nil
}

func (account *Account) TxData(value int64) (auth *bind.TransactOpts) {
	auth, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, account.ChainId)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(account.Nonce()))
	auth.Value = big.NewInt(value) // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = account.Gas()
	return auth
}
