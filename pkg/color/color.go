package color

import (
	"image"
	"image/png"
	"os"
)

func LoadPaletteFromImage(p string) (image.Image, error) {
	ImageFile, err := os.Open("../../files/palettes/" + p + ".png")
	if err != nil {
		return nil, err
	}
	defer ImageFile.Close()

	loadedImage, err := png.Decode(ImageFile)
	if err != nil {
		return nil, err
	}

	return loadedImage, nil
}
