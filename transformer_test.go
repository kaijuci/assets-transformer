package transformer

import (
	"fmt"
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
	faceJPEG string = "testdata/face.jpeg"
	faceGIF  string = "testdata/face.gif"
	faceWEBP string = "testdata/face.webp"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestPNGtoPNGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.png")
	result, err := at.Transform(facePNG, &TransformOption{Format: PNG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
	if err != nil {
		t.Fatalf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Fatalf("result is empty")
	}

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

	t.Logf("Result: %s size: %dx%d format: %s", *result, size.Width, size.Height, *format)

}

func TestPNGtoJPEGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.jpeg")
	result, err := at.Transform(facePNG, &TransformOption{Format: JPEG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
	if err != nil {
		t.Fatalf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Fatalf("result is empty")
	}

	size, format, err := getImageSizeAndFormat(*result)
	if err != nil {
		t.Fatalf("error getting image size and format: %v", err)
	}

	if size.Width != 100 || size.Height != 100 {
		t.Fatalf("expected size 100x100, got %v", *size)
	}

	if *format != JPEG {
		t.Fatalf("expected format png, got %s", *format)
	}

	t.Logf("Result: %s size: %dx%d format: %s", *result, size.Width, size.Height, *format)
}

func TestPNGtoWEBPConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.webp")
	result, err := at.Transform(facePNG, &TransformOption{Format: WEBP, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
	if err != nil {
		t.Fatalf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Fatalf("result is empty")
	}

	size, format, err := getImageSizeAndFormat(*result)
	if err != nil {
		t.Fatalf("error getting image size and format: %v", err)
	}

	if size.Width != 100 || size.Height != 100 {
		t.Fatalf("expected size 100x100, got %v", *size)
	}

	if *format != WEBP {
		t.Fatalf("expected format png, got %s", *format)
	}

	t.Logf("Result: %s size: %dx%d format: %s", *result, size.Width, size.Height, *format)
}

func TestPNGtoGIFConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.gif")
	result, err := at.Transform(facePNG, &TransformOption{Format: GIF, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
	if err != nil {
		t.Fatalf("error transforming asset: %v", err)
	}

	if len(*result) == 0 {
		t.Fatalf("result is empty")
	}

	size, format, err := getImageSizeAndFormat(*result)
	if err != nil {
		t.Fatalf("error getting image size and format: %v", err)
	}

	if size.Width != 100 || size.Height != 100 {
		t.Fatalf("expected size 100x100, got %v", *size)
	}

	if *format != GIF {
		t.Fatalf("expected format png, got %s", *format)
	}

	t.Logf("Result: %s size: %dx%d format: %s", *result, size.Width, size.Height, *format)
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
