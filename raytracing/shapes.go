package raytracing

import (
	"errors"
	"github.com/seantur/ray_tracer_challenge/datatypes"
	"math"
)

type Sphere struct {
	Transform datatypes.Matrix
	Material
}

type Intersection struct {
	T      float64
	Object *Sphere
}

type Computation struct {
	T                               float64
	Object                          *Sphere
	Point, Eyev, Normalv, OverPoint datatypes.Tuple
	IsInside                        bool
}

func (i *Intersection) PrepareComputations(r Ray) Computation {
	c := Computation{}

	c.T = i.T
	c.Object = i.Object
	c.Point = r.Position(c.T)
	c.Eyev = r.Direction.Negate()
	c.Normalv = c.Object.GetNormal(c.Point)

	if datatypes.Dot(c.Normalv, c.Eyev) < 0 {
		c.IsInside = true
		c.Normalv = c.Normalv.Negate()
	} else {
		c.IsInside = false
	}

	c.OverPoint = datatypes.Add(c.Point, c.Normalv.Multiply(datatypes.EPSILON))

	return c
}

// ByT implements sort.Interface for []Intersection based on the T field
type ByT []Intersection

func (in ByT) Len() int           { return len(in) }
func (in ByT) Swap(i, j int)      { in[i], in[j] = in[j], in[i] }
func (in ByT) Less(i, j int) bool { return in[i].T < in[j].T }

func GetSphere() Sphere {
	s := Sphere{}
	s.Transform = datatypes.GetIdentity()
	s.Material = GetMaterial()

	return s
}

func (s *Sphere) SetTransform(m datatypes.Matrix) {
	s.Transform = m
}

func (s *Sphere) GetNormal(world_p datatypes.Tuple) datatypes.Tuple {
	s_transform_inv, _ := s.Transform.Inverse()

	obj_p := datatypes.TupleMultiply(s_transform_inv, world_p)
	obj_normal := datatypes.Subtract(obj_p, datatypes.Point(0, 0, 0))
	world_normal := datatypes.TupleMultiply(s_transform_inv.Transpose(), obj_normal)
	world_normal.W = 0

	return world_normal.Normalize()
}

func (s *Sphere) Intersect(r Ray) []Intersection {
	tInv, _ := s.Transform.Inverse()
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

func Hit(intersections []Intersection) (Intersection, error) {

	if len(intersections) == 0 {
		return Intersection{}, errors.New("did not find hit")
	}

	var hit_val float64
	hit_intersection := intersections[0]

	for _, intersection := range intersections {
		if intersection.T > 0 && (hit_val == 0 || intersection.T < hit_val) {
			hit_intersection = intersection
			hit_val = intersection.T
		}
	}

	if hit_val == 0 {
		return Intersection{}, errors.New("did not find hit")
	}

	return hit_intersection, nil
}
