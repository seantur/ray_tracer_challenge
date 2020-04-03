package data_types

import (
	"math"
	"testing"
)

func TestTuples(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	assertString := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertTupleEqual := func(t *testing.T, got Tuple, want Tuple) {
		if !Equal(got, want) {
			t.Error("wanted equal tuples are not equal")
		}
	}

	t.Run("A tuple with w=1.0 is a point", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 1.0}

		assertVal(t, a.x, 4.3)
		assertVal(t, a.y, -4.2)
		assertVal(t, a.z, 3.1)
		assertVal(t, a.w, 1.0)
		assertString(t, a.label(), "point")
	})

	t.Run("A tuple with w=0.0 is a vector", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 0.0}

		assertVal(t, a.x, 4.3)
		assertVal(t, a.y, -4.2)
		assertVal(t, a.z, 3.1)
		assertVal(t, a.w, 0.0)
		assertString(t, a.label(), "vector")
	})

	t.Run("point() creates a tuple with w=1", func(t *testing.T) {
		p := point(4, -4, 3)

		assertVal(t, p.x, 4)
		assertVal(t, p.y, -4)
		assertVal(t, p.z, 3)
		assertVal(t, p.w, 1)
	})

	t.Run("vector() creates a tuple with w=0", func(t *testing.T) {
		v := vector(4, -4, 3)

		assertVal(t, v.x, 4)
		assertVal(t, v.y, -4)
		assertVal(t, v.z, 3)
		assertVal(t, v.w, 0)
	})

	t.Run("test IsClose", func(t *testing.T) {
		if IsClose(5, 5+(2*EPSILON)) {
			t.Errorf("want false, got true")
		}

		if !IsClose(5, 5+(EPSILON/2)) {
			t.Errorf("want true, got false")
		}
	})

	t.Run("test adding tuples", func(t *testing.T) {
		a1 := Tuple{3, -2, 5, 1}
		a2 := Tuple{-2, 3, 1, 0}

		assertTupleEqual(t, Add(a1, a2), Tuple{1, 1, 6, 1})
	})

	t.Run("test subtracting points", func(t *testing.T) {
		p1 := point(3, 2, 1)
		p2 := point(5, 6, 7)

		assertTupleEqual(t, Subtract(p1, p2), vector(-2, -4, -6))
	})

	t.Run("test subtract vector from a point", func(t *testing.T) {
		p := point(3, 2, 1)
		v := vector(5, 6, 7)

		assertTupleEqual(t, Subtract(p, v), point(-2, -4, -6))
	})

	t.Run("test subtracting vectors", func(t *testing.T) {
		v1 := vector(3, 2, 1)
		v2 := vector(5, 6, 7)

		assertTupleEqual(t, Subtract(v1, v2), vector(-2, -4, -6))
	})

	t.Run("test subtracting from the zero vector", func(t *testing.T) {
		v1 := vector(0, 0, 0)
		v2 := vector(1, 2, 3)

		assertTupleEqual(t, Subtract(v1, v2), vector(-1, -2, -3))
	})

	t.Run("test negating a vector", func(t *testing.T) {
		v := vector(1, 2, 3)

		assertTupleEqual(t, v.negate(), vector(-1, -2, -3))
	})

	t.Run("multiple a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.multiply(3.5), Tuple{3.5, -7, 10.5, -14})
	})

	t.Run("multiple a tuple by a fraction", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.multiply(0.5), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("divide a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.divide(2), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("Compute the magnitude of vector(1, 0, 0)", func(t *testing.T) {
		v := vector(1, 0, 0)

		assertVal(t, v.magnitude(), 1)
	})

	t.Run("Compute the magnitude of vector(0, 1, 0)", func(t *testing.T) {
		v := vector(0, 1, 0)

		assertVal(t, v.magnitude(), 1)
	})

	t.Run("Compute the magnitude of vector(0, 0, 1)", func(t *testing.T) {
		v := vector(0, 0, 1)

		assertVal(t, v.magnitude(), 1)
	})

	t.Run("Compute the magnitude of vector(1, 2, 3)", func(t *testing.T) {
		v := vector(1, 2, 3)

		assertVal(t, v.magnitude(), math.Sqrt(14))
	})

	t.Run("Compute the magnitude of vector(-1, -2, -3)", func(t *testing.T) {
		v := vector(-1, -2, -3)

		assertVal(t, v.magnitude(), math.Sqrt(14))
	})

	t.Run("Normalize vector(4, 0, 0) gives the unit vector", func(t *testing.T) {
		v := vector(4, 0, 0)

		assertTupleEqual(t, v.normalize(), vector(1, 0, 0))
	})

	t.Run("Normalize vector(1, 2, 2)", func(t *testing.T) {
		v := vector(1, 2, 3)

		assertTupleEqual(t, v.normalize(), vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)))
	})

	t.Run("magnitude of a normalized vector is 1", func(t *testing.T) {
		v := vector(1, 2, 3)
		norm := v.normalize()

		assertVal(t, norm.magnitude(), 1)
	})

	t.Run("dot product of 2 tuples", func(t *testing.T) {
		a := vector(1, 2, 3)
		b := vector(2, 3, 4)

		assertVal(t, Dot(a, b), 20)
	})

	t.Run("cross product of 2 tuples", func(t *testing.T) {
		a := vector(1, 2, 3)
		b := vector(2, 3, 4)

		assertTupleEqual(t, Cross(a, b), vector(-1, 2, -1))
		assertTupleEqual(t, Cross(b, a), vector(1, -2, 1))
	})
}
