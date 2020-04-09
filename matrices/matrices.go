package matrices

import (
	"errors"
	"github.com/seantur/ray_tracer_challenge/tuples"
)

const EPSILON = 0.00001

type Matrix struct {
	Row  int
	Col  int
	Vals []float64
}

func (m *Matrix) Init() {
	m.Vals = make([]float64, m.Row*m.Col)
}

func (m *Matrix) At(row int, col int) (float64, error) {

	if row > m.Row || col > m.Col {
		return 0.0, errors.New("trying access out of bounds")
	}

	return m.Vals[row*m.Col+col], nil
}

func (m *Matrix) Set(row int, col int, val float64) error {

	if row > m.Row || col > m.Col {
		return errors.New("trying access out of bounds")
	}

	m.Vals[row*m.Col+col] = val

	return nil
}

func (m *Matrix) Transpose() Matrix {

	M := Matrix{Row: m.Col, Col: m.Row}
	M.Init()

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			val, _ := m.At(i, j)
			M.Set(j, i, val)
		}
	}

	return M
}

// Return the matrix with row/col removed
func (m *Matrix) Submatrix(row int, col int) Matrix {

	M := Matrix{Row: m.Row - 1, Col: m.Col - 1}
	M.Init()

	var Mi, Mj int

	for i := 0; i < m.Row; i++ {
		switch {
		case i < row:
			Mi = i
		case i == row:
			continue
		case i > row:
			Mi = i - 1
		}
		for j := 0; j < m.Col; j++ {
			switch {
			case j < col:
				Mj = j
			case j == col:
				continue
			case j > col:
				Mj = j - 1
			}
			val, _ := m.At(i, j)
			M.Set(Mi, Mj, val)
		}
	}
	return M
}

func (m *Matrix) isInvertible() bool {
	if GetDeterminant(*m) == 0 {
		return false
	}
	return true

}

func (m *Matrix) Inverse() (Matrix, error) {

	det := GetDeterminant(*m)

	if !m.isInvertible() {
		return Matrix{}, errors.New("trying to invert and non-invertible matrix")
	}

	M := Matrix{Row: m.Row, Col: m.Col}
	M.Init()

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			M.Set(j, i, GetCofactor(*m, i, j)/det)
		}
	}

	return M, nil
}

func Equal(m1 Matrix, m2 Matrix) bool {

	if m1.Row != m2.Row || m1.Col != m2.Col {
		return false
	}

	for i := 0; i < len(m1.Vals); i++ {
		if (m1.Vals[i] - m2.Vals[i]) > EPSILON {
			return false
		}
	}
	return true
}

func Multiply(m1 Matrix, m2 Matrix) Matrix {

	//if m1.Col != m2.Row {
	// TODO throw an error if you can't multiply
	//}

	M := Matrix{Row: m1.Row, Col: m2.Col}
	M.Init()

	for i := 0; i < m1.Row; i++ {
		for j := 0; j < m2.Col; j++ {
			var val float64
			//val = 0

			for k := 0; k < m1.Col; k++ {
				m1Val, _ := m1.At(i, k)
				m2Val, _ := m2.At(k, j)

				val += m1Val * m2Val
			}

			M.Set(i, j, val)
		}
	}

	return M
}

func TupleMultiply(m Matrix, t tuples.Tuple) tuples.Tuple {
	tMat := Matrix{Row: 4, Col: 1, Vals: []float64{t.X, t.Y, t.Z, t.W}}

	out := Multiply(m, tMat)

	x, _ := out.At(0, 0)
	y, _ := out.At(1, 0)
	z, _ := out.At(2, 0)
	w, _ := out.At(3, 0)

	return tuples.Tuple{X: x, Y: y, Z: z, W: w}
}

func GetIdentity() Matrix {
	return Matrix{Row: 4, Col: 4, Vals: []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}}
}

func GetDeterminant(m Matrix) float64 {
	if m.Row == 2 && m.Col == 2 {
		a, _ := m.At(0, 0)
		b, _ := m.At(0, 1)
		c, _ := m.At(1, 0)
		d, _ := m.At(1, 1)

		return a*d - b*c
	}

	var det float64

	for i := 0; i < m.Col; i++ {
		val, _ := m.At(0, i)
		det += val * GetCofactor(m, 0, i)
	}

	return det
}

// TODO Enforce 3x3
func GetMinor(m Matrix, row int, col int) float64 {
	sub := m.Submatrix(row, col)
	return GetDeterminant(sub)
}

// TODO Enforce 3x3
func GetCofactor(m Matrix, row int, col int) float64 {
	minor := GetMinor(m, row, col)

	if ((row + col) % 2) == 0 {
		return minor
	} else {
		return -minor
	}
}
