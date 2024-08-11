package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Account struct {
	Id         int
	PrivateKey *ecdsa.PrivateKey
	Client     *ethclient.Client
	ChainId    *big.Int
	Proxy      string
	Active     bool
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
