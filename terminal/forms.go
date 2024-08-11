package terminal

import (
	"base/helpers"
	"github.com/rivo/tview"
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
	form := tview.NewForm()

	for i, module := range modules {
		form.AddCheckbox(module.Title, false, func(checked bool) {
			if checked {
				dataset.Modules[i] = module.Module
			} else {
				delete(dataset.Modules, i)
			}
		})
	}
	form.
		AddButton("Continue", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Choose modules").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
