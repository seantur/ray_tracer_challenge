package datatypes

import "testing"

func AssertVal(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func AssertMatrixEqual(t *testing.T, got Matrix, want Matrix) {
	if !got.equal(want) {
		t.Error("wanted equal matrices are not equal")
	}
}

func AssertTupleEqual(t *testing.T, got Tuple, want Tuple) {
	if !got.equal(want) {
		t.Error("wanted equal tuples are not equal")
	}
}

func AssertString(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
