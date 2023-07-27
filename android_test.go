package transformer

import (
	"testing"
)

const rollerSkatePNG = "testdata/roller-skate.png"

func TestLauncherIcon(t *testing.T) {
	at, err := NewAndroidAssetTransformer()
	if err != nil {
		t.Fatalf("failed to create AndroidAssetTransformer: %v", err)
	}

	gen, err := at.TransformAsset(rollerSkatePNG, "ic_launcher")
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
