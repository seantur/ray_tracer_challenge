package tuples

import "math"

const EPSILON = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t *Tuple) Label() string {
	if t.W == 1.0 {
		return "point"
	} else if t.W == 0.0 {
		return "vector"
	} else {
		return "unknown"
	}
}

func (t *Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t *Tuple) Multiply(a float64) Tuple {
	return Tuple{t.X * a, t.Y * a, t.Z * a, t.W * a}
}

func (t *Tuple) Divide(a float64) Tuple {
	return Tuple{t.X / a, t.Y / a, t.Z / a, t.W / a}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))
}

func (t *Tuple) Normalize() Tuple {
	magnitude := t.Magnitude()
	return Tuple{t.X / magnitude, t.Y / magnitude, t.Z / magnitude, t.W / magnitude}
}

func IsClose(a float64, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	} else {
		return false
	}
}

func Point(X float64, Y float64, Z float64) Tuple {
	return Tuple{X: X, Y: Y, Z: Z, W: 1}
}

func Vector(X float64, Y float64, Z float64) Tuple {
	return Tuple{X: X, Y: Y, Z: Z, W: 0}
}

func Equal(a Tuple, b Tuple) bool {
	return IsClose(a.X, b.X) && IsClose(a.Y, b.Y) && IsClose(a.Z, b.Z) && IsClose(a.W, b.W)
}

func Add(a Tuple, b Tuple) Tuple {
	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

func Subtract(a Tuple, b Tuple) Tuple {
	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

func Dot(a Tuple, b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func Cross(a Tuple, b Tuple) Tuple {
	return Vector(a.Y*b.Z-a.Z*b.Y, a.Z*b.X-a.X*b.Z, a.X*b.Y-a.Y*b.X)
}
