package assets

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/internal/core"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

var (
	coreInstance *core.Core
	filePath     string
	fileType     string
)

func AddAssetsCmd(c *core.Core) *cobra.Command {
	coreInstance = c

	addAssetsCmd := &cobra.Command{
		Use:   "load",
		Short: "Load Assets from a file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if filePath == "" {
				util.Log.Fatalf("no file path provided")
			}
			if err := coreInstance.Machinations.DataStore.LoadAssets(filePath, fileType); err != nil {
				util.Log.Fatalf("failed to load assets: %v", err)
			}

			util.Log.Infof("successfully loaded assets from %s", filePath)
		},
	}

	addAssetsCmd.Flags().StringVarP(&filePath, "path", "p", "", "path to the file containing the assets")
	addAssetsCmd.Flags().StringVarP(&fileType, "type", "t", "", "type of file containing the assets")
	addAssetsCmd.MarkFlagRequired("path")

	return addAssetsCmd
}
