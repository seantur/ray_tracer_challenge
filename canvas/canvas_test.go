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
		c := InitCanvas(20, 10)

		max := c.Bounds().Max
		height, width := max.X, max.Y

		assertVal(t, float64(width), 20)
		assertVal(t, float64(height), 10)

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				val := c.At(i, j)
				r, g, b, _ := val.RGBA()
				output := raytracing.RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}
				raytracing.AssertColorsEqual(t, output, raytracing.RGB{0, 0, 0})
			}
		}

	})

	t.Run("Write a pixel to a canvas", func(t *testing.T) {
		c := InitCanvas(10, 20)

		Red := raytracing.RGB{1, 0, 0}

		c.SetRGBA64(2, 3, Red.Cvt())

		val := c.At(2, 3)
		r, g, b, a := val.RGBA()
		output := raytracing.RGB{float64(r / a), float64(g / a), float64(b / a)}

		raytracing.AssertColorsEqual(t, output, Red)
	})

	t.Run("construct PPM header", func(t *testing.T) {
		c := InitCanvas(3, 5)

		assertStringLine(t, toPPM(c), "P3", 0)
		assertStringLine(t, toPPM(c), "5 3", 1)
		assertStringLine(t, toPPM(c), "255", 2)

	})

	t.Run("construct PPM pixel data", func(t *testing.T) {
		c := InitCanvas(3, 5)

		c1 := raytracing.RGB{Red: 1.5, Green: 0, Blue: 0}
		c2 := raytracing.RGB{Red: 0, Green: 0.5, Blue: 0}
		c3 := raytracing.RGB{Red: -0.5, Green: 0, Blue: 1}

		c.SetRGBA64(0, 0, c1.Cvt())
		c.SetRGBA64(2, 1, c2.Cvt())
		c.SetRGBA64(4, 2, c3.Cvt())

		got := toPPM(c)

		assertStringLine(t, got, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", 3)
		assertStringLine(t, got, "0 0 0 0 0 0 0 127 0 0 0 0 0 0 0", 4)
		assertStringLine(t, got, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", 5)

	})

	t.Run("PPM lines > 70 are on a new line", func(t *testing.T) {
		c := InitCanvas(2, 10)

		max := c.Bounds().Max
		width, height := max.X, max.Y

		color := raytracing.RGB{1, 0.8, 0.6}

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				c.Set(i, j, color)
			}
		}

		got := toPPM(c)

		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 3)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 4)
		assertStringLine(t, got, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", 5)
		assertStringLine(t, got, "153 255 204 153 255 204 153 255 204 153 255 204 153", 6)

	})

	t.Run("ensure ppm string ends with newline", func(t *testing.T) {
		c := InitCanvas(5, 3)

		ppm := toPPM(c)
		endPPM := ppm[len(ppm)-1]

		if endPPM != '\n' {
			t.Error("did not end with newline")
		}
	})
}
