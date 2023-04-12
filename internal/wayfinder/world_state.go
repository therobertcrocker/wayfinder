package wayfinder

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
