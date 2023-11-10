package bubble_tea

import (
	"fmt"
	"strings"

	// Import necessary packages for TUI
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define custom types and constants
type assetModel struct {
	inputs  []textinput.Model // To hold all the textinput models
	focused int               // To keep track of which input is focused
	// Additional fields and methods would go here
}

// Initialize a model with all inputs and their properties
func AssetInputModel() assetModel {
	inputs := make([]textinput.Model, 6) // Update the size to accommodate all inputs
	for i := range inputs {
		inputs[i] = textinput.New() // Create a new textinput.Model for each field
		// Common settings for all inputs can be configured here
	}

	// Configure each input field with specific properties
	inputs[0].Placeholder = "Enter character type..."
	inputs[1].Placeholder = "Enter primary attribute..."
	inputs[2].Placeholder = "Enter character qualities..."
	inputs[3].Placeholder = "Enter description..."
	inputs[4].Placeholder = "Enter credit amount..."
	inputs[5].Placeholder = "Enter current health..."

	// Set the first input as focused
	inputs[0].Focus()

	// Return the model with inputs and focused index set
	return assetModel{
		inputs:  inputs,
		focused: 0,
	}
}

// Model initialization: can return commands that should be run at the start
func (m assetModel) Init() tea.Cmd {
	return textinput.Blink // Start with the blinking cursor in the focused input
}

// Update is called when messages are received
func (m assetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd // To store commands for each input

	switch msg := msg.(type) {
	// Handle keyboard input messages
	case tea.KeyMsg:
		// Key messages to switch input focus or quit
		switch msg.Type {
		// Submit form
		case tea.KeyEnter:
			// Logic to handle form submission
			// For now, we'll just quit the program
			return m, tea.Quit
		// Quit the program
		case tea.KeyCtrlC:
			return m, tea.Quit
			// Navigate between inputs
			// ...
		}

		// Update the focused state of the inputs
		for i := 0; i < len(m.inputs); i++ {
			if i == m.focused {
				m.inputs[i].Focus() // Focus the currently focused input
			} else {
				m.inputs[i].Blur() // Remove focus from all other inputs
			}
			// Update each input and collect their commands
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}
	}

	// Return the model and a batch of collected commands
	return m, tea.Batch(cmds...)
}

// View renders the UI, which is displayed in the terminal
func (m assetModel) View() string {
	// Define styles for inputs
	inputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))

	// Build the form UI with a title and input fields
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n", "Character Creation Form")
	for _, input := range m.inputs {
		fmt.Fprintf(&b, "%s\n", inputStyle.Render(input.Placeholder))
		fmt.Fprintf(&b, "%s\n", input.View())
	}
	return b.String() // Return the UI as a string
}

// Additional methods such as nextInput and prevInput could be implemented here
// ...
