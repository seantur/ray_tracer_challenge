package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/matrices"
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
	"reflect"
	"testing"
)

func TestShapes(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	assertTupleEqual := func(t *testing.T, got tuples.Tuple, want tuples.Tuple) {
		t.Helper()
		if !tuples.Equal(got, want) {
			t.Error("wanted equal tuples are not equal")
		}
	}

	assertMatrixEqual := func(t *testing.T, got matrices.Matrix, want matrices.Matrix) {
		t.Helper()
		if !matrices.Equal(got, want) {
			t.Error("wanted equal matrices are not equal")
		}
	}

	t.Run("Sphere's default transform is identity matrix", func(t *testing.T) {
		s := GetSphere()

		assertMatrixEqual(t, s.transform, matrices.GetIdentity())
	})

	t.Run("Changing a sphere's transformation", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(matrices.GetTranslation(2, 3, 4))

		assertMatrixEqual(t, s.transform, matrices.GetTranslation(2, 3, 4))
	})

	t.Run("Intersecting a scaled sphere with a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, -5), Direction: tuples.Vector(0, 0, 1)}

		s := GetSphere()

		s.SetTransform(matrices.GetScaling(2, 2, 2))
		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 3)
		assertVal(t, xs[1].T, 7)
	})

	t.Run("Intersecting a translated sphere with a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, -5), Direction: tuples.Vector(0, 0, 1)}

		s := GetSphere()

		s.SetTransform(matrices.GetTranslation(5, 0, 0))
		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 0)
	})

	t.Run("The normal on a sphere at a point on the x axis", func(t *testing.T) {
		s := GetSphere()

		n := s.GetNormal(tuples.Point(1, 0, 0))
		assertTupleEqual(t, n, tuples.Vector(1, 0, 0))
	})

	t.Run("The normal on a sphere at a point on the y axis", func(t *testing.T) {
		s := GetSphere()

		n := s.GetNormal(tuples.Point(0, 1, 0))
		assertTupleEqual(t, n, tuples.Vector(0, 1, 0))
	})

	t.Run("The normal on a sphere at a point on the z axis", func(t *testing.T) {
		s := GetSphere()

		n := s.GetNormal(tuples.Point(0, 0, 1))
		assertTupleEqual(t, n, tuples.Vector(0, 0, 1))
	})

	t.Run("The normal on a sphere at a noaxial point", func(t *testing.T) {
		s := GetSphere()

		n := s.GetNormal(tuples.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		assertTupleEqual(t, n, tuples.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	})

	t.Run("The normal is a normalized vector", func(t *testing.T) {
		s := GetSphere()

		n := s.GetNormal(tuples.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
		assertTupleEqual(t, n.Normalize(), n)
	})

	t.Run("Computing the normal on a translated sphere", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(matrices.GetTranslation(0, 1, 0))

		n := s.GetNormal(tuples.Point(0, 1.70711, -.70711))
		assertTupleEqual(t, n, tuples.Vector(0, 0.70711, -.70711))
	})

	t.Run("Computing the normal on a transformed sphere", func(t *testing.T) {
		s := GetSphere()
		s.SetTransform(matrices.Multiply(matrices.GetScaling(1, 0.5, 1), matrices.GetRotationZ(math.Pi/5)))

		n := s.GetNormal(tuples.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
		assertTupleEqual(t, n, tuples.Vector(0, 0.97014, -0.24254))
	})

	t.Run("A sphere has a default material", func(t *testing.T) {
		s := GetSphere()
		m := GetMaterial()

		if !reflect.DeepEqual(s.Material, m) {
			t.Error("expected sphere material is not the default")
		}
	})

	t.Run("A sphere may be assigned a materials", func(t *testing.T) {
		s := GetSphere()
		m := GetMaterial()

		m.Ambient = 1
		s.Material = m

		if !reflect.DeepEqual(s.Material, m) {
			t.Error("sphere material is not expected")
		}
	})

}
