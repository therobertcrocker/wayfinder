package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	layout := tview.NewFlex()

	menu1 := tview.NewList()
	menu1.AddItem("Option 1", "", '1', func() {
		menu2 := tview.NewList()
		menu2.AddItem("Suboption A", "", 'a', nil)
		menu2.AddItem("Suboption B", "", 'b', nil)
		menu2.AddItem("Suboption C", "", 'c', nil)
		menu2.SetSelectedFunc(func(index int, label string, secondaryText string, shortcut rune) {
			// Handle sub-menu selection
			app.Stop()
		})
		app.SetRoot(menu2, true)
	})
	menu1.AddItem("Option 2", "", '2', nil)
	menu1.AddItem("Option 3", "", '3', nil)
	menu1.SetSelectedFunc(func(index int, label string, secondaryText string, shortcut rune) {
		// Handle menu1 selection
	})

	layout.AddItem(menu1, 0, 1, false)

	if err := app.SetRoot(layout, true).SetFocus(menu1).Run(); err != nil {
		panic(err)
	}
}
