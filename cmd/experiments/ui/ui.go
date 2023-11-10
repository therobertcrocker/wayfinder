package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/therobertcrocker/wayfinder/cmd/experiments/ui/bubble_tea"
)

func main() {
	// Start the program with the initial model

	p := tea.NewProgram(bubble_tea.AssetInputModel())

	// Run the program and handle any errors that may occur
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
