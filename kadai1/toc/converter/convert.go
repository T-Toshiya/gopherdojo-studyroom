package converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type ConvertOpt struct {
	BeforeFmt string
	AfterFmt  string
}

func (c ConvertOpt) Convert(directory, filepath string) error {
	var img image.Image

	file, err := os.Open(directory + "/" + filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	switch c.BeforeFmt {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, _, err = image.Decode(file)
	}

	if err != nil {
		return err
	}

	paths := strings.Split(filepath, ".")
	out, err := os.Create(directory + "/" + paths[0] + "." + c.AfterFmt)
	if err != nil {
		return err
	}
	defer out.Close()

	switch c.AfterFmt {
	case "jpg", "jpeg":
		if err := jpeg.Encode(out, img, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	}

	return nil
}
