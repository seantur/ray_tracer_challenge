package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
	"testing"
)

func TestPlanes(t *testing.T) {

	t.Run("Plane's default transform is identity matrix", func(t *testing.T) {
		p := GetPlane()

		datatypes.AssertMatrixEqual(t, p.GetTransform(), datatypes.GetIdentity())
	})

	t.Run("Changing a plane's transformation", func(t *testing.T) {
		p := GetPlane()
		p.SetTransform(datatypes.GetTranslation(2, 3, 4))

		datatypes.AssertMatrixEqual(t, p.GetTransform(), datatypes.GetTranslation(2, 3, 4))
	})

	t.Run("A plane's normal is constant everywhere", func(t *testing.T) {
		p := GetPlane()

		datatypes.AssertTupleEqual(t, p.Normal(datatypes.Point(0, 0, 0)), datatypes.Vector(0, 1, 0))
		datatypes.AssertTupleEqual(t, p.Normal(datatypes.Point(10, 0, -10)), datatypes.Vector(0, 1, 0))
		datatypes.AssertTupleEqual(t, p.Normal(datatypes.Point(-5, 0, 150)), datatypes.Vector(0, 1, 0))
	})

	t.Run("Intersect with a ray parallel to the plane", func(t *testing.T) {
		p := GetPlane()
		r := Ray{Origin: datatypes.Point(0, 10, 0), Direction: datatypes.Vector(0, 0, 1)}
		xs := p.Intersect(r)

		datatypes.AssertVal(t, float64(len(xs)), 0)
	})

	t.Run("Intersect with a coplanar ray", func(t *testing.T) {
		p := GetPlane()
		r := Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 0, 1)}
		xs := p.Intersect(r)

		datatypes.AssertVal(t, float64(len(xs)), 0)
	})

	t.Run("A ray intersecting a plane from above", func(t *testing.T) {
		p := GetPlane()
		r := Ray{Origin: datatypes.Point(0, 1, 0), Direction: datatypes.Vector(0, -1, 0)}
		xs := p.Intersect(r)

		datatypes.AssertVal(t, float64(len(xs)), 1)
		datatypes.AssertVal(t, xs[0].T, 1)
		if xs[0].Object != p {
			t.Error("Expected intersection object didn't match")
		}
	})

	t.Run("A ray intersecting a plane from below", func(t *testing.T) {
		p := GetPlane()
		r := Ray{Origin: datatypes.Point(0, -1, 0), Direction: datatypes.Vector(0, 1, 0)}
		xs := p.Intersect(r)

		datatypes.AssertVal(t, float64(len(xs)), 1)
		datatypes.AssertVal(t, xs[0].T, 1)
		if xs[0].Object != p {
			t.Error("Expected intersection object didn't match")
		}
	})

	t.Run("Ensure we precompute the reflection vector", func(t *testing.T) {
		p := GetPlane()
		r := Ray{Origin: datatypes.Point(0, 1, -1), Direction: datatypes.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2)}
		i := Intersection{T: math.Sqrt(2), Object: p}

		comps := i.PrepareComputations(r, []Intersection{i})
		datatypes.AssertTupleEqual(t, comps.Reflectv, datatypes.Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2))
	})

}
