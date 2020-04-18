package datatypes

import (
	"math"
	"testing"
)

func TestTransformations(t *testing.T) {

	t.Run("Multiplying by a translation matrix", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		p := Point(-3, 4, 5)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(2, 1, 7))

	})

	t.Run("Multiplying by the inverse of translation matrix", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		p := Point(-3, 4, 5)

		T, _ = T.Inverse()

		AssertTupleEqual(t, TupleMultiply(T, p), Point(-8, 7, 3))

	})

	t.Run("Translation does not affect vectors", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		v := Vector(-3, 4, 5)

		AssertTupleEqual(t, TupleMultiply(T, v), v)

	})

	t.Run("A scaling matrix applied to a point", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		p := Point(-4, 6, 8)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(-8, 18, 32))

	})

	t.Run("A scaling matrix applied to a vector", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		v := Vector(-4, 6, 8)

		AssertTupleEqual(t, TupleMultiply(T, v), Vector(-8, 18, 32))

	})

	t.Run("Multiplying by the inverse of a scaling matrix", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		T, _ = T.Inverse()

		v := Vector(-4, 6, 8)
		AssertTupleEqual(t, TupleMultiply(T, v), Vector(-2, 2, 2))

	})

	t.Run("Reflection is scaling by a negative value", func(t *testing.T) {
		T := GetScaling(-1, 1, 1)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(-2, 3, 4))

	})

	t.Run("Rotating a point around the x axis", func(t *testing.T) {
		p := Point(0, 1, 0)
		Rot45 := GetRotationX(math.Pi / 4)
		Rot90 := GetRotationX(math.Pi / 2)

		AssertTupleEqual(t, TupleMultiply(Rot45, p), Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2))
		AssertTupleEqual(t, TupleMultiply(Rot90, p), Point(0, 0, 1))

	})

	t.Run("The inverse of an x-rotation rotates in the opposite direction", func(t *testing.T) {
		p := Point(0, 1, 0)
		Rot45 := GetRotationX(math.Pi / 4)
		inv, _ := Rot45.Inverse()

		AssertTupleEqual(t, TupleMultiply(inv, p), Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	})

	t.Run("Rotating a point around the y axis", func(t *testing.T) {
		p := Point(0, 0, 1)
		Rot45 := GetRotationY(math.Pi / 4)
		Rot90 := GetRotationY(math.Pi / 2)

		AssertTupleEqual(t, TupleMultiply(Rot45, p), Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2))
		AssertTupleEqual(t, TupleMultiply(Rot90, p), Point(1, 0, 0))

	})

	t.Run("Rotating a point around the z axis", func(t *testing.T) {
		p := Point(0, 1, 0)
		Rot45 := GetRotationZ(math.Pi / 4)
		Rot90 := GetRotationZ(math.Pi / 2)

		AssertTupleEqual(t, TupleMultiply(Rot45, p), Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0))
		AssertTupleEqual(t, TupleMultiply(Rot90, p), Point(-1, 0, 0))

	})

	t.Run("A shearing transformation moves x in proportion to y", func(t *testing.T) {
		T := GetShearing(1, 0, 0, 0, 0, 0)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(5, 3, 4))

	})

	t.Run("A shearing transformation moves x in proportion to z", func(t *testing.T) {
		T := GetShearing(0, 1, 0, 0, 0, 0)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(6, 3, 4))

	})

	t.Run("A shearing transformation moves y in proportion to x", func(t *testing.T) {
		T := GetShearing(0, 0, 1, 0, 0, 0)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(2, 5, 4))

	})

	t.Run("A shearing transformation moves y in proportion to z", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 1, 0, 0)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(2, 7, 4))

	})

	t.Run("A shearing transformation moves z in proportion to x", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 0, 1, 0)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(2, 3, 6))

	})

	t.Run("A shearing transformation moves z in proportion to y", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 0, 0, 1)
		p := Point(2, 3, 4)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(2, 3, 7))

	})

	t.Run("Individual transformations are applied in sequence", func(t *testing.T) {
		p := Point(1, 0, 1)
		A := GetRotationX(math.Pi / 2)
		B := GetScaling(5, 5, 5)
		C := GetTranslation(10, 5, 7)

		p2 := TupleMultiply(A, p)
		AssertTupleEqual(t, p2, Point(1, -1, 0))

		p3 := TupleMultiply(B, p2)
		AssertTupleEqual(t, p3, Point(5, -5, 0))

		p4 := TupleMultiply(C, p3)
		AssertTupleEqual(t, p4, Point(15, 0, 7))

	})

	t.Run("Chained transformations must be applied in reverse order", func(t *testing.T) {
		p := Point(1, 0, 1)
		A := GetRotationX(math.Pi / 2)
		B := GetScaling(5, 5, 5)
		C := GetTranslation(10, 5, 7)

		T := Multiply(C, B)
		T = Multiply(T, A)

		AssertTupleEqual(t, TupleMultiply(T, p), Point(15, 0, 7))
	})

}