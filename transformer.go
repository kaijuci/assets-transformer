package transformer

import "errors"

type AssetTransformer interface {
	Transform(string, ...*TransformOptions) (*string, error)
}

func NewAssetTransformer() (AssetTransformer, error) {
	return &impl{}, nil
}

type impl struct{}

func (i *impl) Transform(asset string, options ...*TransformOptions) (*string, error) {
	return nil, errors.New("not implemented")
}
