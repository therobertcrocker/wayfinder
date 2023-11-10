package machinations

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/cmd/wayfinder_cli/commands/machinations/assets"
	"github.com/therobertcrocker/wayfinder/internal/core"
)

var (
	coreInstance *core.Core
)

func AssetsCmd(c *core.Core) *cobra.Command {
	coreInstance = c

	assetsCmd := &cobra.Command{
		Use:   "assets",
		Short: "Manage Assets",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("assets called")
		},
	}

	bulkAddCmd := assets.BulkAddAssetsCmd(coreInstance)

	assetsCmd.AddCommand(bulkAddCmd)

	return assetsCmd
}
