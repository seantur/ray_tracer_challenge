package canvas

import (
	"github.com/seantur/ray_tracer_challenge/raytracing"
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

}
