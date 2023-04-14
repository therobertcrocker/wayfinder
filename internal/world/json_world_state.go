package world

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type JSONWorldState struct {
	PartyLevel int
	PartySize  int
	LevelRange []string
	state      map[string]interface{}
}

func NewJSONWorldState(configPath string) (*JSONWorldState, error) {
	//Load JSON Config
	return &JSONWorldState{
		state: make(map[string]interface{}),
	}, nil
}

func (ws *JSONWorldState) Init(filePath string) error {
	// Calls Load From File and sets up the Party Level and Level Range
	err := ws.LoadFromFile(filePath)
	if err != nil {
		return err
	}

	ws.PartyLevel = int(ws.Get("party_level").(float64))
	ws.PartySize = int(ws.Get("party_size").(float64))
	ws.LevelRange = make([]string, 6)
	for i := -2; i <= 3; i++ {
		ws.LevelRange[i+2] = strconv.Itoa(ws.PartyLevel + i)
	}

	return nil
}

func (ws *JSONWorldState) Get(key string) interface{} {
	// Gets a Key and Value from the state
	return ws.state[key]
}
func (ws *JSONWorldState) Set(key string, value interface{}) {
	// Sets a Key and Value in the state
	ws.state[key] = value
}
func (ws *JSONWorldState) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(ws.state)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}
	return nil
}
func (ws *JSONWorldState) LoadFromFile(filename string) error {
	// Loads the Entire WorldState Struct as a JSON file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&ws.state)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	return nil
}
