package raytracing

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type Pattern interface {
	At(point datatypes.Tuple) RGB
	SetTransform(m datatypes.Matrix)
	GetTransform() datatypes.Matrix
}

type Stripe struct {
	A, B      RGB
	Transform datatypes.Matrix
}

func GetStripe(a, b RGB) Pattern {
	s := Stripe{a, b, datatypes.GetIdentity()}
	return &s
}

func (s *Stripe) At(point datatypes.Tuple) RGB {

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
	A, B      RGB
	Transform datatypes.Matrix
}

func GetGradient(a, b RGB) Pattern {
	g := Gradient{a, b, datatypes.GetIdentity()}
	return &g
}

func (g *Gradient) At(point datatypes.Tuple) RGB {
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

type Ring struct {
	A, B      RGB
	Transform datatypes.Matrix
}

func GetRing(a, b RGB) Pattern {
	r := Ring{a, b, datatypes.GetIdentity()}
	return &r
}

func (r *Ring) At(point datatypes.Tuple) RGB {
	if math.Mod(math.Floor(math.Sqrt(math.Pow(point.X, 2)+math.Pow(point.Z, 2))), 2) == 0 {
		return r.A
	}
	return r.B
}

func (r *Ring) SetTransform(m datatypes.Matrix) {
	r.Transform = m
}

func (r *Ring) GetTransform() datatypes.Matrix {
	return r.Transform
}

type Checkers struct {
	A, B      RGB
	Transform datatypes.Matrix
}

func GetCheckers(a, b RGB) Pattern {
	r := Checkers{a, b, datatypes.GetIdentity()}
	return &r
}

func (c *Checkers) At(point datatypes.Tuple) RGB {

	if math.Mod(math.Floor(point.X)+math.Floor(point.Y)+math.Floor(point.Z), 2) == 0 {
		return c.A
	}
	return c.B
}

func (c *Checkers) SetTransform(m datatypes.Matrix) {
	c.Transform = m
}

func (c *Checkers) GetTransform() datatypes.Matrix {
	return c.Transform
}

type TestPat struct {
	Transform datatypes.Matrix
}

func GetTestPat() Pattern {
	testpat := TestPat{}
	testpat.Transform = datatypes.GetIdentity()
	return &testpat
}

func (t *TestPat) At(point datatypes.Tuple) RGB {
	return RGB{Red: point.X, Green: point.Y, Blue: point.Z}
}

func (t *TestPat) SetTransform(m datatypes.Matrix) {
	t.Transform = m
}

func (t *TestPat) GetTransform() datatypes.Matrix {
	return t.Transform
}
