package transformer

import (
	"fmt"
	"testing"
	"time"

	"github.com/kaijuci/transformer/android"
)

const rollerSkatePNG = "testdata/roller-skate.png"

func TestLauncherIconPNG(t *testing.T) {
	workDir := fmt.Sprintf("/tmp/at-%d", time.Now().UnixMilli())
	at, err := NewAndroidAssetTransformer(workDir)
	if err != nil {
		t.Fatalf("failed to create AndroidAssetTransformer: %v", err)
	}

	opt := AndroidTransformOption{
		IconType: android.AndroidIconTypeLauncher,
		Format:   PNG,
	}
	gen, err := at.TransformAsset(rollerSkatePNG, "app", &opt)
	if err != nil {
		t.Fatalf("failed to transform asset: %v", err)
	}

	if len(gen) != 6 {
		t.Fatalf("expected 6 assets, got %d", len(gen))
	}

	dpiDict := android.NewAndroidAssetDensityDictionary()

	for dpi, g := range gen {
		file := g.GeneratedPath
		size, format, err := getImageSizeAndFormat(file)
		if err != nil {
			t.Fatalf("failed to get image size and format: %v", err)
		}

		if *format != PNG {
			t.Fatalf("expected PNG format, got %s", string(*format))
		}

		density := dpiDict[dpi]
		pixels := density.CalculateDimension(48)
		if size.Width != pixels || size.Height != pixels {
			t.Fatalf("expected %dx%d, got %dx%d", pixels, pixels, size.Width, size.Height)
		}
		t.Logf("generated: %s %dx%d %s", file, size.Width, size.Height, string(*format))
	}
}

func TestLauncherIconWEBP(t *testing.T) {
	workDir := fmt.Sprintf("/tmp/at-%d", time.Now().UnixMilli())
	at, err := NewAndroidAssetTransformer(workDir)
	if err != nil {
		t.Fatalf("failed to create AndroidAssetTransformer: %v", err)
	}

	opt := AndroidTransformOption{
		IconType: android.AndroidIconTypeLauncher,
		Format:   WEBP,
	}
	gen, err := at.TransformAsset(rollerSkatePNG, "app", &opt)
	if err != nil {
		t.Fatalf("failed to transform asset: %v", err)
	}

	if len(gen) != 6 {
		t.Fatalf("expected 6 assets, got %d", len(gen))
	}

	dpiDict := android.NewAndroidAssetDensityDictionary()

	for dpi, g := range gen {
		file := g.GeneratedPath
		size, format, err := getImageSizeAndFormat(file)
		if err != nil {
			t.Fatalf("failed to get image size and format: %v", err)
		}

		if *format != WEBP {
			t.Fatalf("expected PNG format, got %s", string(*format))
		}

		density := dpiDict[dpi]
		pixels := density.CalculateDimension(48)
		if size.Width != pixels || size.Height != pixels {
			t.Fatalf("expected %dx%d, got %dx%d", pixels, pixels, size.Width, size.Height)
		}
		t.Logf("generated: %s %dx%d %s", file, size.Width, size.Height, string(*format))
	}
}

func TestNotificationIconWEBP(t *testing.T) {
	workDir := fmt.Sprintf("/tmp/at-%d", time.Now().UnixMilli())
	at, err := NewAndroidAssetTransformer(workDir)
	if err != nil {
		t.Fatalf("failed to create AndroidAssetTransformer: %v", err)
	}

	opt := AndroidTransformOption{
		IconType: android.AndroidIconTypeNotification,
		Format:   WEBP,
	}
	gen, err := at.TransformAsset(rollerSkatePNG, "app", &opt)
	if err != nil {
		t.Fatalf("failed to transform asset: %v", err)
	}

	if len(gen) != 6 {
		t.Fatalf("expected 6 assets, got %d", len(gen))
	}

	dpiDict := android.NewAndroidAssetDensityDictionary()

	for dpi, generated := range gen {
		file := generated.GeneratedPath
		size, format, err := getImageSizeAndFormat(file)
		if err != nil {
			t.Fatalf("failed to get image size and format: %v", err)
		}

		if *format != WEBP {
			t.Fatalf("expected PNG format, got %s", string(*format))
		}

		density := dpiDict[dpi]
		pixels := density.CalculateDimension(32)
		if size.Width != pixels || size.Height != pixels {
			t.Fatalf("expected %dx%d, got %dx%d", pixels, pixels, size.Width, size.Height)
		}
		t.Logf("generated: %s %dx%d %s", file, size.Width, size.Height, string(*format))
	}
}
