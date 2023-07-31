package catalog

import (
	"log"

	"github.com/kaijuci/assets-transformer/vars"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of xfrmr",
		Long:  `All software has versions. This is xfrmr's`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("version: %s\n", vars.Version)
		},
	}
}
