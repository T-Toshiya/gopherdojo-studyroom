package converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Converter is a struct that has formats, directory, filepath
type Converter struct {
	BeforeFmt string
	AfterFmt  string
	Directory string
	FilePath  string
}

// Convert convert image
func (c Converter) Convert() error {
	var img image.Image

	file, err := os.Open(c.Directory + "/" + c.FilePath)
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

	paths := strings.Split(c.FilePath, ".")
	out, err := os.Create(c.Directory + "/" + paths[0] + "." + c.AfterFmt)
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
