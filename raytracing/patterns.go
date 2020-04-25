package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type Pattern interface {
	At(point datatypes.Tuple) Color
	SetTransform(m datatypes.Matrix)
	GetTransform() datatypes.Matrix
}

func AtObj(p Pattern, shape Shape, point datatypes.Tuple) Color {
	objTransform := shape.GetTransform()
	patternTransform := p.GetTransform()

	patternTransformInv, _ := patternTransform.Inverse()
	objTransformInv, _ := objTransform.Inverse()

	objPoint := datatypes.TupleMultiply(objTransformInv, point)
	patternPoint := datatypes.TupleMultiply(patternTransformInv, objPoint)

	return p.At(patternPoint)
}

type Stripe struct {
	A, B      Color
	Transform datatypes.Matrix
}

func GetStripe(a, b Color) Pattern {
	s := Stripe{a, b, datatypes.GetIdentity()}
	return &s
}

func (s *Stripe) At(point datatypes.Tuple) Color {

	if math.Mod(math.Floor(point.X), 2) == 0 {
		return s.A
	}

	return s.B
}

func (s *Stripe) SetTransform(m datatypes.Matrix) {
	s.Transform = m
}

func (s *Stripe) GetTransform() datatypes.Matrix {
	return s.Transform
}

type Gradient struct {
	A, B      Color
	Transform datatypes.Matrix
}

func GetGradient(a, b Color) Pattern {
	g := Gradient{a, b, datatypes.GetIdentity()}
	return &g
}

func (g *Gradient) At(point datatypes.Tuple) Color {
	distance := Subtract(g.B, g.A)
	frac := point.X - math.Floor(point.X)

	return Add(g.A, distance.Multiply(frac))
}

func (g *Gradient) SetTransform(m datatypes.Matrix) {
	g.Transform = m
}

func (g *Gradient) GetTransform() datatypes.Matrix {
	return g.Transform
}
