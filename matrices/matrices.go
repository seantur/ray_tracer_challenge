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

	//}

	M := Matrix{Row: m1.Row, Col: m2.Col}
	M.Init()
	var val float64

	for i := 0; i < m1.Row; i++ {
		for j := 0; j < m2.Col; j++ {
			val = 0

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
