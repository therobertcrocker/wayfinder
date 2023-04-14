package forge

type QuestConfig struct {
	QuestTypes  []string
	XPTable     map[int][]float64
	RewardTable map[string]map[string]int
	GoldByLevel []int
}

type RewardMod struct {
	Key       string
	ModValues int
}
