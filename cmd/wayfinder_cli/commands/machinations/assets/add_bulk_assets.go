package assets

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/wayfinder/internal/core"
)

var (
	coreInstance *core.Core
)

func AddAssetsCmd(c *core.Core) *cobra.Command {
	coreInstance = c

	addAssetsCmd := &cobra.Command{
		Use:   "load",
		Short: "Load Assets from a file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add assets called")
		},
	}

	return addAssetsCmd
}
