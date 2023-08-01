package android

import (
	"strings"

	"github.com/kaijuci/transformer"
	"github.com/kaijuci/transformer/android"
)

const (
	ErrInvalidAssetType = "invalid asset type"
)

type androidWork struct {
	inputFilename string
	outputDir     string
	assetName     string
	iconTypes     []android.AndroidIconType
	format        transformer.AssetFormat
}

func newAndroidAssetWork(inputFilename, outputDir, assetName, types, format string) *androidWork {
	return &androidWork{
		inputFilename: inputFilename,
		outputDir:     outputDir,
		assetName:     assetName,
		iconTypes:     resolveIconTypes(strings.Split(types, ",")),
		format:        transformer.AssetFormat(strings.ToLower(format)),
	}
}

func (w *androidWork) doWork() ([]*string, error) {
	var result []*string
	for _, t := range w.iconTypes {
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
	at, err := transformer.NewAndroidAssetTransformer(w.outputDir)
	if err != nil {
		return nil, err
	}

	result := []*string{}

	opt := transformer.AndroidTransformOption{
		IconType: iconType,
		Format:   w.format,
	}

	gen, err := at.TransformAsset(w.inputFilename, w.assetName, &opt)
	if err != nil {
		return nil, err
	}

	for _, v := range gen {
		path := v
		result = append(result, &path)
	}

	return result, nil
}
