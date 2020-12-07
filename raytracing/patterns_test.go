package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

func TestPatterns(t *testing.T) {
	black := RGB{Red: 0, Green: 0, Blue: 0}
	white := RGB{Red: 1, Green: 1, Blue: 1}

	t.Run("Creating a striped pattern", func(t *testing.T) {
		stripe := Stripe{A: white, B: black}

		AssertColorsEqual(t, stripe.A, white)
		AssertColorsEqual(t, stripe.B, black)

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
		c := AtObj(pattern, obj, datatypes.Point(1.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

	t.Run("Stripes with a pattern transformation", func(t *testing.T) {
		obj := GetSphere()

		pattern := GetStripe(white, black)
		pattern.SetTransform(datatypes.GetScaling(2, 2, 2))
		c := AtObj(pattern, obj, datatypes.Point(1.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

	t.Run("Stripes with an object and a pattern transformation", func(t *testing.T) {
		obj := GetSphere()
		obj.SetTransform(datatypes.GetScaling(2, 2, 2))

		pattern := GetStripe(white, black)
		pattern.SetTransform(datatypes.GetScaling(2, 2, 2))
		c := AtObj(pattern, obj, datatypes.Point(2.5, 0, 0))

		AssertColorsEqual(t, c, white)
	})

	t.Run("A gradient linearly interpolates between colors", func(t *testing.T) {
		pattern := GetGradient(white, black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.25, 0, 0)), RGB{Red: 0.75, Green: 0.75, Blue: 0.75})
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.5, 0, 0)), RGB{Red: 0.5, Green: 0.5, Blue: 0.5})
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.75, 0, 0)), RGB{Red: 0.25, Green: 0.25, Blue: 0.25})
	})

	t.Run("A ring should extend in both x and z", func(t *testing.T) {
		pattern := GetRing(white, black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(1, 0, 0)), black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 1)), black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.708, 0, 0.708)), black)
	})

	t.Run("Checkers should repeat in x", func(t *testing.T) {
		pattern := GetCheckers(white, black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0.99, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(1.01, 0, 0)), black)
	})

	t.Run("Checkers should repeat in y", func(t *testing.T) {
		pattern := GetCheckers(white, black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0.99, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 1.01, 0)), black)
	})

	t.Run("Checkers should repeat in z", func(t *testing.T) {
		pattern := GetCheckers(white, black)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 0.99)), white)
		AssertColorsEqual(t, pattern.At(datatypes.Point(0, 0, 1.01)), black)
	})

	t.Run("The default pattern transformation", func(t *testing.T) {
		pat := GetTestPat()
		datatypes.AssertMatrixEqual(t, pat.GetTransform(), datatypes.GetIdentity())
	})

}
