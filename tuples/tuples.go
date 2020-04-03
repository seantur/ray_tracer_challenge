package data_types

import "math"

const EPSILON = 0.00001

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (t *Tuple) label() string {
	if t.w == 1.0 {
		return "point"
	} else if t.w == 0.0 {
		return "vector"
	} else {
		return "unknown"
	}
}

func (t *Tuple) negate() Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

func (t *Tuple) multiply(a float64) Tuple {
	return Tuple{t.x * a, t.y * a, t.z * a, t.w * a}
}

func (t *Tuple) divide(a float64) Tuple {
	return Tuple{t.x / a, t.y / a, t.z / a, t.w / a}
}

func (t *Tuple) magnitude() float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2))
}

func (t *Tuple) normalize() Tuple {
	magnitude := t.magnitude()
	return Tuple{t.x / magnitude, t.y / magnitude, t.z / magnitude, t.w / magnitude}
}

func IsClose(a float64, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	} else {
		return false
	}
}

func point(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1}
}

func vector(x float64, y float64, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0}
}

func Equal(a Tuple, b Tuple) bool {
	return IsClose(a.x, b.x) && IsClose(a.y, b.y) && IsClose(a.z, b.z) && IsClose(a.w, b.w)
}

func Add(a Tuple, b Tuple) Tuple {
	return Tuple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func Subtract(a Tuple, b Tuple) Tuple {
	return Tuple{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
}

func Dot(a Tuple, b Tuple) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z + a.w*b.w
}

func Cross(a Tuple, b Tuple) Tuple {
	return vector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}
