package android

import (
	"github.com/kaijuci/transformer"
	"github.com/kaijuci/transformer/android"
)

const (
	ErrInvalidAssetType = "invalid asset type"
)

type AndroidTransformOpts struct {
	InputFilename string                    `json:"filename"`
	OutputDir     string                    `json:"res_dir"`
	AssetName     string                    `json:"asset_name"`
	IconTypes     []android.AndroidIconType `json:"icon_types"`
	Format        transformer.AssetFormat   `json:"-"`
}

type AndroidTransformResult struct {
	AndroidTransformOpts `json:"input"`
	AssetInfo            *transformer.AssetInfo                `json:"asset_info"`
	OutputFormat         string                                `json:"output_format"`
	GeneratedPaths       map[android.AndroidIconType][]*string `json:"generated"`
}

type androidWork struct {
	AndroidTransformOpts
}

func newAndroidAssetWork(opts *AndroidTransformOpts) *androidWork {
	return &androidWork{*opts}
}

func (w *androidWork) doWork() (*AndroidTransformResult, error) {
	info, err := transformer.GetImageInfo(w.InputFilename)
	if err != nil {
		return nil, err
	}

	generated := map[android.AndroidIconType][]*string{}
	for _, iconType := range w.IconTypes {
		paths, err := w.transform(iconType)
		if err != nil {
			return nil, err
		}
		generated[iconType] = paths
	}

	return &AndroidTransformResult{w.AndroidTransformOpts, info, string(w.AndroidTransformOpts.Format), generated}, nil
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
