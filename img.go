package goutil

import (
	"image"
	"os"
	"path"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

// JPG only
func RotateImgOrientation(filename string) error {

	if strings.ToLower(path.Ext(filename)) != "jpg" {
		return nil
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	exif.RegisterParsers(mknote.All...)
	exf, err := exif.Decode(file)
	if err != nil {
		return err
	}

	tag, err := exf.Get(exif.Orientation)
	if err != nil {
		return err
	}

	var i int
	i, err = tag.Int(i)
	if err != nil {
		return err
	}

	r := map[int]func(image.Image) *image.NRGBA{
		6: imaging.Rotate270,
		3: imaging.Rotate180,
		8: imaging.Rotate90,
	}

	rotate, ok := r[i]
	if !ok {
		return nil
	}

	img, err := imaging.Decode(file)
	if err != nil {
		return err
	}

	img = rotate(img)
	return imaging.Save(img, filename)
}
