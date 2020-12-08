package raytracing

import "image/color"

// RGB
type RGB struct {
	Red   float64
	Green float64
	Blue  float64
}

func clip(val float64, lower_bound, upper_bound uint32) uint32 {

	if val < float64(lower_bound) {
		return lower_bound
	} else if val > float64(upper_bound) {
		return upper_bound
	} else {
		return uint32(val)
	}
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	a = 0xFFFF

	r = clip(c.Red*float64(a), 0, a)
	g = clip(c.Green*float64(a), 0, a)
	b = clip(c.Blue*float64(a), 0, a)

	return
}

func (c RGB) Cvt() color.RGBA64 {
	r, g, b, a := c.RGBA()
	return color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
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
