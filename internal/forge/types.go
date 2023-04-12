package forge

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Quest struct {
	Type           string `json:"class"`
	Level          string `json:"level"`
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	Source         string `json:"source"`
	Gold           string `json:"payment"`
	TreasureRating string `json:"treasure_rating"`
	Reputation     string `json:"reputation_bonus"`
}

type QuestData struct {
	Quest
}

func NewQuestData() *QuestData {
	return &QuestData{}
}

func (qd *QuestData) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(qd, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
