package coinbaseforge

import (
	"base/account"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
	"time"
)

func getInstance(client *ethclient.Client, contractAddress string) (instance *Coinbaseforge) {
	instance, err := NewCoinbaseforge(
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

func MintWithComment(
	account *account.Account,
	contractAddress string,
	value int,
	quantity int,
	comment string,
	badgeId string,
) {
	instance := getInstance(account.Client, contractAddress)
	balance := BalanceOf(account, contractAddress)
	log.Printf("[%s] [%s] Current balance is: %s", strconv.Itoa(account.Id), account.Address(), balance.String())
	if balance.Int64() == int64(0) {
		txData := account.TxData(int64(value))
		tx, err := instance.MintWithComment(txData,
			account.Address(),
			big.NewInt(int64(quantity)),
			comment,
		)
		if err != nil {
			panic(fmt.Sprintf("Tx error: %s", err))
		}
		log.Printf("[%s] [%s] Tx is sent: %s", strconv.Itoa(account.Id), account.Address(), tx.Hash().Hex())
		time.Sleep(3 * time.Second)
		err = account.ClaimBadge(badgeId)
		if err != nil {
			log.Printf("[%s] [%s] Claim Badge error - %s", strconv.Itoa(account.Id), account.Address(), err)
		}
	} else {
		log.Printf("[%s] [%s] Already minted. Balance is %s", strconv.Itoa(account.Id), account.Address(), strconv.Itoa(int(balance.Int64())))
		err := account.ClaimBadge(badgeId)
		if err != nil {
			log.Printf("[%s] [%s] Claim Badge error - %s", strconv.Itoa(account.Id), account.Address(), err)
		}
	}
}

func MintNounsSummer(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x01F0D0D40D4BB499e7CB35940E908b74D08BA412",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintWalkingOnSunshine(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x4E4c12451A6e2473Fc4C63f84E175C3D31555F47",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintConsolationOfChromaConvergence(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x7d22F7EC034bB1060E033f132b0b23aA45b2B9B4",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintStrut(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x9FF8Fd82c0ce09caE76e777f47d536579AF2Fe7C",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintBaseNft(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x3b4B32a5c9A01763A0945A8a4a4269052DC3DE2F",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintToshiOnchainSummer(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x44d461Da8A451f05b6b75EdD5C4a2d2f3C14aaB4",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func Mintᐊᔮᑎᓯᐏᐣayâtisiwin(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x7B791EdF061Df65bAC7a9d47668F61F1a9A998C0",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintNounisvary(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xC8f93Ce7A12960466a2e13E70dE5CA41B652e4E6",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintThinkBig(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x752d593b3B8aD1c5d827F5B9AA9b653eE7134ea0",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintWhatchuLookingAt(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x5307c5ee9AeE0B944fA2E0Dba5D35D1D454E4bcE",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintStandWithPizza(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x4beAdC00E2A6b6C4fAc1a43FF340E5D71CBB9F77",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintEndaoment(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x4e4431BDdC2a896b1268ded02807b78c318C82e0",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintStandWithCrypto(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x146B627a763DFaE78f6A409CEF5B8ad84dDD4150",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintStandWithCryptoSongADay(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x2382456097cC12ce54052084e9357612497FD6be",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintEspressoAndMilk(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x52088a4416a9D437fC4a20a147103B4450e1bfd2",
		500000000000000,
		1,
		"",
		badgeId,
	)
}

func MintBaseTurnsOne(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x96C7464c73c24e3cBF3CcA8f3a2BFF917F39dC26",
		400000000000000,
		1,
		"",
		badgeId,
	)
}

func MintToshiVibe(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xbFa3fF9dcdB811037Bbec89f89E2751114ECD299",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintToshiChess(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xd60f13cC3e4d5bC96e7bAE8AAb5F448f3eFF3F0C",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintCryptoWillBloom(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x651b0A2b9FB9C186fB6C9a9CEddf25B791Ad5753",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintCryptoShield(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x651b0A2b9FB9C186fB6C9a9CEddf25B791Ad5753",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintTheVision(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x8605522B075aFeD48f9987E573E0AA8E572B8452",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintCryptoShieldRune(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x13fCcd944B1D88d0670cae18A00abD272256DDeE",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintShieldingTheWonder(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x6A3dA97Dc82c098038940Db5CB2Aa6B1541f2ebe",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintLiveAndLetLive(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x279dFFD2b14a4A60e266bEfb0D2c10E695D58113",
		500000000000000,
		1,
		"",
		badgeId,
	)
}

func MintSatoshiSummerRiddle(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x24da42FF6C4A6d5a6dCd1BecB204Bf82EA91d7B0",
		650000000000000,
		1,
		"",
		badgeId,
	)
}

func MintEarthStandsWithCrypto(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xd1E1da0b62761b0df8135aE4e925052C8f618458",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintGlassesStandsWithCrypto(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x03c6eF731453bfEc65a800F83f026ad011D8Abec",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintEnGarde(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x1f006edBc0Bcc528A743ee7A53b5e3dD393A1Df6",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintDualityWithMotion(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x5b45498D20d24D9c6Da165eDcd0eBcE0636176Ae",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintCreativeShield(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x892Bc2468f20D40F4424eE6A504e354D9D7E1866",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintWeStandWeBuild(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xEb9A3540E6A3dc31d982A47925d5831E02a3Fe1e",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintToshixSwc(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xb620bEdCe2615A3F35273A08b3e45e3431229A60",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintCryptoVibe(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x6a43B7e3ebFc915A8021dd05f07896bc092d1415",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintLetTheShieldShine(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x2a8e46E78BA9667c661326820801695dcf1c403E",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintAllForOne(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x8e50c64310b55729F8EE67c471E052B1Cd7AF5b3",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintLetsStand(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x95ff853A4C66a5068f1ED8Aaf7c6F4e3bDBEBAE1",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintForbesWeb3(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x0821D16eCb68FA7C623f0cD7c83C8D5Bd80bd822",
		0,
		1,
		"",
		badgeId,
	)
}

func MintSWC(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0x95167eB15a94DD048b2028c8d3fA3490f4cf8c76",
		211000000000000,
		1,
		"",
		badgeId,
	)
}

func MintWhatIfSWC(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xea50e58B518435AD2CeCE84d1e099b2e0878B9cF",
		100000000000000,
		1,
		"",
		badgeId,
	)
}

func MintNatureSWC(account *account.Account, badgeId string) {
	MintWithComment(
		account,
		"0xBB8F6319355d223C4a9f89a1b2A1c183B8Bf4EFF",
		100000000000000,
		1,
		"",
		badgeId,
	)
}
