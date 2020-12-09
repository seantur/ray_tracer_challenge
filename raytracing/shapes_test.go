package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
	"reflect"
	"testing"
)

func TestShapes(t *testing.T) {
	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if !datatypes.IsClose(got, want) {
			t.Errorf("got %f want %f", got, want)
		}

	}
	t.Run("A ray intersect a sphere at two points", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, -5), Direction: datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 4.0)
		assertVal(t, xs[1].T, 6.0)

	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 1, -5), Direction: datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 5.0)
		assertVal(t, xs[1].T, 5.0)

	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 2, -5), Direction: datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 0)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, 0), Direction: datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, -1.0)
		assertVal(t, xs[1].T, 1.0)
	})

	t.Run("A ray is behind a ray", func(t *testing.T) {
		r := datatypes.Ray{Origin: datatypes.Point(0, 0, 5), Direction: datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, -6.0)
		assertVal(t, xs[1].T, -4.0)
	})

	t.Run("Aggregating intersections", func(t *testing.T) {
		s := GetSphere()

		i1 := Intersection{1, s}
		i2 := Intersection{2, s}

		xs := [...]Intersection{i1, i2}

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 1)
		assertVal(t, xs[1].T, 2)

	})

	t.Run("Intersect sets the object on the intersection", func(t *testing.T) {
		r := datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 0, 1)}
		s := GetSphere()

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)

		if s != xs[0].Object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[0].Object)
			t.Fatal()
		}

		if s != xs[1].Object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[1].Object)
			t.Fatal()
		}

	})

	t.Run("The hit when all intersections have positive t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{1, s}
		i2 := Intersection{2, s}

		xs := []Intersection{i1, i2}

		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i1) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("The hit where some intersections have negative t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{-1, s}
		i2 := Intersection{1, s}

		xs := []Intersection{i1, i2}

		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i2) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("The hit where all intersections have negative t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{-2, s}
		i2 := Intersection{-1, s}

		xs := []Intersection{i1, i2}

		_, err := Hit(xs)

		if err == nil {
			t.Errorf("expected no hits, but got one")
		}
	})

	t.Run("The hit is always the lower nonnegative intersection", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{5, s}
		i2 := Intersection{7, s}
		i3 := Intersection{-3, s}
		i4 := Intersection{2, s}

		xs := []Intersection{i1, i2, i3, i4}
		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i4) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("The hit returns an error when passed an empty Intersection slice", func(t *testing.T) {
		i := make([]Intersection, 0)
		_, err := Hit(i)

		if err == nil {
			t.Errorf("expected no hits, but got one")
		}
	})

	t.Run("The Schlick approximation under total internal reflection", func(t *testing.T) {
		s := GetGlassSphere()

		r := datatypes.Ray{datatypes.Point(0, 0, math.Sqrt(2)/2), datatypes.Vector(0, 1, 0)}
		xs := []Intersection{Intersection{-math.Sqrt(2) / 2, s}, Intersection{math.Sqrt(2) / 2, s}}
		comps := xs[1].PrepareComputations(r, xs)

		reflectance := Schlick(comps)

		assertVal(t, reflectance, 1.0)
	})

	t.Run("The Schlick approximation with a perpendicular viewing angle", func(t *testing.T) {
		s := GetGlassSphere()

		r := datatypes.Ray{datatypes.Point(0, 0, 0), datatypes.Vector(0, 1, 0)}
		xs := []Intersection{Intersection{-1, s}, Intersection{1, s}}
		comps := xs[1].PrepareComputations(r, xs)

		assertVal(t, Schlick(comps), 0.04)
	})

	t.Run("The Schlick approximation with small angle and n2 > n1", func(t *testing.T) {
		s := GetGlassSphere()

		r := datatypes.Ray{datatypes.Point(0, 0.99, -2), datatypes.Vector(0, 0, 1)}
		xs := []Intersection{Intersection{1.8589, s}}
		comps := xs[0].PrepareComputations(r, xs)

		assertVal(t, Schlick(comps), 0.48873)
	})
}
