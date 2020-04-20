package canvas

import (
	"errors"
	"fmt"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"io/ioutil"
	"math"
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
	pixels []raytracing.Color
}

func (c *Canvas) Init() {
	c.pixels = make([]raytracing.Color, c.Height*c.Width)
}

func (c *Canvas) WritePixel(x int, y int, color raytracing.Color) error {
	if x > c.Width || y > c.Height {
		return errors.New(ErrOutOfBounds)
	}

	c.pixels[x*c.Height+y] = color

	return nil
}

func (c *Canvas) At(x, y int) raytracing.Color {
	return c.pixels[x*c.Height+y]
}

func (c *Canvas) ReadPixel(x int, y int) (raytracing.Color, error) {
	if x > c.Width || y > c.Height {
		return raytracing.Color{}, errors.New(ErrOutOfBounds)
	}
	return c.pixels[x*c.Height+y], nil
}

func scale255(val float64) string {
	output := int(math.Round(val * 255))

	if output < 0 {
		return "0 "
	} else if output > 255 {
		return "255 "
	} else {
		return strconv.Itoa(output) + " "
	}
}

func addColor(color float64, s *strings.Builder, row *string) {
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

			addColor(color.Red, &str, &row)
			addColor(color.Green, &str, &row)
			addColor(color.Blue, &str, &row)

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
