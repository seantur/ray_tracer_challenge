package canvas

type Color struct {
	red   float64
	green float64
	blue  float64
}

func (c *Color) multiply(a float64) Color {
	return Color{c.red * a, c.green * a, c.blue * a}
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
