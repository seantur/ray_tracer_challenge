package canvas

import (
	"github.com/seantur/ray_tracer_challenge/raytracing"
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

	t.Run("Create a canvas", func(t *testing.T) {
		c := Canvas{Height: 10, Width: 20}
		c.Init()

		assertVal(t, float64(c.Width), 20)
		assertVal(t, float64(c.Height), 10)

		for i := 0; i < c.Width; i++ {
			for j := 0; j < c.Height; j++ {
				val, _ := c.ReadPixel(i, j)
				r, g, b, _ := val.RGBA()
				output := raytracing.RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}
				raytracing.AssertColorsEqual(t, output, raytracing.RGB{0, 0, 0})
			}
		}

	})

	t.Run("Write a pixel to a canvas", func(t *testing.T) {
		c := Canvas{Height: 20, Width: 10}
		c.Init()

		Red := raytracing.RGB{1, 0, 0}

		c.WritePixel(2, 3, Red)

		val, _ := c.ReadPixel(2, 3)
		r, g, b, _ := val.RGBA()
		output := raytracing.RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}

		raytracing.AssertColorsEqual(t, output, Red)
	})

	t.Run("construct PPM header", func(t *testing.T) {
		c := Canvas{Height: 3, Width: 5}
		c.Init()

		assertStringLine(t, c.toPPM(), "P3", 0)
		assertStringLine(t, c.toPPM(), "5 3", 1)
		assertStringLine(t, c.toPPM(), "255", 2)

	})

	t.Run("construct PPM pixel data", func(t *testing.T) {
		c := Canvas{Height: 3, Width: 5}
		c.Init()

		c1 := raytracing.RGB{1.5, 0, 0}
		c2 := raytracing.RGB{0, 0.5, 0}
		c3 := raytracing.RGB{-0.5, 0, 1}

		c.WritePixel(0, 0, c1)
		c.WritePixel(2, 1, c2)
		c.WritePixel(4, 2, c3)

		got := c.toPPM()

		assertStringLine(t, got, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", 3)
		assertStringLine(t, got, "0 0 0 0 0 0 0 127 0 0 0 0 0 0 0", 4)
		assertStringLine(t, got, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", 5)

	})

	t.Run("PPM lines > 70 are on a new line", func(t *testing.T) {
		c := Canvas{Height: 2, Width: 10}
		c.Init()

		color := raytracing.RGB{1, 0.8, 0.6}

		for i := 0; i < c.Width; i++ {
			for j := 0; j < c.Height; j++ {
				c.WritePixel(i, j, color)
			}
		}

		got := c.toPPM()

		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 3)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 4)
		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 5)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 6)

	})

	t.Run("ensure ppm string ends with newline", func(t *testing.T) {
		c := Canvas{Height: 3, Width: 5}
		c.Init()

		ppm := c.toPPM()
		endPPM := ppm[len(ppm)-1]

		if endPPM != '\n' {
			t.Error("did not end with newline")
		}
	})

	t.Run("ensure out of bounds doesn't crash", func(t *testing.T) {
		c := Canvas{Height: 3, Width: 5}
		c.Init()

		_, err := c.ReadPixel(3, 5)

		if err.Error() != ErrOutOfBounds {
			t.Error("did not throw appropriate error")
		}

	})
}
