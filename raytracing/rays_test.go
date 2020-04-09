package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/tuples"
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
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].t, 4.0)
		assertVal(t, xs[1].t, 6.0)

	})

	t.Run("A ray intersects a sphere at a tangent", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 1, -5), Direction: tuples.Vector(0, 0, 1)}
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].t, 5.0)
		assertVal(t, xs[1].t, 5.0)

	})

	t.Run("A ray misses a sphere", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 2, -5), Direction: tuples.Vector(0, 0, 1)}
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 0)
	})

	t.Run("A ray originates inside a sphere", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, 0), Direction: tuples.Vector(0, 0, 1)}
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].t, -1.0)
		assertVal(t, xs[1].t, 1.0)
	})

	t.Run("A ray is behind a ray", func(t *testing.T) {
		r := Ray{Origin: tuples.Point(0, 0, 5), Direction: tuples.Vector(0, 0, 1)}
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].t, -6.0)
		assertVal(t, xs[1].t, -4.0)
	})

	t.Run("Aggregating intersections", func(t *testing.T) {
		s := Sphere{}

		var xs []Intersection

		xs = append(xs, Intersection{1, &s}, Intersection{2, &s})

		assertVal(t, float64(len(xs)), 2)
		assertVal(t, xs[0].t, 1)
		assertVal(t, xs[1].t, 2)

	})

	t.Run("Intersect sets the object on the intersection", func(t *testing.T) {
		r := Ray{tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1)}
		s := Sphere{}

		xs := Intersect(s, r)

		assertVal(t, float64(len(xs)), 2)

		if &s != xs[0].object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[0].object)
			t.Fatal()
		}

		if &s != xs[1].object {
			t.Errorf("expected equal pointers were not equal: %p %p", &s, xs[1].object)
			t.Fatal()
		}

	})

}
