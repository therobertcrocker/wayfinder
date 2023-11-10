package main

import (
	log "github.com/sirupsen/logrus"
	cmd "github.com/therobertcrocker/wayfinder/cmd/wayfinder_cli/commands"
	"github.com/therobertcrocker/wayfinder/internal/common/config"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
	"github.com/therobertcrocker/wayfinder/internal/core"
)

var (
	Log = logging.Log
)

const (
	configPath = "config.toml"
)

func main() {

	// Load the configuration file
	config, err := config.LoadConfig(configPath)
	if err != nil {
		Log.Fatalf("failed to load config: %v", err)
	}

	// set the logging level
	switch config.Logging.Level {
	case "debug":
		Log.SetLevel(log.DebugLevel)
	case "info":
		Log.SetLevel(log.InfoLevel)
	case "warn":
		Log.SetLevel(log.WarnLevel)
	case "error":
		Log.SetLevel(log.ErrorLevel)
	case "fatal":
		Log.SetLevel(log.FatalLevel)
	default:
		Log.SetLevel(log.InfoLevel)
	}

	coreInstance := core.NewCore(config)
	cmd.Execute(coreInstance)
}
