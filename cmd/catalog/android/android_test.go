package android

import (
	"encoding/json"
	"testing"

	"github.com/kaijuci/transformer"
	"github.com/kaijuci/transformer/android"
)

func TestReportGeneration(t *testing.T) {
	opts := &AndroidTransformOpts{
		InputFilename: "testdata/face.png",
		OutputDir:     "/tmp/reportTest",
		AssetName:     "face",
		IconTypes: []android.AndroidIconType{
			android.AndroidIconTypeAction,
			android.AndroidIconTypeLauncher,
			android.AndroidIconTypeStatusBar,
			android.AndroidIconTypeNotification,
		},
		Format: transformer.WEBP,
	}
	w := newAndroidAssetWork(opts)
	report, err := w.doWork()
	if err != nil {
		t.Fatalf("report generation error: %v", err)
	}

	b, _ := json.MarshalIndent(report, "", "  ")
	t.Logf("report json: %s", string(b))
}
