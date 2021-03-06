package datatypes

import (
	"math"
	"testing"
)

func TestTuples(t *testing.T) {

	t.Run("A tuple with w=1.0 is a Point", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 1.0}

		AssertVal(t, a.X, 4.3)
		AssertVal(t, a.Y, -4.2)
		AssertVal(t, a.Z, 3.1)
		AssertVal(t, a.W, 1.0)
		AssertString(t, a.Label(), "point")
	})

	t.Run("A tuple with w=0.0 is a Vector", func(t *testing.T) {
		a := Tuple{4.3, -4.2, 3.1, 0.0}

		AssertVal(t, a.X, 4.3)
		AssertVal(t, a.Y, -4.2)
		AssertVal(t, a.Z, 3.1)
		AssertVal(t, a.W, 0.0)
		AssertString(t, a.Label(), "vector")
	})

	t.Run("Point() creates a tuple with w=1", func(t *testing.T) {
		p := Point(4, -4, 3)

		AssertVal(t, p.X, 4)
		AssertVal(t, p.Y, -4)
		AssertVal(t, p.Z, 3)
		AssertVal(t, p.W, 1)
	})

	t.Run("Vector() creates a tuple with w=0", func(t *testing.T) {
		v := Vector(4, -4, 3)

		AssertVal(t, v.X, 4)
		AssertVal(t, v.Y, -4)
		AssertVal(t, v.Z, 3)
		AssertVal(t, v.W, 0)
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

		AssertTupleEqual(t, Add(a1, a2), Tuple{1, 1, 6, 1})
	})

	t.Run("test subtracting Points", func(t *testing.T) {
		p1 := Point(3, 2, 1)
		p2 := Point(5, 6, 7)

		AssertTupleEqual(t, Subtract(p1, p2), Vector(-2, -4, -6))
	})

	t.Run("test subtract Vector from a Point", func(t *testing.T) {
		p := Point(3, 2, 1)
		v := Vector(5, 6, 7)

		AssertTupleEqual(t, Subtract(p, v), Point(-2, -4, -6))
	})

	t.Run("test subtracting Vectors", func(t *testing.T) {
		v1 := Vector(3, 2, 1)
		v2 := Vector(5, 6, 7)

		AssertTupleEqual(t, Subtract(v1, v2), Vector(-2, -4, -6))
	})

	t.Run("test subtracting from the zero Vector", func(t *testing.T) {
		v1 := Vector(0, 0, 0)
		v2 := Vector(1, 2, 3)

		AssertTupleEqual(t, Subtract(v1, v2), Vector(-1, -2, -3))
	})

	t.Run("test negating a Vector", func(t *testing.T) {
		v := Vector(1, 2, 3)

		AssertTupleEqual(t, v.Negate(), Vector(-1, -2, -3))
	})

	t.Run("multiple a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		AssertTupleEqual(t, a.Multiply(3.5), Tuple{3.5, -7, 10.5, -14})
	})

	t.Run("multiple a tuple by a fraction", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		AssertTupleEqual(t, a.Multiply(0.5), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("Divide a tuple by a scalar", func(t *testing.T) {
		a := Tuple{1, -2, 3, -4}

		AssertTupleEqual(t, a.Divide(2), Tuple{0.5, -1, 1.5, -2})
	})

	t.Run("Compute the Magnitude of Vector(1, 0, 0)", func(t *testing.T) {
		v := Vector(1, 0, 0)

		AssertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(0, 1, 0)", func(t *testing.T) {
		v := Vector(0, 1, 0)

		AssertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(0, 0, 1)", func(t *testing.T) {
		v := Vector(0, 0, 1)

		AssertVal(t, v.Magnitude(), 1)
	})

	t.Run("Compute the Magnitude of Vector(1, 2, 3)", func(t *testing.T) {
		v := Vector(1, 2, 3)

		AssertVal(t, v.Magnitude(), math.Sqrt(14))
	})

	t.Run("Compute the Magnitude of Vector(-1, -2, -3)", func(t *testing.T) {
		v := Vector(-1, -2, -3)

		AssertVal(t, v.Magnitude(), math.Sqrt(14))
	})

	t.Run("Normalize Vector(4, 0, 0) gives the unit vector", func(t *testing.T) {
		v := Vector(4, 0, 0)

		AssertTupleEqual(t, v.Normalize(), Vector(1, 0, 0))
	})

	t.Run("Normalize Vector(1, 2, 2)", func(t *testing.T) {
		v := Vector(1, 2, 3)

		AssertTupleEqual(t, v.Normalize(), Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)))
	})

	t.Run("Magnitude of a Normalized Vector is 1", func(t *testing.T) {
		v := Vector(1, 2, 3)
		norm := v.Normalize()

		AssertVal(t, norm.Magnitude(), 1)
	})

	t.Run("dot product of 2 tuples", func(t *testing.T) {
		a := Vector(1, 2, 3)
		b := Vector(2, 3, 4)

		AssertVal(t, Dot(a, b), 20)
	})

	t.Run("cross product of 2 tuples", func(t *testing.T) {
		a := Vector(1, 2, 3)
		b := Vector(2, 3, 4)

		AssertTupleEqual(t, Cross(a, b), Vector(-1, 2, -1))
		AssertTupleEqual(t, Cross(b, a), Vector(1, -2, 1))
	})

	t.Run("Reflecting a vector approaching at 45deg", func(t *testing.T) {
		v := Vector(1, -1, 0)
		n := Vector(0, 1, 0)
		r := v.Reflect(n)

		AssertTupleEqual(t, r, Vector(1, 1, 0))
	})

	t.Run("Reflecting a vector off a slanted surface", func(t *testing.T) {
		v := Vector(0, -1, 0)
		n := Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

		r := v.Reflect(n)

		AssertTupleEqual(t, r, Vector(1, 0, 0))
	})
}
