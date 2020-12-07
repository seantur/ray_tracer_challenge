package canvas

import (
	"errors"
	"fmt"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"image/color"
	"io/ioutil"
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

type Canvas struct {
	Height int
	Width  int
	pixels []color.Color
}

func (c *Canvas) Init() {
	for i := 0; i < c.Height*c.Width; i++ {
		c.pixels = append(c.pixels, raytracing.RGB{})
	}
}

func (c *Canvas) WritePixel(x int, y int, pixel color.Color) error {
	if x > c.Width || y > c.Height {
		return errors.New(ErrOutOfBounds)
	}

	c.pixels[x*c.Height+y] = pixel

	return nil
}

func (c *Canvas) At(x, y int) color.Color {
	return c.pixels[x*c.Height+y]
}

func (c *Canvas) ReadPixel(x int, y int) (color.Color, error) {
	if x > c.Width || y > c.Height {
		return raytracing.RGB{}, errors.New(ErrOutOfBounds)
	}
	return c.pixels[x*c.Height+y], nil
}

func scale255(val uint32) string {
	if val < 0 {
		return "0 "
	} else if val > 255 {
		return "255 "
	} else {
		return strconv.Itoa(int(val)) + " "
	}
}

func addColor(color uint32, s *strings.Builder, row *string) {
	colorStr := scale255(color)

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

func (c *Canvas) toPPM() string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%s\n%d %d\n255\n", PPMHeader, c.Width, c.Height))

	for i := 0; i < c.Height; i++ {
		row := ""
		for j := 0; j < c.Width; j++ {
			color, _ := c.ReadPixel(j, i)

			red, green, blue, _ := color.RGBA()

			addColor(red, &str, &row)
			addColor(green, &str, &row)
			addColor(blue, &str, &row)

		}
		writeRow(&str, &row)
	}

	return str.String()
}

func (c *Canvas) SavePPM(path string) {
	err := ioutil.WriteFile(path, []byte(c.toPPM()), 0644)

	if err != nil {
		panic(err)
	}

}
