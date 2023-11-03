package util

import (
	"github.com/BurntSushi/toml"
)

// DatabaseConfig holds config values for the database.
type DatabaseConfig struct {
	Path string `toml:"path"`
}

// Config brings together all of the config structs into one.
type Config struct {
	Database DatabaseConfig `toml:"database"`
}

// LoadConfig() loads the configuration from the given path.
func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		Log.WithField("path", path).Errorf("failed to load config: %v", err)
		return nil, err
	}
	return &config, nil
}
