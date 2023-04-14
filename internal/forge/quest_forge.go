package forge

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/rivo/tview"
	"github.com/therobertcrocker/wayfinder/internal/world"
)

const (
	configPath = "internal/forge/config.json"
)

type QuestForge struct {
	config    *QuestConfig
	form      *tview.Form
	questData *QuestData
	world     world.JSONWorldState
}

func NewQuestForge(world world.JSONWorldState) (*QuestForge, error) {
	questConfig, err := NewQuestConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load quest config: %v", err)
	}
	return &QuestForge{

		form:      tview.NewForm(),
		questData: NewQuestData(questConfig),
		world:     world,
		config:    questConfig,
	}, nil
}

func NewQuestConfig(configPath string) (*QuestConfig, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config QuestConfig
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}

func (qf *QuestForge) Setup() tview.Primitive {
	qf.form.
		AddDropDown("Type", qf.config.QuestTypes, 0, func(option string, index int) {
			qf.questData.Type = option
		}).
		AddDropDown("Level", qf.world.LevelRange, 0, func(option string, index int) {
			qf.questData.Level, _ = strconv.Atoi(option)
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
	qf.questData.CalculateRewards(&qf.world)
	filePath := fmt.Sprintf("output/quests/%s_%s.json", qf.questData.Type, qf.questData.Title)
	qf.questData.SaveToFile(filePath)
}

func (qf *QuestForge) writeOnImage() {
	// Placeholder code - implement loading an image, drawing text on it, and saving the result
}
