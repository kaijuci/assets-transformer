package transformer

import (
	"errors"

	"github.com/kaijuci/assets-transformer/android"
)

type AndroidTransformOption struct {
	TransformOption
	IconType  android.AndroidIconType
	TargetDPI []android.AssetDPI
}

type AndroidAssetTransformer interface {
	TransformAsset(string, string, ...*AndroidTransformOption) ([]string, error)
}

func NewAndroidAssetTransformer() (AndroidAssetTransformer, error) {
	return &androidimpl{}, nil
}

type androidimpl struct {
}

func (i *androidimpl) TransformAsset(filename string, name string, options ...*AndroidTransformOption) ([]string, error) {
	return nil, errors.New("not implemented")
}
