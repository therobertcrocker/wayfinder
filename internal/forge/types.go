package forge

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/therobertcrocker/wayfinder/internal/wayfinder"
)

var (
	QuestTypes = map[string]int{
		"Hunt":        0,
		"Acquisition": 1,
		"Whisper":     2,
		"Knowledge":   3,
	}

	XPTable = map[int]int{
		-2: 100,
		-1: 200,
		0:  250,
		1:  300,
		2:  400,
		3:  500,
	}
	RewardMod = map[string][]int{
		"gold":       {2, -1, 0, 0},
		"treasure":   {-1, 2, -1, 0},
		"reputation": {0, 0, 2, 0},
	}
)

type Quest struct {
	Type    string `json:"class"`
	Level   int    `json:"level"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Source  string `json:"source"`
	Rewards Reward `json:"rewards"`
}

type Reward struct {
	XP             int `json:"XP_award"`
	Gold           int `json:"payment"`
	TreasureRating int `json:"treasure_rating"`
	Reputation     int `json:"reputation_bonus"`
}

type QuestData struct {
	Quest
}

func NewQuestData() *QuestData {
	return &QuestData{}
}

func (qd *QuestData) CalculateRewards(world wayfinder.WorldState) {

	index := QuestTypes[qd.Type]

	qd.Rewards.TreasureRating = RewardMod["treasure"][index]
	qd.Rewards.Reputation = RewardMod["reputation"][index]

	relativeLevel := qd.Level - world.Get("party_level").(int)
	qd.Rewards.XP = XPTable[relativeLevel]
	qd.Rewards.Gold = world.Get("gold_by_level").([]int)[qd.Level] / world.Get("party_size").(int)

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
