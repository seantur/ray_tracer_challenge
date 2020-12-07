package scene

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
	"testing"
)

func TestCamera(t *testing.T) {

	t.Run("Constructing a camera", func(t *testing.T) {
		c := GetCamera(160, 120, math.Pi/2)

		datatypes.AssertVal(t, float64(c.Hsize), 160)
		datatypes.AssertVal(t, float64(c.Vsize), 120)
		datatypes.AssertVal(t, c.Fov, math.Pi/2)
		datatypes.AssertMatrixEqual(t, c.Transform, datatypes.GetIdentity())
	})

	t.Run("The pixel size for a horizontal canvas", func(t *testing.T) {
		c := GetCamera(200, 125, math.Pi/2)
		datatypes.AssertVal(t, c.PixelSize, 0.01)
	})

	t.Run("The pixel size for a vertical canvas", func(t *testing.T) {
		c := GetCamera(125, 200, math.Pi/2)
		datatypes.AssertVal(t, c.PixelSize, 0.01)
	})

	/*
		t.Run("Constructing a ray through the center of the canvas", func(t *testing.T) {
			c := GetCamera(201, 101, math.Pi/2)
			r := c.RayForPixel(100, 50)

			datatypes.AssertTupleEqual(t, r.Origin, datatypes.Point(0, 0, 0))
			datatypes.AssertTupleEqual(t, r.Direction, datatypes.Vector(0, 0, -1))
		})

		t.Run("Constructing a ray through the corner of the canvas", func(t *testing.T) {
			c := GetCamera(201, 101, math.Pi/2)
			r := c.RayForPixel(0, 0)

			datatypes.AssertTupleEqual(t, r.Origin, datatypes.Point(0, 0, 0))
			datatypes.AssertTupleEqual(t, r.Direction, datatypes.Vector(0.66519, 0.33259, -0.66851))
		})

		t.Run("Constructing a ray when the camera is transformed", func(t *testing.T) {
			c := GetCamera(201, 101, math.Pi/2)
			c.Transform = datatypes.Multiply(datatypes.GetRotationY(math.Pi/4), datatypes.GetTranslation(0, -2, 5))
			r := c.RayForPixel(100, 50)

			datatypes.AssertTupleEqual(t, r.Origin, datatypes.Point(0, 2, -5))
			datatypes.AssertTupleEqual(t, r.Direction, datatypes.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2))
		})
	*/

	t.Run("Rendering a world with the camera", func(t *testing.T) {
		w := GetWorld()
		c := GetCamera(11, 11, math.Pi/2)

		from := datatypes.Point(0, 0, -5)
		to := datatypes.Point(0, 0, 0)
		up := datatypes.Point(0, 1, 0)
		Transform := datatypes.ViewTransform(from, to, up)

		c.Transform = Transform

		im := c.Render(w)
		r, g, b, _ := im.At(5, 5).RGBA()
		output := raytracing.RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}
		raytracing.AssertColorsEqual(t, output, raytracing.RGB{Red: 0.38039, Green: 0.47450, Blue: 0.28235})
	})

}
