package raytracing

import (
	"errors"
	"github.com/seantur/ray_tracer_challenge/matrices"
	"github.com/seantur/ray_tracer_challenge/tuples"
	"math"
)

type Sphere struct {
	transform matrices.Matrix
}

func (s *Sphere) Init() {
	s.transform = matrices.GetIdentity()
}

func (s *Sphere) SetTransform(m matrices.Matrix) {
	s.transform = m
}

func (s *Sphere) Intersect(r Ray) []Intersection {
	tInv, _ := s.transform.Inverse()
	r = r.Transform(tInv)

	sphereToRay := tuples.Subtract(r.Origin, tuples.Point(0, 0, 0))

	a := tuples.Dot(r.Direction, r.Direction)
	b := 2 * tuples.Dot(r.Direction, sphereToRay)
	c := tuples.Dot(sphereToRay, sphereToRay) - 1

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
	t      float64
	object *Sphere
}

func Hit(intersections []Intersection) (Intersection, error) {

	var hit_val float64
	hit_intersection := intersections[0]

	for _, intersection := range intersections {
		if intersection.t > 0 && (hit_val == 0 || intersection.t < hit_val) {
			hit_intersection = intersection
			hit_val = intersection.t
		}
	}

	if hit_val == 0 {
		return hit_intersection, errors.New("did not find hit")
	}

	return hit_intersection, nil
}
