package matrices

import "math"

func GetTranslation(x float64, y float64, z float64) Matrix {
	return Matrix{4, 4, []float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1}}
}

func GetScaling(x float64, y float64, z float64) Matrix {
	return Matrix{4, 4, []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1}}
}

func GetRotationX(rads float64) Matrix {
	return Matrix{4, 4, []float64{
		1, 0, 0, 0,
		0, math.Cos(rads), -math.Sin(rads), 0,
		0, math.Sin(rads), math.Cos(rads), 0,
		0, 0, 0, 1}}
}

func GetRotationY(rads float64) Matrix {
	return Matrix{4, 4, []float64{
		math.Cos(rads), 0, math.Sin(rads), 0,
		0, 1, 0, 0,
		-math.Sin(rads), 0, math.Cos(rads), 0,
		0, 0, 0, 1}}
}

func GetRotationZ(rads float64) Matrix {
	return Matrix{4, 4, []float64{
		math.Cos(rads), -math.Sin(rads), 0, 0,
		math.Sin(rads), math.Cos(rads), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}}
}

func GetShearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	return Matrix{4, 4, []float64{
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1}}
}
