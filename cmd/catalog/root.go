package catalog

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xfrmr",
	Short: "Transforms image files into properly named and sized Android assets",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	// Add commands here
	versionCmd := NewVersionCmd()

	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
