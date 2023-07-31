package android

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func NewAndroidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "android",
		Short: "Android asset transformer",
		Long:  _androidDescription,
		Run: func(cmd *cobra.Command, args []string) {
			inputFilename := cmd.Flag("input").Value.String()
			outputDir := cmd.Flag("output").Value.String()
			assetName := cmd.Flag("name").Value.String()
			types := cmd.Flag("types").Value.String()

			w := newAndroidAssetWork(inputFilename, outputDir, assetName, types)
			destDir, err := w.doWork()
			if err != nil {
				log.Panicf("error: %v\n", err)
			}

			for _, d := range destDir {
				_, err = os.Stdout.WriteString(fmt.Sprintf("%s\n", *d))

				if err != nil {
					log.Panicf("error: %v\n", err)
				}
			}
		},
	}

	cmd.Flags().StringP("types", "t", "all", "Asset types to transform, comma separated. Available: launcher, notification, dialog, tab, toolbar, all")
	cmd.Flags().StringP("input", "i", "", "Input image file, can be any format one of PNG, GIF, JPEG, or WEBP")
	cmd.Flags().StringP("format", "f", "WEBP", "Output format for transformed assets. Available: PNG, GIF, JPEG, WEBP")
	cmd.Flags().StringP("output", "o", "/tmp", "Output directory for transformed assets")
	cmd.Flags().StringP("name", "n", "app_icon", "Name for transformed assets e.g. app_icon, notification_icon, etc.")

	cmd.MarkFlagRequired("input")

	return cmd
}
