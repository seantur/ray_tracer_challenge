package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"testing"
)

type OriginDirectionTestCase struct {
	ray    datatypes.Ray
	T1, T2 float64
}

type OriginDirectionCount struct {
	ray   datatypes.Ray
	count float64
}

func TestCylinders(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if !datatypes.IsClose(got, want) {
			t.Errorf("got %f want %f", got, want)
		}

	}

	t.Run("A ray missing a cylinder", func(t *testing.T) {

		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(1, 0, 0), datatypes.Vector(0, 1, 0)},
			datatypes.Ray{datatypes.Point(0, 0, 0), datatypes.Vector(0, 1, 0)},
			datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(1, 1, 1)}}

		c := GetCylinder()

		for _, ray := range testcases {
			xs := Intersect(c, ray)
			assertVal(t, float64(len(xs)), 0)
		}

	})

	t.Run("A ray strikes a cylinder", func(t *testing.T) {
		testcases := []OriginDirectionTestCase{
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(1, 0, -5), datatypes.Vector(0, 0, 1)}, 5, 5},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 0, 1)}, 4, 6},
			OriginDirectionTestCase{datatypes.Ray{datatypes.Point(0.5, 0, -5), datatypes.Vector(0.1, 1, 1)}, 6.80798, 7.08872}}

		c := GetCylinder()

		for _, testcase := range testcases {
			ray := testcase.ray
			ray.Direction = ray.Direction.Normalize()
			xs := Intersect(c, ray)
			assertVal(t, float64(len(xs)), 2)

			assertVal(t, xs[0].T, testcase.T1)
			assertVal(t, xs[1].T, testcase.T2)
		}

	})

	t.Run("Normal vector on a cylinder", func(t *testing.T) {
		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(1, 0, 0), datatypes.Vector(1, 0, 0)},
			datatypes.Ray{datatypes.Point(0, 5, -1), datatypes.Vector(0, 0, -1)},
			datatypes.Ray{datatypes.Point(0, -2, 1), datatypes.Vector(0, 0, 1)},
			datatypes.Ray{datatypes.Point(-1, 1, 0), datatypes.Vector(-1, 0, 0)}}

		c := GetCylinder()

		for _, ray := range testcases {
			normal := Normal(c, ray.Origin)
			if !datatypes.IsTupleEqual(normal, ray.Direction) {
				t.Fatalf("Actual normal vector %v != expected normal vector %v", normal, ray.Direction)
			}
		}

	})

	t.Run("The default min and max for a cylinder", func(t *testing.T) {
		c := GetCylinder()

		assertVal(t, c.Min, -datatypes.INFINITY)
		assertVal(t, c.Max, datatypes.INFINITY)
	})

	t.Run("Intersecting a constrained cylinder", func(t *testing.T) {
		testcases := []OriginDirectionCount{
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 1.5, 0), datatypes.Vector(0.1, 1, 0)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 3, -5), datatypes.Vector(0, 0, 1)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 0, -5), datatypes.Vector(0, 0, 1)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 2, -5), datatypes.Vector(0, 0, 1)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 1, -5), datatypes.Vector(0, 0, 1)}, 0},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 1.5, -2), datatypes.Vector(0, 0, 1)}, 2}}

		c := GetCylinder()
		c.Min = 1
		c.Max = 2

		for _, testcase := range testcases {
			testcase.ray.Direction = testcase.ray.Direction.Normalize()
			xs := Intersect(c, testcase.ray)
			assertVal(t, float64(len(xs)), testcase.count)
		}
	})

	t.Run("The default closed value for a cylinder", func(t *testing.T) {
		c := GetCylinder()
		if c.Closed {
			t.Fatal("Expected cylinder default to be not closed")
		}

	})

	t.Run("Intersecting the caps of a closed cylinder", func(t *testing.T) {
		testcases := []OriginDirectionCount{
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 3, 0), datatypes.Vector(0, -1, 0)}, 2},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 3, -2), datatypes.Vector(0, -1, 2)}, 2},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 4, -2), datatypes.Vector(0, -1, 1)}, 2},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, 0, -2), datatypes.Vector(0, 1, 2)}, 2},
			OriginDirectionCount{datatypes.Ray{datatypes.Point(0, -1, -2), datatypes.Vector(0, 1, 1)}, 2}}

		c := GetCylinder()
		c.Min = 1
		c.Max = 2
		c.Closed = true

		for _, testcase := range testcases {
			testcase.ray.Direction = testcase.ray.Direction.Normalize()
			xs := Intersect(c, testcase.ray)
			assertVal(t, float64(len(xs)), testcase.count)
		}
	})

	t.Run("The normal vector on a cylinder's end caps", func(t *testing.T) {
		testcases := []datatypes.Ray{
			datatypes.Ray{datatypes.Point(0, 1, 0), datatypes.Vector(0, -1, 0)},
			datatypes.Ray{datatypes.Point(0.5, 1, 0), datatypes.Vector(0, -1, 0)},
			datatypes.Ray{datatypes.Point(0, 1, 0.5), datatypes.Vector(0, -1, 0)},
			datatypes.Ray{datatypes.Point(0, 2, 0), datatypes.Vector(0, 1, 0)},
			datatypes.Ray{datatypes.Point(0.5, 2, 0), datatypes.Vector(0, 1, 0)},
			datatypes.Ray{datatypes.Point(0, 2, 0.5), datatypes.Vector(0, 1, 0)}}

		c := GetCylinder()
		c.Min = 1
		c.Max = 2
		c.Closed = true

		for _, ray := range testcases {
			normal := Normal(c, ray.Origin)
			if !datatypes.IsTupleEqual(normal, ray.Direction) {
				t.Fatalf("Actual normal vector %v != expected normal vector %v", normal, ray.Direction)
			}
		}
	})

}
