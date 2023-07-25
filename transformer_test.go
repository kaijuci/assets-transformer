package transformer

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

const (
	facePNG  string = "testdata/face.png"
	touchPNG string = "testdata/touch.png"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestPNGConverstion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	result, err := at.Transform("test.png", &TransformOptions{string(PNG), "100x100"})
	if err != nil {
		t.Errorf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Errorf("result is empty")
	}

	t.Logf("Result: %s", *result)
	size, format, err := getImageSizeAndFormat(*result)
	if err != nil {
		t.Errorf("error getting image size and format: %v", err)
	}

	if *size != "100x100" {
		t.Errorf("expected size 100x100, got %s", *size)
	}

	if *format != "png" {
		t.Errorf("expected format png, got %s", *format)
	}
}

func TestJPEGConverstion(t *testing.T) {
	t.Errorf("Test not implemented")
}

func TestWEBPConverstion(t *testing.T) {
	t.Errorf("Test not implemented")
}

func getImageSizeAndFormat(filename string) (*string, *string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	im, format, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, nil, err
	}

	size := fmt.Sprintf("%dx%d", im.Width, im.Height)

	return &size, &format, nil
}
