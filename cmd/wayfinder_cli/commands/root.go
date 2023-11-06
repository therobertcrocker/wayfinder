/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/cmd/wayfinder_cli/commands/machinations"
	"github.com/therobertcrocker/wayfinder/internal/core"
)

var coreInstance *core.Core

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wayfinder",
	Short: "wayfinder - a suite of GM Tools",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(c *core.Core) {

	// Set the core instance
	coreInstance = c

	// Add the Machinations commands
	assetsCmd := machinations.AssetsCmd(coreInstance)
	rootCmd.AddCommand(assetsCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
