package tuples

import (
	"math"
	"testing"
)

func TestTuples(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	assertString := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertTupleEqual := func(t *testing.T, got Tuple, want Tuple) {
		t.Helper()
		if !Equal(got, want) {
			t.Error("wanted equal tuples are not equal")
		}
	}

	t.Run("A tuple with w=1.0 is a Point", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 1.0}

		assertVal(t, a.X, 4.3)
		assertVal(t, a.Y, -4.2)
		assertVal(t, a.Z, 3.1)
		assertVal(t, a.W, 1.0)
		assertString(t, a.Label(), "point")
	})

	t.Run("A tuple with w=0.0 is a Vector", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 0.0}

		assertVal(t, a.X, 4.3)
		assertVal(t, a.Y, -4.2)
		assertVal(t, a.Z, 3.1)
		assertVal(t, a.W, 0.0)
		assertString(t, a.Label(), "vector")
	})

	t.Run("Point() creates a tuple with w=1", func(t *testing.T) {
		p := Point(4, -4, 3)

		assertVal(t, p.X, 4)
		assertVal(t, p.Y, -4)
		assertVal(t, p.Z, 3)
		assertVal(t, p.W, 1)
	})

	t.Run("Vector() creates a tuple with w=0", func(t *testing.T) {
		v := Vector(4, -4, 3)

		assertVal(t, v.X, 4)
		assertVal(t, v.Y, -4)
		assertVal(t, v.Z, 3)
		assertVal(t, v.W, 0)
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

	t.Run("test subtracting Points", func(t *testing.T) {
		p1 := Point(3, 2, 1)
		p2 := Point(5, 6, 7)

		assertTupleEqual(t, Subtract(p1, p2), Vector(-2, -4, -6))
	})

	t.Run("test subtract Vector from a Point", func(t *testing.T) {
		p := Point(3, 2, 1)
		v := Vector(5, 6, 7)

		assertTupleEqual(t, Subtract(p, v), Point(-2, -4, -6))
	})

	t.Run("test subtracting Vectors", func(t *testing.T) {
		v1 := Vector(3, 2, 1)
		v2 := Vector(5, 6, 7)

		assertTupleEqual(t, Subtract(v1, v2), Vector(-2, -4, -6))
	})

	t.Run("test subtracting from the zero Vector", func(t *testing.T) {
		v1 := Vector(0, 0, 0)
		v2 := Vector(1, 2, 3)

		assertTupleEqual(t, Subtract(v1, v2), Vector(-1, -2, -3))
	})

	t.Run("test negating a Vector", func(t *testing.T) {
		v := Vector(1, 2, 3)

		assertTupleEqual(t, v.Negate(), Vector(-1, -2, -3))
	})

	t.Run("multiple a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.Multiply(3.5), Tuple{3.5, -7, 10.5, -14})
	})

	t.Run("multiple a tuple by a fraction", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.Multiply(0.5), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("Divide a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		assertTupleEqual(t, a.Divide(2), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("Compute the Magnitude of Vector(1, 0, 0)", func(t *testing.T) {
		v := Vector(1, 0, 0)

		assertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(0, 1, 0)", func(t *testing.T) {
		v := Vector(0, 1, 0)

		assertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(0, 0, 1)", func(t *testing.T) {
		v := Vector(0, 0, 1)

		assertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(1, 2, 3)", func(t *testing.T) {
		v := Vector(1, 2, 3)

		assertVal(t, v.Magnitude(), math.Sqrt(14))
	})

	t.Run("Compute the Magnitude of Vector(-1, -2, -3)", func(t *testing.T) {
		v := Vector(-1, -2, -3)

		assertVal(t, v.Magnitude(), math.Sqrt(14))
	})

	t.Run("Normalize Vector(4, 0, 0) gives the unit vector", func(t *testing.T) {
		v := Vector(4, 0, 0)

		assertTupleEqual(t, v.Normalize(), Vector(1, 0, 0))
	})

	t.Run("Normalize Vector(1, 2, 2)", func(t *testing.T) {
		v := Vector(1, 2, 3)

		assertTupleEqual(t, v.Normalize(), Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)))
	})

	t.Run("Magnitude of a Normalized Vector is 1", func(t *testing.T) {
		v := Vector(1, 2, 3)
		norm := v.Normalize()

		assertVal(t, norm.Magnitude(), 1)
	})

	t.Run("dot product of 2 tuples", func(t *testing.T) {
		a := Vector(1, 2, 3)
		b := Vector(2, 3, 4)

		assertVal(t, Dot(a, b), 20)
	})

	t.Run("cross product of 2 tuples", func(t *testing.T) {
		a := Vector(1, 2, 3)
		b := Vector(2, 3, 4)

		assertTupleEqual(t, Cross(a, b), Vector(-1, 2, -1))
		assertTupleEqual(t, Cross(b, a), Vector(1, -2, 1))
	})

	t.Run("Reflecting a vector approaching at 45deg", func(t *testing.T) {
		v := Vector(1, -1, 0)
		n := Vector(0, 1, 0)
		r := v.Reflect(n)

		assertTupleEqual(t, r, Vector(1, 1, 0))
	})

	t.Run("Reflecting a vector off a slanted surface", func(t *testing.T) {
		v := Vector(0, -1, 0)
		n := Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

		r := v.Reflect(n)

		assertTupleEqual(t, r, Vector(1, 0, 0))
	})
}
