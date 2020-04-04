package canvas

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const (
	PPMHeader  = "P3"
	PPMCharLen = 70
)

const (
	ErrOutOfBounds = "trying to access out of bounds"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func (c *Color) Multiply(a float64) Color {
	return Color{c.Red * a, c.Green * a, c.Blue * a}
}

func Add(a Color, b Color) Color {
	return Color{a.Red + b.Red, a.Green + b.Green, a.Blue + b.Blue}
}

func Subtract(a Color, b Color) Color {
	return Color{a.Red - b.Red, a.Green - b.Green, a.Blue - b.Blue}
}

func Hadamard(a Color, b Color) Color {
	return Color{a.Red * b.Red, a.Green * b.Green, a.Blue * b.Blue}
}

type Canvas struct {
	Height int
	Width  int
	pixels []Color
}

func (c *Canvas) Init() {
	c.pixels = make([]Color, c.Height*c.Width)
}

func (c *Canvas) WritePixel(x int, y int, color Color) error {
	if x > c.Width || y > c.Height {
		return errors.New(ErrOutOfBounds)
	}

	c.pixels[x*c.Height+y] = color

	return nil
}

func (c *Canvas) ReadPixel(x int, y int) (Color, error) {
	if x > c.Width || y > c.Height {
		return Color{}, errors.New(ErrOutOfBounds)
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

func addColor(color float64, s *string, row *string) {
	colorStr := scale255(color)

	if len([]rune(*row))+len([]rune(colorStr)) > PPMCharLen {
		writeRow(s, row)
		*row = colorStr
	} else {
		*row += colorStr
	}

}

func writeRow(s *string, row *string) {
	*s += (*row)[:len((*row))-1] + "\n"
}

func (c *Canvas) toPPM() string {
	s := fmt.Sprintf("%s\n%d %d\n255\n", PPMHeader, c.Width, c.Height)

	for i := 0; i < c.Height; i++ {
		row := ""
		for j := 0; j < c.Width; j++ {
			color, _ := c.ReadPixel(j, i)

			addColor(color.Red, &s, &row)
			addColor(color.Green, &s, &row)
			addColor(color.Blue, &s, &row)

		}
		writeRow(&s, &row)
	}

	return s
}

func (c *Canvas) SavePPM(path string) {
	err := ioutil.WriteFile(path, []byte(c.toPPM()), 0644)

	if err != nil {
		panic(err)
	}

}
