package transformer

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kaijuci/assets-transformer/android"
)

type AndroidTransformOption struct {
	IconType android.AndroidIconType
	Format   AssetFormat
}

type AndroidAssetTransformer interface {
	TransformAsset(string, string, ...*AndroidTransformOption) ([]string, error)
}

func NewAndroidAssetTransformer(workDir string) (AndroidAssetTransformer, error) {
	return &androidimpl{
		workDir:      workDir,
		dpiList:      android.NewDPIList(),
		dpiDict:      android.NewAndroidAssetDensityDictionary(),
		iconSpecDict: android.NewAndroidIconSpecDictionary(),
	}, nil
}

type androidimpl struct {
	workDir      string
	dpiList      []android.AssetDPI
	dpiDict      android.AndroidAssetDensityDictionary
	iconSpecDict android.AndroidIconSpecDictionary
}

func (i *androidimpl) TransformAsset(filename string, name string, options ...*AndroidTransformOption) ([]string, error) {
	opts, err := i.validateOptions(name, options)
	if err != nil {
		return nil, err
	}

	at, err := NewAssetTransformer()
	if err != nil {
		return nil, err
	}

	var paths []string

	log.Printf("opts: %v", opts)

	for dir, opt := range opts {
		err = i.ensureRootDir(dir)
		if err != nil {
			return nil, err
		}

		path, err := at.Transform(filename, opt)
		if err != nil {
			return nil, err
		}
		paths = append(paths, *path)
	}

	return paths, nil
}

func (i *androidimpl) ensureRootDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *androidimpl) validateOptions(name string, options []*AndroidTransformOption) (map[string]*TransformOption, error) {
	if len(options) == 0 {
		return nil, errors.New("no options")
	}

	src := options[0]
	spec := i.iconSpecDict[src.IconType]
	var opts map[string]*TransformOption = make(map[string]*TransformOption)

	for _, dpi := range i.dpiList {
		density := i.dpiDict[dpi]
		path := fmt.Sprintf("%s/%s", i.workDir, density.ResPath)
		filename := fmt.Sprintf("%s/%s/%s_%s.%s", i.workDir, density.ResPath, spec.Prefix, name, src.Format)
		opt := &TransformOption{
			Size:    AssetSize{density.Width, density.Height},
			Format:  src.Format,
			Outfile: filename,
		}
		log.Printf("opt: %v", opt)
		opts[path] = opt
	}

	return opts, nil
}
