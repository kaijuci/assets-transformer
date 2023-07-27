package android

import "errors"

type AndroidTransformOption struct {
	IconType  AndroidIconType
	TargetDPI []AssetDPI
}

type AndroidAssetTransformer interface {
	GetDPIDictionary() AndroidAssetDensityDictionary
	GetIconSpecDictionary() AndroidIconSpecDictionary
	TransformAsset(string, string, ...*AndroidTransformOption) ([]string, error)
}

func NewAndroidAssetTransformer() (AndroidAssetTransformer, error) {
	return &impl{newAndroidAssetDensityDictionary(), newAndroidIconSpecDictionary()}, nil
}

type impl struct {
	dpiDict      AndroidAssetDensityDictionary
	iconSpecDict AndroidIconSpecDictionary
}

func (i *impl) GetDPIDictionary() AndroidAssetDensityDictionary {
	return i.dpiDict
}

func (i *impl) GetIconSpecDictionary() AndroidIconSpecDictionary {
	return i.iconSpecDict
}

func (i *impl) TransformAsset(filename string, name string, options ...*AndroidTransformOption) ([]string, error) {
	return nil, errors.New("not implemented")
}
