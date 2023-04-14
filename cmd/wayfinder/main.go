package main

import (
	"github.com/therobertcrocker/wayfinder/internal/controller"
	"github.com/therobertcrocker/wayfinder/internal/forge"

	// Import other modules
	"github.com/therobertcrocker/wayfinder/internal/world"
)

func main() {
	worldState, _ := world.NewJSONWorldState("internal/world/configuration/config.json")
	worldState.Init("internal/world/configuration/worldstate.json")
	appController := controller.NewAppController(worldState)

	questForge := forge.NewQuestForge(worldState)
	// Create other module instances

	appController.AddModule("questForge", questForge)
	// Add other modules

	appController.SwitchToModule("questForge") // Set the initial module

	if err := appController.Run(); err != nil {
		panic(err)
	}
}
