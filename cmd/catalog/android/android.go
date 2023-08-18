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

type AndroidDPIPath struct {
	IconType  android.AndroidIconType     `json:"icon_type"`
	Generated map[android.AssetDPI]string `json:"generated_paths"`
}

type AndroidTransformResult struct {
	AndroidTransformOpts `json:"input"`
	AssetInfo            *transformer.AssetInfo `json:"asset_info"`
	OutputFormat         string                 `json:"output_format"`
	Generated            []AndroidDPIPath       `json:"generated"`
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

	generated := []AndroidDPIPath{}
	for _, iconType := range w.IconTypes {
		dpiPaths, err := w.transform(iconType)
		if err != nil {
			return nil, err
		}
		generated = append(generated, AndroidDPIPath{iconType, dpiPaths})
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

func (w *androidWork) transform(iconType android.AndroidIconType) (map[android.AssetDPI]string, error) {
	at, err := transformer.NewAndroidAssetTransformer(w.OutputDir)
	if err != nil {
		return nil, err
	}

	opt := transformer.AndroidTransformOption{
		IconType: iconType,
		Format:   w.Format,
	}

	gen, err := at.TransformAsset(w.InputFilename, w.AssetName, &opt)
	if err != nil {
		return nil, err
	}

	return gen, nil
}
