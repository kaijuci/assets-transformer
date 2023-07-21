package catalog

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "xfrmr",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	// Add commands here
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
