package raytracing

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
