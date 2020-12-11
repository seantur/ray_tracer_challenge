package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
	"testing"
)

func TestCones(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if !datatypes.IsClose(got, want) {
			t.Errorf("got %f want %f", got, want)
		}

	}

	t.Run("Intersecting a cone with a ray", func(t *testing.T) {
		testcases := []OriginDirectionTestCase{
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 0, 1)}, 5, 5},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(1, 1, 1)}, 8.66025, 8.66025},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(1, 1, -5), datatypes.Vector(-0.5, -1, 1)}, 4.55006, 49.44994}}

		c := GetCone()

		for _, testcase := range testcases {
			testcase.ray.Direction = testcase.ray.Direction.Normalize()
			xs := Intersect(c, testcase.ray)
			assertVal(t, float64(len(xs)), 2)

			if len(xs) != 2 {
				continue
				//t.Fatal() // don't index into an array if the points don't exist
			}

			assertVal(t, xs[0].T, testcase.T1)
			assertVal(t, xs[1].T, testcase.T2)
		}

	})

	t.Run("Intersecting a cone with a ray parallel to one of its halves", func(t *testing.T) {
		c := GetCone()
		direction := datatypes.Vector(0, 1, 1)
		direction = direction.Normalize()
		r := datatypes.Ray{datatypes.Point(0, 0, -1), direction}
		xs := Intersect(c, r)

		assertVal(t, float64(len(xs)), 1)

		if len(xs) != 1 {
			t.Fatal()
		}

		assertVal(t, xs[0].T, 0.35355)
	})

	t.Run("Intersecting a cone's end cap", func(t *testing.T) {
		c := GetCone()
		c.Min = -0.5
		c.Max = 0.5
		c.Closed = true

		testcases := []OriginDirectionCount{
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 1, 0)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 0, -0.25), datatypes.Vector(0, 1, 1)}, 2},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 0, -0.25), datatypes.Vector(0, 1, 0)}, 4}}

		for _, testcase := range testcases {
			testcase.ray.Direction = testcase.ray.Direction.Normalize()
			xs := Intersect(c, testcase.ray)

			assertVal(t, float64(len(xs)), testcase.count)
		}

	})

	t.Run("Computing the normal vector on a cone", func(t *testing.T) {
		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(0, 0, 0), datatypes.Vector(0, 0, 0)},
			datatypes.Ray{datatypes.Point(1, 1, 1), datatypes.Vector(1, -math.Sqrt(2), 1)},
			datatypes.Ray{datatypes.Point(-1, -1, 0), datatypes.Vector(-1, 1, 0)}}

		c := GetCone()

		for _, ray := range testcases {
			normal := c.Normal(ray.Origin)
			if !datatypes.IsTupleEqual(normal, ray.Direction) {
				t.Fatalf("Actual normal vector %v != expected normal vector %v", normal, ray.Direction)
			}
		}

	})

}
