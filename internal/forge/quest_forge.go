package forge

import (
	"strconv"

	"github.com/rivo/tview"
	"github.com/therobertcrocker/wayfinder/internal/wayfinder"
)

type QuestForge struct {
	form      *tview.Form
	questData *QuestData
	world     wayfinder.WorldState
}

func NewQuestForge(world wayfinder.WorldState) *QuestForge {
	return &QuestForge{
		form:      tview.NewForm(),
		questData: NewQuestData(),
		world:     world,
	}
}

func (qf *QuestForge) Setup() tview.Primitive {
	qf.form.
		AddDropDown("Type", qf.world.Get("quest_types").([]string), 0, func(option string, index int) {
			qf.questData.Type = option
		}).
		AddDropDown("Level", qf.world.Get("level_range").([]string), 0, func(option string, index int) {
			qf.questData.Level, _ = strconv.Atoi(option)

			qf.questData.CalculateRewards(qf.world)
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
