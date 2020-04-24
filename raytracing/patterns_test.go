package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

func TestPatterns(t *testing.T) {
	black := Color{Red: 0, Green: 0, Blue: 0}
	white := Color{Red: 1, Green: 1, Blue: 1}

	t.Run("Creating a striped pattern", func(t *testing.T) {
		pattern := GetStripe(white, black)

		AssertColorsEqual(t, pattern.A, white)
		AssertColorsEqual(t, pattern.B, black)

	})

	t.Run("A stripe pattern is constant in y", func(t *testing.T) {
		pattern := GetStripe(white, black)

		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 1, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 2, 0)), white)
	})

	t.Run("A stripe pattern is constant in z", func(t *testing.T) {
		pattern := GetStripe(white, black)

		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 1)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 2)), white)
	})

	t.Run("A stripe pattern is alternates in x", func(t *testing.T) {
		pattern := GetStripe(white, black)

		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.9, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(1, 0, 0)), black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(-0.1, 0, 0)), black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(-1, 0, 0)), black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(-1.1, 0, 0)), white)
	})

	t.Run("Stripes with an object transformation", func(t *testing.T) {
		obj := GetSphere()
		obj.SetTransform(datatypes.GetScaling(2, 2, 2))

		pattern := GetStripe(white, black)
		c := pattern.AtObj(obj, datatypes.Point(1.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

	t.Run("Stripes with a pattern transformation", func(t *testing.T) {
		obj := GetSphere()

		pattern := GetStripe(white, black)
		pattern.SetTransform(datatypes.GetScaling(2, 2, 2))
		c := pattern.AtObj(obj, datatypes.Point(1.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

	t.Run("Stripes with an object and a pattern transformation", func(t *testing.T) {
		obj := GetSphere()
		obj.SetTransform(datatypes.GetScaling(2, 2, 2))

		pattern := GetStripe(white, black)
		pattern.SetTransform(datatypes.GetScaling(2, 2, 2))
		c := pattern.AtObj(obj, datatypes.Point(2.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

}
