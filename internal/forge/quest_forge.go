package forge

import (
	"github.com/rivo/tview"
	"github.com/therobertcrocker/wayfinder/internal/wayfinder"
)

type QuestForge struct {
	form       *tview.Form
	questData  *QuestData
	worldState wayfinder.WorldState
}

func NewQuestForge(world wayfinder.WorldState) *QuestForge {
	return &QuestForge{
		form:       tview.NewForm(),
		questData:  NewQuestData(),
		worldState: world,
	}
}

func (qf *QuestForge) Setup() tview.Primitive {
	qf.form.
		AddDropDown("Type", []string{"Hunt", "Acquisitions", "Whisper", "Knowledge"}, 0, func(option string, index int) {
			qf.questData.Type = option
		}).
		AddDropDown("Level", []string{"Level1", "Level2", "Level3"}, 0, func(option string, index int) {
			qf.questData.Level = option
		}).
		AddInputField("Title", "", 20, nil, func(value string) {
			qf.questData.Title = value
		}).
		AddTextArea("Summary", "", 0, 3, 0, func(value string) {
			qf.questData.Summary = value
		}).
		AddInputField("Source", "", 20, nil, func(value string) {
			qf.questData.Source = value
		}).
		AddInputField("Gold", "", 20, nil, func(value string) {
			qf.questData.Gold = value
		}).
		AddDropDown("Treasure Rating", []string{"Rating1", "Rating2", "Rating3"}, 0, func(option string, index int) {
			qf.questData.TreasureRating = option
		}).
		AddDropDown("Reputation", []string{"Reputation1", "Reputation2", "Reputation3"}, 0, func(option string, index int) {
			qf.questData.Reputation = option
		}).
		AddButton("Save JSON", qf.save).
		AddButton("Save Image", qf.writeOnImage).
		AddButton("Clear", func() {
			qf.form.Clear(true)
		})

	return qf.form
}

func (qf *QuestForge) save() {
	// Placeholder code - implement saving quest data to a JSON file
}

func (qf *QuestForge) writeOnImage() {
	// Placeholder code - implement loading an image, drawing text on it, and saving the result
}
