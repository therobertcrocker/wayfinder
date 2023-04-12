package wayfinder

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type WorldState interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}

type InMemoryWorldState struct {
	state map[string]interface{}
}

func NewInMemoryWorldState() *InMemoryWorldState {
	return &InMemoryWorldState{
		state: make(map[string]interface{}),
	}
}

func (ws *InMemoryWorldState) Init() {
	ws.Set("quest_types", []string{"Hunt", "Acquisitions", "Whisper", "Knowledge"})
	ws.Set("party_level", 5)
	partyLevel := ws.Get("party_level").(int)
	levelRange := []string{
		strconv.Itoa(partyLevel - 2),
		strconv.Itoa(partyLevel - 1),
		strconv.Itoa(partyLevel),
		strconv.Itoa(partyLevel + 1),
		strconv.Itoa(partyLevel + 2),
		strconv.Itoa(partyLevel + 3)}
	ws.Set("level_range", levelRange)
	ws.Set("party_size", 3)
	ws.Set("gold_by_level", []int{0, 40, 70, 120, 200, 320, 500, 720, 1000, 1400, 2000, 2800, 4000, 6000, 9000, 13000, 20000, 30000, 48000, 80000, 14000})

}

func (ws *InMemoryWorldState) Get(key string) interface{} {
	return ws.state[key]
}

func (ws *InMemoryWorldState) Set(key string, value interface{}) {
	ws.state[key] = value
}

func (ws *InMemoryWorldState) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(ws.state, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (ws *InMemoryWorldState) LoadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &ws.state)
	if err != nil {
		return err
	}

	return nil
}
