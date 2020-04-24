package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type StripePattern struct {
	A, B      Color
	Transform datatypes.Matrix
}

func GetStripe(a, b Color) StripePattern {
	return StripePattern{a, b, datatypes.GetIdentity()}
}

func (s *StripePattern) At(point datatypes.Tuple) Color {

	if math.Mod(math.Floor(point.X), 2) == 0 {
		return s.A
	}

	return s.B
}

func (s *StripePattern) SetTransform(m datatypes.Matrix) {
	s.Transform = m
}

func (s *StripePattern) AtObj(shape Shape, point datatypes.Tuple) Color {
	objTransform := shape.GetTransform()
	patternTransform, _ := s.Transform.Inverse()
	objTransformInv, _ := objTransform.Inverse()

	objPoint := datatypes.TupleMultiply(objTransformInv, point)
	patternPoint := datatypes.TupleMultiply(patternTransform, objPoint)

	return s.At(patternPoint)
}
