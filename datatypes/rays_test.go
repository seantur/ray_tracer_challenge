package datatypes

import (
	"testing"
)

func TestRays(t *testing.T) {

	t.Run("create and query a ray", func(t *testing.T) {
		origin := Point(1, 2, 3)
		direction := Vector(4, 5, 6)

		r := Ray{origin, direction}

		AssertTupleEqual(t, r.Origin, origin)
		AssertTupleEqual(t, r.Direction, direction)

	})

	t.Run("Compute a point from a distance", func(t *testing.T) {
		r := Ray{Origin: Point(2, 3, 4), Direction: Vector(1, 0, 0)}

		AssertTupleEqual(t, r.Position(0), Point(2, 3, 4))
		AssertTupleEqual(t, r.Position(1), Point(3, 3, 4))
		AssertTupleEqual(t, r.Position(-1), Point(1, 3, 4))
		AssertTupleEqual(t, r.Position(2.5), Point(4.5, 3, 4))

	})

	t.Run("Translating a ray", func(t *testing.T) {
		r := Ray{Origin: Point(1, 2, 3), Direction: Vector(0, 1, 0)}
		m := GetTranslation(3, 4, 5)

		r2 := r.Transform(m)

		AssertTupleEqual(t, r2.Origin, Point(4, 6, 8))
		AssertTupleEqual(t, r2.Direction, Vector(0, 1, 0))
	})

	t.Run("Scaling a ray", func(t *testing.T) {
		r := Ray{Origin: Point(1, 2, 3), Direction: Vector(0, 1, 0)}
		m := GetScaling(2, 3, 4)

		r2 := r.Transform(m)

		AssertTupleEqual(t, r2.Origin, Point(2, 6, 12))
		AssertTupleEqual(t, r2.Direction, Vector(0, 3, 0))
	})

}
