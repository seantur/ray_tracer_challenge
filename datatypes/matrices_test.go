package datatypes

import (
	"testing"
)

func TestMatrices(t *testing.T) {

	t.Run("construct and inspect 4x4 matrix", func(t *testing.T) {
		M := Matrix{4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}}

		val, _ := M.At(0, 0)
		AssertVal(t, val, 1)

		val, _ = M.At(0, 3)
		AssertVal(t, val, 4)

		val, _ = M.At(1, 0)
		AssertVal(t, val, 5.5)

		val, _ = M.At(1, 2)
		AssertVal(t, val, 7.5)

		val, _ = M.At(2, 2)
		AssertVal(t, val, 11)

		val, _ = M.At(3, 0)
		AssertVal(t, val, 13.5)

		val, _ = M.At(3, 2)
		AssertVal(t, val, 15.5)
	})

	t.Run("construct a 2x2 matrix", func(t *testing.T) {
		M := Matrix{2, 2, []float64{-3, 5, 1, -2}}

		val, _ := M.At(0, 0)
		AssertVal(t, val, -3)

		val, _ = M.At(0, 1)
		AssertVal(t, val, 5)

		val, _ = M.At(1, 0)
		AssertVal(t, val, 1)

		val, _ = M.At(1, 1)
		AssertVal(t, val, -2)

	})

	t.Run("construct a 3x3 matrix", func(t *testing.T) {
		M := Matrix{3, 3, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}}

		val, _ := M.At(0, 0)
		AssertVal(t, val, -3)

		val, _ = M.At(1, 1)
		AssertVal(t, val, -2)

		val, _ = M.At(2, 2)
		AssertVal(t, val, 1)
	})

	t.Run("test matrix equal with same matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}

		AssertMatrixEqual(t, A, B)
	})

	t.Run("test matrix not-equal with different matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}}

		if A.equal(B) {
			t.Error("Assert matrices are not equal")
		}
	})

	t.Run("multiply 2 4x4 matrices", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}}
		B := Matrix{4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}}

		want := Matrix{4, 4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}}

		AssertMatrixEqual(t, Multiply(A, B), want)
	})

	t.Run("matrix multiplied by a tuple", func(t *testing.T) {
		A := Matrix{4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}}
		tuple := Tuple{X: 1, Y: 2, Z: 3, W: 1}

		want := Tuple{X: 18, Y: 24, Z: 33, W: 1}
		got := TupleMultiply(A, tuple)

		AssertTupleEqual(t, want, got)
	})

	t.Run("matrix multiplied by identity matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32}}
		I := GetIdentity()

		AssertMatrixEqual(t, Multiply(A, I), A)
	})

	t.Run("identity matrix multiplied by a tuple", func(t *testing.T) {
		I := GetIdentity()
		tuple := Tuple{X: 1, Y: 2, Z: 3, W: 1}

		got := TupleMultiply(I, tuple)

		AssertTupleEqual(t, got, tuple)
	})

	t.Run("transpose a matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}}
		want := Matrix{4, 4, []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}}

		got := A.Transpose()

		AssertMatrixEqual(t, want, got)
	})

	t.Run("the transpose of an the identity matrix is the identity matrix", func(t *testing.T) {
		I := GetIdentity()
		It := I.Transpose()

		AssertMatrixEqual(t, I, It)

	})

	t.Run("get determinant of 2x2 matrix", func(t *testing.T) {
		A := Matrix{2, 2, []float64{1, 5, -3, 2}}

		got := GetDeterminant(A)

		AssertVal(t, 17, got)

	})

	t.Run("get 2x2 submatrix of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{1, 5, 0, -3, 2, 7, 0, 6, -3}}
		want := Matrix{2, 2, []float64{-3, 2, 0, 6}}

		got := A.Submatrix(0, 2)

		AssertMatrixEqual(t, want, got)
	})

	t.Run("get 3x3 submatrix of 4x4 matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1}}
		want := Matrix{3, 3, []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1}}

		got := A.Submatrix(2, 1)

		AssertMatrixEqual(t, want, got)
	})

	t.Run("calculate the minor of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}}
		B := A.Submatrix(1, 0)

		det := GetDeterminant(B)
		minor := GetMinor(A, 1, 0)

		AssertVal(t, det, minor)
		AssertVal(t, minor, 25)
	})

	t.Run("calculate the cofactor of 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}}

		minor := GetMinor(A, 0, 0)
		AssertVal(t, minor, -12)

		cofactor := GetCofactor(A, 0, 0)
		AssertVal(t, cofactor, -12)

		minor = GetMinor(A, 1, 0)
		AssertVal(t, minor, 25)

		cofactor = GetCofactor(A, 1, 0)
		AssertVal(t, cofactor, -25)

	})

	t.Run("calculate the determinate of a 3x3 matrix", func(t *testing.T) {
		A := Matrix{3, 3, []float64{1, 2, 6, -5, 8, -4, 2, 6, 4}}

		AssertVal(t, GetCofactor(A, 0, 0), 56)
		AssertVal(t, GetCofactor(A, 0, 1), 12)
		AssertVal(t, GetCofactor(A, 0, 2), -46)
		AssertVal(t, GetDeterminant(A), -196)
	})

	t.Run("calculate the dterminate of a 4x4 matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9}}

		AssertVal(t, GetCofactor(A, 0, 0), 690)
		AssertVal(t, GetCofactor(A, 0, 1), 447)
		AssertVal(t, GetCofactor(A, 0, 2), 210)
		AssertVal(t, GetCofactor(A, 0, 3), 51)
		AssertVal(t, GetDeterminant(A), -4071)

	})

	t.Run("testing an invertible matrix for invertibility", func(t *testing.T) {
		A := Matrix{4, 4, []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6}}

		AssertVal(t, GetDeterminant(A), -2120)

		if !A.isInvertible() {
			t.Error("A is not invertable, but it should be")
		}
	})

	t.Run("testing an uninvertible matrix for invertibility", func(t *testing.T) {
		A := Matrix{4, 4, []float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0}}

		AssertVal(t, GetDeterminant(A), 0)
		if A.isInvertible() {
			t.Error("A is invertable, but it should not be")
		}
	})

	t.Run("calculate the inverse of a matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{
			-5, 2, 6, -8,
			1, -5, 1, 8,
			7, 7, -6, -7,
			1, -3, 7, 4}}
		B := Matrix{4, 4, []float64{
			0.21805, 0.45113, 0.24060, -0.04511,
			-0.80827, -1.45677, -0.44361, 0.52068,
			-0.07895, -0.22368, -0.05263, 0.19737,
			-0.52256, -0.81391, -0.30075, 0.30639}}

		AssertVal(t, GetDeterminant(A), 532)
		AssertVal(t, GetCofactor(A, 2, 3), -160)
		AssertVal(t, GetCofactor(A, 3, 2), 105)

		Ainverse, _ := A.Inverse()

		AssertMatrixEqual(t, Ainverse, B)

	})

	t.Run("calculate the inverse of another matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{
			8, -5, 9, 2,
			7, 5, 6, 1,
			-6, 0, 9, 6,
			-3, 0, -9, -4}}

		B := Matrix{4, 4, []float64{
			-0.15385, -0.15385, -0.28205, -0.53846,
			-0.07692, 0.12308, 0.02564, 0.03077,
			0.35897, 0.35897, 0.43590, 0.92308,
			-0.69231, -0.69231, -0.76923, -1.92308}}

		Ainverse, _ := A.Inverse()

		AssertMatrixEqual(t, Ainverse, B)

	})

	t.Run("calculate the inverse of a third matrix", func(t *testing.T) {
		A := Matrix{4, 4, []float64{
			9, 3, 0, 9,
			-5, -2, -6, -3,
			-4, 9, 6, 4,
			-7, 6, 6, 2}}

		B := Matrix{4, 4, []float64{
			-0.04074, -0.07778, 0.14444, -0.22222,
			-0.07778, 0.03333, 0.36667, -0.33333,
			-0.02901, -0.14630, -0.10926, 0.12963,
			0.17778, 0.06667, -0.26667, 0.33333}}

		Ainverse, _ := A.Inverse()

		AssertMatrixEqual(t, Ainverse, B)

	})

	t.Run("multiply a product by its inverse", func(t *testing.T) {
		A := Matrix{4, 4, []float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1}}
		B := Matrix{4, 4, []float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5}}

		C := Multiply(A, B)

		Binverse, _ := B.Inverse()

		AssertMatrixEqual(t, Multiply(C, Binverse), A)

	})

	t.Run("multiply multiple identity matrices", func(t *testing.T) {
		a := GetIdentity()
		b := GetIdentity()
		c := GetIdentity()

		d := Multiply(a, b, c)
		AssertMatrixEqual(t, d, GetIdentity())
	})

	t.Run("multiply multiple matrices", func(t *testing.T) {
		a := Matrix{4, 4, []float64{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 0, 0, 0, 0}}
		b := Matrix{4, 4, []float64{4, 3, 2, 1, 4, 3, 2, 1, 4, 3, 2, 1, 0, 0, 0, 0}}
		c := Matrix{4, 4, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}}

		m := Multiply(a, b)
		m = Multiply(m, c)

		M := Multiply(a, b, c)

		AssertMatrixEqual(t, m, M)
	})
}
