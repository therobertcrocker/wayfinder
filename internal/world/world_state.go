package world

type WorldState interface {
	Init(filePath string) error
	Get(key string) interface{}
	Set(key string, value interface{})
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
}
