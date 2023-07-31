package android

import (
	"errors"
	"strings"
)

const (
	ErrInvalidAssetType = "invalid asset type"
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
	        /tmp/iconJob/res/drawable-xhdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxhdpi/ic_launcher_app_icon.webp
	        /tmp/iconJob/res/drawable-xxxhdpi/ic_launcher_app_icon.webp
	`
)

type androidAssetWorkFn func(inputFilename, outputDir, assetName string) (*string, error)

type androidWork struct {
	mapFn         map[string]androidAssetWorkFn
	inputFilename string
	outputDir     string
	assetName     string
	types         []string
}

func newAndroidAssetWork(inputFilename, outputDir, assetName, types string) *androidWork {
	return &androidWork{
		mapFn: map[string]androidAssetWorkFn{
			"launcher": handleLauncherAssetType,
			"all":      handleAllAssetTypes,
		},
		inputFilename: inputFilename,
		outputDir:     outputDir,
		assetName:     assetName,
		types:         strings.Split(types, ","),
	}
}

func (w *androidWork) doWork() (*string, error) {
	return nil, errors.New("not implemented")
}

func handleLauncherAssetType(inputFilename, outputDir, assetName string) (*string, error) {
	return nil, errors.New("not implemented")
}

func handleAllAssetTypes(inputFilename, outputDir, assetName string) (*string, error) {
	return nil, errors.New("not implemented")
}
