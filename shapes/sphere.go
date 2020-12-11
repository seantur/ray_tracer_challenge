package shapes

import (
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"github.com/seantur/ray_tracer_challenge/raytracing"
	"math"
)

type Sphere struct {
	Transform datatypes.Matrix
	raytracing.Material
}

func GetSphere() Shape {
	s := Sphere{}
	s.Transform = datatypes.GetIdentity()
	s.Material = raytracing.GetMaterial()

	return &s
}

func (s *Sphere) GetMaterial() raytracing.Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m raytracing.Material) {
	s.Material = m
}

func (s *Sphere) GetTransform() datatypes.Matrix {
	return s.Transform
}

func (s *Sphere) SetTransform(m datatypes.Matrix) {
	s.Transform = m
}

func (s *Sphere) Intersect(r datatypes.Ray) (xs []Intersection) {
	sphereToRay := datatypes.Subtract(r.Origin, datatypes.Point(0, 0, 0))

	a := datatypes.Dot(r.Direction, r.Direction)
	b := 2 * datatypes.Dot(r.Direction, sphereToRay)
	c := datatypes.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return
	}

	xs = []Intersection{
		Intersection{(-b - math.Sqrt(discriminant)) / (2 * a), s},
		Intersection{(-b + math.Sqrt(discriminant)) / (2 * a), s}}

	return
}

func (s *Sphere) Normal(obj_p datatypes.Tuple) datatypes.Tuple {
	return datatypes.Subtract(obj_p, datatypes.Point(0, 0, 0))
}

func GetGlassSphere() Shape {
	s := GetSphere()

	material := s.GetMaterial()
	material.Transparency = 1.0
	material.RefractiveIndex = 1.5
	s.SetMaterial(material)

	return s
}
