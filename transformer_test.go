package transformer

import (
	"image"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"os"
	"testing"

	_ "golang.org/x/image/webp"
)

const (
	facePNG  string = "testdata/face.png"
	touchPNG string = "testdata/touch.png"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestPNGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	result, err := at.Transform(facePNG, &TransformOptions{Format: PNG, Size: AssetSize{Width: 100, Height: 100}})
	if err != nil {
		t.Fatalf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Fatalf("result is empty")
	}

	t.Logf("Result: %s", *result)

	size, format, err := getImageSizeAndFormat(*result)
	if err != nil {
		t.Fatalf("error getting image size and format: %v", err)
	}

	if size.Width != 100 || size.Height != 100 {
		t.Fatalf("expected size 100x100, got %v", *size)
	}

	if *format != PNG {
		t.Fatalf("expected format png, got %s", *format)
	}
}

func TestJPEGConverstion(t *testing.T) {
	t.Errorf("Test not implemented")
}

func TestWEBPConverstion(t *testing.T) {
	t.Errorf("Test not implemented")
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

	size := AssetSize{im.Width, im.Height}
	assetFormat := AssetFormat(format)

	return &size, &assetFormat, nil
}