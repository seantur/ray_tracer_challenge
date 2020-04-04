package canvas

import (
	"fmt"
	"math"
	"strconv"
)

const (
	PPM_HEADER   = "P3"
	PPM_CHAR_LEN = 70
)

type Color struct {
	red   float64
	green float64
	blue  float64
}

func (c *Color) multiply(a float64) Color {
	return Color{c.red * a, c.green * a, c.blue * a}
}

func (c *Color) get_min() float64 {
	return math.Min(math.Min(c.red, c.green), c.blue)
}

func (c *Color) get_max() float64 {
	return math.Max(math.Max(c.red, c.green), c.blue)
}

func Add(a Color, b Color) Color {
	return Color{a.red + b.red, a.green + b.green, a.blue + b.blue}
}

func Subtract(a Color, b Color) Color {
	return Color{a.red - b.red, a.green - b.green, a.blue - b.blue}
}

func Hadamard(a Color, b Color) Color {
	return Color{a.red * b.red, a.green * b.green, a.blue * b.blue}
}

type Canvas struct {
	height int
	width  int
	pixels []Color
}

func (c *Canvas) init() {
	c.pixels = make([]Color, c.height*c.width)
}

func (c *Canvas) write_pixel(x int, y int, color Color) {
	c.pixels[x*c.height+y] = color
}

func (c *Canvas) read_pixel(x int, y int) Color {
	return c.pixels[x*c.height+y]
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

func add_color(color float64, s string, row string) (string, string) {
	color_str := scale255(color)

	if len([]rune(row))+len([]rune(color_str)) > PPM_CHAR_LEN {
		writeRow(&s, &row)
		return s, color_str
	} else {
		return s, row + color_str
	}

}

func writeRow(s *string, row *string) {
	*s += (*row)[:len((*row))-1] + "\n"
}

func (c *Canvas) to_ppm() string {
	s := fmt.Sprintf("%s\n%d %d\n255\n", PPM_HEADER, c.width, c.height)

	for i := 0; i < c.height; i++ {
		row := ""
		for j := 0; j < c.width; j++ {
			color := c.read_pixel(j, i)

			//add_color(color.red, &s, &row)
			//add_color(color.red, &s, &row)
			//add_color(color.red, &s, &row)

			s, row = add_color(color.red, s, row)
			s, row = add_color(color.green, s, row)
			s, row = add_color(color.blue, s, row)
		}
		writeRow(&s, &row)
	}

	return s
}
