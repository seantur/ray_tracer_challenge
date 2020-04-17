package raytracing

import (
	"errors"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type Sphere struct {
	transform datatypes.Matrix
	Material  Material
}

func GetSphere() Sphere {
	s := Sphere{}
	s.transform = datatypes.GetIdentity()
	s.Material = GetMaterial()

	return s
}

func (s *Sphere) SetTransform(m datatypes.Matrix) {
	s.transform = m
}

func (s *Sphere) GetNormal(world_p datatypes.Tuple) datatypes.Tuple {
	s_transform_inv, _ := s.transform.Inverse()

	obj_p := datatypes.TupleMultiply(s_transform_inv, world_p)
	obj_normal := datatypes.Subtract(obj_p, datatypes.Point(0, 0, 0))
	world_normal := datatypes.TupleMultiply(s_transform_inv.Transpose(), obj_normal)
	world_normal.W = 0

	return world_normal.Normalize()
}

func (s *Sphere) Intersect(r Ray) []Intersection {
	tInv, _ := s.transform.Inverse()
	r = r.Transform(tInv)

	sphereToRay := datatypes.Subtract(r.Origin, datatypes.Point(0, 0, 0))

	a := datatypes.Dot(r.Direction, r.Direction)
	b := 2 * datatypes.Dot(r.Direction, sphereToRay)
	c := datatypes.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return make([]Intersection, 0)
	}

	xs := []Intersection{
		Intersection{(-b - math.Sqrt(discriminant)) / (2 * a), s},
		Intersection{(-b + math.Sqrt(discriminant)) / (2 * a), s}}

	return xs
}

type Intersection struct {
	T      float64
	Object *Sphere
}

func Hit(intersections []Intersection) (Intersection, error) {

	var hit_val float64
	hit_intersection := intersections[0]

	for _, intersection := range intersections {
		if intersection.T > 0 && (hit_val == 0 || intersection.T < hit_val) {
			hit_intersection = intersection
			hit_val = intersection.T
		}
	}

	if hit_val == 0 {
		return hit_intersection, errors.New("did not find hit")
	}

	return hit_intersection, nil
}
