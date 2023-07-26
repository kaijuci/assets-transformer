package transformer

import (
	"errors"
	"log"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type AssetTransformer interface {
	Transform(string, ...*TransformOption) (*string, error)
}

func NewAssetTransformer() (AssetTransformer, error) {
	return &impl{}, nil
}

type impl struct{}

func (i *impl) Transform(asset string, options ...*TransformOption) (*string, error) {
	opt, err := i.validateOptions(options)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(asset)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imagick.Initialize()
	defer imagick.Terminate()

	wand := imagick.NewMagickWand()

	err = wand.PingImage(asset)
	if err != nil {
		return nil, err
	}
	log.Printf("original name: %s size:%dx%d format:%s", asset, wand.GetImageWidth(), wand.GetImageHeight(), wand.GetImageFormat())

	err = wand.ReadImageFile(file)
	if err != nil {
		return nil, err
	}

	// wand.SetImageInterpolateMethod(imagick.INTERPOLATE_PIXEL_BICUBIC)
	err = wand.ResizeImage(opt.Size.Width, opt.Size.Height, imagick.FILTER_LANCZOS2_SHARP)
	if err != nil {
		return nil, err
	}

	err = wand.SetImageFormat(string(opt.Format))
	if err != nil {
		return nil, err
	}

	err = wand.WriteImage(opt.Outfile)
	if err != nil {
		return nil, err
	}

	return &opt.Outfile, nil
}

func (i *impl) validateOptions(options []*TransformOption) (*TransformOption, error) {
	if len(options) == 0 {
		return nil, errors.New("no options provided for operations")
	}

	opt := options[0]
	if len(opt.Outfile) == 0 {
		return nil, errors.New("outfile not specified")
	}

	return opt, nil
}
