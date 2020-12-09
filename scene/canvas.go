package scene

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func InitCanvas(height, width int) *image.RGBA64 {
	return image.NewRGBA64(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})
}

func SavePng(c image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, c); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func SaveJpg(c image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := jpeg.Encode(f, c, nil); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
