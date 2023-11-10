package assets

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
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
		Short: "Bulk add assets from a file",
		Long:  `This command will add assets from a file to the database.`,
		Run: func(cmd *cobra.Command, args []string) {

			logging.Log.Info("Bulk adding assets")

			if err := coreInstance.Machinations.AssetStore.BulkAddAssets(filePath, fileType); err != nil {
				logging.Log.Fatalf("failed to bulk add assets: %v", err)
			}

			logging.Log.Info("Bulk adding assets complete")
		},
	}

	bulkAddAssetsCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to file containing assets to add")
	bulkAddAssetsCmd.Flags().StringVarP(&fileType, "type", "t", "", "Type of file containing assets to add")

	bulkAddAssetsCmd.MarkFlagRequired("file")

	return bulkAddAssetsCmd
}
