package transformer

type AssetFormat string

const (
	PNG  AssetFormat = "png"
	JPEG AssetFormat = "jpeg"
	WEBP AssetFormat = "webp"
	GIF  AssetFormat = "gif"
)

type AssetSize struct {
	Width  uint
	Height uint
}

type TransformOption struct {
	Size    AssetSize
	Format  AssetFormat
	Outfile string
}
