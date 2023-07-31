package catalog

import (
	"log"
	"os"

	"github.com/kaijuci/transformer/cmd/catalog/android"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xfrmr",
	Short: "Transforms image files into properly named and sized Android assets",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	versionCmd := NewVersionCmd()
	androidCmd := android.NewAndroidCmd()

	rootCmd.AddCommand(versionCmd, androidCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
