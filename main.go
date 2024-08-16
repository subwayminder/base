package main

import (
	"base/account"
	_ "base/account"
	"base/contracts/coinbaseforge"
	"base/contracts/coinbaseone"
	"base/contracts/olympic"
	ser "base/contracts/seasonal_erosion_relic"
	"base/contracts/zora/zorafirst"
	"base/contracts/zora/zorasecond"
	"base/helpers"
	"base/terminal"
	"github.com/rivo/tview"
)

func main() {
	var modules []helpers.ModuleList
	modules = []helpers.ModuleList{
		{
			Title:  "Mint seasonal erosion relic",
			Module: ser.MintSeasonalErosion,
		},
		{
			Title:  "Mint Olympic relic",
			Module: olympic.MintOlympic,
		},
		{
			Title:  "Mint HBT",
			Module: coinbaseone.MintHbt,
		},
		{
			Title:  "Mint Celebrating Ethereum",
			Module: coinbaseone.MintCelebratingEth,
		},
		{
			Title:  "Mint Eurc eth",
			Module: coinbaseone.MintEurcEth,
		},
		{
			Title:  "Mint Eth can't stop",
			Module: coinbaseone.MintEthCantStop,
		},
		{
			Title:  "Mint Eth Breaking Through",
			Module: coinbaseone.MintEthBreakingThrough,
		},
		{
			Title:  "Mint Post eth era",
			Module: coinbaseone.MintPostEthEra,
		},
		{
			Title:  "Mint Eth etf",
			Module: coinbaseone.MintEthEtf,
		},
		{
			Title:  "Mint Ethfereum",
			Module: coinbaseone.MintEtfereum,
		},
		{
			Title:  "Mint Coinbase - Nouns Machine",
			Module: coinbaseone.MintNounsMachine,
		},
		{
			Title:  "Mint Coinbase - Laser Nouns",
			Module: coinbaseone.MintLaserNouns,
		},
		{
			Title:  "Mint Coinbase - Celebrating nonce",
			Module: coinbaseone.MintCelebratingNonce,
		},
		{
			Title:  "Mint Coinbase - Coffee days",
			Module: coinbaseone.MintCoffeeDays,
		},
		{
			Title:  "Mint Coinbase - Noun moon",
			Module: coinbaseone.MintNounMoon,
		},
		{
			Title:  "Mint Coinbase - Happy born day",
			Module: coinbaseone.MintHappyBornDay,
		},
		{
			Title:  "Mint Coinbase - Nounify the rockies",
			Module: coinbaseone.MintNounifyTheRockies,
		},
		{
			Title:  "Mint Coinbase - Nouns forever",
			Module: coinbaseone.MintNounsForever,
		},
		{
			Title:  "Mint Coinbase - Happy nouniversary",
			Module: coinbaseone.MintHappyNouniversary,
		},
		{
			Title:  "Mint Coinbase - Nounish vibes",
			Module: coinbaseone.MintNounishVibes,
		},
		{
			Title:  "Mint Coinbase - Hand of Nouns",
			Module: coinbaseone.MintHandOfNouns,
		},
		{
			Title:  "Mint Coinbase - Nouns everywhere",
			Module: coinbaseone.MintNounsEverywhere,
		},
		{
			Title:  "Mint Coinbase - Base gods and miggs",
			Module: coinbaseone.MintBaseGodsAndMiggs,
		},
		{
			Title:  "Mint Coinbase - Celebrating of end of nouns",
			Module: coinbaseone.MintCelebratingOfEndOfNouns,
		},
		{
			Title:  "Mint Coinbase - Third nouniversary",
			Module: coinbaseone.MintThirdNouniversary,
		},
		{
			Title:  "Mint Coinbase - Happy nouniversary",
			Module: coinbaseone.MintHappyNouniversaryTwo,
		},
		{
			Title:  "Mint Zora - EIC summer",
			Module: zorafirst.MintEicSummer,
		},
		{
			Title:  "Mint Zora - One Year On Base",
			Module: zorafirst.MintOneYearOnBase,
		},
		{
			Title:  "Mint Zora - Higher",
			Module: zorasecond.MintHigher,
		},
		{
			Title:  "Mint Zora - Bouquet",
			Module: zorasecond.MintBouquet,
		},
		{
			Title:  "Mint Zora - Hodlhq",
			Module: zorasecond.MintHodlhq,
		},
		{
			Title:  "Mint Zora - Fullmetaldroid",
			Module: zorasecond.MintFullmetaldroid,
		},
		{
			Title:  "Mint Zora - Essence",
			Module: zorasecond.MintEssence,
		},
		{
			Title:  "Mint Zora - Base Summer Genesis",
			Module: zorasecond.MintBaseSummerGenesis,
		},
		{
			Title:  "Mint Zora - Base Onchainomics",
			Module: zorasecond.MintOnchainomics,
		},
		{
			Title:  "Mint Zora - Stay Based",
			Module: zorasecond.MintStayBased,
		},
		{
			Title:  "Mint Zora - Palomar Group Archive",
			Module: zorasecond.MintPalomarGroupArchive,
		},
		{
			Title:  "Mint Zora - Building Onchain Podcast",
			Module: zorasecond.MintBuildingOnchainPodcast,
		},
		{
			Title:  "Mint Zora - Onchainomics Essay",
			Module: zorasecond.MintOnchainomicsEssay,
		},
		{
			Title:  "Mint Zora - Based On Happines",
			Module: zorasecond.MintBasedOnHappines,
		},
		{
			Title:  "Mint Zora - Summer Basecamp",
			Module: zorasecond.MintSummerBasecamp,
		},
		{
			Title:  "Mint Zora - Persona On Base",
			Module: zorasecond.MintPersonaOnBase,
		},
		{
			Title:  "Mint Zora - Girl Onchain",
			Module: zorasecond.MintGirlOnchain,
		},
		{
			Title:  "Mint Zora - Zeropod",
			Module: zorasecond.MintZeropod,
		},
		{
			Title:  "Mint Zora - Spiral 360",
			Module: zorasecond.MintSpiral360,
		},
		{
			Title:  "Mint Zora - Memloop",
			Module: zorasecond.MintMemloop,
		},
		{
			Title:  "Mint Coinbase - Nouns summer",
			Module: coinbaseforge.MintNounsSummer,
		},
		{
			Title:  "Mint Coinbase - Walking On Sunshine",
			Module: coinbaseforge.MintWalkingOnSunshine,
		},
		{
			Title:  "Mint Coinbase - Consolation of Chroma Convergence",
			Module: coinbaseforge.MintConsolationOfChromaConvergence,
		},
	}

	filename := terminal.Input("Enter a filename (default: accounts.csv): ")
	if filename == "" {
		filename = "accounts.csv"
	}
	rpc := terminal.Input("Enter RPC server address (default is https://base.drpc.org): ")
	if rpc == "" {
		rpc = "https://base.drpc.org"
	}

	accounts := helpers.LoadAccounts(filename, rpc)
	dataset := helpers.AppDataset{
		MaxRetries:        terminal.InputNumber("Input max retries: "),
		MinSleepInSeconds: terminal.InputNumber("Input min sleep seconds between accounts: "),
		MaxSleepInSeconds: terminal.InputNumber("Input max sleep seconds between accounts: "),
		Module:            nil,
		Modules:           make(map[int]func(account *account.Account)),
	}
	app := tview.NewApplication()
	list := tview.NewList()
	list.
		AddItem("Run one module", "", '1', func() {
			app.Stop()
			terminal.ChooseModule(&dataset, modules)
			helpers.RunModuleLoop(dataset, accounts, nil)
		}).
		AddItem("Run multiple modules", "", '2', func() {
			app.Stop()
			dataset.MinModuleSleepInSeconds = terminal.InputNumber("Input min sleep seconds between modules: ")
			dataset.MaxModuleSleepInSeconds = terminal.InputNumber("Input max sleep seconds between modules: ")
			terminal.ChooseModulesForm(&dataset, modules)
			helpers.RunAccountsModules(dataset, accounts)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	list.SetTitle("Choose mode:")
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
