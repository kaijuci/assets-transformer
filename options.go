package transformer

type AssetFormat string

const (
	PNG  AssetFormat = "png"
	JPEG AssetFormat = "jpeg"
	WEBP AssetFormat = "webp"
	GIF  AssetFormat = "gif"
)

type AssetSize struct {
	Width  int
	Height int
}

type TransformOptions struct {
	Size   AssetSize
	Format AssetFormat
}
