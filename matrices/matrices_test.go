package matrices

import (
	"github.com/seantur/ray_tracer_challenge/tuples"
	"testing"
)

func TestMatrices(t *testing.T) {

	assertVal := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	assertMatrixEqual := func(t *testing.T, got Matrix, want Matrix) {
		t.Helper()
		if !Equal(got, want) {
			t.Error("wanted equal matrices are not equal")
		}
	}

	t.Run("construct and inspect 4x4 matrix", func(t *testing.T) {
		M := Matrix{4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}}

		val, _ := M.At(0, 0)
		assertVal(t, val, 1)

		val, _ = M.At(0, 3)
		assertVal(t, val, 4)

		val, _ = M.At(1, 0)
		assertVal(t, val, 5.5)

		val, _ = M.At(1, 2)
		assertVal(t, val, 7.5)

		val, _ = M.At(2, 2)
		assertVal(t, val, 11)

		val, _ = M.At(3, 0)
		assertVal(t, val, 13.5)

		val, _ = M.At(3, 2)
		assertVal(t, val, 15.5)
	})

	t.Run("construct a 2x2 matrix", func(t *testing.T) {
		M := Matrix{2, 2, []float64{-3, 5, 1, -2}}

		val, _ := M.At(0, 0)
		assertVal(t, val, -3)

		val, _ = M.At(0, 1)
		assertVal(t, val, 5)

		val, _ = M.At(1, 0)
		assertVal(t, val, 1)

		val, _ = M.At(1, 1)
		assertVal(t, val, -2)

	})

	t.Run("construct a 3x3 matrix", func(t *testing.T) {
		M := Matrix{3, 3, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}}

		val, _ := M.At(0, 0)
		assertVal(t, val, -3)

		val, _ = M.At(1, 1)
		assertVal(t, val, -2)

		val, _ = M.At(2, 2)
		assertVal(t, val, 1)
	})

	t.Run("test matrix equal with same matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}

		if !Equal(A, B) {
			t.Error("Assert matrices are not equal")
		}
	})

	t.Run("test matrix not-equal with different matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}}

		if Equal(A, B) {
			t.Error("Assert matrices are not equal")
		}
	})

	t.Run("multiply 2 4x4 matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}}

		want := Matrix{4, 4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}}

		assertMatrixEqual(t, Multiply(A, B), want)
	})

	t.Run("matrix multiplied by a tuple", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}}
		tuple := tuples.Tuple{1, 2, 3, 1}

		want := tuples.Tuple{18, 24, 33, 1}
		got := TupleMultiply(A, tuple)

		if !tuples.Equal(want, got) {
			t.Error("did not get expected tuple")
		}
	})

}
