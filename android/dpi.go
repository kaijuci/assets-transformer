package android

import "fmt"

type AssetDPI string

const BaseDPI float32 = 160

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
	DrawableMDPI    AssetResPath = "drawable-mdpi" // baseline
	DrawableHDPI    AssetResPath = "drawable-hdpi"
	DrawableXHDPI   AssetResPath = "drawable-xhdpi"
	DrawableXXHDPI  AssetResPath = "drawable-xxhdpi"
	DrawableXXXHDPI AssetResPath = "drawable-xxxhdpi"
)

func NewDPIList() []AssetDPI {
	return []AssetDPI{LDPI, MDPI, HDPI, XHDPI, XXHDPI, XXXHDPI}
}

type AndroidAssetDensity struct {
	DPI     AssetDPI     `json:"dpi"`
	ResPath AssetResPath `json:"res_path"`
	Pixels  uint         `json:"pixels"`
}

func (a AndroidAssetDensity) MakeAssetPath(asset string) string {
	return fmt.Sprintf("%s/%s", a.ResPath, asset)
}

func (a AndroidAssetDensity) CalculateDimension(dp uint) uint {
	px := uint(float32(dp) * (float32(a.Pixels) / BaseDPI))
	return px
}

type AndroidAssetDensityDictionary map[AssetDPI]AndroidAssetDensity

func NewAndroidAssetDensityDictionary() AndroidAssetDensityDictionary {
	return AndroidAssetDensityDictionary{
		LDPI:    AndroidAssetDensity{DPI: LDPI, ResPath: DrawableLDPI, Pixels: 120},
		MDPI:    AndroidAssetDensity{DPI: MDPI, ResPath: DrawableMDPI, Pixels: 160},
		HDPI:    AndroidAssetDensity{DPI: HDPI, ResPath: DrawableHDPI, Pixels: 240},
		XHDPI:   AndroidAssetDensity{DPI: XHDPI, ResPath: DrawableXHDPI, Pixels: 320},
		XXHDPI:  AndroidAssetDensity{DPI: XXHDPI, ResPath: DrawableXXHDPI, Pixels: 480},
		XXXHDPI: AndroidAssetDensity{DPI: XXXHDPI, ResPath: DrawableXXXHDPI, Pixels: 640},
	}
}
