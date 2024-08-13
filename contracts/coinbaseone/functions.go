package coinbaseone

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

func getInstance(client *ethclient.Client, contractAddress string) (instance *Coinbaseone) {
	instance, err := NewCoinbaseone(
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

func Mint(account *account.Account, contractAddress string, value int) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress)
	if balance.Int64() == int64(0) {
		txData := account.TxData(int64(value))
		tx, err := instance.MintWithComment(txData, account.Address(), big.NewInt(1), "")
		if err != nil {
			panic(fmt.Sprintf("Mint error - %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), balance)
	}
}

func MintHbt(account *account.Account) {
	Mint(account, "0xE65dFa5C8B531544b5Ae4960AE0345456D87A47D", 100000000000000)
}

func MintCelebratingEth(account *account.Account) {
	Mint(account, "0xb5408b7126142C61f509046868B1273F96191b6d", 100000000000000)
}

func MintEurcEth(account *account.Account) {
	Mint(account, "0x615194d9695d0c02Fc30a897F8dA92E17403D61B", 100000000000000)
}

func MintEthCantStop(account *account.Account) {
	Mint(account, "0xb0FF351AD7b538452306d74fB7767EC019Fa10CF", 100000000000000)
}

func MintEthBreakingThrough(account *account.Account) {
	Mint(account, "0x96E82d88c07eCa6a29B2AD86623397B689380652", 100000000000000)
}

func MintPostEthEra(account *account.Account) {
	Mint(account, "0x955FdFdFd783C89Beb54c85f0a97F0904D85B86C", 100000000000000)
}

func MintEthEtf(account *account.Account) {
	Mint(account, "0xC00F7096357f09d9f5FE335CFD15065326229F66", 100000000000000)
}

func MintEtfereum(account *account.Account) {
	Mint(account, "0xE8aD8b2c5Ec79d4735026f95Ba7C10DCB0D3732B", 100000000000000)
}

func MintNounsMachine(account *account.Account) {
	Mint(account, "0x6A791493D1FC10ED17BB698114e86d279190AE33", 100000000000000)
}

func MintLaserNouns(account *account.Account) {
	Mint(account, "0x656C06dfCEcc978C898d3E80E471B3Ed6cC6b957", 100000000000000)
}

func MintCelebratingNonce(account *account.Account) {
	Mint(account, "0x5680eAD37A60604a12F821Bb9Da42858cbC346Fd", 100000000000000)
}

func MintCoffeeDays(account *account.Account) {
	Mint(account, "0xf16755b43eE1a458161f0faE5a9124729f4f6B1B", 100000000000000)
}

func MintNounMoon(account *account.Account) {
	Mint(account, "0x92440f15451f1058237a83B2fD64C67110C5146B", 600000000000000)
}

func MintHappyBornDay(account *account.Account) {
	Mint(account, "0x037A8002604cB2C1871204Ca868536B7D696df1d", 100000000000000)
}

func MintNounifyTheRockies(account *account.Account) {
	Mint(account, "0x306671092213C4d0da1a7bB5c31D5B4F9aB62246", 100000000000000)
}

func MintNounsForever(account *account.Account) {
	Mint(account, "0xcF74F48B71f2A8160aDa67D1720ce0F2778b5a28", 100000000000000)
}

func MintHappyNouniversary(account *account.Account) {
	Mint(account, "0x6414A4359848d2BF12B93483cd8A6ef6B03779ae", 100000000000000)
}

func MintNounishVibes(account *account.Account) {
	Mint(account, "0xCcbb9DC3FeCAf7a9cAe716eF1C16C8ca2f19a3D1", 100000000000000)
}

func MintHandOfNouns(account *account.Account) {
	Mint(account, "0x250d4678a1175113eC96e7DeB90584267026D443", 100000000000000)
}

func MintNounsEverywhere(account *account.Account) {
	Mint(account, "0x63197bb4dE33DA81FdB311Ef6395237fB0F65C7D", 100000000000000)
}

func MintBaseGodsAndMiggs(account *account.Account) {
	Mint(account, "0x25F98e990B6C0dBa5A109B92542F16DCbbD017C8", 100000000000000)
}

func MintCelebratingOfEndOfNouns(account *account.Account) {
	Mint(account, "0x7B28d9Efa325225666Aa6ddaC20c46420cd75871", 100000000000000)
}

func MintThirdNouniversary(account *account.Account) {
	Mint(account, "0xae954896B4d3B113C9FCe85f64387229291fb5a9", 100000000000000)
}

func MintHappyNouniversaryTwo(account *account.Account) {
	Mint(account, "0xE0fE6DD851187c62a79D00a211953Fe3B5Cec7FE", 100000000000000)
}
