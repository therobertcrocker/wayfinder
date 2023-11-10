package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validConfigPath = "testdata/valid_config.toml"
	validConfig     = &Config{
		Database: DatabaseConfig{
			Path: "test",
		},
	}
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name           string
		configPath     string
		expectedConfig *Config
		wantErr        bool
	}{
		{
			name:           "ValidConfig",
			configPath:     validConfigPath,
			expectedConfig: validConfig,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// load the config
			config, err := LoadConfig(tt.configPath)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedConfig, config)
			}

		})

	}

}
