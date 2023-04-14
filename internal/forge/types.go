package forge

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/therobertcrocker/wayfinder/internal/world"
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
	config *QuestConfig
	Quest
}

func NewQuestData(config *QuestConfig) *QuestData {
	return &QuestData{
		config: config,
	}
}

func (qd *QuestData) CalculateRewards(world *world.JSONWorldState) {
	xpTable := qd.config.XPTable
	goldTable := qd.config.GoldByLevel
	rewardTable := qd.config.RewardTable
	relativeLevel := qd.Level - world.PartyLevel

	qd.Rewards.XP = int(xpTable[relativeLevel][0])

	goldMod := rewardTable["Gold"][qd.Type]
	qd.Rewards.Gold = goldTable[qd.Level+goldMod] / int(xpTable[relativeLevel][1]) / world.PartySize
	qd.Rewards.TreasureRating = rewardTable["Loot"][qd.Type]
	qd.Rewards.Reputation = rewardTable["Reputation"][qd.Type]

}

func (qd *QuestData) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(qd)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	return nil
}
