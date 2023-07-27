package android

import (
	"image"
	"os"
	"testing"
)

const rollerSkatePNG = "testdata/roller-skate.png"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLauncherIcon(t *testing.T) {
	at, err := NewAndroidAssetTransformer()
	if err != nil {
		t.Fatalf("failed to create AndroidAssetTransformer: %v", err)
	}

	gen, err := at.TransformAsset(rollerSkatePNG, "ic_launcher", &AndroidTransformOption{AndroidIconTypeLauncher, []AssetDPI{LDPI, MDPI, HDPI, XHDPI, XXHDPI, XXXHDPI}})
	if err != nil {
		t.Fatalf("failed to transform asset: %v", err)
	}

	if len(gen) != 6 {
		t.Fatalf("expected 6 assets, got %d", len(gen))
	}

	for _, g := range gen {
		t.Logf("generated: %s", g)
	}
}

func getImageSizeAndFormat(filename string) (*AssetSize, *AssetFormat, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	im, format, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, nil, err
	}

	size := AssetSize{uint(im.Width), uint(im.Height)}
	assetFormat := AssetFormat(format)

	return &size, &assetFormat, nil
}
