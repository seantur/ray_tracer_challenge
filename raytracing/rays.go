package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/matrices"
	"github.com/seantur/ray_tracer_challenge/tuples"
)

type Ray struct {
	Origin    tuples.Tuple
	Direction tuples.Tuple
}

func (r *Ray) Position(t float64) tuples.Tuple {
	return tuples.Add(r.Origin, r.Direction.Multiply(t))
}

func (r *Ray) Transform(m matrices.Matrix) Ray {
	origin := matrices.TupleMultiply(m, r.Origin)
	direction := matrices.TupleMultiply(m, r.Direction)
	return Ray{Origin: origin, Direction: direction}
}
