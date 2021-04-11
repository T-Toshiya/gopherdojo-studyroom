package converter

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func Convert(directory, filepath string) error {
	file, err := os.Open(directory + "/" + filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	paths := strings.Split(filepath, ".")
	out, err := os.Create(directory + "/" + paths[0] + ".png")
	if err != nil {
		return err
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		log.Fatal(err)
	}
	return nil
}
