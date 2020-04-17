package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
)

type Ray struct {
	Origin    datatypes.Tuple
	Direction datatypes.Tuple
}

func (r *Ray) Position(t float64) datatypes.Tuple {
	return datatypes.Add(r.Origin, r.Direction.Multiply(t))
}

func (r *Ray) Transform(m datatypes.Matrix) Ray {
	origin := datatypes.TupleMultiply(m, r.Origin)
	direction := datatypes.TupleMultiply(m, r.Direction)
	return Ray{Origin: origin, Direction: direction}
}
