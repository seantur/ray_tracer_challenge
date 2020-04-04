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
		tuple := tuples.Tuple{X: 1, Y: 2, Z: 3, W: 1}

		want := tuples.Tuple{X: 18, Y: 24, Z: 33, W: 1}
		got := TupleMultiply(A, tuple)

		if !tuples.Equal(want, got) {
			t.Error("did not get expected tuple")
		}
	})

	t.Run("matrix multiplied by identity matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32}}
		I := GetIdentity()

		assertMatrixEqual(t, Multiply(A, I), A)
	})

	t.Run("identity matrix multiplied by a tuple", func(t *testing.T) {
		I := GetIdentity()
		tuple := tuples.Tuple{X: 1, Y: 2, Z: 3, W: 1}

		got := TupleMultiply(I, tuple)

		if !tuples.Equal(tuple, got) {
			t.Error("did not get expected tuple")
		}
	})

	t.Run("transpose a matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}}
		want := Matrix{4, 4, []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}}

		got := A.Transpose()

		assertMatrixEqual(t, want, got)
	})

	t.Run("the transpose of an the identity matrix is the identity matrix", func(t *testing.T) {
		I := GetIdentity()
		It := I.Transpose()

		assertMatrixEqual(t, I, It)

	})

	t.Run("get determinant of 2x2 matrix", func(t *testing.T) {
		A := Matrix{2, 2, []float64{1, 5, -3, 2}}

		got := GetDeterminant(A)

		assertVal(t, 17, got)

	})

	t.Run("get 2x2 submatrix of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{1, 5, 0, -3, 2, 7, 0, 6, -3}}
		want := Matrix{2, 2, []float64{-3, 2, 0, 6}}

		got := A.Submatrix(0, 2)

		assertMatrixEqual(t, want, got)
	})

	t.Run("get 3x3 submatrix of 4x4 matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1}}
		want := Matrix{3, 3, []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1}}

		got := A.Submatrix(2, 1)

		assertMatrixEqual(t, want, got)
	})

	t.Run("calculate the minor of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}}
		B := A.Submatrix(1, 0)

		det := GetDeterminant(B)
		minor := GetMinor(A, 1, 0)

		assertVal(t, det, minor)
		assertVal(t, minor, 25)
	})

	t.Run("calculate the cofactor of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}}

		minor := GetMinor(A, 0, 0)
		assertVal(t, minor, -12)

		cofactor := GetCofactor(A, 0, 0)
		assertVal(t, cofactor, -12)

		minor = GetMinor(A, 1, 0)
		assertVal(t, minor, 25)

		cofactor = GetCofactor(A, 1, 0)
		assertVal(t, cofactor, -25)

	})

}
