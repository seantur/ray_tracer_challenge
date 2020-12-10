package datatypes

import "math"

const EPSILON = 0.00001
const INFINITY = 10e6

func IsClose(a float64, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	} else {
		return false
	}
}
