package canvas

import (
	"github.com/seantur/ray_tracer_challenge/tuples"
	"strings"
	"testing"
)

const EPSILON = .00001

func TestCanvas(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	assertColorsEqual := func(t *testing.T, got Color, want Color) {
		t.Helper()

		allClose := tuples.IsClose(got.red, want.red) && tuples.IsClose(got.green, want.green) && tuples.IsClose(got.blue, want.blue)

		if !allClose {
			t.Error("wanted equal colors are not equal")
		}
	}

	assertStringLine := func(t *testing.T, got string, want_line string, line int) {
		t.Helper()

		got_split := strings.Split(got, "\n")

		if len(got_split) <= line {
			t.Fatal("trying to compare line out of bounds")
		}

		got_line := got_split[line]

		if got_line != want_line {
			t.Errorf("got %s want %s", got_line, want_line)
		}

	}

	t.Run("Colors are a tuple", func(t *testing.T) {
		c := Color{-0.5, 0.4, 1.7}

		assertVal(t, c.red, -0.5)
		assertVal(t, c.green, 0.4)
		assertVal(t, c.blue, 1.7)
	})

	t.Run("Adding colors", func(t *testing.T) {
		c1 := Color{0.9, 0.6, 0.75}
		c2 := Color{0.7, 0.1, 0.25}

		assertColorsEqual(t, Add(c1, c2), Color{1.6, 0.7, 1.0})
	})

	t.Run("Subtracting colors", func(t *testing.T) {
		c1 := Color{0.9, 0.6, 0.75}
		c2 := Color{0.7, 0.1, 0.25}

		assertColorsEqual(t, Subtract(c1, c2), Color{0.2, 0.5, 0.5})
	})

	t.Run("Multiply a color by a scalar", func(t *testing.T) {
		c := Color{0.2, 0.3, 0.4}

		assertColorsEqual(t, c.multiply(2.), Color{0.4, 0.6, 0.8})
	})

	t.Run("Multiply 2 colors together", func(t *testing.T) {
		c1 := Color{1, 0.2, 0.4}
		c2 := Color{0.9, 1, 0.1}

		assertColorsEqual(t, Hadamard(c1, c2), Color{0.9, 0.2, 0.04})
	})

	t.Run("Create a canvas", func(t *testing.T) {
		c := Canvas{height: 10, width: 20}
		c.init()

		assertVal(t, float64(c.width), 20)
		assertVal(t, float64(c.height), 10)

		for i := 0; i < c.width; i++ {
			for j := 0; j < c.height; j++ {
				assertColorsEqual(t, c.read_pixel(i, j), Color{0, 0, 0})
			}
		}

	})

	t.Run("Write a pixel to a canvas", func(t *testing.T) {
		c := Canvas{height: 20, width: 10}
		c.init()

		red := Color{1, 0, 0}

		c.write_pixel(2, 3, red)

		assertColorsEqual(t, c.read_pixel(2, 3), red)
	})

	t.Run("construct PPM header", func(t *testing.T) {
		c := Canvas{height: 3, width: 5}
		c.init()

		assertStringLine(t, c.to_ppm(), "P3", 0)
		assertStringLine(t, c.to_ppm(), "5 3", 1)
		assertStringLine(t, c.to_ppm(), "255", 2)

	})

	t.Run("construct PPM pixel data", func(t *testing.T) {
		c := Canvas{height: 3, width: 5}
		c.init()

		c1 := Color{1.5, 0, 0}
		c2 := Color{0, 0.5, 0}
		c3 := Color{-0.5, 0, 1}

		c.write_pixel(0, 0, c1)
		c.write_pixel(2, 1, c2)
		c.write_pixel(4, 2, c3)

		got := c.to_ppm()

		assertStringLine(t, got, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", 3)
		assertStringLine(t, got, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", 4)
		assertStringLine(t, got, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", 5)

	})

	t.Run("PPM lines > 70 are on a new line", func(t *testing.T) {
		c := Canvas{height: 2, width: 10}
		c.init()

		color := Color{1, 0.8, 0.6}

		for i := 0; i < c.width; i++ {
			for j := 0; j < c.height; j++ {
				c.write_pixel(i, j, color)
			}
		}

		got := c.to_ppm()

		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 3)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 4)
		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 5)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 6)

	})

}
