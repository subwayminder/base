package terminal

import (
	"base/helpers"
	"fmt"
	"github.com/rivo/tview"
	"math"
	"os"
	"strconv"
)

func ChooseModule(dataset *helpers.AppDataset, modules []helpers.ModuleList) {
	app := tview.NewApplication()
	list := tview.NewList()

	for _, module := range modules {
		list.AddItem(module.Title, "", '*', func() {
			app.Stop()
			dataset.Module = module.Module
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	list.SetTitle("Choose module")
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func ChooseModulesForm(dataset *helpers.AppDataset, modules []helpers.ModuleList) {
	app := tview.NewApplication()
	pages := tview.NewPages()
	pageCount := int(math.Ceil(float64(len(modules)/10))) + 1
	for i := 1; i <= pageCount; i++ {
		form := tview.NewForm()
		k := 0
		for j, module := range modules {
			k++
			index := (i-1)*10 + k
			form.AddCheckbox(module.Title, false, func(checked bool) {
				if checked {
					dataset.Modules[index] = module.Module
				} else {
					delete(dataset.Modules, index)
				}
			})
			if k >= 10 {
				modules = modules[j+1:]
				break
			}
		}
		form.SetTitle(fmt.Sprintf("Choose module [Page %s/%s]", strconv.Itoa(i), strconv.Itoa(pageCount))).
			SetBorder(true)

		if i != 1 {
			form.AddButton("<", func() {
				pages.SwitchToPage(strconv.Itoa(i - 1))
			})
		}
		if i != pageCount {
			form.AddButton(">", func() {
				pages.SwitchToPage(strconv.Itoa(i + 1))
			})
		}
		form.AddButton("Run", func() {
			app.Stop()
		})
		form.AddButton("Exit", func() {
			defer os.Exit(0)
			app.Stop()
		})
		pages.AddPage(strconv.Itoa(i), form, true, true)
	}
	pages.SwitchToPage("1")
	pages.SetTitle("Choose modules").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
