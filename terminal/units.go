package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
)

func Input(text string) (filename string) {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel(text).
		SetFieldWidth(50).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
	filename = inputField.GetText()
	return filename
}

func InputNumber(text string) (number int) {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel(text).
		SetFieldWidth(15).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
	number, err := strconv.Atoi(inputField.GetText())
	if err != nil {
		panic(err)
	}
	return number
}
