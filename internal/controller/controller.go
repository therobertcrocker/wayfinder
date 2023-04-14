package controller

import (
	"github.com/rivo/tview"
	"github.com/therobertcrocker/wayfinder/internal/world"
)

type ModuleController interface {
	Setup() tview.Primitive
}

type AppController struct {
	app          *tview.Application
	pages        *tview.Pages
	world        world.WorldState
	modules      []string
	currentIndex int
}

func NewAppController(world world.WorldState) *AppController {
	return &AppController{
		app:     tview.NewApplication().EnableMouse(true),
		pages:   tview.NewPages(),
		modules: []string{},
		world:   world,
	}
}

func (c *AppController) AddModule(name string, module ModuleController) {
	view := module.Setup()
	c.pages.AddPage(name, view, true, false)
	c.modules = append(c.modules, name)
}

func (c *AppController) SwitchToModule(name string) {
	c.pages.SwitchToPage(name)
}

func (c *AppController) SetupMenu() {
	menu := tview.NewList()

	for _, moduleName := range c.modules {
		menu.AddItem(moduleName, "", 0, func() {
			c.SwitchToModule(moduleName)
		})
	}
	quitButton := tview.NewButton("Quit").SetSelectedFunc(func() {
		c.app.Stop()
	})

	topBar := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(menu, 0, 3, false).
		AddItem(nil, 0, 1, false)

	bottomBar := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(quitButton, 0, 3, false).
		AddItem(nil, 0, 1, false)

	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(topBar, 1, 1, false).
		AddItem(tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(c.pages, 0, 3, true).
			AddItem(nil, 0, 1, false), 0, 1, false).
		AddItem(bottomBar, 1, 1, false)

	c.app.SetRoot(mainFlex, true).SetFocus(c.pages)
}

func (c *AppController) Run() error {
	c.SetupMenu()
	return c.app.Run()
}
