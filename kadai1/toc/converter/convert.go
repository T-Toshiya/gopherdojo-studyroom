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
	case "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	}

	if err != nil {
		return err
	}

	paths := strings.Split(filepath, ".")
	out, err := os.Create(directory + "/" + paths[0] + ".png")
	if err != nil {
		return err
	}
	defer out.Close()

	switch c.AfterFmt {
	case "jpg":
		if err := jpeg.Encode(out, img, &jpeg.Options{}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	}

	return nil
}
