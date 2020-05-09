package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

func TestColors(t *testing.T) {

	t.Run("Colors are a tuple", func(t *testing.T) {
		c := Color{-0.5, 0.4, 1.7}

		datatypes.AssertVal(t, c.Red, -0.5)
		datatypes.AssertVal(t, c.Green, 0.4)
		datatypes.AssertVal(t, c.Blue, 1.7)
	})

	t.Run("Adding colors", func(t *testing.T) {
		c1 := Color{0.9, 0.6, 0.75}
		c2 := Color{0.7, 0.1, 0.25}

		AssertColorsEqual(t, Add(c1, c2), Color{1.6, 0.7, 1.0})
	})

	t.Run("Subtracting colors", func(t *testing.T) {
		c1 := Color{0.9, 0.6, 0.75}
		c2 := Color{0.7, 0.1, 0.25}

		AssertColorsEqual(t, Subtract(c1, c2), Color{0.2, 0.5, 0.5})
	})

	t.Run("Multiply a color by a scalar", func(t *testing.T) {
		c := Color{0.2, 0.3, 0.4}

		AssertColorsEqual(t, c.Multiply(2.), Color{0.4, 0.6, 0.8})
	})

	t.Run("Multiply 2 colors together", func(t *testing.T) {
		c1 := Color{1, 0.2, 0.4}
		c2 := Color{0.9, 1, 0.1}

		AssertColorsEqual(t, Hadamard(c1, c2), Color{0.9, 0.2, 0.04})
	})

	t.Run("HexColor convert hex to color correctly", func(t *testing.T) {
		c1 := Color{1, 0, 0}
		c2 := HexColor(0xFF0000)
		AssertColorsEqual(t, c1, c2)

		c1 = Color{0, 1, 0}
		c2 = HexColor(0x00FF00)
		AssertColorsEqual(t, c1, c2)

		c1 = Color{0, 0, 1}
		c2 = HexColor(0x0000FF)
		AssertColorsEqual(t, c1, c2)
	})

	t.Run("Color add is variadic", func(t *testing.T) {
		c1 := Color{0, 1, 2}
		c2 := Color{1, 2, 1}
		c3 := Color{4, 3, 2}

		AssertColorsEqual(t, Add(c1, c2, c3), Color{5, 6, 5})
	})

}
