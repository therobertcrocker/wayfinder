package config

import (
	"github.com/BurntSushi/toml"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
)

var Log = logging.Log

// LoggingConfig holds config values for logging.
type LoggingConfig struct {
	Level string `toml:"level"`
}

// DatabaseConfig holds config values for the database.
type DatabaseConfig struct {
	Path    string   `toml:"path"`
	Buckets []string `toml:"buckets"`
}

// Config brings together all of the config structs into one.
type Config struct {
	Database DatabaseConfig `toml:"database"`
	Logging  LoggingConfig  `toml:"logging"`
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
