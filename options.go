package transformer

type AssetFormat string

const (
	PNG  AssetFormat = "png"
	JPEG AssetFormat = "jpeg"
	WEBP AssetFormat = "webp"
	GIF  AssetFormat = "gif"
)

type TransformOptions struct {
	Size   string
	Format AssetFormat
}
