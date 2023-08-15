package android

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kaijuci/transformer"
	"github.com/spf13/cobra"
)

const (
	_androidDescription string = `
	Transforms input image file into properly named and sized Android assets. It uses the input file format for the generated asset files.
	Optionally, the target format can be specified with the --format flag. Available formats are PNG, GIF, JPEG, and WEBP.

	Supported android icon types:
		- launcher
		- action
		- tab
		- menu
		- dialog
		- statusbar
		- toolbar
		- notification

	Examples:
	- xfrmr android -i /tmp/LauncherIcon.png -n app_icon -o /tmp/iconJob/res -t launcher -f WEBP

	  result:
	        /tmp/iconJob/res/drawable-ldpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-mdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-hdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xhdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxhdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxxhdpi/ic_launcher_app_icon.webp
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
			types := cmd.Flag("types").Value.String()
			format := cmd.Flag("format").Value.String()

			iconTypes := resolveIconTypes(strings.Split(types, ","))
			w := newAndroidAssetWork(&AndroidTransformOpts{inputFilename, outputDir, assetName, iconTypes, transformer.AssetFormat(format)})
			result, err := w.doWork()
			if err != nil {
				log.Panicf("error: %v\n", err)
			}

			b, err := json.Marshal(*result)
			if err != nil {
				log.Panicf("error: %v\n", err)
			}
			_, err = os.Stdout.WriteString(fmt.Sprintf("%s\n", string(b)))

			if err != nil {
				log.Panicf("error: %v\n", err)
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
