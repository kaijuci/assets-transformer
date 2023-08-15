package android

import (
	"strings"

	"github.com/kaijuci/transformer"
	"github.com/kaijuci/transformer/android"
)

const (
	ErrInvalidAssetType = "invalid asset type"
)

type AndroidTransformOpts struct {
	InputFilename string
	OutputDir     string
	AssetName     string
	IconTypes     []android.AndroidIconType
	Format        transformer.AssetFormat
}
type AndroidTransformResult struct {
	InputFilename string `json:"input_filename"`
}
type androidWork struct {
	AndroidTransformOpts
}

func newAndroidAssetWork(inputFilename, outputDir, assetName, types, format string) *androidWork {
	return &androidWork{
		AndroidTransformOpts{InputFilename: inputFilename,
			OutputDir: outputDir,
			AssetName: assetName,
			IconTypes: resolveIconTypes(strings.Split(types, ",")),
			Format:    transformer.AssetFormat(strings.ToLower(format)),
		},
	}
}

func (w *androidWork) doWork() ([]*string, error) {
	var result []*string
	for _, t := range w.IconTypes {
		paths, err := w.transform(t)
		if err != nil {
			return nil, err
		}

		result = append(result, paths...)
	}

	return result, nil
}

func resolveIconTypes(types []string) []android.AndroidIconType {
	var result []android.AndroidIconType
	for _, t := range types {
		i := android.AndroidIconType(t)
		result = append(result, i)
	}

	return result
}

func (w *androidWork) transform(iconType android.AndroidIconType) ([]*string, error) {
	at, err := transformer.NewAndroidAssetTransformer(w.OutputDir)
	if err != nil {
		return nil, err
	}

	result := []*string{}

	opt := transformer.AndroidTransformOption{
		IconType: iconType,
		Format:   w.Format,
	}

	gen, err := at.TransformAsset(w.InputFilename, w.AssetName, &opt)
	if err != nil {
		return nil, err
	}

	for _, v := range gen {
		path := v
		result = append(result, &path)
	}

	return result, nil
}
