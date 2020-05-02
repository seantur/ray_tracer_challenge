package datatypes

import (
	"math"
)

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

func ViewTransform(from Tuple, to Tuple, up Tuple) Matrix {
	forward := Subtract(to, from)
	forward = forward.Normalize()

	left := Cross(forward, up.Normalize())

	true_up := Cross(left, forward)

	orientation := Matrix{4, 4, []float64{
		left.X, left.Y, left.Z, 0,
		true_up.X, true_up.Y, true_up.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1}}

	return Multiply(orientation, GetTranslation(-from.X, -from.Y, -from.Z))
}

func GetTransform(transform_matrices ...Matrix) Matrix {
	for i := len(transform_matrices)/2 - 1; i >= 0; i-- {
		opp := len(transform_matrices) - 1 - i
		transform_matrices[i], transform_matrices[opp] = transform_matrices[opp], transform_matrices[i]
	}

	return Multiply(transform_matrices...)
}
