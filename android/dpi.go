package android

import "fmt"

type AssetDPI string

const (
	LDPI    AssetDPI = "ldpi"
	MDPI    AssetDPI = "mdpi"
	HDPI    AssetDPI = "hdpi"
	XHDPI   AssetDPI = "xhdpi"
	XXHDPI  AssetDPI = "xxhdpi"
	XXXHDPI AssetDPI = "xxxhdpi"
)

type AssetResPath string

const (
	DrawableLDPI    AssetResPath = "drawable-ldpi"
	DrawableMDPI    AssetResPath = "drawable-mdpi"
	DrawableHDPI    AssetResPath = "drawable-hdpi"
	DrawableXHDPI   AssetResPath = "drawable-xhdpi"
	DrawableXXHDPI  AssetResPath = "drawable-xxhdpi"
	DrawableXXXHDPI AssetResPath = "drawable-xxxhdpi"
)

type AndroidAssetDensity struct {
	DPI     AssetDPI     `json:"dpi"`
	ResPath AssetResPath `json:"res_path"`
	Width   uint         `json:"width"`
	Height  uint         `json:"height"`
}

func (a AndroidAssetDensity) MakeAssetPath(rootPath, asset string) string {
	return fmt.Sprintf("%s/%s/%s", rootPath, a.ResPath, asset)
}

type AndroidAssetDensityDictionary map[AssetDPI]AndroidAssetDensity

func NewAndroidAssetDensityDictionary() AndroidAssetDensityDictionary {
	return AndroidAssetDensityDictionary{
		LDPI:    AndroidAssetDensity{DPI: LDPI, ResPath: DrawableLDPI, Width: 36, Height: 36},
		MDPI:    AndroidAssetDensity{DPI: MDPI, ResPath: DrawableMDPI, Width: 48, Height: 48},
		HDPI:    AndroidAssetDensity{DPI: HDPI, ResPath: DrawableHDPI, Width: 72, Height: 72},
		XHDPI:   AndroidAssetDensity{DPI: XHDPI, ResPath: DrawableXHDPI, Width: 96, Height: 96},
		XXHDPI:  AndroidAssetDensity{DPI: XXHDPI, ResPath: DrawableXXHDPI, Width: 144, Height: 144},
		XXXHDPI: AndroidAssetDensity{DPI: XXXHDPI, ResPath: DrawableXXXHDPI, Width: 192, Height: 192},
	}
}
