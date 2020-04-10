package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/matrices"
	"github.com/seantur/ray_tracer_challenge/tuples"
	"reflect"
	"testing"
)

func TestRays(t *testing.T) {

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

	t.Run("create and query a ray", func(t *testing.T) {
		origin := tuples.Point(1, 2, 3)
		direction := tuples.Vector(4, 5, 6)

		r := Ray{origin, direction}

		assertTupleEqual(t, r.Origin, origin)
		assertTupleEqual(t, r.Direction, direction)

	})

	t.Run("Compute a point from a distance", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(2, 3, 4), Direction: tuples.Vector(1, 0, 0)}

		assertTupleEqual(t, r.Position(0), tuples.Point(2, 3, 4))
		assertTupleEqual(t, r.Position(1), tuples.Point(3, 3, 4))
		assertTupleEqual(t, r.Position(-1), tuples.Point(1, 3, 4))
		assertTupleEqual(t, r.Position(2.5), tuples.Point(4.5, 3, 4))

	})

	t.Run("A ray intersect a sphere at two points", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, -5), Direction: tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 4.0)
		assertVal(t, xs[1].T, 6.0)

	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 1, -5), Direction: tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 5.0)
		assertVal(t, xs[1].T, 5.0)

	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 2, -5), Direction: tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 0)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, 0), Direction: tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, -1.0)
		assertVal(t, xs[1].T, 1.0)
	})

	t.Run("A ray is behind a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, 5), Direction: tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, -6.0)
		assertVal(t, xs[1].T, -4.0)
	})

	t.Run("Aggregating intersections", func(t *testing.T) {
		s := GetSphere()

		i1 := Intersection{1, &s}
		i2 := Intersection{2, &s}

		xs := [...]Intersection{i1, i2}

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].T, 1)
		assertVal(t, xs[1].T, 2)

	})

	t.Run("Intersect sets the object on the intersection", func(t *testing.T) {
		r := Ray{tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1)}
		s := GetSphere()

		xs := s.Intersect(r)

		assertVal(t, float64(len(xs)), 2)

		if &s != xs[0].Object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[0].Object)
			t.Fatal()
		}

		if &s != xs[1].Object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[1].Object)
			t.Fatal()
		}

	})

	t.Run("The hit when all intersections have positive t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{1, &s}
		i2 := Intersection{2, &s}

		xs := []Intersection{i1, i2}

		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i1) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("The hit where some intersections have negative t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{-1, &s}
		i2 := Intersection{1, &s}

		xs := []Intersection{i1, i2}

		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i2) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("The hit where all intersections have negative t", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{-2, &s}
		i2 := Intersection{-1, &s}

		xs := []Intersection{i1, i2}

		_, err := Hit(xs)

		if err == nil {
			t.Errorf("expected no hits, but got one")
		}
	})

	t.Run("The hit is always the lower nonnegative intersection", func(t *testing.T) {
		s := GetSphere()
		i1 := Intersection{5, &s}
		i2 := Intersection{7, &s}
		i3 := Intersection{-3, &s}
		i4 := Intersection{2, &s}

		xs := []Intersection{i1, i2, i3, i4}
		i, _ := Hit(xs)

		if !reflect.DeepEqual(i, i4) {
			t.Errorf("expected equal intersections were not equal")
		}
	})

	t.Run("Translating a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(1, 2, 3), Direction: tuples.Vector(0, 1, 0)}
		m := matrices.GetTranslation(3, 4, 5)

		r2 := r.Transform(m)

		assertTupleEqual(t, r2.Origin, tuples.Point(4, 6, 8))
		assertTupleEqual(t, r2.Direction, tuples.Vector(0, 1, 0))
	})

	t.Run("Scaling a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(1, 2, 3), Direction: tuples.Vector(0, 1, 0)}
		m := matrices.GetScaling(2, 3, 4)

		r2 := r.Transform(m)

		assertTupleEqual(t, r2.Origin, tuples.Point(2, 6, 12))
		assertTupleEqual(t, r2.Direction, tuples.Vector(0, 3, 0))
	})

}
