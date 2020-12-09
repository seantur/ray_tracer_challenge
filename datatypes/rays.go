package datatypes

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func (r *Ray) Position(t float64) Tuple {
	return Add(r.Origin, r.Direction.Multiply(t))
}

func (r *Ray) Transform(m Matrix) Ray {
	origin := TupleMultiply(m, r.Origin)
	direction := TupleMultiply(m, r.Direction)
	return Ray{Origin: origin, Direction: direction}
}
