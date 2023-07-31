package catalog

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	_androidDescription string = `
	Transforms input image file into properly named and sized Android assets. It uses the input file format for the generated asset files.
	Optionally, the target format can be specified with the --format flag. Available formats are PNG, GIF, JPEG, and WEBP.

	Examples:
	- xfrmr android -i /tmp/LauncherIcon.png -n app_icon -o /tmp/iconJob/res -t launcher -f WEBP

	  result:
	        /tmp/iconJob/res/drawable-ldpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-mdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-hdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxxdpi/ic_launcher_app_icon.webp
	`
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
			types, err := cmd.Flags().GetStringArray("types")
			if err != nil {
				log.Fatalf("error: %v\n", err)
			}
			doAndroidAssetWork(inputFilename, outputDir, assetName, types)
		},
	}

	cmd.Flags().StringArrayP("types", "t", []string{"all"}, "Asset types to transform. Available: launcher, notification, dialog, tab, toolbar, all")
	cmd.Flags().StringP("input", "i", "", "Input image file, can be any format one of PNG, GIF, JPEG, or WEBP")
	cmd.Flags().StringP("format", "f", "WEBP", "Output format for transformed assets. Available: PNG, GIF, JPEG, WEBP")
	cmd.Flags().StringP("output", "o", "/tmp", "Output directory for transformed assets")
	cmd.Flags().StringP("name", "n", "app_icon", "Name for transformed assets e.g. app_icon, notification_icon, etc.")

	cmd.MarkFlagRequired("input")

	return cmd
}

func doAndroidAssetWork(inputFilename, outputDir, assetName string, types []string) {
	log.Printf("inputFilename: %s outputDir: %s assetName: %s types: %s", inputFilename, outputDir, assetName, types)
}
