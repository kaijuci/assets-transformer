package transformer

import (
	"fmt"
	"testing"
	"time"

	"github.com/kaijuci/assets-transformer/android"
)

const rollerSkatePNG = "testdata/roller-skate.png"

func TestLauncherIcon(t *testing.T) {
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

	for _, g := range gen {
		t.Logf("generated: %s", g)
	}
}
