package datatypes

import (
	"errors"
)

type Matrix struct {
	Row  int
	Col  int
	Vals []float64
}

func (m *Matrix) Init() {
	m.Vals = make([]float64, m.Row*m.Col)
}

func GetEmptyMatrix(row, col int) Matrix {
	m := Matrix{Row: row, Col: col}
	m.Init()
	return m
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
	// Transpose, so initial column/row instead of row/column
	M := GetEmptyMatrix(m.Col, m.Row)

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

	M := GetEmptyMatrix(m.Row-1, m.Col-1)
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
		return Matrix{}, errors.New("trying to invert an non-invertible matrix")
	}

	M := GetEmptyMatrix(m.Row, m.Col)

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			M.Set(j, i, GetCofactor(*m, i, j)/det)
		}
	}

	return M, nil
}

func (m *Matrix) equal(m2 Matrix) bool {

	if m.Row != m2.Row || m.Col != m2.Col {
		return false
	}

	for i := 0; i < len(m.Vals); i++ {
		if !IsClose(m.Vals[i], m2.Vals[i]) {
			return false
		}
	}
	return true
}

func Multiply(matrices ...Matrix) Matrix {
	//TODO throw error if any dimensions don't match

	M := Matrix{Row: matrices[0].Row, Col: matrices[len(matrices)-1].Col}
	M.Init()

	var val float64
	for index := 0; index < len(matrices)-1; index++ {
		for i := 0; i < matrices[index].Row; i++ {
			for j := 0; j < matrices[index+1].Col; j++ {
				val = 0
				for k := 0; k < matrices[index].Col; k++ {
					m1Val, _ := matrices[index].At(i, k)
					m2Val, _ := matrices[index+1].At(k, j)

					val += m1Val * m2Val
				}

				M.Set(i, j, val)
			}
		}
	}

	return M
}

func TupleMultiply(m Matrix, t Tuple) Tuple {
	tMat := Matrix{Row: 4, Col: 1, Vals: []float64{t.X, t.Y, t.Z, t.W}}

	out := Multiply(m, tMat)

	x, _ := out.At(0, 0)
	y, _ := out.At(1, 0)
	z, _ := out.At(2, 0)
	w, _ := out.At(3, 0)

	return Tuple{X: x, Y: y, Z: z, W: w}
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
