package helpers

import (
	accpkg "base/account"
	"crypto/ecdsa"
	"encoding/csv"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type AppDataset struct {
	MaxRetries              int
	MinSleepInSeconds       int
	MaxSleepInSeconds       int
	MinModuleSleepInSeconds int
	MaxModuleSleepInSeconds int
	Module                  func(account *accpkg.Account)
	Modules                 map[int]func(account *accpkg.Account)
}

type ModuleList struct {
	Title  string
	Module func(account *accpkg.Account)
}

func LoadPrivateKey(privateKeyHex string) (privateKey *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func GetUtcTimezone() (loc *time.Location) {
	loc, _ = time.LoadLocation("UTC")
	return loc
}

func LoadAccounts(filename string, rpc string) (accounts []accpkg.Account) {
	client, chainId := accpkg.Client(rpc)
	type Record struct {
		Private string `csv:"Private"`
		Proxy   string `csv:"Proxy"`
		Active  int    `csv:"Active"`
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})

	var records []Record
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	for id, record := range records {
		accounts = append(accounts, accpkg.Account{
			Id:         id + 1,
			PrivateKey: LoadPrivateKey(record.Private[2:]),
			Client:     client,
			ChainId:    chainId,
			Proxy:      record.Proxy,
			Active:     record.Active == 1,
		})
	}
	for i := range accounts {
		j := rand.Intn(i + 1)
		accounts[i], accounts[j] = accounts[j], accounts[i]
	}
	return accounts
}

func RunAccountsModules(dataset AppDataset, accounts []accpkg.Account) {
	var wg sync.WaitGroup
	for _, account := range accounts {
		if !account.Active {
			continue
		}
		if dataset.MinSleepInSeconds > 0 {
			time.Sleep(
				time.Duration(
					rand.Intn(dataset.MaxSleepInSeconds-dataset.MinSleepInSeconds)+dataset.MinSleepInSeconds) * time.Second)
		}
		log.Printf("[%s] [%s] Account started", strconv.Itoa(account.Id), account.Address())
		wg.Add(1)
		go RunAccountModules(dataset, account, &wg)
	}
	wg.Wait()
}

func RunAccountModules(dataset AppDataset, account accpkg.Account, accountWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	if accountWg != nil {
		defer accountWg.Done()
	}
	for _, module := range dataset.Modules {
		if dataset.MinModuleSleepInSeconds > 0 {
			time.Sleep(time.Duration(rand.Intn(dataset.MaxModuleSleepInSeconds-dataset.MinModuleSleepInSeconds)+dataset.MinModuleSleepInSeconds) * time.Second)
		}
		log.Printf("[%s] [%s] - %s started",
			strconv.Itoa(account.Id), account.Address(),
			runtime.FuncForPC(reflect.ValueOf(module).Pointer()).Name())
		wg.Add(1)
		go recoveryWrap(dataset.MaxRetries, module, &account, &wg, dataset)
	}
	wg.Wait()
}

func RunModuleLoop(dataset AppDataset, accounts []accpkg.Account, moduleWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	if moduleWg != nil {
		defer moduleWg.Done()
	}
	for _, account := range accounts {
		if !account.Active {
			continue
		}
		if dataset.MinSleepInSeconds > 0 {
			time.Sleep(
				time.Duration(
					rand.Intn(dataset.MaxSleepInSeconds-dataset.MinSleepInSeconds)+dataset.MinSleepInSeconds) * time.Second)
		}
		log.Printf("[%s] [%s] Mint started", strconv.Itoa(account.Id), account.Address())
		wg.Add(1)
		go recoveryWrap(dataset.MaxRetries, dataset.Module, &account, &wg, dataset)
	}
	wg.Wait()
}

func recoveryWrap(
	maxPanics int,
	f func(account *accpkg.Account),
	account *accpkg.Account,
	wg *sync.WaitGroup,
	dataset AppDataset) {

	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[%s] [%s] Panic - %s", strconv.Itoa(account.Id), account.Address(), err)
			if maxPanics == 0 {
				log.Printf("[%s] [%s] Fatal - too many panics", strconv.Itoa(account.Id), account.Address())
				return
			} else {
				wg.Add(1)
				go recoveryWrap(maxPanics-1, f, account, wg, dataset)
			}
		}
	}()
	f(account)
}
