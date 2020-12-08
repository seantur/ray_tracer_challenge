package canvas

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	PPMHeader  = "P3"
	PPMCharLen = 70
)

const (
	ErrOutOfBounds = "trying to access out of bounds"
)

func InitCanvas(height, width int) *image.RGBA64 {
	return image.NewRGBA64(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})
}

// WritePixel -> Set
// At -> At

func scale255(color, alpha uint32) string {
	var val uint8

	if alpha != 0 {
		val = uint8(float64(color) / float64(alpha) * 255)
	}

	if val < 0 {
		return "0 "
	} else if val > 255 {
		return "255 "
	} else {
		return strconv.Itoa(int(val)) + " "
	}
}

func addColor(color, alpha uint32, s *strings.Builder, row *string) {
	colorStr := scale255(color, alpha)

	if len([]rune(*row))+len([]rune(colorStr)) > PPMCharLen {
		writeRow(s, row)
		*row = colorStr
	} else {
		*row += colorStr
	}

}

func writeRow(s *strings.Builder, row *string) {
	s.WriteString((*row)[:len((*row))-1] + "\n")
}

func toPPM(c image.Image) string {
	var str strings.Builder

	max := c.Bounds().Max
	width, height := max.X, max.Y

	str.WriteString(fmt.Sprintf("%s\n%d %d\n255\n", PPMHeader, width, height))

	for i := 0; i < height; i++ {
		row := ""
		for j := 0; j < width; j++ {
			color := c.At(j, i)

			red, green, blue, alpha := color.RGBA()

			addColor(red, alpha, &str, &row)
			addColor(green, alpha, &str, &row)
			addColor(blue, alpha, &str, &row)

		}
		writeRow(&str, &row)
	}

	return str.String()
}

func SavePPM(c image.Image, path string) {
	err := ioutil.WriteFile(path, []byte(toPPM(c)), 0644)

	if err != nil {
		log.Fatal(err)
	}

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
