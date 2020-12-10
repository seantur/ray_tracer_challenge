package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

func TestCubes(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if !datatypes.IsClose(got, want) {
			t.Errorf("got %f want %f", got, want)
		}

	}

	t.Run("A ray intersects a cube", func(t *testing.T) {

		type OriginDirectionTestCase struct {
			ray    datatypes.Ray
			T1, T2 float64
		}

		testcases := []OriginDirectionTestCase{
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(5, 0.5, 0), datatypes.Vector(-1, 0, 0)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(-5, 0.5, 0), datatypes.Vector(1, 0, 0)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0.5, 5, 0), datatypes.Vector(0, -1, 0)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0.5, -5, 0), datatypes.Vector(0, 1, 0)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0.5, 0, 5), datatypes.Vector(0, 0, -1)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0.5, 0, -5), datatypes.Vector(0, 0, 1)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0, 0.5, 0), datatypes.Vector(0, 0, 1)}, -1, 1}}

		c := GetCube()

		for _, testcase := range testcases {
			xs := Intersect(c, testcase.ray)
			assertVal(t, float64(len(xs)), 2)

			assertVal(t, xs[0].T, testcase.T1)
			assertVal(t, xs[1].T, testcase.T2)
		}

	})

	t.Run("A ray misses a cube", func(t *testing.T) {

		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(-2, 0, 0), datatypes.Vector(0.2673, 0.5345, 0.8018)},
			datatypes.Ray{datatypes.Point(0, -2, 0), datatypes.Vector(0.8018, 0.2673, 0.5345)},
			datatypes.Ray{datatypes.Point(0, 0, -2), datatypes.Vector(0.5345, 0.8018, 0.2673)},
			datatypes.Ray{datatypes.Point(2, 0, 2), datatypes.Vector(0, 0, -1)},
			datatypes.Ray{datatypes.Point(0, 2, 2), datatypes.Vector(0, -1, 0)},
			datatypes.Ray{datatypes.Point(2, 2, 0), datatypes.Vector(-1, 0, 0)}}

		c := GetCube()

		for _, ray := range testcases {
			xs := Intersect(c, ray)
			assertVal(t, float64(len(xs)), 0)
		}

	})

	t.Run("The normal on the surface of a cube", func(t *testing.T) {

		// This is a slight abuse of notation, but Ray is the obvious container for a Point + Vector
		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(1, 0.5, -0.8), datatypes.Vector(1, 0, 0)},
			datatypes.Ray{datatypes.Point(-1, -0.2, 0.9), datatypes.Vector(-1, 0, 0)},
			datatypes.Ray{datatypes.Point(-0.4, 1, -0.1), datatypes.Vector(0, 1, 0)},
			datatypes.Ray{datatypes.Point(0.3, -1, -0.7), datatypes.Vector(0, -1, 0)},
			datatypes.Ray{datatypes.Point(-0.6, 0.3, 1), datatypes.Vector(0, 0, 1)},
			datatypes.Ray{datatypes.Point(0.4, 0.4, -1), datatypes.Vector(0, 0, -1)},
			datatypes.Ray{datatypes.Point(1, 1, 1), datatypes.Vector(1, 0, 0)},
			datatypes.Ray{datatypes.Point(-1, -1, -1), datatypes.Vector(-1, 0, 0)}}

		c := GetCube()

		for _, ray := range testcases {
			normal := Normal(c, ray.Origin)
			if !datatypes.IsTupleEqual(normal, ray.Direction) {
				t.Fatalf("Actual normal vector %v != expected normal vector %v", normal, ray.Direction)
			}
		}

	})

}
