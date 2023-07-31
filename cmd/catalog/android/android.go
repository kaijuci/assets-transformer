package android

import (
	"errors"
	"strings"

	transformer "github.com/kaijuci/transformer"
	"github.com/kaijuci/transformer/android"
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

type androidAssetWorkFn func(inputFilename, outputDir, assetName, format string) ([]*string, error)

type androidWork struct {
	mapFn         map[string]androidAssetWorkFn
	inputFilename string
	outputDir     string
	assetName     string
	types         []string
	format        string
}

func newAndroidAssetWork(inputFilename, outputDir, assetName, types, format string) *androidWork {
	return &androidWork{
		mapFn: map[string]androidAssetWorkFn{
			"launcher": handleLauncherAssetType,
			"all":      handleAllAssetTypes,
		},
		inputFilename: inputFilename,
		outputDir:     outputDir,
		assetName:     assetName,
		types:         strings.Split(types, ","),
		format:        format,
	}
}

func (w *androidWork) doWork() ([]*string, error) {
	var result []*string
	for _, t := range w.types {
		fn, ok := w.mapFn[t]
		if !ok {
			return nil, errors.New(ErrInvalidAssetType)
		}
		paths, err := fn(w.inputFilename, w.outputDir, w.assetName, w.format)
		if err != nil {
			return nil, err
		}

		result = append(result, paths...)
	}

	return result, nil
}

func handleLauncherAssetType(inputFilename, outputDir, assetName, format string) ([]*string, error) {
	at, err := transformer.NewAndroidAssetTransformer(outputDir)
	if err != nil {
		return nil, err
	}

	opt := transformer.AndroidTransformOption{
		IconType: android.AndroidIconTypeLauncher,
		Format:   transformer.AssetFormat(strings.ToLower(format)),
	}

	gen, err := at.TransformAsset(inputFilename, assetName, &opt)
	if err != nil {
		return nil, err
	}

	result := []*string{}
	for _, v := range gen {
		path := v
		result = append(result, &path)
	}

	return result, nil
}

func handleAllAssetTypes(inputFilename, outputDir, assetName, format string) ([]*string, error) {
	return nil, errors.New("not implemented")
}
