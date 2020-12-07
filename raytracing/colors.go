package raytracing

// RGB
type RGB struct {
	Red   float64
	Green float64
	Blue  float64
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	alpha := 255.0
	if c.Red < 0 {
		c.Red = 0
	}

	if c.Green < 0 {
		c.Green = 0
	}

	if c.Blue < 0 {
		c.Blue = 0
	}
	return uint32(c.Red * alpha), uint32(c.Green * alpha), uint32(c.Blue * alpha), uint32(alpha)
}

func (c *RGB) Multiply(a float64) RGB {
	return RGB{c.Red * a, c.Green * a, c.Blue * a}
}

func Add(colors ...RGB) (c RGB) {
	for _, color := range colors {
		c.Red += color.Red
		c.Green += color.Green
		c.Blue += color.Blue
	}
	return
}

func Subtract(a RGB, b RGB) RGB {
	return RGB{a.Red - b.Red, a.Green - b.Green, a.Blue - b.Blue}
}

func Hadamard(a RGB, b RGB) RGB {
	return RGB{a.Red * b.Red, a.Green * b.Green, a.Blue * b.Blue}
}

func HexColor(hex int) RGB {
	return RGB{float64((hex&0xFF0000)>>16) / 255.0, float64((hex&0x00FF00)>>8) / 255.0, float64(hex&0x0000FF) / 255.}
}

const (
	Black   = 0x000000
	White   = 0xFFFFFF
	Red     = 0xFF0000
	Green   = 0x00FF00
	Blue    = 0x0000FF
	Yellow  = 0xFFFF00
	Magenta = 0xFF00FF
	Cyan    = 0x00FFFF
	Orange  = 0xFFA500
	Navy    = 0x000080
	Teal    = 0x008080
	Purple  = 0x800080
	Pink    = 0xFF1493
)
