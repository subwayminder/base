package main

import (
	_ "base/account"
	"base/contracts/coinbaseforge"
	"base/contracts/coinbaseone"
	"base/contracts/olympic"
	ser "base/contracts/seasonal_erosion_relic"
	"base/contracts/thirdweb"
	"base/contracts/zora/zorafirst"
	"base/contracts/zora/zorasecond"
	"base/helpers"
	"base/terminal"
	"github.com/rivo/tview"
)

func main() {
	modules := []helpers.ModuleList{
		{
			Title: "Mint seasonal erosion relic",
			Module: helpers.ModuleUnit{
				Function: ser.MintSeasonalErosion,
				BadgeId:  "",
			},
		},
		{
			Title: "Mint Olympic relic",
			Module: helpers.ModuleUnit{
				Function: olympic.MintOlympic,
				BadgeId:  "",
			},
		},
		{
			Title: "Mint HBT",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintHbt,
				BadgeId:  "1pjoNf5onjgsi7r9fWp3ob",
			},
		},
		{
			Title: "Mint Celebrating Ethereum",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintCelebratingEth,
				BadgeId:  "5e383RWcRtGAwGUorkGiYC",
			},
		},
		{
			Title: "Mint Eurc eth",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintEurcEth,
				BadgeId:  "1iZiHPbqaIGW5F08bCit6J",
			},
		},
		{
			Title: "Mint Eth can't stop",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintEthCantStop,
				BadgeId:  "ocsChallenge_c1de2373-35ad-4f3c-ab18-4dfadf15754d",
			},
		},
		{
			Title: "Mint Eth Breaking Through",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintEthBreakingThrough,
				BadgeId:  "78AUXYw8UCyFUPE2zy9yMZ",
			},
		},
		{
			Title: "Mint Post eth era",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintPostEthEra,
				BadgeId:  "ocsChallenge_65c17605-e085-4528-b4f1-76ce5f48da56",
			},
		},
		{
			Title: "Mint Eth etf",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintEthEtf,
				BadgeId:  "ocsChallenge_ee0cf23e-74a1-4bb3-badf-037a6bbf35e8",
			},
		},
		{
			Title: "Mint Ethfereum",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintEtfereum,
				BadgeId:  "ocsChallenge_eba9e6f0-b7b6-4d18-8b99-a64aea045117",
			},
		},
		{
			Title: "Mint Coinbase - Nouns Machine",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounsMachine,
				BadgeId:  "7mRx7h33KuMmTbus0Bv4AE",
			},
		},
		{
			Title: "Mint Coinbase - Laser Nouns",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintLaserNouns,
				BadgeId:  "ocsChallenge_fe16c31b-4c59-45d8-9431-8276f835cb16",
			},
		},
		{
			Title: "Mint Coinbase - Celebrating nonce",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintCelebratingNonce,
				BadgeId:  "6VA9MQosJnPcCwEeNkNVsW",
			},
		},
		{
			Title: "Mint Coinbase - Coffee days",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintCoffeeDays,
				BadgeId:  "ocsChallenge_9142cba1-ec12-4ee8-915e-7976536908cd",
			},
		},
		{
			Title: "Mint Coinbase - Noun moon",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounMoon,
				BadgeId:  "ocsChallenge_d9e15d6a-3c6e-4585-91ff-ba9e157f3789",
			},
		},
		{
			Title: "Mint Coinbase - Happy born day",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintHappyBornDay,
				BadgeId:  "2U6YRYo55Zgfcxull5C1Xn",
			},
		},
		{
			Title: "Mint Coinbase - Nounify the rockies",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounifyTheRockies,
				BadgeId:  "21pui4pvJ0h6YA8EAlvjqh",
			},
		},
		{
			Title: "Mint Coinbase - Nouns forever",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounsForever,
				BadgeId:  "3M9bT5pBKJE6jgolwMFsJU",
			},
		},
		{
			Title: "Mint Coinbase - Happy nouniversary",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintHappyNouniversary,
				BadgeId:  "71QheN8IVzfyoVtE8oeHNU",
			},
		},
		{
			Title: "Mint Coinbase - Nounish vibes",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounishVibes,
				BadgeId:  "2r8tpvuVPkYIuhAWSoMYY1",
			},
		},
		{
			Title: "Mint Coinbase - Hand of Nouns",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintHandOfNouns,
				BadgeId:  "2qOcpUCs12XwgLUpoQQgYTT",
			},
		},
		{
			Title: "Mint Coinbase - Nouns everywhere",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintNounsEverywhere,
				BadgeId:  "ocsChallenge_d122a59b-6a04-4949-9cdb-b9262e843aa6",
			},
		},
		{
			Title: "Mint Coinbase - Base gods and miggs",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintBaseGodsAndMiggs,
				BadgeId:  "1eeRIVPiOVBJ3rlM5sGnpx",
			},
		},
		{
			Title: "Mint Coinbase - Celebrating of end of nouns",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintCelebratingOfEndOfNouns,
				BadgeId:  "ocsChallenge_15259510-7040-4f16-bdfa-f137846b546c",
			},
		},
		{
			Title: "Mint Coinbase - Third nouniversary",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintThirdNouniversary,
				BadgeId:  "7ktuPuO5kUtvQmvzd4T5r3",
			},
		},
		{
			Title: "Mint Coinbase - Happy nouniversary",
			Module: helpers.ModuleUnit{
				Function: coinbaseone.MintHappyNouniversaryTwo,
				BadgeId:  "44wp1P8LSnwkPSz7Ft3q78",
			},
		},
		{
			Title: "Mint Zora - EIC summer",
			Module: helpers.ModuleUnit{
				Function: zorafirst.MintEicSummer,
				BadgeId:  "ocsChallenge_1287a985-e4b1-4f30-b508-306c4d109832",
			},
		},
		{
			Title: "Mint Zora - One Year On Base",
			Module: helpers.ModuleUnit{
				Function: zorafirst.MintOneYearOnBase,
				BadgeId:  "TqnsSrVTggFzdu9vFuSju",
			},
		},
		{
			Title: "Mint Zora - Higher",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintHigher,
				BadgeId:  "ocsChallenge_0c81b9f8-4cd0-4ad1-8eed-692afb988995",
			},
		},
		{
			Title: "Mint Zora - Bouquet",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintBouquet,
				BadgeId:  "ocsChallenge_1d2ae291-a51e-44e9-ab92-bd2540262b1b",
			},
		},
		{
			Title: "Mint Zora - Hodlhq",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintHodlhq,
				BadgeId:  "ocsChallenge_851731af-e339-4a5e-8aa2-60bfca2079bf",
			},
		},
		{
			Title: "Mint Zora - Fullmetaldroid",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintFullmetaldroid,
				BadgeId:  "ocsChallenge_7ef196df-e2cf-4d56-b381-09af0f1984e2",
			},
		},
		{
			Title: "Mint Zora - Essence",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintEssence,
				BadgeId:  "ocsChallenge_092780c6-704c-4ec9-8986-3049cb80acef",
			},
		},
		{
			Title: "Mint Zora - Base Summer Genesis",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintBaseSummerGenesis,
				BadgeId:  "",
			},
		},
		{
			Title: "Mint Zora - Base Onchainomics",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintOnchainomics,
				BadgeId:  "ocsChallenge_fc73d959-f1d6-4630-ab10-8d697678ff0c",
			},
		},
		{
			Title: "Mint Zora - Stay Based",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintStayBased,
				BadgeId:  "ocsChallenge_ba17be7a-b8f3-4213-aa99-d07d8550bff8",
			},
		},
		{
			Title: "Mint Zora - Palomar Group Archive",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintPalomarGroupArchive,
				BadgeId:  "ocsChallenge_1c4878b7-f846-44cd-a2fe-5c91b1e27cd6",
			},
		},
		{
			Title: "Mint Zora - Building Onchain Podcast",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintBuildingOnchainPodcast,
				BadgeId:  "ocsChallenge_11b6a5a1-7180-450b-8c7c-77d37560dd10",
			},
		},
		{
			Title: "Mint Zora - Onchainomics Essay",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintOnchainomicsEssay,
				BadgeId:  "ocsChallenge_d9d8342f-99b1-4e55-98f4-9bb4e0a71ca1",
			},
		},
		{
			Title: "Mint Zora - Based On Happines",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintBasedOnHappines,
				BadgeId:  "ocsChallenge_546d588b-9daa-4e13-bbeb-d1ae9646179e",
			},
		},
		{
			Title: "Mint Zora - Summer Basecamp",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintSummerBasecamp,
				BadgeId:  "ocsChallenge_15b1524f-de7e-43d5-a9cb-03c9f24d0881",
			},
		},
		{
			Title: "Mint Zora - Persona On Base",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintPersonaOnBase,
				BadgeId:  "",
			},
		},
		{
			Title: "Mint Zora - Girl Onchain",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintGirlOnchain,
				BadgeId:  "ocsChallenge_475b4248-2c7d-45fb-a062-45d6eb18fd67",
			},
		},
		{
			Title: "Mint Zora - Zeropod",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintZeropod,
				BadgeId:  "ocsChallenge_acbbf04b-d703-4ce5-b1ed-32bb461168c7",
			},
		},
		{
			Title: "Mint Zora - Spiral 360",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintSpiral360,
				BadgeId:  "ocsChallenge_277d663c-ceb5-424c-ba69-b96334ff5516",
			},
		},
		{
			Title: "Mint Zora - Memloop",
			Module: helpers.ModuleUnit{
				Function: zorasecond.MintMemloop,
				BadgeId:  "ocsChallenge_8995cf46-b23e-4e46-afae-062655648951",
			},
		},
		{
			Title: "Mint Coinbase - Nouns summer",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintNounsSummer,
				BadgeId:  "TqnsSrVTggFzdu9vFuSju",
			},
		},
		{
			Title: "Mint Coinbase - Walking On Sunshine",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintWalkingOnSunshine,
				BadgeId:  "22B38pVH5SN5yAiW6HYdNa",
			},
		},
		{
			Title: "Mint Coinbase - Consolation of Chroma Convergence",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintConsolationOfChromaConvergence,
				BadgeId:  "6FfxTruNSVDFLtvD3hb9sT",
			},
		},
		{
			Title: "Mint Coinbase - Strut",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintStrut,
				BadgeId:  "5c3PqZ2EGVbzQ2CQXL1vWK",
			},
		},
		{
			Title: "Mint Coinbase - STIX",
			Module: helpers.ModuleUnit{
				Function: thirdweb.MintStix,
				BadgeId:  "ocsChallenge_bd5208b5-ff1e-4f5b-8522-c4d4ebb795b7",
			},
		},
		{
			Title: "Mint Coinbase - Base NFT",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintBaseNft,
				BadgeId:  "6UuHdstl9MRFd4cgFf15kk",
			},
		},
		{
			Title: "Mint Coinbase - Toshi Onchain Summer",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintToshiOnchainSummer,
				BadgeId:  "4g9y1NVvIxiOCSeGvPLmJS",
			},
		},
		{
			Title: "Mint Coinbase - ᐊᔮᑎᓯᐏᐣayâtisiwin",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.Mintᐊᔮᑎᓯᐏᐣayâtisiwin,
				BadgeId:  "1BWyKWI2UZHnOEw8E4hpS5",
			},
		},
		{
			Title: "Mint Coinbase - Nounisvary",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintNounisvary,
				BadgeId:  "ocsChallenge_578c4c33-4506-4604-8359-ac0b43a3809c",
			},
		},
		{
			Title: "Mint Coinbase - Summer Serenity",
			Module: helpers.ModuleUnit{
				Function: zorafirst.MintSummerSerenity,
				BadgeId:  "ocsChallenge_e490f19f-1923-479d-83b8-25ecdc3b8c4a",
			},
		},
		{
			Title: "Mint Coinbase - Think Big",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintThinkBig,
				BadgeId:  "3EOQYszODyvZvbQMoKPoDO",
			},
		},
		{
			Title: "Mint Coinbase - Think Big",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintThinkBig,
				BadgeId:  "3EOQYszODyvZvbQMoKPoDO",
			},
		},
		{
			Title: "Mint Coinbase - Whatchu Lookin' At?",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintWhatchuLookingAt,
				BadgeId:  "39XYCR1jsdPwnoFEpwCwhD",
			},
		},
		{
			Title: "Mint Coinbase - Stand with Crypto Pizza?",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintStandWithPizza,
				BadgeId:  "1zbecUKJMKwyYoKOn2OV5n",
			},
		},
		{
			Title: "Mint Coinbase - Endaoment X SWC Shield",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintEndaoment,
				BadgeId:  "359X8U2xzQmVIQRe7xSFk9",
			},
		},
		{
			Title: "Mint Coinbase - Stand with Crypto Pizza 2",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintStandWithCrypto,
				BadgeId:  "3ofLIMuInVt5sKkQOtLWp0",
			},
		},
		{
			Title: "Mint Coinbase - Stand With Crypto | Song A Day #5714",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintStandWithCryptoSongADay,
				BadgeId:  "5Hyw2HMBfOBFDvCBkvdVmX",
			},
		},
		{
			Title: "Mint Coinbase - Espresso and Milk | Coffee Days 2024",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintEspressoAndMilk,
				BadgeId:  "ocsChallenge_668763d8-477c-444b-91e1-346a69a8146d",
			},
		},
		{
			Title: "Mint Coinbase - Base Turns One",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintBaseTurnsOne,
				BadgeId:  "ocsChallenge_3809bf6d-a2d1-4e15-84e7-74beac310661",
			},
		},
		{
			Title: "Mint Coinbase - Toshi Vibe",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintToshiVibe,
				BadgeId:  "3WE9nylUC2bMHz9c6hxFnL",
			},
		},
		{
			Title: "Mint Coinbase - Toshi Chess",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintToshiChess,
				BadgeId:  "1HMONONDaMukjieAOD3PHQ",
			},
		},
		{
			Title: "Mint Coinbase - Crypto with bloom",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintCryptoWillBloom,
				BadgeId:  "S3DyUSaz6mYehsypyOqPD",
			},
		},
		{
			Title: "Mint Coinbase - Introducing: Coinbase Wallet web app",
			Module: helpers.ModuleUnit{
				Function: thirdweb.MintIntroductionCoinbaseWallet,
				BadgeId:  "78zcHkWSABcPWMoacVI9Vs",
			},
		},
		{
			Title: "Mint Coinbase - What if we added a Stand With Crypto Shield....",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintCryptoShield,
				BadgeId:  "71fCEn2cIwqXqLE6wYxGl0",
			},
		},
		{
			Title: "Mint Coinbase - Mint the vision",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintTheVision,
				BadgeId:  "3fYDO7ZCCl91Tg7u6cMHBa",
			},
		},
		{
			Title: "Mint Coinbase - Stand With Crypto Shield Rune",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintCryptoShieldRune,
				BadgeId:  "1FH5jNuTVIRrPBUNtKrFtQ",
			},
		},
		{
			Title: "Mint Coinbase - Shielding the wonder",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintShieldingTheWonder,
				BadgeId:  "4EHmwNf6hrcTFnume2AUhv",
			},
		},
		{
			Title: "Mint Coinbase - Live and Let Live!",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintLiveAndLetLive,
				BadgeId:  "4MMQPGoZviSqLoJgaVDY05",
			},
		},
		{
			Title: "Mint Coinbase - Satoshi's Summer Riddle Starter Pack NFT",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintSatoshiSummerRiddle,
				BadgeId:  "6al1rUlphrMNYqQRbG9l83",
			},
		},
		{
			Title: "Mint Coinbase - Earth Stands with Crypto",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintEarthStandsWithCrypto,
				BadgeId:  "7JZn2HJuvLZRoE8R8a8OBp",
			},
		},
		{
			Title: "Mint Coinbase - ⌐◨-◨ Stand With Crypto",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintGlassesStandsWithCrypto,
				BadgeId:  "4JS8wKnPtZ0lE34C5crIUk",
			},
		},
		{
			Title: "Mint Coinbase - en garde",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintEnGarde,
				BadgeId:  "3wnaF1zwkxXMotK2grz0kO",
			},
		},
		{
			Title: "Mint Coinbase - duality in motion",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintDualityWithMotion,
				BadgeId:  "3Po39fHlC66muE3X5IHNfs",
			},
		},
		{
			Title: "Mint Coinbase - The Creative Shield",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintCreativeShield,
				BadgeId:  "6kv6tqF4mCQiGn5SQUwdps",
			},
		},
		{
			Title: "Mint Coinbase - we stand, we build",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintWeStandWeBuild,
				BadgeId:  "43EAydXs7EVNGkR9UZ5JJH",
			},
		},
		{
			Title: "Mint Coinbase - Toshi x SWC 3",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintToshixSwc,
				BadgeId:  "1VaefmSAUYw5vW1lxc0Viq",
			},
		},
		{
			Title: "Mint Coinbase - Crypto Vibe",
			Module: helpers.ModuleUnit{
				Function: coinbaseforge.MintCryptoVibe,
				BadgeId:  "OE6zO6T5M3COHSFcIIvmA",
			},
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
		Module:            helpers.ModuleList{},
		Modules:           make(map[int]helpers.ModuleList),
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
		AddItem("Claim badges", "", '3', func() {
			app.Stop()
			dataset.MinModuleSleepInSeconds = terminal.InputNumber("Input min sleep seconds between modules: ")
			dataset.MaxModuleSleepInSeconds = terminal.InputNumber("Input max sleep seconds between modules: ")
			terminal.ChooseModulesForm(&dataset, modules)
			for i, module := range dataset.Modules {
				module.Module.Function = helpers.ClaimBadge
				dataset.Modules[i] = module
			}
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
