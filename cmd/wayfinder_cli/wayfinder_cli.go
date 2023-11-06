package main

import (
	cmd "github.com/therobertcrocker/wayfinder/cmd/wayfinder_cli/commands"
	"github.com/therobertcrocker/wayfinder/internal/core"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

const (
	configPath = "config.toml"
)

func main() {

	// Load the configuration file
	config, err := util.LoadConfig(configPath)
	if err != nil {
		util.Log.Fatalf("failed to load config: %v", err)
	}

	coreInstance := core.NewCore(config)
	cmd.Execute(coreInstance)
}
