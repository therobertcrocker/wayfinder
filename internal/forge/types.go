package forge

type Quest struct {
	Type    string
	Level   int
	Title   string
	Summary string
	Source  string
	Reward  reward
}

type reward struct {
	Gold           int
	TreasureRating int
	Reputation     int
}
