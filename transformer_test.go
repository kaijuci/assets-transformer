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

func TestJPEGtoPNGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.png")
	result, err := at.Transform(faceJPEG, &TransformOption{Format: PNG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestJPEGtoJPEGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.jpeg")
	result, err := at.Transform(faceJPEG, &TransformOption{Format: JPEG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestJPEGtoWEBPConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.webp")
	result, err := at.Transform(faceJPEG, &TransformOption{Format: WEBP, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestJPEGtoGIFConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.gif")
	result, err := at.Transform(faceJPEG, &TransformOption{Format: GIF, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestGIFtoPNGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.png")
	result, err := at.Transform(faceGIF, &TransformOption{Format: PNG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestGIFtoJPEGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.jpeg")
	result, err := at.Transform(faceGIF, &TransformOption{Format: JPEG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestGIFtoWEBPConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.webp")
	result, err := at.Transform(faceGIF, &TransformOption{Format: WEBP, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestGIFtoGIFConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.gif")
	result, err := at.Transform(faceGIF, &TransformOption{Format: GIF, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestWEBPtoPNGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.png")
	result, err := at.Transform(faceWEBP, &TransformOption{Format: PNG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestWEBPtoJPEGConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.jpeg")
	result, err := at.Transform(faceWEBP, &TransformOption{Format: JPEG, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestWEBPtoWEBPConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.webp")
	result, err := at.Transform(faceWEBP, &TransformOption{Format: WEBP, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestWEBPtoGIFConversion(t *testing.T) {
	at, err := NewAssetTransformer()
	if err != nil {
		t.Errorf("error creating transformer: %v", err)
	}

	outfile := fmt.Sprintf("/tmp/%s", "face.gif")
	result, err := at.Transform(faceWEBP, &TransformOption{Format: GIF, Size: AssetSize{Width: 100, Height: 100}, Outfile: outfile})
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

func TestImageInfo(t *testing.T) {
	want := AssetInfo{Width: 512, Height: 512, Format: "PNG"}

	got, err := GetImageInfo(facePNG)
	if err != nil {
		t.Fatalf("error getting image info: %v", err)
	}

	if got.Width != want.Width || got.Height != want.Height || got.Format != want.Format {
		t.Fatalf("incorrect asset info got: %v want: %v", got, want)
	}

	t.Logf("asset: %s info: %v", facePNG, got)
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
