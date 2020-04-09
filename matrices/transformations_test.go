package matrices

import (
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
	"testing"
)

func TestTransformations(t *testing.T) {

	assertTupleEqual := func(t *testing.T, got tuples.Tuple, want tuples.Tuple) {
		t.Helper()
		if !tuples.Equal(got, want) {
			t.Error("wanted equal tuples are not equal")
		}
	}

	t.Run("Multiplying by a translation matrix", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		p := tuples.Point(-3, 4, 5)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(2, 1, 7))

	})

	t.Run("Multiplying by the inverse of translation matrix", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		p := tuples.Point(-3, 4, 5)

		T, _ = T.Inverse()

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(-8, 7, 3))

	})

	t.Run("Translation does not affect vectors", func(t *testing.T) {
		T := GetTranslation(5, -3, 2)
		v := tuples.Vector(-3, 4, 5)

		assertTupleEqual(t, TupleMultiply(T, v), v)

	})

	t.Run("A scaling matrix applied to a point", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		p := tuples.Point(-4, 6, 8)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(-8, 18, 32))

	})

	t.Run("A scaling matrix applied to a vector", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		v := tuples.Vector(-4, 6, 8)

		assertTupleEqual(t, TupleMultiply(T, v), tuples.Vector(-8, 18, 32))

	})

	t.Run("Multiplying by the inverse of a scaling matrix", func(t *testing.T) {
		T := GetScaling(2, 3, 4)
		T, _ = T.Inverse()

		v := tuples.Vector(-4, 6, 8)
		assertTupleEqual(t, TupleMultiply(T, v), tuples.Vector(-2, 2, 2))

	})

	t.Run("Reflection is scaling by a negative value", func(t *testing.T) {
		T := GetScaling(-1, 1, 1)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(-2, 3, 4))

	})

	t.Run("Rotating a point around the x axis", func(t *testing.T) {
		p := tuples.Point(0, 1, 0)
		Rot45 := GetRotationX(math.Pi / 4)
		Rot90 := GetRotationX(math.Pi / 2)

		assertTupleEqual(t, TupleMultiply(Rot45, p), tuples.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2))
		assertTupleEqual(t, TupleMultiply(Rot90, p), tuples.Point(0, 0, 1))

	})

	t.Run("The inverse of an x-rotation rotates in the opposite direction", func(t *testing.T) {
		p := tuples.Point(0, 1, 0)
		Rot45 := GetRotationX(math.Pi / 4)
		inv, _ := Rot45.Inverse()

		assertTupleEqual(t, TupleMultiply(inv, p), tuples.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	})

	t.Run("Rotating a point around the y axis", func(t *testing.T) {
		p := tuples.Point(0, 0, 1)
		Rot45 := GetRotationY(math.Pi / 4)
		Rot90 := GetRotationY(math.Pi / 2)

		assertTupleEqual(t, TupleMultiply(Rot45, p), tuples.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2))
		assertTupleEqual(t, TupleMultiply(Rot90, p), tuples.Point(1, 0, 0))

	})

	t.Run("Rotating a point around the z axis", func(t *testing.T) {
		p := tuples.Point(0, 1, 0)
		Rot45 := GetRotationZ(math.Pi / 4)
		Rot90 := GetRotationZ(math.Pi / 2)

		assertTupleEqual(t, TupleMultiply(Rot45, p), tuples.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0))
		assertTupleEqual(t, TupleMultiply(Rot90, p), tuples.Point(-1, 0, 0))

	})

	t.Run("A shearing transformation moves x in proportion to y", func(t *testing.T) {
		T := GetShearing(1, 0, 0, 0, 0, 0)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(5, 3, 4))

	})

	t.Run("A shearing transformation moves x in proportion to z", func(t *testing.T) {
		T := GetShearing(0, 1, 0, 0, 0, 0)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(6, 3, 4))

	})

	t.Run("A shearing transformation moves y in proportion to x", func(t *testing.T) {
		T := GetShearing(0, 0, 1, 0, 0, 0)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(2, 5, 4))

	})

	t.Run("A shearing transformation moves y in proportion to z", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 1, 0, 0)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(2, 7, 4))

	})

	t.Run("A shearing transformation moves z in proportion to x", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 0, 1, 0)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(2, 3, 6))

	})

	t.Run("A shearing transformation moves z in proportion to y", func(t *testing.T) {
		T := GetShearing(0, 0, 0, 0, 0, 1)
		p := tuples.Point(2, 3, 4)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(2, 3, 7))

	})

	t.Run("Individual transformations are applied in sequence", func(t *testing.T) {
		p := tuples.Point(1, 0, 1)
		A := GetRotationX(math.Pi / 2)
		B := GetScaling(5, 5, 5)
		C := GetTranslation(10, 5, 7)

		p2 := TupleMultiply(A, p)
		assertTupleEqual(t, p2, tuples.Point(1, -1, 0))

		p3 := TupleMultiply(B, p2)
		assertTupleEqual(t, p3, tuples.Point(5, -5, 0))

		p4 := TupleMultiply(C, p3)
		assertTupleEqual(t, p4, tuples.Point(15, 0, 7))

	})

	t.Run("Chained transformations must be applied in reverse order", func(t *testing.T) {
		p := tuples.Point(1, 0, 1)
		A := GetRotationX(math.Pi / 2)
		B := GetScaling(5, 5, 5)
		C := GetTranslation(10, 5, 7)

		T := Multiply(C, B)
		T = Multiply(T, A)

		assertTupleEqual(t, TupleMultiply(T, p), tuples.Point(15, 0, 7))
	})

}
