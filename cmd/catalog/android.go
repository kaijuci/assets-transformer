package catalog

import (
	"log"

	"github.com/spf13/cobra"
)

func NewAndroidCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "android",
		Short: "Android asset transformer",
		Long:  "Android asset transformer",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("do the dew")
		},
	}
}
