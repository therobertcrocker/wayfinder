package assets

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/internal/core"
)

var (
	coreInstance *core.Core
	filePath     string
	fileType     string
)

func BulkAddAssetsCmd(c *core.Core) *cobra.Command {
	coreInstance = c

	bulkAddAssetsCmd := &cobra.Command{
		Use:   "bulk-add",
		Short: "Bulk add assets",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			coreInstance.Machinations.AssetStore.BulkAddAssets(filePath, fileType)
		},
	}

	bulkAddAssetsCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to file containing assets to add")
	bulkAddAssetsCmd.Flags().StringVarP(&fileType, "type", "t", "", "Type of file containing assets to add")

	bulkAddAssetsCmd.MarkFlagRequired("file")

	return bulkAddAssetsCmd
}
